// AUTO GENERATED by tools/loadmeta.go
package services_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/session/sessionfakes"
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
		Context("SoftLayer_Survey Set Options", func() {
			It("Set Options properly", func() {
				t_id := 1234
				t_filter := "{'testFilter':{'test'}}"
				t_limit := 100
				t_offset := 5
				sl_service = sl_service.Id(t_id).Filter(t_filter).Offset(t_offset).Limit(t_limit)
				Expect(sl_service.Options.Id).To(HaveValue(Equal(t_id)))
				Expect(sl_service.Options.Filter).To(HaveValue(Equal(t_filter)))
				Expect(sl_service.Options.Limit).To(HaveValue(Equal(t_limit)))
				Expect(sl_service.Options.Offset).To(HaveValue(Equal(t_offset)))
			})
		})
		Context("SoftLayer_Survey Set Mask", func() {
			It("Set Options properly", func() {
				t_mask1 := "mask[test,test2]"
				sl_service = sl_service.Mask(t_mask1)
				Expect(sl_service.Options.Mask).To(HaveValue(Equal(t_mask1)))
				// Mask("test,test2") should set the mask to be "mask[test,test2]" aka t_mask1
				sl_service = sl_service.Mask("test,test2")
				Expect(sl_service.Options.Mask).To(HaveValue(Equal(t_mask1)))
			})
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
