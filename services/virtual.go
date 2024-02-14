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

// This data type presents the structure for a dedicated host. The data type contains relational properties to distinguish a dedicated host and associate an account to it.
type Virtual_DedicatedHost struct {
	Session session.SLSession
	Options sl.Options
}

// GetVirtualDedicatedHostService returns an instance of the Virtual_DedicatedHost SoftLayer service
func GetVirtualDedicatedHostService(sess session.SLSession) Virtual_DedicatedHost {
	return Virtual_DedicatedHost{Session: sess}
}

func (r Virtual_DedicatedHost) Id(id int) Virtual_DedicatedHost {
	r.Options.Id = &id
	return r
}

func (r Virtual_DedicatedHost) Mask(mask string) Virtual_DedicatedHost {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Virtual_DedicatedHost) Filter(filter string) Virtual_DedicatedHost {
	r.Options.Filter = filter
	return r
}

func (r Virtual_DedicatedHost) Limit(limit int) Virtual_DedicatedHost {
	r.Options.Limit = &limit
	return r
}

func (r Virtual_DedicatedHost) Offset(offset int) Virtual_DedicatedHost {
	r.Options.Offset = &offset
	return r
}

// Retrieve The guests associated with the dedicated host.
func (r Virtual_DedicatedHost) GetGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_DedicatedHost", "getGuests", nil, &r.Options, &resp)
	return
}

// no documentation yet
func (r Virtual_DedicatedHost) GetObject() (resp datatypes.Virtual_DedicatedHost, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_DedicatedHost", "getObject", nil, &r.Options, &resp)
	return
}

// The virtual guest data type presents the structure in which all virtual guests will be presented. Internally, the structure supports various virtualization platforms with no change to external interaction.
//
// A guest, also known as a virtual server, represents an allocation of resources on a virtual host.
type Virtual_Guest struct {
	Session session.SLSession
	Options sl.Options
}

// GetVirtualGuestService returns an instance of the Virtual_Guest SoftLayer service
func GetVirtualGuestService(sess session.SLSession) Virtual_Guest {
	return Virtual_Guest{Session: sess}
}

func (r Virtual_Guest) Id(id int) Virtual_Guest {
	r.Options.Id = &id
	return r
}

func (r Virtual_Guest) Mask(mask string) Virtual_Guest {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Virtual_Guest) Filter(filter string) Virtual_Guest {
	r.Options.Filter = filter
	return r
}

func (r Virtual_Guest) Limit(limit int) Virtual_Guest {
	r.Options.Limit = &limit
	return r
}

func (r Virtual_Guest) Offset(offset int) Virtual_Guest {
	r.Options.Offset = &offset
	return r
}

// This method is used to allow access to multiple SoftLayer_Network_Storage volumes that support host- or network-level access control.
func (r Virtual_Guest) AllowAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "allowAccessToNetworkStorageList", params, &r.Options, &resp)
	return
}

// Creates a transaction to attach a guest's disk image. If the disk image is already attached it will be ignored.
//
// WARNING: SoftLayer_Virtual_Guest::checkHostDiskAvailability should be called before this method. If the SoftLayer_Virtual_Guest::checkHostDiskAvailability method is not called before this method, the guest migration will happen automatically.
func (r Virtual_Guest) AttachDiskImage(imageId *int) (resp datatypes.Provisioning_Version1_Transaction, err error) {
	params := []interface{}{
		imageId,
	}
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "attachDiskImage", params, &r.Options, &resp)
	return
}

// Create a transaction to archive a computing instance's block devices
func (r Virtual_Guest) CreateArchiveTemplate(groupName *string, blockDevices []datatypes.Virtual_Guest_Block_Device, note *string) (resp datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	params := []interface{}{
		groupName,
		blockDevices,
		note,
	}
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "createArchiveTemplate", params, &r.Options, &resp)
	return
}

// Create a transaction to archive a computing instance's block devices
// Deprecated: This function has been marked as deprecated.
func (r Virtual_Guest) CreateArchiveTransaction(groupName *string, blockDevices []datatypes.Virtual_Guest_Block_Device, note *string) (resp datatypes.Provisioning_Version1_Transaction, err error) {
	params := []interface{}{
		groupName,
		blockDevices,
		note,
	}
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "createArchiveTransaction", params, &r.Options, &resp)
	return
}

