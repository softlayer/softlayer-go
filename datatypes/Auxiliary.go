package datatypes

import "time"

type Auxiliary_Marketing_Event struct {
	Entity

	CreateDate  *time.Time `json:"createDate,omitempty"`
	EnabledFlag *int       `json:"enabledFlag,omitempty"`
	EndDate     *time.Time `json:"endDate,omitempty"`
	Location    *string    `json:"location,omitempty"`
	ModifyDate  *time.Time `json:"modifyDate,omitempty"`
	StartDate   *time.Time `json:"startDate,omitempty"`
	Title       *string    `json:"title,omitempty"`
	Url         *string    `json:"url,omitempty"`
}

type Auxiliary_Network_Status struct {
	Entity
}

type Auxiliary_Notification_Emergency struct {
	Entity

	CreateDate       *time.Time                                  `json:"createDate,omitempty"`
	Device           *string                                     `json:"device,omitempty"`
	Duration         *string                                     `json:"duration,omitempty"`
	Id               *int                                        `json:"id,omitempty"`
	Location         *string                                     `json:"location,omitempty"`
	Message          *string                                     `json:"message,omitempty"`
	ModifyDate       *time.Time                                  `json:"modifyDate,omitempty"`
	ServicesAffected *string                                     `json:"servicesAffected,omitempty"`
	Signature        *Auxiliary_Notification_Emergency_Signature `json:"signature,omitempty"`
	StartDate        *time.Time                                  `json:"startDate,omitempty"`
	Status           *Auxiliary_Notification_Emergency_Status    `json:"status,omitempty"`
	StatusId         *int                                        `json:"statusId,omitempty"`
}

type Auxiliary_Notification_Emergency_Signature struct {
	Entity

	Name *string `json:"name,omitempty"`
}

type Auxiliary_Notification_Emergency_Status struct {
	Entity

	Name *string `json:"name,omitempty"`
}

type Auxiliary_Press_Release struct {
	Entity

	About                []Auxiliary_Press_Release_About_Press_Release         `json:"about,omitempty"`
	AboutCount           *uint                                                 `json:"aboutCount,omitempty"`
	ContactCount         *uint                                                 `json:"contactCount,omitempty"`
	Contacts             []Auxiliary_Press_Release_Contact_Press_Release       `json:"contacts,omitempty"`
	Id                   *int                                                  `json:"id,omitempty"`
	MediaPartnerCount    *uint                                                 `json:"mediaPartnerCount,omitempty"`
	MediaPartners        []Auxiliary_Press_Release_Media_Partner_Press_Release `json:"mediaPartners,omitempty"`
	PressReleaseContent  *Auxiliary_Press_Release_Content                      `json:"pressReleaseContent,omitempty"`
	PublishDate          *time.Time                                            `json:"publishDate,omitempty"`
	ReleaseLocation      *string                                               `json:"releaseLocation,omitempty"`
	SubTitle             *string                                               `json:"subTitle,omitempty"`
	Title                *string                                               `json:"title,omitempty"`
	WebsiteHighlightFlag *bool                                                 `json:"websiteHighlightFlag,omitempty"`
}

type Auxiliary_Press_Release_About struct {
	Entity

	Content *string `json:"content,omitempty"`
	Id      *int    `json:"id,omitempty"`
	Title   *string `json:"title,omitempty"`
}

type Auxiliary_Press_Release_About_Press_Release struct {
	Entity

	AboutParagraphCount *uint                           `json:"aboutParagraphCount,omitempty"`
	AboutParagraphs     []Auxiliary_Press_Release_About `json:"aboutParagraphs,omitempty"`
	Id                  *int                            `json:"id,omitempty"`
	PressReleaseAboutId *int                            `json:"pressReleaseAboutId,omitempty"`
	PressReleaseCount   *uint                           `json:"pressReleaseCount,omitempty"`
	PressReleaseId      *int                            `json:"pressReleaseId,omitempty"`
	PressReleases       []Auxiliary_Press_Release       `json:"pressReleases,omitempty"`
	SortOrder           *int                            `json:"sortOrder,omitempty"`
}

type Auxiliary_Press_Release_Contact struct {
	Entity

	Email             *string `json:"email,omitempty"`
	FirstName         *string `json:"firstName,omitempty"`
	Id                *int    `json:"id,omitempty"`
	LastName          *string `json:"lastName,omitempty"`
	Phone             *string `json:"phone,omitempty"`
	ProfessionalTitle *string `json:"professionalTitle,omitempty"`
}

type Auxiliary_Press_Release_Contact_Press_Release struct {
	Entity

	ContactCount          *uint                             `json:"contactCount,omitempty"`
	Contacts              []Auxiliary_Press_Release_Contact `json:"contacts,omitempty"`
	Id                    *int                              `json:"id,omitempty"`
	PressReleaseContactId *int                              `json:"pressReleaseContactId,omitempty"`
	PressReleaseCount     *uint                             `json:"pressReleaseCount,omitempty"`
	PressReleaseId        *int                              `json:"pressReleaseId,omitempty"`
	PressReleases         []Auxiliary_Press_Release         `json:"pressReleases,omitempty"`
	SortOrder             *int                              `json:"sortOrder,omitempty"`
}

type Auxiliary_Press_Release_Content struct {
	Entity

	Id             *int    `json:"id,omitempty"`
	PressReleaseId *int    `json:"pressReleaseId,omitempty"`
	Text           *string `json:"text,omitempty"`
}

type Auxiliary_Press_Release_Media_Partner struct {
	Entity

	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Auxiliary_Press_Release_Media_Partner_Press_Release struct {
	Entity

	Id                *int                                    `json:"id,omitempty"`
	MediaPartnerCount *uint                                   `json:"mediaPartnerCount,omitempty"`
	MediaPartnerId    *int                                    `json:"mediaPartnerId,omitempty"`
	MediaPartners     []Auxiliary_Press_Release_Media_Partner `json:"mediaPartners,omitempty"`
	PressReleaseCount *uint                                   `json:"pressReleaseCount,omitempty"`
	PressReleaseId    *int                                    `json:"pressReleaseId,omitempty"`
	PressReleases     []Auxiliary_Press_Release               `json:"pressReleases,omitempty"`
}

type Auxiliary_Shipping_Courier struct {
	Entity

	Id      *int    `json:"id,omitempty"`
	KeyName *string `json:"keyName,omitempty"`
	Name    *string `json:"name,omitempty"`
	Url     *string `json:"url,omitempty"`
}

type Auxiliary_Shipping_Courier_Type struct {
	Entity

	Courier      []Auxiliary_Shipping_Courier `json:"courier,omitempty"`
	CourierCount *uint                        `json:"courierCount,omitempty"`
	Description  *string                      `json:"description,omitempty"`
	Id           *int                         `json:"id,omitempty"`
	KeyName      *string                      `json:"keyName,omitempty"`
	Name         *string                      `json:"name,omitempty"`
}
