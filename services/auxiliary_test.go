
package services_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/softlayer/softlayer-go/session/sessionfakes"
    "github.com/softlayer/softlayer-go/services"
)   

var _ = Describe("Auxiliary Tests", func() {
    var slsession *sessionfakes.FakeSLSession
    BeforeEach(func() {
        slsession = &sessionfakes.FakeSLSession{}
    })

    Context("Testing SoftLayer_Auxiliary_Network_Status service", func() {
        var sl_service services.Auxiliary_Network_Status
        BeforeEach(func() {
            sl_service = services.GetAuxiliaryNetworkStatusService(slsession)
        })


        Context("SoftLayer_Auxiliary_Network_Status::getNetworkStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkStatus(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Auxiliary_Notification_Emergency service", func() {
        var sl_service services.Auxiliary_Notification_Emergency
        BeforeEach(func() {
            sl_service = services.GetAuxiliaryNotificationEmergencyService(slsession)
        })


        Context("SoftLayer_Auxiliary_Notification_Emergency::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Auxiliary_Notification_Emergency::getCurrentNotifications", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentNotifications()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Auxiliary_Notification_Emergency::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Auxiliary_Notification_Emergency::getSignature", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSignature()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Auxiliary_Notification_Emergency::getStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Auxiliary_Shipping_Courier_Type service", func() {
        var sl_service services.Auxiliary_Shipping_Courier_Type
        BeforeEach(func() {
            sl_service = services.GetAuxiliaryShippingCourierTypeService(slsession)
        })


        Context("SoftLayer_Auxiliary_Shipping_Courier_Type::getCourier", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCourier()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Auxiliary_Shipping_Courier_Type::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Auxiliary_Shipping_Courier_Type::getTypeByKeyName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTypeByKeyName(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

})
