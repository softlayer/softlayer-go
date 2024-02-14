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

// SoftLayer_Software_AccountLicense is a class that represents software licenses that are tied only to a customer's account and not to any particular hardware, IP address, etc.
type Software_AccountLicense struct {
	Session session.SLSession
	Options sl.Options
}

// GetSoftwareAccountLicenseService returns an instance of the Software_AccountLicense SoftLayer service
func GetSoftwareAccountLicenseService(sess session.SLSession) Software_AccountLicense {
	return Software_AccountLicense{Session: sess}
}

func (r Software_AccountLicense) Id(id int) Software_AccountLicense {
	r.Options.Id = &id
	return r
}

func (r Software_AccountLicense) Mask(mask string) Software_AccountLicense {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Software_AccountLicense) Filter(filter string) Software_AccountLicense {
	r.Options.Filter = filter
	return r
}

func (r Software_AccountLicense) Limit(limit int) Software_AccountLicense {
	r.Options.Limit = &limit
	return r
}

func (r Software_AccountLicense) Offset(offset int) Software_AccountLicense {
	r.Options.Offset = &offset
	return r
}

// no documentation yet
func (r Software_AccountLicense) GetAllObjects() (resp []datatypes.Software_AccountLicense, err error) {
	err = r.Session.DoRequest("SoftLayer_Software_AccountLicense", "getAllObjects", nil, &r.Options, &resp)
	return
}

// no documentation yet
func (r Software_AccountLicense) GetObject() (resp datatypes.Software_AccountLicense, err error) {
	err = r.Session.DoRequest("SoftLayer_Software_AccountLicense", "getObject", nil, &r.Options, &resp)
	return
}

// This SoftLayer_Software_Component_Password data type contains a password for a specific software component instance.
type Software_Component_Password struct {
	Session session.SLSession
	Options sl.Options
}

// GetSoftwareComponentPasswordService returns an instance of the Software_Component_Password SoftLayer service
func GetSoftwareComponentPasswordService(sess session.SLSession) Software_Component_Password {
	return Software_Component_Password{Session: sess}
}

func (r Software_Component_Password) Id(id int) Software_Component_Password {
	r.Options.Id = &id
	return r
}

func (r Software_Component_Password) Mask(mask string) Software_Component_Password {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Software_Component_Password) Filter(filter string) Software_Component_Password {
	r.Options.Filter = filter
	return r
}

func (r Software_Component_Password) Limit(limit int) Software_Component_Password {
	r.Options.Limit = &limit
	return r
}

func (r Software_Component_Password) Offset(offset int) Software_Component_Password {
	r.Options.Offset = &offset
	return r
}

// Create a password for a software component.
func (r Software_Component_Password) CreateObject(templateObject *datatypes.Software_Component_Password) (resp datatypes.Software_Component_Password, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Software_Component_Password", "createObject", params, &r.Options, &resp)
	return
}

// no documentation yet
func (r Software_Component_Password) GetObject() (resp datatypes.Software_Component_Password, err error) {
	err = r.Session.DoRequest("SoftLayer_Software_Component_Password", "getObject", nil, &r.Options, &resp)
	return
}
