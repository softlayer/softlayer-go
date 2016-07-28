package softlayer

import (
	"time"

	"github.ibm.com/riethm/gopherlayer/datatypes"
)

type Virtual_Disk_Image interface {
	Entity

	EditObject(templateObject datatypes.Virtual_Disk_Image) (bool, error)
	GetObject() (datatypes.Virtual_Disk_Image, error)
	GetPublicIsoImages() (datatypes.Virtual_Disk_Image, error)

	GetBillingItem() (datatypes.Billing_Item_Virtual_Disk_Image, error)
	GetBlockDevices() (datatypes.Virtual_Guest_Block_Device, error)
	GetBootableVolumeFlag() (bool, error)
	GetCoalescedDiskImages() (datatypes.Virtual_Disk_Image, error)
	GetCopyOnWriteFlag() (bool, error)
	GetLocalDiskFlag() (bool, error)
	GetMetadataFlag() (bool, error)
	GetSoftwareReferences() (datatypes.Virtual_Disk_Image_Software, error)
	GetSourceDiskImage() (datatypes.Virtual_Disk_Image, error)
	GetStorageRepository() (datatypes.Virtual_Storage_Repository, error)
	GetStorageRepositoryType() (datatypes.Virtual_Storage_Repository_Type, error)
	GetTemplateBlockDevice() (datatypes.Virtual_Guest_Block_Device_Template, error)
	GetType() (datatypes.Virtual_Disk_Image_Type, error)
}

type Virtual_Disk_Image_Software interface {
	Entity

	GetDiskImage() (datatypes.Virtual_Disk_Image, error)
	GetPasswords() (datatypes.Virtual_Disk_Image_Software_Password, error)
	GetSoftwareDescription() (datatypes.Software_Description, error)
}

type Virtual_Disk_Image_Software_Password interface {
	Entity

	GetSoftware() (datatypes.Virtual_Disk_Image_Software, error)
}

type Virtual_Disk_Image_Type interface {
	Entity
}

