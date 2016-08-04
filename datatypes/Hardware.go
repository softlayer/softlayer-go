package datatypes

import "time"

type Hardware struct {
	Entity

	Account                                     *Account                                     `json:"account,omitempty"`
	AccountId                                   *int                                         `json:"accountId,omitempty"`
	ActiveComponentCount                        *uint                                        `json:"activeComponentCount,omitempty"`
	ActiveComponents                            []Hardware_Component                         `json:"activeComponents,omitempty"`
	ActiveNetworkMonitorIncident                []Network_Monitor_Version1_Incident          `json:"activeNetworkMonitorIncident,omitempty"`
	ActiveNetworkMonitorIncidentCount           *uint                                        `json:"activeNetworkMonitorIncidentCount,omitempty"`
	AllPowerComponentCount                      *uint                                        `json:"allPowerComponentCount,omitempty"`
	AllPowerComponents                          []Hardware_Power_Component                   `json:"allPowerComponents,omitempty"`
	AllowedHost                                 *Network_Storage_Allowed_Host                `json:"allowedHost,omitempty"`
	AllowedNetworkStorage                       []Network_Storage                            `json:"allowedNetworkStorage,omitempty"`
	AllowedNetworkStorageCount                  *uint                                        `json:"allowedNetworkStorageCount,omitempty"`
	AllowedNetworkStorageReplicaCount           *uint                                        `json:"allowedNetworkStorageReplicaCount,omitempty"`
	AllowedNetworkStorageReplicas               []Network_Storage                            `json:"allowedNetworkStorageReplicas,omitempty"`
	AntivirusSpywareSoftwareComponent           *Software_Component                          `json:"antivirusSpywareSoftwareComponent,omitempty"`
	AttributeCount                              *uint                                        `json:"attributeCount,omitempty"`
	Attributes                                  []Hardware_Attribute                         `json:"attributes,omitempty"`
	AverageDailyPublicBandwidthUsage            *float64                                     `json:"averageDailyPublicBandwidthUsage,omitempty"`
	BackendNetworkComponentCount                *uint                                        `json:"backendNetworkComponentCount,omitempty"`
	BackendNetworkComponents                    []Network_Component                          `json:"backendNetworkComponents,omitempty"`
	BackendRouterCount                          *uint                                        `json:"backendRouterCount,omitempty"`
	BackendRouters                              []Hardware                                   `json:"backendRouters,omitempty"`
	BandwidthAllocation                         *float64                                     `json:"bandwidthAllocation,omitempty"`
	BandwidthAllotmentDetail                    *Network_Bandwidth_Version1_Allotment_Detail `json:"bandwidthAllotmentDetail,omitempty"`
	BareMetalInstanceFlag                       *int                                         `json:"bareMetalInstanceFlag,omitempty"`
	BenchmarkCertificationCount                 *uint                                        `json:"benchmarkCertificationCount,omitempty"`
	BenchmarkCertifications                     []Hardware_Benchmark_Certification           `json:"benchmarkCertifications,omitempty"`
	BillingItem                                 *Billing_Item_Hardware                       `json:"billingItem,omitempty"`
	BillingItemFlag                             *bool                                        `json:"billingItemFlag,omitempty"`
	BlockCancelBecauseDisconnectedFlag          *bool                                        `json:"blockCancelBecauseDisconnectedFlag,omitempty"`
	BusinessContinuanceInsuranceFlag            *bool                                        `json:"businessContinuanceInsuranceFlag,omitempty"`
	ComponentCount                              *uint                                        `json:"componentCount,omitempty"`
	Components                                  []Hardware_Component                         `json:"components,omitempty"`
	ContinuousDataProtectionSoftwareComponent   *Software_Component                          `json:"continuousDataProtectionSoftwareComponent,omitempty"`
	CurrentBillableBandwidthUsage               *float64                                     `json:"currentBillableBandwidthUsage,omitempty"`
	Datacenter                                  *Location                                    `json:"datacenter,omitempty"`
	DatacenterName                              *string                                      `json:"datacenterName,omitempty"`
	Domain                                      *string                                      `json:"domain,omitempty"`
	DownlinkHardware                            []Hardware                                   `json:"downlinkHardware,omitempty"`
	DownlinkHardwareCount                       *uint                                        `json:"downlinkHardwareCount,omitempty"`
	DownlinkNetworkHardware                     []Hardware                                   `json:"downlinkNetworkHardware,omitempty"`
	DownlinkNetworkHardwareCount                *uint                                        `json:"downlinkNetworkHardwareCount,omitempty"`
	DownlinkServerCount                         *uint                                        `json:"downlinkServerCount,omitempty"`
	DownlinkServers                             []Hardware                                   `json:"downlinkServers,omitempty"`
	DownlinkVirtualGuestCount                   *uint                                        `json:"downlinkVirtualGuestCount,omitempty"`
	DownlinkVirtualGuests                       []Virtual_Guest                              `json:"downlinkVirtualGuests,omitempty"`
	DownstreamHardwareBindingCount              *uint                                        `json:"downstreamHardwareBindingCount,omitempty"`
	DownstreamHardwareBindings                  []Network_Component_Uplink_Hardware          `json:"downstreamHardwareBindings,omitempty"`
	DownstreamNetworkHardware                   []Hardware                                   `json:"downstreamNetworkHardware,omitempty"`
	DownstreamNetworkHardwareCount              *uint                                        `json:"downstreamNetworkHardwareCount,omitempty"`
	DownstreamNetworkHardwareWithIncidentCount  *uint                                        `json:"downstreamNetworkHardwareWithIncidentCount,omitempty"`
	DownstreamNetworkHardwareWithIncidents      []Hardware                                   `json:"downstreamNetworkHardwareWithIncidents,omitempty"`
	DownstreamServerCount                       *uint                                        `json:"downstreamServerCount,omitempty"`
	DownstreamServers                           []Hardware                                   `json:"downstreamServers,omitempty"`
	DownstreamVirtualGuestCount                 *uint                                        `json:"downstreamVirtualGuestCount,omitempty"`
	DownstreamVirtualGuests                     []Virtual_Guest                              `json:"downstreamVirtualGuests,omitempty"`
	DriveControllerCount                        *uint                                        `json:"driveControllerCount,omitempty"`
	DriveControllers                            []Hardware_Component                         `json:"driveControllers,omitempty"`
	EvaultNetworkStorage                        []Network_Storage                            `json:"evaultNetworkStorage,omitempty"`
	EvaultNetworkStorageCount                   *uint                                        `json:"evaultNetworkStorageCount,omitempty"`
	FirewallServiceComponent                    *Network_Component_Firewall                  `json:"firewallServiceComponent,omitempty"`
	FixedConfigurationPreset                    *Product_Package_Preset                      `json:"fixedConfigurationPreset,omitempty"`
	FrontendNetworkComponentCount               *uint                                        `json:"frontendNetworkComponentCount,omitempty"`
	FrontendNetworkComponents                   []Network_Component                          `json:"frontendNetworkComponents,omitempty"`
	FrontendRouterCount                         *uint                                        `json:"frontendRouterCount,omitempty"`
	FrontendRouters                             []Hardware                                   `json:"frontendRouters,omitempty"`
	FullyQualifiedDomainName                    *string                                      `json:"fullyQualifiedDomainName,omitempty"`
	GlobalIdentifier                            *string                                      `json:"globalIdentifier,omitempty"`
	HardDriveCount                              *uint                                        `json:"hardDriveCount,omitempty"`
	HardDrives                                  []Hardware_Component                         `json:"hardDrives,omitempty"`
	HardwareChassis                             *Hardware_Chassis                            `json:"hardwareChassis,omitempty"`
	HardwareFunction                            *Hardware_Function                           `json:"hardwareFunction,omitempty"`
	HardwareFunctionDescription                 *string                                      `json:"hardwareFunctionDescription,omitempty"`
	HardwareStatus                              *Hardware_Status                             `json:"hardwareStatus,omitempty"`
	HardwareStatusId                            *int                                         `json:"hardwareStatusId,omitempty"`
	HasTrustedPlatformModuleBillingItemFlag     *bool                                        `json:"hasTrustedPlatformModuleBillingItemFlag,omitempty"`
	HostIpsSoftwareComponent                    *Software_Component                          `json:"hostIpsSoftwareComponent,omitempty"`
	Hostname                                    *string                                      `json:"hostname,omitempty"`
	HourlyBillingFlag                           *bool                                        `json:"hourlyBillingFlag,omitempty"`
	Id                                          *int                                         `json:"id,omitempty"`
	InboundBandwidthUsage                       *float64                                     `json:"inboundBandwidthUsage,omitempty"`
	InboundPublicBandwidthUsage                 *float64                                     `json:"inboundPublicBandwidthUsage,omitempty"`
	LastTransaction                             *Provisioning_Version1_Transaction           `json:"lastTransaction,omitempty"`
	LatestNetworkMonitorIncident                *Network_Monitor_Version1_Incident           `json:"latestNetworkMonitorIncident,omitempty"`
	Location                                    *Location                                    `json:"location,omitempty"`
	LocationPathString                          *string                                      `json:"locationPathString,omitempty"`
	LockboxNetworkStorage                       *Network_Storage                             `json:"lockboxNetworkStorage,omitempty"`
	ManagedResourceFlag                         *bool                                        `json:"managedResourceFlag,omitempty"`
	ManufacturerSerialNumber                    *string                                      `json:"manufacturerSerialNumber,omitempty"`
	Memory                                      []Hardware_Component                         `json:"memory,omitempty"`
	MemoryCapacity                              *uint                                        `json:"memoryCapacity,omitempty"`
	MemoryCount                                 *uint                                        `json:"memoryCount,omitempty"`
	MetricTrackingObject                        *Metric_Tracking_Object_HardwareServer       `json:"metricTrackingObject,omitempty"`
	MonitoringAgentCount                        *uint                                        `json:"monitoringAgentCount,omitempty"`
	MonitoringAgents                            []Monitoring_Agent                           `json:"monitoringAgents,omitempty"`
	MonitoringRobot                             *Monitoring_Robot                            `json:"monitoringRobot,omitempty"`
	MonitoringServiceComponent                  *Network_Monitor_Version1_Query_Host_Stratum `json:"monitoringServiceComponent,omitempty"`
	MonitoringServiceEligibilityFlag            *bool                                        `json:"monitoringServiceEligibilityFlag,omitempty"`
	MonitoringServiceFlag                       *bool                                        `json:"monitoringServiceFlag,omitempty"`
	Motherboard                                 *Hardware_Component                          `json:"motherboard,omitempty"`
	NetworkCardCount                            *uint                                        `json:"networkCardCount,omitempty"`
	NetworkCards                                []Hardware_Component                         `json:"networkCards,omitempty"`
	NetworkComponentCount                       *uint                                        `json:"networkComponentCount,omitempty"`
	NetworkComponents                           []Network_Component                          `json:"networkComponents,omitempty"`
	NetworkGatewayMember                        *Network_Gateway_Member                      `json:"networkGatewayMember,omitempty"`
	NetworkGatewayMemberFlag                    *bool                                        `json:"networkGatewayMemberFlag,omitempty"`
	NetworkManagementIpAddress                  *string                                      `json:"networkManagementIpAddress,omitempty"`
	NetworkMonitorAttachedDownHardware          []Hardware                                   `json:"networkMonitorAttachedDownHardware,omitempty"`
	NetworkMonitorAttachedDownHardwareCount     *uint                                        `json:"networkMonitorAttachedDownHardwareCount,omitempty"`
	NetworkMonitorAttachedDownVirtualGuestCount *uint                                        `json:"networkMonitorAttachedDownVirtualGuestCount,omitempty"`
	NetworkMonitorAttachedDownVirtualGuests     []Virtual_Guest                              `json:"networkMonitorAttachedDownVirtualGuests,omitempty"`
	NetworkMonitorCount                         *uint                                        `json:"networkMonitorCount,omitempty"`
	NetworkMonitorIncidentCount                 *uint                                        `json:"networkMonitorIncidentCount,omitempty"`
	NetworkMonitorIncidents                     []Network_Monitor_Version1_Incident          `json:"networkMonitorIncidents,omitempty"`
	NetworkMonitors                             []Network_Monitor_Version1_Query_Host        `json:"networkMonitors,omitempty"`
	NetworkStatus                               *string                                      `json:"networkStatus,omitempty"`
	NetworkStatusAttribute                      *Hardware_Attribute                          `json:"networkStatusAttribute,omitempty"`
	NetworkStorage                              []Network_Storage                            `json:"networkStorage,omitempty"`
	NetworkStorageCount                         *uint                                        `json:"networkStorageCount,omitempty"`
	NetworkVlanCount                            *uint                                        `json:"networkVlanCount,omitempty"`
	NetworkVlans                                []Network_Vlan                               `json:"networkVlans,omitempty"`
	NextBillingCycleBandwidthAllocation         *float64                                     `json:"nextBillingCycleBandwidthAllocation,omitempty"`
	Notes                                       *string                                      `json:"notes,omitempty"`
	NotesHistory                                []Hardware_Note                              `json:"notesHistory,omitempty"`
	NotesHistoryCount                           *uint                                        `json:"notesHistoryCount,omitempty"`
	OperatingSystem                             *Software_Component_OperatingSystem          `json:"operatingSystem,omitempty"`
	OperatingSystemReferenceCode                *string                                      `json:"operatingSystemReferenceCode,omitempty"`
	OutboundBandwidthUsage                      *float64                                     `json:"outboundBandwidthUsage,omitempty"`
	OutboundPublicBandwidthUsage                *float64                                     `json:"outboundPublicBandwidthUsage,omitempty"`
	PointOfPresenceLocation                     *Location                                    `json:"pointOfPresenceLocation,omitempty"`
	PostInstallScriptUri                        *string                                      `json:"postInstallScriptUri,omitempty"`
	PowerComponentCount                         *uint                                        `json:"powerComponentCount,omitempty"`
	PowerComponents                             []Hardware_Power_Component                   `json:"powerComponents,omitempty"`
	PowerSupply                                 []Hardware_Component                         `json:"powerSupply,omitempty"`
	PowerSupplyCount                            *uint                                        `json:"powerSupplyCount,omitempty"`
	PrimaryBackendIpAddress                     *string                                      `json:"primaryBackendIpAddress,omitempty"`
	PrimaryBackendNetworkComponent              *Network_Component                           `json:"primaryBackendNetworkComponent,omitempty"`
	PrimaryIpAddress                            *string                                      `json:"primaryIpAddress,omitempty"`
	PrimaryNetworkComponent                     *Network_Component                           `json:"primaryNetworkComponent,omitempty"`
	PrivateNetworkOnlyFlag                      *bool                                        `json:"privateNetworkOnlyFlag,omitempty"`
	ProcessorCoreAmount                         *uint                                        `json:"processorCoreAmount,omitempty"`
	ProcessorCount                              *uint                                        `json:"processorCount,omitempty"`
	ProcessorPhysicalCoreAmount                 *uint                                        `json:"processorPhysicalCoreAmount,omitempty"`
	Processors                                  []Hardware_Component                         `json:"processors,omitempty"`
	ProvisionDate                               *time.Time                                   `json:"provisionDate,omitempty"`
	Rack                                        *Location                                    `json:"rack,omitempty"`
	RaidControllerCount                         *uint                                        `json:"raidControllerCount,omitempty"`
	RaidControllers                             []Hardware_Component                         `json:"raidControllers,omitempty"`
	RecentEventCount                            *uint                                        `json:"recentEventCount,omitempty"`
	RecentEvents                                []Notification_Occurrence_Event              `json:"recentEvents,omitempty"`
	RemoteManagementAccountCount                *uint                                        `json:"remoteManagementAccountCount,omitempty"`
	RemoteManagementAccounts                    []Hardware_Component_RemoteManagement_User   `json:"remoteManagementAccounts,omitempty"`
	RemoteManagementComponent                   *Network_Component                           `json:"remoteManagementComponent,omitempty"`
	ResourceGroupCount                          *uint                                        `json:"resourceGroupCount,omitempty"`
	ResourceGroupMemberReferenceCount           *uint                                        `json:"resourceGroupMemberReferenceCount,omitempty"`
	ResourceGroupMemberReferences               []Resource_Group_Member                      `json:"resourceGroupMemberReferences,omitempty"`
	ResourceGroupRoleCount                      *uint                                        `json:"resourceGroupRoleCount,omitempty"`
	ResourceGroupRoles                          []Resource_Group_Role                        `json:"resourceGroupRoles,omitempty"`
	ResourceGroups                              []Resource_Group                             `json:"resourceGroups,omitempty"`
	RouterCount                                 *uint                                        `json:"routerCount,omitempty"`
	Routers                                     []Hardware                                   `json:"routers,omitempty"`
	ScaleAssetCount                             *uint                                        `json:"scaleAssetCount,omitempty"`
	ScaleAssets                                 []Scale_Asset                                `json:"scaleAssets,omitempty"`
	SecurityScanRequestCount                    *uint                                        `json:"securityScanRequestCount,omitempty"`
	SecurityScanRequests                        []Network_Security_Scanner_Request           `json:"securityScanRequests,omitempty"`
	SerialNumber                                *string                                      `json:"serialNumber,omitempty"`
	ServerRoom                                  *Location                                    `json:"serverRoom,omitempty"`
	ServiceProvider                             *Service_Provider                            `json:"serviceProvider,omitempty"`
	ServiceProviderId                           *int                                         `json:"serviceProviderId,omitempty"`
	ServiceProviderResourceId                   *int                                         `json:"serviceProviderResourceId,omitempty"`
	SoftwareComponentCount                      *uint                                        `json:"softwareComponentCount,omitempty"`
	SoftwareComponents                          []Software_Component                         `json:"softwareComponents,omitempty"`
	SparePoolBillingItem                        *Billing_Item_Hardware                       `json:"sparePoolBillingItem,omitempty"`
	SshKeyCount                                 *uint                                        `json:"sshKeyCount,omitempty"`
	SshKeys                                     []Security_Ssh_Key                           `json:"sshKeys,omitempty"`
	StorageNetworkComponentCount                *uint                                        `json:"storageNetworkComponentCount,omitempty"`
	StorageNetworkComponents                    []Network_Component                          `json:"storageNetworkComponents,omitempty"`
	TagReferenceCount                           *uint                                        `json:"tagReferenceCount,omitempty"`
	TagReferences                               []Tag_Reference                              `json:"tagReferences,omitempty"`
	TopLevelLocation                            *Location                                    `json:"topLevelLocation,omitempty"`
	UpgradeRequest                              *Product_Upgrade_Request                     `json:"upgradeRequest,omitempty"`
	UplinkHardware                              *Hardware                                    `json:"uplinkHardware,omitempty"`
	UplinkNetworkComponentCount                 *uint                                        `json:"uplinkNetworkComponentCount,omitempty"`
	UplinkNetworkComponents                     []Network_Component                          `json:"uplinkNetworkComponents,omitempty"`
	UserData                                    []Hardware_Attribute                         `json:"userData,omitempty"`
	UserDataCount                               *uint                                        `json:"userDataCount,omitempty"`
	VirtualChassis                              *Hardware_Group                              `json:"virtualChassis,omitempty"`
	VirtualChassisSiblingCount                  *uint                                        `json:"virtualChassisSiblingCount,omitempty"`
	VirtualChassisSiblings                      []Hardware                                   `json:"virtualChassisSiblings,omitempty"`
	VirtualHost                                 *Virtual_Host                                `json:"virtualHost,omitempty"`
	VirtualLicenseCount                         *uint                                        `json:"virtualLicenseCount,omitempty"`
	VirtualLicenses                             []Software_VirtualLicense                    `json:"virtualLicenses,omitempty"`
	VirtualRack                                 *Network_Bandwidth_Version1_Allotment        `json:"virtualRack,omitempty"`
	VirtualRackId                               *int                                         `json:"virtualRackId,omitempty"`
	VirtualRackName                             *string                                      `json:"virtualRackName,omitempty"`
	VirtualizationPlatform                      *Software_Component                          `json:"virtualizationPlatform,omitempty"`
}

