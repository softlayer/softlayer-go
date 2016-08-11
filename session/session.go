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
	"os"
	"os/user"
	"strings"

	"github.ibm.com/riethm/gopherlayer/config"
	"github.ibm.com/riethm/gopherlayer/sl"
)

const DEFAULT_ENDPOINT = "https://api.softlayer.com/rest/v3"

type transportHandlerFunc func(
	sess *Session,
	service string,
	method string,
	args []interface{},
	options *sl.Options,
	pResult interface{}) error

type Session struct {
	UserName         string
	ApiKey           string
	Endpoint         string
	Debug            bool
	transportHandler transportHandlerFunc
}

func New(args ...interface{}) *Session {
	keys := map[string]int{"username": 0, "api_key": 1, "endpoint_url": 2}
	values := []string{"", "", ""}

	for i := 0; i < len(args); i++ {
		values[i] = args[i].(string)
	}

	// Default to the environment variables
	envFallback("SOFTLAYER_USERNAME", &values[keys["username"]])
	envFallback("SOFTLAYER_API_KEY", &values[keys["api_key"]])
	envFallback("SOFTLAYER_ENDPOINT_URL", &values[keys["endpoint_url"]])

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

	endpointUrl := values[keys["endpoint_url"]]
	if endpointUrl == "" || !strings.Contains(endpointUrl, "/rest/") {
		endpointUrl = DEFAULT_ENDPOINT
	}

	return &Session{
		UserName:         values[keys["username"]],
		ApiKey:           values[keys["api_key"]],
		Endpoint:         endpointUrl,
		transportHandler: doRestRequest,
	}
}

func (r *Session) DoRequest(service string, method string, args []interface{}, options *sl.Options, pResult interface{}) error {
	return r.transportHandler(r, service, method, args, options, pResult)
}

func envFallback(keyName string, value *string) {
	if *value == "" {
		*value = os.Getenv(keyName)
	}
}