type Virtual_Guest interface {
	Entity

	ActivatePrivatePort() (bool, error)
	ActivatePublicPort() (bool, error)
	AllowAccessToNetworkStorage(networkStorageTemplateObject datatypes.Network_Storage) (bool, error)
	AllowAccessToNetworkStorageList(networkStorageTemplateObjects datatypes.Network_Storage) (bool, error)
	AttachDiskImage(imageId int) (datatypes.Provisioning_Version1_Transaction, error)
	CancelIsolationForDestructiveAction() error
	CaptureImage(captureTemplate datatypes.Container_Disk_Image_Capture_Template) (datatypes.Virtual_Guest_Block_Device_Template_Group, error)
	CheckHostDiskAvailability(diskCapacity int) (bool, error)
	CloseAlarm(alarmId string) (bool, error)
	ConfigureMetadataDisk() (datatypes.Provisioning_Version1_Transaction, error)
	CreateArchiveTransaction(groupName string, blockDevices datatypes.Virtual_Guest_Block_Device, note string) (datatypes.Provisioning_Version1_Transaction, error)
	CreateObject(templateObject datatypes.Virtual_Guest) (datatypes.Virtual_Guest, error)
	CreateObjects(templateObjects datatypes.Virtual_Guest) (datatypes.Virtual_Guest, error)
	CreatePostSoftwareInstallTransaction(data string, returnBoolean bool) (bool, error)
	DeleteObject() (bool, error)
	DetachDiskImage(imageId int) (datatypes.Provisioning_Version1_Transaction, error)
	EditObject(templateObject datatypes.Virtual_Guest) (bool, error)
	ExecuteIderaBareMetalRestore() (bool, error)
	ExecuteR1SoftBareMetalRestore() (bool, error)
	ExecuteRemoteScript(uri string) error
	ExecuteRescueLayer() (bool, error)
	FindByIpAddress(ipAddress string) (datatypes.Virtual_Guest, error)
	GenerateOrderTemplate(templateObject datatypes.Virtual_Guest) (datatypes.Container_Product_Order, error)
	GetAdditionalRequiredPricesForOsReload(config datatypes.Container_Hardware_Server_Configuration) (datatypes.Product_Item_Price, error)
	GetAlarmHistory(startDate time.Time, endDate time.Time, alarmId string) (datatypes.Container_Monitoring_Alarm_History, error)
	GetAttachedNetworkStorages(nasType string) (datatypes.Network_Storage, error)
	GetAvailableBlockDevicePositions() (string, error)
	GetAvailableNetworkStorages(nasType string) (datatypes.Network_Storage, error)
	GetBandwidthDataByDate(startDateTime time.Time, endDateTime time.Time, networkType string) (datatypes.Metric_Tracking_Object_Data, error)
	GetBandwidthForDateRange(startDate time.Time, endDate time.Time) (datatypes.Metric_Tracking_Object_Data, error)
	GetBandwidthImage(networkType string, snapshotRange string, dateSpecified time.Time, dateSpecifiedEnd time.Time) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetBandwidthImageByDate(startDateTime time.Time, endDateTime time.Time, networkType string) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetBandwidthTotal(startDateTime time.Time, endDateTime time.Time, direction string, side string) (uint, error)
	GetBootOrder() (string, error)
	GetConsoleAccessLog() (datatypes.Network_Logging_Syslog, error)
	GetCoreRestrictedOperatingSystemPrice() (datatypes.Product_Item_Price, error)
	GetCpuMetricDataByDate(startDateTime time.Time, endDateTime time.Time, cpuIndexes int) (datatypes.Metric_Tracking_Object_Data, error)
	GetCpuMetricImage(snapshotRange string, dateSpecified time.Time) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetCpuMetricImageByDate(startDateTime time.Time, endDateTime time.Time, cpuIndexes int) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetCreateObjectOptions() (datatypes.Container_Virtual_Guest_Configuration, error)
	GetCurrentBillingDetail() (datatypes.Billing_Item, error)
	GetCurrentBillingTotal() (float64, error)
	GetCustomBandwidthDataByDate(graphData datatypes.Container_Graph) (datatypes.Container_Graph, error)
	GetCustomMetricDataByDate(graphData datatypes.Container_Graph) (datatypes.Container_Graph, error)
	GetDriveRetentionItemPrice() (datatypes.Product_Item_Price, error)
	GetFirewallProtectableSubnets() (datatypes.Network_Subnet, error)
	GetFirstAvailableBlockDevicePosition() (string, error)
	GetIsoBootImage() (datatypes.Virtual_Disk_Image, error)
	GetItemPricesFromSoftwareDescriptions(softwareDescriptions datatypes.Software_Description, includeTranslationsFlag bool, returnAllPricesFlag bool) (datatypes.Product_Item, error)
	GetMemoryMetricDataByDate(startDateTime time.Time, endDateTime time.Time) (datatypes.Metric_Tracking_Object_Data, error)
	GetMemoryMetricImage(snapshotRange string, dateSpecified time.Time) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetMemoryMetricImageByDate(startDateTime time.Time, endDateTime time.Time) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetMonitoringActiveAlarms(startDate time.Time, endDate time.Time) (datatypes.Container_Monitoring_Alarm_History, error)
	GetMonitoringClosedAlarms(startDate time.Time, endDate time.Time) (datatypes.Container_Monitoring_Alarm_History, error)
	GetNetworkComponentFirewallProtectableIpAddresses() (datatypes.Network_Subnet_IpAddress, error)
	GetObject() (datatypes.Virtual_Guest, error)
	GetOrderTemplate(billingType string, orderPrices datatypes.Product_Item_Price) (datatypes.Container_Product_Order, error)
	GetProvisionDate() (time.Time, error)
	GetRecentMetricData(time uint) (datatypes.Metric_Tracking_Object, error)
	GetRemoteMonitoringActiveAlarms(startDate time.Time, endDate time.Time) (datatypes.Container_Monitoring_Alarm_History, error)
	GetRemoteMonitoringClosedAlarms(startDate time.Time, endDate time.Time) (datatypes.Container_Monitoring_Alarm_History, error)
	GetReverseDomainRecords() (datatypes.Dns_Domain, error)
	GetUpgradeItemPrices(includeDowngradeItemPrices bool) (datatypes.Product_Item_Price, error)
	GetValidBlockDeviceTemplateGroups(visibility string) (datatypes.Virtual_Guest_Block_Device_Template_Group, error)
	IsBackendPingable() (bool, error)
	IsPingable() (bool, error)
	IsolateInstanceForDestructiveAction() error
	MountIsoImage(diskImageId int) (datatypes.Provisioning_Version1_Transaction, error)
	Pause() (bool, error)
	PowerCycle() (bool, error)
	PowerOff() (bool, error)
	PowerOffSoft() (bool, error)
	PowerOn() (bool, error)
	RebootDefault() (bool, error)
	RebootHard() (bool, error)
	RebootSoft() (bool, error)
	ReloadCurrentOperatingSystemConfiguration() (datatypes.Provisioning_Version1_Transaction, error)
	ReloadOperatingSystem(token string, config datatypes.Container_Hardware_Server_Configuration) (string, error)
	RemoveAccessToNetworkStorage(networkStorageTemplateObject datatypes.Network_Storage) (bool, error)
	RemoveAccessToNetworkStorageList(networkStorageTemplateObjects datatypes.Network_Storage) (bool, error)
	Resume() (bool, error)
	SetPrivateNetworkInterfaceSpeed(newSpeed int) (bool, error)
	SetPublicNetworkInterfaceSpeed(newSpeed int) (bool, error)
	SetTags(tags string) (bool, error)
	SetUserMetadata(metadata string) (bool, error)
	ShutdownPrivatePort() (bool, error)
	ShutdownPublicPort() (bool, error)
	UnmountIsoImage() (datatypes.Provisioning_Version1_Transaction, error)
	ValidateImageTemplate(imageTemplateId int) (bool, error)
	VerifyReloadOperatingSystem(config datatypes.Container_Hardware_Server_Configuration) (bool, error)

	GetAccount() (datatypes.Account, error)
	GetAccountOwnedPoolFlag() (bool, error)
	GetActiveNetworkMonitorIncident() (datatypes.Network_Monitor_Version1_Incident, error)
	GetActiveTickets() (datatypes.Ticket, error)
	GetActiveTransaction() (datatypes.Provisioning_Version1_Transaction, error)
	GetActiveTransactions() (datatypes.Provisioning_Version1_Transaction, error)
	GetAllowedHost() (datatypes.Network_Storage_Allowed_Host, error)
	GetAllowedNetworkStorage() (datatypes.Network_Storage, error)
	GetAllowedNetworkStorageReplicas() (datatypes.Network_Storage, error)
	GetAntivirusSpywareSoftwareComponent() (datatypes.Software_Component, error)
	GetApplicationDeliveryController() (datatypes.Network_Application_Delivery_Controller, error)
	GetAttributes() (datatypes.Virtual_Guest_Attribute, error)
	GetAvailableMonitoring() (datatypes.Network_Monitor_Version1_Query_Host_Stratum, error)
	GetAverageDailyPrivateBandwidthUsage() (float64, error)
	GetAverageDailyPublicBandwidthUsage() (float64, error)
	GetBackendNetworkComponents() (datatypes.Virtual_Guest_Network_Component, error)
	GetBackendRouters() (datatypes.Hardware, error)
	GetBandwidthAllocation() (float64, error)
	GetBandwidthAllotmentDetail() (datatypes.Network_Bandwidth_Version1_Allotment_Detail, error)
	GetBillingCycleBandwidthUsage() (datatypes.Network_Bandwidth_Usage, error)
	GetBillingCyclePrivateBandwidthUsage() (datatypes.Network_Bandwidth_Usage, error)
	GetBillingCyclePublicBandwidthUsage() (datatypes.Network_Bandwidth_Usage, error)
	GetBillingItem() (datatypes.Billing_Item_Virtual_Guest, error)
	GetBlockCancelBecauseDisconnectedFlag() (bool, error)
	GetBlockDeviceTemplateGroup() (datatypes.Virtual_Guest_Block_Device_Template_Group, error)
	GetBlockDevices() (datatypes.Virtual_Guest_Block_Device, error)
	GetConsoleIpAddressFlag() (bool, error)
	GetConsoleIpAddressRecord() (datatypes.Virtual_Guest_Network_Component_IpAddress, error)
	GetContinuousDataProtectionSoftwareComponent() (datatypes.Software_Component, error)
	GetControlPanel() (datatypes.Software_Component, error)
	GetCurrentBandwidthSummary() (datatypes.Metric_Tracking_Object_Bandwidth_Summary, error)
	GetDatacenter() (datatypes.Location, error)
	GetEvaultNetworkStorage() (datatypes.Network_Storage, error)
	GetFirewallServiceComponent() (datatypes.Network_Component_Firewall, error)
	GetFrontendNetworkComponents() (datatypes.Virtual_Guest_Network_Component, error)
	GetFrontendRouters() (datatypes.Hardware, error)
	GetGlobalIdentifier() (string, error)
	GetGuestBootParameter() (datatypes.Virtual_Guest_Boot_Parameter, error)
	GetHost() (datatypes.Virtual_Host, error)
	GetHostIpsSoftwareComponent() (datatypes.Software_Component, error)
	GetHourlyBillingFlag() (bool, error)
	GetInboundPrivateBandwidthUsage() (float64, error)
	GetInboundPublicBandwidthUsage() (float64, error)
	GetInternalTagReferences() (datatypes.Tag_Reference, error)
	GetLastKnownPowerState() (datatypes.Virtual_Guest_Power_State, error)
	GetLastOperatingSystemReload() (datatypes.Provisioning_Version1_Transaction, error)
	GetLastTransaction() (datatypes.Provisioning_Version1_Transaction, error)
	GetLatestNetworkMonitorIncident() (datatypes.Network_Monitor_Version1_Incident, error)
	GetLocalDiskFlag() (bool, error)
	GetLocation() (datatypes.Location, error)
	GetManagedResourceFlag() (bool, error)
	GetMetricTrackingObject() (datatypes.Metric_Tracking_Object, error)
	GetMetricTrackingObjectId() (int, error)
	GetMonitoringAgents() (datatypes.Monitoring_Agent, error)
	GetMonitoringRobot() (datatypes.Monitoring_Robot, error)
	GetMonitoringServiceComponent() (datatypes.Network_Monitor_Version1_Query_Host_Stratum, error)
	GetMonitoringServiceEligibilityFlag() (bool, error)
	GetMonitoringServiceFlag() (bool, error)
	GetMonitoringUserNotification() (datatypes.User_Customer_Notification_Virtual_Guest, error)
	GetNetworkComponents() (datatypes.Virtual_Guest_Network_Component, error)
	GetNetworkMonitorIncidents() (datatypes.Network_Monitor_Version1_Incident, error)
	GetNetworkMonitors() (datatypes.Network_Monitor_Version1_Query_Host, error)
	GetNetworkStorage() (datatypes.Network_Storage, error)
	GetNetworkVlans() (datatypes.Network_Vlan, error)
	GetOpenCancellationTicket() (datatypes.Ticket, error)
	GetOperatingSystem() (datatypes.Software_Component_OperatingSystem, error)
	GetOperatingSystemReferenceCode() (string, error)
	GetOrderedPackageId() (string, error)
	GetOutboundPrivateBandwidthUsage() (float64, error)
	GetOutboundPublicBandwidthUsage() (float64, error)
	GetOverBandwidthAllocationFlag() (int, error)
	GetPowerState() (datatypes.Virtual_Guest_Power_State, error)
	GetPrimaryBackendIpAddress() (string, error)
	GetPrimaryBackendNetworkComponent() (datatypes.Virtual_Guest_Network_Component, error)
	GetPrimaryIpAddress() (string, error)
	GetPrimaryNetworkComponent() (datatypes.Virtual_Guest_Network_Component, error)
	GetPrivateNetworkOnlyFlag() (bool, error)
	GetProjectedOverBandwidthAllocationFlag() (int, error)
	GetProjectedPublicBandwidthUsage() (float64, error)
	GetRecentEvents() (datatypes.Notification_Occurrence_Event, error)
	GetRegionalGroup() (datatypes.Location_Group_Regional, error)
	GetRegionalInternetRegistry() (datatypes.Network_Regional_Internet_Registry, error)
	GetScaleAssets() (datatypes.Scale_Asset, error)
	GetScaleMember() (datatypes.Scale_Member_Virtual_Guest, error)
	GetScaledFlag() (bool, error)
	GetSecurityScanRequests() (datatypes.Network_Security_Scanner_Request, error)
	GetServerRoom() (datatypes.Location, error)
	GetSoftwareComponents() (datatypes.Software_Component, error)
	GetSshKeys() (datatypes.Security_Ssh_Key, error)
	GetStatus() (datatypes.Virtual_Guest_Status, error)
	GetTagReferences() (datatypes.Tag_Reference, error)
	GetUpgradeRequest() (datatypes.Product_Upgrade_Request, error)
	GetUserData() (datatypes.Virtual_Guest_Attribute, error)
	GetUsers() (datatypes.User_Customer, error)
	GetVirtualRack() (datatypes.Network_Bandwidth_Version1_Allotment, error)
	GetVirtualRackId() (int, error)
	GetVirtualRackName() (string, error)
}

