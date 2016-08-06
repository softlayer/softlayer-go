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

type Network struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkService() Network {
	return Network{Session: r}
}

func (r *Network) CreateObject(templateObject *datatypes.Network) (resp datatypes.Network, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network) CreateSubnet(subnet *datatypes.Network_Subnet, pod *datatypes.Network_Pod) (resp datatypes.Network_Subnet, err error) {
	params := []interface{}{
		subnet,
		pod,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network) DeleteSubnet(subnet *datatypes.Network_Subnet) (resp bool, err error) {
	params := []interface{}{
		subnet,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network) EditObject(templateObject *datatypes.Network) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network) GetAllObjects() (resp []datatypes.Network, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network) GetCidr() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network) GetName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network) GetNetworkIdentifier() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network) GetNotes() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network) GetObject() (resp datatypes.Network, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network) GetSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Application_Delivery_Controller struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkApplicationDeliveryControllerService() Network_Application_Delivery_Controller {
	return Network_Application_Delivery_Controller{Session: r}
}

func (r *Network_Application_Delivery_Controller) CreateLiveLoadBalancer(loadBalancer *datatypes.Network_LoadBalancer_VirtualIpAddress) (resp bool, err error) {
	params := []interface{}{
		loadBalancer,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) DeleteLiveLoadBalancer(loadBalancer *datatypes.Network_LoadBalancer_VirtualIpAddress) (resp bool, err error) {
	params := []interface{}{
		loadBalancer,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) DeleteLiveLoadBalancerService(service *datatypes.Network_LoadBalancer_Service) (resp bool, err error) {
	params := []interface{}{
		service,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) EditObject(templateObject *datatypes.Network_Application_Delivery_Controller) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) GetAverageDailyPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) GetBandwidthDataByDate(startDateTime *datatypes.Time, endDateTime *datatypes.Time, networkType *string) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
		networkType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) GetBandwidthImageByDate(startDateTime *datatypes.Time, endDateTime *datatypes.Time, networkType *string) (resp datatypes.Container_Bandwidth_GraphOutputs, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
		networkType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) GetBillingItem() (resp datatypes.Billing_Item_Network_Application_Delivery_Controller, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) GetConfigurationHistory() (resp []datatypes.Network_Application_Delivery_Controller_Configuration_History, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) GetCustomBandwidthDataByDate(graphData *datatypes.Container_Graph) (resp datatypes.Container_Graph, err error) {
	params := []interface{}{
		graphData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) GetDatacenter() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) GetDescription() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) GetLicenseExpirationDate() (resp datatypes.Time, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) GetLiveLoadBalancerServiceGraphImage(service *datatypes.Network_LoadBalancer_Service, graphType *string, metric *string) (resp []byte, err error) {
	params := []interface{}{
		service,
		graphType,
		metric,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) GetLoadBalancers() (resp []datatypes.Network_LoadBalancer_VirtualIpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) GetManagedResourceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) GetManagementIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) GetNetworkVlan() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) GetNetworkVlans() (resp []datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) GetObject() (resp datatypes.Network_Application_Delivery_Controller, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) GetOutboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) GetPassword() (resp datatypes.Software_Component_Password, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) GetPrimaryIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) GetProjectedPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) GetSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) GetTagReferences() (resp []datatypes.Tag_Reference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) GetType() (resp datatypes.Network_Application_Delivery_Controller_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) GetVirtualIpAddresses() (resp []datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) RestoreBaseConfiguration() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) RestoreConfiguration(configurationHistoryId *int) (resp bool, err error) {
	params := []interface{}{
		configurationHistoryId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) SaveCurrentConfiguration(notes *string) (resp datatypes.Network_Application_Delivery_Controller_Configuration_History, err error) {
	params := []interface{}{
		notes,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) UpdateLiveLoadBalancer(loadBalancer *datatypes.Network_LoadBalancer_VirtualIpAddress) (resp bool, err error) {
	params := []interface{}{
		loadBalancer,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller) UpdateNetScalerLicense() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Application_Delivery_Controller_Configuration_History struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkApplicationDeliveryControllerConfigurationHistoryService() Network_Application_Delivery_Controller_Configuration_History {
	return Network_Application_Delivery_Controller_Configuration_History{Session: r}
}

func (r *Network_Application_Delivery_Controller_Configuration_History) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_Configuration_History) GetController() (resp datatypes.Network_Application_Delivery_Controller, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_Configuration_History) GetObject() (resp datatypes.Network_Application_Delivery_Controller_Configuration_History, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkApplicationDeliveryControllerLoadBalancerHealthAttributeService() Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute {
	return Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute{Session: r}
}

func (r *Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute) GetHealthCheck() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Check, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute) GetObject() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute) GetType() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkApplicationDeliveryControllerLoadBalancerHealthAttributeTypeService() Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type {
	return Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type{Session: r}
}

func (r *Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type) GetAllObjects() (resp []datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type) GetObject() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Application_Delivery_Controller_LoadBalancer_Health_Check struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkApplicationDeliveryControllerLoadBalancerHealthCheckService() Network_Application_Delivery_Controller_LoadBalancer_Health_Check {
	return Network_Application_Delivery_Controller_LoadBalancer_Health_Check{Session: r}
}

func (r *Network_Application_Delivery_Controller_LoadBalancer_Health_Check) GetAttributes() (resp []datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Health_Check) GetObject() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Check, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Health_Check) GetScaleLoadBalancers() (resp []datatypes.Scale_LoadBalancer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Health_Check) GetServices() (resp []datatypes.Network_Application_Delivery_Controller_LoadBalancer_Service, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Health_Check) GetType() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkApplicationDeliveryControllerLoadBalancerHealthCheckTypeService() Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type {
	return Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type{Session: r}
}

func (r *Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type) GetAllObjects() (resp []datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type) GetObject() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Application_Delivery_Controller_LoadBalancer_Routing_Method struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkApplicationDeliveryControllerLoadBalancerRoutingMethodService() Network_Application_Delivery_Controller_LoadBalancer_Routing_Method {
	return Network_Application_Delivery_Controller_LoadBalancer_Routing_Method{Session: r}
}

func (r *Network_Application_Delivery_Controller_LoadBalancer_Routing_Method) GetAllObjects() (resp []datatypes.Network_Application_Delivery_Controller_LoadBalancer_Routing_Method, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Routing_Method) GetObject() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_Routing_Method, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Application_Delivery_Controller_LoadBalancer_Routing_Type struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkApplicationDeliveryControllerLoadBalancerRoutingTypeService() Network_Application_Delivery_Controller_LoadBalancer_Routing_Type {
	return Network_Application_Delivery_Controller_LoadBalancer_Routing_Type{Session: r}
}

func (r *Network_Application_Delivery_Controller_LoadBalancer_Routing_Type) GetAllObjects() (resp []datatypes.Network_Application_Delivery_Controller_LoadBalancer_Routing_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Routing_Type) GetObject() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_Routing_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Application_Delivery_Controller_LoadBalancer_Service struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkApplicationDeliveryControllerLoadBalancerServiceService() Network_Application_Delivery_Controller_LoadBalancer_Service {
	return Network_Application_Delivery_Controller_LoadBalancer_Service{Session: r}
}

func (r *Network_Application_Delivery_Controller_LoadBalancer_Service) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Service) GetGraphImage(graphType *string, metric *string) (resp []byte, err error) {
	params := []interface{}{
		graphType,
		metric,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Service) GetGroupReferences() (resp []datatypes.Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Service) GetGroups() (resp []datatypes.Network_Application_Delivery_Controller_LoadBalancer_Service_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Service) GetHealthCheck() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Check, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Service) GetHealthChecks() (resp []datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Check, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Service) GetIpAddress() (resp datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Service) GetObject() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_Service, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Service) GetServiceGroup() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_Service_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Service) ToggleStatus() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Application_Delivery_Controller_LoadBalancer_Service_Group struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkApplicationDeliveryControllerLoadBalancerServiceGroupService() Network_Application_Delivery_Controller_LoadBalancer_Service_Group {
	return Network_Application_Delivery_Controller_LoadBalancer_Service_Group{Session: r}
}

func (r *Network_Application_Delivery_Controller_LoadBalancer_Service_Group) GetGraphImage(graphType *string, metric *string) (resp []byte, err error) {
	params := []interface{}{
		graphType,
		metric,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Service_Group) GetObject() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_Service_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Service_Group) GetRoutingMethod() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_Routing_Method, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Service_Group) GetRoutingType() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_Routing_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Service_Group) GetServiceReferences() (resp []datatypes.Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Service_Group) GetServices() (resp []datatypes.Network_Application_Delivery_Controller_LoadBalancer_Service, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Service_Group) GetVirtualServer() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualServer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Service_Group) GetVirtualServers() (resp []datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualServer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_Service_Group) KickAllConnections() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkApplicationDeliveryControllerLoadBalancerVirtualIpAddressService() Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress {
	return Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress{Session: r}
}

func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) EditObject(templateObject *datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) GetApplicationDeliveryController() (resp datatypes.Network_Application_Delivery_Controller, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) GetApplicationDeliveryControllers() (resp []datatypes.Network_Application_Delivery_Controller, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) GetAvailableSecureTransportCiphers() (resp []datatypes.Security_SecureTransportCipher, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) GetAvailableSecureTransportProtocols() (resp []datatypes.Security_SecureTransportProtocol, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) GetDedicatedBillingItem() (resp datatypes.Billing_Item_Network_LoadBalancer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) GetHighAvailabilityFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) GetIpAddress() (resp datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) GetLoadBalancerHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) GetManagedResourceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) GetObject() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) GetSecureTransportCiphers() (resp []datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportCipher, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) GetSecureTransportProtocols() (resp []datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportProtocol, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) GetSecurityCertificate() (resp datatypes.Security_Certificate, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) GetSecurityCertificateEntry() (resp datatypes.Security_Certificate_Entry, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) GetVirtualServers() (resp []datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualServer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) StartSsl() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) StopSsl() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) UpgradeConnectionLimit() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Application_Delivery_Controller_LoadBalancer_VirtualServer struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkApplicationDeliveryControllerLoadBalancerVirtualServerService() Network_Application_Delivery_Controller_LoadBalancer_VirtualServer {
	return Network_Application_Delivery_Controller_LoadBalancer_VirtualServer{Session: r}
}

func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualServer) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualServer) GetObject() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualServer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualServer) GetRoutingMethod() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_Routing_Method, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualServer) GetScaleLoadBalancers() (resp []datatypes.Scale_LoadBalancer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualServer) GetServiceGroups() (resp []datatypes.Network_Application_Delivery_Controller_LoadBalancer_Service_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualServer) GetVirtualIpAddress() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualServer) StartSsl() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Application_Delivery_Controller_LoadBalancer_VirtualServer) StopSsl() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Backbone struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkBackboneService() Network_Backbone {
	return Network_Backbone{Session: r}
}

