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

type Resource_Group struct {
	Session *Session
	Options
}

func (r *Session) GetResourceGroupService() Resource_Group {
	return Resource_Group{Session: r}
}

func (r *Resource_Group) EditObject(templateObject *datatypes.Resource_Group) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Group) GetObject() (resp datatypes.Resource_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Resource_Group) GetAncestorGroups() (resp []datatypes.Resource_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Group) GetAttributes() (resp []datatypes.Resource_Group_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Group) GetHardwareMembers() (resp []datatypes.Resource_Group_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Group) GetMembers() (resp []datatypes.Resource_Group_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Group) GetRootResourceGroup() (resp datatypes.Resource_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Group) GetSubnetMembers() (resp []datatypes.Resource_Group_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Group) GetTemplate() (resp datatypes.Resource_Group_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Group) GetVlanMembers() (resp []datatypes.Resource_Group_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Resource_Group_Attribute struct {
	Session *Session
	Options
}

func (r *Session) GetResourceGroupAttributeService() Resource_Group_Attribute {
	return Resource_Group_Attribute{Session: r}
}

func (r *Resource_Group_Attribute) GetGroup() (resp datatypes.Resource_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Group_Attribute) GetType() (resp datatypes.Resource_Group_Attribute_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Resource_Group_Attribute_Type struct {
	Session *Session
	Options
}

func (r *Session) GetResourceGroupAttributeTypeService() Resource_Group_Attribute_Type {
	return Resource_Group_Attribute_Type{Session: r}
}

type Resource_Group_Descendant_Reference struct {
	Session *Session
	Options
}

func (r *Session) GetResourceGroupDescendantReferenceService() Resource_Group_Descendant_Reference {
	return Resource_Group_Descendant_Reference{Session: r}
}

func (r *Resource_Group_Descendant_Reference) GetGroup() (resp datatypes.Resource_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Group_Descendant_Reference) GetGroupMember() (resp datatypes.Resource_Group_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Resource_Group_Member struct {
	Session *Session
	Options
}

func (r *Session) GetResourceGroupMemberService() Resource_Group_Member {
	return Resource_Group_Member{Session: r}
}

func (r *Resource_Group_Member) GetAttributes() (resp []datatypes.Resource_Group_Member_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Group_Member) GetDescendantMembers() (resp []datatypes.Resource_Group_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Group_Member) GetGroup() (resp datatypes.Resource_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Group_Member) GetRoles() (resp []datatypes.Resource_Group_Role, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Group_Member) GetType() (resp datatypes.Resource_Group_Member_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Resource_Group_Member_Attribute struct {
	Session *Session
	Options
}

func (r *Session) GetResourceGroupMemberAttributeService() Resource_Group_Member_Attribute {
	return Resource_Group_Member_Attribute{Session: r}
}

func (r *Resource_Group_Member_Attribute) GetMember() (resp datatypes.Resource_Group_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Group_Member_Attribute) GetType() (resp datatypes.Resource_Group_Member_Attribute_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Resource_Group_Member_Attribute_Type struct {
	Session *Session
	Options
}

func (r *Session) GetResourceGroupMemberAttributeTypeService() Resource_Group_Member_Attribute_Type {
	return Resource_Group_Member_Attribute_Type{Session: r}
}

type Resource_Group_Member_CloudStack_Version3_Cluster struct {
	Session *Session
	Options
}

func (r *Session) GetResourceGroupMemberCloudStackVersion3ClusterService() Resource_Group_Member_CloudStack_Version3_Cluster {
	return Resource_Group_Member_CloudStack_Version3_Cluster{Session: r}
}

func (r *Resource_Group_Member_CloudStack_Version3_Cluster) GetResource() (resp datatypes.Resource_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Resource_Group_Member_CloudStack_Version3_Pod struct {
	Session *Session
	Options
}

func (r *Session) GetResourceGroupMemberCloudStackVersion3PodService() Resource_Group_Member_CloudStack_Version3_Pod {
	return Resource_Group_Member_CloudStack_Version3_Pod{Session: r}
}

func (r *Resource_Group_Member_CloudStack_Version3_Pod) GetResource() (resp datatypes.Resource_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Resource_Group_Member_CloudStack_Version3_Zone struct {
	Session *Session
	Options
}

func (r *Session) GetResourceGroupMemberCloudStackVersion3ZoneService() Resource_Group_Member_CloudStack_Version3_Zone {
	return Resource_Group_Member_CloudStack_Version3_Zone{Session: r}
}

func (r *Resource_Group_Member_CloudStack_Version3_Zone) GetResource() (resp datatypes.Resource_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Resource_Group_Member_Hardware struct {
	Session *Session
	Options
}

func (r *Session) GetResourceGroupMemberHardwareService() Resource_Group_Member_Hardware {
	return Resource_Group_Member_Hardware{Session: r}
}

func (r *Resource_Group_Member_Hardware) GetResource() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Group_Member_Hardware) GetServerArbiterOnly() (resp datatypes.Resource_Group_Member_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Group_Member_Hardware) GetServerHidden() (resp datatypes.Resource_Group_Member_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Group_Member_Hardware) GetServerPriority() (resp datatypes.Resource_Group_Member_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Group_Member_Hardware) GetServerSlaveDelay() (resp datatypes.Resource_Group_Member_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Group_Member_Hardware) GetServerTags() (resp datatypes.Resource_Group_Member_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Group_Member_Hardware) GetServerVotes() (resp datatypes.Resource_Group_Member_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Resource_Group_Member_Network_Storage struct {
	Session *Session
	Options
}

func (r *Session) GetResourceGroupMemberNetworkStorageService() Resource_Group_Member_Network_Storage {
	return Resource_Group_Member_Network_Storage{Session: r}
}

func (r *Resource_Group_Member_Network_Storage) GetResource() (resp datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Resource_Group_Member_Network_Subnet struct {
	Session *Session
	Options
}

func (r *Session) GetResourceGroupMemberNetworkSubnetService() Resource_Group_Member_Network_Subnet {
	return Resource_Group_Member_Network_Subnet{Session: r}
}

func (r *Resource_Group_Member_Network_Subnet) GetResource() (resp datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Resource_Group_Member_Network_Vlan struct {
	Session *Session
	Options
}

func (r *Session) GetResourceGroupMemberNetworkVlanService() Resource_Group_Member_Network_Vlan {
	return Resource_Group_Member_Network_Vlan{Session: r}
}

func (r *Resource_Group_Member_Network_Vlan) GetResource() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Resource_Group_Member_Resource_Group struct {
	Session *Session
	Options
}

func (r *Session) GetResourceGroupMemberResourceGroupService() Resource_Group_Member_Resource_Group {
	return Resource_Group_Member_Resource_Group{Session: r}
}

func (r *Resource_Group_Member_Resource_Group) GetResource() (resp datatypes.Resource_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Resource_Group_Member_Role_Link struct {
	Session *Session
	Options
}

func (r *Session) GetResourceGroupMemberRoleLinkService() Resource_Group_Member_Role_Link {
	return Resource_Group_Member_Role_Link{Session: r}
}

type Resource_Group_Member_Software_Component_Password struct {
	Session *Session
	Options
}

func (r *Session) GetResourceGroupMemberSoftwareComponentPasswordService() Resource_Group_Member_Software_Component_Password {
	return Resource_Group_Member_Software_Component_Password{Session: r}
}

func (r *Resource_Group_Member_Software_Component_Password) GetResource() (resp datatypes.Software_Component_Password, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Resource_Group_Member_Type struct {
	Session *Session
	Options
}

func (r *Session) GetResourceGroupMemberTypeService() Resource_Group_Member_Type {
	return Resource_Group_Member_Type{Session: r}
}

type Resource_Group_Member_Virtual_Host_Pool struct {
	Session *Session
	Options
}

func (r *Session) GetResourceGroupMemberVirtualHostPoolService() Resource_Group_Member_Virtual_Host_Pool {
	return Resource_Group_Member_Virtual_Host_Pool{Session: r}
}

type Resource_Group_Role struct {
	Session *Session
	Options
}

func (r *Session) GetResourceGroupRoleService() Resource_Group_Role {
	return Resource_Group_Role{Session: r}
}

func (r *Resource_Group_Role) GetMemberLinks() (resp []datatypes.Resource_Group_Member_Role_Link, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Resource_Group_Template struct {
	Session *Session
	Options
}

func (r *Session) GetResourceGroupTemplateService() Resource_Group_Template {
	return Resource_Group_Template{Session: r}
}

func (r *Resource_Group_Template) GetAllObjects() (resp []datatypes.Resource_Group_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Group_Template) GetObject() (resp datatypes.Resource_Group_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Resource_Group_Template) GetChildren() (resp []datatypes.Resource_Group_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Group_Template) GetMembers() (resp []datatypes.Resource_Group_Template_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Group_Template) GetPackage() (resp datatypes.Product_Package, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Resource_Group_Template_Member struct {
	Session *Session
	Options
}

func (r *Session) GetResourceGroupTemplateMemberService() Resource_Group_Template_Member {
	return Resource_Group_Template_Member{Session: r}
}

func (r *Resource_Group_Template_Member) GetRole() (resp datatypes.Resource_Group_Role, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Group_Template_Member) GetTemplate() (resp datatypes.Resource_Group_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Resource_Metadata struct {
	Session *Session
	Options
}

func (r *Session) GetResourceMetadataService() Resource_Metadata {
	return Resource_Metadata{Session: r}
}

func (r *Resource_Metadata) GetBackendMacAddresses() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Metadata) GetDatacenter() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Metadata) GetDatacenterId() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Metadata) GetDomain() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Metadata) GetFrontendMacAddresses() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Metadata) GetFullyQualifiedDomainName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Metadata) GetHostname() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Metadata) GetId() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Metadata) GetPrimaryBackendIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Metadata) GetPrimaryIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Metadata) GetProvisionState() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Metadata) GetRouter(macAddress *string) (resp string, err error) {
	params := []interface{}{
		macAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Metadata) GetServiceResource(serviceName *string, index *int) (resp string, err error) {
	params := []interface{}{
		serviceName,
		index,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Metadata) GetServiceResources() (resp []datatypes.Container_Resource_Metadata_ServiceResource, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Metadata) GetTags() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Metadata) GetUserMetadata() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Metadata) GetVlanIds(macAddress *string) (resp []int, err error) {
	params := []interface{}{
		macAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Resource_Metadata) GetVlans(macAddress *string) (resp []int, err error) {
	params := []interface{}{
		macAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
