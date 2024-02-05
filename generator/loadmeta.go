/**
 * Copyright 2016 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package generator

import (
	"bytes"

	"fmt"
	"github.com/spf13/cobra"
	"go/format"
	"golang.org/x/tools/imports"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
	"text/template"
)

var rootCmd = &cobra.Command{
	Use:   "tools",
	Short: "Generates the services and datatype definitions from SoftLayer API Metadata",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// rootCmd.AddCommand(generateCmd)
	// rootCmd.AddCommand(versionCmd)
}

func NilParam(params []Parameter) string {
	rString := ""
	if len(params) == 0 {
		return rString
	}
	rString = "nil"
	// The place/verify/quoteOrder are special
	if params[0].Type == "SoftLayer_Container_Product_Order" {
		rString = fmt.Sprintf("%s", "GetOrderContainer()")
	}
	for _, _ = range params[1:] {
		// rString = fmt.Sprintf("%s,%s-%s", rString, param.Name, param.Type)
		rString = fmt.Sprintf("%s, nil", rString)
	}
	return rString
}

func RemovePrefix(args ...interface{}) string {
	s := args[0].(string)

	if strings.HasPrefix(s, "SoftLayer_") {
		return s[10:]
	}

	return s
}

// ConvertType takes the name of the type to convert, and the package context.
func ConvertType(args ...interface{}) string {
	t := args[0].(string)
	p := args[1].(string)

	// Convert softlayer types to golang types
	switch t {
	case "unsignedLong", "unsignedInt":
		return "uint"
	case "boolean":
		return "bool"
	case "dateTime":
		if p != "datatypes" {
			return "datatypes.Time"
		} else {
			return "Time"
		}
	case "decimal", "float":
		if p != "datatypes" {
			return "datatypes.Float64"
		} else {
			return "Float64"
		}
	case "base64Binary":
		return "[]byte"
	case "json", "enum":
		return "string"
	}

	if strings.HasPrefix(t, "SoftLayer_") {
		t = RemovePrefix(t)
		if p != "datatypes" {
			return "datatypes." + t
		}
		// A Property called Resource that is an 'Entity' can be multiple types in reality
		// Specifically this is for Container_Search_Result
		if len(args) >= 4 && args[3] == "resource" && t == "Entity" {
			return "interface{}"
		}
		return t
	}

	// Need to update this if SL ever exposes classes that don't start with SoftLayer_
	// We could proboably assume anything that doesn't get type matched above should be a datatype
	// but I feel this is more specific.
	validPrefix := []string{"McAfee_", "Sprint_", "BMS_"}
	for _, prefix := range validPrefix {
		if strings.HasPrefix(t, prefix) {
			if p != "datatypes" {
				return "datatypes." + t
			}
			return t
		}
	}

	return t
}

func RemoveReservedWords(args ...interface{}) string {
	n := args[0].(string)

	// Replace language reserved identifiers with alternatives
	switch n {
	case "type":
		return "typ"
	}

	return n
}

// Remove '_' from Snake_Case values
func Desnake(args ...interface{}) string {
	s := args[0].(string)
	return strings.Replace(s, "_", "", -1)
}

// Formats a string into a comment.  For now, just each comment line with "//"
func GoDoc(args ...interface{}) string {
	s := args[0].(string)
	if s == "" {
		s = "no documentation yet"
	}

	return "// " + strings.Replace(s, "\n", "\n// ", -1)
}

// Remove omitempty tags if required
func Tags(args ...interface{}) string {
	n := args[0].(string)

	switch n {
	case "resourceRecords":
		return n
	default:
		return n + ",omitempty"
	}
}

func CreateGetters(service *Type) {
	for _, p := range service.Properties {
		if p.Form == "relational" {
			m := Method{
				Name:       "get" + strings.Title(p.Name),
				Type:       p.Type,
				TypeArray:  p.TypeArray,
				Doc:        "Retrieve " + p.Doc, // TODO lowercase the first letter
				Parameters: []Parameter{},
			}

			service.Methods[m.Name] = m
		}
	}
}

// Special case for ensuring we can set a complexType on product orders.
func AddComplexType(dataType *Type) {
	// Only adding this to the base product order type. All others embed this one.
	if dataType.Name == "SoftLayer_Container_Product_Order" {
		dataType.Properties["complexType"] = Property{
			Name: "complexType",
			Type: "string",
			Form: "local",
			Doc:  "Added by softlayer-go. This hints to the API what kind of product order this is.",
		}
	} else if dataType.Name == "SoftLayer_Container_User_Customer_External_Binding" {
		dataType.Properties["complexType"] = Property{
			Name: "complexType",
			Type: "string",
			Form: "local",
			Doc:  "Added by softlayer-go. This hints to the API what kind of binding this is.",
		}
	}
}

// Special case for fixing some datatype properties in the metadata so that these properties can be set in their base
func FixDatatype(t *Type, meta map[string]Type) {
	if strings.HasPrefix(t.Name, "SoftLayer_Dns_Domain_ResourceRecord_") {
		baseRecordType, _ := meta["SoftLayer_Dns_Domain_ResourceRecord"]
		for propName, prop := range t.Properties {
			baseRecordType.Properties[propName] = prop
		}
		meta["SoftLayer_Dns_Domain_ResourceRecord"] = baseRecordType
	} else if t.Name == "SoftLayer_Container_User_Customer_External_Binding_Verisign" {
		baseType, _ := meta["SoftLayer_Container_User_Customer_External_Binding"]
		for propName, prop := range t.Properties {
			baseType.Properties[propName] = prop
		}
		meta["SoftLayer_Container_User_Customer_External_Binding"] = baseType
	}
}

// Special case for fixing some broken return types in the metadata
func FixReturnType(service *Type) {
	brokenServices := map[string]string{
		"SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service":       "deleteObject",
		"SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer": "deleteObject",
		"SoftLayer_Network_Application_Delivery_Controller":                            "deleteLiveLoadBalancerService",
	}

	if methodName, ok := brokenServices[service.Name]; ok {
		method := service.Methods[methodName]
		method.Type = "void"
		service.Methods[methodName] = method
	}
}

// Return formatted method argument phrase used by the method generation.
func phraseMethodArg(methodName string, argName string, isArray bool, argType string) string {
	argName = RemoveReservedWords(argName)

	// Handle special case - placeOrder/verifyOrder should take any kind of order type.
	if (methodName == "placeOrder" || methodName == "verifyOrder" || methodName == "placeQuote") &&
		strings.HasPrefix(argType, "SoftLayer_Container_Product_Order") {
		return fmt.Sprintf("%s interface{}, ", argName)
	}

	refPrefix := "*"
	if isArray {
		refPrefix = "[]"
	}

	argType = ConvertType(argType, "services")

	return fmt.Sprintf("%s %s%s, ", argName, refPrefix, argType)
}

func CombineMethods(baseMethods map[string]Method, subclassMethods map[string]Method) map[string]Method {
	r := map[string]Method{}

	// Copy all subclass methods into the result set
	for k, v := range subclassMethods {
		r[k] = v
	}

	// Copy each method from the base class into the result set, but only if a like-named method
	// does not already exist (a method in the child should override a same-named method in the parent)
	for k, v := range baseMethods {
		if _, ok := r[k]; !ok {
			r[k] = v
		}
	}

	return r
}

func GetBaseMethods(s Type, typeMap map[string]Type) map[string]Method {
	var methods, baseMethods map[string]Method

	methods = s.Methods

	if s.Base != "SoftLayer_Entity" {
		baseMethods = GetBaseMethods(typeMap[s.Base], typeMap)

		// Add base methods to current service methods
		methods = CombineMethods(baseMethods, methods)
	}

	// return my methods
	return methods
}

func GetSortedKeys(m map[string]Type) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		if validCheck(key) {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)

	return keys
}

func validCheck(name string) bool {

	if strings.HasPrefix(name, "BluePages_") {
		return false
	}
	if strings.HasPrefix(name, "IntegratedOfferingTeam_") {
		return false
	}
	return true
}
func WritePackage(base string, pkg string, meta []Type, ts string) error {
	var currPrefix string
	var start int

	for i, t := range meta {

		components := strings.Split(RemovePrefix(t.Name), "_")
		meta[i].ServiceGroup = components[0]
		if i == 0 {
			currPrefix = components[0]
			continue
		}

		if components[0] != currPrefix {
			err := WriteGoFile(base, pkg, currPrefix, meta[start:i], ts)
			if err != nil {
				return err
			}

			currPrefix = components[0]
			start = i
		}
	}

	WriteGoFile(base, pkg, currPrefix, meta[start:], ts)

	return nil
}

// Executes a template against the metadata structure, and generates a go source file with the result
func WriteGoFile(base string, pkg string, name string, meta []Type, ts string) error {
	filename := base + "/" + pkg + "/" + strings.ToLower(name) + ".go"
	fmt.Printf("Creating %s\n", filename)
	// Generate the source
	var buf bytes.Buffer
	t := template.New(pkg).Funcs(fMap)
	template.Must(t.Parse(ts)).Execute(&buf, meta)

	/*if pkg == "services" && name == "Account"{
		fmt.Println(string(buf.String()))
		os.Exit(0)
	}*/
	// Add the imports
	src, err := imports.Process(filename, buf.Bytes(), &imports.Options{Comments: true})
	if err != nil {
		fmt.Printf("Error processing imports: %s", err)
	}

	// Format
	pretty, err := format.Source(src)
	if err != nil {
		return fmt.Errorf("Error while formatting source: %s", err)
	}

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Error creating %s: %s", filename, err)
	}
	defer f.Close()
	fmt.Fprintf(f, "%s", pretty)
	if pkg == "services" {
		generateTestFiles(base, meta, strings.ToLower(name))
	}

	return nil
}

