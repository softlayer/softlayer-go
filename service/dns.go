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

type Dns_Domain struct {
	Session *Session
	Options
}

func (r *Session) GetDnsDomainService() Dns_Domain {
	return Dns_Domain{Session: r}
}

func (r *Dns_Domain) CreateARecord(host *string, data *string, ttl *int) (resp datatypes.Dns_Domain_ResourceRecord_AType, err error) {
	params := []interface{}{
		host,
		data,
		ttl,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain) CreateAaaaRecord(host *string, data *string, ttl *int) (resp datatypes.Dns_Domain_ResourceRecord_AaaaType, err error) {
	params := []interface{}{
		host,
		data,
		ttl,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain) CreateCnameRecord(host *string, data *string, ttl *int) (resp datatypes.Dns_Domain_ResourceRecord_CnameType, err error) {
	params := []interface{}{
		host,
		data,
		ttl,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain) CreateMxRecord(host *string, data *string, ttl *int, mxPriority *int) (resp datatypes.Dns_Domain_ResourceRecord_MxType, err error) {
	params := []interface{}{
		host,
		data,
		ttl,
		mxPriority,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain) CreateNsRecord(host *string, data *string, ttl *int) (resp datatypes.Dns_Domain_ResourceRecord_NsType, err error) {
	params := []interface{}{
		host,
		data,
		ttl,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain) CreateObject(templateObject *datatypes.Dns_Domain) (resp datatypes.Dns_Domain, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain) CreateObjects(templateObjects []datatypes.Dns_Domain) (resp []datatypes.Dns_Domain, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain) CreatePtrRecord(ipAddress *string, ptrRecord *string, ttl *int) (resp datatypes.Dns_Domain_ResourceRecord, err error) {
	params := []interface{}{
		ipAddress,
		ptrRecord,
		ttl,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain) CreateSpfRecord(host *string, data *string, ttl *int) (resp datatypes.Dns_Domain_ResourceRecord_SpfType, err error) {
	params := []interface{}{
		host,
		data,
		ttl,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain) CreateTxtRecord(host *string, data *string, ttl *int) (resp datatypes.Dns_Domain_ResourceRecord_TxtType, err error) {
	params := []interface{}{
		host,
		data,
		ttl,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain) GetByDomainName(name *string) (resp []datatypes.Dns_Domain, err error) {
	params := []interface{}{
		name,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain) GetManagedResourceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain) GetObject() (resp datatypes.Dns_Domain, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain) GetResourceRecords() (resp []datatypes.Dns_Domain_ResourceRecord, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain) GetSecondary() (resp datatypes.Dns_Secondary, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain) GetZoneFileContents() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Dns_Domain_Registration struct {
	Session *Session
	Options
}

func (r *Session) GetDnsDomainRegistrationService() Dns_Domain_Registration {
	return Dns_Domain_Registration{Session: r}
}

func (r *Dns_Domain_Registration) AddNameserversToDomain(nameservers []string) (resp bool, err error) {
	params := []interface{}{
		nameservers,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_Registration) DeleteRegisteredNameserver(nameserver *string) (resp bool, err error) {
	params := []interface{}{
		nameserver,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_Registration) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_Registration) GetAuthenticationCode() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_Registration) GetDomainInformation() (resp datatypes.Container_Dns_Domain_Registration_Information, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_Registration) GetDomainNameservers() (resp []datatypes.Container_Dns_Domain_Registration_Nameserver, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_Registration) GetDomainRegistrationStatus() (resp datatypes.Dns_Domain_Registration_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_Registration) GetExtendedAttributes(domainName *string) (resp []datatypes.Container_Dns_Domain_Registration_ExtendedAttribute, err error) {
	params := []interface{}{
		domainName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_Registration) GetObject() (resp datatypes.Dns_Domain_Registration, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_Registration) GetRegisteredNameserver() (resp datatypes.Container_Dns_Domain_Registration_Nameserver, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_Registration) GetRegistrantVerificationStatus() (resp datatypes.Dns_Domain_Registration_Registrant_Verification_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_Registration) GetRegistrantVerificationStatusDetail() (resp datatypes.Container_Dns_Domain_Registration_Registrant_Verification_StatusDetail, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_Registration) GetServiceProvider() (resp datatypes.Service_Provider, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_Registration) GetTransferInformation(domainName *string) (resp datatypes.Container_Dns_Domain_Registration_Transfer_Information, err error) {
	params := []interface{}{
		domainName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_Registration) LockDomain() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_Registration) LookupDomain(domainName *string) (resp []datatypes.Container_Dns_Domain_Registration_Lookup, err error) {
	params := []interface{}{
		domainName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_Registration) ModifyContact(contact *datatypes.Container_Dns_Domain_Registration_Contact) (resp bool, err error) {
	params := []interface{}{
		contact,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_Registration) ModifyRegisteredNameserver(oldNameserver *string, newNameserver *string, ipAddress *string) (resp bool, err error) {
	params := []interface{}{
		oldNameserver,
		newNameserver,
		ipAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_Registration) RegisterNameserver(nameserver *string, ipAddress *string) (resp bool, err error) {
	params := []interface{}{
		nameserver,
		ipAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_Registration) RemoveNameserversFromDomain(nameservers []string) (resp bool, err error) {
	params := []interface{}{
		nameservers,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_Registration) SendAuthenticationCode() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_Registration) SendRegistrantVerificationEmail() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_Registration) SendTransferApprovalEmail() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_Registration) SetAuthenticationCode(authenticationCode *string) (resp bool, err error) {
	params := []interface{}{
		authenticationCode,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_Registration) UnlockDomain() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Dns_Domain_Registration_Registrant_Verification_Status struct {
	Session *Session
	Options
}

func (r *Session) GetDnsDomainRegistrationRegistrantVerificationStatusService() Dns_Domain_Registration_Registrant_Verification_Status {
	return Dns_Domain_Registration_Registrant_Verification_Status{Session: r}
}

func (r *Dns_Domain_Registration_Registrant_Verification_Status) GetAllObjects() (resp []datatypes.Dns_Domain_Registration_Registrant_Verification_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_Registration_Registrant_Verification_Status) GetObject() (resp datatypes.Dns_Domain_Registration_Registrant_Verification_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Dns_Domain_Registration_Status struct {
	Session *Session
	Options
}

func (r *Session) GetDnsDomainRegistrationStatusService() Dns_Domain_Registration_Status {
	return Dns_Domain_Registration_Status{Session: r}
}

func (r *Dns_Domain_Registration_Status) GetAllObjects() (resp []datatypes.Dns_Domain_Registration_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_Registration_Status) GetObject() (resp datatypes.Dns_Domain_Registration_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Dns_Domain_ResourceRecord struct {
	Session *Session
	Options
}

func (r *Session) GetDnsDomainResourceRecordService() Dns_Domain_ResourceRecord {
	return Dns_Domain_ResourceRecord{Session: r}
}

func (r *Dns_Domain_ResourceRecord) CreateObject(templateObject *datatypes.Dns_Domain_ResourceRecord) (resp datatypes.Dns_Domain_ResourceRecord, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_ResourceRecord) CreateObjects(templateObjects []datatypes.Dns_Domain_ResourceRecord) (resp []datatypes.Dns_Domain_ResourceRecord, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_ResourceRecord) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_ResourceRecord) DeleteObjects(templateObjects []datatypes.Dns_Domain_ResourceRecord) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_ResourceRecord) EditObject(templateObject *datatypes.Dns_Domain_ResourceRecord) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_ResourceRecord) EditObjects(templateObjects []datatypes.Dns_Domain_ResourceRecord) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_ResourceRecord) GetDomain() (resp datatypes.Dns_Domain, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_ResourceRecord) GetObject() (resp datatypes.Dns_Domain_ResourceRecord, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Dns_Domain_ResourceRecord_MxType struct {
	Session *Session
	Options
}

func (r *Session) GetDnsDomainResourceRecordMxTypeService() Dns_Domain_ResourceRecord_MxType {
	return Dns_Domain_ResourceRecord_MxType{Session: r}
}

func (r *Dns_Domain_ResourceRecord_MxType) CreateObject(templateObject *datatypes.Dns_Domain_ResourceRecord_MxType) (resp datatypes.Dns_Domain_ResourceRecord_MxType, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_ResourceRecord_MxType) CreateObjects(templateObjects []datatypes.Dns_Domain_ResourceRecord) (resp []datatypes.Dns_Domain_ResourceRecord, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_ResourceRecord_MxType) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_ResourceRecord_MxType) DeleteObjects(templateObjects []datatypes.Dns_Domain_ResourceRecord_MxType) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_ResourceRecord_MxType) EditObject(templateObject *datatypes.Dns_Domain_ResourceRecord_MxType) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_ResourceRecord_MxType) EditObjects(templateObjects []datatypes.Dns_Domain_ResourceRecord_MxType) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_ResourceRecord_MxType) GetDomain() (resp datatypes.Dns_Domain, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_ResourceRecord_MxType) GetObject() (resp datatypes.Dns_Domain_ResourceRecord_MxType, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Dns_Domain_ResourceRecord_SrvType struct {
	Session *Session
	Options
}

func (r *Session) GetDnsDomainResourceRecordSrvTypeService() Dns_Domain_ResourceRecord_SrvType {
	return Dns_Domain_ResourceRecord_SrvType{Session: r}
}

func (r *Dns_Domain_ResourceRecord_SrvType) CreateObject(templateObject *datatypes.Dns_Domain_ResourceRecord_SrvType) (resp datatypes.Dns_Domain_ResourceRecord_SrvType, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_ResourceRecord_SrvType) CreateObjects(templateObjects []datatypes.Dns_Domain_ResourceRecord) (resp []datatypes.Dns_Domain_ResourceRecord, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_ResourceRecord_SrvType) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_ResourceRecord_SrvType) DeleteObjects(templateObjects []datatypes.Dns_Domain_ResourceRecord_SrvType) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_ResourceRecord_SrvType) EditObject(templateObject *datatypes.Dns_Domain_ResourceRecord_SrvType) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_ResourceRecord_SrvType) EditObjects(templateObjects []datatypes.Dns_Domain_ResourceRecord_SrvType) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_ResourceRecord_SrvType) GetDomain() (resp datatypes.Dns_Domain, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Domain_ResourceRecord_SrvType) GetObject() (resp datatypes.Dns_Domain_ResourceRecord_SrvType, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Dns_Secondary struct {
	Session *Session
	Options
}

func (r *Session) GetDnsSecondaryService() Dns_Secondary {
	return Dns_Secondary{Session: r}
}

func (r *Dns_Secondary) ConvertToPrimary() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Secondary) CreateObject(templateObject *datatypes.Dns_Secondary) (resp datatypes.Dns_Secondary, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Secondary) CreateObjects(templateObjects []datatypes.Dns_Secondary) (resp []datatypes.Dns_Secondary, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Secondary) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Secondary) EditObject(templateObject *datatypes.Dns_Secondary) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Secondary) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Secondary) GetByDomainName(name *string) (resp []datatypes.Dns_Secondary, err error) {
	params := []interface{}{
		name,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Secondary) GetDomain() (resp datatypes.Dns_Domain, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Secondary) GetErrorMessages() (resp []datatypes.Dns_Message, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Secondary) GetObject() (resp datatypes.Dns_Secondary, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Secondary) GetStatus() (resp datatypes.Dns_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Dns_Secondary) TransferNow() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
