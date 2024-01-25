
package services_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/softlayer/softlayer-go/session/sessionfakes"
    "github.com/softlayer/softlayer-go/services"
)   

var _ = Describe("Account Tests", func() {
    var slsession *sessionfakes.FakeSLSession
    BeforeEach(func() {
        slsession = &sessionfakes.FakeSLSession{}
    })

    Context("Testing SoftLayer_Account service", func() {
        var sl_service services.Account
        BeforeEach(func() {
            sl_service = services.GetAccountService(slsession)
        })


        Context("SoftLayer_Account::activatePartner", func() {
            It("API Call Test", func() {

                _, err := sl_service.ActivatePartner(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::addAchInformation", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddAchInformation(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::addReferralPartnerPaymentOption", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddReferralPartnerPaymentOption(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::areVdrUpdatesBlockedForBilling", func() {
            It("API Call Test", func() {

                _, err := sl_service.AreVdrUpdatesBlockedForBilling()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::cancelPayPalTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.CancelPayPalTransaction(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::completePayPalTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.CompletePayPalTransaction(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::countHourlyInstances", func() {
            It("API Call Test", func() {

                _, err := sl_service.CountHourlyInstances()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::createUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateUser(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::disableEuSupport", func() {
            It("API Call Test", func() {

				err := sl_service.DisableEuSupport()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::disableVpnConfigRequiresVpnManageAttribute", func() {
            It("API Call Test", func() {

				err := sl_service.DisableVpnConfigRequiresVpnManageAttribute()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::editAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditAccount(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::enableEuSupport", func() {
            It("API Call Test", func() {

				err := sl_service.EnableEuSupport()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::enableVpnConfigRequiresVpnManageAttribute", func() {
            It("API Call Test", func() {

				err := sl_service.EnableVpnConfigRequiresVpnManageAttribute()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAbuseEmail", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAbuseEmail()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAbuseEmails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAbuseEmails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAccountBackupHistory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccountBackupHistory(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAccountContacts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccountContacts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAccountLicenses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccountLicenses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAccountLinks", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccountLinks()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAccountStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccountStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAccountTraitValue", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccountTraitValue(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getActiveAccountDiscountBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveAccountDiscountBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getActiveAccountLicenses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveAccountLicenses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getActiveAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getActiveAgreements", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveAgreements()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getActiveBillingAgreements", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveBillingAgreements()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getActiveCatalystEnrollment", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveCatalystEnrollment()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getActiveColocationContainers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveColocationContainers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getActiveFlexibleCreditEnrollment", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveFlexibleCreditEnrollment()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getActiveFlexibleCreditEnrollments", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveFlexibleCreditEnrollments()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getActiveNotificationSubscribers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveNotificationSubscribers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getActiveOutletPackages", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveOutletPackages()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getActivePackages", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActivePackages()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getActivePackagesByAttribute", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActivePackagesByAttribute(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getActivePrivateHostedCloudPackages", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActivePrivateHostedCloudPackages()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getActiveQuotes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveQuotes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getActiveReservedCapacityAgreements", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveReservedCapacityAgreements()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getActiveVirtualLicenses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveVirtualLicenses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAdcLoadBalancers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAdcLoadBalancers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAffiliateId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAffiliateId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAllBillingItems", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllBillingItems()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAllCommissionBillingItems", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllCommissionBillingItems()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAllRecurringTopLevelBillingItems", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllRecurringTopLevelBillingItems()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAllRecurringTopLevelBillingItemsUnfiltered", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllRecurringTopLevelBillingItemsUnfiltered()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAllSubnetBillingItems", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllSubnetBillingItems()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAllTopLevelBillingItems", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllTopLevelBillingItems()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAllTopLevelBillingItemsUnfiltered", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllTopLevelBillingItemsUnfiltered()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAllowIbmIdSilentMigrationFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowIbmIdSilentMigrationFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAllowsBluemixAccountLinkingFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowsBluemixAccountLinkingFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAlternateCreditCardData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAlternateCreditCardData()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getApplicationDeliveryControllers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetApplicationDeliveryControllers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAttributeByType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttributeByType(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAttributes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttributes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAuxiliaryNotifications", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAuxiliaryNotifications()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAvailablePublicNetworkVlans", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailablePublicNetworkVlans()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAverageArchiveUsageMetricDataByDate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAverageArchiveUsageMetricDataByDate(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getAveragePublicUsageMetricDataByDate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAveragePublicUsageMetricDataByDate(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getBalance", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBalance()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getBandwidthAllotments", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthAllotments()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getBandwidthAllotmentsOverAllocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthAllotmentsOverAllocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getBandwidthAllotmentsProjectedOverAllocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthAllotmentsProjectedOverAllocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getBandwidthList", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthList(nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getBareMetalInstances", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBareMetalInstances()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getBillingAgreements", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingAgreements()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getBillingInfo", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingInfo()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getBlockDeviceTemplateGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBlockDeviceTemplateGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getBluemixAccountId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBluemixAccountId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getBluemixAccountLink", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBluemixAccountLink()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getBluemixLinkedFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBluemixLinkedFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getBrand", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBrand()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getBrandAccountFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBrandAccountFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getBrandKeyName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBrandKeyName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getBusinessPartner", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBusinessPartner()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getCanOrderAdditionalVlansFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCanOrderAdditionalVlansFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getCarts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCarts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getCatalystEnrollments", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCatalystEnrollments()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getClosedTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetClosedTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getCurrentUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getDatacentersWithSubnetAllocations", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDatacentersWithSubnetAllocations()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getDedicatedHosts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDedicatedHosts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getDedicatedHostsForImageTemplate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDedicatedHostsForImageTemplate(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getDisablePaymentProcessingFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDisablePaymentProcessingFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getDisplaySupportRepresentativeAssignments", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDisplaySupportRepresentativeAssignments()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getDomainRegistrations", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDomainRegistrations()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getDomains", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDomains()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getDomainsWithoutSecondaryDnsRecords", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDomainsWithoutSecondaryDnsRecords()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getEuSupportedFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEuSupportedFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getEvaultCapacityGB", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEvaultCapacityGB()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getEvaultMasterUsers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEvaultMasterUsers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getEvaultNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEvaultNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getExpiredSecurityCertificates", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetExpiredSecurityCertificates()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getFacilityLogs", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFacilityLogs()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getFileBlockBetaAccessFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFileBlockBetaAccessFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getFlexibleCreditEnrollments", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFlexibleCreditEnrollments()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getFlexibleCreditProgramInfo", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFlexibleCreditProgramInfo(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getFlexibleCreditProgramsInfo", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFlexibleCreditProgramsInfo(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getForcePaasAccountLinkDate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetForcePaasAccountLinkDate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getGlobalIpRecords", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGlobalIpRecords()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getGlobalIpv4Records", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGlobalIpv4Records()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getGlobalIpv6Records", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGlobalIpv6Records()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getGlobalLoadBalancerAccounts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGlobalLoadBalancerAccounts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getHardwareOverBandwidthAllocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareOverBandwidthAllocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getHardwarePools", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwarePools()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getHardwareProjectedOverBandwidthAllocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareProjectedOverBandwidthAllocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getHardwareWithCpanel", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareWithCpanel()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getHardwareWithHelm", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareWithHelm()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getHardwareWithMcafee", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareWithMcafee()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getHardwareWithMcafeeAntivirusRedhat", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareWithMcafeeAntivirusRedhat()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getHardwareWithMcafeeAntivirusWindows", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareWithMcafeeAntivirusWindows()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getHardwareWithMcafeeIntrusionDetectionSystem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareWithMcafeeIntrusionDetectionSystem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getHardwareWithPlesk", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareWithPlesk()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getHardwareWithQuantastor", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareWithQuantastor()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getHardwareWithUrchin", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareWithUrchin()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getHardwareWithWindows", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareWithWindows()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getHasEvaultBareMetalRestorePluginFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasEvaultBareMetalRestorePluginFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getHasIderaBareMetalRestorePluginFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasIderaBareMetalRestorePluginFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getHasPendingOrder", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasPendingOrder()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getHasR1softBareMetalRestorePluginFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasR1softBareMetalRestorePluginFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getHourlyBareMetalInstances", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHourlyBareMetalInstances()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getHourlyServiceBillingItems", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHourlyServiceBillingItems()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getHourlyVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHourlyVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getHubNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHubNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getIbmCustomerNumber", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIbmCustomerNumber()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getIbmIdAuthenticationRequiredFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIbmIdAuthenticationRequiredFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getIbmIdMigrationExpirationTimestamp", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIbmIdMigrationExpirationTimestamp()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getInProgressExternalAccountSetup", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInProgressExternalAccountSetup()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getInternalCciHostAccountFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInternalCciHostAccountFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getInternalImageTemplateCreationFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInternalImageTemplateCreationFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getInternalNotes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInternalNotes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getInternalRestrictionFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInternalRestrictionFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getInvoices", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInvoices()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getIscsiIsolationDisabled", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIscsiIsolationDisabled()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getIscsiNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIscsiNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getLargestAllowedSubnetCidr", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLargestAllowedSubnetCidr(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getLastCanceledBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLastCanceledBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getLastCancelledServerBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLastCancelledServerBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getLastFiveClosedAbuseTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLastFiveClosedAbuseTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getLastFiveClosedAccountingTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLastFiveClosedAccountingTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getLastFiveClosedOtherTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLastFiveClosedOtherTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getLastFiveClosedSalesTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLastFiveClosedSalesTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getLastFiveClosedSupportTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLastFiveClosedSupportTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getLastFiveClosedTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLastFiveClosedTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getLatestBillDate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLatestBillDate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getLatestRecurringInvoice", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLatestRecurringInvoice()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getLatestRecurringPendingInvoice", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLatestRecurringPendingInvoice()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getLegacyBandwidthAllotments", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLegacyBandwidthAllotments()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getLegacyIscsiCapacityGB", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLegacyIscsiCapacityGB()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getLoadBalancers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLoadBalancers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getLockboxCapacityGB", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLockboxCapacityGB()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getLockboxNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLockboxNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getManualPaymentsUnderReview", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetManualPaymentsUnderReview()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getMasterUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMasterUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getMediaDataTransferRequests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMediaDataTransferRequests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getMetricTrackingObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMetricTrackingObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getMigratedToIbmCloudPortalFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMigratedToIbmCloudPortalFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getMonthlyBareMetalInstances", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMonthlyBareMetalInstances()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getMonthlyVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMonthlyVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNasNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNasNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNetAppActiveAccountLicenseKeys", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetAppActiveAccountLicenseKeys()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNetworkCreationFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkCreationFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNetworkGateways", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkGateways()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNetworkHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNetworkMessageDeliveryAccounts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMessageDeliveryAccounts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNetworkMonitorDownHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitorDownHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNetworkMonitorDownVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitorDownVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNetworkMonitorRecoveringHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitorRecoveringHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNetworkMonitorRecoveringVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitorRecoveringVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNetworkMonitorUpHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitorUpHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNetworkMonitorUpVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitorUpVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNetworkStorageGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkStorageGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNetworkTunnelContexts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkTunnelContexts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNetworkVlanSpan", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkVlanSpan()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNetworkVlans", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkVlans()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNextBillingPublicAllotmentHardwareBandwidthDetails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNextBillingPublicAllotmentHardwareBandwidthDetails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNextInvoiceExcel", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNextInvoiceExcel(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNextInvoiceIncubatorExemptTotal", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNextInvoiceIncubatorExemptTotal()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNextInvoicePdf", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNextInvoicePdf(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNextInvoicePdfDetailed", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNextInvoicePdfDetailed(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNextInvoiceRecurringAmountEligibleForAccountDiscount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNextInvoiceRecurringAmountEligibleForAccountDiscount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNextInvoiceTopLevelBillingItems", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNextInvoiceTopLevelBillingItems()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNextInvoiceTotalAmount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNextInvoiceTotalAmount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNextInvoiceTotalOneTimeAmount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNextInvoiceTotalOneTimeAmount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNextInvoiceTotalOneTimeTaxAmount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNextInvoiceTotalOneTimeTaxAmount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNextInvoiceTotalRecurringAmount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNextInvoiceTotalRecurringAmount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNextInvoiceTotalRecurringAmountBeforeAccountDiscount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNextInvoiceTotalRecurringAmountBeforeAccountDiscount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNextInvoiceTotalRecurringTaxAmount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNextInvoiceTotalRecurringTaxAmount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNextInvoiceTotalTaxableRecurringAmount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNextInvoiceTotalTaxableRecurringAmount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNextInvoiceZeroFeeItemCounts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNextInvoiceZeroFeeItemCounts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getNotificationSubscribers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNotificationSubscribers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getOpenAbuseTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOpenAbuseTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getOpenAccountingTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOpenAccountingTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getOpenBillingTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOpenBillingTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getOpenCancellationRequests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOpenCancellationRequests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getOpenOtherTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOpenOtherTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getOpenRecurringInvoices", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOpenRecurringInvoices()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getOpenSalesTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOpenSalesTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getOpenStackAccountLinks", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOpenStackAccountLinks()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getOpenStackObjectStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOpenStackObjectStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getOpenSupportTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOpenSupportTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getOpenTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOpenTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getOpenTicketsWaitingOnCustomer", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOpenTicketsWaitingOnCustomer()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getOrders", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOrders()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getOrphanBillingItems", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOrphanBillingItems()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getOwnedBrands", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOwnedBrands()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getOwnedHardwareGenericComponentModels", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOwnedHardwareGenericComponentModels()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPaymentProcessors", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPaymentProcessors()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPendingCreditCardChangeRequestData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPendingCreditCardChangeRequestData()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPendingEvents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPendingEvents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPendingInvoice", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPendingInvoice()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPendingInvoiceTopLevelItems", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPendingInvoiceTopLevelItems()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPendingInvoiceTotalAmount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPendingInvoiceTotalAmount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPendingInvoiceTotalOneTimeAmount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPendingInvoiceTotalOneTimeAmount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPendingInvoiceTotalOneTimeTaxAmount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPendingInvoiceTotalOneTimeTaxAmount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPendingInvoiceTotalRecurringAmount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPendingInvoiceTotalRecurringAmount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPendingInvoiceTotalRecurringTaxAmount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPendingInvoiceTotalRecurringTaxAmount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPermissionGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPermissionGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPermissionRoles", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPermissionRoles()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPlacementGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPlacementGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPortableStorageVolumes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPortableStorageVolumes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPostProvisioningHooks", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPostProvisioningHooks()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPptpVpnAllowedFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPptpVpnAllowedFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPptpVpnUsers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPptpVpnUsers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPreviousRecurringRevenue", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPreviousRecurringRevenue()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPriceRestrictions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPriceRestrictions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPriorityOneTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPriorityOneTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPrivateAllotmentHardwareBandwidthDetails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateAllotmentHardwareBandwidthDetails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPrivateBlockDeviceTemplateGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateBlockDeviceTemplateGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPrivateIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPrivateNetworkVlans", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateNetworkVlans()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPrivateSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getProofOfConceptAccountFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProofOfConceptAccountFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPublicAllotmentHardwareBandwidthDetails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicAllotmentHardwareBandwidthDetails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPublicIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPublicNetworkVlans", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicNetworkVlans()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getPublicSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getQuotes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetQuotes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getRecentEvents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRecentEvents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getReferralPartner", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReferralPartner()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getReferralPartnerCommissionForecast", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReferralPartnerCommissionForecast()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getReferralPartnerCommissionHistory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReferralPartnerCommissionHistory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getReferralPartnerCommissionPending", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReferralPartnerCommissionPending()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getReferredAccountFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReferredAccountFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getReferredAccounts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReferredAccounts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getRegulatedWorkloads", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRegulatedWorkloads()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getRemoteManagementCommandRequests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRemoteManagementCommandRequests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getReplicationEvents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReplicationEvents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getRequireSilentIBMidUserCreation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRequireSilentIBMidUserCreation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getReservedCapacityAgreements", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReservedCapacityAgreements()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getReservedCapacityGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReservedCapacityGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getRouters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRouters()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getRwhoisData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRwhoisData()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getSamlAuthentication", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSamlAuthentication()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getSecondaryDomains", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSecondaryDomains()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getSecurityCertificates", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSecurityCertificates()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getSecurityGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSecurityGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getSecurityLevel", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSecurityLevel()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getSecurityScanRequests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSecurityScanRequests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getServiceBillingItems", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceBillingItems()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getSharedBlockDeviceTemplateGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSharedBlockDeviceTemplateGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getShipments", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetShipments()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getSshKeys", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSshKeys()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getSslVpnUsers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSslVpnUsers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getStandardPoolVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStandardPoolVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getSubnetRegistrationDetails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubnetRegistrationDetails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getSubnetRegistrations", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubnetRegistrations()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getSupportRepresentatives", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSupportRepresentatives()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getSupportSubscriptions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSupportSubscriptions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getSupportTier", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSupportTier()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getSuppressInvoicesFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSuppressInvoicesFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getTags", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTags()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getTechIncubatorProgramInfo", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTechIncubatorProgramInfo(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getTestAccountAttributeFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTestAccountAttributeFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getThirdPartyPoliciesAcceptanceStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetThirdPartyPoliciesAcceptanceStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getTicketsClosedInTheLastThreeDays", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTicketsClosedInTheLastThreeDays()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getTicketsClosedToday", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTicketsClosedToday()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getUpgradeRequests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpgradeRequests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getUsers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUsers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getValidSecurityCertificateEntries", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetValidSecurityCertificateEntries()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getValidSecurityCertificates", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetValidSecurityCertificates()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getVdrUpdatesInProgressFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVdrUpdatesInProgressFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getVirtualDedicatedRacks", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualDedicatedRacks()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getVirtualDiskImages", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualDiskImages()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getVirtualGuestsOverBandwidthAllocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuestsOverBandwidthAllocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getVirtualGuestsProjectedOverBandwidthAllocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuestsProjectedOverBandwidthAllocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getVirtualGuestsWithCpanel", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuestsWithCpanel()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getVirtualGuestsWithMcafee", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuestsWithMcafee()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getVirtualGuestsWithMcafeeAntivirusRedhat", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuestsWithMcafeeAntivirusRedhat()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getVirtualGuestsWithMcafeeAntivirusWindows", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuestsWithMcafeeAntivirusWindows()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getVirtualGuestsWithMcafeeIntrusionDetectionSystem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuestsWithMcafeeIntrusionDetectionSystem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getVirtualGuestsWithPlesk", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuestsWithPlesk()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getVirtualGuestsWithQuantastor", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuestsWithQuantastor()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getVirtualGuestsWithUrchin", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuestsWithUrchin()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getVirtualPrivateRack", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualPrivateRack()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getVirtualStorageArchiveRepositories", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualStorageArchiveRepositories()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getVirtualStoragePublicRepositories", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualStoragePublicRepositories()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getVmWareActiveAccountLicenseKeys", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVmWareActiveAccountLicenseKeys()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getVpcVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVpcVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getVpnConfigRequiresVPNManageFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVpnConfigRequiresVPNManageFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::getWindowsUpdateStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetWindowsUpdateStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::hasAttribute", func() {
            It("API Call Test", func() {

                _, err := sl_service.HasAttribute(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::hourlyInstanceLimit", func() {
            It("API Call Test", func() {

                _, err := sl_service.HourlyInstanceLimit()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::hourlyServerLimit", func() {
            It("API Call Test", func() {

                _, err := sl_service.HourlyServerLimit()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::initiatePayerAuthentication", func() {
            It("API Call Test", func() {

                _, err := sl_service.InitiatePayerAuthentication(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::isActiveVmwareCustomer", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsActiveVmwareCustomer()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::isEligibleForLocalCurrencyProgram", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsEligibleForLocalCurrencyProgram()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::isEligibleToLinkWithPaas", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsEligibleToLinkWithPaas()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::linkExternalAccount", func() {
            It("API Call Test", func() {

				err := sl_service.LinkExternalAccount(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::removeAlternateCreditCard", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAlternateCreditCard()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::requestCreditCardChange", func() {
            It("API Call Test", func() {

                _, err := sl_service.RequestCreditCardChange(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::requestManualPayment", func() {
            It("API Call Test", func() {

                _, err := sl_service.RequestManualPayment(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::requestManualPaymentUsingCreditCardOnFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.RequestManualPaymentUsingCreditCardOnFile(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::saveInternalCostRecovery", func() {
            It("API Call Test", func() {

				err := sl_service.SaveInternalCostRecovery(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::setAbuseEmails", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetAbuseEmails(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::setManagedPoolQuantity", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetManagedPoolQuantity(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::setVlanSpan", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetVlanSpan(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::swapCreditCards", func() {
            It("API Call Test", func() {

                _, err := sl_service.SwapCreditCards()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::syncCurrentUserPopulationWithPaas", func() {
            It("API Call Test", func() {

				err := sl_service.SyncCurrentUserPopulationWithPaas()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::updateVpnUsersForResource", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateVpnUsersForResource(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::validate", func() {
            It("API Call Test", func() {

                _, err := sl_service.Validate(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account::validateManualPaymentAmount", func() {
            It("API Call Test", func() {

                _, err := sl_service.ValidateManualPaymentAmount(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Address service", func() {
        var sl_service services.Account_Address
        BeforeEach(func() {
            sl_service = services.GetAccountAddressService(slsession)
        })


        Context("SoftLayer_Account_Address::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Address::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Address::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Address::getAllDataCenters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllDataCenters()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Address::getCreateUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCreateUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Address::getLocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Address::getModifyEmployee", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetModifyEmployee()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Address::getModifyUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetModifyUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Address::getNetworkAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Address::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Address::getType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Address_Type service", func() {
        var sl_service services.Account_Address_Type
        BeforeEach(func() {
            sl_service = services.GetAccountAddressTypeService(slsession)
        })


        Context("SoftLayer_Account_Address_Type::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Affiliation service", func() {
        var sl_service services.Account_Affiliation
        BeforeEach(func() {
            sl_service = services.GetAccountAffiliationService(slsession)
        })


        Context("SoftLayer_Account_Affiliation::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Affiliation::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Affiliation::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Affiliation::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Affiliation::getAccountAffiliationsByAffiliateId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccountAffiliationsByAffiliateId(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Affiliation::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Agreement service", func() {
        var sl_service services.Account_Agreement
        BeforeEach(func() {
            sl_service = services.GetAccountAgreementService(slsession)
        })


        Context("SoftLayer_Account_Agreement::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Agreement::getAgreementType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAgreementType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Agreement::getAttachedBillingAgreementFiles", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachedBillingAgreementFiles()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Agreement::getBillingItems", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItems()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Agreement::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Agreement::getStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Agreement::getTopLevelBillingItems", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTopLevelBillingItems()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Authentication_Attribute service", func() {
        var sl_service services.Account_Authentication_Attribute
        BeforeEach(func() {
            sl_service = services.GetAccountAuthenticationAttributeService(slsession)
        })


        Context("SoftLayer_Account_Authentication_Attribute::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Authentication_Attribute::getAuthenticationRecord", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAuthenticationRecord()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Authentication_Attribute::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Authentication_Attribute::getType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Authentication_Attribute_Type service", func() {
        var sl_service services.Account_Authentication_Attribute_Type
        BeforeEach(func() {
            sl_service = services.GetAccountAuthenticationAttributeTypeService(slsession)
        })


        Context("SoftLayer_Account_Authentication_Attribute_Type::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Authentication_Attribute_Type::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Authentication_Saml service", func() {
        var sl_service services.Account_Authentication_Saml
        BeforeEach(func() {
            sl_service = services.GetAccountAuthenticationSamlService(slsession)
        })


        Context("SoftLayer_Account_Authentication_Saml::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Authentication_Saml::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Authentication_Saml::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Authentication_Saml::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Authentication_Saml::getAttributes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttributes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Authentication_Saml::getMetadata", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMetadata()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Authentication_Saml::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Brand_Migration_Request service", func() {
        var sl_service services.Account_Brand_Migration_Request
        BeforeEach(func() {
            sl_service = services.GetAccountBrandMigrationRequestService(slsession)
        })


        Context("SoftLayer_Account_Brand_Migration_Request::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Brand_Migration_Request::getDestinationBrand", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDestinationBrand()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Brand_Migration_Request::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Brand_Migration_Request::getSourceBrand", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSourceBrand()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Brand_Migration_Request::getUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Business_Partner service", func() {
        var sl_service services.Account_Business_Partner
        BeforeEach(func() {
            sl_service = services.GetAccountBusinessPartnerService(slsession)
        })


        Context("SoftLayer_Account_Business_Partner::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Business_Partner::getChannel", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetChannel()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Business_Partner::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Business_Partner::getSegment", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSegment()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Contact service", func() {
        var sl_service services.Account_Contact
        BeforeEach(func() {
            sl_service = services.GetAccountContactService(slsession)
        })


        Context("SoftLayer_Account_Contact::createComplianceReportRequestorContact", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateComplianceReportRequestorContact(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Contact::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Contact::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Contact::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Contact::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Contact::getAllContactTypes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllContactTypes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Contact::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Contact::getType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_External_Setup service", func() {
        var sl_service services.Account_External_Setup
        BeforeEach(func() {
            sl_service = services.GetAccountExternalSetupService(slsession)
        })


        Context("SoftLayer_Account_External_Setup::finalizeExternalBillingForAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.FinalizeExternalBillingForAccount(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_External_Setup::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_External_Setup::getVerifyCardTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVerifyCardTransaction()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Historical_Report service", func() {
        var sl_service services.Account_Historical_Report
        BeforeEach(func() {
            sl_service = services.GetAccountHistoricalReportService(slsession)
        })


        Context("SoftLayer_Account_Historical_Report::getAccountHostUptimeSummary", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccountHostUptimeSummary(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Historical_Report::getAccountUrlUptimeSummary", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccountUrlUptimeSummary(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Historical_Report::getHostUptimeDetail", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHostUptimeDetail(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Historical_Report::getHostUptimeGraphData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHostUptimeGraphData(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Historical_Report::getUrlUptimeDetail", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUrlUptimeDetail(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Historical_Report::getUrlUptimeGraphData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUrlUptimeGraphData(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Internal_Ibm service", func() {
        var sl_service services.Account_Internal_Ibm
        BeforeEach(func() {
            sl_service = services.GetAccountInternalIbmService(slsession)
        })


        Context("SoftLayer_Account_Internal_Ibm::getAccountTypes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccountTypes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Internal_Ibm::getAuthorizationUrl", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAuthorizationUrl(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Internal_Ibm::getBmsCountries", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBmsCountries()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Internal_Ibm::getBmsCountryList", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBmsCountryList()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Internal_Ibm::getEmployeeAccessToken", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEmployeeAccessToken(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Internal_Ibm::getManagerPreview", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetManagerPreview(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Internal_Ibm::hasExistingRequest", func() {
            It("API Call Test", func() {

                _, err := sl_service.HasExistingRequest(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Internal_Ibm::managerApprove", func() {
            It("API Call Test", func() {

				err := sl_service.ManagerApprove(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Internal_Ibm::managerDeny", func() {
            It("API Call Test", func() {

				err := sl_service.ManagerDeny(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Internal_Ibm::requestAccount", func() {
            It("API Call Test", func() {

				err := sl_service.RequestAccount(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Internal_Ibm_CostRecovery_Validator service", func() {
        var sl_service services.Account_Internal_Ibm_CostRecovery_Validator
        BeforeEach(func() {
            sl_service = services.GetAccountInternalIbmCostRecoveryValidatorService(slsession)
        })


        Context("SoftLayer_Account_Internal_Ibm_CostRecovery_Validator::getSprintContainer", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSprintContainer(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Internal_Ibm_CostRecovery_Validator::validateByAccountAndCountryId", func() {
            It("API Call Test", func() {

                _, err := sl_service.ValidateByAccountAndCountryId(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Link_Bluemix service", func() {
        var sl_service services.Account_Link_Bluemix
        BeforeEach(func() {
            sl_service = services.GetAccountLinkBluemixService(slsession)
        })


        Context("SoftLayer_Account_Link_Bluemix::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Link_Bluemix::getSupportTierType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSupportTierType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Link_OpenStack service", func() {
        var sl_service services.Account_Link_OpenStack
        BeforeEach(func() {
            sl_service = services.GetAccountLinkOpenStackService(slsession)
        })


        Context("SoftLayer_Account_Link_OpenStack::createOSDomain", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateOSDomain(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Link_OpenStack::createOSProject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateOSProject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Link_OpenStack::deleteOSDomain", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteOSDomain(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Link_OpenStack::deleteOSProject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteOSProject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Link_OpenStack::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Link_OpenStack::getOSProject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOSProject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Link_OpenStack::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Link_OpenStack::listOSProjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.ListOSProjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Lockdown_Request service", func() {
        var sl_service services.Account_Lockdown_Request
        BeforeEach(func() {
            sl_service = services.GetAccountLockdownRequestService(slsession)
        })


        Context("SoftLayer_Account_Lockdown_Request::cancelRequest", func() {
            It("API Call Test", func() {

				err := sl_service.CancelRequest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Lockdown_Request::disableLockedAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.DisableLockedAccount(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Lockdown_Request::disconnectCompute", func() {
            It("API Call Test", func() {

                _, err := sl_service.DisconnectCompute(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Lockdown_Request::getAccountHistory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccountHistory(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Lockdown_Request::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Lockdown_Request::reconnectCompute", func() {
            It("API Call Test", func() {

                _, err := sl_service.ReconnectCompute(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_MasterServiceAgreement service", func() {
        var sl_service services.Account_MasterServiceAgreement
        BeforeEach(func() {
            sl_service = services.GetAccountMasterServiceAgreementService(slsession)
        })


        Context("SoftLayer_Account_MasterServiceAgreement::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_MasterServiceAgreement::getFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFile()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_MasterServiceAgreement::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Media service", func() {
        var sl_service services.Account_Media
        BeforeEach(func() {
            sl_service = services.GetAccountMediaService(slsession)
        })


        Context("SoftLayer_Account_Media::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Media::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Media::getAllMediaTypes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllMediaTypes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Media::getCreateUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCreateUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Media::getDatacenter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDatacenter()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Media::getModifyEmployee", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetModifyEmployee()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Media::getModifyUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetModifyUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Media::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Media::getRequest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRequest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Media::getType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Media::getVolume", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVolume()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Media::removeMediaFromList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveMediaFromList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Media_Data_Transfer_Request service", func() {
        var sl_service services.Account_Media_Data_Transfer_Request
        BeforeEach(func() {
            sl_service = services.GetAccountMediaDataTransferRequestService(slsession)
        })


        Context("SoftLayer_Account_Media_Data_Transfer_Request::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Media_Data_Transfer_Request::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Media_Data_Transfer_Request::getActiveTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Media_Data_Transfer_Request::getAllRequestStatuses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllRequestStatuses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Media_Data_Transfer_Request::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Media_Data_Transfer_Request::getCreateUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCreateUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Media_Data_Transfer_Request::getMedia", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMedia()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Media_Data_Transfer_Request::getModifyEmployee", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetModifyEmployee()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Media_Data_Transfer_Request::getModifyUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetModifyUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Media_Data_Transfer_Request::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Media_Data_Transfer_Request::getShipments", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetShipments()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Media_Data_Transfer_Request::getStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Media_Data_Transfer_Request::getTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Note service", func() {
        var sl_service services.Account_Note
        BeforeEach(func() {
            sl_service = services.GetAccountNoteService(slsession)
        })


        Context("SoftLayer_Account_Note::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Note::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Note::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Note::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Note::getCustomer", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCustomer()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Note::getNoteHistory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNoteHistory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Note::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Partner_Referral_Prospect service", func() {
        var sl_service services.Account_Partner_Referral_Prospect
        BeforeEach(func() {
            sl_service = services.GetAccountPartnerReferralProspectService(slsession)
        })


        Context("SoftLayer_Account_Partner_Referral_Prospect::createProspect", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateProspect(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Partner_Referral_Prospect::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Partner_Referral_Prospect::getSurveyQuestions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSurveyQuestions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Password service", func() {
        var sl_service services.Account_Password
        BeforeEach(func() {
            sl_service = services.GetAccountPasswordService(slsession)
        })


        Context("SoftLayer_Account_Password::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Password::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Password::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Password::getType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_ProofOfConcept service", func() {
        var sl_service services.Account_ProofOfConcept
        BeforeEach(func() {
            sl_service = services.GetAccountProofOfConceptService(slsession)
        })


        Context("SoftLayer_Account_ProofOfConcept::approveReview", func() {
            It("API Call Test", func() {

				err := sl_service.ApproveReview(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_ProofOfConcept::denyReview", func() {
            It("API Call Test", func() {

				err := sl_service.DenyReview(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_ProofOfConcept::getAuthenticationUrl", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAuthenticationUrl(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_ProofOfConcept::getRequestsPendingIntegratedOfferingTeamReview", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRequestsPendingIntegratedOfferingTeamReview(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_ProofOfConcept::getRequestsPendingOverThresholdReview", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRequestsPendingOverThresholdReview(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_ProofOfConcept::getReviewerAccessToken", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReviewerAccessToken(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_ProofOfConcept::getReviewerEmailFromAccessToken", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReviewerEmailFromAccessToken(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_ProofOfConcept::getSubmittedRequest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubmittedRequest(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_ProofOfConcept::getSubmittedRequests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubmittedRequests(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_ProofOfConcept::getSupportEmailAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSupportEmailAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_ProofOfConcept::getTotalRequestsPendingIntegratedOfferingTeamReview", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTotalRequestsPendingIntegratedOfferingTeamReview(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_ProofOfConcept::getTotalRequestsPendingOverThresholdReviewCount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTotalRequestsPendingOverThresholdReviewCount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_ProofOfConcept::isCurrentReviewer", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsCurrentReviewer(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_ProofOfConcept::isIntegratedOfferingTeamReviewer", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsIntegratedOfferingTeamReviewer(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_ProofOfConcept::isOverThresholdReviewer", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsOverThresholdReviewer(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_ProofOfConcept::requestAccountTeamFundedAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.RequestAccountTeamFundedAccount(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_ProofOfConcept::requestGlobalFundedAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.RequestGlobalFundedAccount(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_ProofOfConcept::verifyReviewer", func() {
            It("API Call Test", func() {

				err := sl_service.VerifyReviewer(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_ProofOfConcept_Approver service", func() {
        var sl_service services.Account_ProofOfConcept_Approver
        BeforeEach(func() {
            sl_service = services.GetAccountProofOfConceptApproverService(slsession)
        })


        Context("SoftLayer_Account_ProofOfConcept_Approver::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_ProofOfConcept_Approver::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_ProofOfConcept_Approver::getRole", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRole()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_ProofOfConcept_Approver::getType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_ProofOfConcept_Approver_Role service", func() {
        var sl_service services.Account_ProofOfConcept_Approver_Role
        BeforeEach(func() {
            sl_service = services.GetAccountProofOfConceptApproverRoleService(slsession)
        })


        Context("SoftLayer_Account_ProofOfConcept_Approver_Role::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_ProofOfConcept_Approver_Type service", func() {
        var sl_service services.Account_ProofOfConcept_Approver_Type
        BeforeEach(func() {
            sl_service = services.GetAccountProofOfConceptApproverTypeService(slsession)
        })


        Context("SoftLayer_Account_ProofOfConcept_Approver_Type::getApprovers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetApprovers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_ProofOfConcept_Approver_Type::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_ProofOfConcept_Campaign_Code service", func() {
        var sl_service services.Account_ProofOfConcept_Campaign_Code
        BeforeEach(func() {
            sl_service = services.GetAccountProofOfConceptCampaignCodeService(slsession)
        })


        Context("SoftLayer_Account_ProofOfConcept_Campaign_Code::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_ProofOfConcept_Campaign_Code::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_ProofOfConcept_Funding_Type service", func() {
        var sl_service services.Account_ProofOfConcept_Funding_Type
        BeforeEach(func() {
            sl_service = services.GetAccountProofOfConceptFundingTypeService(slsession)
        })


        Context("SoftLayer_Account_ProofOfConcept_Funding_Type::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_ProofOfConcept_Funding_Type::getApproverTypes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetApproverTypes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_ProofOfConcept_Funding_Type::getApprovers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetApprovers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_ProofOfConcept_Funding_Type::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Regional_Registry_Detail service", func() {
        var sl_service services.Account_Regional_Registry_Detail
        BeforeEach(func() {
            sl_service = services.GetAccountRegionalRegistryDetailService(slsession)
        })


        Context("SoftLayer_Account_Regional_Registry_Detail::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Regional_Registry_Detail::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Regional_Registry_Detail::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Regional_Registry_Detail::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Regional_Registry_Detail::getDetailType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDetailType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Regional_Registry_Detail::getDetails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDetails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Regional_Registry_Detail::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Regional_Registry_Detail::getProperties", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProperties()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Regional_Registry_Detail::getRegionalInternetRegistryHandle", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRegionalInternetRegistryHandle()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Regional_Registry_Detail::updateReferencedRegistrations", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateReferencedRegistrations()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Regional_Registry_Detail::validatePersonForAllRegistrars", func() {
            It("API Call Test", func() {

                _, err := sl_service.ValidatePersonForAllRegistrars()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Regional_Registry_Detail_Property service", func() {
        var sl_service services.Account_Regional_Registry_Detail_Property
        BeforeEach(func() {
            sl_service = services.GetAccountRegionalRegistryDetailPropertyService(slsession)
        })


        Context("SoftLayer_Account_Regional_Registry_Detail_Property::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Regional_Registry_Detail_Property::createObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Regional_Registry_Detail_Property::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Regional_Registry_Detail_Property::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Regional_Registry_Detail_Property::editObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Regional_Registry_Detail_Property::getDetail", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDetail()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Regional_Registry_Detail_Property::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Regional_Registry_Detail_Property::getPropertyType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPropertyType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Regional_Registry_Detail_Property_Type service", func() {
        var sl_service services.Account_Regional_Registry_Detail_Property_Type
        BeforeEach(func() {
            sl_service = services.GetAccountRegionalRegistryDetailPropertyTypeService(slsession)
        })


        Context("SoftLayer_Account_Regional_Registry_Detail_Property_Type::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Regional_Registry_Detail_Property_Type::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Regional_Registry_Detail_Type service", func() {
        var sl_service services.Account_Regional_Registry_Detail_Type
        BeforeEach(func() {
            sl_service = services.GetAccountRegionalRegistryDetailTypeService(slsession)
        })


        Context("SoftLayer_Account_Regional_Registry_Detail_Type::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Regional_Registry_Detail_Type::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Reports_Request service", func() {
        var sl_service services.Account_Reports_Request
        BeforeEach(func() {
            sl_service = services.GetAccountReportsRequestService(slsession)
        })


        Context("SoftLayer_Account_Reports_Request::createRequest", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateRequest(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Reports_Request::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Reports_Request::getAccountContact", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccountContact()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Reports_Request::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Reports_Request::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Reports_Request::getReportType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReportType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Reports_Request::getRequestByRequestKey", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRequestByRequestKey(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Reports_Request::getRequestorContact", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRequestorContact()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Reports_Request::getTicket", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTicket()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Reports_Request::getUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Reports_Request::sendReportEmail", func() {
            It("API Call Test", func() {

                _, err := sl_service.SendReportEmail(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Reports_Request::updateTicketOnDecline", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateTicketOnDecline(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Shipment service", func() {
        var sl_service services.Account_Shipment
        BeforeEach(func() {
            sl_service = services.GetAccountShipmentService(slsession)
        })


        Context("SoftLayer_Account_Shipment::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment::getAllCouriers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllCouriers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment::getAllCouriersByType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllCouriersByType(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment::getAllShipmentStatuses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllShipmentStatuses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment::getAllShipmentTypes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllShipmentTypes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment::getCourier", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCourier()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment::getCreateEmployee", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCreateEmployee()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment::getCreateUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCreateUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment::getCurrency", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrency()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment::getDestinationAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDestinationAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment::getMasterTrackingData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMasterTrackingData()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment::getModifyEmployee", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetModifyEmployee()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment::getModifyUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetModifyUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment::getOriginationAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOriginationAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment::getShipmentItems", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetShipmentItems()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment::getStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment::getTrackingData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTrackingData()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment::getType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment::getViaAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetViaAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Shipment_Item service", func() {
        var sl_service services.Account_Shipment_Item
        BeforeEach(func() {
            sl_service = services.GetAccountShipmentItemService(slsession)
        })


        Context("SoftLayer_Account_Shipment_Item::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment_Item::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment_Item::getShipment", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetShipment()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment_Item::getShipmentItemType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetShipmentItemType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Shipment_Item_Type service", func() {
        var sl_service services.Account_Shipment_Item_Type
        BeforeEach(func() {
            sl_service = services.GetAccountShipmentItemTypeService(slsession)
        })


        Context("SoftLayer_Account_Shipment_Item_Type::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Shipment_Resource_Type service", func() {
        var sl_service services.Account_Shipment_Resource_Type
        BeforeEach(func() {
            sl_service = services.GetAccountShipmentResourceTypeService(slsession)
        })


        Context("SoftLayer_Account_Shipment_Resource_Type::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Shipment_Status service", func() {
        var sl_service services.Account_Shipment_Status
        BeforeEach(func() {
            sl_service = services.GetAccountShipmentStatusService(slsession)
        })


        Context("SoftLayer_Account_Shipment_Status::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Shipment_Tracking_Data service", func() {
        var sl_service services.Account_Shipment_Tracking_Data
        BeforeEach(func() {
            sl_service = services.GetAccountShipmentTrackingDataService(slsession)
        })


        Context("SoftLayer_Account_Shipment_Tracking_Data::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment_Tracking_Data::createObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment_Tracking_Data::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment_Tracking_Data::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment_Tracking_Data::getCreateEmployee", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCreateEmployee()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment_Tracking_Data::getCreateUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCreateUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment_Tracking_Data::getModifyEmployee", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetModifyEmployee()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment_Tracking_Data::getModifyUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetModifyUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment_Tracking_Data::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Account_Shipment_Tracking_Data::getShipment", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetShipment()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Account_Shipment_Type service", func() {
        var sl_service services.Account_Shipment_Type
        BeforeEach(func() {
            sl_service = services.GetAccountShipmentTypeService(slsession)
        })


        Context("SoftLayer_Account_Shipment_Type::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

})
