package softlayer

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Brand interface {
	Entity

	CreateCustomerAccount(account datatypes.Account, bypassDuplicateAccountCheck bool) (datatypes.Account, error)
	CreateObject(templateObject datatypes.Brand) (datatypes.Brand, error)
	GetAllTicketSubjects(account datatypes.Account) (datatypes.Ticket_Subject, error)
	GetContactInformation() (datatypes.Brand_Contact, error)
	GetMerchantName() (string, error)
	GetObject() (datatypes.Brand, error)
	GetToken(userId int) (string, error)

	GetAccount() (datatypes.Account, error)
	GetAllOwnedAccounts() (datatypes.Account, error)
	GetAllowAccountCreationFlag() (bool, error)
	GetCatalog() (datatypes.Product_Catalog, error)
	GetContacts() (datatypes.Brand_Contact, error)
	GetCustomerCountryLocationRestrictions() (datatypes.Brand_Restriction_Location_CustomerCountry, error)
	GetDistributor() (datatypes.Brand, error)
	GetDistributorChildFlag() (bool, error)
	GetDistributorFlag() (string, error)
	GetHardware() (datatypes.Hardware, error)
	GetHasAgentSupportFlag() (bool, error)
	GetOpenTickets() (datatypes.Ticket, error)
	GetOwnedAccounts() (datatypes.Account, error)
	GetTicketGroups() (datatypes.Ticket_Group, error)
	GetTickets() (datatypes.Ticket, error)
	GetUsers() (datatypes.User_Customer, error)
	GetVirtualGuests() (datatypes.Virtual_Guest, error)
}

type Brand_Attribute interface {
	Entity

	GetBrand() (datatypes.Brand, error)
}

type Brand_Contact interface {
	Entity

	GetBrand() (datatypes.Brand, error)
	GetBrandContactType() (datatypes.Brand_Contact_Type, error)
}

type Brand_Contact_Type interface {
	Entity
}

type Brand_Payment_Processor interface {
	Entity

	GetBrand() (datatypes.Brand, error)
	GetPaymentProcessor() (datatypes.Billing_Payment_Processor, error)
}

type Brand_Restriction_Location_CustomerCountry interface {
	Entity

	GetAllObjects() (datatypes.Brand_Restriction_Location_CustomerCountry, error)
	GetObject() (datatypes.Brand_Restriction_Location_CustomerCountry, error)

	GetBrand() (datatypes.Brand, error)
	GetLocation() (datatypes.Location, error)
}
