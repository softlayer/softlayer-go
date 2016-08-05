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

type McAfee_Epolicy_Orchestrator_Version36_Agent_Details struct {
	Entity

	AgentVersion  *string                                                     `json:"agentVersion,omitempty"`
	CurrentPolicy *McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details `json:"currentPolicy,omitempty"`
	LastUpdate    *string                                                     `json:"lastUpdate,omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details struct {
	Entity

	CurrentPolicy *McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details `json:"currentPolicy,omitempty"`
	Name          *string                                                     `json:"name,omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event struct {
	Entity

	EventLocalDateTime *Time                                                                     `json:"eventLocalDateTime,omitempty"`
	Filename           *string                                                                   `json:"filename,omitempty"`
	VirusActionTaken   *McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_Filter_Description `json:"virusActionTaken,omitempty"`
	VirusName          *string                                                                   `json:"virusName,omitempty"`
	VirusType          *string                                                                   `json:"virusType,omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_AccessProtection struct {
	Entity

	EventLocalDateTime *Time   `json:"eventLocalDateTime,omitempty"`
	Filename           *string `json:"filename,omitempty"`
	ProcessName        *string `json:"processName,omitempty"`
	RuleName           *string `json:"ruleName,omitempty"`
	Source             *string `json:"source,omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_Filter_Description struct {
	Entity

	Name *string `json:"name,omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_BlockedApplicationEvent struct {
	Entity

	ApplicationDescription *string `json:"applicationDescription,omitempty"`
	IncidentTime           *Time   `json:"incidentTime,omitempty"`
	ProcessName            *string `json:"processName,omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_Event_Signature struct {
	Entity

	SignatureName *string `json:"signatureName,omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_IPSEvent struct {
	Entity

	IncidentTime    *Time                                                                `json:"incidentTime,omitempty"`
	ProcessName     *string                                                              `json:"processName,omitempty"`
	ReactionText    *string                                                              `json:"reactionText,omitempty"`
	RemoteIpAddress *string                                                              `json:"remoteIpAddress,omitempty"`
	SeverityText    *string                                                              `json:"severityText,omitempty"`
	Signature       *McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_Event_Signature `json:"signature,omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_BlockedApplicationEvent struct {
	Entity

	ApplicationDescription *string `json:"applicationDescription,omitempty"`
	IncidentTime           *Time   `json:"incidentTime,omitempty"`
	ProcessName            *string `json:"processName,omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_Event_Signature struct {
	Entity

	SignatureName *string `json:"signatureName,omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent struct {
	Entity

	IncidentTime    *Time                                                                `json:"incidentTime,omitempty"`
	ProcessName     *string                                                              `json:"processName,omitempty"`
	ReactionText    *string                                                              `json:"reactionText,omitempty"`
	RemoteIpAddress *string                                                              `json:"remoteIpAddress,omitempty"`
	SeverityText    *string                                                              `json:"severityText,omitempty"`
	Signature       *McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_Event_Signature `json:"signature,omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version36_Policy_Object struct {
	Entity

	Name *string `json:"name,omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version36_Product_Properties struct {
	Entity

	DatVersion *string `json:"datVersion,omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version45_Agent_Details struct {
	Entity

	AgentVersion *string `json:"agentVersion,omitempty"`
	LastUpdate   *Time   `json:"lastUpdate,omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details struct {
	Entity

	AgentDetails *McAfee_Epolicy_Orchestrator_Version45_Agent_Details         `json:"agentDetails,omitempty"`
	Name         *string                                                      `json:"name,omitempty"`
	Policies     []McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details `json:"policies,omitempty"`
	PolicyCount  *uint                                                        `json:"policyCount,omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version45_Event struct {
	Entity

	AgentDetails        *McAfee_Epolicy_Orchestrator_Version45_Agent_Details            `json:"agentDetails,omitempty"`
	DetectedUtc         *Time                                                           `json:"detectedUtc,omitempty"`
	SourceIpv4          *string                                                         `json:"sourceIpv4,omitempty"`
	SourceProcessName   *string                                                         `json:"sourceProcessName,omitempty"`
	TargetFilename      *string                                                         `json:"targetFilename,omitempty"`
	ThreatActionTaken   *string                                                         `json:"threatActionTaken,omitempty"`
	ThreatName          *string                                                         `json:"threatName,omitempty"`
	ThreatSeverityLabel *string                                                         `json:"threatSeverityLabel,omitempty"`
	ThreatType          *string                                                         `json:"threatType,omitempty"`
	VirusActionTaken    *McAfee_Epolicy_Orchestrator_Version45_Event_Filter_Description `json:"virusActionTaken,omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version45_Event_Filter_Description struct {
	Entity

	Name *string `json:"name,omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version45_Event_Version7 struct {
	McAfee_Epolicy_Orchestrator_Version45_Event

	Signature *McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version7 `json:"signature,omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version45_Event_Version8 struct {
	McAfee_Epolicy_Orchestrator_Version45_Event

	Signature *McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version8 `json:"signature,omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version7 struct {
	Entity

	SignatureName *string `json:"signatureName,omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version8 struct {
	Entity

	SignatureName *string `json:"signatureName,omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version45_Policy_Object struct {
	Entity

	Name *string `json:"name,omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version45_Product_Properties struct {
	Entity

	DatVersion *string `json:"datVersion,omitempty"`
}
