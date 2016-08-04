package datatypes

import "time"

type Virtual_Disk_Image struct {
	Entity

	BillingItem             *Billing_Item_Virtual_Disk_Image     `json:"billingItem,omitempty"`
	BlockDeviceCount        *uint                                `json:"blockDeviceCount,omitempty"`
	BlockDevices            []Virtual_Guest_Block_Device         `json:"blockDevices,omitempty"`
	BootableVolumeFlag      *bool                                `json:"bootableVolumeFlag,omitempty"`
	Capacity                *int                                 `json:"capacity,omitempty"`
	Checksum                *string                              `json:"checksum,omitempty"`
	CoalescedDiskImageCount *uint                                `json:"coalescedDiskImageCount,omitempty"`
	CoalescedDiskImages     []Virtual_Disk_Image                 `json:"coalescedDiskImages,omitempty"`
	CopyOnWriteFlag         *bool                                `json:"copyOnWriteFlag,omitempty"`
	CreateDate              *time.Time                           `json:"createDate,omitempty"`
	Description             *string                              `json:"description,omitempty"`
	Id                      *int                                 `json:"id,omitempty"`
	LocalDiskFlag           *bool                                `json:"localDiskFlag,omitempty"`
	MetadataFlag            *bool                                `json:"metadataFlag,omitempty"`
	ModifyDate              *time.Time                           `json:"modifyDate,omitempty"`
	Name                    *string                              `json:"name,omitempty"`
	ParentId                *int                                 `json:"parentId,omitempty"`
	SoftwareReferenceCount  *uint                                `json:"softwareReferenceCount,omitempty"`
	SoftwareReferences      []Virtual_Disk_Image_Software        `json:"softwareReferences,omitempty"`
	SourceDiskImage         *Virtual_Disk_Image                  `json:"sourceDiskImage,omitempty"`
	StorageRepository       *Virtual_Storage_Repository          `json:"storageRepository,omitempty"`
	StorageRepositoryId     *int                                 `json:"storageRepositoryId,omitempty"`
	StorageRepositoryType   *Virtual_Storage_Repository_Type     `json:"storageRepositoryType,omitempty"`
	TemplateBlockDevice     *Virtual_Guest_Block_Device_Template `json:"templateBlockDevice,omitempty"`
	Type                    *Virtual_Disk_Image_Type             `json:"type,omitempty"`
	TypeId                  *int                                 `json:"typeId,omitempty"`
	Units                   *string                              `json:"units,omitempty"`
	Uuid                    *string                              `json:"uuid,omitempty"`
}

type Virtual_Disk_Image_Software struct {
	Entity

	DiskImage             *Virtual_Disk_Image                    `json:"diskImage,omitempty"`
	Id                    *int                                   `json:"id,omitempty"`
	PasswordCount         *uint                                  `json:"passwordCount,omitempty"`
	Passwords             []Virtual_Disk_Image_Software_Password `json:"passwords,omitempty"`
	SoftwareDescription   *Software_Description                  `json:"softwareDescription,omitempty"`
	SoftwareDescriptionId *int                                   `json:"softwareDescriptionId,omitempty"`
}

type Virtual_Disk_Image_Software_Password struct {
	Entity

	Password *string                      `json:"password,omitempty"`
	Software *Virtual_Disk_Image_Software `json:"software,omitempty"`
	Username *string                      `json:"username,omitempty"`
}

