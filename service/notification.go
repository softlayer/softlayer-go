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

type Notification struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationService() Notification {
	return Notification{Session: r}
}

func (r *Notification) GetAllObjects() (resp []datatypes.Notification, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification) GetObject() (resp datatypes.Notification, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Notification) GetPreferences() (resp []datatypes.Notification_Preference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification) GetRequiredPreferences() (resp []datatypes.Notification_Preference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Notification_Delivery_Method struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationDeliveryMethodService() Notification_Delivery_Method {
	return Notification_Delivery_Method{Session: r}
}

type Notification_Mobile struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationMobileService() Notification_Mobile {
	return Notification_Mobile{Session: r}
}

func (r *Notification_Mobile) CreateSubscriberForMobileDevice(keyName *string, resourceTableId *int, userRecordId *int) (resp bool, err error) {
	params := []interface{}{
		keyName,
		resourceTableId,
		userRecordId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Mobile) GetObject() (resp datatypes.Notification_Mobile, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Notification_Occurrence_Account struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationOccurrenceAccountService() Notification_Occurrence_Account {
	return Notification_Occurrence_Account{Session: r}
}

func (r *Notification_Occurrence_Account) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Occurrence_Account) GetLastNotificationUpdate() (resp datatypes.Notification_Occurrence_Update, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Occurrence_Account) GetNotificationOccurrenceEvent() (resp datatypes.Notification_Occurrence_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Notification_Occurrence_Event struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationOccurrenceEventService() Notification_Occurrence_Event {
	return Notification_Occurrence_Event{Session: r}
}

func (r *Notification_Occurrence_Event) AcknowledgeNotification() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Occurrence_Event) GetAllObjects() (resp []datatypes.Notification_Occurrence_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Occurrence_Event) GetAttachedFile(attachmentId *int) (resp []byte, err error) {
	params := []interface{}{
		attachmentId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Occurrence_Event) GetImpactedAccountCount() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Occurrence_Event) GetImpactedDeviceCount() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Occurrence_Event) GetImpactedDevices() (resp []datatypes.Notification_Occurrence_Resource, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Occurrence_Event) GetObject() (resp datatypes.Notification_Occurrence_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Notification_Occurrence_Event) GetAcknowledgedFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Occurrence_Event) GetAttachments() (resp []datatypes.Notification_Occurrence_Event_Attachment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Occurrence_Event) GetFirstUpdate() (resp datatypes.Notification_Occurrence_Update, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Occurrence_Event) GetImpactedAccounts() (resp []datatypes.Notification_Occurrence_Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Occurrence_Event) GetImpactedResources() (resp []datatypes.Notification_Occurrence_Resource, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Occurrence_Event) GetImpactedUsers() (resp []datatypes.Notification_Occurrence_User, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Occurrence_Event) GetLastUpdate() (resp datatypes.Notification_Occurrence_Update, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Occurrence_Event) GetNotificationOccurrenceEventType() (resp datatypes.Notification_Occurrence_Event_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Occurrence_Event) GetStatusCode() (resp datatypes.Notification_Occurrence_Status_Code, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Occurrence_Event) GetUpdates() (resp []datatypes.Notification_Occurrence_Update, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Notification_Occurrence_Event_Attachment struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationOccurrenceEventAttachmentService() Notification_Occurrence_Event_Attachment {
	return Notification_Occurrence_Event_Attachment{Session: r}
}

func (r *Notification_Occurrence_Event_Attachment) GetNotificationOccurrenceEvent() (resp datatypes.Notification_Occurrence_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Notification_Occurrence_Event_Type struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationOccurrenceEventTypeService() Notification_Occurrence_Event_Type {
	return Notification_Occurrence_Event_Type{Session: r}
}

type Notification_Occurrence_Resource struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationOccurrenceResourceService() Notification_Occurrence_Resource {
	return Notification_Occurrence_Resource{Session: r}
}

func (r *Notification_Occurrence_Resource) GetNotificationOccurrenceEvent() (resp datatypes.Notification_Occurrence_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Occurrence_Resource) GetResource() (resp datatypes.Entity, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Notification_Occurrence_Resource_Hardware struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationOccurrenceResourceHardwareService() Notification_Occurrence_Resource_Hardware {
	return Notification_Occurrence_Resource_Hardware{Session: r}
}

type Notification_Occurrence_Resource_Network_Application_Delivery_Controller struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationOccurrenceResourceNetworkApplicationDeliveryControllerService() Notification_Occurrence_Resource_Network_Application_Delivery_Controller {
	return Notification_Occurrence_Resource_Network_Application_Delivery_Controller{Session: r}
}

type Notification_Occurrence_Resource_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationOccurrenceResourceNetworkApplicationDeliveryControllerLoadBalancerVirtualIpAddressService() Notification_Occurrence_Resource_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress {
	return Notification_Occurrence_Resource_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress{Session: r}
}

type Notification_Occurrence_Resource_Network_Storage_Iscsi_EqualLogic struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationOccurrenceResourceNetworkStorageIscsiEqualLogicService() Notification_Occurrence_Resource_Network_Storage_Iscsi_EqualLogic {
	return Notification_Occurrence_Resource_Network_Storage_Iscsi_EqualLogic{Session: r}
}

type Notification_Occurrence_Resource_Network_Storage_Iscsi_NetApp struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationOccurrenceResourceNetworkStorageIscsiNetAppService() Notification_Occurrence_Resource_Network_Storage_Iscsi_NetApp {
	return Notification_Occurrence_Resource_Network_Storage_Iscsi_NetApp{Session: r}
}

type Notification_Occurrence_Resource_Network_Storage_Lockbox struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationOccurrenceResourceNetworkStorageLockboxService() Notification_Occurrence_Resource_Network_Storage_Lockbox {
	return Notification_Occurrence_Resource_Network_Storage_Lockbox{Session: r}
}

type Notification_Occurrence_Resource_Network_Storage_Nas struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationOccurrenceResourceNetworkStorageNasService() Notification_Occurrence_Resource_Network_Storage_Nas {
	return Notification_Occurrence_Resource_Network_Storage_Nas{Session: r}
}

type Notification_Occurrence_Resource_Network_Storage_NetApp_Volume struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationOccurrenceResourceNetworkStorageNetAppVolumeService() Notification_Occurrence_Resource_Network_Storage_NetApp_Volume {
	return Notification_Occurrence_Resource_Network_Storage_NetApp_Volume{Session: r}
}

type Notification_Occurrence_Resource_Virtual struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationOccurrenceResourceVirtualService() Notification_Occurrence_Resource_Virtual {
	return Notification_Occurrence_Resource_Virtual{Session: r}
}

type Notification_Occurrence_Status_Code struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationOccurrenceStatusCodeService() Notification_Occurrence_Status_Code {
	return Notification_Occurrence_Status_Code{Session: r}
}

type Notification_Occurrence_Update struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationOccurrenceUpdateService() Notification_Occurrence_Update {
	return Notification_Occurrence_Update{Session: r}
}

func (r *Notification_Occurrence_Update) GetEmployee() (resp datatypes.User_Employee, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Occurrence_Update) GetNotificationOccurrenceEvent() (resp datatypes.Notification_Occurrence_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Notification_Occurrence_User struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationOccurrenceUserService() Notification_Occurrence_User {
	return Notification_Occurrence_User{Session: r}
}

func (r *Notification_Occurrence_User) Acknowledge() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Occurrence_User) GetAllObjects() (resp []datatypes.Notification_Occurrence_User, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Occurrence_User) GetImpactedDeviceCount() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Occurrence_User) GetObject() (resp datatypes.Notification_Occurrence_User, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Notification_Occurrence_User) GetImpactedResources() (resp []datatypes.Notification_Occurrence_Resource, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Occurrence_User) GetNotificationOccurrenceEvent() (resp datatypes.Notification_Occurrence_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Occurrence_User) GetUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Notification_Preference struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationPreferenceService() Notification_Preference {
	return Notification_Preference{Session: r}
}

type Notification_Subscriber struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationSubscriberService() Notification_Subscriber {
	return Notification_Subscriber{Session: r}
}

func (r *Notification_Subscriber) GetDeliveryMethods() (resp []datatypes.Notification_Subscriber_Delivery_Method, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Subscriber) GetNotification() (resp datatypes.Notification, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Notification_Subscriber_Customer struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationSubscriberCustomerService() Notification_Subscriber_Customer {
	return Notification_Subscriber_Customer{Session: r}
}

func (r *Notification_Subscriber_Customer) GetSubscriberRecord() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Notification_Subscriber_Delivery_Method struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationSubscriberDeliveryMethodService() Notification_Subscriber_Delivery_Method {
	return Notification_Subscriber_Delivery_Method{Session: r}
}

func (r *Notification_Subscriber_Delivery_Method) GetNotificationDeliveryMethod() (resp datatypes.Notification_Delivery_Method, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_Subscriber_Delivery_Method) GetNotificationSubscriber() (resp datatypes.Notification_Subscriber, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Notification_User_Subscriber struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationUserSubscriberService() Notification_User_Subscriber {
	return Notification_User_Subscriber{Session: r}
}

func (r *Notification_User_Subscriber) CreateObject(templateObject *datatypes.Notification_User_Subscriber) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_User_Subscriber) EditObject(templateObject *datatypes.Notification_User_Subscriber) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_User_Subscriber) GetObject() (resp datatypes.Notification_User_Subscriber, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Notification_User_Subscriber) GetDeliveryMethods() (resp []datatypes.Notification_Delivery_Method, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_User_Subscriber) GetNotification() (resp datatypes.Notification, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_User_Subscriber) GetPreferences() (resp []datatypes.Notification_User_Subscriber_Preference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_User_Subscriber) GetPreferencesDetails() (resp []datatypes.Notification_Preference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_User_Subscriber) GetResourceRecord() (resp datatypes.Notification_User_Subscriber_Resource, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_User_Subscriber) GetUserRecord() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Notification_User_Subscriber_Billing struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationUserSubscriberBillingService() Notification_User_Subscriber_Billing {
	return Notification_User_Subscriber_Billing{Session: r}
}

func (r *Notification_User_Subscriber_Billing) GetObject() (resp datatypes.Notification_User_Subscriber_Billing, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Notification_User_Subscriber_Delivery_Method struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationUserSubscriberDeliveryMethodService() Notification_User_Subscriber_Delivery_Method {
	return Notification_User_Subscriber_Delivery_Method{Session: r}
}

func (r *Notification_User_Subscriber_Delivery_Method) GetDeliveryMethod() (resp datatypes.Notification_Delivery_Method, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_User_Subscriber_Delivery_Method) GetNotificationUserSubscriber() (resp datatypes.Notification_User_Subscriber, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Notification_User_Subscriber_Mobile struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationUserSubscriberMobileService() Notification_User_Subscriber_Mobile {
	return Notification_User_Subscriber_Mobile{Session: r}
}

func (r *Notification_User_Subscriber_Mobile) ClearSnoozeTimer() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_User_Subscriber_Mobile) GetObject() (resp datatypes.Notification_User_Subscriber_Mobile, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_User_Subscriber_Mobile) SetSnoozeTimer(start *int, end *int) (resp bool, err error) {
	params := []interface{}{
		start,
		end,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Notification_User_Subscriber_Preference struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationUserSubscriberPreferenceService() Notification_User_Subscriber_Preference {
	return Notification_User_Subscriber_Preference{Session: r}
}

func (r *Notification_User_Subscriber_Preference) CreateObject(templateObject *datatypes.Notification_User_Subscriber_Preference) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_User_Subscriber_Preference) EditObjects(templateObjects []datatypes.Notification_User_Subscriber_Preference) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_User_Subscriber_Preference) GetObject() (resp datatypes.Notification_User_Subscriber_Preference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Notification_User_Subscriber_Preference) GetDefaultPreference() (resp datatypes.Notification_Preference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Notification_User_Subscriber_Preference) GetNotificationUserSubscriber() (resp datatypes.Notification_User_Subscriber, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Notification_User_Subscriber_Resource struct {
	Session *Session
	Options
}

func (r *Session) GetNotificationUserSubscriberResourceService() Notification_User_Subscriber_Resource {
	return Notification_User_Subscriber_Resource{Session: r}
}

func (r *Notification_User_Subscriber_Resource) GetNotificationUserSubscriber() (resp datatypes.Notification_User_Subscriber, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
