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

type Virtual_Disk_Image struct {
	Session *Session
	Options
}

func (r *Session) GetVirtualDiskImageService() Virtual_Disk_Image {
	return Virtual_Disk_Image{Session: r}
}

func (r *Virtual_Disk_Image) EditObject(templateObject *datatypes.Virtual_Disk_Image) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Disk_Image) GetObject() (resp datatypes.Virtual_Disk_Image, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Disk_Image) GetPublicIsoImages() (resp []datatypes.Virtual_Disk_Image, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Virtual_Disk_Image) GetBillingItem() (resp datatypes.Billing_Item_Virtual_Disk_Image, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Disk_Image) GetBlockDevices() (resp []datatypes.Virtual_Guest_Block_Device, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Disk_Image) GetBootableVolumeFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Disk_Image) GetCoalescedDiskImages() (resp []datatypes.Virtual_Disk_Image, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Disk_Image) GetCopyOnWriteFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Disk_Image) GetLocalDiskFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Disk_Image) GetMetadataFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Disk_Image) GetSoftwareReferences() (resp []datatypes.Virtual_Disk_Image_Software, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Disk_Image) GetSourceDiskImage() (resp datatypes.Virtual_Disk_Image, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Disk_Image) GetStorageRepository() (resp datatypes.Virtual_Storage_Repository, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Disk_Image) GetStorageRepositoryType() (resp datatypes.Virtual_Storage_Repository_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Disk_Image) GetTemplateBlockDevice() (resp datatypes.Virtual_Guest_Block_Device_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Disk_Image) GetType() (resp datatypes.Virtual_Disk_Image_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Virtual_Guest struct {
	Session *Session
	Options
}

func (r *Session) GetVirtualGuestService() Virtual_Guest {
	return Virtual_Guest{Session: r}
}

func (r *Virtual_Guest) ActivatePrivatePort() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) ActivatePublicPort() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) AllowAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) AllowAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) AttachDiskImage(imageId *int) (resp datatypes.Provisioning_Version1_Transaction, err error) {
	params := []interface{}{
		imageId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) CancelIsolationForDestructiveAction() (err error) {
	var resp datatypes.Void
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) CaptureImage(captureTemplate *datatypes.Container_Disk_Image_Capture_Template) (resp datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	params := []interface{}{
		captureTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) CheckHostDiskAvailability(diskCapacity *int) (resp bool, err error) {
	params := []interface{}{
		diskCapacity,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) CloseAlarm(alarmId *string) (resp bool, err error) {
	params := []interface{}{
		alarmId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) ConfigureMetadataDisk() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) CreateArchiveTransaction(groupName *string, blockDevices []datatypes.Virtual_Guest_Block_Device, note *string) (resp datatypes.Provisioning_Version1_Transaction, err error) {
	params := []interface{}{
		groupName,
		blockDevices,
		note,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) CreateObject(templateObject *datatypes.Virtual_Guest) (resp datatypes.Virtual_Guest, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) CreateObjects(templateObjects []datatypes.Virtual_Guest) (resp []datatypes.Virtual_Guest, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) CreatePostSoftwareInstallTransaction(data *string, returnBoolean *bool) (resp bool, err error) {
	params := []interface{}{
		data,
		returnBoolean,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) DetachDiskImage(imageId *int) (resp datatypes.Provisioning_Version1_Transaction, err error) {
	params := []interface{}{
		imageId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) EditObject(templateObject *datatypes.Virtual_Guest) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) ExecuteIderaBareMetalRestore() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) ExecuteR1SoftBareMetalRestore() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) ExecuteRemoteScript(uri *string) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		uri,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) ExecuteRescueLayer() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) FindByIpAddress(ipAddress *string) (resp datatypes.Virtual_Guest, err error) {
	params := []interface{}{
		ipAddress,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GenerateOrderTemplate(templateObject *datatypes.Virtual_Guest) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetAdditionalRequiredPricesForOsReload(config *datatypes.Container_Hardware_Server_Configuration) (resp []datatypes.Product_Item_Price, err error) {
	params := []interface{}{
		config,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetAlarmHistory(startDate *datatypes.Time, endDate *datatypes.Time, alarmId *string) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
		alarmId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetAttachedNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetAvailableBlockDevicePositions() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetAvailableNetworkStorages(nasType *string) (resp []datatypes.Network_Storage, err error) {
	params := []interface{}{
		nasType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetBandwidthDataByDate(startDateTime *datatypes.Time, endDateTime *datatypes.Time, networkType *string) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
		networkType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetBandwidthForDateRange(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetBandwidthImage(networkType *string, snapshotRange *string, dateSpecified *datatypes.Time, dateSpecifiedEnd *datatypes.Time) (resp datatypes.Container_Bandwidth_GraphOutputs, err error) {
	params := []interface{}{
		networkType,
		snapshotRange,
		dateSpecified,
		dateSpecifiedEnd,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetBandwidthImageByDate(startDateTime *datatypes.Time, endDateTime *datatypes.Time, networkType *string) (resp datatypes.Container_Bandwidth_GraphOutputs, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
		networkType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetBandwidthTotal(startDateTime *datatypes.Time, endDateTime *datatypes.Time, direction *string, side *string) (resp uint, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
		direction,
		side,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetBootOrder() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetConsoleAccessLog() (resp []datatypes.Network_Logging_Syslog, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetCoreRestrictedOperatingSystemPrice() (resp datatypes.Product_Item_Price, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetCpuMetricDataByDate(startDateTime *datatypes.Time, endDateTime *datatypes.Time, cpuIndexes []int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
		cpuIndexes,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetCpuMetricImage(snapshotRange *string, dateSpecified *datatypes.Time) (resp datatypes.Container_Bandwidth_GraphOutputs, err error) {
	params := []interface{}{
		snapshotRange,
		dateSpecified,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetCpuMetricImageByDate(startDateTime *datatypes.Time, endDateTime *datatypes.Time, cpuIndexes []int) (resp datatypes.Container_Bandwidth_GraphOutputs, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
		cpuIndexes,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetCreateObjectOptions() (resp datatypes.Container_Virtual_Guest_Configuration, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetCurrentBillingDetail() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetCurrentBillingTotal() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetCustomBandwidthDataByDate(graphData *datatypes.Container_Graph) (resp datatypes.Container_Graph, err error) {
	params := []interface{}{
		graphData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetCustomMetricDataByDate(graphData *datatypes.Container_Graph) (resp datatypes.Container_Graph, err error) {
	params := []interface{}{
		graphData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetDriveRetentionItemPrice() (resp datatypes.Product_Item_Price, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetFirewallProtectableSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetFirstAvailableBlockDevicePosition() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetIsoBootImage() (resp datatypes.Virtual_Disk_Image, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetItemPricesFromSoftwareDescriptions(softwareDescriptions []datatypes.Software_Description, includeTranslationsFlag *bool, returnAllPricesFlag *bool) (resp []datatypes.Product_Item, err error) {
	params := []interface{}{
		softwareDescriptions,
		includeTranslationsFlag,
		returnAllPricesFlag,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetMemoryMetricDataByDate(startDateTime *datatypes.Time, endDateTime *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetMemoryMetricImage(snapshotRange *string, dateSpecified *datatypes.Time) (resp datatypes.Container_Bandwidth_GraphOutputs, err error) {
	params := []interface{}{
		snapshotRange,
		dateSpecified,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetMemoryMetricImageByDate(startDateTime *datatypes.Time, endDateTime *datatypes.Time) (resp datatypes.Container_Bandwidth_GraphOutputs, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetMonitoringActiveAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetMonitoringClosedAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetNetworkComponentFirewallProtectableIpAddresses() (resp []datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetObject() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetOrderTemplate(billingType *string, orderPrices []datatypes.Product_Item_Price) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		billingType,
		orderPrices,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetProvisionDate() (resp datatypes.Time, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetRecentMetricData(time *uint) (resp []datatypes.Metric_Tracking_Object, err error) {
	params := []interface{}{
		time,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetRemoteMonitoringActiveAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetRemoteMonitoringClosedAlarms(startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetReverseDomainRecords() (resp []datatypes.Dns_Domain, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetUpgradeItemPrices(includeDowngradeItemPrices *bool) (resp []datatypes.Product_Item_Price, err error) {
	params := []interface{}{
		includeDowngradeItemPrices,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetValidBlockDeviceTemplateGroups(visibility *string) (resp []datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	params := []interface{}{
		visibility,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) IsBackendPingable() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) IsPingable() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) IsolateInstanceForDestructiveAction() (err error) {
	var resp datatypes.Void
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) MountIsoImage(diskImageId *int) (resp datatypes.Provisioning_Version1_Transaction, err error) {
	params := []interface{}{
		diskImageId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) Pause() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) PowerCycle() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) PowerOff() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) PowerOffSoft() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) PowerOn() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) RebootDefault() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) RebootHard() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) RebootSoft() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) ReloadCurrentOperatingSystemConfiguration() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) ReloadOperatingSystem(token *string, config *datatypes.Container_Hardware_Server_Configuration) (resp string, err error) {
	params := []interface{}{
		token,
		config,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) RemoveAccessToNetworkStorage(networkStorageTemplateObject *datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) RemoveAccessToNetworkStorageList(networkStorageTemplateObjects []datatypes.Network_Storage) (resp bool, err error) {
	params := []interface{}{
		networkStorageTemplateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) Resume() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) SetPrivateNetworkInterfaceSpeed(newSpeed *int) (resp bool, err error) {
	params := []interface{}{
		newSpeed,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) SetPublicNetworkInterfaceSpeed(newSpeed *int) (resp bool, err error) {
	params := []interface{}{
		newSpeed,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) SetTags(tags *string) (resp bool, err error) {
	params := []interface{}{
		tags,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) SetUserMetadata(metadata []string) (resp bool, err error) {
	params := []interface{}{
		metadata,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) ShutdownPrivatePort() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) ShutdownPublicPort() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) UnmountIsoImage() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) ValidateImageTemplate(imageTemplateId *int) (resp bool, err error) {
	params := []interface{}{
		imageTemplateId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) VerifyReloadOperatingSystem(config *datatypes.Container_Hardware_Server_Configuration) (resp bool, err error) {
	params := []interface{}{
		config,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

func (r *Virtual_Guest) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetAccountOwnedPoolFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetActiveNetworkMonitorIncident() (resp []datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetActiveTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetActiveTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetActiveTransactions() (resp []datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetAllowedHost() (resp datatypes.Network_Storage_Allowed_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetAllowedNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetAllowedNetworkStorageReplicas() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetAntivirusSpywareSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetApplicationDeliveryController() (resp datatypes.Network_Application_Delivery_Controller, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetAttributes() (resp []datatypes.Virtual_Guest_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetAvailableMonitoring() (resp []datatypes.Network_Monitor_Version1_Query_Host_Stratum, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetAverageDailyPrivateBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetAverageDailyPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetBackendNetworkComponents() (resp []datatypes.Virtual_Guest_Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetBackendRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetBandwidthAllocation() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetBandwidthAllotmentDetail() (resp datatypes.Network_Bandwidth_Version1_Allotment_Detail, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetBillingCycleBandwidthUsage() (resp []datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetBillingCyclePrivateBandwidthUsage() (resp datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetBillingCyclePublicBandwidthUsage() (resp datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetBillingItem() (resp datatypes.Billing_Item_Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetBlockCancelBecauseDisconnectedFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetBlockDeviceTemplateGroup() (resp datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetBlockDevices() (resp []datatypes.Virtual_Guest_Block_Device, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetConsoleIpAddressFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetConsoleIpAddressRecord() (resp datatypes.Virtual_Guest_Network_Component_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetContinuousDataProtectionSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetControlPanel() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetCurrentBandwidthSummary() (resp datatypes.Metric_Tracking_Object_Bandwidth_Summary, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetDatacenter() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetEvaultNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetFirewallServiceComponent() (resp datatypes.Network_Component_Firewall, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetFrontendNetworkComponents() (resp []datatypes.Virtual_Guest_Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetFrontendRouters() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetGlobalIdentifier() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetGuestBootParameter() (resp datatypes.Virtual_Guest_Boot_Parameter, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetHost() (resp datatypes.Virtual_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetHostIpsSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetHourlyBillingFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetInboundPrivateBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetInboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetInternalTagReferences() (resp []datatypes.Tag_Reference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetLastKnownPowerState() (resp datatypes.Virtual_Guest_Power_State, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetLastOperatingSystemReload() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetLastTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetLatestNetworkMonitorIncident() (resp datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetLocalDiskFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetManagedResourceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetMetricTrackingObject() (resp datatypes.Metric_Tracking_Object, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetMetricTrackingObjectId() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetMonitoringAgents() (resp []datatypes.Monitoring_Agent, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetMonitoringRobot() (resp datatypes.Monitoring_Robot, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetMonitoringServiceComponent() (resp datatypes.Network_Monitor_Version1_Query_Host_Stratum, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetMonitoringServiceEligibilityFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetMonitoringServiceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetMonitoringUserNotification() (resp []datatypes.User_Customer_Notification_Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetNetworkComponents() (resp []datatypes.Virtual_Guest_Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetNetworkMonitorIncidents() (resp []datatypes.Network_Monitor_Version1_Incident, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetNetworkMonitors() (resp []datatypes.Network_Monitor_Version1_Query_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetNetworkVlans() (resp []datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetOpenCancellationTicket() (resp datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetOperatingSystem() (resp datatypes.Software_Component_OperatingSystem, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetOperatingSystemReferenceCode() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetOrderedPackageId() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetOutboundPrivateBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetOutboundPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetOverBandwidthAllocationFlag() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetPowerState() (resp datatypes.Virtual_Guest_Power_State, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetPrimaryBackendIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetPrimaryBackendNetworkComponent() (resp datatypes.Virtual_Guest_Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetPrimaryIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetPrimaryNetworkComponent() (resp datatypes.Virtual_Guest_Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetPrivateNetworkOnlyFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetProjectedOverBandwidthAllocationFlag() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetProjectedPublicBandwidthUsage() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetRecentEvents() (resp []datatypes.Notification_Occurrence_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetRegionalGroup() (resp datatypes.Location_Group_Regional, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetRegionalInternetRegistry() (resp datatypes.Network_Regional_Internet_Registry, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetScaleAssets() (resp []datatypes.Scale_Asset, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetScaleMember() (resp datatypes.Scale_Member_Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetScaledFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetSecurityScanRequests() (resp []datatypes.Network_Security_Scanner_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetServerRoom() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetSoftwareComponents() (resp []datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetSshKeys() (resp []datatypes.Security_Ssh_Key, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetStatus() (resp datatypes.Virtual_Guest_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetTagReferences() (resp []datatypes.Tag_Reference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetUpgradeRequest() (resp datatypes.Product_Upgrade_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetUserData() (resp []datatypes.Virtual_Guest_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetUsers() (resp []datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetVirtualRack() (resp datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetVirtualRackId() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest) GetVirtualRackName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Virtual_Guest_Block_Device_Template_Group struct {
	Session *Session
	Options
}

func (r *Session) GetVirtualGuestBlockDeviceTemplateGroupService() Virtual_Guest_Block_Device_Template_Group {
	return Virtual_Guest_Block_Device_Template_Group{Session: r}
}

func (r *Virtual_Guest_Block_Device_Template_Group) AddLocations(locations []datatypes.Location) (resp bool, err error) {
	params := []interface{}{
		locations,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) CopyToExternalSource(configuration *datatypes.Container_Virtual_Guest_Block_Device_Template_Configuration) (resp bool, err error) {
	params := []interface{}{
		configuration,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) CreateFromExternalSource(configuration *datatypes.Container_Virtual_Guest_Block_Device_Template_Configuration) (resp datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	params := []interface{}{
		configuration,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) CreatePublicArchiveTransaction(groupName *string, summary *string, note *string, locations []datatypes.Location) (resp int, err error) {
	params := []interface{}{
		groupName,
		summary,
		note,
		locations,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) DeleteObject() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) DenySharingAccess(accountId *int) (resp bool, err error) {
	params := []interface{}{
		accountId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) EditObject(templateObject *datatypes.Virtual_Guest_Block_Device_Template_Group) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) GetObject() (resp datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) GetPublicCustomerOwnedImages() (resp []datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) GetPublicImages() (resp []datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) GetStorageLocations() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) GetVhdImportSoftwareDescriptions() (resp []datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) PermitSharingAccess(accountId *int) (resp bool, err error) {
	params := []interface{}{
		accountId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) RemoveLocations(locations []datatypes.Location) (resp bool, err error) {
	params := []interface{}{
		locations,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) SetAvailableLocations(locations []datatypes.Location) (resp bool, err error) {
	params := []interface{}{
		locations,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) SetTags(tags *string) (resp bool, err error) {
	params := []interface{}{
		tags,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

func (r *Virtual_Guest_Block_Device_Template_Group) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) GetAccountContacts() (resp []datatypes.Account_Contact, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) GetAccountReferences() (resp []datatypes.Virtual_Guest_Block_Device_Template_Group_Accounts, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) GetBlockDevices() (resp []datatypes.Virtual_Guest_Block_Device_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) GetBlockDevicesDiskSpaceTotal() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) GetChildren() (resp []datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) GetDatacenter() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) GetDatacenters() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) GetFlexImageFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) GetGlobalIdentifier() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) GetImageType() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) GetImageTypeKeyName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) GetParent() (resp datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) GetSshKeys() (resp []datatypes.Security_Ssh_Key, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) GetStatus() (resp datatypes.Virtual_Guest_Block_Device_Template_Group_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) GetStorageRepository() (resp datatypes.Virtual_Storage_Repository, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) GetTagReferences() (resp []datatypes.Tag_Reference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Block_Device_Template_Group) GetTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Virtual_Guest_Boot_Parameter struct {
	Session *Session
	Options
}

func (r *Session) GetVirtualGuestBootParameterService() Virtual_Guest_Boot_Parameter {
	return Virtual_Guest_Boot_Parameter{Session: r}
}

func (r *Virtual_Guest_Boot_Parameter) CreateObject(templateObject *datatypes.Virtual_Guest_Boot_Parameter) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Boot_Parameter) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Boot_Parameter) EditObject(templateObject *datatypes.Virtual_Guest_Boot_Parameter) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Boot_Parameter) GetObject() (resp datatypes.Virtual_Guest_Boot_Parameter, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Virtual_Guest_Boot_Parameter) GetGuest() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Boot_Parameter) GetGuestBootParameterType() (resp datatypes.Virtual_Guest_Boot_Parameter_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Virtual_Guest_Boot_Parameter_Type struct {
	Session *Session
	Options
}

func (r *Session) GetVirtualGuestBootParameterTypeService() Virtual_Guest_Boot_Parameter_Type {
	return Virtual_Guest_Boot_Parameter_Type{Session: r}
}

func (r *Virtual_Guest_Boot_Parameter_Type) GetAllObjects() (resp []datatypes.Virtual_Guest_Boot_Parameter_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Boot_Parameter_Type) GetObject() (resp datatypes.Virtual_Guest_Boot_Parameter_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Virtual_Guest_Network_Component struct {
	Session *Session
	Options
}

func (r *Session) GetVirtualGuestNetworkComponentService() Virtual_Guest_Network_Component {
	return Virtual_Guest_Network_Component{Session: r}
}

func (r *Virtual_Guest_Network_Component) Disable() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Network_Component) Enable() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Network_Component) GetObject() (resp datatypes.Virtual_Guest_Network_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Network_Component) IsPingable() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Virtual_Guest_Network_Component) GetGuest() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Network_Component) GetHighAvailabilityFirewallFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Network_Component) GetIpAddressBindings() (resp []datatypes.Virtual_Guest_Network_Component_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Network_Component) GetNetworkComponentFirewall() (resp datatypes.Network_Component_Firewall, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Network_Component) GetNetworkVlan() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Network_Component) GetPrimaryIpAddress() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Network_Component) GetPrimaryIpAddressRecord() (resp datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Network_Component) GetPrimarySubnet() (resp datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Network_Component) GetPrimaryVersion6IpAddressRecord() (resp datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Network_Component) GetRouter() (resp datatypes.Hardware_Router, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Guest_Network_Component) GetSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Virtual_Host struct {
	Session *Session
	Options
}

func (r *Session) GetVirtualHostService() Virtual_Host {
	return Virtual_Host{Session: r}
}

func (r *Virtual_Host) GetLiveGuestByUuid(uuid *string) (resp datatypes.Virtual_Guest, err error) {
	params := []interface{}{
		uuid,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Host) GetLiveGuestList() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Host) GetLiveGuestRecentMetricData(uuid *string, time *int, limit *int, interval *int) (resp []datatypes.Metric_Tracking_Object, err error) {
	params := []interface{}{
		uuid,
		time,
		limit,
		interval,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Host) GetObject() (resp datatypes.Virtual_Host, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Host) PauseLiveGuest(uuid *string) (resp bool, err error) {
	params := []interface{}{
		uuid,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Host) PowerCycleLiveGuest(uuid *string) (resp bool, err error) {
	params := []interface{}{
		uuid,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Host) PowerOffLiveGuest(uuid *string) (resp bool, err error) {
	params := []interface{}{
		uuid,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Host) PowerOnLiveGuest(uuid *string) (resp bool, err error) {
	params := []interface{}{
		uuid,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Host) RebootSoftLiveGuest(uuid *string) (resp bool, err error) {
	params := []interface{}{
		uuid,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Host) ResumeLiveGuest(uuid *string) (resp bool, err error) {
	params := []interface{}{
		uuid,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

func (r *Virtual_Host) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Host) GetBilledPerGuestFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Host) GetBilledPerMemoryUsageFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Host) GetGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Host) GetHardware() (resp datatypes.Hardware_Server, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Host) GetMetricTrackingObject() (resp datatypes.Metric_Tracking_Object, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Virtual_Storage_Repository struct {
	Session *Session
	Options
}

func (r *Session) GetVirtualStorageRepositoryService() Virtual_Storage_Repository {
	return Virtual_Storage_Repository{Session: r}
}

func (r *Virtual_Storage_Repository) GetArchiveDiskUsageRatePerGb() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Storage_Repository) GetAverageUsageMetricDataByDate(startDateTime *datatypes.Time, endDateTime *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Storage_Repository) GetObject() (resp datatypes.Virtual_Storage_Repository, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Storage_Repository) GetPublicImageDiskUsageRatePerGb() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Storage_Repository) GetStorageLocations() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Storage_Repository) GetUsageMetricDataByDate(startDateTime *datatypes.Time, endDateTime *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Storage_Repository) GetUsageMetricImageByDate(startDateTime *datatypes.Time, endDateTime *datatypes.Time) (resp datatypes.Container_Bandwidth_GraphOutputs, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

func (r *Virtual_Storage_Repository) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Storage_Repository) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Storage_Repository) GetDatacenter() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Storage_Repository) GetDiskImages() (resp []datatypes.Virtual_Disk_Image, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Storage_Repository) GetGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Storage_Repository) GetMetricTrackingObject() (resp datatypes.Metric_Tracking_Object_Virtual_Storage_Repository, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Storage_Repository) GetPublicImageBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Virtual_Storage_Repository) GetType() (resp datatypes.Virtual_Storage_Repository_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