type Hardware_Attribute struct {
	Entity

	HardwareAttributeType   *Hardware_Attribute_Type `json:"hardwareAttributeType,omitempty"`
	HardwareAttributeTypeId *int                     `json:"hardwareAttributeTypeId,omitempty"`
	Id                      *int                     `json:"id,omitempty"`
	Value                   *string                  `json:"value,omitempty"`
}

type Hardware_Attribute_Type struct {
	Entity

	Keyname *string `json:"keyname,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type Hardware_Benchmark_Certification struct {
	Entity

	Account    *Account   `json:"account,omitempty"`
	AccountId  *int       `json:"accountId,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	Hardware   *Hardware  `json:"hardware,omitempty"`
	HardwareId *int       `json:"hardwareId,omitempty"`
}

type Hardware_Chassis struct {
	Entity

	BackplaneCapacity       *string            `json:"backplaneCapacity,omitempty"`
	BayCapacity             *string            `json:"bayCapacity,omitempty"`
	DriveCapacity           *string            `json:"driveCapacity,omitempty"`
	DriveControllerCapacity *string            `json:"driveControllerCapacity,omitempty"`
	FormFactorId            *int               `json:"formFactorId,omitempty"`
	GpuCapacity             *string            `json:"gpuCapacity,omitempty"`
	HardwareFunction        *Hardware_Function `json:"hardwareFunction,omitempty"`
	Id                      *int               `json:"id,omitempty"`
	Manufacturer            *string            `json:"manufacturer,omitempty"`
	Name                    *string            `json:"name,omitempty"`
	PowerCapacity           *string            `json:"powerCapacity,omitempty"`
	UnitSize                *int               `json:"unitSize,omitempty"`
	Version                 *string            `json:"version,omitempty"`
}

