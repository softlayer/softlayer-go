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

type Provisioning_Hook struct {
	Session *Session
	Options
}

func (r *Session) GetProvisioningHookService() Provisioning_Hook {
	return Provisioning_Hook{Session: r}
}

func (r *Provisioning_Hook) CreateObject(templateObject *datatypes.Provisioning_Hook) (resp datatypes.Provisioning_Hook, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Hook) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Hook) EditObject(templateObject *datatypes.Provisioning_Hook) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Hook) GetObject() (resp datatypes.Provisioning_Hook, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Provisioning_Hook) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Hook) GetHookType() (resp datatypes.Provisioning_Hook_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Provisioning_Hook_Type struct {
	Session *Session
	Options
}

func (r *Session) GetProvisioningHookTypeService() Provisioning_Hook_Type {
	return Provisioning_Hook_Type{Session: r}
}

func (r *Provisioning_Hook_Type) GetAllHookTypes() (resp []datatypes.Provisioning_Hook_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Hook_Type) GetObject() (resp datatypes.Provisioning_Hook_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Provisioning_Maintenance_Classification struct {
	Session *Session
	Options
}

func (r *Session) GetProvisioningMaintenanceClassificationService() Provisioning_Maintenance_Classification {
	return Provisioning_Maintenance_Classification{Session: r}
}

func (r *Provisioning_Maintenance_Classification) GetMaintenanceClassification(maintenanceClassificationId *int) (resp []datatypes.Provisioning_Maintenance_Classification, err error) {
	params := []interface{}{
		maintenanceClassificationId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Maintenance_Classification) GetMaintenanceClassificationsByItemCategory() (resp []datatypes.Provisioning_Maintenance_Classification_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Maintenance_Classification) GetObject() (resp datatypes.Provisioning_Maintenance_Classification, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Provisioning_Maintenance_Classification) GetItemCategories() (resp []datatypes.Provisioning_Maintenance_Classification_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Provisioning_Maintenance_Classification_Item_Category struct {
	Session *Session
	Options
}

func (r *Session) GetProvisioningMaintenanceClassificationItemCategoryService() Provisioning_Maintenance_Classification_Item_Category {
	return Provisioning_Maintenance_Classification_Item_Category{Session: r}
}

func (r *Provisioning_Maintenance_Classification_Item_Category) GetObject() (resp datatypes.Provisioning_Maintenance_Classification_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Provisioning_Maintenance_Classification_Item_Category) GetMaintenanceClassification() (resp datatypes.Provisioning_Maintenance_Classification, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Provisioning_Maintenance_Slots struct {
	Session *Session
	Options
}

func (r *Session) GetProvisioningMaintenanceSlotsService() Provisioning_Maintenance_Slots {
	return Provisioning_Maintenance_Slots{Session: r}
}

func (r *Provisioning_Maintenance_Slots) GetObject() (resp datatypes.Provisioning_Maintenance_Slots, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Provisioning_Maintenance_Ticket struct {
	Session *Session
	Options
}

func (r *Session) GetProvisioningMaintenanceTicketService() Provisioning_Maintenance_Ticket {
	return Provisioning_Maintenance_Ticket{Session: r}
}

func (r *Provisioning_Maintenance_Ticket) GetObject() (resp datatypes.Provisioning_Maintenance_Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Provisioning_Maintenance_Ticket) GetAvailableSlots() (resp datatypes.Provisioning_Maintenance_Slots, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Maintenance_Ticket) GetMaintenanceClass() (resp datatypes.Provisioning_Maintenance_Classification, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Maintenance_Ticket) GetTicket() (resp datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Provisioning_Maintenance_Window struct {
	Session *Session
	Options
}

func (r *Session) GetProvisioningMaintenanceWindowService() Provisioning_Maintenance_Window {
	return Provisioning_Maintenance_Window{Session: r}
}

func (r *Provisioning_Maintenance_Window) AddCustomerUpgradeWindow(customerUpgradeWindow *datatypes.Container_Provisioning_Maintenance_Window) (resp bool, err error) {
	params := []interface{}{
		customerUpgradeWindow,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Maintenance_Window) GetMaintenanceClassifications() (resp []datatypes.Provisioning_Maintenance_Classification, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Maintenance_Window) GetMaintenanceStartEndTime(ticketId *int) (resp datatypes.Provisioning_Maintenance_Window, err error) {
	params := []interface{}{
		ticketId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Maintenance_Window) GetMaintenanceWindowForTicket(maintenanceWindowId *int) (resp []datatypes.Provisioning_Maintenance_Window, err error) {
	params := []interface{}{
		maintenanceWindowId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Maintenance_Window) GetMaintenanceWindowTicketsByTicketId(ticketId *int) (resp []datatypes.Provisioning_Maintenance_Ticket, err error) {
	params := []interface{}{
		ticketId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Maintenance_Window) GetMaintenanceWindows(beginDate *datatypes.Time, endDate *datatypes.Time, locationId *int, slotsNeeded *int) (resp []datatypes.Provisioning_Maintenance_Window, err error) {
	params := []interface{}{
		beginDate,
		endDate,
		locationId,
		slotsNeeded,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Maintenance_Window) GetMaintenceWindows(beginDate *datatypes.Time, endDate *datatypes.Time, locationId *int, slotsNeeded *int) (resp []datatypes.Provisioning_Maintenance_Window, err error) {
	params := []interface{}{
		beginDate,
		endDate,
		locationId,
		slotsNeeded,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Maintenance_Window) UpdateCustomerUpgradeWindow(maintenanceStartTime *datatypes.Time, newMaintenanceWindowId *int, ticketId *int) (resp bool, err error) {
	params := []interface{}{
		maintenanceStartTime,
		newMaintenanceWindowId,
		ticketId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Provisioning_Version1_Transaction_Group struct {
	Session *Session
	Options
}

func (r *Session) GetProvisioningVersion1TransactionGroupService() Provisioning_Version1_Transaction_Group {
	return Provisioning_Version1_Transaction_Group{Session: r}
}

func (r *Provisioning_Version1_Transaction_Group) GetAllObjects() (resp []datatypes.Provisioning_Version1_Transaction_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Provisioning_Version1_Transaction_Group) GetObject() (resp datatypes.Provisioning_Version1_Transaction_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
