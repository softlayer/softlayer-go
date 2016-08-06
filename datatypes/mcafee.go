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

/**
 * AUTOMATICALLY GENERATED CODE - DO NOT MODIFY
 */

package datatypes

// The McAfee_Epolicy_Orchestrator_Version36_Agent_Details data type represents a virus scan agent and contains details about its version.
type McAfee_Epolicy_Orchestrator_Version36_Agent_Details struct {
	Entity

	// Version number of the anti-virus scan agent.
	AgentVersion *string `json:"agentVersion,omitempty"`

	// The current anti-virus policy of an agent.
	CurrentPolicy *McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details `json:"currentPolicy,omitempty"`

	// The date of the last time the anti-virus agent checked in.
	LastUpdate *string `json:"lastUpdate,omitempty"`
}

// The McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details data type contains the name of an anti-virus policy.
type McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details struct {
	Entity

	// The current anti-virus policy of an agent.
	CurrentPolicy *McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details `json:"currentPolicy,omitempty"`

	// The name of a policy.
	Name *string `json:"name,omitempty"`
}

// The McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event data type represents a single anti-virus event. It contains details about the event such as the date the event occurred, the virus that is detected and the action that is taken.
type McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event struct {
	Entity

	// The date when an anti-virus event occurs.
	EventLocalDateTime *Time `json:"eventLocalDateTime,omitempty"`

	// Name of the file found to be infected.
	Filename *string `json:"filename,omitempty"`

	// The action taken when a virus is detected.
	VirusActionTaken *McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_Filter_Description `json:"virusActionTaken,omitempty"`

	// The name of a virus that is found.
	VirusName *string `json:"virusName,omitempty"`

	// The type of virus that is found.
	VirusType *string `json:"virusType,omitempty"`
}

// The McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_AccessProtection data type represents an access protection event. It contains details about the event such as when it occurs, the process that caused it, and the rule that triggered the event.
type McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_AccessProtection struct {
	Entity

	// The date that an access protection event occurs.
	EventLocalDateTime *Time `json:"eventLocalDateTime,omitempty"`

	// The name of the file that was protected from access.
	Filename *string `json:"filename,omitempty"`

	// The name of the process that was protected from access.
	ProcessName *string `json:"processName,omitempty"`

	// The name of the rule that triggered an access protection event.
	RuleName *string `json:"ruleName,omitempty"`

	// The IP address that caused an access protection event.
	Source *string `json:"source,omitempty"`
}

// The McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_Filter_Description data type contains the name of the rule that was triggered by an anti-virus event.
type McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_Filter_Description struct {
	Entity

	// The name of the rule that triggered an anti-virus event.
	Name *string `json:"name,omitempty"`
}

// The McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_BlockedApplicationEvent data type contains a single blocked application event. The details of the event are the time the event occurred, the process that generated the event and a brief description of the application that was blocked.
type McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_BlockedApplicationEvent struct {
	Entity

	// A brief description of the application that is blocked.
	ApplicationDescription *string `json:"applicationDescription,omitempty"`

	// The time that an application is blocked.
	IncidentTime *Time `json:"incidentTime,omitempty"`

	// The name of a process that is blocked.
	ProcessName *string `json:"processName,omitempty"`
}

// The McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_Event_Signature data type contains the signature name of a rule that generated an IPS event.
type McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_Event_Signature struct {
	Entity

	// The name of a rule that triggered an IPS event.
	SignatureName *string `json:"signatureName,omitempty"`
}

// The McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_IPSEvent data type represents a single IPS event.  It contains details about the event such as the date the event occurred, the process that generated it, the severity of the event, and the action taken.
type McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_IPSEvent struct {
	Entity

	// The time when an IPS event occurred.
	IncidentTime *Time `json:"incidentTime,omitempty"`

	// Name of the process that generated an IPS event.
	ProcessName *string `json:"processName,omitempty"`

	// The action taken because of an IPS event.
	ReactionText *string `json:"reactionText,omitempty"`

	// The IP address that generated an IPS event.
	RemoteIpAddress *string `json:"remoteIpAddress,omitempty"`

	// The severity level for an IPS event.
	SeverityText *string `json:"severityText,omitempty"`

	// The signature that generated an IPS event.
	Signature *McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_Event_Signature `json:"signature,omitempty"`
}

// The McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_BlockedApplicationEvent data type contains a single blocked application event. The details of the event are the time the event occurred, the process that generated the event and a brief description of the application that was blocked.
type McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_BlockedApplicationEvent struct {
	Entity

	// A brief description of the application that is blocked.
	ApplicationDescription *string `json:"applicationDescription,omitempty"`

	// The time that an application is blocked.
	IncidentTime *Time `json:"incidentTime,omitempty"`

	// The name of a process that is blocked.
	ProcessName *string `json:"processName,omitempty"`
}

// The McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_Event_Signature data type contains the signature name of a rule that generated an IPS event.
type McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_Event_Signature struct {
	Entity

	// The name of a rule that triggered an IPS event.
	SignatureName *string `json:"signatureName,omitempty"`
}

// The McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent data type represents a single IPS event.  It contains details about the event such as the date the event occurred, the process that generated it, the severity of the event, and the action taken.
type McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent struct {
	Entity

	// The time when an IPS event occurred.
	IncidentTime *Time `json:"incidentTime,omitempty"`

	// Name of the process that generated an IPS event.
	ProcessName *string `json:"processName,omitempty"`

	// The action taken because of an IPS event.
	ReactionText *string `json:"reactionText,omitempty"`

	// The IP address that generated an IPS event.
	RemoteIpAddress *string `json:"remoteIpAddress,omitempty"`

	// The severity level for an IPS event.
	SeverityText *string `json:"severityText,omitempty"`

	// The signature that generated an IPS event.
	Signature *McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_Event_Signature `json:"signature,omitempty"`
}

