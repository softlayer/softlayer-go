package datatypes

import "time"

type User_Access_Facility_Log struct {
	Entity

	Account     *Account                       `json:"account:omitempty"`
	AccountId   *int                           `json:"accountId:omitempty"`
	Datacenter  *Location                      `json:"datacenter:omitempty"`
	Description *string                        `json:"description:omitempty"`
	Hardware    *Hardware                      `json:"hardware:omitempty"`
	HardwareId  *int                           `json:"hardwareId:omitempty"`
	Id          *int                           `json:"id:omitempty"`
	LocationId  *int                           `json:"locationId:omitempty"`
	LogType     *User_Access_Facility_Log_Type `json:"logType:omitempty"`
	TimeIn      *time.Time                     `json:"timeIn:omitempty"`
	TimeOut     *time.Time                     `json:"timeOut:omitempty"`
	Visitor     *Entity                        `json:"visitor:omitempty"`
}

type User_Access_Facility_Log_Type struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type User_Access_Facility_Visitor struct {
	Entity

	CompanyName *string                            `json:"companyName:omitempty"`
	FirstName   *string                            `json:"firstName:omitempty"`
	LastName    *string                            `json:"lastName:omitempty"`
	TypeId      *int                               `json:"typeId:omitempty"`
	VisitorType *User_Access_Facility_Visitor_Type `json:"visitorType:omitempty"`
}

