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

package datatypes

type Notification struct {
	Entity

	Id                      *int                      `json:"id,omitempty"`
	KeyName                 *string                   `json:"keyName,omitempty"`
	Name                    *string                   `json:"name,omitempty"`
	PreferenceCount         *uint                     `json:"preferenceCount,omitempty"`
	Preferences             []Notification_Preference `json:"preferences,omitempty"`
	RequiredPreferenceCount *uint                     `json:"requiredPreferenceCount,omitempty"`
	RequiredPreferences     []Notification_Preference `json:"requiredPreferences,omitempty"`
}

type Notification_Delivery_Method struct {
	Entity

	Active      *int    `json:"active,omitempty"`
	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	KeyName     *string `json:"keyName,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Notification_Mobile struct {
	Notification
}

type Notification_Occurrence_Account struct {
	Entity

	Account                     *Account                        `json:"account,omitempty"`
	Active                      *int                            `json:"active,omitempty"`
	LastNotificationUpdate      *Notification_Occurrence_Update `json:"lastNotificationUpdate,omitempty"`
	NotificationOccurrenceEvent *Notification_Occurrence_Event  `json:"notificationOccurrenceEvent,omitempty"`
}

type Notification_Occurrence_Event struct {
	Entity

	AcknowledgedFlag                *bool                                      `json:"acknowledgedFlag,omitempty"`
	AttachmentCount                 *uint                                      `json:"attachmentCount,omitempty"`
	Attachments                     []Notification_Occurrence_Event_Attachment `json:"attachments,omitempty"`
	EndDate                         *Time                                      `json:"endDate,omitempty"`
	FirstUpdate                     *Notification_Occurrence_Update            `json:"firstUpdate,omitempty"`
	Id                              *int                                       `json:"id,omitempty"`
	ImpactedAccountCount            *uint                                      `json:"impactedAccountCount,omitempty"`
	ImpactedAccounts                []Notification_Occurrence_Account          `json:"impactedAccounts,omitempty"`
	ImpactedResourceCount           *uint                                      `json:"impactedResourceCount,omitempty"`
	ImpactedResources               []Notification_Occurrence_Resource         `json:"impactedResources,omitempty"`
	ImpactedUserCount               *uint                                      `json:"impactedUserCount,omitempty"`
	ImpactedUsers                   []Notification_Occurrence_User             `json:"impactedUsers,omitempty"`
	LastImpactedUserCount           *int                                       `json:"lastImpactedUserCount,omitempty"`
	LastUpdate                      *Notification_Occurrence_Update            `json:"lastUpdate,omitempty"`
	ModifyDate                      *Time                                      `json:"modifyDate,omitempty"`
	NotificationOccurrenceEventType *Notification_Occurrence_Event_Type        `json:"notificationOccurrenceEventType,omitempty"`
	RecoveryTime                    *int                                       `json:"recoveryTime,omitempty"`
	StartDate                       *Time                                      `json:"startDate,omitempty"`
	StatusCode                      *Notification_Occurrence_Status_Code       `json:"statusCode,omitempty"`
	Subject                         *string                                    `json:"subject,omitempty"`
	Summary                         *string                                    `json:"summary,omitempty"`
	SystemTicketId                  *int                                       `json:"systemTicketId,omitempty"`
	UpdateCount                     *uint                                      `json:"updateCount,omitempty"`
	Updates                         []Notification_Occurrence_Update           `json:"updates,omitempty"`
}

type Notification_Occurrence_Event_Attachment struct {
	Entity

	CreateDate                    *Time                          `json:"createDate,omitempty"`
	FileName                      *string                        `json:"fileName,omitempty"`
	FileSize                      *string                        `json:"fileSize,omitempty"`
	Id                            *int                           `json:"id,omitempty"`
	NotificationOccurrenceEvent   *Notification_Occurrence_Event `json:"notificationOccurrenceEvent,omitempty"`
	NotificationOccurrenceEventId *int                           `json:"notificationOccurrenceEventId,omitempty"`
}

type Notification_Occurrence_Event_Type struct {
	Entity

	KeyName *string `json:"keyName,omitempty"`
}

type Notification_Occurrence_Resource struct {
	Entity

	Active                        *int                           `json:"active,omitempty"`
	FilterLabel                   *string                        `json:"filterLabel,omitempty"`
	NotificationOccurrenceEvent   *Notification_Occurrence_Event `json:"notificationOccurrenceEvent,omitempty"`
	NotificationOccurrenceEventId *int                           `json:"notificationOccurrenceEventId,omitempty"`
	Resource                      *Entity                        `json:"resource,omitempty"`
	ResourceAccountId             *int                           `json:"resourceAccountId,omitempty"`
	ResourceName                  *string                        `json:"resourceName,omitempty"`
	ResourceTableId               *int                           `json:"resourceTableId,omitempty"`
}

type Notification_Occurrence_Resource_Hardware struct {
	Notification_Occurrence_Resource

	Hostname     *string `json:"hostname,omitempty"`
	PrivateIp    *string `json:"privateIp,omitempty"`
	PublicIp     *string `json:"publicIp,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
}