// createObject() enables the creation of computing instances on an account. This method is a simplified alternative to interacting with the ordering system directly.
//
// In order to create a computing instance, a template object must be sent in with a few required values.
//
// When this method returns an order will have been placed for a computing instance of the specified configuration.
//
// To determine when the instance is available you can poll the instance via [[SoftLayer_Virtual_Guest/getObject]], with an object mask requesting the `provisionDate` relational property. When `provisionDate` is not `null`, the instance will be ready.
//
// > **Warning:** Computing instances created via this method will incur charges on your account. For testing input parameters see [[SoftLayer_Virtual_Guest/generateOrderTemplate]].
//
// ### Required Input [[SoftLayer_Virtual_Guest]]
//
// - `Hostname`  String **Required**
//   - Hostname for the computing instance.
//
// - `Domain` String **Required**
//   - Domain for the computing instance.
//
// - `startCpus` Integer **Required**
//   - The number of CPU cores to allocate.
//   - See [[SoftLayer_Virtual_Guest/getCreateObjectOptions]] for available options.
//
// - `maxMemory` Integer **Required**
//   - The amount of memory to allocate in megabytes.
//   - See [[SoftLayer_Virtual_Guest/getCreateObjectOptions]] for available options.
//
// - `datacenter.name` *String* **Required**
//   - Specifies which datacenter the instance is to be provisioned in. Needs to be a nested object.
//   - Example: `"datacenter": {"name": "dal05"}`
//   - See [[SoftLayer_Virtual_Guest/getCreateObjectOptions]] for available options.
//
// - `hourlyBillingFlag` Boolean **Required**
//   - Specifies the billing type for the instance.
//   - True for hourly billing, False for monthly billing.
//
// - `localDiskFlag` Boolean **Required**
//   - Specifies the disk type for the instance.
//   - True for local to the instance disks, False for SAN disks.
//
// - `dedicatedAccountHostOnlyFlag` Boolean
//   - When true this flag specifies that a compute instance is to run on hosts that only have guests from the same account.
//   - Default: False
//
// - `operatingSystemReferenceCode` String **Conditionally required**
//   - An identifier for the operating system to provision the computing instance with.
//   - Not required when using a `blockDeviceTemplateGroup.globalIdentifier`, as the template will have its own operating system.
//   - See [[SoftLayer_Virtual_Guest/getCreateObjectOptions]] for available options.
//   - **Notice**: Some operating systems are billed based on the number of CPUs the guest has. The price which is used can be determined by calling
//     [[SoftLayer_Virtual_Guest/generateOrderTemplate]] with your desired device specifications.
//
// - `blockDeviceTemplateGroup.globalIdentifier` String
//   - The GUID for the template to be used to provision the computing instance.
//   - Conflicts with `operatingSystemReferenceCode`
//   - **Notice**: Some operating systems are billed based on the number of CPUs the guest has. The price which is used can be determined by calling
//     [[SoftLayer_Virtual_Guest/generateOrderTemplate]] with your desired device specifications.
//   - A list of public images may be obtained via a request to [[SoftLayer_Virtual_Guest_Block_Device_Template_Group/getPublicImages]]
//   - A list of private images may be obtained via a request to [[SoftLayer_Account/getPrivateBlockDeviceTemplateGroups]]
//   - Example: `"blockDeviceTemplateGroup": { globalIdentifier": "07beadaa-1e11-476e-a188-3f7795feb9fb"`
//
// - `networkComponents.maxSpeed` Integer
//   - Specifies the connection speed for the instance's network components.
//   - The `networkComponents` property is an array with a single [[SoftLayer_Virtual_Guest_Network_Component]] structure.
//     The `maxSpeed` property must be set to specify the network uplink speed, in megabits per second, of the computing instance.
//   - See [[SoftLayer_Virtual_Guest/getCreateObjectOptions]] for available options.
//   - Default: 10
//   - Example: `"networkComponents": [{"maxSpeed": 1000}]`
//
// - `privateNetworkOnlyFlag` Boolean
//   - When true this flag specifies that a compute instance is to only have access to the private network.
//   - Default: False
//
// - `primaryNetworkComponent.networkVlan.id` Integer
//   - Specifies the network vlan which is to be used for the frontend interface of the computing instance.
//   - The `primaryNetworkComponent` property is a [[SoftLayer_Virtual_Guest_Network_Component]] structure with the `networkVlan` property populated with a i
//     [[SoftLayer_Network_Vlan]] structure. The `id` property must be set to specify the frontend network vlan of the computing instance.
//   - *NOTE* This is the VLAN `id`, NOT the vlan number.
//   - Example: `"primaryNetworkComponent":{"networkVlan": {"id": 1234567}}`
//
// - `backendNetworkComponent.networkVlan.id` Integer
//   - Specifies the network vlan which is to be used for the backend interface of the computing instance.
//   - The `backendNetworkComponent` property is a [[SoftLayer_Virtual_Guest_Network_Component]] structure with the `networkVlan` property populated with a
//     [[SoftLayer_Network_Vlan]] structure. The `id` property must be set to specify the backend network vlan of the computing instance.
//   - *NOTE* This is the VLAN `id`, NOT the vlan number.
//   - Example: `"backendNetworkComponent":{"networkVlan": {"id": 1234567}}`
//
// - `primaryNetworkComponent.securityGroupBindings` [[SoftLayer_Virtual_Network_SecurityGroup_NetworkComponentBinding]][]
//   - Specifies the security groups to be attached to this VSI's frontend network adapter
//   - The `primaryNetworkComponent` property is a [[SoftLayer_Virtual_Guest_Network_Component]] structure with the `securityGroupBindings` property populated
//     with an array of [[SoftLayer_Virtual_Network_SecurityGroup_NetworkComponentBinding]] structures. The `securityGroup` property in each must be set to
//     specify the security group to be attached to the primary frontend network component.
//   - Example:
//     ```
//     "primaryNetworkComponent": {
//     "securityGroupBindings": [
//     {"securityGroup":{"id": 5555555}},
//     {"securityGroup":{"id": 1112223}},
//     ]
//     }
//     ```
//
// - `primaryBackendNetworkComponent.securityGroupBindings` [[SoftLayer_Virtual_Network_SecurityGroup_NetworkComponentBinding]][]
//   - Specifies the security groups to be attached to this VSI's backend network adapter
//   - The `primaryNetworkComponent` property is a [[SoftLayer_Virtual_Guest_Network_Component]] structure with the `securityGroupBindings` property populated
//     with an array of [[SoftLayer_Virtual_Network_SecurityGroup_NetworkComponentBinding]] structures. The `securityGroup` property in each must be set to
//     specify the security group to be attached to the primary frontend network component.
//   - Example:
//     ```
//     "primaryBackendNetworkComponent": {
//     "securityGroupBindings": [
//     {"securityGroup":{"id": 33322211}},
//     {"securityGroup":{"id": 77777222}},
//     ]
//     }
//     ```
//
// - `blockDevices` [[SoftLayer_Virtual_Guest_Block_Device]][]
//   - Block device and disk image settings for the computing instance
//   - The `blockDevices` property is an array of [[SoftLayer_Virtual_Guest_Block_Device]] structures. Each block device must specify the `device` property
//     along with the `diskImage`  property, which is a [[SoftLayer_Virtual_Disk_Image]] structure with the `capacity` property set. The `device` number `'1'`
//     is reserved for the SWAP disk attached to the computing instance.
//   - Default: The smallest available capacity for the primary disk will be used. If an image template is specified the disk capacity will be be provided by the template.
//   - Example:
//     ```
//     "blockDevices":[{"device": "0", "diskImage": {"capacity": 100}}],
//     "localDiskFlag": true
//     ```
//   - See [[SoftLayer_Virtual_Guest/getCreateObjectOptions]] for available options.
//
// - `userData.value`  String
//   - Arbitrary data to be made available to the computing instance.
//   - The `userData` property is an array with a single [[SoftLayer_Virtual_Guest_Attribute]] structure with the `value` property set to an arbitrary value.
//     This value can be retrieved via the [[SoftLayer_Resource_Metadata/getUserMetadata]] method from a request originating from the computing instance.
//     This is primarily useful for providing data to software that may be on the instance and configured to execute upon first boot.
//   - Example: `"userData":[{"value": "testData"}]`
//
// - `sshKeys` [[SoftLayer_Security_Ssh_Key]][]
//   - The `sshKeys` property is an array of [[SoftLayer_Security_Ssh_Key]] structures with the `id` property set to the value of an existing SSH key.
//   - To create a new SSH key, call [[SoftLayer_Security_Ssh_Key/createObject|createObject]].
//   - To obtain a list of existing SSH keys, call [[SoftLayer_Account/getSshKeys]]
//   - Example: `"sshKeys":[{"id": 1234567}]`
//
// - `postInstallScriptUri` String
//   - Specifies the uri location of the script to be downloaded and run after installation is complete. Only scripts from HTTPS servers are executed on startup.
//
// REST Example:
// ```
//
//	curl -X POST -d '{
//	    "parameters":[
//	        {
//	            "hostname": "host1",
//	            "domain": "example.com",
//	            "startCpus": 1,
//	            "maxMemory": 1024,
//	            "hourlyBillingFlag": true,
//	            "localDiskFlag": true,
//	            "operatingSystemReferenceCode": "UBUNTU_LATEST"
//	        }
//	}' https://api.softlayer.com/rest/v3.1/SoftLayer_Virtual_Guest/createObject.json
//
// HTTP/1.1 201 Created
// Location: https://api.softlayer.com/rest/v3.1/SoftLayer_Virtual_Guest/1301396/getObject
//
//	{
//	  "accountId": 232298,
//	  "createDate": "2012-11-30T16:28:17-06:00",
//	  "dedicatedAccountHostOnlyFlag": false,
//	  "domain": "example.com",
//	  "hostname": "host1",
//	  "id": 1301396,
//	  "lastPowerStateId": null,
//	  "lastVerifiedDate": null,
//	  "maxCpu": 1,
//	  "maxCpuUnits": "CORE",
//	  "maxMemory": 1024,
//	  "metricPollDate": null,
//	  "modifyDate": null,
//	  "privateNetworkOnlyFlag": false,
//	  "startCpus": 1,
//	  "statusId": 1001,
//	  "globalIdentifier": "2d203774-0ee1-49f5-9599-6ef67358dd31"
//	}
//
// ```
func (r Virtual_Guest) CreateObject(templateObject *datatypes.Virtual_Guest) (resp datatypes.Virtual_Guest, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "createObject", params, &r.Options, &resp)
	return
}

