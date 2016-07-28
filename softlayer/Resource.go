package softlayer

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Resource_Group interface {
	Entity

	EditObject(templateObject datatypes.Resource_Group) (bool, error)
	GetObject() (datatypes.Resource_Group, error)

	GetAncestorGroups() (datatypes.Resource_Group, error)
	GetAttributes() (datatypes.Resource_Group_Attribute, error)
	GetHardwareMembers() (datatypes.Resource_Group_Member, error)
	GetMembers() (datatypes.Resource_Group_Member, error)
	GetRootResourceGroup() (datatypes.Resource_Group, error)
	GetSubnetMembers() (datatypes.Resource_Group_Member, error)
	GetTemplate() (datatypes.Resource_Group_Template, error)
	GetVlanMembers() (datatypes.Resource_Group_Member, error)
}

type Resource_Group_Attribute interface {
	Entity

	GetGroup() (datatypes.Resource_Group, error)
	GetType() (datatypes.Resource_Group_Attribute_Type, error)
}

type Resource_Group_Attribute_Type interface {
	Entity
}

type Resource_Group_Descendant_Reference interface {
	Entity

	GetGroup() (datatypes.Resource_Group, error)
	GetGroupMember() (datatypes.Resource_Group_Member, error)
}

type Resource_Group_Member interface {
	Entity

	GetAttributes() (datatypes.Resource_Group_Member_Attribute, error)
	GetDescendantMembers() (datatypes.Resource_Group_Member, error)
	GetGroup() (datatypes.Resource_Group, error)
	GetRoles() (datatypes.Resource_Group_Role, error)
	GetType() (datatypes.Resource_Group_Member_Type, error)
}

type Resource_Group_Member_Attribute interface {
	Entity

	GetMember() (datatypes.Resource_Group_Member, error)
	GetType() (datatypes.Resource_Group_Member_Attribute_Type, error)
}

type Resource_Group_Member_Attribute_Type interface {
	Entity
}

type Resource_Group_Member_CloudStack_Version3_Cluster interface {
	Resource_Group_Member

	GetResource() (datatypes.Resource_Group, error)
}

type Resource_Group_Member_CloudStack_Version3_Pod interface {
	Resource_Group_Member

	GetResource() (datatypes.Resource_Group, error)
}

type Resource_Group_Member_CloudStack_Version3_Zone interface {
	Resource_Group_Member

	GetResource() (datatypes.Resource_Group, error)
}

type Resource_Group_Member_Hardware interface {
	Resource_Group_Member

	GetResource() (datatypes.Hardware, error)
	GetServerArbiterOnly() (datatypes.Resource_Group_Member_Attribute, error)
	GetServerHidden() (datatypes.Resource_Group_Member_Attribute, error)
	GetServerPriority() (datatypes.Resource_Group_Member_Attribute, error)
	GetServerSlaveDelay() (datatypes.Resource_Group_Member_Attribute, error)
	GetServerTags() (datatypes.Resource_Group_Member_Attribute, error)
	GetServerVotes() (datatypes.Resource_Group_Member_Attribute, error)
}

type Resource_Group_Member_Network_Storage interface {
	Resource_Group_Member

	GetResource() (datatypes.Network_Storage, error)
}

type Resource_Group_Member_Network_Subnet interface {
	Resource_Group_Member

	GetResource() (datatypes.Network_Subnet, error)
}

type Resource_Group_Member_Network_Vlan interface {
	Resource_Group_Member

	GetResource() (datatypes.Network_Vlan, error)
}

type Resource_Group_Member_Resource_Group interface {
	Resource_Group_Member

	GetResource() (datatypes.Resource_Group, error)
}

type Resource_Group_Member_Role_Link interface {
	Entity
}

type Resource_Group_Member_Software_Component_Password interface {
	Resource_Group_Member

	GetResource() (datatypes.Software_Component_Password, error)
}

type Resource_Group_Member_Type interface {
	Entity
}

type Resource_Group_Member_Virtual_Host_Pool interface {
	Resource_Group_Member
}

type Resource_Group_Role interface {
	Entity

	GetMemberLinks() (datatypes.Resource_Group_Member_Role_Link, error)
}

type Resource_Group_Template interface {
	Entity

	GetAllObjects() (datatypes.Resource_Group_Template, error)
	GetObject() (datatypes.Resource_Group_Template, error)

	GetChildren() (datatypes.Resource_Group_Template, error)
	GetMembers() (datatypes.Resource_Group_Template_Member, error)
	GetPackage() (datatypes.Product_Package, error)
}

type Resource_Group_Template_Member interface {
	Entity

	GetRole() (datatypes.Resource_Group_Role, error)
	GetTemplate() (datatypes.Resource_Group_Template, error)
}

type Resource_Metadata interface {
	Entity

	GetBackendMacAddresses() (string, error)
	GetDatacenter() (string, error)
	GetDatacenterId() (int, error)
	GetDomain() (string, error)
	GetFrontendMacAddresses() (string, error)
	GetFullyQualifiedDomainName() (string, error)
	GetHostname() (string, error)
	GetId() (int, error)
	GetPrimaryBackendIpAddress() (string, error)
	GetPrimaryIpAddress() (string, error)
	GetProvisionState() (string, error)
	GetRouter(macAddress string) (string, error)
	GetServiceResource(serviceName string, index int) (string, error)
	GetServiceResources() (datatypes.Container_Resource_Metadata_ServiceResource, error)
	GetTags() (string, error)
	GetUserMetadata() (string, error)
	GetVlanIds(macAddress string) (int, error)
	GetVlans(macAddress string) (int, error)
}
