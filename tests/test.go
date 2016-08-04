package main

import (
	"fmt"
	"os"

	"github.ibm.com/riethm/gopherlayer/datatypes"
	"github.ibm.com/riethm/gopherlayer/service"
	"github.ibm.com/riethm/gopherlayer/sl"
	"time"
)

func main() {
	session := service.NewSession("username", "apikey") // default endpoint

	// Example: List all Virtual Guests for an account

	// Get the Account service
	acctService := session.GetAccountService()

	//Set an object mask and a result limit
	acctService.Mask("id;hostname;domain").Limit(10)

	// List VMs
	vms, err := acctService.GetVirtualGuests()
	if err != nil {
		fmt.Printf("Error retrieving Virtual Guests from Account: %s\n", err)
		os.Exit(1)
	} else {
		fmt.Println("VMs under Account:")
	}

	for _, vm := range vms {
		fmt.Printf("\t[%d]%s.%s\n", *vm.Id, *vm.Hostname, *vm.Domain)
	}

	// Example: Execute a remote script on a Virtual Guest

	// Get the VirtualGuest service
	vGuestService := session.GetVirtualGuestService()

	// Set the object ID for the service to act upon
	vGuestService.Id(22870595)

	// Execute the remote script
	err = vGuestService.ExecuteRemoteScript(sl.String("http://example.com"))
	if err != nil {
		fmt.Printf("Error executing remote script on VM: %s", err)
	} else {
		fmt.Println("Remote script sent for execution on VM")
	}

	doCreateVMTest(&session)
}

func doCreateVMTest(session *service.Session) {
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
	} else {
		fmt.Printf("\nNew Virtual Guest created with ID %d\n", *vGuest.Id)
		fmt.Printf("Domain: %s\n", *vGuest.Domain)
	}

	// Wait for transactions to finish
	fmt.Printf("Waiting for transactions to complete before destroying.")
	service.Id(*vGuest.Id)

	// Delay to allow transactions to be registered
	time.Sleep(5 * time.Second)

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