type Virtual_Guest_Attribute interface {
	Entity

	GetGuest() (datatypes.Virtual_Guest, error)
	GetType() (datatypes.Virtual_Guest_Attribute_Type, error)
}

type Virtual_Guest_Attribute_Type interface {
	Entity
}

type Virtual_Guest_Attribute_UserData interface {
	Virtual_Guest_Attribute
}

type Virtual_Guest_Block_Device interface {
	Entity

	GetDiskImage() (datatypes.Virtual_Disk_Image, error)
	GetGuest() (datatypes.Virtual_Guest, error)
	GetStatus() (datatypes.Virtual_Guest_Block_Device_Status, error)
}

type Virtual_Guest_Block_Device_Status interface {
	Entity
}

type Virtual_Guest_Block_Device_Template interface {
	Entity

	GetDiskImage() (datatypes.Virtual_Disk_Image, error)
	GetGroup() (datatypes.Virtual_Guest_Block_Device_Template_Group, error)
}

type Virtual_Guest_Block_Device_Template_Group interface {
	Entity

	AddLocations(locations datatypes.Location) (bool, error)
	CopyToExternalSource(configuration datatypes.Container_Virtual_Guest_Block_Device_Template_Configuration) (bool, error)
	CreateFromExternalSource(configuration datatypes.Container_Virtual_Guest_Block_Device_Template_Configuration) (datatypes.Virtual_Guest_Block_Device_Template_Group, error)
	CreatePublicArchiveTransaction(groupName string, summary string, note string, locations datatypes.Location) (int, error)
	DeleteObject() (datatypes.Provisioning_Version1_Transaction, error)
	DenySharingAccess(accountId int) (bool, error)
	EditObject(templateObject datatypes.Virtual_Guest_Block_Device_Template_Group) (bool, error)
	GetObject() (datatypes.Virtual_Guest_Block_Device_Template_Group, error)
	GetPublicCustomerOwnedImages() (datatypes.Virtual_Guest_Block_Device_Template_Group, error)
	GetPublicImages() (datatypes.Virtual_Guest_Block_Device_Template_Group, error)
	GetStorageLocations() (datatypes.Location, error)
	GetVhdImportSoftwareDescriptions() (datatypes.Software_Description, error)
	PermitSharingAccess(accountId int) (bool, error)
	RemoveLocations(locations datatypes.Location) (bool, error)
	SetAvailableLocations(locations datatypes.Location) (bool, error)
	SetTags(tags string) (bool, error)

	GetAccount() (datatypes.Account, error)
	GetAccountContacts() (datatypes.Account_Contact, error)
	GetAccountReferences() (datatypes.Virtual_Guest_Block_Device_Template_Group_Accounts, error)
	GetBlockDevices() (datatypes.Virtual_Guest_Block_Device_Template, error)
	GetBlockDevicesDiskSpaceTotal() (float64, error)
	GetChildren() (datatypes.Virtual_Guest_Block_Device_Template_Group, error)
	GetDatacenter() (datatypes.Location, error)
	GetDatacenters() (datatypes.Location, error)
	GetFlexImageFlag() (bool, error)
	GetGlobalIdentifier() (string, error)
	GetImageType() (string, error)
	GetImageTypeKeyName() (string, error)
	GetParent() (datatypes.Virtual_Guest_Block_Device_Template_Group, error)
	GetSshKeys() (datatypes.Security_Ssh_Key, error)
	GetStatus() (datatypes.Virtual_Guest_Block_Device_Template_Group_Status, error)
	GetStorageRepository() (datatypes.Virtual_Storage_Repository, error)
	GetTagReferences() (datatypes.Tag_Reference, error)
	GetTransaction() (datatypes.Provisioning_Version1_Transaction, error)
}

