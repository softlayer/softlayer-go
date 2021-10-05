package cmd

import (
	"errors"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/session"
)

func init() {
	rootCmd.AddCommand(detailVlanCmd)
}

var detailVlanCmd = &cobra.Command{
	Use:   "vlan-detail [vlanId]",
	Short: "Vlan Details",
	Long:  `Gets a lot of information about a VLAN`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("a VLAN ID is required")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return VlanDetailCommand(cmd, args)
	},
}

func VlanDetailCommand(cmd *cobra.Command, args []string) error {
	// resultLimit := 50
	// resultOffset := 0
	objectMask := `mask[id,name,vlanNumber,primaryRouter[id,hostname,datacenterName,datacenter[name]],
subnets[id,networkIdentifier,note, subnetType]]`

	// Sets up the softlayer-client session and Service
	sess := session.New()
	// sess.Debug = true
	service := services.GetNetworkVlanService(sess)

	// Make sure vlanId is an int
	vlanId, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	// Get the VLAN object from the API https://sldn.softlayer.com/reference/services/SoftLayer_Network_Vlan/getObject/
	vlan, err := service.Mask(objectMask).Id(vlanId).GetObject()
	if err != nil {
		return err
	}
	
	// https://github.com/jedib0t/go-pretty/tree/main/table
	// A fancy table builder so our output can look nice.
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Property", "Value"})

	t.AppendRow([]interface{}{"Id", *vlan.Id})
	t.AppendRow([]interface{}{"Number", *vlan.VlanNumber})
	t.AppendRow([]interface{}{"Name", *vlan.Name})
	t.AppendRow([]interface{}{"Datacenter", *vlan.PrimaryRouter.DatacenterName})
	t.AppendRow([]interface{}{"Primary Router", *vlan.PrimaryRouter.Hostname})

	if len(vlan.Subnets) > 0  {
		subnetTable := table.NewWriter()
		subnetTable.AppendHeader(table.Row{"Id", "Network", "Type", "Note"})
		for _, subnet := range vlan.Subnets {
			note := "-"
			// Subnets do not always have notes, this is required to avoid invalid memory address errors
			if subnet.Note != nil {
				note = *subnet.Note
			}
			subnetTable.AppendRow([]interface{}{*subnet.Id, *subnet.NetworkIdentifier, *subnet.SubnetType, note})
		}
		// Renders the subnet table to the main table instead of main output
		t.AppendRow([]interface{}{"Subnets", subnetTable.Render()})
	}
	// Prints out the table
	t.Render()
	return nil
}
