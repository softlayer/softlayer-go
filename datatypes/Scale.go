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

package datatypes

import "time"

type Scale_Asset struct {
	Entity

	CreateDate   *time.Time   `json:"createDate,omitempty"`
	DeleteFlag   *bool        `json:"deleteFlag,omitempty"`
	Id           *int         `json:"id,omitempty"`
	ScaleGroup   *Scale_Group `json:"scaleGroup,omitempty"`
	ScaleGroupId *int         `json:"scaleGroupId,omitempty"`
}

type Scale_Asset_Hardware struct {
	Scale_Asset

	Hardware   *Hardware `json:"hardware,omitempty"`
	HardwareId *int      `json:"hardwareId,omitempty"`
}

type Scale_Asset_Virtual_Guest struct {
	Scale_Asset

	VirtualGuest   *Virtual_Guest `json:"virtualGuest,omitempty"`
	VirtualGuestId *int           `json:"virtualGuestId,omitempty"`
}

type Scale_Group struct {
	Entity

	Account                    *Account                     `json:"account,omitempty"`
	AccountId                  *int                         `json:"accountId,omitempty"`
	BalancedTerminationFlag    *bool                        `json:"balancedTerminationFlag,omitempty"`
	Cooldown                   *int                         `json:"cooldown,omitempty"`
	CreateDate                 *time.Time                   `json:"createDate,omitempty"`
	DesiredMemberCount         *int                         `json:"desiredMemberCount,omitempty"`
	Id                         *int                         `json:"id,omitempty"`
	LastActionDate             *time.Time                   `json:"lastActionDate,omitempty"`
	LoadBalancerCount          *uint                        `json:"loadBalancerCount,omitempty"`
	LoadBalancers              []Scale_LoadBalancer         `json:"loadBalancers,omitempty"`
	LogCount                   *uint                        `json:"logCount,omitempty"`
	Logs                       []Scale_Group_Log            `json:"logs,omitempty"`
	MaximumMemberCount         *int                         `json:"maximumMemberCount,omitempty"`
	MinimumMemberCount         *int                         `json:"minimumMemberCount,omitempty"`
	ModifyDate                 *time.Time                   `json:"modifyDate,omitempty"`
	Name                       *string                      `json:"name,omitempty"`
	NetworkVlanCount           *uint                        `json:"networkVlanCount,omitempty"`
	NetworkVlans               []Scale_Network_Vlan         `json:"networkVlans,omitempty"`
	Policies                   []Scale_Policy               `json:"policies,omitempty"`
	PolicyCount                *uint                        `json:"policyCount,omitempty"`
	RegionalGroup              *Location_Group_Regional     `json:"regionalGroup,omitempty"`
	RegionalGroupId            *int                         `json:"regionalGroupId,omitempty"`
	Status                     *Scale_Group_Status          `json:"status,omitempty"`
	SuspendedFlag              *bool                        `json:"suspendedFlag,omitempty"`
	TerminationPolicy          *Scale_Termination_Policy    `json:"terminationPolicy,omitempty"`
	TerminationPolicyId        *int                         `json:"terminationPolicyId,omitempty"`
	VirtualGuestAssetCount     *uint                        `json:"virtualGuestAssetCount,omitempty"`
	VirtualGuestAssets         []Scale_Asset_Virtual_Guest  `json:"virtualGuestAssets,omitempty"`
	VirtualGuestMemberCount    *uint                        `json:"virtualGuestMemberCount,omitempty"`
	VirtualGuestMemberTemplate *Virtual_Guest               `json:"virtualGuestMemberTemplate,omitempty"`
	VirtualGuestMembers        []Scale_Member_Virtual_Guest `json:"virtualGuestMembers,omitempty"`
}

type Scale_Group_Log struct {
	Entity

	CreateDate   *time.Time   `json:"createDate,omitempty"`
	Description  *string      `json:"description,omitempty"`
	Id           *int         `json:"id,omitempty"`
	ScaleGroup   *Scale_Group `json:"scaleGroup,omitempty"`
	ScaleGroupId *int         `json:"scaleGroupId,omitempty"`
}

