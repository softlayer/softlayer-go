package softlayer

import (
	"time"

	"github.ibm.com/riethm/gopherlayer/datatypes"
)

type Network interface {
	Entity

	CreateObject(templateObject datatypes.Network) (datatypes.Network, error)
	CreateSubnet(subnet datatypes.Network_Subnet, pod datatypes.Network_Pod) (datatypes.Network_Subnet, error)
	DeleteObject() (bool, error)
	DeleteSubnet(subnet datatypes.Network_Subnet) (bool, error)
	EditObject(templateObject datatypes.Network) (bool, error)
	GetAllObjects() (datatypes.Network, error)
	GetObject() (datatypes.Network, error)

	GetCidr() (int, error)
	GetName() (string, error)
	GetNetworkIdentifier() (string, error)
	GetNotes() (string, error)
	GetSubnets() (datatypes.Network_Subnet, error)
}

type Network_Application_Delivery_Controller interface {
	Entity

	CreateLiveLoadBalancer(loadBalancer datatypes.Network_LoadBalancer_VirtualIpAddress) (bool, error)
	DeleteLiveLoadBalancer(loadBalancer datatypes.Network_LoadBalancer_VirtualIpAddress) (bool, error)
	DeleteLiveLoadBalancerService(service datatypes.Network_LoadBalancer_Service) (bool, error)
	EditObject(templateObject datatypes.Network_Application_Delivery_Controller) (bool, error)
	GetBandwidthDataByDate(startDateTime time.Time, endDateTime time.Time, networkType string) (datatypes.Metric_Tracking_Object_Data, error)
	GetBandwidthImageByDate(startDateTime time.Time, endDateTime time.Time, networkType string) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetCustomBandwidthDataByDate(graphData datatypes.Container_Graph) (datatypes.Container_Graph, error)
	GetLiveLoadBalancerServiceGraphImage(service datatypes.Network_LoadBalancer_Service, graphType string, metric string) ([]byte, error)
	GetObject() (datatypes.Network_Application_Delivery_Controller, error)
	RestoreBaseConfiguration() (bool, error)
	RestoreConfiguration(configurationHistoryId int) (bool, error)
	SaveCurrentConfiguration(notes string) (datatypes.Network_Application_Delivery_Controller_Configuration_History, error)
	UpdateLiveLoadBalancer(loadBalancer datatypes.Network_LoadBalancer_VirtualIpAddress) (bool, error)
	UpdateNetScalerLicense() (datatypes.Provisioning_Version1_Transaction, error)

	GetAccount() (datatypes.Account, error)
	GetAverageDailyPublicBandwidthUsage() (float64, error)
	GetBillingItem() (datatypes.Billing_Item_Network_Application_Delivery_Controller, error)
	GetConfigurationHistory() (datatypes.Network_Application_Delivery_Controller_Configuration_History, error)
	GetDatacenter() (datatypes.Location, error)
	GetDescription() (string, error)
	GetLicenseExpirationDate() (time.Time, error)
	GetLoadBalancers() (datatypes.Network_LoadBalancer_VirtualIpAddress, error)
	GetManagedResourceFlag() (bool, error)
	GetManagementIpAddress() (string, error)
	GetNetworkVlan() (datatypes.Network_Vlan, error)
	GetNetworkVlans() (datatypes.Network_Vlan, error)
	GetOutboundPublicBandwidthUsage() (float64, error)
	GetPassword() (datatypes.Software_Component_Password, error)
	GetPrimaryIpAddress() (string, error)
	GetProjectedPublicBandwidthUsage() (float64, error)
	GetSubnets() (datatypes.Network_Subnet, error)
	GetTagReferences() (datatypes.Tag_Reference, error)
	GetType() (datatypes.Network_Application_Delivery_Controller_Type, error)
	GetVirtualIpAddresses() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress, error)
}

type Network_Application_Delivery_Controller_Configuration_History interface {
	Entity

	DeleteObject() (bool, error)
	GetObject() (datatypes.Network_Application_Delivery_Controller_Configuration_History, error)

	GetController() (datatypes.Network_Application_Delivery_Controller, error)
}

type Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute interface {
	Entity

	GetObject() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute, error)

	GetHealthCheck() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Check, error)
	GetType() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type, error)
}

type Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type interface {
	Entity

	GetAllObjects() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type, error)
	GetObject() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type, error)
}

type Network_Application_Delivery_Controller_LoadBalancer_Health_Check interface {
	Entity

	GetObject() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Check, error)

	GetAttributes() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute, error)
	GetScaleLoadBalancers() (datatypes.Scale_LoadBalancer, error)
	GetServices() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Service, error)
	GetType() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type, error)
}

type Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type interface {
	Entity

	GetAllObjects() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type, error)
	GetObject() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type, error)
}

type Network_Application_Delivery_Controller_LoadBalancer_Routing_Method interface {
	Entity

	GetAllObjects() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Routing_Method, error)
	GetObject() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Routing_Method, error)
}

type Network_Application_Delivery_Controller_LoadBalancer_Routing_Type interface {
	Entity

	GetAllObjects() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Routing_Type, error)
	GetObject() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Routing_Type, error)
}

type Network_Application_Delivery_Controller_LoadBalancer_Service interface {
	Entity

	DeleteObject() (bool, error)
	GetGraphImage(graphType string, metric string) ([]byte, error)
	GetObject() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Service, error)
	ToggleStatus() (bool, error)

	GetGroupReferences() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference, error)
	GetGroups() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Service_Group, error)
	GetHealthCheck() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Check, error)
	GetHealthChecks() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Check, error)
	GetIpAddress() (datatypes.Network_Subnet_IpAddress, error)
	GetServiceGroup() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Service_Group, error)
}

type Network_Application_Delivery_Controller_LoadBalancer_Service_Group interface {
	Entity

	GetGraphImage(graphType string, metric string) ([]byte, error)
	GetObject() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Service_Group, error)
	KickAllConnections() (bool, error)

	GetRoutingMethod() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Routing_Method, error)
	GetRoutingType() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Routing_Type, error)
	GetServiceReferences() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference, error)
	GetServices() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Service, error)
	GetVirtualServer() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualServer, error)
	GetVirtualServers() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualServer, error)
}

type Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference interface {
	Entity

	GetService() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Service, error)
	GetServiceGroup() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Service_Group, error)
}

type Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress interface {
	Entity

	EditObject(templateObject datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) (bool, error)
	GetAvailableSecureTransportCiphers() (datatypes.Security_SecureTransportCipher, error)
	GetAvailableSecureTransportProtocols() (datatypes.Security_SecureTransportProtocol, error)
	GetObject() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress, error)
	StartSsl() (bool, error)
	StopSsl() (bool, error)
	UpgradeConnectionLimit() (bool, error)

	GetAccount() (datatypes.Account, error)
	GetApplicationDeliveryController() (datatypes.Network_Application_Delivery_Controller, error)
	GetApplicationDeliveryControllers() (datatypes.Network_Application_Delivery_Controller, error)
	GetBillingItem() (datatypes.Billing_Item, error)
	GetDedicatedBillingItem() (datatypes.Billing_Item_Network_LoadBalancer, error)
	GetHighAvailabilityFlag() (bool, error)
	GetIpAddress() (datatypes.Network_Subnet_IpAddress, error)
	GetLoadBalancerHardware() (datatypes.Hardware, error)
	GetManagedResourceFlag() (bool, error)
	GetSecureTransportCiphers() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportCipher, error)
	GetSecureTransportProtocols() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportProtocol, error)
	GetSecurityCertificate() (datatypes.Security_Certificate, error)
	GetSecurityCertificateEntry() (datatypes.Security_Certificate_Entry, error)
	GetVirtualServers() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualServer, error)
}

type Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportCipher interface {
	Entity

	GetVirtualIpAddress() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress, error)
}

type Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportProtocol interface {
	Entity

	GetVirtualIpAddress() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress, error)
}

type Network_Application_Delivery_Controller_LoadBalancer_VirtualServer interface {
	Entity

	DeleteObject() (bool, error)
	GetObject() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualServer, error)
	StartSsl() (bool, error)
	StopSsl() (bool, error)

	GetRoutingMethod() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Routing_Method, error)
	GetScaleLoadBalancers() (datatypes.Scale_LoadBalancer, error)
	GetServiceGroups() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Service_Group, error)
	GetVirtualIpAddress() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress, error)
}

type Network_Application_Delivery_Controller_Type interface {
	Entity
}

type Network_Backbone interface {
	Entity

	GetAllBackbones() (datatypes.Network_Backbone, error)
	GetBackbonesForLocationName(locationName string) (datatypes.Network_Backbone, error)
	GetGraphImage() ([]byte, error)
	GetObject() (datatypes.Network_Backbone, error)

	GetHealth() (string, error)
	GetLocation() (datatypes.Location, error)
	GetNetworkComponent() (datatypes.Network_Component, error)
}

type Network_Backbone_Location_Dependent interface {
	Entity

	GetAllObjects() (datatypes.Network_Backbone_Location_Dependent, error)
	GetObject() (datatypes.Network_Backbone_Location_Dependent, error)
	GetSourceDependentsByName(locationName string) (datatypes.Location, error)

	GetDependentLocation() (datatypes.Location, error)
	GetSourceLocation() (datatypes.Location, error)
}

type Network_Bandwidth_Usage interface {
	Entity

	GetTrackingObject() (datatypes.Metric_Tracking_Object, error)
	GetType() (datatypes.Network_Bandwidth_Version1_Usage_Detail_Type, error)
}

type Network_Bandwidth_Usage_Detail interface {
	Entity

	GetAccount() (datatypes.Account, error)
	GetTrackingObject() (datatypes.Metric_Tracking_Object, error)
	GetType() (datatypes.Network_Bandwidth_Version1_Usage_Detail_Type, error)
}

type Network_Bandwidth_Version1_Allocation interface {
	Entity

	GetAllotmentDetail() (datatypes.Network_Bandwidth_Version1_Allotment_Detail, error)
	GetBillingItem() (datatypes.Billing_Item_Hardware, error)
}