// The McAfee_Epolicy_Orchestrator_Version36_Policy_Object data type contains the name of a policy that may be assigned to a server.
type McAfee_Epolicy_Orchestrator_Version36_Policy_Object struct {
	Entity

	// The name of a policy.
	Name *string `json:"name,omitempty"`
}

// The McAfee_Epolicy_Orchestrator_Version36_Product_Properties data type contains the virus definition file version.
type McAfee_Epolicy_Orchestrator_Version36_Product_Properties struct {
	Entity

	// The virus definition file version.
	DatVersion *string `json:"datVersion,omitempty"`
}

// The McAfee_Epolicy_Orchestrator_Version45_Agent_Details data type represents a virus scan agent and contains details about its version.
type McAfee_Epolicy_Orchestrator_Version45_Agent_Details struct {
	Entity

	// Version number of the anti-virus scan agent.
	AgentVersion *string `json:"agentVersion,omitempty"`

	// The date of the last time the anti-virus agent checked in.
	LastUpdate *Time `json:"lastUpdate,omitempty"`
}

// The McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details data type contains the name of an anti-virus policy.
type McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details struct {
	Entity

	// Additional information about an agent.
	AgentDetails *McAfee_Epolicy_Orchestrator_Version45_Agent_Details `json:"agentDetails,omitempty"`

	// The name of a policy.
	Name *string `json:"name,omitempty"`

	// The current anti-virus policy of an agent.
	Policies []McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details `json:"policies,omitempty"`

	// A count of the current anti-virus policy of an agent.
	PolicyCount *uint `json:"policyCount,omitempty"`
}

// The McAfee_Epolicy_Orchestrator_Version45_Event data type represents a single event. It contains details about the event such as the date the event occurred, the virus or intrusion that is detected and the action that is taken.
type McAfee_Epolicy_Orchestrator_Version45_Event struct {
	Entity

	// Additional information about an agent.
	AgentDetails *McAfee_Epolicy_Orchestrator_Version45_Agent_Details `json:"agentDetails,omitempty"`

	// The time that an event was detected.
	DetectedUtc *Time `json:"detectedUtc,omitempty"`

	// The IP address of the source that generated an event.
	SourceIpv4 *string `json:"sourceIpv4,omitempty"`

	// The name of the process that generated an event.
	SourceProcessName *string `json:"sourceProcessName,omitempty"`

	// The name of the file that was the target of the event.
	TargetFilename *string `json:"targetFilename,omitempty"`

	// The action taken regarding a threat.
	ThreatActionTaken *string `json:"threatActionTaken,omitempty"`

	// The name of the threat.
	ThreatName *string `json:"threatName,omitempty"`

	// The textual representation of the severity level.
	ThreatSeverityLabel *string `json:"threatSeverityLabel,omitempty"`

	// The type of threat.
	ThreatType *string `json:"threatType,omitempty"`

	// The action taken when a virus is detected.
	VirusActionTaken *McAfee_Epolicy_Orchestrator_Version45_Event_Filter_Description `json:"virusActionTaken,omitempty"`
}

// The McAfee_Epolicy_Orchestrator_Version45_Event_Filter_Description data type contains the name of the rule that was triggered by an event.
type McAfee_Epolicy_Orchestrator_Version45_Event_Filter_Description struct {
	Entity

	// The name of the rule that triggered an event.
	Name *string `json:"name,omitempty"`
}

// The McAfee_Epolicy_Orchestrator_Version45_Event_Version7 data type represents a single event. It contains details about the event such as the date the event occurred, the virus or intrusion that is detected and the action that is taken.
type McAfee_Epolicy_Orchestrator_Version45_Event_Version7 struct {
	McAfee_Epolicy_Orchestrator_Version45_Event

	// The signature information for an event.
	Signature *McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version7 `json:"signature,omitempty"`
}

// The McAfee_Epolicy_Orchestrator_Version45_Event_Version8 data type represents a single event. It contains details about the event such as the date the event occurred, the virus or intrusion that is detected and the action that is taken.
type McAfee_Epolicy_Orchestrator_Version45_Event_Version8 struct {
	McAfee_Epolicy_Orchestrator_Version45_Event

	// The signature information for an event.
	Signature *McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version8 `json:"signature,omitempty"`
}

// The McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version7 data type contains the signature name of a rule that generated an IPS event.
type McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version7 struct {
	Entity

	// The name of a rule that triggered an IPS event.
	SignatureName *string `json:"signatureName,omitempty"`
}

// The McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version8 data type contains the signature name of a rule that generated an IPS event.
type McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version8 struct {
	Entity

	// The name of a rule that triggered an IPS event.
	SignatureName *string `json:"signatureName,omitempty"`
}

// The McAfee_Epolicy_Orchestrator_Version45_Policy_Object data type contains the name of a policy that may be assigned to a server.
type McAfee_Epolicy_Orchestrator_Version45_Policy_Object struct {
	Entity

	// The name of a policy.
	Name *string `json:"name,omitempty"`
}

// The McAfee_Epolicy_Orchestrator_Version45_Product_Properties data type contains the virus definition file version.
type McAfee_Epolicy_Orchestrator_Version45_Product_Properties struct {
	Entity

	// The virus definition file version.
	DatVersion *string `json:"datVersion,omitempty"`
}
