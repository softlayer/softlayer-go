// AUTO GENERATED by tools/loadmeta.go
package services_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/session/sessionfakes"
)

var _ = Describe("Sales Tests", func() {
	var slsession *sessionfakes.FakeSLSession
	BeforeEach(func() {
		slsession = &sessionfakes.FakeSLSession{}
	})

	Context("Testing SoftLayer_Sales_Presale_Event service", func() {
		var sl_service services.Sales_Presale_Event
		BeforeEach(func() {
			sl_service = services.GetSalesPresaleEventService(slsession)
		})
		Context("SoftLayer_Sales_Presale_Event::getActiveFlag", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetActiveFlag()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Sales_Presale_Event::getAllObjects", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetAllObjects()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Sales_Presale_Event::getExpiredFlag", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetExpiredFlag()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Sales_Presale_Event::getItem", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetItem()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Sales_Presale_Event::getLocation", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetLocation()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Sales_Presale_Event::getObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetObject()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Sales_Presale_Event::getOrders", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetOrders()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

})
