package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/filter"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/sl"
)

// Create the session and services
var sess = session.New()
var packageService = services.GetProductPackageService(sess)
var orderService = services.GetProductOrderService(sess)

// Parameters will help us to reduce the number of argumments in a function
// and not all of them are required. e.g. hostID, reservedCapacityID,sshKeys, etc.
type Parameters struct {
	hostname           string
	domain             string
	pkgName            string
	flavor             string
	hourly             bool
	hostID             int // To order in an specific dedicated host
	reservedCapacityID int // To order in a reerved capacity
	location           string
	items              []string
	publicVlanID       int
	privateVlanID      int
	sshKeys            []int
	postScripts        []string
}

func main_one() {
	// KeyNames of the items
	items := []string{
		"1_GBPS_PUBLIC_PRIVATE_NETWORK_UPLINKS",
		"BANDWIDTH_0_GB_2", "1_IP_ADDRESS",
		"OS_UBUNTU_16_04_LTS_XENIAL_XERUS_MINIMAL_64_BIT_FOR_VSI",
		"REBOOT_REMOTE_CONSOLE", "NESSUS_VULNERABILITY_ASSESSMENT_REPORTING",
		"MONITORING_HOST_PING", "NOTIFICATION_EMAIL_AND_TICKET", "AUTOMATED_NOTIFICATION",
		"UNLIMITED_SSL_VPN_USERS_1_PPTP_VPN_USER_PER_ACCOUNT",
	}

	parameters := Parameters{
		hostname: "tester",
		domain:   "cgallo.com",
		// publicVlanID:  2068353,
		// privateVlanID: 2068355,
		// CLOUD_SERVER, SUSPEND_CLOUD_SERVER, PUBLIC_CLOUD_SERVER or TRANSIENT_CLOUD_SERVER
		pkgName:  "SUSPEND_CLOUD_SERVER",
		flavor:   "M1_1X8X25", // Remove this if you want to order a CLOUD_SERVER
		location: "DALLAS13",
		hourly:   true, // Suspended and transient vsis can only be hourly
		items:    items,
	}

	OrderVirtualServer(parameters)
}

// OrderVirtualServer place/verify the order
func OrderVirtualServer(args Parameters) {

	// Get the packageId and prices
	packageID := GetPackageID(args.pkgName)
	prices := GetPrices(*packageID, args.items)

	// Build the Virtual_Guest array object
	virtualGuests := BuildArrayVirtualGuests(args)

	// Build the Container_Product_Order with the basic required data
	containerOrder := datatypes.Container_Product_Order{
		PackageId:        packageID,
		Location:         sl.String(args.location),
		VirtualGuests:    virtualGuests,
		Prices:           prices,
		UseHourlyPricing: sl.Bool(args.hourly),
	}
	// Add presetId to the Container_Product_Order if a flavor exists
	if args.flavor != "" {
		containerOrder.PresetId = GetPresetID(*packageID, args.flavor)
	}

	// Build the order template that will be used to order the virtual server
	orderTemplate := datatypes.Container_Product_Order_Virtual_Guest{
		Container_Product_Order_Hardware_Server: datatypes.Container_Product_Order_Hardware_Server{
			Container_Product_Order: containerOrder,
		},
	}

	// Use verifyOrder() method to check for errors. Replace this with placeOrder() when
	// you are ready to order.
	// receipt, err := orderService.VerifyOrder(&orderTemplate)
	receipt, err := orderService.PlaceQuote(&orderTemplate)
	if err != nil {
		fmt.Printf("Unable to verify/place the order: %s", err)
	}

	// pretty print
	jsonFormat, jsonErr := json.MarshalIndent(receipt, "", "   ")
	if jsonErr != nil {
		fmt.Println(jsonErr)
		return
	}
	fmt.Println(string(jsonFormat))
}

// GetPackageID returns the package id
func GetPackageID(keyName string) *int {
	objFilter := filter.Path("keyName").Eq(keyName).Build()
	packages, err := packageService.Mask("mask[id]").Filter(objFilter).GetAllObjects()
	if err != nil {
		fmt.Printf("Package %s doesn't exists\n", keyName)
		os.Exit(1)
	}
	return packages[0].Id
}

// GetPresetID returns the presetID of a flavor
func GetPresetID(packageID int, flavor string) *int {
	objFilter := filter.Path("keyName").Eq(flavor).Build()
	presets, err := packageService.Filter(objFilter).Id(packageID).GetActivePresets()
	if err != nil {
		fmt.Printf("Flavor %s doesn't exists\n", flavor)
		os.Exit(1)
	}
	return presets[0].Id
}

// GetPrices returns the list of prices
func GetPrices(packageID int, itemKeyNames []string) []datatypes.Product_Item_Price {
	objMask := "mask[id,itemCategory,keyName,prices[categories]]"

	items, err := packageService.Mask(objMask).Id(packageID).GetItems()
	if err != nil {
		fmt.Printf("Unable to retrieve items: %s\n", err)
		os.Exit(1)
	}

	prices := make([]datatypes.Product_Item_Price, len(itemKeyNames))

	for idx, itemKeyName := range itemKeyNames {
		prices[idx].Id, err = GetPriceIDIfMatch(itemKeyName, items)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	return prices
}

// GetPriceIDIfMatch returns a priceID or error if keyName doesnt exists
func GetPriceIDIfMatch(keyName string, items []datatypes.Product_Item) (priceID *int, err error) {
	var matchingItem datatypes.Product_Item
	// find the item keyname
	for _, item := range items {
		if keyName == *item.KeyName {
			matchingItem = item
			break
		}
	}
	// returns error if keyName doesn't exists in the items list
	if matchingItem.KeyName == nil {
		return nil, fmt.Errorf("The item %s doesn't exists in the package", keyName)
	}
	// find the standard priceID
	for _, price := range matchingItem.Prices {
		if price.LocationGroupId == nil {
			priceID = price.Id
			break
		}
	}

	return
}

// BuildArrayVirtualGuests builds the esqueleton of Virtual_Guest array
func BuildArrayVirtualGuests(args Parameters) []datatypes.Virtual_Guest {
	virtualGuests := []datatypes.Virtual_Guest{
		{
			Hostname: sl.String(args.hostname),
			Domain:   sl.String(args.domain),
		},
	}

	if args.privateVlanID > 0 {
		virtualGuests[0].PrimaryBackendNetworkComponent = &datatypes.Virtual_Guest_Network_Component{
			NetworkVlan: &datatypes.Network_Vlan{
				Id: sl.Int(args.privateVlanID),
			},
		}
	}

	if args.publicVlanID > 0 {
		virtualGuests[0].PrimaryNetworkComponent = &datatypes.Virtual_Guest_Network_Component{
			NetworkVlan: &datatypes.Network_Vlan{
				Id: sl.Int(args.publicVlanID),
			},
		}
	}
	return virtualGuests
}
