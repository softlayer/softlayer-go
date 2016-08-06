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

type Software_AccountLicense struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareAccountLicenseService() Software_AccountLicense {
	return Software_AccountLicense{Session: r}
}

func (r *Software_AccountLicense) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_AccountLicense) GetAllObjects() (resp []datatypes.Software_AccountLicense, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_AccountLicense) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_AccountLicense) GetObject() (resp datatypes.Software_AccountLicense, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_AccountLicense) GetSoftwareDescription() (resp datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Software_Component struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentService() Software_Component {
	return Software_Component{Session: r}
}

func (r *Software_Component) GetAverageInstallationDuration() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component) GetLicenseFile() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component) GetObject() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component) GetPasswordHistory() (resp []datatypes.Software_Component_Password_History, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component) GetPasswords() (resp []datatypes.Software_Component_Password, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component) GetSoftwareDescription() (resp datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component) GetSoftwareLicense() (resp datatypes.Software_License, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component) GetVendorSetUpConfiguration() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component) GetVirtualGuest() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Software_Component_AntivirusSpyware struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentAntivirusSpywareService() Software_Component_AntivirusSpyware {
	return Software_Component_AntivirusSpyware{Session: r}
}

func (r *Software_Component_AntivirusSpyware) GetAverageInstallationDuration() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_AntivirusSpyware) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_AntivirusSpyware) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_AntivirusSpyware) GetLicenseFile() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_AntivirusSpyware) GetObject() (resp datatypes.Software_Component_AntivirusSpyware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_AntivirusSpyware) GetPasswordHistory() (resp []datatypes.Software_Component_Password_History, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_AntivirusSpyware) GetPasswords() (resp []datatypes.Software_Component_Password, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_AntivirusSpyware) GetSoftwareDescription() (resp datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_AntivirusSpyware) GetSoftwareLicense() (resp datatypes.Software_License, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_AntivirusSpyware) GetVendorSetUpConfiguration() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_AntivirusSpyware) GetVirtualGuest() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_AntivirusSpyware) UpdateAntivirusSpywarePolicy(newPolicy *string, enforce *bool) (resp bool, err error) {
	params := []interface{}{
		newPolicy,
		enforce,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Software_Component_HostIps struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentHostIpsService() Software_Component_HostIps {
	return Software_Component_HostIps{Session: r}
}

func (r *Software_Component_HostIps) GetAverageInstallationDuration() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps) GetCurrentHostIpsPolicies() (resp []datatypes.Container_Software_Component_HostIps_Policy, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps) GetLicenseFile() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps) GetObject() (resp datatypes.Software_Component_HostIps, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps) GetPasswordHistory() (resp []datatypes.Software_Component_Password_History, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps) GetPasswords() (resp []datatypes.Software_Component_Password, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps) GetSoftwareDescription() (resp datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps) GetSoftwareLicense() (resp datatypes.Software_License, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps) GetVendorSetUpConfiguration() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps) GetVirtualGuest() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
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

type Software_Component_Password struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentPasswordService() Software_Component_Password {
	return Software_Component_Password{Session: r}
}

func (r *Software_Component_Password) CreateObject(templateObject *datatypes.Software_Component_Password) (resp datatypes.Software_Component_Password, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_Password) CreateObjects(templateObjects []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_Password) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_Password) DeleteObjects(templateObjects []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_Password) EditObject(templateObject *datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_Password) EditObjects(templateObjects []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_Password) GetObject() (resp datatypes.Software_Component_Password, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_Password) GetSoftware() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_Password) GetSshKeys() (resp []datatypes.Security_Ssh_Key, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Software_Description struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareDescriptionService() Software_Description {
	return Software_Description{Session: r}
}

func (r *Software_Description) GetAllObjects() (resp []datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Description) GetAttributes() (resp []datatypes.Software_Description_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Description) GetAverageInstallationDuration() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Description) GetCompatibleSoftwareDescriptions() (resp []datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Description) GetFeatures() (resp []datatypes.Software_Description_Feature, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Description) GetLatestVersion() (resp []datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Description) GetObject() (resp datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Description) GetProductItems() (resp []datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Description) GetProvisionTransactionGroup() (resp datatypes.Provisioning_Version1_Transaction_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Description) GetReloadTransactionGroup() (resp datatypes.Provisioning_Version1_Transaction_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Description) GetRequiredUser() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Description) GetSoftwareLicenses() (resp []datatypes.Software_License, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Description) GetUpgradeSoftwareDescription() (resp datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Description) GetUpgradeSwDesc() (resp datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Description) GetValidFilesystemTypes() (resp []datatypes.Configuration_Storage_Filesystem_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Software_VirtualLicense struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareVirtualLicenseService() Software_VirtualLicense {
	return Software_VirtualLicense{Session: r}
}

func (r *Software_VirtualLicense) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_VirtualLicense) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_VirtualLicense) GetHostHardware() (resp datatypes.Hardware_Server, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_VirtualLicense) GetIpAddressRecord() (resp datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_VirtualLicense) GetLicenseFile() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_VirtualLicense) GetObject() (resp datatypes.Software_VirtualLicense, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_VirtualLicense) GetSoftwareDescription() (resp datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_VirtualLicense) GetSubnet() (resp datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
