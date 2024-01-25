
package services_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/softlayer/softlayer-go/session/sessionfakes"
    "github.com/softlayer/softlayer-go/services"
)   

var _ = Describe("Metric Tests", func() {
    var slsession *sessionfakes.FakeSLSession
    BeforeEach(func() {
        slsession = &sessionfakes.FakeSLSession{}
    })

    Context("Testing SoftLayer_Metric_Tracking_Object service", func() {
        var sl_service services.Metric_Tracking_Object
        BeforeEach(func() {
            sl_service = services.GetMetricTrackingObjectService(slsession)
        })


        Context("SoftLayer_Metric_Tracking_Object::getBandwidthData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthData(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Metric_Tracking_Object::getBandwidthGraph", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthGraph(nil,nil,nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Metric_Tracking_Object::getBandwidthTotal", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthTotal(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Metric_Tracking_Object::getDetailsForDateRange", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDetailsForDateRange(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Metric_Tracking_Object::getGraph", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGraph(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Metric_Tracking_Object::getMetricDataTypes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMetricDataTypes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Metric_Tracking_Object::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Metric_Tracking_Object::getSummary", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSummary(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Metric_Tracking_Object::getSummaryData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSummaryData(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Metric_Tracking_Object::getType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Metric_Tracking_Object_Bandwidth_Summary service", func() {
        var sl_service services.Metric_Tracking_Object_Bandwidth_Summary
        BeforeEach(func() {
            sl_service = services.GetMetricTrackingObjectBandwidthSummaryService(slsession)
        })


        Context("SoftLayer_Metric_Tracking_Object_Bandwidth_Summary::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

})