type Network_Bandwidth_Version1_Allotment interface {
	Entity

	CreateObject(templateObject datatypes.Network_Bandwidth_Version1_Allotment) (datatypes.Network_Bandwidth_Version1_Allotment, error)
	EditObject(templateObject datatypes.Network_Bandwidth_Version1_Allotment) (bool, error)
	GetBackendBandwidthByHour(date time.Time) (datatypes.Container_Network_Bandwidth_Version1_Usage, error)
	GetBackendBandwidthUse(startDate time.Time, endDate time.Time) (datatypes.Network_Bandwidth_Version1_Usage_Detail, error)
	GetBandwidthForDateRange(startDate time.Time, endDate time.Time) (datatypes.Metric_Tracking_Object_Data, error)
	GetBandwidthImage(networkType string, snapshotRange string, draw bool, dateSpecified time.Time, dateSpecifiedEnd time.Time) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetCustomBandwidthDataByDate(graphData datatypes.Container_Graph) (datatypes.Container_Graph, error)
	GetFrontendBandwidthByHour(date time.Time) (datatypes.Container_Network_Bandwidth_Version1_Usage, error)
	GetFrontendBandwidthUse(startDate time.Time, endDate time.Time) (datatypes.Network_Bandwidth_Version1_Usage_Detail, error)
	GetObject() (datatypes.Network_Bandwidth_Version1_Allotment, error)
	ReassignServers(templateObjects datatypes.Hardware, newAllotmentId int) (bool, error)
	RequestVdrCancellation() (bool, error)
	RequestVdrContentUpdates(hardwareToAdd datatypes.Hardware, hardwareToRemove datatypes.Hardware, cloudsToAdd datatypes.Virtual_Guest, cloudsToRemove datatypes.Virtual_Guest, optionalAllotmentId int, adcToAdd datatypes.Network_Application_Delivery_Controller, adcToRemove datatypes.Network_Application_Delivery_Controller) (bool, error)
	SetVdrContent(hardware datatypes.Hardware, bareMetalServers datatypes.Hardware, virtualServerInstance datatypes.Virtual_Guest, adc datatypes.Network_Application_Delivery_Controller, optionalAllotmentId int) (bool, error)
	UnassignServers(templateObjects datatypes.Hardware) (bool, error)
	VoidPendingServerMove(id int, typ string) (bool, error)
	VoidPendingVdrCancellation() (bool, error)

	GetAccount() (datatypes.Account, error)
	GetActiveDetails() (datatypes.Network_Bandwidth_Version1_Allotment_Detail, error)
	GetApplicationDeliveryControllers() (datatypes.Network_Application_Delivery_Controller, error)
	GetAverageDailyPublicBandwidthUsage() (float64, error)
	GetBareMetalInstances() (datatypes.Hardware, error)
	GetBillingCycleBandwidthUsage() (datatypes.Network_Bandwidth_Usage, error)
	GetBillingCyclePrivateBandwidthUsage() (datatypes.Network_Bandwidth_Usage, error)
	GetBillingCyclePublicBandwidthUsage() (datatypes.Network_Bandwidth_Usage, error)
	GetBillingCyclePublicUsageTotal() (uint, error)
	GetBillingItem() (datatypes.Billing_Item, error)
	GetCurrentBandwidthSummary() (datatypes.Metric_Tracking_Object_Bandwidth_Summary, error)
	GetDetails() (datatypes.Network_Bandwidth_Version1_Allotment_Detail, error)
	GetHardware() (datatypes.Hardware, error)
	GetInboundPublicBandwidthUsage() (float64, error)
	GetLocationGroup() (datatypes.Location_Group, error)
	GetManagedBareMetalInstances() (datatypes.Hardware, error)
	GetManagedHardware() (datatypes.Hardware, error)
	GetManagedVirtualGuests() (datatypes.Virtual_Guest, error)
	GetMetricTrackingObject() (datatypes.Metric_Tracking_Object_VirtualDedicatedRack, error)
	GetMetricTrackingObjectId() (int, error)
	GetOutboundPublicBandwidthUsage() (float64, error)
	GetOverBandwidthAllocationFlag() (int, error)
	GetPrivateNetworkOnlyHardware() (datatypes.Hardware, error)
	GetProjectedOverBandwidthAllocationFlag() (int, error)
	GetProjectedPublicBandwidthUsage() (float64, error)
	GetServiceProvider() (datatypes.Service_Provider, error)
	GetTotalBandwidthAllocated() (uint, error)
	GetVirtualGuests() (datatypes.Virtual_Guest, error)
}

type Network_Bandwidth_Version1_Allotment_Detail interface {
	Entity

	GetAllocation() (datatypes.Network_Bandwidth_Version1_Allocation, error)
	GetBandwidthAllotment() (datatypes.Network_Bandwidth_Version1_Allotment, error)
	GetBandwidthUsage() (datatypes.Network_Bandwidth_Version1_Usage, error)
}

type Network_Bandwidth_Version1_Host interface {
	Entity
}

type Network_Bandwidth_Version1_Interface interface {
	Entity

	GetHost() (datatypes.Network_Bandwidth_Version1_Host, error)
	GetNetworkComponent() (datatypes.Network_Component, error)
}

type Network_Bandwidth_Version1_Usage interface {
	Entity

	GetBandwidthAllotmentDetail() (datatypes.Network_Bandwidth_Version1_Allotment_Detail, error)
	GetBandwidthUsageDetail() (datatypes.Network_Bandwidth_Version1_Usage_Detail, error)
}

type Network_Bandwidth_Version1_Usage_Detail interface {
	Entity

	GetBandwidthUsage() (datatypes.Network_Bandwidth_Version1_Usage, error)
	GetBandwidthUsageDetailType() (datatypes.Network_Bandwidth_Version1_Usage_Detail_Type, error)
}

type Network_Bandwidth_Version1_Usage_Detail_Total interface {
	Entity

	GetAccount() (datatypes.Account, error)
	GetTrackingObject() (datatypes.Metric_Tracking_Object, error)
	GetType() (datatypes.Network_Bandwidth_Version1_Usage_Detail_Type, error)
}

type Network_Bandwidth_Version1_Usage_Detail_Type interface {
	Entity
}

type Network_Component interface {
	Entity

	AddNetworkVlanTrunks(networkVlans datatypes.Network_Vlan) (datatypes.Network_Vlan, error)
	ClearNetworkVlanTrunks() (datatypes.Network_Vlan, error)
	GetCustomBandwidthDataByDate(graphData datatypes.Container_Graph) (datatypes.Container_Graph, error)
	GetObject() (datatypes.Network_Component, error)
	GetPortStatistics() (datatypes.Container_Network_Port_Statistic, error)
	RemoveNetworkVlanTrunks(networkVlans datatypes.Network_Vlan) (datatypes.Network_Vlan, error)

	GetActiveCommand() (datatypes.Hardware_Component_RemoteManagement_Command_Request, error)
	GetDownlinkComponent() (datatypes.Network_Component, error)
	GetDuplexMode() (datatypes.Network_Component_Duplex_Mode, error)
	GetHardware() (datatypes.Hardware, error)
	GetHighAvailabilityFirewallFlag() (bool, error)
	GetInterface() (datatypes.Network_Bandwidth_Version1_Interface, error)
	GetIpAddressBindings() (datatypes.Network_Component_IpAddress, error)
	GetIpAddresses() (datatypes.Network_Subnet_IpAddress, error)
	GetLastCommand() (datatypes.Hardware_Component_RemoteManagement_Command_Request, error)
	GetMetricTrackingObject() (datatypes.Metric_Tracking_Object, error)
	GetNetworkComponentFirewall() (datatypes.Network_Component_Firewall, error)
	GetNetworkComponentGroup() (datatypes.Network_Component_Group, error)
	GetNetworkHardware() (datatypes.Hardware, error)
	GetNetworkVlan() (datatypes.Network_Vlan, error)
	GetNetworkVlanTrunks() (datatypes.Network_Component_Network_Vlan_Trunk, error)
	GetPrimaryIpAddressRecord() (datatypes.Network_Subnet_IpAddress, error)
	GetPrimarySubnet() (datatypes.Network_Subnet, error)
	GetPrimaryVersion6IpAddressRecord() (datatypes.Network_Subnet_IpAddress, error)
	GetRecentCommands() (datatypes.Hardware_Component_RemoteManagement_Command_Request, error)
	GetRedundancyCapableFlag() (bool, error)
	GetRedundancyEnabledFlag() (bool, error)
	GetRemoteManagementUsers() (datatypes.Hardware_Component_RemoteManagement_User, error)
	GetRouter() (datatypes.Hardware, error)
	GetStorageNetworkFlag() (bool, error)
	GetSubnets() (datatypes.Network_Subnet, error)
	GetUplinkComponent() (datatypes.Network_Component, error)
	GetUplinkDuplexMode() (datatypes.Network_Component_Duplex_Mode, error)
}

type Network_Component_Duplex_Mode interface {
	Entity
}

type Network_Component_Firewall interface {
	Entity

	GetObject() (datatypes.Network_Component_Firewall, error)

	GetApplyServerRuleSubnets() (datatypes.Network_Subnet, error)
	GetBillingItem() (datatypes.Billing_Item, error)
	GetGuestNetworkComponent() (datatypes.Virtual_Guest_Network_Component, error)
	GetNetworkComponent() (datatypes.Network_Component, error)
	GetNetworkFirewallUpdateRequest() (datatypes.Network_Firewall_Update_Request, error)
	GetRules() (datatypes.Network_Component_Firewall_Rule, error)
	GetSubnets() (datatypes.Network_Subnet, error)
}

type Network_Component_Firewall_Rule interface {
	Entity

	GetNetworkComponentFirewall() (datatypes.Network_Component_Firewall, error)
}

type Network_Component_Firewall_Subnets interface {
	Entity

	GetNetworkComponentFirewall() (datatypes.Network_Component_Firewall, error)
	GetSubnet() (datatypes.Network_Subnet, error)
}

type Network_Component_Group interface {
	Entity

	GetNetworkComponents() (datatypes.Network_Component, error)
}

type Network_Component_IpAddress interface {
	Entity

	GetIpAddress() (datatypes.Network_Subnet_IpAddress, error)
	GetNetworkComponent() (datatypes.Network_Component, error)
}

type Network_Component_Network_Vlan_Trunk interface {
	Entity

	GetNetworkComponent() (datatypes.Network_Component, error)
	GetNetworkVlan() (datatypes.Network_Vlan, error)
}

type Network_Component_RemoteManagement interface {
	Network_Component
}

type Network_Component_Uplink_Hardware interface {
	Entity

	GetHardware() (datatypes.Hardware, error)
	GetNetworkComponent() (datatypes.Network_Component, error)
}

type Network_ContentDelivery_Account interface {
	Entity

	AuthenticateResourceRequest(parameter datatypes.Container_Network_ContentDelivery_Authentication_Parameter) (bool, error)
	CreateDirectory(directoryName string) (bool, error)
	CreateFtpUser(newPassword string) (bool, error)
	CreateOriginPullMapping(mappingObject datatypes.Container_Network_ContentDelivery_OriginPull_Mapping) (bool, error)
	CreateOriginPullRule(originDomain string, cnameRecord string) (bool, error)
	CreateTokenAuthenticationDirectory(directory string, mediaType string) (bool, error)
	DeleteFtpUser() (bool, error)
	DeleteOriginPullRule(originMappingId string) (bool, error)
	DisableLogging() (bool, error)
	EnableLogging() (bool, error)
	GetAllPopsBandwidthData(beginDateTime time.Time, endDateTime time.Time) (datatypes.Container_Network_ContentDelivery_Bandwidth_PointsOfPresence_Summary, error)
	GetAllPopsBandwidthImage(title string, beginDateTime time.Time, endDateTime time.Time, unit string) (datatypes.Container_Bandwidth_GraphOutputsExtended, error)
	GetAuthenticationServiceEndpoints() (datatypes.Container_Network_ContentDelivery_Authentication_ServiceEndpoint, error)
	GetBandwidthData(beginDateTime time.Time, endDateTime time.Time) (datatypes.Container_Network_ContentDelivery_Bandwidth_Summary, error)
	GetBandwidthDataWithTypes(beginDateTime time.Time, endDateTime time.Time, period string) (datatypes.Container_Network_ContentDelivery_Report_Usage, error)
	GetBandwidthImage(title string, beginDateTime time.Time, endDateTime time.Time) (datatypes.Container_Bandwidth_GraphOutputsExtended, error)
	GetCustomerOrigins(mediaType string) (datatypes.Container_Network_ContentDelivery_OriginPull_Mapping, error)
	GetDirectoryInformation(directoryName string) (datatypes.Container_Network_Directory_Listing, error)
	GetDiskSpaceUsageDataByDate(beginDateTime time.Time, endDateTime time.Time) (datatypes.Metric_Tracking_Object_Data, error)
	GetDiskSpaceUsageImageByDate(beginDateTime time.Time, endDateTime time.Time) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetFtpAttributes() (datatypes.Container_Network_Authentication_Data, error)
	GetMediaUrls() (datatypes.Container_Network_ContentDelivery_SupportedProtocol, error)
	GetObject() (datatypes.Network_ContentDelivery_Account, error)
	GetOriginPullMappingInformation() (datatypes.Container_Network_ContentDelivery_OriginPull_Mapping, error)
	GetOriginPullSupportedMediaUrls() (datatypes.Container_Network_ContentDelivery_SupportedProtocol, error)
	GetOriginPullUrl() (string, error)
	GetPopNames() (datatypes.Container_Network_ContentDelivery_PointsOfPresence, error)
	GetProviderPortalCredentials() (datatypes.Container_Network_Authentication_Data, error)
	GetTokenAuthenticationDirectories() (datatypes.Container_Network_Directory_Listing, error)
	GetVendorFtpAttributes() (datatypes.Container_Network_Authentication_Data, error)
	LoadContent(objectUrls string) (bool, error)
	ManageHttpCompression(enableFlag bool, mimeTypes string) (bool, error)
	PurgeCache(objectUrls string) (datatypes.Container_Network_ContentDelivery_PurgeService_Response, error)
	RemoveAuthenticationDirectory(directory string, mediaType string) (bool, error)
	RemoveFile(source string) (bool, error)
	SetAuthenticationServiceEndpoint(webserviceEndpoint string, protocol string) (bool, error)
	SetFtpPassword(newPassword string) (bool, error)
	UpdateNote(note string) (bool, error)
	UploadStream(source datatypes.Container_Utility_File_Attachment, target string) (bool, error)

	GetAccount() (datatypes.Account, error)
	GetAssociatedCdnAccountId() (string, error)
	GetAuthenticationIpAddresses() (datatypes.Network_ContentDelivery_Authentication_Address, error)
	GetBillingItem() (datatypes.Billing_Item, error)
	GetCdnAccountName() (string, error)
	GetCdnAccountNote() (string, error)
	GetCdnSolutionName() (string, error)
	GetDependantServiceFlag() (bool, error)
	GetLegacyCdnFlag() (bool, error)
	GetLogEnabledFlag() (string, error)
	GetProviderPortalAccessFlag() (bool, error)
	GetStatus() (datatypes.Network_ContentDelivery_Account_Status, error)
	GetTokenAuthenticationEnabledFlag() (bool, error)
}

