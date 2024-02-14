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
type Security_Certificate struct {
	Session session.SLSession
	Options sl.Options
}

// GetSecurityCertificateService returns an instance of the Security_Certificate SoftLayer service
func GetSecurityCertificateService(sess session.SLSession) Security_Certificate {
	return Security_Certificate{Session: sess}
}

func (r Security_Certificate) Id(id int) Security_Certificate {
	r.Options.Id = &id
	return r
}

func (r Security_Certificate) Mask(mask string) Security_Certificate {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Security_Certificate) Filter(filter string) Security_Certificate {
	r.Options.Filter = filter
	return r
}

func (r Security_Certificate) Limit(limit int) Security_Certificate {
	r.Options.Limit = &limit
	return r
}

func (r Security_Certificate) Offset(offset int) Security_Certificate {
	r.Options.Offset = &offset
	return r
}

// Add a certificate to your account for your records, or for use with various services. Only the certificate and private key are usually required. If your issuer provided an intermediate certificate, you must also provide that certificate. Details will be extracted from the certificate. Validation will be performed between the certificate and the private key as well as the certificate and the intermediate certificate, if provided.
//
// The certificate signing request is not required, but can be provided for your records.
func (r Security_Certificate) CreateObject(templateObject *datatypes.Security_Certificate) (resp datatypes.Security_Certificate, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Security_Certificate", "createObject", params, &r.Options, &resp)
	return
}

// Remove a certificate from your account. You may not remove a certificate with associated services.
func (r Security_Certificate) DeleteObject() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Security_Certificate", "deleteObject", nil, &r.Options, &resp)
	return
}

// Update a certificate. Modifications are restricted to the note and CSR if the are any services associated with the certificate. There are no modification restrictions for a certificate with no associated services.
func (r Security_Certificate) EditObject(templateObject *datatypes.Security_Certificate) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Security_Certificate", "editObject", params, &r.Options, &resp)
	return
}

// no documentation yet
func (r Security_Certificate) GetObject() (resp datatypes.Security_Certificate, err error) {
	err = r.Session.DoRequest("SoftLayer_Security_Certificate", "getObject", nil, &r.Options, &resp)
	return
}

// no documentation yet
type Security_Ssh_Key struct {
	Session session.SLSession
	Options sl.Options
}

// GetSecuritySshKeyService returns an instance of the Security_Ssh_Key SoftLayer service
func GetSecuritySshKeyService(sess session.SLSession) Security_Ssh_Key {
	return Security_Ssh_Key{Session: sess}
}

func (r Security_Ssh_Key) Id(id int) Security_Ssh_Key {
	r.Options.Id = &id
	return r
}

func (r Security_Ssh_Key) Mask(mask string) Security_Ssh_Key {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Security_Ssh_Key) Filter(filter string) Security_Ssh_Key {
	r.Options.Filter = filter
	return r
}

func (r Security_Ssh_Key) Limit(limit int) Security_Ssh_Key {
	r.Options.Limit = &limit
	return r
}

func (r Security_Ssh_Key) Offset(offset int) Security_Ssh_Key {
	r.Options.Offset = &offset
	return r
}

// Add a ssh key to your account for use during server provisioning and os reloads.
func (r Security_Ssh_Key) CreateObject(templateObject *datatypes.Security_Ssh_Key) (resp datatypes.Security_Ssh_Key, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Security_Ssh_Key", "createObject", params, &r.Options, &resp)
	return
}

// Remove a ssh key from your account.
func (r Security_Ssh_Key) DeleteObject() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Security_Ssh_Key", "deleteObject", nil, &r.Options, &resp)
	return
}

// Update a ssh key.
func (r Security_Ssh_Key) EditObject(templateObject *datatypes.Security_Ssh_Key) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Security_Ssh_Key", "editObject", params, &r.Options, &resp)
	return
}

// no documentation yet
func (r Security_Ssh_Key) GetObject() (resp datatypes.Security_Ssh_Key, err error) {
	err = r.Session.DoRequest("SoftLayer_Security_Ssh_Key", "getObject", nil, &r.Options, &resp)
	return
}
