package datatypes

import "time"

type Network struct {
	Entity

	AccountId         *int             `json:"accountId,omitempty"`
	Cidr              *int             `json:"cidr,omitempty"`
	Id                *int             `json:"id,omitempty"`
	Name              *string          `json:"name,omitempty"`
	NetworkIdentifier *string          `json:"networkIdentifier,omitempty"`
	Notes             *string          `json:"notes,omitempty"`
	SubnetCount       *uint            `json:"subnetCount,omitempty"`
	Subnets           []Network_Subnet `json:"subnets,omitempty"`
}

type Network_Application_Delivery_Controller struct {
	Entity

	Account                          *Account                                                                `json:"account,omitempty"`
	AccountId                        *int                                                                    `json:"accountId,omitempty"`
	AverageDailyPublicBandwidthUsage *float64                                                                `json:"averageDailyPublicBandwidthUsage,omitempty"`
	BillingItem                      *Billing_Item_Network_Application_Delivery_Controller                   `json:"billingItem,omitempty"`
	ConfigurationHistory             []Network_Application_Delivery_Controller_Configuration_History         `json:"configurationHistory,omitempty"`
	ConfigurationHistoryCount        *uint                                                                   `json:"configurationHistoryCount,omitempty"`
	CreateDate                       *time.Time                                                              `json:"createDate,omitempty"`
	Datacenter                       *Location                                                               `json:"datacenter,omitempty"`
	Description                      *string                                                                 `json:"description,omitempty"`
	Id                               *int                                                                    `json:"id,omitempty"`
	LicenseExpirationDate            *time.Time                                                              `json:"licenseExpirationDate,omitempty"`
	LoadBalancerCount                *uint                                                                   `json:"loadBalancerCount,omitempty"`
	LoadBalancers                    []Network_LoadBalancer_VirtualIpAddress                                 `json:"loadBalancers,omitempty"`
	ManagedResourceFlag              *bool                                                                   `json:"managedResourceFlag,omitempty"`
	ManagementIpAddress              *string                                                                 `json:"managementIpAddress,omitempty"`
	ModifyDate                       *time.Time                                                              `json:"modifyDate,omitempty"`
	Name                             *string                                                                 `json:"name,omitempty"`
	NetworkVlan                      *Network_Vlan                                                           `json:"networkVlan,omitempty"`
	NetworkVlanCount                 *uint                                                                   `json:"networkVlanCount,omitempty"`
	NetworkVlans                     []Network_Vlan                                                          `json:"networkVlans,omitempty"`
	Notes                            *string                                                                 `json:"notes,omitempty"`
	OutboundPublicBandwidthUsage     *float64                                                                `json:"outboundPublicBandwidthUsage,omitempty"`
	Password                         *Software_Component_Password                                            `json:"password,omitempty"`
	PrimaryIpAddress                 *string                                                                 `json:"primaryIpAddress,omitempty"`
	ProjectedPublicBandwidthUsage    *float64                                                                `json:"projectedPublicBandwidthUsage,omitempty"`
	SubnetCount                      *uint                                                                   `json:"subnetCount,omitempty"`
	Subnets                          []Network_Subnet                                                        `json:"subnets,omitempty"`
	TagReferenceCount                *uint                                                                   `json:"tagReferenceCount,omitempty"`
	TagReferences                    []Tag_Reference                                                         `json:"tagReferences,omitempty"`
	Type                             *Network_Application_Delivery_Controller_Type                           `json:"type,omitempty"`
	TypeId                           *int                                                                    `json:"typeId,omitempty"`
	VirtualIpAddressCount            *uint                                                                   `json:"virtualIpAddressCount,omitempty"`
	VirtualIpAddresses               []Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"virtualIpAddresses,omitempty"`
}

type Network_Application_Delivery_Controller_Configuration_History struct {
	Entity

	Controller *Network_Application_Delivery_Controller `json:"controller,omitempty"`
	CreateDate *time.Time                               `json:"createDate,omitempty"`
	Id         *int                                     `json:"id,omitempty"`
	Notes      *string                                  `json:"notes,omitempty"`
}

type Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute struct {
	Entity

	HealthAttributeTypeId *int                                                                        `json:"healthAttributeTypeId,omitempty"`
	HealthCheck           *Network_Application_Delivery_Controller_LoadBalancer_Health_Check          `json:"healthCheck,omitempty"`
	HealthCheckId         *int                                                                        `json:"healthCheckId,omitempty"`
	Id                    *int                                                                        `json:"id,omitempty"`
	Type                  *Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type `json:"type,omitempty"`
	Value                 *string                                                                     `json:"value,omitempty"`
}

type Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type struct {
	Entity

	Description     *string `json:"description,omitempty"`
	Id              *int    `json:"id,omitempty"`
	Keyname         *string `json:"keyname,omitempty"`
	Name            *string `json:"name,omitempty"`
	ValueExpression *string `json:"valueExpression,omitempty"`
}

type Network_Application_Delivery_Controller_LoadBalancer_Health_Check struct {
	Entity

	AttributeCount         *uint                                                                   `json:"attributeCount,omitempty"`
	Attributes             []Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute `json:"attributes,omitempty"`
	HealthCheckTypeId      *int                                                                    `json:"healthCheckTypeId,omitempty"`
	Id                     *int                                                                    `json:"id,omitempty"`
	Name                   *string                                                                 `json:"name,omitempty"`
	Notes                  *string                                                                 `json:"notes,omitempty"`
	ScaleLoadBalancerCount *uint                                                                   `json:"scaleLoadBalancerCount,omitempty"`
	ScaleLoadBalancers     []Scale_LoadBalancer                                                    `json:"scaleLoadBalancers,omitempty"`
	ServiceCount           *uint                                                                   `json:"serviceCount,omitempty"`
	Services               []Network_Application_Delivery_Controller_LoadBalancer_Service          `json:"services,omitempty"`
	Type                   *Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type `json:"type,omitempty"`
}

type Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type struct {
	Entity

	Id      *int    `json:"id,omitempty"`
	Keyname *string `json:"keyname,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type Network_Application_Delivery_Controller_LoadBalancer_Routing_Method struct {
	Entity

	Id      *int    `json:"id,omitempty"`
	Keyname *string `json:"keyname,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type Network_Application_Delivery_Controller_LoadBalancer_Routing_Type struct {
	Entity

	Id      *int    `json:"id,omitempty"`
	Keyname *string `json:"keyname,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type Network_Application_Delivery_Controller_LoadBalancer_Service struct {
	Entity

	Enabled             *int                                                                                `json:"enabled,omitempty"`
	GroupCount          *uint                                                                               `json:"groupCount,omitempty"`
	GroupReferenceCount *uint                                                                               `json:"groupReferenceCount,omitempty"`
	GroupReferences     []Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference `json:"groupReferences,omitempty"`
	Groups              []Network_Application_Delivery_Controller_LoadBalancer_Service_Group                `json:"groups,omitempty"`
	HealthCheck         *Network_Application_Delivery_Controller_LoadBalancer_Health_Check                  `json:"healthCheck,omitempty"`
	HealthCheckCount    *uint                                                                               `json:"healthCheckCount,omitempty"`
	HealthChecks        []Network_Application_Delivery_Controller_LoadBalancer_Health_Check                 `json:"healthChecks,omitempty"`
	Id                  *int                                                                                `json:"id,omitempty"`
	IpAddress           *Network_Subnet_IpAddress                                                           `json:"ipAddress,omitempty"`
	IpAddressId         *int                                                                                `json:"ipAddressId,omitempty"`
	Name                *string                                                                             `json:"name,omitempty"`
	Notes               *string                                                                             `json:"notes,omitempty"`
	Port                *int                                                                                `json:"port,omitempty"`
	ServiceGroup        *Network_Application_Delivery_Controller_LoadBalancer_Service_Group                 `json:"serviceGroup,omitempty"`
	Status              *string                                                                             `json:"status,omitempty"`
}

type Network_Application_Delivery_Controller_LoadBalancer_Service_Group struct {
	Entity

	Id                    *int                                                                                `json:"id,omitempty"`
	Name                  *string                                                                             `json:"name,omitempty"`
	Notes                 *string                                                                             `json:"notes,omitempty"`
	RoutingMethod         *Network_Application_Delivery_Controller_LoadBalancer_Routing_Method                `json:"routingMethod,omitempty"`
	RoutingMethodId       *int                                                                                `json:"routingMethodId,omitempty"`
	RoutingType           *Network_Application_Delivery_Controller_LoadBalancer_Routing_Type                  `json:"routingType,omitempty"`
	RoutingTypeId         *int                                                                                `json:"routingTypeId,omitempty"`
	ServiceCount          *uint                                                                               `json:"serviceCount,omitempty"`
	ServiceReferenceCount *uint                                                                               `json:"serviceReferenceCount,omitempty"`
	ServiceReferences     []Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference `json:"serviceReferences,omitempty"`
	Services              []Network_Application_Delivery_Controller_LoadBalancer_Service                      `json:"services,omitempty"`
	Timeout               *int                                                                                `json:"timeout,omitempty"`
	VirtualServer         *Network_Application_Delivery_Controller_LoadBalancer_VirtualServer                 `json:"virtualServer,omitempty"`
	VirtualServerCount    *uint                                                                               `json:"virtualServerCount,omitempty"`
	VirtualServers        []Network_Application_Delivery_Controller_LoadBalancer_VirtualServer                `json:"virtualServers,omitempty"`
}

type Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference struct {
	Entity

	Service        *Network_Application_Delivery_Controller_LoadBalancer_Service       `json:"service,omitempty"`
	ServiceGroup   *Network_Application_Delivery_Controller_LoadBalancer_Service_Group `json:"serviceGroup,omitempty"`
	ServiceGroupId *int                                                                `json:"serviceGroupId,omitempty"`
	ServiceId      *int                                                                `json:"serviceId,omitempty"`
	Weight         *int                                                                `json:"weight,omitempty"`
}

type Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress struct {
	Entity

	Account                            *Account                                                                                        `json:"account,omitempty"`
	AccountId                          *int                                                                                            `json:"accountId,omitempty"`
	ApplicationDeliveryController      *Network_Application_Delivery_Controller                                                        `json:"applicationDeliveryController,omitempty"`
	ApplicationDeliveryControllerCount *uint                                                                                           `json:"applicationDeliveryControllerCount,omitempty"`
	ApplicationDeliveryControllers     []Network_Application_Delivery_Controller                                                       `json:"applicationDeliveryControllers,omitempty"`
	BillingItem                        *Billing_Item                                                                                   `json:"billingItem,omitempty"`
	ConnectionLimit                    *int                                                                                            `json:"connectionLimit,omitempty"`
	ConnectionLimitUnits               *string                                                                                         `json:"connectionLimitUnits,omitempty"`
	DedicatedBillingItem               *Billing_Item_Network_LoadBalancer                                                              `json:"dedicatedBillingItem,omitempty"`
	DedicatedFlag                      *bool                                                                                           `json:"dedicatedFlag,omitempty"`
	HighAvailabilityFlag               *bool                                                                                           `json:"highAvailabilityFlag,omitempty"`
	Id                                 *int                                                                                            `json:"id,omitempty"`
	IpAddress                          *Network_Subnet_IpAddress                                                                       `json:"ipAddress,omitempty"`
	IpAddressId                        *int                                                                                            `json:"ipAddressId,omitempty"`
	LoadBalancerHardware               []Hardware                                                                                      `json:"loadBalancerHardware,omitempty"`
	LoadBalancerHardwareCount          *uint                                                                                           `json:"loadBalancerHardwareCount,omitempty"`
	ManagedResourceFlag                *bool                                                                                           `json:"managedResourceFlag,omitempty"`
	Notes                              *string                                                                                         `json:"notes,omitempty"`
	SecureTransportCipherCount         *uint                                                                                           `json:"secureTransportCipherCount,omitempty"`
	SecureTransportCiphers             []Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportCipher   `json:"secureTransportCiphers,omitempty"`
	SecureTransportProtocolCount       *uint                                                                                           `json:"secureTransportProtocolCount,omitempty"`
	SecureTransportProtocols           []Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportProtocol `json:"secureTransportProtocols,omitempty"`
	SecurityCertificate                *Security_Certificate                                                                           `json:"securityCertificate,omitempty"`
	SecurityCertificateEntry           *Security_Certificate_Entry                                                                     `json:"securityCertificateEntry,omitempty"`
	SecurityCertificateId              *int                                                                                            `json:"securityCertificateId,omitempty"`
	SslActiveFlag                      *bool                                                                                           `json:"sslActiveFlag,omitempty"`
	SslEnabledFlag                     *bool                                                                                           `json:"sslEnabledFlag,omitempty"`
	VirtualServerCount                 *uint                                                                                           `json:"virtualServerCount,omitempty"`
	VirtualServers                     []Network_Application_Delivery_Controller_LoadBalancer_VirtualServer                            `json:"virtualServers,omitempty"`
}

type Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportCipher struct {
	Entity

	Id                 *int                                                                   `json:"id,omitempty"`
	KeyName            *string                                                                `json:"keyName,omitempty"`
	VirtualIpAddress   *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"virtualIpAddress,omitempty"`
	VirtualIpAddressId *int                                                                   `json:"virtualIpAddressId,omitempty"`
}

type Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportProtocol struct {
	Entity

	Id                 *int                                                                   `json:"id,omitempty"`
	KeyName            *string                                                                `json:"keyName,omitempty"`
	VirtualIpAddress   *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"virtualIpAddress,omitempty"`
	VirtualIpAddressId *int                                                                   `json:"virtualIpAddressId,omitempty"`
}

type Network_Application_Delivery_Controller_LoadBalancer_VirtualServer struct {
	Entity

	Allocation             *int                                                                   `json:"allocation,omitempty"`
	Id                     *int                                                                   `json:"id,omitempty"`
	Name                   *string                                                                `json:"name,omitempty"`
	Notes                  *string                                                                `json:"notes,omitempty"`
	Port                   *int                                                                   `json:"port,omitempty"`
	RoutingMethod          *Network_Application_Delivery_Controller_LoadBalancer_Routing_Method   `json:"routingMethod,omitempty"`
	RoutingMethodId        *int                                                                   `json:"routingMethodId,omitempty"`
	ScaleLoadBalancerCount *uint                                                                  `json:"scaleLoadBalancerCount,omitempty"`
	ScaleLoadBalancers     []Scale_LoadBalancer                                                   `json:"scaleLoadBalancers,omitempty"`
	ServiceGroupCount      *uint                                                                  `json:"serviceGroupCount,omitempty"`
	ServiceGroups          []Network_Application_Delivery_Controller_LoadBalancer_Service_Group   `json:"serviceGroups,omitempty"`
	VirtualIpAddress       *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"virtualIpAddress,omitempty"`
	VirtualIpAddressId     *int                                                                   `json:"virtualIpAddressId,omitempty"`
}

type Network_Application_Delivery_Controller_Type struct {
	Entity

	KeyName *string `json:"keyName,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type Network_Backbone struct {
	Entity

	Capacity           *int               `json:"capacity,omitempty"`
	CapacityUnits      *string            `json:"capacityUnits,omitempty"`
	Health             *string            `json:"health,omitempty"`
	Id                 *int               `json:"id,omitempty"`
	Location           *Location          `json:"location,omitempty"`
	Name               *string            `json:"name,omitempty"`
	NetworkComponent   *Network_Component `json:"networkComponent,omitempty"`
	NetworkComponentId *int               `json:"networkComponentId,omitempty"`
	Type               *string            `json:"type,omitempty"`
}

type Network_Backbone_Location_Dependent struct {
	Entity

	DependentLocation   *Location `json:"dependentLocation,omitempty"`
	DependentLocationId *int      `json:"dependentLocationId,omitempty"`
	Id                  *int      `json:"id,omitempty"`
	SourceLocation      *Location `json:"sourceLocation,omitempty"`
	SourceLocationId    *int      `json:"sourceLocationId,omitempty"`
}

