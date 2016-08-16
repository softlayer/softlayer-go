/**
 * Copyright 2016 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/**
 * AUTOMATICALLY GENERATED CODE - DO NOT MODIFY
 */

package services

import (
	"fmt"
	"strings"

	"github.ibm.com/riethm/gopherlayer.git/datatypes"
	"github.ibm.com/riethm/gopherlayer.git/session"
	"github.ibm.com/riethm/gopherlayer.git/sl"
)

// The SoftLayer_Hardware data type contains general information relating to a single SoftLayer hardware.
type Hardware struct {
	Session *session.Session
	Options sl.Options
}

func GetHardwareService(sess *session.Session) Hardware {
	return Hardware{Session: sess}
}

func (r Hardware) Id(id int) Hardware {
	r.Options.Id = &id
	return r
}

func (r Hardware) Mask(mask string) Hardware {
	if !strings.HasPrefix(mask, "mask[") && strings.Contains(mask, "[") {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Hardware) Filter(filter string) Hardware {
	r.Options.Filter = filter
	return r
}

func (r Hardware) Limit(limit int) Hardware {
	r.Options.Limit = &limit
	return r
}

func (r Hardware) Offset(offset int) Hardware {
	r.Options.Offset = &offset
	return r
}

// This method is used to allow access to a SoftLayer_Network_Storage volume that supports host- or network-level access control.
func (r Hardware) AllowAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod("SoftLayer_Hardware", "allowAccessToNetworkStorage", params, r.Session, &r.Options, &resp)
	return
}

// This method is used to allow access to multiple SoftLayer_Network_Storage volumes that support host- or network-level access control.
func (r Hardware) AllowAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod("SoftLayer_Hardware", "allowAccessToNetworkStorageList", params, r.Session, &r.Options, &resp)
	return
}

// Captures a Flex Image of the hard disk on the physical machine, based on the capture template parameter. Returns the image template group containing the disk image.
func (r Hardware) CaptureImage(captureTemplate *datatypes.Container_Disk_Image_Capture_Template) (resp datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	params := []interface{}{
		captureTemplate,
	}
	err = invokeMethod("SoftLayer_Hardware", "captureImage", params, r.Session, &r.Options, &resp)
	return
}

// Returns monitoring alarm detailed history
func (r Hardware) CloseAlarm(alarmId *string) (resp bool, err error) {
	params := []interface{}{
		alarmId,
	}
	err = invokeMethod("SoftLayer_Hardware", "closeAlarm", params, r.Session, &r.Options, &resp)
	return
}

//
// <style type="text/css">.create_object > li > div { padding-top: .5em; padding-bottom: .5em}</style>
// createObject() enables the creation of servers on an account. This
// method is a simplified alternative to interacting with the ordering system directly.
//
//
// In order to create a server, a template object must be sent in with a few required
// values.
//
//
// When this method returns an order will have been placed for a server of the specified configuration.
//
//
// To determine when the server is available you can poll the server via [[SoftLayer_Hardware/getObject|getObject]],
// checking the <code>provisionDate</code> property.
// When <code>provisionDate</code> is not null, the server will be ready. Be sure to use the <code>globalIdentifier</code>
// as your initialization parameter.
//
//
// <b>Warning:</b> Servers created via this method will incur charges on your account. For testing input parameters see [[SoftLayer_Hardware/generateOrderTemplate|generateOrderTemplate]].
//
//
// <b>Input</b> - [[SoftLayer_Hardware (type)|SoftLayer_Hardware]]
// <ul class="create_object">
//     <li><code>hostname</code>
//         <div>Hostname for the server.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - string</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>domain</code>
//         <div>Domain for the server.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - string</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>processorCoreAmount</code>
//         <div>The number of logical CPU cores to allocate.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - int</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>memoryCapacity</code>
//         <div>The amount of memory to allocate in gigabytes.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - int</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>hourlyBillingFlag</code>
//         <div>Specifies the billing type for the server.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - boolean</li>
//             <li>When true the server will be billed on hourly usage, otherwise it will be billed on a monthly basis.</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>operatingSystemReferenceCode</code>
//         <div>An identifier for the operating system to provision the server with.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - string</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>datacenter.name</code>
//         <div>Specifies which datacenter the server is to be provisioned in.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - string</li>
//             <li>The <code>datacenter</code> property is a [[SoftLayer_Location (type)|location]] structure with the <code>name</code> field set.</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//         <http title="Example">{
//     "datacenter": {
//         "name": "dal05"
//     }
// }</http>
//         <br />
//     </li>
//     <li><code>networkComponents.maxSpeed</code>
//         <div>Specifies the connection speed for the server's network components.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - int</li>
//             <li><b>Default</b> - The highest available zero cost port speed will be used.</li>
//             <li><b>Description</b> - The <code>networkComponents</code> property is an array with a single [[SoftLayer_Network_Component (type)|network component]] structure. The <code>maxSpeed</code> property must be set to specify the network uplink speed, in megabits per second, of the server.</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//             <http title="Example">{
//     "networkComponents": [
//         {
//             "maxSpeed": 1000
//         }
//     ]
// }</http>
//         <br />
//     </li>
//     <li><code>networkComponents.redundancyEnabledFlag</code>
//         <div>Specifies whether or not the server's network components should be in redundancy groups.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - bool</li>
//             <li><b>Default</b> - <code>false</code></li>
//             <li><b>Description</b> - The <code>networkComponents</code> property is an array with a single [[SoftLayer_Network_Component (type)|network component]] structure. When the <code>redundancyEnabledFlag</code> property is true the server's network components will be in redundancy groups.</li>
//         </ul>
//             <http title="Example">{
//     "networkComponents": [
//         {
//             "redundancyEnabledFlag": false
//         }
//     ]
// }</http>
//         <br />
//     </li>
//     <li><code>privateNetworkOnlyFlag</code>
//         <div>Specifies whether or not the server only has access to the private network</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - boolean</li>
//             <li><b>Default</b> - <code>false</code></li>
//             <li>When true this flag specifies that a server is to only have access to the private network.</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>primaryNetworkComponent.networkVlan.id</code>
//         <div>Specifies the network vlan which is to be used for the frontend interface of the server.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - int</li>
//             <li><b>Description</b> - The <code>primaryNetworkComponent</code> property is a [[SoftLayer_Network_Component (type)|network component]] structure with the <code>networkVlan</code> property populated with a [[SoftLayer_Network_Vlan (type)|vlan]] structure. The <code>id</code> property must be set to specify the frontend network vlan of the server.</li>
//         </ul>
//         <http title="Example">{
//     "primaryNetworkComponent": {
//         "networkVlan": {
//             "id": 1
//         }
//     }
// }</http>
//         <br />
//     </li>
//     <li><code>primaryBackendNetworkComponent.networkVlan.id</code>
//         <div>Specifies the network vlan which is to be used for the backend interface of the server.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - int</li>
//             <li><b>Description</b> - The <code>primaryBackendNetworkComponent</code> property is a [[SoftLayer_Network_Component (type)|network component]] structure with the <code>networkVlan</code> property populated with a [[SoftLayer_Network_Vlan (type)|vlan]] structure. The <code>id</code> property must be set to specify the backend network vlan of the server.</li>
//         </ul>
//         <http title="Example">{
//     "primaryBackendNetworkComponent": {
//         "networkVlan": {
//             "id": 2
//         }
//     }
// }</http>
//         <br />
//     </li>
//     <li><code>fixedConfigurationPreset.keyName</code>
//         <div></div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - string</li>
//             <li><b>Description</b> - The <code>fixedConfigurationPreset</code> property is a [[SoftLayer_Product_Package_Preset (type)|fixed configuration preset]] structure. The <code>keyName</code> property must be set to specify preset to use.</li>
//             <li>If a fixed configuration preset is used <code>processorCoreAmount</code>, <code>memoryCapacity</code> and <code>hardDrives</code> properties must not be set.</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//         <http title="Example">{
//     "fixedConfigurationPreset": {
//         "keyName": "SOME_KEY_NAME"
//     }
// }</http>
//         <br />
//     </li>
//     <li><code>userData.value</code>
//         <div>Arbitrary data to be made available to the server.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - string</li>
//             <li><b>Description</b> - The <code>userData</code> property is an array with a single [[SoftLayer_Hardware_Attribute (type)|attribute]] structure with the <code>value</code> property set to an arbitrary value.</li>
//             <li>This value can be retrieved via the [[SoftLayer_Resource_Metadata/getUserMetadata|getUserMetadata]] method from a request originating from the server. This is primarily useful for providing data to software that may be on the server and configured to execute upon first boot.</li>
//         </ul>
//         <http title="Example">{
//     "userData": [
//         {
//             "value": "someValue"
//         }
//     ]
// }</http>
//         <br />
//     </li>
//     <li><code>hardDrives</code>
//         <div>Hard drive settings for the server</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - SoftLayer_Hardware_Component</li>
//             <li><b>Default</b> - The largest available capacity for a zero cost primary disk will be used.</li>
//             <li><b>Description</b> - The <code>hardDrives</code> property is an array of [[SoftLayer_Hardware_Component (type)|hardware component]] structures.</i>
//             <li>Each hard drive must specify the <code>capacity</code> property.</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//         <http title="Example">{
//     "hardDrives": [
//         {
//             "capacity": 500
//         }
//     ]
// }</http>
//         <br />
//     </li>
//     <li id="hardware-create-object-ssh-keys"><code>sshKeys</code>
//         <div>SSH keys to install on the server upon provisioning.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - array of [[SoftLayer_Security_Ssh_Key (type)|SoftLayer_Security_Ssh_Key]]</li>
//             <li><b>Description</b> - The <code>sshKeys</code> property is an array of [[SoftLayer_Security_Ssh_Key (type)|SSH Key]] structures with the <code>id</code> property set to the value of an existing SSH key.</li>
//             <li>To create a new SSH key, call [[SoftLayer_Security_Ssh_Key/createObject|createObject]] on the [[SoftLayer_Security_Ssh_Key]] service.</li>
//             <li>To obtain a list of existing SSH keys, call [[SoftLayer_Account/getSshKeys|getSshKeys]] on the [[SoftLayer_Account]] service.
//         </ul>
//         <http title="Example">{
//     "sshKeys": [
//         {
//             "id": 123
//         }
//     ]
// }</http>
//         <br />
//     </li>
//     <li><code>postInstallScriptUri</code>
//         <div>Specifies the uri location of the script to be downloaded and run after installation is complete.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - string</li>
//         </ul>
//         <br />
//     </li>
// </ul>
//
//
// <h1>REST Example</h1>
// <http title="Request">curl -X POST -d '{
//  "parameters":[
//      {
//          "hostname": "host1",
//          "domain": "example.com",
//          "processorCoreAmount": 2,
//          "memoryCapacity": 2,
//          "hourlyBillingFlag": true,
//          "operatingSystemReferenceCode": "UBUNTU_LATEST"
//      }
//  ]
// }' https://api.softlayer.com/rest/v3/SoftLayer_Hardware.json
// </http>
// <http title="Response">HTTP/1.1 201 Created
// Location: https://api.softlayer.com/rest/v3/SoftLayer_Hardware/f5a3fcff-db1d-4b7c-9fa0-0349e41c29c5/getObject
//
//
// {
//     "accountId": 232298,
//     "bareMetalInstanceFlag": null,
//     "domain": "example.com",
//     "hardwareStatusId": null,
//     "hostname": "host1",
//     "id": null,
//     "serviceProviderId": null,
//     "serviceProviderResourceId": null,
//     "globalIdentifier": "f5a3fcff-db1d-4b7c-9fa0-0349e41c29c5",
//     "hourlyBillingFlag": true,
//     "memoryCapacity": 2,
//     "operatingSystemReferenceCode": "UBUNTU_LATEST",
//     "processorCoreAmount": 2
// }
// </http>
func (r Hardware) CreateObject(templateObject *datatypes.Hardware) (resp datatypes.Hardware, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod("SoftLayer_Hardware", "createObject", params, r.Session, &r.Options, &resp)
	return
}

//
// This method will cancel a server effective immediately. For servers billed hourly, the charges will stop immediately after the method returns.
func (r Hardware) DeleteObject() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware", "deleteObject", nil, r.Session, &r.Options, &resp)
	return
}

// Delete software component passwords.
func (r Hardware) DeleteSoftwareComponentPasswords(softwareComponentPasswords []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		softwareComponentPasswords,
	}
	err = invokeMethod("SoftLayer_Hardware", "deleteSoftwareComponentPasswords", params, r.Session, &r.Options, &resp)
	return
}

// Edit the properties of a software component password such as the username, password, and notes.
func (r Hardware) EditSoftwareComponentPasswords(softwareComponentPasswords []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		softwareComponentPasswords,
	}
	err = invokeMethod("SoftLayer_Hardware", "editSoftwareComponentPasswords", params, r.Session, &r.Options, &resp)
	return
}

// Download and run remote script from uri on the hardware.
func (r Hardware) ExecuteRemoteScript(uri *string) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		uri,
	}
	err = invokeMethod("SoftLayer_Hardware", "executeRemoteScript", params, r.Session, &r.Options, &resp)
	return
}

// The '''findByIpAddress''' method finds hardware using its primary public or private IP address. IP addresses that have a secondary subnet tied to the hardware will not return the hardware - alternate means of locating the hardware must be used (see '''Associated Methods'''). If no hardware is found, no errors are generated and no data is returned.
func (r Hardware) FindByIpAddress(ipAddress *string) (resp datatypes.Hardware, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod("SoftLayer_Hardware", "findByIpAddress", params, r.Session, &r.Options, &resp)
	return
}

