// AUTO GENERATED by tools/loadmeta.go
package services_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/session/sessionfakes"
)

var _ = Describe("Provisioning Tests", func() {
	var slsession *sessionfakes.FakeSLSession
	BeforeEach(func() {
		slsession = &sessionfakes.FakeSLSession{}
	})

	Context("Testing SoftLayer_Provisioning_Hook service", func() {
		var sl_service services.Provisioning_Hook
		BeforeEach(func() {
			sl_service = services.GetProvisioningHookService(slsession)
		})
		Context("SoftLayer_Provisioning_Hook::createObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.CreateObject(nil)
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Provisioning_Hook::deleteObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.DeleteObject()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Provisioning_Hook::editObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.EditObject(nil)
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Provisioning_Hook::getAccount", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetAccount()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Provisioning_Hook::getHookType", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetHookType()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Provisioning_Hook::getObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetObject()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

	Context("Testing SoftLayer_Provisioning_Hook_Type service", func() {
		var sl_service services.Provisioning_Hook_Type
		BeforeEach(func() {
			sl_service = services.GetProvisioningHookTypeService(slsession)
		})
		Context("SoftLayer_Provisioning_Hook_Type::getAllHookTypes", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetAllHookTypes()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Provisioning_Hook_Type::getObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetObject()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

	Context("Testing SoftLayer_Provisioning_Maintenance_Classification service", func() {
		var sl_service services.Provisioning_Maintenance_Classification
		BeforeEach(func() {
			sl_service = services.GetProvisioningMaintenanceClassificationService(slsession)
		})
		Context("SoftLayer_Provisioning_Maintenance_Classification::getItemCategories", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetItemCategories()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Provisioning_Maintenance_Classification::getMaintenanceClassification", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetMaintenanceClassification(nil)
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Provisioning_Maintenance_Classification::getMaintenanceClassificationsByItemCategory", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetMaintenanceClassificationsByItemCategory()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Provisioning_Maintenance_Classification::getObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetObject()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

	Context("Testing SoftLayer_Provisioning_Maintenance_Classification_Item_Category service", func() {
		var sl_service services.Provisioning_Maintenance_Classification_Item_Category
		BeforeEach(func() {
			sl_service = services.GetProvisioningMaintenanceClassificationItemCategoryService(slsession)
		})
		Context("SoftLayer_Provisioning_Maintenance_Classification_Item_Category::getMaintenanceClassification", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetMaintenanceClassification()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Provisioning_Maintenance_Classification_Item_Category::getObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetObject()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

	Context("Testing SoftLayer_Provisioning_Maintenance_Slots service", func() {
		var sl_service services.Provisioning_Maintenance_Slots
		BeforeEach(func() {
			sl_service = services.GetProvisioningMaintenanceSlotsService(slsession)
		})
		Context("SoftLayer_Provisioning_Maintenance_Slots::getObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetObject()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

	Context("Testing SoftLayer_Provisioning_Maintenance_Ticket service", func() {
		var sl_service services.Provisioning_Maintenance_Ticket
		BeforeEach(func() {
			sl_service = services.GetProvisioningMaintenanceTicketService(slsession)
		})
		Context("SoftLayer_Provisioning_Maintenance_Ticket::getAvailableSlots", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetAvailableSlots()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Provisioning_Maintenance_Ticket::getMaintenanceClass", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetMaintenanceClass()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Provisioning_Maintenance_Ticket::getObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetObject()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Provisioning_Maintenance_Ticket::getTicket", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetTicket()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

	Context("Testing SoftLayer_Provisioning_Maintenance_Window service", func() {
		var sl_service services.Provisioning_Maintenance_Window
		BeforeEach(func() {
			sl_service = services.GetProvisioningMaintenanceWindowService(slsession)
		})
		Context("SoftLayer_Provisioning_Maintenance_Window::addCustomerUpgradeWindow", func() {
			It("API Call Test", func() {
				_, err := sl_service.AddCustomerUpgradeWindow(nil)
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Provisioning_Maintenance_Window::getMaintenanceClassifications", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetMaintenanceClassifications()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Provisioning_Maintenance_Window::getMaintenanceStartEndTime", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetMaintenanceStartEndTime(nil)
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Provisioning_Maintenance_Window::getMaintenanceWindowForTicket", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetMaintenanceWindowForTicket(nil)
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Provisioning_Maintenance_Window::getMaintenanceWindowTicketsByTicketId", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetMaintenanceWindowTicketsByTicketId(nil)
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Provisioning_Maintenance_Window::getMaintenanceWindows", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetMaintenanceWindows(nil, nil, nil, nil)
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Provisioning_Maintenance_Window::getMaintenceWindows", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetMaintenceWindows(nil, nil, nil, nil)
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

	Context("Testing SoftLayer_Provisioning_Version1_Transaction_Group service", func() {
		var sl_service services.Provisioning_Version1_Transaction_Group
		BeforeEach(func() {
			sl_service = services.GetProvisioningVersion1TransactionGroupService(slsession)
		})
		Context("SoftLayer_Provisioning_Version1_Transaction_Group::getAllObjects", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetAllObjects()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Provisioning_Version1_Transaction_Group::getObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetObject()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

	Context("Testing SoftLayer_Provisioning_Version1_Transaction_OrderTracking service", func() {
		var sl_service services.Provisioning_Version1_Transaction_OrderTracking
		BeforeEach(func() {
			sl_service = services.GetProvisioningVersion1TransactionOrderTrackingService(slsession)
		})
		Context("SoftLayer_Provisioning_Version1_Transaction_OrderTracking::getInvoiceId", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetInvoiceId()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Provisioning_Version1_Transaction_OrderTracking::getObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetObject()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Provisioning_Version1_Transaction_OrderTracking::getOrderTrackingState", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetOrderTrackingState()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
		Context("SoftLayer_Provisioning_Version1_Transaction_OrderTracking::getTransaction", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetTransaction()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

	Context("Testing SoftLayer_Provisioning_Version1_Transaction_OrderTrackingState service", func() {
		var sl_service services.Provisioning_Version1_Transaction_OrderTrackingState
		BeforeEach(func() {
			sl_service = services.GetProvisioningVersion1TransactionOrderTrackingStateService(slsession)
		})
		Context("SoftLayer_Provisioning_Version1_Transaction_OrderTrackingState::getObject", func() {
			It("API Call Test", func() {
				_, err := sl_service.GetObject()
				Expect(err).To(Succeed())
				Expect(slsession.DoRequestCallCount()).To(Equal(1))
			})
		})
	})

})
