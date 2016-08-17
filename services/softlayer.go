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
	"runtime"
	"strings"

	"github.ibm.com/riethm/gopherlayer.git/session"
	"github.ibm.com/riethm/gopherlayer.git/sl"
)

func invokeMethod(args []interface{}, session *session.Session, options *sl.Options, pResult interface{}) error {
	// Get the caller information, which gives us the service and method name
	pc, _, _, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	segments := strings.Split(f.Name(), ".")
	service, method := segments[len(segments)-2], segments[len(segments)-1]

	// All services that have methods have the SoftLayer_ prefix
	service = "SoftLayer_" + service

	// camelCase the method name
	method = strings.ToLower(string(method[0])) + method[1:]

	return session.DoRequest(service, method, args, options, pResult)
}
