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
	"testing"

	"fmt"
	"reflect"

	"github.com/jarcoal/httpmock"
	"github.com/softlayer/softlayer-go/sl"
	"github.com/softlayer/xmlrpc"

)


// structure of a single testcase
type xmltestcase struct {
	description string
	service     string
	method      string
	args        []interface{}
	options     sl.Options
	responder   httpmock.Responder
	expected    interface{}
	expectError bool
}

const xmlrpcEndpoint = "https://api.softlayer.com/xmlrpc/v3"

var xmlClient *xmlrpc.Client
var err error

var emptyArray = make([]struct{}, 0)

var xmltestcases = []xmltestcase{
	// Positive tests
	// {
	// 	description: "empty array return",
	// 	service:     "SoftLayer_Account",
	// 	method:      "getVirtualGuests",
	// 	args:        nil,
	// 	options:     sl.Options{},
	// 	responder:   httpmock.NewStringResponder(200, `<?xml version="1.0" encoding="utf-8"?><params><param><value><array><data/></array></value></param></params>`),
	// 	expected:    emptyArray,
	// 	expectError: false,
	// },
	{
		description: "Single Value",
		service:     "SoftLayer_Account",
		method:      "getOneThing",
		args:        nil,
		options:     sl.Options{},
		responder:   httpmock.NewStringResponder(200, `<?xml version="1.0" encoding="utf-8"?><params><param><value><int>100</int></value></param></params>`),
		expected:    100,
		expectError: false,
	},
}

func TestXmlRpc(t *testing.T) {
	// setup session and mock environment
	s = New()
	s.Endpoint = xmlrpcEndpoint
	
	//s.Debug = true
	httpmock.Activate()
	defer httpmock.Deactivate()

	// For each test case:
	// 1. Setup the mock environment
	// 2. Allocate an empty variable for the response, based on the 'expected' response
	// 3. Call DoRequest
	// 4. Check result matches expected
	// 5. Reset the mock
	for _, tc := range xmltestcases {
		setup1(tc)

		pResult := reflect.New(reflect.TypeOf(tc.expected)).Interface()

		fmt.Printf("Test [%s]: ", tc.description)

		err := s.DoRequest(tc.service, tc.method, tc.args, &tc.options, pResult)
		fmt.Printf("\n\tpResult: %v\n", pResult)

		// Report results
		switch {

		// Positive tests - no error expected
		case !tc.expectError && err != nil:
			fmt.Println("Unexpected error:", err.Error())
			t.Fail()
		case !tc.expectError && err == nil:
			result := reflect.Indirect(reflect.ValueOf(pResult)).Interface()
			if !reflect.DeepEqual(tc.expected, result) {
				fmt.Println("FAIL")
				t.Errorf("Expected %#v, got %#v", tc.expected, result)
			} else {
				fmt.Println("OK")
			}

		// Negative tests - error expected
		case tc.expectError && err == nil:
			fmt.Println("FAIL")
			t.Errorf("Expected error not received")
		case tc.expectError && err != nil:
			fmt.Println("OK")
		}

		teardown1()
	}
}

func setup1(tc xmltestcase) {

	fmt.Printf("\n\tRegistering.... %s/%s\n",xmlrpcEndpoint, tc.service)
	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/%s", xmlrpcEndpoint, tc.service),
		tc.responder)
}

// remove any existig mocks (e.g., between tests)
func teardown1() {
	info := httpmock.GetCallCountInfo()
	fmt.Printf("\nCALLS\n%v\n", info)
	httpmock.Reset()
}