// createObjects() enables the creation of multiple computing instances on an account in a single call. This
// method is a simplified alternative to interacting with the ordering system directly.
//
// In order to create a computing instance a set of template objects must be sent in with a few required
// values.
//
// <b>Warning:</b> Computing instances created via this method will incur charges on your account.
//
// See [[SoftLayer_Virtual_Guest/createObject|createObject]] for specifics on the requirements of each template object.
//
// <h1>Example</h1>
//
//	<http title="Request">curl -X POST -d '{
//	 "parameters":[
//	     [
//	         {
//	             "hostname": "host1",
//	             "domain": "example.com",
//	             "startCpus": 1,
//	             "maxMemory": 1024,
//	             "hourlyBillingFlag": true,
//	             "localDiskFlag": true,
//	             "operatingSystemReferenceCode": "UBUNTU_LATEST"
//	         },
//	         {
//	             "hostname": "host2",
//	             "domain": "example.com",
//	             "startCpus": 1,
//	             "maxMemory": 1024,
//	             "hourlyBillingFlag": true,
//	             "localDiskFlag": true,
//	             "operatingSystemReferenceCode": "UBUNTU_LATEST"
//	         }
//	     ]
//	 ]
//	}' https://api.softlayer.com/rest/v3/SoftLayer_Virtual_Guest/createObjects.json
//
// </http>
// <http title="Response">HTTP/1.1 200 OK
//
// [
//
//	{
//	    "accountId": 232298,
//	    "createDate": "2012-11-30T23:56:48-06:00",
//	    "dedicatedAccountHostOnlyFlag": false,
//	    "domain": "softlayer.com",
//	    "hostname": "ubuntu1",
//	    "id": 1301456,
//	    "lastPowerStateId": null,
//	    "lastVerifiedDate": null,
//	    "maxCpu": 1,
//	    "maxCpuUnits": "CORE",
//	    "maxMemory": 1024,
//	    "metricPollDate": null,
//	    "modifyDate": null,
//	    "privateNetworkOnlyFlag": false,
//	    "startCpus": 1,
//	    "statusId": 1001,
//	    "globalIdentifier": "fed4c822-48c0-45d0-85e2-90476aa0c542"
//	},
//	{
//	    "accountId": 232298,
//	    "createDate": "2012-11-30T23:56:49-06:00",
//	    "dedicatedAccountHostOnlyFlag": false,
//	    "domain": "softlayer.com",
//	    "hostname": "ubuntu2",
//	    "id": 1301457,
//	    "lastPowerStateId": null,
//	    "lastVerifiedDate": null,
//	    "maxCpu": 1,
//	    "maxCpuUnits": "CORE",
//	    "maxMemory": 1024,
//	    "metricPollDate": null,
//	    "modifyDate": null,
//	    "privateNetworkOnlyFlag": false,
//	    "startCpus": 1,
//	    "statusId": 1001,
//	    "globalIdentifier": "bed4c686-9562-4ade-9049-dc4d5b6b200c"
//	}
//
// ]
// </http>
func (r Virtual_Guest) CreateObjects(templateObjects []datatypes.Virtual_Guest) (resp []datatypes.Virtual_Guest, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "createObjects", params, &r.Options, &resp)
	return
}

