package softlayer

import (
	"time"

	"github.ibm.com/riethm/gopherlayer/datatypes"
)

type Software_AccountLicense interface {
	Entity

	GetAllObjects() (datatypes.Software_AccountLicense, error)
	GetObject() (datatypes.Software_AccountLicense, error)

	GetAccount() (datatypes.Account, error)
	GetBillingItem() (datatypes.Billing_Item, error)
	GetSoftwareDescription() (datatypes.Software_Description, error)
}

type Software_Component interface {
	Entity

	GetLicenseFile() (string, error)
	GetObject() (datatypes.Software_Component, error)
	GetVendorSetUpConfiguration() (string, error)

	GetAverageInstallationDuration() (uint, error)
	GetBillingItem() (datatypes.Billing_Item, error)
	GetHardware() (datatypes.Hardware, error)
	GetPasswordHistory() (datatypes.Software_Component_Password_History, error)
	GetPasswords() (datatypes.Software_Component_Password, error)
	GetSoftwareDescription() (datatypes.Software_Description, error)
	GetSoftwareLicense() (datatypes.Software_License, error)
	GetVirtualGuest() (datatypes.Virtual_Guest, error)
}

type Software_Component_Analytics interface {
	Software_Component
}

type Software_Component_Analytics_Urchin interface {
	Software_Component_Analytics
}

type Software_Component_AntivirusSpyware interface {
	Software_Component

	GetObject() (datatypes.Software_Component_AntivirusSpyware, error)
	UpdateAntivirusSpywarePolicy(newPolicy string, enforce bool) (bool, error)
}

type Software_Component_AntivirusSpyware_Mcafee interface {
	Software_Component_AntivirusSpyware
}

type Software_Component_AntivirusSpyware_Mcafee_Epo_Version36 interface {
	Software_Component_AntivirusSpyware_Mcafee

	GetAgentDetails() (McAfee_Epolicy_Orchestrator_Version36_Agent_Details, error)
	GetCurrentAntivirusPolicy() (int, error)
	GetDataFileVersion() (McAfee_Epolicy_Orchestrator_Version36_Product_Properties, error)
	GetLatestAccessProtectionEvents() (McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_AccessProtection, error)
	GetLatestAntivirusEvents() (McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event, error)
	GetLatestSpywareEvents() (McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event, error)
	GetTransactionStatus() (string, error)
}

type Software_Component_AntivirusSpyware_Mcafee_Epo_Version45 interface {
	Software_Component_AntivirusSpyware_Mcafee

	GetAgentDetails() (McAfee_Epolicy_Orchestrator_Version45_Agent_Details, error)
	GetCurrentAntivirusPolicy() (int, error)
	GetDataFileVersion() (McAfee_Epolicy_Orchestrator_Version45_Product_Properties, error)
	GetLatestAccessProtectionEvents() (McAfee_Epolicy_Orchestrator_Version45_Event, error)
	GetLatestAntivirusEvents() (McAfee_Epolicy_Orchestrator_Version45_Event, error)
	GetLatestSpywareEvents() (McAfee_Epolicy_Orchestrator_Version45_Event, error)
	GetTransactionStatus() (string, error)
}

type Software_Component_ControlPanel interface {
	Software_Component
}

type Software_Component_ControlPanel_Cpanel interface {
	Software_Component
}

type Software_Component_ControlPanel_Idera interface {
	Software_Component
}

type Software_Component_ControlPanel_Idera_ServerBackup interface {
	Software_Component_ControlPanel_Idera
}

type Software_Component_ControlPanel_Microsoft interface {
	Software_Component
}

type Software_Component_ControlPanel_Microsoft_WebPlatform interface {
	Software_Component_ControlPanel_Microsoft
}

type Software_Component_ControlPanel_Parallels interface {
	Software_Component
}

type Software_Component_ControlPanel_Parallels_Plesk interface {
	Software_Component_ControlPanel_Parallels
}

type Software_Component_ControlPanel_R1soft interface {
	Software_Component
}

