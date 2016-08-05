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

type Location struct {
	Session *Session
	Options
}

func (r *Session) GetLocationService() Location {
	return Location{Session: r}
}

func (r *Location) GetAvailableObjectStorageDatacenters() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location) GetDatacenters() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location) GetDatacentersWithVirtualImageStoreServiceResourceRecord() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location) GetObject() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location) GetViewableDatacenters() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location) GetViewablePopsAndDataCenters() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location) GetViewablepointOfPresence() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location) GetpointOfPresence() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Location) GetBackboneDependents() (resp []datatypes.Network_Backbone_Location_Dependent, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location) GetGroups() (resp []datatypes.Location_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location) GetHardwareFirewalls() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location) GetLocationAddress() (resp datatypes.Account_Address, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location) GetLocationReservationMember() (resp datatypes.Location_Reservation_Rack_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location) GetLocationStatus() (resp datatypes.Location_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location) GetNetworkConfigurationAttribute() (resp datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location) GetOnlinePptpVpnUserCount() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location) GetOnlineSslVpnUserCount() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location) GetPathString() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location) GetPriceGroups() (resp []datatypes.Location_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location) GetRegions() (resp []datatypes.Location_Region, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location) GetTimezone() (resp datatypes.Locale_Timezone, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location) GetVdrGroup() (resp datatypes.Location_Group_Location_CrossReference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Location_Datacenter struct {
	Session *Session
	Options
}

func (r *Session) GetLocationDatacenterService() Location_Datacenter {
	return Location_Datacenter{Session: r}
}

func (r *Location_Datacenter) GetObject() (resp datatypes.Location_Datacenter, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location_Datacenter) GetStatisticsGraphImage() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Location_Datacenter) GetActiveItemPresaleEvents() (resp []datatypes.Sales_Presale_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location_Datacenter) GetActivePresaleEvents() (resp []datatypes.Sales_Presale_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location_Datacenter) GetBackendHardwareRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location_Datacenter) GetBoundSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location_Datacenter) GetBrandCountryRestrictions() (resp []datatypes.Brand_Restriction_Location_CustomerCountry, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location_Datacenter) GetFrontendHardwareRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location_Datacenter) GetHardwareRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location_Datacenter) GetPresaleEvents() (resp []datatypes.Sales_Presale_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location_Datacenter) GetRegionalGroup() (resp datatypes.Location_Group_Regional, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location_Datacenter) GetRegionalInternetRegistry() (resp datatypes.Network_Regional_Internet_Registry, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location_Datacenter) GetRoutableBoundSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Location_Group struct {
	Session *Session
	Options
}

func (r *Session) GetLocationGroupService() Location_Group {
	return Location_Group{Session: r}
}

func (r *Location_Group) GetAllObjects() (resp []datatypes.Location_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location_Group) GetObject() (resp datatypes.Location_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Location_Group) GetLocationGroupType() (resp datatypes.Location_Group_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location_Group) GetLocations() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Location_Group_Pricing struct {
	Session *Session
	Options
}

func (r *Session) GetLocationGroupPricingService() Location_Group_Pricing {
	return Location_Group_Pricing{Session: r}
}

func (r *Location_Group_Pricing) GetAllObjects() (resp []datatypes.Location_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location_Group_Pricing) GetObject() (resp datatypes.Location_Group_Pricing, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Location_Group_Pricing) GetPrices() (resp []datatypes.Product_Item_Price, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Location_Group_Regional struct {
	Session *Session
	Options
}

func (r *Session) GetLocationGroupRegionalService() Location_Group_Regional {
	return Location_Group_Regional{Session: r}
}

func (r *Location_Group_Regional) GetAllObjects() (resp []datatypes.Location_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location_Group_Regional) GetObject() (resp datatypes.Location_Group_Regional, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Location_Group_Regional) GetDatacenters() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location_Group_Regional) GetPreferredDatacenter() (resp datatypes.Location_Datacenter, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Location_Reservation struct {
	Session *Session
	Options
}

func (r *Session) GetLocationReservationService() Location_Reservation {
	return Location_Reservation{Session: r}
}

func (r *Location_Reservation) GetAccountReservations() (resp []datatypes.Location_Reservation, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location_Reservation) GetObject() (resp datatypes.Location_Reservation, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Location_Reservation) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location_Reservation) GetAllotment() (resp datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location_Reservation) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location_Reservation) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location_Reservation) GetLocationReservationRack() (resp datatypes.Location_Reservation_Rack, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Location_Reservation_Rack struct {
	Session *Session
	Options
}

func (r *Session) GetLocationReservationRackService() Location_Reservation_Rack {
	return Location_Reservation_Rack{Session: r}
}

func (r *Location_Reservation_Rack) GetObject() (resp datatypes.Location_Reservation_Rack, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Location_Reservation_Rack) GetAllotment() (resp datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location_Reservation_Rack) GetChildren() (resp []datatypes.Location_Reservation_Rack_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location_Reservation_Rack) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location_Reservation_Rack) GetLocationReservation() (resp datatypes.Location_Reservation, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Location_Reservation_Rack_Member struct {
	Session *Session
	Options
}

func (r *Session) GetLocationReservationRackMemberService() Location_Reservation_Rack_Member {
	return Location_Reservation_Rack_Member{Session: r}
}

func (r *Location_Reservation_Rack_Member) GetObject() (resp datatypes.Location_Reservation_Rack_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Location_Reservation_Rack_Member) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Location_Reservation_Rack_Member) GetLocationReservationRack() (resp datatypes.Location_Reservation, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
