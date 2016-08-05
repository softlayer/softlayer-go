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

type Ticket struct {
	Entity

	Account                              *Account                            `json:"account,omitempty"`
	AccountId                            *int                                `json:"accountId,omitempty"`
	AssignedAgentCount                   *uint                               `json:"assignedAgentCount,omitempty"`
	AssignedAgents                       []User_Customer                     `json:"assignedAgents,omitempty"`
	AssignedUser                         *User_Customer                      `json:"assignedUser,omitempty"`
	AssignedUserId                       *int                                `json:"assignedUserId,omitempty"`
	AttachedAdditionalEmailCount         *uint                               `json:"attachedAdditionalEmailCount,omitempty"`
	AttachedAdditionalEmails             []User_Customer_AdditionalEmail     `json:"attachedAdditionalEmails,omitempty"`
	AttachedFileCount                    *uint                               `json:"attachedFileCount,omitempty"`
	AttachedFiles                        []Ticket_Attachment_File            `json:"attachedFiles,omitempty"`
	AttachedHardware                     []Hardware                          `json:"attachedHardware,omitempty"`
	AttachedHardwareCount                *uint                               `json:"attachedHardwareCount,omitempty"`
	AttachedResourceCount                *uint                               `json:"attachedResourceCount,omitempty"`
	AttachedResources                    []Ticket_Attachment                 `json:"attachedResources,omitempty"`
	AttachedVirtualGuestCount            *uint                               `json:"attachedVirtualGuestCount,omitempty"`
	AttachedVirtualGuests                []Virtual_Guest                     `json:"attachedVirtualGuests,omitempty"`
	AwaitingUserResponseFlag             *bool                               `json:"awaitingUserResponseFlag,omitempty"`
	BillableFlag                         *bool                               `json:"billableFlag,omitempty"`
	CancellationRequest                  *Billing_Item_Cancellation_Request  `json:"cancellationRequest,omitempty"`
	ChangeOwnerFlag                      *bool                               `json:"changeOwnerFlag,omitempty"`
	CreateDate                           *Time                               `json:"createDate,omitempty"`
	EmployeeAttachmentCount              *uint                               `json:"employeeAttachmentCount,omitempty"`
	EmployeeAttachments                  []User_Employee                     `json:"employeeAttachments,omitempty"`
	FinalComments                        *string                             `json:"finalComments,omitempty"`
	FirstAttachedResource                *Ticket_Attachment                  `json:"firstAttachedResource,omitempty"`
	FirstUpdate                          *Ticket_Update                      `json:"firstUpdate,omitempty"`
	Group                                *Ticket_Group                       `json:"group,omitempty"`
	GroupId                              *int                                `json:"groupId,omitempty"`
	Id                                   *int                                `json:"id,omitempty"`
	InvoiceItemCount                     *uint                               `json:"invoiceItemCount,omitempty"`
	InvoiceItems                         []Billing_Invoice_Item              `json:"invoiceItems,omitempty"`
	LastActivity                         *Ticket_Activity                    `json:"lastActivity,omitempty"`
	LastEditDate                         *Time                               `json:"lastEditDate,omitempty"`
	LastEditType                         *string                             `json:"lastEditType,omitempty"`
	LastEditor                           *User_Interface                     `json:"lastEditor,omitempty"`
	LastResponseDate                     *Time                               `json:"lastResponseDate,omitempty"`
	LastUpdate                           *Ticket_Update                      `json:"lastUpdate,omitempty"`
	LastViewedDate                       *Time                               `json:"lastViewedDate,omitempty"`
	Location                             *Location                           `json:"location,omitempty"`
	LocationId                           *int                                `json:"locationId,omitempty"`
	ModifyDate                           *Time                               `json:"modifyDate,omitempty"`
	NewUpdatesFlag                       *bool                               `json:"newUpdatesFlag,omitempty"`
	NotifyUserOnUpdateFlag               *bool                               `json:"notifyUserOnUpdateFlag,omitempty"`
	OriginatingIpAddress                 *string                             `json:"originatingIpAddress,omitempty"`
	Priority                             *int                                `json:"priority,omitempty"`
	ResponsibleBrandId                   *int                                `json:"responsibleBrandId,omitempty"`
	ScheduledActionCount                 *uint                               `json:"scheduledActionCount,omitempty"`
	ScheduledActions                     []Provisioning_Version1_Transaction `json:"scheduledActions,omitempty"`
	ServerAdministrationBillingAmount    *int                                `json:"serverAdministrationBillingAmount,omitempty"`
	ServerAdministrationBillingInvoice   *Billing_Invoice                    `json:"serverAdministrationBillingInvoice,omitempty"`
	ServerAdministrationBillingInvoiceId *int                                `json:"serverAdministrationBillingInvoiceId,omitempty"`
	ServerAdministrationFlag             *int                                `json:"serverAdministrationFlag,omitempty"`
	ServerAdministrationRefundInvoice    *Billing_Invoice                    `json:"serverAdministrationRefundInvoice,omitempty"`
	ServerAdministrationRefundInvoiceId  *int                                `json:"serverAdministrationRefundInvoiceId,omitempty"`
	ServiceProvider                      *Service_Provider                   `json:"serviceProvider,omitempty"`
	ServiceProviderId                    *int                                `json:"serviceProviderId,omitempty"`
	ServiceProviderResourceId            *int                                `json:"serviceProviderResourceId,omitempty"`
	State                                []Ticket_State                      `json:"state,omitempty"`
	StateCount                           *uint                               `json:"stateCount,omitempty"`
	Status                               *Ticket_Status                      `json:"status,omitempty"`
	StatusId                             *int                                `json:"statusId,omitempty"`
	Subject                              *Ticket_Subject                     `json:"subject,omitempty"`
	SubjectId                            *int                                `json:"subjectId,omitempty"`
	TagReferenceCount                    *uint                               `json:"tagReferenceCount,omitempty"`
	TagReferences                        []Tag_Reference                     `json:"tagReferences,omitempty"`
	Title                                *string                             `json:"title,omitempty"`
	TotalUpdateCount                     *int                                `json:"totalUpdateCount,omitempty"`
	UpdateCount                          *uint                               `json:"updateCount,omitempty"`
	Updates                              []Ticket_Update                     `json:"updates,omitempty"`
	UserEditableFlag                     *bool                               `json:"userEditableFlag,omitempty"`
}

