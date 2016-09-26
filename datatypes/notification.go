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

package datatypes

// Details provided for the notification are basic.  Details such as the related preferences, name and keyname for the notification can be retrieved.  The keyname property for the notification can be used to refer to a notification when integrating into the SoftLayer Notification system.  The name property can used more for display purposes.
type Notification struct {
	Entity

	// Unique identifier for the notification.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// Name that can be used by external systems to refer to a notification.
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// Friendly name for the notification.
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// A count of the preferences related to the notification. These are preferences are configurable and optional for subscribers to use.
	PreferenceCount *uint `json:"preferenceCount,omitempty" xmlrpc:"preferenceCount"`

	// The preferences related to the notification. These are preferences are configurable and optional for subscribers to use.
	Preferences []Notification_Preference `json:"preferences,omitempty" xmlrpc:"preferences"`

	// A count of the required preferences related to the notification. While configurable, the subscriber does not have the option whether to use the preference.
	RequiredPreferenceCount *uint `json:"requiredPreferenceCount,omitempty" xmlrpc:"requiredPreferenceCount"`

	// The required preferences related to the notification. While configurable, the subscriber does not have the option whether to use the preference.
	RequiredPreferences []Notification_Preference `json:"requiredPreferences,omitempty" xmlrpc:"requiredPreferences"`
}

