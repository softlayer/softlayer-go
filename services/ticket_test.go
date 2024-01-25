
package services_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/softlayer/softlayer-go/session/sessionfakes"
    "github.com/softlayer/softlayer-go/services"
)   

var _ = Describe("Ticket Tests", func() {
    var slsession *sessionfakes.FakeSLSession
    BeforeEach(func() {
        slsession = &sessionfakes.FakeSLSession{}
    })

    Context("Testing SoftLayer_Ticket service", func() {
        var sl_service services.Ticket
        BeforeEach(func() {
            sl_service = services.GetTicketService(slsession)
        })


        Context("SoftLayer_Ticket::addAssignedAgent", func() {
            It("API Call Test", func() {

				err := sl_service.AddAssignedAgent(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::addAttachedAdditionalEmails", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddAttachedAdditionalEmails(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::addAttachedDedicatedHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddAttachedDedicatedHost(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::addAttachedFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddAttachedFile(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::addAttachedHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddAttachedHardware(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::addAttachedVirtualGuest", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddAttachedVirtualGuest(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::addFinalComments", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddFinalComments(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::addScheduledAlert", func() {
            It("API Call Test", func() {

				err := sl_service.AddScheduledAlert(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::addScheduledAutoClose", func() {
            It("API Call Test", func() {

				err := sl_service.AddScheduledAutoClose(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::addUpdate", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddUpdate(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::createAdministrativeTicket", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateAdministrativeTicket(nil,nil,nil,nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::createCancelServerTicket", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateCancelServerTicket(nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::createCancelServiceTicket", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateCancelServiceTicket(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::createStandardTicket", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateStandardTicket(nil,nil,nil,nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::createUpgradeTicket", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateUpgradeTicket(nil,nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::edit", func() {
            It("API Call Test", func() {

                _, err := sl_service.Edit(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getAllTicketGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllTicketGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getAllTicketStatuses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllTicketStatuses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getAssignedAgents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedAgents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getAssignedUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getAttachedAdditionalEmails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachedAdditionalEmails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getAttachedDedicatedHosts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachedDedicatedHosts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getAttachedFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachedFile(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getAttachedFiles", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachedFiles()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getAttachedHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachedHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getAttachedHardwareCount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachedHardwareCount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getAttachedResources", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachedResources()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getAttachedVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachedVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getAwaitingUserResponseFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAwaitingUserResponseFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getBnppSupportedFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBnppSupportedFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getCancellationRequest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCancellationRequest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getEmployeeAttachments", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEmployeeAttachments()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getEuSupportedFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEuSupportedFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getFirstAttachedResource", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirstAttachedResource()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getFirstUpdate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirstUpdate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getFsboaSupportedFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFsboaSupportedFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getGroup", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGroup()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getInvoiceItems", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInvoiceItems()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getLastActivity", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLastActivity()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getLastEditor", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLastEditor()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getLastUpdate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLastUpdate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getLocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getNewUpdatesFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNewUpdatesFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getScheduledActions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetScheduledActions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getServerAdministrationBillingInvoice", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerAdministrationBillingInvoice()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getServerAdministrationRefundInvoice", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerAdministrationRefundInvoice()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getServiceProvider", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceProvider()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getState", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetState()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getSubject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getTagReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTagReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getTicketsClosedSinceDate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTicketsClosedSinceDate(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getUpdateRatingFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpdateRatingFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::getUpdates", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpdates()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::markAsViewed", func() {
            It("API Call Test", func() {

				err := sl_service.MarkAsViewed()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::removeAssignedAgent", func() {
            It("API Call Test", func() {

				err := sl_service.RemoveAssignedAgent(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::removeAttachedAdditionalEmails", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAttachedAdditionalEmails(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::removeAttachedHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAttachedHardware(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::removeAttachedVirtualGuest", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAttachedVirtualGuest(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::removeScheduledAlert", func() {
            It("API Call Test", func() {

				err := sl_service.RemoveScheduledAlert()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::removeScheduledAutoClose", func() {
            It("API Call Test", func() {

				err := sl_service.RemoveScheduledAutoClose()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::setTags", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetTags(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::surveyEligible", func() {
            It("API Call Test", func() {

                _, err := sl_service.SurveyEligible()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket::updateAttachedAdditionalEmails", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateAttachedAdditionalEmails(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Ticket_Attachment_File service", func() {
        var sl_service services.Ticket_Attachment_File
        BeforeEach(func() {
            sl_service = services.GetTicketAttachmentFileService(slsession)
        })


        Context("SoftLayer_Ticket_Attachment_File::getExtensionWhitelist", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetExtensionWhitelist()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket_Attachment_File::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket_Attachment_File::getTicket", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTicket()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket_Attachment_File::getUpdate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpdate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Ticket_Attachment_File_ServiceNow service", func() {
        var sl_service services.Ticket_Attachment_File_ServiceNow
        BeforeEach(func() {
            sl_service = services.GetTicketAttachmentFileServiceNowService(slsession)
        })


        Context("SoftLayer_Ticket_Attachment_File_ServiceNow::getExtensionWhitelist", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetExtensionWhitelist()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket_Attachment_File_ServiceNow::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket_Attachment_File_ServiceNow::getTicket", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTicket()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket_Attachment_File_ServiceNow::getUpdate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpdate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Ticket_Priority service", func() {
        var sl_service services.Ticket_Priority
        BeforeEach(func() {
            sl_service = services.GetTicketPriorityService(slsession)
        })


        Context("SoftLayer_Ticket_Priority::getPriorities", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPriorities()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Ticket_Subject service", func() {
        var sl_service services.Ticket_Subject
        BeforeEach(func() {
            sl_service = services.GetTicketSubjectService(slsession)
        })


        Context("SoftLayer_Ticket_Subject::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket_Subject::getCategory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCategory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket_Subject::getChildren", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetChildren()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket_Subject::getGroup", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGroup()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket_Subject::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket_Subject::getParent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetParent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket_Subject::getTopFiveKnowledgeLayerQuestions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTopFiveKnowledgeLayerQuestions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Ticket_Subject_Category service", func() {
        var sl_service services.Ticket_Subject_Category
        BeforeEach(func() {
            sl_service = services.GetTicketSubjectCategoryService(slsession)
        })


        Context("SoftLayer_Ticket_Subject_Category::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket_Subject_Category::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket_Subject_Category::getSubjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Ticket_Survey service", func() {
        var sl_service services.Ticket_Survey
        BeforeEach(func() {
            sl_service = services.GetTicketSurveyService(slsession)
        })


        Context("SoftLayer_Ticket_Survey::getPreference", func() {
            It("API Call Test", func() {

				err := sl_service.GetPreference()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket_Survey::optIn", func() {
            It("API Call Test", func() {

				err := sl_service.OptIn()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket_Survey::optOut", func() {
            It("API Call Test", func() {

				err := sl_service.OptOut()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Ticket_Update_Employee service", func() {
        var sl_service services.Ticket_Update_Employee
        BeforeEach(func() {
            sl_service = services.GetTicketUpdateEmployeeService(slsession)
        })


        Context("SoftLayer_Ticket_Update_Employee::addResponseRating", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddResponseRating(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Ticket_Update_Employee::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

})
