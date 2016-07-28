package softlayer

import (
	"time"

	"github.ibm.com/riethm/gopherlayer/datatypes"
)

type Hardware interface {
	Entity

	AllowAccessToNetworkStorage(networkStorageTemplateObject datatypes.Network_Storage) (bool, error)
	AllowAccessToNetworkStorageList(networkStorageTemplateObjects datatypes.Network_Storage) (bool, error)
	CaptureImage(captureTemplate datatypes.Container_Disk_Image_Capture_Template) (datatypes.Virtual_Guest_Block_Device_Template_Group, error)
	CloseAlarm(alarmId string) (bool, error)
	CreateObject(templateObject datatypes.Hardware) (datatypes.Hardware, error)
	DeleteObject() (bool, error)
	DeleteSoftwareComponentPasswords(softwareComponentPasswords datatypes.Software_Component_Password) (bool, error)
	EditSoftwareComponentPasswords(softwareComponentPasswords datatypes.Software_Component_Password) (bool, error)
	ExecuteRemoteScript(uri string) error
	FindByIpAddress(ipAddress string) (datatypes.Hardware, error)
	GenerateOrderTemplate(templateObject datatypes.Hardware) (datatypes.Container_Product_Order, error)
	GetAlarmHistory(startDate time.Time, endDate time.Time, alarmId string) (datatypes.Container_Monitoring_Alarm_History, error)
	GetAttachedNetworkStorages(nasType string) (datatypes.Network_Storage, error)
	GetAvailableNetworkStorages(nasType string) (datatypes.Network_Storage, error)
	GetBackendIncomingBandwidth(startDate time.Time, endDate time.Time) (float64, error)
	GetBackendOutgoingBandwidth(startDate time.Time, endDate time.Time) (float64, error)
	GetCreateObjectOptions() (datatypes.Container_Hardware_Configuration, error)
	GetCurrentBillingDetail() (datatypes.Billing_Item, error)
	GetCurrentBillingTotal() (float64, error)
	GetDailyAverage(startDate time.Time, endDate time.Time) (float64, error)
	GetFrontendIncomingBandwidth(startDate time.Time, endDate time.Time) (float64, error)
	GetFrontendOutgoingBandwidth(startDate time.Time, endDate time.Time) (float64, error)
	GetHourlyBandwidth(mode string, day time.Time) (datatypes.Metric_Tracking_Object_Data, error)
	GetMonitoringActiveAlarms(startDate time.Time, endDate time.Time) (datatypes.Container_Monitoring_Alarm_History, error)
	GetMonitoringClosedAlarms(startDate time.Time, endDate time.Time) (datatypes.Container_Monitoring_Alarm_History, error)
	GetObject() (datatypes.Hardware, error)
	GetPrivateBandwidthData(startTime int, endTime int) (datatypes.Metric_Tracking_Object_Data, error)
	GetPublicBandwidthData(startTime int, endTime int) (datatypes.Metric_Tracking_Object_Data, error)
	GetSensorData() (datatypes.Container_RemoteManagement_SensorReading, error)
	GetSensorDataWithGraphs() (datatypes.Container_RemoteManagement_SensorReadingsWithGraphs, error)
	GetServerFanSpeedGraphs() (datatypes.Container_RemoteManagement_Graphs_SensorSpeed, error)
	GetServerPowerState() (string, error)
	GetServerTemperatureGraphs() (datatypes.Container_RemoteManagement_Graphs_SensorTemperature, error)
	GetTransactionHistory() (datatypes.Provisioning_Version1_Transaction_History, error)
	GetUpgradeItemPrices() (datatypes.Product_Item_Price, error)
	ImportVirtualHost() (datatypes.Virtual_Host, error)
	IsPingable() (bool, error)
	Ping() (string, error)
	PowerCycle() (bool, error)
	PowerOff() (bool, error)
	PowerOn() (bool, error)
	RebootDefault() (bool, error)
	RebootHard() (bool, error)
	RebootSoft() (bool, error)
	RemoveAccessToNetworkStorage(networkStorageTemplateObject datatypes.Network_Storage) (bool, error)
	RemoveAccessToNetworkStorageList(networkStorageTemplateObjects datatypes.Network_Storage) (bool, error)
	SetTags(tags string) (bool, error)

	GetAccount() (datatypes.Account, error)
	GetActiveComponents() (datatypes.Hardware_Component, error)
	GetActiveNetworkMonitorIncident() (datatypes.Network_Monitor_Version1_Incident, error)
	GetAllPowerComponents() (datatypes.Hardware_Power_Component, error)
	GetAllowedHost() (datatypes.Network_Storage_Allowed_Host, error)
	GetAllowedNetworkStorage() (datatypes.Network_Storage, error)
	GetAllowedNetworkStorageReplicas() (datatypes.Network_Storage, error)
	GetAntivirusSpywareSoftwareComponent() (datatypes.Software_Component, error)
	GetAttributes() (datatypes.Hardware_Attribute, error)
	GetAverageDailyPublicBandwidthUsage() (float64, error)
	GetBackendNetworkComponents() (datatypes.Network_Component, error)
	GetBackendRouters() (datatypes.Hardware, error)
	GetBandwidthAllocation() (float64, error)
	GetBandwidthAllotmentDetail() (datatypes.Network_Bandwidth_Version1_Allotment_Detail, error)
	GetBenchmarkCertifications() (datatypes.Hardware_Benchmark_Certification, error)
	GetBillingItem() (datatypes.Billing_Item_Hardware, error)
	GetBillingItemFlag() (bool, error)
	GetBlockCancelBecauseDisconnectedFlag() (bool, error)
	GetBusinessContinuanceInsuranceFlag() (bool, error)
	GetComponents() (datatypes.Hardware_Component, error)
	GetContinuousDataProtectionSoftwareComponent() (datatypes.Software_Component, error)
	GetCurrentBillableBandwidthUsage() (float64, error)
	GetDatacenter() (datatypes.Location, error)
	GetDatacenterName() (string, error)
	GetDownlinkHardware() (datatypes.Hardware, error)
	GetDownlinkNetworkHardware() (datatypes.Hardware, error)
	GetDownlinkServers() (datatypes.Hardware, error)
	GetDownlinkVirtualGuests() (datatypes.Virtual_Guest, error)
	GetDownstreamHardwareBindings() (datatypes.Network_Component_Uplink_Hardware, error)
	GetDownstreamNetworkHardware() (datatypes.Hardware, error)
	GetDownstreamNetworkHardwareWithIncidents() (datatypes.Hardware, error)
	GetDownstreamServers() (datatypes.Hardware, error)
	GetDownstreamVirtualGuests() (datatypes.Virtual_Guest, error)
	GetDriveControllers() (datatypes.Hardware_Component, error)
	GetEvaultNetworkStorage() (datatypes.Network_Storage, error)
	GetFirewallServiceComponent() (datatypes.Network_Component_Firewall, error)
	GetFixedConfigurationPreset() (datatypes.Product_Package_Preset, error)
	GetFrontendNetworkComponents() (datatypes.Network_Component, error)
	GetFrontendRouters() (datatypes.Hardware, error)
	GetGlobalIdentifier() (string, error)
	GetHardDrives() (datatypes.Hardware_Component, error)
	GetHardwareChassis() (datatypes.Hardware_Chassis, error)
	GetHardwareFunction() (datatypes.Hardware_Function, error)
	GetHardwareFunctionDescription() (string, error)
	GetHardwareStatus() (datatypes.Hardware_Status, error)
	GetHasTrustedPlatformModuleBillingItemFlag() (bool, error)
	GetHostIpsSoftwareComponent() (datatypes.Software_Component, error)
	GetHourlyBillingFlag() (bool, error)
	GetInboundBandwidthUsage() (float64, error)
	GetInboundPublicBandwidthUsage() (float64, error)
	GetLastTransaction() (datatypes.Provisioning_Version1_Transaction, error)
	GetLatestNetworkMonitorIncident() (datatypes.Network_Monitor_Version1_Incident, error)
	GetLocation() (datatypes.Location, error)
	GetLocationPathString() (string, error)
	GetLockboxNetworkStorage() (datatypes.Network_Storage, error)
	GetManagedResourceFlag() (bool, error)
	GetMemory() (datatypes.Hardware_Component, error)
	GetMemoryCapacity() (uint, error)
	GetMetricTrackingObject() (datatypes.Metric_Tracking_Object_HardwareServer, error)
	GetMonitoringAgents() (datatypes.Monitoring_Agent, error)
	GetMonitoringRobot() (datatypes.Monitoring_Robot, error)
	GetMonitoringServiceComponent() (datatypes.Network_Monitor_Version1_Query_Host_Stratum, error)
	GetMonitoringServiceEligibilityFlag() (bool, error)
	GetMonitoringServiceFlag() (bool, error)
	GetMotherboard() (datatypes.Hardware_Component, error)
	GetNetworkCards() (datatypes.Hardware_Component, error)
	GetNetworkComponents() (datatypes.Network_Component, error)
	GetNetworkGatewayMember() (datatypes.Network_Gateway_Member, error)
	GetNetworkGatewayMemberFlag() (bool, error)
	GetNetworkManagementIpAddress() (string, error)
	GetNetworkMonitorAttachedDownHardware() (datatypes.Hardware, error)
	GetNetworkMonitorAttachedDownVirtualGuests() (datatypes.Virtual_Guest, error)
	GetNetworkMonitorIncidents() (datatypes.Network_Monitor_Version1_Incident, error)
	GetNetworkMonitors() (datatypes.Network_Monitor_Version1_Query_Host, error)
	GetNetworkStatus() (string, error)
	GetNetworkStatusAttribute() (datatypes.Hardware_Attribute, error)
	GetNetworkStorage() (datatypes.Network_Storage, error)
	GetNetworkVlans() (datatypes.Network_Vlan, error)
	GetNextBillingCycleBandwidthAllocation() (float64, error)
	GetNotesHistory() (datatypes.Hardware_Note, error)
	GetOperatingSystem() (datatypes.Software_Component_OperatingSystem, error)
	GetOperatingSystemReferenceCode() (string, error)
	GetOutboundBandwidthUsage() (float64, error)
	GetOutboundPublicBandwidthUsage() (float64, error)
	GetPointOfPresenceLocation() (datatypes.Location, error)
	GetPowerComponents() (datatypes.Hardware_Power_Component, error)
	GetPowerSupply() (datatypes.Hardware_Component, error)
	GetPrimaryBackendIpAddress() (string, error)
	GetPrimaryBackendNetworkComponent() (datatypes.Network_Component, error)
	GetPrimaryIpAddress() (string, error)
	GetPrimaryNetworkComponent() (datatypes.Network_Component, error)
	GetPrivateNetworkOnlyFlag() (bool, error)
	GetProcessorCoreAmount() (uint, error)
	GetProcessorPhysicalCoreAmount() (uint, error)
	GetProcessors() (datatypes.Hardware_Component, error)
	GetRack() (datatypes.Location, error)
	GetRaidControllers() (datatypes.Hardware_Component, error)
	GetRecentEvents() (datatypes.Notification_Occurrence_Event, error)
	GetRemoteManagementAccounts() (datatypes.Hardware_Component_RemoteManagement_User, error)
	GetRemoteManagementComponent() (datatypes.Network_Component, error)
	GetResourceGroupMemberReferences() (datatypes.Resource_Group_Member, error)
	GetResourceGroupRoles() (datatypes.Resource_Group_Role, error)
	GetResourceGroups() (datatypes.Resource_Group, error)
	GetRouters() (datatypes.Hardware, error)
	GetScaleAssets() (datatypes.Scale_Asset, error)
	GetSecurityScanRequests() (datatypes.Network_Security_Scanner_Request, error)
	GetServerRoom() (datatypes.Location, error)
	GetServiceProvider() (datatypes.Service_Provider, error)
	GetSoftwareComponents() (datatypes.Software_Component, error)
	GetSparePoolBillingItem() (datatypes.Billing_Item_Hardware, error)
	GetSshKeys() (datatypes.Security_Ssh_Key, error)
	GetStorageNetworkComponents() (datatypes.Network_Component, error)
	GetTagReferences() (datatypes.Tag_Reference, error)
	GetTopLevelLocation() (datatypes.Location, error)
	GetUpgradeRequest() (datatypes.Product_Upgrade_Request, error)
	GetUplinkHardware() (datatypes.Hardware, error)
	GetUplinkNetworkComponents() (datatypes.Network_Component, error)
	GetUserData() (datatypes.Hardware_Attribute, error)
	GetVirtualChassis() (datatypes.Hardware_Group, error)
	GetVirtualChassisSiblings() (datatypes.Hardware, error)
	GetVirtualHost() (datatypes.Virtual_Host, error)
	GetVirtualLicenses() (datatypes.Software_VirtualLicense, error)
	GetVirtualRack() (datatypes.Network_Bandwidth_Version1_Allotment, error)
	GetVirtualRackId() (int, error)
	GetVirtualRackName() (string, error)
	GetVirtualizationPlatform() (datatypes.Software_Component, error)
}

