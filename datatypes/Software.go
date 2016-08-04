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

package datatypes

import "time"

type Software_AccountLicense struct {
	Entity

	Account             *Account              `json:"account,omitempty"`
	AccountId           *int                  `json:"accountId,omitempty"`
	BillingItem         *Billing_Item         `json:"billingItem,omitempty"`
	Capacity            *string               `json:"capacity,omitempty"`
	Key                 *string               `json:"key,omitempty"`
	SoftwareDescription *Software_Description `json:"softwareDescription,omitempty"`
	Units               *string               `json:"units,omitempty"`
}

type Software_Component struct {
	Entity

	AverageInstallationDuration *uint                                 `json:"averageInstallationDuration,omitempty"`
	BillingItem                 *Billing_Item                         `json:"billingItem,omitempty"`
	Hardware                    *Hardware                             `json:"hardware,omitempty"`
	HardwareId                  *int                                  `json:"hardwareId,omitempty"`
	Id                          *int                                  `json:"id,omitempty"`
	ManufacturerActivationCode  *string                               `json:"manufacturerActivationCode,omitempty"`
	ManufacturerLicenseInstance *string                               `json:"manufacturerLicenseInstance,omitempty"`
	PasswordCount               *uint                                 `json:"passwordCount,omitempty"`
	PasswordHistory             []Software_Component_Password_History `json:"passwordHistory,omitempty"`
	PasswordHistoryCount        *uint                                 `json:"passwordHistoryCount,omitempty"`
	Passwords                   []Software_Component_Password         `json:"passwords,omitempty"`
	SoftwareDescription         *Software_Description                 `json:"softwareDescription,omitempty"`
	SoftwareLicense             *Software_License                     `json:"softwareLicense,omitempty"`
	VirtualGuest                *Virtual_Guest                        `json:"virtualGuest,omitempty"`
}

type Software_Component_Analytics struct {
	Software_Component
}

type Software_Component_Analytics_Urchin struct {
	Software_Component_Analytics
}

type Software_Component_AntivirusSpyware struct {
	Software_Component
}

type Software_Component_AntivirusSpyware_Mcafee struct {
	Software_Component_AntivirusSpyware
}

type Software_Component_AntivirusSpyware_Mcafee_Epo_Version36 struct {
	Software_Component_AntivirusSpyware_Mcafee

	AgentDetails                     *McAfee_Epolicy_Orchestrator_Version36_Agent_Details                     `json:"agentDetails,omitempty"`
	CurrentAntivirusPolicy           *int                                                                     `json:"currentAntivirusPolicy,omitempty"`
	DataFileVersion                  *McAfee_Epolicy_Orchestrator_Version36_Product_Properties                `json:"dataFileVersion,omitempty"`
	EpoVersion                       *string                                                                  `json:"epoVersion,omitempty"`
	LatestAccessProtectionEventCount *uint                                                                    `json:"latestAccessProtectionEventCount,omitempty"`
	LatestAccessProtectionEvents     []McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_AccessProtection `json:"latestAccessProtectionEvents,omitempty"`
	LatestAntivirusEventCount        *uint                                                                    `json:"latestAntivirusEventCount,omitempty"`
	LatestAntivirusEvents            []McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event                  `json:"latestAntivirusEvents,omitempty"`
	LatestSpywareEventCount          *uint                                                                    `json:"latestSpywareEventCount,omitempty"`
	LatestSpywareEvents              []McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event                  `json:"latestSpywareEvents,omitempty"`
	TransactionStatus                *string                                                                  `json:"transactionStatus,omitempty"`
}

