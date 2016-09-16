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

/**
 * AUTOMATICALLY GENERATED CODE - DO NOT MODIFY
 */

package datatypes

// SoftLayer_Container_Account_Discount_Program models a single outbound object for a graph of given data sets.
type Container_Account_Discount_Program struct {
	Entity

	// The credit allowance that has already been applied during the current billing cycle. If the lifetime limit has been or soon will be reached, this amount may included credit applied in previous billing cycles.
	AppliedCredit *Float64 `json:"appliedCredit,omitempty" xmlrpc:"appliedCredit"`

	// Flag to signify whether the account is a participant in the discount program.
	IsParticipant *bool `json:"isParticipant,omitempty" xmlrpc:"isParticipant"`

	// Credit allowance applied over the course of the entire program enrollment. For enrollments without a lifetime restriction, this property will not be populated as credit will be tracked on a purely monthly basis.
	LifetimeAppliedCredit *Float64 `json:"lifetimeAppliedCredit,omitempty" xmlrpc:"lifetimeAppliedCredit"`

	// Credit allowance available over the course of the entire program enrollment. If null, enrollment credit is applied on a strictly monthly basis and there is no lifetime maximum. Enrollments with non-null lifetime credit will receive the lesser of the remaining monthly credit or the remaining lifetime credit.
	LifetimeCredit *Float64 `json:"lifetimeCredit,omitempty" xmlrpc:"lifetimeCredit"`

	// Remaining credit allowance available over the remaining duration of the program enrollment. If null, enrollment credit is applied on a strictly monthly basis and there is no lifetime maximum. Enrollments with non-null remaining lifetime credit will receive the lesser of the remaining monthly credit or the remaining lifetime credit.
	LifetimeRemainingCredit *Float64 `json:"lifetimeRemainingCredit,omitempty" xmlrpc:"lifetimeRemainingCredit"`

	// Maximum number of orders the enrolled account is allowed to have open at one time. If null, then the Flexible Credit Program does not impose an order limit.
	MaximumActiveOrders *Float64 `json:"maximumActiveOrders,omitempty" xmlrpc:"maximumActiveOrders"`

	// The monthly credit allowance that is available at the beginning of the billing cycle.
	MonthlyCredit *Float64 `json:"monthlyCredit,omitempty" xmlrpc:"monthlyCredit"`

	// DEPRECATED: Taxes are calculated in real time and discount amounts are shown pre-tax in all cases. Tax values in the SoftLayer_Container_Account_Discount_Program container are now populated with the related pre-tax values.
	PostTaxRemainingCredit *Float64 `json:"postTaxRemainingCredit,omitempty" xmlrpc:"postTaxRemainingCredit"`

	// The date at which the program expires in MM/DD/YYYY format.
	ProgramEndDate *Time `json:"programEndDate,omitempty" xmlrpc:"programEndDate"`

	// Name of the Flexible Credit Program the account is enrolled in.
	ProgramName *string `json:"programName,omitempty" xmlrpc:"programName"`

	// The credit allowance that is available during the current billing cycle. If the lifetime limit has been or soon will be reached, this amount may be reduced by credit applied in previous billing cycles.
	RemainingCredit *Float64 `json:"remainingCredit,omitempty" xmlrpc:"remainingCredit"`

	// DEPRECATED: Taxes are calculated in real time and discount amounts are shown pre-tax in all cases. Tax values in the SoftLayer_Container_Account_Discount_Program container are now populated with the related pre-tax values.
	RemainingCreditTax *Float64 `json:"remainingCreditTax,omitempty" xmlrpc:"remainingCreditTax"`
}

// SoftLayer_Container_Account_Graph_Outputs <<< EOT
type Container_Account_Graph_Outputs struct {
	Entity

	// The count of closed tickets included in this graph.
	ClosedTickets *string `json:"closedTickets,omitempty" xmlrpc:"closedTickets"`

	// The count of completed backups included in this graph.
	CompletedBackupCount *string `json:"completedBackupCount,omitempty" xmlrpc:"completedBackupCount"`

	// The count of conflicted backups included in this graph.
	ConflictBackupCount *string `json:"conflictBackupCount,omitempty" xmlrpc:"conflictBackupCount"`

	// The maximum date included in this graph.
	EndDate *Time `json:"endDate,omitempty" xmlrpc:"endDate"`

	// The count of failed backups included in this graph.
	FailedBackupCount *string `json:"failedBackupCount,omitempty" xmlrpc:"failedBackupCount"`

	// Error message encountered during graphing
	GraphError *string `json:"graphError,omitempty" xmlrpc:"graphError"`

	// The raw PNG binary data to be displayed once the graph is drawn.
	GraphImage *[]byte `json:"graphImage,omitempty" xmlrpc:"graphImage"`

	// The average of hardware uptime included in this graph.
	HardwareUptime *string `json:"hardwareUptime,omitempty" xmlrpc:"hardwareUptime"`

	// The inbound bandwidth usage shown in this graph.
	InboundUsage *string `json:"inboundUsage,omitempty" xmlrpc:"inboundUsage"`

	// The count of open tickets included in this graph.
	OpenTickets *string `json:"openTickets,omitempty" xmlrpc:"openTickets"`

	// The outbound bandwidth usage shown in this graph.
	OutboundUsage *string `json:"outboundUsage,omitempty" xmlrpc:"outboundUsage"`

	// The count of tickets included in this graph.
	PendingCustomerResponseTicketCount *string `json:"pendingCustomerResponseTicketCount,omitempty" xmlrpc:"pendingCustomerResponseTicketCount"`

	// The minimum date included in this graph.
	StartDate *Time `json:"startDate,omitempty" xmlrpc:"startDate"`

	// The average of url uptime included in this graph.
	UrlUptime *string `json:"urlUptime,omitempty" xmlrpc:"urlUptime"`

	// The count of tickets included in this graph.
	WaitingEmployeeResponseTicketCount *string `json:"waitingEmployeeResponseTicketCount,omitempty" xmlrpc:"waitingEmployeeResponseTicketCount"`
}

// Historical Summary Container for account resource details
type Container_Account_Historical_Summary struct {
	Entity

	// Array of server uptime detail containers
	Details []Container_Account_Historical_Summary_Detail `json:"details,omitempty" xmlrpc:"details"`

	// The maximum date included in the summary.
	EndDate *Time `json:"endDate,omitempty" xmlrpc:"endDate"`

	// The minimum date included in the summary.
	StartDate *Time `json:"startDate,omitempty" xmlrpc:"startDate"`
}

// Historical Summary Details Container for a resource's data
type Container_Account_Historical_Summary_Detail struct {
	Entity

	// The maximum date included in the detail.
	EndDate *Time `json:"endDate,omitempty" xmlrpc:"endDate"`

	// The minimum date included in the detail.
	StartDate *Time `json:"startDate,omitempty" xmlrpc:"startDate"`
}

// Historical Summary Details Container for a host resource uptime
type Container_Account_Historical_Summary_Detail_Uptime struct {
	Container_Account_Historical_Summary_Detail

	// The hardware for uptime details.
	CloudComputingInstance *Virtual_Guest `json:"cloudComputingInstance,omitempty" xmlrpc:"cloudComputingInstance"`

	// The configuration value for the detail's resource.
	ConfigurationValue *Monitoring_Agent_Configuration_Value `json:"configurationValue,omitempty" xmlrpc:"configurationValue"`

	// The data associated with a host uptime details.
	Data []Metric_Tracking_Object_Data `json:"data,omitempty" xmlrpc:"data"`

	// The hardware for uptime details.
	Hardware *Hardware `json:"hardware,omitempty" xmlrpc:"hardware"`
}

// Historical Summary Container for account host's resource uptime details
type Container_Account_Historical_Summary_Uptime struct {
	Container_Account_Historical_Summary
}

// no documentation yet
type Container_Account_Payment_Method_CreditCard struct {
	Entity

	// no documentation yet
	Address1 *string `json:"address1,omitempty" xmlrpc:"address1"`

	// no documentation yet
	Address2 *string `json:"address2,omitempty" xmlrpc:"address2"`

	// no documentation yet
	City *string `json:"city,omitempty" xmlrpc:"city"`

	// no documentation yet
	Country *string `json:"country,omitempty" xmlrpc:"country"`

	// no documentation yet
	CurrencyShortName *string `json:"currencyShortName,omitempty" xmlrpc:"currencyShortName"`

	// no documentation yet
	CybersourceAssignedCardType *string `json:"cybersourceAssignedCardType,omitempty" xmlrpc:"cybersourceAssignedCardType"`

	// no documentation yet
	ExpireMonth *string `json:"expireMonth,omitempty" xmlrpc:"expireMonth"`

	// no documentation yet
	ExpireYear *string `json:"expireYear,omitempty" xmlrpc:"expireYear"`

	// no documentation yet
	FirstName *string `json:"firstName,omitempty" xmlrpc:"firstName"`

	// no documentation yet
	LastFourDigits *string `json:"lastFourDigits,omitempty" xmlrpc:"lastFourDigits"`

	// no documentation yet
	LastName *string `json:"lastName,omitempty" xmlrpc:"lastName"`

	// no documentation yet
	Nickname *string `json:"nickname,omitempty" xmlrpc:"nickname"`

	// no documentation yet
	PaymentMethodRoleName *string `json:"paymentMethodRoleName,omitempty" xmlrpc:"paymentMethodRoleName"`

	// no documentation yet
	PaymentTypeId *string `json:"paymentTypeId,omitempty" xmlrpc:"paymentTypeId"`

	// no documentation yet
	PaymentTypeName *string `json:"paymentTypeName,omitempty" xmlrpc:"paymentTypeName"`

	// no documentation yet
	PostalCode *string `json:"postalCode,omitempty" xmlrpc:"postalCode"`

	// no documentation yet
	State *string `json:"state,omitempty" xmlrpc:"state"`
}

// The SoftLayer_Container_Auxiliary_Network_Status_Reading data type contains information relating to an object being monitored from outside the SoftLayer network.  It is primarily used to check the status of our edge routers from multiple locations around the world.
type Container_Auxiliary_Network_Status_Reading struct {
	Entity

	// Average packet round-trip time.
	AveragePing *Float64 `json:"averagePing,omitempty" xmlrpc:"averagePing"`

	// Number of failures since the target was last detected to be working properly.
	Fails *int `json:"fails,omitempty" xmlrpc:"fails"`

	// Monitoring frequency in minutes.
	Frequency *int `json:"frequency,omitempty" xmlrpc:"frequency"`

	// The target babel.
	Label *string `json:"label,omitempty" xmlrpc:"label"`

	// Last check date and time.
	LastCheckDate *Time `json:"lastCheckDate,omitempty" xmlrpc:"lastCheckDate"`

	// Date and time of the last problem detected.
	LastDownDate *Time `json:"lastDownDate,omitempty" xmlrpc:"lastDownDate"`

	// The total response time in seconds calculated during the last check.
	Latency *Float64 `json:"latency,omitempty" xmlrpc:"latency"`

	// The monitoring location name.
	Location *string `json:"location,omitempty" xmlrpc:"location"`

	// Maximum packet round-trip time.
	MaximumPing *Float64 `json:"maximumPing,omitempty" xmlrpc:"maximumPing"`

	// Minimum packet round-trip time.
	MinimumPing *Float64 `json:"minimumPing,omitempty" xmlrpc:"minimumPing"`

	// Packet loss percentage.
	PingLoss *Float64 `json:"pingLoss,omitempty" xmlrpc:"pingLoss"`

	// The date monitoring first began
	StartDate *Time `json:"startDate,omitempty" xmlrpc:"startDate"`

	// Status Code - one of UP, Down, Test pending.
	StatusCode *string `json:"statusCode,omitempty" xmlrpc:"statusCode"`

	// The status message from the last effective check.
	StatusMessage *string `json:"statusMessage,omitempty" xmlrpc:"statusMessage"`

	// The target object.
	Target *string `json:"target,omitempty" xmlrpc:"target"`

	// A letter indicating the target type.
	TargetType *string `json:"targetType,omitempty" xmlrpc:"targetType"`
}

// SoftLayer_Container_Bandwidth_GraphInputs models a single inbound object for a given bandwidth graph.
type Container_Bandwidth_GraphInputs struct {
	Entity

	// This is a unix timestamp that represents the stop date/time for a graph.
	EndDate *Time `json:"endDate,omitempty" xmlrpc:"endDate"`

	// The front-end or back-end network uplink interface associated with this server.
	NetworkInterfaceId *int `json:"networkInterfaceId,omitempty" xmlrpc:"networkInterfaceId"`

	// *
	Pod *int `json:"pod,omitempty" xmlrpc:"pod"`

	// This is a human readable name for the server or rack being graphed.
	ServerName *string `json:"serverName,omitempty" xmlrpc:"serverName"`

	// This is a unix timestamp that represents the begin date/time for a graph.
	StartDate *Time `json:"startDate,omitempty" xmlrpc:"startDate"`
}

// SoftLayer_Container_Bandwidth_GraphOutputs models a single outbound object for a given bandwidth graph.
type Container_Bandwidth_GraphOutputs struct {
	Entity

	// The raw PNG binary data to be displayed once the graph is drawn.
	GraphImage *[]byte `json:"graphImage,omitempty" xmlrpc:"graphImage"`

	// The title that ended up being displayed as part of the graph image.
	GraphTitle *string `json:"graphTitle,omitempty" xmlrpc:"graphTitle"`

	// The maximum date included in this graph.
	MaxEndDate *Time `json:"maxEndDate,omitempty" xmlrpc:"maxEndDate"`

	// The minimum date included in this graph.
	MinStartDate *Time `json:"minStartDate,omitempty" xmlrpc:"minStartDate"`
}

// SoftLayer_Container_Bandwidth_GraphOutputs models an individual bandwidth graph image and certain details about that graph image.
type Container_Bandwidth_GraphOutputsExtended struct {
	Entity

	// The raw PNG binary data of a bandwidth graph image.
	GraphImage *[]byte `json:"graphImage,omitempty" xmlrpc:"graphImage"`

	// A bandwidth graph's title.
	GraphTitle *string `json:"graphTitle,omitempty" xmlrpc:"graphTitle"`

	// The amount of inbound traffic reported on a bandwidth graph image.
	InBoundTotalBytes *uint `json:"inBoundTotalBytes,omitempty" xmlrpc:"inBoundTotalBytes"`

	// The ending date of the data represented in a bandwidth graph.
	MaxEndDate *Time `json:"maxEndDate,omitempty" xmlrpc:"maxEndDate"`

	// The beginning date of the data represented in a bandwidth graph.
	MinStartDate *Time `json:"minStartDate,omitempty" xmlrpc:"minStartDate"`

	// The amount of outbound traffic reported on a bandwidth graph image.
	OutBoundTotalBytes *uint `json:"outBoundTotalBytes,omitempty" xmlrpc:"outBoundTotalBytes"`
}

// SoftLayer_Container_Bandwidth_Projection models projected bandwidth use over a time range.
type Container_Bandwidth_Projection struct {
	Entity

	// Bandwidth limit for this hardware.
	AllowedUsage *string `json:"allowedUsage,omitempty" xmlrpc:"allowedUsage"`

	// Estimated bandwidth usage so far this billing cycle.
	EstimatedUsage *string `json:"estimatedUsage,omitempty" xmlrpc:"estimatedUsage"`

	// Hardware ID of server to monitor.
	HardwareId *int `json:"hardwareId,omitempty" xmlrpc:"hardwareId"`

	// Projected usage for this hardware based on previous usage this billing cycle.
	ProjectedUsage *string `json:"projectedUsage,omitempty" xmlrpc:"projectedUsage"`

	// the text name of the server being monitored.
	ServerName *string `json:"serverName,omitempty" xmlrpc:"serverName"`

	// The minimum date included in this list.
	StartDate *Time `json:"startDate,omitempty" xmlrpc:"startDate"`
}

// no documentation yet
type Container_Billing_Currency_Format struct {
	Entity

	// no documentation yet
	Currency *string `json:"currency,omitempty" xmlrpc:"currency"`

	// no documentation yet
	Display *int `json:"display,omitempty" xmlrpc:"display"`

	// no documentation yet
	Format *string `json:"format,omitempty" xmlrpc:"format"`

	// no documentation yet
	Locale *string `json:"locale,omitempty" xmlrpc:"locale"`

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// no documentation yet
	Position *int `json:"position,omitempty" xmlrpc:"position"`

	// no documentation yet
	Precision *int `json:"precision,omitempty" xmlrpc:"precision"`

	// no documentation yet
	Script *string `json:"script,omitempty" xmlrpc:"script"`

	// no documentation yet
	Service *string `json:"service,omitempty" xmlrpc:"service"`

	// no documentation yet
	Symbol *string `json:"symbol,omitempty" xmlrpc:"symbol"`

	// no documentation yet
	Tag *string `json:"tag,omitempty" xmlrpc:"tag"`

	// no documentation yet
	Value *Float64 `json:"value,omitempty" xmlrpc:"value"`
}

// no documentation yet
type Container_Billing_Info_Ach struct {
	Entity

	// no documentation yet
	AccountNumber *string `json:"accountNumber,omitempty" xmlrpc:"accountNumber"`

	// no documentation yet
	AccountType *string `json:"accountType,omitempty" xmlrpc:"accountType"`

	// no documentation yet
	BankTransitNumber *string `json:"bankTransitNumber,omitempty" xmlrpc:"bankTransitNumber"`

	// no documentation yet
	City *string `json:"city,omitempty" xmlrpc:"city"`

	// no documentation yet
	Country *string `json:"country,omitempty" xmlrpc:"country"`

	// no documentation yet
	FederalTaxId *string `json:"federalTaxId,omitempty" xmlrpc:"federalTaxId"`

	// no documentation yet
	FirstName *string `json:"firstName,omitempty" xmlrpc:"firstName"`

	// no documentation yet
	LastName *string `json:"lastName,omitempty" xmlrpc:"lastName"`

	// no documentation yet
	PhoneNumber *string `json:"phoneNumber,omitempty" xmlrpc:"phoneNumber"`

	// no documentation yet
	PostalCode *string `json:"postalCode,omitempty" xmlrpc:"postalCode"`

	// no documentation yet
	State *string `json:"state,omitempty" xmlrpc:"state"`

	// no documentation yet
	Street1 *string `json:"street1,omitempty" xmlrpc:"street1"`

	// no documentation yet
	Street2 *string `json:"street2,omitempty" xmlrpc:"street2"`
}

// This container is used to provide all the options for [[SoftLayer_Billing_Invoice/emailInvoices|emailInvoices]] in order to have the necessary invoices generated and links sent to the user's email.
type Container_Billing_Invoice_Email struct {
	Entity

	// Excel Invoices to email
	ExcelInvoiceIds []int `json:"excelInvoiceIds,omitempty" xmlrpc:"excelInvoiceIds"`

	// PDF Invoice Details to email
	PdfDetailedInvoiceIds []int `json:"pdfDetailedInvoiceIds,omitempty" xmlrpc:"pdfDetailedInvoiceIds"`

	// PDF Invoices to email
	PdfInvoiceIds []int `json:"pdfInvoiceIds,omitempty" xmlrpc:"pdfInvoiceIds"`

	// The type of Invoices to be emailed [current|next]. If next is selected, the account id will be used.
	Type *string `json:"type,omitempty" xmlrpc:"type"`
}

// SoftLayer_Container_Billing_Order_Status models an order status.
type Container_Billing_Order_Status struct {
	Entity

	// The description of the status.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// The keyname of the status.
	Status *string `json:"status,omitempty" xmlrpc:"status"`
}

// Contains user information used to request a manual Catalyst enrollment.
type Container_Catalyst_ManualEnrollmentRequest struct {
	Entity

	// Applicant's email address
	CustomerEmail *string `json:"customerEmail,omitempty" xmlrpc:"customerEmail"`

	// Applicant's first and last name
	CustomerName *string `json:"customerName,omitempty" xmlrpc:"customerName"`

	// Name of applicant's startup company
	StartupName *string `json:"startupName,omitempty" xmlrpc:"startupName"`

	// Flag indicating whether (true) or not (false) and applicant is
	VentureAffiliationFlag *bool `json:"ventureAffiliationFlag,omitempty" xmlrpc:"ventureAffiliationFlag"`

	// Name of the venture capital fund, if any, applicant is affiliated with
	VentureFundName *string `json:"ventureFundName,omitempty" xmlrpc:"ventureFundName"`
}

// This container is used to hold country locale information.
type Container_Collection_Locale_CountryCode struct {
	Entity

	// no documentation yet
	LongName *string `json:"longName,omitempty" xmlrpc:"longName"`

	// no documentation yet
	ShortName *string `json:"shortName,omitempty" xmlrpc:"shortName"`

	// no documentation yet
	StateCodes []Container_Collection_Locale_StateCode `json:"stateCodes,omitempty" xmlrpc:"stateCodes"`
}

// This container is used to hold information regarding a state or province.
type Container_Collection_Locale_StateCode struct {
	Entity

	// no documentation yet
	LongName *string `json:"longName,omitempty" xmlrpc:"longName"`

	// no documentation yet
	ShortName *string `json:"shortName,omitempty" xmlrpc:"shortName"`
}

// no documentation yet
type Container_Disk_Image_Capture_Template struct {
	Entity

	// no documentation yet
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// no documentation yet
	Summary *string `json:"summary,omitempty" xmlrpc:"summary"`

	// no documentation yet
	Volumes []Container_Disk_Image_Capture_Template_Volume `json:"volumes,omitempty" xmlrpc:"volumes"`
}

// no documentation yet
type Container_Disk_Image_Capture_Template_Volume struct {
	Entity

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// no documentation yet
	Partitions []Container_Disk_Image_Capture_Template_Volume_Partition `json:"partitions,omitempty" xmlrpc:"partitions"`
}

