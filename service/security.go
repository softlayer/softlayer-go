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

type Security_Certificate struct {
	Session *Session
	Options
}

func (r *Session) GetSecurityCertificateService() Security_Certificate {
	return Security_Certificate{Session: r}
}

func (r *Security_Certificate) CreateObject(templateObject *datatypes.Security_Certificate) (resp datatypes.Security_Certificate, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate) EditObject(templateObject *datatypes.Security_Certificate) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate) FindByCommonName(commonName *string) (resp []datatypes.Security_Certificate, err error) {
	params := []interface{}{
		commonName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate) GetAssociatedServiceCount() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate) GetLoadBalancerVirtualIpAddresses() (resp []datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate) GetObject() (resp datatypes.Security_Certificate, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate) GetPemFormat() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Security_Certificate_Request struct {
	Session *Session
	Options
}

func (r *Session) GetSecurityCertificateRequestService() Security_Certificate_Request {
	return Security_Certificate_Request{Session: r}
}

func (r *Security_Certificate_Request) CancelSslOrder() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request) GetAdministratorEmailDomains(commonName *string) (resp []string, err error) {
	params := []interface{}{
		commonName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request) GetAdministratorEmailPrefixes() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request) GetCertificateAuthorityName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request) GetObject() (resp datatypes.Security_Certificate_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request) GetOrder() (resp datatypes.Billing_Order, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request) GetOrderItem() (resp datatypes.Billing_Order_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request) GetPreviousOrderData() (resp datatypes.Container_Product_Order_Security_Certificate, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request) GetSslCertificateRequests(accountId *int) (resp []datatypes.Security_Certificate_Request, err error) {
	params := []interface{}{
		accountId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request) GetStatus() (resp datatypes.Security_Certificate_Request_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request) ResendEmail(emailType *string) (resp bool, err error) {
	params := []interface{}{
		emailType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request) ValidateCsr(csr *string, validityMonths *int, itemId *int, serverType *string) (resp bool, err error) {
	params := []interface{}{
		csr,
		validityMonths,
		itemId,
		serverType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Security_Certificate_Request_ServerType struct {
	Session *Session
	Options
}

func (r *Session) GetSecurityCertificateRequestServerTypeService() Security_Certificate_Request_ServerType {
	return Security_Certificate_Request_ServerType{Session: r}
}

func (r *Security_Certificate_Request_ServerType) GetAllObjects() (resp []datatypes.Security_Certificate_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request_ServerType) GetObject() (resp datatypes.Security_Certificate_Request_ServerType, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Security_Certificate_Request_Status struct {
	Session *Session
	Options
}

func (r *Session) GetSecurityCertificateRequestStatusService() Security_Certificate_Request_Status {
	return Security_Certificate_Request_Status{Session: r}
}

func (r *Security_Certificate_Request_Status) GetObject() (resp datatypes.Security_Certificate_Request_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request_Status) GetSslRequestStatuses() (resp []datatypes.Security_Certificate_Request_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Security_Ssh_Key struct {
	Session *Session
	Options
}

func (r *Session) GetSecuritySshKeyService() Security_Ssh_Key {
	return Security_Ssh_Key{Session: r}
}

func (r *Security_Ssh_Key) CreateObject(templateObject *datatypes.Security_Ssh_Key) (resp datatypes.Security_Ssh_Key, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Ssh_Key) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Ssh_Key) EditObject(templateObject *datatypes.Security_Ssh_Key) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Ssh_Key) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Ssh_Key) GetBlockDeviceTemplateGroups() (resp []datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Ssh_Key) GetObject() (resp datatypes.Security_Ssh_Key, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Ssh_Key) GetSoftwarePasswords() (resp []datatypes.Software_Component_Password, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
