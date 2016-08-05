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

package datatypes

type Location struct {
	Entity

	BackboneDependentCount        *uint                                   `json:"backboneDependentCount,omitempty"`
	BackboneDependents            []Network_Backbone_Location_Dependent   `json:"backboneDependents,omitempty"`
	GroupCount                    *uint                                   `json:"groupCount,omitempty"`
	Groups                        []Location_Group                        `json:"groups,omitempty"`
	HardwareFirewallCount         *uint                                   `json:"hardwareFirewallCount,omitempty"`
	HardwareFirewalls             []Hardware                              `json:"hardwareFirewalls,omitempty"`
	Id                            *int                                    `json:"id,omitempty"`
	LocationAddress               *Account_Address                        `json:"locationAddress,omitempty"`
	LocationReservationMember     *Location_Reservation_Rack_Member       `json:"locationReservationMember,omitempty"`
	LocationStatus                *Location_Status                        `json:"locationStatus,omitempty"`
	LongName                      *string                                 `json:"longName,omitempty"`
	Name                          *string                                 `json:"name,omitempty"`
	NetworkConfigurationAttribute *Hardware_Attribute                     `json:"networkConfigurationAttribute,omitempty"`
	OnlinePptpVpnUserCount        *int                                    `json:"onlinePptpVpnUserCount,omitempty"`
	OnlineSslVpnUserCount         *int                                    `json:"onlineSslVpnUserCount,omitempty"`
	PathString                    *string                                 `json:"pathString,omitempty"`
	PriceGroupCount               *uint                                   `json:"priceGroupCount,omitempty"`
	PriceGroups                   []Location_Group                        `json:"priceGroups,omitempty"`
	RegionCount                   *uint                                   `json:"regionCount,omitempty"`
	Regions                       []Location_Region                       `json:"regions,omitempty"`
	StatusId                      *int                                    `json:"statusId,omitempty"`
	Timezone                      *Locale_Timezone                        `json:"timezone,omitempty"`
	VdrGroup                      *Location_Group_Location_CrossReference `json:"vdrGroup,omitempty"`
}

type Location_Datacenter struct {
	Location

	ActiveItemPresaleEventCount  *uint                                        `json:"activeItemPresaleEventCount,omitempty"`
	ActiveItemPresaleEvents      []Sales_Presale_Event                        `json:"activeItemPresaleEvents,omitempty"`
	ActivePresaleEventCount      *uint                                        `json:"activePresaleEventCount,omitempty"`
	ActivePresaleEvents          []Sales_Presale_Event                        `json:"activePresaleEvents,omitempty"`
	BackendHardwareRouterCount   *uint                                        `json:"backendHardwareRouterCount,omitempty"`
	BackendHardwareRouters       []Hardware                                   `json:"backendHardwareRouters,omitempty"`
	BoundSubnetCount             *uint                                        `json:"boundSubnetCount,omitempty"`
	BoundSubnets                 []Network_Subnet                             `json:"boundSubnets,omitempty"`
	BrandCountryRestrictionCount *uint                                        `json:"brandCountryRestrictionCount,omitempty"`
	BrandCountryRestrictions     []Brand_Restriction_Location_CustomerCountry `json:"brandCountryRestrictions,omitempty"`
	FrontendHardwareRouterCount  *uint                                        `json:"frontendHardwareRouterCount,omitempty"`
	FrontendHardwareRouters      []Hardware                                   `json:"frontendHardwareRouters,omitempty"`
	HardwareRouterCount          *uint                                        `json:"hardwareRouterCount,omitempty"`
	HardwareRouters              []Hardware                                   `json:"hardwareRouters,omitempty"`
	PresaleEventCount            *uint                                        `json:"presaleEventCount,omitempty"`
	PresaleEvents                []Sales_Presale_Event                        `json:"presaleEvents,omitempty"`
	RegionalGroup                *Location_Group_Regional                     `json:"regionalGroup,omitempty"`
	RegionalInternetRegistry     *Network_Regional_Internet_Registry          `json:"regionalInternetRegistry,omitempty"`
	RoutableBoundSubnetCount     *uint                                        `json:"routableBoundSubnetCount,omitempty"`
	RoutableBoundSubnets         []Network_Subnet                             `json:"routableBoundSubnets,omitempty"`
}

type Location_Group struct {
	Entity

	Description         *string              `json:"description,omitempty"`
	Id                  *int                 `json:"id,omitempty"`
	LocationCount       *uint                `json:"locationCount,omitempty"`
	LocationGroupType   *Location_Group_Type `json:"locationGroupType,omitempty"`
	LocationGroupTypeId *int                 `json:"locationGroupTypeId,omitempty"`
	Locations           []Location           `json:"locations,omitempty"`
	Name                *string              `json:"name,omitempty"`
	SecurityLevelId     *int                 `json:"securityLevelId,omitempty"`
}

