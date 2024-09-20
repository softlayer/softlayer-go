package cmd

import (
	"fmt"
	"time"
	"github.com/spf13/cobra"

	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/session"
)

func init() {
	rootCmd.AddCommand(iamDemoCmd)
}

var iamDemoCmd = &cobra.Command{
	Use:   "iam-demo",
	Short: "Will make 1 API call per minute and refresh API key when needed.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return RunIamCmd(cmd, args)
	},
}

func RunIamCmd(cmd *cobra.Command, args []string) error {
	objectMask := "mask[id,companyName]"

	// Sets up the session with authentication headers.
	sess := &session.Session{
		Endpoint: session.DefaultEndpoint,
		IAMToken: "Bearer TOKEN",
		IAMRefreshToken: "REFRESH TOKEN",
		Debug: true,
	}

	// creates a reference to the service object (SoftLayer_Account)
	service := services.GetAccountService(sess)

	// Sets the mask, filter, result limit, and then makes the API call SoftLayer_Account::getHardware()


	for {
		account, err := service.Mask(objectMask).GetObject()
		if err != nil {
			fmt.Printf("======= ERROR ======")
			return err
		}
		fmt.Printf("AccountId: %v, CompanyName: %v\n", *account.Id, *account.CompanyName)
		fmt.Printf("Refreshing Token for no reason...\n")
		sess.RefreshToken()
		fmt.Printf("%s\n", sess.IAMToken)
		fmt.Printf("Sleeping for 60s.......\n")
		time.Sleep(60 * time.Second)
	}

	return nil
}
