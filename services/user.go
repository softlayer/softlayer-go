/**
 * Copyright 2016-2024 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
 * the License. You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed
 * on an "AS IS" BASIS,WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and limitations under the License.
 */

// AUTOMATICALLY GENERATED CODE - DO NOT MODIFY

package services

import (
	"fmt"
	"strings"

	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/sl"
)

// The SoftLayer_User_Customer data type contains general information relating to a single SoftLayer customer portal user. Personal information in this type such as names, addresses, and phone numbers are not necessarily associated with the customer account the user is assigned to.
type User_Customer struct {
	Session session.SLSession
	Options sl.Options
}

// GetUserCustomerService returns an instance of the User_Customer SoftLayer service
func GetUserCustomerService(sess session.SLSession) User_Customer {
	return User_Customer{Session: sess}
}

func (r User_Customer) Id(id int) User_Customer {
	r.Options.Id = &id
	return r
}

func (r User_Customer) Mask(mask string) User_Customer {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r User_Customer) Filter(filter string) User_Customer {
	r.Options.Filter = filter
	return r
}

func (r User_Customer) Limit(limit int) User_Customer {
	r.Options.Limit = &limit
	return r
}

func (r User_Customer) Offset(offset int) User_Customer {
	r.Options.Offset = &offset
	return r
}

// Create a user's API authentication key, allowing that user access to query the SoftLayer API. addApiAuthenticationKey() returns the user's new API key. Each portal user is allowed only one API key.
func (r User_Customer) AddApiAuthenticationKey() (resp string, err error) {
	err = r.Session.DoRequest("SoftLayer_User_Customer", "addApiAuthenticationKey", nil, &r.Options, &resp)
	return
}

// Add multiple permissions to a portal user's permission set. [[SoftLayer_User_Customer_CustomerPermission_Permission]] control which features in the SoftLayer customer portal and API a user may use. addBulkPortalPermission() does not attempt to add permissions already assigned to the user.
//
// Users can assign permissions to their child users, but not to themselves. An account's master has all portal permissions and can set permissions for any of the other users on their account.
//
// Use the [[SoftLayer_User_Customer_CustomerPermission_Permission::getAllObjects]] method to retrieve a list of all permissions available in the SoftLayer customer portal and API. Permissions are removed based on the keyName property of the permission objects within the permissions parameter.
func (r User_Customer) AddBulkPortalPermission(permissions []datatypes.User_Customer_CustomerPermission_Permission) (resp bool, err error) {
	params := []interface{}{
		permissions,
	}
	err = r.Session.DoRequest("SoftLayer_User_Customer", "addBulkPortalPermission", params, &r.Options, &resp)
	return
}

// Grants the user access to a single dedicated host device.  The user will only be allowed to see and access devices in both the portal and the API to which they have been granted access.  If the user's account has devices to which the user has not been granted access, then "not found" exceptions are thrown if the user attempts to access any of these devices.
//
// Users can assign device access to their child users, but not to themselves. An account's master has access to all devices on their customer account and can set dedicated host access for any of the other users on their account.
//
// Only the USER_MANAGE permission is required to execute this.
func (r User_Customer) AddDedicatedHostAccess(dedicatedHostId *int) (resp bool, err error) {
	params := []interface{}{
		dedicatedHostId,
	}
	err = r.Session.DoRequest("SoftLayer_User_Customer", "addDedicatedHostAccess", params, &r.Options, &resp)
	return
}

// Add hardware to a portal user's hardware access list. A user's hardware access list controls which of an account's hardware objects a user has access to in the SoftLayer customer portal and API. Hardware does not exist in the SoftLayer portal and returns "not found" exceptions in the API if the user doesn't have access to it. If a user already has access to the hardware you're attempting to add then addHardwareAccess() returns true.
//
// Users can assign hardware access to their child users, but not to themselves. An account's master has access to all hardware on their customer account and can set hardware access for any of the other users on their account.
//
// Only the USER_MANAGE permission is required to execute this.
func (r User_Customer) AddHardwareAccess(hardwareId *int) (resp bool, err error) {
	params := []interface{}{
		hardwareId,
	}
	err = r.Session.DoRequest("SoftLayer_User_Customer", "addHardwareAccess", params, &r.Options, &resp)
	return
}

