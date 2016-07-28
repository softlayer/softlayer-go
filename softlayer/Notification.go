package softlayer

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Notification interface {
	Entity

	GetAllObjects() (datatypes.Notification, error)
	GetObject() (datatypes.Notification, error)

	GetPreferences() (datatypes.Notification_Preference, error)
	GetRequiredPreferences() (datatypes.Notification_Preference, error)
}

type Notification_Delivery_Method interface {
	Entity
}

type Notification_Mobile interface {
	Notification

	CreateSubscriberForMobileDevice(keyName string, resourceTableId int, userRecordId int) (bool, error)
	GetObject() (datatypes.Notification_Mobile, error)
}

type Notification_Occurrence_Account interface {
	Entity

	GetAccount() (datatypes.Account, error)
	GetLastNotificationUpdate() (datatypes.Notification_Occurrence_Update, error)
	GetNotificationOccurrenceEvent() (datatypes.Notification_Occurrence_Event, error)
}

type Notification_Occurrence_Event interface {
	Entity

	AcknowledgeNotification() (bool, error)
	GetAllObjects() (datatypes.Notification_Occurrence_Event, error)
	GetAttachedFile(attachmentId int) ([]byte, error)
	GetImpactedAccountCount() (int, error)
	GetImpactedDeviceCount() (int, error)
	GetImpactedDevices() (datatypes.Notification_Occurrence_Resource, error)
	GetObject() (datatypes.Notification_Occurrence_Event, error)

	GetAcknowledgedFlag() (bool, error)
	GetAttachments() (datatypes.Notification_Occurrence_Event_Attachment, error)
	GetFirstUpdate() (datatypes.Notification_Occurrence_Update, error)
	GetImpactedAccounts() (datatypes.Notification_Occurrence_Account, error)
	GetImpactedResources() (datatypes.Notification_Occurrence_Resource, error)
	GetImpactedUsers() (datatypes.Notification_Occurrence_User, error)
	GetLastUpdate() (datatypes.Notification_Occurrence_Update, error)
	GetNotificationOccurrenceEventType() (datatypes.Notification_Occurrence_Event_Type, error)
	GetStatusCode() (datatypes.Notification_Occurrence_Status_Code, error)
	GetUpdates() (datatypes.Notification_Occurrence_Update, error)
}

type Notification_Occurrence_Event_Attachment interface {
	Entity

	GetNotificationOccurrenceEvent() (datatypes.Notification_Occurrence_Event, error)
}

type Notification_Occurrence_Event_Type interface {
	Entity
}

type Notification_Occurrence_Resource interface {
	Entity

	GetNotificationOccurrenceEvent() (datatypes.Notification_Occurrence_Event, error)
	GetResource() (datatypes.Entity, error)
}

type Notification_Occurrence_Resource_Hardware interface {
	Notification_Occurrence_Resource
}

type Notification_Occurrence_Resource_Network_Application_Delivery_Controller interface {
	Notification_Occurrence_Resource
}

type Notification_Occurrence_Resource_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress interface {
	Notification_Occurrence_Resource
}

type Notification_Occurrence_Resource_Network_Storage_Iscsi_EqualLogic interface {
	Notification_Occurrence_Resource
}

type Notification_Occurrence_Resource_Network_Storage_Iscsi_NetApp interface {
	Notification_Occurrence_Resource
}

type Notification_Occurrence_Resource_Network_Storage_Lockbox interface {
	Notification_Occurrence_Resource
}

type Notification_Occurrence_Resource_Network_Storage_Nas interface {
	Notification_Occurrence_Resource
}

type Notification_Occurrence_Resource_Network_Storage_NetApp_Volume interface {
	Notification_Occurrence_Resource
}

type Notification_Occurrence_Resource_Virtual interface {
	Notification_Occurrence_Resource
}

type Notification_Occurrence_Status_Code interface {
	Entity
}

type Notification_Occurrence_Update interface {
	Entity

	GetEmployee() (datatypes.User_Employee, error)
	GetNotificationOccurrenceEvent() (datatypes.Notification_Occurrence_Event, error)
}

type Notification_Occurrence_User interface {
	Entity

	Acknowledge() (bool, error)
	GetAllObjects() (datatypes.Notification_Occurrence_User, error)
	GetImpactedDeviceCount() (int, error)
	GetObject() (datatypes.Notification_Occurrence_User, error)

	GetImpactedResources() (datatypes.Notification_Occurrence_Resource, error)
	GetNotificationOccurrenceEvent() (datatypes.Notification_Occurrence_Event, error)
	GetUser() (datatypes.User_Customer, error)
}

type Notification_Preference interface {
	Entity
}

type Notification_Subscriber interface {
	Entity

	GetDeliveryMethods() (datatypes.Notification_Subscriber_Delivery_Method, error)
	GetNotification() (datatypes.Notification, error)
}

type Notification_Subscriber_Customer interface {
	Notification_Subscriber

	GetSubscriberRecord() (datatypes.User_Customer, error)
}

type Notification_Subscriber_Delivery_Method interface {
	Entity

	GetNotificationDeliveryMethod() (datatypes.Notification_Delivery_Method, error)
	GetNotificationSubscriber() (datatypes.Notification_Subscriber, error)
}

type Notification_User_Subscriber interface {
	Entity

	CreateObject(templateObject datatypes.Notification_User_Subscriber) (bool, error)
	EditObject(templateObject datatypes.Notification_User_Subscriber) (bool, error)
	GetObject() (datatypes.Notification_User_Subscriber, error)

	GetDeliveryMethods() (datatypes.Notification_Delivery_Method, error)
	GetNotification() (datatypes.Notification, error)
	GetPreferences() (datatypes.Notification_User_Subscriber_Preference, error)
	GetPreferencesDetails() (datatypes.Notification_Preference, error)
	GetResourceRecord() (datatypes.Notification_User_Subscriber_Resource, error)
	GetUserRecord() (datatypes.User_Customer, error)
}

type Notification_User_Subscriber_Billing interface {
	Notification_User_Subscriber

	GetObject() (datatypes.Notification_User_Subscriber_Billing, error)
}

type Notification_User_Subscriber_Delivery_Method interface {
	Entity

	GetDeliveryMethod() (datatypes.Notification_Delivery_Method, error)
	GetNotificationUserSubscriber() (datatypes.Notification_User_Subscriber, error)
}

type Notification_User_Subscriber_Mobile interface {
	Notification_User_Subscriber

	ClearSnoozeTimer() (bool, error)
	GetObject() (datatypes.Notification_User_Subscriber_Mobile, error)
	SetSnoozeTimer(start int, end int) (bool, error)
}

type Notification_User_Subscriber_Preference interface {
	Entity

	CreateObject(templateObject datatypes.Notification_User_Subscriber_Preference) (bool, error)
	EditObjects(templateObjects datatypes.Notification_User_Subscriber_Preference) (bool, error)
	GetObject() (datatypes.Notification_User_Subscriber_Preference, error)

	GetDefaultPreference() (datatypes.Notification_Preference, error)
	GetNotificationUserSubscriber() (datatypes.Notification_User_Subscriber, error)
}

type Notification_User_Subscriber_Resource interface {
	Entity

	GetNotificationUserSubscriber() (datatypes.Notification_User_Subscriber, error)
}
