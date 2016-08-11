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
	"github.ibm.com/riethm/gopherlayer/datatypes"
	"github.ibm.com/riethm/gopherlayer/session"
	"github.ibm.com/riethm/gopherlayer/sl"
)

// The SoftLayer_Hardware data type contains general information relating to a single SoftLayer hardware.
type Hardware struct {
	Session *session.Session
	sl.Options
}

func GetHardwareService(sess *session.Session) Hardware {
	return Hardware{Session: sess}
}

// This method is used to allow access to a SoftLayer_Network_Storage volume that supports host- or network-level access control.
func (r *Hardware) AllowAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// This method is used to allow access to multiple SoftLayer_Network_Storage volumes that support host- or network-level access control.
func (r *Hardware) AllowAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Captures a Flex Image of the hard disk on the physical machine, based on the capture template parameter. Returns the image template group containing the disk image.
func (r *Hardware) CaptureImage(captureTemplate *datatypes.Container_Disk_Image_Capture_Template) (resp datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	params := []interface{}{
		captureTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Returns monitoring alarm detailed history
func (r *Hardware) CloseAlarm(alarmId *string) (resp bool, err error) {
	params := []interface{}{
		alarmId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
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
func (r *Hardware) CreateObject(templateObject *datatypes.Hardware) (resp datatypes.Hardware, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
// This method will cancel a server effective immediately. For servers billed hourly, the charges will stop immediately after the method returns.
func (r *Hardware) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Delete software component passwords.
func (r *Hardware) DeleteSoftwareComponentPasswords(softwareComponentPasswords []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		softwareComponentPasswords,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Edit the properties of a software component password such as the username, password, and notes.
func (r *Hardware) EditSoftwareComponentPasswords(softwareComponentPasswords []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		softwareComponentPasswords,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Download and run remote script from uri on the hardware.
func (r *Hardware) ExecuteRemoteScript(uri *string) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		uri,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// The '''findByIpAddress''' method finds hardware using its primary public or private IP address. IP addresses that have a secondary subnet tied to the hardware will not return the hardware - alternate means of locating the hardware must be used (see '''Associated Methods'''). If no hardware is found, no errors are generated and no data is returned.
func (r *Hardware) FindByIpAddress(ipAddress *string) (resp datatypes.Hardware, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
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
func (r *Hardware) GenerateOrderTemplate(templateObject *datatypes.Hardware) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The account associated with a piece of hardware.
func (r *Hardware) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's active physical components.
func (r *Hardware) GetActiveComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's active network monitoring incidents.
func (r *Hardware) GetActiveNetworkMonitorIncident() (resp []datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getAlarmHistory''' method retrieves a detailed history for the monitoring alarm. When calling this method, a start and end date for the history to be retrieved must be entered.
func (r *Hardware) GetAlarmHistory(startDate *datatypes.Time, endDate *datatypes.Time, alarmId *string) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
		alarmId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware) GetAllPowerComponents() (resp []datatypes.Hardware_Power_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Storage_Allowed_Host information to connect this server to Network Storage volumes that require access control lists.
func (r *Hardware) GetAllowedHost() (resp datatypes.Network_Storage_Allowed_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Storage objects that this SoftLayer_Hardware has access to.
func (r *Hardware) GetAllowedNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Storage objects whose Replica that this SoftLayer_Hardware has access to.
func (r *Hardware) GetAllowedNetworkStorageReplicas() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding an antivirus/spyware software component object.
func (r *Hardware) GetAntivirusSpywareSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This method is retrieve a list of SoftLayer_Network_Storage volumes that are authorized access to this SoftLayer_Hardware.
func (r *Hardware) GetAttachedNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's specific attributes.
func (r *Hardware) GetAttributes() (resp []datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This method retrieves a list of SoftLayer_Network_Storage volumes that can be authorized to this SoftLayer_Hardware.
func (r *Hardware) GetAvailableNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The average daily public bandwidth usage for the current billing cycle.
func (r *Hardware) GetAverageDailyPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getBackendIncomingBandwidth''' method retrieves the amount of incoming private network traffic used between the given start date and end date parameters. When entering start and end dates, only the month, day and year are used to calculate bandwidth totals - the time (HH:MM:SS) is ignored and defaults to midnight. The amount of bandwidth retrieved is measured in gigabytes.
func (r *Hardware) GetBackendIncomingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's back-end or private network components.
func (r *Hardware) GetBackendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getBackendOutgoingBandwidth''' method retrieves the amount of outgoing private network traffic used between the given start date and end date parameters. When entering start and end dates, only the month, day and year are used to calculate bandwidth totals - the time (HH:MM:SS) is ignored and defaults to midnight. The amount of bandwidth retrieved is measured in gigabytes.
func (r *Hardware) GetBackendOutgoingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's backend or private router.
func (r *Hardware) GetBackendRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's allotted bandwidth (measured in GB).
func (r *Hardware) GetBandwidthAllocation() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's allotted detail record. Allotment details link bandwidth allocation with allotments.
func (r *Hardware) GetBandwidthAllotmentDetail() (resp datatypes.Network_Bandwidth_Version1_Allotment_Detail, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's benchmark certifications.
func (r *Hardware) GetBenchmarkCertifications() (resp []datatypes.Hardware_Benchmark_Certification, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the billing item for a server.
func (r *Hardware) GetBillingItem() (resp datatypes.Billing_Item_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A flag indicating that a billing item exists.
func (r *Hardware) GetBillingItemFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Determines whether the hardware is ineligible for cancellation because it is disconnected.
func (r *Hardware) GetBlockCancelBecauseDisconnectedFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Status indicating whether or not a piece of hardware has business continuance insurance.
func (r *Hardware) GetBusinessContinuanceInsuranceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's components.
func (r *Hardware) GetComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A continuous data protection/server backup software component object.
func (r *Hardware) GetContinuousDataProtectionSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
// There are many options that may be provided while ordering a server, this method can be used to determine what these options are.
//
//
// Detailed information on the return value can be found on the data type page for [[SoftLayer_Container_Hardware_Configuration (type)]].
func (r *Hardware) GetCreateObjectOptions() (resp datatypes.Container_Hardware_Configuration, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The current billable public outbound bandwidth for this hardware for the current billing cycle.
func (r *Hardware) GetCurrentBillableBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Get the billing detail for this instance for the current billing period. This does not include bandwidth usage.
func (r *Hardware) GetCurrentBillingDetail() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getCurrentBillingTotal''' method retrieves the total bill amount in US Dollars ($) for the current billing period. In addition to the total bill amount, the billing detail also includes all bandwidth used up to the point the method is called on the piece of hardware.
func (r *Hardware) GetCurrentBillingTotal() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getDailyAverage''' method calculates the average daily network traffic used by the selected server. Using the required parameter ''dateTime'' to enter a start and end date, the user retrieves this average, measure in gigabytes (GB) for the specified date range. When entering parameters, only the month, day and year are required - time entries are omitted as this method defaults the time to midnight in order to account for the entire day.
func (r *Hardware) GetDailyAverage(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the datacenter in which a piece of hardware resides.
func (r *Hardware) GetDatacenter() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The name of the datacenter in which a piece of hardware resides.
func (r *Hardware) GetDatacenterName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All hardware that has uplink network connections to a piece of hardware.
func (r *Hardware) GetDownlinkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All hardware that has uplink network connections to a piece of hardware.
func (r *Hardware) GetDownlinkNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all servers attached to a piece of network hardware.
func (r *Hardware) GetDownlinkServers() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all virtual guests attached to a piece of network hardware.
func (r *Hardware) GetDownlinkVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All hardware downstream from a network device.
func (r *Hardware) GetDownstreamHardwareBindings() (resp []datatypes.Network_Component_Uplink_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All network hardware downstream from the selected piece of hardware.
func (r *Hardware) GetDownstreamNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All network hardware with monitoring warnings or errors that are downstream from the selected piece of hardware.
func (r *Hardware) GetDownstreamNetworkHardwareWithIncidents() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all servers attached downstream to a piece of network hardware.
func (r *Hardware) GetDownstreamServers() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all virtual guests attached to a piece of network hardware.
func (r *Hardware) GetDownstreamVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The drive controllers contained within a piece of hardware.
func (r *Hardware) GetDriveControllers() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's associated EVault network storage service account.
func (r *Hardware) GetEvaultNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's firewall services.
func (r *Hardware) GetFirewallServiceComponent() (resp datatypes.Network_Component_Firewall, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Defines the fixed components in a fixed configuration bare metal server.
func (r *Hardware) GetFixedConfigurationPreset() (resp datatypes.Product_Package_Preset, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getFrontendIncomingBandwidth''' method retrieves the amount of incoming public network traffic used by a server between the given start and end date parameters. When entering the ''dateTime'' parameter, only the month, day and year of the start and end dates are required - the time (hour, minute and second) are set to midnight by default and cannot be changed. The amount of bandwidth retrieved is measured in gigabytes (GB).
func (r *Hardware) GetFrontendIncomingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's front-end or public network components.
func (r *Hardware) GetFrontendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getFrontendOutgoingBandwidth''' method retrieves the amount of outgoing public network traffic used by a server between the given start and end date parameters. The ''dateTime'' parameter requires only the day, month and year to be entered - the time (hour, minute and second) are set to midnight be default in order to gather the data for the entire start and end date indicated in the parameter. The amount of bandwidth retrieved is measured in gigabytes (GB).
func (r *Hardware) GetFrontendOutgoingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's frontend or public router.
func (r *Hardware) GetFrontendRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's universally unique identifier.
func (r *Hardware) GetGlobalIdentifier() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hard drives contained within a piece of hardware.
func (r *Hardware) GetHardDrives() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The chassis that a piece of hardware is housed in.
func (r *Hardware) GetHardwareChassis() (resp datatypes.Hardware_Chassis, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's function.
func (r *Hardware) GetHardwareFunction() (resp datatypes.Hardware_Function, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's function.
func (r *Hardware) GetHardwareFunctionDescription() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's status.
func (r *Hardware) GetHardwareStatus() (resp datatypes.Hardware_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Determine in hardware object has TPM enabled.
func (r *Hardware) GetHasTrustedPlatformModuleBillingItemFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a host IPS software component object.
func (r *Hardware) GetHostIpsSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getHourlyBandwidth''' method retrieves all bandwidth updates hourly for the specified hardware. Because the potential number of data points can become excessive, the method limits the user to obtain data in 24-hour intervals. The required ''dateTime'' parameter is used as the starting point for the query and will be calculated for the 24-hour period starting with the specified date and time. For example, entering a parameter of
//
// '02/01/2008 0:00'
//
// results in a return of all bandwidth data for the entire day of February 1, 2008, as 0:00 specifies a midnight start date. Please note that the time entered should be completed using a 24-hour clock (military time, astronomical time).
//
// For data spanning more than a single 24-hour period, refer to the getBandwidthData function on the metricTrackingObject for the piece of hardware.
func (r *Hardware) GetHourlyBandwidth(mode *string, day *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		mode,
		day,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A server's hourly billing status.
func (r *Hardware) GetHourlyBillingFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The sum of all the inbound network traffic data for the last 30 days.
func (r *Hardware) GetInboundBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total public inbound bandwidth for this hardware for the current billing cycle.
func (r *Hardware) GetInboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the last transaction a server performed.
func (r *Hardware) GetLastTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's latest network monitoring incident.
func (r *Hardware) GetLatestNetworkMonitorIncident() (resp datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Where a piece of hardware is located within SoftLayer's location hierarchy.
func (r *Hardware) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware) GetLocationPathString() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a lockbox account associated with a server.
func (r *Hardware) GetLockboxNetworkStorage() (resp datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A flag indicating that the hardware is a managed resource.
func (r *Hardware) GetManagedResourceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's memory.
func (r *Hardware) GetMemory() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The amount of memory a piece of hardware has, measured in gigabytes.
func (r *Hardware) GetMemoryCapacity() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's metric tracking object.
func (r *Hardware) GetMetricTrackingObject() (resp datatypes.Metric_Tracking_Object_HardwareServer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Returns open monitoring alarms for a given time period
func (r *Hardware) GetMonitoringActiveAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the monitoring agents associated with a piece of hardware.
func (r *Hardware) GetMonitoringAgents() (resp []datatypes.Monitoring_Agent, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Returns closed monitoring alarms for a given time period
func (r *Hardware) GetMonitoringClosedAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the hardware's monitoring robot.
func (r *Hardware) GetMonitoringRobot() (resp datatypes.Monitoring_Robot, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's network monitoring services.
func (r *Hardware) GetMonitoringServiceComponent() (resp datatypes.Network_Monitor_Version1_Query_Host_Stratum, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The monitoring service flag eligibility status for a piece of hardware.
func (r *Hardware) GetMonitoringServiceEligibilityFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The service flag status for a piece of hardware.
func (r *Hardware) GetMonitoringServiceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's motherboard.
func (r *Hardware) GetMotherboard() (resp datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's network cards.
func (r *Hardware) GetNetworkCards() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Returns a hardware's network components.
func (r *Hardware) GetNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The gateway member if this device is part of a network gateway.
func (r *Hardware) GetNetworkGatewayMember() (resp datatypes.Network_Gateway_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Whether or not this device is part of a network gateway.
func (r *Hardware) GetNetworkGatewayMemberFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's network management IP address.
func (r *Hardware) GetNetworkManagementIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All servers with failed monitoring that are attached downstream to a piece of hardware.
func (r *Hardware) GetNetworkMonitorAttachedDownHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Virtual guests that are attached downstream to a hardware that have failed monitoring
func (r *Hardware) GetNetworkMonitorAttachedDownVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The status of all of a piece of hardware's network monitoring incidents.
func (r *Hardware) GetNetworkMonitorIncidents() (resp []datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's network monitors.
func (r *Hardware) GetNetworkMonitors() (resp []datatypes.Network_Monitor_Version1_Query_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The value of a hardware's network status attribute.
func (r *Hardware) GetNetworkStatus() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware's related network status attribute.
func (r *Hardware) GetNetworkStatusAttribute() (resp datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's associated network storage service account.
func (r *Hardware) GetNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The network virtual LANs (VLANs) associated with a piece of hardware's network components.
func (r *Hardware) GetNetworkVlans() (resp []datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's allotted bandwidth for the next billing cycle (measured in GB).
func (r *Hardware) GetNextBillingCycleBandwidthAllocation() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware) GetNotesHistory() (resp []datatypes.Hardware_Note, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Hardware object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Hardware service. You can only retrieve the account that your portal user is assigned to.
func (r *Hardware) GetObject() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's operating system.
func (r *Hardware) GetOperatingSystem() (resp datatypes.Software_Component_OperatingSystem, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's operating system software description.
func (r *Hardware) GetOperatingSystemReferenceCode() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The sum of all the outbound network traffic data for the last 30 days.
func (r *Hardware) GetOutboundBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total public outbound bandwidth for this hardware for the current billing cycle.
func (r *Hardware) GetOutboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the Point of Presence (PoP) location in which a piece of hardware resides.
func (r *Hardware) GetPointOfPresenceLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The power components for a hardware object.
func (r *Hardware) GetPowerComponents() (resp []datatypes.Hardware_Power_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's power supply.
func (r *Hardware) GetPowerSupply() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware's primary private IP address.
func (r *Hardware) GetPrimaryBackendIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the hardware's primary back-end network component.
func (r *Hardware) GetPrimaryBackendNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware's primary public IP address.
func (r *Hardware) GetPrimaryIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the hardware's primary public network component.
func (r *Hardware) GetPrimaryNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a graph of a server's private network bandwidth usage over the specified timeframe. If no timeframe is specified then getPublicBandwidthGraphImage retrieves the last 24 hours of public bandwidth usage. getPrivateBandwidthGraphImage returns a PNG image measuring 827 pixels by 293 pixels.
func (r *Hardware) GetPrivateBandwidthData(startTime *int, endTime *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Whether the hardware only has access to the private network.
func (r *Hardware) GetPrivateNetworkOnlyFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total number of processor cores, summed from all processors that are attached to a piece of hardware
func (r *Hardware) GetProcessorCoreAmount() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total number of physical processor cores, summed from all processors that are attached to a piece of hardware
func (r *Hardware) GetProcessorPhysicalCoreAmount() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's processors.
func (r *Hardware) GetProcessors() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a graph of a server's public network bandwidth usage over the specified timeframe. If no timeframe is specified then getPublicBandwidthGraphImage retrieves the last 24 hours of public bandwidth usage. getPublicBandwidthGraphImage returns a PNG image measuring 827 pixels by 293 pixels.
func (r *Hardware) GetPublicBandwidthData(startTime *int, endTime *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware) GetRack() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The RAID controllers contained within a piece of hardware.
func (r *Hardware) GetRaidControllers() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Recent events that impact this hardware.
func (r *Hardware) GetRecentEvents() (resp []datatypes.Notification_Occurrence_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve User credentials to issue commands and/or interact with the server's remote management card.
func (r *Hardware) GetRemoteManagementAccounts() (resp []datatypes.Hardware_Component_RemoteManagement_User, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's associated remote management component. This is normally IPMI.
func (r *Hardware) GetRemoteManagementComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware) GetResourceGroupMemberReferences() (resp []datatypes.Resource_Group_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware) GetResourceGroupRoles() (resp []datatypes.Resource_Group_Role, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The resource groups in which this hardware is a member.
func (r *Hardware) GetResourceGroups() (resp []datatypes.Resource_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's routers.
func (r *Hardware) GetRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Collection of scale assets this hardware corresponds to.
func (r *Hardware) GetScaleAssets() (resp []datatypes.Scale_Asset, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's vulnerability scan requests.
func (r *Hardware) GetSecurityScanRequests() (resp []datatypes.Network_Security_Scanner_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getSensorData''' method retrieves a server's hardware state via its internal sensors. Remote sensor data is transmitted to the SoftLayer API by way of the server's remote management card. Sensor data measures various information, including system temperatures, voltages and other local server settings. Sensor data is cached for 30 second; calls made to this method for the same server within 30 seconds of each other will result in the same data being returned. To ensure that the data retrieved retrieves snapshot of varied data, make calls greater than 30 seconds apart.
func (r *Hardware) GetSensorData() (resp []datatypes.Container_RemoteManagement_SensorReading, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getSensorDataWithGraphs''' method retrieves the raw data returned from the server's remote management card. Along with raw data, graphs for the CPU and system temperatures and fan speeds are also returned. For more details on what information is returned, refer to the ''getSensorData'' method.
func (r *Hardware) GetSensorDataWithGraphs() (resp datatypes.Container_RemoteManagement_SensorReadingsWithGraphs, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getServerFanSpeedGraphs''' method retrieves the server's fan speeds and displays the speeds using tachometer graphs. data used to construct these graphs is retrieved from the server's remote management card. Each graph returned will have an associated title.
func (r *Hardware) GetServerFanSpeedGraphs() (resp []datatypes.Container_RemoteManagement_Graphs_SensorSpeed, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getPowerState''' method retrieves the power state for the selected server. The server's power status is retrieved from its remote management card. This method returns "on", for a server that has been powered on, or "off" for servers powered off.
func (r *Hardware) GetServerPowerState() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the server room in which the hardware is located.
func (r *Hardware) GetServerRoom() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getServerTemperatureGraphs''' retrieves the server's temperatures and displays the various temperatures using thermometer graphs. Temperatures retrieved are CPU temperature(s) and system temperatures. Data used to construct the graphs is retrieved from the server's remote management card. All graphs returned will have an associated title.
func (r *Hardware) GetServerTemperatureGraphs() (resp []datatypes.Container_RemoteManagement_Graphs_SensorTemperature, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the piece of hardware's service provider.
func (r *Hardware) GetServiceProvider() (resp datatypes.Service_Provider, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's installed software.
func (r *Hardware) GetSoftwareComponents() (resp []datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the billing item for a spare pool server.
func (r *Hardware) GetSparePoolBillingItem() (resp datatypes.Billing_Item_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve SSH keys to be installed on the server during provisioning or an OS reload.
func (r *Hardware) GetSshKeys() (resp []datatypes.Security_Ssh_Key, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware) GetStorageNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware) GetTagReferences() (resp []datatypes.Tag_Reference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware) GetTopLevelLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
// This method will query transaction history for a piece of hardware.
func (r *Hardware) GetTransactionHistory() (resp []datatypes.Provisioning_Version1_Transaction_History, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a list of upgradeable items available to this piece of hardware. Currently, getUpgradeItemPrices retrieves upgrades available for a server's memory, hard drives, network port speed, bandwidth allocation and GPUs.
func (r *Hardware) GetUpgradeItemPrices() (resp []datatypes.Product_Item_Price, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An account's associated upgrade request object, if any.
func (r *Hardware) GetUpgradeRequest() (resp datatypes.Product_Upgrade_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The network device connected to a piece of hardware.
func (r *Hardware) GetUplinkHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the network component that is one level higher than a piece of hardware on the network infrastructure.
func (r *Hardware) GetUplinkNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A string containing custom user data for a hardware order.
func (r *Hardware) GetUserData() (resp []datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the virtual chassis for a piece of hardware.
func (r *Hardware) GetVirtualChassis() (resp datatypes.Hardware_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the virtual chassis siblings for a piece of hardware.
func (r *Hardware) GetVirtualChassisSiblings() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's virtual host record.
func (r *Hardware) GetVirtualHost() (resp datatypes.Virtual_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's virtual software licenses.
func (r *Hardware) GetVirtualLicenses() (resp []datatypes.Software_VirtualLicense, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the bandwidth allotment to which a piece of hardware belongs.
func (r *Hardware) GetVirtualRack() (resp datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The name of the bandwidth allotment belonging to a piece of hardware.
func (r *Hardware) GetVirtualRackId() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The name of the bandwidth allotment belonging to a piece of hardware.
func (r *Hardware) GetVirtualRackName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's virtualization platform software.
func (r *Hardware) GetVirtualizationPlatform() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''importVirtualHost''' method attempts to import the host record for the virtualization platform running on a server.
func (r *Hardware) ImportVirtualHost() (resp datatypes.Virtual_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''isPingable''' method issues a ping command to the selected server and returns the result of the ping command. This boolean return value displays ''true'' upon successful ping or ''false'' for a failed ping.
func (r *Hardware) IsPingable() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Issues a ping command to the server and returns the ping response.
func (r *Hardware) Ping() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''powerCycle''' method completes a power off and power on of the server successively in one command. The power cycle command is equivalent to unplugging the server from the power strip and then plugging the server back in. '''This method should only be used when all other options have been exhausted'''. Additional remote management commands may not be executed if this command was successfully issued within the last 20 minutes to avoid server failure. Remote management commands include:
//
// rebootSoft rebootHard powerOn powerOff powerCycle
//
//
func (r *Hardware) PowerCycle() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This method will power off the server via the server's remote management card.
func (r *Hardware) PowerOff() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''powerOn''' method powers on a server via its remote management card. This boolean return value returns ''true'' upon successful execution and ''false'' if unsuccessful. Other remote management commands may not be issued in this command was successfully completed within the last 20 minutes to avoid server failure. Remote management commands include:
//
// rebootSoft rebootHard powerOn powerOff powerCycle
//
//
func (r *Hardware) PowerOn() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''rebootDefault''' method attempts to reboot the server by issuing a soft reboot, or reset, command to the server's remote management card. if the reset attempt is unsuccessful, a power cycle command will be issued via the power strip. The power cycle command is equivalent to unplugging the server from the power strip and then plugging the server back in. If the reset was successful within the last 20 minutes, another remote management command cannot be completed to avoid server failure. Remote management commands include:
//
// rebootSoft rebootHard powerOn powerOff powerCycle
//
//
func (r *Hardware) RebootDefault() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''rebootHard''' method reboots the server by issuing a cycle command to the server's remote management card. A hard reboot is equivalent to pressing the ''Reset'' button on a server - it is issued immediately and will not allow processes to shut down prior to the reboot. Completing a hard reboot may initiate system disk checks upon server reboot, causing the boot up to take longer than normally expected.
//
// Remote management commands are unable to be executed if a reboot has been issued successfully within the last 20 minutes to avoid server failure. Remote management commands include:
//
// rebootSoft rebootHard powerOn powerOff powerCycle
//
//
func (r *Hardware) RebootHard() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''rebootSoft''' method reboots the server by issuing a reset command to the server's remote management card via soft reboot. When executing a soft reboot, servers allow all processes to shut down completely before rebooting. Remote management commands are unable to be issued within 20 minutes of issuing a successful soft reboot in order to avoid server failure. Remote management commands include:
//
// rebootSoft rebootHard powerOn powerOff powerCycle
//
//
func (r *Hardware) RebootSoft() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This method is used to remove access to s SoftLayer_Network_Storage volumes that supports host- or network-level access control.
func (r *Hardware) RemoveAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// This method is used to allow access to multiple SoftLayer_Network_Storage volumes that support host- or network-level access control.
func (r *Hardware) RemoveAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Hardware) SetTags(tags *string) (resp bool, err error) {
	params := []interface{}{
		tags,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// The SoftLayer_Hardware_Benchmark_Certification data type contains general information relating to a single SoftLayer hardware benchmark certification document.
type Hardware_Benchmark_Certification struct {
	Session *session.Session
	sl.Options
}

func GetHardwareBenchmarkCertificationService(sess *session.Session) Hardware_Benchmark_Certification {
	return Hardware_Benchmark_Certification{Session: sess}
}

// Retrieve Information regarding a benchmark certification result's associated SoftLayer customer account.
func (r *Hardware_Benchmark_Certification) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the piece of hardware on which a benchmark certification test was performed.
func (r *Hardware_Benchmark_Certification) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Hardware_Benchmark_Certification object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Hardware_Benchmark_Certification service.
func (r *Hardware_Benchmark_Certification) GetObject() (resp datatypes.Hardware_Benchmark_Certification, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Attempt to retrieve the file associated with a benchmark certification result, if such a file exists.  If there is no file for this benchmark certification result, calling this method throws an exception. The "getResultFile" method attempts to retrieve the file associated with a benchmark certification result, if such a file exists. If no file exists for the benchmark certification, an exception is thrown.
func (r *Hardware_Benchmark_Certification) GetResultFile() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The SoftLayer_Hardware_Component_Model data type contains general information relating to a single SoftLayer component model.  A component model represents a vendor specific representation of a hardware component.  Every piece of hardware on a server will have a specific hardware component model.
type Hardware_Component_Model struct {
	Session *session.Session
	sl.Options
}

func GetHardwareComponentModelService(sess *session.Session) Hardware_Component_Model {
	return Hardware_Component_Model{Session: sess}
}

// Retrieve
func (r *Hardware_Component_Model) GetArchitectureType() (resp datatypes.Hardware_Component_Model_Architecture_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_Component_Model) GetAttributes() (resp []datatypes.Hardware_Component_Model_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_Component_Model) GetCompatibleArrayTypes() (resp []datatypes.Configuration_Storage_Group_Array_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All the component models that are compatible with a hardware component model.
func (r *Hardware_Component_Model) GetCompatibleChildComponentModels() (resp []datatypes.Hardware_Component_Model, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All the component models that a hardware component model is compatible with.
func (r *Hardware_Component_Model) GetCompatibleParentComponentModels() (resp []datatypes.Hardware_Component_Model, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware component model's physical components in inventory.
func (r *Hardware_Component_Model) GetHardwareComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The non-vendor specific generic component model for a hardware component model.
func (r *Hardware_Component_Model) GetHardwareGenericComponentModel() (resp datatypes.Hardware_Component_Model_Generic, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_Component_Model) GetInfinibandCompatibleAttribute() (resp datatypes.Hardware_Component_Model_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_Component_Model) GetIsInfinibandCompatible() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Hardware_Component_Model object.
func (r *Hardware_Component_Model) GetObject() (resp datatypes.Hardware_Component_Model, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A motherboard's average reboot time.
func (r *Hardware_Component_Model) GetRebootTime() (resp datatypes.Hardware_Component_Motherboard_Reboot_Time, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware component model's type.
func (r *Hardware_Component_Model) GetType() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The types of attributes that are allowed for a given hardware component model.
func (r *Hardware_Component_Model) GetValidAttributeTypes() (resp []datatypes.Hardware_Component_Model_Attribute_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The SoftLayer_Hardware_Component_Partition_OperatingSystem data type contains general information relating to a single SoftLayer Operating System Partition Template.
type Hardware_Component_Partition_OperatingSystem struct {
	Session *session.Session
	sl.Options
}

func GetHardwareComponentPartitionOperatingSystemService(sess *session.Session) Hardware_Component_Partition_OperatingSystem {
	return Hardware_Component_Partition_OperatingSystem{Session: sess}
}

//
func (r *Hardware_Component_Partition_OperatingSystem) GetAllObjects() (resp []datatypes.Hardware_Component_Partition_OperatingSystem, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getByDescription''' method retrieves all possible partition templates based on the description (required parameter) entered when calling the method. The description is typically the operating system's name. Current recognized values include 'linux', 'windows', 'freebsd', and 'Debian'.
func (r *Hardware_Component_Partition_OperatingSystem) GetByDescription(description *string) (resp datatypes.Hardware_Component_Partition_OperatingSystem, err error) {
	params := []interface{}{
		description,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Hardware_Component_Partition_OperatingSystem object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Hardware_Component_Partition_OperatingSystem service.s
func (r *Hardware_Component_Partition_OperatingSystem) GetObject() (resp datatypes.Hardware_Component_Partition_OperatingSystem, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding an operating system's [[SoftLayer_Hardware_Component_Partition_Template|Partition Templates]].
func (r *Hardware_Component_Partition_OperatingSystem) GetPartitionTemplates() (resp []datatypes.Hardware_Component_Partition_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The SoftLayer_Hardware_Component_Partition_Template data type contains general information relating to a single SoftLayer partition template.  Partition templates group 1 or more partition configurations that can be used to predefined how a hard drives partitions will be configured.
type Hardware_Component_Partition_Template struct {
	Session *session.Session
	sl.Options
}

func GetHardwareComponentPartitionTemplateService(sess *session.Session) Hardware_Component_Partition_Template {
	return Hardware_Component_Partition_Template{Session: sess}
}

// Retrieve A partition template's associated [[SoftLayer_Account|Account]].
func (r *Hardware_Component_Partition_Template) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An individual partition for a partition template. This is identical to 'partitionTemplatePartition' except this will sort unix partitions.
func (r *Hardware_Component_Partition_Template) GetData() (resp []datatypes.Hardware_Component_Partition_Template_Partition, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_Component_Partition_Template) GetExpireDate() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Hardware_Component_Partition_Template object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Hardware_Component_Partition_Template service. You can only retrieve the partition templates that your account created or the templates predefined by SoftLayer.
func (r *Hardware_Component_Partition_Template) GetObject() (resp datatypes.Hardware_Component_Partition_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A partition template's associated [[SoftLayer_Hardware_Component_Partition_OperatingSystem|Operating System]].
func (r *Hardware_Component_Partition_Template) GetPartitionOperatingSystem() (resp datatypes.Hardware_Component_Partition_OperatingSystem, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An individual partition for a partition template.
func (r *Hardware_Component_Partition_Template) GetPartitionTemplatePartition() (resp []datatypes.Hardware_Component_Partition_Template_Partition, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The SoftLayer_Hardware_Router data type contains general information relating to a single SoftLayer router.
type Hardware_Router struct {
	Session *session.Session
	sl.Options
}

func GetHardwareRouterService(sess *session.Session) Hardware_Router {
	return Hardware_Router{Session: sess}
}

// This method is used to allow access to a SoftLayer_Network_Storage volume that supports host- or network-level access control.
func (r *Hardware_Router) AllowAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// This method is used to allow access to multiple SoftLayer_Network_Storage volumes that support host- or network-level access control.
func (r *Hardware_Router) AllowAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Captures a Flex Image of the hard disk on the physical machine, based on the capture template parameter. Returns the image template group containing the disk image.
func (r *Hardware_Router) CaptureImage(captureTemplate *datatypes.Container_Disk_Image_Capture_Template) (resp datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	params := []interface{}{
		captureTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Returns monitoring alarm detailed history
func (r *Hardware_Router) CloseAlarm(alarmId *string) (resp bool, err error) {
	params := []interface{}{
		alarmId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
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
func (r *Hardware_Router) CreateObject(templateObject *datatypes.Hardware) (resp datatypes.Hardware, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
// This method will cancel a server effective immediately. For servers billed hourly, the charges will stop immediately after the method returns.
func (r *Hardware_Router) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Delete software component passwords.
func (r *Hardware_Router) DeleteSoftwareComponentPasswords(softwareComponentPasswords []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		softwareComponentPasswords,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Edit the properties of a software component password such as the username, password, and notes.
func (r *Hardware_Router) EditSoftwareComponentPasswords(softwareComponentPasswords []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		softwareComponentPasswords,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Download and run remote script from uri on the hardware.
func (r *Hardware_Router) ExecuteRemoteScript(uri *string) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		uri,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// The '''findByIpAddress''' method finds hardware using its primary public or private IP address. IP addresses that have a secondary subnet tied to the hardware will not return the hardware - alternate means of locating the hardware must be used (see '''Associated Methods'''). If no hardware is found, no errors are generated and no data is returned.
func (r *Hardware_Router) FindByIpAddress(ipAddress *string) (resp datatypes.Hardware, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
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
func (r *Hardware_Router) GenerateOrderTemplate(templateObject *datatypes.Hardware) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The account associated with a piece of hardware.
func (r *Hardware_Router) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's active physical components.
func (r *Hardware_Router) GetActiveComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's active network monitoring incidents.
func (r *Hardware_Router) GetActiveNetworkMonitorIncident() (resp []datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getAlarmHistory''' method retrieves a detailed history for the monitoring alarm. When calling this method, a start and end date for the history to be retrieved must be entered.
func (r *Hardware_Router) GetAlarmHistory(startDate *datatypes.Time, endDate *datatypes.Time, alarmId *string) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
		alarmId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_Router) GetAllPowerComponents() (resp []datatypes.Hardware_Power_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Storage_Allowed_Host information to connect this server to Network Storage volumes that require access control lists.
func (r *Hardware_Router) GetAllowedHost() (resp datatypes.Network_Storage_Allowed_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Storage objects that this SoftLayer_Hardware has access to.
func (r *Hardware_Router) GetAllowedNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Storage objects whose Replica that this SoftLayer_Hardware has access to.
func (r *Hardware_Router) GetAllowedNetworkStorageReplicas() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding an antivirus/spyware software component object.
func (r *Hardware_Router) GetAntivirusSpywareSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This method is retrieve a list of SoftLayer_Network_Storage volumes that are authorized access to this SoftLayer_Hardware.
func (r *Hardware_Router) GetAttachedNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's specific attributes.
func (r *Hardware_Router) GetAttributes() (resp []datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This method retrieves a list of SoftLayer_Network_Storage volumes that can be authorized to this SoftLayer_Hardware.
func (r *Hardware_Router) GetAvailableNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The average daily public bandwidth usage for the current billing cycle.
func (r *Hardware_Router) GetAverageDailyPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getBackendIncomingBandwidth''' method retrieves the amount of incoming private network traffic used between the given start date and end date parameters. When entering start and end dates, only the month, day and year are used to calculate bandwidth totals - the time (HH:MM:SS) is ignored and defaults to midnight. The amount of bandwidth retrieved is measured in gigabytes.
func (r *Hardware_Router) GetBackendIncomingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's back-end or private network components.
func (r *Hardware_Router) GetBackendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getBackendOutgoingBandwidth''' method retrieves the amount of outgoing private network traffic used between the given start date and end date parameters. When entering start and end dates, only the month, day and year are used to calculate bandwidth totals - the time (HH:MM:SS) is ignored and defaults to midnight. The amount of bandwidth retrieved is measured in gigabytes.
func (r *Hardware_Router) GetBackendOutgoingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's backend or private router.
func (r *Hardware_Router) GetBackendRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's allotted bandwidth (measured in GB).
func (r *Hardware_Router) GetBandwidthAllocation() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's allotted detail record. Allotment details link bandwidth allocation with allotments.
func (r *Hardware_Router) GetBandwidthAllotmentDetail() (resp datatypes.Network_Bandwidth_Version1_Allotment_Detail, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's benchmark certifications.
func (r *Hardware_Router) GetBenchmarkCertifications() (resp []datatypes.Hardware_Benchmark_Certification, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the billing item for a server.
func (r *Hardware_Router) GetBillingItem() (resp datatypes.Billing_Item_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A flag indicating that a billing item exists.
func (r *Hardware_Router) GetBillingItemFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Determines whether the hardware is ineligible for cancellation because it is disconnected.
func (r *Hardware_Router) GetBlockCancelBecauseDisconnectedFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Associated subnets for a router object.
func (r *Hardware_Router) GetBoundSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Status indicating whether or not a piece of hardware has business continuance insurance.
func (r *Hardware_Router) GetBusinessContinuanceInsuranceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's components.
func (r *Hardware_Router) GetComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A continuous data protection/server backup software component object.
func (r *Hardware_Router) GetContinuousDataProtectionSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
// There are many options that may be provided while ordering a server, this method can be used to determine what these options are.
//
//
// Detailed information on the return value can be found on the data type page for [[SoftLayer_Container_Hardware_Configuration (type)]].
func (r *Hardware_Router) GetCreateObjectOptions() (resp datatypes.Container_Hardware_Configuration, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The current billable public outbound bandwidth for this hardware for the current billing cycle.
func (r *Hardware_Router) GetCurrentBillableBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Get the billing detail for this instance for the current billing period. This does not include bandwidth usage.
func (r *Hardware_Router) GetCurrentBillingDetail() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getCurrentBillingTotal''' method retrieves the total bill amount in US Dollars ($) for the current billing period. In addition to the total bill amount, the billing detail also includes all bandwidth used up to the point the method is called on the piece of hardware.
func (r *Hardware_Router) GetCurrentBillingTotal() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getDailyAverage''' method calculates the average daily network traffic used by the selected server. Using the required parameter ''dateTime'' to enter a start and end date, the user retrieves this average, measure in gigabytes (GB) for the specified date range. When entering parameters, only the month, day and year are required - time entries are omitted as this method defaults the time to midnight in order to account for the entire day.
func (r *Hardware_Router) GetDailyAverage(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the datacenter in which a piece of hardware resides.
func (r *Hardware_Router) GetDatacenter() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The name of the datacenter in which a piece of hardware resides.
func (r *Hardware_Router) GetDatacenterName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All hardware that has uplink network connections to a piece of hardware.
func (r *Hardware_Router) GetDownlinkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All hardware that has uplink network connections to a piece of hardware.
func (r *Hardware_Router) GetDownlinkNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all servers attached to a piece of network hardware.
func (r *Hardware_Router) GetDownlinkServers() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all virtual guests attached to a piece of network hardware.
func (r *Hardware_Router) GetDownlinkVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All hardware downstream from a network device.
func (r *Hardware_Router) GetDownstreamHardwareBindings() (resp []datatypes.Network_Component_Uplink_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All network hardware downstream from the selected piece of hardware.
func (r *Hardware_Router) GetDownstreamNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All network hardware with monitoring warnings or errors that are downstream from the selected piece of hardware.
func (r *Hardware_Router) GetDownstreamNetworkHardwareWithIncidents() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all servers attached downstream to a piece of network hardware.
func (r *Hardware_Router) GetDownstreamServers() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all virtual guests attached to a piece of network hardware.
func (r *Hardware_Router) GetDownstreamVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The drive controllers contained within a piece of hardware.
func (r *Hardware_Router) GetDriveControllers() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's associated EVault network storage service account.
func (r *Hardware_Router) GetEvaultNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's firewall services.
func (r *Hardware_Router) GetFirewallServiceComponent() (resp datatypes.Network_Component_Firewall, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Defines the fixed components in a fixed configuration bare metal server.
func (r *Hardware_Router) GetFixedConfigurationPreset() (resp datatypes.Product_Package_Preset, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getFrontendIncomingBandwidth''' method retrieves the amount of incoming public network traffic used by a server between the given start and end date parameters. When entering the ''dateTime'' parameter, only the month, day and year of the start and end dates are required - the time (hour, minute and second) are set to midnight by default and cannot be changed. The amount of bandwidth retrieved is measured in gigabytes (GB).
func (r *Hardware_Router) GetFrontendIncomingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's front-end or public network components.
func (r *Hardware_Router) GetFrontendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getFrontendOutgoingBandwidth''' method retrieves the amount of outgoing public network traffic used by a server between the given start and end date parameters. The ''dateTime'' parameter requires only the day, month and year to be entered - the time (hour, minute and second) are set to midnight be default in order to gather the data for the entire start and end date indicated in the parameter. The amount of bandwidth retrieved is measured in gigabytes (GB).
func (r *Hardware_Router) GetFrontendOutgoingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's frontend or public router.
func (r *Hardware_Router) GetFrontendRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's universally unique identifier.
func (r *Hardware_Router) GetGlobalIdentifier() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hard drives contained within a piece of hardware.
func (r *Hardware_Router) GetHardDrives() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The chassis that a piece of hardware is housed in.
func (r *Hardware_Router) GetHardwareChassis() (resp datatypes.Hardware_Chassis, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's function.
func (r *Hardware_Router) GetHardwareFunction() (resp datatypes.Hardware_Function, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's function.
func (r *Hardware_Router) GetHardwareFunctionDescription() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's status.
func (r *Hardware_Router) GetHardwareStatus() (resp datatypes.Hardware_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Determine in hardware object has TPM enabled.
func (r *Hardware_Router) GetHasTrustedPlatformModuleBillingItemFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a host IPS software component object.
func (r *Hardware_Router) GetHostIpsSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getHourlyBandwidth''' method retrieves all bandwidth updates hourly for the specified hardware. Because the potential number of data points can become excessive, the method limits the user to obtain data in 24-hour intervals. The required ''dateTime'' parameter is used as the starting point for the query and will be calculated for the 24-hour period starting with the specified date and time. For example, entering a parameter of
//
// '02/01/2008 0:00'
//
// results in a return of all bandwidth data for the entire day of February 1, 2008, as 0:00 specifies a midnight start date. Please note that the time entered should be completed using a 24-hour clock (military time, astronomical time).
//
// For data spanning more than a single 24-hour period, refer to the getBandwidthData function on the metricTrackingObject for the piece of hardware.
func (r *Hardware_Router) GetHourlyBandwidth(mode *string, day *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		mode,
		day,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A server's hourly billing status.
func (r *Hardware_Router) GetHourlyBillingFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The sum of all the inbound network traffic data for the last 30 days.
func (r *Hardware_Router) GetInboundBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total public inbound bandwidth for this hardware for the current billing cycle.
func (r *Hardware_Router) GetInboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the last transaction a server performed.
func (r *Hardware_Router) GetLastTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's latest network monitoring incident.
func (r *Hardware_Router) GetLatestNetworkMonitorIncident() (resp datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A flag indicating that a VLAN on the router can be assigned to a host that has local disk functionality.
func (r *Hardware_Router) GetLocalDiskStorageCapabilityFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Where a piece of hardware is located within SoftLayer's location hierarchy.
func (r *Hardware_Router) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_Router) GetLocationPathString() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a lockbox account associated with a server.
func (r *Hardware_Router) GetLockboxNetworkStorage() (resp datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A flag indicating that the hardware is a managed resource.
func (r *Hardware_Router) GetManagedResourceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's memory.
func (r *Hardware_Router) GetMemory() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The amount of memory a piece of hardware has, measured in gigabytes.
func (r *Hardware_Router) GetMemoryCapacity() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's metric tracking object.
func (r *Hardware_Router) GetMetricTrackingObject() (resp datatypes.Metric_Tracking_Object_HardwareServer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Returns open monitoring alarms for a given time period
func (r *Hardware_Router) GetMonitoringActiveAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the monitoring agents associated with a piece of hardware.
func (r *Hardware_Router) GetMonitoringAgents() (resp []datatypes.Monitoring_Agent, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Returns closed monitoring alarms for a given time period
func (r *Hardware_Router) GetMonitoringClosedAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the hardware's monitoring robot.
func (r *Hardware_Router) GetMonitoringRobot() (resp datatypes.Monitoring_Robot, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's network monitoring services.
func (r *Hardware_Router) GetMonitoringServiceComponent() (resp datatypes.Network_Monitor_Version1_Query_Host_Stratum, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The monitoring service flag eligibility status for a piece of hardware.
func (r *Hardware_Router) GetMonitoringServiceEligibilityFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The service flag status for a piece of hardware.
func (r *Hardware_Router) GetMonitoringServiceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's motherboard.
func (r *Hardware_Router) GetMotherboard() (resp datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's network cards.
func (r *Hardware_Router) GetNetworkCards() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Returns a hardware's network components.
func (r *Hardware_Router) GetNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The gateway member if this device is part of a network gateway.
func (r *Hardware_Router) GetNetworkGatewayMember() (resp datatypes.Network_Gateway_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Whether or not this device is part of a network gateway.
func (r *Hardware_Router) GetNetworkGatewayMemberFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's network management IP address.
func (r *Hardware_Router) GetNetworkManagementIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All servers with failed monitoring that are attached downstream to a piece of hardware.
func (r *Hardware_Router) GetNetworkMonitorAttachedDownHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Virtual guests that are attached downstream to a hardware that have failed monitoring
func (r *Hardware_Router) GetNetworkMonitorAttachedDownVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The status of all of a piece of hardware's network monitoring incidents.
func (r *Hardware_Router) GetNetworkMonitorIncidents() (resp []datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's network monitors.
func (r *Hardware_Router) GetNetworkMonitors() (resp []datatypes.Network_Monitor_Version1_Query_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The value of a hardware's network status attribute.
func (r *Hardware_Router) GetNetworkStatus() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware's related network status attribute.
func (r *Hardware_Router) GetNetworkStatusAttribute() (resp datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's associated network storage service account.
func (r *Hardware_Router) GetNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The network virtual LANs (VLANs) associated with a piece of hardware's network components.
func (r *Hardware_Router) GetNetworkVlans() (resp []datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's allotted bandwidth for the next billing cycle (measured in GB).
func (r *Hardware_Router) GetNextBillingCycleBandwidthAllocation() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_Router) GetNotesHistory() (resp []datatypes.Hardware_Note, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Hardware_Router) GetObject() (resp datatypes.Hardware_Router, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's operating system.
func (r *Hardware_Router) GetOperatingSystem() (resp datatypes.Software_Component_OperatingSystem, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's operating system software description.
func (r *Hardware_Router) GetOperatingSystemReferenceCode() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The sum of all the outbound network traffic data for the last 30 days.
func (r *Hardware_Router) GetOutboundBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total public outbound bandwidth for this hardware for the current billing cycle.
func (r *Hardware_Router) GetOutboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the Point of Presence (PoP) location in which a piece of hardware resides.
func (r *Hardware_Router) GetPointOfPresenceLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The power components for a hardware object.
func (r *Hardware_Router) GetPowerComponents() (resp []datatypes.Hardware_Power_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's power supply.
func (r *Hardware_Router) GetPowerSupply() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware's primary private IP address.
func (r *Hardware_Router) GetPrimaryBackendIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the hardware's primary back-end network component.
func (r *Hardware_Router) GetPrimaryBackendNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware's primary public IP address.
func (r *Hardware_Router) GetPrimaryIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the hardware's primary public network component.
func (r *Hardware_Router) GetPrimaryNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a graph of a server's private network bandwidth usage over the specified timeframe. If no timeframe is specified then getPublicBandwidthGraphImage retrieves the last 24 hours of public bandwidth usage. getPrivateBandwidthGraphImage returns a PNG image measuring 827 pixels by 293 pixels.
func (r *Hardware_Router) GetPrivateBandwidthData(startTime *int, endTime *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Whether the hardware only has access to the private network.
func (r *Hardware_Router) GetPrivateNetworkOnlyFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total number of processor cores, summed from all processors that are attached to a piece of hardware
func (r *Hardware_Router) GetProcessorCoreAmount() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total number of physical processor cores, summed from all processors that are attached to a piece of hardware
func (r *Hardware_Router) GetProcessorPhysicalCoreAmount() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's processors.
func (r *Hardware_Router) GetProcessors() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a graph of a server's public network bandwidth usage over the specified timeframe. If no timeframe is specified then getPublicBandwidthGraphImage retrieves the last 24 hours of public bandwidth usage. getPublicBandwidthGraphImage returns a PNG image measuring 827 pixels by 293 pixels.
func (r *Hardware_Router) GetPublicBandwidthData(startTime *int, endTime *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_Router) GetRack() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The RAID controllers contained within a piece of hardware.
func (r *Hardware_Router) GetRaidControllers() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Recent events that impact this hardware.
func (r *Hardware_Router) GetRecentEvents() (resp []datatypes.Notification_Occurrence_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve User credentials to issue commands and/or interact with the server's remote management card.
func (r *Hardware_Router) GetRemoteManagementAccounts() (resp []datatypes.Hardware_Component_RemoteManagement_User, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's associated remote management component. This is normally IPMI.
func (r *Hardware_Router) GetRemoteManagementComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_Router) GetResourceGroupMemberReferences() (resp []datatypes.Resource_Group_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_Router) GetResourceGroupRoles() (resp []datatypes.Resource_Group_Role, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The resource groups in which this hardware is a member.
func (r *Hardware_Router) GetResourceGroups() (resp []datatypes.Resource_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's routers.
func (r *Hardware_Router) GetRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A flag indicating that a VLAN on the router can be assigned to a host that has SAN disk functionality.
func (r *Hardware_Router) GetSanStorageCapabilityFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Collection of scale assets this hardware corresponds to.
func (r *Hardware_Router) GetScaleAssets() (resp []datatypes.Scale_Asset, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's vulnerability scan requests.
func (r *Hardware_Router) GetSecurityScanRequests() (resp []datatypes.Network_Security_Scanner_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getSensorData''' method retrieves a server's hardware state via its internal sensors. Remote sensor data is transmitted to the SoftLayer API by way of the server's remote management card. Sensor data measures various information, including system temperatures, voltages and other local server settings. Sensor data is cached for 30 second; calls made to this method for the same server within 30 seconds of each other will result in the same data being returned. To ensure that the data retrieved retrieves snapshot of varied data, make calls greater than 30 seconds apart.
func (r *Hardware_Router) GetSensorData() (resp []datatypes.Container_RemoteManagement_SensorReading, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getSensorDataWithGraphs''' method retrieves the raw data returned from the server's remote management card. Along with raw data, graphs for the CPU and system temperatures and fan speeds are also returned. For more details on what information is returned, refer to the ''getSensorData'' method.
func (r *Hardware_Router) GetSensorDataWithGraphs() (resp datatypes.Container_RemoteManagement_SensorReadingsWithGraphs, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getServerFanSpeedGraphs''' method retrieves the server's fan speeds and displays the speeds using tachometer graphs. data used to construct these graphs is retrieved from the server's remote management card. Each graph returned will have an associated title.
func (r *Hardware_Router) GetServerFanSpeedGraphs() (resp []datatypes.Container_RemoteManagement_Graphs_SensorSpeed, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getPowerState''' method retrieves the power state for the selected server. The server's power status is retrieved from its remote management card. This method returns "on", for a server that has been powered on, or "off" for servers powered off.
func (r *Hardware_Router) GetServerPowerState() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the server room in which the hardware is located.
func (r *Hardware_Router) GetServerRoom() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getServerTemperatureGraphs''' retrieves the server's temperatures and displays the various temperatures using thermometer graphs. Temperatures retrieved are CPU temperature(s) and system temperatures. Data used to construct the graphs is retrieved from the server's remote management card. All graphs returned will have an associated title.
func (r *Hardware_Router) GetServerTemperatureGraphs() (resp []datatypes.Container_RemoteManagement_Graphs_SensorTemperature, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the piece of hardware's service provider.
func (r *Hardware_Router) GetServiceProvider() (resp datatypes.Service_Provider, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's installed software.
func (r *Hardware_Router) GetSoftwareComponents() (resp []datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the billing item for a spare pool server.
func (r *Hardware_Router) GetSparePoolBillingItem() (resp datatypes.Billing_Item_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve SSH keys to be installed on the server during provisioning or an OS reload.
func (r *Hardware_Router) GetSshKeys() (resp []datatypes.Security_Ssh_Key, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_Router) GetStorageNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_Router) GetTagReferences() (resp []datatypes.Tag_Reference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_Router) GetTopLevelLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
// This method will query transaction history for a piece of hardware.
func (r *Hardware_Router) GetTransactionHistory() (resp []datatypes.Provisioning_Version1_Transaction_History, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a list of upgradeable items available to this piece of hardware. Currently, getUpgradeItemPrices retrieves upgrades available for a server's memory, hard drives, network port speed, bandwidth allocation and GPUs.
func (r *Hardware_Router) GetUpgradeItemPrices() (resp []datatypes.Product_Item_Price, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An account's associated upgrade request object, if any.
func (r *Hardware_Router) GetUpgradeRequest() (resp datatypes.Product_Upgrade_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The network device connected to a piece of hardware.
func (r *Hardware_Router) GetUplinkHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the network component that is one level higher than a piece of hardware on the network infrastructure.
func (r *Hardware_Router) GetUplinkNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A string containing custom user data for a hardware order.
func (r *Hardware_Router) GetUserData() (resp []datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the virtual chassis for a piece of hardware.
func (r *Hardware_Router) GetVirtualChassis() (resp datatypes.Hardware_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the virtual chassis siblings for a piece of hardware.
func (r *Hardware_Router) GetVirtualChassisSiblings() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's virtual host record.
func (r *Hardware_Router) GetVirtualHost() (resp datatypes.Virtual_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's virtual software licenses.
func (r *Hardware_Router) GetVirtualLicenses() (resp []datatypes.Software_VirtualLicense, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the bandwidth allotment to which a piece of hardware belongs.
func (r *Hardware_Router) GetVirtualRack() (resp datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The name of the bandwidth allotment belonging to a piece of hardware.
func (r *Hardware_Router) GetVirtualRackId() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The name of the bandwidth allotment belonging to a piece of hardware.
func (r *Hardware_Router) GetVirtualRackName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's virtualization platform software.
func (r *Hardware_Router) GetVirtualizationPlatform() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''importVirtualHost''' method attempts to import the host record for the virtualization platform running on a server.
func (r *Hardware_Router) ImportVirtualHost() (resp datatypes.Virtual_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''isPingable''' method issues a ping command to the selected server and returns the result of the ping command. This boolean return value displays ''true'' upon successful ping or ''false'' for a failed ping.
func (r *Hardware_Router) IsPingable() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Issues a ping command to the server and returns the ping response.
func (r *Hardware_Router) Ping() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''powerCycle''' method completes a power off and power on of the server successively in one command. The power cycle command is equivalent to unplugging the server from the power strip and then plugging the server back in. '''This method should only be used when all other options have been exhausted'''. Additional remote management commands may not be executed if this command was successfully issued within the last 20 minutes to avoid server failure. Remote management commands include:
//
// rebootSoft rebootHard powerOn powerOff powerCycle
//
//
func (r *Hardware_Router) PowerCycle() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This method will power off the server via the server's remote management card.
func (r *Hardware_Router) PowerOff() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''powerOn''' method powers on a server via its remote management card. This boolean return value returns ''true'' upon successful execution and ''false'' if unsuccessful. Other remote management commands may not be issued in this command was successfully completed within the last 20 minutes to avoid server failure. Remote management commands include:
//
// rebootSoft rebootHard powerOn powerOff powerCycle
//
//
func (r *Hardware_Router) PowerOn() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''rebootDefault''' method attempts to reboot the server by issuing a soft reboot, or reset, command to the server's remote management card. if the reset attempt is unsuccessful, a power cycle command will be issued via the power strip. The power cycle command is equivalent to unplugging the server from the power strip and then plugging the server back in. If the reset was successful within the last 20 minutes, another remote management command cannot be completed to avoid server failure. Remote management commands include:
//
// rebootSoft rebootHard powerOn powerOff powerCycle
//
//
func (r *Hardware_Router) RebootDefault() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''rebootHard''' method reboots the server by issuing a cycle command to the server's remote management card. A hard reboot is equivalent to pressing the ''Reset'' button on a server - it is issued immediately and will not allow processes to shut down prior to the reboot. Completing a hard reboot may initiate system disk checks upon server reboot, causing the boot up to take longer than normally expected.
//
// Remote management commands are unable to be executed if a reboot has been issued successfully within the last 20 minutes to avoid server failure. Remote management commands include:
//
// rebootSoft rebootHard powerOn powerOff powerCycle
//
//
func (r *Hardware_Router) RebootHard() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''rebootSoft''' method reboots the server by issuing a reset command to the server's remote management card via soft reboot. When executing a soft reboot, servers allow all processes to shut down completely before rebooting. Remote management commands are unable to be issued within 20 minutes of issuing a successful soft reboot in order to avoid server failure. Remote management commands include:
//
// rebootSoft rebootHard powerOn powerOff powerCycle
//
//
func (r *Hardware_Router) RebootSoft() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This method is used to remove access to s SoftLayer_Network_Storage volumes that supports host- or network-level access control.
func (r *Hardware_Router) RemoveAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// This method is used to allow access to multiple SoftLayer_Network_Storage volumes that support host- or network-level access control.
func (r *Hardware_Router) RemoveAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Hardware_Router) SetTags(tags *string) (resp bool, err error) {
	params := []interface{}{
		tags,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
type Hardware_SecurityModule struct {
	Session *session.Session
	sl.Options
}

func GetHardwareSecurityModuleService(sess *session.Session) Hardware_SecurityModule {
	return Hardware_SecurityModule{Session: sess}
}

// Activates the private network port
func (r *Hardware_SecurityModule) ActivatePrivatePort() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Activates the public network port
func (r *Hardware_SecurityModule) ActivatePublicPort() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This method is used to allow access to a SoftLayer_Network_Storage volume that supports host- or network-level access control.
func (r *Hardware_SecurityModule) AllowAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// This method is used to allow access to multiple SoftLayer_Network_Storage volumes that support host- or network-level access control.
func (r *Hardware_SecurityModule) AllowAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// The Rescue Kernel is designed to provide you with the ability to bring a server online in order to troubleshoot system problems that would normally only be resolved by an OS Reload. The correct Rescue Kernel will be selected based upon the currently installed operating system. When the rescue kernel process is initiated, the server will shutdown and reboot on to the public network with the same IP's assigned to the server to allow for remote connections. It will bring your server offline for approximately 10 minutes while the rescue is in progress. The root/administrator password will be the same as what is listed in the portal for the server.
func (r *Hardware_SecurityModule) BootToRescueLayer(noOsBootEnvironment *string) (resp bool, err error) {
	params := []interface{}{
		noOsBootEnvironment,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Captures a Flex Image of the hard disk on the physical machine, based on the capture template parameter. Returns the image template group containing the disk image.
func (r *Hardware_SecurityModule) CaptureImage(captureTemplate *datatypes.Container_Disk_Image_Capture_Template) (resp datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	params := []interface{}{
		captureTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Returns monitoring alarm detailed history
func (r *Hardware_SecurityModule) CloseAlarm(alarmId *string) (resp bool, err error) {
	params := []interface{}{
		alarmId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// You can launch firmware updates by selecting from your server list. It will bring your server offline for approximately 20 minutes while the updates are in progress.
//
// In the event of a hardware failure during this test our datacenter engineers will be notified of the problem automatically. They will then replace any failed components to bring your server back online, and will be contacting you to ensure that impact on your server is minimal.
func (r *Hardware_SecurityModule) CreateFirmwareUpdateTransaction(ipmi *int, raidController *int, bios *int, harddrive *int) (resp bool, err error) {
	params := []interface{}{
		ipmi,
		raidController,
		bios,
		harddrive,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
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
func (r *Hardware_SecurityModule) CreateObject(templateObject *datatypes.Hardware_SecurityModule) (resp datatypes.Hardware_SecurityModule, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Hardware_SecurityModule) CreatePostSoftwareInstallTransaction(installCodes []string, returnBoolean *bool) (resp bool, err error) {
	params := []interface{}{
		installCodes,
		returnBoolean,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
// This method will cancel a server effective immediately. For servers billed hourly, the charges will stop immediately after the method returns.
func (r *Hardware_SecurityModule) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Delete software component passwords.
func (r *Hardware_SecurityModule) DeleteSoftwareComponentPasswords(softwareComponentPasswords []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		softwareComponentPasswords,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Edit a server's properties
func (r *Hardware_SecurityModule) EditObject(templateObject *datatypes.Hardware_Server) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Edit the properties of a software component password such as the username, password, and notes.
func (r *Hardware_SecurityModule) EditSoftwareComponentPasswords(softwareComponentPasswords []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		softwareComponentPasswords,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Download and run remote script from uri on the hardware.
func (r *Hardware_SecurityModule) ExecuteRemoteScript(uri *string) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		uri,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// The '''findByIpAddress''' method finds hardware using its primary public or private IP address. IP addresses that have a secondary subnet tied to the hardware will not return the hardware - alternate means of locating the hardware must be used (see '''Associated Methods'''). If no hardware is found, no errors are generated and no data is returned.
func (r *Hardware_SecurityModule) FindByIpAddress(ipAddress *string) (resp datatypes.Hardware, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
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
func (r *Hardware_SecurityModule) GenerateOrderTemplate(templateObject *datatypes.Hardware) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The account associated with a piece of hardware.
func (r *Hardware_SecurityModule) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's active physical components.
func (r *Hardware_SecurityModule) GetActiveComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The billing item for a server's attached network firewall.
func (r *Hardware_SecurityModule) GetActiveNetworkFirewallBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's active network monitoring incidents.
func (r *Hardware_SecurityModule) GetActiveNetworkMonitorIncident() (resp []datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_SecurityModule) GetActiveTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Transaction currently running for server.
func (r *Hardware_SecurityModule) GetActiveTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Any active transaction(s) that are currently running for the server (example: os reload).
func (r *Hardware_SecurityModule) GetActiveTransactions() (resp []datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getAlarmHistory''' method retrieves a detailed history for the monitoring alarm. When calling this method, a start and end date for the history to be retrieved must be entered.
func (r *Hardware_SecurityModule) GetAlarmHistory(startDate *datatypes.Time, endDate *datatypes.Time, alarmId *string) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
		alarmId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_SecurityModule) GetAllPowerComponents() (resp []datatypes.Hardware_Power_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Storage_Allowed_Host information to connect this server to Network Storage volumes that require access control lists.
func (r *Hardware_SecurityModule) GetAllowedHost() (resp datatypes.Network_Storage_Allowed_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Storage objects that this SoftLayer_Hardware has access to.
func (r *Hardware_SecurityModule) GetAllowedNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Storage objects whose Replica that this SoftLayer_Hardware has access to.
func (r *Hardware_SecurityModule) GetAllowedNetworkStorageReplicas() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding an antivirus/spyware software component object.
func (r *Hardware_SecurityModule) GetAntivirusSpywareSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This method is retrieve a list of SoftLayer_Network_Storage volumes that are authorized access to this SoftLayer_Hardware.
func (r *Hardware_SecurityModule) GetAttachedNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's specific attributes.
func (r *Hardware_SecurityModule) GetAttributes() (resp []datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An object that stores the maximum level for the monitoring query types and response types.
func (r *Hardware_SecurityModule) GetAvailableMonitoring() (resp []datatypes.Network_Monitor_Version1_Query_Host_Stratum, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This method retrieves a list of SoftLayer_Network_Storage volumes that can be authorized to this SoftLayer_Hardware.
func (r *Hardware_SecurityModule) GetAvailableNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The average daily total bandwidth usage for the current billing cycle.
func (r *Hardware_SecurityModule) GetAverageDailyBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The average daily private bandwidth usage for the current billing cycle.
func (r *Hardware_SecurityModule) GetAverageDailyPrivateBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The average daily public bandwidth usage for the current billing cycle.
func (r *Hardware_SecurityModule) GetAverageDailyPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Use this method to return an array of private bandwidth utilization records between a given date range.
//
// This method represents the NEW version of getFrontendBandwidthUse
func (r *Hardware_SecurityModule) GetBackendBandwidthUsage(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Use this method to return an array of private bandwidth utilization records between a given date range.
func (r *Hardware_SecurityModule) GetBackendBandwidthUse(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Network_Bandwidth_Version1_Usage_Detail, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// The '''getBackendIncomingBandwidth''' method retrieves the amount of incoming private network traffic used between the given start date and end date parameters. When entering start and end dates, only the month, day and year are used to calculate bandwidth totals - the time (HH:MM:SS) is ignored and defaults to midnight. The amount of bandwidth retrieved is measured in gigabytes.
func (r *Hardware_SecurityModule) GetBackendIncomingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's back-end or private network components.
func (r *Hardware_SecurityModule) GetBackendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getBackendOutgoingBandwidth''' method retrieves the amount of outgoing private network traffic used between the given start date and end date parameters. When entering start and end dates, only the month, day and year are used to calculate bandwidth totals - the time (HH:MM:SS) is ignored and defaults to midnight. The amount of bandwidth retrieved is measured in gigabytes.
func (r *Hardware_SecurityModule) GetBackendOutgoingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's backend or private router.
func (r *Hardware_SecurityModule) GetBackendRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's allotted bandwidth (measured in GB).
func (r *Hardware_SecurityModule) GetBandwidthAllocation() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's allotted detail record. Allotment details link bandwidth allocation with allotments.
func (r *Hardware_SecurityModule) GetBandwidthAllotmentDetail() (resp datatypes.Network_Bandwidth_Version1_Allotment_Detail, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a collection of bandwidth data from an individual public or private network tracking object. Data is ideal if you with to employ your own traffic storage and graphing systems.
func (r *Hardware_SecurityModule) GetBandwidthForDateRange(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Use this method when needing a bandwidth image for a single server.  It will gather the correct input parameters for the generic graphing utility automatically based on the snapshot specified.  Use the $draw flag to suppress the generation of the actual binary PNG image.
func (r *Hardware_SecurityModule) GetBandwidthImage(networkType *string, snapshotRange *string, draw *bool, dateSpecified *datatypes.Time, dateSpecifiedEnd *datatypes.Time) (resp datatypes.Container_Bandwidth_GraphOutputs, err error) {
	params := []interface{}{
		networkType,
		snapshotRange,
		draw,
		dateSpecified,
		dateSpecifiedEnd,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's benchmark certifications.
func (r *Hardware_SecurityModule) GetBenchmarkCertifications() (resp []datatypes.Hardware_Benchmark_Certification, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The raw bandwidth usage data for the current billing cycle. One object will be returned for each network this server is attached to.
func (r *Hardware_SecurityModule) GetBillingCycleBandwidthUsage() (resp []datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The raw private bandwidth usage data for the current billing cycle.
func (r *Hardware_SecurityModule) GetBillingCyclePrivateBandwidthUsage() (resp datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The raw public bandwidth usage data for the current billing cycle.
func (r *Hardware_SecurityModule) GetBillingCyclePublicBandwidthUsage() (resp datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the billing item for a server.
func (r *Hardware_SecurityModule) GetBillingItem() (resp datatypes.Billing_Item_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A flag indicating that a billing item exists.
func (r *Hardware_SecurityModule) GetBillingItemFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Determines whether the hardware is ineligible for cancellation because it is disconnected.
func (r *Hardware_SecurityModule) GetBlockCancelBecauseDisconnectedFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Status indicating whether or not a piece of hardware has business continuance insurance.
func (r *Hardware_SecurityModule) GetBusinessContinuanceInsuranceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's components.
func (r *Hardware_SecurityModule) GetComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_SecurityModule) GetContainsSolidStateDrivesFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A continuous data protection/server backup software component object.
func (r *Hardware_SecurityModule) GetContinuousDataProtectionSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A server's control panel.
func (r *Hardware_SecurityModule) GetControlPanel() (resp datatypes.Software_Component_ControlPanel, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total cost of a server, measured in US Dollars ($USD).
func (r *Hardware_SecurityModule) GetCost() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
// There are many options that may be provided while ordering a server, this method can be used to determine what these options are.
//
//
// Detailed information on the return value can be found on the data type page for [[SoftLayer_Container_Hardware_Configuration (type)]].
func (r *Hardware_SecurityModule) GetCreateObjectOptions() (resp datatypes.Container_Hardware_Configuration, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An object that provides commonly used bandwidth summary components for the current billing cycle.
func (r *Hardware_SecurityModule) GetCurrentBandwidthSummary() (resp datatypes.Metric_Tracking_Object_Bandwidth_Summary, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Attempt to retrieve the file associated with the current benchmark certification result, if such a file exists.  If there is no file for this benchmark certification result, calling this method throws an exception.
func (r *Hardware_SecurityModule) GetCurrentBenchmarkCertificationResultFile() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The current billable public outbound bandwidth for this hardware for the current billing cycle.
func (r *Hardware_SecurityModule) GetCurrentBillableBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Get the billing detail for this instance for the current billing period. This does not include bandwidth usage.
func (r *Hardware_SecurityModule) GetCurrentBillingDetail() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getCurrentBillingTotal''' method retrieves the total bill amount in US Dollars ($) for the current billing period. In addition to the total bill amount, the billing detail also includes all bandwidth used up to the point the method is called on the piece of hardware.
func (r *Hardware_SecurityModule) GetCurrentBillingTotal() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve bandwidth graph by date.
func (r *Hardware_SecurityModule) GetCustomBandwidthDataByDate(graphData *datatypes.Container_Graph) (resp datatypes.Container_Graph, err error) {
	params := []interface{}{
		graphData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Indicates if a server has a Customer Installed OS
func (r *Hardware_SecurityModule) GetCustomerInstalledOperatingSystemFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Indicates if a server is a customer owned device.
func (r *Hardware_SecurityModule) GetCustomerOwnedFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getDailyAverage''' method calculates the average daily network traffic used by the selected server. Using the required parameter ''dateTime'' to enter a start and end date, the user retrieves this average, measure in gigabytes (GB) for the specified date range. When entering parameters, only the month, day and year are required - time entries are omitted as this method defaults the time to midnight in order to account for the entire day.
func (r *Hardware_SecurityModule) GetDailyAverage(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the datacenter in which a piece of hardware resides.
func (r *Hardware_SecurityModule) GetDatacenter() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The name of the datacenter in which a piece of hardware resides.
func (r *Hardware_SecurityModule) GetDatacenterName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All hardware that has uplink network connections to a piece of hardware.
func (r *Hardware_SecurityModule) GetDownlinkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All hardware that has uplink network connections to a piece of hardware.
func (r *Hardware_SecurityModule) GetDownlinkNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all servers attached to a piece of network hardware.
func (r *Hardware_SecurityModule) GetDownlinkServers() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all virtual guests attached to a piece of network hardware.
func (r *Hardware_SecurityModule) GetDownlinkVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All hardware downstream from a network device.
func (r *Hardware_SecurityModule) GetDownstreamHardwareBindings() (resp []datatypes.Network_Component_Uplink_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All network hardware downstream from the selected piece of hardware.
func (r *Hardware_SecurityModule) GetDownstreamNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All network hardware with monitoring warnings or errors that are downstream from the selected piece of hardware.
func (r *Hardware_SecurityModule) GetDownstreamNetworkHardwareWithIncidents() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all servers attached downstream to a piece of network hardware.
func (r *Hardware_SecurityModule) GetDownstreamServers() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all virtual guests attached to a piece of network hardware.
func (r *Hardware_SecurityModule) GetDownstreamVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The drive controllers contained within a piece of hardware.
func (r *Hardware_SecurityModule) GetDriveControllers() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's associated EVault network storage service account.
func (r *Hardware_SecurityModule) GetEvaultNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Get the subnets associated with this server that are protectable by a network component firewall.
func (r *Hardware_SecurityModule) GetFirewallProtectableSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's firewall services.
func (r *Hardware_SecurityModule) GetFirewallServiceComponent() (resp datatypes.Network_Component_Firewall, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Defines the fixed components in a fixed configuration bare metal server.
func (r *Hardware_SecurityModule) GetFixedConfigurationPreset() (resp datatypes.Product_Package_Preset, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Use this method to return an array of public bandwidth utilization records between a given date range.
//
// This method represents the NEW version of getFrontendBandwidthUse
func (r *Hardware_SecurityModule) GetFrontendBandwidthUsage(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Use this method to return an array of public bandwidth utilization records between a given date range.
func (r *Hardware_SecurityModule) GetFrontendBandwidthUse(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Network_Bandwidth_Version1_Usage_Detail, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// The '''getFrontendIncomingBandwidth''' method retrieves the amount of incoming public network traffic used by a server between the given start and end date parameters. When entering the ''dateTime'' parameter, only the month, day and year of the start and end dates are required - the time (hour, minute and second) are set to midnight by default and cannot be changed. The amount of bandwidth retrieved is measured in gigabytes (GB).
func (r *Hardware_SecurityModule) GetFrontendIncomingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's front-end or public network components.
func (r *Hardware_SecurityModule) GetFrontendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getFrontendOutgoingBandwidth''' method retrieves the amount of outgoing public network traffic used by a server between the given start and end date parameters. The ''dateTime'' parameter requires only the day, month and year to be entered - the time (hour, minute and second) are set to midnight be default in order to gather the data for the entire start and end date indicated in the parameter. The amount of bandwidth retrieved is measured in gigabytes (GB).
func (r *Hardware_SecurityModule) GetFrontendOutgoingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's frontend or public router.
func (r *Hardware_SecurityModule) GetFrontendRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's universally unique identifier.
func (r *Hardware_SecurityModule) GetGlobalIdentifier() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hard drives contained within a piece of hardware.
func (r *Hardware_SecurityModule) GetHardDrives() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a server by searching for the primary IP address.
func (r *Hardware_SecurityModule) GetHardwareByIpAddress(ipAddress *string) (resp datatypes.Hardware_Server, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The chassis that a piece of hardware is housed in.
func (r *Hardware_SecurityModule) GetHardwareChassis() (resp datatypes.Hardware_Chassis, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's function.
func (r *Hardware_SecurityModule) GetHardwareFunction() (resp datatypes.Hardware_Function, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's function.
func (r *Hardware_SecurityModule) GetHardwareFunctionDescription() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's status.
func (r *Hardware_SecurityModule) GetHardwareStatus() (resp datatypes.Hardware_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Determine in hardware object has TPM enabled.
func (r *Hardware_SecurityModule) GetHasTrustedPlatformModuleBillingItemFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a host IPS software component object.
func (r *Hardware_SecurityModule) GetHostIpsSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getHourlyBandwidth''' method retrieves all bandwidth updates hourly for the specified hardware. Because the potential number of data points can become excessive, the method limits the user to obtain data in 24-hour intervals. The required ''dateTime'' parameter is used as the starting point for the query and will be calculated for the 24-hour period starting with the specified date and time. For example, entering a parameter of
//
// '02/01/2008 0:00'
//
// results in a return of all bandwidth data for the entire day of February 1, 2008, as 0:00 specifies a midnight start date. Please note that the time entered should be completed using a 24-hour clock (military time, astronomical time).
//
// For data spanning more than a single 24-hour period, refer to the getBandwidthData function on the metricTrackingObject for the piece of hardware.
func (r *Hardware_SecurityModule) GetHourlyBandwidth(mode *string, day *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		mode,
		day,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A server's hourly billing status.
func (r *Hardware_SecurityModule) GetHourlyBillingFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The sum of all the inbound network traffic data for the last 30 days.
func (r *Hardware_SecurityModule) GetInboundBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total private inbound bandwidth for this hardware for the current billing cycle.
func (r *Hardware_SecurityModule) GetInboundPrivateBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total public inbound bandwidth for this hardware for the current billing cycle.
func (r *Hardware_SecurityModule) GetInboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Return a collection of SoftLayer_Item_Price objects from a collection of SoftLayer_Software_Description
func (r *Hardware_SecurityModule) GetItemPricesFromSoftwareDescriptions(softwareDescriptions []datatypes.Software_Description, includeTranslationsFlag *bool, returnAllPricesFlag *bool) (resp []datatypes.Product_Item, err error) {
	params := []interface{}{
		softwareDescriptions,
		includeTranslationsFlag,
		returnAllPricesFlag,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The last transaction that a server's operating system was loaded.
func (r *Hardware_SecurityModule) GetLastOperatingSystemReload() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the last transaction a server performed.
func (r *Hardware_SecurityModule) GetLastTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's latest network monitoring incident.
func (r *Hardware_SecurityModule) GetLatestNetworkMonitorIncident() (resp datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Where a piece of hardware is located within SoftLayer's location hierarchy.
func (r *Hardware_SecurityModule) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_SecurityModule) GetLocationPathString() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a lockbox account associated with a server.
func (r *Hardware_SecurityModule) GetLockboxNetworkStorage() (resp datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A flag indicating that the hardware is a managed resource.
func (r *Hardware_SecurityModule) GetManagedResourceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the remote management network component attached with this server.
func (r *Hardware_SecurityModule) GetManagementNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's memory.
func (r *Hardware_SecurityModule) GetMemory() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The amount of memory a piece of hardware has, measured in gigabytes.
func (r *Hardware_SecurityModule) GetMemoryCapacity() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's metric tracking object.
func (r *Hardware_SecurityModule) GetMetricTrackingObject() (resp datatypes.Metric_Tracking_Object_HardwareServer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The metric tracking object id for this server.
func (r *Hardware_SecurityModule) GetMetricTrackingObjectId() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Returns open monitoring alarms for a given time period
func (r *Hardware_SecurityModule) GetMonitoringActiveAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the monitoring agents associated with a piece of hardware.
func (r *Hardware_SecurityModule) GetMonitoringAgents() (resp []datatypes.Monitoring_Agent, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Returns closed monitoring alarms for a given time period
func (r *Hardware_SecurityModule) GetMonitoringClosedAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the hardware's monitoring robot.
func (r *Hardware_SecurityModule) GetMonitoringRobot() (resp datatypes.Monitoring_Robot, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's network monitoring services.
func (r *Hardware_SecurityModule) GetMonitoringServiceComponent() (resp datatypes.Network_Monitor_Version1_Query_Host_Stratum, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The monitoring service flag eligibility status for a piece of hardware.
func (r *Hardware_SecurityModule) GetMonitoringServiceEligibilityFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The service flag status for a piece of hardware.
func (r *Hardware_SecurityModule) GetMonitoringServiceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The monitoring notification objects for this hardware. Each object links this hardware instance to a user account that will be notified if monitoring on this hardware object fails
func (r *Hardware_SecurityModule) GetMonitoringUserNotification() (resp []datatypes.User_Customer_Notification_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's motherboard.
func (r *Hardware_SecurityModule) GetMotherboard() (resp datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's network cards.
func (r *Hardware_SecurityModule) GetNetworkCards() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Get the IP addresses associated with this server that are protectable by a network component firewall. Note, this may not return all values for IPv6 subnets for this server. Please use getFirewallProtectableSubnets to get all protectable subnets.
func (r *Hardware_SecurityModule) GetNetworkComponentFirewallProtectableIpAddresses() (resp []datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Returns a hardware's network components.
func (r *Hardware_SecurityModule) GetNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The gateway member if this device is part of a network gateway.
func (r *Hardware_SecurityModule) GetNetworkGatewayMember() (resp datatypes.Network_Gateway_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Whether or not this device is part of a network gateway.
func (r *Hardware_SecurityModule) GetNetworkGatewayMemberFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's network management IP address.
func (r *Hardware_SecurityModule) GetNetworkManagementIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All servers with failed monitoring that are attached downstream to a piece of hardware.
func (r *Hardware_SecurityModule) GetNetworkMonitorAttachedDownHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Virtual guests that are attached downstream to a hardware that have failed monitoring
func (r *Hardware_SecurityModule) GetNetworkMonitorAttachedDownVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The status of all of a piece of hardware's network monitoring incidents.
func (r *Hardware_SecurityModule) GetNetworkMonitorIncidents() (resp []datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's network monitors.
func (r *Hardware_SecurityModule) GetNetworkMonitors() (resp []datatypes.Network_Monitor_Version1_Query_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The value of a hardware's network status attribute.
func (r *Hardware_SecurityModule) GetNetworkStatus() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware's related network status attribute.
func (r *Hardware_SecurityModule) GetNetworkStatusAttribute() (resp datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's associated network storage service account.
func (r *Hardware_SecurityModule) GetNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The network virtual LANs (VLANs) associated with a piece of hardware's network components.
func (r *Hardware_SecurityModule) GetNetworkVlans() (resp []datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's allotted bandwidth for the next billing cycle (measured in GB).
func (r *Hardware_SecurityModule) GetNextBillingCycleBandwidthAllocation() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_SecurityModule) GetNotesHistory() (resp []datatypes.Hardware_Note, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Hardware_SecurityModule) GetObject() (resp datatypes.Hardware_SecurityModule, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An open ticket requesting cancellation of this server, if one exists.
func (r *Hardware_SecurityModule) GetOpenCancellationTicket() (resp datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's operating system.
func (r *Hardware_SecurityModule) GetOperatingSystem() (resp datatypes.Software_Component_OperatingSystem, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's operating system software description.
func (r *Hardware_SecurityModule) GetOperatingSystemReferenceCode() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The sum of all the outbound network traffic data for the last 30 days.
func (r *Hardware_SecurityModule) GetOutboundBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total private outbound bandwidth for this hardware for the current billing cycle.
func (r *Hardware_SecurityModule) GetOutboundPrivateBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total public outbound bandwidth for this hardware for the current billing cycle.
func (r *Hardware_SecurityModule) GetOutboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Whether the bandwidth usage for this hardware for the current billing cycle exceeds the allocation.
func (r *Hardware_SecurityModule) GetOverBandwidthAllocationFlag() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a server's hardware state via its internal sensors. Remote sensor data is transmitted to the SoftLayer API by way of the server's remote management card. Sensor data measures system temperatures, voltages, and other local server settings. Sensor data is cached for 30 seconds. Calls made to getSensorData for the same server within 30 seconds of each other will return the same data. Subsequent calls will return new data once the cache expires.
func (r *Hardware_SecurityModule) GetPMInfo() (resp []datatypes.Container_RemoteManagement_PmInfo, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the Point of Presence (PoP) location in which a piece of hardware resides.
func (r *Hardware_SecurityModule) GetPointOfPresenceLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The power components for a hardware object.
func (r *Hardware_SecurityModule) GetPowerComponents() (resp []datatypes.Hardware_Power_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's power supply.
func (r *Hardware_SecurityModule) GetPowerSupply() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware's primary private IP address.
func (r *Hardware_SecurityModule) GetPrimaryBackendIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the hardware's primary back-end network component.
func (r *Hardware_SecurityModule) GetPrimaryBackendNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Hardware_SecurityModule) GetPrimaryDriveSize() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware's primary public IP address.
func (r *Hardware_SecurityModule) GetPrimaryIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the hardware's primary public network component.
func (r *Hardware_SecurityModule) GetPrimaryNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a graph of a server's private network bandwidth usage over the specified timeframe. If no timeframe is specified then getPublicBandwidthGraphImage retrieves the last 24 hours of public bandwidth usage. getPrivateBandwidthGraphImage returns a PNG image measuring 827 pixels by 293 pixels.
func (r *Hardware_SecurityModule) GetPrivateBandwidthData(startTime *int, endTime *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve a brief summary of a server's private network bandwidth usage. getPrivateBandwidthDataSummary retrieves a server's bandwidth allocation for its billing period, its estimated usage during its billing period, and an estimation of how much bandwidth it will use during its billing period based on its current usage. A server's projected bandwidth usage increases in accuracy as it progresses through its billing period.
func (r *Hardware_SecurityModule) GetPrivateBandwidthDataSummary() (resp datatypes.Container_Network_Bandwidth_Data_Summary, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a graph of a server's private network bandwidth usage over the specified time frame. If no time frame is specified then getPublicBandwidthGraphImage retrieves the last 24 hours of public bandwidth usage. getPublicBandwidthGraphImage returns a PNG image
func (r *Hardware_SecurityModule) GetPrivateBandwidthGraphImage(startTime *string, endTime *string) (resp []byte, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A server's primary private IP address.
func (r *Hardware_SecurityModule) GetPrivateIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the private network component attached with this server.
func (r *Hardware_SecurityModule) GetPrivateNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Whether the hardware only has access to the private network.
func (r *Hardware_SecurityModule) GetPrivateNetworkOnlyFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the backend VLAN for the primary IP address of the server
func (r *Hardware_SecurityModule) GetPrivateVlan() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a backend network VLAN by searching for an IP address
func (r *Hardware_SecurityModule) GetPrivateVlanByIpAddress(ipAddress *string) (resp datatypes.Network_Vlan, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total number of processor cores, summed from all processors that are attached to a piece of hardware
func (r *Hardware_SecurityModule) GetProcessorCoreAmount() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total number of physical processor cores, summed from all processors that are attached to a piece of hardware
func (r *Hardware_SecurityModule) GetProcessorPhysicalCoreAmount() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's processors.
func (r *Hardware_SecurityModule) GetProcessors() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Whether the bandwidth usage for this hardware for the current billing cycle is projected to exceed the allocation.
func (r *Hardware_SecurityModule) GetProjectedOverBandwidthAllocationFlag() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The projected public outbound bandwidth for this hardware for the current billing cycle.
func (r *Hardware_SecurityModule) GetProjectedPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Hardware_SecurityModule) GetProvisionDate() (resp datatypes.Time, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a graph of a server's public network bandwidth usage over the specified timeframe. If no timeframe is specified then getPublicBandwidthGraphImage retrieves the last 24 hours of public bandwidth usage. getPublicBandwidthGraphImage returns a PNG image measuring 827 pixels by 293 pixels.
func (r *Hardware_SecurityModule) GetPublicBandwidthData(startTime *int, endTime *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve a brief summary of a server's public network bandwidth usage. getPublicBandwidthDataSummary retrieves a server's bandwidth allocation for its billing period, its estimated usage during its billing period, and an estimation of how much bandwidth it will use during its billing period based on its current usage. A server's projected bandwidth usage increases in accuracy as it progresses through its billing period.
func (r *Hardware_SecurityModule) GetPublicBandwidthDataSummary() (resp datatypes.Container_Network_Bandwidth_Data_Summary, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a graph of a server's public network bandwidth usage over the specified time frame. If no time frame is specified then getPublicBandwidthGraphImage retrieves the last 24 hours of public bandwidth usage. getPublicBandwidthGraphImage returns a PNG image measuring 827 pixels by 293 pixels.  THIS METHOD GENERATES GRAPHS BASED ON THE NEW DATA WAREHOUSE REPOSITORY.
func (r *Hardware_SecurityModule) GetPublicBandwidthGraphImage(startTime *datatypes.Time, endTime *datatypes.Time) (resp []byte, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve the total number of bytes used by a server over a specified time period via the data warehouse tracking objects for this hardware.
func (r *Hardware_SecurityModule) GetPublicBandwidthTotal(startTime *int, endTime *int) (resp uint, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve a SoftLayer server's public network component. Some servers are only connected to the private network and may not have a public network component. In that case getPublicNetworkComponent returns a null object.
func (r *Hardware_SecurityModule) GetPublicNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the frontend VLAN for the primary IP address of the server
func (r *Hardware_SecurityModule) GetPublicVlan() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the frontend network Vlan by searching the hostname of a server
func (r *Hardware_SecurityModule) GetPublicVlanByHostname(hostname *string) (resp datatypes.Network_Vlan, err error) {
	params := []interface{}{
		hostname,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_SecurityModule) GetRack() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The RAID controllers contained within a piece of hardware.
func (r *Hardware_SecurityModule) GetRaidControllers() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Recent events that impact this hardware.
func (r *Hardware_SecurityModule) GetRecentEvents() (resp []datatypes.Notification_Occurrence_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The last five commands issued to the server's remote management card.
func (r *Hardware_SecurityModule) GetRecentRemoteManagementCommands() (resp []datatypes.Hardware_Component_RemoteManagement_Command_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_SecurityModule) GetRegionalInternetRegistry() (resp datatypes.Network_Regional_Internet_Registry, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A server's remote management card.
func (r *Hardware_SecurityModule) GetRemoteManagement() (resp datatypes.Hardware_Component_RemoteManagement, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve User credentials to issue commands and/or interact with the server's remote management card.
func (r *Hardware_SecurityModule) GetRemoteManagementAccounts() (resp []datatypes.Hardware_Component_RemoteManagement_User, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's associated remote management component. This is normally IPMI.
func (r *Hardware_SecurityModule) GetRemoteManagementComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve User(s) who have access to issue commands and/or interact with the server's remote management card.
func (r *Hardware_SecurityModule) GetRemoteManagementUsers() (resp []datatypes.Hardware_Component_RemoteManagement_User, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_SecurityModule) GetResourceGroupMemberReferences() (resp []datatypes.Resource_Group_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_SecurityModule) GetResourceGroupRoles() (resp []datatypes.Resource_Group_Role, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The resource groups in which this hardware is a member.
func (r *Hardware_SecurityModule) GetResourceGroups() (resp []datatypes.Resource_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the reverse domain records associated with this server.
func (r *Hardware_SecurityModule) GetReverseDomainRecords() (resp []datatypes.Dns_Domain, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's routers.
func (r *Hardware_SecurityModule) GetRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Collection of scale assets this hardware corresponds to.
func (r *Hardware_SecurityModule) GetScaleAssets() (resp []datatypes.Scale_Asset, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's vulnerability scan requests.
func (r *Hardware_SecurityModule) GetSecurityScanRequests() (resp []datatypes.Network_Security_Scanner_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a server's hardware state via its internal sensors. Remote sensor data is transmitted to the SoftLayer API by way of the server's remote management card. Sensor data measures system temperatures, voltages, and other local server settings. Sensor data is cached for 30 seconds. Calls made to getSensorData for the same server within 30 seconds of each other will return the same data. Subsequent calls will return new data once the cache expires.
func (r *Hardware_SecurityModule) GetSensorData() (resp []datatypes.Container_RemoteManagement_SensorReading, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieves the raw data returned from the server's remote management card.  For more details of what is returned please refer to the getSensorData method.  Along with the raw data, graphs for the cpu and system temperatures and fan speeds are also returned.
func (r *Hardware_SecurityModule) GetSensorDataWithGraphs() (resp datatypes.Container_RemoteManagement_SensorReadingsWithGraphs, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a server's hardware components, software, and network components. getServerDetails is an aggregation function that combines the results of [[SoftLayer_Hardware_Server::getComponents]], [[SoftLayer_Hardware_Server::getSoftware]], and [[SoftLayer_Hardware_Server::getNetworkComponents]] in a single container.
func (r *Hardware_SecurityModule) GetServerDetails() (resp datatypes.Container_Hardware_Server_Details, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the server's fan speeds and displays them using tachometer graphs.  Data used to construct graphs is retrieved from the server's remote management card.  All graphs returned will have a title associated with it.
func (r *Hardware_SecurityModule) GetServerFanSpeedGraphs() (resp []datatypes.Container_RemoteManagement_Graphs_SensorSpeed, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieves the power state for the server.  The server's power status is retrieved from its remote management card.  This will return 'on' or 'off'.
func (r *Hardware_SecurityModule) GetServerPowerState() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the server room in which the hardware is located.
func (r *Hardware_SecurityModule) GetServerRoom() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the server's temperature and displays them using thermometer graphs.  Temperatures retrieved are CPU(s) and system temperatures.  Data used to construct graphs is retrieved from the server's remote management card.  All graphs returned will have a title associated with it.
func (r *Hardware_SecurityModule) GetServerTemperatureGraphs() (resp []datatypes.Container_RemoteManagement_Graphs_SensorTemperature, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the piece of hardware's service provider.
func (r *Hardware_SecurityModule) GetServiceProvider() (resp datatypes.Service_Provider, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's installed software.
func (r *Hardware_SecurityModule) GetSoftwareComponents() (resp []datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the billing item for a spare pool server.
func (r *Hardware_SecurityModule) GetSparePoolBillingItem() (resp datatypes.Billing_Item_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve SSH keys to be installed on the server during provisioning or an OS reload.
func (r *Hardware_SecurityModule) GetSshKeys() (resp []datatypes.Security_Ssh_Key, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A server's remote management card used for statistics.
func (r *Hardware_SecurityModule) GetStatisticsRemoteManagement() (resp datatypes.Hardware_Component_RemoteManagement, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_SecurityModule) GetStorageNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_SecurityModule) GetTagReferences() (resp []datatypes.Tag_Reference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_SecurityModule) GetTopLevelLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
// This method will query transaction history for a piece of hardware.
func (r *Hardware_SecurityModule) GetTransactionHistory() (resp []datatypes.Provisioning_Version1_Transaction_History, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a list of upgradeable items available to this piece of hardware. Currently, getUpgradeItemPrices retrieves upgrades available for a server's memory, hard drives, network port speed, bandwidth allocation and GPUs.
func (r *Hardware_SecurityModule) GetUpgradeItemPrices() (resp []datatypes.Product_Item_Price, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An account's associated upgrade request object, if any.
func (r *Hardware_SecurityModule) GetUpgradeRequest() (resp datatypes.Product_Upgrade_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The network device connected to a piece of hardware.
func (r *Hardware_SecurityModule) GetUplinkHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the network component that is one level higher than a piece of hardware on the network infrastructure.
func (r *Hardware_SecurityModule) GetUplinkNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A string containing custom user data for a hardware order.
func (r *Hardware_SecurityModule) GetUserData() (resp []datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A list of users that have access to this computing instance.
func (r *Hardware_SecurityModule) GetUsers() (resp []datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This method will return the list of block device template groups that are valid to the host. For instance, it will only retrieve FLEX images.
func (r *Hardware_SecurityModule) GetValidBlockDeviceTemplateGroups(visibility *string) (resp []datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	params := []interface{}{
		visibility,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the virtual chassis for a piece of hardware.
func (r *Hardware_SecurityModule) GetVirtualChassis() (resp datatypes.Hardware_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the virtual chassis siblings for a piece of hardware.
func (r *Hardware_SecurityModule) GetVirtualChassisSiblings() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware server's virtual servers.
func (r *Hardware_SecurityModule) GetVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's virtual host record.
func (r *Hardware_SecurityModule) GetVirtualHost() (resp datatypes.Virtual_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's virtual software licenses.
func (r *Hardware_SecurityModule) GetVirtualLicenses() (resp []datatypes.Software_VirtualLicense, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the bandwidth allotment to which a piece of hardware belongs.
func (r *Hardware_SecurityModule) GetVirtualRack() (resp datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The name of the bandwidth allotment belonging to a piece of hardware.
func (r *Hardware_SecurityModule) GetVirtualRackId() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The name of the bandwidth allotment belonging to a piece of hardware.
func (r *Hardware_SecurityModule) GetVirtualRackName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's virtualization platform software.
func (r *Hardware_SecurityModule) GetVirtualizationPlatform() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a list of Windows updates available for a server from the local SoftLayer Windows Server Update Services (WSUS) server. Windows servers provisioned by SoftLayer are configured to use the local WSUS server via the private network by default.
func (r *Hardware_SecurityModule) GetWindowsUpdateAvailableUpdates() (resp []datatypes.Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a list of Windows updates installed on a server as reported by the local SoftLayer Windows Server Update Services (WSUS) server. Windows servers provisioned by SoftLayer are configured to use the local WSUS server via the private network by default.
func (r *Hardware_SecurityModule) GetWindowsUpdateInstalledUpdates() (resp []datatypes.Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This method returns an update status record for this server.  That record will specify if the server is missing updates, or has updates that must be reinstalled or require a reboot to go into affect.
func (r *Hardware_SecurityModule) GetWindowsUpdateStatus() (resp datatypes.Container_Utility_Microsoft_Windows_UpdateServices_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''importVirtualHost''' method attempts to import the host record for the virtualization platform running on a server.
func (r *Hardware_SecurityModule) ImportVirtualHost() (resp datatypes.Virtual_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Idera Bare Metal Server Restore is a backup agent designed specifically for making full system restores made with Idera Server Backup.
func (r *Hardware_SecurityModule) InitiateIderaBareMetalRestore() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// R1Soft Bare Metal Server Restore is an R1Soft disk agent designed specifically for making full system restores made with R1Soft CDP Server backup.
func (r *Hardware_SecurityModule) InitiateR1SoftBareMetalRestore() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Issues a ping command and returns the success (true) or failure (false) of the ping command.
func (r *Hardware_SecurityModule) IsBackendPingable() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Issues a ping command and returns the success (true) or failure (false) of the ping command.
func (r *Hardware_SecurityModule) IsPingable() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Determine if the server runs any version of the Microsoft Windows operating systems. Return ''true'' if it does and ''false if otherwise.
func (r *Hardware_SecurityModule) IsWindowsServer() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Issues a ping command to the server and returns the ping response.
func (r *Hardware_SecurityModule) Ping() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Power off then power on the server via powerstrip.  The power cycle command is equivalent to unplugging the server from the powerstrip and then plugging the server back into the powerstrip.  This should only be used as a last resort.  If a reboot command has been issued successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard, powerOn, powerOff and powerCycle) will not be allowed. This is to avoid any type of server failures.
func (r *Hardware_SecurityModule) PowerCycle() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This method will power off the server via the server's remote management card.
func (r *Hardware_SecurityModule) PowerOff() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Power on server via its remote management card.  If a reboot command has been issued successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard, powerOn, powerOff and powerCycle) will not be allowed.  This is to avoid any type of server failures.
func (r *Hardware_SecurityModule) PowerOn() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Attempts to reboot the server by issuing a reset (soft reboot) command to the server's remote management card. If the reset (soft reboot) attempt is unsuccessful, a power cycle command will be issued via the powerstrip. The power cycle command is equivalent to unplugging the server from the powerstrip and then plugging the server back into the powerstrip.  If a reboot command has been issued successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard, powerOn, powerOff and powerCycle) will not be allowed.  This is to avoid any type of server failures.
func (r *Hardware_SecurityModule) RebootDefault() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Reboot the server by issuing a cycle command to the server's remote management card.  This is equivalent to pressing the 'Reset' button on the server.  This command is issued immediately and will not wait for processes to shutdown. After this command is issued, the server may take a few moments to boot up as server may run system disks checks. If a reboot command has been issued successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard, powerOn, powerOff and powerCycle) will not be allowed.  This is to avoid any type of server failures.
func (r *Hardware_SecurityModule) RebootHard() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Reboot the server by issuing a reset command to the server's remote management card.  This is a graceful reboot. The servers will allow all process to shutdown gracefully before rebooting.  If a reboot command has been issued successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard, powerOn, powerOff and powerCycle) will not be allowed.  This is to avoid any type of server failures.
func (r *Hardware_SecurityModule) RebootSoft() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Reloads current operating system configuration.
//
// This service has a confirmation protocol for proceeding with the reload. To proceed with the reload without confirmation, simply pass in 'FORCE' as the token parameter. To proceed with the reload with confirmation, simply call the service with no parameter. A token string will be returned by this service. The token will remain active for 10 minutes. Use this token as the parameter to confirm that a reload is to be performed for the server.
//
// As a precaution, we strongly  recommend backing up all data before reloading the operating system. The reload will format the primary disk and will reconfigure the server to the current specifications on record.
//
// The reload will take AT MINIMUM 66 minutes.
func (r *Hardware_SecurityModule) ReloadCurrentOperatingSystemConfiguration(token *string) (resp string, err error) {
	params := []interface{}{
		token,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Reloads current or customer specified operating system configuration.
//
// This service has a confirmation protocol for proceeding with the reload. To proceed with the reload without confirmation, simply pass in 'FORCE' as the token parameter. To proceed with the reload with confirmation, simply call the service with no parameter. A token string will be returned by this service. The token will remain active for 10 minutes. Use this token as the parameter to confirm that a reload is to be performed for the server.
//
// As a precaution, we strongly  recommend backing up all data before reloading the operating system. The reload will format the primary disk and will reconfigure the server to the current specifications on record.
//
// The reload will take AT MINIMUM 66 minutes.
func (r *Hardware_SecurityModule) ReloadOperatingSystem(token *string, config *datatypes.Container_Hardware_Server_Configuration) (resp string, err error) {
	params := []interface{}{
		token,
		config,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// This method is used to remove access to s SoftLayer_Network_Storage volumes that supports host- or network-level access control.
func (r *Hardware_SecurityModule) RemoveAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// This method is used to allow access to multiple SoftLayer_Network_Storage volumes that support host- or network-level access control.
func (r *Hardware_SecurityModule) RemoveAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// You can launch a new Passmark hardware test by selecting from your server list. It will bring your server offline for approximately 20 minutes while the testing is in progress, and will publish a certificate with the results to your hardware details page.
//
// While the hard drives are tested for the initial deployment, the Passmark Certificate utility will not test the hard drives on your live server. This is to ensure that no data is overwritten. If you would like to test the server's hard drives, you can have the full Passmark suite installed to your server free of charge through a new Support ticket.
//
// While the test itself does not overwrite any data on the server, it is recommended that you make full off-server backups of all data prior to launching the test. The Passmark hardware test is designed to force any latent hardware issues to the surface, so hardware failure is possible.
//
// In the event of a hardware failure during this test our datacenter engineers will be notified of the problem automatically. They will then replace any failed components to bring your server back online, and will be contacting you to ensure that impact on your server is minimal.
func (r *Hardware_SecurityModule) RunPassmarkCertificationBenchmark() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Changes the password that we have stored in our database for a servers' Operating System
func (r *Hardware_SecurityModule) SetOperatingSystemPassword(newPassword *string) (resp bool, err error) {
	params := []interface{}{
		newPassword,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Sets the private network interface speed to the new speed. Speed values can only be 0 (Disconnect), 10, 100, 1000, and 10000. The new speed must be equal to or less than the max speed of the interface.
//
// It will take less than a minute to update the switch port speed. The server uplink will not be operational again until the server interface speed is updated.
func (r *Hardware_SecurityModule) SetPrivateNetworkInterfaceSpeed(newSpeed *int) (resp bool, err error) {
	params := []interface{}{
		newSpeed,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Sets the public network interface speed to the new speed. Speed values can only be 0 (Disconnect), 10, 100, 1000, and 10000. The new speed must be equal to or less than the max speed of the interface.
//
// It will take less than a minute to update the switch port speed. The server uplink will not be operational again until the server interface speed is updated.
func (r *Hardware_SecurityModule) SetPublicNetworkInterfaceSpeed(newSpeed *int) (resp bool, err error) {
	params := []interface{}{
		newSpeed,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Hardware_SecurityModule) SetTags(tags *string) (resp bool, err error) {
	params := []interface{}{
		tags,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Sets the data that will be written to the configuration drive.
func (r *Hardware_SecurityModule) SetUserMetadata(metadata []string) (resp []datatypes.Hardware_Attribute, err error) {
	params := []interface{}{
		metadata,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Shuts down the public network port
func (r *Hardware_SecurityModule) ShutdownPrivatePort() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Shuts down the public network port
func (r *Hardware_SecurityModule) ShutdownPublicPort() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The ability to place bare metal servers in a state where they are powered down, and ports closed yet still allocated to the customer as a part of the Spare Pool program.
func (r *Hardware_SecurityModule) SparePool(action *string, newOrder *bool) (resp bool, err error) {
	params := []interface{}{
		action,
		newOrder,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Validates a collection of partitions for an operating system
func (r *Hardware_SecurityModule) ValidatePartitionsForOperatingSystem(operatingSystem *datatypes.Software_Description, partitions []datatypes.Hardware_Component_Partition) (resp bool, err error) {
	params := []interface{}{
		operatingSystem,
		partitions,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// The SoftLayer_Hardware_Server data type contains general information relating to a single SoftLayer server.
type Hardware_Server struct {
	Session *session.Session
	sl.Options
}

func GetHardwareServerService(sess *session.Session) Hardware_Server {
	return Hardware_Server{Session: sess}
}

// Activates the private network port
func (r *Hardware_Server) ActivatePrivatePort() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Activates the public network port
func (r *Hardware_Server) ActivatePublicPort() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This method is used to allow access to a SoftLayer_Network_Storage volume that supports host- or network-level access control.
func (r *Hardware_Server) AllowAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// This method is used to allow access to multiple SoftLayer_Network_Storage volumes that support host- or network-level access control.
func (r *Hardware_Server) AllowAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// The Rescue Kernel is designed to provide you with the ability to bring a server online in order to troubleshoot system problems that would normally only be resolved by an OS Reload. The correct Rescue Kernel will be selected based upon the currently installed operating system. When the rescue kernel process is initiated, the server will shutdown and reboot on to the public network with the same IP's assigned to the server to allow for remote connections. It will bring your server offline for approximately 10 minutes while the rescue is in progress. The root/administrator password will be the same as what is listed in the portal for the server.
func (r *Hardware_Server) BootToRescueLayer(noOsBootEnvironment *string) (resp bool, err error) {
	params := []interface{}{
		noOsBootEnvironment,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Captures a Flex Image of the hard disk on the physical machine, based on the capture template parameter. Returns the image template group containing the disk image.
func (r *Hardware_Server) CaptureImage(captureTemplate *datatypes.Container_Disk_Image_Capture_Template) (resp datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	params := []interface{}{
		captureTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Returns monitoring alarm detailed history
func (r *Hardware_Server) CloseAlarm(alarmId *string) (resp bool, err error) {
	params := []interface{}{
		alarmId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// You can launch firmware updates by selecting from your server list. It will bring your server offline for approximately 20 minutes while the updates are in progress.
//
// In the event of a hardware failure during this test our datacenter engineers will be notified of the problem automatically. They will then replace any failed components to bring your server back online, and will be contacting you to ensure that impact on your server is minimal.
func (r *Hardware_Server) CreateFirmwareUpdateTransaction(ipmi *int, raidController *int, bios *int, harddrive *int) (resp bool, err error) {
	params := []interface{}{
		ipmi,
		raidController,
		bios,
		harddrive,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
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
func (r *Hardware_Server) CreateObject(templateObject *datatypes.Hardware_Server) (resp datatypes.Hardware_Server, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Hardware_Server) CreatePostSoftwareInstallTransaction(installCodes []string, returnBoolean *bool) (resp bool, err error) {
	params := []interface{}{
		installCodes,
		returnBoolean,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
// This method will cancel a server effective immediately. For servers billed hourly, the charges will stop immediately after the method returns.
func (r *Hardware_Server) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Delete software component passwords.
func (r *Hardware_Server) DeleteSoftwareComponentPasswords(softwareComponentPasswords []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		softwareComponentPasswords,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Edit a server's properties
func (r *Hardware_Server) EditObject(templateObject *datatypes.Hardware_Server) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Edit the properties of a software component password such as the username, password, and notes.
func (r *Hardware_Server) EditSoftwareComponentPasswords(softwareComponentPasswords []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		softwareComponentPasswords,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Download and run remote script from uri on the hardware.
func (r *Hardware_Server) ExecuteRemoteScript(uri *string) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		uri,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// The '''findByIpAddress''' method finds hardware using its primary public or private IP address. IP addresses that have a secondary subnet tied to the hardware will not return the hardware - alternate means of locating the hardware must be used (see '''Associated Methods'''). If no hardware is found, no errors are generated and no data is returned.
func (r *Hardware_Server) FindByIpAddress(ipAddress *string) (resp datatypes.Hardware, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
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
func (r *Hardware_Server) GenerateOrderTemplate(templateObject *datatypes.Hardware) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The account associated with a piece of hardware.
func (r *Hardware_Server) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's active physical components.
func (r *Hardware_Server) GetActiveComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The billing item for a server's attached network firewall.
func (r *Hardware_Server) GetActiveNetworkFirewallBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's active network monitoring incidents.
func (r *Hardware_Server) GetActiveNetworkMonitorIncident() (resp []datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_Server) GetActiveTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Transaction currently running for server.
func (r *Hardware_Server) GetActiveTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Any active transaction(s) that are currently running for the server (example: os reload).
func (r *Hardware_Server) GetActiveTransactions() (resp []datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getAlarmHistory''' method retrieves a detailed history for the monitoring alarm. When calling this method, a start and end date for the history to be retrieved must be entered.
func (r *Hardware_Server) GetAlarmHistory(startDate *datatypes.Time, endDate *datatypes.Time, alarmId *string) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
		alarmId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_Server) GetAllPowerComponents() (resp []datatypes.Hardware_Power_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Storage_Allowed_Host information to connect this server to Network Storage volumes that require access control lists.
func (r *Hardware_Server) GetAllowedHost() (resp datatypes.Network_Storage_Allowed_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Storage objects that this SoftLayer_Hardware has access to.
func (r *Hardware_Server) GetAllowedNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Storage objects whose Replica that this SoftLayer_Hardware has access to.
func (r *Hardware_Server) GetAllowedNetworkStorageReplicas() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding an antivirus/spyware software component object.
func (r *Hardware_Server) GetAntivirusSpywareSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This method is retrieve a list of SoftLayer_Network_Storage volumes that are authorized access to this SoftLayer_Hardware.
func (r *Hardware_Server) GetAttachedNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's specific attributes.
func (r *Hardware_Server) GetAttributes() (resp []datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An object that stores the maximum level for the monitoring query types and response types.
func (r *Hardware_Server) GetAvailableMonitoring() (resp []datatypes.Network_Monitor_Version1_Query_Host_Stratum, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This method retrieves a list of SoftLayer_Network_Storage volumes that can be authorized to this SoftLayer_Hardware.
func (r *Hardware_Server) GetAvailableNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The average daily total bandwidth usage for the current billing cycle.
func (r *Hardware_Server) GetAverageDailyBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The average daily private bandwidth usage for the current billing cycle.
func (r *Hardware_Server) GetAverageDailyPrivateBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The average daily public bandwidth usage for the current billing cycle.
func (r *Hardware_Server) GetAverageDailyPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Use this method to return an array of private bandwidth utilization records between a given date range.
//
// This method represents the NEW version of getFrontendBandwidthUse
func (r *Hardware_Server) GetBackendBandwidthUsage(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Use this method to return an array of private bandwidth utilization records between a given date range.
func (r *Hardware_Server) GetBackendBandwidthUse(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Network_Bandwidth_Version1_Usage_Detail, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// The '''getBackendIncomingBandwidth''' method retrieves the amount of incoming private network traffic used between the given start date and end date parameters. When entering start and end dates, only the month, day and year are used to calculate bandwidth totals - the time (HH:MM:SS) is ignored and defaults to midnight. The amount of bandwidth retrieved is measured in gigabytes.
func (r *Hardware_Server) GetBackendIncomingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's back-end or private network components.
func (r *Hardware_Server) GetBackendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getBackendOutgoingBandwidth''' method retrieves the amount of outgoing private network traffic used between the given start date and end date parameters. When entering start and end dates, only the month, day and year are used to calculate bandwidth totals - the time (HH:MM:SS) is ignored and defaults to midnight. The amount of bandwidth retrieved is measured in gigabytes.
func (r *Hardware_Server) GetBackendOutgoingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's backend or private router.
func (r *Hardware_Server) GetBackendRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's allotted bandwidth (measured in GB).
func (r *Hardware_Server) GetBandwidthAllocation() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's allotted detail record. Allotment details link bandwidth allocation with allotments.
func (r *Hardware_Server) GetBandwidthAllotmentDetail() (resp datatypes.Network_Bandwidth_Version1_Allotment_Detail, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a collection of bandwidth data from an individual public or private network tracking object. Data is ideal if you with to employ your own traffic storage and graphing systems.
func (r *Hardware_Server) GetBandwidthForDateRange(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Use this method when needing a bandwidth image for a single server.  It will gather the correct input parameters for the generic graphing utility automatically based on the snapshot specified.  Use the $draw flag to suppress the generation of the actual binary PNG image.
func (r *Hardware_Server) GetBandwidthImage(networkType *string, snapshotRange *string, draw *bool, dateSpecified *datatypes.Time, dateSpecifiedEnd *datatypes.Time) (resp datatypes.Container_Bandwidth_GraphOutputs, err error) {
	params := []interface{}{
		networkType,
		snapshotRange,
		draw,
		dateSpecified,
		dateSpecifiedEnd,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's benchmark certifications.
func (r *Hardware_Server) GetBenchmarkCertifications() (resp []datatypes.Hardware_Benchmark_Certification, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The raw bandwidth usage data for the current billing cycle. One object will be returned for each network this server is attached to.
func (r *Hardware_Server) GetBillingCycleBandwidthUsage() (resp []datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The raw private bandwidth usage data for the current billing cycle.
func (r *Hardware_Server) GetBillingCyclePrivateBandwidthUsage() (resp datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The raw public bandwidth usage data for the current billing cycle.
func (r *Hardware_Server) GetBillingCyclePublicBandwidthUsage() (resp datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the billing item for a server.
func (r *Hardware_Server) GetBillingItem() (resp datatypes.Billing_Item_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A flag indicating that a billing item exists.
func (r *Hardware_Server) GetBillingItemFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Determines whether the hardware is ineligible for cancellation because it is disconnected.
func (r *Hardware_Server) GetBlockCancelBecauseDisconnectedFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Status indicating whether or not a piece of hardware has business continuance insurance.
func (r *Hardware_Server) GetBusinessContinuanceInsuranceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's components.
func (r *Hardware_Server) GetComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_Server) GetContainsSolidStateDrivesFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A continuous data protection/server backup software component object.
func (r *Hardware_Server) GetContinuousDataProtectionSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A server's control panel.
func (r *Hardware_Server) GetControlPanel() (resp datatypes.Software_Component_ControlPanel, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total cost of a server, measured in US Dollars ($USD).
func (r *Hardware_Server) GetCost() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
// There are many options that may be provided while ordering a server, this method can be used to determine what these options are.
//
//
// Detailed information on the return value can be found on the data type page for [[SoftLayer_Container_Hardware_Configuration (type)]].
func (r *Hardware_Server) GetCreateObjectOptions() (resp datatypes.Container_Hardware_Configuration, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An object that provides commonly used bandwidth summary components for the current billing cycle.
func (r *Hardware_Server) GetCurrentBandwidthSummary() (resp datatypes.Metric_Tracking_Object_Bandwidth_Summary, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Attempt to retrieve the file associated with the current benchmark certification result, if such a file exists.  If there is no file for this benchmark certification result, calling this method throws an exception.
func (r *Hardware_Server) GetCurrentBenchmarkCertificationResultFile() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The current billable public outbound bandwidth for this hardware for the current billing cycle.
func (r *Hardware_Server) GetCurrentBillableBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Get the billing detail for this instance for the current billing period. This does not include bandwidth usage.
func (r *Hardware_Server) GetCurrentBillingDetail() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getCurrentBillingTotal''' method retrieves the total bill amount in US Dollars ($) for the current billing period. In addition to the total bill amount, the billing detail also includes all bandwidth used up to the point the method is called on the piece of hardware.
func (r *Hardware_Server) GetCurrentBillingTotal() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve bandwidth graph by date.
func (r *Hardware_Server) GetCustomBandwidthDataByDate(graphData *datatypes.Container_Graph) (resp datatypes.Container_Graph, err error) {
	params := []interface{}{
		graphData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Indicates if a server has a Customer Installed OS
func (r *Hardware_Server) GetCustomerInstalledOperatingSystemFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Indicates if a server is a customer owned device.
func (r *Hardware_Server) GetCustomerOwnedFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getDailyAverage''' method calculates the average daily network traffic used by the selected server. Using the required parameter ''dateTime'' to enter a start and end date, the user retrieves this average, measure in gigabytes (GB) for the specified date range. When entering parameters, only the month, day and year are required - time entries are omitted as this method defaults the time to midnight in order to account for the entire day.
func (r *Hardware_Server) GetDailyAverage(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the datacenter in which a piece of hardware resides.
func (r *Hardware_Server) GetDatacenter() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The name of the datacenter in which a piece of hardware resides.
func (r *Hardware_Server) GetDatacenterName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All hardware that has uplink network connections to a piece of hardware.
func (r *Hardware_Server) GetDownlinkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All hardware that has uplink network connections to a piece of hardware.
func (r *Hardware_Server) GetDownlinkNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all servers attached to a piece of network hardware.
func (r *Hardware_Server) GetDownlinkServers() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all virtual guests attached to a piece of network hardware.
func (r *Hardware_Server) GetDownlinkVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All hardware downstream from a network device.
func (r *Hardware_Server) GetDownstreamHardwareBindings() (resp []datatypes.Network_Component_Uplink_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All network hardware downstream from the selected piece of hardware.
func (r *Hardware_Server) GetDownstreamNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All network hardware with monitoring warnings or errors that are downstream from the selected piece of hardware.
func (r *Hardware_Server) GetDownstreamNetworkHardwareWithIncidents() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all servers attached downstream to a piece of network hardware.
func (r *Hardware_Server) GetDownstreamServers() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding all virtual guests attached to a piece of network hardware.
func (r *Hardware_Server) GetDownstreamVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The drive controllers contained within a piece of hardware.
func (r *Hardware_Server) GetDriveControllers() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's associated EVault network storage service account.
func (r *Hardware_Server) GetEvaultNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Get the subnets associated with this server that are protectable by a network component firewall.
func (r *Hardware_Server) GetFirewallProtectableSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's firewall services.
func (r *Hardware_Server) GetFirewallServiceComponent() (resp datatypes.Network_Component_Firewall, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Defines the fixed components in a fixed configuration bare metal server.
func (r *Hardware_Server) GetFixedConfigurationPreset() (resp datatypes.Product_Package_Preset, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Use this method to return an array of public bandwidth utilization records between a given date range.
//
// This method represents the NEW version of getFrontendBandwidthUse
func (r *Hardware_Server) GetFrontendBandwidthUsage(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Use this method to return an array of public bandwidth utilization records between a given date range.
func (r *Hardware_Server) GetFrontendBandwidthUse(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Network_Bandwidth_Version1_Usage_Detail, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// The '''getFrontendIncomingBandwidth''' method retrieves the amount of incoming public network traffic used by a server between the given start and end date parameters. When entering the ''dateTime'' parameter, only the month, day and year of the start and end dates are required - the time (hour, minute and second) are set to midnight by default and cannot be changed. The amount of bandwidth retrieved is measured in gigabytes (GB).
func (r *Hardware_Server) GetFrontendIncomingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's front-end or public network components.
func (r *Hardware_Server) GetFrontendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getFrontendOutgoingBandwidth''' method retrieves the amount of outgoing public network traffic used by a server between the given start and end date parameters. The ''dateTime'' parameter requires only the day, month and year to be entered - the time (hour, minute and second) are set to midnight be default in order to gather the data for the entire start and end date indicated in the parameter. The amount of bandwidth retrieved is measured in gigabytes (GB).
func (r *Hardware_Server) GetFrontendOutgoingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's frontend or public router.
func (r *Hardware_Server) GetFrontendRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's universally unique identifier.
func (r *Hardware_Server) GetGlobalIdentifier() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hard drives contained within a piece of hardware.
func (r *Hardware_Server) GetHardDrives() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a server by searching for the primary IP address.
func (r *Hardware_Server) GetHardwareByIpAddress(ipAddress *string) (resp datatypes.Hardware_Server, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The chassis that a piece of hardware is housed in.
func (r *Hardware_Server) GetHardwareChassis() (resp datatypes.Hardware_Chassis, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's function.
func (r *Hardware_Server) GetHardwareFunction() (resp datatypes.Hardware_Function, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's function.
func (r *Hardware_Server) GetHardwareFunctionDescription() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's status.
func (r *Hardware_Server) GetHardwareStatus() (resp datatypes.Hardware_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Determine in hardware object has TPM enabled.
func (r *Hardware_Server) GetHasTrustedPlatformModuleBillingItemFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a host IPS software component object.
func (r *Hardware_Server) GetHostIpsSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''getHourlyBandwidth''' method retrieves all bandwidth updates hourly for the specified hardware. Because the potential number of data points can become excessive, the method limits the user to obtain data in 24-hour intervals. The required ''dateTime'' parameter is used as the starting point for the query and will be calculated for the 24-hour period starting with the specified date and time. For example, entering a parameter of
//
// '02/01/2008 0:00'
//
// results in a return of all bandwidth data for the entire day of February 1, 2008, as 0:00 specifies a midnight start date. Please note that the time entered should be completed using a 24-hour clock (military time, astronomical time).
//
// For data spanning more than a single 24-hour period, refer to the getBandwidthData function on the metricTrackingObject for the piece of hardware.
func (r *Hardware_Server) GetHourlyBandwidth(mode *string, day *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		mode,
		day,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A server's hourly billing status.
func (r *Hardware_Server) GetHourlyBillingFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The sum of all the inbound network traffic data for the last 30 days.
func (r *Hardware_Server) GetInboundBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total private inbound bandwidth for this hardware for the current billing cycle.
func (r *Hardware_Server) GetInboundPrivateBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total public inbound bandwidth for this hardware for the current billing cycle.
func (r *Hardware_Server) GetInboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Return a collection of SoftLayer_Item_Price objects from a collection of SoftLayer_Software_Description
func (r *Hardware_Server) GetItemPricesFromSoftwareDescriptions(softwareDescriptions []datatypes.Software_Description, includeTranslationsFlag *bool, returnAllPricesFlag *bool) (resp []datatypes.Product_Item, err error) {
	params := []interface{}{
		softwareDescriptions,
		includeTranslationsFlag,
		returnAllPricesFlag,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The last transaction that a server's operating system was loaded.
func (r *Hardware_Server) GetLastOperatingSystemReload() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the last transaction a server performed.
func (r *Hardware_Server) GetLastTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's latest network monitoring incident.
func (r *Hardware_Server) GetLatestNetworkMonitorIncident() (resp datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Where a piece of hardware is located within SoftLayer's location hierarchy.
func (r *Hardware_Server) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_Server) GetLocationPathString() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a lockbox account associated with a server.
func (r *Hardware_Server) GetLockboxNetworkStorage() (resp datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A flag indicating that the hardware is a managed resource.
func (r *Hardware_Server) GetManagedResourceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the remote management network component attached with this server.
func (r *Hardware_Server) GetManagementNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's memory.
func (r *Hardware_Server) GetMemory() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The amount of memory a piece of hardware has, measured in gigabytes.
func (r *Hardware_Server) GetMemoryCapacity() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's metric tracking object.
func (r *Hardware_Server) GetMetricTrackingObject() (resp datatypes.Metric_Tracking_Object_HardwareServer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The metric tracking object id for this server.
func (r *Hardware_Server) GetMetricTrackingObjectId() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Returns open monitoring alarms for a given time period
func (r *Hardware_Server) GetMonitoringActiveAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the monitoring agents associated with a piece of hardware.
func (r *Hardware_Server) GetMonitoringAgents() (resp []datatypes.Monitoring_Agent, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Returns closed monitoring alarms for a given time period
func (r *Hardware_Server) GetMonitoringClosedAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the hardware's monitoring robot.
func (r *Hardware_Server) GetMonitoringRobot() (resp datatypes.Monitoring_Robot, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's network monitoring services.
func (r *Hardware_Server) GetMonitoringServiceComponent() (resp datatypes.Network_Monitor_Version1_Query_Host_Stratum, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The monitoring service flag eligibility status for a piece of hardware.
func (r *Hardware_Server) GetMonitoringServiceEligibilityFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The service flag status for a piece of hardware.
func (r *Hardware_Server) GetMonitoringServiceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The monitoring notification objects for this hardware. Each object links this hardware instance to a user account that will be notified if monitoring on this hardware object fails
func (r *Hardware_Server) GetMonitoringUserNotification() (resp []datatypes.User_Customer_Notification_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's motherboard.
func (r *Hardware_Server) GetMotherboard() (resp datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's network cards.
func (r *Hardware_Server) GetNetworkCards() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Get the IP addresses associated with this server that are protectable by a network component firewall. Note, this may not return all values for IPv6 subnets for this server. Please use getFirewallProtectableSubnets to get all protectable subnets.
func (r *Hardware_Server) GetNetworkComponentFirewallProtectableIpAddresses() (resp []datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Returns a hardware's network components.
func (r *Hardware_Server) GetNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The gateway member if this device is part of a network gateway.
func (r *Hardware_Server) GetNetworkGatewayMember() (resp datatypes.Network_Gateway_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Whether or not this device is part of a network gateway.
func (r *Hardware_Server) GetNetworkGatewayMemberFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's network management IP address.
func (r *Hardware_Server) GetNetworkManagementIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All servers with failed monitoring that are attached downstream to a piece of hardware.
func (r *Hardware_Server) GetNetworkMonitorAttachedDownHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Virtual guests that are attached downstream to a hardware that have failed monitoring
func (r *Hardware_Server) GetNetworkMonitorAttachedDownVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The status of all of a piece of hardware's network monitoring incidents.
func (r *Hardware_Server) GetNetworkMonitorIncidents() (resp []datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's network monitors.
func (r *Hardware_Server) GetNetworkMonitors() (resp []datatypes.Network_Monitor_Version1_Query_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The value of a hardware's network status attribute.
func (r *Hardware_Server) GetNetworkStatus() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware's related network status attribute.
func (r *Hardware_Server) GetNetworkStatusAttribute() (resp datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's associated network storage service account.
func (r *Hardware_Server) GetNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The network virtual LANs (VLANs) associated with a piece of hardware's network components.
func (r *Hardware_Server) GetNetworkVlans() (resp []datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's allotted bandwidth for the next billing cycle (measured in GB).
func (r *Hardware_Server) GetNextBillingCycleBandwidthAllocation() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_Server) GetNotesHistory() (resp []datatypes.Hardware_Note, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Hardware_Server object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Hardware service. You can only retrieve servers from the account that your portal user is assigned to.
func (r *Hardware_Server) GetObject() (resp datatypes.Hardware_Server, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An open ticket requesting cancellation of this server, if one exists.
func (r *Hardware_Server) GetOpenCancellationTicket() (resp datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's operating system.
func (r *Hardware_Server) GetOperatingSystem() (resp datatypes.Software_Component_OperatingSystem, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's operating system software description.
func (r *Hardware_Server) GetOperatingSystemReferenceCode() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The sum of all the outbound network traffic data for the last 30 days.
func (r *Hardware_Server) GetOutboundBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total private outbound bandwidth for this hardware for the current billing cycle.
func (r *Hardware_Server) GetOutboundPrivateBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total public outbound bandwidth for this hardware for the current billing cycle.
func (r *Hardware_Server) GetOutboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Whether the bandwidth usage for this hardware for the current billing cycle exceeds the allocation.
func (r *Hardware_Server) GetOverBandwidthAllocationFlag() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a server's hardware state via its internal sensors. Remote sensor data is transmitted to the SoftLayer API by way of the server's remote management card. Sensor data measures system temperatures, voltages, and other local server settings. Sensor data is cached for 30 seconds. Calls made to getSensorData for the same server within 30 seconds of each other will return the same data. Subsequent calls will return new data once the cache expires.
func (r *Hardware_Server) GetPMInfo() (resp []datatypes.Container_RemoteManagement_PmInfo, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the Point of Presence (PoP) location in which a piece of hardware resides.
func (r *Hardware_Server) GetPointOfPresenceLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The power components for a hardware object.
func (r *Hardware_Server) GetPowerComponents() (resp []datatypes.Hardware_Power_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's power supply.
func (r *Hardware_Server) GetPowerSupply() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware's primary private IP address.
func (r *Hardware_Server) GetPrimaryBackendIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the hardware's primary back-end network component.
func (r *Hardware_Server) GetPrimaryBackendNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Hardware_Server) GetPrimaryDriveSize() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The hardware's primary public IP address.
func (r *Hardware_Server) GetPrimaryIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the hardware's primary public network component.
func (r *Hardware_Server) GetPrimaryNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a graph of a server's private network bandwidth usage over the specified timeframe. If no timeframe is specified then getPublicBandwidthGraphImage retrieves the last 24 hours of public bandwidth usage. getPrivateBandwidthGraphImage returns a PNG image measuring 827 pixels by 293 pixels.
func (r *Hardware_Server) GetPrivateBandwidthData(startTime *int, endTime *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve a brief summary of a server's private network bandwidth usage. getPrivateBandwidthDataSummary retrieves a server's bandwidth allocation for its billing period, its estimated usage during its billing period, and an estimation of how much bandwidth it will use during its billing period based on its current usage. A server's projected bandwidth usage increases in accuracy as it progresses through its billing period.
func (r *Hardware_Server) GetPrivateBandwidthDataSummary() (resp datatypes.Container_Network_Bandwidth_Data_Summary, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a graph of a server's private network bandwidth usage over the specified time frame. If no time frame is specified then getPublicBandwidthGraphImage retrieves the last 24 hours of public bandwidth usage. getPublicBandwidthGraphImage returns a PNG image
func (r *Hardware_Server) GetPrivateBandwidthGraphImage(startTime *string, endTime *string) (resp []byte, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A server's primary private IP address.
func (r *Hardware_Server) GetPrivateIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the private network component attached with this server.
func (r *Hardware_Server) GetPrivateNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Whether the hardware only has access to the private network.
func (r *Hardware_Server) GetPrivateNetworkOnlyFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the backend VLAN for the primary IP address of the server
func (r *Hardware_Server) GetPrivateVlan() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a backend network VLAN by searching for an IP address
func (r *Hardware_Server) GetPrivateVlanByIpAddress(ipAddress *string) (resp datatypes.Network_Vlan, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total number of processor cores, summed from all processors that are attached to a piece of hardware
func (r *Hardware_Server) GetProcessorCoreAmount() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total number of physical processor cores, summed from all processors that are attached to a piece of hardware
func (r *Hardware_Server) GetProcessorPhysicalCoreAmount() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's processors.
func (r *Hardware_Server) GetProcessors() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Whether the bandwidth usage for this hardware for the current billing cycle is projected to exceed the allocation.
func (r *Hardware_Server) GetProjectedOverBandwidthAllocationFlag() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The projected public outbound bandwidth for this hardware for the current billing cycle.
func (r *Hardware_Server) GetProjectedPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Hardware_Server) GetProvisionDate() (resp datatypes.Time, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a graph of a server's public network bandwidth usage over the specified timeframe. If no timeframe is specified then getPublicBandwidthGraphImage retrieves the last 24 hours of public bandwidth usage. getPublicBandwidthGraphImage returns a PNG image measuring 827 pixels by 293 pixels.
func (r *Hardware_Server) GetPublicBandwidthData(startTime *int, endTime *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve a brief summary of a server's public network bandwidth usage. getPublicBandwidthDataSummary retrieves a server's bandwidth allocation for its billing period, its estimated usage during its billing period, and an estimation of how much bandwidth it will use during its billing period based on its current usage. A server's projected bandwidth usage increases in accuracy as it progresses through its billing period.
func (r *Hardware_Server) GetPublicBandwidthDataSummary() (resp datatypes.Container_Network_Bandwidth_Data_Summary, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a graph of a server's public network bandwidth usage over the specified time frame. If no time frame is specified then getPublicBandwidthGraphImage retrieves the last 24 hours of public bandwidth usage. getPublicBandwidthGraphImage returns a PNG image measuring 827 pixels by 293 pixels.  THIS METHOD GENERATES GRAPHS BASED ON THE NEW DATA WAREHOUSE REPOSITORY.
func (r *Hardware_Server) GetPublicBandwidthGraphImage(startTime *datatypes.Time, endTime *datatypes.Time) (resp []byte, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve the total number of bytes used by a server over a specified time period via the data warehouse tracking objects for this hardware.
func (r *Hardware_Server) GetPublicBandwidthTotal(startTime *int, endTime *int) (resp uint, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve a SoftLayer server's public network component. Some servers are only connected to the private network and may not have a public network component. In that case getPublicNetworkComponent returns a null object.
func (r *Hardware_Server) GetPublicNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the frontend VLAN for the primary IP address of the server
func (r *Hardware_Server) GetPublicVlan() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the frontend network Vlan by searching the hostname of a server
func (r *Hardware_Server) GetPublicVlanByHostname(hostname *string) (resp datatypes.Network_Vlan, err error) {
	params := []interface{}{
		hostname,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_Server) GetRack() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The RAID controllers contained within a piece of hardware.
func (r *Hardware_Server) GetRaidControllers() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Recent events that impact this hardware.
func (r *Hardware_Server) GetRecentEvents() (resp []datatypes.Notification_Occurrence_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The last five commands issued to the server's remote management card.
func (r *Hardware_Server) GetRecentRemoteManagementCommands() (resp []datatypes.Hardware_Component_RemoteManagement_Command_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_Server) GetRegionalInternetRegistry() (resp datatypes.Network_Regional_Internet_Registry, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A server's remote management card.
func (r *Hardware_Server) GetRemoteManagement() (resp datatypes.Hardware_Component_RemoteManagement, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve User credentials to issue commands and/or interact with the server's remote management card.
func (r *Hardware_Server) GetRemoteManagementAccounts() (resp []datatypes.Hardware_Component_RemoteManagement_User, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's associated remote management component. This is normally IPMI.
func (r *Hardware_Server) GetRemoteManagementComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve User(s) who have access to issue commands and/or interact with the server's remote management card.
func (r *Hardware_Server) GetRemoteManagementUsers() (resp []datatypes.Hardware_Component_RemoteManagement_User, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_Server) GetResourceGroupMemberReferences() (resp []datatypes.Resource_Group_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_Server) GetResourceGroupRoles() (resp []datatypes.Resource_Group_Role, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The resource groups in which this hardware is a member.
func (r *Hardware_Server) GetResourceGroups() (resp []datatypes.Resource_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the reverse domain records associated with this server.
func (r *Hardware_Server) GetReverseDomainRecords() (resp []datatypes.Dns_Domain, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's routers.
func (r *Hardware_Server) GetRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Collection of scale assets this hardware corresponds to.
func (r *Hardware_Server) GetScaleAssets() (resp []datatypes.Scale_Asset, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's vulnerability scan requests.
func (r *Hardware_Server) GetSecurityScanRequests() (resp []datatypes.Network_Security_Scanner_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a server's hardware state via its internal sensors. Remote sensor data is transmitted to the SoftLayer API by way of the server's remote management card. Sensor data measures system temperatures, voltages, and other local server settings. Sensor data is cached for 30 seconds. Calls made to getSensorData for the same server within 30 seconds of each other will return the same data. Subsequent calls will return new data once the cache expires.
func (r *Hardware_Server) GetSensorData() (resp []datatypes.Container_RemoteManagement_SensorReading, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieves the raw data returned from the server's remote management card.  For more details of what is returned please refer to the getSensorData method.  Along with the raw data, graphs for the cpu and system temperatures and fan speeds are also returned.
func (r *Hardware_Server) GetSensorDataWithGraphs() (resp datatypes.Container_RemoteManagement_SensorReadingsWithGraphs, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a server's hardware components, software, and network components. getServerDetails is an aggregation function that combines the results of [[SoftLayer_Hardware_Server::getComponents]], [[SoftLayer_Hardware_Server::getSoftware]], and [[SoftLayer_Hardware_Server::getNetworkComponents]] in a single container.
func (r *Hardware_Server) GetServerDetails() (resp datatypes.Container_Hardware_Server_Details, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the server's fan speeds and displays them using tachometer graphs.  Data used to construct graphs is retrieved from the server's remote management card.  All graphs returned will have a title associated with it.
func (r *Hardware_Server) GetServerFanSpeedGraphs() (resp []datatypes.Container_RemoteManagement_Graphs_SensorSpeed, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieves the power state for the server.  The server's power status is retrieved from its remote management card.  This will return 'on' or 'off'.
func (r *Hardware_Server) GetServerPowerState() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the server room in which the hardware is located.
func (r *Hardware_Server) GetServerRoom() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the server's temperature and displays them using thermometer graphs.  Temperatures retrieved are CPU(s) and system temperatures.  Data used to construct graphs is retrieved from the server's remote management card.  All graphs returned will have a title associated with it.
func (r *Hardware_Server) GetServerTemperatureGraphs() (resp []datatypes.Container_RemoteManagement_Graphs_SensorTemperature, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the piece of hardware's service provider.
func (r *Hardware_Server) GetServiceProvider() (resp datatypes.Service_Provider, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's installed software.
func (r *Hardware_Server) GetSoftwareComponents() (resp []datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the billing item for a spare pool server.
func (r *Hardware_Server) GetSparePoolBillingItem() (resp datatypes.Billing_Item_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve SSH keys to be installed on the server during provisioning or an OS reload.
func (r *Hardware_Server) GetSshKeys() (resp []datatypes.Security_Ssh_Key, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A server's remote management card used for statistics.
func (r *Hardware_Server) GetStatisticsRemoteManagement() (resp datatypes.Hardware_Component_RemoteManagement, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_Server) GetStorageNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_Server) GetTagReferences() (resp []datatypes.Tag_Reference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Hardware_Server) GetTopLevelLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
// This method will query transaction history for a piece of hardware.
func (r *Hardware_Server) GetTransactionHistory() (resp []datatypes.Provisioning_Version1_Transaction_History, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a list of upgradeable items available to this piece of hardware. Currently, getUpgradeItemPrices retrieves upgrades available for a server's memory, hard drives, network port speed, bandwidth allocation and GPUs.
func (r *Hardware_Server) GetUpgradeItemPrices() (resp []datatypes.Product_Item_Price, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An account's associated upgrade request object, if any.
func (r *Hardware_Server) GetUpgradeRequest() (resp datatypes.Product_Upgrade_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The network device connected to a piece of hardware.
func (r *Hardware_Server) GetUplinkHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the network component that is one level higher than a piece of hardware on the network infrastructure.
func (r *Hardware_Server) GetUplinkNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A string containing custom user data for a hardware order.
func (r *Hardware_Server) GetUserData() (resp []datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A list of users that have access to this computing instance.
func (r *Hardware_Server) GetUsers() (resp []datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This method will return the list of block device template groups that are valid to the host. For instance, it will only retrieve FLEX images.
func (r *Hardware_Server) GetValidBlockDeviceTemplateGroups(visibility *string) (resp []datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	params := []interface{}{
		visibility,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the virtual chassis for a piece of hardware.
func (r *Hardware_Server) GetVirtualChassis() (resp datatypes.Hardware_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the virtual chassis siblings for a piece of hardware.
func (r *Hardware_Server) GetVirtualChassisSiblings() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware server's virtual servers.
func (r *Hardware_Server) GetVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's virtual host record.
func (r *Hardware_Server) GetVirtualHost() (resp datatypes.Virtual_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's virtual software licenses.
func (r *Hardware_Server) GetVirtualLicenses() (resp []datatypes.Software_VirtualLicense, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information regarding the bandwidth allotment to which a piece of hardware belongs.
func (r *Hardware_Server) GetVirtualRack() (resp datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The name of the bandwidth allotment belonging to a piece of hardware.
func (r *Hardware_Server) GetVirtualRackId() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The name of the bandwidth allotment belonging to a piece of hardware.
func (r *Hardware_Server) GetVirtualRackName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's virtualization platform software.
func (r *Hardware_Server) GetVirtualizationPlatform() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a list of Windows updates available for a server from the local SoftLayer Windows Server Update Services (WSUS) server. Windows servers provisioned by SoftLayer are configured to use the local WSUS server via the private network by default.
func (r *Hardware_Server) GetWindowsUpdateAvailableUpdates() (resp []datatypes.Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a list of Windows updates installed on a server as reported by the local SoftLayer Windows Server Update Services (WSUS) server. Windows servers provisioned by SoftLayer are configured to use the local WSUS server via the private network by default.
func (r *Hardware_Server) GetWindowsUpdateInstalledUpdates() (resp []datatypes.Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This method returns an update status record for this server.  That record will specify if the server is missing updates, or has updates that must be reinstalled or require a reboot to go into affect.
func (r *Hardware_Server) GetWindowsUpdateStatus() (resp datatypes.Container_Utility_Microsoft_Windows_UpdateServices_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The '''importVirtualHost''' method attempts to import the host record for the virtualization platform running on a server.
func (r *Hardware_Server) ImportVirtualHost() (resp datatypes.Virtual_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Idera Bare Metal Server Restore is a backup agent designed specifically for making full system restores made with Idera Server Backup.
func (r *Hardware_Server) InitiateIderaBareMetalRestore() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// R1Soft Bare Metal Server Restore is an R1Soft disk agent designed specifically for making full system restores made with R1Soft CDP Server backup.
func (r *Hardware_Server) InitiateR1SoftBareMetalRestore() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Issues a ping command and returns the success (true) or failure (false) of the ping command.
func (r *Hardware_Server) IsBackendPingable() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Issues a ping command and returns the success (true) or failure (false) of the ping command.
func (r *Hardware_Server) IsPingable() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Determine if the server runs any version of the Microsoft Windows operating systems. Return ''true'' if it does and ''false if otherwise.
func (r *Hardware_Server) IsWindowsServer() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Issues a ping command to the server and returns the ping response.
func (r *Hardware_Server) Ping() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Power off then power on the server via powerstrip.  The power cycle command is equivalent to unplugging the server from the powerstrip and then plugging the server back into the powerstrip.  This should only be used as a last resort.  If a reboot command has been issued successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard, powerOn, powerOff and powerCycle) will not be allowed. This is to avoid any type of server failures.
func (r *Hardware_Server) PowerCycle() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This method will power off the server via the server's remote management card.
func (r *Hardware_Server) PowerOff() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Power on server via its remote management card.  If a reboot command has been issued successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard, powerOn, powerOff and powerCycle) will not be allowed.  This is to avoid any type of server failures.
func (r *Hardware_Server) PowerOn() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Attempts to reboot the server by issuing a reset (soft reboot) command to the server's remote management card. If the reset (soft reboot) attempt is unsuccessful, a power cycle command will be issued via the powerstrip. The power cycle command is equivalent to unplugging the server from the powerstrip and then plugging the server back into the powerstrip.  If a reboot command has been issued successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard, powerOn, powerOff and powerCycle) will not be allowed.  This is to avoid any type of server failures.
func (r *Hardware_Server) RebootDefault() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Reboot the server by issuing a cycle command to the server's remote management card.  This is equivalent to pressing the 'Reset' button on the server.  This command is issued immediately and will not wait for processes to shutdown. After this command is issued, the server may take a few moments to boot up as server may run system disks checks. If a reboot command has been issued successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard, powerOn, powerOff and powerCycle) will not be allowed.  This is to avoid any type of server failures.
func (r *Hardware_Server) RebootHard() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Reboot the server by issuing a reset command to the server's remote management card.  This is a graceful reboot. The servers will allow all process to shutdown gracefully before rebooting.  If a reboot command has been issued successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard, powerOn, powerOff and powerCycle) will not be allowed.  This is to avoid any type of server failures.
func (r *Hardware_Server) RebootSoft() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Reloads current operating system configuration.
//
// This service has a confirmation protocol for proceeding with the reload. To proceed with the reload without confirmation, simply pass in 'FORCE' as the token parameter. To proceed with the reload with confirmation, simply call the service with no parameter. A token string will be returned by this service. The token will remain active for 10 minutes. Use this token as the parameter to confirm that a reload is to be performed for the server.
//
// As a precaution, we strongly  recommend backing up all data before reloading the operating system. The reload will format the primary disk and will reconfigure the server to the current specifications on record.
//
// The reload will take AT MINIMUM 66 minutes.
func (r *Hardware_Server) ReloadCurrentOperatingSystemConfiguration(token *string) (resp string, err error) {
	params := []interface{}{
		token,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Reloads current or customer specified operating system configuration.
//
// This service has a confirmation protocol for proceeding with the reload. To proceed with the reload without confirmation, simply pass in 'FORCE' as the token parameter. To proceed with the reload with confirmation, simply call the service with no parameter. A token string will be returned by this service. The token will remain active for 10 minutes. Use this token as the parameter to confirm that a reload is to be performed for the server.
//
// As a precaution, we strongly  recommend backing up all data before reloading the operating system. The reload will format the primary disk and will reconfigure the server to the current specifications on record.
//
// The reload will take AT MINIMUM 66 minutes.
func (r *Hardware_Server) ReloadOperatingSystem(token *string, config *datatypes.Container_Hardware_Server_Configuration) (resp string, err error) {
	params := []interface{}{
		token,
		config,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// This method is used to remove access to s SoftLayer_Network_Storage volumes that supports host- or network-level access control.
func (r *Hardware_Server) RemoveAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// This method is used to allow access to multiple SoftLayer_Network_Storage volumes that support host- or network-level access control.
func (r *Hardware_Server) RemoveAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// You can launch a new Passmark hardware test by selecting from your server list. It will bring your server offline for approximately 20 minutes while the testing is in progress, and will publish a certificate with the results to your hardware details page.
//
// While the hard drives are tested for the initial deployment, the Passmark Certificate utility will not test the hard drives on your live server. This is to ensure that no data is overwritten. If you would like to test the server's hard drives, you can have the full Passmark suite installed to your server free of charge through a new Support ticket.
//
// While the test itself does not overwrite any data on the server, it is recommended that you make full off-server backups of all data prior to launching the test. The Passmark hardware test is designed to force any latent hardware issues to the surface, so hardware failure is possible.
//
// In the event of a hardware failure during this test our datacenter engineers will be notified of the problem automatically. They will then replace any failed components to bring your server back online, and will be contacting you to ensure that impact on your server is minimal.
func (r *Hardware_Server) RunPassmarkCertificationBenchmark() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Changes the password that we have stored in our database for a servers' Operating System
func (r *Hardware_Server) SetOperatingSystemPassword(newPassword *string) (resp bool, err error) {
	params := []interface{}{
		newPassword,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Sets the private network interface speed to the new speed. Speed values can only be 0 (Disconnect), 10, 100, 1000, and 10000. The new speed must be equal to or less than the max speed of the interface.
//
// It will take less than a minute to update the switch port speed. The server uplink will not be operational again until the server interface speed is updated.
func (r *Hardware_Server) SetPrivateNetworkInterfaceSpeed(newSpeed *int) (resp bool, err error) {
	params := []interface{}{
		newSpeed,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Sets the public network interface speed to the new speed. Speed values can only be 0 (Disconnect), 10, 100, 1000, and 10000. The new speed must be equal to or less than the max speed of the interface.
//
// It will take less than a minute to update the switch port speed. The server uplink will not be operational again until the server interface speed is updated.
func (r *Hardware_Server) SetPublicNetworkInterfaceSpeed(newSpeed *int) (resp bool, err error) {
	params := []interface{}{
		newSpeed,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Hardware_Server) SetTags(tags *string) (resp bool, err error) {
	params := []interface{}{
		tags,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Sets the data that will be written to the configuration drive.
func (r *Hardware_Server) SetUserMetadata(metadata []string) (resp []datatypes.Hardware_Attribute, err error) {
	params := []interface{}{
		metadata,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Shuts down the public network port
func (r *Hardware_Server) ShutdownPrivatePort() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Shuts down the public network port
func (r *Hardware_Server) ShutdownPublicPort() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The ability to place bare metal servers in a state where they are powered down, and ports closed yet still allocated to the customer as a part of the Spare Pool program.
func (r *Hardware_Server) SparePool(action *string, newOrder *bool) (resp bool, err error) {
	params := []interface{}{
		action,
		newOrder,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Validates a collection of partitions for an operating system
func (r *Hardware_Server) ValidatePartitionsForOperatingSystem(operatingSystem *datatypes.Software_Description, partitions []datatypes.Hardware_Component_Partition) (resp bool, err error) {
	params := []interface{}{
		operatingSystem,
		partitions,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
