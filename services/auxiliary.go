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
type Auxiliary_Marketing_Event struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryMarketingEventService() Auxiliary_Marketing_Event {
	return Auxiliary_Marketing_Event{Session: r}
}

// This method will return a collection of SoftLayer_Auxiliary_Marketing_Event objects ordered in ascending order by start date.
func (r *Auxiliary_Marketing_Event) GetMarketingEvents() (resp []datatypes.Auxiliary_Marketing_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Auxiliary_Marketing_Event) GetObject() (resp datatypes.Auxiliary_Marketing_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Auxiliary_Network_Status struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryNetworkStatusService() Auxiliary_Network_Status {
	return Auxiliary_Network_Status{Session: r}
}

// Return the current network status of and latency information for a given target from numerous points around the world. Valid Targets:
// * ALL
// * NETWORK_DALLAS
// * NETWORK_SEATTLE
// * NETWORK_PUBLIC
// * NETWORK_PUBLIC_DALLAS
// * NETWORK_PUBLIC_SEATTLE
// * NETWORK_PUBLIC_WDC
// * NETWORK_PRIVATE
// * NETWORK_PRIVATE_DALLAS
// * NETWORK_PRIVATE_SEATTLE
// * NETWORK_PRIVATE_WDC
func (r *Auxiliary_Network_Status) GetNetworkStatus(target *string) (resp []datatypes.Container_Auxiliary_Network_Status_Reading, err error) {
	params := []interface{}{
		target,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// A SoftLayer_Auxiliary_Notification_Emergency data object represents a notification event being broadcast to the SoftLayer customer base. It is used to provide information regarding outages or current known issues.
type Auxiliary_Notification_Emergency struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryNotificationEmergencyService() Auxiliary_Notification_Emergency {
	return Auxiliary_Notification_Emergency{Session: r}
}

// Retrieve an array of SoftLayer_Auxiliary_Notification_Emergency data types, which contain all notification events regardless of status.
func (r *Auxiliary_Notification_Emergency) GetAllObjects() (resp []datatypes.Auxiliary_Notification_Emergency, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve an array of SoftLayer_Auxiliary_Notification_Emergency data types, which contain all current notification events.
func (r *Auxiliary_Notification_Emergency) GetCurrentNotifications() (resp []datatypes.Auxiliary_Notification_Emergency, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Auxiliary_Notification_Emergency object, it can be used to check for current notifications being broadcast by SoftLayer.
func (r *Auxiliary_Notification_Emergency) GetObject() (resp datatypes.Auxiliary_Notification_Emergency, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The signature of the SoftLayer employee department associated with this notification.
func (r *Auxiliary_Notification_Emergency) GetSignature() (resp datatypes.Auxiliary_Notification_Emergency_Signature, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The status of this notification.
func (r *Auxiliary_Notification_Emergency) GetStatus() (resp datatypes.Auxiliary_Notification_Emergency_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Auxiliary_Press_Release struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryPressReleaseService() Auxiliary_Press_Release {
	return Auxiliary_Press_Release{Session: r}
}

// Retrieve
func (r *Auxiliary_Press_Release) GetAbout() (resp []datatypes.Auxiliary_Press_Release_About_Press_Release, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve an array of SoftLayer_Auxiliary_Press_Release data types, which contain all press releases.
func (r *Auxiliary_Press_Release) GetAllObjects() (resp []datatypes.Auxiliary_Press_Release, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Auxiliary_Press_Release) GetContacts() (resp []datatypes.Auxiliary_Press_Release_Contact_Press_Release, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Auxiliary_Press_Release) GetMediaPartners() (resp []datatypes.Auxiliary_Press_Release_Media_Partner_Press_Release, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Auxiliary_Press_Release object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Auxiliary_Press_Release service.
func (r *Auxiliary_Press_Release) GetObject() (resp datatypes.Auxiliary_Press_Release, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Auxiliary_Press_Release) GetPressReleaseContent() (resp datatypes.Auxiliary_Press_Release_Content, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve an array of SoftLayer_Auxiliary_Press_Release data types, which contain all press releases.
func (r *Auxiliary_Press_Release) GetRenderedPressRelease() (resp []datatypes.Auxiliary_Press_Release, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve an array of SoftLayer_Auxiliary_Press_Release data types, which contain all press releases for a given year and or result limit.
func (r *Auxiliary_Press_Release) GetRenderedPressReleases(resultLimit *string, year *string) (resp []datatypes.Auxiliary_Press_Release, err error) {
	params := []interface{}{
		resultLimit,
		year,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve an array of SoftLayer_Auxiliary_Press_Release data types, which have the website highlight flag set.
func (r *Auxiliary_Press_Release) GetWebsiteHighlightPressReleases() (resp []datatypes.Auxiliary_Press_Release, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Auxiliary_Press_Release_About struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryPressReleaseAboutService() Auxiliary_Press_Release_About {
	return Auxiliary_Press_Release_About{Session: r}
}

// getObject retrieves the SoftLayer_Auxiliary_Press_Release_About object whose about id number corresponds to the ID number of the init parameter passed to the SoftLayer_Auxiliary_Press_Release service.
func (r *Auxiliary_Press_Release_About) GetObject() (resp datatypes.Auxiliary_Press_Release_About, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Auxiliary_Press_Release_About_Press_Release struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryPressReleaseAboutPressReleaseService() Auxiliary_Press_Release_About_Press_Release {
	return Auxiliary_Press_Release_About_Press_Release{Session: r}
}

// Retrieve
func (r *Auxiliary_Press_Release_About_Press_Release) GetAboutParagraphs() (resp []datatypes.Auxiliary_Press_Release_About, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Auxiliary_Press_Release_About_Press_Release object whose contact id number corresponds to the ID number of the init parameter passed to the SoftLayer_Auxiliary_Press_Release service.
func (r *Auxiliary_Press_Release_About_Press_Release) GetObject() (resp datatypes.Auxiliary_Press_Release_About_Press_Release, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Auxiliary_Press_Release_About_Press_Release) GetPressReleases() (resp []datatypes.Auxiliary_Press_Release, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Auxiliary_Press_Release_Contact struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryPressReleaseContactService() Auxiliary_Press_Release_Contact {
	return Auxiliary_Press_Release_Contact{Session: r}
}

// getObject retrieves the SoftLayer_Auxiliary_Press_Release_Contact object whose contact id number corresponds to the ID number of the init parameter passed to the SoftLayer_Auxiliary_Press_Release service.
func (r *Auxiliary_Press_Release_Contact) GetObject() (resp datatypes.Auxiliary_Press_Release_Contact, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Auxiliary_Press_Release_Contact_Press_Release struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryPressReleaseContactPressReleaseService() Auxiliary_Press_Release_Contact_Press_Release {
	return Auxiliary_Press_Release_Contact_Press_Release{Session: r}
}

// Retrieve
func (r *Auxiliary_Press_Release_Contact_Press_Release) GetContacts() (resp []datatypes.Auxiliary_Press_Release_Contact, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Auxiliary_Press_Release_Contact object whose contact id number corresponds to the ID number of the init parameter passed to the SoftLayer_Auxiliary_Press_Release service.
func (r *Auxiliary_Press_Release_Contact_Press_Release) GetObject() (resp datatypes.Auxiliary_Press_Release_Contact_Press_Release, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Auxiliary_Press_Release_Contact_Press_Release) GetPressReleases() (resp []datatypes.Auxiliary_Press_Release, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Auxiliary_Press_Release_Content struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryPressReleaseContentService() Auxiliary_Press_Release_Content {
	return Auxiliary_Press_Release_Content{Session: r}
}

// getObject retrieves the SoftLayer_Auxiliary_Press_Release_Content object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Auxiliary_Press_Release service.
func (r *Auxiliary_Press_Release_Content) GetObject() (resp datatypes.Auxiliary_Press_Release_Content, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Auxiliary_Press_Release_Media_Partner struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryPressReleaseMediaPartnerService() Auxiliary_Press_Release_Media_Partner {
	return Auxiliary_Press_Release_Media_Partner{Session: r}
}

// getObject retrieves the SoftLayer_Auxiliary_Press_Release_Contact object whose contact id number corresponds to the ID number of the init parameter passed to the SoftLayer_Auxiliary_Press_Release service.
func (r *Auxiliary_Press_Release_Media_Partner) GetObject() (resp datatypes.Auxiliary_Press_Release_Media_Partner, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Auxiliary_Press_Release_Media_Partner_Press_Release struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryPressReleaseMediaPartnerPressReleaseService() Auxiliary_Press_Release_Media_Partner_Press_Release {
	return Auxiliary_Press_Release_Media_Partner_Press_Release{Session: r}
}

// Retrieve
func (r *Auxiliary_Press_Release_Media_Partner_Press_Release) GetMediaPartners() (resp []datatypes.Auxiliary_Press_Release_Media_Partner, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Auxiliary_Press_Release_Media_Partner_Press_Release object whose media partner id number corresponds to the ID number of the init parameter passed to the SoftLayer_Auxiliary_Press_Release service.
func (r *Auxiliary_Press_Release_Media_Partner_Press_Release) GetObject() (resp datatypes.Auxiliary_Press_Release_Media_Partner_Press_Release, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Auxiliary_Press_Release_Media_Partner_Press_Release) GetPressReleases() (resp []datatypes.Auxiliary_Press_Release, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Auxiliary_Shipping_Courier_Type struct {
	Session *Session
	Options
}

func (r *Session) GetAuxiliaryShippingCourierTypeService() Auxiliary_Shipping_Courier_Type {
	return Auxiliary_Shipping_Courier_Type{Session: r}
}

// Retrieve
func (r *Auxiliary_Shipping_Courier_Type) GetCourier() (resp []datatypes.Auxiliary_Shipping_Courier, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Auxiliary_Shipping_Courier_Type) GetObject() (resp datatypes.Auxiliary_Shipping_Courier_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Auxiliary_Shipping_Courier_Type) GetTypeByKeyName(keyName *string) (resp datatypes.Auxiliary_Shipping_Courier_Type, err error) {
	params := []interface{}{
		keyName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
