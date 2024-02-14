/**
 * Copyright 2016-2024 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
 * the License. You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed
 * on an "AS IS" BASIS,WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and limitations under the License.
 */

// AUTOMATICALLY GENERATED CODE - DO NOT MODIFY

package services

import (
	"fmt"
	"strings"

	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/sl"
)

// The SoftLayer_Account data type contains general information relating to a single SoftLayer customer account. Personal information in this type such as names, addresses, and phone numbers are assigned to the account only and not to users belonging to the account. The SoftLayer_Account data type contains a number of relational properties that are used by the SoftLayer customer portal to quickly present a variety of account related services to it's users.
//
// SoftLayer customers are unable to change their company account information in the portal or the API. If you need to change this information please open a sales ticket in our customer portal and our account management staff will assist you.
type Account struct {
	Session session.SLSession
	Options sl.Options
}

// GetAccountService returns an instance of the Account SoftLayer service
func GetAccountService(sess session.SLSession) Account {
	return Account{Session: sess}
}

func (r Account) Id(id int) Account {
	r.Options.Id = &id
	return r
}

func (r Account) Mask(mask string) Account {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Account) Filter(filter string) Account {
	r.Options.Filter = filter
	return r
}

func (r Account) Limit(limit int) Account {
	r.Options.Limit = &limit
	return r
}

func (r Account) Offset(offset int) Account {
	r.Options.Offset = &offset
	return r
}

// Retrieve The active account software licenses owned by an account
func (r Account) GetActiveAccountLicenses() (resp []datatypes.Software_AccountLicense, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getActiveAccountLicenses", nil, &r.Options, &resp)
	return
}

// Retrieve An account's non-expired quotes.
func (r Account) GetActiveQuotes() (resp []datatypes.Billing_Order_Quote, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getActiveQuotes", nil, &r.Options, &resp)
	return
}

// Retrieve The virtual software licenses controlled by an account
func (r Account) GetActiveVirtualLicenses() (resp []datatypes.Software_VirtualLicense, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getActiveVirtualLicenses", nil, &r.Options, &resp)
	return
}

// Retrieve All billing items of an account.
func (r Account) GetAllTopLevelBillingItems() (resp []datatypes.Billing_Item, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getAllTopLevelBillingItems", nil, &r.Options, &resp)
	return
}

// Retrieve An account's associated application delivery controller records.
func (r Account) GetApplicationDeliveryControllers() (resp []datatypes.Network_Application_Delivery_Controller, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getApplicationDeliveryControllers", nil, &r.Options, &resp)
	return
}

// Retrieve The bandwidth allotments for an account.
func (r Account) GetBandwidthAllotments() (resp []datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getBandwidthAllotments", nil, &r.Options, &resp)
	return
}

// Retrieve All closed tickets associated with an account.
func (r Account) GetClosedTickets() (resp []datatypes.Ticket, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getClosedTickets", nil, &r.Options, &resp)
	return
}

// Retrieve the user record of the user calling the SoftLayer API.
func (r Account) GetCurrentUser() (resp datatypes.User_Customer, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getCurrentUser", nil, &r.Options, &resp)
	return
}

// Retrieve An account's associated virtual dedicated host objects.
func (r Account) GetDedicatedHosts() (resp []datatypes.Virtual_DedicatedHost, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getDedicatedHosts", nil, &r.Options, &resp)
	return
}

// Retrieve The DNS domains associated with an account.
func (r Account) GetDomains() (resp []datatypes.Dns_Domain, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getDomains", nil, &r.Options, &resp)
	return
}

// Retrieve Stored security certificates that are expired (ie. SSL)
func (r Account) GetExpiredSecurityCertificates() (resp []datatypes.Security_Certificate, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getExpiredSecurityCertificates", nil, &r.Options, &resp)
	return
}

// Retrieve
func (r Account) GetGlobalIpRecords() (resp []datatypes.Network_Subnet_IpAddress_Global, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getGlobalIpRecords", nil, &r.Options, &resp)
	return
}

// Retrieve
func (r Account) GetGlobalIpv4Records() (resp []datatypes.Network_Subnet_IpAddress_Global, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getGlobalIpv4Records", nil, &r.Options, &resp)
	return
}

// Retrieve
func (r Account) GetGlobalIpv6Records() (resp []datatypes.Network_Subnet_IpAddress_Global, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getGlobalIpv6Records", nil, &r.Options, &resp)
	return
}

// Retrieve An account's associated hardware objects.
func (r Account) GetHardware() (resp []datatypes.Hardware, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getHardware", nil, &r.Options, &resp)
	return
}

// Retrieve An account's associated hourly virtual guest objects.
func (r Account) GetHourlyVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getHourlyVirtualGuests", nil, &r.Options, &resp)
	return
}

// Retrieve An account's associated Virtual Storage volumes.
func (r Account) GetHubNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getHubNetworkStorage", nil, &r.Options, &resp)
	return
}

// Retrieve An account's associated billing invoices.
func (r Account) GetInvoices() (resp []datatypes.Billing_Invoice, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getInvoices", nil, &r.Options, &resp)
	return
}

// Retrieve An account's associated iSCSI storage volumes.
func (r Account) GetIscsiNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getIscsiNetworkStorage", nil, &r.Options, &resp)
	return
}

// Retrieve An account's associated monthly virtual guest objects.
func (r Account) GetMonthlyVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getMonthlyVirtualGuests", nil, &r.Options, &resp)
	return
}

// Retrieve An account's associated NAS storage volumes.
func (r Account) GetNasNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getNasNetworkStorage", nil, &r.Options, &resp)
	return
}

