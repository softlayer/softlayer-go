package softlayer

import "github.ibm.com/riethm/gopherlayer/datatypes"

type User_Access_Facility_Log interface {
	Entity

	GetAccount() (datatypes.Account, error)
	GetDatacenter() (datatypes.Location, error)
	GetHardware() (datatypes.Hardware, error)
	GetLogType() (datatypes.User_Access_Facility_Log_Type, error)
	GetVisitor() (datatypes.Entity, error)
}

type User_Access_Facility_Log_Type interface {
	Entity
}

type User_Access_Facility_Visitor interface {
	Entity

	GetVisitorType() (datatypes.User_Access_Facility_Visitor_Type, error)
}

type User_Access_Facility_Visitor_Type interface {
	Entity
}

type User_Customer interface {
	User_Interface

	AcknowledgeSupportPolicy() error
	AddApiAuthenticationKey() (string, error)
	AddBulkHardwareAccess(hardwareIds int) (bool, error)
	AddBulkPortalPermission(permissions datatypes.User_Customer_CustomerPermission_Permission) (bool, error)
	AddBulkRoles(roles datatypes.User_Permission_Role) error
	AddBulkVirtualGuestAccess(virtualGuestIds int) (bool, error)
	AddExternalBinding(externalBinding datatypes.User_External_Binding) (datatypes.User_Customer_External_Binding, error)
	AddHardwareAccess(hardwareId int) (bool, error)
	AddNotificationSubscriber(notificationKeyName string) (bool, error)
	AddPortalPermission(permission datatypes.User_Customer_CustomerPermission_Permission) (bool, error)
	AddRole(role datatypes.User_Permission_Role) error
	AddVirtualGuestAccess(virtualGuestId int) (bool, error)
	ChangePreference(preferenceTypeKeyName string, value string) (datatypes.User_Preference, error)
	CheckExternalAuthenticationStatus(authenticationContainer datatypes.Container_User_Customer_External_Binding) (datatypes.Container_User_Customer_Portal_Token, error)
	CheckPhoneFactorAuthenticationForPasswordSet(passwordSet datatypes.Container_User_Customer_PasswordSet, authenticationContainer datatypes.Container_User_Customer_External_Binding) (bool, error)
	CreateNotificationSubscriber(keyName string, resourceTableId int) (bool, error)
	CreateObject(templateObject datatypes.User_Customer, password string, vpnPassword string) (datatypes.User_Customer, error)
	CreateSubscriberDeliveryMethods(notificationKeyName string, deliveryMethodKeyNames string) (bool, error)
	DeactivateNotificationSubscriber(keyName string, resourceTableId int) (bool, error)
	EditObject(templateObject datatypes.User_Customer) (bool, error)
	EditObjects(templateObjects datatypes.User_Customer) (bool, error)
	FindUserPreference(profileName string, containerKeyname string, preferenceKeyname string) (datatypes.Layout_Profile, error)
	GetActiveExternalAuthenticationVendors() (datatypes.Container_User_Customer_External_Binding_Vendor, error)
	GetAllowedHardwareIds() (int, error)
	GetAllowedVirtualGuestIds() (int, error)
	GetAuthenticationToken(token datatypes.Container_User_Authentication_Token) (datatypes.Container_User_Authentication_Token, error)
	GetDefaultSecurityQuestions(key string) (datatypes.User_Security_Question, error)
	GetHardwareCount() (int, error)
	GetImpersonationToken() (string, error)
	GetObject() (datatypes.User_Customer, error)
	GetPortalLoginToken(username string, password string, securityQuestionId int, securityQuestionAnswer string) (datatypes.Container_User_Customer_Portal_Token, error)
	GetPreference(preferenceTypeKeyName string) (datatypes.User_Preference, error)
	GetPreferenceTypes() (datatypes.User_Preference_Type, error)
	GetRequirementsForPasswordSet(passwordSet datatypes.Container_User_Customer_PasswordSet) (datatypes.Container_User_Customer_PasswordSet, error)
	GetSupportPolicyDocument() ([]byte, error)
	GetSupportPolicyName() (string, error)
	GetSupportedLocales() (datatypes.Locale, error)
	GetUserFromLostPasswordRequest(key string) (datatypes.User_Security_Question, error)
	GetUserIdForPasswordSet(key string) (int, error)
	GetUserPreferences(profileName string, containerKeyname string) (datatypes.Layout_Profile, error)
	GetVirtualGuestCount() (int, error)
	InTerminalStatus() (bool, error)
	InitiateExternalAuthentication(authenticationContainer datatypes.Container_User_Customer_External_Binding) (string, error)
	InitiatePortalPasswordChange(username string) (bool, error)
	InitiatePortalPasswordChangeByBrandAgent(username string) (bool, error)
	InviteUserToLinkOpenIdConnect(providerType string) (bool, error)
	IsMasterUser() (bool, error)
	IsValidForumPassword(password string) (bool, error)
	IsValidPortalPassword(password string) (bool, error)
	LostPassword(username string, email string) (bool, error)
	PerformExternalAuthentication(authenticationContainer datatypes.Container_User_Customer_External_Binding) (datatypes.Container_User_Customer_Portal_Token, error)
	ProcessPasswordSetRequest(passwordSet datatypes.Container_User_Customer_PasswordSet, authenticationContainer datatypes.Container_User_Customer_External_Binding) (datatypes.Container_User_Customer_PasswordSet, error)
	RemoveAllHardwareAccessForThisUser() (bool, error)
	RemoveAllVirtualAccessForThisUser() (bool, error)
	RemoveApiAuthenticationKey(keyId int) (bool, error)
	RemoveBulkHardwareAccess(hardwareIds int) (bool, error)
	RemoveBulkPortalPermission(permissions datatypes.User_Customer_CustomerPermission_Permission) (bool, error)
	RemoveBulkRoles(roles datatypes.User_Permission_Role) error
	RemoveBulkVirtualGuestAccess(virtualGuestIds int) (bool, error)
	RemoveExternalBinding(externalBinding datatypes.User_External_Binding) (bool, error)
	RemoveHardwareAccess(hardwareId int) (bool, error)
	RemovePortalPermission(permission datatypes.User_Customer_CustomerPermission_Permission) (bool, error)
	RemoveRole(role datatypes.User_Permission_Role) error
	RemoveVirtualGuestAccess(virtualGuestId int) (bool, error)
	ResetExpiredPassword(username string, password string, newPassword string, securityQuestionId int, securityQuestionAnswer string) (bool, error)
	SamlAuthenticate(accountId string, samlResponse string) (datatypes.Container_User_Customer_Portal_Token, error)
	SamlBeginAuthentication(accountId int) (string, error)
	SamlBeginLogout() (string, error)
	SamlLogout(samlResponse string) error
	SetPasswordFromLostPasswordRequest(key string, password string, securityAnswers datatypes.User_Customer_Security_Answer) (bool, error)
	UpdateForumPassword(password string) (bool, error)
	UpdateNotificationSubscriber(notificationKeyName string, active int) (bool, error)
	UpdatePassword(password string) (bool, error)
	UpdateSecurityAnswers(questions datatypes.User_Security_Question, answers string) (bool, error)
	UpdateSubscriberDeliveryMethod(notificationKeyName string, deliveryMethodKeyNames string, active int) (bool, error)
	UpdateVpnPassword(password string) (bool, error)
	UpdateVpnUser() (bool, error)
	ValidateAuthenticationToken(authenticationToken datatypes.Container_User_Authentication_Token) (datatypes.Container_User_Customer_Portal_Token, error)

	GetAccount() (datatypes.Account, error)
	GetActions() (datatypes.User_Permission_Action, error)
	GetAdditionalEmails() (datatypes.User_Customer_AdditionalEmail, error)
	GetApiAuthenticationKeys() (datatypes.User_Customer_ApiAuthentication, error)
	GetCdnAccounts() (datatypes.Network_ContentDelivery_Account, error)
	GetChildUsers() (datatypes.User_Customer, error)
	GetClosedTickets() (datatypes.Ticket, error)
	GetExternalBindings() (datatypes.User_External_Binding, error)
	GetHardware() (datatypes.Hardware, error)
	GetHardwareNotifications() (datatypes.User_Customer_Notification_Hardware, error)
	GetHasAcknowledgedSupportPolicyFlag() (bool, error)
	GetHasFullHardwareAccessFlag() (bool, error)
	GetHasFullVirtualGuestAccessFlag() (bool, error)
	GetLayoutProfiles() (datatypes.Layout_Profile, error)
	GetLocale() (datatypes.Locale, error)
	GetLoginAttempts() (datatypes.User_Customer_Access_Authentication, error)
	GetMobileDevices() (datatypes.User_Customer_MobileDevice, error)
	GetNotificationSubscribers() (datatypes.Notification_Subscriber, error)
	GetOpenTickets() (datatypes.Ticket, error)
	GetOverrides() (datatypes.Network_Service_Vpn_Overrides, error)
	GetParent() (datatypes.User_Customer, error)
	GetPermissions() (datatypes.User_Customer_CustomerPermission_Permission, error)
	GetPreferences() (datatypes.User_Preference, error)
	GetRoles() (datatypes.User_Permission_Role, error)
	GetSalesforceUserLink() (datatypes.User_Customer_Link, error)
	GetSecurityAnswers() (datatypes.User_Customer_Security_Answer, error)
	GetSubscribers() (datatypes.Notification_User_Subscriber, error)
	GetSuccessfulLogins() (datatypes.User_Customer_Access_Authentication, error)
	GetSupportPolicyAcknowledgementRequiredFlag() (int, error)
	GetSurveyRequiredFlag() (bool, error)
	GetSurveys() (datatypes.Survey, error)
	GetTickets() (datatypes.Ticket, error)
	GetTimezone() (datatypes.Locale_Timezone, error)
	GetUnsuccessfulLogins() (datatypes.User_Customer_Access_Authentication, error)
	GetUserLinks() (datatypes.User_Customer_Link, error)
	GetUserStatus() (datatypes.User_Customer_Status, error)
	GetVirtualGuests() (datatypes.Virtual_Guest, error)
}

