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

package datatypes

import "time"

type Metric_Tracking_Object struct {
	Entity

	Data            []Metric_Tracking_Object_Data `json:"data,omitempty"`
	Id              *int                          `json:"id,omitempty"`
	Label           *string                       `json:"label,omitempty"`
	ResourceTableId *int                          `json:"resourceTableId,omitempty"`
	StartDate       *time.Time                    `json:"startDate,omitempty"`
	Type            *Metric_Tracking_Object_Type  `json:"type,omitempty"`
}

type Metric_Tracking_Object_Abstract struct {
	Metric_Tracking_Object
}

type Metric_Tracking_Object_Bandwidth_Summary struct {
	Entity

	AllocationAmount            *float64 `json:"allocationAmount,omitempty"`
	AllocationId                *int     `json:"allocationId,omitempty"`
	AmountOut                   *float64 `json:"amountOut,omitempty"`
	AverageDailyUsage           *float64 `json:"averageDailyUsage,omitempty"`
	CurrentlyOverAllocationFlag *int     `json:"currentlyOverAllocationFlag,omitempty"`
	Id                          *int     `json:"id,omitempty"`
	OutboundBandwidthAmount     *float64 `json:"outboundBandwidthAmount,omitempty"`
	ProjectedBandwidthUsage     *float64 `json:"projectedBandwidthUsage,omitempty"`
	ProjectedOverAllocationFlag *int     `json:"projectedOverAllocationFlag,omitempty"`
}

type Metric_Tracking_Object_Data struct {
	Entity

	Counter  *float64   `json:"counter,omitempty"`
	DateTime *time.Time `json:"dateTime,omitempty"`
	Type     *string    `json:"type,omitempty"`
}

type Metric_Tracking_Object_Data_Network_ContentDelivery_Account struct {
	Metric_Tracking_Object_Data

	FileName *string `json:"fileName,omitempty"`
	PopId    *int    `json:"popId,omitempty"`
}

type Metric_Tracking_Object_HardwareServer struct {
	Metric_Tracking_Object_Abstract

	BillingCycleBandwidthUsage             []Network_Bandwidth_Usage `json:"billingCycleBandwidthUsage,omitempty"`
	BillingCycleBandwidthUsageCount        *uint                     `json:"billingCycleBandwidthUsageCount,omitempty"`
	BillingCyclePrivateBandwidthUsage      []Network_Bandwidth_Usage `json:"billingCyclePrivateBandwidthUsage,omitempty"`
	BillingCyclePrivateBandwidthUsageCount *uint                     `json:"billingCyclePrivateBandwidthUsageCount,omitempty"`
	BillingCyclePrivateUsageIn             *float64                  `json:"billingCyclePrivateUsageIn,omitempty"`
	BillingCyclePrivateUsageOut            *float64                  `json:"billingCyclePrivateUsageOut,omitempty"`
	BillingCyclePrivateUsageTotal          *uint                     `json:"billingCyclePrivateUsageTotal,omitempty"`
	BillingCyclePublicBandwidthUsage       *Network_Bandwidth_Usage  `json:"billingCyclePublicBandwidthUsage,omitempty"`
	BillingCyclePublicUsageIn              *float64                  `json:"billingCyclePublicUsageIn,omitempty"`
	BillingCyclePublicUsageOut             *float64                  `json:"billingCyclePublicUsageOut,omitempty"`
	BillingCyclePublicUsageTotal           *uint                     `json:"billingCyclePublicUsageTotal,omitempty"`
	Resource                               *Hardware_Server          `json:"resource,omitempty"`
}

type Metric_Tracking_Object_Type struct {
	Entity

	Keyname *string `json:"keyname,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type Metric_Tracking_Object_VirtualDedicatedRack struct {
	Metric_Tracking_Object_Abstract

	BillingCycleBandwidthUsage             []Network_Bandwidth_Usage             `json:"billingCycleBandwidthUsage,omitempty"`
	BillingCycleBandwidthUsageCount        *uint                                 `json:"billingCycleBandwidthUsageCount,omitempty"`
	BillingCyclePrivateBandwidthUsage      []Network_Bandwidth_Usage             `json:"billingCyclePrivateBandwidthUsage,omitempty"`
	BillingCyclePrivateBandwidthUsageCount *uint                                 `json:"billingCyclePrivateBandwidthUsageCount,omitempty"`
	BillingCyclePrivateUsageIn             *float64                              `json:"billingCyclePrivateUsageIn,omitempty"`
	BillingCyclePrivateUsageOut            *float64                              `json:"billingCyclePrivateUsageOut,omitempty"`
	BillingCyclePrivateUsageTotal          *uint                                 `json:"billingCyclePrivateUsageTotal,omitempty"`
	BillingCyclePublicBandwidthUsage       *Network_Bandwidth_Usage              `json:"billingCyclePublicBandwidthUsage,omitempty"`
	BillingCyclePublicUsageIn              *float64                              `json:"billingCyclePublicUsageIn,omitempty"`
	BillingCyclePublicUsageOut             *float64                              `json:"billingCyclePublicUsageOut,omitempty"`
	BillingCyclePublicUsageTotal           *uint                                 `json:"billingCyclePublicUsageTotal,omitempty"`
	Resource                               *Network_Bandwidth_Version1_Allotment `json:"resource,omitempty"`
}

type Metric_Tracking_Object_Virtual_Storage_Repository struct {
	Metric_Tracking_Object_Abstract

	Resource *Virtual_Storage_Repository `json:"resource,omitempty"`
}
