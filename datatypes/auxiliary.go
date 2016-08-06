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

//
type Auxiliary_Marketing_Event struct {
	Entity

	//
	CreateDate *Time `json:"createDate,omitempty"`

	//
	EnabledFlag *int `json:"enabledFlag,omitempty"`

	//
	EndDate *Time `json:"endDate,omitempty"`

	//
	Location *string `json:"location,omitempty"`

	//
	ModifyDate *Time `json:"modifyDate,omitempty"`

	//
	StartDate *Time `json:"startDate,omitempty"`

	//
	Title *string `json:"title,omitempty"`

	//
	Url *string `json:"url,omitempty"`
}

//
type Auxiliary_Network_Status struct {
	Entity
}

// A SoftLayer_Auxiliary_Notification_Emergency data object represents a notification event being broadcast to the SoftLayer customer base. It is used to provide information regarding outages or current known issues.
type Auxiliary_Notification_Emergency struct {
	Entity

	// The date this event was created.
	CreateDate *Time `json:"createDate,omitempty"`

	// The device (if any) effected by this event.
	Device *string `json:"device,omitempty"`

	// The duration of this event.
	Duration *string `json:"duration,omitempty"`

	// The device (if any) effected by this event.
	Id *int `json:"id,omitempty"`

	// The location effected by this event.
	Location *string `json:"location,omitempty"`

	// A message describing this event.
	Message *string `json:"message,omitempty"`

	// The last date this event was modified.
	ModifyDate *Time `json:"modifyDate,omitempty"`

	// The service(s) (if any) effected by this event.
	ServicesAffected *string `json:"servicesAffected,omitempty"`

	// The signature of the SoftLayer employee department associated with this notification.
	Signature *Auxiliary_Notification_Emergency_Signature `json:"signature,omitempty"`

	// The date this event will start.
	StartDate *Time `json:"startDate,omitempty"`

	// The status of this notification.
	Status *Auxiliary_Notification_Emergency_Status `json:"status,omitempty"`

	// Current status record for this event.
	StatusId *int `json:"statusId,omitempty"`
}

// Every SoftLayer_Auxiliary_Notification_Emergency has a signatureId that references a SoftLayer_Auxiliary_Notification_Emergency_Signature data type.  The signature is the user or group  responsible for the current event.
type Auxiliary_Notification_Emergency_Signature struct {
	Entity

	// The name or signature for the current Emergency Notification.
	Name *string `json:"name,omitempty"`
}

// Every SoftLayer_Auxiliary_Notification_Emergency has a statusId that references a SoftLayer_Auxiliary_Notification_Emergency_Status data type.  The status is used to determine the current state of the event.
type Auxiliary_Notification_Emergency_Status struct {
	Entity

	// A name describing the status of the current Emergency Notification.
	Name *string `json:"name,omitempty"`
}

//
type Auxiliary_Press_Release struct {
	Entity

	//
	About []Auxiliary_Press_Release_About_Press_Release `json:"about,omitempty"`

	// A count of
	AboutCount *uint `json:"aboutCount,omitempty"`

	// A count of
	ContactCount *uint `json:"contactCount,omitempty"`

	//
	Contacts []Auxiliary_Press_Release_Contact_Press_Release `json:"contacts,omitempty"`

	// A press release's internal identifier.
	Id *int `json:"id,omitempty"`

	// A count of
	MediaPartnerCount *uint `json:"mediaPartnerCount,omitempty"`

	//
	MediaPartners []Auxiliary_Press_Release_Media_Partner_Press_Release `json:"mediaPartners,omitempty"`

	//
	PressReleaseContent *Auxiliary_Press_Release_Content `json:"pressReleaseContent,omitempty"`

	// The data a press release was published.
	PublishDate *Time `json:"publishDate,omitempty"`

	// A press release's location.
	ReleaseLocation *string `json:"releaseLocation,omitempty"`

	// A press release's sub-title.
	SubTitle *string `json:"subTitle,omitempty"`

	// A press release's title.
	Title *string `json:"title,omitempty"`

	// Whether or not a press release is highlighted on the SoftLayer Website.
	WebsiteHighlightFlag *bool `json:"websiteHighlightFlag,omitempty"`
}

//
type Auxiliary_Press_Release_About struct {
	Entity

	// A press release about's content.
	Content *string `json:"content,omitempty"`

	// A press release about's internal
	Id *int `json:"id,omitempty"`

	// A press release about's title.
	Title *string `json:"title,omitempty"`
}