type Hardware_Attribute interface {
	Entity

	GetHardwareAttributeType() (datatypes.Hardware_Attribute_Type, error)
}

type Hardware_Attribute_Type interface {
	Entity
}

type Hardware_Benchmark_Certification interface {
	Entity

	GetObject() (datatypes.Hardware_Benchmark_Certification, error)
	GetResultFile() ([]byte, error)

	GetAccount() (datatypes.Account, error)
	GetHardware() (datatypes.Hardware, error)
}

type Hardware_Chassis interface {
	Entity

	GetBackplaneCapacity() (string, error)
	GetBayCapacity() (string, error)
	GetDriveCapacity() (string, error)
	GetDriveControllerCapacity() (string, error)
	GetGpuCapacity() (string, error)
	GetHardwareFunction() (datatypes.Hardware_Function, error)
	GetPowerCapacity() (string, error)
}

type Hardware_Component interface {
	Entity

	GetCapacity() (float64, error)
	GetChildren() (datatypes.Hardware_Component, error)
	GetDownlinkHardwareComponents() (datatypes.Hardware_Component, error)
	GetHardware() (datatypes.Hardware, error)
	GetHardwareComponentModel() (datatypes.Hardware_Component_Model, error)
	GetHardwareComponentType() (datatypes.Hardware_Component_Type, error)
	GetNetworkComponents() (datatypes.Network_Component, error)
	GetOwner() (datatypes.Account, error)
	GetParent() (datatypes.Hardware_Component, error)
	GetRaidMode() (string, error)
	GetServiceProvider() (datatypes.Service_Provider, error)
	GetUplinkHardwareComponents() (datatypes.Hardware_Component, error)
}

