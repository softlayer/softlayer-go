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

type Scale_Asset struct {
	Session *Session
	Options
}

func (r *Session) GetScaleAssetService() Scale_Asset {
	return Scale_Asset{Session: r}
}

func (r *Scale_Asset) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Asset) GetObject() (resp datatypes.Scale_Asset, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Scale_Asset) GetScaleGroup() (resp datatypes.Scale_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Scale_Asset_Hardware struct {
	Session *Session
	Options
}

func (r *Session) GetScaleAssetHardwareService() Scale_Asset_Hardware {
	return Scale_Asset_Hardware{Session: r}
}

func (r *Scale_Asset_Hardware) CreateObject(templateObject *datatypes.Scale_Asset_Hardware) (resp datatypes.Scale_Asset_Hardware, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Asset_Hardware) GetObject() (resp datatypes.Scale_Asset_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Scale_Asset_Hardware) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Asset_Hardware) GetHardwareId() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Scale_Asset_Virtual_Guest struct {
	Session *Session
	Options
}

func (r *Session) GetScaleAssetVirtualGuestService() Scale_Asset_Virtual_Guest {
	return Scale_Asset_Virtual_Guest{Session: r}
}

func (r *Scale_Asset_Virtual_Guest) CreateObject(templateObject *datatypes.Scale_Asset_Virtual_Guest) (resp datatypes.Scale_Asset_Virtual_Guest, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Asset_Virtual_Guest) GetObject() (resp datatypes.Scale_Asset_Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Scale_Asset_Virtual_Guest) GetVirtualGuest() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Asset_Virtual_Guest) GetVirtualGuestId() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Scale_Group struct {
	Session *Session
	Options
}

func (r *Session) GetScaleGroupService() Scale_Group {
	return Scale_Group{Session: r}
}

func (r *Scale_Group) CreateObject(templateObject *datatypes.Scale_Group) (resp datatypes.Scale_Group, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Group) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Group) EditObject(templateObject *datatypes.Scale_Group) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Group) ForceDeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Group) GetAvailableHourlyInstanceLimit() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Group) GetAvailableRegionalGroups() (resp []datatypes.Location_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Group) GetObject() (resp datatypes.Scale_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Group) Resume() (err error) {
	var resp datatypes.Void
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Group) Scale(delta *int) (resp []datatypes.Scale_Member, err error) {
	params := []interface{}{
		delta,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Group) ScaleTo(number *int) (resp []datatypes.Scale_Member, err error) {
	params := []interface{}{
		number,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Group) Suspend() (err error) {
	var resp datatypes.Void
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Scale_Group) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Group) GetLoadBalancers() (resp []datatypes.Scale_LoadBalancer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Group) GetLogs() (resp []datatypes.Scale_Group_Log, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Group) GetNetworkVlans() (resp []datatypes.Scale_Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Group) GetPolicies() (resp []datatypes.Scale_Policy, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Group) GetRegionalGroup() (resp datatypes.Location_Group_Regional, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Group) GetStatus() (resp datatypes.Scale_Group_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Group) GetTerminationPolicy() (resp datatypes.Scale_Termination_Policy, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Group) GetVirtualGuestAssets() (resp []datatypes.Scale_Asset_Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Group) GetVirtualGuestMembers() (resp []datatypes.Scale_Member_Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Scale_Group_Log struct {
	Session *Session
	Options
}

func (r *Session) GetScaleGroupLogService() Scale_Group_Log {
	return Scale_Group_Log{Session: r}
}

func (r *Scale_Group_Log) GetScaleGroup() (resp datatypes.Scale_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Scale_Group_Status struct {
	Session *Session
	Options
}

func (r *Session) GetScaleGroupStatusService() Scale_Group_Status {
	return Scale_Group_Status{Session: r}
}

func (r *Scale_Group_Status) GetAllObjects() (resp []datatypes.Scale_Group_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Group_Status) GetObject() (resp datatypes.Scale_Group_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Scale_LoadBalancer struct {
	Session *Session
	Options
}

func (r *Session) GetScaleLoadBalancerService() Scale_LoadBalancer {
	return Scale_LoadBalancer{Session: r}
}

func (r *Scale_LoadBalancer) CreateObject(templateObject *datatypes.Scale_LoadBalancer) (resp datatypes.Scale_LoadBalancer, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_LoadBalancer) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_LoadBalancer) EditObject(templateObject *datatypes.Scale_LoadBalancer) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_LoadBalancer) GetObject() (resp datatypes.Scale_LoadBalancer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Scale_LoadBalancer) GetAllocationPercent() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_LoadBalancer) GetHealthCheck() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Check, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_LoadBalancer) GetRoutingMethod() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_Routing_Method, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_LoadBalancer) GetRoutingType() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_Routing_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_LoadBalancer) GetScaleGroup() (resp datatypes.Scale_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_LoadBalancer) GetVirtualIpAddressId() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_LoadBalancer) GetVirtualServer() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualServer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_LoadBalancer) GetVirtualServerPort() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Scale_Member struct {
	Session *Session
	Options
}