type Network_ContentDelivery_Account_Status interface {
	Entity
}

type Network_ContentDelivery_Authentication_Address interface {
	Entity

	CreateObject(templateObject datatypes.Network_ContentDelivery_Authentication_Address) (datatypes.Network_ContentDelivery_Authentication_Address, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Network_ContentDelivery_Authentication_Address) (bool, error)
	GetObject() (datatypes.Network_ContentDelivery_Authentication_Address, error)
	RearrangeAuthenticationIp(cdnAccountId int, templateObjects datatypes.Network_ContentDelivery_Authentication_Address) (bool, error)
}

type Network_ContentDelivery_Authentication_Token interface {
	Entity

	CreateObject(templateObject datatypes.Network_ContentDelivery_Authentication_Token) (datatypes.Network_ContentDelivery_Authentication_Token, error)
	GetAllManagedTokens(cdnAccountId int) (datatypes.Network_ContentDelivery_Authentication_Token, error)
	GetObject() (datatypes.Network_ContentDelivery_Authentication_Token, error)
	GetTimedToken(cdnAccountId int, tokenLife int, clientIp string, referrer string, mediaType string) (string, error)
	RevokeAllManagedTokens(cdnAccountId int) (bool, error)
	RevokeAllTokens(cdnAccountId int, mediaType string) (bool, error)
	RevokeManagedToken(cdnAccountId int, token string) (bool, error)
	RevokeManagedTokens(templateObjects datatypes.Network_ContentDelivery_Authentication_Token) (bool, error)
}

type Network_Customer_Subnet interface {
	Entity

	CreateObject(templateObject datatypes.Network_Customer_Subnet) (datatypes.Network_Customer_Subnet, error)
	GetObject() (datatypes.Network_Customer_Subnet, error)

	GetIpAddresses() (datatypes.Network_Customer_Subnet_IpAddress, error)
}

type Network_Customer_Subnet_IpAddress interface {
	Entity

	GetSubnet() (datatypes.Network_Customer_Subnet, error)
	GetTranslations() (datatypes.Network_Tunnel_Module_Context_Address_Translation, error)
}

type Network_Firewall_AccessControlList interface {
	Entity

	GetObject() (datatypes.Network_Firewall_AccessControlList, error)

	GetNetworkFirewallUpdateRequests() (datatypes.Network_Firewall_Update_Request, error)
	GetNetworkVlan() (datatypes.Network_Vlan, error)
	GetRules() (datatypes.Network_Vlan_Firewall_Rule, error)
}

type Network_Firewall_Interface interface {
	Network_Firewall_Module_Context_Interface

	GetObject() (datatypes.Network_Firewall_Interface, error)
}

type Network_Firewall_Module_Context_Interface interface {
	Entity

	GetObject() (datatypes.Network_Firewall_Module_Context_Interface, error)

	GetFirewallContextAccessControlLists() (datatypes.Network_Firewall_AccessControlList, error)
	GetNetworkVlan() (datatypes.Network_Vlan, error)
}

type Network_Firewall_Template interface {
	Entity

	GetAllObjects() (datatypes.Network_Firewall_Template, error)
	GetObject() (datatypes.Network_Firewall_Template, error)

	GetRules() (datatypes.Network_Firewall_Template_Rule, error)
}

type Network_Firewall_Template_Rule interface {
	Entity

	GetFirewallTemplate() (datatypes.Network_Firewall_Template, error)
}

type Network_Firewall_Update_Request interface {
	Entity

	CreateObject(templateObject datatypes.Network_Firewall_Update_Request) (datatypes.Network_Firewall_Update_Request, error)
	GetFirewallUpdateRequestRuleAttributes() (datatypes.Container_Utility_Network_Firewall_Rule_Attribute, error)
	GetObject() (datatypes.Network_Firewall_Update_Request, error)
	UpdateRuleNote(fwRule datatypes.Network_Component_Firewall_Rule, note string) (bool, error)

	GetAuthorizingUser() (datatypes.User_Interface, error)
	GetGuest() (datatypes.Virtual_Guest, error)
	GetHardware() (datatypes.Hardware, error)
	GetNetworkComponentFirewall() (datatypes.Network_Component_Firewall, error)
	GetRules() (datatypes.Network_Firewall_Update_Request_Rule, error)
}

type Network_Firewall_Update_Request_Customer interface {
	Network_Firewall_Update_Request
}

type Network_Firewall_Update_Request_Employee interface {
	Network_Firewall_Update_Request
}

type Network_Firewall_Update_Request_Rule interface {
	Entity

	CreateObject(templateObject datatypes.Network_Firewall_Update_Request_Rule) (datatypes.Network_Firewall_Update_Request_Rule, error)
	GetObject() (datatypes.Network_Firewall_Update_Request_Rule, error)
	ValidateRule(rule datatypes.Network_Firewall_Update_Request_Rule, applyToComponentId int, applyToAclId int) error

	GetFirewallUpdateRequest() (datatypes.Network_Firewall_Update_Request, error)
}

type Network_Firewall_Update_Request_Rule_Version6 interface {
	Network_Firewall_Update_Request_Rule
}

type Network_Gateway interface {
	Entity

	BypassAllVlans() error
	BypassVlans(vlans datatypes.Network_Gateway_Vlan) error
	CreateObject(templateObject datatypes.Network_Gateway) (datatypes.Network_Gateway, error)
	EditObject(templateObject datatypes.Network_Gateway) (bool, error)
	GetObject() (datatypes.Network_Gateway, error)
	GetPossibleInsideVlans() (datatypes.Network_Vlan, error)
	UnbypassAllVlans() error
	UnbypassVlans(vlans datatypes.Network_Gateway_Vlan) error

	GetAccount() (datatypes.Account, error)
	GetInsideVlans() (datatypes.Network_Gateway_Vlan, error)
	GetMembers() (datatypes.Network_Gateway_Member, error)
	GetPrivateIpAddress() (datatypes.Network_Subnet_IpAddress, error)
	GetPrivateVlan() (datatypes.Network_Vlan, error)
	GetPublicIpAddress() (datatypes.Network_Subnet_IpAddress, error)
	GetPublicIpv6Address() (datatypes.Network_Subnet_IpAddress, error)
	GetPublicVlan() (datatypes.Network_Vlan, error)
	GetStatus() (datatypes.Network_Gateway_Status, error)
}

type Network_Gateway_Member interface {
	Entity

	CreateObject(templateObject datatypes.Network_Gateway_Member) (datatypes.Network_Gateway_Member, error)
	CreateObjects(templateObjects datatypes.Network_Gateway_Member) (datatypes.Network_Gateway_Member, error)
	GetObject() (datatypes.Network_Gateway_Member, error)

	GetHardware() (datatypes.Hardware, error)
	GetNetworkGateway() (datatypes.Network_Gateway, error)
}

type Network_Gateway_Status interface {
	Entity

	GetObject() (datatypes.Network_Gateway_Status, error)
}

type Network_Gateway_Vlan interface {
	Entity

	Bypass() error
	CreateObject(templateObject datatypes.Network_Gateway_Vlan) (datatypes.Network_Gateway_Vlan, error)
	CreateObjects(templateObjects datatypes.Network_Gateway_Vlan) (datatypes.Network_Gateway_Vlan, error)
	DeleteObject() error
	DeleteObjects(templateObjects datatypes.Network_Gateway_Vlan) (bool, error)
	GetObject() (datatypes.Network_Gateway_Vlan, error)
	Unbypass() error

	GetNetworkGateway() (datatypes.Network_Gateway, error)
	GetNetworkVlan() (datatypes.Network_Vlan, error)
}

type Network_LoadBalancer_Global_Account interface {
	Entity

	AddNsRecord() (bool, error)
	EditObject(templateObject datatypes.Network_LoadBalancer_Global_Account) (bool, error)
	GetObject() (datatypes.Network_LoadBalancer_Global_Account, error)
	RemoveNsRecord() (bool, error)

	GetAccount() (datatypes.Account, error)
	GetBillingItem() (datatypes.Billing_Item, error)
	GetHosts() (datatypes.Network_LoadBalancer_Global_Host, error)
	GetLoadBalanceType() (datatypes.Network_LoadBalancer_Global_Type, error)
	GetManagedResourceFlag() (bool, error)
}

type Network_LoadBalancer_Global_Host interface {
	Entity

	DeleteObject() (bool, error)
	GetObject() (datatypes.Network_LoadBalancer_Global_Host, error)

	GetLoadBalancerAccount() (datatypes.Network_LoadBalancer_Global_Account, error)
}

type Network_LoadBalancer_Global_Type interface {
	Entity
}

type Network_LoadBalancer_Service interface {
	Entity

	DeleteObject() (bool, error)
	GetGraphImage(graphType string, metric string) ([]byte, error)
	GetObject() (datatypes.Network_LoadBalancer_Service, error)
	GetStatus() (datatypes.Container_Network_LoadBalancer_StatusEntry, error)
	ResetPeakConnections() (bool, error)

	GetVip() (datatypes.Network_LoadBalancer_VirtualIpAddress, error)
}

type Network_LoadBalancer_VirtualIpAddress interface {
	Entity

	Disable() (bool, error)
	EditObject(templateObject datatypes.Network_LoadBalancer_VirtualIpAddress) (bool, error)
	Enable() (bool, error)
	GetObject() (datatypes.Network_LoadBalancer_VirtualIpAddress, error)
	KickAllConnections() (bool, error)
	UpgradeConnectionLimit() (bool, error)

	GetAccount() (datatypes.Account, error)
	GetBillingItem() (datatypes.Billing_Item, error)
	GetCustomerManagedFlag() (int, error)
	GetManagedResourceFlag() (bool, error)
	GetServices() (datatypes.Network_LoadBalancer_Service, error)
}

type Network_Logging_Syslog interface {
	Entity
}

