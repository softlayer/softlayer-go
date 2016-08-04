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

package service

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Layout_Container struct {
	Session *Session
	Options
}

func (r *Session) GetLayoutContainerService() Layout_Container {
	return Layout_Container{Session: r}
}

func (r *Layout_Container) GetAllObjects() (resp []datatypes.Layout_Container, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Layout_Container) GetObject() (resp datatypes.Layout_Container, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Layout_Container) GetLayoutContainerType() (resp datatypes.Layout_Container_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Layout_Container) GetLayoutItems() (resp []datatypes.Layout_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Layout_Container_Type struct {
	Session *Session
	Options
}

func (r *Session) GetLayoutContainerTypeService() Layout_Container_Type {
	return Layout_Container_Type{Session: r}
}

type Layout_Item struct {
	Session *Session
	Options
}

func (r *Session) GetLayoutItemService() Layout_Item {
	return Layout_Item{Session: r}
}

func (r *Layout_Item) GetObject() (resp datatypes.Layout_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Layout_Item) GetLayoutItemPreferences() (resp []datatypes.Layout_Preference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Layout_Item) GetLayoutItemType() (resp datatypes.Layout_Item_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Layout_Item_Type struct {
	Session *Session
	Options
}

func (r *Session) GetLayoutItemTypeService() Layout_Item_Type {
	return Layout_Item_Type{Session: r}
}

type Layout_Preference struct {
	Session *Session
	Options
}

func (r *Session) GetLayoutPreferenceService() Layout_Preference {
	return Layout_Preference{Session: r}
}

func (r *Layout_Preference) GetLayoutPreferenceType() (resp datatypes.Layout_Preference_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Layout_Preference_Type struct {
	Session *Session
	Options
}

func (r *Session) GetLayoutPreferenceTypeService() Layout_Preference_Type {
	return Layout_Preference_Type{Session: r}
}

type Layout_Profile struct {
	Session *Session
	Options
}

func (r *Session) GetLayoutProfileService() Layout_Profile {
	return Layout_Profile{Session: r}
}

func (r *Layout_Profile) CreateObject(templateObject *datatypes.Layout_Profile) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Layout_Profile) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Layout_Profile) EditObject(templateObject *datatypes.Layout_Profile) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Layout_Profile) GetObject() (resp datatypes.Layout_Profile, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Layout_Profile) ModifyPreference(templateObject *datatypes.Layout_Profile_Preference) (resp datatypes.Layout_Profile_Preference, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Layout_Profile) ModifyPreferences(layoutPreferenceObjects []datatypes.Layout_Profile_Preference) (resp []datatypes.Layout_Profile_Preference, err error) {
	params := []interface{}{
		layoutPreferenceObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

func (r *Layout_Profile) GetLayoutContainers() (resp []datatypes.Layout_Container, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Layout_Profile) GetLayoutPreferences() (resp []datatypes.Layout_Profile_Preference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Layout_Profile_Containers struct {
	Session *Session
	Options
}

func (r *Session) GetLayoutProfileContainersService() Layout_Profile_Containers {
	return Layout_Profile_Containers{Session: r}
}

func (r *Layout_Profile_Containers) CreateObject(templateObject *datatypes.Layout_Profile_Containers) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Layout_Profile_Containers) EditObject(templateObject *datatypes.Layout_Profile_Containers) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Layout_Profile_Containers) GetObject() (resp datatypes.Layout_Profile_Containers, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Layout_Profile_Containers) GetLayoutContainerType() (resp datatypes.Layout_Container, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Layout_Profile_Containers) GetLayoutProfile() (resp datatypes.Layout_Profile, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Layout_Profile_Customer struct {
	Session *Session
	Options
}

func (r *Session) GetLayoutProfileCustomerService() Layout_Profile_Customer {
	return Layout_Profile_Customer{Session: r}
}

func (r *Layout_Profile_Customer) GetObject() (resp datatypes.Layout_Profile_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Layout_Profile_Customer) GetUserRecord() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Layout_Profile_Preference struct {
	Session *Session
	Options
}

func (r *Session) GetLayoutProfilePreferenceService() Layout_Profile_Preference {
	return Layout_Profile_Preference{Session: r}
}

func (r *Layout_Profile_Preference) GetObject() (resp datatypes.Layout_Profile_Preference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Layout_Profile_Preference) GetLayoutContainer() (resp datatypes.Layout_Container, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Layout_Profile_Preference) GetLayoutItem() (resp datatypes.Layout_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Layout_Profile_Preference) GetLayoutPreference() (resp datatypes.Layout_Preference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Layout_Profile_Preference) GetLayoutProfile() (resp datatypes.Layout_Profile, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
