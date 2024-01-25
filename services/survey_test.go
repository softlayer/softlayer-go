
package services_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/softlayer/softlayer-go/session/sessionfakes"
    "github.com/softlayer/softlayer-go/services"
)   

var _ = Describe("Survey Tests", func() {
    var slsession *sessionfakes.FakeSLSession
    BeforeEach(func() {
        slsession = &sessionfakes.FakeSLSession{}
    })

    Context("Testing SoftLayer_Survey service", func() {
        var sl_service services.Survey
        BeforeEach(func() {
            sl_service = services.GetSurveyService(slsession)
        })


        Context("SoftLayer_Survey::getActiveSurveyByType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveSurveyByType(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Survey::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Survey::getQuestions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetQuestions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Survey::getStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Survey::getType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Survey::takeSurvey", func() {
            It("API Call Test", func() {

                _, err := sl_service.TakeSurvey(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

})
