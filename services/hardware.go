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

// The SoftLayer_Hardware data type contains general information relating to a single SoftLayer hardware.
type Hardware struct {
	Session session.SLSession
	Options sl.Options
}

// GetHardwareService returns an instance of the Hardware SoftLayer service
func GetHardwareService(sess session.SLSession) Hardware {
	return Hardware{Session: sess}
}

func (r Hardware) Id(id int) Hardware {
	r.Options.Id = &id
	return r
}

func (r Hardware) Mask(mask string) Hardware {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
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

// getObject retrieves the SoftLayer_Hardware object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Hardware service. You can only retrieve the account that your portal user is assigned to.
func (r Hardware) GetObject() (resp datatypes.Hardware, err error) {
	err = r.Session.DoRequest("SoftLayer_Hardware", "getObject", nil, &r.Options, &resp)
	return
}

// The SoftLayer_Hardware_Server data type contains general information relating to a single SoftLayer server.
type Hardware_Server struct {
	Session session.SLSession
	Options sl.Options
}

// GetHardwareServerService returns an instance of the Hardware_Server SoftLayer service
func GetHardwareServerService(sess session.SLSession) Hardware_Server {
	return Hardware_Server{Session: sess}
}

func (r Hardware_Server) Id(id int) Hardware_Server {
	r.Options.Id = &id
	return r
}

func (r Hardware_Server) Mask(mask string) Hardware_Server {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
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

// This method is used to allow access to multiple SoftLayer_Network_Storage volumes that support host- or network-level access control.
func (r Hardware_Server) AllowAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "allowAccessToNetworkStorageList", params, &r.Options, &resp)
	return
}

// The Rescue Kernel is designed to provide you with the ability to bring a server online in order to troubleshoot system problems that would normally only be resolved by an OS Reload. The correct Rescue Kernel will be selected based upon the currently installed operating system. When the rescue kernel process is initiated, the server will shutdown and reboot on to the public network with the same IP's assigned to the server to allow for remote connections. It will bring your server offline for approximately 10 minutes while the rescue is in progress. The root/administrator password will be the same as what is listed in the portal for the server.
func (r Hardware_Server) BootToRescueLayer(noOsBootEnvironment *string) (resp bool, err error) {
	params := []interface{}{
		noOsBootEnvironment,
	}
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "bootToRescueLayer", params, &r.Options, &resp)
	return
}

// You can launch firmware reflash by selecting from your server list. It will bring your server offline for approximately 60 minutes while the flashes are in progress.
//
// In the event of a hardware failure during this our datacenter engineers will be notified of the problem automatically. They will then replace any failed components to bring your server back online, and will be contacting you to ensure that impact on your server is minimal.
func (r Hardware_Server) CreateFirmwareReflashTransaction(ipmi *int, raidController *int, bios *int) (resp bool, err error) {
	params := []interface{}{
		ipmi,
		raidController,
		bios,
	}
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "createFirmwareReflashTransaction", params, &r.Options, &resp)
	return
}

// You can launch firmware updates by selecting from your server list. It will bring your server offline for approximately 20 minutes while the updates are in progress.
//
// In the event of a hardware failure during this test our datacenter engineers will be notified of the problem automatically. They will then replace any failed components to bring your server back online, and will be contacting you to ensure that impact on your server is minimal.
func (r Hardware_Server) CreateFirmwareUpdateTransaction(ipmi *int, raidController *int, bios *int, harddrive *int, networkCard *int) (resp bool, err error) {
	params := []interface{}{
		ipmi,
		raidController,
		bios,
		harddrive,
		networkCard,
	}
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "createFirmwareUpdateTransaction", params, &r.Options, &resp)
	return
}