type Hardware_Component_Attribute interface {
	Entity

	GetHardwareComponent() (datatypes.Hardware_Component, error)
	GetHardwareComponentAttributeType() (datatypes.Hardware_Component_Attribute_Type, error)
}

type Hardware_Component_Attribute_Type interface {
	Entity
}

type Hardware_Component_DriveController interface {
	Hardware_Component
}

type Hardware_Component_HardDrive interface {
	Hardware_Component

	GetPartitions() (datatypes.Hardware_Component_Partition, error)
}

type Hardware_Component_Model interface {
	Entity

	GetObject() (datatypes.Hardware_Component_Model, error)

	GetArchitectureType() (datatypes.Hardware_Component_Model_Architecture_Type, error)
	GetAttributes() (datatypes.Hardware_Component_Model_Attribute, error)
	GetCompatibleArrayTypes() (datatypes.Configuration_Storage_Group_Array_Type, error)
	GetCompatibleChildComponentModels() (datatypes.Hardware_Component_Model, error)
	GetCompatibleParentComponentModels() (datatypes.Hardware_Component_Model, error)
	GetHardwareComponents() (datatypes.Hardware_Component, error)
	GetHardwareGenericComponentModel() (datatypes.Hardware_Component_Model_Generic, error)
	GetInfinibandCompatibleAttribute() (datatypes.Hardware_Component_Model_Attribute, error)
	GetIsInfinibandCompatible() (bool, error)
	GetRebootTime() (datatypes.Hardware_Component_Motherboard_Reboot_Time, error)
	GetType() (string, error)
	GetValidAttributeTypes() (datatypes.Hardware_Component_Model_Attribute_Type, error)
}

