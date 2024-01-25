
package services_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/softlayer/softlayer-go/session/sessionfakes"
    "github.com/softlayer/softlayer-go/services"
)   

var _ = Describe("Configuration Tests", func() {
    var slsession *sessionfakes.FakeSLSession
    BeforeEach(func() {
        slsession = &sessionfakes.FakeSLSession{}
    })

    Context("Testing SoftLayer_Configuration_Storage_Group_Array_Type service", func() {
        var sl_service services.Configuration_Storage_Group_Array_Type
        BeforeEach(func() {
            sl_service = services.GetConfigurationStorageGroupArrayTypeService(slsession)
        })


        Context("SoftLayer_Configuration_Storage_Group_Array_Type::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Storage_Group_Array_Type::getHardwareComponentModels", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetHardwareComponentModels()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Storage_Group_Array_Type::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Configuration_Template service", func() {
        var sl_service services.Configuration_Template
        BeforeEach(func() {
            sl_service = services.GetConfigurationTemplateService(slsession)
        })


        Context("SoftLayer_Configuration_Template::copyTemplate", func() {
            It("API Call Test", func() {

                _, err := sl_service.CopyTemplate(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template::deleteObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.DeleteObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template::editObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.EditObject(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template::getAccount", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAccount()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template::getAllObjects", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllObjects()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template::getConfigurationSections", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetConfigurationSections()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template::getDefaultValues", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDefaultValues()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template::getDefinitions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDefinitions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template::getItem", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetItem()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template::getLinkedSectionReferences", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLinkedSectionReferences()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template::getParent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetParent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template::getUser", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetUser()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template::updateDefaultValues", func() {
            It("API Call Test", func() {

                _, err := sl_service.UpdateDefaultValues(nil)

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Configuration_Template_Section service", func() {
        var sl_service services.Configuration_Template_Section
        BeforeEach(func() {
            sl_service = services.GetConfigurationTemplateSectionService(slsession)
        })


        Context("SoftLayer_Configuration_Template_Section::getDefinitions", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDefinitions()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template_Section::getDisallowedDeletionFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDisallowedDeletionFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template_Section::getLinkedTemplate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLinkedTemplate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template_Section::getLinkedTemplateReference", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetLinkedTemplateReference()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template_Section::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template_Section::getProfiles", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetProfiles()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template_Section::getSectionType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSectionType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template_Section::getSectionTypeName", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSectionTypeName()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template_Section::getSubSections", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSubSections()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template_Section::getTemplate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTemplate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template_Section::hasSubSections", func() {
            It("API Call Test", func() {

                _, err := sl_service.HasSubSections()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Configuration_Template_Section_Definition service", func() {
        var sl_service services.Configuration_Template_Section_Definition
        BeforeEach(func() {
            sl_service = services.GetConfigurationTemplateSectionDefinitionService(slsession)
        })


        Context("SoftLayer_Configuration_Template_Section_Definition::getAttributes", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAttributes()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template_Section_Definition::getDefaultValue", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDefaultValue()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template_Section_Definition::getGroup", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetGroup()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template_Section_Definition::getMonitoringDataFlag", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetMonitoringDataFlag()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template_Section_Definition::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template_Section_Definition::getSection", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSection()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template_Section_Definition::getValueType", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetValueType()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Configuration_Template_Section_Definition_Group service", func() {
        var sl_service services.Configuration_Template_Section_Definition_Group
        BeforeEach(func() {
            sl_service = services.GetConfigurationTemplateSectionDefinitionGroupService(slsession)
        })


        Context("SoftLayer_Configuration_Template_Section_Definition_Group::getAllGroups", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetAllGroups()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template_Section_Definition_Group::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template_Section_Definition_Group::getParent", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetParent()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Configuration_Template_Section_Definition_Type service", func() {
        var sl_service services.Configuration_Template_Section_Definition_Type
        BeforeEach(func() {
            sl_service = services.GetConfigurationTemplateSectionDefinitionTypeService(slsession)
        })


        Context("SoftLayer_Configuration_Template_Section_Definition_Type::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Configuration_Template_Section_Definition_Value service", func() {
        var sl_service services.Configuration_Template_Section_Definition_Value
        BeforeEach(func() {
            sl_service = services.GetConfigurationTemplateSectionDefinitionValueService(slsession)
        })


        Context("SoftLayer_Configuration_Template_Section_Definition_Value::getDefinition", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetDefinition()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template_Section_Definition_Value::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template_Section_Definition_Value::getTemplate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTemplate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Configuration_Template_Section_Profile service", func() {
        var sl_service services.Configuration_Template_Section_Profile
        BeforeEach(func() {
            sl_service = services.GetConfigurationTemplateSectionProfileService(slsession)
        })


        Context("SoftLayer_Configuration_Template_Section_Profile::getConfigurationSection", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetConfigurationSection()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template_Section_Profile::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Configuration_Template_Section_Reference service", func() {
        var sl_service services.Configuration_Template_Section_Reference
        BeforeEach(func() {
            sl_service = services.GetConfigurationTemplateSectionReferenceService(slsession)
        })


        Context("SoftLayer_Configuration_Template_Section_Reference::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template_Section_Reference::getSection", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetSection()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

        Context("SoftLayer_Configuration_Template_Section_Reference::getTemplate", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetTemplate()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Configuration_Template_Section_Type service", func() {
        var sl_service services.Configuration_Template_Section_Type
        BeforeEach(func() {
            sl_service = services.GetConfigurationTemplateSectionTypeService(slsession)
        })


        Context("SoftLayer_Configuration_Template_Section_Type::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

    Context("Testing SoftLayer_Configuration_Template_Type service", func() {
        var sl_service services.Configuration_Template_Type
        BeforeEach(func() {
            sl_service = services.GetConfigurationTemplateTypeService(slsession)
        })


        Context("SoftLayer_Configuration_Template_Type::getObject", func() {
            It("API Call Test", func() {

                _, err := sl_service.GetObject()

                Expect(err).To(Succeed())
                Expect(slsession.DoRequestCallCount()).To(Equal(1))
            })
        })

    })

})
