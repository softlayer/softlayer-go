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

type Ticket struct {
	Session *Session
	Options
}

func (r *Session) GetTicketService() Ticket {
	return Ticket{Session: r}
}

func (r *Ticket) AddAssignedAgent(agentId *int) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		agentId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) AddAttachedAdditionalEmails(emails []string) (resp bool, err error) {
	params := []interface{}{
		emails,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) AddAttachedFile(fileAttachment *datatypes.Container_Utility_File_Attachment) (resp datatypes.Ticket_Attachment_File, err error) {
	params := []interface{}{
		fileAttachment,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) AddAttachedHardware(hardwareId *int) (resp datatypes.Ticket_Attachment_Hardware, err error) {
	params := []interface{}{
		hardwareId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) AddAttachedVirtualGuest(guestId *int) (resp datatypes.Ticket_Attachment_Virtual_Guest, err error) {
	params := []interface{}{
		guestId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) AddFinalComments(finalComments *string) (resp bool, err error) {
	params := []interface{}{
		finalComments,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) AddScheduledAlert(activationTime *string) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		activationTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) AddScheduledAutoClose(activationTime *string) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		activationTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) AddUpdate(templateObject *datatypes.Ticket_Update, attachedFiles []datatypes.Container_Utility_File_Attachment) (resp []datatypes.Ticket_Update, err error) {
	params := []interface{}{
		templateObject,
		attachedFiles,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) CreateAdministrativeTicket(templateObject *datatypes.Ticket, contents *string, attachmentId *int, rootPassword *string, controlPanelPassword *string, accessPort *string, attachedFiles []datatypes.Container_Utility_File_Attachment, attachmentType *string) (resp datatypes.Ticket, err error) {
	params := []interface{}{
		templateObject,
		contents,
		attachmentId,
		rootPassword,
		controlPanelPassword,
		accessPort,
		attachedFiles,
		attachmentType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) CreateCancelServerTicket(attachmentId *int, reason *string, content *string, cancelAssociatedItems *bool, attachmentType *string) (resp datatypes.Ticket, err error) {
	params := []interface{}{
		attachmentId,
		reason,
		content,
		cancelAssociatedItems,
		attachmentType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) CreateCancelServiceTicket(attachmentId *int, reason *string, content *string, attachmentType *string) (resp datatypes.Ticket, err error) {
	params := []interface{}{
		attachmentId,
		reason,
		content,
		attachmentType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) CreateStandardTicket(templateObject *datatypes.Ticket, contents *string, attachmentId *int, rootPassword *string, controlPanelPassword *string, accessPort *string, attachedFiles []datatypes.Container_Utility_File_Attachment, attachmentType *string) (resp datatypes.Ticket, err error) {
	params := []interface{}{
		templateObject,
		contents,
		attachmentId,
		rootPassword,
		controlPanelPassword,
		accessPort,
		attachedFiles,
		attachmentType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) CreateUpgradeTicket(attachmentId *int, genericUpgrade *string, upgradeMaintenanceWindow *string, details *string, attachmentType *string, title *string) (resp datatypes.Ticket, err error) {
	params := []interface{}{
		attachmentId,
		genericUpgrade,
		upgradeMaintenanceWindow,
		details,
		attachmentType,
		title,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) Edit(templateObject *datatypes.Ticket, contents *string, attachedFiles []datatypes.Container_Utility_File_Attachment) (resp datatypes.Ticket, err error) {
	params := []interface{}{
		templateObject,
		contents,
		attachedFiles,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetAllTicketGroups() (resp []datatypes.Ticket_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetAllTicketStatuses() (resp []datatypes.Ticket_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetAttachedFile(attachmentId *int) (resp []byte, err error) {
	params := []interface{}{
		attachmentId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetObject() (resp datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetTicketsClosedSinceDate(closeDate *datatypes.Time) (resp []datatypes.Ticket, err error) {
	params := []interface{}{
		closeDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) MarkAsViewed() (err error) {
	var resp datatypes.Void
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) RemoveAssignedAgent(agentId *int) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		agentId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) RemoveAttachedAdditionalEmails(emails []string) (resp bool, err error) {
	params := []interface{}{
		emails,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) RemoveAttachedHardware(hardwareId *int) (resp bool, err error) {
	params := []interface{}{
		hardwareId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) RemoveAttachedVirtualGuest(guestId *int) (resp bool, err error) {
	params := []interface{}{
		guestId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) RemoveScheduledAlert() (err error) {
	var resp datatypes.Void
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) RemoveScheduledAutoClose() (err error) {
	var resp datatypes.Void
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) SetTags(tags *string) (resp bool, err error) {
	params := []interface{}{
		tags,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) SurveyEligible() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) UpdateAttachedAdditionalEmails(emails []string) (resp bool, err error) {
	params := []interface{}{
		emails,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

func (r *Ticket) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetAssignedAgents() (resp []datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetAssignedUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetAttachedAdditionalEmails() (resp []datatypes.User_Customer_AdditionalEmail, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetAttachedFiles() (resp []datatypes.Ticket_Attachment_File, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetAttachedHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetAttachedHardwareCount() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetAttachedResources() (resp []datatypes.Ticket_Attachment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetAttachedVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetAwaitingUserResponseFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetCancellationRequest() (resp datatypes.Billing_Item_Cancellation_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetEmployeeAttachments() (resp []datatypes.User_Employee, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetFirstAttachedResource() (resp datatypes.Ticket_Attachment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetFirstUpdate() (resp datatypes.Ticket_Update, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetGroup() (resp datatypes.Ticket_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetInvoiceItems() (resp []datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetLastActivity() (resp datatypes.Ticket_Activity, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetLastEditor() (resp datatypes.User_Interface, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetLastUpdate() (resp datatypes.Ticket_Update, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetLastViewedDate() (resp datatypes.Time, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetNewUpdatesFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetScheduledActions() (resp []datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetServerAdministrationBillingInvoice() (resp datatypes.Billing_Invoice, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetServerAdministrationRefundInvoice() (resp datatypes.Billing_Invoice, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetServiceProvider() (resp datatypes.Service_Provider, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetState() (resp []datatypes.Ticket_State, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetStatus() (resp datatypes.Ticket_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetSubject() (resp datatypes.Ticket_Subject, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetTagReferences() (resp []datatypes.Tag_Reference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket) GetUpdates() (resp []datatypes.Ticket_Update, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Ticket_Attachment_File struct {
	Session *Session
	Options
}

func (r *Session) GetTicketAttachmentFileService() Ticket_Attachment_File {
	return Ticket_Attachment_File{Session: r}
}

func (r *Ticket_Attachment_File) GetExtensionWhitelist() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket_Attachment_File) GetObject() (resp datatypes.Ticket_Attachment_File, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Ticket_Attachment_File) GetTicket() (resp datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket_Attachment_File) GetUpdate() (resp datatypes.Ticket_Update, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Ticket_Priority struct {
	Session *Session
	Options
}

func (r *Session) GetTicketPriorityService() Ticket_Priority {
	return Ticket_Priority{Session: r}
}

func (r *Ticket_Priority) GetPriorities() (resp []datatypes.Container_Ticket_Priority, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Ticket_Subject struct {
	Session *Session
	Options
}

func (r *Session) GetTicketSubjectService() Ticket_Subject {
	return Ticket_Subject{Session: r}
}

func (r *Ticket_Subject) GetAllObjects() (resp []datatypes.Ticket_Subject, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket_Subject) GetObject() (resp datatypes.Ticket_Subject, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket_Subject) GetTopFiveKnowledgeLayerQuestions() (resp []datatypes.Container_KnowledgeLayer_QuestionAnswer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Ticket_Subject) GetGroup() (resp datatypes.Ticket_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Ticket_Survey struct {
	Session *Session
	Options
}

func (r *Session) GetTicketSurveyService() Ticket_Survey {
	return Ticket_Survey{Session: r}
}

func (r *Ticket_Survey) GetPreference() (resp datatypes.Container_Ticket_Survey_Preference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket_Survey) OptIn() (resp datatypes.Container_Ticket_Survey_Preference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket_Survey) OptOut() (resp datatypes.Container_Ticket_Survey_Preference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Ticket_Update_Employee struct {
	Session *Session
	Options
}

func (r *Session) GetTicketUpdateEmployeeService() Ticket_Update_Employee {
	return Ticket_Update_Employee{Session: r}
}

func (r *Ticket_Update_Employee) AddResponseRating(responseRating *int) (resp bool, err error) {
	params := []interface{}{
		responseRating,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket_Update_Employee) GetObject() (resp datatypes.Ticket_Update_Employee, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
