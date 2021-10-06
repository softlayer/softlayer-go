package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/softlayer/softlayer-go/filter"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/session"
)

func init() {
	listVlanCmd.Flags().StringVarP(&Datacenter, "datacenter", "d", "", "List only VLANs from this datacenter")
	rootCmd.AddCommand(listVlanCmd)
}

var Datacenter string
var listVlanCmd = &cobra.Command{
	Use:   "vlan-list",
	Short: "Lists all vlans on the account",
	Long:  `Lists all vlans on the account`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return VlanListCommand(cmd, args)
	},
}

func VlanListCommand(cmd *cobra.Command, args []string) error {
	resultLimit := 50
	resultOffset := 0
	objectMask := "mask[id,name,vlanNumber,primaryRouter[id,hostname,datacenterName,datacenter[name]]]"
	// When using a result Limit to break up your API request, its important to include an orderBy objectFilter
	// to enforce an order on the query, as the database might not always return results in the same order between
	// queries otherwise
	filters := filter.New()
	filters = append(filters, filter.Path("networkVlans.id").OrderBy("DESC"))
	if Datacenter != "" {
		filters = append(filters, filter.Path("networkVlans.primaryRouter.datacenter.name").Eq(Datacenter))
	}
	objectFilter := filters.Build()
	// fmt.Println("FILTER: ", objectFilter)

	// Sets up the session with authentication headers.
	sess := session.New()
	// uncomment to output API calls as they are made.
	// sess.Debug = true

	// creates a reference to the service object (SoftLayer_Account)
	service := services.GetAccountService(sess)

	// Sets the mask, filter, result limit, and then makes the API call SoftLayer_Account::getHardware()
	vlans, err := service.Mask(objectMask).Filter(objectFilter).
		Offset(resultOffset).Limit(resultLimit).GetNetworkVlans()
	if err != nil {
		return err
	}
	fmt.Printf("Id, VLAN, Name, Datacenter\n")
	for {

		for _, s := range vlans {
			name := "-"
			// The Name property doesn't get returned if the vlan hasn't been named, which will cause a
			// 'panic: runtime error: invalid memory address' error if we dont check first.
			if s.Name != nil {
				name = *s.Name
			}
			fmt.Printf("%v, %v, %v, %v\n", *s.Id, name, *s.VlanNumber, *s.PrimaryRouter.Datacenter.Name)
		}
		// If we get less than the number of results we asked for, we are at the end of our server list
		if len(vlans) < resultLimit {
			break
		}
		// Increment the offset to get next set of results
		resultOffset = resultOffset + resultLimit

		vlans, err = service.Mask(objectMask).Filter(objectFilter).
			Offset(resultOffset).Limit(resultLimit).GetNetworkVlans()
		if err != nil {
			return err
		}
	}

	return nil
}
