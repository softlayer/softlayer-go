
package services_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/softlayer/softlayer-go/session/sessionfakes"
    "github.com/softlayer/softlayer-go/services"
)   

var _ = Describe("Email Tests", func() {
    var slsession *sessionfakes.FakeSLSession
    BeforeEach(func() {
        slsession = &sessionfakes.FakeSLSession{}
    })

    Context("Testing SoftLayer_Email_Subscription service", func() {
        var sl_service services.Email_Subscription
        BeforeEach(func() {
            sl_service = services.GetEmailSubscriptionService(slsession)
        })


        Context("SoftLayer_Email_Subscription::disable", func() {
            It("API Call Test", func() {

                _, err := sl_service.Disable()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Email_Subscription::enable", func() {
            It("API Call Test", func() {

                _, err := sl_service.Enable()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Email_Subscription::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Email_Subscription::getEnabled", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEnabled()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Email_Subscription::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Email_Subscription_Group service", func() {
        var sl_service services.Email_Subscription_Group
        BeforeEach(func() {
            sl_service = services.GetEmailSubscriptionGroupService(slsession)
        })


        Context("SoftLayer_Email_Subscription_Group::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Email_Subscription_Group::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Email_Subscription_Group::getSubscriptions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubscriptions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

})
