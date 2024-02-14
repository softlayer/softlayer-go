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

// The SoftLayer_Provisioning_Hook contains all the information needed to add a hook into a server/Virtual provision and os reload.
type Provisioning_Hook struct {
	Session session.SLSession
	Options sl.Options
}

// GetProvisioningHookService returns an instance of the Provisioning_Hook SoftLayer service
func GetProvisioningHookService(sess session.SLSession) Provisioning_Hook {
	return Provisioning_Hook{Session: sess}
}

func (r Provisioning_Hook) Id(id int) Provisioning_Hook {
	r.Options.Id = &id
	return r
}

func (r Provisioning_Hook) Mask(mask string) Provisioning_Hook {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Provisioning_Hook) Filter(filter string) Provisioning_Hook {
	r.Options.Filter = filter
	return r
}

func (r Provisioning_Hook) Limit(limit int) Provisioning_Hook {
	r.Options.Limit = &limit
	return r
}

func (r Provisioning_Hook) Offset(offset int) Provisioning_Hook {
	r.Options.Offset = &offset
	return r
}

// no documentation yet
func (r Provisioning_Hook) CreateObject(templateObject *datatypes.Provisioning_Hook) (resp datatypes.Provisioning_Hook, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Provisioning_Hook", "createObject", params, &r.Options, &resp)
	return
}

// no documentation yet
func (r Provisioning_Hook) DeleteObject() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Provisioning_Hook", "deleteObject", nil, &r.Options, &resp)
	return
}
