/**
 * Copyright 2016-2024 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
 * the License. You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed
 * on an "AS IS" BASIS,WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and limitations under the License.
 */

// AUTOMATICALLY GENERATED CODE - DO NOT MODIFY

package services

import (
	"fmt"
	"strings"

	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/sl"
)

// The SoftLayer_Network_Application_Delivery_Controller data type models a single instance of an application delivery controller. Local properties are read only, except for a ”notes” property, which can be used to describe your application delivery controller service. The type's relational properties provide more information to the service's function and login information to the controller's backend management if advanced view is enabled.
type Network_Application_Delivery_Controller struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkApplicationDeliveryControllerService returns an instance of the Network_Application_Delivery_Controller SoftLayer service
func GetNetworkApplicationDeliveryControllerService(sess session.SLSession) Network_Application_Delivery_Controller {
	return Network_Application_Delivery_Controller{Session: sess}
}

func (r Network_Application_Delivery_Controller) Id(id int) Network_Application_Delivery_Controller {
	r.Options.Id = &id
	return r
}

func (r Network_Application_Delivery_Controller) Mask(mask string) Network_Application_Delivery_Controller {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_Application_Delivery_Controller) Filter(filter string) Network_Application_Delivery_Controller {
	r.Options.Filter = filter
	return r
}

func (r Network_Application_Delivery_Controller) Limit(limit int) Network_Application_Delivery_Controller {
	r.Options.Limit = &limit
	return r
}

func (r Network_Application_Delivery_Controller) Offset(offset int) Network_Application_Delivery_Controller {
	r.Options.Offset = &offset
	return r
}

// getObject retrieves the SoftLayer_Network_Application_Delivery_Controller object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Network_Application_Delivery_Controller service. You can only retrieve application delivery controllers that are associated with your SoftLayer customer account.
func (r Network_Application_Delivery_Controller) GetObject() (resp datatypes.Network_Application_Delivery_Controller, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Application_Delivery_Controller", "getObject", nil, &r.Options, &resp)
	return
}

// The SoftLayer_Network_Bandwidth_Version1_Allotment class provides methods and data structures necessary to work with an array of hardware objects associated with a single Bandwidth Pooling.
type Network_Bandwidth_Version1_Allotment struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkBandwidthVersion1AllotmentService returns an instance of the Network_Bandwidth_Version1_Allotment SoftLayer service
func GetNetworkBandwidthVersion1AllotmentService(sess session.SLSession) Network_Bandwidth_Version1_Allotment {
	return Network_Bandwidth_Version1_Allotment{Session: sess}
}

func (r Network_Bandwidth_Version1_Allotment) Id(id int) Network_Bandwidth_Version1_Allotment {
	r.Options.Id = &id
	return r
}

func (r Network_Bandwidth_Version1_Allotment) Mask(mask string) Network_Bandwidth_Version1_Allotment {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_Bandwidth_Version1_Allotment) Filter(filter string) Network_Bandwidth_Version1_Allotment {
	r.Options.Filter = filter
	return r
}

func (r Network_Bandwidth_Version1_Allotment) Limit(limit int) Network_Bandwidth_Version1_Allotment {
	r.Options.Limit = &limit
	return r
}

func (r Network_Bandwidth_Version1_Allotment) Offset(offset int) Network_Bandwidth_Version1_Allotment {
	r.Options.Offset = &offset
	return r
}

// Create a allotment for servers to pool bandwidth and avoid overages in billing if they use more than there allocated bandwidth.
func (r Network_Bandwidth_Version1_Allotment) CreateObject(templateObject *datatypes.Network_Bandwidth_Version1_Allotment) (resp datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Bandwidth_Version1_Allotment", "createObject", params, &r.Options, &resp)
	return
}

// Edit a bandwidth allotment's local properties. Currently you may only change an allotment's name. Use the [[SoftLayer_Network_Bandwidth_Version1_Allotment::reassignServers|reassignServers()]] and [[SoftLayer_Network_Bandwidth_Version1_Allotment::unassignServers|unassignServers()]] methods to move servers in and out of your allotments.
func (r Network_Bandwidth_Version1_Allotment) EditObject(templateObject *datatypes.Network_Bandwidth_Version1_Allotment) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Bandwidth_Version1_Allotment", "editObject", params, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Network_Bandwidth_Version1_Allotment object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Hardware service. You can only retrieve an allotment associated with the account that your portal user is assigned to.
func (r Network_Bandwidth_Version1_Allotment) GetObject() (resp datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Bandwidth_Version1_Allotment", "getObject", nil, &r.Options, &resp)
	return
}

// This will remove a bandwidth pooling from a customer's allotments by cancelling the billing item.  All servers in that allotment will get moved to the account's vpr.
func (r Network_Bandwidth_Version1_Allotment) RequestVdrCancellation() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Bandwidth_Version1_Allotment", "requestVdrCancellation", nil, &r.Options, &resp)
	return
}

// This data type models a purge event that occurs in caching server. It contains a reference to a mapping configuration, the path to execute the purge on, the status of the purge, and flag that enables saving the purge information for future use.
type Network_CdnMarketplace_Configuration_Cache_Purge struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkCdnMarketplaceConfigurationCachePurgeService returns an instance of the Network_CdnMarketplace_Configuration_Cache_Purge SoftLayer service
func GetNetworkCdnMarketplaceConfigurationCachePurgeService(sess session.SLSession) Network_CdnMarketplace_Configuration_Cache_Purge {
	return Network_CdnMarketplace_Configuration_Cache_Purge{Session: sess}
}

func (r Network_CdnMarketplace_Configuration_Cache_Purge) Id(id int) Network_CdnMarketplace_Configuration_Cache_Purge {
	r.Options.Id = &id
	return r
}

func (r Network_CdnMarketplace_Configuration_Cache_Purge) Mask(mask string) Network_CdnMarketplace_Configuration_Cache_Purge {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_CdnMarketplace_Configuration_Cache_Purge) Filter(filter string) Network_CdnMarketplace_Configuration_Cache_Purge {
	r.Options.Filter = filter
	return r
}

func (r Network_CdnMarketplace_Configuration_Cache_Purge) Limit(limit int) Network_CdnMarketplace_Configuration_Cache_Purge {
	r.Options.Limit = &limit
	return r
}

func (r Network_CdnMarketplace_Configuration_Cache_Purge) Offset(offset int) Network_CdnMarketplace_Configuration_Cache_Purge {
	r.Options.Offset = &offset
	return r
}

// no documentation yet
func (r Network_CdnMarketplace_Configuration_Cache_Purge) CreatePurge(uniqueId *string, path *string) (resp []datatypes.Container_Network_CdnMarketplace_Configuration_Cache_Purge, err error) {
	params := []interface{}{
		uniqueId,
		path,
	}
	err = r.Session.DoRequest("SoftLayer_Network_CdnMarketplace_Configuration_Cache_Purge", "createPurge", params, &r.Options, &resp)
	return
}

// This data type represents the mapping Configuration settings for enabling CDN services. Each instance contains a reference to a CDN account, and CDN configuration properties such as a domain, an origin host and its port, a cname we generate, a cname the vendor generates, and a status. Other properties include the type of content to be cached (static or dynamic), the origin type (a host server or an object storage account), and the protocol to be used for caching.
type Network_CdnMarketplace_Configuration_Mapping struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkCdnMarketplaceConfigurationMappingService returns an instance of the Network_CdnMarketplace_Configuration_Mapping SoftLayer service
func GetNetworkCdnMarketplaceConfigurationMappingService(sess session.SLSession) Network_CdnMarketplace_Configuration_Mapping {
	return Network_CdnMarketplace_Configuration_Mapping{Session: sess}
}

func (r Network_CdnMarketplace_Configuration_Mapping) Id(id int) Network_CdnMarketplace_Configuration_Mapping {
	r.Options.Id = &id
	return r
}

func (r Network_CdnMarketplace_Configuration_Mapping) Mask(mask string) Network_CdnMarketplace_Configuration_Mapping {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_CdnMarketplace_Configuration_Mapping) Filter(filter string) Network_CdnMarketplace_Configuration_Mapping {
	r.Options.Filter = filter
	return r
}

func (r Network_CdnMarketplace_Configuration_Mapping) Limit(limit int) Network_CdnMarketplace_Configuration_Mapping {
	r.Options.Limit = &limit
	return r
}

func (r Network_CdnMarketplace_Configuration_Mapping) Offset(offset int) Network_CdnMarketplace_Configuration_Mapping {
	r.Options.Offset = &offset
	return r
}

// no documentation yet
func (r Network_CdnMarketplace_Configuration_Mapping) CreateDomainMapping(input *datatypes.Container_Network_CdnMarketplace_Configuration_Input) (resp []datatypes.Container_Network_CdnMarketplace_Configuration_Mapping, err error) {
	params := []interface{}{
		input,
	}
	err = r.Session.DoRequest("SoftLayer_Network_CdnMarketplace_Configuration_Mapping", "createDomainMapping", params, &r.Options, &resp)
	return
}

// no documentation yet
func (r Network_CdnMarketplace_Configuration_Mapping) DeleteDomainMapping(uniqueId *string) (resp []datatypes.Container_Network_CdnMarketplace_Configuration_Mapping, err error) {
	params := []interface{}{
		uniqueId,
	}
	err = r.Session.DoRequest("SoftLayer_Network_CdnMarketplace_Configuration_Mapping", "deleteDomainMapping", params, &r.Options, &resp)
	return
}

// no documentation yet
func (r Network_CdnMarketplace_Configuration_Mapping) ListDomainMappingByUniqueId(uniqueId *string) (resp []datatypes.Container_Network_CdnMarketplace_Configuration_Mapping, err error) {
	params := []interface{}{
		uniqueId,
	}
	err = r.Session.DoRequest("SoftLayer_Network_CdnMarketplace_Configuration_Mapping", "listDomainMappingByUniqueId", params, &r.Options, &resp)
	return
}

// no documentation yet
func (r Network_CdnMarketplace_Configuration_Mapping) ListDomainMappings() (resp []datatypes.Container_Network_CdnMarketplace_Configuration_Mapping, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_CdnMarketplace_Configuration_Mapping", "listDomainMappings", nil, &r.Options, &resp)
	return
}

// no documentation yet
func (r Network_CdnMarketplace_Configuration_Mapping) UpdateDomainMapping(input *datatypes.Container_Network_CdnMarketplace_Configuration_Input) (resp []datatypes.Container_Network_CdnMarketplace_Configuration_Mapping, err error) {
	params := []interface{}{
		input,
	}
	err = r.Session.DoRequest("SoftLayer_Network_CdnMarketplace_Configuration_Mapping", "updateDomainMapping", params, &r.Options, &resp)
	return
}

// no documentation yet
type Network_CdnMarketplace_Configuration_Mapping_Path struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkCdnMarketplaceConfigurationMappingPathService returns an instance of the Network_CdnMarketplace_Configuration_Mapping_Path SoftLayer service
func GetNetworkCdnMarketplaceConfigurationMappingPathService(sess session.SLSession) Network_CdnMarketplace_Configuration_Mapping_Path {
	return Network_CdnMarketplace_Configuration_Mapping_Path{Session: sess}
}

func (r Network_CdnMarketplace_Configuration_Mapping_Path) Id(id int) Network_CdnMarketplace_Configuration_Mapping_Path {
	r.Options.Id = &id
	return r
}

func (r Network_CdnMarketplace_Configuration_Mapping_Path) Mask(mask string) Network_CdnMarketplace_Configuration_Mapping_Path {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_CdnMarketplace_Configuration_Mapping_Path) Filter(filter string) Network_CdnMarketplace_Configuration_Mapping_Path {
	r.Options.Filter = filter
	return r
}

func (r Network_CdnMarketplace_Configuration_Mapping_Path) Limit(limit int) Network_CdnMarketplace_Configuration_Mapping_Path {
	r.Options.Limit = &limit
	return r
}

func (r Network_CdnMarketplace_Configuration_Mapping_Path) Offset(offset int) Network_CdnMarketplace_Configuration_Mapping_Path {
	r.Options.Offset = &offset
	return r
}

