
package services_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/softlayer/softlayer-go/session/sessionfakes"
    "github.com/softlayer/softlayer-go/services"
)   

var _ = Describe("Network Tests", func() {
    var slsession *sessionfakes.FakeSLSession
    BeforeEach(func() {
        slsession = &sessionfakes.FakeSLSession{}
    })

    Context("Testing SoftLayer_Network service", func() {
        var sl_service services.Network
        BeforeEach(func() {
            sl_service = services.GetNetworkService(slsession)
        })


        Context("SoftLayer_Network::connectPrivateEndpointService", func() {
            It("API Call Test", func() {

                _, err := sl_service.ConnectPrivateEndpointService()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network::disconnectPrivateEndpointService", func() {
            It("API Call Test", func() {

                _, err := sl_service.DisconnectPrivateEndpointService()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network::isConnectedToPrivateEndpointService", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsConnectedToPrivateEndpointService()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Application_Delivery_Controller service", func() {
        var sl_service services.Network_Application_Delivery_Controller
        BeforeEach(func() {
            sl_service = services.GetNetworkApplicationDeliveryControllerService(slsession)
        })


        Context("SoftLayer_Network_Application_Delivery_Controller::createLiveLoadBalancer", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateLiveLoadBalancer(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::deleteLiveLoadBalancer", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteLiveLoadBalancer(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::deleteLiveLoadBalancerService", func() {
            It("API Call Test", func() {

				err := sl_service.DeleteLiveLoadBalancerService(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::getAverageDailyPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAverageDailyPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::getBandwidthDataByDate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthDataByDate(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::getBandwidthImageByDate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthImageByDate(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::getConfigurationHistory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetConfigurationHistory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::getDatacenter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDatacenter()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::getDescription", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDescription()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::getInboundPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInboundPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::getLicenseExpirationDate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLicenseExpirationDate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::getLiveLoadBalancerServiceGraphImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLiveLoadBalancerServiceGraphImage(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::getLoadBalancers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLoadBalancers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::getManagedResourceFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetManagedResourceFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::getManagementIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetManagementIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::getNetworkVlan", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkVlan()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::getNetworkVlans", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkVlans()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::getOutboundPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOutboundPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::getPassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPassword()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::getPrimaryIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::getProjectedPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProjectedPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::getSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::getTagReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTagReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::getType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::getVirtualIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::restoreBaseConfiguration", func() {
            It("API Call Test", func() {

                _, err := sl_service.RestoreBaseConfiguration()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::restoreConfiguration", func() {
            It("API Call Test", func() {

                _, err := sl_service.RestoreConfiguration(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::saveCurrentConfiguration", func() {
            It("API Call Test", func() {

                _, err := sl_service.SaveCurrentConfiguration(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::updateLiveLoadBalancer", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateLiveLoadBalancer(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller::updateNetScalerLicense", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateNetScalerLicense()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Application_Delivery_Controller_Configuration_History service", func() {
        var sl_service services.Network_Application_Delivery_Controller_Configuration_History
        BeforeEach(func() {
            sl_service = services.GetNetworkApplicationDeliveryControllerConfigurationHistoryService(slsession)
        })


        Context("SoftLayer_Network_Application_Delivery_Controller_Configuration_History::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_Configuration_History::getController", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetController()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_Configuration_History::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute service", func() {
        var sl_service services.Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute
        BeforeEach(func() {
            sl_service = services.GetNetworkApplicationDeliveryControllerLoadBalancerHealthAttributeService(slsession)
        })


        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute::getHealthCheck", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHealthCheck()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute::getType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type service", func() {
        var sl_service services.Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type
        BeforeEach(func() {
            sl_service = services.GetNetworkApplicationDeliveryControllerLoadBalancerHealthAttributeTypeService(slsession)
        })


        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check service", func() {
        var sl_service services.Network_Application_Delivery_Controller_LoadBalancer_Health_Check
        BeforeEach(func() {
            sl_service = services.GetNetworkApplicationDeliveryControllerLoadBalancerHealthCheckService(slsession)
        })


        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check::getAttributes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttributes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check::getServices", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServices()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check::getType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type service", func() {
        var sl_service services.Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type
        BeforeEach(func() {
            sl_service = services.GetNetworkApplicationDeliveryControllerLoadBalancerHealthCheckTypeService(slsession)
        })


        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Method service", func() {
        var sl_service services.Network_Application_Delivery_Controller_LoadBalancer_Routing_Method
        BeforeEach(func() {
            sl_service = services.GetNetworkApplicationDeliveryControllerLoadBalancerRoutingMethodService(slsession)
        })


        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Method::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Method::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Type service", func() {
        var sl_service services.Network_Application_Delivery_Controller_LoadBalancer_Routing_Type
        BeforeEach(func() {
            sl_service = services.GetNetworkApplicationDeliveryControllerLoadBalancerRoutingTypeService(slsession)
        })


        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Type::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Type::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service service", func() {
        var sl_service services.Network_Application_Delivery_Controller_LoadBalancer_Service
        BeforeEach(func() {
            sl_service = services.GetNetworkApplicationDeliveryControllerLoadBalancerServiceService(slsession)
        })


        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service::deleteObject", func() {
            It("API Call Test", func() {

				err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service::getGraphImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGraphImage(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service::getGroupReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGroupReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service::getGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service::getHealthCheck", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHealthCheck()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service::getHealthChecks", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHealthChecks()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service::getIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service::getServiceGroup", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceGroup()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service::toggleStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.ToggleStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group service", func() {
        var sl_service services.Network_Application_Delivery_Controller_LoadBalancer_Service_Group
        BeforeEach(func() {
            sl_service = services.GetNetworkApplicationDeliveryControllerLoadBalancerServiceGroupService(slsession)
        })


        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group::getGraphImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGraphImage(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group::getRoutingMethod", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRoutingMethod()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group::getRoutingType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRoutingType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group::getServiceReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group::getServices", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServices()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group::getVirtualServer", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualServer()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group::getVirtualServers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualServers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group::kickAllConnections", func() {
            It("API Call Test", func() {

                _, err := sl_service.KickAllConnections()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress service", func() {
        var sl_service services.Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress
        BeforeEach(func() {
            sl_service = services.GetNetworkApplicationDeliveryControllerLoadBalancerVirtualIpAddressService(slsession)
        })


        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress::getApplicationDeliveryController", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetApplicationDeliveryController()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress::getApplicationDeliveryControllers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetApplicationDeliveryControllers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress::getAvailableSecureTransportCiphers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailableSecureTransportCiphers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress::getAvailableSecureTransportProtocols", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailableSecureTransportProtocols()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress::getDedicatedBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDedicatedBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress::getHighAvailabilityFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHighAvailabilityFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress::getIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress::getLoadBalancerHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLoadBalancerHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress::getManagedResourceFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetManagedResourceFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress::getSecureTransportCiphers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSecureTransportCiphers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress::getSecureTransportProtocols", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSecureTransportProtocols()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress::getSecurityCertificate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSecurityCertificate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress::getSecurityCertificateEntry", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSecurityCertificateEntry()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress::getVirtualServers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualServers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress::startSsl", func() {
            It("API Call Test", func() {

                _, err := sl_service.StartSsl()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress::stopSsl", func() {
            It("API Call Test", func() {

                _, err := sl_service.StopSsl()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress::upgradeConnectionLimit", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpgradeConnectionLimit()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer service", func() {
        var sl_service services.Network_Application_Delivery_Controller_LoadBalancer_VirtualServer
        BeforeEach(func() {
            sl_service = services.GetNetworkApplicationDeliveryControllerLoadBalancerVirtualServerService(slsession)
        })


        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer::deleteObject", func() {
            It("API Call Test", func() {

				err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer::getRoutingMethod", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRoutingMethod()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer::getServiceGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer::getVirtualIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer::startSsl", func() {
            It("API Call Test", func() {

                _, err := sl_service.StartSsl()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer::stopSsl", func() {
            It("API Call Test", func() {

                _, err := sl_service.StopSsl()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Backbone service", func() {
        var sl_service services.Network_Backbone
        BeforeEach(func() {
            sl_service = services.GetNetworkBackboneService(slsession)
        })


        Context("SoftLayer_Network_Backbone::getAllBackbones", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllBackbones()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Backbone::getBackbonesForLocationName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBackbonesForLocationName(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Backbone::getHealth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHealth()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Backbone::getLocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Backbone::getNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Backbone::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Backbone_Location_Dependent service", func() {
        var sl_service services.Network_Backbone_Location_Dependent
        BeforeEach(func() {
            sl_service = services.GetNetworkBackboneLocationDependentService(slsession)
        })


        Context("SoftLayer_Network_Backbone_Location_Dependent::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Backbone_Location_Dependent::getDependentLocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDependentLocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Backbone_Location_Dependent::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Backbone_Location_Dependent::getSourceDependentsByName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSourceDependentsByName(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Backbone_Location_Dependent::getSourceLocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSourceLocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Bandwidth_Version1_Allotment service", func() {
        var sl_service services.Network_Bandwidth_Version1_Allotment
        BeforeEach(func() {
            sl_service = services.GetNetworkBandwidthVersion1AllotmentService(slsession)
        })


        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getActiveDetails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveDetails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getApplicationDeliveryControllers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetApplicationDeliveryControllers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getAverageDailyPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAverageDailyPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getBandwidthAllotmentType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthAllotmentType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getBandwidthForDateRange", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthForDateRange(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getBandwidthImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthImage(nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getBareMetalInstances", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBareMetalInstances()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getBillingCycleBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingCycleBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getBillingCyclePrivateBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingCyclePrivateBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getBillingCyclePublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingCyclePublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getBillingCyclePublicUsageTotal", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingCyclePublicUsageTotal()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getCurrentBandwidthSummary", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCurrentBandwidthSummary()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getDetails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDetails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getInboundPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInboundPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getLocationGroup", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLocationGroup()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getManagedBareMetalInstances", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetManagedBareMetalInstances()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getManagedHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetManagedHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getManagedVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetManagedVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getMetricTrackingObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMetricTrackingObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getMetricTrackingObjectId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMetricTrackingObjectId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getOutboundPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOutboundPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getOverBandwidthAllocationFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOverBandwidthAllocationFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getPrivateNetworkOnlyHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateNetworkOnlyHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getProjectedOverBandwidthAllocationFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProjectedOverBandwidthAllocationFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getProjectedPublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProjectedPublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getServiceProvider", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceProvider()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getTotalBandwidthAllocated", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTotalBandwidthAllocated()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getVdrMemberRecurringFee", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVdrMemberRecurringFee()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::getVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::reassignServers", func() {
            It("API Call Test", func() {

                _, err := sl_service.ReassignServers(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::requestVdrCancellation", func() {
            It("API Call Test", func() {

                _, err := sl_service.RequestVdrCancellation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::requestVdrContentUpdates", func() {
            It("API Call Test", func() {

                _, err := sl_service.RequestVdrContentUpdates(nil,nil,nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::setVdrContent", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetVdrContent(nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::unassignServers", func() {
            It("API Call Test", func() {

                _, err := sl_service.UnassignServers(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::voidPendingServerMove", func() {
            It("API Call Test", func() {

                _, err := sl_service.VoidPendingServerMove(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Bandwidth_Version1_Allotment::voidPendingVdrCancellation", func() {
            It("API Call Test", func() {

                _, err := sl_service.VoidPendingVdrCancellation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_CdnMarketplace_Account service", func() {
        var sl_service services.Network_CdnMarketplace_Account
        BeforeEach(func() {
            sl_service = services.GetNetworkCdnMarketplaceAccountService(slsession)
        })


        Context("SoftLayer_Network_CdnMarketplace_Account::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Account::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Account::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Account::verifyCdnAccountExists", func() {
            It("API Call Test", func() {

                _, err := sl_service.VerifyCdnAccountExists(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_CdnMarketplace_Configuration_Behavior_Geoblocking service", func() {
        var sl_service services.Network_CdnMarketplace_Configuration_Behavior_Geoblocking
        BeforeEach(func() {
            sl_service = services.GetNetworkCdnMarketplaceConfigurationBehaviorGeoblockingService(slsession)
        })


        Context("SoftLayer_Network_CdnMarketplace_Configuration_Behavior_Geoblocking::createGeoblocking", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateGeoblocking(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Behavior_Geoblocking::deleteGeoblocking", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteGeoblocking(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Behavior_Geoblocking::getGeoblocking", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGeoblocking(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Behavior_Geoblocking::getGeoblockingAllowedTypesAndRegions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGeoblockingAllowedTypesAndRegions(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Behavior_Geoblocking::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Behavior_Geoblocking::updateGeoblocking", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateGeoblocking(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_CdnMarketplace_Configuration_Behavior_HotlinkProtection service", func() {
        var sl_service services.Network_CdnMarketplace_Configuration_Behavior_HotlinkProtection
        BeforeEach(func() {
            sl_service = services.GetNetworkCdnMarketplaceConfigurationBehaviorHotlinkProtectionService(slsession)
        })


        Context("SoftLayer_Network_CdnMarketplace_Configuration_Behavior_HotlinkProtection::createHotlinkProtection", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateHotlinkProtection(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Behavior_HotlinkProtection::deleteHotlinkProtection", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteHotlinkProtection(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Behavior_HotlinkProtection::getHotlinkProtection", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHotlinkProtection(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Behavior_HotlinkProtection::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Behavior_HotlinkProtection::updateHotlinkProtection", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateHotlinkProtection(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_CdnMarketplace_Configuration_Behavior_ModifyResponseHeader service", func() {
        var sl_service services.Network_CdnMarketplace_Configuration_Behavior_ModifyResponseHeader
        BeforeEach(func() {
            sl_service = services.GetNetworkCdnMarketplaceConfigurationBehaviorModifyResponseHeaderService(slsession)
        })


        Context("SoftLayer_Network_CdnMarketplace_Configuration_Behavior_ModifyResponseHeader::createModifyResponseHeader", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateModifyResponseHeader(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Behavior_ModifyResponseHeader::deleteModifyResponseHeader", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteModifyResponseHeader(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Behavior_ModifyResponseHeader::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Behavior_ModifyResponseHeader::listModifyResponseHeader", func() {
            It("API Call Test", func() {

                _, err := sl_service.ListModifyResponseHeader(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Behavior_ModifyResponseHeader::updateModifyResponseHeader", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateModifyResponseHeader(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_CdnMarketplace_Configuration_Behavior_TokenAuth service", func() {
        var sl_service services.Network_CdnMarketplace_Configuration_Behavior_TokenAuth
        BeforeEach(func() {
            sl_service = services.GetNetworkCdnMarketplaceConfigurationBehaviorTokenAuthService(slsession)
        })


        Context("SoftLayer_Network_CdnMarketplace_Configuration_Behavior_TokenAuth::createTokenAuthPath", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateTokenAuthPath(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Behavior_TokenAuth::deleteTokenAuthPath", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteTokenAuthPath(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Behavior_TokenAuth::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Behavior_TokenAuth::listTokenAuthPath", func() {
            It("API Call Test", func() {

                _, err := sl_service.ListTokenAuthPath(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Behavior_TokenAuth::updateTokenAuthPath", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateTokenAuthPath(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_CdnMarketplace_Configuration_Cache_Purge service", func() {
        var sl_service services.Network_CdnMarketplace_Configuration_Cache_Purge
        BeforeEach(func() {
            sl_service = services.GetNetworkCdnMarketplaceConfigurationCachePurgeService(slsession)
        })


        Context("SoftLayer_Network_CdnMarketplace_Configuration_Cache_Purge::createPurge", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreatePurge(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Cache_Purge::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Cache_Purge::getPurgeHistoryPerMapping", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPurgeHistoryPerMapping(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Cache_Purge::getPurgeStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPurgeStatus(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Cache_Purge::saveOrUnsavePurgePath", func() {
            It("API Call Test", func() {

                _, err := sl_service.SaveOrUnsavePurgePath(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_CdnMarketplace_Configuration_Cache_PurgeGroup service", func() {
        var sl_service services.Network_CdnMarketplace_Configuration_Cache_PurgeGroup
        BeforeEach(func() {
            sl_service = services.GetNetworkCdnMarketplaceConfigurationCachePurgeGroupService(slsession)
        })


        Context("SoftLayer_Network_CdnMarketplace_Configuration_Cache_PurgeGroup::createPurgeGroup", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreatePurgeGroup(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Cache_PurgeGroup::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Cache_PurgeGroup::getPurgeGroupByGroupId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPurgeGroupByGroupId(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Cache_PurgeGroup::getPurgeGroupQuota", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPurgeGroupQuota()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Cache_PurgeGroup::listFavoriteGroup", func() {
            It("API Call Test", func() {

                _, err := sl_service.ListFavoriteGroup(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Cache_PurgeGroup::listUnfavoriteGroup", func() {
            It("API Call Test", func() {

                _, err := sl_service.ListUnfavoriteGroup(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Cache_PurgeGroup::purgeByGroupIds", func() {
            It("API Call Test", func() {

                _, err := sl_service.PurgeByGroupIds(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Cache_PurgeGroup::removePurgeGroupFromFavorite", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemovePurgeGroupFromFavorite(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Cache_PurgeGroup::savePurgeGroupAsFavorite", func() {
            It("API Call Test", func() {

                _, err := sl_service.SavePurgeGroupAsFavorite(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_CdnMarketplace_Configuration_Cache_PurgeHistory service", func() {
        var sl_service services.Network_CdnMarketplace_Configuration_Cache_PurgeHistory
        BeforeEach(func() {
            sl_service = services.GetNetworkCdnMarketplaceConfigurationCachePurgeHistoryService(slsession)
        })


        Context("SoftLayer_Network_CdnMarketplace_Configuration_Cache_PurgeHistory::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Cache_PurgeHistory::listPurgeGroupHistory", func() {
            It("API Call Test", func() {

                _, err := sl_service.ListPurgeGroupHistory(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_CdnMarketplace_Configuration_Cache_TimeToLive service", func() {
        var sl_service services.Network_CdnMarketplace_Configuration_Cache_TimeToLive
        BeforeEach(func() {
            sl_service = services.GetNetworkCdnMarketplaceConfigurationCacheTimeToLiveService(slsession)
        })


        Context("SoftLayer_Network_CdnMarketplace_Configuration_Cache_TimeToLive::createTimeToLive", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateTimeToLive(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Cache_TimeToLive::deleteTimeToLive", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteTimeToLive(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Cache_TimeToLive::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Cache_TimeToLive::listTimeToLive", func() {
            It("API Call Test", func() {

                _, err := sl_service.ListTimeToLive(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Cache_TimeToLive::updateTimeToLive", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateTimeToLive(nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_CdnMarketplace_Configuration_Mapping service", func() {
        var sl_service services.Network_CdnMarketplace_Configuration_Mapping
        BeforeEach(func() {
            sl_service = services.GetNetworkCdnMarketplaceConfigurationMappingService(slsession)
        })


        Context("SoftLayer_Network_CdnMarketplace_Configuration_Mapping::createDomainMapping", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateDomainMapping(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Mapping::deleteDomainMapping", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteDomainMapping(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Mapping::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Mapping::listDomainMappingByUniqueId", func() {
            It("API Call Test", func() {

                _, err := sl_service.ListDomainMappingByUniqueId(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Mapping::listDomainMappings", func() {
            It("API Call Test", func() {

                _, err := sl_service.ListDomainMappings()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Mapping::retryHttpsActionRequest", func() {
            It("API Call Test", func() {

                _, err := sl_service.RetryHttpsActionRequest(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Mapping::startDomainMapping", func() {
            It("API Call Test", func() {

                _, err := sl_service.StartDomainMapping(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Mapping::stopDomainMapping", func() {
            It("API Call Test", func() {

                _, err := sl_service.StopDomainMapping(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Mapping::updateDomainMapping", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateDomainMapping(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Mapping::verifyCname", func() {
            It("API Call Test", func() {

                _, err := sl_service.VerifyCname(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Mapping::verifyDomainMapping", func() {
            It("API Call Test", func() {

                _, err := sl_service.VerifyDomainMapping(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_CdnMarketplace_Configuration_Mapping_Path service", func() {
        var sl_service services.Network_CdnMarketplace_Configuration_Mapping_Path
        BeforeEach(func() {
            sl_service = services.GetNetworkCdnMarketplaceConfigurationMappingPathService(slsession)
        })


        Context("SoftLayer_Network_CdnMarketplace_Configuration_Mapping_Path::createOriginPath", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateOriginPath(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Mapping_Path::deleteOriginPath", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteOriginPath(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Mapping_Path::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Mapping_Path::listOriginPath", func() {
            It("API Call Test", func() {

                _, err := sl_service.ListOriginPath(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Configuration_Mapping_Path::updateOriginPath", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateOriginPath(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_CdnMarketplace_Metrics service", func() {
        var sl_service services.Network_CdnMarketplace_Metrics
        BeforeEach(func() {
            sl_service = services.GetNetworkCdnMarketplaceMetricsService(slsession)
        })


        Context("SoftLayer_Network_CdnMarketplace_Metrics::getCustomerInvoicingMetrics", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCustomerInvoicingMetrics(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Metrics::getCustomerRealTimeMetrics", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCustomerRealTimeMetrics(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Metrics::getCustomerUsageMetrics", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCustomerUsageMetrics(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Metrics::getMappingBandwidthByRegionMetrics", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMappingBandwidthByRegionMetrics(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Metrics::getMappingBandwidthMetrics", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMappingBandwidthMetrics(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Metrics::getMappingHitsByTypeMetrics", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMappingHitsByTypeMetrics(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Metrics::getMappingHitsMetrics", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMappingHitsMetrics(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Metrics::getMappingIntegratedMetrics", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMappingIntegratedMetrics(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Metrics::getMappingRealTimeMetrics", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMappingRealTimeMetrics(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Metrics::getMappingUsageMetrics", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMappingUsageMetrics(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_CdnMarketplace_Vendor service", func() {
        var sl_service services.Network_CdnMarketplace_Vendor
        BeforeEach(func() {
            sl_service = services.GetNetworkCdnMarketplaceVendorService(slsession)
        })


        Context("SoftLayer_Network_CdnMarketplace_Vendor::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_CdnMarketplace_Vendor::listVendors", func() {
            It("API Call Test", func() {

                _, err := sl_service.ListVendors()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Component service", func() {
        var sl_service services.Network_Component
        BeforeEach(func() {
            sl_service = services.GetNetworkComponentService(slsession)
        })


        Context("SoftLayer_Network_Component::addNetworkVlanTrunks", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddNetworkVlanTrunks(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::clearNetworkVlanTrunks", func() {
            It("API Call Test", func() {

                _, err := sl_service.ClearNetworkVlanTrunks()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getActiveCommand", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveCommand()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getDownlinkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDownlinkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getDuplexMode", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDuplexMode()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getHighAvailabilityFirewallFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHighAvailabilityFirewallFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getIpAddressBindings", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIpAddressBindings()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getLastCommand", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLastCommand()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getMetricTrackingObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMetricTrackingObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getNetworkComponentFirewall", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkComponentFirewall()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getNetworkComponentGroup", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkComponentGroup()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getNetworkHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getNetworkVlan", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkVlan()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getNetworkVlanTrunks", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkVlanTrunks()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getNetworkVlansTrunkable", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkVlansTrunkable()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getPortStatistics", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPortStatistics()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getPrimaryIpAddressRecord", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryIpAddressRecord()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getPrimarySubnet", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimarySubnet()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getPrimaryVersion6IpAddressRecord", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryVersion6IpAddressRecord()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getRecentCommands", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRecentCommands()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getRedundancyCapableFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRedundancyCapableFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getRedundancyEnabledFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRedundancyEnabledFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getRemoteManagementUsers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRemoteManagementUsers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getRouter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRouter()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getStorageNetworkFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageNetworkFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getUplinkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUplinkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::getUplinkDuplexMode", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUplinkDuplexMode()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component::removeNetworkVlanTrunks", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveNetworkVlanTrunks(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Component_Firewall service", func() {
        var sl_service services.Network_Component_Firewall
        BeforeEach(func() {
            sl_service = services.GetNetworkComponentFirewallService(slsession)
        })


        Context("SoftLayer_Network_Component_Firewall::getApplyServerRuleSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetApplyServerRuleSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component_Firewall::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component_Firewall::getGuestNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGuestNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component_Firewall::getNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component_Firewall::getNetworkFirewallUpdateRequest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkFirewallUpdateRequest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component_Firewall::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component_Firewall::getRules", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRules()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component_Firewall::getSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Component_Firewall::hasActiveTransactions", func() {
            It("API Call Test", func() {

                _, err := sl_service.HasActiveTransactions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Customer_Subnet service", func() {
        var sl_service services.Network_Customer_Subnet
        BeforeEach(func() {
            sl_service = services.GetNetworkCustomerSubnetService(slsession)
        })


        Context("SoftLayer_Network_Customer_Subnet::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Customer_Subnet::getIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Customer_Subnet::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_DirectLink_Location service", func() {
        var sl_service services.Network_DirectLink_Location
        BeforeEach(func() {
            sl_service = services.GetNetworkDirectLinkLocationService(slsession)
        })


        Context("SoftLayer_Network_DirectLink_Location::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_DirectLink_Location::getLocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_DirectLink_Location::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_DirectLink_Location::getProvider", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProvider()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_DirectLink_Location::getServiceType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_DirectLink_Provider service", func() {
        var sl_service services.Network_DirectLink_Provider
        BeforeEach(func() {
            sl_service = services.GetNetworkDirectLinkProviderService(slsession)
        })


        Context("SoftLayer_Network_DirectLink_Provider::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_DirectLink_ServiceType service", func() {
        var sl_service services.Network_DirectLink_ServiceType
        BeforeEach(func() {
            sl_service = services.GetNetworkDirectLinkServiceTypeService(slsession)
        })


        Context("SoftLayer_Network_DirectLink_ServiceType::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Firewall_AccessControlList service", func() {
        var sl_service services.Network_Firewall_AccessControlList
        BeforeEach(func() {
            sl_service = services.GetNetworkFirewallAccessControlListService(slsession)
        })


        Context("SoftLayer_Network_Firewall_AccessControlList::getNetworkFirewallUpdateRequests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkFirewallUpdateRequests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Firewall_AccessControlList::getNetworkVlan", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkVlan()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Firewall_AccessControlList::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Firewall_AccessControlList::getRules", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRules()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Firewall_Interface service", func() {
        var sl_service services.Network_Firewall_Interface
        BeforeEach(func() {
            sl_service = services.GetNetworkFirewallInterfaceService(slsession)
        })


        Context("SoftLayer_Network_Firewall_Interface::getFirewallContextAccessControlLists", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirewallContextAccessControlLists()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Firewall_Interface::getNetworkVlan", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkVlan()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Firewall_Interface::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Firewall_Module_Context_Interface service", func() {
        var sl_service services.Network_Firewall_Module_Context_Interface
        BeforeEach(func() {
            sl_service = services.GetNetworkFirewallModuleContextInterfaceService(slsession)
        })


        Context("SoftLayer_Network_Firewall_Module_Context_Interface::getFirewallContextAccessControlLists", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirewallContextAccessControlLists()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Firewall_Module_Context_Interface::getNetworkVlan", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkVlan()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Firewall_Module_Context_Interface::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Firewall_Template service", func() {
        var sl_service services.Network_Firewall_Template
        BeforeEach(func() {
            sl_service = services.GetNetworkFirewallTemplateService(slsession)
        })


        Context("SoftLayer_Network_Firewall_Template::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Firewall_Template::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Firewall_Template::getRules", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRules()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Firewall_Update_Request service", func() {
        var sl_service services.Network_Firewall_Update_Request
        BeforeEach(func() {
            sl_service = services.GetNetworkFirewallUpdateRequestService(slsession)
        })


        Context("SoftLayer_Network_Firewall_Update_Request::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Firewall_Update_Request::getAuthorizingUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAuthorizingUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Firewall_Update_Request::getFirewallUpdateRequestRuleAttributes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirewallUpdateRequestRuleAttributes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Firewall_Update_Request::getGuest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGuest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Firewall_Update_Request::getHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Firewall_Update_Request::getNetworkComponentFirewall", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkComponentFirewall()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Firewall_Update_Request::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Firewall_Update_Request::getRules", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRules()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Firewall_Update_Request::updateRuleNote", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateRuleNote(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Firewall_Update_Request_Rule service", func() {
        var sl_service services.Network_Firewall_Update_Request_Rule
        BeforeEach(func() {
            sl_service = services.GetNetworkFirewallUpdateRequestRuleService(slsession)
        })


        Context("SoftLayer_Network_Firewall_Update_Request_Rule::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Firewall_Update_Request_Rule::getFirewallUpdateRequest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirewallUpdateRequest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Firewall_Update_Request_Rule::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Firewall_Update_Request_Rule::validateRule", func() {
            It("API Call Test", func() {

				err := sl_service.ValidateRule(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Gateway service", func() {
        var sl_service services.Network_Gateway
        BeforeEach(func() {
            sl_service = services.GetNetworkGatewayService(slsession)
        })


        Context("SoftLayer_Network_Gateway::bypassAllVlans", func() {
            It("API Call Test", func() {

				err := sl_service.BypassAllVlans()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::bypassVlans", func() {
            It("API Call Test", func() {

				err := sl_service.BypassVlans(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::changeGatewayVersion", func() {
            It("API Call Test", func() {

                _, err := sl_service.ChangeGatewayVersion(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::checkAccountWhiteList", func() {
            It("API Call Test", func() {

                _, err := sl_service.CheckAccountWhiteList(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::forceRebuildCluster", func() {
            It("API Call Test", func() {

                _, err := sl_service.ForceRebuildCluster(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::getAllowedOsPriceIds", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedOsPriceIds(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::getCapacity", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCapacity()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::getInsideVlans", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInsideVlans()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::getManufacturer", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetManufacturer(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::getMemberGatewayImagesMatch", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMemberGatewayImagesMatch()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::getMembers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMembers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::getNetworkFirewall", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkFirewall()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::getNetworkFirewallFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkFirewallFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::getPossibleInsideVlans", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPossibleInsideVlans()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::getPrivateIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::getPrivateVlan", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateVlan()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::getPublicIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::getPublicIpv6Address", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicIpv6Address()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::getPublicVlan", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicVlan()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::getRollbackSupport", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRollbackSupport()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::getStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::getUpgradeItemPrices", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpgradeItemPrices()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::isAccountWhiteListed", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsAccountWhiteListed(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::isLicenseServerAllowed", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsLicenseServerAllowed(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::isRollbackAllowed", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsRollbackAllowed()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::manageLicenses", func() {
            It("API Call Test", func() {

                _, err := sl_service.ManageLicenses(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::rebuildHACluster", func() {
            It("API Call Test", func() {

                _, err := sl_service.RebuildHACluster()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::rebuildvSRXHACluster", func() {
            It("API Call Test", func() {

                _, err := sl_service.RebuildvSRXHACluster()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::refreshGatewayLicense", func() {
            It("API Call Test", func() {

                _, err := sl_service.RefreshGatewayLicense()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::rename", func() {
            It("API Call Test", func() {

                _, err := sl_service.Rename(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::setGatewayPassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetGatewayPassword(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::unbypassAllVlans", func() {
            It("API Call Test", func() {

				err := sl_service.UnbypassAllVlans()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::unbypassVlans", func() {
            It("API Call Test", func() {

				err := sl_service.UnbypassVlans(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway::updateGatewayUserPassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateGatewayUserPassword(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Gateway_Member service", func() {
        var sl_service services.Network_Gateway_Member
        BeforeEach(func() {
            sl_service = services.GetNetworkGatewayMemberService(slsession)
        })


        Context("SoftLayer_Network_Gateway_Member::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_Member::createObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_Member::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_Member::getAttributes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttributes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_Member::getGatewaySoftwareDescription", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGatewaySoftwareDescription()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_Member::getHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_Member::getLicenses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLicenses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_Member::getNetworkGateway", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkGateway()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_Member::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_Member::getPasswords", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPasswords()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_Member::getPublicIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Gateway_Member_Attribute service", func() {
        var sl_service services.Network_Gateway_Member_Attribute
        BeforeEach(func() {
            sl_service = services.GetNetworkGatewayMemberAttributeService(slsession)
        })


        Context("SoftLayer_Network_Gateway_Member_Attribute::getGatewayMember", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGatewayMember()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_Member_Attribute::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Gateway_Precheck service", func() {
        var sl_service services.Network_Gateway_Precheck
        BeforeEach(func() {
            sl_service = services.GetNetworkGatewayPrecheckService(slsession)
        })


        Context("SoftLayer_Network_Gateway_Precheck::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_Precheck::getPrecheckStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrecheckStatus(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_Precheck::licenseManagementPrecheck", func() {
            It("API Call Test", func() {

                _, err := sl_service.LicenseManagementPrecheck(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_Precheck::osReloadPrecheck", func() {
            It("API Call Test", func() {

                _, err := sl_service.OsReloadPrecheck(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_Precheck::upgradePrecheck", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpgradePrecheck(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Gateway_Status service", func() {
        var sl_service services.Network_Gateway_Status
        BeforeEach(func() {
            sl_service = services.GetNetworkGatewayStatusService(slsession)
        })


        Context("SoftLayer_Network_Gateway_Status::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Gateway_VersionUpgrade service", func() {
        var sl_service services.Network_Gateway_VersionUpgrade
        BeforeEach(func() {
            sl_service = services.GetNetworkGatewayVersionUpgradeService(slsession)
        })


        Context("SoftLayer_Network_Gateway_VersionUpgrade::getAllByUpgradePkgUrlId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllByUpgradePkgUrlId(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_VersionUpgrade::getAllUpgradesByGatewayId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllUpgradesByGatewayId(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_VersionUpgrade::getGwOrdersAllowedLicenses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGwOrdersAllowedLicenses(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_VersionUpgrade::getGwOrdersAllowedOS", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGwOrdersAllowedOS(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_VersionUpgrade::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_VersionUpgrade::validateVersionChange", func() {
            It("API Call Test", func() {

                _, err := sl_service.ValidateVersionChange(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Gateway_Vlan service", func() {
        var sl_service services.Network_Gateway_Vlan
        BeforeEach(func() {
            sl_service = services.GetNetworkGatewayVlanService(slsession)
        })


        Context("SoftLayer_Network_Gateway_Vlan::bypass", func() {
            It("API Call Test", func() {

				err := sl_service.Bypass()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_Vlan::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_Vlan::createObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_Vlan::deleteObject", func() {
            It("API Call Test", func() {

				err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_Vlan::deleteObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_Vlan::getNetworkGateway", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkGateway()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_Vlan::getNetworkVlan", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkVlan()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_Vlan::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Gateway_Vlan::unbypass", func() {
            It("API Call Test", func() {

				err := sl_service.Unbypass()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Interconnect_Tenant service", func() {
        var sl_service services.Network_Interconnect_Tenant
        BeforeEach(func() {
            sl_service = services.GetNetworkInterconnectTenantService(slsession)
        })


        Context("SoftLayer_Network_Interconnect_Tenant::allowDeleteConnection", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowDeleteConnection(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Interconnect_Tenant::createConnection", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateConnection(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Interconnect_Tenant::deleteConnection", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteConnection(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Interconnect_Tenant::editConnection", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditConnection(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Interconnect_Tenant::getAllConnections", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllConnections()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Interconnect_Tenant::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Interconnect_Tenant::getAllPortLabelsWithCurrentUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllPortLabelsWithCurrentUsage(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Interconnect_Tenant::getBgpIpRange", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBgpIpRange()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Interconnect_Tenant::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Interconnect_Tenant::getConnection", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetConnection(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Interconnect_Tenant::getDatacenterName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDatacenterName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Interconnect_Tenant::getDirectLinkSpeeds", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDirectLinkSpeeds(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Interconnect_Tenant::getNetworkZones", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkZones()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Interconnect_Tenant::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Interconnect_Tenant::getPortLabel", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPortLabel()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Interconnect_Tenant::getPorts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPorts(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Interconnect_Tenant::getServiceType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Interconnect_Tenant::getVendorName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVendorName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Interconnect_Tenant::getZoneName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetZoneName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Interconnect_Tenant::rejectApprovalRequests", func() {
            It("API Call Test", func() {

                _, err := sl_service.RejectApprovalRequests(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Interconnect_Tenant::updateConnectionStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateConnectionStatus(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_LBaaS_HealthMonitor service", func() {
        var sl_service services.Network_LBaaS_HealthMonitor
        BeforeEach(func() {
            sl_service = services.GetNetworkLBaaSHealthMonitorService(slsession)
        })


        Context("SoftLayer_Network_LBaaS_HealthMonitor::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_HealthMonitor::updateLoadBalancerHealthMonitors", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateLoadBalancerHealthMonitors(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_LBaaS_L7Member service", func() {
        var sl_service services.Network_LBaaS_L7Member
        BeforeEach(func() {
            sl_service = services.GetNetworkLBaaSL7MemberService(slsession)
        })


        Context("SoftLayer_Network_LBaaS_L7Member::addL7PoolMembers", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddL7PoolMembers(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_L7Member::deleteL7PoolMembers", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteL7PoolMembers(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_L7Member::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_L7Member::updateL7PoolMembers", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateL7PoolMembers(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_LBaaS_L7Policy service", func() {
        var sl_service services.Network_LBaaS_L7Policy
        BeforeEach(func() {
            sl_service = services.GetNetworkLBaaSL7PolicyService(slsession)
        })


        Context("SoftLayer_Network_LBaaS_L7Policy::addL7Policies", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddL7Policies(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_L7Policy::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_L7Policy::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_L7Policy::getL7Rules", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetL7Rules()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_L7Policy::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_LBaaS_L7Pool service", func() {
        var sl_service services.Network_LBaaS_L7Pool
        BeforeEach(func() {
            sl_service = services.GetNetworkLBaaSL7PoolService(slsession)
        })


        Context("SoftLayer_Network_LBaaS_L7Pool::createL7Pool", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateL7Pool(nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_L7Pool::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_L7Pool::getL7HealthMonitor", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetL7HealthMonitor()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_L7Pool::getL7Members", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetL7Members()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_L7Pool::getL7Policies", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetL7Policies()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_L7Pool::getL7PoolMemberHealth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetL7PoolMemberHealth(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_L7Pool::getL7SessionAffinity", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetL7SessionAffinity()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_L7Pool::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_L7Pool::updateL7Pool", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateL7Pool(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_LBaaS_L7Rule service", func() {
        var sl_service services.Network_LBaaS_L7Rule
        BeforeEach(func() {
            sl_service = services.GetNetworkLBaaSL7RuleService(slsession)
        })


        Context("SoftLayer_Network_LBaaS_L7Rule::addL7Rules", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddL7Rules(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_L7Rule::deleteL7Rules", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteL7Rules(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_L7Rule::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_L7Rule::updateL7Rules", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateL7Rules(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_LBaaS_Listener service", func() {
        var sl_service services.Network_LBaaS_Listener
        BeforeEach(func() {
            sl_service = services.GetNetworkLBaaSListenerService(slsession)
        })


        Context("SoftLayer_Network_LBaaS_Listener::deleteLoadBalancerProtocols", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteLoadBalancerProtocols(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_Listener::getDefaultPool", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDefaultPool()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_Listener::getL7Policies", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetL7Policies()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_Listener::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_Listener::updateLoadBalancerProtocols", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateLoadBalancerProtocols(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_LBaaS_LoadBalancer service", func() {
        var sl_service services.Network_LBaaS_LoadBalancer
        BeforeEach(func() {
            sl_service = services.GetNetworkLBaaSLoadBalancerService(slsession)
        })


        Context("SoftLayer_Network_LBaaS_LoadBalancer::cancelLoadBalancer", func() {
            It("API Call Test", func() {

                _, err := sl_service.CancelLoadBalancer(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_LoadBalancer::enableOrDisableDataLogs", func() {
            It("API Call Test", func() {

                _, err := sl_service.EnableOrDisableDataLogs(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_LoadBalancer::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_LoadBalancer::getAppliances", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAppliances(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_LoadBalancer::getDatacenter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDatacenter()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_LoadBalancer::getHealthMonitors", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHealthMonitors()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_LoadBalancer::getL7Pools", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetL7Pools()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_LoadBalancer::getListenerTimeSeriesData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetListenerTimeSeriesData(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_LoadBalancer::getListeners", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetListeners()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_LoadBalancer::getLoadBalancer", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLoadBalancer(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_LoadBalancer::getLoadBalancerMemberHealth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLoadBalancerMemberHealth(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_LoadBalancer::getLoadBalancerStatistics", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLoadBalancerStatistics(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_LoadBalancer::getLoadBalancers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLoadBalancers(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_LoadBalancer::getMembers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMembers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_LoadBalancer::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_LoadBalancer::getSslCiphers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSslCiphers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_LoadBalancer::serviceDNS", func() {
            It("API Call Test", func() {

				err := sl_service.ServiceDNS(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_LoadBalancer::serviceLoadBalancer", func() {
            It("API Call Test", func() {

                _, err := sl_service.ServiceLoadBalancer(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_LoadBalancer::updateLoadBalancer", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateLoadBalancer(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_LoadBalancer::updateSslCiphers", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateSslCiphers(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_LBaaS_LoadBalancerAppliance service", func() {
        var sl_service services.Network_LBaaS_LoadBalancerAppliance
        BeforeEach(func() {
            sl_service = services.GetNetworkLBaaSLoadBalancerApplianceService(slsession)
        })


        Context("SoftLayer_Network_LBaaS_LoadBalancerAppliance::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_LBaaS_Member service", func() {
        var sl_service services.Network_LBaaS_Member
        BeforeEach(func() {
            sl_service = services.GetNetworkLBaaSMemberService(slsession)
        })


        Context("SoftLayer_Network_LBaaS_Member::addLoadBalancerMembers", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddLoadBalancerMembers(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_Member::deleteLoadBalancerMembers", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteLoadBalancerMembers(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_Member::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_Member::updateLoadBalancerMembers", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateLoadBalancerMembers(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_LBaaS_SSLCipher service", func() {
        var sl_service services.Network_LBaaS_SSLCipher
        BeforeEach(func() {
            sl_service = services.GetNetworkLBaaSSSLCipherService(slsession)
        })


        Context("SoftLayer_Network_LBaaS_SSLCipher::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LBaaS_SSLCipher::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_LoadBalancer_Global_Account service", func() {
        var sl_service services.Network_LoadBalancer_Global_Account
        BeforeEach(func() {
            sl_service = services.GetNetworkLoadBalancerGlobalAccountService(slsession)
        })


        Context("SoftLayer_Network_LoadBalancer_Global_Account::addNsRecord", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddNsRecord()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LoadBalancer_Global_Account::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LoadBalancer_Global_Account::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LoadBalancer_Global_Account::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LoadBalancer_Global_Account::getHosts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHosts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LoadBalancer_Global_Account::getLoadBalanceType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLoadBalanceType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LoadBalancer_Global_Account::getManagedResourceFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetManagedResourceFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LoadBalancer_Global_Account::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LoadBalancer_Global_Account::removeNsRecord", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveNsRecord()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_LoadBalancer_Global_Host service", func() {
        var sl_service services.Network_LoadBalancer_Global_Host
        BeforeEach(func() {
            sl_service = services.GetNetworkLoadBalancerGlobalHostService(slsession)
        })


        Context("SoftLayer_Network_LoadBalancer_Global_Host::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LoadBalancer_Global_Host::getLoadBalancerAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLoadBalancerAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LoadBalancer_Global_Host::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_LoadBalancer_Service service", func() {
        var sl_service services.Network_LoadBalancer_Service
        BeforeEach(func() {
            sl_service = services.GetNetworkLoadBalancerServiceService(slsession)
        })


        Context("SoftLayer_Network_LoadBalancer_Service::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LoadBalancer_Service::getGraphImage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGraphImage(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LoadBalancer_Service::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LoadBalancer_Service::getStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LoadBalancer_Service::getVip", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVip()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LoadBalancer_Service::resetPeakConnections", func() {
            It("API Call Test", func() {

                _, err := sl_service.ResetPeakConnections()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_LoadBalancer_VirtualIpAddress service", func() {
        var sl_service services.Network_LoadBalancer_VirtualIpAddress
        BeforeEach(func() {
            sl_service = services.GetNetworkLoadBalancerVirtualIpAddressService(slsession)
        })


        Context("SoftLayer_Network_LoadBalancer_VirtualIpAddress::disable", func() {
            It("API Call Test", func() {

                _, err := sl_service.Disable()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LoadBalancer_VirtualIpAddress::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LoadBalancer_VirtualIpAddress::enable", func() {
            It("API Call Test", func() {

                _, err := sl_service.Enable()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LoadBalancer_VirtualIpAddress::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LoadBalancer_VirtualIpAddress::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LoadBalancer_VirtualIpAddress::getCustomerManagedFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCustomerManagedFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LoadBalancer_VirtualIpAddress::getManagedResourceFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetManagedResourceFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LoadBalancer_VirtualIpAddress::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LoadBalancer_VirtualIpAddress::getServices", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServices()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LoadBalancer_VirtualIpAddress::kickAllConnections", func() {
            It("API Call Test", func() {

                _, err := sl_service.KickAllConnections()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_LoadBalancer_VirtualIpAddress::upgradeConnectionLimit", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpgradeConnectionLimit()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Message_Delivery service", func() {
        var sl_service services.Network_Message_Delivery
        BeforeEach(func() {
            sl_service = services.GetNetworkMessageDeliveryService(slsession)
        })


        Context("SoftLayer_Network_Message_Delivery::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Message_Delivery::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Message_Delivery::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Message_Delivery::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Message_Delivery::getType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Message_Delivery::getUpgradeItemPrices", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpgradeItemPrices()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Message_Delivery::getVendor", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVendor()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Message_Delivery_Email_Sendgrid service", func() {
        var sl_service services.Network_Message_Delivery_Email_Sendgrid
        BeforeEach(func() {
            sl_service = services.GetNetworkMessageDeliveryEmailSendgridService(slsession)
        })


        Context("SoftLayer_Network_Message_Delivery_Email_Sendgrid::addUnsubscribeEmailAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddUnsubscribeEmailAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Message_Delivery_Email_Sendgrid::deleteEmailListEntries", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteEmailListEntries(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Message_Delivery_Email_Sendgrid::disableSmtpAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.DisableSmtpAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Message_Delivery_Email_Sendgrid::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Message_Delivery_Email_Sendgrid::enableSmtpAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.EnableSmtpAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Message_Delivery_Email_Sendgrid::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Message_Delivery_Email_Sendgrid::getAccountOverview", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccountOverview()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Message_Delivery_Email_Sendgrid::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Message_Delivery_Email_Sendgrid::getEmailAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEmailAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Message_Delivery_Email_Sendgrid::getEmailList", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEmailList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Message_Delivery_Email_Sendgrid::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Message_Delivery_Email_Sendgrid::getOfferingsList", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOfferingsList()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Message_Delivery_Email_Sendgrid::getSmtpAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSmtpAccess()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Message_Delivery_Email_Sendgrid::getStatistics", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStatistics(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Message_Delivery_Email_Sendgrid::getStatisticsGraph", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStatisticsGraph(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Message_Delivery_Email_Sendgrid::getType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Message_Delivery_Email_Sendgrid::getUpgradeItemPrices", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpgradeItemPrices()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Message_Delivery_Email_Sendgrid::getVendor", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVendor()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Message_Delivery_Email_Sendgrid::singleSignOn", func() {
            It("API Call Test", func() {

                _, err := sl_service.SingleSignOn()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Message_Delivery_Email_Sendgrid::updateEmailAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateEmailAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Monitor service", func() {
        var sl_service services.Network_Monitor
        BeforeEach(func() {
            sl_service = services.GetNetworkMonitorService(slsession)
        })


        Context("SoftLayer_Network_Monitor::getIpAddressesByHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIpAddressesByHardware(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Monitor::getIpAddressesByVirtualGuest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIpAddressesByVirtualGuest(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Monitor_Version1_Query_Host service", func() {
        var sl_service services.Network_Monitor_Version1_Query_Host
        BeforeEach(func() {
            sl_service = services.GetNetworkMonitorVersion1QueryHostService(slsession)
        })


        Context("SoftLayer_Network_Monitor_Version1_Query_Host::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Monitor_Version1_Query_Host::createObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Monitor_Version1_Query_Host::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Monitor_Version1_Query_Host::deleteObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Monitor_Version1_Query_Host::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Monitor_Version1_Query_Host::editObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Monitor_Version1_Query_Host::findByHardwareId", func() {
            It("API Call Test", func() {

                _, err := sl_service.FindByHardwareId(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Monitor_Version1_Query_Host::getHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Monitor_Version1_Query_Host::getLastResult", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLastResult()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Monitor_Version1_Query_Host::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Monitor_Version1_Query_Host::getQueryType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetQueryType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Monitor_Version1_Query_Host::getResponseAction", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResponseAction()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Monitor_Version1_Query_Host_Stratum service", func() {
        var sl_service services.Network_Monitor_Version1_Query_Host_Stratum
        BeforeEach(func() {
            sl_service = services.GetNetworkMonitorVersion1QueryHostStratumService(slsession)
        })


        Context("SoftLayer_Network_Monitor_Version1_Query_Host_Stratum::getAllQueryTypes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllQueryTypes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Monitor_Version1_Query_Host_Stratum::getAllResponseTypes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllResponseTypes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Monitor_Version1_Query_Host_Stratum::getHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Monitor_Version1_Query_Host_Stratum::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Pod service", func() {
        var sl_service services.Network_Pod
        BeforeEach(func() {
            sl_service = services.GetNetworkPodService(slsession)
        })


        Context("SoftLayer_Network_Pod::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Pod::getCapabilities", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCapabilities()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Pod::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Pod::listCapabilities", func() {
            It("API Call Test", func() {

                _, err := sl_service.ListCapabilities()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_SecurityGroup service", func() {
        var sl_service services.Network_SecurityGroup
        BeforeEach(func() {
            sl_service = services.GetNetworkSecurityGroupService(slsession)
        })


        Context("SoftLayer_Network_SecurityGroup::addRules", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddRules(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_SecurityGroup::attachNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.AttachNetworkComponents(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_SecurityGroup::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_SecurityGroup::createObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_SecurityGroup::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_SecurityGroup::deleteObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_SecurityGroup::detachNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.DetachNetworkComponents(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_SecurityGroup::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_SecurityGroup::editObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_SecurityGroup::editRules", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditRules(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_SecurityGroup::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_SecurityGroup::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_SecurityGroup::getLimits", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLimits()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_SecurityGroup::getNetworkComponentBindings", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkComponentBindings()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_SecurityGroup::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_SecurityGroup::getOrderBindings", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOrderBindings()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_SecurityGroup::getRules", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRules()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_SecurityGroup::getSupportedDataCenters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSupportedDataCenters()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_SecurityGroup::removeRules", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveRules(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Security_Scanner_Request service", func() {
        var sl_service services.Network_Security_Scanner_Request
        BeforeEach(func() {
            sl_service = services.GetNetworkSecurityScannerRequestService(slsession)
        })


        Context("SoftLayer_Network_Security_Scanner_Request::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Security_Scanner_Request::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Security_Scanner_Request::getGuest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGuest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Security_Scanner_Request::getHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Security_Scanner_Request::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Security_Scanner_Request::getReport", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReport()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Security_Scanner_Request::getRequestorOwnedFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRequestorOwnedFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Security_Scanner_Request::getStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Service_Vpn_Overrides service", func() {
        var sl_service services.Network_Service_Vpn_Overrides
        BeforeEach(func() {
            sl_service = services.GetNetworkServiceVpnOverridesService(slsession)
        })


        Context("SoftLayer_Network_Service_Vpn_Overrides::createObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Service_Vpn_Overrides::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Service_Vpn_Overrides::deleteObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Service_Vpn_Overrides::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Service_Vpn_Overrides::getSubnet", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubnet()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Service_Vpn_Overrides::getUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Storage service", func() {
        var sl_service services.Network_Storage
        BeforeEach(func() {
            sl_service = services.GetNetworkStorageService(slsession)
        })


        Context("SoftLayer_Network_Storage::allowAccessFromHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromHardware(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::allowAccessFromHardwareList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromHardwareList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::allowAccessFromHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromHost(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::allowAccessFromHostList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromHostList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::allowAccessFromIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromIpAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::allowAccessFromIpAddressList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromIpAddressList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::allowAccessFromSubnet", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromSubnet(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::allowAccessFromSubnetList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromSubnetList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::allowAccessFromVirtualGuest", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromVirtualGuest(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::allowAccessFromVirtualGuestList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromVirtualGuestList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::allowAccessToReplicantFromHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToReplicantFromHardware(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::allowAccessToReplicantFromHardwareList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToReplicantFromHardwareList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::allowAccessToReplicantFromIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToReplicantFromIpAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::allowAccessToReplicantFromIpAddressList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToReplicantFromIpAddressList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::allowAccessToReplicantFromSubnet", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToReplicantFromSubnet(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::allowAccessToReplicantFromSubnetList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToReplicantFromSubnetList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::allowAccessToReplicantFromVirtualGuest", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToReplicantFromVirtualGuest(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::allowAccessToReplicantFromVirtualGuestList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToReplicantFromVirtualGuestList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::assignCredential", func() {
            It("API Call Test", func() {

                _, err := sl_service.AssignCredential(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::assignNewCredential", func() {
            It("API Call Test", func() {

                _, err := sl_service.AssignNewCredential(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::changePassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.ChangePassword(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::collectBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.CollectBandwidth(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::collectBytesUsed", func() {
            It("API Call Test", func() {

                _, err := sl_service.CollectBytesUsed()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::convertCloneDependentToIndependent", func() {
            It("API Call Test", func() {

                _, err := sl_service.ConvertCloneDependentToIndependent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::createFolder", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateFolder(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::createOrUpdateLunId", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateOrUpdateLunId(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::createSnapshot", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateSnapshot(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::deleteAllFiles", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteAllFiles()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::deleteFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteFile(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::deleteFiles", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteFiles(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::deleteFolder", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteFolder(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::disableSnapshots", func() {
            It("API Call Test", func() {

                _, err := sl_service.DisableSnapshots(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::disasterRecoveryFailoverToReplicant", func() {
            It("API Call Test", func() {

                _, err := sl_service.DisasterRecoveryFailoverToReplicant(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::downloadFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.DownloadFile(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::editCredential", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditCredential(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::enableSnapshots", func() {
            It("API Call Test", func() {

                _, err := sl_service.EnableSnapshots(nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::failbackFromReplicant", func() {
            It("API Call Test", func() {

                _, err := sl_service.FailbackFromReplicant()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::failoverToReplicant", func() {
            It("API Call Test", func() {

                _, err := sl_service.FailoverToReplicant(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getAccountPassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccountPassword()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getActiveTransactions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveTransactions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getAllFiles", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllFiles()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getAllFilesByFilter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllFilesByFilter(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getAllowDisasterRecoveryFailback", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowDisasterRecoveryFailback()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getAllowDisasterRecoveryFailover", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowDisasterRecoveryFailover()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getAllowableHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowableHardware(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getAllowableIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowableIpAddresses(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getAllowableSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowableSubnets(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getAllowableVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowableVirtualGuests(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getAllowedHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getAllowedHostsLimit", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedHostsLimit()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getAllowedIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getAllowedReplicationHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedReplicationHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getAllowedReplicationIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedReplicationIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getAllowedReplicationSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedReplicationSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getAllowedReplicationVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedReplicationVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getAllowedSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getAllowedVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getBillingItemCategory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItemCategory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getByUsername", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetByUsername(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getBytesUsed", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBytesUsed()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getCdnUrls", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCdnUrls()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getClusterResource", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetClusterResource()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getCreationScheduleId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCreationScheduleId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getCredentials", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCredentials()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getDailySchedule", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDailySchedule()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getDependentDuplicate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDependentDuplicate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getDependentDuplicates", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDependentDuplicates()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getDuplicateConversionStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDuplicateConversionStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getEvents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEvents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getFailbackNotAllowed", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFailbackNotAllowed()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getFailoverNotAllowed", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFailoverNotAllowed()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getFileBlockEncryptedLocations", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFileBlockEncryptedLocations()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getFileByIdentifier", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFileByIdentifier(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getFileCount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFileCount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getFileList", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFileList(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getFileNetworkMountAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFileNetworkMountAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getFilePendingDeleteCount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFilePendingDeleteCount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getFilesPendingDelete", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFilesPendingDelete()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getFixReplicationCurrentStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFixReplicationCurrentStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getFolderList", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFolderList()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getGraph", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGraph(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getHasEncryptionAtRest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasEncryptionAtRest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getHourlySchedule", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHourlySchedule()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getIntervalSchedule", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIntervalSchedule()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getIops", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIops()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getIsConvertToIndependentTransactionInProgress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsConvertToIndependentTransactionInProgress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getIsDependentDuplicateProvisionCompleted", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsDependentDuplicateProvisionCompleted()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getIsInDedicatedServiceResource", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsInDedicatedServiceResource()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getIsMagneticStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsMagneticStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getIsProvisionInProgress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsProvisionInProgress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getIsReadyForSnapshot", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsReadyForSnapshot()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getIsReadyToMount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsReadyToMount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getIscsiLuns", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIscsiLuns()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getIscsiReplicatingVolume", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIscsiReplicatingVolume()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getIscsiTargetIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIscsiTargetIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getLunId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLunId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getManualSnapshots", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetManualSnapshots()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getMaximumExpansionSize", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMaximumExpansionSize()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getMetricTrackingObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMetricTrackingObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getMountPath", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMountPath()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getMountableFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMountableFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getMoveAndSplitStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMoveAndSplitStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getNetworkConnectionDetails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkConnectionDetails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getNetworkMountAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMountAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getNetworkMountPath", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMountPath()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getNotificationSubscribers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNotificationSubscribers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getObjectStorageConnectionInformation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObjectStorageConnectionInformation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getObjectsByCredential", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObjectsByCredential(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getOriginalSnapshotName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOriginalSnapshotName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getOriginalVolumeId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOriginalVolumeId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getOriginalVolumeName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOriginalVolumeName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getOriginalVolumeSize", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOriginalVolumeSize()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getOsType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOsType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getOsTypeId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOsTypeId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getParentPartnerships", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetParentPartnerships()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getParentVolume", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetParentVolume()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getPartnerships", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPartnerships()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getPermissionsGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPermissionsGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getProperties", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProperties()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getProvisionedIops", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProvisionedIops()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getRecycleBinFileByIdentifier", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRecycleBinFileByIdentifier(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getRemainingAllowedHosts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRemainingAllowedHosts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getRemainingAllowedHostsForReplicant", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRemainingAllowedHostsForReplicant()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getReplicatingLuns", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReplicatingLuns()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getReplicatingVolume", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReplicatingVolume()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getReplicationEvents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReplicationEvents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getReplicationPartners", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReplicationPartners()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getReplicationSchedule", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReplicationSchedule()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getReplicationStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReplicationStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getReplicationTimestamp", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReplicationTimestamp()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getSchedules", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSchedules()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getServiceResource", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceResource()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getServiceResourceBackendIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceResourceBackendIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getServiceResourceName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceResourceName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getSnapshotCapacityGb", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSnapshotCapacityGb()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getSnapshotCreationTimestamp", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSnapshotCreationTimestamp()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getSnapshotDeletionThresholdPercentage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSnapshotDeletionThresholdPercentage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getSnapshotNotificationStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSnapshotNotificationStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getSnapshotSizeBytes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSnapshotSizeBytes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getSnapshotSpaceAvailable", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSnapshotSpaceAvailable()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getSnapshots", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSnapshots()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getSnapshotsForVolume", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSnapshotsForVolume()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getStaasVersion", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStaasVersion()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getStorageGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getStorageGroupsNetworkConnectionDetails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageGroupsNetworkConnectionDetails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getStorageTierLevel", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageTierLevel()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getStorageType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getTargetIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTargetIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getTotalBytesUsed", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTotalBytesUsed()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getTotalScheduleSnapshotRetentionCount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTotalScheduleSnapshotRetentionCount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getUsageNotification", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUsageNotification()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getValidReplicationTargetDatacenterLocations", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetValidReplicationTargetDatacenterLocations()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getVendorName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVendorName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getVirtualGuest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getVolumeCountLimits", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVolumeCountLimits()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getVolumeDuplicateParameters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVolumeDuplicateParameters()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getVolumeHistory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVolumeHistory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getVolumeStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVolumeStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getWebccAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetWebccAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::getWeeklySchedule", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetWeeklySchedule()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::immediateFailoverToReplicant", func() {
            It("API Call Test", func() {

                _, err := sl_service.ImmediateFailoverToReplicant(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::initiateOriginVolumeReclaim", func() {
            It("API Call Test", func() {

                _, err := sl_service.InitiateOriginVolumeReclaim()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::initiateVolumeCutover", func() {
            It("API Call Test", func() {

                _, err := sl_service.InitiateVolumeCutover()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::isBlockingOperationInProgress", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsBlockingOperationInProgress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::isDuplicateReadyForSnapshot", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsDuplicateReadyForSnapshot()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::isDuplicateReadyToMount", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsDuplicateReadyToMount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::isVolumeActive", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsVolumeActive()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::refreshDependentDuplicate", func() {
            It("API Call Test", func() {

                _, err := sl_service.RefreshDependentDuplicate(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::refreshDuplicate", func() {
            It("API Call Test", func() {

                _, err := sl_service.RefreshDuplicate(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::removeAccessFromHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromHardware(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::removeAccessFromHardwareList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromHardwareList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::removeAccessFromHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromHost(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::removeAccessFromHostList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromHostList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::removeAccessFromIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromIpAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::removeAccessFromIpAddressList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromIpAddressList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::removeAccessFromSubnet", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromSubnet(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::removeAccessFromSubnetList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromSubnetList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::removeAccessFromVirtualGuest", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromVirtualGuest(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::removeAccessFromVirtualGuestList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromVirtualGuestList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::removeAccessToReplicantFromHardwareList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToReplicantFromHardwareList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::removeAccessToReplicantFromIpAddressList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToReplicantFromIpAddressList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::removeAccessToReplicantFromSubnet", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToReplicantFromSubnet(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::removeAccessToReplicantFromSubnetList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToReplicantFromSubnetList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::removeAccessToReplicantFromVirtualGuestList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToReplicantFromVirtualGuestList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::removeCredential", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveCredential(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::restoreFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.RestoreFile(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::restoreFromSnapshot", func() {
            It("API Call Test", func() {

                _, err := sl_service.RestoreFromSnapshot(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::sendPasswordReminderEmail", func() {
            It("API Call Test", func() {

                _, err := sl_service.SendPasswordReminderEmail(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::setMountable", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetMountable(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::setSnapshotAllocation", func() {
            It("API Call Test", func() {

				err := sl_service.SetSnapshotAllocation(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::setSnapshotNotification", func() {
            It("API Call Test", func() {

				err := sl_service.SetSnapshotNotification(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::upgradeVolumeCapacity", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpgradeVolumeCapacity(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::uploadFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.UploadFile(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage::validateHostsAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.ValidateHostsAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Storage_Allowed_Host service", func() {
        var sl_service services.Network_Storage_Allowed_Host
        BeforeEach(func() {
            sl_service = services.GetNetworkStorageAllowedHostService(slsession)
        })


        Context("SoftLayer_Network_Storage_Allowed_Host::assignSubnetsToAcl", func() {
            It("API Call Test", func() {

                _, err := sl_service.AssignSubnetsToAcl(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host::getAssignedGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host::getAssignedIscsiVolumes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedIscsiVolumes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host::getAssignedNfsVolumes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedNfsVolumes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host::getAssignedReplicationVolumes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedReplicationVolumes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host::getAssignedVolumes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedVolumes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host::getCredential", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCredential()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host::getSourceSubnet", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSourceSubnet()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host::getSubnetsInAcl", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubnetsInAcl()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host::removeSubnetsFromAcl", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveSubnetsFromAcl(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host::setCredentialPassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetCredentialPassword(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Storage_Allowed_Host_Hardware service", func() {
        var sl_service services.Network_Storage_Allowed_Host_Hardware
        BeforeEach(func() {
            sl_service = services.GetNetworkStorageAllowedHostHardwareService(slsession)
        })


        Context("SoftLayer_Network_Storage_Allowed_Host_Hardware::assignSubnetsToAcl", func() {
            It("API Call Test", func() {

                _, err := sl_service.AssignSubnetsToAcl(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Hardware::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Hardware::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Hardware::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Hardware::getAssignedGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Hardware::getAssignedIscsiVolumes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedIscsiVolumes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Hardware::getAssignedNfsVolumes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedNfsVolumes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Hardware::getAssignedReplicationVolumes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedReplicationVolumes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Hardware::getAssignedVolumes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedVolumes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Hardware::getCredential", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCredential()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Hardware::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Hardware::getResource", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResource()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Hardware::getSourceSubnet", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSourceSubnet()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Hardware::getSubnetsInAcl", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubnetsInAcl()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Hardware::removeSubnetsFromAcl", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveSubnetsFromAcl(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Hardware::setCredentialPassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetCredentialPassword(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Storage_Allowed_Host_IpAddress service", func() {
        var sl_service services.Network_Storage_Allowed_Host_IpAddress
        BeforeEach(func() {
            sl_service = services.GetNetworkStorageAllowedHostIpAddressService(slsession)
        })


        Context("SoftLayer_Network_Storage_Allowed_Host_IpAddress::assignSubnetsToAcl", func() {
            It("API Call Test", func() {

                _, err := sl_service.AssignSubnetsToAcl(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_IpAddress::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_IpAddress::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_IpAddress::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_IpAddress::getAssignedGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_IpAddress::getAssignedIscsiVolumes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedIscsiVolumes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_IpAddress::getAssignedNfsVolumes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedNfsVolumes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_IpAddress::getAssignedReplicationVolumes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedReplicationVolumes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_IpAddress::getAssignedVolumes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedVolumes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_IpAddress::getCredential", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCredential()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_IpAddress::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_IpAddress::getResource", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResource()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_IpAddress::getSourceSubnet", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSourceSubnet()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_IpAddress::getSubnetsInAcl", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubnetsInAcl()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_IpAddress::removeSubnetsFromAcl", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveSubnetsFromAcl(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_IpAddress::setCredentialPassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetCredentialPassword(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Storage_Allowed_Host_Subnet service", func() {
        var sl_service services.Network_Storage_Allowed_Host_Subnet
        BeforeEach(func() {
            sl_service = services.GetNetworkStorageAllowedHostSubnetService(slsession)
        })


        Context("SoftLayer_Network_Storage_Allowed_Host_Subnet::assignSubnetsToAcl", func() {
            It("API Call Test", func() {

                _, err := sl_service.AssignSubnetsToAcl(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Subnet::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Subnet::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Subnet::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Subnet::getAssignedGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Subnet::getAssignedIscsiVolumes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedIscsiVolumes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Subnet::getAssignedNfsVolumes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedNfsVolumes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Subnet::getAssignedReplicationVolumes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedReplicationVolumes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Subnet::getAssignedVolumes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedVolumes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Subnet::getCredential", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCredential()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Subnet::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Subnet::getResource", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResource()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Subnet::getSourceSubnet", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSourceSubnet()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Subnet::getSubnetsInAcl", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubnetsInAcl()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Subnet::removeSubnetsFromAcl", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveSubnetsFromAcl(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_Subnet::setCredentialPassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetCredentialPassword(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Storage_Allowed_Host_VirtualGuest service", func() {
        var sl_service services.Network_Storage_Allowed_Host_VirtualGuest
        BeforeEach(func() {
            sl_service = services.GetNetworkStorageAllowedHostVirtualGuestService(slsession)
        })


        Context("SoftLayer_Network_Storage_Allowed_Host_VirtualGuest::assignSubnetsToAcl", func() {
            It("API Call Test", func() {

                _, err := sl_service.AssignSubnetsToAcl(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_VirtualGuest::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_VirtualGuest::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_VirtualGuest::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_VirtualGuest::getAssignedGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_VirtualGuest::getAssignedIscsiVolumes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedIscsiVolumes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_VirtualGuest::getAssignedNfsVolumes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedNfsVolumes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_VirtualGuest::getAssignedReplicationVolumes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedReplicationVolumes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_VirtualGuest::getAssignedVolumes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAssignedVolumes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_VirtualGuest::getCredential", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCredential()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_VirtualGuest::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_VirtualGuest::getResource", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetResource()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_VirtualGuest::getSourceSubnet", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSourceSubnet()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_VirtualGuest::getSubnetsInAcl", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubnetsInAcl()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_VirtualGuest::removeSubnetsFromAcl", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveSubnetsFromAcl(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Allowed_Host_VirtualGuest::setCredentialPassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetCredentialPassword(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Storage_Backup_Evault service", func() {
        var sl_service services.Network_Storage_Backup_Evault
        BeforeEach(func() {
            sl_service = services.GetNetworkStorageBackupEvaultService(slsession)
        })


        Context("SoftLayer_Network_Storage_Backup_Evault::allowAccessFromHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromHardware(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::allowAccessFromHardwareList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromHardwareList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::allowAccessFromHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromHost(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::allowAccessFromHostList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromHostList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::allowAccessFromIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromIpAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::allowAccessFromIpAddressList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromIpAddressList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::allowAccessFromSubnet", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromSubnet(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::allowAccessFromSubnetList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromSubnetList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::allowAccessFromVirtualGuest", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromVirtualGuest(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::allowAccessFromVirtualGuestList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromVirtualGuestList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::allowAccessToReplicantFromHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToReplicantFromHardware(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::allowAccessToReplicantFromHardwareList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToReplicantFromHardwareList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::allowAccessToReplicantFromIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToReplicantFromIpAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::allowAccessToReplicantFromIpAddressList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToReplicantFromIpAddressList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::allowAccessToReplicantFromSubnet", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToReplicantFromSubnet(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::allowAccessToReplicantFromSubnetList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToReplicantFromSubnetList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::allowAccessToReplicantFromVirtualGuest", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToReplicantFromVirtualGuest(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::allowAccessToReplicantFromVirtualGuestList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToReplicantFromVirtualGuestList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::assignCredential", func() {
            It("API Call Test", func() {

                _, err := sl_service.AssignCredential(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::assignNewCredential", func() {
            It("API Call Test", func() {

                _, err := sl_service.AssignNewCredential(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::changePassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.ChangePassword(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::collectBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.CollectBandwidth(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::collectBytesUsed", func() {
            It("API Call Test", func() {

                _, err := sl_service.CollectBytesUsed()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::convertCloneDependentToIndependent", func() {
            It("API Call Test", func() {

                _, err := sl_service.ConvertCloneDependentToIndependent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::createFolder", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateFolder(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::createOrUpdateLunId", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateOrUpdateLunId(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::createSnapshot", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateSnapshot(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::deleteAllFiles", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteAllFiles()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::deleteFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteFile(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::deleteFiles", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteFiles(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::deleteFolder", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteFolder(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::deleteTasks", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteTasks(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::disableSnapshots", func() {
            It("API Call Test", func() {

                _, err := sl_service.DisableSnapshots(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::disasterRecoveryFailoverToReplicant", func() {
            It("API Call Test", func() {

                _, err := sl_service.DisasterRecoveryFailoverToReplicant(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::downloadFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.DownloadFile(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::editCredential", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditCredential(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::enableSnapshots", func() {
            It("API Call Test", func() {

                _, err := sl_service.EnableSnapshots(nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::failbackFromReplicant", func() {
            It("API Call Test", func() {

                _, err := sl_service.FailbackFromReplicant()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::failoverToReplicant", func() {
            It("API Call Test", func() {

                _, err := sl_service.FailoverToReplicant(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getAccountPassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccountPassword()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getActiveTransactions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveTransactions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getAllFiles", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllFiles()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getAllFilesByFilter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllFilesByFilter(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getAllowDisasterRecoveryFailback", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowDisasterRecoveryFailback()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getAllowDisasterRecoveryFailover", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowDisasterRecoveryFailover()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getAllowableHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowableHardware(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getAllowableIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowableIpAddresses(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getAllowableSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowableSubnets(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getAllowableVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowableVirtualGuests(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getAllowedHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getAllowedHostsLimit", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedHostsLimit()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getAllowedIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getAllowedReplicationHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedReplicationHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getAllowedReplicationIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedReplicationIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getAllowedReplicationSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedReplicationSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getAllowedReplicationVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedReplicationVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getAllowedSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getAllowedVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getBillingItemCategory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItemCategory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getByUsername", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetByUsername(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getBytesUsed", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBytesUsed()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getCdnUrls", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCdnUrls()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getClusterResource", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetClusterResource()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getCreationScheduleId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCreationScheduleId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getCredentials", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCredentials()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getDailySchedule", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDailySchedule()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getDependentDuplicate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDependentDuplicate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getDependentDuplicates", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDependentDuplicates()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getDuplicateConversionStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDuplicateConversionStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getEvents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEvents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getFailbackNotAllowed", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFailbackNotAllowed()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getFailoverNotAllowed", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFailoverNotAllowed()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getFileBlockEncryptedLocations", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFileBlockEncryptedLocations()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getFileByIdentifier", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFileByIdentifier(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getFileCount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFileCount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getFileList", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFileList(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getFileNetworkMountAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFileNetworkMountAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getFilePendingDeleteCount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFilePendingDeleteCount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getFilesPendingDelete", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFilesPendingDelete()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getFixReplicationCurrentStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFixReplicationCurrentStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getFolderList", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFolderList()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getGraph", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGraph(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getHardwareWithEvaultFirst", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareWithEvaultFirst(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getHasEncryptionAtRest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasEncryptionAtRest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getHourlySchedule", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHourlySchedule()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getIntervalSchedule", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIntervalSchedule()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getIops", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIops()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getIsConvertToIndependentTransactionInProgress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsConvertToIndependentTransactionInProgress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getIsDependentDuplicateProvisionCompleted", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsDependentDuplicateProvisionCompleted()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getIsInDedicatedServiceResource", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsInDedicatedServiceResource()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getIsMagneticStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsMagneticStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getIsProvisionInProgress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsProvisionInProgress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getIsReadyForSnapshot", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsReadyForSnapshot()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getIsReadyToMount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsReadyToMount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getIscsiLuns", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIscsiLuns()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getIscsiReplicatingVolume", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIscsiReplicatingVolume()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getIscsiTargetIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIscsiTargetIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getLunId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLunId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getManualSnapshots", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetManualSnapshots()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getMaximumExpansionSize", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMaximumExpansionSize()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getMetricTrackingObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMetricTrackingObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getMountPath", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMountPath()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getMountableFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMountableFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getMoveAndSplitStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMoveAndSplitStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getNetworkConnectionDetails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkConnectionDetails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getNetworkMountAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMountAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getNetworkMountPath", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMountPath()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getNotificationSubscribers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNotificationSubscribers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getObjectStorageConnectionInformation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObjectStorageConnectionInformation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getObjectsByCredential", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObjectsByCredential(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getOriginalSnapshotName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOriginalSnapshotName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getOriginalVolumeId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOriginalVolumeId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getOriginalVolumeName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOriginalVolumeName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getOriginalVolumeSize", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOriginalVolumeSize()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getOsType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOsType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getOsTypeId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOsTypeId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getParentPartnerships", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetParentPartnerships()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getParentVolume", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetParentVolume()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getPartnerships", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPartnerships()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getPermissionsGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPermissionsGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getProperties", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProperties()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getProvisionedIops", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProvisionedIops()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getRecycleBinFileByIdentifier", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRecycleBinFileByIdentifier(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getRemainingAllowedHosts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRemainingAllowedHosts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getRemainingAllowedHostsForReplicant", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRemainingAllowedHostsForReplicant()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getReplicatingLuns", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReplicatingLuns()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getReplicatingVolume", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReplicatingVolume()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getReplicationEvents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReplicationEvents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getReplicationPartners", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReplicationPartners()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getReplicationSchedule", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReplicationSchedule()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getReplicationStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReplicationStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getReplicationTimestamp", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReplicationTimestamp()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getSchedules", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSchedules()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getServiceResource", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceResource()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getServiceResourceBackendIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceResourceBackendIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getServiceResourceName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceResourceName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getSnapshotCapacityGb", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSnapshotCapacityGb()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getSnapshotCreationTimestamp", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSnapshotCreationTimestamp()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getSnapshotDeletionThresholdPercentage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSnapshotDeletionThresholdPercentage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getSnapshotNotificationStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSnapshotNotificationStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getSnapshotSizeBytes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSnapshotSizeBytes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getSnapshotSpaceAvailable", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSnapshotSpaceAvailable()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getSnapshots", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSnapshots()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getSnapshotsForVolume", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSnapshotsForVolume()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getStaasVersion", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStaasVersion()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getStorageGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getStorageGroupsNetworkConnectionDetails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageGroupsNetworkConnectionDetails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getStorageTierLevel", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageTierLevel()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getStorageType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getTargetIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTargetIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getTotalBytesUsed", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTotalBytesUsed()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getTotalScheduleSnapshotRetentionCount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTotalScheduleSnapshotRetentionCount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getUsageNotification", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUsageNotification()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getValidReplicationTargetDatacenterLocations", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetValidReplicationTargetDatacenterLocations()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getVendorName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVendorName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getVirtualGuest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getVolumeCountLimits", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVolumeCountLimits()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getVolumeDuplicateParameters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVolumeDuplicateParameters()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getVolumeHistory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVolumeHistory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getVolumeStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVolumeStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getWebCCAuthenticationDetails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetWebCCAuthenticationDetails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getWebccAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetWebccAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::getWeeklySchedule", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetWeeklySchedule()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::immediateFailoverToReplicant", func() {
            It("API Call Test", func() {

                _, err := sl_service.ImmediateFailoverToReplicant(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::initiateBareMetalRestore", func() {
            It("API Call Test", func() {

                _, err := sl_service.InitiateBareMetalRestore()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::initiateBareMetalRestoreForServer", func() {
            It("API Call Test", func() {

                _, err := sl_service.InitiateBareMetalRestoreForServer(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::initiateOriginVolumeReclaim", func() {
            It("API Call Test", func() {

                _, err := sl_service.InitiateOriginVolumeReclaim()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::initiateVolumeCutover", func() {
            It("API Call Test", func() {

                _, err := sl_service.InitiateVolumeCutover()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::isBlockingOperationInProgress", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsBlockingOperationInProgress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::isDuplicateReadyForSnapshot", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsDuplicateReadyForSnapshot()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::isDuplicateReadyToMount", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsDuplicateReadyToMount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::isVolumeActive", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsVolumeActive()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::refreshDependentDuplicate", func() {
            It("API Call Test", func() {

                _, err := sl_service.RefreshDependentDuplicate(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::refreshDuplicate", func() {
            It("API Call Test", func() {

                _, err := sl_service.RefreshDuplicate(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::removeAccessFromHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromHardware(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::removeAccessFromHardwareList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromHardwareList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::removeAccessFromHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromHost(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::removeAccessFromHostList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromHostList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::removeAccessFromIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromIpAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::removeAccessFromIpAddressList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromIpAddressList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::removeAccessFromSubnet", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromSubnet(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::removeAccessFromSubnetList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromSubnetList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::removeAccessFromVirtualGuest", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromVirtualGuest(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::removeAccessFromVirtualGuestList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromVirtualGuestList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::removeAccessToReplicantFromHardwareList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToReplicantFromHardwareList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::removeAccessToReplicantFromIpAddressList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToReplicantFromIpAddressList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::removeAccessToReplicantFromSubnet", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToReplicantFromSubnet(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::removeAccessToReplicantFromSubnetList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToReplicantFromSubnetList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::removeAccessToReplicantFromVirtualGuestList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToReplicantFromVirtualGuestList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::removeCredential", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveCredential(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::restoreFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.RestoreFile(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::restoreFromSnapshot", func() {
            It("API Call Test", func() {

                _, err := sl_service.RestoreFromSnapshot(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::sendPasswordReminderEmail", func() {
            It("API Call Test", func() {

                _, err := sl_service.SendPasswordReminderEmail(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::setMountable", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetMountable(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::setSnapshotAllocation", func() {
            It("API Call Test", func() {

				err := sl_service.SetSnapshotAllocation(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::setSnapshotNotification", func() {
            It("API Call Test", func() {

				err := sl_service.SetSnapshotNotification(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::upgradeVolumeCapacity", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpgradeVolumeCapacity(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::uploadFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.UploadFile(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Backup_Evault::validateHostsAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.ValidateHostsAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Storage_DedicatedCluster service", func() {
        var sl_service services.Network_Storage_DedicatedCluster
        BeforeEach(func() {
            sl_service = services.GetNetworkStorageDedicatedClusterService(slsession)
        })


        Context("SoftLayer_Network_Storage_DedicatedCluster::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_DedicatedCluster::getDedicatedClusterList", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDedicatedClusterList()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_DedicatedCluster::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_DedicatedCluster::getServiceResource", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceResource()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Storage_Group service", func() {
        var sl_service services.Network_Storage_Group
        BeforeEach(func() {
            sl_service = services.GetNetworkStorageGroupService(slsession)
        })


        Context("SoftLayer_Network_Storage_Group::addAllowedHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddAllowedHost(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group::attachToVolume", func() {
            It("API Call Test", func() {

                _, err := sl_service.AttachToVolume(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group::getAllowedHosts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedHosts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group::getAttachedVolumes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachedVolumes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group::getGroupType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGroupType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group::getNetworkConnectionDetails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkConnectionDetails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group::getOsType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOsType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group::getServiceResource", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceResource()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group::removeAllowedHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAllowedHost(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group::removeFromVolume", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveFromVolume(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Storage_Group_Iscsi service", func() {
        var sl_service services.Network_Storage_Group_Iscsi
        BeforeEach(func() {
            sl_service = services.GetNetworkStorageGroupIscsiService(slsession)
        })


        Context("SoftLayer_Network_Storage_Group_Iscsi::addAllowedHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddAllowedHost(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Iscsi::attachToVolume", func() {
            It("API Call Test", func() {

                _, err := sl_service.AttachToVolume(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Iscsi::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Iscsi::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Iscsi::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Iscsi::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Iscsi::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Iscsi::getAllowedHosts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedHosts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Iscsi::getAttachedVolumes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachedVolumes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Iscsi::getGroupType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGroupType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Iscsi::getNetworkConnectionDetails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkConnectionDetails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Iscsi::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Iscsi::getOsType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOsType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Iscsi::getServiceResource", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceResource()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Iscsi::removeAllowedHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAllowedHost(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Iscsi::removeFromVolume", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveFromVolume(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Storage_Group_Nfs service", func() {
        var sl_service services.Network_Storage_Group_Nfs
        BeforeEach(func() {
            sl_service = services.GetNetworkStorageGroupNfsService(slsession)
        })


        Context("SoftLayer_Network_Storage_Group_Nfs::addAllowedHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddAllowedHost(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Nfs::attachToVolume", func() {
            It("API Call Test", func() {

                _, err := sl_service.AttachToVolume(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Nfs::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Nfs::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Nfs::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Nfs::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Nfs::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Nfs::getAllowedHosts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedHosts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Nfs::getAttachedVolumes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachedVolumes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Nfs::getGroupType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGroupType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Nfs::getNetworkConnectionDetails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkConnectionDetails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Nfs::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Nfs::getOsType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOsType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Nfs::getServiceResource", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceResource()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Nfs::removeAllowedHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAllowedHost(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Nfs::removeFromVolume", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveFromVolume(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Storage_Group_Type service", func() {
        var sl_service services.Network_Storage_Group_Type
        BeforeEach(func() {
            sl_service = services.GetNetworkStorageGroupTypeService(slsession)
        })


        Context("SoftLayer_Network_Storage_Group_Type::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Group_Type::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Storage_Hub_Cleversafe_Account service", func() {
        var sl_service services.Network_Storage_Hub_Cleversafe_Account
        BeforeEach(func() {
            sl_service = services.GetNetworkStorageHubCleversafeAccountService(slsession)
        })


        Context("SoftLayer_Network_Storage_Hub_Cleversafe_Account::credentialCreate", func() {
            It("API Call Test", func() {

                _, err := sl_service.CredentialCreate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Hub_Cleversafe_Account::credentialDelete", func() {
            It("API Call Test", func() {

                _, err := sl_service.CredentialDelete(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Hub_Cleversafe_Account::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Hub_Cleversafe_Account::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Hub_Cleversafe_Account::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Hub_Cleversafe_Account::getBuckets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBuckets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Hub_Cleversafe_Account::getCancelledBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCancelledBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Hub_Cleversafe_Account::getCapacityUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCapacityUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Hub_Cleversafe_Account::getCloudObjectStorageMetrics", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCloudObjectStorageMetrics(nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Hub_Cleversafe_Account::getCredentialLimit", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCredentialLimit()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Hub_Cleversafe_Account::getCredentials", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCredentials()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Hub_Cleversafe_Account::getEndpoints", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEndpoints(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Hub_Cleversafe_Account::getEndpointsWithRefetch", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEndpointsWithRefetch(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Hub_Cleversafe_Account::getMetricTrackingObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMetricTrackingObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Hub_Cleversafe_Account::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Hub_Cleversafe_Account::getUuid", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUuid()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Storage_Hub_Swift_Metrics service", func() {
        var sl_service services.Network_Storage_Hub_Swift_Metrics
        BeforeEach(func() {
            sl_service = services.GetNetworkStorageHubSwiftMetricsService(slsession)
        })


        Context("SoftLayer_Network_Storage_Hub_Swift_Metrics::getMetricData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMetricData(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Hub_Swift_Metrics::getSummaryData", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSummaryData(nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Storage_Hub_Swift_Share service", func() {
        var sl_service services.Network_Storage_Hub_Swift_Share
        BeforeEach(func() {
            sl_service = services.GetNetworkStorageHubSwiftShareService(slsession)
        })


        Context("SoftLayer_Network_Storage_Hub_Swift_Share::getContainerList", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetContainerList()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Hub_Swift_Share::getFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFile(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Hub_Swift_Share::getFileList", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFileList(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Storage_Iscsi service", func() {
        var sl_service services.Network_Storage_Iscsi
        BeforeEach(func() {
            sl_service = services.GetNetworkStorageIscsiService(slsession)
        })


        Context("SoftLayer_Network_Storage_Iscsi::allowAccessFromHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromHardware(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::allowAccessFromHardwareList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromHardwareList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::allowAccessFromHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromHost(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::allowAccessFromHostList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromHostList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::allowAccessFromIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromIpAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::allowAccessFromIpAddressList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromIpAddressList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::allowAccessFromSubnet", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromSubnet(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::allowAccessFromSubnetList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromSubnetList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::allowAccessFromVirtualGuest", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromVirtualGuest(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::allowAccessFromVirtualGuestList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessFromVirtualGuestList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::allowAccessToReplicantFromHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToReplicantFromHardware(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::allowAccessToReplicantFromHardwareList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToReplicantFromHardwareList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::allowAccessToReplicantFromIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToReplicantFromIpAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::allowAccessToReplicantFromIpAddressList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToReplicantFromIpAddressList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::allowAccessToReplicantFromSubnet", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToReplicantFromSubnet(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::allowAccessToReplicantFromSubnetList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToReplicantFromSubnetList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::allowAccessToReplicantFromVirtualGuest", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToReplicantFromVirtualGuest(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::allowAccessToReplicantFromVirtualGuestList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToReplicantFromVirtualGuestList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::assignCredential", func() {
            It("API Call Test", func() {

                _, err := sl_service.AssignCredential(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::assignNewCredential", func() {
            It("API Call Test", func() {

                _, err := sl_service.AssignNewCredential(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::changePassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.ChangePassword(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::collectBandwidth", func() {
            It("API Call Test", func() {

                _, err := sl_service.CollectBandwidth(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::collectBytesUsed", func() {
            It("API Call Test", func() {

                _, err := sl_service.CollectBytesUsed()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::convertCloneDependentToIndependent", func() {
            It("API Call Test", func() {

                _, err := sl_service.ConvertCloneDependentToIndependent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::createFolder", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateFolder(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::createOrUpdateLunId", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateOrUpdateLunId(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::createSnapshot", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateSnapshot(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::deleteAllFiles", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteAllFiles()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::deleteFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteFile(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::deleteFiles", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteFiles(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::deleteFolder", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteFolder(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::disableSnapshots", func() {
            It("API Call Test", func() {

                _, err := sl_service.DisableSnapshots(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::disasterRecoveryFailoverToReplicant", func() {
            It("API Call Test", func() {

                _, err := sl_service.DisasterRecoveryFailoverToReplicant(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::downloadFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.DownloadFile(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::editCredential", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditCredential(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::enableSnapshots", func() {
            It("API Call Test", func() {

                _, err := sl_service.EnableSnapshots(nil,nil,nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::failbackFromReplicant", func() {
            It("API Call Test", func() {

                _, err := sl_service.FailbackFromReplicant()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::failoverToReplicant", func() {
            It("API Call Test", func() {

                _, err := sl_service.FailoverToReplicant(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getAccountPassword", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccountPassword()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getActiveTransactions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveTransactions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getAllFiles", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllFiles()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getAllFilesByFilter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllFilesByFilter(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getAllowDisasterRecoveryFailback", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowDisasterRecoveryFailback()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getAllowDisasterRecoveryFailover", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowDisasterRecoveryFailover()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getAllowableHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowableHardware(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getAllowableIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowableIpAddresses(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getAllowableSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowableSubnets(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getAllowableVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowableVirtualGuests(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getAllowedHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getAllowedHostsLimit", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedHostsLimit()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getAllowedIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getAllowedReplicationHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedReplicationHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getAllowedReplicationIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedReplicationIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getAllowedReplicationSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedReplicationSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getAllowedReplicationVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedReplicationVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getAllowedSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getAllowedVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getBillingItemCategory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItemCategory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getByUsername", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetByUsername(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getBytesUsed", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBytesUsed()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getCdnUrls", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCdnUrls()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getClusterResource", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetClusterResource()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getCreationScheduleId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCreationScheduleId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getCredentials", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCredentials()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getDailySchedule", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDailySchedule()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getDependentDuplicate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDependentDuplicate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getDependentDuplicates", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDependentDuplicates()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getDuplicateConversionStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDuplicateConversionStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getEvents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEvents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getFailbackNotAllowed", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFailbackNotAllowed()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getFailoverNotAllowed", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFailoverNotAllowed()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getFileBlockEncryptedLocations", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFileBlockEncryptedLocations()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getFileByIdentifier", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFileByIdentifier(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getFileCount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFileCount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getFileList", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFileList(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getFileNetworkMountAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFileNetworkMountAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getFilePendingDeleteCount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFilePendingDeleteCount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getFilesPendingDelete", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFilesPendingDelete()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getFixReplicationCurrentStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFixReplicationCurrentStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getFolderList", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFolderList()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getGraph", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGraph(nil,nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getHasEncryptionAtRest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHasEncryptionAtRest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getHourlySchedule", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHourlySchedule()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getIntervalSchedule", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIntervalSchedule()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getIops", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIops()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getIsConvertToIndependentTransactionInProgress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsConvertToIndependentTransactionInProgress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getIsDependentDuplicateProvisionCompleted", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsDependentDuplicateProvisionCompleted()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getIsInDedicatedServiceResource", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsInDedicatedServiceResource()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getIsMagneticStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsMagneticStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getIsProvisionInProgress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsProvisionInProgress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getIsReadyForSnapshot", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsReadyForSnapshot()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getIsReadyToMount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIsReadyToMount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getIscsiLuns", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIscsiLuns()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getIscsiReplicatingVolume", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIscsiReplicatingVolume()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getIscsiTargetIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIscsiTargetIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getLunId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLunId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getManualSnapshots", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetManualSnapshots()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getMaximumExpansionSize", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMaximumExpansionSize()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getMetricTrackingObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMetricTrackingObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getMountPath", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMountPath()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getMountableFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMountableFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getMoveAndSplitStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMoveAndSplitStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getNetworkConnectionDetails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkConnectionDetails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getNetworkMountAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMountAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getNetworkMountPath", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkMountPath()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getNotificationSubscribers", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNotificationSubscribers()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getObjectStorageConnectionInformation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObjectStorageConnectionInformation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getObjectsByCredential", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObjectsByCredential(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getOriginalSnapshotName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOriginalSnapshotName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getOriginalVolumeId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOriginalVolumeId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getOriginalVolumeName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOriginalVolumeName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getOriginalVolumeSize", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOriginalVolumeSize()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getOsType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOsType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getOsTypeId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetOsTypeId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getParentPartnerships", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetParentPartnerships()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getParentVolume", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetParentVolume()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getPartnerships", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPartnerships()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getPermissionsGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPermissionsGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getProperties", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProperties()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getProvisionedIops", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProvisionedIops()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getRecycleBinFileByIdentifier", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRecycleBinFileByIdentifier(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getRemainingAllowedHosts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRemainingAllowedHosts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getRemainingAllowedHostsForReplicant", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRemainingAllowedHostsForReplicant()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getReplicatingLuns", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReplicatingLuns()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getReplicatingVolume", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReplicatingVolume()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getReplicationEvents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReplicationEvents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getReplicationPartners", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReplicationPartners()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getReplicationSchedule", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReplicationSchedule()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getReplicationStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReplicationStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getReplicationTimestamp", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReplicationTimestamp()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getSchedules", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSchedules()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getServiceResource", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceResource()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getServiceResourceBackendIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceResourceBackendIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getServiceResourceName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceResourceName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getSnapshotCapacityGb", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSnapshotCapacityGb()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getSnapshotCreationTimestamp", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSnapshotCreationTimestamp()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getSnapshotDeletionThresholdPercentage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSnapshotDeletionThresholdPercentage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getSnapshotNotificationStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSnapshotNotificationStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getSnapshotSizeBytes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSnapshotSizeBytes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getSnapshotSpaceAvailable", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSnapshotSpaceAvailable()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getSnapshots", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSnapshots()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getSnapshotsForVolume", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSnapshotsForVolume()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getStaasVersion", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStaasVersion()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getStorageGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getStorageGroupsNetworkConnectionDetails", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageGroupsNetworkConnectionDetails()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getStorageTierLevel", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageTierLevel()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getStorageType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStorageType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getTargetIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTargetIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getTotalBytesUsed", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTotalBytesUsed()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getTotalScheduleSnapshotRetentionCount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTotalScheduleSnapshotRetentionCount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getUsageNotification", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUsageNotification()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getValidReplicationTargetDatacenterLocations", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetValidReplicationTargetDatacenterLocations()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getVendorName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVendorName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getVirtualGuest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getVolumeCountLimits", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVolumeCountLimits()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getVolumeDuplicateParameters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVolumeDuplicateParameters()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getVolumeHistory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVolumeHistory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getVolumeStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVolumeStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getWebccAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetWebccAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::getWeeklySchedule", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetWeeklySchedule()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::immediateFailoverToReplicant", func() {
            It("API Call Test", func() {

                _, err := sl_service.ImmediateFailoverToReplicant(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::initiateOriginVolumeReclaim", func() {
            It("API Call Test", func() {

                _, err := sl_service.InitiateOriginVolumeReclaim()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::initiateVolumeCutover", func() {
            It("API Call Test", func() {

                _, err := sl_service.InitiateVolumeCutover()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::isBlockingOperationInProgress", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsBlockingOperationInProgress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::isDuplicateReadyForSnapshot", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsDuplicateReadyForSnapshot()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::isDuplicateReadyToMount", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsDuplicateReadyToMount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::isVolumeActive", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsVolumeActive()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::refreshDependentDuplicate", func() {
            It("API Call Test", func() {

                _, err := sl_service.RefreshDependentDuplicate(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::refreshDuplicate", func() {
            It("API Call Test", func() {

                _, err := sl_service.RefreshDuplicate(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::removeAccessFromHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromHardware(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::removeAccessFromHardwareList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromHardwareList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::removeAccessFromHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromHost(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::removeAccessFromHostList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromHostList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::removeAccessFromIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromIpAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::removeAccessFromIpAddressList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromIpAddressList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::removeAccessFromSubnet", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromSubnet(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::removeAccessFromSubnetList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromSubnetList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::removeAccessFromVirtualGuest", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromVirtualGuest(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::removeAccessFromVirtualGuestList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessFromVirtualGuestList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::removeAccessToReplicantFromHardwareList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToReplicantFromHardwareList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::removeAccessToReplicantFromIpAddressList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToReplicantFromIpAddressList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::removeAccessToReplicantFromSubnet", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToReplicantFromSubnet(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::removeAccessToReplicantFromSubnetList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToReplicantFromSubnetList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::removeAccessToReplicantFromVirtualGuestList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToReplicantFromVirtualGuestList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::removeCredential", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveCredential(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::restoreFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.RestoreFile(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::restoreFromSnapshot", func() {
            It("API Call Test", func() {

                _, err := sl_service.RestoreFromSnapshot(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::sendPasswordReminderEmail", func() {
            It("API Call Test", func() {

                _, err := sl_service.SendPasswordReminderEmail(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::setMountable", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetMountable(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::setSnapshotAllocation", func() {
            It("API Call Test", func() {

				err := sl_service.SetSnapshotAllocation(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::setSnapshotNotification", func() {
            It("API Call Test", func() {

				err := sl_service.SetSnapshotNotification(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::upgradeVolumeCapacity", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpgradeVolumeCapacity(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::uploadFile", func() {
            It("API Call Test", func() {

                _, err := sl_service.UploadFile(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi::validateHostsAccess", func() {
            It("API Call Test", func() {

                _, err := sl_service.ValidateHostsAccess(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Storage_Iscsi_OS_Type service", func() {
        var sl_service services.Network_Storage_Iscsi_OS_Type
        BeforeEach(func() {
            sl_service = services.GetNetworkStorageIscsiOSTypeService(slsession)
        })


        Context("SoftLayer_Network_Storage_Iscsi_OS_Type::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Iscsi_OS_Type::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Storage_MassDataMigration_CrossRegion_Country_Xref service", func() {
        var sl_service services.Network_Storage_MassDataMigration_CrossRegion_Country_Xref
        BeforeEach(func() {
            sl_service = services.GetNetworkStorageMassDataMigrationCrossRegionCountryXrefService(slsession)
        })


        Context("SoftLayer_Network_Storage_MassDataMigration_CrossRegion_Country_Xref::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_MassDataMigration_CrossRegion_Country_Xref::getCountry", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCountry()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_MassDataMigration_CrossRegion_Country_Xref::getLocationGroup", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLocationGroup()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_MassDataMigration_CrossRegion_Country_Xref::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_MassDataMigration_CrossRegion_Country_Xref::getValidCountriesForRegion", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetValidCountriesForRegion(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Storage_MassDataMigration_Request service", func() {
        var sl_service services.Network_Storage_MassDataMigration_Request
        BeforeEach(func() {
            sl_service = services.GetNetworkStorageMassDataMigrationRequestService(slsession)
        })


        Context("SoftLayer_Network_Storage_MassDataMigration_Request::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_MassDataMigration_Request::getActiveTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_MassDataMigration_Request::getAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_MassDataMigration_Request::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_MassDataMigration_Request::getAllRequestStatuses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllRequestStatuses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_MassDataMigration_Request::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_MassDataMigration_Request::getCreateEmployee", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCreateEmployee()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_MassDataMigration_Request::getCreateUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCreateUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_MassDataMigration_Request::getDeviceConfiguration", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDeviceConfiguration()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_MassDataMigration_Request::getDeviceModel", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDeviceModel()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_MassDataMigration_Request::getKeyContacts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetKeyContacts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_MassDataMigration_Request::getModifyEmployee", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetModifyEmployee()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_MassDataMigration_Request::getModifyUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetModifyUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_MassDataMigration_Request::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_MassDataMigration_Request::getPendingRequests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPendingRequests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_MassDataMigration_Request::getShipments", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetShipments()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_MassDataMigration_Request::getStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_MassDataMigration_Request::getTicket", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTicket()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_MassDataMigration_Request::getTickets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTickets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Storage_MassDataMigration_Request_KeyContact service", func() {
        var sl_service services.Network_Storage_MassDataMigration_Request_KeyContact
        BeforeEach(func() {
            sl_service = services.GetNetworkStorageMassDataMigrationRequestKeyContactService(slsession)
        })


        Context("SoftLayer_Network_Storage_MassDataMigration_Request_KeyContact::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_MassDataMigration_Request_KeyContact::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_MassDataMigration_Request_KeyContact::getRequest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRequest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Storage_MassDataMigration_Request_Status service", func() {
        var sl_service services.Network_Storage_MassDataMigration_Request_Status
        BeforeEach(func() {
            sl_service = services.GetNetworkStorageMassDataMigrationRequestStatusService(slsession)
        })


        Context("SoftLayer_Network_Storage_MassDataMigration_Request_Status::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Storage_Schedule service", func() {
        var sl_service services.Network_Storage_Schedule
        BeforeEach(func() {
            sl_service = services.GetNetworkStorageScheduleService(slsession)
        })


        Context("SoftLayer_Network_Storage_Schedule::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Schedule::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Schedule::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Schedule::getDay", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDay()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Schedule::getDayOfMonth", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDayOfMonth()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Schedule::getDayOfWeek", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDayOfWeek()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Schedule::getEvents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEvents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Schedule::getHour", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHour()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Schedule::getMinute", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMinute()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Schedule::getMonthOfYear", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMonthOfYear()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Schedule::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Schedule::getPartnership", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPartnership()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Schedule::getProperties", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProperties()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Schedule::getReplicaSnapshots", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReplicaSnapshots()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Schedule::getRetentionCount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRetentionCount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Schedule::getSecond", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSecond()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Schedule::getSnapshots", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSnapshots()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Schedule::getType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Schedule::getVolume", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVolume()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Storage_Schedule_Property_Type service", func() {
        var sl_service services.Network_Storage_Schedule_Property_Type
        BeforeEach(func() {
            sl_service = services.GetNetworkStorageSchedulePropertyTypeService(slsession)
        })


        Context("SoftLayer_Network_Storage_Schedule_Property_Type::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Storage_Schedule_Property_Type::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Subnet service", func() {
        var sl_service services.Network_Subnet
        BeforeEach(func() {
            sl_service = services.GetNetworkSubnetService(slsession)
        })


        Context("SoftLayer_Network_Subnet::allowAccessToNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToNetworkStorage(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::allowAccessToNetworkStorageList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToNetworkStorageList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::clearRoute", func() {
            It("API Call Test", func() {

                _, err := sl_service.ClearRoute()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::createReverseDomainRecords", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateReverseDomainRecords()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::createSubnetRouteUpdateTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateSubnetRouteUpdateTransaction(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::createSwipTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateSwipTransaction()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::editNote", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditNote(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::findAllSubnetsAndActiveSwipTransactionStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.FindAllSubnetsAndActiveSwipTransactionStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getActiveRegistration", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveRegistration()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getActiveSwipTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveSwipTransaction()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getActiveTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveTransaction()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getAddressSpace", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAddressSpace()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getAllowedHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedHost()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getAllowedNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getAllowedNetworkStorageReplicas", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedNetworkStorageReplicas()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getAttachedNetworkStorages", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachedNetworkStorages(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getAvailableNetworkStorages", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailableNetworkStorages(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getBoundDescendants", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBoundDescendants()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getBoundRouterFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBoundRouterFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getBoundRouters", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBoundRouters()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getChildren", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetChildren()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getDatacenter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDatacenter()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getDescendants", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDescendants()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getDisplayLabel", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDisplayLabel()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getEndPointIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEndPointIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getGlobalIpRecord", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGlobalIpRecord()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getIpAddressUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIpAddressUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getNetworkComponentFirewall", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkComponentFirewall()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getNetworkProtectionAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkProtectionAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getNetworkTunnelContexts", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkTunnelContexts()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getNetworkVlan", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkVlan()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getPodName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPodName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getProtectedIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProtectedIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getRegionalInternetRegistry", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRegionalInternetRegistry()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getRegistrations", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRegistrations()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getReverseDomain", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReverseDomain()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getReverseDomainRecords", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReverseDomainRecords()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getRoleKeyName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRoleKeyName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getRoleName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRoleName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getRoutableEndpointIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRoutableEndpointIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getRoutingTypeKeyName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRoutingTypeKeyName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getRoutingTypeName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRoutingTypeName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getSubnetForIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubnetForIpAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getSwipTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSwipTransaction()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getTagReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTagReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getUnboundDescendants", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUnboundDescendants()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getUtilizedIpAddressCount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUtilizedIpAddressCount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::getVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::removeAccessToNetworkStorageList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToNetworkStorageList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::route", func() {
            It("API Call Test", func() {

                _, err := sl_service.Route(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet::setTags", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetTags(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Subnet_IpAddress service", func() {
        var sl_service services.Network_Subnet_IpAddress
        BeforeEach(func() {
            sl_service = services.GetNetworkSubnetIpAddressService(slsession)
        })


        Context("SoftLayer_Network_Subnet_IpAddress::allowAccessToNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToNetworkStorage(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::allowAccessToNetworkStorageList", func() {
            It("API Call Test", func() {

                _, err := sl_service.AllowAccessToNetworkStorageList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::editObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::findByIpv4Address", func() {
            It("API Call Test", func() {

                _, err := sl_service.FindByIpv4Address(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::findUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.FindUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getAllowedHost", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedHost()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getAllowedNetworkStorage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedNetworkStorage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getAllowedNetworkStorageReplicas", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllowedNetworkStorageReplicas()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getApplicationDeliveryController", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetApplicationDeliveryController()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getAttachedNetworkStorages", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachedNetworkStorages(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getAvailableNetworkStorages", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAvailableNetworkStorages(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getByIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetByIpAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getContextTunnelTranslations", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetContextTunnelTranslations()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getEndpointSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEndpointSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getGuestNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGuestNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getGuestNetworkComponentBinding", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGuestNetworkComponentBinding()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getPrivateNetworkGateway", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateNetworkGateway()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getProtectionAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProtectionAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getPublicNetworkGateway", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicNetworkGateway()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getRemoteManagementNetworkComponent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRemoteManagementNetworkComponent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getSubnet", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubnet()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getSyslogEventsOneDay", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSyslogEventsOneDay()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getSyslogEventsSevenDays", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSyslogEventsSevenDays()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getTopTenSyslogEventsByDestinationPortOneDay", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTopTenSyslogEventsByDestinationPortOneDay()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getTopTenSyslogEventsByDestinationPortSevenDays", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTopTenSyslogEventsByDestinationPortSevenDays()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getTopTenSyslogEventsByProtocolsOneDay", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTopTenSyslogEventsByProtocolsOneDay()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getTopTenSyslogEventsByProtocolsSevenDays", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTopTenSyslogEventsByProtocolsSevenDays()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getTopTenSyslogEventsBySourceIpOneDay", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTopTenSyslogEventsBySourceIpOneDay()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getTopTenSyslogEventsBySourceIpSevenDays", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTopTenSyslogEventsBySourceIpSevenDays()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getTopTenSyslogEventsBySourcePortOneDay", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTopTenSyslogEventsBySourcePortOneDay()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getTopTenSyslogEventsBySourcePortSevenDays", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTopTenSyslogEventsBySourcePortSevenDays()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getVirtualGuest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::getVirtualLicenses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualLicenses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress::removeAccessToNetworkStorageList", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAccessToNetworkStorageList(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Subnet_IpAddress_Global service", func() {
        var sl_service services.Network_Subnet_IpAddress_Global
        BeforeEach(func() {
            sl_service = services.GetNetworkSubnetIpAddressGlobalService(slsession)
        })


        Context("SoftLayer_Network_Subnet_IpAddress_Global::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress_Global::getActiveTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveTransaction()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress_Global::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress_Global::getDestinationIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDestinationIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress_Global::getIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIpAddress()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress_Global::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress_Global::route", func() {
            It("API Call Test", func() {

                _, err := sl_service.Route(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_IpAddress_Global::unroute", func() {
            It("API Call Test", func() {

                _, err := sl_service.Unroute()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Subnet_Registration service", func() {
        var sl_service services.Network_Subnet_Registration
        BeforeEach(func() {
            sl_service = services.GetNetworkSubnetRegistrationService(slsession)
        })


        Context("SoftLayer_Network_Subnet_Registration::clearRegistration", func() {
            It("API Call Test", func() {

                _, err := sl_service.ClearRegistration()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Registration::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Registration::createObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObjects(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Registration::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Registration::editRegistrationAttachedDetails", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditRegistrationAttachedDetails(nil,nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Registration::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Registration::getDetailReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDetailReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Registration::getEvents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEvents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Registration::getNetworkDetail", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkDetail()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Registration::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Registration::getPersonDetail", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPersonDetail()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Registration::getRegionalInternetRegistry", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRegionalInternetRegistry()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Registration::getRegionalInternetRegistryHandle", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRegionalInternetRegistryHandle()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Registration::getStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Registration::getSubnet", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubnet()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Subnet_Registration_Details service", func() {
        var sl_service services.Network_Subnet_Registration_Details
        BeforeEach(func() {
            sl_service = services.GetNetworkSubnetRegistrationDetailsService(slsession)
        })


        Context("SoftLayer_Network_Subnet_Registration_Details::createObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Registration_Details::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Registration_Details::getDetail", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDetail()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Registration_Details::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Registration_Details::getRegistration", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRegistration()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Subnet_Registration_Status service", func() {
        var sl_service services.Network_Subnet_Registration_Status
        BeforeEach(func() {
            sl_service = services.GetNetworkSubnetRegistrationStatusService(slsession)
        })


        Context("SoftLayer_Network_Subnet_Registration_Status::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Registration_Status::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Subnet_Rwhois_Data service", func() {
        var sl_service services.Network_Subnet_Rwhois_Data
        BeforeEach(func() {
            sl_service = services.GetNetworkSubnetRwhoisDataService(slsession)
        })


        Context("SoftLayer_Network_Subnet_Rwhois_Data::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Rwhois_Data::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Rwhois_Data::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Subnet_Swip_Transaction service", func() {
        var sl_service services.Network_Subnet_Swip_Transaction
        BeforeEach(func() {
            sl_service = services.GetNetworkSubnetSwipTransactionService(slsession)
        })


        Context("SoftLayer_Network_Subnet_Swip_Transaction::findMyTransactions", func() {
            It("API Call Test", func() {

                _, err := sl_service.FindMyTransactions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Swip_Transaction::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Swip_Transaction::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Swip_Transaction::getSubnet", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubnet()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Swip_Transaction::removeAllSubnetSwips", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveAllSubnetSwips()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Swip_Transaction::removeSwipData", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveSwipData()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Swip_Transaction::resendSwipData", func() {
            It("API Call Test", func() {

                _, err := sl_service.ResendSwipData()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Swip_Transaction::swipAllSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.SwipAllSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Subnet_Swip_Transaction::updateAllSubnetSwips", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateAllSubnetSwips()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Tunnel_Module_Context service", func() {
        var sl_service services.Network_Tunnel_Module_Context
        BeforeEach(func() {
            sl_service = services.GetNetworkTunnelModuleContextService(slsession)
        })


        Context("SoftLayer_Network_Tunnel_Module_Context::addCustomerSubnetToNetworkTunnel", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddCustomerSubnetToNetworkTunnel(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::addPrivateSubnetToNetworkTunnel", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddPrivateSubnetToNetworkTunnel(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::addServiceSubnetToNetworkTunnel", func() {
            It("API Call Test", func() {

                _, err := sl_service.AddServiceSubnetToNetworkTunnel(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::applyConfigurationsToDevice", func() {
            It("API Call Test", func() {

                _, err := sl_service.ApplyConfigurationsToDevice()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::createAddressTranslation", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateAddressTranslation(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::createAddressTranslations", func() {
            It("API Call Test", func() {

                _, err := sl_service.CreateAddressTranslations(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::deleteAddressTranslation", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteAddressTranslation(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::downloadAddressTranslationConfigurations", func() {
            It("API Call Test", func() {

                _, err := sl_service.DownloadAddressTranslationConfigurations()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::downloadParameterConfigurations", func() {
            It("API Call Test", func() {

                _, err := sl_service.DownloadParameterConfigurations()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::editAddressTranslation", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditAddressTranslation(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::editAddressTranslations", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditAddressTranslations(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::getActiveTransaction", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetActiveTransaction()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::getAddressTranslationConfigurations", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAddressTranslationConfigurations()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::getAddressTranslations", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAddressTranslations()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::getAllAvailableServiceSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllAvailableServiceSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::getAuthenticationDefault", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAuthenticationDefault()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::getAuthenticationOptions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAuthenticationOptions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::getCustomerSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCustomerSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::getDatacenter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDatacenter()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::getDiffieHellmanGroupDefault", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDiffieHellmanGroupDefault()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::getDiffieHellmanGroupOptions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDiffieHellmanGroupOptions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::getEncryptionDefault", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEncryptionDefault()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::getEncryptionOptions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetEncryptionOptions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::getInternalSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetInternalSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::getKeylifeLimits", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetKeylifeLimits()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::getParameterConfigurationsForCustomerView", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetParameterConfigurationsForCustomerView()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::getPhaseOneKeylifeDefault", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPhaseOneKeylifeDefault()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::getPhaseTwoKeylifeDefault", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPhaseTwoKeylifeDefault()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::getServiceSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetServiceSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::getStaticRouteSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetStaticRouteSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::getTransactionHistory", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTransactionHistory()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::removeCustomerSubnetFromNetworkTunnel", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveCustomerSubnetFromNetworkTunnel(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::removePrivateSubnetFromNetworkTunnel", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemovePrivateSubnetFromNetworkTunnel(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Tunnel_Module_Context::removeServiceSubnetFromNetworkTunnel", func() {
            It("API Call Test", func() {

                _, err := sl_service.RemoveServiceSubnetFromNetworkTunnel(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Vlan service", func() {
        var sl_service services.Network_Vlan
        BeforeEach(func() {
            sl_service = services.GetNetworkVlanService(slsession)
        })


        Context("SoftLayer_Network_Vlan::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getAdditionalPrimarySubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAdditionalPrimarySubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getAttachedNetworkGateway", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachedNetworkGateway()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getAttachedNetworkGatewayFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachedNetworkGatewayFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getAttachedNetworkGatewayVlan", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttachedNetworkGatewayVlan()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getCancelFailureReasons", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetCancelFailureReasons()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getDatacenter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDatacenter()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getDedicatedFirewallFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDedicatedFirewallFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getExtensionRouter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetExtensionRouter()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getFirewallGuestNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirewallGuestNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getFirewallInterfaces", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirewallInterfaces()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getFirewallNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirewallNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getFirewallProtectableIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirewallProtectableIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getFirewallProtectableSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirewallProtectableSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getFirewallRules", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirewallRules()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getGuestNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGuestNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getHardware", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardware()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getHighAvailabilityFirewallFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHighAvailabilityFirewallFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getIpAddressUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetIpAddressUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getLocalDiskStorageCapabilityFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLocalDiskStorageCapabilityFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getNetworkComponentTrunks", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkComponentTrunks()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getNetworkComponents", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkComponents()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getNetworkComponentsTrunkable", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkComponentsTrunkable()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getNetworkSpace", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkSpace()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getNetworkVlanFirewall", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkVlanFirewall()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getPodName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPodName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getPrimaryRouter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimaryRouter()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getPrimarySubnet", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimarySubnet()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getPrimarySubnetVersion6", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimarySubnetVersion6()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getPrimarySubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrimarySubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getPrivateNetworkGateways", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateNetworkGateways()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getPrivateVlan", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateVlan()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getPrivateVlanByIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPrivateVlanByIpAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getProtectedIpAddresses", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProtectedIpAddresses()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getPublicNetworkGateways", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicNetworkGateways()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getPublicVlanByFqdn", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetPublicVlanByFqdn(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getReverseDomainRecords", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetReverseDomainRecords()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getSanStorageCapabilityFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSanStorageCapabilityFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getSecondaryRouter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSecondaryRouter()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getSecondarySubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSecondarySubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getSubnets", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubnets()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getTagReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTagReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getTotalPrimaryIpAddressCount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTotalPrimaryIpAddressCount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getVirtualGuests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVirtualGuests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::getVlanForIpAddress", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetVlanForIpAddress(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::setTags", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetTags(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::updateFirewallIntraVlanCommunication", func() {
            It("API Call Test", func() {

				err := sl_service.UpdateFirewallIntraVlanCommunication(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan::upgrade", func() {
            It("API Call Test", func() {

                _, err := sl_service.Upgrade()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Vlan_Firewall service", func() {
        var sl_service services.Network_Vlan_Firewall
        BeforeEach(func() {
            sl_service = services.GetNetworkVlanFirewallService(slsession)
        })


        Context("SoftLayer_Network_Vlan_Firewall::approveBypassRequest", func() {
            It("API Call Test", func() {

				err := sl_service.ApproveBypassRequest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::getAccountId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccountId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::getBandwidthAllocation", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBandwidthAllocation()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::getBillingCycleBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingCycleBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::getBillingCyclePrivateBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingCyclePrivateBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::getBillingCyclePublicBandwidthUsage", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingCyclePublicBandwidthUsage()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::getBillingItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBillingItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::getBypassRequestStatus", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetBypassRequestStatus()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::getDatacenter", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDatacenter()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::getFirewallFirmwareVersion", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirewallFirmwareVersion()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::getFirewallType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFirewallType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::getFullyQualifiedDomainName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetFullyQualifiedDomainName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::getManagementCredentials", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetManagementCredentials()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::getMetricTrackingObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMetricTrackingObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::getMetricTrackingObjectId", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMetricTrackingObjectId()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::getNetworkFirewallUpdateRequests", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkFirewallUpdateRequests()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::getNetworkGateway", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkGateway()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::getNetworkVlan", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkVlan()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::getNetworkVlans", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetNetworkVlans()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::getRules", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetRules()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::getTagReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTagReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::getUpgradeRequest", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUpgradeRequest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::hasActiveTransactions", func() {
            It("API Call Test", func() {

                _, err := sl_service.HasActiveTransactions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::isAccountAllowed", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsAccountAllowed()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::isHighAvailabilityUpgradeAvailable", func() {
            It("API Call Test", func() {

                _, err := sl_service.IsHighAvailabilityUpgradeAvailable()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::rejectBypassRequest", func() {
            It("API Call Test", func() {

				err := sl_service.RejectBypassRequest()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::restoreDefaults", func() {
            It("API Call Test", func() {

                _, err := sl_service.RestoreDefaults()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::setTags", func() {
            It("API Call Test", func() {

                _, err := sl_service.SetTags(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Network_Vlan_Firewall::updateRouteBypass", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateRouteBypass(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Network_Vlan_Type service", func() {
        var sl_service services.Network_Vlan_Type
        BeforeEach(func() {
            sl_service = services.GetNetworkVlanTypeService(slsession)
        })


        Context("SoftLayer_Network_Vlan_Type::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

})