type User_Customer_Access_Authentication interface {
	Entity

	GetUser() (datatypes.User_Customer, error)
}

type User_Customer_AdditionalEmail interface {
	Entity

	GetUser() (datatypes.User_Customer, error)
}

type User_Customer_ApiAuthentication interface {
	Entity

	EditObject(templateObject datatypes.User_Customer_ApiAuthentication) (datatypes.User_Customer_ApiAuthentication, error)
	GetObject() (datatypes.User_Customer_ApiAuthentication, error)

	GetUser() (datatypes.User_Customer, error)
}

type User_Customer_CustomerPermission_Permission interface {
	Entity

	GetAllObjects() (datatypes.User_Customer_CustomerPermission_Permission, error)
	GetObject() (datatypes.User_Customer_CustomerPermission_Permission, error)
}

type User_Customer_External_Binding interface {
	User_External_Binding

	Disable(reason string) (bool, error)
	Enable() (bool, error)
	GetObject() (datatypes.User_Customer_External_Binding, error)

	GetUser() (datatypes.User_Customer, error)
}

type User_Customer_External_Binding_Attribute interface {
	User_External_Binding_Attribute
}

type User_Customer_External_Binding_Phone interface {
	User_Customer_External_Binding

	CheckPhoneValidationResult(token string) (bool, error)
	Disable(reason string) (bool, error)
	Enable() (bool, error)
	GetAllAuthenticationModes() (string, error)
	GetAllAuthenticationPinModes(authenticationModeKeyName string) (string, error)
	GetAuthenticationMode() (datatypes.Container_User_Customer_External_Binding_Phone_Mode, error)
	GetObject() (datatypes.User_Customer_External_Binding_Phone, error)
	GetPhoneAppActivationCode() (datatypes.User_External_Binding_Attribute, error)
	GetPhoneData() (datatypes.Container_User_Data_Phone, error)
	RequestPhoneValidation(phoneData datatypes.Container_User_Data_Phone) (string, error)
	UpdateAuthenticationMode(mode datatypes.Container_User_Customer_External_Binding_Phone_Mode) (bool, error)
	UpdatePhone(phoneData datatypes.Container_User_Data_Phone) (bool, error)

	GetBindingStatus() (string, error)
	GetPinLength() (string, error)
}