func (r *Session) GetScaleMemberService() Scale_Member {
	return Scale_Member{Session: r}
}

func (r *Scale_Member) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Member) GetObject() (resp datatypes.Scale_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Scale_Member) GetScaleGroup() (resp datatypes.Scale_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Scale_Member_Virtual_Guest struct {
	Session *Session
	Options
}

func (r *Session) GetScaleMemberVirtualGuestService() Scale_Member_Virtual_Guest {
	return Scale_Member_Virtual_Guest{Session: r}
}

func (r *Scale_Member_Virtual_Guest) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Member_Virtual_Guest) GetObject() (resp datatypes.Scale_Member_Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Scale_Member_Virtual_Guest) GetVirtualGuest() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Member_Virtual_Guest) GetVirtualGuestId() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Scale_Network_Vlan struct {
	Session *Session
	Options
}

func (r *Session) GetScaleNetworkVlanService() Scale_Network_Vlan {
	return Scale_Network_Vlan{Session: r}
}

func (r *Scale_Network_Vlan) CreateObject(templateObject *datatypes.Scale_Network_Vlan) (resp datatypes.Scale_Network_Vlan, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Network_Vlan) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Network_Vlan) GetObject() (resp datatypes.Scale_Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Scale_Network_Vlan) GetNetworkVlan() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Network_Vlan) GetScaleGroup() (resp datatypes.Scale_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Scale_Policy struct {
	Session *Session
	Options
}

func (r *Session) GetScalePolicyService() Scale_Policy {
	return Scale_Policy{Session: r}
}

func (r *Scale_Policy) CreateObject(templateObject *datatypes.Scale_Policy) (resp datatypes.Scale_Policy, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy) EditObject(templateObject *datatypes.Scale_Policy) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy) GetObject() (resp datatypes.Scale_Policy, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy) Trigger() (resp []datatypes.Scale_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Scale_Policy) GetActions() (resp []datatypes.Scale_Policy_Action, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy) GetOneTimeTriggers() (resp []datatypes.Scale_Policy_Trigger_OneTime, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy) GetRepeatingTriggers() (resp []datatypes.Scale_Policy_Trigger_Repeating, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy) GetResourceUseTriggers() (resp []datatypes.Scale_Policy_Trigger_ResourceUse, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy) GetScaleActions() (resp []datatypes.Scale_Policy_Action_Scale, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy) GetScaleGroup() (resp datatypes.Scale_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy) GetTriggers() (resp []datatypes.Scale_Policy_Trigger, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Scale_Policy_Action struct {
	Session *Session
	Options
}

func (r *Session) GetScalePolicyActionService() Scale_Policy_Action {
	return Scale_Policy_Action{Session: r}
}

func (r *Scale_Policy_Action) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy_Action) EditObject(templateObject *datatypes.Scale_Policy_Action) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy_Action) GetObject() (resp datatypes.Scale_Policy_Action, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Scale_Policy_Action) GetScalePolicy() (resp datatypes.Scale_Policy, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy_Action) GetType() (resp datatypes.Scale_Policy_Action_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Scale_Policy_Action_Scale struct {
	Session *Session
	Options
}

func (r *Session) GetScalePolicyActionScaleService() Scale_Policy_Action_Scale {
	return Scale_Policy_Action_Scale{Session: r}
}

func (r *Scale_Policy_Action_Scale) CreateObject(templateObject *datatypes.Scale_Policy_Action_Scale) (resp datatypes.Scale_Policy_Action_Scale, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy_Action_Scale) GetObject() (resp datatypes.Scale_Policy_Action_Scale, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Scale_Policy_Action_Type struct {
	Session *Session
	Options
}

func (r *Session) GetScalePolicyActionTypeService() Scale_Policy_Action_Type {
	return Scale_Policy_Action_Type{Session: r}
}

func (r *Scale_Policy_Action_Type) GetAllObjects() (resp []datatypes.Scale_Policy_Action_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy_Action_Type) GetObject() (resp datatypes.Scale_Policy_Action_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Scale_Policy_Trigger struct {
	Session *Session
	Options
}

func (r *Session) GetScalePolicyTriggerService() Scale_Policy_Trigger {
	return Scale_Policy_Trigger{Session: r}
}

func (r *Scale_Policy_Trigger) CreateObject(templateObject *datatypes.Scale_Policy_Trigger) (resp datatypes.Scale_Policy_Trigger, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy_Trigger) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy_Trigger) EditObject(templateObject *datatypes.Scale_Policy_Trigger) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy_Trigger) GetObject() (resp datatypes.Scale_Policy_Trigger, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Scale_Policy_Trigger) GetScalePolicy() (resp datatypes.Scale_Policy, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy_Trigger) GetType() (resp datatypes.Scale_Policy_Trigger_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Scale_Policy_Trigger_OneTime struct {
	Session *Session
	Options
}

func (r *Session) GetScalePolicyTriggerOneTimeService() Scale_Policy_Trigger_OneTime {
	return Scale_Policy_Trigger_OneTime{Session: r}
}

func (r *Scale_Policy_Trigger_OneTime) CreateObject(templateObject *datatypes.Scale_Policy_Trigger_OneTime) (resp datatypes.Scale_Policy_Trigger_OneTime, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy_Trigger_OneTime) GetObject() (resp datatypes.Scale_Policy_Trigger_OneTime, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Scale_Policy_Trigger_Repeating struct {
	Session *Session
	Options
}

func (r *Session) GetScalePolicyTriggerRepeatingService() Scale_Policy_Trigger_Repeating {
	return Scale_Policy_Trigger_Repeating{Session: r}
}

func (r *Scale_Policy_Trigger_Repeating) CreateObject(templateObject *datatypes.Scale_Policy_Trigger_Repeating) (resp datatypes.Scale_Policy_Trigger_Repeating, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy_Trigger_Repeating) GetObject() (resp datatypes.Scale_Policy_Trigger_Repeating, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy_Trigger_Repeating) ValidateCronExpression(expression *string) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		expression,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Scale_Policy_Trigger_ResourceUse struct {
	Session *Session
	Options
}

func (r *Session) GetScalePolicyTriggerResourceUseService() Scale_Policy_Trigger_ResourceUse {
	return Scale_Policy_Trigger_ResourceUse{Session: r}
}

func (r *Scale_Policy_Trigger_ResourceUse) CreateObject(templateObject *datatypes.Scale_Policy_Trigger_ResourceUse) (resp datatypes.Scale_Policy_Trigger_ResourceUse, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy_Trigger_ResourceUse) GetObject() (resp datatypes.Scale_Policy_Trigger_ResourceUse, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Scale_Policy_Trigger_ResourceUse) GetWatches() (resp []datatypes.Scale_Policy_Trigger_ResourceUse_Watch, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Scale_Policy_Trigger_ResourceUse_Watch struct {
	Session *Session
	Options
}

func (r *Session) GetScalePolicyTriggerResourceUseWatchService() Scale_Policy_Trigger_ResourceUse_Watch {
	return Scale_Policy_Trigger_ResourceUse_Watch{Session: r}
}

func (r *Scale_Policy_Trigger_ResourceUse_Watch) CreateObject(templateObject *datatypes.Scale_Policy_Trigger_ResourceUse_Watch) (resp datatypes.Scale_Policy_Trigger_ResourceUse_Watch, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy_Trigger_ResourceUse_Watch) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy_Trigger_ResourceUse_Watch) EditObject(templateObject *datatypes.Scale_Policy_Trigger_ResourceUse_Watch) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy_Trigger_ResourceUse_Watch) GetAllPossibleAlgorithms() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy_Trigger_ResourceUse_Watch) GetAllPossibleMetrics() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy_Trigger_ResourceUse_Watch) GetAllPossibleOperators() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy_Trigger_ResourceUse_Watch) GetObject() (resp datatypes.Scale_Policy_Trigger_ResourceUse_Watch, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Scale_Policy_Trigger_ResourceUse_Watch) GetScalePolicyTrigger() (resp datatypes.Scale_Policy_Trigger_ResourceUse, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Scale_Policy_Trigger_Type struct {
	Session *Session
	Options
}

func (r *Session) GetScalePolicyTriggerTypeService() Scale_Policy_Trigger_Type {
	return Scale_Policy_Trigger_Type{Session: r}
}

func (r *Scale_Policy_Trigger_Type) GetAllObjects() (resp []datatypes.Scale_Policy_Trigger_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Policy_Trigger_Type) GetObject() (resp datatypes.Scale_Policy_Trigger_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Scale_Termination_Policy struct {
	Session *Session
	Options
}

func (r *Session) GetScaleTerminationPolicyService() Scale_Termination_Policy {
	return Scale_Termination_Policy{Session: r}
}

func (r *Scale_Termination_Policy) GetAllObjects() (resp []datatypes.Scale_Termination_Policy, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Scale_Termination_Policy) GetObject() (resp datatypes.Scale_Termination_Policy, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
