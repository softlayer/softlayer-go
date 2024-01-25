
package services_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/softlayer/softlayer-go/session/sessionfakes"
    "github.com/softlayer/softlayer-go/services"
)   

var _ = Describe("Business Tests", func() {
    var slsession *sessionfakes.FakeSLSession
    BeforeEach(func() {
        slsession = &sessionfakes.FakeSLSession{}
    })

    Context("Testing SoftLayer_Business_Partner_Channel service", func() {
        var sl_service services.Business_Partner_Channel
        BeforeEach(func() {
            sl_service = services.GetBusinessPartnerChannelService(slsession)
        })


        Context("SoftLayer_Business_Partner_Channel::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Business_Partner_Segment service", func() {
        var sl_service services.Business_Partner_Segment
        BeforeEach(func() {
            sl_service = services.GetBusinessPartnerSegmentService(slsession)
        })


        Context("SoftLayer_Business_Partner_Segment::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

})