// no documentation yet
type Container_Disk_Image_Capture_Template_Volume_Partition struct {
	Entity

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// Contact information container for domain registration
type Container_Dns_Domain_Registration_Contact struct {
	Entity

	// The street address of the contact.
	Address1 *string `json:"address1,omitempty" xmlrpc:"address1"`

	// The second line in the address of the contact.
	Address2 *string `json:"address2,omitempty" xmlrpc:"address2"`

	// The third line in the address of the contact.
	Address3 *string `json:"address3,omitempty" xmlrpc:"address3"`

	// The city of the contact.
	City *string `json:"city,omitempty" xmlrpc:"city"`

	// The 2-character Country code. (i.e. US)
	Country *string `json:"country,omitempty" xmlrpc:"country"`

	// The email address of the contact.
	Email *string `json:"email,omitempty" xmlrpc:"email"`

	// The fax number of the contact.
	Fax *string `json:"fax,omitempty" xmlrpc:"fax"`

	// The first name of the contact.
	FirstName *string `json:"firstName,omitempty" xmlrpc:"firstName"`

	// The last name of the contact.
	LastName *string `json:"lastName,omitempty" xmlrpc:"lastName"`

	// The organization name of the contact.
	OrganizationName *string `json:"organizationName,omitempty" xmlrpc:"organizationName"`

	// The phone number of the contact.
	Phone *string `json:"phone,omitempty" xmlrpc:"phone"`

	// The postal code of the contact.
	PostalCode *string `json:"postalCode,omitempty" xmlrpc:"postalCode"`

	// The state of the contact.
	State *string `json:"state,omitempty" xmlrpc:"state"`

	// The type of contact. The following are the valid types of contacts:
	// * admin
	// * owner
	// * billing
	// * tech
	Type *string `json:"type,omitempty" xmlrpc:"type"`
}

// This container data type contains extended attributes information for a domain of country code TLD.
type Container_Dns_Domain_Registration_ExtendedAttribute struct {
	Entity

	// Indicates if this is a child of another extended attribute.
	ChildFlag *bool `json:"childFlag,omitempty" xmlrpc:"childFlag"`

	// The description of an extended attribute.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// The name of an extended attribute.
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// The collection of options for an extended attribute.
	Options []Container_Dns_Domain_Registration_ExtendedAttribute_Option `json:"options,omitempty" xmlrpc:"options"`

	// Indicates if extended attribute is required.
	RequiredFlag *int `json:"requiredFlag,omitempty" xmlrpc:"requiredFlag"`

	// User defined indicates that the value is required from outside sources.
	UserDefinedFlag *bool `json:"userDefinedFlag,omitempty" xmlrpc:"userDefinedFlag"`
}

// This is the data type that may need to be populated to complete registraton for domains that are country code TLD's.
type Container_Dns_Domain_Registration_ExtendedAttribute_Configuration struct {
	Entity

	// The extended attribute name.
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// The extended attribute option value.
	Value *string `json:"value,omitempty" xmlrpc:"value"`
}

// This container data type contains extended attribute options information for a domain of country code TLD.
type Container_Dns_Domain_Registration_ExtendedAttribute_Option struct {
	Entity

	// The description of an option.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// Extended Attribute that is required for an option.
	RequireExtendedAttributes []Container_Dns_Domain_Registration_ExtendedAttribute_Option_Require `json:"requireExtendedAttributes,omitempty" xmlrpc:"requireExtendedAttributes"`

	// The title of an option.
	Title *string `json:"title,omitempty" xmlrpc:"title"`

	// The value of an option.
	Value *string `json:"value,omitempty" xmlrpc:"value"`
}

// This container data type contains the extended attribute name that is required by an extended attribute option.
type Container_Dns_Domain_Registration_ExtendedAttribute_Option_Require struct {
	Entity

	// The name of an extended attribute that is required by an extended attribute option.
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// Information container for domain registration
type Container_Dns_Domain_Registration_Information struct {
	Entity

	// The information of the registered domain.
	Contacts []Container_Dns_Domain_Registration_Contact `json:"contacts,omitempty" xmlrpc:"contacts"`

	// The date that a domain is set to expire.
	ExpireDate *Time `json:"expireDate,omitempty" xmlrpc:"expireDate"`

	// The list of nameservers for the domain.
	Nameservers []Container_Dns_Domain_Registration_Nameserver `json:"nameservers,omitempty" xmlrpc:"nameservers"`

	// no documentation yet
	RegistryCreateDate *Time `json:"registryCreateDate,omitempty" xmlrpc:"registryCreateDate"`

	// no documentation yet
	RegistryExpireDate *Time `json:"registryExpireDate,omitempty" xmlrpc:"registryExpireDate"`

	// no documentation yet
	RegistryUpdateDate *Time `json:"registryUpdateDate,omitempty" xmlrpc:"registryUpdateDate"`
}

// no documentation yet
type Container_Dns_Domain_Registration_List struct {
	Entity

	// The domain name.
	DomainName *string `json:"domainName,omitempty" xmlrpc:"domainName"`

	// Three-character language tag for the IDN domain that you're trying to register. This is only required for IDN domains.
	EncodingType *string `json:"encodingType,omitempty" xmlrpc:"encodingType"`

	// Data required by the Registry for some country code top level domains (i.e. example.us).
	//
	// In order to determine if a domain requires extended attributes, use [[SoftLayer_Dns_Domain_Registration::getExtendedAttributes|domain registration]] service.
	ExtendedAttributeConfiguration []Container_Dns_Domain_Registration_ExtendedAttribute_Configuration `json:"extendedAttributeConfiguration,omitempty" xmlrpc:"extendedAttributeConfiguration"`

	// The length of the registration period in years. Valid values are 1 – 10.
	RegistrationPeriod *int `json:"registrationPeriod,omitempty" xmlrpc:"registrationPeriod"`
}

// Lookup domain container for domain registration
type Container_Dns_Domain_Registration_Lookup struct {
	Entity

	// The list of available and taken domain names.
	Items []Container_Dns_Domain_Registration_Lookup_Items `json:"items,omitempty" xmlrpc:"items"`
}

// Lookup items container for domain registration
type Container_Dns_Domain_Registration_Lookup_Items struct {
	Entity

	// The domain name.
	DomainName *string `json:"domainName,omitempty" xmlrpc:"domainName"`

	// The status of the domain name if available and can be registered.
	Status *string `json:"status,omitempty" xmlrpc:"status"`
}

// Nameserver container for domain registration
type Container_Dns_Domain_Registration_Nameserver struct {
	Entity

	// The list of fully qualified names of the nameserver.
	Nameservers []Container_Dns_Domain_Registration_Nameserver_List `json:"nameservers,omitempty" xmlrpc:"nameservers"`
}

// Nameservers list container for domain registration
type Container_Dns_Domain_Registration_Nameserver_List struct {
	Entity

	// The IPv4 address of the nameserver.
	Ipv4Address *string `json:"ipv4Address,omitempty" xmlrpc:"ipv4Address"`

	// The IPv6 address of the nameserver.
	Ipv6Address *string `json:"ipv6Address,omitempty" xmlrpc:"ipv6Address"`

	// The fully qualified name of the nameserver
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// The sort order of the nameserver
	SortOrder *int `json:"sortOrder,omitempty" xmlrpc:"sortOrder"`
}

// no documentation yet
type Container_Dns_Domain_Registration_Registrant_Verification_StatusDetail struct {
	Entity

	// The current status of the verification.
	Status *Dns_Domain_Registration_Registrant_Verification_Status `json:"status,omitempty" xmlrpc:"status"`

	// The adate when the domain will be suspended.
	VerificationDeadlineDate *Time `json:"verificationDeadlineDate,omitempty" xmlrpc:"verificationDeadlineDate"`
}

// Transfer Information container for domain registration
type Container_Dns_Domain_Registration_Transfer_Information struct {
	Entity

	// The reason why a domain is not transferable.
	Reason *string `json:"reason,omitempty" xmlrpc:"reason"`

	// The registrant email.
	RegistrantEmail *string `json:"registrantEmail,omitempty" xmlrpc:"registrantEmail"`

	// The status of the latest transfer on the domain.
	Status *string `json:"status,omitempty" xmlrpc:"status"`

	// The date and time of the most recent update to the state of the transfer.
	TimeStamp *Time `json:"timeStamp,omitempty" xmlrpc:"timeStamp"`

	// Indicates if the domain can be transferred.
	Transferrable *int `json:"transferrable,omitempty" xmlrpc:"transferrable"`
}

// The SoftLayer_Container_Exception data type represents a SoftLayer_Exception.
type Container_Exception struct {
	Entity

	// The SoftLayer_Exception class that the error is.
	ExceptionClass *string `json:"exceptionClass,omitempty" xmlrpc:"exceptionClass"`

	// The exception message.
	ExceptionMessage *string `json:"exceptionMessage,omitempty" xmlrpc:"exceptionMessage"`
}

// no documentation yet
type Container_Graph struct {
	Entity

	// base units associated with the graph.
	BaseUnit *string `json:"baseUnit,omitempty" xmlrpc:"baseUnit"`

	// Graph range end datetime.
	EndDatetime *string `json:"endDatetime,omitempty" xmlrpc:"endDatetime"`

	// The height of the graph image.
	Height *int `json:"height,omitempty" xmlrpc:"height"`

	// The graph image.
	Image *[]byte `json:"image,omitempty" xmlrpc:"image"`

	// The graph interval in seconds.
	Interval *int `json:"interval,omitempty" xmlrpc:"interval"`

	// Metric types associated with the graph.
	Metrics []Container_Metric_Data_Type `json:"metrics,omitempty" xmlrpc:"metrics"`

	// Indicator to control whether the graph data is normalized.
	NormalizeFlag *[]byte `json:"normalizeFlag,omitempty" xmlrpc:"normalizeFlag"`

	// The options used to control the graph appearance.
	Options []Container_Graph_Option `json:"options,omitempty" xmlrpc:"options"`

	// A collection of graph plots.
	Plots []Container_Graph_Plot `json:"plots,omitempty" xmlrpc:"plots"`

	// option to not return the image.
	ReturnImage *bool `json:"returnImage,omitempty" xmlrpc:"returnImage"`

	// Graph range start datetime.
	StartDatetime *string `json:"startDatetime,omitempty" xmlrpc:"startDatetime"`

	// The name of the template to use; may be null.
	Template *string `json:"template,omitempty" xmlrpc:"template"`

	// The title of the graph image.
	Title *string `json:"title,omitempty" xmlrpc:"title"`

	// The width of the graph image.
	Width *int `json:"width,omitempty" xmlrpc:"width"`
}

// no documentation yet
type Container_Graph_Option struct {
	Entity

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// no documentation yet
	Value *string `json:"value,omitempty" xmlrpc:"value"`
}

// no documentation yet
type Container_Graph_Plot struct {
	Entity

	// no documentation yet
	Data []Container_Graph_Plot_Coordinate `json:"data,omitempty" xmlrpc:"data"`

	// no documentation yet
	Metric *Container_Metric_Data_Type `json:"metric,omitempty" xmlrpc:"metric"`

	// no documentation yet
	Unit *string `json:"unit,omitempty" xmlrpc:"unit"`
}

// no documentation yet
type Container_Graph_Plot_Coordinate struct {
	Entity

	// no documentation yet
	XValue *Float64 `json:"xValue,omitempty" xmlrpc:"xValue"`

	// no documentation yet
	YValue *Float64 `json:"yValue,omitempty" xmlrpc:"yValue"`

	// no documentation yet
	ZValue *Float64 `json:"zValue,omitempty" xmlrpc:"zValue"`
}

// The hardware configuration container is used to provide configuration options for servers.
//
// Each configuration option will include both an <code>itemPrice</code> and a <code>template</code>.
//
// The <code>itemPrice</code> value will provide hourly and monthly costs (if either are applicable), and a description of the option.
//
// The <code>template</code> will provide a fragment of the request with the properties and values that must be sent when creating a server with the option.
//
// The [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] method returns this data structure.
//
// <style type="text/css">#properties .views-field-body p { margin-top: 1.5em; };</style>
type Container_Hardware_Configuration struct {
	Entity

	//
	// <div style="width: 200%">
	// Available datacenter options.
	//
	//
	// The <code>datacenter.name</code> value in the template represents which datacenter the server will be provisioned in.
	// </div>
	Datacenters []Container_Hardware_Configuration_Option `json:"datacenters,omitempty" xmlrpc:"datacenters"`

	//
	// <div style="width: 200%">
	// Available fixed configuration preset options.
	//
	//
	// The <code>fixedConfigurationPreset.keyName</code> value in the template is an identifier for a particular fixed configuration. When provided exactly as shown in the template, that fixed configuration will be used.
	//
	//
	// When providing a <code>fixedConfigurationPreset.keyName</code> while ordering a server the <code>processors</code> and <code>hardDrives</code> configuration options cannot be used.
	// </div>
	FixedConfigurationPresets []Container_Hardware_Configuration_Option `json:"fixedConfigurationPresets,omitempty" xmlrpc:"fixedConfigurationPresets"`

	//
	// <div style="width: 200%">
	// Available hard drive options.
	//
	//
	// A server will have at least one hard drive.
	//
	//
	// The <code>hardDrives.capacity</code> value in the template represents the size, in gigabytes, of the disk.
	// </div>
	HardDrives []Container_Hardware_Configuration_Option `json:"hardDrives,omitempty" xmlrpc:"hardDrives"`

	//
	// <div style="width: 200%">
	// Available network component options.
	//
	//
	// The <code>networkComponent.maxSpeed</code> value in the template represents the link speed, in megabits per second, of the network connections for a server.
	// </div>
	NetworkComponents []Container_Hardware_Configuration_Option `json:"networkComponents,omitempty" xmlrpc:"networkComponents"`

	//
	// <div style="width: 200%">
	// Available operating system options.
	//
	//
	// The <code>operatingSystemReferenceCode</code> value in the template is an identifier for a particular operating system. When provided exactly as shown in the template, that operating system will be used.
	//
	//
	// A reference code is structured as three tokens separated by underscores. The first token represents the product, the second is the version of the product, and the third is whether the OS is 32 or 64bit.
	//
	//
	// When providing an <code>operatingSystemReferenceCode</code> while ordering a server the only token required to match exactly is the product. The version token may be given as 'LATEST', else it will require an exact match as well. When the bits token is not provided, 64 bits will be assumed.
	//
	//
	// Providing the value of 'LATEST' for a version will select the latest release of that product for the operating system. As this may change over time, you should be sure that the release version is irrelevant for your applications.
	//
	//
	// For Windows based operating systems the version will represent both the release version (2008, 2012, etc) and the edition (Standard, Enterprise, etc). For all other operating systems the version will represent the major version (Centos 6, Ubuntu 12, etc) of that operating system, minor versions are represented in few reference codes where they are significant.
	// </div>
	OperatingSystems []Container_Hardware_Configuration_Option `json:"operatingSystems,omitempty" xmlrpc:"operatingSystems"`

	//
	// <div style="width: 200%">
	// Available processor options.
	//
	//
	// The <code>processorCoreAmount</code> value in the template represents the number of cores allocated to the server.
	// The <code>memoryCapacity</code> value in the template represents the amount of memory, in gigabytes, allocated to the server.
	// </div>
	Processors []Container_Hardware_Configuration_Option `json:"processors,omitempty" xmlrpc:"processors"`
}

// An option found within a [[SoftLayer_Container_Hardware_Configuration (type)]] structure.
type Container_Hardware_Configuration_Option struct {
	Entity

	//
	// Provides hourly and monthly costs (if either are applicable), and a description of the option.
	ItemPrice *Product_Item_Price `json:"itemPrice,omitempty" xmlrpc:"itemPrice"`

	//
	// Provides a description of a fixed configuration preset with monthly and hourly costs.
	Preset *Product_Package_Preset `json:"preset,omitempty" xmlrpc:"preset"`

	//
	// Provides a fragment of the request with the properties and values that must be sent when creating a server with the option.
	Template *Hardware `json:"template,omitempty" xmlrpc:"template"`
}

// no documentation yet
type Container_Hardware_MassUpdate struct {
	Entity

	// The hardwares updated by the mass update tool
	HardwareId *int `json:"hardwareId,omitempty" xmlrpc:"hardwareId"`

	// Errors encountered while mass updating hardwares
	Message *string `json:"message,omitempty" xmlrpc:"message"`

	// The hardwares that failed to update
	SuccessFlag *string `json:"successFlag,omitempty" xmlrpc:"successFlag"`
}

// The SoftLayer_Container_Hardware_Server_Configuration data type contains information relating to a server's item price information, and hard drive partition information.
type Container_Hardware_Server_Configuration struct {
	Entity

	// A flag indicating that the server will be moved into the spare pool after an Operating system reload.
	AddToSparePoolAfterOsReload *int `json:"addToSparePoolAfterOsReload,omitempty" xmlrpc:"addToSparePoolAfterOsReload"`

	// The customer provision uri will be used to download and execute a customer defined script on the host at the end of provisioning.
	CustomProvisionScriptUri *string `json:"customProvisionScriptUri,omitempty" xmlrpc:"customProvisionScriptUri"`

	// A flag indicating that the primary drive will be converted to a portable storage volume during an Operating System reload.
	DriveRetentionFlag *bool `json:"driveRetentionFlag,omitempty" xmlrpc:"driveRetentionFlag"`

	// A flag indicating that all data will be erased from drives during an Operating System reload.
	EraseHardDrives *int `json:"eraseHardDrives,omitempty" xmlrpc:"eraseHardDrives"`

	// The hard drive partitions that a server can be partitioned with.
	HardDrives []Hardware_Component `json:"hardDrives,omitempty" xmlrpc:"hardDrives"`

	// An Image Template ID [[SoftLayer_Virtual_Guest_Block_Device_Template_Group]] that will be deployed to the host.  If provided no item prices are required.
	ImageTemplateId *int `json:"imageTemplateId,omitempty" xmlrpc:"imageTemplateId"`

	// The item prices that a server can be configured with.
	ItemPrices []Product_Item_Price `json:"itemPrices,omitempty" xmlrpc:"itemPrices"`

	// A flag indicating that the remote management cards password will be reset.
	ResetIpmiPassword *int `json:"resetIpmiPassword,omitempty" xmlrpc:"resetIpmiPassword"`

	// IDs to SoftLayer_Security_Ssh_Key objects on the current account which will be added to the server for authentication. SSH Keys will not be added to servers with Microsoft Windows.
	SshKeyIds []int `json:"sshKeyIds,omitempty" xmlrpc:"sshKeyIds"`

	// A flag indicating that the the BIOS will not be updated when installing the operating system.
	UpgradeBios *int `json:"upgradeBios,omitempty" xmlrpc:"upgradeBios"`

	// A flag indicating that the firmware on all hard drives will be updated when installing the operating system.
	UpgradeHardDriveFirmware *int `json:"upgradeHardDriveFirmware,omitempty" xmlrpc:"upgradeHardDriveFirmware"`
}

// The SoftLayer_Container_Hardware_Server_Details data type contains information relating to a server's component information, network information, and software information.
type Container_Hardware_Server_Details struct {
	Entity

	// The components that belong to a piece of hardware.
	Components []Hardware_Component `json:"components,omitempty" xmlrpc:"components"`

	// The network components that belong to a piece of hardware.
	NetworkComponents []Network_Component `json:"networkComponents,omitempty" xmlrpc:"networkComponents"`

	// The software that belong to a piece of hardware.
	Software []Software_Component `json:"software,omitempty" xmlrpc:"software"`
}

// SoftLayer_Container_KnowledgeLayer_QuestionAnswer models a single question and answer pair from SoftLayer's KnowledgeLayer knowledge base. SoftLayer's backend network interfaces with the KnowledgeLayer to recommend helpful articles when support tickets are created.
type Container_KnowledgeLayer_QuestionAnswer struct {
	Entity

	// The answer to a question asked on the SoftLayer KnowledgeLayer.
	Answer *string `json:"answer,omitempty" xmlrpc:"answer"`

	// The link to a question asked on the SoftLayer KnowledgeLayer.
	Link *string `json:"link,omitempty" xmlrpc:"link"`

	// A question asked on the SoftLayer KnowledgeLayer.
	Question *string `json:"question,omitempty" xmlrpc:"question"`
}

// no documentation yet
type Container_Message struct {
	Entity

	// no documentation yet
	Message *string `json:"message,omitempty" xmlrpc:"message"`

	// no documentation yet
	Type *string `json:"type,omitempty" xmlrpc:"type"`
}

// no documentation yet
type Container_Metric_Data_Type struct {
	Entity

	// no documentation yet
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// no documentation yet
	SummaryType *string `json:"summaryType,omitempty" xmlrpc:"summaryType"`

	// no documentation yet
	Unit *string `json:"unit,omitempty" xmlrpc:"unit"`
}

// SoftLayer_Container_Metric_Tracking_Object_Details This container is a parent class for detailing diverse metrics.
type Container_Metric_Tracking_Object_Details struct {
	Entity

	// The name that best describes the metric being collected.
	MetricName *string `json:"metricName,omitempty" xmlrpc:"metricName"`
}

// SoftLayer_Container_Metric_Tracking_Object_Summary This container is a parent class for summarizing diverse metrics.
type Container_Metric_Tracking_Object_Summary struct {
	Entity

	// The name that best describes the metric being collected.
	MetricName *string `json:"metricName,omitempty" xmlrpc:"metricName"`
}

// SoftLayer_Container_Metric_Tracking_Object_Virtual_Host_Details This container details a virtual host's metric data.
type Container_Metric_Tracking_Object_Virtual_Host_Details struct {
	Container_Metric_Tracking_Object_Details

	// The day this metric was collected.
	Day *Time `json:"day,omitempty" xmlrpc:"day"`

	// The maximum number of guests hosted by this platform for the given day.
	MaxInstances *int `json:"maxInstances,omitempty" xmlrpc:"maxInstances"`

	// The maximum amount of memory utilized by this platform for the given day.
	MaxMemoryUsage *int `json:"maxMemoryUsage,omitempty" xmlrpc:"maxMemoryUsage"`

	// The mean number of guests hosted by this platform for the given day.
	MeanInstances *Float64 `json:"meanInstances,omitempty" xmlrpc:"meanInstances"`

	// The mean amount of memory utilized by this platform for the given day.
	MeanMemoryUsage *Float64 `json:"meanMemoryUsage,omitempty" xmlrpc:"meanMemoryUsage"`

	// The minimum number of guests hosted by this platform for the given day.
	MinInstances *int `json:"minInstances,omitempty" xmlrpc:"minInstances"`

	// The minimum amount of memory utilized by this platform for the given day.
	MinMemoryUsage *int `json:"minMemoryUsage,omitempty" xmlrpc:"minMemoryUsage"`
}

// SoftLayer_Container_Metric_Tracking_Object_Virtual_Host_Summary This container summarizes a virtual host's metric data.
type Container_Metric_Tracking_Object_Virtual_Host_Summary struct {
	Container_Metric_Tracking_Object_Summary

	// The average amount of memory usage thus far in this billing cycle.
	AvgMemoryUsageInBillingCycle *int `json:"avgMemoryUsageInBillingCycle,omitempty" xmlrpc:"avgMemoryUsageInBillingCycle"`

	// Current bill cycle end date.
	CurrentBillCycleEnd *Time `json:"currentBillCycleEnd,omitempty" xmlrpc:"currentBillCycleEnd"`

	// Current bill cycle start date.
	CurrentBillCycleStart *Time `json:"currentBillCycleStart,omitempty" xmlrpc:"currentBillCycleStart"`

	// The last count of instances this platform was hosting.
	LastInstanceCount *int `json:"lastInstanceCount,omitempty" xmlrpc:"lastInstanceCount"`

	// The last amount of memory this platform was using.
	LastMemoryUsageAmount *int `json:"lastMemoryUsageAmount,omitempty" xmlrpc:"lastMemoryUsageAmount"`

	// The last time this virtual host was polled for metrics.
	LastPollTime *Time `json:"lastPollTime,omitempty" xmlrpc:"lastPollTime"`

	// The max number of instances hosted thus far in this billing cycle.
	MaxInstanceInBillingCycle *int `json:"maxInstanceInBillingCycle,omitempty" xmlrpc:"maxInstanceInBillingCycle"`

	// Previous bill cycle end date.
	PreviousBillCycleEnd *Time `json:"previousBillCycleEnd,omitempty" xmlrpc:"previousBillCycleEnd"`

	// Previous bill cycle start date.
	PreviousBillCycleStart *Time `json:"previousBillCycleStart,omitempty" xmlrpc:"previousBillCycleStart"`

	// This virtual hosting platform name.
	VirtualPlatformName *string `json:"virtualPlatformName,omitempty" xmlrpc:"virtualPlatformName"`
}

// The SoftLayer_Container_Monitoring_Alarm_History data type contains information relating to SoftLayer monitoring alarm history.
type Container_Monitoring_Alarm_History struct {
	Entity

	// Account ID that this alarm belongs to
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// ID of the monitoring agent that triggered this alarm
	AgentId *int `json:"agentId,omitempty" xmlrpc:"agentId"`

	// Alarm ID
	AlarmId *string `json:"alarmId,omitempty" xmlrpc:"alarmId"`

	// Time that an alarm was closed.
	ClosedDate *Time `json:"closedDate,omitempty" xmlrpc:"closedDate"`

	// Time that an alarm was triggered
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// Alarm message
	Message *string `json:"message,omitempty" xmlrpc:"message"`

	// Robot ID
	RobotId *int `json:"robotId,omitempty" xmlrpc:"robotId"`

	// Severity of an alarm
	Severity *string `json:"severity,omitempty" xmlrpc:"severity"`
}

// SoftLayer_Container_Monitoring_Graph_Outputs models a single outbound object for a graph of given data sets.
type Container_Monitoring_Graph_Outputs struct {
	Entity

	// The maximum date included in this graph.
	EndDate *Time `json:"endDate,omitempty" xmlrpc:"endDate"`

	// Error message encountered during graphing
	GraphError *string `json:"graphError,omitempty" xmlrpc:"graphError"`

	// The raw PNG binary data to be displayed once the graph is drawn.
	GraphImage *[]byte `json:"graphImage,omitempty" xmlrpc:"graphImage"`

	// The minimum date included in this graph.
	StartDate *Time `json:"startDate,omitempty" xmlrpc:"startDate"`
}

// This object holds authentication data to a server.
type Container_Network_Authentication_Data struct {
	Entity

	// The name of a host
	Host *string `json:"host,omitempty" xmlrpc:"host"`

	// The authentication password
	Password *string `json:"password,omitempty" xmlrpc:"password"`

	// The port number
	Port *int `json:"port,omitempty" xmlrpc:"port"`

	// The type of network protocol. This can be ftp, ssh and so on.
	Type *string `json:"type,omitempty" xmlrpc:"type"`

	// The authentication username
	Username *string `json:"username,omitempty" xmlrpc:"username"`
}

// SoftLayer_Container_Network_Bandwidth_Data_Summary models an interface's overall bandwidth usage during it's current billing cycle.
type Container_Network_Bandwidth_Data_Summary struct {
	Entity

	// The amount of bandwidth a server has allocated to it in it's current billing period.
	AllowedUsage *Float64 `json:"allowedUsage,omitempty" xmlrpc:"allowedUsage"`

	// The amount of bandwidth that a server has used within it's current billing period.
	EstimatedUsage *Float64 `json:"estimatedUsage,omitempty" xmlrpc:"estimatedUsage"`

	// The amount of bandwidth a server is projected to use within its billing period, based on it's current usage.
	ProjectedUsage *Float64 `json:"projectedUsage,omitempty" xmlrpc:"projectedUsage"`

	// The unit of measurement used in a bandwidth data summary.
	UsageUnits *string `json:"usageUnits,omitempty" xmlrpc:"usageUnits"`
}

// SoftLayer_Container_Network_Bandwidth_Version1_Usage models an hourly bandwidth record.
type Container_Network_Bandwidth_Version1_Usage struct {
	Entity

	// The amount of incoming bandwidth that a server has used within the hour of the recordedDate.
	IncomingAmount *Float64 `json:"incomingAmount,omitempty" xmlrpc:"incomingAmount"`

	// The amount of outgoing bandwidth that a server has used within the hour of the recordedDate.
	OutgoingAmount *Float64 `json:"outgoingAmount,omitempty" xmlrpc:"outgoingAmount"`

	// The date and time that the bandwidth was used by a piece of hardware
	RecordedDate *Time `json:"recordedDate,omitempty" xmlrpc:"recordedDate"`
}

// SoftLayer_Container_Network_ContentDelivery_Authentication_Directory represents a token authentication directory on your CDN FTP or on your origin server.
type Container_Network_ContentDelivery_Authentication_Directory struct {
	Entity

	// The date that a token authentication directory was created.
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// The name of a directory or a file within a directory listing.
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// The type of platform that a token authentication directory is defined for. Possible types are HTTP Large, HTTP Small, Flash and Windows Media
	Type *string `json:"type,omitempty" xmlrpc:"type"`
}

// This container is used for CDN content authentication service.
type Container_Network_ContentDelivery_Authentication_Parameter struct {
	Entity

	// A CDN account name
	CdnAccountName *string `json:"cdnAccountName,omitempty" xmlrpc:"cdnAccountName"`

	// A client IP address
	ClientIp *string `json:"clientIp,omitempty" xmlrpc:"clientIp"`

	// A client referrer information
	Referrer *string `json:"referrer,omitempty" xmlrpc:"referrer"`

	// A source URL
	SourceUrl *string `json:"sourceUrl,omitempty" xmlrpc:"sourceUrl"`

	// An authentication token string
	Token *string `json:"token,omitempty" xmlrpc:"token"`
}

// CDN supports the content authentication service. With the content authentication service, customers can control access to their contents. There are several scenarios where this authentication capability could be useful. Websites can prevent other rogue websites from linking to their videos. Content owners can prevent users from passing around http links, thus forcing them to login to view contents. It is also possible to authenticate via the client IP address. Referrer information is also checked if provided by a client's browser. servers will invoke a web service method to validate a content authentication token.
//
// CDN uses the default authentication web service provided by SoftLayer to validate a token. A customer can use their own implementation of the token authentication web service by using [[SoftLayer_Network_ContentDelivery_Account::setAuthenticationServiceEndpoint|setAuthenticationServiceEndpoint]] method.
//
// This container class holds the token validation web service endpoint information. CDN supports 3 different protocols: HTTP, RTMP (streaming Flash), and MMS (streaming Windows Media)
type Container_Network_ContentDelivery_Authentication_ServiceEndpoint struct {
	Entity

	// The authentication web service endpoint that CDN servers will use to validate a token
	Endpoint *string `json:"endpoint,omitempty" xmlrpc:"endpoint"`

	// The protocol that the WSDL will be used for.  This can be HTTP, WINDOWSMEDIA, or FLASH
	Protocol *string `json:"protocol,omitempty" xmlrpc:"protocol"`
}

// SoftLayer_Container_Network_ContentDelivery_Bandwidth_PointsOfPresence_Summary models an individual CDN point of presence's bandwidth usage for a CDN account within a given date range. CDN POPs are located throughout the world, so individual POP usage may be beneficial in determining who is downloading your CDN hosted content.
type Container_Network_ContentDelivery_Bandwidth_PointsOfPresence_Summary struct {
	Entity

	// The amount of bandwidth used by a CDN POP.
	Bandwidth *uint `json:"bandwidth,omitempty" xmlrpc:"bandwidth"`

	// The ending date of a CDN POP bandwidth summary.
	EndDateTime *Time `json:"endDateTime,omitempty" xmlrpc:"endDateTime"`

	// A CDN POP's name. This is typically the city the POP resides in.
	PopName *string `json:"popName,omitempty" xmlrpc:"popName"`

	// The starting date of a CDN POP bandwidth summary.
	StartDateTime *Time `json:"startDateTime,omitempty" xmlrpc:"startDateTime"`

	// The unit of measurement used in a CDN POP bandwidth summary.
	UsageUnits *string `json:"usageUnits,omitempty" xmlrpc:"usageUnits"`

	// The view count
	ViewCount *uint `json:"viewCount,omitempty" xmlrpc:"viewCount"`
}

// SoftLayer_Container_Network_ContentDelivery_Bandwidth_Summary models a CDN account's overall bandwidth usage and overages within a given date range.
type Container_Network_ContentDelivery_Bandwidth_Summary struct {
	Entity

	// The CDN account id
	CdnAccountId *int `json:"cdnAccountId,omitempty" xmlrpc:"cdnAccountId"`

	// The ending date of a CDN bandwidth summary.
	EndDateTime *Time `json:"endDateTime,omitempty" xmlrpc:"endDateTime"`

	// The name of a file that is requested by visitors
	FileName *string `json:"fileName,omitempty" xmlrpc:"fileName"`

	// The media type
	MediaType *string `json:"mediaType,omitempty" xmlrpc:"mediaType"`

	// The starting date of a CDN bandwidth summary.
	StartDateTime *Time `json:"startDateTime,omitempty" xmlrpc:"startDateTime"`

	// The amount of bandwidth used by a CDN account in between a given starting and ending date.
	Usage *Float64 `json:"usage,omitempty" xmlrpc:"usage"`

	// The unit of measurement used in a CDN bandwidth summary.
	UsageUnits *string `json:"usageUnits,omitempty" xmlrpc:"usageUnits"`
}

// SoftLayer_Container_Network_ContentDelivery_Bandwidth_Summary_File models a CDN account's overall bandwidth usage and overages within a given date range.
type Container_Network_ContentDelivery_Bandwidth_Summary_Detail struct {
	Container_Network_ContentDelivery_Bandwidth_Summary

	// The duration of a file that is viewed.
	Duration *Float64 `json:"duration,omitempty" xmlrpc:"duration"`

	// The number of times that a file is viewed.
	ViewCount *int `json:"viewCount,omitempty" xmlrpc:"viewCount"`
}

// SoftLayer's CDN allows for multiple origin pull domains and CNAME records. This container holds the origin pull configuration details. CDN currently supports origin pull method for HTTP content.
type Container_Network_ContentDelivery_OriginPull_Mapping struct {
	Entity

	// The CNAME record.
	Cname *string `json:"cname,omitempty" xmlrpc:"cname"`

	// The unique identifier of an origin pull configuration
	Id *string `json:"id,omitempty" xmlrpc:"id"`

	// This indicates if an origin pull mapping is for the secure content or not.
	IsSecureContent *bool `json:"isSecureContent,omitempty" xmlrpc:"isSecureContent"`

	// The type of a media supported by CDN. Supported media types are: "HTTP", "FLASH" and "WM"
	MediaType *string `json:"mediaType,omitempty" xmlrpc:"mediaType"`

	// The URL of a origin server.  A URL can contain a directory path.
	OriginUrl *string `json:"originUrl,omitempty" xmlrpc:"originUrl"`
}

// SoftLayer's CDN content delivery network offering replicates your data to a number of Points of Presence (POP's) around the world. SoftLayer_Container_Network_ContentDelivery_PointsOfPresence models one of these POP locations.
type Container_Network_ContentDelivery_PointsOfPresence struct {
	Entity

	// A CDN Point of Presence's internal identifier.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// A CDN Point of Presence's name. This is typically the city that the POP is located in.
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// This container holds information on a purge request. [[SoftLayer_Network_ContentDelivery_Account::purgeCache|Purge method]] for more details.
//
// Status code can be "SUCCESS", "FAILED", or "INVALID_URL" "INVALID_URL" code is returned when a URL is malformed or does not belong to customer. "FAILED" is returned in case there was an internal error.
type Container_Network_ContentDelivery_PurgeService_Response struct {
	Entity

	// A status code indicates whether your request was successful or not
	StatusCode *string `json:"statusCode,omitempty" xmlrpc:"statusCode"`

	// A URL that you wish to purge its cache object
	Url *string `json:"url,omitempty" xmlrpc:"url"`
}

// no documentation yet
type Container_Network_ContentDelivery_Report_Usage struct {
	Entity

	// no documentation yet
	ApplicationDeliveryNetwork *Float64 `json:"applicationDeliveryNetwork,omitempty" xmlrpc:"applicationDeliveryNetwork"`

	// no documentation yet
	ApplicationDeliveryNetworkSsl *Float64 `json:"applicationDeliveryNetworkSsl,omitempty" xmlrpc:"applicationDeliveryNetworkSsl"`

	// no documentation yet
	DiskSpace *Float64 `json:"diskSpace,omitempty" xmlrpc:"diskSpace"`

	// no documentation yet
	EndDate *Time `json:"endDate,omitempty" xmlrpc:"endDate"`

	// no documentation yet
	Flash *Float64 `json:"flash,omitempty" xmlrpc:"flash"`

	// no documentation yet
	Http *Float64 `json:"http,omitempty" xmlrpc:"http"`

	// no documentation yet
	HttpSmall *Float64 `json:"httpSmall,omitempty" xmlrpc:"httpSmall"`

	// no documentation yet
	Https *Float64 `json:"https,omitempty" xmlrpc:"https"`

	// no documentation yet
	HttpsSmall *Float64 `json:"httpsSmall,omitempty" xmlrpc:"httpsSmall"`

	// no documentation yet
	Region *string `json:"region,omitempty" xmlrpc:"region"`

	// no documentation yet
	SslTotal *Float64 `json:"sslTotal,omitempty" xmlrpc:"sslTotal"`

	// no documentation yet
	StandardTotal *Float64 `json:"standardTotal,omitempty" xmlrpc:"standardTotal"`

	// no documentation yet
	StartDate *Time `json:"startDate,omitempty" xmlrpc:"startDate"`

	// no documentation yet
	WindowsMedia *Float64 `json:"windowsMedia,omitempty" xmlrpc:"windowsMedia"`
}

// SoftLayer's CDN content delivery network allows for multiple types of media hosting in addition to traditional HTTP hosting. Each of these media types are accessible form a different URL. SoftLayer_Container_Network_ContentDelivery_SupportedProtocol holds details about CDN supported media types and their associated URLs.
//
// CDN media URLs follow the standard <protocol>://<cdn-name>.<platform-name>.cdn.softlayer.net
//
// Flash streaming, Windows Media streaming and HTTP protocols are supported: Flash streaming: <nowiki>rtmp://<cdn-name>.flash.cdn.softlayer.net</nowiki> Windows Media streaming: <nowiki>mms://<cdn-name>.wm.cdn.softlayer.net</nowiki> HTTP: <nowiki>http://<cdn-name>.http.cdn.softlayer.net</nowiki>
type Container_Network_ContentDelivery_SupportedProtocol struct {
	Entity

	// The host name related to CDN supported media, and is represented in the hostname portion of a CDN URL.
	Host *string `json:"host,omitempty" xmlrpc:"host"`

	// The type of a media supported by CDN such as "FLASH", "WINDOWSMEDIA" or "HTTP"
	MediaType *string `json:"mediaType,omitempty" xmlrpc:"mediaType"`

	// The platform name. It's a friendly name of media type.
	Platform *string `json:"platform,omitempty" xmlrpc:"platform"`

	// The media protocol supported by CDN. This represents the media portion of a CDN URL.  Currently supported protocols are: rtmp, mms and http
	Protocol *string `json:"protocol,omitempty" xmlrpc:"protocol"`
}

// SoftLayer_Container_Network_Directory_Listing represents a single entry in a listing of files within a remote directory. API methods that return remote directory listings typically return arrays of SoftLayer_Container_Network_Directory_Listing objects.
type Container_Network_Directory_Listing struct {
	Entity

	// If the file in a directory listing is a directory itself then fileCount is the number of files within the directory.
	FileCount *int `json:"fileCount,omitempty" xmlrpc:"fileCount"`

	// The name of a directory or a file within a directory listing.
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// The type of file in a directory listing. If a directory listing entry is a directory itself then type is set to "directory". Otherwise, type is a blank string.
	Type *string `json:"type,omitempty" xmlrpc:"type"`
}

// The IntrusionProtection_Event object stores information about individual intrusion protection events.
//
// It is a data container that cannot be edited, deleted, or saved.
//
// It is returned by many methods in the TippingPointReporting object, but never directly, always as a child of another container object.
type Container_Network_IntrusionProtection_Event struct {
	Entity

	// The CVE ID(s), if any, associated with this attack signature.
	CVEId *string `json:"CVEId,omitempty" xmlrpc:"CVEId"`

	// The action that was taken when this attack was discovered.  Can be either "Block" or "Permit"
	ActionTaken *string `json:"actionTaken,omitempty" xmlrpc:"actionTaken"`

	// The number of attacks in this block.  Attacks are grouped differently based on the query performed on the tippingPointReporting object.
	AttackCount *int `json:"attackCount,omitempty" xmlrpc:"attackCount"`

	// Long description of the attack.  May contain links to more information
	AttackLongDescription *string `json:"attackLongDescription,omitempty" xmlrpc:"attackLongDescription"`

	// Name of the attack
	AttackName *string `json:"attackName,omitempty" xmlrpc:"attackName"`

	// The starting timestamp of the attack recorded, in Y-m-d H:i:s format.  May not be set, depending on the type of query performed.
	BeginTime *string `json:"beginTime,omitempty" xmlrpc:"beginTime"`

	// The BugTraq ID(s), if any, associated with this attack signature.
	BugtraqId *string `json:"bugtraqId,omitempty" xmlrpc:"bugtraqId"`

	// The human-readable classification of the attack
	Classification *string `json:"classification,omitempty" xmlrpc:"classification"`

	// The IP Address (as a dotted decimal string) of the machine that was the target of the attack
	DestinationIpAddress *string `json:"destinationIpAddress,omitempty" xmlrpc:"destinationIpAddress"`

	// The port the attack was directed at
	DestinationPort *int `json:"destinationPort,omitempty" xmlrpc:"destinationPort"`

	// The ending timestamp of the attack recorded, in Y-m-d H:i:s format.  May not be set, depending on the type of query performed.
	EndTime *string `json:"endTime,omitempty" xmlrpc:"endTime"`

	// The platform affected by the attack
	Platform *string `json:"platform,omitempty" xmlrpc:"platform"`

	// The protocol used in the attack
	Protocol *string `json:"protocol,omitempty" xmlrpc:"protocol"`

	// The human-readable severity of this attack, from "Low" to "Critical"
	Severity *string `json:"severity,omitempty" xmlrpc:"severity"`

	// Unique ID of the "Signature" in question.  The signature determines the type of attack recorded.  SignatureId is used in the drillDown() function on the TippingPointReporting service
	SignatureId *string `json:"signatureId,omitempty" xmlrpc:"signatureId"`

	// The IP Address (as a dotted decimal string) of the machine originating the attack
	SourceIpAddress *string `json:"sourceIpAddress,omitempty" xmlrpc:"sourceIpAddress"`

	// The port the attack originated from
	SourcePort *int `json:"sourcePort,omitempty" xmlrpc:"sourcePort"`
}

// The IntrusionProtection_Statistic is used exclusively by the getMainStatistics method on the TippingPointReporting service, and serves mainly as a pair object, storing a name and an attack count.  Name is usually the name of an attack, but it can also be an attacking IP Address
type Container_Network_IntrusionProtection_Statistic struct {
	Entity

	// The number of attacks effecting this name over the time period
	AttackCount *int `json:"attackCount,omitempty" xmlrpc:"attackCount"`

	// Either the name of the attack in question, or the attacking IP Address
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// The IntrusionProtection_Statistics Type is used as a container for SoftLayer_Container_Network_IntrusionProtection_Statistic objects.  The SoftLayer_Container_Network_IntrusionProtection_Statistics class holds the "header" information, like the item being queried (either account or data center), the time frame, and the grand total of the attacks.
type Container_Network_IntrusionProtection_Statistics struct {
	Entity

	// The actual target, either a datacenter name, an account ID, or a subnet IP
	Target *string `json:"target,omitempty" xmlrpc:"target"`

	// The type of the target, right now either "datacenter", "account", or "subnet"
	TargetType *string `json:"targetType,omitempty" xmlrpc:"targetType"`

	// The time frame of the attack, in string form, like "Last 24 hours"
	TimeFrame *string `json:"timeFrame,omitempty" xmlrpc:"timeFrame"`

	// The top attacks for this target over this time frame
	TopAttacks []Container_Network_IntrusionProtection_Statistic `json:"topAttacks,omitempty" xmlrpc:"topAttacks"`

	// Total attacks for this $target over this time frame
	TotalAttacks *int `json:"totalAttacks,omitempty" xmlrpc:"totalAttacks"`
}

// The IntrusionProtection_SubnetReport object is the container that holds the SoftLayer_Container_Network_IntrusionProtection_Event objects for a particular subnet, or "All Subnets", whatever the case may be.  Subnet, subnet mask, direction, and the individual events are returned by this object.
type Container_Network_IntrusionProtection_SubnetReport struct {
	Entity

	// cidr for this report.  If the subnetIpAddress is "All Subnets", this is set to 32 and should be ignored.
	Cidr *int `json:"cidr,omitempty" xmlrpc:"cidr"`

	// Direction of the attack, either 'Inbound' or 'Outbound'
	Direction *string `json:"direction,omitempty" xmlrpc:"direction"`

	// The class SoftLayer_Container_Network_IntrusionProtection_Event objects on this report.
	Events []Container_Network_IntrusionProtection_Event `json:"events,omitempty" xmlrpc:"events"`

	// The "target" of this report, could be an IP address, a subnet's network identifier, or "All Subnets"
	SubnetIpAddress *string `json:"subnetIpAddress,omitempty" xmlrpc:"subnetIpAddress"`
}

// The LoadBalancer_StatusEntry object stores information about the current status of a particular load balancer service.
//
// It is a data container that cannot be edited, deleted, or saved.
//
// It is returned exclusively by the getStatus method on the [[SoftLayer_Network_LoadBalancer_Service]] service
type Container_Network_LoadBalancer_StatusEntry struct {
	Entity

	// The value of the entry.
	Content *string `json:"content,omitempty" xmlrpc:"content"`

	// Text description of the status entry
	Label *string `json:"label,omitempty" xmlrpc:"label"`
}

// This container class holds information on a media file such as file name, codec, frame rate and so on
type Container_Network_Media_Information struct {
	Entity

	// The audio bit rate
	AudioBitRate *int `json:"audioBitRate,omitempty" xmlrpc:"audioBitRate"`

	// The audio channel mode
	AudioChannelMode *string `json:"audioChannelMode,omitempty" xmlrpc:"audioChannelMode"`

	// The number of audio channels
	AudioChannels *int `json:"audioChannels,omitempty" xmlrpc:"audioChannels"`

	// The audio codec name
	AudioCodec *string `json:"audioCodec,omitempty" xmlrpc:"audioCodec"`

	// The audio sample rate
	AudioSampleRate *int `json:"audioSampleRate,omitempty" xmlrpc:"audioSampleRate"`

	// The duration of a media
	Duration *Float64 `json:"duration,omitempty" xmlrpc:"duration"`

	// The error message if any.
	ErrorMessage *string `json:"errorMessage,omitempty" xmlrpc:"errorMessage"`

	// The name of a media file
	File *string `json:"file,omitempty" xmlrpc:"file"`

	// The file format
	FileFormat *string `json:"fileFormat,omitempty" xmlrpc:"fileFormat"`

	// The size of a media file in byte
	FileSize *uint `json:"fileSize,omitempty" xmlrpc:"fileSize"`

	// The frame rate
	FrameRate *Float64 `json:"frameRate,omitempty" xmlrpc:"frameRate"`

	// The width of a media in pixel
	SizeX *int `json:"sizeX,omitempty" xmlrpc:"sizeX"`

	// The height of a media in pixel
	SizeY *int `json:"sizeY,omitempty" xmlrpc:"sizeY"`

	// The total of frames
	TotalFrames *uint `json:"totalFrames,omitempty" xmlrpc:"totalFrames"`

	// The width in a video's width to height aspect ratio
	VideoAspectX *Float64 `json:"videoAspectX,omitempty" xmlrpc:"videoAspectX"`

	// The height in a video's width to height aspect ratio
	VideoAspectY *int `json:"videoAspectY,omitempty" xmlrpc:"videoAspectY"`

	// The video codec name
	VideoCodec *string `json:"videoCodec,omitempty" xmlrpc:"videoCodec"`
}

// no documentation yet
type Container_Network_Media_Transcode_Job_Watermark struct {
	Entity

	// Time to stop showing watermark in milliseconds
	EndTime *int `json:"endTime,omitempty" xmlrpc:"endTime"`

	// Filename of image to use as watermark in transcoding job
	FileName *string `json:"fileName,omitempty" xmlrpc:"fileName"`

	// Position to place watermark at
	Position *Container_Network_Media_Transcode_Job_Watermark_Position `json:"position,omitempty" xmlrpc:"position"`

	// Time to start showing watermark in milliseconds
	StartTime *int `json:"startTime,omitempty" xmlrpc:"startTime"`

	// Text to Place in Watermark
	Text *string `json:"text,omitempty" xmlrpc:"text"`

	// Percentage Transparent watermark should be
	TransparencyPercentage *int `json:"transparencyPercentage,omitempty" xmlrpc:"transparencyPercentage"`
}

// no documentation yet
type Container_Network_Media_Transcode_Job_Watermark_Position struct {
	Entity

	// X Coordinate of Watermark
	X *int `json:"x,omitempty" xmlrpc:"x"`

	// vertical Coordinate of Watermark
	Y *int `json:"y,omitempty" xmlrpc:"y"`
}

// Transcode preset is a set of configuration parameters that defines a Transcode output format. SoftLayer_Container_Network_Media_Transcode_Preset contains a preset information defined on a Transcode server
type Container_Network_Media_Transcode_Preset struct {
	Entity

	// The unique id that is used by a Transcode server
	GUID *string `json:"GUID,omitempty" xmlrpc:"GUID"`

	// The category that a preset belongs to
	Category *string `json:"category,omitempty" xmlrpc:"category"`

	// The description of the preset
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// The friendly name of a preset
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// Transcode preset element
type Container_Network_Media_Transcode_Preset_Element struct {
	Entity

	// The additional elements for DROPDOWNLIST element
	AdditionalElements []Container_Network_Media_Transcode_Preset_Element_Option `json:"additionalElements,omitempty" xmlrpc:"additionalElements"`

	// The default value of an element.
	DefaultValue *string `json:"defaultValue,omitempty" xmlrpc:"defaultValue"`

	// The description of a preset element
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// The flag that indicates whether an element is enabled or not
	Enabled *bool `json:"enabled,omitempty" xmlrpc:"enabled"`

	// The extended description of a preset element
	ExtendedDescription *string `json:"extendedDescription,omitempty" xmlrpc:"extendedDescription"`

	// The flag that indicates whether an element is hidden or not
	Hidden *bool `json:"hidden,omitempty" xmlrpc:"hidden"`

	// The maximum value of an element
	MaximumValue *int `json:"maximumValue,omitempty" xmlrpc:"maximumValue"`

	// The minimum value of an element
	MinimumValue *int `json:"minimumValue,omitempty" xmlrpc:"minimumValue"`

	// The name of an preset element
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// The name of a parent element
	ParentName *string `json:"parentName,omitempty" xmlrpc:"parentName"`

	// The type of an preset element.
	Type *string `json:"type,omitempty" xmlrpc:"type"`
}

// Transcode preset element
type Container_Network_Media_Transcode_Preset_Element_Option struct {
	Entity

	// The name of a additional preset element
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// The value of a additional preset element
	Value *string `json:"value,omitempty" xmlrpc:"value"`
}

// no documentation yet
type Container_Network_Message_Delivery_Email struct {
	Entity

	// no documentation yet
	Body *string `json:"body,omitempty" xmlrpc:"body"`

	// no documentation yet
	ContainsHtml *bool `json:"containsHtml,omitempty" xmlrpc:"containsHtml"`

	// no documentation yet
	From *string `json:"from,omitempty" xmlrpc:"from"`

	// no documentation yet
	Subject *string `json:"subject,omitempty" xmlrpc:"subject"`

	// no documentation yet
	To *string `json:"to,omitempty" xmlrpc:"to"`
}

// no documentation yet
type Container_Network_Message_Delivery_Email_Sendgrid_Account_Overview struct {
	Entity

	// no documentation yet
	CreditsAllowed *int `json:"creditsAllowed,omitempty" xmlrpc:"creditsAllowed"`

	// no documentation yet
	CreditsOverage *int `json:"creditsOverage,omitempty" xmlrpc:"creditsOverage"`

	// no documentation yet
	CreditsRemain *int `json:"creditsRemain,omitempty" xmlrpc:"creditsRemain"`

	// no documentation yet
	CreditsUsed *int `json:"creditsUsed,omitempty" xmlrpc:"creditsUsed"`

	// no documentation yet
	Package *string `json:"package,omitempty" xmlrpc:"package"`

	// no documentation yet
	Reputation *int `json:"reputation,omitempty" xmlrpc:"reputation"`

	// no documentation yet
	Requests *int `json:"requests,omitempty" xmlrpc:"requests"`
}

// no documentation yet
type Container_Network_Message_Delivery_Email_Sendgrid_Customer_Profile struct {
	Entity

	// no documentation yet
	Address *string `json:"address,omitempty" xmlrpc:"address"`

	// no documentation yet
	City *string `json:"city,omitempty" xmlrpc:"city"`

	// no documentation yet
	Country *string `json:"country,omitempty" xmlrpc:"country"`

	// no documentation yet
	Email *string `json:"email,omitempty" xmlrpc:"email"`

	// no documentation yet
	FirstName *string `json:"firstName,omitempty" xmlrpc:"firstName"`

	// no documentation yet
	LastName *string `json:"lastName,omitempty" xmlrpc:"lastName"`

	// no documentation yet
	Phone *string `json:"phone,omitempty" xmlrpc:"phone"`

	// no documentation yet
	State *string `json:"state,omitempty" xmlrpc:"state"`

	// no documentation yet
	Website *string `json:"website,omitempty" xmlrpc:"website"`

	// no documentation yet
	Zip *string `json:"zip,omitempty" xmlrpc:"zip"`
}

// no documentation yet
type Container_Network_Message_Delivery_Email_Sendgrid_List_Entry struct {
	Entity

	// no documentation yet
	Created *string `json:"created,omitempty" xmlrpc:"created"`

	// no documentation yet
	Email *string `json:"email,omitempty" xmlrpc:"email"`

	// no documentation yet
	Reason *string `json:"reason,omitempty" xmlrpc:"reason"`

	// no documentation yet
	Status *string `json:"status,omitempty" xmlrpc:"status"`
}

// no documentation yet
type Container_Network_Message_Delivery_Email_Sendgrid_Statistics struct {
	Entity

	// no documentation yet
	Blocks *int `json:"blocks,omitempty" xmlrpc:"blocks"`

	// no documentation yet
	Bounces *int `json:"bounces,omitempty" xmlrpc:"bounces"`

	// no documentation yet
	Clicks *int `json:"clicks,omitempty" xmlrpc:"clicks"`

	// no documentation yet
	Date *string `json:"date,omitempty" xmlrpc:"date"`

	// no documentation yet
	Delivered *int `json:"delivered,omitempty" xmlrpc:"delivered"`

	// no documentation yet
	InvalidEmail *int `json:"invalidEmail,omitempty" xmlrpc:"invalidEmail"`

	// no documentation yet
	Opens *int `json:"opens,omitempty" xmlrpc:"opens"`

	// no documentation yet
	RepeatBounces *int `json:"repeatBounces,omitempty" xmlrpc:"repeatBounces"`

	// no documentation yet
	RepeatSpamReports *int `json:"repeatSpamReports,omitempty" xmlrpc:"repeatSpamReports"`

	// no documentation yet
	RepeatUnsubscribes *int `json:"repeatUnsubscribes,omitempty" xmlrpc:"repeatUnsubscribes"`

	// no documentation yet
	Requests *int `json:"requests,omitempty" xmlrpc:"requests"`

	// no documentation yet
	SpamReports *int `json:"spamReports,omitempty" xmlrpc:"spamReports"`

	// no documentation yet
	UniqueClicks *int `json:"uniqueClicks,omitempty" xmlrpc:"uniqueClicks"`

	// no documentation yet
	UniqueOpens *int `json:"uniqueOpens,omitempty" xmlrpc:"uniqueOpens"`

	// no documentation yet
	Unsubscribes *int `json:"unsubscribes,omitempty" xmlrpc:"unsubscribes"`
}

// no documentation yet
type Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Graph struct {
	Entity

	// no documentation yet
	GraphError *string `json:"graphError,omitempty" xmlrpc:"graphError"`

	// no documentation yet
	GraphImage *[]byte `json:"graphImage,omitempty" xmlrpc:"graphImage"`

	// no documentation yet
	GraphTitle *string `json:"graphTitle,omitempty" xmlrpc:"graphTitle"`
}

// no documentation yet
type Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Options struct {
	Entity

	// no documentation yet
	AggregatesOnly *bool `json:"aggregatesOnly,omitempty" xmlrpc:"aggregatesOnly"`

	// no documentation yet
	Category *string `json:"category,omitempty" xmlrpc:"category"`

	// no documentation yet
	Days *int `json:"days,omitempty" xmlrpc:"days"`

	// no documentation yet
	EndDate *Time `json:"endDate,omitempty" xmlrpc:"endDate"`

	// no documentation yet
	SelectedStatistics []string `json:"selectedStatistics,omitempty" xmlrpc:"selectedStatistics"`

	// no documentation yet
	StartDate *Time `json:"startDate,omitempty" xmlrpc:"startDate"`
}

// no documentation yet
type Container_Network_Port_Statistic struct {
	Entity

	// no documentation yet
	AdministrativeStatus *int `json:"administrativeStatus,omitempty" xmlrpc:"administrativeStatus"`

	// no documentation yet
	InDiscardPackets *uint `json:"inDiscardPackets,omitempty" xmlrpc:"inDiscardPackets"`

	// no documentation yet
	InErrorPackets *uint `json:"inErrorPackets,omitempty" xmlrpc:"inErrorPackets"`

	// no documentation yet
	InOctets *uint `json:"inOctets,omitempty" xmlrpc:"inOctets"`

	// no documentation yet
	InUnicastPackets *uint `json:"inUnicastPackets,omitempty" xmlrpc:"inUnicastPackets"`

	// no documentation yet
	MaximumTransmissionUnit *uint `json:"maximumTransmissionUnit,omitempty" xmlrpc:"maximumTransmissionUnit"`

	// no documentation yet
	OperationalStatus *int `json:"operationalStatus,omitempty" xmlrpc:"operationalStatus"`

	// no documentation yet
	OutDiscardPackets *uint `json:"outDiscardPackets,omitempty" xmlrpc:"outDiscardPackets"`

	// no documentation yet
	OutErrorPackets *uint `json:"outErrorPackets,omitempty" xmlrpc:"outErrorPackets"`

	// no documentation yet
	OutOctets *uint `json:"outOctets,omitempty" xmlrpc:"outOctets"`

	// no documentation yet
	OutUnicastPackets *uint `json:"outUnicastPackets,omitempty" xmlrpc:"outUnicastPackets"`

	// no documentation yet
	PortDuplex *uint `json:"portDuplex,omitempty" xmlrpc:"portDuplex"`

	// no documentation yet
	Speed *uint `json:"speed,omitempty" xmlrpc:"speed"`
}

// no documentation yet
type Container_Network_Service_Resource_ObjectStorage_ConnectionInformation struct {
	Entity

	// no documentation yet
	Datacenter *string `json:"datacenter,omitempty" xmlrpc:"datacenter"`

	// no documentation yet
	DatacenterShortName *string `json:"datacenterShortName,omitempty" xmlrpc:"datacenterShortName"`

	// no documentation yet
	PrivateEndpoint *string `json:"privateEndpoint,omitempty" xmlrpc:"privateEndpoint"`

	// no documentation yet
	PublicEndpoint *string `json:"publicEndpoint,omitempty" xmlrpc:"publicEndpoint"`
}

// no documentation yet
type Container_Network_Storage_Backup_Evault_WebCc_Authentication_Details struct {
	Entity

	// no documentation yet
	EventValidation *string `json:"eventValidation,omitempty" xmlrpc:"eventValidation"`

	// no documentation yet
	ViewState *string `json:"viewState,omitempty" xmlrpc:"viewState"`

	// no documentation yet
	WebCcUrl *string `json:"webCcUrl,omitempty" xmlrpc:"webCcUrl"`
}

// SoftLayer's StorageLayer Evault services provides details regarding the the purchased vault.
//
// When a job is created using the Webcc Console, the job created is identified as a task on the vault. Using this service, information regarding the task can be retrieved.
//
//
type Container_Network_Storage_Evault_Vault_Task struct {
	Entity

	// Unique identifier for the task.
	Id *uint `json:"id,omitempty" xmlrpc:"id"`

	// The hostname provided at time of agent registration.
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// Total compressed bytes used for the task.
	UsedPoolsize *uint `json:"usedPoolsize,omitempty" xmlrpc:"usedPoolsize"`
}

// The SoftLayer_Container_Network_Storage_Evault_WebCc_AgentStatus will contain the timestamp of the last backup performed by the EVault agent.  The agent status will also be returned.
type Container_Network_Storage_Evault_WebCc_AgentStatus struct {
	Entity

	// Timestamp of last backup performed by the EVault backup agent
	LastBackup *Time `json:"lastBackup,omitempty" xmlrpc:"lastBackup"`

	// Status indicating the accumulative status result of all jobs performed by the evault agent.  For example, if one job out three jobs failed agent status will by "Failed Backup(s)".
	Status *string `json:"status,omitempty" xmlrpc:"status"`
}

// The SoftLayer_Container_Network_Storage_Evault_WebCc_BackupResults will contain the timeframe of backups and the results will also be returned.
type Container_Network_Storage_Evault_WebCc_BackupResults struct {
	Entity

	// Timestamp of begin time
	BeginTime *Time `json:"beginTime,omitempty" xmlrpc:"beginTime"`

	// Count of backups with conflicts.
	Conflict *string `json:"conflict,omitempty" xmlrpc:"conflict"`

	// Timestamp of end time
	EndTime *Time `json:"endTime,omitempty" xmlrpc:"endTime"`

	// Count of failed backups.
	Failed *string `json:"failed,omitempty" xmlrpc:"failed"`

	// Count of successfull backups.
	Success *string `json:"success,omitempty" xmlrpc:"success"`
}

// The SoftLayer_Container_Network_Storage_Evault_WebCc_JobDetails will contain basic details for all backup and restore jobs performed by the StorageLayer EVault service offering.
type Container_Network_Storage_Evault_WebCc_JobDetails struct {
	Entity

	// The number of bytes currently used by the backup job. (provided only for backup jobs)
	BytesUsed *uint `json:"bytesUsed,omitempty" xmlrpc:"bytesUsed"`

	// Description of the backup/restore job
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// hardware id
	HardwareId *int `json:"hardwareId,omitempty" xmlrpc:"hardwareId"`

	// Date of the last jobrun.
	LastRunDate *Time `json:"lastRunDate,omitempty" xmlrpc:"lastRunDate"`

	// Name of the backup/restore job
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// Size of backup job when it was first run. (provided only for backup jobs)
	OriginalSize *uint `json:"originalSize,omitempty" xmlrpc:"originalSize"`

	// Percentage of overall used space allocated by the job. (provided only for backup jobs)
	PercentageOfTotalUsage *int `json:"percentageOfTotalUsage,omitempty" xmlrpc:"percentageOfTotalUsage"`

	// Result of the latest jobrun.
	Result *string `json:"result,omitempty" xmlrpc:"result"`

	// virtual guest id
	VirtualGuestId *int `json:"virtualGuestId,omitempty" xmlrpc:"virtualGuestId"`
}

// The SoftLayer_Container_Network_Storage_Host will contain the reference id field for the object associated with the host.  The host object type will also be returned.
type Container_Network_Storage_Host struct {
	Entity

	// Reference id field for object associated with host.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// Type for the object associated with host
	ObjectType *string `json:"objectType,omitempty" xmlrpc:"objectType"`
}

// SoftLayer_Container_Network_Storage_Hub_ObjectStorage_ContentDeliveryUrl provides specific details is a container which contains the cdn urls associated with an object storage account
type Container_Network_Storage_Hub_ObjectStorage_ContentDeliveryUrl struct {
	Entity

	// no documentation yet
	Datacenter *string `json:"datacenter,omitempty" xmlrpc:"datacenter"`

	// no documentation yet
	FlashUrl *string `json:"flashUrl,omitempty" xmlrpc:"flashUrl"`

	// no documentation yet
	HttpUrl *string `json:"httpUrl,omitempty" xmlrpc:"httpUrl"`
}

// SoftLayer_Container_Network_Storage_Hub_ObjectStorage_Endpoint provides specific details on available endpoint URLs and locations.
type Container_Network_Storage_Hub_ObjectStorage_Endpoint struct {
	Entity

	// no documentation yet
	Location *string `json:"location,omitempty" xmlrpc:"location"`

	// no documentation yet
	Region *string `json:"region,omitempty" xmlrpc:"region"`

	// no documentation yet
	Type *string `json:"type,omitempty" xmlrpc:"type"`

	// no documentation yet
	Url *string `json:"url,omitempty" xmlrpc:"url"`
}

// SoftLayer_Container_Network_Storage_Hub_ObjectStorage_File provides specific details that only apply to files that are sent or received from CloudLayer storage resources.
type Container_Network_Storage_Hub_ObjectStorage_File struct {
	Container_Utility_File_Entity

	// no documentation yet
	Folder *string `json:"folder,omitempty" xmlrpc:"folder"`

	// no documentation yet
	Hash *string `json:"hash,omitempty" xmlrpc:"hash"`
}

// SoftLayer_Container_Network_Storage_Hub_Container provides details about containers which store collections of files.
type Container_Network_Storage_Hub_ObjectStorage_Folder struct {
	Entity

	// no documentation yet
	Bytes *uint `json:"bytes,omitempty" xmlrpc:"bytes"`

	// no documentation yet
	Count *uint `json:"count,omitempty" xmlrpc:"count"`

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// SoftLayer_Container_Network_Storage_Hub_ObjectStorage_Node provides detailed information for a particular object storage node
type Container_Network_Storage_Hub_ObjectStorage_Node struct {
	Entity

	// no documentation yet
	DeviceName *string `json:"deviceName,omitempty" xmlrpc:"deviceName"`

	// no documentation yet
	ResourceName *string `json:"resourceName,omitempty" xmlrpc:"resourceName"`

	// no documentation yet
	UserAuthUrl *string `json:"userAuthUrl,omitempty" xmlrpc:"userAuthUrl"`
}

// SoftLayer_Container_Network_Storage_Hub_ObjectStorage_Policy provides specific details on available storage policies.
type Container_Network_Storage_Hub_ObjectStorage_Policy struct {
	Entity

	// no documentation yet
	PolicyCode *string `json:"policyCode,omitempty" xmlrpc:"policyCode"`
}

// no documentation yet
type Container_Network_Storage_NetworkConnectionInformation struct {
	Entity

	// no documentation yet
	Id *string `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	IpAddress *string `json:"ipAddress,omitempty" xmlrpc:"ipAddress"`

	// no documentation yet
	StorageType *string `json:"storageType,omitempty" xmlrpc:"storageType"`
}

// SoftLayer_Container_Subnet_IPAddress models an IP v4 address as it exists as a member of it's subnet, letting the user know if it is a network identifier, gateway, broadcast, or useable address. Addresses that are neither the network identifier nor the gateway nor the broadcast addresses are usable by SoftLayer servers.
type Container_Network_Subnet_IpAddress struct {
	Entity

	// The hardware that an IP address is associated with.
	Hardware *Hardware `json:"hardware,omitempty" xmlrpc:"hardware"`

	// An IP address expressed in dotted-quad notation.
	IpAddress *string `json:"ipAddress,omitempty" xmlrpc:"ipAddress"`

	// Whether an IP address is its subnet's broadcast address.
	IsBroadcastAddress *bool `json:"isBroadcastAddress,omitempty" xmlrpc:"isBroadcastAddress"`

	// Whether an IP address is its subnet's gateway address. Gateway addresses exist on SoftLayer's routers and many not be assigned to servers.
	IsGatewayAddress *bool `json:"isGatewayAddress,omitempty" xmlrpc:"isGatewayAddress"`

	// Whether an IP address is its subnet's network identifier address.
	IsNetworkAddress *bool `json:"isNetworkAddress,omitempty" xmlrpc:"isNetworkAddress"`
}

// SoftLayer_Container_Network_Subnet_Registration_SubnetReference is provided to reference [[SoftLayer_Network_Subnet_Registration]] object and the [[SoftLayer_Network_Subnet]] it references, in CIDR form.
type Container_Network_Subnet_Registration_SubnetReference struct {
	Entity

	// The ID of the [[SoftLayer_Network_Subnet_Registration]] object.
	RegistrationId *int `json:"registrationId,omitempty" xmlrpc:"registrationId"`

	// The subnet address in CIDR form.
	SubnetCidr *string `json:"subnetCidr,omitempty" xmlrpc:"subnetCidr"`
}

// SoftLayer_Container_Subnet_Registration_TransactionDetails is provided to return details of a newly created Subnet Registration Transaction.
type Container_Network_Subnet_Registration_TransactionDetails struct {
	Entity

	// The IDs and Subnets of the [[SoftLayer_Network_Subnet_Registration]] object.
	SubnetReferences []Container_Network_Subnet_Registration_SubnetReference `json:"subnetReferences,omitempty" xmlrpc:"subnetReferences"`

	// The ID of the Transaction object.
	TransactionId *int `json:"transactionId,omitempty" xmlrpc:"transactionId"`
}

// no documentation yet
type Container_Notification_Mass_Filter_TemplateKey struct {
	Entity
}

// no documentation yet
type Container_Notification_Mass_Filter_TemplateValue struct {
	Entity
}

// Represents the acceptance status of a Policy.
type Container_Policy_Acceptance struct {
	Entity

	// Flag to indicate if a policy has been previously accepted.
	AcceptedFlag *bool `json:"acceptedFlag,omitempty" xmlrpc:"acceptedFlag"`

	// Name of the policy for which we are representing it's acceptance status.
	PolicyName *string `json:"policyName,omitempty" xmlrpc:"policyName"`

	// ID of the [[SoftLayer_Product_Item_Policy_Assignment]].
	ProductPolicyAssignmentId *int `json:"productPolicyAssignmentId,omitempty" xmlrpc:"productPolicyAssignmentId"`
}

// The SoftLayer_Container_Product_Item_Category data type represents a single product item category.
type Container_Product_Item_Category struct {
	Entity

	// identifier for category.
	Id *int `json:"id,omitempty" xmlrpc:"id"`
}

// The SoftLayer_Container_Product_Item_Category_Question_Answer data type represents an answer to an item category question.  It contains the category, the question being answered, and the answer.
type Container_Product_Item_Category_Question_Answer struct {
	Entity

	// The answer to the question.
	Answer *string `json:"answer,omitempty" xmlrpc:"answer"`

	// The product item category code.
	CategoryCode *string `json:"categoryCode,omitempty" xmlrpc:"categoryCode"`

	// The product item category id.
	CategoryId *int `json:"categoryId,omitempty" xmlrpc:"categoryId"`

	// The product item category question id.
	QuestionId *int `json:"questionId,omitempty" xmlrpc:"questionId"`
}

// The SoftLayer_Container_Product_Item_Category_ZeroFee_Count data type represents a count of zero fee billing/invoice items.
type Container_Product_Item_Category_ZeroFee_Count struct {
	Entity

	// The product item category code.
	CategoryCode *string `json:"categoryCode,omitempty" xmlrpc:"categoryCode"`

	// The product item category id.
	CategoryId *int `json:"categoryId,omitempty" xmlrpc:"categoryId"`

	// The product item category name.
	CategoryName *string `json:"categoryName,omitempty" xmlrpc:"categoryName"`

	// The count of zero fee items for this category.
	Count *int `json:"count,omitempty" xmlrpc:"count"`
}

// The SoftLayer_Container_Product_Item_Discount_Program data type represents the information about a discount that is related to a specific product item.
type Container_Product_Item_Discount_Program struct {
	Entity

	// The number of times the item discount(s) may be applied for that order container.  At a minimum the number will be 1 and at most, it will match the quantity of the order container.
	ApplicableQuantity *int `json:"applicableQuantity,omitempty" xmlrpc:"applicableQuantity"`

	// The product item that the discount applies to.
	Item *Product_Item `json:"item,omitempty" xmlrpc:"item"`

	// The sum of the one time fees (one time, setup and labor) of the prices of this container multiplied by the applicable quantity of this container.
	OneTimeAmount *Float64 `json:"oneTimeAmount,omitempty" xmlrpc:"oneTimeAmount"`

	// The tax amount on the one time fees (one time, setup and labor) of the prices of this container mulitiplied by the applicable quantity of this container.
	OneTimeTax *Float64 `json:"oneTimeTax,omitempty" xmlrpc:"oneTimeTax"`

	// The item prices that contain the amount of the discount in the recurringFee field.  There may be one or more prices.
	Prices []Product_Item_Price `json:"prices,omitempty" xmlrpc:"prices"`

	// The sum of the one time fees (one time, setup and labor) of the prices of this container multiplied by the applicable quantity of this container with the proration factor applied.
	ProratedOneTimeAmount *Float64 `json:"proratedOneTimeAmount,omitempty" xmlrpc:"proratedOneTimeAmount"`

	// The tax amount on the one time fees (one time, setup and labor) of the prices of this container mulitiplied by the applicable quantity of this container with the proration factor applied.
	ProratedOneTimeTax *Float64 `json:"proratedOneTimeTax,omitempty" xmlrpc:"proratedOneTimeTax"`

	// The sum of the recurring fees of the prices of this container multiplied by the applicable quantity of this container with the proration factor applied.
	ProratedRecurringAmount *Float64 `json:"proratedRecurringAmount,omitempty" xmlrpc:"proratedRecurringAmount"`

	// The tax amount on the recurring fees of the prices of this container mulitiplied by the applicable quantity of this container with the proration factor applied.
	ProratedRecurringTax *Float64 `json:"proratedRecurringTax,omitempty" xmlrpc:"proratedRecurringTax"`

	// The sum of the recurring fees of the prices of this container multiplied by the applicable quantity of this container.
	RecurringAmount *Float64 `json:"recurringAmount,omitempty" xmlrpc:"recurringAmount"`

	// The tax amount on the recurring fees of the prices of this container mulitiplied by the applicable quantity of this container.
	RecurringTax *Float64 `json:"recurringTax,omitempty" xmlrpc:"recurringTax"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place an order with SoftLayer.
type Container_Product_Order struct {
	Entity

	// Flag for identifying an order for Big Data Deployment.
	BigDataOrderFlag *bool `json:"bigDataOrderFlag,omitempty" xmlrpc:"bigDataOrderFlag"`

	// Billing Information associated with an order. For existing customers this information is completely ignored. Do not send this information for existing customers.
	BillingInformation *Container_Product_Order_Billing_Information `json:"billingInformation,omitempty" xmlrpc:"billingInformation"`

	// This is the ID of the [[SoftLayer_Billing_Order_Item]] of this configuration/container. It is used for rebuilding an order container from a quote and is set automatically.
	BillingOrderItemId *int `json:"billingOrderItemId,omitempty" xmlrpc:"billingOrderItemId"`

	// The URL to which PayPal redirects browser after checkout has been canceled before completion of a payment.
	CancelUrl *string `json:"cancelUrl,omitempty" xmlrpc:"cancelUrl"`

	// Added by Gopherlayer. This hints to the API what kind of product order this is.
	ComplexType *string `json:"complexType,omitempty" xmlrpc:"complexType"`

	// User-specified description to identify a particular order container. This is useful if you have a multi-configuration order (multiple <code>orderContainers</code>) and you want to be able to easily determine one from another. Populating this value may be helpful if an exception is thrown when placing an order and it's tied to a specific order container.
	ContainerIdentifier *string `json:"containerIdentifier,omitempty" xmlrpc:"containerIdentifier"`

	// This hash is internally-generated and is used to for tracking order containers.
	ContainerSplHash *string `json:"containerSplHash,omitempty" xmlrpc:"containerSplHash"`

	// The currency type chosen at checkout.
	CurrencyShortName *string `json:"currencyShortName,omitempty" xmlrpc:"currencyShortName"`

	// Device Fingerprint Identifier - Optional.
	DeviceFingerprintId *string `json:"deviceFingerprintId,omitempty" xmlrpc:"deviceFingerprintId"`

	// This is the configuration identifier for tracking orders on the HTML order forms.
	DisplayLayerSessionId *string `json:"displayLayerSessionId,omitempty" xmlrpc:"displayLayerSessionId"`

	// no documentation yet
	ExtendedHardwareTesting *bool `json:"extendedHardwareTesting,omitempty" xmlrpc:"extendedHardwareTesting"`

	// The [[SoftLayer_Product_Item_Price]] for the Flexible Credit Program discount.  The <code>oneTimeFee</code> field contains the calculated discount being applied to the order.
	FlexibleCreditProgramPrice *Product_Item_Price `json:"flexibleCreditProgramPrice,omitempty" xmlrpc:"flexibleCreditProgramPrice"`

	// For orders that contain servers (bare metal, virtual server, big data, etc.), the hardware property is required. This property is an array of [[SoftLayer_Hardware]] objects. The <code>hostname</code> and <code>domain</code> properties are required for each hardware object. Note that virtual server ([[SoftLayer_Container_Product_Order_Virtual_Guest]]) orders may populate this field instead of the <code>virtualGuests</code> property.
	Hardware []Hardware `json:"hardware,omitempty" xmlrpc:"hardware"`

	// An optional virtual disk image template identifier to be used as an installation base for a computing instance order
	ImageTemplateGlobalIdentifier *string `json:"imageTemplateGlobalIdentifier,omitempty" xmlrpc:"imageTemplateGlobalIdentifier"`

	// An optional virtual disk image template identifier to be used as an installation base for a computing instance order
	ImageTemplateId *int `json:"imageTemplateId,omitempty" xmlrpc:"imageTemplateId"`

	// Flag to identify a "managed" order. This value is set internally.
	IsManagedOrder *int `json:"isManagedOrder,omitempty" xmlrpc:"isManagedOrder"`

	// The collection of [[SoftLayer_Container_Product_Item_Category_Question_Answer]] for any product category that has additional questions requiring user input.
	ItemCategoryQuestionAnswers []Container_Product_Item_Category_Question_Answer `json:"itemCategoryQuestionAnswers,omitempty" xmlrpc:"itemCategoryQuestionAnswers"`

	// The [[SoftLayer_Location_Region]] keyname or specific [[SoftLayer_Location_Datacenter]] id where the order should be provisioned. If this value is provided and the <code>regionalGroup</code> property is also specified, an exception will be thrown indicating that only 1 is allowed.
	Location *string `json:"location,omitempty" xmlrpc:"location"`

	// This [[SoftLayer_Location]] object will be determined from the <code>location</code> property and will be returned in the order verification or placement response. Any value specified here will get overwritten by the verification process.
	LocationObject *Location `json:"locationObject,omitempty" xmlrpc:"locationObject"`

	// A generic message about the order. Does not need to be sent in with any orders.
	Message *string `json:"message,omitempty" xmlrpc:"message"`

	// Orders may contain an array of configurations. Populating this property allows you to purchase multiple configurations within a single order. Each order container will have its own individual settings independent of the other order containers. For example, it is possible to order a bare metal server in one configuration and a virtual server in another.
	//
	// If <code>orderContainers</code> is populated on the base order container, most of the configuration-specific properties are ignored on the base container. For example, <code>prices</code>, <code>location</code> and <code>packageId</code> will be ignored on the base container, but since the <code>billingInformation</code> is a property that's not specific to a single order container (but the order as a whole) it must be populated on the base container.
	OrderContainers []Container_Product_Order `json:"orderContainers,omitempty" xmlrpc:"orderContainers"`

	// This is deprecated and does not do anything.
	OrderHostnames []string `json:"orderHostnames,omitempty" xmlrpc:"orderHostnames"`

	// Collection of exceptions resulting from the verification of the order. This value is set internally and is not required for end users when placing an order. When placing API orders, users can use this value to determine the container-specific exception that was thrown.
	OrderVerificationExceptions []Container_Exception `json:"orderVerificationExceptions,omitempty" xmlrpc:"orderVerificationExceptions"`

	// The [[SoftLayer_Product_Package]] id for an order container. This is required to place an order.
	PackageId *int `json:"packageId,omitempty" xmlrpc:"packageId"`

	// The Payment Type is Optional. If nothing is sent in, then the normal method of payment will be used. For paypal customers, this means a paypalToken will be returned in the receipt. This token is to be used on the paypal website to complete the order. For Credit Card customers, the card on file in our system will be used to make an initial authorization. To force the order to use a payment type, use one of the following: CARD_ON_FILE or PAYPAL
	PaymentType *string `json:"paymentType,omitempty" xmlrpc:"paymentType"`

	// The post-tax recurring charge for the order. This is the sum of preTaxRecurring + totalRecurringTax.
	PostTaxRecurring *Float64 `json:"postTaxRecurring,omitempty" xmlrpc:"postTaxRecurring"`

	// The post-tax recurring hourly charge for the order. Since taxes are not calculated for hourly orders, this value will be the same as preTaxRecurringHourly.
	PostTaxRecurringHourly *Float64 `json:"postTaxRecurringHourly,omitempty" xmlrpc:"postTaxRecurringHourly"`

	// The post-tax recurring monthly charge for the order. This is the sum of preTaxRecurringMonthly + totalRecurringTax.
	PostTaxRecurringMonthly *Float64 `json:"postTaxRecurringMonthly,omitempty" xmlrpc:"postTaxRecurringMonthly"`

	// The post-tax setup fees of the order. This is the sum of preTaxSetup + totalSetupTax;
	PostTaxSetup *Float64 `json:"postTaxSetup,omitempty" xmlrpc:"postTaxSetup"`

	// The pre-tax recurring total of the order. If there are mixed monthly and hourly prices on the order, this will be the sum of preTaxRecurringHourly and preTaxRecurringMonthly.
	PreTaxRecurring *Float64 `json:"preTaxRecurring,omitempty" xmlrpc:"preTaxRecurring"`

	// The pre-tax hourly recurring total of the order. If there are only monthly prices on the order, this value will be 0.
	PreTaxRecurringHourly *Float64 `json:"preTaxRecurringHourly,omitempty" xmlrpc:"preTaxRecurringHourly"`

	// The pre-tax monthly recurring total of the order. If there are only hourly prices on the order, this value will be 0.
	PreTaxRecurringMonthly *Float64 `json:"preTaxRecurringMonthly,omitempty" xmlrpc:"preTaxRecurringMonthly"`

	// The pre-tax setup fee total of the order.
	PreTaxSetup *Float64 `json:"preTaxSetup,omitempty" xmlrpc:"preTaxSetup"`

	// If there are any presale events available for an order, this value will be populated. It is set internally and is not required for end users when placing an order. See [[SoftLayer_Sales_Presale_Event]] for more info.
	PresaleEvent *Sales_Presale_Event `json:"presaleEvent,omitempty" xmlrpc:"presaleEvent"`

	// A preset configuration id for the package. Is required if not submitting any prices.
	PresetId *int `json:"presetId,omitempty" xmlrpc:"presetId"`

	// This is a collection of [[SoftLayer_Product_Item_Price]] objects. The only required property to populate for an item price object when ordering is its <code>id</code> - all other supplied information about the price (e.g., recurringFee, setupFee, etc.) will be ignored. Unless the [[SoftLayer_Product_Package]] associated with the order allows for preset prices, this property is required to place an order.
	Prices []Product_Item_Price `json:"prices,omitempty" xmlrpc:"prices"`

	// The id of a [[SoftLayer_Hardware_Component_Partition_Template]]. This property is optional. If no partition template is provided, a default will be used according to the operating system chosen with the order. Using the [[SoftLayer_Hardware_Component_Partition_OperatingSystem]] service, getPartitionTemplates will return those available for the particular operating system.
	PrimaryDiskPartitionId *int `json:"primaryDiskPartitionId,omitempty" xmlrpc:"primaryDiskPartitionId"`

	// Priorities to set on replication set servers.
	Priorities []string `json:"priorities,omitempty" xmlrpc:"priorities"`

	// Flag for identifying a container as Virtual Server (Private Node).
	PrivateCloudOrderFlag *bool `json:"privateCloudOrderFlag,omitempty" xmlrpc:"privateCloudOrderFlag"`

	// Type of Virtual Server (Private Node) order. Potential values: INITIAL, ADDHOST, ADDIPS, ADDZONE
	PrivateCloudOrderType *string `json:"privateCloudOrderType,omitempty" xmlrpc:"privateCloudOrderType"`

	// Optional promotion code for an order.
	PromotionCode *string `json:"promotionCode,omitempty" xmlrpc:"promotionCode"`

	// Generic properties.
	Properties []Container_Product_Order_Property `json:"properties,omitempty" xmlrpc:"properties"`

	// The Prorated Initial Charge plus the balance on the account. Only the recurring fees are prorated. Here's how the calculation works: We take the postTaxRecurring value and we prorate it based on the time between now and the next bill date for this account. After this, we add in the setup fee since this is not prorated. Then, if there is a balance on the account, we add that to the account. In the event that there is a credit balance on the account, we will subtract this amount from the order total. If the credit balance on the account is greater than the prorated initial charge, the order will go through without a charge to the credit card on the account or requiring a paypal payment. The credit on the account will be reduced by the order total, and the order will await approval from sales, as normal. If there is a pending order already in the system, We will ignore the balance on the account completely, in the calculation of the initial charge. This is to protect against two orders coming into the system and getting the benefit of a credit balance, or worse, both orders being charged the order amount + the balance on the account.
	ProratedInitialCharge *Float64 `json:"proratedInitialCharge,omitempty" xmlrpc:"proratedInitialCharge"`

	// This is the same as the proratedInitialCharge, except the balance on the account is ignored. This is the prorated total amount of the order.
	ProratedOrderTotal *Float64 `json:"proratedOrderTotal,omitempty" xmlrpc:"proratedOrderTotal"`

	// The URLs for scripts to execute on their respective servers after they have been provisioned. Provision scripts are not available for Microsoft Windows servers.
	ProvisionScripts []string `json:"provisionScripts,omitempty" xmlrpc:"provisionScripts"`

	// The quantity of the item being ordered
	Quantity *int `json:"quantity,omitempty" xmlrpc:"quantity"`

	// A custom name to be assigned to the quote.
	QuoteName *string `json:"quoteName,omitempty" xmlrpc:"quoteName"`

	// Specifying a regional group name allows you to not worry about placing your server or service at a specific datacenter, but to any datacenter within that regional group. See [[SoftLayer_Location_Group_Regional]] to get a list of available regional group names.
	//
	// <code>location</code> and <code>regionalGroup</code> are mutually exclusive on an order container. If both location and regionalGroup are provided, an exception will be thrown indicating that only 1 is allowed.
	//
	// If a regional group is provided and VLANs are specified (within the <code>hardware</code> or <code>virtualGuests</code> properties), we will use the datacenter where the VLANs are located. If no VLANs are specified, we will use the preferred datacenter on the regional group object.
	RegionalGroup *string `json:"regionalGroup,omitempty" xmlrpc:"regionalGroup"`

	// An optional resource group identifier specifying the resource group to attach the order to
	ResourceGroupId *int `json:"resourceGroupId,omitempty" xmlrpc:"resourceGroupId"`

	// This variable specifies the name of the resource group the server configuration belongs to. For MongoDB Replica sets, it would be the replica set name.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" xmlrpc:"resourceGroupName"`

	// An optional resource group template identifier to be used as a deployment base for a Virtual Server (Private Node) order.
	ResourceGroupTemplateId *int `json:"resourceGroupTemplateId,omitempty" xmlrpc:"resourceGroupTemplateId"`

	// The URL to which PayPal redirects browser after a payment is completed.
	ReturnUrl *string `json:"returnUrl,omitempty" xmlrpc:"returnUrl"`

	// This flag indicates that the quote should be sent to the email address associated with the account or order.
	SendQuoteEmailFlag *bool `json:"sendQuoteEmailFlag,omitempty" xmlrpc:"sendQuoteEmailFlag"`

	// The number of cores for the server being ordered. This value is set internally.
	ServerCoreCount *int `json:"serverCoreCount,omitempty" xmlrpc:"serverCoreCount"`

	// An optional computing instance identifier to be used as an installation base for a computing instance order
	SourceVirtualGuestId *int `json:"sourceVirtualGuestId,omitempty" xmlrpc:"sourceVirtualGuestId"`

	// The containers which hold SoftLayer_Security_Ssh_Key IDs to add to their respective servers. The order of containers passed in needs to match the order they are assigned to either hardware or virtualGuests. SSH Keys will not be assigned for servers with Microsoft Windows.
	SshKeys []Container_Product_Order_SshKeys `json:"sshKeys,omitempty" xmlrpc:"sshKeys"`

	// An optional parameter for step-based order processing.
	StepId *int `json:"stepId,omitempty" xmlrpc:"stepId"`

	//
	//
	// For orders that want to add storage groups such as RAID across multiple disks, simply add [[SoftLayer_Container_Product_Order_Storage_Group]] objects to this array. Storage groups will only be used if the 'RAID' disk controller price is selected. Any other disk controller types will ignore the storage groups set here.
	//
	// The first storage group in this array will be considered the primary storage group, which is used for the OS. Any other storage groups will act as data storage.
	//
	//
	StorageGroups []Container_Product_Order_Storage_Group `json:"storageGroups,omitempty" xmlrpc:"storageGroups"`

	// The order container may not contain the final tax rates when it is returned from [[SoftLayer_Product_Order/verifyOrder|verifyOrder]]. This hash will facilitate checking if the tax rates have finished being calculated and retrieving the accurate tax rate values.
	TaxCacheHash *string `json:"taxCacheHash,omitempty" xmlrpc:"taxCacheHash"`

	// Flag to indicate if the order container has the final tax rates for the order. Some tax rates are calculated in the background because they take longer, and they might not be finished when the container is returned from [[SoftLayer_Product_Order/verifyOrder|verifyOrder]].
	TaxCompletedFlag *bool `json:"taxCompletedFlag,omitempty" xmlrpc:"taxCompletedFlag"`

	// The SoftLayer_Product_Item_Price for the Tech Incubator discount.  The oneTimeFee field contain the calculated discount being applied to the order.
	TechIncubatorItemPrice *Product_Item_Price `json:"techIncubatorItemPrice,omitempty" xmlrpc:"techIncubatorItemPrice"`

	// The total tax portion of the recurring fees.
	TotalRecurringTax *Float64 `json:"totalRecurringTax,omitempty" xmlrpc:"totalRecurringTax"`

	// The tax amount of the setup fees.
	TotalSetupTax *Float64 `json:"totalSetupTax,omitempty" xmlrpc:"totalSetupTax"`

	// An optional flag to use hourly pricing instead of standard monthly pricing.
	UseHourlyPricing *bool `json:"useHourlyPricing,omitempty" xmlrpc:"useHourlyPricing"`

	// For virtual guest (virtual server) orders, this property is required if you did not specify data in the <code>hardware</code> property. This is an array of [[SoftLayer_Virtual_Guest]] objects. The <code>hostname</code> and <code>domain</code> properties are required for each virtual guest object. There is no need to specify data in this property and the <code>hardware</code> property - only one is required for virtual server orders.
	VirtualGuests []Virtual_Guest `json:"virtualGuests,omitempty" xmlrpc:"virtualGuests"`
}

// This datatype is to be used for data transfer requests.
type Container_Product_Order_Account_Media_Data_Transfer_Request struct {
	Container_Product_Order

	// An instance of [[SoftLayer_Account_Media_Data_Transfer_Request]]
	Request *Account_Media_Data_Transfer_Request `json:"request,omitempty" xmlrpc:"request"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. The SoftLayer_Container_Product_Order_Attribute_Address datatype contains the address information.
type Container_Product_Order_Attribute_Address struct {
	Entity

	// The physical street address.
	AddressLine1 *string `json:"addressLine1,omitempty" xmlrpc:"addressLine1"`

	// The second line in the address. Information such as suite number goes here.
	AddressLine2 *string `json:"addressLine2,omitempty" xmlrpc:"addressLine2"`

	// The city name
	City *string `json:"city,omitempty" xmlrpc:"city"`

	// The 2-character Country code. (i.e. US)
	CountryCode *string `json:"countryCode,omitempty" xmlrpc:"countryCode"`

	// The non US/Canadian state or region.
	NonUsState *string `json:"nonUsState,omitempty" xmlrpc:"nonUsState"`

	// The Zip or Postal Code.
	PostalCode *string `json:"postalCode,omitempty" xmlrpc:"postalCode"`

	// The state or region.
	State *string `json:"state,omitempty" xmlrpc:"state"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. The SoftLayer_Container_Product_Order_Attribute_Contact datatype contains the contact information.
type Container_Product_Order_Attribute_Contact struct {
	Entity

	// The address information of the contact.
	Address *Container_Product_Order_Attribute_Address `json:"address,omitempty" xmlrpc:"address"`

	// The email address of the contact.
	EmailAddress *string `json:"emailAddress,omitempty" xmlrpc:"emailAddress"`

	// The fax number associated with a contact. This is an optional value.
	FaxNumber *string `json:"faxNumber,omitempty" xmlrpc:"faxNumber"`

	// The first name of the contact.
	FirstName *string `json:"firstName,omitempty" xmlrpc:"firstName"`

	// The last name of the contact.
	LastName *string `json:"lastName,omitempty" xmlrpc:"lastName"`

	// The organization name of the contact.
	OrganizationName *string `json:"organizationName,omitempty" xmlrpc:"organizationName"`

	// The phone number associated with a contact.
	PhoneNumber *string `json:"phoneNumber,omitempty" xmlrpc:"phoneNumber"`

	// The title of the contact.
	Title *string `json:"title,omitempty" xmlrpc:"title"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. The SoftLayer_Container_Product_Order_Attribute_Organization datatype contains the organization information.
type Container_Product_Order_Attribute_Organization struct {
	Entity

	// The address information of the contact.
	Address *Container_Product_Order_Attribute_Address `json:"address,omitempty" xmlrpc:"address"`

	// The fax number associated with an organization. This is an optional value.
	FaxNumber *string `json:"faxNumber,omitempty" xmlrpc:"faxNumber"`

	// The name of an organization.
	OrganizationName *string `json:"organizationName,omitempty" xmlrpc:"organizationName"`

	// The phone number associated with an organization.
	PhoneNumber *string `json:"phoneNumber,omitempty" xmlrpc:"phoneNumber"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place an order with SoftLayer.
type Container_Product_Order_Billing_Information struct {
	Entity

	// The physical street address. Reserve information such as "apartment #123" or "Suite 2" for line 1.
	BillingAddressLine1 *string `json:"billingAddressLine1,omitempty" xmlrpc:"billingAddressLine1"`

	// The second line in the address. Information such as suite number goes here.
	BillingAddressLine2 *string `json:"billingAddressLine2,omitempty" xmlrpc:"billingAddressLine2"`

	// The city in which a customer's account resides.
	BillingCity *string `json:"billingCity,omitempty" xmlrpc:"billingCity"`

	// The 2-character Country code for an account's address. (i.e. US)
	BillingCountryCode *string `json:"billingCountryCode,omitempty" xmlrpc:"billingCountryCode"`

	// The email address associated with a customer account.
	BillingEmail *string `json:"billingEmail,omitempty" xmlrpc:"billingEmail"`

	// the company name for an account.
	BillingNameCompany *string `json:"billingNameCompany,omitempty" xmlrpc:"billingNameCompany"`

	// The first name of the customer account owner.
	BillingNameFirst *string `json:"billingNameFirst,omitempty" xmlrpc:"billingNameFirst"`

	// The last name of the customer account owner
	BillingNameLast *string `json:"billingNameLast,omitempty" xmlrpc:"billingNameLast"`

	// The fax number associated with a customer account.
	BillingPhoneFax *string `json:"billingPhoneFax,omitempty" xmlrpc:"billingPhoneFax"`

	// The phone number associated with a customer account.
	BillingPhoneVoice *string `json:"billingPhoneVoice,omitempty" xmlrpc:"billingPhoneVoice"`

	// The Zip or Postal Code for the billing address on an account.
	BillingPostalCode *string `json:"billingPostalCode,omitempty" xmlrpc:"billingPostalCode"`

	// The State for the account.
	BillingState *string `json:"billingState,omitempty" xmlrpc:"billingState"`

	// The credit card number to use.
	CardAccountNumber *string `json:"cardAccountNumber,omitempty" xmlrpc:"cardAccountNumber"`

	// The payment card expiration month
	CardExpirationMonth *int `json:"cardExpirationMonth,omitempty" xmlrpc:"cardExpirationMonth"`

	// The payment card expiration year
	CardExpirationYear *int `json:"cardExpirationYear,omitempty" xmlrpc:"cardExpirationYear"`

	// The Card Verification Value Code (CVV) number
	CreditCardVerificationNumber *string `json:"creditCardVerificationNumber,omitempty" xmlrpc:"creditCardVerificationNumber"`

	// Tax exempt status. 1 = exempt (not taxable),  0 = not exempt (taxable)
	TaxExempt *int `json:"taxExempt,omitempty" xmlrpc:"taxExempt"`

	// The VAT ID entered at checkout
	VatId *string `json:"vatId,omitempty" xmlrpc:"vatId"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. The SoftLayer_Container_Product_Order_Dns_Domain_Registration datatype contains everything required to place a domain registration order with SoftLayer.
type Container_Product_Order_Dns_Domain_Registration struct {
	Container_Product_Order

	// Administrative contact information associated with an registraton or transfer. This is required if registration type is 'new' or 'transfer'.
	AdministrativeContact *Container_Dns_Domain_Registration_Contact `json:"administrativeContact,omitempty" xmlrpc:"administrativeContact"`

	// Billing contact information associated with an registraton or transfer. This is required if registration type is 'new' or 'transfer'.
	BillingContact *Container_Dns_Domain_Registration_Contact `json:"billingContact,omitempty" xmlrpc:"billingContact"`

	// The list of domains to be registered. This is required if registration type is 'new', 'renew', or 'transfer'.
	DomainRegistrationList []Container_Dns_Domain_Registration_List `json:"domainRegistrationList,omitempty" xmlrpc:"domainRegistrationList"`

	// Owner contact information associated with an registraton or transfer. This is required if registration type is 'new' or 'transfer'.
	OwnerContact *Container_Dns_Domain_Registration_Contact `json:"ownerContact,omitempty" xmlrpc:"ownerContact"`

	// The type of a domain registration order. The registration type is Required. Allowed values are new, transfer, and renew
	RegistrationType *string `json:"registrationType,omitempty" xmlrpc:"registrationType"`

	// Technical contact information associated with an registraton or transfer. This is required if registration type is 'new' or 'transfer'.
	TechnicalContact *Container_Dns_Domain_Registration_Contact `json:"technicalContact,omitempty" xmlrpc:"technicalContact"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. The SoftLayer_Container_Product_Order_Dns_Domain_Reseller datatype contains everything required to place a domain reseller credit order with SoftLayer.
type Container_Product_Order_Dns_Domain_Reseller struct {
	Container_Product_Order

	// Amount to be credited to the domain reseller account.
	CreditAmount *Float64 `json:"creditAmount,omitempty" xmlrpc:"creditAmount"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place a Gateway Appliance Cluster order with SoftLayer.
type Container_Product_Order_Gateway_Appliance_Cluster struct {
	Container_Product_Order

	// Used to identify which items on an order belong in the same cluster.
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" xmlrpc:"clusterIdentifier"`

	// Indicates what type of cluster order is being placed (HA, Provision).
	ClusterOrderType *string `json:"clusterOrderType,omitempty" xmlrpc:"clusterOrderType"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place a hardware security module order with SoftLayer.
type Container_Product_Order_Hardware_Security_Module struct {
	Container_Product_Order_Hardware_Server
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place an order with SoftLayer.
type Container_Product_Order_Hardware_Server struct {
	Container_Product_Order

	// Used to identify which items on an order belong in the same cluster.
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" xmlrpc:"clusterIdentifier"`

	// Indicates what type of cluster order is being placed (HA, Provision).
	ClusterOrderType *string `json:"clusterOrderType,omitempty" xmlrpc:"clusterOrderType"`

	// Used to identify which gateway is being upgraded to HA.
	ClusterResourceId *int `json:"clusterResourceId,omitempty" xmlrpc:"clusterResourceId"`

	// Id of the [[SoftLayer_Monitoring_Agent_Configuration_Template_Group]] to be used with the monitoring package
	MonitoringAgentConfigurationTemplateGroupId *int `json:"monitoringAgentConfigurationTemplateGroupId,omitempty" xmlrpc:"monitoringAgentConfigurationTemplateGroupId"`

	// When ordering Virtual Server (Private Node), this variable specifies the role of the server configuration. (Deprecated)
	PrivateCloudServerRole *string `json:"privateCloudServerRole,omitempty" xmlrpc:"privateCloudServerRole"`

	// Used to identify which device the new server should be attached to.
	RequiredUpstreamDeviceId *int `json:"requiredUpstreamDeviceId,omitempty" xmlrpc:"requiredUpstreamDeviceId"`

	// tags (used in MongoDB deployments). (Deprecated)
	Tags []Container_Product_Order_Property `json:"tags,omitempty" xmlrpc:"tags"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place an order with SoftLayer.
type Container_Product_Order_Hardware_Server_Colocation struct {
	Container_Product_Order_Hardware_Server
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place a Gateway Appliance order.
type Container_Product_Order_Hardware_Server_Gateway_Appliance struct {
	Container_Product_Order_Hardware_Server
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place an order with SoftLayer.
type Container_Product_Order_Hardware_Server_Upgrade struct {
	Container_Product_Order_Hardware_Server
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place a Monitoring Package order with SoftLayer.
type Container_Product_Order_Monitoring_Package struct {
	Container_Product_Order

	// no documentation yet
	ConfigurationTemplateGroups []Monitoring_Agent_Configuration_Template_Group `json:"configurationTemplateGroups,omitempty" xmlrpc:"configurationTemplateGroups"`

	// no documentation yet
	ServerType *string `json:"serverType,omitempty" xmlrpc:"serverType"`
}

// This is a datatype used with multi-configuration deployments. Multi-configuration deployments also have a deployment specific datatype that should be used in lieu of this one.
type Container_Product_Order_MultiConfiguration struct {
	Container_Product_Order
}

// no documentation yet
type Container_Product_Order_MultiConfiguration_Tornado struct {
	Container_Product_Order_MultiConfiguration
}

// This type contains the structure of network-related objects that may be specified when ordering services.
type Container_Product_Order_Network struct {
	Entity

	// The [[SoftLayer_Network]] object.
	Network *Network `json:"network,omitempty" xmlrpc:"network"`

	// The list of public [[SoftLayer_Container_Product_Order_Network_Vlan|vlans]] available for ordering. Each VLAN will have list of public subnets that are accessible to the VLAN.
	PublicVlans []Container_Product_Order `json:"publicVlans,omitempty" xmlrpc:"publicVlans"`

	// The list of private [[SoftLayer_Container_Product_Order_Network_Subnet|subnets]] available for ordering with a description of their available IP space.
	Subnets []Container_Product_Order `json:"subnets,omitempty" xmlrpc:"subnets"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place an application delivery controller order with SoftLayer.
type Container_Product_Order_Network_Application_Delivery_Controller struct {
	Container_Product_Order

	// An optional [[SoftLayer_Network_Application_Delivery_Controller]] identifier that is used for upgrading an existing application delivery controller.
	ApplicationDeliveryControllerId *int `json:"applicationDeliveryControllerId,omitempty" xmlrpc:"applicationDeliveryControllerId"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place a CDN order with SoftLayer.
type Container_Product_Order_Network_ContentDelivery_Account struct {
	Container_Product_Order

	// The CDN account name
	CdnAccountName *string `json:"cdnAccountName,omitempty" xmlrpc:"cdnAccountName"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place a CDN order with SoftLayer.
type Container_Product_Order_Network_ContentDelivery_Account_Upgrade struct {
	Container_Product_Order

	// ID of an existing CDN account. You can use this to upgrade an existing CDN account.
	CdnAccountId *string `json:"cdnAccountId,omitempty" xmlrpc:"cdnAccountId"`
}

// This is the default container type for network load balancer orders.
type Container_Product_Order_Network_LoadBalancer struct {
	Container_Product_Order
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place a global load balancer order with SoftLayer.
type Container_Product_Order_Network_LoadBalancer_Global struct {
	Container_Product_Order

	// The domain name that will be load balanced.
	Domain *string `json:"domain,omitempty" xmlrpc:"domain"`

	// The hostname that will be load balanced.
	Hostname *string `json:"hostname,omitempty" xmlrpc:"hostname"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place a network message delivery order with SoftLayer.
type Container_Product_Order_Network_Message_Delivery struct {
	Container_Product_Order

	// The account password for SendGrid enrollment.
	AccountPassword *string `json:"accountPassword,omitempty" xmlrpc:"accountPassword"`

	// The username for SendGrid enrollment.
	AccountUsername *string `json:"accountUsername,omitempty" xmlrpc:"accountUsername"`

	// The email address for SendGrid enrollment.
	EmailAddress *string `json:"emailAddress,omitempty" xmlrpc:"emailAddress"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place a Message Queue order with SoftLayer.
type Container_Product_Order_Network_Message_Queue struct {
	Container_Product_Order
}

// This is the base data type for Performance storage order containers. If you wish to place an order you must not use this class and instead use the appropriate child container for the type of storage you would like to order: [[SoftLayer_Container_Product_Order_Network_PerformanceStorage_Nfs]] for File and [[SoftLayer_Container_Product_Order_Network_PerformanceStorage_Iscsi]] for Block storage.
type Container_Product_Order_Network_PerformanceStorage struct {
	Container_Product_Order
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place an order for iSCSI (Block) Performance Storage
type Container_Product_Order_Network_PerformanceStorage_Iscsi struct {
	Container_Product_Order_Network_PerformanceStorage

	// OS Type to be used when formatting the storage space, this should match the OS type that will be connecting to the LUN. The only required property its the keyName of the OS type.
	OsFormatType *Network_Storage_Iscsi_OS_Type `json:"osFormatType,omitempty" xmlrpc:"osFormatType"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place an order for NFS (File) Performance Storage
type Container_Product_Order_Network_PerformanceStorage_Nfs struct {
	Container_Product_Order_Network_PerformanceStorage
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place a hardware firewall order with SoftLayer.
type Container_Product_Order_Network_Protection_Firewall struct {
	Container_Product_Order
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place a hardware (dedicated) firewall order with SoftLayer.
type Container_Product_Order_Network_Protection_Firewall_Dedicated struct {
	Container_Product_Order

	// generic properties.
	Vlan *Network_Vlan `json:"vlan,omitempty" xmlrpc:"vlan"`

	// generic properties.
	VlanId *int `json:"vlanId,omitempty" xmlrpc:"vlanId"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place an order for additional Evault plugins.
type Container_Product_Order_Network_Storage_Backup_Evault_Plugin struct {
	Container_Product_Order
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place an Evault order with SoftLayer.
type Container_Product_Order_Network_Storage_Backup_Evault_Vault struct {
	Container_Product_Order
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place an order for Enterprise Storage
type Container_Product_Order_Network_Storage_Enterprise struct {
	Container_Product_Order

	// This must be populated only for replicant volume ordering. It represents the identifier of the origin [[SoftLayer_Network_Storage]].
	OriginVolumeId *int `json:"originVolumeId,omitempty" xmlrpc:"originVolumeId"`

	// This must be populated only for replicant volume ordering. It represents the [[SoftLayer_Network_Storage_Schedule]] that will be be used to replicate the origin [[SoftLayer_Network_Storage]] volume.
	OriginVolumeScheduleId *int `json:"originVolumeScheduleId,omitempty" xmlrpc:"originVolumeScheduleId"`

	// This must be populated for block storage orders. This should match the OS type of the host(s) that will connect to the volume. The only required property is the keyName of the OS type. This property is ignored for file storage orders.
	OsFormatType *Network_Storage_Iscsi_OS_Type `json:"osFormatType,omitempty" xmlrpc:"osFormatType"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place an order for Enterprise Storage Snapshot Space.
type Container_Product_Order_Network_Storage_Enterprise_SnapshotSpace struct {
	Container_Product_Order

	// The [[SoftLayer_Network_Storage]] id for which snapshot space is being ordered for.
	VolumeId *int `json:"volumeId,omitempty" xmlrpc:"volumeId"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place an upgrade order for Enterprise Storage Snapshot Space.
type Container_Product_Order_Network_Storage_Enterprise_SnapshotSpace_Upgrade struct {
	Container_Product_Order_Network_Storage_Enterprise_SnapshotSpace
}

// This datatype is to be used for object storage orders.
type Container_Product_Order_Network_Storage_Hub struct {
	Container_Product_Order
}

// This class is used to contain a datacenter location and its associated active usage rate prices for object storage ordering.
type Container_Product_Order_Network_Storage_Hub_Datacenter struct {
	Entity

	// The datacenter location where object storage is available.
	Location *Location `json:"location,omitempty" xmlrpc:"location"`

	// The collection of active usage rate item prices.
	UsageRatePrices []Product_Item_Price `json:"usageRatePrices,omitempty" xmlrpc:"usageRatePrices"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place an ISCSI order with SoftLayer.
type Container_Product_Order_Network_Storage_Iscsi struct {
	Container_Product_Order
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place an ISCSI Replication order with SoftLayer.
type Container_Product_Order_Network_Storage_Iscsi_Replication struct {
	Container_Product_Order

	// the [[SoftLayer_Network_Storage_Iscsi_EqualLogic_Version3]] Id.
	VolumeId *int `json:"volumeId,omitempty" xmlrpc:"volumeId"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place an ISCSI Snapshot Space order with SoftLayer.
type Container_Product_Order_Network_Storage_Iscsi_SnapshotSpace struct {
	Container_Product_Order

	// the [[SoftLayer_Network_Storage_Iscsi_EqualLogic_Version3]] Id.
	VolumeId *int `json:"volumeId,omitempty" xmlrpc:"volumeId"`
}

// The SoftLayer_Container_Product_Order_Network_Storage_Modification datatype has everything required to place a modification to an existing StorageLayer account with SoftLayer. Modifications, at present time, include upgrade and downgrades only. The ''volumeId'' property must be set to the network storage volume id to be upgraded. Once populated send this container to the [[SoftLayer_Product_Order::placeOrder]] method.
//
// The ''packageId'' property passed in for CloudLayer storage accounts must be set to 0 (zero) and the ''quantity'' property must be set to 1. The location does not have to be set. Please use the [[SoftLayer_Product_Package]] service to retrieve a list of CloudLayer items.
//
// NOTE: When upgrading CloudLayer storage service from a metered plan (pay as you go) to a non-metered plan, make sure the chosen plan's storage allotment has enough space to cover the current usage. If the chosen plan's usage allotment is less than the CloudLayer storage's usage the order will be rejected.
type Container_Product_Order_Network_Storage_Modification struct {
	Container_Product_Order

	// The id of the StorageLayer account to modify.
	VolumeId *int `json:"volumeId,omitempty" xmlrpc:"volumeId"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder when placing network attached storage orders.
type Container_Product_Order_Network_Storage_Nas struct {
	Container_Product_Order
}

// This datatype is to be used for ordering object storage products using the object_storage [[SoftLayer_Product_Item_Category|category]]. For object storage products using hub [[SoftLayer_Product_Item_Category|category]] use the [[SoftLayer_Container_Product_Order_Network_Storage_Hub]] order container.
type Container_Product_Order_Network_Storage_Object struct {
	Container_Product_Order
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place a subnet order with SoftLayer.
type Container_Product_Order_Network_Subnet struct {
	Container_Product_Order

	// The description which includes the network identifier, Classless Inter-Domain Routing prefix and the available slot count.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// The [[SoftLayer_Network_Subnet_IpAddress]] id.
	EndPointIpAddressId *int `json:"endPointIpAddressId,omitempty" xmlrpc:"endPointIpAddressId"`

	// The [[SoftLayer_Network_Vlan]] id.
	EndPointVlanId *int `json:"endPointVlanId,omitempty" xmlrpc:"endPointVlanId"`

	// The [[SoftLayer_Network_Subnet]] id.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// This is the hostname for the router associated with the [[SoftLayer_Network_Subnet|subnet]]. This is a readonly property.
	RouterHostname *string `json:"routerHostname,omitempty" xmlrpc:"routerHostname"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place a network ipsec vpn order with SoftLayer.
type Container_Product_Order_Network_Tunnel_Ipsec struct {
	Container_Product_Order
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place a network vlan order with SoftLayer.
type Container_Product_Order_Network_Vlan struct {
	Container_Product_Order

	// The description which includes the primary router's hostname plus the vlan number.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// The datacenter portion of the hostname.
	HostnameDatacenter *string `json:"hostnameDatacenter,omitempty" xmlrpc:"hostnameDatacenter"`

	// The router portion of the hostname.
	HostnameRouter *string `json:"hostnameRouter,omitempty" xmlrpc:"hostnameRouter"`

	// The [[SoftLayer_Network_Vlan]] id.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// The optional name for this VLAN
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// The router object on which the new VLAN should be created.
	Router *Hardware `json:"router,omitempty" xmlrpc:"router"`

	// The ID of the [[SoftLayer_Hardware_Router]] object on which the new VLAN should be created.
	RouterId *int `json:"routerId,omitempty" xmlrpc:"routerId"`

	// The collection of subnets associated with this vlan.
	Subnets []Container_Product_Order `json:"subnets,omitempty" xmlrpc:"subnets"`

	// The vlan number.
	VlanNumber *int `json:"vlanNumber,omitempty" xmlrpc:"vlanNumber"`
}

// This class contains the collections of public and private VLANs that are available during the ordering process.
type Container_Product_Order_Network_Vlans struct {
	Entity

	// The collection of private vlans available during ordering.
	PrivateVlans []Container_Product_Order `json:"privateVlans,omitempty" xmlrpc:"privateVlans"`

	// The collection of public vlans available during ordering.
	PublicVlans []Container_Product_Order `json:"publicVlans,omitempty" xmlrpc:"publicVlans"`
}

// This is used for storing various items about the order. Currently used for storing additional raid information when ordering servers. This is optional
type Container_Product_Order_Property struct {
	Entity

	// The property name
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// The property value
	Value *string `json:"value,omitempty" xmlrpc:"value"`
}

// When an order is placed (SoftLayer_Product_Order::placeOrder), a receipt is returned when the order is created successfully. The information in the receipt helps explain information about the order. It's order ID, and all the data within the order as well.
//
// For PayPal Orders, an URL is also returned to the user so that the user can complete the transaction. Users paying with PayPal must continue on to this URL, login and pay. When doing this, PayPal will redirect the user back to a SoftLayer page which will then "finalize" the authorization process. From here, Sales will verify the order by contacting the user in some way, unless sales has already spoken to the user about approving the order.
//
// For users paying with a credit card, a receipt means the order has gone to sales and is awaiting approval.
type Container_Product_Order_Receipt struct {
	Entity

	// This URL refers to the location where you will visit to complete the payment authorization for an external service, such as PayPal. This property is associated with <code>externalPaymentToken</code> and will only be populated when purchasing products with an external service.
	//
	// Once you visit this location, you will be presented with the options to confirm payment or deny payment. If you confirm payment, you will be redirected back to the receipt for your order. If you deny, you will be redirected back to the cancel order page where you do not need to take any additional action.
	//
	// Until you confirm payment with the external service, your products will not be provisioned or accessible for your consumption. Upon successfully confirming payment, our system will be notified and the order approval and provisioning systems will begin processing. After provisioning is complete, your services will be available.
	ExternalPaymentCheckoutUrl *string `json:"externalPaymentCheckoutUrl,omitempty" xmlrpc:"externalPaymentCheckoutUrl"`

	// This token refers to the identifier for the external payment authorization. This token is associated with the <code>externalPaymentCheckoutUrl</code> and is only populated when purchasing products with an external service like PayPal.
	ExternalPaymentToken *string `json:"externalPaymentToken,omitempty" xmlrpc:"externalPaymentToken"`

	// The date when SoftLayer received the order.
	OrderDate *Time `json:"orderDate,omitempty" xmlrpc:"orderDate"`

	// This is a copy of the order container (SoftLayer_Container_Product_Order) which holds all the data related to an order. This will only return when an order is processed successfully. It will contain all the items in an order as well as the order totals.
	OrderDetails *Container_Product_Order `json:"orderDetails,omitempty" xmlrpc:"orderDetails"`

	// SoftLayer's unique identifier for the order.
	OrderId *int `json:"orderId,omitempty" xmlrpc:"orderId"`

	// Deprecation notice: use <code>externalPaymentCheckoutUrl</code> instead of this property.
	//
	// This URL refers to the location where you will visit to complete the payment authorization for PayPal. This property is associated with <code>paypalToken</code> and will only be populated when purchasing products with PayPal.
	//
	// Once you visit PayPal's site, you will be presented with the options to confirm payment or deny payment. If you confirm payment, you will be redirected back to the receipt for your order. If you deny, you will be redirected back to the cancel order page where you do not need to take any additional action.
	//
	// Until you confirm payment with PayPal, your products will not be provisioned or accessible for your consumption. Upon successfully confirming payment, our system will be notified and the order approval and provisioning systems will begin processing. After provisioning is complete, your services will be available.
	PaypalCheckoutUrl *string `json:"paypalCheckoutUrl,omitempty" xmlrpc:"paypalCheckoutUrl"`

	// Deprecation notice: use <code>externalPaymentToken</code> instead of this property.
	//
	// This token refers to the identifier provided when payment is processed via PayPal. This token is associated with the <code>paypalCheckoutUrl</code>.
	PaypalToken *string `json:"paypalToken,omitempty" xmlrpc:"paypalToken"`

	// This is a copy of the order that was successfully placed (SoftLayer_Billing_Order). This will only return when an order is processed successfully.
	PlacedOrder *Billing_Order `json:"placedOrder,omitempty" xmlrpc:"placedOrder"`

	// This is a copy of the quote container (SoftLayer_Billing_Order_Quote) which holds all the data related to a quote. This will only return when a quote is processed successfully.
	Quote *Billing_Order_Quote `json:"quote,omitempty" xmlrpc:"quote"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype contains everything required to place a secure certificate order with SoftLayer.
type Container_Product_Order_Security_Certificate struct {
	Container_Product_Order

	// The administrator contact associated with a SSL certificate. If the contact is not provided the technical contact will be used. If the address is not provided the organization information address will be used.
	AdministrativeContact *Container_Product_Order_Attribute_Contact `json:"administrativeContact,omitempty" xmlrpc:"administrativeContact"`

	// The billing contact associated with a SSL certificate. If the contact is not provided the technical contact will be used. If the address is not provided the organization information address will be used.
	BillingContact *Container_Product_Order_Attribute_Contact `json:"billingContact,omitempty" xmlrpc:"billingContact"`

	// The base64 encoded string that sent from an applicant to a certificate authority. The CSR contains information identifying the applicant and the public key chosen by the applicant. The corresponding private key should not be included.
	CertificateSigningRequest *string `json:"certificateSigningRequest,omitempty" xmlrpc:"certificateSigningRequest"`

	// The email address that can approve a secure certificate order.
	OrderApproverEmailAddress *string `json:"orderApproverEmailAddress,omitempty" xmlrpc:"orderApproverEmailAddress"`

	// The organization information associated with a SSL certificate.
	OrganizationInformation *Container_Product_Order_Attribute_Organization `json:"organizationInformation,omitempty" xmlrpc:"organizationInformation"`

	// Indicates if it is an renewal order of an existing SSL certificate.
	RenewalFlag *bool `json:"renewalFlag,omitempty" xmlrpc:"renewalFlag"`

	// The number of servers.
	ServerCount *int `json:"serverCount,omitempty" xmlrpc:"serverCount"`

	// The server type. This is the name from a [[SoftLayer_Security_Certificate_Request_ServerType]] object.
	ServerType *string `json:"serverType,omitempty" xmlrpc:"serverType"`

	// The technical contact associated with a SSL certificate. If the address is not provided the organization information address will be used.
	TechnicalContact *Container_Product_Order_Attribute_Contact `json:"technicalContact,omitempty" xmlrpc:"technicalContact"`

	// The period that a SSL certificate is valid for.  For example, 12, 24, 36
	ValidityMonths *int `json:"validityMonths,omitempty" xmlrpc:"validityMonths"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder.
type Container_Product_Order_Service struct {
	Container_Product_Order
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place a virtual license order with SoftLayer.
type Container_Product_Order_Software_Component_Virtual struct {
	Container_Product_Order

	// array of ip address ids for which a license should be allocated for.
	EndPointIpAddressIds []int `json:"endPointIpAddressIds,omitempty" xmlrpc:"endPointIpAddressIds"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place a hardware security module order with SoftLayer.
type Container_Product_Order_Software_License struct {
	Container_Product_Order
}

// This object holds all of the ssh key ids that will allow authentication to a single server.
type Container_Product_Order_SshKeys struct {
	Entity

	// An array of SoftLayer_Security_Ssh_Key IDs to assign to a server.
	SshKeyIds []int `json:"sshKeyIds,omitempty" xmlrpc:"sshKeyIds"`
}

// A single storage group container used for a hardware server order.
//
// This object describes a single storage group that can be added to an order container.
type Container_Product_Order_Storage_Group struct {
	Entity

	// Size of the array in gigabytes. Must be within limitations of the smallest drive assigned to the storage group and the storage group type.
	ArraySize *Float64 `json:"arraySize,omitempty" xmlrpc:"arraySize"`

	// The array type id from a [[SoftLayer_Configuration_Storage_Group_Array_Type]] object.
	ArrayTypeId *int `json:"arrayTypeId,omitempty" xmlrpc:"arrayTypeId"`

	// Integer array of drive indexes to use in the storage group.
	HardDrives []int `json:"hardDrives,omitempty" xmlrpc:"hardDrives"`

	// If an array should be protected by an hotspare, the drive index of the hotspares should be here.
	//
	// If a drive is a hotspare for all arrays then a separate storage group with array type GLOBAL_HOT_SPARE should be used
	HotSpareDrives []int `json:"hotSpareDrives,omitempty" xmlrpc:"hotSpareDrives"`

	// The id for a [[SoftLayer_Hardware_Component_Partition_Template]] object, which will determine the partitions to add to the storage group.
	//
	// If this storage group is not a primary storage group, then this will not be used.
	PartitionTemplateId *int `json:"partitionTemplateId,omitempty" xmlrpc:"partitionTemplateId"`

	// Defines the partitions for the storage group.
	//
	// If this storage group is not a secondary storage group, then this will not be used.
	Partitions []Container_Product_Order_Storage_Group_Partition `json:"partitions,omitempty" xmlrpc:"partitions"`
}

// A storage group partition container used for a hardware server order.
//
// This object describes the partitions for a single storage group that can be added to an order container.
type Container_Product_Order_Storage_Group_Partition struct {
	Entity

	// Is this a grow partition
	IsGrow *bool `json:"isGrow,omitempty" xmlrpc:"isGrow"`

	// The name of this partition
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// The size of this partition
	Size *Float64 `json:"size,omitempty" xmlrpc:"size"`
}

// This container type is used for placing orders for external authentication, such as phone-based authentication.
type Container_Product_Order_User_Customer_External_Binding struct {
	Container_Product_Order

	// The external id that access to external authentication is being purchased for.
	ExternalId *string `json:"externalId,omitempty" xmlrpc:"externalId"`

	// The SoftLayer [[SoftLayer_User_Customer|user]] identifier that an external binding is being purchased for.
	UserId *int `json:"userId,omitempty" xmlrpc:"userId"`

	// The [[SoftLayer_User_Customer_External_Binding_Vendor|vendor]] identifier for the external binding being purchased.
	VendorId *int `json:"vendorId,omitempty" xmlrpc:"vendorId"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place a Portable Storage order with SoftLayer.
type Container_Product_Order_Virtual_Disk_Image struct {
	Container_Product_Order

	// Label for the portable storage volume.
	DiskDescription *string `json:"diskDescription,omitempty" xmlrpc:"diskDescription"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place an order with SoftLayer.
type Container_Product_Order_Virtual_Guest struct {
	Container_Product_Order_Hardware_Server

	// Identifier of the [[SoftLayer_Virtual_Disk_Image]] to boot from.
	BootableDiskId *int `json:"bootableDiskId,omitempty" xmlrpc:"bootableDiskId"`
}

// This is the datatype that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to place an order with SoftLayer.
type Container_Product_Order_Virtual_Guest_Upgrade struct {
	Container_Product_Order_Virtual_Guest
}

// This is the datatype that needs to be populated and sent to SoftLayer_Provisioning_Maintenance_Window::addCustomerUpgradeWindow. This datatype has everything required to place an order with SoftLayer.
type Container_Provisioning_Maintenance_Window struct {
	Entity

	// Maintenance classifications.
	ClassificationIds []Provisioning_Maintenance_Classification `json:"classificationIds,omitempty" xmlrpc:"classificationIds"`

	// Maintenance classifications.
	ItemCategoryIds []Product_Item_Category `json:"itemCategoryIds,omitempty" xmlrpc:"itemCategoryIds"`

	// The maintenance window id
	MaintenanceWindowId *int `json:"maintenanceWindowId,omitempty" xmlrpc:"maintenanceWindowId"`

	// Maintenance window ticket id
	TicketId *int `json:"ticketId,omitempty" xmlrpc:"ticketId"`

	// Maintenance window date
	WindowMaintenanceDate *Time `json:"windowMaintenanceDate,omitempty" xmlrpc:"windowMaintenanceDate"`
}

// no documentation yet
type Container_Referral_Partner_Commission struct {
	Entity

	// no documentation yet
	CommissionAmount *Float64 `json:"commissionAmount,omitempty" xmlrpc:"commissionAmount"`

	// no documentation yet
	CommissionRate *Float64 `json:"commissionRate,omitempty" xmlrpc:"commissionRate"`

	// no documentation yet
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// no documentation yet
	ReferralAccountId *int `json:"referralAccountId,omitempty" xmlrpc:"referralAccountId"`

	// no documentation yet
	ReferralCompanyName *string `json:"referralCompanyName,omitempty" xmlrpc:"referralCompanyName"`

	// no documentation yet
	ReferralPartnerAccountId *int `json:"referralPartnerAccountId,omitempty" xmlrpc:"referralPartnerAccountId"`

	// no documentation yet
	ReferralRevenue *Float64 `json:"referralRevenue,omitempty" xmlrpc:"referralRevenue"`
}

// no documentation yet
type Container_Referral_Partner_Payment_Option struct {
	Entity

	// no documentation yet
	AccountNumber *string `json:"accountNumber,omitempty" xmlrpc:"accountNumber"`

	// no documentation yet
	AccountType *string `json:"accountType,omitempty" xmlrpc:"accountType"`

	// no documentation yet
	Address1 *string `json:"address1,omitempty" xmlrpc:"address1"`

	// no documentation yet
	Address2 *string `json:"address2,omitempty" xmlrpc:"address2"`

	// no documentation yet
	BankTransitNumber *string `json:"bankTransitNumber,omitempty" xmlrpc:"bankTransitNumber"`

	// no documentation yet
	City *string `json:"city,omitempty" xmlrpc:"city"`

	// no documentation yet
	CompanyName *string `json:"companyName,omitempty" xmlrpc:"companyName"`

	// no documentation yet
	Country *string `json:"country,omitempty" xmlrpc:"country"`

	// no documentation yet
	FederalTaxId *string `json:"federalTaxId,omitempty" xmlrpc:"federalTaxId"`

	// no documentation yet
	FirstName *string `json:"firstName,omitempty" xmlrpc:"firstName"`

	// no documentation yet
	LastName *string `json:"lastName,omitempty" xmlrpc:"lastName"`

	// no documentation yet
	PaymentType *string `json:"paymentType,omitempty" xmlrpc:"paymentType"`

	// no documentation yet
	PaypalEmail *string `json:"paypalEmail,omitempty" xmlrpc:"paypalEmail"`

	// no documentation yet
	PhoneNumber *string `json:"phoneNumber,omitempty" xmlrpc:"phoneNumber"`

	// no documentation yet
	PostalCode *string `json:"postalCode,omitempty" xmlrpc:"postalCode"`

	// no documentation yet
	State *string `json:"state,omitempty" xmlrpc:"state"`
}

// no documentation yet
type Container_Referral_Partner_Prospect struct {
	Entity

	// no documentation yet
	Address1 *string `json:"address1,omitempty" xmlrpc:"address1"`

	// no documentation yet
	Address2 *string `json:"address2,omitempty" xmlrpc:"address2"`

	// no documentation yet
	City *string `json:"city,omitempty" xmlrpc:"city"`

	// no documentation yet
	CompanyName *string `json:"companyName,omitempty" xmlrpc:"companyName"`

	// no documentation yet
	Country *string `json:"country,omitempty" xmlrpc:"country"`

	// no documentation yet
	Email *string `json:"email,omitempty" xmlrpc:"email"`

	// no documentation yet
	FirstName *string `json:"firstName,omitempty" xmlrpc:"firstName"`

	// no documentation yet
	LastName *string `json:"lastName,omitempty" xmlrpc:"lastName"`

	// no documentation yet
	OfficePhone *string `json:"officePhone,omitempty" xmlrpc:"officePhone"`

	// no documentation yet
	PostalCode *string `json:"postalCode,omitempty" xmlrpc:"postalCode"`

	// no documentation yet
	Questions []string `json:"questions,omitempty" xmlrpc:"questions"`

	// no documentation yet
	Responses []Survey_Response `json:"responses,omitempty" xmlrpc:"responses"`

	// no documentation yet
	State *string `json:"state,omitempty" xmlrpc:"state"`

	// no documentation yet
	SurveyId *string `json:"surveyId,omitempty" xmlrpc:"surveyId"`
}

// The SoftLayer_Container_RemoteManagement_Graphs_SensorSpeed contains graphs to  display speed for each of the server's fans.  Fan speeds are gathered from the server's remote management card.
type Container_RemoteManagement_Graphs_SensorSpeed struct {
	Entity

	// The graph to display the server's fan speed.
	Graph *[]byte `json:"graph,omitempty" xmlrpc:"graph"`

	// A title that may be used to display for the graph.
	Title *string `json:"title,omitempty" xmlrpc:"title"`
}

// The SoftLayer_Container_RemoteManagement_Graphs_SensorTemperature contains graphs to display the cpu(s) and system temperatures retrieved from the management card using thermometer graphs.
type Container_RemoteManagement_Graphs_SensorTemperature struct {
	Entity

	// The graph to display the server's cpu(s) and system temperatures.
	Graph *[]byte `json:"graph,omitempty" xmlrpc:"graph"`

	// A title that may be used to display for the graph.
	Title *string `json:"title,omitempty" xmlrpc:"title"`
}

// The SoftLayer_Container_RemoteManagement_PmInfo contains pminfo information retrieved from a server's remote management card.
type Container_RemoteManagement_PmInfo struct {
	Entity

	// PmInfo ID
	PmInfoId *string `json:"pmInfoId,omitempty" xmlrpc:"pmInfoId"`

	// PmInfo Reading
	PmInfoReading *string `json:"pmInfoReading,omitempty" xmlrpc:"pmInfoReading"`
}

// The SoftLayer_Container_RemoteManagement_SensorReadings contains sensor information retrieved from a server's remote management card.
type Container_RemoteManagement_SensorReading struct {
	Entity

	// Lower Non-Recoverable threshold
	LowerCritical *string `json:"lowerCritical,omitempty" xmlrpc:"lowerCritical"`

	// Lower Non-Critical threshold
	LowerNonCritical *string `json:"lowerNonCritical,omitempty" xmlrpc:"lowerNonCritical"`

	// Lower Non-Recoverable threshold
	LowerNonRecoverable *string `json:"lowerNonRecoverable,omitempty" xmlrpc:"lowerNonRecoverable"`

	// Sensor ID
	SensorId *string `json:"sensorId,omitempty" xmlrpc:"sensorId"`

	// Sensor Reading
	SensorReading *string `json:"sensorReading,omitempty" xmlrpc:"sensorReading"`

	// Sensor Units
	SensorUnits *string `json:"sensorUnits,omitempty" xmlrpc:"sensorUnits"`

	// Sensor Status
	Status *string `json:"status,omitempty" xmlrpc:"status"`

	// Upper Critical threshold
	UpperCritical *string `json:"upperCritical,omitempty" xmlrpc:"upperCritical"`

	// Upper Non-Critical threshold
	UpperNonCritical *string `json:"upperNonCritical,omitempty" xmlrpc:"upperNonCritical"`

	// Upper Non-Recoverable threshold
	UpperNonRecoverable *string `json:"upperNonRecoverable,omitempty" xmlrpc:"upperNonRecoverable"`
}

// The SoftLayer_Container_RemoteManagement_SensorReadingsWithGraphs contains the raw data retrieved from a server's remote management card.  Along with the raw data, two sets of graphs will be returned.  One set of graphs is used to display, using thermometer graphs, the temperatures (cpu(s) and system) retrieved from the management card.  The other set is used to display speed for each of the server's fans.
type Container_RemoteManagement_SensorReadingsWithGraphs struct {
	Entity

	// The raw data returned from the server's remote management card.
	RawData []Container_RemoteManagement_SensorReading `json:"rawData,omitempty" xmlrpc:"rawData"`

	// The graph(s) to display the server's fan speeds.
	SpeedGraphs []Container_RemoteManagement_Graphs_SensorSpeed `json:"speedGraphs,omitempty" xmlrpc:"speedGraphs"`

	// The graph(s) to display the server's cpu(s) and system temperatures.
	TemperatureGraphs []Container_RemoteManagement_Graphs_SensorTemperature `json:"temperatureGraphs,omitempty" xmlrpc:"temperatureGraphs"`
}

// The metadata service resource container is used to store information about a single service resource.
type Container_Resource_Metadata_ServiceResource struct {
	Entity

	// The backend IP address for this resource
	BackendIpAddress *string `json:"backendIpAddress,omitempty" xmlrpc:"backendIpAddress"`

	// The type for this resource
	Type *Network_Service_Resource_Type `json:"type,omitempty" xmlrpc:"type"`
}

// This data type is a container that stores information about a single indexed object type.  Object type information can be used for discovery of searchable data and for creation or validation of object index search strings.  Each of these containers holds a collection of <b>[[SoftLayer_Container_Search_ObjectType_Property (type)|SoftLayer_Container_Search_ObjectType_Property]]</b> objects, specifying which object properties are exposed for the current user.  Refer to the the documentation for the <b>[[SoftLayer_Search/search|search()]]</b> method for information on using object types in search strings.
type Container_Search_ObjectType struct {
	Entity

	// Name of object type.
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// A collection of [[SoftLayer_Container_Search_ObjectType_Property|object properties]].
	Properties []Container_Search_ObjectType_Property `json:"properties,omitempty" xmlrpc:"properties"`
}

// This data type is a container that stores information about a single property of a searchable object type.  Each <b>[[SoftLayer_Container_Search_ObjectType (type)|SoftLayer_Container_Search_ObjectType]]</b> object holds a collection of these properties.  Property information can be used for discovery of searchable data and for the creation or validation of object index search strings.  Note that properties are only understood by the <b>[[SoftLayer_Search/advancedSearch|advancedSearch()]]</b> method.  Refer to the <b>advancedSearch()</b> method for information on using properties in search strings.
type Container_Search_ObjectType_Property struct {
	Entity

	// Name of property.
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// Indicates if this property can be sorted.
	SortableFlag *bool `json:"sortableFlag,omitempty" xmlrpc:"sortableFlag"`

	// Property data type.  Valid values include 'boolean', 'integer', 'date', 'string' or 'text'.
	Type *string `json:"type,omitempty" xmlrpc:"type"`
}

// The SoftLayer_Container_Search_Result data type represents a result row from an execution of Search service.
type Container_Search_Result struct {
	Entity

	// An array of terms that were matched in the resource object.
	MatchedTerms []string `json:"matchedTerms,omitempty" xmlrpc:"matchedTerms"`

	// The score ratio of the result for relevance to the search criteria.
	RelevanceScore *Float64 `json:"relevanceScore,omitempty" xmlrpc:"relevanceScore"`

	// A search results resource object that matched search criteria.
	Resource *Entity `json:"resource,omitempty" xmlrpc:"resource"`

	// The type of the resource object that matched search criteria.
	ResourceType *string `json:"resourceType,omitempty" xmlrpc:"resourceType"`
}

// The SoftLayer_Container_Software_Component_HostIps_Policy container holds the title and value of a current host ips policy.
type Container_Software_Component_HostIps_Policy struct {
	Entity

	// The value of a host ips category.
	Policy *string `json:"policy,omitempty" xmlrpc:"policy"`

	// The category title of a host ips policy.
	PolicyTitle *string `json:"policyTitle,omitempty" xmlrpc:"policyTitle"`
}

// These are the results of a tax calculation. The tax calculation was kicked off but allowed to run in the background. This type stores the results so that an interface can be updated with up-to-date information.
type Container_Tax_Cache struct {
	Entity

	// The percentage of the final total that should be tax.
	EffectiveTaxRate *Float64 `json:"effectiveTaxRate,omitempty" xmlrpc:"effectiveTaxRate"`

	// The container that holds the four actual tax rates, one for each fee type.
	Items []Container_Tax_Cache_Item `json:"items,omitempty" xmlrpc:"items"`

	// The status of the tax request. This should be PENDING, FAILED, or COMPLETED.
	Status *string `json:"status,omitempty" xmlrpc:"status"`

	// The final amount of tax for the order.
	TotalTaxAmount *Float64 `json:"totalTaxAmount,omitempty" xmlrpc:"totalTaxAmount"`
}

// This represents one order item in a tax calculation.
type Container_Tax_Cache_Item struct {
	Entity

	// The category code for the referenced product.
	CategoryCode *string `json:"categoryCode,omitempty" xmlrpc:"categoryCode"`

	// This hash will match to the hash on an order container.
	ContainerHash *string `json:"containerHash,omitempty" xmlrpc:"containerHash"`

	// The reference to the price for this order item.
	ItemPriceId *int `json:"itemPriceId,omitempty" xmlrpc:"itemPriceId"`

	// This is the container containing the individual tax rates.
	TaxRates *Container_Tax_Rates `json:"taxRates,omitempty" xmlrpc:"taxRates"`
}

// This contains the four tax rates, one for each fee type.
type Container_Tax_Rates struct {
	Entity

	// The tax rate associated with the labor fee.
	LaborTaxRate *Float64 `json:"laborTaxRate,omitempty" xmlrpc:"laborTaxRate"`

	// A reference to a location.
	LocationId *Float64 `json:"locationId,omitempty" xmlrpc:"locationId"`

	// The tax rate associated with the one-time fee.
	OneTimeTaxRate *Float64 `json:"oneTimeTaxRate,omitempty" xmlrpc:"oneTimeTaxRate"`

	// The tax rate associated with the recurring fee.
	RecurringTaxRate *Float64 `json:"recurringTaxRate,omitempty" xmlrpc:"recurringTaxRate"`

	// The tax rate associated with the setup fee.
	SetupTaxRate *Float64 `json:"setupTaxRate,omitempty" xmlrpc:"setupTaxRate"`
}

// SoftLayer_Container_Ticket_GraphInputs models a single inbound object for a given ticket graph.
type Container_Ticket_GraphInputs struct {
	Entity

	// This is a unix timestamp that represents the stop date/time for a graph.
	EndDate *Time `json:"endDate,omitempty" xmlrpc:"endDate"`

	// The front-end or back-end network uplink interface associated with this server.
	NetworkInterfaceId *int `json:"networkInterfaceId,omitempty" xmlrpc:"networkInterfaceId"`

	// *
	Pod *int `json:"pod,omitempty" xmlrpc:"pod"`

	// This is a human readable name for the server or rack being graphed.
	ServerName *string `json:"serverName,omitempty" xmlrpc:"serverName"`

	// This is a unix timestamp that represents the begin date/time for a graph.
	StartDate *Time `json:"startDate,omitempty" xmlrpc:"startDate"`
}

// SoftLayer_Container_Ticket_GraphOutputs models a single outbound object for a given bandwidth graph.
type Container_Ticket_GraphOutputs struct {
	Entity

	// The raw PNG binary data to be displayed once the graph is drawn.
	GraphImage *[]byte `json:"graphImage,omitempty" xmlrpc:"graphImage"`

	// The title that ended up being displayed as part of the graph image.
	GraphTitle *string `json:"graphTitle,omitempty" xmlrpc:"graphTitle"`

	// The maximum date included in this graph.
	MaxEndDate *Time `json:"maxEndDate,omitempty" xmlrpc:"maxEndDate"`

	// The minimum date included in this graph.
	MinStartDate *Time `json:"minStartDate,omitempty" xmlrpc:"minStartDate"`
}

// no documentation yet
type Container_Ticket_Priority struct {
	Entity

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// no documentation yet
	Value *int `json:"value,omitempty" xmlrpc:"value"`
}

// no documentation yet
type Container_Ticket_Survey_Preference struct {
	Entity

	// no documentation yet
	Applicable *bool `json:"applicable,omitempty" xmlrpc:"applicable"`

	// no documentation yet
	OptedOut *bool `json:"optedOut,omitempty" xmlrpc:"optedOut"`

	// no documentation yet
	OptedOutDate *Time `json:"optedOutDate,omitempty" xmlrpc:"optedOutDate"`
}

// Container class used to hold user authentication token
type Container_User_Authentication_Token struct {
	Entity

	// hash that gets populated for user authentication
	Hash *string `json:"hash,omitempty" xmlrpc:"hash"`

	// the user authenticated object
	User *User_Customer `json:"user,omitempty" xmlrpc:"user"`

	// the id of the user to authenticate
	UserId *int `json:"userId,omitempty" xmlrpc:"userId"`
}

// Container classed used to hold external authentication information
type Container_User_Customer_External_Binding struct {
	Entity

	// The unique token that is created by an external authentication request.
	AuthenticationToken *string `json:"authenticationToken,omitempty" xmlrpc:"authenticationToken"`

	// The OpenID Connect access token which provides access to a resource by the OpenID Connect provider.
	OpenIdConnectAccessToken *string `json:"openIdConnectAccessToken,omitempty" xmlrpc:"openIdConnectAccessToken"`

	// The account to login to, if not provided a default will be used.
	OpenIdConnectAccountId *int `json:"openIdConnectAccountId,omitempty" xmlrpc:"openIdConnectAccountId"`

	// The OpenID Connect provider type, as a string.
	OpenIdConnectProvider *string `json:"openIdConnectProvider,omitempty" xmlrpc:"openIdConnectProvider"`

	// Your SoftLayer customer portal user's portal password.
	Password *string `json:"password,omitempty" xmlrpc:"password"`

	// The answer to your security question.
	SecurityQuestionAnswer *string `json:"securityQuestionAnswer,omitempty" xmlrpc:"securityQuestionAnswer"`

	// A security question you wish to answer when authenticating to the SoftLayer customer portal. This parameter isn't required if no security questions are set on your portal account or if your account is configured to not require answering a security account upon login.
	SecurityQuestionId *int `json:"securityQuestionId,omitempty" xmlrpc:"securityQuestionId"`

	// The username you wish to authenticate to the SoftLayer customer portal with.
	Username *string `json:"username,omitempty" xmlrpc:"username"`

	// The name of the vendor that will be used for external authentication
	Vendor *string `json:"vendor,omitempty" xmlrpc:"vendor"`
}

// Container classed used to hold portal token
type Container_User_Customer_External_Binding_Phone struct {
	Container_User_Customer_External_Binding
}

// This container can be used to configure the phone authentication mode. By default, "VOICE_CALL" in "STANDARD" mode with no Pin number will be used. With the default mode, you will have to answer a phone call from a trusted 2 form factor vendor during authentication process. You have to answer the call and follow the instruction in order to complete the authentication.
//
// You can also use SMS text message or PhoneFactor mobile app modes (in case you're using PhoneFactor). Additionally, you can set up a Pin number. By requiring you to verify your secret PIN, you can ensure that you have possession of your phone.
type Container_User_Customer_External_Binding_Phone_Mode struct {
	Entity

	// Authentication mode. Valid modes are: VOICE_CALL, SMS_TEXT, PHONE_APP
	//
	//
	// *VOICE_CALL
	// In this mode, users will receive a phone call to authenticate. Using PIN can enhance the security of the phone authentication by requiring the user to enter a PIN during the authentication call. Valid Pin modes are: PIN, VOICE_PRINT, STANDARD
	//
	//
	// **STANDARD: (default) No PIN is used.
	// **PIN: 4 to 10 digit numeric value
	// **VOICE_PRINT: The user's voice will be used to identify the user.
	//
	//
	// *SMS_TEXT
	// SMS Text mode will send a SMS text message to the user's phone to complete the authentication.  There are 2 different pin modes:
	//
	//
	// **OTP: (default) A text message containing a One-Time Passcode (OTP) is sent to the user. The user must reply to the text message entering this OTP to complete the authentication.
	// **OTP_PIN: This mode enhances the security of the authentication by requiring the user to enter the OTP + their PIN in the text reply.
	//
	//
	//
	//
	// *PHONE_APP
	// This mode is applicable for PhoneFactor. Phone App mode results in a notification being sent to the user's PhoneFactor phone app. There are 2 different pin modes for the mobile app authentication.
	// **STANDARD: (default) The first authentication is when the user signs on using a username and password.
	// The second authentication is when the user receives a notification in the PhoneFactor phone app. In Standard Mode, users will prompted to authenticate, deny, or deny and report fraud.
	// **PIN: This mode enhances the security of the authentication by requiring the user to enter their PIN in the phone app.
	Mode *string `json:"mode,omitempty" xmlrpc:"mode"`

	// Optional authentication pin.
	Pin *string `json:"pin,omitempty" xmlrpc:"pin"`

	// Available Pin modes are: PIN, VOICE_PRINT, STANDARD Default: STANDARD (Pin is not used)
	PinMode *string `json:"pinMode,omitempty" xmlrpc:"pinMode"`
}

// Container classed used to hold portal token
type Container_User_Customer_External_Binding_Totp struct {
	Container_User_Customer_External_Binding

	// The security code used to validate a Totp credential.
	SecurityCode *string `json:"securityCode,omitempty" xmlrpc:"securityCode"`
}

// Container classed used to hold details about an external authentication vendor.
type Container_User_Customer_External_Binding_Vendor struct {
	Entity

	// The keyname used to identify an external authentication vendor.
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// The name of an external authentication vendor.
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// Container classed used to hold portal token
type Container_User_Customer_External_Binding_Verisign struct {
	Container_User_Customer_External_Binding

	// A second security code that is only required if your credential has become unsynchronized.
	SecondSecurityCode *string `json:"secondSecurityCode,omitempty" xmlrpc:"secondSecurityCode"`

	// The security code used to validate a VeriSign credential.
	SecurityCode *string `json:"securityCode,omitempty" xmlrpc:"securityCode"`
}

// Container for holding information necessary for the setting and resetting of customer passwords
//
//
type Container_User_Customer_PasswordSet struct {
	Entity

	// id of SoftLayer_User_Security_Question
	AnsweredSecurityQuestionId *int `json:"answeredSecurityQuestionId,omitempty" xmlrpc:"answeredSecurityQuestionId"`

	// the authentication methods required
	AuthenticationMethods []int `json:"authenticationMethods,omitempty" xmlrpc:"authenticationMethods"`

	// the password key provided to user in the password set url link sent via email
	Key *string `json:"key,omitempty" xmlrpc:"key"`

	// the user's new password
	Password *string `json:"password,omitempty" xmlrpc:"password"`

	// answer to security question provided by the user
	SecurityAnswer *string `json:"securityAnswer,omitempty" xmlrpc:"securityAnswer"`

	// array of SoftLayer_User_Security_Question
	SecurityQuestions []User_Security_Question `json:"securityQuestions,omitempty" xmlrpc:"securityQuestions"`

	// the id of the user to authenticate
	UserId *int `json:"userId,omitempty" xmlrpc:"userId"`
}

// Container classed used to hold mobile portal token
type Container_User_Customer_Portal_MobileToken struct {
	Container_User_Customer_Portal_Token

	// True if this user login required an external binding.
	HasExternalBinding *bool `json:"hasExternalBinding,omitempty" xmlrpc:"hasExternalBinding"`
}

// Container classed used to hold portal token
type Container_User_Customer_Portal_Token struct {
	Entity

	// hash of logged in user session id
	Hash *string `json:"hash,omitempty" xmlrpc:"hash"`

	// the logged in user data
	User *User_Customer `json:"user,omitempty" xmlrpc:"user"`

	// the id of the logged in user
	UserId *int `json:"userId,omitempty" xmlrpc:"userId"`
}

// This container holds user's phone information.
type Container_User_Data_Phone struct {
	Entity

	// Country code number for the phone number Default: 1 (United States & Canada +1)
	CountryCode *int `json:"countryCode,omitempty" xmlrpc:"countryCode"`

	// Phone extension code. It can be digits, commas, *, and # are allowed.
	Extension *string `json:"extension,omitempty" xmlrpc:"extension"`

	// Phone number can be a mobile phone number, desk phone number, or some other option. The phone number format must match the format selected in the country code.
	Phone *string `json:"phone,omitempty" xmlrpc:"phone"`

	// Type of phone number such as "primary", "office" or "home"
	PhoneType *string `json:"phoneType,omitempty" xmlrpc:"phoneType"`
}

// Container classed used to hold portal token
type Container_User_Employee_External_Binding_Verisign struct {
	Entity
}

// At times,such as when attaching files to tickets, it is necessary to send files to SoftLayer API methods. The SoftLayer_Container_Utility_File_Attachment data type models a single file to upload to the API.
type Container_Utility_File_Attachment struct {
	Entity

	// The contents of a file that is uploaded to the SoftLayer API.
	Data *[]byte `json:"data,omitempty" xmlrpc:"data"`

	// The name of a file that is uploaded to the SoftLayer API.
	Filename *string `json:"filename,omitempty" xmlrpc:"filename"`
}

// Used to describe a document in the file system on the file server
type Container_Utility_File_Descriptor struct {
	Entity

	// The name of a file as it exists on the file server.
	FileName *string `json:"fileName,omitempty" xmlrpc:"fileName"`

	// The friendly name of a file as it exists on the file server.
	FriendlyName *string `json:"friendlyName,omitempty" xmlrpc:"friendlyName"`

	// The date the file was last modified on the file server.
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`
}

// SoftLayer_Container_Utility_File_Entity data type models a single entity on a storage resource. Entities can include anything within a storage volume including files, folders, directories, and CloudLayer storage projects.
type Container_Utility_File_Entity struct {
	Entity

	// A file entity's raw content.
	Content *[]byte `json:"content,omitempty" xmlrpc:"content"`

	// A file entity's MIME content type.
	ContentType *string `json:"contentType,omitempty" xmlrpc:"contentType"`

	// The date a file entity was created.
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// The date a CloudLayer storage file entity was moved into the recycle bin. This field applies to files that are pending deletion in the recycle bin.
	DeleteDate *Time `json:"deleteDate,omitempty" xmlrpc:"deleteDate"`

	// Unique identifier for the file. This can be either a number or guid.
	Id *string `json:"id,omitempty" xmlrpc:"id"`

	// Whether a CloudLayer storage file entity is shared with another CloudLayer user.
	IsShared *int `json:"isShared,omitempty" xmlrpc:"isShared"`

	// The date a file entity was last changed.
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// A file entity's name.
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// The owner is usually the account who first upload or created the file on the resource or the account who is responsible for the file at the moment.
	Owner *string `json:"owner,omitempty" xmlrpc:"owner"`

	// The size of a file entity in bytes.
	Size *uint `json:"size,omitempty" xmlrpc:"size"`

	// A CloudLayer storage file entity's type. Types can include "file", "folder", "dir", and "project".
	Type *string `json:"type,omitempty" xmlrpc:"type"`

	// The latest revision of a file on a CloudLayer storage volume. This number increments each time a new revision of the file is uploaded.
	Version *int `json:"version,omitempty" xmlrpc:"version"`
}

// no documentation yet
type Container_Utility_Message struct {
	Entity

	// no documentation yet
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// no documentation yet
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	Message *string `json:"message,omitempty" xmlrpc:"message"`

	// no documentation yet
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// no documentation yet
	Summary *string `json:"summary,omitempty" xmlrpc:"summary"`
}

// SoftLayer customer servers that are purchased with the Microsoft Windows operating system are configured by default to retrieve updates from SoftLayer's local Windows Server Update Services (WSUS) server. Periodically, these servers synchronize and check for new updates from their local WSUS server. SoftLayer_Container_Utility_Microsoft_Windows_UpdateServices_Status models the results of a server's last synchronization attempt as queried from SoftLayer's WSUS servers.
type Container_Utility_Microsoft_Windows_UpdateServices_Status struct {
	Entity

	// The last time a server rebooted due to a Windows Update.
	LastRebootDate *Time `json:"lastRebootDate,omitempty" xmlrpc:"lastRebootDate"`

	// The last time that SoftLayer's local WSUS server received a status update from a customer server.
	LastStatusDate *Time `json:"lastStatusDate,omitempty" xmlrpc:"lastStatusDate"`

	// The last time a server synchronized with SoftLayer's local WSUS server.
	LastSyncDate *Time `json:"lastSyncDate,omitempty" xmlrpc:"lastSyncDate"`

	// This is the private IP address for this server.
	PrivateIPAddress *string `json:"privateIPAddress,omitempty" xmlrpc:"privateIPAddress"`

	// The status message returned from a server's last synchronization with SoftLayer's local WSUS server.
	SyncStatus *string `json:"syncStatus,omitempty" xmlrpc:"syncStatus"`

	// A server's update status, as retrieved form SoftLayer's local WSUS server.
	UpdateStatus *string `json:"updateStatus,omitempty" xmlrpc:"updateStatus"`
}

// SoftLayer_Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem models a single Microsoft Update as reported by SoftLayer's private Windows Server Update Services (WSUS) services. All servers purchased with Microsoft Windows retrieve updates from SoftLayer's WSUS servers by default.
type Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem struct {
	Entity

	// A short description of a Microsoft Windows Update.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// Flag indicating that this patch failed to properly install
	Failed *bool `json:"failed,omitempty" xmlrpc:"failed"`

	// A Windows Update's knowledge base article number. Every Windows Update can be referenced on the Microsoft Help and Support site at the URL <nowiki>http://support.microsoft.com/kb/<article number></nowiki>.
	KbArticleNumber *int `json:"kbArticleNumber,omitempty" xmlrpc:"kbArticleNumber"`

	// Flag indicating that the update is entirely optionals
	Optional *bool `json:"optional,omitempty" xmlrpc:"optional"`

	// Flag indicating that a reboot is needed for this update to be fully applied
	RequiresReboot *bool `json:"requiresReboot,omitempty" xmlrpc:"requiresReboot"`
}

// The SoftLayer_Container_Utility_Network_Firewall_Rule_Attribute data type contains information relating to a single firewall rule.
type Container_Utility_Network_Firewall_Rule_Attribute struct {
	Entity

	// The valid actions for use with rules.
	Actions []string `json:"actions,omitempty" xmlrpc:"actions"`

	// Maximum allowed number of rules.
	MaximumRuleCount *int `json:"maximumRuleCount,omitempty" xmlrpc:"maximumRuleCount"`

	// The valid protocols for use with rules.
	Protocols []string `json:"protocols,omitempty" xmlrpc:"protocols"`

	// The valid source ip subnet masks for use with rules.
	SourceIpSubnetMasks []Container_Utility_Network_Subnet_Mask_Generic_Detail `json:"sourceIpSubnetMasks,omitempty" xmlrpc:"sourceIpSubnetMasks"`
}

// The SoftLayer_Container_Utility_Network_Subnet_Mask_Generic_Detail data type contains information relating to a subnet mask and details associated with that object.
type Container_Utility_Network_Subnet_Mask_Generic_Detail struct {
	Entity

	// The subnet cidr prefix.
	Cidr *string `json:"cidr,omitempty" xmlrpc:"cidr"`

	// The subnet mask description.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// The subnet mask.
	Mask *string `json:"mask,omitempty" xmlrpc:"mask"`
}

// The SoftLayer_Container_Virtual_Guest_Block_Device_Template_Configuration data type contains information relating to a template's external location for importing and exporting
type Container_Virtual_Guest_Block_Device_Template_Configuration struct {
	Entity

	// The group name to be applied to the imported template
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// The note to be applied to the imported template
	Note *string `json:"note,omitempty" xmlrpc:"note"`

	//
	// The referenceCode of the operating system software description for the imported VHD
	OperatingSystemReferenceCode *string `json:"operatingSystemReferenceCode,omitempty" xmlrpc:"operatingSystemReferenceCode"`

	//
	// The URI for an object storage object (.vhd/.iso file)
	// <code>swift://<ObjectStorageAccountName>@<clusterName>/<containerName>/<fileName.(vhd|iso)></code>
	Uri *string `json:"uri,omitempty" xmlrpc:"uri"`
}

// The guest configuration container is used to provide configuration options for creating computing instances.
//
// Each configuration option will include both an <code>itemPrice</code> and a <code>template</code>.
//
// The <code>itemPrice</code> value will provide hourly and monthly costs (if either are applicable), and a description of the option.
//
// The <code>template</code> will provide a fragment of the request with the properties and values that must be sent when creating a computing instance with the option.
//
// The [[SoftLayer_Virtual_Guest/getCreateObjectOptions|getCreateObjectOptions]] method returns this data structure.
//
// <style type="text/css">#properties .views-field-body p { margin-top: 1.5em; };</style>
type Container_Virtual_Guest_Configuration struct {
	Entity

	//
	// <div style="width: 200%">
	// Available block device options.
	//
	//
	// A computing instance will have at least one block device represented by a <code>device</code> number of <code>'0'</code>.
	//
	//
	// The <code>blockDevices.device</code> value in the template represents which device the option is for.
	// The <code>blockDevices.diskImage.capacity</code> value in the template represents the size, in gigabytes, of the disk.
	// The <code>localDiskFlag</code> value in the template represents whether the option is a local or SAN based disk.
	//
	//
	// Note: The block device number <code>'1'</code> is reserved for the SWAP disk attached to the computing instance.
	// </div>
	BlockDevices []Container_Virtual_Guest_Configuration_Option `json:"blockDevices,omitempty" xmlrpc:"blockDevices"`

	//
	// <div style="width: 200%">
	// Available datacenter options.
	//
	//
	// The <code>datacenter.name</code> value in the template represents which datacenter the computing instance will be provisioned in.
	// </div>
	Datacenters []Container_Virtual_Guest_Configuration_Option `json:"datacenters,omitempty" xmlrpc:"datacenters"`

	//
	// <div style="width: 200%">
	// Available memory options.
	//
	//
	// The <code>maxMemory</code> value in the template represents the amount of memory, in megabytes, allocated to the computing instance.
	// </div>
	Memory []Container_Virtual_Guest_Configuration_Option `json:"memory,omitempty" xmlrpc:"memory"`

	//
	// <div style="width: 200%">
	// Available network component options.
	//
	//
	// The <code>networkComponent.maxSpeed</code> value in the template represents the link speed, in megabits per second, of the network connections for a computing instance.
	// </div>
	NetworkComponents []Container_Virtual_Guest_Configuration_Option `json:"networkComponents,omitempty" xmlrpc:"networkComponents"`

	//
	// <div style="width: 200%">
	// Available operating system options.
	//
	//
	// The <code>operatingSystemReferenceCode</code> value in the template is an identifier for a particular operating system. When provided exactly as shown in the template, that operating system will be used.
	//
	//
	// A reference code is structured as three tokens separated by underscores. The first token represents the product, the second is the version of the product, and the third is whether the OS is 32 or 64bit.
	//
	//
	// When providing an <code>operatingSystemReferenceCode</code> while ordering a computing instance the only token required to match exactly is the product. The version token may be given as 'LATEST', else it will require an exact match as well. When the bits token is not provided, 64 bits will be assumed.
	//
	//
	// Providing the value of 'LATEST' for a version will select the latest release of that product for the operating system. As this may change over time, you should be sure that the release version is irrelevant for your applications.
	//
	//
	// For Windows based operating systems the version will represent both the release version (2008, 2012, etc) and the edition (Standard, Enterprise, etc). For all other operating systems the version will represent the major version (Centos 6, Ubuntu 12, etc) of that operating system, minor versions are not represented in a reference code.
	//
	//
	// <b>Notice</b> - Some operating systems are charged based on the value specified in <code>startCpus</code>. The price which is used can be determined by calling [[SoftLayer_Virtual_Guest/generateOrderTemplate|generateOrderTemplate]] with your desired device specifications.
	// </div>
	OperatingSystems []Container_Virtual_Guest_Configuration_Option `json:"operatingSystems,omitempty" xmlrpc:"operatingSystems"`

	//
	// <div style="width: 200%">
	// Available processor options.
	//
	//
	// The <code>startCpus</code> value in the template represents the number of cores allocated to the computing instance.
	// The <code>dedicatedAccountHostOnlyFlag</code> value in the template represents whether the instance will run on hosts with instances belonging to other accounts.
	// </div>
	Processors []Container_Virtual_Guest_Configuration_Option `json:"processors,omitempty" xmlrpc:"processors"`
}

// An option found within a [[SoftLayer_Container_Virtual_Guest_Configuration (type)]] structure.
type Container_Virtual_Guest_Configuration_Option struct {
	Entity

	//
	// Provides hourly and monthly costs (if either are applicable), and a description of the option.
	ItemPrice *Product_Item_Price `json:"itemPrice,omitempty" xmlrpc:"itemPrice"`

	//
	// Provides a fragment of the request with the properties and values that must be sent when creating a computing instance with the option.
	Template *Virtual_Guest `json:"template,omitempty" xmlrpc:"template"`
}
