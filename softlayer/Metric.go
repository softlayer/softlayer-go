package softlayer

import (
	"time"

	"github.ibm.com/riethm/gopherlayer/datatypes"
)

type Metric_Tracking_Object interface {
	Entity

	GetBackboneBandwidthGraph(graphTitle string) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetBandwidthData(startDateTime time.Time, endDateTime time.Time, typ string, rollupSeconds int) (datatypes.Metric_Tracking_Object_Data, error)
	GetBandwidthGraph(startDateTime time.Time, endDateTime time.Time, graphType string, fontSize int, graphWidth int, graphHeight int, doNotShowTimeZone bool) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetBandwidthTotal(startDateTime time.Time, endDateTime time.Time, direction string, typ string) (uint, error)
	GetCustomGraphData(graphContainer datatypes.Container_Graph) (datatypes.Container_Graph, error)
	GetDetailsForDateRange(startDate time.Time, endDate time.Time, graphType string) (datatypes.Container_Metric_Tracking_Object_Details, error)
	GetGraph(startDateTime time.Time, endDateTime time.Time, graphType string) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetMetricDataTypes() (datatypes.Container_Metric_Data_Type, error)
	GetObject() (datatypes.Metric_Tracking_Object, error)
	GetSummary(graphType string) (datatypes.Container_Metric_Tracking_Object_Summary, error)
	GetSummaryData(startDateTime time.Time, endDateTime time.Time, validTypes datatypes.Container_Metric_Data_Type, summaryPeriod int) (datatypes.Metric_Tracking_Object_Data, error)

	GetType() (datatypes.Metric_Tracking_Object_Type, error)
}

type Metric_Tracking_Object_Abstract interface {
	Metric_Tracking_Object
}

type Metric_Tracking_Object_Bandwidth_Summary interface {
	Entity

	GetObject() (datatypes.Metric_Tracking_Object_Bandwidth_Summary, error)
}

type Metric_Tracking_Object_Data interface {
	Entity
}

type Metric_Tracking_Object_Data_Network_ContentDelivery_Account interface {
	Metric_Tracking_Object_Data
}

type Metric_Tracking_Object_HardwareServer interface {
	Metric_Tracking_Object_Abstract

	GetBillingCycleBandwidthUsage() (datatypes.Network_Bandwidth_Usage, error)
	GetBillingCyclePrivateBandwidthUsage() (datatypes.Network_Bandwidth_Usage, error)
	GetBillingCyclePrivateUsageIn() (float64, error)
	GetBillingCyclePrivateUsageOut() (float64, error)
	GetBillingCyclePrivateUsageTotal() (uint, error)
	GetBillingCyclePublicBandwidthUsage() (datatypes.Network_Bandwidth_Usage, error)
	GetBillingCyclePublicUsageIn() (float64, error)
	GetBillingCyclePublicUsageOut() (float64, error)
	GetBillingCyclePublicUsageTotal() (uint, error)
	GetResource() (datatypes.Hardware_Server, error)
}

type Metric_Tracking_Object_Type interface {
	Entity
}

type Metric_Tracking_Object_VirtualDedicatedRack interface {
	Metric_Tracking_Object_Abstract

	GetBillingCycleBandwidthUsage() (datatypes.Network_Bandwidth_Usage, error)
	GetBillingCyclePrivateBandwidthUsage() (datatypes.Network_Bandwidth_Usage, error)
	GetBillingCyclePrivateUsageIn() (float64, error)
	GetBillingCyclePrivateUsageOut() (float64, error)
	GetBillingCyclePrivateUsageTotal() (uint, error)
	GetBillingCyclePublicBandwidthUsage() (datatypes.Network_Bandwidth_Usage, error)
	GetBillingCyclePublicUsageIn() (float64, error)
	GetBillingCyclePublicUsageOut() (float64, error)
	GetBillingCyclePublicUsageTotal() (uint, error)
	GetResource() (datatypes.Network_Bandwidth_Version1_Allotment, error)
}

type Metric_Tracking_Object_Virtual_Storage_Repository interface {
	Metric_Tracking_Object_Abstract

	GetResource() (datatypes.Virtual_Storage_Repository, error)
}
