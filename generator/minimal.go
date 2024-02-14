package generator

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"os"
	"bufio"
)

func init() {
	minCmd.Flags().StringVarP(
		&varMetadata,
		"metadata", "m",
		SoftLayerMetadataAPIURL,
		"A JSON file or HTTPS url to read metadata from.",
	)
	minCmd.Flags().StringVarP(
		&varOutput,
		"output", "o",
		"./",
		"Output directory. A 'services' and 'datatypes' directory will be created here.",
	)
	rootCmd.AddCommand(minCmd)
}

// These are defined in the generate.go file
// var varMetadata string
// var varOutput string
var minCmd = &cobra.Command{
	Use:   "minimal",
	Short: "Generates the MINIMAL service and datatype definitions, used mostly for the softlayer-cli",
	Run: func(cmd *cobra.Command, args []string) {
		varMetadata = strings.ToLower(varMetadata)
		GenerateMinAPI(varMetadata, varOutput)
	},
}

func BuildApiFilter(filterPath string) map[string][]string {
	filteredApi :=  map[string][]string{
		"SoftLayer_Account": {
			"getObject",
			"getVirtualGuests",
			"getHardware",
			"getCurrentUser",
			"getDomains",
			"setVlanSpan",
		},
		"SoftLayer_Virtual_Guest": {
			"getObject",
			"deleteObject",
			"editObject",
		},
		"SoftLayer_Hardware_Server": {
			"getObject",
			"deleteObject",
			"rebootHard",
		},
	}

	file, err := os.Open(filterPath)
	errCheck(err, "Problem opening filter file")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		this_line := scanner.Text()
		apicall := strings.Split(this_line, "::")
		if len(apicall) != 2 {
			fmt.Printf("%v badly formatted, skipping\n", this_line)
		}
		t_service := apicall[0]
		t_method, _ := strings.CutSuffix(apicall[1], ".json")
		if _, ok := filteredApi[t_service]; !ok {
			filteredApi[t_service] = []string{}
		}
		filteredApi[t_service] = append(filteredApi[t_service], t_method)
	}
	errCheck(scanner.Err(), "Scanner Error")
	return filteredApi
}