//
// Obtain an [[SoftLayer_Container_Product_Order_Hardware_Server (type)|order container]] that can be sent to [[SoftLayer_Product_Order/verifyOrder|verifyOrder]] or [[SoftLayer_Product_Order/placeOrder|placeOrder]].
//
//
// This is primarily useful when there is a necessity to confirm the price which will be charged for an order.
//
//
// See [[SoftLayer_Hardware/createObject|createObject]] for specifics on the requirements of the template object parameter.
func (r Hardware) GenerateOrderTemplate(templateObject *datatypes.Hardware) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod("SoftLayer_Hardware", "generateOrderTemplate", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The account associated with a piece of hardware.
func (r Hardware) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getAccount", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's active physical components.
func (r Hardware) GetActiveComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getActiveComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's active network monitoring incidents.
func (r Hardware) GetActiveNetworkMonitorIncident() (resp []datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getActiveNetworkMonitorIncident", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getAlarmHistory''' method retrieves a detailed history for the monitoring alarm. When calling this method, a start and end date for the history to be retrieved must be entered.
func (r Hardware) GetAlarmHistory(startDate *datatypes.Time, endDate *datatypes.Time, alarmId *string) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
		alarmId,
	}
	err = invokeMethod("SoftLayer_Hardware", "getAlarmHistory", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware) GetAllPowerComponents() (resp []datatypes.Hardware_Power_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getAllPowerComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Storage_Allowed_Host information to connect this server to Network Storage volumes that require access control lists.
func (r Hardware) GetAllowedHost() (resp datatypes.Network_Storage_Allowed_Host, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getAllowedHost", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Storage objects that this SoftLayer_Hardware has access to.
func (r Hardware) GetAllowedNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getAllowedNetworkStorage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Storage objects whose Replica that this SoftLayer_Hardware has access to.
func (r Hardware) GetAllowedNetworkStorageReplicas() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getAllowedNetworkStorageReplicas", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding an antivirus/spyware software component object.
func (r Hardware) GetAntivirusSpywareSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getAntivirusSpywareSoftwareComponent", nil, r.Session, &r.Options, &resp)
	return
}

// This method is retrieve a list of SoftLayer_Network_Storage volumes that are authorized access to this SoftLayer_Hardware.
func (r Hardware) GetAttachedNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod("SoftLayer_Hardware", "getAttachedNetworkStorages", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's specific attributes.
func (r Hardware) GetAttributes() (resp []datatypes.Hardware_Attribute, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getAttributes", nil, r.Session, &r.Options, &resp)
	return
}

// This method retrieves a list of SoftLayer_Network_Storage volumes that can be authorized to this SoftLayer_Hardware.
func (r Hardware) GetAvailableNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod("SoftLayer_Hardware", "getAvailableNetworkStorages", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The average daily public bandwidth usage for the current billing cycle.
func (r Hardware) GetAverageDailyPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getAverageDailyPublicBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getBackendIncomingBandwidth''' method retrieves the amount of incoming private network traffic used between the given start date and end date parameters. When entering start and end dates, only the month, day and year are used to calculate bandwidth totals - the time (HH:MM:SS) is ignored and defaults to midnight. The amount of bandwidth retrieved is measured in gigabytes.
func (r Hardware) GetBackendIncomingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware", "getBackendIncomingBandwidth", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's back-end or private network components.
func (r Hardware) GetBackendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getBackendNetworkComponents", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getBackendOutgoingBandwidth''' method retrieves the amount of outgoing private network traffic used between the given start date and end date parameters. When entering start and end dates, only the month, day and year are used to calculate bandwidth totals - the time (HH:MM:SS) is ignored and defaults to midnight. The amount of bandwidth retrieved is measured in gigabytes.
func (r Hardware) GetBackendOutgoingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware", "getBackendOutgoingBandwidth", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's backend or private router.
func (r Hardware) GetBackendRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getBackendRouters", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's allotted bandwidth (measured in GB).
func (r Hardware) GetBandwidthAllocation() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getBandwidthAllocation", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's allotted detail record. Allotment details link bandwidth allocation with allotments.
func (r Hardware) GetBandwidthAllotmentDetail() (resp datatypes.Network_Bandwidth_Version1_Allotment_Detail, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getBandwidthAllotmentDetail", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's benchmark certifications.
func (r Hardware) GetBenchmarkCertifications() (resp []datatypes.Hardware_Benchmark_Certification, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getBenchmarkCertifications", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the billing item for a server.
func (r Hardware) GetBillingItem() (resp datatypes.Billing_Item_Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getBillingItem", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A flag indicating that a billing item exists.
func (r Hardware) GetBillingItemFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getBillingItemFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Determines whether the hardware is ineligible for cancellation because it is disconnected.
func (r Hardware) GetBlockCancelBecauseDisconnectedFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getBlockCancelBecauseDisconnectedFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Status indicating whether or not a piece of hardware has business continuance insurance.
func (r Hardware) GetBusinessContinuanceInsuranceFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getBusinessContinuanceInsuranceFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's components.
func (r Hardware) GetComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A continuous data protection/server backup software component object.
func (r Hardware) GetContinuousDataProtectionSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getContinuousDataProtectionSoftwareComponent", nil, r.Session, &r.Options, &resp)
	return
}

//
// There are many options that may be provided while ordering a server, this method can be used to determine what these options are.
//
//
// Detailed information on the return value can be found on the data type page for [[SoftLayer_Container_Hardware_Configuration (type)]].
func (r Hardware) GetCreateObjectOptions() (resp datatypes.Container_Hardware_Configuration, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getCreateObjectOptions", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The current billable public outbound bandwidth for this hardware for the current billing cycle.
func (r Hardware) GetCurrentBillableBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getCurrentBillableBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Get the billing detail for this instance for the current billing period. This does not include bandwidth usage.
func (r Hardware) GetCurrentBillingDetail() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getCurrentBillingDetail", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getCurrentBillingTotal''' method retrieves the total bill amount in US Dollars ($) for the current billing period. In addition to the total bill amount, the billing detail also includes all bandwidth used up to the point the method is called on the piece of hardware.
func (r Hardware) GetCurrentBillingTotal() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getCurrentBillingTotal", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getDailyAverage''' method calculates the average daily network traffic used by the selected server. Using the required parameter ''dateTime'' to enter a start and end date, the user retrieves this average, measure in gigabytes (GB) for the specified date range. When entering parameters, only the month, day and year are required - time entries are omitted as this method defaults the time to midnight in order to account for the entire day.
func (r Hardware) GetDailyAverage(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware", "getDailyAverage", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the datacenter in which a piece of hardware resides.
func (r Hardware) GetDatacenter() (resp datatypes.Location, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getDatacenter", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The name of the datacenter in which a piece of hardware resides.
func (r Hardware) GetDatacenterName() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getDatacenterName", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All hardware that has uplink network connections to a piece of hardware.
func (r Hardware) GetDownlinkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getDownlinkHardware", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All hardware that has uplink network connections to a piece of hardware.
func (r Hardware) GetDownlinkNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getDownlinkNetworkHardware", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all servers attached to a piece of network hardware.
func (r Hardware) GetDownlinkServers() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getDownlinkServers", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all virtual guests attached to a piece of network hardware.
func (r Hardware) GetDownlinkVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getDownlinkVirtualGuests", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All hardware downstream from a network device.
func (r Hardware) GetDownstreamHardwareBindings() (resp []datatypes.Network_Component_Uplink_Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getDownstreamHardwareBindings", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All network hardware downstream from the selected piece of hardware.
func (r Hardware) GetDownstreamNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getDownstreamNetworkHardware", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All network hardware with monitoring warnings or errors that are downstream from the selected piece of hardware.
func (r Hardware) GetDownstreamNetworkHardwareWithIncidents() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getDownstreamNetworkHardwareWithIncidents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all servers attached downstream to a piece of network hardware.
func (r Hardware) GetDownstreamServers() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getDownstreamServers", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all virtual guests attached to a piece of network hardware.
func (r Hardware) GetDownstreamVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getDownstreamVirtualGuests", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The drive controllers contained within a piece of hardware.
func (r Hardware) GetDriveControllers() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getDriveControllers", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's associated EVault network storage service account.
func (r Hardware) GetEvaultNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getEvaultNetworkStorage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's firewall services.
func (r Hardware) GetFirewallServiceComponent() (resp datatypes.Network_Component_Firewall, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getFirewallServiceComponent", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Defines the fixed components in a fixed configuration bare metal server.
func (r Hardware) GetFixedConfigurationPreset() (resp datatypes.Product_Package_Preset, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getFixedConfigurationPreset", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getFrontendIncomingBandwidth''' method retrieves the amount of incoming public network traffic used by a server between the given start and end date parameters. When entering the ''dateTime'' parameter, only the month, day and year of the start and end dates are required - the time (hour, minute and second) are set to midnight by default and cannot be changed. The amount of bandwidth retrieved is measured in gigabytes (GB).
func (r Hardware) GetFrontendIncomingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware", "getFrontendIncomingBandwidth", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's front-end or public network components.
func (r Hardware) GetFrontendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getFrontendNetworkComponents", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getFrontendOutgoingBandwidth''' method retrieves the amount of outgoing public network traffic used by a server between the given start and end date parameters. The ''dateTime'' parameter requires only the day, month and year to be entered - the time (hour, minute and second) are set to midnight be default in order to gather the data for the entire start and end date indicated in the parameter. The amount of bandwidth retrieved is measured in gigabytes (GB).
func (r Hardware) GetFrontendOutgoingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware", "getFrontendOutgoingBandwidth", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's frontend or public router.
func (r Hardware) GetFrontendRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getFrontendRouters", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's universally unique identifier.
func (r Hardware) GetGlobalIdentifier() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getGlobalIdentifier", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hard drives contained within a piece of hardware.
func (r Hardware) GetHardDrives() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getHardDrives", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The chassis that a piece of hardware is housed in.
func (r Hardware) GetHardwareChassis() (resp datatypes.Hardware_Chassis, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getHardwareChassis", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's function.
func (r Hardware) GetHardwareFunction() (resp datatypes.Hardware_Function, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getHardwareFunction", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's function.
func (r Hardware) GetHardwareFunctionDescription() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getHardwareFunctionDescription", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's status.
func (r Hardware) GetHardwareStatus() (resp datatypes.Hardware_Status, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getHardwareStatus", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Determine in hardware object has TPM enabled.
func (r Hardware) GetHasTrustedPlatformModuleBillingItemFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getHasTrustedPlatformModuleBillingItemFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a host IPS software component object.
func (r Hardware) GetHostIpsSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getHostIpsSoftwareComponent", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getHourlyBandwidth''' method retrieves all bandwidth updates hourly for the specified hardware. Because the potential number of data points can become excessive, the method limits the user to obtain data in 24-hour intervals. The required ''dateTime'' parameter is used as the starting point for the query and will be calculated for the 24-hour period starting with the specified date and time. For example, entering a parameter of
//
// '02/01/2008 0:00'
//
// results in a return of all bandwidth data for the entire day of February 1, 2008, as 0:00 specifies a midnight start date. Please note that the time entered should be completed using a 24-hour clock (military time, astronomical time).
//
// For data spanning more than a single 24-hour period, refer to the getBandwidthData function on the metricTrackingObject for the piece of hardware.
func (r Hardware) GetHourlyBandwidth(mode *string, day *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		mode,
		day,
	}
	err = invokeMethod("SoftLayer_Hardware", "getHourlyBandwidth", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A server's hourly billing status.
func (r Hardware) GetHourlyBillingFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getHourlyBillingFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The sum of all the inbound network traffic data for the last 30 days.
func (r Hardware) GetInboundBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getInboundBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total public inbound bandwidth for this hardware for the current billing cycle.
func (r Hardware) GetInboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getInboundPublicBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the last transaction a server performed.
func (r Hardware) GetLastTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getLastTransaction", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's latest network monitoring incident.
func (r Hardware) GetLatestNetworkMonitorIncident() (resp datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getLatestNetworkMonitorIncident", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Where a piece of hardware is located within SoftLayer's location hierarchy.
func (r Hardware) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getLocation", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware) GetLocationPathString() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getLocationPathString", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a lockbox account associated with a server.
func (r Hardware) GetLockboxNetworkStorage() (resp datatypes.Network_Storage, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getLockboxNetworkStorage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A flag indicating that the hardware is a managed resource.
func (r Hardware) GetManagedResourceFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getManagedResourceFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's memory.
func (r Hardware) GetMemory() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getMemory", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The amount of memory a piece of hardware has, measured in gigabytes.
func (r Hardware) GetMemoryCapacity() (resp uint, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getMemoryCapacity", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's metric tracking object.
func (r Hardware) GetMetricTrackingObject() (resp datatypes.Metric_Tracking_Object_HardwareServer, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getMetricTrackingObject", nil, r.Session, &r.Options, &resp)
	return
}

// Returns open monitoring alarms for a given time period
func (r Hardware) GetMonitoringActiveAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware", "getMonitoringActiveAlarms", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the monitoring agents associated with a piece of hardware.
func (r Hardware) GetMonitoringAgents() (resp []datatypes.Monitoring_Agent, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getMonitoringAgents", nil, r.Session, &r.Options, &resp)
	return
}

// Returns closed monitoring alarms for a given time period
func (r Hardware) GetMonitoringClosedAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware", "getMonitoringClosedAlarms", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the hardware's monitoring robot.
func (r Hardware) GetMonitoringRobot() (resp datatypes.Monitoring_Robot, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getMonitoringRobot", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's network monitoring services.
func (r Hardware) GetMonitoringServiceComponent() (resp datatypes.Network_Monitor_Version1_Query_Host_Stratum, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getMonitoringServiceComponent", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The monitoring service flag eligibility status for a piece of hardware.
func (r Hardware) GetMonitoringServiceEligibilityFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getMonitoringServiceEligibilityFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The service flag status for a piece of hardware.
func (r Hardware) GetMonitoringServiceFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getMonitoringServiceFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's motherboard.
func (r Hardware) GetMotherboard() (resp datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getMotherboard", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's network cards.
func (r Hardware) GetNetworkCards() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getNetworkCards", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Returns a hardware's network components.
func (r Hardware) GetNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getNetworkComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The gateway member if this device is part of a network gateway.
func (r Hardware) GetNetworkGatewayMember() (resp datatypes.Network_Gateway_Member, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getNetworkGatewayMember", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Whether or not this device is part of a network gateway.
func (r Hardware) GetNetworkGatewayMemberFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getNetworkGatewayMemberFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's network management IP address.
func (r Hardware) GetNetworkManagementIpAddress() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getNetworkManagementIpAddress", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All servers with failed monitoring that are attached downstream to a piece of hardware.
func (r Hardware) GetNetworkMonitorAttachedDownHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getNetworkMonitorAttachedDownHardware", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Virtual guests that are attached downstream to a hardware that have failed monitoring
func (r Hardware) GetNetworkMonitorAttachedDownVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getNetworkMonitorAttachedDownVirtualGuests", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The status of all of a piece of hardware's network monitoring incidents.
func (r Hardware) GetNetworkMonitorIncidents() (resp []datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getNetworkMonitorIncidents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's network monitors.
func (r Hardware) GetNetworkMonitors() (resp []datatypes.Network_Monitor_Version1_Query_Host, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getNetworkMonitors", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The value of a hardware's network status attribute.
func (r Hardware) GetNetworkStatus() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getNetworkStatus", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware's related network status attribute.
func (r Hardware) GetNetworkStatusAttribute() (resp datatypes.Hardware_Attribute, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getNetworkStatusAttribute", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's associated network storage service account.
func (r Hardware) GetNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getNetworkStorage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The network virtual LANs (VLANs) associated with a piece of hardware's network components.
func (r Hardware) GetNetworkVlans() (resp []datatypes.Network_Vlan, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getNetworkVlans", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's allotted bandwidth for the next billing cycle (measured in GB).
func (r Hardware) GetNextBillingCycleBandwidthAllocation() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getNextBillingCycleBandwidthAllocation", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware) GetNotesHistory() (resp []datatypes.Hardware_Note, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getNotesHistory", nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Hardware object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Hardware service. You can only retrieve the account that your portal user is assigned to.
func (r Hardware) GetObject() (resp datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getObject", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's operating system.
func (r Hardware) GetOperatingSystem() (resp datatypes.Software_Component_OperatingSystem, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getOperatingSystem", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's operating system software description.
func (r Hardware) GetOperatingSystemReferenceCode() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getOperatingSystemReferenceCode", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The sum of all the outbound network traffic data for the last 30 days.
func (r Hardware) GetOutboundBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getOutboundBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total public outbound bandwidth for this hardware for the current billing cycle.
func (r Hardware) GetOutboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getOutboundPublicBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the Point of Presence (PoP) location in which a piece of hardware resides.
func (r Hardware) GetPointOfPresenceLocation() (resp datatypes.Location, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getPointOfPresenceLocation", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The power components for a hardware object.
func (r Hardware) GetPowerComponents() (resp []datatypes.Hardware_Power_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getPowerComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's power supply.
func (r Hardware) GetPowerSupply() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getPowerSupply", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware's primary private IP address.
func (r Hardware) GetPrimaryBackendIpAddress() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getPrimaryBackendIpAddress", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the hardware's primary back-end network component.
func (r Hardware) GetPrimaryBackendNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getPrimaryBackendNetworkComponent", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware's primary public IP address.
func (r Hardware) GetPrimaryIpAddress() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getPrimaryIpAddress", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the hardware's primary public network component.
func (r Hardware) GetPrimaryNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getPrimaryNetworkComponent", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a graph of a server's private network bandwidth usage over the specified timeframe. If no timeframe is specified then getPublicBandwidthGraphImage retrieves the last 24 hours of public bandwidth usage. getPrivateBandwidthGraphImage returns a PNG image measuring 827 pixels by 293 pixels.
func (r Hardware) GetPrivateBandwidthData(startTime *int, endTime *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod("SoftLayer_Hardware", "getPrivateBandwidthData", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Whether the hardware only has access to the private network.
func (r Hardware) GetPrivateNetworkOnlyFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getPrivateNetworkOnlyFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total number of processor cores, summed from all processors that are attached to a piece of hardware
func (r Hardware) GetProcessorCoreAmount() (resp uint, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getProcessorCoreAmount", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total number of physical processor cores, summed from all processors that are attached to a piece of hardware
func (r Hardware) GetProcessorPhysicalCoreAmount() (resp uint, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getProcessorPhysicalCoreAmount", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's processors.
func (r Hardware) GetProcessors() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getProcessors", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a graph of a server's public network bandwidth usage over the specified timeframe. If no timeframe is specified then getPublicBandwidthGraphImage retrieves the last 24 hours of public bandwidth usage. getPublicBandwidthGraphImage returns a PNG image measuring 827 pixels by 293 pixels.
func (r Hardware) GetPublicBandwidthData(startTime *int, endTime *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod("SoftLayer_Hardware", "getPublicBandwidthData", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware) GetRack() (resp datatypes.Location, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getRack", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The RAID controllers contained within a piece of hardware.
func (r Hardware) GetRaidControllers() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getRaidControllers", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Recent events that impact this hardware.
func (r Hardware) GetRecentEvents() (resp []datatypes.Notification_Occurrence_Event, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getRecentEvents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve User credentials to issue commands and/or interact with the server's remote management card.
func (r Hardware) GetRemoteManagementAccounts() (resp []datatypes.Hardware_Component_RemoteManagement_User, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getRemoteManagementAccounts", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's associated remote management component. This is normally IPMI.
func (r Hardware) GetRemoteManagementComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getRemoteManagementComponent", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware) GetResourceGroupMemberReferences() (resp []datatypes.Resource_Group_Member, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getResourceGroupMemberReferences", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware) GetResourceGroupRoles() (resp []datatypes.Resource_Group_Role, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getResourceGroupRoles", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The resource groups in which this hardware is a member.
func (r Hardware) GetResourceGroups() (resp []datatypes.Resource_Group, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getResourceGroups", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's routers.
func (r Hardware) GetRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getRouters", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Collection of scale assets this hardware corresponds to.
func (r Hardware) GetScaleAssets() (resp []datatypes.Scale_Asset, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getScaleAssets", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's vulnerability scan requests.
func (r Hardware) GetSecurityScanRequests() (resp []datatypes.Network_Security_Scanner_Request, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getSecurityScanRequests", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getSensorData''' method retrieves a server's hardware state via its internal sensors. Remote sensor data is transmitted to the SoftLayer API by way of the server's remote management card. Sensor data measures various information, including system temperatures, voltages and other local server settings. Sensor data is cached for 30 second; calls made to this method for the same server within 30 seconds of each other will result in the same data being returned. To ensure that the data retrieved retrieves snapshot of varied data, make calls greater than 30 seconds apart.
func (r Hardware) GetSensorData() (resp []datatypes.Container_RemoteManagement_SensorReading, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getSensorData", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getSensorDataWithGraphs''' method retrieves the raw data returned from the server's remote management card. Along with raw data, graphs for the CPU and system temperatures and fan speeds are also returned. For more details on what information is returned, refer to the ''getSensorData'' method.
func (r Hardware) GetSensorDataWithGraphs() (resp datatypes.Container_RemoteManagement_SensorReadingsWithGraphs, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getSensorDataWithGraphs", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getServerFanSpeedGraphs''' method retrieves the server's fan speeds and displays the speeds using tachometer graphs. data used to construct these graphs is retrieved from the server's remote management card. Each graph returned will have an associated title.
func (r Hardware) GetServerFanSpeedGraphs() (resp []datatypes.Container_RemoteManagement_Graphs_SensorSpeed, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getServerFanSpeedGraphs", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getPowerState''' method retrieves the power state for the selected server. The server's power status is retrieved from its remote management card. This method returns "on", for a server that has been powered on, or "off" for servers powered off.
func (r Hardware) GetServerPowerState() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getServerPowerState", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the server room in which the hardware is located.
func (r Hardware) GetServerRoom() (resp datatypes.Location, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getServerRoom", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getServerTemperatureGraphs''' retrieves the server's temperatures and displays the various temperatures using thermometer graphs. Temperatures retrieved are CPU temperature(s) and system temperatures. Data used to construct the graphs is retrieved from the server's remote management card. All graphs returned will have an associated title.
func (r Hardware) GetServerTemperatureGraphs() (resp []datatypes.Container_RemoteManagement_Graphs_SensorTemperature, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getServerTemperatureGraphs", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the piece of hardware's service provider.
func (r Hardware) GetServiceProvider() (resp datatypes.Service_Provider, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getServiceProvider", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's installed software.
func (r Hardware) GetSoftwareComponents() (resp []datatypes.Software_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getSoftwareComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the billing item for a spare pool server.
func (r Hardware) GetSparePoolBillingItem() (resp datatypes.Billing_Item_Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getSparePoolBillingItem", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve SSH keys to be installed on the server during provisioning or an OS reload.
func (r Hardware) GetSshKeys() (resp []datatypes.Security_Ssh_Key, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getSshKeys", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware) GetStorageNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getStorageNetworkComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware) GetTagReferences() (resp []datatypes.Tag_Reference, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getTagReferences", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware) GetTopLevelLocation() (resp datatypes.Location, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getTopLevelLocation", nil, r.Session, &r.Options, &resp)
	return
}

//
// This method will query transaction history for a piece of hardware.
func (r Hardware) GetTransactionHistory() (resp []datatypes.Provisioning_Version1_Transaction_History, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getTransactionHistory", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a list of upgradeable items available to this piece of hardware. Currently, getUpgradeItemPrices retrieves upgrades available for a server's memory, hard drives, network port speed, bandwidth allocation and GPUs.
func (r Hardware) GetUpgradeItemPrices() (resp []datatypes.Product_Item_Price, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getUpgradeItemPrices", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An account's associated upgrade request object, if any.
func (r Hardware) GetUpgradeRequest() (resp datatypes.Product_Upgrade_Request, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getUpgradeRequest", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The network device connected to a piece of hardware.
func (r Hardware) GetUplinkHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getUplinkHardware", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the network component that is one level higher than a piece of hardware on the network infrastructure.
func (r Hardware) GetUplinkNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getUplinkNetworkComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A string containing custom user data for a hardware order.
func (r Hardware) GetUserData() (resp []datatypes.Hardware_Attribute, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getUserData", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the virtual chassis for a piece of hardware.
func (r Hardware) GetVirtualChassis() (resp datatypes.Hardware_Group, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getVirtualChassis", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the virtual chassis siblings for a piece of hardware.
func (r Hardware) GetVirtualChassisSiblings() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getVirtualChassisSiblings", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's virtual host record.
func (r Hardware) GetVirtualHost() (resp datatypes.Virtual_Host, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getVirtualHost", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's virtual software licenses.
func (r Hardware) GetVirtualLicenses() (resp []datatypes.Software_VirtualLicense, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getVirtualLicenses", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the bandwidth allotment to which a piece of hardware belongs.
func (r Hardware) GetVirtualRack() (resp datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getVirtualRack", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The name of the bandwidth allotment belonging to a piece of hardware.
func (r Hardware) GetVirtualRackId() (resp int, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getVirtualRackId", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The name of the bandwidth allotment belonging to a piece of hardware.
func (r Hardware) GetVirtualRackName() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getVirtualRackName", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's virtualization platform software.
func (r Hardware) GetVirtualizationPlatform() (resp datatypes.Software_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware", "getVirtualizationPlatform", nil, r.Session, &r.Options, &resp)
	return
}

// The '''importVirtualHost''' method attempts to import the host record for the virtualization platform running on a server.
func (r Hardware) ImportVirtualHost() (resp datatypes.Virtual_Host, err error) {
	err = invokeMethod("SoftLayer_Hardware", "importVirtualHost", nil, r.Session, &r.Options, &resp)
	return
}

// The '''isPingable''' method issues a ping command to the selected server and returns the result of the ping command. This boolean return value displays ''true'' upon successful ping or ''false'' for a failed ping.
func (r Hardware) IsPingable() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware", "isPingable", nil, r.Session, &r.Options, &resp)
	return
}

// Issues a ping command to the server and returns the ping response.
func (r Hardware) Ping() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware", "ping", nil, r.Session, &r.Options, &resp)
	return
}

// The '''powerCycle''' method completes a power off and power on of the server successively in one command. The power cycle command is equivalent to unplugging the server from the power strip and then plugging the server back in. '''This method should only be used when all other options have been exhausted'''. Additional remote management commands may not be executed if this command was successfully issued within the last 20 minutes to avoid server failure. Remote management commands include:
//
// rebootSoft rebootHard powerOn powerOff powerCycle
//
//
func (r Hardware) PowerCycle() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware", "powerCycle", nil, r.Session, &r.Options, &resp)
	return
}

// This method will power off the server via the server's remote management card.
func (r Hardware) PowerOff() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware", "powerOff", nil, r.Session, &r.Options, &resp)
	return
}

// The '''powerOn''' method powers on a server via its remote management card. This boolean return value returns ''true'' upon successful execution and ''false'' if unsuccessful. Other remote management commands may not be issued in this command was successfully completed within the last 20 minutes to avoid server failure. Remote management commands include:
//
// rebootSoft rebootHard powerOn powerOff powerCycle
//
//
func (r Hardware) PowerOn() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware", "powerOn", nil, r.Session, &r.Options, &resp)
	return
}

// The '''rebootDefault''' method attempts to reboot the server by issuing a soft reboot, or reset, command to the server's remote management card. if the reset attempt is unsuccessful, a power cycle command will be issued via the power strip. The power cycle command is equivalent to unplugging the server from the power strip and then plugging the server back in. If the reset was successful within the last 20 minutes, another remote management command cannot be completed to avoid server failure. Remote management commands include:
//
// rebootSoft rebootHard powerOn powerOff powerCycle
//
//
func (r Hardware) RebootDefault() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware", "rebootDefault", nil, r.Session, &r.Options, &resp)
	return
}

// The '''rebootHard''' method reboots the server by issuing a cycle command to the server's remote management card. A hard reboot is equivalent to pressing the ''Reset'' button on a server - it is issued immediately and will not allow processes to shut down prior to the reboot. Completing a hard reboot may initiate system disk checks upon server reboot, causing the boot up to take longer than normally expected.
//
// Remote management commands are unable to be executed if a reboot has been issued successfully within the last 20 minutes to avoid server failure. Remote management commands include:
//
// rebootSoft rebootHard powerOn powerOff powerCycle
//
//
func (r Hardware) RebootHard() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware", "rebootHard", nil, r.Session, &r.Options, &resp)
	return
}

// The '''rebootSoft''' method reboots the server by issuing a reset command to the server's remote management card via soft reboot. When executing a soft reboot, servers allow all processes to shut down completely before rebooting. Remote management commands are unable to be issued within 20 minutes of issuing a successful soft reboot in order to avoid server failure. Remote management commands include:
//
// rebootSoft rebootHard powerOn powerOff powerCycle
//
//
func (r Hardware) RebootSoft() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware", "rebootSoft", nil, r.Session, &r.Options, &resp)
	return
}

// This method is used to remove access to s SoftLayer_Network_Storage volumes that supports host- or network-level access control.
func (r Hardware) RemoveAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod("SoftLayer_Hardware", "removeAccessToNetworkStorage", params, r.Session, &r.Options, &resp)
	return
}

// This method is used to allow access to multiple SoftLayer_Network_Storage volumes that support host- or network-level access control.
func (r Hardware) RemoveAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod("SoftLayer_Hardware", "removeAccessToNetworkStorageList", params, r.Session, &r.Options, &resp)
	return
}

// no documentation yet
func (r Hardware) SetTags(tags *string) (resp bool, err error) {
	params := []interface{}{
		tags,
	}
	err = invokeMethod("SoftLayer_Hardware", "setTags", params, r.Session, &r.Options, &resp)
	return
}

// The SoftLayer_Hardware_Benchmark_Certification data type contains general information relating to a single SoftLayer hardware benchmark certification document.
type Hardware_Benchmark_Certification struct {
	Session *session.Session
	Options sl.Options
}

func GetHardwareBenchmarkCertificationService(sess *session.Session) Hardware_Benchmark_Certification {
	return Hardware_Benchmark_Certification{Session: sess}
}

func (r Hardware_Benchmark_Certification) Id(id int) Hardware_Benchmark_Certification {
	r.Options.Id = &id
	return r
}

func (r Hardware_Benchmark_Certification) Mask(mask string) Hardware_Benchmark_Certification {
	if !strings.HasPrefix(mask, "mask[") && strings.Contains(mask, "[") {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Hardware_Benchmark_Certification) Filter(filter string) Hardware_Benchmark_Certification {
	r.Options.Filter = filter
	return r
}

func (r Hardware_Benchmark_Certification) Limit(limit int) Hardware_Benchmark_Certification {
	r.Options.Limit = &limit
	return r
}

func (r Hardware_Benchmark_Certification) Offset(offset int) Hardware_Benchmark_Certification {
	r.Options.Offset = &offset
	return r
}

// Retrieve Information regarding a benchmark certification result's associated SoftLayer customer account.
func (r Hardware_Benchmark_Certification) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod("SoftLayer_Hardware_Benchmark_Certification", "getAccount", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the piece of hardware on which a benchmark certification test was performed.
func (r Hardware_Benchmark_Certification) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Benchmark_Certification", "getHardware", nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Hardware_Benchmark_Certification object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Hardware_Benchmark_Certification service.
func (r Hardware_Benchmark_Certification) GetObject() (resp datatypes.Hardware_Benchmark_Certification, err error) {
	err = invokeMethod("SoftLayer_Hardware_Benchmark_Certification", "getObject", nil, r.Session, &r.Options, &resp)
	return
}

// Attempt to retrieve the file associated with a benchmark certification result, if such a file exists.  If there is no file for this benchmark certification result, calling this method throws an exception. The "getResultFile" method attempts to retrieve the file associated with a benchmark certification result, if such a file exists. If no file exists for the benchmark certification, an exception is thrown.
func (r Hardware_Benchmark_Certification) GetResultFile() (resp []byte, err error) {
	err = invokeMethod("SoftLayer_Hardware_Benchmark_Certification", "getResultFile", nil, r.Session, &r.Options, &resp)
	return
}

// The SoftLayer_Hardware_Component_Model data type contains general information relating to a single SoftLayer component model.  A component model represents a vendor specific representation of a hardware component.  Every piece of hardware on a server will have a specific hardware component model.
type Hardware_Component_Model struct {
	Session *session.Session
	Options sl.Options
}

func GetHardwareComponentModelService(sess *session.Session) Hardware_Component_Model {
	return Hardware_Component_Model{Session: sess}
}

func (r Hardware_Component_Model) Id(id int) Hardware_Component_Model {
	r.Options.Id = &id
	return r
}

func (r Hardware_Component_Model) Mask(mask string) Hardware_Component_Model {
	if !strings.HasPrefix(mask, "mask[") && strings.Contains(mask, "[") {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Hardware_Component_Model) Filter(filter string) Hardware_Component_Model {
	r.Options.Filter = filter
	return r
}

func (r Hardware_Component_Model) Limit(limit int) Hardware_Component_Model {
	r.Options.Limit = &limit
	return r
}

func (r Hardware_Component_Model) Offset(offset int) Hardware_Component_Model {
	r.Options.Offset = &offset
	return r
}

// Retrieve
func (r Hardware_Component_Model) GetArchitectureType() (resp datatypes.Hardware_Component_Model_Architecture_Type, err error) {
	err = invokeMethod("SoftLayer_Hardware_Component_Model", "getArchitectureType", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Component_Model) GetAttributes() (resp []datatypes.Hardware_Component_Model_Attribute, err error) {
	err = invokeMethod("SoftLayer_Hardware_Component_Model", "getAttributes", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Component_Model) GetCompatibleArrayTypes() (resp []datatypes.Configuration_Storage_Group_Array_Type, err error) {
	err = invokeMethod("SoftLayer_Hardware_Component_Model", "getCompatibleArrayTypes", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All the component models that are compatible with a hardware component model.
func (r Hardware_Component_Model) GetCompatibleChildComponentModels() (resp []datatypes.Hardware_Component_Model, err error) {
	err = invokeMethod("SoftLayer_Hardware_Component_Model", "getCompatibleChildComponentModels", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All the component models that a hardware component model is compatible with.
func (r Hardware_Component_Model) GetCompatibleParentComponentModels() (resp []datatypes.Hardware_Component_Model, err error) {
	err = invokeMethod("SoftLayer_Hardware_Component_Model", "getCompatibleParentComponentModels", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware component model's physical components in inventory.
func (r Hardware_Component_Model) GetHardwareComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Component_Model", "getHardwareComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The non-vendor specific generic component model for a hardware component model.
func (r Hardware_Component_Model) GetHardwareGenericComponentModel() (resp datatypes.Hardware_Component_Model_Generic, err error) {
	err = invokeMethod("SoftLayer_Hardware_Component_Model", "getHardwareGenericComponentModel", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Component_Model) GetInfinibandCompatibleAttribute() (resp datatypes.Hardware_Component_Model_Attribute, err error) {
	err = invokeMethod("SoftLayer_Hardware_Component_Model", "getInfinibandCompatibleAttribute", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Component_Model) GetIsInfinibandCompatible() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Component_Model", "getIsInfinibandCompatible", nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Hardware_Component_Model object.
func (r Hardware_Component_Model) GetObject() (resp datatypes.Hardware_Component_Model, err error) {
	err = invokeMethod("SoftLayer_Hardware_Component_Model", "getObject", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A motherboard's average reboot time.
func (r Hardware_Component_Model) GetRebootTime() (resp datatypes.Hardware_Component_Motherboard_Reboot_Time, err error) {
	err = invokeMethod("SoftLayer_Hardware_Component_Model", "getRebootTime", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware component model's type.
func (r Hardware_Component_Model) GetType() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Component_Model", "getType", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The types of attributes that are allowed for a given hardware component model.
func (r Hardware_Component_Model) GetValidAttributeTypes() (resp []datatypes.Hardware_Component_Model_Attribute_Type, err error) {
	err = invokeMethod("SoftLayer_Hardware_Component_Model", "getValidAttributeTypes", nil, r.Session, &r.Options, &resp)
	return
}

// The SoftLayer_Hardware_Component_Partition_OperatingSystem data type contains general information relating to a single SoftLayer Operating System Partition Template.
type Hardware_Component_Partition_OperatingSystem struct {
	Session *session.Session
	Options sl.Options
}

func GetHardwareComponentPartitionOperatingSystemService(sess *session.Session) Hardware_Component_Partition_OperatingSystem {
	return Hardware_Component_Partition_OperatingSystem{Session: sess}
}

func (r Hardware_Component_Partition_OperatingSystem) Id(id int) Hardware_Component_Partition_OperatingSystem {
	r.Options.Id = &id
	return r
}

func (r Hardware_Component_Partition_OperatingSystem) Mask(mask string) Hardware_Component_Partition_OperatingSystem {
	if !strings.HasPrefix(mask, "mask[") && strings.Contains(mask, "[") {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Hardware_Component_Partition_OperatingSystem) Filter(filter string) Hardware_Component_Partition_OperatingSystem {
	r.Options.Filter = filter
	return r
}

func (r Hardware_Component_Partition_OperatingSystem) Limit(limit int) Hardware_Component_Partition_OperatingSystem {
	r.Options.Limit = &limit
	return r
}

func (r Hardware_Component_Partition_OperatingSystem) Offset(offset int) Hardware_Component_Partition_OperatingSystem {
	r.Options.Offset = &offset
	return r
}

// no documentation yet
func (r Hardware_Component_Partition_OperatingSystem) GetAllObjects() (resp []datatypes.Hardware_Component_Partition_OperatingSystem, err error) {
	err = invokeMethod("SoftLayer_Hardware_Component_Partition_OperatingSystem", "getAllObjects", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getByDescription''' method retrieves all possible partition templates based on the description (required parameter) entered when calling the method. The description is typically the operating system's name. Current recognized values include 'linux', 'windows', 'freebsd', and 'Debian'.
func (r Hardware_Component_Partition_OperatingSystem) GetByDescription(description *string) (resp datatypes.Hardware_Component_Partition_OperatingSystem, err error) {
	params := []interface{}{
		description,
	}
	err = invokeMethod("SoftLayer_Hardware_Component_Partition_OperatingSystem", "getByDescription", params, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Hardware_Component_Partition_OperatingSystem object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Hardware_Component_Partition_OperatingSystem service.s
func (r Hardware_Component_Partition_OperatingSystem) GetObject() (resp datatypes.Hardware_Component_Partition_OperatingSystem, err error) {
	err = invokeMethod("SoftLayer_Hardware_Component_Partition_OperatingSystem", "getObject", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding an operating system's [[SoftLayer_Hardware_Component_Partition_Template|Partition Templates]].
func (r Hardware_Component_Partition_OperatingSystem) GetPartitionTemplates() (resp []datatypes.Hardware_Component_Partition_Template, err error) {
	err = invokeMethod("SoftLayer_Hardware_Component_Partition_OperatingSystem", "getPartitionTemplates", nil, r.Session, &r.Options, &resp)
	return
}

// The SoftLayer_Hardware_Component_Partition_Template data type contains general information relating to a single SoftLayer partition template.  Partition templates group 1 or more partition configurations that can be used to predefined how a hard drives partitions will be configured.
type Hardware_Component_Partition_Template struct {
	Session *session.Session
	Options sl.Options
}

func GetHardwareComponentPartitionTemplateService(sess *session.Session) Hardware_Component_Partition_Template {
	return Hardware_Component_Partition_Template{Session: sess}
}

func (r Hardware_Component_Partition_Template) Id(id int) Hardware_Component_Partition_Template {
	r.Options.Id = &id
	return r
}

func (r Hardware_Component_Partition_Template) Mask(mask string) Hardware_Component_Partition_Template {
	if !strings.HasPrefix(mask, "mask[") && strings.Contains(mask, "[") {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Hardware_Component_Partition_Template) Filter(filter string) Hardware_Component_Partition_Template {
	r.Options.Filter = filter
	return r
}

func (r Hardware_Component_Partition_Template) Limit(limit int) Hardware_Component_Partition_Template {
	r.Options.Limit = &limit
	return r
}

func (r Hardware_Component_Partition_Template) Offset(offset int) Hardware_Component_Partition_Template {
	r.Options.Offset = &offset
	return r
}

// Retrieve A partition template's associated [[SoftLayer_Account|Account]].
func (r Hardware_Component_Partition_Template) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod("SoftLayer_Hardware_Component_Partition_Template", "getAccount", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An individual partition for a partition template. This is identical to 'partitionTemplatePartition' except this will sort unix partitions.
func (r Hardware_Component_Partition_Template) GetData() (resp []datatypes.Hardware_Component_Partition_Template_Partition, err error) {
	err = invokeMethod("SoftLayer_Hardware_Component_Partition_Template", "getData", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Component_Partition_Template) GetExpireDate() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Component_Partition_Template", "getExpireDate", nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Hardware_Component_Partition_Template object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Hardware_Component_Partition_Template service. You can only retrieve the partition templates that your account created or the templates predefined by SoftLayer.
func (r Hardware_Component_Partition_Template) GetObject() (resp datatypes.Hardware_Component_Partition_Template, err error) {
	err = invokeMethod("SoftLayer_Hardware_Component_Partition_Template", "getObject", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A partition template's associated [[SoftLayer_Hardware_Component_Partition_OperatingSystem|Operating System]].
func (r Hardware_Component_Partition_Template) GetPartitionOperatingSystem() (resp datatypes.Hardware_Component_Partition_OperatingSystem, err error) {
	err = invokeMethod("SoftLayer_Hardware_Component_Partition_Template", "getPartitionOperatingSystem", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An individual partition for a partition template.
func (r Hardware_Component_Partition_Template) GetPartitionTemplatePartition() (resp []datatypes.Hardware_Component_Partition_Template_Partition, err error) {
	err = invokeMethod("SoftLayer_Hardware_Component_Partition_Template", "getPartitionTemplatePartition", nil, r.Session, &r.Options, &resp)
	return
}

// The SoftLayer_Hardware_Router data type contains general information relating to a single SoftLayer router.
type Hardware_Router struct {
	Session *session.Session
	Options sl.Options
}

func GetHardwareRouterService(sess *session.Session) Hardware_Router {
	return Hardware_Router{Session: sess}
}

func (r Hardware_Router) Id(id int) Hardware_Router {
	r.Options.Id = &id
	return r
}

func (r Hardware_Router) Mask(mask string) Hardware_Router {
	if !strings.HasPrefix(mask, "mask[") && strings.Contains(mask, "[") {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Hardware_Router) Filter(filter string) Hardware_Router {
	r.Options.Filter = filter
	return r
}

func (r Hardware_Router) Limit(limit int) Hardware_Router {
	r.Options.Limit = &limit
	return r
}

func (r Hardware_Router) Offset(offset int) Hardware_Router {
	r.Options.Offset = &offset
	return r
}

// This method is used to allow access to a SoftLayer_Network_Storage volume that supports host- or network-level access control.
func (r Hardware_Router) AllowAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod("SoftLayer_Hardware_Router", "allowAccessToNetworkStorage", params, r.Session, &r.Options, &resp)
	return
}

// This method is used to allow access to multiple SoftLayer_Network_Storage volumes that support host- or network-level access control.
func (r Hardware_Router) AllowAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod("SoftLayer_Hardware_Router", "allowAccessToNetworkStorageList", params, r.Session, &r.Options, &resp)
	return
}

// Captures a Flex Image of the hard disk on the physical machine, based on the capture template parameter. Returns the image template group containing the disk image.
func (r Hardware_Router) CaptureImage(captureTemplate *datatypes.Container_Disk_Image_Capture_Template) (resp datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	params := []interface{}{
		captureTemplate,
	}
	err = invokeMethod("SoftLayer_Hardware_Router", "captureImage", params, r.Session, &r.Options, &resp)
	return
}

// Returns monitoring alarm detailed history
func (r Hardware_Router) CloseAlarm(alarmId *string) (resp bool, err error) {
	params := []interface{}{
		alarmId,
	}
	err = invokeMethod("SoftLayer_Hardware_Router", "closeAlarm", params, r.Session, &r.Options, &resp)
	return
}

//
// <style type="text/css">.create_object > li > div { padding-top: .5em; padding-bottom: .5em}</style>
// createObject() enables the creation of servers on an account. This
// method is a simplified alternative to interacting with the ordering system directly.
//
//
// In order to create a server, a template object must be sent in with a few required
// values.
//
//
// When this method returns an order will have been placed for a server of the specified configuration.
//
//
// To determine when the server is available you can poll the server via [[SoftLayer_Hardware/getObject|getObject]],
// checking the <code>provisionDate</code> property.
// When <code>provisionDate</code> is not null, the server will be ready. Be sure to use the <code>globalIdentifier</code>
// as your initialization parameter.
//
//
// <b>Warning:</b> Servers created via this method will incur charges on your account. For testing input parameters see [[SoftLayer_Hardware/generateOrderTemplate|generateOrderTemplate]].
//
//
// <b>Input</b> - [[SoftLayer_Hardware (type)|SoftLayer_Hardware]]
// <ul class="create_object">
//     <li><code>hostname</code>
//         <div>Hostname for the server.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - string</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>domain</code>
//         <div>Domain for the server.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - string</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>processorCoreAmount</code>
//         <div>The number of logical CPU cores to allocate.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - int</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>memoryCapacity</code>
//         <div>The amount of memory to allocate in gigabytes.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - int</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>hourlyBillingFlag</code>
//         <div>Specifies the billing type for the server.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - boolean</li>
//             <li>When true the server will be billed on hourly usage, otherwise it will be billed on a monthly basis.</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>operatingSystemReferenceCode</code>
//         <div>An identifier for the operating system to provision the server with.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - string</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>datacenter.name</code>
//         <div>Specifies which datacenter the server is to be provisioned in.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - string</li>
//             <li>The <code>datacenter</code> property is a [[SoftLayer_Location (type)|location]] structure with the <code>name</code> field set.</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//         <http title="Example">{
//     "datacenter": {
//         "name": "dal05"
//     }
// }</http>
//         <br />
//     </li>
//     <li><code>networkComponents.maxSpeed</code>
//         <div>Specifies the connection speed for the server's network components.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - int</li>
//             <li><b>Default</b> - The highest available zero cost port speed will be used.</li>
//             <li><b>Description</b> - The <code>networkComponents</code> property is an array with a single [[SoftLayer_Network_Component (type)|network component]] structure. The <code>maxSpeed</code> property must be set to specify the network uplink speed, in megabits per second, of the server.</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//             <http title="Example">{
//     "networkComponents": [
//         {
//             "maxSpeed": 1000
//         }
//     ]
// }</http>
//         <br />
//     </li>
//     <li><code>networkComponents.redundancyEnabledFlag</code>
//         <div>Specifies whether or not the server's network components should be in redundancy groups.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - bool</li>
//             <li><b>Default</b> - <code>false</code></li>
//             <li><b>Description</b> - The <code>networkComponents</code> property is an array with a single [[SoftLayer_Network_Component (type)|network component]] structure. When the <code>redundancyEnabledFlag</code> property is true the server's network components will be in redundancy groups.</li>
//         </ul>
//             <http title="Example">{
//     "networkComponents": [
//         {
//             "redundancyEnabledFlag": false
//         }
//     ]
// }</http>
//         <br />
//     </li>
//     <li><code>privateNetworkOnlyFlag</code>
//         <div>Specifies whether or not the server only has access to the private network</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - boolean</li>
//             <li><b>Default</b> - <code>false</code></li>
//             <li>When true this flag specifies that a server is to only have access to the private network.</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>primaryNetworkComponent.networkVlan.id</code>
//         <div>Specifies the network vlan which is to be used for the frontend interface of the server.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - int</li>
//             <li><b>Description</b> - The <code>primaryNetworkComponent</code> property is a [[SoftLayer_Network_Component (type)|network component]] structure with the <code>networkVlan</code> property populated with a [[SoftLayer_Network_Vlan (type)|vlan]] structure. The <code>id</code> property must be set to specify the frontend network vlan of the server.</li>
//         </ul>
//         <http title="Example">{
//     "primaryNetworkComponent": {
//         "networkVlan": {
//             "id": 1
//         }
//     }
// }</http>
//         <br />
//     </li>
//     <li><code>primaryBackendNetworkComponent.networkVlan.id</code>
//         <div>Specifies the network vlan which is to be used for the backend interface of the server.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - int</li>
//             <li><b>Description</b> - The <code>primaryBackendNetworkComponent</code> property is a [[SoftLayer_Network_Component (type)|network component]] structure with the <code>networkVlan</code> property populated with a [[SoftLayer_Network_Vlan (type)|vlan]] structure. The <code>id</code> property must be set to specify the backend network vlan of the server.</li>
//         </ul>
//         <http title="Example">{
//     "primaryBackendNetworkComponent": {
//         "networkVlan": {
//             "id": 2
//         }
//     }
// }</http>
//         <br />
//     </li>
//     <li><code>fixedConfigurationPreset.keyName</code>
//         <div></div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - string</li>
//             <li><b>Description</b> - The <code>fixedConfigurationPreset</code> property is a [[SoftLayer_Product_Package_Preset (type)|fixed configuration preset]] structure. The <code>keyName</code> property must be set to specify preset to use.</li>
//             <li>If a fixed configuration preset is used <code>processorCoreAmount</code>, <code>memoryCapacity</code> and <code>hardDrives</code> properties must not be set.</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//         <http title="Example">{
//     "fixedConfigurationPreset": {
//         "keyName": "SOME_KEY_NAME"
//     }
// }</http>
//         <br />
//     </li>
//     <li><code>userData.value</code>
//         <div>Arbitrary data to be made available to the server.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - string</li>
//             <li><b>Description</b> - The <code>userData</code> property is an array with a single [[SoftLayer_Hardware_Attribute (type)|attribute]] structure with the <code>value</code> property set to an arbitrary value.</li>
//             <li>This value can be retrieved via the [[SoftLayer_Resource_Metadata/getUserMetadata|getUserMetadata]] method from a request originating from the server. This is primarily useful for providing data to software that may be on the server and configured to execute upon first boot.</li>
//         </ul>
//         <http title="Example">{
//     "userData": [
//         {
//             "value": "someValue"
//         }
//     ]
// }</http>
//         <br />
//     </li>
//     <li><code>hardDrives</code>
//         <div>Hard drive settings for the server</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - SoftLayer_Hardware_Component</li>
//             <li><b>Default</b> - The largest available capacity for a zero cost primary disk will be used.</li>
//             <li><b>Description</b> - The <code>hardDrives</code> property is an array of [[SoftLayer_Hardware_Component (type)|hardware component]] structures.</i>
//             <li>Each hard drive must specify the <code>capacity</code> property.</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//         <http title="Example">{
//     "hardDrives": [
//         {
//             "capacity": 500
//         }
//     ]
// }</http>
//         <br />
//     </li>
//     <li id="hardware-create-object-ssh-keys"><code>sshKeys</code>
//         <div>SSH keys to install on the server upon provisioning.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - array of [[SoftLayer_Security_Ssh_Key (type)|SoftLayer_Security_Ssh_Key]]</li>
//             <li><b>Description</b> - The <code>sshKeys</code> property is an array of [[SoftLayer_Security_Ssh_Key (type)|SSH Key]] structures with the <code>id</code> property set to the value of an existing SSH key.</li>
//             <li>To create a new SSH key, call [[SoftLayer_Security_Ssh_Key/createObject|createObject]] on the [[SoftLayer_Security_Ssh_Key]] service.</li>
//             <li>To obtain a list of existing SSH keys, call [[SoftLayer_Account/getSshKeys|getSshKeys]] on the [[SoftLayer_Account]] service.
//         </ul>
//         <http title="Example">{
//     "sshKeys": [
//         {
//             "id": 123
//         }
//     ]
// }</http>
//         <br />
//     </li>
//     <li><code>postInstallScriptUri</code>
//         <div>Specifies the uri location of the script to be downloaded and run after installation is complete.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - string</li>
//         </ul>
//         <br />
//     </li>
// </ul>
//
//
// <h1>REST Example</h1>
// <http title="Request">curl -X POST -d '{
//  "parameters":[
//      {
//          "hostname": "host1",
//          "domain": "example.com",
//          "processorCoreAmount": 2,
//          "memoryCapacity": 2,
//          "hourlyBillingFlag": true,
//          "operatingSystemReferenceCode": "UBUNTU_LATEST"
//      }
//  ]
// }' https://api.softlayer.com/rest/v3/SoftLayer_Hardware.json
// </http>
// <http title="Response">HTTP/1.1 201 Created
// Location: https://api.softlayer.com/rest/v3/SoftLayer_Hardware/f5a3fcff-db1d-4b7c-9fa0-0349e41c29c5/getObject
//
//
// {
//     "accountId": 232298,
//     "bareMetalInstanceFlag": null,
//     "domain": "example.com",
//     "hardwareStatusId": null,
//     "hostname": "host1",
//     "id": null,
//     "serviceProviderId": null,
//     "serviceProviderResourceId": null,
//     "globalIdentifier": "f5a3fcff-db1d-4b7c-9fa0-0349e41c29c5",
//     "hourlyBillingFlag": true,
//     "memoryCapacity": 2,
//     "operatingSystemReferenceCode": "UBUNTU_LATEST",
//     "processorCoreAmount": 2
// }
// </http>
func (r Hardware_Router) CreateObject(templateObject *datatypes.Hardware) (resp datatypes.Hardware, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod("SoftLayer_Hardware_Router", "createObject", params, r.Session, &r.Options, &resp)
	return
}

//
// This method will cancel a server effective immediately. For servers billed hourly, the charges will stop immediately after the method returns.
func (r Hardware_Router) DeleteObject() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "deleteObject", nil, r.Session, &r.Options, &resp)
	return
}

// Delete software component passwords.
func (r Hardware_Router) DeleteSoftwareComponentPasswords(softwareComponentPasswords []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		softwareComponentPasswords,
	}
	err = invokeMethod("SoftLayer_Hardware_Router", "deleteSoftwareComponentPasswords", params, r.Session, &r.Options, &resp)
	return
}

// Edit the properties of a software component password such as the username, password, and notes.
func (r Hardware_Router) EditSoftwareComponentPasswords(softwareComponentPasswords []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		softwareComponentPasswords,
	}
	err = invokeMethod("SoftLayer_Hardware_Router", "editSoftwareComponentPasswords", params, r.Session, &r.Options, &resp)
	return
}

// Download and run remote script from uri on the hardware.
func (r Hardware_Router) ExecuteRemoteScript(uri *string) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		uri,
	}
	err = invokeMethod("SoftLayer_Hardware_Router", "executeRemoteScript", params, r.Session, &r.Options, &resp)
	return
}

// The '''findByIpAddress''' method finds hardware using its primary public or private IP address. IP addresses that have a secondary subnet tied to the hardware will not return the hardware - alternate means of locating the hardware must be used (see '''Associated Methods'''). If no hardware is found, no errors are generated and no data is returned.
func (r Hardware_Router) FindByIpAddress(ipAddress *string) (resp datatypes.Hardware, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod("SoftLayer_Hardware_Router", "findByIpAddress", params, r.Session, &r.Options, &resp)
	return
}

//
// Obtain an [[SoftLayer_Container_Product_Order_Hardware_Server (type)|order container]] that can be sent to [[SoftLayer_Product_Order/verifyOrder|verifyOrder]] or [[SoftLayer_Product_Order/placeOrder|placeOrder]].
//
//
// This is primarily useful when there is a necessity to confirm the price which will be charged for an order.
//
//
// See [[SoftLayer_Hardware/createObject|createObject]] for specifics on the requirements of the template object parameter.
func (r Hardware_Router) GenerateOrderTemplate(templateObject *datatypes.Hardware) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod("SoftLayer_Hardware_Router", "generateOrderTemplate", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The account associated with a piece of hardware.
func (r Hardware_Router) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getAccount", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's active physical components.
func (r Hardware_Router) GetActiveComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getActiveComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's active network monitoring incidents.
func (r Hardware_Router) GetActiveNetworkMonitorIncident() (resp []datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getActiveNetworkMonitorIncident", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getAlarmHistory''' method retrieves a detailed history for the monitoring alarm. When calling this method, a start and end date for the history to be retrieved must be entered.
func (r Hardware_Router) GetAlarmHistory(startDate *datatypes.Time, endDate *datatypes.Time, alarmId *string) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
		alarmId,
	}
	err = invokeMethod("SoftLayer_Hardware_Router", "getAlarmHistory", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Router) GetAllPowerComponents() (resp []datatypes.Hardware_Power_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getAllPowerComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Storage_Allowed_Host information to connect this server to Network Storage volumes that require access control lists.
func (r Hardware_Router) GetAllowedHost() (resp datatypes.Network_Storage_Allowed_Host, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getAllowedHost", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Storage objects that this SoftLayer_Hardware has access to.
func (r Hardware_Router) GetAllowedNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getAllowedNetworkStorage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Storage objects whose Replica that this SoftLayer_Hardware has access to.
func (r Hardware_Router) GetAllowedNetworkStorageReplicas() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getAllowedNetworkStorageReplicas", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding an antivirus/spyware software component object.
func (r Hardware_Router) GetAntivirusSpywareSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getAntivirusSpywareSoftwareComponent", nil, r.Session, &r.Options, &resp)
	return
}

// This method is retrieve a list of SoftLayer_Network_Storage volumes that are authorized access to this SoftLayer_Hardware.
func (r Hardware_Router) GetAttachedNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod("SoftLayer_Hardware_Router", "getAttachedNetworkStorages", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's specific attributes.
func (r Hardware_Router) GetAttributes() (resp []datatypes.Hardware_Attribute, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getAttributes", nil, r.Session, &r.Options, &resp)
	return
}

// This method retrieves a list of SoftLayer_Network_Storage volumes that can be authorized to this SoftLayer_Hardware.
func (r Hardware_Router) GetAvailableNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod("SoftLayer_Hardware_Router", "getAvailableNetworkStorages", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The average daily public bandwidth usage for the current billing cycle.
func (r Hardware_Router) GetAverageDailyPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getAverageDailyPublicBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getBackendIncomingBandwidth''' method retrieves the amount of incoming private network traffic used between the given start date and end date parameters. When entering start and end dates, only the month, day and year are used to calculate bandwidth totals - the time (HH:MM:SS) is ignored and defaults to midnight. The amount of bandwidth retrieved is measured in gigabytes.
func (r Hardware_Router) GetBackendIncomingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_Router", "getBackendIncomingBandwidth", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's back-end or private network components.
func (r Hardware_Router) GetBackendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getBackendNetworkComponents", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getBackendOutgoingBandwidth''' method retrieves the amount of outgoing private network traffic used between the given start date and end date parameters. When entering start and end dates, only the month, day and year are used to calculate bandwidth totals - the time (HH:MM:SS) is ignored and defaults to midnight. The amount of bandwidth retrieved is measured in gigabytes.
func (r Hardware_Router) GetBackendOutgoingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_Router", "getBackendOutgoingBandwidth", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's backend or private router.
func (r Hardware_Router) GetBackendRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getBackendRouters", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's allotted bandwidth (measured in GB).
func (r Hardware_Router) GetBandwidthAllocation() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getBandwidthAllocation", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's allotted detail record. Allotment details link bandwidth allocation with allotments.
func (r Hardware_Router) GetBandwidthAllotmentDetail() (resp datatypes.Network_Bandwidth_Version1_Allotment_Detail, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getBandwidthAllotmentDetail", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's benchmark certifications.
func (r Hardware_Router) GetBenchmarkCertifications() (resp []datatypes.Hardware_Benchmark_Certification, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getBenchmarkCertifications", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the billing item for a server.
func (r Hardware_Router) GetBillingItem() (resp datatypes.Billing_Item_Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getBillingItem", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A flag indicating that a billing item exists.
func (r Hardware_Router) GetBillingItemFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getBillingItemFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Determines whether the hardware is ineligible for cancellation because it is disconnected.
func (r Hardware_Router) GetBlockCancelBecauseDisconnectedFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getBlockCancelBecauseDisconnectedFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Associated subnets for a router object.
func (r Hardware_Router) GetBoundSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getBoundSubnets", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Status indicating whether or not a piece of hardware has business continuance insurance.
func (r Hardware_Router) GetBusinessContinuanceInsuranceFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getBusinessContinuanceInsuranceFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's components.
func (r Hardware_Router) GetComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A continuous data protection/server backup software component object.
func (r Hardware_Router) GetContinuousDataProtectionSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getContinuousDataProtectionSoftwareComponent", nil, r.Session, &r.Options, &resp)
	return
}

//
// There are many options that may be provided while ordering a server, this method can be used to determine what these options are.
//
//
// Detailed information on the return value can be found on the data type page for [[SoftLayer_Container_Hardware_Configuration (type)]].
func (r Hardware_Router) GetCreateObjectOptions() (resp datatypes.Container_Hardware_Configuration, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getCreateObjectOptions", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The current billable public outbound bandwidth for this hardware for the current billing cycle.
func (r Hardware_Router) GetCurrentBillableBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getCurrentBillableBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Get the billing detail for this instance for the current billing period. This does not include bandwidth usage.
func (r Hardware_Router) GetCurrentBillingDetail() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getCurrentBillingDetail", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getCurrentBillingTotal''' method retrieves the total bill amount in US Dollars ($) for the current billing period. In addition to the total bill amount, the billing detail also includes all bandwidth used up to the point the method is called on the piece of hardware.
func (r Hardware_Router) GetCurrentBillingTotal() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getCurrentBillingTotal", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getDailyAverage''' method calculates the average daily network traffic used by the selected server. Using the required parameter ''dateTime'' to enter a start and end date, the user retrieves this average, measure in gigabytes (GB) for the specified date range. When entering parameters, only the month, day and year are required - time entries are omitted as this method defaults the time to midnight in order to account for the entire day.
func (r Hardware_Router) GetDailyAverage(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_Router", "getDailyAverage", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the datacenter in which a piece of hardware resides.
func (r Hardware_Router) GetDatacenter() (resp datatypes.Location, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getDatacenter", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The name of the datacenter in which a piece of hardware resides.
func (r Hardware_Router) GetDatacenterName() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getDatacenterName", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All hardware that has uplink network connections to a piece of hardware.
func (r Hardware_Router) GetDownlinkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getDownlinkHardware", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All hardware that has uplink network connections to a piece of hardware.
func (r Hardware_Router) GetDownlinkNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getDownlinkNetworkHardware", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all servers attached to a piece of network hardware.
func (r Hardware_Router) GetDownlinkServers() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getDownlinkServers", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all virtual guests attached to a piece of network hardware.
func (r Hardware_Router) GetDownlinkVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getDownlinkVirtualGuests", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All hardware downstream from a network device.
func (r Hardware_Router) GetDownstreamHardwareBindings() (resp []datatypes.Network_Component_Uplink_Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getDownstreamHardwareBindings", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All network hardware downstream from the selected piece of hardware.
func (r Hardware_Router) GetDownstreamNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getDownstreamNetworkHardware", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All network hardware with monitoring warnings or errors that are downstream from the selected piece of hardware.
func (r Hardware_Router) GetDownstreamNetworkHardwareWithIncidents() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getDownstreamNetworkHardwareWithIncidents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all servers attached downstream to a piece of network hardware.
func (r Hardware_Router) GetDownstreamServers() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getDownstreamServers", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all virtual guests attached to a piece of network hardware.
func (r Hardware_Router) GetDownstreamVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getDownstreamVirtualGuests", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The drive controllers contained within a piece of hardware.
func (r Hardware_Router) GetDriveControllers() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getDriveControllers", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's associated EVault network storage service account.
func (r Hardware_Router) GetEvaultNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getEvaultNetworkStorage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's firewall services.
func (r Hardware_Router) GetFirewallServiceComponent() (resp datatypes.Network_Component_Firewall, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getFirewallServiceComponent", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Defines the fixed components in a fixed configuration bare metal server.
func (r Hardware_Router) GetFixedConfigurationPreset() (resp datatypes.Product_Package_Preset, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getFixedConfigurationPreset", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getFrontendIncomingBandwidth''' method retrieves the amount of incoming public network traffic used by a server between the given start and end date parameters. When entering the ''dateTime'' parameter, only the month, day and year of the start and end dates are required - the time (hour, minute and second) are set to midnight by default and cannot be changed. The amount of bandwidth retrieved is measured in gigabytes (GB).
func (r Hardware_Router) GetFrontendIncomingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_Router", "getFrontendIncomingBandwidth", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's front-end or public network components.
func (r Hardware_Router) GetFrontendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getFrontendNetworkComponents", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getFrontendOutgoingBandwidth''' method retrieves the amount of outgoing public network traffic used by a server between the given start and end date parameters. The ''dateTime'' parameter requires only the day, month and year to be entered - the time (hour, minute and second) are set to midnight be default in order to gather the data for the entire start and end date indicated in the parameter. The amount of bandwidth retrieved is measured in gigabytes (GB).
func (r Hardware_Router) GetFrontendOutgoingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_Router", "getFrontendOutgoingBandwidth", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's frontend or public router.
func (r Hardware_Router) GetFrontendRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getFrontendRouters", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's universally unique identifier.
func (r Hardware_Router) GetGlobalIdentifier() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getGlobalIdentifier", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hard drives contained within a piece of hardware.
func (r Hardware_Router) GetHardDrives() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getHardDrives", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The chassis that a piece of hardware is housed in.
func (r Hardware_Router) GetHardwareChassis() (resp datatypes.Hardware_Chassis, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getHardwareChassis", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's function.
func (r Hardware_Router) GetHardwareFunction() (resp datatypes.Hardware_Function, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getHardwareFunction", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's function.
func (r Hardware_Router) GetHardwareFunctionDescription() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getHardwareFunctionDescription", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's status.
func (r Hardware_Router) GetHardwareStatus() (resp datatypes.Hardware_Status, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getHardwareStatus", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Determine in hardware object has TPM enabled.
func (r Hardware_Router) GetHasTrustedPlatformModuleBillingItemFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getHasTrustedPlatformModuleBillingItemFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a host IPS software component object.
func (r Hardware_Router) GetHostIpsSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getHostIpsSoftwareComponent", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getHourlyBandwidth''' method retrieves all bandwidth updates hourly for the specified hardware. Because the potential number of data points can become excessive, the method limits the user to obtain data in 24-hour intervals. The required ''dateTime'' parameter is used as the starting point for the query and will be calculated for the 24-hour period starting with the specified date and time. For example, entering a parameter of
//
// '02/01/2008 0:00'
//
// results in a return of all bandwidth data for the entire day of February 1, 2008, as 0:00 specifies a midnight start date. Please note that the time entered should be completed using a 24-hour clock (military time, astronomical time).
//
// For data spanning more than a single 24-hour period, refer to the getBandwidthData function on the metricTrackingObject for the piece of hardware.
func (r Hardware_Router) GetHourlyBandwidth(mode *string, day *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		mode,
		day,
	}
	err = invokeMethod("SoftLayer_Hardware_Router", "getHourlyBandwidth", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A server's hourly billing status.
func (r Hardware_Router) GetHourlyBillingFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getHourlyBillingFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The sum of all the inbound network traffic data for the last 30 days.
func (r Hardware_Router) GetInboundBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getInboundBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total public inbound bandwidth for this hardware for the current billing cycle.
func (r Hardware_Router) GetInboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getInboundPublicBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the last transaction a server performed.
func (r Hardware_Router) GetLastTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getLastTransaction", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's latest network monitoring incident.
func (r Hardware_Router) GetLatestNetworkMonitorIncident() (resp datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getLatestNetworkMonitorIncident", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A flag indicating that a VLAN on the router can be assigned to a host that has local disk functionality.
func (r Hardware_Router) GetLocalDiskStorageCapabilityFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getLocalDiskStorageCapabilityFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Where a piece of hardware is located within SoftLayer's location hierarchy.
func (r Hardware_Router) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getLocation", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Router) GetLocationPathString() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getLocationPathString", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a lockbox account associated with a server.
func (r Hardware_Router) GetLockboxNetworkStorage() (resp datatypes.Network_Storage, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getLockboxNetworkStorage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A flag indicating that the hardware is a managed resource.
func (r Hardware_Router) GetManagedResourceFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getManagedResourceFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's memory.
func (r Hardware_Router) GetMemory() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getMemory", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The amount of memory a piece of hardware has, measured in gigabytes.
func (r Hardware_Router) GetMemoryCapacity() (resp uint, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getMemoryCapacity", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's metric tracking object.
func (r Hardware_Router) GetMetricTrackingObject() (resp datatypes.Metric_Tracking_Object_HardwareServer, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getMetricTrackingObject", nil, r.Session, &r.Options, &resp)
	return
}

// Returns open monitoring alarms for a given time period
func (r Hardware_Router) GetMonitoringActiveAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_Router", "getMonitoringActiveAlarms", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the monitoring agents associated with a piece of hardware.
func (r Hardware_Router) GetMonitoringAgents() (resp []datatypes.Monitoring_Agent, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getMonitoringAgents", nil, r.Session, &r.Options, &resp)
	return
}

// Returns closed monitoring alarms for a given time period
func (r Hardware_Router) GetMonitoringClosedAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_Router", "getMonitoringClosedAlarms", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the hardware's monitoring robot.
func (r Hardware_Router) GetMonitoringRobot() (resp datatypes.Monitoring_Robot, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getMonitoringRobot", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's network monitoring services.
func (r Hardware_Router) GetMonitoringServiceComponent() (resp datatypes.Network_Monitor_Version1_Query_Host_Stratum, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getMonitoringServiceComponent", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The monitoring service flag eligibility status for a piece of hardware.
func (r Hardware_Router) GetMonitoringServiceEligibilityFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getMonitoringServiceEligibilityFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The service flag status for a piece of hardware.
func (r Hardware_Router) GetMonitoringServiceFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getMonitoringServiceFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's motherboard.
func (r Hardware_Router) GetMotherboard() (resp datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getMotherboard", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's network cards.
func (r Hardware_Router) GetNetworkCards() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getNetworkCards", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Returns a hardware's network components.
func (r Hardware_Router) GetNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getNetworkComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The gateway member if this device is part of a network gateway.
func (r Hardware_Router) GetNetworkGatewayMember() (resp datatypes.Network_Gateway_Member, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getNetworkGatewayMember", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Whether or not this device is part of a network gateway.
func (r Hardware_Router) GetNetworkGatewayMemberFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getNetworkGatewayMemberFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's network management IP address.
func (r Hardware_Router) GetNetworkManagementIpAddress() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getNetworkManagementIpAddress", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All servers with failed monitoring that are attached downstream to a piece of hardware.
func (r Hardware_Router) GetNetworkMonitorAttachedDownHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getNetworkMonitorAttachedDownHardware", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Virtual guests that are attached downstream to a hardware that have failed monitoring
func (r Hardware_Router) GetNetworkMonitorAttachedDownVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getNetworkMonitorAttachedDownVirtualGuests", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The status of all of a piece of hardware's network monitoring incidents.
func (r Hardware_Router) GetNetworkMonitorIncidents() (resp []datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getNetworkMonitorIncidents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's network monitors.
func (r Hardware_Router) GetNetworkMonitors() (resp []datatypes.Network_Monitor_Version1_Query_Host, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getNetworkMonitors", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The value of a hardware's network status attribute.
func (r Hardware_Router) GetNetworkStatus() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getNetworkStatus", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware's related network status attribute.
func (r Hardware_Router) GetNetworkStatusAttribute() (resp datatypes.Hardware_Attribute, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getNetworkStatusAttribute", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's associated network storage service account.
func (r Hardware_Router) GetNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getNetworkStorage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The network virtual LANs (VLANs) associated with a piece of hardware's network components.
func (r Hardware_Router) GetNetworkVlans() (resp []datatypes.Network_Vlan, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getNetworkVlans", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's allotted bandwidth for the next billing cycle (measured in GB).
func (r Hardware_Router) GetNextBillingCycleBandwidthAllocation() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getNextBillingCycleBandwidthAllocation", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Router) GetNotesHistory() (resp []datatypes.Hardware_Note, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getNotesHistory", nil, r.Session, &r.Options, &resp)
	return
}

// no documentation yet
func (r Hardware_Router) GetObject() (resp datatypes.Hardware_Router, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getObject", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's operating system.
func (r Hardware_Router) GetOperatingSystem() (resp datatypes.Software_Component_OperatingSystem, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getOperatingSystem", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's operating system software description.
func (r Hardware_Router) GetOperatingSystemReferenceCode() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getOperatingSystemReferenceCode", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The sum of all the outbound network traffic data for the last 30 days.
func (r Hardware_Router) GetOutboundBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getOutboundBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total public outbound bandwidth for this hardware for the current billing cycle.
func (r Hardware_Router) GetOutboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getOutboundPublicBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the Point of Presence (PoP) location in which a piece of hardware resides.
func (r Hardware_Router) GetPointOfPresenceLocation() (resp datatypes.Location, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getPointOfPresenceLocation", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The power components for a hardware object.
func (r Hardware_Router) GetPowerComponents() (resp []datatypes.Hardware_Power_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getPowerComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's power supply.
func (r Hardware_Router) GetPowerSupply() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getPowerSupply", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware's primary private IP address.
func (r Hardware_Router) GetPrimaryBackendIpAddress() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getPrimaryBackendIpAddress", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the hardware's primary back-end network component.
func (r Hardware_Router) GetPrimaryBackendNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getPrimaryBackendNetworkComponent", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware's primary public IP address.
func (r Hardware_Router) GetPrimaryIpAddress() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getPrimaryIpAddress", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the hardware's primary public network component.
func (r Hardware_Router) GetPrimaryNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getPrimaryNetworkComponent", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a graph of a server's private network bandwidth usage over the specified timeframe. If no timeframe is specified then getPublicBandwidthGraphImage retrieves the last 24 hours of public bandwidth usage. getPrivateBandwidthGraphImage returns a PNG image measuring 827 pixels by 293 pixels.
func (r Hardware_Router) GetPrivateBandwidthData(startTime *int, endTime *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod("SoftLayer_Hardware_Router", "getPrivateBandwidthData", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Whether the hardware only has access to the private network.
func (r Hardware_Router) GetPrivateNetworkOnlyFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getPrivateNetworkOnlyFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total number of processor cores, summed from all processors that are attached to a piece of hardware
func (r Hardware_Router) GetProcessorCoreAmount() (resp uint, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getProcessorCoreAmount", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total number of physical processor cores, summed from all processors that are attached to a piece of hardware
func (r Hardware_Router) GetProcessorPhysicalCoreAmount() (resp uint, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getProcessorPhysicalCoreAmount", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's processors.
func (r Hardware_Router) GetProcessors() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getProcessors", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a graph of a server's public network bandwidth usage over the specified timeframe. If no timeframe is specified then getPublicBandwidthGraphImage retrieves the last 24 hours of public bandwidth usage. getPublicBandwidthGraphImage returns a PNG image measuring 827 pixels by 293 pixels.
func (r Hardware_Router) GetPublicBandwidthData(startTime *int, endTime *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod("SoftLayer_Hardware_Router", "getPublicBandwidthData", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Router) GetRack() (resp datatypes.Location, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getRack", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The RAID controllers contained within a piece of hardware.
func (r Hardware_Router) GetRaidControllers() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getRaidControllers", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Recent events that impact this hardware.
func (r Hardware_Router) GetRecentEvents() (resp []datatypes.Notification_Occurrence_Event, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getRecentEvents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve User credentials to issue commands and/or interact with the server's remote management card.
func (r Hardware_Router) GetRemoteManagementAccounts() (resp []datatypes.Hardware_Component_RemoteManagement_User, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getRemoteManagementAccounts", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's associated remote management component. This is normally IPMI.
func (r Hardware_Router) GetRemoteManagementComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getRemoteManagementComponent", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Router) GetResourceGroupMemberReferences() (resp []datatypes.Resource_Group_Member, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getResourceGroupMemberReferences", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Router) GetResourceGroupRoles() (resp []datatypes.Resource_Group_Role, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getResourceGroupRoles", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The resource groups in which this hardware is a member.
func (r Hardware_Router) GetResourceGroups() (resp []datatypes.Resource_Group, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getResourceGroups", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's routers.
func (r Hardware_Router) GetRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getRouters", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A flag indicating that a VLAN on the router can be assigned to a host that has SAN disk functionality.
func (r Hardware_Router) GetSanStorageCapabilityFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getSanStorageCapabilityFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Collection of scale assets this hardware corresponds to.
func (r Hardware_Router) GetScaleAssets() (resp []datatypes.Scale_Asset, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getScaleAssets", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's vulnerability scan requests.
func (r Hardware_Router) GetSecurityScanRequests() (resp []datatypes.Network_Security_Scanner_Request, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getSecurityScanRequests", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getSensorData''' method retrieves a server's hardware state via its internal sensors. Remote sensor data is transmitted to the SoftLayer API by way of the server's remote management card. Sensor data measures various information, including system temperatures, voltages and other local server settings. Sensor data is cached for 30 second; calls made to this method for the same server within 30 seconds of each other will result in the same data being returned. To ensure that the data retrieved retrieves snapshot of varied data, make calls greater than 30 seconds apart.
func (r Hardware_Router) GetSensorData() (resp []datatypes.Container_RemoteManagement_SensorReading, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getSensorData", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getSensorDataWithGraphs''' method retrieves the raw data returned from the server's remote management card. Along with raw data, graphs for the CPU and system temperatures and fan speeds are also returned. For more details on what information is returned, refer to the ''getSensorData'' method.
func (r Hardware_Router) GetSensorDataWithGraphs() (resp datatypes.Container_RemoteManagement_SensorReadingsWithGraphs, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getSensorDataWithGraphs", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getServerFanSpeedGraphs''' method retrieves the server's fan speeds and displays the speeds using tachometer graphs. data used to construct these graphs is retrieved from the server's remote management card. Each graph returned will have an associated title.
func (r Hardware_Router) GetServerFanSpeedGraphs() (resp []datatypes.Container_RemoteManagement_Graphs_SensorSpeed, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getServerFanSpeedGraphs", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getPowerState''' method retrieves the power state for the selected server. The server's power status is retrieved from its remote management card. This method returns "on", for a server that has been powered on, or "off" for servers powered off.
func (r Hardware_Router) GetServerPowerState() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getServerPowerState", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the server room in which the hardware is located.
func (r Hardware_Router) GetServerRoom() (resp datatypes.Location, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getServerRoom", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getServerTemperatureGraphs''' retrieves the server's temperatures and displays the various temperatures using thermometer graphs. Temperatures retrieved are CPU temperature(s) and system temperatures. Data used to construct the graphs is retrieved from the server's remote management card. All graphs returned will have an associated title.
func (r Hardware_Router) GetServerTemperatureGraphs() (resp []datatypes.Container_RemoteManagement_Graphs_SensorTemperature, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getServerTemperatureGraphs", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the piece of hardware's service provider.
func (r Hardware_Router) GetServiceProvider() (resp datatypes.Service_Provider, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getServiceProvider", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's installed software.
func (r Hardware_Router) GetSoftwareComponents() (resp []datatypes.Software_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getSoftwareComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the billing item for a spare pool server.
func (r Hardware_Router) GetSparePoolBillingItem() (resp datatypes.Billing_Item_Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getSparePoolBillingItem", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve SSH keys to be installed on the server during provisioning or an OS reload.
func (r Hardware_Router) GetSshKeys() (resp []datatypes.Security_Ssh_Key, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getSshKeys", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Router) GetStorageNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getStorageNetworkComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Router) GetTagReferences() (resp []datatypes.Tag_Reference, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getTagReferences", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Router) GetTopLevelLocation() (resp datatypes.Location, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getTopLevelLocation", nil, r.Session, &r.Options, &resp)
	return
}

//
// This method will query transaction history for a piece of hardware.
func (r Hardware_Router) GetTransactionHistory() (resp []datatypes.Provisioning_Version1_Transaction_History, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getTransactionHistory", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a list of upgradeable items available to this piece of hardware. Currently, getUpgradeItemPrices retrieves upgrades available for a server's memory, hard drives, network port speed, bandwidth allocation and GPUs.
func (r Hardware_Router) GetUpgradeItemPrices() (resp []datatypes.Product_Item_Price, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getUpgradeItemPrices", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An account's associated upgrade request object, if any.
func (r Hardware_Router) GetUpgradeRequest() (resp datatypes.Product_Upgrade_Request, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getUpgradeRequest", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The network device connected to a piece of hardware.
func (r Hardware_Router) GetUplinkHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getUplinkHardware", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the network component that is one level higher than a piece of hardware on the network infrastructure.
func (r Hardware_Router) GetUplinkNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getUplinkNetworkComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A string containing custom user data for a hardware order.
func (r Hardware_Router) GetUserData() (resp []datatypes.Hardware_Attribute, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getUserData", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the virtual chassis for a piece of hardware.
func (r Hardware_Router) GetVirtualChassis() (resp datatypes.Hardware_Group, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getVirtualChassis", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the virtual chassis siblings for a piece of hardware.
func (r Hardware_Router) GetVirtualChassisSiblings() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getVirtualChassisSiblings", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's virtual host record.
func (r Hardware_Router) GetVirtualHost() (resp datatypes.Virtual_Host, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getVirtualHost", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's virtual software licenses.
func (r Hardware_Router) GetVirtualLicenses() (resp []datatypes.Software_VirtualLicense, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getVirtualLicenses", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the bandwidth allotment to which a piece of hardware belongs.
func (r Hardware_Router) GetVirtualRack() (resp datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getVirtualRack", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The name of the bandwidth allotment belonging to a piece of hardware.
func (r Hardware_Router) GetVirtualRackId() (resp int, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getVirtualRackId", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The name of the bandwidth allotment belonging to a piece of hardware.
func (r Hardware_Router) GetVirtualRackName() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getVirtualRackName", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's virtualization platform software.
func (r Hardware_Router) GetVirtualizationPlatform() (resp datatypes.Software_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "getVirtualizationPlatform", nil, r.Session, &r.Options, &resp)
	return
}

// The '''importVirtualHost''' method attempts to import the host record for the virtualization platform running on a server.
func (r Hardware_Router) ImportVirtualHost() (resp datatypes.Virtual_Host, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "importVirtualHost", nil, r.Session, &r.Options, &resp)
	return
}

// The '''isPingable''' method issues a ping command to the selected server and returns the result of the ping command. This boolean return value displays ''true'' upon successful ping or ''false'' for a failed ping.
func (r Hardware_Router) IsPingable() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "isPingable", nil, r.Session, &r.Options, &resp)
	return
}

// Issues a ping command to the server and returns the ping response.
func (r Hardware_Router) Ping() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "ping", nil, r.Session, &r.Options, &resp)
	return
}

// The '''powerCycle''' method completes a power off and power on of the server successively in one command. The power cycle command is equivalent to unplugging the server from the power strip and then plugging the server back in. '''This method should only be used when all other options have been exhausted'''. Additional remote management commands may not be executed if this command was successfully issued within the last 20 minutes to avoid server failure. Remote management commands include:
//
// rebootSoft rebootHard powerOn powerOff powerCycle
//
//
func (r Hardware_Router) PowerCycle() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "powerCycle", nil, r.Session, &r.Options, &resp)
	return
}

// This method will power off the server via the server's remote management card.
func (r Hardware_Router) PowerOff() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "powerOff", nil, r.Session, &r.Options, &resp)
	return
}

// The '''powerOn''' method powers on a server via its remote management card. This boolean return value returns ''true'' upon successful execution and ''false'' if unsuccessful. Other remote management commands may not be issued in this command was successfully completed within the last 20 minutes to avoid server failure. Remote management commands include:
//
// rebootSoft rebootHard powerOn powerOff powerCycle
//
//
func (r Hardware_Router) PowerOn() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "powerOn", nil, r.Session, &r.Options, &resp)
	return
}

// The '''rebootDefault''' method attempts to reboot the server by issuing a soft reboot, or reset, command to the server's remote management card. if the reset attempt is unsuccessful, a power cycle command will be issued via the power strip. The power cycle command is equivalent to unplugging the server from the power strip and then plugging the server back in. If the reset was successful within the last 20 minutes, another remote management command cannot be completed to avoid server failure. Remote management commands include:
//
// rebootSoft rebootHard powerOn powerOff powerCycle
//
//
func (r Hardware_Router) RebootDefault() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "rebootDefault", nil, r.Session, &r.Options, &resp)
	return
}

// The '''rebootHard''' method reboots the server by issuing a cycle command to the server's remote management card. A hard reboot is equivalent to pressing the ''Reset'' button on a server - it is issued immediately and will not allow processes to shut down prior to the reboot. Completing a hard reboot may initiate system disk checks upon server reboot, causing the boot up to take longer than normally expected.
//
// Remote management commands are unable to be executed if a reboot has been issued successfully within the last 20 minutes to avoid server failure. Remote management commands include:
//
// rebootSoft rebootHard powerOn powerOff powerCycle
//
//
func (r Hardware_Router) RebootHard() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "rebootHard", nil, r.Session, &r.Options, &resp)
	return
}

// The '''rebootSoft''' method reboots the server by issuing a reset command to the server's remote management card via soft reboot. When executing a soft reboot, servers allow all processes to shut down completely before rebooting. Remote management commands are unable to be issued within 20 minutes of issuing a successful soft reboot in order to avoid server failure. Remote management commands include:
//
// rebootSoft rebootHard powerOn powerOff powerCycle
//
//
func (r Hardware_Router) RebootSoft() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Router", "rebootSoft", nil, r.Session, &r.Options, &resp)
	return
}

// This method is used to remove access to s SoftLayer_Network_Storage volumes that supports host- or network-level access control.
func (r Hardware_Router) RemoveAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod("SoftLayer_Hardware_Router", "removeAccessToNetworkStorage", params, r.Session, &r.Options, &resp)
	return
}

// This method is used to allow access to multiple SoftLayer_Network_Storage volumes that support host- or network-level access control.
func (r Hardware_Router) RemoveAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod("SoftLayer_Hardware_Router", "removeAccessToNetworkStorageList", params, r.Session, &r.Options, &resp)
	return
}

// no documentation yet
func (r Hardware_Router) SetTags(tags *string) (resp bool, err error) {
	params := []interface{}{
		tags,
	}
	err = invokeMethod("SoftLayer_Hardware_Router", "setTags", params, r.Session, &r.Options, &resp)
	return
}

// no documentation yet
type Hardware_SecurityModule struct {
	Session *session.Session
	Options sl.Options
}

func GetHardwareSecurityModuleService(sess *session.Session) Hardware_SecurityModule {
	return Hardware_SecurityModule{Session: sess}
}

func (r Hardware_SecurityModule) Id(id int) Hardware_SecurityModule {
	r.Options.Id = &id
	return r
}

func (r Hardware_SecurityModule) Mask(mask string) Hardware_SecurityModule {
	if !strings.HasPrefix(mask, "mask[") && strings.Contains(mask, "[") {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Hardware_SecurityModule) Filter(filter string) Hardware_SecurityModule {
	r.Options.Filter = filter
	return r
}

func (r Hardware_SecurityModule) Limit(limit int) Hardware_SecurityModule {
	r.Options.Limit = &limit
	return r
}

func (r Hardware_SecurityModule) Offset(offset int) Hardware_SecurityModule {
	r.Options.Offset = &offset
	return r
}

// Activates the private network port
func (r Hardware_SecurityModule) ActivatePrivatePort() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "activatePrivatePort", nil, r.Session, &r.Options, &resp)
	return
}

// Activates the public network port
func (r Hardware_SecurityModule) ActivatePublicPort() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "activatePublicPort", nil, r.Session, &r.Options, &resp)
	return
}

// This method is used to allow access to a SoftLayer_Network_Storage volume that supports host- or network-level access control.
func (r Hardware_SecurityModule) AllowAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "allowAccessToNetworkStorage", params, r.Session, &r.Options, &resp)
	return
}

// This method is used to allow access to multiple SoftLayer_Network_Storage volumes that support host- or network-level access control.
func (r Hardware_SecurityModule) AllowAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "allowAccessToNetworkStorageList", params, r.Session, &r.Options, &resp)
	return
}

// The Rescue Kernel is designed to provide you with the ability to bring a server online in order to troubleshoot system problems that would normally only be resolved by an OS Reload. The correct Rescue Kernel will be selected based upon the currently installed operating system. When the rescue kernel process is initiated, the server will shutdown and reboot on to the public network with the same IP's assigned to the server to allow for remote connections. It will bring your server offline for approximately 10 minutes while the rescue is in progress. The root/administrator password will be the same as what is listed in the portal for the server.
func (r Hardware_SecurityModule) BootToRescueLayer(noOsBootEnvironment *string) (resp bool, err error) {
	params := []interface{}{
		noOsBootEnvironment,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "bootToRescueLayer", params, r.Session, &r.Options, &resp)
	return
}

// Captures a Flex Image of the hard disk on the physical machine, based on the capture template parameter. Returns the image template group containing the disk image.
func (r Hardware_SecurityModule) CaptureImage(captureTemplate *datatypes.Container_Disk_Image_Capture_Template) (resp datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	params := []interface{}{
		captureTemplate,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "captureImage", params, r.Session, &r.Options, &resp)
	return
}

// Returns monitoring alarm detailed history
func (r Hardware_SecurityModule) CloseAlarm(alarmId *string) (resp bool, err error) {
	params := []interface{}{
		alarmId,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "closeAlarm", params, r.Session, &r.Options, &resp)
	return
}

// You can launch firmware updates by selecting from your server list. It will bring your server offline for approximately 20 minutes while the updates are in progress.
//
// In the event of a hardware failure during this test our datacenter engineers will be notified of the problem automatically. They will then replace any failed components to bring your server back online, and will be contacting you to ensure that impact on your server is minimal.
func (r Hardware_SecurityModule) CreateFirmwareUpdateTransaction(ipmi *int, raidController *int, bios *int, harddrive *int) (resp bool, err error) {
	params := []interface{}{
		ipmi,
		raidController,
		bios,
		harddrive,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "createFirmwareUpdateTransaction", params, r.Session, &r.Options, &resp)
	return
}

//
// <style type="text/css">.create_object > li > div { padding-top: .5em; padding-bottom: .5em}</style>
// createObject() enables the creation of servers on an account. This
// method is a simplified alternative to interacting with the ordering system directly.
//
//
// In order to create a server, a template object must be sent in with a few required
// values.
//
//
// When this method returns an order will have been placed for a server of the specified configuration.
//
//
// To determine when the server is available you can poll the server via [[SoftLayer_Hardware/getObject|getObject]],
// checking the <code>provisionDate</code> property.
// When <code>provisionDate</code> is not null, the server will be ready. Be sure to use the <code>globalIdentifier</code>
// as your initialization parameter.
//
//
// <b>Warning:</b> Servers created via this method will incur charges on your account. For testing input parameters see [[SoftLayer_Hardware/generateOrderTemplate|generateOrderTemplate]].
//
//
// <b>Input</b> - [[SoftLayer_Hardware (type)|SoftLayer_Hardware]]
// <ul class="create_object">
//     <li><code>hostname</code>
//         <div>Hostname for the server.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - string</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>domain</code>
//         <div>Domain for the server.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - string</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>processorCoreAmount</code>
//         <div>The number of logical CPU cores to allocate.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - int</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>memoryCapacity</code>
//         <div>The amount of memory to allocate in gigabytes.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - int</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>hourlyBillingFlag</code>
//         <div>Specifies the billing type for the server.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - boolean</li>
//             <li>When true the server will be billed on hourly usage, otherwise it will be billed on a monthly basis.</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>operatingSystemReferenceCode</code>
//         <div>An identifier for the operating system to provision the server with.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - string</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>datacenter.name</code>
//         <div>Specifies which datacenter the server is to be provisioned in.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - string</li>
//             <li>The <code>datacenter</code> property is a [[SoftLayer_Location (type)|location]] structure with the <code>name</code> field set.</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//         <http title="Example">{
//     "datacenter": {
//         "name": "dal05"
//     }
// }</http>
//         <br />
//     </li>
//     <li><code>networkComponents.maxSpeed</code>
//         <div>Specifies the connection speed for the server's network components.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - int</li>
//             <li><b>Default</b> - The highest available zero cost port speed will be used.</li>
//             <li><b>Description</b> - The <code>networkComponents</code> property is an array with a single [[SoftLayer_Network_Component (type)|network component]] structure. The <code>maxSpeed</code> property must be set to specify the network uplink speed, in megabits per second, of the server.</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//             <http title="Example">{
//     "networkComponents": [
//         {
//             "maxSpeed": 1000
//         }
//     ]
// }</http>
//         <br />
//     </li>
//     <li><code>networkComponents.redundancyEnabledFlag</code>
//         <div>Specifies whether or not the server's network components should be in redundancy groups.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - bool</li>
//             <li><b>Default</b> - <code>false</code></li>
//             <li><b>Description</b> - The <code>networkComponents</code> property is an array with a single [[SoftLayer_Network_Component (type)|network component]] structure. When the <code>redundancyEnabledFlag</code> property is true the server's network components will be in redundancy groups.</li>
//         </ul>
//             <http title="Example">{
//     "networkComponents": [
//         {
//             "redundancyEnabledFlag": false
//         }
//     ]
// }</http>
//         <br />
//     </li>
//     <li><code>privateNetworkOnlyFlag</code>
//         <div>Specifies whether or not the server only has access to the private network</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - boolean</li>
//             <li><b>Default</b> - <code>false</code></li>
//             <li>When true this flag specifies that a server is to only have access to the private network.</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>primaryNetworkComponent.networkVlan.id</code>
//         <div>Specifies the network vlan which is to be used for the frontend interface of the server.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - int</li>
//             <li><b>Description</b> - The <code>primaryNetworkComponent</code> property is a [[SoftLayer_Network_Component (type)|network component]] structure with the <code>networkVlan</code> property populated with a [[SoftLayer_Network_Vlan (type)|vlan]] structure. The <code>id</code> property must be set to specify the frontend network vlan of the server.</li>
//         </ul>
//         <http title="Example">{
//     "primaryNetworkComponent": {
//         "networkVlan": {
//             "id": 1
//         }
//     }
// }</http>
//         <br />
//     </li>
//     <li><code>primaryBackendNetworkComponent.networkVlan.id</code>
//         <div>Specifies the network vlan which is to be used for the backend interface of the server.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - int</li>
//             <li><b>Description</b> - The <code>primaryBackendNetworkComponent</code> property is a [[SoftLayer_Network_Component (type)|network component]] structure with the <code>networkVlan</code> property populated with a [[SoftLayer_Network_Vlan (type)|vlan]] structure. The <code>id</code> property must be set to specify the backend network vlan of the server.</li>
//         </ul>
//         <http title="Example">{
//     "primaryBackendNetworkComponent": {
//         "networkVlan": {
//             "id": 2
//         }
//     }
// }</http>
//         <br />
//     </li>
//     <li><code>fixedConfigurationPreset.keyName</code>
//         <div></div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - string</li>
//             <li><b>Description</b> - The <code>fixedConfigurationPreset</code> property is a [[SoftLayer_Product_Package_Preset (type)|fixed configuration preset]] structure. The <code>keyName</code> property must be set to specify preset to use.</li>
//             <li>If a fixed configuration preset is used <code>processorCoreAmount</code>, <code>memoryCapacity</code> and <code>hardDrives</code> properties must not be set.</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//         <http title="Example">{
//     "fixedConfigurationPreset": {
//         "keyName": "SOME_KEY_NAME"
//     }
// }</http>
//         <br />
//     </li>
//     <li><code>userData.value</code>
//         <div>Arbitrary data to be made available to the server.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - string</li>
//             <li><b>Description</b> - The <code>userData</code> property is an array with a single [[SoftLayer_Hardware_Attribute (type)|attribute]] structure with the <code>value</code> property set to an arbitrary value.</li>
//             <li>This value can be retrieved via the [[SoftLayer_Resource_Metadata/getUserMetadata|getUserMetadata]] method from a request originating from the server. This is primarily useful for providing data to software that may be on the server and configured to execute upon first boot.</li>
//         </ul>
//         <http title="Example">{
//     "userData": [
//         {
//             "value": "someValue"
//         }
//     ]
// }</http>
//         <br />
//     </li>
//     <li><code>hardDrives</code>
//         <div>Hard drive settings for the server</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - SoftLayer_Hardware_Component</li>
//             <li><b>Default</b> - The largest available capacity for a zero cost primary disk will be used.</li>
//             <li><b>Description</b> - The <code>hardDrives</code> property is an array of [[SoftLayer_Hardware_Component (type)|hardware component]] structures.</i>
//             <li>Each hard drive must specify the <code>capacity</code> property.</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//         <http title="Example">{
//     "hardDrives": [
//         {
//             "capacity": 500
//         }
//     ]
// }</http>
//         <br />
//     </li>
//     <li id="hardware-create-object-ssh-keys"><code>sshKeys</code>
//         <div>SSH keys to install on the server upon provisioning.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - array of [[SoftLayer_Security_Ssh_Key (type)|SoftLayer_Security_Ssh_Key]]</li>
//             <li><b>Description</b> - The <code>sshKeys</code> property is an array of [[SoftLayer_Security_Ssh_Key (type)|SSH Key]] structures with the <code>id</code> property set to the value of an existing SSH key.</li>
//             <li>To create a new SSH key, call [[SoftLayer_Security_Ssh_Key/createObject|createObject]] on the [[SoftLayer_Security_Ssh_Key]] service.</li>
//             <li>To obtain a list of existing SSH keys, call [[SoftLayer_Account/getSshKeys|getSshKeys]] on the [[SoftLayer_Account]] service.
//         </ul>
//         <http title="Example">{
//     "sshKeys": [
//         {
//             "id": 123
//         }
//     ]
// }</http>
//         <br />
//     </li>
//     <li><code>postInstallScriptUri</code>
//         <div>Specifies the uri location of the script to be downloaded and run after installation is complete.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - string</li>
//         </ul>
//         <br />
//     </li>
// </ul>
//
//
// <h1>REST Example</h1>
// <http title="Request">curl -X POST -d '{
//  "parameters":[
//      {
//          "hostname": "host1",
//          "domain": "example.com",
//          "processorCoreAmount": 2,
//          "memoryCapacity": 2,
//          "hourlyBillingFlag": true,
//          "operatingSystemReferenceCode": "UBUNTU_LATEST"
//      }
//  ]
// }' https://api.softlayer.com/rest/v3/SoftLayer_Hardware.json
// </http>
// <http title="Response">HTTP/1.1 201 Created
// Location: https://api.softlayer.com/rest/v3/SoftLayer_Hardware/f5a3fcff-db1d-4b7c-9fa0-0349e41c29c5/getObject
//
//
// {
//     "accountId": 232298,
//     "bareMetalInstanceFlag": null,
//     "domain": "example.com",
//     "hardwareStatusId": null,
//     "hostname": "host1",
//     "id": null,
//     "serviceProviderId": null,
//     "serviceProviderResourceId": null,
//     "globalIdentifier": "f5a3fcff-db1d-4b7c-9fa0-0349e41c29c5",
//     "hourlyBillingFlag": true,
//     "memoryCapacity": 2,
//     "operatingSystemReferenceCode": "UBUNTU_LATEST",
//     "processorCoreAmount": 2
// }
// </http>
func (r Hardware_SecurityModule) CreateObject(templateObject *datatypes.Hardware_SecurityModule) (resp datatypes.Hardware_SecurityModule, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "createObject", params, r.Session, &r.Options, &resp)
	return
}

// no documentation yet
func (r Hardware_SecurityModule) CreatePostSoftwareInstallTransaction(installCodes []string, returnBoolean *bool) (resp bool, err error) {
	params := []interface{}{
		installCodes,
		returnBoolean,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "createPostSoftwareInstallTransaction", params, r.Session, &r.Options, &resp)
	return
}

//
// This method will cancel a server effective immediately. For servers billed hourly, the charges will stop immediately after the method returns.
func (r Hardware_SecurityModule) DeleteObject() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "deleteObject", nil, r.Session, &r.Options, &resp)
	return
}

// Delete software component passwords.
func (r Hardware_SecurityModule) DeleteSoftwareComponentPasswords(softwareComponentPasswords []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		softwareComponentPasswords,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "deleteSoftwareComponentPasswords", params, r.Session, &r.Options, &resp)
	return
}

// Edit a server's properties
func (r Hardware_SecurityModule) EditObject(templateObject *datatypes.Hardware_Server) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "editObject", params, r.Session, &r.Options, &resp)
	return
}

// Edit the properties of a software component password such as the username, password, and notes.
func (r Hardware_SecurityModule) EditSoftwareComponentPasswords(softwareComponentPasswords []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		softwareComponentPasswords,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "editSoftwareComponentPasswords", params, r.Session, &r.Options, &resp)
	return
}

// Download and run remote script from uri on the hardware.
func (r Hardware_SecurityModule) ExecuteRemoteScript(uri *string) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		uri,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "executeRemoteScript", params, r.Session, &r.Options, &resp)
	return
}

// The '''findByIpAddress''' method finds hardware using its primary public or private IP address. IP addresses that have a secondary subnet tied to the hardware will not return the hardware - alternate means of locating the hardware must be used (see '''Associated Methods'''). If no hardware is found, no errors are generated and no data is returned.
func (r Hardware_SecurityModule) FindByIpAddress(ipAddress *string) (resp datatypes.Hardware, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "findByIpAddress", params, r.Session, &r.Options, &resp)
	return
}

//
// Obtain an [[SoftLayer_Container_Product_Order_Hardware_Server (type)|order container]] that can be sent to [[SoftLayer_Product_Order/verifyOrder|verifyOrder]] or [[SoftLayer_Product_Order/placeOrder|placeOrder]].
//
//
// This is primarily useful when there is a necessity to confirm the price which will be charged for an order.
//
//
// See [[SoftLayer_Hardware/createObject|createObject]] for specifics on the requirements of the template object parameter.
func (r Hardware_SecurityModule) GenerateOrderTemplate(templateObject *datatypes.Hardware) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "generateOrderTemplate", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The account associated with a piece of hardware.
func (r Hardware_SecurityModule) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getAccount", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's active physical components.
func (r Hardware_SecurityModule) GetActiveComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getActiveComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The billing item for a server's attached network firewall.
func (r Hardware_SecurityModule) GetActiveNetworkFirewallBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getActiveNetworkFirewallBillingItem", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's active network monitoring incidents.
func (r Hardware_SecurityModule) GetActiveNetworkMonitorIncident() (resp []datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getActiveNetworkMonitorIncident", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_SecurityModule) GetActiveTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getActiveTickets", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Transaction currently running for server.
func (r Hardware_SecurityModule) GetActiveTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getActiveTransaction", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Any active transaction(s) that are currently running for the server (example: os reload).
func (r Hardware_SecurityModule) GetActiveTransactions() (resp []datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getActiveTransactions", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getAlarmHistory''' method retrieves a detailed history for the monitoring alarm. When calling this method, a start and end date for the history to be retrieved must be entered.
func (r Hardware_SecurityModule) GetAlarmHistory(startDate *datatypes.Time, endDate *datatypes.Time, alarmId *string) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
		alarmId,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getAlarmHistory", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_SecurityModule) GetAllPowerComponents() (resp []datatypes.Hardware_Power_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getAllPowerComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Storage_Allowed_Host information to connect this server to Network Storage volumes that require access control lists.
func (r Hardware_SecurityModule) GetAllowedHost() (resp datatypes.Network_Storage_Allowed_Host, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getAllowedHost", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Storage objects that this SoftLayer_Hardware has access to.
func (r Hardware_SecurityModule) GetAllowedNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getAllowedNetworkStorage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Storage objects whose Replica that this SoftLayer_Hardware has access to.
func (r Hardware_SecurityModule) GetAllowedNetworkStorageReplicas() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getAllowedNetworkStorageReplicas", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding an antivirus/spyware software component object.
func (r Hardware_SecurityModule) GetAntivirusSpywareSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getAntivirusSpywareSoftwareComponent", nil, r.Session, &r.Options, &resp)
	return
}

// This method is retrieve a list of SoftLayer_Network_Storage volumes that are authorized access to this SoftLayer_Hardware.
func (r Hardware_SecurityModule) GetAttachedNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getAttachedNetworkStorages", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's specific attributes.
func (r Hardware_SecurityModule) GetAttributes() (resp []datatypes.Hardware_Attribute, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getAttributes", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An object that stores the maximum level for the monitoring query types and response types.
func (r Hardware_SecurityModule) GetAvailableMonitoring() (resp []datatypes.Network_Monitor_Version1_Query_Host_Stratum, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getAvailableMonitoring", nil, r.Session, &r.Options, &resp)
	return
}

// This method retrieves a list of SoftLayer_Network_Storage volumes that can be authorized to this SoftLayer_Hardware.
func (r Hardware_SecurityModule) GetAvailableNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getAvailableNetworkStorages", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The average daily total bandwidth usage for the current billing cycle.
func (r Hardware_SecurityModule) GetAverageDailyBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getAverageDailyBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The average daily private bandwidth usage for the current billing cycle.
func (r Hardware_SecurityModule) GetAverageDailyPrivateBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getAverageDailyPrivateBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The average daily public bandwidth usage for the current billing cycle.
func (r Hardware_SecurityModule) GetAverageDailyPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getAverageDailyPublicBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Use this method to return an array of private bandwidth utilization records between a given date range.
//
// This method represents the NEW version of getFrontendBandwidthUse
func (r Hardware_SecurityModule) GetBackendBandwidthUsage(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getBackendBandwidthUsage", params, r.Session, &r.Options, &resp)
	return
}

// Use this method to return an array of private bandwidth utilization records between a given date range.
func (r Hardware_SecurityModule) GetBackendBandwidthUse(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Network_Bandwidth_Version1_Usage_Detail, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getBackendBandwidthUse", params, r.Session, &r.Options, &resp)
	return
}

// The '''getBackendIncomingBandwidth''' method retrieves the amount of incoming private network traffic used between the given start date and end date parameters. When entering start and end dates, only the month, day and year are used to calculate bandwidth totals - the time (HH:MM:SS) is ignored and defaults to midnight. The amount of bandwidth retrieved is measured in gigabytes.
func (r Hardware_SecurityModule) GetBackendIncomingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getBackendIncomingBandwidth", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's back-end or private network components.
func (r Hardware_SecurityModule) GetBackendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getBackendNetworkComponents", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getBackendOutgoingBandwidth''' method retrieves the amount of outgoing private network traffic used between the given start date and end date parameters. When entering start and end dates, only the month, day and year are used to calculate bandwidth totals - the time (HH:MM:SS) is ignored and defaults to midnight. The amount of bandwidth retrieved is measured in gigabytes.
func (r Hardware_SecurityModule) GetBackendOutgoingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getBackendOutgoingBandwidth", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's backend or private router.
func (r Hardware_SecurityModule) GetBackendRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getBackendRouters", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's allotted bandwidth (measured in GB).
func (r Hardware_SecurityModule) GetBandwidthAllocation() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getBandwidthAllocation", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's allotted detail record. Allotment details link bandwidth allocation with allotments.
func (r Hardware_SecurityModule) GetBandwidthAllotmentDetail() (resp datatypes.Network_Bandwidth_Version1_Allotment_Detail, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getBandwidthAllotmentDetail", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a collection of bandwidth data from an individual public or private network tracking object. Data is ideal if you with to employ your own traffic storage and graphing systems.
func (r Hardware_SecurityModule) GetBandwidthForDateRange(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getBandwidthForDateRange", params, r.Session, &r.Options, &resp)
	return
}

// Use this method when needing a bandwidth image for a single server.  It will gather the correct input parameters for the generic graphing utility automatically based on the snapshot specified.  Use the $draw flag to suppress the generation of the actual binary PNG image.
func (r Hardware_SecurityModule) GetBandwidthImage(networkType *string, snapshotRange *string, draw *bool, dateSpecified *datatypes.Time, dateSpecifiedEnd *datatypes.Time) (resp datatypes.Container_Bandwidth_GraphOutputs, err error) {
	params := []interface{}{
		networkType,
		snapshotRange,
		draw,
		dateSpecified,
		dateSpecifiedEnd,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getBandwidthImage", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's benchmark certifications.
func (r Hardware_SecurityModule) GetBenchmarkCertifications() (resp []datatypes.Hardware_Benchmark_Certification, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getBenchmarkCertifications", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The raw bandwidth usage data for the current billing cycle. One object will be returned for each network this server is attached to.
func (r Hardware_SecurityModule) GetBillingCycleBandwidthUsage() (resp []datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getBillingCycleBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The raw private bandwidth usage data for the current billing cycle.
func (r Hardware_SecurityModule) GetBillingCyclePrivateBandwidthUsage() (resp datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getBillingCyclePrivateBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The raw public bandwidth usage data for the current billing cycle.
func (r Hardware_SecurityModule) GetBillingCyclePublicBandwidthUsage() (resp datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getBillingCyclePublicBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the billing item for a server.
func (r Hardware_SecurityModule) GetBillingItem() (resp datatypes.Billing_Item_Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getBillingItem", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A flag indicating that a billing item exists.
func (r Hardware_SecurityModule) GetBillingItemFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getBillingItemFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Determines whether the hardware is ineligible for cancellation because it is disconnected.
func (r Hardware_SecurityModule) GetBlockCancelBecauseDisconnectedFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getBlockCancelBecauseDisconnectedFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Status indicating whether or not a piece of hardware has business continuance insurance.
func (r Hardware_SecurityModule) GetBusinessContinuanceInsuranceFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getBusinessContinuanceInsuranceFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's components.
func (r Hardware_SecurityModule) GetComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_SecurityModule) GetContainsSolidStateDrivesFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getContainsSolidStateDrivesFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A continuous data protection/server backup software component object.
func (r Hardware_SecurityModule) GetContinuousDataProtectionSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getContinuousDataProtectionSoftwareComponent", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A server's control panel.
func (r Hardware_SecurityModule) GetControlPanel() (resp datatypes.Software_Component_ControlPanel, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getControlPanel", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total cost of a server, measured in US Dollars ($USD).
func (r Hardware_SecurityModule) GetCost() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getCost", nil, r.Session, &r.Options, &resp)
	return
}

//
// There are many options that may be provided while ordering a server, this method can be used to determine what these options are.
//
//
// Detailed information on the return value can be found on the data type page for [[SoftLayer_Container_Hardware_Configuration (type)]].
func (r Hardware_SecurityModule) GetCreateObjectOptions() (resp datatypes.Container_Hardware_Configuration, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getCreateObjectOptions", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An object that provides commonly used bandwidth summary components for the current billing cycle.
func (r Hardware_SecurityModule) GetCurrentBandwidthSummary() (resp datatypes.Metric_Tracking_Object_Bandwidth_Summary, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getCurrentBandwidthSummary", nil, r.Session, &r.Options, &resp)
	return
}

// Attempt to retrieve the file associated with the current benchmark certification result, if such a file exists.  If there is no file for this benchmark certification result, calling this method throws an exception.
func (r Hardware_SecurityModule) GetCurrentBenchmarkCertificationResultFile() (resp []byte, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getCurrentBenchmarkCertificationResultFile", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The current billable public outbound bandwidth for this hardware for the current billing cycle.
func (r Hardware_SecurityModule) GetCurrentBillableBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getCurrentBillableBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Get the billing detail for this instance for the current billing period. This does not include bandwidth usage.
func (r Hardware_SecurityModule) GetCurrentBillingDetail() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getCurrentBillingDetail", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getCurrentBillingTotal''' method retrieves the total bill amount in US Dollars ($) for the current billing period. In addition to the total bill amount, the billing detail also includes all bandwidth used up to the point the method is called on the piece of hardware.
func (r Hardware_SecurityModule) GetCurrentBillingTotal() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getCurrentBillingTotal", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve bandwidth graph by date.
func (r Hardware_SecurityModule) GetCustomBandwidthDataByDate(graphData *datatypes.Container_Graph) (resp datatypes.Container_Graph, err error) {
	params := []interface{}{
		graphData,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getCustomBandwidthDataByDate", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Indicates if a server has a Customer Installed OS
func (r Hardware_SecurityModule) GetCustomerInstalledOperatingSystemFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getCustomerInstalledOperatingSystemFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Indicates if a server is a customer owned device.
func (r Hardware_SecurityModule) GetCustomerOwnedFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getCustomerOwnedFlag", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getDailyAverage''' method calculates the average daily network traffic used by the selected server. Using the required parameter ''dateTime'' to enter a start and end date, the user retrieves this average, measure in gigabytes (GB) for the specified date range. When entering parameters, only the month, day and year are required - time entries are omitted as this method defaults the time to midnight in order to account for the entire day.
func (r Hardware_SecurityModule) GetDailyAverage(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getDailyAverage", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the datacenter in which a piece of hardware resides.
func (r Hardware_SecurityModule) GetDatacenter() (resp datatypes.Location, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getDatacenter", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The name of the datacenter in which a piece of hardware resides.
func (r Hardware_SecurityModule) GetDatacenterName() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getDatacenterName", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All hardware that has uplink network connections to a piece of hardware.
func (r Hardware_SecurityModule) GetDownlinkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getDownlinkHardware", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All hardware that has uplink network connections to a piece of hardware.
func (r Hardware_SecurityModule) GetDownlinkNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getDownlinkNetworkHardware", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all servers attached to a piece of network hardware.
func (r Hardware_SecurityModule) GetDownlinkServers() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getDownlinkServers", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all virtual guests attached to a piece of network hardware.
func (r Hardware_SecurityModule) GetDownlinkVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getDownlinkVirtualGuests", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All hardware downstream from a network device.
func (r Hardware_SecurityModule) GetDownstreamHardwareBindings() (resp []datatypes.Network_Component_Uplink_Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getDownstreamHardwareBindings", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All network hardware downstream from the selected piece of hardware.
func (r Hardware_SecurityModule) GetDownstreamNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getDownstreamNetworkHardware", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All network hardware with monitoring warnings or errors that are downstream from the selected piece of hardware.
func (r Hardware_SecurityModule) GetDownstreamNetworkHardwareWithIncidents() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getDownstreamNetworkHardwareWithIncidents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all servers attached downstream to a piece of network hardware.
func (r Hardware_SecurityModule) GetDownstreamServers() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getDownstreamServers", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all virtual guests attached to a piece of network hardware.
func (r Hardware_SecurityModule) GetDownstreamVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getDownstreamVirtualGuests", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The drive controllers contained within a piece of hardware.
func (r Hardware_SecurityModule) GetDriveControllers() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getDriveControllers", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's associated EVault network storage service account.
func (r Hardware_SecurityModule) GetEvaultNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getEvaultNetworkStorage", nil, r.Session, &r.Options, &resp)
	return
}

// Get the subnets associated with this server that are protectable by a network component firewall.
func (r Hardware_SecurityModule) GetFirewallProtectableSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getFirewallProtectableSubnets", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's firewall services.
func (r Hardware_SecurityModule) GetFirewallServiceComponent() (resp datatypes.Network_Component_Firewall, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getFirewallServiceComponent", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Defines the fixed components in a fixed configuration bare metal server.
func (r Hardware_SecurityModule) GetFixedConfigurationPreset() (resp datatypes.Product_Package_Preset, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getFixedConfigurationPreset", nil, r.Session, &r.Options, &resp)
	return
}

// Use this method to return an array of public bandwidth utilization records between a given date range.
//
// This method represents the NEW version of getFrontendBandwidthUse
func (r Hardware_SecurityModule) GetFrontendBandwidthUsage(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getFrontendBandwidthUsage", params, r.Session, &r.Options, &resp)
	return
}

// Use this method to return an array of public bandwidth utilization records between a given date range.
func (r Hardware_SecurityModule) GetFrontendBandwidthUse(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Network_Bandwidth_Version1_Usage_Detail, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getFrontendBandwidthUse", params, r.Session, &r.Options, &resp)
	return
}

// The '''getFrontendIncomingBandwidth''' method retrieves the amount of incoming public network traffic used by a server between the given start and end date parameters. When entering the ''dateTime'' parameter, only the month, day and year of the start and end dates are required - the time (hour, minute and second) are set to midnight by default and cannot be changed. The amount of bandwidth retrieved is measured in gigabytes (GB).
func (r Hardware_SecurityModule) GetFrontendIncomingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getFrontendIncomingBandwidth", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's front-end or public network components.
func (r Hardware_SecurityModule) GetFrontendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getFrontendNetworkComponents", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getFrontendOutgoingBandwidth''' method retrieves the amount of outgoing public network traffic used by a server between the given start and end date parameters. The ''dateTime'' parameter requires only the day, month and year to be entered - the time (hour, minute and second) are set to midnight be default in order to gather the data for the entire start and end date indicated in the parameter. The amount of bandwidth retrieved is measured in gigabytes (GB).
func (r Hardware_SecurityModule) GetFrontendOutgoingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getFrontendOutgoingBandwidth", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's frontend or public router.
func (r Hardware_SecurityModule) GetFrontendRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getFrontendRouters", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's universally unique identifier.
func (r Hardware_SecurityModule) GetGlobalIdentifier() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getGlobalIdentifier", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hard drives contained within a piece of hardware.
func (r Hardware_SecurityModule) GetHardDrives() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getHardDrives", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a server by searching for the primary IP address.
func (r Hardware_SecurityModule) GetHardwareByIpAddress(ipAddress *string) (resp datatypes.Hardware_Server, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getHardwareByIpAddress", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The chassis that a piece of hardware is housed in.
func (r Hardware_SecurityModule) GetHardwareChassis() (resp datatypes.Hardware_Chassis, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getHardwareChassis", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's function.
func (r Hardware_SecurityModule) GetHardwareFunction() (resp datatypes.Hardware_Function, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getHardwareFunction", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's function.
func (r Hardware_SecurityModule) GetHardwareFunctionDescription() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getHardwareFunctionDescription", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's status.
func (r Hardware_SecurityModule) GetHardwareStatus() (resp datatypes.Hardware_Status, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getHardwareStatus", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Determine in hardware object has TPM enabled.
func (r Hardware_SecurityModule) GetHasTrustedPlatformModuleBillingItemFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getHasTrustedPlatformModuleBillingItemFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a host IPS software component object.
func (r Hardware_SecurityModule) GetHostIpsSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getHostIpsSoftwareComponent", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getHourlyBandwidth''' method retrieves all bandwidth updates hourly for the specified hardware. Because the potential number of data points can become excessive, the method limits the user to obtain data in 24-hour intervals. The required ''dateTime'' parameter is used as the starting point for the query and will be calculated for the 24-hour period starting with the specified date and time. For example, entering a parameter of
//
// '02/01/2008 0:00'
//
// results in a return of all bandwidth data for the entire day of February 1, 2008, as 0:00 specifies a midnight start date. Please note that the time entered should be completed using a 24-hour clock (military time, astronomical time).
//
// For data spanning more than a single 24-hour period, refer to the getBandwidthData function on the metricTrackingObject for the piece of hardware.
func (r Hardware_SecurityModule) GetHourlyBandwidth(mode *string, day *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		mode,
		day,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getHourlyBandwidth", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A server's hourly billing status.
func (r Hardware_SecurityModule) GetHourlyBillingFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getHourlyBillingFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The sum of all the inbound network traffic data for the last 30 days.
func (r Hardware_SecurityModule) GetInboundBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getInboundBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total private inbound bandwidth for this hardware for the current billing cycle.
func (r Hardware_SecurityModule) GetInboundPrivateBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getInboundPrivateBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total public inbound bandwidth for this hardware for the current billing cycle.
func (r Hardware_SecurityModule) GetInboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getInboundPublicBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Return a collection of SoftLayer_Item_Price objects from a collection of SoftLayer_Software_Description
func (r Hardware_SecurityModule) GetItemPricesFromSoftwareDescriptions(softwareDescriptions []datatypes.Software_Description, includeTranslationsFlag *bool, returnAllPricesFlag *bool) (resp []datatypes.Product_Item, err error) {
	params := []interface{}{
		softwareDescriptions,
		includeTranslationsFlag,
		returnAllPricesFlag,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getItemPricesFromSoftwareDescriptions", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The last transaction that a server's operating system was loaded.
func (r Hardware_SecurityModule) GetLastOperatingSystemReload() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getLastOperatingSystemReload", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the last transaction a server performed.
func (r Hardware_SecurityModule) GetLastTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getLastTransaction", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's latest network monitoring incident.
func (r Hardware_SecurityModule) GetLatestNetworkMonitorIncident() (resp datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getLatestNetworkMonitorIncident", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Where a piece of hardware is located within SoftLayer's location hierarchy.
func (r Hardware_SecurityModule) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getLocation", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_SecurityModule) GetLocationPathString() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getLocationPathString", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a lockbox account associated with a server.
func (r Hardware_SecurityModule) GetLockboxNetworkStorage() (resp datatypes.Network_Storage, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getLockboxNetworkStorage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A flag indicating that the hardware is a managed resource.
func (r Hardware_SecurityModule) GetManagedResourceFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getManagedResourceFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the remote management network component attached with this server.
func (r Hardware_SecurityModule) GetManagementNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getManagementNetworkComponent", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's memory.
func (r Hardware_SecurityModule) GetMemory() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getMemory", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The amount of memory a piece of hardware has, measured in gigabytes.
func (r Hardware_SecurityModule) GetMemoryCapacity() (resp uint, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getMemoryCapacity", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's metric tracking object.
func (r Hardware_SecurityModule) GetMetricTrackingObject() (resp datatypes.Metric_Tracking_Object_HardwareServer, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getMetricTrackingObject", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The metric tracking object id for this server.
func (r Hardware_SecurityModule) GetMetricTrackingObjectId() (resp int, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getMetricTrackingObjectId", nil, r.Session, &r.Options, &resp)
	return
}

// Returns open monitoring alarms for a given time period
func (r Hardware_SecurityModule) GetMonitoringActiveAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getMonitoringActiveAlarms", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the monitoring agents associated with a piece of hardware.
func (r Hardware_SecurityModule) GetMonitoringAgents() (resp []datatypes.Monitoring_Agent, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getMonitoringAgents", nil, r.Session, &r.Options, &resp)
	return
}

// Returns closed monitoring alarms for a given time period
func (r Hardware_SecurityModule) GetMonitoringClosedAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getMonitoringClosedAlarms", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the hardware's monitoring robot.
func (r Hardware_SecurityModule) GetMonitoringRobot() (resp datatypes.Monitoring_Robot, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getMonitoringRobot", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's network monitoring services.
func (r Hardware_SecurityModule) GetMonitoringServiceComponent() (resp datatypes.Network_Monitor_Version1_Query_Host_Stratum, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getMonitoringServiceComponent", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The monitoring service flag eligibility status for a piece of hardware.
func (r Hardware_SecurityModule) GetMonitoringServiceEligibilityFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getMonitoringServiceEligibilityFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The service flag status for a piece of hardware.
func (r Hardware_SecurityModule) GetMonitoringServiceFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getMonitoringServiceFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The monitoring notification objects for this hardware. Each object links this hardware instance to a user account that will be notified if monitoring on this hardware object fails
func (r Hardware_SecurityModule) GetMonitoringUserNotification() (resp []datatypes.User_Customer_Notification_Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getMonitoringUserNotification", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's motherboard.
func (r Hardware_SecurityModule) GetMotherboard() (resp datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getMotherboard", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's network cards.
func (r Hardware_SecurityModule) GetNetworkCards() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getNetworkCards", nil, r.Session, &r.Options, &resp)
	return
}

// Get the IP addresses associated with this server that are protectable by a network component firewall. Note, this may not return all values for IPv6 subnets for this server. Please use getFirewallProtectableSubnets to get all protectable subnets.
func (r Hardware_SecurityModule) GetNetworkComponentFirewallProtectableIpAddresses() (resp []datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getNetworkComponentFirewallProtectableIpAddresses", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Returns a hardware's network components.
func (r Hardware_SecurityModule) GetNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getNetworkComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The gateway member if this device is part of a network gateway.
func (r Hardware_SecurityModule) GetNetworkGatewayMember() (resp datatypes.Network_Gateway_Member, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getNetworkGatewayMember", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Whether or not this device is part of a network gateway.
func (r Hardware_SecurityModule) GetNetworkGatewayMemberFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getNetworkGatewayMemberFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's network management IP address.
func (r Hardware_SecurityModule) GetNetworkManagementIpAddress() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getNetworkManagementIpAddress", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All servers with failed monitoring that are attached downstream to a piece of hardware.
func (r Hardware_SecurityModule) GetNetworkMonitorAttachedDownHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getNetworkMonitorAttachedDownHardware", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Virtual guests that are attached downstream to a hardware that have failed monitoring
func (r Hardware_SecurityModule) GetNetworkMonitorAttachedDownVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getNetworkMonitorAttachedDownVirtualGuests", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The status of all of a piece of hardware's network monitoring incidents.
func (r Hardware_SecurityModule) GetNetworkMonitorIncidents() (resp []datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getNetworkMonitorIncidents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's network monitors.
func (r Hardware_SecurityModule) GetNetworkMonitors() (resp []datatypes.Network_Monitor_Version1_Query_Host, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getNetworkMonitors", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The value of a hardware's network status attribute.
func (r Hardware_SecurityModule) GetNetworkStatus() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getNetworkStatus", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware's related network status attribute.
func (r Hardware_SecurityModule) GetNetworkStatusAttribute() (resp datatypes.Hardware_Attribute, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getNetworkStatusAttribute", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's associated network storage service account.
func (r Hardware_SecurityModule) GetNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getNetworkStorage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The network virtual LANs (VLANs) associated with a piece of hardware's network components.
func (r Hardware_SecurityModule) GetNetworkVlans() (resp []datatypes.Network_Vlan, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getNetworkVlans", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's allotted bandwidth for the next billing cycle (measured in GB).
func (r Hardware_SecurityModule) GetNextBillingCycleBandwidthAllocation() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getNextBillingCycleBandwidthAllocation", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_SecurityModule) GetNotesHistory() (resp []datatypes.Hardware_Note, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getNotesHistory", nil, r.Session, &r.Options, &resp)
	return
}

// no documentation yet
func (r Hardware_SecurityModule) GetObject() (resp datatypes.Hardware_SecurityModule, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getObject", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An open ticket requesting cancellation of this server, if one exists.
func (r Hardware_SecurityModule) GetOpenCancellationTicket() (resp datatypes.Ticket, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getOpenCancellationTicket", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's operating system.
func (r Hardware_SecurityModule) GetOperatingSystem() (resp datatypes.Software_Component_OperatingSystem, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getOperatingSystem", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's operating system software description.
func (r Hardware_SecurityModule) GetOperatingSystemReferenceCode() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getOperatingSystemReferenceCode", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The sum of all the outbound network traffic data for the last 30 days.
func (r Hardware_SecurityModule) GetOutboundBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getOutboundBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total private outbound bandwidth for this hardware for the current billing cycle.
func (r Hardware_SecurityModule) GetOutboundPrivateBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getOutboundPrivateBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total public outbound bandwidth for this hardware for the current billing cycle.
func (r Hardware_SecurityModule) GetOutboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getOutboundPublicBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Whether the bandwidth usage for this hardware for the current billing cycle exceeds the allocation.
func (r Hardware_SecurityModule) GetOverBandwidthAllocationFlag() (resp int, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getOverBandwidthAllocationFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a server's hardware state via its internal sensors. Remote sensor data is transmitted to the SoftLayer API by way of the server's remote management card. Sensor data measures system temperatures, voltages, and other local server settings. Sensor data is cached for 30 seconds. Calls made to getSensorData for the same server within 30 seconds of each other will return the same data. Subsequent calls will return new data once the cache expires.
func (r Hardware_SecurityModule) GetPMInfo() (resp []datatypes.Container_RemoteManagement_PmInfo, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getPMInfo", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the Point of Presence (PoP) location in which a piece of hardware resides.
func (r Hardware_SecurityModule) GetPointOfPresenceLocation() (resp datatypes.Location, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getPointOfPresenceLocation", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The power components for a hardware object.
func (r Hardware_SecurityModule) GetPowerComponents() (resp []datatypes.Hardware_Power_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getPowerComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's power supply.
func (r Hardware_SecurityModule) GetPowerSupply() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getPowerSupply", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware's primary private IP address.
func (r Hardware_SecurityModule) GetPrimaryBackendIpAddress() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getPrimaryBackendIpAddress", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the hardware's primary back-end network component.
func (r Hardware_SecurityModule) GetPrimaryBackendNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getPrimaryBackendNetworkComponent", nil, r.Session, &r.Options, &resp)
	return
}

// no documentation yet
func (r Hardware_SecurityModule) GetPrimaryDriveSize() (resp int, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getPrimaryDriveSize", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware's primary public IP address.
func (r Hardware_SecurityModule) GetPrimaryIpAddress() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getPrimaryIpAddress", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the hardware's primary public network component.
func (r Hardware_SecurityModule) GetPrimaryNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getPrimaryNetworkComponent", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a graph of a server's private network bandwidth usage over the specified timeframe. If no timeframe is specified then getPublicBandwidthGraphImage retrieves the last 24 hours of public bandwidth usage. getPrivateBandwidthGraphImage returns a PNG image measuring 827 pixels by 293 pixels.
func (r Hardware_SecurityModule) GetPrivateBandwidthData(startTime *int, endTime *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getPrivateBandwidthData", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve a brief summary of a server's private network bandwidth usage. getPrivateBandwidthDataSummary retrieves a server's bandwidth allocation for its billing period, its estimated usage during its billing period, and an estimation of how much bandwidth it will use during its billing period based on its current usage. A server's projected bandwidth usage increases in accuracy as it progresses through its billing period.
func (r Hardware_SecurityModule) GetPrivateBandwidthDataSummary() (resp datatypes.Container_Network_Bandwidth_Data_Summary, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getPrivateBandwidthDataSummary", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a graph of a server's private network bandwidth usage over the specified time frame. If no time frame is specified then getPublicBandwidthGraphImage retrieves the last 24 hours of public bandwidth usage. getPublicBandwidthGraphImage returns a PNG image
func (r Hardware_SecurityModule) GetPrivateBandwidthGraphImage(startTime *string, endTime *string) (resp []byte, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getPrivateBandwidthGraphImage", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A server's primary private IP address.
func (r Hardware_SecurityModule) GetPrivateIpAddress() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getPrivateIpAddress", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the private network component attached with this server.
func (r Hardware_SecurityModule) GetPrivateNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getPrivateNetworkComponent", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Whether the hardware only has access to the private network.
func (r Hardware_SecurityModule) GetPrivateNetworkOnlyFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getPrivateNetworkOnlyFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the backend VLAN for the primary IP address of the server
func (r Hardware_SecurityModule) GetPrivateVlan() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getPrivateVlan", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a backend network VLAN by searching for an IP address
func (r Hardware_SecurityModule) GetPrivateVlanByIpAddress(ipAddress *string) (resp datatypes.Network_Vlan, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getPrivateVlanByIpAddress", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total number of processor cores, summed from all processors that are attached to a piece of hardware
func (r Hardware_SecurityModule) GetProcessorCoreAmount() (resp uint, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getProcessorCoreAmount", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total number of physical processor cores, summed from all processors that are attached to a piece of hardware
func (r Hardware_SecurityModule) GetProcessorPhysicalCoreAmount() (resp uint, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getProcessorPhysicalCoreAmount", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's processors.
func (r Hardware_SecurityModule) GetProcessors() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getProcessors", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Whether the bandwidth usage for this hardware for the current billing cycle is projected to exceed the allocation.
func (r Hardware_SecurityModule) GetProjectedOverBandwidthAllocationFlag() (resp int, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getProjectedOverBandwidthAllocationFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The projected public outbound bandwidth for this hardware for the current billing cycle.
func (r Hardware_SecurityModule) GetProjectedPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getProjectedPublicBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// no documentation yet
func (r Hardware_SecurityModule) GetProvisionDate() (resp datatypes.Time, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getProvisionDate", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a graph of a server's public network bandwidth usage over the specified timeframe. If no timeframe is specified then getPublicBandwidthGraphImage retrieves the last 24 hours of public bandwidth usage. getPublicBandwidthGraphImage returns a PNG image measuring 827 pixels by 293 pixels.
func (r Hardware_SecurityModule) GetPublicBandwidthData(startTime *int, endTime *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getPublicBandwidthData", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve a brief summary of a server's public network bandwidth usage. getPublicBandwidthDataSummary retrieves a server's bandwidth allocation for its billing period, its estimated usage during its billing period, and an estimation of how much bandwidth it will use during its billing period based on its current usage. A server's projected bandwidth usage increases in accuracy as it progresses through its billing period.
func (r Hardware_SecurityModule) GetPublicBandwidthDataSummary() (resp datatypes.Container_Network_Bandwidth_Data_Summary, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getPublicBandwidthDataSummary", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a graph of a server's public network bandwidth usage over the specified time frame. If no time frame is specified then getPublicBandwidthGraphImage retrieves the last 24 hours of public bandwidth usage. getPublicBandwidthGraphImage returns a PNG image measuring 827 pixels by 293 pixels.  THIS METHOD GENERATES GRAPHS BASED ON THE NEW DATA WAREHOUSE REPOSITORY.
func (r Hardware_SecurityModule) GetPublicBandwidthGraphImage(startTime *datatypes.Time, endTime *datatypes.Time) (resp []byte, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getPublicBandwidthGraphImage", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve the total number of bytes used by a server over a specified time period via the data warehouse tracking objects for this hardware.
func (r Hardware_SecurityModule) GetPublicBandwidthTotal(startTime *int, endTime *int) (resp uint, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getPublicBandwidthTotal", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve a SoftLayer server's public network component. Some servers are only connected to the private network and may not have a public network component. In that case getPublicNetworkComponent returns a null object.
func (r Hardware_SecurityModule) GetPublicNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getPublicNetworkComponent", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the frontend VLAN for the primary IP address of the server
func (r Hardware_SecurityModule) GetPublicVlan() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getPublicVlan", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the frontend network Vlan by searching the hostname of a server
func (r Hardware_SecurityModule) GetPublicVlanByHostname(hostname *string) (resp datatypes.Network_Vlan, err error) {
	params := []interface{}{
		hostname,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getPublicVlanByHostname", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_SecurityModule) GetRack() (resp datatypes.Location, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getRack", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The RAID controllers contained within a piece of hardware.
func (r Hardware_SecurityModule) GetRaidControllers() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getRaidControllers", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Recent events that impact this hardware.
func (r Hardware_SecurityModule) GetRecentEvents() (resp []datatypes.Notification_Occurrence_Event, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getRecentEvents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The last five commands issued to the server's remote management card.
func (r Hardware_SecurityModule) GetRecentRemoteManagementCommands() (resp []datatypes.Hardware_Component_RemoteManagement_Command_Request, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getRecentRemoteManagementCommands", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_SecurityModule) GetRegionalInternetRegistry() (resp datatypes.Network_Regional_Internet_Registry, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getRegionalInternetRegistry", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A server's remote management card.
func (r Hardware_SecurityModule) GetRemoteManagement() (resp datatypes.Hardware_Component_RemoteManagement, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getRemoteManagement", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve User credentials to issue commands and/or interact with the server's remote management card.
func (r Hardware_SecurityModule) GetRemoteManagementAccounts() (resp []datatypes.Hardware_Component_RemoteManagement_User, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getRemoteManagementAccounts", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's associated remote management component. This is normally IPMI.
func (r Hardware_SecurityModule) GetRemoteManagementComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getRemoteManagementComponent", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve User(s) who have access to issue commands and/or interact with the server's remote management card.
func (r Hardware_SecurityModule) GetRemoteManagementUsers() (resp []datatypes.Hardware_Component_RemoteManagement_User, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getRemoteManagementUsers", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_SecurityModule) GetResourceGroupMemberReferences() (resp []datatypes.Resource_Group_Member, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getResourceGroupMemberReferences", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_SecurityModule) GetResourceGroupRoles() (resp []datatypes.Resource_Group_Role, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getResourceGroupRoles", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The resource groups in which this hardware is a member.
func (r Hardware_SecurityModule) GetResourceGroups() (resp []datatypes.Resource_Group, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getResourceGroups", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the reverse domain records associated with this server.
func (r Hardware_SecurityModule) GetReverseDomainRecords() (resp []datatypes.Dns_Domain, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getReverseDomainRecords", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's routers.
func (r Hardware_SecurityModule) GetRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getRouters", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Collection of scale assets this hardware corresponds to.
func (r Hardware_SecurityModule) GetScaleAssets() (resp []datatypes.Scale_Asset, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getScaleAssets", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's vulnerability scan requests.
func (r Hardware_SecurityModule) GetSecurityScanRequests() (resp []datatypes.Network_Security_Scanner_Request, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getSecurityScanRequests", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a server's hardware state via its internal sensors. Remote sensor data is transmitted to the SoftLayer API by way of the server's remote management card. Sensor data measures system temperatures, voltages, and other local server settings. Sensor data is cached for 30 seconds. Calls made to getSensorData for the same server within 30 seconds of each other will return the same data. Subsequent calls will return new data once the cache expires.
func (r Hardware_SecurityModule) GetSensorData() (resp []datatypes.Container_RemoteManagement_SensorReading, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getSensorData", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieves the raw data returned from the server's remote management card.  For more details of what is returned please refer to the getSensorData method.  Along with the raw data, graphs for the cpu and system temperatures and fan speeds are also returned.
func (r Hardware_SecurityModule) GetSensorDataWithGraphs() (resp datatypes.Container_RemoteManagement_SensorReadingsWithGraphs, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getSensorDataWithGraphs", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a server's hardware components, software, and network components. getServerDetails is an aggregation function that combines the results of [[SoftLayer_Hardware_Server::getComponents]], [[SoftLayer_Hardware_Server::getSoftware]], and [[SoftLayer_Hardware_Server::getNetworkComponents]] in a single container.
func (r Hardware_SecurityModule) GetServerDetails() (resp datatypes.Container_Hardware_Server_Details, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getServerDetails", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the server's fan speeds and displays them using tachometer graphs.  Data used to construct graphs is retrieved from the server's remote management card.  All graphs returned will have a title associated with it.
func (r Hardware_SecurityModule) GetServerFanSpeedGraphs() (resp []datatypes.Container_RemoteManagement_Graphs_SensorSpeed, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getServerFanSpeedGraphs", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieves the power state for the server.  The server's power status is retrieved from its remote management card.  This will return 'on' or 'off'.
func (r Hardware_SecurityModule) GetServerPowerState() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getServerPowerState", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the server room in which the hardware is located.
func (r Hardware_SecurityModule) GetServerRoom() (resp datatypes.Location, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getServerRoom", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the server's temperature and displays them using thermometer graphs.  Temperatures retrieved are CPU(s) and system temperatures.  Data used to construct graphs is retrieved from the server's remote management card.  All graphs returned will have a title associated with it.
func (r Hardware_SecurityModule) GetServerTemperatureGraphs() (resp []datatypes.Container_RemoteManagement_Graphs_SensorTemperature, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getServerTemperatureGraphs", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the piece of hardware's service provider.
func (r Hardware_SecurityModule) GetServiceProvider() (resp datatypes.Service_Provider, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getServiceProvider", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's installed software.
func (r Hardware_SecurityModule) GetSoftwareComponents() (resp []datatypes.Software_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getSoftwareComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the billing item for a spare pool server.
func (r Hardware_SecurityModule) GetSparePoolBillingItem() (resp datatypes.Billing_Item_Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getSparePoolBillingItem", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve SSH keys to be installed on the server during provisioning or an OS reload.
func (r Hardware_SecurityModule) GetSshKeys() (resp []datatypes.Security_Ssh_Key, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getSshKeys", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A server's remote management card used for statistics.
func (r Hardware_SecurityModule) GetStatisticsRemoteManagement() (resp datatypes.Hardware_Component_RemoteManagement, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getStatisticsRemoteManagement", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_SecurityModule) GetStorageNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getStorageNetworkComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_SecurityModule) GetTagReferences() (resp []datatypes.Tag_Reference, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getTagReferences", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_SecurityModule) GetTopLevelLocation() (resp datatypes.Location, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getTopLevelLocation", nil, r.Session, &r.Options, &resp)
	return
}

//
// This method will query transaction history for a piece of hardware.
func (r Hardware_SecurityModule) GetTransactionHistory() (resp []datatypes.Provisioning_Version1_Transaction_History, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getTransactionHistory", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a list of upgradeable items available to this piece of hardware. Currently, getUpgradeItemPrices retrieves upgrades available for a server's memory, hard drives, network port speed, bandwidth allocation and GPUs.
func (r Hardware_SecurityModule) GetUpgradeItemPrices() (resp []datatypes.Product_Item_Price, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getUpgradeItemPrices", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An account's associated upgrade request object, if any.
func (r Hardware_SecurityModule) GetUpgradeRequest() (resp datatypes.Product_Upgrade_Request, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getUpgradeRequest", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The network device connected to a piece of hardware.
func (r Hardware_SecurityModule) GetUplinkHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getUplinkHardware", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the network component that is one level higher than a piece of hardware on the network infrastructure.
func (r Hardware_SecurityModule) GetUplinkNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getUplinkNetworkComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A string containing custom user data for a hardware order.
func (r Hardware_SecurityModule) GetUserData() (resp []datatypes.Hardware_Attribute, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getUserData", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A list of users that have access to this computing instance.
func (r Hardware_SecurityModule) GetUsers() (resp []datatypes.User_Customer, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getUsers", nil, r.Session, &r.Options, &resp)
	return
}

// This method will return the list of block device template groups that are valid to the host. For instance, it will only retrieve FLEX images.
func (r Hardware_SecurityModule) GetValidBlockDeviceTemplateGroups(visibility *string) (resp []datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	params := []interface{}{
		visibility,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getValidBlockDeviceTemplateGroups", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the virtual chassis for a piece of hardware.
func (r Hardware_SecurityModule) GetVirtualChassis() (resp datatypes.Hardware_Group, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getVirtualChassis", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the virtual chassis siblings for a piece of hardware.
func (r Hardware_SecurityModule) GetVirtualChassisSiblings() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getVirtualChassisSiblings", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware server's virtual servers.
func (r Hardware_SecurityModule) GetVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getVirtualGuests", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's virtual host record.
func (r Hardware_SecurityModule) GetVirtualHost() (resp datatypes.Virtual_Host, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getVirtualHost", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's virtual software licenses.
func (r Hardware_SecurityModule) GetVirtualLicenses() (resp []datatypes.Software_VirtualLicense, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getVirtualLicenses", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the bandwidth allotment to which a piece of hardware belongs.
func (r Hardware_SecurityModule) GetVirtualRack() (resp datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getVirtualRack", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The name of the bandwidth allotment belonging to a piece of hardware.
func (r Hardware_SecurityModule) GetVirtualRackId() (resp int, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getVirtualRackId", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The name of the bandwidth allotment belonging to a piece of hardware.
func (r Hardware_SecurityModule) GetVirtualRackName() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getVirtualRackName", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's virtualization platform software.
func (r Hardware_SecurityModule) GetVirtualizationPlatform() (resp datatypes.Software_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getVirtualizationPlatform", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a list of Windows updates available for a server from the local SoftLayer Windows Server Update Services (WSUS) server. Windows servers provisioned by SoftLayer are configured to use the local WSUS server via the private network by default.
func (r Hardware_SecurityModule) GetWindowsUpdateAvailableUpdates() (resp []datatypes.Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getWindowsUpdateAvailableUpdates", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a list of Windows updates installed on a server as reported by the local SoftLayer Windows Server Update Services (WSUS) server. Windows servers provisioned by SoftLayer are configured to use the local WSUS server via the private network by default.
func (r Hardware_SecurityModule) GetWindowsUpdateInstalledUpdates() (resp []datatypes.Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getWindowsUpdateInstalledUpdates", nil, r.Session, &r.Options, &resp)
	return
}

// This method returns an update status record for this server.  That record will specify if the server is missing updates, or has updates that must be reinstalled or require a reboot to go into affect.
func (r Hardware_SecurityModule) GetWindowsUpdateStatus() (resp datatypes.Container_Utility_Microsoft_Windows_UpdateServices_Status, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "getWindowsUpdateStatus", nil, r.Session, &r.Options, &resp)
	return
}

// The '''importVirtualHost''' method attempts to import the host record for the virtualization platform running on a server.
func (r Hardware_SecurityModule) ImportVirtualHost() (resp datatypes.Virtual_Host, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "importVirtualHost", nil, r.Session, &r.Options, &resp)
	return
}

// Idera Bare Metal Server Restore is a backup agent designed specifically for making full system restores made with Idera Server Backup.
func (r Hardware_SecurityModule) InitiateIderaBareMetalRestore() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "initiateIderaBareMetalRestore", nil, r.Session, &r.Options, &resp)
	return
}

// R1Soft Bare Metal Server Restore is an R1Soft disk agent designed specifically for making full system restores made with R1Soft CDP Server backup.
func (r Hardware_SecurityModule) InitiateR1SoftBareMetalRestore() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "initiateR1SoftBareMetalRestore", nil, r.Session, &r.Options, &resp)
	return
}

// Issues a ping command and returns the success (true) or failure (false) of the ping command.
func (r Hardware_SecurityModule) IsBackendPingable() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "isBackendPingable", nil, r.Session, &r.Options, &resp)
	return
}

// Issues a ping command and returns the success (true) or failure (false) of the ping command.
func (r Hardware_SecurityModule) IsPingable() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "isPingable", nil, r.Session, &r.Options, &resp)
	return
}

// Determine if the server runs any version of the Microsoft Windows operating systems. Return ''true'' if it does and ''false if otherwise.
func (r Hardware_SecurityModule) IsWindowsServer() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "isWindowsServer", nil, r.Session, &r.Options, &resp)
	return
}

// Issues a ping command to the server and returns the ping response.
func (r Hardware_SecurityModule) Ping() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "ping", nil, r.Session, &r.Options, &resp)
	return
}

// Power off then power on the server via powerstrip.  The power cycle command is equivalent to unplugging the server from the powerstrip and then plugging the server back into the powerstrip.  This should only be used as a last resort.  If a reboot command has been issued successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard, powerOn, powerOff and powerCycle) will not be allowed. This is to avoid any type of server failures.
func (r Hardware_SecurityModule) PowerCycle() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "powerCycle", nil, r.Session, &r.Options, &resp)
	return
}

// This method will power off the server via the server's remote management card.
func (r Hardware_SecurityModule) PowerOff() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "powerOff", nil, r.Session, &r.Options, &resp)
	return
}

// Power on server via its remote management card.  If a reboot command has been issued successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard, powerOn, powerOff and powerCycle) will not be allowed.  This is to avoid any type of server failures.
func (r Hardware_SecurityModule) PowerOn() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "powerOn", nil, r.Session, &r.Options, &resp)
	return
}

// Attempts to reboot the server by issuing a reset (soft reboot) command to the server's remote management card. If the reset (soft reboot) attempt is unsuccessful, a power cycle command will be issued via the powerstrip. The power cycle command is equivalent to unplugging the server from the powerstrip and then plugging the server back into the powerstrip.  If a reboot command has been issued successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard, powerOn, powerOff and powerCycle) will not be allowed.  This is to avoid any type of server failures.
func (r Hardware_SecurityModule) RebootDefault() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "rebootDefault", nil, r.Session, &r.Options, &resp)
	return
}

// Reboot the server by issuing a cycle command to the server's remote management card.  This is equivalent to pressing the 'Reset' button on the server.  This command is issued immediately and will not wait for processes to shutdown. After this command is issued, the server may take a few moments to boot up as server may run system disks checks. If a reboot command has been issued successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard, powerOn, powerOff and powerCycle) will not be allowed.  This is to avoid any type of server failures.
func (r Hardware_SecurityModule) RebootHard() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "rebootHard", nil, r.Session, &r.Options, &resp)
	return
}

// Reboot the server by issuing a reset command to the server's remote management card.  This is a graceful reboot. The servers will allow all process to shutdown gracefully before rebooting.  If a reboot command has been issued successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard, powerOn, powerOff and powerCycle) will not be allowed.  This is to avoid any type of server failures.
func (r Hardware_SecurityModule) RebootSoft() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "rebootSoft", nil, r.Session, &r.Options, &resp)
	return
}

// Reloads current operating system configuration.
//
// This service has a confirmation protocol for proceeding with the reload. To proceed with the reload without confirmation, simply pass in 'FORCE' as the token parameter. To proceed with the reload with confirmation, simply call the service with no parameter. A token string will be returned by this service. The token will remain active for 10 minutes. Use this token as the parameter to confirm that a reload is to be performed for the server.
//
// As a precaution, we strongly  recommend backing up all data before reloading the operating system. The reload will format the primary disk and will reconfigure the server to the current specifications on record.
//
// The reload will take AT MINIMUM 66 minutes.
func (r Hardware_SecurityModule) ReloadCurrentOperatingSystemConfiguration(token *string) (resp string, err error) {
	params := []interface{}{
		token,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "reloadCurrentOperatingSystemConfiguration", params, r.Session, &r.Options, &resp)
	return
}

// Reloads current or customer specified operating system configuration.
//
// This service has a confirmation protocol for proceeding with the reload. To proceed with the reload without confirmation, simply pass in 'FORCE' as the token parameter. To proceed with the reload with confirmation, simply call the service with no parameter. A token string will be returned by this service. The token will remain active for 10 minutes. Use this token as the parameter to confirm that a reload is to be performed for the server.
//
// As a precaution, we strongly  recommend backing up all data before reloading the operating system. The reload will format the primary disk and will reconfigure the server to the current specifications on record.
//
// The reload will take AT MINIMUM 66 minutes.
func (r Hardware_SecurityModule) ReloadOperatingSystem(token *string, config *datatypes.Container_Hardware_Server_Configuration) (resp string, err error) {
	params := []interface{}{
		token,
		config,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "reloadOperatingSystem", params, r.Session, &r.Options, &resp)
	return
}

// This method is used to remove access to s SoftLayer_Network_Storage volumes that supports host- or network-level access control.
func (r Hardware_SecurityModule) RemoveAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "removeAccessToNetworkStorage", params, r.Session, &r.Options, &resp)
	return
}

// This method is used to allow access to multiple SoftLayer_Network_Storage volumes that support host- or network-level access control.
func (r Hardware_SecurityModule) RemoveAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "removeAccessToNetworkStorageList", params, r.Session, &r.Options, &resp)
	return
}

// You can launch a new Passmark hardware test by selecting from your server list. It will bring your server offline for approximately 20 minutes while the testing is in progress, and will publish a certificate with the results to your hardware details page.
//
// While the hard drives are tested for the initial deployment, the Passmark Certificate utility will not test the hard drives on your live server. This is to ensure that no data is overwritten. If you would like to test the server's hard drives, you can have the full Passmark suite installed to your server free of charge through a new Support ticket.
//
// While the test itself does not overwrite any data on the server, it is recommended that you make full off-server backups of all data prior to launching the test. The Passmark hardware test is designed to force any latent hardware issues to the surface, so hardware failure is possible.
//
// In the event of a hardware failure during this test our datacenter engineers will be notified of the problem automatically. They will then replace any failed components to bring your server back online, and will be contacting you to ensure that impact on your server is minimal.
func (r Hardware_SecurityModule) RunPassmarkCertificationBenchmark() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "runPassmarkCertificationBenchmark", nil, r.Session, &r.Options, &resp)
	return
}

// Changes the password that we have stored in our database for a servers' Operating System
func (r Hardware_SecurityModule) SetOperatingSystemPassword(newPassword *string) (resp bool, err error) {
	params := []interface{}{
		newPassword,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "setOperatingSystemPassword", params, r.Session, &r.Options, &resp)
	return
}

// Sets the private network interface speed to the new speed. Speed values can only be 0 (Disconnect), 10, 100, 1000, and 10000. The new speed must be equal to or less than the max speed of the interface.
//
// It will take less than a minute to update the switch port speed. The server uplink will not be operational again until the server interface speed is updated.
func (r Hardware_SecurityModule) SetPrivateNetworkInterfaceSpeed(newSpeed *int) (resp bool, err error) {
	params := []interface{}{
		newSpeed,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "setPrivateNetworkInterfaceSpeed", params, r.Session, &r.Options, &resp)
	return
}

// Sets the public network interface speed to the new speed. Speed values can only be 0 (Disconnect), 10, 100, 1000, and 10000. The new speed must be equal to or less than the max speed of the interface.
//
// It will take less than a minute to update the switch port speed. The server uplink will not be operational again until the server interface speed is updated.
func (r Hardware_SecurityModule) SetPublicNetworkInterfaceSpeed(newSpeed *int) (resp bool, err error) {
	params := []interface{}{
		newSpeed,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "setPublicNetworkInterfaceSpeed", params, r.Session, &r.Options, &resp)
	return
}

// no documentation yet
func (r Hardware_SecurityModule) SetTags(tags *string) (resp bool, err error) {
	params := []interface{}{
		tags,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "setTags", params, r.Session, &r.Options, &resp)
	return
}

// Sets the data that will be written to the configuration drive.
func (r Hardware_SecurityModule) SetUserMetadata(metadata []string) (resp []datatypes.Hardware_Attribute, err error) {
	params := []interface{}{
		metadata,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "setUserMetadata", params, r.Session, &r.Options, &resp)
	return
}

// Shuts down the public network port
func (r Hardware_SecurityModule) ShutdownPrivatePort() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "shutdownPrivatePort", nil, r.Session, &r.Options, &resp)
	return
}

// Shuts down the public network port
func (r Hardware_SecurityModule) ShutdownPublicPort() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "shutdownPublicPort", nil, r.Session, &r.Options, &resp)
	return
}

// The ability to place bare metal servers in a state where they are powered down, and ports closed yet still allocated to the customer as a part of the Spare Pool program.
func (r Hardware_SecurityModule) SparePool(action *string, newOrder *bool) (resp bool, err error) {
	params := []interface{}{
		action,
		newOrder,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "sparePool", params, r.Session, &r.Options, &resp)
	return
}

// Validates a collection of partitions for an operating system
func (r Hardware_SecurityModule) ValidatePartitionsForOperatingSystem(operatingSystem *datatypes.Software_Description, partitions []datatypes.Hardware_Component_Partition) (resp bool, err error) {
	params := []interface{}{
		operatingSystem,
		partitions,
	}
	err = invokeMethod("SoftLayer_Hardware_SecurityModule", "validatePartitionsForOperatingSystem", params, r.Session, &r.Options, &resp)
	return
}

// The SoftLayer_Hardware_Server data type contains general information relating to a single SoftLayer server.
type Hardware_Server struct {
	Session *session.Session
	Options sl.Options
}

func GetHardwareServerService(sess *session.Session) Hardware_Server {
	return Hardware_Server{Session: sess}
}

func (r Hardware_Server) Id(id int) Hardware_Server {
	r.Options.Id = &id
	return r
}

func (r Hardware_Server) Mask(mask string) Hardware_Server {
	if !strings.HasPrefix(mask, "mask[") && strings.Contains(mask, "[") {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Hardware_Server) Filter(filter string) Hardware_Server {
	r.Options.Filter = filter
	return r
}

func (r Hardware_Server) Limit(limit int) Hardware_Server {
	r.Options.Limit = &limit
	return r
}

func (r Hardware_Server) Offset(offset int) Hardware_Server {
	r.Options.Offset = &offset
	return r
}

// Activates the private network port
func (r Hardware_Server) ActivatePrivatePort() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "activatePrivatePort", nil, r.Session, &r.Options, &resp)
	return
}

// Activates the public network port
func (r Hardware_Server) ActivatePublicPort() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "activatePublicPort", nil, r.Session, &r.Options, &resp)
	return
}

// This method is used to allow access to a SoftLayer_Network_Storage volume that supports host- or network-level access control.
func (r Hardware_Server) AllowAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "allowAccessToNetworkStorage", params, r.Session, &r.Options, &resp)
	return
}

// This method is used to allow access to multiple SoftLayer_Network_Storage volumes that support host- or network-level access control.
func (r Hardware_Server) AllowAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "allowAccessToNetworkStorageList", params, r.Session, &r.Options, &resp)
	return
}

// The Rescue Kernel is designed to provide you with the ability to bring a server online in order to troubleshoot system problems that would normally only be resolved by an OS Reload. The correct Rescue Kernel will be selected based upon the currently installed operating system. When the rescue kernel process is initiated, the server will shutdown and reboot on to the public network with the same IP's assigned to the server to allow for remote connections. It will bring your server offline for approximately 10 minutes while the rescue is in progress. The root/administrator password will be the same as what is listed in the portal for the server.
func (r Hardware_Server) BootToRescueLayer(noOsBootEnvironment *string) (resp bool, err error) {
	params := []interface{}{
		noOsBootEnvironment,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "bootToRescueLayer", params, r.Session, &r.Options, &resp)
	return
}

// Captures a Flex Image of the hard disk on the physical machine, based on the capture template parameter. Returns the image template group containing the disk image.
func (r Hardware_Server) CaptureImage(captureTemplate *datatypes.Container_Disk_Image_Capture_Template) (resp datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	params := []interface{}{
		captureTemplate,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "captureImage", params, r.Session, &r.Options, &resp)
	return
}

// Returns monitoring alarm detailed history
func (r Hardware_Server) CloseAlarm(alarmId *string) (resp bool, err error) {
	params := []interface{}{
		alarmId,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "closeAlarm", params, r.Session, &r.Options, &resp)
	return
}

// You can launch firmware updates by selecting from your server list. It will bring your server offline for approximately 20 minutes while the updates are in progress.
//
// In the event of a hardware failure during this test our datacenter engineers will be notified of the problem automatically. They will then replace any failed components to bring your server back online, and will be contacting you to ensure that impact on your server is minimal.
func (r Hardware_Server) CreateFirmwareUpdateTransaction(ipmi *int, raidController *int, bios *int, harddrive *int) (resp bool, err error) {
	params := []interface{}{
		ipmi,
		raidController,
		bios,
		harddrive,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "createFirmwareUpdateTransaction", params, r.Session, &r.Options, &resp)
	return
}

//
// <style type="text/css">.create_object > li > div { padding-top: .5em; padding-bottom: .5em}</style>
// createObject() enables the creation of servers on an account. This
// method is a simplified alternative to interacting with the ordering system directly.
//
//
// In order to create a server, a template object must be sent in with a few required
// values.
//
//
// When this method returns an order will have been placed for a server of the specified configuration.
//
//
// To determine when the server is available you can poll the server via [[SoftLayer_Hardware/getObject|getObject]],
// checking the <code>provisionDate</code> property.
// When <code>provisionDate</code> is not null, the server will be ready. Be sure to use the <code>globalIdentifier</code>
// as your initialization parameter.
//
//
// <b>Warning:</b> Servers created via this method will incur charges on your account. For testing input parameters see [[SoftLayer_Hardware/generateOrderTemplate|generateOrderTemplate]].
//
//
// <b>Input</b> - [[SoftLayer_Hardware (type)|SoftLayer_Hardware]]
// <ul class="create_object">
//     <li><code>hostname</code>
//         <div>Hostname for the server.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - string</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>domain</code>
//         <div>Domain for the server.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - string</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>processorCoreAmount</code>
//         <div>The number of logical CPU cores to allocate.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - int</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>memoryCapacity</code>
//         <div>The amount of memory to allocate in gigabytes.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - int</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>hourlyBillingFlag</code>
//         <div>Specifies the billing type for the server.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - boolean</li>
//             <li>When true the server will be billed on hourly usage, otherwise it will be billed on a monthly basis.</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>operatingSystemReferenceCode</code>
//         <div>An identifier for the operating system to provision the server with.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - string</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>datacenter.name</code>
//         <div>Specifies which datacenter the server is to be provisioned in.</div><ul>
//             <li><b>Required</b></li>
//             <li><b>Type</b> - string</li>
//             <li>The <code>datacenter</code> property is a [[SoftLayer_Location (type)|location]] structure with the <code>name</code> field set.</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//         <http title="Example">{
//     "datacenter": {
//         "name": "dal05"
//     }
// }</http>
//         <br />
//     </li>
//     <li><code>networkComponents.maxSpeed</code>
//         <div>Specifies the connection speed for the server's network components.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - int</li>
//             <li><b>Default</b> - The highest available zero cost port speed will be used.</li>
//             <li><b>Description</b> - The <code>networkComponents</code> property is an array with a single [[SoftLayer_Network_Component (type)|network component]] structure. The <code>maxSpeed</code> property must be set to specify the network uplink speed, in megabits per second, of the server.</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//             <http title="Example">{
//     "networkComponents": [
//         {
//             "maxSpeed": 1000
//         }
//     ]
// }</http>
//         <br />
//     </li>
//     <li><code>networkComponents.redundancyEnabledFlag</code>
//         <div>Specifies whether or not the server's network components should be in redundancy groups.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - bool</li>
//             <li><b>Default</b> - <code>false</code></li>
//             <li><b>Description</b> - The <code>networkComponents</code> property is an array with a single [[SoftLayer_Network_Component (type)|network component]] structure. When the <code>redundancyEnabledFlag</code> property is true the server's network components will be in redundancy groups.</li>
//         </ul>
//             <http title="Example">{
//     "networkComponents": [
//         {
//             "redundancyEnabledFlag": false
//         }
//     ]
// }</http>
//         <br />
//     </li>
//     <li><code>privateNetworkOnlyFlag</code>
//         <div>Specifies whether or not the server only has access to the private network</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - boolean</li>
//             <li><b>Default</b> - <code>false</code></li>
//             <li>When true this flag specifies that a server is to only have access to the private network.</li>
//         </ul>
//         <br />
//     </li>
//     <li><code>primaryNetworkComponent.networkVlan.id</code>
//         <div>Specifies the network vlan which is to be used for the frontend interface of the server.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - int</li>
//             <li><b>Description</b> - The <code>primaryNetworkComponent</code> property is a [[SoftLayer_Network_Component (type)|network component]] structure with the <code>networkVlan</code> property populated with a [[SoftLayer_Network_Vlan (type)|vlan]] structure. The <code>id</code> property must be set to specify the frontend network vlan of the server.</li>
//         </ul>
//         <http title="Example">{
//     "primaryNetworkComponent": {
//         "networkVlan": {
//             "id": 1
//         }
//     }
// }</http>
//         <br />
//     </li>
//     <li><code>primaryBackendNetworkComponent.networkVlan.id</code>
//         <div>Specifies the network vlan which is to be used for the backend interface of the server.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - int</li>
//             <li><b>Description</b> - The <code>primaryBackendNetworkComponent</code> property is a [[SoftLayer_Network_Component (type)|network component]] structure with the <code>networkVlan</code> property populated with a [[SoftLayer_Network_Vlan (type)|vlan]] structure. The <code>id</code> property must be set to specify the backend network vlan of the server.</li>
//         </ul>
//         <http title="Example">{
//     "primaryBackendNetworkComponent": {
//         "networkVlan": {
//             "id": 2
//         }
//     }
// }</http>
//         <br />
//     </li>
//     <li><code>fixedConfigurationPreset.keyName</code>
//         <div></div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - string</li>
//             <li><b>Description</b> - The <code>fixedConfigurationPreset</code> property is a [[SoftLayer_Product_Package_Preset (type)|fixed configuration preset]] structure. The <code>keyName</code> property must be set to specify preset to use.</li>
//             <li>If a fixed configuration preset is used <code>processorCoreAmount</code>, <code>memoryCapacity</code> and <code>hardDrives</code> properties must not be set.</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//         <http title="Example">{
//     "fixedConfigurationPreset": {
//         "keyName": "SOME_KEY_NAME"
//     }
// }</http>
//         <br />
//     </li>
//     <li><code>userData.value</code>
//         <div>Arbitrary data to be made available to the server.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - string</li>
//             <li><b>Description</b> - The <code>userData</code> property is an array with a single [[SoftLayer_Hardware_Attribute (type)|attribute]] structure with the <code>value</code> property set to an arbitrary value.</li>
//             <li>This value can be retrieved via the [[SoftLayer_Resource_Metadata/getUserMetadata|getUserMetadata]] method from a request originating from the server. This is primarily useful for providing data to software that may be on the server and configured to execute upon first boot.</li>
//         </ul>
//         <http title="Example">{
//     "userData": [
//         {
//             "value": "someValue"
//         }
//     ]
// }</http>
//         <br />
//     </li>
//     <li><code>hardDrives</code>
//         <div>Hard drive settings for the server</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - SoftLayer_Hardware_Component</li>
//             <li><b>Default</b> - The largest available capacity for a zero cost primary disk will be used.</li>
//             <li><b>Description</b> - The <code>hardDrives</code> property is an array of [[SoftLayer_Hardware_Component (type)|hardware component]] structures.</i>
//             <li>Each hard drive must specify the <code>capacity</code> property.</li>
//             <li>See [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] for available options.</li>
//         </ul>
//         <http title="Example">{
//     "hardDrives": [
//         {
//             "capacity": 500
//         }
//     ]
// }</http>
//         <br />
//     </li>
//     <li id="hardware-create-object-ssh-keys"><code>sshKeys</code>
//         <div>SSH keys to install on the server upon provisioning.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - array of [[SoftLayer_Security_Ssh_Key (type)|SoftLayer_Security_Ssh_Key]]</li>
//             <li><b>Description</b> - The <code>sshKeys</code> property is an array of [[SoftLayer_Security_Ssh_Key (type)|SSH Key]] structures with the <code>id</code> property set to the value of an existing SSH key.</li>
//             <li>To create a new SSH key, call [[SoftLayer_Security_Ssh_Key/createObject|createObject]] on the [[SoftLayer_Security_Ssh_Key]] service.</li>
//             <li>To obtain a list of existing SSH keys, call [[SoftLayer_Account/getSshKeys|getSshKeys]] on the [[SoftLayer_Account]] service.
//         </ul>
//         <http title="Example">{
//     "sshKeys": [
//         {
//             "id": 123
//         }
//     ]
// }</http>
//         <br />
//     </li>
//     <li><code>postInstallScriptUri</code>
//         <div>Specifies the uri location of the script to be downloaded and run after installation is complete.</div><ul>
//             <li><b>Optional</b></li>
//             <li><b>Type</b> - string</li>
//         </ul>
//         <br />
//     </li>
// </ul>
//
//
// <h1>REST Example</h1>
// <http title="Request">curl -X POST -d '{
//  "parameters":[
//      {
//          "hostname": "host1",
//          "domain": "example.com",
//          "processorCoreAmount": 2,
//          "memoryCapacity": 2,
//          "hourlyBillingFlag": true,
//          "operatingSystemReferenceCode": "UBUNTU_LATEST"
//      }
//  ]
// }' https://api.softlayer.com/rest/v3/SoftLayer_Hardware.json
// </http>
// <http title="Response">HTTP/1.1 201 Created
// Location: https://api.softlayer.com/rest/v3/SoftLayer_Hardware/f5a3fcff-db1d-4b7c-9fa0-0349e41c29c5/getObject
//
//
// {
//     "accountId": 232298,
//     "bareMetalInstanceFlag": null,
//     "domain": "example.com",
//     "hardwareStatusId": null,
//     "hostname": "host1",
//     "id": null,
//     "serviceProviderId": null,
//     "serviceProviderResourceId": null,
//     "globalIdentifier": "f5a3fcff-db1d-4b7c-9fa0-0349e41c29c5",
//     "hourlyBillingFlag": true,
//     "memoryCapacity": 2,
//     "operatingSystemReferenceCode": "UBUNTU_LATEST",
//     "processorCoreAmount": 2
// }
// </http>
func (r Hardware_Server) CreateObject(templateObject *datatypes.Hardware_Server) (resp datatypes.Hardware_Server, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "createObject", params, r.Session, &r.Options, &resp)
	return
}

// no documentation yet
func (r Hardware_Server) CreatePostSoftwareInstallTransaction(installCodes []string, returnBoolean *bool) (resp bool, err error) {
	params := []interface{}{
		installCodes,
		returnBoolean,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "createPostSoftwareInstallTransaction", params, r.Session, &r.Options, &resp)
	return
}

//
// This method will cancel a server effective immediately. For servers billed hourly, the charges will stop immediately after the method returns.
func (r Hardware_Server) DeleteObject() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "deleteObject", nil, r.Session, &r.Options, &resp)
	return
}

// Delete software component passwords.
func (r Hardware_Server) DeleteSoftwareComponentPasswords(softwareComponentPasswords []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		softwareComponentPasswords,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "deleteSoftwareComponentPasswords", params, r.Session, &r.Options, &resp)
	return
}

// Edit a server's properties
func (r Hardware_Server) EditObject(templateObject *datatypes.Hardware_Server) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "editObject", params, r.Session, &r.Options, &resp)
	return
}

// Edit the properties of a software component password such as the username, password, and notes.
func (r Hardware_Server) EditSoftwareComponentPasswords(softwareComponentPasswords []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		softwareComponentPasswords,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "editSoftwareComponentPasswords", params, r.Session, &r.Options, &resp)
	return
}

// Download and run remote script from uri on the hardware.
func (r Hardware_Server) ExecuteRemoteScript(uri *string) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		uri,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "executeRemoteScript", params, r.Session, &r.Options, &resp)
	return
}

// The '''findByIpAddress''' method finds hardware using its primary public or private IP address. IP addresses that have a secondary subnet tied to the hardware will not return the hardware - alternate means of locating the hardware must be used (see '''Associated Methods'''). If no hardware is found, no errors are generated and no data is returned.
func (r Hardware_Server) FindByIpAddress(ipAddress *string) (resp datatypes.Hardware, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "findByIpAddress", params, r.Session, &r.Options, &resp)
	return
}

//
// Obtain an [[SoftLayer_Container_Product_Order_Hardware_Server (type)|order container]] that can be sent to [[SoftLayer_Product_Order/verifyOrder|verifyOrder]] or [[SoftLayer_Product_Order/placeOrder|placeOrder]].
//
//
// This is primarily useful when there is a necessity to confirm the price which will be charged for an order.
//
//
// See [[SoftLayer_Hardware/createObject|createObject]] for specifics on the requirements of the template object parameter.
func (r Hardware_Server) GenerateOrderTemplate(templateObject *datatypes.Hardware) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "generateOrderTemplate", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The account associated with a piece of hardware.
func (r Hardware_Server) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getAccount", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's active physical components.
func (r Hardware_Server) GetActiveComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getActiveComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The billing item for a server's attached network firewall.
func (r Hardware_Server) GetActiveNetworkFirewallBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getActiveNetworkFirewallBillingItem", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's active network monitoring incidents.
func (r Hardware_Server) GetActiveNetworkMonitorIncident() (resp []datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getActiveNetworkMonitorIncident", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Server) GetActiveTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getActiveTickets", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Transaction currently running for server.
func (r Hardware_Server) GetActiveTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getActiveTransaction", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Any active transaction(s) that are currently running for the server (example: os reload).
func (r Hardware_Server) GetActiveTransactions() (resp []datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getActiveTransactions", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getAlarmHistory''' method retrieves a detailed history for the monitoring alarm. When calling this method, a start and end date for the history to be retrieved must be entered.
func (r Hardware_Server) GetAlarmHistory(startDate *datatypes.Time, endDate *datatypes.Time, alarmId *string) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
		alarmId,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getAlarmHistory", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Server) GetAllPowerComponents() (resp []datatypes.Hardware_Power_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getAllPowerComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Storage_Allowed_Host information to connect this server to Network Storage volumes that require access control lists.
func (r Hardware_Server) GetAllowedHost() (resp datatypes.Network_Storage_Allowed_Host, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getAllowedHost", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Storage objects that this SoftLayer_Hardware has access to.
func (r Hardware_Server) GetAllowedNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getAllowedNetworkStorage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Storage objects whose Replica that this SoftLayer_Hardware has access to.
func (r Hardware_Server) GetAllowedNetworkStorageReplicas() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getAllowedNetworkStorageReplicas", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding an antivirus/spyware software component object.
func (r Hardware_Server) GetAntivirusSpywareSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getAntivirusSpywareSoftwareComponent", nil, r.Session, &r.Options, &resp)
	return
}

// This method is retrieve a list of SoftLayer_Network_Storage volumes that are authorized access to this SoftLayer_Hardware.
func (r Hardware_Server) GetAttachedNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getAttachedNetworkStorages", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's specific attributes.
func (r Hardware_Server) GetAttributes() (resp []datatypes.Hardware_Attribute, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getAttributes", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An object that stores the maximum level for the monitoring query types and response types.
func (r Hardware_Server) GetAvailableMonitoring() (resp []datatypes.Network_Monitor_Version1_Query_Host_Stratum, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getAvailableMonitoring", nil, r.Session, &r.Options, &resp)
	return
}

// This method retrieves a list of SoftLayer_Network_Storage volumes that can be authorized to this SoftLayer_Hardware.
func (r Hardware_Server) GetAvailableNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getAvailableNetworkStorages", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The average daily total bandwidth usage for the current billing cycle.
func (r Hardware_Server) GetAverageDailyBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getAverageDailyBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The average daily private bandwidth usage for the current billing cycle.
func (r Hardware_Server) GetAverageDailyPrivateBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getAverageDailyPrivateBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The average daily public bandwidth usage for the current billing cycle.
func (r Hardware_Server) GetAverageDailyPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getAverageDailyPublicBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Use this method to return an array of private bandwidth utilization records between a given date range.
//
// This method represents the NEW version of getFrontendBandwidthUse
func (r Hardware_Server) GetBackendBandwidthUsage(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getBackendBandwidthUsage", params, r.Session, &r.Options, &resp)
	return
}

// Use this method to return an array of private bandwidth utilization records between a given date range.
func (r Hardware_Server) GetBackendBandwidthUse(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Network_Bandwidth_Version1_Usage_Detail, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getBackendBandwidthUse", params, r.Session, &r.Options, &resp)
	return
}

// The '''getBackendIncomingBandwidth''' method retrieves the amount of incoming private network traffic used between the given start date and end date parameters. When entering start and end dates, only the month, day and year are used to calculate bandwidth totals - the time (HH:MM:SS) is ignored and defaults to midnight. The amount of bandwidth retrieved is measured in gigabytes.
func (r Hardware_Server) GetBackendIncomingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getBackendIncomingBandwidth", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's back-end or private network components.
func (r Hardware_Server) GetBackendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getBackendNetworkComponents", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getBackendOutgoingBandwidth''' method retrieves the amount of outgoing private network traffic used between the given start date and end date parameters. When entering start and end dates, only the month, day and year are used to calculate bandwidth totals - the time (HH:MM:SS) is ignored and defaults to midnight. The amount of bandwidth retrieved is measured in gigabytes.
func (r Hardware_Server) GetBackendOutgoingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getBackendOutgoingBandwidth", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's backend or private router.
func (r Hardware_Server) GetBackendRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getBackendRouters", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's allotted bandwidth (measured in GB).
func (r Hardware_Server) GetBandwidthAllocation() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getBandwidthAllocation", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's allotted detail record. Allotment details link bandwidth allocation with allotments.
func (r Hardware_Server) GetBandwidthAllotmentDetail() (resp datatypes.Network_Bandwidth_Version1_Allotment_Detail, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getBandwidthAllotmentDetail", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a collection of bandwidth data from an individual public or private network tracking object. Data is ideal if you with to employ your own traffic storage and graphing systems.
func (r Hardware_Server) GetBandwidthForDateRange(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getBandwidthForDateRange", params, r.Session, &r.Options, &resp)
	return
}

// Use this method when needing a bandwidth image for a single server.  It will gather the correct input parameters for the generic graphing utility automatically based on the snapshot specified.  Use the $draw flag to suppress the generation of the actual binary PNG image.
func (r Hardware_Server) GetBandwidthImage(networkType *string, snapshotRange *string, draw *bool, dateSpecified *datatypes.Time, dateSpecifiedEnd *datatypes.Time) (resp datatypes.Container_Bandwidth_GraphOutputs, err error) {
	params := []interface{}{
		networkType,
		snapshotRange,
		draw,
		dateSpecified,
		dateSpecifiedEnd,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getBandwidthImage", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's benchmark certifications.
func (r Hardware_Server) GetBenchmarkCertifications() (resp []datatypes.Hardware_Benchmark_Certification, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getBenchmarkCertifications", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The raw bandwidth usage data for the current billing cycle. One object will be returned for each network this server is attached to.
func (r Hardware_Server) GetBillingCycleBandwidthUsage() (resp []datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getBillingCycleBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The raw private bandwidth usage data for the current billing cycle.
func (r Hardware_Server) GetBillingCyclePrivateBandwidthUsage() (resp datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getBillingCyclePrivateBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The raw public bandwidth usage data for the current billing cycle.
func (r Hardware_Server) GetBillingCyclePublicBandwidthUsage() (resp datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getBillingCyclePublicBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the billing item for a server.
func (r Hardware_Server) GetBillingItem() (resp datatypes.Billing_Item_Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getBillingItem", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A flag indicating that a billing item exists.
func (r Hardware_Server) GetBillingItemFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getBillingItemFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Determines whether the hardware is ineligible for cancellation because it is disconnected.
func (r Hardware_Server) GetBlockCancelBecauseDisconnectedFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getBlockCancelBecauseDisconnectedFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Status indicating whether or not a piece of hardware has business continuance insurance.
func (r Hardware_Server) GetBusinessContinuanceInsuranceFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getBusinessContinuanceInsuranceFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's components.
func (r Hardware_Server) GetComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Server) GetContainsSolidStateDrivesFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getContainsSolidStateDrivesFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A continuous data protection/server backup software component object.
func (r Hardware_Server) GetContinuousDataProtectionSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getContinuousDataProtectionSoftwareComponent", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A server's control panel.
func (r Hardware_Server) GetControlPanel() (resp datatypes.Software_Component_ControlPanel, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getControlPanel", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total cost of a server, measured in US Dollars ($USD).
func (r Hardware_Server) GetCost() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getCost", nil, r.Session, &r.Options, &resp)
	return
}

//
// There are many options that may be provided while ordering a server, this method can be used to determine what these options are.
//
//
// Detailed information on the return value can be found on the data type page for [[SoftLayer_Container_Hardware_Configuration (type)]].
func (r Hardware_Server) GetCreateObjectOptions() (resp datatypes.Container_Hardware_Configuration, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getCreateObjectOptions", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An object that provides commonly used bandwidth summary components for the current billing cycle.
func (r Hardware_Server) GetCurrentBandwidthSummary() (resp datatypes.Metric_Tracking_Object_Bandwidth_Summary, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getCurrentBandwidthSummary", nil, r.Session, &r.Options, &resp)
	return
}

// Attempt to retrieve the file associated with the current benchmark certification result, if such a file exists.  If there is no file for this benchmark certification result, calling this method throws an exception.
func (r Hardware_Server) GetCurrentBenchmarkCertificationResultFile() (resp []byte, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getCurrentBenchmarkCertificationResultFile", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The current billable public outbound bandwidth for this hardware for the current billing cycle.
func (r Hardware_Server) GetCurrentBillableBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getCurrentBillableBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Get the billing detail for this instance for the current billing period. This does not include bandwidth usage.
func (r Hardware_Server) GetCurrentBillingDetail() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getCurrentBillingDetail", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getCurrentBillingTotal''' method retrieves the total bill amount in US Dollars ($) for the current billing period. In addition to the total bill amount, the billing detail also includes all bandwidth used up to the point the method is called on the piece of hardware.
func (r Hardware_Server) GetCurrentBillingTotal() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getCurrentBillingTotal", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve bandwidth graph by date.
func (r Hardware_Server) GetCustomBandwidthDataByDate(graphData *datatypes.Container_Graph) (resp datatypes.Container_Graph, err error) {
	params := []interface{}{
		graphData,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getCustomBandwidthDataByDate", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Indicates if a server has a Customer Installed OS
func (r Hardware_Server) GetCustomerInstalledOperatingSystemFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getCustomerInstalledOperatingSystemFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Indicates if a server is a customer owned device.
func (r Hardware_Server) GetCustomerOwnedFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getCustomerOwnedFlag", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getDailyAverage''' method calculates the average daily network traffic used by the selected server. Using the required parameter ''dateTime'' to enter a start and end date, the user retrieves this average, measure in gigabytes (GB) for the specified date range. When entering parameters, only the month, day and year are required - time entries are omitted as this method defaults the time to midnight in order to account for the entire day.
func (r Hardware_Server) GetDailyAverage(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getDailyAverage", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the datacenter in which a piece of hardware resides.
func (r Hardware_Server) GetDatacenter() (resp datatypes.Location, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getDatacenter", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The name of the datacenter in which a piece of hardware resides.
func (r Hardware_Server) GetDatacenterName() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getDatacenterName", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All hardware that has uplink network connections to a piece of hardware.
func (r Hardware_Server) GetDownlinkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getDownlinkHardware", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All hardware that has uplink network connections to a piece of hardware.
func (r Hardware_Server) GetDownlinkNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getDownlinkNetworkHardware", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all servers attached to a piece of network hardware.
func (r Hardware_Server) GetDownlinkServers() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getDownlinkServers", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all virtual guests attached to a piece of network hardware.
func (r Hardware_Server) GetDownlinkVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getDownlinkVirtualGuests", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All hardware downstream from a network device.
func (r Hardware_Server) GetDownstreamHardwareBindings() (resp []datatypes.Network_Component_Uplink_Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getDownstreamHardwareBindings", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All network hardware downstream from the selected piece of hardware.
func (r Hardware_Server) GetDownstreamNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getDownstreamNetworkHardware", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All network hardware with monitoring warnings or errors that are downstream from the selected piece of hardware.
func (r Hardware_Server) GetDownstreamNetworkHardwareWithIncidents() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getDownstreamNetworkHardwareWithIncidents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all servers attached downstream to a piece of network hardware.
func (r Hardware_Server) GetDownstreamServers() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getDownstreamServers", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all virtual guests attached to a piece of network hardware.
func (r Hardware_Server) GetDownstreamVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getDownstreamVirtualGuests", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The drive controllers contained within a piece of hardware.
func (r Hardware_Server) GetDriveControllers() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getDriveControllers", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's associated EVault network storage service account.
func (r Hardware_Server) GetEvaultNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getEvaultNetworkStorage", nil, r.Session, &r.Options, &resp)
	return
}

// Get the subnets associated with this server that are protectable by a network component firewall.
func (r Hardware_Server) GetFirewallProtectableSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getFirewallProtectableSubnets", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's firewall services.
func (r Hardware_Server) GetFirewallServiceComponent() (resp datatypes.Network_Component_Firewall, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getFirewallServiceComponent", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Defines the fixed components in a fixed configuration bare metal server.
func (r Hardware_Server) GetFixedConfigurationPreset() (resp datatypes.Product_Package_Preset, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getFixedConfigurationPreset", nil, r.Session, &r.Options, &resp)
	return
}

// Use this method to return an array of public bandwidth utilization records between a given date range.
//
// This method represents the NEW version of getFrontendBandwidthUse
func (r Hardware_Server) GetFrontendBandwidthUsage(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getFrontendBandwidthUsage", params, r.Session, &r.Options, &resp)
	return
}

// Use this method to return an array of public bandwidth utilization records between a given date range.
func (r Hardware_Server) GetFrontendBandwidthUse(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Network_Bandwidth_Version1_Usage_Detail, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getFrontendBandwidthUse", params, r.Session, &r.Options, &resp)
	return
}

// The '''getFrontendIncomingBandwidth''' method retrieves the amount of incoming public network traffic used by a server between the given start and end date parameters. When entering the ''dateTime'' parameter, only the month, day and year of the start and end dates are required - the time (hour, minute and second) are set to midnight by default and cannot be changed. The amount of bandwidth retrieved is measured in gigabytes (GB).
func (r Hardware_Server) GetFrontendIncomingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getFrontendIncomingBandwidth", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's front-end or public network components.
func (r Hardware_Server) GetFrontendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getFrontendNetworkComponents", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getFrontendOutgoingBandwidth''' method retrieves the amount of outgoing public network traffic used by a server between the given start and end date parameters. The ''dateTime'' parameter requires only the day, month and year to be entered - the time (hour, minute and second) are set to midnight be default in order to gather the data for the entire start and end date indicated in the parameter. The amount of bandwidth retrieved is measured in gigabytes (GB).
func (r Hardware_Server) GetFrontendOutgoingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getFrontendOutgoingBandwidth", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's frontend or public router.
func (r Hardware_Server) GetFrontendRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getFrontendRouters", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's universally unique identifier.
func (r Hardware_Server) GetGlobalIdentifier() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getGlobalIdentifier", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hard drives contained within a piece of hardware.
func (r Hardware_Server) GetHardDrives() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getHardDrives", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a server by searching for the primary IP address.
func (r Hardware_Server) GetHardwareByIpAddress(ipAddress *string) (resp datatypes.Hardware_Server, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getHardwareByIpAddress", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The chassis that a piece of hardware is housed in.
func (r Hardware_Server) GetHardwareChassis() (resp datatypes.Hardware_Chassis, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getHardwareChassis", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's function.
func (r Hardware_Server) GetHardwareFunction() (resp datatypes.Hardware_Function, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getHardwareFunction", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's function.
func (r Hardware_Server) GetHardwareFunctionDescription() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getHardwareFunctionDescription", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's status.
func (r Hardware_Server) GetHardwareStatus() (resp datatypes.Hardware_Status, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getHardwareStatus", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Determine in hardware object has TPM enabled.
func (r Hardware_Server) GetHasTrustedPlatformModuleBillingItemFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getHasTrustedPlatformModuleBillingItemFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a host IPS software component object.
func (r Hardware_Server) GetHostIpsSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getHostIpsSoftwareComponent", nil, r.Session, &r.Options, &resp)
	return
}

// The '''getHourlyBandwidth''' method retrieves all bandwidth updates hourly for the specified hardware. Because the potential number of data points can become excessive, the method limits the user to obtain data in 24-hour intervals. The required ''dateTime'' parameter is used as the starting point for the query and will be calculated for the 24-hour period starting with the specified date and time. For example, entering a parameter of
//
// '02/01/2008 0:00'
//
// results in a return of all bandwidth data for the entire day of February 1, 2008, as 0:00 specifies a midnight start date. Please note that the time entered should be completed using a 24-hour clock (military time, astronomical time).
//
// For data spanning more than a single 24-hour period, refer to the getBandwidthData function on the metricTrackingObject for the piece of hardware.
func (r Hardware_Server) GetHourlyBandwidth(mode *string, day *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		mode,
		day,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getHourlyBandwidth", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A server's hourly billing status.
func (r Hardware_Server) GetHourlyBillingFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getHourlyBillingFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The sum of all the inbound network traffic data for the last 30 days.
func (r Hardware_Server) GetInboundBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getInboundBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total private inbound bandwidth for this hardware for the current billing cycle.
func (r Hardware_Server) GetInboundPrivateBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getInboundPrivateBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total public inbound bandwidth for this hardware for the current billing cycle.
func (r Hardware_Server) GetInboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getInboundPublicBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Return a collection of SoftLayer_Item_Price objects from a collection of SoftLayer_Software_Description
func (r Hardware_Server) GetItemPricesFromSoftwareDescriptions(softwareDescriptions []datatypes.Software_Description, includeTranslationsFlag *bool, returnAllPricesFlag *bool) (resp []datatypes.Product_Item, err error) {
	params := []interface{}{
		softwareDescriptions,
		includeTranslationsFlag,
		returnAllPricesFlag,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getItemPricesFromSoftwareDescriptions", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The last transaction that a server's operating system was loaded.
func (r Hardware_Server) GetLastOperatingSystemReload() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getLastOperatingSystemReload", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the last transaction a server performed.
func (r Hardware_Server) GetLastTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getLastTransaction", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's latest network monitoring incident.
func (r Hardware_Server) GetLatestNetworkMonitorIncident() (resp datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getLatestNetworkMonitorIncident", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Where a piece of hardware is located within SoftLayer's location hierarchy.
func (r Hardware_Server) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getLocation", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Server) GetLocationPathString() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getLocationPathString", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a lockbox account associated with a server.
func (r Hardware_Server) GetLockboxNetworkStorage() (resp datatypes.Network_Storage, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getLockboxNetworkStorage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A flag indicating that the hardware is a managed resource.
func (r Hardware_Server) GetManagedResourceFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getManagedResourceFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the remote management network component attached with this server.
func (r Hardware_Server) GetManagementNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getManagementNetworkComponent", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's memory.
func (r Hardware_Server) GetMemory() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getMemory", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The amount of memory a piece of hardware has, measured in gigabytes.
func (r Hardware_Server) GetMemoryCapacity() (resp uint, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getMemoryCapacity", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's metric tracking object.
func (r Hardware_Server) GetMetricTrackingObject() (resp datatypes.Metric_Tracking_Object_HardwareServer, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getMetricTrackingObject", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The metric tracking object id for this server.
func (r Hardware_Server) GetMetricTrackingObjectId() (resp int, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getMetricTrackingObjectId", nil, r.Session, &r.Options, &resp)
	return
}

// Returns open monitoring alarms for a given time period
func (r Hardware_Server) GetMonitoringActiveAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getMonitoringActiveAlarms", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the monitoring agents associated with a piece of hardware.
func (r Hardware_Server) GetMonitoringAgents() (resp []datatypes.Monitoring_Agent, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getMonitoringAgents", nil, r.Session, &r.Options, &resp)
	return
}

// Returns closed monitoring alarms for a given time period
func (r Hardware_Server) GetMonitoringClosedAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getMonitoringClosedAlarms", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the hardware's monitoring robot.
func (r Hardware_Server) GetMonitoringRobot() (resp datatypes.Monitoring_Robot, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getMonitoringRobot", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's network monitoring services.
func (r Hardware_Server) GetMonitoringServiceComponent() (resp datatypes.Network_Monitor_Version1_Query_Host_Stratum, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getMonitoringServiceComponent", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The monitoring service flag eligibility status for a piece of hardware.
func (r Hardware_Server) GetMonitoringServiceEligibilityFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getMonitoringServiceEligibilityFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The service flag status for a piece of hardware.
func (r Hardware_Server) GetMonitoringServiceFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getMonitoringServiceFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The monitoring notification objects for this hardware. Each object links this hardware instance to a user account that will be notified if monitoring on this hardware object fails
func (r Hardware_Server) GetMonitoringUserNotification() (resp []datatypes.User_Customer_Notification_Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getMonitoringUserNotification", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's motherboard.
func (r Hardware_Server) GetMotherboard() (resp datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getMotherboard", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's network cards.
func (r Hardware_Server) GetNetworkCards() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getNetworkCards", nil, r.Session, &r.Options, &resp)
	return
}

// Get the IP addresses associated with this server that are protectable by a network component firewall. Note, this may not return all values for IPv6 subnets for this server. Please use getFirewallProtectableSubnets to get all protectable subnets.
func (r Hardware_Server) GetNetworkComponentFirewallProtectableIpAddresses() (resp []datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getNetworkComponentFirewallProtectableIpAddresses", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Returns a hardware's network components.
func (r Hardware_Server) GetNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getNetworkComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The gateway member if this device is part of a network gateway.
func (r Hardware_Server) GetNetworkGatewayMember() (resp datatypes.Network_Gateway_Member, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getNetworkGatewayMember", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Whether or not this device is part of a network gateway.
func (r Hardware_Server) GetNetworkGatewayMemberFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getNetworkGatewayMemberFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's network management IP address.
func (r Hardware_Server) GetNetworkManagementIpAddress() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getNetworkManagementIpAddress", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All servers with failed monitoring that are attached downstream to a piece of hardware.
func (r Hardware_Server) GetNetworkMonitorAttachedDownHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getNetworkMonitorAttachedDownHardware", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Virtual guests that are attached downstream to a hardware that have failed monitoring
func (r Hardware_Server) GetNetworkMonitorAttachedDownVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getNetworkMonitorAttachedDownVirtualGuests", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The status of all of a piece of hardware's network monitoring incidents.
func (r Hardware_Server) GetNetworkMonitorIncidents() (resp []datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getNetworkMonitorIncidents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's network monitors.
func (r Hardware_Server) GetNetworkMonitors() (resp []datatypes.Network_Monitor_Version1_Query_Host, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getNetworkMonitors", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The value of a hardware's network status attribute.
func (r Hardware_Server) GetNetworkStatus() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getNetworkStatus", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware's related network status attribute.
func (r Hardware_Server) GetNetworkStatusAttribute() (resp datatypes.Hardware_Attribute, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getNetworkStatusAttribute", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's associated network storage service account.
func (r Hardware_Server) GetNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getNetworkStorage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The network virtual LANs (VLANs) associated with a piece of hardware's network components.
func (r Hardware_Server) GetNetworkVlans() (resp []datatypes.Network_Vlan, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getNetworkVlans", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's allotted bandwidth for the next billing cycle (measured in GB).
func (r Hardware_Server) GetNextBillingCycleBandwidthAllocation() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getNextBillingCycleBandwidthAllocation", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Server) GetNotesHistory() (resp []datatypes.Hardware_Note, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getNotesHistory", nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Hardware_Server object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Hardware service. You can only retrieve servers from the account that your portal user is assigned to.
func (r Hardware_Server) GetObject() (resp datatypes.Hardware_Server, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getObject", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An open ticket requesting cancellation of this server, if one exists.
func (r Hardware_Server) GetOpenCancellationTicket() (resp datatypes.Ticket, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getOpenCancellationTicket", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's operating system.
func (r Hardware_Server) GetOperatingSystem() (resp datatypes.Software_Component_OperatingSystem, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getOperatingSystem", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's operating system software description.
func (r Hardware_Server) GetOperatingSystemReferenceCode() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getOperatingSystemReferenceCode", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The sum of all the outbound network traffic data for the last 30 days.
func (r Hardware_Server) GetOutboundBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getOutboundBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total private outbound bandwidth for this hardware for the current billing cycle.
func (r Hardware_Server) GetOutboundPrivateBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getOutboundPrivateBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total public outbound bandwidth for this hardware for the current billing cycle.
func (r Hardware_Server) GetOutboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getOutboundPublicBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Whether the bandwidth usage for this hardware for the current billing cycle exceeds the allocation.
func (r Hardware_Server) GetOverBandwidthAllocationFlag() (resp int, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getOverBandwidthAllocationFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a server's hardware state via its internal sensors. Remote sensor data is transmitted to the SoftLayer API by way of the server's remote management card. Sensor data measures system temperatures, voltages, and other local server settings. Sensor data is cached for 30 seconds. Calls made to getSensorData for the same server within 30 seconds of each other will return the same data. Subsequent calls will return new data once the cache expires.
func (r Hardware_Server) GetPMInfo() (resp []datatypes.Container_RemoteManagement_PmInfo, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getPMInfo", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the Point of Presence (PoP) location in which a piece of hardware resides.
func (r Hardware_Server) GetPointOfPresenceLocation() (resp datatypes.Location, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getPointOfPresenceLocation", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The power components for a hardware object.
func (r Hardware_Server) GetPowerComponents() (resp []datatypes.Hardware_Power_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getPowerComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's power supply.
func (r Hardware_Server) GetPowerSupply() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getPowerSupply", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware's primary private IP address.
func (r Hardware_Server) GetPrimaryBackendIpAddress() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getPrimaryBackendIpAddress", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the hardware's primary back-end network component.
func (r Hardware_Server) GetPrimaryBackendNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getPrimaryBackendNetworkComponent", nil, r.Session, &r.Options, &resp)
	return
}

// no documentation yet
func (r Hardware_Server) GetPrimaryDriveSize() (resp int, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getPrimaryDriveSize", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware's primary public IP address.
func (r Hardware_Server) GetPrimaryIpAddress() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getPrimaryIpAddress", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the hardware's primary public network component.
func (r Hardware_Server) GetPrimaryNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getPrimaryNetworkComponent", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a graph of a server's private network bandwidth usage over the specified timeframe. If no timeframe is specified then getPublicBandwidthGraphImage retrieves the last 24 hours of public bandwidth usage. getPrivateBandwidthGraphImage returns a PNG image measuring 827 pixels by 293 pixels.
func (r Hardware_Server) GetPrivateBandwidthData(startTime *int, endTime *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getPrivateBandwidthData", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve a brief summary of a server's private network bandwidth usage. getPrivateBandwidthDataSummary retrieves a server's bandwidth allocation for its billing period, its estimated usage during its billing period, and an estimation of how much bandwidth it will use during its billing period based on its current usage. A server's projected bandwidth usage increases in accuracy as it progresses through its billing period.
func (r Hardware_Server) GetPrivateBandwidthDataSummary() (resp datatypes.Container_Network_Bandwidth_Data_Summary, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getPrivateBandwidthDataSummary", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a graph of a server's private network bandwidth usage over the specified time frame. If no time frame is specified then getPublicBandwidthGraphImage retrieves the last 24 hours of public bandwidth usage. getPublicBandwidthGraphImage returns a PNG image
func (r Hardware_Server) GetPrivateBandwidthGraphImage(startTime *string, endTime *string) (resp []byte, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getPrivateBandwidthGraphImage", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A server's primary private IP address.
func (r Hardware_Server) GetPrivateIpAddress() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getPrivateIpAddress", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the private network component attached with this server.
func (r Hardware_Server) GetPrivateNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getPrivateNetworkComponent", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Whether the hardware only has access to the private network.
func (r Hardware_Server) GetPrivateNetworkOnlyFlag() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getPrivateNetworkOnlyFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the backend VLAN for the primary IP address of the server
func (r Hardware_Server) GetPrivateVlan() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getPrivateVlan", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a backend network VLAN by searching for an IP address
func (r Hardware_Server) GetPrivateVlanByIpAddress(ipAddress *string) (resp datatypes.Network_Vlan, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getPrivateVlanByIpAddress", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total number of processor cores, summed from all processors that are attached to a piece of hardware
func (r Hardware_Server) GetProcessorCoreAmount() (resp uint, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getProcessorCoreAmount", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total number of physical processor cores, summed from all processors that are attached to a piece of hardware
func (r Hardware_Server) GetProcessorPhysicalCoreAmount() (resp uint, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getProcessorPhysicalCoreAmount", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's processors.
func (r Hardware_Server) GetProcessors() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getProcessors", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Whether the bandwidth usage for this hardware for the current billing cycle is projected to exceed the allocation.
func (r Hardware_Server) GetProjectedOverBandwidthAllocationFlag() (resp int, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getProjectedOverBandwidthAllocationFlag", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The projected public outbound bandwidth for this hardware for the current billing cycle.
func (r Hardware_Server) GetProjectedPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getProjectedPublicBandwidthUsage", nil, r.Session, &r.Options, &resp)
	return
}

// no documentation yet
func (r Hardware_Server) GetProvisionDate() (resp datatypes.Time, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getProvisionDate", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a graph of a server's public network bandwidth usage over the specified timeframe. If no timeframe is specified then getPublicBandwidthGraphImage retrieves the last 24 hours of public bandwidth usage. getPublicBandwidthGraphImage returns a PNG image measuring 827 pixels by 293 pixels.
func (r Hardware_Server) GetPublicBandwidthData(startTime *int, endTime *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getPublicBandwidthData", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve a brief summary of a server's public network bandwidth usage. getPublicBandwidthDataSummary retrieves a server's bandwidth allocation for its billing period, its estimated usage during its billing period, and an estimation of how much bandwidth it will use during its billing period based on its current usage. A server's projected bandwidth usage increases in accuracy as it progresses through its billing period.
func (r Hardware_Server) GetPublicBandwidthDataSummary() (resp datatypes.Container_Network_Bandwidth_Data_Summary, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getPublicBandwidthDataSummary", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a graph of a server's public network bandwidth usage over the specified time frame. If no time frame is specified then getPublicBandwidthGraphImage retrieves the last 24 hours of public bandwidth usage. getPublicBandwidthGraphImage returns a PNG image measuring 827 pixels by 293 pixels.  THIS METHOD GENERATES GRAPHS BASED ON THE NEW DATA WAREHOUSE REPOSITORY.
func (r Hardware_Server) GetPublicBandwidthGraphImage(startTime *datatypes.Time, endTime *datatypes.Time) (resp []byte, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getPublicBandwidthGraphImage", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve the total number of bytes used by a server over a specified time period via the data warehouse tracking objects for this hardware.
func (r Hardware_Server) GetPublicBandwidthTotal(startTime *int, endTime *int) (resp uint, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getPublicBandwidthTotal", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve a SoftLayer server's public network component. Some servers are only connected to the private network and may not have a public network component. In that case getPublicNetworkComponent returns a null object.
func (r Hardware_Server) GetPublicNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getPublicNetworkComponent", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the frontend VLAN for the primary IP address of the server
func (r Hardware_Server) GetPublicVlan() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getPublicVlan", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the frontend network Vlan by searching the hostname of a server
func (r Hardware_Server) GetPublicVlanByHostname(hostname *string) (resp datatypes.Network_Vlan, err error) {
	params := []interface{}{
		hostname,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getPublicVlanByHostname", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Server) GetRack() (resp datatypes.Location, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getRack", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The RAID controllers contained within a piece of hardware.
func (r Hardware_Server) GetRaidControllers() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getRaidControllers", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Recent events that impact this hardware.
func (r Hardware_Server) GetRecentEvents() (resp []datatypes.Notification_Occurrence_Event, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getRecentEvents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The last five commands issued to the server's remote management card.
func (r Hardware_Server) GetRecentRemoteManagementCommands() (resp []datatypes.Hardware_Component_RemoteManagement_Command_Request, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getRecentRemoteManagementCommands", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Server) GetRegionalInternetRegistry() (resp datatypes.Network_Regional_Internet_Registry, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getRegionalInternetRegistry", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A server's remote management card.
func (r Hardware_Server) GetRemoteManagement() (resp datatypes.Hardware_Component_RemoteManagement, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getRemoteManagement", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve User credentials to issue commands and/or interact with the server's remote management card.
func (r Hardware_Server) GetRemoteManagementAccounts() (resp []datatypes.Hardware_Component_RemoteManagement_User, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getRemoteManagementAccounts", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's associated remote management component. This is normally IPMI.
func (r Hardware_Server) GetRemoteManagementComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getRemoteManagementComponent", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve User(s) who have access to issue commands and/or interact with the server's remote management card.
func (r Hardware_Server) GetRemoteManagementUsers() (resp []datatypes.Hardware_Component_RemoteManagement_User, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getRemoteManagementUsers", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Server) GetResourceGroupMemberReferences() (resp []datatypes.Resource_Group_Member, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getResourceGroupMemberReferences", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Server) GetResourceGroupRoles() (resp []datatypes.Resource_Group_Role, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getResourceGroupRoles", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The resource groups in which this hardware is a member.
func (r Hardware_Server) GetResourceGroups() (resp []datatypes.Resource_Group, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getResourceGroups", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the reverse domain records associated with this server.
func (r Hardware_Server) GetReverseDomainRecords() (resp []datatypes.Dns_Domain, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getReverseDomainRecords", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's routers.
func (r Hardware_Server) GetRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getRouters", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Collection of scale assets this hardware corresponds to.
func (r Hardware_Server) GetScaleAssets() (resp []datatypes.Scale_Asset, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getScaleAssets", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's vulnerability scan requests.
func (r Hardware_Server) GetSecurityScanRequests() (resp []datatypes.Network_Security_Scanner_Request, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getSecurityScanRequests", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a server's hardware state via its internal sensors. Remote sensor data is transmitted to the SoftLayer API by way of the server's remote management card. Sensor data measures system temperatures, voltages, and other local server settings. Sensor data is cached for 30 seconds. Calls made to getSensorData for the same server within 30 seconds of each other will return the same data. Subsequent calls will return new data once the cache expires.
func (r Hardware_Server) GetSensorData() (resp []datatypes.Container_RemoteManagement_SensorReading, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getSensorData", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieves the raw data returned from the server's remote management card.  For more details of what is returned please refer to the getSensorData method.  Along with the raw data, graphs for the cpu and system temperatures and fan speeds are also returned.
func (r Hardware_Server) GetSensorDataWithGraphs() (resp datatypes.Container_RemoteManagement_SensorReadingsWithGraphs, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getSensorDataWithGraphs", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a server's hardware components, software, and network components. getServerDetails is an aggregation function that combines the results of [[SoftLayer_Hardware_Server::getComponents]], [[SoftLayer_Hardware_Server::getSoftware]], and [[SoftLayer_Hardware_Server::getNetworkComponents]] in a single container.
func (r Hardware_Server) GetServerDetails() (resp datatypes.Container_Hardware_Server_Details, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getServerDetails", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the server's fan speeds and displays them using tachometer graphs.  Data used to construct graphs is retrieved from the server's remote management card.  All graphs returned will have a title associated with it.
func (r Hardware_Server) GetServerFanSpeedGraphs() (resp []datatypes.Container_RemoteManagement_Graphs_SensorSpeed, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getServerFanSpeedGraphs", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieves the power state for the server.  The server's power status is retrieved from its remote management card.  This will return 'on' or 'off'.
func (r Hardware_Server) GetServerPowerState() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getServerPowerState", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the server room in which the hardware is located.
func (r Hardware_Server) GetServerRoom() (resp datatypes.Location, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getServerRoom", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the server's temperature and displays them using thermometer graphs.  Temperatures retrieved are CPU(s) and system temperatures.  Data used to construct graphs is retrieved from the server's remote management card.  All graphs returned will have a title associated with it.
func (r Hardware_Server) GetServerTemperatureGraphs() (resp []datatypes.Container_RemoteManagement_Graphs_SensorTemperature, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getServerTemperatureGraphs", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the piece of hardware's service provider.
func (r Hardware_Server) GetServiceProvider() (resp datatypes.Service_Provider, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getServiceProvider", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's installed software.
func (r Hardware_Server) GetSoftwareComponents() (resp []datatypes.Software_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getSoftwareComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the billing item for a spare pool server.
func (r Hardware_Server) GetSparePoolBillingItem() (resp datatypes.Billing_Item_Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getSparePoolBillingItem", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve SSH keys to be installed on the server during provisioning or an OS reload.
func (r Hardware_Server) GetSshKeys() (resp []datatypes.Security_Ssh_Key, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getSshKeys", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A server's remote management card used for statistics.
func (r Hardware_Server) GetStatisticsRemoteManagement() (resp datatypes.Hardware_Component_RemoteManagement, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getStatisticsRemoteManagement", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Server) GetStorageNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getStorageNetworkComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Server) GetTagReferences() (resp []datatypes.Tag_Reference, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getTagReferences", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Server) GetTopLevelLocation() (resp datatypes.Location, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getTopLevelLocation", nil, r.Session, &r.Options, &resp)
	return
}

//
// This method will query transaction history for a piece of hardware.
func (r Hardware_Server) GetTransactionHistory() (resp []datatypes.Provisioning_Version1_Transaction_History, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getTransactionHistory", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a list of upgradeable items available to this piece of hardware. Currently, getUpgradeItemPrices retrieves upgrades available for a server's memory, hard drives, network port speed, bandwidth allocation and GPUs.
func (r Hardware_Server) GetUpgradeItemPrices() (resp []datatypes.Product_Item_Price, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getUpgradeItemPrices", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An account's associated upgrade request object, if any.
func (r Hardware_Server) GetUpgradeRequest() (resp datatypes.Product_Upgrade_Request, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getUpgradeRequest", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The network device connected to a piece of hardware.
func (r Hardware_Server) GetUplinkHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getUplinkHardware", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the network component that is one level higher than a piece of hardware on the network infrastructure.
func (r Hardware_Server) GetUplinkNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getUplinkNetworkComponents", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A string containing custom user data for a hardware order.
func (r Hardware_Server) GetUserData() (resp []datatypes.Hardware_Attribute, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getUserData", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A list of users that have access to this computing instance.
func (r Hardware_Server) GetUsers() (resp []datatypes.User_Customer, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getUsers", nil, r.Session, &r.Options, &resp)
	return
}

// This method will return the list of block device template groups that are valid to the host. For instance, it will only retrieve FLEX images.
func (r Hardware_Server) GetValidBlockDeviceTemplateGroups(visibility *string) (resp []datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	params := []interface{}{
		visibility,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "getValidBlockDeviceTemplateGroups", params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the virtual chassis for a piece of hardware.
func (r Hardware_Server) GetVirtualChassis() (resp datatypes.Hardware_Group, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getVirtualChassis", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the virtual chassis siblings for a piece of hardware.
func (r Hardware_Server) GetVirtualChassisSiblings() (resp []datatypes.Hardware, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getVirtualChassisSiblings", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware server's virtual servers.
func (r Hardware_Server) GetVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getVirtualGuests", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's virtual host record.
func (r Hardware_Server) GetVirtualHost() (resp datatypes.Virtual_Host, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getVirtualHost", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's virtual software licenses.
func (r Hardware_Server) GetVirtualLicenses() (resp []datatypes.Software_VirtualLicense, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getVirtualLicenses", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the bandwidth allotment to which a piece of hardware belongs.
func (r Hardware_Server) GetVirtualRack() (resp datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getVirtualRack", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The name of the bandwidth allotment belonging to a piece of hardware.
func (r Hardware_Server) GetVirtualRackId() (resp int, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getVirtualRackId", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The name of the bandwidth allotment belonging to a piece of hardware.
func (r Hardware_Server) GetVirtualRackName() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getVirtualRackName", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's virtualization platform software.
func (r Hardware_Server) GetVirtualizationPlatform() (resp datatypes.Software_Component, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getVirtualizationPlatform", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a list of Windows updates available for a server from the local SoftLayer Windows Server Update Services (WSUS) server. Windows servers provisioned by SoftLayer are configured to use the local WSUS server via the private network by default.
func (r Hardware_Server) GetWindowsUpdateAvailableUpdates() (resp []datatypes.Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getWindowsUpdateAvailableUpdates", nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a list of Windows updates installed on a server as reported by the local SoftLayer Windows Server Update Services (WSUS) server. Windows servers provisioned by SoftLayer are configured to use the local WSUS server via the private network by default.
func (r Hardware_Server) GetWindowsUpdateInstalledUpdates() (resp []datatypes.Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getWindowsUpdateInstalledUpdates", nil, r.Session, &r.Options, &resp)
	return
}

// This method returns an update status record for this server.  That record will specify if the server is missing updates, or has updates that must be reinstalled or require a reboot to go into affect.
func (r Hardware_Server) GetWindowsUpdateStatus() (resp datatypes.Container_Utility_Microsoft_Windows_UpdateServices_Status, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "getWindowsUpdateStatus", nil, r.Session, &r.Options, &resp)
	return
}

// The '''importVirtualHost''' method attempts to import the host record for the virtualization platform running on a server.
func (r Hardware_Server) ImportVirtualHost() (resp datatypes.Virtual_Host, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "importVirtualHost", nil, r.Session, &r.Options, &resp)
	return
}

// Idera Bare Metal Server Restore is a backup agent designed specifically for making full system restores made with Idera Server Backup.
func (r Hardware_Server) InitiateIderaBareMetalRestore() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "initiateIderaBareMetalRestore", nil, r.Session, &r.Options, &resp)
	return
}

// R1Soft Bare Metal Server Restore is an R1Soft disk agent designed specifically for making full system restores made with R1Soft CDP Server backup.
func (r Hardware_Server) InitiateR1SoftBareMetalRestore() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "initiateR1SoftBareMetalRestore", nil, r.Session, &r.Options, &resp)
	return
}

// Issues a ping command and returns the success (true) or failure (false) of the ping command.
func (r Hardware_Server) IsBackendPingable() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "isBackendPingable", nil, r.Session, &r.Options, &resp)
	return
}

// Issues a ping command and returns the success (true) or failure (false) of the ping command.
func (r Hardware_Server) IsPingable() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "isPingable", nil, r.Session, &r.Options, &resp)
	return
}

// Determine if the server runs any version of the Microsoft Windows operating systems. Return ''true'' if it does and ''false if otherwise.
func (r Hardware_Server) IsWindowsServer() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "isWindowsServer", nil, r.Session, &r.Options, &resp)
	return
}

// Issues a ping command to the server and returns the ping response.
func (r Hardware_Server) Ping() (resp string, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "ping", nil, r.Session, &r.Options, &resp)
	return
}

// Power off then power on the server via powerstrip.  The power cycle command is equivalent to unplugging the server from the powerstrip and then plugging the server back into the powerstrip.  This should only be used as a last resort.  If a reboot command has been issued successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard, powerOn, powerOff and powerCycle) will not be allowed. This is to avoid any type of server failures.
func (r Hardware_Server) PowerCycle() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "powerCycle", nil, r.Session, &r.Options, &resp)
	return
}

// This method will power off the server via the server's remote management card.
func (r Hardware_Server) PowerOff() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "powerOff", nil, r.Session, &r.Options, &resp)
	return
}

// Power on server via its remote management card.  If a reboot command has been issued successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard, powerOn, powerOff and powerCycle) will not be allowed.  This is to avoid any type of server failures.
func (r Hardware_Server) PowerOn() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "powerOn", nil, r.Session, &r.Options, &resp)
	return
}

// Attempts to reboot the server by issuing a reset (soft reboot) command to the server's remote management card. If the reset (soft reboot) attempt is unsuccessful, a power cycle command will be issued via the powerstrip. The power cycle command is equivalent to unplugging the server from the powerstrip and then plugging the server back into the powerstrip.  If a reboot command has been issued successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard, powerOn, powerOff and powerCycle) will not be allowed.  This is to avoid any type of server failures.
func (r Hardware_Server) RebootDefault() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "rebootDefault", nil, r.Session, &r.Options, &resp)
	return
}

// Reboot the server by issuing a cycle command to the server's remote management card.  This is equivalent to pressing the 'Reset' button on the server.  This command is issued immediately and will not wait for processes to shutdown. After this command is issued, the server may take a few moments to boot up as server may run system disks checks. If a reboot command has been issued successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard, powerOn, powerOff and powerCycle) will not be allowed.  This is to avoid any type of server failures.
func (r Hardware_Server) RebootHard() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "rebootHard", nil, r.Session, &r.Options, &resp)
	return
}

// Reboot the server by issuing a reset command to the server's remote management card.  This is a graceful reboot. The servers will allow all process to shutdown gracefully before rebooting.  If a reboot command has been issued successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard, powerOn, powerOff and powerCycle) will not be allowed.  This is to avoid any type of server failures.
func (r Hardware_Server) RebootSoft() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "rebootSoft", nil, r.Session, &r.Options, &resp)
	return
}

// Reloads current operating system configuration.
//
// This service has a confirmation protocol for proceeding with the reload. To proceed with the reload without confirmation, simply pass in 'FORCE' as the token parameter. To proceed with the reload with confirmation, simply call the service with no parameter. A token string will be returned by this service. The token will remain active for 10 minutes. Use this token as the parameter to confirm that a reload is to be performed for the server.
//
// As a precaution, we strongly  recommend backing up all data before reloading the operating system. The reload will format the primary disk and will reconfigure the server to the current specifications on record.
//
// The reload will take AT MINIMUM 66 minutes.
func (r Hardware_Server) ReloadCurrentOperatingSystemConfiguration(token *string) (resp string, err error) {
	params := []interface{}{
		token,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "reloadCurrentOperatingSystemConfiguration", params, r.Session, &r.Options, &resp)
	return
}

// Reloads current or customer specified operating system configuration.
//
// This service has a confirmation protocol for proceeding with the reload. To proceed with the reload without confirmation, simply pass in 'FORCE' as the token parameter. To proceed with the reload with confirmation, simply call the service with no parameter. A token string will be returned by this service. The token will remain active for 10 minutes. Use this token as the parameter to confirm that a reload is to be performed for the server.
//
// As a precaution, we strongly  recommend backing up all data before reloading the operating system. The reload will format the primary disk and will reconfigure the server to the current specifications on record.
//
// The reload will take AT MINIMUM 66 minutes.
func (r Hardware_Server) ReloadOperatingSystem(token *string, config *datatypes.Container_Hardware_Server_Configuration) (resp string, err error) {
	params := []interface{}{
		token,
		config,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "reloadOperatingSystem", params, r.Session, &r.Options, &resp)
	return
}

// This method is used to remove access to s SoftLayer_Network_Storage volumes that supports host- or network-level access control.
func (r Hardware_Server) RemoveAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "removeAccessToNetworkStorage", params, r.Session, &r.Options, &resp)
	return
}

// This method is used to allow access to multiple SoftLayer_Network_Storage volumes that support host- or network-level access control.
func (r Hardware_Server) RemoveAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "removeAccessToNetworkStorageList", params, r.Session, &r.Options, &resp)
	return
}

// You can launch a new Passmark hardware test by selecting from your server list. It will bring your server offline for approximately 20 minutes while the testing is in progress, and will publish a certificate with the results to your hardware details page.
//
// While the hard drives are tested for the initial deployment, the Passmark Certificate utility will not test the hard drives on your live server. This is to ensure that no data is overwritten. If you would like to test the server's hard drives, you can have the full Passmark suite installed to your server free of charge through a new Support ticket.
//
// While the test itself does not overwrite any data on the server, it is recommended that you make full off-server backups of all data prior to launching the test. The Passmark hardware test is designed to force any latent hardware issues to the surface, so hardware failure is possible.
//
// In the event of a hardware failure during this test our datacenter engineers will be notified of the problem automatically. They will then replace any failed components to bring your server back online, and will be contacting you to ensure that impact on your server is minimal.
func (r Hardware_Server) RunPassmarkCertificationBenchmark() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "runPassmarkCertificationBenchmark", nil, r.Session, &r.Options, &resp)
	return
}

// Changes the password that we have stored in our database for a servers' Operating System
func (r Hardware_Server) SetOperatingSystemPassword(newPassword *string) (resp bool, err error) {
	params := []interface{}{
		newPassword,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "setOperatingSystemPassword", params, r.Session, &r.Options, &resp)
	return
}

// Sets the private network interface speed to the new speed. Speed values can only be 0 (Disconnect), 10, 100, 1000, and 10000. The new speed must be equal to or less than the max speed of the interface.
//
// It will take less than a minute to update the switch port speed. The server uplink will not be operational again until the server interface speed is updated.
func (r Hardware_Server) SetPrivateNetworkInterfaceSpeed(newSpeed *int) (resp bool, err error) {
	params := []interface{}{
		newSpeed,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "setPrivateNetworkInterfaceSpeed", params, r.Session, &r.Options, &resp)
	return
}

// Sets the public network interface speed to the new speed. Speed values can only be 0 (Disconnect), 10, 100, 1000, and 10000. The new speed must be equal to or less than the max speed of the interface.
//
// It will take less than a minute to update the switch port speed. The server uplink will not be operational again until the server interface speed is updated.
func (r Hardware_Server) SetPublicNetworkInterfaceSpeed(newSpeed *int) (resp bool, err error) {
	params := []interface{}{
		newSpeed,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "setPublicNetworkInterfaceSpeed", params, r.Session, &r.Options, &resp)
	return
}

// no documentation yet
func (r Hardware_Server) SetTags(tags *string) (resp bool, err error) {
	params := []interface{}{
		tags,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "setTags", params, r.Session, &r.Options, &resp)
	return
}

// Sets the data that will be written to the configuration drive.
func (r Hardware_Server) SetUserMetadata(metadata []string) (resp []datatypes.Hardware_Attribute, err error) {
	params := []interface{}{
		metadata,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "setUserMetadata", params, r.Session, &r.Options, &resp)
	return
}

// Shuts down the public network port
func (r Hardware_Server) ShutdownPrivatePort() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "shutdownPrivatePort", nil, r.Session, &r.Options, &resp)
	return
}

// Shuts down the public network port
func (r Hardware_Server) ShutdownPublicPort() (resp bool, err error) {
	err = invokeMethod("SoftLayer_Hardware_Server", "shutdownPublicPort", nil, r.Session, &r.Options, &resp)
	return
}

// The ability to place bare metal servers in a state where they are powered down, and ports closed yet still allocated to the customer as a part of the Spare Pool program.
func (r Hardware_Server) SparePool(action *string, newOrder *bool) (resp bool, err error) {
	params := []interface{}{
		action,
		newOrder,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "sparePool", params, r.Session, &r.Options, &resp)
	return
}

// Validates a collection of partitions for an operating system
func (r Hardware_Server) ValidatePartitionsForOperatingSystem(operatingSystem *datatypes.Software_Description, partitions []datatypes.Hardware_Component_Partition) (resp bool, err error) {
	params := []interface{}{
		operatingSystem,
		partitions,
	}
	err = invokeMethod("SoftLayer_Hardware_Server", "validatePartitionsForOperatingSystem", params, r.Session, &r.Options, &resp)
	return
}
