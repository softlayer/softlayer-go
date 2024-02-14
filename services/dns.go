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

// The SoftLayer_Dns_Domain data type represents a single DNS domain record hosted on the SoftLayer nameservers. Domains contain general information about the domain name such as name and serial. Individual records such as A, AAAA, CTYPE, and MX records are stored in the domain's associated [[SoftLayer_Dns_Domain_ResourceRecord (type)|SoftLayer_Dns_Domain_ResourceRecord]] records.
type Dns_Domain struct {
	Session session.SLSession
	Options sl.Options
}

// GetDnsDomainService returns an instance of the Dns_Domain SoftLayer service
func GetDnsDomainService(sess session.SLSession) Dns_Domain {
	return Dns_Domain{Session: sess}
}

func (r Dns_Domain) Id(id int) Dns_Domain {
	r.Options.Id = &id
	return r
}

func (r Dns_Domain) Mask(mask string) Dns_Domain {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Dns_Domain) Filter(filter string) Dns_Domain {
	r.Options.Filter = filter
	return r
}

func (r Dns_Domain) Limit(limit int) Dns_Domain {
	r.Options.Limit = &limit
	return r
}

func (r Dns_Domain) Offset(offset int) Dns_Domain {
	r.Options.Offset = &offset
	return r
}

// Create a new domain on the SoftLayer name servers. The SoftLayer_Dns_Domain object passed to this function must have at least one A or AAAA resource record.
//
// createObject creates a default SOA record with the data:
// * ”'host”': "@"
// * ”'data”': "ns1.softlayer.com."
// * ”'responsible person”': "root.[your domain name]."
// * ”'expire”': 604800 seconds
// * ”'refresh”': 3600 seconds
// * ”'retry”': 300 seconds
// * ”'minimum”': 3600 seconds
//
// If your new domain uses the .de top-level domain then SOA refresh is set to 10000 seconds, retry is set to 1800 seconds, and minimum to 10000 seconds.
//
// If your domain doesn't contain NS resource records for ns1.softlayer.com or ns2.softlayer.com then ”createObject” will create them for you.
//
// ”createObject” returns a Boolean ”true” on successful object creation or ”false” if your domain was unable to be created..
func (r Dns_Domain) CreateObject(templateObject *datatypes.Dns_Domain) (resp datatypes.Dns_Domain, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Dns_Domain", "createObject", params, &r.Options, &resp)
	return
}

// deleteObject permanently removes a domain and all of it's associated resource records from the softlayer name servers. ”'This cannot be undone.”' Be wary of running this method. If you remove a domain in error you will need to re-create it by creating a new SoftLayer_Dns_Domain object.
func (r Dns_Domain) DeleteObject() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Dns_Domain", "deleteObject", nil, &r.Options, &resp)
	return
}