type Network_Media_Transcode_Account interface {
	Entity

	CreateTranscodeAccount() (bool, error)
	CreateTranscodeJob(newJob datatypes.Network_Media_Transcode_Job) (bool, error)
	GetDirectoryInformation(directoryName string, extensionFilter string) (datatypes.Container_Network_Directory_Listing, error)
	GetFileDetail(source string) (datatypes.Container_Network_Media_Information, error)
	GetFtpAttributes() (datatypes.Container_Network_Authentication_Data, error)
	GetObject() (datatypes.Network_Media_Transcode_Account, error)
	GetPresetDetail(guid string) (datatypes.Container_Network_Media_Transcode_Preset_Element, error)
	GetPresets() (datatypes.Container_Network_Media_Transcode_Preset, error)

	GetAccount() (datatypes.Account, error)
	GetTranscodeJobs() (datatypes.Network_Media_Transcode_Job, error)
}

type Network_Media_Transcode_Job interface {
	Entity

	CreateObject(templateObject datatypes.Network_Media_Transcode_Job) (datatypes.Network_Media_Transcode_Job, error)
	GetObject() (datatypes.Network_Media_Transcode_Job, error)

	GetHistory() (datatypes.Network_Media_Transcode_Job_History, error)
	GetTranscodeAccount() (datatypes.Network_Media_Transcode_Account, error)
	GetTranscodeStatus() (datatypes.Network_Media_Transcode_Job_Status, error)
	GetTranscodeStatusName() (string, error)
	GetUser() (datatypes.User_Customer, error)
}

type Network_Media_Transcode_Job_History interface {
	Entity

	GetTranscodeStatusName() (string, error)
}

type Network_Media_Transcode_Job_Status interface {
	Entity

	GetAllStatuses() (datatypes.Network_Media_Transcode_Job_Status, error)
	GetObject() (datatypes.Network_Media_Transcode_Job_Status, error)
}

type Network_Message_Delivery interface {
	Entity

	EditObject(templateObject datatypes.Network_Message_Delivery) (bool, error)
	GetObject() (datatypes.Network_Message_Delivery, error)

	GetAccount() (datatypes.Account, error)
	GetBillingItem() (datatypes.Billing_Item, error)
	GetType() (datatypes.Network_Message_Delivery_Type, error)
	GetVendor() (datatypes.Network_Message_Delivery_Vendor, error)
}

type Network_Message_Delivery_Attribute interface {
	Entity

	GetNetworkMessageDelivery() (datatypes.Network_Message_Delivery, error)
}

type Network_Message_Delivery_Email_Sendgrid interface {
	Network_Message_Delivery

	AddUnsubscribeEmailAddress(emailAddress string) (bool, error)
	DeleteEmailListEntries(list string, entries string) (bool, error)
	DisableSmtpAccess() (bool, error)
	EnableSmtpAccess() (bool, error)
	GetAccountOverview() (datatypes.Container_Network_Message_Delivery_Email_Sendgrid_Account_Overview, error)
	GetCategoryList() (string, error)
	GetEmailList(list string) (datatypes.Container_Network_Message_Delivery_Email_Sendgrid_List_Entry, error)
	GetObject() (datatypes.Network_Message_Delivery_Email_Sendgrid, error)
	GetStatistics(options datatypes.Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Options) (datatypes.Container_Network_Message_Delivery_Email_Sendgrid_Statistics, error)
	GetStatisticsGraph(options datatypes.Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Options) (datatypes.Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Graph, error)
	GetVendorPortalUrl() (string, error)
	SendEmail(emailContainer datatypes.Container_Network_Message_Delivery_Email) (bool, error)
	UpdateEmailAddress(emailAddress string) (bool, error)

	GetEmailAddress() (string, error)
	GetSmtpAccess() (string, error)
}

type Network_Message_Delivery_Type interface {
	Entity
}

type Network_Message_Delivery_Vendor interface {
	Entity
}

type Network_Message_Queue interface {
	Entity

	GetObject() (datatypes.Network_Message_Queue, error)

	GetAccount() (datatypes.Account, error)
	GetBillingItem() (datatypes.Billing_Item, error)
	GetNodes() (datatypes.Network_Message_Queue_Node, error)
	GetStatus() (datatypes.Network_Message_Queue_Status, error)
}

type Network_Message_Queue_Node interface {
	Entity

	AddUser(username string) (bool, error)
	DeleteUser(username string) (bool, error)
	GetAllUsers() (string, error)
	GetObject() (datatypes.Network_Message_Queue_Node, error)
	GetUsage(startDate time.Time, endDate time.Time) (datatypes.Metric_Tracking_Object_Data, error)
	GetUsageGraph(graphData datatypes.Container_Graph) (datatypes.Container_Graph, error)

	GetMessageQueue() (datatypes.Network_Message_Queue, error)
	GetMetricTrackingObject() (datatypes.Metric_Tracking_Object, error)
	GetServiceResource() (datatypes.Network_Service_Resource, error)
}

type Network_Message_Queue_Status interface {
	Entity

	GetObject() (datatypes.Network_Message_Queue_Status, error)
}

type Network_Monitor interface {
	Entity

	GetIpAddressesByHardware(hardware datatypes.Hardware, partialIpAddress string) (datatypes.Network_Subnet_IpAddress, error)
	GetIpAddressesByVirtualGuest(guest datatypes.Virtual_Guest, partialIpAddress string) (datatypes.Network_Subnet_IpAddress, error)
}

type Network_Monitor_Version1_Incident interface {
	Entity
}

type Network_Monitor_Version1_Query_Host interface {
	Entity

	CreateObject(templateObject datatypes.Network_Monitor_Version1_Query_Host) (datatypes.Network_Monitor_Version1_Query_Host, error)
	CreateObjects(templateObjects datatypes.Network_Monitor_Version1_Query_Host) (datatypes.Network_Monitor_Version1_Query_Host, error)
	DeleteObject() (bool, error)
	DeleteObjects(templateObjects datatypes.Network_Monitor_Version1_Query_Host) (bool, error)
	EditObject(templateObject datatypes.Network_Monitor_Version1_Query_Host) (bool, error)
	EditObjects(templateObjects datatypes.Network_Monitor_Version1_Query_Host) (bool, error)
	FindByHardwareId(hardwareId int) (datatypes.Network_Monitor_Version1_Query_Host, error)
	GetObject() (datatypes.Network_Monitor_Version1_Query_Host, error)

	GetHardware() (datatypes.Hardware, error)
	GetLastResult() (datatypes.Network_Monitor_Version1_Query_Result, error)
	GetQueryType() (datatypes.Network_Monitor_Version1_Query_Type, error)
	GetResponseAction() (datatypes.Network_Monitor_Version1_Query_ResponseType, error)
}

type Network_Monitor_Version1_Query_Host_Stratum interface {
	Entity

	GetAllQueryTypes() (datatypes.Network_Monitor_Version1_Query_Type, error)
	GetAllResponseTypes() (datatypes.Network_Monitor_Version1_Query_ResponseType, error)
	GetObject() (datatypes.Network_Monitor_Version1_Query_Host_Stratum, error)

	GetHardware() (datatypes.Hardware, error)
}

type Network_Monitor_Version1_Query_ResponseType interface {
	Entity
}

type Network_Monitor_Version1_Query_Result interface {
	Entity

	GetQueryHost() (datatypes.Network_Monitor_Version1_Query_Host, error)
}

type Network_Monitor_Version1_Query_Type interface {
	Entity
}

type Network_Pod interface {
	Entity

	GetAllObjects() (datatypes.Network_Pod, error)
	GetCapabilities() (string, error)
	GetObject() (datatypes.Network_Pod, error)
	ListCapabilities() (string, error)
}

type Network_Protection_Address interface {
	Entity

	GetAccount() (datatypes.Account, error)
	GetLocation() (datatypes.Location, error)
	GetModifiedUser() (datatypes.User_Employee, error)
	GetPrimaryRouter() (datatypes.Hardware_Router, error)
	GetServiceProvider() (datatypes.Service_Provider, error)
	GetSubnet() (datatypes.Network_Subnet, error)
	GetSubnetIpAddress() (datatypes.Network_Subnet_IpAddress, error)
	GetTerminatedUser() (datatypes.User_Employee, error)
	GetTicket() (datatypes.Ticket, error)
	GetTransactions() (datatypes.Provisioning_Version1_Transaction, error)
	GetUserDepartment() (datatypes.User_Employee_Department, error)
	GetUserRecord() (datatypes.User_Employee, error)
}

type Network_Regional_Internet_Registry interface {
	Entity
}

type Network_Security_Scanner_Request interface {
	Entity

	CreateObject(templateObject datatypes.Network_Security_Scanner_Request) (datatypes.Network_Security_Scanner_Request, error)
	GetObject() (datatypes.Network_Security_Scanner_Request, error)
	GetReport() (string, error)

	GetAccount() (datatypes.Account, error)
	GetGuest() (datatypes.Virtual_Guest, error)
	GetHardware() (datatypes.Hardware, error)
	GetRequestorOwnedFlag() (bool, error)
	GetStatus() (datatypes.Network_Security_Scanner_Request_Status, error)
}

type Network_Security_Scanner_Request_Status interface {
	Entity
}

type Network_Service_Health interface {
	Entity

	GetLocation() (datatypes.Location, error)
	GetStatus() (datatypes.Network_Service_Health_Status, error)
}

type Network_Service_Health_Status interface {
	Entity
}

type Network_Service_Resource interface {
	Entity

	GetApiHost() (string, error)
	GetApiPassword() (string, error)
	GetApiPath() (string, error)
	GetApiPort() (string, error)
	GetApiProtocol() (string, error)
	GetApiUsername() (string, error)
	GetApiVersion() (string, error)
	GetAttributes() (datatypes.Network_Service_Resource_Attribute, error)
	GetDatacenter() (datatypes.Location, error)
	GetNetworkDevice() (datatypes.Hardware, error)
	GetSshUsername() (string, error)
	GetType() (datatypes.Network_Service_Resource_Type, error)
}

type Network_Service_Resource_Attribute interface {
	Entity

	GetAttributeType() (datatypes.Network_Service_Resource_Attribute_Type, error)
	GetServiceResource() (datatypes.Network_Service_Resource, error)
}

type Network_Service_Resource_Attribute_Type interface {
	Entity
}

type Network_Service_Resource_Hub interface {
	Network_Service_Resource
}

type Network_Service_Resource_Hub_Swift interface {
	Network_Service_Resource_Hub
}

type Network_Service_Resource_MonitoringHub interface {
	Network_Service_Resource

	GetAdnServicesIp() (string, error)
	GetHubAddress() (string, error)
	GetHubConnectionTimeout() (string, error)
	GetRobotsCount() (string, error)
	GetRobotsMax() (string, error)
}

type Network_Service_Resource_NimsoftLandingHub interface {
	Network_Service_Resource_MonitoringHub
}

type Network_Service_Resource_Type interface {
	Entity

	GetServiceResources() (datatypes.Network_Service_Resource, error)
}

type Network_Service_Vpn_Overrides interface {
	Entity

	CreateObjects(templateObjects datatypes.Network_Service_Vpn_Overrides) (bool, error)
	DeleteObject() (bool, error)
	DeleteObjects(templateObjects datatypes.Network_Service_Vpn_Overrides) (bool, error)
	GetObject() (datatypes.Network_Service_Vpn_Overrides, error)

	GetSubnet() (datatypes.Network_Subnet, error)
	GetUser() (datatypes.User_Customer, error)
}

