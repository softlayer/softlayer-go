/**
 * Copyright 2016-2024 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
 * the License. You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed
 * on an "AS IS" BASIS,WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and limitations under the License.
 */

// AUTOMATICALLY GENERATED CODE - DO NOT MODIFY

package services

import (
	"fmt"
	"strings"

	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/sl"
)

// no documentation yet
type Notification_Occurrence_Event struct {
	Session session.SLSession
	Options sl.Options
}

// GetNotificationOccurrenceEventService returns an instance of the Notification_Occurrence_Event SoftLayer service
func GetNotificationOccurrenceEventService(sess session.SLSession) Notification_Occurrence_Event {
	return Notification_Occurrence_Event{Session: sess}
}

func (r Notification_Occurrence_Event) Id(id int) Notification_Occurrence_Event {
	r.Options.Id = &id
	return r
}

func (r Notification_Occurrence_Event) Mask(mask string) Notification_Occurrence_Event {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Notification_Occurrence_Event) Filter(filter string) Notification_Occurrence_Event {
	r.Options.Filter = filter
	return r
}

func (r Notification_Occurrence_Event) Limit(limit int) Notification_Occurrence_Event {
	r.Options.Limit = &limit
	return r
}

func (r Notification_Occurrence_Event) Offset(offset int) Notification_Occurrence_Event {
	r.Options.Offset = &offset
	return r
}

// <<<< EOT
func (r Notification_Occurrence_Event) AcknowledgeNotification() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Notification_Occurrence_Event", "acknowledgeNotification", nil, &r.Options, &resp)
	return
}

// no documentation yet
func (r Notification_Occurrence_Event) GetAllObjects() (resp []datatypes.Notification_Occurrence_Event, err error) {
	err = r.Session.DoRequest("SoftLayer_Notification_Occurrence_Event", "getAllObjects", nil, &r.Options, &resp)
	return
}

// no documentation yet
func (r Notification_Occurrence_Event) GetObject() (resp datatypes.Notification_Occurrence_Event, err error) {
	err = r.Session.DoRequest("SoftLayer_Notification_Occurrence_Event", "getObject", nil, &r.Options, &resp)
	return
}
