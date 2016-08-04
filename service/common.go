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

package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

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

func httpMethod(name string) string {
	if name == "deleteObject" {
		return "DELETE"
	} else if name == "createObject" || name == "createObjects" {
		return "POST"
	} else if name == "editObject" || name == "editObjects" {
		return "PUT"
	}

	return "GET"
}

func queryString(opts *Options) string {
	sep := "?"

	var qstr string

	if opts.ObjectMask != "" {
		qstr = sep + "objectMask=" + opts.ObjectMask
		sep = "&"
	}

	if opts.ObjectFilter != "" {
		qstr = qstr + sep + "objectFilter=" + opts.ObjectFilter
		sep = "&"
	}

	// resultLimit=<offset>,<limit>
	// If offset unspecified, default to 0
	if opts.ResultLimit != nil {
		qstr = qstr + sep + "resultLimit="
		if opts.StartOffset != nil {
			qstr = qstr + strconv.Itoa(*opts.StartOffset)
		} else {
			qstr = qstr + "0"
		}
		qstr = qstr + "," + strconv.Itoa(*opts.ResultLimit)
	}

	return qstr
}

func invokeMethod(args []interface{}, session *Session, options *Options, pResult interface{}) error {

	// Static Method Call:
	// <service_name>/<method_name>.json?objectMask=<mask>&objectFilter=<>&limit=<limit>&offset=<offset>

	// Instance Method Call:
	// <service_name>/<id>/<method_name>.json?objectMask=<mask>&objectFilter=<>&limit=<limit>&offset=<offset>

	// Parameters:
	// {"parameters": [ <param>, ... ]}

	// If Parameters array is len 0, method is GET, otherwise POST

	// Service name = receiver of caller + 'SoftLayer_' (except special case of McAfee* resources)
	// Method name = caller method name camelCased
	// Id, Mask, filter, offset, limit - from entity

	// method params: map[string] interface{}

	// Get the caller information, which gives us the service and method name
	pc, _, _, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	segments := strings.Split(f.Name(), ".")
	service, apiMethod := segments[len(segments)-2], segments[len(segments)-1]

	// The receiver has the form "(*Type)".  Strip the unnecessary characters
	service = service[2 : len(service)-1]

	// Most services need to be prefixed with "SoftLayer_"
	if service[:6] != "McAfee" {
		service = "SoftLayer_" + service
	}

	// camelCase the method name
	apiMethod = strings.ToLower(string(apiMethod[0])) + apiMethod[1:]

	// Parse any method parameters and determine the HTTP method
	var parameters []byte
	var restMethod string
	if len(args) > 0 {
		// parse the parameters
		parameters, _ = json.Marshal(
			map[string]interface{}{
				"parameters": args,
			})
		restMethod = "POST"
	} else {
		restMethod = httpMethod(apiMethod)

	}

	path := service

	if options.ObjectId != nil {
		path = path + "/" + strconv.Itoa(*options.ObjectId)
	}

	if apiMethod != "getObject" && apiMethod != "deleteObject" && apiMethod != "createObject" &&
		apiMethod != "createObjects" && apiMethod != "editObject" && apiMethod != "editObjects" {
		path = path + "/" + apiMethod
	}

	path = path + ".json"

	path = path + queryString(options)

	fmt.Printf("Path: %s\n", path)
	fmt.Printf("Parameters: %s\n", parameters)

	resp, code, err := makeHttpRequest(
		session,
		path,
		restMethod,
		bytes.NewBuffer(parameters))

	if err != nil {
		return fmt.Errorf("Error during HTTP request: %s", err)
	}

	if code < 200 || code > 299 {
		return fmt.Errorf("An HTTP Error was returned: %d: %s", code, string(resp))
	}

	returnType := reflect.TypeOf(pResult).String()
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
	//case "*time.Time":
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