// This method will cancel a computing instance effective immediately. For instances billed hourly, the charges will stop immediately after the method returns.
func (r Virtual_Guest) DeleteObject() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "deleteObject", nil, &r.Options, &resp)
	return
}

// Edit a computing instance's properties
func (r Virtual_Guest) EditObject(templateObject *datatypes.Virtual_Guest) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "editObject", params, &r.Options, &resp)
	return
}

// Reboot a Linux guest into the Xen rescue image.
func (r Virtual_Guest) ExecuteRescueLayer() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "executeRescueLayer", nil, &r.Options, &resp)
	return
}

// Obtain an [[SoftLayer_Container_Product_Order_Virtual_Guest (type)|order container]] that can be sent to [[SoftLayer_Product_Order/verifyOrder|verifyOrder]] or [[SoftLayer_Product_Order/placeOrder|placeOrder]].
//
// This is primarily useful when there is a necessity to confirm the price which will be charged for an order.
//
// See [[SoftLayer_Virtual_Guest/createObject|createObject]] for specifics on the requirements of the template object parameter.
func (r Virtual_Guest) GenerateOrderTemplate(templateObject *datatypes.Virtual_Guest) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "generateOrderTemplate", params, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Storage_Allowed_Host information to connect this Virtual Guest to Network Storage volumes that require access control lists.
func (r Virtual_Guest) GetAllowedHost() (resp datatypes.Network_Storage_Allowed_Host, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "getAllowedHost", nil, &r.Options, &resp)
	return
}