type Software_Component_AntivirusSpyware_Mcafee_Epo_Version45 struct {
	Software_Component_AntivirusSpyware_Mcafee

	AgentDetails                     *McAfee_Epolicy_Orchestrator_Version45_Agent_Details      `json:"agentDetails,omitempty"`
	CurrentAntivirusPolicy           *int                                                      `json:"currentAntivirusPolicy,omitempty"`
	DataFileVersion                  *McAfee_Epolicy_Orchestrator_Version45_Product_Properties `json:"dataFileVersion,omitempty"`
	EpoVersion                       *string                                                   `json:"epoVersion,omitempty"`
	LatestAccessProtectionEventCount *uint                                                     `json:"latestAccessProtectionEventCount,omitempty"`
	LatestAccessProtectionEvents     []McAfee_Epolicy_Orchestrator_Version45_Event             `json:"latestAccessProtectionEvents,omitempty"`
	LatestAntivirusEventCount        *uint                                                     `json:"latestAntivirusEventCount,omitempty"`
	LatestAntivirusEvents            []McAfee_Epolicy_Orchestrator_Version45_Event             `json:"latestAntivirusEvents,omitempty"`
	LatestSpywareEventCount          *uint                                                     `json:"latestSpywareEventCount,omitempty"`
	LatestSpywareEvents              []McAfee_Epolicy_Orchestrator_Version45_Event             `json:"latestSpywareEvents,omitempty"`
	TransactionStatus                *string                                                   `json:"transactionStatus,omitempty"`
}

type Software_Component_ControlPanel struct {
	Software_Component
}

type Software_Component_ControlPanel_Cpanel struct {
	Software_Component
}

type Software_Component_ControlPanel_Idera struct {
	Software_Component
}

type Software_Component_ControlPanel_Idera_ServerBackup struct {
	Software_Component_ControlPanel_Idera
}

type Software_Component_ControlPanel_Microsoft struct {
	Software_Component
}

type Software_Component_ControlPanel_Microsoft_WebPlatform struct {
	Software_Component_ControlPanel_Microsoft
}

type Software_Component_ControlPanel_Parallels struct {
	Software_Component
}

type Software_Component_ControlPanel_Parallels_Plesk struct {
	Software_Component_ControlPanel_Parallels
}

type Software_Component_ControlPanel_R1soft struct {
	Software_Component
}

type Software_Component_ControlPanel_R1soft_Cdp struct {
	Software_Component_ControlPanel_R1soft
}

type Software_Component_ControlPanel_R1soft_ServerBackup struct {
	Software_Component_ControlPanel_R1soft
}

type Software_Component_ControlPanel_Swsoft struct {
	Software_Component
}

type Software_Component_ControlPanel_WebhostAutomation struct {
	Software_Component
}

type Software_Component_HostIps struct {
	Software_Component
}

type Software_Component_HostIps_Mcafee struct {
	Software_Component_HostIps
}

