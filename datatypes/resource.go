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

package datatypes

// no documentation yet
type Resource_Group struct {
	Entity

	// A count of a resource group's associated group ancestors.
	AncestorGroupCount *uint `json:"ancestorGroupCount,omitempty" xmlrpc:"ancestorGroupCount"`

	// A resource group's associated group ancestors.
	AncestorGroups []Resource_Group `json:"ancestorGroups,omitempty" xmlrpc:"ancestorGroups"`

	// A count of a resource group's associated attributes.
	AttributeCount *uint `json:"attributeCount,omitempty" xmlrpc:"attributeCount"`

	// A resource group's associated attributes.
	Attributes []Resource_Group_Attribute `json:"attributes,omitempty" xmlrpc:"attributes"`

	// A resource group's creation date.
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// A resource group's description.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// A count of a resource group's associated hardware members.
	HardwareMemberCount *uint `json:"hardwareMemberCount,omitempty" xmlrpc:"hardwareMemberCount"`

	// A resource group's associated hardware members.
	HardwareMembers []Resource_Group_Member `json:"hardwareMembers,omitempty" xmlrpc:"hardwareMembers"`

	// A resource group's ID.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// A resource group's keyname.
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// A count of a resource group's associated members.
	MemberCount *uint `json:"memberCount,omitempty" xmlrpc:"memberCount"`

	// A resource group's associated members.
	Members []Resource_Group_Member `json:"members,omitempty" xmlrpc:"members"`

	// A resource group's name.
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// A resource group's associated root resource group.
	RootResourceGroup *Resource_Group `json:"rootResourceGroup,omitempty" xmlrpc:"rootResourceGroup"`

	// no documentation yet
	RootResourceGroupId *int `json:"rootResourceGroupId,omitempty" xmlrpc:"rootResourceGroupId"`

	// A count of a resource group's associated subnet members.
	SubnetMemberCount *uint `json:"subnetMemberCount,omitempty" xmlrpc:"subnetMemberCount"`

	// A resource group's associated subnet members.
	SubnetMembers []Resource_Group_Member `json:"subnetMembers,omitempty" xmlrpc:"subnetMembers"`

	// A resource group's associated template.
	Template *Resource_Group_Template `json:"template,omitempty" xmlrpc:"template"`

	// A resource group's template ID.
	TemplateId *int `json:"templateId,omitempty" xmlrpc:"templateId"`

	// A count of a resource group's associated VLAN members.
	VlanMemberCount *uint `json:"vlanMemberCount,omitempty" xmlrpc:"vlanMemberCount"`

	// A resource group's associated VLAN members.
	VlanMembers []Resource_Group_Member `json:"vlanMembers,omitempty" xmlrpc:"vlanMembers"`
}

// no documentation yet
type Resource_Group_Attribute struct {
	Entity

	// A resource group attribute's creation date.
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// A resource group attribute's resource group.
	Group *Resource_Group `json:"group,omitempty" xmlrpc:"group"`

	// A resource group attribute's ID.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// A resource group attribute's type.
	Type *Resource_Group_Attribute_Type `json:"type,omitempty" xmlrpc:"type"`

	// A resource group attribute's value.
	Value *string `json:"value,omitempty" xmlrpc:"value"`
}

