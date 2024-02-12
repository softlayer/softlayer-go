package generator

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
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

	apiFilter := map[string][]string{
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
	filteredServices := make([]Type, 0, len(apiFilter))
	// Services can be subclasses of other services. Copy methods from each service's 'Base' entity to
	// the child service, only if a same-named method does not already exist (i.e., overridden by the
	// child service)
	for i, service := range sortedServices {
		sortedServices[i].Methods = GetBaseMethods(service, meta)
		FixReturnType(&sortedServices[i])
	}

	// Purge anything not in the apiFilter
	for _, service := range sortedServices {
		if _, ok := apiFilter[service.Name]; ok {
			fmt.Printf("%v is in apiFilter!\n", service.Name)
			filteredMethods := map[string]Method{}
			for mi, method := range service.Methods {
				for x := range apiFilter[service.Name] {
					if mi == apiFilter[service.Name][x] {
						fmt.Printf("%v::%v is in apiFilter\n", service.Name, method.Name)
						filteredMethods[mi] = method
						break
					}					
				}

			}
			service.Methods = filteredMethods
			filteredServices = append(filteredServices, service)
		}
	}

	// fmt.Printf("Creating Datatypes...\n")
	// err = WritePackage(outputDir, "datatypes", sortedTypes, datatype)
	// errCheck(err, "Error creating Datatypes")

	fmt.Printf("Creating Servicess...\n")
	err = WritePackage(outputDir, "services", filteredServices, services)
	errCheck(err, "Error creating Services")
}