type User_Customer_External_Binding_Totp interface {
	User_Customer_External_Binding

	Activate() (bool, error)
	Deactivate() (bool, error)
	GenerateSecretKey() (string, error)
	GetObject() (datatypes.User_Customer_External_Binding_Totp, error)
}

type User_Customer_External_Binding_Type interface {
	User_External_Binding_Type
}

type User_Customer_External_Binding_Vendor interface {
	User_External_Binding_Vendor

	GetObject() (datatypes.User_Customer_External_Binding_Vendor, error)
}

type User_Customer_External_Binding_Verisign interface {
	User_Customer_External_Binding

	DeleteObject() (bool, error)
	Disable(reason string) (bool, error)
	Enable() (bool, error)
	GetActivationCodeForMobileClient() (string, error)
	GetObject() (datatypes.User_Customer_External_Binding_Verisign, error)
	Unlock(securityCode string) (bool, error)
	ValidateCredentialId(userId int, externalId string) (bool, error)

	GetCredentialExpirationDate() (string, error)
	GetCredentialLastUpdateDate() (string, error)
	GetCredentialState() (string, error)
	GetCredentialType() (string, error)
}

type User_Customer_Invitation interface {
	Entity

	GetObject() (datatypes.User_Customer_Invitation, error)

	GetUser() (datatypes.User_Customer, error)
}