// no documentation yet
func (r Network_CdnMarketplace_Configuration_Mapping_Path) CreateOriginPath(input *datatypes.Container_Network_CdnMarketplace_Configuration_Input) (resp []datatypes.Container_Network_CdnMarketplace_Configuration_Mapping_Path, err error) {
	params := []interface{}{
		input,
	}
	err = r.Session.DoRequest("SoftLayer_Network_CdnMarketplace_Configuration_Mapping_Path", "createOriginPath", params, &r.Options, &resp)
	return
}

// no documentation yet
func (r Network_CdnMarketplace_Configuration_Mapping_Path) DeleteOriginPath(uniqueId *string, path *string) (resp string, err error) {
	params := []interface{}{
		uniqueId,
		path,
	}
	err = r.Session.DoRequest("SoftLayer_Network_CdnMarketplace_Configuration_Mapping_Path", "deleteOriginPath", params, &r.Options, &resp)
	return
}

// no documentation yet
func (r Network_CdnMarketplace_Configuration_Mapping_Path) ListOriginPath(uniqueId *string) (resp []datatypes.Container_Network_CdnMarketplace_Configuration_Mapping_Path, err error) {
	params := []interface{}{
		uniqueId,
	}
	err = r.Session.DoRequest("SoftLayer_Network_CdnMarketplace_Configuration_Mapping_Path", "listOriginPath", params, &r.Options, &resp)
	return
}

// This Metrics class provides methods to get CDN metrics based on account or mapping unique id.
type Network_CdnMarketplace_Metrics struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkCdnMarketplaceMetricsService returns an instance of the Network_CdnMarketplace_Metrics SoftLayer service
func GetNetworkCdnMarketplaceMetricsService(sess session.SLSession) Network_CdnMarketplace_Metrics {
	return Network_CdnMarketplace_Metrics{Session: sess}
}

func (r Network_CdnMarketplace_Metrics) Id(id int) Network_CdnMarketplace_Metrics {
	r.Options.Id = &id
	return r
}

func (r Network_CdnMarketplace_Metrics) Mask(mask string) Network_CdnMarketplace_Metrics {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_CdnMarketplace_Metrics) Filter(filter string) Network_CdnMarketplace_Metrics {
	r.Options.Filter = filter
	return r
}

func (r Network_CdnMarketplace_Metrics) Limit(limit int) Network_CdnMarketplace_Metrics {
	r.Options.Limit = &limit
	return r
}

func (r Network_CdnMarketplace_Metrics) Offset(offset int) Network_CdnMarketplace_Metrics {
	r.Options.Offset = &offset
	return r
}

// no documentation yet
func (r Network_CdnMarketplace_Metrics) GetMappingUsageMetrics(mappingUniqueId *string, startDate *int, endDate *int, frequency *string) (resp []datatypes.Container_Network_CdnMarketplace_Metrics, err error) {
	params := []interface{}{
		mappingUniqueId,
		startDate,
		endDate,
		frequency,
	}
	err = r.Session.DoRequest("SoftLayer_Network_CdnMarketplace_Metrics", "getMappingUsageMetrics", params, &r.Options, &resp)
	return
}

// Every piece of hardware running in SoftLayer's datacenters connected to the public, private, or management networks (where applicable) have a corresponding network component. These network components are modeled by the SoftLayer_Network_Component data type. These data types reflect the servers' local ethernet and remote management interfaces.
type Network_Component struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkComponentService returns an instance of the Network_Component SoftLayer service
func GetNetworkComponentService(sess session.SLSession) Network_Component {
	return Network_Component{Session: sess}
}

func (r Network_Component) Id(id int) Network_Component {
	r.Options.Id = &id
	return r
}

func (r Network_Component) Mask(mask string) Network_Component {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_Component) Filter(filter string) Network_Component {
	r.Options.Filter = filter
	return r
}

func (r Network_Component) Limit(limit int) Network_Component {
	r.Options.Limit = &limit
	return r
}

func (r Network_Component) Offset(offset int) Network_Component {
	r.Options.Offset = &offset
	return r
}

