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

type McAfee_Epolicy_Orchestrator_Version36_Agent_Details struct {
	Session *Session
	Options
}

func (r *Session) GetMcAfeeEpolicyOrchestratorVersion36AgentDetailsService() McAfee_Epolicy_Orchestrator_Version36_Agent_Details {
	return McAfee_Epolicy_Orchestrator_Version36_Agent_Details{Session: r}
}

func (r *McAfee_Epolicy_Orchestrator_Version36_Agent_Details) GetCurrentPolicy() (resp McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details struct {
	Session *Session
	Options
}

func (r *Session) GetMcAfeeEpolicyOrchestratorVersion36AgentParentDetailsService() McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details {
	return McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details{Session: r}
}

func (r *McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details) GetCurrentPolicy() (resp McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event struct {
	Session *Session
	Options
}

func (r *Session) GetMcAfeeEpolicyOrchestratorVersion36AntivirusEventService() McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event {
	return McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event{Session: r}
}

func (r *McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event) GetVirusActionTaken() (resp McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_Filter_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_AccessProtection struct {
	Session *Session
	Options
}

func (r *Session) GetMcAfeeEpolicyOrchestratorVersion36AntivirusEventAccessProtectionService() McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_AccessProtection {
	return McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_AccessProtection{Session: r}
}

type McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_Filter_Description struct {
	Session *Session
	Options
}

func (r *Session) GetMcAfeeEpolicyOrchestratorVersion36AntivirusEventFilterDescriptionService() McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_Filter_Description {
	return McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_Filter_Description{Session: r}
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_BlockedApplicationEvent struct {
	Session *Session
	Options
}

func (r *Session) GetMcAfeeEpolicyOrchestratorVersion36HipsVersion6BlockedApplicationEventService() McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_BlockedApplicationEvent {
	return McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_BlockedApplicationEvent{Session: r}
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_Event_Signature struct {
	Session *Session
	Options
}

func (r *Session) GetMcAfeeEpolicyOrchestratorVersion36HipsVersion6EventSignatureService() McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_Event_Signature {
	return McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_Event_Signature{Session: r}
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_IPSEvent struct {
	Session *Session
	Options
}

func (r *Session) GetMcAfeeEpolicyOrchestratorVersion36HipsVersion6IPSEventService() McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_IPSEvent {
	return McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_IPSEvent{Session: r}
}

func (r *McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_IPSEvent) GetSignature() (resp McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_Event_Signature, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_BlockedApplicationEvent struct {
	Session *Session
	Options
}

func (r *Session) GetMcAfeeEpolicyOrchestratorVersion36HipsVersion7BlockedApplicationEventService() McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_BlockedApplicationEvent {
	return McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_BlockedApplicationEvent{Session: r}
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_Event_Signature struct {
	Session *Session
	Options
}

func (r *Session) GetMcAfeeEpolicyOrchestratorVersion36HipsVersion7EventSignatureService() McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_Event_Signature {
	return McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_Event_Signature{Session: r}
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent struct {
	Session *Session
	Options
}

func (r *Session) GetMcAfeeEpolicyOrchestratorVersion36HipsVersion7IPSEventService() McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent {
	return McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent{Session: r}
}

func (r *McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent) GetSignature() (resp McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_Event_Signature, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type McAfee_Epolicy_Orchestrator_Version36_Policy_Object struct {
	Session *Session
	Options
}

func (r *Session) GetMcAfeeEpolicyOrchestratorVersion36PolicyObjectService() McAfee_Epolicy_Orchestrator_Version36_Policy_Object {
	return McAfee_Epolicy_Orchestrator_Version36_Policy_Object{Session: r}
}

type McAfee_Epolicy_Orchestrator_Version36_Product_Properties struct {
	Session *Session
	Options
}

func (r *Session) GetMcAfeeEpolicyOrchestratorVersion36ProductPropertiesService() McAfee_Epolicy_Orchestrator_Version36_Product_Properties {
	return McAfee_Epolicy_Orchestrator_Version36_Product_Properties{Session: r}
}

type McAfee_Epolicy_Orchestrator_Version45_Agent_Details struct {
	Session *Session
	Options
}

func (r *Session) GetMcAfeeEpolicyOrchestratorVersion45AgentDetailsService() McAfee_Epolicy_Orchestrator_Version45_Agent_Details {
	return McAfee_Epolicy_Orchestrator_Version45_Agent_Details{Session: r}
}

type McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details struct {
	Session *Session
	Options
}

func (r *Session) GetMcAfeeEpolicyOrchestratorVersion45AgentParentDetailsService() McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details {
	return McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details{Session: r}
}

func (r *McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details) GetAgentDetails() (resp McAfee_Epolicy_Orchestrator_Version45_Agent_Details, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details) GetPolicies() (resp []McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type McAfee_Epolicy_Orchestrator_Version45_Event struct {
	Session *Session
	Options
}

func (r *Session) GetMcAfeeEpolicyOrchestratorVersion45EventService() McAfee_Epolicy_Orchestrator_Version45_Event {
	return McAfee_Epolicy_Orchestrator_Version45_Event{Session: r}
}

func (r *McAfee_Epolicy_Orchestrator_Version45_Event) GetAgentDetails() (resp McAfee_Epolicy_Orchestrator_Version45_Agent_Details, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *McAfee_Epolicy_Orchestrator_Version45_Event) GetVirusActionTaken() (resp McAfee_Epolicy_Orchestrator_Version45_Event_Filter_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type McAfee_Epolicy_Orchestrator_Version45_Event_Filter_Description struct {
	Session *Session
	Options
}

func (r *Session) GetMcAfeeEpolicyOrchestratorVersion45EventFilterDescriptionService() McAfee_Epolicy_Orchestrator_Version45_Event_Filter_Description {
	return McAfee_Epolicy_Orchestrator_Version45_Event_Filter_Description{Session: r}
}

type McAfee_Epolicy_Orchestrator_Version45_Event_Version7 struct {
	Session *Session
	Options
}

func (r *Session) GetMcAfeeEpolicyOrchestratorVersion45EventVersion7Service() McAfee_Epolicy_Orchestrator_Version45_Event_Version7 {
	return McAfee_Epolicy_Orchestrator_Version45_Event_Version7{Session: r}
}

func (r *McAfee_Epolicy_Orchestrator_Version45_Event_Version7) GetSignature() (resp McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version7, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type McAfee_Epolicy_Orchestrator_Version45_Event_Version8 struct {
	Session *Session
	Options
}

func (r *Session) GetMcAfeeEpolicyOrchestratorVersion45EventVersion8Service() McAfee_Epolicy_Orchestrator_Version45_Event_Version8 {
	return McAfee_Epolicy_Orchestrator_Version45_Event_Version8{Session: r}
}

func (r *McAfee_Epolicy_Orchestrator_Version45_Event_Version8) GetSignature() (resp McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version8, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version7 struct {
	Session *Session
	Options
}

func (r *Session) GetMcAfeeEpolicyOrchestratorVersion45HipsEventSignatureVersion7Service() McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version7 {
	return McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version7{Session: r}
}

type McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version8 struct {
	Session *Session
	Options
}

func (r *Session) GetMcAfeeEpolicyOrchestratorVersion45HipsEventSignatureVersion8Service() McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version8 {
	return McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version8{Session: r}
}

type McAfee_Epolicy_Orchestrator_Version45_Policy_Object struct {
	Session *Session
	Options
}

func (r *Session) GetMcAfeeEpolicyOrchestratorVersion45PolicyObjectService() McAfee_Epolicy_Orchestrator_Version45_Policy_Object {
	return McAfee_Epolicy_Orchestrator_Version45_Policy_Object{Session: r}
}

type McAfee_Epolicy_Orchestrator_Version45_Product_Properties struct {
	Session *Session
	Options
}

func (r *Session) GetMcAfeeEpolicyOrchestratorVersion45ProductPropertiesService() McAfee_Epolicy_Orchestrator_Version45_Product_Properties {
	return McAfee_Epolicy_Orchestrator_Version45_Product_Properties{Session: r}
}