type Scale_Group_Status struct {
	Entity

	Id      *int    `json:"id,omitempty"`
	KeyName *string `json:"keyName,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type Scale_LoadBalancer struct {
	Entity

	AllocationPercent  *int                                                                 `json:"allocationPercent,omitempty"`
	CreateDate         *time.Time                                                           `json:"createDate,omitempty"`
	DeleteFlag         *bool                                                                `json:"deleteFlag,omitempty"`
	HealthCheck        *Network_Application_Delivery_Controller_LoadBalancer_Health_Check   `json:"healthCheck,omitempty"`
	HealthCheckId      *int                                                                 `json:"healthCheckId,omitempty"`
	Id                 *int                                                                 `json:"id,omitempty"`
	ModifyDate         *time.Time                                                           `json:"modifyDate,omitempty"`
	Port               *int                                                                 `json:"port,omitempty"`
	RoutingMethod      *Network_Application_Delivery_Controller_LoadBalancer_Routing_Method `json:"routingMethod,omitempty"`
	RoutingType        *Network_Application_Delivery_Controller_LoadBalancer_Routing_Type   `json:"routingType,omitempty"`
	ScaleGroup         *Scale_Group                                                         `json:"scaleGroup,omitempty"`
	ScaleGroupId       *int                                                                 `json:"scaleGroupId,omitempty"`
	VirtualIpAddressId *int                                                                 `json:"virtualIpAddressId,omitempty"`
	VirtualServer      *Network_Application_Delivery_Controller_LoadBalancer_VirtualServer  `json:"virtualServer,omitempty"`
	VirtualServerId    *int                                                                 `json:"virtualServerId,omitempty"`
	VirtualServerPort  *int                                                                 `json:"virtualServerPort,omitempty"`
}

type Scale_Member struct {
	Entity

	CreateDate   *time.Time   `json:"createDate,omitempty"`
	Id           *int         `json:"id,omitempty"`
	ScaleGroup   *Scale_Group `json:"scaleGroup,omitempty"`
	ScaleGroupId *int         `json:"scaleGroupId,omitempty"`
}

type Scale_Member_Virtual_Guest struct {
	Scale_Member

	VirtualGuest   *Virtual_Guest `json:"virtualGuest,omitempty"`
	VirtualGuestId *int           `json:"virtualGuestId,omitempty"`
}

type Scale_Network_Vlan struct {
	Entity

	CreateDate    *time.Time    `json:"createDate,omitempty"`
	DeleteFlag    *bool         `json:"deleteFlag,omitempty"`
	Id            *int          `json:"id,omitempty"`
	NetworkVlan   *Network_Vlan `json:"networkVlan,omitempty"`
	NetworkVlanId *int          `json:"networkVlanId,omitempty"`
	ScaleGroup    *Scale_Group  `json:"scaleGroup,omitempty"`
	ScaleGroupId  *int          `json:"scaleGroupId,omitempty"`
}

type Scale_Policy struct {
	Entity

	ActionCount             *uint                              `json:"actionCount,omitempty"`
	Actions                 []Scale_Policy_Action              `json:"actions,omitempty"`
	Cooldown                *int                               `json:"cooldown,omitempty"`
	CreateDate              *time.Time                         `json:"createDate,omitempty"`
	DeleteFlag              *bool                              `json:"deleteFlag,omitempty"`
	Id                      *int                               `json:"id,omitempty"`
	ModifyDate              *time.Time                         `json:"modifyDate,omitempty"`
	Name                    *string                            `json:"name,omitempty"`
	OneTimeTriggerCount     *uint                              `json:"oneTimeTriggerCount,omitempty"`
	OneTimeTriggers         []Scale_Policy_Trigger_OneTime     `json:"oneTimeTriggers,omitempty"`
	RepeatingTriggerCount   *uint                              `json:"repeatingTriggerCount,omitempty"`
	RepeatingTriggers       []Scale_Policy_Trigger_Repeating   `json:"repeatingTriggers,omitempty"`
	ResourceUseTriggerCount *uint                              `json:"resourceUseTriggerCount,omitempty"`
	ResourceUseTriggers     []Scale_Policy_Trigger_ResourceUse `json:"resourceUseTriggers,omitempty"`
	ScaleActionCount        *uint                              `json:"scaleActionCount,omitempty"`
	ScaleActions            []Scale_Policy_Action_Scale        `json:"scaleActions,omitempty"`
	ScaleGroup              *Scale_Group                       `json:"scaleGroup,omitempty"`
	ScaleGroupId            *int                               `json:"scaleGroupId,omitempty"`
	TriggerCount            *uint                              `json:"triggerCount,omitempty"`
	Triggers                []Scale_Policy_Trigger             `json:"triggers,omitempty"`
}

type Scale_Policy_Action struct {
	Entity

	CreateDate    *time.Time                `json:"createDate,omitempty"`
	DeleteFlag    *bool                     `json:"deleteFlag,omitempty"`
	Id            *int                      `json:"id,omitempty"`
	ModifyDate    *time.Time                `json:"modifyDate,omitempty"`
	ScalePolicy   *Scale_Policy             `json:"scalePolicy,omitempty"`
	ScalePolicyId *int                      `json:"scalePolicyId,omitempty"`
	Type          *Scale_Policy_Action_Type `json:"type,omitempty"`
	TypeId        *int                      `json:"typeId,omitempty"`
}

type Scale_Policy_Action_Scale struct {
	Scale_Policy_Action

	Amount    *int    `json:"amount,omitempty"`
	ScaleType *string `json:"scaleType,omitempty"`
}

type Scale_Policy_Action_Type struct {
	Entity

	Id      *int    `json:"id,omitempty"`
	KeyName *string `json:"keyName,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type Scale_Policy_Trigger struct {
	Entity

	CreateDate    *time.Time                 `json:"createDate,omitempty"`
	DeleteFlag    *bool                      `json:"deleteFlag,omitempty"`
	Id            *int                       `json:"id,omitempty"`
	ModifyDate    *time.Time                 `json:"modifyDate,omitempty"`
	ScalePolicy   *Scale_Policy              `json:"scalePolicy,omitempty"`
	ScalePolicyId *int                       `json:"scalePolicyId,omitempty"`
	Type          *Scale_Policy_Trigger_Type `json:"type,omitempty"`
	TypeId        *int                       `json:"typeId,omitempty"`
}