type Hardware_Component_Model_Architecture_Type interface {
	Entity

	GetChildren() (datatypes.Hardware_Component_Model_Architecture_Type, error)
	GetParent() (datatypes.Hardware_Component_Model_Architecture_Type, error)
}

type Hardware_Component_Model_Attribute interface {
	Entity

	GetHardwareComponent() (datatypes.Hardware_Component_Model, error)
	GetHardwareComponentAttributeType() (datatypes.Hardware_Component_Model_Attribute_Type, error)
}

type Hardware_Component_Model_Attribute_Type interface {
	Entity

	GetValidComponentTypes() (datatypes.Hardware_Component_Type, error)
}

type Hardware_Component_Model_Generic interface {
	Entity

	GetHardwareComponentModels() (datatypes.Hardware_Component_Model, error)
	GetHardwareComponentType() (datatypes.Hardware_Component_Type, error)
	GetMarketingFeatures() (datatypes.Hardware_Component_Model_Generic_MarketingFeature, error)
}

type Hardware_Component_Model_Generic_Attribute interface {
	Entity

	GetHardwareGenericComponentModel() (datatypes.Hardware_Component_Model_Generic, error)
}

type Hardware_Component_Model_Generic_MarketingFeature interface {
	Entity

	GetHardwareGenericComponentModel() (datatypes.Hardware_Component_Model_Generic, error)
}

