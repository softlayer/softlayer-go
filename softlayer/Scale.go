package softlayer

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Scale_Asset interface {
	Entity

	DeleteObject() (bool, error)
	GetObject() (datatypes.Scale_Asset, error)

	GetScaleGroup() (datatypes.Scale_Group, error)
}

type Scale_Asset_Hardware interface {
	Scale_Asset

	CreateObject(templateObject datatypes.Scale_Asset_Hardware) (datatypes.Scale_Asset_Hardware, error)
	GetObject() (datatypes.Scale_Asset_Hardware, error)

	GetHardware() (datatypes.Hardware, error)
	GetHardwareId() (int, error)
}

type Scale_Asset_Virtual_Guest interface {
	Scale_Asset

	CreateObject(templateObject datatypes.Scale_Asset_Virtual_Guest) (datatypes.Scale_Asset_Virtual_Guest, error)
	GetObject() (datatypes.Scale_Asset_Virtual_Guest, error)

	GetVirtualGuest() (datatypes.Virtual_Guest, error)
	GetVirtualGuestId() (int, error)
}

type Scale_Group interface {
	Entity

	CreateObject(templateObject datatypes.Scale_Group) (datatypes.Scale_Group, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Scale_Group) (bool, error)
	ForceDeleteObject() (bool, error)
	GetAvailableHourlyInstanceLimit() (int, error)
	GetAvailableRegionalGroups() (datatypes.Location_Group, error)
	GetObject() (datatypes.Scale_Group, error)
	Resume() error
	Scale(delta int) (datatypes.Scale_Member, error)
	ScaleTo(number int) (datatypes.Scale_Member, error)
	Suspend() error

	GetAccount() (datatypes.Account, error)
	GetLoadBalancers() (datatypes.Scale_LoadBalancer, error)
	GetLogs() (datatypes.Scale_Group_Log, error)
	GetNetworkVlans() (datatypes.Scale_Network_Vlan, error)
	GetPolicies() (datatypes.Scale_Policy, error)
	GetRegionalGroup() (datatypes.Location_Group_Regional, error)
	GetStatus() (datatypes.Scale_Group_Status, error)
	GetTerminationPolicy() (datatypes.Scale_Termination_Policy, error)
	GetVirtualGuestAssets() (datatypes.Scale_Asset_Virtual_Guest, error)
	GetVirtualGuestMembers() (datatypes.Scale_Member_Virtual_Guest, error)
}

type Scale_Group_Log interface {
	Entity

	GetScaleGroup() (datatypes.Scale_Group, error)
}

type Scale_Group_Status interface {
	Entity

	GetAllObjects() (datatypes.Scale_Group_Status, error)
	GetObject() (datatypes.Scale_Group_Status, error)
}

type Scale_LoadBalancer interface {
	Entity

	CreateObject(templateObject datatypes.Scale_LoadBalancer) (datatypes.Scale_LoadBalancer, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Scale_LoadBalancer) (bool, error)
	GetObject() (datatypes.Scale_LoadBalancer, error)

	GetAllocationPercent() (int, error)
	GetHealthCheck() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Check, error)
	GetRoutingMethod() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Routing_Method, error)
	GetRoutingType() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Routing_Type, error)
	GetScaleGroup() (datatypes.Scale_Group, error)
	GetVirtualIpAddressId() (int, error)
	GetVirtualServer() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualServer, error)
	GetVirtualServerPort() (int, error)
}

type Scale_Member interface {
	Entity

	DeleteObject() (bool, error)
	GetObject() (datatypes.Scale_Member, error)

	GetScaleGroup() (datatypes.Scale_Group, error)
}

type Scale_Member_Virtual_Guest interface {
	Scale_Member

	DeleteObject() (bool, error)
	GetObject() (datatypes.Scale_Member_Virtual_Guest, error)

	GetVirtualGuest() (datatypes.Virtual_Guest, error)
	GetVirtualGuestId() (int, error)
}

type Scale_Network_Vlan interface {
	Entity

	CreateObject(templateObject datatypes.Scale_Network_Vlan) (datatypes.Scale_Network_Vlan, error)
	DeleteObject() (bool, error)
	GetObject() (datatypes.Scale_Network_Vlan, error)

	GetNetworkVlan() (datatypes.Network_Vlan, error)
	GetScaleGroup() (datatypes.Scale_Group, error)
}

