
package services_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/softlayer/softlayer-go/session/sessionfakes"
    "github.com/softlayer/softlayer-go/services"
)   

var _ = Describe("Software Tests", func() {
    var slsession *sessionfakes.FakeSLSession
    BeforeEach(func() {
        slsession = &sessionfakes.FakeSLSession{}
    })

    Context("Testing SoftLayer_Software_AccountLicense service", func() {
        var sl_service services.Software_AccountLicense
        BeforeEach(func() {
            sl_service = services.GetSoftwareAccountLicenseService(slsession)
        })


        Context("SoftLayer_Software_AccountLicense::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_AccountLicense::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_AccountLicense::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_AccountLicense::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_AccountLicense::getSoftwareDescription", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSoftwareDescription()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Software_Component service", func() {
        var sl_service services.Software_Component
        BeforeEach(func() {
            sl_service = services.GetSoftwareComponentService(slsession)
        })


        Context("SoftLayer_Software_Component::getAverageInstallationDuration", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAverageInstallationDuration()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component::getHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component::getLicenseFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLicenseFile()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component::getPasswordHistory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPasswordHistory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component::getPasswords", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPasswords()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component::getSoftwareDescription", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSoftwareDescription()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component::getSoftwareLicense", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSoftwareLicense()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component::getVendorSetUpConfiguration", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVendorSetUpConfiguration()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component::getVirtualGuest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Software_Component_AntivirusSpyware service", func() {
        var sl_service services.Software_Component_AntivirusSpyware
        BeforeEach(func() {
            sl_service = services.GetSoftwareComponentAntivirusSpywareService(slsession)
        })


        Context("SoftLayer_Software_Component_AntivirusSpyware::getAverageInstallationDuration", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAverageInstallationDuration()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_AntivirusSpyware::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_AntivirusSpyware::getHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_AntivirusSpyware::getLicenseFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLicenseFile()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_AntivirusSpyware::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_AntivirusSpyware::getPasswordHistory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPasswordHistory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_AntivirusSpyware::getPasswords", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPasswords()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_AntivirusSpyware::getSoftwareDescription", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSoftwareDescription()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_AntivirusSpyware::getSoftwareLicense", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSoftwareLicense()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_AntivirusSpyware::getVendorSetUpConfiguration", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVendorSetUpConfiguration()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_AntivirusSpyware::getVirtualGuest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_AntivirusSpyware::updateAntivirusSpywarePolicy", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateAntivirusSpywarePolicy(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Software_Component_HostIps service", func() {
        var sl_service services.Software_Component_HostIps
        BeforeEach(func() {
            sl_service = services.GetSoftwareComponentHostIpsService(slsession)
        })


        Context("SoftLayer_Software_Component_HostIps::getAverageInstallationDuration", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAverageInstallationDuration()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_HostIps::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_HostIps::getCurrentHostIpsPolicies", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentHostIpsPolicies()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_HostIps::getHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_HostIps::getLicenseFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLicenseFile()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_HostIps::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_HostIps::getPasswordHistory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPasswordHistory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_HostIps::getPasswords", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPasswords()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_HostIps::getSoftwareDescription", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSoftwareDescription()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_HostIps::getSoftwareLicense", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSoftwareLicense()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_HostIps::getVendorSetUpConfiguration", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVendorSetUpConfiguration()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_HostIps::getVirtualGuest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_HostIps::updateHipsPolicies", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateHipsPolicies(nil,nil,nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Software_Component_Password service", func() {
        var sl_service services.Software_Component_Password
        BeforeEach(func() {
            sl_service = services.GetSoftwareComponentPasswordService(slsession)
        })


        Context("SoftLayer_Software_Component_Password::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_Password::createObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_Password::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_Password::deleteObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_Password::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_Password::editObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_Password::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_Password::getSoftware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSoftware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_Password::getSshKeys", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSshKeys()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Software_Component_Trellix service", func() {
        var sl_service services.Software_Component_Trellix
        BeforeEach(func() {
            sl_service = services.GetSoftwareComponentTrellixService(slsession)
        })


        Context("SoftLayer_Software_Component_Trellix::getAverageInstallationDuration", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAverageInstallationDuration()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_Trellix::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_Trellix::getCurrentHostIpsPolicies", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentHostIpsPolicies()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_Trellix::getHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_Trellix::getLicenseFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLicenseFile()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_Trellix::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_Trellix::getPasswordHistory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPasswordHistory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_Trellix::getPasswords", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPasswords()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_Trellix::getSoftwareDescription", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSoftwareDescription()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_Trellix::getSoftwareLicense", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSoftwareLicense()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_Trellix::getVendorSetUpConfiguration", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVendorSetUpConfiguration()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_Trellix::getVirtualGuest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_Trellix::updateAntivirusSpywarePolicy", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateAntivirusSpywarePolicy(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Component_Trellix::updateHipsPolicies", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateHipsPolicies(nil,nil,nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Software_Description service", func() {
        var sl_service services.Software_Description
        BeforeEach(func() {
            sl_service = services.GetSoftwareDescriptionService(slsession)
        })


        Context("SoftLayer_Software_Description::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Description::getAttributes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttributes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Description::getAverageInstallationDuration", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAverageInstallationDuration()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Description::getCompatibleSoftwareDescriptions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCompatibleSoftwareDescriptions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Description::getCustomerOwnedLicenseDescriptions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCustomerOwnedLicenseDescriptions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Description::getFeatures", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFeatures()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Description::getLatestVersion", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLatestVersion()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Description::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Description::getProductItems", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProductItems()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Description::getProvisionTransactionGroup", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProvisionTransactionGroup()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Description::getReloadTransactionGroup", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReloadTransactionGroup()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Description::getRequiredUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRequiredUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Description::getSoftwareLicenses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSoftwareLicenses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Description::getUpgradeSoftwareDescription", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpgradeSoftwareDescription()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Description::getUpgradeSwDesc", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpgradeSwDesc()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_Description::getValidFilesystemTypes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetValidFilesystemTypes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Software_VirtualLicense service", func() {
        var sl_service services.Software_VirtualLicense
        BeforeEach(func() {
            sl_service = services.GetSoftwareVirtualLicenseService(slsession)
        })


        Context("SoftLayer_Software_VirtualLicense::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_VirtualLicense::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_VirtualLicense::getHostHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHostHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_VirtualLicense::getIpAddressRecord", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIpAddressRecord()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_VirtualLicense::getLicenseFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLicenseFile()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_VirtualLicense::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_VirtualLicense::getSoftwareDescription", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSoftwareDescription()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Software_VirtualLicense::getSubnet", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubnet()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

})