// Edit a server's properties
func (r Hardware_Server) EditObject(templateObject *datatypes.Hardware_Server) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "editObject", params, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's active physical components.
func (r Hardware_Server) GetActiveComponents() (resp []datatypes.Hardware_Component, err error) {
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "getActiveComponents", nil, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Network_Storage_Allowed_Host information to connect this server to Network Storage volumes that require access control lists.
func (r Hardware_Server) GetAllowedHost() (resp datatypes.Network_Storage_Allowed_Host, err error) {
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "getAllowedHost", nil, &r.Options, &resp)
	return
}

// This method is retrieve a list of SoftLayer_Network_Storage volumes that are authorized access to this SoftLayer_Hardware.
func (r Hardware_Server) GetAttachedNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "getAttachedNetworkStorages", params, &r.Options, &resp)
	return
}

// Retrieve A hardware's allotted detail record. Allotment details link bandwidth allocation with allotments.
func (r Hardware_Server) GetBandwidthAllotmentDetail() (resp datatypes.Network_Bandwidth_Version1_Allotment_Detail, err error) {
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "getBandwidthAllotmentDetail", nil, &r.Options, &resp)
	return
}

// Retrieve The raw bandwidth usage data for the current billing cycle. One object will be returned for each network this server is attached to.
func (r Hardware_Server) GetBillingCycleBandwidthUsage() (resp []datatypes.Network_Bandwidth_Usage, err error) {
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "getBillingCycleBandwidthUsage", nil, &r.Options, &resp)
	return
}

// Retrieve Information regarding the billing item for a server.
func (r Hardware_Server) GetBillingItem() (resp datatypes.Billing_Item_Hardware, err error) {
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "getBillingItem", nil, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's components.
func (r Hardware_Server) GetComponents() (resp []datatypes.Hardware_Component, err error) {
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "getComponents", nil, &r.Options, &resp)
	return
}

// Retrieve A piece of hardware's front-end or public network components.
func (r Hardware_Server) GetFrontendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "getFrontendNetworkComponents", nil, &r.Options, &resp)
	return
}

// Retrieve The hard drives contained within a piece of hardware.
func (r Hardware_Server) GetHardDrives() (resp []datatypes.Hardware_Component, err error) {
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "getHardDrives", nil, &r.Options, &resp)
	return
}

// Retrieve The metric tracking object id for this server.
func (r Hardware_Server) GetMetricTrackingObjectId() (resp int, err error) {
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "getMetricTrackingObjectId", nil, &r.Options, &resp)
	return
}

// Retrieve Returns a hardware's network components.
func (r Hardware_Server) GetNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "getNetworkComponents", nil, &r.Options, &resp)
	return
}

// Retrieve The network virtual LANs (VLANs) associated with a piece of hardware's network components.
func (r Hardware_Server) GetNetworkVlans() (resp []datatypes.Network_Vlan, err error) {
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "getNetworkVlans", nil, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Hardware_Server object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Hardware service. You can only retrieve servers from the account that your portal user is assigned to.
func (r Hardware_Server) GetObject() (resp datatypes.Hardware_Server, err error) {
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "getObject", nil, &r.Options, &resp)
	return
}

// Retrieve Information regarding a piece of hardware's operating system.
func (r Hardware_Server) GetOperatingSystem() (resp datatypes.Software_Component_OperatingSystem, err error) {
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "getOperatingSystem", nil, &r.Options, &resp)
	return
}

// Retrieve User credentials to issue commands and/or interact with the server's remote management card.
func (r Hardware_Server) GetRemoteManagementAccounts() (resp []datatypes.Hardware_Component_RemoteManagement_User, err error) {
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "getRemoteManagementAccounts", nil, &r.Options, &resp)
	return
}

// Retrieve a server's hardware state via its internal sensors. Remote sensor data is transmitted to the SoftLayer API by way of the server's remote management card. Sensor data measures system temperatures, voltages, and other local server settings. Sensor data is cached for 30 seconds. Calls made to getSensorData for the same server within 30 seconds of each other will return the same data. Subsequent calls will return new data once the cache expires.
func (r Hardware_Server) GetSensorData() (resp []datatypes.Container_RemoteManagement_SensorReading, err error) {
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "getSensorData", nil, &r.Options, &resp)
	return
}