type Software_Component_HostIps_Mcafee_Epo_Version36_Hips struct {
	Software_Component_HostIps_Mcafee

	AgentDetails                      *McAfee_Epolicy_Orchestrator_Version36_Agent_Details  `json:"agentDetails,omitempty"`
	ApplicationModePolicyNameCount    *uint                                                 `json:"applicationModePolicyNameCount,omitempty"`
	ApplicationModePolicyNames        []McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"applicationModePolicyNames,omitempty"`
	ApplicationRuleSetPolicyNameCount *uint                                                 `json:"applicationRuleSetPolicyNameCount,omitempty"`
	ApplicationRuleSetPolicyNames     []McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"applicationRuleSetPolicyNames,omitempty"`
	EnforcementPolicyNameCount        *uint                                                 `json:"enforcementPolicyNameCount,omitempty"`
	EnforcementPolicyNames            []McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"enforcementPolicyNames,omitempty"`
	EpoVersion                        *string                                               `json:"epoVersion,omitempty"`
	FirewallModePolicyNameCount       *uint                                                 `json:"firewallModePolicyNameCount,omitempty"`
	FirewallModePolicyNames           []McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"firewallModePolicyNames,omitempty"`
	FirewallRuleSetPolicyNameCount    *uint                                                 `json:"firewallRuleSetPolicyNameCount,omitempty"`
	FirewallRuleSetPolicyNames        []McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"firewallRuleSetPolicyNames,omitempty"`
	IpsModePolicyNameCount            *uint                                                 `json:"ipsModePolicyNameCount,omitempty"`
	IpsModePolicyNames                []McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"ipsModePolicyNames,omitempty"`
	IpsProtectionPolicyNameCount      *uint                                                 `json:"ipsProtectionPolicyNameCount,omitempty"`
	IpsProtectionPolicyNames          []McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"ipsProtectionPolicyNames,omitempty"`
	TransactionStatus                 *string                                               `json:"transactionStatus,omitempty"`
}

type Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version6 struct {
	Software_Component_HostIps_Mcafee_Epo_Version36_Hips

	BlockedApplicationEventCount *uint                                                                         `json:"blockedApplicationEventCount,omitempty"`
	BlockedApplicationEvents     []McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_BlockedApplicationEvent `json:"blockedApplicationEvents,omitempty"`
	IpsEventCount                *uint                                                                         `json:"ipsEventCount,omitempty"`
	IpsEvents                    []McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_IPSEvent                `json:"ipsEvents,omitempty"`
}

type Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version7 struct {
	Software_Component_HostIps_Mcafee_Epo_Version36_Hips

	BlockedApplicationEventCount *uint                                                                         `json:"blockedApplicationEventCount,omitempty"`
	BlockedApplicationEvents     []McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_BlockedApplicationEvent `json:"blockedApplicationEvents,omitempty"`
	IpsEventCount                *uint                                                                         `json:"ipsEventCount,omitempty"`
	IpsEvents                    []McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent                `json:"ipsEvents,omitempty"`
}

type Software_Component_HostIps_Mcafee_Epo_Version45_Hips struct {
	Software_Component_HostIps_Mcafee

	AgentDetails                      *McAfee_Epolicy_Orchestrator_Version45_Agent_Details  `json:"agentDetails,omitempty"`
	ApplicationModePolicyNameCount    *uint                                                 `json:"applicationModePolicyNameCount,omitempty"`
	ApplicationModePolicyNames        []McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"applicationModePolicyNames,omitempty"`
	ApplicationRuleSetPolicyNameCount *uint                                                 `json:"applicationRuleSetPolicyNameCount,omitempty"`
	ApplicationRuleSetPolicyNames     []McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"applicationRuleSetPolicyNames,omitempty"`
	BlockedApplicationEventCount      *uint                                                 `json:"blockedApplicationEventCount,omitempty"`
	BlockedApplicationEvents          []McAfee_Epolicy_Orchestrator_Version45_Event         `json:"blockedApplicationEvents,omitempty"`
	EnforcementPolicyNameCount        *uint                                                 `json:"enforcementPolicyNameCount,omitempty"`
	EnforcementPolicyNames            []McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"enforcementPolicyNames,omitempty"`
	EpoVersion                        *string                                               `json:"epoVersion,omitempty"`
	FirewallModePolicyNameCount       *uint                                                 `json:"firewallModePolicyNameCount,omitempty"`
	FirewallModePolicyNames           []McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"firewallModePolicyNames,omitempty"`
	FirewallRuleSetPolicyNameCount    *uint                                                 `json:"firewallRuleSetPolicyNameCount,omitempty"`
	FirewallRuleSetPolicyNames        []McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"firewallRuleSetPolicyNames,omitempty"`
	IpsEventCount                     *uint                                                 `json:"ipsEventCount,omitempty"`
	IpsEvents                         []McAfee_Epolicy_Orchestrator_Version45_Event         `json:"ipsEvents,omitempty"`
	IpsModePolicyNameCount            *uint                                                 `json:"ipsModePolicyNameCount,omitempty"`
	IpsModePolicyNames                []McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"ipsModePolicyNames,omitempty"`
	IpsProtectionPolicyNameCount      *uint                                                 `json:"ipsProtectionPolicyNameCount,omitempty"`
	IpsProtectionPolicyNames          []McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"ipsProtectionPolicyNames,omitempty"`
	TransactionStatus                 *string                                               `json:"transactionStatus,omitempty"`
}