func (r *Network_Backbone) GetAllBackbones() (resp []datatypes.Network_Backbone, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Backbone) GetBackbonesForLocationName(locationName *string) (resp []datatypes.Network_Backbone, err error) {
	params := []interface{}{
		locationName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Backbone) GetGraphImage() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Backbone) GetHealth() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Backbone) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Backbone) GetNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Backbone) GetObject() (resp datatypes.Network_Backbone, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Backbone_Location_Dependent struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkBackboneLocationDependentService() Network_Backbone_Location_Dependent {
	return Network_Backbone_Location_Dependent{Session: r}
}

func (r *Network_Backbone_Location_Dependent) GetAllObjects() (resp []datatypes.Network_Backbone_Location_Dependent, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Backbone_Location_Dependent) GetDependentLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Backbone_Location_Dependent) GetObject() (resp datatypes.Network_Backbone_Location_Dependent, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Backbone_Location_Dependent) GetSourceDependentsByName(locationName *string) (resp datatypes.Location, err error) {
	params := []interface{}{
		locationName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Backbone_Location_Dependent) GetSourceLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Bandwidth_Version1_Allotment struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkBandwidthVersion1AllotmentService() Network_Bandwidth_Version1_Allotment {
	return Network_Bandwidth_Version1_Allotment{Session: r}
}

func (r *Network_Bandwidth_Version1_Allotment) CreateObject(templateObject *datatypes.Network_Bandwidth_Version1_Allotment) (resp datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) EditObject(templateObject *datatypes.Network_Bandwidth_Version1_Allotment) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetActiveDetails() (resp []datatypes.Network_Bandwidth_Version1_Allotment_Detail, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetApplicationDeliveryControllers() (resp []datatypes.Network_Application_Delivery_Controller, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetAverageDailyPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetBackendBandwidthByHour(date *datatypes.Time) (resp []datatypes.Container_Network_Bandwidth_Version1_Usage, err error) {
	params := []interface{}{
		date,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetBackendBandwidthUse(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Network_Bandwidth_Version1_Usage_Detail, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetBandwidthForDateRange(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetBandwidthImage(networkType *string, snapshotRange *string, draw *bool, dateSpecified *datatypes.Time, dateSpecifiedEnd *datatypes.Time) (resp datatypes.Container_Bandwidth_GraphOutputs, err error) {
	params := []interface{}{
		networkType,
		snapshotRange,
		draw,
		dateSpecified,
		dateSpecifiedEnd,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetBareMetalInstances() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetBillingCycleBandwidthUsage() (resp []datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetBillingCyclePrivateBandwidthUsage() (resp datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetBillingCyclePublicBandwidthUsage() (resp datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetBillingCyclePublicUsageTotal() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetCurrentBandwidthSummary() (resp datatypes.Metric_Tracking_Object_Bandwidth_Summary, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetCustomBandwidthDataByDate(graphData *datatypes.Container_Graph) (resp datatypes.Container_Graph, err error) {
	params := []interface{}{
		graphData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetDetails() (resp []datatypes.Network_Bandwidth_Version1_Allotment_Detail, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetFrontendBandwidthByHour(date *datatypes.Time) (resp []datatypes.Container_Network_Bandwidth_Version1_Usage, err error) {
	params := []interface{}{
		date,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetFrontendBandwidthUse(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Network_Bandwidth_Version1_Usage_Detail, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetInboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetLocationGroup() (resp datatypes.Location_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetManagedBareMetalInstances() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetManagedHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetManagedVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetMetricTrackingObject() (resp datatypes.Metric_Tracking_Object_VirtualDedicatedRack, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetMetricTrackingObjectId() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetObject() (resp datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetOutboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetOverBandwidthAllocationFlag() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetPrivateNetworkOnlyHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetProjectedOverBandwidthAllocationFlag() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetProjectedPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetServiceProvider() (resp datatypes.Service_Provider, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetTotalBandwidthAllocated() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) GetVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) ReassignServers(templateObjects []datatypes.Hardware, newAllotmentId *int) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
		newAllotmentId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) RequestVdrCancellation() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) RequestVdrContentUpdates(hardwareToAdd []datatypes.Hardware, hardwareToRemove []datatypes.Hardware, cloudsToAdd []datatypes.Virtual_Guest, cloudsToRemove []datatypes.Virtual_Guest, optionalAllotmentId *int, adcToAdd []datatypes.Network_Application_Delivery_Controller, adcToRemove []datatypes.Network_Application_Delivery_Controller) (resp bool, err error) {
	params := []interface{}{
		hardwareToAdd,
		hardwareToRemove,
		cloudsToAdd,
		cloudsToRemove,
		optionalAllotmentId,
		adcToAdd,
		adcToRemove,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) SetVdrContent(hardware []datatypes.Hardware, bareMetalServers []datatypes.Hardware, virtualServerInstance []datatypes.Virtual_Guest, adc []datatypes.Network_Application_Delivery_Controller, optionalAllotmentId *int) (resp bool, err error) {
	params := []interface{}{
		hardware,
		bareMetalServers,
		virtualServerInstance,
		adc,
		optionalAllotmentId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) UnassignServers(templateObjects []datatypes.Hardware) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) VoidPendingServerMove(id *int, typ *string) (resp bool, err error) {
	params := []interface{}{
		id,
		typ,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Bandwidth_Version1_Allotment) VoidPendingVdrCancellation() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Component struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkComponentService() Network_Component {
	return Network_Component{Session: r}
}

func (r *Network_Component) AddNetworkVlanTrunks(networkVlans []datatypes.Network_Vlan) (resp []datatypes.Network_Vlan, err error) {
	params := []interface{}{
		networkVlans,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) ClearNetworkVlanTrunks() (resp []datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetActiveCommand() (resp datatypes.Hardware_Component_RemoteManagement_Command_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetCustomBandwidthDataByDate(graphData *datatypes.Container_Graph) (resp datatypes.Container_Graph, err error) {
	params := []interface{}{
		graphData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetDownlinkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetDuplexMode() (resp datatypes.Network_Component_Duplex_Mode, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetHighAvailabilityFirewallFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetInterface() (resp datatypes.Network_Bandwidth_Version1_Interface, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetIpAddressBindings() (resp []datatypes.Network_Component_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetIpAddresses() (resp []datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetLastCommand() (resp datatypes.Hardware_Component_RemoteManagement_Command_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetMetricTrackingObject() (resp datatypes.Metric_Tracking_Object, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetNetworkComponentFirewall() (resp datatypes.Network_Component_Firewall, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetNetworkComponentGroup() (resp datatypes.Network_Component_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetNetworkVlan() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetNetworkVlanTrunks() (resp []datatypes.Network_Component_Network_Vlan_Trunk, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetObject() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetPortStatistics() (resp datatypes.Container_Network_Port_Statistic, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetPrimaryIpAddressRecord() (resp datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetPrimarySubnet() (resp datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetPrimaryVersion6IpAddressRecord() (resp datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetRecentCommands() (resp []datatypes.Hardware_Component_RemoteManagement_Command_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetRedundancyCapableFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetRedundancyEnabledFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetRemoteManagementUsers() (resp []datatypes.Hardware_Component_RemoteManagement_User, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetRouter() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetStorageNetworkFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetUplinkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) GetUplinkDuplexMode() (resp datatypes.Network_Component_Duplex_Mode, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component) RemoveNetworkVlanTrunks(networkVlans []datatypes.Network_Vlan) (resp []datatypes.Network_Vlan, err error) {
	params := []interface{}{
		networkVlans,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_Component_Firewall struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkComponentFirewallService() Network_Component_Firewall {
	return Network_Component_Firewall{Session: r}
}

func (r *Network_Component_Firewall) GetApplyServerRuleSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component_Firewall) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component_Firewall) GetGuestNetworkComponent() (resp datatypes.Virtual_Guest_Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component_Firewall) GetNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component_Firewall) GetNetworkFirewallUpdateRequest() (resp []datatypes.Network_Firewall_Update_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component_Firewall) GetObject() (resp datatypes.Network_Component_Firewall, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component_Firewall) GetRules() (resp []datatypes.Network_Component_Firewall_Rule, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Component_Firewall) GetSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_ContentDelivery_Account struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkContentDeliveryAccountService() Network_ContentDelivery_Account {
	return Network_ContentDelivery_Account{Session: r}
}

func (r *Network_ContentDelivery_Account) AuthenticateResourceRequest(parameter *datatypes.Container_Network_ContentDelivery_Authentication_Parameter) (resp bool, err error) {
	params := []interface{}{
		parameter,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) CreateDirectory(directoryName *string) (resp bool, err error) {
	params := []interface{}{
		directoryName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) CreateFtpUser(newPassword *string) (resp bool, err error) {
	params := []interface{}{
		newPassword,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) CreateOriginPullMapping(mappingObject *datatypes.Container_Network_ContentDelivery_OriginPull_Mapping) (resp bool, err error) {
	params := []interface{}{
		mappingObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) CreateOriginPullRule(originDomain *string, cnameRecord *string) (resp bool, err error) {
	params := []interface{}{
		originDomain,
		cnameRecord,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) CreateTokenAuthenticationDirectory(directory *string, mediaType *string) (resp bool, err error) {
	params := []interface{}{
		directory,
		mediaType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) DeleteFtpUser() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) DeleteOriginPullRule(originMappingId *string) (resp bool, err error) {
	params := []interface{}{
		originMappingId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) DisableLogging() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) EnableLogging() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetAllPopsBandwidthData(beginDateTime *datatypes.Time, endDateTime *datatypes.Time) (resp []datatypes.Container_Network_ContentDelivery_Bandwidth_PointsOfPresence_Summary, err error) {
	params := []interface{}{
		beginDateTime,
		endDateTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetAllPopsBandwidthImage(title *string, beginDateTime *datatypes.Time, endDateTime *datatypes.Time, unit *string) (resp datatypes.Container_Bandwidth_GraphOutputsExtended, err error) {
	params := []interface{}{
		title,
		beginDateTime,
		endDateTime,
		unit,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetAssociatedCdnAccountId() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetAuthenticationIpAddresses() (resp []datatypes.Network_ContentDelivery_Authentication_Address, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetAuthenticationServiceEndpoints() (resp []datatypes.Container_Network_ContentDelivery_Authentication_ServiceEndpoint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetBandwidthData(beginDateTime *datatypes.Time, endDateTime *datatypes.Time) (resp []datatypes.Container_Network_ContentDelivery_Bandwidth_Summary, err error) {
	params := []interface{}{
		beginDateTime,
		endDateTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetBandwidthDataWithTypes(beginDateTime *datatypes.Time, endDateTime *datatypes.Time, period *string) (resp []datatypes.Container_Network_ContentDelivery_Report_Usage, err error) {
	params := []interface{}{
		beginDateTime,
		endDateTime,
		period,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetBandwidthImage(title *string, beginDateTime *datatypes.Time, endDateTime *datatypes.Time) (resp datatypes.Container_Bandwidth_GraphOutputsExtended, err error) {
	params := []interface{}{
		title,
		beginDateTime,
		endDateTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetCdnAccountName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetCdnAccountNote() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetCdnSolutionName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetCustomerOrigins(mediaType *string) (resp []datatypes.Container_Network_ContentDelivery_OriginPull_Mapping, err error) {
	params := []interface{}{
		mediaType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetDependantServiceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetDirectoryInformation(directoryName *string) (resp []datatypes.Container_Network_Directory_Listing, err error) {
	params := []interface{}{
		directoryName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetDiskSpaceUsageDataByDate(beginDateTime *datatypes.Time, endDateTime *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		beginDateTime,
		endDateTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetDiskSpaceUsageImageByDate(beginDateTime *datatypes.Time, endDateTime *datatypes.Time) (resp datatypes.Container_Bandwidth_GraphOutputs, err error) {
	params := []interface{}{
		beginDateTime,
		endDateTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetFtpAttributes() (resp datatypes.Container_Network_Authentication_Data, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetLegacyCdnFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetLogEnabledFlag() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetMediaUrls() (resp []datatypes.Container_Network_ContentDelivery_SupportedProtocol, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetObject() (resp datatypes.Network_ContentDelivery_Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetOriginPullMappingInformation() (resp []datatypes.Container_Network_ContentDelivery_OriginPull_Mapping, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetOriginPullSupportedMediaUrls() (resp []datatypes.Container_Network_ContentDelivery_SupportedProtocol, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetOriginPullUrl() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetPopNames() (resp []datatypes.Container_Network_ContentDelivery_PointsOfPresence, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetProviderPortalAccessFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetProviderPortalCredentials() (resp datatypes.Container_Network_Authentication_Data, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetStatus() (resp datatypes.Network_ContentDelivery_Account_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetTokenAuthenticationDirectories() (resp []datatypes.Container_Network_Directory_Listing, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetTokenAuthenticationEnabledFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) GetVendorFtpAttributes() (resp datatypes.Container_Network_Authentication_Data, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) LoadContent(objectUrls []string) (resp bool, err error) {
	params := []interface{}{
		objectUrls,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) ManageHttpCompression(enableFlag *bool, mimeTypes []string) (resp bool, err error) {
	params := []interface{}{
		enableFlag,
		mimeTypes,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) PurgeCache(objectUrls []string) (resp []datatypes.Container_Network_ContentDelivery_PurgeService_Response, err error) {
	params := []interface{}{
		objectUrls,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) RemoveAuthenticationDirectory(directory *string, mediaType *string) (resp bool, err error) {
	params := []interface{}{
		directory,
		mediaType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) RemoveFile(source *string) (resp bool, err error) {
	params := []interface{}{
		source,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) SetAuthenticationServiceEndpoint(webserviceEndpoint *string, protocol *string) (resp bool, err error) {
	params := []interface{}{
		webserviceEndpoint,
		protocol,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) SetFtpPassword(newPassword *string) (resp bool, err error) {
	params := []interface{}{
		newPassword,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) UpdateNote(note *string) (resp bool, err error) {
	params := []interface{}{
		note,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Account) UploadStream(source *datatypes.Container_Utility_File_Attachment, target *string) (resp bool, err error) {
	params := []interface{}{
		source,
		target,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_ContentDelivery_Authentication_Address struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkContentDeliveryAuthenticationAddressService() Network_ContentDelivery_Authentication_Address {
	return Network_ContentDelivery_Authentication_Address{Session: r}
}

func (r *Network_ContentDelivery_Authentication_Address) CreateObject(templateObject *datatypes.Network_ContentDelivery_Authentication_Address) (resp datatypes.Network_ContentDelivery_Authentication_Address, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Authentication_Address) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Authentication_Address) EditObject(templateObject *datatypes.Network_ContentDelivery_Authentication_Address) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Authentication_Address) GetObject() (resp datatypes.Network_ContentDelivery_Authentication_Address, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Authentication_Address) RearrangeAuthenticationIp(cdnAccountId *int, templateObjects []datatypes.Network_ContentDelivery_Authentication_Address) (resp bool, err error) {
	params := []interface{}{
		cdnAccountId,
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_ContentDelivery_Authentication_Token struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkContentDeliveryAuthenticationTokenService() Network_ContentDelivery_Authentication_Token {
	return Network_ContentDelivery_Authentication_Token{Session: r}
}

func (r *Network_ContentDelivery_Authentication_Token) CreateObject(templateObject *datatypes.Network_ContentDelivery_Authentication_Token) (resp datatypes.Network_ContentDelivery_Authentication_Token, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Authentication_Token) GetAllManagedTokens(cdnAccountId *int) (resp []datatypes.Network_ContentDelivery_Authentication_Token, err error) {
	params := []interface{}{
		cdnAccountId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Authentication_Token) GetObject() (resp datatypes.Network_ContentDelivery_Authentication_Token, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Authentication_Token) GetTimedToken(cdnAccountId *int, tokenLife *int, clientIp *string, referrer *string, mediaType *string) (resp string, err error) {
	params := []interface{}{
		cdnAccountId,
		tokenLife,
		clientIp,
		referrer,
		mediaType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Authentication_Token) RevokeAllManagedTokens(cdnAccountId *int) (resp bool, err error) {
	params := []interface{}{
		cdnAccountId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Authentication_Token) RevokeAllTokens(cdnAccountId *int, mediaType *string) (resp bool, err error) {
	params := []interface{}{
		cdnAccountId,
		mediaType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Authentication_Token) RevokeManagedToken(cdnAccountId *int, token *string) (resp bool, err error) {
	params := []interface{}{
		cdnAccountId,
		token,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_ContentDelivery_Authentication_Token) RevokeManagedTokens(templateObjects []datatypes.Network_ContentDelivery_Authentication_Token) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_Customer_Subnet struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkCustomerSubnetService() Network_Customer_Subnet {
	return Network_Customer_Subnet{Session: r}
}

func (r *Network_Customer_Subnet) CreateObject(templateObject *datatypes.Network_Customer_Subnet) (resp datatypes.Network_Customer_Subnet, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Customer_Subnet) GetIpAddresses() (resp []datatypes.Network_Customer_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Customer_Subnet) GetObject() (resp datatypes.Network_Customer_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Firewall_AccessControlList struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkFirewallAccessControlListService() Network_Firewall_AccessControlList {
	return Network_Firewall_AccessControlList{Session: r}
}

func (r *Network_Firewall_AccessControlList) GetNetworkFirewallUpdateRequests() (resp []datatypes.Network_Firewall_Update_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Firewall_AccessControlList) GetNetworkVlan() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Firewall_AccessControlList) GetObject() (resp datatypes.Network_Firewall_AccessControlList, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Firewall_AccessControlList) GetRules() (resp []datatypes.Network_Vlan_Firewall_Rule, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Firewall_Interface struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkFirewallInterfaceService() Network_Firewall_Interface {
	return Network_Firewall_Interface{Session: r}
}

func (r *Network_Firewall_Interface) GetFirewallContextAccessControlLists() (resp []datatypes.Network_Firewall_AccessControlList, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Firewall_Interface) GetNetworkVlan() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Firewall_Interface) GetObject() (resp datatypes.Network_Firewall_Interface, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Firewall_Module_Context_Interface struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkFirewallModuleContextInterfaceService() Network_Firewall_Module_Context_Interface {
	return Network_Firewall_Module_Context_Interface{Session: r}
}

func (r *Network_Firewall_Module_Context_Interface) GetFirewallContextAccessControlLists() (resp []datatypes.Network_Firewall_AccessControlList, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Firewall_Module_Context_Interface) GetNetworkVlan() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Firewall_Module_Context_Interface) GetObject() (resp datatypes.Network_Firewall_Module_Context_Interface, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Firewall_Template struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkFirewallTemplateService() Network_Firewall_Template {
	return Network_Firewall_Template{Session: r}
}

func (r *Network_Firewall_Template) GetAllObjects() (resp []datatypes.Network_Firewall_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Firewall_Template) GetObject() (resp datatypes.Network_Firewall_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Firewall_Template) GetRules() (resp []datatypes.Network_Firewall_Template_Rule, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Firewall_Update_Request struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkFirewallUpdateRequestService() Network_Firewall_Update_Request {
	return Network_Firewall_Update_Request{Session: r}
}

func (r *Network_Firewall_Update_Request) CreateObject(templateObject *datatypes.Network_Firewall_Update_Request) (resp datatypes.Network_Firewall_Update_Request, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Firewall_Update_Request) GetAuthorizingUser() (resp datatypes.User_Interface, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Firewall_Update_Request) GetFirewallUpdateRequestRuleAttributes() (resp datatypes.Container_Utility_Network_Firewall_Rule_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Firewall_Update_Request) GetGuest() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Firewall_Update_Request) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Firewall_Update_Request) GetNetworkComponentFirewall() (resp datatypes.Network_Component_Firewall, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Firewall_Update_Request) GetObject() (resp datatypes.Network_Firewall_Update_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Firewall_Update_Request) GetRules() (resp []datatypes.Network_Firewall_Update_Request_Rule, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Firewall_Update_Request) UpdateRuleNote(fwRule *datatypes.Network_Component_Firewall_Rule, note *string) (resp bool, err error) {
	params := []interface{}{
		fwRule,
		note,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_Firewall_Update_Request_Rule struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkFirewallUpdateRequestRuleService() Network_Firewall_Update_Request_Rule {
	return Network_Firewall_Update_Request_Rule{Session: r}
}

func (r *Network_Firewall_Update_Request_Rule) CreateObject(templateObject *datatypes.Network_Firewall_Update_Request_Rule) (resp datatypes.Network_Firewall_Update_Request_Rule, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Firewall_Update_Request_Rule) GetFirewallUpdateRequest() (resp datatypes.Network_Firewall_Update_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Firewall_Update_Request_Rule) GetObject() (resp datatypes.Network_Firewall_Update_Request_Rule, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Firewall_Update_Request_Rule) ValidateRule(rule *datatypes.Network_Firewall_Update_Request_Rule, applyToComponentId *int, applyToAclId *int) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		rule,
		applyToComponentId,
		applyToAclId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_Gateway struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkGatewayService() Network_Gateway {
	return Network_Gateway{Session: r}
}

func (r *Network_Gateway) BypassAllVlans() (err error) {
	var resp datatypes.Void
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway) BypassVlans(vlans []datatypes.Network_Gateway_Vlan) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		vlans,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway) CreateObject(templateObject *datatypes.Network_Gateway) (resp datatypes.Network_Gateway, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway) EditObject(templateObject *datatypes.Network_Gateway) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway) GetInsideVlans() (resp []datatypes.Network_Gateway_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway) GetMembers() (resp []datatypes.Network_Gateway_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway) GetObject() (resp datatypes.Network_Gateway, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway) GetPossibleInsideVlans() (resp []datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway) GetPrivateIpAddress() (resp datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway) GetPrivateVlan() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway) GetPublicIpAddress() (resp datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway) GetPublicIpv6Address() (resp datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway) GetPublicVlan() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway) GetStatus() (resp datatypes.Network_Gateway_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway) UnbypassAllVlans() (err error) {
	var resp datatypes.Void
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway) UnbypassVlans(vlans []datatypes.Network_Gateway_Vlan) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		vlans,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_Gateway_Member struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkGatewayMemberService() Network_Gateway_Member {
	return Network_Gateway_Member{Session: r}
}

func (r *Network_Gateway_Member) CreateObject(templateObject *datatypes.Network_Gateway_Member) (resp datatypes.Network_Gateway_Member, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway_Member) CreateObjects(templateObjects []datatypes.Network_Gateway_Member) (resp []datatypes.Network_Gateway_Member, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway_Member) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway_Member) GetNetworkGateway() (resp datatypes.Network_Gateway, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway_Member) GetObject() (resp datatypes.Network_Gateway_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Gateway_Status struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkGatewayStatusService() Network_Gateway_Status {
	return Network_Gateway_Status{Session: r}
}

func (r *Network_Gateway_Status) GetObject() (resp datatypes.Network_Gateway_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Gateway_Vlan struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkGatewayVlanService() Network_Gateway_Vlan {
	return Network_Gateway_Vlan{Session: r}
}

func (r *Network_Gateway_Vlan) Bypass() (err error) {
	var resp datatypes.Void
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway_Vlan) CreateObject(templateObject *datatypes.Network_Gateway_Vlan) (resp datatypes.Network_Gateway_Vlan, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway_Vlan) CreateObjects(templateObjects []datatypes.Network_Gateway_Vlan) (resp []datatypes.Network_Gateway_Vlan, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway_Vlan) DeleteObject() (err error) {
	var resp datatypes.Void
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway_Vlan) DeleteObjects(templateObjects []datatypes.Network_Gateway_Vlan) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway_Vlan) GetNetworkGateway() (resp datatypes.Network_Gateway, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway_Vlan) GetNetworkVlan() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway_Vlan) GetObject() (resp datatypes.Network_Gateway_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Gateway_Vlan) Unbypass() (err error) {
	var resp datatypes.Void
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_LoadBalancer_Global_Account struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkLoadBalancerGlobalAccountService() Network_LoadBalancer_Global_Account {
	return Network_LoadBalancer_Global_Account{Session: r}
}

func (r *Network_LoadBalancer_Global_Account) AddNsRecord() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_LoadBalancer_Global_Account) EditObject(templateObject *datatypes.Network_LoadBalancer_Global_Account) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_LoadBalancer_Global_Account) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_LoadBalancer_Global_Account) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_LoadBalancer_Global_Account) GetHosts() (resp []datatypes.Network_LoadBalancer_Global_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_LoadBalancer_Global_Account) GetLoadBalanceType() (resp datatypes.Network_LoadBalancer_Global_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_LoadBalancer_Global_Account) GetManagedResourceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_LoadBalancer_Global_Account) GetObject() (resp datatypes.Network_LoadBalancer_Global_Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_LoadBalancer_Global_Account) RemoveNsRecord() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_LoadBalancer_Global_Host struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkLoadBalancerGlobalHostService() Network_LoadBalancer_Global_Host {
	return Network_LoadBalancer_Global_Host{Session: r}
}

func (r *Network_LoadBalancer_Global_Host) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_LoadBalancer_Global_Host) GetLoadBalancerAccount() (resp datatypes.Network_LoadBalancer_Global_Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_LoadBalancer_Global_Host) GetObject() (resp datatypes.Network_LoadBalancer_Global_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_LoadBalancer_Service struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkLoadBalancerServiceService() Network_LoadBalancer_Service {
	return Network_LoadBalancer_Service{Session: r}
}

func (r *Network_LoadBalancer_Service) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_LoadBalancer_Service) GetGraphImage(graphType *string, metric *string) (resp []byte, err error) {
	params := []interface{}{
		graphType,
		metric,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_LoadBalancer_Service) GetObject() (resp datatypes.Network_LoadBalancer_Service, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_LoadBalancer_Service) GetStatus() (resp []datatypes.Container_Network_LoadBalancer_StatusEntry, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_LoadBalancer_Service) GetVip() (resp datatypes.Network_LoadBalancer_VirtualIpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_LoadBalancer_Service) ResetPeakConnections() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_LoadBalancer_VirtualIpAddress struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkLoadBalancerVirtualIpAddressService() Network_LoadBalancer_VirtualIpAddress {
	return Network_LoadBalancer_VirtualIpAddress{Session: r}
}

func (r *Network_LoadBalancer_VirtualIpAddress) Disable() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_LoadBalancer_VirtualIpAddress) EditObject(templateObject *datatypes.Network_LoadBalancer_VirtualIpAddress) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_LoadBalancer_VirtualIpAddress) Enable() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_LoadBalancer_VirtualIpAddress) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_LoadBalancer_VirtualIpAddress) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_LoadBalancer_VirtualIpAddress) GetCustomerManagedFlag() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_LoadBalancer_VirtualIpAddress) GetManagedResourceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_LoadBalancer_VirtualIpAddress) GetObject() (resp datatypes.Network_LoadBalancer_VirtualIpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_LoadBalancer_VirtualIpAddress) GetServices() (resp []datatypes.Network_LoadBalancer_Service, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_LoadBalancer_VirtualIpAddress) KickAllConnections() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_LoadBalancer_VirtualIpAddress) UpgradeConnectionLimit() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Media_Transcode_Account struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkMediaTranscodeAccountService() Network_Media_Transcode_Account {
	return Network_Media_Transcode_Account{Session: r}
}

func (r *Network_Media_Transcode_Account) CreateTranscodeAccount() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Media_Transcode_Account) CreateTranscodeJob(newJob *datatypes.Network_Media_Transcode_Job) (resp bool, err error) {
	params := []interface{}{
		newJob,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Media_Transcode_Account) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Media_Transcode_Account) GetDirectoryInformation(directoryName *string, extensionFilter *string) (resp []datatypes.Container_Network_Directory_Listing, err error) {
	params := []interface{}{
		directoryName,
		extensionFilter,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Media_Transcode_Account) GetFileDetail(source *string) (resp datatypes.Container_Network_Media_Information, err error) {
	params := []interface{}{
		source,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Media_Transcode_Account) GetFtpAttributes() (resp datatypes.Container_Network_Authentication_Data, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Media_Transcode_Account) GetObject() (resp datatypes.Network_Media_Transcode_Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Media_Transcode_Account) GetPresetDetail(guid *string) (resp []datatypes.Container_Network_Media_Transcode_Preset_Element, err error) {
	params := []interface{}{
		guid,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Media_Transcode_Account) GetPresets() (resp []datatypes.Container_Network_Media_Transcode_Preset, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Media_Transcode_Account) GetTranscodeJobs() (resp []datatypes.Network_Media_Transcode_Job, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Media_Transcode_Job struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkMediaTranscodeJobService() Network_Media_Transcode_Job {
	return Network_Media_Transcode_Job{Session: r}
}

func (r *Network_Media_Transcode_Job) CreateObject(templateObject *datatypes.Network_Media_Transcode_Job) (resp datatypes.Network_Media_Transcode_Job, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Media_Transcode_Job) GetHistory() (resp []datatypes.Network_Media_Transcode_Job_History, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Media_Transcode_Job) GetObject() (resp datatypes.Network_Media_Transcode_Job, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Media_Transcode_Job) GetTranscodeAccount() (resp datatypes.Network_Media_Transcode_Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Media_Transcode_Job) GetTranscodeStatus() (resp datatypes.Network_Media_Transcode_Job_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Media_Transcode_Job) GetTranscodeStatusName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Media_Transcode_Job) GetUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Media_Transcode_Job_Status struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkMediaTranscodeJobStatusService() Network_Media_Transcode_Job_Status {
	return Network_Media_Transcode_Job_Status{Session: r}
}

func (r *Network_Media_Transcode_Job_Status) GetAllStatuses() (resp []datatypes.Network_Media_Transcode_Job_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Media_Transcode_Job_Status) GetObject() (resp datatypes.Network_Media_Transcode_Job_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Message_Delivery struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkMessageDeliveryService() Network_Message_Delivery {
	return Network_Message_Delivery{Session: r}
}

func (r *Network_Message_Delivery) EditObject(templateObject *datatypes.Network_Message_Delivery) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Delivery) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Delivery) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Delivery) GetObject() (resp datatypes.Network_Message_Delivery, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Delivery) GetType() (resp datatypes.Network_Message_Delivery_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Delivery) GetVendor() (resp datatypes.Network_Message_Delivery_Vendor, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Message_Delivery_Email_Sendgrid struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkMessageDeliveryEmailSendgridService() Network_Message_Delivery_Email_Sendgrid {
	return Network_Message_Delivery_Email_Sendgrid{Session: r}
}

func (r *Network_Message_Delivery_Email_Sendgrid) AddUnsubscribeEmailAddress(emailAddress *string) (resp bool, err error) {
	params := []interface{}{
		emailAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Delivery_Email_Sendgrid) DeleteEmailListEntries(list *string, entries []string) (resp bool, err error) {
	params := []interface{}{
		list,
		entries,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Delivery_Email_Sendgrid) DisableSmtpAccess() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Delivery_Email_Sendgrid) EditObject(templateObject *datatypes.Network_Message_Delivery) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Delivery_Email_Sendgrid) EnableSmtpAccess() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Delivery_Email_Sendgrid) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Delivery_Email_Sendgrid) GetAccountOverview() (resp datatypes.Container_Network_Message_Delivery_Email_Sendgrid_Account_Overview, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Delivery_Email_Sendgrid) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Delivery_Email_Sendgrid) GetCategoryList() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Delivery_Email_Sendgrid) GetEmailAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Delivery_Email_Sendgrid) GetEmailList(list *string) (resp []datatypes.Container_Network_Message_Delivery_Email_Sendgrid_List_Entry, err error) {
	params := []interface{}{
		list,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Delivery_Email_Sendgrid) GetObject() (resp datatypes.Network_Message_Delivery_Email_Sendgrid, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Delivery_Email_Sendgrid) GetSmtpAccess() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Delivery_Email_Sendgrid) GetStatistics(options *datatypes.Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Options) (resp []datatypes.Container_Network_Message_Delivery_Email_Sendgrid_Statistics, err error) {
	params := []interface{}{
		options,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Delivery_Email_Sendgrid) GetStatisticsGraph(options *datatypes.Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Options) (resp datatypes.Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Graph, err error) {
	params := []interface{}{
		options,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Delivery_Email_Sendgrid) GetType() (resp datatypes.Network_Message_Delivery_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Delivery_Email_Sendgrid) GetVendor() (resp datatypes.Network_Message_Delivery_Vendor, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Delivery_Email_Sendgrid) GetVendorPortalUrl() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Delivery_Email_Sendgrid) SendEmail(emailContainer *datatypes.Container_Network_Message_Delivery_Email) (resp bool, err error) {
	params := []interface{}{
		emailContainer,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Delivery_Email_Sendgrid) UpdateEmailAddress(emailAddress *string) (resp bool, err error) {
	params := []interface{}{
		emailAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_Message_Queue struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkMessageQueueService() Network_Message_Queue {
	return Network_Message_Queue{Session: r}
}

func (r *Network_Message_Queue) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Queue) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Queue) GetNodes() (resp []datatypes.Network_Message_Queue_Node, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Queue) GetObject() (resp datatypes.Network_Message_Queue, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Queue) GetStatus() (resp datatypes.Network_Message_Queue_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Message_Queue_Node struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkMessageQueueNodeService() Network_Message_Queue_Node {
	return Network_Message_Queue_Node{Session: r}
}

func (r *Network_Message_Queue_Node) AddUser(username *string) (resp bool, err error) {
	params := []interface{}{
		username,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Queue_Node) DeleteUser(username *string) (resp bool, err error) {
	params := []interface{}{
		username,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Queue_Node) GetAllUsers() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Queue_Node) GetMessageQueue() (resp datatypes.Network_Message_Queue, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Queue_Node) GetMetricTrackingObject() (resp datatypes.Metric_Tracking_Object, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Queue_Node) GetObject() (resp datatypes.Network_Message_Queue_Node, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Queue_Node) GetServiceResource() (resp datatypes.Network_Service_Resource, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Queue_Node) GetUsage(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Message_Queue_Node) GetUsageGraph(graphData *datatypes.Container_Graph) (resp datatypes.Container_Graph, err error) {
	params := []interface{}{
		graphData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_Message_Queue_Status struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkMessageQueueStatusService() Network_Message_Queue_Status {
	return Network_Message_Queue_Status{Session: r}
}

func (r *Network_Message_Queue_Status) GetObject() (resp datatypes.Network_Message_Queue_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Monitor struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkMonitorService() Network_Monitor {
	return Network_Monitor{Session: r}
}

func (r *Network_Monitor) GetIpAddressesByHardware(hardware *datatypes.Hardware, partialIpAddress *string) (resp []datatypes.Network_Subnet_IpAddress, err error) {
	params := []interface{}{
		hardware,
		partialIpAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Monitor) GetIpAddressesByVirtualGuest(guest *datatypes.Virtual_Guest, partialIpAddress *string) (resp []datatypes.Network_Subnet_IpAddress, err error) {
	params := []interface{}{
		guest,
		partialIpAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_Monitor_Version1_Query_Host struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkMonitorVersion1QueryHostService() Network_Monitor_Version1_Query_Host {
	return Network_Monitor_Version1_Query_Host{Session: r}
}

func (r *Network_Monitor_Version1_Query_Host) CreateObject(templateObject *datatypes.Network_Monitor_Version1_Query_Host) (resp datatypes.Network_Monitor_Version1_Query_Host, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Monitor_Version1_Query_Host) CreateObjects(templateObjects []datatypes.Network_Monitor_Version1_Query_Host) (resp []datatypes.Network_Monitor_Version1_Query_Host, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Monitor_Version1_Query_Host) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Monitor_Version1_Query_Host) DeleteObjects(templateObjects []datatypes.Network_Monitor_Version1_Query_Host) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Monitor_Version1_Query_Host) EditObject(templateObject *datatypes.Network_Monitor_Version1_Query_Host) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Monitor_Version1_Query_Host) EditObjects(templateObjects []datatypes.Network_Monitor_Version1_Query_Host) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Monitor_Version1_Query_Host) FindByHardwareId(hardwareId *int) (resp []datatypes.Network_Monitor_Version1_Query_Host, err error) {
	params := []interface{}{
		hardwareId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Monitor_Version1_Query_Host) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Monitor_Version1_Query_Host) GetLastResult() (resp datatypes.Network_Monitor_Version1_Query_Result, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Monitor_Version1_Query_Host) GetObject() (resp datatypes.Network_Monitor_Version1_Query_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Monitor_Version1_Query_Host) GetQueryType() (resp datatypes.Network_Monitor_Version1_Query_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Monitor_Version1_Query_Host) GetResponseAction() (resp datatypes.Network_Monitor_Version1_Query_ResponseType, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Monitor_Version1_Query_Host_Stratum struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkMonitorVersion1QueryHostStratumService() Network_Monitor_Version1_Query_Host_Stratum {
	return Network_Monitor_Version1_Query_Host_Stratum{Session: r}
}

func (r *Network_Monitor_Version1_Query_Host_Stratum) GetAllQueryTypes() (resp []datatypes.Network_Monitor_Version1_Query_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Monitor_Version1_Query_Host_Stratum) GetAllResponseTypes() (resp []datatypes.Network_Monitor_Version1_Query_ResponseType, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Monitor_Version1_Query_Host_Stratum) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Monitor_Version1_Query_Host_Stratum) GetObject() (resp datatypes.Network_Monitor_Version1_Query_Host_Stratum, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Pod struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkPodService() Network_Pod {
	return Network_Pod{Session: r}
}

func (r *Network_Pod) GetAllObjects() (resp []datatypes.Network_Pod, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Pod) GetCapabilities() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Pod) GetObject() (resp datatypes.Network_Pod, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Pod) ListCapabilities() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Security_Scanner_Request struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkSecurityScannerRequestService() Network_Security_Scanner_Request {
	return Network_Security_Scanner_Request{Session: r}
}

func (r *Network_Security_Scanner_Request) CreateObject(templateObject *datatypes.Network_Security_Scanner_Request) (resp datatypes.Network_Security_Scanner_Request, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Security_Scanner_Request) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Security_Scanner_Request) GetGuest() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Security_Scanner_Request) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Security_Scanner_Request) GetObject() (resp datatypes.Network_Security_Scanner_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Security_Scanner_Request) GetReport() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Security_Scanner_Request) GetRequestorOwnedFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Security_Scanner_Request) GetStatus() (resp datatypes.Network_Security_Scanner_Request_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Service_Vpn_Overrides struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkServiceVpnOverridesService() Network_Service_Vpn_Overrides {
	return Network_Service_Vpn_Overrides{Session: r}
}

func (r *Network_Service_Vpn_Overrides) CreateObjects(templateObjects []datatypes.Network_Service_Vpn_Overrides) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Service_Vpn_Overrides) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Service_Vpn_Overrides) DeleteObjects(templateObjects []datatypes.Network_Service_Vpn_Overrides) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Service_Vpn_Overrides) GetObject() (resp datatypes.Network_Service_Vpn_Overrides, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Service_Vpn_Overrides) GetSubnet() (resp datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Service_Vpn_Overrides) GetUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Storage struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkStorageService() Network_Storage {
	return Network_Storage{Session: r}
}

func (r *Network_Storage) AllowAccessFromHardware(hardwareObjectTemplate *datatypes.Hardware) (resp bool, err error) {
	params := []interface{}{
		hardwareObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) AllowAccessFromHardwareList(hardwareObjectTemplates []datatypes.Hardware) (resp bool, err error) {
	params := []interface{}{
		hardwareObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) AllowAccessFromHost(typeClassName *string, hostId *int) (resp datatypes.Network_Storage_Allowed_Host, err error) {
	params := []interface{}{
		typeClassName,
		hostId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) AllowAccessFromHostList(hostObjectTemplates []datatypes.Container_Network_Storage_Host) (resp []datatypes.Network_Storage_Allowed_Host, err error) {
	params := []interface{}{
		hostObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) AllowAccessFromIpAddress(ipAddressObjectTemplate *datatypes.Network_Subnet_IpAddress) (resp bool, err error) {
	params := []interface{}{
		ipAddressObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) AllowAccessFromIpAddressList(ipAddressObjectTemplates []datatypes.Network_Subnet_IpAddress) (resp bool, err error) {
	params := []interface{}{
		ipAddressObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) AllowAccessFromSubnet(subnetObjectTemplate *datatypes.Network_Subnet) (resp bool, err error) {
	params := []interface{}{
		subnetObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) AllowAccessFromSubnetList(subnetObjectTemplates []datatypes.Network_Subnet) (resp bool, err error) {
	params := []interface{}{
		subnetObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) AllowAccessFromVirtualGuest(virtualGuestObjectTemplate *datatypes.Virtual_Guest) (resp bool, err error) {
	params := []interface{}{
		virtualGuestObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) AllowAccessFromVirtualGuestList(virtualGuestObjectTemplates []datatypes.Virtual_Guest) (resp bool, err error) {
	params := []interface{}{
		virtualGuestObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) AllowAccessToReplicantFromHardware(hardwareObjectTemplate *datatypes.Hardware) (resp bool, err error) {
	params := []interface{}{
		hardwareObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) AllowAccessToReplicantFromHardwareList(hardwareObjectTemplates []datatypes.Hardware) (resp bool, err error) {
	params := []interface{}{
		hardwareObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) AllowAccessToReplicantFromIpAddress(ipAddressObjectTemplate *datatypes.Network_Subnet_IpAddress) (resp bool, err error) {
	params := []interface{}{
		ipAddressObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) AllowAccessToReplicantFromIpAddressList(ipAddressObjectTemplates []datatypes.Network_Subnet_IpAddress) (resp bool, err error) {
	params := []interface{}{
		ipAddressObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) AllowAccessToReplicantFromSubnet(subnetObjectTemplate *datatypes.Network_Subnet) (resp bool, err error) {
	params := []interface{}{
		subnetObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) AllowAccessToReplicantFromSubnetList(subnetObjectTemplates []datatypes.Network_Subnet) (resp bool, err error) {
	params := []interface{}{
		subnetObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) AllowAccessToReplicantFromVirtualGuest(virtualGuestObjectTemplate *datatypes.Virtual_Guest) (resp bool, err error) {
	params := []interface{}{
		virtualGuestObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) AllowAccessToReplicantFromVirtualGuestList(virtualGuestObjectTemplates []datatypes.Virtual_Guest) (resp bool, err error) {
	params := []interface{}{
		virtualGuestObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) AssignCredential(username *string) (resp bool, err error) {
	params := []interface{}{
		username,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) AssignNewCredential(typ *string) (resp datatypes.Network_Storage_Credential, err error) {
	params := []interface{}{
		typ,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) ChangePassword(username *string, currentPassword *string, newPassword *string) (resp bool, err error) {
	params := []interface{}{
		username,
		currentPassword,
		newPassword,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) CollectBandwidth(typ *string, startDate *datatypes.Time, endDate *datatypes.Time) (resp uint, err error) {
	params := []interface{}{
		typ,
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) CollectBytesUsed() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) CreateFolder(folder *string) (resp bool, err error) {
	params := []interface{}{
		folder,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) CreateSnapshot(notes *string) (resp datatypes.Network_Storage, err error) {
	params := []interface{}{
		notes,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) DeleteAllFiles() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) DeleteFile(fileId *string) (resp bool, err error) {
	params := []interface{}{
		fileId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) DeleteFiles(fileIds []string) (resp bool, err error) {
	params := []interface{}{
		fileIds,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) DeleteFolder(folder *string) (resp bool, err error) {
	params := []interface{}{
		folder,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) DisableSnapshots(scheduleType *string) (resp bool, err error) {
	params := []interface{}{
		scheduleType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) DownloadFile(fileId *string) (resp datatypes.Container_Utility_File_Entity, err error) {
	params := []interface{}{
		fileId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) EditCredential(username *string, newPassword *string) (resp bool, err error) {
	params := []interface{}{
		username,
		newPassword,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) EditObject(templateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) EnableSnapshots(scheduleType *string, retentionCount *int, minute *int, hour *int, dayOfWeek *string) (resp bool, err error) {
	params := []interface{}{
		scheduleType,
		retentionCount,
		minute,
		hour,
		dayOfWeek,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) FailbackFromReplicant() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) FailoverToReplicant(replicantId *int) (resp bool, err error) {
	params := []interface{}{
		replicantId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetAccountPassword() (resp datatypes.Account_Password, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetActiveTransactions() (resp []datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetAllFiles() (resp []datatypes.Container_Utility_File_Entity, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetAllFilesByFilter(filter *datatypes.Container_Utility_File_Entity) (resp []datatypes.Container_Utility_File_Entity, err error) {
	params := []interface{}{
		filter,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetAllowableHardware(filterHostname *string) (resp []datatypes.Hardware, err error) {
	params := []interface{}{
		filterHostname,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetAllowableIpAddresses(subnetId *int, filterIpAddress *string) (resp []datatypes.Network_Subnet_IpAddress, err error) {
	params := []interface{}{
		subnetId,
		filterIpAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetAllowableSubnets(filterNetworkIdentifier *string) (resp []datatypes.Network_Subnet, err error) {
	params := []interface{}{
		filterNetworkIdentifier,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetAllowableVirtualGuests(filterHostname *string) (resp []datatypes.Virtual_Guest, err error) {
	params := []interface{}{
		filterHostname,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetAllowedHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetAllowedHostsLimit() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetAllowedIpAddresses() (resp []datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetAllowedReplicationHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetAllowedReplicationIpAddresses() (resp []datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetAllowedReplicationSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetAllowedReplicationVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetAllowedSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetAllowedVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetBillingItemCategory() (resp datatypes.Product_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetByUsername(username *string, typ *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		username,
		typ,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetBytesUsed() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetCdnUrls() (resp []datatypes.Container_Network_Storage_Hub_ObjectStorage_ContentDeliveryUrl, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetClusterResource() (resp datatypes.Network_Service_Resource, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetCreationScheduleId() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetCredentials() (resp []datatypes.Network_Storage_Credential, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetDailySchedule() (resp datatypes.Network_Storage_Schedule, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetEvents() (resp []datatypes.Network_Storage_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetFileByIdentifier(identifier *string) (resp datatypes.Container_Utility_File_Entity, err error) {
	params := []interface{}{
		identifier,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetFileCount() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetFileList(folder *string, path *string) (resp []datatypes.Container_Utility_File_Entity, err error) {
	params := []interface{}{
		folder,
		path,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetFilePendingDeleteCount() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetFilesPendingDelete() (resp []datatypes.Container_Utility_File_Entity, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetFolderList() (resp []datatypes.Container_Network_Storage_Hub_ObjectStorage_Folder, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetGraph(startDate *datatypes.Time, endDate *datatypes.Time, typ *string) (resp datatypes.Container_Bandwidth_GraphOutputs, err error) {
	params := []interface{}{
		startDate,
		endDate,
		typ,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetHourlySchedule() (resp datatypes.Network_Storage_Schedule, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetIops() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetIscsiLuns() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetLunId() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetManualSnapshots() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetMetricTrackingObject() (resp datatypes.Metric_Tracking_Object, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetMountableFlag() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetNetworkConnectionDetails() (resp datatypes.Container_Network_Storage_NetworkConnectionInformation, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetNotificationSubscribers() (resp []datatypes.Notification_User_Subscriber, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetObject() (resp datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetObjectStorageConnectionInformation() (resp []datatypes.Container_Network_Service_Resource_ObjectStorage_ConnectionInformation, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetObjectsByCredential(credentialObject *datatypes.Network_Storage_Credential) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		credentialObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetOsType() (resp datatypes.Network_Storage_Iscsi_OS_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetOsTypeId() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetParentPartnerships() (resp []datatypes.Network_Storage_Partnership, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetParentVolume() (resp datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetPartnerships() (resp []datatypes.Network_Storage_Partnership, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetPermissionsGroups() (resp []datatypes.Network_Storage_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetProperties() (resp []datatypes.Network_Storage_Property, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetRecycleBinFileByIdentifier(fileId *string) (resp datatypes.Container_Utility_File_Entity, err error) {
	params := []interface{}{
		fileId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetRemainingAllowedHosts() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetReplicatingLuns() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetReplicatingVolume() (resp datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetReplicationEvents() (resp []datatypes.Network_Storage_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetReplicationPartners() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetReplicationSchedule() (resp datatypes.Network_Storage_Schedule, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetReplicationStatus() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetSchedules() (resp []datatypes.Network_Storage_Schedule, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetServiceResource() (resp datatypes.Network_Service_Resource, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetServiceResourceBackendIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetServiceResourceName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetSnapshotCapacityGb() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetSnapshotCreationTimestamp() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetSnapshotDeletionThresholdPercentage() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetSnapshotSizeBytes() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetSnapshotSpaceAvailable() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetSnapshots() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetSnapshotsForVolume() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetStorageGroups() (resp []datatypes.Network_Storage_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetStorageGroupsNetworkConnectionDetails() (resp []datatypes.Container_Network_Storage_NetworkConnectionInformation, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetStorageTierLevel() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetStorageType() (resp datatypes.Network_Storage_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetTotalBytesUsed() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetTotalScheduleSnapshotRetentionCount() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetUsageNotification() (resp datatypes.Notification, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetValidReplicationTargetDatacenterLocations() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetVendorName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetVirtualGuest() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetVolumeHistory() (resp []datatypes.Network_Storage_History, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetVolumeStatus() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetWebccAccount() (resp datatypes.Account_Password, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) GetWeeklySchedule() (resp datatypes.Network_Storage_Schedule, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) ImmediateFailoverToReplicant(replicantId *int) (resp bool, err error) {
	params := []interface{}{
		replicantId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) IsBlockingOperationInProgress(exemptStatusKeyNames []string) (resp bool, err error) {
	params := []interface{}{
		exemptStatusKeyNames,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) RemoveAccessFromHardware(hardwareObjectTemplate *datatypes.Hardware) (resp bool, err error) {
	params := []interface{}{
		hardwareObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) RemoveAccessFromHardwareList(hardwareObjectTemplates []datatypes.Hardware) (resp bool, err error) {
	params := []interface{}{
		hardwareObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) RemoveAccessFromHost(typeClassName *string, hostId *int) (resp datatypes.Network_Storage_Allowed_Host, err error) {
	params := []interface{}{
		typeClassName,
		hostId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) RemoveAccessFromHostList(hostObjectTemplates []datatypes.Container_Network_Storage_Host) (resp []datatypes.Network_Storage_Allowed_Host, err error) {
	params := []interface{}{
		hostObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) RemoveAccessFromIpAddress(ipAddressObjectTemplate *datatypes.Network_Subnet_IpAddress) (resp bool, err error) {
	params := []interface{}{
		ipAddressObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) RemoveAccessFromIpAddressList(ipAddressObjectTemplates []datatypes.Network_Subnet_IpAddress) (resp bool, err error) {
	params := []interface{}{
		ipAddressObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) RemoveAccessFromSubnet(subnetObjectTemplate *datatypes.Network_Subnet) (resp bool, err error) {
	params := []interface{}{
		subnetObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) RemoveAccessFromSubnetList(subnetObjectTemplates []datatypes.Network_Subnet) (resp bool, err error) {
	params := []interface{}{
		subnetObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) RemoveAccessFromVirtualGuest(virtualGuestObjectTemplate *datatypes.Virtual_Guest) (resp bool, err error) {
	params := []interface{}{
		virtualGuestObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) RemoveAccessFromVirtualGuestList(virtualGuestObjectTemplates []datatypes.Virtual_Guest) (resp bool, err error) {
	params := []interface{}{
		virtualGuestObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) RemoveAccessToReplicantFromHardwareList(hardwareObjectTemplates []datatypes.Hardware) (resp bool, err error) {
	params := []interface{}{
		hardwareObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) RemoveAccessToReplicantFromIpAddressList(ipAddressObjectTemplates []datatypes.Network_Subnet_IpAddress) (resp bool, err error) {
	params := []interface{}{
		ipAddressObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) RemoveAccessToReplicantFromSubnet(subnetObjectTemplate *datatypes.Network_Subnet) (resp bool, err error) {
	params := []interface{}{
		subnetObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) RemoveAccessToReplicantFromSubnetList(subnetObjectTemplates []datatypes.Network_Subnet) (resp bool, err error) {
	params := []interface{}{
		subnetObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) RemoveAccessToReplicantFromVirtualGuestList(virtualGuestObjectTemplates []datatypes.Virtual_Guest) (resp bool, err error) {
	params := []interface{}{
		virtualGuestObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) RemoveCredential(username *string) (resp bool, err error) {
	params := []interface{}{
		username,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) RestoreFile(fileId *string) (resp datatypes.Container_Utility_File_Entity, err error) {
	params := []interface{}{
		fileId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) RestoreFromSnapshot(snapshotId *int) (resp bool, err error) {
	params := []interface{}{
		snapshotId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) SendPasswordReminderEmail(username *string) (resp bool, err error) {
	params := []interface{}{
		username,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) SetMountable(mountable *bool) (resp bool, err error) {
	params := []interface{}{
		mountable,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) SetSnapshotAllocation(capacityGb *int) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		capacityGb,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) UpgradeVolumeCapacity(itemId *int) (resp bool, err error) {
	params := []interface{}{
		itemId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage) UploadFile(file *datatypes.Container_Utility_File_Entity) (resp datatypes.Container_Utility_File_Entity, err error) {
	params := []interface{}{
		file,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_Storage_Allowed_Host struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkStorageAllowedHostService() Network_Storage_Allowed_Host {
	return Network_Storage_Allowed_Host{Session: r}
}

func (r *Network_Storage_Allowed_Host) CreateObject(templateObject *datatypes.Network_Storage_Allowed_Host) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host) EditObject(templateObject *datatypes.Network_Storage_Allowed_Host) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host) GetAssignedGroups() (resp []datatypes.Network_Storage_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host) GetAssignedReplicationVolumes() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host) GetAssignedVolumes() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host) GetCredential() (resp datatypes.Network_Storage_Credential, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host) GetObject() (resp datatypes.Network_Storage_Allowed_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host) SetCredentialPassword(password *string) (resp bool, err error) {
	params := []interface{}{
		password,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_Storage_Allowed_Host_Hardware struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkStorageAllowedHostHardwareService() Network_Storage_Allowed_Host_Hardware {
	return Network_Storage_Allowed_Host_Hardware{Session: r}
}

func (r *Network_Storage_Allowed_Host_Hardware) CreateObject(templateObject *datatypes.Network_Storage_Allowed_Host) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_Hardware) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_Hardware) EditObject(templateObject *datatypes.Network_Storage_Allowed_Host) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_Hardware) GetAssignedGroups() (resp []datatypes.Network_Storage_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_Hardware) GetAssignedReplicationVolumes() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_Hardware) GetAssignedVolumes() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_Hardware) GetCredential() (resp datatypes.Network_Storage_Credential, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_Hardware) GetObject() (resp datatypes.Network_Storage_Allowed_Host_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_Hardware) GetResource() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_Hardware) SetCredentialPassword(password *string) (resp bool, err error) {
	params := []interface{}{
		password,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_Storage_Allowed_Host_IpAddress struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkStorageAllowedHostIpAddressService() Network_Storage_Allowed_Host_IpAddress {
	return Network_Storage_Allowed_Host_IpAddress{Session: r}
}

func (r *Network_Storage_Allowed_Host_IpAddress) CreateObject(templateObject *datatypes.Network_Storage_Allowed_Host) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_IpAddress) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_IpAddress) EditObject(templateObject *datatypes.Network_Storage_Allowed_Host) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_IpAddress) GetAssignedGroups() (resp []datatypes.Network_Storage_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_IpAddress) GetAssignedReplicationVolumes() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_IpAddress) GetAssignedVolumes() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_IpAddress) GetCredential() (resp datatypes.Network_Storage_Credential, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_IpAddress) GetObject() (resp datatypes.Network_Storage_Allowed_Host_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_IpAddress) GetResource() (resp datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_IpAddress) SetCredentialPassword(password *string) (resp bool, err error) {
	params := []interface{}{
		password,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_Storage_Allowed_Host_Subnet struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkStorageAllowedHostSubnetService() Network_Storage_Allowed_Host_Subnet {
	return Network_Storage_Allowed_Host_Subnet{Session: r}
}

func (r *Network_Storage_Allowed_Host_Subnet) CreateObject(templateObject *datatypes.Network_Storage_Allowed_Host) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_Subnet) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_Subnet) EditObject(templateObject *datatypes.Network_Storage_Allowed_Host) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_Subnet) GetAssignedGroups() (resp []datatypes.Network_Storage_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_Subnet) GetAssignedReplicationVolumes() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_Subnet) GetAssignedVolumes() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_Subnet) GetCredential() (resp datatypes.Network_Storage_Credential, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_Subnet) GetObject() (resp datatypes.Network_Storage_Allowed_Host_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_Subnet) GetResource() (resp datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_Subnet) SetCredentialPassword(password *string) (resp bool, err error) {
	params := []interface{}{
		password,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_Storage_Allowed_Host_VirtualGuest struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkStorageAllowedHostVirtualGuestService() Network_Storage_Allowed_Host_VirtualGuest {
	return Network_Storage_Allowed_Host_VirtualGuest{Session: r}
}

func (r *Network_Storage_Allowed_Host_VirtualGuest) CreateObject(templateObject *datatypes.Network_Storage_Allowed_Host) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_VirtualGuest) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_VirtualGuest) EditObject(templateObject *datatypes.Network_Storage_Allowed_Host) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_VirtualGuest) GetAssignedGroups() (resp []datatypes.Network_Storage_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_VirtualGuest) GetAssignedReplicationVolumes() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_VirtualGuest) GetAssignedVolumes() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_VirtualGuest) GetCredential() (resp datatypes.Network_Storage_Credential, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_VirtualGuest) GetObject() (resp datatypes.Network_Storage_Allowed_Host_VirtualGuest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_VirtualGuest) GetResource() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Allowed_Host_VirtualGuest) SetCredentialPassword(password *string) (resp bool, err error) {
	params := []interface{}{
		password,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_Storage_Backup_Evault struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkStorageBackupEvaultService() Network_Storage_Backup_Evault {
	return Network_Storage_Backup_Evault{Session: r}
}

func (r *Network_Storage_Backup_Evault) AllowAccessFromHardware(hardwareObjectTemplate *datatypes.Hardware) (resp bool, err error) {
	params := []interface{}{
		hardwareObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) AllowAccessFromHardwareList(hardwareObjectTemplates []datatypes.Hardware) (resp bool, err error) {
	params := []interface{}{
		hardwareObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) AllowAccessFromHost(typeClassName *string, hostId *int) (resp datatypes.Network_Storage_Allowed_Host, err error) {
	params := []interface{}{
		typeClassName,
		hostId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) AllowAccessFromHostList(hostObjectTemplates []datatypes.Container_Network_Storage_Host) (resp []datatypes.Network_Storage_Allowed_Host, err error) {
	params := []interface{}{
		hostObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) AllowAccessFromIpAddress(ipAddressObjectTemplate *datatypes.Network_Subnet_IpAddress) (resp bool, err error) {
	params := []interface{}{
		ipAddressObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) AllowAccessFromIpAddressList(ipAddressObjectTemplates []datatypes.Network_Subnet_IpAddress) (resp bool, err error) {
	params := []interface{}{
		ipAddressObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) AllowAccessFromSubnet(subnetObjectTemplate *datatypes.Network_Subnet) (resp bool, err error) {
	params := []interface{}{
		subnetObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) AllowAccessFromSubnetList(subnetObjectTemplates []datatypes.Network_Subnet) (resp bool, err error) {
	params := []interface{}{
		subnetObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) AllowAccessFromVirtualGuest(virtualGuestObjectTemplate *datatypes.Virtual_Guest) (resp bool, err error) {
	params := []interface{}{
		virtualGuestObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) AllowAccessFromVirtualGuestList(virtualGuestObjectTemplates []datatypes.Virtual_Guest) (resp bool, err error) {
	params := []interface{}{
		virtualGuestObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) AllowAccessToReplicantFromHardware(hardwareObjectTemplate *datatypes.Hardware) (resp bool, err error) {
	params := []interface{}{
		hardwareObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) AllowAccessToReplicantFromHardwareList(hardwareObjectTemplates []datatypes.Hardware) (resp bool, err error) {
	params := []interface{}{
		hardwareObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) AllowAccessToReplicantFromIpAddress(ipAddressObjectTemplate *datatypes.Network_Subnet_IpAddress) (resp bool, err error) {
	params := []interface{}{
		ipAddressObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) AllowAccessToReplicantFromIpAddressList(ipAddressObjectTemplates []datatypes.Network_Subnet_IpAddress) (resp bool, err error) {
	params := []interface{}{
		ipAddressObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) AllowAccessToReplicantFromSubnet(subnetObjectTemplate *datatypes.Network_Subnet) (resp bool, err error) {
	params := []interface{}{
		subnetObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) AllowAccessToReplicantFromSubnetList(subnetObjectTemplates []datatypes.Network_Subnet) (resp bool, err error) {
	params := []interface{}{
		subnetObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) AllowAccessToReplicantFromVirtualGuest(virtualGuestObjectTemplate *datatypes.Virtual_Guest) (resp bool, err error) {
	params := []interface{}{
		virtualGuestObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) AllowAccessToReplicantFromVirtualGuestList(virtualGuestObjectTemplates []datatypes.Virtual_Guest) (resp bool, err error) {
	params := []interface{}{
		virtualGuestObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) AssignCredential(username *string) (resp bool, err error) {
	params := []interface{}{
		username,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) AssignNewCredential(typ *string) (resp datatypes.Network_Storage_Credential, err error) {
	params := []interface{}{
		typ,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) ChangePassword(username *string, currentPassword *string, newPassword *string) (resp bool, err error) {
	params := []interface{}{
		username,
		currentPassword,
		newPassword,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) CollectBandwidth(typ *string, startDate *datatypes.Time, endDate *datatypes.Time) (resp uint, err error) {
	params := []interface{}{
		typ,
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) CollectBytesUsed() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) CreateFolder(folder *string) (resp bool, err error) {
	params := []interface{}{
		folder,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) CreateSnapshot(notes *string) (resp datatypes.Network_Storage, err error) {
	params := []interface{}{
		notes,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) DeleteAllFiles() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) DeleteFile(fileId *string) (resp bool, err error) {
	params := []interface{}{
		fileId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) DeleteFiles(fileIds []string) (resp bool, err error) {
	params := []interface{}{
		fileIds,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) DeleteFolder(folder *string) (resp bool, err error) {
	params := []interface{}{
		folder,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) DeleteTasks(tasks []int) (resp bool, err error) {
	params := []interface{}{
		tasks,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) DisableSnapshots(scheduleType *string) (resp bool, err error) {
	params := []interface{}{
		scheduleType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) DownloadFile(fileId *string) (resp datatypes.Container_Utility_File_Entity, err error) {
	params := []interface{}{
		fileId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) EditCredential(username *string, newPassword *string) (resp bool, err error) {
	params := []interface{}{
		username,
		newPassword,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) EditObject(templateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) EnableSnapshots(scheduleType *string, retentionCount *int, minute *int, hour *int, dayOfWeek *string) (resp bool, err error) {
	params := []interface{}{
		scheduleType,
		retentionCount,
		minute,
		hour,
		dayOfWeek,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) FailbackFromReplicant() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) FailoverToReplicant(replicantId *int) (resp bool, err error) {
	params := []interface{}{
		replicantId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetAccountPassword() (resp datatypes.Account_Password, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetActiveTransactions() (resp []datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetAllFiles() (resp []datatypes.Container_Utility_File_Entity, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetAllFilesByFilter(filter *datatypes.Container_Utility_File_Entity) (resp []datatypes.Container_Utility_File_Entity, err error) {
	params := []interface{}{
		filter,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetAllowableHardware(filterHostname *string) (resp []datatypes.Hardware, err error) {
	params := []interface{}{
		filterHostname,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetAllowableIpAddresses(subnetId *int, filterIpAddress *string) (resp []datatypes.Network_Subnet_IpAddress, err error) {
	params := []interface{}{
		subnetId,
		filterIpAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetAllowableSubnets(filterNetworkIdentifier *string) (resp []datatypes.Network_Subnet, err error) {
	params := []interface{}{
		filterNetworkIdentifier,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetAllowableVirtualGuests(filterHostname *string) (resp []datatypes.Virtual_Guest, err error) {
	params := []interface{}{
		filterHostname,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetAllowedHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetAllowedHostsLimit() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetAllowedIpAddresses() (resp []datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetAllowedReplicationHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetAllowedReplicationIpAddresses() (resp []datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetAllowedReplicationSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetAllowedReplicationVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetAllowedSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetAllowedVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetBillingItemCategory() (resp datatypes.Product_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetByUsername(username *string, typ *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		username,
		typ,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetBytesUsed() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetCdnUrls() (resp []datatypes.Container_Network_Storage_Hub_ObjectStorage_ContentDeliveryUrl, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetClusterResource() (resp datatypes.Network_Service_Resource, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetCreationScheduleId() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetCredentials() (resp []datatypes.Network_Storage_Credential, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetDailySchedule() (resp datatypes.Network_Storage_Schedule, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetEvents() (resp []datatypes.Network_Storage_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetFileByIdentifier(identifier *string) (resp datatypes.Container_Utility_File_Entity, err error) {
	params := []interface{}{
		identifier,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetFileCount() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetFileList(folder *string, path *string) (resp []datatypes.Container_Utility_File_Entity, err error) {
	params := []interface{}{
		folder,
		path,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetFilePendingDeleteCount() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetFilesPendingDelete() (resp []datatypes.Container_Utility_File_Entity, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetFolderList() (resp []datatypes.Container_Network_Storage_Hub_ObjectStorage_Folder, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetGraph(startDate *datatypes.Time, endDate *datatypes.Time, typ *string) (resp datatypes.Container_Bandwidth_GraphOutputs, err error) {
	params := []interface{}{
		startDate,
		endDate,
		typ,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetHardwareWithEvaultFirst(option *string, exactMatch *bool, criteria *string, mode *string) (resp []datatypes.Hardware, err error) {
	params := []interface{}{
		option,
		exactMatch,
		criteria,
		mode,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetHourlySchedule() (resp datatypes.Network_Storage_Schedule, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetIops() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetIscsiLuns() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetLunId() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetManualSnapshots() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetMetricTrackingObject() (resp datatypes.Metric_Tracking_Object, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetMountableFlag() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetNetworkConnectionDetails() (resp datatypes.Container_Network_Storage_NetworkConnectionInformation, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetNotificationSubscribers() (resp []datatypes.Notification_User_Subscriber, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetObject() (resp datatypes.Network_Storage_Backup_Evault, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetObjectStorageConnectionInformation() (resp []datatypes.Container_Network_Service_Resource_ObjectStorage_ConnectionInformation, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetObjectsByCredential(credentialObject *datatypes.Network_Storage_Credential) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		credentialObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetOsType() (resp datatypes.Network_Storage_Iscsi_OS_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetOsTypeId() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetParentPartnerships() (resp []datatypes.Network_Storage_Partnership, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetParentVolume() (resp datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetPartnerships() (resp []datatypes.Network_Storage_Partnership, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetPermissionsGroups() (resp []datatypes.Network_Storage_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetProperties() (resp []datatypes.Network_Storage_Property, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetRecycleBinFileByIdentifier(fileId *string) (resp datatypes.Container_Utility_File_Entity, err error) {
	params := []interface{}{
		fileId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetRemainingAllowedHosts() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetReplicatingLuns() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetReplicatingVolume() (resp datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetReplicationEvents() (resp []datatypes.Network_Storage_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetReplicationPartners() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetReplicationSchedule() (resp datatypes.Network_Storage_Schedule, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetReplicationStatus() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetSchedules() (resp []datatypes.Network_Storage_Schedule, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetServiceResource() (resp datatypes.Network_Service_Resource, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetServiceResourceBackendIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetServiceResourceName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetSnapshotCapacityGb() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetSnapshotCreationTimestamp() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetSnapshotDeletionThresholdPercentage() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetSnapshotSizeBytes() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetSnapshotSpaceAvailable() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetSnapshots() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetSnapshotsForVolume() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetStorageGroups() (resp []datatypes.Network_Storage_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetStorageGroupsNetworkConnectionDetails() (resp []datatypes.Container_Network_Storage_NetworkConnectionInformation, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetStorageTierLevel() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetStorageType() (resp datatypes.Network_Storage_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetTotalBytesUsed() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetTotalScheduleSnapshotRetentionCount() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetUsageNotification() (resp datatypes.Notification, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetValidReplicationTargetDatacenterLocations() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetVendorName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetVirtualGuest() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetVolumeHistory() (resp []datatypes.Network_Storage_History, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetVolumeStatus() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetWebCCAuthenticationDetails() (resp datatypes.Container_Network_Storage_Backup_Evault_WebCc_Authentication_Details, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetWebccAccount() (resp datatypes.Account_Password, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) GetWeeklySchedule() (resp datatypes.Network_Storage_Schedule, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) ImmediateFailoverToReplicant(replicantId *int) (resp bool, err error) {
	params := []interface{}{
		replicantId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) InitiateBareMetalRestore() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) InitiateBareMetalRestoreForServer(hardwareId *int) (resp bool, err error) {
	params := []interface{}{
		hardwareId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) IsBlockingOperationInProgress(exemptStatusKeyNames []string) (resp bool, err error) {
	params := []interface{}{
		exemptStatusKeyNames,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) RemoveAccessFromHardware(hardwareObjectTemplate *datatypes.Hardware) (resp bool, err error) {
	params := []interface{}{
		hardwareObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) RemoveAccessFromHardwareList(hardwareObjectTemplates []datatypes.Hardware) (resp bool, err error) {
	params := []interface{}{
		hardwareObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) RemoveAccessFromHost(typeClassName *string, hostId *int) (resp datatypes.Network_Storage_Allowed_Host, err error) {
	params := []interface{}{
		typeClassName,
		hostId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) RemoveAccessFromHostList(hostObjectTemplates []datatypes.Container_Network_Storage_Host) (resp []datatypes.Network_Storage_Allowed_Host, err error) {
	params := []interface{}{
		hostObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) RemoveAccessFromIpAddress(ipAddressObjectTemplate *datatypes.Network_Subnet_IpAddress) (resp bool, err error) {
	params := []interface{}{
		ipAddressObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) RemoveAccessFromIpAddressList(ipAddressObjectTemplates []datatypes.Network_Subnet_IpAddress) (resp bool, err error) {
	params := []interface{}{
		ipAddressObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) RemoveAccessFromSubnet(subnetObjectTemplate *datatypes.Network_Subnet) (resp bool, err error) {
	params := []interface{}{
		subnetObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) RemoveAccessFromSubnetList(subnetObjectTemplates []datatypes.Network_Subnet) (resp bool, err error) {
	params := []interface{}{
		subnetObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) RemoveAccessFromVirtualGuest(virtualGuestObjectTemplate *datatypes.Virtual_Guest) (resp bool, err error) {
	params := []interface{}{
		virtualGuestObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) RemoveAccessFromVirtualGuestList(virtualGuestObjectTemplates []datatypes.Virtual_Guest) (resp bool, err error) {
	params := []interface{}{
		virtualGuestObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) RemoveAccessToReplicantFromHardwareList(hardwareObjectTemplates []datatypes.Hardware) (resp bool, err error) {
	params := []interface{}{
		hardwareObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) RemoveAccessToReplicantFromIpAddressList(ipAddressObjectTemplates []datatypes.Network_Subnet_IpAddress) (resp bool, err error) {
	params := []interface{}{
		ipAddressObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) RemoveAccessToReplicantFromSubnet(subnetObjectTemplate *datatypes.Network_Subnet) (resp bool, err error) {
	params := []interface{}{
		subnetObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) RemoveAccessToReplicantFromSubnetList(subnetObjectTemplates []datatypes.Network_Subnet) (resp bool, err error) {
	params := []interface{}{
		subnetObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) RemoveAccessToReplicantFromVirtualGuestList(virtualGuestObjectTemplates []datatypes.Virtual_Guest) (resp bool, err error) {
	params := []interface{}{
		virtualGuestObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) RemoveCredential(username *string) (resp bool, err error) {
	params := []interface{}{
		username,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) RestoreFile(fileId *string) (resp datatypes.Container_Utility_File_Entity, err error) {
	params := []interface{}{
		fileId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) RestoreFromSnapshot(snapshotId *int) (resp bool, err error) {
	params := []interface{}{
		snapshotId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) SendPasswordReminderEmail(username *string) (resp bool, err error) {
	params := []interface{}{
		username,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) SetMountable(mountable *bool) (resp bool, err error) {
	params := []interface{}{
		mountable,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) SetSnapshotAllocation(capacityGb *int) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		capacityGb,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) UpgradeVolumeCapacity(itemId *int) (resp bool, err error) {
	params := []interface{}{
		itemId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Backup_Evault) UploadFile(file *datatypes.Container_Utility_File_Entity) (resp datatypes.Container_Utility_File_Entity, err error) {
	params := []interface{}{
		file,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_Storage_Group struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkStorageGroupService() Network_Storage_Group {
	return Network_Storage_Group{Session: r}
}

func (r *Network_Storage_Group) AddAllowedHost(allowedHost *datatypes.Network_Storage_Allowed_Host) (resp bool, err error) {
	params := []interface{}{
		allowedHost,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group) AttachToVolume(volume *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		volume,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group) CreateObject(templateObject *datatypes.Network_Storage_Group) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group) EditObject(templateObject *datatypes.Network_Storage_Group) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group) GetAllObjects() (resp []datatypes.Network_Storage_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group) GetAllowedHosts() (resp []datatypes.Network_Storage_Allowed_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group) GetAttachedVolumes() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group) GetGroupType() (resp datatypes.Network_Storage_Group_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group) GetNetworkConnectionDetails() (resp datatypes.Container_Network_Storage_NetworkConnectionInformation, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group) GetObject() (resp datatypes.Network_Storage_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group) GetOsType() (resp datatypes.Network_Storage_Iscsi_OS_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group) GetServiceResource() (resp datatypes.Network_Service_Resource, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group) RemoveAllowedHost(allowedHost *datatypes.Network_Storage_Allowed_Host) (resp bool, err error) {
	params := []interface{}{
		allowedHost,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group) RemoveFromVolume(volume *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		volume,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_Storage_Group_Iscsi struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkStorageGroupIscsiService() Network_Storage_Group_Iscsi {
	return Network_Storage_Group_Iscsi{Session: r}
}

func (r *Network_Storage_Group_Iscsi) AddAllowedHost(allowedHost *datatypes.Network_Storage_Allowed_Host) (resp bool, err error) {
	params := []interface{}{
		allowedHost,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Iscsi) AttachToVolume(volume *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		volume,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Iscsi) CreateObject(templateObject *datatypes.Network_Storage_Group) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Iscsi) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Iscsi) EditObject(templateObject *datatypes.Network_Storage_Group) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Iscsi) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Iscsi) GetAllObjects() (resp []datatypes.Network_Storage_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Iscsi) GetAllowedHosts() (resp []datatypes.Network_Storage_Allowed_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Iscsi) GetAttachedVolumes() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Iscsi) GetGroupType() (resp datatypes.Network_Storage_Group_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Iscsi) GetNetworkConnectionDetails() (resp datatypes.Container_Network_Storage_NetworkConnectionInformation, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Iscsi) GetObject() (resp datatypes.Network_Storage_Group_Iscsi, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Iscsi) GetOsType() (resp datatypes.Network_Storage_Iscsi_OS_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Iscsi) GetServiceResource() (resp datatypes.Network_Service_Resource, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Iscsi) RemoveAllowedHost(allowedHost *datatypes.Network_Storage_Allowed_Host) (resp bool, err error) {
	params := []interface{}{
		allowedHost,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Iscsi) RemoveFromVolume(volume *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		volume,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_Storage_Group_Nfs struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkStorageGroupNfsService() Network_Storage_Group_Nfs {
	return Network_Storage_Group_Nfs{Session: r}
}

func (r *Network_Storage_Group_Nfs) AddAllowedHost(allowedHost *datatypes.Network_Storage_Allowed_Host) (resp bool, err error) {
	params := []interface{}{
		allowedHost,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Nfs) AttachToVolume(volume *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		volume,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Nfs) CreateObject(templateObject *datatypes.Network_Storage_Group) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Nfs) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Nfs) EditObject(templateObject *datatypes.Network_Storage_Group) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Nfs) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Nfs) GetAllObjects() (resp []datatypes.Network_Storage_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Nfs) GetAllowedHosts() (resp []datatypes.Network_Storage_Allowed_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Nfs) GetAttachedVolumes() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Nfs) GetGroupType() (resp datatypes.Network_Storage_Group_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Nfs) GetNetworkConnectionDetails() (resp datatypes.Container_Network_Storage_NetworkConnectionInformation, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Nfs) GetObject() (resp datatypes.Network_Storage_Group_Nfs, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Nfs) GetOsType() (resp datatypes.Network_Storage_Iscsi_OS_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Nfs) GetServiceResource() (resp datatypes.Network_Service_Resource, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Nfs) RemoveAllowedHost(allowedHost *datatypes.Network_Storage_Allowed_Host) (resp bool, err error) {
	params := []interface{}{
		allowedHost,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Nfs) RemoveFromVolume(volume *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		volume,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_Storage_Group_Type struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkStorageGroupTypeService() Network_Storage_Group_Type {
	return Network_Storage_Group_Type{Session: r}
}

func (r *Network_Storage_Group_Type) GetAllObjects() (resp []datatypes.Network_Storage_Group_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Group_Type) GetObject() (resp datatypes.Network_Storage_Group_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Storage_Hub_Cleversafe_Account struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkStorageHubCleversafeAccountService() Network_Storage_Hub_Cleversafe_Account {
	return Network_Storage_Hub_Cleversafe_Account{Session: r}
}

func (r *Network_Storage_Hub_Cleversafe_Account) CredentialCreate() (resp []datatypes.Network_Storage_Credential, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Hub_Cleversafe_Account) CredentialDelete(credential *datatypes.Network_Storage_Credential) (resp bool, err error) {
	params := []interface{}{
		credential,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Hub_Cleversafe_Account) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Hub_Cleversafe_Account) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Hub_Cleversafe_Account) GetCancelledBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Hub_Cleversafe_Account) GetCapacityUsage() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Hub_Cleversafe_Account) GetCredentialLimit() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Hub_Cleversafe_Account) GetCredentials() (resp []datatypes.Network_Storage_Credential, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Hub_Cleversafe_Account) GetEndpoints() (resp []datatypes.Container_Network_Storage_Hub_ObjectStorage_Endpoint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Hub_Cleversafe_Account) GetEvents() (resp []datatypes.Network_Storage_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Hub_Cleversafe_Account) GetMetricTrackingObject() (resp datatypes.Metric_Tracking_Object, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Hub_Cleversafe_Account) GetObject() (resp datatypes.Network_Storage_Hub_Cleversafe_Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Hub_Cleversafe_Account) GetUuid() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Storage_Hub_Swift_Share struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkStorageHubSwiftShareService() Network_Storage_Hub_Swift_Share {
	return Network_Storage_Hub_Swift_Share{Session: r}
}

func (r *Network_Storage_Hub_Swift_Share) GetContainerList() (resp []datatypes.Container_Network_Storage_Hub_ObjectStorage_Folder, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Hub_Swift_Share) GetFile(fileName *string, container *string) (resp datatypes.Container_Network_Storage_Hub_ObjectStorage_File, err error) {
	params := []interface{}{
		fileName,
		container,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Hub_Swift_Share) GetFileList(container *string, path *string) (resp []datatypes.Container_Utility_File_Entity, err error) {
	params := []interface{}{
		container,
		path,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_Storage_Iscsi struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkStorageIscsiService() Network_Storage_Iscsi {
	return Network_Storage_Iscsi{Session: r}
}

func (r *Network_Storage_Iscsi) AllowAccessFromHardware(hardwareObjectTemplate *datatypes.Hardware) (resp bool, err error) {
	params := []interface{}{
		hardwareObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) AllowAccessFromHardwareList(hardwareObjectTemplates []datatypes.Hardware) (resp bool, err error) {
	params := []interface{}{
		hardwareObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) AllowAccessFromHost(typeClassName *string, hostId *int) (resp datatypes.Network_Storage_Allowed_Host, err error) {
	params := []interface{}{
		typeClassName,
		hostId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) AllowAccessFromHostList(hostObjectTemplates []datatypes.Container_Network_Storage_Host) (resp []datatypes.Network_Storage_Allowed_Host, err error) {
	params := []interface{}{
		hostObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) AllowAccessFromIpAddress(ipAddressObjectTemplate *datatypes.Network_Subnet_IpAddress) (resp bool, err error) {
	params := []interface{}{
		ipAddressObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) AllowAccessFromIpAddressList(ipAddressObjectTemplates []datatypes.Network_Subnet_IpAddress) (resp bool, err error) {
	params := []interface{}{
		ipAddressObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) AllowAccessFromSubnet(subnetObjectTemplate *datatypes.Network_Subnet) (resp bool, err error) {
	params := []interface{}{
		subnetObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) AllowAccessFromSubnetList(subnetObjectTemplates []datatypes.Network_Subnet) (resp bool, err error) {
	params := []interface{}{
		subnetObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) AllowAccessFromVirtualGuest(virtualGuestObjectTemplate *datatypes.Virtual_Guest) (resp bool, err error) {
	params := []interface{}{
		virtualGuestObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) AllowAccessFromVirtualGuestList(virtualGuestObjectTemplates []datatypes.Virtual_Guest) (resp bool, err error) {
	params := []interface{}{
		virtualGuestObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) AllowAccessToReplicantFromHardware(hardwareObjectTemplate *datatypes.Hardware) (resp bool, err error) {
	params := []interface{}{
		hardwareObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) AllowAccessToReplicantFromHardwareList(hardwareObjectTemplates []datatypes.Hardware) (resp bool, err error) {
	params := []interface{}{
		hardwareObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) AllowAccessToReplicantFromIpAddress(ipAddressObjectTemplate *datatypes.Network_Subnet_IpAddress) (resp bool, err error) {
	params := []interface{}{
		ipAddressObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) AllowAccessToReplicantFromIpAddressList(ipAddressObjectTemplates []datatypes.Network_Subnet_IpAddress) (resp bool, err error) {
	params := []interface{}{
		ipAddressObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) AllowAccessToReplicantFromSubnet(subnetObjectTemplate *datatypes.Network_Subnet) (resp bool, err error) {
	params := []interface{}{
		subnetObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) AllowAccessToReplicantFromSubnetList(subnetObjectTemplates []datatypes.Network_Subnet) (resp bool, err error) {
	params := []interface{}{
		subnetObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) AllowAccessToReplicantFromVirtualGuest(virtualGuestObjectTemplate *datatypes.Virtual_Guest) (resp bool, err error) {
	params := []interface{}{
		virtualGuestObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) AllowAccessToReplicantFromVirtualGuestList(virtualGuestObjectTemplates []datatypes.Virtual_Guest) (resp bool, err error) {
	params := []interface{}{
		virtualGuestObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) AssignCredential(username *string) (resp bool, err error) {
	params := []interface{}{
		username,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) AssignNewCredential(typ *string) (resp datatypes.Network_Storage_Credential, err error) {
	params := []interface{}{
		typ,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) ChangePassword(username *string, currentPassword *string, newPassword *string) (resp bool, err error) {
	params := []interface{}{
		username,
		currentPassword,
		newPassword,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) CollectBandwidth(typ *string, startDate *datatypes.Time, endDate *datatypes.Time) (resp uint, err error) {
	params := []interface{}{
		typ,
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) CollectBytesUsed() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) CreateFolder(folder *string) (resp bool, err error) {
	params := []interface{}{
		folder,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) CreateSnapshot(notes *string) (resp datatypes.Network_Storage, err error) {
	params := []interface{}{
		notes,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) DeleteAllFiles() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) DeleteFile(fileId *string) (resp bool, err error) {
	params := []interface{}{
		fileId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) DeleteFiles(fileIds []string) (resp bool, err error) {
	params := []interface{}{
		fileIds,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) DeleteFolder(folder *string) (resp bool, err error) {
	params := []interface{}{
		folder,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) DisableSnapshots(scheduleType *string) (resp bool, err error) {
	params := []interface{}{
		scheduleType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) DownloadFile(fileId *string) (resp datatypes.Container_Utility_File_Entity, err error) {
	params := []interface{}{
		fileId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) EditCredential(username *string, newPassword *string) (resp bool, err error) {
	params := []interface{}{
		username,
		newPassword,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) EditObject(templateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) EnableSnapshots(scheduleType *string, retentionCount *int, minute *int, hour *int, dayOfWeek *string) (resp bool, err error) {
	params := []interface{}{
		scheduleType,
		retentionCount,
		minute,
		hour,
		dayOfWeek,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) FailbackFromReplicant() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) FailoverToReplicant(replicantId *int) (resp bool, err error) {
	params := []interface{}{
		replicantId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetAccountPassword() (resp datatypes.Account_Password, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetActiveTransactions() (resp []datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetAllFiles() (resp []datatypes.Container_Utility_File_Entity, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetAllFilesByFilter(filter *datatypes.Container_Utility_File_Entity) (resp []datatypes.Container_Utility_File_Entity, err error) {
	params := []interface{}{
		filter,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetAllowableHardware(filterHostname *string) (resp []datatypes.Hardware, err error) {
	params := []interface{}{
		filterHostname,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetAllowableIpAddresses(subnetId *int, filterIpAddress *string) (resp []datatypes.Network_Subnet_IpAddress, err error) {
	params := []interface{}{
		subnetId,
		filterIpAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetAllowableSubnets(filterNetworkIdentifier *string) (resp []datatypes.Network_Subnet, err error) {
	params := []interface{}{
		filterNetworkIdentifier,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetAllowableVirtualGuests(filterHostname *string) (resp []datatypes.Virtual_Guest, err error) {
	params := []interface{}{
		filterHostname,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetAllowedHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetAllowedHostsLimit() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetAllowedIpAddresses() (resp []datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetAllowedReplicationHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetAllowedReplicationIpAddresses() (resp []datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetAllowedReplicationSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetAllowedReplicationVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetAllowedSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetAllowedVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetBillingItemCategory() (resp datatypes.Product_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetByUsername(username *string, typ *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		username,
		typ,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetBytesUsed() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetCdnUrls() (resp []datatypes.Container_Network_Storage_Hub_ObjectStorage_ContentDeliveryUrl, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetClusterResource() (resp datatypes.Network_Service_Resource, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetCreationScheduleId() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetCredentials() (resp []datatypes.Network_Storage_Credential, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetDailySchedule() (resp datatypes.Network_Storage_Schedule, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetEvents() (resp []datatypes.Network_Storage_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetFileByIdentifier(identifier *string) (resp datatypes.Container_Utility_File_Entity, err error) {
	params := []interface{}{
		identifier,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetFileCount() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetFileList(folder *string, path *string) (resp []datatypes.Container_Utility_File_Entity, err error) {
	params := []interface{}{
		folder,
		path,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetFilePendingDeleteCount() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetFilesPendingDelete() (resp []datatypes.Container_Utility_File_Entity, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetFolderList() (resp []datatypes.Container_Network_Storage_Hub_ObjectStorage_Folder, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetGraph(startDate *datatypes.Time, endDate *datatypes.Time, typ *string) (resp datatypes.Container_Bandwidth_GraphOutputs, err error) {
	params := []interface{}{
		startDate,
		endDate,
		typ,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetHourlySchedule() (resp datatypes.Network_Storage_Schedule, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetIops() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetIscsiLuns() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetLunId() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetManualSnapshots() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetMetricTrackingObject() (resp datatypes.Metric_Tracking_Object, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetMountableFlag() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetNetworkConnectionDetails() (resp datatypes.Container_Network_Storage_NetworkConnectionInformation, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetNotificationSubscribers() (resp []datatypes.Notification_User_Subscriber, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetObject() (resp datatypes.Network_Storage_Iscsi, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetObjectStorageConnectionInformation() (resp []datatypes.Container_Network_Service_Resource_ObjectStorage_ConnectionInformation, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetObjectsByCredential(credentialObject *datatypes.Network_Storage_Credential) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		credentialObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetOsType() (resp datatypes.Network_Storage_Iscsi_OS_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetOsTypeId() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetParentPartnerships() (resp []datatypes.Network_Storage_Partnership, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetParentVolume() (resp datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetPartnerships() (resp []datatypes.Network_Storage_Partnership, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetPermissionsGroups() (resp []datatypes.Network_Storage_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetProperties() (resp []datatypes.Network_Storage_Property, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetRecycleBinFileByIdentifier(fileId *string) (resp datatypes.Container_Utility_File_Entity, err error) {
	params := []interface{}{
		fileId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetRemainingAllowedHosts() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetReplicatingLuns() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetReplicatingVolume() (resp datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetReplicationEvents() (resp []datatypes.Network_Storage_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetReplicationPartners() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetReplicationSchedule() (resp datatypes.Network_Storage_Schedule, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetReplicationStatus() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetSchedules() (resp []datatypes.Network_Storage_Schedule, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetServiceResource() (resp datatypes.Network_Service_Resource, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetServiceResourceBackendIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetServiceResourceName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetSnapshotCapacityGb() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetSnapshotCreationTimestamp() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetSnapshotDeletionThresholdPercentage() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetSnapshotSizeBytes() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetSnapshotSpaceAvailable() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetSnapshots() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetSnapshotsForVolume() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetStorageGroups() (resp []datatypes.Network_Storage_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetStorageGroupsNetworkConnectionDetails() (resp []datatypes.Container_Network_Storage_NetworkConnectionInformation, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetStorageTierLevel() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetStorageType() (resp datatypes.Network_Storage_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetTotalBytesUsed() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetTotalScheduleSnapshotRetentionCount() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetUsageNotification() (resp datatypes.Notification, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetValidReplicationTargetDatacenterLocations() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetVendorName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetVirtualGuest() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetVolumeHistory() (resp []datatypes.Network_Storage_History, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetVolumeStatus() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetWebccAccount() (resp datatypes.Account_Password, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) GetWeeklySchedule() (resp datatypes.Network_Storage_Schedule, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) ImmediateFailoverToReplicant(replicantId *int) (resp bool, err error) {
	params := []interface{}{
		replicantId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) IsBlockingOperationInProgress(exemptStatusKeyNames []string) (resp bool, err error) {
	params := []interface{}{
		exemptStatusKeyNames,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) RemoveAccessFromHardware(hardwareObjectTemplate *datatypes.Hardware) (resp bool, err error) {
	params := []interface{}{
		hardwareObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) RemoveAccessFromHardwareList(hardwareObjectTemplates []datatypes.Hardware) (resp bool, err error) {
	params := []interface{}{
		hardwareObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) RemoveAccessFromHost(typeClassName *string, hostId *int) (resp datatypes.Network_Storage_Allowed_Host, err error) {
	params := []interface{}{
		typeClassName,
		hostId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) RemoveAccessFromHostList(hostObjectTemplates []datatypes.Container_Network_Storage_Host) (resp []datatypes.Network_Storage_Allowed_Host, err error) {
	params := []interface{}{
		hostObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) RemoveAccessFromIpAddress(ipAddressObjectTemplate *datatypes.Network_Subnet_IpAddress) (resp bool, err error) {
	params := []interface{}{
		ipAddressObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) RemoveAccessFromIpAddressList(ipAddressObjectTemplates []datatypes.Network_Subnet_IpAddress) (resp bool, err error) {
	params := []interface{}{
		ipAddressObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) RemoveAccessFromSubnet(subnetObjectTemplate *datatypes.Network_Subnet) (resp bool, err error) {
	params := []interface{}{
		subnetObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) RemoveAccessFromSubnetList(subnetObjectTemplates []datatypes.Network_Subnet) (resp bool, err error) {
	params := []interface{}{
		subnetObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) RemoveAccessFromVirtualGuest(virtualGuestObjectTemplate *datatypes.Virtual_Guest) (resp bool, err error) {
	params := []interface{}{
		virtualGuestObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) RemoveAccessFromVirtualGuestList(virtualGuestObjectTemplates []datatypes.Virtual_Guest) (resp bool, err error) {
	params := []interface{}{
		virtualGuestObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) RemoveAccessToReplicantFromHardwareList(hardwareObjectTemplates []datatypes.Hardware) (resp bool, err error) {
	params := []interface{}{
		hardwareObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) RemoveAccessToReplicantFromIpAddressList(ipAddressObjectTemplates []datatypes.Network_Subnet_IpAddress) (resp bool, err error) {
	params := []interface{}{
		ipAddressObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) RemoveAccessToReplicantFromSubnet(subnetObjectTemplate *datatypes.Network_Subnet) (resp bool, err error) {
	params := []interface{}{
		subnetObjectTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) RemoveAccessToReplicantFromSubnetList(subnetObjectTemplates []datatypes.Network_Subnet) (resp bool, err error) {
	params := []interface{}{
		subnetObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) RemoveAccessToReplicantFromVirtualGuestList(virtualGuestObjectTemplates []datatypes.Virtual_Guest) (resp bool, err error) {
	params := []interface{}{
		virtualGuestObjectTemplates,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) RemoveCredential(username *string) (resp bool, err error) {
	params := []interface{}{
		username,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) RestoreFile(fileId *string) (resp datatypes.Container_Utility_File_Entity, err error) {
	params := []interface{}{
		fileId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) RestoreFromSnapshot(snapshotId *int) (resp bool, err error) {
	params := []interface{}{
		snapshotId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) SendPasswordReminderEmail(username *string) (resp bool, err error) {
	params := []interface{}{
		username,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) SetMountable(mountable *bool) (resp bool, err error) {
	params := []interface{}{
		mountable,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) SetSnapshotAllocation(capacityGb *int) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		capacityGb,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) UpgradeVolumeCapacity(itemId *int) (resp bool, err error) {
	params := []interface{}{
		itemId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi) UploadFile(file *datatypes.Container_Utility_File_Entity) (resp datatypes.Container_Utility_File_Entity, err error) {
	params := []interface{}{
		file,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_Storage_Iscsi_OS_Type struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkStorageIscsiOSTypeService() Network_Storage_Iscsi_OS_Type {
	return Network_Storage_Iscsi_OS_Type{Session: r}
}

func (r *Network_Storage_Iscsi_OS_Type) GetAllObjects() (resp []datatypes.Network_Storage_Iscsi_OS_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Iscsi_OS_Type) GetObject() (resp datatypes.Network_Storage_Iscsi_OS_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Storage_Schedule struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkStorageScheduleService() Network_Storage_Schedule {
	return Network_Storage_Schedule{Session: r}
}

func (r *Network_Storage_Schedule) CreateObject(templateObject *datatypes.Network_Storage_Schedule) (resp datatypes.Network_Storage_Schedule, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Schedule) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Schedule) EditObject(templateObject *datatypes.Network_Storage_Schedule) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Schedule) GetDayOfMonth() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Schedule) GetDayOfWeek() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Schedule) GetEvents() (resp []datatypes.Network_Storage_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Schedule) GetHour() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Schedule) GetMinute() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Schedule) GetMonthOfYear() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Schedule) GetObject() (resp datatypes.Network_Storage_Schedule, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Schedule) GetPartnership() (resp datatypes.Network_Storage_Partnership, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Schedule) GetProperties() (resp []datatypes.Network_Storage_Schedule_Property, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Schedule) GetReplicaSnapshots() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Schedule) GetRetentionCount() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Schedule) GetSnapshots() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Schedule) GetType() (resp datatypes.Network_Storage_Schedule_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Schedule) GetVolume() (resp datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Storage_Schedule_Property_Type struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkStorageSchedulePropertyTypeService() Network_Storage_Schedule_Property_Type {
	return Network_Storage_Schedule_Property_Type{Session: r}
}

func (r *Network_Storage_Schedule_Property_Type) GetAllObjects() (resp []datatypes.Network_Storage_Schedule_Property_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Storage_Schedule_Property_Type) GetObject() (resp datatypes.Network_Storage_Schedule_Property_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Subnet struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkSubnetService() Network_Subnet {
	return Network_Subnet{Session: r}
}

func (r *Network_Subnet) AllowAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) AllowAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) CreateReverseDomainRecords() (resp datatypes.Dns_Domain_Reverse, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) CreateSubnetRouteUpdateTransaction(newEndPointIpAddress *string) (resp bool, err error) {
	params := []interface{}{
		newEndPointIpAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) CreateSwipTransaction() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) EditNote(note *string) (resp bool, err error) {
	params := []interface{}{
		note,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) FindAllSubnetsAndActiveSwipTransactionStatus() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetActiveRegistration() (resp datatypes.Network_Subnet_Registration, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetActiveSwipTransaction() (resp datatypes.Network_Subnet_Swip_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetActiveTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetAddressSpace() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetAllowedHost() (resp datatypes.Network_Storage_Allowed_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetAllowedNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetAllowedNetworkStorageReplicas() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetAttachedNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetAvailableNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetBoundDescendants() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetBoundRouterFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetBoundRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetChildren() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetDatacenter() (resp datatypes.Location_Datacenter, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetDescendants() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetDisplayLabel() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetEndPointIpAddress() (resp datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetGlobalIpRecord() (resp datatypes.Network_Subnet_IpAddress_Global, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetIpAddresses() (resp []datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetNetworkComponentFirewall() (resp datatypes.Network_Component_Firewall, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetNetworkProtectionAddresses() (resp []datatypes.Network_Protection_Address, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetNetworkTunnelContexts() (resp []datatypes.Network_Tunnel_Module_Context, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetNetworkVlan() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetObject() (resp datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetPodName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetProtectedIpAddresses() (resp []datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetRegionalInternetRegistry() (resp datatypes.Network_Regional_Internet_Registry, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetRegistrations() (resp []datatypes.Network_Subnet_Registration, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetResourceGroups() (resp []datatypes.Resource_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetReverseDomain() (resp datatypes.Dns_Domain, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetReverseDomainRecords() (resp []datatypes.Dns_Domain, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetRoleKeyName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetRoleName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetRoutableEndpointIpAddresses() (resp []datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetRoutingTypeKeyName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetRoutingTypeName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetSubnetForIpAddress(ipAddress *string) (resp datatypes.Network_Subnet, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetSwipTransaction() (resp []datatypes.Network_Subnet_Swip_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetUnboundDescendants() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) GetVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet) RemoveAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_Subnet_IpAddress struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkSubnetIpAddressService() Network_Subnet_IpAddress {
	return Network_Subnet_IpAddress{Session: r}
}

func (r *Network_Subnet_IpAddress) AllowAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) AllowAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) EditObject(templateObject *datatypes.Network_Subnet_IpAddress) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) EditObjects(templateObjects []datatypes.Network_Subnet_IpAddress) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) FindByIpv4Address(ipAddress *string) (resp datatypes.Network_Subnet_IpAddress, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetAllowedHost() (resp datatypes.Network_Storage_Allowed_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetAllowedNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetAllowedNetworkStorageReplicas() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetApplicationDeliveryController() (resp datatypes.Network_Application_Delivery_Controller, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetAttachedNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetAvailableNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetByIpAddress(ipAddress *string) (resp datatypes.Network_Subnet_IpAddress, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetContextTunnelTranslations() (resp []datatypes.Network_Tunnel_Module_Context_Address_Translation, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetEndpointSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetGuestNetworkComponent() (resp datatypes.Virtual_Guest_Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetGuestNetworkComponentBinding() (resp datatypes.Virtual_Guest_Network_Component_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetObject() (resp datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetPrivateNetworkGateway() (resp datatypes.Network_Gateway, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetProtectionAddress() (resp []datatypes.Network_Protection_Address, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetPublicNetworkGateway() (resp datatypes.Network_Gateway, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetRemoteManagementNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetSubnet() (resp datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetSyslogEventsOneDay() (resp []datatypes.Network_Logging_Syslog, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetSyslogEventsSevenDays() (resp []datatypes.Network_Logging_Syslog, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetTopTenSyslogEventsByDestinationPortOneDay() (resp []datatypes.Network_Logging_Syslog, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetTopTenSyslogEventsByDestinationPortSevenDays() (resp []datatypes.Network_Logging_Syslog, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetTopTenSyslogEventsByProtocolsOneDay() (resp []datatypes.Network_Logging_Syslog, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetTopTenSyslogEventsByProtocolsSevenDays() (resp []datatypes.Network_Logging_Syslog, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetTopTenSyslogEventsBySourceIpOneDay() (resp []datatypes.Network_Logging_Syslog, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetTopTenSyslogEventsBySourceIpSevenDays() (resp []datatypes.Network_Logging_Syslog, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetTopTenSyslogEventsBySourcePortOneDay() (resp []datatypes.Network_Logging_Syslog, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetTopTenSyslogEventsBySourcePortSevenDays() (resp []datatypes.Network_Logging_Syslog, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetVirtualGuest() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) GetVirtualLicenses() (resp []datatypes.Software_VirtualLicense, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress) RemoveAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_Subnet_IpAddress_Global struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkSubnetIpAddressGlobalService() Network_Subnet_IpAddress_Global {
	return Network_Subnet_IpAddress_Global{Session: r}
}

func (r *Network_Subnet_IpAddress_Global) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress_Global) GetActiveTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress_Global) GetBillingItem() (resp datatypes.Billing_Item_Network_Subnet_IpAddress_Global, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress_Global) GetDestinationIpAddress() (resp datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress_Global) GetIpAddress() (resp datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress_Global) GetObject() (resp datatypes.Network_Subnet_IpAddress_Global, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress_Global) Route(newEndPointIpAddress *string) (resp datatypes.Provisioning_Version1_Transaction, err error) {
	params := []interface{}{
		newEndPointIpAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_IpAddress_Global) Unroute() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Subnet_Registration struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkSubnetRegistrationService() Network_Subnet_Registration {
	return Network_Subnet_Registration{Session: r}
}

func (r *Network_Subnet_Registration) ClearRegistration() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Registration) CreateObject(templateObject *datatypes.Network_Subnet_Registration) (resp datatypes.Network_Subnet_Registration, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Registration) EditObject(templateObject *datatypes.Network_Subnet_Registration) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Registration) EditRegistrationAttachedDetails(personObjectSkeleton *datatypes.Network_Subnet_Registration_Details, networkObjectSkeleton *datatypes.Network_Subnet_Registration_Details) (resp bool, err error) {
	params := []interface{}{
		personObjectSkeleton,
		networkObjectSkeleton,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Registration) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Registration) GetDetailReferences() (resp []datatypes.Network_Subnet_Registration_Details, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Registration) GetEvents() (resp []datatypes.Network_Subnet_Registration_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Registration) GetNetworkDetail() (resp datatypes.Account_Regional_Registry_Detail, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Registration) GetObject() (resp datatypes.Network_Subnet_Registration, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Registration) GetPersonDetail() (resp datatypes.Account_Regional_Registry_Detail, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Registration) GetRegionalInternetRegistry() (resp datatypes.Network_Regional_Internet_Registry, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Registration) GetRegionalInternetRegistryHandle() (resp datatypes.Account_Rwhois_Handle, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Registration) GetStatus() (resp datatypes.Network_Subnet_Registration_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Registration) GetSubnet() (resp datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Subnet_Registration_Details struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkSubnetRegistrationDetailsService() Network_Subnet_Registration_Details {
	return Network_Subnet_Registration_Details{Session: r}
}

func (r *Network_Subnet_Registration_Details) CreateObject(templateObject *datatypes.Network_Subnet_Registration_Details) (resp datatypes.Network_Subnet_Registration_Details, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Registration_Details) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Registration_Details) GetDetail() (resp datatypes.Account_Regional_Registry_Detail, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Registration_Details) GetObject() (resp datatypes.Network_Subnet_Registration_Details, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Registration_Details) GetRegistration() (resp datatypes.Network_Subnet_Registration, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Subnet_Registration_Status struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkSubnetRegistrationStatusService() Network_Subnet_Registration_Status {
	return Network_Subnet_Registration_Status{Session: r}
}

func (r *Network_Subnet_Registration_Status) GetAllObjects() (resp []datatypes.Network_Subnet_Registration_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Registration_Status) GetObject() (resp datatypes.Network_Subnet_Registration_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Subnet_Rwhois_Data struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkSubnetRwhoisDataService() Network_Subnet_Rwhois_Data {
	return Network_Subnet_Rwhois_Data{Session: r}
}

func (r *Network_Subnet_Rwhois_Data) EditObject(templateObject *datatypes.Network_Subnet_Rwhois_Data) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Rwhois_Data) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Rwhois_Data) GetObject() (resp datatypes.Network_Subnet_Rwhois_Data, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_Subnet_Swip_Transaction struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkSubnetSwipTransactionService() Network_Subnet_Swip_Transaction {
	return Network_Subnet_Swip_Transaction{Session: r}
}

func (r *Network_Subnet_Swip_Transaction) FindMyTransactions() (resp []datatypes.Network_Subnet_Swip_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Swip_Transaction) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Swip_Transaction) GetObject() (resp datatypes.Network_Subnet_Swip_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Swip_Transaction) GetSubnet() (resp datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Swip_Transaction) RemoveAllSubnetSwips() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Swip_Transaction) RemoveSwipData() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Swip_Transaction) ResendSwipData() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Swip_Transaction) SwipAllSubnets() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Subnet_Swip_Transaction) UpdateAllSubnetSwips() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Network_TippingPointReporting struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkTippingPointReportingService() Network_TippingPointReporting {
	return Network_TippingPointReporting{Session: r}
}

func (r *Network_TippingPointReporting) DrillDownAttack(signatureId *string, IpAddress *string, subnetMask *int, timeFrame *int, direction *string) (resp datatypes.Container_Network_IntrusionProtection_SubnetReport, err error) {
	params := []interface{}{
		signatureId,
		IpAddress,
		subnetMask,
		timeFrame,
		direction,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_TippingPointReporting) GetMainStatistics(numberOfAttacks *int) (resp []datatypes.Container_Network_IntrusionProtection_Statistics, err error) {
	params := []interface{}{
		numberOfAttacks,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_TippingPointReporting) GetReportForIpAddressOrSubnet(IpAddress *string, subnetMask *int, timeFrame *int, orderBy *string, orderDirection *string) (resp []datatypes.Container_Network_IntrusionProtection_SubnetReport, err error) {
	params := []interface{}{
		IpAddress,
		subnetMask,
		timeFrame,
		orderBy,
		orderDirection,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_TippingPointReporting) GetSubnetReportForEntireAccount(timeFrame *int, orderBy *string, orderDirection *string, returnSubnetGroups *bool) (resp []datatypes.Container_Network_IntrusionProtection_SubnetReport, err error) {
	params := []interface{}{
		timeFrame,
		orderBy,
		orderDirection,
		returnSubnetGroups,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_Tunnel_Module_Context struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkTunnelModuleContextService() Network_Tunnel_Module_Context {
	return Network_Tunnel_Module_Context{Session: r}
}

func (r *Network_Tunnel_Module_Context) AddCustomerSubnetToNetworkTunnel(subnetId *int) (resp bool, err error) {
	params := []interface{}{
		subnetId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) AddPrivateSubnetToNetworkTunnel(subnetId *int) (resp bool, err error) {
	params := []interface{}{
		subnetId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) AddServiceSubnetToNetworkTunnel(subnetId *int) (resp bool, err error) {
	params := []interface{}{
		subnetId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) ApplyConfigurationsToDevice() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) CreateAddressTranslation(translation *datatypes.Network_Tunnel_Module_Context_Address_Translation) (resp datatypes.Network_Tunnel_Module_Context_Address_Translation, err error) {
	params := []interface{}{
		translation,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) CreateAddressTranslations(translations []datatypes.Network_Tunnel_Module_Context_Address_Translation) (resp []datatypes.Network_Tunnel_Module_Context_Address_Translation, err error) {
	params := []interface{}{
		translations,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) DeleteAddressTranslation(translationId *int) (resp bool, err error) {
	params := []interface{}{
		translationId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) DownloadAddressTranslationConfigurations() (resp datatypes.Container_Utility_File_Entity, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) DownloadParameterConfigurations() (resp datatypes.Container_Utility_File_Entity, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) EditAddressTranslation(translation *datatypes.Network_Tunnel_Module_Context_Address_Translation) (resp datatypes.Network_Tunnel_Module_Context_Address_Translation, err error) {
	params := []interface{}{
		translation,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) EditAddressTranslations(translations []datatypes.Network_Tunnel_Module_Context_Address_Translation) (resp []datatypes.Network_Tunnel_Module_Context_Address_Translation, err error) {
	params := []interface{}{
		translations,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) EditObject(templateObject *datatypes.Network_Tunnel_Module_Context) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) GetActiveTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) GetAddressTranslationConfigurations() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) GetAddressTranslations() (resp []datatypes.Network_Tunnel_Module_Context_Address_Translation, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) GetAllAvailableServiceSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) GetAuthenticationDefault() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) GetAuthenticationOptions() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) GetCustomerSubnets() (resp []datatypes.Network_Customer_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) GetDatacenter() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) GetDiffieHellmanGroupDefault() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) GetDiffieHellmanGroupOptions() (resp []int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) GetEncryptionDefault() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) GetEncryptionOptions() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) GetInternalSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) GetKeylifeLimits() (resp []int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) GetObject() (resp datatypes.Network_Tunnel_Module_Context, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) GetParameterConfigurationsForCustomerView() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) GetPhaseOneKeylifeDefault() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) GetPhaseTwoKeylifeDefault() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) GetServiceSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) GetStaticRouteSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) GetTransactionHistory() (resp []datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) RemoveCustomerSubnetFromNetworkTunnel(subnetId *int) (resp bool, err error) {
	params := []interface{}{
		subnetId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) RemovePrivateSubnetFromNetworkTunnel(subnetId *int) (resp bool, err error) {
	params := []interface{}{
		subnetId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Tunnel_Module_Context) RemoveServiceSubnetFromNetworkTunnel(subnetId *int) (resp bool, err error) {
	params := []interface{}{
		subnetId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_Vlan struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkVlanService() Network_Vlan {
	return Network_Vlan{Session: r}
}

func (r *Network_Vlan) EditObject(templateObject *datatypes.Network_Vlan) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetAdditionalPrimarySubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetAttachedNetworkGateway() (resp datatypes.Network_Gateway, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetAttachedNetworkGatewayFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetAttachedNetworkGatewayVlan() (resp datatypes.Network_Gateway_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetCancelFailureReasons() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetDedicatedFirewallFlag() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetExtensionRouter() (resp datatypes.Hardware_Router, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetFirewallGuestNetworkComponents() (resp []datatypes.Network_Component_Firewall, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetFirewallInterfaces() (resp []datatypes.Network_Firewall_Module_Context_Interface, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetFirewallNetworkComponents() (resp []datatypes.Network_Component_Firewall, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetFirewallProtectableIpAddresses() (resp []datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetFirewallProtectableSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetFirewallRules() (resp []datatypes.Network_Vlan_Firewall_Rule, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetGuestNetworkComponents() (resp []datatypes.Virtual_Guest_Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetHighAvailabilityFirewallFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetLocalDiskStorageCapabilityFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetNetwork() (resp datatypes.Network, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetNetworkComponentTrunks() (resp []datatypes.Network_Component_Network_Vlan_Trunk, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetNetworkSpace() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetNetworkVlanFirewall() (resp datatypes.Network_Vlan_Firewall, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetObject() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetPrimaryRouter() (resp datatypes.Hardware_Router, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetPrimarySubnet() (resp datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetPrimarySubnetVersion6() (resp datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetPrimarySubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetPrivateNetworkGateways() (resp []datatypes.Network_Gateway, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetPrivateVlan() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetPrivateVlanByIpAddress(ipAddress *string) (resp datatypes.Network_Vlan, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetProtectedIpAddresses() (resp []datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetPublicNetworkGateways() (resp []datatypes.Network_Gateway, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetPublicVlanByFqdn(fqdn *string) (resp datatypes.Network_Vlan, err error) {
	params := []interface{}{
		fqdn,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetResourceGroupMember() (resp []datatypes.Resource_Group_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetResourceGroups() (resp []datatypes.Resource_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetReverseDomainRecords() (resp []datatypes.Dns_Domain, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetSanStorageCapabilityFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetScaleVlans() (resp []datatypes.Scale_Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetSecondaryRouter() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetSecondarySubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetTagReferences() (resp []datatypes.Tag_Reference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetTotalPrimaryIpAddressCount() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetType() (resp datatypes.Network_Vlan_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) GetVlanForIpAddress(ipAddress *string) (resp datatypes.Network_Vlan, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) SetTags(tags *string) (resp bool, err error) {
	params := []interface{}{
		tags,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan) UpdateFirewallIntraVlanCommunication(enabled *bool) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		enabled,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_Vlan_Firewall struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkVlanFirewallService() Network_Vlan_Firewall {
	return Network_Vlan_Firewall{Session: r}
}

func (r *Network_Vlan_Firewall) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan_Firewall) GetDatacenter() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan_Firewall) GetFirewallType() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan_Firewall) GetFullyQualifiedDomainName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan_Firewall) GetManagementCredentials() (resp datatypes.Software_Component_Password, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan_Firewall) GetNetworkFirewallUpdateRequests() (resp []datatypes.Network_Firewall_Update_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan_Firewall) GetNetworkVlan() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan_Firewall) GetNetworkVlans() (resp []datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan_Firewall) GetObject() (resp datatypes.Network_Vlan_Firewall, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan_Firewall) GetRules() (resp []datatypes.Network_Vlan_Firewall_Rule, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan_Firewall) GetTagReferences() (resp []datatypes.Tag_Reference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan_Firewall) RestoreDefaults() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan_Firewall) SetTags(tags *string) (resp bool, err error) {
	params := []interface{}{
		tags,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Network_Vlan_Firewall) UpdateRouteBypass(bypass *bool) (resp datatypes.Provisioning_Version1_Transaction, err error) {
	params := []interface{}{
		bypass,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Network_Vlan_Type struct {
	Session *Session
	Options
}

func (r *Session) GetNetworkVlanTypeService() Network_Vlan_Type {
	return Network_Vlan_Type{Session: r}
}

func (r *Network_Vlan_Type) GetObject() (resp datatypes.Network_Vlan_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
