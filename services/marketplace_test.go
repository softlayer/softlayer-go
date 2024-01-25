
package services_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/softlayer/softlayer-go/session/sessionfakes"
    "github.com/softlayer/softlayer-go/services"
)   

var _ = Describe("Marketplace Tests", func() {
    var slsession *sessionfakes.FakeSLSession
    BeforeEach(func() {
        slsession = &sessionfakes.FakeSLSession{}
    })

    Context("Testing SoftLayer_Marketplace_Partner service", func() {
        var sl_service services.Marketplace_Partner
        BeforeEach(func() {
            sl_service = services.GetMarketplacePartnerService(slsession)
        })


        Context("SoftLayer_Marketplace_Partner::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Marketplace_Partner::getAllPublishedPartners", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllPublishedPartners(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Marketplace_Partner::getAttachments", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachments()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Marketplace_Partner::getFeaturedPartners", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFeaturedPartners(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Marketplace_Partner::getFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFile(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Marketplace_Partner::getLogoMedium", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLogoMedium()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Marketplace_Partner::getLogoMediumTemp", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLogoMediumTemp()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Marketplace_Partner::getLogoSmall", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLogoSmall()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Marketplace_Partner::getLogoSmallTemp", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLogoSmallTemp()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Marketplace_Partner::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Marketplace_Partner::getPartnerByUrlIdentifier", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPartnerByUrlIdentifier(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

})
