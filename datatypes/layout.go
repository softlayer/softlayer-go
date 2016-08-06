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

// The SoftLayer_Layout_Container contains definitions for default page layouts
type Layout_Container struct {
	Entity

	// The internal identifier of a layout container
	Id *int `json:"id,omitempty"`

	// The unique key name of the layout container, used primarily for programmatic purposes
	Keyname *string `json:"keyname,omitempty"`

	// The type of the layout container object
	LayoutContainerType *Layout_Container_Type `json:"layoutContainerType,omitempty"`

	// The internal identifier of the related [[SoftLayer_Layout_Container_Type]]
	LayoutContainerTypeId *int `json:"layoutContainerTypeId,omitempty"`

	// A count of the layout items assigned to this layout container
	LayoutItemCount *uint `json:"layoutItemCount,omitempty"`

	// The layout items assigned to this layout container
	LayoutItems []Layout_Item `json:"layoutItems,omitempty"`

	// The friendly name of the layout container
	Name *string `json:"name,omitempty"`
}

// The SoftLayer_Layout_Container_Type contains definitions for container types
type Layout_Container_Type struct {
	Entity

	// The internal identifier of the container type
	Id *int `json:"id,omitempty"`

	// The unique key name of the container type, used primarily for programmatic purposes
	Keyname *string `json:"keyname,omitempty"`

	// The friendly name of the container type
	Name *string `json:"name,omitempty"`
}

// The SoftLayer_Layout_Item contains definitions for default layout items
type Layout_Item struct {
	Entity

	// The internal identifier of a layout item
	Id *int `json:"id,omitempty"`

	// The unique key name of the layout item, used primarily for programmatic purposes
	Keyname *string `json:"keyname,omitempty"`

	// A count of the layout preferences assigned to this layout item
	LayoutItemPreferenceCount *uint `json:"layoutItemPreferenceCount,omitempty"`

	// The layout preferences assigned to this layout item
	LayoutItemPreferences []Layout_Preference `json:"layoutItemPreferences,omitempty"`

	// The type of the layout item object
	LayoutItemType *Layout_Item_Type `json:"layoutItemType,omitempty"`

	// The internal identifier of the related [[SoftLayer_Layout_Item_Type]]
	LayoutItemTypeId *int `json:"layoutItemTypeId,omitempty"`

	// The friendly name of the layout item
	Name *string `json:"name,omitempty"`
}

// The SoftLayer_Layout_Item_Type contains definitions for item types
type Layout_Item_Type struct {
	Entity

	// The internal identifier of the item type
	Id *int `json:"id,omitempty"`

	// The unique key name of the item type, used primarily for programmatic purposes
	Keyname *string `json:"keyname,omitempty"`

	// The friendly name of the item type
	Name *string `json:"name,omitempty"`
}

// The SoftLayer_Layout_Preference contains definitions for default layout item preferences
type Layout_Preference struct {
	Entity

	// The internal identifier of a layout preference
	Id *int `json:"id,omitempty"`

	// The type of the preference object
	LayoutPreferenceType *Layout_Preference_Type `json:"layoutPreferenceType,omitempty"`

	// The internal identifier of the related [[SoftLayer_Layout_Preference_Type]]
	LayoutPreferenceTypeId *int `json:"layoutPreferenceTypeId,omitempty"`

	// The default value of the preference
	Value *string `json:"value,omitempty"`
}

// The SoftLayer_Layout_Preference_Type contains definitions for preference types
type Layout_Preference_Type struct {
	Entity

	// The internal identifier of the item type
	Id *int `json:"id,omitempty"`

	// The unique key name of the item type, used primarily for programmatic purposes
	Keyname *string `json:"keyname,omitempty"`

	// The friendly name of the item type
	Name *string `json:"name,omitempty"`

	// A regular expression used to validate the related [[SoftLayer_Layout_Preference]]
	ValueExpression *string `json:"valueExpression,omitempty"`
}

// The SoftLayer_Layout_Profile contains the definition of the layout profile
type Layout_Profile struct {
	Entity

	// Active status of the layout profile
	ActiveFlag *int `json:"activeFlag,omitempty"`

	// Timestamp of when the layout profile was created
	CreateDate *Time `json:"createDate,omitempty"`

	// The internal identifier of a layout profile
	Id *int `json:"id,omitempty"`

	// A count of
	LayoutContainerCount *uint `json:"layoutContainerCount,omitempty"`

	//
	LayoutContainers []Layout_Container `json:"layoutContainers,omitempty"`

	// A count of
	LayoutPreferenceCount *uint `json:"layoutPreferenceCount,omitempty"`

	//
	LayoutPreferences []Layout_Profile_Preference `json:"layoutPreferences,omitempty"`

	// Timestamp of when the layout profile was last updated
	ModifyDate *Time `json:"modifyDate,omitempty"`

	// The friendly name of the layout profile
	Name *string `json:"name,omitempty"`

	// The [[SoftLayer_User_Customer]] owning this layout profile
	UserRecordId *int `json:"userRecordId,omitempty"`
}

//
type Layout_Profile_Containers struct {
	Entity

	// Timestamp of when the reference was created
	CreateDate *Time `json:"createDate,omitempty"`

	// The internal identifier of the container reference
	Id *int `json:"id,omitempty"`

	// The id of the referenced [[SoftLayer_Layout_Container]]
	LayoutContainerId *int `json:"layoutContainerId,omitempty"`

	// The container to be contained
	LayoutContainerType *Layout_Container `json:"layoutContainerType,omitempty"`

	// The profile containing this container
	LayoutProfile *Layout_Profile `json:"layoutProfile,omitempty"`

	// The id of the referenced [[SoftLayer_Layout_Profile]]
	LayoutProfileId *int `json:"layoutProfileId,omitempty"`

	// Timestamp of when the reference was last updated
	ModifyDate *Time `json:"modifyDate,omitempty"`
}

//
type Layout_Profile_Customer struct {
	Layout_Profile

	//
	UserRecord *User_Customer `json:"userRecord,omitempty"`
}

// The SoftLayer_Layout_Profile_Preference contains definitions for layout preferences
type Layout_Profile_Preference struct {
	Entity

	// Timestamp of when the preference was created
	CreateDate *Time `json:"createDate,omitempty"`

	// Indicates whether this is a default value or not
	DefaultValueFlag *int `json:"defaultValueFlag,omitempty"`

	//
	LayoutContainer *Layout_Container `json:"layoutContainer,omitempty"`

	// The id of the related [[SoftLayer_Layout_Container]]
	LayoutContainerId *int `json:"layoutContainerId,omitempty"`

	//
	LayoutItem *Layout_Item `json:"layoutItem,omitempty"`

	// The id of the related [[SoftLayer_Layout_Item]]
	LayoutItemId *int `json:"layoutItemId,omitempty"`

	//
	LayoutPreference *Layout_Preference `json:"layoutPreference,omitempty"`

	// The internal identifier of the overridden [[SoftLayer_Layout_Preference]]
	LayoutPreferenceId *int `json:"layoutPreferenceId,omitempty"`

	//
	LayoutProfile *Layout_Profile `json:"layoutProfile,omitempty"`

	// The internal identifier of the related [[SoftLayer_Layout_Profile]]
	LayoutProfileId *int `json:"layoutProfileId,omitempty"`

	// Timestamp of when the preference was last updated
	ModifyDate *Time `json:"modifyDate,omitempty"`

	// The value overriding the default value
	Value *string `json:"value,omitempty"`
}