type Location_Group_Location_CrossReference struct {
	Entity

	Location        *Location       `json:"location,omitempty"`
	LocationGroup   *Location_Group `json:"locationGroup,omitempty"`
	LocationGroupId *int            `json:"locationGroupId,omitempty"`
	LocationId      *int            `json:"locationId,omitempty"`
	Priority        *int            `json:"priority,omitempty"`
}

type Location_Group_Pricing struct {
	Location_Group

	PriceCount *uint                `json:"priceCount,omitempty"`
	Prices     []Product_Item_Price `json:"prices,omitempty"`
}

type Location_Group_Regional struct {
	Location_Group

	DatacenterCount     *uint                `json:"datacenterCount,omitempty"`
	Datacenters         []Location           `json:"datacenters,omitempty"`
	PreferredDatacenter *Location_Datacenter `json:"preferredDatacenter,omitempty"`
}

type Location_Group_Type struct {
	Entity

	Name *string `json:"name,omitempty"`
}

type Location_Inventory_Room struct {
	Location
}

type Location_Network_Operations_Center struct {
	Location
}

type Location_Office struct {
	Location
}

type Location_Rack struct {
	Location
}

type Location_Region struct {
	Entity

	Description *string                   `json:"description,omitempty"`
	Keyname     *string                   `json:"keyname,omitempty"`
	Location    *Location_Region_Location `json:"location,omitempty"`
	SortOrder   *int                      `json:"sortOrder,omitempty"`
}

type Location_Region_Location struct {
	Entity

	Location                   *Location                   `json:"location,omitempty"`
	LocationPackageDetailCount *uint                       `json:"locationPackageDetailCount,omitempty"`
	LocationPackageDetails     []Product_Package_Locations `json:"locationPackageDetails,omitempty"`
	Region                     *Location_Region            `json:"region,omitempty"`
}

type Location_Reservation struct {
	Entity

	Account                 *Account                              `json:"account,omitempty"`
	Allotment               *Network_Bandwidth_Version1_Allotment `json:"allotment,omitempty"`
	AllotmentId             *int                                  `json:"allotmentId,omitempty"`
	BillingItem             *Billing_Item                         `json:"billingItem,omitempty"`
	Id                      *int                                  `json:"id,omitempty"`
	Location                *Location                             `json:"location,omitempty"`
	LocationId              *int                                  `json:"locationId,omitempty"`
	LocationReservationRack *Location_Reservation_Rack            `json:"locationReservationRack,omitempty"`
	Name                    *string                               `json:"name,omitempty"`
	Notes                   *string                               `json:"notes,omitempty"`
}

type Location_Reservation_Rack struct {
	Entity

	Allotment                    *Network_Bandwidth_Version1_Allotment `json:"allotment,omitempty"`
	Children                     []Location_Reservation_Rack_Member    `json:"children,omitempty"`
	ChildrenCount                *uint                                 `json:"childrenCount,omitempty"`
	Location                     *Location                             `json:"location,omitempty"`
	LocationId                   *int                                  `json:"locationId,omitempty"`
	LocationReservation          *Location_Reservation                 `json:"locationReservation,omitempty"`
	LocationReservationId        *int                                  `json:"locationReservationId,omitempty"`
	NetworkConnectionCapacity    *int                                  `json:"networkConnectionCapacity,omitempty"`
	NetworkConnectionReservation *int                                  `json:"networkConnectionReservation,omitempty"`
	PowerConnectionCapacity      *int                                  `json:"powerConnectionCapacity,omitempty"`
	PowerConnectionReservation   *int                                  `json:"powerConnectionReservation,omitempty"`
	SlotCapacity                 *int                                  `json:"slotCapacity,omitempty"`
	SlotReservation              *int                                  `json:"slotReservation,omitempty"`
}

type Location_Reservation_Rack_Member struct {
	Entity

	Id                      *int                  `json:"id,omitempty"`
	Location                *Location             `json:"location,omitempty"`
	LocationId              *int                  `json:"locationId,omitempty"`
	LocationReservationRack *Location_Reservation `json:"locationReservationRack,omitempty"`
}

type Location_Root struct {
	Location
}

type Location_Server_Room struct {
	Location
}

type Location_Slot struct {
	Location
}

type Location_Status struct {
	Entity

	Id     *int    `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
}

type Location_Storage_Room struct {
	Location
}
