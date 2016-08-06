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

//
type Resource_Group struct {
	Entity

	// A count of a resource group's associated group ancestors.
	AncestorGroupCount *uint `json:"ancestorGroupCount,omitempty"`

	// A resource group's associated group ancestors.
	AncestorGroups []Resource_Group `json:"ancestorGroups,omitempty"`

	// A count of a resource group's associated attributes.
	AttributeCount *uint `json:"attributeCount,omitempty"`

	// A resource group's associated attributes.
	Attributes []Resource_Group_Attribute `json:"attributes,omitempty"`

	// A resource group's creation date.
	CreateDate *Time `json:"createDate,omitempty"`

	// A resource group's description.
	Description *string `json:"description,omitempty"`

	// A count of a resource group's associated hardware members.
	HardwareMemberCount *uint `json:"hardwareMemberCount,omitempty"`

	// A resource group's associated hardware members.
	HardwareMembers []Resource_Group_Member `json:"hardwareMembers,omitempty"`

	// A resource group's ID.
	Id *int `json:"id,omitempty"`

	// A resource group's keyname.
	KeyName *string `json:"keyName,omitempty"`

	// A count of a resource group's associated members.
	MemberCount *uint `json:"memberCount,omitempty"`

	// A resource group's associated members.
	Members []Resource_Group_Member `json:"members,omitempty"`

	// A resource group's name.
	Name *string `json:"name,omitempty"`

	// A resource group's associated root resource group.
	RootResourceGroup *Resource_Group `json:"rootResourceGroup,omitempty"`

	//
	RootResourceGroupId *int `json:"rootResourceGroupId,omitempty"`

	// A count of a resource group's associated subnet members.
	SubnetMemberCount *uint `json:"subnetMemberCount,omitempty"`

	// A resource group's associated subnet members.
	SubnetMembers []Resource_Group_Member `json:"subnetMembers,omitempty"`

	// A resource group's associated template.
	Template *Resource_Group_Template `json:"template,omitempty"`

	// A resource group's template ID.
	TemplateId *int `json:"templateId,omitempty"`

	// A count of a resource group's associated VLAN members.
	VlanMemberCount *uint `json:"vlanMemberCount,omitempty"`

	// A resource group's associated VLAN members.
	VlanMembers []Resource_Group_Member `json:"vlanMembers,omitempty"`
}

//
type Resource_Group_Attribute struct {
	Entity

	// A resource group attribute's creation date.
	CreateDate *Time `json:"createDate,omitempty"`

	// A resource group attribute's resource group.
	Group *Resource_Group `json:"group,omitempty"`

	// A resource group attribute's ID.
	Id *int `json:"id,omitempty"`

	// A resource group attribute's type.
	Type *Resource_Group_Attribute_Type `json:"type,omitempty"`

	// A resource group attribute's value.
	Value *string `json:"value,omitempty"`
}

//
type Resource_Group_Attribute_Type struct {
	Entity

	//
	Description *string `json:"description,omitempty"`

	//
	Id *int `json:"id,omitempty"`

	//
	KeyName *string `json:"keyName,omitempty"`

	//
	Name *string `json:"name,omitempty"`
}

// The SoftLayer_Resource_Group_Descendant_Reference data type simplifies the link between one SoftLayer_Resource_Group_Member object and all of its parents.
//
//
type Resource_Group_Descendant_Reference struct {
	Entity

	//
	Group *Resource_Group `json:"group,omitempty"`

	//
	GroupMember *Resource_Group_Member `json:"groupMember,omitempty"`
}

//
type Resource_Group_Member struct {
	Entity

	// A count of a resource group member's associated attributes.
	AttributeCount *uint `json:"attributeCount,omitempty"`

	// A resource group member's associated attributes.
	Attributes []Resource_Group_Member_Attribute `json:"attributes,omitempty"`

	// A resource group member's creation date.
	CreateDate *Time `json:"createDate,omitempty"`

	// A count of a resource group member's associated member descendants.
	DescendantMemberCount *uint `json:"descendantMemberCount,omitempty"`

	// A resource group member's associated member descendants.
	DescendantMembers []Resource_Group_Member `json:"descendantMembers,omitempty"`

	// A resource group member's resource group.
	Group *Resource_Group `json:"group,omitempty"`

	// A resource group member's ID.
	Id *int `json:"id,omitempty"`

	// A count of a resource group member's associated roles.
	RoleCount *uint `json:"roleCount,omitempty"`

	// A resource group member's associated roles.
	Roles []Resource_Group_Role `json:"roles,omitempty"`

	// A resource group member's status.
	Status *string `json:"status,omitempty"`

	// A resource group member's type.
	Type *Resource_Group_Member_Type `json:"type,omitempty"`
}

//
type Resource_Group_Member_Attribute struct {
	Entity

	// A resource group member attribute's creation date.
	CreateDate *Time `json:"createDate,omitempty"`

	// A resource group member attribute's ID.
	Id *int `json:"id,omitempty"`

	// A resource group member attribute's resource group member.
	Member *Resource_Group_Member `json:"member,omitempty"`

	// A resource group member attribute's type.
	Type *Resource_Group_Member_Attribute_Type `json:"type,omitempty"`

	// A resource group member attribute's value.
	Value *string `json:"value,omitempty"`
}

