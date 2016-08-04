package service

import (
	"time"

	"github.ibm.com/riethm/gopherlayer/datatypes"
)

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
func (r *Ticket) GetTicketsClosedSinceDate(closeDate *time.Time) (resp []datatypes.Ticket, err error) {
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
func (r *Ticket) GetLastViewedDate() (resp time.Time, err error) {
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

type Ticket_Activity struct {
	Session *Session
	Options
}

func (r *Session) GetTicketActivityService() Ticket_Activity {
	return Ticket_Activity{Session: r}
}

func (r *Ticket_Activity) GetEditor() (resp datatypes.User_Interface, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket_Activity) GetTicket() (resp datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket_Activity) GetTicketUpdate() (resp datatypes.Ticket_Update, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Ticket_Attachment struct {
	Session *Session
	Options
}

func (r *Session) GetTicketAttachmentService() Ticket_Attachment {
	return Ticket_Attachment{Session: r}
}

func (r *Ticket_Attachment) GetAssignedAgent() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket_Attachment) GetScheduledAction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket_Attachment) GetTicket() (resp datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Ticket_Attachment_Assigned_Agent struct {
	Session *Session
	Options
}

func (r *Session) GetTicketAttachmentAssignedAgentService() Ticket_Attachment_Assigned_Agent {
	return Ticket_Attachment_Assigned_Agent{Session: r}
}

func (r *Ticket_Attachment_Assigned_Agent) GetResource() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Ticket_Attachment_CardChangeRequest struct {
	Session *Session
	Options
}

func (r *Session) GetTicketAttachmentCardChangeRequestService() Ticket_Attachment_CardChangeRequest {
	return Ticket_Attachment_CardChangeRequest{Session: r}
}

func (r *Ticket_Attachment_CardChangeRequest) GetResource() (resp datatypes.Billing_Payment_Card_ChangeRequest, err error) {
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

type Ticket_Attachment_Hardware struct {
	Session *Session
	Options
}

func (r *Session) GetTicketAttachmentHardwareService() Ticket_Attachment_Hardware {
	return Ticket_Attachment_Hardware{Session: r}
}

func (r *Ticket_Attachment_Hardware) GetHardware() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket_Attachment_Hardware) GetResource() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Ticket_Attachment_ManualPayment struct {
	Session *Session
	Options
}

func (r *Session) GetTicketAttachmentManualPaymentService() Ticket_Attachment_ManualPayment {
	return Ticket_Attachment_ManualPayment{Session: r}
}

func (r *Ticket_Attachment_ManualPayment) GetResource() (resp datatypes.Billing_Payment_Card_ManualPayment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Ticket_Attachment_Scheduled_Action struct {
	Session *Session
	Options
}

func (r *Session) GetTicketAttachmentScheduledActionService() Ticket_Attachment_Scheduled_Action {
	return Ticket_Attachment_Scheduled_Action{Session: r}
}

func (r *Ticket_Attachment_Scheduled_Action) GetResource() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket_Attachment_Scheduled_Action) GetTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Ticket_Attachment_Virtual_Guest struct {
	Session *Session
	Options
}

func (r *Session) GetTicketAttachmentVirtualGuestService() Ticket_Attachment_Virtual_Guest {
	return Ticket_Attachment_Virtual_Guest{Session: r}
}

func (r *Ticket_Attachment_Virtual_Guest) GetResource() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket_Attachment_Virtual_Guest) GetVirtualGuest() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Ticket_Chat struct {
	Session *Session
	Options
}

func (r *Session) GetTicketChatService() Ticket_Chat {
	return Ticket_Chat{Session: r}
}

func (r *Ticket_Chat) GetAgent() (resp datatypes.User_Employee, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket_Chat) GetCustomer() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket_Chat) GetTicketUpdate() (resp datatypes.Ticket_Update_Chat, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Ticket_Chat_Liveperson struct {
	Session *Session
	Options
}

func (r *Session) GetTicketChatLivepersonService() Ticket_Chat_Liveperson {
	return Ticket_Chat_Liveperson{Session: r}
}

type Ticket_Chat_TranscriptLine struct {
	Session *Session
	Options
}

func (r *Session) GetTicketChatTranscriptLineService() Ticket_Chat_TranscriptLine {
	return Ticket_Chat_TranscriptLine{Session: r}
}

func (r *Ticket_Chat_TranscriptLine) GetSpeaker() (resp datatypes.User_Interface, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Ticket_Chat_TranscriptLine_Customer struct {
	Session *Session
	Options
}

func (r *Session) GetTicketChatTranscriptLineCustomerService() Ticket_Chat_TranscriptLine_Customer {
	return Ticket_Chat_TranscriptLine_Customer{Session: r}
}

type Ticket_Chat_TranscriptLine_Employee struct {
	Session *Session
	Options
}

func (r *Session) GetTicketChatTranscriptLineEmployeeService() Ticket_Chat_TranscriptLine_Employee {
	return Ticket_Chat_TranscriptLine_Employee{Session: r}
}

type Ticket_Group struct {
	Session *Session
	Options
}

func (r *Session) GetTicketGroupService() Ticket_Group {
	return Ticket_Group{Session: r}
}

func (r *Ticket_Group) GetAssignedBrands() (resp []datatypes.Brand, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket_Group) GetCategory() (resp datatypes.Ticket_Group_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Ticket_Group_Category struct {
	Session *Session
	Options
}

func (r *Session) GetTicketGroupCategoryService() Ticket_Group_Category {
	return Ticket_Group_Category{Session: r}
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

type Ticket_State struct {
	Session *Session
	Options
}

func (r *Session) GetTicketStateService() Ticket_State {
	return Ticket_State{Session: r}
}

func (r *Ticket_State) GetStateType() (resp datatypes.Ticket_State_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket_State) GetTicket() (resp datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Ticket_State_Type struct {
	Session *Session
	Options
}

func (r *Session) GetTicketStateTypeService() Ticket_State_Type {
	return Ticket_State_Type{Session: r}
}

type Ticket_Status struct {
	Session *Session
	Options
}

func (r *Session) GetTicketStatusService() Ticket_Status {
	return Ticket_Status{Session: r}
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

type Ticket_Type struct {
	Session *Session
	Options
}

func (r *Session) GetTicketTypeService() Ticket_Type {
	return Ticket_Type{Session: r}
}

type Ticket_Update struct {
	Session *Session
	Options
}

func (r *Session) GetTicketUpdateService() Ticket_Update {
	return Ticket_Update{Session: r}
}

func (r *Ticket_Update) GetChangeOwnerActivity() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket_Update) GetEditor() (resp datatypes.User_Interface, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket_Update) GetFileAttachment() (resp []datatypes.Ticket_Attachment_File, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket_Update) GetTicket() (resp datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Ticket_Update) GetType() (resp datatypes.Ticket_Update_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Ticket_Update_Agent struct {
	Session *Session
	Options
}

func (r *Session) GetTicketUpdateAgentService() Ticket_Update_Agent {
	return Ticket_Update_Agent{Session: r}
}

type Ticket_Update_Chat struct {
	Session *Session
	Options
}

func (r *Session) GetTicketUpdateChatService() Ticket_Update_Chat {
	return Ticket_Update_Chat{Session: r}
}

func (r *Ticket_Update_Chat) GetChat() (resp datatypes.Ticket_Chat_Liveperson, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Ticket_Update_Customer struct {
	Session *Session
	Options
}

func (r *Session) GetTicketUpdateCustomerService() Ticket_Update_Customer {
	return Ticket_Update_Customer{Session: r}
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

type Ticket_Update_Type struct {
	Session *Session
	Options
}

func (r *Session) GetTicketUpdateTypeService() Ticket_Update_Type {
	return Ticket_Update_Type{Session: r}
}

func (r *Ticket_Update_Type) GetTicket() (resp datatypes.Ticket_Update, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
