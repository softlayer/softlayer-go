package datatypes

import "time"

type Provisioning_Hook struct {
	Entity

	Account    *Account                `json:"account:omitempty"`
	AccountId  *int                    `json:"accountId:omitempty"`
	CreateDate *time.Time              `json:"createDate:omitempty"`
	HookType   *Provisioning_Hook_Type `json:"hookType:omitempty"`
	Id         *int                    `json:"id:omitempty"`
	ModifyDate *time.Time              `json:"modifyDate:omitempty"`
	Name       *string                 `json:"name:omitempty"`
	TypeId     *int                    `json:"typeId:omitempty"`
	Uri        *string                 `json:"uri:omitempty"`
}

type Provisioning_Hook_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Provisioning_Maintenance_Classification struct {
	Entity

	Id                *int                                                    `json:"id:omitempty"`
	ItemCategories    []Provisioning_Maintenance_Classification_Item_Category `json:"itemCategories:omitempty"`
	ItemCategoryCount *uint                                                   `json:"itemCategoryCount:omitempty"`
	Slots             *int                                                    `json:"slots:omitempty"`
	Type              *string                                                 `json:"type:omitempty"`
}

type Provisioning_Maintenance_Classification_Item_Category struct {
	Entity

	ItemCategoryId              *int                                     `json:"itemCategoryId:omitempty"`
	MaintenanceClassification   *Provisioning_Maintenance_Classification `json:"maintenanceClassification:omitempty"`
	MaintenanceClassificationId *int                                     `json:"maintenanceClassificationId:omitempty"`
}

type Provisioning_Maintenance_Slots struct {
	Entity

	AvailableSlots *int `json:"availableSlots:omitempty"`
}

type Provisioning_Maintenance_Ticket struct {
	Entity

	AvailableSlots   *Provisioning_Maintenance_Slots          `json:"availableSlots:omitempty"`
	MaintClassId     *int                                     `json:"maintClassId:omitempty"`
	MaintWindowId    *int                                     `json:"maintWindowId:omitempty"`
	MaintenanceClass *Provisioning_Maintenance_Classification `json:"maintenanceClass:omitempty"`
	MaintenanceDate  *time.Time                               `json:"maintenanceDate:omitempty"`
	Ticket           *Ticket                                  `json:"ticket:omitempty"`
	TicketId         *int                                     `json:"ticketId:omitempty"`
}

type Provisioning_Maintenance_Window struct {
	Entity

	BeginDate  *time.Time `json:"beginDate:omitempty"`
	DayOfWeek  *int       `json:"dayOfWeek:omitempty"`
	EndDate    *time.Time `json:"endDate:omitempty"`
	Id         *int       `json:"id:omitempty"`
	LocationId *int       `json:"locationId:omitempty"`
	PortalTzId *int       `json:"portalTzId:omitempty"`
}

type Provisioning_Version1_Transaction struct {
	Entity

	Account                             *Account                                  `json:"account:omitempty"`
	CreateDate                          *time.Time                                `json:"createDate:omitempty"`
	ElapsedSeconds                      *int                                      `json:"elapsedSeconds:omitempty"`
	Guest                               *Virtual_Guest                            `json:"guest:omitempty"`
	GuestId                             *int                                      `json:"guestId:omitempty"`
	Hardware                            *Hardware                                 `json:"hardware:omitempty"`
	HardwareId                          *int                                      `json:"hardwareId:omitempty"`
	Id                                  *int                                      `json:"id:omitempty"`
	Loopback                            []Provisioning_Version1_Transaction       `json:"loopback:omitempty"`
	LoopbackCount                       *uint                                     `json:"loopbackCount:omitempty"`
	ModifyDate                          *time.Time                                `json:"modifyDate:omitempty"`
	PendingTransactionCount             *uint                                     `json:"pendingTransactionCount:omitempty"`
	PendingTransactions                 []Provisioning_Version1_Transaction       `json:"pendingTransactions:omitempty"`
	StatusChangeDate                    *time.Time                                `json:"statusChangeDate:omitempty"`
	TicketScheduledActionReference      []Ticket_Attachment                       `json:"ticketScheduledActionReference:omitempty"`
	TicketScheduledActionReferenceCount *uint                                     `json:"ticketScheduledActionReferenceCount:omitempty"`
	TransactionGroup                    *Provisioning_Version1_Transaction_Group  `json:"transactionGroup:omitempty"`
	TransactionStatus                   *Provisioning_Version1_Transaction_Status `json:"transactionStatus:omitempty"`
}

type Provisioning_Version1_Transaction_Group struct {
	Entity

	AverageTimeToComplete *float64 `json:"averageTimeToComplete:omitempty"`
	Name                  *string  `json:"name:omitempty"`
}

type Provisioning_Version1_Transaction_History struct {
	Entity

	FinishDate          *time.Time                                `json:"finishDate:omitempty"`
	Guest               *Virtual_Guest                            `json:"guest:omitempty"`
	GuestId             *int                                      `json:"guestId:omitempty"`
	Hardware            *Hardware                                 `json:"hardware:omitempty"`
	HardwareId          *int                                      `json:"hardwareId:omitempty"`
	HostId              *int                                      `json:"hostId:omitempty"`
	Id                  *int                                      `json:"id:omitempty"`
	StartDate           *time.Time                                `json:"startDate:omitempty"`
	Transaction         *Provisioning_Version1_Transaction        `json:"transaction:omitempty"`
	TransactionId       *int                                      `json:"transactionId:omitempty"`
	TransactionStatus   *Provisioning_Version1_Transaction_Status `json:"transactionStatus:omitempty"`
	TransactionStatusId *int                                      `json:"transactionStatusId:omitempty"`
}

type Provisioning_Version1_Transaction_Status struct {
	Entity

	AverageDuration              *float64                            `json:"averageDuration:omitempty"`
	FriendlyName                 *string                             `json:"friendlyName:omitempty"`
	Name                         *string                             `json:"name:omitempty"`
	NonCompletedTransactionCount *uint                               `json:"nonCompletedTransactionCount:omitempty"`
	NonCompletedTransactions     []Provisioning_Version1_Transaction `json:"nonCompletedTransactions:omitempty"`
}

type Provisioning_Version1_Transaction_SubnetMigration struct {
	Provisioning_Version1_Transaction
}