// This method is retrieve a list of SoftLayer_Network_Storage volumes that are authorized access to this SoftLayer_Virtual_Guest.
func (r Virtual_Guest) GetAttachedNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "getAttachedNetworkStorages", params, &r.Options, &resp)
	return
}

// Retrieve A computing instance's block devices. Block devices link [[SoftLayer_Virtual_Disk_Image|disk images]] to computing instances.
func (r Virtual_Guest) GetBlockDevices() (resp []datatypes.Virtual_Guest_Block_Device, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "getBlockDevices", nil, &r.Options, &resp)
	return
}

// There are many options that may be provided while ordering a computing instance, this method can be used to determine what these options are.
//
// Detailed information on the return value can be found on the data type page for [[SoftLayer_Container_Virtual_Guest_Configuration (type)]].
func (r Virtual_Guest) GetCreateObjectOptions() (resp datatypes.Container_Virtual_Guest_Configuration, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "getCreateObjectOptions", nil, &r.Options, &resp)
	return
}

// Retrieve A guest's metric tracking object.
func (r Virtual_Guest) GetMetricTrackingObject() (resp datatypes.Metric_Tracking_Object, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "getMetricTrackingObject", nil, &r.Options, &resp)
	return
}

// Retrieve The metric tracking object id for this guest.
func (r Virtual_Guest) GetMetricTrackingObjectId() (resp int, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "getMetricTrackingObjectId", nil, &r.Options, &resp)
	return
}

// no documentation yet
func (r Virtual_Guest) GetObject() (resp datatypes.Virtual_Guest, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "getObject", nil, &r.Options, &resp)
	return
}

// Retrieve the reverse domain records associated with this server.
func (r Virtual_Guest) GetReverseDomainRecords() (resp []datatypes.Dns_Domain, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "getReverseDomainRecords", nil, &r.Options, &resp)
	return
}

// Retrieves a list of all upgrades available to a virtual server. Upgradeable items include, but are not limited to, number of cores, amount of RAM, storage configuration, and network port speed.
//
// This method exclude downgrade item prices by default. You can set the "includeDowngradeItemPrices" parameter to true so that it can include downgrade item prices.
func (r Virtual_Guest) GetUpgradeItemPrices(includeDowngradeItemPrices *bool) (resp []datatypes.Product_Item_Price, err error) {
	params := []interface{}{
		includeDowngradeItemPrices,
	}
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "getUpgradeItemPrices", params, &r.Options, &resp)
	return
}

// Issues a ping command and returns the success (true) or failure (false) of the ping command.
func (r Virtual_Guest) IsPingable() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "isPingable", nil, &r.Options, &resp)
	return
}

// Creates a transaction to migrate a virtual guest to a new host. NOTE: Will only migrate if SoftLayer_Virtual_Guest property pendingMigrationFlag = true
func (r Virtual_Guest) Migrate() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "migrate", nil, &r.Options, &resp)
	return
}

// Create a transaction to migrate an instance from one dedicated host to another dedicated host
func (r Virtual_Guest) MigrateDedicatedHost(destinationHostId *int) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		destinationHostId,
	}
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "migrateDedicatedHost", params, &r.Options, &resp)
	return
}

// Pause a virtual guest. This can only be called when the specified VM is in the Running state.
func (r Virtual_Guest) Pause() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "pause", nil, &r.Options, &resp)
	return
}

// Power off a virtual guest
func (r Virtual_Guest) PowerOff() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "powerOff", nil, &r.Options, &resp)
	return
}

// Power off a virtual guest
func (r Virtual_Guest) PowerOffSoft() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "powerOffSoft", nil, &r.Options, &resp)
	return
}

// Power on a virtual guest
func (r Virtual_Guest) PowerOn() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "powerOn", nil, &r.Options, &resp)
	return
}

// Power cycle a virtual guest
func (r Virtual_Guest) RebootDefault() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "rebootDefault", nil, &r.Options, &resp)
	return
}

