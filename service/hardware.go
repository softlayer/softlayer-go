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

package service

import (
	"time"

	"github.ibm.com/riethm/gopherlayer/datatypes"
)

type Hardware struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareService() Hardware {
	return Hardware{Session: r}
}

func (r *Hardware) AllowAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) AllowAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) CaptureImage(captureTemplate *datatypes.Container_Disk_Image_Capture_Template) (resp datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	params := []interface{}{
		captureTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) CloseAlarm(alarmId *string) (resp bool, err error) {
	params := []interface{}{
		alarmId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) CreateObject(templateObject *datatypes.Hardware) (resp datatypes.Hardware, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) DeleteSoftwareComponentPasswords(softwareComponentPasswords []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		softwareComponentPasswords,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) EditSoftwareComponentPasswords(softwareComponentPasswords []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		softwareComponentPasswords,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) ExecuteRemoteScript(uri *string) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		uri,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) FindByIpAddress(ipAddress *string) (resp datatypes.Hardware, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GenerateOrderTemplate(templateObject *datatypes.Hardware) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetAlarmHistory(startDate *time.Time, endDate *time.Time, alarmId *string) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
		alarmId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetAttachedNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetAvailableNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetBackendIncomingBandwidth(startDate *time.Time, endDate *time.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetBackendOutgoingBandwidth(startDate *time.Time, endDate *time.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetCreateObjectOptions() (resp datatypes.Container_Hardware_Configuration, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetCurrentBillingDetail() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetCurrentBillingTotal() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetDailyAverage(startDate *time.Time, endDate *time.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetFrontendIncomingBandwidth(startDate *time.Time, endDate *time.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetFrontendOutgoingBandwidth(startDate *time.Time, endDate *time.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetHourlyBandwidth(mode *string, day *time.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		mode,
		day,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetMonitoringActiveAlarms(startDate *time.Time, endDate *time.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetMonitoringClosedAlarms(startDate *time.Time, endDate *time.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetObject() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetPrivateBandwidthData(startTime *int, endTime *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetPublicBandwidthData(startTime *int, endTime *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetSensorData() (resp []datatypes.Container_RemoteManagement_SensorReading, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetSensorDataWithGraphs() (resp datatypes.Container_RemoteManagement_SensorReadingsWithGraphs, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetServerFanSpeedGraphs() (resp []datatypes.Container_RemoteManagement_Graphs_SensorSpeed, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetServerPowerState() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetServerTemperatureGraphs() (resp []datatypes.Container_RemoteManagement_Graphs_SensorTemperature, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetTransactionHistory() (resp []datatypes.Provisioning_Version1_Transaction_History, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetUpgradeItemPrices() (resp []datatypes.Product_Item_Price, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) ImportVirtualHost() (resp datatypes.Virtual_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) IsPingable() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) Ping() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) PowerCycle() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) PowerOff() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) PowerOn() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) RebootDefault() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) RebootHard() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) RebootSoft() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) RemoveAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) RemoveAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) SetTags(tags *string) (resp bool, err error) {
	params := []interface{}{
		tags,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

func (r *Hardware) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetActiveComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetActiveNetworkMonitorIncident() (resp []datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetAllPowerComponents() (resp []datatypes.Hardware_Power_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetAllowedHost() (resp datatypes.Network_Storage_Allowed_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetAllowedNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetAllowedNetworkStorageReplicas() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetAntivirusSpywareSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetAttributes() (resp []datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetAverageDailyPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetBackendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetBackendRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetBandwidthAllocation() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetBandwidthAllotmentDetail() (resp datatypes.Network_Bandwidth_Version1_Allotment_Detail, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetBenchmarkCertifications() (resp []datatypes.Hardware_Benchmark_Certification, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetBillingItem() (resp datatypes.Billing_Item_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetBillingItemFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetBlockCancelBecauseDisconnectedFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetBusinessContinuanceInsuranceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetContinuousDataProtectionSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetCurrentBillableBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetDatacenter() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetDatacenterName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetDownlinkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetDownlinkNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetDownlinkServers() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetDownlinkVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetDownstreamHardwareBindings() (resp []datatypes.Network_Component_Uplink_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetDownstreamNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetDownstreamNetworkHardwareWithIncidents() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetDownstreamServers() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetDownstreamVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetDriveControllers() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetEvaultNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetFirewallServiceComponent() (resp datatypes.Network_Component_Firewall, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetFixedConfigurationPreset() (resp datatypes.Product_Package_Preset, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetFrontendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetFrontendRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetGlobalIdentifier() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetHardDrives() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetHardwareChassis() (resp datatypes.Hardware_Chassis, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetHardwareFunction() (resp datatypes.Hardware_Function, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetHardwareFunctionDescription() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetHardwareStatus() (resp datatypes.Hardware_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetHasTrustedPlatformModuleBillingItemFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetHostIpsSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetHourlyBillingFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetInboundBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetInboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetLastTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetLatestNetworkMonitorIncident() (resp datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetLocationPathString() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetLockboxNetworkStorage() (resp datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetManagedResourceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetMemory() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetMemoryCapacity() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetMetricTrackingObject() (resp datatypes.Metric_Tracking_Object_HardwareServer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetMonitoringAgents() (resp []datatypes.Monitoring_Agent, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetMonitoringRobot() (resp datatypes.Monitoring_Robot, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetMonitoringServiceComponent() (resp datatypes.Network_Monitor_Version1_Query_Host_Stratum, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetMonitoringServiceEligibilityFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetMonitoringServiceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetMotherboard() (resp datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetNetworkCards() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetNetworkGatewayMember() (resp datatypes.Network_Gateway_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetNetworkGatewayMemberFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetNetworkManagementIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetNetworkMonitorAttachedDownHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetNetworkMonitorAttachedDownVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetNetworkMonitorIncidents() (resp []datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetNetworkMonitors() (resp []datatypes.Network_Monitor_Version1_Query_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetNetworkStatus() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetNetworkStatusAttribute() (resp datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetNetworkVlans() (resp []datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetNextBillingCycleBandwidthAllocation() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetNotesHistory() (resp []datatypes.Hardware_Note, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetOperatingSystem() (resp datatypes.Software_Component_OperatingSystem, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetOperatingSystemReferenceCode() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetOutboundBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetOutboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetPointOfPresenceLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetPowerComponents() (resp []datatypes.Hardware_Power_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetPowerSupply() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetPrimaryBackendIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetPrimaryBackendNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetPrimaryIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetPrimaryNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetPrivateNetworkOnlyFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetProcessorCoreAmount() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetProcessorPhysicalCoreAmount() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetProcessors() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetRack() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetRaidControllers() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetRecentEvents() (resp []datatypes.Notification_Occurrence_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetRemoteManagementAccounts() (resp []datatypes.Hardware_Component_RemoteManagement_User, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetRemoteManagementComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetResourceGroupMemberReferences() (resp []datatypes.Resource_Group_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetResourceGroupRoles() (resp []datatypes.Resource_Group_Role, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetResourceGroups() (resp []datatypes.Resource_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetScaleAssets() (resp []datatypes.Scale_Asset, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetSecurityScanRequests() (resp []datatypes.Network_Security_Scanner_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetServerRoom() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetServiceProvider() (resp datatypes.Service_Provider, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetSoftwareComponents() (resp []datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetSparePoolBillingItem() (resp datatypes.Billing_Item_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetSshKeys() (resp []datatypes.Security_Ssh_Key, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetStorageNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetTagReferences() (resp []datatypes.Tag_Reference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetTopLevelLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetUpgradeRequest() (resp datatypes.Product_Upgrade_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetUplinkHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetUplinkNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetUserData() (resp []datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetVirtualChassis() (resp datatypes.Hardware_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetVirtualChassisSiblings() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetVirtualHost() (resp datatypes.Virtual_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetVirtualLicenses() (resp []datatypes.Software_VirtualLicense, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetVirtualRack() (resp datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetVirtualRackId() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetVirtualRackName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetVirtualizationPlatform() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Attribute struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareAttributeService() Hardware_Attribute {
	return Hardware_Attribute{Session: r}
}

func (r *Hardware_Attribute) GetHardwareAttributeType() (resp datatypes.Hardware_Attribute_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Attribute_Type struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareAttributeTypeService() Hardware_Attribute_Type {
	return Hardware_Attribute_Type{Session: r}
}

type Hardware_Benchmark_Certification struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareBenchmarkCertificationService() Hardware_Benchmark_Certification {
	return Hardware_Benchmark_Certification{Session: r}
}

func (r *Hardware_Benchmark_Certification) GetObject() (resp datatypes.Hardware_Benchmark_Certification, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Benchmark_Certification) GetResultFile() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Hardware_Benchmark_Certification) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Benchmark_Certification) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Chassis struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareChassisService() Hardware_Chassis {
	return Hardware_Chassis{Session: r}
}

func (r *Hardware_Chassis) GetBackplaneCapacity() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Chassis) GetBayCapacity() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Chassis) GetDriveCapacity() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Chassis) GetDriveControllerCapacity() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Chassis) GetGpuCapacity() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Chassis) GetHardwareFunction() (resp datatypes.Hardware_Function, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Chassis) GetPowerCapacity() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Component struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentService() Hardware_Component {
	return Hardware_Component{Session: r}
}

func (r *Hardware_Component) GetCapacity() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component) GetChildren() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component) GetDownlinkHardwareComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component) GetHardwareComponentModel() (resp datatypes.Hardware_Component_Model, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component) GetHardwareComponentType() (resp datatypes.Hardware_Component_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component) GetNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component) GetOwner() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component) GetParent() (resp datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component) GetRaidMode() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component) GetServiceProvider() (resp datatypes.Service_Provider, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component) GetUplinkHardwareComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Component_Attribute struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentAttributeService() Hardware_Component_Attribute {
	return Hardware_Component_Attribute{Session: r}
}

func (r *Hardware_Component_Attribute) GetHardwareComponent() (resp datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_Attribute) GetHardwareComponentAttributeType() (resp datatypes.Hardware_Component_Attribute_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Component_Attribute_Type struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentAttributeTypeService() Hardware_Component_Attribute_Type {
	return Hardware_Component_Attribute_Type{Session: r}
}

type Hardware_Component_DriveController struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentDriveControllerService() Hardware_Component_DriveController {
	return Hardware_Component_DriveController{Session: r}
}

type Hardware_Component_HardDrive struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentHardDriveService() Hardware_Component_HardDrive {
	return Hardware_Component_HardDrive{Session: r}
}

func (r *Hardware_Component_HardDrive) GetPartitions() (resp []datatypes.Hardware_Component_Partition, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Component_Model struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentModelService() Hardware_Component_Model {
	return Hardware_Component_Model{Session: r}
}

func (r *Hardware_Component_Model) GetObject() (resp datatypes.Hardware_Component_Model, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Hardware_Component_Model) GetArchitectureType() (resp datatypes.Hardware_Component_Model_Architecture_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_Model) GetAttributes() (resp []datatypes.Hardware_Component_Model_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_Model) GetCompatibleArrayTypes() (resp []datatypes.Configuration_Storage_Group_Array_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_Model) GetCompatibleChildComponentModels() (resp []datatypes.Hardware_Component_Model, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_Model) GetCompatibleParentComponentModels() (resp []datatypes.Hardware_Component_Model, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_Model) GetHardwareComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_Model) GetHardwareGenericComponentModel() (resp datatypes.Hardware_Component_Model_Generic, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_Model) GetInfinibandCompatibleAttribute() (resp datatypes.Hardware_Component_Model_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_Model) GetIsInfinibandCompatible() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_Model) GetRebootTime() (resp datatypes.Hardware_Component_Motherboard_Reboot_Time, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_Model) GetType() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_Model) GetValidAttributeTypes() (resp []datatypes.Hardware_Component_Model_Attribute_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Component_Model_Architecture_Type struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentModelArchitectureTypeService() Hardware_Component_Model_Architecture_Type {
	return Hardware_Component_Model_Architecture_Type{Session: r}
}

func (r *Hardware_Component_Model_Architecture_Type) GetChildren() (resp []datatypes.Hardware_Component_Model_Architecture_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_Model_Architecture_Type) GetParent() (resp datatypes.Hardware_Component_Model_Architecture_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Component_Model_Attribute struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentModelAttributeService() Hardware_Component_Model_Attribute {
	return Hardware_Component_Model_Attribute{Session: r}
}

func (r *Hardware_Component_Model_Attribute) GetHardwareComponent() (resp datatypes.Hardware_Component_Model, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_Model_Attribute) GetHardwareComponentAttributeType() (resp datatypes.Hardware_Component_Model_Attribute_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Component_Model_Attribute_Type struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentModelAttributeTypeService() Hardware_Component_Model_Attribute_Type {
	return Hardware_Component_Model_Attribute_Type{Session: r}
}

func (r *Hardware_Component_Model_Attribute_Type) GetValidComponentTypes() (resp []datatypes.Hardware_Component_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Component_Model_Generic struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentModelGenericService() Hardware_Component_Model_Generic {
	return Hardware_Component_Model_Generic{Session: r}
}

func (r *Hardware_Component_Model_Generic) GetHardwareComponentModels() (resp []datatypes.Hardware_Component_Model, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_Model_Generic) GetHardwareComponentType() (resp datatypes.Hardware_Component_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_Model_Generic) GetMarketingFeatures() (resp datatypes.Hardware_Component_Model_Generic_MarketingFeature, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Component_Model_Generic_Attribute struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentModelGenericAttributeService() Hardware_Component_Model_Generic_Attribute {
	return Hardware_Component_Model_Generic_Attribute{Session: r}
}

func (r *Hardware_Component_Model_Generic_Attribute) GetHardwareGenericComponentModel() (resp datatypes.Hardware_Component_Model_Generic, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Component_Model_Generic_MarketingFeature struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentModelGenericMarketingFeatureService() Hardware_Component_Model_Generic_MarketingFeature {
	return Hardware_Component_Model_Generic_MarketingFeature{Session: r}
}

func (r *Hardware_Component_Model_Generic_MarketingFeature) GetHardwareGenericComponentModel() (resp datatypes.Hardware_Component_Model_Generic, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Component_Motherboard struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentMotherboardService() Hardware_Component_Motherboard {
	return Hardware_Component_Motherboard{Session: r}
}

type Hardware_Component_Motherboard_Reboot_Time struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentMotherboardRebootTimeService() Hardware_Component_Motherboard_Reboot_Time {
	return Hardware_Component_Motherboard_Reboot_Time{Session: r}
}

func (r *Hardware_Component_Motherboard_Reboot_Time) GetHardwareComponentModel() (resp datatypes.Hardware_Component_Model, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Component_NetworkCard struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentNetworkCardService() Hardware_Component_NetworkCard {
	return Hardware_Component_NetworkCard{Session: r}
}

type Hardware_Component_Partition struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentPartitionService() Hardware_Component_Partition {
	return Hardware_Component_Partition{Session: r}
}

func (r *Hardware_Component_Partition) GetHardwareComponent() (resp datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Component_Partition_OperatingSystem struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentPartitionOperatingSystemService() Hardware_Component_Partition_OperatingSystem {
	return Hardware_Component_Partition_OperatingSystem{Session: r}
}

func (r *Hardware_Component_Partition_OperatingSystem) GetAllObjects() (resp []datatypes.Hardware_Component_Partition_OperatingSystem, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_Partition_OperatingSystem) GetByDescription(description *string) (resp datatypes.Hardware_Component_Partition_OperatingSystem, err error) {
	params := []interface{}{
		description,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_Partition_OperatingSystem) GetObject() (resp datatypes.Hardware_Component_Partition_OperatingSystem, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Hardware_Component_Partition_OperatingSystem) GetPartitionTemplates() (resp []datatypes.Hardware_Component_Partition_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Component_Partition_Template struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentPartitionTemplateService() Hardware_Component_Partition_Template {
	return Hardware_Component_Partition_Template{Session: r}
}

func (r *Hardware_Component_Partition_Template) GetObject() (resp datatypes.Hardware_Component_Partition_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Hardware_Component_Partition_Template) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_Partition_Template) GetData() (resp []datatypes.Hardware_Component_Partition_Template_Partition, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_Partition_Template) GetExpireDate() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_Partition_Template) GetPartitionOperatingSystem() (resp datatypes.Hardware_Component_Partition_OperatingSystem, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_Partition_Template) GetPartitionTemplatePartition() (resp []datatypes.Hardware_Component_Partition_Template_Partition, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Component_Partition_Template_Partition struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentPartitionTemplatePartitionService() Hardware_Component_Partition_Template_Partition {
	return Hardware_Component_Partition_Template_Partition{Session: r}
}

func (r *Hardware_Component_Partition_Template_Partition) GetFilesystemType() (resp datatypes.Configuration_Storage_Filesystem_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_Partition_Template_Partition) GetPartitionTemplate() (resp datatypes.Hardware_Component_Partition_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Component_Processor struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentProcessorService() Hardware_Component_Processor {
	return Hardware_Component_Processor{Session: r}
}

type Hardware_Component_Ram struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentRamService() Hardware_Component_Ram {
	return Hardware_Component_Ram{Session: r}
}

type Hardware_Component_RemoteManagement struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentRemoteManagementService() Hardware_Component_RemoteManagement {
	return Hardware_Component_RemoteManagement{Session: r}
}

func (r *Hardware_Component_RemoteManagement) GetNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Component_RemoteManagement_Command struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentRemoteManagementCommandService() Hardware_Component_RemoteManagement_Command {
	return Hardware_Component_RemoteManagement_Command{Session: r}
}

func (r *Hardware_Component_RemoteManagement_Command) GetRequests() (resp []datatypes.Hardware_Component_RemoteManagement_Command_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Component_RemoteManagement_Command_Request struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentRemoteManagementCommandRequestService() Hardware_Component_RemoteManagement_Command_Request {
	return Hardware_Component_RemoteManagement_Command_Request{Session: r}
}

func (r *Hardware_Component_RemoteManagement_Command_Request) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_RemoteManagement_Command_Request) GetNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_RemoteManagement_Command_Request) GetRemoteManagementCommand() (resp datatypes.Hardware_Component_RemoteManagement_Command, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_RemoteManagement_Command_Request) GetUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Component_RemoteManagement_User struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentRemoteManagementUserService() Hardware_Component_RemoteManagement_User {
	return Hardware_Component_RemoteManagement_User{Session: r}
}

func (r *Hardware_Component_RemoteManagement_User) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_RemoteManagement_User) GetNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Component_SecurityDevice struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentSecurityDeviceService() Hardware_Component_SecurityDevice {
	return Hardware_Component_SecurityDevice{Session: r}
}

type Hardware_Component_SecurityDevice_Infineon struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentSecurityDeviceInfineonService() Hardware_Component_SecurityDevice_Infineon {
	return Hardware_Component_SecurityDevice_Infineon{Session: r}
}

type Hardware_Component_Type struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareComponentTypeService() Hardware_Component_Type {
	return Hardware_Component_Type{Session: r}
}

func (r *Hardware_Component_Type) GetHardwareGenericComponentModels() (resp []datatypes.Hardware_Component_Model_Generic, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Component_Type) GetTypeParent() (resp datatypes.Hardware_Component_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Firewall struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareFirewallService() Hardware_Firewall {
	return Hardware_Firewall{Session: r}
}

func (r *Hardware_Firewall) GetUsers() (resp []datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Function struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareFunctionService() Hardware_Function {
	return Hardware_Function{Session: r}
}

type Hardware_Group struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareGroupService() Hardware_Group {
	return Hardware_Group{Session: r}
}

func (r *Hardware_Group) GetDownlinkServers() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Group) GetDownlinkVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Group) GetDownstreamNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Group) GetDownstreamNetworkHardwareWithIncidents() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Group) GetHardwareChassis() (resp datatypes.Hardware_Chassis, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Group) GetNetworkMonitorAttachedDownHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Group) GetNetworkMonitorAttachedDownVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Group) GetNetworkStatus() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_LoadBalancer struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareLoadBalancerService() Hardware_LoadBalancer {
	return Hardware_LoadBalancer{Session: r}
}

func (r *Hardware_LoadBalancer) GetModelFamily() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_LoadBalancer) GetUsers() (resp []datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Note struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareNoteService() Hardware_Note {
	return Hardware_Note{Session: r}
}

func (r *Hardware_Note) GetEmployee() (resp datatypes.User_Employee, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Note) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Note) GetType() (resp datatypes.Hardware_Note_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Note) GetUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Note_Type struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareNoteTypeService() Hardware_Note_Type {
	return Hardware_Note_Type{Session: r}
}

type Hardware_Power_Component struct {
	Session *Session
	Options
}

func (r *Session) GetHardwarePowerComponentService() Hardware_Power_Component {
	return Hardware_Power_Component{Session: r}
}

func (r *Hardware_Power_Component) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Router struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareRouterService() Hardware_Router {
	return Hardware_Router{Session: r}
}

func (r *Hardware_Router) GetObject() (resp datatypes.Hardware_Router, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Hardware_Router) GetBoundSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetLocalDiskStorageCapabilityFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetSanStorageCapabilityFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Router_Backend struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareRouterBackendService() Hardware_Router_Backend {
	return Hardware_Router_Backend{Session: r}
}

type Hardware_Router_Frontend struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareRouterFrontendService() Hardware_Router_Frontend {
	return Hardware_Router_Frontend{Session: r}
}

type Hardware_SecurityModule struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareSecurityModuleService() Hardware_SecurityModule {
	return Hardware_SecurityModule{Session: r}
}

func (r *Hardware_SecurityModule) CreateObject(templateObject *datatypes.Hardware_SecurityModule) (resp datatypes.Hardware_SecurityModule, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetObject() (resp datatypes.Hardware_SecurityModule, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Server struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareServerService() Hardware_Server {
	return Hardware_Server{Session: r}
}

func (r *Hardware_Server) ActivatePrivatePort() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) ActivatePublicPort() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) BootToRescueLayer(noOsBootEnvironment *string) (resp bool, err error) {
	params := []interface{}{
		noOsBootEnvironment,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
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
func (r *Hardware_Server) CreateObject(templateObject *datatypes.Hardware_Server) (resp datatypes.Hardware_Server, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) CreatePostSoftwareInstallTransaction(installCodes []string, returnBoolean *bool) (resp bool, err error) {
	params := []interface{}{
		installCodes,
		returnBoolean,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) EditObject(templateObject *datatypes.Hardware_Server) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetBackendBandwidthUsage(startDate *time.Time, endDate *time.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetBackendBandwidthUse(startDate *time.Time, endDate *time.Time) (resp []datatypes.Network_Bandwidth_Version1_Usage_Detail, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetBandwidthForDateRange(startDate *time.Time, endDate *time.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetBandwidthImage(networkType *string, snapshotRange *string, draw *bool, dateSpecified *time.Time, dateSpecifiedEnd *time.Time) (resp datatypes.Container_Bandwidth_GraphOutputs, err error) {
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
func (r *Hardware_Server) GetCurrentBenchmarkCertificationResultFile() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetCustomBandwidthDataByDate(graphData *datatypes.Container_Graph) (resp datatypes.Container_Graph, err error) {
	params := []interface{}{
		graphData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetFirewallProtectableSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetFrontendBandwidthUsage(startDate *time.Time, endDate *time.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetFrontendBandwidthUse(startDate *time.Time, endDate *time.Time) (resp []datatypes.Network_Bandwidth_Version1_Usage_Detail, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetHardwareByIpAddress(ipAddress *string) (resp datatypes.Hardware_Server, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetItemPricesFromSoftwareDescriptions(softwareDescriptions []datatypes.Software_Description, includeTranslationsFlag *bool, returnAllPricesFlag *bool) (resp []datatypes.Product_Item, err error) {
	params := []interface{}{
		softwareDescriptions,
		includeTranslationsFlag,
		returnAllPricesFlag,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetManagementNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetNetworkComponentFirewallProtectableIpAddresses() (resp []datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetObject() (resp datatypes.Hardware_Server, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPMInfo() (resp []datatypes.Container_RemoteManagement_PmInfo, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPrimaryDriveSize() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPrivateBandwidthDataSummary() (resp datatypes.Container_Network_Bandwidth_Data_Summary, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPrivateBandwidthGraphImage(startTime *string, endTime *string) (resp []byte, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPrivateNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPrivateVlan() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPrivateVlanByIpAddress(ipAddress *string) (resp datatypes.Network_Vlan, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetProvisionDate() (resp time.Time, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPublicBandwidthDataSummary() (resp datatypes.Container_Network_Bandwidth_Data_Summary, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPublicBandwidthGraphImage(startTime *time.Time, endTime *time.Time) (resp []byte, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPublicBandwidthTotal(startTime *int, endTime *int) (resp uint, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPublicNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPublicVlan() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPublicVlanByHostname(hostname *string) (resp datatypes.Network_Vlan, err error) {
	params := []interface{}{
		hostname,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetReverseDomainRecords() (resp []datatypes.Dns_Domain, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetSensorData() (resp []datatypes.Container_RemoteManagement_SensorReading, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetSensorDataWithGraphs() (resp datatypes.Container_RemoteManagement_SensorReadingsWithGraphs, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetServerDetails() (resp datatypes.Container_Hardware_Server_Details, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetServerFanSpeedGraphs() (resp []datatypes.Container_RemoteManagement_Graphs_SensorSpeed, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetServerPowerState() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetServerTemperatureGraphs() (resp []datatypes.Container_RemoteManagement_Graphs_SensorTemperature, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetValidBlockDeviceTemplateGroups(visibility *string) (resp []datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	params := []interface{}{
		visibility,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetWindowsUpdateAvailableUpdates() (resp []datatypes.Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetWindowsUpdateInstalledUpdates() (resp []datatypes.Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetWindowsUpdateStatus() (resp datatypes.Container_Utility_Microsoft_Windows_UpdateServices_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) InitiateIderaBareMetalRestore() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) InitiateR1SoftBareMetalRestore() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) IsBackendPingable() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) IsPingable() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) IsWindowsServer() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) Ping() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) PowerCycle() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) PowerOff() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) PowerOn() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) RebootDefault() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) RebootHard() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) RebootSoft() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) ReloadCurrentOperatingSystemConfiguration(token *string) (resp string, err error) {
	params := []interface{}{
		token,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) ReloadOperatingSystem(token *string, config *datatypes.Container_Hardware_Server_Configuration) (resp string, err error) {
	params := []interface{}{
		token,
		config,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) RunPassmarkCertificationBenchmark() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) SetOperatingSystemPassword(newPassword *string) (resp bool, err error) {
	params := []interface{}{
		newPassword,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) SetPrivateNetworkInterfaceSpeed(newSpeed *int) (resp bool, err error) {
	params := []interface{}{
		newSpeed,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) SetPublicNetworkInterfaceSpeed(newSpeed *int) (resp bool, err error) {
	params := []interface{}{
		newSpeed,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) SetUserMetadata(metadata []string) (resp []datatypes.Hardware_Attribute, err error) {
	params := []interface{}{
		metadata,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) ShutdownPrivatePort() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) ShutdownPublicPort() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) SparePool(action *string, newOrder *bool) (resp bool, err error) {
	params := []interface{}{
		action,
		newOrder,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) ValidatePartitionsForOperatingSystem(operatingSystem *datatypes.Software_Description, partitions []datatypes.Hardware_Component_Partition) (resp bool, err error) {
	params := []interface{}{
		operatingSystem,
		partitions,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

func (r *Hardware_Server) GetActiveNetworkFirewallBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetActiveTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetActiveTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetActiveTransactions() (resp []datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetAvailableMonitoring() (resp []datatypes.Network_Monitor_Version1_Query_Host_Stratum, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetAverageDailyBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetAverageDailyPrivateBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetBillingCycleBandwidthUsage() (resp []datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetBillingCyclePrivateBandwidthUsage() (resp datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetBillingCyclePublicBandwidthUsage() (resp datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetContainsSolidStateDrivesFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetControlPanel() (resp datatypes.Software_Component_ControlPanel, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetCost() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetCurrentBandwidthSummary() (resp datatypes.Metric_Tracking_Object_Bandwidth_Summary, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetCustomerInstalledOperatingSystemFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetCustomerOwnedFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetInboundPrivateBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetLastOperatingSystemReload() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetMetricTrackingObjectId() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetMonitoringUserNotification() (resp []datatypes.User_Customer_Notification_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetOpenCancellationTicket() (resp datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetOutboundPrivateBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetOverBandwidthAllocationFlag() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPrivateIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetProjectedOverBandwidthAllocationFlag() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetProjectedPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetRecentRemoteManagementCommands() (resp []datatypes.Hardware_Component_RemoteManagement_Command_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetRegionalInternetRegistry() (resp datatypes.Network_Regional_Internet_Registry, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetRemoteManagement() (resp datatypes.Hardware_Component_RemoteManagement, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetRemoteManagementUsers() (resp []datatypes.Hardware_Component_RemoteManagement_User, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetStatisticsRemoteManagement() (resp datatypes.Hardware_Component_RemoteManagement, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetUsers() (resp []datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Hardware_Status struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareStatusService() Hardware_Status {
	return Hardware_Status{Session: r}
}

type Hardware_Switch struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareSwitchService() Hardware_Switch {
	return Hardware_Switch{Session: r}
}
