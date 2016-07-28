package softlayer

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Configuration_Storage_Filesystem_Type interface {
	Entity
}

type Configuration_Storage_Group_Array_Type interface {
	Entity

	GetAllObjects() (datatypes.Configuration_Storage_Group_Array_Type, error)
	GetObject() (datatypes.Configuration_Storage_Group_Array_Type, error)

	GetHardwareComponentModels() (datatypes.Hardware_Component_Model, error)
}

type Configuration_Storage_Group_Order interface {
	Entity

	GetArrayType() (datatypes.Configuration_Storage_Group_Array_Type, error)
	GetBillingOrderItem() (datatypes.Billing_Order_Item, error)
}

type Configuration_Storage_Group_Template_Group interface {
	Entity

	GetType() (datatypes.Configuration_Storage_Group_Array_Type, error)
}

type Configuration_Template interface {
	Entity

	CopyTemplate(templateObject datatypes.Configuration_Template) (datatypes.Configuration_Template, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Configuration_Template) (bool, error)
	GetAllObjects() (datatypes.Configuration_Template, error)
	GetObject() (datatypes.Configuration_Template, error)
	UpdateDefaultValues(configurationValues datatypes.Configuration_Template_Section_Definition_Value) (bool, error)

	GetAccount() (datatypes.Account, error)
	GetConfigurationSections() (datatypes.Configuration_Template_Section, error)
	GetConfigurationTemplateReference() (datatypes.Monitoring_Agent_Configuration_Template_Group_Reference, error)
	GetDefaultValues() (datatypes.Configuration_Template_Section_Definition_Value, error)
	GetDefinitions() (datatypes.Configuration_Template_Section_Definition, error)
	GetItem() (datatypes.Product_Item, error)
	GetLinkedSectionReferences() (datatypes.Configuration_Template_Section_Reference, error)
	GetParent() (datatypes.Configuration_Template, error)
	GetUser() (datatypes.User_Customer, error)
}

type Configuration_Template_Attribute interface {
	Entity

	GetConfigurationTemplate() (datatypes.Configuration_Template, error)
}

type Configuration_Template_Section interface {
	Entity

	GetObject() (datatypes.Configuration_Template_Section, error)
	HasSubSections() (bool, error)

	GetDefinitions() (datatypes.Configuration_Template_Section_Definition, error)
	GetDisallowedDeletionFlag() (bool, error)
	GetLinkedTemplate() (datatypes.Configuration_Template, error)
	GetLinkedTemplateReference() (datatypes.Configuration_Template_Section_Reference, error)
	GetProfiles() (datatypes.Configuration_Template_Section_Profile, error)
	GetSectionType() (datatypes.Configuration_Template_Section_Type, error)
	GetSectionTypeName() (string, error)
	GetSubSections() (datatypes.Configuration_Template_Section, error)
	GetTemplate() (datatypes.Configuration_Template, error)
}

type Configuration_Template_Section_Attribute interface {
	Entity

	GetConfigurationSection() (datatypes.Configuration_Template_Section, error)
}

type Configuration_Template_Section_Definition interface {
	Entity

	GetObject() (datatypes.Configuration_Template_Section_Definition, error)

	GetAttributes() (datatypes.Configuration_Template_Section_Definition_Attribute, error)
	GetDefaultValue() (datatypes.Configuration_Template_Section_Definition_Value, error)
	GetGroup() (datatypes.Configuration_Template_Section_Definition_Group, error)
	GetMonitoringDataFlag() (bool, error)
	GetSection() (datatypes.Configuration_Template_Section, error)
	GetValueType() (datatypes.Configuration_Template_Section_Definition_Type, error)
}

type Configuration_Template_Section_Definition_Attribute interface {
	Entity

	GetAttributeType() (datatypes.Configuration_Template_Section_Definition_Attribute_Type, error)
	GetConfigurationDefinition() (datatypes.Configuration_Template_Section_Definition, error)
}

type Configuration_Template_Section_Definition_Attribute_Type interface {
	Entity
}

type Configuration_Template_Section_Definition_Group interface {
	Entity

	GetAllGroups() (datatypes.Configuration_Template_Section_Definition_Group, error)
	GetObject() (datatypes.Configuration_Template_Section_Definition_Group, error)

	GetParent() (datatypes.Configuration_Template_Section_Definition_Group, error)
}

type Configuration_Template_Section_Definition_Type interface {
	Entity

	GetObject() (datatypes.Configuration_Template_Section_Definition_Type, error)
}

type Configuration_Template_Section_Definition_Value interface {
	Entity

	GetObject() (datatypes.Configuration_Template_Section_Definition_Value, error)

	GetDefinition() (datatypes.Configuration_Template_Section_Definition, error)
	GetTemplate() (datatypes.Configuration_Template, error)
}

type Configuration_Template_Section_Profile interface {
	Entity

	GetObject() (datatypes.Configuration_Template_Section_Profile, error)

	GetConfigurationSection() (datatypes.Configuration_Template_Section, error)
	GetMonitoringAgent() (datatypes.Monitoring_Agent, error)
}

type Configuration_Template_Section_Reference interface {
	Entity

	GetObject() (datatypes.Configuration_Template_Section_Reference, error)

	GetSection() (datatypes.Configuration_Template_Section, error)
	GetTemplate() (datatypes.Configuration_Template, error)
}

type Configuration_Template_Section_Type interface {
	Entity

	GetObject() (datatypes.Configuration_Template_Section_Type, error)
}

type Configuration_Template_Type interface {
	Entity

	GetObject() (datatypes.Configuration_Template_Type, error)
}