type Network_Bandwidth_Usage struct {
	Entity

	AmountIn                   *float64                                      `json:"amountIn,omitempty"`
	AmountOut                  *float64                                      `json:"amountOut,omitempty"`
	BandwidthUsageDetailTypeId *float64                                      `json:"bandwidthUsageDetailTypeId,omitempty"`
	TrackingObject             *Metric_Tracking_Object                       `json:"trackingObject,omitempty"`
	Type                       *Network_Bandwidth_Version1_Usage_Detail_Type `json:"type,omitempty"`
}

type Network_Bandwidth_Usage_Detail struct {
	Entity

	Account                    *Account                                      `json:"account,omitempty"`
	AmountIn                   *float64                                      `json:"amountIn,omitempty"`
	AmountOut                  *float64                                      `json:"amountOut,omitempty"`
	BandwidthUsageDetailTypeId *float64                                      `json:"bandwidthUsageDetailTypeId,omitempty"`
	TrackingObject             *Metric_Tracking_Object                       `json:"trackingObject,omitempty"`
	Type                       *Network_Bandwidth_Version1_Usage_Detail_Type `json:"type,omitempty"`
}

type Network_Bandwidth_Version1_Allocation struct {
	Entity

	AllotmentDetail *Network_Bandwidth_Version1_Allotment_Detail `json:"allotmentDetail,omitempty"`
	Amount          *float64                                     `json:"amount,omitempty"`
	BillingItem     *Billing_Item_Hardware                       `json:"billingItem,omitempty"`
	Id              *int                                         `json:"id,omitempty"`
}

type Network_Bandwidth_Version1_Allotment struct {
	Entity

	Account                              *Account                                      `json:"account,omitempty"`
	AccountId                            *int                                          `json:"accountId,omitempty"`
	ActiveDetailCount                    *uint                                         `json:"activeDetailCount,omitempty"`
	ActiveDetails                        []Network_Bandwidth_Version1_Allotment_Detail `json:"activeDetails,omitempty"`
	ApplicationDeliveryControllerCount   *uint                                         `json:"applicationDeliveryControllerCount,omitempty"`
	ApplicationDeliveryControllers       []Network_Application_Delivery_Controller     `json:"applicationDeliveryControllers,omitempty"`
	AverageDailyPublicBandwidthUsage     *float64                                      `json:"averageDailyPublicBandwidthUsage,omitempty"`
	BandwidthAllotmentTypeId             *int                                          `json:"bandwidthAllotmentTypeId,omitempty"`
	BareMetalInstanceCount               *uint                                         `json:"bareMetalInstanceCount,omitempty"`
	BareMetalInstances                   []Hardware                                    `json:"bareMetalInstances,omitempty"`
	BillingCycleBandwidthUsage           []Network_Bandwidth_Usage                     `json:"billingCycleBandwidthUsage,omitempty"`
	BillingCycleBandwidthUsageCount      *uint                                         `json:"billingCycleBandwidthUsageCount,omitempty"`
	BillingCyclePrivateBandwidthUsage    *Network_Bandwidth_Usage                      `json:"billingCyclePrivateBandwidthUsage,omitempty"`
	BillingCyclePublicBandwidthUsage     *Network_Bandwidth_Usage                      `json:"billingCyclePublicBandwidthUsage,omitempty"`
	BillingCyclePublicUsageTotal         *uint                                         `json:"billingCyclePublicUsageTotal,omitempty"`
	BillingItem                          *Billing_Item                                 `json:"billingItem,omitempty"`
	CreateDate                           *time.Time                                    `json:"createDate,omitempty"`
	CurrentBandwidthSummary              *Metric_Tracking_Object_Bandwidth_Summary     `json:"currentBandwidthSummary,omitempty"`
	DetailCount                          *uint                                         `json:"detailCount,omitempty"`
	Details                              []Network_Bandwidth_Version1_Allotment_Detail `json:"details,omitempty"`
	EndDate                              *time.Time                                    `json:"endDate,omitempty"`
	Hardware                             []Hardware                                    `json:"hardware,omitempty"`
	HardwareCount                        *uint                                         `json:"hardwareCount,omitempty"`
	Id                                   *int                                          `json:"id,omitempty"`
	InboundPublicBandwidthUsage          *float64                                      `json:"inboundPublicBandwidthUsage,omitempty"`
	LocationGroup                        *Location_Group                               `json:"locationGroup,omitempty"`
	LocationGroupId                      *int                                          `json:"locationGroupId,omitempty"`
	ManagedBareMetalInstanceCount        *uint                                         `json:"managedBareMetalInstanceCount,omitempty"`
	ManagedBareMetalInstances            []Hardware                                    `json:"managedBareMetalInstances,omitempty"`
	ManagedHardware                      []Hardware                                    `json:"managedHardware,omitempty"`
	ManagedHardwareCount                 *uint                                         `json:"managedHardwareCount,omitempty"`
	ManagedVirtualGuestCount             *uint                                         `json:"managedVirtualGuestCount,omitempty"`
	ManagedVirtualGuests                 []Virtual_Guest                               `json:"managedVirtualGuests,omitempty"`
	MetricTrackingObject                 *Metric_Tracking_Object_VirtualDedicatedRack  `json:"metricTrackingObject,omitempty"`
	MetricTrackingObjectId               *int                                          `json:"metricTrackingObjectId,omitempty"`
	Name                                 *string                                       `json:"name,omitempty"`
	OutboundPublicBandwidthUsage         *float64                                      `json:"outboundPublicBandwidthUsage,omitempty"`
	OverBandwidthAllocationFlag          *int                                          `json:"overBandwidthAllocationFlag,omitempty"`
	PrivateNetworkOnlyHardware           []Hardware                                    `json:"privateNetworkOnlyHardware,omitempty"`
	PrivateNetworkOnlyHardwareCount      *uint                                         `json:"privateNetworkOnlyHardwareCount,omitempty"`
	ProjectedOverBandwidthAllocationFlag *int                                          `json:"projectedOverBandwidthAllocationFlag,omitempty"`
	ProjectedPublicBandwidthUsage        *float64                                      `json:"projectedPublicBandwidthUsage,omitempty"`
	ServiceProvider                      *Service_Provider                             `json:"serviceProvider,omitempty"`
	ServiceProviderId                    *int                                          `json:"serviceProviderId,omitempty"`
	TotalBandwidthAllocated              *uint                                         `json:"totalBandwidthAllocated,omitempty"`
	VirtualGuestCount                    *uint                                         `json:"virtualGuestCount,omitempty"`
	VirtualGuests                        []Virtual_Guest                               `json:"virtualGuests,omitempty"`
}

type Network_Bandwidth_Version1_Allotment_Detail struct {
	Entity

	Allocation           *Network_Bandwidth_Version1_Allocation `json:"allocation,omitempty"`
	AllocationId         *int                                   `json:"allocationId,omitempty"`
	BandwidthAllotment   *Network_Bandwidth_Version1_Allotment  `json:"bandwidthAllotment,omitempty"`
	BandwidthAllotmentId *int                                   `json:"bandwidthAllotmentId,omitempty"`
	BandwidthUsage       []Network_Bandwidth_Version1_Usage     `json:"bandwidthUsage,omitempty"`
	BandwidthUsageCount  *uint                                  `json:"bandwidthUsageCount,omitempty"`
	EffectiveDate        *time.Time                             `json:"effectiveDate,omitempty"`
	EndEffectiveDate     *time.Time                             `json:"endEffectiveDate,omitempty"`
	Id                   *int                                   `json:"id,omitempty"`
	ServiceProviderId    *int                                   `json:"serviceProviderId,omitempty"`
}

type Network_Bandwidth_Version1_Host struct {
	Entity

	PodId *int `json:"podId,omitempty"`
}

type Network_Bandwidth_Version1_Interface struct {
	Entity

	Host               *Network_Bandwidth_Version1_Host `json:"host,omitempty"`
	HostId             *int                             `json:"hostId,omitempty"`
	NetworkComponent   *Network_Component               `json:"networkComponent,omitempty"`
	NetworkComponentId *int                             `json:"networkComponentId,omitempty"`
}

type Network_Bandwidth_Version1_Usage struct {
	Entity

	BandwidthAllotmentDetail  *Network_Bandwidth_Version1_Allotment_Detail `json:"bandwidthAllotmentDetail,omitempty"`
	BandwidthUsageDetail      []Network_Bandwidth_Version1_Usage_Detail    `json:"bandwidthUsageDetail,omitempty"`
	BandwidthUsageDetailCount *uint                                        `json:"bandwidthUsageDetailCount,omitempty"`
}

type Network_Bandwidth_Version1_Usage_Detail struct {
	Entity

	AmountIn                 *float64                                      `json:"amountIn,omitempty"`
	AmountOut                *float64                                      `json:"amountOut,omitempty"`
	BandwidthUsage           *Network_Bandwidth_Version1_Usage             `json:"bandwidthUsage,omitempty"`
	BandwidthUsageDetailType *Network_Bandwidth_Version1_Usage_Detail_Type `json:"bandwidthUsageDetailType,omitempty"`
	Day                      *time.Time                                    `json:"day,omitempty"`
}

type Network_Bandwidth_Version1_Usage_Detail_Total struct {
	Entity

	Account                    *Account                                      `json:"account,omitempty"`
	AmountIn                   *float64                                      `json:"amountIn,omitempty"`
	AmountOut                  *float64                                      `json:"amountOut,omitempty"`
	BandwidthUsageDetailTypeId *float64                                      `json:"bandwidthUsageDetailTypeId,omitempty"`
	TrackingObject             *Metric_Tracking_Object                       `json:"trackingObject,omitempty"`
	Type                       *Network_Bandwidth_Version1_Usage_Detail_Type `json:"type,omitempty"`
}

type Network_Bandwidth_Version1_Usage_Detail_Type struct {
	Entity

	Alias *string `json:"alias,omitempty"`
}

type Network_Component struct {
	Entity

	ActiveCommand                  *Hardware_Component_RemoteManagement_Command_Request  `json:"activeCommand,omitempty"`
	DownlinkComponent              *Network_Component                                    `json:"downlinkComponent,omitempty"`
	DuplexMode                     *Network_Component_Duplex_Mode                        `json:"duplexMode,omitempty"`
	DuplexModeId                   *string                                               `json:"duplexModeId,omitempty"`
	Hardware                       *Hardware                                             `json:"hardware,omitempty"`
	HardwareId                     *int                                                  `json:"hardwareId,omitempty"`
	HighAvailabilityFirewallFlag   *bool                                                 `json:"highAvailabilityFirewallFlag,omitempty"`
	Id                             *int                                                  `json:"id,omitempty"`
	Interface                      *Network_Bandwidth_Version1_Interface                 `json:"interface,omitempty"`
	IpAddressBindingCount          *uint                                                 `json:"ipAddressBindingCount,omitempty"`
	IpAddressBindings              []Network_Component_IpAddress                         `json:"ipAddressBindings,omitempty"`
	IpAddressCount                 *uint                                                 `json:"ipAddressCount,omitempty"`
	IpAddresses                    []Network_Subnet_IpAddress                            `json:"ipAddresses,omitempty"`
	IpmiIpAddress                  *string                                               `json:"ipmiIpAddress,omitempty"`
	IpmiMacAddress                 *string                                               `json:"ipmiMacAddress,omitempty"`
	LastCommand                    *Hardware_Component_RemoteManagement_Command_Request  `json:"lastCommand,omitempty"`
	MacAddress                     *string                                               `json:"macAddress,omitempty"`
	MaxSpeed                       *int                                                  `json:"maxSpeed,omitempty"`
	MetricTrackingObject           *Metric_Tracking_Object                               `json:"metricTrackingObject,omitempty"`
	ModifyDate                     *time.Time                                            `json:"modifyDate,omitempty"`
	Name                           *string                                               `json:"name,omitempty"`
	NetworkComponentFirewall       *Network_Component_Firewall                           `json:"networkComponentFirewall,omitempty"`
	NetworkComponentGroup          *Network_Component_Group                              `json:"networkComponentGroup,omitempty"`
	NetworkHardware                []Hardware                                            `json:"networkHardware,omitempty"`
	NetworkHardwareCount           *uint                                                 `json:"networkHardwareCount,omitempty"`
	NetworkVlan                    *Network_Vlan                                         `json:"networkVlan,omitempty"`
	NetworkVlanId                  *int                                                  `json:"networkVlanId,omitempty"`
	NetworkVlanTrunkCount          *uint                                                 `json:"networkVlanTrunkCount,omitempty"`
	NetworkVlanTrunks              []Network_Component_Network_Vlan_Trunk                `json:"networkVlanTrunks,omitempty"`
	Port                           *int                                                  `json:"port,omitempty"`
	PrimaryIpAddress               *string                                               `json:"primaryIpAddress,omitempty"`
	PrimaryIpAddressRecord         *Network_Subnet_IpAddress                             `json:"primaryIpAddressRecord,omitempty"`
	PrimarySubnet                  *Network_Subnet                                       `json:"primarySubnet,omitempty"`
	PrimaryVersion6IpAddressRecord *Network_Subnet_IpAddress                             `json:"primaryVersion6IpAddressRecord,omitempty"`
	RecentCommandCount             *uint                                                 `json:"recentCommandCount,omitempty"`
	RecentCommands                 []Hardware_Component_RemoteManagement_Command_Request `json:"recentCommands,omitempty"`
	RedundancyCapableFlag          *bool                                                 `json:"redundancyCapableFlag,omitempty"`
	RedundancyEnabledFlag          *bool                                                 `json:"redundancyEnabledFlag,omitempty"`
	RemoteManagementUserCount      *uint                                                 `json:"remoteManagementUserCount,omitempty"`
	RemoteManagementUsers          []Hardware_Component_RemoteManagement_User            `json:"remoteManagementUsers,omitempty"`
	Router                         *Hardware                                             `json:"router,omitempty"`
	Speed                          *int                                                  `json:"speed,omitempty"`
	Status                         *string                                               `json:"status,omitempty"`
	StorageNetworkFlag             *bool                                                 `json:"storageNetworkFlag,omitempty"`
	SubnetCount                    *uint                                                 `json:"subnetCount,omitempty"`
	Subnets                        []Network_Subnet                                      `json:"subnets,omitempty"`
	UplinkComponent                *Network_Component                                    `json:"uplinkComponent,omitempty"`
	UplinkDuplexMode               *Network_Component_Duplex_Mode                        `json:"uplinkDuplexMode,omitempty"`
}

