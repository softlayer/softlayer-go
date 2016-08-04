package datatypes

import "time"

type Layout_Container struct {
	Entity

	Id                    *int                   `json:"id,omitempty"`
	Keyname               *string                `json:"keyname,omitempty"`
	LayoutContainerType   *Layout_Container_Type `json:"layoutContainerType,omitempty"`
	LayoutContainerTypeId *int                   `json:"layoutContainerTypeId,omitempty"`
	LayoutItemCount       *uint                  `json:"layoutItemCount,omitempty"`
	LayoutItems           []Layout_Item          `json:"layoutItems,omitempty"`
	Name                  *string                `json:"name,omitempty"`
}

type Layout_Container_Type struct {
	Entity

	Id      *int    `json:"id,omitempty"`
	Keyname *string `json:"keyname,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type Layout_Item struct {
	Entity

	Id                        *int                `json:"id,omitempty"`
	Keyname                   *string             `json:"keyname,omitempty"`
	LayoutItemPreferenceCount *uint               `json:"layoutItemPreferenceCount,omitempty"`
	LayoutItemPreferences     []Layout_Preference `json:"layoutItemPreferences,omitempty"`
	LayoutItemType            *Layout_Item_Type   `json:"layoutItemType,omitempty"`
	LayoutItemTypeId          *int                `json:"layoutItemTypeId,omitempty"`
	Name                      *string             `json:"name,omitempty"`
}

type Layout_Item_Type struct {
	Entity

	Id      *int    `json:"id,omitempty"`
	Keyname *string `json:"keyname,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type Layout_Preference struct {
	Entity

	Id                     *int                    `json:"id,omitempty"`
	LayoutPreferenceType   *Layout_Preference_Type `json:"layoutPreferenceType,omitempty"`
	LayoutPreferenceTypeId *int                    `json:"layoutPreferenceTypeId,omitempty"`
	Value                  *string                 `json:"value,omitempty"`
}

type Layout_Preference_Type struct {
	Entity

	Id              *int    `json:"id,omitempty"`
	Keyname         *string `json:"keyname,omitempty"`
	Name            *string `json:"name,omitempty"`
	ValueExpression *string `json:"valueExpression,omitempty"`
}

type Layout_Profile struct {
	Entity

	ActiveFlag            *int                        `json:"activeFlag,omitempty"`
	CreateDate            *time.Time                  `json:"createDate,omitempty"`
	Id                    *int                        `json:"id,omitempty"`
	LayoutContainerCount  *uint                       `json:"layoutContainerCount,omitempty"`
	LayoutContainers      []Layout_Container          `json:"layoutContainers,omitempty"`
	LayoutPreferenceCount *uint                       `json:"layoutPreferenceCount,omitempty"`
	LayoutPreferences     []Layout_Profile_Preference `json:"layoutPreferences,omitempty"`
	ModifyDate            *time.Time                  `json:"modifyDate,omitempty"`
	Name                  *string                     `json:"name,omitempty"`
	UserRecordId          *int                        `json:"userRecordId,omitempty"`
}

type Layout_Profile_Containers struct {
	Entity

	CreateDate          *time.Time        `json:"createDate,omitempty"`
	Id                  *int              `json:"id,omitempty"`
	LayoutContainerId   *int              `json:"layoutContainerId,omitempty"`
	LayoutContainerType *Layout_Container `json:"layoutContainerType,omitempty"`
	LayoutProfile       *Layout_Profile   `json:"layoutProfile,omitempty"`
	LayoutProfileId     *int              `json:"layoutProfileId,omitempty"`
	ModifyDate          *time.Time        `json:"modifyDate,omitempty"`
}

type Layout_Profile_Customer struct {
	Layout_Profile

	UserRecord *User_Customer `json:"userRecord,omitempty"`
}

type Layout_Profile_Preference struct {
	Entity

	CreateDate         *time.Time         `json:"createDate,omitempty"`
	DefaultValueFlag   *int               `json:"defaultValueFlag,omitempty"`
	LayoutContainer    *Layout_Container  `json:"layoutContainer,omitempty"`
	LayoutContainerId  *int               `json:"layoutContainerId,omitempty"`
	LayoutItem         *Layout_Item       `json:"layoutItem,omitempty"`
	LayoutItemId       *int               `json:"layoutItemId,omitempty"`
	LayoutPreference   *Layout_Preference `json:"layoutPreference,omitempty"`
	LayoutPreferenceId *int               `json:"layoutPreferenceId,omitempty"`
	LayoutProfile      *Layout_Profile    `json:"layoutProfile,omitempty"`
	LayoutProfileId    *int               `json:"layoutProfileId,omitempty"`
	ModifyDate         *time.Time         `json:"modifyDate,omitempty"`
	Value              *string            `json:"value,omitempty"`
}
