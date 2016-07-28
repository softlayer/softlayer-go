package softlayer

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Layout_Container interface {
	Entity

	GetAllObjects() (datatypes.Layout_Container, error)
	GetObject() (datatypes.Layout_Container, error)

	GetLayoutContainerType() (datatypes.Layout_Container_Type, error)
	GetLayoutItems() (datatypes.Layout_Item, error)
}

type Layout_Container_Type interface {
	Entity
}

type Layout_Item interface {
	Entity

	GetObject() (datatypes.Layout_Item, error)

	GetLayoutItemPreferences() (datatypes.Layout_Preference, error)
	GetLayoutItemType() (datatypes.Layout_Item_Type, error)
}

type Layout_Item_Type interface {
	Entity
}

type Layout_Preference interface {
	Entity

	GetLayoutPreferenceType() (datatypes.Layout_Preference_Type, error)
}

type Layout_Preference_Type interface {
	Entity
}

type Layout_Profile interface {
	Entity

	CreateObject(templateObject datatypes.Layout_Profile) (bool, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Layout_Profile) (bool, error)
	GetObject() (datatypes.Layout_Profile, error)
	ModifyPreference(templateObject datatypes.Layout_Profile_Preference) (datatypes.Layout_Profile_Preference, error)
	ModifyPreferences(layoutPreferenceObjects datatypes.Layout_Profile_Preference) (datatypes.Layout_Profile_Preference, error)

	GetLayoutContainers() (datatypes.Layout_Container, error)
	GetLayoutPreferences() (datatypes.Layout_Profile_Preference, error)
}

type Layout_Profile_Containers interface {
	Entity

	CreateObject(templateObject datatypes.Layout_Profile_Containers) (bool, error)
	EditObject(templateObject datatypes.Layout_Profile_Containers) (bool, error)
	GetObject() (datatypes.Layout_Profile_Containers, error)

	GetLayoutContainerType() (datatypes.Layout_Container, error)
	GetLayoutProfile() (datatypes.Layout_Profile, error)
}

type Layout_Profile_Customer interface {
	Layout_Profile

	GetObject() (datatypes.Layout_Profile_Customer, error)

	GetUserRecord() (datatypes.User_Customer, error)
}

type Layout_Profile_Preference interface {
	Entity

	GetObject() (datatypes.Layout_Profile_Preference, error)

	GetLayoutContainer() (datatypes.Layout_Container, error)
	GetLayoutItem() (datatypes.Layout_Item, error)
	GetLayoutPreference() (datatypes.Layout_Preference, error)
	GetLayoutProfile() (datatypes.Layout_Profile, error)
}
