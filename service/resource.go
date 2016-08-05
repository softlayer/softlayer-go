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
