package softlayer

import (
	"time"

	"github.ibm.com/riethm/gopherlayer/datatypes"
)

type Provisioning_Hook interface {
	Entity

	CreateObject(templateObject datatypes.Provisioning_Hook) (datatypes.Provisioning_Hook, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Provisioning_Hook) (bool, error)
	GetObject() (datatypes.Provisioning_Hook, error)

	GetAccount() (datatypes.Account, error)
	GetHookType() (datatypes.Provisioning_Hook_Type, error)
}

type Provisioning_Hook_Type interface {
	Entity

	GetAllHookTypes() (datatypes.Provisioning_Hook_Type, error)
	GetObject() (datatypes.Provisioning_Hook_Type, error)
}

type Provisioning_Maintenance_Classification interface {
	Entity

	GetMaintenanceClassification(maintenanceClassificationId int) (datatypes.Provisioning_Maintenance_Classification, error)
	GetMaintenanceClassificationsByItemCategory() (datatypes.Provisioning_Maintenance_Classification_Item_Category, error)
	GetObject() (datatypes.Provisioning_Maintenance_Classification, error)

	GetItemCategories() (datatypes.Provisioning_Maintenance_Classification_Item_Category, error)
}

type Provisioning_Maintenance_Classification_Item_Category interface {
	Entity

	GetObject() (datatypes.Provisioning_Maintenance_Classification_Item_Category, error)

	GetMaintenanceClassification() (datatypes.Provisioning_Maintenance_Classification, error)
}

type Provisioning_Maintenance_Slots interface {
	Entity

	GetObject() (datatypes.Provisioning_Maintenance_Slots, error)
}

type Provisioning_Maintenance_Ticket interface {
	Entity

	GetObject() (datatypes.Provisioning_Maintenance_Ticket, error)

	GetAvailableSlots() (datatypes.Provisioning_Maintenance_Slots, error)
	GetMaintenanceClass() (datatypes.Provisioning_Maintenance_Classification, error)
	GetTicket() (datatypes.Ticket, error)
}

type Provisioning_Maintenance_Window interface {
	Entity

	AddCustomerUpgradeWindow(customerUpgradeWindow datatypes.Container_Provisioning_Maintenance_Window) (bool, error)
	GetMaintenanceClassifications() (datatypes.Provisioning_Maintenance_Classification, error)
	GetMaintenanceStartEndTime(ticketId int) (datatypes.Provisioning_Maintenance_Window, error)
	GetMaintenanceWindowForTicket(maintenanceWindowId int) (datatypes.Provisioning_Maintenance_Window, error)
	GetMaintenanceWindowTicketsByTicketId(ticketId int) (datatypes.Provisioning_Maintenance_Ticket, error)
	GetMaintenanceWindows(beginDate time.Time, endDate time.Time, locationId int, slotsNeeded int) (datatypes.Provisioning_Maintenance_Window, error)
	GetMaintenceWindows(beginDate time.Time, endDate time.Time, locationId int, slotsNeeded int) (datatypes.Provisioning_Maintenance_Window, error)
	UpdateCustomerUpgradeWindow(maintenanceStartTime time.Time, newMaintenanceWindowId int, ticketId int) (bool, error)
}

type Provisioning_Version1_Transaction interface {
	Entity

	GetAccount() (datatypes.Account, error)
	GetGuest() (datatypes.Virtual_Guest, error)
	GetHardware() (datatypes.Hardware, error)
	GetLoopback() (datatypes.Provisioning_Version1_Transaction, error)
	GetPendingTransactions() (datatypes.Provisioning_Version1_Transaction, error)
	GetTicketScheduledActionReference() (datatypes.Ticket_Attachment, error)
	GetTransactionGroup() (datatypes.Provisioning_Version1_Transaction_Group, error)
	GetTransactionStatus() (datatypes.Provisioning_Version1_Transaction_Status, error)
}

type Provisioning_Version1_Transaction_Group interface {
	Entity

	GetAllObjects() (datatypes.Provisioning_Version1_Transaction_Group, error)
	GetObject() (datatypes.Provisioning_Version1_Transaction_Group, error)
}

type Provisioning_Version1_Transaction_History interface {
	Entity

	GetGuest() (datatypes.Virtual_Guest, error)
	GetHardware() (datatypes.Hardware, error)
	GetTransaction() (datatypes.Provisioning_Version1_Transaction, error)
	GetTransactionStatus() (datatypes.Provisioning_Version1_Transaction_Status, error)
}

type Provisioning_Version1_Transaction_Status interface {
	Entity

	GetNonCompletedTransactions() (datatypes.Provisioning_Version1_Transaction, error)
}

type Provisioning_Version1_Transaction_SubnetMigration interface {
	Provisioning_Version1_Transaction
}