type Network_Storage interface {
	Entity

	AllowAccessFromHardware(hardwareObjectTemplate datatypes.Hardware) (bool, error)
	AllowAccessFromHardwareList(hardwareObjectTemplates datatypes.Hardware) (bool, error)
	AllowAccessFromHost(typeClassName string, hostId int) (datatypes.Network_Storage_Allowed_Host, error)
	AllowAccessFromHostList(hostObjectTemplates datatypes.Container_Network_Storage_Host) (datatypes.Network_Storage_Allowed_Host, error)
	AllowAccessFromIpAddress(ipAddressObjectTemplate datatypes.Network_Subnet_IpAddress) (bool, error)
	AllowAccessFromIpAddressList(ipAddressObjectTemplates datatypes.Network_Subnet_IpAddress) (bool, error)
	AllowAccessFromSubnet(subnetObjectTemplate datatypes.Network_Subnet) (bool, error)
	AllowAccessFromSubnetList(subnetObjectTemplates datatypes.Network_Subnet) (bool, error)
	AllowAccessFromVirtualGuest(virtualGuestObjectTemplate datatypes.Virtual_Guest) (bool, error)
	AllowAccessFromVirtualGuestList(virtualGuestObjectTemplates datatypes.Virtual_Guest) (bool, error)
	AllowAccessToReplicantFromHardware(hardwareObjectTemplate datatypes.Hardware) (bool, error)
	AllowAccessToReplicantFromHardwareList(hardwareObjectTemplates datatypes.Hardware) (bool, error)
	AllowAccessToReplicantFromIpAddress(ipAddressObjectTemplate datatypes.Network_Subnet_IpAddress) (bool, error)
	AllowAccessToReplicantFromIpAddressList(ipAddressObjectTemplates datatypes.Network_Subnet_IpAddress) (bool, error)
	AllowAccessToReplicantFromSubnet(subnetObjectTemplate datatypes.Network_Subnet) (bool, error)
	AllowAccessToReplicantFromSubnetList(subnetObjectTemplates datatypes.Network_Subnet) (bool, error)
	AllowAccessToReplicantFromVirtualGuest(virtualGuestObjectTemplate datatypes.Virtual_Guest) (bool, error)
	AllowAccessToReplicantFromVirtualGuestList(virtualGuestObjectTemplates datatypes.Virtual_Guest) (bool, error)
	AssignCredential(username string) (bool, error)
	AssignNewCredential(typ string) (datatypes.Network_Storage_Credential, error)
	ChangePassword(username string, currentPassword string, newPassword string) (bool, error)
	CollectBandwidth(typ string, startDate time.Time, endDate time.Time) (uint, error)
	CollectBytesUsed() (uint, error)
	CreateFolder(folder string) (bool, error)
	CreateSnapshot(notes string) (datatypes.Network_Storage, error)
	DeleteAllFiles() (bool, error)
	DeleteFile(fileId string) (bool, error)
	DeleteFiles(fileIds string) (bool, error)
	DeleteFolder(folder string) (bool, error)
	DeleteObject() (bool, error)
	DisableSnapshots(scheduleType string) (bool, error)
	DownloadFile(fileId string) (datatypes.Container_Utility_File_Entity, error)
	EditCredential(username string, newPassword string) (bool, error)
	EditObject(templateObject datatypes.Network_Storage) (bool, error)
	EnableSnapshots(scheduleType string, retentionCount int, minute int, hour int, dayOfWeek string) (bool, error)
	FailbackFromReplicant() (bool, error)
	FailoverToReplicant(replicantId int) (bool, error)
	GetAllFiles() (datatypes.Container_Utility_File_Entity, error)
	GetAllFilesByFilter(filter datatypes.Container_Utility_File_Entity) (datatypes.Container_Utility_File_Entity, error)
	GetAllowableHardware(filterHostname string) (datatypes.Hardware, error)
	GetAllowableIpAddresses(subnetId int, filterIpAddress string) (datatypes.Network_Subnet_IpAddress, error)
	GetAllowableSubnets(filterNetworkIdentifier string) (datatypes.Network_Subnet, error)
	GetAllowableVirtualGuests(filterHostname string) (datatypes.Virtual_Guest, error)
	GetAllowedHostsLimit() (int, error)
	GetByUsername(username string, typ string) (datatypes.Network_Storage, error)
	GetCdnUrls() (datatypes.Container_Network_Storage_Hub_ObjectStorage_ContentDeliveryUrl, error)
	GetClusterResource() (datatypes.Network_Service_Resource, error)
	GetFileByIdentifier(identifier string) (datatypes.Container_Utility_File_Entity, error)
	GetFileCount() (int, error)
	GetFileList(folder string, path string) (datatypes.Container_Utility_File_Entity, error)
	GetFilePendingDeleteCount() (int, error)
	GetFilesPendingDelete() (datatypes.Container_Utility_File_Entity, error)
	GetFolderList() (datatypes.Container_Network_Storage_Hub_ObjectStorage_Folder, error)
	GetGraph(startDate time.Time, endDate time.Time, typ string) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetNetworkConnectionDetails() (datatypes.Container_Network_Storage_NetworkConnectionInformation, error)
	GetObject() (datatypes.Network_Storage, error)
	GetObjectStorageConnectionInformation() (datatypes.Container_Network_Service_Resource_ObjectStorage_ConnectionInformation, error)
	GetObjectsByCredential(credentialObject datatypes.Network_Storage_Credential) (datatypes.Network_Storage, error)
	GetRecycleBinFileByIdentifier(fileId string) (datatypes.Container_Utility_File_Entity, error)
	GetRemainingAllowedHosts() (int, error)
	GetSnapshotsForVolume() (datatypes.Network_Storage, error)
	GetStorageGroupsNetworkConnectionDetails() (datatypes.Container_Network_Storage_NetworkConnectionInformation, error)
	GetValidReplicationTargetDatacenterLocations() (datatypes.Location, error)
	ImmediateFailoverToReplicant(replicantId int) (bool, error)
	IsBlockingOperationInProgress(exemptStatusKeyNames string) (bool, error)
	RemoveAccessFromHardware(hardwareObjectTemplate datatypes.Hardware) (bool, error)
	RemoveAccessFromHardwareList(hardwareObjectTemplates datatypes.Hardware) (bool, error)
	RemoveAccessFromHost(typeClassName string, hostId int) (datatypes.Network_Storage_Allowed_Host, error)
	RemoveAccessFromHostList(hostObjectTemplates datatypes.Container_Network_Storage_Host) (datatypes.Network_Storage_Allowed_Host, error)
	RemoveAccessFromIpAddress(ipAddressObjectTemplate datatypes.Network_Subnet_IpAddress) (bool, error)
	RemoveAccessFromIpAddressList(ipAddressObjectTemplates datatypes.Network_Subnet_IpAddress) (bool, error)
	RemoveAccessFromSubnet(subnetObjectTemplate datatypes.Network_Subnet) (bool, error)
	RemoveAccessFromSubnetList(subnetObjectTemplates datatypes.Network_Subnet) (bool, error)
	RemoveAccessFromVirtualGuest(virtualGuestObjectTemplate datatypes.Virtual_Guest) (bool, error)
	RemoveAccessFromVirtualGuestList(virtualGuestObjectTemplates datatypes.Virtual_Guest) (bool, error)
	RemoveAccessToReplicantFromHardwareList(hardwareObjectTemplates datatypes.Hardware) (bool, error)
	RemoveAccessToReplicantFromIpAddressList(ipAddressObjectTemplates datatypes.Network_Subnet_IpAddress) (bool, error)
	RemoveAccessToReplicantFromSubnet(subnetObjectTemplate datatypes.Network_Subnet) (bool, error)
	RemoveAccessToReplicantFromSubnetList(subnetObjectTemplates datatypes.Network_Subnet) (bool, error)
	RemoveAccessToReplicantFromVirtualGuestList(virtualGuestObjectTemplates datatypes.Virtual_Guest) (bool, error)
	RemoveCredential(username string) (bool, error)
	RestoreFile(fileId string) (datatypes.Container_Utility_File_Entity, error)
	RestoreFromSnapshot(snapshotId int) (bool, error)
	SendPasswordReminderEmail(username string) (bool, error)
	SetMountable(mountable bool) (bool, error)
	SetSnapshotAllocation(capacityGb int) error
	UpgradeVolumeCapacity(itemId int) (bool, error)
	UploadFile(file datatypes.Container_Utility_File_Entity) (datatypes.Container_Utility_File_Entity, error)

	GetAccount() (datatypes.Account, error)
	GetAccountPassword() (datatypes.Account_Password, error)
	GetActiveTransactions() (datatypes.Provisioning_Version1_Transaction, error)
	GetAllowedHardware() (datatypes.Hardware, error)
	GetAllowedIpAddresses() (datatypes.Network_Subnet_IpAddress, error)
	GetAllowedReplicationHardware() (datatypes.Hardware, error)
	GetAllowedReplicationIpAddresses() (datatypes.Network_Subnet_IpAddress, error)
	GetAllowedReplicationSubnets() (datatypes.Network_Subnet, error)
	GetAllowedReplicationVirtualGuests() (datatypes.Virtual_Guest, error)
	GetAllowedSubnets() (datatypes.Network_Subnet, error)
	GetAllowedVirtualGuests() (datatypes.Virtual_Guest, error)
	GetBillingItem() (datatypes.Billing_Item, error)
	GetBillingItemCategory() (datatypes.Product_Item_Category, error)
	GetBytesUsed() (string, error)
	GetCreationScheduleId() (string, error)
	GetCredentials() (datatypes.Network_Storage_Credential, error)
	GetDailySchedule() (datatypes.Network_Storage_Schedule, error)
	GetEvents() (datatypes.Network_Storage_Event, error)
	GetHardware() (datatypes.Hardware, error)
	GetHourlySchedule() (datatypes.Network_Storage_Schedule, error)
	GetIops() (string, error)
	GetIscsiLuns() (datatypes.Network_Storage, error)
	GetLunId() (string, error)
	GetManualSnapshots() (datatypes.Network_Storage, error)
	GetMetricTrackingObject() (datatypes.Metric_Tracking_Object, error)
	GetMountableFlag() (string, error)
	GetNotificationSubscribers() (datatypes.Notification_User_Subscriber, error)
	GetOsType() (datatypes.Network_Storage_Iscsi_OS_Type, error)
	GetOsTypeId() (string, error)
	GetParentPartnerships() (datatypes.Network_Storage_Partnership, error)
	GetParentVolume() (datatypes.Network_Storage, error)
	GetPartnerships() (datatypes.Network_Storage_Partnership, error)
	GetPermissionsGroups() (datatypes.Network_Storage_Group, error)
	GetProperties() (datatypes.Network_Storage_Property, error)
	GetReplicatingLuns() (datatypes.Network_Storage, error)
	GetReplicatingVolume() (datatypes.Network_Storage, error)
	GetReplicationEvents() (datatypes.Network_Storage_Event, error)
	GetReplicationPartners() (datatypes.Network_Storage, error)
	GetReplicationSchedule() (datatypes.Network_Storage_Schedule, error)
	GetReplicationStatus() (string, error)
	GetSchedules() (datatypes.Network_Storage_Schedule, error)
	GetServiceResource() (datatypes.Network_Service_Resource, error)
	GetServiceResourceBackendIpAddress() (string, error)
	GetServiceResourceName() (string, error)
	GetSnapshotCapacityGb() (string, error)
	GetSnapshotCreationTimestamp() (string, error)
	GetSnapshotDeletionThresholdPercentage() (string, error)
	GetSnapshotSizeBytes() (string, error)
	GetSnapshotSpaceAvailable() (string, error)
	GetSnapshots() (datatypes.Network_Storage, error)
	GetStorageGroups() (datatypes.Network_Storage_Group, error)
	GetStorageTierLevel() (string, error)
	GetStorageType() (datatypes.Network_Storage_Type, error)
	GetTotalBytesUsed() (string, error)
	GetTotalScheduleSnapshotRetentionCount() (uint, error)
	GetUsageNotification() (datatypes.Notification, error)
	GetVendorName() (string, error)
	GetVirtualGuest() (datatypes.Virtual_Guest, error)
	GetVolumeHistory() (datatypes.Network_Storage_History, error)
	GetVolumeStatus() (string, error)
	GetWebccAccount() (datatypes.Account_Password, error)
	GetWeeklySchedule() (datatypes.Network_Storage_Schedule, error)
}

