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

type Billing_Currency struct {
	Session *Session
	Options
}

func (r *Session) GetBillingCurrencyService() Billing_Currency {
	return Billing_Currency{Session: r}
}

func (r *Billing_Currency) GetAllObjects() (resp []datatypes.Billing_Currency, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Currency) GetObject() (resp datatypes.Billing_Currency, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Currency) GetPrice(price *float64, formatOptions *datatypes.Container_Billing_Currency_Format) (resp string, err error) {
	params := []interface{}{
		price,
		formatOptions,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Billing_Currency_ExchangeRate struct {
	Session *Session
	Options
}

func (r *Session) GetBillingCurrencyExchangeRateService() Billing_Currency_ExchangeRate {
	return Billing_Currency_ExchangeRate{Session: r}
}

func (r *Billing_Currency_ExchangeRate) GetAllCurrencyExchangeRates(stringDate *string) (resp []datatypes.Billing_Currency_ExchangeRate, err error) {
	params := []interface{}{
		stringDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Currency_ExchangeRate) GetCurrencies() (resp []datatypes.Billing_Currency, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Currency_ExchangeRate) GetExchangeRate(to *string, from *string, effectiveDate *datatypes.Time) (resp datatypes.Billing_Currency_ExchangeRate, err error) {
	params := []interface{}{
		to,
		from,
		effectiveDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Currency_ExchangeRate) GetObject() (resp datatypes.Billing_Currency_ExchangeRate, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Currency_ExchangeRate) GetPrice(price *float64, formatOptions *datatypes.Container_Billing_Currency_Format) (resp string, err error) {
	params := []interface{}{
		price,
		formatOptions,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

func (r *Billing_Currency_ExchangeRate) GetFundingCurrency() (resp datatypes.Billing_Currency, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Currency_ExchangeRate) GetLocalCurrency() (resp datatypes.Billing_Currency, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Info struct {
	Session *Session
	Options
}

func (r *Session) GetBillingInfoService() Billing_Info {
	return Billing_Info{Session: r}
}

func (r *Billing_Info) GetObject() (resp datatypes.Billing_Info, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Billing_Info) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Info) GetAchInformation() (resp []datatypes.Billing_Info_Ach, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Info) GetCurrency() (resp datatypes.Billing_Currency, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Info) GetCurrentBillingCycle() (resp datatypes.Billing_Info_Cycle, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Info) GetLastBillDate() (resp datatypes.Time, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Info) GetNextBillDate() (resp datatypes.Time, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Info_Ach struct {
	Session *Session
	Options
}

func (r *Session) GetBillingInfoAchService() Billing_Info_Ach {
	return Billing_Info_Ach{Session: r}
}

func (r *Billing_Info_Ach) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Info_Cycle struct {
	Session *Session
	Options
}

func (r *Session) GetBillingInfoCycleService() Billing_Info_Cycle {
	return Billing_Info_Cycle{Session: r}
}

func (r *Billing_Info_Cycle) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Invoice struct {
	Session *Session
	Options
}

func (r *Session) GetBillingInvoiceService() Billing_Invoice {
	return Billing_Invoice{Session: r}
}

func (r *Billing_Invoice) EmailInvoices(options *datatypes.Container_Billing_Invoice_Email) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		options,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetExcel() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetObject() (resp datatypes.Billing_Invoice, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetPdf() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetPdfDetailed() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetPdfDetailedFilename() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetPdfFileSize() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetPdfFilename() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetPreliminaryExcel() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetPreliminaryPdf() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetPreliminaryPdfDetailed() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetXlsFilename() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetZeroFeeItemCounts() (resp []datatypes.Container_Product_Item_Category_ZeroFee_Count, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Billing_Invoice) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetBrandAtInvoiceCreation() (resp datatypes.Brand, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetDetailedPdfGeneratedFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetInvoiceTopLevelItems() (resp []datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetInvoiceTotalAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetInvoiceTotalOneTimeAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetInvoiceTotalOneTimeTaxAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetInvoiceTotalPreTaxAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetInvoiceTotalRecurringAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetInvoiceTotalRecurringTaxAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetItems() (resp []datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetPayment() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetPayments() (resp []datatypes.Billing_Invoice_Receivable_Payment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetSellerRegistration() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetTaxInfo() (resp datatypes.Billing_Invoice_Tax_Info, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetTaxInfoHistory() (resp []datatypes.Billing_Invoice_Tax_Info, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetTaxMessage() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice) GetTaxType() (resp datatypes.Billing_Invoice_Tax_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Invoice_Item struct {
	Session *Session
	Options
}

func (r *Session) GetBillingInvoiceItemService() Billing_Invoice_Item {
	return Billing_Invoice_Item{Session: r}
}

func (r *Billing_Invoice_Item) GetObject() (resp datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Billing_Invoice_Item) GetAssociatedChildren() (resp []datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Item) GetAssociatedInvoiceItem() (resp datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Item) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Item) GetCategory() (resp datatypes.Product_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Item) GetChildren() (resp []datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Item) GetFilteredAssociatedChildren() (resp []datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Item) GetInvoice() (resp datatypes.Billing_Invoice, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Item) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Item) GetNonZeroAssociatedChildren() (resp []datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Item) GetParent() (resp datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Item) GetProduct() (resp datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Item) GetTotalOneTimeAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Item) GetTotalOneTimeTaxAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Item) GetTotalRecurringAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Item) GetTotalRecurringTaxAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Invoice_Item_Hardware struct {
	Session *Session
	Options
}

func (r *Session) GetBillingInvoiceItemHardwareService() Billing_Invoice_Item_Hardware {
	return Billing_Invoice_Item_Hardware{Session: r}
}

func (r *Billing_Invoice_Item_Hardware) GetResource() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Invoice_Item_Tax_Info struct {
	Session *Session
	Options
}

func (r *Session) GetBillingInvoiceItemTaxInfoService() Billing_Invoice_Item_Tax_Info {
	return Billing_Invoice_Item_Tax_Info{Session: r}
}

func (r *Billing_Invoice_Item_Tax_Info) GetInvoiceItem() (resp datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Item_Tax_Info) GetInvoiceTaxInfo() (resp datatypes.Billing_Invoice_Tax_Info, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Item_Tax_Info) GetToCurrency() (resp datatypes.Billing_Currency, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Invoice_Next struct {
	Session *Session
	Options
}

func (r *Session) GetBillingInvoiceNextService() Billing_Invoice_Next {
	return Billing_Invoice_Next{Session: r}
}

func (r *Billing_Invoice_Next) GetExcel(documentCreateDate *datatypes.Time) (resp []byte, err error) {
	params := []interface{}{
		documentCreateDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Next) GetPdf(documentCreateDate *datatypes.Time) (resp []byte, err error) {
	params := []interface{}{
		documentCreateDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Next) GetPdfDetailed(documentCreateDate *datatypes.Time) (resp []byte, err error) {
	params := []interface{}{
		documentCreateDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Billing_Invoice_Receivable_Payment struct {
	Session *Session
	Options
}

func (r *Session) GetBillingInvoiceReceivablePaymentService() Billing_Invoice_Receivable_Payment {
	return Billing_Invoice_Receivable_Payment{Session: r}
}

func (r *Billing_Invoice_Receivable_Payment) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Receivable_Payment) GetCreditCardLastFourDigits() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Receivable_Payment) GetCreditCardRequestId() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Receivable_Payment) GetCreditCardTransaction() (resp datatypes.Billing_Payment_Card_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Receivable_Payment) GetExchangeRate() (resp datatypes.Billing_Currency_ExchangeRate, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Receivable_Payment) GetInvoice() (resp datatypes.Billing_Invoice, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Receivable_Payment) GetPaypalTransaction() (resp datatypes.Billing_Payment_PayPal_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Invoice_Tax_Info struct {
	Session *Session
	Options
}

func (r *Session) GetBillingInvoiceTaxInfoService() Billing_Invoice_Tax_Info {
	return Billing_Invoice_Tax_Info{Session: r}
}

func (r *Billing_Invoice_Tax_Info) GetCurrency() (resp datatypes.Billing_Currency, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Tax_Info) GetFunctionalCurrency() (resp datatypes.Billing_Currency, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Tax_Info) GetInvoice() (resp datatypes.Billing_Invoice, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Tax_Info) GetItemWithCurrencyInfo() (resp datatypes.Billing_Invoice_Item_Tax_Info, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Tax_Info) GetItems() (resp []datatypes.Billing_Invoice_Item_Tax_Info, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Tax_Info) GetTotalTaxAmountToCurrency() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Invoice_Tax_Status struct {
	Session *Session
	Options
}

func (r *Session) GetBillingInvoiceTaxStatusService() Billing_Invoice_Tax_Status {
	return Billing_Invoice_Tax_Status{Session: r}
}

func (r *Billing_Invoice_Tax_Status) GetAllObjects() (resp []datatypes.Billing_Invoice_Tax_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Tax_Status) GetObject() (resp datatypes.Billing_Invoice_Tax_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Invoice_Tax_Type struct {
	Session *Session
	Options
}

func (r *Session) GetBillingInvoiceTaxTypeService() Billing_Invoice_Tax_Type {
	return Billing_Invoice_Tax_Type{Session: r}
}

func (r *Billing_Invoice_Tax_Type) GetAllObjects() (resp []datatypes.Billing_Invoice_Tax_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Invoice_Tax_Type) GetObject() (resp datatypes.Billing_Invoice_Tax_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemService() Billing_Item {
	return Billing_Item{Session: r}
}

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
func (r *Billing_Item) CancelService() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) CancelServiceOnAnniversaryDate() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetObject() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetServiceBillingItemsByCategory(categoryCode *string, includeZeroRecurringFee *bool) (resp []datatypes.Billing_Item, err error) {
	params := []interface{}{
		categoryCode,
		includeZeroRecurringFee,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) RemoveAssociationId() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) SetAssociationId(associatedId *int) (resp bool, err error) {
	params := []interface{}{
		associatedId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) VoidCancelService() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Billing_Item) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetActiveAgreement() (resp datatypes.Account_Agreement, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetActiveAgreementFlag() (resp datatypes.Account_Agreement, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetActiveAssociatedChildren() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetActiveAssociatedGuestDiskBillingItems() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetActiveBundledItems() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetActiveCancellationItem() (resp datatypes.Billing_Item_Cancellation_Request_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetActiveChildren() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetActiveFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetActiveSparePoolAssociatedGuestDiskBillingItems() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetActiveSparePoolBundledItems() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetAssociatedBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetAssociatedBillingItemHistory() (resp []datatypes.Billing_Item_Association_History, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetAssociatedChildren() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetAssociatedParent() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetAvailableMatchingVlans() (resp []datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetBandwidthAllocation() (resp datatypes.Network_Bandwidth_Version1_Allocation, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetBillableChildren() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetBundleItems() (resp []datatypes.Product_Item_Bundles, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetBundledItems() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetCanceledChildren() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetCancellationReason() (resp datatypes.Billing_Item_Cancellation_Reason, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetCancellationRequests() (resp []datatypes.Billing_Item_Cancellation_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetCategory() (resp datatypes.Product_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetChildren() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetChildrenWithActiveAgreement() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetDowngradeItems() (resp []datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetFilteredNextInvoiceChildren() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetHourlyFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetInvoiceItem() (resp datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetInvoiceItems() (resp []datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetItem() (resp datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetNextInvoiceChildren() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetNextInvoiceTotalOneTimeAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetNextInvoiceTotalOneTimeTaxAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetNextInvoiceTotalRecurringAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetNextInvoiceTotalRecurringTaxAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetNonZeroNextInvoiceChildren() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetOrderItem() (resp datatypes.Billing_Order_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetOriginalLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetPackage() (resp datatypes.Product_Package, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetParent() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetParentVirtualGuestBillingItem() (resp datatypes.Billing_Item_Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetPendingCancellationFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetPendingOrderItem() (resp datatypes.Billing_Order_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetProvisionTransaction() (resp datatypes.Provisioning_Version1_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetSoftwareDescription() (resp datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetUpgradeItem() (resp datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item) GetUpgradeItems() (resp []datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Account_Media_Data_Transfer_Request struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemAccountMediaDataTransferRequestService() Billing_Item_Account_Media_Data_Transfer_Request {
	return Billing_Item_Account_Media_Data_Transfer_Request{Session: r}
}

func (r *Billing_Item_Account_Media_Data_Transfer_Request) GetResource() (resp datatypes.Account_Media_Data_Transfer_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Association_History struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemAssociationHistoryService() Billing_Item_Association_History {
	return Billing_Item_Association_History{Session: r}
}

func (r *Billing_Item_Association_History) GetAssociatedBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Association_History) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Cancellation_Reason struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemCancellationReasonService() Billing_Item_Cancellation_Reason {
	return Billing_Item_Cancellation_Reason{Session: r}
}

func (r *Billing_Item_Cancellation_Reason) GetAllCancellationReasons() (resp []datatypes.Billing_Item_Cancellation_Reason, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Cancellation_Reason) GetObject() (resp datatypes.Billing_Item_Cancellation_Reason, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Billing_Item_Cancellation_Reason) GetBillingCancellationReasonCategory() (resp datatypes.Billing_Item_Cancellation_Reason_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Cancellation_Reason) GetBillingItems() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Cancellation_Reason) GetTranslatedReason() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Cancellation_Reason_Category struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemCancellationReasonCategoryService() Billing_Item_Cancellation_Reason_Category {
	return Billing_Item_Cancellation_Reason_Category{Session: r}
}

func (r *Billing_Item_Cancellation_Reason_Category) GetAllCancellationReasonCategories() (resp []datatypes.Billing_Item_Cancellation_Reason_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Cancellation_Reason_Category) GetObject() (resp datatypes.Billing_Item_Cancellation_Reason_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Billing_Item_Cancellation_Reason_Category) GetBillingCancellationReasons() (resp []datatypes.Billing_Item_Cancellation_Reason, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Cancellation_Request struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemCancellationRequestService() Billing_Item_Cancellation_Request {
	return Billing_Item_Cancellation_Request{Session: r}
}

func (r *Billing_Item_Cancellation_Request) CreateObject(templateObject *datatypes.Billing_Item_Cancellation_Request) (resp datatypes.Billing_Item_Cancellation_Request, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Cancellation_Request) GetAllCancellationRequests() (resp []datatypes.Billing_Item_Cancellation_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Cancellation_Request) GetCancellationCutoffDate(accountId *int, categoryCode *string) (resp datatypes.Time, err error) {
	params := []interface{}{
		accountId,
		categoryCode,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Cancellation_Request) GetObject() (resp datatypes.Billing_Item_Cancellation_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Cancellation_Request) RemoveCancellationItem(itemId *int) (resp bool, err error) {
	params := []interface{}{
		itemId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Cancellation_Request) ValidateBillingItemForCancellation(billingItemId *int) (resp bool, err error) {
	params := []interface{}{
		billingItemId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Cancellation_Request) Void(closeRelatedTicketFlag *bool) (resp bool, err error) {
	params := []interface{}{
		closeRelatedTicketFlag,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

func (r *Billing_Item_Cancellation_Request) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Cancellation_Request) GetItems() (resp []datatypes.Billing_Item_Cancellation_Request_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Cancellation_Request) GetStatus() (resp datatypes.Billing_Item_Cancellation_Request_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Cancellation_Request) GetTicket() (resp datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Cancellation_Request) GetUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Cancellation_Request_Item struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemCancellationRequestItemService() Billing_Item_Cancellation_Request_Item {
	return Billing_Item_Cancellation_Request_Item{Session: r}
}

func (r *Billing_Item_Cancellation_Request_Item) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Cancellation_Request_Item) GetCancellationRequest() (resp datatypes.Billing_Item_Cancellation_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Cancellation_Request_Status struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemCancellationRequestStatusService() Billing_Item_Cancellation_Request_Status {
	return Billing_Item_Cancellation_Request_Status{Session: r}
}

type Billing_Item_Ctc_Account struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemCtcAccountService() Billing_Item_Ctc_Account {
	return Billing_Item_Ctc_Account{Session: r}
}

type Billing_Item_Gateway_Appliance_Cluster struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemGatewayApplianceClusterService() Billing_Item_Gateway_Appliance_Cluster {
	return Billing_Item_Gateway_Appliance_Cluster{Session: r}
}

func (r *Billing_Item_Gateway_Appliance_Cluster) GetResource() (resp datatypes.Resource_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Hardware struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemHardwareService() Billing_Item_Hardware {
	return Billing_Item_Hardware{Session: r}
}

func (r *Billing_Item_Hardware) GetBillingCycleBandwidthUsage() (resp []datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Hardware) GetBillingCyclePrivateBandwidthUsage() (resp []datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Hardware) GetBillingCyclePrivateUsageIn() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Hardware) GetBillingCyclePrivateUsageOut() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Hardware) GetBillingCyclePrivateUsageTotal() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Hardware) GetBillingCyclePublicBandwidthUsage() (resp []datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Hardware) GetBillingCyclePublicUsageIn() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Hardware) GetBillingCyclePublicUsageOut() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Hardware) GetBillingCyclePublicUsageTotal() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Hardware) GetLockboxNetworkStorage() (resp datatypes.Billing_Item_Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Hardware) GetMonitoringBillingItems() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Hardware) GetResource() (resp datatypes.Hardware_Server, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Hardware_Colocation struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemHardwareColocationService() Billing_Item_Hardware_Colocation {
	return Billing_Item_Hardware_Colocation{Session: r}
}

type Billing_Item_Hardware_Component struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemHardwareComponentService() Billing_Item_Hardware_Component {
	return Billing_Item_Hardware_Component{Session: r}
}

func (r *Billing_Item_Hardware_Component) GetResource() (resp []datatypes.Hardware_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Hardware_Security_Module struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemHardwareSecurityModuleService() Billing_Item_Hardware_Security_Module {
	return Billing_Item_Hardware_Security_Module{Session: r}
}

type Billing_Item_Hardware_Server struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemHardwareServerService() Billing_Item_Hardware_Server {
	return Billing_Item_Hardware_Server{Session: r}
}

type Billing_Item_Link_ThePlanet struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemLinkThePlanetService() Billing_Item_Link_ThePlanet {
	return Billing_Item_Link_ThePlanet{Session: r}
}

func (r *Billing_Item_Link_ThePlanet) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Link_ThePlanet) GetServiceProvider() (resp datatypes.Service_Provider, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Network_Application_Delivery_Controller struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemNetworkApplicationDeliveryControllerService() Billing_Item_Network_Application_Delivery_Controller {
	return Billing_Item_Network_Application_Delivery_Controller{Session: r}
}

func (r *Billing_Item_Network_Application_Delivery_Controller) GetBandwidthAllotmentDetail() (resp datatypes.Network_Bandwidth_Version1_Allotment_Detail, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Network_Application_Delivery_Controller) GetResource() (resp datatypes.Network_Application_Delivery_Controller, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemNetworkApplicationDeliveryControllerLoadBalancerVirtualIpAddressService() Billing_Item_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress {
	return Billing_Item_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress{Session: r}
}

func (r *Billing_Item_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) GetResource() (resp datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Network_Bandwidth struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemNetworkBandwidthService() Billing_Item_Network_Bandwidth {
	return Billing_Item_Network_Bandwidth{Session: r}
}

type Billing_Item_Network_Firewall struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemNetworkFirewallService() Billing_Item_Network_Firewall {
	return Billing_Item_Network_Firewall{Session: r}
}

func (r *Billing_Item_Network_Firewall) GetResource() (resp datatypes.Network_Component_Firewall, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Network_Firewall_Module_Context struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemNetworkFirewallModuleContextService() Billing_Item_Network_Firewall_Module_Context {
	return Billing_Item_Network_Firewall_Module_Context{Session: r}
}

type Billing_Item_Network_Interconnect struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemNetworkInterconnectService() Billing_Item_Network_Interconnect {
	return Billing_Item_Network_Interconnect{Session: r}
}

type Billing_Item_Network_LoadBalancer struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemNetworkLoadBalancerService() Billing_Item_Network_LoadBalancer {
	return Billing_Item_Network_LoadBalancer{Session: r}
}

type Billing_Item_Network_LoadBalancer_Global struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemNetworkLoadBalancerGlobalService() Billing_Item_Network_LoadBalancer_Global {
	return Billing_Item_Network_LoadBalancer_Global{Session: r}
}

func (r *Billing_Item_Network_LoadBalancer_Global) GetResource() (resp datatypes.Network_LoadBalancer_Global_Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Network_LoadBalancer_VirtualIpAddress struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemNetworkLoadBalancerVirtualIpAddressService() Billing_Item_Network_LoadBalancer_VirtualIpAddress {
	return Billing_Item_Network_LoadBalancer_VirtualIpAddress{Session: r}
}

func (r *Billing_Item_Network_LoadBalancer_VirtualIpAddress) GetResource() (resp datatypes.Network_LoadBalancer_VirtualIpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Network_Message_Delivery struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemNetworkMessageDeliveryService() Billing_Item_Network_Message_Delivery {
	return Billing_Item_Network_Message_Delivery{Session: r}
}

func (r *Billing_Item_Network_Message_Delivery) GetResource() (resp datatypes.Network_Message_Delivery, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Network_Message_Queue struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemNetworkMessageQueueService() Billing_Item_Network_Message_Queue {
	return Billing_Item_Network_Message_Queue{Session: r}
}

func (r *Billing_Item_Network_Message_Queue) GetResource() (resp datatypes.Network_Message_Queue, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Network_Message_Queue_Delivery struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemNetworkMessageQueueDeliveryService() Billing_Item_Network_Message_Queue_Delivery {
	return Billing_Item_Network_Message_Queue_Delivery{Session: r}
}

type Billing_Item_Network_PerformanceStorage_Iscsi struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemNetworkPerformanceStorageIscsiService() Billing_Item_Network_PerformanceStorage_Iscsi {
	return Billing_Item_Network_PerformanceStorage_Iscsi{Session: r}
}

type Billing_Item_Network_PerformanceStorage_Nfs struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemNetworkPerformanceStorageNfsService() Billing_Item_Network_PerformanceStorage_Nfs {
	return Billing_Item_Network_PerformanceStorage_Nfs{Session: r}
}

type Billing_Item_Network_Storage struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemNetworkStorageService() Billing_Item_Network_Storage {
	return Billing_Item_Network_Storage{Session: r}
}

func (r *Billing_Item_Network_Storage) GetResource() (resp datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Network_Storage_Hub struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemNetworkStorageHubService() Billing_Item_Network_Storage_Hub {
	return Billing_Item_Network_Storage_Hub{Session: r}
}

type Billing_Item_Network_Storage_Hub_Bandwidth struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemNetworkStorageHubBandwidthService() Billing_Item_Network_Storage_Hub_Bandwidth {
	return Billing_Item_Network_Storage_Hub_Bandwidth{Session: r}
}

type Billing_Item_Network_Subnet struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemNetworkSubnetService() Billing_Item_Network_Subnet {
	return Billing_Item_Network_Subnet{Session: r}
}

func (r *Billing_Item_Network_Subnet) GetResource() (resp datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Network_Subnet_IpAddress_Global struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemNetworkSubnetIpAddressGlobalService() Billing_Item_Network_Subnet_IpAddress_Global {
	return Billing_Item_Network_Subnet_IpAddress_Global{Session: r}
}

type Billing_Item_Network_Tunnel struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemNetworkTunnelService() Billing_Item_Network_Tunnel {
	return Billing_Item_Network_Tunnel{Session: r}
}

func (r *Billing_Item_Network_Tunnel) GetResource() (resp datatypes.Network_Tunnel_Module_Context, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Network_Vlan struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemNetworkVlanService() Billing_Item_Network_Vlan {
	return Billing_Item_Network_Vlan{Session: r}
}

func (r *Billing_Item_Network_Vlan) GetResource() (resp datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Software_Component struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemSoftwareComponentService() Billing_Item_Software_Component {
	return Billing_Item_Software_Component{Session: r}
}

func (r *Billing_Item_Software_Component) GetResource() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Software_Component_Analytics_Urchin struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemSoftwareComponentAnalyticsUrchinService() Billing_Item_Software_Component_Analytics_Urchin {
	return Billing_Item_Software_Component_Analytics_Urchin{Session: r}
}

type Billing_Item_Software_Component_ControlPanel struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemSoftwareComponentControlPanelService() Billing_Item_Software_Component_ControlPanel {
	return Billing_Item_Software_Component_ControlPanel{Session: r}
}

type Billing_Item_Software_Component_ControlPanel_Parallels_Plesk_Billing struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemSoftwareComponentControlPanelParallelsPleskBillingService() Billing_Item_Software_Component_ControlPanel_Parallels_Plesk_Billing {
	return Billing_Item_Software_Component_ControlPanel_Parallels_Plesk_Billing{Session: r}
}

type Billing_Item_Software_Component_OperatingSystem_Addon struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemSoftwareComponentOperatingSystemAddonService() Billing_Item_Software_Component_OperatingSystem_Addon {
	return Billing_Item_Software_Component_OperatingSystem_Addon{Session: r}
}

type Billing_Item_Software_Component_OperatingSystem_Addon_Citrix_Essentials struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemSoftwareComponentOperatingSystemAddonCitrixEssentialsService() Billing_Item_Software_Component_OperatingSystem_Addon_Citrix_Essentials {
	return Billing_Item_Software_Component_OperatingSystem_Addon_Citrix_Essentials{Session: r}
}

func (r *Billing_Item_Software_Component_OperatingSystem_Addon_Citrix_Essentials) GetResource() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Software_Component_Virtual_OperatingSystem struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemSoftwareComponentVirtualOperatingSystemService() Billing_Item_Software_Component_Virtual_OperatingSystem {
	return Billing_Item_Software_Component_Virtual_OperatingSystem{Session: r}
}

type Billing_Item_Software_Component_Virtual_OperatingSystem_Microsoft struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemSoftwareComponentVirtualOperatingSystemMicrosoftService() Billing_Item_Software_Component_Virtual_OperatingSystem_Microsoft {
	return Billing_Item_Software_Component_Virtual_OperatingSystem_Microsoft{Session: r}
}

func (r *Billing_Item_Software_Component_Virtual_OperatingSystem_Microsoft) GetResource() (resp datatypes.Software_VirtualLicense, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Software_Component_Virtual_OperatingSystem_Redhat struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemSoftwareComponentVirtualOperatingSystemRedhatService() Billing_Item_Software_Component_Virtual_OperatingSystem_Redhat {
	return Billing_Item_Software_Component_Virtual_OperatingSystem_Redhat{Session: r}
}

func (r *Billing_Item_Software_Component_Virtual_OperatingSystem_Redhat) GetResource() (resp datatypes.Software_Component, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Software_License struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemSoftwareLicenseService() Billing_Item_Software_License {
	return Billing_Item_Software_License{Session: r}
}

func (r *Billing_Item_Software_License) GetResource() (resp datatypes.Software_AccountLicense, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Support struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemSupportService() Billing_Item_Support {
	return Billing_Item_Support{Session: r}
}

type Billing_Item_User_Customer_External_Binding struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemUserCustomerExternalBindingService() Billing_Item_User_Customer_External_Binding {
	return Billing_Item_User_Customer_External_Binding{Session: r}
}

func (r *Billing_Item_User_Customer_External_Binding) GetResource() (resp datatypes.User_Customer_External_Binding, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Virtual_Dedicated_Rack struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemVirtualDedicatedRackService() Billing_Item_Virtual_Dedicated_Rack {
	return Billing_Item_Virtual_Dedicated_Rack{Session: r}
}

func (r *Billing_Item_Virtual_Dedicated_Rack) GetBillingCycleBandwidthUsage() (resp []datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Virtual_Dedicated_Rack) GetBillingCyclePrivateBandwidthUsage() (resp []datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Virtual_Dedicated_Rack) GetBillingCyclePrivateUsageIn() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Virtual_Dedicated_Rack) GetBillingCyclePrivateUsageOut() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Virtual_Dedicated_Rack) GetBillingCyclePrivateUsageTotal() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Virtual_Dedicated_Rack) GetBillingCyclePublicBandwidthUsage() (resp []datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Virtual_Dedicated_Rack) GetBillingCyclePublicUsageIn() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Virtual_Dedicated_Rack) GetBillingCyclePublicUsageOut() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Virtual_Dedicated_Rack) GetBillingCyclePublicUsageTotal() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Virtual_Dedicated_Rack) GetResource() (resp datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Virtual_Disk_Image struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemVirtualDiskImageService() Billing_Item_Virtual_Disk_Image {
	return Billing_Item_Virtual_Disk_Image{Session: r}
}

func (r *Billing_Item_Virtual_Disk_Image) GetResource() (resp datatypes.Virtual_Disk_Image, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Virtual_Guest struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemVirtualGuestService() Billing_Item_Virtual_Guest {
	return Billing_Item_Virtual_Guest{Session: r}
}

func (r *Billing_Item_Virtual_Guest) GetBillingCycleBandwidthUsage() (resp []datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Virtual_Guest) GetBillingCyclePrivateBandwidthUsage() (resp []datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Virtual_Guest) GetBillingCyclePrivateUsageIn() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Virtual_Guest) GetBillingCyclePrivateUsageOut() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Virtual_Guest) GetBillingCyclePrivateUsageTotal() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Virtual_Guest) GetBillingCyclePublicBandwidthUsage() (resp []datatypes.Network_Bandwidth_Usage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Virtual_Guest) GetBillingCyclePublicUsageIn() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Virtual_Guest) GetBillingCyclePublicUsageOut() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Virtual_Guest) GetBillingCyclePublicUsageTotal() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Virtual_Guest) GetMonitoringBillingItems() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Item_Virtual_Guest) GetResource() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Virtual_Host_Usage struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemVirtualHostUsageService() Billing_Item_Virtual_Host_Usage {
	return Billing_Item_Virtual_Host_Usage{Session: r}
}

func (r *Billing_Item_Virtual_Host_Usage) GetResource() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Item_Workspace struct {
	Session *Session
	Options
}

func (r *Session) GetBillingItemWorkspaceService() Billing_Item_Workspace {
	return Billing_Item_Workspace{Session: r}
}

type Billing_Order struct {
	Session *Session
	Options
}

func (r *Session) GetBillingOrderService() Billing_Order {
	return Billing_Order{Session: r}
}

func (r *Billing_Order) ApproveModifiedOrder() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetAllObjects() (resp []datatypes.Billing_Order, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetObject() (resp datatypes.Billing_Order, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetOrderStatuses() (resp []datatypes.Container_Billing_Order_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetPdf() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetPdfFilename() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetRecalculatedOrderContainer(message *string, ignoreDiscountsFlag *bool) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		message,
		ignoreDiscountsFlag,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetReceipt() (resp datatypes.Container_Product_Order_Receipt, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) IsPendingEditApproval() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Billing_Order) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetBrand() (resp datatypes.Brand, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetCart() (resp datatypes.Billing_Order_Cart, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetCoreRestrictedItems() (resp []datatypes.Billing_Order_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetCreditCardTransactions() (resp []datatypes.Billing_Payment_Card_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetExchangeRate() (resp datatypes.Billing_Currency_ExchangeRate, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetInitialInvoice() (resp datatypes.Billing_Invoice, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetItems() (resp []datatypes.Billing_Order_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetOrderApprovalDate() (resp datatypes.Time, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetOrderNonServerMonthlyAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetOrderServerMonthlyAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetOrderTopLevelItems() (resp []datatypes.Billing_Order_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetOrderTotalAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetOrderTotalOneTime() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetOrderTotalOneTimeAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetOrderTotalOneTimeTaxAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetOrderTotalRecurring() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetOrderTotalRecurringAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetOrderTotalRecurringTaxAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetOrderTotalSetupAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetOrderType() (resp datatypes.Billing_Order_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetPaypalTransactions() (resp []datatypes.Billing_Payment_PayPal_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetPresaleEvent() (resp datatypes.Sales_Presale_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetQuote() (resp datatypes.Billing_Order_Quote, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetReferralPartner() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetUpgradeRequestFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order) GetUserRecord() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Order_Cart struct {
	Session *Session
	Options
}

func (r *Session) GetBillingOrderCartService() Billing_Order_Cart {
	return Billing_Order_Cart{Session: r}
}

func (r *Billing_Order_Cart) CreateCart(orderData *datatypes.Container_Product_Order) (resp int, err error) {
	params := []interface{}{
		orderData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Cart) DeleteCart() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Cart) GetCartByCartKey(cartKey *string) (resp datatypes.Billing_Order_Cart, err error) {
	params := []interface{}{
		cartKey,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Cart) GetObject() (resp datatypes.Billing_Order_Cart, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Cart) GetPdf() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Cart) GetRecalculatedOrderContainer(orderData *datatypes.Container_Product_Order, orderBeingPlacedFlag *bool) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		orderData,
		orderBeingPlacedFlag,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Cart) UpdateCart(orderData *datatypes.Container_Product_Order) (resp int, err error) {
	params := []interface{}{
		orderData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Billing_Order_Item struct {
	Session *Session
	Options
}

func (r *Session) GetBillingOrderItemService() Billing_Order_Item {
	return Billing_Order_Item{Session: r}
}

func (r *Billing_Order_Item) GetObject() (resp datatypes.Billing_Order_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Billing_Order_Item) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Item) GetBundledItems() (resp []datatypes.Billing_Order_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Item) GetCategory() (resp datatypes.Product_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Item) GetChildren() (resp []datatypes.Billing_Order_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Item) GetGlobalIdentifier() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Item) GetHardwareGenericComponent() (resp datatypes.Hardware_Component_Model_Generic, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Item) GetItem() (resp datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Item) GetItemCategoryAnswers() (resp []datatypes.Billing_Order_Item_Category_Answer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Item) GetItemPrice() (resp datatypes.Product_Item_Price, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Item) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Item) GetNextOrderChildren() (resp []datatypes.Billing_Order_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Item) GetOldBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Item) GetOrder() (resp datatypes.Billing_Order, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Item) GetOrderApprovalDate() (resp datatypes.Time, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Item) GetPackage() (resp datatypes.Product_Package, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Item) GetParent() (resp datatypes.Billing_Order_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Item) GetRedundantPowerSupplyCount() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Item) GetSoftwareDescription() (resp datatypes.Software_Description, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Item) GetStorageGroups() (resp []datatypes.Configuration_Storage_Group_Order, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Item) GetTotalRecurringAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Item) GetUpgradeItem() (resp datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Order_Item_Category_Answer struct {
	Session *Session
	Options
}

func (r *Session) GetBillingOrderItemCategoryAnswerService() Billing_Order_Item_Category_Answer {
	return Billing_Order_Item_Category_Answer{Session: r}
}

func (r *Billing_Order_Item_Category_Answer) GetOrderItem() (resp datatypes.Billing_Order_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Item_Category_Answer) GetQuestion() (resp datatypes.Product_Item_Category_Question, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Order_Note struct {
	Session *Session
	Options
}

func (r *Session) GetBillingOrderNoteService() Billing_Order_Note {
	return Billing_Order_Note{Session: r}
}

func (r *Billing_Order_Note) GetEmployee() (resp datatypes.User_Employee, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Note) GetOrder() (resp datatypes.Billing_Order, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Order_Quote struct {
	Session *Session
	Options
}

func (r *Session) GetBillingOrderQuoteService() Billing_Order_Quote {
	return Billing_Order_Quote{Session: r}
}

func (r *Billing_Order_Quote) Claim(quoteKey *string, quoteId *int) (resp datatypes.Billing_Order_Quote, err error) {
	params := []interface{}{
		quoteKey,
		quoteId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Quote) DeleteQuote() (resp datatypes.Billing_Order_Quote, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Quote) GetObject() (resp datatypes.Billing_Order_Quote, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Quote) GetPdf() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Quote) GetQuoteByQuoteKey(quoteKey *string) (resp datatypes.Billing_Order_Quote, err error) {
	params := []interface{}{
		quoteKey,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Quote) GetRecalculatedOrderContainer(orderData *datatypes.Container_Product_Order, orderBeingPlacedFlag *bool) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		orderData,
		orderBeingPlacedFlag,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Quote) PlaceOrder(orderData *datatypes.Container_Product_Order) (resp datatypes.Container_Product_Order_Receipt, err error) {
	params := []interface{}{
		orderData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Quote) PlaceQuote(orderData *datatypes.Container_Product_Order) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		orderData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Quote) SaveQuote() (resp datatypes.Billing_Order_Quote, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Quote) VerifyOrder(orderData *datatypes.Container_Product_Order) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		orderData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

func (r *Billing_Order_Quote) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Quote) GetOrder() (resp datatypes.Billing_Order, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Order_Quote) GetOrdersFromQuote() (resp []datatypes.Billing_Order, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Order_Type struct {
	Session *Session
	Options
}

func (r *Session) GetBillingOrderTypeService() Billing_Order_Type {
	return Billing_Order_Type{Session: r}
}

type Billing_Payment_Card_ChangeRequest struct {
	Session *Session
	Options
}

func (r *Session) GetBillingPaymentCardChangeRequestService() Billing_Payment_Card_ChangeRequest {
	return Billing_Payment_Card_ChangeRequest{Session: r}
}

func (r *Billing_Payment_Card_ChangeRequest) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Payment_Card_ChangeRequest) GetAuthorizedCreditCardTransaction() (resp datatypes.Billing_Payment_Card_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Payment_Card_ChangeRequest) GetCaptureCreditCardTransaction() (resp datatypes.Billing_Payment_Card_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Payment_Card_ChangeRequest) GetTicketAttachmentReferences() (resp []datatypes.Ticket_Attachment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Payment_Card_ManualPayment struct {
	Session *Session
	Options
}

func (r *Session) GetBillingPaymentCardManualPaymentService() Billing_Payment_Card_ManualPayment {
	return Billing_Payment_Card_ManualPayment{Session: r}
}

func (r *Billing_Payment_Card_ManualPayment) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Payment_Card_ManualPayment) GetAuthorizedCreditCardTransaction() (resp datatypes.Billing_Payment_Card_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Payment_Card_ManualPayment) GetAuthorizedPayPalTransaction() (resp datatypes.Billing_Payment_PayPal_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Payment_Card_ManualPayment) GetCaptureCreditCardTransaction() (resp datatypes.Billing_Payment_Card_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Payment_Card_ManualPayment) GetCapturePayPalTransaction() (resp datatypes.Billing_Payment_PayPal_Transaction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Payment_Card_ManualPayment) GetTicketAttachmentReferences() (resp []datatypes.Ticket_Attachment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Payment_Card_Transaction struct {
	Session *Session
	Options
}

func (r *Session) GetBillingPaymentCardTransactionService() Billing_Payment_Card_Transaction {
	return Billing_Payment_Card_Transaction{Session: r}
}

func (r *Billing_Payment_Card_Transaction) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Payment_Card_Transaction) GetOrder() (resp datatypes.Billing_Order, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Payment_PayPal_Transaction struct {
	Session *Session
	Options
}

func (r *Session) GetBillingPaymentPayPalTransactionService() Billing_Payment_PayPal_Transaction {
	return Billing_Payment_PayPal_Transaction{Session: r}
}

func (r *Billing_Payment_PayPal_Transaction) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Payment_PayPal_Transaction) GetOrder() (resp datatypes.Billing_Order, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Payment_Processor struct {
	Session *Session
	Options
}

func (r *Session) GetBillingPaymentProcessorService() Billing_Payment_Processor {
	return Billing_Payment_Processor{Session: r}
}

func (r *Billing_Payment_Processor) GetBrandAssignments() (resp []datatypes.Brand_Payment_Processor, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Payment_Processor) GetOwnerAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Payment_Processor) GetPaymentMethods() (resp []datatypes.Billing_Payment_Processor_Method, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Payment_Processor) GetType() (resp datatypes.Billing_Payment_Processor_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Payment_Processor_Method struct {
	Session *Session
	Options
}

func (r *Session) GetBillingPaymentProcessorMethodService() Billing_Payment_Processor_Method {
	return Billing_Payment_Processor_Method{Session: r}
}

func (r *Billing_Payment_Processor_Method) GetPaymentProcessor() (resp datatypes.Billing_Payment_Processor, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Billing_Payment_Processor_Method) GetPaymentType() (resp datatypes.Billing_Payment_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Payment_Processor_Type struct {
	Session *Session
	Options
}

func (r *Session) GetBillingPaymentProcessorTypeService() Billing_Payment_Processor_Type {
	return Billing_Payment_Processor_Type{Session: r}
}

func (r *Billing_Payment_Processor_Type) GetPaymentProcessors() (resp []datatypes.Billing_Payment_Processor, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Billing_Payment_Transaction struct {
	Session *Session
	Options
}

func (r *Session) GetBillingPaymentTransactionService() Billing_Payment_Transaction {
	return Billing_Payment_Transaction{Session: r}
}

type Billing_Payment_Type struct {
	Session *Session
	Options
}

func (r *Session) GetBillingPaymentTypeService() Billing_Payment_Type {
	return Billing_Payment_Type{Session: r}
}