type Hardware_Component struct {
	Entity

	Capacity                       *float64                  `json:"capacity,omitempty"`
	Children                       []Hardware_Component      `json:"children,omitempty"`
	ChildrenCount                  *uint                     `json:"childrenCount,omitempty"`
	DownlinkHardwareComponentCount *uint                     `json:"downlinkHardwareComponentCount,omitempty"`
	DownlinkHardwareComponents     []Hardware_Component      `json:"downlinkHardwareComponents,omitempty"`
	Hardware                       *Hardware                 `json:"hardware,omitempty"`
	HardwareComponentModel         *Hardware_Component_Model `json:"hardwareComponentModel,omitempty"`
	HardwareComponentModelId       *int                      `json:"hardwareComponentModelId,omitempty"`
	HardwareComponentType          *Hardware_Component_Type  `json:"hardwareComponentType,omitempty"`
	HardwareId                     *int                      `json:"hardwareId,omitempty"`
	Id                             *int                      `json:"id,omitempty"`
	ModifyDate                     *time.Time                `json:"modifyDate,omitempty"`
	Name                           *string                   `json:"name,omitempty"`
	NetworkComponentCount          *uint                     `json:"networkComponentCount,omitempty"`
	NetworkComponents              []Network_Component       `json:"networkComponents,omitempty"`
	Owner                          *Account                  `json:"owner,omitempty"`
	Parent                         *Hardware_Component       `json:"parent,omitempty"`
	RaidMode                       *string                   `json:"raidMode,omitempty"`
	SerialNumber                   *string                   `json:"serialNumber,omitempty"`
	ServiceProvider                *Service_Provider         `json:"serviceProvider,omitempty"`
	ServiceProviderId              *int                      `json:"serviceProviderId,omitempty"`
	UplinkHardwareComponentCount   *uint                     `json:"uplinkHardwareComponentCount,omitempty"`
	UplinkHardwareComponents       []Hardware_Component      `json:"uplinkHardwareComponents,omitempty"`
}

