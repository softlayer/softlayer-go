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

type Metric_Tracking_Object struct {
	Session *Session
	Options
}

func (r *Session) GetMetricTrackingObjectService() Metric_Tracking_Object {
	return Metric_Tracking_Object{Session: r}
}

func (r *Metric_Tracking_Object) GetBackboneBandwidthGraph(graphTitle *string) (resp datatypes.Container_Bandwidth_GraphOutputs, err error) {
	params := []interface{}{
		graphTitle,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object) GetBandwidthData(startDateTime *time.Time, endDateTime *time.Time, typ *string, rollupSeconds *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
		typ,
		rollupSeconds,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object) GetBandwidthGraph(startDateTime *time.Time, endDateTime *time.Time, graphType *string, fontSize *int, graphWidth *int, graphHeight *int, doNotShowTimeZone *bool) (resp datatypes.Container_Bandwidth_GraphOutputs, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
		graphType,
		fontSize,
		graphWidth,
		graphHeight,
		doNotShowTimeZone,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object) GetBandwidthTotal(startDateTime *time.Time, endDateTime *time.Time, direction *string, typ *string) (resp uint, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
		direction,
		typ,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object) GetCustomGraphData(graphContainer *datatypes.Container_Graph) (resp datatypes.Container_Graph, err error) {
	params := []interface{}{
		graphContainer,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object) GetDetailsForDateRange(startDate *time.Time, endDate *time.Time, graphType []string) (resp []datatypes.Container_Metric_Tracking_Object_Details, err error) {
	params := []interface{}{
		startDate,
		endDate,
		graphType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object) GetGraph(startDateTime *time.Time, endDateTime *time.Time, graphType []string) (resp datatypes.Container_Bandwidth_GraphOutputs, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
		graphType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object) GetMetricDataTypes() (resp []datatypes.Container_Metric_Data_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object) GetObject() (resp datatypes.Metric_Tracking_Object, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object) GetSummary(graphType *string) (resp datatypes.Container_Metric_Tracking_Object_Summary, err error) {
	params := []interface{}{
		graphType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object) GetSummaryData(startDateTime *time.Time, endDateTime *time.Time, validTypes []datatypes.Container_Metric_Data_Type, summaryPeriod *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
		validTypes,
		summaryPeriod,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

func (r *Metric_Tracking_Object) GetType() (resp datatypes.Metric_Tracking_Object_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Metric_Tracking_Object_Abstract struct {
	Session *Session
	Options
}

func (r *Session) GetMetricTrackingObjectAbstractService() Metric_Tracking_Object_Abstract {
	return Metric_Tracking_Object_Abstract{Session: r}
}

type Metric_Tracking_Object_Bandwidth_Summary struct {
	Session *Session
	Options
}

func (r *Session) GetMetricTrackingObjectBandwidthSummaryService() Metric_Tracking_Object_Bandwidth_Summary {
	return Metric_Tracking_Object_Bandwidth_Summary{Session: r}
}

func (r *Metric_Tracking_Object_Bandwidth_Summary) GetObject() (resp datatypes.Metric_Tracking_Object_Bandwidth_Summary, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Metric_Tracking_Object_Data struct {
	Session *Session
	Options
}

func (r *Session) GetMetricTrackingObjectDataService() Metric_Tracking_Object_Data {
	return Metric_Tracking_Object_Data{Session: r}
}

type Metric_Tracking_Object_Data_Network_ContentDelivery_Account struct {
	Session *Session
	Options
}

func (r *Session) GetMetricTrackingObjectDataNetworkContentDeliveryAccountService() Metric_Tracking_Object_Data_Network_ContentDelivery_Account {
	return Metric_Tracking_Object_Data_Network_ContentDelivery_Account{Session: r}
}

type Metric_Tracking_Object_HardwareServer struct {
	Session *Session
	Options
}

func (r *Session) GetMetricTrackingObjectHardwareServerService() Metric_Tracking_Object_HardwareServer {
	return Metric_Tracking_Object_HardwareServer{Session: r}
}

func (r *Metric_Tracking_Object_HardwareServer) GetBillingCycleBandwidthUsage() (resp []datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object_HardwareServer) GetBillingCyclePrivateBandwidthUsage() (resp []datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object_HardwareServer) GetBillingCyclePrivateUsageIn() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object_HardwareServer) GetBillingCyclePrivateUsageOut() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object_HardwareServer) GetBillingCyclePrivateUsageTotal() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object_HardwareServer) GetBillingCyclePublicBandwidthUsage() (resp datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object_HardwareServer) GetBillingCyclePublicUsageIn() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object_HardwareServer) GetBillingCyclePublicUsageOut() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object_HardwareServer) GetBillingCyclePublicUsageTotal() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object_HardwareServer) GetResource() (resp datatypes.Hardware_Server, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Metric_Tracking_Object_Type struct {
	Session *Session
	Options
}

func (r *Session) GetMetricTrackingObjectTypeService() Metric_Tracking_Object_Type {
	return Metric_Tracking_Object_Type{Session: r}
}

type Metric_Tracking_Object_VirtualDedicatedRack struct {
	Session *Session
	Options
}

func (r *Session) GetMetricTrackingObjectVirtualDedicatedRackService() Metric_Tracking_Object_VirtualDedicatedRack {
	return Metric_Tracking_Object_VirtualDedicatedRack{Session: r}
}

func (r *Metric_Tracking_Object_VirtualDedicatedRack) GetBillingCycleBandwidthUsage() (resp []datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object_VirtualDedicatedRack) GetBillingCyclePrivateBandwidthUsage() (resp []datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object_VirtualDedicatedRack) GetBillingCyclePrivateUsageIn() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object_VirtualDedicatedRack) GetBillingCyclePrivateUsageOut() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object_VirtualDedicatedRack) GetBillingCyclePrivateUsageTotal() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object_VirtualDedicatedRack) GetBillingCyclePublicBandwidthUsage() (resp datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object_VirtualDedicatedRack) GetBillingCyclePublicUsageIn() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object_VirtualDedicatedRack) GetBillingCyclePublicUsageOut() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object_VirtualDedicatedRack) GetBillingCyclePublicUsageTotal() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object_VirtualDedicatedRack) GetResource() (resp datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Metric_Tracking_Object_Virtual_Storage_Repository struct {
	Session *Session
	Options
}

func (r *Session) GetMetricTrackingObjectVirtualStorageRepositoryService() Metric_Tracking_Object_Virtual_Storage_Repository {
	return Metric_Tracking_Object_Virtual_Storage_Repository{Session: r}
}

func (r *Metric_Tracking_Object_Virtual_Storage_Repository) GetResource() (resp datatypes.Virtual_Storage_Repository, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