type Software_Component_HostIps_Mcafee_Epo_Version45_Hips_Version7 struct {
	Software_Component_HostIps_Mcafee_Epo_Version45_Hips
}

type Software_Component_HostIps_Mcafee_Epo_Version45_Hips_Version8 struct {
	Software_Component_HostIps_Mcafee_Epo_Version45_Hips
}

type Software_Component_OperatingSystem struct {
	Software_Component

	LicenseExpirationDate  *time.Time                               `json:"licenseExpirationDate,omitempty"`
	PartitionTemplateCount *uint                                    `json:"partitionTemplateCount,omitempty"`
	PartitionTemplates     []Hardware_Component_Partition_Template  `json:"partitionTemplates,omitempty"`
	ReloadTransactionGroup *Provisioning_Version1_Transaction_Group `json:"reloadTransactionGroup,omitempty"`
}

type Software_Component_Package struct {
	Software_Component
}

type Software_Component_Package_Management struct {
	Software_Component_Package
}

type Software_Component_Package_Management_Ksplice struct {
	Software_Component_Package_Management
}

type Software_Component_Password struct {
	Entity

	CreateDate  *time.Time          `json:"createDate,omitempty"`
	Id          *int                `json:"id,omitempty"`
	ModifyDate  *time.Time          `json:"modifyDate,omitempty"`
	Notes       *string             `json:"notes,omitempty"`
	Password    *string             `json:"password,omitempty"`
	Port        *int                `json:"port,omitempty"`
	Software    *Software_Component `json:"software,omitempty"`
	SoftwareId  *int                `json:"softwareId,omitempty"`
	SshKeyCount *uint               `json:"sshKeyCount,omitempty"`
	SshKeys     []Security_Ssh_Key  `json:"sshKeys,omitempty"`
	Username    *string             `json:"username,omitempty"`
}

type Software_Component_Password_History struct {
	Entity

	CreateDate          *time.Time          `json:"createDate,omitempty"`
	Notes               *string             `json:"notes,omitempty"`
	Password            *string             `json:"password,omitempty"`
	SoftwareComponent   *Software_Component `json:"softwareComponent,omitempty"`
	SoftwareComponentId *int                `json:"softwareComponentId,omitempty"`
	Username            *string             `json:"username,omitempty"`
}

type Software_Component_Security struct {
	Software_Component
}

type Software_Component_Security_SafeNet struct {
	Software_Component_Security
}

