package softlayer

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Auxiliary_Marketing_Event interface {
	Entity

	GetMarketingEvents() (datatypes.Auxiliary_Marketing_Event, error)
	GetObject() (datatypes.Auxiliary_Marketing_Event, error)
}

type Auxiliary_Network_Status interface {
	Entity

	GetNetworkStatus(target string) (datatypes.Container_Auxiliary_Network_Status_Reading, error)
}

type Auxiliary_Notification_Emergency interface {
	Entity

	GetAllObjects() (datatypes.Auxiliary_Notification_Emergency, error)
	GetCurrentNotifications() (datatypes.Auxiliary_Notification_Emergency, error)
	GetObject() (datatypes.Auxiliary_Notification_Emergency, error)

	GetSignature() (datatypes.Auxiliary_Notification_Emergency_Signature, error)
	GetStatus() (datatypes.Auxiliary_Notification_Emergency_Status, error)
}

type Auxiliary_Notification_Emergency_Signature interface {
	Entity
}

type Auxiliary_Notification_Emergency_Status interface {
	Entity
}

type Auxiliary_Press_Release interface {
	Entity

	GetAllObjects() (datatypes.Auxiliary_Press_Release, error)
	GetObject() (datatypes.Auxiliary_Press_Release, error)
	GetRenderedPressRelease() (datatypes.Auxiliary_Press_Release, error)
	GetRenderedPressReleases(resultLimit string, year string) (datatypes.Auxiliary_Press_Release, error)
	GetWebsiteHighlightPressReleases() (datatypes.Auxiliary_Press_Release, error)

	GetAbout() (datatypes.Auxiliary_Press_Release_About_Press_Release, error)
	GetContacts() (datatypes.Auxiliary_Press_Release_Contact_Press_Release, error)
	GetMediaPartners() (datatypes.Auxiliary_Press_Release_Media_Partner_Press_Release, error)
	GetPressReleaseContent() (datatypes.Auxiliary_Press_Release_Content, error)
}

type Auxiliary_Press_Release_About interface {
	Entity

	GetObject() (datatypes.Auxiliary_Press_Release_About, error)
}

type Auxiliary_Press_Release_About_Press_Release interface {
	Entity

	GetObject() (datatypes.Auxiliary_Press_Release_About_Press_Release, error)

	GetAboutParagraphs() (datatypes.Auxiliary_Press_Release_About, error)
	GetPressReleases() (datatypes.Auxiliary_Press_Release, error)
}

type Auxiliary_Press_Release_Contact interface {
	Entity

	GetObject() (datatypes.Auxiliary_Press_Release_Contact, error)
}

type Auxiliary_Press_Release_Contact_Press_Release interface {
	Entity

	GetObject() (datatypes.Auxiliary_Press_Release_Contact_Press_Release, error)

	GetContacts() (datatypes.Auxiliary_Press_Release_Contact, error)
	GetPressReleases() (datatypes.Auxiliary_Press_Release, error)
}

type Auxiliary_Press_Release_Content interface {
	Entity

	GetObject() (datatypes.Auxiliary_Press_Release_Content, error)
}

type Auxiliary_Press_Release_Media_Partner interface {
	Entity

	GetObject() (datatypes.Auxiliary_Press_Release_Media_Partner, error)
}

type Auxiliary_Press_Release_Media_Partner_Press_Release interface {
	Entity

	GetObject() (datatypes.Auxiliary_Press_Release_Media_Partner_Press_Release, error)

	GetMediaPartners() (datatypes.Auxiliary_Press_Release_Media_Partner, error)
	GetPressReleases() (datatypes.Auxiliary_Press_Release, error)
}

type Auxiliary_Shipping_Courier interface {
	Entity
}

type Auxiliary_Shipping_Courier_Type interface {
	Entity

	GetObject() (datatypes.Auxiliary_Shipping_Courier_Type, error)
	GetTypeByKeyName(keyName string) (datatypes.Auxiliary_Shipping_Courier_Type, error)

	GetCourier() (datatypes.Auxiliary_Shipping_Courier, error)
}