type Network_Storage_Allowed_Host interface {
	Entity

	CreateObject(templateObject datatypes.Network_Storage_Allowed_Host) (bool, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Network_Storage_Allowed_Host) (bool, error)
	GetObject() (datatypes.Network_Storage_Allowed_Host, error)
	SetCredentialPassword(password string) (bool, error)

	GetAssignedGroups() (datatypes.Network_Storage_Group, error)
	GetAssignedReplicationVolumes() (datatypes.Network_Storage, error)
	GetAssignedVolumes() (datatypes.Network_Storage, error)
	GetCredential() (datatypes.Network_Storage_Credential, error)
}

type Network_Storage_Allowed_Host_Hardware interface {
	Network_Storage_Allowed_Host

	GetObject() (datatypes.Network_Storage_Allowed_Host_Hardware, error)

	GetResource() (datatypes.Hardware, error)
}

type Network_Storage_Allowed_Host_IpAddress interface {
	Network_Storage_Allowed_Host

	GetObject() (datatypes.Network_Storage_Allowed_Host_IpAddress, error)

	GetResource() (datatypes.Network_Subnet_IpAddress, error)
}

type Network_Storage_Allowed_Host_Subnet interface {
	Network_Storage_Allowed_Host

	GetObject() (datatypes.Network_Storage_Allowed_Host_Subnet, error)

	GetResource() (datatypes.Network_Subnet, error)
}

type Network_Storage_Allowed_Host_VirtualGuest interface {
	Network_Storage_Allowed_Host

	GetObject() (datatypes.Network_Storage_Allowed_Host_VirtualGuest, error)

	GetResource() (datatypes.Virtual_Guest, error)
}

type Network_Storage_Backup interface {
	Network_Storage

	GetCurrentCyclePeakUsage() (uint, error)
	GetPreviousCyclePeakUsage() (uint, error)
}

type Network_Storage_Backup_Evault interface {
	Network_Storage_Backup

	DeleteTasks(tasks int) (bool, error)
	GetHardwareWithEvaultFirst(option string, exactMatch bool, criteria string, mode string) (datatypes.Hardware, error)
	GetObject() (datatypes.Network_Storage_Backup_Evault, error)
	GetWebCCAuthenticationDetails() (datatypes.Container_Network_Storage_Backup_Evault_WebCc_Authentication_Details, error)
	InitiateBareMetalRestore() (bool, error)
	InitiateBareMetalRestoreForServer(hardwareId int) (bool, error)
}

type Network_Storage_Backup_Evault_Version6 interface {
	Network_Storage_Backup_Evault

	GetAgentStatuses() (datatypes.Container_Network_Storage_Evault_WebCc_AgentStatus, error)
	GetBackupJobDetails() (datatypes.Container_Network_Storage_Evault_WebCc_JobDetails, error)
	GetPluginBillingItems() (datatypes.Billing_Item, error)
	GetRestoreJobDetails() (datatypes.Container_Network_Storage_Evault_WebCc_JobDetails, error)
	GetSoftwareComponent() (datatypes.Software_Component, error)
	GetTasks() (datatypes.Container_Network_Storage_Evault_Vault_Task, error)
}

type Network_Storage_Credential interface {
	Entity

	GetAccount() (datatypes.Account, error)
	GetNetworkStorageAllowedHosts() (datatypes.Network_Storage_Allowed_Host, error)
	GetType() (datatypes.Network_Storage_Credential_Type, error)
	GetVolumes() (datatypes.Network_Storage, error)
}

type Network_Storage_Credential_Type interface {
	Entity
}

type Network_Storage_Daily_Usage interface {
	Entity

	GetNasVolume() (datatypes.Network_Storage, error)
}

type Network_Storage_Event interface {
	Entity

	GetSchedule() (datatypes.Network_Storage_Schedule, error)
	GetVolume() (datatypes.Network_Storage, error)
}

type Network_Storage_Group interface {
	Entity

	AddAllowedHost(allowedHost datatypes.Network_Storage_Allowed_Host) (bool, error)
	AttachToVolume(volume datatypes.Network_Storage) (bool, error)
	CreateObject(templateObject datatypes.Network_Storage_Group) (bool, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Network_Storage_Group) (bool, error)
	GetAllObjects() (datatypes.Network_Storage_Group, error)
	GetNetworkConnectionDetails() (datatypes.Container_Network_Storage_NetworkConnectionInformation, error)
	GetObject() (datatypes.Network_Storage_Group, error)
	RemoveAllowedHost(allowedHost datatypes.Network_Storage_Allowed_Host) (bool, error)
	RemoveFromVolume(volume datatypes.Network_Storage) (bool, error)

	GetAccount() (datatypes.Account, error)
	GetAllowedHosts() (datatypes.Network_Storage_Allowed_Host, error)
	GetAttachedVolumes() (datatypes.Network_Storage, error)
	GetGroupType() (datatypes.Network_Storage_Group_Type, error)
	GetOsType() (datatypes.Network_Storage_Iscsi_OS_Type, error)
	GetServiceResource() (datatypes.Network_Service_Resource, error)
}

type Network_Storage_Group_Iscsi interface {
	Network_Storage_Group

	AddAllowedHost(allowedHost datatypes.Network_Storage_Allowed_Host) (bool, error)
	AttachToVolume(volume datatypes.Network_Storage) (bool, error)
	GetObject() (datatypes.Network_Storage_Group_Iscsi, error)
	RemoveAllowedHost(allowedHost datatypes.Network_Storage_Allowed_Host) (bool, error)
	RemoveFromVolume(volume datatypes.Network_Storage) (bool, error)
}

type Network_Storage_Group_Nfs interface {
	Network_Storage_Group

	AddAllowedHost(allowedHost datatypes.Network_Storage_Allowed_Host) (bool, error)
	AttachToVolume(volume datatypes.Network_Storage) (bool, error)
	GetObject() (datatypes.Network_Storage_Group_Nfs, error)
	RemoveAllowedHost(allowedHost datatypes.Network_Storage_Allowed_Host) (bool, error)
	RemoveFromVolume(volume datatypes.Network_Storage) (bool, error)
}

type Network_Storage_Group_Type interface {
	Entity

	GetAllObjects() (datatypes.Network_Storage_Group_Type, error)
	GetObject() (datatypes.Network_Storage_Group_Type, error)
}

type Network_Storage_History interface {
	Entity

	GetAccount() (datatypes.Account, error)
	GetNasVolume() (datatypes.Network_Storage, error)
}

type Network_Storage_Hub interface {
	Network_Storage

	GetBandwidthBillingItems() (datatypes.Billing_Item, error)
}

type Network_Storage_Hub_Cleversafe_Account interface {
	Entity

	CredentialCreate() (datatypes.Network_Storage_Credential, error)
	CredentialDelete(credential datatypes.Network_Storage_Credential) (bool, error)
	GetCapacityUsage() (int, error)
	GetCredentialLimit() (int, error)
	GetEndpoints() (datatypes.Container_Network_Storage_Hub_ObjectStorage_Endpoint, error)
	GetObject() (datatypes.Network_Storage_Hub_Cleversafe_Account, error)

	GetAccount() (datatypes.Account, error)
	GetBillingItem() (datatypes.Billing_Item, error)
	GetCancelledBillingItem() (datatypes.Billing_Item, error)
	GetCredentials() (datatypes.Network_Storage_Credential, error)
	GetEvents() (datatypes.Network_Storage_Event, error)
	GetMetricTrackingObject() (datatypes.Metric_Tracking_Object, error)
	GetUuid() (string, error)
}

type Network_Storage_Hub_Swift interface {
	Network_Storage_Hub

	GetStorageNodes() (datatypes.Network_Service_Resource, error)
}

type Network_Storage_Hub_Swift_Container interface {
	Network_Storage_Hub_Swift
}

type Network_Storage_Hub_Swift_Share interface {
	Entity

	GetContainerList() (datatypes.Container_Network_Storage_Hub_ObjectStorage_Folder, error)
	GetFile(fileName string, container string) (datatypes.Container_Network_Storage_Hub_ObjectStorage_File, error)
	GetFileList(container string, path string) (datatypes.Container_Utility_File_Entity, error)
}

type Network_Storage_Hub_Swift_Version1 interface {
	Network_Storage_Hub_Swift
}

type Network_Storage_Iscsi interface {
	Network_Storage

	AllowAccessFromHardware(hardwareObjectTemplate datatypes.Hardware) (bool, error)
	AllowAccessFromIpAddress(ipAddressObjectTemplate datatypes.Network_Subnet_IpAddress) (bool, error)
	AllowAccessFromVirtualGuest(virtualGuestObjectTemplate datatypes.Virtual_Guest) (bool, error)
	AllowAccessToReplicantFromHardwareList(hardwareObjectTemplates datatypes.Hardware) (bool, error)
	AllowAccessToReplicantFromIpAddressList(ipAddressObjectTemplates datatypes.Network_Subnet_IpAddress) (bool, error)
	AllowAccessToReplicantFromVirtualGuestList(virtualGuestObjectTemplates datatypes.Virtual_Guest) (bool, error)
	GetObject() (datatypes.Network_Storage_Iscsi, error)
	GetSnapshotsForVolume() (datatypes.Network_Storage, error)
	RemoveAccessFromHardware(hardwareObjectTemplate datatypes.Hardware) (bool, error)
	RemoveAccessFromIpAddress(ipAddressObjectTemplate datatypes.Network_Subnet_IpAddress) (bool, error)
	RemoveAccessFromVirtualGuest(virtualGuestObjectTemplate datatypes.Virtual_Guest) (bool, error)
	RemoveAccessToReplicantFromHardwareList(hardwareObjectTemplates datatypes.Hardware) (bool, error)
	RemoveAccessToReplicantFromIpAddressList(ipAddressObjectTemplates datatypes.Network_Subnet_IpAddress) (bool, error)
	RemoveAccessToReplicantFromVirtualGuestList(virtualGuestObjectTemplates datatypes.Virtual_Guest) (bool, error)
}

type Network_Storage_Iscsi_EqualLogic_Version3 interface {
	Network_Storage_Iscsi
}

type Network_Storage_Iscsi_EqualLogic_Version3_Replicant interface {
	Network_Storage_Iscsi_EqualLogic_Version3

	GetFailbackInProgressFlag() (bool, error)
	GetVolumeName() (string, error)
}

type Network_Storage_Iscsi_EqualLogic_Version3_Snapshot interface {
	Network_Storage_Iscsi_EqualLogic_Version3

	GetCreationSchedule() (datatypes.Network_Storage_Schedule, error)
	GetVolumeName() (string, error)
}

type Network_Storage_Iscsi_OS_Type interface {
	Entity

	GetAllObjects() (datatypes.Network_Storage_Iscsi_OS_Type, error)
	GetObject() (datatypes.Network_Storage_Iscsi_OS_Type, error)
}

type Network_Storage_Nas interface {
	Network_Storage

	GetRecentBytesUsed() (datatypes.Network_Storage_Daily_Usage, error)
}

type Network_Storage_OpenStack_Object interface {
	Network_Storage

	GetBandwidthBillingItems() (datatypes.Billing_Item, error)
}

