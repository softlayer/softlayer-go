
package services_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/softlayer/softlayer-go/session/sessionfakes"
    "github.com/softlayer/softlayer-go/services"
)   

var _ = Describe("Resource Tests", func() {
    var slsession *sessionfakes.FakeSLSession
    BeforeEach(func() {
        slsession = &sessionfakes.FakeSLSession{}
    })

    Context("Testing SoftLayer_Resource_Configuration service", func() {
        var sl_service services.Resource_Configuration
        BeforeEach(func() {
            sl_service = services.GetResourceConfigurationService(slsession)
        })


        Context("SoftLayer_Resource_Configuration::setOsPasswordFromEncrypted", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetOsPasswordFromEncrypted(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Resource_Group service", func() {
        var sl_service services.Resource_Group
        BeforeEach(func() {
            sl_service = services.GetResourceGroupService(slsession)
        })


        Context("SoftLayer_Resource_Group::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Group::getAncestorGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAncestorGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Group::getAttributes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttributes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Group::getHardwareMembers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareMembers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Group::getMembers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMembers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Group::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Group::getRootResourceGroup", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRootResourceGroup()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Group::getSubnetMembers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubnetMembers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Group::getTemplate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTemplate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Group::getVlanMembers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVlanMembers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Resource_Group_Template service", func() {
        var sl_service services.Resource_Group_Template
        BeforeEach(func() {
            sl_service = services.GetResourceGroupTemplateService(slsession)
        })


        Context("SoftLayer_Resource_Group_Template::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Group_Template::getChildren", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetChildren()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Group_Template::getMembers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMembers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Group_Template::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Resource_Metadata service", func() {
        var sl_service services.Resource_Metadata
        BeforeEach(func() {
            sl_service = services.GetResourceMetadataService(slsession)
        })


        Context("SoftLayer_Resource_Metadata::getAccountId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccountId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Metadata::getBackendMacAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendMacAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Metadata::getDatacenter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDatacenter()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Metadata::getDatacenterId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDatacenterId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Metadata::getDomain", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDomain()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Metadata::getFrontendMacAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFrontendMacAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Metadata::getFullyQualifiedDomainName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFullyQualifiedDomainName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Metadata::getGlobalIdentifier", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGlobalIdentifier()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Metadata::getHostname", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHostname()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Metadata::getId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Metadata::getPrimaryBackendIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryBackendIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Metadata::getPrimaryIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Metadata::getProvisionState", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProvisionState()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Metadata::getRouter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRouter(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Metadata::getServiceResource", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceResource(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Metadata::getServiceResources", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceResources()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Metadata::getTags", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTags()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Metadata::getUserMetadata", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUserMetadata()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Metadata::getVlanIds", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVlanIds(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Resource_Metadata::getVlans", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVlans(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

})
