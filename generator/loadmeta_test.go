package generator_test

import (
	"encoding/json"
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/softlayer/softlayer-go/generator"
)

var _ = Describe("LoadMeta Tests", func() {
	BeforeEach(func() {

	})
	Context("Testing Load Meta methods", func() {
		Context("RemovePrefix() Tests", func() {
			It("Test SoftLayer_ is removed", func() {
				result := RemovePrefix("SoftLayer_Test")
				Expect(result).To(Equal("Test"))
			})
			It("Unexpected Prefix not removed", func() {
				result := RemovePrefix("Something_Else")
				Expect(result).To(Equal("Something_Else"))
			})
		})
		// https://onsi.github.io/ginkgo/#table-specs
		Context("ConvertType() Tests", func() {
			DescribeTable("Testing converting SL types to Go types",
				func(i_string string, expected string) {
					result := ConvertType(i_string, "datatypes")
					Expect(result).To(Equal(expected))
				},
				// This handles printing out the Entry description.
				func(i_string string, expected string) string {
					return fmt.Sprintf("ConvertType(%s) = %s", i_string, expected)
				},
				Entry(nil, "unsignedLong", "uint"),
				Entry(nil, "boolean", "bool"),
				Entry(nil, "dateTime", "Time"),
				Entry(nil, "decimal", "Float64"),
				Entry(nil, "base64Binary", "[]byte"),
				Entry(nil, "json", "string"),
				Entry(nil, "SoftLayer_Hardware", "Hardware"),
				Entry(nil, "SoftLayer_Entity", "Entity"),
				Entry(nil, "McAfee_Something", "McAfee_Something"),
				Entry(nil, "test", "test"),
			)
			Describe("Testing converting to an interface", func() {
				result := ConvertType("SoftLayer_Entity", "datatypes", "SoftLayer_Test", "resource")
				Expect(result).To(Equal("interface{}"))
			})
		})
		Context("RemoveReservedWords() Tests", func() {
			DescribeTable("Testing reserveed word removal",
				func(i_string string, expected string) {
					Expect(RemoveReservedWords(i_string)).To(Equal(expected))
				},
				func(i_string string, expected string) string {
					return fmt.Sprintf("RemovedReservedWords(%s) = %s", i_string, expected)
				},
				Entry(nil, "type", "typ"),
				Entry(nil, "other", "other"),
				Entry(nil, "SoftLayer", "SoftLayer"),
			)
		})
		Context("Desnake() Tests", func() {
			Describe("Test Desnake", func() {
				It("Desnake(SoftLayer_Test_This_Thing)", func() {
					Expect(Desnake("SoftLayer_Test_This_Thing")).To(Equal("SoftLayerTestThisThing"))
				})

			})
		})
		Context("GoDoc() Tests", func() {
			Describe("GoDoc Tests", func() {
				It("GoDoc('')", func() {
					Expect(GoDoc("")).To(Equal("// no documentation yet"))
				})
				It("GoDoc('TEST')", func() {
					Expect(GoDoc("TEST")).To(Equal("// TEST"))
				})
			})
		})
		Context("Tags() Tests", func() {
			Describe("Tags Test", func() {
				It("Tags('resourceRecords')", func() {
					Expect(Tags("resourceRecords")).To(Equal("resourceRecords"))
				})
				It("Tags('SomethingElse')", func() {
					Expect(Tags("SomethingElse")).To(Equal("SomethingElse,omitempty"))
				})
			})
		})
		Context("Data Fixing Tests", func() {
			var sl_services map[string]Type
			BeforeEach(func() {
				jsonResp, err := GetMetaFromFile(metaPath) // from generate_test.go
				Expect(err).NotTo(HaveOccurred())
				err = json.Unmarshal(jsonResp, &sl_services)
				Expect(err).NotTo(HaveOccurred())
			})
			Describe("CreateGetters() Basic Tests", func() {
				It("No Properties", func() {
					test_service := sl_services["SoftLayer_Test_Service"]
					CreateGetters(&test_service)
					Expect(test_service.Name).To(Equal("SoftLayer_Test_Service"))
					Expect(test_service.Methods).To(BeNil())
				})
				It("Happy Path", func() {
					test_service := sl_services["SoftLayer_Account"]
					CreateGetters(&test_service)
					Expect(test_service.Name).To(Equal("SoftLayer_Account"))
					Expect(test_service.Methods).ToNot(BeNil())
					Expect(len(test_service.Methods)).To(Equal(4))
					Expect(test_service.Methods["getAbuseEmail"].Name).To(Equal("getAbuseEmail"))
					Expect(test_service.Methods["getAccountLinks"].Name).To(Equal("getAccountLinks"))
				})
			})
			Describe("AddComplexType() Basic Tests", func() {
				It("SoftLayer_Container_Product_Order Complex Type", func() {
					test_datatype := sl_services["SoftLayer_Container_Product_Order"]
					AddComplexType(&test_datatype)
					Expect(len(test_datatype.Properties)).To(Equal(2))
					Expect(test_datatype.Properties["complexType"]).ToNot(BeNil())
					Expect(test_datatype.Properties["complexType"].Name).To(Equal("complexType"))
				})
				It("SoftLayer_Container_User_Customer_External_Binding Complex Type", func() {
					test_datatype := sl_services["SoftLayer_Container_User_Customer_External_Binding"]
					AddComplexType(&test_datatype)
					Expect(len(test_datatype.Properties)).To(Equal(2))
					Expect(test_datatype.Properties["complexType"]).ToNot(BeNil())
					Expect(test_datatype.Properties["complexType"].Name).To(Equal("complexType"))
				})
			})
			Describe("FixDatatype() Basic Tests", func() {
				It("Test SoftLayer_Dns_Domain_ResourceRecord Fixes", func() {
					test_datatype1 := sl_services["SoftLayer_Dns_Domain_ResourceRecord_Test"]
					test_datatype2 := sl_services["SoftLayer_Dns_Domain_ResourceRecord"]
					Expect(len(test_datatype1.Properties)).To(Equal(1))
					Expect(len(test_datatype2.Properties)).To(Equal(1))
					FixDatatype(&test_datatype1, sl_services)
					Expect(len(test_datatype2.Properties)).To(Equal(2))
					Expect(test_datatype2.Properties["testProp"]).ToNot(BeNil())
					Expect(test_datatype2.Properties["testProp"].Name).To(Equal("testProp1"))
				})
				It("Test Container_User_Customer_External_Binding Fixes", func() {
					test_datatype1 := sl_services["SoftLayer_Container_User_Customer_External_Binding_Verisign"]
					test_datatype2 := sl_services["SoftLayer_Container_User_Customer_External_Binding"]
					Expect(len(test_datatype1.Properties)).To(Equal(1))
					Expect(len(test_datatype2.Properties)).To(Equal(1))
					FixDatatype(&test_datatype1, sl_services)
					Expect(len(test_datatype2.Properties)).To(Equal(2))
					Expect(test_datatype2.Properties["securityCode"]).ToNot(BeNil())
					Expect(test_datatype2.Properties["securityCode"].Name).To(Equal("securityCode"))
				})
			})
		})
	})
})
