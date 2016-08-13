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

import (
	"github.ibm.com/riethm/gopherlayer/datatypes"
	"github.ibm.com/riethm/gopherlayer/session"
	"github.ibm.com/riethm/gopherlayer/sl"
)

//
type Catalyst_Company_Type struct {
	Session *session.Session
	Options sl.Options
}

func GetCatalystCompanyTypeService(sess *session.Session) Catalyst_Company_Type {
	return Catalyst_Company_Type{Session: sess}
}

func (r Catalyst_Company_Type) Id(id int) Catalyst_Company_Type {
	r.Options.Id = &id
	return r
}

func (r Catalyst_Company_Type) Mask(mask string) Catalyst_Company_Type {
	r.Options.Mask = mask
	return r
}

func (r Catalyst_Company_Type) Filter(filter string) Catalyst_Company_Type {
	r.Options.Filter = filter
	return r
}

func (r Catalyst_Company_Type) Limit(limit int) Catalyst_Company_Type {
	r.Options.Limit = &limit
	return r
}

func (r Catalyst_Company_Type) Offset(offset int) Catalyst_Company_Type {
	r.Options.Offset = &offset
	return r
}

// <<<EOT
func (r Catalyst_Company_Type) GetAllObjects() (resp []datatypes.Catalyst_Company_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r Catalyst_Company_Type) GetObject() (resp datatypes.Catalyst_Company_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Catalyst_Enrollment struct {
	Session *session.Session
	Options sl.Options
}

func GetCatalystEnrollmentService(sess *session.Session) Catalyst_Enrollment {
	return Catalyst_Enrollment{Session: sess}
}

func (r Catalyst_Enrollment) Id(id int) Catalyst_Enrollment {
	r.Options.Id = &id
	return r
}

func (r Catalyst_Enrollment) Mask(mask string) Catalyst_Enrollment {
	r.Options.Mask = mask
	return r
}

func (r Catalyst_Enrollment) Filter(filter string) Catalyst_Enrollment {
	r.Options.Filter = filter
	return r
}

func (r Catalyst_Enrollment) Limit(limit int) Catalyst_Enrollment {
	r.Options.Limit = &limit
	return r
}

func (r Catalyst_Enrollment) Offset(offset int) Catalyst_Enrollment {
	r.Options.Offset = &offset
	return r
}

// Retrieve
func (r Catalyst_Enrollment) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Catalyst_Enrollment) GetAffiliate() (resp datatypes.Catalyst_Affiliate, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r Catalyst_Enrollment) GetAffiliates() (resp []datatypes.Catalyst_Affiliate, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Catalyst_Enrollment) GetCompanyType() (resp datatypes.Catalyst_Company_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r Catalyst_Enrollment) GetCompanyTypes() (resp []datatypes.Catalyst_Company_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r Catalyst_Enrollment) GetEnrollmentRequestAnnualRevenueOptions() (resp []datatypes.Catalyst_Enrollment_Request_Container_AnswerOption, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r Catalyst_Enrollment) GetEnrollmentRequestUserCountOptions() (resp []datatypes.Catalyst_Enrollment_Request_Container_AnswerOption, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r Catalyst_Enrollment) GetEnrollmentRequestYearsInOperationOptions() (resp []datatypes.Catalyst_Enrollment_Request_Container_AnswerOption, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Catalyst_Enrollment) GetIsActiveFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r Catalyst_Enrollment) GetObject() (resp datatypes.Catalyst_Enrollment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Catalyst_Enrollment) GetRepresentative() (resp datatypes.User_Employee, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r Catalyst_Enrollment) RequestManualEnrollment(request *datatypes.Container_Catalyst_ManualEnrollmentRequest) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		request,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r Catalyst_Enrollment) RequestSelfEnrollment(enrollmentRequest *datatypes.Catalyst_Enrollment_Request) (resp datatypes.Account, err error) {
	params := []interface{}{
		enrollmentRequest,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
