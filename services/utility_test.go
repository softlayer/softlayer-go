
package services_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/softlayer/softlayer-go/session/sessionfakes"
    "github.com/softlayer/softlayer-go/services"
)   

var _ = Describe("Utility Tests", func() {
    var slsession *sessionfakes.FakeSLSession
    BeforeEach(func() {
        slsession = &sessionfakes.FakeSLSession{}
    })

    Context("Testing SoftLayer_Utility_Network service", func() {
        var sl_service services.Utility_Network
        BeforeEach(func() {
            sl_service = services.GetUtilityNetworkService(slsession)
        })


        Context("SoftLayer_Utility_Network::nsLookup", func() {
            It("API Call Test", func() {

                _, err := sl_service.NsLookup(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Utility_Network::whois", func() {
            It("API Call Test", func() {

                _, err := sl_service.Whois(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

})