type Notification_Occurrence_Resource_Network_Application_Delivery_Controller struct {
	Notification_Occurrence_Resource

	Hostname     *string `json:"hostname,omitempty"`
	PrivateIp    *string `json:"privateIp,omitempty"`
	PublicIp     *string `json:"publicIp,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
}

type Notification_Occurrence_Resource_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress struct {
	Notification_Occurrence_Resource

	Hostname     *string `json:"hostname,omitempty"`
	PublicIp     *string `json:"publicIp,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
}

type Notification_Occurrence_Resource_Network_Storage_Iscsi_EqualLogic struct {
	Notification_Occurrence_Resource

	Hostname     *string `json:"hostname,omitempty"`
	PrivateIp    *string `json:"privateIp,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
}

type Notification_Occurrence_Resource_Network_Storage_Iscsi_NetApp struct {
	Notification_Occurrence_Resource

	Hostname     *string `json:"hostname,omitempty"`
	PrivateIp    *string `json:"privateIp,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
}

type Notification_Occurrence_Resource_Network_Storage_Lockbox struct {
	Notification_Occurrence_Resource

	Hostname     *string `json:"hostname,omitempty"`
	PrivateIp    *string `json:"privateIp,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
}

type Notification_Occurrence_Resource_Network_Storage_Nas struct {
	Notification_Occurrence_Resource

	Hostname     *string `json:"hostname,omitempty"`
	PrivateIp    *string `json:"privateIp,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
}

type Notification_Occurrence_Resource_Network_Storage_NetApp_Volume struct {
	Notification_Occurrence_Resource

	Hostname     *string `json:"hostname,omitempty"`
	PrivateIp    *string `json:"privateIp,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
}

type Notification_Occurrence_Resource_Virtual struct {
	Notification_Occurrence_Resource

	Hostname     *string `json:"hostname,omitempty"`
	PrivateIp    *string `json:"privateIp,omitempty"`
	PublicIp     *string `json:"publicIp,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
}