// Add a CloudLayer Computing Instance to a portal user's access list. A user's CloudLayer Computing Instance access list controls which of an account's CloudLayer Computing Instance objects a user has access to in the SoftLayer customer portal and API. CloudLayer Computing Instances do not exist in the SoftLayer portal and returns "not found" exceptions in the API if the user doesn't have access to it. If a user already has access to the CloudLayer Computing Instance you're attempting to add then addVirtualGuestAccess() returns true.
//
// Users can assign CloudLayer Computing Instance access to their child users, but not to themselves. An account's master has access to all CloudLayer Computing Instances on their customer account and can set CloudLayer Computing Instance access for any of the other users on their account.
//
// Only the USER_MANAGE permission is required to execute this.
func (r User_Customer) AddVirtualGuestAccess(virtualGuestId *int) (resp bool, err error) {
	params := []interface{}{
		virtualGuestId,
	}
	err = r.Session.DoRequest("SoftLayer_User_Customer", "addVirtualGuestAccess", params, &r.Options, &resp)
	return
}

// Create a new user in the SoftLayer customer portal. It is not possible to set up SLL enable flags during object creation. These flags are ignored during object creation. You will need to make a subsequent call to edit object in order to enable VPN access.
//
// An account's master user and sub-users who have the User Manage permission can add new users.
//
// Users are created with a default permission set. After adding a user it may be helpful to set their permissions and device access.
//
// secondaryPasswordTimeoutDays will be set to the system configured default value if the attribute is not provided or the attribute is not a valid value.
//
// Note, neither password nor vpnPassword parameters are required.
//
// Password When a new user is created, an email will be sent to the new user's email address with a link to a url that will allow the new user to create or change their password for the SoftLayer customer portal.
//
// If the password parameter is provided and is not null, then that value will be validated. If it is a valid password, then the user will be created with this password.  This user will still receive a portal password email.  It can be used within 24 hours to change their password, or it can be allowed to expire, and the password provided during user creation will remain as the user's password.
//
// If the password parameter is not provided or the value is null, the user must set their portal password using the link sent in email within 24 hours.  If the user fails to set their password within 24 hours, then a non-master user can use the "Reset Password" link on the login page of the portal to request a new email.  A master user can use the link to retrieve a phone number to call to assist in resetting their password.
//
// The password parameter is ignored for VPN_ONLY users or for IBMid authenticated users.
//
// vpnPassword If the vpnPassword is provided, then the user's vpnPassword will be set to the provided password.  When creating a vpn only user, the vpnPassword MUST be supplied.  If the vpnPassword is not provided, then the user will need to use the portal to edit their profile and set the vpnPassword.
//
// IBMid considerations When a SoftLayer account is linked to a Platform Services (PaaS, formerly Bluemix) account, AND the trait on the SoftLayer Account indicating IBMid authentication is set, then SoftLayer will delegate the creation of an ACTIVE user to PaaS. This means that even though the request to create a new user in such an account may start at the IMS API, via this delegation we effectively turn it into a request that is driven by PaaS. In particular this means that any "invitation email" that comes to the user, will come from PaaS, not from IMS via IBMid.
//
// Users created in states other than ACTIVE (for example, a VPN_ONLY user) will be created directly in IMS without delegation (but note that no invitation is sent for a user created in any state other than ACTIVE).
func (r User_Customer) CreateObject(templateObject *datatypes.User_Customer, password *string, vpnPassword *string) (resp datatypes.User_Customer, err error) {
	params := []interface{}{
		templateObject,
		password,
		vpnPassword,
	}
	err = r.Session.DoRequest("SoftLayer_User_Customer", "createObject", params, &r.Options, &resp)
	return
}

// Account master users and sub-users who have the User Manage permission in the SoftLayer customer portal can update other user's information. Use editObject() if you wish to edit a single user account. Users who do not have the User Manage permission can only update their own information.
func (r User_Customer) EditObject(templateObject *datatypes.User_Customer) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_User_Customer", "editObject", params, &r.Options, &resp)
	return
}

// Retrieve A portal user's API Authentication keys. There is a max limit of one API key per user.
func (r User_Customer) GetApiAuthenticationKeys() (resp []datatypes.User_Customer_ApiAuthentication, err error) {
	err = r.Session.DoRequest("SoftLayer_User_Customer", "getApiAuthenticationKeys", nil, &r.Options, &resp)
	return
}

