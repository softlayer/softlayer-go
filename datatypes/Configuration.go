package datatypes

import "time"

type Configuration_Storage_Filesystem_Type struct {
	Entity

	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Configuration_Storage_Group_Array_Type struct {
	Entity

	Description                 *string                    `json:"description:omitempty"`
	DriveMultiplier             *int                       `json:"driveMultiplier:omitempty"`
	HardwareComponentModelCount *uint                      `json:"hardwareComponentModelCount:omitempty"`
	HardwareComponentModels     []Hardware_Component_Model `json:"hardwareComponentModels:omitempty"`
	HotspareAllow               *bool                      `json:"hotspareAllow:omitempty"`
	Id                          *int                       `json:"id:omitempty"`
	KeyName                     *string                    `json:"keyName:omitempty"`
	MaximumDrives               *int                       `json:"maximumDrives:omitempty"`
	MinimumDrives               *int                       `json:"minimumDrives:omitempty"`
	Name                        *string                    `json:"name:omitempty"`
}

type Configuration_Storage_Group_Order struct {
	Entity

	ArrayNumber        *int                                    `json:"arrayNumber:omitempty"`
	ArraySize          *float64                                `json:"arraySize:omitempty"`
	ArrayType          *Configuration_Storage_Group_Array_Type `json:"arrayType:omitempty"`
	ArrayTypeId        *int                                    `json:"arrayTypeId:omitempty"`
	BillingOrderItem   *Billing_Order_Item                     `json:"billingOrderItem:omitempty"`
	BillingOrderItemId *int                                    `json:"billingOrderItemId:omitempty"`
	HardDrives         []int                                   `json:"hardDrives:omitempty"`
	HotSpareDrives     []int                                   `json:"hotSpareDrives:omitempty"`
	PartitionData      *string                                 `json:"partitionData:omitempty"`
}

type Configuration_Storage_Group_Template_Group struct {
	Entity

	Grow             *bool                                   `json:"grow:omitempty"`
	HardDrivesString *string                                 `json:"hardDrivesString:omitempty"`
	OrderIndex       *int                                    `json:"orderIndex:omitempty"`
	Size             *float64                                `json:"size:omitempty"`
	Type             *Configuration_Storage_Group_Array_Type `json:"type:omitempty"`
}

type Configuration_Template struct {
	Entity

	Account                             *Account                                                  `json:"account:omitempty"`
	AccountId                           *int                                                      `json:"accountId:omitempty"`
	ConfigurationSectionCount           *uint                                                     `json:"configurationSectionCount:omitempty"`
	ConfigurationSections               []Configuration_Template_Section                          `json:"configurationSections:omitempty"`
	ConfigurationTemplateReference      []Monitoring_Agent_Configuration_Template_Group_Reference `json:"configurationTemplateReference:omitempty"`
	ConfigurationTemplateReferenceCount *uint                                                     `json:"configurationTemplateReferenceCount:omitempty"`
	CreateDate                          *time.Time                                                `json:"createDate:omitempty"`
	DefaultValueCount                   *uint                                                     `json:"defaultValueCount:omitempty"`
	DefaultValues                       []Configuration_Template_Section_Definition_Value         `json:"defaultValues:omitempty"`
	DefinitionCount                     *uint                                                     `json:"definitionCount:omitempty"`
	Definitions                         []Configuration_Template_Section_Definition               `json:"definitions:omitempty"`
	Description                         *string                                                   `json:"description:omitempty"`
	Id                                  *int                                                      `json:"id:omitempty"`
	Item                                *Product_Item                                             `json:"item:omitempty"`
	ItemId                              *int                                                      `json:"itemId:omitempty"`
	LinkedSectionReferences             *Configuration_Template_Section_Reference                 `json:"linkedSectionReferences:omitempty"`
	ModifyDate                          *time.Time                                                `json:"modifyDate:omitempty"`
	Name                                *string                                                   `json:"name:omitempty"`
	Parent                              *Configuration_Template                                   `json:"parent:omitempty"`
	ParentId                            *int                                                      `json:"parentId:omitempty"`
	User                                *User_Customer                                            `json:"user:omitempty"`
	UserRecordId                        *int                                                      `json:"userRecordId:omitempty"`
}

type Configuration_Template_Attribute struct {
	Entity

	ConfigurationTemplate *Configuration_Template `json:"configurationTemplate:omitempty"`
	Value                 *string                 `json:"value:omitempty"`
}

type Configuration_Template_Section struct {
	Entity

	CreateDate              *time.Time                                  `json:"createDate:omitempty"`
	DefinitionCount         *uint                                       `json:"definitionCount:omitempty"`
	Definitions             []Configuration_Template_Section_Definition `json:"definitions:omitempty"`
	Description             *string                                     `json:"description:omitempty"`
	DisallowedDeletionFlag  *bool                                       `json:"disallowedDeletionFlag:omitempty"`
	Id                      *int                                        `json:"id:omitempty"`
	LinkedTemplate          *Configuration_Template                     `json:"linkedTemplate:omitempty"`
	LinkedTemplateId        *string                                     `json:"linkedTemplateId:omitempty"`
	LinkedTemplateReference *Configuration_Template_Section_Reference   `json:"linkedTemplateReference:omitempty"`
	ModifyDate              *time.Time                                  `json:"modifyDate:omitempty"`
	Name                    *string                                     `json:"name:omitempty"`
	ParentId                *int                                        `json:"parentId:omitempty"`
	ProfileCount            *uint                                       `json:"profileCount:omitempty"`
	Profiles                []Configuration_Template_Section_Profile    `json:"profiles:omitempty"`
	SectionType             *Configuration_Template_Section_Type        `json:"sectionType:omitempty"`
	SectionTypeName         *string                                     `json:"sectionTypeName:omitempty"`
	Sort                    *int                                        `json:"sort:omitempty"`
	SubSectionCount         *uint                                       `json:"subSectionCount:omitempty"`
	SubSections             []Configuration_Template_Section            `json:"subSections:omitempty"`
	Template                *Configuration_Template                     `json:"template:omitempty"`
	TemplateId              *string                                     `json:"templateId:omitempty"`
	TypeId                  *int                                        `json:"typeId:omitempty"`
}

type Configuration_Template_Section_Attribute struct {
	Entity

	ConfigurationSection *Configuration_Template_Section `json:"configurationSection:omitempty"`
	Value                *string                         `json:"value:omitempty"`
}

type Configuration_Template_Section_Definition struct {
	Entity

	AttributeCount     *uint                                                 `json:"attributeCount:omitempty"`
	Attributes         []Configuration_Template_Section_Definition_Attribute `json:"attributes:omitempty"`
	CreateDate         *time.Time                                            `json:"createDate:omitempty"`
	DefaultValue       *Configuration_Template_Section_Definition_Value      `json:"defaultValue:omitempty"`
	Description        *string                                               `json:"description:omitempty"`
	EnumerationValues  *string                                               `json:"enumerationValues:omitempty"`
	Group              *Configuration_Template_Section_Definition_Group      `json:"group:omitempty"`
	GroupId            *string                                               `json:"groupId:omitempty"`
	Id                 *int                                                  `json:"id:omitempty"`
	MaximumValue       *string                                               `json:"maximumValue:omitempty"`
	MinimumValue       *string                                               `json:"minimumValue:omitempty"`
	ModifyDate         *time.Time                                            `json:"modifyDate:omitempty"`
	MonitoringDataFlag *bool                                                 `json:"monitoringDataFlag:omitempty"`
	Name               *string                                               `json:"name:omitempty"`
	Path               *string                                               `json:"path:omitempty"`
	RequireValueFlag   *int                                                  `json:"requireValueFlag:omitempty"`
	Section            *Configuration_Template_Section                       `json:"section:omitempty"`
	SectionId          *int                                                  `json:"sectionId:omitempty"`
	ShortName          *string                                               `json:"shortName:omitempty"`
	Sort               *int                                                  `json:"sort:omitempty"`
	TypeId             *int                                                  `json:"typeId:omitempty"`
	ValueType          *Configuration_Template_Section_Definition_Type       `json:"valueType:omitempty"`
}

type Configuration_Template_Section_Definition_Attribute struct {
	Entity

	AttributeType           *Configuration_Template_Section_Definition_Attribute_Type `json:"attributeType:omitempty"`
	ConfigurationDefinition *Configuration_Template_Section_Definition                `json:"configurationDefinition:omitempty"`
	Value                   *string                                                   `json:"value:omitempty"`
}

type Configuration_Template_Section_Definition_Attribute_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Configuration_Template_Section_Definition_Group struct {
	Entity

	CreateDate  *time.Time                                       `json:"createDate:omitempty"`
	Description *string                                          `json:"description:omitempty"`
	Id          *int                                             `json:"id:omitempty"`
	Name        *string                                          `json:"name:omitempty"`
	Parent      *Configuration_Template_Section_Definition_Group `json:"parent:omitempty"`
	SortOrder   *int                                             `json:"sortOrder:omitempty"`
}

type Configuration_Template_Section_Definition_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Configuration_Template_Section_Definition_Value struct {
	Entity

	CreateDate   *time.Time                                 `json:"createDate:omitempty"`
	Definition   *Configuration_Template_Section_Definition `json:"definition:omitempty"`
	DefinitionId *int                                       `json:"definitionId:omitempty"`
	ModifyDate   *time.Time                                 `json:"modifyDate:omitempty"`
	Template     *Configuration_Template                    `json:"template:omitempty"`
	TemplateId   *int                                       `json:"templateId:omitempty"`
	Value        *string                                    `json:"value:omitempty"`
}

type Configuration_Template_Section_Profile struct {
	Entity

	AgentId              *int                            `json:"agentId:omitempty"`
	ConfigurationSection *Configuration_Template_Section `json:"configurationSection:omitempty"`
	CreateDate           *time.Time                      `json:"createDate:omitempty"`
	Id                   *int                            `json:"id:omitempty"`
	MonitoringAgent      *Monitoring_Agent               `json:"monitoringAgent:omitempty"`
	Name                 *string                         `json:"name:omitempty"`
	SectionId            *int                            `json:"sectionId:omitempty"`
}

type Configuration_Template_Section_Reference struct {
	Entity

	CreateDate *time.Time                      `json:"createDate:omitempty"`
	Id         *int                            `json:"id:omitempty"`
	ModifyDate *time.Time                      `json:"modifyDate:omitempty"`
	Section    *Configuration_Template_Section `json:"section:omitempty"`
	SectionId  *int                            `json:"sectionId:omitempty"`
	Template   *Configuration_Template         `json:"template:omitempty"`
	TemplateId *int                            `json:"templateId:omitempty"`
}

type Configuration_Template_Section_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Configuration_Template_Type struct {
	Entity

	CreateDate  *time.Time `json:"createDate:omitempty"`
	Description *string    `json:"description:omitempty"`
	Id          *int       `json:"id:omitempty"`
	Name        *string    `json:"name:omitempty"`
}
