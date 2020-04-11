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
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

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

var xmltestcases = []xmltestcase{
	// Positive tests
	{
		description: "empty array return",
		service:     "SoftLayer_Account",
		method:      "getVirtualGuests",
		args:        nil,
		options:     sl.Options{},
		responder:   httpmock.NewStringResponder(200, `<?xml version="1.0" encoding="utf-8"?><params><param><value><array><data></data></array></value></param></params>`),
		expected:    make([]struct{}, 0),
		expectError: false,
	},
	{
		description: "Values in array",
		service:     "SoftLayer_Account",
		method:      "getVirtualGuests",
		args:        nil,
		options:     sl.Options{},
		responder:   httpmock.NewStringResponder(200, `<?xml version="1.0" encoding="utf-8"?><params><param><value><array><data><value><int>1</int></value><value><int>5</int></value></data></array></value></param></params>`),
		expected:    []int{1, 5},
		expectError: false,
	},
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
	{
		description: "Test Object Mask mask[]",
		service:     "SoftLayer_Account",
		method:      "getVirtualGuests",
		args:        nil,
		options:     sl.Options{Mask: "mask[id,hostname]"},
		responder:   responderCheckObjectMask,
		expected:    "OK",
		expectError: false,
	},
	{
		description: "Test Object Mask no mask[]",
		service:     "SoftLayer_Account",
		method:      "getVirtualGuests",
		args:        nil,
		options:     sl.Options{Mask: "id,hostname"},
		responder:   responderCheckObjectMask,
		expected:    "OK",
		expectError: false,
	},
}

func TestXmlRpc(t *testing.T) {
	// The structure for this test is inspired heavily from here:
	// https://github.com/softlayer/xmlrpc/blob/master/decoder_test.go

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
		setupxml(tc)

		// Do some reflecting to make comparisions of type and value easier
		pResult := reflect.New(reflect.TypeOf(tc.expected))
		expected := reflect.ValueOf(tc.expected)

		fmt.Printf("Test [%s]: ", tc.description)

		// Actually make the request
		err := s.DoRequest(tc.service, tc.method, tc.args, &tc.options, pResult.Interface())
		// Removes the pointer bits I think.
		pResult = pResult.Elem()

		// Report results
		switch {
		// Positive tests - no error expected
		case !tc.expectError && err != nil:
			fmt.Println("Unexpected error:", err.Error())
			t.Fail()
		case !tc.expectError && err == nil:
			// Slices comparisons gave me a lot of trouble with DeepEqual, check each element instead
			if expected.Kind() == reflect.Slice {
				if pResult.Len() != expected.Len() {
					fmt.Println("FAIL")
					t.Errorf("Slice length mismatch. Expcected %v, got %v", expected.Len(), pResult.Len())
				}
				for i := 0; i < pResult.Len(); i++ {
					if expected.Index(i).Interface() != pResult.Index(i).Interface() {
						fmt.Println("FAIL")
						t.Errorf("Expected %#v, got %#v", expected.Index(i), pResult.Index(i))
					}
				}
				// If we haven't failed until now, assume its ok.
				fmt.Println("OK")
			} else if !reflect.DeepEqual(expected.Interface(), pResult.Interface()) {
				fmt.Println("FAIL")
				t.Errorf("Expected %#v, got %#v", expected, pResult)
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
		teardownxml()
	}
}

func responderCheckObjectMask(req *http.Request) (*http.Response, error) {
	body, _ := ioutil.ReadAll(req.Body)
	string_body := string(body)
	// fmt.Printf("Body: %v\n", string_body)
	message := "FAILED"
	if strings.Contains(string_body, "<name>mask</name><value><string>mask[id,hostname]</string>") {
		message = "OK"
	}

	resp := httpmock.NewStringResponse(200, `<?xml version="1.0" encoding="utf-8"?><params><param><value>`+message+`</value></param></params>`)
	return resp, nil
}

func setupxml(tc xmltestcase) {
	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/%s", xmlrpcEndpoint, tc.service),
		tc.responder)
}

// remove any existig mocks (e.g., between tests)
func teardownxml() {
	httpmock.Reset()
}
