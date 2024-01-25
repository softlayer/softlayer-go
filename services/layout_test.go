
package services_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/softlayer/softlayer-go/session/sessionfakes"
    "github.com/softlayer/softlayer-go/services"
)   

var _ = Describe("Layout Tests", func() {
    var slsession *sessionfakes.FakeSLSession
    BeforeEach(func() {
        slsession = &sessionfakes.FakeSLSession{}
    })

    Context("Testing SoftLayer_Layout_Container service", func() {
        var sl_service services.Layout_Container
        BeforeEach(func() {
            sl_service = services.GetLayoutContainerService(slsession)
        })


        Context("SoftLayer_Layout_Container::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Container::getLayoutContainerType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLayoutContainerType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Container::getLayoutItems", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLayoutItems()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Container::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Layout_Item service", func() {
        var sl_service services.Layout_Item
        BeforeEach(func() {
            sl_service = services.GetLayoutItemService(slsession)
        })


        Context("SoftLayer_Layout_Item::getLayoutItemPreferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLayoutItemPreferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Item::getLayoutItemType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLayoutItemType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Item::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Layout_Profile service", func() {
        var sl_service services.Layout_Profile
        BeforeEach(func() {
            sl_service = services.GetLayoutProfileService(slsession)
        })


        Context("SoftLayer_Layout_Profile::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Profile::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Profile::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Profile::getLayoutContainers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLayoutContainers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Profile::getLayoutPreferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLayoutPreferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Profile::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Profile::modifyPreference", func() {
            It("API Call Test", func() {

                _, err := sl_service.ModifyPreference(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Profile::modifyPreferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.ModifyPreferences(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Layout_Profile_Containers service", func() {
        var sl_service services.Layout_Profile_Containers
        BeforeEach(func() {
            sl_service = services.GetLayoutProfileContainersService(slsession)
        })


        Context("SoftLayer_Layout_Profile_Containers::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Profile_Containers::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Profile_Containers::getLayoutContainerType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLayoutContainerType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Profile_Containers::getLayoutProfile", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLayoutProfile()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Profile_Containers::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Layout_Profile_Customer service", func() {
        var sl_service services.Layout_Profile_Customer
        BeforeEach(func() {
            sl_service = services.GetLayoutProfileCustomerService(slsession)
        })


        Context("SoftLayer_Layout_Profile_Customer::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Profile_Customer::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Profile_Customer::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Profile_Customer::getLayoutContainers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLayoutContainers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Profile_Customer::getLayoutPreferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLayoutPreferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Profile_Customer::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Profile_Customer::getUserRecord", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUserRecord()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Profile_Customer::modifyPreference", func() {
            It("API Call Test", func() {

                _, err := sl_service.ModifyPreference(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Profile_Customer::modifyPreferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.ModifyPreferences(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Layout_Profile_Preference service", func() {
        var sl_service services.Layout_Profile_Preference
        BeforeEach(func() {
            sl_service = services.GetLayoutProfilePreferenceService(slsession)
        })


        Context("SoftLayer_Layout_Profile_Preference::getLayoutContainer", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLayoutContainer()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Profile_Preference::getLayoutItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLayoutItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Profile_Preference::getLayoutPreference", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLayoutPreference()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Profile_Preference::getLayoutProfile", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLayoutProfile()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Layout_Profile_Preference::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

})
