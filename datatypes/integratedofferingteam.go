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
type IntegratedOfferingTeam_Container_Region struct {
	Entity

	// no documentation yet
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName,omitempty"`

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name,omitempty"`
}

// no documentation yet
type IntegratedOfferingTeam_Container_Region_Lead struct {
	Entity

	// Regional lead's email address
	EmailAddress *string `json:"emailAddress,omitempty" xmlrpc:"emailAddress,omitempty"`

	// Regional lead's first name
	FirstName *string `json:"firstName,omitempty" xmlrpc:"firstName,omitempty"`

	// Regional lead's last name
	LastName *string `json:"lastName,omitempty" xmlrpc:"lastName,omitempty"`

	// Key name of the region this lead is in charge of
	RegionKeyName *string `json:"regionKeyName,omitempty" xmlrpc:"regionKeyName,omitempty"`

	// Full name of the region this lead is in charge of
	RegionName *string `json:"regionName,omitempty" xmlrpc:"regionName,omitempty"`
}

// This class represents an Integrated Offering Team region.
type IntegratedOfferingTeam_Region struct {
	Entity
}
