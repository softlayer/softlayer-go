
package services_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/softlayer/softlayer-go/session/sessionfakes"
    "github.com/softlayer/softlayer-go/services"
)   

var _ = Describe("Notification Tests", func() {
    var slsession *sessionfakes.FakeSLSession
    BeforeEach(func() {
        slsession = &sessionfakes.FakeSLSession{}
    })

    Context("Testing SoftLayer_Notification service", func() {
        var sl_service services.Notification
        BeforeEach(func() {
            sl_service = services.GetNotificationService(slsession)
        })


        Context("SoftLayer_Notification::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification::getPreferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPreferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification::getRequiredPreferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRequiredPreferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Notification_Mobile service", func() {
        var sl_service services.Notification_Mobile
        BeforeEach(func() {
            sl_service = services.GetNotificationMobileService(slsession)
        })


        Context("SoftLayer_Notification_Mobile::createSubscriberForMobileDevice", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateSubscriberForMobileDevice(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_Mobile::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_Mobile::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_Mobile::getPreferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPreferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_Mobile::getRequiredPreferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRequiredPreferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Notification_Occurrence_Event service", func() {
        var sl_service services.Notification_Occurrence_Event
        BeforeEach(func() {
            sl_service = services.GetNotificationOccurrenceEventService(slsession)
        })


        Context("SoftLayer_Notification_Occurrence_Event::acknowledgeNotification", func() {
            It("API Call Test", func() {

                _, err := sl_service.AcknowledgeNotification()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_Occurrence_Event::getAcknowledgedFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAcknowledgedFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_Occurrence_Event::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_Occurrence_Event::getAttachedFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachedFile(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_Occurrence_Event::getAttachments", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachments()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_Occurrence_Event::getFirstUpdate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirstUpdate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_Occurrence_Event::getImpactedAccountCount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetImpactedAccountCount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_Occurrence_Event::getImpactedAccounts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetImpactedAccounts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_Occurrence_Event::getImpactedDeviceCount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetImpactedDeviceCount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_Occurrence_Event::getImpactedDevices", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetImpactedDevices()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_Occurrence_Event::getImpactedResources", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetImpactedResources()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_Occurrence_Event::getImpactedUsers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetImpactedUsers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_Occurrence_Event::getLastUpdate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLastUpdate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_Occurrence_Event::getNotificationOccurrenceEventType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNotificationOccurrenceEventType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_Occurrence_Event::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_Occurrence_Event::getStatusCode", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStatusCode()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_Occurrence_Event::getUpdates", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpdates()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Notification_Occurrence_User service", func() {
        var sl_service services.Notification_Occurrence_User
        BeforeEach(func() {
            sl_service = services.GetNotificationOccurrenceUserService(slsession)
        })


        Context("SoftLayer_Notification_Occurrence_User::acknowledge", func() {
            It("API Call Test", func() {

                _, err := sl_service.Acknowledge()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_Occurrence_User::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_Occurrence_User::getImpactedDeviceCount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetImpactedDeviceCount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_Occurrence_User::getImpactedResources", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetImpactedResources()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_Occurrence_User::getNotificationOccurrenceEvent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNotificationOccurrenceEvent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_Occurrence_User::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_Occurrence_User::getUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Notification_User_Subscriber service", func() {
        var sl_service services.Notification_User_Subscriber
        BeforeEach(func() {
            sl_service = services.GetNotificationUserSubscriberService(slsession)
        })


        Context("SoftLayer_Notification_User_Subscriber::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber::getDeliveryMethods", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDeliveryMethods()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber::getNotification", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNotification()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber::getPreferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPreferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber::getPreferencesDetails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPreferencesDetails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber::getResourceRecord", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResourceRecord()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber::getUserRecord", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUserRecord()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Notification_User_Subscriber_Billing service", func() {
        var sl_service services.Notification_User_Subscriber_Billing
        BeforeEach(func() {
            sl_service = services.GetNotificationUserSubscriberBillingService(slsession)
        })


        Context("SoftLayer_Notification_User_Subscriber_Billing::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber_Billing::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber_Billing::getDeliveryMethods", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDeliveryMethods()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber_Billing::getNotification", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNotification()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber_Billing::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber_Billing::getPreferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPreferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber_Billing::getPreferencesDetails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPreferencesDetails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber_Billing::getResourceRecord", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResourceRecord()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber_Billing::getUserRecord", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUserRecord()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Notification_User_Subscriber_Mobile service", func() {
        var sl_service services.Notification_User_Subscriber_Mobile
        BeforeEach(func() {
            sl_service = services.GetNotificationUserSubscriberMobileService(slsession)
        })


        Context("SoftLayer_Notification_User_Subscriber_Mobile::clearSnoozeTimer", func() {
            It("API Call Test", func() {

                _, err := sl_service.ClearSnoozeTimer()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber_Mobile::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber_Mobile::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber_Mobile::getDeliveryMethods", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDeliveryMethods()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber_Mobile::getNotification", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNotification()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber_Mobile::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber_Mobile::getPreferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPreferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber_Mobile::getPreferencesDetails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPreferencesDetails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber_Mobile::getResourceRecord", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResourceRecord()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber_Mobile::getUserRecord", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUserRecord()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber_Mobile::setSnoozeTimer", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetSnoozeTimer(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Notification_User_Subscriber_Preference service", func() {
        var sl_service services.Notification_User_Subscriber_Preference
        BeforeEach(func() {
            sl_service = services.GetNotificationUserSubscriberPreferenceService(slsession)
        })


        Context("SoftLayer_Notification_User_Subscriber_Preference::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber_Preference::editObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber_Preference::getDefaultPreference", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDefaultPreference()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber_Preference::getNotificationUserSubscriber", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNotificationUserSubscriber()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Notification_User_Subscriber_Preference::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

})
