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

// no documentation yet
type Billing_Currency struct {
	Entity

	// no documentation yet
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// no documentation yet
type Billing_Currency_ExchangeRate struct {
	Entity

	// no documentation yet
	EffectiveDate *Time `json:"effectiveDate,omitempty" xmlrpc:"effectiveDate"`

	// no documentation yet
	ExpirationDate *Time `json:"expirationDate,omitempty" xmlrpc:"expirationDate"`

	// no documentation yet
	FundingCurrency *Billing_Currency `json:"fundingCurrency,omitempty" xmlrpc:"fundingCurrency"`

	// The id of the exchange rate record.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	LocalCurrency *Billing_Currency `json:"localCurrency,omitempty" xmlrpc:"localCurrency"`

	// no documentation yet
	Rate *Float64 `json:"rate,omitempty" xmlrpc:"rate"`
}

// Every SoftLayer customer account has billing specific information which is kept in the SoftLayer_Billing_Info data type. This information is used by the SoftLayer accounting group when sending invoices and making billing inquiries.
type Billing_Info struct {
	Entity

	// The SoftLayer customer account associated with this billing information.
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// A SoftLayer account's identifier.
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// no documentation yet
	AchInformation []Billing_Info_Ach `json:"achInformation,omitempty" xmlrpc:"achInformation"`

	// A count of
	AchInformationCount *uint `json:"achInformationCount,omitempty" xmlrpc:"achInformationCount"`

	// The day of the month that a SoftLayer customer is billed.
	AnniversaryDayOfMonth *int `json:"anniversaryDayOfMonth,omitempty" xmlrpc:"anniversaryDayOfMonth"`

	// This value doesn't persist to this object. It's used as part of the account creation process only;
	CardAccountNumber *string `json:"cardAccountNumber,omitempty" xmlrpc:"cardAccountNumber"`

	// the expiration month of the credit card on file
	CardExpirationMonth *int `json:"cardExpirationMonth,omitempty" xmlrpc:"cardExpirationMonth"`

	// the expiration year of the credit card on file
	CardExpirationYear *int `json:"cardExpirationYear,omitempty" xmlrpc:"cardExpirationYear"`

	// no documentation yet
	CardNickname *string `json:"cardNickname,omitempty" xmlrpc:"cardNickname"`

	// the type of the credit card on file
	CardType *string `json:"cardType,omitempty" xmlrpc:"cardType"`

	// This value doesn't persist to this object. It's used as part of the account creation process only.
	CardVerificationNumber *string `json:"cardVerificationNumber,omitempty" xmlrpc:"cardVerificationNumber"`

	// The date a customer's billing information was created.
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// Currency to be used by this customer account.
	Currency *Billing_Currency `json:"currency,omitempty" xmlrpc:"currency"`

	// Information related to an account's current and previous billing cycles.
	CurrentBillingCycle *Billing_Info_Cycle `json:"currentBillingCycle,omitempty" xmlrpc:"currentBillingCycle"`

	// A SoftLayer customer's billing information identifier.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// The date on which an account was last billed.
	LastBillDate *Time `json:"lastBillDate,omitempty" xmlrpc:"lastBillDate"`

	// The last four digits of the credit card currently on the account. This is the only portion of the card that we store. For Paypal customers, this value will be empty.
	LastFourPaymentCardDigits *int `json:"lastFourPaymentCardDigits,omitempty" xmlrpc:"lastFourPaymentCardDigits"`

	// The date of the last payment received by SoftLayer from the account holder.
	LastPaymentDate *Time `json:"lastPaymentDate,omitempty" xmlrpc:"lastPaymentDate"`

	// The date a customer's billing information was last modified.
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// The date on which an account will be billed next.
	NextBillDate *Time `json:"nextBillDate,omitempty" xmlrpc:"nextBillDate"`

	// The payment terms for an account.
	PaymentTerms *int `json:"paymentTerms,omitempty" xmlrpc:"paymentTerms"`

	// The percentage discount received on all one-time charges on a customer's monthly bill.
	PercentDiscountOnetime *int `json:"percentDiscountOnetime,omitempty" xmlrpc:"percentDiscountOnetime"`

	// The percentage discount received on all recurring charges on a customer's monthly bill.
	PercentDiscountRecurring *int `json:"percentDiscountRecurring,omitempty" xmlrpc:"percentDiscountRecurring"`

	// The total recurring fee amount for servers that are in the spare pool status.
	SparePoolAmount *int `json:"sparePoolAmount,omitempty" xmlrpc:"sparePoolAmount"`

	// no documentation yet
	VatId *string `json:"vatId,omitempty" xmlrpc:"vatId"`
}

// no documentation yet
type Billing_Info_Ach struct {
	Entity

	// no documentation yet
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// no documentation yet
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

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
	FirstName *string `json:"firstName,omitempty" xmlrpc:"firstName"`

	// no documentation yet
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	LastName *string `json:"lastName,omitempty" xmlrpc:"lastName"`

	// no documentation yet
	PhoneNumber *string `json:"phoneNumber,omitempty" xmlrpc:"phoneNumber"`

	// no documentation yet
	Postalcode *string `json:"postalcode,omitempty" xmlrpc:"postalcode"`

	// no documentation yet
	State *string `json:"state,omitempty" xmlrpc:"state"`

	// no documentation yet
	Status *string `json:"status,omitempty" xmlrpc:"status"`

	// no documentation yet
	Street1 *string `json:"street1,omitempty" xmlrpc:"street1"`

	// no documentation yet
	Street2 *string `json:"street2,omitempty" xmlrpc:"street2"`

	// no documentation yet
	VerifiedDate *Time `json:"verifiedDate,omitempty" xmlrpc:"verifiedDate"`
}

// The SoftLayer_Billing_Info_Cycle data type models basic information concerning a SoftLayer account's previous and current billing cycles. The information in this class is only populated for SoftLayer customers who are billed monthly.
type Billing_Info_Cycle struct {
	Entity

	// The account that a current billing cycle is associated with.
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// The ending date of an account's current billing cycle.
	CurrentCycleEndDate *Time `json:"currentCycleEndDate,omitempty" xmlrpc:"currentCycleEndDate"`

	// The starting date of an account's current billing cycle.
	CurrentCycleStartDate *Time `json:"currentCycleStartDate,omitempty" xmlrpc:"currentCycleStartDate"`

	// The start date of an account's next billing cycle.
	NextCycleStartDate *Time `json:"nextCycleStartDate,omitempty" xmlrpc:"nextCycleStartDate"`

	// The ending date of an account's previous billing cycle.
	PreviousCycleEndDate *Time `json:"previousCycleEndDate,omitempty" xmlrpc:"previousCycleEndDate"`

	// The starting date of an account's previous billing cycle.
	PreviousCycleStartDate *Time `json:"previousCycleStartDate,omitempty" xmlrpc:"previousCycleStartDate"`
}

// The SoftLayer_Billing_Invoice data type contains general information relating to an individual invoice applied to a SoftLayer customer account. Personal information in this type such as names, addresses, and phone numbers are taken from the account's contact information at the time the invoice is generated.
type Billing_Invoice struct {
	Entity

	// The account that an invoice belongs to.
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// The SoftLayer customer account that an invoice belongs to.
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// The first line of an address belonging to an account at the time an invoice is created.
	Address1 *string `json:"address1,omitempty" xmlrpc:"address1"`

	// The second line of an address belonging to an account at the time an invoice is created.
	Address2 *string `json:"address2,omitempty" xmlrpc:"address2"`

	// This is the amount of this invoice.
	Amount *Float64 `json:"amount,omitempty" xmlrpc:"amount"`

	// no documentation yet
	BrandAtInvoiceCreation *Brand `json:"brandAtInvoiceCreation,omitempty" xmlrpc:"brandAtInvoiceCreation"`

	// The city portion of an address belonging to an account at the time an invoice is created.
	City *string `json:"city,omitempty" xmlrpc:"city"`

	// Whether an account was exempt from taxes on their invoices at the time an invoice is created.
	ClaimedTaxExemptTxFlag *bool `json:"claimedTaxExemptTxFlag,omitempty" xmlrpc:"claimedTaxExemptTxFlag"`

	// The date an invoice was closed. Open invoices have a null closed date.
	ClosedDate *Time `json:"closedDate,omitempty" xmlrpc:"closedDate"`

	// The company name belonging to an account at the time an invoice is created.
	CompanyName *string `json:"companyName,omitempty" xmlrpc:"companyName"`

	// A two-letter abbreviation of the country portion of an address belonging to an account at the time an invoice is created.
	Country *string `json:"country,omitempty" xmlrpc:"country"`

	// The date an invoice was created.
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// A flag that will reflect whether the detailed version of the pdf has been generated.
	DetailedPdfGeneratedFlag *bool `json:"detailedPdfGeneratedFlag,omitempty" xmlrpc:"detailedPdfGeneratedFlag"`

	// no documentation yet
	DocumentsGeneratedFlag *bool `json:"documentsGeneratedFlag,omitempty" xmlrpc:"documentsGeneratedFlag"`

	// The email address belonging to an account at the time an invoice is created.
	Email *string `json:"email,omitempty" xmlrpc:"email"`

	// An SoftLayer account's balance at the time an invoice is closed. This value is measured in US Dollar ($USD) currency.
	EndingBalance *Float64 `json:"endingBalance,omitempty" xmlrpc:"endingBalance"`

	// The fax telephone number belonging to an account at the time an invoice is created.
	FaxPhone *string `json:"faxPhone,omitempty" xmlrpc:"faxPhone"`

	// The first name of the account holder at the time an invoice is created.
	FirstName *string `json:"firstName,omitempty" xmlrpc:"firstName"`

	// An invoice's internal identifier.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// A count of a list of top-level invoice items that are on the currently pending invoice.
	InvoiceTopLevelItemCount *uint `json:"invoiceTopLevelItemCount,omitempty" xmlrpc:"invoiceTopLevelItemCount"`

	// A list of top-level invoice items that are on the currently pending invoice.
	InvoiceTopLevelItems []Billing_Invoice_Item `json:"invoiceTopLevelItems,omitempty" xmlrpc:"invoiceTopLevelItems"`

	// The total amount of this invoice.
	InvoiceTotalAmount *Float64 `json:"invoiceTotalAmount,omitempty" xmlrpc:"invoiceTotalAmount"`

	// The total one-time charges for this invoice. This is the sum of one-time charges + setup fees + labor fees. This does not include taxes.
	InvoiceTotalOneTimeAmount *Float64 `json:"invoiceTotalOneTimeAmount,omitempty" xmlrpc:"invoiceTotalOneTimeAmount"`

	// A sum of all the taxes related to one time charges for this invoice.
	InvoiceTotalOneTimeTaxAmount *Float64 `json:"invoiceTotalOneTimeTaxAmount,omitempty" xmlrpc:"invoiceTotalOneTimeTaxAmount"`

	// The total amount of this invoice. This does not include taxes.
	InvoiceTotalPreTaxAmount *Float64 `json:"invoiceTotalPreTaxAmount,omitempty" xmlrpc:"invoiceTotalPreTaxAmount"`

	// The total Recurring amount of this invoice. This amount does not include taxes or one time charges.
	InvoiceTotalRecurringAmount *Float64 `json:"invoiceTotalRecurringAmount,omitempty" xmlrpc:"invoiceTotalRecurringAmount"`

	// The total amount of the recurring taxes on this invoice.
	InvoiceTotalRecurringTaxAmount *Float64 `json:"invoiceTotalRecurringTaxAmount,omitempty" xmlrpc:"invoiceTotalRecurringTaxAmount"`

	// A count of the items that belong to this invoice.
	ItemCount *uint `json:"itemCount,omitempty" xmlrpc:"itemCount"`

	// The items that belong to this invoice.
	Items []Billing_Invoice_Item `json:"items,omitempty" xmlrpc:"items"`

	// The last name of the account holder at the time an invoice is created.
	LastName *string `json:"lastName,omitempty" xmlrpc:"lastName"`

	// The date an invoice was last modified.
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// The telephone number belonging to an account at the time an invoice is created.
	OfficePhone *string `json:"officePhone,omitempty" xmlrpc:"officePhone"`

	// This is the total payment made on this invoice.
	Payment *Float64 `json:"payment,omitempty" xmlrpc:"payment"`

	// A count of the payments for the invoice.
	PaymentCount *uint `json:"paymentCount,omitempty" xmlrpc:"paymentCount"`

	// The payments for the invoice.
	Payments []Billing_Invoice_Receivable_Payment `json:"payments,omitempty" xmlrpc:"payments"`

	// The postal code portion of an address belonging to an account at the time an invoice is created.
	PostalCode *string `json:"postalCode,omitempty" xmlrpc:"postalCode"`

	// no documentation yet
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty" xmlrpc:"purchaseOrderNumber"`

	// This is the seller's tax registration.
	SellerRegistration *string `json:"sellerRegistration,omitempty" xmlrpc:"sellerRegistration"`

	// An SoftLayer account's balance at the time an invoice is created. This value is measured in US Dollar ($USD) currency.
	StartingBalance *Float64 `json:"startingBalance,omitempty" xmlrpc:"startingBalance"`

	// A two-letter abbreviation of the state portion of an address belonging to an account at the time an invoice is created. If the account that the invoice was generated for resides outside a province then this is set to "other".
	State *string `json:"state,omitempty" xmlrpc:"state"`

	// An invoice's status. The "OPEN" status means SoftLayer has not yet received payment for this invoice. "CLOSED" status means that SoftLayer has received payment and closed the invoice. The "CLOSED_FAILED" status code means SoftLayer closed the invoice without receiving a payment. Invoices are usually set to CLOSED_FAILED status in cases where customer accounts are terminated for non-payment.
	StatusCode *string `json:"statusCode,omitempty" xmlrpc:"statusCode"`

	// This is the tax information that applies to tax auditing. This is the official tax record for this invoice.
	TaxInfo *Billing_Invoice_Tax_Info `json:"taxInfo,omitempty" xmlrpc:"taxInfo"`

	// This is the set of tax information for any tax calculation for this invoice. Note that not all of these are necessarily official, so use the taxInfo key to get the final information.
	TaxInfoHistory []Billing_Invoice_Tax_Info `json:"taxInfoHistory,omitempty" xmlrpc:"taxInfoHistory"`

	// A count of this is the set of tax information for any tax calculation for this invoice. Note that not all of these are necessarily official, so use the taxInfo key to get the final information.
	TaxInfoHistoryCount *uint `json:"taxInfoHistoryCount,omitempty" xmlrpc:"taxInfoHistoryCount"`

	// This is a message explaining the tax treatment for this invoice.
	TaxMessage *string `json:"taxMessage,omitempty" xmlrpc:"taxMessage"`

	// no documentation yet
	TaxStatusId *int `json:"taxStatusId,omitempty" xmlrpc:"taxStatusId"`

	// This is the strategy used to calculate tax on this invoice.
	TaxType *Billing_Invoice_Tax_Type `json:"taxType,omitempty" xmlrpc:"taxType"`

	// no documentation yet
	TaxTypeId *int `json:"taxTypeId,omitempty" xmlrpc:"taxTypeId"`

	// An invoice's type. SoftLayer invoices and service credits are differentiated by their type. The "NEW" type code signifies an invoice for new service. A SoftLayer customer's first invoice has the NEW type code. "RECURRING" invoices are generated on a SoftLayer customer's anniversary billing date for monthly services. "ONE-TIME-CHARGE" invoices are generated when one-time charges are applied to an account. "CREDIT" invoices are generated whenever SoftLayer applies a credit against an account's balance. There are two special types of service credits. "REFUND" type credits are applied against a customer's account balance along with the receivables on their account. "MANUAL_PAYMENT_CREDIT" invoice credits are generated whenever a customer makes an unscheduled payment.
	TypeCode *string `json:"typeCode,omitempty" xmlrpc:"typeCode"`
}

// Each billing invoice item makes up a record within an invoice. This provides you with a detailed record of everything related to an invoice item. When you are billed, our system takes active billing items and creates an invoice. These invoice items are a copy of your active billing items, and make up the contents of your invoice.
type Billing_Invoice_Item struct {
	Entity

	// An Invoice Item's associated child invoice items. Only parent invoice items have associated children. For instance, a server invoice item may have associated children.
	AssociatedChildren []Billing_Invoice_Item `json:"associatedChildren,omitempty" xmlrpc:"associatedChildren"`

	// A count of an Invoice Item's associated child invoice items. Only parent invoice items have associated children. For instance, a server invoice item may have associated children.
	AssociatedChildrenCount *uint `json:"associatedChildrenCount,omitempty" xmlrpc:"associatedChildrenCount"`

	// An Invoice Item's associated invoice item. If this is populated, it means this is an orphaned invoice item, but logically belongs to the associated invoice item.
	AssociatedInvoiceItem *Billing_Invoice_Item `json:"associatedInvoiceItem,omitempty" xmlrpc:"associatedInvoiceItem"`

	// The associated invoice Item ID.
	AssociatedInvoiceItemId *int `json:"associatedInvoiceItemId,omitempty" xmlrpc:"associatedInvoiceItemId"`

	// An Invoice Item's billing item, from which this item was generated.
	BillingItem *Billing_Item `json:"billingItem,omitempty" xmlrpc:"billingItem"`

	// The billing item from which this invoice item was generated.
	BillingItemId *int `json:"billingItemId,omitempty" xmlrpc:"billingItemId"`

	// This invoice item's "item category".
	Category *Product_Item_Category `json:"category,omitempty" xmlrpc:"category"`

	// The item category of the invoice item being invoiced.
	CategoryCode *string `json:"categoryCode,omitempty" xmlrpc:"categoryCode"`

	// An Invoice Item's child invoice items. Only parent invoice items have children. For instance, a server invoice item will have children.
	Children []Billing_Invoice_Item `json:"children,omitempty" xmlrpc:"children"`

	// A count of an Invoice Item's child invoice items. Only parent invoice items have children. For instance, a server invoice item will have children.
	ChildrenCount *uint `json:"childrenCount,omitempty" xmlrpc:"childrenCount"`

	// The date the invoice item was created.
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// The item description for this invoice item.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// The domain name of the invoiced item. This is only used on invoice items whose category is "server".
	DomainName *string `json:"domainName,omitempty" xmlrpc:"domainName"`

	// An Invoice Item's associated child invoice items, excluding some items with a $0.00 recurring fee. Only parent invoice items have associated children. For instance, a server invoice item may have associated children.
	FilteredAssociatedChildren []Billing_Invoice_Item `json:"filteredAssociatedChildren,omitempty" xmlrpc:"filteredAssociatedChildren"`

	// A count of an Invoice Item's associated child invoice items, excluding some items with a $0.00 recurring fee. Only parent invoice items have associated children. For instance, a server invoice item may have associated children.
	FilteredAssociatedChildrenCount *uint `json:"filteredAssociatedChildrenCount,omitempty" xmlrpc:"filteredAssociatedChildrenCount"`

	// The Host name of the invoiced item. This is only used on invoice items whose category is "server".
	HostName *string `json:"hostName,omitempty" xmlrpc:"hostName"`

	// The hourly recurring fee of the invoice item represented by a floating point decimal in US Dollars ($USD)
	HourlyRecurringFee *Float64 `json:"hourlyRecurringFee,omitempty" xmlrpc:"hourlyRecurringFee"`

	// The ID of the invoice item.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// The invoice to which this item belongs.
	Invoice *Billing_Invoice `json:"invoice,omitempty" xmlrpc:"invoice"`

	// The invoice to which this invoice item belongs.
	InvoiceId *int `json:"invoiceId,omitempty" xmlrpc:"invoiceId"`

	// An invoice item's labor fee total after taxes. This does not include any child invoice items.
	LaborAfterTaxAmount *Float64 `json:"laborAfterTaxAmount,omitempty" xmlrpc:"laborAfterTaxAmount"`

	// This also a one-time fee of a special type.
	LaborFee *Float64 `json:"laborFee,omitempty" xmlrpc:"laborFee"`

	// The tax rate at which the labor fee is taxed.
	LaborFeeTaxRate *Float64 `json:"laborFeeTaxRate,omitempty" xmlrpc:"laborFeeTaxRate"`

	// An invoice item's labor tax amount. This does not include any child invoice items.
	LaborTaxAmount *Float64 `json:"laborTaxAmount,omitempty" xmlrpc:"laborTaxAmount"`

	// An invoice item's location, if one exists.'
	Location *Location `json:"location,omitempty" xmlrpc:"location"`

	// An Invoice Item's associated child invoice items, excluding ALL items with a $0.00 recurring fee. Only parent invoice items have associated children. For instance, a server invoice item may have associated children.
	NonZeroAssociatedChildren []Billing_Invoice_Item `json:"nonZeroAssociatedChildren,omitempty" xmlrpc:"nonZeroAssociatedChildren"`

	// A count of an Invoice Item's associated child invoice items, excluding ALL items with a $0.00 recurring fee. Only parent invoice items have associated children. For instance, a server invoice item may have associated children.
	NonZeroAssociatedChildrenCount *uint `json:"nonZeroAssociatedChildrenCount,omitempty" xmlrpc:"nonZeroAssociatedChildrenCount"`

	// A note to help describe more about the item. This normally holds usernames, or some other bit of extra information.
	Notes *string `json:"notes,omitempty" xmlrpc:"notes"`

	// An invoice item's one-time fee total after taxes. This does not include any child invoice items.
	OneTimeAfterTaxAmount *Float64 `json:"oneTimeAfterTaxAmount,omitempty" xmlrpc:"oneTimeAfterTaxAmount"`

	// If there are any one-time charges assessed, it will show up here represented by a floating point decimal in US Dollars ($USD)
	OneTimeFee *Float64 `json:"oneTimeFee,omitempty" xmlrpc:"oneTimeFee"`

	// The rate at which the one-time fee is taxed.
	OneTimeFeeTaxRate *Float64 `json:"oneTimeFeeTaxRate,omitempty" xmlrpc:"oneTimeFeeTaxRate"`

	// An invoice item's one-time tax amount. This does not include any child invoice items.
	OneTimeTaxAmount *Float64 `json:"oneTimeTaxAmount,omitempty" xmlrpc:"oneTimeTaxAmount"`

	// Every item tied to a server should have a parent invoice item which is the server line item. This is how we associate items to a server.
	Parent *Billing_Invoice_Item `json:"parent,omitempty" xmlrpc:"parent"`

	// The parent invoice item, usually the server invoice item.
	ParentId *int `json:"parentId,omitempty" xmlrpc:"parentId"`

	// The entry in the product catalog that a invoice item is based upon.
	Product *Product_Item `json:"product,omitempty" xmlrpc:"product"`

	// The entry in the product catalog that a invoice item is based upon.
	ProductItemId *int `json:"productItemId,omitempty" xmlrpc:"productItemId"`

	// An invoice item's recurring fee total after taxes. This does not include any child invoice items.
	RecurringAfterTaxAmount *Float64 `json:"recurringAfterTaxAmount,omitempty" xmlrpc:"recurringAfterTaxAmount"`

	// The recurring fee of the invoice item represented by a floating point decimal in US Dollars ($USD)
	RecurringFee *Float64 `json:"recurringFee,omitempty" xmlrpc:"recurringFee"`

	// the rate at which the recurring fee is taxed.
	RecurringFeeTaxRate *Float64 `json:"recurringFeeTaxRate,omitempty" xmlrpc:"recurringFeeTaxRate"`

	// An invoice item's recurring tax amount. This does not include any child invoice items.
	RecurringTaxAmount *Float64 `json:"recurringTaxAmount,omitempty" xmlrpc:"recurringTaxAmount"`

	// A unique identifier for a SoftLayer Service that is associated to an invoice item.
	ResourceTableId *int `json:"resourceTableId,omitempty" xmlrpc:"resourceTableId"`

	// An invoice item's setup fee total after taxes. This does not include any child invoice items.
	SetupAfterTaxAmount *Float64 `json:"setupAfterTaxAmount,omitempty" xmlrpc:"setupAfterTaxAmount"`

	// If there were any setup fees they will show up here. These are normally a one-time fee.
	SetupFee *Float64 `json:"setupFee,omitempty" xmlrpc:"setupFee"`

	// The tax rate at which the setup fee is taxed.
	SetupFeeTaxRate *Float64 `json:"setupFeeTaxRate,omitempty" xmlrpc:"setupFeeTaxRate"`

	// An invoice item's setup tax amount. This does not include any child invoice items.
	SetupTaxAmount *Float64 `json:"setupTaxAmount,omitempty" xmlrpc:"setupTaxAmount"`

	// An invoice Item's total, including any child invoice items if they exist.
	TotalOneTimeAmount *Float64 `json:"totalOneTimeAmount,omitempty" xmlrpc:"totalOneTimeAmount"`

	// An invoice Item's total, including any child invoice items if they exist.
	TotalOneTimeTaxAmount *Float64 `json:"totalOneTimeTaxAmount,omitempty" xmlrpc:"totalOneTimeTaxAmount"`

	// An invoice Item's total, including any child invoice items if they exist.
	TotalRecurringAmount *Float64 `json:"totalRecurringAmount,omitempty" xmlrpc:"totalRecurringAmount"`

	// A Billing Item's total, including any child billing items if they exist.'
	TotalRecurringTaxAmount *Float64 `json:"totalRecurringTaxAmount,omitempty" xmlrpc:"totalRecurringTaxAmount"`
}

// The SoftLayer_Billing_Invoice_Item_Hardware data type contains a "resource". This resource is a link to the hardware tied to a SoftLayer_Billing_item whose category code is "server".
type Billing_Invoice_Item_Hardware struct {
	Billing_Invoice_Item

	// The resource for a server invoice item.
	Resource *Hardware `json:"resource,omitempty" xmlrpc:"resource"`
}

// Information about the tax rates that apply to a particular invoice item.
type Billing_Invoice_Item_Tax_Info struct {
	Entity

	// The date and time the tax information was recorded.
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// The invoice description with special information about the invoice.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// The tax rate that can be multiplied by the subtotal to get the
	EffectiveTaxRate *Float64 `json:"effectiveTaxRate,omitempty" xmlrpc:"effectiveTaxRate"`

	// The amount that is exempt from tax.
	ExemptAmount *Float64 `json:"exemptAmount,omitempty" xmlrpc:"exemptAmount"`

	// The type of fee being tracked for this particular set of tax information.
	FeeProperty *string `json:"feeProperty,omitempty" xmlrpc:"feeProperty"`

	// An invoice item's tax information internal identifier.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	InvoiceItem *Billing_Invoice_Item `json:"invoiceItem,omitempty" xmlrpc:"invoiceItem"`

	// A reference to the related invoice item.
	InvoiceItemId *int `json:"invoiceItemId,omitempty" xmlrpc:"invoiceItemId"`

	// no documentation yet
	InvoiceTaxInfo *Billing_Invoice_Tax_Info `json:"invoiceTaxInfo,omitempty" xmlrpc:"invoiceTaxInfo"`

	// A reference to the tax information for the parent invoice.
	InvoiceTaxInfoId *int `json:"invoiceTaxInfoId,omitempty" xmlrpc:"invoiceTaxInfoId"`

	// The date and time the tax information was modified.
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// The amount that is exempt from tax.
	NonTaxableBasis *Float64 `json:"nonTaxableBasis,omitempty" xmlrpc:"nonTaxableBasis"`

	// A flag to indicate whether this is the official record for this invoice item.
	ReportedFlag *bool `json:"reportedFlag,omitempty" xmlrpc:"reportedFlag"`

	// The registration that the seller will use to report the invoice.
	SellerRegistration *string `json:"sellerRegistration,omitempty" xmlrpc:"sellerRegistration"`

	// The tax amount associated with this line item.
	TaxAmount *Float64 `json:"taxAmount,omitempty" xmlrpc:"taxAmount"`

	// The tax amount (converted to the 'to' currency) associated with this line item.
	TaxAmountToCurrency *Float64 `json:"taxAmountToCurrency,omitempty" xmlrpc:"taxAmountToCurrency"`

	// The tax rate used. Note that this might apply to only part of the
	TaxRate *Float64 `json:"taxRate,omitempty" xmlrpc:"taxRate"`

	// The amount that is subject to tax.
	TaxableBasis *Float64 `json:"taxableBasis,omitempty" xmlrpc:"taxableBasis"`

	// This is the currency the invoice will be converted to.
	ToCurrency *Billing_Currency `json:"toCurrency,omitempty" xmlrpc:"toCurrency"`

	// The currency code that the invoice is being converted to.
	ToCurrencyId *int `json:"toCurrencyId,omitempty" xmlrpc:"toCurrencyId"`
}

// no documentation yet
type Billing_Invoice_Next struct {
	Entity
}

// The SoftLayer_Billing_Invoice_Receivable_Payment data type contains general information relating to payments made against invoices.
type Billing_Invoice_Receivable_Payment struct {
	Entity

	// no documentation yet
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// The amount of the payment.
	Amount *Float64 `json:"amount,omitempty" xmlrpc:"amount"`

	// The date of the payment.
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// no documentation yet
	CreditCardLastFourDigits *int `json:"creditCardLastFourDigits,omitempty" xmlrpc:"creditCardLastFourDigits"`

	// no documentation yet
	CreditCardRequestId *string `json:"creditCardRequestId,omitempty" xmlrpc:"creditCardRequestId"`

	// no documentation yet
	CreditCardTransaction *Billing_Payment_Card_Transaction `json:"creditCardTransaction,omitempty" xmlrpc:"creditCardTransaction"`

	// no documentation yet
	ExchangeRate *Billing_Currency_ExchangeRate `json:"exchangeRate,omitempty" xmlrpc:"exchangeRate"`

	// no documentation yet
	Invoice *Billing_Invoice `json:"invoice,omitempty" xmlrpc:"invoice"`

	// The invoice that the payment is for.
	InvoiceId *int `json:"invoiceId,omitempty" xmlrpc:"invoiceId"`

	// no documentation yet
	PaypalTransaction *Billing_Payment_PayPal_Transaction `json:"paypalTransaction,omitempty" xmlrpc:"paypalTransaction"`

	// The type of payment.
	TypeCode *string `json:"typeCode,omitempty" xmlrpc:"typeCode"`
}

// Invoice tax information contains top-level information about the taxes recorded for a particular invoice.
type Billing_Invoice_Tax_Info struct {
	Entity

	// The date and time this tax information was recorded.
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// This is the currency used for the invoice.
	Currency *Billing_Currency `json:"currency,omitempty" xmlrpc:"currency"`

	// The currency code that the invoice should be recorded in.
	CurrencyId *int `json:"currencyId,omitempty" xmlrpc:"currencyId"`

	// This is the functional currency used for the invoice.
	FunctionalCurrency *Billing_Currency `json:"functionalCurrency,omitempty" xmlrpc:"functionalCurrency"`

	// The internal identifier for this invoice tax information.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// This is the related invoice for this tax-related information.
	Invoice *Billing_Invoice `json:"invoice,omitempty" xmlrpc:"invoice"`

	// A reference to the related invoice.
	InvoiceId *int `json:"invoiceId,omitempty" xmlrpc:"invoiceId"`

	// A count of this is the collection of tax information for each of the related invoice items.
	ItemCount *uint `json:"itemCount,omitempty" xmlrpc:"itemCount"`

	// This tax information on the invoice item that includes currency details.
	ItemWithCurrencyInfo *Billing_Invoice_Item_Tax_Info `json:"itemWithCurrencyInfo,omitempty" xmlrpc:"itemWithCurrencyInfo"`

	// This is the collection of tax information for each of the related invoice items.
	Items []Billing_Invoice_Item_Tax_Info `json:"items,omitempty" xmlrpc:"items"`

	// The date and time this tax information was updated.
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// A flag to indicate whether the invoice will be auditable.
	ReportedFlag *bool `json:"reportedFlag,omitempty" xmlrpc:"reportedFlag"`

	// This the total tax amount (converted to the 'to' currency) for the invoice.
	TotalTaxAmountToCurrency *Float64 `json:"totalTaxAmountToCurrency,omitempty" xmlrpc:"totalTaxAmountToCurrency"`
}

// The invoice tax status data type models a single status or state that an invoice can reflect in regard to an integration with a third-party tax calculation service.
type Billing_Invoice_Tax_Status struct {
	Entity

	// no documentation yet
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// no documentation yet
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// no documentation yet
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// The invoice tax type data type models a single strategy for handling tax calculations.
type Billing_Invoice_Tax_Type struct {
	Entity

	// A tax type's internal identifier. Each type of tax calculation strategy has a unique ID value.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// A unique string that identifies each strategy and is guaranteed to be stable over time.
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// A human-readable label for each tax strategy.
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// Every individual item that a SoftLayer customer is billed for is recorded in the SoftLayer_Billing_Item data type. Billing items range from server chassis to hard drives to control panels, bandwidth quota upgrades and port upgrade charges. Softlayer [[SoftLayer_Billing_Invoice|invoices]] are generated from the cost of a customer's billing items. Billing items are copied from the product catalog as they're ordered by customers to create a reference between an account and the billable items they own.
//
// Billing items exist in a tree relationship. Items are associated with each other by parent/child relationships. Component items such as CPU's, RAM, and software each have a parent billing item for the server chassis they're associated with. Billing Items with a null parent item do not have an associated parent item.
type Billing_Item struct {
	Entity

	// The account that a billing item belongs to.
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// no documentation yet
	ActiveAgreement *Account_Agreement `json:"activeAgreement,omitempty" xmlrpc:"activeAgreement"`

	// A flag indicating that the billing item is under an active agreement.
	ActiveAgreementFlag *Account_Agreement `json:"activeAgreementFlag,omitempty" xmlrpc:"activeAgreementFlag"`

	// A billing item's active associated child billing items. This includes "floating" items that are not necessarily child items of this billing item.
	ActiveAssociatedChildren []Billing_Item `json:"activeAssociatedChildren,omitempty" xmlrpc:"activeAssociatedChildren"`

	// A count of a billing item's active associated child billing items. This includes "floating" items that are not necessarily child items of this billing item.
	ActiveAssociatedChildrenCount *uint `json:"activeAssociatedChildrenCount,omitempty" xmlrpc:"activeAssociatedChildrenCount"`

	// A count of
	ActiveAssociatedGuestDiskBillingItemCount *uint `json:"activeAssociatedGuestDiskBillingItemCount,omitempty" xmlrpc:"activeAssociatedGuestDiskBillingItemCount"`

	// no documentation yet
	ActiveAssociatedGuestDiskBillingItems []Billing_Item `json:"activeAssociatedGuestDiskBillingItems,omitempty" xmlrpc:"activeAssociatedGuestDiskBillingItems"`

	// A count of a Billing Item's active bundled billing items.
	ActiveBundledItemCount *uint `json:"activeBundledItemCount,omitempty" xmlrpc:"activeBundledItemCount"`

	// A Billing Item's active bundled billing items.
	ActiveBundledItems []Billing_Item `json:"activeBundledItems,omitempty" xmlrpc:"activeBundledItems"`

	// A service cancellation request item that corresponds to the billing item.
	ActiveCancellationItem *Billing_Item_Cancellation_Request_Item `json:"activeCancellationItem,omitempty" xmlrpc:"activeCancellationItem"`

	// A Billing Item's active child billing items.
	ActiveChildren []Billing_Item `json:"activeChildren,omitempty" xmlrpc:"activeChildren"`

	// A count of a Billing Item's active child billing items.
	ActiveChildrenCount *uint `json:"activeChildrenCount,omitempty" xmlrpc:"activeChildrenCount"`

	// no documentation yet
	ActiveFlag *bool `json:"activeFlag,omitempty" xmlrpc:"activeFlag"`

	// A count of
	ActiveSparePoolAssociatedGuestDiskBillingItemCount *uint `json:"activeSparePoolAssociatedGuestDiskBillingItemCount,omitempty" xmlrpc:"activeSparePoolAssociatedGuestDiskBillingItemCount"`

	// no documentation yet
	ActiveSparePoolAssociatedGuestDiskBillingItems []Billing_Item `json:"activeSparePoolAssociatedGuestDiskBillingItems,omitempty" xmlrpc:"activeSparePoolAssociatedGuestDiskBillingItems"`

	// A count of a Billing Item's spare pool bundled billing items.
	ActiveSparePoolBundledItemCount *uint `json:"activeSparePoolBundledItemCount,omitempty" xmlrpc:"activeSparePoolBundledItemCount"`

	// A Billing Item's spare pool bundled billing items.
	ActiveSparePoolBundledItems []Billing_Item `json:"activeSparePoolBundledItems,omitempty" xmlrpc:"activeSparePoolBundledItems"`

	// Flag to check if a billing item can be cancelled. 1 = yes. 0 = no.
	AllowCancellationFlag *int `json:"allowCancellationFlag,omitempty" xmlrpc:"allowCancellationFlag"`

	// A billing item's associated parent. This is to be used for billing items that are "floating", and therefore are not child items of any parent billing item. If it is desired to associate an item to another, populate this with the SoftLayer_Billing_Item ID of that associated parent item.
	AssociatedBillingItem *Billing_Item `json:"associatedBillingItem,omitempty" xmlrpc:"associatedBillingItem"`

	// A history of billing items which a billing item has been associated with.
	AssociatedBillingItemHistory []Billing_Item_Association_History `json:"associatedBillingItemHistory,omitempty" xmlrpc:"associatedBillingItemHistory"`

	// A count of a history of billing items which a billing item has been associated with.
	AssociatedBillingItemHistoryCount *uint `json:"associatedBillingItemHistoryCount,omitempty" xmlrpc:"associatedBillingItemHistoryCount"`

	// This is sometimes populated for orphan billing items that are not attached to servers. Billing items like secondary portable IP addresses fit into this category. A user may set an association by calling [[SoftLayer_Billing_Item::setAssociationId]]. This will cause this orphan item to appear under its associated server billing item on future invoices. You may only attach orphaned billing items to server billing items without cancellation dates set.
	AssociatedBillingItemId *string `json:"associatedBillingItemId,omitempty" xmlrpc:"associatedBillingItemId"`

	// A Billing Item's associated child billing items. This includes "floating" items that are not necessarily child billing items of this billing item.
	AssociatedChildren []Billing_Item `json:"associatedChildren,omitempty" xmlrpc:"associatedChildren"`

	// A count of a Billing Item's associated child billing items. This includes "floating" items that are not necessarily child billing items of this billing item.
	AssociatedChildrenCount *uint `json:"associatedChildrenCount,omitempty" xmlrpc:"associatedChildrenCount"`

	// A billing item's associated parent billing item. This object will be the same as the parent billing item if parentId is set.
	AssociatedParent []Billing_Item `json:"associatedParent,omitempty" xmlrpc:"associatedParent"`

	// A count of a billing item's associated parent billing item. This object will be the same as the parent billing item if parentId is set.
	AssociatedParentCount *uint `json:"associatedParentCount,omitempty" xmlrpc:"associatedParentCount"`

	// A count of
	AvailableMatchingVlanCount *uint `json:"availableMatchingVlanCount,omitempty" xmlrpc:"availableMatchingVlanCount"`

	// no documentation yet
	AvailableMatchingVlans []Network_Vlan `json:"availableMatchingVlans,omitempty" xmlrpc:"availableMatchingVlans"`

	// The bandwidth allocation for a billing item.
	BandwidthAllocation *Network_Bandwidth_Version1_Allocation `json:"bandwidthAllocation,omitempty" xmlrpc:"bandwidthAllocation"`

	// A billing item's recurring child items that have once been billed and are scheduled to be billed in the future.
	BillableChildren []Billing_Item `json:"billableChildren,omitempty" xmlrpc:"billableChildren"`

	// A count of a billing item's recurring child items that have once been billed and are scheduled to be billed in the future.
	BillableChildrenCount *uint `json:"billableChildrenCount,omitempty" xmlrpc:"billableChildrenCount"`

	// A count of a Billing Item's bundled billing items
	BundleItemCount *uint `json:"bundleItemCount,omitempty" xmlrpc:"bundleItemCount"`

	// A Billing Item's bundled billing items
	BundleItems []Product_Item_Bundles `json:"bundleItems,omitempty" xmlrpc:"bundleItems"`

	// A count of a Billing Item's bundled billing items'
	BundledItemCount *uint `json:"bundledItemCount,omitempty" xmlrpc:"bundledItemCount"`

	// A Billing Item's bundled billing items'
	BundledItems []Billing_Item `json:"bundledItems,omitempty" xmlrpc:"bundledItems"`

	// A Billing Item's active child billing items.
	CanceledChildren []Billing_Item `json:"canceledChildren,omitempty" xmlrpc:"canceledChildren"`

	// A count of a Billing Item's active child billing items.
	CanceledChildrenCount *uint `json:"canceledChildrenCount,omitempty" xmlrpc:"canceledChildrenCount"`

	// A billing item's cancellation date. A billing item with a cancellation date in the past is not charged on your SoftLayer invoice. Cancellation dates in the future indicate the current billing item is active, but will be cancelled and not charged for in the future. A billing item with a null cancellation date is also considered an active billing item and is charged once every billing cycle.
	CancellationDate *Time `json:"cancellationDate,omitempty" xmlrpc:"cancellationDate"`

	// The billing item's cancellation reason.
	CancellationReason *Billing_Item_Cancellation_Reason `json:"cancellationReason,omitempty" xmlrpc:"cancellationReason"`

	// A count of this will return any cancellation requests that are associated with this billing item.
	CancellationRequestCount *uint `json:"cancellationRequestCount,omitempty" xmlrpc:"cancellationRequestCount"`

	// This will return any cancellation requests that are associated with this billing item.
	CancellationRequests []Billing_Item_Cancellation_Request `json:"cancellationRequests,omitempty" xmlrpc:"cancellationRequests"`

	// The item category to which the billing item's item belongs.
	Category *Product_Item_Category `json:"category,omitempty" xmlrpc:"category"`

	// The category code of this billing item. It is used to tell us the difference between a primary disk and a secondary disk, for instance.
	CategoryCode *string `json:"categoryCode,omitempty" xmlrpc:"categoryCode"`

	// A Billing Item's child billing items'
	Children []Billing_Item `json:"children,omitempty" xmlrpc:"children"`

	// A count of a Billing Item's child billing items'
	ChildrenCount *uint `json:"childrenCount,omitempty" xmlrpc:"childrenCount"`

	// A Billing Item's active child billing items.
	ChildrenWithActiveAgreement []Billing_Item `json:"childrenWithActiveAgreement,omitempty" xmlrpc:"childrenWithActiveAgreement"`

	// A count of a Billing Item's active child billing items.
	ChildrenWithActiveAgreementCount *uint `json:"childrenWithActiveAgreementCount,omitempty" xmlrpc:"childrenWithActiveAgreementCount"`

	// The date the billing item was created. You can see this date on the invoice.
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// This is the total charge for the billing item for this billing item. It is calculated based on the hourlyRecurringFee * hoursUsed.
	CurrentHourlyCharge *string `json:"currentHourlyCharge,omitempty" xmlrpc:"currentHourlyCharge"`

	// The last time this billing item was charged.
	CycleStartDate *Time `json:"cycleStartDate,omitempty" xmlrpc:"cycleStartDate"`

	// A brief description of a billing item.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// The domain name is provided for server billing items.
	DomainName *string `json:"domainName,omitempty" xmlrpc:"domainName"`

	// A count of for product items which have a downgrade path defined, this will return those product items.
	DowngradeItemCount *uint `json:"downgradeItemCount,omitempty" xmlrpc:"downgradeItemCount"`

	// For product items which have a downgrade path defined, this will return those product items.
	DowngradeItems []Product_Item `json:"downgradeItems,omitempty" xmlrpc:"downgradeItems"`

	// A Billing Item's associated child billing items, excluding some items with a $0.00 recurring fee.
	FilteredNextInvoiceChildren []Billing_Item `json:"filteredNextInvoiceChildren,omitempty" xmlrpc:"filteredNextInvoiceChildren"`

	// A count of a Billing Item's associated child billing items, excluding some items with a $0.00 recurring fee.
	FilteredNextInvoiceChildrenCount *uint `json:"filteredNextInvoiceChildrenCount,omitempty" xmlrpc:"filteredNextInvoiceChildrenCount"`

	// The hostname is provided for server billing items
	HostName *string `json:"hostName,omitempty" xmlrpc:"hostName"`

	// A flag that will reflect whether this billing item is billed on an hourly basis or not.
	HourlyFlag *bool `json:"hourlyFlag,omitempty" xmlrpc:"hourlyFlag"`

	// The amount of money charged per hour for a billing item, if applicable. hourlyRecurringFee is measured in US Dollars ($USD).
	HourlyRecurringFee *Float64 `json:"hourlyRecurringFee,omitempty" xmlrpc:"hourlyRecurringFee"`

	// This is the number of hours the hourly billing item has been in use this billing period. For virtual servers, this means running, paused or stopped.
	HoursUsed *string `json:"hoursUsed,omitempty" xmlrpc:"hoursUsed"`

	// The unique identifier for this billing item.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// Invoice items associated with this billing item
	InvoiceItem *Billing_Invoice_Item `json:"invoiceItem,omitempty" xmlrpc:"invoiceItem"`

	// A count of all invoice items associated with the billing item
	InvoiceItemCount *uint `json:"invoiceItemCount,omitempty" xmlrpc:"invoiceItemCount"`

	// All invoice items associated with the billing item
	InvoiceItems []Billing_Invoice_Item `json:"invoiceItems,omitempty" xmlrpc:"invoiceItems"`

	// The entry in the SoftLayer product catalog that a billing item is based upon.
	Item *Product_Item `json:"item,omitempty" xmlrpc:"item"`

	// The labor fee, if any. This is a one time charge.
	LaborFee *Float64 `json:"laborFee,omitempty" xmlrpc:"laborFee"`

	// The rate at which labor fees are taxed if you are a taxable customer.
	LaborFeeTaxRate *Float64 `json:"laborFeeTaxRate,omitempty" xmlrpc:"laborFeeTaxRate"`

	// The last time this billing item was charged.
	LastBillDate *Time `json:"lastBillDate,omitempty" xmlrpc:"lastBillDate"`

	// The location of the billing item. Some billing items have physical properties such as the server itself. For items such as these, we provide location information.
	Location *Location `json:"location,omitempty" xmlrpc:"location"`

	// The date that a billing item was last modified.
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// The date on which your account will be charged for this billing item.
	NextBillDate *Time `json:"nextBillDate,omitempty" xmlrpc:"nextBillDate"`

	// A Billing Item's child billing items and associated items'
	NextInvoiceChildren []Billing_Item `json:"nextInvoiceChildren,omitempty" xmlrpc:"nextInvoiceChildren"`

	// A count of a Billing Item's child billing items and associated items'
	NextInvoiceChildrenCount *uint `json:"nextInvoiceChildrenCount,omitempty" xmlrpc:"nextInvoiceChildrenCount"`

	// A Billing Item's total, including any child billing items if they exist.'
	NextInvoiceTotalOneTimeAmount *Float64 `json:"nextInvoiceTotalOneTimeAmount,omitempty" xmlrpc:"nextInvoiceTotalOneTimeAmount"`

	// A Billing Item's total, including any child billing items if they exist.'
	NextInvoiceTotalOneTimeTaxAmount *Float64 `json:"nextInvoiceTotalOneTimeTaxAmount,omitempty" xmlrpc:"nextInvoiceTotalOneTimeTaxAmount"`

	// A Billing Item's total, including any child billing items and associated billing items if they exist.'
	NextInvoiceTotalRecurringAmount *Float64 `json:"nextInvoiceTotalRecurringAmount,omitempty" xmlrpc:"nextInvoiceTotalRecurringAmount"`

	// This is deprecated and will always be zero. Because tax is calculated in real-time, previewing the next recurring invoice is pre-tax only.
	NextInvoiceTotalRecurringTaxAmount *Float64 `json:"nextInvoiceTotalRecurringTaxAmount,omitempty" xmlrpc:"nextInvoiceTotalRecurringTaxAmount"`

	// A Billing Item's associated child billing items, excluding ALL items with a $0.00 recurring fee.
	NonZeroNextInvoiceChildren []Billing_Item `json:"nonZeroNextInvoiceChildren,omitempty" xmlrpc:"nonZeroNextInvoiceChildren"`

	// A count of a Billing Item's associated child billing items, excluding ALL items with a $0.00 recurring fee.
	NonZeroNextInvoiceChildrenCount *uint `json:"nonZeroNextInvoiceChildrenCount,omitempty" xmlrpc:"nonZeroNextInvoiceChildrenCount"`

	// Extra information provided to help you identify this billing item. This is often a username or something to help identify items that customers have more than one of.
	Notes *string `json:"notes,omitempty" xmlrpc:"notes"`

	// The amount of money charged as a one-time charge for a billing item, if applicable. oneTimeFee is measured in US Dollars ($USD).
	OneTimeFee *Float64 `json:"oneTimeFee,omitempty" xmlrpc:"oneTimeFee"`

	// The rate at which one time fees are taxed if you are a taxable customer.
	OneTimeFeeTaxRate *Float64 `json:"oneTimeFeeTaxRate,omitempty" xmlrpc:"oneTimeFeeTaxRate"`

	// A billing item's original order item. Simply a reference to the original order from which this billing item was created.
	OrderItem *Billing_Order_Item `json:"orderItem,omitempty" xmlrpc:"orderItem"`

	// the SoftLayer_Billing_Order_Item ID. This is a reference to the original order item from which this billing item was originally created.
	OrderItemId *int `json:"orderItemId,omitempty" xmlrpc:"orderItemId"`

	// The original physical location for this billing item--may differ from current.
	OriginalLocation *Location `json:"originalLocation,omitempty" xmlrpc:"originalLocation"`

	// The package under which this billing item was sold. A Package is the general grouping of products as seen on our order forms.
	Package *Product_Package `json:"package,omitempty" xmlrpc:"package"`

	// A billing item's parent item. If a billing item has no parent item then this value is null.
	Parent *Billing_Item `json:"parent,omitempty" xmlrpc:"parent"`

	// The unique identifier of the parent of this billing item.
	ParentId *int `json:"parentId,omitempty" xmlrpc:"parentId"`

	// A billing item's parent item. If a billing item has no parent item then this value is null.
	ParentVirtualGuestBillingItem *Billing_Item_Virtual_Guest `json:"parentVirtualGuestBillingItem,omitempty" xmlrpc:"parentVirtualGuestBillingItem"`

	// This flag indicates whether a billing item is scheduled to be canceled or not.
	PendingCancellationFlag *bool `json:"pendingCancellationFlag,omitempty" xmlrpc:"pendingCancellationFlag"`

	// The new order item that will replace this billing item.
	PendingOrderItem *Billing_Order_Item `json:"pendingOrderItem,omitempty" xmlrpc:"pendingOrderItem"`

	// Provisioning transaction for this billing item
	ProvisionTransaction *Provisioning_Version1_Transaction `json:"provisionTransaction,omitempty" xmlrpc:"provisionTransaction"`

	// The amount of money charged per month for a billing item, if applicable. recurringFee is measured in US Dollars ($USD).
	RecurringFee *Float64 `json:"recurringFee,omitempty" xmlrpc:"recurringFee"`

	// The rate at which recurring fees are taxed if you are a taxable customer.
	RecurringFeeTaxRate *Float64 `json:"recurringFeeTaxRate,omitempty" xmlrpc:"recurringFeeTaxRate"`

	// The number of months in which the recurring fees will be incurred.
	RecurringMonths *int `json:"recurringMonths,omitempty" xmlrpc:"recurringMonths"`

	// This is the service provider for this billing item.
	ServiceProviderId *int `json:"serviceProviderId,omitempty" xmlrpc:"serviceProviderId"`

	// The setup fee, if any. This is a one time charge.
	SetupFee *Float64 `json:"setupFee,omitempty" xmlrpc:"setupFee"`

	// The rate at which setup fees are taxed if you are a taxable customer.
	SetupFeeTaxRate *Float64 `json:"setupFeeTaxRate,omitempty" xmlrpc:"setupFeeTaxRate"`

	// A friendly description of software component
	SoftwareDescription *Software_Description `json:"softwareDescription,omitempty" xmlrpc:"softwareDescription"`

	// Billing items whose product item has an upgrade path defined in our system will return the next product item in the upgrade path.
	UpgradeItem *Product_Item `json:"upgradeItem,omitempty" xmlrpc:"upgradeItem"`

	// A count of billing items whose product item has an upgrade path defined in our system will return all the product items in the upgrade path.
	UpgradeItemCount *uint `json:"upgradeItemCount,omitempty" xmlrpc:"upgradeItemCount"`

	// Billing items whose product item has an upgrade path defined in our system will return all the product items in the upgrade path.
	UpgradeItems []Product_Item `json:"upgradeItems,omitempty" xmlrpc:"upgradeItems"`
}

// The SoftLayer_Billing_Item_Account_Media_Data_Transfer_Request data type contains general information relating to a single SoftLayer billing item for a data transfer request.
type Billing_Item_Account_Media_Data_Transfer_Request struct {
	Billing_Item

	// The data transfer request to which the billing item points.
	Resource *Account_Media_Data_Transfer_Request `json:"resource,omitempty" xmlrpc:"resource"`
}

// The SoftLayer_Billing_Item_Association_History type keeps a record of which server billing items an "orphan" item has been associated with. Orphan billing items are billable items for secondary portable services (such as secondary subnets and StorageLayer accounts) that are not associated with a server and appear at the bottom of a SoftLayer invoice. The [[SoftLayer_Billing_Item::setAssociationId]] method allows you to associate these kinds of items with servers, making them appear as a child item of the server on your invoice. A SoftLayer_Billing_Item_Association_History record is created every time one of these associations are set.
type Billing_Item_Association_History struct {
	Entity

	// The server billing item that an orphaned billing item was associated with.
	AssociatedBillingItem *Billing_Item `json:"associatedBillingItem,omitempty" xmlrpc:"associatedBillingItem"`

	// The internal identifier of the server billing item that an orphaned billing item was associated with.
	AssociatedBillingItemId *int `json:"associatedBillingItemId,omitempty" xmlrpc:"associatedBillingItemId"`

	// The billing item that was associated with a server billing item.
	BillingItem *Billing_Item `json:"billingItem,omitempty" xmlrpc:"billingItem"`

	// The internal identifier of the billing item that was associated with a server billing item.
	BillingItemId *int `json:"billingItemId,omitempty" xmlrpc:"billingItemId"`

	// The date that a billing item association was last changed.
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// A billing item association history's internal identifier.
	Id *int `json:"id,omitempty" xmlrpc:"id"`
}

// The SoftLayer_Billing_Item_Cancellation_Reason data type contains cancellation reasons.
type Billing_Item_Cancellation_Reason struct {
	Entity

	// A cancel reason category internal identifier.
	BillingCancelReasonCategoryId *int `json:"billingCancelReasonCategoryId,omitempty" xmlrpc:"billingCancelReasonCategoryId"`

	// An billing cancellation reason category.
	BillingCancellationReasonCategory *Billing_Item_Cancellation_Reason_Category `json:"billingCancellationReasonCategory,omitempty" xmlrpc:"billingCancellationReasonCategory"`

	// A count of the corresponding billing items having the specific cancellation reason.
	BillingItemCount *uint `json:"billingItemCount,omitempty" xmlrpc:"billingItemCount"`

	// The corresponding billing items having the specific cancellation reason.
	BillingItems []Billing_Item `json:"billingItems,omitempty" xmlrpc:"billingItems"`

	// A reason internal identifier.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// A standardized reason internal identifier.
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// The descriptoin of the reason
	Reason *string `json:"reason,omitempty" xmlrpc:"reason"`

	// no documentation yet
	TranslatedReason *string `json:"translatedReason,omitempty" xmlrpc:"translatedReason"`
}

// The SoftLayer_Billing_Item_Cancellation_Reason_Category data type contains cancellation reason categories.
type Billing_Item_Cancellation_Reason_Category struct {
	Entity

	// A count of the corresponding billing cancellation reasons having the specific billing cancellation reason category.
	BillingCancellationReasonCount *uint `json:"billingCancellationReasonCount,omitempty" xmlrpc:"billingCancellationReasonCount"`

	// The corresponding billing cancellation reasons having the specific billing cancellation reason category.
	BillingCancellationReasons []Billing_Item_Cancellation_Reason `json:"billingCancellationReasons,omitempty" xmlrpc:"billingCancellationReasons"`

	// A category internal identifier.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// The description of the category
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// SoftLayer_Billing_Item_Cancellation_Request data type is used to cancel service billing items.
type Billing_Item_Cancellation_Request struct {
	Entity

	// The SoftLayer account that a service cancellation request belongs to.
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// The internal identifier of the customer account that a service cancellation record belongs to.
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// The last modified date.
	BillingCancelReasonId *int `json:"billingCancelReasonId,omitempty" xmlrpc:"billingCancelReasonId"`

	// The date that a cancellation request was created.
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// A cancellation record's internal identifier.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// A count of a collection of service cancellation items.
	ItemCount *uint `json:"itemCount,omitempty" xmlrpc:"itemCount"`

	// A collection of service cancellation items.
	Items []Billing_Item_Cancellation_Request_Item `json:"items,omitempty" xmlrpc:"items"`

	// The last modified date.
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// Brief cancellation note.
	Notes *string `json:"notes,omitempty" xmlrpc:"notes"`

	// The status of a service cancellation request.
	Status *Billing_Item_Cancellation_Request_Status `json:"status,omitempty" xmlrpc:"status"`

	// An internal identifier of the service cancellation status that this request is associated with. When a service cancellation is submitted, it will be in "Pending" status until SoftLayer Sales team reviews it. The status of a cancellation request will be updated to "Approved" or "Voided" by SoftLayer Sales.
	//
	// It will be updated to "Complete" when all services are reclaimed.
	StatusId *int `json:"statusId,omitempty" xmlrpc:"statusId"`

	// The ticket that is associated with the service cancellation request.
	Ticket *Ticket `json:"ticket,omitempty" xmlrpc:"ticket"`

	// An internal identifier of the ticket that is associated with a service cancellation request. When a service cancellation is submitted, a support ticket will be created. This ticket contains details on your service cancellation details and SoftLayer Sales team will use it to further communicate with you.
	TicketId *int `json:"ticketId,omitempty" xmlrpc:"ticketId"`

	// The user that initiated a service cancellation request.
	User *User_Customer `json:"user,omitempty" xmlrpc:"user"`
}

// SoftLayer_Billing_Item_Cancellation_Request_Item data type contains a billing item for cancellation. This data type is used to harness billing items to the associated service.
type Billing_Item_Cancellation_Request_Item struct {
	Entity

	// The billing item for cancellation.
	BillingItem *Billing_Item `json:"billingItem,omitempty" xmlrpc:"billingItem"`

	// The internal identifier of a billing item
	BillingItemId *int `json:"billingItemId,omitempty" xmlrpc:"billingItemId"`

	// The service cancellation request that a cancellation item belongs to.
	CancellationRequest *Billing_Item_Cancellation_Request `json:"cancellationRequest,omitempty" xmlrpc:"cancellationRequest"`

	// A cancellation request's internal identifier.
	CancellationRequestId *int `json:"cancellationRequestId,omitempty" xmlrpc:"cancellationRequestId"`

	// A cancellation request item's internal identifier.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// This flag indicated if a billing item should be canceled immediately or not.  Set this flag to true when creating a cancellation request.
	ImmediateCancellationFlag *bool `json:"immediateCancellationFlag,omitempty" xmlrpc:"immediateCancellationFlag"`

	// The scheduled cancellation date
	ScheduledCancellationDate *Time `json:"scheduledCancellationDate,omitempty" xmlrpc:"scheduledCancellationDate"`

	// The reclaim status of a service.
	ServiceReclaimStatusCode *string `json:"serviceReclaimStatusCode,omitempty" xmlrpc:"serviceReclaimStatusCode"`
}

// SoftLayer_Billing_Item_Cancellation_Request_Status data type represents the status of a service cancellation request.
type Billing_Item_Cancellation_Request_Status struct {
	Entity

	// The short description of a cancellation request status
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// The internal identifier of a cancellation request status.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// status key name
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// The status name
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// The SoftLayer_Billing_Item_Ctc_Account data type contains general information relating to a single SoftLayer billing item for a CTC client account creation
type Billing_Item_Ctc_Account struct {
	Billing_Item
}

// The SoftLayer_Billing_Item_Big_Data_Cluster data type contains general information relating to a single SoftLayer billing item for a big data cluster.
type Billing_Item_Gateway_Appliance_Cluster struct {
	Billing_Item

	// The resource for a resource group billing item.
	Resource *Resource_Group `json:"resource,omitempty" xmlrpc:"resource"`
}

// The SoftLayer_Billing_Item_Hardware data type contains general information relating to a single SoftLayer billing item for hardware.
type Billing_Item_Hardware struct {
	Billing_Item

	// The raw bandwidth usage data for the current billing cycle. One object will be returned for each network this server is attached to.
	BillingCycleBandwidthUsage []Network_Bandwidth_Usage `json:"billingCycleBandwidthUsage,omitempty" xmlrpc:"billingCycleBandwidthUsage"`

	// A count of the raw bandwidth usage data for the current billing cycle. One object will be returned for each network this server is attached to.
	BillingCycleBandwidthUsageCount *uint `json:"billingCycleBandwidthUsageCount,omitempty" xmlrpc:"billingCycleBandwidthUsageCount"`

	// The raw private bandwidth usage data for the current billing cycle.
	BillingCyclePrivateBandwidthUsage []Network_Bandwidth_Usage `json:"billingCyclePrivateBandwidthUsage,omitempty" xmlrpc:"billingCyclePrivateBandwidthUsage"`

	// A count of the raw private bandwidth usage data for the current billing cycle.
	BillingCyclePrivateBandwidthUsageCount *uint `json:"billingCyclePrivateBandwidthUsageCount,omitempty" xmlrpc:"billingCyclePrivateBandwidthUsageCount"`

	// The total private inbound bandwidth for this hardware for the current billing cycle.
	BillingCyclePrivateUsageIn *Float64 `json:"billingCyclePrivateUsageIn,omitempty" xmlrpc:"billingCyclePrivateUsageIn"`

	// The total private outbound bandwidth for this hardware for the current billing cycle.
	BillingCyclePrivateUsageOut *Float64 `json:"billingCyclePrivateUsageOut,omitempty" xmlrpc:"billingCyclePrivateUsageOut"`

	// The total private bandwidth for this hardware for the current billing cycle.
	BillingCyclePrivateUsageTotal *uint `json:"billingCyclePrivateUsageTotal,omitempty" xmlrpc:"billingCyclePrivateUsageTotal"`

	// The raw public bandwidth usage data for the current billing cycle.
	BillingCyclePublicBandwidthUsage []Network_Bandwidth_Usage `json:"billingCyclePublicBandwidthUsage,omitempty" xmlrpc:"billingCyclePublicBandwidthUsage"`

	// A count of the raw public bandwidth usage data for the current billing cycle.
	BillingCyclePublicBandwidthUsageCount *uint `json:"billingCyclePublicBandwidthUsageCount,omitempty" xmlrpc:"billingCyclePublicBandwidthUsageCount"`

	// The total public inbound bandwidth for this hardware for the current billing cycle.
	BillingCyclePublicUsageIn *Float64 `json:"billingCyclePublicUsageIn,omitempty" xmlrpc:"billingCyclePublicUsageIn"`

	// The total public outbound bandwidth for this hardware for the current billing cycle.
	BillingCyclePublicUsageOut *Float64 `json:"billingCyclePublicUsageOut,omitempty" xmlrpc:"billingCyclePublicUsageOut"`

	// The total public bandwidth for this hardware for the current billing cycle.
	BillingCyclePublicUsageTotal *uint `json:"billingCyclePublicUsageTotal,omitempty" xmlrpc:"billingCyclePublicUsageTotal"`

	// A lockbox account associated with a server.
	LockboxNetworkStorage *Billing_Item_Network_Storage `json:"lockboxNetworkStorage,omitempty" xmlrpc:"lockboxNetworkStorage"`

	// A count of
	MonitoringBillingItemCount *uint `json:"monitoringBillingItemCount,omitempty" xmlrpc:"monitoringBillingItemCount"`

	// no documentation yet
	MonitoringBillingItems []Billing_Item `json:"monitoringBillingItems,omitempty" xmlrpc:"monitoringBillingItems"`

	// The resource for a server billing item.
	Resource *Hardware_Server `json:"resource,omitempty" xmlrpc:"resource"`

	// The resource (unique identifier) for a server billing item.
	ResourceTableId *int `json:"resourceTableId,omitempty" xmlrpc:"resourceTableId"`
}

// The SoftLayer_Billing_Item_Hardware data type contains general information relating to a single SoftLayer billing item for hardware.
type Billing_Item_Hardware_Colocation struct {
	Billing_Item_Hardware
}

// The SoftLayer_Billing_Item_Hardware data type contains general information relating to a single SoftLayer billing item for hardware components.
type Billing_Item_Hardware_Component struct {
	Billing_Item

	// The hardware component that this billing item points to.
	Resource []Hardware_Component `json:"resource,omitempty" xmlrpc:"resource"`

	// A count of the hardware component that this billing item points to.
	ResourceCount *uint `json:"resourceCount,omitempty" xmlrpc:"resourceCount"`

	// The resource (unique identifier) for a server billing item.
	ResourceTableId *int `json:"resourceTableId,omitempty" xmlrpc:"resourceTableId"`
}

// The SoftLayer_Billing_Item_Hardware_Security_Module data type contains general information relating to a single SoftLayer billing item for a hardware security module.
type Billing_Item_Hardware_Security_Module struct {
	Billing_Item_Hardware
}

// The SoftLayer_Billing_Item_Hardware_Server data type contains billing information about a bare metal server and its relationship to a particular customer account.
type Billing_Item_Hardware_Server struct {
	Billing_Item_Hardware
}

// no documentation yet
type Billing_Item_Link_ThePlanet struct {
	Entity

	// no documentation yet
	BillingItem *Billing_Item `json:"billingItem,omitempty" xmlrpc:"billingItem"`

	// no documentation yet
	ServiceProvider *Service_Provider `json:"serviceProvider,omitempty" xmlrpc:"serviceProvider"`
}

// The SoftLayer_Billing_Item_Network_Application_Delivery_Controller data type describes the billing item related to a NetScaler VPX
type Billing_Item_Network_Application_Delivery_Controller struct {
	Billing_Item

	// The bandwidth allotment detail for a billing item.
	BandwidthAllotmentDetail *Network_Bandwidth_Version1_Allotment_Detail `json:"bandwidthAllotmentDetail,omitempty" xmlrpc:"bandwidthAllotmentDetail"`

	// The network application controller that a billing item is associated with.
	Resource *Network_Application_Delivery_Controller `json:"resource,omitempty" xmlrpc:"resource"`
}

// A SoftLayer_Billing_Item_Network_Application_Delivery_Controller_LoadBalancer represents the [[SoftLayer_Billing_Item|billing item]] related to a single [[SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress|load balancer]] instance.
type Billing_Item_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress struct {
	Billing_Item

	// The load balancer that a load balancer billing item is associated with.
	Resource *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"resource,omitempty" xmlrpc:"resource"`
}

// The SoftLayer_Billing_Item_Hardware data type contains general information relating to a single SoftLayer billing item for hardware.
type Billing_Item_Network_Bandwidth struct {
	Billing_Item
}

// The SoftLayer_Billing_Item_Network_Firewall data type contains general information relating to a single SoftLayer billing item whose item category code is 'firewall'
type Billing_Item_Network_Firewall struct {
	Billing_Item

	// The VLAN firewall that a VLAN firewall billing item is associated with.
	Resource *Network_Component_Firewall `json:"resource,omitempty" xmlrpc:"resource"`
}

// The SoftLayer_Billing_Item_Network_Firewall_Module_Context data type describes the billing items related to VLAN Firewalls.
type Billing_Item_Network_Firewall_Module_Context struct {
	Billing_Item
}

// A SoftLayer_Billing_Item_Network_Interconnect represents the [[SoftLayer_Billing_Item|billing item]] related to a network interconnect instance.
type Billing_Item_Network_Interconnect struct {
	Billing_Item
}

// A SoftLayer_Billing_Item_Network_LoadBalancer represents the [[SoftLayer_Billing_Item|billing item]] related to a single [[SoftLayer_Network_LoadBalancer|load balancer]] instance.
type Billing_Item_Network_LoadBalancer struct {
	Billing_Item
}

// The SoftLayer_Billing_Item_Network_LoadBalancer_Global data type contains general information relating to a single SoftLayer billing item whose item category code is 'global_load_balancer'
type Billing_Item_Network_LoadBalancer_Global struct {
	Billing_Item

	// The resource for a global load balancer billing item.
	Resource *Network_LoadBalancer_Global_Account `json:"resource,omitempty" xmlrpc:"resource"`
}

// A SoftLayer_Billing_Item_Network_LoadBalancer_VirtualIpAddress represents the [[SoftLayer_Billing_Item|billing item]] related to a single [[SoftLayer_Network_LoadBalancer_VirtualIpAddress|load balancer]] instance.
type Billing_Item_Network_LoadBalancer_VirtualIpAddress struct {
	Billing_Item

	// The load balancer's virtual IP address that the billing item is associated with.
	Resource *Network_LoadBalancer_VirtualIpAddress `json:"resource,omitempty" xmlrpc:"resource"`
}

// The SoftLayer_Billing_Item_Network_Message_Delivery data describes the related billing item.
type Billing_Item_Network_Message_Delivery struct {
	Billing_Item

	// The object this billing item is associated with.
	Resource *Network_Message_Delivery `json:"resource,omitempty" xmlrpc:"resource"`
}

// The SoftLayer_Billing_Item_Network_Message_Queue data describes the related billing item.
type Billing_Item_Network_Message_Queue struct {
	Billing_Item

	// The object this billing item is associated with.
	Resource *Network_Message_Queue `json:"resource,omitempty" xmlrpc:"resource"`
}

// The SoftLayer_Billing_Item_Network_Message_Queue data describes the related billing item.
type Billing_Item_Network_Message_Queue_Delivery struct {
	Billing_Item_Network_Message_Queue
}

// The SoftLayer_Billing_Item_Network_PerformanceStorage_Iscsi data type contains general information relating to a single SoftLayer billing item whose item category code is 'performance_storage_iscsi'
type Billing_Item_Network_PerformanceStorage_Iscsi struct {
	Billing_Item_Network_Storage
}

// The SoftLayer_Billing_Item_Network_PerformanceStorage_Nfs data type contains general information relating to a single SoftLayer billing item whose item category code is 'performance_storage_nfs'
type Billing_Item_Network_PerformanceStorage_Nfs struct {
	Billing_Item_Network_Storage
}

// The SoftLayer_Billing_Item_Network_Storage data type describes the billing items related to StorageLayer accounts.
type Billing_Item_Network_Storage struct {
	Billing_Item

	// The StorageLayer account that a network storage billing item is associated with.
	Resource *Network_Storage `json:"resource,omitempty" xmlrpc:"resource"`
}

// The SoftLayer_Billing_Item_Network_Storage_Hub models all billing items related to hub-based StorageLayer offerings, such as CloudLayer storage.
type Billing_Item_Network_Storage_Hub struct {
	Billing_Item_Network_Storage
}

// The SoftLayer_Billing_Item_Network_Storage_Hub_Bandwidth data type models the billing items created when a CloudLayer storage account generates a bandwidth overage charge.
type Billing_Item_Network_Storage_Hub_Bandwidth struct {
	Billing_Item_Network_Storage
}

// The SoftLayer_Billing_Item_Network_Subnet data type contains general information relating to a single SoftLayer billing item whose item category code is one of the following:
// * pri_ip_address
// * static_sec_ip_addresses (static secondary)
// * sov_sec_ip_addresses (secondary on vlan, also known as "portable ips")
// * sov_sec_ip_addresses_pub (sov_sec_ip_addresses public only)
// * sov_sec_ip_addresses_priv (sov_sec_ip_addresses private only)
// * sec_ip_addresses (old style, secondary ip addresses)
//
//
// These item categories denote that the billing item has subnet information attached.
type Billing_Item_Network_Subnet struct {
	Billing_Item

	// The resource for a subnet-related billing item.
	Resource *Network_Subnet `json:"resource,omitempty" xmlrpc:"resource"`

	// The resource name for a subnet billing item.
	ResourceName *string `json:"resourceName,omitempty" xmlrpc:"resourceName"`

	// The resource (unique identifier) for a server billing item.
	ResourceTableId *int `json:"resourceTableId,omitempty" xmlrpc:"resourceTableId"`
}

// The SoftLayer_Billing_Item_Network_Subnet_IpAddress_Global data type contains general information relating to a single SoftLayer billing item whose item category code is one of the following:
// * global_ipv4
// * global_ipv6
//
//
// These item categories denote that the billing item has subnet information attached.
type Billing_Item_Network_Subnet_IpAddress_Global struct {
	Billing_Item_Network_Subnet
}

// The SoftLayer_Billing_Item_Network_Storage data type describes the billing items related to StorageLayer accounts.
type Billing_Item_Network_Tunnel struct {
	Billing_Item

	// The IPsec VPN that a network tunnel billing item is associated with.
	Resource *Network_Tunnel_Module_Context `json:"resource,omitempty" xmlrpc:"resource"`
}

// The SoftLayer_Billing_Item_Network_Vlant data type contains general information relating to a single SoftLayer billing item whose item category code is one of the following:
// * network_vlan
//
//
// These item categories denote that the billing item has network vlan information attached.
type Billing_Item_Network_Vlan struct {
	Billing_Item

	// The resource for a network vlan related billing item.
	Resource *Network_Vlan `json:"resource,omitempty" xmlrpc:"resource"`
}

// The SoftLayer_Billing_Item_Hardware data type contains general information relating to a single SoftLayer billing item for hardware components.
type Billing_Item_Software_Component struct {
	Billing_Item

	// The software component that this billing item points to.
	Resource *Software_Component `json:"resource,omitempty" xmlrpc:"resource"`

	// The resource (unique identifier) for a software component billing item.
	ResourceTableId *int `json:"resourceTableId,omitempty" xmlrpc:"resourceTableId"`
}

// The SoftLayer_Billing_Item_Software_Component_Analytics_Urchin data type contains general information relating to a single SoftLayer billing item for Urchin software components.
type Billing_Item_Software_Component_Analytics_Urchin struct {
	Billing_Item
}

// The SoftLayer_Billing_Item_Software_Component_ControlPanel data type contains general information relating to a single SoftLayer billing item for control panel software components.
type Billing_Item_Software_Component_ControlPanel struct {
	Billing_Item
}

// The SoftLayer_Billing_Item_Software_Component_ControlPanel data type contains general information relating to a single SoftLayer billing item for control panel software components.
type Billing_Item_Software_Component_ControlPanel_Parallels_Plesk_Billing struct {
	Billing_Item
}

// The SoftLayer_Billing_Item_Software_Component_OperatingSystem_Addon data type contains general information relating to a single SoftLayer billing item for operating system add-on software components.
type Billing_Item_Software_Component_OperatingSystem_Addon struct {
	Billing_Item
}

// The SoftLayer_Billing_Item_Software_Component_OperatingSystem_Addon_Citrix_Essentials data type contains general information relating to a single SoftLayer billing item for Citrix Essentials software components.
type Billing_Item_Software_Component_OperatingSystem_Addon_Citrix_Essentials struct {
	Billing_Item_Software_Component_OperatingSystem_Addon

	// The Citrix Essentials software component that a billing item is associated with.
	Resource *Software_Component `json:"resource,omitempty" xmlrpc:"resource"`
}

// The SoftLayer_Billing_Item_Software_Component_Virtual_OperatingSystem data type contains general information relating to a single SoftLayer billing item for operating system software components on virtual machines.
type Billing_Item_Software_Component_Virtual_OperatingSystem struct {
	Billing_Item
}

// The SoftLayer_Billing_Item_Software_Component_Virtual_OperatingSystem_Microsoft data type contains general information relating to a single SoftLayer billing item for a Microsoft operating system software components on virtual machines.
type Billing_Item_Software_Component_Virtual_OperatingSystem_Microsoft struct {
	Billing_Item_Software_Component_Virtual_OperatingSystem

	// The software virtual license to which this billing item points.
	Resource *Software_VirtualLicense `json:"resource,omitempty" xmlrpc:"resource"`

	// The resource (unique identifier) for a software virtual license billing item.
	ResourceTableId *int `json:"resourceTableId,omitempty" xmlrpc:"resourceTableId"`
}

// The SoftLayer_Billing_Item_Software_Component_Virtual_OperatingSystem_Microsoft data type contains general information relating to a single SoftLayer billing item for a Microsoft operating system software components on virtual machines.
type Billing_Item_Software_Component_Virtual_OperatingSystem_Redhat struct {
	Billing_Item_Software_Component_Virtual_OperatingSystem

	// The software component to which this billing item points.
	Resource *Software_Component `json:"resource,omitempty" xmlrpc:"resource"`

	// The resource (unique identifier) for a software component billing item.
	ResourceTableId *int `json:"resourceTableId,omitempty" xmlrpc:"resourceTableId"`
}

// The SoftLayer_Billing_Item_Software_License data type contains general information relating to a single SoftLayer billing item for a software license.
type Billing_Item_Software_License struct {
	Billing_Item

	// The resource for a software license billing item.
	Resource *Software_AccountLicense `json:"resource,omitempty" xmlrpc:"resource"`
}

// The SoftLayer_Billing_Item_Support data type contains general information relating to a premium support offering
type Billing_Item_Support struct {
	Billing_Item
}

// The SoftLayer_Billing_Item_Network_Application_Delivery_Controller data type describes the billing item related to an external authentication binding
type Billing_Item_User_Customer_External_Binding struct {
	Billing_Item

	// The external authentication binding that a billing item is associated with.
	Resource *User_Customer_External_Binding `json:"resource,omitempty" xmlrpc:"resource"`
}

// A SoftLayer_Billing_Item_Virtual_Dedicated_Rack data type models the billing information for a single bandwidth pooling. Bandwidth pooling members share their public bandwidth allocations, and incur overage charges instead of the overages on individual rack members. Virtual rack billing items are the parent items for all of it's rack membership billing items.
type Billing_Item_Virtual_Dedicated_Rack struct {
	Billing_Item

	// The raw bandwidth usage data for the current billing cycle. One object is returned for each network a virtual rack is attached to.
	BillingCycleBandwidthUsage []Network_Bandwidth_Usage `json:"billingCycleBandwidthUsage,omitempty" xmlrpc:"billingCycleBandwidthUsage"`

	// A count of the raw bandwidth usage data for the current billing cycle. One object is returned for each network a virtual rack is attached to.
	BillingCycleBandwidthUsageCount *uint `json:"billingCycleBandwidthUsageCount,omitempty" xmlrpc:"billingCycleBandwidthUsageCount"`

	// The raw private bandwidth usage data for the current billing cycle.
	BillingCyclePrivateBandwidthUsage []Network_Bandwidth_Usage `json:"billingCyclePrivateBandwidthUsage,omitempty" xmlrpc:"billingCyclePrivateBandwidthUsage"`

	// A count of the raw private bandwidth usage data for the current billing cycle.
	BillingCyclePrivateBandwidthUsageCount *uint `json:"billingCyclePrivateBandwidthUsageCount,omitempty" xmlrpc:"billingCyclePrivateBandwidthUsageCount"`

	// The total private network inbound bandwidth for this virtual rack for the current billing cycle.
	BillingCyclePrivateUsageIn *Float64 `json:"billingCyclePrivateUsageIn,omitempty" xmlrpc:"billingCyclePrivateUsageIn"`

	// The total private network outbound bandwidth for this virtual rack for the current billing cycle.
	BillingCyclePrivateUsageOut *Float64 `json:"billingCyclePrivateUsageOut,omitempty" xmlrpc:"billingCyclePrivateUsageOut"`

	// The total private network bandwidth for this virtual rack for the current billing cycle.
	BillingCyclePrivateUsageTotal *uint `json:"billingCyclePrivateUsageTotal,omitempty" xmlrpc:"billingCyclePrivateUsageTotal"`

	// The raw public bandwidth usage data for the current billing cycle.
	BillingCyclePublicBandwidthUsage []Network_Bandwidth_Usage `json:"billingCyclePublicBandwidthUsage,omitempty" xmlrpc:"billingCyclePublicBandwidthUsage"`

	// A count of the raw public bandwidth usage data for the current billing cycle.
	BillingCyclePublicBandwidthUsageCount *uint `json:"billingCyclePublicBandwidthUsageCount,omitempty" xmlrpc:"billingCyclePublicBandwidthUsageCount"`

	// The total public inbound bandwidth for this virtual rack for the current billing cycle.
	BillingCyclePublicUsageIn *Float64 `json:"billingCyclePublicUsageIn,omitempty" xmlrpc:"billingCyclePublicUsageIn"`

	// The total public outbound bandwidth for this virtual rack for the current billing cycle.
	BillingCyclePublicUsageOut *Float64 `json:"billingCyclePublicUsageOut,omitempty" xmlrpc:"billingCyclePublicUsageOut"`

	// The total public bandwidth for this virtual rack for the current billing cycle.
	BillingCyclePublicUsageTotal *uint `json:"billingCyclePublicUsageTotal,omitempty" xmlrpc:"billingCyclePublicUsageTotal"`

	// The virtual rack that a virtual rack billing item is associated with.
	Resource *Network_Bandwidth_Version1_Allotment `json:"resource,omitempty" xmlrpc:"resource"`
}

// The SoftLayer_Billing_Item_Virtual_Disk_Image data type contains general information relating to a single SoftLayer billing item for disk images.
type Billing_Item_Virtual_Disk_Image struct {
	Billing_Item

	// The disk image to which the billing item points.
	Resource *Virtual_Disk_Image `json:"resource,omitempty" xmlrpc:"resource"`

	// The resource (unique identifier) for a disk image billing item.
	ResourceTableId *int `json:"resourceTableId,omitempty" xmlrpc:"resourceTableId"`
}

// The SoftLayer_Billing_Item_Virtual_Guest data type contains general information relating to a single SoftLayer billing item for guests.
type Billing_Item_Virtual_Guest struct {
	Billing_Item

	// The raw bandwidth usage data for the current billing cycle. One object will be returned for each network this server is attached to.
	BillingCycleBandwidthUsage []Network_Bandwidth_Usage `json:"billingCycleBandwidthUsage,omitempty" xmlrpc:"billingCycleBandwidthUsage"`

	// A count of the raw bandwidth usage data for the current billing cycle. One object will be returned for each network this server is attached to.
	BillingCycleBandwidthUsageCount *uint `json:"billingCycleBandwidthUsageCount,omitempty" xmlrpc:"billingCycleBandwidthUsageCount"`

	// The raw private bandwidth usage data for the current billing cycle.
	BillingCyclePrivateBandwidthUsage []Network_Bandwidth_Usage `json:"billingCyclePrivateBandwidthUsage,omitempty" xmlrpc:"billingCyclePrivateBandwidthUsage"`

	// A count of the raw private bandwidth usage data for the current billing cycle.
	BillingCyclePrivateBandwidthUsageCount *uint `json:"billingCyclePrivateBandwidthUsageCount,omitempty" xmlrpc:"billingCyclePrivateBandwidthUsageCount"`

	// The total private inbound bandwidth for this virtual server for the current billing cycle.
	BillingCyclePrivateUsageIn *Float64 `json:"billingCyclePrivateUsageIn,omitempty" xmlrpc:"billingCyclePrivateUsageIn"`

	// The total private outbound bandwidth for this virtual server for the current billing cycle.
	BillingCyclePrivateUsageOut *Float64 `json:"billingCyclePrivateUsageOut,omitempty" xmlrpc:"billingCyclePrivateUsageOut"`

	// The total private bandwidth for this virtual server for the current billing cycle.
	BillingCyclePrivateUsageTotal *uint `json:"billingCyclePrivateUsageTotal,omitempty" xmlrpc:"billingCyclePrivateUsageTotal"`

	// The raw public bandwidth usage data for the current billing cycle.
	BillingCyclePublicBandwidthUsage []Network_Bandwidth_Usage `json:"billingCyclePublicBandwidthUsage,omitempty" xmlrpc:"billingCyclePublicBandwidthUsage"`

	// A count of the raw public bandwidth usage data for the current billing cycle.
	BillingCyclePublicBandwidthUsageCount *uint `json:"billingCyclePublicBandwidthUsageCount,omitempty" xmlrpc:"billingCyclePublicBandwidthUsageCount"`

	// The total public inbound bandwidth for this virtual server for the current billing cycle.
	BillingCyclePublicUsageIn *Float64 `json:"billingCyclePublicUsageIn,omitempty" xmlrpc:"billingCyclePublicUsageIn"`

	// The total public outbound bandwidth for this virtual server for the current billing cycle.
	BillingCyclePublicUsageOut *Float64 `json:"billingCyclePublicUsageOut,omitempty" xmlrpc:"billingCyclePublicUsageOut"`

	// The total public bandwidth for this virtual server for the current billing cycle.
	BillingCyclePublicUsageTotal *uint `json:"billingCyclePublicUsageTotal,omitempty" xmlrpc:"billingCyclePublicUsageTotal"`

	// A count of
	MonitoringBillingItemCount *uint `json:"monitoringBillingItemCount,omitempty" xmlrpc:"monitoringBillingItemCount"`

	// no documentation yet
	MonitoringBillingItems []Billing_Item `json:"monitoringBillingItems,omitempty" xmlrpc:"monitoringBillingItems"`

	// The resource for a cloud server billing item.
	Resource *Virtual_Guest `json:"resource,omitempty" xmlrpc:"resource"`

	// The resource (unique identifier) for a server billing item.
	ResourceTableId *int `json:"resourceTableId,omitempty" xmlrpc:"resourceTableId"`
}

// The SoftLayer_Billing_Item_Virtual_Host_Usage data type contains general information relating to a single SoftLayer billing item for virtual machine peak usage
type Billing_Item_Virtual_Host_Usage struct {
	Billing_Item

	// The resource for a peak virtual machine usage billing item.
	Resource *Hardware `json:"resource,omitempty" xmlrpc:"resource"`

	// The resource (unique identifier) for a server billing item.
	ResourceTableId *int `json:"resourceTableId,omitempty" xmlrpc:"resourceTableId"`
}

// The SoftLayer_Billing_Item_Workspace data type contains general information relating to a single SoftLayer billing item whose item category code is 'workspace'
type Billing_Item_Workspace struct {
	Billing_Item
}

// The SoftLayer_Billing_Order data type contains general information relating to an individual order applied to a SoftLayer customer account or to a new customer. Personal information in this type such as names, addresses, and phone numbers are taken from the account's contact information at the time the order is generated for existing SoftLayer customer.
type Billing_Order struct {
	Entity

	// The [[SoftLayer_Account|account]] to which an order belongs.
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// The account ID to which an order belongs.
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// no documentation yet
	Brand *Brand `json:"brand,omitempty" xmlrpc:"brand"`

	// A cart is similar to a quote, except that it can be continually modified by the customer and does not have locked-in prices. Not all orders will have a cart associated with them. See [[SoftLayer_Billing_Order_Cart]] for more information.
	Cart *Billing_Order_Cart `json:"cart,omitempty" xmlrpc:"cart"`

	// A count of the [[SoftLayer_Billing_Order_Item (type)|order items]] that are core restricted
	CoreRestrictedItemCount *uint `json:"coreRestrictedItemCount,omitempty" xmlrpc:"coreRestrictedItemCount"`

	// The [[SoftLayer_Billing_Order_Item (type)|order items]] that are core restricted
	CoreRestrictedItems []Billing_Order_Item `json:"coreRestrictedItems,omitempty" xmlrpc:"coreRestrictedItems"`

	// The point in time at which a billing item was created.
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// A count of all credit card transactions associated with this order. If this order was not placed with a credit card, this will be empty.
	CreditCardTransactionCount *uint `json:"creditCardTransactionCount,omitempty" xmlrpc:"creditCardTransactionCount"`

	// All credit card transactions associated with this order. If this order was not placed with a credit card, this will be empty.
	CreditCardTransactions []Billing_Payment_Card_Transaction `json:"creditCardTransactions,omitempty" xmlrpc:"creditCardTransactions"`

	// no documentation yet
	ExchangeRate *Billing_Currency_ExchangeRate `json:"exchangeRate,omitempty" xmlrpc:"exchangeRate"`

	// *
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// The SoftLayer_User_Customer id of the portal or API user who impersonated the user which submitted an order.
	ImpersonatingUserRecordId *int `json:"impersonatingUserRecordId,omitempty" xmlrpc:"impersonatingUserRecordId"`

	// no documentation yet
	InitialInvoice *Billing_Invoice `json:"initialInvoice,omitempty" xmlrpc:"initialInvoice"`

	// A count of the SoftLayer_Billing_Order_items included in an order.
	ItemCount *uint `json:"itemCount,omitempty" xmlrpc:"itemCount"`

	// The SoftLayer_Billing_Order_items included in an order.
	Items []Billing_Order_Item `json:"items,omitempty" xmlrpc:"items"`

	// The last time an order was updated.
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// no documentation yet
	OrderApprovalDate *Time `json:"orderApprovalDate,omitempty" xmlrpc:"orderApprovalDate"`

	// An order's non-server items total monthly fee.
	OrderNonServerMonthlyAmount *Float64 `json:"orderNonServerMonthlyAmount,omitempty" xmlrpc:"orderNonServerMonthlyAmount"`

	// The SoftLayer_Billing_Order_Quote id of the quote's user who finalized an order.
	OrderQuoteId *int `json:"orderQuoteId,omitempty" xmlrpc:"orderQuoteId"`

	// An order's server items total monthly fee.
	OrderServerMonthlyAmount *Float64 `json:"orderServerMonthlyAmount,omitempty" xmlrpc:"orderServerMonthlyAmount"`

	// A count of an order's top level items. This normally includes the server line item and any non-server additional services such as NAS or ISCSI.
	OrderTopLevelItemCount *uint `json:"orderTopLevelItemCount,omitempty" xmlrpc:"orderTopLevelItemCount"`

	// An order's top level items. This normally includes the server line item and any non-server additional services such as NAS or ISCSI.
	OrderTopLevelItems []Billing_Order_Item `json:"orderTopLevelItems,omitempty" xmlrpc:"orderTopLevelItems"`

	// This amount represents the order's initial charge including set up fee and taxes.
	OrderTotalAmount *Float64 `json:"orderTotalAmount,omitempty" xmlrpc:"orderTotalAmount"`

	// An order's total one time amount summing all the set up fees, the labor fees and the one time fees. Taxes will be applied for non-tax-exempt. This amount represents the initial fees that will be charged.
	OrderTotalOneTime *Float64 `json:"orderTotalOneTime,omitempty" xmlrpc:"orderTotalOneTime"`

	// An order's total one time amount. This amount represents the initial fees before tax.
	OrderTotalOneTimeAmount *Float64 `json:"orderTotalOneTimeAmount,omitempty" xmlrpc:"orderTotalOneTimeAmount"`

	// An order's total one time tax amount. This amount represents the tax that will be applied to the total charge, if the SoftLayer_Account tied to a SoftLayer_Billing_Order is a taxable account.
	OrderTotalOneTimeTaxAmount *Float64 `json:"orderTotalOneTimeTaxAmount,omitempty" xmlrpc:"orderTotalOneTimeTaxAmount"`

	// An order's total recurring amount. Taxes will be applied for non-tax-exempt. This amount represents the fees that will be charged on a recurring (usually monthly) basis.
	OrderTotalRecurring *Float64 `json:"orderTotalRecurring,omitempty" xmlrpc:"orderTotalRecurring"`

	// An order's total recurring amount. This amount represents the fees that will be charged on a recurring (usually monthly) basis.
	OrderTotalRecurringAmount *Float64 `json:"orderTotalRecurringAmount,omitempty" xmlrpc:"orderTotalRecurringAmount"`

	// The total tax amount of the recurring fees, if the SoftLayer_Account tied to a SoftLayer_Billing_Order is a taxable account.
	OrderTotalRecurringTaxAmount *Float64 `json:"orderTotalRecurringTaxAmount,omitempty" xmlrpc:"orderTotalRecurringTaxAmount"`

	// An order's total setup fee.
	OrderTotalSetupAmount *Float64 `json:"orderTotalSetupAmount,omitempty" xmlrpc:"orderTotalSetupAmount"`

	// The type of an order. This lets you know where this order was generated from.
	OrderType *Billing_Order_Type `json:"orderType,omitempty" xmlrpc:"orderType"`

	// The SoftLayer_Billing_Order_Type id of the order.
	OrderTypeId *int `json:"orderTypeId,omitempty" xmlrpc:"orderTypeId"`

	// A count of all PayPal transactions associated with this order. If this order was not placed with PayPal, this will be empty.
	PaypalTransactionCount *uint `json:"paypalTransactionCount,omitempty" xmlrpc:"paypalTransactionCount"`

	// All PayPal transactions associated with this order. If this order was not placed with PayPal, this will be empty.
	PaypalTransactions []Billing_Payment_PayPal_Transaction `json:"paypalTransactions,omitempty" xmlrpc:"paypalTransactions"`

	// no documentation yet
	PresaleEvent *Sales_Presale_Event `json:"presaleEvent,omitempty" xmlrpc:"presaleEvent"`

	// no documentation yet
	PresaleEventId *int `json:"presaleEventId,omitempty" xmlrpc:"presaleEventId"`

	// Flag indicating a private cloud solution order (Deprecated)
	PrivateCloudOrderFlag *bool `json:"privateCloudOrderFlag,omitempty" xmlrpc:"privateCloudOrderFlag"`

	// The quote of an order. This quote holds information about its expiration date, creation date, name and status. This information is tied to an order having the status 'QUOTE'
	Quote *Billing_Order_Quote `json:"quote,omitempty" xmlrpc:"quote"`

	// The Referral Partner who referred this order. (Only necessary for new customer orders)
	ReferralPartner *Account `json:"referralPartner,omitempty" xmlrpc:"referralPartner"`

	// Purchaser current status e.i. Approved, Pending_Approval
	Status *string `json:"status,omitempty" xmlrpc:"status"`

	// This flag indicates an order is an upgrade.
	UpgradeRequestFlag *bool `json:"upgradeRequestFlag,omitempty" xmlrpc:"upgradeRequestFlag"`

	// The SoftLayer_User_Customer object tied to an order.
	UserRecord *User_Customer `json:"userRecord,omitempty" xmlrpc:"userRecord"`

	// The SoftLayer_User_Customer id of the portal or API user who submitted an order.
	UserRecordId *int `json:"userRecordId,omitempty" xmlrpc:"userRecordId"`
}

// no documentation yet
type Billing_Order_Cart struct {
	Billing_Order_Quote
}

// Every individual item that a SoftLayer customer is billed for is recorded in the SoftLayer_Billing_Item data type. Billing items range from server chassis to hard drives to control panels, bandwidth quota upgrades and port upgrade charges. Softlayer [[SoftLayer_Billing_Invoice|invoices]] are generated from the cost of a customer's billing items. Billing items are copied from the product catalog as they're ordered by customers to create a reference between an account and the billable items they own.
//
// Billing items exist in a tree relationship. Items are associated with each other by parent/child relationships. Component items such as CPU's, RAM, and software each have a parent billing item for the server chassis they're associated with. Billing Items with a null parent item do not have an associated parent item.
type Billing_Order_Item struct {
	Entity

	// The SoftLayer_Billing_Item tied to the order item.
	BillingItem *Billing_Item `json:"billingItem,omitempty" xmlrpc:"billingItem"`

	// A count of the other items included with an ordered item.
	BundledItemCount *uint `json:"bundledItemCount,omitempty" xmlrpc:"bundledItemCount"`

	// The other items included with an ordered item.
	BundledItems []Billing_Order_Item `json:"bundledItems,omitempty" xmlrpc:"bundledItems"`

	// The item category tied to an order item.
	Category *Product_Item_Category `json:"category,omitempty" xmlrpc:"category"`

	// The category code for the order item.
	CategoryCode *string `json:"categoryCode,omitempty" xmlrpc:"categoryCode"`

	// The child order items for an order item. All server order items should have children. These children are considered a part of the server.
	Children []Billing_Order_Item `json:"children,omitempty" xmlrpc:"children"`

	// A count of the child order items for an order item. All server order items should have children. These children are considered a part of the server.
	ChildrenCount *uint `json:"childrenCount,omitempty" xmlrpc:"childrenCount"`

	// friendly description of purchase item.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// The domain name of the server as designated by the purchaser at the time of order placement.
	DomainName *string `json:"domainName,omitempty" xmlrpc:"domainName"`

	// A hardware's universally unique identifier.
	GlobalIdentifier *string `json:"globalIdentifier,omitempty" xmlrpc:"globalIdentifier"`

	// The component type tied to an order item. All hardware-specific items should have a generic hardware component.
	HardwareGenericComponent *Hardware_Component_Model_Generic `json:"hardwareGenericComponent,omitempty" xmlrpc:"hardwareGenericComponent"`

	// The hostname of the server as designated by the purchaser at the time of order placement.
	HostName *string `json:"hostName,omitempty" xmlrpc:"hostName"`

	// The amount of money charged per hourly for an order item, if applicable, and only if it was ordered this day. hourlyRecurringFee is measured in US Dollars ($USD).
	HourlyRecurringFee *Float64 `json:"hourlyRecurringFee,omitempty" xmlrpc:"hourlyRecurringFee"`

	// no documentation yet
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// The SoftLayer_Product_Item tied to an order item. The item is the actual definition of the product being sold.
	Item *Product_Item `json:"item,omitempty" xmlrpc:"item"`

	// A count of this is an item's category answers.
	ItemCategoryAnswerCount *uint `json:"itemCategoryAnswerCount,omitempty" xmlrpc:"itemCategoryAnswerCount"`

	// This is an item's category answers.
	ItemCategoryAnswers []Billing_Order_Item_Category_Answer `json:"itemCategoryAnswers,omitempty" xmlrpc:"itemCategoryAnswers"`

	// The SoftLayer_Product_Item ID for this order item.
	ItemId *int `json:"itemId,omitempty" xmlrpc:"itemId"`

	// The SoftLayer_Product_Item_Price tied to an order item. The item price object describes the cost of an item.
	ItemPrice *Product_Item_Price `json:"itemPrice,omitempty" xmlrpc:"itemPrice"`

	// the item price id (SoftLayer_Product_Item_Price->id) of the ordered item.
	ItemPriceId *Float64 `json:"itemPriceId,omitempty" xmlrpc:"itemPriceId"`

	// An order item's labor fee total after taxes. This does not include any child invoice items.
	LaborAfterTaxAmount *Float64 `json:"laborAfterTaxAmount,omitempty" xmlrpc:"laborAfterTaxAmount"`

	// The labor fee, if any. This is a one time charge.
	LaborFee *Float64 `json:"laborFee,omitempty" xmlrpc:"laborFee"`

	// The rate at which labor fees are taxed if you are a taxable customer.
	LaborFeeTaxRate *Float64 `json:"laborFeeTaxRate,omitempty" xmlrpc:"laborFeeTaxRate"`

	// An order item's labor tax amount. This does not include any child invoice items.
	LaborTaxAmount *Float64 `json:"laborTaxAmount,omitempty" xmlrpc:"laborTaxAmount"`

	// The location of an ordered item. This is usually the same as the server it is being ordered with. Otherwise it describes the location of the additional service being ordered.
	Location *Location `json:"location,omitempty" xmlrpc:"location"`

	// no documentation yet
	NextOrderChildren []Billing_Order_Item `json:"nextOrderChildren,omitempty" xmlrpc:"nextOrderChildren"`

	// A count of
	NextOrderChildrenCount *uint `json:"nextOrderChildrenCount,omitempty" xmlrpc:"nextOrderChildrenCount"`

	// This is only populated when an upgrade order is placed. The old billing item represents what the billing was before the upgrade happened.
	OldBillingItem *Billing_Item `json:"oldBillingItem,omitempty" xmlrpc:"oldBillingItem"`

	// An order item's one-time fee total after taxes. This does not include any child invoice items.
	OneTimeAfterTaxAmount *Float64 `json:"oneTimeAfterTaxAmount,omitempty" xmlrpc:"oneTimeAfterTaxAmount"`

	// The amount of money charged as a one-time charge for an order item, if applicable. oneTimeFee is measured in US Dollars ($USD).
	OneTimeFee *Float64 `json:"oneTimeFee,omitempty" xmlrpc:"oneTimeFee"`

	// The rate at which one time fees are taxed if you are a taxable customer.
	OneTimeFeeTaxRate *Float64 `json:"oneTimeFeeTaxRate,omitempty" xmlrpc:"oneTimeFeeTaxRate"`

	// An order item's one-time tax amount. This does not include any child invoice items.
	OneTimeTaxAmount *Float64 `json:"oneTimeTaxAmount,omitempty" xmlrpc:"oneTimeTaxAmount"`

	// The order to which this item belongs. The order contains all the information related to the items included in an order
	Order *Billing_Order `json:"order,omitempty" xmlrpc:"order"`

	// no documentation yet
	OrderApprovalDate *Time `json:"orderApprovalDate,omitempty" xmlrpc:"orderApprovalDate"`

	// The SoftLayer_Product_Package an order item is a part of.
	Package *Product_Package `json:"package,omitempty" xmlrpc:"package"`

	// The parent order item ID for an item. Items that are associated with a server will have a parent. The parent will be the server item itself.
	Parent *Billing_Order_Item `json:"parent,omitempty" xmlrpc:"parent"`

	// no documentation yet
	ParentId *int `json:"parentId,omitempty" xmlrpc:"parentId"`

	// no documentation yet
	PromoCodeId *int `json:"promoCodeId,omitempty" xmlrpc:"promoCodeId"`

	// the quantity of the ordered item in a quote.
	Quantity *int `json:"quantity,omitempty" xmlrpc:"quantity"`

	// An order item's recurring fee total after taxes. This does not include any child invoice items.
	RecurringAfterTaxAmount *Float64 `json:"recurringAfterTaxAmount,omitempty" xmlrpc:"recurringAfterTaxAmount"`

	// The amount of money charged per month for an order item, if applicable. recurringFee is measured in US Dollars ($USD).
	RecurringFee *Float64 `json:"recurringFee,omitempty" xmlrpc:"recurringFee"`

	// An order item's recurring tax amount. This does not include any child invoice items.
	RecurringTaxAmount *Float64 `json:"recurringTaxAmount,omitempty" xmlrpc:"recurringTaxAmount"`

	// A count of power supplies contained within this SoftLayer_Billing_Order
	RedundantPowerSupplyCount *uint `json:"redundantPowerSupplyCount,omitempty" xmlrpc:"redundantPowerSupplyCount"`

	// An order item's setup fee total after taxes. This does not include any child invoice items.
	SetupAfterTaxAmount *Float64 `json:"setupAfterTaxAmount,omitempty" xmlrpc:"setupAfterTaxAmount"`

	// The setup fee, if any. This is a one time charge.
	SetupFee *Float64 `json:"setupFee,omitempty" xmlrpc:"setupFee"`

	// The month set up fee deferral.
	SetupFeeDeferralMonths *int `json:"setupFeeDeferralMonths,omitempty" xmlrpc:"setupFeeDeferralMonths"`

	// The rate at which setup fees are taxed if you are a taxable customer.
	SetupFeeTaxRate *Float64 `json:"setupFeeTaxRate,omitempty" xmlrpc:"setupFeeTaxRate"`

	// An order item's setup tax amount. This does not include any child invoice items.
	SetupTaxAmount *Float64 `json:"setupTaxAmount,omitempty" xmlrpc:"setupTaxAmount"`

	// For ordered items that are software items, a full description of that software can be found with this property.
	SoftwareDescription *Software_Description `json:"softwareDescription,omitempty" xmlrpc:"softwareDescription"`

	// A count of the drive storage groups that are attached to this billing order item.
	StorageGroupCount *uint `json:"storageGroupCount,omitempty" xmlrpc:"storageGroupCount"`

	// The drive storage groups that are attached to this billing order item.
	StorageGroups []Configuration_Storage_Group_Order `json:"storageGroups,omitempty" xmlrpc:"storageGroups"`

	// The recurring fee of an ordered item. This amount represents the fees that will be charged on a recurring (usually monthly) basis.
	TotalRecurringAmount *Float64 `json:"totalRecurringAmount,omitempty" xmlrpc:"totalRecurringAmount"`

	// The next SoftLayer_Product_Item in the upgrade path for this order item.
	UpgradeItem *Product_Item `json:"upgradeItem,omitempty" xmlrpc:"upgradeItem"`
}

// The SoftLayer_Billing_Order_Item_Category_Answer data type represents a single answer to an item category question.
type Billing_Order_Item_Category_Answer struct {
	Entity

	// The answer to the question.
	Answer *string `json:"answer,omitempty" xmlrpc:"answer"`

	// The date that the answer was created.
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// The billing order item that the answer is for.
	OrderItem *Billing_Order_Item `json:"orderItem,omitempty" xmlrpc:"orderItem"`

	// The question that is being answered.
	Question *Product_Item_Category_Question `json:"question,omitempty" xmlrpc:"question"`

	// The identifier for the question that the answer belongs to.
	QuestionId *int `json:"questionId,omitempty" xmlrpc:"questionId"`
}

// no documentation yet
type Billing_Order_Note struct {
	Entity

	// no documentation yet
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// no documentation yet
	Employee *User_Employee `json:"employee,omitempty" xmlrpc:"employee"`

	// no documentation yet
	Order *Billing_Order `json:"order,omitempty" xmlrpc:"order"`
}

// The SoftLayer_Billing_Oder_Quote data type contains general information relating to an individual order applied to a SoftLayer customer account or to a new customer. Personal information in this type such as names, addresses, and phone numbers are taken from the account's contact information at the time the quote is generated for existing SoftLayer customer.
type Billing_Order_Quote struct {
	Entity

	// A quote's corresponding account.
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// Identification Number of the account record tied to the quote
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// Identification Number of the order record tied to the quote.
	CompletedPurchaseDataId *int `json:"completedPurchaseDataId,omitempty" xmlrpc:"completedPurchaseDataId"`

	// Holds the date the quote record was created
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// This property holds the date of expiration of a quote, after that date the quote would be deem expired
	ExpirationDate *Time `json:"expirationDate,omitempty" xmlrpc:"expirationDate"`

	// The id use to identify a quote.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// Holds the date when the quote record was modified with reference to its creation date
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// The name given to quote by the initiator
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// This order contains the records for which products were selected for this quote.
	Order *Billing_Order `json:"order,omitempty" xmlrpc:"order"`

	// These are all the orders that were created from this quote.
	OrdersFromQuote []Billing_Order `json:"ordersFromQuote,omitempty" xmlrpc:"ordersFromQuote"`

	// A count of these are all the orders that were created from this quote.
	OrdersFromQuoteCount *uint `json:"ordersFromQuoteCount,omitempty" xmlrpc:"ordersFromQuoteCount"`

	// This property Holds system generated notes. In our case if a quote is tied to an order where one of the order item has an inactive promotion code, the quote will be considered invalid.
	PublicNote *string `json:"publicNote,omitempty" xmlrpc:"publicNote"`

	// Holds system generated hash password for the Quote
	QuoteKey *string `json:"quoteKey,omitempty" xmlrpc:"quoteKey"`

	// This property Holds the current status of a Quote: pending,expired, saved or deleted
	Status *string `json:"status,omitempty" xmlrpc:"status"`
}

// The SoftLayer_Billing_Oder_Type data type contains general information relating to all the different types of orders that exist. This data pertains only to where an order was generated from, from any of the SoftLayer websites with ordering interfaces or directly through the SoftLayer API.
type Billing_Order_Type struct {
	Entity

	// A brief description of where a SoftLayer order originated from.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// A SoftLayer order type's internal identifier.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// A simple keyname stating where a SoftLayer order originated from.
	Type *string `json:"type,omitempty" xmlrpc:"type"`
}

// The SoftLayer_Billing_Payment_Card_ChangeRequest data type contains general information relating to attempted credit card information changes.
type Billing_Payment_Card_ChangeRequest struct {
	Entity

	// no documentation yet
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// The account ID to which the credit card and billing information is associated with.
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// The total amount of the attempted transaction, represented in decimal format as US Dollars ($USD).
	Amount *Float64 `json:"amount,omitempty" xmlrpc:"amount"`

	// The SoftLayer_Billing_Payment_Card_Transaction tied to the authorization performed as part of this change request.
	AuthorizedCreditCardTransaction *Billing_Payment_Card_Transaction `json:"authorizedCreditCardTransaction,omitempty" xmlrpc:"authorizedCreditCardTransaction"`

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

	// The SoftLayer_Billing_Payment_Card_Transaction tied to the capture of funds performed as part of this change request.
	CaptureCreditCardTransaction *Billing_Payment_Card_Transaction `json:"captureCreditCardTransaction,omitempty" xmlrpc:"captureCreditCardTransaction"`

	// The last 4 digits of a customer's credit card.
	CardAccountLast4 *string `json:"cardAccountLast4,omitempty" xmlrpc:"cardAccountLast4"`

	// The card number submitted in the change request.
	CardAccountNumber *string `json:"cardAccountNumber,omitempty" xmlrpc:"cardAccountNumber"`

	// The month (MM) in which a customer's payment card will expire.
	CardExpirationMonth *string `json:"cardExpirationMonth,omitempty" xmlrpc:"cardExpirationMonth"`

	// The year (YYYY) in which a customer's payment card will expire.
	CardExpirationYear *string `json:"cardExpirationYear,omitempty" xmlrpc:"cardExpirationYear"`

	// no documentation yet
	CardNickname *string `json:"cardNickname,omitempty" xmlrpc:"cardNickname"`

	// The type of payment card a customer has. (i.e. Visa, MasterCard, American Express).
	CardType *string `json:"cardType,omitempty" xmlrpc:"cardType"`

	// The credit card verification number submitted in the change request.
	CreditCardVerificationNumber *string `json:"creditCardVerificationNumber,omitempty" xmlrpc:"creditCardVerificationNumber"`

	// Describes the currency selected for payment
	CurrencyShortName *string `json:"currencyShortName,omitempty" xmlrpc:"currencyShortName"`

	// Device Fingerprint Identifier - Used internally and can safely be ignored.
	DeviceFingerprintId *string `json:"deviceFingerprintId,omitempty" xmlrpc:"deviceFingerprintId"`

	// The unique identifier for a single change request.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// the notes stored about a customer's payment card.
	Notes *string `json:"notes,omitempty" xmlrpc:"notes"`

	// no documentation yet
	PaymentRoleId *int `json:"paymentRoleId,omitempty" xmlrpc:"paymentRoleId"`

	// The description of the type of payment sent in a change transaction.
	PaymentType *string `json:"paymentType,omitempty" xmlrpc:"paymentType"`

	// A count of these are tickets tied to a credit card change request.
	TicketAttachmentReferenceCount *uint `json:"ticketAttachmentReferenceCount,omitempty" xmlrpc:"ticketAttachmentReferenceCount"`

	// These are tickets tied to a credit card change request.
	TicketAttachmentReferences []Ticket_Attachment `json:"ticketAttachmentReferences,omitempty" xmlrpc:"ticketAttachmentReferences"`

	// Unique identifier for a ticket discussing the switch between payment methods.
	TicketId *int `json:"ticketId,omitempty" xmlrpc:"ticketId"`
}

// The SoftLayer_Billing_Payment_Card_ManualPayment data type contains general information relating to attempted credit card information changes.
type Billing_Payment_Card_ManualPayment struct {
	Entity

	// no documentation yet
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// The account ID to which the credit card and billing information is associated with.
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// The total amount of the attempted transaction, represented in decimal format as US Dollars ($USD).
	Amount *Float64 `json:"amount,omitempty" xmlrpc:"amount"`

	// This is the credit card transaction data tied to a credit card manual payment.
	AuthorizedCreditCardTransaction *Billing_Payment_Card_Transaction `json:"authorizedCreditCardTransaction,omitempty" xmlrpc:"authorizedCreditCardTransaction"`

	// The unique identifier of an attempted credit card transaction.
	AuthorizedCreditCardTransactionId *int `json:"authorizedCreditCardTransactionId,omitempty" xmlrpc:"authorizedCreditCardTransactionId"`

	// This is the PayPal transaction data tied to a PayPal manual payment.
	AuthorizedPayPalTransaction *Billing_Payment_PayPal_Transaction `json:"authorizedPayPalTransaction,omitempty" xmlrpc:"authorizedPayPalTransaction"`

	// The unique identifier of an attempted PayPal transaction.
	AuthorizedPayPalTransactionId *int `json:"authorizedPayPalTransactionId,omitempty" xmlrpc:"authorizedPayPalTransactionId"`

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

	// The last name of the customer account owner.
	BillingNameLast *string `json:"billingNameLast,omitempty" xmlrpc:"billingNameLast"`

	// The fax number associated with a customer account.
	BillingPhoneFax *string `json:"billingPhoneFax,omitempty" xmlrpc:"billingPhoneFax"`

	// The phone number associated with a customer account.
	BillingPhoneVoice *string `json:"billingPhoneVoice,omitempty" xmlrpc:"billingPhoneVoice"`

	// The Zip or Postal Code for the billing address on an account.
	BillingPostalCode *string `json:"billingPostalCode,omitempty" xmlrpc:"billingPostalCode"`

	// The State for the account.
	BillingState *string `json:"billingState,omitempty" xmlrpc:"billingState"`

	// The cancel URL is the page to which PayPal redirects if payment is not approved.
	CancelUrl *string `json:"cancelUrl,omitempty" xmlrpc:"cancelUrl"`

	// The SoftLayer_Billing_Payment_Card_Transaction tied to the capture performed as part of this manual payment. This will only exist if the manual payment was performed with a credit card.
	CaptureCreditCardTransaction *Billing_Payment_Card_Transaction `json:"captureCreditCardTransaction,omitempty" xmlrpc:"captureCreditCardTransaction"`

	// The SoftLayer_Billing_Payment_PayPal_Transaction tied to the capture performed as part of this manual payment. This will only exist if the manual payment was performed via PayPal.
	CapturePayPalTransaction *Billing_Payment_PayPal_Transaction `json:"capturePayPalTransaction,omitempty" xmlrpc:"capturePayPalTransaction"`

	// A hash value of the credit card number.
	CardAccountHash *string `json:"cardAccountHash,omitempty" xmlrpc:"cardAccountHash"`

	// The last 4 digits of a customer's credit card.
	CardAccountLast4 *string `json:"cardAccountLast4,omitempty" xmlrpc:"cardAccountLast4"`

	// The card number submitted in the change request.
	CardAccountNumber *string `json:"cardAccountNumber,omitempty" xmlrpc:"cardAccountNumber"`

	// The month (MM) in which a customer's payment card will expire.
	CardExpirationMonth *string `json:"cardExpirationMonth,omitempty" xmlrpc:"cardExpirationMonth"`

	// The year (YYYY) in which a customer's payment card will expire.
	CardExpirationYear *string `json:"cardExpirationYear,omitempty" xmlrpc:"cardExpirationYear"`

	// The method key of the type payment issued (Visa - 001, Mastercard - 002, American Express - 003, Discover - 004, PayPal - paypal).
	CardType *string `json:"cardType,omitempty" xmlrpc:"cardType"`

	// The credit card verification number submitted in the change request.
	CreditCardVerificationNumber *string `json:"creditCardVerificationNumber,omitempty" xmlrpc:"creditCardVerificationNumber"`

	// Describes the currency selected for payment
	CurrencyShortName *string `json:"currencyShortName,omitempty" xmlrpc:"currencyShortName"`

	// Device Fingerprint Identifier - Used internally and can safely be ignored.
	DeviceFingerprintId *string `json:"deviceFingerprintId,omitempty" xmlrpc:"deviceFingerprintId"`

	// The IP address from which the transaction originates.
	FromIpAddress *string `json:"fromIpAddress,omitempty" xmlrpc:"fromIpAddress"`

	// The unique identifier for a single manual payment request.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// Notes generated as a result of the payment request.
	Notes *string `json:"notes,omitempty" xmlrpc:"notes"`

	// The description of the type of payment sent in a change transaction.
	PaymentType *string `json:"paymentType,omitempty" xmlrpc:"paymentType"`

	// The return URL is the page to which PayPal redirects after payment is approved.
	ReturnUrl *string `json:"returnUrl,omitempty" xmlrpc:"returnUrl"`

	// A count of these are tickets tied to a credit card manual payment.
	TicketAttachmentReferenceCount *uint `json:"ticketAttachmentReferenceCount,omitempty" xmlrpc:"ticketAttachmentReferenceCount"`

	// These are tickets tied to a credit card manual payment.
	TicketAttachmentReferences []Ticket_Attachment `json:"ticketAttachmentReferences,omitempty" xmlrpc:"ticketAttachmentReferences"`

	// Describes the type of manual payment (PAYPAL or CREDIT_CARD).
	Type *string `json:"type,omitempty" xmlrpc:"type"`
}

// The SoftLayer_Billing_Payment_Card_Transaction data type contains general information relating to attempted credit card transactions.
type Billing_Payment_Card_Transaction struct {
	Billing_Payment_Transaction

	// The account to which a transaction belongs.
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// The account ID to which the credit card and billing information is associated with.
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// The total amount of the attempted transaction, represented in decimal format as US Dollars ($USD).
	Amount *Float64 `json:"amount,omitempty" xmlrpc:"amount"`

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

	// The last name of the customer account owner.
	BillingNameLast *string `json:"billingNameLast,omitempty" xmlrpc:"billingNameLast"`

	// The fax number associated with a customer account.
	BillingPhoneFax *string `json:"billingPhoneFax,omitempty" xmlrpc:"billingPhoneFax"`

	// The phone number associated with a customer account.
	BillingPhoneVoice *string `json:"billingPhoneVoice,omitempty" xmlrpc:"billingPhoneVoice"`

	// The Zip or Postal Code for the billing address on an account.
	BillingPostalCode *string `json:"billingPostalCode,omitempty" xmlrpc:"billingPostalCode"`

	// The State for the account.
	BillingState *string `json:"billingState,omitempty" xmlrpc:"billingState"`

	// The last 4 digits of a customer's credit card.
	CardAccountLast4 *int `json:"cardAccountLast4,omitempty" xmlrpc:"cardAccountLast4"`

	// The month (MM) in which a customer's payment card will expire.
	CardExpirationMonth *int `json:"cardExpirationMonth,omitempty" xmlrpc:"cardExpirationMonth"`

	// The year (YYYY) in which a customer's payment card will expire.
	CardExpirationYear *int `json:"cardExpirationYear,omitempty" xmlrpc:"cardExpirationYear"`

	// The type of payment issued (i.e. Visa, MasterCard, American Express).
	CardType *string `json:"cardType,omitempty" xmlrpc:"cardType"`

	// The date that the transaction was attempted.
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// The unique identifier for a single credit card transaction request.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// Unique identifier of the invoice to which funds will be applied.
	InvoiceId *int `json:"invoiceId,omitempty" xmlrpc:"invoiceId"`

	// The date that the transaction was modified.
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// no documentation yet
	Order *Billing_Order `json:"order,omitempty" xmlrpc:"order"`

	// The IP address from which the transaction originates.
	OrderFromIpAddress *string `json:"orderFromIpAddress,omitempty" xmlrpc:"orderFromIpAddress"`

	// A code used by the financial institution to refer to the requested transaction.
	ReferenceCode *string `json:"referenceCode,omitempty" xmlrpc:"referenceCode"`

	// The unique identifier of the request submitted to the financial institution.
	RequestId *string `json:"requestId,omitempty" xmlrpc:"requestId"`

	// The status code returned from the financial institution.
	ReturnStatus *int `json:"returnStatus,omitempty" xmlrpc:"returnStatus"`

	// A serialized, delimited string of the transaction request sent to the financial institution.
	SerializedReply *string `json:"serializedReply,omitempty" xmlrpc:"serializedReply"`

	// A serialized, delimited string of the transaction request sent to the financial institution.
	SerializedRequest *string `json:"serializedRequest,omitempty" xmlrpc:"serializedRequest"`
}

// The SoftLayer_Billing_Payment_PayPal_Transaction data type contains general information relating to attempted PayPal transactions.
type Billing_Payment_PayPal_Transaction struct {
	Billing_Payment_Transaction

	// The account to which a transaction belongs.
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// The account ID to which the PayPal and billing information is associated with.
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// City given in the address of the PayPal user.
	AddressCityName *string `json:"addressCityName,omitempty" xmlrpc:"addressCityName"`

	// Country given in the named address of the PayPal user.
	AddressCountry *string `json:"addressCountry,omitempty" xmlrpc:"addressCountry"`

	// Name given to the address provided for the PayPal user.
	AddressName *string `json:"addressName,omitempty" xmlrpc:"addressName"`

	// Postal Code of the address of the PayPal user.
	AddressPostalCode *string `json:"addressPostalCode,omitempty" xmlrpc:"addressPostalCode"`

	// State or Province in the address of the PayPal user.
	AddressStateProvence *string `json:"addressStateProvence,omitempty" xmlrpc:"addressStateProvence"`

	// PayPal defined status of the address of the PayPal user.
	AddressStatus *string `json:"addressStatus,omitempty" xmlrpc:"addressStatus"`

	// First line of the street address of the PayPal user.
	AddressStreet1 *string `json:"addressStreet1,omitempty" xmlrpc:"addressStreet1"`

	// Second line of the street address of the PayPal user.
	AddressStreet2 *string `json:"addressStreet2,omitempty" xmlrpc:"addressStreet2"`

	// Phone number provided for the PayPal user.
	ContactPhone *string `json:"contactPhone,omitempty" xmlrpc:"contactPhone"`

	// The date that the transaction was attempted.
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// Exchange rate imposed on the payment amount.
	ExchangeRate *string `json:"exchangeRate,omitempty" xmlrpc:"exchangeRate"`

	// PayPal fee applied to the payment.
	FeeAmount *Float64 `json:"feeAmount,omitempty" xmlrpc:"feeAmount"`

	// The total amount of the payment executed by PayPal, represented in decimal format as US Dollars ($USD).
	GrossAmount *Float64 `json:"grossAmount,omitempty" xmlrpc:"grossAmount"`

	// The unique identifier for a single PayPal transaction request.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// Unique identifier of the invoice to which funds will be applied.
	InvoiceId *int `json:"invoiceId,omitempty" xmlrpc:"invoiceId"`

	// The name of the command issued to PayPal with regards to the attempted transaction.
	LastPaypalCommand *string `json:"lastPaypalCommand,omitempty" xmlrpc:"lastPaypalCommand"`

	// The date that the transaction was modified.
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// no documentation yet
	Order *Billing_Order `json:"order,omitempty" xmlrpc:"order"`

	// The IP address from where the PayPal payment request originated.
	OrderFromIpAddress *string `json:"orderFromIpAddress,omitempty" xmlrpc:"orderFromIpAddress"`

	// The amount of the payment submitted through the SoftLayer interface, represented in decimal format as US Dollars ($USD).
	OrderTotal *Float64 `json:"orderTotal,omitempty" xmlrpc:"orderTotal"`

	// The PayPal user account name (email address) associated with the customer account.
	Payer *string `json:"payer,omitempty" xmlrpc:"payer"`

	// The name of the business associated with the PayPal user.
	PayerBusiness *string `json:"payerBusiness,omitempty" xmlrpc:"payerBusiness"`

	// Country given in the address of the PayPal user.
	PayerCountry *string `json:"payerCountry,omitempty" xmlrpc:"payerCountry"`

	// First name of the PayPal user.
	PayerFirstName *string `json:"payerFirstName,omitempty" xmlrpc:"payerFirstName"`

	// Unique PayPal user account identifier.
	PayerId *string `json:"payerId,omitempty" xmlrpc:"payerId"`

	// Last name of the PayPal user.
	PayerLastName *string `json:"payerLastName,omitempty" xmlrpc:"payerLastName"`

	// Current PayPal status associated with the user account.
	PayerStatus *string `json:"payerStatus,omitempty" xmlrpc:"payerStatus"`

	// Date that the payment was confirmed in PayPal by the user.
	PaymentDate *Time `json:"paymentDate,omitempty" xmlrpc:"paymentDate"`

	// PayPal defined status of the attempted payment.
	PaymentStatus *string `json:"paymentStatus,omitempty" xmlrpc:"paymentStatus"`

	// PayPal defined code used to identify the type of payment.  Provided in a PayPal response.
	PaymentType *string `json:"paymentType,omitempty" xmlrpc:"paymentType"`

	// Reason provided by PayPal for a payment given a pending status.
	PendingReason *string `json:"pendingReason,omitempty" xmlrpc:"pendingReason"`

	// A serialized, delimited string of the reply received from PayPal.
	SerializedReply *string `json:"serializedReply,omitempty" xmlrpc:"serializedReply"`

	// A serialized, delimited string of the request submitted to PayPal.
	SerializedRequest *string `json:"serializedRequest,omitempty" xmlrpc:"serializedRequest"`

	// PayPal defined fee.
	SettleAmount *Float64 `json:"settleAmount,omitempty" xmlrpc:"settleAmount"`

	// Tax applied by PayPal to the payment amount.
	TaxAmount *Float64 `json:"taxAmount,omitempty" xmlrpc:"taxAmount"`

	// Value issued by PayPal for referencing the attempted transaction.
	Token *string `json:"token,omitempty" xmlrpc:"token"`

	// Unique transaction ID provided in a PayPal response.
	TransactionId *string `json:"transactionId,omitempty" xmlrpc:"transactionId"`

	// PayPal defined code used to identify the type of transaction.  Provided in a PayPal response.
	TransactionType *string `json:"transactionType,omitempty" xmlrpc:"transactionType"`
}

// no documentation yet
type Billing_Payment_Processor struct {
	Entity

	// A count of
	BrandAssignmentCount *uint `json:"brandAssignmentCount,omitempty" xmlrpc:"brandAssignmentCount"`

	// no documentation yet
	BrandAssignments []Brand_Payment_Processor `json:"brandAssignments,omitempty" xmlrpc:"brandAssignments"`

	// no documentation yet
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// no documentation yet
	OwnerAccount *Account `json:"ownerAccount,omitempty" xmlrpc:"ownerAccount"`

	// A count of
	PaymentMethodCount *uint `json:"paymentMethodCount,omitempty" xmlrpc:"paymentMethodCount"`

	// no documentation yet
	PaymentMethods []Billing_Payment_Processor_Method `json:"paymentMethods,omitempty" xmlrpc:"paymentMethods"`

	// no documentation yet
	Type *Billing_Payment_Processor_Type `json:"type,omitempty" xmlrpc:"type"`
}

// no documentation yet
type Billing_Payment_Processor_Method struct {
	Entity

	// no documentation yet
	MethodKey *string `json:"methodKey,omitempty" xmlrpc:"methodKey"`

	// no documentation yet
	MultipleCurrencyFlag *bool `json:"multipleCurrencyFlag,omitempty" xmlrpc:"multipleCurrencyFlag"`

	// no documentation yet
	PaymentProcessor *Billing_Payment_Processor `json:"paymentProcessor,omitempty" xmlrpc:"paymentProcessor"`

	// no documentation yet
	PaymentType *Billing_Payment_Type `json:"paymentType,omitempty" xmlrpc:"paymentType"`
}

// no documentation yet
type Billing_Payment_Processor_Type struct {
	Entity

	// no documentation yet
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// no documentation yet
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// A count of
	PaymentProcessorCount *uint `json:"paymentProcessorCount,omitempty" xmlrpc:"paymentProcessorCount"`

	// no documentation yet
	PaymentProcessors []Billing_Payment_Processor `json:"paymentProcessors,omitempty" xmlrpc:"paymentProcessors"`
}

// Implementation for payment transactions.
type Billing_Payment_Transaction struct {
	Entity
}

// no documentation yet
type Billing_Payment_Type struct {
	Entity

	// no documentation yet
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// no documentation yet
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}