type Hardware_Component_Motherboard interface {
	Hardware_Component
}

type Hardware_Component_Motherboard_Reboot_Time interface {
	Entity

	GetHardwareComponentModel() (datatypes.Hardware_Component_Model, error)
}

type Hardware_Component_NetworkCard interface {
	Hardware_Component
}

type Hardware_Component_Partition interface {
	Entity

	GetHardwareComponent() (datatypes.Hardware_Component, error)
}

type Hardware_Component_Partition_OperatingSystem interface {
	Entity

	GetAllObjects() (datatypes.Hardware_Component_Partition_OperatingSystem, error)
	GetByDescription(description string) (datatypes.Hardware_Component_Partition_OperatingSystem, error)
	GetObject() (datatypes.Hardware_Component_Partition_OperatingSystem, error)

	GetPartitionTemplates() (datatypes.Hardware_Component_Partition_Template, error)
}

type Hardware_Component_Partition_Template interface {
	Entity

	GetObject() (datatypes.Hardware_Component_Partition_Template, error)

	GetAccount() (datatypes.Account, error)
	GetData() (datatypes.Hardware_Component_Partition_Template_Partition, error)
	GetExpireDate() (string, error)
	GetPartitionOperatingSystem() (datatypes.Hardware_Component_Partition_OperatingSystem, error)
	GetPartitionTemplatePartition() (datatypes.Hardware_Component_Partition_Template_Partition, error)
}

type Hardware_Component_Partition_Template_Partition interface {
	Entity

	GetFilesystemType() (datatypes.Configuration_Storage_Filesystem_Type, error)
	GetPartitionTemplate() (datatypes.Hardware_Component_Partition_Template, error)
}

type Hardware_Component_Processor interface {
	Hardware_Component
}

type Hardware_Component_Ram interface {
	Hardware_Component
}

type Hardware_Component_RemoteManagement interface {
	Hardware_Component

	GetNetworkComponent() (datatypes.Network_Component, error)
}

