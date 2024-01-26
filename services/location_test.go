// AUTO GENERATED by tools/loadmeta.go
package services_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/session/sessionfakes"
)

var _ = Describe("Location Tests", func() {
	var slsession *sessionfakes.FakeSLSession
	BeforeEach(func() {
		slsession = &sessionfakes.FakeSLSession{}
	})

	Context("Testing SoftLayer_Location service", func() {
		var sl_service services.Location
		BeforeEach(func() {
			sl_service = services.GetLocationService(slsession)
		})
		Context("SoftLayer_Location::getActivePresaleEvents", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetActivePresaleEvents()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location::getAvailableObjectStorageDatacenters", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetAvailableObjectStorageDatacenters()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location::getBackboneDependents", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetBackboneDependents()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location::getBnppCompliantFlag", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetBnppCompliantFlag()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location::getDatacenters", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetDatacenters()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location::getDatacentersWithVirtualImageStoreServiceResourceRecord", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetDatacentersWithVirtualImageStoreServiceResourceRecord()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location::getEuCompliantFlag", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetEuCompliantFlag()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location::getGroups", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetGroups()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location::getHardwareFirewalls", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetHardwareFirewalls()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location::getLocationAddress", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetLocationAddress()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location::getLocationAddresses", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetLocationAddresses()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location::getLocationReservationMember", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetLocationReservationMember()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location::getLocationStatus", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetLocationStatus()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location::getNetworkConfigurationAttribute", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetNetworkConfigurationAttribute()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location::getObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetObject()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location::getOnlineSslVpnUserCount", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetOnlineSslVpnUserCount()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location::getPathString", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetPathString()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location::getPriceGroups", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetPriceGroups()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location::getRegions", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetRegions()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location::getTimezone", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetTimezone()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location::getVdrGroup", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetVdrGroup()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location::getViewableDatacenters", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetViewableDatacenters()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location::getViewablePopsAndDataCenters", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetViewablePopsAndDataCenters()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location::getViewablepointOfPresence", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetViewablepointOfPresence()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location::getpointOfPresence", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetpointOfPresence()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

	Context("Testing SoftLayer_Location_Datacenter service", func() {
		var sl_service services.Location_Datacenter
		BeforeEach(func() {
			sl_service = services.GetLocationDatacenterService(slsession)
		})
		Context("SoftLayer_Location_Datacenter::getActiveItemPresaleEvents", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetActiveItemPresaleEvents()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getActivePresaleEvents", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetActivePresaleEvents()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getAvailableObjectStorageDatacenters", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetAvailableObjectStorageDatacenters()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getBackboneDependents", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetBackboneDependents()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getBackendHardwareRouters", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetBackendHardwareRouters()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getBnppCompliantFlag", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetBnppCompliantFlag()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getBoundSubnets", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetBoundSubnets()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getBrandCountryRestrictions", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetBrandCountryRestrictions()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getDatacenters", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetDatacenters()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getDatacentersWithVirtualImageStoreServiceResourceRecord", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetDatacentersWithVirtualImageStoreServiceResourceRecord()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getEuCompliantFlag", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetEuCompliantFlag()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getFrontendHardwareRouters", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetFrontendHardwareRouters()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getGroups", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetGroups()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getHardwareFirewalls", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetHardwareFirewalls()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getHardwareRouters", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetHardwareRouters()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getLocationAddress", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetLocationAddress()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getLocationAddresses", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetLocationAddresses()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getLocationReservationMember", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetLocationReservationMember()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getLocationStatus", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetLocationStatus()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getNetworkConfigurationAttribute", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetNetworkConfigurationAttribute()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetObject()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getOnlineSslVpnUserCount", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetOnlineSslVpnUserCount()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getPathString", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetPathString()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getPresaleEvents", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetPresaleEvents()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getPriceGroups", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetPriceGroups()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getRegionalGroup", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetRegionalGroup()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getRegionalInternetRegistry", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetRegionalInternetRegistry()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getRegions", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetRegions()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getRoutableBoundSubnets", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetRoutableBoundSubnets()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getStatisticsGraphImage", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetStatisticsGraphImage()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getTimezone", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetTimezone()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getVdrGroup", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetVdrGroup()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getViewableDatacenters", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetViewableDatacenters()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getViewablePopsAndDataCenters", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetViewablePopsAndDataCenters()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getViewablepointOfPresence", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetViewablepointOfPresence()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Datacenter::getpointOfPresence", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetpointOfPresence()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

	Context("Testing SoftLayer_Location_Group service", func() {
		var sl_service services.Location_Group
		BeforeEach(func() {
			sl_service = services.GetLocationGroupService(slsession)
		})
		Context("SoftLayer_Location_Group::getAllObjects", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetAllObjects()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Group::getLocationGroupType", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetLocationGroupType()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Group::getLocations", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetLocations()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Group::getObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetObject()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

	Context("Testing SoftLayer_Location_Group_Pricing service", func() {
		var sl_service services.Location_Group_Pricing
		BeforeEach(func() {
			sl_service = services.GetLocationGroupPricingService(slsession)
		})
		Context("SoftLayer_Location_Group_Pricing::getAllObjects", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetAllObjects()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Group_Pricing::getLocationGroupType", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetLocationGroupType()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Group_Pricing::getLocations", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetLocations()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Group_Pricing::getObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetObject()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Group_Pricing::getPrices", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetPrices()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

	Context("Testing SoftLayer_Location_Group_Regional service", func() {
		var sl_service services.Location_Group_Regional
		BeforeEach(func() {
			sl_service = services.GetLocationGroupRegionalService(slsession)
		})
		Context("SoftLayer_Location_Group_Regional::getAllObjects", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetAllObjects()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Group_Regional::getDatacenters", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetDatacenters()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Group_Regional::getLocationGroupType", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetLocationGroupType()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Group_Regional::getLocations", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetLocations()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Group_Regional::getObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetObject()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Group_Regional::getPreferredDatacenter", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetPreferredDatacenter()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

	Context("Testing SoftLayer_Location_Reservation service", func() {
		var sl_service services.Location_Reservation
		BeforeEach(func() {
			sl_service = services.GetLocationReservationService(slsession)
		})
		Context("SoftLayer_Location_Reservation::getAccount", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetAccount()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Reservation::getAccountReservations", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetAccountReservations()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Reservation::getAllotment", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetAllotment()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Reservation::getBillingItem", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetBillingItem()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Reservation::getLocation", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetLocation()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Reservation::getLocationReservationRack", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetLocationReservationRack()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Reservation::getObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetObject()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

	Context("Testing SoftLayer_Location_Reservation_Rack service", func() {
		var sl_service services.Location_Reservation_Rack
		BeforeEach(func() {
			sl_service = services.GetLocationReservationRackService(slsession)
		})
		Context("SoftLayer_Location_Reservation_Rack::getAllotment", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetAllotment()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Reservation_Rack::getChildren", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetChildren()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Reservation_Rack::getLocation", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetLocation()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Reservation_Rack::getLocationReservation", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetLocationReservation()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Reservation_Rack::getObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetObject()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

	Context("Testing SoftLayer_Location_Reservation_Rack_Member service", func() {
		var sl_service services.Location_Reservation_Rack_Member
		BeforeEach(func() {
			sl_service = services.GetLocationReservationRackMemberService(slsession)
		})
		Context("SoftLayer_Location_Reservation_Rack_Member::getLocation", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetLocation()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Reservation_Rack_Member::getLocationReservationRack", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetLocationReservationRack()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Location_Reservation_Rack_Member::getObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetObject()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

})