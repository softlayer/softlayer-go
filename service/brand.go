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

type Brand struct {
	Session *Session
	Options
}

func (r *Session) GetBrandService() Brand {
	return Brand{Session: r}
}

func (r *Brand) CreateCustomerAccount(account *datatypes.Account, bypassDuplicateAccountCheck *bool) (resp datatypes.Account, err error) {
	params := []interface{}{
		account,
		bypassDuplicateAccountCheck,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Brand) CreateObject(templateObject *datatypes.Brand) (resp datatypes.Brand, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Brand) GetAllTicketSubjects(account *datatypes.Account) (resp []datatypes.Ticket_Subject, err error) {
	params := []interface{}{
		account,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Brand) GetContactInformation() (resp []datatypes.Brand_Contact, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Brand) GetMerchantName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Brand) GetObject() (resp datatypes.Brand, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Brand) GetToken(userId *int) (resp string, err error) {
	params := []interface{}{
		userId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

func (r *Brand) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Brand) GetAllOwnedAccounts() (resp []datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Brand) GetAllowAccountCreationFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Brand) GetCatalog() (resp datatypes.Product_Catalog, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Brand) GetContacts() (resp []datatypes.Brand_Contact, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Brand) GetCustomerCountryLocationRestrictions() (resp []datatypes.Brand_Restriction_Location_CustomerCountry, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Brand) GetDistributor() (resp datatypes.Brand, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Brand) GetDistributorChildFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Brand) GetDistributorFlag() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Brand) GetHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Brand) GetHasAgentSupportFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Brand) GetOpenTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Brand) GetOwnedAccounts() (resp []datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Brand) GetTicketGroups() (resp []datatypes.Ticket_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Brand) GetTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Brand) GetUsers() (resp []datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Brand) GetVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Brand_Restriction_Location_CustomerCountry struct {
	Session *Session
	Options
}

func (r *Session) GetBrandRestrictionLocationCustomerCountryService() Brand_Restriction_Location_CustomerCountry {
	return Brand_Restriction_Location_CustomerCountry{Session: r}
}

func (r *Brand_Restriction_Location_CustomerCountry) GetAllObjects() (resp []datatypes.Brand_Restriction_Location_CustomerCountry, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Brand_Restriction_Location_CustomerCountry) GetObject() (resp datatypes.Brand_Restriction_Location_CustomerCountry, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Brand_Restriction_Location_CustomerCountry) GetBrand() (resp datatypes.Brand, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Brand_Restriction_Location_CustomerCountry) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
