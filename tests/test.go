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

package main

import (
	"fmt"

	"github.ibm.com/riethm/gopherlayer/datatypes"
	"github.ibm.com/riethm/gopherlayer/services"
	"github.ibm.com/riethm/gopherlayer/sl"
	"reflect"
	"time"
)

func main() {
	session := services.NewSession("username", "apikey") // default endpoint

	session.Debug = true

	// List all Virtual Guests for an account
	//doListAccountVMsTest(&session)

	// Execute a remote script on a Virtual Guest
	//doExecuteRemoteScriptTest(&session)

	// Example: Provision and destroy a Virtual Guest
	//doCreateVMTest(&session)

	// Example: Get disk usage metrics by date
	//doGetDiskUsageMetricsTest(&session)

	// Example: Get the last bill date
	//doGetLatestBillDate(&session)

	// Demonstrate API Error
	//doError(&session)
}

func doListAccountVMsTest(session *services.Session) {
	// Get the Account service
	service := session.GetAccountService()

	//Set an object mask and a result limit
	service.Mask("id;hostname;domain").Limit(10)

	// List VMs
	vms, err := service.GetVirtualGuests()
	if err != nil {
		fmt.Printf("Error retrieving Virtual Guests from Account: %s\n", err)
		return
	} else {
		fmt.Println("VMs under Account:")
	}

	for _, vm := range vms {
		fmt.Printf("\t[%d]%s.%s\n", *vm.Id, *vm.Hostname, *vm.Domain)
	}
}

func doExecuteRemoteScriptTest(session *services.Session) {
	// Get the VirtualGuest service
	service := session.GetVirtualGuestService()

	// Set the object ID for the service to act upon
	service.Id(22870595)

	// Execute the remote script
	err := service.ExecuteRemoteScript(sl.String("http://example.com"))
	if err != nil {
		fmt.Println("Error executing remote script on VM:", err)
	} else {
		fmt.Println("Remote script sent for execution on VM")
	}
}

func doCreateVMTest(session *services.Session) {
	service := session.GetVirtualGuestService()

	// Create a Virtual_Guest instance as a template
	vGuestTemplate := datatypes.Virtual_Guest{}

	//Set Creation values - use helpers from the sl package to set pointer values
	vGuestTemplate.Hostname = sl.String("sample")
	vGuestTemplate.Domain = sl.String("example.com")
	vGuestTemplate.MaxMemory = sl.Int(4096)
	vGuestTemplate.StartCpus = sl.Int(1)
	vGuestTemplate.Datacenter = &datatypes.Location{Name: sl.String("wdc01")}
	vGuestTemplate.OperatingSystemReferenceCode = sl.String("UBUNTU_LATEST")
	vGuestTemplate.LocalDiskFlag = sl.Bool(true)

	service.Mask("id;domain")

	vGuest, err := service.CreateObject(&vGuestTemplate)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	} else {
		fmt.Printf("\nNew Virtual Guest created with ID %d\n", *vGuest.Id)
		fmt.Printf("Domain: %s\n", *vGuest.Domain)
	}

	// Wait for transactions to finish
	fmt.Printf("Waiting for transactions to complete before destroying.")
	service.Id(*vGuest.Id)

	// Delay to allow transactions to be registered
	time.Sleep(10 * time.Second)

	for transactions, _ := service.GetActiveTransactions(); len(transactions) > 0; {
		fmt.Print(".")
		time.Sleep(10 * time.Second)
		transactions, err = service.GetActiveTransactions()
	}

	fmt.Println("Deleting virtual guest")

	success, err := service.DeleteObject()
	if err != nil {
		fmt.Printf("Error deleting virtual guest: %s", err)
	} else if success == false {
		fmt.Printf("Error deleting virtual guest")
	} else {
		fmt.Printf("Virtual Guest deleted successfully")
	}
}

func doGetDiskUsageMetricsTest(session *services.Session) {
	service := session.GetAccountService()

	tEnd := sl.Time(time.Now())
	tStart := sl.Time(tEnd.AddDate(0, -6, 0))

	data, err := service.GetDiskUsageMetricDataByDate(tStart, tEnd)
	if err != nil {
		fmt.Println("Error calling GetDiskUsageMetricDataByDate: ", err)
		return
	}

	fmt.Printf("Number of elements returned: %d\n", len(data))

	// Retrieve and print a DateTime (sl.Time) value
	if len(data) > 0 {
		fmt.Printf("item.DateTime = %s\n", data[0].DateTime)
	}
}

func doGetLatestBillDate(session *services.Session) {
	service := session.GetAccountService()

	d, _ := service.GetLatestBillDate()

	fmt.Printf("date of last bill: %s\n", d)
	fmt.Printf("type of date field: %s\n", reflect.TypeOf(d))
}

func handleError(err error) {
	apiErr := err.(services.Error)
	fmt.Printf(
		"Exception: %s\nMessage: %s\nHTTP Status Code: %d\n",
		apiErr.Exception,
		apiErr.Message,
		apiErr.StatusCode)

	// Note that we could instead just dump the error, if we are not interested
	// in the individual fields
	//fmt.Println("Error:", err)
}

func doError(session *services.Session) {
	service := session.GetVirtualGuestService()

	// Example of an API error
	service.Id(0)  // invalid object ID
	_, err := service.GetObject()
	if err != nil {
		handleError(err)
	}

	// Example of an HTTP, but non-API error
	session.Endpoint = "http://example.com"  // invalid endpoint
	_, err = service.GetObject()
	if err != nil {
		handleError(err)
	}

	// Example of a non-HTTP, non-API error
	session.Endpoint = services.DEFAULT_ENDPOINT
	var result struct {
		Id string `json:"id"`  // type mismatch (unmarshal an integer value into a string)
	}
	err = session.DoRequest("SoftLayer_Account", "getObject", nil, &services.Options{}, &result)
	if err != nil {
		handleError(err)
	}
}