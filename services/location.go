/**
 * Copyright 2016-2024 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
 * the License. You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed
 * on an "AS IS" BASIS,WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and limitations under the License.
 */

// AUTOMATICALLY GENERATED CODE - DO NOT MODIFY

package services

import (
	"fmt"
	"strings"

	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/sl"
)

// Every piece of hardware and network connection owned by SoftLayer is tracked physically by location and stored in the SoftLayer_Location data type. SoftLayer locations exist in parent/child relationships, a convenient way to track equipment from it's city, datacenter, server room, rack, then slot. Network backbones are tied to datacenters only, not to a room, rack, or slot.
type Location struct {
	Session session.SLSession
	Options sl.Options
}

// GetLocationService returns an instance of the Location SoftLayer service
func GetLocationService(sess session.SLSession) Location {
	return Location{Session: sess}
}

func (r Location) Id(id int) Location {
	r.Options.Id = &id
	return r
}

func (r Location) Mask(mask string) Location {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Location) Filter(filter string) Location {
	r.Options.Filter = filter
	return r
}

func (r Location) Limit(limit int) Location {
	r.Options.Limit = &limit
	return r
}

func (r Location) Offset(offset int) Location {
	r.Options.Offset = &offset
	return r
}

// Retrieve all datacenter locations. SoftLayer's datacenters exist in various cities and each contain one or more server rooms which house network and server infrastructure.
func (r Location) GetDatacenters() (resp []datatypes.Location, err error) {
	err = r.Session.DoRequest("SoftLayer_Location", "getDatacenters", nil, &r.Options, &resp)
	return
}

// SoftLayer_Location_Datacenter extends the [[SoftLayer_Location]] data type to include datacenter-specific properties.
type Location_Datacenter struct {
	Session session.SLSession
	Options sl.Options
}

// GetLocationDatacenterService returns an instance of the Location_Datacenter SoftLayer service
func GetLocationDatacenterService(sess session.SLSession) Location_Datacenter {
	return Location_Datacenter{Session: sess}
}

func (r Location_Datacenter) Id(id int) Location_Datacenter {
	r.Options.Id = &id
	return r
}

func (r Location_Datacenter) Mask(mask string) Location_Datacenter {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Location_Datacenter) Filter(filter string) Location_Datacenter {
	r.Options.Filter = filter
	return r
}

func (r Location_Datacenter) Limit(limit int) Location_Datacenter {
	r.Options.Limit = &limit
	return r
}

func (r Location_Datacenter) Offset(offset int) Location_Datacenter {
	r.Options.Offset = &offset
	return r
}

// Retrieve all datacenter locations. SoftLayer's datacenters exist in various cities and each contain one or more server rooms which house network and server infrastructure.
func (r Location_Datacenter) GetDatacenters() (resp []datatypes.Location, err error) {
	err = r.Session.DoRequest("SoftLayer_Location_Datacenter", "getDatacenters", nil, &r.Options, &resp)
	return
}

// Retrieve
func (r Location_Datacenter) GetHardwareRouters() (resp []datatypes.Hardware, err error) {
	err = r.Session.DoRequest("SoftLayer_Location_Datacenter", "getHardwareRouters", nil, &r.Options, &resp)
	return
}

// no documentation yet
type Location_Group struct {
	Session session.SLSession
	Options sl.Options
}

// GetLocationGroupService returns an instance of the Location_Group SoftLayer service
func GetLocationGroupService(sess session.SLSession) Location_Group {
	return Location_Group{Session: sess}
}

func (r Location_Group) Id(id int) Location_Group {
	r.Options.Id = &id
	return r
}

func (r Location_Group) Mask(mask string) Location_Group {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Location_Group) Filter(filter string) Location_Group {
	r.Options.Filter = filter
	return r
}

func (r Location_Group) Limit(limit int) Location_Group {
	r.Options.Limit = &limit
	return r
}

func (r Location_Group) Offset(offset int) Location_Group {
	r.Options.Offset = &offset
	return r
}

// no documentation yet
func (r Location_Group) GetAllObjects() (resp []datatypes.Location_Group, err error) {
	err = r.Session.DoRequest("SoftLayer_Location_Group", "getAllObjects", nil, &r.Options, &resp)
	return
}