// Retrieve The dedicated hosts to which the user has been granted access.
func (r User_Customer) GetDedicatedHosts() (resp []datatypes.Virtual_DedicatedHost, err error) {
	err = r.Session.DoRequest("SoftLayer_User_Customer", "getDedicatedHosts", nil, &r.Options, &resp)
	return
}

// Retrieve A portal user's accessible hardware. These permissions control which hardware a user has access to in the SoftLayer customer portal.
func (r User_Customer) GetHardware() (resp []datatypes.Hardware, err error) {
	err = r.Session.DoRequest("SoftLayer_User_Customer", "getHardware", nil, &r.Options, &resp)
	return
}

// Retrieve A user's attempts to log into the SoftLayer customer portal.
func (r User_Customer) GetLoginAttempts() (resp []datatypes.User_Customer_Access_Authentication, err error) {
	err = r.Session.DoRequest("SoftLayer_User_Customer", "getLoginAttempts", nil, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_User_Customer object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_User_Customer service. You can only retrieve users that are assigned to the customer account belonging to the user making the API call.
func (r User_Customer) GetObject() (resp datatypes.User_Customer, err error) {
	err = r.Session.DoRequest("SoftLayer_User_Customer", "getObject", nil, &r.Options, &resp)
	return
}

// Retrieve A portal user's vpn accessible subnets.
func (r User_Customer) GetOverrides() (resp []datatypes.Network_Service_Vpn_Overrides, err error) {
	err = r.Session.DoRequest("SoftLayer_User_Customer", "getOverrides", nil, &r.Options, &resp)
	return
}

// Retrieve A portal user's permissions. These permissions control that user's access to functions within the SoftLayer customer portal and API.
func (r User_Customer) GetPermissions() (resp []datatypes.User_Customer_CustomerPermission_Permission, err error) {
	err = r.Session.DoRequest("SoftLayer_User_Customer", "getPermissions", nil, &r.Options, &resp)
	return
}

// Retrieve A portal user's accessible CloudLayer Computing Instances. These permissions control which CloudLayer Computing Instances a user has access to in the SoftLayer customer portal.
func (r User_Customer) GetVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = r.Session.DoRequest("SoftLayer_User_Customer", "getVirtualGuests", nil, &r.Options, &resp)
	return
}

// Portal users are considered master users if they don't have an associated parent user. The only users who don't have parent users are users whose username matches their SoftLayer account name. Master users have special permissions throughout the SoftLayer customer portal.
// Deprecated: This function has been marked as deprecated.
func (r User_Customer) IsMasterUser() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_User_Customer", "isMasterUser", nil, &r.Options, &resp)
	return
}

// Remove a user's API authentication key, removing that user's access to query the SoftLayer API.
func (r User_Customer) RemoveApiAuthenticationKey(keyId *int) (resp bool, err error) {
	params := []interface{}{
		keyId,
	}
	err = r.Session.DoRequest("SoftLayer_User_Customer", "removeApiAuthenticationKey", params, &r.Options, &resp)
	return
}

// Remove (revoke) multiple permissions from a portal user's permission set. [[SoftLayer_User_Customer_CustomerPermission_Permission]] control which features in the SoftLayer customer portal and API a user may use. Removing a user's permission will affect that user's portal and API access. removePortalPermission() does not attempt to remove permissions that are not assigned to the user.
//
// Users can grant or revoke permissions to their child users, but not to themselves. An account's master has all portal permissions and can grant permissions for any of the other users on their account.
//
// If the cascadePermissionsFlag is set to true, then removing the permissions from a user will cascade down the child hierarchy and remove the permissions from this user along with all child users who also have the permission.
//
// If the cascadePermissionsFlag is not provided or is set to false and the user has children users who have the permission, then an exception will be thrown, and the permission will not be removed from this user.
//
// Use the [[SoftLayer_User_Customer_CustomerPermission_Permission::getAllObjects]] method to retrieve a list of all permissions available in the SoftLayer customer portal and API. Permissions are removed based on the keyName property of the permission objects within the permissions parameter.
func (r User_Customer) RemoveBulkPortalPermission(permissions []datatypes.User_Customer_CustomerPermission_Permission, cascadePermissionsFlag *bool) (resp bool, err error) {
	params := []interface{}{
		permissions,
		cascadePermissionsFlag,
	}
	err = r.Session.DoRequest("SoftLayer_User_Customer", "removeBulkPortalPermission", params, &r.Options, &resp)
	return
}