type User_Customer_Link interface {
	Entity

	GetServiceProvider() (datatypes.Service_Provider, error)
	GetUser() (datatypes.User_Customer, error)
}

type User_Customer_Link_ThePlanet interface {
	User_Customer_Link
}

type User_Customer_MobileDevice interface {
	Entity

	CreateObject(templateObject datatypes.User_Customer_MobileDevice) (datatypes.User_Customer_MobileDevice, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.User_Customer_MobileDevice) (bool, error)
	GetObject() (datatypes.User_Customer_MobileDevice, error)

	GetAvailablePushNotificationSubscriptions() (datatypes.Notification, error)
	GetCustomer() (datatypes.User_Customer, error)
	GetOperatingSystem() (datatypes.User_Customer_MobileDevice_OperatingSystem, error)
	GetPushNotificationSubscriptions() (datatypes.Notification_User_Subscriber, error)
	GetType() (datatypes.User_Customer_MobileDevice_Type, error)
}

type User_Customer_MobileDevice_OperatingSystem interface {
	Entity

	GetAllObjects() (datatypes.User_Customer_MobileDevice_OperatingSystem, error)
	GetObject() (datatypes.User_Customer_MobileDevice_OperatingSystem, error)
}

type User_Customer_MobileDevice_Type interface {
	Entity

	GetAllObjects() (datatypes.User_Customer_MobileDevice_Type, error)
	GetObject() (datatypes.User_Customer_MobileDevice_Type, error)
}

type User_Customer_Notification_Hardware interface {
	Entity

	CreateObject(templateObject datatypes.User_Customer_Notification_Hardware) (datatypes.User_Customer_Notification_Hardware, error)
	CreateObjects(templateObjects datatypes.User_Customer_Notification_Hardware) (datatypes.Dns_Domain, error)
	DeleteObjects(templateObjects datatypes.User_Customer_Notification_Hardware) (bool, error)
	FindByHardwareId(hardwareId int) (datatypes.User_Customer_Notification_Hardware, error)
	GetObject() (datatypes.User_Customer_Notification_Hardware, error)

	GetHardware() (datatypes.Hardware, error)
	GetUser() (datatypes.User_Customer, error)
}

