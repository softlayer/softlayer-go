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

type Resource_Group struct {
	Entity

	AncestorGroupCount  *uint                      `json:"ancestorGroupCount,omitempty"`
	AncestorGroups      []Resource_Group           `json:"ancestorGroups,omitempty"`
	AttributeCount      *uint                      `json:"attributeCount,omitempty"`
	Attributes          []Resource_Group_Attribute `json:"attributes,omitempty"`
	CreateDate          *time.Time                 `json:"createDate,omitempty"`
	Description         *string                    `json:"description,omitempty"`
	HardwareMemberCount *uint                      `json:"hardwareMemberCount,omitempty"`
	HardwareMembers     []Resource_Group_Member    `json:"hardwareMembers,omitempty"`
	Id                  *int                       `json:"id,omitempty"`
	KeyName             *string                    `json:"keyName,omitempty"`
	MemberCount         *uint                      `json:"memberCount,omitempty"`
	Members             []Resource_Group_Member    `json:"members,omitempty"`
	Name                *string                    `json:"name,omitempty"`
	RootResourceGroup   *Resource_Group            `json:"rootResourceGroup,omitempty"`
	RootResourceGroupId *int                       `json:"rootResourceGroupId,omitempty"`
	SubnetMemberCount   *uint                      `json:"subnetMemberCount,omitempty"`
	SubnetMembers       []Resource_Group_Member    `json:"subnetMembers,omitempty"`
	Template            *Resource_Group_Template   `json:"template,omitempty"`
	TemplateId          *int                       `json:"templateId,omitempty"`
	VlanMemberCount     *uint                      `json:"vlanMemberCount,omitempty"`
	VlanMembers         []Resource_Group_Member    `json:"vlanMembers,omitempty"`
}

type Resource_Group_Attribute struct {
	Entity

	CreateDate *time.Time                     `json:"createDate,omitempty"`
	Group      *Resource_Group                `json:"group,omitempty"`
	Id         *int                           `json:"id,omitempty"`
	Type       *Resource_Group_Attribute_Type `json:"type,omitempty"`
	Value      *string                        `json:"value,omitempty"`
}

type Resource_Group_Attribute_Type struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	KeyName     *string `json:"keyName,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Resource_Group_Descendant_Reference struct {
	Entity

	Group       *Resource_Group        `json:"group,omitempty"`
	GroupMember *Resource_Group_Member `json:"groupMember,omitempty"`
}

type Resource_Group_Member struct {
	Entity

	AttributeCount        *uint                             `json:"attributeCount,omitempty"`
	Attributes            []Resource_Group_Member_Attribute `json:"attributes,omitempty"`
	CreateDate            *time.Time                        `json:"createDate,omitempty"`
	DescendantMemberCount *uint                             `json:"descendantMemberCount,omitempty"`
	DescendantMembers     []Resource_Group_Member           `json:"descendantMembers,omitempty"`
	Group                 *Resource_Group                   `json:"group,omitempty"`
	Id                    *int                              `json:"id,omitempty"`
	RoleCount             *uint                             `json:"roleCount,omitempty"`
	Roles                 []Resource_Group_Role             `json:"roles,omitempty"`
	Status                *string                           `json:"status,omitempty"`
	Type                  *Resource_Group_Member_Type       `json:"type,omitempty"`
}

type Resource_Group_Member_Attribute struct {
	Entity

	CreateDate *time.Time                            `json:"createDate,omitempty"`
	Id         *int                                  `json:"id,omitempty"`
	Member     *Resource_Group_Member                `json:"member,omitempty"`
	Type       *Resource_Group_Member_Attribute_Type `json:"type,omitempty"`
	Value      *string                               `json:"value,omitempty"`
}

type Resource_Group_Member_Attribute_Type struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	KeyName     *string `json:"keyName,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Resource_Group_Member_CloudStack_Version3_Cluster struct {
	Resource_Group_Member

	Resource *Resource_Group `json:"resource,omitempty"`
}

