
package services_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/softlayer/softlayer-go/session/sessionfakes"
    "github.com/softlayer/softlayer-go/services"
)   

var _ = Describe("User Tests", func() {
    var slsession *sessionfakes.FakeSLSession
    BeforeEach(func() {
        slsession = &sessionfakes.FakeSLSession{}
    })

    Context("Testing SoftLayer_User_Customer service", func() {
        var sl_service services.User_Customer
        BeforeEach(func() {
            sl_service = services.GetUserCustomerService(slsession)
        })


        Context("SoftLayer_User_Customer::acknowledgeSupportPolicy", func() {
            It("API Call Test", func() {

				err := sl_service.AcknowledgeSupportPolicy()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::addApiAuthenticationKey", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddApiAuthenticationKey()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::addBulkDedicatedHostAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddBulkDedicatedHostAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::addBulkHardwareAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddBulkHardwareAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::addBulkPortalPermission", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddBulkPortalPermission(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::addBulkRoles", func() {
            It("API Call Test", func() {

				err := sl_service.AddBulkRoles(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::addBulkVirtualGuestAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddBulkVirtualGuestAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::addDedicatedHostAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddDedicatedHostAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::addExternalBinding", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddExternalBinding(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::addHardwareAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddHardwareAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::addNotificationSubscriber", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddNotificationSubscriber(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::addPortalPermission", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddPortalPermission(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::addRole", func() {
            It("API Call Test", func() {

				err := sl_service.AddRole(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::addVirtualGuestAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddVirtualGuestAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::assignNewParentId", func() {
            It("API Call Test", func() {

                _, err := sl_service.AssignNewParentId(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::changePreference", func() {
            It("API Call Test", func() {

                _, err := sl_service.ChangePreference(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::createNotificationSubscriber", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateNotificationSubscriber(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::createSubscriberDeliveryMethods", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateSubscriberDeliveryMethods(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::deactivateNotificationSubscriber", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeactivateNotificationSubscriber(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::editObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::findUserPreference", func() {
            It("API Call Test", func() {

                _, err := sl_service.FindUserPreference(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getActions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getActiveExternalAuthenticationVendors", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveExternalAuthenticationVendors()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getAdditionalEmails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAdditionalEmails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getAgentImpersonationToken", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAgentImpersonationToken()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getAllowedDedicatedHostIds", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedDedicatedHostIds()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getAllowedHardwareIds", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedHardwareIds()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getAllowedVirtualGuestIds", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedVirtualGuestIds()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getApiAuthenticationKeys", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetApiAuthenticationKeys()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getAuthenticationToken", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAuthenticationToken(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getChildUsers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetChildUsers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getClosedTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetClosedTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getDedicatedHosts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDedicatedHosts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getDefaultAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDefaultAccount(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getExternalBindings", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetExternalBindings()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getHardwareCount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareCount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getHardwareNotifications", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareNotifications()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getHasAcknowledgedSupportPolicyFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasAcknowledgedSupportPolicyFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getHasFullDedicatedHostAccessFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasFullDedicatedHostAccessFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getHasFullHardwareAccessFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasFullHardwareAccessFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getHasFullVirtualGuestAccessFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasFullVirtualGuestAccessFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getIbmIdLink", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIbmIdLink()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getImpersonationToken", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetImpersonationToken()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getLayoutProfiles", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLayoutProfiles()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getLocale", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLocale()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getLoginAttempts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLoginAttempts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getLoginToken", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLoginToken(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getMappedAccounts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMappedAccounts(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getNotificationSubscribers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNotificationSubscribers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getOpenIdConnectMigrationState", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOpenIdConnectMigrationState()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getOpenTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOpenTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getOverrides", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOverrides()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getParent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetParent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getPasswordRequirements", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPasswordRequirements(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getPermissions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPermissions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getPortalLoginToken", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPortalLoginToken(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getPreference", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPreference(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getPreferenceTypes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPreferenceTypes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getPreferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPreferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getRequirementsForPasswordSet", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRequirementsForPasswordSet(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getRoles", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRoles()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getSecurityAnswers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSecurityAnswers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getSubscribers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubscribers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getSuccessfulLogins", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSuccessfulLogins()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getSupportPolicyAcknowledgementRequiredFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSupportPolicyAcknowledgementRequiredFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getSupportPolicyDocument", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSupportPolicyDocument()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getSupportPolicyName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSupportPolicyName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getSupportedLocales", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSupportedLocales()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getSurveyRequiredFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSurveyRequiredFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getSurveys", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSurveys()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getTimezone", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTimezone()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getUnsuccessfulLogins", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUnsuccessfulLogins()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getUserIdForPasswordSet", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUserIdForPasswordSet(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getUserLinks", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUserLinks()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getUserPreferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUserPreferences(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getUserStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUserStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getVirtualGuestCount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuestCount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::getVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::inTerminalStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.InTerminalStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::initiatePortalPasswordChange", func() {
            It("API Call Test", func() {

                _, err := sl_service.InitiatePortalPasswordChange(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::initiatePortalPasswordChangeByBrandAgent", func() {
            It("API Call Test", func() {

                _, err := sl_service.InitiatePortalPasswordChangeByBrandAgent(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::inviteUserToLinkOpenIdConnect", func() {
            It("API Call Test", func() {

				err := sl_service.InviteUserToLinkOpenIdConnect(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::isMasterUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsMasterUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::isValidPortalPassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsValidPortalPassword(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::performExternalAuthentication", func() {
            It("API Call Test", func() {

                _, err := sl_service.PerformExternalAuthentication(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::processPasswordSetRequest", func() {
            It("API Call Test", func() {

                _, err := sl_service.ProcessPasswordSetRequest(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::removeAllDedicatedHostAccessForThisUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAllDedicatedHostAccessForThisUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::removeAllHardwareAccessForThisUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAllHardwareAccessForThisUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::removeAllVirtualAccessForThisUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAllVirtualAccessForThisUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::removeApiAuthenticationKey", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveApiAuthenticationKey(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::removeBulkDedicatedHostAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveBulkDedicatedHostAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::removeBulkHardwareAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveBulkHardwareAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::removeBulkPortalPermission", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveBulkPortalPermission(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::removeBulkRoles", func() {
            It("API Call Test", func() {

				err := sl_service.RemoveBulkRoles(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::removeBulkVirtualGuestAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveBulkVirtualGuestAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::removeDedicatedHostAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveDedicatedHostAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::removeExternalBinding", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveExternalBinding(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::removeHardwareAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveHardwareAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::removePortalPermission", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemovePortalPermission(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::removeRole", func() {
            It("API Call Test", func() {

				err := sl_service.RemoveRole(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::removeSecurityAnswers", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveSecurityAnswers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::removeVirtualGuestAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveVirtualGuestAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::resetOpenIdConnectLink", func() {
            It("API Call Test", func() {

				err := sl_service.ResetOpenIdConnectLink(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::resetOpenIdConnectLinkUnifiedUserManagementMode", func() {
            It("API Call Test", func() {

				err := sl_service.ResetOpenIdConnectLinkUnifiedUserManagementMode(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::samlAuthenticate", func() {
            It("API Call Test", func() {

                _, err := sl_service.SamlAuthenticate(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::samlBeginAuthentication", func() {
            It("API Call Test", func() {

                _, err := sl_service.SamlBeginAuthentication(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::samlBeginLogout", func() {
            It("API Call Test", func() {

                _, err := sl_service.SamlBeginLogout()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::samlLogout", func() {
            It("API Call Test", func() {

				err := sl_service.SamlLogout(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::selfPasswordChange", func() {
            It("API Call Test", func() {

				err := sl_service.SelfPasswordChange(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::setDefaultAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetDefaultAccount(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::silentlyMigrateUserOpenIdConnect", func() {
            It("API Call Test", func() {

                _, err := sl_service.SilentlyMigrateUserOpenIdConnect(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::turnOffMasterUserPermissionCheckMode", func() {
            It("API Call Test", func() {

				err := sl_service.TurnOffMasterUserPermissionCheckMode()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::turnOnMasterUserPermissionCheckMode", func() {
            It("API Call Test", func() {

				err := sl_service.TurnOnMasterUserPermissionCheckMode()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::updateNotificationSubscriber", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateNotificationSubscriber(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::updateSecurityAnswers", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateSecurityAnswers(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::updateSubscriberDeliveryMethod", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateSubscriberDeliveryMethod(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::updateVpnPassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateVpnPassword(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::updateVpnUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateVpnUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer::validateAuthenticationToken", func() {
            It("API Call Test", func() {

                _, err := sl_service.ValidateAuthenticationToken(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_User_Customer_ApiAuthentication service", func() {
        var sl_service services.User_Customer_ApiAuthentication
        BeforeEach(func() {
            sl_service = services.GetUserCustomerApiAuthenticationService(slsession)
        })


        Context("SoftLayer_User_Customer_ApiAuthentication::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_ApiAuthentication::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_ApiAuthentication::getUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_User_Customer_CustomerPermission_Permission service", func() {
        var sl_service services.User_Customer_CustomerPermission_Permission
        BeforeEach(func() {
            sl_service = services.GetUserCustomerCustomerPermissionPermissionService(slsession)
        })


        Context("SoftLayer_User_Customer_CustomerPermission_Permission::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_CustomerPermission_Permission::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_User_Customer_External_Binding service", func() {
        var sl_service services.User_Customer_External_Binding
        BeforeEach(func() {
            sl_service = services.GetUserCustomerExternalBindingService(slsession)
        })


        Context("SoftLayer_User_Customer_External_Binding::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding::disable", func() {
            It("API Call Test", func() {

                _, err := sl_service.Disable(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding::enable", func() {
            It("API Call Test", func() {

                _, err := sl_service.Enable()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding::getAttributes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttributes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding::getNote", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNote()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding::getType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding::getUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding::getVendor", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVendor()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding::updateNote", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateNote(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_User_Customer_External_Binding_Totp service", func() {
        var sl_service services.User_Customer_External_Binding_Totp
        BeforeEach(func() {
            sl_service = services.GetUserCustomerExternalBindingTotpService(slsession)
        })


        Context("SoftLayer_User_Customer_External_Binding_Totp::activate", func() {
            It("API Call Test", func() {

                _, err := sl_service.Activate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Totp::deactivate", func() {
            It("API Call Test", func() {

                _, err := sl_service.Deactivate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Totp::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Totp::disable", func() {
            It("API Call Test", func() {

                _, err := sl_service.Disable(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Totp::enable", func() {
            It("API Call Test", func() {

                _, err := sl_service.Enable()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Totp::generateSecretKey", func() {
            It("API Call Test", func() {

                _, err := sl_service.GenerateSecretKey()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Totp::getAttributes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttributes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Totp::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Totp::getNote", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNote()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Totp::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Totp::getType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Totp::getUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Totp::getVendor", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVendor()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Totp::updateNote", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateNote(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_User_Customer_External_Binding_Vendor service", func() {
        var sl_service services.User_Customer_External_Binding_Vendor
        BeforeEach(func() {
            sl_service = services.GetUserCustomerExternalBindingVendorService(slsession)
        })


        Context("SoftLayer_User_Customer_External_Binding_Vendor::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Vendor::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_User_Customer_External_Binding_Verisign service", func() {
        var sl_service services.User_Customer_External_Binding_Verisign
        BeforeEach(func() {
            sl_service = services.GetUserCustomerExternalBindingVerisignService(slsession)
        })


        Context("SoftLayer_User_Customer_External_Binding_Verisign::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Verisign::disable", func() {
            It("API Call Test", func() {

                _, err := sl_service.Disable(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Verisign::enable", func() {
            It("API Call Test", func() {

                _, err := sl_service.Enable()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Verisign::getActivationCodeForMobileClient", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActivationCodeForMobileClient()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Verisign::getAttributes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttributes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Verisign::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Verisign::getCredentialExpirationDate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCredentialExpirationDate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Verisign::getCredentialLastUpdateDate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCredentialLastUpdateDate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Verisign::getCredentialState", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCredentialState()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Verisign::getCredentialType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCredentialType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Verisign::getNote", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNote()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Verisign::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Verisign::getType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Verisign::getUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Verisign::getVendor", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVendor()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Verisign::unlock", func() {
            It("API Call Test", func() {

                _, err := sl_service.Unlock(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Verisign::updateNote", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateNote(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_External_Binding_Verisign::validateCredentialId", func() {
            It("API Call Test", func() {

                _, err := sl_service.ValidateCredentialId(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_User_Customer_Invitation service", func() {
        var sl_service services.User_Customer_Invitation
        BeforeEach(func() {
            sl_service = services.GetUserCustomerInvitationService(slsession)
        })


        Context("SoftLayer_User_Customer_Invitation::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_Invitation::getUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_User_Customer_Notification_Hardware service", func() {
        var sl_service services.User_Customer_Notification_Hardware
        BeforeEach(func() {
            sl_service = services.GetUserCustomerNotificationHardwareService(slsession)
        })


        Context("SoftLayer_User_Customer_Notification_Hardware::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_Notification_Hardware::createObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_Notification_Hardware::deleteObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_Notification_Hardware::findByHardwareId", func() {
            It("API Call Test", func() {

                _, err := sl_service.FindByHardwareId(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_Notification_Hardware::getHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_Notification_Hardware::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_Notification_Hardware::getUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_User_Customer_Notification_Virtual_Guest service", func() {
        var sl_service services.User_Customer_Notification_Virtual_Guest
        BeforeEach(func() {
            sl_service = services.GetUserCustomerNotificationVirtualGuestService(slsession)
        })


        Context("SoftLayer_User_Customer_Notification_Virtual_Guest::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_Notification_Virtual_Guest::createObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_Notification_Virtual_Guest::deleteObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_Notification_Virtual_Guest::findByGuestId", func() {
            It("API Call Test", func() {

                _, err := sl_service.FindByGuestId(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_Notification_Virtual_Guest::getGuest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGuest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_Notification_Virtual_Guest::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_Notification_Virtual_Guest::getUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_User_Customer_OpenIdConnect service", func() {
        var sl_service services.User_Customer_OpenIdConnect
        BeforeEach(func() {
            sl_service = services.GetUserCustomerOpenIdConnectService(slsession)
        })


        Context("SoftLayer_User_Customer_OpenIdConnect::acknowledgeSupportPolicy", func() {
            It("API Call Test", func() {

				err := sl_service.AcknowledgeSupportPolicy()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::activateOpenIdConnectUser", func() {
            It("API Call Test", func() {

				err := sl_service.ActivateOpenIdConnectUser(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::addApiAuthenticationKey", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddApiAuthenticationKey()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::addBulkDedicatedHostAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddBulkDedicatedHostAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::addBulkHardwareAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddBulkHardwareAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::addBulkPortalPermission", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddBulkPortalPermission(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::addBulkRoles", func() {
            It("API Call Test", func() {

				err := sl_service.AddBulkRoles(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::addBulkVirtualGuestAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddBulkVirtualGuestAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::addDedicatedHostAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddDedicatedHostAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::addExternalBinding", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddExternalBinding(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::addHardwareAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddHardwareAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::addNotificationSubscriber", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddNotificationSubscriber(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::addPortalPermission", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddPortalPermission(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::addRole", func() {
            It("API Call Test", func() {

				err := sl_service.AddRole(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::addVirtualGuestAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddVirtualGuestAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::assignNewParentId", func() {
            It("API Call Test", func() {

                _, err := sl_service.AssignNewParentId(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::changePreference", func() {
            It("API Call Test", func() {

                _, err := sl_service.ChangePreference(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::completeInvitationAfterLogin", func() {
            It("API Call Test", func() {

				err := sl_service.CompleteInvitationAfterLogin(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::createNotificationSubscriber", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateNotificationSubscriber(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::createOpenIdConnectUserAndCompleteInvitation", func() {
            It("API Call Test", func() {

				err := sl_service.CreateOpenIdConnectUserAndCompleteInvitation(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::createSubscriberDeliveryMethods", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateSubscriberDeliveryMethods(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::deactivateNotificationSubscriber", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeactivateNotificationSubscriber(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::declineInvitation", func() {
            It("API Call Test", func() {

				err := sl_service.DeclineInvitation(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::editObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::findUserPreference", func() {
            It("API Call Test", func() {

                _, err := sl_service.FindUserPreference(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getActions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getActiveExternalAuthenticationVendors", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveExternalAuthenticationVendors()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getAdditionalEmails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAdditionalEmails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getAgentImpersonationToken", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAgentImpersonationToken()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getAllowedDedicatedHostIds", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedDedicatedHostIds()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getAllowedHardwareIds", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedHardwareIds()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getAllowedVirtualGuestIds", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedVirtualGuestIds()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getApiAuthenticationKeys", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetApiAuthenticationKeys()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getAuthenticationToken", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAuthenticationToken(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getChildUsers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetChildUsers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getClosedTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetClosedTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getDedicatedHosts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDedicatedHosts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getDefaultAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDefaultAccount(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getExternalBindings", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetExternalBindings()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getHardwareCount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareCount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getHardwareNotifications", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareNotifications()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getHasAcknowledgedSupportPolicyFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasAcknowledgedSupportPolicyFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getHasFullDedicatedHostAccessFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasFullDedicatedHostAccessFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getHasFullHardwareAccessFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasFullHardwareAccessFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getHasFullVirtualGuestAccessFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasFullVirtualGuestAccessFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getIbmIdLink", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIbmIdLink()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getImpersonationToken", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetImpersonationToken()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getLayoutProfiles", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLayoutProfiles()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getLocale", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLocale()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getLoginAccountInfoOpenIdConnect", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLoginAccountInfoOpenIdConnect(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getLoginAttempts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLoginAttempts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getLoginToken", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLoginToken(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getMappedAccounts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMappedAccounts(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getNotificationSubscribers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNotificationSubscribers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getOpenIdConnectMigrationState", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOpenIdConnectMigrationState()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getOpenIdRegistrationInfoFromCode", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOpenIdRegistrationInfoFromCode(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getOpenTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOpenTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getOverrides", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOverrides()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getParent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetParent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getPasswordRequirements", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPasswordRequirements(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getPermissions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPermissions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getPortalLoginToken", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPortalLoginToken(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getPortalLoginTokenOpenIdConnect", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPortalLoginTokenOpenIdConnect(nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getPreference", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPreference(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getPreferenceTypes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPreferenceTypes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getPreferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPreferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getRequirementsForPasswordSet", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRequirementsForPasswordSet(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getRoles", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRoles()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getSecurityAnswers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSecurityAnswers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getSubscribers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubscribers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getSuccessfulLogins", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSuccessfulLogins()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getSupportPolicyAcknowledgementRequiredFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSupportPolicyAcknowledgementRequiredFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getSupportPolicyDocument", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSupportPolicyDocument()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getSupportPolicyName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSupportPolicyName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getSupportedLocales", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSupportedLocales()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getSurveyRequiredFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSurveyRequiredFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getSurveys", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSurveys()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getTimezone", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTimezone()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getUnsuccessfulLogins", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUnsuccessfulLogins()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getUserForUnifiedInvitation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUserForUnifiedInvitation(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getUserIdForPasswordSet", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUserIdForPasswordSet(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getUserLinks", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUserLinks()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getUserPreferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUserPreferences(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getUserStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUserStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getVirtualGuestCount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuestCount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::getVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::inTerminalStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.InTerminalStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::initiatePortalPasswordChange", func() {
            It("API Call Test", func() {

                _, err := sl_service.InitiatePortalPasswordChange(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::initiatePortalPasswordChangeByBrandAgent", func() {
            It("API Call Test", func() {

                _, err := sl_service.InitiatePortalPasswordChangeByBrandAgent(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::inviteUserToLinkOpenIdConnect", func() {
            It("API Call Test", func() {

				err := sl_service.InviteUserToLinkOpenIdConnect(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::isMasterUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsMasterUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::isValidPortalPassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsValidPortalPassword(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::performExternalAuthentication", func() {
            It("API Call Test", func() {

                _, err := sl_service.PerformExternalAuthentication(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::processPasswordSetRequest", func() {
            It("API Call Test", func() {

                _, err := sl_service.ProcessPasswordSetRequest(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::removeAllDedicatedHostAccessForThisUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAllDedicatedHostAccessForThisUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::removeAllHardwareAccessForThisUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAllHardwareAccessForThisUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::removeAllVirtualAccessForThisUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAllVirtualAccessForThisUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::removeApiAuthenticationKey", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveApiAuthenticationKey(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::removeBulkDedicatedHostAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveBulkDedicatedHostAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::removeBulkHardwareAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveBulkHardwareAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::removeBulkPortalPermission", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveBulkPortalPermission(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::removeBulkRoles", func() {
            It("API Call Test", func() {

				err := sl_service.RemoveBulkRoles(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::removeBulkVirtualGuestAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveBulkVirtualGuestAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::removeDedicatedHostAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveDedicatedHostAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::removeExternalBinding", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveExternalBinding(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::removeHardwareAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveHardwareAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::removePortalPermission", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemovePortalPermission(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::removeRole", func() {
            It("API Call Test", func() {

				err := sl_service.RemoveRole(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::removeSecurityAnswers", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveSecurityAnswers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::removeVirtualGuestAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveVirtualGuestAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::resetOpenIdConnectLink", func() {
            It("API Call Test", func() {

				err := sl_service.ResetOpenIdConnectLink(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::resetOpenIdConnectLinkUnifiedUserManagementMode", func() {
            It("API Call Test", func() {

				err := sl_service.ResetOpenIdConnectLinkUnifiedUserManagementMode(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::samlAuthenticate", func() {
            It("API Call Test", func() {

                _, err := sl_service.SamlAuthenticate(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::samlBeginAuthentication", func() {
            It("API Call Test", func() {

                _, err := sl_service.SamlBeginAuthentication(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::samlBeginLogout", func() {
            It("API Call Test", func() {

                _, err := sl_service.SamlBeginLogout()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::samlLogout", func() {
            It("API Call Test", func() {

				err := sl_service.SamlLogout(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::selfPasswordChange", func() {
            It("API Call Test", func() {

				err := sl_service.SelfPasswordChange(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::setDefaultAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetDefaultAccount(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::silentlyMigrateUserOpenIdConnect", func() {
            It("API Call Test", func() {

                _, err := sl_service.SilentlyMigrateUserOpenIdConnect(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::turnOffMasterUserPermissionCheckMode", func() {
            It("API Call Test", func() {

				err := sl_service.TurnOffMasterUserPermissionCheckMode()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::turnOnMasterUserPermissionCheckMode", func() {
            It("API Call Test", func() {

				err := sl_service.TurnOnMasterUserPermissionCheckMode()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::updateNotificationSubscriber", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateNotificationSubscriber(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::updateSecurityAnswers", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateSecurityAnswers(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::updateSubscriberDeliveryMethod", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateSubscriberDeliveryMethod(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::updateVpnPassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateVpnPassword(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::updateVpnUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateVpnUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect::validateAuthenticationToken", func() {
            It("API Call Test", func() {

                _, err := sl_service.ValidateAuthenticationToken(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_User_Customer_OpenIdConnect_TrustedProfile service", func() {
        var sl_service services.User_Customer_OpenIdConnect_TrustedProfile
        BeforeEach(func() {
            sl_service = services.GetUserCustomerOpenIdConnectTrustedProfileService(slsession)
        })


        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::acknowledgeSupportPolicy", func() {
            It("API Call Test", func() {

				err := sl_service.AcknowledgeSupportPolicy()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::activateOpenIdConnectUser", func() {
            It("API Call Test", func() {

				err := sl_service.ActivateOpenIdConnectUser(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::addApiAuthenticationKey", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddApiAuthenticationKey()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::addBulkDedicatedHostAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddBulkDedicatedHostAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::addBulkHardwareAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddBulkHardwareAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::addBulkPortalPermission", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddBulkPortalPermission(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::addBulkRoles", func() {
            It("API Call Test", func() {

				err := sl_service.AddBulkRoles(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::addBulkVirtualGuestAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddBulkVirtualGuestAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::addDedicatedHostAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddDedicatedHostAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::addExternalBinding", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddExternalBinding(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::addHardwareAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddHardwareAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::addNotificationSubscriber", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddNotificationSubscriber(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::addPortalPermission", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddPortalPermission(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::addRole", func() {
            It("API Call Test", func() {

				err := sl_service.AddRole(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::addVirtualGuestAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddVirtualGuestAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::assignNewParentId", func() {
            It("API Call Test", func() {

                _, err := sl_service.AssignNewParentId(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::changePreference", func() {
            It("API Call Test", func() {

                _, err := sl_service.ChangePreference(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::completeInvitationAfterLogin", func() {
            It("API Call Test", func() {

				err := sl_service.CompleteInvitationAfterLogin(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::createNotificationSubscriber", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateNotificationSubscriber(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::createOpenIdConnectUserAndCompleteInvitation", func() {
            It("API Call Test", func() {

				err := sl_service.CreateOpenIdConnectUserAndCompleteInvitation(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::createSubscriberDeliveryMethods", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateSubscriberDeliveryMethods(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::deactivateNotificationSubscriber", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeactivateNotificationSubscriber(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::declineInvitation", func() {
            It("API Call Test", func() {

				err := sl_service.DeclineInvitation(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::editObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::findUserPreference", func() {
            It("API Call Test", func() {

                _, err := sl_service.FindUserPreference(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getActions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getActiveExternalAuthenticationVendors", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveExternalAuthenticationVendors()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getAdditionalEmails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAdditionalEmails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getAgentImpersonationToken", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAgentImpersonationToken()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getAllowedDedicatedHostIds", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedDedicatedHostIds()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getAllowedHardwareIds", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedHardwareIds()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getAllowedVirtualGuestIds", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedVirtualGuestIds()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getApiAuthenticationKeys", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetApiAuthenticationKeys()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getAuthenticationToken", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAuthenticationToken(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getChildUsers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetChildUsers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getClosedTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetClosedTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getDedicatedHosts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDedicatedHosts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getDefaultAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDefaultAccount(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getExternalBindings", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetExternalBindings()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getHardwareCount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareCount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getHardwareNotifications", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareNotifications()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getHasAcknowledgedSupportPolicyFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasAcknowledgedSupportPolicyFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getHasFullDedicatedHostAccessFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasFullDedicatedHostAccessFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getHasFullHardwareAccessFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasFullHardwareAccessFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getHasFullVirtualGuestAccessFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasFullVirtualGuestAccessFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getIbmIdLink", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIbmIdLink()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getImpersonationToken", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetImpersonationToken()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getLayoutProfiles", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLayoutProfiles()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getLocale", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLocale()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getLoginAccountInfoOpenIdConnect", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLoginAccountInfoOpenIdConnect(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getLoginAttempts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLoginAttempts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getLoginToken", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLoginToken(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getMappedAccounts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMappedAccounts(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getNotificationSubscribers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNotificationSubscribers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getOpenIdConnectMigrationState", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOpenIdConnectMigrationState()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getOpenIdRegistrationInfoFromCode", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOpenIdRegistrationInfoFromCode(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getOpenTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOpenTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getOverrides", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOverrides()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getParent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetParent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getPasswordRequirements", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPasswordRequirements(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getPermissions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPermissions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getPortalLoginToken", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPortalLoginToken(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getPortalLoginTokenOpenIdConnect", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPortalLoginTokenOpenIdConnect(nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getPreference", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPreference(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getPreferenceTypes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPreferenceTypes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getPreferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPreferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getRequirementsForPasswordSet", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRequirementsForPasswordSet(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getRoles", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRoles()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getSecurityAnswers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSecurityAnswers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getSubscribers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubscribers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getSuccessfulLogins", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSuccessfulLogins()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getSupportPolicyAcknowledgementRequiredFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSupportPolicyAcknowledgementRequiredFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getSupportPolicyDocument", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSupportPolicyDocument()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getSupportPolicyName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSupportPolicyName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getSupportedLocales", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSupportedLocales()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getSurveyRequiredFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSurveyRequiredFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getSurveys", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSurveys()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getTimezone", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTimezone()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getUnsuccessfulLogins", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUnsuccessfulLogins()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getUserForUnifiedInvitation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUserForUnifiedInvitation(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getUserIdForPasswordSet", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUserIdForPasswordSet(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getUserLinks", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUserLinks()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getUserPreferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUserPreferences(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getUserStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUserStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getVirtualGuestCount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuestCount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::getVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::inTerminalStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.InTerminalStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::initiatePortalPasswordChange", func() {
            It("API Call Test", func() {

                _, err := sl_service.InitiatePortalPasswordChange(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::initiatePortalPasswordChangeByBrandAgent", func() {
            It("API Call Test", func() {

                _, err := sl_service.InitiatePortalPasswordChangeByBrandAgent(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::inviteUserToLinkOpenIdConnect", func() {
            It("API Call Test", func() {

				err := sl_service.InviteUserToLinkOpenIdConnect(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::isMasterUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsMasterUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::isValidPortalPassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsValidPortalPassword(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::performExternalAuthentication", func() {
            It("API Call Test", func() {

                _, err := sl_service.PerformExternalAuthentication(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::processPasswordSetRequest", func() {
            It("API Call Test", func() {

                _, err := sl_service.ProcessPasswordSetRequest(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::removeAllDedicatedHostAccessForThisUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAllDedicatedHostAccessForThisUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::removeAllHardwareAccessForThisUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAllHardwareAccessForThisUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::removeAllVirtualAccessForThisUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAllVirtualAccessForThisUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::removeApiAuthenticationKey", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveApiAuthenticationKey(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::removeBulkDedicatedHostAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveBulkDedicatedHostAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::removeBulkHardwareAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveBulkHardwareAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::removeBulkPortalPermission", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveBulkPortalPermission(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::removeBulkRoles", func() {
            It("API Call Test", func() {

				err := sl_service.RemoveBulkRoles(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::removeBulkVirtualGuestAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveBulkVirtualGuestAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::removeDedicatedHostAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveDedicatedHostAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::removeExternalBinding", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveExternalBinding(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::removeHardwareAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveHardwareAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::removePortalPermission", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemovePortalPermission(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::removeRole", func() {
            It("API Call Test", func() {

				err := sl_service.RemoveRole(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::removeSecurityAnswers", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveSecurityAnswers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::removeVirtualGuestAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveVirtualGuestAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::resetOpenIdConnectLink", func() {
            It("API Call Test", func() {

				err := sl_service.ResetOpenIdConnectLink(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::resetOpenIdConnectLinkUnifiedUserManagementMode", func() {
            It("API Call Test", func() {

				err := sl_service.ResetOpenIdConnectLinkUnifiedUserManagementMode(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::samlAuthenticate", func() {
            It("API Call Test", func() {

                _, err := sl_service.SamlAuthenticate(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::samlBeginAuthentication", func() {
            It("API Call Test", func() {

                _, err := sl_service.SamlBeginAuthentication(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::samlBeginLogout", func() {
            It("API Call Test", func() {

                _, err := sl_service.SamlBeginLogout()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::samlLogout", func() {
            It("API Call Test", func() {

				err := sl_service.SamlLogout(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::selfPasswordChange", func() {
            It("API Call Test", func() {

				err := sl_service.SelfPasswordChange(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::setDefaultAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetDefaultAccount(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::silentlyMigrateUserOpenIdConnect", func() {
            It("API Call Test", func() {

                _, err := sl_service.SilentlyMigrateUserOpenIdConnect(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::turnOffMasterUserPermissionCheckMode", func() {
            It("API Call Test", func() {

				err := sl_service.TurnOffMasterUserPermissionCheckMode()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::turnOnMasterUserPermissionCheckMode", func() {
            It("API Call Test", func() {

				err := sl_service.TurnOnMasterUserPermissionCheckMode()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::updateNotificationSubscriber", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateNotificationSubscriber(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::updateSecurityAnswers", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateSecurityAnswers(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::updateSubscriberDeliveryMethod", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateSubscriberDeliveryMethod(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::updateVpnPassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateVpnPassword(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::updateVpnUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateVpnUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_OpenIdConnect_TrustedProfile::validateAuthenticationToken", func() {
            It("API Call Test", func() {

                _, err := sl_service.ValidateAuthenticationToken(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_User_Customer_Profile_Event_HyperWarp service", func() {
        var sl_service services.User_Customer_Profile_Event_HyperWarp
        BeforeEach(func() {
            sl_service = services.GetUserCustomerProfileEventHyperWarpService(slsession)
        })


        Context("SoftLayer_User_Customer_Profile_Event_HyperWarp::receiveEventDirect", func() {
            It("API Call Test", func() {

                _, err := sl_service.ReceiveEventDirect(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_User_Customer_Prospect_ServiceProvider_EnrollRequest service", func() {
        var sl_service services.User_Customer_Prospect_ServiceProvider_EnrollRequest
        BeforeEach(func() {
            sl_service = services.GetUserCustomerProspectServiceProviderEnrollRequestService(slsession)
        })


        Context("SoftLayer_User_Customer_Prospect_ServiceProvider_EnrollRequest::enroll", func() {
            It("API Call Test", func() {

                _, err := sl_service.Enroll(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_Prospect_ServiceProvider_EnrollRequest::getCompanyType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCompanyType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_Prospect_ServiceProvider_EnrollRequest::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_User_Customer_Security_Answer service", func() {
        var sl_service services.User_Customer_Security_Answer
        BeforeEach(func() {
            sl_service = services.GetUserCustomerSecurityAnswerService(slsession)
        })


        Context("SoftLayer_User_Customer_Security_Answer::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_Security_Answer::getQuestion", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetQuestion()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_Security_Answer::getUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_User_Customer_Status service", func() {
        var sl_service services.User_Customer_Status
        BeforeEach(func() {
            sl_service = services.GetUserCustomerStatusService(slsession)
        })


        Context("SoftLayer_User_Customer_Status::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Customer_Status::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_User_External_Binding service", func() {
        var sl_service services.User_External_Binding
        BeforeEach(func() {
            sl_service = services.GetUserExternalBindingService(slsession)
        })


        Context("SoftLayer_User_External_Binding::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_External_Binding::getAttributes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttributes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_External_Binding::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_External_Binding::getNote", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNote()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_External_Binding::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_External_Binding::getType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_External_Binding::getVendor", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVendor()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_External_Binding::updateNote", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateNote(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_User_External_Binding_Vendor service", func() {
        var sl_service services.User_External_Binding_Vendor
        BeforeEach(func() {
            sl_service = services.GetUserExternalBindingVendorService(slsession)
        })


        Context("SoftLayer_User_External_Binding_Vendor::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_External_Binding_Vendor::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_User_Permission_Action service", func() {
        var sl_service services.User_Permission_Action
        BeforeEach(func() {
            sl_service = services.GetUserPermissionActionService(slsession)
        })


        Context("SoftLayer_User_Permission_Action::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Action::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_User_Permission_Group service", func() {
        var sl_service services.User_Permission_Group
        BeforeEach(func() {
            sl_service = services.GetUserPermissionGroupService(slsession)
        })


        Context("SoftLayer_User_Permission_Group::addAction", func() {
            It("API Call Test", func() {

				err := sl_service.AddAction(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Group::addBulkActions", func() {
            It("API Call Test", func() {

				err := sl_service.AddBulkActions(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Group::addBulkResourceObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddBulkResourceObjects(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Group::addResourceObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddResourceObject(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Group::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Group::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Group::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Group::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Group::getActions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Group::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Group::getRoles", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRoles()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Group::getType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Group::linkRole", func() {
            It("API Call Test", func() {

				err := sl_service.LinkRole(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Group::removeAction", func() {
            It("API Call Test", func() {

				err := sl_service.RemoveAction(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Group::removeBulkActions", func() {
            It("API Call Test", func() {

				err := sl_service.RemoveBulkActions(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Group::removeBulkResourceObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveBulkResourceObjects(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Group::removeResourceObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveResourceObject(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Group::unlinkRole", func() {
            It("API Call Test", func() {

				err := sl_service.UnlinkRole(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_User_Permission_Group_Type service", func() {
        var sl_service services.User_Permission_Group_Type
        BeforeEach(func() {
            sl_service = services.GetUserPermissionGroupTypeService(slsession)
        })


        Context("SoftLayer_User_Permission_Group_Type::getGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Group_Type::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_User_Permission_Resource_Type service", func() {
        var sl_service services.User_Permission_Resource_Type
        BeforeEach(func() {
            sl_service = services.GetUserPermissionResourceTypeService(slsession)
        })


        Context("SoftLayer_User_Permission_Resource_Type::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Resource_Type::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_User_Permission_Role service", func() {
        var sl_service services.User_Permission_Role
        BeforeEach(func() {
            sl_service = services.GetUserPermissionRoleService(slsession)
        })


        Context("SoftLayer_User_Permission_Role::addUser", func() {
            It("API Call Test", func() {

				err := sl_service.AddUser(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Role::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Role::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Role::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Role::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Role::getActions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Role::getGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Role::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Role::getUsers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUsers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Role::linkGroup", func() {
            It("API Call Test", func() {

				err := sl_service.LinkGroup(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Role::removeUser", func() {
            It("API Call Test", func() {

				err := sl_service.RemoveUser(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Permission_Role::unlinkGroup", func() {
            It("API Call Test", func() {

				err := sl_service.UnlinkGroup(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_User_Security_Question service", func() {
        var sl_service services.User_Security_Question
        BeforeEach(func() {
            sl_service = services.GetUserSecurityQuestionService(slsession)
        })


        Context("SoftLayer_User_Security_Question::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_User_Security_Question::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

})