// Provides details for the delivery methods available.
type Notification_Delivery_Method struct {
	Entity

	// Determines if the delivery method is still used by the system.
	Active *int `json:"active,omitempty" xmlrpc:"active"`

	// Description used for the delivery method.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// Unique identifier for the various notification delivery methods.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// Name that can be used by external systems to refer to delivery method.
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// Friendly name used for the delivery method.
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// This is an extension of the SoftLayer_Notification class.  These are implementation details specific to those notifications which can be subscribed to and received on a mobile device.
type Notification_Mobile struct {
	Notification
}

// no documentation yet
type Notification_Occurrence_Account struct {
	Entity

	// no documentation yet
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// no documentation yet
	Active *int `json:"active,omitempty" xmlrpc:"active"`

	// no documentation yet
	LastNotificationUpdate *Notification_Occurrence_Update `json:"lastNotificationUpdate,omitempty" xmlrpc:"lastNotificationUpdate"`

	// no documentation yet
	NotificationOccurrenceEvent *Notification_Occurrence_Event `json:"notificationOccurrenceEvent,omitempty" xmlrpc:"notificationOccurrenceEvent"`
}

// no documentation yet
type Notification_Occurrence_Event struct {
	Entity

	// Indicates whether or not this event has been acknowledged by the user.
	AcknowledgedFlag *bool `json:"acknowledgedFlag,omitempty" xmlrpc:"acknowledgedFlag"`

	// A count of a collection of attachments for this event which provide supplementary information to impacted users some examples are RFO (Reason For Outage) and root cause analysis documents.
	AttachmentCount *uint `json:"attachmentCount,omitempty" xmlrpc:"attachmentCount"`

	// A collection of attachments for this event which provide supplementary information to impacted users some examples are RFO (Reason For Outage) and root cause analysis documents.
	Attachments []Notification_Occurrence_Event_Attachment `json:"attachments,omitempty" xmlrpc:"attachments"`

	// When this event will end.
	EndDate *Time `json:"endDate,omitempty" xmlrpc:"endDate"`

	// The first update for this event.
	FirstUpdate *Notification_Occurrence_Update `json:"firstUpdate,omitempty" xmlrpc:"firstUpdate"`

	// Unique identifier for this event.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// A count of a collection of accounts impacted by this event. Each impacted account record relates directly to a [[SoftLayer_Account]].
	ImpactedAccountCount *uint `json:"impactedAccountCount,omitempty" xmlrpc:"impactedAccountCount"`

	// A collection of accounts impacted by this event. Each impacted account record relates directly to a [[SoftLayer_Account]].
	ImpactedAccounts []Notification_Occurrence_Account `json:"impactedAccounts,omitempty" xmlrpc:"impactedAccounts"`

	// A count of a collection of resources impacted by this event. Each record will relate to some physical resource that the user has access to such as [[SoftLayer_Hardware]] or [[SoftLayer_Virtual_Guest]].
	ImpactedResourceCount *uint `json:"impactedResourceCount,omitempty" xmlrpc:"impactedResourceCount"`

	// A collection of resources impacted by this event. Each record will relate to some physical resource that the user has access to such as [[SoftLayer_Hardware]] or [[SoftLayer_Virtual_Guest]].
	ImpactedResources []Notification_Occurrence_Resource `json:"impactedResources,omitempty" xmlrpc:"impactedResources"`

	// A count of a collection of users impacted by this event. Each impacted user record relates directly to a [[SoftLayer_User_Customer]].
	ImpactedUserCount *uint `json:"impactedUserCount,omitempty" xmlrpc:"impactedUserCount"`

	// A collection of users impacted by this event. Each impacted user record relates directly to a [[SoftLayer_User_Customer]].
	ImpactedUsers []Notification_Occurrence_User `json:"impactedUsers,omitempty" xmlrpc:"impactedUsers"`

	// Latest count of users impacted by this event.
	LastImpactedUserCount *int `json:"lastImpactedUserCount,omitempty" xmlrpc:"lastImpactedUserCount"`

	// The last update for this event.
	LastUpdate *Notification_Occurrence_Update `json:"lastUpdate,omitempty" xmlrpc:"lastUpdate"`

	// When this event was last updated.
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// The type of event such as planned or unplanned maintenance.
	NotificationOccurrenceEventType *Notification_Occurrence_Event_Type `json:"notificationOccurrenceEventType,omitempty" xmlrpc:"notificationOccurrenceEventType"`

	// no documentation yet
	RecoveryTime *int `json:"recoveryTime,omitempty" xmlrpc:"recoveryTime"`

	// When this event started.
	StartDate *Time `json:"startDate,omitempty" xmlrpc:"startDate"`

	// no documentation yet
	StatusCode *Notification_Occurrence_Status_Code `json:"statusCode,omitempty" xmlrpc:"statusCode"`

	// Brief description of this event.
	Subject *string `json:"subject,omitempty" xmlrpc:"subject"`

	// Details of this event.
	Summary *string `json:"summary,omitempty" xmlrpc:"summary"`

	// Unique identifier for the [[SoftLayer_Ticket]] associated with this event.
	SystemTicketId *int `json:"systemTicketId,omitempty" xmlrpc:"systemTicketId"`

	// A count of all updates for this event.
	UpdateCount *uint `json:"updateCount,omitempty" xmlrpc:"updateCount"`

	// All updates for this event.
	Updates []Notification_Occurrence_Update `json:"updates,omitempty" xmlrpc:"updates"`
}

// SoftLayer events can have have files attached to them by a SoftLayer employee. Attaching a file to a event is a way to provide supplementary information such as a RFO (reason for outage) document or root cause analysis. The SoftLayer_Notification_Occurrence_Event_Attachment data type models a single file attached to a event.
type Notification_Occurrence_Event_Attachment struct {
	Entity

	// The date the file was attached to the event.
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// The name of the file attached to the event.
	FileName *string `json:"fileName,omitempty" xmlrpc:"fileName"`

	// The size of the file, measured in bytes.
	FileSize *string `json:"fileSize,omitempty" xmlrpc:"fileSize"`

	// A event attachments' unique identifier.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	NotificationOccurrenceEvent *Notification_Occurrence_Event `json:"notificationOccurrenceEvent,omitempty" xmlrpc:"notificationOccurrenceEvent"`

	// The unique event identifier that the file is attached to.
	NotificationOccurrenceEventId *int `json:"notificationOccurrenceEventId,omitempty" xmlrpc:"notificationOccurrenceEventId"`
}

// This represents the type of SoftLayer_Notification_Occurrence_Event.
type Notification_Occurrence_Event_Type struct {
	Entity

	// The friendly unique identifier for this event type.
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`
}

// This type contains general information relating to any hardware or services that may be impacted by a SoftLayer_Notification_Occurrence_Event.
type Notification_Occurrence_Resource struct {
	Entity

	// no documentation yet
	Active *int `json:"active,omitempty" xmlrpc:"active"`

	// <<< EOT A label which gives some background as to what piece of
	FilterLabel *string `json:"filterLabel,omitempty" xmlrpc:"filterLabel"`

	// The associated event.
	NotificationOccurrenceEvent *Notification_Occurrence_Event `json:"notificationOccurrenceEvent,omitempty" xmlrpc:"notificationOccurrenceEvent"`

	// <<< EOT The unique identifier for the associated
	NotificationOccurrenceEventId *int `json:"notificationOccurrenceEventId,omitempty" xmlrpc:"notificationOccurrenceEventId"`

	// The physical resource.
	Resource *Entity `json:"resource,omitempty" xmlrpc:"resource"`

	// <<< EOT The unique identifier for the [[SoftLayer_Account]] associated with
	ResourceAccountId *int `json:"resourceAccountId,omitempty" xmlrpc:"resourceAccountId"`

	// no documentation yet
	ResourceName *string `json:"resourceName,omitempty" xmlrpc:"resourceName"`

	// <<< EOT The unique identifier for the physical resource that is associated
	ResourceTableId *int `json:"resourceTableId,omitempty" xmlrpc:"resourceTableId"`
}

// This type contains general information related to a [[SoftLayer_Hardware]] resource that is impacted by a [[SoftLayer_Notification_Occurrence_Event]].
type Notification_Occurrence_Resource_Hardware struct {
	Notification_Occurrence_Resource

	// no documentation yet
	Hostname *string `json:"hostname,omitempty" xmlrpc:"hostname"`

	// no documentation yet
	PrivateIp *string `json:"privateIp,omitempty" xmlrpc:"privateIp"`

	// no documentation yet
	PublicIp *string `json:"publicIp,omitempty" xmlrpc:"publicIp"`

	// no documentation yet
	ResourceType *string `json:"resourceType,omitempty" xmlrpc:"resourceType"`
}

// This type contains general information related to a [[SoftLayer_Network_Application_Delivery_Controller]] resource that is impacted by a [[SoftLayer_Notification_Occurrence_Event]].
type Notification_Occurrence_Resource_Network_Application_Delivery_Controller struct {
	Notification_Occurrence_Resource

	// no documentation yet
	Hostname *string `json:"hostname,omitempty" xmlrpc:"hostname"`

	// no documentation yet
	PrivateIp *string `json:"privateIp,omitempty" xmlrpc:"privateIp"`

	// no documentation yet
	PublicIp *string `json:"publicIp,omitempty" xmlrpc:"publicIp"`

	// no documentation yet
	ResourceType *string `json:"resourceType,omitempty" xmlrpc:"resourceType"`
}

// This type contains general information related to a [[SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress]] resource that is impacted by a [[SoftLayer_Notification_Occurrence_Event]].
type Notification_Occurrence_Resource_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress struct {
	Notification_Occurrence_Resource

	// no documentation yet
	Hostname *string `json:"hostname,omitempty" xmlrpc:"hostname"`

	// no documentation yet
	PublicIp *string `json:"publicIp,omitempty" xmlrpc:"publicIp"`

	// no documentation yet
	ResourceType *string `json:"resourceType,omitempty" xmlrpc:"resourceType"`
}

// This type contains general information related to a [[SoftLayer_Network_Storage_Iscsi_EqualLogic]] resource that is impacted by a [[SoftLayer_Notification_Occurrence_Event]].
type Notification_Occurrence_Resource_Network_Storage_Iscsi_EqualLogic struct {
	Notification_Occurrence_Resource

	// no documentation yet
	Hostname *string `json:"hostname,omitempty" xmlrpc:"hostname"`

	// no documentation yet
	PrivateIp *string `json:"privateIp,omitempty" xmlrpc:"privateIp"`

	// no documentation yet
	ResourceType *string `json:"resourceType,omitempty" xmlrpc:"resourceType"`
}

// This type contains general information related to a [[SoftLayer_Network_Storage_Iscsi_NetApp]] resource that is impacted by a [[SoftLayer_Notification_Occurrence_Event]].
type Notification_Occurrence_Resource_Network_Storage_Iscsi_NetApp struct {
	Notification_Occurrence_Resource

	// no documentation yet
	Hostname *string `json:"hostname,omitempty" xmlrpc:"hostname"`

	// no documentation yet
	PrivateIp *string `json:"privateIp,omitempty" xmlrpc:"privateIp"`

	// no documentation yet
	ResourceType *string `json:"resourceType,omitempty" xmlrpc:"resourceType"`
}

// This type contains general information related to a [[SoftLayer_Network_Storage_Lockbox]] resource that is impacted by a [[SoftLayer_Notification_Occurrence_Event]].
type Notification_Occurrence_Resource_Network_Storage_Lockbox struct {
	Notification_Occurrence_Resource

	// no documentation yet
	Hostname *string `json:"hostname,omitempty" xmlrpc:"hostname"`

	// no documentation yet
	PrivateIp *string `json:"privateIp,omitempty" xmlrpc:"privateIp"`

	// no documentation yet
	ResourceType *string `json:"resourceType,omitempty" xmlrpc:"resourceType"`
}

// This type contains general information related to a [[SoftLayer_Network_Storage_Nas]] resource that is impacted by a [[SoftLayer_Notification_Occurrence_Event]].
type Notification_Occurrence_Resource_Network_Storage_Nas struct {
	Notification_Occurrence_Resource

	// no documentation yet
	Hostname *string `json:"hostname,omitempty" xmlrpc:"hostname"`

	// no documentation yet
	PrivateIp *string `json:"privateIp,omitempty" xmlrpc:"privateIp"`

	// no documentation yet
	ResourceType *string `json:"resourceType,omitempty" xmlrpc:"resourceType"`
}

// This type contains general information related to a [[SoftLayer_Network_Storage_NetApp_Volume]] resource that is impacted by a [[SoftLayer_Notification_Occurrence_Event]].
type Notification_Occurrence_Resource_Network_Storage_NetApp_Volume struct {
	Notification_Occurrence_Resource

	// no documentation yet
	Hostname *string `json:"hostname,omitempty" xmlrpc:"hostname"`

	// no documentation yet
	PrivateIp *string `json:"privateIp,omitempty" xmlrpc:"privateIp"`

	// no documentation yet
	ResourceType *string `json:"resourceType,omitempty" xmlrpc:"resourceType"`
}

// This type contains general information related to a [[SoftLayer_Virtual_Guest]] resource that is impacted by a [[SoftLayer_Notification_Occurrence_Event]].
type Notification_Occurrence_Resource_Virtual struct {
	Notification_Occurrence_Resource

	// no documentation yet
	Hostname *string `json:"hostname,omitempty" xmlrpc:"hostname"`

	// no documentation yet
	PrivateIp *string `json:"privateIp,omitempty" xmlrpc:"privateIp"`

	// no documentation yet
	PublicIp *string `json:"publicIp,omitempty" xmlrpc:"publicIp"`

	// no documentation yet
	ResourceType *string `json:"resourceType,omitempty" xmlrpc:"resourceType"`
}

// no documentation yet
type Notification_Occurrence_Status_Code struct {
	Entity

	// no documentation yet
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// no documentation yet
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// no documentation yet
type Notification_Occurrence_Update struct {
	Entity

	// no documentation yet
	Contents *string `json:"contents,omitempty" xmlrpc:"contents"`

	// no documentation yet
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// no documentation yet
	Employee *User_Employee `json:"employee,omitempty" xmlrpc:"employee"`

	// no documentation yet
	EndDate *Time `json:"endDate,omitempty" xmlrpc:"endDate"`

	// no documentation yet
	NotificationOccurrenceEvent *Notification_Occurrence_Event `json:"notificationOccurrenceEvent,omitempty" xmlrpc:"notificationOccurrenceEvent"`

	// no documentation yet
	StartDate *Time `json:"startDate,omitempty" xmlrpc:"startDate"`
}

// This type contains general information relating to a user that may be impacted by a [[SoftLayer_Notification_Occurrence_Event]].
type Notification_Occurrence_User struct {
	Entity

	// no documentation yet
	AcknowledgedFlag *int `json:"acknowledgedFlag,omitempty" xmlrpc:"acknowledgedFlag"`

	// no documentation yet
	Active *int `json:"active,omitempty" xmlrpc:"active"`

	// no documentation yet
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// A count of a collection of resources impacted by the associated event.
	ImpactedResourceCount *uint `json:"impactedResourceCount,omitempty" xmlrpc:"impactedResourceCount"`

	// A collection of resources impacted by the associated event.
	ImpactedResources []Notification_Occurrence_Resource `json:"impactedResources,omitempty" xmlrpc:"impactedResources"`

	// The associated event.
	NotificationOccurrenceEvent *Notification_Occurrence_Event `json:"notificationOccurrenceEvent,omitempty" xmlrpc:"notificationOccurrenceEvent"`

	// The impacted user.
	User *User_Customer `json:"user,omitempty" xmlrpc:"user"`

	// no documentation yet
	UsrRecordId *int `json:"usrRecordId,omitempty" xmlrpc:"usrRecordId"`
}

// Retrieve details for preferences.  Preferences are used to allow the subscriber to modify their subscription in various ways.  Details such as friendly name, keyname maximum and minimum values can be retrieved.  These provide details to help configure subscriber preferences correctly.
type Notification_Preference struct {
	Entity

	// A description of what the preference is used for.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// Unique identifier for the notification preference.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// Name that can be used by external systems to refer to preference.
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// Largest value allowed for the preference.
	MaximumValue *string `json:"maximumValue,omitempty" xmlrpc:"maximumValue"`

	// Smallest value allowed for the preference.
	MinimumValue *string `json:"minimumValue,omitempty" xmlrpc:"minimumValue"`

	// Friendly name for the notification.
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// The unit of measure used for the preference's value, minimum and maximum as well.
	Units *string `json:"units,omitempty" xmlrpc:"units"`

	// Default value used when setting up preferences for a new subscriber.
	Value *string `json:"value,omitempty" xmlrpc:"value"`
}

// no documentation yet
type Notification_Subscriber struct {
	Entity

	// no documentation yet
	Active *int `json:"active,omitempty" xmlrpc:"active"`

	// no documentation yet
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// A count of
	DeliveryMethodCount *uint `json:"deliveryMethodCount,omitempty" xmlrpc:"deliveryMethodCount"`

	// no documentation yet
	DeliveryMethods []Notification_Subscriber_Delivery_Method `json:"deliveryMethods,omitempty" xmlrpc:"deliveryMethods"`

	// no documentation yet
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// no documentation yet
	Notification *Notification `json:"notification,omitempty" xmlrpc:"notification"`

	// no documentation yet
	NotificationId *int `json:"notificationId,omitempty" xmlrpc:"notificationId"`

	// no documentation yet
	NotificationSubscriberTypeId *int `json:"notificationSubscriberTypeId,omitempty" xmlrpc:"notificationSubscriberTypeId"`

	// no documentation yet
	NotificationSubscriberTypeResourceId *int `json:"notificationSubscriberTypeResourceId,omitempty" xmlrpc:"notificationSubscriberTypeResourceId"`
}

// no documentation yet
type Notification_Subscriber_Customer struct {
	Notification_Subscriber

	// no documentation yet
	SubscriberRecord *User_Customer `json:"subscriberRecord,omitempty" xmlrpc:"subscriberRecord"`
}

// Provides details for the subscriber's delivery methods.
type Notification_Subscriber_Delivery_Method struct {
	Entity

	// Indicates the subscriber's delivery method availability for notifications.
	Active *int `json:"active,omitempty" xmlrpc:"active"`

	// Date the subscriber's delivery method was created.
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// Date the subscriber's delivery method was last modified.
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// no documentation yet
	NotificationDeliveryMethod *Notification_Delivery_Method `json:"notificationDeliveryMethod,omitempty" xmlrpc:"notificationDeliveryMethod"`

	// Identifier for the notification delivery method.
	NotificationDeliveryMethodId *int `json:"notificationDeliveryMethodId,omitempty" xmlrpc:"notificationDeliveryMethodId"`

	// no documentation yet
	NotificationSubscriber *Notification_Subscriber `json:"notificationSubscriber,omitempty" xmlrpc:"notificationSubscriber"`

	// Identifier for the subscriber.
	NotificationSubscriberId *int `json:"notificationSubscriberId,omitempty" xmlrpc:"notificationSubscriberId"`
}

// A notification subscriber will have details pertaining to the subscriber's notification subscription.  You can receive details such as preferences, details of the preferences, delivery methods and the delivery methods for the subscriber.
//
// NOTE: There are preferences and delivery methods that cannot be modified.  Also, there are some subscriptions that are required.
type Notification_User_Subscriber struct {
	Entity

	// The current status of the subscription.
	Active *int `json:"active,omitempty" xmlrpc:"active"`

	// A count of the delivery methods used to send the subscribed notification.
	DeliveryMethodCount *uint `json:"deliveryMethodCount,omitempty" xmlrpc:"deliveryMethodCount"`

	// The delivery methods used to send the subscribed notification.
	DeliveryMethods []Notification_Delivery_Method `json:"deliveryMethods,omitempty" xmlrpc:"deliveryMethods"`

	// Unique identifier of the subscriber that will receive the alerts.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// Notification subscribed to.
	Notification *Notification `json:"notification,omitempty" xmlrpc:"notification"`

	// Unique identifier of the notification subscribed to.
	NotificationId *int `json:"notificationId,omitempty" xmlrpc:"notificationId"`

	// A count of associated subscriber preferences used for the notification subscription. For example, preferences include number of deliveries (limit) and threshold.
	PreferenceCount *uint `json:"preferenceCount,omitempty" xmlrpc:"preferenceCount"`

	// Associated subscriber preferences used for the notification subscription. For example, preferences include number of deliveries (limit) and threshold.
	Preferences []Notification_User_Subscriber_Preference `json:"preferences,omitempty" xmlrpc:"preferences"`

	// A count of preference details such as description, minimum and maximum limits, default value and unit of measure.
	PreferencesDetailCount *uint `json:"preferencesDetailCount,omitempty" xmlrpc:"preferencesDetailCount"`

	// Preference details such as description, minimum and maximum limits, default value and unit of measure.
	PreferencesDetails []Notification_Preference `json:"preferencesDetails,omitempty" xmlrpc:"preferencesDetails"`

	// The subscriber id to resource id mapping.
	ResourceRecord *Notification_User_Subscriber_Resource `json:"resourceRecord,omitempty" xmlrpc:"resourceRecord"`

	// User record for the subscription.
	UserRecord *User_Customer `json:"userRecord,omitempty" xmlrpc:"userRecord"`

	// Unique identifier of the user the subscription is for.
	UserRecordId *int `json:"userRecordId,omitempty" xmlrpc:"userRecordId"`
}

// A notification subscriber will have details pertaining to the subscriber's notification subscription.  You can receive details such as preferences, details of the preferences, delivery methods and the delivery methods for the subscriber.
//
// NOTE: There are preferences and delivery methods that cannot be modified.  Also, there are some subscriptions that are required.
type Notification_User_Subscriber_Billing struct {
	Notification_User_Subscriber
}

// Provides mapping details of how the subscriber's notification will be delivered.  This maps the subscriber's id with all the delivery method ids used to delivery the notification.
type Notification_User_Subscriber_Delivery_Method struct {
	Entity

	// Determines if the delivery method is active for the user.
	Active *int `json:"active,omitempty" xmlrpc:"active"`

	// Provides details for the method used to deliver the notification (email, sms, ticket).
	DeliveryMethod *Notification_Delivery_Method `json:"deliveryMethod,omitempty" xmlrpc:"deliveryMethod"`

	// Unique identifier of the method used to deliver notification.
	NotificationMethodId *int `json:"notificationMethodId,omitempty" xmlrpc:"notificationMethodId"`

	// The Subscriber information tied to the delivery method.
	NotificationUserSubscriber *Notification_User_Subscriber `json:"notificationUserSubscriber,omitempty" xmlrpc:"notificationUserSubscriber"`

	// Unique identifier of the subscriber tied to the delivery method.
	NotificationUserSubscriberId *int `json:"notificationUserSubscriberId,omitempty" xmlrpc:"notificationUserSubscriberId"`
}

// A notification subscriber will have details pertaining to the subscriber's notification subscription.  You can receive details such as preferences, details of the preferences, delivery methods and the delivery methods for the subscriber.
//
// NOTE: There are preferences and delivery methods that cannot be modified.  Also, there are some subscriptions that are required.
type Notification_User_Subscriber_Mobile struct {
	Notification_User_Subscriber
}

// Preferences are settings that can be modified to change the behavior of the subscription.  For example, modify the limit preference to only receive notifications 10 times instead of 1 during a billing cycle.
//
// NOTE: Some preferences have certain restrictions on values that can be set.
type Notification_User_Subscriber_Preference struct {
	Entity

	// Details such name, keyname, minimum and maximum values for the preference.
	DefaultPreference *Notification_Preference `json:"defaultPreference,omitempty" xmlrpc:"defaultPreference"`

	// Unique identifier for the subscriber's preferences.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// Unique identifier of the default preference for which the subscriber preference is based on.  For example, if no preferences are supplied during the creation of a subscriber.  The default values are pulled using this property.
	NotificationPreferenceId *int `json:"notificationPreferenceId,omitempty" xmlrpc:"notificationPreferenceId"`

	// Details of the subscriber tied to the preference.
	NotificationUserSubscriber *Notification_User_Subscriber `json:"notificationUserSubscriber,omitempty" xmlrpc:"notificationUserSubscriber"`

	// Unique identifier of the subscriber tied to the subscriber preference.
	NotificationUserSubscriberId *int `json:"notificationUserSubscriberId,omitempty" xmlrpc:"notificationUserSubscriberId"`

	// The user supplied value to "override" the "default" preference's value.
	Value *string `json:"value,omitempty" xmlrpc:"value"`
}

// Retrieve identifier cross-reference information.  SoftLayer_Notification_User_Subscriber_Resource provides the resource table id and subscriber id relation. The resource table id is the id of the service the subscriber receives alerts for.  This resource table id could be the unique identifier for a Storage Evault service, Global Load Balancer or CDN service.
type Notification_User_Subscriber_Resource struct {
	Entity

	// The Subscriber information tied to the resource service.
	NotificationUserSubscriber *Notification_User_Subscriber `json:"notificationUserSubscriber,omitempty" xmlrpc:"notificationUserSubscriber"`

	// Unique identifier of the subscriber that will receive the alerts for the resource subscribed to a notification.
	NotificationUserSubscriberId *int `json:"notificationUserSubscriberId,omitempty" xmlrpc:"notificationUserSubscriberId"`

	// Unique identifier for a SoftLayer service that is subscribed to a notification.  Currently, the SoftLayer services that can be subscribed to notifications are:
	//
	// Storage EVault CDN Global Load Balancer
	//
	//
	ResourceTableId *int `json:"resourceTableId,omitempty" xmlrpc:"resourceTableId"`
}