type Software_Component_ControlPanel_R1soft_Cdp interface {
	Software_Component_ControlPanel_R1soft
}

type Software_Component_ControlPanel_R1soft_ServerBackup interface {
	Software_Component_ControlPanel_R1soft
}

type Software_Component_ControlPanel_Swsoft interface {
	Software_Component
}

type Software_Component_ControlPanel_WebhostAutomation interface {
	Software_Component
}

type Software_Component_HostIps interface {
	Software_Component

	GetCurrentHostIpsPolicies() (datatypes.Container_Software_Component_HostIps_Policy, error)
	GetObject() (datatypes.Software_Component_HostIps, error)
	UpdateHipsPolicies(newIpsMode string, newIpsProtection string, newFirewallMode string, newFirewallRuleset string, newApplicationMode string, newApplicationRuleset string, newEnforcementPolicy string) (bool, error)
}

type Software_Component_HostIps_Mcafee interface {
	Software_Component_HostIps
}

type Software_Component_HostIps_Mcafee_Epo_Version36_Hips interface {
	Software_Component_HostIps_Mcafee

	GetAgentDetails() (McAfee_Epolicy_Orchestrator_Version36_Agent_Details, error)
	GetApplicationModePolicyNames() (McAfee_Epolicy_Orchestrator_Version36_Policy_Object, error)
	GetApplicationRuleSetPolicyNames() (McAfee_Epolicy_Orchestrator_Version36_Policy_Object, error)
	GetEnforcementPolicyNames() (McAfee_Epolicy_Orchestrator_Version36_Policy_Object, error)
	GetFirewallModePolicyNames() (McAfee_Epolicy_Orchestrator_Version36_Policy_Object, error)
	GetFirewallRuleSetPolicyNames() (McAfee_Epolicy_Orchestrator_Version36_Policy_Object, error)
	GetIpsModePolicyNames() (McAfee_Epolicy_Orchestrator_Version36_Policy_Object, error)
	GetIpsProtectionPolicyNames() (McAfee_Epolicy_Orchestrator_Version36_Policy_Object, error)
	GetTransactionStatus() (string, error)
}

type Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version6 interface {
	Software_Component_HostIps_Mcafee_Epo_Version36_Hips

	GetBlockedApplicationEvents() (McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_BlockedApplicationEvent, error)
	GetIpsEvents() (McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_IPSEvent, error)
}

type Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version7 interface {
	Software_Component_HostIps_Mcafee_Epo_Version36_Hips

	GetBlockedApplicationEvents() (McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_BlockedApplicationEvent, error)
	GetIpsEvents() (McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent, error)
}

type Software_Component_HostIps_Mcafee_Epo_Version45_Hips interface {
	Software_Component_HostIps_Mcafee

	GetAgentDetails() (McAfee_Epolicy_Orchestrator_Version45_Agent_Details, error)
	GetApplicationModePolicyNames() (McAfee_Epolicy_Orchestrator_Version45_Policy_Object, error)
	GetApplicationRuleSetPolicyNames() (McAfee_Epolicy_Orchestrator_Version45_Policy_Object, error)
	GetBlockedApplicationEvents() (McAfee_Epolicy_Orchestrator_Version45_Event, error)
	GetEnforcementPolicyNames() (McAfee_Epolicy_Orchestrator_Version45_Policy_Object, error)
	GetFirewallModePolicyNames() (McAfee_Epolicy_Orchestrator_Version45_Policy_Object, error)
	GetFirewallRuleSetPolicyNames() (McAfee_Epolicy_Orchestrator_Version45_Policy_Object, error)
	GetIpsEvents() (McAfee_Epolicy_Orchestrator_Version45_Event, error)
	GetIpsModePolicyNames() (McAfee_Epolicy_Orchestrator_Version45_Policy_Object, error)
	GetIpsProtectionPolicyNames() (McAfee_Epolicy_Orchestrator_Version45_Policy_Object, error)
	GetTransactionStatus() (string, error)
}

type Software_Component_HostIps_Mcafee_Epo_Version45_Hips_Version7 interface {
	Software_Component_HostIps_Mcafee_Epo_Version45_Hips
}

