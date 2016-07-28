package softlayer

import (
	"time"

	"github.ibm.com/riethm/gopherlayer/datatypes"
)

type Ticket interface {
	Entity

	AddAssignedAgent(agentId int) error
	AddAttachedAdditionalEmails(emails string) (bool, error)
	AddAttachedFile(fileAttachment datatypes.Container_Utility_File_Attachment) (datatypes.Ticket_Attachment_File, error)
	AddAttachedHardware(hardwareId int) (datatypes.Ticket_Attachment_Hardware, error)
	AddAttachedVirtualGuest(guestId int) (datatypes.Ticket_Attachment_Virtual_Guest, error)
	AddFinalComments(finalComments string) (bool, error)
	AddScheduledAlert(activationTime string) error
	AddScheduledAutoClose(activationTime string) error
	AddUpdate(templateObject datatypes.Ticket_Update, attachedFiles datatypes.Container_Utility_File_Attachment) (datatypes.Ticket_Update, error)
	CreateAdministrativeTicket(templateObject datatypes.Ticket, contents string, attachmentId int, rootPassword string, controlPanelPassword string, accessPort string, attachedFiles datatypes.Container_Utility_File_Attachment, attachmentType string) (datatypes.Ticket, error)
	CreateCancelServerTicket(attachmentId int, reason string, content string, cancelAssociatedItems bool, attachmentType string) (datatypes.Ticket, error)
	CreateCancelServiceTicket(attachmentId int, reason string, content string, attachmentType string) (datatypes.Ticket, error)
	CreateStandardTicket(templateObject datatypes.Ticket, contents string, attachmentId int, rootPassword string, controlPanelPassword string, accessPort string, attachedFiles datatypes.Container_Utility_File_Attachment, attachmentType string) (datatypes.Ticket, error)
	CreateUpgradeTicket(attachmentId int, genericUpgrade string, upgradeMaintenanceWindow string, details string, attachmentType string) (datatypes.Ticket, error)
	Edit(templateObject datatypes.Ticket, contents string, attachedFiles datatypes.Container_Utility_File_Attachment) (datatypes.Ticket, error)
	GetAllTicketGroups() (datatypes.Ticket_Group, error)
	GetAllTicketStatuses() (datatypes.Ticket_Status, error)
	GetAttachedFile(attachmentId int) ([]byte, error)
	GetObject() (datatypes.Ticket, error)
	GetTicketsClosedSinceDate(closeDate time.Time) (datatypes.Ticket, error)
	MarkAsViewed() error
	RemoveAssignedAgent(agentId int) error
	RemoveAttachedAdditionalEmails(emails string) (bool, error)
	RemoveAttachedHardware(hardwareId int) (bool, error)
	RemoveAttachedVirtualGuest(guestId int) (bool, error)
	RemoveScheduledAlert() error
	RemoveScheduledAutoClose() error
	SetTags(tags string) (bool, error)
	SurveyEligible() (bool, error)
	UpdateAttachedAdditionalEmails(emails string) (bool, error)

	GetAccount() (datatypes.Account, error)
	GetAssignedAgents() (datatypes.User_Customer, error)
	GetAssignedUser() (datatypes.User_Customer, error)
	GetAttachedAdditionalEmails() (datatypes.User_Customer_AdditionalEmail, error)
	GetAttachedFiles() (datatypes.Ticket_Attachment_File, error)
	GetAttachedHardware() (datatypes.Hardware, error)
	GetAttachedHardwareCount() (uint, error)
	GetAttachedResources() (datatypes.Ticket_Attachment, error)
	GetAttachedVirtualGuests() (datatypes.Virtual_Guest, error)
	GetAwaitingUserResponseFlag() (bool, error)
	GetCancellationRequest() (datatypes.Billing_Item_Cancellation_Request, error)
	GetEmployeeAttachments() (datatypes.User_Employee, error)
	GetFirstAttachedResource() (datatypes.Ticket_Attachment, error)
	GetFirstUpdate() (datatypes.Ticket_Update, error)
	GetGroup() (datatypes.Ticket_Group, error)
	GetInvoiceItems() (datatypes.Billing_Invoice_Item, error)
	GetLastActivity() (datatypes.Ticket_Activity, error)
	GetLastEditor() (datatypes.User_Interface, error)
	GetLastUpdate() (datatypes.Ticket_Update, error)
	GetLastViewedDate() (time.Time, error)
	GetLocation() (datatypes.Location, error)
	GetNewUpdatesFlag() (bool, error)
	GetScheduledActions() (datatypes.Provisioning_Version1_Transaction, error)
	GetServerAdministrationBillingInvoice() (datatypes.Billing_Invoice, error)
	GetServerAdministrationRefundInvoice() (datatypes.Billing_Invoice, error)
	GetServiceProvider() (datatypes.Service_Provider, error)
	GetState() (datatypes.Ticket_State, error)
	GetStatus() (datatypes.Ticket_Status, error)
	GetSubject() (datatypes.Ticket_Subject, error)
	GetTagReferences() (datatypes.Tag_Reference, error)
	GetUpdates() (datatypes.Ticket_Update, error)
}

