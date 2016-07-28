package softlayer

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Location interface {
	Entity

	GetAvailableObjectStorageDatacenters() (datatypes.Location, error)
	GetDatacenters() (datatypes.Location, error)
	GetDatacentersWithVirtualImageStoreServiceResourceRecord() (datatypes.Location, error)
	GetObject() (datatypes.Location, error)
	GetViewableDatacenters() (datatypes.Location, error)
	GetViewablePopsAndDataCenters() (datatypes.Location, error)
	GetViewablepointOfPresence() (datatypes.Location, error)
	GetpointOfPresence() (datatypes.Location, error)

	GetBackboneDependents() (datatypes.Network_Backbone_Location_Dependent, error)
	GetGroups() (datatypes.Location_Group, error)
	GetHardwareFirewalls() (datatypes.Hardware, error)
	GetLocationAddress() (datatypes.Account_Address, error)
	GetLocationReservationMember() (datatypes.Location_Reservation_Rack_Member, error)
	GetLocationStatus() (datatypes.Location_Status, error)
	GetNetworkConfigurationAttribute() (datatypes.Hardware_Attribute, error)
	GetOnlinePptpVpnUserCount() (int, error)
	GetOnlineSslVpnUserCount() (int, error)
	GetPathString() (string, error)
	GetPriceGroups() (datatypes.Location_Group, error)
	GetRegions() (datatypes.Location_Region, error)
	GetTimezone() (datatypes.Locale_Timezone, error)
	GetVdrGroup() (datatypes.Location_Group_Location_CrossReference, error)
}

type Location_Datacenter interface {
	Location

	GetObject() (datatypes.Location_Datacenter, error)
	GetStatisticsGraphImage() ([]byte, error)

	GetActiveItemPresaleEvents() (datatypes.Sales_Presale_Event, error)
	GetActivePresaleEvents() (datatypes.Sales_Presale_Event, error)
	GetBackendHardwareRouters() (datatypes.Hardware, error)
	GetBoundSubnets() (datatypes.Network_Subnet, error)
	GetBrandCountryRestrictions() (datatypes.Brand_Restriction_Location_CustomerCountry, error)
	GetFrontendHardwareRouters() (datatypes.Hardware, error)
	GetHardwareRouters() (datatypes.Hardware, error)
	GetPresaleEvents() (datatypes.Sales_Presale_Event, error)
	GetRegionalGroup() (datatypes.Location_Group_Regional, error)
	GetRegionalInternetRegistry() (datatypes.Network_Regional_Internet_Registry, error)
	GetRoutableBoundSubnets() (datatypes.Network_Subnet, error)
}

type Location_Group interface {
	Entity

	GetAllObjects() (datatypes.Location_Group, error)
	GetObject() (datatypes.Location_Group, error)

	GetLocationGroupType() (datatypes.Location_Group_Type, error)
	GetLocations() (datatypes.Location, error)
}

type Location_Group_Location_CrossReference interface {
	Entity

	GetLocation() (datatypes.Location, error)
	GetLocationGroup() (datatypes.Location_Group, error)
}

type Location_Group_Pricing interface {
	Location_Group

	GetAllObjects() (datatypes.Location_Group, error)
	GetObject() (datatypes.Location_Group_Pricing, error)

	GetPrices() (datatypes.Product_Item_Price, error)
}

type Location_Group_Regional interface {
	Location_Group

	GetAllObjects() (datatypes.Location_Group, error)
	GetObject() (datatypes.Location_Group_Regional, error)

	GetDatacenters() (datatypes.Location, error)
	GetPreferredDatacenter() (datatypes.Location_Datacenter, error)
}

type Location_Group_Type interface {
	Entity
}

type Location_Inventory_Room interface {
	Location
}

type Location_Network_Operations_Center interface {
	Location
}

type Location_Office interface {
	Location
}

type Location_Rack interface {
	Location
}

type Location_Region interface {
	Entity

	GetLocation() (datatypes.Location_Region_Location, error)
}

type Location_Region_Location interface {
	Entity

	GetLocation() (datatypes.Location, error)
	GetLocationPackageDetails() (datatypes.Product_Package_Locations, error)
	GetRegion() (datatypes.Location_Region, error)
}

type Location_Reservation interface {
	Entity

	GetAccountReservations() (datatypes.Location_Reservation, error)
	GetObject() (datatypes.Location_Reservation, error)

	GetAccount() (datatypes.Account, error)
	GetAllotment() (datatypes.Network_Bandwidth_Version1_Allotment, error)
	GetBillingItem() (datatypes.Billing_Item, error)
	GetLocation() (datatypes.Location, error)
	GetLocationReservationRack() (datatypes.Location_Reservation_Rack, error)
}

type Location_Reservation_Rack interface {
	Entity

	GetObject() (datatypes.Location_Reservation_Rack, error)

	GetAllotment() (datatypes.Network_Bandwidth_Version1_Allotment, error)
	GetChildren() (datatypes.Location_Reservation_Rack_Member, error)
	GetLocation() (datatypes.Location, error)
	GetLocationReservation() (datatypes.Location_Reservation, error)
}

type Location_Reservation_Rack_Member interface {
	Entity

	GetObject() (datatypes.Location_Reservation_Rack_Member, error)

	GetLocation() (datatypes.Location, error)
	GetLocationReservationRack() (datatypes.Location_Reservation, error)
}

type Location_Root interface {
	Location
}

type Location_Server_Room interface {
	Location
}

type Location_Slot interface {
	Location
}

type Location_Status interface {
	Entity
}

type Location_Storage_Room interface {
	Location
}