type Network_Storage_Partnership interface {
	Entity

	GetPartnerVolume() (datatypes.Network_Storage, error)
	GetType() (datatypes.Network_Storage_Partnership_Type, error)
	GetVolume() (datatypes.Network_Storage, error)
}

type Network_Storage_Partnership_Type interface {
	Entity
}

type Network_Storage_Property interface {
	Entity

	GetType() (datatypes.Network_Storage_Property_Type, error)
	GetVolume() (datatypes.Network_Storage, error)
}

type Network_Storage_Property_Type interface {
	Entity
}

type Network_Storage_Replicant interface {
	Network_Storage

	GetFailbackInProgressFlag() (string, error)
	GetVolumeName() (string, error)
}

type Network_Storage_Schedule interface {
	Entity

	CreateObject(templateObject datatypes.Network_Storage_Schedule) (datatypes.Network_Storage_Schedule, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Network_Storage_Schedule) (bool, error)
	GetObject() (datatypes.Network_Storage_Schedule, error)

	GetDayOfMonth() (string, error)
	GetDayOfWeek() (string, error)
	GetEvents() (datatypes.Network_Storage_Event, error)
	GetHour() (string, error)
	GetMinute() (string, error)
	GetMonthOfYear() (string, error)
	GetPartnership() (datatypes.Network_Storage_Partnership, error)
	GetProperties() (datatypes.Network_Storage_Schedule_Property, error)
	GetReplicaSnapshots() (datatypes.Network_Storage, error)
	GetRetentionCount() (string, error)
	GetSnapshots() (datatypes.Network_Storage, error)
	GetType() (datatypes.Network_Storage_Schedule_Type, error)
	GetVolume() (datatypes.Network_Storage, error)
}

type Network_Storage_Schedule_Property interface {
	Entity

	GetSchedule() (datatypes.Network_Storage_Schedule, error)
	GetType() (datatypes.Network_Storage_Schedule_Property_Type, error)
}

type Network_Storage_Schedule_Property_Type interface {
	Entity

	GetAllObjects() (datatypes.Network_Storage_Schedule_Property_Type, error)
	GetObject() (datatypes.Network_Storage_Schedule_Property_Type, error)
}

type Network_Storage_Schedule_Type interface {
	Entity
}

type Network_Storage_Snapshot interface {
	Network_Storage

	GetCreationSchedule() (datatypes.Network_Storage_Schedule, error)
	GetVolumeName() (string, error)
}

type Network_Storage_Type interface {
	Entity

	GetVolumes() (datatypes.Network_Storage, error)
}

type Network_Subnet interface {
	Entity

	AllowAccessToNetworkStorage(networkStorageTemplateObject datatypes.Network_Storage) (bool, error)
	AllowAccessToNetworkStorageList(networkStorageTemplateObjects datatypes.Network_Storage) (bool, error)
	CreateReverseDomainRecords() (datatypes.Dns_Domain_Reverse, error)
	CreateSubnetRouteUpdateTransaction(newEndPointIpAddress string) (bool, error)
	CreateSwipTransaction() (bool, error)
	EditNote(note string) (bool, error)
	FindAllSubnetsAndActiveSwipTransactionStatus() (datatypes.Network_Subnet, error)
	GetAttachedNetworkStorages(nasType string) (datatypes.Network_Storage, error)
	GetAvailableNetworkStorages(nasType string) (datatypes.Network_Storage, error)
	GetObject() (datatypes.Network_Subnet, error)
	GetReverseDomainRecords() (datatypes.Dns_Domain, error)
	GetRoutableEndpointIpAddresses() (datatypes.Network_Subnet_IpAddress, error)
	GetSubnetForIpAddress(ipAddress string) (datatypes.Network_Subnet, error)
	RemoveAccessToNetworkStorageList(networkStorageTemplateObjects datatypes.Network_Storage) (bool, error)

	GetAccount() (datatypes.Account, error)
	GetActiveRegistration() (datatypes.Network_Subnet_Registration, error)
	GetActiveSwipTransaction() (datatypes.Network_Subnet_Swip_Transaction, error)
	GetActiveTransaction() (datatypes.Provisioning_Version1_Transaction, error)
	GetAddressSpace() (string, error)
	GetAllowedHost() (datatypes.Network_Storage_Allowed_Host, error)
	GetAllowedNetworkStorage() (datatypes.Network_Storage, error)
	GetAllowedNetworkStorageReplicas() (datatypes.Network_Storage, error)
	GetBillingItem() (datatypes.Billing_Item, error)
	GetBoundDescendants() (datatypes.Network_Subnet, error)
	GetBoundRouterFlag() (bool, error)
	GetBoundRouters() (datatypes.Hardware, error)
	GetChildren() (datatypes.Network_Subnet, error)
	GetDatacenter() (datatypes.Location_Datacenter, error)
	GetDescendants() (datatypes.Network_Subnet, error)
	GetDisplayLabel() (string, error)
	GetEndPointIpAddress() (datatypes.Network_Subnet_IpAddress, error)
	GetGlobalIpRecord() (datatypes.Network_Subnet_IpAddress_Global, error)
	GetHardware() (datatypes.Hardware, error)
	GetIpAddresses() (datatypes.Network_Subnet_IpAddress, error)
	GetNetworkComponent() (datatypes.Network_Component, error)
	GetNetworkComponentFirewall() (datatypes.Network_Component_Firewall, error)
	GetNetworkProtectionAddresses() (datatypes.Network_Protection_Address, error)
	GetNetworkTunnelContexts() (datatypes.Network_Tunnel_Module_Context, error)
	GetNetworkVlan() (datatypes.Network_Vlan, error)
	GetPodName() (string, error)
	GetProtectedIpAddresses() (datatypes.Network_Subnet_IpAddress, error)
	GetRegionalInternetRegistry() (datatypes.Network_Regional_Internet_Registry, error)
	GetRegistrations() (datatypes.Network_Subnet_Registration, error)
	GetResourceGroups() (datatypes.Resource_Group, error)
	GetReverseDomain() (datatypes.Dns_Domain, error)
	GetRoleKeyName() (string, error)
	GetRoleName() (string, error)
	GetRoutingTypeKeyName() (string, error)
	GetRoutingTypeName() (string, error)
	GetSwipTransaction() (datatypes.Network_Subnet_Swip_Transaction, error)
	GetUnboundDescendants() (datatypes.Network_Subnet, error)
	GetVirtualGuests() (datatypes.Virtual_Guest, error)
}

type Network_Subnet_IpAddress interface {
	Entity

	AllowAccessToNetworkStorage(networkStorageTemplateObject datatypes.Network_Storage) (bool, error)
	AllowAccessToNetworkStorageList(networkStorageTemplateObjects datatypes.Network_Storage) (bool, error)
	EditObject(templateObject datatypes.Network_Subnet_IpAddress) (bool, error)
	EditObjects(templateObjects datatypes.Network_Subnet_IpAddress) (bool, error)
	FindByIpv4Address(ipAddress string) (datatypes.Network_Subnet_IpAddress, error)
	GetAttachedNetworkStorages(nasType string) (datatypes.Network_Storage, error)
	GetAvailableNetworkStorages(nasType string) (datatypes.Network_Storage, error)
	GetByIpAddress(ipAddress string) (datatypes.Network_Subnet_IpAddress, error)
	GetObject() (datatypes.Network_Subnet_IpAddress, error)
	RemoveAccessToNetworkStorageList(networkStorageTemplateObjects datatypes.Network_Storage) (bool, error)

	GetAllowedHost() (datatypes.Network_Storage_Allowed_Host, error)
	GetAllowedNetworkStorage() (datatypes.Network_Storage, error)
	GetAllowedNetworkStorageReplicas() (datatypes.Network_Storage, error)
	GetApplicationDeliveryController() (datatypes.Network_Application_Delivery_Controller, error)
	GetContextTunnelTranslations() (datatypes.Network_Tunnel_Module_Context_Address_Translation, error)
	GetEndpointSubnets() (datatypes.Network_Subnet, error)
	GetGuestNetworkComponent() (datatypes.Virtual_Guest_Network_Component, error)
	GetGuestNetworkComponentBinding() (datatypes.Virtual_Guest_Network_Component_IpAddress, error)
	GetHardware() (datatypes.Hardware, error)
	GetNetworkComponent() (datatypes.Network_Component, error)
	GetPrivateNetworkGateway() (datatypes.Network_Gateway, error)
	GetProtectionAddress() (datatypes.Network_Protection_Address, error)
	GetPublicNetworkGateway() (datatypes.Network_Gateway, error)
	GetRemoteManagementNetworkComponent() (datatypes.Network_Component, error)
	GetSubnet() (datatypes.Network_Subnet, error)
	GetSyslogEventsOneDay() (datatypes.Network_Logging_Syslog, error)
	GetSyslogEventsSevenDays() (datatypes.Network_Logging_Syslog, error)
	GetTopTenSyslogEventsByDestinationPortOneDay() (datatypes.Network_Logging_Syslog, error)
	GetTopTenSyslogEventsByDestinationPortSevenDays() (datatypes.Network_Logging_Syslog, error)
	GetTopTenSyslogEventsByProtocolsOneDay() (datatypes.Network_Logging_Syslog, error)
	GetTopTenSyslogEventsByProtocolsSevenDays() (datatypes.Network_Logging_Syslog, error)
	GetTopTenSyslogEventsBySourceIpOneDay() (datatypes.Network_Logging_Syslog, error)
	GetTopTenSyslogEventsBySourceIpSevenDays() (datatypes.Network_Logging_Syslog, error)
	GetTopTenSyslogEventsBySourcePortOneDay() (datatypes.Network_Logging_Syslog, error)
	GetTopTenSyslogEventsBySourcePortSevenDays() (datatypes.Network_Logging_Syslog, error)
	GetVirtualGuest() (datatypes.Virtual_Guest, error)
	GetVirtualLicenses() (datatypes.Software_VirtualLicense, error)
}

type Network_Subnet_IpAddress_Global interface {
	Entity

	GetObject() (datatypes.Network_Subnet_IpAddress_Global, error)
	Route(newEndPointIpAddress string) (datatypes.Provisioning_Version1_Transaction, error)
	Unroute() (datatypes.Provisioning_Version1_Transaction, error)

	GetAccount() (datatypes.Account, error)
	GetActiveTransaction() (datatypes.Provisioning_Version1_Transaction, error)
	GetBillingItem() (datatypes.Billing_Item_Network_Subnet_IpAddress_Global, error)
	GetDestinationIpAddress() (datatypes.Network_Subnet_IpAddress, error)
	GetIpAddress() (datatypes.Network_Subnet_IpAddress, error)
}

type Network_Subnet_IpAddress_Version6 interface {
	Network_Subnet_IpAddress

	GetPublicVersion6NetworkGateway() (datatypes.Network_Gateway, error)
}

type Network_Subnet_Registration interface {
	Entity

	ClearRegistration() (bool, error)
	CreateObject(templateObject datatypes.Network_Subnet_Registration) (datatypes.Network_Subnet_Registration, error)
	EditObject(templateObject datatypes.Network_Subnet_Registration) (bool, error)
	EditRegistrationAttachedDetails(personObjectSkeleton datatypes.Network_Subnet_Registration_Details, networkObjectSkeleton datatypes.Network_Subnet_Registration_Details) (bool, error)
	GetObject() (datatypes.Network_Subnet_Registration, error)

	GetAccount() (datatypes.Account, error)
	GetDetailReferences() (datatypes.Network_Subnet_Registration_Details, error)
	GetEvents() (datatypes.Network_Subnet_Registration_Event, error)
	GetNetworkDetail() (datatypes.Account_Regional_Registry_Detail, error)
	GetPersonDetail() (datatypes.Account_Regional_Registry_Detail, error)
	GetRegionalInternetRegistry() (datatypes.Network_Regional_Internet_Registry, error)
	GetRegionalInternetRegistryHandle() (datatypes.Account_Rwhois_Handle, error)
	GetStatus() (datatypes.Network_Subnet_Registration_Status, error)
	GetSubnet() (datatypes.Network_Subnet, error)
}