// Revokes access for the user to a single dedicated host device.  The user will only be allowed to see and access devices in both the portal and the API to which they have been granted access.  If the user's account has devices to which the user has not been granted access or the access has been revoked, then "not found" exceptions are thrown if the user attempts to access any of these devices.
//
// Users can assign device access to their child users, but not to themselves. An account's master has access to all devices on their customer account and can set dedicated host access for any of the other users on their account.
func (r User_Customer) RemoveDedicatedHostAccess(dedicatedHostId *int) (resp bool, err error) {
	params := []interface{}{
		dedicatedHostId,
	}
	err = r.Session.DoRequest("SoftLayer_User_Customer", "removeDedicatedHostAccess", params, &r.Options, &resp)
	return
}

// Remove hardware from a portal user's hardware access list. A user's hardware access list controls which of an account's hardware objects a user has access to in the SoftLayer customer portal and API. Hardware does not exist in the SoftLayer portal and returns "not found" exceptions in the API if the user doesn't have access to it. If a user does not has access to the hardware you're attempting remove add then removeHardwareAccess() returns true.
//
// Users can assign hardware access to their child users, but not to themselves. An account's master has access to all hardware on their customer account and can set hardware access for any of the other users on their account.
func (r User_Customer) RemoveHardwareAccess(hardwareId *int) (resp bool, err error) {
	params := []interface{}{
		hardwareId,
	}
	err = r.Session.DoRequest("SoftLayer_User_Customer", "removeHardwareAccess", params, &r.Options, &resp)
	return
}

// Remove a CloudLayer Computing Instance from a portal user's access list. A user's CloudLayer Computing Instance access list controls which of an account's computing instances a user has access to in the SoftLayer customer portal and API. CloudLayer Computing Instances do not exist in the SoftLayer portal and returns "not found" exceptions in the API if the user doesn't have access to it. If a user does not has access to the CloudLayer Computing Instance you're attempting remove add then removeVirtualGuestAccess() returns true.
//
// Users can assign CloudLayer Computing Instance access to their child users, but not to themselves. An account's master has access to all CloudLayer Computing Instances on their customer account and can set instance access for any of the other users on their account.
func (r User_Customer) RemoveVirtualGuestAccess(virtualGuestId *int) (resp bool, err error) {
	params := []interface{}{
		virtualGuestId,
	}
	err = r.Session.DoRequest("SoftLayer_User_Customer", "removeVirtualGuestAccess", params, &r.Options, &resp)
	return
}

// Update a user's VPN password on the SoftLayer customer portal. As with portal passwords, VPN passwords must match the following restrictions. VPN passwords must...
// * ...be over eight characters long.
// * ...be under twenty characters long.
// * ...contain at least one uppercase letter
// * ...contain at least one lowercase letter
// * ...contain at least one number
// * ...contain one of the special characters _ - | @ . , ? / ! ~ # $ % ^ & * ( ) { } [ ] \ =
// * ...not match your username
// Finally, users can only update their own VPN password. An account's master user can update any of their account users' VPN passwords.
func (r User_Customer) UpdateVpnPassword(password *string) (resp bool, err error) {
	params := []interface{}{
		password,
	}
	err = r.Session.DoRequest("SoftLayer_User_Customer", "updateVpnPassword", params, &r.Options, &resp)
	return
}

// Always call this function to enable changes when manually configuring VPN subnet access.
func (r User_Customer) UpdateVpnUser() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_User_Customer", "updateVpnUser", nil, &r.Options, &resp)
	return
}

// Each SoftLayer portal account is assigned a series of permissions that determine what access the user has to functions within the SoftLayer customer portal. This status is reflected in the SoftLayer_User_Customer_Status data type. Permissions differ from user status in that user status applies globally to the portal while user permissions are applied to specific portal functions.
type User_Customer_CustomerPermission_Permission struct {
	Session session.SLSession
	Options sl.Options
}

// GetUserCustomerCustomerPermissionPermissionService returns an instance of the User_Customer_CustomerPermission_Permission SoftLayer service
func GetUserCustomerCustomerPermissionPermissionService(sess session.SLSession) User_Customer_CustomerPermission_Permission {
	return User_Customer_CustomerPermission_Permission{Session: sess}
}

func (r User_Customer_CustomerPermission_Permission) Id(id int) User_Customer_CustomerPermission_Permission {
	r.Options.Id = &id
	return r
}

