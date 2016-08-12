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

/**
 * AUTOMATICALLY GENERATED CODE - DO NOT MODIFY
 */

package services

import (
	"github.ibm.com/riethm/gopherlayer/datatypes"
	"github.ibm.com/riethm/gopherlayer/session"
	"github.ibm.com/riethm/gopherlayer/sl"
)

//
type Scale_Asset struct {
	Session *session.Session
	Options sl.Options
}

func GetScaleAssetService(sess *session.Session) *Scale_Asset {
	return &Scale_Asset{Session: sess}
}

func (r Scale_Asset) Id(id int) *Scale_Asset {
	r.Options.Id = &id
	return &r
}

func (r Scale_Asset) Mask(mask string) *Scale_Asset {
	r.Options.Mask = mask
	return &r
}

func (r Scale_Asset) Filter(filter string) *Scale_Asset {
	r.Options.Filter = filter
	return &r
}

func (r Scale_Asset) Limit(limit int) *Scale_Asset {
	r.Options.Limit = &limit
	return &r
}

func (r Scale_Asset) Offset(offset int) *Scale_Asset {
	r.Options.Offset = &offset
	return &r
}

//
func (r *Scale_Asset) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Asset) GetObject() (resp datatypes.Scale_Asset, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The group this asset belongs to.
func (r *Scale_Asset) GetScaleGroup() (resp datatypes.Scale_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Scale_Asset_Hardware struct {
	Session *session.Session
	Options sl.Options
}

func GetScaleAssetHardwareService(sess *session.Session) *Scale_Asset_Hardware {
	return &Scale_Asset_Hardware{Session: sess}
}

func (r Scale_Asset_Hardware) Id(id int) *Scale_Asset_Hardware {
	r.Options.Id = &id
	return &r
}

func (r Scale_Asset_Hardware) Mask(mask string) *Scale_Asset_Hardware {
	r.Options.Mask = mask
	return &r
}

func (r Scale_Asset_Hardware) Filter(filter string) *Scale_Asset_Hardware {
	r.Options.Filter = filter
	return &r
}

func (r Scale_Asset_Hardware) Limit(limit int) *Scale_Asset_Hardware {
	r.Options.Limit = &limit
	return &r
}

func (r Scale_Asset_Hardware) Offset(offset int) *Scale_Asset_Hardware {
	r.Options.Offset = &offset
	return &r
}

//
func (r *Scale_Asset_Hardware) CreateObject(templateObject *datatypes.Scale_Asset_Hardware) (resp datatypes.Scale_Asset_Hardware, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Asset_Hardware) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware for this asset.
func (r *Scale_Asset_Hardware) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The identifier of the hardware for this asset.
func (r *Scale_Asset_Hardware) GetHardwareId() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Asset_Hardware) GetObject() (resp datatypes.Scale_Asset_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The group this asset belongs to.
func (r *Scale_Asset_Hardware) GetScaleGroup() (resp datatypes.Scale_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Scale_Asset_Virtual_Guest struct {
	Session *session.Session
	Options sl.Options
}

func GetScaleAssetVirtualGuestService(sess *session.Session) *Scale_Asset_Virtual_Guest {
	return &Scale_Asset_Virtual_Guest{Session: sess}
}

func (r Scale_Asset_Virtual_Guest) Id(id int) *Scale_Asset_Virtual_Guest {
	r.Options.Id = &id
	return &r
}

func (r Scale_Asset_Virtual_Guest) Mask(mask string) *Scale_Asset_Virtual_Guest {
	r.Options.Mask = mask
	return &r
}

func (r Scale_Asset_Virtual_Guest) Filter(filter string) *Scale_Asset_Virtual_Guest {
	r.Options.Filter = filter
	return &r
}

func (r Scale_Asset_Virtual_Guest) Limit(limit int) *Scale_Asset_Virtual_Guest {
	r.Options.Limit = &limit
	return &r
}

func (r Scale_Asset_Virtual_Guest) Offset(offset int) *Scale_Asset_Virtual_Guest {
	r.Options.Offset = &offset
	return &r
}

//
func (r *Scale_Asset_Virtual_Guest) CreateObject(templateObject *datatypes.Scale_Asset_Virtual_Guest) (resp datatypes.Scale_Asset_Virtual_Guest, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Asset_Virtual_Guest) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Asset_Virtual_Guest) GetObject() (resp datatypes.Scale_Asset_Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The group this asset belongs to.
func (r *Scale_Asset_Virtual_Guest) GetScaleGroup() (resp datatypes.Scale_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The guest for this asset.
func (r *Scale_Asset_Virtual_Guest) GetVirtualGuest() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The identifier of the guest for this asset.
func (r *Scale_Asset_Virtual_Guest) GetVirtualGuestId() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Scale_Group struct {
	Session *session.Session
	Options sl.Options
}

func GetScaleGroupService(sess *session.Session) *Scale_Group {
	return &Scale_Group{Session: sess}
}

func (r Scale_Group) Id(id int) *Scale_Group {
	r.Options.Id = &id
	return &r
}

func (r Scale_Group) Mask(mask string) *Scale_Group {
	r.Options.Mask = mask
	return &r
}

func (r Scale_Group) Filter(filter string) *Scale_Group {
	r.Options.Filter = filter
	return &r
}

func (r Scale_Group) Limit(limit int) *Scale_Group {
	r.Options.Limit = &limit
	return &r
}

func (r Scale_Group) Offset(offset int) *Scale_Group {
	r.Options.Offset = &offset
	return &r
}

//
func (r *Scale_Group) CreateObject(templateObject *datatypes.Scale_Group) (resp datatypes.Scale_Group, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Group) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Group) EditObject(templateObject *datatypes.Scale_Group) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Group) ForceDeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The account for this scaling group.
func (r *Scale_Group) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Group) GetAvailableHourlyInstanceLimit() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Group) GetAvailableRegionalGroups() (resp []datatypes.Location_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Collection of load balancers for this auto scale group.
func (r *Scale_Group) GetLoadBalancers() (resp []datatypes.Scale_LoadBalancer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Collection of log entries for this group.
func (r *Scale_Group) GetLogs() (resp []datatypes.Scale_Group_Log, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Collection of VLANs for this auto scale group. VLANs are optional. This can contain a public or private VLAN or both. When a single VLAN for a public/private type is given it can be a non-purchased VLAN only if the minimumMemberCount on the group is >= 1. This can also contain any number of public/private purchased VLANs and members are staggered across them when scaled up.
func (r *Scale_Group) GetNetworkVlans() (resp []datatypes.Scale_Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Group) GetObject() (resp datatypes.Scale_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Collection of policies for this group. This can be empty.
func (r *Scale_Group) GetPolicies() (resp []datatypes.Scale_Policy, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The regional group for this scale group.
func (r *Scale_Group) GetRegionalGroup() (resp datatypes.Location_Group_Regional, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The status for this scale group.
func (r *Scale_Group) GetStatus() (resp datatypes.Scale_Group_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The termination policy for this scaling group.
func (r *Scale_Group) GetTerminationPolicy() (resp datatypes.Scale_Termination_Policy, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Collection of guests that have been pinned to this group. Guest assets are only used for certain trigger checks such as resource watches. They do not count towards the auto scaling guest counts of this group in anyway and are never automatically added or removed.
func (r *Scale_Group) GetVirtualGuestAssets() (resp []datatypes.Scale_Asset_Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Collection of guests that have been scaled with the group. When this group is active, the count of guests here is guaranteed to be between minimumMemberCount and maximumMemberCount inclusively.
func (r *Scale_Group) GetVirtualGuestMembers() (resp []datatypes.Scale_Member_Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Group) Resume() (err error) {
	var resp datatypes.Void
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Group) Scale(delta *int) (resp []datatypes.Scale_Member, err error) {
	params := []interface{}{
		delta,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Group) ScaleTo(number *int) (resp []datatypes.Scale_Member, err error) {
	params := []interface{}{
		number,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Group) Suspend() (err error) {
	var resp datatypes.Void
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Scale_Group_Status struct {
	Session *session.Session
	Options sl.Options
}

func GetScaleGroupStatusService(sess *session.Session) *Scale_Group_Status {
	return &Scale_Group_Status{Session: sess}
}

func (r Scale_Group_Status) Id(id int) *Scale_Group_Status {
	r.Options.Id = &id
	return &r
}

func (r Scale_Group_Status) Mask(mask string) *Scale_Group_Status {
	r.Options.Mask = mask
	return &r
}

func (r Scale_Group_Status) Filter(filter string) *Scale_Group_Status {
	r.Options.Filter = filter
	return &r
}

func (r Scale_Group_Status) Limit(limit int) *Scale_Group_Status {
	r.Options.Limit = &limit
	return &r
}

func (r Scale_Group_Status) Offset(offset int) *Scale_Group_Status {
	r.Options.Offset = &offset
	return &r
}

//
func (r *Scale_Group_Status) GetAllObjects() (resp []datatypes.Scale_Group_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Group_Status) GetObject() (resp datatypes.Scale_Group_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Scale_LoadBalancer struct {
	Session *session.Session
	Options sl.Options
}

func GetScaleLoadBalancerService(sess *session.Session) *Scale_LoadBalancer {
	return &Scale_LoadBalancer{Session: sess}
}

func (r Scale_LoadBalancer) Id(id int) *Scale_LoadBalancer {
	r.Options.Id = &id
	return &r
}

func (r Scale_LoadBalancer) Mask(mask string) *Scale_LoadBalancer {
	r.Options.Mask = mask
	return &r
}

func (r Scale_LoadBalancer) Filter(filter string) *Scale_LoadBalancer {
	r.Options.Filter = filter
	return &r
}

func (r Scale_LoadBalancer) Limit(limit int) *Scale_LoadBalancer {
	r.Options.Limit = &limit
	return &r
}

func (r Scale_LoadBalancer) Offset(offset int) *Scale_LoadBalancer {
	r.Options.Offset = &offset
	return &r
}

//
func (r *Scale_LoadBalancer) CreateObject(templateObject *datatypes.Scale_LoadBalancer) (resp datatypes.Scale_LoadBalancer, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_LoadBalancer) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_LoadBalancer) EditObject(templateObject *datatypes.Scale_LoadBalancer) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The percentage of connections allocated to this virtual server.
func (r *Scale_LoadBalancer) GetAllocationPercent() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The health check for this configuration.
func (r *Scale_LoadBalancer) GetHealthCheck() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Check, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_LoadBalancer) GetObject() (resp datatypes.Scale_LoadBalancer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The routing method.
func (r *Scale_LoadBalancer) GetRoutingMethod() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_Routing_Method, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The routing type.
func (r *Scale_LoadBalancer) GetRoutingType() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_Routing_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The group this load balancer configuration is for.
func (r *Scale_LoadBalancer) GetScaleGroup() (resp datatypes.Scale_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The ID of the virtual IP address.
func (r *Scale_LoadBalancer) GetVirtualIpAddressId() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The virtual server for this configuration.
func (r *Scale_LoadBalancer) GetVirtualServer() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualServer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The port on the virtual server.
func (r *Scale_LoadBalancer) GetVirtualServerPort() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Scale_Member struct {
	Session *session.Session
	Options sl.Options
}

func GetScaleMemberService(sess *session.Session) *Scale_Member {
	return &Scale_Member{Session: sess}
}

func (r Scale_Member) Id(id int) *Scale_Member {
	r.Options.Id = &id
	return &r
}

func (r Scale_Member) Mask(mask string) *Scale_Member {
	r.Options.Mask = mask
	return &r
}

func (r Scale_Member) Filter(filter string) *Scale_Member {
	r.Options.Filter = filter
	return &r
}

func (r Scale_Member) Limit(limit int) *Scale_Member {
	r.Options.Limit = &limit
	return &r
}

func (r Scale_Member) Offset(offset int) *Scale_Member {
	r.Options.Offset = &offset
	return &r
}

//
func (r *Scale_Member) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Member) GetObject() (resp datatypes.Scale_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The group this member belongs to.
func (r *Scale_Member) GetScaleGroup() (resp datatypes.Scale_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Scale_Member_Virtual_Guest struct {
	Session *session.Session
	Options sl.Options
}

func GetScaleMemberVirtualGuestService(sess *session.Session) *Scale_Member_Virtual_Guest {
	return &Scale_Member_Virtual_Guest{Session: sess}
}

func (r Scale_Member_Virtual_Guest) Id(id int) *Scale_Member_Virtual_Guest {
	r.Options.Id = &id
	return &r
}

func (r Scale_Member_Virtual_Guest) Mask(mask string) *Scale_Member_Virtual_Guest {
	r.Options.Mask = mask
	return &r
}

func (r Scale_Member_Virtual_Guest) Filter(filter string) *Scale_Member_Virtual_Guest {
	r.Options.Filter = filter
	return &r
}

func (r Scale_Member_Virtual_Guest) Limit(limit int) *Scale_Member_Virtual_Guest {
	r.Options.Limit = &limit
	return &r
}

func (r Scale_Member_Virtual_Guest) Offset(offset int) *Scale_Member_Virtual_Guest {
	r.Options.Offset = &offset
	return &r
}

//
func (r *Scale_Member_Virtual_Guest) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Member_Virtual_Guest) GetObject() (resp datatypes.Scale_Member_Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The group this member belongs to.
func (r *Scale_Member_Virtual_Guest) GetScaleGroup() (resp datatypes.Scale_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The guest for this member.
func (r *Scale_Member_Virtual_Guest) GetVirtualGuest() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The identifier of the guest for this member.
func (r *Scale_Member_Virtual_Guest) GetVirtualGuestId() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Scale_Network_Vlan struct {
	Session *session.Session
	Options sl.Options
}

func GetScaleNetworkVlanService(sess *session.Session) *Scale_Network_Vlan {
	return &Scale_Network_Vlan{Session: sess}
}

func (r Scale_Network_Vlan) Id(id int) *Scale_Network_Vlan {
	r.Options.Id = &id
	return &r
}

func (r Scale_Network_Vlan) Mask(mask string) *Scale_Network_Vlan {
	r.Options.Mask = mask
	return &r
}

func (r Scale_Network_Vlan) Filter(filter string) *Scale_Network_Vlan {
	r.Options.Filter = filter
	return &r
}

func (r Scale_Network_Vlan) Limit(limit int) *Scale_Network_Vlan {
	r.Options.Limit = &limit
	return &r
}

func (r Scale_Network_Vlan) Offset(offset int) *Scale_Network_Vlan {
	r.Options.Offset = &offset
	return &r
}

//
func (r *Scale_Network_Vlan) CreateObject(templateObject *datatypes.Scale_Network_Vlan) (resp datatypes.Scale_Network_Vlan, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Network_Vlan) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The network VLAN to scale with.
func (r *Scale_Network_Vlan) GetNetworkVlan() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Network_Vlan) GetObject() (resp datatypes.Scale_Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The group this network VLAN is for.
func (r *Scale_Network_Vlan) GetScaleGroup() (resp datatypes.Scale_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Scale_Policy struct {
	Session *session.Session
	Options sl.Options
}

func GetScalePolicyService(sess *session.Session) *Scale_Policy {
	return &Scale_Policy{Session: sess}
}

func (r Scale_Policy) Id(id int) *Scale_Policy {
	r.Options.Id = &id
	return &r
}

func (r Scale_Policy) Mask(mask string) *Scale_Policy {
	r.Options.Mask = mask
	return &r
}

func (r Scale_Policy) Filter(filter string) *Scale_Policy {
	r.Options.Filter = filter
	return &r
}

func (r Scale_Policy) Limit(limit int) *Scale_Policy {
	r.Options.Limit = &limit
	return &r
}

func (r Scale_Policy) Offset(offset int) *Scale_Policy {
	r.Options.Offset = &offset
	return &r
}

//
func (r *Scale_Policy) CreateObject(templateObject *datatypes.Scale_Policy) (resp datatypes.Scale_Policy, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy) EditObject(templateObject *datatypes.Scale_Policy) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The actions to perform upon any trigger hit. Currently this must be a single value.
func (r *Scale_Policy) GetActions() (resp []datatypes.Scale_Policy_Action, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy) GetObject() (resp datatypes.Scale_Policy, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The one-time triggers to check for this group.
func (r *Scale_Policy) GetOneTimeTriggers() (resp []datatypes.Scale_Policy_Trigger_OneTime, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The repeating triggers to check for this group.
func (r *Scale_Policy) GetRepeatingTriggers() (resp []datatypes.Scale_Policy_Trigger_Repeating, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The resource-use triggers to check for this group.
func (r *Scale_Policy) GetResourceUseTriggers() (resp []datatypes.Scale_Policy_Trigger_ResourceUse, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The scale actions to perform upon any trigger hit. Currently this must be a single value.
func (r *Scale_Policy) GetScaleActions() (resp []datatypes.Scale_Policy_Action_Scale, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The group this policy is on.
func (r *Scale_Policy) GetScaleGroup() (resp datatypes.Scale_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The triggers to check for this group.
func (r *Scale_Policy) GetTriggers() (resp []datatypes.Scale_Policy_Trigger, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy) Trigger() (resp []datatypes.Scale_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Scale_Policy_Action struct {
	Session *session.Session
	Options sl.Options
}

func GetScalePolicyActionService(sess *session.Session) *Scale_Policy_Action {
	return &Scale_Policy_Action{Session: sess}
}

func (r Scale_Policy_Action) Id(id int) *Scale_Policy_Action {
	r.Options.Id = &id
	return &r
}

func (r Scale_Policy_Action) Mask(mask string) *Scale_Policy_Action {
	r.Options.Mask = mask
	return &r
}

func (r Scale_Policy_Action) Filter(filter string) *Scale_Policy_Action {
	r.Options.Filter = filter
	return &r
}

func (r Scale_Policy_Action) Limit(limit int) *Scale_Policy_Action {
	r.Options.Limit = &limit
	return &r
}

func (r Scale_Policy_Action) Offset(offset int) *Scale_Policy_Action {
	r.Options.Offset = &offset
	return &r
}

//
func (r *Scale_Policy_Action) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy_Action) EditObject(templateObject *datatypes.Scale_Policy_Action) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy_Action) GetObject() (resp datatypes.Scale_Policy_Action, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The policy this action is on.
func (r *Scale_Policy_Action) GetScalePolicy() (resp datatypes.Scale_Policy, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The type of action.
func (r *Scale_Policy_Action) GetType() (resp datatypes.Scale_Policy_Action_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Scale_Policy_Action_Scale struct {
	Session *session.Session
	Options sl.Options
}

func GetScalePolicyActionScaleService(sess *session.Session) *Scale_Policy_Action_Scale {
	return &Scale_Policy_Action_Scale{Session: sess}
}

func (r Scale_Policy_Action_Scale) Id(id int) *Scale_Policy_Action_Scale {
	r.Options.Id = &id
	return &r
}

func (r Scale_Policy_Action_Scale) Mask(mask string) *Scale_Policy_Action_Scale {
	r.Options.Mask = mask
	return &r
}

func (r Scale_Policy_Action_Scale) Filter(filter string) *Scale_Policy_Action_Scale {
	r.Options.Filter = filter
	return &r
}

func (r Scale_Policy_Action_Scale) Limit(limit int) *Scale_Policy_Action_Scale {
	r.Options.Limit = &limit
	return &r
}

func (r Scale_Policy_Action_Scale) Offset(offset int) *Scale_Policy_Action_Scale {
	r.Options.Offset = &offset
	return &r
}

//
func (r *Scale_Policy_Action_Scale) CreateObject(templateObject *datatypes.Scale_Policy_Action_Scale) (resp datatypes.Scale_Policy_Action_Scale, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy_Action_Scale) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy_Action_Scale) EditObject(templateObject *datatypes.Scale_Policy_Action) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy_Action_Scale) GetObject() (resp datatypes.Scale_Policy_Action_Scale, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The policy this action is on.
func (r *Scale_Policy_Action_Scale) GetScalePolicy() (resp datatypes.Scale_Policy, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The type of action.
func (r *Scale_Policy_Action_Scale) GetType() (resp datatypes.Scale_Policy_Action_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Scale_Policy_Action_Type struct {
	Session *session.Session
	Options sl.Options
}

func GetScalePolicyActionTypeService(sess *session.Session) *Scale_Policy_Action_Type {
	return &Scale_Policy_Action_Type{Session: sess}
}

func (r Scale_Policy_Action_Type) Id(id int) *Scale_Policy_Action_Type {
	r.Options.Id = &id
	return &r
}

func (r Scale_Policy_Action_Type) Mask(mask string) *Scale_Policy_Action_Type {
	r.Options.Mask = mask
	return &r
}

func (r Scale_Policy_Action_Type) Filter(filter string) *Scale_Policy_Action_Type {
	r.Options.Filter = filter
	return &r
}

func (r Scale_Policy_Action_Type) Limit(limit int) *Scale_Policy_Action_Type {
	r.Options.Limit = &limit
	return &r
}

func (r Scale_Policy_Action_Type) Offset(offset int) *Scale_Policy_Action_Type {
	r.Options.Offset = &offset
	return &r
}

//
func (r *Scale_Policy_Action_Type) GetAllObjects() (resp []datatypes.Scale_Policy_Action_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy_Action_Type) GetObject() (resp datatypes.Scale_Policy_Action_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Scale_Policy_Trigger struct {
	Session *session.Session
	Options sl.Options
}

func GetScalePolicyTriggerService(sess *session.Session) *Scale_Policy_Trigger {
	return &Scale_Policy_Trigger{Session: sess}
}

func (r Scale_Policy_Trigger) Id(id int) *Scale_Policy_Trigger {
	r.Options.Id = &id
	return &r
}

func (r Scale_Policy_Trigger) Mask(mask string) *Scale_Policy_Trigger {
	r.Options.Mask = mask
	return &r
}

func (r Scale_Policy_Trigger) Filter(filter string) *Scale_Policy_Trigger {
	r.Options.Filter = filter
	return &r
}

func (r Scale_Policy_Trigger) Limit(limit int) *Scale_Policy_Trigger {
	r.Options.Limit = &limit
	return &r
}

func (r Scale_Policy_Trigger) Offset(offset int) *Scale_Policy_Trigger {
	r.Options.Offset = &offset
	return &r
}

//
func (r *Scale_Policy_Trigger) CreateObject(templateObject *datatypes.Scale_Policy_Trigger) (resp datatypes.Scale_Policy_Trigger, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy_Trigger) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy_Trigger) EditObject(templateObject *datatypes.Scale_Policy_Trigger) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy_Trigger) GetObject() (resp datatypes.Scale_Policy_Trigger, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The policy this trigger is on.
func (r *Scale_Policy_Trigger) GetScalePolicy() (resp datatypes.Scale_Policy, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The type of trigger.
func (r *Scale_Policy_Trigger) GetType() (resp datatypes.Scale_Policy_Trigger_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Scale_Policy_Trigger_OneTime struct {
	Session *session.Session
	Options sl.Options
}

func GetScalePolicyTriggerOneTimeService(sess *session.Session) *Scale_Policy_Trigger_OneTime {
	return &Scale_Policy_Trigger_OneTime{Session: sess}
}

func (r Scale_Policy_Trigger_OneTime) Id(id int) *Scale_Policy_Trigger_OneTime {
	r.Options.Id = &id
	return &r
}

func (r Scale_Policy_Trigger_OneTime) Mask(mask string) *Scale_Policy_Trigger_OneTime {
	r.Options.Mask = mask
	return &r
}

func (r Scale_Policy_Trigger_OneTime) Filter(filter string) *Scale_Policy_Trigger_OneTime {
	r.Options.Filter = filter
	return &r
}

func (r Scale_Policy_Trigger_OneTime) Limit(limit int) *Scale_Policy_Trigger_OneTime {
	r.Options.Limit = &limit
	return &r
}

func (r Scale_Policy_Trigger_OneTime) Offset(offset int) *Scale_Policy_Trigger_OneTime {
	r.Options.Offset = &offset
	return &r
}

//
func (r *Scale_Policy_Trigger_OneTime) CreateObject(templateObject *datatypes.Scale_Policy_Trigger_OneTime) (resp datatypes.Scale_Policy_Trigger_OneTime, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy_Trigger_OneTime) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy_Trigger_OneTime) EditObject(templateObject *datatypes.Scale_Policy_Trigger) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy_Trigger_OneTime) GetObject() (resp datatypes.Scale_Policy_Trigger_OneTime, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The policy this trigger is on.
func (r *Scale_Policy_Trigger_OneTime) GetScalePolicy() (resp datatypes.Scale_Policy, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The type of trigger.
func (r *Scale_Policy_Trigger_OneTime) GetType() (resp datatypes.Scale_Policy_Trigger_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Scale_Policy_Trigger_Repeating struct {
	Session *session.Session
	Options sl.Options
}

func GetScalePolicyTriggerRepeatingService(sess *session.Session) *Scale_Policy_Trigger_Repeating {
	return &Scale_Policy_Trigger_Repeating{Session: sess}
}

func (r Scale_Policy_Trigger_Repeating) Id(id int) *Scale_Policy_Trigger_Repeating {
	r.Options.Id = &id
	return &r
}

func (r Scale_Policy_Trigger_Repeating) Mask(mask string) *Scale_Policy_Trigger_Repeating {
	r.Options.Mask = mask
	return &r
}

func (r Scale_Policy_Trigger_Repeating) Filter(filter string) *Scale_Policy_Trigger_Repeating {
	r.Options.Filter = filter
	return &r
}

func (r Scale_Policy_Trigger_Repeating) Limit(limit int) *Scale_Policy_Trigger_Repeating {
	r.Options.Limit = &limit
	return &r
}

func (r Scale_Policy_Trigger_Repeating) Offset(offset int) *Scale_Policy_Trigger_Repeating {
	r.Options.Offset = &offset
	return &r
}

//
func (r *Scale_Policy_Trigger_Repeating) CreateObject(templateObject *datatypes.Scale_Policy_Trigger_Repeating) (resp datatypes.Scale_Policy_Trigger_Repeating, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy_Trigger_Repeating) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy_Trigger_Repeating) EditObject(templateObject *datatypes.Scale_Policy_Trigger) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy_Trigger_Repeating) GetObject() (resp datatypes.Scale_Policy_Trigger_Repeating, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The policy this trigger is on.
func (r *Scale_Policy_Trigger_Repeating) GetScalePolicy() (resp datatypes.Scale_Policy, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The type of trigger.
func (r *Scale_Policy_Trigger_Repeating) GetType() (resp datatypes.Scale_Policy_Trigger_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy_Trigger_Repeating) ValidateCronExpression(expression *string) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		expression,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
type Scale_Policy_Trigger_ResourceUse struct {
	Session *session.Session
	Options sl.Options
}

func GetScalePolicyTriggerResourceUseService(sess *session.Session) *Scale_Policy_Trigger_ResourceUse {
	return &Scale_Policy_Trigger_ResourceUse{Session: sess}
}

func (r Scale_Policy_Trigger_ResourceUse) Id(id int) *Scale_Policy_Trigger_ResourceUse {
	r.Options.Id = &id
	return &r
}

func (r Scale_Policy_Trigger_ResourceUse) Mask(mask string) *Scale_Policy_Trigger_ResourceUse {
	r.Options.Mask = mask
	return &r
}

func (r Scale_Policy_Trigger_ResourceUse) Filter(filter string) *Scale_Policy_Trigger_ResourceUse {
	r.Options.Filter = filter
	return &r
}

func (r Scale_Policy_Trigger_ResourceUse) Limit(limit int) *Scale_Policy_Trigger_ResourceUse {
	r.Options.Limit = &limit
	return &r
}

func (r Scale_Policy_Trigger_ResourceUse) Offset(offset int) *Scale_Policy_Trigger_ResourceUse {
	r.Options.Offset = &offset
	return &r
}

//
func (r *Scale_Policy_Trigger_ResourceUse) CreateObject(templateObject *datatypes.Scale_Policy_Trigger_ResourceUse) (resp datatypes.Scale_Policy_Trigger_ResourceUse, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy_Trigger_ResourceUse) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy_Trigger_ResourceUse) EditObject(templateObject *datatypes.Scale_Policy_Trigger) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy_Trigger_ResourceUse) GetObject() (resp datatypes.Scale_Policy_Trigger_ResourceUse, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The policy this trigger is on.
func (r *Scale_Policy_Trigger_ResourceUse) GetScalePolicy() (resp datatypes.Scale_Policy, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The type of trigger.
func (r *Scale_Policy_Trigger_ResourceUse) GetType() (resp datatypes.Scale_Policy_Trigger_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The resource watches for this trigger.
func (r *Scale_Policy_Trigger_ResourceUse) GetWatches() (resp []datatypes.Scale_Policy_Trigger_ResourceUse_Watch, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Scale_Policy_Trigger_ResourceUse_Watch struct {
	Session *session.Session
	Options sl.Options
}

func GetScalePolicyTriggerResourceUseWatchService(sess *session.Session) *Scale_Policy_Trigger_ResourceUse_Watch {
	return &Scale_Policy_Trigger_ResourceUse_Watch{Session: sess}
}

func (r Scale_Policy_Trigger_ResourceUse_Watch) Id(id int) *Scale_Policy_Trigger_ResourceUse_Watch {
	r.Options.Id = &id
	return &r
}

func (r Scale_Policy_Trigger_ResourceUse_Watch) Mask(mask string) *Scale_Policy_Trigger_ResourceUse_Watch {
	r.Options.Mask = mask
	return &r
}

func (r Scale_Policy_Trigger_ResourceUse_Watch) Filter(filter string) *Scale_Policy_Trigger_ResourceUse_Watch {
	r.Options.Filter = filter
	return &r
}

func (r Scale_Policy_Trigger_ResourceUse_Watch) Limit(limit int) *Scale_Policy_Trigger_ResourceUse_Watch {
	r.Options.Limit = &limit
	return &r
}

func (r Scale_Policy_Trigger_ResourceUse_Watch) Offset(offset int) *Scale_Policy_Trigger_ResourceUse_Watch {
	r.Options.Offset = &offset
	return &r
}

//
func (r *Scale_Policy_Trigger_ResourceUse_Watch) CreateObject(templateObject *datatypes.Scale_Policy_Trigger_ResourceUse_Watch) (resp datatypes.Scale_Policy_Trigger_ResourceUse_Watch, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy_Trigger_ResourceUse_Watch) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy_Trigger_ResourceUse_Watch) EditObject(templateObject *datatypes.Scale_Policy_Trigger_ResourceUse_Watch) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy_Trigger_ResourceUse_Watch) GetAllPossibleAlgorithms() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy_Trigger_ResourceUse_Watch) GetAllPossibleMetrics() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy_Trigger_ResourceUse_Watch) GetAllPossibleOperators() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy_Trigger_ResourceUse_Watch) GetObject() (resp datatypes.Scale_Policy_Trigger_ResourceUse_Watch, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The trigger this watch is on.
func (r *Scale_Policy_Trigger_ResourceUse_Watch) GetScalePolicyTrigger() (resp datatypes.Scale_Policy_Trigger_ResourceUse, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Scale_Policy_Trigger_Type struct {
	Session *session.Session
	Options sl.Options
}

func GetScalePolicyTriggerTypeService(sess *session.Session) *Scale_Policy_Trigger_Type {
	return &Scale_Policy_Trigger_Type{Session: sess}
}

func (r Scale_Policy_Trigger_Type) Id(id int) *Scale_Policy_Trigger_Type {
	r.Options.Id = &id
	return &r
}

func (r Scale_Policy_Trigger_Type) Mask(mask string) *Scale_Policy_Trigger_Type {
	r.Options.Mask = mask
	return &r
}

func (r Scale_Policy_Trigger_Type) Filter(filter string) *Scale_Policy_Trigger_Type {
	r.Options.Filter = filter
	return &r
}

func (r Scale_Policy_Trigger_Type) Limit(limit int) *Scale_Policy_Trigger_Type {
	r.Options.Limit = &limit
	return &r
}

func (r Scale_Policy_Trigger_Type) Offset(offset int) *Scale_Policy_Trigger_Type {
	r.Options.Offset = &offset
	return &r
}

//
func (r *Scale_Policy_Trigger_Type) GetAllObjects() (resp []datatypes.Scale_Policy_Trigger_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Policy_Trigger_Type) GetObject() (resp datatypes.Scale_Policy_Trigger_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Scale_Termination_Policy struct {
	Session *session.Session
	Options sl.Options
}

func GetScaleTerminationPolicyService(sess *session.Session) *Scale_Termination_Policy {
	return &Scale_Termination_Policy{Session: sess}
}

func (r Scale_Termination_Policy) Id(id int) *Scale_Termination_Policy {
	r.Options.Id = &id
	return &r
}

func (r Scale_Termination_Policy) Mask(mask string) *Scale_Termination_Policy {
	r.Options.Mask = mask
	return &r
}

func (r Scale_Termination_Policy) Filter(filter string) *Scale_Termination_Policy {
	r.Options.Filter = filter
	return &r
}

func (r Scale_Termination_Policy) Limit(limit int) *Scale_Termination_Policy {
	r.Options.Limit = &limit
	return &r
}

func (r Scale_Termination_Policy) Offset(offset int) *Scale_Termination_Policy {
	r.Options.Offset = &offset
	return &r
}

//
func (r *Scale_Termination_Policy) GetAllObjects() (resp []datatypes.Scale_Termination_Policy, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Scale_Termination_Policy) GetObject() (resp datatypes.Scale_Termination_Policy, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
