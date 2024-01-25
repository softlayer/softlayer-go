
package services_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/softlayer/softlayer-go/session/sessionfakes"
    "github.com/softlayer/softlayer-go/services"
)   

var _ = Describe("Dns Tests", func() {
    var slsession *sessionfakes.FakeSLSession
    BeforeEach(func() {
        slsession = &sessionfakes.FakeSLSession{}
    })

    Context("Testing SoftLayer_Dns_Domain service", func() {
        var sl_service services.Dns_Domain
        BeforeEach(func() {
            sl_service = services.GetDnsDomainService(slsession)
        })


        Context("SoftLayer_Dns_Domain::createARecord", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateARecord(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain::createAaaaRecord", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateAaaaRecord(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain::createCnameRecord", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateCnameRecord(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain::createMxRecord", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateMxRecord(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain::createNsRecord", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateNsRecord(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain::createObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain::createPtrRecord", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreatePtrRecord(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain::createSpfRecord", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateSpfRecord(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain::createTxtRecord", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateTxtRecord(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain::getByDomainName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetByDomainName(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain::getManagedResourceFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetManagedResourceFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain::getResourceRecords", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResourceRecords()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain::getSecondary", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSecondary()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain::getSoaResourceRecord", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSoaResourceRecord()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain::getZoneFileContents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetZoneFileContents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Dns_Domain_Registration service", func() {
        var sl_service services.Dns_Domain_Registration
        BeforeEach(func() {
            sl_service = services.GetDnsDomainRegistrationService(slsession)
        })


        Context("SoftLayer_Dns_Domain_Registration::addNameserversToDomain", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddNameserversToDomain(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_Registration::deleteRegisteredNameserver", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteRegisteredNameserver(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_Registration::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_Registration::getAuthenticationCode", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAuthenticationCode()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_Registration::getDomainInformation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDomainInformation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_Registration::getDomainNameservers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDomainNameservers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_Registration::getDomainRegistrationStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDomainRegistrationStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_Registration::getExtendedAttributes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetExtendedAttributes(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_Registration::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_Registration::getRegisteredNameserver", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRegisteredNameserver()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_Registration::getRegistrantVerificationStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRegistrantVerificationStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_Registration::getRegistrantVerificationStatusDetail", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRegistrantVerificationStatusDetail()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_Registration::getServiceProvider", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceProvider()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_Registration::getTransferInformation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTransferInformation(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_Registration::lockDomain", func() {
            It("API Call Test", func() {

                _, err := sl_service.LockDomain()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_Registration::lookupDomain", func() {
            It("API Call Test", func() {

                _, err := sl_service.LookupDomain(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_Registration::modifyContact", func() {
            It("API Call Test", func() {

                _, err := sl_service.ModifyContact(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_Registration::modifyRegisteredNameserver", func() {
            It("API Call Test", func() {

                _, err := sl_service.ModifyRegisteredNameserver(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_Registration::registerNameserver", func() {
            It("API Call Test", func() {

                _, err := sl_service.RegisterNameserver(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_Registration::removeNameserversFromDomain", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveNameserversFromDomain(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_Registration::sendAuthenticationCode", func() {
            It("API Call Test", func() {

                _, err := sl_service.SendAuthenticationCode()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_Registration::sendRegistrantVerificationEmail", func() {
            It("API Call Test", func() {

                _, err := sl_service.SendRegistrantVerificationEmail()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_Registration::sendTransferApprovalEmail", func() {
            It("API Call Test", func() {

                _, err := sl_service.SendTransferApprovalEmail()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_Registration::setAuthenticationCode", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetAuthenticationCode(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_Registration::unlockDomain", func() {
            It("API Call Test", func() {

                _, err := sl_service.UnlockDomain()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Dns_Domain_Registration_Registrant_Verification_Status service", func() {
        var sl_service services.Dns_Domain_Registration_Registrant_Verification_Status
        BeforeEach(func() {
            sl_service = services.GetDnsDomainRegistrationRegistrantVerificationStatusService(slsession)
        })


        Context("SoftLayer_Dns_Domain_Registration_Registrant_Verification_Status::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_Registration_Registrant_Verification_Status::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Dns_Domain_Registration_Status service", func() {
        var sl_service services.Dns_Domain_Registration_Status
        BeforeEach(func() {
            sl_service = services.GetDnsDomainRegistrationStatusService(slsession)
        })


        Context("SoftLayer_Dns_Domain_Registration_Status::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_Registration_Status::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Dns_Domain_ResourceRecord service", func() {
        var sl_service services.Dns_Domain_ResourceRecord
        BeforeEach(func() {
            sl_service = services.GetDnsDomainResourceRecordService(slsession)
        })


        Context("SoftLayer_Dns_Domain_ResourceRecord::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_ResourceRecord::createObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_ResourceRecord::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_ResourceRecord::deleteObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_ResourceRecord::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_ResourceRecord::editObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_ResourceRecord::getDomain", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDomain()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_ResourceRecord::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Dns_Domain_ResourceRecord_MxType service", func() {
        var sl_service services.Dns_Domain_ResourceRecord_MxType
        BeforeEach(func() {
            sl_service = services.GetDnsDomainResourceRecordMxTypeService(slsession)
        })


        Context("SoftLayer_Dns_Domain_ResourceRecord_MxType::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_ResourceRecord_MxType::createObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_ResourceRecord_MxType::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_ResourceRecord_MxType::deleteObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_ResourceRecord_MxType::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_ResourceRecord_MxType::editObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_ResourceRecord_MxType::getDomain", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDomain()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_ResourceRecord_MxType::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Dns_Domain_ResourceRecord_SrvType service", func() {
        var sl_service services.Dns_Domain_ResourceRecord_SrvType
        BeforeEach(func() {
            sl_service = services.GetDnsDomainResourceRecordSrvTypeService(slsession)
        })


        Context("SoftLayer_Dns_Domain_ResourceRecord_SrvType::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_ResourceRecord_SrvType::createObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_ResourceRecord_SrvType::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_ResourceRecord_SrvType::deleteObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_ResourceRecord_SrvType::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_ResourceRecord_SrvType::editObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_ResourceRecord_SrvType::getDomain", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDomain()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Domain_ResourceRecord_SrvType::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Dns_Secondary service", func() {
        var sl_service services.Dns_Secondary
        BeforeEach(func() {
            sl_service = services.GetDnsSecondaryService(slsession)
        })


        Context("SoftLayer_Dns_Secondary::convertToPrimary", func() {
            It("API Call Test", func() {

                _, err := sl_service.ConvertToPrimary()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Secondary::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Secondary::createObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Secondary::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Secondary::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Secondary::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Secondary::getByDomainName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetByDomainName(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Secondary::getDomain", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDomain()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Secondary::getErrorMessages", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetErrorMessages()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Secondary::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Secondary::getStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Dns_Secondary::transferNow", func() {
            It("API Call Test", func() {

                _, err := sl_service.TransferNow()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

})