type Scale_Policy interface {
	Entity

	CreateObject(templateObject datatypes.Scale_Policy) (datatypes.Scale_Policy, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Scale_Policy) (bool, error)
	GetObject() (datatypes.Scale_Policy, error)
	Trigger() (datatypes.Scale_Member, error)

	GetActions() (datatypes.Scale_Policy_Action, error)
	GetOneTimeTriggers() (datatypes.Scale_Policy_Trigger_OneTime, error)
	GetRepeatingTriggers() (datatypes.Scale_Policy_Trigger_Repeating, error)
	GetResourceUseTriggers() (datatypes.Scale_Policy_Trigger_ResourceUse, error)
	GetScaleActions() (datatypes.Scale_Policy_Action_Scale, error)
	GetScaleGroup() (datatypes.Scale_Group, error)
	GetTriggers() (datatypes.Scale_Policy_Trigger, error)
}

type Scale_Policy_Action interface {
	Entity

	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Scale_Policy_Action) (bool, error)
	GetObject() (datatypes.Scale_Policy_Action, error)

	GetScalePolicy() (datatypes.Scale_Policy, error)
	GetType() (datatypes.Scale_Policy_Action_Type, error)
}

type Scale_Policy_Action_Scale interface {
	Scale_Policy_Action

	CreateObject(templateObject datatypes.Scale_Policy_Action_Scale) (datatypes.Scale_Policy_Action_Scale, error)
	GetObject() (datatypes.Scale_Policy_Action_Scale, error)
}

type Scale_Policy_Action_Type interface {
	Entity

	GetAllObjects() (datatypes.Scale_Policy_Action_Type, error)
	GetObject() (datatypes.Scale_Policy_Action_Type, error)
}

type Scale_Policy_Trigger interface {
	Entity

	CreateObject(templateObject datatypes.Scale_Policy_Trigger) (datatypes.Scale_Policy_Trigger, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Scale_Policy_Trigger) (bool, error)
	GetObject() (datatypes.Scale_Policy_Trigger, error)

	GetScalePolicy() (datatypes.Scale_Policy, error)
	GetType() (datatypes.Scale_Policy_Trigger_Type, error)
}

type Scale_Policy_Trigger_OneTime interface {
	Scale_Policy_Trigger

	CreateObject(templateObject datatypes.Scale_Policy_Trigger_OneTime) (datatypes.Scale_Policy_Trigger_OneTime, error)
	GetObject() (datatypes.Scale_Policy_Trigger_OneTime, error)
}

type Scale_Policy_Trigger_Repeating interface {
	Scale_Policy_Trigger

	CreateObject(templateObject datatypes.Scale_Policy_Trigger_Repeating) (datatypes.Scale_Policy_Trigger_Repeating, error)
	GetObject() (datatypes.Scale_Policy_Trigger_Repeating, error)
	ValidateCronExpression(expression string) error
}

type Scale_Policy_Trigger_ResourceUse interface {
	Scale_Policy_Trigger

	CreateObject(templateObject datatypes.Scale_Policy_Trigger_ResourceUse) (datatypes.Scale_Policy_Trigger_ResourceUse, error)
	GetObject() (datatypes.Scale_Policy_Trigger_ResourceUse, error)

	GetWatches() (datatypes.Scale_Policy_Trigger_ResourceUse_Watch, error)
}

type Scale_Policy_Trigger_ResourceUse_Watch interface {
	Entity

	CreateObject(templateObject datatypes.Scale_Policy_Trigger_ResourceUse_Watch) (datatypes.Scale_Policy_Trigger_ResourceUse_Watch, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Scale_Policy_Trigger_ResourceUse_Watch) (bool, error)
	GetAllPossibleAlgorithms() (string, error)
	GetAllPossibleMetrics() (string, error)
	GetAllPossibleOperators() (string, error)
	GetObject() (datatypes.Scale_Policy_Trigger_ResourceUse_Watch, error)

	GetScalePolicyTrigger() (datatypes.Scale_Policy_Trigger_ResourceUse, error)
}

type Scale_Policy_Trigger_Type interface {
	Entity

	GetAllObjects() (datatypes.Scale_Policy_Trigger_Type, error)
	GetObject() (datatypes.Scale_Policy_Trigger_Type, error)
}

type Scale_Termination_Policy interface {
	Entity

	GetAllObjects() (datatypes.Scale_Termination_Policy, error)
	GetObject() (datatypes.Scale_Termination_Policy, error)
}
