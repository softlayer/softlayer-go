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
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

const DEFAULT_ENDPOINT = "https://api.softlayer.com/rest/v3"

type Session struct {
	UserName string
	ApiKey   string
	Endpoint string
	Debug    bool
}

func (r *Session) String() string {
	return "Username: " + r.UserName +
		", ApiKey: " + r.ApiKey +
		", Endpoint: " + r.Endpoint
}

func NewSession(u string, k string, args ...interface{}) Session {
	var e string

	if len(args) > 0 {
		e = args[0].(string)
	} else {
		e = DEFAULT_ENDPOINT
	}

	// TODO: Read credentials from ~/.softlayer. Requires dependency for parsing ini file
	envFallback("SOFTLAYER_USERNAME", &u)
	envFallback("SOFTLAYER_API_KEY", &k)

	return Session{UserName: u, ApiKey: k, Endpoint: e}
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

	url := session.Endpoint + "/" + path
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