// Add VLANs as trunks to a network component. The VLANs given must be assigned to your account and belong to the same pod in which this network component and its hardware reside. The current native VLAN cannot be added as a trunk.
//
// This method should be called on a network component of assigned hardware. A current list of VLAN trunks for a network component on a customer server can be found at 'uplinkComponent->networkVlanTrunks'.
//
// This method returns an array of SoftLayer_Network_Vlans which were added as trunks. Any requested VLANs which are already trunked will be ignored and will not be returned.
//
// Affected VLANs will not yet be operational as trunks on the network upon return of this call, but activation will have been scheduled and should be considered imminent. The trunking records associated with the affected VLANs will maintain an 'isUpdating' value of '1' so long as this is the case.
//
// Note that in the event of an "internal system error" some VLANs may still have been affected and scheduled for activation.
func (r Network_Component) AddNetworkVlanTrunks(networkVlans []datatypes.Network_Vlan) (resp []datatypes.Network_Vlan, err error) {
	params := []interface{}{
		networkVlans,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Component", "addNetworkVlanTrunks", params, &r.Options, &resp)
	return
}

// no documentation yet
func (r Network_Component) GetObject() (resp datatypes.Network_Component, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Component", "getObject", nil, &r.Options, &resp)
	return
}

// Remove one or more VLANs currently attached as trunks to this network component.
//
// If any VLANs are given which are not attached as trunks, they will be ignored.
//
// This method should be called on a network component of assigned hardware. A current list of VLAN trunks for a network component on a customer server can be found at 'uplinkComponent->networkVlanTrunks'.
//
// This method returns an array of SoftLayer_Network_Vlans which will be removed as trunks. Any requested VLANs which were not trunked will be ignored and will not be returned.
//
// Affected VLANs will not yet be removed as trunks upon return of this call, but deactivation and removal will have been scheduled and should be considered imminent. The trunking records associated with the affected VLANs will maintain an 'isUpdating' value of '1' so long as this is the case.
//
// Note that in the event of a "pending API request" error some VLANs may still have been affected and scheduled for deactivation.
func (r Network_Component) RemoveNetworkVlanTrunks(networkVlans []datatypes.Network_Vlan) (resp []datatypes.Network_Vlan, err error) {
	params := []interface{}{
		networkVlans,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Component", "removeNetworkVlanTrunks", params, &r.Options, &resp)
	return
}

// The SoftLayer_Network_Component_Firewall data type contains general information relating to a single SoftLayer network component firewall. This is the object which ties the running rules to a specific downstream server. Use the [[SoftLayer Network Firewall Template]] service to pull SoftLayer recommended rule set templates. Use the [[SoftLayer Network Firewall Update Request]] service to submit a firewall update request.
type Network_Component_Firewall struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkComponentFirewallService returns an instance of the Network_Component_Firewall SoftLayer service
func GetNetworkComponentFirewallService(sess session.SLSession) Network_Component_Firewall {
	return Network_Component_Firewall{Session: sess}
}

func (r Network_Component_Firewall) Id(id int) Network_Component_Firewall {
	r.Options.Id = &id
	return r
}

func (r Network_Component_Firewall) Mask(mask string) Network_Component_Firewall {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_Component_Firewall) Filter(filter string) Network_Component_Firewall {
	r.Options.Filter = filter
	return r
}

func (r Network_Component_Firewall) Limit(limit int) Network_Component_Firewall {
	r.Options.Limit = &limit
	return r
}

func (r Network_Component_Firewall) Offset(offset int) Network_Component_Firewall {
	r.Options.Offset = &offset
	return r
}

// getObject returns a SoftLayer_Network_Firewall_Module_Context_Interface_AccessControlList_Network_Component object. You can only get objects for servers attached to your account that have a network firewall enabled.
func (r Network_Component_Firewall) GetObject() (resp datatypes.Network_Component_Firewall, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Component_Firewall", "getObject", nil, &r.Options, &resp)
	return
}

// Retrieve The currently running rule set of this network component firewall.
func (r Network_Component_Firewall) GetRules() (resp []datatypes.Network_Component_Firewall_Rule, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Component_Firewall", "getRules", nil, &r.Options, &resp)
	return
}

// The SoftLayer_Network_Customer_Subnet data type contains general information relating to a single customer subnet (remote).
type Network_Customer_Subnet struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkCustomerSubnetService returns an instance of the Network_Customer_Subnet SoftLayer service
func GetNetworkCustomerSubnetService(sess session.SLSession) Network_Customer_Subnet {
	return Network_Customer_Subnet{Session: sess}
}

func (r Network_Customer_Subnet) Id(id int) Network_Customer_Subnet {
	r.Options.Id = &id
	return r
}

func (r Network_Customer_Subnet) Mask(mask string) Network_Customer_Subnet {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_Customer_Subnet) Filter(filter string) Network_Customer_Subnet {
	r.Options.Filter = filter
	return r
}

func (r Network_Customer_Subnet) Limit(limit int) Network_Customer_Subnet {
	r.Options.Limit = &limit
	return r
}

func (r Network_Customer_Subnet) Offset(offset int) Network_Customer_Subnet {
	r.Options.Offset = &offset
	return r
}

// For IPSec network tunnels, customers can create their local subnets using this method.  After the customer is created successfully, the customer subnet can then be added to the IPSec network tunnel.
func (r Network_Customer_Subnet) CreateObject(templateObject *datatypes.Network_Customer_Subnet) (resp datatypes.Network_Customer_Subnet, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Customer_Subnet", "createObject", params, &r.Options, &resp)
	return
}

// The SoftLayer_Network_Firewall_Update_Request data type contains information relating to a SoftLayer network firewall update request. Use the [[SoftLayer Network Component Firewall]] service to view current rules. Use the [[SoftLayer Network Firewall Template]] service to pull SoftLayer recommended rule set templates.
type Network_Firewall_Update_Request struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkFirewallUpdateRequestService returns an instance of the Network_Firewall_Update_Request SoftLayer service
func GetNetworkFirewallUpdateRequestService(sess session.SLSession) Network_Firewall_Update_Request {
	return Network_Firewall_Update_Request{Session: sess}
}

func (r Network_Firewall_Update_Request) Id(id int) Network_Firewall_Update_Request {
	r.Options.Id = &id
	return r
}

func (r Network_Firewall_Update_Request) Mask(mask string) Network_Firewall_Update_Request {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_Firewall_Update_Request) Filter(filter string) Network_Firewall_Update_Request {
	r.Options.Filter = filter
	return r
}

func (r Network_Firewall_Update_Request) Limit(limit int) Network_Firewall_Update_Request {
	r.Options.Limit = &limit
	return r
}

func (r Network_Firewall_Update_Request) Offset(offset int) Network_Firewall_Update_Request {
	r.Options.Offset = &offset
	return r
}

// Create a new firewall update request. If the SoftLayer_Network_Firewall_Update_Request object passed to this function has no rule, the firewall be set to bypass state and all the existing firewall rule(s) will be deleted.
//
// ”createObject” returns a Boolean ”true” on successful object creation or ”false” if your firewall update request was unable to be created.
func (r Network_Firewall_Update_Request) CreateObject(templateObject *datatypes.Network_Firewall_Update_Request) (resp datatypes.Network_Firewall_Update_Request, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Firewall_Update_Request", "createObject", params, &r.Options, &resp)
	return
}

// The SoftLayer_Network_LBaaS_HealthMonitor type presents a structure containing attributes of a health monitor object associated with load balancer instance. Note that the relationship between backend (pool) and health monitor is N-to-1, especially that the pools object associated with a health monitor must have the same pair of protocol and port. Example: frontend FA: http, 80   - backend BA: tcp, 3456 - healthmonitor HM_tcp3456 frontend FB: https, 443 - backend BB: tcp, 3456 - healthmonitor HM_tcp3456 In above example both backends BA and BB share the same healthmonitor HM_tcp3456
type Network_LBaaS_HealthMonitor struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkLBaaSHealthMonitorService returns an instance of the Network_LBaaS_HealthMonitor SoftLayer service
func GetNetworkLBaaSHealthMonitorService(sess session.SLSession) Network_LBaaS_HealthMonitor {
	return Network_LBaaS_HealthMonitor{Session: sess}
}

func (r Network_LBaaS_HealthMonitor) Id(id int) Network_LBaaS_HealthMonitor {
	r.Options.Id = &id
	return r
}

func (r Network_LBaaS_HealthMonitor) Mask(mask string) Network_LBaaS_HealthMonitor {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_LBaaS_HealthMonitor) Filter(filter string) Network_LBaaS_HealthMonitor {
	r.Options.Filter = filter
	return r
}

func (r Network_LBaaS_HealthMonitor) Limit(limit int) Network_LBaaS_HealthMonitor {
	r.Options.Limit = &limit
	return r
}

func (r Network_LBaaS_HealthMonitor) Offset(offset int) Network_LBaaS_HealthMonitor {
	r.Options.Offset = &offset
	return r
}

// Update load balancers health monitor and return load balancer object with listeners (frontend), pools (backend), health monitor server instances (members) and datacenter populated
func (r Network_LBaaS_HealthMonitor) UpdateLoadBalancerHealthMonitors(loadBalancerUuid *string, healthMonitorConfigurations []datatypes.Network_LBaaS_LoadBalancerHealthMonitorConfiguration) (resp datatypes.Network_LBaaS_LoadBalancer, err error) {
	params := []interface{}{
		loadBalancerUuid,
		healthMonitorConfigurations,
	}
	err = r.Session.DoRequest("SoftLayer_Network_LBaaS_HealthMonitor", "updateLoadBalancerHealthMonitors", params, &r.Options, &resp)
	return
}

// The SoftLayer_Network_LBaaS_L7Member represents the backend member for a L7 pool. It can be either a virtual server or a bare metal machine.
type Network_LBaaS_L7Member struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkLBaaSL7MemberService returns an instance of the Network_LBaaS_L7Member SoftLayer service
func GetNetworkLBaaSL7MemberService(sess session.SLSession) Network_LBaaS_L7Member {
	return Network_LBaaS_L7Member{Session: sess}
}

func (r Network_LBaaS_L7Member) Id(id int) Network_LBaaS_L7Member {
	r.Options.Id = &id
	return r
}

func (r Network_LBaaS_L7Member) Mask(mask string) Network_LBaaS_L7Member {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_LBaaS_L7Member) Filter(filter string) Network_LBaaS_L7Member {
	r.Options.Filter = filter
	return r
}

func (r Network_LBaaS_L7Member) Limit(limit int) Network_LBaaS_L7Member {
	r.Options.Limit = &limit
	return r
}

func (r Network_LBaaS_L7Member) Offset(offset int) Network_LBaaS_L7Member {
	r.Options.Offset = &offset
	return r
}

// Add server instances as members to a L7pool and return the LoadBalancer Object with listeners, pools and members populated
func (r Network_LBaaS_L7Member) AddL7PoolMembers(l7PoolUuid *string, memberInstances []datatypes.Network_LBaaS_L7Member) (resp datatypes.Network_LBaaS_LoadBalancer, err error) {
	params := []interface{}{
		l7PoolUuid,
		memberInstances,
	}
	err = r.Session.DoRequest("SoftLayer_Network_LBaaS_L7Member", "addL7PoolMembers", params, &r.Options, &resp)
	return
}

// Delete given members from load balancer and return load balancer object with listeners, pools and members populated
func (r Network_LBaaS_L7Member) DeleteL7PoolMembers(l7PoolUuid *string, memberUuids []string) (resp datatypes.Network_LBaaS_LoadBalancer, err error) {
	params := []interface{}{
		l7PoolUuid,
		memberUuids,
	}
	err = r.Session.DoRequest("SoftLayer_Network_LBaaS_L7Member", "deleteL7PoolMembers", params, &r.Options, &resp)
	return
}

// The SoftLayer_Network_LBaaS_L7Policy represents the policy for a listener.
type Network_LBaaS_L7Policy struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkLBaaSL7PolicyService returns an instance of the Network_LBaaS_L7Policy SoftLayer service
func GetNetworkLBaaSL7PolicyService(sess session.SLSession) Network_LBaaS_L7Policy {
	return Network_LBaaS_L7Policy{Session: sess}
}

func (r Network_LBaaS_L7Policy) Id(id int) Network_LBaaS_L7Policy {
	r.Options.Id = &id
	return r
}

func (r Network_LBaaS_L7Policy) Mask(mask string) Network_LBaaS_L7Policy {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_LBaaS_L7Policy) Filter(filter string) Network_LBaaS_L7Policy {
	r.Options.Filter = filter
	return r
}

func (r Network_LBaaS_L7Policy) Limit(limit int) Network_LBaaS_L7Policy {
	r.Options.Limit = &limit
	return r
}

func (r Network_LBaaS_L7Policy) Offset(offset int) Network_LBaaS_L7Policy {
	r.Options.Offset = &offset
	return r
}

// This function creates multiple policies with rules for the given listener.
func (r Network_LBaaS_L7Policy) AddL7Policies(listenerUuid *string, policiesRules []datatypes.Network_LBaaS_PolicyRule) (resp datatypes.Network_LBaaS_LoadBalancer, err error) {
	params := []interface{}{
		listenerUuid,
		policiesRules,
	}
	err = r.Session.DoRequest("SoftLayer_Network_LBaaS_L7Policy", "addL7Policies", params, &r.Options, &resp)
	return
}

// Deletes a l7 policy instance and the rules associated with the policy
func (r Network_LBaaS_L7Policy) DeleteObject() (resp datatypes.Network_LBaaS_LoadBalancer, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_LBaaS_L7Policy", "deleteObject", nil, &r.Options, &resp)
	return
}

// Edit a l7 policy instance's properties
func (r Network_LBaaS_L7Policy) EditObject(templateObject *datatypes.Network_LBaaS_L7Policy) (resp datatypes.Network_LBaaS_LoadBalancer, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Network_LBaaS_L7Policy", "editObject", params, &r.Options, &resp)
	return
}

// Retrieve
func (r Network_LBaaS_L7Policy) GetL7Rules() (resp []datatypes.Network_LBaaS_L7Rule, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_LBaaS_L7Policy", "getL7Rules", nil, &r.Options, &resp)
	return
}

// no documentation yet
func (r Network_LBaaS_L7Policy) GetObject() (resp datatypes.Network_LBaaS_L7Policy, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_LBaaS_L7Policy", "getObject", nil, &r.Options, &resp)
	return
}

// The SoftLayer_Network_LBaaS_L7Pool type presents a structure containing attributes of a load balancer's L7 pool such as the protocol, and the load balancing algorithm used. L7 pool is used for redirect_pool action of the L7 policy and is different from the default pool
type Network_LBaaS_L7Pool struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkLBaaSL7PoolService returns an instance of the Network_LBaaS_L7Pool SoftLayer service
func GetNetworkLBaaSL7PoolService(sess session.SLSession) Network_LBaaS_L7Pool {
	return Network_LBaaS_L7Pool{Session: sess}
}

func (r Network_LBaaS_L7Pool) Id(id int) Network_LBaaS_L7Pool {
	r.Options.Id = &id
	return r
}

func (r Network_LBaaS_L7Pool) Mask(mask string) Network_LBaaS_L7Pool {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_LBaaS_L7Pool) Filter(filter string) Network_LBaaS_L7Pool {
	r.Options.Filter = filter
	return r
}

func (r Network_LBaaS_L7Pool) Limit(limit int) Network_LBaaS_L7Pool {
	r.Options.Limit = &limit
	return r
}

func (r Network_LBaaS_L7Pool) Offset(offset int) Network_LBaaS_L7Pool {
	r.Options.Offset = &offset
	return r
}

// Create a backend to be used for L7 load balancing. This L7 pool has backend protocol, L7 members, L7 health monitor and session affinity. L7 pool is associated with L7 policies.
func (r Network_LBaaS_L7Pool) CreateL7Pool(loadBalancerUuid *string, l7Pool *datatypes.Network_LBaaS_L7Pool, l7Members []datatypes.Network_LBaaS_L7Member, l7HealthMonitor *datatypes.Network_LBaaS_L7HealthMonitor, l7SessionAffinity *datatypes.Network_LBaaS_L7SessionAffinity) (resp datatypes.Network_LBaaS_LoadBalancer, err error) {
	params := []interface{}{
		loadBalancerUuid,
		l7Pool,
		l7Members,
		l7HealthMonitor,
		l7SessionAffinity,
	}
	err = r.Session.DoRequest("SoftLayer_Network_LBaaS_L7Pool", "createL7Pool", params, &r.Options, &resp)
	return
}

// Deletes an existing L7 pool along with L7 members, L7 health monitor, and L7 session affinity.
func (r Network_LBaaS_L7Pool) DeleteObject() (resp datatypes.Network_LBaaS_LoadBalancer, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_LBaaS_L7Pool", "deleteObject", nil, &r.Options, &resp)
	return
}

// Retrieve
func (r Network_LBaaS_L7Pool) GetL7HealthMonitor() (resp datatypes.Network_LBaaS_L7HealthMonitor, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_LBaaS_L7Pool", "getL7HealthMonitor", nil, &r.Options, &resp)
	return
}

// Retrieve
func (r Network_LBaaS_L7Pool) GetL7Members() (resp []datatypes.Network_LBaaS_L7Member, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_LBaaS_L7Pool", "getL7Members", nil, &r.Options, &resp)
	return
}

// Retrieve
func (r Network_LBaaS_L7Pool) GetL7SessionAffinity() (resp datatypes.Network_LBaaS_L7SessionAffinity, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_LBaaS_L7Pool", "getL7SessionAffinity", nil, &r.Options, &resp)
	return
}

// no documentation yet
func (r Network_LBaaS_L7Pool) GetObject() (resp datatypes.Network_LBaaS_L7Pool, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_LBaaS_L7Pool", "getObject", nil, &r.Options, &resp)
	return
}

// Updates an existing L7 pool, L7 health monitor and L7 session affinity.
func (r Network_LBaaS_L7Pool) UpdateL7Pool(l7PoolUuid *string, l7Pool *datatypes.Network_LBaaS_L7Pool, l7HealthMonitor *datatypes.Network_LBaaS_L7HealthMonitor, l7SessionAffinity *datatypes.Network_LBaaS_L7SessionAffinity) (resp datatypes.Network_LBaaS_LoadBalancer, err error) {
	params := []interface{}{
		l7PoolUuid,
		l7Pool,
		l7HealthMonitor,
		l7SessionAffinity,
	}
	err = r.Session.DoRequest("SoftLayer_Network_LBaaS_L7Pool", "updateL7Pool", params, &r.Options, &resp)
	return
}

// The SoftLayer_Network_LBaaS_L7Rule represents the Rules that can be attached to a a L7 policy.
type Network_LBaaS_L7Rule struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkLBaaSL7RuleService returns an instance of the Network_LBaaS_L7Rule SoftLayer service
func GetNetworkLBaaSL7RuleService(sess session.SLSession) Network_LBaaS_L7Rule {
	return Network_LBaaS_L7Rule{Session: sess}
}

func (r Network_LBaaS_L7Rule) Id(id int) Network_LBaaS_L7Rule {
	r.Options.Id = &id
	return r
}

func (r Network_LBaaS_L7Rule) Mask(mask string) Network_LBaaS_L7Rule {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_LBaaS_L7Rule) Filter(filter string) Network_LBaaS_L7Rule {
	r.Options.Filter = filter
	return r
}

func (r Network_LBaaS_L7Rule) Limit(limit int) Network_LBaaS_L7Rule {
	r.Options.Limit = &limit
	return r
}

func (r Network_LBaaS_L7Rule) Offset(offset int) Network_LBaaS_L7Rule {
	r.Options.Offset = &offset
	return r
}

// This function creates and adds multiple Rules to a given L7 policy with all the details provided for rules
func (r Network_LBaaS_L7Rule) AddL7Rules(policyUuid *string, rules []datatypes.Network_LBaaS_L7Rule) (resp datatypes.Network_LBaaS_LoadBalancer, err error) {
	params := []interface{}{
		policyUuid,
		rules,
	}
	err = r.Session.DoRequest("SoftLayer_Network_LBaaS_L7Rule", "addL7Rules", params, &r.Options, &resp)
	return
}

// This function deletes multiple rules aassociated with the same policy.
func (r Network_LBaaS_L7Rule) DeleteL7Rules(policyUuid *string, ruleUuids []string) (resp datatypes.Network_LBaaS_LoadBalancer, err error) {
	params := []interface{}{
		policyUuid,
		ruleUuids,
	}
	err = r.Session.DoRequest("SoftLayer_Network_LBaaS_L7Rule", "deleteL7Rules", params, &r.Options, &resp)
	return
}

// The SoftLayer_Network_LBaaS_Listener type presents a data structure for a load balancers listener, also called frontend.
type Network_LBaaS_Listener struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkLBaaSListenerService returns an instance of the Network_LBaaS_Listener SoftLayer service
func GetNetworkLBaaSListenerService(sess session.SLSession) Network_LBaaS_Listener {
	return Network_LBaaS_Listener{Session: sess}
}

func (r Network_LBaaS_Listener) Id(id int) Network_LBaaS_Listener {
	r.Options.Id = &id
	return r
}

func (r Network_LBaaS_Listener) Mask(mask string) Network_LBaaS_Listener {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_LBaaS_Listener) Filter(filter string) Network_LBaaS_Listener {
	r.Options.Filter = filter
	return r
}

func (r Network_LBaaS_Listener) Limit(limit int) Network_LBaaS_Listener {
	r.Options.Limit = &limit
	return r
}

func (r Network_LBaaS_Listener) Offset(offset int) Network_LBaaS_Listener {
	r.Options.Offset = &offset
	return r
}

// Delete load balancers front- and backend protocols and return load balancer object with listeners (frontend), pools (backend), server instances (members) and datacenter populated.
func (r Network_LBaaS_Listener) DeleteLoadBalancerProtocols(loadBalancerUuid *string, listenerUuids []string) (resp datatypes.Network_LBaaS_LoadBalancer, err error) {
	params := []interface{}{
		loadBalancerUuid,
		listenerUuids,
	}
	err = r.Session.DoRequest("SoftLayer_Network_LBaaS_Listener", "deleteLoadBalancerProtocols", params, &r.Options, &resp)
	return
}

// Retrieve
func (r Network_LBaaS_Listener) GetL7Policies() (resp []datatypes.Network_LBaaS_L7Policy, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_LBaaS_Listener", "getL7Policies", nil, &r.Options, &resp)
	return
}

// Update (create) load balancers front- and backend protocols and return load balancer object with listeners (frontend), pools (backend), server instances (members) and datacenter populated. Note if a protocolConfiguration has no listenerUuid set, this function will create the specified front- and backend accordingly. Otherwise the given front- and backend will be updated with the new protocol and port.
func (r Network_LBaaS_Listener) UpdateLoadBalancerProtocols(loadBalancerUuid *string, protocolConfigurations []datatypes.Network_LBaaS_LoadBalancerProtocolConfiguration) (resp datatypes.Network_LBaaS_LoadBalancer, err error) {
	params := []interface{}{
		loadBalancerUuid,
		protocolConfigurations,
	}
	err = r.Session.DoRequest("SoftLayer_Network_LBaaS_Listener", "updateLoadBalancerProtocols", params, &r.Options, &resp)
	return
}

// The SoftLayer_Network_LBaaS_LoadBalancer type presents a structure containing attributes of a load balancer, and its related objects including listeners, pools and members.
type Network_LBaaS_LoadBalancer struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkLBaaSLoadBalancerService returns an instance of the Network_LBaaS_LoadBalancer SoftLayer service
func GetNetworkLBaaSLoadBalancerService(sess session.SLSession) Network_LBaaS_LoadBalancer {
	return Network_LBaaS_LoadBalancer{Session: sess}
}

func (r Network_LBaaS_LoadBalancer) Id(id int) Network_LBaaS_LoadBalancer {
	r.Options.Id = &id
	return r
}

func (r Network_LBaaS_LoadBalancer) Mask(mask string) Network_LBaaS_LoadBalancer {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_LBaaS_LoadBalancer) Filter(filter string) Network_LBaaS_LoadBalancer {
	r.Options.Filter = filter
	return r
}

func (r Network_LBaaS_LoadBalancer) Limit(limit int) Network_LBaaS_LoadBalancer {
	r.Options.Limit = &limit
	return r
}

func (r Network_LBaaS_LoadBalancer) Offset(offset int) Network_LBaaS_LoadBalancer {
	r.Options.Offset = &offset
	return r
}

// Cancel a load balancer with the given uuid. The billing system will execute the deletion of load balancer and all objects associated with it such as load balancer appliances, listeners, pools and members in the background.
func (r Network_LBaaS_LoadBalancer) CancelLoadBalancer(uuid *string) (resp bool, err error) {
	params := []interface{}{
		uuid,
	}
	err = r.Session.DoRequest("SoftLayer_Network_LBaaS_LoadBalancer", "cancelLoadBalancer", params, &r.Options, &resp)
	return
}

// Return all existing load balancers
func (r Network_LBaaS_LoadBalancer) GetAllObjects() (resp []datatypes.Network_LBaaS_LoadBalancer, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_LBaaS_LoadBalancer", "getAllObjects", nil, &r.Options, &resp)
	return
}

// no documentation yet
func (r Network_LBaaS_LoadBalancer) GetObject() (resp datatypes.Network_LBaaS_LoadBalancer, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_LBaaS_LoadBalancer", "getObject", nil, &r.Options, &resp)
	return
}

// The SoftLayer_Network_LBaaS_Member represents the backend member for a load balancer. It can be either a virtual server or a bare metal machine.
type Network_LBaaS_Member struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkLBaaSMemberService returns an instance of the Network_LBaaS_Member SoftLayer service
func GetNetworkLBaaSMemberService(sess session.SLSession) Network_LBaaS_Member {
	return Network_LBaaS_Member{Session: sess}
}

func (r Network_LBaaS_Member) Id(id int) Network_LBaaS_Member {
	r.Options.Id = &id
	return r
}

func (r Network_LBaaS_Member) Mask(mask string) Network_LBaaS_Member {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_LBaaS_Member) Filter(filter string) Network_LBaaS_Member {
	r.Options.Filter = filter
	return r
}

func (r Network_LBaaS_Member) Limit(limit int) Network_LBaaS_Member {
	r.Options.Limit = &limit
	return r
}

func (r Network_LBaaS_Member) Offset(offset int) Network_LBaaS_Member {
	r.Options.Offset = &offset
	return r
}

// Add server instances as members to load balancer and return it with listeners, pools and members populated
func (r Network_LBaaS_Member) AddLoadBalancerMembers(loadBalancerUuid *string, serverInstances []datatypes.Network_LBaaS_LoadBalancerServerInstanceInfo) (resp datatypes.Network_LBaaS_LoadBalancer, err error) {
	params := []interface{}{
		loadBalancerUuid,
		serverInstances,
	}
	err = r.Session.DoRequest("SoftLayer_Network_LBaaS_Member", "addLoadBalancerMembers", params, &r.Options, &resp)
	return
}

// Delete given members from load balancer and return load balancer object with listeners, pools and members populated
func (r Network_LBaaS_Member) DeleteLoadBalancerMembers(loadBalancerUuid *string, memberUuids []string) (resp datatypes.Network_LBaaS_LoadBalancer, err error) {
	params := []interface{}{
		loadBalancerUuid,
		memberUuids,
	}
	err = r.Session.DoRequest("SoftLayer_Network_LBaaS_Member", "deleteLoadBalancerMembers", params, &r.Options, &resp)
	return
}

// no documentation yet
type Network_Message_Delivery_Email_Sendgrid struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkMessageDeliveryEmailSendgridService returns an instance of the Network_Message_Delivery_Email_Sendgrid SoftLayer service
func GetNetworkMessageDeliveryEmailSendgridService(sess session.SLSession) Network_Message_Delivery_Email_Sendgrid {
	return Network_Message_Delivery_Email_Sendgrid{Session: sess}
}

func (r Network_Message_Delivery_Email_Sendgrid) Id(id int) Network_Message_Delivery_Email_Sendgrid {
	r.Options.Id = &id
	return r
}

func (r Network_Message_Delivery_Email_Sendgrid) Mask(mask string) Network_Message_Delivery_Email_Sendgrid {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_Message_Delivery_Email_Sendgrid) Filter(filter string) Network_Message_Delivery_Email_Sendgrid {
	r.Options.Filter = filter
	return r
}

func (r Network_Message_Delivery_Email_Sendgrid) Limit(limit int) Network_Message_Delivery_Email_Sendgrid {
	r.Options.Limit = &limit
	return r
}

func (r Network_Message_Delivery_Email_Sendgrid) Offset(offset int) Network_Message_Delivery_Email_Sendgrid {
	r.Options.Offset = &offset
	return r
}

// no documentation yet
func (r Network_Message_Delivery_Email_Sendgrid) EditObject(templateObject *datatypes.Network_Message_Delivery) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Message_Delivery_Email_Sendgrid", "editObject", params, &r.Options, &resp)
	return
}

// no documentation yet
func (r Network_Message_Delivery_Email_Sendgrid) GetAccountOverview() (resp datatypes.Container_Network_Message_Delivery_Email_Sendgrid_Account, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Message_Delivery_Email_Sendgrid", "getAccountOverview", nil, &r.Options, &resp)
	return
}

// no documentation yet
func (r Network_Message_Delivery_Email_Sendgrid) GetObject() (resp datatypes.Network_Message_Delivery_Email_Sendgrid, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Message_Delivery_Email_Sendgrid", "getObject", nil, &r.Options, &resp)
	return
}

// no documentation yet
func (r Network_Message_Delivery_Email_Sendgrid) GetStatistics(options *datatypes.Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Options) (resp []datatypes.Container_Network_Message_Delivery_Email_Sendgrid_Statistics, err error) {
	params := []interface{}{
		options,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Message_Delivery_Email_Sendgrid", "getStatistics", params, &r.Options, &resp)
	return
}

// no documentation yet
func (r Network_Message_Delivery_Email_Sendgrid) UpdateEmailAddress(emailAddress *string) (resp bool, err error) {
	params := []interface{}{
		emailAddress,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Message_Delivery_Email_Sendgrid", "updateEmailAddress", params, &r.Options, &resp)
	return
}

// no documentation yet
type Network_Pod struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkPodService returns an instance of the Network_Pod SoftLayer service
func GetNetworkPodService(sess session.SLSession) Network_Pod {
	return Network_Pod{Session: sess}
}

func (r Network_Pod) Id(id int) Network_Pod {
	r.Options.Id = &id
	return r
}

func (r Network_Pod) Mask(mask string) Network_Pod {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_Pod) Filter(filter string) Network_Pod {
	r.Options.Filter = filter
	return r
}

func (r Network_Pod) Limit(limit int) Network_Pod {
	r.Options.Limit = &limit
	return r
}

func (r Network_Pod) Offset(offset int) Network_Pod {
	r.Options.Offset = &offset
	return r
}

// Filtering is supported for “datacenterName“ and “capabilities“. When filtering on capabilities, use the “in“ operation. Pods fulfilling all capabilities provided will be returned. “datacenterName“ represents an operation against “SoftLayer_Location_Datacenter.name`, such as dal05 when referring to Dallas 5.
//
// ```Examples:```
//
// List Pods in a specific datacenter. <pre> datacenterName.operation = 'dal06' </pre>
//
// List Pods in a geographical area. <pre> datacenterName.operation = '^= dal' </pre>
//
// List Pods in a region fulfilling capabilities. <pre> datacenterName.operation = '^= dal' capabilities.operation = 'in' capabilities.options = [ { name = data, value = [SOME_CAPABILITY, ANOTHER_CAPABILITY] } ] </pre>
func (r Network_Pod) GetAllObjects() (resp []datatypes.Network_Pod, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Pod", "getAllObjects", nil, &r.Options, &resp)
	return
}

// The SoftLayer_Network_SecurityGroup data type contains general information for a single security group. A security group contains a set of IP filter [[SoftLayer_Network_SecurityGroup_Rule (type)|rules]] that define how to handle incoming (ingress) and outgoing (egress) traffic to both the public and private interfaces of a virtual server instance and a set of [[SoftLayer_Virtual_Network_SecurityGroup_NetworkComponentBinding (type)|bindings]] to associate virtual guest network components with the security group.
type Network_SecurityGroup struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkSecurityGroupService returns an instance of the Network_SecurityGroup SoftLayer service
func GetNetworkSecurityGroupService(sess session.SLSession) Network_SecurityGroup {
	return Network_SecurityGroup{Session: sess}
}

func (r Network_SecurityGroup) Id(id int) Network_SecurityGroup {
	r.Options.Id = &id
	return r
}

func (r Network_SecurityGroup) Mask(mask string) Network_SecurityGroup {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_SecurityGroup) Filter(filter string) Network_SecurityGroup {
	r.Options.Filter = filter
	return r
}

func (r Network_SecurityGroup) Limit(limit int) Network_SecurityGroup {
	r.Options.Limit = &limit
	return r
}

func (r Network_SecurityGroup) Offset(offset int) Network_SecurityGroup {
	r.Options.Offset = &offset
	return r
}

// Add new rules to a security group by sending in an array of template [[SoftLayer_Network_SecurityGroup_Rule (type)]] objects to be created.
func (r Network_SecurityGroup) AddRules(ruleTemplates []datatypes.Network_SecurityGroup_Rule) (resp datatypes.Network_SecurityGroup_RequestRules, err error) {
	params := []interface{}{
		ruleTemplates,
	}
	err = r.Session.DoRequest("SoftLayer_Network_SecurityGroup", "addRules", params, &r.Options, &resp)
	return
}

// Attach virtual guest network components to a security group by creating [[SoftLayer_Virtual_Network_SecurityGroup_NetworkComponentBinding (type)]] objects.
func (r Network_SecurityGroup) AttachNetworkComponents(networkComponentIds []int) (resp datatypes.Network_SecurityGroup_Request, err error) {
	params := []interface{}{
		networkComponentIds,
	}
	err = r.Session.DoRequest("SoftLayer_Network_SecurityGroup", "attachNetworkComponents", params, &r.Options, &resp)
	return
}

// Create a new security group.
func (r Network_SecurityGroup) CreateObject(templateObject *datatypes.Network_SecurityGroup) (resp datatypes.Network_SecurityGroup, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Network_SecurityGroup", "createObject", params, &r.Options, &resp)
	return
}

// Delete a security group for an account. A security group cannot be deleted if any network components are attached or if the security group is a remote security group for a [[SoftLayer_Network_SecurityGroup_Rule (type)|rule]].
func (r Network_SecurityGroup) DeleteObject() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_SecurityGroup", "deleteObject", nil, &r.Options, &resp)
	return
}

// Detach virtual guest network components from a security group by deleting its [[SoftLayer_Virtual_Network_SecurityGroup_NetworkComponentBinding (type)]].
func (r Network_SecurityGroup) DetachNetworkComponents(networkComponentIds []int) (resp datatypes.Network_SecurityGroup_Request, err error) {
	params := []interface{}{
		networkComponentIds,
	}
	err = r.Session.DoRequest("SoftLayer_Network_SecurityGroup", "detachNetworkComponents", params, &r.Options, &resp)
	return
}

// Edit a security group.
func (r Network_SecurityGroup) EditObject(templateObject *datatypes.Network_SecurityGroup) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Network_SecurityGroup", "editObject", params, &r.Options, &resp)
	return
}

// Edit rules that belong to the security group. An array of skeleton [[SoftLayer_Network_SecurityGroup_Rule]] objects must be sent in with only the properties defined that you want to change. To edit a property to null, send in -1 for integer properties and "" for string properties. Unchanged properties are left alone.
func (r Network_SecurityGroup) EditRules(ruleTemplates []datatypes.Network_SecurityGroup_Rule) (resp datatypes.Network_SecurityGroup_RequestRules, err error) {
	params := []interface{}{
		ruleTemplates,
	}
	err = r.Session.DoRequest("SoftLayer_Network_SecurityGroup", "editRules", params, &r.Options, &resp)
	return
}

// no documentation yet
func (r Network_SecurityGroup) GetAllObjects() (resp []datatypes.Network_SecurityGroup, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_SecurityGroup", "getAllObjects", nil, &r.Options, &resp)
	return
}

// no documentation yet
func (r Network_SecurityGroup) GetObject() (resp datatypes.Network_SecurityGroup, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_SecurityGroup", "getObject", nil, &r.Options, &resp)
	return
}

// Retrieve The rules for this security group.
func (r Network_SecurityGroup) GetRules() (resp []datatypes.Network_SecurityGroup_Rule, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_SecurityGroup", "getRules", nil, &r.Options, &resp)
	return
}

// Remove rules from a security group.
func (r Network_SecurityGroup) RemoveRules(ruleIds []int) (resp datatypes.Network_SecurityGroup_RequestRules, err error) {
	params := []interface{}{
		ruleIds,
	}
	err = r.Session.DoRequest("SoftLayer_Network_SecurityGroup", "removeRules", params, &r.Options, &resp)
	return
}

// The SoftLayer_Network_Service_Vpn_Overrides data type contains information relating user ids to subnet ids when VPN access is manually configured.  It is essentially an entry in a 'white list' of subnets a SoftLayer portal VPN user may access.
type Network_Service_Vpn_Overrides struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkServiceVpnOverridesService returns an instance of the Network_Service_Vpn_Overrides SoftLayer service
func GetNetworkServiceVpnOverridesService(sess session.SLSession) Network_Service_Vpn_Overrides {
	return Network_Service_Vpn_Overrides{Session: sess}
}

func (r Network_Service_Vpn_Overrides) Id(id int) Network_Service_Vpn_Overrides {
	r.Options.Id = &id
	return r
}

func (r Network_Service_Vpn_Overrides) Mask(mask string) Network_Service_Vpn_Overrides {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_Service_Vpn_Overrides) Filter(filter string) Network_Service_Vpn_Overrides {
	r.Options.Filter = filter
	return r
}

func (r Network_Service_Vpn_Overrides) Limit(limit int) Network_Service_Vpn_Overrides {
	r.Options.Limit = &limit
	return r
}

func (r Network_Service_Vpn_Overrides) Offset(offset int) Network_Service_Vpn_Overrides {
	r.Options.Offset = &offset
	return r
}

// Create Softlayer portal user VPN overrides.
func (r Network_Service_Vpn_Overrides) CreateObjects(templateObjects []datatypes.Network_Service_Vpn_Overrides) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Service_Vpn_Overrides", "createObjects", params, &r.Options, &resp)
	return
}

