
package services_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/softlayer/softlayer-go/session/sessionfakes"
    "github.com/softlayer/softlayer-go/services"
)   

var _ = Describe("Locale Tests", func() {
    var slsession *sessionfakes.FakeSLSession
    BeforeEach(func() {
        slsession = &sessionfakes.FakeSLSession{}
    })

    Context("Testing SoftLayer_Locale service", func() {
        var sl_service services.Locale
        BeforeEach(func() {
            sl_service = services.GetLocaleService(slsession)
        })


        Context("SoftLayer_Locale::getClosestToLanguageTag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetClosestToLanguageTag(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Locale::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Locale_Country service", func() {
        var sl_service services.Locale_Country
        BeforeEach(func() {
            sl_service = services.GetLocaleCountryService(slsession)
        })


        Context("SoftLayer_Locale_Country::getAllVatCountryCodesAndVatIdRegexes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllVatCountryCodesAndVatIdRegexes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Locale_Country::getAvailableCountries", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailableCountries()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Locale_Country::getCountries", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCountries()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Locale_Country::getCountriesAndStates", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCountriesAndStates(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Locale_Country::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Locale_Country::getPostalCodeRequiredCountryCodes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPostalCodeRequiredCountryCodes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Locale_Country::getStates", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStates()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Locale_Country::getVatCountries", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVatCountries()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Locale_Country::getVatRequiredCountryCodes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVatRequiredCountryCodes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Locale_Country::isEuropeanUnionCountry", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsEuropeanUnionCountry(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Locale_Timezone service", func() {
        var sl_service services.Locale_Timezone
        BeforeEach(func() {
            sl_service = services.GetLocaleTimezoneService(slsession)
        })


        Context("SoftLayer_Locale_Timezone::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Locale_Timezone::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

})
