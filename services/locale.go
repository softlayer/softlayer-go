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

package services

import (
	"github.ibm.com/riethm/gopherlayer/datatypes"
	"github.ibm.com/riethm/gopherlayer/session"
	"github.ibm.com/riethm/gopherlayer/sl"
)

//
type Locale struct {
	Session *session.Session
	Options sl.Options
}

func GetLocaleService(sess *session.Session) Locale {
	return Locale{Session: sess}
}

func (r Locale) Id(id int) Locale {
	r.Options.Id = &id
	return r
}

func (r Locale) Mask(mask string) Locale {
	r.Options.Mask = mask
	return r
}

func (r Locale) Filter(filter string) Locale {
	r.Options.Filter = filter
	return r
}

func (r Locale) Limit(limit int) Locale {
	r.Options.Limit = &limit
	return r
}

func (r Locale) Offset(offset int) Locale {
	r.Options.Offset = &offset
	return r
}

//
func (r Locale) GetClosestToLanguageTag(languageTag *string) (resp datatypes.Locale, err error) {
	params := []interface{}{
		languageTag,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r Locale) GetObject() (resp datatypes.Locale, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Locale_Country struct {
	Session *session.Session
	Options sl.Options
}

func GetLocaleCountryService(sess *session.Session) Locale_Country {
	return Locale_Country{Session: sess}
}

func (r Locale_Country) Id(id int) Locale_Country {
	r.Options.Id = &id
	return r
}

func (r Locale_Country) Mask(mask string) Locale_Country {
	r.Options.Mask = mask
	return r
}

func (r Locale_Country) Filter(filter string) Locale_Country {
	r.Options.Filter = filter
	return r
}

func (r Locale_Country) Limit(limit int) Locale_Country {
	r.Options.Limit = &limit
	return r
}

func (r Locale_Country) Offset(offset int) Locale_Country {
	r.Options.Offset = &offset
	return r
}

// Use this method to retrieve a list of countries and locale information available to the current user.
func (r Locale_Country) GetAvailableCountries() (resp []datatypes.Locale_Country, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Use this method to retrieve a list of countries and locale information such as country code and state/provinces.
func (r Locale_Country) GetCountries() (resp []datatypes.Locale_Country, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r Locale_Country) GetObject() (resp datatypes.Locale_Country, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve States that belong to this country.
func (r Locale_Country) GetStates() (resp []datatypes.Locale_StateProvince, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Each User is assigned a timezone allowing for a precise local timestamp.
type Locale_Timezone struct {
	Session *session.Session
	Options sl.Options
}

func GetLocaleTimezoneService(sess *session.Session) Locale_Timezone {
	return Locale_Timezone{Session: sess}
}

func (r Locale_Timezone) Id(id int) Locale_Timezone {
	r.Options.Id = &id
	return r
}

func (r Locale_Timezone) Mask(mask string) Locale_Timezone {
	r.Options.Mask = mask
	return r
}

func (r Locale_Timezone) Filter(filter string) Locale_Timezone {
	r.Options.Filter = filter
	return r
}

func (r Locale_Timezone) Limit(limit int) Locale_Timezone {
	r.Options.Limit = &limit
	return r
}

func (r Locale_Timezone) Offset(offset int) Locale_Timezone {
	r.Options.Offset = &offset
	return r
}

// Retrieve all timezone objects.
func (r Locale_Timezone) GetAllObjects() (resp []datatypes.Locale_Timezone, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Locale_Timezone object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Locale_Timezone service.
func (r Locale_Timezone) GetObject() (resp datatypes.Locale_Timezone, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