type Virtual_Disk_Image_Type struct {
	Entity

	Description *string `json:"description,omitempty"`
	KeyName     *string `json:"keyName,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Virtual_Guest struct {
	Entity

	Account                                   *Account                                       `json:"account,omitempty"`
	AccountId                                 *int                                           `json:"accountId,omitempty"`
	AccountOwnedPoolFlag                      *bool                                          `json:"accountOwnedPoolFlag,omitempty"`
	ActiveNetworkMonitorIncident              []Network_Monitor_Version1_Incident            `json:"activeNetworkMonitorIncident,omitempty"`
	ActiveNetworkMonitorIncidentCount         *uint                                          `json:"activeNetworkMonitorIncidentCount,omitempty"`
	ActiveTicketCount                         *uint                                          `json:"activeTicketCount,omitempty"`
	ActiveTickets                             []Ticket                                       `json:"activeTickets,omitempty"`
	ActiveTransaction                         *Provisioning_Version1_Transaction             `json:"activeTransaction,omitempty"`
	ActiveTransactionCount                    *uint                                          `json:"activeTransactionCount,omitempty"`
	ActiveTransactions                        []Provisioning_Version1_Transaction            `json:"activeTransactions,omitempty"`
	AllowedHost                               *Network_Storage_Allowed_Host                  `json:"allowedHost,omitempty"`
	AllowedNetworkStorage                     []Network_Storage                              `json:"allowedNetworkStorage,omitempty"`
	AllowedNetworkStorageCount                *uint                                          `json:"allowedNetworkStorageCount,omitempty"`
	AllowedNetworkStorageReplicaCount         *uint                                          `json:"allowedNetworkStorageReplicaCount,omitempty"`
	AllowedNetworkStorageReplicas             []Network_Storage                              `json:"allowedNetworkStorageReplicas,omitempty"`
	AntivirusSpywareSoftwareComponent         *Software_Component                            `json:"antivirusSpywareSoftwareComponent,omitempty"`
	ApplicationDeliveryController             *Network_Application_Delivery_Controller       `json:"applicationDeliveryController,omitempty"`
	AttributeCount                            *uint                                          `json:"attributeCount,omitempty"`
	Attributes                                []Virtual_Guest_Attribute                      `json:"attributes,omitempty"`
	AvailableMonitoring                       []Network_Monitor_Version1_Query_Host_Stratum  `json:"availableMonitoring,omitempty"`
	AvailableMonitoringCount                  *uint                                          `json:"availableMonitoringCount,omitempty"`
	AverageDailyPrivateBandwidthUsage         *float64                                       `json:"averageDailyPrivateBandwidthUsage,omitempty"`
	AverageDailyPublicBandwidthUsage          *float64                                       `json:"averageDailyPublicBandwidthUsage,omitempty"`
	BackendNetworkComponentCount              *uint                                          `json:"backendNetworkComponentCount,omitempty"`
	BackendNetworkComponents                  []Virtual_Guest_Network_Component              `json:"backendNetworkComponents,omitempty"`
	BackendRouterCount                        *uint                                          `json:"backendRouterCount,omitempty"`
	BackendRouters                            []Hardware                                     `json:"backendRouters,omitempty"`
	BandwidthAllocation                       *float64                                       `json:"bandwidthAllocation,omitempty"`
	BandwidthAllotmentDetail                  *Network_Bandwidth_Version1_Allotment_Detail   `json:"bandwidthAllotmentDetail,omitempty"`
	BillingCycleBandwidthUsage                []Network_Bandwidth_Usage                      `json:"billingCycleBandwidthUsage,omitempty"`
	BillingCycleBandwidthUsageCount           *uint                                          `json:"billingCycleBandwidthUsageCount,omitempty"`
	BillingCyclePrivateBandwidthUsage         *Network_Bandwidth_Usage                       `json:"billingCyclePrivateBandwidthUsage,omitempty"`
	BillingCyclePublicBandwidthUsage          *Network_Bandwidth_Usage                       `json:"billingCyclePublicBandwidthUsage,omitempty"`
	BillingItem                               *Billing_Item_Virtual_Guest                    `json:"billingItem,omitempty"`
	BlockCancelBecauseDisconnectedFlag        *bool                                          `json:"blockCancelBecauseDisconnectedFlag,omitempty"`
	BlockDeviceCount                          *uint                                          `json:"blockDeviceCount,omitempty"`
	BlockDeviceTemplateGroup                  *Virtual_Guest_Block_Device_Template_Group     `json:"blockDeviceTemplateGroup,omitempty"`
	BlockDevices                              []Virtual_Guest_Block_Device                   `json:"blockDevices,omitempty"`
	ConsoleIpAddressFlag                      *bool                                          `json:"consoleIpAddressFlag,omitempty"`
	ConsoleIpAddressRecord                    *Virtual_Guest_Network_Component_IpAddress     `json:"consoleIpAddressRecord,omitempty"`
	ContinuousDataProtectionSoftwareComponent *Software_Component                            `json:"continuousDataProtectionSoftwareComponent,omitempty"`
	ControlPanel                              *Software_Component                            `json:"controlPanel,omitempty"`
	CreateDate                                *time.Time                                     `json:"createDate,omitempty"`
	CurrentBandwidthSummary                   *Metric_Tracking_Object_Bandwidth_Summary      `json:"currentBandwidthSummary,omitempty"`
	Datacenter                                *Location                                      `json:"datacenter,omitempty"`
	DedicatedAccountHostOnlyFlag              *bool                                          `json:"dedicatedAccountHostOnlyFlag,omitempty"`
	Domain                                    *string                                        `json:"domain,omitempty"`
	EvaultNetworkStorage                      []Network_Storage                              `json:"evaultNetworkStorage,omitempty"`
	EvaultNetworkStorageCount                 *uint                                          `json:"evaultNetworkStorageCount,omitempty"`
	FirewallServiceComponent                  *Network_Component_Firewall                    `json:"firewallServiceComponent,omitempty"`
	FrontendNetworkComponentCount             *uint                                          `json:"frontendNetworkComponentCount,omitempty"`
	FrontendNetworkComponents                 []Virtual_Guest_Network_Component              `json:"frontendNetworkComponents,omitempty"`
	FrontendRouters                           *Hardware                                      `json:"frontendRouters,omitempty"`
	FullyQualifiedDomainName                  *string                                        `json:"fullyQualifiedDomainName,omitempty"`
	GlobalIdentifier                          *string                                        `json:"globalIdentifier,omitempty"`
	GuestBootParameter                        *Virtual_Guest_Boot_Parameter                  `json:"guestBootParameter,omitempty"`
	Host                                      *Virtual_Host                                  `json:"host,omitempty"`
	HostIpsSoftwareComponent                  *Software_Component                            `json:"hostIpsSoftwareComponent,omitempty"`
	Hostname                                  *string                                        `json:"hostname,omitempty"`
	HourlyBillingFlag                         *bool                                          `json:"hourlyBillingFlag,omitempty"`
	Id                                        *int                                           `json:"id,omitempty"`
	InboundPrivateBandwidthUsage              *float64                                       `json:"inboundPrivateBandwidthUsage,omitempty"`
	InboundPublicBandwidthUsage               *float64                                       `json:"inboundPublicBandwidthUsage,omitempty"`
	InternalTagReferenceCount                 *uint                                          `json:"internalTagReferenceCount,omitempty"`
	InternalTagReferences                     []Tag_Reference                                `json:"internalTagReferences,omitempty"`
	LastKnownPowerState                       *Virtual_Guest_Power_State                     `json:"lastKnownPowerState,omitempty"`
	LastOperatingSystemReload                 *Provisioning_Version1_Transaction             `json:"lastOperatingSystemReload,omitempty"`
	LastPowerStateId                          *int                                           `json:"lastPowerStateId,omitempty"`
	LastTransaction                           *Provisioning_Version1_Transaction             `json:"lastTransaction,omitempty"`
	LastVerifiedDate                          *time.Time                                     `json:"lastVerifiedDate,omitempty"`
	LatestNetworkMonitorIncident              *Network_Monitor_Version1_Incident             `json:"latestNetworkMonitorIncident,omitempty"`
	LocalDiskFlag                             *bool                                          `json:"localDiskFlag,omitempty"`
	Location                                  *Location                                      `json:"location,omitempty"`
	ManagedResourceFlag                       *bool                                          `json:"managedResourceFlag,omitempty"`
	MaxCpu                                    *int                                           `json:"maxCpu,omitempty"`
	MaxCpuUnits                               *string                                        `json:"maxCpuUnits,omitempty"`
	MaxMemory                                 *int                                           `json:"maxMemory,omitempty"`
	MetricPollDate                            *time.Time                                     `json:"metricPollDate,omitempty"`
	MetricTrackingObject                      *Metric_Tracking_Object                        `json:"metricTrackingObject,omitempty"`
	MetricTrackingObjectId                    *int                                           `json:"metricTrackingObjectId,omitempty"`
	ModifyDate                                *time.Time                                     `json:"modifyDate,omitempty"`
	MonitoringAgentCount                      *uint                                          `json:"monitoringAgentCount,omitempty"`
	MonitoringAgents                          []Monitoring_Agent                             `json:"monitoringAgents,omitempty"`
	MonitoringRobot                           *Monitoring_Robot                              `json:"monitoringRobot,omitempty"`
	MonitoringServiceComponent                *Network_Monitor_Version1_Query_Host_Stratum   `json:"monitoringServiceComponent,omitempty"`
	MonitoringServiceEligibilityFlag          *bool                                          `json:"monitoringServiceEligibilityFlag,omitempty"`
	MonitoringServiceFlag                     *bool                                          `json:"monitoringServiceFlag,omitempty"`
	MonitoringUserNotification                []User_Customer_Notification_Virtual_Guest     `json:"monitoringUserNotification,omitempty"`
	MonitoringUserNotificationCount           *uint                                          `json:"monitoringUserNotificationCount,omitempty"`
	NetworkComponentCount                     *uint                                          `json:"networkComponentCount,omitempty"`
	NetworkComponents                         []Virtual_Guest_Network_Component              `json:"networkComponents,omitempty"`
	NetworkMonitorCount                       *uint                                          `json:"networkMonitorCount,omitempty"`
	NetworkMonitorIncidentCount               *uint                                          `json:"networkMonitorIncidentCount,omitempty"`
	NetworkMonitorIncidents                   []Network_Monitor_Version1_Incident            `json:"networkMonitorIncidents,omitempty"`
	NetworkMonitors                           []Network_Monitor_Version1_Query_Host          `json:"networkMonitors,omitempty"`
	NetworkStorage                            []Network_Storage                              `json:"networkStorage,omitempty"`
	NetworkStorageCount                       *uint                                          `json:"networkStorageCount,omitempty"`
	NetworkVlanCount                          *uint                                          `json:"networkVlanCount,omitempty"`
	NetworkVlans                              []Network_Vlan                                 `json:"networkVlans,omitempty"`
	Notes                                     *string                                        `json:"notes,omitempty"`
	OpenCancellationTicket                    *Ticket                                        `json:"openCancellationTicket,omitempty"`
	OperatingSystem                           *Software_Component_OperatingSystem            `json:"operatingSystem,omitempty"`
	OperatingSystemReferenceCode              *string                                        `json:"operatingSystemReferenceCode,omitempty"`
	OrderedPackageId                          *string                                        `json:"orderedPackageId,omitempty"`
	OutboundPrivateBandwidthUsage             *float64                                       `json:"outboundPrivateBandwidthUsage,omitempty"`
	OutboundPublicBandwidthUsage              *float64                                       `json:"outboundPublicBandwidthUsage,omitempty"`
	OverBandwidthAllocationFlag               *int                                           `json:"overBandwidthAllocationFlag,omitempty"`
	PostInstallScriptUri                      *string                                        `json:"postInstallScriptUri,omitempty"`
	PowerState                                *Virtual_Guest_Power_State                     `json:"powerState,omitempty"`
	PrimaryBackendIpAddress                   *string                                        `json:"primaryBackendIpAddress,omitempty"`
	PrimaryBackendNetworkComponent            *Virtual_Guest_Network_Component               `json:"primaryBackendNetworkComponent,omitempty"`
	PrimaryIpAddress                          *string                                        `json:"primaryIpAddress,omitempty"`
	PrimaryNetworkComponent                   *Virtual_Guest_Network_Component               `json:"primaryNetworkComponent,omitempty"`
	PrivateNetworkOnlyFlag                    *bool                                          `json:"privateNetworkOnlyFlag,omitempty"`
	ProjectedOverBandwidthAllocationFlag      *int                                           `json:"projectedOverBandwidthAllocationFlag,omitempty"`
	ProjectedPublicBandwidthUsage             *float64                                       `json:"projectedPublicBandwidthUsage,omitempty"`
	ProvisionDate                             *time.Time                                     `json:"provisionDate,omitempty"`
	RecentEventCount                          *uint                                          `json:"recentEventCount,omitempty"`
	RecentEvents                              []Notification_Occurrence_Event                `json:"recentEvents,omitempty"`
	RegionalGroup                             *Location_Group_Regional                       `json:"regionalGroup,omitempty"`
	RegionalInternetRegistry                  *Network_Regional_Internet_Registry            `json:"regionalInternetRegistry,omitempty"`
	ScaleAssetCount                           *uint                                          `json:"scaleAssetCount,omitempty"`
	ScaleAssets                               []Scale_Asset                                  `json:"scaleAssets,omitempty"`
	ScaleMember                               *Scale_Member_Virtual_Guest                    `json:"scaleMember,omitempty"`
	ScaledFlag                                *bool                                          `json:"scaledFlag,omitempty"`
	SecurityScanRequestCount                  *uint                                          `json:"securityScanRequestCount,omitempty"`
	SecurityScanRequests                      []Network_Security_Scanner_Request             `json:"securityScanRequests,omitempty"`
	ServerRoom                                *Location                                      `json:"serverRoom,omitempty"`
	SoftwareComponentCount                    *uint                                          `json:"softwareComponentCount,omitempty"`
	SoftwareComponents                        []Software_Component                           `json:"softwareComponents,omitempty"`
	SshKeyCount                               *uint                                          `json:"sshKeyCount,omitempty"`
	SshKeys                                   []Security_Ssh_Key                             `json:"sshKeys,omitempty"`
	StartCpus                                 *int                                           `json:"startCpus,omitempty"`
	Status                                    *Virtual_Guest_Status                          `json:"status,omitempty"`
	StatusId                                  *int                                           `json:"statusId,omitempty"`
	SupplementalCreateObjectOptions           *Virtual_Guest_SupplementalCreateObjectOptions `json:"supplementalCreateObjectOptions,omitempty"`
	TagReferenceCount                         *uint                                          `json:"tagReferenceCount,omitempty"`
	TagReferences                             []Tag_Reference                                `json:"tagReferences,omitempty"`
	UpgradeRequest                            *Product_Upgrade_Request                       `json:"upgradeRequest,omitempty"`
	UserCount                                 *uint                                          `json:"userCount,omitempty"`
	UserData                                  []Virtual_Guest_Attribute                      `json:"userData,omitempty"`
	UserDataCount                             *uint                                          `json:"userDataCount,omitempty"`
	Users                                     []User_Customer                                `json:"users,omitempty"`
	Uuid                                      *string                                        `json:"uuid,omitempty"`
	VirtualRack                               *Network_Bandwidth_Version1_Allotment          `json:"virtualRack,omitempty"`
	VirtualRackId                             *int                                           `json:"virtualRackId,omitempty"`
	VirtualRackName                           *string                                        `json:"virtualRackName,omitempty"`
}

type Virtual_Guest_Attribute struct {
	Entity

	Guest *Virtual_Guest                `json:"guest,omitempty"`
	Type  *Virtual_Guest_Attribute_Type `json:"type,omitempty"`
	Value *string                       `json:"value,omitempty"`
}

type Virtual_Guest_Attribute_Type struct {
	Entity

	Keyname *string `json:"keyname,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type Virtual_Guest_Attribute_UserData struct {
	Virtual_Guest_Attribute
}