type Network_Subnet_Registration_Apnic interface {
	Network_Subnet_Registration
}

type Network_Subnet_Registration_Arin interface {
	Network_Subnet_Registration
}

type Network_Subnet_Registration_Details interface {
	Entity

	CreateObject(templateObject datatypes.Network_Subnet_Registration_Details) (datatypes.Network_Subnet_Registration_Details, error)
	DeleteObject() (bool, error)
	GetObject() (datatypes.Network_Subnet_Registration_Details, error)

	GetDetail() (datatypes.Account_Regional_Registry_Detail, error)
	GetRegistration() (datatypes.Network_Subnet_Registration, error)
}

type Network_Subnet_Registration_Event interface {
	Entity

	GetRegistration() (datatypes.Network_Subnet_Registration, error)
	GetType() (datatypes.Network_Subnet_Registration_Event_Type, error)
}

type Network_Subnet_Registration_Event_Type interface {
	Entity
}

type Network_Subnet_Registration_Ripe interface {
	Network_Subnet_Registration
}

type Network_Subnet_Registration_Status interface {
	Entity

	GetAllObjects() (datatypes.Network_Subnet_Registration_Status, error)
	GetObject() (datatypes.Network_Subnet_Registration_Status, error)
}

type Network_Subnet_Rwhois_Data interface {
	Entity

	EditObject(templateObject datatypes.Network_Subnet_Rwhois_Data) (bool, error)
	GetObject() (datatypes.Network_Subnet_Rwhois_Data, error)

	GetAccount() (datatypes.Account, error)
}

type Network_Subnet_Swip_Transaction interface {
	Entity

	FindMyTransactions() (datatypes.Network_Subnet_Swip_Transaction, error)
	GetObject() (datatypes.Network_Subnet_Swip_Transaction, error)
	RemoveAllSubnetSwips() (int, error)
	RemoveSwipData() (bool, error)
	ResendSwipData() (bool, error)
	SwipAllSubnets() (int, error)
	UpdateAllSubnetSwips() (int, error)

	GetAccount() (datatypes.Account, error)
	GetSubnet() (datatypes.Network_Subnet, error)
}

type Network_TippingPointReporting interface {
	Entity

	DrillDownAttack(signatureId string, IpAddress string, subnetMask int, timeFrame int, direction string) (datatypes.Container_Network_IntrusionProtection_SubnetReport, error)
	GetMainStatistics(numberOfAttacks int) (datatypes.Container_Network_IntrusionProtection_Statistics, error)
	GetReportForIpAddressOrSubnet(IpAddress string, subnetMask int, timeFrame int, orderBy string, orderDirection string) (datatypes.Container_Network_IntrusionProtection_SubnetReport, error)
	GetSubnetReportForEntireAccount(timeFrame int, orderBy string, orderDirection string, returnSubnetGroups bool) (datatypes.Container_Network_IntrusionProtection_SubnetReport, error)
}

type Network_Tunnel_Module_Context interface {
	Entity

	AddCustomerSubnetToNetworkTunnel(subnetId int) (bool, error)
	AddPrivateSubnetToNetworkTunnel(subnetId int) (bool, error)
	AddServiceSubnetToNetworkTunnel(subnetId int) (bool, error)
	ApplyConfigurationsToDevice() (bool, error)
	CreateAddressTranslation(translation datatypes.Network_Tunnel_Module_Context_Address_Translation) (datatypes.Network_Tunnel_Module_Context_Address_Translation, error)
	CreateAddressTranslations(translations datatypes.Network_Tunnel_Module_Context_Address_Translation) (datatypes.Network_Tunnel_Module_Context_Address_Translation, error)
	DeleteAddressTranslation(translationId int) (bool, error)
	DownloadAddressTranslationConfigurations() (datatypes.Container_Utility_File_Entity, error)
	DownloadParameterConfigurations() (datatypes.Container_Utility_File_Entity, error)
	EditAddressTranslation(translation datatypes.Network_Tunnel_Module_Context_Address_Translation) (datatypes.Network_Tunnel_Module_Context_Address_Translation, error)
	EditAddressTranslations(translations datatypes.Network_Tunnel_Module_Context_Address_Translation) (datatypes.Network_Tunnel_Module_Context_Address_Translation, error)
	EditObject(templateObject datatypes.Network_Tunnel_Module_Context) (bool, error)
	GetAddressTranslationConfigurations() (string, error)
	GetAuthenticationDefault() (string, error)
	GetAuthenticationOptions() (string, error)
	GetDiffieHellmanGroupDefault() (int, error)
	GetDiffieHellmanGroupOptions() (int, error)
	GetEncryptionDefault() (string, error)
	GetEncryptionOptions() (string, error)
	GetKeylifeLimits() (int, error)
	GetObject() (datatypes.Network_Tunnel_Module_Context, error)
	GetParameterConfigurationsForCustomerView() (string, error)
	GetPhaseOneKeylifeDefault() (string, error)
	GetPhaseTwoKeylifeDefault() (string, error)
	RemoveCustomerSubnetFromNetworkTunnel(subnetId int) (bool, error)
	RemovePrivateSubnetFromNetworkTunnel(subnetId int) (bool, error)
	RemoveServiceSubnetFromNetworkTunnel(subnetId int) (bool, error)

	GetAccount() (datatypes.Account, error)
	GetActiveTransaction() (datatypes.Provisioning_Version1_Transaction, error)
	GetAddressTranslations() (datatypes.Network_Tunnel_Module_Context_Address_Translation, error)
	GetAllAvailableServiceSubnets() (datatypes.Network_Subnet, error)
	GetBillingItem() (datatypes.Billing_Item, error)
	GetCustomerSubnets() (datatypes.Network_Customer_Subnet, error)
	GetDatacenter() (datatypes.Location, error)
	GetInternalSubnets() (datatypes.Network_Subnet, error)
	GetServiceSubnets() (datatypes.Network_Subnet, error)
	GetStaticRouteSubnets() (datatypes.Network_Subnet, error)
	GetTransactionHistory() (datatypes.Provisioning_Version1_Transaction, error)
}

type Network_Tunnel_Module_Context_Address_Translation interface {
	Entity

	GetCustomerIpAddressRecord() (datatypes.Network_Customer_Subnet_IpAddress, error)
	GetInternalIpAddressRecord() (datatypes.Network_Subnet_IpAddress, error)
	GetNetworkTunnelContext() (datatypes.Network_Tunnel_Module_Context, error)
}

type Network_Vlan interface {
	Entity

	EditObject(templateObject datatypes.Network_Vlan) (bool, error)
	GetCancelFailureReasons() (string, error)
	GetFirewallProtectableIpAddresses() (datatypes.Network_Subnet_IpAddress, error)
	GetFirewallProtectableSubnets() (datatypes.Network_Subnet, error)
	GetObject() (datatypes.Network_Vlan, error)
	GetPrivateVlan() (datatypes.Network_Vlan, error)
	GetPrivateVlanByIpAddress(ipAddress string) (datatypes.Network_Vlan, error)
	GetPublicVlanByFqdn(fqdn string) (datatypes.Network_Vlan, error)
	GetReverseDomainRecords() (datatypes.Dns_Domain, error)
	GetVlanForIpAddress(ipAddress string) (datatypes.Network_Vlan, error)
	SetTags(tags string) (bool, error)
	UpdateFirewallIntraVlanCommunication(enabled bool) error

	GetAccount() (datatypes.Account, error)
	GetAdditionalPrimarySubnets() (datatypes.Network_Subnet, error)
	GetAttachedNetworkGateway() (datatypes.Network_Gateway, error)
	GetAttachedNetworkGatewayFlag() (bool, error)
	GetAttachedNetworkGatewayVlan() (datatypes.Network_Gateway_Vlan, error)
	GetBillingItem() (datatypes.Billing_Item, error)
	GetDedicatedFirewallFlag() (int, error)
	GetExtensionRouter() (datatypes.Hardware_Router, error)
	GetFirewallGuestNetworkComponents() (datatypes.Network_Component_Firewall, error)
	GetFirewallInterfaces() (datatypes.Network_Firewall_Module_Context_Interface, error)
	GetFirewallNetworkComponents() (datatypes.Network_Component_Firewall, error)
	GetFirewallRules() (datatypes.Network_Vlan_Firewall_Rule, error)
	GetGuestNetworkComponents() (datatypes.Virtual_Guest_Network_Component, error)
	GetHardware() (datatypes.Hardware, error)
	GetHighAvailabilityFirewallFlag() (bool, error)
	GetLocalDiskStorageCapabilityFlag() (bool, error)
	GetNetwork() (datatypes.Network, error)
	GetNetworkComponentTrunks() (datatypes.Network_Component_Network_Vlan_Trunk, error)
	GetNetworkComponents() (datatypes.Network_Component, error)
	GetNetworkSpace() (string, error)
	GetNetworkVlanFirewall() (datatypes.Network_Vlan_Firewall, error)
	GetPrimaryRouter() (datatypes.Hardware_Router, error)
	GetPrimarySubnet() (datatypes.Network_Subnet, error)
	GetPrimarySubnetVersion6() (datatypes.Network_Subnet, error)
	GetPrimarySubnets() (datatypes.Network_Subnet, error)
	GetPrivateNetworkGateways() (datatypes.Network_Gateway, error)
	GetProtectedIpAddresses() (datatypes.Network_Subnet_IpAddress, error)
	GetPublicNetworkGateways() (datatypes.Network_Gateway, error)
	GetResourceGroupMember() (datatypes.Resource_Group_Member, error)
	GetResourceGroups() (datatypes.Resource_Group, error)
	GetSanStorageCapabilityFlag() (bool, error)
	GetScaleVlans() (datatypes.Scale_Network_Vlan, error)
	GetSecondaryRouter() (datatypes.Hardware, error)
	GetSecondarySubnets() (datatypes.Network_Subnet, error)
	GetSubnets() (datatypes.Network_Subnet, error)
	GetTagReferences() (datatypes.Tag_Reference, error)
	GetTotalPrimaryIpAddressCount() (uint, error)
	GetType() (datatypes.Network_Vlan_Type, error)
	GetVirtualGuests() (datatypes.Virtual_Guest, error)
}

type Network_Vlan_Firewall interface {
	Entity

	GetObject() (datatypes.Network_Vlan_Firewall, error)
	RestoreDefaults() (datatypes.Provisioning_Version1_Transaction, error)
	SetTags(tags string) (bool, error)
	UpdateRouteBypass(bypass bool) (datatypes.Provisioning_Version1_Transaction, error)

	GetBillingItem() (datatypes.Billing_Item, error)
	GetDatacenter() (datatypes.Location, error)
	GetFirewallType() (string, error)
	GetFullyQualifiedDomainName() (string, error)
	GetManagementCredentials() (datatypes.Software_Component_Password, error)
	GetNetworkFirewallUpdateRequests() (datatypes.Network_Firewall_Update_Request, error)
	GetNetworkVlan() (datatypes.Network_Vlan, error)
	GetNetworkVlans() (datatypes.Network_Vlan, error)
	GetRules() (datatypes.Network_Vlan_Firewall_Rule, error)
	GetTagReferences() (datatypes.Tag_Reference, error)
}

type Network_Vlan_Firewall_Rule interface {
	Entity

	GetNetworkComponentFirewall() (datatypes.Network_Component_Firewall, error)
}

type Network_Vlan_Type interface {
	Entity

	GetObject() (datatypes.Network_Vlan_Type, error)
}
