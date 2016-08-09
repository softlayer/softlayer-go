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

package services

import "github.ibm.com/riethm/gopherlayer/datatypes"

//
type Billing_Currency struct {
	Session *Session
	Options
}

func (r *Session) GetBillingCurrencyService() Billing_Currency {
	return Billing_Currency{Session: r}
}

//
func (r *Billing_Currency) GetAllObjects() (resp []datatypes.Billing_Currency, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Billing_Currency) GetObject() (resp datatypes.Billing_Currency, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Billing_Currency) GetPrice(price *float64, formatOptions *datatypes.Container_Billing_Currency_Format) (resp string, err error) {
	params := []interface{}{
		price,
		formatOptions,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
type Billing_Currency_ExchangeRate struct {
	Session *Session
	Options
}

func (r *Session) GetBillingCurrencyExchangeRateService() Billing_Currency_ExchangeRate {
	return Billing_Currency_ExchangeRate{Session: r}
}

//
func (r *Billing_Currency_ExchangeRate) GetAllCurrencyExchangeRates(stringDate *string) (resp []datatypes.Billing_Currency_ExchangeRate, err error) {
	params := []interface{}{
		stringDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Billing_Currency_ExchangeRate) GetCurrencies() (resp []datatypes.Billing_Currency, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Billing_Currency_ExchangeRate) GetExchangeRate(to *string, from *string, effectiveDate *datatypes.Time) (resp datatypes.Billing_Currency_ExchangeRate, err error) {
	params := []interface{}{
		to,
		from,
		effectiveDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Billing_Currency_ExchangeRate) GetFundingCurrency() (resp datatypes.Billing_Currency, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Billing_Currency_ExchangeRate) GetLocalCurrency() (resp datatypes.Billing_Currency, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Billing_Currency_ExchangeRate) GetObject() (resp datatypes.Billing_Currency_ExchangeRate, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Billing_Currency_ExchangeRate) GetPrice(price *float64, formatOptions *datatypes.Container_Billing_Currency_Format) (resp string, err error) {
	params := []interface{}{
		price,
		formatOptions,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Every SoftLayer customer account has billing specific information which is kept in the SoftLayer_Billing_Info data type. This information is used by the SoftLayer accounting group when sending invoices and making billing inquiries.
type Billing_Info struct {
	Session *Session
	Options
}

func (r *Session) GetBillingInfoService() Billing_Info {
	return Billing_Info{Session: r}
}

// Retrieve The SoftLayer customer account associated with this billing information.
func (r *Billing_Info) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Billing_Info) GetAchInformation() (resp []datatypes.Billing_Info_Ach, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Currency to be used by this customer account.
func (r *Billing_Info) GetCurrency() (resp datatypes.Billing_Currency, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Information related to an account's current and previous billing cycles.
func (r *Billing_Info) GetCurrentBillingCycle() (resp datatypes.Billing_Info_Cycle, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The date on which an account was last billed.
func (r *Billing_Info) GetLastBillDate() (resp datatypes.Time, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The date on which an account will be billed next.
func (r *Billing_Info) GetNextBillDate() (resp datatypes.Time, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Billing_Info object whose data corresponds to the account to which your portal user is tied.
func (r *Billing_Info) GetObject() (resp datatypes.Billing_Info, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The SoftLayer_Billing_Invoice data type contains general information relating to an individual invoice applied to a SoftLayer customer account. Personal information in this type such as names, addresses, and phone numbers are taken from the account's contact information at the time the invoice is generated.
type Billing_Invoice struct {
	Session *Session
	Options
}

func (r *Session) GetBillingInvoiceService() Billing_Invoice {
	return Billing_Invoice{Session: r}
}

// Create a transaction to email PDF and/or Excel invoice links to the requesting user's email address. You must have a PDF reader installed in order to view these files.
func (r *Billing_Invoice) EmailInvoices(options *datatypes.Container_Billing_Invoice_Email) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		options,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The account that an invoice belongs to.
func (r *Billing_Invoice) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve This is the amount of this invoice.
func (r *Billing_Invoice) GetAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Billing_Invoice) GetBrandAtInvoiceCreation() (resp datatypes.Brand, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A flag that will reflect whether the detailed version of the pdf has been generated.
func (r *Billing_Invoice) GetDetailedPdfGeneratedFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a Microsoft Excel spreadsheet of a SoftLayer invoice. You must have a Microsoft Excel reader installed in order to view these invoice files.
func (r *Billing_Invoice) GetExcel() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A list of top-level invoice items that are on the currently pending invoice.
func (r *Billing_Invoice) GetInvoiceTopLevelItems() (resp []datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total amount of this invoice.
func (r *Billing_Invoice) GetInvoiceTotalAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total one-time charges for this invoice. This is the sum of one-time charges + setup fees + labor fees. This does not include taxes.
func (r *Billing_Invoice) GetInvoiceTotalOneTimeAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A sum of all the taxes related to one time charges for this invoice.
func (r *Billing_Invoice) GetInvoiceTotalOneTimeTaxAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total amount of this invoice. This does not include taxes.
func (r *Billing_Invoice) GetInvoiceTotalPreTaxAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total Recurring amount of this invoice. This amount does not include taxes or one time charges.
func (r *Billing_Invoice) GetInvoiceTotalRecurringAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total amount of the recurring taxes on this invoice.
func (r *Billing_Invoice) GetInvoiceTotalRecurringTaxAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The items that belong to this invoice.
func (r *Billing_Invoice) GetItems() (resp []datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Billing_Invoice object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Billing_Invoice service. You can only retrieve invoices that are assigned to your portal user's account.
func (r *Billing_Invoice) GetObject() (resp datatypes.Billing_Invoice, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve This is the total payment made on this invoice.
func (r *Billing_Invoice) GetPayment() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The payments for the invoice.
func (r *Billing_Invoice) GetPayments() (resp []datatypes.Billing_Invoice_Receivable_Payment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a PDF record of a SoftLayer invoice. SoftLayer keeps PDF records of all closed invoices for customer retrieval from the portal and API. You must have a PDF reader installed in order to view these invoice files.
func (r *Billing_Invoice) GetPdf() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a PDF record of a SoftLayer detailed invoice summary. SoftLayer keeps PDF records of all closed invoices for customer retrieval from the portal and API. You must have a PDF reader installed in order to view these files.
func (r *Billing_Invoice) GetPdfDetailed() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Billing_Invoice) GetPdfDetailedFilename() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the size of a PDF record of a SoftLayer invoice. SoftLayer keeps PDF records of all closed invoices for customer retrieval from the portal and API.
func (r *Billing_Invoice) GetPdfFileSize() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Billing_Invoice) GetPdfFilename() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a Microsoft Excel record of a SoftLayer invoice. SoftLayer generates Microsoft Excel records of all closed invoices for customer retrieval from the portal and API. You must have a Microsoft Excel reader installed in order to view these invoice files.
func (r *Billing_Invoice) GetPreliminaryExcel() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a PDF record of a SoftLayer invoice. SoftLayer keeps PDF records of all closed invoices for customer retrieval from the portal and API. You must have a PDF reader installed in order to view these invoice files.
func (r *Billing_Invoice) GetPreliminaryPdf() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a PDF record of the detailed version of a SoftLayer invoice. SoftLayer keeps PDF records of all closed invoices for customer retrieval from the portal and API.
func (r *Billing_Invoice) GetPreliminaryPdfDetailed() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve This is the seller's tax registration.
func (r *Billing_Invoice) GetSellerRegistration() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve This is the tax information that applies to tax auditing. This is the official tax record for this invoice.
func (r *Billing_Invoice) GetTaxInfo() (resp datatypes.Billing_Invoice_Tax_Info, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve This is the set of tax information for any tax calculation for this invoice. Note that not all of these are necessarily official, so use the taxInfo key to get the final information.
func (r *Billing_Invoice) GetTaxInfoHistory() (resp []datatypes.Billing_Invoice_Tax_Info, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve This is a message explaining the tax treatment for this invoice.
func (r *Billing_Invoice) GetTaxMessage() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve This is the strategy used to calculate tax on this invoice.
func (r *Billing_Invoice) GetTaxType() (resp datatypes.Billing_Invoice_Tax_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Billing_Invoice) GetXlsFilename() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Billing_Invoice) GetZeroFeeItemCounts() (resp []datatypes.Container_Product_Item_Category_ZeroFee_Count, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Each billing invoice item makes up a record within an invoice. This provides you with a detailed record of everything related to an invoice item. When you are billed, our system takes active billing items and creates an invoice. These invoice items are a copy of your active billing items, and make up the contents of your invoice.
type Billing_Invoice_Item struct {
	Session *Session
	Options
}

func (r *Session) GetBillingInvoiceItemService() Billing_Invoice_Item {
	return Billing_Invoice_Item{Session: r}
}

// Retrieve An Invoice Item's associated child invoice items. Only parent invoice items have associated children. For instance, a server invoice item may have associated children.
func (r *Billing_Invoice_Item) GetAssociatedChildren() (resp []datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An Invoice Item's associated invoice item. If this is populated, it means this is an orphaned invoice item, but logically belongs to the associated invoice item.
func (r *Billing_Invoice_Item) GetAssociatedInvoiceItem() (resp datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An Invoice Item's billing item, from which this item was generated.
func (r *Billing_Invoice_Item) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve This invoice item's "item category".
func (r *Billing_Invoice_Item) GetCategory() (resp datatypes.Product_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An Invoice Item's child invoice items. Only parent invoice items have children. For instance, a server invoice item will have children.
func (r *Billing_Invoice_Item) GetChildren() (resp []datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An Invoice Item's associated child invoice items, excluding some items with a $0.00 recurring fee. Only parent invoice items have associated children. For instance, a server invoice item may have associated children.
func (r *Billing_Invoice_Item) GetFilteredAssociatedChildren() (resp []datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The invoice to which this item belongs.
func (r *Billing_Invoice_Item) GetInvoice() (resp datatypes.Billing_Invoice, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An invoice item's location, if one exists.'
func (r *Billing_Invoice_Item) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An Invoice Item's associated child invoice items, excluding ALL items with a $0.00 recurring fee. Only parent invoice items have associated children. For instance, a server invoice item may have associated children.
func (r *Billing_Invoice_Item) GetNonZeroAssociatedChildren() (resp []datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Billing_Invoice_Item object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Billing_Invoice_Item service. You can only retrieve the items tied to the account that your portal user is assigned to.
func (r *Billing_Invoice_Item) GetObject() (resp datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Every item tied to a server should have a parent invoice item which is the server line item. This is how we associate items to a server.
func (r *Billing_Invoice_Item) GetParent() (resp datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The entry in the product catalog that a invoice item is based upon.
func (r *Billing_Invoice_Item) GetProduct() (resp datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An invoice Item's total, including any child invoice items if they exist.
func (r *Billing_Invoice_Item) GetTotalOneTimeAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An invoice Item's total, including any child invoice items if they exist.
func (r *Billing_Invoice_Item) GetTotalOneTimeTaxAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An invoice Item's total, including any child invoice items if they exist.
func (r *Billing_Invoice_Item) GetTotalRecurringAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A Billing Item's total, including any child billing items if they exist.'
func (r *Billing_Invoice_Item) GetTotalRecurringTaxAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Billing_Invoice_Next struct {
	Session *Session
	Options
}

func (r *Session) GetBillingInvoiceNextService() Billing_Invoice_Next {
	return Billing_Invoice_Next{Session: r}
}

// Return an account's next invoice in a Microsoft excel format.
func (r *Billing_Invoice_Next) GetExcel(documentCreateDate *datatypes.Time) (resp []byte, err error) {
	params := []interface{}{
		documentCreateDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Return an account's next invoice in PDF format.
func (r *Billing_Invoice_Next) GetPdf(documentCreateDate *datatypes.Time) (resp []byte, err error) {
	params := []interface{}{
		documentCreateDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Return an account's next invoice detailed portion in PDF format.
func (r *Billing_Invoice_Next) GetPdfDetailed(documentCreateDate *datatypes.Time) (resp []byte, err error) {
	params := []interface{}{
		documentCreateDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// The invoice tax status data type models a single status or state that an invoice can reflect in regard to an integration with a third-party tax calculation service.
type Billing_Invoice_Tax_Status struct {
	Session *Session
	Options
}

func (r *Session) GetBillingInvoiceTaxStatusService() Billing_Invoice_Tax_Status {
	return Billing_Invoice_Tax_Status{Session: r}
}

//
func (r *Billing_Invoice_Tax_Status) GetAllObjects() (resp []datatypes.Billing_Invoice_Tax_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Billing_Invoice_Tax_Status) GetObject() (resp datatypes.Billing_Invoice_Tax_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The invoice tax type data type models a single strategy for handling tax calculations.
type Billing_Invoice_Tax_Type struct {
	Session *Session
	Options
}

func (r *Session) GetBillingInvoiceTaxTypeService() Billing_Invoice_Tax_Type {
	return Billing_Invoice_Tax_Type{Session: r}
}

//
func (r *Billing_Invoice_Tax_Type) GetAllObjects() (resp []datatypes.Billing_Invoice_Tax_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Billing_Invoice_Tax_Type) GetObject() (resp datatypes.Billing_Invoice_Tax_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Every individual item that a SoftLayer customer is billed for is recorded in the SoftLayer_Billing_Item data type. Billing items range from server chassis to hard drives to control panels, bandwidth quota upgrades and port upgrade charges. Softlayer [[SoftLayer_Billing_Invoice|invoices]] are generated from the cost of a customer's billing items. Billing items are copied from the product catalog as they're ordered by customers to create a reference between an account and the billable items they own.
//
// Billing items exist in a tree relationship. Items are associated with each other by parent/child relationships. Component items such as CPU's, RAM, and software each have a parent billing item for the server chassis they're associated with. Billing Items with a null parent item do not have an associated parent item.
type Billing_Item struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemService() Billing_Item {
	return Billing_Item{Session: r}
}

// Cancel the resource or service for a billing Item. By default the billing item will be cancelled immediately and reclaim of the resource will begin shortly. Setting the "cancelImmediately" property to false will delay the cancellation until the next bill date.
//
//
// * The reason parameter could be from the list below:
// * "No longer needed"
// * "Business closing down"
// * "Server / Upgrade Costs"
// * "Migrating to larger server"
// * "Migrating to smaller server"
// * "Migrating to a different SoftLayer datacenter"
// * "Network performance / latency"
// * "Support response / timing"
// * "Sales process / upgrades"
// * "Moving to competitor"
func (r *Billing_Item) CancelItem(cancelImmediately *bool, cancelAssociatedBillingItems *bool, reason *string, customerNote *string) (resp bool, err error) {
	params := []interface{}{
		cancelImmediately,
		cancelAssociatedBillingItems,
		reason,
		customerNote,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Cancel the resource or service (excluding bare metal servers) for a billing Item. The billing item will be cancelled immediately and reclaim of the resource will begin shortly.
func (r *Billing_Item) CancelService() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Cancel the resource or service for a billing Item
func (r *Billing_Item) CancelServiceOnAnniversaryDate() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The account that a billing item belongs to.
func (r *Billing_Item) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Billing_Item) GetActiveAgreement() (resp datatypes.Account_Agreement, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A flag indicating that the billing item is under an active agreement.
func (r *Billing_Item) GetActiveAgreementFlag() (resp datatypes.Account_Agreement, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A billing item's active associated child billing items. This includes "floating" items that are not necessarily child items of this billing item.
func (r *Billing_Item) GetActiveAssociatedChildren() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Billing_Item) GetActiveAssociatedGuestDiskBillingItems() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A Billing Item's active bundled billing items.
func (r *Billing_Item) GetActiveBundledItems() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A service cancellation request item that corresponds to the billing item.
func (r *Billing_Item) GetActiveCancellationItem() (resp datatypes.Billing_Item_Cancellation_Request_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A Billing Item's active child billing items.
func (r *Billing_Item) GetActiveChildren() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Billing_Item) GetActiveFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Billing_Item) GetActiveSparePoolAssociatedGuestDiskBillingItems() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A Billing Item's spare pool bundled billing items.
func (r *Billing_Item) GetActiveSparePoolBundledItems() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A billing item's associated parent. This is to be used for billing items that are "floating", and therefore are not child items of any parent billing item. If it is desired to associate an item to another, populate this with the SoftLayer_Billing_Item ID of that associated parent item.
func (r *Billing_Item) GetAssociatedBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A history of billing items which a billing item has been associated with.
func (r *Billing_Item) GetAssociatedBillingItemHistory() (resp []datatypes.Billing_Item_Association_History, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A Billing Item's associated child billing items. This includes "floating" items that are not necessarily child billing items of this billing item.
func (r *Billing_Item) GetAssociatedChildren() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A billing item's associated parent billing item. This object will be the same as the parent billing item if parentId is set.
func (r *Billing_Item) GetAssociatedParent() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Billing_Item) GetAvailableMatchingVlans() (resp []datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The bandwidth allocation for a billing item.
func (r *Billing_Item) GetBandwidthAllocation() (resp datatypes.Network_Bandwidth_Version1_Allocation, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A billing item's recurring child items that have once been billed and are scheduled to be billed in the future.
func (r *Billing_Item) GetBillableChildren() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A Billing Item's bundled billing items
func (r *Billing_Item) GetBundleItems() (resp []datatypes.Product_Item_Bundles, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A Billing Item's bundled billing items'
func (r *Billing_Item) GetBundledItems() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A Billing Item's active child billing items.
func (r *Billing_Item) GetCanceledChildren() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The billing item's cancellation reason.
func (r *Billing_Item) GetCancellationReason() (resp datatypes.Billing_Item_Cancellation_Reason, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve This will return any cancellation requests that are associated with this billing item.
func (r *Billing_Item) GetCancellationRequests() (resp []datatypes.Billing_Item_Cancellation_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The item category to which the billing item's item belongs.
func (r *Billing_Item) GetCategory() (resp datatypes.Product_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A Billing Item's child billing items'
func (r *Billing_Item) GetChildren() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A Billing Item's active child billing items.
func (r *Billing_Item) GetChildrenWithActiveAgreement() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve For product items which have a downgrade path defined, this will return those product items.
func (r *Billing_Item) GetDowngradeItems() (resp []datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A Billing Item's associated child billing items, excluding some items with a $0.00 recurring fee.
func (r *Billing_Item) GetFilteredNextInvoiceChildren() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A flag that will reflect whether this billing item is billed on an hourly basis or not.
func (r *Billing_Item) GetHourlyFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Invoice items associated with this billing item
func (r *Billing_Item) GetInvoiceItem() (resp datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All invoice items associated with the billing item
func (r *Billing_Item) GetInvoiceItems() (resp []datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The entry in the SoftLayer product catalog that a billing item is based upon.
func (r *Billing_Item) GetItem() (resp datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The location of the billing item. Some billing items have physical properties such as the server itself. For items such as these, we provide location information.
func (r *Billing_Item) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A Billing Item's child billing items and associated items'
func (r *Billing_Item) GetNextInvoiceChildren() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A Billing Item's total, including any child billing items if they exist.'
func (r *Billing_Item) GetNextInvoiceTotalOneTimeAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A Billing Item's total, including any child billing items if they exist.'
func (r *Billing_Item) GetNextInvoiceTotalOneTimeTaxAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A Billing Item's total, including any child billing items and associated billing items if they exist.'
func (r *Billing_Item) GetNextInvoiceTotalRecurringAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve This is deprecated and will always be zero. Because tax is calculated in real-time, previewing the next recurring invoice is pre-tax only.
func (r *Billing_Item) GetNextInvoiceTotalRecurringTaxAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A Billing Item's associated child billing items, excluding ALL items with a $0.00 recurring fee.
func (r *Billing_Item) GetNonZeroNextInvoiceChildren() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Billing_Item object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Billing_Item service. You can only retrieve billing items tied to the account that your portal user is assigned to. Billing items are an account's items of billable items. There are "parent" billing items and "child" billing items. The server billing item is generally referred to as a parent billing item. The items tied to a server, such as ram, harddrives, and operating systems are considered "child" billing items.
func (r *Billing_Item) GetObject() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A billing item's original order item. Simply a reference to the original order from which this billing item was created.
func (r *Billing_Item) GetOrderItem() (resp datatypes.Billing_Order_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The original physical location for this billing item--may differ from current.
func (r *Billing_Item) GetOriginalLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The package under which this billing item was sold. A Package is the general grouping of products as seen on our order forms.
func (r *Billing_Item) GetPackage() (resp datatypes.Product_Package, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A billing item's parent item. If a billing item has no parent item then this value is null.
func (r *Billing_Item) GetParent() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A billing item's parent item. If a billing item has no parent item then this value is null.
func (r *Billing_Item) GetParentVirtualGuestBillingItem() (resp datatypes.Billing_Item_Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve This flag indicates whether a billing item is scheduled to be canceled or not.
func (r *Billing_Item) GetPendingCancellationFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The new order item that will replace this billing item.
func (r *Billing_Item) GetPendingOrderItem() (resp datatypes.Billing_Order_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Provisioning transaction for this billing item
func (r *Billing_Item) GetProvisionTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This service returns billing items of a specified category code. This service should be used to retrieve billing items that you wish to cancel. Some billing items can be canceled via [[SoftLayer_Security_Certificate_Request|service cancellation]] service.
//
// In order to find billing items for cancellation, use [[SoftLayer_Product_Item_Category::getValidCancelableServiceItemCategories|product categories]] service to retrieve category codes that are eligible for cancellation.
func (r *Billing_Item) GetServiceBillingItemsByCategory(categoryCode *string, includeZeroRecurringFee *bool) (resp []datatypes.Billing_Item, err error) {
	params := []interface{}{
		categoryCode,
		includeZeroRecurringFee,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A friendly description of software component
func (r *Billing_Item) GetSoftwareDescription() (resp datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Billing items whose product item has an upgrade path defined in our system will return the next product item in the upgrade path.
func (r *Billing_Item) GetUpgradeItem() (resp datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Billing items whose product item has an upgrade path defined in our system will return all the product items in the upgrade path.
func (r *Billing_Item) GetUpgradeItems() (resp []datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Remove the association from a billing item.
func (r *Billing_Item) RemoveAssociationId() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Set an associated billing item to an orphan billing item. Associations allow you to tie an "orphaned" billing item, any non-server billing item that doesn't have a parent item such as secondary IP subnets or StorageLayer accounts, to a server billing item. You may only set an association for an orphan to a server. You cannot associate a server to an orphan if the either the server or orphan billing items have a cancellation date set.
func (r *Billing_Item) SetAssociationId(associatedId *int) (resp bool, err error) {
	params := []interface{}{
		associatedId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Void a previously made cancellation for a service
func (r *Billing_Item) VoidCancelService() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The SoftLayer_Billing_Item_Cancellation_Reason data type contains cancellation reasons.
type Billing_Item_Cancellation_Reason struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemCancellationReasonService() Billing_Item_Cancellation_Reason {
	return Billing_Item_Cancellation_Reason{Session: r}
}

// getAllCancellationReasons() retrieves a list of all cancellation reasons that a server/service may be assigned to.
func (r *Billing_Item_Cancellation_Reason) GetAllCancellationReasons() (resp []datatypes.Billing_Item_Cancellation_Reason, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An billing cancellation reason category.
func (r *Billing_Item_Cancellation_Reason) GetBillingCancellationReasonCategory() (resp datatypes.Billing_Item_Cancellation_Reason_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The corresponding billing items having the specific cancellation reason.
func (r *Billing_Item_Cancellation_Reason) GetBillingItems() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Billing_Item_Cancellation_Reason) GetObject() (resp datatypes.Billing_Item_Cancellation_Reason, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Billing_Item_Cancellation_Reason) GetTranslatedReason() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The SoftLayer_Billing_Item_Cancellation_Reason_Category data type contains cancellation reason categories.
type Billing_Item_Cancellation_Reason_Category struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemCancellationReasonCategoryService() Billing_Item_Cancellation_Reason_Category {
	return Billing_Item_Cancellation_Reason_Category{Session: r}
}

// getAllCancellationReasonCategories() retrieves a list of all cancellation reason categories
func (r *Billing_Item_Cancellation_Reason_Category) GetAllCancellationReasonCategories() (resp []datatypes.Billing_Item_Cancellation_Reason_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The corresponding billing cancellation reasons having the specific billing cancellation reason category.
func (r *Billing_Item_Cancellation_Reason_Category) GetBillingCancellationReasons() (resp []datatypes.Billing_Item_Cancellation_Reason, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Billing_Item_Cancellation_Reason_Category) GetObject() (resp datatypes.Billing_Item_Cancellation_Reason_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// SoftLayer_Billing_Item_Cancellation_Request data type is used to cancel service billing items.
type Billing_Item_Cancellation_Request struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemCancellationRequestService() Billing_Item_Cancellation_Request {
	return Billing_Item_Cancellation_Request{Session: r}
}

// This method creates a service cancellation request.
//
// You need to have "Cancel Services" privilege to create a cancellation request. You have to provide at least one SoftLayer_Billing_Item_Cancellation_Request_Item in the "items" property. Make sure billing item's category code belongs to the cancelable product codes. You can retrieve the cancelable product category by the [[SoftLayer_Product_Item_Category::getValidCancelableServiceItemCategories|product category]] service.
func (r *Billing_Item_Cancellation_Request) CreateObject(templateObject *datatypes.Billing_Item_Cancellation_Request) (resp datatypes.Billing_Item_Cancellation_Request, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer account that a service cancellation request belongs to.
func (r *Billing_Item_Cancellation_Request) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This method returns all service cancellation requests.
//
// Make sure to include the "resultLimit" in the SOAP request header for quicker response. If there is no result limit header is passed, it will return the latest 25 results by default.
func (r *Billing_Item_Cancellation_Request) GetAllCancellationRequests() (resp []datatypes.Billing_Item_Cancellation_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Services can be canceled 2 or 3 days prior to your next bill date. This service returns the time by which a cancellation request submission is permitted in the current billing cycle. If the current time falls into the cut off date, this will return next earliest cancellation cut off date.
//
// Available category codes are: service, server
func (r *Billing_Item_Cancellation_Request) GetCancellationCutoffDate(accountId *int, categoryCode *string) (resp datatypes.Time, err error) {
	params := []interface{}{
		accountId,
		categoryCode,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve A collection of service cancellation items.
func (r *Billing_Item_Cancellation_Request) GetItems() (resp []datatypes.Billing_Item_Cancellation_Request_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Billing_Item_Cancellation_Request object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Billing_Item_Cancellation_Request service. You can only retrieve cancellation request records that are assigned to your SoftLayer account.
func (r *Billing_Item_Cancellation_Request) GetObject() (resp datatypes.Billing_Item_Cancellation_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The status of a service cancellation request.
func (r *Billing_Item_Cancellation_Request) GetStatus() (resp datatypes.Billing_Item_Cancellation_Request_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The ticket that is associated with the service cancellation request.
func (r *Billing_Item_Cancellation_Request) GetTicket() (resp datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The user that initiated a service cancellation request.
func (r *Billing_Item_Cancellation_Request) GetUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This method removes a cancellation item from a cancellation request that is in "Pending" or "Approved" status.
func (r *Billing_Item_Cancellation_Request) RemoveCancellationItem(itemId *int) (resp bool, err error) {
	params := []interface{}{
		itemId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// This method examined if a billing item is eligible for cancellation. It checks if the billing item you provided is already in your existing cancellation request.
func (r *Billing_Item_Cancellation_Request) ValidateBillingItemForCancellation(billingItemId *int) (resp bool, err error) {
	params := []interface{}{
		billingItemId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// This method voids a service cancellation request in "Pending" or "Approved" status.
func (r *Billing_Item_Cancellation_Request) Void(closeRelatedTicketFlag *bool) (resp bool, err error) {
	params := []interface{}{
		closeRelatedTicketFlag,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// The SoftLayer_Billing_Order data type contains general information relating to an individual order applied to a SoftLayer customer account or to a new customer. Personal information in this type such as names, addresses, and phone numbers are taken from the account's contact information at the time the order is generated for existing SoftLayer customer.
type Billing_Order struct {
	Session *Session
	Options
}

func (r *Session) GetBillingOrderService() Billing_Order {
	return Billing_Order{Session: r}
}

// When an order has been modified, the customer will need to approve the changes. This method will allow the customer to approve the changes.
func (r *Billing_Order) ApproveModifiedOrder() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The account to which an order belongs.
func (r *Billing_Order) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// This will get all billing orders for your account.
func (r *Billing_Order) GetAllObjects() (resp []datatypes.Billing_Order, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Billing_Order) GetBrand() (resp datatypes.Brand, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A cart is similar to a quote, except that it can be continually modified by the customer and does not have locked-in prices. Not all orders will have a cart associated with them. See [[SoftLayer_Billing_Order_Cart]] for more information.
func (r *Billing_Order) GetCart() (resp datatypes.Billing_Order_Cart, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The [[SoftLayer_Billing_Order_Item (type)|order items]] that are core restricted
func (r *Billing_Order) GetCoreRestrictedItems() (resp []datatypes.Billing_Order_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All credit card transactions associated with this order. If this order was not placed with a credit card, this will be empty.
func (r *Billing_Order) GetCreditCardTransactions() (resp []datatypes.Billing_Payment_Card_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Billing_Order) GetExchangeRate() (resp datatypes.Billing_Currency_ExchangeRate, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Billing_Order) GetInitialInvoice() (resp datatypes.Billing_Invoice, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Billing_Order_items included in an order.
func (r *Billing_Order) GetItems() (resp []datatypes.Billing_Order_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Billing_Order object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Billing_Order service. You can only retrieve orders that are assigned to your portal user's account.
func (r *Billing_Order) GetObject() (resp datatypes.Billing_Order, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Billing_Order) GetOrderApprovalDate() (resp datatypes.Time, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An order's non-server items total monthly fee.
func (r *Billing_Order) GetOrderNonServerMonthlyAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An order's server items total monthly fee.
func (r *Billing_Order) GetOrderServerMonthlyAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Get a list of [[SoftLayer_Container_Billing_Order_Status]] objects.
func (r *Billing_Order) GetOrderStatuses() (resp []datatypes.Container_Billing_Order_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An order's top level items. This normally includes the server line item and any non-server additional services such as NAS or ISCSI.
func (r *Billing_Order) GetOrderTopLevelItems() (resp []datatypes.Billing_Order_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve This amount represents the order's initial charge including set up fee and taxes.
func (r *Billing_Order) GetOrderTotalAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An order's total one time amount summing all the set up fees, the labor fees and the one time fees. Taxes will be applied for non-tax-exempt. This amount represents the initial fees that will be charged.
func (r *Billing_Order) GetOrderTotalOneTime() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An order's total one time amount. This amount represents the initial fees before tax.
func (r *Billing_Order) GetOrderTotalOneTimeAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An order's total one time tax amount. This amount represents the tax that will be applied to the total charge, if the SoftLayer_Account tied to a SoftLayer_Billing_Order is a taxable account.
func (r *Billing_Order) GetOrderTotalOneTimeTaxAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An order's total recurring amount. Taxes will be applied for non-tax-exempt. This amount represents the fees that will be charged on a recurring (usually monthly) basis.
func (r *Billing_Order) GetOrderTotalRecurring() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An order's total recurring amount. This amount represents the fees that will be charged on a recurring (usually monthly) basis.
func (r *Billing_Order) GetOrderTotalRecurringAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total tax amount of the recurring fees, if the SoftLayer_Account tied to a SoftLayer_Billing_Order is a taxable account.
func (r *Billing_Order) GetOrderTotalRecurringTaxAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve An order's total setup fee.
func (r *Billing_Order) GetOrderTotalSetupAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The type of an order. This lets you know where this order was generated from.
func (r *Billing_Order) GetOrderType() (resp datatypes.Billing_Order_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve All PayPal transactions associated with this order. If this order was not placed with PayPal, this will be empty.
func (r *Billing_Order) GetPaypalTransactions() (resp []datatypes.Billing_Payment_PayPal_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a PDF record of a SoftLayer quote. If the order is not a quote, an error will be thrown.
func (r *Billing_Order) GetPdf() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve the default filename of an order PDF.
func (r *Billing_Order) GetPdfFilename() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Billing_Order) GetPresaleEvent() (resp datatypes.Sales_Presale_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The quote of an order. This quote holds information about its expiration date, creation date, name and status. This information is tied to an order having the status 'QUOTE'
func (r *Billing_Order) GetQuote() (resp datatypes.Billing_Order_Quote, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Generate an [[SoftLayer_Container_Product_Order|order container]] from a billing order. This will take into account promotions, reseller status, estimated taxes and all other standard order verification processes.
func (r *Billing_Order) GetRecalculatedOrderContainer(message *string, ignoreDiscountsFlag *bool) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		message,
		ignoreDiscountsFlag,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Generate a [[SoftLayer_Container_Product_Order_Receipt]] object with all the order information.
func (r *Billing_Order) GetReceipt() (resp datatypes.Container_Product_Order_Receipt, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The Referral Partner who referred this order. (Only necessary for new customer orders)
func (r *Billing_Order) GetReferralPartner() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve This flag indicates an order is an upgrade.
func (r *Billing_Order) GetUpgradeRequestFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_User_Customer object tied to an order.
func (r *Billing_Order) GetUserRecord() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// When an order has been modified, it will contain a status indicating so. This method checks that status and also verifies that the active user's account is the same as the account on the order.
func (r *Billing_Order) IsPendingEditApproval() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Billing_Order_Cart struct {
	Session *Session
	Options
}

func (r *Session) GetBillingOrderCartService() Billing_Order_Cart {
	return Billing_Order_Cart{Session: r}
}

// This method is used to transfer an anonymous quote to the active user and associated account. An anonymous quote is one that was created by a user without being authenticated. If a quote was created anonymously and then the customer attempts to access that anonymous quote via the API (which requires authentication), the customer will be unable to retrieve the quote due to the security restrictions in place. By providing the ability for a customer to claim a quote, s/he will be able to pull the anonymous quote onto his/her account and successfully view the quote.
//
// To claim a quote, both the quote id and the quote key (the 32-character random string) must be provided.
func (r *Billing_Order_Cart) Claim(quoteKey *string, quoteId *int) (resp datatypes.Billing_Order_Quote, err error) {
	params := []interface{}{
		quoteKey,
		quoteId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// When creating a new cart, the order data is sent through SoftLayer_Product_Order::verifyOrder to make sure that the cart contains valid data. If an issue is found with the order, an exception will be thrown and you will receive the same response as if SoftLayer_Product_Order::verifyOrder were called directly. Once the order verification is complete, the cart will be created.
//
// The response is the new cart id.
func (r *Billing_Order_Cart) CreateCart(orderData *datatypes.Container_Product_Order) (resp int, err error) {
	params := []interface{}{
		orderData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// If a cart is no longer needed, it can be deleted using this service. Once a cart has been deleted, it cannot be retrieved again.
func (r *Billing_Order_Cart) DeleteCart() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Account master users and sub-users in the SoftLayer customer portal can delete the quote of an order.
func (r *Billing_Order_Cart) DeleteQuote() (resp datatypes.Billing_Order_Quote, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A quote's corresponding account.
func (r *Billing_Order_Cart) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a valid cart record of a SoftLayer order.
func (r *Billing_Order_Cart) GetCartByCartKey(cartKey *string) (resp datatypes.Billing_Order_Cart, err error) {
	params := []interface{}{
		cartKey,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Billing_Order_Cart) GetObject() (resp datatypes.Billing_Order_Cart, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve This order contains the records for which products were selected for this quote.
func (r *Billing_Order_Cart) GetOrder() (resp datatypes.Billing_Order, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve These are all the orders that were created from this quote.
func (r *Billing_Order_Cart) GetOrdersFromQuote() (resp []datatypes.Billing_Order, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a PDF copy of the cart.
func (r *Billing_Order_Cart) GetPdf() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a valid quote record of a SoftLayer order. Quote whose promotion code expired or one of the items is no longer available will not be retrieved.
func (r *Billing_Order_Cart) GetQuoteByQuoteKey(quoteKey *string) (resp datatypes.Billing_Order_Quote, err error) {
	params := []interface{}{
		quoteKey,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// This method allows the customer to retrieve a saved cart and put it in a format that's suitable to be sent to SoftLayer_Billing_Order_Cart::createCart to create a new cart or to SoftLayer_Billing_Order_Cart::updateCart to update an existing cart.
func (r *Billing_Order_Cart) GetRecalculatedOrderContainer(orderData *datatypes.Container_Product_Order, orderBeingPlacedFlag *bool) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		orderData,
		orderBeingPlacedFlag,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Use this method for placing server orders and additional services orders. The same applies for this as with verifyOrder. Send in the SoftLayer_Container_Product_Order_Hardware_Server for server orders. In addition to verifying the order, placeOrder() also makes an initial authorization on the SoftLayer_Account tied to this order, if a credit card is on file. If the account tied to this order is a paypal customer, an URL will also be returned to the customer. After placing the order, you must go to this URL to finish the authorization process. This tells paypal that you indeed want to place the order. After going to this URL, it will direct you back to a SoftLayer webpage that tells us you have finished the process. After this, it will go to sales for final approval.
func (r *Billing_Order_Cart) PlaceOrder(orderData *datatypes.Container_Product_Order) (resp datatypes.Container_Product_Order_Receipt, err error) {
	params := []interface{}{
		orderData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Use this method for placing server quotes and additional services quotes. The same applies for this as with verifyOrder. Send in the SoftLayer_Container_Product_Order_Hardware_Server for server quotes. In addition to verifying the quote, placeQuote() also makes an initial authorization on the SoftLayer_Account tied to this order, if a credit card is on file. If the account tied to this order is a paypal customer, an URL will also be returned to the customer. After placing the order, you must go to this URL to finish the authorization process. This tells paypal that you indeed want to place the order. After going to this URL, it will direct you back to a SoftLayer webpage that tells us you have finished the process.
func (r *Billing_Order_Cart) PlaceQuote(orderData *datatypes.Container_Product_Order) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		orderData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Account master users and sub-users in the SoftLayer customer portal can save the quote of an order to avoid its deletion after 5 days or its expiration after 2 days.
func (r *Billing_Order_Cart) SaveQuote() (resp datatypes.Billing_Order_Quote, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Like SoftLayer_Billing_Order_Cart::createCart, the order data will be sent through SoftLayer_Product_Order::verifyOrder to make sure that the updated cart information is valid. Once it has been verified, the new order data will be saved.
//
// This will return the cart id.
func (r *Billing_Order_Cart) UpdateCart(orderData *datatypes.Container_Product_Order) (resp int, err error) {
	params := []interface{}{
		orderData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Use this method for placing server orders and additional services orders. The same applies for this as with verifyOrder. Send in the SoftLayer_Container_Product_Order_Hardware_Server for server orders. In addition to verifying the order, placeOrder() also makes an initial authorization on the SoftLayer_Account tied to this order, if a credit card is on file. If the account tied to this order is a paypal customer, an URL will also be returned to the customer. After placing the order, you must go to this URL to finish the authorization process. This tells paypal that you indeed want to place the order. After going to this URL, it will direct you back to a SoftLayer webpage that tells us you have finished the process. After this, it will go to sales for final approval.
func (r *Billing_Order_Cart) VerifyOrder(orderData *datatypes.Container_Product_Order) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		orderData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Every individual item that a SoftLayer customer is billed for is recorded in the SoftLayer_Billing_Item data type. Billing items range from server chassis to hard drives to control panels, bandwidth quota upgrades and port upgrade charges. Softlayer [[SoftLayer_Billing_Invoice|invoices]] are generated from the cost of a customer's billing items. Billing items are copied from the product catalog as they're ordered by customers to create a reference between an account and the billable items they own.
//
// Billing items exist in a tree relationship. Items are associated with each other by parent/child relationships. Component items such as CPU's, RAM, and software each have a parent billing item for the server chassis they're associated with. Billing Items with a null parent item do not have an associated parent item.
type Billing_Order_Item struct {
	Session *Session
	Options
}

func (r *Session) GetBillingOrderItemService() Billing_Order_Item {
	return Billing_Order_Item{Session: r}
}

// Retrieve The SoftLayer_Billing_Item tied to the order item.
func (r *Billing_Order_Item) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The other items included with an ordered item.
func (r *Billing_Order_Item) GetBundledItems() (resp []datatypes.Billing_Order_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The item category tied to an order item.
func (r *Billing_Order_Item) GetCategory() (resp datatypes.Product_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The child order items for an order item. All server order items should have children. These children are considered a part of the server.
func (r *Billing_Order_Item) GetChildren() (resp []datatypes.Billing_Order_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A hardware's universally unique identifier.
func (r *Billing_Order_Item) GetGlobalIdentifier() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The component type tied to an order item. All hardware-specific items should have a generic hardware component.
func (r *Billing_Order_Item) GetHardwareGenericComponent() (resp datatypes.Hardware_Component_Model_Generic, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Product_Item tied to an order item. The item is the actual definition of the product being sold.
func (r *Billing_Order_Item) GetItem() (resp datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve This is an item's category answers.
func (r *Billing_Order_Item) GetItemCategoryAnswers() (resp []datatypes.Billing_Order_Item_Category_Answer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Product_Item_Price tied to an order item. The item price object describes the cost of an item.
func (r *Billing_Order_Item) GetItemPrice() (resp datatypes.Product_Item_Price, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The location of an ordered item. This is usually the same as the server it is being ordered with. Otherwise it describes the location of the additional service being ordered.
func (r *Billing_Order_Item) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Billing_Order_Item) GetNextOrderChildren() (resp []datatypes.Billing_Order_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Billing_Item object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Billing_Item service. You can only retrieve billing items tied to the account that your portal user is assigned to. Billing items are an account's items of billable items. There are "parent" billing items and "child" billing items. The server billing item is generally referred to as a parent billing item. The items tied to a server, such as ram, harddrives, and operating systems are considered "child" billing items.
func (r *Billing_Order_Item) GetObject() (resp datatypes.Billing_Order_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve This is only populated when an upgrade order is placed. The old billing item represents what the billing was before the upgrade happened.
func (r *Billing_Order_Item) GetOldBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The order to which this item belongs. The order contains all the information related to the items included in an order
func (r *Billing_Order_Item) GetOrder() (resp datatypes.Billing_Order, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Billing_Order_Item) GetOrderApprovalDate() (resp datatypes.Time, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The SoftLayer_Product_Package an order item is a part of.
func (r *Billing_Order_Item) GetPackage() (resp datatypes.Product_Package, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The parent order item ID for an item. Items that are associated with a server will have a parent. The parent will be the server item itself.
func (r *Billing_Order_Item) GetParent() (resp datatypes.Billing_Order_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A count of power supplies contained within this SoftLayer_Billing_Order
func (r *Billing_Order_Item) GetRedundantPowerSupplyCount() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve For ordered items that are software items, a full description of that software can be found with this property.
func (r *Billing_Order_Item) GetSoftwareDescription() (resp datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The drive storage groups that are attached to this billing order item.
func (r *Billing_Order_Item) GetStorageGroups() (resp []datatypes.Configuration_Storage_Group_Order, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The recurring fee of an ordered item. This amount represents the fees that will be charged on a recurring (usually monthly) basis.
func (r *Billing_Order_Item) GetTotalRecurringAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The next SoftLayer_Product_Item in the upgrade path for this order item.
func (r *Billing_Order_Item) GetUpgradeItem() (resp datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// The SoftLayer_Billing_Oder_Quote data type contains general information relating to an individual order applied to a SoftLayer customer account or to a new customer. Personal information in this type such as names, addresses, and phone numbers are taken from the account's contact information at the time the quote is generated for existing SoftLayer customer.
type Billing_Order_Quote struct {
	Session *Session
	Options
}

func (r *Session) GetBillingOrderQuoteService() Billing_Order_Quote {
	return Billing_Order_Quote{Session: r}
}

// This method is used to transfer an anonymous quote to the active user and associated account. An anonymous quote is one that was created by a user without being authenticated. If a quote was created anonymously and then the customer attempts to access that anonymous quote via the API (which requires authentication), the customer will be unable to retrieve the quote due to the security restrictions in place. By providing the ability for a customer to claim a quote, s/he will be able to pull the anonymous quote onto his/her account and successfully view the quote.
//
// To claim a quote, both the quote id and the quote key (the 32-character random string) must be provided.
func (r *Billing_Order_Quote) Claim(quoteKey *string, quoteId *int) (resp datatypes.Billing_Order_Quote, err error) {
	params := []interface{}{
		quoteKey,
		quoteId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Account master users and sub-users in the SoftLayer customer portal can delete the quote of an order.
func (r *Billing_Order_Quote) DeleteQuote() (resp datatypes.Billing_Order_Quote, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A quote's corresponding account.
func (r *Billing_Order_Quote) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Billing_Order_Quote object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Billing_Order_Quote service. You can only retrieve quotes that are assigned to your portal user's account.
func (r *Billing_Order_Quote) GetObject() (resp datatypes.Billing_Order_Quote, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve This order contains the records for which products were selected for this quote.
func (r *Billing_Order_Quote) GetOrder() (resp datatypes.Billing_Order, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve These are all the orders that were created from this quote.
func (r *Billing_Order_Quote) GetOrdersFromQuote() (resp []datatypes.Billing_Order, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a PDF record of a SoftLayer quoted order. SoftLayer keeps PDF records of all quoted orders for customer retrieval from the portal and API. You must have a PDF reader installed in order to view these quoted order files.
func (r *Billing_Order_Quote) GetPdf() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a valid quote record of a SoftLayer order. Quote whose promotion code expired or one of the items is no longer available will not be retrieved.
func (r *Billing_Order_Quote) GetQuoteByQuoteKey(quoteKey *string) (resp datatypes.Billing_Order_Quote, err error) {
	params := []interface{}{
		quoteKey,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Generate an [[SoftLayer_Container_Product_Order|order container]] from the previously-created quote. This will take into account promotions, reseller status, estimated taxes and all other standard order verification processes.
func (r *Billing_Order_Quote) GetRecalculatedOrderContainer(orderData *datatypes.Container_Product_Order, orderBeingPlacedFlag *bool) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		orderData,
		orderBeingPlacedFlag,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Use this method for placing server orders and additional services orders. The same applies for this as with verifyOrder. Send in the SoftLayer_Container_Product_Order_Hardware_Server for server orders. In addition to verifying the order, placeOrder() also makes an initial authorization on the SoftLayer_Account tied to this order, if a credit card is on file. If the account tied to this order is a paypal customer, an URL will also be returned to the customer. After placing the order, you must go to this URL to finish the authorization process. This tells paypal that you indeed want to place the order. After going to this URL, it will direct you back to a SoftLayer webpage that tells us you have finished the process. After this, it will go to sales for final approval.
func (r *Billing_Order_Quote) PlaceOrder(orderData *datatypes.Container_Product_Order) (resp datatypes.Container_Product_Order_Receipt, err error) {
	params := []interface{}{
		orderData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Use this method for placing server quotes and additional services quotes. The same applies for this as with verifyOrder. Send in the SoftLayer_Container_Product_Order_Hardware_Server for server quotes. In addition to verifying the quote, placeQuote() also makes an initial authorization on the SoftLayer_Account tied to this order, if a credit card is on file. If the account tied to this order is a paypal customer, an URL will also be returned to the customer. After placing the order, you must go to this URL to finish the authorization process. This tells paypal that you indeed want to place the order. After going to this URL, it will direct you back to a SoftLayer webpage that tells us you have finished the process.
func (r *Billing_Order_Quote) PlaceQuote(orderData *datatypes.Container_Product_Order) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		orderData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Account master users and sub-users in the SoftLayer customer portal can save the quote of an order to avoid its deletion after 5 days or its expiration after 2 days.
func (r *Billing_Order_Quote) SaveQuote() (resp datatypes.Billing_Order_Quote, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Use this method for placing server orders and additional services orders. The same applies for this as with verifyOrder. Send in the SoftLayer_Container_Product_Order_Hardware_Server for server orders. In addition to verifying the order, placeOrder() also makes an initial authorization on the SoftLayer_Account tied to this order, if a credit card is on file. If the account tied to this order is a paypal customer, an URL will also be returned to the customer. After placing the order, you must go to this URL to finish the authorization process. This tells paypal that you indeed want to place the order. After going to this URL, it will direct you back to a SoftLayer webpage that tells us you have finished the process. After this, it will go to sales for final approval.
func (r *Billing_Order_Quote) VerifyOrder(orderData *datatypes.Container_Product_Order) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		orderData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