type Virtual_Guest_Block_Device_Template_Group_Accounts interface {
	Entity

	GetAccount() (datatypes.Account, error)
	GetGroup() (datatypes.Virtual_Guest_Block_Device_Template_Group, error)
}

type Virtual_Guest_Block_Device_Template_Group_Status interface {
	Entity
}

type Virtual_Guest_Boot_Parameter interface {
	Entity

	CreateObject(templateObject datatypes.Virtual_Guest_Boot_Parameter) (bool, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Virtual_Guest_Boot_Parameter) (bool, error)
	GetObject() (datatypes.Virtual_Guest_Boot_Parameter, error)

	GetGuest() (datatypes.Virtual_Guest, error)
	GetGuestBootParameterType() (datatypes.Virtual_Guest_Boot_Parameter_Type, error)
}

type Virtual_Guest_Boot_Parameter_Type interface {
	Entity

	GetAllObjects() (datatypes.Virtual_Guest_Boot_Parameter_Type, error)
	GetObject() (datatypes.Virtual_Guest_Boot_Parameter_Type, error)
}

type Virtual_Guest_Network_Component interface {
	Entity

	Disable() (bool, error)
	Enable() (bool, error)
	GetObject() (datatypes.Virtual_Guest_Network_Component, error)
	IsPingable() (bool, error)

	GetGuest() (datatypes.Virtual_Guest, error)
	GetHighAvailabilityFirewallFlag() (bool, error)
	GetIpAddressBindings() (datatypes.Virtual_Guest_Network_Component_IpAddress, error)
	GetNetworkComponentFirewall() (datatypes.Network_Component_Firewall, error)
	GetNetworkVlan() (datatypes.Network_Vlan, error)
	GetPrimaryIpAddress() (string, error)
	GetPrimaryIpAddressRecord() (datatypes.Network_Subnet_IpAddress, error)
	GetPrimarySubnet() (datatypes.Network_Subnet, error)
	GetPrimaryVersion6IpAddressRecord() (datatypes.Network_Subnet_IpAddress, error)
	GetRouter() (datatypes.Hardware_Router, error)
	GetSubnets() (datatypes.Network_Subnet, error)
}