// Retrieve All network gateway devices on this account.
func (r Account) GetNetworkGateways() (resp []datatypes.Network_Gateway, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getNetworkGateways", nil, &r.Options, &resp)
	return
}

// Retrieve
func (r Account) GetNetworkMessageDeliveryAccounts() (resp []datatypes.Network_Message_Delivery, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getNetworkMessageDeliveryAccounts", nil, &r.Options, &resp)
	return
}

// Retrieve An account's associated storage volumes. This includes Lockbox, NAS, EVault, and iSCSI volumes.
func (r Account) GetNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getNetworkStorage", nil, &r.Options, &resp)
	return
}

// Retrieve IPSec network tunnels for an account.
func (r Account) GetNetworkTunnelContexts() (resp []datatypes.Network_Tunnel_Module_Context, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getNetworkTunnelContexts", nil, &r.Options, &resp)
	return
}

// Retrieve All network VLANs assigned to an account.
func (r Account) GetNetworkVlans() (resp []datatypes.Network_Vlan, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getNetworkVlans", nil, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Account object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Account service. You can only retrieve the account that your portal user is assigned to.
func (r Account) GetObject() (resp datatypes.Account, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getObject", nil, &r.Options, &resp)
	return
}

// Retrieve All open tickets associated with an account.
func (r Account) GetOpenTickets() (resp []datatypes.Ticket, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getOpenTickets", nil, &r.Options, &resp)
	return
}

// Retrieve An account's associated virtual placement groups.
func (r Account) GetPlacementGroups() (resp []datatypes.Virtual_PlacementGroup, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getPlacementGroups", nil, &r.Options, &resp)
	return
}

// Retrieve
func (r Account) GetPortableStorageVolumes() (resp []datatypes.Virtual_Disk_Image, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getPortableStorageVolumes", nil, &r.Options, &resp)
	return
}

// Retrieve Customer specified URIs that are downloaded onto a newly provisioned or reloaded server. If the URI is sent over https it will be executed directly on the server.
func (r Account) GetPostProvisioningHooks() (resp []datatypes.Provisioning_Hook, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getPostProvisioningHooks", nil, &r.Options, &resp)
	return
}

// Retrieve Private and shared template group objects (parent only) for an account.
func (r Account) GetPrivateBlockDeviceTemplateGroups() (resp []datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getPrivateBlockDeviceTemplateGroups", nil, &r.Options, &resp)
	return
}

// Retrieve The private network VLANs assigned to an account.
func (r Account) GetPrivateNetworkVlans() (resp []datatypes.Network_Vlan, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getPrivateNetworkVlans", nil, &r.Options, &resp)
	return
}

// Retrieve The reserved capacity groups owned by this account.
func (r Account) GetReservedCapacityGroups() (resp []datatypes.Virtual_ReservedCapacityGroup, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getReservedCapacityGroups", nil, &r.Options, &resp)
	return
}

// Retrieve All Routers that an accounts VLANs reside on
func (r Account) GetRouters() (resp []datatypes.Hardware, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getRouters", nil, &r.Options, &resp)
	return
}

// Retrieve DEPRECATED
func (r Account) GetRwhoisData() (resp []datatypes.Network_Subnet_Rwhois_Data, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getRwhoisData", nil, &r.Options, &resp)
	return
}

// Retrieve Stored security certificates (ie. SSL)
func (r Account) GetSecurityCertificates() (resp []datatypes.Security_Certificate, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getSecurityCertificates", nil, &r.Options, &resp)
	return
}

// Retrieve Customer specified SSH keys that can be implemented onto a newly provisioned or reloaded server.
func (r Account) GetSshKeys() (resp []datatypes.Security_Ssh_Key, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getSshKeys", nil, &r.Options, &resp)
	return
}

// Retrieve All network subnets associated with an account.
func (r Account) GetSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getSubnets", nil, &r.Options, &resp)
	return
}

// Retrieve An account's associated tickets.
func (r Account) GetTickets() (resp []datatypes.Ticket, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getTickets", nil, &r.Options, &resp)
	return
}

// Retrieve An account's associated upgrade requests.
func (r Account) GetUpgradeRequests() (resp []datatypes.Product_Upgrade_Request, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getUpgradeRequests", nil, &r.Options, &resp)
	return
}

// Retrieve An account's portal users.
func (r Account) GetUsers() (resp []datatypes.User_Customer, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getUsers", nil, &r.Options, &resp)
	return
}

// Retrieve Stored security certificates that are not expired (ie. SSL)
func (r Account) GetValidSecurityCertificates() (resp []datatypes.Security_Certificate, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getValidSecurityCertificates", nil, &r.Options, &resp)
	return
}

// Retrieve The bandwidth pooling for this account.
func (r Account) GetVirtualDedicatedRacks() (resp []datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getVirtualDedicatedRacks", nil, &r.Options, &resp)
	return
}

// Retrieve An account's associated virtual guest objects.
func (r Account) GetVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = r.Session.DoRequest("SoftLayer_Account", "getVirtualGuests", nil, &r.Options, &resp)
	return
}

// Set the flag that enables or disables automatic private network VLAN spanning for a SoftLayer customer account. Enabling VLAN spanning allows an account's servers to talk on the same broadcast domain even if they reside within different private vlans.
func (r Account) SetVlanSpan(enabled *bool) (resp bool, err error) {
	params := []interface{}{
		enabled,
	}
	err = r.Session.DoRequest("SoftLayer_Account", "setVlanSpan", params, &r.Options, &resp)
	return
}
