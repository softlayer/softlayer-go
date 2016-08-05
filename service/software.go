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

func (r *Software_AccountLicense) GetAllObjects() (resp []datatypes.Software_AccountLicense, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_AccountLicense) GetObject() (resp datatypes.Software_AccountLicense, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Software_AccountLicense) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_AccountLicense) GetBillingItem() (resp datatypes.Billing_Item, err error) {
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

func (r *Software_Component) GetLicenseFile() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component) GetObject() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component) GetVendorSetUpConfiguration() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
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
func (r *Software_Component) GetVirtualGuest() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Software_Component_Analytics struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentAnalyticsService() Software_Component_Analytics {
	return Software_Component_Analytics{Session: r}
}

type Software_Component_Analytics_Urchin struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentAnalyticsUrchinService() Software_Component_Analytics_Urchin {
	return Software_Component_Analytics_Urchin{Session: r}
}

type Software_Component_AntivirusSpyware struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentAntivirusSpywareService() Software_Component_AntivirusSpyware {
	return Software_Component_AntivirusSpyware{Session: r}
}

func (r *Software_Component_AntivirusSpyware) GetObject() (resp datatypes.Software_Component_AntivirusSpyware, err error) {
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

type Software_Component_AntivirusSpyware_Mcafee struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentAntivirusSpywareMcafeeService() Software_Component_AntivirusSpyware_Mcafee {
	return Software_Component_AntivirusSpyware_Mcafee{Session: r}
}

type Software_Component_AntivirusSpyware_Mcafee_Epo_Version36 struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentAntivirusSpywareMcafeeEpoVersion36Service() Software_Component_AntivirusSpyware_Mcafee_Epo_Version36 {
	return Software_Component_AntivirusSpyware_Mcafee_Epo_Version36{Session: r}
}

func (r *Software_Component_AntivirusSpyware_Mcafee_Epo_Version36) GetAgentDetails() (resp McAfee_Epolicy_Orchestrator_Version36_Agent_Details, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_AntivirusSpyware_Mcafee_Epo_Version36) GetCurrentAntivirusPolicy() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_AntivirusSpyware_Mcafee_Epo_Version36) GetDataFileVersion() (resp McAfee_Epolicy_Orchestrator_Version36_Product_Properties, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_AntivirusSpyware_Mcafee_Epo_Version36) GetLatestAccessProtectionEvents() (resp []McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_AccessProtection, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_AntivirusSpyware_Mcafee_Epo_Version36) GetLatestAntivirusEvents() (resp []McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_AntivirusSpyware_Mcafee_Epo_Version36) GetLatestSpywareEvents() (resp []McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_AntivirusSpyware_Mcafee_Epo_Version36) GetTransactionStatus() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Software_Component_AntivirusSpyware_Mcafee_Epo_Version45 struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentAntivirusSpywareMcafeeEpoVersion45Service() Software_Component_AntivirusSpyware_Mcafee_Epo_Version45 {
	return Software_Component_AntivirusSpyware_Mcafee_Epo_Version45{Session: r}
}

func (r *Software_Component_AntivirusSpyware_Mcafee_Epo_Version45) GetAgentDetails() (resp McAfee_Epolicy_Orchestrator_Version45_Agent_Details, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_AntivirusSpyware_Mcafee_Epo_Version45) GetCurrentAntivirusPolicy() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_AntivirusSpyware_Mcafee_Epo_Version45) GetDataFileVersion() (resp McAfee_Epolicy_Orchestrator_Version45_Product_Properties, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_AntivirusSpyware_Mcafee_Epo_Version45) GetLatestAccessProtectionEvents() (resp []McAfee_Epolicy_Orchestrator_Version45_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_AntivirusSpyware_Mcafee_Epo_Version45) GetLatestAntivirusEvents() (resp []McAfee_Epolicy_Orchestrator_Version45_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_AntivirusSpyware_Mcafee_Epo_Version45) GetLatestSpywareEvents() (resp []McAfee_Epolicy_Orchestrator_Version45_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_AntivirusSpyware_Mcafee_Epo_Version45) GetTransactionStatus() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Software_Component_ControlPanel struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentControlPanelService() Software_Component_ControlPanel {
	return Software_Component_ControlPanel{Session: r}
}

