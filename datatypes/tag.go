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

// The SoftLayer_Tag data type is an optional type associated with hardware. The account ID that the tag is tied to, and the tag itself are stored in this data type. There is also a flag to denote whether the tag is internal or not.
type Tag struct {
	Entity

	// The account to which the tag is tied.
	Account *Account `json:"account,omitempty"`

	// Account the tag belongs to.
	AccountId *int `json:"accountId,omitempty"`

	// Unique identifier for a tag.
	Id *int `json:"id,omitempty"`

	// Indicates whether a tag is internal.
	Internal *int `json:"internal,omitempty"`

	// Name of the tag. The characters permitted are A-Z, 0-9, whitespace, _ (underscore),
	Name *string `json:"name,omitempty"`

	// A count of references that tie object to the tag.
	ReferenceCount *uint `json:"referenceCount,omitempty"`

	// References that tie object to the tag.
	References []Tag_Reference `json:"references,omitempty"`
}

//
type Tag_Reference struct {
	Entity

	//
	Customer *User_Customer `json:"customer,omitempty"`

	//
	EmpRecordId *int `json:"empRecordId,omitempty"`

	//
	Employee *User_Employee `json:"employee,omitempty"`

	//
	Id *int `json:"id,omitempty"`

	//
	ResourceTableId *int `json:"resourceTableId,omitempty"`

	//
	Tag *Tag `json:"tag,omitempty"`

	//
	TagId *int `json:"tagId,omitempty"`

	//
	TagType *Tag_Type `json:"tagType,omitempty"`

	//
	TagTypeId *int `json:"tagTypeId,omitempty"`

	//
	UsrRecordId *int `json:"usrRecordId,omitempty"`
}

//
type Tag_Reference_Hardware struct {
	Tag_Reference

	//
	Resource *Hardware `json:"resource,omitempty"`
}

//
type Tag_Reference_Network_Application_Delivery_Controller struct {
	Tag_Reference

	//
	Resource *Network_Application_Delivery_Controller `json:"resource,omitempty"`
}

//
type Tag_Reference_Network_Vlan struct {
	Tag_Reference

	//
	Resource *Network_Vlan `json:"resource,omitempty"`
}

//
type Tag_Reference_Network_Vlan_Firewall struct {
	Tag_Reference

	//
	Resource *Network_Vlan_Firewall `json:"resource,omitempty"`
}

//
type Tag_Reference_Resource_Group struct {
	Tag_Reference

	//
	Resource *Resource_Group `json:"resource,omitempty"`
}

//
type Tag_Reference_Virtual_Guest struct {
	Tag_Reference

	//
	Resource *Virtual_Guest `json:"resource,omitempty"`
}

//
type Tag_Reference_Virtual_Guest_Block_Device_Template_Group struct {
	Tag_Reference

	//
	Resource *Virtual_Guest_Block_Device_Template_Group `json:"resource,omitempty"`
}

//
type Tag_Type struct {
	Entity

	//
	Description *string `json:"description,omitempty"`

	//
	KeyName *string `json:"keyName,omitempty"`
}
