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

type Monitoring_Agent struct {
	Entity

	AgentStatus               *Monitoring_Agent_Status                 `json:"agentStatus,omitempty"`
	ConfigurationProfileCount *uint                                    `json:"configurationProfileCount,omitempty"`
	ConfigurationProfiles     []Configuration_Template_Section_Profile `json:"configurationProfiles,omitempty"`
	ConfigurationTemplate     *Configuration_Template                  `json:"configurationTemplate,omitempty"`
	ConfigurationTemplateId   *int                                     `json:"configurationTemplateId,omitempty"`
	ConfigurationValueCount   *uint                                    `json:"configurationValueCount,omitempty"`
	ConfigurationValues       []Monitoring_Agent_Configuration_Value   `json:"configurationValues,omitempty"`
	CreateDate                *time.Time                               `json:"createDate,omitempty"`
	Hardware                  *Hardware                                `json:"hardware,omitempty"`
	Id                        *int                                     `json:"id,omitempty"`
	ModifyDate                *time.Time                               `json:"modifyDate,omitempty"`
	Name                      *string                                  `json:"name,omitempty"`
	ProductItem               *Product_Item                            `json:"productItem,omitempty"`
	RemoteMonitoringAgentFlag *bool                                    `json:"remoteMonitoringAgentFlag,omitempty"`
	RobotId                   *int                                     `json:"robotId,omitempty"`
	SoftwareDescription       *Software_Description                    `json:"softwareDescription,omitempty"`
	StatusId                  *int                                     `json:"statusId,omitempty"`
	StatusName                *string                                  `json:"statusName,omitempty"`
	VirtualGuest              *Virtual_Guest                           `json:"virtualGuest,omitempty"`
}

type Monitoring_Agent_Configuration_Template_Group struct {
	Entity

	Account                             *Account                                                  `json:"account,omitempty"`
	AccountId                           *int                                                      `json:"accountId,omitempty"`
	ConfigurationTemplateCount          *uint                                                     `json:"configurationTemplateCount,omitempty"`
	ConfigurationTemplateReferenceCount *uint                                                     `json:"configurationTemplateReferenceCount,omitempty"`
	ConfigurationTemplateReferences     []Monitoring_Agent_Configuration_Template_Group_Reference `json:"configurationTemplateReferences,omitempty"`
	ConfigurationTemplates              []Configuration_Template                                  `json:"configurationTemplates,omitempty"`
	CreateDate                          *time.Time                                                `json:"createDate,omitempty"`
	Description                         *string                                                   `json:"description,omitempty"`
	Id                                  *int                                                      `json:"id,omitempty"`
	Item                                *Product_Item                                             `json:"item,omitempty"`
	ItemId                              *int                                                      `json:"itemId,omitempty"`
	ModifyDate                          *time.Time                                                `json:"modifyDate,omitempty"`
	Name                                *string                                                   `json:"name,omitempty"`
}

type Monitoring_Agent_Configuration_Template_Group_Reference struct {
	Entity

	ConfigurationTemplate   *Configuration_Template                        `json:"configurationTemplate,omitempty"`
	ConfigurationTemplateId *int                                           `json:"configurationTemplateId,omitempty"`
	Id                      *int                                           `json:"id,omitempty"`
	TemplateGroup           *Monitoring_Agent_Configuration_Template_Group `json:"templateGroup,omitempty"`
	TemplateGroupId         *int                                           `json:"templateGroupId,omitempty"`
}

type Monitoring_Agent_Configuration_Value struct {
	Entity

	AgentId                   *int                                       `json:"agentId,omitempty"`
	ConfigurationDefinitionId *int                                       `json:"configurationDefinitionId,omitempty"`
	Definition                *Configuration_Template_Section_Definition `json:"definition,omitempty"`
	Description               *string                                    `json:"description,omitempty"`
	Id                        *int                                       `json:"id,omitempty"`
	MetricDataType            *Container_Metric_Data_Type                `json:"metricDataType,omitempty"`
	MonitoringAgent           *Monitoring_Agent                          `json:"monitoringAgent,omitempty"`
	Profile                   *Configuration_Template_Section_Profile    `json:"profile,omitempty"`
	ProfileId                 *int                                       `json:"profileId,omitempty"`
	Value                     *string                                    `json:"value,omitempty"`
}

type Monitoring_Agent_Status struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Monitoring_Robot struct {
	Entity

	Account              *Account                 `json:"account,omitempty"`
	AccountId            *int                     `json:"accountId,omitempty"`
	Id                   *int                     `json:"id,omitempty"`
	MonitoringAgentCount *uint                    `json:"monitoringAgentCount,omitempty"`
	MonitoringAgents     []Monitoring_Agent       `json:"monitoringAgents,omitempty"`
	Name                 *string                  `json:"name,omitempty"`
	RobotStatus          *Monitoring_Robot_Status `json:"robotStatus,omitempty"`
	SoftwareComponent    *Software_Component      `json:"softwareComponent,omitempty"`
	StatusId             *int                     `json:"statusId,omitempty"`
}

type Monitoring_Robot_Status struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
}