func generateTestFiles(base string, meta []Type, groupName string) {
	// These are technically a service... but dont have any methods to test...
	if groupName == "exception" {
		return
	}
	path := fmt.Sprintf("%s/services/%s_test.go", base, groupName)
	fmt.Printf("Creating %s\n", path)
	template := template.New(path)
	template.Funcs(fMap)
	_, err := template.Parse(testTemplate)
	if err != nil {
		fmt.Printf("Error Parsing Template: %s\n", err)
		return
	}
	outfile, err := os.Create(path)
	if err != nil {
		fmt.Printf("Error Creating File %s: %s\n", path, err)
		return
	}
	defer outfile.Close()
	err = template.Execute(outfile, meta)
	if err != nil {
		fmt.Printf("Error Executing Template: %s\n", err)
		return
	}
}

func GetMetaFromURL(url string) ([]byte, error) {
	client := http.DefaultClient

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func GetMetaFromFile(file string) ([]byte, error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()
	return ioutil.ReadAll(jsonFile)
}

func getTypePrefix(isArray bool, theType string) string {
	if isArray {
		return "[]"
	}
	if theType != "interface{}" {
		return "*"
	}
	return ""
}

func deprecatedDoc(isDeprecated bool) string {
	if isDeprecated {
		return "\n// Deprecated: This function has been marked as deprecated."
	}
	return ""
}
