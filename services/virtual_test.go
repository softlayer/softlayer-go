
package services_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/softlayer/softlayer-go/session/sessionfakes"
    "github.com/softlayer/softlayer-go/services"
)   

var _ = Describe("Virtual Tests", func() {
    var slsession *sessionfakes.FakeSLSession
    BeforeEach(func() {
        slsession = &sessionfakes.FakeSLSession{}
    })

    Context("Testing SoftLayer_Virtual_DedicatedHost service", func() {
        var sl_service services.Virtual_DedicatedHost
        BeforeEach(func() {
            sl_service = services.GetVirtualDedicatedHostService(slsession)
        })


        Context("SoftLayer_Virtual_DedicatedHost::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_DedicatedHost::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_DedicatedHost::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_DedicatedHost::getAllocationStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllocationStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_DedicatedHost::getAvailableRouters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailableRouters(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_DedicatedHost::getBackendRouter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendRouter()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_DedicatedHost::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_DedicatedHost::getDatacenter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDatacenter()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_DedicatedHost::getGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_DedicatedHost::getInternalTagReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInternalTagReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_DedicatedHost::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_DedicatedHost::getPciDeviceAllocationStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPciDeviceAllocationStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_DedicatedHost::getPciDevices", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPciDevices()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_DedicatedHost::getTagReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTagReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_DedicatedHost::setTags", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetTags(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Virtual_Disk_Image service", func() {
        var sl_service services.Virtual_Disk_Image
        BeforeEach(func() {
            sl_service = services.GetVirtualDiskImageService(slsession)
        })


        Context("SoftLayer_Virtual_Disk_Image::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Disk_Image::getAvailableBootModes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailableBootModes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Disk_Image::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Disk_Image::getBlockDevices", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBlockDevices()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Disk_Image::getBootableVolumeFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBootableVolumeFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Disk_Image::getCloudInitFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCloudInitFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Disk_Image::getCoalescedDiskImages", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCoalescedDiskImages()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Disk_Image::getCopyOnWriteFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCopyOnWriteFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Disk_Image::getDiskFileExtension", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDiskFileExtension()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Disk_Image::getDiskImageStorageGroup", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDiskImageStorageGroup()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Disk_Image::getImportedDiskType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetImportedDiskType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Disk_Image::getIsEncrypted", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsEncrypted()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Disk_Image::getLocalDiskFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLocalDiskFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Disk_Image::getMetadataFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMetadataFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Disk_Image::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Disk_Image::getPublicIsoImages", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicIsoImages()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Disk_Image::getSoftwareReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSoftwareReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Disk_Image::getSourceDiskImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSourceDiskImage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Disk_Image::getStorageGroupDetails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageGroupDetails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Disk_Image::getStorageGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Disk_Image::getStorageRepository", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageRepository()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Disk_Image::getStorageRepositoryType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageRepositoryType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Disk_Image::getSupportedHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSupportedHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Disk_Image::getTemplateBlockDevice", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTemplateBlockDevice()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Disk_Image::getType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Virtual_Guest service", func() {
        var sl_service services.Virtual_Guest
        BeforeEach(func() {
            sl_service = services.GetVirtualGuestService(slsession)
        })


        Context("SoftLayer_Virtual_Guest::activatePrivatePort", func() {
            It("API Call Test", func() {

                _, err := sl_service.ActivatePrivatePort()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::activatePublicPort", func() {
            It("API Call Test", func() {

                _, err := sl_service.ActivatePublicPort()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::allowAccessToNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToNetworkStorage(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::allowAccessToNetworkStorageList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToNetworkStorageList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::attachDiskImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.AttachDiskImage(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::cancelIsolationForDestructiveAction", func() {
            It("API Call Test", func() {

				err := sl_service.CancelIsolationForDestructiveAction()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::captureImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.CaptureImage(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::checkHostDiskAvailability", func() {
            It("API Call Test", func() {

                _, err := sl_service.CheckHostDiskAvailability(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::configureMetadataDisk", func() {
            It("API Call Test", func() {

                _, err := sl_service.ConfigureMetadataDisk()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::createArchiveTemplate", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateArchiveTemplate(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::createArchiveTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateArchiveTransaction(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::createObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::createPostSoftwareInstallTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreatePostSoftwareInstallTransaction(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::deleteTag", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteTag(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::deleteTransientWebhook", func() {
            It("API Call Test", func() {

				err := sl_service.DeleteTransientWebhook()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::detachDiskImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.DetachDiskImage(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::executeIderaBareMetalRestore", func() {
            It("API Call Test", func() {

                _, err := sl_service.ExecuteIderaBareMetalRestore()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::executeR1SoftBareMetalRestore", func() {
            It("API Call Test", func() {

                _, err := sl_service.ExecuteR1SoftBareMetalRestore()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::executeRemoteScript", func() {
            It("API Call Test", func() {

				err := sl_service.ExecuteRemoteScript(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::executeRescueLayer", func() {
            It("API Call Test", func() {

                _, err := sl_service.ExecuteRescueLayer()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::findByHostname", func() {
            It("API Call Test", func() {

                _, err := sl_service.FindByHostname(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::findByIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.FindByIpAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::generateOrderTemplate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GenerateOrderTemplate(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getAccountOwnedPoolFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccountOwnedPoolFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getActiveNetworkMonitorIncident", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveNetworkMonitorIncident()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getActiveTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getActiveTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveTransaction()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getActiveTransactions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveTransactions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getAdditionalRequiredPricesForOsReload", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAdditionalRequiredPricesForOsReload(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getAllowedHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedHost()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getAllowedNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getAllowedNetworkStorageReplicas", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedNetworkStorageReplicas()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getAntivirusSpywareSoftwareComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAntivirusSpywareSoftwareComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getApplicationDeliveryController", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetApplicationDeliveryController()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getAttachedNetworkStorages", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachedNetworkStorages(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getAttributes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttributes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getAvailableBlockDevicePositions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailableBlockDevicePositions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getAvailableMonitoring", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailableMonitoring()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getAvailableNetworkStorages", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailableNetworkStorages(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getAverageDailyPrivateBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAverageDailyPrivateBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getAverageDailyPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAverageDailyPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getBackendNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getBackendRouters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendRouters()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getBandwidthAllocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthAllocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getBandwidthAllotmentDetail", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthAllotmentDetail()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getBandwidthDataByDate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthDataByDate(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getBandwidthForDateRange", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthForDateRange(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getBandwidthImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthImage(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getBandwidthImageByDate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthImageByDate(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getBandwidthTotal", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthTotal(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getBillingCycleBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingCycleBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getBillingCyclePrivateBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingCyclePrivateBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getBillingCyclePublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingCyclePublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getBlockCancelBecauseDisconnectedFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBlockCancelBecauseDisconnectedFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getBlockDeviceTemplateGroup", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBlockDeviceTemplateGroup()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getBlockDevices", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBlockDevices()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getBootMode", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBootMode()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getBootOrder", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBootOrder()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getBrowserConsoleAccessLogs", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBrowserConsoleAccessLogs()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getConsoleAccessLog", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetConsoleAccessLog()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getConsoleData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetConsoleData()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getConsoleIpAddressFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetConsoleIpAddressFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getConsoleIpAddressRecord", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetConsoleIpAddressRecord()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getContinuousDataProtectionSoftwareComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetContinuousDataProtectionSoftwareComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getControlPanel", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetControlPanel()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getCoreRestrictedOperatingSystemPrice", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCoreRestrictedOperatingSystemPrice()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getCpuMetricDataByDate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCpuMetricDataByDate(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getCpuMetricImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCpuMetricImage(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getCpuMetricImageByDate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCpuMetricImageByDate(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getCreateObjectOptions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCreateObjectOptions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getCurrentBandwidthSummary", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentBandwidthSummary()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getCurrentBillingDetail", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentBillingDetail()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getCurrentBillingTotal", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentBillingTotal()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getDatacenter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDatacenter()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getDedicatedHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDedicatedHost()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getDeviceStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDeviceStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getDriveRetentionItemPrice", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDriveRetentionItemPrice()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getEvaultNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEvaultNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getFirewallProtectableSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirewallProtectableSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getFirewallServiceComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirewallServiceComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getFirstAvailableBlockDevicePosition", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirstAvailableBlockDevicePosition()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getFrontendNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFrontendNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getFrontendRouters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFrontendRouters()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getGlobalIdentifier", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGlobalIdentifier()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getGpuCount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGpuCount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getGpuType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGpuType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getGuestBootParameter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGuestBootParameter()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getHardwareFunctionDescription", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareFunctionDescription()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHost()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getHostIpsSoftwareComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHostIpsSoftwareComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getHourlyBillingFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHourlyBillingFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getInboundPrivateBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInboundPrivateBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getInboundPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInboundPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getInternalTagReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInternalTagReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getIsoBootImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsoBootImage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getItemPricesFromSoftwareDescriptions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetItemPricesFromSoftwareDescriptions(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getLastKnownPowerState", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLastKnownPowerState()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getLastOperatingSystemReload", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLastOperatingSystemReload()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getLastTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLastTransaction()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getLatestNetworkMonitorIncident", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLatestNetworkMonitorIncident()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getLocalDiskFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLocalDiskFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getLocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getManagedResourceFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetManagedResourceFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getMemoryMetricDataByDate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMemoryMetricDataByDate(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getMemoryMetricImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMemoryMetricImage(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getMemoryMetricImageByDate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMemoryMetricImageByDate(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getMetricTrackingObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMetricTrackingObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getMetricTrackingObjectId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMetricTrackingObjectId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getMonitoringRobot", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMonitoringRobot()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getMonitoringServiceComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMonitoringServiceComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getMonitoringServiceEligibilityFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMonitoringServiceEligibilityFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getMonitoringUserNotification", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMonitoringUserNotification()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getNetworkComponentFirewallProtectableIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkComponentFirewallProtectableIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getNetworkMonitorIncidents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitorIncidents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getNetworkMonitors", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMonitors()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getNetworkVlans", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkVlans()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getOpenCancellationTicket", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOpenCancellationTicket()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getOperatingSystem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOperatingSystem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getOperatingSystemReferenceCode", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOperatingSystemReferenceCode()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getOrderTemplate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOrderTemplate(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getOrderedPackageId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOrderedPackageId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getOutboundPrivateBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOutboundPrivateBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getOutboundPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOutboundPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getOverBandwidthAllocationFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOverBandwidthAllocationFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getPendingMaintenanceActions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPendingMaintenanceActions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getPendingMigrationFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPendingMigrationFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getPlacementGroup", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPlacementGroup()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getPowerState", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPowerState()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getPrimaryBackendIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryBackendIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getPrimaryBackendNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryBackendNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getPrimaryIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getPrimaryNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getPrivateNetworkOnlyFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateNetworkOnlyFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getProjectedOverBandwidthAllocationFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProjectedOverBandwidthAllocationFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getProjectedPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProjectedPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getProvisionDate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProvisionDate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getRecentEvents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRecentEvents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getRecentMetricData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRecentMetricData(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getRegionalGroup", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRegionalGroup()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getRegionalInternetRegistry", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRegionalInternetRegistry()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getReservedCapacityGroup", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReservedCapacityGroup()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getReservedCapacityGroupFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReservedCapacityGroupFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getReservedCapacityGroupInstance", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReservedCapacityGroupInstance()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getReverseDomainRecords", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReverseDomainRecords()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getSecurityScanRequests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSecurityScanRequests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getServerRoom", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServerRoom()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getSoftwareComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSoftwareComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getSshKeys", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSshKeys()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getTagReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTagReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getTransientGuestFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTransientGuestFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getTransientWebhookURI", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTransientWebhookURI()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getUpgradeItemPrices", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpgradeItemPrices(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getUpgradeRequest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpgradeRequest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getUserData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUserData()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getUsers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUsers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getValidBlockDeviceTemplateGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetValidBlockDeviceTemplateGroups(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getVirtualRack", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualRack()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getVirtualRackId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualRackId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::getVirtualRackName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualRackName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::isBackendPingable", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsBackendPingable()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::isCloudInit", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsCloudInit()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::isPingable", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsPingable()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::isolateInstanceForDestructiveAction", func() {
            It("API Call Test", func() {

				err := sl_service.IsolateInstanceForDestructiveAction()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::migrate", func() {
            It("API Call Test", func() {

                _, err := sl_service.Migrate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::migrateDedicatedHost", func() {
            It("API Call Test", func() {

				err := sl_service.MigrateDedicatedHost(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::mountIsoImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.MountIsoImage(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::pause", func() {
            It("API Call Test", func() {

                _, err := sl_service.Pause()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::powerCycle", func() {
            It("API Call Test", func() {

                _, err := sl_service.PowerCycle()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::powerOff", func() {
            It("API Call Test", func() {

                _, err := sl_service.PowerOff()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::powerOffSoft", func() {
            It("API Call Test", func() {

                _, err := sl_service.PowerOffSoft()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::powerOn", func() {
            It("API Call Test", func() {

                _, err := sl_service.PowerOn()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::rebootDefault", func() {
            It("API Call Test", func() {

                _, err := sl_service.RebootDefault()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::rebootHard", func() {
            It("API Call Test", func() {

                _, err := sl_service.RebootHard()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::rebootSoft", func() {
            It("API Call Test", func() {

                _, err := sl_service.RebootSoft()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::reloadCurrentOperatingSystemConfiguration", func() {
            It("API Call Test", func() {

                _, err := sl_service.ReloadCurrentOperatingSystemConfiguration()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::reloadOperatingSystem", func() {
            It("API Call Test", func() {

                _, err := sl_service.ReloadOperatingSystem(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::removeAccessToNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToNetworkStorage(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::removeAccessToNetworkStorageList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToNetworkStorageList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::removeTags", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveTags(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::resume", func() {
            It("API Call Test", func() {

                _, err := sl_service.Resume()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::sendTestReclaimScheduledAlert", func() {
            It("API Call Test", func() {

				err := sl_service.SendTestReclaimScheduledAlert()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::setPrivateNetworkInterfaceSpeed", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetPrivateNetworkInterfaceSpeed(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::setPublicNetworkInterfaceSpeed", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetPublicNetworkInterfaceSpeed(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::setTags", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetTags(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::setTransientWebhook", func() {
            It("API Call Test", func() {

				err := sl_service.SetTransientWebhook(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::setUserMetadata", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetUserMetadata(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::shutdownPrivatePort", func() {
            It("API Call Test", func() {

                _, err := sl_service.ShutdownPrivatePort()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::shutdownPublicPort", func() {
            It("API Call Test", func() {

                _, err := sl_service.ShutdownPublicPort()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::unmountIsoImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.UnmountIsoImage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::validateImageTemplate", func() {
            It("API Call Test", func() {

                _, err := sl_service.ValidateImageTemplate(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest::verifyReloadOperatingSystem", func() {
            It("API Call Test", func() {

                _, err := sl_service.VerifyReloadOperatingSystem(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Virtual_Guest_Block_Device_Template_Group service", func() {
        var sl_service services.Virtual_Guest_Block_Device_Template_Group
        BeforeEach(func() {
            sl_service = services.GetVirtualGuestBlockDeviceTemplateGroupService(slsession)
        })


        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::addByolAttribute", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddByolAttribute()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::addCloudInitAttribute", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddCloudInitAttribute()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::addLocations", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddLocations(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::addSupportedBootMode", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddSupportedBootMode(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::copyToExternalSource", func() {
            It("API Call Test", func() {

                _, err := sl_service.CopyToExternalSource(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::copyToIcos", func() {
            It("API Call Test", func() {

                _, err := sl_service.CopyToIcos(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::createFromExternalSource", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateFromExternalSource(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::createFromIcos", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateFromIcos(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::createPublicArchiveTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreatePublicArchiveTransaction(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::deleteByolAttribute", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteByolAttribute()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::deleteCloudInitAttribute", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteCloudInitAttribute()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::denySharingAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.DenySharingAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::findGcImagesByCurrentUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.FindGcImagesByCurrentUser(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getAccountContacts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccountContacts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getAccountReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccountReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getAllAvailableCompatiblePlatformNames", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllAvailableCompatiblePlatformNames()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getBlockDevices", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBlockDevices()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getBlockDevicesDiskSpaceTotal", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBlockDevicesDiskSpaceTotal()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getBootMode", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBootMode()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getByolFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetByolFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getChildren", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetChildren()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getCurrentCompatiblePlatformNames", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentCompatiblePlatformNames()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getDatacenter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDatacenter()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getDatacenters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDatacenters()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getDefaultBootMode", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDefaultBootMode()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getEncryptionAttributes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEncryptionAttributes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getFirstChild", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirstChild()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getFlexImageFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFlexImageFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getGlobalIdentifier", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGlobalIdentifier()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getImageType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetImageType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getImageTypeKeyName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetImageTypeKeyName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getNextGenFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNextGenFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getParent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetParent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getPublicCustomerOwnedImages", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicCustomerOwnedImages()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getPublicImages", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicImages()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getRegion", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRegion()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getRegions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRegions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getRiasAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRiasAccount(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getSshKeys", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSshKeys()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getStorageLocations", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageLocations()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getStorageRepository", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageRepository()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getSupportedBootModes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSupportedBootModes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getTagReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTagReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getTemplateDataCenterName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTemplateDataCenterName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTransaction()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::getVhdImportSoftwareDescriptions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVhdImportSoftwareDescriptions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::isByol", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsByol()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::isByolCapableOperatingSystem", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsByolCapableOperatingSystem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::isByolOnlyOperatingSystem", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsByolOnlyOperatingSystem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::isCloudInit", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsCloudInit()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::isCloudInitOnlyOperatingSystem", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsCloudInitOnlyOperatingSystem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::isEncrypted", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsEncrypted()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::permitSharingAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.PermitSharingAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::removeCompatiblePlatforms", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveCompatiblePlatforms(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::removeLocations", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveLocations(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::removeSupportedBootMode", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveSupportedBootMode(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::setAvailableLocations", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetAvailableLocations(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::setBootMode", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetBootMode(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::setCompatiblePlatforms", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetCompatiblePlatforms(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Block_Device_Template_Group::setTags", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetTags(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Virtual_Guest_Boot_Parameter service", func() {
        var sl_service services.Virtual_Guest_Boot_Parameter
        BeforeEach(func() {
            sl_service = services.GetVirtualGuestBootParameterService(slsession)
        })


        Context("SoftLayer_Virtual_Guest_Boot_Parameter::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Boot_Parameter::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Boot_Parameter::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Boot_Parameter::getGuest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGuest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Boot_Parameter::getGuestBootParameterType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGuestBootParameterType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Boot_Parameter::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Virtual_Guest_Boot_Parameter_Type service", func() {
        var sl_service services.Virtual_Guest_Boot_Parameter_Type
        BeforeEach(func() {
            sl_service = services.GetVirtualGuestBootParameterTypeService(slsession)
        })


        Context("SoftLayer_Virtual_Guest_Boot_Parameter_Type::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Boot_Parameter_Type::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Virtual_Guest_Network_Component service", func() {
        var sl_service services.Virtual_Guest_Network_Component
        BeforeEach(func() {
            sl_service = services.GetVirtualGuestNetworkComponentService(slsession)
        })


        Context("SoftLayer_Virtual_Guest_Network_Component::disable", func() {
            It("API Call Test", func() {

                _, err := sl_service.Disable()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Network_Component::enable", func() {
            It("API Call Test", func() {

                _, err := sl_service.Enable()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Network_Component::getGuest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGuest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Network_Component::getHighAvailabilityFirewallFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHighAvailabilityFirewallFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Network_Component::getIcpBinding", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIcpBinding()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Network_Component::getIpAddressBindings", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIpAddressBindings()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Network_Component::getNetworkComponentFirewall", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkComponentFirewall()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Network_Component::getNetworkVlan", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkVlan()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Network_Component::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Network_Component::getPrimaryIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Network_Component::getPrimaryIpAddressRecord", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryIpAddressRecord()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Network_Component::getPrimarySubnet", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimarySubnet()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Network_Component::getPrimaryVersion6IpAddressRecord", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryVersion6IpAddressRecord()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Network_Component::getRouter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRouter()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Network_Component::getSecurityGroupBindings", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSecurityGroupBindings()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Network_Component::getSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Network_Component::isPingable", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsPingable()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Guest_Network_Component::securityGroupsReady", func() {
            It("API Call Test", func() {

                _, err := sl_service.SecurityGroupsReady()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Virtual_Host service", func() {
        var sl_service services.Virtual_Host
        BeforeEach(func() {
            sl_service = services.GetVirtualHostService(slsession)
        })


        Context("SoftLayer_Virtual_Host::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Host::getHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Host::getMetricTrackingObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMetricTrackingObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Host::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Host::getPciDevices", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPciDevices()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Virtual_PlacementGroup service", func() {
        var sl_service services.Virtual_PlacementGroup
        BeforeEach(func() {
            sl_service = services.GetVirtualPlacementGroupService(slsession)
        })


        Context("SoftLayer_Virtual_PlacementGroup::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_PlacementGroup::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_PlacementGroup::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_PlacementGroup::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_PlacementGroup::getAvailableRouters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailableRouters(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_PlacementGroup::getBackendRouter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendRouter()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_PlacementGroup::getGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_PlacementGroup::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_PlacementGroup::getRule", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRule()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Virtual_PlacementGroup_Rule service", func() {
        var sl_service services.Virtual_PlacementGroup_Rule
        BeforeEach(func() {
            sl_service = services.GetVirtualPlacementGroupRuleService(slsession)
        })


        Context("SoftLayer_Virtual_PlacementGroup_Rule::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_PlacementGroup_Rule::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Virtual_ReservedCapacityGroup service", func() {
        var sl_service services.Virtual_ReservedCapacityGroup
        BeforeEach(func() {
            sl_service = services.GetVirtualReservedCapacityGroupService(slsession)
        })


        Context("SoftLayer_Virtual_ReservedCapacityGroup::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_ReservedCapacityGroup::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_ReservedCapacityGroup::getAvailableInstances", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailableInstances()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_ReservedCapacityGroup::getBackendRouter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackendRouter()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_ReservedCapacityGroup::getInstances", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInstances()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_ReservedCapacityGroup::getInstancesCount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInstancesCount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_ReservedCapacityGroup::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_ReservedCapacityGroup::getOccupiedInstances", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOccupiedInstances()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Virtual_ReservedCapacityGroup_Instance service", func() {
        var sl_service services.Virtual_ReservedCapacityGroup_Instance
        BeforeEach(func() {
            sl_service = services.GetVirtualReservedCapacityGroupInstanceService(slsession)
        })


        Context("SoftLayer_Virtual_ReservedCapacityGroup_Instance::getAvailableFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailableFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_ReservedCapacityGroup_Instance::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_ReservedCapacityGroup_Instance::getGuest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGuest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_ReservedCapacityGroup_Instance::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_ReservedCapacityGroup_Instance::getReservedCapacityGroup", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReservedCapacityGroup()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Virtual_Storage_Repository service", func() {
        var sl_service services.Virtual_Storage_Repository
        BeforeEach(func() {
            sl_service = services.GetVirtualStorageRepositoryService(slsession)
        })


        Context("SoftLayer_Virtual_Storage_Repository::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Storage_Repository::getArchiveDiskUsageRatePerGb", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetArchiveDiskUsageRatePerGb()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Storage_Repository::getAverageDiskUsageMetricDataFromInfluxByDate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAverageDiskUsageMetricDataFromInfluxByDate(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Storage_Repository::getAverageUsageMetricDataByDate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAverageUsageMetricDataByDate(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Storage_Repository::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Storage_Repository::getDatacenter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDatacenter()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Storage_Repository::getDiskImages", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDiskImages()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Storage_Repository::getGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Storage_Repository::getMetricTrackingObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMetricTrackingObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Storage_Repository::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Storage_Repository::getPublicImageBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicImageBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Storage_Repository::getPublicImageDiskUsageRatePerGb", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicImageDiskUsageRatePerGb()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Storage_Repository::getStorageLocations", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageLocations()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Storage_Repository::getType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Storage_Repository::getUsageMetricDataByDate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUsageMetricDataByDate(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Virtual_Storage_Repository::getUsageMetricImageByDate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUsageMetricImageByDate(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

})