type Network_Component_Duplex_Mode struct {
	Entity

	Description *string `json:"description,omitempty"`
	Keyname     *string `json:"keyname,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Network_Component_Firewall struct {
	Entity

	ApplyServerRuleSubnetCount        *uint                             `json:"applyServerRuleSubnetCount,omitempty"`
	ApplyServerRuleSubnets            []Network_Subnet                  `json:"applyServerRuleSubnets,omitempty"`
	BillingItem                       *Billing_Item                     `json:"billingItem,omitempty"`
	GuestNetworkComponent             *Virtual_Guest_Network_Component  `json:"guestNetworkComponent,omitempty"`
	GuestNetworkComponentId           *int                              `json:"guestNetworkComponentId,omitempty"`
	Id                                *int                              `json:"id,omitempty"`
	NetworkComponent                  *Network_Component                `json:"networkComponent,omitempty"`
	NetworkComponentId                *int                              `json:"networkComponentId,omitempty"`
	NetworkFirewallUpdateRequest      []Network_Firewall_Update_Request `json:"networkFirewallUpdateRequest,omitempty"`
	NetworkFirewallUpdateRequestCount *uint                             `json:"networkFirewallUpdateRequestCount,omitempty"`
	RuleCount                         *uint                             `json:"ruleCount,omitempty"`
	Rules                             []Network_Component_Firewall_Rule `json:"rules,omitempty"`
	Status                            *string                           `json:"status,omitempty"`
	SubnetCount                       *uint                             `json:"subnetCount,omitempty"`
	Subnets                           []Network_Subnet                  `json:"subnets,omitempty"`
}

type Network_Component_Firewall_Rule struct {
	Entity

	Action                    *string                     `json:"action,omitempty"`
	DestinationIpAddress      *string                     `json:"destinationIpAddress,omitempty"`
	DestinationIpCidr         *int                        `json:"destinationIpCidr,omitempty"`
	DestinationIpSubnetMask   *string                     `json:"destinationIpSubnetMask,omitempty"`
	DestinationPortRangeEnd   *int                        `json:"destinationPortRangeEnd,omitempty"`
	DestinationPortRangeStart *int                        `json:"destinationPortRangeStart,omitempty"`
	Id                        *int                        `json:"id,omitempty"`
	NetworkComponentFirewall  *Network_Component_Firewall `json:"networkComponentFirewall,omitempty"`
	Notes                     *string                     `json:"notes,omitempty"`
	OrderValue                *int                        `json:"orderValue,omitempty"`
	Protocol                  *string                     `json:"protocol,omitempty"`
	SourceIpAddress           *string                     `json:"sourceIpAddress,omitempty"`
	SourceIpCidr              *int                        `json:"sourceIpCidr,omitempty"`
	SourceIpSubnetMask        *string                     `json:"sourceIpSubnetMask,omitempty"`
	Status                    *string                     `json:"status,omitempty"`
	Version                   *int                        `json:"version,omitempty"`
}

type Network_Component_Firewall_Subnets struct {
	Entity

	ApplyServerRulesFlag     *bool                       `json:"applyServerRulesFlag,omitempty"`
	NetworkComponentFirewall *Network_Component_Firewall `json:"networkComponentFirewall,omitempty"`
	Subnet                   *Network_Subnet             `json:"subnet,omitempty"`
	SubnetId                 *int                        `json:"subnetId,omitempty"`
}

type Network_Component_Group struct {
	Entity

	GroupTypeId           *int                `json:"groupTypeId,omitempty"`
	NetworkComponentCount *uint               `json:"networkComponentCount,omitempty"`
	NetworkComponents     []Network_Component `json:"networkComponents,omitempty"`
}

type Network_Component_IpAddress struct {
	Entity

	IpAddress        *Network_Subnet_IpAddress `json:"ipAddress,omitempty"`
	NetworkComponent *Network_Component        `json:"networkComponent,omitempty"`
}

type Network_Component_Network_Vlan_Trunk struct {
	Entity

	NetworkComponent   *Network_Component `json:"networkComponent,omitempty"`
	NetworkComponentId *int               `json:"networkComponentId,omitempty"`
	NetworkVlan        *Network_Vlan      `json:"networkVlan,omitempty"`
	NetworkVlanId      *int               `json:"networkVlanId,omitempty"`
}

type Network_Component_RemoteManagement struct {
	Network_Component
}

type Network_Component_Uplink_Hardware struct {
	Entity

	Hardware         *Hardware          `json:"hardware,omitempty"`
	NetworkComponent *Network_Component `json:"networkComponent,omitempty"`
}

type Network_ContentDelivery_Account struct {
	Entity

	Account                        *Account                                         `json:"account,omitempty"`
	AccountId                      *int                                             `json:"accountId,omitempty"`
	AssociatedCdnAccountId         *string                                          `json:"associatedCdnAccountId,omitempty"`
	AuthenticationIpAddressCount   *uint                                            `json:"authenticationIpAddressCount,omitempty"`
	AuthenticationIpAddresses      []Network_ContentDelivery_Authentication_Address `json:"authenticationIpAddresses,omitempty"`
	BillingItem                    *Billing_Item                                    `json:"billingItem,omitempty"`
	CdnAccountName                 *string                                          `json:"cdnAccountName,omitempty"`
	CdnAccountNote                 *string                                          `json:"cdnAccountNote,omitempty"`
	CdnSolutionName                *string                                          `json:"cdnSolutionName,omitempty"`
	CreateDate                     *time.Time                                       `json:"createDate,omitempty"`
	DependantServiceFlag           *bool                                            `json:"dependantServiceFlag,omitempty"`
	Id                             *int                                             `json:"id,omitempty"`
	LegacyCdnFlag                  *bool                                            `json:"legacyCdnFlag,omitempty"`
	LogEnabledFlag                 *string                                          `json:"logEnabledFlag,omitempty"`
	ProviderPortalAccessFlag       *bool                                            `json:"providerPortalAccessFlag,omitempty"`
	Status                         *Network_ContentDelivery_Account_Status          `json:"status,omitempty"`
	StatusId                       *int                                             `json:"statusId,omitempty"`
	TokenAuthenticationEnabledFlag *bool                                            `json:"tokenAuthenticationEnabledFlag,omitempty"`
}

type Network_ContentDelivery_Account_Status struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Network_ContentDelivery_Authentication_Address struct {
	Entity

	AccessType   *string    `json:"accessType,omitempty"`
	CdnAccountId *int       `json:"cdnAccountId,omitempty"`
	CreateDate   *time.Time `json:"createDate,omitempty"`
	Id           *int       `json:"id,omitempty"`
	IpAddress    *string    `json:"ipAddress,omitempty"`
	ModifyDate   *time.Time `json:"modifyDate,omitempty"`
	Name         *string    `json:"name,omitempty"`
	Priority     *int       `json:"priority,omitempty"`
	UserId       *int       `json:"userId,omitempty"`
}

type Network_ContentDelivery_Authentication_Token struct {
	Entity

	CdnAccountId *int       `json:"cdnAccountId,omitempty"`
	ClientIp     *string    `json:"clientIp,omitempty"`
	CreateDate   *time.Time `json:"createDate,omitempty"`
	Name         *string    `json:"name,omitempty"`
	Referrer     *string    `json:"referrer,omitempty"`
	Token        *string    `json:"token,omitempty"`
}

type Network_Customer_Subnet struct {
	Entity

	AccountId         *int                                `json:"accountId,omitempty"`
	Cidr              *int                                `json:"cidr,omitempty"`
	Id                *int                                `json:"id,omitempty"`
	IpAddressCount    *uint                               `json:"ipAddressCount,omitempty"`
	IpAddresses       []Network_Customer_Subnet_IpAddress `json:"ipAddresses,omitempty"`
	Netmask           *string                             `json:"netmask,omitempty"`
	NetworkIdentifier *string                             `json:"networkIdentifier,omitempty"`
	TotalIpAddresses  *int                                `json:"totalIpAddresses,omitempty"`
}

type Network_Customer_Subnet_IpAddress struct {
	Entity

	Id               *int                                                `json:"id,omitempty"`
	IpAddress        *string                                             `json:"ipAddress,omitempty"`
	Notes            *string                                             `json:"notes,omitempty"`
	Subnet           *Network_Customer_Subnet                            `json:"subnet,omitempty"`
	SubnetId         *int                                                `json:"subnetId,omitempty"`
	TranslationCount *uint                                               `json:"translationCount,omitempty"`
	Translations     []Network_Tunnel_Module_Context_Address_Translation `json:"translations,omitempty"`
}

type Network_Firewall_AccessControlList struct {
	Entity

	Direction                         *string                           `json:"direction,omitempty"`
	FirewallContextInterfaceId        *int                              `json:"firewallContextInterfaceId,omitempty"`
	Id                                *int                              `json:"id,omitempty"`
	NetworkFirewallUpdateRequestCount *uint                             `json:"networkFirewallUpdateRequestCount,omitempty"`
	NetworkFirewallUpdateRequests     []Network_Firewall_Update_Request `json:"networkFirewallUpdateRequests,omitempty"`
	NetworkVlan                       *Network_Vlan                     `json:"networkVlan,omitempty"`
	RuleCount                         *uint                             `json:"ruleCount,omitempty"`
	Rules                             []Network_Vlan_Firewall_Rule      `json:"rules,omitempty"`
}

type Network_Firewall_Interface struct {
	Network_Firewall_Module_Context_Interface
}

type Network_Firewall_Module_Context_Interface struct {
	Entity

	FirewallContextAccessControlListCount *uint                                `json:"firewallContextAccessControlListCount,omitempty"`
	FirewallContextAccessControlLists     []Network_Firewall_AccessControlList `json:"firewallContextAccessControlLists,omitempty"`
	Id                                    *int                                 `json:"id,omitempty"`
	Name                                  *string                              `json:"name,omitempty"`
	NetworkVlan                           *Network_Vlan                        `json:"networkVlan,omitempty"`
}

type Network_Firewall_Template struct {
	Entity

	Id        *int                             `json:"id,omitempty"`
	Name      *string                          `json:"name,omitempty"`
	RuleCount *uint                            `json:"ruleCount,omitempty"`
	Rules     []Network_Firewall_Template_Rule `json:"rules,omitempty"`
}

type Network_Firewall_Template_Rule struct {
	Entity

	Action                    *string                    `json:"action,omitempty"`
	DestinationIpAddress      *string                    `json:"destinationIpAddress,omitempty"`
	DestinationIpSubnetMask   *string                    `json:"destinationIpSubnetMask,omitempty"`
	DestinationPortRangeEnd   *int                       `json:"destinationPortRangeEnd,omitempty"`
	DestinationPortRangeStart *int                       `json:"destinationPortRangeStart,omitempty"`
	FirewallTemplate          *Network_Firewall_Template `json:"firewallTemplate,omitempty"`
	FirewallTemplateId        *int                       `json:"firewallTemplateId,omitempty"`
	Id                        *int                       `json:"id,omitempty"`
	Notes                     *string                    `json:"notes,omitempty"`
	OrderValue                *int                       `json:"orderValue,omitempty"`
	Protocol                  *string                    `json:"protocol,omitempty"`
	SourceIpAddress           *string                    `json:"sourceIpAddress,omitempty"`
	SourceIpSubnetMask        *string                    `json:"sourceIpSubnetMask,omitempty"`
}

type Network_Firewall_Update_Request struct {
	Entity

	ApplyDate                          *time.Time                             `json:"applyDate,omitempty"`
	AuthorizingUser                    *User_Interface                        `json:"authorizingUser,omitempty"`
	AuthorizingUserId                  *int                                   `json:"authorizingUserId,omitempty"`
	AuthorizingUserType                *string                                `json:"authorizingUserType,omitempty"`
	BypassFlag                         *bool                                  `json:"bypassFlag,omitempty"`
	CreateDate                         *time.Time                             `json:"createDate,omitempty"`
	FirewallContextAccessControlListId *int                                   `json:"firewallContextAccessControlListId,omitempty"`
	Guest                              *Virtual_Guest                         `json:"guest,omitempty"`
	Hardware                           *Hardware                              `json:"hardware,omitempty"`
	HardwareId                         *int                                   `json:"hardwareId,omitempty"`
	Id                                 *int                                   `json:"id,omitempty"`
	NetworkComponentFirewall           *Network_Component_Firewall            `json:"networkComponentFirewall,omitempty"`
	NetworkComponentFirewallId         *int                                   `json:"networkComponentFirewallId,omitempty"`
	RuleCount                          *uint                                  `json:"ruleCount,omitempty"`
	Rules                              []Network_Firewall_Update_Request_Rule `json:"rules,omitempty"`
}

type Network_Firewall_Update_Request_Customer struct {
	Network_Firewall_Update_Request
}

type Network_Firewall_Update_Request_Employee struct {
	Network_Firewall_Update_Request
}

type Network_Firewall_Update_Request_Rule struct {
	Entity

	Action                    *string                          `json:"action,omitempty"`
	DestinationIpAddress      *string                          `json:"destinationIpAddress,omitempty"`
	DestinationIpCidr         *int                             `json:"destinationIpCidr,omitempty"`
	DestinationIpSubnetMask   *string                          `json:"destinationIpSubnetMask,omitempty"`
	DestinationPortRangeEnd   *int                             `json:"destinationPortRangeEnd,omitempty"`
	DestinationPortRangeStart *int                             `json:"destinationPortRangeStart,omitempty"`
	FirewallUpdateRequest     *Network_Firewall_Update_Request `json:"firewallUpdateRequest,omitempty"`
	FirewallUpdateRequestId   *int                             `json:"firewallUpdateRequestId,omitempty"`
	Id                        *int                             `json:"id,omitempty"`
	Notes                     *string                          `json:"notes,omitempty"`
	OrderValue                *int                             `json:"orderValue,omitempty"`
	Protocol                  *string                          `json:"protocol,omitempty"`
	SourceIpAddress           *string                          `json:"sourceIpAddress,omitempty"`
	SourceIpCidr              *int                             `json:"sourceIpCidr,omitempty"`
	SourceIpSubnetMask        *string                          `json:"sourceIpSubnetMask,omitempty"`
	Version                   *int                             `json:"version,omitempty"`
}

type Network_Firewall_Update_Request_Rule_Version6 struct {
	Network_Firewall_Update_Request_Rule
}

type Network_Gateway struct {
	Entity

	Account             *Account                  `json:"account,omitempty"`
	AccountId           *int                      `json:"accountId,omitempty"`
	GroupNumber         *int                      `json:"groupNumber,omitempty"`
	Id                  *int                      `json:"id,omitempty"`
	InsideVlanCount     *uint                     `json:"insideVlanCount,omitempty"`
	InsideVlans         []Network_Gateway_Vlan    `json:"insideVlans,omitempty"`
	MemberCount         *uint                     `json:"memberCount,omitempty"`
	Members             []Network_Gateway_Member  `json:"members,omitempty"`
	Name                *string                   `json:"name,omitempty"`
	NetworkSpace        *string                   `json:"networkSpace,omitempty"`
	PrivateIpAddress    *Network_Subnet_IpAddress `json:"privateIpAddress,omitempty"`
	PrivateIpAddressId  *int                      `json:"privateIpAddressId,omitempty"`
	PrivateVlan         *Network_Vlan             `json:"privateVlan,omitempty"`
	PrivateVlanId       *int                      `json:"privateVlanId,omitempty"`
	PublicIpAddress     *Network_Subnet_IpAddress `json:"publicIpAddress,omitempty"`
	PublicIpAddressId   *int                      `json:"publicIpAddressId,omitempty"`
	PublicIpv6Address   *Network_Subnet_IpAddress `json:"publicIpv6Address,omitempty"`
	PublicIpv6AddressId *int                      `json:"publicIpv6AddressId,omitempty"`
	PublicVlan          *Network_Vlan             `json:"publicVlan,omitempty"`
	PublicVlanId        *int                      `json:"publicVlanId,omitempty"`
	Status              *Network_Gateway_Status   `json:"status,omitempty"`
	StatusId            *int                      `json:"statusId,omitempty"`
}

type Network_Gateway_Member struct {
	Entity

	Hardware         *Hardware        `json:"hardware,omitempty"`
	HardwareId       *int             `json:"hardwareId,omitempty"`
	Id               *int             `json:"id,omitempty"`
	NetworkGateway   *Network_Gateway `json:"networkGateway,omitempty"`
	NetworkGatewayId *int             `json:"networkGatewayId,omitempty"`
	Priority         *int             `json:"priority,omitempty"`
}

type Network_Gateway_Status struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	KeyName     *string `json:"keyName,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Network_Gateway_Vlan struct {
	Entity

	BypassFlag       *bool            `json:"bypassFlag,omitempty"`
	Id               *int             `json:"id,omitempty"`
	NetworkGateway   *Network_Gateway `json:"networkGateway,omitempty"`
	NetworkGatewayId *int             `json:"networkGatewayId,omitempty"`
	NetworkVlan      *Network_Vlan    `json:"networkVlan,omitempty"`
	NetworkVlanId    *int             `json:"networkVlanId,omitempty"`
}