func (r User_Customer_CustomerPermission_Permission) Mask(mask string) User_Customer_CustomerPermission_Permission {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r User_Customer_CustomerPermission_Permission) Filter(filter string) User_Customer_CustomerPermission_Permission {
	r.Options.Filter = filter
	return r
}

func (r User_Customer_CustomerPermission_Permission) Limit(limit int) User_Customer_CustomerPermission_Permission {
	r.Options.Limit = &limit
	return r
}

func (r User_Customer_CustomerPermission_Permission) Offset(offset int) User_Customer_CustomerPermission_Permission {
	r.Options.Offset = &offset
	return r
}

// Retrieve all available permissions.
// Deprecated: This function has been marked as deprecated.
func (r User_Customer_CustomerPermission_Permission) GetAllObjects() (resp []datatypes.User_Customer_CustomerPermission_Permission, err error) {
	err = r.Session.DoRequest("SoftLayer_User_Customer_CustomerPermission_Permission", "getAllObjects", nil, &r.Options, &resp)
	return
}

// The Customer_Notification_Hardware object stores links between customers and the hardware devices they wish to monitor.  This link is not enough, the user must be sure to also create SoftLayer_Network_Monitor_Version1_Query_Host instance with the response action set to "notify users" in order for the users linked to that hardware object to be notified on failure.
type User_Customer_Notification_Hardware struct {
	Session session.SLSession
	Options sl.Options
}

// GetUserCustomerNotificationHardwareService returns an instance of the User_Customer_Notification_Hardware SoftLayer service
func GetUserCustomerNotificationHardwareService(sess session.SLSession) User_Customer_Notification_Hardware {
	return User_Customer_Notification_Hardware{Session: sess}
}

func (r User_Customer_Notification_Hardware) Id(id int) User_Customer_Notification_Hardware {
	r.Options.Id = &id
	return r
}

func (r User_Customer_Notification_Hardware) Mask(mask string) User_Customer_Notification_Hardware {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r User_Customer_Notification_Hardware) Filter(filter string) User_Customer_Notification_Hardware {
	r.Options.Filter = filter
	return r
}

func (r User_Customer_Notification_Hardware) Limit(limit int) User_Customer_Notification_Hardware {
	r.Options.Limit = &limit
	return r
}

func (r User_Customer_Notification_Hardware) Offset(offset int) User_Customer_Notification_Hardware {
	r.Options.Offset = &offset
	return r
}

// Passing in an unsaved instances of a Customer_Notification_Hardware object into this function will create the object and return the results to the user.
func (r User_Customer_Notification_Hardware) CreateObject(templateObject *datatypes.User_Customer_Notification_Hardware) (resp datatypes.User_Customer_Notification_Hardware, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_User_Customer_Notification_Hardware", "createObject", params, &r.Options, &resp)
	return
}

// Like any other API object, the customer notification objects can be deleted by passing an instance of them into this function.  The ID on the object must be set.
func (r User_Customer_Notification_Hardware) DeleteObjects(templateObjects []datatypes.User_Customer_Notification_Hardware) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = r.Session.DoRequest("SoftLayer_User_Customer_Notification_Hardware", "deleteObjects", params, &r.Options, &resp)
	return
}

