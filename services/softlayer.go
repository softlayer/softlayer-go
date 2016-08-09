/**
 * Copyright 2016 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"github.ibm.com/riethm/gopherlayer/config"
)

const DEFAULT_ENDPOINT = "https://api.softlayer.com/rest/v3"

type Options struct {
	ObjectId     *int
	ObjectMask   string
	ObjectFilter string
	ResultLimit  *int
	StartOffset  *int
}

func (r *Options) Id(id int) *Options {
	r.ObjectId = &id
	return r
}

func (r *Options) Mask(mask string) *Options {
	r.ObjectMask = mask
	return r
}

func (r *Options) Filter(filter string) *Options {
	r.ObjectFilter = filter
	return r
}

func (r *Options) Limit(limit int) *Options {
	r.ResultLimit = &limit
	return r
}

func (r *Options) Offset(offset int) *Options {
	r.StartOffset = &offset
	return r
}

type Session struct {
	UserName string
	ApiKey   string
	Endpoint string
	Debug    bool
}

func NewSession(args ...interface{}) Session {
	keys := map[string]int{"username": 0, "api_key": 1, "endpoint_url": 2}
	values := []string{"", "", DEFAULT_ENDPOINT}

	if endpoint_url := os.Getenv("SOFTLAYER_ENDPOINT_URL"); endpoint_url != "" {
		values[keys["endpoint_url"]] = endpoint_url
	}

	for i := 0; i < len(args); i++ {
		values[i] = args[i].(string)
	}

	// Default to the environment variables
	envFallback("SOFTLAYER_USERNAME", &values[keys["username"]])
	envFallback("SOFTLAYER_API_KEY", &values[keys["api_key"]])

	// Read ~/.softlayer for configuration
	u, err := user.Current()
	if err != nil {
		panic("session: Could not determine current user.")
	}

	configPath := fmt.Sprintf("%s/.softlayer", u.HomeDir)
	if _, err = os.Stat(configPath); !os.IsNotExist(err) {
		// config file exists
		file, err := config.LoadFile(configPath)
		if err != nil {
			log.Println(fmt.Sprintf("[WARN] session: Could not parse %s : %s", configPath, err))
		} else {
			for k, v := range keys {
				value, ok := file.Get("softlayer", k)
				if ok && values[v] == "" {
					values[v] = value
				}
			}
		}
	}

	return Session{
		UserName: values[keys["username"]],
		ApiKey:   values[keys["api_key"]],
		Endpoint: values[keys["endpoint_url"]],
	}
}

func (r *Session) String() string {
	return "Username: " + r.UserName +
		", ApiKey: " + r.ApiKey +
		", Endpoint: " + r.Endpoint
}

func (r *Session) DoRequest(service string, method string, args []interface{}, options *Options, pResult interface{}) error {
	restMethod := httpMethod(method, args)

	// Parse any method parameters and determine the HTTP method
	var parameters []byte
	if len(args) > 0 {
		// parse the parameters
		parameters, _ = json.Marshal(
			map[string]interface{}{
				"parameters": args,
			})
	}

	// Start building the request path
	path := service

	if options.ObjectId != nil {
		path = path + "/" + strconv.Itoa(*options.ObjectId)
	}

	// omit the API method name if the method represents one of the basic REST methods
	if method != "getObject" && method != "deleteObject" && method != "createObject" &&
		method != "createObjects" && method != "editObject" && method != "editObjects" {
		path = path + "/" + method
	}

	path = path + ".json"

	resp, code, err := makeHttpRequest(
		r,
		path,
		restMethod,
		bytes.NewBuffer(parameters),
		options)

	if err != nil {
		return fmt.Errorf("Error during HTTP request: %s", err)
	}

	if code < 200 || code > 299 {
		return fmt.Errorf("An HTTP Error was returned: %d: %s", code, string(resp))
	}

	returnType := reflect.TypeOf(pResult).String()

	// Some APIs that normally return a collection, omit the []'s when the API returns a single value
	if strings.Index(returnType, "[]") == 1 && strings.Index(string(resp), "[") != 0 {
		resp = []byte("[" + string(resp) + "]")
	}

	// At this point, all that's left to do is parse the return value to the appropriate type, and return
	// any parse errors (or nil if successful)

	switch returnType {
	case "[]byte":
		pResult = &resp
		return nil
	case "*void":
		return nil
	case "*uint":
		*pResult.(*int), err = strconv.Atoi(string(resp))
		return err
	case "*bool":
		*pResult.(*bool), err = strconv.ParseBool(string(resp))
		return err
	case "float64":
		*pResult.(*float64), err = strconv.ParseFloat(string(resp), 64)
		return err
	case "string":
		*pResult.(*string) = string(resp)
		return nil
	}

	// Must be a json representation of one of the many softlayer datatypes
	return json.Unmarshal(resp, pResult)
}

func invokeMethod(args []interface{}, session *Session, options *Options, pResult interface{}) error {
	// Get the caller information, which gives us the service and method name
	pc, _, _, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	segments := strings.Split(f.Name(), ".")
	service, method := segments[len(segments)-2], segments[len(segments)-1]

	// The receiver has the form "(*Type)".  Strip the unnecessary characters
	service = service[2 : len(service)-1]

	// Most services need to be prefixed with "SoftLayer_"
	if service[:6] != "McAfee" {
		service = "SoftLayer_" + service
	}

	// camelCase the method name
	method = strings.ToLower(string(method[0])) + method[1:]

	return session.DoRequest(service, method, args, options, pResult)
}

func envFallback(keyName string, value *string) {
	if *value == "" {
		*value = os.Getenv(keyName)
	}
}

func encodeQuery(opts *Options) string {
	query := new(url.URL).Query()

	if opts.ObjectMask != "" {
		query.Add("objectMask", opts.ObjectMask)
	}

	if opts.ObjectFilter != "" {
		query.Add("objectFilter", opts.ObjectFilter)
	}

	// resultLimit=<offset>,<limit>
	// If offset unspecified, default to 0
	if opts.ResultLimit != nil {
		startOffset := 0
		if opts.StartOffset != nil {
			startOffset = *opts.StartOffset
		}

		query.Add("resultLimit", fmt.Sprintf("%d,%d", startOffset, *opts.ResultLimit))
	}

	return query.Encode()
}

func makeHttpRequest(session *Session, path string, requestType string, requestBody *bytes.Buffer, options *Options) ([]byte, int, error) {
	client := http.DefaultClient

	var url string
	if session.Endpoint == "" {
		url = url + DEFAULT_ENDPOINT
	} else {
		url = url + session.Endpoint
	}
	url = url + "/" + path
	req, err := http.NewRequest(requestType, url, requestBody)
	if err != nil {
		return nil, 0, err
	}

	req.SetBasicAuth(session.UserName, session.ApiKey)

	req.URL.RawQuery = encodeQuery(options)

	if session.Debug {
		log.Println("[DEBUG] Path: ", req.URL)
		log.Println("[DEBUG] Parameters: ", requestBody.String())
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, 520, err
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return responseBody, resp.StatusCode, nil
}

func httpMethod(name string, args []interface{}) string {
	if name == "deleteObject" {
		return "DELETE"
	} else if name == "editObject" || name == "editObjects" {
		return "PUT"
	} else if name == "createObject" || name == "createObjects" || len(args) > 0 {
		return "POST"
	}

	return "GET"
}