type Network_LoadBalancer_Global_Account struct {
	Entity

	Account                     *Account                           `json:"account,omitempty"`
	AllowedNumberOfHosts        *int                               `json:"allowedNumberOfHosts,omitempty"`
	AverageConnectionsPerSecond *float64                           `json:"averageConnectionsPerSecond,omitempty"`
	BillingItem                 *Billing_Item                      `json:"billingItem,omitempty"`
	ConnectionsPerSecond        *int                               `json:"connectionsPerSecond,omitempty"`
	FallbackIp                  *string                            `json:"fallbackIp,omitempty"`
	HostCount                   *uint                              `json:"hostCount,omitempty"`
	Hostname                    *string                            `json:"hostname,omitempty"`
	Hosts                       []Network_LoadBalancer_Global_Host `json:"hosts,omitempty"`
	Id                          *int                               `json:"id,omitempty"`
	LoadBalanceType             *Network_LoadBalancer_Global_Type  `json:"loadBalanceType,omitempty"`
	LoadBalanceTypeId           *int                               `json:"loadBalanceTypeId,omitempty"`
	ManagedResourceFlag         *bool                              `json:"managedResourceFlag,omitempty"`
	Notes                       *string                            `json:"notes,omitempty"`
}

type Network_LoadBalancer_Global_Host struct {
	Entity

	DestinationIp       *string                              `json:"destinationIp,omitempty"`
	DestinationPort     *int                                 `json:"destinationPort,omitempty"`
	Enabled             *int                                 `json:"enabled,omitempty"`
	HealthCheck         *string                              `json:"healthCheck,omitempty"`
	Hits                *float64                             `json:"hits,omitempty"`
	Id                  *int                                 `json:"id,omitempty"`
	LoadBalanceOrder    *int                                 `json:"loadBalanceOrder,omitempty"`
	LoadBalancerAccount *Network_LoadBalancer_Global_Account `json:"loadBalancerAccount,omitempty"`
	Location            *string                              `json:"location,omitempty"`
	Status              *string                              `json:"status,omitempty"`
	Weight              *int                                 `json:"weight,omitempty"`
}

type Network_LoadBalancer_Global_Type struct {
	Entity

	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Network_LoadBalancer_Service struct {
	Entity

	ConnectionLimit      *int                                   `json:"connectionLimit,omitempty"`
	CreateDate           *time.Time                             `json:"createDate,omitempty"`
	DestinationIpAddress *string                                `json:"destinationIpAddress,omitempty"`
	DestinationPort      *int                                   `json:"destinationPort,omitempty"`
	Enabled              *bool                                  `json:"enabled,omitempty"`
	HealthCheck          *string                                `json:"healthCheck,omitempty"`
	HealthCheckURL       *string                                `json:"healthCheckURL,omitempty"`
	HealthResponse       *string                                `json:"healthResponse,omitempty"`
	Id                   *int                                   `json:"id,omitempty"`
	ModifyDate           *time.Time                             `json:"modifyDate,omitempty"`
	Name                 *string                                `json:"name,omitempty"`
	Notes                *string                                `json:"notes,omitempty"`
	PeakConnections      *int                                   `json:"peakConnections,omitempty"`
	SourcePort           *int                                   `json:"sourcePort,omitempty"`
	Type                 *string                                `json:"type,omitempty"`
	Vip                  *Network_LoadBalancer_VirtualIpAddress `json:"vip,omitempty"`
	VipId                *int                                   `json:"vipId,omitempty"`
	Weight               *int                                   `json:"weight,omitempty"`
}

type Network_LoadBalancer_VirtualIpAddress struct {
	Entity

	Account                     *Account                       `json:"account,omitempty"`
	BillingItem                 *Billing_Item                  `json:"billingItem,omitempty"`
	ConnectionLimit             *int                           `json:"connectionLimit,omitempty"`
	CustomerManagedFlag         *int                           `json:"customerManagedFlag,omitempty"`
	Id                          *int                           `json:"id,omitempty"`
	LoadBalancingMethod         *string                        `json:"loadBalancingMethod,omitempty"`
	LoadBalancingMethodFullName *string                        `json:"loadBalancingMethodFullName,omitempty"`
	ManagedResourceFlag         *bool                          `json:"managedResourceFlag,omitempty"`
	ModifyDate                  *time.Time                     `json:"modifyDate,omitempty"`
	Name                        *string                        `json:"name,omitempty"`
	Notes                       *string                        `json:"notes,omitempty"`
	SecurityCertificateId       *int                           `json:"securityCertificateId,omitempty"`
	ServiceCount                *uint                          `json:"serviceCount,omitempty"`
	Services                    []Network_LoadBalancer_Service `json:"services,omitempty"`
	SourcePort                  *int                           `json:"sourcePort,omitempty"`
	Type                        *string                        `json:"type,omitempty"`
	VirtualIpAddress            *string                        `json:"virtualIpAddress,omitempty"`
}

type Network_Logging_Syslog struct {
	Entity

	CreateDate           *time.Time `json:"createDate,omitempty"`
	DestinationIpAddress *string    `json:"destinationIpAddress,omitempty"`
	DestinationPort      *int       `json:"destinationPort,omitempty"`
	EventType            *string    `json:"eventType,omitempty"`
	Message              *string    `json:"message,omitempty"`
	Protocol             *string    `json:"protocol,omitempty"`
	SourceIpAddress      *string    `json:"sourceIpAddress,omitempty"`
	SourcePort           *int       `json:"sourcePort,omitempty"`
	TotalEvents          *int       `json:"totalEvents,omitempty"`
}

type Network_Media_Transcode_Account struct {
	Entity

	Account           *Account                      `json:"account,omitempty"`
	AccountId         *int                          `json:"accountId,omitempty"`
	CreateDate        *time.Time                    `json:"createDate,omitempty"`
	Id                *int                          `json:"id,omitempty"`
	ModifyDate        *time.Time                    `json:"modifyDate,omitempty"`
	TranscodeJobCount *uint                         `json:"transcodeJobCount,omitempty"`
	TranscodeJobs     []Network_Media_Transcode_Job `json:"transcodeJobs,omitempty"`
}

type Network_Media_Transcode_Job struct {
	Entity

	AutoDeleteDuration  *int                                             `json:"autoDeleteDuration,omitempty"`
	ByteIn              *int                                             `json:"byteIn,omitempty"`
	CreateDate          *time.Time                                       `json:"createDate,omitempty"`
	History             []Network_Media_Transcode_Job_History            `json:"history,omitempty"`
	HistoryCount        *uint                                            `json:"historyCount,omitempty"`
	Id                  *int                                             `json:"id,omitempty"`
	InputFile           *string                                          `json:"inputFile,omitempty"`
	ModifyDate          *time.Time                                       `json:"modifyDate,omitempty"`
	Name                *string                                          `json:"name,omitempty"`
	OutputFile          *string                                          `json:"outputFile,omitempty"`
	TranscodeAccount    *Network_Media_Transcode_Account                 `json:"transcodeAccount,omitempty"`
	TranscodeAccountId  *int                                             `json:"transcodeAccountId,omitempty"`
	TranscodeJobGuid    *string                                          `json:"transcodeJobGuid,omitempty"`
	TranscodePresetGuid *string                                          `json:"transcodePresetGuid,omitempty"`
	TranscodePresetName *string                                          `json:"transcodePresetName,omitempty"`
	TranscodeStatus     *Network_Media_Transcode_Job_Status              `json:"transcodeStatus,omitempty"`
	TranscodeStatusId   *int                                             `json:"transcodeStatusId,omitempty"`
	TranscodeStatusName *string                                          `json:"transcodeStatusName,omitempty"`
	User                *User_Customer                                   `json:"user,omitempty"`
	UserId              *int                                             `json:"userId,omitempty"`
	Watermark           *Container_Network_Media_Transcode_Job_Watermark `json:"watermark,omitempty"`
}

type Network_Media_Transcode_Job_History struct {
	Entity

	CreateDate          *time.Time `json:"createDate,omitempty"`
	PublicNotes         *string    `json:"publicNotes,omitempty"`
	TranscodeJobId      *int       `json:"transcodeJobId,omitempty"`
	TranscodeStatusName *string    `json:"transcodeStatusName,omitempty"`
}

type Network_Media_Transcode_Job_Status struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Network_Message_Delivery struct {
	Entity

	Account     *Account                         `json:"account,omitempty"`
	AccountId   *int                             `json:"accountId,omitempty"`
	BillingItem *Billing_Item                    `json:"billingItem,omitempty"`
	CreateDate  *time.Time                       `json:"createDate,omitempty"`
	Id          *int                             `json:"id,omitempty"`
	ModifyDate  *time.Time                       `json:"modifyDate,omitempty"`
	Password    *string                          `json:"password,omitempty"`
	Type        *Network_Message_Delivery_Type   `json:"type,omitempty"`
	TypeId      *int                             `json:"typeId,omitempty"`
	Username    *string                          `json:"username,omitempty"`
	Vendor      *Network_Message_Delivery_Vendor `json:"vendor,omitempty"`
	VendorId    *int                             `json:"vendorId,omitempty"`
}

type Network_Message_Delivery_Attribute struct {
	Entity

	NetworkMessageDelivery *Network_Message_Delivery `json:"networkMessageDelivery,omitempty"`
	Value                  *string                   `json:"value,omitempty"`
}

type Network_Message_Delivery_Email_Sendgrid struct {
	Network_Message_Delivery

	EmailAddress *string `json:"emailAddress,omitempty"`
	SmtpAccess   *string `json:"smtpAccess,omitempty"`
}

type Network_Message_Delivery_Type struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	KeyName     *string `json:"keyName,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Network_Message_Delivery_Vendor struct {
	Entity

	Id      *int    `json:"id,omitempty"`
	KeyName *string `json:"keyName,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type Network_Message_Queue struct {
	Entity

	Account              *Account                      `json:"account,omitempty"`
	AccountId            *int                          `json:"accountId,omitempty"`
	BillingItem          *Billing_Item                 `json:"billingItem,omitempty"`
	CreateDate           *time.Time                    `json:"createDate,omitempty"`
	Id                   *int                          `json:"id,omitempty"`
	MessageQueueStatusId *int                          `json:"messageQueueStatusId,omitempty"`
	Name                 *string                       `json:"name,omitempty"`
	NodeCount            *uint                         `json:"nodeCount,omitempty"`
	Nodes                []Network_Message_Queue_Node  `json:"nodes,omitempty"`
	Notes                *string                       `json:"notes,omitempty"`
	Status               *Network_Message_Queue_Status `json:"status,omitempty"`
}

type Network_Message_Queue_Node struct {
	Entity

	AccountName          *string                   `json:"accountName,omitempty"`
	Id                   *int                      `json:"id,omitempty"`
	MessageQueue         *Network_Message_Queue    `json:"messageQueue,omitempty"`
	MessageQueueId       *int                      `json:"messageQueueId,omitempty"`
	MetricTrackingObject *Metric_Tracking_Object   `json:"metricTrackingObject,omitempty"`
	Name                 *string                   `json:"name,omitempty"`
	Notes                *string                   `json:"notes,omitempty"`
	ServiceResource      *Network_Service_Resource `json:"serviceResource,omitempty"`
}

type Network_Message_Queue_Status struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Network_Monitor struct {
	Entity
}

type Network_Monitor_Version1_Incident struct {
	Entity

	Status *string `json:"status,omitempty"`
}

type Network_Monitor_Version1_Query_Host struct {
	Entity

	Arg1Value        *string                                      `json:"arg1Value,omitempty"`
	GuestId          *int                                         `json:"guestId,omitempty"`
	Hardware         *Hardware                                    `json:"hardware,omitempty"`
	HardwareId       *int                                         `json:"hardwareId,omitempty"`
	HostId           *int                                         `json:"hostId,omitempty"`
	Id               *int                                         `json:"id,omitempty"`
	IpAddress        *string                                      `json:"ipAddress,omitempty"`
	LastResult       *Network_Monitor_Version1_Query_Result       `json:"lastResult,omitempty"`
	QueryType        *Network_Monitor_Version1_Query_Type         `json:"queryType,omitempty"`
	QueryTypeId      *int                                         `json:"queryTypeId,omitempty"`
	ResponseAction   *Network_Monitor_Version1_Query_ResponseType `json:"responseAction,omitempty"`
	ResponseActionId *int                                         `json:"responseActionId,omitempty"`
	Status           *string                                      `json:"status,omitempty"`
	WaitCycles       *int                                         `json:"waitCycles,omitempty"`
}

type Network_Monitor_Version1_Query_Host_Stratum struct {
	Entity

	Hardware      *Hardware `json:"hardware,omitempty"`
	MonitorLevel  *int      `json:"monitorLevel,omitempty"`
	ResponseLevel *int      `json:"responseLevel,omitempty"`
}

type Network_Monitor_Version1_Query_ResponseType struct {
	Entity

	ActionDescription *string `json:"actionDescription,omitempty"`
	Id                *int    `json:"id,omitempty"`
	Level             *int    `json:"level,omitempty"`
}

type Network_Monitor_Version1_Query_Result struct {
	Entity

	FinishTime     *time.Time                           `json:"finishTime,omitempty"`
	QueryHost      *Network_Monitor_Version1_Query_Host `json:"queryHost,omitempty"`
	ResponseStatus *int                                 `json:"responseStatus,omitempty"`
	ResponseTime   *float64                             `json:"responseTime,omitempty"`
}

type Network_Monitor_Version1_Query_Type struct {
	Entity

	ArgumentDescription *string `json:"argumentDescription,omitempty"`
	Description         *string `json:"description,omitempty"`
	Id                  *int    `json:"id,omitempty"`
	MonitorLevel        *int    `json:"monitorLevel,omitempty"`
	Name                *string `json:"name,omitempty"`
}

type Network_Pod struct {
	Entity

	BackendRouterId    *int     `json:"backendRouterId,omitempty"`
	BackendRouterName  *string  `json:"backendRouterName,omitempty"`
	Capabilities       []string `json:"capabilities,omitempty"`
	DatacenterLongName *string  `json:"datacenterLongName,omitempty"`
	DatacenterName     *string  `json:"datacenterName,omitempty"`
	FrontendRouterId   *int     `json:"frontendRouterId,omitempty"`
	FrontendRouterName *string  `json:"frontendRouterName,omitempty"`
	Name               *string  `json:"name,omitempty"`
}

type Network_Protection_Address struct {
	Entity

	Account              *Account                            `json:"account,omitempty"`
	DepartmentId         *int                                `json:"departmentId,omitempty"`
	IpAddress            *string                             `json:"ipAddress,omitempty"`
	Location             *Location                           `json:"location,omitempty"`
	ManagementMethodType *string                             `json:"managementMethodType,omitempty"`
	ModifiedUser         *User_Employee                      `json:"modifiedUser,omitempty"`
	PrimaryRouter        *Hardware_Router                    `json:"primaryRouter,omitempty"`
	ServiceProvider      *Service_Provider                   `json:"serviceProvider,omitempty"`
	Subnet               *Network_Subnet                     `json:"subnet,omitempty"`
	SubnetIpAddress      *Network_Subnet_IpAddress           `json:"subnetIpAddress,omitempty"`
	TerminatedUser       *User_Employee                      `json:"terminatedUser,omitempty"`
	Ticket               *Ticket                             `json:"ticket,omitempty"`
	TransactionCount     *uint                               `json:"transactionCount,omitempty"`
	Transactions         []Provisioning_Version1_Transaction `json:"transactions,omitempty"`
	UserDepartment       *User_Employee_Department           `json:"userDepartment,omitempty"`
	UserRecord           *User_Employee                      `json:"userRecord,omitempty"`
}