type Software_Component_HostIps_Mcafee_Epo_Version45_Hips_Version8 interface {
	Software_Component_HostIps_Mcafee_Epo_Version45_Hips
}

type Software_Component_OperatingSystem interface {
	Software_Component

	GetLicenseExpirationDate() (time.Time, error)
	GetPartitionTemplates() (datatypes.Hardware_Component_Partition_Template, error)
	GetReloadTransactionGroup() (datatypes.Provisioning_Version1_Transaction_Group, error)
}

type Software_Component_Package interface {
	Software_Component
}

type Software_Component_Package_Management interface {
	Software_Component_Package
}

type Software_Component_Package_Management_Ksplice interface {
	Software_Component_Package_Management
}

type Software_Component_Password interface {
	Entity

	CreateObject(templateObject datatypes.Software_Component_Password) (datatypes.Software_Component_Password, error)
	CreateObjects(templateObjects datatypes.Software_Component_Password) (bool, error)
	DeleteObject() (bool, error)
	DeleteObjects(templateObjects datatypes.Software_Component_Password) (bool, error)
	EditObject(templateObject datatypes.Software_Component_Password) (bool, error)
	EditObjects(templateObjects datatypes.Software_Component_Password) (bool, error)
	GetObject() (datatypes.Software_Component_Password, error)

	GetSoftware() (datatypes.Software_Component, error)
	GetSshKeys() (datatypes.Security_Ssh_Key, error)
}

type Software_Component_Password_History interface {
	Entity

	GetSoftwareComponent() (datatypes.Software_Component, error)
}

type Software_Component_Security interface {
	Software_Component
}

type Software_Component_Security_SafeNet interface {
	Software_Component_Security
}

type Software_Description interface {
	Entity

	GetAllObjects() (datatypes.Software_Description, error)
	GetObject() (datatypes.Software_Description, error)

	GetAttributes() (datatypes.Software_Description_Attribute, error)
	GetAverageInstallationDuration() (int, error)
	GetCompatibleSoftwareDescriptions() (datatypes.Software_Description, error)
	GetFeatures() (datatypes.Software_Description_Feature, error)
	GetLatestVersion() (datatypes.Software_Description, error)
	GetProductItems() (datatypes.Product_Item, error)
	GetProvisionTransactionGroup() (datatypes.Provisioning_Version1_Transaction_Group, error)
	GetReloadTransactionGroup() (datatypes.Provisioning_Version1_Transaction_Group, error)
	GetRequiredUser() (string, error)
	GetSoftwareLicenses() (datatypes.Software_License, error)
	GetUpgradeSoftwareDescription() (datatypes.Software_Description, error)
	GetUpgradeSwDesc() (datatypes.Software_Description, error)
	GetValidFilesystemTypes() (datatypes.Configuration_Storage_Filesystem_Type, error)
}

type Software_Description_Attribute interface {
	Entity

	GetSoftwareDescription() (datatypes.Software_Description, error)
	GetType() (datatypes.Software_Description_Attribute_Type, error)
}

type Software_Description_Attribute_Type interface {
	Entity
}

type Software_Description_Feature interface {
	Entity
}

type Software_Description_RequiredUser interface {
	Entity
}

type Software_License interface {
	Entity

	GetAccount() (datatypes.Account, error)
	GetOwner() (datatypes.Account, error)
	GetSoftwareDescription() (datatypes.Software_Description, error)
}

type Software_VirtualLicense interface {
	Entity

	GetLicenseFile() ([]byte, error)
	GetObject() (datatypes.Software_VirtualLicense, error)

	GetAccount() (datatypes.Account, error)
	GetBillingItem() (datatypes.Billing_Item, error)
	GetHostHardware() (datatypes.Hardware_Server, error)
	GetIpAddressRecord() (datatypes.Network_Subnet_IpAddress, error)
	GetSoftwareDescription() (datatypes.Software_Description, error)
	GetSubnet() (datatypes.Network_Subnet, error)
}
