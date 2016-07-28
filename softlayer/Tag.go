package softlayer

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Tag interface {
	Entity

	AutoComplete(tag string) (datatypes.Tag, error)
	GetAllTagTypes() (datatypes.Tag_Type, error)
	GetObject() (datatypes.Tag, error)
	GetTagByTagName(tagList string) (datatypes.Tag, error)
	SetTags(tags string, keyName string, resourceTableId int) (bool, error)

	GetAccount() (datatypes.Account, error)
	GetReferences() (datatypes.Tag_Reference, error)
}

type Tag_Reference interface {
	Entity

	GetCustomer() (datatypes.User_Customer, error)
	GetEmployee() (datatypes.User_Employee, error)
	GetTag() (datatypes.Tag, error)
	GetTagType() (datatypes.Tag_Type, error)
}

type Tag_Reference_Hardware interface {
	Tag_Reference

	GetResource() (datatypes.Hardware, error)
}

type Tag_Reference_Network_Application_Delivery_Controller interface {
	Tag_Reference

	GetResource() (datatypes.Network_Application_Delivery_Controller, error)
}

type Tag_Reference_Network_Vlan interface {
	Tag_Reference

	GetResource() (datatypes.Network_Vlan, error)
}

type Tag_Reference_Network_Vlan_Firewall interface {
	Tag_Reference

	GetResource() (datatypes.Network_Vlan_Firewall, error)
}

type Tag_Reference_Resource_Group interface {
	Tag_Reference

	GetResource() (datatypes.Resource_Group, error)
}

type Tag_Reference_Virtual_Guest interface {
	Tag_Reference

	GetResource() (datatypes.Virtual_Guest, error)
}

type Tag_Reference_Virtual_Guest_Block_Device_Template_Group interface {
	Tag_Reference

	GetResource() (datatypes.Virtual_Guest_Block_Device_Template_Group, error)
}

type Tag_Type interface {
	Entity
}