type Virtual_Guest_Block_Device struct {
	Entity

	BootableFlag *int                               `json:"bootableFlag,omitempty"`
	CreateDate   *time.Time                         `json:"createDate,omitempty"`
	Device       *string                            `json:"device,omitempty"`
	DiskImage    *Virtual_Disk_Image                `json:"diskImage,omitempty"`
	DiskImageId  *int                               `json:"diskImageId,omitempty"`
	Guest        *Virtual_Guest                     `json:"guest,omitempty"`
	GuestId      *int                               `json:"guestId,omitempty"`
	HotPlugFlag  *int                               `json:"hotPlugFlag,omitempty"`
	Id           *int                               `json:"id,omitempty"`
	ModifyDate   *time.Time                         `json:"modifyDate,omitempty"`
	MountMode    *string                            `json:"mountMode,omitempty"`
	MountType    *string                            `json:"mountType,omitempty"`
	Status       *Virtual_Guest_Block_Device_Status `json:"status,omitempty"`
	StatusId     *int                               `json:"statusId,omitempty"`
	Uuid         *string                            `json:"uuid,omitempty"`
}

type Virtual_Guest_Block_Device_Status struct {
	Entity

	KeyName *string `json:"keyName,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type Virtual_Guest_Block_Device_Template struct {
	Entity

	Device      *string                                    `json:"device,omitempty"`
	DiskImage   *Virtual_Disk_Image                        `json:"diskImage,omitempty"`
	DiskImageId *int                                       `json:"diskImageId,omitempty"`
	DiskSpace   *float64                                   `json:"diskSpace,omitempty"`
	Group       *Virtual_Guest_Block_Device_Template_Group `json:"group,omitempty"`
	GroupId     *int                                       `json:"groupId,omitempty"`
	Id          *int                                       `json:"id,omitempty"`
	Units       *string                                    `json:"units,omitempty"`
}

type Virtual_Guest_Block_Device_Template_Group struct {
	Entity

	Account                    *Account                                             `json:"account,omitempty"`
	AccountContactCount        *uint                                                `json:"accountContactCount,omitempty"`
	AccountContacts            []Account_Contact                                    `json:"accountContacts,omitempty"`
	AccountId                  *int                                                 `json:"accountId,omitempty"`
	AccountReferenceCount      *uint                                                `json:"accountReferenceCount,omitempty"`
	AccountReferences          []Virtual_Guest_Block_Device_Template_Group_Accounts `json:"accountReferences,omitempty"`
	BlockDeviceCount           *uint                                                `json:"blockDeviceCount,omitempty"`
	BlockDevices               []Virtual_Guest_Block_Device_Template                `json:"blockDevices,omitempty"`
	BlockDevicesDiskSpaceTotal *float64                                             `json:"blockDevicesDiskSpaceTotal,omitempty"`
	Children                   []Virtual_Guest_Block_Device_Template_Group          `json:"children,omitempty"`
	ChildrenCount              *uint                                                `json:"childrenCount,omitempty"`
	CreateDate                 *time.Time                                           `json:"createDate,omitempty"`
	Datacenter                 *Location                                            `json:"datacenter,omitempty"`
	DatacenterCount            *uint                                                `json:"datacenterCount,omitempty"`
	Datacenters                []Location                                           `json:"datacenters,omitempty"`
	FlexImageFlag              *bool                                                `json:"flexImageFlag,omitempty"`
	GlobalIdentifier           *string                                              `json:"globalIdentifier,omitempty"`
	Id                         *int                                                 `json:"id,omitempty"`
	ImageType                  *string                                              `json:"imageType,omitempty"`
	ImageTypeKeyName           *string                                              `json:"imageTypeKeyName,omitempty"`
	Name                       *string                                              `json:"name,omitempty"`
	Note                       *string                                              `json:"note,omitempty"`
	Parent                     *Virtual_Guest_Block_Device_Template_Group           `json:"parent,omitempty"`
	ParentId                   *int                                                 `json:"parentId,omitempty"`
	PublicFlag                 *int                                                 `json:"publicFlag,omitempty"`
	SshKeyCount                *uint                                                `json:"sshKeyCount,omitempty"`
	SshKeys                    []Security_Ssh_Key                                   `json:"sshKeys,omitempty"`
	Status                     *Virtual_Guest_Block_Device_Template_Group_Status    `json:"status,omitempty"`
	StatusId                   *int                                                 `json:"statusId,omitempty"`
	StorageRepository          *Virtual_Storage_Repository                          `json:"storageRepository,omitempty"`
	Summary                    *string                                              `json:"summary,omitempty"`
	TagReferenceCount          *uint                                                `json:"tagReferenceCount,omitempty"`
	TagReferences              []Tag_Reference                                      `json:"tagReferences,omitempty"`
	Transaction                *Provisioning_Version1_Transaction                   `json:"transaction,omitempty"`
	TransactionId              *int                                                 `json:"transactionId,omitempty"`
	UserRecordId               *int                                                 `json:"userRecordId,omitempty"`
}

type Virtual_Guest_Block_Device_Template_Group_Accounts struct {
	Entity

	Account    *Account                                   `json:"account,omitempty"`
	AccountId  *int                                       `json:"accountId,omitempty"`
	CreateDate *time.Time                                 `json:"createDate,omitempty"`
	Group      *Virtual_Guest_Block_Device_Template_Group `json:"group,omitempty"`
	GroupId    *int                                       `json:"groupId,omitempty"`
}

type Virtual_Guest_Block_Device_Template_Group_Status struct {
	Entity

	Description *string `json:"description,omitempty"`
	KeyName     *string `json:"keyName,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Virtual_Guest_Boot_Parameter struct {
	Entity

	CreateDate               *time.Time                         `json:"createDate,omitempty"`
	Guest                    *Virtual_Guest                     `json:"guest,omitempty"`
	GuestBootParameterType   *Virtual_Guest_Boot_Parameter_Type `json:"guestBootParameterType,omitempty"`
	GuestBootParameterTypeId *int                               `json:"guestBootParameterTypeId,omitempty"`
	GuestId                  *int                               `json:"guestId,omitempty"`
	Id                       *int                               `json:"id,omitempty"`
	ModifyDate               *time.Time                         `json:"modifyDate,omitempty"`
}

