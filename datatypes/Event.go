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

package datatypes

import "time"

type Event_Log struct {
	Entity

	AccountId       *int           `json:"accountId,omitempty"`
	EventCreateDate *time.Time     `json:"eventCreateDate,omitempty"`
	EventName       *string        `json:"eventName,omitempty"`
	IpAddress       *string        `json:"ipAddress,omitempty"`
	Label           *string        `json:"label,omitempty"`
	MetaData        *string        `json:"metaData,omitempty"`
	ObjectId        *int           `json:"objectId,omitempty"`
	ObjectName      *string        `json:"objectName,omitempty"`
	Resource        *Entity        `json:"resource,omitempty"`
	TraceId         *string        `json:"traceId,omitempty"`
	User            *User_Customer `json:"user,omitempty"`
	UserId          *int           `json:"userId,omitempty"`
	UserType        *string        `json:"userType,omitempty"`
	Username        *string        `json:"username,omitempty"`
}