type User_Customer_Notification_Virtual_Guest interface {
	Entity

	CreateObject(templateObject datatypes.User_Customer_Notification_Virtual_Guest) (datatypes.User_Customer_Notification_Virtual_Guest, error)
	CreateObjects(templateObjects datatypes.User_Customer_Notification_Virtual_Guest) (datatypes.User_Customer_Notification_Virtual_Guest, error)
	DeleteObjects(templateObjects datatypes.User_Customer_Notification_Virtual_Guest) (bool, error)
	FindByGuestId(id int) (datatypes.User_Customer_Notification_Virtual_Guest, error)
	GetObject() (datatypes.User_Customer_Notification_Virtual_Guest, error)

	GetGuest() (datatypes.Virtual_Guest, error)
	GetUser() (datatypes.User_Customer, error)
}

type User_Customer_OpenIdConnect interface {
	User_Customer

	CompleteInvitationAfterLogin(providerType string, accessToken string, emailRegistrationCode string) error
	CreateOpenIdConnectUserAndCompleteInvitation(providerType string, user datatypes.User_Customer, password string, registrationCode string) (string, error)
	DeclineInvitation(providerType string, registrationCode string) error
	GetDefaultAccount(providerType string) (datatypes.Account, error)
	GetMappedAccounts(providerType string) (datatypes.Account, error)
	GetObject() (datatypes.User_Customer_OpenIdConnect, error)
	GetOpenIdRegistrationInfoFromCode(providerType string, registrationCode string) (datatypes.Account_Authentication_OpenIdConnect_RegistrationInformation, error)
	GetPortalLoginTokenOpenIdConnect(providerType string, accessToken string, accountId int) (datatypes.Container_User_Customer_Portal_Token, error)
	GetUserFromLostPasswordRequest(key string) (datatypes.User_Security_Question, error)
	IsValidPortalPassword(password string) (bool, error)
	LostPassword(username string, email string) (bool, error)
	ResetExpiredPassword(username string, password string, newPassword string, securityQuestionId int, securityQuestionAnswer string) (bool, error)
	SetDefaultAccount(providerType string, accountId int) (datatypes.Account, error)
	SetPasswordFromLostPasswordRequest(key string, password string, securityAnswers datatypes.User_Customer_Security_Answer) (bool, error)
	UpdatePassword(password string) (bool, error)
}

type User_Customer_Prospect interface {
	Entity

	GetAccount() (datatypes.Account, error)
	GetAssignedEmployees() (datatypes.User_Employee, error)
	GetQuotes() (datatypes.Billing_Order_Quote, error)
	GetType() (datatypes.User_Customer_Prospect_Type, error)
}

type User_Customer_Prospect_ServiceProvider_EnrollRequest interface {
	Entity

	Enroll(templateObject datatypes.User_Customer_Prospect_ServiceProvider_EnrollRequest) (datatypes.User_Customer_Prospect_ServiceProvider_EnrollRequest, error)
	GetObject() (datatypes.User_Customer_Prospect_ServiceProvider_EnrollRequest, error)

	GetCompanyType() (datatypes.Catalyst_Company_Type, error)
}

type User_Customer_Prospect_Type interface {
	Entity
}

type User_Customer_Security_Answer interface {
	Entity

	GetObject() (datatypes.User_Customer_Security_Answer, error)

	GetQuestion() (datatypes.User_Security_Question, error)
	GetUser() (datatypes.User_Customer, error)
}

type User_Customer_Status interface {
	Entity

	GetAllObjects() (datatypes.User_Customer_Status, error)
	GetObject() (datatypes.User_Customer_Status, error)
}

type User_Employee interface {
	User_Interface

	GetActions() (datatypes.User_Permission_Action, error)
	GetChatTranscript() (datatypes.Ticket_Chat, error)
	GetEmployeeDepartment() (datatypes.User_Employee_Department, error)
	GetLayoutProfiles() (datatypes.Layout_Profile, error)
	GetMetricTrackingObject() (datatypes.Metric_Tracking_Object, error)
	GetRoles() (datatypes.User_Permission_Role, error)
	GetTicketActivities() (datatypes.Ticket_Activity, error)
	GetTicketAttachmentReferences() (datatypes.Ticket_Attachment, error)
}