//
type Auxiliary_Press_Release_About_Press_Release struct {
	Entity

	// A count of
	AboutParagraphCount *uint `json:"aboutParagraphCount,omitempty"`

	//
	AboutParagraphs []Auxiliary_Press_Release_About `json:"aboutParagraphs,omitempty"`

	// A press release about cross
	Id *int `json:"id,omitempty"`

	// A press release about's internal
	PressReleaseAboutId *int `json:"pressReleaseAboutId,omitempty"`

	// A count of
	PressReleaseCount *uint `json:"pressReleaseCount,omitempty"`

	// A press release internal identifier.
	PressReleaseId *int `json:"pressReleaseId,omitempty"`

	//
	PressReleases []Auxiliary_Press_Release `json:"pressReleases,omitempty"`

	// The number that associated an about
	SortOrder *int `json:"sortOrder,omitempty"`
}

//
type Auxiliary_Press_Release_Contact struct {
	Entity

	// A press release contact's email
	Email *string `json:"email,omitempty"`

	// A press release contact's first name.
	FirstName *string `json:"firstName,omitempty"`

	// A press release contact's internal
	Id *int `json:"id,omitempty"`

	// A press release contact's last name.
	LastName *string `json:"lastName,omitempty"`

	// A press release contact's phone
	Phone *string `json:"phone,omitempty"`

	// A press release contact's
	ProfessionalTitle *string `json:"professionalTitle,omitempty"`
}

//
type Auxiliary_Press_Release_Contact_Press_Release struct {
	Entity

	// A count of
	ContactCount *uint `json:"contactCount,omitempty"`

	//
	Contacts []Auxiliary_Press_Release_Contact `json:"contacts,omitempty"`

	// A press release contact cross
	Id *int `json:"id,omitempty"`

	// A press release contact's internal
	PressReleaseContactId *int `json:"pressReleaseContactId,omitempty"`

	// A count of
	PressReleaseCount *uint `json:"pressReleaseCount,omitempty"`

	// A press release internal identifier.
	PressReleaseId *int `json:"pressReleaseId,omitempty"`

	//
	PressReleases []Auxiliary_Press_Release `json:"pressReleases,omitempty"`

	// The number that associated a contact
	SortOrder *int `json:"sortOrder,omitempty"`
}

//
type Auxiliary_Press_Release_Content struct {
	Entity

	// the id of a single press release
	Id *int `json:"id,omitempty"`

	// the press release id that the content
	PressReleaseId *int `json:"pressReleaseId,omitempty"`

	// the content of a press release
	Text *string `json:"text,omitempty"`
}

//
type Auxiliary_Press_Release_Media_Partner struct {
	Entity

	// A press release media partner's
	Id *int `json:"id,omitempty"`

	// A press release media partner's name.
	Name *string `json:"name,omitempty"`
}

//
type Auxiliary_Press_Release_Media_Partner_Press_Release struct {
	Entity

	// A press release media partner cross
	Id *int `json:"id,omitempty"`

	// A count of
	MediaPartnerCount *uint `json:"mediaPartnerCount,omitempty"`

	// A press release media partner's
	MediaPartnerId *int `json:"mediaPartnerId,omitempty"`

	//
	MediaPartners []Auxiliary_Press_Release_Media_Partner `json:"mediaPartners,omitempty"`

	// A count of
	PressReleaseCount *uint `json:"pressReleaseCount,omitempty"`

	// A press release internal identifier.
	PressReleaseId *int `json:"pressReleaseId,omitempty"`

	//
	PressReleases []Auxiliary_Press_Release `json:"pressReleases,omitempty"`
}

// The SoftLayer_Auxiliary_Shipping_Courier data type contains general information relating the different (major) couriers that SoftLayer may use for shipping.
type Auxiliary_Shipping_Courier struct {
	Entity

	// The unique id of the shipping courier.
	Id *int `json:"id,omitempty"`

	// The unique keyname of the shipping courier.
	KeyName *string `json:"keyName,omitempty"`

	// The name of the shipping courier.
	Name *string `json:"name,omitempty"`

	// The url to shipping courier's website.
	Url *string `json:"url,omitempty"`
}

//
type Auxiliary_Shipping_Courier_Type struct {
	Entity

	//
	Courier []Auxiliary_Shipping_Courier `json:"courier,omitempty"`

	// A count of
	CourierCount *uint `json:"courierCount,omitempty"`

	//
	Description *string `json:"description,omitempty"`

	//
	Id *int `json:"id,omitempty"`

	//
	KeyName *string `json:"keyName,omitempty"`

	//
	Name *string `json:"name,omitempty"`
}
