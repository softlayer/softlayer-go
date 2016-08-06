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

//
type Locale struct {
	Session *Session
	Options
}

func (r *Session) GetLocaleService() Locale {
	return Locale{Session: r}
}

//
func (r *Locale) GetClosestToLanguageTag(languageTag *string) (resp datatypes.Locale, err error) {
	params := []interface{}{
		languageTag,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Locale) GetObject() (resp datatypes.Locale, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Locale_Country struct {
	Session *Session
	Options
}

func (r *Session) GetLocaleCountryService() Locale_Country {
	return Locale_Country{Session: r}
}

// Use this method to retrieve a list of countries and locale information available to the current user.
func (r *Locale_Country) GetAvailableCountries() (resp []datatypes.Locale_Country, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Use this method to retrieve a list of countries and locale information such as country code and state/provinces.
func (r *Locale_Country) GetCountries() (resp []datatypes.Locale_Country, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Locale_Country) GetObject() (resp datatypes.Locale_Country, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve States that belong to this country.
func (r *Locale_Country) GetStates() (resp []datatypes.Locale_StateProvince, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Each User is assigned a timezone allowing for a precise local timestamp.
type Locale_Timezone struct {
	Session *Session
	Options
}

func (r *Session) GetLocaleTimezoneService() Locale_Timezone {
	return Locale_Timezone{Session: r}
}

// Retrieve all timezone objects.
func (r *Locale_Timezone) GetAllObjects() (resp []datatypes.Locale_Timezone, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Locale_Timezone object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Locale_Timezone service.
func (r *Locale_Timezone) GetObject() (resp datatypes.Locale_Timezone, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
