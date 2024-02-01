package generator

import (
	"fmt"
	"strings"
	"os"
	"encoding/json"
	"github.com/spf13/cobra"
)

func init() {
	generateCmd.Flags().StringVarP(
		&varMetadata,
		"metadata", "m",
		SoftLayerMetadataAPIURL,
		"A JSON file or HTTPS url to read metadata from.",
	)
	generateCmd.Flags().StringVarP(
		&varOutput,
		"output", "o",
		"./",
		"Output directory. A 'services' and 'datatypes' directory will be created here.",
	)
	rootCmd.AddCommand(generateCmd)
}

const SoftLayerMetadataAPIURL = "https://api.softlayer.com/metadata/v3.1"
var varMetadata string
var varOutput string
var generateCmd = &cobra.Command{
	Use: "generate",
	Short: "Generates the service and datatype definitions",
	Run: func(cmd *cobra.Command, args []string) {
		varMetadata = strings.ToLower(varMetadata)
		GenerateAPI(varMetadata, varOutput)
	},
}


func GenerateAPI(metaPath string, outputDir string) {
	fmt.Printf("Getting metadata from %s ...\n", metaPath)
	var jsonResp []byte
	var err error
	if strings.HasPrefix(metaPath, "http") {
		jsonResp, err = GetMetaFromURL(metaPath)
		errCheck(err, fmt.Sprintf("Error getting metadata from %s", metaPath))
	} else {
		fmt.Printf("Going to read from a file |%s|!\n", metaPath)
	}
	var meta map[string]Type
	err = json.Unmarshal(jsonResp, &meta)
	errCheck(err, "Error unmarshaling json response")

	// Build an array of Types, sorted by name
	// This will ensure consistency in the order that code is later emitted
	keys := getSortedKeys(meta)

	sortedTypes := make([]Type, 0, len(keys))
	sortedServices := make([]Type, 0, len(keys))

	for _, name := range keys {
		t := meta[name]
		sortedTypes = append(sortedTypes, t)
		addComplexType(&t)
		fixDatatype(&t, meta)

		// Not every datatype is also a service
		if !t.NoService {
			createGetters(&t)
			sortedServices = append(sortedServices, t)
		}
	}

	// Services can be subclasses of other services. Copy methods from each service's 'Base' entity to
	// the child service, only if a same-named method does not already exist (i.e., overridden by the
	// child service)
	for i, service := range sortedServices {
		sortedServices[i].Methods = getBaseMethods(service, meta)
		fixReturnType(&sortedServices[i])
	}

	fmt.Printf("Creating Datatypes...\n")
	err = writePackage(outputDir, "datatypes", sortedTypes, datatype)
	errCheck(err, "Error creating Datatypes")

	fmt.Printf("Creating Servicess...\n")
	err = writePackage(outputDir, "services", sortedServices, services)
	errCheck(err, "Error creating Services")
}

func errCheck(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s\n", msg, err)
		os.Exit(1)
	}
}