type Ticket_Activity interface {
	Entity

	GetEditor() (datatypes.User_Interface, error)
	GetTicket() (datatypes.Ticket, error)
	GetTicketUpdate() (datatypes.Ticket_Update, error)
}

type Ticket_Attachment interface {
	Entity

	GetAssignedAgent() (datatypes.User_Customer, error)
	GetScheduledAction() (datatypes.Provisioning_Version1_Transaction, error)
	GetTicket() (datatypes.Ticket, error)
}

type Ticket_Attachment_Assigned_Agent interface {
	Ticket_Attachment

	GetResource() (datatypes.User_Customer, error)
}

type Ticket_Attachment_CardChangeRequest interface {
	Ticket_Attachment

	GetResource() (datatypes.Billing_Payment_Card_ChangeRequest, error)
}

type Ticket_Attachment_File interface {
	Entity

	GetExtensionWhitelist() (string, error)
	GetObject() (datatypes.Ticket_Attachment_File, error)

	GetTicket() (datatypes.Ticket, error)
	GetUpdate() (datatypes.Ticket_Update, error)
}

type Ticket_Attachment_Hardware interface {
	Ticket_Attachment

	GetHardware() (datatypes.Hardware, error)
	GetResource() (datatypes.Hardware, error)
}

type Ticket_Attachment_ManualPayment interface {
	Ticket_Attachment

	GetResource() (datatypes.Billing_Payment_Card_ManualPayment, error)
}

type Ticket_Attachment_Scheduled_Action interface {
	Ticket_Attachment

	GetResource() (datatypes.Provisioning_Version1_Transaction, error)
	GetTransaction() (datatypes.Provisioning_Version1_Transaction, error)
}

type Ticket_Attachment_Virtual_Guest interface {
	Ticket_Attachment

	GetResource() (datatypes.Virtual_Guest, error)
	GetVirtualGuest() (datatypes.Virtual_Guest, error)
}

type Ticket_Chat interface {
	Entity

	GetAgent() (datatypes.User_Employee, error)
	GetCustomer() (datatypes.User_Customer, error)
	GetTicketUpdate() (datatypes.Ticket_Update_Chat, error)
}

type Ticket_Chat_Liveperson interface {
	Ticket_Chat
}

type Ticket_Chat_TranscriptLine interface {
	Entity

	GetSpeaker() (datatypes.User_Interface, error)
}

type Ticket_Chat_TranscriptLine_Customer interface {
	Ticket_Chat_TranscriptLine
}

type Ticket_Chat_TranscriptLine_Employee interface {
	Ticket_Chat_TranscriptLine
}

type Ticket_Group interface {
	Entity

	GetAssignedBrands() (datatypes.Brand, error)
	GetCategory() (datatypes.Ticket_Group_Category, error)
}

type Ticket_Group_Category interface {
	Entity
}

type Ticket_Priority interface {
	Entity

	GetPriorities() (datatypes.Container_Ticket_Priority, error)
}

type Ticket_State interface {
	Entity

	GetStateType() (datatypes.Ticket_State_Type, error)
	GetTicket() (datatypes.Ticket, error)
}

type Ticket_State_Type interface {
	Entity
}

type Ticket_Status interface {
	Entity
}

type Ticket_Subject interface {
	Entity

	GetAllObjects() (datatypes.Ticket_Subject, error)
	GetObject() (datatypes.Ticket_Subject, error)
	GetTopFiveKnowledgeLayerQuestions() (datatypes.Container_KnowledgeLayer_QuestionAnswer, error)

	GetGroup() (datatypes.Ticket_Group, error)
}

type Ticket_Survey interface {
	Entity

	GetPreference() (datatypes.Container_Ticket_Survey_Preference, error)
	OptIn() (datatypes.Container_Ticket_Survey_Preference, error)
	OptOut() (datatypes.Container_Ticket_Survey_Preference, error)
}

type Ticket_Type interface {
	Entity
}

type Ticket_Update interface {
	Entity

	GetChangeOwnerActivity() (string, error)
	GetEditor() (datatypes.User_Interface, error)
	GetFileAttachment() (datatypes.Ticket_Attachment_File, error)
	GetTicket() (datatypes.Ticket, error)
	GetType() (datatypes.Ticket_Update_Type, error)
}

type Ticket_Update_Agent interface {
	Ticket_Update
}

type Ticket_Update_Chat interface {
	Ticket_Update

	GetChat() (datatypes.Ticket_Chat_Liveperson, error)
}

type Ticket_Update_Customer interface {
	Ticket_Update
}

type Ticket_Update_Employee interface {
	Ticket_Update

	AddResponseRating(responseRating int) (bool, error)
	GetObject() (datatypes.Ticket_Update_Employee, error)
}

type Ticket_Update_Type interface {
	Entity

	GetTicket() (datatypes.Ticket_Update, error)
}