// Retrieve
func (r Hardware_Server) GetTagReferences() (resp []datatypes.Tag_Reference, err error) {
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "getTagReferences", nil, &r.Options, &resp)
	return
}

// Power off then power on the server via powerstrip.  The power cycle command is equivalent to unplugging the server from the powerstrip and then plugging the server back into the powerstrip.  This should only be used as a last resort.  If a reboot command has been issued successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard, powerOn, powerOff and powerCycle) will not be allowed. This is to avoid any type of server failures.
func (r Hardware_Server) PowerCycle() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "powerCycle", nil, &r.Options, &resp)
	return
}

// This method will power off the server via the server's remote management card.
func (r Hardware_Server) PowerOff() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "powerOff", nil, &r.Options, &resp)
	return
}

// Power on server via its remote management card.  If a reboot command has been issued successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard, powerOn, powerOff and powerCycle) will not be allowed.  This is to avoid any type of server failures.
func (r Hardware_Server) PowerOn() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "powerOn", nil, &r.Options, &resp)
	return
}

// Attempts to reboot the server by issuing a reset (soft reboot) command to the server's remote management card. If the reset (soft reboot) attempt is unsuccessful, a power cycle command will be issued via the powerstrip. The power cycle command is equivalent to unplugging the server from the powerstrip and then plugging the server back into the powerstrip.  If a reboot command has been issued successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard, powerOn, powerOff and powerCycle) will not be allowed.  This is to avoid any type of server failures.
func (r Hardware_Server) RebootDefault() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "rebootDefault", nil, &r.Options, &resp)
	return
}

// Reboot the server by issuing a reset command to the server's remote management card.  This is a graceful reboot. The servers will allow all process to shutdown gracefully before rebooting.  If a reboot command has been issued successfully in the past 20 minutes, another remote management command (rebootSoft, rebootHard, powerOn, powerOff and powerCycle) will not be allowed.  This is to avoid any type of server failures.
func (r Hardware_Server) RebootSoft() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "rebootSoft", nil, &r.Options, &resp)
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
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "reloadOperatingSystem", params, &r.Options, &resp)
	return
}

// Set the private network interface speed and redundancy configuration.
//
// Possible $newSpeed values are -1 (maximum available), 0 (disconnect), 10, 100, 1000, and 10000; not all values are available to every server. The maximum speed is limited by the speed requested during provisioning. All intermediate speeds are limited by the capability of the pod the server is deployed in. No guarantee is made that a speed other than what was requested during provisioning will be available.
//
// If specified, possible $redundancy values are either "redundant" or "degraded". Not specifying a redundancy mode will use the best possible redundancy available to the server. However, specifying a redundacy mode that is not available to the server will result in an error. "redundant" indicates all available interfaces should be active. "degraded" indicates only the primary interface should be active. Irrespective of the number of interfaces available to a server, it is only possible to have either a single interface or all interfaces active.
//
// Receipt of a response does not indicate completion of the configuration change. Any subsequent attempts to request the interface change speed or state, while changes are pending, will result in a busy error.
//
// A response of true indicates a change was required to achieve the desired interface configuration; thus changes are pending. A response of false indicates the current interface configuration matches the desired configuration, and thus no changes are pending.
//
// <h4>Backwards Compatibility Until February 27th, 2019</h4>
//
// In order to provide a period of transition to the new API, some backwards compatible behaviors will be active during this period. <ul> <li> A "doubled" (eg. 200) speed value will be translated to a redundancy value of "redundant". If a redundancy value is specified, it is assumed no translation is needed and will result in an error due to doubled speeds no longer being valid.</li> <li> A non-doubled (eg. 100) speed value <i>without</i> a redundancy value will be translated to a redundancy value of "degraded".</li> </ul> After the compatibility period, a doubled speed value will result in an error, and a non-doubled speed value without a redundancy value specified will result in the best available redundancy state. An exception is made for the new relative speed value -1. When using -1 without a redundancy value, the best possible redundancy will be used. Please transition away from using doubled speed values in favor of specifying redundancy (when applicable) or using relative speed values 0 and -1.
func (r Hardware_Server) SetPrivateNetworkInterfaceSpeed(newSpeed *int, redundancy *string) (resp bool, err error) {
	params := []interface{}{
		newSpeed,
		redundancy,
	}
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "setPrivateNetworkInterfaceSpeed", params, &r.Options, &resp)
	return
}