type Virtual_Guest_Boot_Parameter_Type struct {
	Entity

	BootOption  *string    `json:"bootOption,omitempty"`
	CreateDate  *time.Time `json:"createDate,omitempty"`
	Description *string    `json:"description,omitempty"`
	Id          *int       `json:"id,omitempty"`
	KeyName     *string    `json:"keyName,omitempty"`
	ModifyDate  *time.Time `json:"modifyDate,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Value       *string    `json:"value,omitempty"`
}

type Virtual_Guest_Network_Component struct {
	Entity

	CreateDate                     *time.Time                                  `json:"createDate,omitempty"`
	Guest                          *Virtual_Guest                              `json:"guest,omitempty"`
	GuestId                        *int                                        `json:"guestId,omitempty"`
	HighAvailabilityFirewallFlag   *bool                                       `json:"highAvailabilityFirewallFlag,omitempty"`
	Id                             *int                                        `json:"id,omitempty"`
	IpAddressBindingCount          *uint                                       `json:"ipAddressBindingCount,omitempty"`
	IpAddressBindings              []Virtual_Guest_Network_Component_IpAddress `json:"ipAddressBindings,omitempty"`
	MacAddress                     *string                                     `json:"macAddress,omitempty"`
	MaxSpeed                       *int                                        `json:"maxSpeed,omitempty"`
	ModifyDate                     *time.Time                                  `json:"modifyDate,omitempty"`
	Name                           *string                                     `json:"name,omitempty"`
	NetworkComponentFirewall       *Network_Component_Firewall                 `json:"networkComponentFirewall,omitempty"`
	NetworkId                      *int                                        `json:"networkId,omitempty"`
	NetworkVlan                    *Network_Vlan                               `json:"networkVlan,omitempty"`
	Port                           *int                                        `json:"port,omitempty"`
	PrimaryIpAddress               *string                                     `json:"primaryIpAddress,omitempty"`
	PrimaryIpAddressRecord         *Network_Subnet_IpAddress                   `json:"primaryIpAddressRecord,omitempty"`
	PrimarySubnet                  *Network_Subnet                             `json:"primarySubnet,omitempty"`
	PrimaryVersion6IpAddressRecord *Network_Subnet_IpAddress                   `json:"primaryVersion6IpAddressRecord,omitempty"`
	Router                         *Hardware_Router                            `json:"router,omitempty"`
	Speed                          *int                                        `json:"speed,omitempty"`
	Status                         *string                                     `json:"status,omitempty"`
	SubnetCount                    *uint                                       `json:"subnetCount,omitempty"`
	Subnets                        []Network_Subnet                            `json:"subnets,omitempty"`
	Uuid                           *string                                     `json:"uuid,omitempty"`
}

type Virtual_Guest_Network_Component_IpAddress struct {
	Entity

	IpAddress        *Network_Subnet_IpAddress        `json:"ipAddress,omitempty"`
	IpAddressId      *int                             `json:"ipAddressId,omitempty"`
	NetworkComponent *Virtual_Guest_Network_Component `json:"networkComponent,omitempty"`
	Port             *int                             `json:"port,omitempty"`
	Type             *string                          `json:"type,omitempty"`
}

type Virtual_Guest_Power_State struct {
	Entity

	Description *string `json:"description,omitempty"`
	KeyName     *string `json:"keyName,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Virtual_Guest_Status struct {
	Entity

	KeyName *string `json:"keyName,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type Virtual_Guest_SupplementalCreateObjectOptions struct {
	Entity

	ImmediateApprovalOnlyFlag *bool   `json:"immediateApprovalOnlyFlag,omitempty"`
	PostInstallScriptUri      *string `json:"postInstallScriptUri,omitempty"`
}

type Virtual_Host struct {
	Entity

	Account                  *Account                `json:"account,omitempty"`
	AccountId                *int                    `json:"accountId,omitempty"`
	BilledPerGuestFlag       *bool                   `json:"billedPerGuestFlag,omitempty"`
	BilledPerMemoryUsageFlag *bool                   `json:"billedPerMemoryUsageFlag,omitempty"`
	CreateDate               *time.Time              `json:"createDate,omitempty"`
	Description              *string                 `json:"description,omitempty"`
	EnabledFlag              *int                    `json:"enabledFlag,omitempty"`
	GuestCount               *uint                   `json:"guestCount,omitempty"`
	Guests                   []Virtual_Guest         `json:"guests,omitempty"`
	Hardware                 *Hardware_Server        `json:"hardware,omitempty"`
	HardwareId               *int                    `json:"hardwareId,omitempty"`
	Id                       *int                    `json:"id,omitempty"`
	MetricTrackingObject     *Metric_Tracking_Object `json:"metricTrackingObject,omitempty"`
	ModifyDate               *time.Time              `json:"modifyDate,omitempty"`
	Name                     *string                 `json:"name,omitempty"`
	PhysicalMemoryCapacity   *int                    `json:"physicalMemoryCapacity,omitempty"`
	Uuid                     *string                 `json:"uuid,omitempty"`
}

type Virtual_Storage_Repository struct {
	Entity

	Account                *Account                                           `json:"account,omitempty"`
	BillingItem            *Billing_Item                                      `json:"billingItem,omitempty"`
	Capacity               *float64                                           `json:"capacity,omitempty"`
	Datacenter             *Location                                          `json:"datacenter,omitempty"`
	Description            *string                                            `json:"description,omitempty"`
	DiskImageCount         *uint                                              `json:"diskImageCount,omitempty"`
	DiskImages             []Virtual_Disk_Image                               `json:"diskImages,omitempty"`
	GuestCount             *uint                                              `json:"guestCount,omitempty"`
	Guests                 []Virtual_Guest                                    `json:"guests,omitempty"`
	Id                     *int                                               `json:"id,omitempty"`
	MetricTrackingObject   *Metric_Tracking_Object_Virtual_Storage_Repository `json:"metricTrackingObject,omitempty"`
	Name                   *string                                            `json:"name,omitempty"`
	PublicFlag             *int                                               `json:"publicFlag,omitempty"`
	PublicImageBillingItem *Billing_Item                                      `json:"publicImageBillingItem,omitempty"`
	Type                   *Virtual_Storage_Repository_Type                   `json:"type,omitempty"`
	TypeId                 *int                                               `json:"typeId,omitempty"`
}

type Virtual_Storage_Repository_Type struct {
	Entity

	Description            *string                      `json:"description,omitempty"`
	Name                   *string                      `json:"name,omitempty"`
	StorageRepositories    []Virtual_Storage_Repository `json:"storageRepositories,omitempty"`
	StorageRepositoryCount *uint                        `json:"storageRepositoryCount,omitempty"`
}
