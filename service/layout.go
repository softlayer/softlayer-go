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

package service

import "github.ibm.com/riethm/gopherlayer/datatypes"

// The SoftLayer_Layout_Container contains definitions for default page layouts
type Layout_Container struct {
	Session *Session
	Options
}

func (r *Session) GetLayoutContainerService() Layout_Container {
	return Layout_Container{Session: r}
}

// Use this method to retrieve all active layout containers that can be customized.
func (r *Layout_Container) GetAllObjects() (resp []datatypes.Layout_Container, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The type of the layout container object
func (r *Layout_Container) GetLayoutContainerType() (resp datatypes.Layout_Container_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The layout items assigned to this layout container
func (r *Layout_Container) GetLayoutItems() (resp []datatypes.Layout_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Layout_Container) GetObject() (resp datatypes.Layout_Container, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The SoftLayer_Layout_Item contains definitions for default layout items
type Layout_Item struct {
	Session *Session
	Options
}

func (r *Session) GetLayoutItemService() Layout_Item {
	return Layout_Item{Session: r}
}

// Retrieve The layout preferences assigned to this layout item
func (r *Layout_Item) GetLayoutItemPreferences() (resp []datatypes.Layout_Preference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The type of the layout item object
func (r *Layout_Item) GetLayoutItemType() (resp datatypes.Layout_Item_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Layout_Item) GetObject() (resp datatypes.Layout_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The SoftLayer_Layout_Profile contains the definition of the layout profile
type Layout_Profile struct {
	Session *Session
	Options
}

func (r *Session) GetLayoutProfileService() Layout_Profile {
	return Layout_Profile{Session: r}
}

// This method creates a new layout profile object.
func (r *Layout_Profile) CreateObject(templateObject *datatypes.Layout_Profile) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// This method deletes an existing layout profile and associated custom preferences
func (r *Layout_Profile) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This method edits an existing layout profile object by passing in a modified instance of the object.
func (r *Layout_Profile) EditObject(templateObject *datatypes.Layout_Profile) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Layout_Profile) GetLayoutContainers() (resp []datatypes.Layout_Container, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Layout_Profile) GetLayoutPreferences() (resp []datatypes.Layout_Profile_Preference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Layout_Profile) GetObject() (resp datatypes.Layout_Profile, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This method modifies an existing associated [[SoftLayer_Layout_Profile_Preference]] object. If the preference object being modified is a default value object, a new record is created to override the default value.
//
// Only preferences that are assigned to a profile may be updated. Attempts to update a non-existent preference object will result in an exception being thrown.
func (r *Layout_Profile) ModifyPreference(templateObject *datatypes.Layout_Profile_Preference) (resp datatypes.Layout_Profile_Preference, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Using this method, multiple [[SoftLayer_Layout_Profile_Preference]] objects may be updated at once.
//
// Refer to [[SoftLayer_Layout_Profile::modifyPreference()]] for more information.
func (r *Layout_Profile) ModifyPreferences(layoutPreferenceObjects []datatypes.Layout_Profile_Preference) (resp []datatypes.Layout_Profile_Preference, err error) {
	params := []interface{}{
		layoutPreferenceObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
type Layout_Profile_Containers struct {
	Session *Session
	Options
}

func (r *Session) GetLayoutProfileContainersService() Layout_Profile_Containers {
	return Layout_Profile_Containers{Session: r}
}

//
func (r *Layout_Profile_Containers) CreateObject(templateObject *datatypes.Layout_Profile_Containers) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Layout_Profile_Containers) EditObject(templateObject *datatypes.Layout_Profile_Containers) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The container to be contained
func (r *Layout_Profile_Containers) GetLayoutContainerType() (resp datatypes.Layout_Container, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The profile containing this container
func (r *Layout_Profile_Containers) GetLayoutProfile() (resp datatypes.Layout_Profile, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Layout_Profile_Containers) GetObject() (resp datatypes.Layout_Profile_Containers, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Layout_Profile_Customer struct {
	Session *Session
	Options
}

func (r *Session) GetLayoutProfileCustomerService() Layout_Profile_Customer {
	return Layout_Profile_Customer{Session: r}
}

// This method creates a new layout profile object.
func (r *Layout_Profile_Customer) CreateObject(templateObject *datatypes.Layout_Profile) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// This method deletes an existing layout profile and associated custom preferences
func (r *Layout_Profile_Customer) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This method edits an existing layout profile object by passing in a modified instance of the object.
func (r *Layout_Profile_Customer) EditObject(templateObject *datatypes.Layout_Profile) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Layout_Profile_Customer) GetLayoutContainers() (resp []datatypes.Layout_Container, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Layout_Profile_Customer) GetLayoutPreferences() (resp []datatypes.Layout_Profile_Preference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Layout_Profile_Customer) GetObject() (resp datatypes.Layout_Profile_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Layout_Profile_Customer) GetUserRecord() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This method modifies an existing associated [[SoftLayer_Layout_Profile_Preference]] object. If the preference object being modified is a default value object, a new record is created to override the default value.
//
// Only preferences that are assigned to a profile may be updated. Attempts to update a non-existent preference object will result in an exception being thrown.
func (r *Layout_Profile_Customer) ModifyPreference(templateObject *datatypes.Layout_Profile_Preference) (resp datatypes.Layout_Profile_Preference, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Using this method, multiple [[SoftLayer_Layout_Profile_Preference]] objects may be updated at once.
//
// Refer to [[SoftLayer_Layout_Profile::modifyPreference()]] for more information.
func (r *Layout_Profile_Customer) ModifyPreferences(layoutPreferenceObjects []datatypes.Layout_Profile_Preference) (resp []datatypes.Layout_Profile_Preference, err error) {
	params := []interface{}{
		layoutPreferenceObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// The SoftLayer_Layout_Profile_Preference contains definitions for layout preferences
type Layout_Profile_Preference struct {
	Session *Session
	Options
}

func (r *Session) GetLayoutProfilePreferenceService() Layout_Profile_Preference {
	return Layout_Profile_Preference{Session: r}
}

// Retrieve
func (r *Layout_Profile_Preference) GetLayoutContainer() (resp datatypes.Layout_Container, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Layout_Profile_Preference) GetLayoutItem() (resp datatypes.Layout_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Layout_Profile_Preference) GetLayoutPreference() (resp datatypes.Layout_Preference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Layout_Profile_Preference) GetLayoutProfile() (resp datatypes.Layout_Profile, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Layout_Profile_Preference) GetObject() (resp datatypes.Layout_Profile_Preference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