type Notification_Occurrence_Status_Code struct {
	Entity

	Description *string `json:"description,omitempty"`
	KeyName     *string `json:"keyName,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Notification_Occurrence_Update struct {
	Entity

	Contents                    *string                        `json:"contents,omitempty"`
	CreateDate                  *Time                          `json:"createDate,omitempty"`
	Employee                    *User_Employee                 `json:"employee,omitempty"`
	EndDate                     *Time                          `json:"endDate,omitempty"`
	NotificationOccurrenceEvent *Notification_Occurrence_Event `json:"notificationOccurrenceEvent,omitempty"`
	StartDate                   *Time                          `json:"startDate,omitempty"`
}

type Notification_Occurrence_User struct {
	Entity

	AcknowledgedFlag            *int                               `json:"acknowledgedFlag,omitempty"`
	Active                      *int                               `json:"active,omitempty"`
	Id                          *int                               `json:"id,omitempty"`
	ImpactedResourceCount       *uint                              `json:"impactedResourceCount,omitempty"`
	ImpactedResources           []Notification_Occurrence_Resource `json:"impactedResources,omitempty"`
	NotificationOccurrenceEvent *Notification_Occurrence_Event     `json:"notificationOccurrenceEvent,omitempty"`
	User                        *User_Customer                     `json:"user,omitempty"`
	UsrRecordId                 *int                               `json:"usrRecordId,omitempty"`
}

type Notification_Preference struct {
	Entity

	Description  *string `json:"description,omitempty"`
	Id           *int    `json:"id,omitempty"`
	KeyName      *string `json:"keyName,omitempty"`
	MaximumValue *string `json:"maximumValue,omitempty"`
	MinimumValue *string `json:"minimumValue,omitempty"`
	Name         *string `json:"name,omitempty"`
	Units        *string `json:"units,omitempty"`
	Value        *string `json:"value,omitempty"`
}

type Notification_Subscriber struct {
	Entity

	Active                               *int                                      `json:"active,omitempty"`
	CreateDate                           *Time                                     `json:"createDate,omitempty"`
	DeliveryMethodCount                  *uint                                     `json:"deliveryMethodCount,omitempty"`
	DeliveryMethods                      []Notification_Subscriber_Delivery_Method `json:"deliveryMethods,omitempty"`
	Id                                   *int                                      `json:"id,omitempty"`
	ModifyDate                           *Time                                     `json:"modifyDate,omitempty"`
	Notification                         *Notification                             `json:"notification,omitempty"`
	NotificationId                       *int                                      `json:"notificationId,omitempty"`
	NotificationSubscriberTypeId         *int                                      `json:"notificationSubscriberTypeId,omitempty"`
	NotificationSubscriberTypeResourceId *int                                      `json:"notificationSubscriberTypeResourceId,omitempty"`
}

type Notification_Subscriber_Customer struct {
	Notification_Subscriber

	SubscriberRecord *User_Customer `json:"subscriberRecord,omitempty"`
}

type Notification_Subscriber_Delivery_Method struct {
	Entity

	Active                       *int                          `json:"active,omitempty"`
	CreateDate                   *Time                         `json:"createDate,omitempty"`
	ModifyDate                   *Time                         `json:"modifyDate,omitempty"`
	NotificationDeliveryMethod   *Notification_Delivery_Method `json:"notificationDeliveryMethod,omitempty"`
	NotificationDeliveryMethodId *int                          `json:"notificationDeliveryMethodId,omitempty"`
	NotificationSubscriber       *Notification_Subscriber      `json:"notificationSubscriber,omitempty"`
	NotificationSubscriberId     *int                          `json:"notificationSubscriberId,omitempty"`
}

type Notification_User_Subscriber struct {
	Entity

	Active                 *int                                      `json:"active,omitempty"`
	DeliveryMethodCount    *uint                                     `json:"deliveryMethodCount,omitempty"`
	DeliveryMethods        []Notification_Delivery_Method            `json:"deliveryMethods,omitempty"`
	Id                     *int                                      `json:"id,omitempty"`
	Notification           *Notification                             `json:"notification,omitempty"`
	NotificationId         *int                                      `json:"notificationId,omitempty"`
	PreferenceCount        *uint                                     `json:"preferenceCount,omitempty"`
	Preferences            []Notification_User_Subscriber_Preference `json:"preferences,omitempty"`
	PreferencesDetailCount *uint                                     `json:"preferencesDetailCount,omitempty"`
	PreferencesDetails     []Notification_Preference                 `json:"preferencesDetails,omitempty"`
	ResourceRecord         *Notification_User_Subscriber_Resource    `json:"resourceRecord,omitempty"`
	UserRecord             *User_Customer                            `json:"userRecord,omitempty"`
	UserRecordId           *int                                      `json:"userRecordId,omitempty"`
}

type Notification_User_Subscriber_Billing struct {
	Notification_User_Subscriber
}

type Notification_User_Subscriber_Delivery_Method struct {
	Entity

	Active                       *int                          `json:"active,omitempty"`
	DeliveryMethod               *Notification_Delivery_Method `json:"deliveryMethod,omitempty"`
	NotificationMethodId         *int                          `json:"notificationMethodId,omitempty"`
	NotificationUserSubscriber   *Notification_User_Subscriber `json:"notificationUserSubscriber,omitempty"`
	NotificationUserSubscriberId *int                          `json:"notificationUserSubscriberId,omitempty"`
}

type Notification_User_Subscriber_Mobile struct {
	Notification_User_Subscriber
}

type Notification_User_Subscriber_Preference struct {
	Entity

	DefaultPreference            *Notification_Preference      `json:"defaultPreference,omitempty"`
	Id                           *int                          `json:"id,omitempty"`
	NotificationPreferenceId     *int                          `json:"notificationPreferenceId,omitempty"`
	NotificationUserSubscriber   *Notification_User_Subscriber `json:"notificationUserSubscriber,omitempty"`
	NotificationUserSubscriberId *int                          `json:"notificationUserSubscriberId,omitempty"`
	Value                        *string                       `json:"value,omitempty"`
}

type Notification_User_Subscriber_Resource struct {
	Entity

	NotificationUserSubscriber   *Notification_User_Subscriber `json:"notificationUserSubscriber,omitempty"`
	NotificationUserSubscriberId *int                          `json:"notificationUserSubscriberId,omitempty"`
	ResourceTableId              *int                          `json:"resourceTableId,omitempty"`
}
