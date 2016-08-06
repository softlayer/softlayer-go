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

type Tag struct {
	Session *Session
	Options
}

func (r *Session) GetTagService() Tag {
	return Tag{Session: r}
}

func (r *Tag) AutoComplete(tag *string) (resp []datatypes.Tag, err error) {
	params := []interface{}{
		tag,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Tag) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Tag) GetAllTagTypes() (resp []datatypes.Tag_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Tag) GetObject() (resp datatypes.Tag, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Tag) GetReferences() (resp []datatypes.Tag_Reference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Tag) GetTagByTagName(tagList *string) (resp []datatypes.Tag, err error) {
	params := []interface{}{
		tagList,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Tag) SetTags(tags *string, keyName *string, resourceTableId *int) (resp bool, err error) {
	params := []interface{}{
		tags,
		keyName,
		resourceTableId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
