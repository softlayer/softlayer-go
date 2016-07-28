package softlayer

import (
	"time"

	"github.ibm.com/riethm/gopherlayer/datatypes"
)

type Monitoring_Agent interface {
	Entity

	Activate() (bool, error)
	AddConfigurationProfile(configurationValues datatypes.Monitoring_Agent_Configuration_Value) (datatypes.Provisioning_Version1_Transaction, error)
	ApplyConfigurationValues(configurationValues datatypes.Monitoring_Agent_Configuration_Value) (datatypes.Provisioning_Version1_Transaction, error)
	Deactivate() (bool, error)
	DeleteConfigurationProfile(sectionId int, profileId int) (bool, error)
	DeployMonitoringAgent(configurationTemplateId int) (datatypes.Provisioning_Version1_Transaction, error)
	GetActiveAlarmSubscribers() (datatypes.Notification_User_Subscriber, error)
	GetAvailableConfigurationTemplates() (datatypes.Configuration_Template, error)
	GetAvailableConfigurationValues(configurationDefinitionId int, configValues datatypes.Monitoring_Agent_Configuration_Value) (datatypes.Monitoring_Agent_Configuration_Value, error)
	GetEligibleAlarmSubscibers() (datatypes.User_Customer, error)
	GetGraph(configurationValues datatypes.Monitoring_Agent_Configuration_Value, beginDate time.Time, endDate time.Time) (datatypes.Container_Monitoring_Graph_Outputs, error)
	GetGraphData(metricDataTypes datatypes.Container_Metric_Data_Type, startDate time.Time, endDate time.Time) (datatypes.Metric_Tracking_Object_Data, error)
	GetObject() (datatypes.Monitoring_Agent, error)
	RemoveActiveAlarmSubscriber(userRecordId int) (bool, error)
	RemoveAllAlarmSubscribers() (bool, error)
	RestartMonitoringAgent() (bool, error)
	SetActiveAlarmSubscriber(userRecordId int) (bool, error)

	GetAgentStatus() (datatypes.Monitoring_Agent_Status, error)
	GetConfigurationProfiles() (datatypes.Configuration_Template_Section_Profile, error)
	GetConfigurationTemplate() (datatypes.Configuration_Template, error)
	GetConfigurationValues() (datatypes.Monitoring_Agent_Configuration_Value, error)
	GetHardware() (datatypes.Hardware, error)
	GetProductItem() (datatypes.Product_Item, error)
	GetSoftwareDescription() (datatypes.Software_Description, error)
	GetStatusName() (string, error)
	GetVirtualGuest() (datatypes.Virtual_Guest, error)
}

type Monitoring_Agent_Configuration_Template_Group interface {
	Entity

	CreateObject(templateObject datatypes.Monitoring_Agent_Configuration_Template_Group) (datatypes.Monitoring_Agent_Configuration_Template_Group, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Monitoring_Agent_Configuration_Template_Group) (bool, error)
	GetAllObjects() (datatypes.Monitoring_Agent_Configuration_Template_Group, error)
	GetConfigurationGroups(packageId int) (datatypes.Monitoring_Agent_Configuration_Template_Group, error)
	GetObject() (datatypes.Monitoring_Agent_Configuration_Template_Group, error)

	GetAccount() (datatypes.Account, error)
	GetConfigurationTemplateReferences() (datatypes.Monitoring_Agent_Configuration_Template_Group_Reference, error)
	GetConfigurationTemplates() (datatypes.Configuration_Template, error)
	GetItem() (datatypes.Product_Item, error)
}

type Monitoring_Agent_Configuration_Template_Group_Reference interface {
	Entity

	CreateObject(templateObject datatypes.Monitoring_Agent_Configuration_Template_Group_Reference) (datatypes.Monitoring_Agent_Configuration_Template_Group_Reference, error)
	CreateObjects(templateObjects datatypes.Monitoring_Agent_Configuration_Template_Group_Reference) (bool, error)
	EditObject(templateObject datatypes.Monitoring_Agent_Configuration_Template_Group_Reference) (bool, error)
	EditObjects(templateObjects datatypes.Monitoring_Agent_Configuration_Template_Group_Reference) (bool, error)
	GetAllObjects() (datatypes.Monitoring_Agent_Configuration_Template_Group_Reference, error)
	GetObject() (datatypes.Monitoring_Agent_Configuration_Template_Group_Reference, error)

	GetConfigurationTemplate() (datatypes.Configuration_Template, error)
	GetTemplateGroup() (datatypes.Monitoring_Agent_Configuration_Template_Group, error)
}

type Monitoring_Agent_Configuration_Value interface {
	Entity

	GetObject() (datatypes.Monitoring_Agent_Configuration_Value, error)

	GetDefinition() (datatypes.Configuration_Template_Section_Definition, error)
	GetMetricDataType() (datatypes.Container_Metric_Data_Type, error)
	GetMonitoringAgent() (datatypes.Monitoring_Agent, error)
	GetProfile() (datatypes.Configuration_Template_Section_Profile, error)
}

type Monitoring_Agent_Status interface {
	Entity

	GetObject() (datatypes.Monitoring_Agent_Status, error)
}

type Monitoring_Robot interface {
	Entity

	CheckConnection() (bool, error)
	DeployMonitoringAgents(configurationTemplateGroup datatypes.Monitoring_Agent_Configuration_Template_Group) (datatypes.Provisioning_Version1_Transaction, error)
	GetAvailableConfigurationGroups() (datatypes.Monitoring_Agent_Configuration_Template_Group, error)
	GetObject() (datatypes.Monitoring_Robot, error)
	ResetStatus() (bool, error)

	GetAccount() (datatypes.Account, error)
	GetMonitoringAgents() (datatypes.Monitoring_Agent, error)
	GetRobotStatus() (datatypes.Monitoring_Robot_Status, error)
	GetSoftwareComponent() (datatypes.Software_Component, error)
}

type Monitoring_Robot_Status interface {
	Entity
}