type Hardware_Component_Attribute struct {
	Entity

	HardwareComponent                *Hardware_Component                `json:"hardwareComponent,omitempty"`
	HardwareComponentAttributeType   *Hardware_Component_Attribute_Type `json:"hardwareComponentAttributeType,omitempty"`
	HardwareComponentAttributeTypeId *int                               `json:"hardwareComponentAttributeTypeId,omitempty"`
	HardwareComponentId              *int                               `json:"hardwareComponentId,omitempty"`
	Value                            *string                            `json:"value,omitempty"`
}

type Hardware_Component_Attribute_Type struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	KeyName     *string `json:"keyName,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Hardware_Component_DriveController struct {
	Hardware_Component
}

type Hardware_Component_HardDrive struct {
	Hardware_Component

	PartitionCount *uint                          `json:"partitionCount,omitempty"`
	Partitions     []Hardware_Component_Partition `json:"partitions,omitempty"`
}

type Hardware_Component_Model struct {
	Entity

	ArchitectureType                    *Hardware_Component_Model_Architecture_Type `json:"architectureType,omitempty"`
	ArchitectureTypeId                  *string                                     `json:"architectureTypeId,omitempty"`
	AttributeCount                      *uint                                       `json:"attributeCount,omitempty"`
	Attributes                          []Hardware_Component_Model_Attribute        `json:"attributes,omitempty"`
	Capacity                            *float64                                    `json:"capacity,omitempty"`
	CompatibleArrayTypeCount            *uint                                       `json:"compatibleArrayTypeCount,omitempty"`
	CompatibleArrayTypes                []Configuration_Storage_Group_Array_Type    `json:"compatibleArrayTypes,omitempty"`
	CompatibleChildComponentModelCount  *uint                                       `json:"compatibleChildComponentModelCount,omitempty"`
	CompatibleChildComponentModels      []Hardware_Component_Model                  `json:"compatibleChildComponentModels,omitempty"`
	CompatibleParentComponentModelCount *uint                                       `json:"compatibleParentComponentModelCount,omitempty"`
	CompatibleParentComponentModels     []Hardware_Component_Model                  `json:"compatibleParentComponentModels,omitempty"`
	Description                         *string                                     `json:"description,omitempty"`
	HardwareComponents                  []Hardware_Component                        `json:"hardwareComponents,omitempty"`
	HardwareGenericComponentModel       *Hardware_Component_Model_Generic           `json:"hardwareGenericComponentModel,omitempty"`
	HardwareGenericComponentModelId     *int                                        `json:"hardwareGenericComponentModelId,omitempty"`
	Id                                  *int                                        `json:"id,omitempty"`
	InfinibandCompatibleAttribute       *Hardware_Component_Model_Attribute         `json:"infinibandCompatibleAttribute,omitempty"`
	IsInfinibandCompatible              *bool                                       `json:"isInfinibandCompatible,omitempty"`
	LongDescription                     *string                                     `json:"longDescription,omitempty"`
	Manufacturer                        *string                                     `json:"manufacturer,omitempty"`
	Name                                *string                                     `json:"name,omitempty"`
	RebootTime                          *Hardware_Component_Motherboard_Reboot_Time `json:"rebootTime,omitempty"`
	Type                                *string                                     `json:"type,omitempty"`
	ValidAttributeTypeCount             *uint                                       `json:"validAttributeTypeCount,omitempty"`
	ValidAttributeTypes                 []Hardware_Component_Model_Attribute_Type   `json:"validAttributeTypes,omitempty"`
	Version                             *string                                     `json:"version,omitempty"`
}

type Hardware_Component_Model_Architecture_Type struct {
	Entity

	Children      []Hardware_Component_Model_Architecture_Type `json:"children,omitempty"`
	ChildrenCount *uint                                        `json:"childrenCount,omitempty"`
	Id            *int                                         `json:"id,omitempty"`
	KeyName       *string                                      `json:"keyName,omitempty"`
	Name          *string                                      `json:"name,omitempty"`
	Parent        *Hardware_Component_Model_Architecture_Type  `json:"parent,omitempty"`
	ParentId      *string                                      `json:"parentId,omitempty"`
}

type Hardware_Component_Model_Attribute struct {
	Entity

	AttributeTypeId                *int                                     `json:"attributeTypeId,omitempty"`
	HardwareComponent              *Hardware_Component_Model                `json:"hardwareComponent,omitempty"`
	HardwareComponentAttributeType *Hardware_Component_Model_Attribute_Type `json:"hardwareComponentAttributeType,omitempty"`
	HardwareComponentModelId       *int                                     `json:"hardwareComponentModelId,omitempty"`
	Value                          *string                                  `json:"value,omitempty"`
}

type Hardware_Component_Model_Attribute_Type struct {
	Entity

	Description             *string                   `json:"description,omitempty"`
	Id                      *int                      `json:"id,omitempty"`
	KeyName                 *string                   `json:"keyName,omitempty"`
	Name                    *string                   `json:"name,omitempty"`
	ValidComponentTypeCount *uint                     `json:"validComponentTypeCount,omitempty"`
	ValidComponentTypes     []Hardware_Component_Type `json:"validComponentTypes,omitempty"`
}

type Hardware_Component_Model_Generic struct {
	Entity

	Capacity                    *float64                                           `json:"capacity,omitempty"`
	Description                 *string                                            `json:"description,omitempty"`
	HardwareComponentModelCount *uint                                              `json:"hardwareComponentModelCount,omitempty"`
	HardwareComponentModels     []Hardware_Component_Model                         `json:"hardwareComponentModels,omitempty"`
	HardwareComponentType       *Hardware_Component_Type                           `json:"hardwareComponentType,omitempty"`
	HardwareComponentTypeId     *int                                               `json:"hardwareComponentTypeId,omitempty"`
	Id                          *int                                               `json:"id,omitempty"`
	MarketingFeatures           *Hardware_Component_Model_Generic_MarketingFeature `json:"marketingFeatures,omitempty"`
	Units                       *string                                            `json:"units,omitempty"`
	UpgradePriority             *int                                               `json:"upgradePriority,omitempty"`
}

type Hardware_Component_Model_Generic_Attribute struct {
	Entity

	HardwareGenericComponentModel *Hardware_Component_Model_Generic `json:"hardwareGenericComponentModel,omitempty"`
	Value                         *string                           `json:"value,omitempty"`
}

type Hardware_Component_Model_Generic_MarketingFeature struct {
	Entity

	Features                      *string                           `json:"features,omitempty"`
	HardwareGenericComponentModel *Hardware_Component_Model_Generic `json:"hardwareGenericComponentModel,omitempty"`
	Price                         *string                           `json:"price,omitempty"`
}

type Hardware_Component_Motherboard struct {
	Hardware_Component
}

type Hardware_Component_Motherboard_Reboot_Time struct {
	Entity

	HardwareComponentModel *Hardware_Component_Model `json:"hardwareComponentModel,omitempty"`
	WithRaid               *int                      `json:"withRaid,omitempty"`
	WithoutRaid            *int                      `json:"withoutRaid,omitempty"`
}

type Hardware_Component_NetworkCard struct {
	Hardware_Component
}

type Hardware_Component_Partition struct {
	Entity

	DiskNumber          *int                `json:"diskNumber,omitempty"`
	Grow                *int                `json:"grow,omitempty"`
	HardwareComponent   *Hardware_Component `json:"hardwareComponent,omitempty"`
	HardwareComponentId *int                `json:"hardwareComponentId,omitempty"`
	MinimumSize         *float64            `json:"minimumSize,omitempty"`
	Name                *string             `json:"name,omitempty"`
}

type Hardware_Component_Partition_OperatingSystem struct {
	Entity

	Description            *string                                 `json:"description,omitempty"`
	Id                     *int                                    `json:"id,omitempty"`
	Notes                  *string                                 `json:"notes,omitempty"`
	PartitionTemplateCount *uint                                   `json:"partitionTemplateCount,omitempty"`
	PartitionTemplates     []Hardware_Component_Partition_Template `json:"partitionTemplates,omitempty"`
}

type Hardware_Component_Partition_Template struct {
	Entity

	Account                         *Account                                          `json:"account,omitempty"`
	AccountId                       *int                                              `json:"accountId,omitempty"`
	Data                            []Hardware_Component_Partition_Template_Partition `json:"data,omitempty"`
	DataCount                       *uint                                             `json:"dataCount,omitempty"`
	Description                     *string                                           `json:"description,omitempty"`
	ExpireDate                      *string                                           `json:"expireDate,omitempty"`
	Id                              *int                                              `json:"id,omitempty"`
	PartitionOperatingSystem        *Hardware_Component_Partition_OperatingSystem     `json:"partitionOperatingSystem,omitempty"`
	PartitionOperatingSystemId      *int                                              `json:"partitionOperatingSystemId,omitempty"`
	PartitionTemplatePartition      []Hardware_Component_Partition_Template_Partition `json:"partitionTemplatePartition,omitempty"`
	PartitionTemplatePartitionCount *uint                                             `json:"partitionTemplatePartitionCount,omitempty"`
	StatusCode                      *string                                           `json:"statusCode,omitempty"`
	TemplateType                    *string                                           `json:"templateType,omitempty"`
}

type Hardware_Component_Partition_Template_Partition struct {
	Entity

	FilesystemType      *Configuration_Storage_Filesystem_Type `json:"filesystemType,omitempty"`
	Id                  *int                                   `json:"id,omitempty"`
	IsGrow              *bool                                  `json:"isGrow,omitempty"`
	PartitionName       *string                                `json:"partitionName,omitempty"`
	PartitionSize       *float64                               `json:"partitionSize,omitempty"`
	PartitionTemplate   *Hardware_Component_Partition_Template `json:"partitionTemplate,omitempty"`
	PartitionTemplateId *int                                   `json:"partitionTemplateId,omitempty"`
}

type Hardware_Component_Processor struct {
	Hardware_Component
}

type Hardware_Component_Ram struct {
	Hardware_Component
}

type Hardware_Component_RemoteManagement struct {
	Hardware_Component

	NetworkComponent *Network_Component `json:"networkComponent,omitempty"`
}

type Hardware_Component_RemoteManagement_Command struct {
	Entity

	KeyName      *string                                               `json:"keyName,omitempty"`
	RequestCount *uint                                                 `json:"requestCount,omitempty"`
	Requests     []Hardware_Component_RemoteManagement_Command_Request `json:"requests,omitempty"`
}

type Hardware_Component_RemoteManagement_Command_Request struct {
	Entity

	CreateDate              *time.Time                                   `json:"createDate,omitempty"`
	Hardware                *Hardware                                    `json:"hardware,omitempty"`
	HardwareId              *int                                         `json:"hardwareId,omitempty"`
	ModifyDate              *time.Time                                   `json:"modifyDate,omitempty"`
	NetworkComponent        *Network_Component                           `json:"networkComponent,omitempty"`
	Processed               *bool                                        `json:"processed,omitempty"`
	RemoteManagementCommand *Hardware_Component_RemoteManagement_Command `json:"remoteManagementCommand,omitempty"`
	User                    *User_Customer                               `json:"user,omitempty"`
}

type Hardware_Component_RemoteManagement_User struct {
	Entity

	Hardware         *Hardware          `json:"hardware,omitempty"`
	NetworkComponent *Network_Component `json:"networkComponent,omitempty"`
	Password         *string            `json:"password,omitempty"`
	Username         *string            `json:"username,omitempty"`
}

type Hardware_Component_SecurityDevice struct {
	Hardware_Component
}

type Hardware_Component_SecurityDevice_Infineon struct {
	Hardware_Component_SecurityDevice
}

type Hardware_Component_Type struct {
	Entity

	HardwareGenericComponentModelCount *uint                              `json:"hardwareGenericComponentModelCount,omitempty"`
	HardwareGenericComponentModels     []Hardware_Component_Model_Generic `json:"hardwareGenericComponentModels,omitempty"`
	Id                                 *int                               `json:"id,omitempty"`
	KeyName                            *string                            `json:"keyName,omitempty"`
	Type                               *string                            `json:"type,omitempty"`
	TypeParent                         *Hardware_Component_Type           `json:"typeParent,omitempty"`
	TypeParentId                       *int                               `json:"typeParentId,omitempty"`
}

type Hardware_Firewall struct {
	Hardware_Switch

	UserCount *uint           `json:"userCount,omitempty"`
	Users     []User_Customer `json:"users,omitempty"`
}

type Hardware_Function struct {
	Entity

	Code        *string `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
}

