package datatypes

type Brand struct {
	Entity

	Account                                 *Account                                     `json:"account:omitempty"`
	AllOwnedAccountCount                    *uint                                        `json:"allOwnedAccountCount:omitempty"`
	AllOwnedAccounts                        []Account                                    `json:"allOwnedAccounts:omitempty"`
	AllowAccountCreationFlag                *bool                                        `json:"allowAccountCreationFlag:omitempty"`
	Catalog                                 *Product_Catalog                             `json:"catalog:omitempty"`
	CatalogId                               *int                                         `json:"catalogId:omitempty"`
	ContactCount                            *uint                                        `json:"contactCount:omitempty"`
	Contacts                                []Brand_Contact                              `json:"contacts:omitempty"`
	CustomerCountryLocationRestrictionCount *uint                                        `json:"customerCountryLocationRestrictionCount:omitempty"`
	CustomerCountryLocationRestrictions     []Brand_Restriction_Location_CustomerCountry `json:"customerCountryLocationRestrictions:omitempty"`
	Distributor                             *Brand                                       `json:"distributor:omitempty"`
	DistributorChildFlag                    *bool                                        `json:"distributorChildFlag:omitempty"`
	DistributorFlag                         *string                                      `json:"distributorFlag:omitempty"`
	Hardware                                []Hardware                                   `json:"hardware:omitempty"`
	HardwareCount                           *uint                                        `json:"hardwareCount:omitempty"`
	HasAgentSupportFlag                     *bool                                        `json:"hasAgentSupportFlag:omitempty"`
	Id                                      *int                                         `json:"id:omitempty"`
	KeyName                                 *string                                      `json:"keyName:omitempty"`
	LongName                                *string                                      `json:"longName:omitempty"`
	Name                                    *string                                      `json:"name:omitempty"`
	OpenTicketCount                         *uint                                        `json:"openTicketCount:omitempty"`
	OpenTickets                             []Ticket                                     `json:"openTickets:omitempty"`
	OwnedAccountCount                       *uint                                        `json:"ownedAccountCount:omitempty"`
	OwnedAccounts                           []Account                                    `json:"ownedAccounts:omitempty"`
	TicketCount                             *uint                                        `json:"ticketCount:omitempty"`
	TicketGroupCount                        *uint                                        `json:"ticketGroupCount:omitempty"`
	TicketGroups                            []Ticket_Group                               `json:"ticketGroups:omitempty"`
	Tickets                                 []Ticket                                     `json:"tickets:omitempty"`
	UserCount                               *uint                                        `json:"userCount:omitempty"`
	Users                                   []User_Customer                              `json:"users:omitempty"`
	VirtualGuestCount                       *uint                                        `json:"virtualGuestCount:omitempty"`
	VirtualGuests                           []Virtual_Guest                              `json:"virtualGuests:omitempty"`
}

type Brand_Attribute struct {
	Entity

	Brand *Brand `json:"brand:omitempty"`
}

type Brand_Contact struct {
	Entity

	Address1           *string             `json:"address1:omitempty"`
	Address2           *string             `json:"address2:omitempty"`
	AlternatePhone     *string             `json:"alternatePhone:omitempty"`
	Brand              *Brand              `json:"brand:omitempty"`
	BrandContactType   *Brand_Contact_Type `json:"brandContactType:omitempty"`
	BrandContactTypeId *int                `json:"brandContactTypeId:omitempty"`
	City               *string             `json:"city:omitempty"`
	Country            *string             `json:"country:omitempty"`
	Email              *string             `json:"email:omitempty"`
	FaxPhone           *string             `json:"faxPhone:omitempty"`
	FirstName          *string             `json:"firstName:omitempty"`
	LastName           *string             `json:"lastName:omitempty"`
	OfficePhone        *string             `json:"officePhone:omitempty"`
	PostalCode         *string             `json:"postalCode:omitempty"`
	State              *string             `json:"state:omitempty"`
}

type Brand_Contact_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Brand_Payment_Processor struct {
	Entity

	Brand            *Brand                     `json:"brand:omitempty"`
	PaymentProcessor *Billing_Payment_Processor `json:"paymentProcessor:omitempty"`
}

type Brand_Restriction_Location_CustomerCountry struct {
	Entity

	Brand               *Brand    `json:"brand:omitempty"`
	BrandId             *int      `json:"brandId:omitempty"`
	CustomerCountryCode *string   `json:"customerCountryCode:omitempty"`
	Location            *Location `json:"location:omitempty"`
	LocationId          *int      `json:"locationId:omitempty"`
}
