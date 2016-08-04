package datatypes

type Tag struct {
	Entity

	Account        *Account        `json:"account,omitempty"`
	AccountId      *int            `json:"accountId,omitempty"`
	Id             *int            `json:"id,omitempty"`
	Internal       *int            `json:"internal,omitempty"`
	Name           *string         `json:"name,omitempty"`
	ReferenceCount *uint           `json:"referenceCount,omitempty"`
	References     []Tag_Reference `json:"references,omitempty"`
}

type Tag_Reference struct {
	Entity

	Customer        *User_Customer `json:"customer,omitempty"`
	EmpRecordId     *int           `json:"empRecordId,omitempty"`
	Employee        *User_Employee `json:"employee,omitempty"`
	Id              *int           `json:"id,omitempty"`
	ResourceTableId *int           `json:"resourceTableId,omitempty"`
	Tag             *Tag           `json:"tag,omitempty"`
	TagId           *int           `json:"tagId,omitempty"`
	TagType         *Tag_Type      `json:"tagType,omitempty"`
	TagTypeId       *int           `json:"tagTypeId,omitempty"`
	UsrRecordId     *int           `json:"usrRecordId,omitempty"`
}

type Tag_Reference_Hardware struct {
	Tag_Reference

	Resource *Hardware `json:"resource,omitempty"`
}

type Tag_Reference_Network_Application_Delivery_Controller struct {
	Tag_Reference

	Resource *Network_Application_Delivery_Controller `json:"resource,omitempty"`
}

type Tag_Reference_Network_Vlan struct {
	Tag_Reference

	Resource *Network_Vlan `json:"resource,omitempty"`
}

type Tag_Reference_Network_Vlan_Firewall struct {
	Tag_Reference

	Resource *Network_Vlan_Firewall `json:"resource,omitempty"`
}

type Tag_Reference_Resource_Group struct {
	Tag_Reference

	Resource *Resource_Group `json:"resource,omitempty"`
}

type Tag_Reference_Virtual_Guest struct {
	Tag_Reference

	Resource *Virtual_Guest `json:"resource,omitempty"`
}

type Tag_Reference_Virtual_Guest_Block_Device_Template_Group struct {
	Tag_Reference

	Resource *Virtual_Guest_Block_Device_Template_Group `json:"resource,omitempty"`
}

type Tag_Type struct {
	Entity

	Description *string `json:"description,omitempty"`
	KeyName     *string `json:"keyName,omitempty"`
}
