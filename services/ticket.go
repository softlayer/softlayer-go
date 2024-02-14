/**
 * Copyright 2016-2024 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
 * the License. You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed
 * on an "AS IS" BASIS,WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and limitations under the License.
 */

// AUTOMATICALLY GENERATED CODE - DO NOT MODIFY

package services

import (
	"fmt"
	"strings"

	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/sl"
)

// The SoftLayer_Ticket data type models a single SoftLayer customer support or notification ticket. Each ticket object contains references to it's updates, the user it's assigned to, the SoftLayer department and employee that it's assigned to, and any hardware objects or attached files associated with the ticket. Tickets are described in further detail on the [[SoftLayer_Ticket]] service page.
//
// To create a support ticket execute the [[SoftLayer_Ticket::createStandardTicket|createStandardTicket]] or [[SoftLayer_Ticket::createAdministrativeTicket|createAdministrativeTicket]] methods in the SoftLayer_Ticket service. To create an upgrade ticket for the SoftLayer sales group execute the [[SoftLayer_Ticket::createUpgradeTicket|createUpgradeTicket]].
type Ticket struct {
	Session session.SLSession
	Options sl.Options
}

// GetTicketService returns an instance of the Ticket SoftLayer service
func GetTicketService(sess session.SLSession) Ticket {
	return Ticket{Session: sess}
}

func (r Ticket) Id(id int) Ticket {
	r.Options.Id = &id
	return r
}

func (r Ticket) Mask(mask string) Ticket {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Ticket) Filter(filter string) Ticket {
	r.Options.Filter = filter
	return r
}

func (r Ticket) Limit(limit int) Ticket {
	r.Options.Limit = &limit
	return r
}

func (r Ticket) Offset(offset int) Ticket {
	r.Options.Offset = &offset
	return r
}

// Attach the given file to a SoftLayer ticket. A file attachment is a convenient way to submit non-textual error reports to SoftLayer employees in a ticket. File attachments to tickets must have a unique name.
func (r Ticket) AddAttachedFile(fileAttachment *datatypes.Container_Utility_File_Attachment) (resp datatypes.Ticket_Attachment_File, err error) {
	params := []interface{}{
		fileAttachment,
	}
	err = r.Session.DoRequest("SoftLayer_Ticket", "addAttachedFile", params, &r.Options, &resp)
	return
}

// Attach the given hardware to a SoftLayer ticket. A hardware attachment provides an easy way for SoftLayer's employees to quickly look up your hardware records in the case of hardware-specific issues.
func (r Ticket) AddAttachedHardware(hardwareId *int) (resp datatypes.Ticket_Attachment_Hardware, err error) {
	params := []interface{}{
		hardwareId,
	}
	err = r.Session.DoRequest("SoftLayer_Ticket", "addAttachedHardware", params, &r.Options, &resp)
	return
}

// Attach the given CloudLayer Computing Instance to a SoftLayer ticket. An attachment provides an easy way for SoftLayer's employees to quickly look up your records in the case of specific issues.
func (r Ticket) AddAttachedVirtualGuest(guestId *int, callCommit *bool) (resp datatypes.Ticket_Attachment_Virtual_Guest, err error) {
	params := []interface{}{
		guestId,
		callCommit,
	}
	err = r.Session.DoRequest("SoftLayer_Ticket", "addAttachedVirtualGuest", params, &r.Options, &resp)
	return
}

// Add an update to a ticket. A ticket update's entry has a maximum length of 4000 characters, so ”addUpdate()” splits the ”entry” property in the ”templateObject” parameter into 3900 character blocks and creates one entry per 3900 character block. Once complete ”addUpdate()” emails the ticket's owner and additional email addresses with an update message if the ticket's ”notifyUserOnUpdateFlag” is set. If the ticket is a Legal or Abuse ticket, then the account's abuse emails are also notified when the updates are processed. Finally, ”addUpdate()” returns an array of the newly created ticket updates.
func (r Ticket) AddUpdate(templateObject *datatypes.Ticket_Update, attachedFiles []datatypes.Container_Utility_File_Attachment) (resp []datatypes.Ticket_Update, err error) {
	params := []interface{}{
		templateObject,
		attachedFiles,
	}
	err = r.Session.DoRequest("SoftLayer_Ticket", "addUpdate", params, &r.Options, &resp)
	return
}

