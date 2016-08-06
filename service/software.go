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

package service

import "github.ibm.com/riethm/gopherlayer/datatypes"

// SoftLayer_Software_AccountLicense is a class that represents software licenses that are tied only to a customer's account and not to any particular hardware, IP address, etc.
type Software_AccountLicense struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareAccountLicenseService() Software_AccountLicense {
	return Software_AccountLicense{Session: r}
}

// Retrieve The customer account this Account License belongs to.
func (r *Software_AccountLicense) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Software_AccountLicense) GetAllObjects() (resp []datatypes.Software_AccountLicense, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The billing item for a software account license.
func (r *Software_AccountLicense) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Software_AccountLicense) GetObject() (resp datatypes.Software_AccountLicense, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Software_Description that this account license is for.
func (r *Software_AccountLicense) GetSoftwareDescription() (resp datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// A SoftLayer_Software_Component ties the installation of a specific piece of software onto a specific piece of hardware.
//
// SoftLayer_Software_Component works with SoftLayer_Software_License and SoftLayer_Software_Description to tie this all together.
//
// <ul> <li>SoftLayer_Software_Component is the installation of a specific piece of software onto a specific piece of hardware in accordance to a software license. <ul> <li>SoftLayer_Software_License dictates when and how a specific piece of software may be installed onto a piece of hardware. <ul> <li>SoftLayer_Software_Description describes a specific piece of software which can be installed onto hardware in accordance with it's license agreement. </li></ul></li></ul></li></ul>
type Software_Component struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentService() Software_Component {
	return Software_Component{Session: r}
}

// Retrieve The average amount of time that a software component takes to install.
func (r *Software_Component) GetAverageInstallationDuration() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The billing item for a software component.
func (r *Software_Component) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware this Software Component is installed upon.
func (r *Software_Component) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Attempt to retrieve the file associated with a software component.  If the software component does not support downloading license files an exception will be thrown.
func (r *Software_Component) GetLicenseFile() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Software_Component object whose ID corresponds to the ID number of the init parameter passed to the SoftLayer_Software_Component service.
//
// The best way to get software components is through getSoftwareComponents from the Hardware service.
func (r *Software_Component) GetObject() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve History Records for Software Passwords.
func (r *Software_Component) GetPasswordHistory() (resp []datatypes.Software_Component_Password_History, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Username/Password pairs used for access to this Software Installation.
func (r *Software_Component) GetPasswords() (resp []datatypes.Software_Component_Password, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The Software Description of this Software Component.
func (r *Software_Component) GetSoftwareDescription() (resp datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The License this Software Component uses.
func (r *Software_Component) GetSoftwareLicense() (resp datatypes.Software_License, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Software_Component) GetVendorSetUpConfiguration() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The virtual guest this software component is installed upon.
func (r *Software_Component) GetVirtualGuest() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This object specifies a specific type of Software Component:  An Anti-virus/spyware instance. Anti-virus/spyware installations have specific properties and methods such as SoftLayer_Software_Component_AntivirusSpyware::updateAntivirusSpywarePolicy. Defaults are initiated by this object.
type Software_Component_AntivirusSpyware struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentAntivirusSpywareService() Software_Component_AntivirusSpyware {
	return Software_Component_AntivirusSpyware{Session: r}
}

// Retrieve The average amount of time that a software component takes to install.
func (r *Software_Component_AntivirusSpyware) GetAverageInstallationDuration() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The billing item for a software component.
func (r *Software_Component_AntivirusSpyware) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware this Software Component is installed upon.
func (r *Software_Component_AntivirusSpyware) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Attempt to retrieve the file associated with a software component.  If the software component does not support downloading license files an exception will be thrown.
func (r *Software_Component_AntivirusSpyware) GetLicenseFile() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Software_Component_AntivirusSpyware) GetObject() (resp datatypes.Software_Component_AntivirusSpyware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve History Records for Software Passwords.
func (r *Software_Component_AntivirusSpyware) GetPasswordHistory() (resp []datatypes.Software_Component_Password_History, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Username/Password pairs used for access to this Software Installation.
func (r *Software_Component_AntivirusSpyware) GetPasswords() (resp []datatypes.Software_Component_Password, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The Software Description of this Software Component.
func (r *Software_Component_AntivirusSpyware) GetSoftwareDescription() (resp datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The License this Software Component uses.
func (r *Software_Component_AntivirusSpyware) GetSoftwareLicense() (resp datatypes.Software_License, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Software_Component_AntivirusSpyware) GetVendorSetUpConfiguration() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The virtual guest this software component is installed upon.
func (r *Software_Component_AntivirusSpyware) GetVirtualGuest() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Update an anti-virus/spyware policy. The policy options that it accepts are the following:
// *1 - Minimal
// *2 - Relaxed
// *3 - Default
// *4 - High
// *5 - Ultimate
func (r *Software_Component_AntivirusSpyware) UpdateAntivirusSpywarePolicy(newPolicy *string, enforce *bool) (resp bool, err error) {
	params := []interface{}{
		newPolicy,
		enforce,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// This object specifies a specific type of Software Component:  A Host Intrusion Protection System instance.
type Software_Component_HostIps struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentHostIpsService() Software_Component_HostIps {
	return Software_Component_HostIps{Session: r}
}

// Retrieve The average amount of time that a software component takes to install.
func (r *Software_Component_HostIps) GetAverageInstallationDuration() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The billing item for a software component.
func (r *Software_Component_HostIps) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Get the current Host IPS policies.
func (r *Software_Component_HostIps) GetCurrentHostIpsPolicies() (resp []datatypes.Container_Software_Component_HostIps_Policy, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware this Software Component is installed upon.
func (r *Software_Component_HostIps) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Attempt to retrieve the file associated with a software component.  If the software component does not support downloading license files an exception will be thrown.
func (r *Software_Component_HostIps) GetLicenseFile() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Software_Component_HostIps) GetObject() (resp datatypes.Software_Component_HostIps, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve History Records for Software Passwords.
func (r *Software_Component_HostIps) GetPasswordHistory() (resp []datatypes.Software_Component_Password_History, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Username/Password pairs used for access to this Software Installation.
func (r *Software_Component_HostIps) GetPasswords() (resp []datatypes.Software_Component_Password, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The Software Description of this Software Component.
func (r *Software_Component_HostIps) GetSoftwareDescription() (resp datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The License this Software Component uses.
func (r *Software_Component_HostIps) GetSoftwareLicense() (resp datatypes.Software_License, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Software_Component_HostIps) GetVendorSetUpConfiguration() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The virtual guest this software component is installed upon.
func (r *Software_Component_HostIps) GetVirtualGuest() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Update the Host IPS policies. To retrieve valid policy options you must use the provided relationships.
func (r *Software_Component_HostIps) UpdateHipsPolicies(newIpsMode *string, newIpsProtection *string, newFirewallMode *string, newFirewallRuleset *string, newApplicationMode *string, newApplicationRuleset *string, newEnforcementPolicy *string) (resp bool, err error) {
	params := []interface{}{
		newIpsMode,
		newIpsProtection,
		newFirewallMode,
		newFirewallRuleset,
		newApplicationMode,
		newApplicationRuleset,
		newEnforcementPolicy,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// This SoftLayer_Software_Component_Password data type contains a password for a specific software component instance.
type Software_Component_Password struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentPasswordService() Software_Component_Password {
	return Software_Component_Password{Session: r}
}

// Create a password for a software component.
func (r *Software_Component_Password) CreateObject(templateObject *datatypes.Software_Component_Password) (resp datatypes.Software_Component_Password, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Create more than one password for a software component.
func (r *Software_Component_Password) CreateObjects(templateObjects []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Delete a password from a software component.
func (r *Software_Component_Password) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Delete more than one passwords from a software component.
func (r *Software_Component_Password) DeleteObjects(templateObjects []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Edit the properties of a software component password such as the username, password, port, and notes.
func (r *Software_Component_Password) EditObject(templateObject *datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Edit more than one password from a software component.
func (r *Software_Component_Password) EditObjects(templateObjects []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Software_Component_Password) GetObject() (resp datatypes.Software_Component_Password, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Software_Component instance that this username/password pair is valid for.
func (r *Software_Component_Password) GetSoftware() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve SSH keys to be installed on the server during provisioning or an OS reload.
func (r *Software_Component_Password) GetSshKeys() (resp []datatypes.Security_Ssh_Key, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This class holds a description for a specific installation of a Software Component.
//
// SoftLayer_Software_Licenses tie a Software Component (A specific installation on a piece of hardware) to it's description.
//
// The "Manufacturer" and "Name" properties of a SoftLayer_Software_Description are used by the framework to factory specific objects, objects that may have special methods for that specific piece of software, or objects that contain application specific data, such as default ports.  For example, if you create a SoftLayer_Software_Component who's SoftLayer_Software_License points to the SoftLayer_Software_Description for "Swsoft" "Plesk", you'll actually get a SoftLayer_Software_Component_Swsoft_Plesk object.
type Software_Description struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareDescriptionService() Software_Description {
	return Software_Description{Session: r}
}

//
func (r *Software_Description) GetAllObjects() (resp []datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Software_Description) GetAttributes() (resp []datatypes.Software_Description_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The average amount of time that a software description takes to install.
func (r *Software_Description) GetAverageInstallationDuration() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A list of the software descriptions that are compatible with this software description.
func (r *Software_Description) GetCompatibleSoftwareDescriptions() (resp []datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The feature attributes of a software description.
func (r *Software_Description) GetFeatures() (resp []datatypes.Software_Description_Feature, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The latest version of a software description.
func (r *Software_Description) GetLatestVersion() (resp []datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Software_Description) GetObject() (resp datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The various product items to which this software description is linked.
func (r *Software_Description) GetProductItems() (resp []datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve This details the provisioning transaction group for this software. This is only valid for Operating System software.
func (r *Software_Description) GetProvisionTransactionGroup() (resp datatypes.Provisioning_Version1_Transaction_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The transaction group that a software description belongs to. A transaction group is a sequence of transactions that must be performed in a specific order for the installation of software.
func (r *Software_Description) GetReloadTransactionGroup() (resp datatypes.Provisioning_Version1_Transaction_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The default user created for a given a software description.
func (r *Software_Description) GetRequiredUser() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Software Licenses that govern this Software Description.
func (r *Software_Description) GetSoftwareLicenses() (resp []datatypes.Software_License, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A suggestion for an upgrade path from this Software Description
func (r *Software_Description) GetUpgradeSoftwareDescription() (resp datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A suggestion for an upgrade path from this Software Description (Deprecated - Use upgradeSoftwareDescription)
func (r *Software_Description) GetUpgradeSwDesc() (resp datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Software_Description) GetValidFilesystemTypes() (resp []datatypes.Configuration_Storage_Filesystem_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// SoftLayer_Software_VirtualLicense is the application class that handles a special type of Software License.  Most software licenses are licensed to a specific hardware ID;  virtual licenses are designed for virtual machines and therefore are assigned to an IP Address.  Not all software packages can be "virtual licensed".
type Software_VirtualLicense struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareVirtualLicenseService() Software_VirtualLicense {
	return Software_VirtualLicense{Session: r}
}

// Retrieve The customer account this Virtual License belongs to.
func (r *Software_VirtualLicense) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The billing item for a software virtual license.
func (r *Software_VirtualLicense) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware record to which the software virtual license is assigned.
func (r *Software_VirtualLicense) GetHostHardware() (resp datatypes.Hardware_Server, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The IP Address record associated with a virtual license.
func (r *Software_VirtualLicense) GetIpAddressRecord() (resp datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Attempt to retrieve the file associated with a virtual license, if such a file exists.  If there is no file for this virtual license, calling this method will either throw an exception or return false.
func (r *Software_VirtualLicense) GetLicenseFile() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Software_VirtualLicense object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Software_VirtualLicense service. You can only retrieve Virtual Licenses assigned to your account number.
func (r *Software_VirtualLicense) GetObject() (resp datatypes.Software_VirtualLicense, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Software_Description that this virtual license is for.
func (r *Software_VirtualLicense) GetSoftwareDescription() (resp datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The subnet this Virtual License's IP address belongs to.
func (r *Software_VirtualLicense) GetSubnet() (resp datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
