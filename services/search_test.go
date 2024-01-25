
package services_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/softlayer/softlayer-go/session/sessionfakes"
    "github.com/softlayer/softlayer-go/services"
)   

var _ = Describe("Search Tests", func() {
    var slsession *sessionfakes.FakeSLSession
    BeforeEach(func() {
        slsession = &sessionfakes.FakeSLSession{}
    })

    Context("Testing SoftLayer_Search service", func() {
        var sl_service services.Search
        BeforeEach(func() {
            sl_service = services.GetSearchService(slsession)
        })


        Context("SoftLayer_Search::advancedSearch", func() {
            It("API Call Test", func() {

                _, err := sl_service.AdvancedSearch(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Search::getObjectTypes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObjectTypes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Search::search", func() {
            It("API Call Test", func() {

                _, err := sl_service.Search(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

})