type Resource_Group_Member_CloudStack_Version3_Pod struct {
	Resource_Group_Member

	Resource *Resource_Group `json:"resource,omitempty"`
}

type Resource_Group_Member_CloudStack_Version3_Zone struct {
	Resource_Group_Member

	Resource *Resource_Group `json:"resource,omitempty"`
}

type Resource_Group_Member_Hardware struct {
	Resource_Group_Member

	Resource          *Hardware                        `json:"resource,omitempty"`
	ServerArbiterOnly *Resource_Group_Member_Attribute `json:"serverArbiterOnly,omitempty"`
	ServerHidden      *Resource_Group_Member_Attribute `json:"serverHidden,omitempty"`
	ServerPriority    *Resource_Group_Member_Attribute `json:"serverPriority,omitempty"`
	ServerSlaveDelay  *Resource_Group_Member_Attribute `json:"serverSlaveDelay,omitempty"`
	ServerTags        *Resource_Group_Member_Attribute `json:"serverTags,omitempty"`
	ServerVotes       *Resource_Group_Member_Attribute `json:"serverVotes,omitempty"`
}

type Resource_Group_Member_Network_Storage struct {
	Resource_Group_Member

	Resource *Network_Storage `json:"resource,omitempty"`
}

type Resource_Group_Member_Network_Subnet struct {
	Resource_Group_Member

	Resource *Network_Subnet `json:"resource,omitempty"`
}

type Resource_Group_Member_Network_Vlan struct {
	Resource_Group_Member

	Resource *Network_Vlan `json:"resource,omitempty"`
}

type Resource_Group_Member_Resource_Group struct {
	Resource_Group_Member

	Resource *Resource_Group `json:"resource,omitempty"`
}

type Resource_Group_Member_Role_Link struct {
	Entity

	GroupMemberId       *int `json:"groupMemberId,omitempty"`
	GroupTemplateRoleId *int `json:"groupTemplateRoleId,omitempty"`
}

type Resource_Group_Member_Software_Component_Password struct {
	Resource_Group_Member

	Resource *Software_Component_Password `json:"resource,omitempty"`
}

type Resource_Group_Member_Type struct {
	Entity

	Description *string `json:"description,omitempty"`
	KeyName     *string `json:"keyName,omitempty"`
}

type Resource_Group_Member_Virtual_Host_Pool struct {
	Resource_Group_Member
}

type Resource_Group_Role struct {
	Entity

	Description     *string                           `json:"description,omitempty"`
	Id              *int                              `json:"id,omitempty"`
	KeyName         *string                           `json:"keyName,omitempty"`
	MemberLinkCount *uint                             `json:"memberLinkCount,omitempty"`
	MemberLinks     []Resource_Group_Member_Role_Link `json:"memberLinks,omitempty"`
}

type Resource_Group_Template struct {
	Entity

	Children      []Resource_Group_Template        `json:"children,omitempty"`
	ChildrenCount *uint                            `json:"childrenCount,omitempty"`
	Description   *string                          `json:"description,omitempty"`
	Id            *int                             `json:"id,omitempty"`
	KeyName       *string                          `json:"keyName,omitempty"`
	MemberCount   *uint                            `json:"memberCount,omitempty"`
	Members       []Resource_Group_Template_Member `json:"members,omitempty"`
	Package       *Product_Package                 `json:"package,omitempty"`
}

type Resource_Group_Template_Member struct {
	Entity

	MaxQuantity *int                     `json:"maxQuantity,omitempty"`
	MinQuantity *int                     `json:"minQuantity,omitempty"`
	Role        *Resource_Group_Role     `json:"role,omitempty"`
	RoleId      *int                     `json:"roleId,omitempty"`
	Template    *Resource_Group_Template `json:"template,omitempty"`
	TemplateId  *int                     `json:"templateId,omitempty"`
}

type Resource_Metadata struct {
	Entity
}