type Software_Component_ControlPanel_Cpanel struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentControlPanelCpanelService() Software_Component_ControlPanel_Cpanel {
	return Software_Component_ControlPanel_Cpanel{Session: r}
}

type Software_Component_ControlPanel_Idera struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentControlPanelIderaService() Software_Component_ControlPanel_Idera {
	return Software_Component_ControlPanel_Idera{Session: r}
}

type Software_Component_ControlPanel_Idera_ServerBackup struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentControlPanelIderaServerBackupService() Software_Component_ControlPanel_Idera_ServerBackup {
	return Software_Component_ControlPanel_Idera_ServerBackup{Session: r}
}

type Software_Component_ControlPanel_Microsoft struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentControlPanelMicrosoftService() Software_Component_ControlPanel_Microsoft {
	return Software_Component_ControlPanel_Microsoft{Session: r}
}

type Software_Component_ControlPanel_Microsoft_WebPlatform struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentControlPanelMicrosoftWebPlatformService() Software_Component_ControlPanel_Microsoft_WebPlatform {
	return Software_Component_ControlPanel_Microsoft_WebPlatform{Session: r}
}

type Software_Component_ControlPanel_Parallels struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentControlPanelParallelsService() Software_Component_ControlPanel_Parallels {
	return Software_Component_ControlPanel_Parallels{Session: r}
}

type Software_Component_ControlPanel_Parallels_Plesk struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentControlPanelParallelsPleskService() Software_Component_ControlPanel_Parallels_Plesk {
	return Software_Component_ControlPanel_Parallels_Plesk{Session: r}
}

type Software_Component_ControlPanel_R1soft struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentControlPanelR1softService() Software_Component_ControlPanel_R1soft {
	return Software_Component_ControlPanel_R1soft{Session: r}
}

type Software_Component_ControlPanel_R1soft_Cdp struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentControlPanelR1softCdpService() Software_Component_ControlPanel_R1soft_Cdp {
	return Software_Component_ControlPanel_R1soft_Cdp{Session: r}
}

type Software_Component_ControlPanel_R1soft_ServerBackup struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentControlPanelR1softServerBackupService() Software_Component_ControlPanel_R1soft_ServerBackup {
	return Software_Component_ControlPanel_R1soft_ServerBackup{Session: r}
}

type Software_Component_ControlPanel_Swsoft struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentControlPanelSwsoftService() Software_Component_ControlPanel_Swsoft {
	return Software_Component_ControlPanel_Swsoft{Session: r}
}

type Software_Component_ControlPanel_WebhostAutomation struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentControlPanelWebhostAutomationService() Software_Component_ControlPanel_WebhostAutomation {
	return Software_Component_ControlPanel_WebhostAutomation{Session: r}
}

type Software_Component_HostIps struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentHostIpsService() Software_Component_HostIps {
	return Software_Component_HostIps{Session: r}
}