type Network_Regional_Internet_Registry struct {
	Entity

	Id      *int    `json:"id,omitempty"`
	KeyName *string `json:"keyName,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type Network_Security_Scanner_Request struct {
	Entity

	Account            *Account                                 `json:"account,omitempty"`
	AccountId          *int                                     `json:"accountId,omitempty"`
	CreateDate         *time.Time                               `json:"createDate,omitempty"`
	Guest              *Virtual_Guest                           `json:"guest,omitempty"`
	GuestId            *int                                     `json:"guestId,omitempty"`
	Hardware           *Hardware                                `json:"hardware,omitempty"`
	HardwareId         *int                                     `json:"hardwareId,omitempty"`
	HostId             *int                                     `json:"hostId,omitempty"`
	Id                 *int                                     `json:"id,omitempty"`
	IpAddress          *string                                  `json:"ipAddress,omitempty"`
	ModifyDate         *time.Time                               `json:"modifyDate,omitempty"`
	RequestorOwnedFlag *bool                                    `json:"requestorOwnedFlag,omitempty"`
	Status             *Network_Security_Scanner_Request_Status `json:"status,omitempty"`
	StatusId           *int                                     `json:"statusId,omitempty"`
}

type Network_Security_Scanner_Request_Status struct {
	Entity

	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Network_Service_Health struct {
	Entity

	CreateDate *time.Time                     `json:"createDate,omitempty"`
	Location   *Location                      `json:"location,omitempty"`
	LocationId *int                           `json:"locationId,omitempty"`
	ModifyDate *time.Time                     `json:"modifyDate,omitempty"`
	Status     *Network_Service_Health_Status `json:"status,omitempty"`
	StatusId   *int                           `json:"statusId,omitempty"`
}

type Network_Service_Health_Status struct {
	Entity

	Name *string `json:"name,omitempty"`
}

type Network_Service_Resource struct {
	Entity

	ApiHost           *string                              `json:"apiHost,omitempty"`
	ApiPassword       *string                              `json:"apiPassword,omitempty"`
	ApiPath           *string                              `json:"apiPath,omitempty"`
	ApiPort           *string                              `json:"apiPort,omitempty"`
	ApiProtocol       *string                              `json:"apiProtocol,omitempty"`
	ApiUsername       *string                              `json:"apiUsername,omitempty"`
	ApiVersion        *string                              `json:"apiVersion,omitempty"`
	AttributeCount    *uint                                `json:"attributeCount,omitempty"`
	Attributes        []Network_Service_Resource_Attribute `json:"attributes,omitempty"`
	BackendIpAddress  *string                              `json:"backendIpAddress,omitempty"`
	Datacenter        *Location                            `json:"datacenter,omitempty"`
	FrontendIpAddress *string                              `json:"frontendIpAddress,omitempty"`
	Id                *int                                 `json:"id,omitempty"`
	Name              *string                              `json:"name,omitempty"`
	NetworkDevice     *Hardware                            `json:"networkDevice,omitempty"`
	SshUsername       *string                              `json:"sshUsername,omitempty"`
	Type              *Network_Service_Resource_Type       `json:"type,omitempty"`
}

type Network_Service_Resource_Attribute struct {
	Entity

	AttributeType   *Network_Service_Resource_Attribute_Type `json:"attributeType,omitempty"`
	ServiceResource *Network_Service_Resource                `json:"serviceResource,omitempty"`
	Value           *string                                  `json:"value,omitempty"`
}

type Network_Service_Resource_Attribute_Type struct {
	Entity

	Keyname *string `json:"keyname,omitempty"`
}

type Network_Service_Resource_Hub struct {
	Network_Service_Resource
}

type Network_Service_Resource_Hub_Swift struct {
	Network_Service_Resource_Hub
}

type Network_Service_Resource_MonitoringHub struct {
	Network_Service_Resource

	AdnServicesIp        *string `json:"adnServicesIp,omitempty"`
	HubAddress           *string `json:"hubAddress,omitempty"`
	HubConnectionTimeout *string `json:"hubConnectionTimeout,omitempty"`
	RobotsCount          *string `json:"robotsCount,omitempty"`
	RobotsMax            *string `json:"robotsMax,omitempty"`
}

type Network_Service_Resource_NimsoftLandingHub struct {
	Network_Service_Resource_MonitoringHub
}

type Network_Service_Resource_Type struct {
	Entity

	ServiceResourceCount *uint                      `json:"serviceResourceCount,omitempty"`
	ServiceResources     []Network_Service_Resource `json:"serviceResources,omitempty"`
	Type                 *string                    `json:"type,omitempty"`
}

type Network_Service_Vpn_Overrides struct {
	Entity

	Id       *int            `json:"id,omitempty"`
	Subnet   *Network_Subnet `json:"subnet,omitempty"`
	SubnetId *int            `json:"subnetId,omitempty"`
	User     *User_Customer  `json:"user,omitempty"`
	UserId   *int            `json:"userId,omitempty"`
}

type Network_Storage struct {
	Entity

	Account                             *Account                            `json:"account,omitempty"`
	AccountId                           *int                                `json:"accountId,omitempty"`
	AccountPassword                     *Account_Password                   `json:"accountPassword,omitempty"`
	ActiveTransactionCount              *uint                               `json:"activeTransactionCount,omitempty"`
	ActiveTransactions                  []Provisioning_Version1_Transaction `json:"activeTransactions,omitempty"`
	AllowedHardware                     []Hardware                          `json:"allowedHardware,omitempty"`
	AllowedHardwareCount                *uint                               `json:"allowedHardwareCount,omitempty"`
	AllowedIpAddressCount               *uint                               `json:"allowedIpAddressCount,omitempty"`
	AllowedIpAddresses                  []Network_Subnet_IpAddress          `json:"allowedIpAddresses,omitempty"`
	AllowedReplicationHardware          []Hardware                          `json:"allowedReplicationHardware,omitempty"`
	AllowedReplicationHardwareCount     *uint                               `json:"allowedReplicationHardwareCount,omitempty"`
	AllowedReplicationIpAddressCount    *uint                               `json:"allowedReplicationIpAddressCount,omitempty"`
	AllowedReplicationIpAddresses       []Network_Subnet_IpAddress          `json:"allowedReplicationIpAddresses,omitempty"`
	AllowedReplicationSubnetCount       *uint                               `json:"allowedReplicationSubnetCount,omitempty"`
	AllowedReplicationSubnets           []Network_Subnet                    `json:"allowedReplicationSubnets,omitempty"`
	AllowedReplicationVirtualGuestCount *uint                               `json:"allowedReplicationVirtualGuestCount,omitempty"`
	AllowedReplicationVirtualGuests     []Virtual_Guest                     `json:"allowedReplicationVirtualGuests,omitempty"`
	AllowedSubnetCount                  *uint                               `json:"allowedSubnetCount,omitempty"`
	AllowedSubnets                      []Network_Subnet                    `json:"allowedSubnets,omitempty"`
	AllowedVirtualGuestCount            *uint                               `json:"allowedVirtualGuestCount,omitempty"`
	AllowedVirtualGuests                []Virtual_Guest                     `json:"allowedVirtualGuests,omitempty"`
	BillingItem                         *Billing_Item                       `json:"billingItem,omitempty"`
	BillingItemCategory                 *Product_Item_Category              `json:"billingItemCategory,omitempty"`
	BytesUsed                           *string                             `json:"bytesUsed,omitempty"`
	CapacityGb                          *int                                `json:"capacityGb,omitempty"`
	CreateDate                          *time.Time                          `json:"createDate,omitempty"`
	CreationScheduleId                  *string                             `json:"creationScheduleId,omitempty"`
	CredentialCount                     *uint                               `json:"credentialCount,omitempty"`
	Credentials                         []Network_Storage_Credential        `json:"credentials,omitempty"`
	DailySchedule                       *Network_Storage_Schedule           `json:"dailySchedule,omitempty"`
	EventCount                          *uint                               `json:"eventCount,omitempty"`
	Events                              []Network_Storage_Event             `json:"events,omitempty"`
	GuestId                             *int                                `json:"guestId,omitempty"`
	Hardware                            *Hardware                           `json:"hardware,omitempty"`
	HardwareId                          *int                                `json:"hardwareId,omitempty"`
	HostId                              *int                                `json:"hostId,omitempty"`
	HourlySchedule                      *Network_Storage_Schedule           `json:"hourlySchedule,omitempty"`
	Id                                  *int                                `json:"id,omitempty"`
	Iops                                *string                             `json:"iops,omitempty"`
	IscsiLunCount                       *uint                               `json:"iscsiLunCount,omitempty"`
	IscsiLuns                           []Network_Storage                   `json:"iscsiLuns,omitempty"`
	LunId                               *string                             `json:"lunId,omitempty"`
	ManualSnapshotCount                 *uint                               `json:"manualSnapshotCount,omitempty"`
	ManualSnapshots                     []Network_Storage                   `json:"manualSnapshots,omitempty"`
	MetricTrackingObject                *Metric_Tracking_Object             `json:"metricTrackingObject,omitempty"`
	MountableFlag                       *string                             `json:"mountableFlag,omitempty"`
	NasType                             *string                             `json:"nasType,omitempty"`
	Notes                               *string                             `json:"notes,omitempty"`
	NotificationSubscriberCount         *uint                               `json:"notificationSubscriberCount,omitempty"`
	NotificationSubscribers             []Notification_User_Subscriber      `json:"notificationSubscribers,omitempty"`
	OsType                              *Network_Storage_Iscsi_OS_Type      `json:"osType,omitempty"`
	OsTypeId                            *string                             `json:"osTypeId,omitempty"`
	ParentPartnershipCount              *uint                               `json:"parentPartnershipCount,omitempty"`
	ParentPartnerships                  []Network_Storage_Partnership       `json:"parentPartnerships,omitempty"`
	ParentVolume                        *Network_Storage                    `json:"parentVolume,omitempty"`
	PartnershipCount                    *uint                               `json:"partnershipCount,omitempty"`
	Partnerships                        []Network_Storage_Partnership       `json:"partnerships,omitempty"`
	Password                            *string                             `json:"password,omitempty"`
	PermissionsGroupCount               *uint                               `json:"permissionsGroupCount,omitempty"`
	PermissionsGroups                   []Network_Storage_Group             `json:"permissionsGroups,omitempty"`
	Properties                          []Network_Storage_Property          `json:"properties,omitempty"`
	PropertyCount                       *uint                               `json:"propertyCount,omitempty"`
	ReplicatingLunCount                 *uint                               `json:"replicatingLunCount,omitempty"`
	ReplicatingLuns                     []Network_Storage                   `json:"replicatingLuns,omitempty"`
	ReplicatingVolume                   *Network_Storage                    `json:"replicatingVolume,omitempty"`
	ReplicationEventCount               *uint                               `json:"replicationEventCount,omitempty"`
	ReplicationEvents                   []Network_Storage_Event             `json:"replicationEvents,omitempty"`
	ReplicationPartnerCount             *uint                               `json:"replicationPartnerCount,omitempty"`
	ReplicationPartners                 []Network_Storage                   `json:"replicationPartners,omitempty"`
	ReplicationSchedule                 *Network_Storage_Schedule           `json:"replicationSchedule,omitempty"`
	ReplicationStatus                   *string                             `json:"replicationStatus,omitempty"`
	ScheduleCount                       *uint                               `json:"scheduleCount,omitempty"`
	Schedules                           []Network_Storage_Schedule          `json:"schedules,omitempty"`
	ServiceProviderId                   *int                                `json:"serviceProviderId,omitempty"`
	ServiceResource                     *Network_Service_Resource           `json:"serviceResource,omitempty"`
	ServiceResourceBackendIpAddress     *string                             `json:"serviceResourceBackendIpAddress,omitempty"`
	ServiceResourceName                 *string                             `json:"serviceResourceName,omitempty"`
	SnapshotCapacityGb                  *string                             `json:"snapshotCapacityGb,omitempty"`
	SnapshotCount                       *uint                               `json:"snapshotCount,omitempty"`
	SnapshotCreationTimestamp           *string                             `json:"snapshotCreationTimestamp,omitempty"`
	SnapshotDeletionThresholdPercentage *string                             `json:"snapshotDeletionThresholdPercentage,omitempty"`
	SnapshotSizeBytes                   *string                             `json:"snapshotSizeBytes,omitempty"`
	SnapshotSpaceAvailable              *string                             `json:"snapshotSpaceAvailable,omitempty"`
	Snapshots                           []Network_Storage                   `json:"snapshots,omitempty"`
	StorageGroupCount                   *uint                               `json:"storageGroupCount,omitempty"`
	StorageGroups                       []Network_Storage_Group             `json:"storageGroups,omitempty"`
	StorageTierLevel                    *string                             `json:"storageTierLevel,omitempty"`
	StorageType                         *Network_Storage_Type               `json:"storageType,omitempty"`
	StorageTypeId                       *string                             `json:"storageTypeId,omitempty"`
	TotalBytesUsed                      *string                             `json:"totalBytesUsed,omitempty"`
	TotalScheduleSnapshotRetentionCount *uint                               `json:"totalScheduleSnapshotRetentionCount,omitempty"`
	UpgradableFlag                      *bool                               `json:"upgradableFlag,omitempty"`
	UsageNotification                   *Notification                       `json:"usageNotification,omitempty"`
	Username                            *string                             `json:"username,omitempty"`
	VendorName                          *string                             `json:"vendorName,omitempty"`
	VirtualGuest                        *Virtual_Guest                      `json:"virtualGuest,omitempty"`
	VolumeHistory                       []Network_Storage_History           `json:"volumeHistory,omitempty"`
	VolumeHistoryCount                  *uint                               `json:"volumeHistoryCount,omitempty"`
	VolumeStatus                        *string                             `json:"volumeStatus,omitempty"`
	WebccAccount                        *Account_Password                   `json:"webccAccount,omitempty"`
	WeeklySchedule                      *Network_Storage_Schedule           `json:"weeklySchedule,omitempty"`
}

type Network_Storage_Allowed_Host struct {
	Entity

	AssignedGroupCount             *uint                       `json:"assignedGroupCount,omitempty"`
	AssignedGroups                 []Network_Storage_Group     `json:"assignedGroups,omitempty"`
	AssignedReplicationVolumeCount *uint                       `json:"assignedReplicationVolumeCount,omitempty"`
	AssignedReplicationVolumes     []Network_Storage           `json:"assignedReplicationVolumes,omitempty"`
	AssignedVolumeCount            *uint                       `json:"assignedVolumeCount,omitempty"`
	AssignedVolumes                []Network_Storage           `json:"assignedVolumes,omitempty"`
	Credential                     *Network_Storage_Credential `json:"credential,omitempty"`
	CredentialId                   *int                        `json:"credentialId,omitempty"`
	Id                             *int                        `json:"id,omitempty"`
	Name                           *string                     `json:"name,omitempty"`
	ResourceTableId                *int                        `json:"resourceTableId,omitempty"`
	ResourceTableName              *string                     `json:"resourceTableName,omitempty"`
}

type Network_Storage_Allowed_Host_Hardware struct {
	Network_Storage_Allowed_Host

	Resource *Hardware `json:"resource,omitempty"`
}

type Network_Storage_Allowed_Host_IpAddress struct {
	Network_Storage_Allowed_Host

	Resource *Network_Subnet_IpAddress `json:"resource,omitempty"`
}

type Network_Storage_Allowed_Host_Subnet struct {
	Network_Storage_Allowed_Host

	Resource *Network_Subnet `json:"resource,omitempty"`
}

type Network_Storage_Allowed_Host_VirtualGuest struct {
	Network_Storage_Allowed_Host

	Resource *Virtual_Guest `json:"resource,omitempty"`
}

type Network_Storage_Backup struct {
	Network_Storage

	CurrentCyclePeakUsage  *uint `json:"currentCyclePeakUsage,omitempty"`
	PreviousCyclePeakUsage *uint `json:"previousCyclePeakUsage,omitempty"`
}

type Network_Storage_Backup_Evault struct {
	Network_Storage_Backup
}

type Network_Storage_Backup_Evault_Version6 struct {
	Network_Storage_Backup_Evault

	AgentStatusCount       *uint                                                `json:"agentStatusCount,omitempty"`
	AgentStatuses          []Container_Network_Storage_Evault_WebCc_AgentStatus `json:"agentStatuses,omitempty"`
	BackupJobDetailCount   *uint                                                `json:"backupJobDetailCount,omitempty"`
	BackupJobDetails       []Container_Network_Storage_Evault_WebCc_JobDetails  `json:"backupJobDetails,omitempty"`
	PluginBillingItemCount *uint                                                `json:"pluginBillingItemCount,omitempty"`
	PluginBillingItems     []Billing_Item                                       `json:"pluginBillingItems,omitempty"`
	RestoreJobDetailCount  *uint                                                `json:"restoreJobDetailCount,omitempty"`
	RestoreJobDetails      []Container_Network_Storage_Evault_WebCc_JobDetails  `json:"restoreJobDetails,omitempty"`
	SoftwareComponent      *Software_Component                                  `json:"softwareComponent,omitempty"`
	TaskCount              *uint                                                `json:"taskCount,omitempty"`
	Tasks                  []Container_Network_Storage_Evault_Vault_Task        `json:"tasks,omitempty"`
}

type Network_Storage_Credential struct {
	Entity

	Account                    *Account                         `json:"account,omitempty"`
	AccountId                  *string                          `json:"accountId,omitempty"`
	CreateDate                 *time.Time                       `json:"createDate,omitempty"`
	Id                         *int                             `json:"id,omitempty"`
	ModifyDate                 *time.Time                       `json:"modifyDate,omitempty"`
	NasCredentialTypeId        *int                             `json:"nasCredentialTypeId,omitempty"`
	NetworkStorageAllowedHosts *Network_Storage_Allowed_Host    `json:"networkStorageAllowedHosts,omitempty"`
	Password                   *string                          `json:"password,omitempty"`
	Type                       *Network_Storage_Credential_Type `json:"type,omitempty"`
	Username                   *string                          `json:"username,omitempty"`
	VolumeCount                *uint                            `json:"volumeCount,omitempty"`
	Volumes                    []Network_Storage                `json:"volumes,omitempty"`
}

type Network_Storage_Credential_Type struct {
	Entity

	CreateDate  *time.Time `json:"createDate,omitempty"`
	Description *string    `json:"description,omitempty"`
	KeyName     *string    `json:"keyName,omitempty"`
	ModifyDate  *time.Time `json:"modifyDate,omitempty"`
	Name        *string    `json:"name,omitempty"`
}

type Network_Storage_Daily_Usage struct {
	Entity

	BytesUsed          *uint            `json:"bytesUsed,omitempty"`
	CdnHttpBandwidth   *uint            `json:"cdnHttpBandwidth,omitempty"`
	CreateDate         *time.Time       `json:"createDate,omitempty"`
	NasVolume          *Network_Storage `json:"nasVolume,omitempty"`
	NasVolumeId        *int             `json:"nasVolumeId,omitempty"`
	PublicBandwidthOut *uint            `json:"publicBandwidthOut,omitempty"`
}

type Network_Storage_Event struct {
	Entity

	CreateDate *time.Time                `json:"createDate,omitempty"`
	Message    *string                   `json:"message,omitempty"`
	Schedule   *Network_Storage_Schedule `json:"schedule,omitempty"`
	ScheduleId *int                      `json:"scheduleId,omitempty"`
	TypeId     *int                      `json:"typeId,omitempty"`
	Volume     *Network_Storage          `json:"volume,omitempty"`
	VolumeId   *int                      `json:"volumeId,omitempty"`
}

type Network_Storage_Group struct {
	Entity

	Account             *Account                       `json:"account,omitempty"`
	AccountId           *int                           `json:"accountId,omitempty"`
	Alias               *string                        `json:"alias,omitempty"`
	AllowedHostCount    *uint                          `json:"allowedHostCount,omitempty"`
	AllowedHosts        []Network_Storage_Allowed_Host `json:"allowedHosts,omitempty"`
	AttachedVolumeCount *uint                          `json:"attachedVolumeCount,omitempty"`
	AttachedVolumes     []Network_Storage              `json:"attachedVolumes,omitempty"`
	CreateDate          *time.Time                     `json:"createDate,omitempty"`
	GroupType           *Network_Storage_Group_Type    `json:"groupType,omitempty"`
	GroupTypeId         *int                           `json:"groupTypeId,omitempty"`
	Id                  *int                           `json:"id,omitempty"`
	ModifyDate          *time.Time                     `json:"modifyDate,omitempty"`
	OsType              *Network_Storage_Iscsi_OS_Type `json:"osType,omitempty"`
	OsTypeId            *int                           `json:"osTypeId,omitempty"`
	ServiceResource     *Network_Service_Resource      `json:"serviceResource,omitempty"`
	ServiceResourceId   *int                           `json:"serviceResourceId,omitempty"`
}

type Network_Storage_Group_Iscsi struct {
	Network_Storage_Group
}

type Network_Storage_Group_Nfs struct {
	Network_Storage_Group
}

type Network_Storage_Group_Type struct {
	Entity

	Id      *int    `json:"id,omitempty"`
	KeyName *string `json:"keyName,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type Network_Storage_History struct {
	Entity

	Account    *Account         `json:"account,omitempty"`
	CreateDate *time.Time       `json:"createDate,omitempty"`
	NasVolume  *Network_Storage `json:"nasVolume,omitempty"`
	Notes      *string          `json:"notes,omitempty"`
	Password   *string          `json:"password,omitempty"`
	Username   *string          `json:"username,omitempty"`
}

