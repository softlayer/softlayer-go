package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/softlayer/softlayer-go/filter"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/session"
)

func init() {
	rootCmd.AddCommand(listServerCmd)
}

var listServerCmd = &cobra.Command{
	Use:   "server-list",
	Short: "Lists all servers on the account",
	Long:  `Lists all servers on the account`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return Run(cmd, args)
	},
}

func Run(cmd *cobra.Command, args []string) error {
	resultLimit := 5
	resultOffset := 0
	objectMask := "mask[id,hostname,domain,primaryIpAddress,primaryBackendIpAddress]"
	// When using a result Limit to break up your API request, its important to include an orderBy objectFilter
	// to enforce an order on the query, as the database might not always return results in the same order between
	// queries otherwise
	filters := filter.New()
	filters = append(filters, filter.Path("hardware.id").OrderBy("ASC"))
	objectFilter := filters.Build()
	// Sets up the session with authentication headers.
	sess := session.New()
	// uncomment to output API calls as they are made.
	// sess.Debug = true

	// creates a reference to the service object (SoftLayer_Account)
	service := services.GetAccountService(sess)

	// Sets the mask, filter, result limit, and then makes the API call SoftLayer_Account::getHardware()
	servers, err := service.Mask(objectMask).Filter(objectFilter).
		Offset(resultOffset).Limit(resultLimit).GetHardware()
	if err != nil {
		return err
	}
	fmt.Printf("Id, Hostname, Domain, IP Address\n")
	for {
		for _, server := range servers {
			ipAddress := "-"
			if *server.PrimaryIpAddress != "" {
				ipAddress = *server.PrimaryIpAddress
			}
			fmt.Printf("%v, %v, %v, %v\n", *server.Id, *server.Hostname, *server.Domain, ipAddress)
		}
		// If we get less than the number of results we asked for, we are at the end of our server list
		if len(servers) < resultLimit {
			break
		}
		// Increment the offset to get next set of results
		resultOffset = resultOffset + resultLimit
		servers, err = service.Mask(objectMask).Filter(objectFilter).
			Offset(resultOffset).Limit(resultLimit).GetHardware()
		if err != nil {
			return err
		}
	}

	return nil
}
