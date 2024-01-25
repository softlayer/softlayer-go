
package services_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/softlayer/softlayer-go/session/sessionfakes"
    "github.com/softlayer/softlayer-go/services"
)   

var _ = Describe("Tag Tests", func() {
    var slsession *sessionfakes.FakeSLSession
    BeforeEach(func() {
        slsession = &sessionfakes.FakeSLSession{}
    })

    Context("Testing SoftLayer_Tag service", func() {
        var sl_service services.Tag
        BeforeEach(func() {
            sl_service = services.GetTagService(slsession)
        })


        Context("SoftLayer_Tag::autoComplete", func() {
            It("API Call Test", func() {

                _, err := sl_service.AutoComplete(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Tag::deleteTag", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteTag(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Tag::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Tag::getAllTagTypes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllTagTypes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Tag::getAttachedTagsForCurrentUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachedTagsForCurrentUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Tag::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Tag::getReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Tag::getTagByTagName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTagByTagName(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Tag::getUnattachedTagsForCurrentUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUnattachedTagsForCurrentUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Tag::setTags", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetTags(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

})