type Network_Storage_Hub struct {
	Network_Storage

	BandwidthBillingItemCount *uint          `json:"bandwidthBillingItemCount,omitempty"`
	BandwidthBillingItems     []Billing_Item `json:"bandwidthBillingItems,omitempty"`
}

type Network_Storage_Hub_Cleversafe_Account struct {
	Entity

	Account              *Account                     `json:"account,omitempty"`
	AccountId            *int                         `json:"accountId,omitempty"`
	BillingItem          *Billing_Item                `json:"billingItem,omitempty"`
	CancelledBillingItem *Billing_Item                `json:"cancelledBillingItem,omitempty"`
	CredentialCount      *uint                        `json:"credentialCount,omitempty"`
	Credentials          []Network_Storage_Credential `json:"credentials,omitempty"`
	EventCount           *uint                        `json:"eventCount,omitempty"`
	Events               []Network_Storage_Event      `json:"events,omitempty"`
	Id                   *int                         `json:"id,omitempty"`
	MetricTrackingObject *Metric_Tracking_Object      `json:"metricTrackingObject,omitempty"`
	NasType              *string                      `json:"nasType,omitempty"`
	Notes                *string                      `json:"notes,omitempty"`
	StorageTypeId        *int                         `json:"storageTypeId,omitempty"`
	Username             *string                      `json:"username,omitempty"`
	Uuid                 *string                      `json:"uuid,omitempty"`
}

type Network_Storage_Hub_Swift struct {
	Network_Storage_Hub

	StorageNodeCount *uint                      `json:"storageNodeCount,omitempty"`
	StorageNodes     []Network_Service_Resource `json:"storageNodes,omitempty"`
}

type Network_Storage_Hub_Swift_Container struct {
	Network_Storage_Hub_Swift
}

type Network_Storage_Hub_Swift_Share struct {
	Entity
}

type Network_Storage_Hub_Swift_Version1 struct {
	Network_Storage_Hub_Swift
}

type Network_Storage_Iscsi struct {
	Network_Storage
}

type Network_Storage_Iscsi_EqualLogic_Version3 struct {
	Network_Storage_Iscsi
}

type Network_Storage_Iscsi_EqualLogic_Version3_Replicant struct {
	Network_Storage_Iscsi_EqualLogic_Version3

	FailbackInProgressFlag *bool   `json:"failbackInProgressFlag,omitempty"`
	VolumeName             *string `json:"volumeName,omitempty"`
}

type Network_Storage_Iscsi_EqualLogic_Version3_Snapshot struct {
	Network_Storage_Iscsi_EqualLogic_Version3

	CreationSchedule *Network_Storage_Schedule `json:"creationSchedule,omitempty"`
	VolumeName       *string                   `json:"volumeName,omitempty"`
}

type Network_Storage_Iscsi_OS_Type struct {
	Entity

	CreateDate  *time.Time `json:"createDate,omitempty"`
	Description *string    `json:"description,omitempty"`
	Id          *int       `json:"id,omitempty"`
	KeyName     *string    `json:"keyName,omitempty"`
	Name        *string    `json:"name,omitempty"`
}

type Network_Storage_Nas struct {
	Network_Storage

	RecentBytesUsed *Network_Storage_Daily_Usage `json:"recentBytesUsed,omitempty"`
}

type Network_Storage_OpenStack_Object struct {
	Network_Storage

	BandwidthBillingItemCount *uint          `json:"bandwidthBillingItemCount,omitempty"`
	BandwidthBillingItems     []Billing_Item `json:"bandwidthBillingItems,omitempty"`
}

type Network_Storage_Partnership struct {
	Entity

	CreateDate      *time.Time                        `json:"createDate,omitempty"`
	ModifyDate      *time.Time                        `json:"modifyDate,omitempty"`
	PartnerVolume   *Network_Storage                  `json:"partnerVolume,omitempty"`
	PartnerVolumeId *int                              `json:"partnerVolumeId,omitempty"`
	Type            *Network_Storage_Partnership_Type `json:"type,omitempty"`
	Volume          *Network_Storage                  `json:"volume,omitempty"`
	VolumeId        *int                              `json:"volumeId,omitempty"`
}

type Network_Storage_Partnership_Type struct {
	Entity

	Description *string `json:"description,omitempty"`
	Keyname     *string `json:"keyname,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Network_Storage_Property struct {
	Entity

	CreateDate *time.Time                     `json:"createDate,omitempty"`
	ModifyDate *time.Time                     `json:"modifyDate,omitempty"`
	Type       *Network_Storage_Property_Type `json:"type,omitempty"`
	Value      *string                        `json:"value,omitempty"`
	Volume     *Network_Storage               `json:"volume,omitempty"`
	VolumeId   *int                           `json:"volumeId,omitempty"`
}

type Network_Storage_Property_Type struct {
	Entity

	Description *string `json:"description,omitempty"`
	Keyname     *string `json:"keyname,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Network_Storage_Replicant struct {
	Network_Storage

	FailbackInProgressFlag *string `json:"failbackInProgressFlag,omitempty"`
	VolumeName             *string `json:"volumeName,omitempty"`
}

type Network_Storage_Schedule struct {
	Entity

	Active               *int                                `json:"active,omitempty"`
	CreateDate           *time.Time                          `json:"createDate,omitempty"`
	DayOfMonth           *string                             `json:"dayOfMonth,omitempty"`
	DayOfWeek            *string                             `json:"dayOfWeek,omitempty"`
	EventCount           *uint                               `json:"eventCount,omitempty"`
	Events               []Network_Storage_Event             `json:"events,omitempty"`
	Hour                 *string                             `json:"hour,omitempty"`
	Id                   *int                                `json:"id,omitempty"`
	Minute               *string                             `json:"minute,omitempty"`
	ModifyDate           *time.Time                          `json:"modifyDate,omitempty"`
	MonthOfYear          *string                             `json:"monthOfYear,omitempty"`
	Name                 *string                             `json:"name,omitempty"`
	Partnership          *Network_Storage_Partnership        `json:"partnership,omitempty"`
	PartnershipId        *int                                `json:"partnershipId,omitempty"`
	Properties           []Network_Storage_Schedule_Property `json:"properties,omitempty"`
	PropertyCount        *uint                               `json:"propertyCount,omitempty"`
	ReplicaSnapshotCount *uint                               `json:"replicaSnapshotCount,omitempty"`
	ReplicaSnapshots     []Network_Storage                   `json:"replicaSnapshots,omitempty"`
	RetentionCount       *string                             `json:"retentionCount,omitempty"`
	SnapshotCount        *uint                               `json:"snapshotCount,omitempty"`
	Snapshots            []Network_Storage                   `json:"snapshots,omitempty"`
	Type                 *Network_Storage_Schedule_Type      `json:"type,omitempty"`
	TypeId               *int                                `json:"typeId,omitempty"`
	Volume               *Network_Storage                    `json:"volume,omitempty"`
	VolumeId             *int                                `json:"volumeId,omitempty"`
}

type Network_Storage_Schedule_Property struct {
	Entity

	CreateDate *time.Time                              `json:"createDate,omitempty"`
	Id         *int                                    `json:"id,omitempty"`
	ModifyDate *time.Time                              `json:"modifyDate,omitempty"`
	Schedule   *Network_Storage_Schedule               `json:"schedule,omitempty"`
	Type       *Network_Storage_Schedule_Property_Type `json:"type,omitempty"`
	TypeId     *int                                    `json:"typeId,omitempty"`
	Value      *string                                 `json:"value,omitempty"`
}

type Network_Storage_Schedule_Property_Type struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	Keyname     *string `json:"keyname,omitempty"`
	Name        *string `json:"name,omitempty"`
	NasType     *string `json:"nasType,omitempty"`
}

