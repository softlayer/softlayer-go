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

import "github.ibm.com/riethm/gopherlayer/datatypes"

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
func (r *Hardware) GetAlarmHistory(startDate *datatypes.Time, endDate *datatypes.Time, alarmId *string) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
		alarmId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
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
func (r *Hardware) GetAttachedNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetAttributes() (resp []datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetAvailableNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetAverageDailyPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetBackendIncomingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetBackendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetBackendOutgoingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
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
func (r *Hardware) GetCreateObjectOptions() (resp datatypes.Container_Hardware_Configuration, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetCurrentBillableBandwidthUsage() (resp float64, err error) {
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
func (r *Hardware) GetDailyAverage(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
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
func (r *Hardware) GetFrontendIncomingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetFrontendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetFrontendOutgoingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
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
func (r *Hardware) GetHourlyBandwidth(mode *string, day *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		mode,
		day,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
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
func (r *Hardware) GetMonitoringActiveAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetMonitoringAgents() (resp []datatypes.Monitoring_Agent, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetMonitoringClosedAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
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
func (r *Hardware) GetObject() (resp datatypes.Hardware, err error) {
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
func (r *Hardware) GetPrivateBandwidthData(startTime *int, endTime *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
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
func (r *Hardware) GetPublicBandwidthData(startTime *int, endTime *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
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
func (r *Hardware) GetServerRoom() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetServerTemperatureGraphs() (resp []datatypes.Container_RemoteManagement_Graphs_SensorTemperature, err error) {
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
func (r *Hardware) GetTransactionHistory() (resp []datatypes.Provisioning_Version1_Transaction_History, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware) GetUpgradeItemPrices() (resp []datatypes.Product_Item_Price, err error) {
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

type Hardware_Benchmark_Certification struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareBenchmarkCertificationService() Hardware_Benchmark_Certification {
	return Hardware_Benchmark_Certification{Session: r}
}

func (r *Hardware_Benchmark_Certification) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Benchmark_Certification) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Benchmark_Certification) GetObject() (resp datatypes.Hardware_Benchmark_Certification, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Benchmark_Certification) GetResultFile() (resp []byte, err error) {
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
func (r *Hardware_Component_Model) GetObject() (resp datatypes.Hardware_Component_Model, err error) {
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
func (r *Hardware_Component_Partition_Template) GetObject() (resp datatypes.Hardware_Component_Partition_Template, err error) {
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

type Hardware_Router struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareRouterService() Hardware_Router {
	return Hardware_Router{Session: r}
}

func (r *Hardware_Router) AllowAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) AllowAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) CaptureImage(captureTemplate *datatypes.Container_Disk_Image_Capture_Template) (resp datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	params := []interface{}{
		captureTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) CloseAlarm(alarmId *string) (resp bool, err error) {
	params := []interface{}{
		alarmId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) CreateObject(templateObject *datatypes.Hardware) (resp datatypes.Hardware, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) DeleteSoftwareComponentPasswords(softwareComponentPasswords []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		softwareComponentPasswords,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) EditSoftwareComponentPasswords(softwareComponentPasswords []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		softwareComponentPasswords,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) ExecuteRemoteScript(uri *string) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		uri,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) FindByIpAddress(ipAddress *string) (resp datatypes.Hardware, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GenerateOrderTemplate(templateObject *datatypes.Hardware) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetActiveComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetActiveNetworkMonitorIncident() (resp []datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetAlarmHistory(startDate *datatypes.Time, endDate *datatypes.Time, alarmId *string) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
		alarmId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetAllPowerComponents() (resp []datatypes.Hardware_Power_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetAllowedHost() (resp datatypes.Network_Storage_Allowed_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetAllowedNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetAllowedNetworkStorageReplicas() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetAntivirusSpywareSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetAttachedNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetAttributes() (resp []datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetAvailableNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetAverageDailyPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetBackendIncomingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetBackendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetBackendOutgoingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetBackendRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetBandwidthAllocation() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetBandwidthAllotmentDetail() (resp datatypes.Network_Bandwidth_Version1_Allotment_Detail, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetBenchmarkCertifications() (resp []datatypes.Hardware_Benchmark_Certification, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetBillingItem() (resp datatypes.Billing_Item_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetBillingItemFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetBlockCancelBecauseDisconnectedFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetBoundSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetBusinessContinuanceInsuranceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetContinuousDataProtectionSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetCreateObjectOptions() (resp datatypes.Container_Hardware_Configuration, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetCurrentBillableBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetCurrentBillingDetail() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetCurrentBillingTotal() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetDailyAverage(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetDatacenter() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetDatacenterName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetDownlinkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetDownlinkNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetDownlinkServers() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetDownlinkVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetDownstreamHardwareBindings() (resp []datatypes.Network_Component_Uplink_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetDownstreamNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetDownstreamNetworkHardwareWithIncidents() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetDownstreamServers() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetDownstreamVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetDriveControllers() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetEvaultNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetFirewallServiceComponent() (resp datatypes.Network_Component_Firewall, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetFixedConfigurationPreset() (resp datatypes.Product_Package_Preset, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetFrontendIncomingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetFrontendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetFrontendOutgoingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetFrontendRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetGlobalIdentifier() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetHardDrives() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetHardwareChassis() (resp datatypes.Hardware_Chassis, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetHardwareFunction() (resp datatypes.Hardware_Function, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetHardwareFunctionDescription() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetHardwareStatus() (resp datatypes.Hardware_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetHasTrustedPlatformModuleBillingItemFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetHostIpsSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetHourlyBandwidth(mode *string, day *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		mode,
		day,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetHourlyBillingFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetInboundBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetInboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetLastTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetLatestNetworkMonitorIncident() (resp datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetLocalDiskStorageCapabilityFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetLocationPathString() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetLockboxNetworkStorage() (resp datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetManagedResourceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetMemory() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetMemoryCapacity() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetMetricTrackingObject() (resp datatypes.Metric_Tracking_Object_HardwareServer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetMonitoringActiveAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetMonitoringAgents() (resp []datatypes.Monitoring_Agent, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetMonitoringClosedAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetMonitoringRobot() (resp datatypes.Monitoring_Robot, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetMonitoringServiceComponent() (resp datatypes.Network_Monitor_Version1_Query_Host_Stratum, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetMonitoringServiceEligibilityFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetMonitoringServiceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetMotherboard() (resp datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetNetworkCards() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetNetworkGatewayMember() (resp datatypes.Network_Gateway_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetNetworkGatewayMemberFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetNetworkManagementIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetNetworkMonitorAttachedDownHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetNetworkMonitorAttachedDownVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetNetworkMonitorIncidents() (resp []datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetNetworkMonitors() (resp []datatypes.Network_Monitor_Version1_Query_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetNetworkStatus() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetNetworkStatusAttribute() (resp datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetNetworkVlans() (resp []datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetNextBillingCycleBandwidthAllocation() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetNotesHistory() (resp []datatypes.Hardware_Note, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetObject() (resp datatypes.Hardware_Router, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetOperatingSystem() (resp datatypes.Software_Component_OperatingSystem, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetOperatingSystemReferenceCode() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetOutboundBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetOutboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetPointOfPresenceLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetPowerComponents() (resp []datatypes.Hardware_Power_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetPowerSupply() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetPrimaryBackendIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetPrimaryBackendNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetPrimaryIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetPrimaryNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetPrivateBandwidthData(startTime *int, endTime *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetPrivateNetworkOnlyFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetProcessorCoreAmount() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetProcessorPhysicalCoreAmount() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetProcessors() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetPublicBandwidthData(startTime *int, endTime *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetRack() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetRaidControllers() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetRecentEvents() (resp []datatypes.Notification_Occurrence_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetRemoteManagementAccounts() (resp []datatypes.Hardware_Component_RemoteManagement_User, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetRemoteManagementComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetResourceGroupMemberReferences() (resp []datatypes.Resource_Group_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetResourceGroupRoles() (resp []datatypes.Resource_Group_Role, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetResourceGroups() (resp []datatypes.Resource_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetSanStorageCapabilityFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetScaleAssets() (resp []datatypes.Scale_Asset, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetSecurityScanRequests() (resp []datatypes.Network_Security_Scanner_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetSensorData() (resp []datatypes.Container_RemoteManagement_SensorReading, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetSensorDataWithGraphs() (resp datatypes.Container_RemoteManagement_SensorReadingsWithGraphs, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetServerFanSpeedGraphs() (resp []datatypes.Container_RemoteManagement_Graphs_SensorSpeed, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetServerPowerState() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetServerRoom() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetServerTemperatureGraphs() (resp []datatypes.Container_RemoteManagement_Graphs_SensorTemperature, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetServiceProvider() (resp datatypes.Service_Provider, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetSoftwareComponents() (resp []datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetSparePoolBillingItem() (resp datatypes.Billing_Item_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetSshKeys() (resp []datatypes.Security_Ssh_Key, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetStorageNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetTagReferences() (resp []datatypes.Tag_Reference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetTopLevelLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetTransactionHistory() (resp []datatypes.Provisioning_Version1_Transaction_History, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetUpgradeItemPrices() (resp []datatypes.Product_Item_Price, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetUpgradeRequest() (resp datatypes.Product_Upgrade_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetUplinkHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetUplinkNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetUserData() (resp []datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetVirtualChassis() (resp datatypes.Hardware_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetVirtualChassisSiblings() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetVirtualHost() (resp datatypes.Virtual_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetVirtualLicenses() (resp []datatypes.Software_VirtualLicense, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetVirtualRack() (resp datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetVirtualRackId() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetVirtualRackName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) GetVirtualizationPlatform() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) ImportVirtualHost() (resp datatypes.Virtual_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) IsPingable() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) Ping() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) PowerCycle() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) PowerOff() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) PowerOn() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) RebootDefault() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) RebootHard() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) RebootSoft() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) RemoveAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) RemoveAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Router) SetTags(tags *string) (resp bool, err error) {
	params := []interface{}{
		tags,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Hardware_SecurityModule struct {
	Session *Session
	Options
}

func (r *Session) GetHardwareSecurityModuleService() Hardware_SecurityModule {
	return Hardware_SecurityModule{Session: r}
}

func (r *Hardware_SecurityModule) ActivatePrivatePort() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) ActivatePublicPort() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) AllowAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) AllowAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) BootToRescueLayer(noOsBootEnvironment *string) (resp bool, err error) {
	params := []interface{}{
		noOsBootEnvironment,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) CaptureImage(captureTemplate *datatypes.Container_Disk_Image_Capture_Template) (resp datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	params := []interface{}{
		captureTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) CloseAlarm(alarmId *string) (resp bool, err error) {
	params := []interface{}{
		alarmId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
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
func (r *Hardware_SecurityModule) CreateObject(templateObject *datatypes.Hardware_SecurityModule) (resp datatypes.Hardware_SecurityModule, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) CreatePostSoftwareInstallTransaction(installCodes []string, returnBoolean *bool) (resp bool, err error) {
	params := []interface{}{
		installCodes,
		returnBoolean,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) DeleteSoftwareComponentPasswords(softwareComponentPasswords []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		softwareComponentPasswords,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) EditObject(templateObject *datatypes.Hardware_Server) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) EditSoftwareComponentPasswords(softwareComponentPasswords []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		softwareComponentPasswords,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) ExecuteRemoteScript(uri *string) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		uri,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) FindByIpAddress(ipAddress *string) (resp datatypes.Hardware, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GenerateOrderTemplate(templateObject *datatypes.Hardware) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetActiveComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetActiveNetworkFirewallBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetActiveNetworkMonitorIncident() (resp []datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetActiveTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetActiveTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetActiveTransactions() (resp []datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetAlarmHistory(startDate *datatypes.Time, endDate *datatypes.Time, alarmId *string) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
		alarmId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetAllPowerComponents() (resp []datatypes.Hardware_Power_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetAllowedHost() (resp datatypes.Network_Storage_Allowed_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetAllowedNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetAllowedNetworkStorageReplicas() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetAntivirusSpywareSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetAttachedNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetAttributes() (resp []datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetAvailableMonitoring() (resp []datatypes.Network_Monitor_Version1_Query_Host_Stratum, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetAvailableNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetAverageDailyBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetAverageDailyPrivateBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetAverageDailyPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetBackendBandwidthUsage(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetBackendBandwidthUse(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Network_Bandwidth_Version1_Usage_Detail, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetBackendIncomingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetBackendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetBackendOutgoingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetBackendRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetBandwidthAllocation() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetBandwidthAllotmentDetail() (resp datatypes.Network_Bandwidth_Version1_Allotment_Detail, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetBandwidthForDateRange(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
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
func (r *Hardware_SecurityModule) GetBenchmarkCertifications() (resp []datatypes.Hardware_Benchmark_Certification, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetBillingCycleBandwidthUsage() (resp []datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetBillingCyclePrivateBandwidthUsage() (resp datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetBillingCyclePublicBandwidthUsage() (resp datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetBillingItem() (resp datatypes.Billing_Item_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetBillingItemFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetBlockCancelBecauseDisconnectedFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetBusinessContinuanceInsuranceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetContainsSolidStateDrivesFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetContinuousDataProtectionSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetControlPanel() (resp datatypes.Software_Component_ControlPanel, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetCost() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetCreateObjectOptions() (resp datatypes.Container_Hardware_Configuration, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetCurrentBandwidthSummary() (resp datatypes.Metric_Tracking_Object_Bandwidth_Summary, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetCurrentBenchmarkCertificationResultFile() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetCurrentBillableBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetCurrentBillingDetail() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetCurrentBillingTotal() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetCustomBandwidthDataByDate(graphData *datatypes.Container_Graph) (resp datatypes.Container_Graph, err error) {
	params := []interface{}{
		graphData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetCustomerInstalledOperatingSystemFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetCustomerOwnedFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetDailyAverage(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetDatacenter() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetDatacenterName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetDownlinkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetDownlinkNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetDownlinkServers() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetDownlinkVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetDownstreamHardwareBindings() (resp []datatypes.Network_Component_Uplink_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetDownstreamNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetDownstreamNetworkHardwareWithIncidents() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetDownstreamServers() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetDownstreamVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetDriveControllers() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetEvaultNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetFirewallProtectableSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetFirewallServiceComponent() (resp datatypes.Network_Component_Firewall, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetFixedConfigurationPreset() (resp datatypes.Product_Package_Preset, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetFrontendBandwidthUsage(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetFrontendBandwidthUse(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Network_Bandwidth_Version1_Usage_Detail, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetFrontendIncomingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetFrontendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetFrontendOutgoingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetFrontendRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetGlobalIdentifier() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetHardDrives() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetHardwareByIpAddress(ipAddress *string) (resp datatypes.Hardware_Server, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetHardwareChassis() (resp datatypes.Hardware_Chassis, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetHardwareFunction() (resp datatypes.Hardware_Function, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetHardwareFunctionDescription() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetHardwareStatus() (resp datatypes.Hardware_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetHasTrustedPlatformModuleBillingItemFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetHostIpsSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetHourlyBandwidth(mode *string, day *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		mode,
		day,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetHourlyBillingFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetInboundBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetInboundPrivateBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetInboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetItemPricesFromSoftwareDescriptions(softwareDescriptions []datatypes.Software_Description, includeTranslationsFlag *bool, returnAllPricesFlag *bool) (resp []datatypes.Product_Item, err error) {
	params := []interface{}{
		softwareDescriptions,
		includeTranslationsFlag,
		returnAllPricesFlag,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetLastOperatingSystemReload() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetLastTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetLatestNetworkMonitorIncident() (resp datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetLocationPathString() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetLockboxNetworkStorage() (resp datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetManagedResourceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetManagementNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetMemory() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetMemoryCapacity() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetMetricTrackingObject() (resp datatypes.Metric_Tracking_Object_HardwareServer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetMetricTrackingObjectId() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetMonitoringActiveAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetMonitoringAgents() (resp []datatypes.Monitoring_Agent, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetMonitoringClosedAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetMonitoringRobot() (resp datatypes.Monitoring_Robot, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetMonitoringServiceComponent() (resp datatypes.Network_Monitor_Version1_Query_Host_Stratum, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetMonitoringServiceEligibilityFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetMonitoringServiceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetMonitoringUserNotification() (resp []datatypes.User_Customer_Notification_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetMotherboard() (resp datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetNetworkCards() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetNetworkComponentFirewallProtectableIpAddresses() (resp []datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetNetworkGatewayMember() (resp datatypes.Network_Gateway_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetNetworkGatewayMemberFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetNetworkManagementIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetNetworkMonitorAttachedDownHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetNetworkMonitorAttachedDownVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetNetworkMonitorIncidents() (resp []datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetNetworkMonitors() (resp []datatypes.Network_Monitor_Version1_Query_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetNetworkStatus() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetNetworkStatusAttribute() (resp datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetNetworkVlans() (resp []datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetNextBillingCycleBandwidthAllocation() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetNotesHistory() (resp []datatypes.Hardware_Note, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetObject() (resp datatypes.Hardware_SecurityModule, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetOpenCancellationTicket() (resp datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetOperatingSystem() (resp datatypes.Software_Component_OperatingSystem, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetOperatingSystemReferenceCode() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetOutboundBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetOutboundPrivateBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetOutboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetOverBandwidthAllocationFlag() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetPMInfo() (resp []datatypes.Container_RemoteManagement_PmInfo, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetPointOfPresenceLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetPowerComponents() (resp []datatypes.Hardware_Power_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetPowerSupply() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetPrimaryBackendIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetPrimaryBackendNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetPrimaryDriveSize() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetPrimaryIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetPrimaryNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetPrivateBandwidthData(startTime *int, endTime *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetPrivateBandwidthDataSummary() (resp datatypes.Container_Network_Bandwidth_Data_Summary, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetPrivateBandwidthGraphImage(startTime *string, endTime *string) (resp []byte, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetPrivateIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetPrivateNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetPrivateNetworkOnlyFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetPrivateVlan() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetPrivateVlanByIpAddress(ipAddress *string) (resp datatypes.Network_Vlan, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetProcessorCoreAmount() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetProcessorPhysicalCoreAmount() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetProcessors() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetProjectedOverBandwidthAllocationFlag() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetProjectedPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetProvisionDate() (resp datatypes.Time, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetPublicBandwidthData(startTime *int, endTime *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetPublicBandwidthDataSummary() (resp datatypes.Container_Network_Bandwidth_Data_Summary, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetPublicBandwidthGraphImage(startTime *datatypes.Time, endTime *datatypes.Time) (resp []byte, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetPublicBandwidthTotal(startTime *int, endTime *int) (resp uint, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetPublicNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetPublicVlan() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetPublicVlanByHostname(hostname *string) (resp datatypes.Network_Vlan, err error) {
	params := []interface{}{
		hostname,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetRack() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetRaidControllers() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetRecentEvents() (resp []datatypes.Notification_Occurrence_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetRecentRemoteManagementCommands() (resp []datatypes.Hardware_Component_RemoteManagement_Command_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetRegionalInternetRegistry() (resp datatypes.Network_Regional_Internet_Registry, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetRemoteManagement() (resp datatypes.Hardware_Component_RemoteManagement, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetRemoteManagementAccounts() (resp []datatypes.Hardware_Component_RemoteManagement_User, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetRemoteManagementComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetRemoteManagementUsers() (resp []datatypes.Hardware_Component_RemoteManagement_User, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetResourceGroupMemberReferences() (resp []datatypes.Resource_Group_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetResourceGroupRoles() (resp []datatypes.Resource_Group_Role, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetResourceGroups() (resp []datatypes.Resource_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetReverseDomainRecords() (resp []datatypes.Dns_Domain, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetScaleAssets() (resp []datatypes.Scale_Asset, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetSecurityScanRequests() (resp []datatypes.Network_Security_Scanner_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetSensorData() (resp []datatypes.Container_RemoteManagement_SensorReading, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetSensorDataWithGraphs() (resp datatypes.Container_RemoteManagement_SensorReadingsWithGraphs, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetServerDetails() (resp datatypes.Container_Hardware_Server_Details, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetServerFanSpeedGraphs() (resp []datatypes.Container_RemoteManagement_Graphs_SensorSpeed, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetServerPowerState() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetServerRoom() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetServerTemperatureGraphs() (resp []datatypes.Container_RemoteManagement_Graphs_SensorTemperature, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetServiceProvider() (resp datatypes.Service_Provider, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetSoftwareComponents() (resp []datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetSparePoolBillingItem() (resp datatypes.Billing_Item_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetSshKeys() (resp []datatypes.Security_Ssh_Key, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetStatisticsRemoteManagement() (resp datatypes.Hardware_Component_RemoteManagement, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetStorageNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetTagReferences() (resp []datatypes.Tag_Reference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetTopLevelLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetTransactionHistory() (resp []datatypes.Provisioning_Version1_Transaction_History, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetUpgradeItemPrices() (resp []datatypes.Product_Item_Price, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetUpgradeRequest() (resp datatypes.Product_Upgrade_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetUplinkHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetUplinkNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetUserData() (resp []datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetUsers() (resp []datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetValidBlockDeviceTemplateGroups(visibility *string) (resp []datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	params := []interface{}{
		visibility,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetVirtualChassis() (resp datatypes.Hardware_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetVirtualChassisSiblings() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetVirtualHost() (resp datatypes.Virtual_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetVirtualLicenses() (resp []datatypes.Software_VirtualLicense, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetVirtualRack() (resp datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetVirtualRackId() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetVirtualRackName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetVirtualizationPlatform() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetWindowsUpdateAvailableUpdates() (resp []datatypes.Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetWindowsUpdateInstalledUpdates() (resp []datatypes.Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) GetWindowsUpdateStatus() (resp datatypes.Container_Utility_Microsoft_Windows_UpdateServices_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) ImportVirtualHost() (resp datatypes.Virtual_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) InitiateIderaBareMetalRestore() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) InitiateR1SoftBareMetalRestore() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) IsBackendPingable() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) IsPingable() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) IsWindowsServer() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) Ping() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) PowerCycle() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) PowerOff() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) PowerOn() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) RebootDefault() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) RebootHard() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) RebootSoft() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) ReloadCurrentOperatingSystemConfiguration(token *string) (resp string, err error) {
	params := []interface{}{
		token,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) ReloadOperatingSystem(token *string, config *datatypes.Container_Hardware_Server_Configuration) (resp string, err error) {
	params := []interface{}{
		token,
		config,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) RemoveAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) RemoveAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) RunPassmarkCertificationBenchmark() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) SetOperatingSystemPassword(newPassword *string) (resp bool, err error) {
	params := []interface{}{
		newPassword,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) SetPrivateNetworkInterfaceSpeed(newSpeed *int) (resp bool, err error) {
	params := []interface{}{
		newSpeed,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) SetPublicNetworkInterfaceSpeed(newSpeed *int) (resp bool, err error) {
	params := []interface{}{
		newSpeed,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) SetTags(tags *string) (resp bool, err error) {
	params := []interface{}{
		tags,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) SetUserMetadata(metadata []string) (resp []datatypes.Hardware_Attribute, err error) {
	params := []interface{}{
		metadata,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) ShutdownPrivatePort() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) ShutdownPublicPort() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) SparePool(action *string, newOrder *bool) (resp bool, err error) {
	params := []interface{}{
		action,
		newOrder,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_SecurityModule) ValidatePartitionsForOperatingSystem(operatingSystem *datatypes.Software_Description, partitions []datatypes.Hardware_Component_Partition) (resp bool, err error) {
	params := []interface{}{
		operatingSystem,
		partitions,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
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
func (r *Hardware_Server) AllowAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) AllowAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) BootToRescueLayer(noOsBootEnvironment *string) (resp bool, err error) {
	params := []interface{}{
		noOsBootEnvironment,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) CaptureImage(captureTemplate *datatypes.Container_Disk_Image_Capture_Template) (resp datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	params := []interface{}{
		captureTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) CloseAlarm(alarmId *string) (resp bool, err error) {
	params := []interface{}{
		alarmId,
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
func (r *Hardware_Server) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) DeleteSoftwareComponentPasswords(softwareComponentPasswords []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		softwareComponentPasswords,
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
func (r *Hardware_Server) EditSoftwareComponentPasswords(softwareComponentPasswords []datatypes.Software_Component_Password) (resp bool, err error) {
	params := []interface{}{
		softwareComponentPasswords,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) ExecuteRemoteScript(uri *string) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		uri,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) FindByIpAddress(ipAddress *string) (resp datatypes.Hardware, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GenerateOrderTemplate(templateObject *datatypes.Hardware) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetActiveComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetActiveNetworkFirewallBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetActiveNetworkMonitorIncident() (resp []datatypes.Network_Monitor_Version1_Incident, err error) {
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
func (r *Hardware_Server) GetAlarmHistory(startDate *datatypes.Time, endDate *datatypes.Time, alarmId *string) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
		alarmId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetAllPowerComponents() (resp []datatypes.Hardware_Power_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetAllowedHost() (resp datatypes.Network_Storage_Allowed_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetAllowedNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetAllowedNetworkStorageReplicas() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetAntivirusSpywareSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetAttachedNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetAttributes() (resp []datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetAvailableMonitoring() (resp []datatypes.Network_Monitor_Version1_Query_Host_Stratum, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetAvailableNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
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
func (r *Hardware_Server) GetAverageDailyPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetBackendBandwidthUsage(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetBackendBandwidthUse(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Network_Bandwidth_Version1_Usage_Detail, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetBackendIncomingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetBackendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetBackendOutgoingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetBackendRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetBandwidthAllocation() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetBandwidthAllotmentDetail() (resp datatypes.Network_Bandwidth_Version1_Allotment_Detail, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetBandwidthForDateRange(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
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
func (r *Hardware_Server) GetBenchmarkCertifications() (resp []datatypes.Hardware_Benchmark_Certification, err error) {
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
func (r *Hardware_Server) GetBillingItem() (resp datatypes.Billing_Item_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetBillingItemFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetBlockCancelBecauseDisconnectedFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetBusinessContinuanceInsuranceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetComponents() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetContainsSolidStateDrivesFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetContinuousDataProtectionSoftwareComponent() (resp datatypes.Software_Component, err error) {
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
func (r *Hardware_Server) GetCreateObjectOptions() (resp datatypes.Container_Hardware_Configuration, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetCurrentBandwidthSummary() (resp datatypes.Metric_Tracking_Object_Bandwidth_Summary, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetCurrentBenchmarkCertificationResultFile() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetCurrentBillableBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetCurrentBillingDetail() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetCurrentBillingTotal() (resp float64, err error) {
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
func (r *Hardware_Server) GetCustomerInstalledOperatingSystemFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetCustomerOwnedFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetDailyAverage(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetDatacenter() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetDatacenterName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetDownlinkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetDownlinkNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetDownlinkServers() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetDownlinkVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetDownstreamHardwareBindings() (resp []datatypes.Network_Component_Uplink_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetDownstreamNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetDownstreamNetworkHardwareWithIncidents() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetDownstreamServers() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetDownstreamVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetDriveControllers() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetEvaultNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetFirewallProtectableSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetFirewallServiceComponent() (resp datatypes.Network_Component_Firewall, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetFixedConfigurationPreset() (resp datatypes.Product_Package_Preset, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetFrontendBandwidthUsage(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetFrontendBandwidthUse(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Network_Bandwidth_Version1_Usage_Detail, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetFrontendIncomingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetFrontendNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetFrontendOutgoingBandwidth(startDate *datatypes.Time, endDate *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetFrontendRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetGlobalIdentifier() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetHardDrives() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetHardwareByIpAddress(ipAddress *string) (resp datatypes.Hardware_Server, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetHardwareChassis() (resp datatypes.Hardware_Chassis, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetHardwareFunction() (resp datatypes.Hardware_Function, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetHardwareFunctionDescription() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetHardwareStatus() (resp datatypes.Hardware_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetHasTrustedPlatformModuleBillingItemFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetHostIpsSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetHourlyBandwidth(mode *string, day *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		mode,
		day,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetHourlyBillingFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetInboundBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetInboundPrivateBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetInboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
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
func (r *Hardware_Server) GetLastOperatingSystemReload() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetLastTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetLatestNetworkMonitorIncident() (resp datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetLocationPathString() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetLockboxNetworkStorage() (resp datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetManagedResourceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetManagementNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetMemory() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetMemoryCapacity() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetMetricTrackingObject() (resp datatypes.Metric_Tracking_Object_HardwareServer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetMetricTrackingObjectId() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetMonitoringActiveAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetMonitoringAgents() (resp []datatypes.Monitoring_Agent, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetMonitoringClosedAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetMonitoringRobot() (resp datatypes.Monitoring_Robot, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetMonitoringServiceComponent() (resp datatypes.Network_Monitor_Version1_Query_Host_Stratum, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetMonitoringServiceEligibilityFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetMonitoringServiceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetMonitoringUserNotification() (resp []datatypes.User_Customer_Notification_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetMotherboard() (resp datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetNetworkCards() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetNetworkComponentFirewallProtectableIpAddresses() (resp []datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetNetworkGatewayMember() (resp datatypes.Network_Gateway_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetNetworkGatewayMemberFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetNetworkManagementIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetNetworkMonitorAttachedDownHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetNetworkMonitorAttachedDownVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetNetworkMonitorIncidents() (resp []datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetNetworkMonitors() (resp []datatypes.Network_Monitor_Version1_Query_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetNetworkStatus() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetNetworkStatusAttribute() (resp datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetNetworkVlans() (resp []datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetNextBillingCycleBandwidthAllocation() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetNotesHistory() (resp []datatypes.Hardware_Note, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetObject() (resp datatypes.Hardware_Server, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetOpenCancellationTicket() (resp datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetOperatingSystem() (resp datatypes.Software_Component_OperatingSystem, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetOperatingSystemReferenceCode() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetOutboundBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetOutboundPrivateBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetOutboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetOverBandwidthAllocationFlag() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPMInfo() (resp []datatypes.Container_RemoteManagement_PmInfo, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPointOfPresenceLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPowerComponents() (resp []datatypes.Hardware_Power_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPowerSupply() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPrimaryBackendIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPrimaryBackendNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPrimaryDriveSize() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPrimaryIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPrimaryNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPrivateBandwidthData(startTime *int, endTime *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
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
func (r *Hardware_Server) GetPrivateIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPrivateNetworkComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPrivateNetworkOnlyFlag() (resp bool, err error) {
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
func (r *Hardware_Server) GetProcessorCoreAmount() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetProcessorPhysicalCoreAmount() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetProcessors() (resp []datatypes.Hardware_Component, err error) {
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
func (r *Hardware_Server) GetProvisionDate() (resp datatypes.Time, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPublicBandwidthData(startTime *int, endTime *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startTime,
		endTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPublicBandwidthDataSummary() (resp datatypes.Container_Network_Bandwidth_Data_Summary, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetPublicBandwidthGraphImage(startTime *datatypes.Time, endTime *datatypes.Time) (resp []byte, err error) {
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
func (r *Hardware_Server) GetRack() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetRaidControllers() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetRecentEvents() (resp []datatypes.Notification_Occurrence_Event, err error) {
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
func (r *Hardware_Server) GetRemoteManagementAccounts() (resp []datatypes.Hardware_Component_RemoteManagement_User, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetRemoteManagementComponent() (resp datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetRemoteManagementUsers() (resp []datatypes.Hardware_Component_RemoteManagement_User, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetResourceGroupMemberReferences() (resp []datatypes.Resource_Group_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetResourceGroupRoles() (resp []datatypes.Resource_Group_Role, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetResourceGroups() (resp []datatypes.Resource_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetReverseDomainRecords() (resp []datatypes.Dns_Domain, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetScaleAssets() (resp []datatypes.Scale_Asset, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetSecurityScanRequests() (resp []datatypes.Network_Security_Scanner_Request, err error) {
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
func (r *Hardware_Server) GetServerRoom() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetServerTemperatureGraphs() (resp []datatypes.Container_RemoteManagement_Graphs_SensorTemperature, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetServiceProvider() (resp datatypes.Service_Provider, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetSoftwareComponents() (resp []datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetSparePoolBillingItem() (resp datatypes.Billing_Item_Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetSshKeys() (resp []datatypes.Security_Ssh_Key, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetStatisticsRemoteManagement() (resp datatypes.Hardware_Component_RemoteManagement, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetStorageNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetTagReferences() (resp []datatypes.Tag_Reference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetTopLevelLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetTransactionHistory() (resp []datatypes.Provisioning_Version1_Transaction_History, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetUpgradeItemPrices() (resp []datatypes.Product_Item_Price, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetUpgradeRequest() (resp datatypes.Product_Upgrade_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetUplinkHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetUplinkNetworkComponents() (resp []datatypes.Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetUserData() (resp []datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetUsers() (resp []datatypes.User_Customer, err error) {
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
func (r *Hardware_Server) GetVirtualChassis() (resp datatypes.Hardware_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetVirtualChassisSiblings() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetVirtualHost() (resp datatypes.Virtual_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetVirtualLicenses() (resp []datatypes.Software_VirtualLicense, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetVirtualRack() (resp datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetVirtualRackId() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetVirtualRackName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) GetVirtualizationPlatform() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
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
func (r *Hardware_Server) ImportVirtualHost() (resp datatypes.Virtual_Host, err error) {
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
func (r *Hardware_Server) RemoveAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Hardware_Server) RemoveAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
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
func (r *Hardware_Server) SetTags(tags *string) (resp bool, err error) {
	params := []interface{}{
		tags,
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