type Hardware_Group struct {
	Entity

	Domain                                      *string           `json:"domain,omitempty"`
	DownlinkServerCount                         *uint             `json:"downlinkServerCount,omitempty"`
	DownlinkServers                             []Hardware        `json:"downlinkServers,omitempty"`
	DownlinkVirtualGuestCount                   *uint             `json:"downlinkVirtualGuestCount,omitempty"`
	DownlinkVirtualGuests                       []Virtual_Guest   `json:"downlinkVirtualGuests,omitempty"`
	DownstreamNetworkHardware                   []Hardware        `json:"downstreamNetworkHardware,omitempty"`
	DownstreamNetworkHardwareCount              *uint             `json:"downstreamNetworkHardwareCount,omitempty"`
	DownstreamNetworkHardwareWithIncidentCount  *uint             `json:"downstreamNetworkHardwareWithIncidentCount,omitempty"`
	DownstreamNetworkHardwareWithIncidents      []Hardware        `json:"downstreamNetworkHardwareWithIncidents,omitempty"`
	HardwareChassis                             *Hardware_Chassis `json:"hardwareChassis,omitempty"`
	Hostname                                    *string           `json:"hostname,omitempty"`
	Id                                          *int              `json:"id,omitempty"`
	NetworkMonitorAttachedDownHardware          []Hardware        `json:"networkMonitorAttachedDownHardware,omitempty"`
	NetworkMonitorAttachedDownHardwareCount     *uint             `json:"networkMonitorAttachedDownHardwareCount,omitempty"`
	NetworkMonitorAttachedDownVirtualGuestCount *uint             `json:"networkMonitorAttachedDownVirtualGuestCount,omitempty"`
	NetworkMonitorAttachedDownVirtualGuests     []Virtual_Guest   `json:"networkMonitorAttachedDownVirtualGuests,omitempty"`
	NetworkStatus                               *string           `json:"networkStatus,omitempty"`
}