// Search for [[SoftLayer_Dns_Domain]] records by domain name. getByDomainName() performs an inclusive search for domain records, returning multiple records based on partial name matches. Use this method to locate domain records if you don't have access to their id numbers.
func (r Dns_Domain) GetByDomainName(name *string) (resp []datatypes.Dns_Domain, err error) {
	params := []interface{}{
		name,
	}
	err = r.Session.DoRequest("SoftLayer_Dns_Domain", "getByDomainName", params, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Dns_Domain object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Dns_Domain service. You can only retrieve domains that are assigned to your SoftLayer account.
func (r Dns_Domain) GetObject() (resp datatypes.Dns_Domain, err error) {
	err = r.Session.DoRequest("SoftLayer_Dns_Domain", "getObject", nil, &r.Options, &resp)
	return
}

// Retrieve The individual records contained within a domain record. These include but are not limited to A, AAAA, MX, CTYPE, SPF and TXT records.
func (r Dns_Domain) GetResourceRecords() (resp []datatypes.Dns_Domain_ResourceRecord, err error) {
	err = r.Session.DoRequest("SoftLayer_Dns_Domain", "getResourceRecords", nil, &r.Options, &resp)
	return
}

// Return a SoftLayer hosted domain and resource records' data formatted as zone file.
func (r Dns_Domain) GetZoneFileContents() (resp string, err error) {
	err = r.Session.DoRequest("SoftLayer_Dns_Domain", "getZoneFileContents", nil, &r.Options, &resp)
	return
}

// The SoftLayer_Dns_Domain_ResourceRecord data type represents a single resource record entry in a SoftLayer hosted domain. Each resource record contains a ”host” and ”data” property, defining a resource's name and it's target data. Domains contain multiple types of resource records. The ”type” property separates out resource records by type. ”Type” can take one of the following values:
// * ”'"a"”' for [[SoftLayer_Dns_Domain_ResourceRecord_AType|address]] records
// * ”'"aaaa"”' for [[SoftLayer_Dns_Domain_ResourceRecord_AaaaType|address]] records
// * ”'"cname"”' for [[SoftLayer_Dns_Domain_ResourceRecord_CnameType|canonical name]] records
// * ”'"mx"”' for [[SoftLayer_Dns_Domain_ResourceRecord_MxType|mail exchanger]] records
// * ”'"ns"”' for [[SoftLayer_Dns_Domain_ResourceRecord_NsType|name server]] records
// * ”'"ptr"”' for [[SoftLayer_Dns_Domain_ResourceRecord_PtrType|pointer]] records in reverse domains
// * ”'"soa"”' for a domain's [[SoftLayer_Dns_Domain_ResourceRecord_SoaType|start of authority]] record
// * ”'"spf"”' for [[SoftLayer_Dns_Domain_ResourceRecord_SpfType|sender policy framework]] records
// * ”'"srv"”' for [[SoftLayer_Dns_Domain_ResourceRecord_SrvType|service]] records
// * ”'"txt"”' for [[SoftLayer_Dns_Domain_ResourceRecord_TxtType|text]] records
//
// As ”SoftLayer_Dns_Domain_ResourceRecord” objects are created and loaded, the API verifies the ”type” property and casts the object as the appropriate type.
type Dns_Domain_ResourceRecord struct {
	Session session.SLSession
	Options sl.Options
}

// GetDnsDomainResourceRecordService returns an instance of the Dns_Domain_ResourceRecord SoftLayer service
func GetDnsDomainResourceRecordService(sess session.SLSession) Dns_Domain_ResourceRecord {
	return Dns_Domain_ResourceRecord{Session: sess}
}

func (r Dns_Domain_ResourceRecord) Id(id int) Dns_Domain_ResourceRecord {
	r.Options.Id = &id
	return r
}

func (r Dns_Domain_ResourceRecord) Mask(mask string) Dns_Domain_ResourceRecord {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Dns_Domain_ResourceRecord) Filter(filter string) Dns_Domain_ResourceRecord {
	r.Options.Filter = filter
	return r
}

func (r Dns_Domain_ResourceRecord) Limit(limit int) Dns_Domain_ResourceRecord {
	r.Options.Limit = &limit
	return r
}

func (r Dns_Domain_ResourceRecord) Offset(offset int) Dns_Domain_ResourceRecord {
	r.Options.Offset = &offset
	return r
}

// createObject creates a new domain resource record. The ”host” property of the templateObject parameter is scrubbed to remove all non-alpha numeric characters except for "@", "_", ".", "*", and "-". The ”data” property of the templateObject parameter is scrubbed to remove all non-alphanumeric characters for "." and "-". Creating a resource record updates the serial number of the domain the resource record is associated with.
//
// ”createObject” returns Boolean ”true” on successful create or ”false” if it was unable to create a resource record.
func (r Dns_Domain_ResourceRecord) CreateObject(templateObject *datatypes.Dns_Domain_ResourceRecord) (resp datatypes.Dns_Domain_ResourceRecord, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Dns_Domain_ResourceRecord", "createObject", params, &r.Options, &resp)
	return
}

// Delete a domain's resource record. ”'This cannot be undone.”' Be wary of running this method. If you remove a resource record in error you will need to re-create it by creating a new SoftLayer_Dns_Domain_ResourceRecord object. The serial number of the domain associated with this resource record is updated upon deletion. You may not delete SOA, NS, or PTR resource records.
//
// ”deleteObject” returns Boolean ”true” on successful deletion or ”false” if it was unable to remove a resource record.
func (r Dns_Domain_ResourceRecord) DeleteObject() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Dns_Domain_ResourceRecord", "deleteObject", nil, &r.Options, &resp)
	return
}

// editObject edits an existing domain resource record. The ”host” property of the templateObject parameter is scrubbed to remove all non-alpha numeric characters except for "@", "_", ".", "*", and "-". The ”data” property of the templateObject parameter is scrubbed to remove all non-alphanumeric characters for "." and "-". Editing a resource record updates the serial number of the domain the resource record is associated with.
//
// ”editObject” returns Boolean ”true” on a successful edit or ”false” if it was unable to edit the resource record.
func (r Dns_Domain_ResourceRecord) EditObject(templateObject *datatypes.Dns_Domain_ResourceRecord) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Dns_Domain_ResourceRecord", "editObject", params, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Dns_Domain_ResourceRecord object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Dns_Domain_ResourceRecord service. You can only retrieve resource records belonging to domains that are assigned to your SoftLayer account.
func (r Dns_Domain_ResourceRecord) GetObject() (resp datatypes.Dns_Domain_ResourceRecord, err error) {
	err = r.Session.DoRequest("SoftLayer_Dns_Domain_ResourceRecord", "getObject", nil, &r.Options, &resp)
	return
}

// SoftLayer_Dns_Domain_ResourceRecord_SrvType is a SoftLayer_Dns_Domain_ResourceRecord object whose ”type” property is set to "srv" and defines a DNS SRV record on a SoftLayer hosted domain.
type Dns_Domain_ResourceRecord_SrvType struct {
	Session session.SLSession
	Options sl.Options
}

// GetDnsDomainResourceRecordSrvTypeService returns an instance of the Dns_Domain_ResourceRecord_SrvType SoftLayer service
func GetDnsDomainResourceRecordSrvTypeService(sess session.SLSession) Dns_Domain_ResourceRecord_SrvType {
	return Dns_Domain_ResourceRecord_SrvType{Session: sess}
}

func (r Dns_Domain_ResourceRecord_SrvType) Id(id int) Dns_Domain_ResourceRecord_SrvType {
	r.Options.Id = &id
	return r
}

func (r Dns_Domain_ResourceRecord_SrvType) Mask(mask string) Dns_Domain_ResourceRecord_SrvType {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Dns_Domain_ResourceRecord_SrvType) Filter(filter string) Dns_Domain_ResourceRecord_SrvType {
	r.Options.Filter = filter
	return r
}

func (r Dns_Domain_ResourceRecord_SrvType) Limit(limit int) Dns_Domain_ResourceRecord_SrvType {
	r.Options.Limit = &limit
	return r
}

func (r Dns_Domain_ResourceRecord_SrvType) Offset(offset int) Dns_Domain_ResourceRecord_SrvType {
	r.Options.Offset = &offset
	return r
}

// createObject creates a new SRV record. The ”host” property of the templateObject parameter is scrubbed to remove all non-alpha numeric characters except for "@", "_", ".", "*", and "-". The ”data” property of the templateObject parameter is scrubbed to remove all non-alphanumeric characters for "." and "-". Creating an SRV record updates the serial number of the domain the resource record is associated with.
func (r Dns_Domain_ResourceRecord_SrvType) CreateObject(templateObject *datatypes.Dns_Domain_ResourceRecord_SrvType) (resp datatypes.Dns_Domain_ResourceRecord_SrvType, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Dns_Domain_ResourceRecord_SrvType", "createObject", params, &r.Options, &resp)
	return
}
