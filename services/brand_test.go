// AUTO GENERATED by tools/loadmeta.go
package services_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/session/sessionfakes"
)

var _ = Describe("Brand Tests", func() {
	var slsession *sessionfakes.FakeSLSession
	BeforeEach(func() {
		slsession = &sessionfakes.FakeSLSession{}
	})

	Context("Testing SoftLayer_Brand service", func() {
		var sl_service services.Brand
		BeforeEach(func() {
			sl_service = services.GetBrandService(slsession)
		})
		Context("SoftLayer_Brand Set Options", func() {
			It("Set Options properly", func() {
				t_id := 1234
				t_filter := "{'testFilter':{'test'}}"
				t_limit := 100
				t_offset := 5
				sl_service = sl_service.Id(t_id).Filter(t_filter).Offset(t_offset).Limit(t_limit)
				Expect(sl_service.Options.Id).To(HaveValue(Equal(t_id)))
				Expect(sl_service.Options.Filter).To(HaveValue(Equal(t_filter)))
				Expect(sl_service.Options.Limit).To(HaveValue(Equal(t_limit)))
				Expect(sl_service.Options.Offset).To(HaveValue(Equal(t_offset)))
			})
		})
		Context("SoftLayer_Brand Set Mask", func() {
			It("Set Options properly", func() {
				t_mask1 := "mask[test,test2]"
				sl_service = sl_service.Mask(t_mask1)
				Expect(sl_service.Options.Mask).To(HaveValue(Equal(t_mask1)))
				// Mask("test,test2") should set the mask to be "mask[test,test2]" aka t_mask1
				sl_service = sl_service.Mask("test,test2")
				Expect(sl_service.Options.Mask).To(HaveValue(Equal(t_mask1)))
			})
		})
		Context("SoftLayer_Brand::createCustomerAccount", func() {
			It("API Call Test", func() {
				_, err := sl_service.CreateCustomerAccount(nil, nil)
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::createObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.CreateObject(nil)
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::disableAccount", func() {
			It("API Call Test", func() {
				err := sl_service.DisableAccount(nil)
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getAccount", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetAccount()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getAllOwnedAccounts", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetAllOwnedAccounts()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getAllTicketSubjects", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetAllTicketSubjects(nil)
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getAllowAccountCreationFlag", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetAllowAccountCreationFlag()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getBillingItemSnapshots", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetBillingItemSnapshots()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getBillingItemSnapshotsForSingleOwnedAccount", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetBillingItemSnapshotsForSingleOwnedAccount(nil)
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getBillingItemSnapshotsWithExternalAccountId", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetBillingItemSnapshotsWithExternalAccountId(nil)
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getBusinessPartner", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetBusinessPartner()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getBusinessPartnerFlag", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetBusinessPartnerFlag()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getCatalog", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetCatalog()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getContactInformation", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetContactInformation()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getContacts", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetContacts()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getCustomerCountryLocationRestrictions", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetCustomerCountryLocationRestrictions()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getDistributor", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetDistributor()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getDistributorChildFlag", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetDistributorChildFlag()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getDistributorFlag", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetDistributorFlag()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getHardware", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetHardware()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getHasAgentAdvancedSupportFlag", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetHasAgentAdvancedSupportFlag()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getHasAgentSupportFlag", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetHasAgentSupportFlag()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getMerchantName", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetMerchantName()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetObject()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getOpenTickets", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetOpenTickets()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getOwnedAccounts", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetOwnedAccounts()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getSecurityLevel", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetSecurityLevel()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getTicketGroups", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetTicketGroups()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getTickets", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetTickets()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getToken", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetToken(nil)
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getUsers", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetUsers()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::getVirtualGuests", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetVirtualGuests()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::isIbmSlicBrand", func() {
			It("API Call Test", func() {
				_, err := sl_service.IsIbmSlicBrand()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::isPlatformServicesBrand", func() {
			It("API Call Test", func() {
				_, err := sl_service.IsPlatformServicesBrand()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::migrateExternalAccount", func() {
			It("API Call Test", func() {
				_, err := sl_service.MigrateExternalAccount(nil)
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::reactivateAccount", func() {
			It("API Call Test", func() {
				err := sl_service.ReactivateAccount(nil)
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::refreshBillingItemSnapshot", func() {
			It("API Call Test", func() {
				_, err := sl_service.RefreshBillingItemSnapshot(nil)
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::verifyCanDisableAccount", func() {
			It("API Call Test", func() {
				err := sl_service.VerifyCanDisableAccount(nil)
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand::verifyCanReactivateAccount", func() {
			It("API Call Test", func() {
				err := sl_service.VerifyCanReactivateAccount(nil)
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

	Context("Testing SoftLayer_Brand_Business_Partner service", func() {
		var sl_service services.Brand_Business_Partner
		BeforeEach(func() {
			sl_service = services.GetBrandBusinessPartnerService(slsession)
		})
		Context("SoftLayer_Brand_Business_Partner Set Options", func() {
			It("Set Options properly", func() {
				t_id := 1234
				t_filter := "{'testFilter':{'test'}}"
				t_limit := 100
				t_offset := 5
				sl_service = sl_service.Id(t_id).Filter(t_filter).Offset(t_offset).Limit(t_limit)
				Expect(sl_service.Options.Id).To(HaveValue(Equal(t_id)))
				Expect(sl_service.Options.Filter).To(HaveValue(Equal(t_filter)))
				Expect(sl_service.Options.Limit).To(HaveValue(Equal(t_limit)))
				Expect(sl_service.Options.Offset).To(HaveValue(Equal(t_offset)))
			})
		})
		Context("SoftLayer_Brand_Business_Partner Set Mask", func() {
			It("Set Options properly", func() {
				t_mask1 := "mask[test,test2]"
				sl_service = sl_service.Mask(t_mask1)
				Expect(sl_service.Options.Mask).To(HaveValue(Equal(t_mask1)))
				// Mask("test,test2") should set the mask to be "mask[test,test2]" aka t_mask1
				sl_service = sl_service.Mask("test,test2")
				Expect(sl_service.Options.Mask).To(HaveValue(Equal(t_mask1)))
			})
		})
		Context("SoftLayer_Brand_Business_Partner::getBrand", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetBrand()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand_Business_Partner::getChannel", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetChannel()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand_Business_Partner::getObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetObject()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand_Business_Partner::getSegment", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetSegment()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

	Context("Testing SoftLayer_Brand_Restriction_Location_CustomerCountry service", func() {
		var sl_service services.Brand_Restriction_Location_CustomerCountry
		BeforeEach(func() {
			sl_service = services.GetBrandRestrictionLocationCustomerCountryService(slsession)
		})
		Context("SoftLayer_Brand_Restriction_Location_CustomerCountry Set Options", func() {
			It("Set Options properly", func() {
				t_id := 1234
				t_filter := "{'testFilter':{'test'}}"
				t_limit := 100
				t_offset := 5
				sl_service = sl_service.Id(t_id).Filter(t_filter).Offset(t_offset).Limit(t_limit)
				Expect(sl_service.Options.Id).To(HaveValue(Equal(t_id)))
				Expect(sl_service.Options.Filter).To(HaveValue(Equal(t_filter)))
				Expect(sl_service.Options.Limit).To(HaveValue(Equal(t_limit)))
				Expect(sl_service.Options.Offset).To(HaveValue(Equal(t_offset)))
			})
		})
		Context("SoftLayer_Brand_Restriction_Location_CustomerCountry Set Mask", func() {
			It("Set Options properly", func() {
				t_mask1 := "mask[test,test2]"
				sl_service = sl_service.Mask(t_mask1)
				Expect(sl_service.Options.Mask).To(HaveValue(Equal(t_mask1)))
				// Mask("test,test2") should set the mask to be "mask[test,test2]" aka t_mask1
				sl_service = sl_service.Mask("test,test2")
				Expect(sl_service.Options.Mask).To(HaveValue(Equal(t_mask1)))
			})
		})
		Context("SoftLayer_Brand_Restriction_Location_CustomerCountry::getAllObjects", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetAllObjects()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand_Restriction_Location_CustomerCountry::getBrand", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetBrand()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand_Restriction_Location_CustomerCountry::getLocation", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetLocation()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Brand_Restriction_Location_CustomerCountry::getObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetObject()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

})
