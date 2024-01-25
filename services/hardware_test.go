
package services_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/softlayer/softlayer-go/session/sessionfakes"
    "github.com/softlayer/softlayer-go/services"
)   

var _ = Describe("Hardware Tests", func() {
    var slsession *sessionfakes.FakeSLSession
    BeforeEach(func() {
        slsession = &sessionfakes.FakeSLSession{}
    })

    Context("Testing SoftLayer_Hardware service", func() {
        var sl_service services.Hardware
        BeforeEach(func() {
            sl_service = services.GetHardwareService(slsession)
        })


        Context("SoftLayer_Hardware::allowAccessToNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToNetworkStorage(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::allowAccessToNetworkStorageList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToNetworkStorageList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::captureImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.CaptureImage(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::deleteSoftwareComponentPasswords", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteSoftwareComponentPasswords(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::deleteTag", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteTag(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::editSoftwareComponentPasswords", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditSoftwareComponentPasswords(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::executeRemoteScript", func() {
            It("API Call Test", func() {

				err := sl_service.ExecuteRemoteScript(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::findByIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.FindByIpAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::generateOrderTemplate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GenerateOrderTemplate(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getActiveComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getActiveNetworkMonitorIncident", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveNetworkMonitorIncident()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getAllPowerComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllPowerComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getAllowedHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedHost()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getAllowedNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getAllowedNetworkStorageReplicas", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedNetworkStorageReplicas()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getAntivirusSpywareSoftwareComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAntivirusSpywareSoftwareComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getAttachedNetworkStorages", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachedNetworkStorages(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getAttributes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttributes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getAvailableBillingTermChangePrices", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailableBillingTermChangePrices()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getAvailableNetworkStorages", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailableNetworkStorages(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getAverageDailyPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAverageDailyPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getBackendIncomingBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendIncomingBandwidth(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getBackendNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getBackendOutgoingBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendOutgoingBandwidth(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getBackendRouters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendRouters()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getBandwidthAllocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthAllocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getBandwidthAllotmentDetail", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthAllotmentDetail()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getBenchmarkCertifications", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBenchmarkCertifications()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getBillingItemFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItemFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getBlockCancelBecauseDisconnectedFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBlockCancelBecauseDisconnectedFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getBusinessContinuanceInsuranceFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBusinessContinuanceInsuranceFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getChildrenHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetChildrenHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getComponentDetailsXML", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetComponentDetailsXML()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getContinuousDataProtectionSoftwareComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetContinuousDataProtectionSoftwareComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getCreateObjectOptions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCreateObjectOptions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getCurrentBillableBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentBillableBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getCurrentBillingDetail", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentBillingDetail()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getCurrentBillingTotal", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentBillingTotal()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getDailyAverage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDailyAverage(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getDatacenter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDatacenter()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getDatacenterName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDatacenterName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getDaysInSparePool", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDaysInSparePool()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getDownlinkHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownlinkHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getDownlinkNetworkHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownlinkNetworkHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getDownlinkServers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownlinkServers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getDownlinkVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownlinkVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getDownstreamHardwareBindings", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownstreamHardwareBindings()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getDownstreamNetworkHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownstreamNetworkHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getDownstreamNetworkHardwareWithIncidents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownstreamNetworkHardwareWithIncidents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getDownstreamServers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownstreamServers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getDownstreamVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownstreamVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getDriveControllers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDriveControllers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getEvaultNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEvaultNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getFirewallServiceComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirewallServiceComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getFixedConfigurationPreset", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFixedConfigurationPreset()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getFrontendIncomingBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFrontendIncomingBandwidth(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getFrontendNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFrontendNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getFrontendOutgoingBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFrontendOutgoingBandwidth(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getFrontendRouters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFrontendRouters()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getFutureBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFutureBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getGlobalIdentifier", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGlobalIdentifier()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getHardDrives", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardDrives()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getHardwareChassis", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareChassis()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getHardwareFunction", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareFunction()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getHardwareFunctionDescription", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareFunctionDescription()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getHardwareState", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareState()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getHardwareStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getHasTrustedPlatformModuleBillingItemFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasTrustedPlatformModuleBillingItemFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getHostIpsSoftwareComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHostIpsSoftwareComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getHourlyBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHourlyBandwidth(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getHourlyBillingFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHourlyBillingFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getInboundBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInboundBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getInboundPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInboundPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getIsBillingTermChangeAvailableFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsBillingTermChangeAvailableFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getLastTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLastTransaction()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getLatestNetworkMonitorIncident", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLatestNetworkMonitorIncident()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getLocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getLocationPathString", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLocationPathString()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getLockboxNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLockboxNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getManagedResourceFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetManagedResourceFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getMemory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMemory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getMemoryCapacity", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMemoryCapacity()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getMetricTrackingObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMetricTrackingObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getModules", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetModules()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getMonitoringRobot", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMonitoringRobot()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getMonitoringServiceComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMonitoringServiceComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getMonitoringServiceEligibilityFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMonitoringServiceEligibilityFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getMotherboard", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMotherboard()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getNetworkCards", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkCards()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getNetworkGatewayMember", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkGatewayMember()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getNetworkGatewayMemberFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkGatewayMemberFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getNetworkManagementIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkManagementIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getNetworkMonitorAttachedDownHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitorAttachedDownHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getNetworkMonitorAttachedDownVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitorAttachedDownVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getNetworkMonitorIncidents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitorIncidents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getNetworkMonitors", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitors()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getNetworkStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getNetworkStatusAttribute", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkStatusAttribute()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getNetworkVlans", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkVlans()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getNextBillingCycleBandwidthAllocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNextBillingCycleBandwidthAllocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getNotesHistory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNotesHistory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getNvRamCapacity", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNvRamCapacity()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getNvRamComponentModels", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNvRamComponentModels()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getOperatingSystem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOperatingSystem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getOperatingSystemReferenceCode", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOperatingSystemReferenceCode()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getOutboundBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOutboundBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getOutboundPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOutboundPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getParentBay", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetParentBay()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getParentHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetParentHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getPointOfPresenceLocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPointOfPresenceLocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getPowerComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPowerComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getPowerSupply", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPowerSupply()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getPrimaryBackendIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryBackendIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getPrimaryBackendNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryBackendNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getPrimaryIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getPrimaryNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getPrivateBandwidthData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateBandwidthData(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getPrivateNetworkOnlyFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateNetworkOnlyFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getProcessorCoreAmount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProcessorCoreAmount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getProcessorPhysicalCoreAmount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProcessorPhysicalCoreAmount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getProcessors", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProcessors()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getPublicBandwidthData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicBandwidthData(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getRack", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRack()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getRaidControllers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRaidControllers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getRecentEvents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRecentEvents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getRemoteManagementAccounts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRemoteManagementAccounts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getRemoteManagementComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRemoteManagementComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getResourceConfigurations", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResourceConfigurations()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getResourceGroupMemberReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResourceGroupMemberReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getResourceGroupRoles", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResourceGroupRoles()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getResourceGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResourceGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getRouters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRouters()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getSecurityScanRequests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSecurityScanRequests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getSensorData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSensorData()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getSensorDataWithGraphs", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSensorDataWithGraphs()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getServerFanSpeedGraphs", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerFanSpeedGraphs()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getServerPowerState", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerPowerState()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getServerRoom", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerRoom()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getServerTemperatureGraphs", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerTemperatureGraphs()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getServiceProvider", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceProvider()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getSoftwareComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSoftwareComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getSparePoolBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSparePoolBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getSshKeys", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSshKeys()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getStorageGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getStorageNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getTagReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTagReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getTopLevelLocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTopLevelLocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getTransactionHistory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTransactionHistory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getUpgradeItemPrices", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpgradeItemPrices()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getUpgradeRequest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpgradeRequest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getUpgradeableActiveComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpgradeableActiveComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getUplinkHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUplinkHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getUplinkNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUplinkNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getUserData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUserData()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getVirtualChassis", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualChassis()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getVirtualChassisSiblings", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualChassisSiblings()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getVirtualHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualHost()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getVirtualLicenses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualLicenses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getVirtualRack", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualRack()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getVirtualRackId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualRackId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getVirtualRackName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualRackName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::getVirtualizationPlatform", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualizationPlatform()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::importVirtualHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.ImportVirtualHost()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::isPingable", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsPingable()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::ping", func() {
            It("API Call Test", func() {

                _, err := sl_service.Ping()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::powerCycle", func() {
            It("API Call Test", func() {

                _, err := sl_service.PowerCycle()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::powerOff", func() {
            It("API Call Test", func() {

                _, err := sl_service.PowerOff()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::powerOn", func() {
            It("API Call Test", func() {

                _, err := sl_service.PowerOn()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::rebootDefault", func() {
            It("API Call Test", func() {

                _, err := sl_service.RebootDefault()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::rebootHard", func() {
            It("API Call Test", func() {

                _, err := sl_service.RebootHard()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::rebootSoft", func() {
            It("API Call Test", func() {

                _, err := sl_service.RebootSoft()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::refreshDeviceStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.RefreshDeviceStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::removeAccessToNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToNetworkStorage(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::removeAccessToNetworkStorageList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToNetworkStorageList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::removeTags", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveTags(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::setTags", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetTags(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware::updateIpmiPassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateIpmiPassword(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Hardware_Benchmark_Certification service", func() {
        var sl_service services.Hardware_Benchmark_Certification
        BeforeEach(func() {
            sl_service = services.GetHardwareBenchmarkCertificationService(slsession)
        })


        Context("SoftLayer_Hardware_Benchmark_Certification::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Benchmark_Certification::getHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Benchmark_Certification::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Benchmark_Certification::getResultFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResultFile()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Hardware_Blade service", func() {
        var sl_service services.Hardware_Blade
        BeforeEach(func() {
            sl_service = services.GetHardwareBladeService(slsession)
        })


        Context("SoftLayer_Hardware_Blade::getHardwareChild", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareChild()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Blade::getHardwareParent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareParent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Blade::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Hardware_Component_Locator service", func() {
        var sl_service services.Hardware_Component_Locator
        BeforeEach(func() {
            sl_service = services.GetHardwareComponentLocatorService(slsession)
        })


        Context("SoftLayer_Hardware_Component_Locator::getGenericComponentModelAvailability", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGenericComponentModelAvailability(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Component_Locator::getPackageItemsAvailability", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPackageItemsAvailability(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Component_Locator::getServerPackageAvailability", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerPackageAvailability()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Hardware_Component_Model service", func() {
        var sl_service services.Hardware_Component_Model
        BeforeEach(func() {
            sl_service = services.GetHardwareComponentModelService(slsession)
        })


        Context("SoftLayer_Hardware_Component_Model::getArchitectureType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetArchitectureType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Component_Model::getAttributes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttributes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Component_Model::getCompatibleArrayTypes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCompatibleArrayTypes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Component_Model::getCompatibleChildComponentModels", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCompatibleChildComponentModels()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Component_Model::getCompatibleParentComponentModels", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCompatibleParentComponentModels()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Component_Model::getFirmwareQuantity", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirmwareQuantity()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Component_Model::getFirmwares", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirmwares()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Component_Model::getHardwareComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Component_Model::getHardwareGenericComponentModel", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareGenericComponentModel()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Component_Model::getInfinibandCompatibleAttribute", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInfinibandCompatibleAttribute()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Component_Model::getIsFlexSkuCompatible", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsFlexSkuCompatible()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Component_Model::getIsInfinibandCompatible", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsInfinibandCompatible()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Component_Model::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Component_Model::getQualifiedFirmwares", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetQualifiedFirmwares()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Component_Model::getRebootTime", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRebootTime()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Component_Model::getType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Component_Model::getValidAttributeTypes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetValidAttributeTypes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Hardware_Component_Partition_OperatingSystem service", func() {
        var sl_service services.Hardware_Component_Partition_OperatingSystem
        BeforeEach(func() {
            sl_service = services.GetHardwareComponentPartitionOperatingSystemService(slsession)
        })


        Context("SoftLayer_Hardware_Component_Partition_OperatingSystem::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Component_Partition_OperatingSystem::getByDescription", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetByDescription(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Component_Partition_OperatingSystem::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Component_Partition_OperatingSystem::getPartitionTemplates", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPartitionTemplates()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Hardware_Component_Partition_Template service", func() {
        var sl_service services.Hardware_Component_Partition_Template
        BeforeEach(func() {
            sl_service = services.GetHardwareComponentPartitionTemplateService(slsession)
        })


        Context("SoftLayer_Hardware_Component_Partition_Template::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Component_Partition_Template::getData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetData()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Component_Partition_Template::getExpireDate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetExpireDate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Component_Partition_Template::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Component_Partition_Template::getPartitionOperatingSystem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPartitionOperatingSystem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Component_Partition_Template::getPartitionTemplatePartition", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPartitionTemplatePartition()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Hardware_Router service", func() {
        var sl_service services.Hardware_Router
        BeforeEach(func() {
            sl_service = services.GetHardwareRouterService(slsession)
        })


        Context("SoftLayer_Hardware_Router::allowAccessToNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToNetworkStorage(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::allowAccessToNetworkStorageList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToNetworkStorageList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::captureImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.CaptureImage(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::deleteSoftwareComponentPasswords", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteSoftwareComponentPasswords(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::deleteTag", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteTag(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::editSoftwareComponentPasswords", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditSoftwareComponentPasswords(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::executeRemoteScript", func() {
            It("API Call Test", func() {

				err := sl_service.ExecuteRemoteScript(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::findByIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.FindByIpAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::generateOrderTemplate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GenerateOrderTemplate(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getActiveComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getActiveNetworkMonitorIncident", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveNetworkMonitorIncident()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getAllPowerComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllPowerComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getAllowedHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedHost()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getAllowedNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getAllowedNetworkStorageReplicas", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedNetworkStorageReplicas()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getAntivirusSpywareSoftwareComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAntivirusSpywareSoftwareComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getAttachedNetworkStorages", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachedNetworkStorages(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getAttributes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttributes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getAvailableBillingTermChangePrices", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailableBillingTermChangePrices()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getAvailableNetworkStorages", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailableNetworkStorages(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getAverageDailyPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAverageDailyPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getBackendIncomingBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendIncomingBandwidth(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getBackendNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getBackendOutgoingBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendOutgoingBandwidth(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getBackendRouters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendRouters()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getBandwidthAllocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthAllocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getBandwidthAllotmentDetail", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthAllotmentDetail()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getBenchmarkCertifications", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBenchmarkCertifications()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getBillingItemFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItemFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getBlockCancelBecauseDisconnectedFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBlockCancelBecauseDisconnectedFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getBoundSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBoundSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getBusinessContinuanceInsuranceFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBusinessContinuanceInsuranceFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getChildrenHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetChildrenHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getComponentDetailsXML", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetComponentDetailsXML()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getContinuousDataProtectionSoftwareComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetContinuousDataProtectionSoftwareComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getCreateObjectOptions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCreateObjectOptions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getCurrentBillableBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentBillableBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getCurrentBillingDetail", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentBillingDetail()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getCurrentBillingTotal", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentBillingTotal()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getDailyAverage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDailyAverage(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getDatacenter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDatacenter()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getDatacenterName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDatacenterName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getDaysInSparePool", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDaysInSparePool()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getDownlinkHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownlinkHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getDownlinkNetworkHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownlinkNetworkHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getDownlinkServers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownlinkServers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getDownlinkVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownlinkVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getDownstreamHardwareBindings", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownstreamHardwareBindings()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getDownstreamNetworkHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownstreamNetworkHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getDownstreamNetworkHardwareWithIncidents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownstreamNetworkHardwareWithIncidents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getDownstreamServers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownstreamServers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getDownstreamVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownstreamVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getDriveControllers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDriveControllers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getEvaultNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEvaultNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getFirewallServiceComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirewallServiceComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getFixedConfigurationPreset", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFixedConfigurationPreset()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getFrontendIncomingBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFrontendIncomingBandwidth(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getFrontendNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFrontendNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getFrontendOutgoingBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFrontendOutgoingBandwidth(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getFrontendRouters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFrontendRouters()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getFutureBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFutureBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getGlobalIdentifier", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGlobalIdentifier()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getHardDrives", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardDrives()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getHardwareChassis", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareChassis()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getHardwareFunction", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareFunction()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getHardwareFunctionDescription", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareFunctionDescription()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getHardwareState", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareState()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getHardwareStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getHasTrustedPlatformModuleBillingItemFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasTrustedPlatformModuleBillingItemFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getHostIpsSoftwareComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHostIpsSoftwareComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getHourlyBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHourlyBandwidth(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getHourlyBillingFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHourlyBillingFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getInboundBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInboundBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getInboundPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInboundPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getIsBillingTermChangeAvailableFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsBillingTermChangeAvailableFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getLastTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLastTransaction()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getLatestNetworkMonitorIncident", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLatestNetworkMonitorIncident()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getLocalDiskStorageCapabilityFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLocalDiskStorageCapabilityFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getLocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getLocationPathString", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLocationPathString()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getLockboxNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLockboxNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getManagedResourceFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetManagedResourceFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getMemory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMemory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getMemoryCapacity", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMemoryCapacity()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getMetricTrackingObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMetricTrackingObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getModules", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetModules()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getMonitoringRobot", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMonitoringRobot()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getMonitoringServiceComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMonitoringServiceComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getMonitoringServiceEligibilityFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMonitoringServiceEligibilityFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getMotherboard", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMotherboard()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getNetworkCards", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkCards()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getNetworkGatewayMember", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkGatewayMember()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getNetworkGatewayMemberFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkGatewayMemberFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getNetworkManagementIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkManagementIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getNetworkMonitorAttachedDownHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitorAttachedDownHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getNetworkMonitorAttachedDownVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitorAttachedDownVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getNetworkMonitorIncidents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitorIncidents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getNetworkMonitors", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitors()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getNetworkStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getNetworkStatusAttribute", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkStatusAttribute()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getNetworkVlans", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkVlans()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getNextBillingCycleBandwidthAllocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNextBillingCycleBandwidthAllocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getNotesHistory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNotesHistory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getNvRamCapacity", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNvRamCapacity()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getNvRamComponentModels", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNvRamComponentModels()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getOperatingSystem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOperatingSystem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getOperatingSystemReferenceCode", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOperatingSystemReferenceCode()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getOutboundBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOutboundBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getOutboundPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOutboundPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getParentBay", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetParentBay()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getParentHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetParentHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getPointOfPresenceLocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPointOfPresenceLocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getPowerComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPowerComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getPowerSupply", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPowerSupply()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getPrimaryBackendIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryBackendIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getPrimaryBackendNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryBackendNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getPrimaryIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getPrimaryNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getPrivateBandwidthData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateBandwidthData(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getPrivateNetworkOnlyFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateNetworkOnlyFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getProcessorCoreAmount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProcessorCoreAmount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getProcessorPhysicalCoreAmount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProcessorPhysicalCoreAmount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getProcessors", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProcessors()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getPublicBandwidthData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicBandwidthData(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getRack", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRack()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getRaidControllers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRaidControllers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getRecentEvents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRecentEvents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getRemoteManagementAccounts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRemoteManagementAccounts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getRemoteManagementComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRemoteManagementComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getResourceConfigurations", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResourceConfigurations()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getResourceGroupMemberReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResourceGroupMemberReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getResourceGroupRoles", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResourceGroupRoles()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getResourceGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResourceGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getRouters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRouters()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getSanStorageCapabilityFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSanStorageCapabilityFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getSecurityScanRequests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSecurityScanRequests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getSensorData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSensorData()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getSensorDataWithGraphs", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSensorDataWithGraphs()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getServerFanSpeedGraphs", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerFanSpeedGraphs()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getServerPowerState", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerPowerState()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getServerRoom", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerRoom()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getServerTemperatureGraphs", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerTemperatureGraphs()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getServiceProvider", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceProvider()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getSoftwareComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSoftwareComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getSparePoolBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSparePoolBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getSshKeys", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSshKeys()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getStorageGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getStorageNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getTagReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTagReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getTopLevelLocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTopLevelLocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getTransactionHistory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTransactionHistory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getUpgradeItemPrices", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpgradeItemPrices()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getUpgradeRequest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpgradeRequest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getUpgradeableActiveComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpgradeableActiveComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getUplinkHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUplinkHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getUplinkNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUplinkNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getUserData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUserData()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getVirtualChassis", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualChassis()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getVirtualChassisSiblings", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualChassisSiblings()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getVirtualHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualHost()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getVirtualLicenses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualLicenses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getVirtualRack", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualRack()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getVirtualRackId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualRackId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getVirtualRackName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualRackName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::getVirtualizationPlatform", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualizationPlatform()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::importVirtualHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.ImportVirtualHost()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::isPingable", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsPingable()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::ping", func() {
            It("API Call Test", func() {

                _, err := sl_service.Ping()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::powerCycle", func() {
            It("API Call Test", func() {

                _, err := sl_service.PowerCycle()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::powerOff", func() {
            It("API Call Test", func() {

                _, err := sl_service.PowerOff()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::powerOn", func() {
            It("API Call Test", func() {

                _, err := sl_service.PowerOn()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::rebootDefault", func() {
            It("API Call Test", func() {

                _, err := sl_service.RebootDefault()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::rebootHard", func() {
            It("API Call Test", func() {

                _, err := sl_service.RebootHard()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::rebootSoft", func() {
            It("API Call Test", func() {

                _, err := sl_service.RebootSoft()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::refreshDeviceStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.RefreshDeviceStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::removeAccessToNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToNetworkStorage(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::removeAccessToNetworkStorageList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToNetworkStorageList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::removeTags", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveTags(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::setTags", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetTags(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Router::updateIpmiPassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateIpmiPassword(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Hardware_SecurityModule service", func() {
        var sl_service services.Hardware_SecurityModule
        BeforeEach(func() {
            sl_service = services.GetHardwareSecurityModuleService(slsession)
        })


        Context("SoftLayer_Hardware_SecurityModule::activatePrivatePort", func() {
            It("API Call Test", func() {

                _, err := sl_service.ActivatePrivatePort()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::activatePublicPort", func() {
            It("API Call Test", func() {

                _, err := sl_service.ActivatePublicPort()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::allowAccessToNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToNetworkStorage(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::allowAccessToNetworkStorageList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToNetworkStorageList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::bootToRescueLayer", func() {
            It("API Call Test", func() {

                _, err := sl_service.BootToRescueLayer(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::captureImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.CaptureImage(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::createFirmwareReflashTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateFirmwareReflashTransaction(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::createFirmwareUpdateTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateFirmwareUpdateTransaction(nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::createHyperThreadingUpdateTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateHyperThreadingUpdateTransaction(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::createPostSoftwareInstallTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreatePostSoftwareInstallTransaction(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::deleteSoftwareComponentPasswords", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteSoftwareComponentPasswords(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::deleteTag", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteTag(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::editSoftwareComponentPasswords", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditSoftwareComponentPasswords(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::executeRemoteScript", func() {
            It("API Call Test", func() {

				err := sl_service.ExecuteRemoteScript(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::findByIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.FindByIpAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::generateOrderTemplate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GenerateOrderTemplate(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getActiveComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getActiveNetworkFirewallBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveNetworkFirewallBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getActiveNetworkMonitorIncident", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveNetworkMonitorIncident()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getActiveTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getActiveTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveTransaction()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getActiveTransactions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveTransactions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getAllPowerComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllPowerComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getAllowedHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedHost()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getAllowedNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getAllowedNetworkStorageReplicas", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedNetworkStorageReplicas()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getAntivirusSpywareSoftwareComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAntivirusSpywareSoftwareComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getAttachedNetworkStorages", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachedNetworkStorages(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getAttributes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttributes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getAvailableBillingTermChangePrices", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailableBillingTermChangePrices()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getAvailableMonitoring", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailableMonitoring()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getAvailableNetworkStorages", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailableNetworkStorages(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getAverageDailyBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAverageDailyBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getAverageDailyPrivateBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAverageDailyPrivateBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getAverageDailyPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAverageDailyPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getBackendBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendBandwidthUsage(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getBackendIncomingBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendIncomingBandwidth(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getBackendNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getBackendOutgoingBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendOutgoingBandwidth(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getBackendRouters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendRouters()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getBandwidthAllocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthAllocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getBandwidthAllotmentDetail", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthAllotmentDetail()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getBandwidthForDateRange", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthForDateRange(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getBandwidthImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthImage(nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getBenchmarkCertifications", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBenchmarkCertifications()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getBillingCycleBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingCycleBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getBillingCyclePrivateBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingCyclePrivateBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getBillingCyclePublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingCyclePublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getBillingItemFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItemFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getBiosPasswordNullFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBiosPasswordNullFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getBlockCancelBecauseDisconnectedFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBlockCancelBecauseDisconnectedFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getBootModeOptions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBootModeOptions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getBusinessContinuanceInsuranceFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBusinessContinuanceInsuranceFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getCaptureEnabledFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCaptureEnabledFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getChildrenHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetChildrenHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getComponentDetailsXML", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetComponentDetailsXML()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getContainsSolidStateDrivesFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetContainsSolidStateDrivesFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getContinuousDataProtectionSoftwareComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetContinuousDataProtectionSoftwareComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getControlPanel", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetControlPanel()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getCost", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCost()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getCreateObjectOptions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCreateObjectOptions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getCurrentBandwidthSummary", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentBandwidthSummary()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getCurrentBenchmarkCertificationResultFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentBenchmarkCertificationResultFile()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getCurrentBillableBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentBillableBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getCurrentBillingDetail", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentBillingDetail()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getCurrentBillingTotal", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentBillingTotal()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getCustomerInstalledOperatingSystemFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCustomerInstalledOperatingSystemFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getCustomerOwnedFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCustomerOwnedFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getDailyAverage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDailyAverage(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getDatacenter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDatacenter()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getDatacenterName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDatacenterName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getDaysInSparePool", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDaysInSparePool()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getDownlinkHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownlinkHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getDownlinkNetworkHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownlinkNetworkHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getDownlinkServers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownlinkServers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getDownlinkVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownlinkVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getDownstreamHardwareBindings", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownstreamHardwareBindings()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getDownstreamNetworkHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownstreamNetworkHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getDownstreamNetworkHardwareWithIncidents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownstreamNetworkHardwareWithIncidents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getDownstreamServers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownstreamServers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getDownstreamVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownstreamVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getDriveControllers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDriveControllers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getEvaultNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEvaultNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getFirewallProtectableSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirewallProtectableSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getFirewallServiceComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirewallServiceComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getFixedConfigurationPreset", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFixedConfigurationPreset()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getFrontendBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFrontendBandwidthUsage(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getFrontendIncomingBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFrontendIncomingBandwidth(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getFrontendNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFrontendNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getFrontendOutgoingBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFrontendOutgoingBandwidth(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getFrontendRouters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFrontendRouters()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getFutureBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFutureBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getGlobalIdentifier", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGlobalIdentifier()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getHardDrives", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardDrives()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getHardwareByIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareByIpAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getHardwareChassis", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareChassis()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getHardwareFunction", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareFunction()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getHardwareFunctionDescription", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareFunctionDescription()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getHardwareState", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareState()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getHardwareStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getHasSingleRootVirtualizationBillingItemFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasSingleRootVirtualizationBillingItemFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getHasTrustedPlatformModuleBillingItemFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasTrustedPlatformModuleBillingItemFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getHostIpsSoftwareComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHostIpsSoftwareComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getHourlyBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHourlyBandwidth(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getHourlyBillingFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHourlyBillingFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getInboundBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInboundBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getInboundPrivateBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInboundPrivateBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getInboundPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInboundPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getIsBillingTermChangeAvailableFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsBillingTermChangeAvailableFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getIsCloudReadyNodeCertified", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsCloudReadyNodeCertified()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getIsIpmiDisabled", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsIpmiDisabled()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getIsQeInternalServer", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsQeInternalServer()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getIsVirtualPrivateCloudNode", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsVirtualPrivateCloudNode()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getItemPricesFromSoftwareDescriptions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetItemPricesFromSoftwareDescriptions(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getLastOperatingSystemReload", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLastOperatingSystemReload()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getLastTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLastTransaction()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getLatestNetworkMonitorIncident", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLatestNetworkMonitorIncident()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getLocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getLocationPathString", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLocationPathString()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getLockboxNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLockboxNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getLogicalVolumeStorageGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLogicalVolumeStorageGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getManagedResourceFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetManagedResourceFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getManagementNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetManagementNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getMemory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMemory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getMemoryCapacity", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMemoryCapacity()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getMetricTrackingObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMetricTrackingObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getMetricTrackingObjectId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMetricTrackingObjectId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getModules", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetModules()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getMonitoringRobot", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMonitoringRobot()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getMonitoringServiceComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMonitoringServiceComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getMonitoringServiceEligibilityFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMonitoringServiceEligibilityFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getMonitoringUserNotification", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMonitoringUserNotification()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getMotherboard", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMotherboard()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getNetworkCards", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkCards()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getNetworkComponentFirewallProtectableIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkComponentFirewallProtectableIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getNetworkGatewayMember", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkGatewayMember()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getNetworkGatewayMemberFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkGatewayMemberFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getNetworkManagementIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkManagementIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getNetworkMonitorAttachedDownHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitorAttachedDownHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getNetworkMonitorAttachedDownVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitorAttachedDownVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getNetworkMonitorIncidents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitorIncidents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getNetworkMonitors", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitors()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getNetworkStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getNetworkStatusAttribute", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkStatusAttribute()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getNetworkVlans", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkVlans()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getNextBillingCycleBandwidthAllocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNextBillingCycleBandwidthAllocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getNotesHistory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNotesHistory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getNvRamCapacity", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNvRamCapacity()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getNvRamComponentModels", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNvRamComponentModels()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getOpenCancellationTicket", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOpenCancellationTicket()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getOperatingSystem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOperatingSystem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getOperatingSystemReferenceCode", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOperatingSystemReferenceCode()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getOutboundBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOutboundBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getOutboundPrivateBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOutboundPrivateBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getOutboundPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOutboundPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getOverBandwidthAllocationFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOverBandwidthAllocationFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getPMInfo", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPMInfo()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getParentBay", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetParentBay()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getParentHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetParentHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getPartitions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPartitions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getPointOfPresenceLocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPointOfPresenceLocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getPowerComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPowerComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getPowerSupply", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPowerSupply()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getPrimaryBackendIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryBackendIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getPrimaryBackendNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryBackendNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getPrimaryDriveSize", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryDriveSize()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getPrimaryIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getPrimaryNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getPrivateBackendNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateBackendNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getPrivateBandwidthData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateBandwidthData(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getPrivateBandwidthDataSummary", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateBandwidthDataSummary()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getPrivateBandwidthGraphImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateBandwidthGraphImage(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getPrivateIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getPrivateNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getPrivateNetworkOnlyFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateNetworkOnlyFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getPrivateVlan", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateVlan()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getPrivateVlanByIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateVlanByIpAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getProcessorCoreAmount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProcessorCoreAmount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getProcessorPhysicalCoreAmount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProcessorPhysicalCoreAmount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getProcessors", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProcessors()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getProjectedOverBandwidthAllocationFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProjectedOverBandwidthAllocationFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getProjectedPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProjectedPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getProvisionDate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProvisionDate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getPublicBandwidthData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicBandwidthData(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getPublicBandwidthDataSummary", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicBandwidthDataSummary()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getPublicBandwidthGraphImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicBandwidthGraphImage(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getPublicBandwidthTotal", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicBandwidthTotal(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getPublicNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getPublicVlan", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicVlan()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getPublicVlanByHostname", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicVlanByHostname(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getRack", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRack()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getRaidControllers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRaidControllers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getReadyNodeFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReadyNodeFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getRecentEvents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRecentEvents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getRecentRemoteManagementCommands", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRecentRemoteManagementCommands()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getRegionalInternetRegistry", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRegionalInternetRegistry()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getRemoteManagement", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRemoteManagement()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getRemoteManagementAccounts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRemoteManagementAccounts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getRemoteManagementComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRemoteManagementComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getRemoteManagementUsers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRemoteManagementUsers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getResourceConfigurations", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResourceConfigurations()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getResourceGroupMemberReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResourceGroupMemberReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getResourceGroupRoles", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResourceGroupRoles()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getResourceGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResourceGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getReverseDomainRecords", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReverseDomainRecords()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getRouters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRouters()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getSecurityScanRequests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSecurityScanRequests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getSensorData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSensorData()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getSensorDataWithGraphs", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSensorDataWithGraphs()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getServerDetails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerDetails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getServerFanSpeedGraphs", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerFanSpeedGraphs()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getServerPowerState", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerPowerState()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getServerRoom", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerRoom()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getServerTemperatureGraphs", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerTemperatureGraphs()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getServiceProvider", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceProvider()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getSoftwareComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSoftwareComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getSoftwareGuardExtensionEnabled", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSoftwareGuardExtensionEnabled()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getSparePoolBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSparePoolBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getSshKeys", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSshKeys()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getStatisticsRemoteManagement", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStatisticsRemoteManagement()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getStorageGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getStorageNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getTagReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTagReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getTopLevelLocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTopLevelLocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getTransactionHistory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTransactionHistory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getUefiBootFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUefiBootFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getUpgradeItemPrices", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpgradeItemPrices()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getUpgradeRequest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpgradeRequest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getUpgradeableActiveComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpgradeableActiveComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getUplinkHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUplinkHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getUplinkNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUplinkNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getUserData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUserData()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getUsers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUsers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getValidBlockDeviceTemplateGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetValidBlockDeviceTemplateGroups(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getVirtualChassis", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualChassis()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getVirtualChassisSiblings", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualChassisSiblings()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getVirtualHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualHost()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getVirtualLicenses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualLicenses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getVirtualRack", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualRack()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getVirtualRackId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualRackId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getVirtualRackName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualRackName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getVirtualizationPlatform", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualizationPlatform()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getWindowsUpdateAvailableUpdates", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetWindowsUpdateAvailableUpdates()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getWindowsUpdateInstalledUpdates", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetWindowsUpdateInstalledUpdates()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::getWindowsUpdateStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetWindowsUpdateStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::importVirtualHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.ImportVirtualHost()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::initiateIderaBareMetalRestore", func() {
            It("API Call Test", func() {

                _, err := sl_service.InitiateIderaBareMetalRestore()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::initiateR1SoftBareMetalRestore", func() {
            It("API Call Test", func() {

                _, err := sl_service.InitiateR1SoftBareMetalRestore()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::isBackendPingable", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsBackendPingable()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::isPingable", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsPingable()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::isWindowsServer", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsWindowsServer()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::massFirmwareReflash", func() {
            It("API Call Test", func() {

                _, err := sl_service.MassFirmwareReflash(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::massFirmwareUpdate", func() {
            It("API Call Test", func() {

                _, err := sl_service.MassFirmwareUpdate(nil,nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::massHyperThreadingUpdate", func() {
            It("API Call Test", func() {

                _, err := sl_service.MassHyperThreadingUpdate(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::massReloadOperatingSystem", func() {
            It("API Call Test", func() {

                _, err := sl_service.MassReloadOperatingSystem(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::massSparePool", func() {
            It("API Call Test", func() {

                _, err := sl_service.MassSparePool(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::ping", func() {
            It("API Call Test", func() {

                _, err := sl_service.Ping()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::populateServerRam", func() {
            It("API Call Test", func() {

				err := sl_service.PopulateServerRam(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::powerCycle", func() {
            It("API Call Test", func() {

                _, err := sl_service.PowerCycle()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::powerOff", func() {
            It("API Call Test", func() {

                _, err := sl_service.PowerOff()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::powerOn", func() {
            It("API Call Test", func() {

                _, err := sl_service.PowerOn()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::rebootDefault", func() {
            It("API Call Test", func() {

                _, err := sl_service.RebootDefault()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::rebootHard", func() {
            It("API Call Test", func() {

                _, err := sl_service.RebootHard()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::rebootSoft", func() {
            It("API Call Test", func() {

                _, err := sl_service.RebootSoft()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::refreshDeviceStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.RefreshDeviceStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::reloadCurrentOperatingSystemConfiguration", func() {
            It("API Call Test", func() {

                _, err := sl_service.ReloadCurrentOperatingSystemConfiguration(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::reloadOperatingSystem", func() {
            It("API Call Test", func() {

                _, err := sl_service.ReloadOperatingSystem(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::removeAccessToNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToNetworkStorage(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::removeAccessToNetworkStorageList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToNetworkStorageList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::removeTags", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveTags(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::runPassmarkCertificationBenchmark", func() {
            It("API Call Test", func() {

                _, err := sl_service.RunPassmarkCertificationBenchmark()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::setOperatingSystemPassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetOperatingSystemPassword(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::setPrivateNetworkInterfaceSpeed", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetPrivateNetworkInterfaceSpeed(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::setPublicNetworkInterfaceSpeed", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetPublicNetworkInterfaceSpeed(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::setTags", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetTags(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::setUserMetadata", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetUserMetadata(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::shutdownPrivatePort", func() {
            It("API Call Test", func() {

                _, err := sl_service.ShutdownPrivatePort()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::shutdownPublicPort", func() {
            It("API Call Test", func() {

                _, err := sl_service.ShutdownPublicPort()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::sparePool", func() {
            It("API Call Test", func() {

                _, err := sl_service.SparePool(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::testRaidAlertService", func() {
            It("API Call Test", func() {

                _, err := sl_service.TestRaidAlertService()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::toggleManagementInterface", func() {
            It("API Call Test", func() {

                _, err := sl_service.ToggleManagementInterface(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::updateIpmiPassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateIpmiPassword(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule::validatePartitionsForOperatingSystem", func() {
            It("API Call Test", func() {

                _, err := sl_service.ValidatePartitionsForOperatingSystem(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Hardware_SecurityModule750 service", func() {
        var sl_service services.Hardware_SecurityModule750
        BeforeEach(func() {
            sl_service = services.GetHardwareSecurityModule750Service(slsession)
        })


        Context("SoftLayer_Hardware_SecurityModule750::activatePrivatePort", func() {
            It("API Call Test", func() {

                _, err := sl_service.ActivatePrivatePort()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::activatePublicPort", func() {
            It("API Call Test", func() {

                _, err := sl_service.ActivatePublicPort()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::allowAccessToNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToNetworkStorage(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::allowAccessToNetworkStorageList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToNetworkStorageList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::bootToRescueLayer", func() {
            It("API Call Test", func() {

                _, err := sl_service.BootToRescueLayer(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::captureImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.CaptureImage(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::createFirmwareReflashTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateFirmwareReflashTransaction(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::createFirmwareUpdateTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateFirmwareUpdateTransaction(nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::createHyperThreadingUpdateTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateHyperThreadingUpdateTransaction(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::createPostSoftwareInstallTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreatePostSoftwareInstallTransaction(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::deleteSoftwareComponentPasswords", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteSoftwareComponentPasswords(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::deleteTag", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteTag(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::editSoftwareComponentPasswords", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditSoftwareComponentPasswords(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::executeRemoteScript", func() {
            It("API Call Test", func() {

				err := sl_service.ExecuteRemoteScript(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::findByIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.FindByIpAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::generateOrderTemplate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GenerateOrderTemplate(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getActiveComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getActiveNetworkFirewallBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveNetworkFirewallBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getActiveNetworkMonitorIncident", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveNetworkMonitorIncident()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getActiveTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getActiveTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveTransaction()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getActiveTransactions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveTransactions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getAllPowerComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllPowerComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getAllowedHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedHost()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getAllowedNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getAllowedNetworkStorageReplicas", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedNetworkStorageReplicas()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getAntivirusSpywareSoftwareComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAntivirusSpywareSoftwareComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getAttachedNetworkStorages", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachedNetworkStorages(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getAttributes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttributes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getAvailableBillingTermChangePrices", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailableBillingTermChangePrices()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getAvailableMonitoring", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailableMonitoring()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getAvailableNetworkStorages", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailableNetworkStorages(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getAverageDailyBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAverageDailyBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getAverageDailyPrivateBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAverageDailyPrivateBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getAverageDailyPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAverageDailyPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getBackendBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendBandwidthUsage(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getBackendIncomingBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendIncomingBandwidth(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getBackendNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getBackendOutgoingBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendOutgoingBandwidth(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getBackendRouters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendRouters()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getBandwidthAllocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthAllocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getBandwidthAllotmentDetail", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthAllotmentDetail()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getBandwidthForDateRange", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthForDateRange(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getBandwidthImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthImage(nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getBenchmarkCertifications", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBenchmarkCertifications()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getBillingCycleBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingCycleBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getBillingCyclePrivateBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingCyclePrivateBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getBillingCyclePublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingCyclePublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getBillingItemFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItemFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getBiosPasswordNullFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBiosPasswordNullFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getBlockCancelBecauseDisconnectedFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBlockCancelBecauseDisconnectedFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getBootModeOptions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBootModeOptions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getBusinessContinuanceInsuranceFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBusinessContinuanceInsuranceFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getCaptureEnabledFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCaptureEnabledFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getChildrenHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetChildrenHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getComponentDetailsXML", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetComponentDetailsXML()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getContainsSolidStateDrivesFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetContainsSolidStateDrivesFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getContinuousDataProtectionSoftwareComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetContinuousDataProtectionSoftwareComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getControlPanel", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetControlPanel()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getCost", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCost()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getCreateObjectOptions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCreateObjectOptions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getCurrentBandwidthSummary", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentBandwidthSummary()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getCurrentBenchmarkCertificationResultFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentBenchmarkCertificationResultFile()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getCurrentBillableBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentBillableBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getCurrentBillingDetail", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentBillingDetail()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getCurrentBillingTotal", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentBillingTotal()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getCustomerInstalledOperatingSystemFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCustomerInstalledOperatingSystemFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getCustomerOwnedFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCustomerOwnedFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getDailyAverage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDailyAverage(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getDatacenter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDatacenter()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getDatacenterName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDatacenterName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getDaysInSparePool", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDaysInSparePool()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getDownlinkHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownlinkHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getDownlinkNetworkHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownlinkNetworkHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getDownlinkServers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownlinkServers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getDownlinkVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownlinkVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getDownstreamHardwareBindings", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownstreamHardwareBindings()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getDownstreamNetworkHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownstreamNetworkHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getDownstreamNetworkHardwareWithIncidents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownstreamNetworkHardwareWithIncidents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getDownstreamServers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownstreamServers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getDownstreamVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownstreamVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getDriveControllers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDriveControllers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getEvaultNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEvaultNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getFirewallProtectableSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirewallProtectableSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getFirewallServiceComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirewallServiceComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getFixedConfigurationPreset", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFixedConfigurationPreset()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getFrontendBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFrontendBandwidthUsage(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getFrontendIncomingBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFrontendIncomingBandwidth(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getFrontendNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFrontendNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getFrontendOutgoingBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFrontendOutgoingBandwidth(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getFrontendRouters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFrontendRouters()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getFutureBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFutureBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getGlobalIdentifier", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGlobalIdentifier()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getHardDrives", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardDrives()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getHardwareByIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareByIpAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getHardwareChassis", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareChassis()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getHardwareFunction", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareFunction()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getHardwareFunctionDescription", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareFunctionDescription()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getHardwareState", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareState()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getHardwareStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getHasSingleRootVirtualizationBillingItemFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasSingleRootVirtualizationBillingItemFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getHasTrustedPlatformModuleBillingItemFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasTrustedPlatformModuleBillingItemFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getHostIpsSoftwareComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHostIpsSoftwareComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getHourlyBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHourlyBandwidth(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getHourlyBillingFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHourlyBillingFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getInboundBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInboundBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getInboundPrivateBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInboundPrivateBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getInboundPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInboundPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getIsBillingTermChangeAvailableFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsBillingTermChangeAvailableFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getIsCloudReadyNodeCertified", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsCloudReadyNodeCertified()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getIsIpmiDisabled", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsIpmiDisabled()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getIsQeInternalServer", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsQeInternalServer()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getIsVirtualPrivateCloudNode", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsVirtualPrivateCloudNode()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getItemPricesFromSoftwareDescriptions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetItemPricesFromSoftwareDescriptions(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getLastOperatingSystemReload", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLastOperatingSystemReload()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getLastTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLastTransaction()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getLatestNetworkMonitorIncident", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLatestNetworkMonitorIncident()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getLocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getLocationPathString", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLocationPathString()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getLockboxNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLockboxNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getLogicalVolumeStorageGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLogicalVolumeStorageGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getManagedResourceFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetManagedResourceFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getManagementNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetManagementNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getMemory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMemory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getMemoryCapacity", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMemoryCapacity()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getMetricTrackingObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMetricTrackingObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getMetricTrackingObjectId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMetricTrackingObjectId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getModules", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetModules()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getMonitoringRobot", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMonitoringRobot()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getMonitoringServiceComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMonitoringServiceComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getMonitoringServiceEligibilityFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMonitoringServiceEligibilityFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getMonitoringUserNotification", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMonitoringUserNotification()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getMotherboard", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMotherboard()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getNetworkCards", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkCards()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getNetworkComponentFirewallProtectableIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkComponentFirewallProtectableIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getNetworkGatewayMember", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkGatewayMember()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getNetworkGatewayMemberFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkGatewayMemberFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getNetworkManagementIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkManagementIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getNetworkMonitorAttachedDownHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitorAttachedDownHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getNetworkMonitorAttachedDownVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitorAttachedDownVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getNetworkMonitorIncidents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitorIncidents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getNetworkMonitors", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitors()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getNetworkStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getNetworkStatusAttribute", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkStatusAttribute()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getNetworkVlans", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkVlans()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getNextBillingCycleBandwidthAllocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNextBillingCycleBandwidthAllocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getNotesHistory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNotesHistory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getNvRamCapacity", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNvRamCapacity()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getNvRamComponentModels", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNvRamComponentModels()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getOpenCancellationTicket", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOpenCancellationTicket()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getOperatingSystem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOperatingSystem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getOperatingSystemReferenceCode", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOperatingSystemReferenceCode()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getOutboundBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOutboundBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getOutboundPrivateBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOutboundPrivateBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getOutboundPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOutboundPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getOverBandwidthAllocationFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOverBandwidthAllocationFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getPMInfo", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPMInfo()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getParentBay", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetParentBay()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getParentHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetParentHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getPartitions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPartitions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getPointOfPresenceLocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPointOfPresenceLocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getPowerComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPowerComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getPowerSupply", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPowerSupply()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getPrimaryBackendIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryBackendIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getPrimaryBackendNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryBackendNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getPrimaryDriveSize", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryDriveSize()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getPrimaryIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getPrimaryNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getPrivateBackendNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateBackendNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getPrivateBandwidthData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateBandwidthData(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getPrivateBandwidthDataSummary", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateBandwidthDataSummary()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getPrivateBandwidthGraphImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateBandwidthGraphImage(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getPrivateIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getPrivateNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getPrivateNetworkOnlyFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateNetworkOnlyFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getPrivateVlan", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateVlan()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getPrivateVlanByIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateVlanByIpAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getProcessorCoreAmount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProcessorCoreAmount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getProcessorPhysicalCoreAmount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProcessorPhysicalCoreAmount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getProcessors", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProcessors()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getProjectedOverBandwidthAllocationFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProjectedOverBandwidthAllocationFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getProjectedPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProjectedPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getProvisionDate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProvisionDate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getPublicBandwidthData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicBandwidthData(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getPublicBandwidthDataSummary", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicBandwidthDataSummary()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getPublicBandwidthGraphImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicBandwidthGraphImage(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getPublicBandwidthTotal", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicBandwidthTotal(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getPublicNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getPublicVlan", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicVlan()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getPublicVlanByHostname", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicVlanByHostname(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getRack", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRack()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getRaidControllers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRaidControllers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getReadyNodeFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReadyNodeFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getRecentEvents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRecentEvents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getRecentRemoteManagementCommands", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRecentRemoteManagementCommands()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getRegionalInternetRegistry", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRegionalInternetRegistry()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getRemoteManagement", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRemoteManagement()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getRemoteManagementAccounts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRemoteManagementAccounts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getRemoteManagementComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRemoteManagementComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getRemoteManagementUsers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRemoteManagementUsers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getResourceConfigurations", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResourceConfigurations()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getResourceGroupMemberReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResourceGroupMemberReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getResourceGroupRoles", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResourceGroupRoles()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getResourceGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResourceGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getReverseDomainRecords", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReverseDomainRecords()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getRouters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRouters()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getSecurityScanRequests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSecurityScanRequests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getSensorData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSensorData()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getSensorDataWithGraphs", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSensorDataWithGraphs()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getServerDetails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerDetails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getServerFanSpeedGraphs", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerFanSpeedGraphs()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getServerPowerState", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerPowerState()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getServerRoom", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerRoom()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getServerTemperatureGraphs", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerTemperatureGraphs()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getServiceProvider", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceProvider()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getSoftwareComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSoftwareComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getSoftwareGuardExtensionEnabled", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSoftwareGuardExtensionEnabled()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getSparePoolBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSparePoolBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getSshKeys", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSshKeys()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getStatisticsRemoteManagement", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStatisticsRemoteManagement()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getStorageGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getStorageNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getTagReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTagReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getTopLevelLocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTopLevelLocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getTransactionHistory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTransactionHistory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getUefiBootFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUefiBootFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getUpgradeItemPrices", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpgradeItemPrices()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getUpgradeRequest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpgradeRequest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getUpgradeableActiveComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpgradeableActiveComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getUplinkHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUplinkHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getUplinkNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUplinkNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getUserData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUserData()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getUsers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUsers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getValidBlockDeviceTemplateGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetValidBlockDeviceTemplateGroups(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getVirtualChassis", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualChassis()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getVirtualChassisSiblings", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualChassisSiblings()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getVirtualHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualHost()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getVirtualLicenses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualLicenses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getVirtualRack", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualRack()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getVirtualRackId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualRackId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getVirtualRackName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualRackName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getVirtualizationPlatform", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualizationPlatform()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getWindowsUpdateAvailableUpdates", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetWindowsUpdateAvailableUpdates()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getWindowsUpdateInstalledUpdates", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetWindowsUpdateInstalledUpdates()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::getWindowsUpdateStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetWindowsUpdateStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::importVirtualHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.ImportVirtualHost()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::initiateIderaBareMetalRestore", func() {
            It("API Call Test", func() {

                _, err := sl_service.InitiateIderaBareMetalRestore()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::initiateR1SoftBareMetalRestore", func() {
            It("API Call Test", func() {

                _, err := sl_service.InitiateR1SoftBareMetalRestore()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::isBackendPingable", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsBackendPingable()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::isPingable", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsPingable()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::isWindowsServer", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsWindowsServer()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::massFirmwareReflash", func() {
            It("API Call Test", func() {

                _, err := sl_service.MassFirmwareReflash(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::massFirmwareUpdate", func() {
            It("API Call Test", func() {

                _, err := sl_service.MassFirmwareUpdate(nil,nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::massHyperThreadingUpdate", func() {
            It("API Call Test", func() {

                _, err := sl_service.MassHyperThreadingUpdate(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::massReloadOperatingSystem", func() {
            It("API Call Test", func() {

                _, err := sl_service.MassReloadOperatingSystem(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::massSparePool", func() {
            It("API Call Test", func() {

                _, err := sl_service.MassSparePool(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::ping", func() {
            It("API Call Test", func() {

                _, err := sl_service.Ping()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::populateServerRam", func() {
            It("API Call Test", func() {

				err := sl_service.PopulateServerRam(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::powerCycle", func() {
            It("API Call Test", func() {

                _, err := sl_service.PowerCycle()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::powerOff", func() {
            It("API Call Test", func() {

                _, err := sl_service.PowerOff()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::powerOn", func() {
            It("API Call Test", func() {

                _, err := sl_service.PowerOn()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::rebootDefault", func() {
            It("API Call Test", func() {

                _, err := sl_service.RebootDefault()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::rebootHard", func() {
            It("API Call Test", func() {

                _, err := sl_service.RebootHard()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::rebootSoft", func() {
            It("API Call Test", func() {

                _, err := sl_service.RebootSoft()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::refreshDeviceStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.RefreshDeviceStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::reloadCurrentOperatingSystemConfiguration", func() {
            It("API Call Test", func() {

                _, err := sl_service.ReloadCurrentOperatingSystemConfiguration(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::reloadOperatingSystem", func() {
            It("API Call Test", func() {

                _, err := sl_service.ReloadOperatingSystem(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::removeAccessToNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToNetworkStorage(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::removeAccessToNetworkStorageList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToNetworkStorageList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::removeTags", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveTags(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::runPassmarkCertificationBenchmark", func() {
            It("API Call Test", func() {

                _, err := sl_service.RunPassmarkCertificationBenchmark()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::setOperatingSystemPassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetOperatingSystemPassword(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::setPrivateNetworkInterfaceSpeed", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetPrivateNetworkInterfaceSpeed(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::setPublicNetworkInterfaceSpeed", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetPublicNetworkInterfaceSpeed(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::setTags", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetTags(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::setUserMetadata", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetUserMetadata(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::shutdownPrivatePort", func() {
            It("API Call Test", func() {

                _, err := sl_service.ShutdownPrivatePort()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::shutdownPublicPort", func() {
            It("API Call Test", func() {

                _, err := sl_service.ShutdownPublicPort()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::sparePool", func() {
            It("API Call Test", func() {

                _, err := sl_service.SparePool(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::testRaidAlertService", func() {
            It("API Call Test", func() {

                _, err := sl_service.TestRaidAlertService()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::toggleManagementInterface", func() {
            It("API Call Test", func() {

                _, err := sl_service.ToggleManagementInterface(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::updateIpmiPassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateIpmiPassword(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_SecurityModule750::validatePartitionsForOperatingSystem", func() {
            It("API Call Test", func() {

                _, err := sl_service.ValidatePartitionsForOperatingSystem(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Hardware_Server service", func() {
        var sl_service services.Hardware_Server
        BeforeEach(func() {
            sl_service = services.GetHardwareServerService(slsession)
        })


        Context("SoftLayer_Hardware_Server::activatePrivatePort", func() {
            It("API Call Test", func() {

                _, err := sl_service.ActivatePrivatePort()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::activatePublicPort", func() {
            It("API Call Test", func() {

                _, err := sl_service.ActivatePublicPort()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::allowAccessToNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToNetworkStorage(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::allowAccessToNetworkStorageList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToNetworkStorageList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::bootToRescueLayer", func() {
            It("API Call Test", func() {

                _, err := sl_service.BootToRescueLayer(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::captureImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.CaptureImage(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::createFirmwareReflashTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateFirmwareReflashTransaction(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::createFirmwareUpdateTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateFirmwareUpdateTransaction(nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::createHyperThreadingUpdateTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateHyperThreadingUpdateTransaction(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::createPostSoftwareInstallTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreatePostSoftwareInstallTransaction(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::deleteSoftwareComponentPasswords", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteSoftwareComponentPasswords(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::deleteTag", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteTag(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::editSoftwareComponentPasswords", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditSoftwareComponentPasswords(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::executeRemoteScript", func() {
            It("API Call Test", func() {

				err := sl_service.ExecuteRemoteScript(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::findByIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.FindByIpAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::generateOrderTemplate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GenerateOrderTemplate(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getActiveComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getActiveNetworkFirewallBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveNetworkFirewallBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getActiveNetworkMonitorIncident", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveNetworkMonitorIncident()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getActiveTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getActiveTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveTransaction()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getActiveTransactions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveTransactions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getAllPowerComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllPowerComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getAllowedHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedHost()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getAllowedNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getAllowedNetworkStorageReplicas", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedNetworkStorageReplicas()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getAntivirusSpywareSoftwareComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAntivirusSpywareSoftwareComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getAttachedNetworkStorages", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachedNetworkStorages(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getAttributes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttributes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getAvailableBillingTermChangePrices", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailableBillingTermChangePrices()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getAvailableMonitoring", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailableMonitoring()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getAvailableNetworkStorages", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailableNetworkStorages(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getAverageDailyBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAverageDailyBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getAverageDailyPrivateBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAverageDailyPrivateBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getAverageDailyPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAverageDailyPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getBackendBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendBandwidthUsage(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getBackendIncomingBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendIncomingBandwidth(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getBackendNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getBackendOutgoingBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendOutgoingBandwidth(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getBackendRouters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendRouters()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getBandwidthAllocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthAllocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getBandwidthAllotmentDetail", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthAllotmentDetail()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getBandwidthForDateRange", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthForDateRange(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getBandwidthImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthImage(nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getBenchmarkCertifications", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBenchmarkCertifications()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getBillingCycleBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingCycleBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getBillingCyclePrivateBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingCyclePrivateBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getBillingCyclePublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingCyclePublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getBillingItemFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItemFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getBiosPasswordNullFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBiosPasswordNullFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getBlockCancelBecauseDisconnectedFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBlockCancelBecauseDisconnectedFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getBootModeOptions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBootModeOptions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getBusinessContinuanceInsuranceFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBusinessContinuanceInsuranceFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getCaptureEnabledFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCaptureEnabledFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getChildrenHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetChildrenHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getComponentDetailsXML", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetComponentDetailsXML()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getContainsSolidStateDrivesFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetContainsSolidStateDrivesFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getContinuousDataProtectionSoftwareComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetContinuousDataProtectionSoftwareComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getControlPanel", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetControlPanel()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getCost", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCost()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getCreateObjectOptions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCreateObjectOptions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getCurrentBandwidthSummary", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentBandwidthSummary()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getCurrentBenchmarkCertificationResultFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentBenchmarkCertificationResultFile()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getCurrentBillableBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentBillableBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getCurrentBillingDetail", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentBillingDetail()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getCurrentBillingTotal", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentBillingTotal()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getCustomerInstalledOperatingSystemFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCustomerInstalledOperatingSystemFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getCustomerOwnedFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCustomerOwnedFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getDailyAverage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDailyAverage(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getDatacenter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDatacenter()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getDatacenterName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDatacenterName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getDaysInSparePool", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDaysInSparePool()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getDownlinkHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownlinkHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getDownlinkNetworkHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownlinkNetworkHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getDownlinkServers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownlinkServers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getDownlinkVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownlinkVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getDownstreamHardwareBindings", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownstreamHardwareBindings()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getDownstreamNetworkHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownstreamNetworkHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getDownstreamNetworkHardwareWithIncidents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownstreamNetworkHardwareWithIncidents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getDownstreamServers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownstreamServers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getDownstreamVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownstreamVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getDriveControllers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDriveControllers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getEvaultNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEvaultNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getFirewallProtectableSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirewallProtectableSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getFirewallServiceComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirewallServiceComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getFixedConfigurationPreset", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFixedConfigurationPreset()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getFrontendBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFrontendBandwidthUsage(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getFrontendIncomingBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFrontendIncomingBandwidth(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getFrontendNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFrontendNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getFrontendOutgoingBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFrontendOutgoingBandwidth(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getFrontendRouters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFrontendRouters()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getFutureBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFutureBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getGlobalIdentifier", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGlobalIdentifier()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getHardDrives", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardDrives()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getHardwareByIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareByIpAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getHardwareChassis", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareChassis()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getHardwareFunction", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareFunction()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getHardwareFunctionDescription", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareFunctionDescription()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getHardwareState", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareState()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getHardwareStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getHasSingleRootVirtualizationBillingItemFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasSingleRootVirtualizationBillingItemFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getHasTrustedPlatformModuleBillingItemFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasTrustedPlatformModuleBillingItemFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getHostIpsSoftwareComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHostIpsSoftwareComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getHourlyBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHourlyBandwidth(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getHourlyBillingFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHourlyBillingFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getInboundBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInboundBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getInboundPrivateBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInboundPrivateBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getInboundPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInboundPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getIsBillingTermChangeAvailableFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsBillingTermChangeAvailableFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getIsCloudReadyNodeCertified", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsCloudReadyNodeCertified()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getIsIpmiDisabled", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsIpmiDisabled()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getIsQeInternalServer", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsQeInternalServer()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getIsVirtualPrivateCloudNode", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsVirtualPrivateCloudNode()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getItemPricesFromSoftwareDescriptions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetItemPricesFromSoftwareDescriptions(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getLastOperatingSystemReload", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLastOperatingSystemReload()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getLastTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLastTransaction()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getLatestNetworkMonitorIncident", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLatestNetworkMonitorIncident()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getLocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getLocationPathString", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLocationPathString()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getLockboxNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLockboxNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getLogicalVolumeStorageGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLogicalVolumeStorageGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getManagedResourceFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetManagedResourceFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getManagementNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetManagementNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getMemory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMemory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getMemoryCapacity", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMemoryCapacity()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getMetricTrackingObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMetricTrackingObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getMetricTrackingObjectId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMetricTrackingObjectId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getModules", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetModules()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getMonitoringRobot", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMonitoringRobot()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getMonitoringServiceComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMonitoringServiceComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getMonitoringServiceEligibilityFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMonitoringServiceEligibilityFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getMonitoringUserNotification", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMonitoringUserNotification()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getMotherboard", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMotherboard()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getNetworkCards", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkCards()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getNetworkComponentFirewallProtectableIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkComponentFirewallProtectableIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getNetworkGatewayMember", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkGatewayMember()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getNetworkGatewayMemberFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkGatewayMemberFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getNetworkManagementIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkManagementIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getNetworkMonitorAttachedDownHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitorAttachedDownHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getNetworkMonitorAttachedDownVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitorAttachedDownVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getNetworkMonitorIncidents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitorIncidents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getNetworkMonitors", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitors()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getNetworkStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getNetworkStatusAttribute", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkStatusAttribute()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getNetworkVlans", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkVlans()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getNextBillingCycleBandwidthAllocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNextBillingCycleBandwidthAllocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getNotesHistory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNotesHistory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getNvRamCapacity", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNvRamCapacity()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getNvRamComponentModels", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNvRamComponentModels()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getOpenCancellationTicket", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOpenCancellationTicket()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getOperatingSystem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOperatingSystem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getOperatingSystemReferenceCode", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOperatingSystemReferenceCode()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getOutboundBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOutboundBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getOutboundPrivateBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOutboundPrivateBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getOutboundPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOutboundPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getOverBandwidthAllocationFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOverBandwidthAllocationFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getPMInfo", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPMInfo()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getParentBay", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetParentBay()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getParentHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetParentHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getPartitions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPartitions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getPointOfPresenceLocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPointOfPresenceLocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getPowerComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPowerComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getPowerSupply", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPowerSupply()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getPrimaryBackendIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryBackendIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getPrimaryBackendNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryBackendNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getPrimaryDriveSize", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryDriveSize()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getPrimaryIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getPrimaryNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getPrivateBackendNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateBackendNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getPrivateBandwidthData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateBandwidthData(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getPrivateBandwidthDataSummary", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateBandwidthDataSummary()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getPrivateBandwidthGraphImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateBandwidthGraphImage(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getPrivateIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getPrivateNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getPrivateNetworkOnlyFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateNetworkOnlyFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getPrivateVlan", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateVlan()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getPrivateVlanByIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateVlanByIpAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getProcessorCoreAmount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProcessorCoreAmount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getProcessorPhysicalCoreAmount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProcessorPhysicalCoreAmount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getProcessors", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProcessors()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getProjectedOverBandwidthAllocationFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProjectedOverBandwidthAllocationFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getProjectedPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProjectedPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getProvisionDate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProvisionDate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getPublicBandwidthData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicBandwidthData(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getPublicBandwidthDataSummary", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicBandwidthDataSummary()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getPublicBandwidthGraphImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicBandwidthGraphImage(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getPublicBandwidthTotal", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicBandwidthTotal(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getPublicNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getPublicVlan", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicVlan()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getPublicVlanByHostname", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicVlanByHostname(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getRack", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRack()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getRaidControllers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRaidControllers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getReadyNodeFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReadyNodeFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getRecentEvents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRecentEvents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getRecentRemoteManagementCommands", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRecentRemoteManagementCommands()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getRegionalInternetRegistry", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRegionalInternetRegistry()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getRemoteManagement", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRemoteManagement()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getRemoteManagementAccounts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRemoteManagementAccounts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getRemoteManagementComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRemoteManagementComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getRemoteManagementUsers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRemoteManagementUsers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getResourceConfigurations", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResourceConfigurations()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getResourceGroupMemberReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResourceGroupMemberReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getResourceGroupRoles", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResourceGroupRoles()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getResourceGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResourceGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getReverseDomainRecords", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReverseDomainRecords()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getRouters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRouters()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getSecurityScanRequests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSecurityScanRequests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getSensorData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSensorData()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getSensorDataWithGraphs", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSensorDataWithGraphs()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getServerDetails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerDetails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getServerFanSpeedGraphs", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerFanSpeedGraphs()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getServerPowerState", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerPowerState()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getServerRoom", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerRoom()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getServerTemperatureGraphs", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerTemperatureGraphs()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getServiceProvider", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceProvider()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getSoftwareComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSoftwareComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getSoftwareGuardExtensionEnabled", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSoftwareGuardExtensionEnabled()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getSparePoolBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSparePoolBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getSshKeys", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSshKeys()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getStatisticsRemoteManagement", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStatisticsRemoteManagement()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getStorageGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getStorageNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getTagReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTagReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getTopLevelLocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTopLevelLocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getTransactionHistory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTransactionHistory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getUefiBootFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUefiBootFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getUpgradeItemPrices", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpgradeItemPrices()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getUpgradeRequest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpgradeRequest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getUpgradeableActiveComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpgradeableActiveComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getUplinkHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUplinkHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getUplinkNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUplinkNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getUserData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUserData()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getUsers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUsers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getValidBlockDeviceTemplateGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetValidBlockDeviceTemplateGroups(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getVirtualChassis", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualChassis()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getVirtualChassisSiblings", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualChassisSiblings()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getVirtualHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualHost()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getVirtualLicenses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualLicenses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getVirtualRack", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualRack()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getVirtualRackId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualRackId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getVirtualRackName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualRackName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getVirtualizationPlatform", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualizationPlatform()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getWindowsUpdateAvailableUpdates", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetWindowsUpdateAvailableUpdates()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getWindowsUpdateInstalledUpdates", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetWindowsUpdateInstalledUpdates()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::getWindowsUpdateStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetWindowsUpdateStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::importVirtualHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.ImportVirtualHost()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::initiateIderaBareMetalRestore", func() {
            It("API Call Test", func() {

                _, err := sl_service.InitiateIderaBareMetalRestore()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::initiateR1SoftBareMetalRestore", func() {
            It("API Call Test", func() {

                _, err := sl_service.InitiateR1SoftBareMetalRestore()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::isBackendPingable", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsBackendPingable()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::isPingable", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsPingable()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::isWindowsServer", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsWindowsServer()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::massFirmwareReflash", func() {
            It("API Call Test", func() {

                _, err := sl_service.MassFirmwareReflash(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::massFirmwareUpdate", func() {
            It("API Call Test", func() {

                _, err := sl_service.MassFirmwareUpdate(nil,nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::massHyperThreadingUpdate", func() {
            It("API Call Test", func() {

                _, err := sl_service.MassHyperThreadingUpdate(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::massReloadOperatingSystem", func() {
            It("API Call Test", func() {

                _, err := sl_service.MassReloadOperatingSystem(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::massSparePool", func() {
            It("API Call Test", func() {

                _, err := sl_service.MassSparePool(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::ping", func() {
            It("API Call Test", func() {

                _, err := sl_service.Ping()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::populateServerRam", func() {
            It("API Call Test", func() {

				err := sl_service.PopulateServerRam(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::powerCycle", func() {
            It("API Call Test", func() {

                _, err := sl_service.PowerCycle()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::powerOff", func() {
            It("API Call Test", func() {

                _, err := sl_service.PowerOff()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::powerOn", func() {
            It("API Call Test", func() {

                _, err := sl_service.PowerOn()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::rebootDefault", func() {
            It("API Call Test", func() {

                _, err := sl_service.RebootDefault()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::rebootHard", func() {
            It("API Call Test", func() {

                _, err := sl_service.RebootHard()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::rebootSoft", func() {
            It("API Call Test", func() {

                _, err := sl_service.RebootSoft()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::refreshDeviceStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.RefreshDeviceStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::reloadCurrentOperatingSystemConfiguration", func() {
            It("API Call Test", func() {

                _, err := sl_service.ReloadCurrentOperatingSystemConfiguration(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::reloadOperatingSystem", func() {
            It("API Call Test", func() {

                _, err := sl_service.ReloadOperatingSystem(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::removeAccessToNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToNetworkStorage(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::removeAccessToNetworkStorageList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToNetworkStorageList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::removeTags", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveTags(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::runPassmarkCertificationBenchmark", func() {
            It("API Call Test", func() {

                _, err := sl_service.RunPassmarkCertificationBenchmark()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::setOperatingSystemPassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetOperatingSystemPassword(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::setPrivateNetworkInterfaceSpeed", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetPrivateNetworkInterfaceSpeed(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::setPublicNetworkInterfaceSpeed", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetPublicNetworkInterfaceSpeed(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::setTags", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetTags(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::setUserMetadata", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetUserMetadata(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::shutdownPrivatePort", func() {
            It("API Call Test", func() {

                _, err := sl_service.ShutdownPrivatePort()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::shutdownPublicPort", func() {
            It("API Call Test", func() {

                _, err := sl_service.ShutdownPublicPort()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::sparePool", func() {
            It("API Call Test", func() {

                _, err := sl_service.SparePool(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::testRaidAlertService", func() {
            It("API Call Test", func() {

                _, err := sl_service.TestRaidAlertService()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::toggleManagementInterface", func() {
            It("API Call Test", func() {

                _, err := sl_service.ToggleManagementInterface(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::updateIpmiPassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateIpmiPassword(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Hardware_Server::validatePartitionsForOperatingSystem", func() {
            It("API Call Test", func() {

                _, err := sl_service.ValidatePartitionsForOperatingSystem(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

})