type Scale_Policy_Trigger_OneTime struct {
	Scale_Policy_Trigger

	Date *time.Time `json:"date,omitempty"`
}

type Scale_Policy_Trigger_Repeating struct {
	Scale_Policy_Trigger

	Schedule *string `json:"schedule,omitempty"`
}

type Scale_Policy_Trigger_ResourceUse struct {
	Scale_Policy_Trigger

	WatchCount *uint                                    `json:"watchCount,omitempty"`
	Watches    []Scale_Policy_Trigger_ResourceUse_Watch `json:"watches,omitempty"`
}

type Scale_Policy_Trigger_ResourceUse_Watch struct {
	Entity

	Algorithm            *string                           `json:"algorithm,omitempty"`
	CreateDate           *time.Time                        `json:"createDate,omitempty"`
	DeleteFlag           *bool                             `json:"deleteFlag,omitempty"`
	Id                   *int                              `json:"id,omitempty"`
	Metric               *string                           `json:"metric,omitempty"`
	ModifyDate           *time.Time                        `json:"modifyDate,omitempty"`
	Operator             *string                           `json:"operator,omitempty"`
	Period               *int                              `json:"period,omitempty"`
	ScalePolicyTrigger   *Scale_Policy_Trigger_ResourceUse `json:"scalePolicyTrigger,omitempty"`
	ScalePolicyTriggerId *int                              `json:"scalePolicyTriggerId,omitempty"`
	Value                *string                           `json:"value,omitempty"`
}

type Scale_Policy_Trigger_Type struct {
	Entity

	Id      *int    `json:"id,omitempty"`
	KeyName *string `json:"keyName,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type Scale_Termination_Policy struct {
	Entity

	Id      *int    `json:"id,omitempty"`
	KeyName *string `json:"keyName,omitempty"`
	Name    *string `json:"name,omitempty"`
}
