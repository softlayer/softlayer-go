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

type Monitoring_Agent struct {
	Session *Session
	Options
}

func (r *Session) GetMonitoringAgentService() Monitoring_Agent {
	return Monitoring_Agent{Session: r}
}

func (r *Monitoring_Agent) Activate() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent) AddConfigurationProfile(configurationValues []datatypes.Monitoring_Agent_Configuration_Value) (resp datatypes.Provisioning_Version1_Transaction, err error) {
	params := []interface{}{
		configurationValues,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent) ApplyConfigurationValues(configurationValues []datatypes.Monitoring_Agent_Configuration_Value) (resp datatypes.Provisioning_Version1_Transaction, err error) {
	params := []interface{}{
		configurationValues,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent) Deactivate() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent) DeleteConfigurationProfile(sectionId *int, profileId *int) (resp bool, err error) {
	params := []interface{}{
		sectionId,
		profileId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent) DeployMonitoringAgent(configurationTemplateId *int) (resp datatypes.Provisioning_Version1_Transaction, err error) {
	params := []interface{}{
		configurationTemplateId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent) GetActiveAlarmSubscribers() (resp []datatypes.Notification_User_Subscriber, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent) GetAgentStatus() (resp datatypes.Monitoring_Agent_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent) GetAvailableConfigurationTemplates() (resp []datatypes.Configuration_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent) GetAvailableConfigurationValues(configurationDefinitionId *int, configValues []datatypes.Monitoring_Agent_Configuration_Value) (resp []datatypes.Monitoring_Agent_Configuration_Value, err error) {
	params := []interface{}{
		configurationDefinitionId,
		configValues,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent) GetConfigurationProfiles() (resp []datatypes.Configuration_Template_Section_Profile, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent) GetConfigurationTemplate() (resp datatypes.Configuration_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent) GetConfigurationValues() (resp []datatypes.Monitoring_Agent_Configuration_Value, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent) GetEligibleAlarmSubscibers() (resp []datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent) GetGraph(configurationValues []datatypes.Monitoring_Agent_Configuration_Value, beginDate *datatypes.Time, endDate *datatypes.Time) (resp datatypes.Container_Monitoring_Graph_Outputs, err error) {
	params := []interface{}{
		configurationValues,
		beginDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent) GetGraphData(metricDataTypes []datatypes.Container_Metric_Data_Type, startDate *datatypes.Time, endDate *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		metricDataTypes,
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent) GetObject() (resp datatypes.Monitoring_Agent, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent) GetProductItem() (resp datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent) GetSoftwareDescription() (resp datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent) GetStatusName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent) GetVirtualGuest() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent) RemoveActiveAlarmSubscriber(userRecordId *int) (resp bool, err error) {
	params := []interface{}{
		userRecordId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent) RemoveAllAlarmSubscribers() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent) RestartMonitoringAgent() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent) SetActiveAlarmSubscriber(userRecordId *int) (resp bool, err error) {
	params := []interface{}{
		userRecordId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Monitoring_Agent_Configuration_Template_Group struct {
	Session *Session
	Options
}

func (r *Session) GetMonitoringAgentConfigurationTemplateGroupService() Monitoring_Agent_Configuration_Template_Group {
	return Monitoring_Agent_Configuration_Template_Group{Session: r}
}

func (r *Monitoring_Agent_Configuration_Template_Group) CreateObject(templateObject *datatypes.Monitoring_Agent_Configuration_Template_Group) (resp datatypes.Monitoring_Agent_Configuration_Template_Group, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent_Configuration_Template_Group) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent_Configuration_Template_Group) EditObject(templateObject *datatypes.Monitoring_Agent_Configuration_Template_Group) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent_Configuration_Template_Group) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent_Configuration_Template_Group) GetAllObjects() (resp []datatypes.Monitoring_Agent_Configuration_Template_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent_Configuration_Template_Group) GetConfigurationGroups(packageId *int) (resp []datatypes.Monitoring_Agent_Configuration_Template_Group, err error) {
	params := []interface{}{
		packageId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent_Configuration_Template_Group) GetConfigurationTemplateReferences() (resp []datatypes.Monitoring_Agent_Configuration_Template_Group_Reference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent_Configuration_Template_Group) GetConfigurationTemplates() (resp []datatypes.Configuration_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent_Configuration_Template_Group) GetItem() (resp datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent_Configuration_Template_Group) GetObject() (resp datatypes.Monitoring_Agent_Configuration_Template_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Monitoring_Agent_Configuration_Template_Group_Reference struct {
	Session *Session
	Options
}

func (r *Session) GetMonitoringAgentConfigurationTemplateGroupReferenceService() Monitoring_Agent_Configuration_Template_Group_Reference {
	return Monitoring_Agent_Configuration_Template_Group_Reference{Session: r}
}

func (r *Monitoring_Agent_Configuration_Template_Group_Reference) CreateObject(templateObject *datatypes.Monitoring_Agent_Configuration_Template_Group_Reference) (resp datatypes.Monitoring_Agent_Configuration_Template_Group_Reference, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent_Configuration_Template_Group_Reference) CreateObjects(templateObjects []datatypes.Monitoring_Agent_Configuration_Template_Group_Reference) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent_Configuration_Template_Group_Reference) EditObject(templateObject *datatypes.Monitoring_Agent_Configuration_Template_Group_Reference) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent_Configuration_Template_Group_Reference) EditObjects(templateObjects []datatypes.Monitoring_Agent_Configuration_Template_Group_Reference) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent_Configuration_Template_Group_Reference) GetAllObjects() (resp []datatypes.Monitoring_Agent_Configuration_Template_Group_Reference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent_Configuration_Template_Group_Reference) GetConfigurationTemplate() (resp datatypes.Configuration_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent_Configuration_Template_Group_Reference) GetObject() (resp datatypes.Monitoring_Agent_Configuration_Template_Group_Reference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent_Configuration_Template_Group_Reference) GetTemplateGroup() (resp datatypes.Monitoring_Agent_Configuration_Template_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Monitoring_Agent_Configuration_Value struct {
	Session *Session
	Options
}

func (r *Session) GetMonitoringAgentConfigurationValueService() Monitoring_Agent_Configuration_Value {
	return Monitoring_Agent_Configuration_Value{Session: r}
}

func (r *Monitoring_Agent_Configuration_Value) GetDefinition() (resp datatypes.Configuration_Template_Section_Definition, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent_Configuration_Value) GetMetricDataType() (resp datatypes.Container_Metric_Data_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent_Configuration_Value) GetMonitoringAgent() (resp datatypes.Monitoring_Agent, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent_Configuration_Value) GetObject() (resp datatypes.Monitoring_Agent_Configuration_Value, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Agent_Configuration_Value) GetProfile() (resp datatypes.Configuration_Template_Section_Profile, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Monitoring_Agent_Status struct {
	Session *Session
	Options
}

func (r *Session) GetMonitoringAgentStatusService() Monitoring_Agent_Status {
	return Monitoring_Agent_Status{Session: r}
}

func (r *Monitoring_Agent_Status) GetObject() (resp datatypes.Monitoring_Agent_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Monitoring_Robot struct {
	Session *Session
	Options
}

func (r *Session) GetMonitoringRobotService() Monitoring_Robot {
	return Monitoring_Robot{Session: r}
}

func (r *Monitoring_Robot) CheckConnection() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Robot) DeployMonitoringAgents(configurationTemplateGroup *datatypes.Monitoring_Agent_Configuration_Template_Group) (resp datatypes.Provisioning_Version1_Transaction, err error) {
	params := []interface{}{
		configurationTemplateGroup,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Robot) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Robot) GetAvailableConfigurationGroups() (resp []datatypes.Monitoring_Agent_Configuration_Template_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Robot) GetMonitoringAgents() (resp []datatypes.Monitoring_Agent, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Robot) GetObject() (resp datatypes.Monitoring_Robot, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Robot) GetRobotStatus() (resp datatypes.Monitoring_Robot_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Robot) GetSoftwareComponent() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Monitoring_Robot) ResetStatus() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