type User_Employee_Department interface {
	Entity
}

type User_External_Binding interface {
	Entity

	DeleteObject() (bool, error)
	GetObject() (datatypes.User_External_Binding, error)
	UpdateNote(text string) (bool, error)

	GetAttributes() (datatypes.User_External_Binding_Attribute, error)
	GetBillingItem() (datatypes.Billing_Item, error)
	GetNote() (string, error)
	GetType() (datatypes.User_External_Binding_Type, error)
	GetVendor() (datatypes.User_External_Binding_Vendor, error)
}

type User_External_Binding_Attribute interface {
	Entity

	GetExternalBinding() (datatypes.User_External_Binding, error)
}

type User_External_Binding_Type interface {
	Entity
}

type User_External_Binding_Vendor interface {
	Entity

	GetAllObjects() (datatypes.User_External_Binding_Vendor, error)
	GetObject() (datatypes.User_External_Binding_Vendor, error)
}

type User_Interface interface {
	Entity
}

type User_Permission_Action interface {
	Entity

	GetAllObjects() (datatypes.User_Permission_Action, error)
	GetObject() (datatypes.User_Permission_Action, error)
}

type User_Permission_Group interface {
	Entity

	AddAction(action datatypes.User_Permission_Action) error
	AddBulkActions(actions datatypes.User_Permission_Action) error
	AddBulkResourceObjects(resourceObjects datatypes.Entity, resourceTypeKeyName string) (bool, error)
	AddResourceObject(resourceObject datatypes.Entity, resourceTypeKeyName string) (bool, error)
	CreateObject(templateObject datatypes.User_Permission_Group) (datatypes.User_Permission_Group, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.User_Permission_Group) (datatypes.User_Permission_Group, error)
	GetObject() (datatypes.User_Permission_Group, error)
	LinkRole(role datatypes.User_Permission_Role) error
	RemoveAction(action datatypes.User_Permission_Action) error
	RemoveBulkActions(actions datatypes.User_Permission_Action) error
	RemoveBulkResourceObjects(resourceObjects datatypes.Entity, resourceTypeKeyName string) (bool, error)
	RemoveResourceObject(resourceObject datatypes.Entity, resourceTypeKeyName string) (bool, error)
	UnlinkRole(role datatypes.User_Permission_Role) error

	GetAccount() (datatypes.Account, error)
	GetActions() (datatypes.User_Permission_Action, error)
	GetRoles() (datatypes.User_Permission_Role, error)
	GetType() (datatypes.User_Permission_Group_Type, error)
}

type User_Permission_Group_Type interface {
	Entity

	GetObject() (datatypes.User_Permission_Group_Type, error)

	GetGroups() (datatypes.User_Permission_Group, error)
}

type User_Permission_Role interface {
	Entity

	AddUser(user datatypes.User_Customer) error
	CreateObject(templateObject datatypes.User_Permission_Role) (datatypes.User_Permission_Role, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.User_Permission_Role) (datatypes.User_Permission_Role, error)
	GetObject() (datatypes.User_Permission_Role, error)
	LinkGroup(group datatypes.User_Permission_Group) error
	RemoveUser(user datatypes.User_Customer) error
	UnlinkGroup(group datatypes.User_Permission_Group) error

	GetAccount() (datatypes.Account, error)
	GetActions() (datatypes.User_Permission_Action, error)
	GetGroups() (datatypes.User_Permission_Group, error)
	GetUsers() (datatypes.User_Customer, error)
}

type User_Preference interface {
	Entity

	GetDescription() (string, error)
	GetType() (datatypes.User_Preference_Type, error)
}

type User_Preference_Type interface {
	Entity
}

type User_Security_Question interface {
	Entity

	GetAllObjects() (datatypes.User_Security_Question, error)
	GetObject() (datatypes.User_Security_Question, error)
}
