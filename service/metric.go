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
func (r *Metric_Tracking_Object) GetBandwidthData(startDateTime *datatypes.Time, endDateTime *datatypes.Time, typ *string, rollupSeconds *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
		typ,
		rollupSeconds,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object) GetBandwidthGraph(startDateTime *datatypes.Time, endDateTime *datatypes.Time, graphType *string, fontSize *int, graphWidth *int, graphHeight *int, doNotShowTimeZone *bool) (resp datatypes.Container_Bandwidth_GraphOutputs, err error) {
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
func (r *Metric_Tracking_Object) GetBandwidthTotal(startDateTime *datatypes.Time, endDateTime *datatypes.Time, direction *string, typ *string) (resp uint, err error) {
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
func (r *Metric_Tracking_Object) GetDetailsForDateRange(startDate *datatypes.Time, endDate *datatypes.Time, graphType []string) (resp []datatypes.Container_Metric_Tracking_Object_Details, err error) {
	params := []interface{}{
		startDate,
		endDate,
		graphType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Metric_Tracking_Object) GetGraph(startDateTime *datatypes.Time, endDateTime *datatypes.Time, graphType []string) (resp datatypes.Container_Bandwidth_GraphOutputs, err error) {
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
func (r *Metric_Tracking_Object) GetSummaryData(startDateTime *datatypes.Time, endDateTime *datatypes.Time, validTypes []datatypes.Container_Metric_Data_Type, summaryPeriod *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
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