// Use this method to delete a single SoftLayer portal VPN user subnet override.
func (r Network_Service_Vpn_Overrides) DeleteObject() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Service_Vpn_Overrides", "deleteObject", nil, &r.Options, &resp)
	return
}

// The SoftLayer_Network_Storage data type contains general information regarding a Storage product such as account id, access username and password, the Storage product type, and the server the Storage service is associated with. Currently, only EVault backup storage has an associated server.
type Network_Storage struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkStorageService returns an instance of the Network_Storage SoftLayer service
func GetNetworkStorageService(sess session.SLSession) Network_Storage {
	return Network_Storage{Session: sess}
}

func (r Network_Storage) Id(id int) Network_Storage {
	r.Options.Id = &id
	return r
}

func (r Network_Storage) Mask(mask string) Network_Storage {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_Storage) Filter(filter string) Network_Storage {
	r.Options.Filter = filter
	return r
}

func (r Network_Storage) Limit(limit int) Network_Storage {
	r.Options.Limit = &limit
	return r
}

func (r Network_Storage) Offset(offset int) Network_Storage {
	r.Options.Offset = &offset
	return r
}

// This method is used to modify the access control list for this Storage volume.  The [[SoftLayer_Hardware|SoftLayer_Virtual_Guest|SoftLayer_Network_Subnet|SoftLayer_Network_Subnet_IpAddress]] objects which have been allowed access to this storage volume will be listed in the [[allowedHardware|allowedVirtualGuests|allowedSubnets|allowedIpAddresses]] property of this storage volume.
func (r Network_Storage) AllowAccessFromHostList(hostObjectTemplates []datatypes.Container_Network_Storage_Host) (resp []datatypes.Network_Storage_Allowed_Host, err error) {
	params := []interface{}{
		hostObjectTemplates,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Storage", "allowAccessFromHostList", params, &r.Options, &resp)
	return
}

// no documentation yet
func (r Network_Storage) ConvertCloneDependentToIndependent() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Storage", "convertCloneDependentToIndependent", nil, &r.Options, &resp)
	return
}