type Hardware_LoadBalancer struct {
	Hardware

	ModelFamily *string         `json:"modelFamily,omitempty"`
	UserCount   *uint           `json:"userCount,omitempty"`
	Users       []User_Customer `json:"users,omitempty"`
}

type Hardware_Note struct {
	Entity

	CreateDate   *time.Time          `json:"createDate,omitempty"`
	Employee     *User_Employee      `json:"employee,omitempty"`
	Hardware     *Hardware           `json:"hardware,omitempty"`
	HardwareId   *int                `json:"hardwareId,omitempty"`
	Id           *int                `json:"id,omitempty"`
	ModifyDate   *time.Time          `json:"modifyDate,omitempty"`
	Note         *string             `json:"note,omitempty"`
	Type         *Hardware_Note_Type `json:"type,omitempty"`
	TypeId       *int                `json:"typeId,omitempty"`
	User         *User_Customer      `json:"user,omitempty"`
	UserRecordId *int                `json:"userRecordId,omitempty"`
}

type Hardware_Note_Type struct {
	Entity

	KeyName *string `json:"keyName,omitempty"`
}

type Hardware_Power_Component struct {
	Entity

	Hardware   *Hardware `json:"hardware,omitempty"`
	HardwareId *int      `json:"hardwareId,omitempty"`
	Id         *int      `json:"id,omitempty"`
}

