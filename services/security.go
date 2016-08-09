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

//
type Security_Certificate struct {
	Session *Session
	Options
}

func (r *Session) GetSecurityCertificateService() Security_Certificate {
	return Security_Certificate{Session: r}
}

// Add a certificate to your account for your records, or for use with various services. Only the certificate and private key are usually required. If your issuer provided an intermediate certificate, you must also provide that certificate. Details will be extracted from the certificate. Validation will be performed between the certificate and the private key as well as the certificate and the intermediate certificate, if provided.
//
// The certificate signing request is not required, but can be provided for your records.
func (r *Security_Certificate) CreateObject(templateObject *datatypes.Security_Certificate) (resp datatypes.Security_Certificate, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Remove a certificate from your account. You may not remove a certificate with associated services.
func (r *Security_Certificate) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Update a certificate. Modifications are restricted to the note and CSR if the are any services associated with the certificate. There are no modification restrictions for a certificate with no associated services.
func (r *Security_Certificate) EditObject(templateObject *datatypes.Security_Certificate) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Locate certificates by their common name, traditionally a domain name.
func (r *Security_Certificate) FindByCommonName(commonName *string) (resp []datatypes.Security_Certificate, err error) {
	params := []interface{}{
		commonName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The number of services currently associated with the certificate.
func (r *Security_Certificate) GetAssociatedServiceCount() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The load balancers virtual IP addresses currently associated with the certificate.
func (r *Security_Certificate) GetLoadBalancerVirtualIpAddresses() (resp []datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Security_Certificate) GetObject() (resp datatypes.Security_Certificate, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the certificate in PEM (Privacy Enhanced Mail) format, which is a string containing all base64 encoded (DER) certificates delimited by -----BEGIN/END *----- clauses.
func (r *Security_Certificate) GetPemFormat() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// SoftLayer_Security_Certificate_Request data type is used to harness your SSL certificate order to a Certificate Authority. This contains data that is required by a Certificate Authority to place an SSL certificate order.
type Security_Certificate_Request struct {
	Session *Session
	Options
}

func (r *Session) GetSecurityCertificateRequestService() Security_Certificate_Request {
	return Security_Certificate_Request{Session: r}
}

// Cancels a pending SSL certificate order at Certificate Authority
func (r *Security_Certificate_Request) CancelSslOrder() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The account to which a SSL certificate request belongs.
func (r *Security_Certificate_Request) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Gets the email domains that can be used to validate a certificate to a domain.
func (r *Security_Certificate_Request) GetAdministratorEmailDomains(commonName *string) (resp []string, err error) {
	params := []interface{}{
		commonName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Gets the email accounts that can be used to validate a certificate to a domain.
func (r *Security_Certificate_Request) GetAdministratorEmailPrefixes() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The Certificate Authority name
func (r *Security_Certificate_Request) GetCertificateAuthorityName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Security_Certificate_Request) GetObject() (resp datatypes.Security_Certificate_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The order contains the information related to a SSL certificate request.
func (r *Security_Certificate_Request) GetOrder() (resp datatypes.Billing_Order, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The associated order item for this SSL certificate request.
func (r *Security_Certificate_Request) GetOrderItem() (resp datatypes.Billing_Order_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Returns previous SSL certificate order data. You can use this data for to place a renewal order for a completed SSL certificate.
func (r *Security_Certificate_Request) GetPreviousOrderData() (resp datatypes.Container_Product_Order_Security_Certificate, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Returns all the SSL certificate requests.
func (r *Security_Certificate_Request) GetSslCertificateRequests(accountId *int) (resp []datatypes.Security_Certificate_Request, err error) {
	params := []interface{}{
		accountId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The status of a SSL certificate request.
func (r *Security_Certificate_Request) GetStatus() (resp datatypes.Security_Certificate_Request_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// A Certificate Authority sends out various emails to your domain administrator or your technical contact. Use this service to re-send these emails if you did not receive them initially.
func (r *Security_Certificate_Request) ResendEmail(emailType *string) (resp bool, err error) {
	params := []interface{}{
		emailType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Allows you to validate a Certificate Signing Request (CSR) required for an SSL certificate with the certificate authority (CA).  This method sends the CSR, the length of the subscription in months, the certificate type, and the server type for validation against requirements of the CA.  Returns true if valid.
//
// More information on CSR generation can be found at: [http://en.wikipedia.org/wiki/Certificate_signing_request Wikipedia] [https://knowledge.verisign.com/support/ssl-certificates-support/index?page=content&id=AR235&actp=LIST&viewlocale=en_US VeriSign]
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

// Represents a server type that can be specified when ordering an SSL certificate.
type Security_Certificate_Request_ServerType struct {
	Session *Session
	Options
}

func (r *Session) GetSecurityCertificateRequestServerTypeService() Security_Certificate_Request_ServerType {
	return Security_Certificate_Request_ServerType{Session: r}
}

// Returns all SSL certificate server types, which  passed in on a certificate order.
func (r *Security_Certificate_Request_ServerType) GetAllObjects() (resp []datatypes.Security_Certificate_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Security_Certificate_Request_ServerType) GetObject() (resp datatypes.Security_Certificate_Request_ServerType, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// SoftLayer_Security_Certificate_Request_Status data type represents the status of an SSL certificate request.
type Security_Certificate_Request_Status struct {
	Session *Session
	Options
}

func (r *Session) GetSecurityCertificateRequestStatusService() Security_Certificate_Request_Status {
	return Security_Certificate_Request_Status{Session: r}
}

//
func (r *Security_Certificate_Request_Status) GetObject() (resp datatypes.Security_Certificate_Request_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Returns all SSL certificate request status objects
func (r *Security_Certificate_Request_Status) GetSslRequestStatuses() (resp []datatypes.Security_Certificate_Request_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Security_Ssh_Key struct {
	Session *Session
	Options
}

func (r *Session) GetSecuritySshKeyService() Security_Ssh_Key {
	return Security_Ssh_Key{Session: r}
}

// Add a ssh key to your account for use during server provisioning and os reloads.
func (r *Security_Ssh_Key) CreateObject(templateObject *datatypes.Security_Ssh_Key) (resp datatypes.Security_Ssh_Key, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Remove a ssh key from your account.
func (r *Security_Ssh_Key) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Update a ssh key.
func (r *Security_Ssh_Key) EditObject(templateObject *datatypes.Security_Ssh_Key) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Security_Ssh_Key) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The image template groups that are linked to an SSH key.
func (r *Security_Ssh_Key) GetBlockDeviceTemplateGroups() (resp []datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Security_Ssh_Key) GetObject() (resp datatypes.Security_Ssh_Key, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The OS root users that are linked to an SSH key.
func (r *Security_Ssh_Key) GetSoftwarePasswords() (resp []datatypes.Software_Component_Password, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
