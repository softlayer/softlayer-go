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
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/renier/xmlrpc"
	"github.com/softlayer/softlayer-go/sl"
)

// Used to pool the clients created per service
// so as to re-use service clients created previously
var xmlRpcClients = map[string]*xmlrpc.Client{}

// Debugging RoundTripper
type debugRoundTripper struct{}
func (mrt debugRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	log.Println("->>>Request:")
	dumpedReq, _ := httputil.DumpRequestOut(request, true)
	log.Println(string(dumpedReq))

	response, err := http.DefaultTransport.RoundTrip(request)
	if err != nil {
		log.Println("Error:", err)
		return response, err
	}

	log.Println("\n\n<<<-Response:")
	dumpedResp, _ := httputil.DumpResponse(response, true)
	log.Println(string(dumpedResp))

	return response, err
}

// XML-RPC Transport
type XmlRpcTransport struct {
	Timeout time.Duration
}
const DefaultXmlRpcTimeout = time.Second*30

// err = r.Session.DoRequest("SoftLayer_Account", "activatePartner", params, &r.Options, &resp)
func (x *XmlRpcTransport) DoRequest(
	sess *Session,
	service string,
	method string,
	args []interface{},
	options *sl.Options,
	pResult interface{},
) error {

	client, ok := xmlRpcClients[service]
	if !ok {
		var roundTripper http.RoundTripper
		var err error

		if sess.Debug {
			roundTripper = debugRoundTripper{}
		}

		timeout := DefaultXmlRpcTimeout
		if x.Timeout != 0 {
			timeout = x.Timeout
		}

		client, err = xmlrpc.NewClient(
			fmt.Sprintf("%s/%s", sess.Endpoint, service),
			roundTripper,
			timeout,
		)
		if err != nil {
			return fmt.Errorf("Could not create an xmlrpc client for %s: %s", service, err)
		}

		xmlRpcClients[service] = client
	}

	// TODO: Pass args into parameters.
	// TODO: Support token auth: complexType(PortalLoginToken), userId, and authToken under authenticate.
	return client.Call(method, map[string]interface{}{
		"headers": map[string]interface{}{
			"authenticate": map[string]string{
				"username": sess.UserName,
				"apiKey":   sess.APIKey,
			},
		},
	}, pResult)
}
