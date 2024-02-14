package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/softlayer/softlayer-go/filter"
	"github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/helpers/virtual"
	"github.com/softlayer/softlayer-go/sl"
)

func init() {
	rootCmd.AddCommand(listVirtCmd)
}

var listVirtCmd = &cobra.Command{
	Use:   "virt-list",
	Short: "Lists all VSI on the account",
	Long:  `Lists all VSI on the account using an iterative aproach.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return RunListVirtCmd(cmd, args)
	},
}

func RunListVirtCmd(cmd *cobra.Command, args []string) error {

	objectMask := "mask[id,hostname,domain,primaryIpAddress,primaryBackendIpAddress]"
	// When using a result Limit to break up your API request, its important to include an orderBy objectFilter
	// to enforce an order on the query, as the database might not always return results in the same order between
	// queries otherwise
	filters := filter.New()
	filters = append(filters, filter.Path("virtualGuests.id").OrderBy("ASC"))
	objectFilter := filters.Build()
	// Sets up the session with authentication headers.
	sess := session.New()
	// uncomment to output API calls as they are made.
	sess.Debug = true

	// Sets the mask, filter, result limit, and then makes the API call SoftLayer_Account::getHardware()
	limit := 5
	options := sl.Options{
		Mask: objectMask,
		Filter: objectFilter,
		Limit: &limit,
	}

	servers, err := virtual.GetVirtualGuestsIter(sess, &options)
	if err != nil {
		return err
	}
	fmt.Printf("Id, Hostname, Domain, IP Address\n")

	for _, server := range servers {
		ipAddress := "-"
		// Servers with a private only connection will not have a primary IP
		if server.PrimaryIpAddress != nil {
			ipAddress = *server.PrimaryIpAddress
		}
		fmt.Printf("%v, %v, %v, %v\n", *server.Id, *server.Hostname, *server.Domain, ipAddress)
	}


	return nil
}
