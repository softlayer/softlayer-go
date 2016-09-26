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

// no documentation yet
type Locale struct {
	Entity

	// no documentation yet
	FriendlyName *string `json:"friendlyName,omitempty" xmlrpc:"friendlyName"`

	// Internal identification number of a locale
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	LanguageTag *string `json:"languageTag,omitempty" xmlrpc:"languageTag"`

	// Locale name
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// no documentation yet
type Locale_Country struct {
	Entity

	// Binary flag denoting if this country is part of the European Union
	IsEuropeanUnionFlag *int `json:"isEuropeanUnionFlag,omitempty" xmlrpc:"isEuropeanUnionFlag"`

	// no documentation yet
	LongName *string `json:"longName,omitempty" xmlrpc:"longName"`

	// no documentation yet
	ShortName *string `json:"shortName,omitempty" xmlrpc:"shortName"`

	// A count of states that belong to this country.
	StateCount *uint `json:"stateCount,omitempty" xmlrpc:"stateCount"`

	// States that belong to this country.
	States []Locale_StateProvince `json:"states,omitempty" xmlrpc:"states"`
}

// This object represents a state or province for a country.
type Locale_StateProvince struct {
	Entity

	// no documentation yet
	LongName *string `json:"longName,omitempty" xmlrpc:"longName"`

	// no documentation yet
	ShortName *string `json:"shortName,omitempty" xmlrpc:"shortName"`
}

// Each User is assigned a timezone allowing for a precise local timestamp.
type Locale_Timezone struct {
	Entity

	// A timezone's identifying number.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// A timezone's long name. For example, "(GMT-06:00) America/Dallas - CST".
	LongName *string `json:"longName,omitempty" xmlrpc:"longName"`

	// A timezone's name. For example, "America/Dallas".
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// A timezone's offset based on the GMT standard. For example, Central Standard Time's offset is "-0600" from GMT=0000.
	Offset *string `json:"offset,omitempty" xmlrpc:"offset"`

	// A timezone's common abbreviation. For example, Central Standard Time's abbreviation is "CST".
	ShortName *string `json:"shortName,omitempty" xmlrpc:"shortName"`
}