// Create a standard support ticket. Use a standard support ticket if you need to work out a problem related to SoftLayer's hardware, network, or services. If you require SoftLayer's assistance managing your server or content then please open an administrative ticket.
//
// Support tickets may only be created in the open state. The SoftLayer API defaults new ticket properties ”userEditableFlag” to true, ”accountId” to the id of the account that your API user belongs to, and ”statusId” to 1001 (or "open"). You may not assign your new to ticket to users that your API user does not have access to.
//
// Once your ticket is created it is placed in a queue for SoftLayer employees to work. As they update the ticket new [[SoftLayer_Ticket_Update]] entries are added to the ticket object.
func (r Ticket) CreateStandardTicket(templateObject *datatypes.Ticket, contents *string, attachmentId *int, rootPassword *string, controlPanelPassword *string, accessPort *string, attachedFiles []datatypes.Container_Utility_File_Attachment, attachmentType *string) (resp datatypes.Ticket, err error) {
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
	err = r.Session.DoRequest("SoftLayer_Ticket", "createStandardTicket", params, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Ticket object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Ticket service. You can only retrieve tickets that are associated with your SoftLayer customer account.
func (r Ticket) GetObject() (resp datatypes.Ticket, err error) {
	err = r.Session.DoRequest("SoftLayer_Ticket", "getObject", nil, &r.Options, &resp)
	return
}

// Retrieve A ticket's updates.
func (r Ticket) GetUpdates() (resp []datatypes.Ticket_Update, err error) {
	err = r.Session.DoRequest("SoftLayer_Ticket", "getUpdates", nil, &r.Options, &resp)
	return
}

// detach the given hardware from a SoftLayer ticket. Removing a hardware attachment may delay ticket processing time if the hardware removed is relevant to the ticket's issue. Return a boolean true upon successful hardware detachment.
func (r Ticket) RemoveAttachedHardware(hardwareId *int) (resp bool, err error) {
	params := []interface{}{
		hardwareId,
	}
	err = r.Session.DoRequest("SoftLayer_Ticket", "removeAttachedHardware", params, &r.Options, &resp)
	return
}

// Detach the given CloudLayer Computing Instance from a SoftLayer ticket. Removing an attachment may delay ticket processing time if the instance removed is relevant to the ticket's issue. Return a boolean true upon successful detachment.
func (r Ticket) RemoveAttachedVirtualGuest(guestId *int) (resp bool, err error) {
	params := []interface{}{
		guestId,
	}
	err = r.Session.DoRequest("SoftLayer_Ticket", "removeAttachedVirtualGuest", params, &r.Options, &resp)
	return
}

// The SoftLayer_Ticket_Subject data type models one of the possible subjects that a standard support ticket may belong to. A basic support ticket's title matches it's corresponding subject's name.
type Ticket_Subject struct {
	Session session.SLSession
	Options sl.Options
}

// GetTicketSubjectService returns an instance of the Ticket_Subject SoftLayer service
func GetTicketSubjectService(sess session.SLSession) Ticket_Subject {
	return Ticket_Subject{Session: sess}
}

func (r Ticket_Subject) Id(id int) Ticket_Subject {
	r.Options.Id = &id
	return r
}

func (r Ticket_Subject) Mask(mask string) Ticket_Subject {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Ticket_Subject) Filter(filter string) Ticket_Subject {
	r.Options.Filter = filter
	return r
}

func (r Ticket_Subject) Limit(limit int) Ticket_Subject {
	r.Options.Limit = &limit
	return r
}

func (r Ticket_Subject) Offset(offset int) Ticket_Subject {
	r.Options.Offset = &offset
	return r
}

// Retrieve all possible ticket subjects. The SoftLayer customer portal uses this method in the add standard support ticket form.
func (r Ticket_Subject) GetAllObjects() (resp []datatypes.Ticket_Subject, err error) {
	err = r.Session.DoRequest("SoftLayer_Ticket_Subject", "getAllObjects", nil, &r.Options, &resp)
	return
}
