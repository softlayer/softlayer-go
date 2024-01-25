
package services_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/softlayer/softlayer-go/session/sessionfakes"
    "github.com/softlayer/softlayer-go/services"
)   

var _ = Describe("FlexibleCredit Tests", func() {
    var slsession *sessionfakes.FakeSLSession
    BeforeEach(func() {
        slsession = &sessionfakes.FakeSLSession{}
    })

    Context("Testing SoftLayer_FlexibleCredit_Program service", func() {
        var sl_service services.FlexibleCredit_Program
        BeforeEach(func() {
            sl_service = services.GetFlexibleCreditProgramService(slsession)
        })


        Context("SoftLayer_FlexibleCredit_Program::getAffiliatesAvailableForSelfEnrollmentByVerificationType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAffiliatesAvailableForSelfEnrollmentByVerificationType(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_FlexibleCredit_Program::getCompanyTypes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCompanyTypes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_FlexibleCredit_Program::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_FlexibleCredit_Program::selfEnrollNewAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.SelfEnrollNewAccount(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

})