func GenerateMinAPI(metaPath string, outputDir string) {
	fmt.Printf("Getting metadata from %s ...\n", metaPath)
	var jsonResp []byte
	var err error
	if strings.HasPrefix(metaPath, "http") {
		jsonResp, err = GetMetaFromURL(metaPath)
		errCheck(err, fmt.Sprintf("Error getting metadata from %s", metaPath))
	} else {
		jsonResp, err = GetMetaFromFile(metaPath)
		errCheck(err, fmt.Sprintf("Error reading file %s", metaPath))
	}
	var meta map[string]Type
	err = json.Unmarshal(jsonResp, &meta)
	errCheck(err, "Error unmarshaling json response")

	// Build an array of Types, sorted by name
	// This will ensure consistency in the order that code is later emitted
	keys := GetSortedKeys(meta)

	sortedTypes := make([]Type, 0, len(keys))
	sortedServices := make([]Type, 0, len(keys))

	for _, name := range keys {
		t := meta[name]
		sortedTypes = append(sortedTypes, t)
		AddComplexType(&t)
		FixDatatype(&t, meta)

		// Not every datatype is also a service
		if !t.NoService {
			CreateGetters(&t)
			sortedServices = append(sortedServices, t)
		}
	}

	apiFilter := BuildApiFilter("min_filter.txt")

	filteredServices := make([]Type, 0, len(apiFilter))
	// Services can be subclasses of other services. Copy methods from each service's 'Base' entity to
	// the child service, only if a same-named method does not already exist (i.e., overridden by the
	// child service)
	for i, service := range sortedServices {
		sortedServices[i].Methods = GetBaseMethods(service, meta)
		FixReturnType(&sortedServices[i])
	}

	// Purge anything not in the apiFilter
	// Go through all the sortedServices
	for _, service := range sortedServices {
		// if that service is in the apiFIlter
		if _, ok := apiFilter[service.Name]; ok {
			// Create an empty method list
			filteredMethods := map[string]Method{}
			// go through all the official methods
			for mi, method := range service.Methods {
				// see if the method is in the apiFilter
				for x := range apiFilter[service.Name] {
					// Add it to the filteredMethods slice
					if mi == apiFilter[service.Name][x] {
						filteredMethods[mi] = method
						break
					}					
				}
			}
			service.Methods = filteredMethods
			filteredServices = append(filteredServices, service)
		}
	}

	/*
	// Find all the datatypes we actually need...
	requiredDatatypes := map[string]Type{
		"SoftLayer_Entity": Type{Name: "SoftLayer_Entity"},
		"SoftLayer_User_Interface": Type{Name: "SoftLayer_User_Interface", Base: "Entity"},
		"SoftLayer_Software_Component": meta["SoftLayer_Software_Component"],
		"SoftLayer_Ticket_Attachment": meta["SoftLayer_Ticket_Attachment"],
		"SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type": meta["SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type"],
		"SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Type": meta["SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Type"],
		"SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Method": meta["SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Method"],
		"SoftLayer_Virtual_Network_SecurityGroup_NetworkComponentBinding": meta["SoftLayer_Virtual_Network_SecurityGroup_NetworkComponentBinding"],
		"SoftLayer_Network_Service_Resource": meta["SoftLayer_Network_Service_Resource"],
		"SoftLayer_Network_Storage_Type": meta["SoftLayer_Network_Storage_Type"],
		"SoftLayer_Order_Item": meta["SoftLayer_Order_Item"],
		"SoftLayer_Billing_Order_Item": meta["SoftLayer_Billing_Order_Item"],
		"SoftLayer_Hardware_Status": meta["SoftLayer_Hardware_Status"],
		"SoftLayer_Software_License": meta["SoftLayer_Software_License"],
		"SoftLayer_Software_Description": meta["SoftLayer_Software_Description"],
		"SoftLayer_Virtual_Guest_Network_Component": meta["SoftLayer_Virtual_Guest_Network_Component"],
		"SoftLayer_Location_Datacenter": meta["SoftLayer_Location_Datacenter"],
		"SoftLayer_Network_Firewall_Module_Context_Interface": meta["SoftLayer_Network_Firewall_Module_Context_Interface"],
		"SoftLayer_Hardware_Router": meta["SoftLayer_Hardware_Router"],
		"SoftLayer_Hardware_Switch": meta["SoftLayer_Hardware_Switch"],
		"SoftLayer_Container_Product_Order_Virtual_DedicatedHost": meta["SoftLayer_Container_Product_Order_Virtual_DedicatedHost"],
		"SoftLayer_Container_Product_Order_Virtual_ReservedCapacity": meta["SoftLayer_Container_Product_Order_Virtual_ReservedCapacity"],
		"SoftLayer_Network_Firewall_Update_Request_Rule": meta["SoftLayer_Network_Firewall_Update_Request_Rule"],
		"SoftLayer_Product_Item_Category": meta["SoftLayer_Product_Item_Category"],
		"SoftLayer_Container_Product_Order_Network_Storage_AsAService": meta["SoftLayer_Container_Product_Order_Network_Storage_AsAService"],
		"SoftLayer_Container_Network_CdnMarketplace_Configuration_Performance_DynamicContentAcceleration": meta["SoftLayer_Container_Network_CdnMarketplace_Configuration_Performance_DynamicContentAcceleration"],
		"SoftLayer_Location_Region_Location": meta["SoftLayer_Location_Region_Location"],
		"SoftLayer_Location_Group_Pricing": meta["SoftLayer_Location_Group_Pricing"],
		"SoftLayer_Container_Product_Order_Network_Protection_Firewall_Dedicated": meta["SoftLayer_Container_Product_Order_Network_Protection_Firewall_Dedicated"],
		"SoftLayer_Container_Product_Order_Network_Protection_Firewall": meta["SoftLayer_Container_Product_Order_Network_Protection_Firewall"],
		"SoftLayer_Network_Component_Group": meta["SoftLayer_Network_Component_Group"],
		"SoftLayer_Network_Firewall_AccessControlList":  meta["SoftLayer_Network_Firewall_AccessControlList"],
		"SoftLayer_Container_Product_Order_SshKeys": meta["SoftLayer_Container_Product_Order_SshKeys"],
		"SoftLayer_Network_Service_Resource_Attribute": meta["SoftLayer_Network_Service_Resource_Attribute"],
		"SoftLayer_Product_Item_Attribute": meta["SoftLayer_Product_Item_Attribute"],
		"SoftLayer_Network_Customer_Subnet_IpAddress": meta["SoftLayer_Network_Customer_Subnet_IpAddress"],
		"SoftLayer_Container_Product_Order_Network_Tunnel_Ipsec": meta["SoftLayer_Container_Product_Order_Network_Tunnel_Ipsec"],
		"SoftLayer_Container_Product_Order_Software_License": meta["SoftLayer_Container_Product_Order_Software_License"],
		"SoftLayer_Container_Product_Order_Network_LoadBalancer_AsAService": meta["SoftLayer_Container_Product_Order_Network_LoadBalancer_AsAService"],
		"SoftLayer_Container_Product_Order_Network_Vlan": meta["SoftLayer_Container_Product_Order_Network_Vlan"],
		"SoftLayer_Container_Product_Order_Network_Subnet": meta["SoftLayer_Container_Product_Order_Network_Subnet"],
		"SoftLayer_Billing_Item_Network_Subnet_IpAddress_Global": meta["SoftLayer_Billing_Item_Network_Subnet_IpAddress_Global"],
		"SoftLayer_Container_Product_Order_Network_Storage_AsAService_Upgrade": meta["SoftLayer_Container_Product_Order_Network_Storage_AsAService_Upgrade"],
		"SoftLayer_Container_Product_Order_Network_Storage_Enterprise_SnapshotSpace_Upgrade": meta["SoftLayer_Container_Product_Order_Network_Storage_Enterprise_SnapshotSpace_Upgrade"],
		"SoftLayer_Container_Product_Order_Network_Storage_Enterprise_SnapshotSpace": meta["SoftLayer_Container_Product_Order_Network_Storage_Enterprise_SnapshotSpace"],
		"SoftLayer_Network_Storage_Schedule": meta["SoftLayer_Network_Storage_Schedule"],
		"SoftLayer_Virtual_Guest_SupplementalCreateObjectOptions": meta["SoftLayer_Virtual_Guest_SupplementalCreateObjectOptions"],
		"SoftLayer_Billing_Item_Network_Subnet": meta["SoftLayer_Billing_Item_Network_Subnet"],
		"SoftLayer_Network_Storage_Schedule_Type": meta["SoftLayer_Network_Storage_Schedule_Type"],
		"SoftLayer_Virtual_Guest_Attribute": meta["SoftLayer_Virtual_Guest_Attribute"],
		"SoftLayer_Container_Product_Order_Property": meta["SoftLayer_Container_Product_Order_Property"],
		"SoftLayer_Billing_Item_Virtual_Guest": meta["SoftLayer_Billing_Item_Virtual_Guest"],
		"SoftLayer_Virtual_Guest_Power_State": meta["SoftLayer_Virtual_Guest_Power_State"],
		"SoftLayer_Container_Product_Order_Virtual_Guest_Upgrade": meta["SoftLayer_Container_Product_Order_Virtual_Guest_Upgrade"],
		"SoftLayer_Container_Product_Order_Virtual_Guest": meta["SoftLayer_Container_Product_Order_Virtual_Guest"],
		"SoftLayer_Container_Product_Order_Hardware_Server": meta["SoftLayer_Container_Product_Order_Hardware_Server"],
		"SoftLayer_Network_Bandwidth_Version1_Allocation": meta["SoftLayer_Network_Bandwidth_Version1_Allocation"],
		"SoftLayer_Network_Gateway_Member": meta["SoftLayer_Network_Gateway_Member"],
		"SoftLayer_Network_Gateway_Vlan": meta["SoftLayer_Network_Gateway_Vlan"],
		"SoftLayer_Network_Gateway_Status": meta["SoftLayer_Network_Gateway_Status"],
		"SoftLayer_Hardware_Router_Backend": meta["SoftLayer_Hardware_Router_Backend"],
		"SoftLayer_Billing_Item_Virtual_DedicatedHost": meta["SoftLayer_Billing_Item_Virtual_DedicatedHost"],
		"SoftLayer_Network_LBaaS_Listener": meta["SoftLayer_Network_LBaaS_Listener"],
		"SoftLayer_Network_LBaaS_Pool": meta["SoftLayer_Network_LBaaS_Pool"],
		"SoftLayer_Network_LBaaS_Member": meta["SoftLayer_Network_LBaaS_Member"],
		"SoftLayer_Network_LBaaS_HealthMonitor": meta["SoftLayer_Network_LBaaS_HealthMonitor"],
		"SoftLayer_Product_Item_Resource_Conflict": meta["SoftLayer_Product_Item_Resource_Conflict"],
		"SoftLayer_Network_Service_Resource_Type": meta["SoftLayer_Network_Service_Resource_Type"],
		"SoftLayer_Ticket_Status": meta["SoftLayer_Ticket_Status"],
		"SoftLayer_Network_Component_Network_Vlan_Trunk": meta["SoftLayer_Network_Component_Network_Vlan_Trunk"],
		"SoftLayer_User_Customer_Status": meta["SoftLayer_User_Customer_Status"],
		"SoftLayer_Container_Search_Result": meta["SoftLayer_Container_Search_Result"],

	}
	for _, datatype := range sortedTypes {
		if strings.HasPrefix(datatype.Name, "SoftLayer_Container_") {
			requiredDatatypes[datatype.Name] = datatype
		}
	}
	// Go through all the filtered Services
	for _, service := range filteredServices {
		fmt.Printf("Checking on %v\n", service.Name)
		// And each of their methods
		for _, method := range service.Methods {
			// if we don't already have this datatype
			fmt.Printf("\tChecking on %v\n", method.Name)
			if _, ok := requiredDatatypes[method.Type]; !ok && strings.HasPrefix(method.Type, "SoftLayer_"){
				fmt.Printf("\t\t%v => %v needs to be added\n", method.Name, method.Type)
				// find the type in the sorted list
				for _, sl_type := range sortedTypes {
					// add it to the requiredDatatypes
					if sl_type.Name == method.Type  {
						fmt.Printf("\t\t\tAdding %v\n", sl_type.Name)
						requiredDatatypes[method.Type] = sl_type
						break
					}
				}
			} else {
				this := requiredDatatypes[method.Type]
				fmt.Printf("\t%v exists in requiredDatatypes[%v] already\n", this.Name, method.Type)
			}
			for _, param := range method.Parameters {
				if _, ok := requiredDatatypes[param.Type]; !ok && strings.HasPrefix(param.Type, "SoftLayer_"){
					fmt.Printf("\t\t%v => %v needs to be added\n", param.Name, param.Type)
					// find the type in the sorted list
					for _, sl_type := range sortedTypes {
						// add it to the requiredDatatypes
						if sl_type.Name == param.Type {
							fmt.Printf("\t\t\tAdding %v\n", sl_type.Name)
							requiredDatatypes[param.Type] = sl_type
							break
						}
					}
				} else {
					this := requiredDatatypes[param.Type]
					fmt.Printf("\t%v exists in requiredDatatypes[%v] already\n", this.Name, param.Type)
				}
			}
		}
	}


	// Remove any properties that don't have one of our required datatypes and hope for the best.
	
	for x, datatype := range requiredDatatypes {
		requiredProperties := map[string]Property{}
		for _, property := range datatype.Properties {
			if _, ok := requiredDatatypes[property.Type]; ok {
				requiredProperties[property.Name] = property
			}  else if  !strings.HasPrefix(property.Type, "SoftLayer_")  {
				requiredProperties[property.Name] = property
			} else {
				fmt.Printf("DT: %v->%v (%v) not needed\n", datatype.Name, property.Name, property.Type)
			}
		}
		datatype.Properties = requiredProperties
		requiredDatatypes[x] = datatype
	}
	
	// resort types.
	keys = GetSortedKeys(requiredDatatypes)
	sortedTypes = make([]Type, 0, len(keys))
	for _, name := range keys {
		t := requiredDatatypes[name]
		sortedTypes = append(sortedTypes, t)
	}

	*/
	fmt.Printf("Creating Datatypes...\n")
	err = WritePackage(outputDir, "datatypes", sortedTypes, datatype)
	errCheck(err, "Error creating Datatypes")

	fmt.Printf("Creating Servicess...\n")
	err = WritePackage(outputDir, "services", filteredServices, services)
	errCheck(err, "Error creating Services")
}