// Set the public network interface speed and redundancy configuration.
//
// Possible $newSpeed values are -1 (maximum available), 0 (disconnect), 10, 100, 1000, and 10000; not all values are available to every server. The maximum speed is limited by the speed requested during provisioning. All intermediate speeds are limited by the capability of the pod the server is deployed in. No guarantee is made that a speed other than what was requested during provisioning will be available.
//
// If specified, possible $redundancy values are either "redundant" or "degraded". Not specifying a redundancy mode will use the best possible redundancy available to the server. However, specifying a redundacy mode that is not available to the server will result in an error. "redundant" indicates all available interfaces should be active. "degraded" indicates only the primary interface should be active. Irrespective of the number of interfaces available to a server, it is only possible to have either a single interface or all interfaces active.
//
// Receipt of a response does not indicate completion of the configuration change. Any subsequent attempts to request the interface change speed or state, while changes are pending, will result in a busy error.
//
// A response of true indicates a change was required to achieve the desired interface configuration; thus changes are pending. A response of false indicates the current interface configuration matches the desired configuration, and thus no changes are pending.
//
// <h4>Backwards Compatibility Until February 27th, 2019</h4>
//
// In order to provide a period of transition to the new API, some backwards compatible behaviors will be active during this period. <ul> <li> A "doubled" (eg. 200) speed value will be translated to a redundancy value of "redundant". If a redundancy value is specified, it is assumed no translation is needed and will result in an error due to doubled speeds no longer being valid.</li> <li> A non-doubled (eg. 100) speed value <i>without</i> a redundancy value will be translated to a redundancy value of "degraded".</li> </ul> After the compatibility period, a doubled speed value will result in an error, and a non-doubled speed value without a redundancy value specified will result in the best available redundancy state. An exception is made for the new relative speed value -1. When using -1 without a redundancy value, the best possible redundancy will be used. Please transition away from using doubled speed values in favor of specifying redundancy (when applicable) or using relative speed values 0 and -1.
func (r Hardware_Server) SetPublicNetworkInterfaceSpeed(newSpeed *int, redundancy *string) (resp bool, err error) {
	params := []interface{}{
		newSpeed,
		redundancy,
	}
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "setPublicNetworkInterfaceSpeed", params, &r.Options, &resp)
	return
}

// no documentation yet
func (r Hardware_Server) SetTags(tags *string) (resp bool, err error) {
	params := []interface{}{
		tags,
	}
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "setTags", params, &r.Options, &resp)
	return
}

// Sets the data that will be written to the configuration drive.
func (r Hardware_Server) SetUserMetadata(metadata []string) (resp []datatypes.Hardware_Attribute, err error) {
	params := []interface{}{
		metadata,
	}
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "setUserMetadata", params, &r.Options, &resp)
	return
}

// Attempt to toggle the IPMI interface.  If there is an active transaction on the server, it will throw an exception. This method creates a job to toggle the interface.  It is not instant.
func (r Hardware_Server) ToggleManagementInterface(enabled *bool) (resp bool, err error) {
	params := []interface{}{
		enabled,
	}
	err = r.Session.DoRequest("SoftLayer_Hardware_Server", "toggleManagementInterface", params, &r.Options, &resp)
	return
}
