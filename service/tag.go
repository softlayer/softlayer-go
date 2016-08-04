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

type Tag struct {
	Session *Session
	Options
}

func (r *Session) GetTagService() Tag {
	return Tag{Session: r}
}

func (r *Tag) AutoComplete(tag *string) (resp []datatypes.Tag, err error) {
	params := []interface{}{
		tag,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Tag) GetAllTagTypes() (resp []datatypes.Tag_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Tag) GetObject() (resp datatypes.Tag, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Tag) GetTagByTagName(tagList *string) (resp []datatypes.Tag, err error) {
	params := []interface{}{
		tagList,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Tag) SetTags(tags *string, keyName *string, resourceTableId *int) (resp bool, err error) {
	params := []interface{}{
		tags,
		keyName,
		resourceTableId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

func (r *Tag) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Tag) GetReferences() (resp []datatypes.Tag_Reference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Tag_Reference struct {
	Session *Session
	Options
}

func (r *Session) GetTagReferenceService() Tag_Reference {
	return Tag_Reference{Session: r}
}

func (r *Tag_Reference) GetCustomer() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Tag_Reference) GetEmployee() (resp datatypes.User_Employee, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Tag_Reference) GetTag() (resp datatypes.Tag, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Tag_Reference) GetTagType() (resp datatypes.Tag_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Tag_Reference_Hardware struct {
	Session *Session
	Options
}

func (r *Session) GetTagReferenceHardwareService() Tag_Reference_Hardware {
	return Tag_Reference_Hardware{Session: r}
}

func (r *Tag_Reference_Hardware) GetResource() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Tag_Reference_Network_Application_Delivery_Controller struct {
	Session *Session
	Options
}

func (r *Session) GetTagReferenceNetworkApplicationDeliveryControllerService() Tag_Reference_Network_Application_Delivery_Controller {
	return Tag_Reference_Network_Application_Delivery_Controller{Session: r}
}

func (r *Tag_Reference_Network_Application_Delivery_Controller) GetResource() (resp datatypes.Network_Application_Delivery_Controller, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Tag_Reference_Network_Vlan struct {
	Session *Session
	Options
}

func (r *Session) GetTagReferenceNetworkVlanService() Tag_Reference_Network_Vlan {
	return Tag_Reference_Network_Vlan{Session: r}
}

func (r *Tag_Reference_Network_Vlan) GetResource() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Tag_Reference_Network_Vlan_Firewall struct {
	Session *Session
	Options
}

func (r *Session) GetTagReferenceNetworkVlanFirewallService() Tag_Reference_Network_Vlan_Firewall {
	return Tag_Reference_Network_Vlan_Firewall{Session: r}
}

func (r *Tag_Reference_Network_Vlan_Firewall) GetResource() (resp datatypes.Network_Vlan_Firewall, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Tag_Reference_Resource_Group struct {
	Session *Session
	Options
}

func (r *Session) GetTagReferenceResourceGroupService() Tag_Reference_Resource_Group {
	return Tag_Reference_Resource_Group{Session: r}
}

func (r *Tag_Reference_Resource_Group) GetResource() (resp datatypes.Resource_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Tag_Reference_Virtual_Guest struct {
	Session *Session
	Options
}

func (r *Session) GetTagReferenceVirtualGuestService() Tag_Reference_Virtual_Guest {
	return Tag_Reference_Virtual_Guest{Session: r}
}

func (r *Tag_Reference_Virtual_Guest) GetResource() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Tag_Reference_Virtual_Guest_Block_Device_Template_Group struct {
	Session *Session
	Options
}

func (r *Session) GetTagReferenceVirtualGuestBlockDeviceTemplateGroupService() Tag_Reference_Virtual_Guest_Block_Device_Template_Group {
	return Tag_Reference_Virtual_Guest_Block_Device_Template_Group{Session: r}
}

func (r *Tag_Reference_Virtual_Guest_Block_Device_Template_Group) GetResource() (resp datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Tag_Type struct {
	Session *Session
	Options
}

func (r *Session) GetTagTypeService() Tag_Type {
	return Tag_Type{Session: r}
}
