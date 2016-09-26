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
type Configuration_Storage_Filesystem_Type struct {
	Entity

	// no documentation yet
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// Supported hardware raid modes
type Configuration_Storage_Group_Array_Type struct {
	Entity

	// no documentation yet
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// no documentation yet
	DriveMultiplier *int `json:"driveMultiplier,omitempty" xmlrpc:"driveMultiplier"`

	// A count of
	HardwareComponentModelCount *uint `json:"hardwareComponentModelCount,omitempty" xmlrpc:"hardwareComponentModelCount"`

	// no documentation yet
	HardwareComponentModels []Hardware_Component_Model `json:"hardwareComponentModels,omitempty" xmlrpc:"hardwareComponentModels"`

	// no documentation yet
	HotspareAllow *bool `json:"hotspareAllow,omitempty" xmlrpc:"hotspareAllow"`

	// no documentation yet
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// no documentation yet
	MaximumDrives *int `json:"maximumDrives,omitempty" xmlrpc:"maximumDrives"`

	// no documentation yet
	MinimumDrives *int `json:"minimumDrives,omitempty" xmlrpc:"minimumDrives"`

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// Single storage group(array) used for a hardware server order.
//
// If a raid configuration is required this object will describe a single array that will be configured on the server. If the server requires more than one array, a storage group will need to be created for each array.
type Configuration_Storage_Group_Order struct {
	Entity

	// no documentation yet
	ArrayNumber *int `json:"arrayNumber,omitempty" xmlrpc:"arrayNumber"`

	// no documentation yet
	ArraySize *Float64 `json:"arraySize,omitempty" xmlrpc:"arraySize"`

	// Raid mode for the storage group.
	ArrayType *Configuration_Storage_Group_Array_Type `json:"arrayType,omitempty" xmlrpc:"arrayType"`

	// no documentation yet
	ArrayTypeId *int `json:"arrayTypeId,omitempty" xmlrpc:"arrayTypeId"`

	// The order item that relates to this storage group.
	BillingOrderItem *Billing_Order_Item `json:"billingOrderItem,omitempty" xmlrpc:"billingOrderItem"`

	// no documentation yet
	BillingOrderItemId *int `json:"billingOrderItemId,omitempty" xmlrpc:"billingOrderItemId"`

	// no documentation yet
	HardDrives []int `json:"hardDrives,omitempty" xmlrpc:"hardDrives"`

	// no documentation yet
	HotSpareDrives []int `json:"hotSpareDrives,omitempty" xmlrpc:"hotSpareDrives"`

	// no documentation yet
	PartitionData *string `json:"partitionData,omitempty" xmlrpc:"partitionData"`
}

// Single storage group(array) used in a storage group template.
//
// If a server configuration requires a raid configuration this object will describe a single array to be configured.
type Configuration_Storage_Group_Template_Group struct {
	Entity

	// Flag to use all available space.
	Grow *bool `json:"grow,omitempty" xmlrpc:"grow"`

	// Comma delimited integers of drive indexes for the array. This can also be the string 'all' to specify all drives in the server
	HardDrivesString *string `json:"hardDrivesString,omitempty" xmlrpc:"hardDrivesString"`

	// The order of the arrays in the template.
	OrderIndex *int `json:"orderIndex,omitempty" xmlrpc:"orderIndex"`

	// Size of array. Must be within limitations of the smallest drive and raid mode
	Size *Float64 `json:"size,omitempty" xmlrpc:"size"`

	// no documentation yet
	Type *Configuration_Storage_Group_Array_Type `json:"type,omitempty" xmlrpc:"type"`
}

// The SoftLayer_Configuration_Template data type contains general information of an arbitrary resource.
type Configuration_Template struct {
	Entity

	// no documentation yet
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// Internal identifier of a SoftLayer account that this configuration template belongs to
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// A count of
	ConfigurationSectionCount *uint `json:"configurationSectionCount,omitempty" xmlrpc:"configurationSectionCount"`

	// no documentation yet
	ConfigurationSections []Configuration_Template_Section `json:"configurationSections,omitempty" xmlrpc:"configurationSections"`

	// no documentation yet
	ConfigurationTemplateReference []Monitoring_Agent_Configuration_Template_Group_Reference `json:"configurationTemplateReference,omitempty" xmlrpc:"configurationTemplateReference"`

	// A count of
	ConfigurationTemplateReferenceCount *uint `json:"configurationTemplateReferenceCount,omitempty" xmlrpc:"configurationTemplateReferenceCount"`

	// Created date
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// A count of
	DefaultValueCount *uint `json:"defaultValueCount,omitempty" xmlrpc:"defaultValueCount"`

	// no documentation yet
	DefaultValues []Configuration_Template_Section_Definition_Value `json:"defaultValues,omitempty" xmlrpc:"defaultValues"`

	// A count of
	DefinitionCount *uint `json:"definitionCount,omitempty" xmlrpc:"definitionCount"`

	// no documentation yet
	Definitions []Configuration_Template_Section_Definition `json:"definitions,omitempty" xmlrpc:"definitions"`

	// Configuration template description
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// Internal identifier of a configuration template.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	Item *Product_Item `json:"item,omitempty" xmlrpc:"item"`

	// Internal identifier of a product item that this configuration template is associated with
	ItemId *int `json:"itemId,omitempty" xmlrpc:"itemId"`

	// no documentation yet
	LinkedSectionReferences *Configuration_Template_Section_Reference `json:"linkedSectionReferences,omitempty" xmlrpc:"linkedSectionReferences"`

	// Last modified date
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// Configuration template name
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// no documentation yet
	Parent *Configuration_Template `json:"parent,omitempty" xmlrpc:"parent"`

	// Internal identifier of the parent configuration template
	ParentId *int `json:"parentId,omitempty" xmlrpc:"parentId"`

	// no documentation yet
	User *User_Customer `json:"user,omitempty" xmlrpc:"user"`

	// Internal identifier of a user that last modified this configuration template
	UserRecordId *int `json:"userRecordId,omitempty" xmlrpc:"userRecordId"`
}

// Configuration template attribute class contains supplementary information for a configuration template.
type Configuration_Template_Attribute struct {
	Entity

	// no documentation yet
	ConfigurationTemplate *Configuration_Template `json:"configurationTemplate,omitempty" xmlrpc:"configurationTemplate"`

	// Value of a configuration template attribute
	Value *string `json:"value,omitempty" xmlrpc:"value"`
}

// The SoftLayer_Configuration_Template_Section data type contains information of a configuration section.
//
// Configuration can contain sub-sections.
type Configuration_Template_Section struct {
	Entity

	// Created date
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// A count of
	DefinitionCount *uint `json:"definitionCount,omitempty" xmlrpc:"definitionCount"`

	// no documentation yet
	Definitions []Configuration_Template_Section_Definition `json:"definitions,omitempty" xmlrpc:"definitions"`

	// Configuration section description
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// no documentation yet
	DisallowedDeletionFlag *bool `json:"disallowedDeletionFlag,omitempty" xmlrpc:"disallowedDeletionFlag"`

	// Internal identifier of a configuration section.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	LinkedTemplate *Configuration_Template `json:"linkedTemplate,omitempty" xmlrpc:"linkedTemplate"`

	// Internal identifier of a sub configuration template that this section points to. Use this property if you wish to create a reference to a sub configuration template when creating a linked section.
	LinkedTemplateId *string `json:"linkedTemplateId,omitempty" xmlrpc:"linkedTemplateId"`

	// no documentation yet
	LinkedTemplateReference *Configuration_Template_Section_Reference `json:"linkedTemplateReference,omitempty" xmlrpc:"linkedTemplateReference"`

	// Last modified date
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// Configuration section name
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// Internal identifier of the parent configuration section
	ParentId *int `json:"parentId,omitempty" xmlrpc:"parentId"`

	// A count of
	ProfileCount *uint `json:"profileCount,omitempty" xmlrpc:"profileCount"`

	// no documentation yet
	Profiles []Configuration_Template_Section_Profile `json:"profiles,omitempty" xmlrpc:"profiles"`

	// no documentation yet
	SectionType *Configuration_Template_Section_Type `json:"sectionType,omitempty" xmlrpc:"sectionType"`

	// no documentation yet
	SectionTypeName *string `json:"sectionTypeName,omitempty" xmlrpc:"sectionTypeName"`

	// Sort order
	Sort *int `json:"sort,omitempty" xmlrpc:"sort"`

	// A count of
	SubSectionCount *uint `json:"subSectionCount,omitempty" xmlrpc:"subSectionCount"`

	// no documentation yet
	SubSections []Configuration_Template_Section `json:"subSections,omitempty" xmlrpc:"subSections"`

	// no documentation yet
	Template *Configuration_Template `json:"template,omitempty" xmlrpc:"template"`

	// Internal identifier of a configuration template that this section belongs to
	TemplateId *string `json:"templateId,omitempty" xmlrpc:"templateId"`

	// Internal identifier of the configuration section type
	TypeId *int `json:"typeId,omitempty" xmlrpc:"typeId"`
}

// Configuration section attribute class contains supplementary information for a configuration section.
type Configuration_Template_Section_Attribute struct {
	Entity

	// no documentation yet
	ConfigurationSection *Configuration_Template_Section `json:"configurationSection,omitempty" xmlrpc:"configurationSection"`

	// Value of a configuration section attribute
	Value *string `json:"value,omitempty" xmlrpc:"value"`
}

// Configuration definition gives you details of the value that you're setting.
//
// Some monitoring agents requires values unique to your system. If value type is defined as "Resource Specific Values", you will have to make an additional API call to retrieve your system specific values.
//
// See [[SoftLayer_Monitoring_Agent::getAvailableConfigurationValues|Monitoring Agent]] service to retrieve your system specific values.
type Configuration_Template_Section_Definition struct {
	Entity

	// A count of
	AttributeCount *uint `json:"attributeCount,omitempty" xmlrpc:"attributeCount"`

	// no documentation yet
	Attributes []Configuration_Template_Section_Definition_Attribute `json:"attributes,omitempty" xmlrpc:"attributes"`

	// Created date
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// no documentation yet
	DefaultValue *Configuration_Template_Section_Definition_Value `json:"defaultValue,omitempty" xmlrpc:"defaultValue"`

	// Description of a configuration definition.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// Enumeration values separated by comma.
	EnumerationValues *string `json:"enumerationValues,omitempty" xmlrpc:"enumerationValues"`

	// no documentation yet
	Group *Configuration_Template_Section_Definition_Group `json:"group,omitempty" xmlrpc:"group"`

	// Definition group id.
	GroupId *string `json:"groupId,omitempty" xmlrpc:"groupId"`

	// Internal identifier of a configuration definition.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// Maximum value of a configuration definition.
	MaximumValue *string `json:"maximumValue,omitempty" xmlrpc:"maximumValue"`

	// Minimum value of a configuration definition.
	MinimumValue *string `json:"minimumValue,omitempty" xmlrpc:"minimumValue"`

	// Last modify date
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// no documentation yet
	MonitoringDataFlag *bool `json:"monitoringDataFlag,omitempty" xmlrpc:"monitoringDataFlag"`

	// Configuration definition name.
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// Definition path.
	Path *string `json:"path,omitempty" xmlrpc:"path"`

	// Indicates if a configuration value is required for this definition.
	RequireValueFlag *int `json:"requireValueFlag,omitempty" xmlrpc:"requireValueFlag"`

	// no documentation yet
	Section *Configuration_Template_Section `json:"section,omitempty" xmlrpc:"section"`

	// Internal identifier of a configuration section.
	SectionId *int `json:"sectionId,omitempty" xmlrpc:"sectionId"`

	// Shortened configuration definition name.
	ShortName *string `json:"shortName,omitempty" xmlrpc:"shortName"`

	// Sort order
	Sort *int `json:"sort,omitempty" xmlrpc:"sort"`

	// Internal identifier of a configuration definition type.
	TypeId *int `json:"typeId,omitempty" xmlrpc:"typeId"`

	// no documentation yet
	ValueType *Configuration_Template_Section_Definition_Type `json:"valueType,omitempty" xmlrpc:"valueType"`
}

// Configuration definition attribute class contains supplementary information for a configuration definition.
type Configuration_Template_Section_Definition_Attribute struct {
	Entity

	// no documentation yet
	AttributeType *Configuration_Template_Section_Definition_Attribute_Type `json:"attributeType,omitempty" xmlrpc:"attributeType"`

	// no documentation yet
	ConfigurationDefinition *Configuration_Template_Section_Definition `json:"configurationDefinition,omitempty" xmlrpc:"configurationDefinition"`

	// Value of a configuration definition attribute
	Value *string `json:"value,omitempty" xmlrpc:"value"`
}

// SoftLayer_Configuration_Template_Attribute_Type models the type of attribute that can be assigned to a configuration definition.
type Configuration_Template_Section_Definition_Attribute_Type struct {
	Entity

	// Description of a definition attribute type
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// Name of a definition attribute type
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// Configuration definition group gives you details of the definition and allows extra functionality.
//
//
type Configuration_Template_Section_Definition_Group struct {
	Entity

	// Created date
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// Internal Description of a definition group.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// Internal identifier of a definition group.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// Internal Definition group name.
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// no documentation yet
	Parent *Configuration_Template_Section_Definition_Group `json:"parent,omitempty" xmlrpc:"parent"`

	// Sort order
	SortOrder *int `json:"sortOrder,omitempty" xmlrpc:"sortOrder"`
}

// SoftLayer_Configuration_Template_Section_Definition_Type further defines the value of a configuration definition.
type Configuration_Template_Section_Definition_Type struct {
	Entity

	// Description of a configuration value type
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// Internal identifier of a configuration value type
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// Name of a configuration value type
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// SoftLayer_Configuration_Section_Value is used to set the value for a configuration definition
type Configuration_Template_Section_Definition_Value struct {
	Entity

	// Created date
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// no documentation yet
	Definition *Configuration_Template_Section_Definition `json:"definition,omitempty" xmlrpc:"definition"`

	// Internal identifier of a configuration definition that this configuration value if defined by
	DefinitionId *int `json:"definitionId,omitempty" xmlrpc:"definitionId"`

	// Internal Last modified date
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// no documentation yet
	Template *Configuration_Template `json:"template,omitempty" xmlrpc:"template"`

	// Internal identifier of a configuration template that this configuration value belongs to
	TemplateId *int `json:"templateId,omitempty" xmlrpc:"templateId"`

	// Internal Configuration value
	Value *string `json:"value,omitempty" xmlrpc:"value"`
}

// Some configuration templates let you create a unique configuration profiles.
//
// For example, you can create multiple configuration profiles to monitor multiple hard drives with "CPU/Memory/Disk Monitoring Agent". SoftLayer_Configuration_Template_Section_Profile help you keep track of custom configuration profiles.
type Configuration_Template_Section_Profile struct {
	Entity

	// Internal identifier of a monitoring agent this profile belongs to.
	AgentId *int `json:"agentId,omitempty" xmlrpc:"agentId"`

	// no documentation yet
	ConfigurationSection *Configuration_Template_Section `json:"configurationSection,omitempty" xmlrpc:"configurationSection"`

	// Created date
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// Internal identifier of a configuration profile.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	MonitoringAgent *Monitoring_Agent `json:"monitoringAgent,omitempty" xmlrpc:"monitoringAgent"`

	// Name of a configuration profile
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// Internal identifier of a configuration section that this profile belongs to.
	SectionId *int `json:"sectionId,omitempty" xmlrpc:"sectionId"`
}

// The SoftLayer_Configuration_Template_Section_Reference data type contains information of a configuration section and its associated configuration template.
type Configuration_Template_Section_Reference struct {
	Entity

	// Created date
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// Internal identifier of a configuration section reference.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// Modified date
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// no documentation yet
	Section *Configuration_Template_Section `json:"section,omitempty" xmlrpc:"section"`

	// Internal identifier of a configuration section.
	SectionId *int `json:"sectionId,omitempty" xmlrpc:"sectionId"`

	// no documentation yet
	Template *Configuration_Template `json:"template,omitempty" xmlrpc:"template"`

	// Internal identifier of a configuration template.
	TemplateId *int `json:"templateId,omitempty" xmlrpc:"templateId"`
}

// The SoftLayer_Configuration_Template_Section_Type data type contains information of a configuration section type.
//
// Configuration can contain sub-sections.
type Configuration_Template_Section_Type struct {
	Entity

	// Configuration section type description
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// Internal identifier of a configuration section type
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// Configuration section type name
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// The SoftLayer_Configuration_Template_Type data type contains configuration template type information.
type Configuration_Template_Type struct {
	Entity

	// Created date. This is deprecated now.
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// Description of a configuration template
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// Internal identifier of a configuration template type
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// Name of a configuration template type
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}
