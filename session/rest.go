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

package session

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.ibm.com/riethm/gopherlayer/sl"
)

func doRestRequest(sess *Session, service string, method string, args []interface{}, options *sl.Options, pResult interface{}) error {
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
		sess,
		path,
		restMethod,
		bytes.NewBuffer(parameters),
		options)

	if err != nil {
		return fmt.Errorf("Error during HTTP request: %s", err)
	}

	if code < 200 || code > 299 {
		e := sl.Error{StatusCode: code}

		err = json.Unmarshal(resp, &e)

		// If unparseable, wrap the json error
		if err != nil {
			e.Wrapped = err
			e.Message = err.Error()
		}

		return e
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
		if err != nil {
			return sl.Error{Message: err.Error(), Wrapped: err}
		}
		return nil
	case "*bool":
		*pResult.(*bool), err = strconv.ParseBool(string(resp))
		if err != nil {
			return sl.Error{Message: err.Error(), Wrapped: err}
		}
	case "float64":
		*pResult.(*float64), err = strconv.ParseFloat(string(resp), 64)
		if err != nil {
			return sl.Error{Message: err.Error(), Wrapped: err}
		}
	case "string":
		*pResult.(*string) = string(resp)
		return nil
	}

	// Must be a json representation of one of the many softlayer datatypes
	err = json.Unmarshal(resp, pResult)
	if err != nil {
		return sl.Error{Message: err.Error(), Wrapped: err}
	}
	return nil
}

func encodeQuery(opts *sl.Options) string {
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

func makeHttpRequest(session *Session, path string, requestType string, requestBody *bytes.Buffer, options *sl.Options) ([]byte, int, error) {
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