// This method returns all Customer_Notification_Hardware objects associated with the passed in hardware ID as long as that hardware ID is owned by the current user's account.
//
// This behavior can also be accomplished by simply tapping monitoringUserNotification on the Hardware_Server object.
func (r User_Customer_Notification_Hardware) FindByHardwareId(hardwareId *int) (resp []datatypes.User_Customer_Notification_Hardware, err error) {
	params := []interface{}{
		hardwareId,
	}
	err = r.Session.DoRequest("SoftLayer_User_Customer_Notification_Hardware", "findByHardwareId", params, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_User_Customer_Notification_Hardware object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_User_Customer_Notification_Hardware service. You can only retrieve hardware notifications attached to hardware and users that belong to your account
func (r User_Customer_Notification_Hardware) GetObject() (resp datatypes.User_Customer_Notification_Hardware, err error) {
	err = r.Session.DoRequest("SoftLayer_User_Customer_Notification_Hardware", "getObject", nil, &r.Options, &resp)
	return
}

// The SoftLayer_User_Customer_Notification_Virtual_Guest object stores links between customers and the virtual guests they wish to monitor.  This link is not enough, the user must be sure to also create SoftLayer_Network_Monitor_Version1_Query_Host instance with the response action set to "notify users" in order for the users linked to that Virtual Guest object to be notified on failure.
type User_Customer_Notification_Virtual_Guest struct {
	Session session.SLSession
	Options sl.Options
}

// GetUserCustomerNotificationVirtualGuestService returns an instance of the User_Customer_Notification_Virtual_Guest SoftLayer service
func GetUserCustomerNotificationVirtualGuestService(sess session.SLSession) User_Customer_Notification_Virtual_Guest {
	return User_Customer_Notification_Virtual_Guest{Session: sess}
}

func (r User_Customer_Notification_Virtual_Guest) Id(id int) User_Customer_Notification_Virtual_Guest {
	r.Options.Id = &id
	return r
}

func (r User_Customer_Notification_Virtual_Guest) Mask(mask string) User_Customer_Notification_Virtual_Guest {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r User_Customer_Notification_Virtual_Guest) Filter(filter string) User_Customer_Notification_Virtual_Guest {
	r.Options.Filter = filter
	return r
}

func (r User_Customer_Notification_Virtual_Guest) Limit(limit int) User_Customer_Notification_Virtual_Guest {
	r.Options.Limit = &limit
	return r
}

func (r User_Customer_Notification_Virtual_Guest) Offset(offset int) User_Customer_Notification_Virtual_Guest {
	r.Options.Offset = &offset
	return r
}

// Passing in an unsaved instance of a SoftLayer_Customer_Notification_Virtual_Guest object into this function will create the object and return the results to the user.
func (r User_Customer_Notification_Virtual_Guest) CreateObject(templateObject *datatypes.User_Customer_Notification_Virtual_Guest) (resp datatypes.User_Customer_Notification_Virtual_Guest, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_User_Customer_Notification_Virtual_Guest", "createObject", params, &r.Options, &resp)
	return
}

// Like any other API object, the customer notification objects can be deleted by passing an instance of them into this function.  The ID on the object must be set.
func (r User_Customer_Notification_Virtual_Guest) DeleteObjects(templateObjects []datatypes.User_Customer_Notification_Virtual_Guest) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = r.Session.DoRequest("SoftLayer_User_Customer_Notification_Virtual_Guest", "deleteObjects", params, &r.Options, &resp)
	return
}

// This method returns all SoftLayer_User_Customer_Notification_Virtual_Guest objects associated with the passed in ID as long as that Virtual Guest ID is owned by the current user's account.
//
// This behavior can also be accomplished by simply tapping monitoringUserNotification on the Virtual_Guest object.
func (r User_Customer_Notification_Virtual_Guest) FindByGuestId(id *int) (resp []datatypes.User_Customer_Notification_Virtual_Guest, err error) {
	params := []interface{}{
		id,
	}
	err = r.Session.DoRequest("SoftLayer_User_Customer_Notification_Virtual_Guest", "findByGuestId", params, &r.Options, &resp)
	return
}

// The SoftLayer_User_Permission_Action data type contains local attributes to identify and describe the valid actions a customer user can perform within IMS.  This includes a name, key name, and description.  This data can not be modified by users of IMS.
//
// It also contains relational attributes that indicate which SoftLayer_User_Permission_Group's include the action.
type User_Permission_Action struct {
	Session session.SLSession
	Options sl.Options
}

// GetUserPermissionActionService returns an instance of the User_Permission_Action SoftLayer service
func GetUserPermissionActionService(sess session.SLSession) User_Permission_Action {
	return User_Permission_Action{Session: sess}
}

func (r User_Permission_Action) Id(id int) User_Permission_Action {
	r.Options.Id = &id
	return r
}

func (r User_Permission_Action) Mask(mask string) User_Permission_Action {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r User_Permission_Action) Filter(filter string) User_Permission_Action {
	r.Options.Filter = filter
	return r
}

func (r User_Permission_Action) Limit(limit int) User_Permission_Action {
	r.Options.Limit = &limit
	return r
}

func (r User_Permission_Action) Offset(offset int) User_Permission_Action {
	r.Options.Offset = &offset
	return r
}

// Object filters and result limits are enabled on this method.
func (r User_Permission_Action) GetAllObjects() (resp []datatypes.User_Permission_Action, err error) {
	err = r.Session.DoRequest("SoftLayer_User_Permission_Action", "getAllObjects", nil, &r.Options, &resp)
	return
}