type Hardware_Component_RemoteManagement_Command interface {
	Entity

	GetRequests() (datatypes.Hardware_Component_RemoteManagement_Command_Request, error)
}

type Hardware_Component_RemoteManagement_Command_Request interface {
	Entity

	GetHardware() (datatypes.Hardware, error)
	GetNetworkComponent() (datatypes.Network_Component, error)
	GetRemoteManagementCommand() (datatypes.Hardware_Component_RemoteManagement_Command, error)
	GetUser() (datatypes.User_Customer, error)
}

type Hardware_Component_RemoteManagement_User interface {
	Entity

	GetHardware() (datatypes.Hardware, error)
	GetNetworkComponent() (datatypes.Network_Component, error)
}

type Hardware_Component_SecurityDevice interface {
	Hardware_Component
}

type Hardware_Component_SecurityDevice_Infineon interface {
	Hardware_Component_SecurityDevice
}

type Hardware_Component_Type interface {
	Entity

	GetHardwareGenericComponentModels() (datatypes.Hardware_Component_Model_Generic, error)
	GetTypeParent() (datatypes.Hardware_Component_Type, error)
}

type Hardware_Firewall interface {
	Hardware_Switch

	GetUsers() (datatypes.User_Customer, error)
}

type Hardware_Function interface {
	Entity
}

type Hardware_Group interface {
	Entity

	GetDownlinkServers() (datatypes.Hardware, error)
	GetDownlinkVirtualGuests() (datatypes.Virtual_Guest, error)
	GetDownstreamNetworkHardware() (datatypes.Hardware, error)
	GetDownstreamNetworkHardwareWithIncidents() (datatypes.Hardware, error)
	GetHardwareChassis() (datatypes.Hardware_Chassis, error)
	GetNetworkMonitorAttachedDownHardware() (datatypes.Hardware, error)
	GetNetworkMonitorAttachedDownVirtualGuests() (datatypes.Virtual_Guest, error)
	GetNetworkStatus() (string, error)
}

type Hardware_LoadBalancer interface {
	Hardware

	GetModelFamily() (string, error)
	GetUsers() (datatypes.User_Customer, error)
}

type Hardware_Note interface {
	Entity

	GetEmployee() (datatypes.User_Employee, error)
	GetHardware() (datatypes.Hardware, error)
	GetType() (datatypes.Hardware_Note_Type, error)
	GetUser() (datatypes.User_Customer, error)
}

type Hardware_Note_Type interface {
	Entity
}

type Hardware_Power_Component interface {
	Entity

	GetHardware() (datatypes.Hardware, error)
}

type Hardware_Router interface {
	Hardware_Switch

	GetObject() (datatypes.Hardware_Router, error)

	GetBoundSubnets() (datatypes.Network_Subnet, error)
	GetLocalDiskStorageCapabilityFlag() (bool, error)
	GetSanStorageCapabilityFlag() (bool, error)
}

type Hardware_Router_Backend interface {
	Hardware_Router
}

type Hardware_Router_Frontend interface {
	Hardware_Router
}

type Hardware_SecurityModule interface {
	Hardware_Server

	CreateObject(templateObject datatypes.Hardware_SecurityModule) (datatypes.Hardware_SecurityModule, error)
	GetObject() (datatypes.Hardware_SecurityModule, error)
}

