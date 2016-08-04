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

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Search struct {
	Session *Session
	Options
}

func (r *Session) GetSearchService() Search {
	return Search{Session: r}
}

func (r *Search) AdvancedSearch(searchString *string) (resp []datatypes.Container_Search_Result, err error) {
	params := []interface{}{
		searchString,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Search) GetObjectTypes() (resp []datatypes.Container_Search_ObjectType, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Search) Search(searchString *string) (resp []datatypes.Container_Search_Result, err error) {
	params := []interface{}{
		searchString,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