type Software_Description struct {
	Entity

	AttributeCount                     *uint                                    `json:"attributeCount,omitempty"`
	Attributes                         []Software_Description_Attribute         `json:"attributes,omitempty"`
	AverageInstallationDuration        *int                                     `json:"averageInstallationDuration,omitempty"`
	CompatibleSoftwareDescriptionCount *uint                                    `json:"compatibleSoftwareDescriptionCount,omitempty"`
	CompatibleSoftwareDescriptions     []Software_Description                   `json:"compatibleSoftwareDescriptions,omitempty"`
	ControlPanel                       *int                                     `json:"controlPanel,omitempty"`
	FeatureCount                       *uint                                    `json:"featureCount,omitempty"`
	Features                           []Software_Description_Feature           `json:"features,omitempty"`
	Id                                 *int                                     `json:"id,omitempty"`
	LatestVersion                      []Software_Description                   `json:"latestVersion,omitempty"`
	LatestVersionCount                 *uint                                    `json:"latestVersionCount,omitempty"`
	LicenseTermUnit                    *string                                  `json:"licenseTermUnit,omitempty"`
	LicenseTermValue                   *int                                     `json:"licenseTermValue,omitempty"`
	LongDescription                    *string                                  `json:"longDescription,omitempty"`
	Manufacturer                       *string                                  `json:"manufacturer,omitempty"`
	Name                               *string                                  `json:"name,omitempty"`
	OperatingSystem                    *int                                     `json:"operatingSystem,omitempty"`
	ProductItemCount                   *uint                                    `json:"productItemCount,omitempty"`
	ProductItems                       []Product_Item                           `json:"productItems,omitempty"`
	ProvisionTransactionGroup          *Provisioning_Version1_Transaction_Group `json:"provisionTransactionGroup,omitempty"`
	ReferenceCode                      *string                                  `json:"referenceCode,omitempty"`
	ReloadTransactionGroup             *Provisioning_Version1_Transaction_Group `json:"reloadTransactionGroup,omitempty"`
	RequiredUser                       *string                                  `json:"requiredUser,omitempty"`
	SoftwareLicenseCount               *uint                                    `json:"softwareLicenseCount,omitempty"`
	SoftwareLicenses                   []Software_License                       `json:"softwareLicenses,omitempty"`
	UpgradeSoftwareDescription         *Software_Description                    `json:"upgradeSoftwareDescription,omitempty"`
	UpgradeSoftwareDescriptionId       *int                                     `json:"upgradeSoftwareDescriptionId,omitempty"`
	UpgradeSwDesc                      *Software_Description                    `json:"upgradeSwDesc,omitempty"`
	UpgradeSwDescId                    *int                                     `json:"upgradeSwDescId,omitempty"`
	ValidFilesystemTypeCount           *uint                                    `json:"validFilesystemTypeCount,omitempty"`
	ValidFilesystemTypes               []Configuration_Storage_Filesystem_Type  `json:"validFilesystemTypes,omitempty"`
	Version                            *string                                  `json:"version,omitempty"`
	VirtualLicense                     *int                                     `json:"virtualLicense,omitempty"`
	VirtualizationPlatform             *int                                     `json:"virtualizationPlatform,omitempty"`
}

type Software_Description_Attribute struct {
	Entity

	SoftwareDescription *Software_Description                `json:"softwareDescription,omitempty"`
	Type                *Software_Description_Attribute_Type `json:"type,omitempty"`
	Value               *string                              `json:"value,omitempty"`
}

type Software_Description_Attribute_Type struct {
	Entity

	Keyname *string `json:"keyname,omitempty"`
}

type Software_Description_Feature struct {
	Entity

	Id      *int    `json:"id,omitempty"`
	KeyName *string `json:"keyName,omitempty"`
	Name    *string `json:"name,omitempty"`
	Vendor  *string `json:"vendor,omitempty"`
}

type Software_Description_RequiredUser struct {
	Entity

	DefaultPassword *string `json:"defaultPassword,omitempty"`
	Username        *string `json:"username,omitempty"`
}

type Software_License struct {
	Entity

	Account               *Account              `json:"account,omitempty"`
	Id                    *int                  `json:"id,omitempty"`
	Owner                 *Account              `json:"owner,omitempty"`
	SoftwareDescription   *Software_Description `json:"softwareDescription,omitempty"`
	SoftwareDescriptionId *int                  `json:"softwareDescriptionId,omitempty"`
}

type Software_VirtualLicense struct {
	Entity

	Account               *Account                  `json:"account,omitempty"`
	AccountId             *int                      `json:"accountId,omitempty"`
	BillingItem           *Billing_Item             `json:"billingItem,omitempty"`
	HostHardware          *Hardware_Server          `json:"hostHardware,omitempty"`
	HostHardwareId        *int                      `json:"hostHardwareId,omitempty"`
	Id                    *int                      `json:"id,omitempty"`
	IpAddress             *string                   `json:"ipAddress,omitempty"`
	IpAddressRecord       *Network_Subnet_IpAddress `json:"ipAddressRecord,omitempty"`
	Key                   *string                   `json:"key,omitempty"`
	Notes                 *string                   `json:"notes,omitempty"`
	SoftwareDescription   *Software_Description     `json:"softwareDescription,omitempty"`
	SoftwareDescriptionId *int                      `json:"softwareDescriptionId,omitempty"`
	Subnet                *Network_Subnet           `json:"subnet,omitempty"`
	SubnetId              *int                      `json:"subnetId,omitempty"`
}