//
type Resource_Group_Member_Attribute_Type struct {
	Entity

	//
	Description *string `json:"description,omitempty"`

	//
	Id *int `json:"id,omitempty"`

	//
	KeyName *string `json:"keyName,omitempty"`

	//
	Name *string `json:"name,omitempty"`
}

//
type Resource_Group_Member_CloudStack_Version3_Cluster struct {
	Resource_Group_Member

	// A resource group member's associated cluster.
	Resource *Resource_Group `json:"resource,omitempty"`
}

//
type Resource_Group_Member_CloudStack_Version3_Pod struct {
	Resource_Group_Member

	// A resource group member's associated pod.
	Resource *Resource_Group `json:"resource,omitempty"`
}

//
type Resource_Group_Member_CloudStack_Version3_Zone struct {
	Resource_Group_Member

	// A resource group member's associated zone.
	Resource *Resource_Group `json:"resource,omitempty"`
}

//
type Resource_Group_Member_Hardware struct {
	Resource_Group_Member

	// A resource group member's associated hardware.
	Resource *Hardware `json:"resource,omitempty"`

	// A resource group hardware member's associated server arbiter-only state.
	ServerArbiterOnly *Resource_Group_Member_Attribute `json:"serverArbiterOnly,omitempty"`

	// A resource group hardware member's associated server hidden state.
	ServerHidden *Resource_Group_Member_Attribute `json:"serverHidden,omitempty"`

	// A resource group hardware member's associated server priority.
	ServerPriority *Resource_Group_Member_Attribute `json:"serverPriority,omitempty"`

	// A resource group hardware member's associated server slave delay (in seconds).
	ServerSlaveDelay *Resource_Group_Member_Attribute `json:"serverSlaveDelay,omitempty"`

	// A resource group hardware member's associated server tags (in JSON format).
	ServerTags *Resource_Group_Member_Attribute `json:"serverTags,omitempty"`

	// A resource group hardware member's associated server vote count.
	ServerVotes *Resource_Group_Member_Attribute `json:"serverVotes,omitempty"`
}

//
type Resource_Group_Member_Network_Storage struct {
	Resource_Group_Member

	// A resource group member's associated network storage.
	Resource *Network_Storage `json:"resource,omitempty"`
}

//
type Resource_Group_Member_Network_Subnet struct {
	Resource_Group_Member

	// A resource group member's associated network subnet.
	Resource *Network_Subnet `json:"resource,omitempty"`
}

//
type Resource_Group_Member_Network_Vlan struct {
	Resource_Group_Member

	// A resource group member's associated network VLAN.
	Resource *Network_Vlan `json:"resource,omitempty"`
}

//
type Resource_Group_Member_Resource_Group struct {
	Resource_Group_Member

	// A resource group member's associated resource group.
	Resource *Resource_Group `json:"resource,omitempty"`
}

//
type Resource_Group_Member_Role_Link struct {
	Entity

	// A resource group member's ID.
	GroupMemberId *int `json:"groupMemberId,omitempty"`

	// A resource group's template role ID.
	GroupTemplateRoleId *int `json:"groupTemplateRoleId,omitempty"`
}

//
type Resource_Group_Member_Software_Component_Password struct {
	Resource_Group_Member

	// A resource group member's associated software component password.
	Resource *Software_Component_Password `json:"resource,omitempty"`
}

//
type Resource_Group_Member_Type struct {
	Entity

	// A resource group member's type description.
	Description *string `json:"description,omitempty"`

	// A resource group member's type keyname.
	KeyName *string `json:"keyName,omitempty"`
}

//
type Resource_Group_Member_Virtual_Host_Pool struct {
	Resource_Group_Member
}

//
type Resource_Group_Role struct {
	Entity

	// A resource group role's description.
	Description *string `json:"description,omitempty"`

	// A resource group role's ID.
	Id *int `json:"id,omitempty"`

	// A resource group role's keyname.
	KeyName *string `json:"keyName,omitempty"`

	// A count of a resource group's role.
	MemberLinkCount *uint `json:"memberLinkCount,omitempty"`

	// A resource group's role.
	MemberLinks []Resource_Group_Member_Role_Link `json:"memberLinks,omitempty"`
}

//
type Resource_Group_Template struct {
	Entity

	//
	Children []Resource_Group_Template `json:"children,omitempty"`

	// A count of
	ChildrenCount *uint `json:"childrenCount,omitempty"`

	// A resource group template's description.
	Description *string `json:"description,omitempty"`

	//
	Id *int `json:"id,omitempty"`

	// A resource group template's keyname.
	KeyName *string `json:"keyName,omitempty"`

	// A count of
	MemberCount *uint `json:"memberCount,omitempty"`

	//
	Members []Resource_Group_Template_Member `json:"members,omitempty"`

	//
	Package *Product_Package `json:"package,omitempty"`
}

//
type Resource_Group_Template_Member struct {
	Entity

	//
	MaxQuantity *int `json:"maxQuantity,omitempty"`

	//
	MinQuantity *int `json:"minQuantity,omitempty"`

	//
	Role *Resource_Group_Role `json:"role,omitempty"`

	//
	RoleId *int `json:"roleId,omitempty"`

	//
	Template *Resource_Group_Template `json:"template,omitempty"`

	//
	TemplateId *int `json:"templateId,omitempty"`
}

//
type Resource_Metadata struct {
	Entity
}