type Virtual_Guest_Network_Component_IpAddress interface {
	Entity

	GetIpAddress() (datatypes.Network_Subnet_IpAddress, error)
	GetNetworkComponent() (datatypes.Virtual_Guest_Network_Component, error)
}

type Virtual_Guest_Power_State interface {
	Entity
}

type Virtual_Guest_Status interface {
	Entity
}

type Virtual_Guest_SupplementalCreateObjectOptions interface {
	Entity
}

type Virtual_Host interface {
	Entity

	GetLiveGuestByUuid(uuid string) (datatypes.Virtual_Guest, error)
	GetLiveGuestList() (datatypes.Virtual_Guest, error)
	GetLiveGuestRecentMetricData(uuid string, time int, limit int, interval int) (datatypes.Metric_Tracking_Object, error)
	GetObject() (datatypes.Virtual_Host, error)
	PauseLiveGuest(uuid string) (bool, error)
	PowerCycleLiveGuest(uuid string) (bool, error)
	PowerOffLiveGuest(uuid string) (bool, error)
	PowerOnLiveGuest(uuid string) (bool, error)
	RebootSoftLiveGuest(uuid string) (bool, error)
	ResumeLiveGuest(uuid string) (bool, error)

	GetAccount() (datatypes.Account, error)
	GetBilledPerGuestFlag() (bool, error)
	GetBilledPerMemoryUsageFlag() (bool, error)
	GetGuests() (datatypes.Virtual_Guest, error)
	GetHardware() (datatypes.Hardware_Server, error)
	GetMetricTrackingObject() (datatypes.Metric_Tracking_Object, error)
}