type Hardware_Router struct {
	Hardware_Switch

	BoundSubnetCount               *uint            `json:"boundSubnetCount,omitempty"`
	BoundSubnets                   []Network_Subnet `json:"boundSubnets,omitempty"`
	LocalDiskStorageCapabilityFlag *bool            `json:"localDiskStorageCapabilityFlag,omitempty"`
	SanStorageCapabilityFlag       *bool            `json:"sanStorageCapabilityFlag,omitempty"`
}

type Hardware_Router_Backend struct {
	Hardware_Router
}

type Hardware_Router_Frontend struct {
	Hardware_Router
}

type Hardware_SecurityModule struct {
	Hardware_Server
}

type Hardware_Server struct {
	Hardware

	ActiveNetworkFirewallBillingItem     *Billing_Item                                         `json:"activeNetworkFirewallBillingItem,omitempty"`
	ActiveTicketCount                    *uint                                                 `json:"activeTicketCount,omitempty"`
	ActiveTickets                        []Ticket                                              `json:"activeTickets,omitempty"`
	ActiveTransaction                    *Provisioning_Version1_Transaction                    `json:"activeTransaction,omitempty"`
	ActiveTransactionCount               *uint                                                 `json:"activeTransactionCount,omitempty"`
	ActiveTransactions                   []Provisioning_Version1_Transaction                   `json:"activeTransactions,omitempty"`
	AvailableMonitoring                  []Network_Monitor_Version1_Query_Host_Stratum         `json:"availableMonitoring,omitempty"`
	AvailableMonitoringCount             *uint                                                 `json:"availableMonitoringCount,omitempty"`
	AverageDailyBandwidthUsage           *float64                                              `json:"averageDailyBandwidthUsage,omitempty"`
	AverageDailyPrivateBandwidthUsage    *float64                                              `json:"averageDailyPrivateBandwidthUsage,omitempty"`
	BillingCycleBandwidthUsage           []Network_Bandwidth_Usage                             `json:"billingCycleBandwidthUsage,omitempty"`
	BillingCycleBandwidthUsageCount      *uint                                                 `json:"billingCycleBandwidthUsageCount,omitempty"`
	BillingCyclePrivateBandwidthUsage    *Network_Bandwidth_Usage                              `json:"billingCyclePrivateBandwidthUsage,omitempty"`
	BillingCyclePublicBandwidthUsage     *Network_Bandwidth_Usage                              `json:"billingCyclePublicBandwidthUsage,omitempty"`
	ContainsSolidStateDrivesFlag         *bool                                                 `json:"containsSolidStateDrivesFlag,omitempty"`
	ControlPanel                         *Software_Component_ControlPanel                      `json:"controlPanel,omitempty"`
	Cost                                 *float64                                              `json:"cost,omitempty"`
	CurrentBandwidthSummary              *Metric_Tracking_Object_Bandwidth_Summary             `json:"currentBandwidthSummary,omitempty"`
	CustomerInstalledOperatingSystemFlag *bool                                                 `json:"customerInstalledOperatingSystemFlag,omitempty"`
	CustomerOwnedFlag                    *bool                                                 `json:"customerOwnedFlag,omitempty"`
	InboundPrivateBandwidthUsage         *float64                                              `json:"inboundPrivateBandwidthUsage,omitempty"`
	LastOperatingSystemReload            *Provisioning_Version1_Transaction                    `json:"lastOperatingSystemReload,omitempty"`
	MetricTrackingObjectId               *int                                                  `json:"metricTrackingObjectId,omitempty"`
	MonitoringUserNotification           []User_Customer_Notification_Hardware                 `json:"monitoringUserNotification,omitempty"`
	MonitoringUserNotificationCount      *uint                                                 `json:"monitoringUserNotificationCount,omitempty"`
	OpenCancellationTicket               *Ticket                                               `json:"openCancellationTicket,omitempty"`
	OutboundPrivateBandwidthUsage        *float64                                              `json:"outboundPrivateBandwidthUsage,omitempty"`
	OverBandwidthAllocationFlag          *int                                                  `json:"overBandwidthAllocationFlag,omitempty"`
	PrivateIpAddress                     *string                                               `json:"privateIpAddress,omitempty"`
	ProjectedOverBandwidthAllocationFlag *int                                                  `json:"projectedOverBandwidthAllocationFlag,omitempty"`
	ProjectedPublicBandwidthUsage        *float64                                              `json:"projectedPublicBandwidthUsage,omitempty"`
	RecentRemoteManagementCommandCount   *uint                                                 `json:"recentRemoteManagementCommandCount,omitempty"`
	RecentRemoteManagementCommands       []Hardware_Component_RemoteManagement_Command_Request `json:"recentRemoteManagementCommands,omitempty"`
	RegionalInternetRegistry             *Network_Regional_Internet_Registry                   `json:"regionalInternetRegistry,omitempty"`
	RemoteManagement                     *Hardware_Component_RemoteManagement                  `json:"remoteManagement,omitempty"`
	RemoteManagementUserCount            *uint                                                 `json:"remoteManagementUserCount,omitempty"`
	RemoteManagementUsers                []Hardware_Component_RemoteManagement_User            `json:"remoteManagementUsers,omitempty"`
	StatisticsRemoteManagement           *Hardware_Component_RemoteManagement                  `json:"statisticsRemoteManagement,omitempty"`
	UserCount                            *uint                                                 `json:"userCount,omitempty"`
	Users                                []User_Customer                                       `json:"users,omitempty"`
	VirtualGuestCount                    *uint                                                 `json:"virtualGuestCount,omitempty"`
	VirtualGuests                        []Virtual_Guest                                       `json:"virtualGuests,omitempty"`
}

type Hardware_Status struct {
	Entity

	Id     *int    `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
}

type Hardware_Switch struct {
	Hardware
}
