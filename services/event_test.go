// AUTO GENERATED by tools/loadmeta.go
package services_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/session/sessionfakes"
)   

var _ = Describe("Event Tests", func() {
	var slsession *sessionfakes.FakeSLSession
	BeforeEach(func() {
		slsession = &sessionfakes.FakeSLSession{}
	})

	Context("Testing SoftLayer_Event_Log service", func() {
		var sl_service services.Event_Log
		BeforeEach(func() {
			sl_service = services.GetEventLogService(slsession)
		})
		Context("SoftLayer_Event_Log::getAllEventNames", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetAllEventNames(nil)
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Event_Log::getAllEventObjectNames", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetAllEventObjectNames()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Event_Log::getAllObjects", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetAllObjects()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Event_Log::getAllUserTypes", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetAllUserTypes()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Event_Log::getUser", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetUser()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

})

