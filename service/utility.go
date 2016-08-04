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

type Utility_Bandwidth_Graph struct {
	Session *Session
	Options
}

func (r *Session) GetUtilityBandwidthGraphService() Utility_Bandwidth_Graph {
	return Utility_Bandwidth_Graph{Session: r}
}

type Utility_Network struct {
	Session *Session
	Options
}

func (r *Session) GetUtilityNetworkService() Utility_Network {
	return Utility_Network{Session: r}
}

func (r *Utility_Network) NsLookup(address *string, typ *string) (resp string, err error) {
	params := []interface{}{
		address,
		typ,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Utility_Network) Whois(address *string) (resp string, err error) {
	params := []interface{}{
		address,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Utility_ObjectFilter struct {
	Session *Session
	Options
}

func (r *Session) GetUtilityObjectFilterService() Utility_ObjectFilter {
	return Utility_ObjectFilter{Session: r}
}

type Utility_ObjectFilter_Operation struct {
	Session *Session
	Options
}

func (r *Session) GetUtilityObjectFilterOperationService() Utility_ObjectFilter_Operation {
	return Utility_ObjectFilter_Operation{Session: r}
}

type Utility_ObjectFilter_Operation_Option struct {
	Session *Session
	Options
}

func (r *Session) GetUtilityObjectFilterOperationOptionService() Utility_ObjectFilter_Operation_Option {
	return Utility_ObjectFilter_Operation_Option{Session: r}
}