type Network_Storage_Schedule_Type struct {
	Entity

	Id      *int    `json:"id,omitempty"`
	Keyname *string `json:"keyname,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type Network_Storage_Snapshot struct {
	Network_Storage

	CreationSchedule *Network_Storage_Schedule `json:"creationSchedule,omitempty"`
	VolumeName       *string                   `json:"volumeName,omitempty"`
}

type Network_Storage_Type struct {
	Entity

	Description *string           `json:"description,omitempty"`
	Id          *int              `json:"id,omitempty"`
	KeyName     *string           `json:"keyName,omitempty"`
	VolumeCount *uint             `json:"volumeCount,omitempty"`
	Volumes     []Network_Storage `json:"volumes,omitempty"`
}

type Network_Subnet struct {
	Entity

	Account                           *Account                            `json:"account,omitempty"`
	ActiveRegistration                *Network_Subnet_Registration        `json:"activeRegistration,omitempty"`
	ActiveSwipTransaction             *Network_Subnet_Swip_Transaction    `json:"activeSwipTransaction,omitempty"`
	ActiveTransaction                 *Provisioning_Version1_Transaction  `json:"activeTransaction,omitempty"`
	AddressSpace                      *string                             `json:"addressSpace,omitempty"`
	AllowedHost                       *Network_Storage_Allowed_Host       `json:"allowedHost,omitempty"`
	AllowedNetworkStorage             []Network_Storage                   `json:"allowedNetworkStorage,omitempty"`
	AllowedNetworkStorageCount        *uint                               `json:"allowedNetworkStorageCount,omitempty"`
	AllowedNetworkStorageReplicaCount *uint                               `json:"allowedNetworkStorageReplicaCount,omitempty"`
	AllowedNetworkStorageReplicas     []Network_Storage                   `json:"allowedNetworkStorageReplicas,omitempty"`
	BillingItem                       *Billing_Item                       `json:"billingItem,omitempty"`
	BoundDescendantCount              *uint                               `json:"boundDescendantCount,omitempty"`
	BoundDescendants                  []Network_Subnet                    `json:"boundDescendants,omitempty"`
	BoundRouterCount                  *uint                               `json:"boundRouterCount,omitempty"`
	BoundRouterFlag                   *bool                               `json:"boundRouterFlag,omitempty"`
	BoundRouters                      []Hardware                          `json:"boundRouters,omitempty"`
	BroadcastAddress                  *string                             `json:"broadcastAddress,omitempty"`
	Children                          []Network_Subnet                    `json:"children,omitempty"`
	ChildrenCount                     *uint                               `json:"childrenCount,omitempty"`
	Cidr                              *int                                `json:"cidr,omitempty"`
	Datacenter                        *Location_Datacenter                `json:"datacenter,omitempty"`
	DescendantCount                   *uint                               `json:"descendantCount,omitempty"`
	Descendants                       []Network_Subnet                    `json:"descendants,omitempty"`
	DisplayLabel                      *string                             `json:"displayLabel,omitempty"`
	EndPointIpAddress                 *Network_Subnet_IpAddress           `json:"endPointIpAddress,omitempty"`
	Gateway                           *string                             `json:"gateway,omitempty"`
	GlobalIpRecord                    *Network_Subnet_IpAddress_Global    `json:"globalIpRecord,omitempty"`
	Hardware                          []Hardware                          `json:"hardware,omitempty"`
	HardwareCount                     *uint                               `json:"hardwareCount,omitempty"`
	Id                                *int                                `json:"id,omitempty"`
	IpAddressCount                    *uint                               `json:"ipAddressCount,omitempty"`
	IpAddresses                       []Network_Subnet_IpAddress          `json:"ipAddresses,omitempty"`
	IsCustomerOwned                   *bool                               `json:"isCustomerOwned,omitempty"`
	IsCustomerRoutable                *bool                               `json:"isCustomerRoutable,omitempty"`
	ModifyDate                        *time.Time                          `json:"modifyDate,omitempty"`
	Netmask                           *string                             `json:"netmask,omitempty"`
	NetworkComponent                  *Network_Component                  `json:"networkComponent,omitempty"`
	NetworkComponentFirewall          *Network_Component_Firewall         `json:"networkComponentFirewall,omitempty"`
	NetworkIdentifier                 *string                             `json:"networkIdentifier,omitempty"`
	NetworkProtectionAddressCount     *uint                               `json:"networkProtectionAddressCount,omitempty"`
	NetworkProtectionAddresses        []Network_Protection_Address        `json:"networkProtectionAddresses,omitempty"`
	NetworkTunnelContextCount         *uint                               `json:"networkTunnelContextCount,omitempty"`
	NetworkTunnelContexts             []Network_Tunnel_Module_Context     `json:"networkTunnelContexts,omitempty"`
	NetworkVlan                       *Network_Vlan                       `json:"networkVlan,omitempty"`
	NetworkVlanId                     *int                                `json:"networkVlanId,omitempty"`
	Note                              *string                             `json:"note,omitempty"`
	PodName                           *string                             `json:"podName,omitempty"`
	ProtectedIpAddressCount           *uint                               `json:"protectedIpAddressCount,omitempty"`
	ProtectedIpAddresses              []Network_Subnet_IpAddress          `json:"protectedIpAddresses,omitempty"`
	RegionalInternetRegistry          *Network_Regional_Internet_Registry `json:"regionalInternetRegistry,omitempty"`
	RegistrationCount                 *uint                               `json:"registrationCount,omitempty"`
	Registrations                     []Network_Subnet_Registration       `json:"registrations,omitempty"`
	ResourceGroupCount                *uint                               `json:"resourceGroupCount,omitempty"`
	ResourceGroups                    []Resource_Group                    `json:"resourceGroups,omitempty"`
	ReverseDomain                     *Dns_Domain                         `json:"reverseDomain,omitempty"`
	RoleKeyName                       *string                             `json:"roleKeyName,omitempty"`
	RoleName                          *string                             `json:"roleName,omitempty"`
	RoutingTypeKeyName                *string                             `json:"routingTypeKeyName,omitempty"`
	RoutingTypeName                   *string                             `json:"routingTypeName,omitempty"`
	SortOrder                         *string                             `json:"sortOrder,omitempty"`
	SubnetType                        *string                             `json:"subnetType,omitempty"`
	SwipTransaction                   []Network_Subnet_Swip_Transaction   `json:"swipTransaction,omitempty"`
	SwipTransactionCount              *uint                               `json:"swipTransactionCount,omitempty"`
	TotalIpAddresses                  *float64                            `json:"totalIpAddresses,omitempty"`
	UnboundDescendantCount            *uint                               `json:"unboundDescendantCount,omitempty"`
	UnboundDescendants                []Network_Subnet                    `json:"unboundDescendants,omitempty"`
	UsableIpAddressCount              *float64                            `json:"usableIpAddressCount,omitempty"`
	Version                           *int                                `json:"version,omitempty"`
	VirtualGuestCount                 *uint                               `json:"virtualGuestCount,omitempty"`
	VirtualGuests                     []Virtual_Guest                     `json:"virtualGuests,omitempty"`
}

type Network_Subnet_IpAddress struct {
	Entity

	AllowedHost                                      *Network_Storage_Allowed_Host                       `json:"allowedHost,omitempty"`
	AllowedNetworkStorage                            []Network_Storage                                   `json:"allowedNetworkStorage,omitempty"`
	AllowedNetworkStorageCount                       *uint                                               `json:"allowedNetworkStorageCount,omitempty"`
	AllowedNetworkStorageReplicaCount                *uint                                               `json:"allowedNetworkStorageReplicaCount,omitempty"`
	AllowedNetworkStorageReplicas                    []Network_Storage                                   `json:"allowedNetworkStorageReplicas,omitempty"`
	ApplicationDeliveryController                    *Network_Application_Delivery_Controller            `json:"applicationDeliveryController,omitempty"`
	ContextTunnelTranslationCount                    *uint                                               `json:"contextTunnelTranslationCount,omitempty"`
	ContextTunnelTranslations                        []Network_Tunnel_Module_Context_Address_Translation `json:"contextTunnelTranslations,omitempty"`
	EndpointSubnetCount                              *uint                                               `json:"endpointSubnetCount,omitempty"`
	EndpointSubnets                                  []Network_Subnet                                    `json:"endpointSubnets,omitempty"`
	GuestNetworkComponent                            *Virtual_Guest_Network_Component                    `json:"guestNetworkComponent,omitempty"`
	GuestNetworkComponentBinding                     *Virtual_Guest_Network_Component_IpAddress          `json:"guestNetworkComponentBinding,omitempty"`
	Hardware                                         *Hardware                                           `json:"hardware,omitempty"`
	Id                                               *int                                                `json:"id,omitempty"`
	IpAddress                                        *string                                             `json:"ipAddress,omitempty"`
	IsBroadcast                                      *bool                                               `json:"isBroadcast,omitempty"`
	IsGateway                                        *bool                                               `json:"isGateway,omitempty"`
	IsNetwork                                        *bool                                               `json:"isNetwork,omitempty"`
	IsReserved                                       *bool                                               `json:"isReserved,omitempty"`
	NetworkComponent                                 *Network_Component                                  `json:"networkComponent,omitempty"`
	Note                                             *string                                             `json:"note,omitempty"`
	PrivateNetworkGateway                            *Network_Gateway                                    `json:"privateNetworkGateway,omitempty"`
	ProtectionAddress                                []Network_Protection_Address                        `json:"protectionAddress,omitempty"`
	ProtectionAddressCount                           *uint                                               `json:"protectionAddressCount,omitempty"`
	PublicNetworkGateway                             *Network_Gateway                                    `json:"publicNetworkGateway,omitempty"`
	RemoteManagementNetworkComponent                 *Network_Component                                  `json:"remoteManagementNetworkComponent,omitempty"`
	Subnet                                           *Network_Subnet                                     `json:"subnet,omitempty"`
	SubnetId                                         *int                                                `json:"subnetId,omitempty"`
	SyslogEventsOneDay                               []Network_Logging_Syslog                            `json:"syslogEventsOneDay,omitempty"`
	SyslogEventsOneDayCount                          *uint                                               `json:"syslogEventsOneDayCount,omitempty"`
	SyslogEventsSevenDayCount                        *uint                                               `json:"syslogEventsSevenDayCount,omitempty"`
	SyslogEventsSevenDays                            []Network_Logging_Syslog                            `json:"syslogEventsSevenDays,omitempty"`
	TopTenSyslogEventsByDestinationPortOneDay        []Network_Logging_Syslog                            `json:"topTenSyslogEventsByDestinationPortOneDay,omitempty"`
	TopTenSyslogEventsByDestinationPortOneDayCount   *uint                                               `json:"topTenSyslogEventsByDestinationPortOneDayCount,omitempty"`
	TopTenSyslogEventsByDestinationPortSevenDayCount *uint                                               `json:"topTenSyslogEventsByDestinationPortSevenDayCount,omitempty"`
	TopTenSyslogEventsByDestinationPortSevenDays     []Network_Logging_Syslog                            `json:"topTenSyslogEventsByDestinationPortSevenDays,omitempty"`
	TopTenSyslogEventsByProtocolsOneDay              []Network_Logging_Syslog                            `json:"topTenSyslogEventsByProtocolsOneDay,omitempty"`
	TopTenSyslogEventsByProtocolsOneDayCount         *uint                                               `json:"topTenSyslogEventsByProtocolsOneDayCount,omitempty"`
	TopTenSyslogEventsByProtocolsSevenDayCount       *uint                                               `json:"topTenSyslogEventsByProtocolsSevenDayCount,omitempty"`
	TopTenSyslogEventsByProtocolsSevenDays           []Network_Logging_Syslog                            `json:"topTenSyslogEventsByProtocolsSevenDays,omitempty"`
	TopTenSyslogEventsBySourceIpOneDay               []Network_Logging_Syslog                            `json:"topTenSyslogEventsBySourceIpOneDay,omitempty"`
	TopTenSyslogEventsBySourceIpOneDayCount          *uint                                               `json:"topTenSyslogEventsBySourceIpOneDayCount,omitempty"`
	TopTenSyslogEventsBySourceIpSevenDayCount        *uint                                               `json:"topTenSyslogEventsBySourceIpSevenDayCount,omitempty"`
	TopTenSyslogEventsBySourceIpSevenDays            []Network_Logging_Syslog                            `json:"topTenSyslogEventsBySourceIpSevenDays,omitempty"`
	TopTenSyslogEventsBySourcePortOneDay             []Network_Logging_Syslog                            `json:"topTenSyslogEventsBySourcePortOneDay,omitempty"`
	TopTenSyslogEventsBySourcePortOneDayCount        *uint                                               `json:"topTenSyslogEventsBySourcePortOneDayCount,omitempty"`
	TopTenSyslogEventsBySourcePortSevenDayCount      *uint                                               `json:"topTenSyslogEventsBySourcePortSevenDayCount,omitempty"`
	TopTenSyslogEventsBySourcePortSevenDays          []Network_Logging_Syslog                            `json:"topTenSyslogEventsBySourcePortSevenDays,omitempty"`
	VirtualGuest                                     *Virtual_Guest                                      `json:"virtualGuest,omitempty"`
	VirtualLicenseCount                              *uint                                               `json:"virtualLicenseCount,omitempty"`
	VirtualLicenses                                  []Software_VirtualLicense                           `json:"virtualLicenses,omitempty"`
}

type Network_Subnet_IpAddress_Global struct {
	Entity

	Account                *Account                                      `json:"account,omitempty"`
	ActiveTransaction      *Provisioning_Version1_Transaction            `json:"activeTransaction,omitempty"`
	BillingItem            *Billing_Item_Network_Subnet_IpAddress_Global `json:"billingItem,omitempty"`
	Description            *int                                          `json:"description,omitempty"`
	DestinationIpAddress   *Network_Subnet_IpAddress                     `json:"destinationIpAddress,omitempty"`
	DestinationIpAddressId *int                                          `json:"destinationIpAddressId,omitempty"`
	Id                     *int                                          `json:"id,omitempty"`
	IpAddress              *Network_Subnet_IpAddress                     `json:"ipAddress,omitempty"`
	IpAddressId            *int                                          `json:"ipAddressId,omitempty"`
	TypeId                 *int                                          `json:"typeId,omitempty"`
}

type Network_Subnet_IpAddress_Version6 struct {
	Network_Subnet_IpAddress

	PublicVersion6NetworkGateway *Network_Gateway `json:"publicVersion6NetworkGateway,omitempty"`
}

type Network_Subnet_Registration struct {
	Entity

	Account                          *Account                              `json:"account,omitempty"`
	AccountId                        *int                                  `json:"accountId,omitempty"`
	Cidr                             *int                                  `json:"cidr,omitempty"`
	CreateDate                       *time.Time                            `json:"createDate,omitempty"`
	DetailReferenceCount             *uint                                 `json:"detailReferenceCount,omitempty"`
	DetailReferences                 []Network_Subnet_Registration_Details `json:"detailReferences,omitempty"`
	EventCount                       *uint                                 `json:"eventCount,omitempty"`
	Events                           []Network_Subnet_Registration_Event   `json:"events,omitempty"`
	Id                               *int                                  `json:"id,omitempty"`
	ModifyDate                       *time.Time                            `json:"modifyDate,omitempty"`
	NetworkDetail                    *Account_Regional_Registry_Detail     `json:"networkDetail,omitempty"`
	NetworkHandle                    *string                               `json:"networkHandle,omitempty"`
	NetworkIdentifier                *string                               `json:"networkIdentifier,omitempty"`
	PersonDetail                     *Account_Regional_Registry_Detail     `json:"personDetail,omitempty"`
	RegionalInternetRegistry         *Network_Regional_Internet_Registry   `json:"regionalInternetRegistry,omitempty"`
	RegionalInternetRegistryHandle   *Account_Rwhois_Handle                `json:"regionalInternetRegistryHandle,omitempty"`
	RegionalInternetRegistryHandleId *int                                  `json:"regionalInternetRegistryHandleId,omitempty"`
	RegionalInternetRegistryId       *int                                  `json:"regionalInternetRegistryId,omitempty"`
	Status                           *Network_Subnet_Registration_Status   `json:"status,omitempty"`
	StatusId                         *int                                  `json:"statusId,omitempty"`
	Subnet                           *Network_Subnet                       `json:"subnet,omitempty"`
}

type Network_Subnet_Registration_Apnic struct {
	Network_Subnet_Registration
}

type Network_Subnet_Registration_Arin struct {
	Network_Subnet_Registration
}

type Network_Subnet_Registration_Details struct {
	Entity

	CreateDate     *time.Time                        `json:"createDate,omitempty"`
	Detail         *Account_Regional_Registry_Detail `json:"detail,omitempty"`
	DetailId       *int                              `json:"detailId,omitempty"`
	Id             *int                              `json:"id,omitempty"`
	ModifyDate     *time.Time                        `json:"modifyDate,omitempty"`
	Registration   *Network_Subnet_Registration      `json:"registration,omitempty"`
	RegistrationId *int                              `json:"registrationId,omitempty"`
}

type Network_Subnet_Registration_Event struct {
	Entity

	CreateDate     *time.Time                              `json:"createDate,omitempty"`
	Id             *int                                    `json:"id,omitempty"`
	Message        *string                                 `json:"message,omitempty"`
	ModifyDate     *time.Time                              `json:"modifyDate,omitempty"`
	Registration   *Network_Subnet_Registration            `json:"registration,omitempty"`
	RegistrationId *int                                    `json:"registrationId,omitempty"`
	Type           *Network_Subnet_Registration_Event_Type `json:"type,omitempty"`
	TypeId         *int                                    `json:"typeId,omitempty"`
}

type Network_Subnet_Registration_Event_Type struct {
	Entity

	CreateDate *time.Time `json:"createDate,omitempty"`
	Id         *int       `json:"id,omitempty"`
	KeyName    *string    `json:"keyName,omitempty"`
	ModifyDate *time.Time `json:"modifyDate,omitempty"`
	Name       *string    `json:"name,omitempty"`
}

type Network_Subnet_Registration_Ripe struct {
	Network_Subnet_Registration
}

type Network_Subnet_Registration_Status struct {
	Entity

	CreateDate *time.Time `json:"createDate,omitempty"`
	Id         *int       `json:"id,omitempty"`
	KeyName    *string    `json:"keyName,omitempty"`
	ModifyDate *time.Time `json:"modifyDate,omitempty"`
	Name       *string    `json:"name,omitempty"`
}

type Network_Subnet_Rwhois_Data struct {
	Entity

	AbuseEmail           *string    `json:"abuseEmail,omitempty"`
	Account              *Account   `json:"account,omitempty"`
	AccountId            *int       `json:"accountId,omitempty"`
	Address1             *string    `json:"address1,omitempty"`
	Address2             *string    `json:"address2,omitempty"`
	City                 *string    `json:"city,omitempty"`
	CompanyName          *string    `json:"companyName,omitempty"`
	Country              *string    `json:"country,omitempty"`
	CreateDate           *time.Time `json:"createDate,omitempty"`
	FirstName            *string    `json:"firstName,omitempty"`
	Id                   *int       `json:"id,omitempty"`
	LastName             *string    `json:"lastName,omitempty"`
	ModifyDate           *time.Time `json:"modifyDate,omitempty"`
	PostalCode           *string    `json:"postalCode,omitempty"`
	PrivateResidenceFlag *bool      `json:"privateResidenceFlag,omitempty"`
	State                *string    `json:"state,omitempty"`
}

type Network_Subnet_Swip_Transaction struct {
	Entity

	Account    *Account        `json:"account,omitempty"`
	Id         *int            `json:"id,omitempty"`
	StatusName *string         `json:"statusName,omitempty"`
	Subnet     *Network_Subnet `json:"subnet,omitempty"`
	SubnetId   *int            `json:"subnetId,omitempty"`
}

type Network_TippingPointReporting struct {
	Entity
}

type Network_Tunnel_Module_Context struct {
	Entity

	Account                        *Account                                            `json:"account,omitempty"`
	AccountId                      *int                                                `json:"accountId,omitempty"`
	ActiveTransaction              *Provisioning_Version1_Transaction                  `json:"activeTransaction,omitempty"`
	AddressTranslationCount        *uint                                               `json:"addressTranslationCount,omitempty"`
	AddressTranslations            []Network_Tunnel_Module_Context_Address_Translation `json:"addressTranslations,omitempty"`
	AdvancedConfigurationFlag      *int                                                `json:"advancedConfigurationFlag,omitempty"`
	AllAvailableServiceSubnetCount *uint                                               `json:"allAvailableServiceSubnetCount,omitempty"`
	AllAvailableServiceSubnets     []Network_Subnet                                    `json:"allAvailableServiceSubnets,omitempty"`
	BillingItem                    *Billing_Item                                       `json:"billingItem,omitempty"`
	CreateDate                     *time.Time                                          `json:"createDate,omitempty"`
	CustomerPeerIpAddress          *string                                             `json:"customerPeerIpAddress,omitempty"`
	CustomerSubnetCount            *uint                                               `json:"customerSubnetCount,omitempty"`
	CustomerSubnets                []Network_Customer_Subnet                           `json:"customerSubnets,omitempty"`
	Datacenter                     *Location                                           `json:"datacenter,omitempty"`
	FriendlyName                   *string                                             `json:"friendlyName,omitempty"`
	Id                             *int                                                `json:"id,omitempty"`
	InternalPeerIpAddress          *string                                             `json:"internalPeerIpAddress,omitempty"`
	InternalSubnetCount            *uint                                               `json:"internalSubnetCount,omitempty"`
	InternalSubnets                []Network_Subnet                                    `json:"internalSubnets,omitempty"`
	ModifyDate                     *time.Time                                          `json:"modifyDate,omitempty"`
	Name                           *string                                             `json:"name,omitempty"`
	PhaseOneAuthentication         *string                                             `json:"phaseOneAuthentication,omitempty"`
	PhaseOneDiffieHellmanGroup     *int                                                `json:"phaseOneDiffieHellmanGroup,omitempty"`
	PhaseOneEncryption             *string                                             `json:"phaseOneEncryption,omitempty"`
	PhaseOneKeylife                *int                                                `json:"phaseOneKeylife,omitempty"`
	PhaseTwoAuthentication         *string                                             `json:"phaseTwoAuthentication,omitempty"`
	PhaseTwoDiffieHellmanGroup     *int                                                `json:"phaseTwoDiffieHellmanGroup,omitempty"`
	PhaseTwoEncryption             *string                                             `json:"phaseTwoEncryption,omitempty"`
	PhaseTwoKeylife                *int                                                `json:"phaseTwoKeylife,omitempty"`
	PhaseTwoPerfectForwardSecrecy  *int                                                `json:"phaseTwoPerfectForwardSecrecy,omitempty"`
	PresharedKey                   *string                                             `json:"presharedKey,omitempty"`
	ServiceSubnetCount             *uint                                               `json:"serviceSubnetCount,omitempty"`
	ServiceSubnets                 []Network_Subnet                                    `json:"serviceSubnets,omitempty"`
	StaticRouteSubnetCount         *uint                                               `json:"staticRouteSubnetCount,omitempty"`
	StaticRouteSubnets             []Network_Subnet                                    `json:"staticRouteSubnets,omitempty"`
	TransactionHistory             []Provisioning_Version1_Transaction                 `json:"transactionHistory,omitempty"`
	TransactionHistoryCount        *uint                                               `json:"transactionHistoryCount,omitempty"`
}

type Network_Tunnel_Module_Context_Address_Translation struct {
	Entity

	CustomerIpAddress       *string                            `json:"customerIpAddress,omitempty"`
	CustomerIpAddressId     *int                               `json:"customerIpAddressId,omitempty"`
	CustomerIpAddressRecord *Network_Customer_Subnet_IpAddress `json:"customerIpAddressRecord,omitempty"`
	Id                      *int                               `json:"id,omitempty"`
	InternalIpAddress       *string                            `json:"internalIpAddress,omitempty"`
	InternalIpAddressId     *int                               `json:"internalIpAddressId,omitempty"`
	InternalIpAddressRecord *Network_Subnet_IpAddress          `json:"internalIpAddressRecord,omitempty"`
	NetworkTunnelContext    *Network_Tunnel_Module_Context     `json:"networkTunnelContext,omitempty"`
	NetworkTunnelContextId  *int                               `json:"networkTunnelContextId,omitempty"`
	Notes                   *string                            `json:"notes,omitempty"`
}

type Network_Vlan struct {
	Entity

	Account                            *Account                                    `json:"account,omitempty"`
	AccountId                          *int                                        `json:"accountId,omitempty"`
	AdditionalPrimarySubnetCount       *uint                                       `json:"additionalPrimarySubnetCount,omitempty"`
	AdditionalPrimarySubnets           []Network_Subnet                            `json:"additionalPrimarySubnets,omitempty"`
	AttachedNetworkGateway             *Network_Gateway                            `json:"attachedNetworkGateway,omitempty"`
	AttachedNetworkGatewayFlag         *bool                                       `json:"attachedNetworkGatewayFlag,omitempty"`
	AttachedNetworkGatewayVlan         *Network_Gateway_Vlan                       `json:"attachedNetworkGatewayVlan,omitempty"`
	BillingItem                        *Billing_Item                               `json:"billingItem,omitempty"`
	DedicatedFirewallFlag              *int                                        `json:"dedicatedFirewallFlag,omitempty"`
	ExtensionRouter                    *Hardware_Router                            `json:"extensionRouter,omitempty"`
	FirewallGuestNetworkComponentCount *uint                                       `json:"firewallGuestNetworkComponentCount,omitempty"`
	FirewallGuestNetworkComponents     []Network_Component_Firewall                `json:"firewallGuestNetworkComponents,omitempty"`
	FirewallInterfaceCount             *uint                                       `json:"firewallInterfaceCount,omitempty"`
	FirewallInterfaces                 []Network_Firewall_Module_Context_Interface `json:"firewallInterfaces,omitempty"`
	FirewallNetworkComponentCount      *uint                                       `json:"firewallNetworkComponentCount,omitempty"`
	FirewallNetworkComponents          []Network_Component_Firewall                `json:"firewallNetworkComponents,omitempty"`
	FirewallRuleCount                  *uint                                       `json:"firewallRuleCount,omitempty"`
	FirewallRules                      []Network_Vlan_Firewall_Rule                `json:"firewallRules,omitempty"`
	GuestNetworkComponentCount         *uint                                       `json:"guestNetworkComponentCount,omitempty"`
	GuestNetworkComponents             []Virtual_Guest_Network_Component           `json:"guestNetworkComponents,omitempty"`
	Hardware                           []Hardware                                  `json:"hardware,omitempty"`
	HardwareCount                      *uint                                       `json:"hardwareCount,omitempty"`
	HighAvailabilityFirewallFlag       *bool                                       `json:"highAvailabilityFirewallFlag,omitempty"`
	Id                                 *int                                        `json:"id,omitempty"`
	LocalDiskStorageCapabilityFlag     *bool                                       `json:"localDiskStorageCapabilityFlag,omitempty"`
	ModifyDate                         *time.Time                                  `json:"modifyDate,omitempty"`
	Name                               *string                                     `json:"name,omitempty"`
	Network                            *Network                                    `json:"network,omitempty"`
	NetworkComponentCount              *uint                                       `json:"networkComponentCount,omitempty"`
	NetworkComponentTrunkCount         *uint                                       `json:"networkComponentTrunkCount,omitempty"`
	NetworkComponentTrunks             []Network_Component_Network_Vlan_Trunk      `json:"networkComponentTrunks,omitempty"`
	NetworkComponents                  []Network_Component                         `json:"networkComponents,omitempty"`
	NetworkSpace                       *string                                     `json:"networkSpace,omitempty"`
	NetworkVlanFirewall                *Network_Vlan_Firewall                      `json:"networkVlanFirewall,omitempty"`
	Note                               *string                                     `json:"note,omitempty"`
	PrimaryRouter                      *Hardware_Router                            `json:"primaryRouter,omitempty"`
	PrimarySubnet                      *Network_Subnet                             `json:"primarySubnet,omitempty"`
	PrimarySubnetCount                 *uint                                       `json:"primarySubnetCount,omitempty"`
	PrimarySubnetId                    *int                                        `json:"primarySubnetId,omitempty"`
	PrimarySubnetVersion6              *Network_Subnet                             `json:"primarySubnetVersion6,omitempty"`
	PrimarySubnets                     []Network_Subnet                            `json:"primarySubnets,omitempty"`
	PrivateNetworkGatewayCount         *uint                                       `json:"privateNetworkGatewayCount,omitempty"`
	PrivateNetworkGateways             []Network_Gateway                           `json:"privateNetworkGateways,omitempty"`
	ProtectedIpAddressCount            *uint                                       `json:"protectedIpAddressCount,omitempty"`
	ProtectedIpAddresses               []Network_Subnet_IpAddress                  `json:"protectedIpAddresses,omitempty"`
	PublicNetworkGatewayCount          *uint                                       `json:"publicNetworkGatewayCount,omitempty"`
	PublicNetworkGateways              []Network_Gateway                           `json:"publicNetworkGateways,omitempty"`
	ResourceGroupCount                 *uint                                       `json:"resourceGroupCount,omitempty"`
	ResourceGroupMember                []Resource_Group_Member                     `json:"resourceGroupMember,omitempty"`
	ResourceGroupMemberCount           *uint                                       `json:"resourceGroupMemberCount,omitempty"`
	ResourceGroups                     []Resource_Group                            `json:"resourceGroups,omitempty"`
	SanStorageCapabilityFlag           *bool                                       `json:"sanStorageCapabilityFlag,omitempty"`
	ScaleVlanCount                     *uint                                       `json:"scaleVlanCount,omitempty"`
	ScaleVlans                         []Scale_Network_Vlan                        `json:"scaleVlans,omitempty"`
	SecondaryRouter                    *Hardware                                   `json:"secondaryRouter,omitempty"`
	SecondarySubnetCount               *uint                                       `json:"secondarySubnetCount,omitempty"`
	SecondarySubnets                   []Network_Subnet                            `json:"secondarySubnets,omitempty"`
	SubnetCount                        *uint                                       `json:"subnetCount,omitempty"`
	Subnets                            []Network_Subnet                            `json:"subnets,omitempty"`
	TagReferenceCount                  *uint                                       `json:"tagReferenceCount,omitempty"`
	TagReferences                      []Tag_Reference                             `json:"tagReferences,omitempty"`
	TotalPrimaryIpAddressCount         *uint                                       `json:"totalPrimaryIpAddressCount,omitempty"`
	Type                               *Network_Vlan_Type                          `json:"type,omitempty"`
	VirtualGuestCount                  *uint                                       `json:"virtualGuestCount,omitempty"`
	VirtualGuests                      []Virtual_Guest                             `json:"virtualGuests,omitempty"`
	VlanNumber                         *int                                        `json:"vlanNumber,omitempty"`
}

type Network_Vlan_Firewall struct {
	Entity

	AdministrativeBypassFlag          *string                           `json:"administrativeBypassFlag,omitempty"`
	BillingItem                       *Billing_Item                     `json:"billingItem,omitempty"`
	CustomerManagedFlag               *bool                             `json:"customerManagedFlag,omitempty"`
	Datacenter                        *Location                         `json:"datacenter,omitempty"`
	FirewallType                      *string                           `json:"firewallType,omitempty"`
	FullyQualifiedDomainName          *string                           `json:"fullyQualifiedDomainName,omitempty"`
	Id                                *int                              `json:"id,omitempty"`
	ManagementCredentials             *Software_Component_Password      `json:"managementCredentials,omitempty"`
	NetworkFirewallUpdateRequestCount *uint                             `json:"networkFirewallUpdateRequestCount,omitempty"`
	NetworkFirewallUpdateRequests     []Network_Firewall_Update_Request `json:"networkFirewallUpdateRequests,omitempty"`
	NetworkVlan                       *Network_Vlan                     `json:"networkVlan,omitempty"`
	NetworkVlanCount                  *uint                             `json:"networkVlanCount,omitempty"`
	NetworkVlans                      []Network_Vlan                    `json:"networkVlans,omitempty"`
	PrimaryIpAddress                  *string                           `json:"primaryIpAddress,omitempty"`
	RuleCount                         *uint                             `json:"ruleCount,omitempty"`
	Rules                             []Network_Vlan_Firewall_Rule      `json:"rules,omitempty"`
	TagReferenceCount                 *uint                             `json:"tagReferenceCount,omitempty"`
	TagReferences                     []Tag_Reference                   `json:"tagReferences,omitempty"`
}

type Network_Vlan_Firewall_Rule struct {
	Entity

	Action                    *string                     `json:"action,omitempty"`
	DestinationIpAddress      *string                     `json:"destinationIpAddress,omitempty"`
	DestinationIpCidr         *int                        `json:"destinationIpCidr,omitempty"`
	DestinationIpSubnetMask   *string                     `json:"destinationIpSubnetMask,omitempty"`
	DestinationPortRangeEnd   *int                        `json:"destinationPortRangeEnd,omitempty"`
	DestinationPortRangeStart *int                        `json:"destinationPortRangeStart,omitempty"`
	Id                        *int                        `json:"id,omitempty"`
	NetworkComponentFirewall  *Network_Component_Firewall `json:"networkComponentFirewall,omitempty"`
	Notes                     *string                     `json:"notes,omitempty"`
	OrderValue                *int                        `json:"orderValue,omitempty"`
	Protocol                  *string                     `json:"protocol,omitempty"`
	SourceIpAddress           *string                     `json:"sourceIpAddress,omitempty"`
	SourceIpCidr              *int                        `json:"sourceIpCidr,omitempty"`
	SourceIpSubnetMask        *string                     `json:"sourceIpSubnetMask,omitempty"`
	Status                    *string                     `json:"status,omitempty"`
	Version                   *int                        `json:"version,omitempty"`
}

type Network_Vlan_Type struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	KeyName     *string `json:"keyName,omitempty"`
	Name        *string `json:"name,omitempty"`
}