func (r *Software_Component_HostIps) GetCurrentHostIpsPolicies() (resp []datatypes.Container_Software_Component_HostIps_Policy, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps) GetObject() (resp datatypes.Software_Component_HostIps, err error) {
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

type Software_Component_HostIps_Mcafee struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentHostIpsMcafeeService() Software_Component_HostIps_Mcafee {
	return Software_Component_HostIps_Mcafee{Session: r}
}

type Software_Component_HostIps_Mcafee_Epo_Version36_Hips struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentHostIpsMcafeeEpoVersion36HipsService() Software_Component_HostIps_Mcafee_Epo_Version36_Hips {
	return Software_Component_HostIps_Mcafee_Epo_Version36_Hips{Session: r}
}

func (r *Software_Component_HostIps_Mcafee_Epo_Version36_Hips) GetAgentDetails() (resp McAfee_Epolicy_Orchestrator_Version36_Agent_Details, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps_Mcafee_Epo_Version36_Hips) GetApplicationModePolicyNames() (resp []McAfee_Epolicy_Orchestrator_Version36_Policy_Object, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps_Mcafee_Epo_Version36_Hips) GetApplicationRuleSetPolicyNames() (resp []McAfee_Epolicy_Orchestrator_Version36_Policy_Object, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps_Mcafee_Epo_Version36_Hips) GetEnforcementPolicyNames() (resp []McAfee_Epolicy_Orchestrator_Version36_Policy_Object, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps_Mcafee_Epo_Version36_Hips) GetFirewallModePolicyNames() (resp []McAfee_Epolicy_Orchestrator_Version36_Policy_Object, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps_Mcafee_Epo_Version36_Hips) GetFirewallRuleSetPolicyNames() (resp []McAfee_Epolicy_Orchestrator_Version36_Policy_Object, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps_Mcafee_Epo_Version36_Hips) GetIpsModePolicyNames() (resp []McAfee_Epolicy_Orchestrator_Version36_Policy_Object, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps_Mcafee_Epo_Version36_Hips) GetIpsProtectionPolicyNames() (resp []McAfee_Epolicy_Orchestrator_Version36_Policy_Object, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps_Mcafee_Epo_Version36_Hips) GetTransactionStatus() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version6 struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentHostIpsMcafeeEpoVersion36HipsVersion6Service() Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version6 {
	return Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version6{Session: r}
}

func (r *Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version6) GetBlockedApplicationEvents() (resp []McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_BlockedApplicationEvent, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version6) GetIpsEvents() (resp []McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_IPSEvent, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version7 struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentHostIpsMcafeeEpoVersion36HipsVersion7Service() Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version7 {
	return Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version7{Session: r}
}

func (r *Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version7) GetBlockedApplicationEvents() (resp []McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_BlockedApplicationEvent, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version7) GetIpsEvents() (resp []McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Software_Component_HostIps_Mcafee_Epo_Version45_Hips struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentHostIpsMcafeeEpoVersion45HipsService() Software_Component_HostIps_Mcafee_Epo_Version45_Hips {
	return Software_Component_HostIps_Mcafee_Epo_Version45_Hips{Session: r}
}

func (r *Software_Component_HostIps_Mcafee_Epo_Version45_Hips) GetAgentDetails() (resp McAfee_Epolicy_Orchestrator_Version45_Agent_Details, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps_Mcafee_Epo_Version45_Hips) GetApplicationModePolicyNames() (resp []McAfee_Epolicy_Orchestrator_Version45_Policy_Object, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps_Mcafee_Epo_Version45_Hips) GetApplicationRuleSetPolicyNames() (resp []McAfee_Epolicy_Orchestrator_Version45_Policy_Object, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps_Mcafee_Epo_Version45_Hips) GetBlockedApplicationEvents() (resp []McAfee_Epolicy_Orchestrator_Version45_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps_Mcafee_Epo_Version45_Hips) GetEnforcementPolicyNames() (resp []McAfee_Epolicy_Orchestrator_Version45_Policy_Object, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps_Mcafee_Epo_Version45_Hips) GetFirewallModePolicyNames() (resp []McAfee_Epolicy_Orchestrator_Version45_Policy_Object, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps_Mcafee_Epo_Version45_Hips) GetFirewallRuleSetPolicyNames() (resp []McAfee_Epolicy_Orchestrator_Version45_Policy_Object, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps_Mcafee_Epo_Version45_Hips) GetIpsEvents() (resp []McAfee_Epolicy_Orchestrator_Version45_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps_Mcafee_Epo_Version45_Hips) GetIpsModePolicyNames() (resp []McAfee_Epolicy_Orchestrator_Version45_Policy_Object, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps_Mcafee_Epo_Version45_Hips) GetIpsProtectionPolicyNames() (resp []McAfee_Epolicy_Orchestrator_Version45_Policy_Object, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_HostIps_Mcafee_Epo_Version45_Hips) GetTransactionStatus() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Software_Component_HostIps_Mcafee_Epo_Version45_Hips_Version7 struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentHostIpsMcafeeEpoVersion45HipsVersion7Service() Software_Component_HostIps_Mcafee_Epo_Version45_Hips_Version7 {
	return Software_Component_HostIps_Mcafee_Epo_Version45_Hips_Version7{Session: r}
}

type Software_Component_HostIps_Mcafee_Epo_Version45_Hips_Version8 struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentHostIpsMcafeeEpoVersion45HipsVersion8Service() Software_Component_HostIps_Mcafee_Epo_Version45_Hips_Version8 {
	return Software_Component_HostIps_Mcafee_Epo_Version45_Hips_Version8{Session: r}
}

type Software_Component_OperatingSystem struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentOperatingSystemService() Software_Component_OperatingSystem {
	return Software_Component_OperatingSystem{Session: r}
}

func (r *Software_Component_OperatingSystem) GetLicenseExpirationDate() (resp datatypes.Time, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_OperatingSystem) GetPartitionTemplates() (resp []datatypes.Hardware_Component_Partition_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Component_OperatingSystem) GetReloadTransactionGroup() (resp datatypes.Provisioning_Version1_Transaction_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Software_Component_Package struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentPackageService() Software_Component_Package {
	return Software_Component_Package{Session: r}
}

type Software_Component_Package_Management struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentPackageManagementService() Software_Component_Package_Management {
	return Software_Component_Package_Management{Session: r}
}

type Software_Component_Package_Management_Ksplice struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentPackageManagementKspliceService() Software_Component_Package_Management_Ksplice {
	return Software_Component_Package_Management_Ksplice{Session: r}
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

type Software_Component_Password_History struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentPasswordHistoryService() Software_Component_Password_History {
	return Software_Component_Password_History{Session: r}
}

func (r *Software_Component_Password_History) GetSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Software_Component_Security struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentSecurityService() Software_Component_Security {
	return Software_Component_Security{Session: r}
}

type Software_Component_Security_SafeNet struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareComponentSecuritySafeNetService() Software_Component_Security_SafeNet {
	return Software_Component_Security_SafeNet{Session: r}
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
func (r *Software_Description) GetObject() (resp datatypes.Software_Description, err error) {
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

type Software_Description_Attribute struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareDescriptionAttributeService() Software_Description_Attribute {
	return Software_Description_Attribute{Session: r}
}

func (r *Software_Description_Attribute) GetSoftwareDescription() (resp datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_Description_Attribute) GetType() (resp datatypes.Software_Description_Attribute_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Software_Description_Attribute_Type struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareDescriptionAttributeTypeService() Software_Description_Attribute_Type {
	return Software_Description_Attribute_Type{Session: r}
}

type Software_Description_Feature struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareDescriptionFeatureService() Software_Description_Feature {
	return Software_Description_Feature{Session: r}
}

type Software_Description_RequiredUser struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareDescriptionRequiredUserService() Software_Description_RequiredUser {
	return Software_Description_RequiredUser{Session: r}
}

type Software_License struct {
	Session *Session
	Options
}

func (r *Session) GetSoftwareLicenseService() Software_License {
	return Software_License{Session: r}
}

func (r *Software_License) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_License) GetOwner() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_License) GetSoftwareDescription() (resp datatypes.Software_Description, err error) {
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

func (r *Software_VirtualLicense) GetLicenseFile() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_VirtualLicense) GetObject() (resp datatypes.Software_VirtualLicense, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
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
func (r *Software_VirtualLicense) GetSoftwareDescription() (resp datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Software_VirtualLicense) GetSubnet() (resp datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