type User_Access_Facility_Visitor_Type struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type User_Customer struct {
	User_Interface

	Account                                  *Account                                      `json:"account:omitempty"`
	AccountId                                *int                                          `json:"accountId:omitempty"`
	ActionCount                              *uint                                         `json:"actionCount:omitempty"`
	Actions                                  []User_Permission_Action                      `json:"actions:omitempty"`
	AdditionalEmailCount                     *uint                                         `json:"additionalEmailCount:omitempty"`
	AdditionalEmails                         []User_Customer_AdditionalEmail               `json:"additionalEmails:omitempty"`
	Address1                                 *string                                       `json:"address1:omitempty"`
	Address2                                 *string                                       `json:"address2:omitempty"`
	Aim                                      *string                                       `json:"aim:omitempty"`
	AlternatePhone                           *string                                       `json:"alternatePhone:omitempty"`
	ApiAuthenticationKeyCount                *uint                                         `json:"apiAuthenticationKeyCount:omitempty"`
	ApiAuthenticationKeys                    []User_Customer_ApiAuthentication             `json:"apiAuthenticationKeys:omitempty"`
	AuthenticationToken                      *Container_User_Authentication_Token          `json:"authenticationToken:omitempty"`
	CdnAccountCount                          *uint                                         `json:"cdnAccountCount:omitempty"`
	CdnAccounts                              []Network_ContentDelivery_Account             `json:"cdnAccounts:omitempty"`
	ChildUserCount                           *uint                                         `json:"childUserCount:omitempty"`
	ChildUsers                               []User_Customer                               `json:"childUsers:omitempty"`
	City                                     *string                                       `json:"city:omitempty"`
	ClosedTicketCount                        *uint                                         `json:"closedTicketCount:omitempty"`
	ClosedTickets                            []Ticket                                      `json:"closedTickets:omitempty"`
	CompanyName                              *string                                       `json:"companyName:omitempty"`
	Country                                  *string                                       `json:"country:omitempty"`
	CreateDate                               *time.Time                                    `json:"createDate:omitempty"`
	DaylightSavingsTimeFlag                  *bool                                         `json:"daylightSavingsTimeFlag:omitempty"`
	DenyAllResourceAccessOnCreateFlag        *bool                                         `json:"denyAllResourceAccessOnCreateFlag:omitempty"`
	DisplayName                              *string                                       `json:"displayName:omitempty"`
	Email                                    *string                                       `json:"email:omitempty"`
	ExternalBindingCount                     *uint                                         `json:"externalBindingCount:omitempty"`
	ExternalBindings                         []User_External_Binding                       `json:"externalBindings:omitempty"`
	FirstName                                *string                                       `json:"firstName:omitempty"`
	ForumPasswordHash                        *string                                       `json:"forumPasswordHash:omitempty"`
	Hardware                                 []Hardware                                    `json:"hardware:omitempty"`
	HardwareCount                            *uint                                         `json:"hardwareCount:omitempty"`
	HardwareNotificationCount                *uint                                         `json:"hardwareNotificationCount:omitempty"`
	HardwareNotifications                    []User_Customer_Notification_Hardware         `json:"hardwareNotifications:omitempty"`
	HasAcknowledgedSupportPolicyFlag         *bool                                         `json:"hasAcknowledgedSupportPolicyFlag:omitempty"`
	HasFullHardwareAccessFlag                *bool                                         `json:"hasFullHardwareAccessFlag:omitempty"`
	HasFullVirtualGuestAccessFlag            *bool                                         `json:"hasFullVirtualGuestAccessFlag:omitempty"`
	Icq                                      *string                                       `json:"icq:omitempty"`
	Id                                       *int                                          `json:"id:omitempty"`
	IpAddressRestriction                     *string                                       `json:"ipAddressRestriction:omitempty"`
	LastName                                 *string                                       `json:"lastName:omitempty"`
	LayoutProfileCount                       *uint                                         `json:"layoutProfileCount:omitempty"`
	LayoutProfiles                           []Layout_Profile                              `json:"layoutProfiles:omitempty"`
	Locale                                   *Locale                                       `json:"locale:omitempty"`
	LocaleId                                 *int                                          `json:"localeId:omitempty"`
	LoginAttemptCount                        *uint                                         `json:"loginAttemptCount:omitempty"`
	LoginAttempts                            []User_Customer_Access_Authentication         `json:"loginAttempts:omitempty"`
	ManagedByFederationFlag                  *bool                                         `json:"managedByFederationFlag:omitempty"`
	ManagedByOpenIdConnectFlag               *bool                                         `json:"managedByOpenIdConnectFlag:omitempty"`
	MobileDeviceCount                        *uint                                         `json:"mobileDeviceCount:omitempty"`
	MobileDevices                            []User_Customer_MobileDevice                  `json:"mobileDevices:omitempty"`
	ModifyDate                               *time.Time                                    `json:"modifyDate:omitempty"`
	Msn                                      *string                                       `json:"msn:omitempty"`
	NameId                                   *string                                       `json:"nameId:omitempty"`
	NotificationSubscriberCount              *uint                                         `json:"notificationSubscriberCount:omitempty"`
	NotificationSubscribers                  []Notification_Subscriber                     `json:"notificationSubscribers:omitempty"`
	OfficePhone                              *string                                       `json:"officePhone:omitempty"`
	OpenIdConnectUserName                    *string                                       `json:"openIdConnectUserName:omitempty"`
	OpenTicketCount                          *uint                                         `json:"openTicketCount:omitempty"`
	OpenTickets                              []Ticket                                      `json:"openTickets:omitempty"`
	OverrideCount                            *uint                                         `json:"overrideCount:omitempty"`
	Overrides                                []Network_Service_Vpn_Overrides               `json:"overrides:omitempty"`
	Parent                                   *User_Customer                                `json:"parent:omitempty"`
	ParentId                                 *int                                          `json:"parentId:omitempty"`
	PasswordExpireDate                       *time.Time                                    `json:"passwordExpireDate:omitempty"`
	PermissionCount                          *uint                                         `json:"permissionCount:omitempty"`
	PermissionSystemVersion                  *int                                          `json:"permissionSystemVersion:omitempty"`
	Permissions                              []User_Customer_CustomerPermission_Permission `json:"permissions:omitempty"`
	PostalCode                               *string                                       `json:"postalCode:omitempty"`
	PptpVpnAllowedFlag                       *bool                                         `json:"pptpVpnAllowedFlag:omitempty"`
	PreferenceCount                          *uint                                         `json:"preferenceCount:omitempty"`
	Preferences                              []User_Preference                             `json:"preferences:omitempty"`
	RoleCount                                *uint                                         `json:"roleCount:omitempty"`
	Roles                                    []User_Permission_Role                        `json:"roles:omitempty"`
	SalesforceUserLink                       *User_Customer_Link                           `json:"salesforceUserLink:omitempty"`
	SavedId                                  *string                                       `json:"savedId:omitempty"`
	SecondaryLoginManagementFlag             *bool                                         `json:"secondaryLoginManagementFlag:omitempty"`
	SecondaryLoginRequiredFlag               *bool                                         `json:"secondaryLoginRequiredFlag:omitempty"`
	SecondaryPasswordModifyDate              *time.Time                                    `json:"secondaryPasswordModifyDate:omitempty"`
	SecondaryPasswordTimeoutDays             *int                                          `json:"secondaryPasswordTimeoutDays:omitempty"`
	SecurityAnswerCount                      *uint                                         `json:"securityAnswerCount:omitempty"`
	SecurityAnswers                          []User_Customer_Security_Answer               `json:"securityAnswers:omitempty"`
	Sms                                      *string                                       `json:"sms:omitempty"`
	SslVpnAllowedFlag                        *bool                                         `json:"sslVpnAllowedFlag:omitempty"`
	State                                    *string                                       `json:"state:omitempty"`
	StatusDate                               *time.Time                                    `json:"statusDate:omitempty"`
	SubscriberCount                          *uint                                         `json:"subscriberCount:omitempty"`
	Subscribers                              []Notification_User_Subscriber                `json:"subscribers:omitempty"`
	SuccessfulLoginCount                     *uint                                         `json:"successfulLoginCount:omitempty"`
	SuccessfulLogins                         []User_Customer_Access_Authentication         `json:"successfulLogins:omitempty"`
	SupportPolicyAcknowledgementRequiredFlag *int                                          `json:"supportPolicyAcknowledgementRequiredFlag:omitempty"`
	SurveyCount                              *uint                                         `json:"surveyCount:omitempty"`
	SurveyRequiredFlag                       *bool                                         `json:"surveyRequiredFlag:omitempty"`
	Surveys                                  []Survey                                      `json:"surveys:omitempty"`
	TicketCount                              *uint                                         `json:"ticketCount:omitempty"`
	Tickets                                  []Ticket                                      `json:"tickets:omitempty"`
	Timezone                                 *Locale_Timezone                              `json:"timezone:omitempty"`
	TimezoneId                               *int                                          `json:"timezoneId:omitempty"`
	UnsuccessfulLoginCount                   *uint                                         `json:"unsuccessfulLoginCount:omitempty"`
	UnsuccessfulLogins                       []User_Customer_Access_Authentication         `json:"unsuccessfulLogins:omitempty"`
	UserLinkCount                            *uint                                         `json:"userLinkCount:omitempty"`
	UserLinks                                []User_Customer_Link                          `json:"userLinks:omitempty"`
	UserStatus                               *User_Customer_Status                         `json:"userStatus:omitempty"`
	UserStatusId                             *int                                          `json:"userStatusId:omitempty"`
	Username                                 *string                                       `json:"username:omitempty"`
	VirtualGuestCount                        *uint                                         `json:"virtualGuestCount:omitempty"`
	VirtualGuests                            []Virtual_Guest                               `json:"virtualGuests:omitempty"`
	VpnManualConfig                          *bool                                         `json:"vpnManualConfig:omitempty"`
	Yahoo                                    *string                                       `json:"yahoo:omitempty"`
}

