package cmd

import (
	"fmt"
	"errors"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/session"
)

func init() {
	vlanNameCmd.Flags().StringVarP(&VlanName, "name", "n", "", "New VLAN name")
	vlanNameCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(vlanNameCmd)
}

var VlanName string
var vlanNameCmd = &cobra.Command{
	Use:   "vlan-name [vlanId]",
	Short: "Set a VLAN's name",
	Long:  `Set a VLAN's name.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("a VLAN ID is required")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return VlanNameCommand(cmd, args)
	},
}

func VlanNameCommand(cmd *cobra.Command, args []string) error {
	// Sets up the softlayer-client session and Service
	sess := session.New()
	// sess.Debug = true
	service := services.GetNetworkVlanService(sess)

	// Make sure vlanId is an int
	vlanId, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	// https://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_Vlan/
	// only Name and Note are editable on a Vlan object.
	vlanSkel := datatypes.Network_Vlan{
		Name: &VlanName,
	}
	// https://sldn.softlayer.com/reference/services/SoftLayer_Network_Vlan/editObject/
	_, err = service.Id(vlanId).EditObject(&vlanSkel)
	if err != nil {
		return err
	}
	fmt.Printf("Set name of VLan %v to %v\n", vlanId, VlanName)
	return nil
}
