
package services_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/softlayer/softlayer-go/session/sessionfakes"
    "github.com/softlayer/softlayer-go/services"
)   

var _ = Describe("Security Tests", func() {
    var slsession *sessionfakes.FakeSLSession
    BeforeEach(func() {
        slsession = &sessionfakes.FakeSLSession{}
    })

    Context("Testing SoftLayer_Security_Certificate service", func() {
        var sl_service services.Security_Certificate
        BeforeEach(func() {
            sl_service = services.GetSecurityCertificateService(slsession)
        })


        Context("SoftLayer_Security_Certificate::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Certificate::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Certificate::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Certificate::findByCommonName", func() {
            It("API Call Test", func() {

                _, err := sl_service.FindByCommonName(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Certificate::getAssociatedServiceCount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssociatedServiceCount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Certificate::getLbaasListeners", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLbaasListeners()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Certificate::getLoadBalancerVirtualIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLoadBalancerVirtualIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Certificate::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Certificate::getPemFormat", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPemFormat()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Security_Certificate_Request service", func() {
        var sl_service services.Security_Certificate_Request
        BeforeEach(func() {
            sl_service = services.GetSecurityCertificateRequestService(slsession)
        })


        Context("SoftLayer_Security_Certificate_Request::cancelSslOrder", func() {
            It("API Call Test", func() {

                _, err := sl_service.CancelSslOrder()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Certificate_Request::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Certificate_Request::getAdministratorEmailDomains", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAdministratorEmailDomains(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Certificate_Request::getAdministratorEmailPrefixes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAdministratorEmailPrefixes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Certificate_Request::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Certificate_Request::getOrder", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOrder()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Certificate_Request::getOrderItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOrderItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Certificate_Request::getPreviousOrderData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPreviousOrderData()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Certificate_Request::getSslCertificateRequests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSslCertificateRequests(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Certificate_Request::getStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Certificate_Request::resendEmail", func() {
            It("API Call Test", func() {

                _, err := sl_service.ResendEmail(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Certificate_Request::validateCsr", func() {
            It("API Call Test", func() {

                _, err := sl_service.ValidateCsr(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Security_Certificate_Request_ServerType service", func() {
        var sl_service services.Security_Certificate_Request_ServerType
        BeforeEach(func() {
            sl_service = services.GetSecurityCertificateRequestServerTypeService(slsession)
        })


        Context("SoftLayer_Security_Certificate_Request_ServerType::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Certificate_Request_ServerType::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Security_Certificate_Request_Status service", func() {
        var sl_service services.Security_Certificate_Request_Status
        BeforeEach(func() {
            sl_service = services.GetSecurityCertificateRequestStatusService(slsession)
        })


        Context("SoftLayer_Security_Certificate_Request_Status::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Certificate_Request_Status::getSslRequestStatuses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSslRequestStatuses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Security_Ssh_Key service", func() {
        var sl_service services.Security_Ssh_Key
        BeforeEach(func() {
            sl_service = services.GetSecuritySshKeyService(slsession)
        })


        Context("SoftLayer_Security_Ssh_Key::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Ssh_Key::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Ssh_Key::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Ssh_Key::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Ssh_Key::getBlockDeviceTemplateGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBlockDeviceTemplateGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Ssh_Key::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Security_Ssh_Key::getSoftwarePasswords", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSoftwarePasswords()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

})