// no documentation yet
type Resource_Group_Attribute_Type struct {
	Entity

	// no documentation yet
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// no documentation yet
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// The SoftLayer_Resource_Group_Descendant_Reference data type simplifies the link between one SoftLayer_Resource_Group_Member object and all of its parents.
//
//
type Resource_Group_Descendant_Reference struct {
	Entity

	// no documentation yet
	Group *Resource_Group `json:"group,omitempty" xmlrpc:"group"`

	// no documentation yet
	GroupMember *Resource_Group_Member `json:"groupMember,omitempty" xmlrpc:"groupMember"`
}

// no documentation yet
type Resource_Group_Member struct {
	Entity

	// A count of a resource group member's associated attributes.
	AttributeCount *uint `json:"attributeCount,omitempty" xmlrpc:"attributeCount"`

	// A resource group member's associated attributes.
	Attributes []Resource_Group_Member_Attribute `json:"attributes,omitempty" xmlrpc:"attributes"`

	// A resource group member's creation date.
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// A count of a resource group member's associated member descendants.
	DescendantMemberCount *uint `json:"descendantMemberCount,omitempty" xmlrpc:"descendantMemberCount"`

	// A resource group member's associated member descendants.
	DescendantMembers []Resource_Group_Member `json:"descendantMembers,omitempty" xmlrpc:"descendantMembers"`

	// A resource group member's resource group.
	Group *Resource_Group `json:"group,omitempty" xmlrpc:"group"`

	// A resource group member's ID.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// A count of a resource group member's associated roles.
	RoleCount *uint `json:"roleCount,omitempty" xmlrpc:"roleCount"`

	// A resource group member's associated roles.
	Roles []Resource_Group_Role `json:"roles,omitempty" xmlrpc:"roles"`

	// A resource group member's status.
	Status *string `json:"status,omitempty" xmlrpc:"status"`

	// A resource group member's type.
	Type *Resource_Group_Member_Type `json:"type,omitempty" xmlrpc:"type"`
}

// no documentation yet
type Resource_Group_Member_Attribute struct {
	Entity

	// A resource group member attribute's creation date.
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// A resource group member attribute's ID.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// A resource group member attribute's resource group member.
	Member *Resource_Group_Member `json:"member,omitempty" xmlrpc:"member"`

	// A resource group member attribute's type.
	Type *Resource_Group_Member_Attribute_Type `json:"type,omitempty" xmlrpc:"type"`

	// A resource group member attribute's value.
	Value *string `json:"value,omitempty" xmlrpc:"value"`
}

// no documentation yet
type Resource_Group_Member_Attribute_Type struct {
	Entity

	// no documentation yet
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// no documentation yet
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// no documentation yet
type Resource_Group_Member_CloudStack_Version3_Cluster struct {
	Resource_Group_Member

	// A resource group member's associated cluster.
	Resource *Resource_Group `json:"resource,omitempty" xmlrpc:"resource"`
}

// no documentation yet
type Resource_Group_Member_CloudStack_Version3_Pod struct {
	Resource_Group_Member

	// A resource group member's associated pod.
	Resource *Resource_Group `json:"resource,omitempty" xmlrpc:"resource"`
}

// no documentation yet
type Resource_Group_Member_CloudStack_Version3_Zone struct {
	Resource_Group_Member

	// A resource group member's associated zone.
	Resource *Resource_Group `json:"resource,omitempty" xmlrpc:"resource"`
}

// no documentation yet
type Resource_Group_Member_Hardware struct {
	Resource_Group_Member

	// A resource group member's associated hardware.
	Resource *Hardware `json:"resource,omitempty" xmlrpc:"resource"`

	// A resource group hardware member's associated server arbiter-only state.
	ServerArbiterOnly *Resource_Group_Member_Attribute `json:"serverArbiterOnly,omitempty" xmlrpc:"serverArbiterOnly"`

	// A resource group hardware member's associated server hidden state.
	ServerHidden *Resource_Group_Member_Attribute `json:"serverHidden,omitempty" xmlrpc:"serverHidden"`

	// A resource group hardware member's associated server priority.
	ServerPriority *Resource_Group_Member_Attribute `json:"serverPriority,omitempty" xmlrpc:"serverPriority"`

	// A resource group hardware member's associated server slave delay (in seconds).
	ServerSlaveDelay *Resource_Group_Member_Attribute `json:"serverSlaveDelay,omitempty" xmlrpc:"serverSlaveDelay"`

	// A resource group hardware member's associated server tags (in JSON format).
	ServerTags *Resource_Group_Member_Attribute `json:"serverTags,omitempty" xmlrpc:"serverTags"`

	// A resource group hardware member's associated server vote count.
	ServerVotes *Resource_Group_Member_Attribute `json:"serverVotes,omitempty" xmlrpc:"serverVotes"`
}

// no documentation yet
type Resource_Group_Member_Network_Storage struct {
	Resource_Group_Member

	// A resource group member's associated network storage.
	Resource *Network_Storage `json:"resource,omitempty" xmlrpc:"resource"`
}

// no documentation yet
type Resource_Group_Member_Network_Subnet struct {
	Resource_Group_Member

	// A resource group member's associated network subnet.
	Resource *Network_Subnet `json:"resource,omitempty" xmlrpc:"resource"`
}

// no documentation yet
type Resource_Group_Member_Network_Vlan struct {
	Resource_Group_Member

	// A resource group member's associated network VLAN.
	Resource *Network_Vlan `json:"resource,omitempty" xmlrpc:"resource"`
}

// no documentation yet
type Resource_Group_Member_Resource_Group struct {
	Resource_Group_Member

	// A resource group member's associated resource group.
	Resource *Resource_Group `json:"resource,omitempty" xmlrpc:"resource"`
}

// no documentation yet
type Resource_Group_Member_Role_Link struct {
	Entity

	// A resource group member's ID.
	GroupMemberId *int `json:"groupMemberId,omitempty" xmlrpc:"groupMemberId"`

	// A resource group's template role ID.
	GroupTemplateRoleId *int `json:"groupTemplateRoleId,omitempty" xmlrpc:"groupTemplateRoleId"`
}

// no documentation yet
type Resource_Group_Member_Software_Component_Password struct {
	Resource_Group_Member

	// A resource group member's associated software component password.
	Resource *Software_Component_Password `json:"resource,omitempty" xmlrpc:"resource"`
}

// no documentation yet
type Resource_Group_Member_Type struct {
	Entity

	// A resource group member's type description.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// A resource group member's type keyname.
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`
}

// no documentation yet
type Resource_Group_Member_Virtual_Host_Pool struct {
	Resource_Group_Member
}

// no documentation yet
type Resource_Group_Role struct {
	Entity

	// A resource group role's description.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// A resource group role's ID.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// A resource group role's keyname.
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// A count of a resource group's role.
	MemberLinkCount *uint `json:"memberLinkCount,omitempty" xmlrpc:"memberLinkCount"`

	// A resource group's role.
	MemberLinks []Resource_Group_Member_Role_Link `json:"memberLinks,omitempty" xmlrpc:"memberLinks"`
}

// no documentation yet
type Resource_Group_Template struct {
	Entity

	// no documentation yet
	Children []Resource_Group_Template `json:"children,omitempty" xmlrpc:"children"`

	// A count of
	ChildrenCount *uint `json:"childrenCount,omitempty" xmlrpc:"childrenCount"`

	// A resource group template's description.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// no documentation yet
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// A resource group template's keyname.
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// A count of
	MemberCount *uint `json:"memberCount,omitempty" xmlrpc:"memberCount"`

	// no documentation yet
	Members []Resource_Group_Template_Member `json:"members,omitempty" xmlrpc:"members"`

	// no documentation yet
	Package *Product_Package `json:"package,omitempty" xmlrpc:"package"`
}

// no documentation yet
type Resource_Group_Template_Member struct {
	Entity

	// no documentation yet
	MaxQuantity *int `json:"maxQuantity,omitempty" xmlrpc:"maxQuantity"`

	// no documentation yet
	MinQuantity *int `json:"minQuantity,omitempty" xmlrpc:"minQuantity"`

	// no documentation yet
	Role *Resource_Group_Role `json:"role,omitempty" xmlrpc:"role"`

	// no documentation yet
	RoleId *int `json:"roleId,omitempty" xmlrpc:"roleId"`

	// no documentation yet
	Template *Resource_Group_Template `json:"template,omitempty" xmlrpc:"template"`

	// no documentation yet
	TemplateId *int `json:"templateId,omitempty" xmlrpc:"templateId"`
}

// no documentation yet
type Resource_Metadata struct {
	Entity
}
