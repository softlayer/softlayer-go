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

/**
 * AUTOMATICALLY GENERATED CODE - DO NOT MODIFY
 */

package services

import "github.ibm.com/riethm/gopherlayer/datatypes"

// The SoftLayer_Event_Log data type contains an event detail occurred upon various SoftLayer resources.
type Event_Log struct {
	Session *Session
	Options
}

func (r *Session) GetEventLogService() Event_Log {
	return Event_Log{Session: r}
}

// This all indexed event names.
func (r *Event_Log) GetAllEventNames(objectName *string) (resp []string, err error) {
	params := []interface{}{
		objectName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// This all indexed event object names.
func (r *Event_Log) GetAllEventObjectNames() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Event_Log) GetAllObjects() (resp []datatypes.Event_Log, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Event_Log) GetAllUserTypes() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Event_Log) GetUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
