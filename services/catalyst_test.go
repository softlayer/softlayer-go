
package services_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/softlayer/softlayer-go/session/sessionfakes"
    "github.com/softlayer/softlayer-go/services"
)   

var _ = Describe("Catalyst Tests", func() {
    var slsession *sessionfakes.FakeSLSession
    BeforeEach(func() {
        slsession = &sessionfakes.FakeSLSession{}
    })

    Context("Testing SoftLayer_Catalyst_Company_Type service", func() {
        var sl_service services.Catalyst_Company_Type
        BeforeEach(func() {
            sl_service = services.GetCatalystCompanyTypeService(slsession)
        })


        Context("SoftLayer_Catalyst_Company_Type::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Catalyst_Company_Type::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Catalyst_Enrollment service", func() {
        var sl_service services.Catalyst_Enrollment
        BeforeEach(func() {
            sl_service = services.GetCatalystEnrollmentService(slsession)
        })


        Context("SoftLayer_Catalyst_Enrollment::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Catalyst_Enrollment::getAffiliate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAffiliate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Catalyst_Enrollment::getAffiliates", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAffiliates()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Catalyst_Enrollment::getCompanyType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCompanyType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Catalyst_Enrollment::getCompanyTypes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCompanyTypes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Catalyst_Enrollment::getEnrollmentRequestAnnualRevenueOptions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEnrollmentRequestAnnualRevenueOptions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Catalyst_Enrollment::getEnrollmentRequestUserCountOptions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEnrollmentRequestUserCountOptions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Catalyst_Enrollment::getEnrollmentRequestYearsInOperationOptions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEnrollmentRequestYearsInOperationOptions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Catalyst_Enrollment::getIsActiveFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsActiveFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Catalyst_Enrollment::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Catalyst_Enrollment::getRepresentative", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRepresentative()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Catalyst_Enrollment::requestManualEnrollment", func() {
            It("API Call Test", func() {

				err := sl_service.RequestManualEnrollment(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Catalyst_Enrollment::requestSelfEnrollment", func() {
            It("API Call Test", func() {

                _, err := sl_service.RequestSelfEnrollment(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

})