type Virtual_Storage_Repository interface {
	Entity

	GetArchiveDiskUsageRatePerGb() (float64, error)
	GetAverageUsageMetricDataByDate(startDateTime time.Time, endDateTime time.Time) (float64, error)
	GetObject() (datatypes.Virtual_Storage_Repository, error)
	GetPublicImageDiskUsageRatePerGb() (float64, error)
	GetStorageLocations() (datatypes.Location, error)
	GetUsageMetricDataByDate(startDateTime time.Time, endDateTime time.Time) (datatypes.Metric_Tracking_Object_Data, error)
	GetUsageMetricImageByDate(startDateTime time.Time, endDateTime time.Time) (datatypes.Container_Bandwidth_GraphOutputs, error)

	GetAccount() (datatypes.Account, error)
	GetBillingItem() (datatypes.Billing_Item, error)
	GetDatacenter() (datatypes.Location, error)
	GetDiskImages() (datatypes.Virtual_Disk_Image, error)
	GetGuests() (datatypes.Virtual_Guest, error)
	GetMetricTrackingObject() (datatypes.Metric_Tracking_Object_Virtual_Storage_Repository, error)
	GetPublicImageBillingItem() (datatypes.Billing_Item, error)
	GetType() (datatypes.Virtual_Storage_Repository_Type, error)
}

type Virtual_Storage_Repository_Type interface {
	Entity

	GetStorageRepositories() (datatypes.Virtual_Storage_Repository, error)
}