// Power cycle a guest.
func (r Virtual_Guest) RebootHard() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "rebootHard", nil, &r.Options, &resp)
	return
}

// Attempt to complete a soft reboot of a guest by shutting down the operating system.
func (r Virtual_Guest) RebootSoft() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "rebootSoft", nil, &r.Options, &resp)
	return
}

// Reloads current operating system configuration.
//
// This service has a confirmation protocol for proceeding with the reload. To proceed with the reload without confirmation, simply pass in 'FORCE' as the token parameter. To proceed with the reload with confirmation, simply call the service with no parameter. A token string will be returned by this service. The token will remain active for 10 minutes. Use this token as the parameter to confirm that a reload is to be performed for the server.
//
// As a precaution, we strongly  recommend backing up all data before reloading the operating system. The reload will format the primary disk and will reconfigure the computing instance to the current specifications on record.
//
// If reloading from an image template, we recommend first getting the list of valid private block device template groups, by calling the getOperatingSystemReloadImages method.
func (r Virtual_Guest) ReloadOperatingSystem(token *string, config *datatypes.Container_Hardware_Server_Configuration) (resp string, err error) {
	params := []interface{}{
		token,
		config,
	}
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "reloadOperatingSystem", params, &r.Options, &resp)
	return
}

// Resume a virtual guest, this can only be called when a VSI is in Suspended state.
func (r Virtual_Guest) Resume() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "resume", nil, &r.Options, &resp)
	return
}

// Sets the private network interface speed to the new speed. Speed values can only be 0 (Disconnect), 10, 100, or 1000. The new speed must be equal to or less than the max speed of the interface.
//
// It will take less than a minute to update the port speed.
func (r Virtual_Guest) SetPrivateNetworkInterfaceSpeed(newSpeed *int) (resp bool, err error) {
	params := []interface{}{
		newSpeed,
	}
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "setPrivateNetworkInterfaceSpeed", params, &r.Options, &resp)
	return
}

// Sets the public network interface speed to the new speed. Speed values can only be 0 (Disconnect), 10, 100, or 1000. The new speed must be equal to or less than the max speed of the interface.
//
// It will take less than a minute to update the port speed.
func (r Virtual_Guest) SetPublicNetworkInterfaceSpeed(newSpeed *int) (resp bool, err error) {
	params := []interface{}{
		newSpeed,
	}
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "setPublicNetworkInterfaceSpeed", params, &r.Options, &resp)
	return
}

// no documentation yet
func (r Virtual_Guest) SetTags(tags *string) (resp bool, err error) {
	params := []interface{}{
		tags,
	}
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "setTags", params, &r.Options, &resp)
	return
}

// Sets the data that will be written to the configuration drive.
func (r Virtual_Guest) SetUserMetadata(metadata []string) (resp bool, err error) {
	params := []interface{}{
		metadata,
	}
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest", "setUserMetadata", params, &r.Options, &resp)
	return
}

// The virtual block device template group data type presents the structure in which a group of archived image templates will be presented. The structure consists of a parent template group which contain multiple child template group objects.  Each child template group object represents the image template in a particular location. Unless editing/deleting a specific child template group object, it is best to use the parent object.
//
// A virtual block device template group, also known as an image template group, represents an image of a virtual guest instance.
type Virtual_Guest_Block_Device_Template_Group struct {
	Session session.SLSession
	Options sl.Options
}

// GetVirtualGuestBlockDeviceTemplateGroupService returns an instance of the Virtual_Guest_Block_Device_Template_Group SoftLayer service
func GetVirtualGuestBlockDeviceTemplateGroupService(sess session.SLSession) Virtual_Guest_Block_Device_Template_Group {
	return Virtual_Guest_Block_Device_Template_Group{Session: sess}
}

func (r Virtual_Guest_Block_Device_Template_Group) Id(id int) Virtual_Guest_Block_Device_Template_Group {
	r.Options.Id = &id
	return r
}

func (r Virtual_Guest_Block_Device_Template_Group) Mask(mask string) Virtual_Guest_Block_Device_Template_Group {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Virtual_Guest_Block_Device_Template_Group) Filter(filter string) Virtual_Guest_Block_Device_Template_Group {
	r.Options.Filter = filter
	return r
}

func (r Virtual_Guest_Block_Device_Template_Group) Limit(limit int) Virtual_Guest_Block_Device_Template_Group {
	r.Options.Limit = &limit
	return r
}

func (r Virtual_Guest_Block_Device_Template_Group) Offset(offset int) Virtual_Guest_Block_Device_Template_Group {
	r.Options.Offset = &offset
	return r
}

