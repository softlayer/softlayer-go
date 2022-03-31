package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/sl"
)

func init() {
	orderVlanCmd.Flags().StringVarP(&Location, "datacenter", "d", "", "Datacenter name to order in (like dal13)")
	orderVlanCmd.Flags().StringVarP(&Name, "name", "n", "", "Name of this new vlan")
	orderVlanCmd.Flags().IntVarP(&RouterId, "routerId", "r", 0, "RouterId if you want this vlan in a specific POD in a datacenter")
	rootCmd.AddCommand(orderVlanCmd)
}

var Location string
var Name string
var RouterId int

var orderVlanCmd = &cobra.Command{
	Use:   "vlan-order",
	Short: "Orders a Vlan",
	Long:  `Orders a vlan with 32 static public IP addresses.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return VlanOrderCommand(cmd, args)
	},
}

func VlanOrderCommand(cmd *cobra.Command, args []string) error {
	// Declare the properties like complexType, location, packageId and quantity for the
	// vlan you wish to order
	quantity := 1
	packageId := 0
	sendQuoteEmailFlag := true

	// Build a skeleton SoftLayer_Product_Item_Price objects. To get the list of valid
	// prices for the package use the SoftLayer_Product_Package:getItems method
	itemPrices := []datatypes.Product_Item_Price{
		{Id: sl.Int(2018)}, // Price for the new Public Network Vlan
		{Id: sl.Int(1092)}, // Price for 32 Static Public IP Addresses
	}

	// Create a session
	sess := session.New()

	// Get SoftLayer_Product_Order service
	service := services.GetProductOrderService(sess)

	// Build the Container_Product_Order_Network_Vlan containing the order you wish to place.
	orderTemplate := datatypes.Container_Product_Order_Network_Vlan{
		Container_Product_Order: datatypes.Container_Product_Order{
			Prices:             itemPrices,
			PackageId:          sl.Int(packageId),
			Location:           sl.String(Location),
			Quantity:           sl.Int(quantity),
			SendQuoteEmailFlag: sl.Bool(sendQuoteEmailFlag),
		},
		Name: sl.String(Name),
	}

	if RouterId > 0 {
		orderTemplate.RouterId = sl.Int(RouterId)
	}

	// Use verifyOrder() method to check for errors. Replace this with placeOrder() when
	// you are ready to order.
	receipt, err := service.VerifyOrder(&orderTemplate)
	if err != nil {
		fmt.Printf("Unable to place order:\n - %s\n", err)
		return err
	}

	// Following helps to print the result in json format.
	jsonFormat, jsonErr := json.MarshalIndent(receipt, "", "    ")
	if jsonErr != nil {
		fmt.Println(jsonErr)
		return jsonErr
	}

	fmt.Println(string(jsonFormat))
	return nil
}