type User_Customer_Access_Authentication struct {
	Entity

	CreateDate  *time.Time     `json:"createDate:omitempty"`
	IpAddress   *string        `json:"ipAddress:omitempty"`
	SuccessFlag *bool          `json:"successFlag:omitempty"`
	User        *User_Customer `json:"user:omitempty"`
	UserId      *int           `json:"userId:omitempty"`
	Username    *string        `json:"username:omitempty"`
}

type User_Customer_AdditionalEmail struct {
	Entity

	Email  *string        `json:"email:omitempty"`
	User   *User_Customer `json:"user:omitempty"`
	UserId *int           `json:"userId:omitempty"`
}

type User_Customer_ApiAuthentication struct {
	Entity

	AuthenticationKey    *string        `json:"authenticationKey:omitempty"`
	Id                   *int           `json:"id:omitempty"`
	IpAddressRestriction *string        `json:"ipAddressRestriction:omitempty"`
	TimestampKey         *int           `json:"timestampKey:omitempty"`
	User                 *User_Customer `json:"user:omitempty"`
	UserId               *int           `json:"userId:omitempty"`
}

type User_Customer_CustomerPermission_Permission struct {
	Entity

	Key     *string `json:"key:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type User_Customer_External_Binding struct {
	User_External_Binding

	User *User_Customer `json:"user:omitempty"`
}

type User_Customer_External_Binding_Attribute struct {
	User_External_Binding_Attribute
}

type User_Customer_External_Binding_Phone struct {
	User_Customer_External_Binding

	BindingStatus *string `json:"bindingStatus:omitempty"`
	PinLength     *string `json:"pinLength:omitempty"`
}

type User_Customer_External_Binding_Totp struct {
	User_Customer_External_Binding
}

type User_Customer_External_Binding_Type struct {
	User_External_Binding_Type
}

type User_Customer_External_Binding_Vendor struct {
	User_External_Binding_Vendor
}

type User_Customer_External_Binding_Verisign struct {
	User_Customer_External_Binding

	CredentialExpirationDate *string `json:"credentialExpirationDate:omitempty"`
	CredentialLastUpdateDate *string `json:"credentialLastUpdateDate:omitempty"`
	CredentialState          *string `json:"credentialState:omitempty"`
	CredentialType           *string `json:"credentialType:omitempty"`
}

type User_Customer_Invitation struct {
	Entity

	Code                       *string        `json:"code:omitempty"`
	CreateDate                 *time.Time     `json:"createDate:omitempty"`
	CreatorId                  *int           `json:"creatorId:omitempty"`
	CreatorType                *string        `json:"creatorType:omitempty"`
	Email                      *string        `json:"email:omitempty"`
	ExistingBlueIdFlag         *int           `json:"existingBlueIdFlag:omitempty"`
	ExpirationDate             *time.Time     `json:"expirationDate:omitempty"`
	Id                         *int           `json:"id:omitempty"`
	IsFederatedEmailDomainFlag *int           `json:"isFederatedEmailDomainFlag:omitempty"`
	ModifyDate                 *time.Time     `json:"modifyDate:omitempty"`
	ResponseDate               *time.Time     `json:"responseDate:omitempty"`
	StatusId                   *int           `json:"statusId:omitempty"`
	User                       *User_Customer `json:"user:omitempty"`
	UserId                     *int           `json:"userId:omitempty"`
}

type User_Customer_Link struct {
	Entity

	CreateDate                    *time.Time        `json:"createDate:omitempty"`
	DefaultFlag                   *int              `json:"defaultFlag:omitempty"`
	DestinationUserAlphanumericId *string           `json:"destinationUserAlphanumericId:omitempty"`
	DestinationUserId             *int              `json:"destinationUserId:omitempty"`
	Id                            *int              `json:"id:omitempty"`
	ServiceProvider               *Service_Provider `json:"serviceProvider:omitempty"`
	ServiceProviderId             *int              `json:"serviceProviderId:omitempty"`
	User                          *User_Customer    `json:"user:omitempty"`
	UserId                        *int              `json:"userId:omitempty"`
}

type User_Customer_Link_ThePlanet struct {
	User_Customer_Link
}

type User_Customer_MobileDevice struct {
	Entity

	AvailablePushNotificationSubscriptionCount *uint                                       `json:"availablePushNotificationSubscriptionCount:omitempty"`
	AvailablePushNotificationSubscriptions     []Notification                              `json:"availablePushNotificationSubscriptions:omitempty"`
	CreateDate                                 *time.Time                                  `json:"createDate:omitempty"`
	Customer                                   *User_Customer                              `json:"customer:omitempty"`
	DisplayResolutionXxY                       *string                                     `json:"displayResolutionXxY:omitempty"`
	Id                                         *int                                        `json:"id:omitempty"`
	MobileDeviceTypeId                         *int                                        `json:"mobileDeviceTypeId:omitempty"`
	MobileOperatingSystemId                    *int                                        `json:"mobileOperatingSystemId:omitempty"`
	ModelNumber                                *string                                     `json:"modelNumber:omitempty"`
	ModifyDate                                 *time.Time                                  `json:"modifyDate:omitempty"`
	Name                                       *string                                     `json:"name:omitempty"`
	OperatingSystem                            *User_Customer_MobileDevice_OperatingSystem `json:"operatingSystem:omitempty"`
	PhoneNumber                                *string                                     `json:"phoneNumber:omitempty"`
	PushNotificationSubscriptionCount          *uint                                       `json:"pushNotificationSubscriptionCount:omitempty"`
	PushNotificationSubscriptions              []Notification_User_Subscriber              `json:"pushNotificationSubscriptions:omitempty"`
	SerialNumber                               *string                                     `json:"serialNumber:omitempty"`
	Token                                      *string                                     `json:"token:omitempty"`
	Type                                       *User_Customer_MobileDevice_Type            `json:"type:omitempty"`
	UserId                                     *int                                        `json:"userId:omitempty"`
}

type User_Customer_MobileDevice_OperatingSystem struct {
	Entity

	BuildVersion *int       `json:"buildVersion:omitempty"`
	CreateDate   *time.Time `json:"createDate:omitempty"`
	Description  *string    `json:"description:omitempty"`
	Id           *int       `json:"id:omitempty"`
	MajorVersion *int       `json:"majorVersion:omitempty"`
	MinorVersion *int       `json:"minorVersion:omitempty"`
	ModifyDate   *time.Time `json:"modifyDate:omitempty"`
	Name         *string    `json:"name:omitempty"`
}

type User_Customer_MobileDevice_Type struct {
	Entity

	CreateDate  *time.Time `json:"createDate:omitempty"`
	Description *string    `json:"description:omitempty"`
	Id          *int       `json:"id:omitempty"`
	ModifyDate  *time.Time `json:"modifyDate:omitempty"`
	Name        *string    `json:"name:omitempty"`
}

type User_Customer_Notification_Hardware struct {
	Entity

	Hardware   *Hardware      `json:"hardware:omitempty"`
	HardwareId *int           `json:"hardwareId:omitempty"`
	Id         *int           `json:"id:omitempty"`
	User       *User_Customer `json:"user:omitempty"`
	UserId     *int           `json:"userId:omitempty"`
}

type User_Customer_Notification_Virtual_Guest struct {
	Entity

	Guest   *Virtual_Guest `json:"guest:omitempty"`
	GuestId *int           `json:"guestId:omitempty"`
	Id      *int           `json:"id:omitempty"`
	User    *User_Customer `json:"user:omitempty"`
	UserId  *int           `json:"userId:omitempty"`
}

type User_Customer_OpenIdConnect struct {
	User_Customer
}

type User_Customer_Prospect struct {
	Entity

	Account               *Account                     `json:"account:omitempty"`
	AssignedEmployeeCount *uint                        `json:"assignedEmployeeCount:omitempty"`
	AssignedEmployees     []User_Employee              `json:"assignedEmployees:omitempty"`
	QuoteCount            *uint                        `json:"quoteCount:omitempty"`
	Quotes                []Billing_Order_Quote        `json:"quotes:omitempty"`
	Type                  *User_Customer_Prospect_Type `json:"type:omitempty"`
}

type User_Customer_Prospect_ServiceProvider_EnrollRequest struct {
	Entity

	AccountId                   *int                   `json:"accountId:omitempty"`
	Address1                    *string                `json:"address1:omitempty"`
	Address2                    *string                `json:"address2:omitempty"`
	CardAccountNumber           *string                `json:"cardAccountNumber:omitempty"`
	CardExpirationMonth         *string                `json:"cardExpirationMonth:omitempty"`
	CardExpirationYear          *string                `json:"cardExpirationYear:omitempty"`
	CardType                    *string                `json:"cardType:omitempty"`
	CardVerificationNumber      *string                `json:"cardVerificationNumber:omitempty"`
	City                        *string                `json:"city:omitempty"`
	CompanyName                 *string                `json:"companyName:omitempty"`
	CompanyType                 *Catalyst_Company_Type `json:"companyType:omitempty"`
	CompanyTypeId               *int                   `json:"companyTypeId:omitempty"`
	CompanyUrl                  *string                `json:"companyUrl:omitempty"`
	ContactEmail                *string                `json:"contactEmail:omitempty"`
	ContactFirstName            *string                `json:"contactFirstName:omitempty"`
	ContactLastName             *string                `json:"contactLastName:omitempty"`
	ContactPhone                *string                `json:"contactPhone:omitempty"`
	Country                     *string                `json:"country:omitempty"`
	CustomerProspectId          *int                   `json:"customerProspectId:omitempty"`
	DeviceFingerprintId         *string                `json:"deviceFingerprintId:omitempty"`
	Email                       *string                `json:"email:omitempty"`
	ExistingCustomerFlag        *bool                  `json:"existingCustomerFlag:omitempty"`
	FirstName                   *string                `json:"firstName:omitempty"`
	IbmPartnerWorldId           *string                `json:"ibmPartnerWorldId:omitempty"`
	IbmPartnerWorldMemberFlag   *bool                  `json:"ibmPartnerWorldMemberFlag:omitempty"`
	LastName                    *string                `json:"lastName:omitempty"`
	MasterAgreementCompleteFlag *bool                  `json:"masterAgreementCompleteFlag:omitempty"`
	OfficePhone                 *string                `json:"officePhone:omitempty"`
	PostalCode                  *string                `json:"postalCode:omitempty"`
	ServiceProviderAddendumFlag *bool                  `json:"serviceProviderAddendumFlag:omitempty"`
	State                       *string                `json:"state:omitempty"`
	SurveyResponses             []Survey_Response      `json:"surveyResponses:omitempty"`
	VatId                       *string                `json:"vatId:omitempty"`
}

type User_Customer_Prospect_Type struct {
	Entity

	CreateDate  *time.Time `json:"createDate:omitempty"`
	Description *string    `json:"description:omitempty"`
	Id          *int       `json:"id:omitempty"`
	KeyName     *string    `json:"keyName:omitempty"`
	ModifyDate  *time.Time `json:"modifyDate:omitempty"`
	Name        *string    `json:"name:omitempty"`
}

type User_Customer_Security_Answer struct {
	Entity

	Answer     *string                 `json:"answer:omitempty"`
	Id         *int                    `json:"id:omitempty"`
	Question   *User_Security_Question `json:"question:omitempty"`
	QuestionId *int                    `json:"questionId:omitempty"`
	User       *User_Customer          `json:"user:omitempty"`
	UserId     *int                    `json:"userId:omitempty"`
}

type User_Customer_Status struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type User_Employee struct {
	User_Interface

	ActionCount                    *uint                     `json:"actionCount:omitempty"`
	Actions                        []User_Permission_Action  `json:"actions:omitempty"`
	ChatTranscript                 []Ticket_Chat             `json:"chatTranscript:omitempty"`
	ChatTranscriptCount            *uint                     `json:"chatTranscriptCount:omitempty"`
	DisplayName                    *string                   `json:"displayName:omitempty"`
	Email                          *string                   `json:"email:omitempty"`
	EmployeeDepartment             *User_Employee_Department `json:"employeeDepartment:omitempty"`
	EmployeeDepartmentId           *int                      `json:"employeeDepartmentId:omitempty"`
	FirstName                      *string                   `json:"firstName:omitempty"`
	LastName                       *string                   `json:"lastName:omitempty"`
	LayoutProfileCount             *uint                     `json:"layoutProfileCount:omitempty"`
	LayoutProfiles                 []Layout_Profile          `json:"layoutProfiles:omitempty"`
	MetricTrackingObject           *Metric_Tracking_Object   `json:"metricTrackingObject:omitempty"`
	OfficePhone                    *string                   `json:"officePhone:omitempty"`
	RoleCount                      *uint                     `json:"roleCount:omitempty"`
	Roles                          []User_Permission_Role    `json:"roles:omitempty"`
	TicketActivities               []Ticket_Activity         `json:"ticketActivities:omitempty"`
	TicketActivityCount            *uint                     `json:"ticketActivityCount:omitempty"`
	TicketAttachmentReferenceCount *uint                     `json:"ticketAttachmentReferenceCount:omitempty"`
	TicketAttachmentReferences     []Ticket_Attachment       `json:"ticketAttachmentReferences:omitempty"`
	Username                       *string                   `json:"username:omitempty"`
}

type User_Employee_Department struct {
	Entity

	Name *string `json:"name:omitempty"`
}

type User_External_Binding struct {
	Entity

	Active         *bool                             `json:"active:omitempty"`
	AttributeCount *uint                             `json:"attributeCount:omitempty"`
	Attributes     []User_External_Binding_Attribute `json:"attributes:omitempty"`
	BillingItem    *Billing_Item                     `json:"billingItem:omitempty"`
	CreateDate     *time.Time                        `json:"createDate:omitempty"`
	ExternalId     *string                           `json:"externalId:omitempty"`
	Id             *int                              `json:"id:omitempty"`
	Note           *string                           `json:"note:omitempty"`
	Password       *string                           `json:"password:omitempty"`
	Type           *User_External_Binding_Type       `json:"type:omitempty"`
	TypeId         *int                              `json:"typeId:omitempty"`
	UserId         *int                              `json:"userId:omitempty"`
	Vendor         *User_External_Binding_Vendor     `json:"vendor:omitempty"`
	VendorId       *int                              `json:"vendorId:omitempty"`
}

type User_External_Binding_Attribute struct {
	Entity

	ExternalBinding *User_External_Binding `json:"externalBinding:omitempty"`
	Value           *string                `json:"value:omitempty"`
}

type User_External_Binding_Type struct {
	Entity

	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type User_External_Binding_Vendor struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type User_Interface struct {
	Entity
}

type User_Permission_Action struct {
	Entity

	CreateDate  *time.Time `json:"createDate:omitempty"`
	Description *string    `json:"description:omitempty"`
	Id          *int       `json:"id:omitempty"`
	KeyName     *string    `json:"keyName:omitempty"`
	ModifyDate  *time.Time `json:"modifyDate:omitempty"`
	Name        *string    `json:"name:omitempty"`
}

type User_Permission_Group struct {
	Entity

	Account        *Account                    `json:"account:omitempty"`
	AccountId      *int                        `json:"accountId:omitempty"`
	ActionCount    *uint                       `json:"actionCount:omitempty"`
	Actions        []User_Permission_Action    `json:"actions:omitempty"`
	CreateDate     *time.Time                  `json:"createDate:omitempty"`
	Description    *string                     `json:"description:omitempty"`
	ExpirationDate *time.Time                  `json:"expirationDate:omitempty"`
	Id             *int                        `json:"id:omitempty"`
	ModifyDate     *time.Time                  `json:"modifyDate:omitempty"`
	Name           *string                     `json:"name:omitempty"`
	RoleCount      *uint                       `json:"roleCount:omitempty"`
	Roles          []User_Permission_Role      `json:"roles:omitempty"`
	Type           *User_Permission_Group_Type `json:"type:omitempty"`
	TypeId         *int                        `json:"typeId:omitempty"`
}

type User_Permission_Group_Type struct {
	Entity

	CreateDate *time.Time              `json:"createDate:omitempty"`
	GroupCount *uint                   `json:"groupCount:omitempty"`
	Groups     []User_Permission_Group `json:"groups:omitempty"`
	Id         *int                    `json:"id:omitempty"`
	KeyName    *string                 `json:"keyName:omitempty"`
	ModifyDate *time.Time              `json:"modifyDate:omitempty"`
	Name       *string                 `json:"name:omitempty"`
}

type User_Permission_Role struct {
	Entity

	Account            *Account                 `json:"account:omitempty"`
	AccountId          *int                     `json:"accountId:omitempty"`
	ActionCount        *uint                    `json:"actionCount:omitempty"`
	Actions            []User_Permission_Action `json:"actions:omitempty"`
	CreateDate         *time.Time               `json:"createDate:omitempty"`
	Description        *string                  `json:"description:omitempty"`
	GroupCount         *uint                    `json:"groupCount:omitempty"`
	Groups             []User_Permission_Group  `json:"groups:omitempty"`
	Id                 *int                     `json:"id:omitempty"`
	ModifyDate         *time.Time               `json:"modifyDate:omitempty"`
	Name               *string                  `json:"name:omitempty"`
	NewUserDefaultFlag *int                     `json:"newUserDefaultFlag:omitempty"`
	SystemFlag         *int                     `json:"systemFlag:omitempty"`
	UserCount          *uint                    `json:"userCount:omitempty"`
	Users              []User_Customer          `json:"users:omitempty"`
}

type User_Preference struct {
	Entity

	Description *string               `json:"description:omitempty"`
	Type        *User_Preference_Type `json:"type:omitempty"`
	Value       *string               `json:"value:omitempty"`
}

type User_Preference_Type struct {
	Entity

	Description  *string `json:"description:omitempty"`
	KeyName      *string `json:"keyName:omitempty"`
	Name         *string `json:"name:omitempty"`
	ValueExample *string `json:"valueExample:omitempty"`
}

type User_Security_Question struct {
	Entity

	DisplayOrder *int    `json:"displayOrder:omitempty"`
	Id           *int    `json:"id:omitempty"`
	Question     *string `json:"question:omitempty"`
	Viewable     *int    `json:"viewable:omitempty"`
}
