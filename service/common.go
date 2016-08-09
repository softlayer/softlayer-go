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
	"runtime"
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
