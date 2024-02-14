// AUTO GENERATED by tools/loadmeta.go
package services_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/session/sessionfakes"
)

var _ = Describe("Billing Tests", func() {
	var slsession *sessionfakes.FakeSLSession
	BeforeEach(func() {
		slsession = &sessionfakes.FakeSLSession{}
	})

	Context("Testing SoftLayer_Billing_Invoice service", func() {
		var sl_service services.Billing_Invoice
		BeforeEach(func() {
			sl_service = services.GetBillingInvoiceService(slsession)
		})
		Context("SoftLayer_Billing_Invoice Set Options", func() {
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
		Context("SoftLayer_Billing_Invoice Set Mask", func() {
			It("Set Options properly", func() {
				t_mask1 := "mask[test,test2]"
				sl_service = sl_service.Mask(t_mask1)
				Expect(sl_service.Options.Mask).To(HaveValue(Equal(t_mask1)))
				// Mask("test,test2") should set the mask to be "mask[test,test2]" aka t_mask1
				sl_service = sl_service.Mask("test,test2")
				Expect(sl_service.Options.Mask).To(HaveValue(Equal(t_mask1)))
			})
		})
		Context("SoftLayer_Billing_Invoice::getInvoiceTopLevelItems", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetInvoiceTopLevelItems()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

	Context("Testing SoftLayer_Billing_Item service", func() {
		var sl_service services.Billing_Item
		BeforeEach(func() {
			sl_service = services.GetBillingItemService(slsession)
		})
		Context("SoftLayer_Billing_Item Set Options", func() {
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
		Context("SoftLayer_Billing_Item Set Mask", func() {
			It("Set Options properly", func() {
				t_mask1 := "mask[test,test2]"
				sl_service = sl_service.Mask(t_mask1)
				Expect(sl_service.Options.Mask).To(HaveValue(Equal(t_mask1)))
				// Mask("test,test2") should set the mask to be "mask[test,test2]" aka t_mask1
				sl_service = sl_service.Mask("test,test2")
				Expect(sl_service.Options.Mask).To(HaveValue(Equal(t_mask1)))
			})
		})
		Context("SoftLayer_Billing_Item::cancelItem", func() {
			It("API Call Test", func() {
				_, err := sl_service.CancelItem(nil, nil, nil, nil)
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Billing_Item::cancelService", func() {
			It("API Call Test", func() {
				_, err := sl_service.CancelService()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Billing_Item::getObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetObject()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

	Context("Testing SoftLayer_Billing_Item_Cancellation_Request service", func() {
		var sl_service services.Billing_Item_Cancellation_Request
		BeforeEach(func() {
			sl_service = services.GetBillingItemCancellationRequestService(slsession)
		})
		Context("SoftLayer_Billing_Item_Cancellation_Request Set Options", func() {
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
		Context("SoftLayer_Billing_Item_Cancellation_Request Set Mask", func() {
			It("Set Options properly", func() {
				t_mask1 := "mask[test,test2]"
				sl_service = sl_service.Mask(t_mask1)
				Expect(sl_service.Options.Mask).To(HaveValue(Equal(t_mask1)))
				// Mask("test,test2") should set the mask to be "mask[test,test2]" aka t_mask1
				sl_service = sl_service.Mask("test,test2")
				Expect(sl_service.Options.Mask).To(HaveValue(Equal(t_mask1)))
			})
		})
		Context("SoftLayer_Billing_Item_Cancellation_Request::getAllCancellationRequests", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetAllCancellationRequests()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

	Context("Testing SoftLayer_Billing_Order service", func() {
		var sl_service services.Billing_Order
		BeforeEach(func() {
			sl_service = services.GetBillingOrderService(slsession)
		})
		Context("SoftLayer_Billing_Order Set Options", func() {
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
		Context("SoftLayer_Billing_Order Set Mask", func() {
			It("Set Options properly", func() {
				t_mask1 := "mask[test,test2]"
				sl_service = sl_service.Mask(t_mask1)
				Expect(sl_service.Options.Mask).To(HaveValue(Equal(t_mask1)))
				// Mask("test,test2") should set the mask to be "mask[test,test2]" aka t_mask1
				sl_service = sl_service.Mask("test,test2")
				Expect(sl_service.Options.Mask).To(HaveValue(Equal(t_mask1)))
			})
		})
		Context("SoftLayer_Billing_Order::getAllObjects", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetAllObjects()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Billing_Order::getObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetObject()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

	Context("Testing SoftLayer_Billing_Order_Quote service", func() {
		var sl_service services.Billing_Order_Quote
		BeforeEach(func() {
			sl_service = services.GetBillingOrderQuoteService(slsession)
		})
		Context("SoftLayer_Billing_Order_Quote Set Options", func() {
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
		Context("SoftLayer_Billing_Order_Quote Set Mask", func() {
			It("Set Options properly", func() {
				t_mask1 := "mask[test,test2]"
				sl_service = sl_service.Mask(t_mask1)
				Expect(sl_service.Options.Mask).To(HaveValue(Equal(t_mask1)))
				// Mask("test,test2") should set the mask to be "mask[test,test2]" aka t_mask1
				sl_service = sl_service.Mask("test,test2")
				Expect(sl_service.Options.Mask).To(HaveValue(Equal(t_mask1)))
			})
		})
		Context("SoftLayer_Billing_Order_Quote::deleteQuote", func() {
			It("API Call Test", func() {
				_, err := sl_service.DeleteQuote()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Billing_Order_Quote::getObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetObject()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Billing_Order_Quote::getRecalculatedOrderContainer", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetRecalculatedOrderContainer(GetOrderContainer(), nil)
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Billing_Order_Quote::placeOrder", func() {
			It("API Call Test", func() {
				_, err := sl_service.PlaceOrder(GetOrderContainer())
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Billing_Order_Quote::saveQuote", func() {
			It("API Call Test", func() {
				_, err := sl_service.SaveQuote()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Billing_Order_Quote::verifyOrder", func() {
			It("API Call Test", func() {
				_, err := sl_service.VerifyOrder(GetOrderContainer())
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

})
