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

type User_Customer struct {
	Session *Session
	Options
}

func (r *Session) GetUserCustomerService() User_Customer {
	return User_Customer{Session: r}
}

func (r *User_Customer) AcknowledgeSupportPolicy() (err error) {
	var resp datatypes.Void
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) AddApiAuthenticationKey() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) AddBulkHardwareAccess(hardwareIds []int) (resp bool, err error) {
	params := []interface{}{
		hardwareIds,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) AddBulkPortalPermission(permissions []datatypes.User_Customer_CustomerPermission_Permission) (resp bool, err error) {
	params := []interface{}{
		permissions,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) AddBulkRoles(roles []datatypes.User_Permission_Role) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		roles,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) AddBulkVirtualGuestAccess(virtualGuestIds []int) (resp bool, err error) {
	params := []interface{}{
		virtualGuestIds,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) AddExternalBinding(externalBinding *datatypes.User_External_Binding) (resp datatypes.User_Customer_External_Binding, err error) {
	params := []interface{}{
		externalBinding,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) AddHardwareAccess(hardwareId *int) (resp bool, err error) {
	params := []interface{}{
		hardwareId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) AddNotificationSubscriber(notificationKeyName *string) (resp bool, err error) {
	params := []interface{}{
		notificationKeyName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) AddPortalPermission(permission *datatypes.User_Customer_CustomerPermission_Permission) (resp bool, err error) {
	params := []interface{}{
		permission,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) AddRole(role *datatypes.User_Permission_Role) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		role,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) AddVirtualGuestAccess(virtualGuestId *int) (resp bool, err error) {
	params := []interface{}{
		virtualGuestId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) ChangePreference(preferenceTypeKeyName *string, value *string) (resp []datatypes.User_Preference, err error) {
	params := []interface{}{
		preferenceTypeKeyName,
		value,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) CheckExternalAuthenticationStatus(authenticationContainer *datatypes.Container_User_Customer_External_Binding) (resp datatypes.Container_User_Customer_Portal_Token, err error) {
	params := []interface{}{
		authenticationContainer,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) CheckPhoneFactorAuthenticationForPasswordSet(passwordSet *datatypes.Container_User_Customer_PasswordSet, authenticationContainer *datatypes.Container_User_Customer_External_Binding) (resp bool, err error) {
	params := []interface{}{
		passwordSet,
		authenticationContainer,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) CreateNotificationSubscriber(keyName *string, resourceTableId *int) (resp bool, err error) {
	params := []interface{}{
		keyName,
		resourceTableId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) CreateObject(templateObject *datatypes.User_Customer, password *string, vpnPassword *string) (resp datatypes.User_Customer, err error) {
	params := []interface{}{
		templateObject,
		password,
		vpnPassword,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) CreateSubscriberDeliveryMethods(notificationKeyName *string, deliveryMethodKeyNames []string) (resp bool, err error) {
	params := []interface{}{
		notificationKeyName,
		deliveryMethodKeyNames,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) DeactivateNotificationSubscriber(keyName *string, resourceTableId *int) (resp bool, err error) {
	params := []interface{}{
		keyName,
		resourceTableId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) EditObject(templateObject *datatypes.User_Customer) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) EditObjects(templateObjects []datatypes.User_Customer) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) FindUserPreference(profileName *string, containerKeyname *string, preferenceKeyname *string) (resp []datatypes.Layout_Profile, err error) {
	params := []interface{}{
		profileName,
		containerKeyname,
		preferenceKeyname,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetActions() (resp []datatypes.User_Permission_Action, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetActiveExternalAuthenticationVendors() (resp []datatypes.Container_User_Customer_External_Binding_Vendor, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetAdditionalEmails() (resp []datatypes.User_Customer_AdditionalEmail, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetAllowedHardwareIds() (resp []int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetAllowedVirtualGuestIds() (resp []int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetApiAuthenticationKeys() (resp []datatypes.User_Customer_ApiAuthentication, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetAuthenticationToken(token *datatypes.Container_User_Authentication_Token) (resp datatypes.Container_User_Authentication_Token, err error) {
	params := []interface{}{
		token,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetCdnAccounts() (resp []datatypes.Network_ContentDelivery_Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetChildUsers() (resp []datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetClosedTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetDefaultSecurityQuestions(key *string) (resp []datatypes.User_Security_Question, err error) {
	params := []interface{}{
		key,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetExternalBindings() (resp []datatypes.User_External_Binding, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetHardwareCount() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetHardwareNotifications() (resp []datatypes.User_Customer_Notification_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetHasAcknowledgedSupportPolicyFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetHasFullHardwareAccessFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetHasFullVirtualGuestAccessFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetImpersonationToken() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetLayoutProfiles() (resp []datatypes.Layout_Profile, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetLocale() (resp datatypes.Locale, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetLoginAttempts() (resp []datatypes.User_Customer_Access_Authentication, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetMobileDevices() (resp []datatypes.User_Customer_MobileDevice, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetNotificationSubscribers() (resp []datatypes.Notification_Subscriber, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetObject() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetOpenTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetOverrides() (resp []datatypes.Network_Service_Vpn_Overrides, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetParent() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetPermissions() (resp []datatypes.User_Customer_CustomerPermission_Permission, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetPortalLoginToken(username *string, password *string, securityQuestionId *int, securityQuestionAnswer *string) (resp datatypes.Container_User_Customer_Portal_Token, err error) {
	params := []interface{}{
		username,
		password,
		securityQuestionId,
		securityQuestionAnswer,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetPreference(preferenceTypeKeyName *string) (resp datatypes.User_Preference, err error) {
	params := []interface{}{
		preferenceTypeKeyName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetPreferenceTypes() (resp []datatypes.User_Preference_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetPreferences() (resp []datatypes.User_Preference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetRequirementsForPasswordSet(passwordSet *datatypes.Container_User_Customer_PasswordSet) (resp datatypes.Container_User_Customer_PasswordSet, err error) {
	params := []interface{}{
		passwordSet,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetRoles() (resp []datatypes.User_Permission_Role, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetSalesforceUserLink() (resp datatypes.User_Customer_Link, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetSecurityAnswers() (resp []datatypes.User_Customer_Security_Answer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetSubscribers() (resp []datatypes.Notification_User_Subscriber, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetSuccessfulLogins() (resp []datatypes.User_Customer_Access_Authentication, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetSupportPolicyAcknowledgementRequiredFlag() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetSupportPolicyDocument() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetSupportPolicyName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetSupportedLocales() (resp []datatypes.Locale, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetSurveyRequiredFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetSurveys() (resp []datatypes.Survey, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetTimezone() (resp datatypes.Locale_Timezone, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetUnsuccessfulLogins() (resp []datatypes.User_Customer_Access_Authentication, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetUserFromLostPasswordRequest(key *string) (resp []datatypes.User_Security_Question, err error) {
	params := []interface{}{
		key,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetUserIdForPasswordSet(key *string) (resp int, err error) {
	params := []interface{}{
		key,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetUserLinks() (resp []datatypes.User_Customer_Link, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetUserPreferences(profileName *string, containerKeyname *string) (resp []datatypes.Layout_Profile, err error) {
	params := []interface{}{
		profileName,
		containerKeyname,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetUserStatus() (resp datatypes.User_Customer_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetVirtualGuestCount() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) GetVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) InTerminalStatus() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) InitiateExternalAuthentication(authenticationContainer *datatypes.Container_User_Customer_External_Binding) (resp string, err error) {
	params := []interface{}{
		authenticationContainer,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) InitiatePortalPasswordChange(username *string) (resp bool, err error) {
	params := []interface{}{
		username,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) InitiatePortalPasswordChangeByBrandAgent(username *string) (resp bool, err error) {
	params := []interface{}{
		username,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) InviteUserToLinkOpenIdConnect(providerType *string) (resp bool, err error) {
	params := []interface{}{
		providerType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) IsMasterUser() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) IsValidForumPassword(password *string) (resp bool, err error) {
	params := []interface{}{
		password,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) IsValidPortalPassword(password *string) (resp bool, err error) {
	params := []interface{}{
		password,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) LostPassword(username *string, email *string) (resp bool, err error) {
	params := []interface{}{
		username,
		email,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) PerformExternalAuthentication(authenticationContainer *datatypes.Container_User_Customer_External_Binding) (resp datatypes.Container_User_Customer_Portal_Token, err error) {
	params := []interface{}{
		authenticationContainer,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) ProcessPasswordSetRequest(passwordSet *datatypes.Container_User_Customer_PasswordSet, authenticationContainer *datatypes.Container_User_Customer_External_Binding) (resp datatypes.Container_User_Customer_PasswordSet, err error) {
	params := []interface{}{
		passwordSet,
		authenticationContainer,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) RemoveAllHardwareAccessForThisUser() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) RemoveAllVirtualAccessForThisUser() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) RemoveApiAuthenticationKey(keyId *int) (resp bool, err error) {
	params := []interface{}{
		keyId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) RemoveBulkHardwareAccess(hardwareIds []int) (resp bool, err error) {
	params := []interface{}{
		hardwareIds,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) RemoveBulkPortalPermission(permissions []datatypes.User_Customer_CustomerPermission_Permission) (resp bool, err error) {
	params := []interface{}{
		permissions,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) RemoveBulkRoles(roles []datatypes.User_Permission_Role) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		roles,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) RemoveBulkVirtualGuestAccess(virtualGuestIds []int) (resp bool, err error) {
	params := []interface{}{
		virtualGuestIds,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) RemoveExternalBinding(externalBinding *datatypes.User_External_Binding) (resp bool, err error) {
	params := []interface{}{
		externalBinding,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) RemoveHardwareAccess(hardwareId *int) (resp bool, err error) {
	params := []interface{}{
		hardwareId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) RemovePortalPermission(permission *datatypes.User_Customer_CustomerPermission_Permission) (resp bool, err error) {
	params := []interface{}{
		permission,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) RemoveRole(role *datatypes.User_Permission_Role) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		role,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) RemoveVirtualGuestAccess(virtualGuestId *int) (resp bool, err error) {
	params := []interface{}{
		virtualGuestId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) ResetExpiredPassword(username *string, password *string, newPassword *string, securityQuestionId *int, securityQuestionAnswer *string) (resp bool, err error) {
	params := []interface{}{
		username,
		password,
		newPassword,
		securityQuestionId,
		securityQuestionAnswer,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) SamlAuthenticate(accountId *string, samlResponse *string) (resp datatypes.Container_User_Customer_Portal_Token, err error) {
	params := []interface{}{
		accountId,
		samlResponse,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) SamlBeginAuthentication(accountId *int) (resp string, err error) {
	params := []interface{}{
		accountId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) SamlBeginLogout() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) SamlLogout(samlResponse *string) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		samlResponse,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) SetPasswordFromLostPasswordRequest(key *string, password *string, securityAnswers []datatypes.User_Customer_Security_Answer) (resp bool, err error) {
	params := []interface{}{
		key,
		password,
		securityAnswers,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) UpdateForumPassword(password *string) (resp bool, err error) {
	params := []interface{}{
		password,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) UpdateNotificationSubscriber(notificationKeyName *string, active *int) (resp bool, err error) {
	params := []interface{}{
		notificationKeyName,
		active,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) UpdatePassword(password *string) (resp bool, err error) {
	params := []interface{}{
		password,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) UpdateSecurityAnswers(questions []datatypes.User_Security_Question, answers []string) (resp bool, err error) {
	params := []interface{}{
		questions,
		answers,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) UpdateSubscriberDeliveryMethod(notificationKeyName *string, deliveryMethodKeyNames []string, active *int) (resp bool, err error) {
	params := []interface{}{
		notificationKeyName,
		deliveryMethodKeyNames,
		active,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) UpdateVpnPassword(password *string) (resp bool, err error) {
	params := []interface{}{
		password,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) UpdateVpnUser() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer) ValidateAuthenticationToken(authenticationToken *datatypes.Container_User_Authentication_Token) (resp datatypes.Container_User_Customer_Portal_Token, err error) {
	params := []interface{}{
		authenticationToken,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type User_Customer_ApiAuthentication struct {
	Session *Session
	Options
}

func (r *Session) GetUserCustomerApiAuthenticationService() User_Customer_ApiAuthentication {
	return User_Customer_ApiAuthentication{Session: r}
}

func (r *User_Customer_ApiAuthentication) EditObject(templateObject *datatypes.User_Customer_ApiAuthentication) (resp datatypes.User_Customer_ApiAuthentication, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_ApiAuthentication) GetObject() (resp datatypes.User_Customer_ApiAuthentication, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_ApiAuthentication) GetUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type User_Customer_CustomerPermission_Permission struct {
	Session *Session
	Options
}

func (r *Session) GetUserCustomerCustomerPermissionPermissionService() User_Customer_CustomerPermission_Permission {
	return User_Customer_CustomerPermission_Permission{Session: r}
}

func (r *User_Customer_CustomerPermission_Permission) GetAllObjects() (resp []datatypes.User_Customer_CustomerPermission_Permission, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_CustomerPermission_Permission) GetObject() (resp datatypes.User_Customer_CustomerPermission_Permission, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type User_Customer_External_Binding struct {
	Session *Session
	Options
}

func (r *Session) GetUserCustomerExternalBindingService() User_Customer_External_Binding {
	return User_Customer_External_Binding{Session: r}
}

func (r *User_Customer_External_Binding) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding) Disable(reason *string) (resp bool, err error) {
	params := []interface{}{
		reason,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding) Enable() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding) GetAttributes() (resp []datatypes.User_External_Binding_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding) GetNote() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding) GetObject() (resp datatypes.User_Customer_External_Binding, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding) GetType() (resp datatypes.User_External_Binding_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding) GetUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding) GetVendor() (resp datatypes.User_External_Binding_Vendor, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding) UpdateNote(text *string) (resp bool, err error) {
	params := []interface{}{
		text,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type User_Customer_External_Binding_Phone struct {
	Session *Session
	Options
}

func (r *Session) GetUserCustomerExternalBindingPhoneService() User_Customer_External_Binding_Phone {
	return User_Customer_External_Binding_Phone{Session: r}
}

func (r *User_Customer_External_Binding_Phone) CheckPhoneValidationResult(token *string) (resp bool, err error) {
	params := []interface{}{
		token,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Phone) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Phone) Disable(reason *string) (resp bool, err error) {
	params := []interface{}{
		reason,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Phone) Enable() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Phone) GetAllAuthenticationModes() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Phone) GetAllAuthenticationPinModes(authenticationModeKeyName *string) (resp []string, err error) {
	params := []interface{}{
		authenticationModeKeyName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Phone) GetAttributes() (resp []datatypes.User_External_Binding_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Phone) GetAuthenticationMode() (resp datatypes.Container_User_Customer_External_Binding_Phone_Mode, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Phone) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Phone) GetBindingStatus() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Phone) GetNote() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Phone) GetObject() (resp datatypes.User_Customer_External_Binding_Phone, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Phone) GetPhoneAppActivationCode() (resp []datatypes.User_External_Binding_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Phone) GetPhoneData() (resp []datatypes.Container_User_Data_Phone, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Phone) GetPinLength() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Phone) GetType() (resp datatypes.User_External_Binding_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Phone) GetUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Phone) GetVendor() (resp datatypes.User_External_Binding_Vendor, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Phone) RequestPhoneValidation(phoneData *datatypes.Container_User_Data_Phone) (resp string, err error) {
	params := []interface{}{
		phoneData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Phone) UpdateAuthenticationMode(mode *datatypes.Container_User_Customer_External_Binding_Phone_Mode) (resp bool, err error) {
	params := []interface{}{
		mode,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Phone) UpdateNote(text *string) (resp bool, err error) {
	params := []interface{}{
		text,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Phone) UpdatePhone(phoneData []datatypes.Container_User_Data_Phone) (resp bool, err error) {
	params := []interface{}{
		phoneData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type User_Customer_External_Binding_Totp struct {
	Session *Session
	Options
}

func (r *Session) GetUserCustomerExternalBindingTotpService() User_Customer_External_Binding_Totp {
	return User_Customer_External_Binding_Totp{Session: r}
}

func (r *User_Customer_External_Binding_Totp) Activate() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Totp) Deactivate() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Totp) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Totp) Disable(reason *string) (resp bool, err error) {
	params := []interface{}{
		reason,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Totp) Enable() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Totp) GenerateSecretKey() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Totp) GetAttributes() (resp []datatypes.User_External_Binding_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Totp) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Totp) GetNote() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Totp) GetObject() (resp datatypes.User_Customer_External_Binding_Totp, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Totp) GetType() (resp datatypes.User_External_Binding_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Totp) GetUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Totp) GetVendor() (resp datatypes.User_External_Binding_Vendor, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Totp) UpdateNote(text *string) (resp bool, err error) {
	params := []interface{}{
		text,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type User_Customer_External_Binding_Vendor struct {
	Session *Session
	Options
}

func (r *Session) GetUserCustomerExternalBindingVendorService() User_Customer_External_Binding_Vendor {
	return User_Customer_External_Binding_Vendor{Session: r}
}

func (r *User_Customer_External_Binding_Vendor) GetAllObjects() (resp []datatypes.User_External_Binding_Vendor, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Vendor) GetObject() (resp datatypes.User_Customer_External_Binding_Vendor, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type User_Customer_External_Binding_Verisign struct {
	Session *Session
	Options
}

func (r *Session) GetUserCustomerExternalBindingVerisignService() User_Customer_External_Binding_Verisign {
	return User_Customer_External_Binding_Verisign{Session: r}
}

func (r *User_Customer_External_Binding_Verisign) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Verisign) Disable(reason *string) (resp bool, err error) {
	params := []interface{}{
		reason,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Verisign) Enable() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Verisign) GetActivationCodeForMobileClient() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Verisign) GetAttributes() (resp []datatypes.User_External_Binding_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Verisign) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Verisign) GetCredentialExpirationDate() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Verisign) GetCredentialLastUpdateDate() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Verisign) GetCredentialState() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Verisign) GetCredentialType() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Verisign) GetNote() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Verisign) GetObject() (resp datatypes.User_Customer_External_Binding_Verisign, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Verisign) GetType() (resp datatypes.User_External_Binding_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Verisign) GetUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Verisign) GetVendor() (resp datatypes.User_External_Binding_Vendor, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Verisign) Unlock(securityCode *string) (resp bool, err error) {
	params := []interface{}{
		securityCode,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Verisign) UpdateNote(text *string) (resp bool, err error) {
	params := []interface{}{
		text,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_External_Binding_Verisign) ValidateCredentialId(userId *int, externalId *string) (resp bool, err error) {
	params := []interface{}{
		userId,
		externalId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type User_Customer_Invitation struct {
	Session *Session
	Options
}

func (r *Session) GetUserCustomerInvitationService() User_Customer_Invitation {
	return User_Customer_Invitation{Session: r}
}

func (r *User_Customer_Invitation) GetObject() (resp datatypes.User_Customer_Invitation, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_Invitation) GetUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type User_Customer_MobileDevice struct {
	Session *Session
	Options
}

func (r *Session) GetUserCustomerMobileDeviceService() User_Customer_MobileDevice {
	return User_Customer_MobileDevice{Session: r}
}

func (r *User_Customer_MobileDevice) CreateObject(templateObject *datatypes.User_Customer_MobileDevice) (resp datatypes.User_Customer_MobileDevice, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_MobileDevice) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_MobileDevice) EditObject(templateObject *datatypes.User_Customer_MobileDevice) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_MobileDevice) GetAvailablePushNotificationSubscriptions() (resp []datatypes.Notification, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_MobileDevice) GetCustomer() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_MobileDevice) GetObject() (resp datatypes.User_Customer_MobileDevice, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_MobileDevice) GetOperatingSystem() (resp datatypes.User_Customer_MobileDevice_OperatingSystem, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_MobileDevice) GetPushNotificationSubscriptions() (resp []datatypes.Notification_User_Subscriber, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_MobileDevice) GetType() (resp datatypes.User_Customer_MobileDevice_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type User_Customer_MobileDevice_OperatingSystem struct {
	Session *Session
	Options
}

func (r *Session) GetUserCustomerMobileDeviceOperatingSystemService() User_Customer_MobileDevice_OperatingSystem {
	return User_Customer_MobileDevice_OperatingSystem{Session: r}
}

func (r *User_Customer_MobileDevice_OperatingSystem) GetAllObjects() (resp []datatypes.User_Customer_MobileDevice_OperatingSystem, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_MobileDevice_OperatingSystem) GetObject() (resp datatypes.User_Customer_MobileDevice_OperatingSystem, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type User_Customer_MobileDevice_Type struct {
	Session *Session
	Options
}

func (r *Session) GetUserCustomerMobileDeviceTypeService() User_Customer_MobileDevice_Type {
	return User_Customer_MobileDevice_Type{Session: r}
}

func (r *User_Customer_MobileDevice_Type) GetAllObjects() (resp []datatypes.User_Customer_MobileDevice_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_MobileDevice_Type) GetObject() (resp datatypes.User_Customer_MobileDevice_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type User_Customer_Notification_Hardware struct {
	Session *Session
	Options
}

func (r *Session) GetUserCustomerNotificationHardwareService() User_Customer_Notification_Hardware {
	return User_Customer_Notification_Hardware{Session: r}
}

func (r *User_Customer_Notification_Hardware) CreateObject(templateObject *datatypes.User_Customer_Notification_Hardware) (resp datatypes.User_Customer_Notification_Hardware, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_Notification_Hardware) CreateObjects(templateObjects []datatypes.User_Customer_Notification_Hardware) (resp []datatypes.Dns_Domain, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_Notification_Hardware) DeleteObjects(templateObjects []datatypes.User_Customer_Notification_Hardware) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_Notification_Hardware) FindByHardwareId(hardwareId *int) (resp []datatypes.User_Customer_Notification_Hardware, err error) {
	params := []interface{}{
		hardwareId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_Notification_Hardware) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_Notification_Hardware) GetObject() (resp datatypes.User_Customer_Notification_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_Notification_Hardware) GetUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type User_Customer_Notification_Virtual_Guest struct {
	Session *Session
	Options
}

func (r *Session) GetUserCustomerNotificationVirtualGuestService() User_Customer_Notification_Virtual_Guest {
	return User_Customer_Notification_Virtual_Guest{Session: r}
}

func (r *User_Customer_Notification_Virtual_Guest) CreateObject(templateObject *datatypes.User_Customer_Notification_Virtual_Guest) (resp datatypes.User_Customer_Notification_Virtual_Guest, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_Notification_Virtual_Guest) CreateObjects(templateObjects []datatypes.User_Customer_Notification_Virtual_Guest) (resp []datatypes.User_Customer_Notification_Virtual_Guest, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_Notification_Virtual_Guest) DeleteObjects(templateObjects []datatypes.User_Customer_Notification_Virtual_Guest) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_Notification_Virtual_Guest) FindByGuestId(id *int) (resp []datatypes.User_Customer_Notification_Virtual_Guest, err error) {
	params := []interface{}{
		id,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_Notification_Virtual_Guest) GetGuest() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_Notification_Virtual_Guest) GetObject() (resp datatypes.User_Customer_Notification_Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_Notification_Virtual_Guest) GetUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type User_Customer_OpenIdConnect struct {
	Session *Session
	Options
}

func (r *Session) GetUserCustomerOpenIdConnectService() User_Customer_OpenIdConnect {
	return User_Customer_OpenIdConnect{Session: r}
}

func (r *User_Customer_OpenIdConnect) AcknowledgeSupportPolicy() (err error) {
	var resp datatypes.Void
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) AddApiAuthenticationKey() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) AddBulkHardwareAccess(hardwareIds []int) (resp bool, err error) {
	params := []interface{}{
		hardwareIds,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) AddBulkPortalPermission(permissions []datatypes.User_Customer_CustomerPermission_Permission) (resp bool, err error) {
	params := []interface{}{
		permissions,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) AddBulkRoles(roles []datatypes.User_Permission_Role) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		roles,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) AddBulkVirtualGuestAccess(virtualGuestIds []int) (resp bool, err error) {
	params := []interface{}{
		virtualGuestIds,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) AddExternalBinding(externalBinding *datatypes.User_External_Binding) (resp datatypes.User_Customer_External_Binding, err error) {
	params := []interface{}{
		externalBinding,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) AddHardwareAccess(hardwareId *int) (resp bool, err error) {
	params := []interface{}{
		hardwareId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) AddNotificationSubscriber(notificationKeyName *string) (resp bool, err error) {
	params := []interface{}{
		notificationKeyName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) AddPortalPermission(permission *datatypes.User_Customer_CustomerPermission_Permission) (resp bool, err error) {
	params := []interface{}{
		permission,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) AddRole(role *datatypes.User_Permission_Role) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		role,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) AddVirtualGuestAccess(virtualGuestId *int) (resp bool, err error) {
	params := []interface{}{
		virtualGuestId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) ChangePreference(preferenceTypeKeyName *string, value *string) (resp []datatypes.User_Preference, err error) {
	params := []interface{}{
		preferenceTypeKeyName,
		value,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) CheckExternalAuthenticationStatus(authenticationContainer *datatypes.Container_User_Customer_External_Binding) (resp datatypes.Container_User_Customer_Portal_Token, err error) {
	params := []interface{}{
		authenticationContainer,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) CheckPhoneFactorAuthenticationForPasswordSet(passwordSet *datatypes.Container_User_Customer_PasswordSet, authenticationContainer *datatypes.Container_User_Customer_External_Binding) (resp bool, err error) {
	params := []interface{}{
		passwordSet,
		authenticationContainer,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) CompleteInvitationAfterLogin(providerType *string, accessToken *string, emailRegistrationCode *string) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		providerType,
		accessToken,
		emailRegistrationCode,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) CreateNotificationSubscriber(keyName *string, resourceTableId *int) (resp bool, err error) {
	params := []interface{}{
		keyName,
		resourceTableId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) CreateObject(templateObject *datatypes.User_Customer, password *string, vpnPassword *string) (resp datatypes.User_Customer, err error) {
	params := []interface{}{
		templateObject,
		password,
		vpnPassword,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) CreateOpenIdConnectUserAndCompleteInvitation(providerType *string, user *datatypes.User_Customer, password *string, registrationCode *string) (resp string, err error) {
	params := []interface{}{
		providerType,
		user,
		password,
		registrationCode,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) CreateSubscriberDeliveryMethods(notificationKeyName *string, deliveryMethodKeyNames []string) (resp bool, err error) {
	params := []interface{}{
		notificationKeyName,
		deliveryMethodKeyNames,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) DeactivateNotificationSubscriber(keyName *string, resourceTableId *int) (resp bool, err error) {
	params := []interface{}{
		keyName,
		resourceTableId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) DeclineInvitation(providerType *string, registrationCode *string) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		providerType,
		registrationCode,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) EditObject(templateObject *datatypes.User_Customer) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) EditObjects(templateObjects []datatypes.User_Customer) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) FindUserPreference(profileName *string, containerKeyname *string, preferenceKeyname *string) (resp []datatypes.Layout_Profile, err error) {
	params := []interface{}{
		profileName,
		containerKeyname,
		preferenceKeyname,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetActions() (resp []datatypes.User_Permission_Action, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetActiveExternalAuthenticationVendors() (resp []datatypes.Container_User_Customer_External_Binding_Vendor, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetAdditionalEmails() (resp []datatypes.User_Customer_AdditionalEmail, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetAllowedHardwareIds() (resp []int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetAllowedVirtualGuestIds() (resp []int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetApiAuthenticationKeys() (resp []datatypes.User_Customer_ApiAuthentication, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetAuthenticationToken(token *datatypes.Container_User_Authentication_Token) (resp datatypes.Container_User_Authentication_Token, err error) {
	params := []interface{}{
		token,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetCdnAccounts() (resp []datatypes.Network_ContentDelivery_Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetChildUsers() (resp []datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetClosedTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetDefaultAccount(providerType *string) (resp datatypes.Account, err error) {
	params := []interface{}{
		providerType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetDefaultSecurityQuestions(key *string) (resp []datatypes.User_Security_Question, err error) {
	params := []interface{}{
		key,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetExternalBindings() (resp []datatypes.User_External_Binding, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetHardwareCount() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetHardwareNotifications() (resp []datatypes.User_Customer_Notification_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetHasAcknowledgedSupportPolicyFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetHasFullHardwareAccessFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetHasFullVirtualGuestAccessFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetImpersonationToken() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetLayoutProfiles() (resp []datatypes.Layout_Profile, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetLocale() (resp datatypes.Locale, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetLoginAttempts() (resp []datatypes.User_Customer_Access_Authentication, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetMappedAccounts(providerType *string) (resp []datatypes.Account, err error) {
	params := []interface{}{
		providerType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetMobileDevices() (resp []datatypes.User_Customer_MobileDevice, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetNotificationSubscribers() (resp []datatypes.Notification_Subscriber, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetObject() (resp datatypes.User_Customer_OpenIdConnect, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetOpenIdRegistrationInfoFromCode(providerType *string, registrationCode *string) (resp datatypes.Account_Authentication_OpenIdConnect_RegistrationInformation, err error) {
	params := []interface{}{
		providerType,
		registrationCode,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetOpenTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetOverrides() (resp []datatypes.Network_Service_Vpn_Overrides, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetParent() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetPermissions() (resp []datatypes.User_Customer_CustomerPermission_Permission, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetPortalLoginToken(username *string, password *string, securityQuestionId *int, securityQuestionAnswer *string) (resp datatypes.Container_User_Customer_Portal_Token, err error) {
	params := []interface{}{
		username,
		password,
		securityQuestionId,
		securityQuestionAnswer,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetPortalLoginTokenOpenIdConnect(providerType *string, accessToken *string, accountId *int) (resp datatypes.Container_User_Customer_Portal_Token, err error) {
	params := []interface{}{
		providerType,
		accessToken,
		accountId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetPreference(preferenceTypeKeyName *string) (resp datatypes.User_Preference, err error) {
	params := []interface{}{
		preferenceTypeKeyName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetPreferenceTypes() (resp []datatypes.User_Preference_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetPreferences() (resp []datatypes.User_Preference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetRequirementsForPasswordSet(passwordSet *datatypes.Container_User_Customer_PasswordSet) (resp datatypes.Container_User_Customer_PasswordSet, err error) {
	params := []interface{}{
		passwordSet,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetRoles() (resp []datatypes.User_Permission_Role, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetSalesforceUserLink() (resp datatypes.User_Customer_Link, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetSecurityAnswers() (resp []datatypes.User_Customer_Security_Answer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetSubscribers() (resp []datatypes.Notification_User_Subscriber, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetSuccessfulLogins() (resp []datatypes.User_Customer_Access_Authentication, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetSupportPolicyAcknowledgementRequiredFlag() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetSupportPolicyDocument() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetSupportPolicyName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetSupportedLocales() (resp []datatypes.Locale, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetSurveyRequiredFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetSurveys() (resp []datatypes.Survey, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetTimezone() (resp datatypes.Locale_Timezone, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetUnsuccessfulLogins() (resp []datatypes.User_Customer_Access_Authentication, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetUserFromLostPasswordRequest(key *string) (resp []datatypes.User_Security_Question, err error) {
	params := []interface{}{
		key,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetUserIdForPasswordSet(key *string) (resp int, err error) {
	params := []interface{}{
		key,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetUserLinks() (resp []datatypes.User_Customer_Link, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetUserPreferences(profileName *string, containerKeyname *string) (resp []datatypes.Layout_Profile, err error) {
	params := []interface{}{
		profileName,
		containerKeyname,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetUserStatus() (resp datatypes.User_Customer_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetVirtualGuestCount() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) GetVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) InTerminalStatus() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) InitiateExternalAuthentication(authenticationContainer *datatypes.Container_User_Customer_External_Binding) (resp string, err error) {
	params := []interface{}{
		authenticationContainer,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) InitiatePortalPasswordChange(username *string) (resp bool, err error) {
	params := []interface{}{
		username,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) InitiatePortalPasswordChangeByBrandAgent(username *string) (resp bool, err error) {
	params := []interface{}{
		username,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) InviteUserToLinkOpenIdConnect(providerType *string) (resp bool, err error) {
	params := []interface{}{
		providerType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) IsMasterUser() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) IsValidForumPassword(password *string) (resp bool, err error) {
	params := []interface{}{
		password,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) IsValidPortalPassword(password *string) (resp bool, err error) {
	params := []interface{}{
		password,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) LostPassword(username *string, email *string) (resp bool, err error) {
	params := []interface{}{
		username,
		email,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) PerformExternalAuthentication(authenticationContainer *datatypes.Container_User_Customer_External_Binding) (resp datatypes.Container_User_Customer_Portal_Token, err error) {
	params := []interface{}{
		authenticationContainer,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) ProcessPasswordSetRequest(passwordSet *datatypes.Container_User_Customer_PasswordSet, authenticationContainer *datatypes.Container_User_Customer_External_Binding) (resp datatypes.Container_User_Customer_PasswordSet, err error) {
	params := []interface{}{
		passwordSet,
		authenticationContainer,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) RemoveAllHardwareAccessForThisUser() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) RemoveAllVirtualAccessForThisUser() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) RemoveApiAuthenticationKey(keyId *int) (resp bool, err error) {
	params := []interface{}{
		keyId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) RemoveBulkHardwareAccess(hardwareIds []int) (resp bool, err error) {
	params := []interface{}{
		hardwareIds,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) RemoveBulkPortalPermission(permissions []datatypes.User_Customer_CustomerPermission_Permission) (resp bool, err error) {
	params := []interface{}{
		permissions,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) RemoveBulkRoles(roles []datatypes.User_Permission_Role) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		roles,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) RemoveBulkVirtualGuestAccess(virtualGuestIds []int) (resp bool, err error) {
	params := []interface{}{
		virtualGuestIds,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) RemoveExternalBinding(externalBinding *datatypes.User_External_Binding) (resp bool, err error) {
	params := []interface{}{
		externalBinding,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) RemoveHardwareAccess(hardwareId *int) (resp bool, err error) {
	params := []interface{}{
		hardwareId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) RemovePortalPermission(permission *datatypes.User_Customer_CustomerPermission_Permission) (resp bool, err error) {
	params := []interface{}{
		permission,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) RemoveRole(role *datatypes.User_Permission_Role) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		role,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) RemoveVirtualGuestAccess(virtualGuestId *int) (resp bool, err error) {
	params := []interface{}{
		virtualGuestId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) ResetExpiredPassword(username *string, password *string, newPassword *string, securityQuestionId *int, securityQuestionAnswer *string) (resp bool, err error) {
	params := []interface{}{
		username,
		password,
		newPassword,
		securityQuestionId,
		securityQuestionAnswer,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) SamlAuthenticate(accountId *string, samlResponse *string) (resp datatypes.Container_User_Customer_Portal_Token, err error) {
	params := []interface{}{
		accountId,
		samlResponse,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) SamlBeginAuthentication(accountId *int) (resp string, err error) {
	params := []interface{}{
		accountId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) SamlBeginLogout() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) SamlLogout(samlResponse *string) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		samlResponse,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) SetDefaultAccount(providerType *string, accountId *int) (resp datatypes.Account, err error) {
	params := []interface{}{
		providerType,
		accountId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) SetPasswordFromLostPasswordRequest(key *string, password *string, securityAnswers []datatypes.User_Customer_Security_Answer) (resp bool, err error) {
	params := []interface{}{
		key,
		password,
		securityAnswers,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) UpdateForumPassword(password *string) (resp bool, err error) {
	params := []interface{}{
		password,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) UpdateNotificationSubscriber(notificationKeyName *string, active *int) (resp bool, err error) {
	params := []interface{}{
		notificationKeyName,
		active,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) UpdatePassword(password *string) (resp bool, err error) {
	params := []interface{}{
		password,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) UpdateSecurityAnswers(questions []datatypes.User_Security_Question, answers []string) (resp bool, err error) {
	params := []interface{}{
		questions,
		answers,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) UpdateSubscriberDeliveryMethod(notificationKeyName *string, deliveryMethodKeyNames []string, active *int) (resp bool, err error) {
	params := []interface{}{
		notificationKeyName,
		deliveryMethodKeyNames,
		active,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) UpdateVpnPassword(password *string) (resp bool, err error) {
	params := []interface{}{
		password,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) UpdateVpnUser() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_OpenIdConnect) ValidateAuthenticationToken(authenticationToken *datatypes.Container_User_Authentication_Token) (resp datatypes.Container_User_Customer_Portal_Token, err error) {
	params := []interface{}{
		authenticationToken,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type User_Customer_Prospect_ServiceProvider_EnrollRequest struct {
	Session *Session
	Options
}

func (r *Session) GetUserCustomerProspectServiceProviderEnrollRequestService() User_Customer_Prospect_ServiceProvider_EnrollRequest {
	return User_Customer_Prospect_ServiceProvider_EnrollRequest{Session: r}
}

func (r *User_Customer_Prospect_ServiceProvider_EnrollRequest) Enroll(templateObject *datatypes.User_Customer_Prospect_ServiceProvider_EnrollRequest) (resp datatypes.User_Customer_Prospect_ServiceProvider_EnrollRequest, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_Prospect_ServiceProvider_EnrollRequest) GetCompanyType() (resp datatypes.Catalyst_Company_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_Prospect_ServiceProvider_EnrollRequest) GetObject() (resp datatypes.User_Customer_Prospect_ServiceProvider_EnrollRequest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type User_Customer_Security_Answer struct {
	Session *Session
	Options
}

func (r *Session) GetUserCustomerSecurityAnswerService() User_Customer_Security_Answer {
	return User_Customer_Security_Answer{Session: r}
}

func (r *User_Customer_Security_Answer) GetObject() (resp datatypes.User_Customer_Security_Answer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_Security_Answer) GetQuestion() (resp datatypes.User_Security_Question, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_Security_Answer) GetUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type User_Customer_Status struct {
	Session *Session
	Options
}

func (r *Session) GetUserCustomerStatusService() User_Customer_Status {
	return User_Customer_Status{Session: r}
}

func (r *User_Customer_Status) GetAllObjects() (resp []datatypes.User_Customer_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Customer_Status) GetObject() (resp datatypes.User_Customer_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type User_External_Binding struct {
	Session *Session
	Options
}

func (r *Session) GetUserExternalBindingService() User_External_Binding {
	return User_External_Binding{Session: r}
}

func (r *User_External_Binding) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_External_Binding) GetAttributes() (resp []datatypes.User_External_Binding_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_External_Binding) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_External_Binding) GetNote() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_External_Binding) GetObject() (resp datatypes.User_External_Binding, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_External_Binding) GetType() (resp datatypes.User_External_Binding_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_External_Binding) GetVendor() (resp datatypes.User_External_Binding_Vendor, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_External_Binding) UpdateNote(text *string) (resp bool, err error) {
	params := []interface{}{
		text,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type User_External_Binding_Vendor struct {
	Session *Session
	Options
}

func (r *Session) GetUserExternalBindingVendorService() User_External_Binding_Vendor {
	return User_External_Binding_Vendor{Session: r}
}

func (r *User_External_Binding_Vendor) GetAllObjects() (resp []datatypes.User_External_Binding_Vendor, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_External_Binding_Vendor) GetObject() (resp datatypes.User_External_Binding_Vendor, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type User_Permission_Action struct {
	Session *Session
	Options
}

func (r *Session) GetUserPermissionActionService() User_Permission_Action {
	return User_Permission_Action{Session: r}
}

func (r *User_Permission_Action) GetAllObjects() (resp []datatypes.User_Permission_Action, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Action) GetObject() (resp datatypes.User_Permission_Action, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type User_Permission_Group struct {
	Session *Session
	Options
}

func (r *Session) GetUserPermissionGroupService() User_Permission_Group {
	return User_Permission_Group{Session: r}
}

func (r *User_Permission_Group) AddAction(action *datatypes.User_Permission_Action) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		action,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Group) AddBulkActions(actions []datatypes.User_Permission_Action) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		actions,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Group) AddBulkResourceObjects(resourceObjects []datatypes.Entity, resourceTypeKeyName *string) (resp bool, err error) {
	params := []interface{}{
		resourceObjects,
		resourceTypeKeyName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Group) AddResourceObject(resourceObject *datatypes.Entity, resourceTypeKeyName *string) (resp bool, err error) {
	params := []interface{}{
		resourceObject,
		resourceTypeKeyName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Group) CreateObject(templateObject *datatypes.User_Permission_Group) (resp datatypes.User_Permission_Group, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Group) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Group) EditObject(templateObject *datatypes.User_Permission_Group) (resp datatypes.User_Permission_Group, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Group) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Group) GetActions() (resp []datatypes.User_Permission_Action, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Group) GetObject() (resp datatypes.User_Permission_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Group) GetRoles() (resp []datatypes.User_Permission_Role, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Group) GetType() (resp datatypes.User_Permission_Group_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Group) LinkRole(role *datatypes.User_Permission_Role) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		role,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Group) RemoveAction(action *datatypes.User_Permission_Action) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		action,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Group) RemoveBulkActions(actions []datatypes.User_Permission_Action) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		actions,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Group) RemoveBulkResourceObjects(resourceObjects []datatypes.Entity, resourceTypeKeyName *string) (resp bool, err error) {
	params := []interface{}{
		resourceObjects,
		resourceTypeKeyName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Group) RemoveResourceObject(resourceObject *datatypes.Entity, resourceTypeKeyName *string) (resp bool, err error) {
	params := []interface{}{
		resourceObject,
		resourceTypeKeyName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Group) UnlinkRole(role *datatypes.User_Permission_Role) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		role,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type User_Permission_Group_Type struct {
	Session *Session
	Options
}

func (r *Session) GetUserPermissionGroupTypeService() User_Permission_Group_Type {
	return User_Permission_Group_Type{Session: r}
}

func (r *User_Permission_Group_Type) GetGroups() (resp []datatypes.User_Permission_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Group_Type) GetObject() (resp datatypes.User_Permission_Group_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type User_Permission_Role struct {
	Session *Session
	Options
}

func (r *Session) GetUserPermissionRoleService() User_Permission_Role {
	return User_Permission_Role{Session: r}
}

func (r *User_Permission_Role) AddUser(user *datatypes.User_Customer) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		user,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Role) CreateObject(templateObject *datatypes.User_Permission_Role) (resp datatypes.User_Permission_Role, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Role) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Role) EditObject(templateObject *datatypes.User_Permission_Role) (resp datatypes.User_Permission_Role, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Role) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Role) GetActions() (resp []datatypes.User_Permission_Action, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Role) GetGroups() (resp []datatypes.User_Permission_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Role) GetObject() (resp datatypes.User_Permission_Role, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Role) GetUsers() (resp []datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Role) LinkGroup(group *datatypes.User_Permission_Group) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		group,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Role) RemoveUser(user *datatypes.User_Customer) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		user,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *User_Permission_Role) UnlinkGroup(group *datatypes.User_Permission_Group) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		group,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type User_Security_Question struct {
	Session *Session
	Options
}

func (r *Session) GetUserSecurityQuestionService() User_Security_Question {
	return User_Security_Question{Session: r}
}

func (r *User_Security_Question) GetAllObjects() (resp []datatypes.User_Security_Question, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *User_Security_Question) GetObject() (resp datatypes.User_Security_Question, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