// This method will create transaction(s) to add available locations to an archive image template.
func (r Virtual_Guest_Block_Device_Template_Group) AddLocations(locations []datatypes.Location) (resp bool, err error) {
	params := []interface{}{
		locations,
	}
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest_Block_Device_Template_Group", "addLocations", params, &r.Options, &resp)
	return
}

// Create a transaction to export/copy a template to an ICOS.
func (r Virtual_Guest_Block_Device_Template_Group) CopyToIcos(configuration *datatypes.Container_Virtual_Guest_Block_Device_Template_Configuration) (resp bool, err error) {
	params := []interface{}{
		configuration,
	}
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest_Block_Device_Template_Group", "copyToIcos", params, &r.Options, &resp)
	return
}

// Create a process to import a disk image from ICOS and create a standard
func (r Virtual_Guest_Block_Device_Template_Group) CreateFromIcos(configuration *datatypes.Container_Virtual_Guest_Block_Device_Template_Configuration) (resp datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	params := []interface{}{
		configuration,
	}
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest_Block_Device_Template_Group", "createFromIcos", params, &r.Options, &resp)
	return
}

// Deleting a block device template group is different from the deletion of other objects.  A block device template group can contain several gigabytes of data in its disk images.  This may take some time to delete and requires a transaction to be created.  This method creates a transaction that will delete all resources associated with the block device template group.
func (r Virtual_Guest_Block_Device_Template_Group) DeleteObject() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest_Block_Device_Template_Group", "deleteObject", nil, &r.Options, &resp)
	return
}

// This method will deny another SoftLayer customer account's previously given access to provision CloudLayer Computing Instances from an image template group. Template access should only be removed from the parent template group object, not the child.
func (r Virtual_Guest_Block_Device_Template_Group) DenySharingAccess(accountId *int) (resp bool, err error) {
	params := []interface{}{
		accountId,
	}
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest_Block_Device_Template_Group", "denySharingAccess", params, &r.Options, &resp)
	return
}

// Edit an image template group's associated name and note. All other properties in the SoftLayer_Virtual_Guest_Block_Device_Template_Group data type are read-only.
func (r Virtual_Guest_Block_Device_Template_Group) EditObject(templateObject *datatypes.Virtual_Guest_Block_Device_Template_Group) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest_Block_Device_Template_Group", "editObject", params, &r.Options, &resp)
	return
}

// no documentation yet
func (r Virtual_Guest_Block_Device_Template_Group) GetObject() (resp datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest_Block_Device_Template_Group", "getObject", nil, &r.Options, &resp)
	return
}

// This method gets all public image templates that the user is allowed to see.
func (r Virtual_Guest_Block_Device_Template_Group) GetPublicImages() (resp []datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest_Block_Device_Template_Group", "getPublicImages", nil, &r.Options, &resp)
	return
}

// Returns the image storage locations.
func (r Virtual_Guest_Block_Device_Template_Group) GetStorageLocations() (resp []datatypes.Location, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest_Block_Device_Template_Group", "getStorageLocations", nil, &r.Options, &resp)
	return
}

// This method will permit another SoftLayer customer account access to provision CloudLayer Computing Instances from an image template group. Template access should only be given to the parent template group object, not the child.
func (r Virtual_Guest_Block_Device_Template_Group) PermitSharingAccess(accountId *int) (resp bool, err error) {
	params := []interface{}{
		accountId,
	}
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest_Block_Device_Template_Group", "permitSharingAccess", params, &r.Options, &resp)
	return
}

// This method will create transaction(s) to remove available locations from an archive image template.
func (r Virtual_Guest_Block_Device_Template_Group) RemoveLocations(locations []datatypes.Location) (resp bool, err error) {
	params := []interface{}{
		locations,
	}
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest_Block_Device_Template_Group", "removeLocations", params, &r.Options, &resp)
	return
}

// Set the tags for this template group.
func (r Virtual_Guest_Block_Device_Template_Group) SetTags(tags *string) (resp bool, err error) {
	params := []interface{}{
		tags,
	}
	err = r.Session.DoRequest("SoftLayer_Virtual_Guest_Block_Device_Template_Group", "setTags", params, &r.Options, &resp)
	return
}

// This data type presents the structure for a virtual guest placement group. The data type contains relational properties to the virtual guest placement group rule class.
type Virtual_PlacementGroup struct {
	Session session.SLSession
	Options sl.Options
}

// GetVirtualPlacementGroupService returns an instance of the Virtual_PlacementGroup SoftLayer service
func GetVirtualPlacementGroupService(sess session.SLSession) Virtual_PlacementGroup {
	return Virtual_PlacementGroup{Session: sess}
}

func (r Virtual_PlacementGroup) Id(id int) Virtual_PlacementGroup {
	r.Options.Id = &id
	return r
}