// The LUN ID only takes effect during the Host Authorization process. It is required to de-authorize all hosts before using this method.
func (r Network_Storage) CreateOrUpdateLunId(lunId *int) (resp datatypes.Network_Storage_Property, err error) {
	params := []interface{}{
		lunId,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Storage", "createOrUpdateLunId", params, &r.Options, &resp)
	return
}

// no documentation yet
func (r Network_Storage) CreateSnapshot(notes *string) (resp datatypes.Network_Storage, err error) {
	params := []interface{}{
		notes,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Storage", "createSnapshot", params, &r.Options, &resp)
	return
}

// Delete a network storage volume. ”'This cannot be undone.”' At this time only network storage snapshots may be deleted with this method.
//
// ”deleteObject” returns Boolean ”true” on successful deletion or ”false” if it was unable to remove a volume;
func (r Network_Storage) DeleteObject() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Storage", "deleteObject", nil, &r.Options, &resp)
	return
}

// This method is not valid for Legacy iSCSI Storage Volumes.
//
// Disable scheduled snapshots of this storage volume. Scheduling options include 'INTERVAL', HOURLY, DAILY and WEEKLY schedules.
func (r Network_Storage) DisableSnapshots(scheduleType *string) (resp bool, err error) {
	params := []interface{}{
		scheduleType,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Storage", "disableSnapshots", params, &r.Options, &resp)
	return
}

// If a volume (with replication) becomes inaccessible due to a disaster event, this method can be used to immediately failover to an available replica in another location. This method does not allow for fail back via the API. To fail back to the original volume after using this method, open a support ticket. To test failover, use [[SoftLayer_Network_Storage::failoverToReplicant]] instead.
func (r Network_Storage) DisasterRecoveryFailoverToReplicant(replicantId *int) (resp bool, err error) {
	params := []interface{}{
		replicantId,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Storage", "disasterRecoveryFailoverToReplicant", params, &r.Options, &resp)
	return
}

// The password and/or notes may be modified for the Storage service except evault passwords and notes.
func (r Network_Storage) EditObject(templateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Storage", "editObject", params, &r.Options, &resp)
	return
}

// This method is not valid for Legacy iSCSI Storage Volumes.
//
// Enable scheduled snapshots of this storage volume. Scheduling options include HOURLY, DAILY and WEEKLY schedules. For HOURLY schedules, provide relevant data for $scheduleType, $retentionCount and $minute. For DAILY schedules, provide relevant data for $scheduleType, $retentionCount, $minute, and $hour. For WEEKLY schedules, provide relevant data for all parameters of this method.
func (r Network_Storage) EnableSnapshots(scheduleType *string, retentionCount *int, minute *int, hour *int, dayOfWeek *string) (resp bool, err error) {
	params := []interface{}{
		scheduleType,
		retentionCount,
		minute,
		hour,
		dayOfWeek,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Storage", "enableSnapshots", params, &r.Options, &resp)
	return
}

// Failback from a volume replicant. In order to failback the volume must have already been failed over to a replicant.
func (r Network_Storage) FailbackFromReplicant() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Storage", "failbackFromReplicant", nil, &r.Options, &resp)
	return
}

// Failover to a volume replicant.  During the time which the replicant is in use the local nas volume will not be available.
func (r Network_Storage) FailoverToReplicant(replicantId *int) (resp bool, err error) {
	params := []interface{}{
		replicantId,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Storage", "failoverToReplicant", params, &r.Options, &resp)
	return
}

// This method is used to check, if for the given classic file block storage volume, a transaction performing dependent to independent duplicate conversion is active. If yes, then this returns the current percentage of its progress along with its start time as [SoftLayer_Container_Network_Storage_DuplicateConversionStatusInformation] object with its name, percentage and transaction start timestamp.
func (r Network_Storage) GetDuplicateConversionStatus() (resp datatypes.Container_Network_Storage_DuplicateConversionStatusInformation, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Storage", "getDuplicateConversionStatus", nil, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Network_Storage object whose ID corresponds to the ID number of the init parameter passed to the SoftLayer_Network_Storage service.
//
// Please use the associated methods in the [[SoftLayer_Network_Storage]] service to retrieve a Storage account's id.
func (r Network_Storage) GetObject() (resp datatypes.Network_Storage, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Storage", "getObject", nil, &r.Options, &resp)
	return
}

// Retrieve The network storage volumes configured to be replicants of a volume.
func (r Network_Storage) GetReplicationPartners() (resp []datatypes.Network_Storage, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Storage", "getReplicationPartners", nil, &r.Options, &resp)
	return
}

// Retrieve Whether or not a network storage volume may be mounted.
func (r Network_Storage) GetSnapshotNotificationStatus() (resp string, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Storage", "getSnapshotNotificationStatus", nil, &r.Options, &resp)
	return
}

// Retrieve The snapshots associated with this SoftLayer_Network_Storage volume.
func (r Network_Storage) GetSnapshots() (resp []datatypes.Network_Storage, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Storage", "getSnapshots", nil, &r.Options, &resp)
	return
}

// no documentation yet
func (r Network_Storage) GetValidReplicationTargetDatacenterLocations() (resp []datatypes.Location, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Storage", "getValidReplicationTargetDatacenterLocations", nil, &r.Options, &resp)
	return
}

// Retrieves an array of volume count limits per location and globally.
func (r Network_Storage) GetVolumeCountLimits() (resp []datatypes.Container_Network_Storage_DataCenterLimits_VolumeCountLimitContainer, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Storage", "getVolumeCountLimits", nil, &r.Options, &resp)
	return
}

// no documentation yet
func (r Network_Storage) RefreshDuplicate(snapshotId *int, forceRefresh *bool) (resp bool, err error) {
	params := []interface{}{
		snapshotId,
		forceRefresh,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Storage", "refreshDuplicate", params, &r.Options, &resp)
	return
}

// This method is used to modify the access control list for this Storage volume.  The [[SoftLayer_Hardware|SoftLayer_Virtual_Guest|SoftLayer_Network_Subnet|SoftLayer_Network_Subnet_IpAddress]] objects which have been allowed access to this storage will be listed in the [[allowedHardware|allowedVirtualGuests|allowedSubnets|allowedIpAddresses]] property of this storage volume.
func (r Network_Storage) RemoveAccessFromHostList(hostObjectTemplates []datatypes.Container_Network_Storage_Host) (resp []datatypes.Network_Storage_Allowed_Host, err error) {
	params := []interface{}{
		hostObjectTemplates,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Storage", "removeAccessFromHostList", params, &r.Options, &resp)
	return
}

// Restore the volume from a snapshot that was previously taken.
func (r Network_Storage) RestoreFromSnapshot(snapshotId *int) (resp bool, err error) {
	params := []interface{}{
		snapshotId,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Storage", "restoreFromSnapshot", params, &r.Options, &resp)
	return
}

// Function to enable/disable snapshot warning notification.
func (r Network_Storage) SetSnapshotNotification(notificationFlag *bool) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		notificationFlag,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Storage", "setSnapshotNotification", params, &r.Options, &resp)
	return
}

// no documentation yet
type Network_Storage_Allowed_Host struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkStorageAllowedHostService returns an instance of the Network_Storage_Allowed_Host SoftLayer service
func GetNetworkStorageAllowedHostService(sess session.SLSession) Network_Storage_Allowed_Host {
	return Network_Storage_Allowed_Host{Session: sess}
}

func (r Network_Storage_Allowed_Host) Id(id int) Network_Storage_Allowed_Host {
	r.Options.Id = &id
	return r
}

func (r Network_Storage_Allowed_Host) Mask(mask string) Network_Storage_Allowed_Host {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_Storage_Allowed_Host) Filter(filter string) Network_Storage_Allowed_Host {
	r.Options.Filter = filter
	return r
}

func (r Network_Storage_Allowed_Host) Limit(limit int) Network_Storage_Allowed_Host {
	r.Options.Limit = &limit
	return r
}

func (r Network_Storage_Allowed_Host) Offset(offset int) Network_Storage_Allowed_Host {
	r.Options.Offset = &offset
	return r
}

// no documentation yet
func (r Network_Storage_Allowed_Host) AssignSubnetsToAcl(subnetIds []int) (resp []int, err error) {
	params := []interface{}{
		subnetIds,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Storage_Allowed_Host", "assignSubnetsToAcl", params, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Subnet records assigned to the ACL for this allowed host.
func (r Network_Storage_Allowed_Host) GetSubnetsInAcl() (resp []datatypes.Network_Subnet, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Storage_Allowed_Host", "getSubnetsInAcl", nil, &r.Options, &resp)
	return
}

// no documentation yet
func (r Network_Storage_Allowed_Host) RemoveSubnetsFromAcl(subnetIds []int) (resp []int, err error) {
	params := []interface{}{
		subnetIds,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Storage_Allowed_Host", "removeSubnetsFromAcl", params, &r.Options, &resp)
	return
}

// Use this method to modify the credential password for a SoftLayer_Network_Storage_Allowed_Host object.
func (r Network_Storage_Allowed_Host) SetCredentialPassword(password *string) (resp bool, err error) {
	params := []interface{}{
		password,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Storage_Allowed_Host", "setCredentialPassword", params, &r.Options, &resp)
	return
}

// no documentation yet
type Network_Storage_Hub_Cleversafe_Account struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkStorageHubCleversafeAccountService returns an instance of the Network_Storage_Hub_Cleversafe_Account SoftLayer service
func GetNetworkStorageHubCleversafeAccountService(sess session.SLSession) Network_Storage_Hub_Cleversafe_Account {
	return Network_Storage_Hub_Cleversafe_Account{Session: sess}
}

func (r Network_Storage_Hub_Cleversafe_Account) Id(id int) Network_Storage_Hub_Cleversafe_Account {
	r.Options.Id = &id
	return r
}

func (r Network_Storage_Hub_Cleversafe_Account) Mask(mask string) Network_Storage_Hub_Cleversafe_Account {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_Storage_Hub_Cleversafe_Account) Filter(filter string) Network_Storage_Hub_Cleversafe_Account {
	r.Options.Filter = filter
	return r
}

func (r Network_Storage_Hub_Cleversafe_Account) Limit(limit int) Network_Storage_Hub_Cleversafe_Account {
	r.Options.Limit = &limit
	return r
}

func (r Network_Storage_Hub_Cleversafe_Account) Offset(offset int) Network_Storage_Hub_Cleversafe_Account {
	r.Options.Offset = &offset
	return r
}

// Create credentials for an IBM Cloud Object Storage Account
func (r Network_Storage_Hub_Cleversafe_Account) CredentialCreate() (resp []datatypes.Network_Storage_Credential, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Storage_Hub_Cleversafe_Account", "credentialCreate", nil, &r.Options, &resp)
	return
}

// Delete a credential
func (r Network_Storage_Hub_Cleversafe_Account) CredentialDelete(credential *datatypes.Network_Storage_Credential) (resp bool, err error) {
	params := []interface{}{
		credential,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Storage_Hub_Cleversafe_Account", "credentialDelete", params, &r.Options, &resp)
	return
}

// Get buckets
func (r Network_Storage_Hub_Cleversafe_Account) GetBuckets() (resp []datatypes.Container_Network_Storage_Hub_ObjectStorage_Bucket, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Storage_Hub_Cleversafe_Account", "getBuckets", nil, &r.Options, &resp)
	return
}

// Returns credential limits for this IBM Cloud Object Storage account.
func (r Network_Storage_Hub_Cleversafe_Account) GetCredentialLimit() (resp int, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Storage_Hub_Cleversafe_Account", "getCredentialLimit", nil, &r.Options, &resp)
	return
}

// Retrieve Credentials used for generating an AWS signature. Max of 2.
func (r Network_Storage_Hub_Cleversafe_Account) GetCredentials() (resp []datatypes.Network_Storage_Credential, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Storage_Hub_Cleversafe_Account", "getCredentials", nil, &r.Options, &resp)
	return
}

// Returns a collection of endpoint URLs available to this IBM Cloud Object Storage account.
func (r Network_Storage_Hub_Cleversafe_Account) GetEndpoints(accountId *int) (resp []datatypes.Container_Network_Storage_Hub_ObjectStorage_Endpoint, err error) {
	params := []interface{}{
		accountId,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Storage_Hub_Cleversafe_Account", "getEndpoints", params, &r.Options, &resp)
	return
}

// no documentation yet
func (r Network_Storage_Hub_Cleversafe_Account) GetObject() (resp datatypes.Network_Storage_Hub_Cleversafe_Account, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Storage_Hub_Cleversafe_Account", "getObject", nil, &r.Options, &resp)
	return
}

// no documentation yet
type Network_Storage_Iscsi_OS_Type struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkStorageIscsiOSTypeService returns an instance of the Network_Storage_Iscsi_OS_Type SoftLayer service
func GetNetworkStorageIscsiOSTypeService(sess session.SLSession) Network_Storage_Iscsi_OS_Type {
	return Network_Storage_Iscsi_OS_Type{Session: sess}
}

func (r Network_Storage_Iscsi_OS_Type) Id(id int) Network_Storage_Iscsi_OS_Type {
	r.Options.Id = &id
	return r
}

func (r Network_Storage_Iscsi_OS_Type) Mask(mask string) Network_Storage_Iscsi_OS_Type {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_Storage_Iscsi_OS_Type) Filter(filter string) Network_Storage_Iscsi_OS_Type {
	r.Options.Filter = filter
	return r
}

func (r Network_Storage_Iscsi_OS_Type) Limit(limit int) Network_Storage_Iscsi_OS_Type {
	r.Options.Limit = &limit
	return r
}

func (r Network_Storage_Iscsi_OS_Type) Offset(offset int) Network_Storage_Iscsi_OS_Type {
	r.Options.Offset = &offset
	return r
}

// Use this method to retrieve all iSCSI OS Types.
func (r Network_Storage_Iscsi_OS_Type) GetAllObjects() (resp []datatypes.Network_Storage_Iscsi_OS_Type, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Storage_Iscsi_OS_Type", "getAllObjects", nil, &r.Options, &resp)
	return
}

// A subnet represents a continguous range of IP addresses. The range is represented by the networkIdentifer and cidr/netmask properties. The version of a subnet, whether IPv4 or IPv6, is represented by the version property.
//
// When routed, a subnet is associated to a VLAN on your account, which defines its scope on the network. Depending on a subnet's route type, IP addresses may be reserved for network and internal functions, the most common of which is the allocation of network, gateway and broadcast IP addresses.
//
// An unrouted subnet is not active on the network and may generally be routed within the datacenter in which it resides.
//
// [Subnetwork at Wikipedia](http://en.wikipedia.org/wiki/Subnetwork)
//
// [RFC950:Internet Standard Subnetting Procedure](http://datatracker.ietf.org/doc/html/rfc950)
type Network_Subnet struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkSubnetService returns an instance of the Network_Subnet SoftLayer service
func GetNetworkSubnetService(sess session.SLSession) Network_Subnet {
	return Network_Subnet{Session: sess}
}

func (r Network_Subnet) Id(id int) Network_Subnet {
	r.Options.Id = &id
	return r
}

func (r Network_Subnet) Mask(mask string) Network_Subnet {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_Subnet) Filter(filter string) Network_Subnet {
	r.Options.Filter = filter
	return r
}

func (r Network_Subnet) Limit(limit int) Network_Subnet {
	r.Options.Limit = &limit
	return r
}

func (r Network_Subnet) Offset(offset int) Network_Subnet {
	r.Options.Offset = &offset
	return r
}

// This interface allows you to remove the route of your secondary subnets. The result will be a subnet that is no longer routed on the network. Remove the route of subnets you are not actively using, as it will make it easier to identify available subnets later.
//
// ”'Important:”' When removing the route of ”Portable” subnets, know that any subnet depending on an IP address provided by the Portable subnet will also have their routes removed!
//
// To review what subnets are routed to IP addresses provided by a ”Portable” subnet, you can utilize the following object mask: 'mask[ipAddresses[endpointSubnets]]'. Any subnet present in conjunction with ”endpointSubnets” is a subnet which depends on the respective IP address.
//
// The behavior of this interface is such that either true or false is returned. A result of false can be interpreted as the clear route request having already been completed. In contrast, a result of true means the subnet is currently routed and will be transitioned. This route change is asynchronous to the request. A response of true does not mean the subnet's route has changed, but simply that it will change. In order to monitor for the completion of the change, you may either attempt a clear route again until the result is false, or monitor one or more SoftLayer_Network_Subnet properties: subnetType, networkVlanId, and or endPointIpAddress to determine if routing of the subnet has been removed.
func (r Network_Subnet) ClearRoute() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Subnet", "clearRoute", nil, &r.Options, &resp)
	return
}

// Edit the note for this subnet.
func (r Network_Subnet) EditNote(note *string) (resp bool, err error) {
	params := []interface{}{
		note,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Subnet", "editNote", params, &r.Options, &resp)
	return
}

// Retrieves a subnet by its id value. Only subnets assigned to your account are accessible.
func (r Network_Subnet) GetObject() (resp datatypes.Network_Subnet, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Subnet", "getObject", nil, &r.Options, &resp)
	return
}

// This interface allows you to change the route of your secondary subnets. It accommodates a number of ways to identify your desired routing destination through the use of a 'type' and 'identifier'. Subnets may be routed as either Static or Portable, and that designation is dictated by the routing destination specified.
//
// Static subnets have an ultimate routing destination of a single IP address but may not be routed to an existing subnet's IP address whose subnet is routed as a Static. Portable subnets have an ultimate routing destination of a VLAN.
//
// A subnet can be routed to any resource within the same "routing region" as the subnet itself. A subnet's routing region can be diverse but is usually limited to a single data center.
//
// The following identifier 'type' values will result in Static routing: <ul> <li>SoftLayer_Network_Subnet_IpAddress</li> <li>SoftLayer_Hardware_Server</li> <li>SoftLayer_Virtual_Guest</li> </ul>
//
// The following identifier 'type' values will result in Portable routing: <ul> <li>SoftLayer_Network_Vlan</li> </ul>
//
// For each identifier type, one or more 'identifier' formats are possible.
//
// ”SoftLayer_Network_Subnet_IpAddress” will accept the following identifier formats: <ul> <li>An entirely numeric value will be treated as a SoftLayer_Network_Subnet_IpAddress.id value of the desired IP address object.</li> <li>A dotted-quad IPv4 address.</li> <li>A full or compressed IPv6 address.</li> </ul>
//
// ”SoftLayer_Network_Vlan” will accept the following identifier formats: <ul> <li>An entirely numeric value will be treated as a SoftLayer_Network_Vlan.id value of the desired VLAN object.</li> <li>A semantic VLAN identifier of the form &lt;data center short name&gt;.&lt;router&gt;.&lt;vlan number&gt;, where &lt; and &gt; are literal, eg. dal13.fcr01.1234 - the router name may optionally contain the 'a' or 'b' redundancy qualifier (which has no meaning in this context).</li> </ul>
//
// ”SoftLayer_Hardware_Server” will accept the following identifier formats: <ul> <li>An entirely numeric value will be treated as a SoftLayer_Hardware_Server.id value of the desired server.</li> <li>A UUID corresponding to a server's SoftLayer_Hardware_Server.globalIdentifier.</li> <li>A value corresponding to a unique SoftLayer_Hardware_Server.hostname.</li> <li>A value corresponding to a unique fully-qualified domain name in the format 'hostname&lt;domain&gt;' where &lt; and &gt; are literal, e.g. myhost&lt;mydomain.com&gt;, hostname refers to SoftLayer_Hardware_Server.hostname and domain to SoftLayer_Hardware_Server.domain, respectively.</li> </ul>
//
// ”SoftLayer_Virtual_Guest” will accept the following identifier formats: <ul> <li>An entirely numeric value will be treated as a SoftLayer_Virtual_Guest.id value of the desired server.</li> <li>A UUID corresponding to a server's SoftLayer_Virtual_Guest.globalIdentifier.</li> <li>A value corresponding to a unique SoftLayer_Virtual_Guest.hostname.</li> <li>A value corresponding to a unique fully-qualified domain name in the format 'hostname&lt;domain&gt;' where &lt; and &gt; are literal, e.g. myhost&lt;mydomain.com&gt;, hostname refers to SoftLayer_Virtual_Guest.hostname and domain to SoftLayer_Virtual_Guest.domain, respectively.</li> </ul>
//
// The routing destination result of specifying a SoftLayer_Hardware_Server or SoftLayer_Virtual_Guest type will be the primary IP address of the server for the same network segment the subnet is on. Thus, a public subnet will be routed to the server's public, primary IP address. Additionally, this IP address resolution will match the subnet's IP version; routing a IPv6 subnet to a server will result in selection of the primary IPv6 address of the respective network segment, if available.
//
// Subnets may only be routed to the IP version they themselves represent. That means an IPv4 subnet can only be routed to IPv4 addresses. Any type/identifier combination that resolves to an IP address must be able to locate an IP address of the same version as the subnet being routed.
//
// When routing to an IP address on a Primary subnet, only those addresses actively assigned to resources may be targeted. Additionally, the network, gateway, or broadcast address of any Portable subnet may not be a routing destination. For some VLANs utilizing the HSRP redundancy strategy, there are additional addresses which cannot be a route destination.
//
// When routing a subnet that is already routed, note that the subnet first has its route removed; this procedure is the same as what will occur when using SoftLayer_Network_Subnet::clearRoute. Special consideration should be made for subnets routed as Portable. Please refer to the documentation for SoftLayer_Network_Subnet::clearRoute for details.
//
// The behavior of this interface is such that either true or false is returned. A response of false indicates the route request would not result in the route of the subnet changing; attempts to route the subnet to the same destination, even if identified by differing means, will result in no changes. A result of false can be interpreted as the route request having already been completed. In contrast, a result of true means the requested destination is different from the current destination and the subnet's routing will be transitioned. This route change is asynchronous to the request. A response of true does not mean the subnet's route has changed, but simply that it will change. In order to monitor for the completion of the change, you may either attempt a route change again until the result is false, or monitor one or more SoftLayer_Network_Subnet properties: subnetType, networkVlanId, and or endPointIpAddress to determine if routing of the subnet has become the desired route destination.
//
// Use of this operation is limited to a single active request per subnet. If a previous route request is not yet complete, a "not ready" message will be returned upon subsequent requests.
func (r Network_Subnet) Route(typ *string, identifier *string) (resp bool, err error) {
	params := []interface{}{
		typ,
		identifier,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Subnet", "route", params, &r.Options, &resp)
	return
}

// Tag a subnet by passing in one or more tags separated by a comma. Any existing tags you wish to keep should be included in the set of tags, as any missing tags will be removed. To remove all tags from the subnet, send an empty string.
func (r Network_Subnet) SetTags(tags *string) (resp bool, err error) {
	params := []interface{}{
		tags,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Subnet", "setTags", params, &r.Options, &resp)
	return
}

// The SoftLayer_Network_Subnet_IpAddress data type contains general information relating to a single SoftLayer IPv4 address.
type Network_Subnet_IpAddress struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkSubnetIpAddressService returns an instance of the Network_Subnet_IpAddress SoftLayer service
func GetNetworkSubnetIpAddressService(sess session.SLSession) Network_Subnet_IpAddress {
	return Network_Subnet_IpAddress{Session: sess}
}

func (r Network_Subnet_IpAddress) Id(id int) Network_Subnet_IpAddress {
	r.Options.Id = &id
	return r
}

func (r Network_Subnet_IpAddress) Mask(mask string) Network_Subnet_IpAddress {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_Subnet_IpAddress) Filter(filter string) Network_Subnet_IpAddress {
	r.Options.Filter = filter
	return r
}

func (r Network_Subnet_IpAddress) Limit(limit int) Network_Subnet_IpAddress {
	r.Options.Limit = &limit
	return r
}

func (r Network_Subnet_IpAddress) Offset(offset int) Network_Subnet_IpAddress {
	r.Options.Offset = &offset
	return r
}

// Edit a subnet IP address.
func (r Network_Subnet_IpAddress) EditObject(templateObject *datatypes.Network_Subnet_IpAddress) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Subnet_IpAddress", "editObject", params, &r.Options, &resp)
	return
}

// Search for an IP address record by IP address.
func (r Network_Subnet_IpAddress) GetByIpAddress(ipAddress *string) (resp datatypes.Network_Subnet_IpAddress, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Subnet_IpAddress", "getByIpAddress", params, &r.Options, &resp)
	return
}

// no documentation yet
type Network_Subnet_IpAddress_Global struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkSubnetIpAddressGlobalService returns an instance of the Network_Subnet_IpAddress_Global SoftLayer service
func GetNetworkSubnetIpAddressGlobalService(sess session.SLSession) Network_Subnet_IpAddress_Global {
	return Network_Subnet_IpAddress_Global{Session: sess}
}

func (r Network_Subnet_IpAddress_Global) Id(id int) Network_Subnet_IpAddress_Global {
	r.Options.Id = &id
	return r
}

func (r Network_Subnet_IpAddress_Global) Mask(mask string) Network_Subnet_IpAddress_Global {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_Subnet_IpAddress_Global) Filter(filter string) Network_Subnet_IpAddress_Global {
	r.Options.Filter = filter
	return r
}

func (r Network_Subnet_IpAddress_Global) Limit(limit int) Network_Subnet_IpAddress_Global {
	r.Options.Limit = &limit
	return r
}

func (r Network_Subnet_IpAddress_Global) Offset(offset int) Network_Subnet_IpAddress_Global {
	r.Options.Offset = &offset
	return r
}

// no documentation yet
func (r Network_Subnet_IpAddress_Global) GetObject() (resp datatypes.Network_Subnet_IpAddress_Global, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Subnet_IpAddress_Global", "getObject", nil, &r.Options, &resp)
	return
}

// ***DEPRECATED***
// This endpoint is deprecated in favor of the more expressive and capable SoftLayer_Network_Subnet::route, to which this endpoint now proxies. Refer to it for more information.
//
// Similarly, unroute requests are proxied to SoftLayer_Network_Subnet::clearRoute.
// Deprecated: This function has been marked as deprecated.
func (r Network_Subnet_IpAddress_Global) Route(newEndPointIpAddress *string) (resp bool, err error) {
	params := []interface{}{
		newEndPointIpAddress,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Subnet_IpAddress_Global", "route", params, &r.Options, &resp)
	return
}

// ***DEPRECATED***
// This endpoint is deprecated in favor of SoftLayer_Network_Subnet::clearRoute, to which this endpoint now proxies. Refer to it for more information.
// Deprecated: This function has been marked as deprecated.
func (r Network_Subnet_IpAddress_Global) Unroute() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Subnet_IpAddress_Global", "unroute", nil, &r.Options, &resp)
	return
}

// The SoftLayer_Network_Tunnel_Module_Context data type contains general information relating to a single SoftLayer network tunnel.  The SoftLayer_Network_Tunnel_Module_Context is useful to gather information such as related customer subnets (remote) and internal subnets (local) associated with the network tunnel as well as other information needed to manage the network tunnel.  Account and billing information related to the network tunnel can also be retrieved.
type Network_Tunnel_Module_Context struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkTunnelModuleContextService returns an instance of the Network_Tunnel_Module_Context SoftLayer service
func GetNetworkTunnelModuleContextService(sess session.SLSession) Network_Tunnel_Module_Context {
	return Network_Tunnel_Module_Context{Session: sess}
}

func (r Network_Tunnel_Module_Context) Id(id int) Network_Tunnel_Module_Context {
	r.Options.Id = &id
	return r
}

func (r Network_Tunnel_Module_Context) Mask(mask string) Network_Tunnel_Module_Context {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_Tunnel_Module_Context) Filter(filter string) Network_Tunnel_Module_Context {
	r.Options.Filter = filter
	return r
}

func (r Network_Tunnel_Module_Context) Limit(limit int) Network_Tunnel_Module_Context {
	r.Options.Limit = &limit
	return r
}

func (r Network_Tunnel_Module_Context) Offset(offset int) Network_Tunnel_Module_Context {
	r.Options.Offset = &offset
	return r
}

// Associates a remote subnet to the network tunnel.  When a remote subnet is associated, a network tunnel will allow the customer (remote) network to communicate with the private and service subnets on the SoftLayer network which are on the other end of this network tunnel.
//
// NOTE:  A network tunnel's configurations must be applied to the network device in order for the association described above to take effect.
func (r Network_Tunnel_Module_Context) AddCustomerSubnetToNetworkTunnel(subnetId *int) (resp bool, err error) {
	params := []interface{}{
		subnetId,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Tunnel_Module_Context", "addCustomerSubnetToNetworkTunnel", params, &r.Options, &resp)
	return
}

// Associates a private subnet to the network tunnel.  When a private subnet is associated, the network tunnel will allow the customer (remote) network to access the private subnet.
//
// NOTE:  A network tunnel's configurations must be applied to the network device in order for the association described above to take effect.
func (r Network_Tunnel_Module_Context) AddPrivateSubnetToNetworkTunnel(subnetId *int) (resp bool, err error) {
	params := []interface{}{
		subnetId,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Tunnel_Module_Context", "addPrivateSubnetToNetworkTunnel", params, &r.Options, &resp)
	return
}

// Associates a service subnet to the network tunnel.  When a service subnet is associated, a network tunnel will allow the customer (remote) network to communicate with the private and service subnets on the SoftLayer network which are on the other end of this network tunnel.  Service subnets provide access to SoftLayer services such as the customer management portal and the SoftLayer API.
//
// NOTE:  A network tunnel's configurations must be applied to the network device in order for the association described above to take effect.
func (r Network_Tunnel_Module_Context) AddServiceSubnetToNetworkTunnel(subnetId *int) (resp bool, err error) {
	params := []interface{}{
		subnetId,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Tunnel_Module_Context", "addServiceSubnetToNetworkTunnel", params, &r.Options, &resp)
	return
}

// An asynchronous task will be created to apply the IPSec network tunnel's configuration to network devices. During this time, an IPSec network tunnel cannot be modified in anyway. Only one network tunnel configuration task can be created at a time. If a task has already been created and has not completed, a new task cannot be created.
func (r Network_Tunnel_Module_Context) ApplyConfigurationsToDevice() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Tunnel_Module_Context", "applyConfigurationsToDevice", nil, &r.Options, &resp)
	return
}

// Create an address translation for a network tunnel.
//
// To create an address translation, ip addresses from an assigned /30 static route subnet are used.  Address translations deliver packets to a destination ip address that is on a customer (remote) subnet.
//
// NOTE:  A network tunnel's configurations must be applied to the network device in order for an address translation to be created.
func (r Network_Tunnel_Module_Context) CreateAddressTranslation(translation *datatypes.Network_Tunnel_Module_Context_Address_Translation) (resp datatypes.Network_Tunnel_Module_Context_Address_Translation, err error) {
	params := []interface{}{
		translation,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Tunnel_Module_Context", "createAddressTranslation", params, &r.Options, &resp)
	return
}

// Remove an existing address translation from a network tunnel.
//
// Address translations deliver packets to a destination ip address that is on a customer subnet (remote).
//
// NOTE:  A network tunnel's configurations must be applied to the network device in order for an address translation to be deleted.
func (r Network_Tunnel_Module_Context) DeleteAddressTranslation(translationId *int) (resp bool, err error) {
	params := []interface{}{
		translationId,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Tunnel_Module_Context", "deleteAddressTranslation", params, &r.Options, &resp)
	return
}

// Edit name, source (SoftLayer IP) ip address and/or destination (Customer IP) ip address for an existing address translation for a network tunnel.
//
// Address translations deliver packets to a destination ip address that is on a customer (remote) subnet.
//
// NOTE:  A network tunnel's configurations must be applied to the network device in order for an address translation to be created.
func (r Network_Tunnel_Module_Context) EditAddressTranslation(translation *datatypes.Network_Tunnel_Module_Context_Address_Translation) (resp datatypes.Network_Tunnel_Module_Context_Address_Translation, err error) {
	params := []interface{}{
		translation,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Tunnel_Module_Context", "editAddressTranslation", params, &r.Options, &resp)
	return
}

// Negotiation parameters for both phases one and two are editable. Here are the phase one and two parameters that can modified:
//
// *Phase One
// **Authentication
// ***Default value is set to MD5.
// ***Valid Options are: MD5, SHA1, SHA256.
// **Encryption
// ***Default value is set to 3DES.
// ***Valid Options are: DES, 3DES, AES128, AES192, AES256.
// **Diffie-Hellman Group
// ***Default value is set to 2.
// ***Valid Options are: 0 (None), 1, 2, 5.
// **Keylife
// ***Default value is set to 3600.
// ***Limits are:  MIN = 120, MAX = 172800
// **Preshared Key
// *Phase Two
// **Authentication
// ***Default value is set to MD5.
// ***Valid Options are: MD5, SHA1, SHA256.
// **Encryption
// ***Default value is set to 3DES.
// ***Valid Options are: DES, 3DES, AES128, AES192, AES256.
// **Diffie-Hellman Group
// ***Default value is set to 2.
// ***Valid Options are: 0 (None), 1, 2, 5.
// **Keylife
// ***Default value is set to 28800.
// ***Limits are:  MIN = 120, MAX = 172800
// **Perfect Forward Secrecy
// ***Valid Options are: Off = 0, On = 1.
// ***NOTE:  If perfect forward secrecy is turned On (set to 1), then a phase 2 diffie-hellman group is required.
//
// The remote peer address for the network tunnel may also be modified if needed.  Invalid options will not be accepted and will cause an exception to be thrown.  There are properties that provide valid options and limits for each negotiation parameter.  Those properties are as follows:
// * encryptionDefault
// * encryptionOptions
// * authenticationDefault
// * authenticationOptions
// * diffieHellmanGroupDefault
// * diffieHellmanGroupOptions
// * phaseOneKeylifeDefault
// * phaseTwoKeylifeDefault
// * keylifeLimits
//
// Configurations cannot be modified if a network tunnel's requires complex manual setups/configuration modifications by the SoftLayer Network department.  If the former is required, the configurations for the network tunnel will be locked until the manual configurations are complete. A network tunnel's configurations are applied via a transaction.  If a network tunnel configuration change transaction is currently running, the network tunnel's setting cannot be modified until the running transaction completes.
//
// NOTE:  A network tunnel's configurations must be applied to the network device in order for the modifications made to take effect.
func (r Network_Tunnel_Module_Context) EditObject(templateObject *datatypes.Network_Tunnel_Module_Context) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Tunnel_Module_Context", "editObject", params, &r.Options, &resp)
	return
}

// Disassociate a customer subnet (remote) from a network tunnel.  When a remote subnet is disassociated, that subnet will not able to communicate with private and service subnets on the SoftLayer network.
//
// NOTE:  A network tunnel's configurations must be applied to the network device in order for the disassociation described above to take effect.
func (r Network_Tunnel_Module_Context) RemoveCustomerSubnetFromNetworkTunnel(subnetId *int) (resp bool, err error) {
	params := []interface{}{
		subnetId,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Tunnel_Module_Context", "removeCustomerSubnetFromNetworkTunnel", params, &r.Options, &resp)
	return
}

// Disassociate a private subnet from a network tunnel.  When a private subnet is disassociated, the customer (remote) subnet on the other end of the tunnel will not able to communicate with the private subnet that was just disassociated.
//
// NOTE:  A network tunnel's configurations must be applied to the network device in order for the disassociation described above to take effect.
func (r Network_Tunnel_Module_Context) RemovePrivateSubnetFromNetworkTunnel(subnetId *int) (resp bool, err error) {
	params := []interface{}{
		subnetId,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Tunnel_Module_Context", "removePrivateSubnetFromNetworkTunnel", params, &r.Options, &resp)
	return
}

// Disassociate a service subnet from a network tunnel.  When a service subnet is disassociated, that customer (remote) subnet on the other end of the network tunnel will not able to communicate with that service subnet on the SoftLayer network.
//
// NOTE:  A network tunnel's configurations must be applied to the network device in order for the disassociation described above to take effect.
func (r Network_Tunnel_Module_Context) RemoveServiceSubnetFromNetworkTunnel(subnetId *int) (resp bool, err error) {
	params := []interface{}{
		subnetId,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Tunnel_Module_Context", "removeServiceSubnetFromNetworkTunnel", params, &r.Options, &resp)
	return
}

// VLANs comprise the fundamental segmentation model on the network, isolating customer networks from one another.
//
// VLANs are scoped to a single network, generally public or private, and a pod. Through association to a single VLAN, assigned subnets are routed on the network to provide IP address connectivity.
//
// Compute devices are associated to a single VLAN per active network, to which the Primary IP Address and containing Primary Subnet belongs. Additional VLANs may be associated to bare metal devices using VLAN trunking.
//
// [VLAN at Wikipedia](https://en.wikipedia.org/wiki/VLAN)
type Network_Vlan struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkVlanService returns an instance of the Network_Vlan SoftLayer service
func GetNetworkVlanService(sess session.SLSession) Network_Vlan {
	return Network_Vlan{Session: sess}
}

func (r Network_Vlan) Id(id int) Network_Vlan {
	r.Options.Id = &id
	return r
}

func (r Network_Vlan) Mask(mask string) Network_Vlan {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_Vlan) Filter(filter string) Network_Vlan {
	r.Options.Filter = filter
	return r
}

func (r Network_Vlan) Limit(limit int) Network_Vlan {
	r.Options.Limit = &limit
	return r
}

func (r Network_Vlan) Offset(offset int) Network_Vlan {
	r.Options.Offset = &offset
	return r
}

// Updates this VLAN using the provided VLAN template.
//
// The following properties may be modified.
//
// - "name" - A description no more than 20 characters in length.
func (r Network_Vlan) EditObject(templateObject *datatypes.Network_Vlan) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Network_Vlan", "editObject", params, &r.Options, &resp)
	return
}

// Evaluates this VLAN for cancellation and returns a list of descriptions why this VLAN may not be cancelled. If the result is empty, this VLAN may be cancelled.
func (r Network_Vlan) GetCancelFailureReasons() (resp []string, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Vlan", "getCancelFailureReasons", nil, &r.Options, &resp)
	return
}

// Retrieves a VLAN by its id value. Only VLANs assigned to your account are accessible.
func (r Network_Vlan) GetObject() (resp datatypes.Network_Vlan, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Vlan", "getObject", nil, &r.Options, &resp)
	return
}

// The SoftLayer_Network_Vlan_Firewall data type contains general information relating to a single SoftLayer VLAN firewall. This is the object which ties the running rules to a specific downstream server. Use the [[SoftLayer Network Firewall Template]] service to pull SoftLayer recommended rule set templates. Use the [[SoftLayer Network Firewall Update Request]] service to submit a firewall update request.
type Network_Vlan_Firewall struct {
	Session session.SLSession
	Options sl.Options
}

// GetNetworkVlanFirewallService returns an instance of the Network_Vlan_Firewall SoftLayer service
func GetNetworkVlanFirewallService(sess session.SLSession) Network_Vlan_Firewall {
	return Network_Vlan_Firewall{Session: sess}
}

func (r Network_Vlan_Firewall) Id(id int) Network_Vlan_Firewall {
	r.Options.Id = &id
	return r
}

func (r Network_Vlan_Firewall) Mask(mask string) Network_Vlan_Firewall {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Network_Vlan_Firewall) Filter(filter string) Network_Vlan_Firewall {
	r.Options.Filter = filter
	return r
}

func (r Network_Vlan_Firewall) Limit(limit int) Network_Vlan_Firewall {
	r.Options.Limit = &limit
	return r
}

func (r Network_Vlan_Firewall) Offset(offset int) Network_Vlan_Firewall {
	r.Options.Offset = &offset
	return r
}

// getObject returns a SoftLayer_Network_Vlan_Firewall object. You can only get objects for vlans attached to your account that have a network firewall enabled.
func (r Network_Vlan_Firewall) GetObject() (resp datatypes.Network_Vlan_Firewall, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Vlan_Firewall", "getObject", nil, &r.Options, &resp)
	return
}

// Retrieve The currently running rule set of this network component firewall.
func (r Network_Vlan_Firewall) GetRules() (resp []datatypes.Network_Vlan_Firewall_Rule, err error) {
	err = r.Session.DoRequest("SoftLayer_Network_Vlan_Firewall", "getRules", nil, &r.Options, &resp)
	return
}