type Hardware_Server interface {
	Hardware

	ActivatePrivatePort() (bool, error)
	ActivatePublicPort() (bool, error)
	BootToRescueLayer(noOsBootEnvironment string) (bool, error)
	CreateFirmwareUpdateTransaction(ipmi int, raidController int, bios int, harddrive int) (bool, error)
	CreateObject(templateObject datatypes.Hardware_Server) (datatypes.Hardware_Server, error)
	CreatePostSoftwareInstallTransaction(installCodes string, returnBoolean bool) (bool, error)
	EditObject(templateObject datatypes.Hardware_Server) (bool, error)
	GetBackendBandwidthUsage(startDate time.Time, endDate time.Time) (datatypes.Metric_Tracking_Object_Data, error)
	GetBackendBandwidthUse(startDate time.Time, endDate time.Time) (datatypes.Network_Bandwidth_Version1_Usage_Detail, error)
	GetBandwidthForDateRange(startDate time.Time, endDate time.Time) (datatypes.Metric_Tracking_Object_Data, error)
	GetBandwidthImage(networkType string, snapshotRange string, draw bool, dateSpecified time.Time, dateSpecifiedEnd time.Time) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetCurrentBenchmarkCertificationResultFile() ([]byte, error)
	GetCustomBandwidthDataByDate(graphData datatypes.Container_Graph) (datatypes.Container_Graph, error)
	GetFirewallProtectableSubnets() (datatypes.Network_Subnet, error)
	GetFrontendBandwidthUsage(startDate time.Time, endDate time.Time) (datatypes.Metric_Tracking_Object_Data, error)
	GetFrontendBandwidthUse(startDate time.Time, endDate time.Time) (datatypes.Network_Bandwidth_Version1_Usage_Detail, error)
	GetHardwareByIpAddress(ipAddress string) (datatypes.Hardware_Server, error)
	GetItemPricesFromSoftwareDescriptions(softwareDescriptions datatypes.Software_Description, includeTranslationsFlag bool, returnAllPricesFlag bool) (datatypes.Product_Item, error)
	GetManagementNetworkComponent() (datatypes.Network_Component, error)
	GetNetworkComponentFirewallProtectableIpAddresses() (datatypes.Network_Subnet_IpAddress, error)
	GetObject() (datatypes.Hardware_Server, error)
	GetPMInfo() (datatypes.Container_RemoteManagement_PmInfo, error)
	GetPrimaryDriveSize() (int, error)
	GetPrivateBandwidthDataSummary() (datatypes.Container_Network_Bandwidth_Data_Summary, error)
	GetPrivateBandwidthGraphImage(startTime string, endTime string) ([]byte, error)
	GetPrivateNetworkComponent() (datatypes.Network_Component, error)
	GetPrivateVlan() (datatypes.Network_Vlan, error)
	GetPrivateVlanByIpAddress(ipAddress string) (datatypes.Network_Vlan, error)
	GetProvisionDate() (time.Time, error)
	GetPublicBandwidthDataSummary() (datatypes.Container_Network_Bandwidth_Data_Summary, error)
	GetPublicBandwidthGraphImage(startTime time.Time, endTime time.Time) ([]byte, error)
	GetPublicBandwidthTotal(startTime int, endTime int) (uint, error)
	GetPublicNetworkComponent() (datatypes.Network_Component, error)
	GetPublicVlan() (datatypes.Network_Vlan, error)
	GetPublicVlanByHostname(hostname string) (datatypes.Network_Vlan, error)
	GetReverseDomainRecords() (datatypes.Dns_Domain, error)
	GetSensorData() (datatypes.Container_RemoteManagement_SensorReading, error)
	GetSensorDataWithGraphs() (datatypes.Container_RemoteManagement_SensorReadingsWithGraphs, error)
	GetServerDetails() (datatypes.Container_Hardware_Server_Details, error)
	GetServerFanSpeedGraphs() (datatypes.Container_RemoteManagement_Graphs_SensorSpeed, error)
	GetServerPowerState() (string, error)
	GetServerTemperatureGraphs() (datatypes.Container_RemoteManagement_Graphs_SensorTemperature, error)
	GetValidBlockDeviceTemplateGroups(visibility string) (datatypes.Virtual_Guest_Block_Device_Template_Group, error)
	GetWindowsUpdateAvailableUpdates() (datatypes.Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem, error)
	GetWindowsUpdateInstalledUpdates() (datatypes.Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem, error)
	GetWindowsUpdateStatus() (datatypes.Container_Utility_Microsoft_Windows_UpdateServices_Status, error)
	InitiateIderaBareMetalRestore() (bool, error)
	InitiateR1SoftBareMetalRestore() (bool, error)
	IsBackendPingable() (bool, error)
	IsPingable() (bool, error)
	IsWindowsServer() (bool, error)
	Ping() (string, error)
	PowerCycle() (bool, error)
	PowerOff() (bool, error)
	PowerOn() (bool, error)
	RebootDefault() (bool, error)
	RebootHard() (bool, error)
	RebootSoft() (bool, error)
	ReloadCurrentOperatingSystemConfiguration(token string) (string, error)
	ReloadOperatingSystem(token string, config datatypes.Container_Hardware_Server_Configuration) (string, error)
	RunPassmarkCertificationBenchmark() (bool, error)
	SetOperatingSystemPassword(newPassword string) (bool, error)
	SetPrivateNetworkInterfaceSpeed(newSpeed int) (bool, error)
	SetPublicNetworkInterfaceSpeed(newSpeed int) (bool, error)
	SetUserMetadata(metadata string) (datatypes.Hardware_Attribute, error)
	ShutdownPrivatePort() (bool, error)
	ShutdownPublicPort() (bool, error)
	SparePool(action string, newOrder bool) (bool, error)
	ValidatePartitionsForOperatingSystem(operatingSystem datatypes.Software_Description, partitions datatypes.Hardware_Component_Partition) (bool, error)

	GetActiveNetworkFirewallBillingItem() (datatypes.Billing_Item, error)
	GetActiveTickets() (datatypes.Ticket, error)
	GetActiveTransaction() (datatypes.Provisioning_Version1_Transaction, error)
	GetActiveTransactions() (datatypes.Provisioning_Version1_Transaction, error)
	GetAvailableMonitoring() (datatypes.Network_Monitor_Version1_Query_Host_Stratum, error)
	GetAverageDailyBandwidthUsage() (float64, error)
	GetAverageDailyPrivateBandwidthUsage() (float64, error)
	GetBillingCycleBandwidthUsage() (datatypes.Network_Bandwidth_Usage, error)
	GetBillingCyclePrivateBandwidthUsage() (datatypes.Network_Bandwidth_Usage, error)
	GetBillingCyclePublicBandwidthUsage() (datatypes.Network_Bandwidth_Usage, error)
	GetContainsSolidStateDrivesFlag() (bool, error)
	GetControlPanel() (datatypes.Software_Component_ControlPanel, error)
	GetCost() (float64, error)
	GetCurrentBandwidthSummary() (datatypes.Metric_Tracking_Object_Bandwidth_Summary, error)
	GetCustomerInstalledOperatingSystemFlag() (bool, error)
	GetCustomerOwnedFlag() (bool, error)
	GetInboundPrivateBandwidthUsage() (float64, error)
	GetLastOperatingSystemReload() (datatypes.Provisioning_Version1_Transaction, error)
	GetMetricTrackingObjectId() (int, error)
	GetMonitoringUserNotification() (datatypes.User_Customer_Notification_Hardware, error)
	GetOpenCancellationTicket() (datatypes.Ticket, error)
	GetOutboundPrivateBandwidthUsage() (float64, error)
	GetOverBandwidthAllocationFlag() (int, error)
	GetPrivateIpAddress() (string, error)
	GetProjectedOverBandwidthAllocationFlag() (int, error)
	GetProjectedPublicBandwidthUsage() (float64, error)
	GetRecentRemoteManagementCommands() (datatypes.Hardware_Component_RemoteManagement_Command_Request, error)
	GetRegionalInternetRegistry() (datatypes.Network_Regional_Internet_Registry, error)
	GetRemoteManagement() (datatypes.Hardware_Component_RemoteManagement, error)
	GetRemoteManagementUsers() (datatypes.Hardware_Component_RemoteManagement_User, error)
	GetStatisticsRemoteManagement() (datatypes.Hardware_Component_RemoteManagement, error)
	GetUsers() (datatypes.User_Customer, error)
	GetVirtualGuests() (datatypes.Virtual_Guest, error)
}

type Hardware_Status interface {
	Entity
}

type Hardware_Switch interface {
	Hardware
}