func (r Virtual_PlacementGroup) Mask(mask string) Virtual_PlacementGroup {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Virtual_PlacementGroup) Filter(filter string) Virtual_PlacementGroup {
	r.Options.Filter = filter
	return r
}

func (r Virtual_PlacementGroup) Limit(limit int) Virtual_PlacementGroup {
	r.Options.Limit = &limit
	return r
}

func (r Virtual_PlacementGroup) Offset(offset int) Virtual_PlacementGroup {
	r.Options.Offset = &offset
	return r
}

// Add a placement group to your account for use during VSI provisioning.
func (r Virtual_PlacementGroup) CreateObject(templateObject *datatypes.Virtual_PlacementGroup) (resp datatypes.Virtual_PlacementGroup, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Virtual_PlacementGroup", "createObject", params, &r.Options, &resp)
	return
}

// Delete a placement group from your account.
func (r Virtual_PlacementGroup) DeleteObject() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_PlacementGroup", "deleteObject", nil, &r.Options, &resp)
	return
}

// Returns all routers available for use with placement groups. If a datacenter location ID is provided, this method will further restrict the list of routers to ones contained within that datacenter.
func (r Virtual_PlacementGroup) GetAvailableRouters(datacenterId *int) (resp []datatypes.Hardware, err error) {
	params := []interface{}{
		datacenterId,
	}
	err = r.Session.DoRequest("SoftLayer_Virtual_PlacementGroup", "getAvailableRouters", params, &r.Options, &resp)
	return
}

// no documentation yet
func (r Virtual_PlacementGroup) GetObject() (resp datatypes.Virtual_PlacementGroup, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_PlacementGroup", "getObject", nil, &r.Options, &resp)
	return
}

// This data type presents the structure of a virtual guest placement group rule.
type Virtual_PlacementGroup_Rule struct {
	Session session.SLSession
	Options sl.Options
}

// GetVirtualPlacementGroupRuleService returns an instance of the Virtual_PlacementGroup_Rule SoftLayer service
func GetVirtualPlacementGroupRuleService(sess session.SLSession) Virtual_PlacementGroup_Rule {
	return Virtual_PlacementGroup_Rule{Session: sess}
}

func (r Virtual_PlacementGroup_Rule) Id(id int) Virtual_PlacementGroup_Rule {
	r.Options.Id = &id
	return r
}

func (r Virtual_PlacementGroup_Rule) Mask(mask string) Virtual_PlacementGroup_Rule {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Virtual_PlacementGroup_Rule) Filter(filter string) Virtual_PlacementGroup_Rule {
	r.Options.Filter = filter
	return r
}

func (r Virtual_PlacementGroup_Rule) Limit(limit int) Virtual_PlacementGroup_Rule {
	r.Options.Limit = &limit
	return r
}

func (r Virtual_PlacementGroup_Rule) Offset(offset int) Virtual_PlacementGroup_Rule {
	r.Options.Offset = &offset
	return r
}

// Get all placement group rules.
func (r Virtual_PlacementGroup_Rule) GetAllObjects() (resp []datatypes.Virtual_PlacementGroup_Rule, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_PlacementGroup_Rule", "getAllObjects", nil, &r.Options, &resp)
	return
}

// This data type presents the structure for a virtual reserved capacity group.
type Virtual_ReservedCapacityGroup struct {
	Session session.SLSession
	Options sl.Options
}

// GetVirtualReservedCapacityGroupService returns an instance of the Virtual_ReservedCapacityGroup SoftLayer service
func GetVirtualReservedCapacityGroupService(sess session.SLSession) Virtual_ReservedCapacityGroup {
	return Virtual_ReservedCapacityGroup{Session: sess}
}

func (r Virtual_ReservedCapacityGroup) Id(id int) Virtual_ReservedCapacityGroup {
	r.Options.Id = &id
	return r
}

func (r Virtual_ReservedCapacityGroup) Mask(mask string) Virtual_ReservedCapacityGroup {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Virtual_ReservedCapacityGroup) Filter(filter string) Virtual_ReservedCapacityGroup {
	r.Options.Filter = filter
	return r
}

func (r Virtual_ReservedCapacityGroup) Limit(limit int) Virtual_ReservedCapacityGroup {
	r.Options.Limit = &limit
	return r
}

func (r Virtual_ReservedCapacityGroup) Offset(offset int) Virtual_ReservedCapacityGroup {
	r.Options.Offset = &offset
	return r
}

// no documentation yet
func (r Virtual_ReservedCapacityGroup) GetObject() (resp datatypes.Virtual_ReservedCapacityGroup, err error) {
	err = r.Session.DoRequest("SoftLayer_Virtual_ReservedCapacityGroup", "getObject", nil, &r.Options, &resp)
	return
}