type Ticket_Activity struct {
	Entity

	CreateDate      *Time           `json:"createDate,omitempty"`
	CreateTimestamp *Time           `json:"createTimestamp,omitempty"`
	Editor          *User_Interface `json:"editor,omitempty"`
	Id              *int            `json:"id,omitempty"`
	Ticket          *Ticket         `json:"ticket,omitempty"`
	TicketUpdate    *Ticket_Update  `json:"ticketUpdate,omitempty"`
	Value           *string         `json:"value,omitempty"`
}

type Ticket_Attachment struct {
	Entity

	AssignedAgent   *User_Customer                     `json:"assignedAgent,omitempty"`
	AttachmentId    *int                               `json:"attachmentId,omitempty"`
	CreateDate      *Time                              `json:"createDate,omitempty"`
	Id              *int                               `json:"id,omitempty"`
	ScheduledAction *Provisioning_Version1_Transaction `json:"scheduledAction,omitempty"`
	Ticket          *Ticket                            `json:"ticket,omitempty"`
	TicketId        *int                               `json:"ticketId,omitempty"`
}

type Ticket_Attachment_Assigned_Agent struct {
	Ticket_Attachment

	AssignedAgentId *int           `json:"assignedAgentId,omitempty"`
	Resource        *User_Customer `json:"resource,omitempty"`
}

type Ticket_Attachment_CardChangeRequest struct {
	Ticket_Attachment

	Resource *Billing_Payment_Card_ChangeRequest `json:"resource,omitempty"`
}

type Ticket_Attachment_File struct {
	Entity

	CreateDate   *Time          `json:"createDate,omitempty"`
	FileName     *string        `json:"fileName,omitempty"`
	FileSize     *string        `json:"fileSize,omitempty"`
	Id           *int           `json:"id,omitempty"`
	ModifyDate   *Time          `json:"modifyDate,omitempty"`
	Ticket       *Ticket        `json:"ticket,omitempty"`
	TicketId     *int           `json:"ticketId,omitempty"`
	Update       *Ticket_Update `json:"update,omitempty"`
	UpdateId     *int           `json:"updateId,omitempty"`
	UploaderId   *string        `json:"uploaderId,omitempty"`
	UploaderType *string        `json:"uploaderType,omitempty"`
}

type Ticket_Attachment_Hardware struct {
	Ticket_Attachment

	Hardware   *Hardware `json:"hardware,omitempty"`
	HardwareId *int      `json:"hardwareId,omitempty"`
	Resource   *Hardware `json:"resource,omitempty"`
}

type Ticket_Attachment_ManualPayment struct {
	Ticket_Attachment

	Resource *Billing_Payment_Card_ManualPayment `json:"resource,omitempty"`
}

