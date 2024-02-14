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

// Metric tracking objects provides a common interface to all metrics provided by SoftLayer. These metrics range from network component traffic for a server to aggregated Bandwidth Pooling traffic and more. Every object within SoftLayer's range of objects that has data that can be tracked over time has an associated tracking object. Use the [[SoftLayer_Metric_Tracking_Object]] service to retrieve raw and graph data from a tracking object.
type Metric_Tracking_Object struct {
	Session session.SLSession
	Options sl.Options
}

// GetMetricTrackingObjectService returns an instance of the Metric_Tracking_Object SoftLayer service
func GetMetricTrackingObjectService(sess session.SLSession) Metric_Tracking_Object {
	return Metric_Tracking_Object{Session: sess}
}

func (r Metric_Tracking_Object) Id(id int) Metric_Tracking_Object {
	r.Options.Id = &id
	return r
}

func (r Metric_Tracking_Object) Mask(mask string) Metric_Tracking_Object {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Metric_Tracking_Object) Filter(filter string) Metric_Tracking_Object {
	r.Options.Filter = filter
	return r
}

func (r Metric_Tracking_Object) Limit(limit int) Metric_Tracking_Object {
	r.Options.Limit = &limit
	return r
}

func (r Metric_Tracking_Object) Offset(offset int) Metric_Tracking_Object {
	r.Options.Offset = &offset
	return r
}

// Retrieve a collection of raw bandwidth data from an individual public or private network tracking object. Raw data is ideal if you with to employ your own traffic storage and graphing systems.
func (r Metric_Tracking_Object) GetBandwidthData(startDateTime *datatypes.Time, endDateTime *datatypes.Time, typ *string, rollupSeconds *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
		typ,
		rollupSeconds,
	}
	err = r.Session.DoRequest("SoftLayer_Metric_Tracking_Object", "getBandwidthData", params, &r.Options, &resp)
	return
}

// Returns summarized metric data for the date range, metric type and summary period provided.
func (r Metric_Tracking_Object) GetSummaryData(startDateTime *datatypes.Time, endDateTime *datatypes.Time, validTypes []datatypes.Container_Metric_Data_Type, summaryPeriod *int) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
		validTypes,
		summaryPeriod,
	}
	err = r.Session.DoRequest("SoftLayer_Metric_Tracking_Object", "getSummaryData", params, &r.Options, &resp)
	return
}