type Ticket_Attachment_Scheduled_Action struct {
	Ticket_Attachment

	Resource      *Provisioning_Version1_Transaction `json:"resource,omitempty"`
	RunDate       *Time                              `json:"runDate,omitempty"`
	Transaction   *Provisioning_Version1_Transaction `json:"transaction,omitempty"`
	TransactionId *int                               `json:"transactionId,omitempty"`
}

type Ticket_Attachment_Virtual_Guest struct {
	Ticket_Attachment

	Resource       *Virtual_Guest `json:"resource,omitempty"`
	VirtualGuest   *Virtual_Guest `json:"virtualGuest,omitempty"`
	VirtualGuestId *int           `json:"virtualGuestId,omitempty"`
}

type Ticket_Chat struct {
	Entity

	Agent        *User_Employee      `json:"agent,omitempty"`
	Customer     *User_Customer      `json:"customer,omitempty"`
	CustomerId   *int                `json:"customerId,omitempty"`
	EndDate      *Time               `json:"endDate,omitempty"`
	StartDate    *Time               `json:"startDate,omitempty"`
	TicketUpdate *Ticket_Update_Chat `json:"ticketUpdate,omitempty"`
	Transcript   *string             `json:"transcript,omitempty"`
}

type Ticket_Chat_Liveperson struct {
	Ticket_Chat
}

type Ticket_Chat_TranscriptLine struct {
	Entity

	Speaker *User_Interface `json:"speaker,omitempty"`
}

type Ticket_Chat_TranscriptLine_Customer struct {
	Ticket_Chat_TranscriptLine
}

type Ticket_Chat_TranscriptLine_Employee struct {
	Ticket_Chat_TranscriptLine
}

type Ticket_Group struct {
	Entity

	AssignedBrandCount    *uint                  `json:"assignedBrandCount,omitempty"`
	AssignedBrands        []Brand                `json:"assignedBrands,omitempty"`
	Category              *Ticket_Group_Category `json:"category,omitempty"`
	Id                    *int                   `json:"id,omitempty"`
	Name                  *string                `json:"name,omitempty"`
	TicketGroupCategoryId *int                   `json:"ticketGroupCategoryId,omitempty"`
}

type Ticket_Group_Category struct {
	Entity

	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Ticket_Priority struct {
	Entity
}

type Ticket_State struct {
	Entity

	Id          *int               `json:"id,omitempty"`
	StateType   *Ticket_State_Type `json:"stateType,omitempty"`
	StateTypeId *int               `json:"stateTypeId,omitempty"`
	Ticket      *Ticket            `json:"ticket,omitempty"`
	TicketId    *int               `json:"ticketId,omitempty"`
}

type Ticket_State_Type struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	KeyName     *string `json:"keyName,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Ticket_Status struct {
	Entity

	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Ticket_Subject struct {
	Entity

	Group *Ticket_Group `json:"group,omitempty"`
	Id    *int          `json:"id,omitempty"`
	Name  *string       `json:"name,omitempty"`
}

type Ticket_Survey struct {
	Entity
}

type Ticket_Type struct {
	Entity

	Id      *int    `json:"id,omitempty"`
	KeyName *string `json:"keyName,omitempty"`
}

type Ticket_Update struct {
	Entity

	ChangeOwnerActivity *string                  `json:"changeOwnerActivity,omitempty"`
	CreateDate          *Time                    `json:"createDate,omitempty"`
	Editor              *User_Interface          `json:"editor,omitempty"`
	EditorId            *int                     `json:"editorId,omitempty"`
	EditorType          *string                  `json:"editorType,omitempty"`
	Entry               *string                  `json:"entry,omitempty"`
	FileAttachment      []Ticket_Attachment_File `json:"fileAttachment,omitempty"`
	FileAttachmentCount *uint                    `json:"fileAttachmentCount,omitempty"`
	Id                  *int                     `json:"id,omitempty"`
	Ticket              *Ticket                  `json:"ticket,omitempty"`
	TicketId            *int                     `json:"ticketId,omitempty"`
	Type                *Ticket_Update_Type      `json:"type,omitempty"`
}

type Ticket_Update_Agent struct {
	Ticket_Update
}

type Ticket_Update_Chat struct {
	Ticket_Update

	Chat *Ticket_Chat_Liveperson `json:"chat,omitempty"`
}

type Ticket_Update_Customer struct {
	Ticket_Update
}

type Ticket_Update_Employee struct {
	Ticket_Update

	ResponseRating *int `json:"responseRating,omitempty"`
}

type Ticket_Update_Type struct {
	Entity

	Description *string        `json:"description,omitempty"`
	KeyName     *string        `json:"keyName,omitempty"`
	Ticket      *Ticket_Update `json:"ticket,omitempty"`
}
