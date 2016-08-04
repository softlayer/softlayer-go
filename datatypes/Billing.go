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

package datatypes

import "time"

type Billing_Currency struct {
	Entity

	Id      *int    `json:"id,omitempty"`
	KeyName *string `json:"keyName,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type Billing_Currency_ExchangeRate struct {
	Entity

	EffectiveDate   *time.Time        `json:"effectiveDate,omitempty"`
	ExpirationDate  *time.Time        `json:"expirationDate,omitempty"`
	FundingCurrency *Billing_Currency `json:"fundingCurrency,omitempty"`
	Id              *int              `json:"id,omitempty"`
	LocalCurrency   *Billing_Currency `json:"localCurrency,omitempty"`
	Rate            *float64          `json:"rate,omitempty"`
}

type Billing_Info struct {
	Entity

	Account                   *Account            `json:"account,omitempty"`
	AccountId                 *int                `json:"accountId,omitempty"`
	AchInformation            []Billing_Info_Ach  `json:"achInformation,omitempty"`
	AchInformationCount       *uint               `json:"achInformationCount,omitempty"`
	AnniversaryDayOfMonth     *int                `json:"anniversaryDayOfMonth,omitempty"`
	CardAccountNumber         *string             `json:"cardAccountNumber,omitempty"`
	CardExpirationMonth       *int                `json:"cardExpirationMonth,omitempty"`
	CardExpirationYear        *int                `json:"cardExpirationYear,omitempty"`
	CardNickname              *string             `json:"cardNickname,omitempty"`
	CardType                  *string             `json:"cardType,omitempty"`
	CardVerificationNumber    *string             `json:"cardVerificationNumber,omitempty"`
	CreateDate                *time.Time          `json:"createDate,omitempty"`
	Currency                  *Billing_Currency   `json:"currency,omitempty"`
	CurrentBillingCycle       *Billing_Info_Cycle `json:"currentBillingCycle,omitempty"`
	Id                        *int                `json:"id,omitempty"`
	LastBillDate              *time.Time          `json:"lastBillDate,omitempty"`
	LastFourPaymentCardDigits *int                `json:"lastFourPaymentCardDigits,omitempty"`
	LastPaymentDate           *time.Time          `json:"lastPaymentDate,omitempty"`
	ModifyDate                *time.Time          `json:"modifyDate,omitempty"`
	NextBillDate              *time.Time          `json:"nextBillDate,omitempty"`
	PaymentTerms              *int                `json:"paymentTerms,omitempty"`
	PercentDiscountOnetime    *int                `json:"percentDiscountOnetime,omitempty"`
	PercentDiscountRecurring  *int                `json:"percentDiscountRecurring,omitempty"`
	SparePoolAmount           *int                `json:"sparePoolAmount,omitempty"`
	VatId                     *string             `json:"vatId,omitempty"`
}

type Billing_Info_Ach struct {
	Entity

	Account           *Account   `json:"account,omitempty"`
	AccountId         *int       `json:"accountId,omitempty"`
	AccountNumber     *string    `json:"accountNumber,omitempty"`
	AccountType       *string    `json:"accountType,omitempty"`
	BankTransitNumber *string    `json:"bankTransitNumber,omitempty"`
	City              *string    `json:"city,omitempty"`
	Country           *string    `json:"country,omitempty"`
	FirstName         *string    `json:"firstName,omitempty"`
	Id                *int       `json:"id,omitempty"`
	LastName          *string    `json:"lastName,omitempty"`
	PhoneNumber       *string    `json:"phoneNumber,omitempty"`
	Postalcode        *string    `json:"postalcode,omitempty"`
	State             *string    `json:"state,omitempty"`
	Status            *string    `json:"status,omitempty"`
	Street1           *string    `json:"street1,omitempty"`
	Street2           *string    `json:"street2,omitempty"`
	VerifiedDate      *time.Time `json:"verifiedDate,omitempty"`
}

type Billing_Info_Cycle struct {
	Entity

	Account                *Account   `json:"account,omitempty"`
	CurrentCycleEndDate    *time.Time `json:"currentCycleEndDate,omitempty"`
	CurrentCycleStartDate  *time.Time `json:"currentCycleStartDate,omitempty"`
	NextCycleStartDate     *time.Time `json:"nextCycleStartDate,omitempty"`
	PreviousCycleEndDate   *time.Time `json:"previousCycleEndDate,omitempty"`
	PreviousCycleStartDate *time.Time `json:"previousCycleStartDate,omitempty"`
}

type Billing_Invoice struct {
	Entity

	Account                        *Account                             `json:"account,omitempty"`
	AccountId                      *int                                 `json:"accountId,omitempty"`
	Address1                       *string                              `json:"address1,omitempty"`
	Address2                       *string                              `json:"address2,omitempty"`
	Amount                         *float64                             `json:"amount,omitempty"`
	BrandAtInvoiceCreation         *Brand                               `json:"brandAtInvoiceCreation,omitempty"`
	City                           *string                              `json:"city,omitempty"`
	ClaimedTaxExemptTxFlag         *bool                                `json:"claimedTaxExemptTxFlag,omitempty"`
	ClosedDate                     *time.Time                           `json:"closedDate,omitempty"`
	CompanyName                    *string                              `json:"companyName,omitempty"`
	Country                        *string                              `json:"country,omitempty"`
	CreateDate                     *time.Time                           `json:"createDate,omitempty"`
	DetailedPdfGeneratedFlag       *bool                                `json:"detailedPdfGeneratedFlag,omitempty"`
	DocumentsGeneratedFlag         *bool                                `json:"documentsGeneratedFlag,omitempty"`
	Email                          *string                              `json:"email,omitempty"`
	EndingBalance                  *float64                             `json:"endingBalance,omitempty"`
	FaxPhone                       *string                              `json:"faxPhone,omitempty"`
	FirstName                      *string                              `json:"firstName,omitempty"`
	Id                             *int                                 `json:"id,omitempty"`
	InvoiceTopLevelItemCount       *uint                                `json:"invoiceTopLevelItemCount,omitempty"`
	InvoiceTopLevelItems           []Billing_Invoice_Item               `json:"invoiceTopLevelItems,omitempty"`
	InvoiceTotalAmount             *float64                             `json:"invoiceTotalAmount,omitempty"`
	InvoiceTotalOneTimeAmount      *float64                             `json:"invoiceTotalOneTimeAmount,omitempty"`
	InvoiceTotalOneTimeTaxAmount   *float64                             `json:"invoiceTotalOneTimeTaxAmount,omitempty"`
	InvoiceTotalPreTaxAmount       *float64                             `json:"invoiceTotalPreTaxAmount,omitempty"`
	InvoiceTotalRecurringAmount    *float64                             `json:"invoiceTotalRecurringAmount,omitempty"`
	InvoiceTotalRecurringTaxAmount *float64                             `json:"invoiceTotalRecurringTaxAmount,omitempty"`
	ItemCount                      *uint                                `json:"itemCount,omitempty"`
	Items                          []Billing_Invoice_Item               `json:"items,omitempty"`
	LastName                       *string                              `json:"lastName,omitempty"`
	ModifyDate                     *time.Time                           `json:"modifyDate,omitempty"`
	OfficePhone                    *string                              `json:"officePhone,omitempty"`
	Payment                        *float64                             `json:"payment,omitempty"`
	PaymentCount                   *uint                                `json:"paymentCount,omitempty"`
	Payments                       []Billing_Invoice_Receivable_Payment `json:"payments,omitempty"`
	PostalCode                     *string                              `json:"postalCode,omitempty"`
	PurchaseOrderNumber            *string                              `json:"purchaseOrderNumber,omitempty"`
	SellerRegistration             *string                              `json:"sellerRegistration,omitempty"`
	StartingBalance                *float64                             `json:"startingBalance,omitempty"`
	State                          *string                              `json:"state,omitempty"`
	StatusCode                     *string                              `json:"statusCode,omitempty"`
	TaxInfo                        *Billing_Invoice_Tax_Info            `json:"taxInfo,omitempty"`
	TaxInfoHistory                 []Billing_Invoice_Tax_Info           `json:"taxInfoHistory,omitempty"`
	TaxInfoHistoryCount            *uint                                `json:"taxInfoHistoryCount,omitempty"`
	TaxMessage                     *string                              `json:"taxMessage,omitempty"`
	TaxStatusId                    *int                                 `json:"taxStatusId,omitempty"`
	TaxType                        *Billing_Invoice_Tax_Type            `json:"taxType,omitempty"`
	TaxTypeId                      *int                                 `json:"taxTypeId,omitempty"`
	TypeCode                       *string                              `json:"typeCode,omitempty"`
}

type Billing_Invoice_Item struct {
	Entity

	AssociatedChildren              []Billing_Invoice_Item `json:"associatedChildren,omitempty"`
	AssociatedChildrenCount         *uint                  `json:"associatedChildrenCount,omitempty"`
	AssociatedInvoiceItem           *Billing_Invoice_Item  `json:"associatedInvoiceItem,omitempty"`
	AssociatedInvoiceItemId         *int                   `json:"associatedInvoiceItemId,omitempty"`
	BillingItem                     *Billing_Item          `json:"billingItem,omitempty"`
	BillingItemId                   *int                   `json:"billingItemId,omitempty"`
	Category                        *Product_Item_Category `json:"category,omitempty"`
	CategoryCode                    *string                `json:"categoryCode,omitempty"`
	Children                        []Billing_Invoice_Item `json:"children,omitempty"`
	ChildrenCount                   *uint                  `json:"childrenCount,omitempty"`
	CreateDate                      *time.Time             `json:"createDate,omitempty"`
	Description                     *string                `json:"description,omitempty"`
	DomainName                      *string                `json:"domainName,omitempty"`
	FilteredAssociatedChildren      []Billing_Invoice_Item `json:"filteredAssociatedChildren,omitempty"`
	FilteredAssociatedChildrenCount *uint                  `json:"filteredAssociatedChildrenCount,omitempty"`
	HostName                        *string                `json:"hostName,omitempty"`
	HourlyRecurringFee              *float64               `json:"hourlyRecurringFee,omitempty"`
	Id                              *int                   `json:"id,omitempty"`
	Invoice                         *Billing_Invoice       `json:"invoice,omitempty"`
	InvoiceId                       *int                   `json:"invoiceId,omitempty"`
	LaborAfterTaxAmount             *float64               `json:"laborAfterTaxAmount,omitempty"`
	LaborFee                        *float64               `json:"laborFee,omitempty"`
	LaborFeeTaxRate                 *float64               `json:"laborFeeTaxRate,omitempty"`
	LaborTaxAmount                  *float64               `json:"laborTaxAmount,omitempty"`
	Location                        *Location              `json:"location,omitempty"`
	NonZeroAssociatedChildren       []Billing_Invoice_Item `json:"nonZeroAssociatedChildren,omitempty"`
	NonZeroAssociatedChildrenCount  *uint                  `json:"nonZeroAssociatedChildrenCount,omitempty"`
	Notes                           *string                `json:"notes,omitempty"`
	OneTimeAfterTaxAmount           *float64               `json:"oneTimeAfterTaxAmount,omitempty"`
	OneTimeFee                      *float64               `json:"oneTimeFee,omitempty"`
	OneTimeFeeTaxRate               *float64               `json:"oneTimeFeeTaxRate,omitempty"`
	OneTimeTaxAmount                *float64               `json:"oneTimeTaxAmount,omitempty"`
	Parent                          *Billing_Invoice_Item  `json:"parent,omitempty"`
	ParentId                        *int                   `json:"parentId,omitempty"`
	Product                         *Product_Item          `json:"product,omitempty"`
	ProductItemId                   *int                   `json:"productItemId,omitempty"`
	RecurringAfterTaxAmount         *float64               `json:"recurringAfterTaxAmount,omitempty"`
	RecurringFee                    *float64               `json:"recurringFee,omitempty"`
	RecurringFeeTaxRate             *float64               `json:"recurringFeeTaxRate,omitempty"`
	RecurringTaxAmount              *float64               `json:"recurringTaxAmount,omitempty"`
	ResourceTableId                 *int                   `json:"resourceTableId,omitempty"`
	SetupAfterTaxAmount             *float64               `json:"setupAfterTaxAmount,omitempty"`
	SetupFee                        *float64               `json:"setupFee,omitempty"`
	SetupFeeTaxRate                 *float64               `json:"setupFeeTaxRate,omitempty"`
	SetupTaxAmount                  *float64               `json:"setupTaxAmount,omitempty"`
	TotalOneTimeAmount              *float64               `json:"totalOneTimeAmount,omitempty"`
	TotalOneTimeTaxAmount           *float64               `json:"totalOneTimeTaxAmount,omitempty"`
	TotalRecurringAmount            *float64               `json:"totalRecurringAmount,omitempty"`
	TotalRecurringTaxAmount         *float64               `json:"totalRecurringTaxAmount,omitempty"`
}

type Billing_Invoice_Item_Hardware struct {
	Billing_Invoice_Item

	Resource *Hardware `json:"resource,omitempty"`
}

type Billing_Invoice_Item_Tax_Info struct {
	Entity

	CreateDate          *time.Time                `json:"createDate,omitempty"`
	Description         *string                   `json:"description,omitempty"`
	EffectiveTaxRate    *float64                  `json:"effectiveTaxRate,omitempty"`
	ExemptAmount        *float64                  `json:"exemptAmount,omitempty"`
	FeeProperty         *string                   `json:"feeProperty,omitempty"`
	Id                  *int                      `json:"id,omitempty"`
	InvoiceItem         *Billing_Invoice_Item     `json:"invoiceItem,omitempty"`
	InvoiceItemId       *int                      `json:"invoiceItemId,omitempty"`
	InvoiceTaxInfo      *Billing_Invoice_Tax_Info `json:"invoiceTaxInfo,omitempty"`
	InvoiceTaxInfoId    *int                      `json:"invoiceTaxInfoId,omitempty"`
	ModifyDate          *time.Time                `json:"modifyDate,omitempty"`
	NonTaxableBasis     *float64                  `json:"nonTaxableBasis,omitempty"`
	ReportedFlag        *bool                     `json:"reportedFlag,omitempty"`
	SellerRegistration  *string                   `json:"sellerRegistration,omitempty"`
	TaxAmount           *float64                  `json:"taxAmount,omitempty"`
	TaxAmountToCurrency *float64                  `json:"taxAmountToCurrency,omitempty"`
	TaxRate             *float64                  `json:"taxRate,omitempty"`
	TaxableBasis        *float64                  `json:"taxableBasis,omitempty"`
	ToCurrency          *Billing_Currency         `json:"toCurrency,omitempty"`
	ToCurrencyId        *int                      `json:"toCurrencyId,omitempty"`
}

type Billing_Invoice_Next struct {
	Entity
}

type Billing_Invoice_Receivable_Payment struct {
	Entity

	Account                  *Account                            `json:"account,omitempty"`
	Amount                   *float64                            `json:"amount,omitempty"`
	CreateDate               *time.Time                          `json:"createDate,omitempty"`
	CreditCardLastFourDigits *int                                `json:"creditCardLastFourDigits,omitempty"`
	CreditCardRequestId      *string                             `json:"creditCardRequestId,omitempty"`
	CreditCardTransaction    *Billing_Payment_Card_Transaction   `json:"creditCardTransaction,omitempty"`
	ExchangeRate             *Billing_Currency_ExchangeRate      `json:"exchangeRate,omitempty"`
	Invoice                  *Billing_Invoice                    `json:"invoice,omitempty"`
	InvoiceId                *int                                `json:"invoiceId,omitempty"`
	PaypalTransaction        *Billing_Payment_PayPal_Transaction `json:"paypalTransaction,omitempty"`
	TypeCode                 *string                             `json:"typeCode,omitempty"`
}

type Billing_Invoice_Tax_Info struct {
	Entity

	CreateDate               *time.Time                      `json:"createDate,omitempty"`
	Currency                 *Billing_Currency               `json:"currency,omitempty"`
	CurrencyId               *int                            `json:"currencyId,omitempty"`
	FunctionalCurrency       *Billing_Currency               `json:"functionalCurrency,omitempty"`
	Id                       *int                            `json:"id,omitempty"`
	Invoice                  *Billing_Invoice                `json:"invoice,omitempty"`
	InvoiceId                *int                            `json:"invoiceId,omitempty"`
	ItemCount                *uint                           `json:"itemCount,omitempty"`
	ItemWithCurrencyInfo     *Billing_Invoice_Item_Tax_Info  `json:"itemWithCurrencyInfo,omitempty"`
	Items                    []Billing_Invoice_Item_Tax_Info `json:"items,omitempty"`
	ModifyDate               *time.Time                      `json:"modifyDate,omitempty"`
	ReportedFlag             *bool                           `json:"reportedFlag,omitempty"`
	TotalTaxAmountToCurrency *float64                        `json:"totalTaxAmountToCurrency,omitempty"`
}

type Billing_Invoice_Tax_Status struct {
	Entity

	CreateDate *time.Time `json:"createDate,omitempty"`
	Id         *int       `json:"id,omitempty"`
	KeyName    *string    `json:"keyName,omitempty"`
	ModifyDate *time.Time `json:"modifyDate,omitempty"`
	Name       *string    `json:"name,omitempty"`
}

type Billing_Invoice_Tax_Type struct {
	Entity

	Id      *int    `json:"id,omitempty"`
	KeyName *string `json:"keyName,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type Billing_Item struct {
	Entity

	Account                                            *Account                                `json:"account,omitempty"`
	ActiveAgreement                                    *Account_Agreement                      `json:"activeAgreement,omitempty"`
	ActiveAgreementFlag                                *Account_Agreement                      `json:"activeAgreementFlag,omitempty"`
	ActiveAssociatedChildren                           []Billing_Item                          `json:"activeAssociatedChildren,omitempty"`
	ActiveAssociatedChildrenCount                      *uint                                   `json:"activeAssociatedChildrenCount,omitempty"`
	ActiveAssociatedGuestDiskBillingItemCount          *uint                                   `json:"activeAssociatedGuestDiskBillingItemCount,omitempty"`
	ActiveAssociatedGuestDiskBillingItems              []Billing_Item                          `json:"activeAssociatedGuestDiskBillingItems,omitempty"`
	ActiveBundledItemCount                             *uint                                   `json:"activeBundledItemCount,omitempty"`
	ActiveBundledItems                                 []Billing_Item                          `json:"activeBundledItems,omitempty"`
	ActiveCancellationItem                             *Billing_Item_Cancellation_Request_Item `json:"activeCancellationItem,omitempty"`
	ActiveChildren                                     []Billing_Item                          `json:"activeChildren,omitempty"`
	ActiveChildrenCount                                *uint                                   `json:"activeChildrenCount,omitempty"`
	ActiveFlag                                         *bool                                   `json:"activeFlag,omitempty"`
	ActiveSparePoolAssociatedGuestDiskBillingItemCount *uint                                   `json:"activeSparePoolAssociatedGuestDiskBillingItemCount,omitempty"`
	ActiveSparePoolAssociatedGuestDiskBillingItems     []Billing_Item                          `json:"activeSparePoolAssociatedGuestDiskBillingItems,omitempty"`
	ActiveSparePoolBundledItemCount                    *uint                                   `json:"activeSparePoolBundledItemCount,omitempty"`
	ActiveSparePoolBundledItems                        []Billing_Item                          `json:"activeSparePoolBundledItems,omitempty"`
	AllowCancellationFlag                              *int                                    `json:"allowCancellationFlag,omitempty"`
	AssociatedBillingItem                              *Billing_Item                           `json:"associatedBillingItem,omitempty"`
	AssociatedBillingItemHistory                       []Billing_Item_Association_History      `json:"associatedBillingItemHistory,omitempty"`
	AssociatedBillingItemHistoryCount                  *uint                                   `json:"associatedBillingItemHistoryCount,omitempty"`
	AssociatedBillingItemId                            *string                                 `json:"associatedBillingItemId,omitempty"`
	AssociatedChildren                                 []Billing_Item                          `json:"associatedChildren,omitempty"`
	AssociatedChildrenCount                            *uint                                   `json:"associatedChildrenCount,omitempty"`
	AssociatedParent                                   []Billing_Item                          `json:"associatedParent,omitempty"`
	AssociatedParentCount                              *uint                                   `json:"associatedParentCount,omitempty"`
	AvailableMatchingVlanCount                         *uint                                   `json:"availableMatchingVlanCount,omitempty"`
	AvailableMatchingVlans                             []Network_Vlan                          `json:"availableMatchingVlans,omitempty"`
	BandwidthAllocation                                *Network_Bandwidth_Version1_Allocation  `json:"bandwidthAllocation,omitempty"`
	BillableChildren                                   []Billing_Item                          `json:"billableChildren,omitempty"`
	BillableChildrenCount                              *uint                                   `json:"billableChildrenCount,omitempty"`
	BundleItemCount                                    *uint                                   `json:"bundleItemCount,omitempty"`
	BundleItems                                        []Product_Item_Bundles                  `json:"bundleItems,omitempty"`
	BundledItemCount                                   *uint                                   `json:"bundledItemCount,omitempty"`
	BundledItems                                       []Billing_Item                          `json:"bundledItems,omitempty"`
	CanceledChildren                                   []Billing_Item                          `json:"canceledChildren,omitempty"`
	CanceledChildrenCount                              *uint                                   `json:"canceledChildrenCount,omitempty"`
	CancellationDate                                   *time.Time                              `json:"cancellationDate,omitempty"`
	CancellationReason                                 *Billing_Item_Cancellation_Reason       `json:"cancellationReason,omitempty"`
	CancellationRequestCount                           *uint                                   `json:"cancellationRequestCount,omitempty"`
	CancellationRequests                               []Billing_Item_Cancellation_Request     `json:"cancellationRequests,omitempty"`
	Category                                           *Product_Item_Category                  `json:"category,omitempty"`
	CategoryCode                                       *string                                 `json:"categoryCode,omitempty"`
	Children                                           []Billing_Item                          `json:"children,omitempty"`
	ChildrenCount                                      *uint                                   `json:"childrenCount,omitempty"`
	ChildrenWithActiveAgreement                        []Billing_Item                          `json:"childrenWithActiveAgreement,omitempty"`
	ChildrenWithActiveAgreementCount                   *uint                                   `json:"childrenWithActiveAgreementCount,omitempty"`
	CreateDate                                         *time.Time                              `json:"createDate,omitempty"`
	CurrentHourlyCharge                                *string                                 `json:"currentHourlyCharge,omitempty"`
	CycleStartDate                                     *time.Time                              `json:"cycleStartDate,omitempty"`
	Description                                        *string                                 `json:"description,omitempty"`
	DomainName                                         *string                                 `json:"domainName,omitempty"`
	DowngradeItemCount                                 *uint                                   `json:"downgradeItemCount,omitempty"`
	DowngradeItems                                     []Product_Item                          `json:"downgradeItems,omitempty"`
	FilteredNextInvoiceChildren                        []Billing_Item                          `json:"filteredNextInvoiceChildren,omitempty"`
	FilteredNextInvoiceChildrenCount                   *uint                                   `json:"filteredNextInvoiceChildrenCount,omitempty"`
	HostName                                           *string                                 `json:"hostName,omitempty"`
	HourlyFlag                                         *bool                                   `json:"hourlyFlag,omitempty"`
	HourlyRecurringFee                                 *float64                                `json:"hourlyRecurringFee,omitempty"`
	HoursUsed                                          *string                                 `json:"hoursUsed,omitempty"`
	Id                                                 *int                                    `json:"id,omitempty"`
	InvoiceItem                                        *Billing_Invoice_Item                   `json:"invoiceItem,omitempty"`
	InvoiceItemCount                                   *uint                                   `json:"invoiceItemCount,omitempty"`
	InvoiceItems                                       []Billing_Invoice_Item                  `json:"invoiceItems,omitempty"`
	Item                                               *Product_Item                           `json:"item,omitempty"`
	LaborFee                                           *float64                                `json:"laborFee,omitempty"`
	LaborFeeTaxRate                                    *float64                                `json:"laborFeeTaxRate,omitempty"`
	LastBillDate                                       *time.Time                              `json:"lastBillDate,omitempty"`
	Location                                           *Location                               `json:"location,omitempty"`
	ModifyDate                                         *time.Time                              `json:"modifyDate,omitempty"`
	NextBillDate                                       *time.Time                              `json:"nextBillDate,omitempty"`
	NextInvoiceChildren                                []Billing_Item                          `json:"nextInvoiceChildren,omitempty"`
	NextInvoiceChildrenCount                           *uint                                   `json:"nextInvoiceChildrenCount,omitempty"`
	NextInvoiceTotalOneTimeAmount                      *float64                                `json:"nextInvoiceTotalOneTimeAmount,omitempty"`
	NextInvoiceTotalOneTimeTaxAmount                   *float64                                `json:"nextInvoiceTotalOneTimeTaxAmount,omitempty"`
	NextInvoiceTotalRecurringAmount                    *float64                                `json:"nextInvoiceTotalRecurringAmount,omitempty"`
	NextInvoiceTotalRecurringTaxAmount                 *float64                                `json:"nextInvoiceTotalRecurringTaxAmount,omitempty"`
	NonZeroNextInvoiceChildren                         []Billing_Item                          `json:"nonZeroNextInvoiceChildren,omitempty"`
	NonZeroNextInvoiceChildrenCount                    *uint                                   `json:"nonZeroNextInvoiceChildrenCount,omitempty"`
	Notes                                              *string                                 `json:"notes,omitempty"`
	OneTimeFee                                         *float64                                `json:"oneTimeFee,omitempty"`
	OneTimeFeeTaxRate                                  *float64                                `json:"oneTimeFeeTaxRate,omitempty"`
	OrderItem                                          *Billing_Order_Item                     `json:"orderItem,omitempty"`
	OrderItemId                                        *int                                    `json:"orderItemId,omitempty"`
	OriginalLocation                                   *Location                               `json:"originalLocation,omitempty"`
	Package                                            *Product_Package                        `json:"package,omitempty"`
	Parent                                             *Billing_Item                           `json:"parent,omitempty"`
	ParentId                                           *int                                    `json:"parentId,omitempty"`
	ParentVirtualGuestBillingItem                      *Billing_Item_Virtual_Guest             `json:"parentVirtualGuestBillingItem,omitempty"`
	PendingCancellationFlag                            *bool                                   `json:"pendingCancellationFlag,omitempty"`
	PendingOrderItem                                   *Billing_Order_Item                     `json:"pendingOrderItem,omitempty"`
	ProvisionTransaction                               *Provisioning_Version1_Transaction      `json:"provisionTransaction,omitempty"`
	RecurringFee                                       *float64                                `json:"recurringFee,omitempty"`
	RecurringFeeTaxRate                                *float64                                `json:"recurringFeeTaxRate,omitempty"`
	RecurringMonths                                    *int                                    `json:"recurringMonths,omitempty"`
	ServiceProviderId                                  *int                                    `json:"serviceProviderId,omitempty"`
	SetupFee                                           *float64                                `json:"setupFee,omitempty"`
	SetupFeeTaxRate                                    *float64                                `json:"setupFeeTaxRate,omitempty"`
	SoftwareDescription                                *Software_Description                   `json:"softwareDescription,omitempty"`
	UpgradeItem                                        *Product_Item                           `json:"upgradeItem,omitempty"`
	UpgradeItemCount                                   *uint                                   `json:"upgradeItemCount,omitempty"`
	UpgradeItems                                       []Product_Item                          `json:"upgradeItems,omitempty"`
}

type Billing_Item_Account_Media_Data_Transfer_Request struct {
	Billing_Item

	Resource *Account_Media_Data_Transfer_Request `json:"resource,omitempty"`
}

type Billing_Item_Association_History struct {
	Entity

	AssociatedBillingItem   *Billing_Item `json:"associatedBillingItem,omitempty"`
	AssociatedBillingItemId *int          `json:"associatedBillingItemId,omitempty"`
	BillingItem             *Billing_Item `json:"billingItem,omitempty"`
	BillingItemId           *int          `json:"billingItemId,omitempty"`
	CreateDate              *time.Time    `json:"createDate,omitempty"`
	Id                      *int          `json:"id,omitempty"`
}

type Billing_Item_Cancellation_Reason struct {
	Entity

	BillingCancelReasonCategoryId     *int                                       `json:"billingCancelReasonCategoryId,omitempty"`
	BillingCancellationReasonCategory *Billing_Item_Cancellation_Reason_Category `json:"billingCancellationReasonCategory,omitempty"`
	BillingItemCount                  *uint                                      `json:"billingItemCount,omitempty"`
	BillingItems                      []Billing_Item                             `json:"billingItems,omitempty"`
	Id                                *int                                       `json:"id,omitempty"`
	KeyName                           *string                                    `json:"keyName,omitempty"`
	Reason                            *string                                    `json:"reason,omitempty"`
	TranslatedReason                  *string                                    `json:"translatedReason,omitempty"`
}

type Billing_Item_Cancellation_Reason_Category struct {
	Entity

	BillingCancellationReasonCount *uint                              `json:"billingCancellationReasonCount,omitempty"`
	BillingCancellationReasons     []Billing_Item_Cancellation_Reason `json:"billingCancellationReasons,omitempty"`
	Id                             *int                               `json:"id,omitempty"`
	Name                           *string                            `json:"name,omitempty"`
}

type Billing_Item_Cancellation_Request struct {
	Entity

	Account               *Account                                  `json:"account,omitempty"`
	AccountId             *int                                      `json:"accountId,omitempty"`
	BillingCancelReasonId *int                                      `json:"billingCancelReasonId,omitempty"`
	CreateDate            *time.Time                                `json:"createDate,omitempty"`
	Id                    *int                                      `json:"id,omitempty"`
	ItemCount             *uint                                     `json:"itemCount,omitempty"`
	Items                 []Billing_Item_Cancellation_Request_Item  `json:"items,omitempty"`
	ModifyDate            *time.Time                                `json:"modifyDate,omitempty"`
	Notes                 *string                                   `json:"notes,omitempty"`
	Status                *Billing_Item_Cancellation_Request_Status `json:"status,omitempty"`
	StatusId              *int                                      `json:"statusId,omitempty"`
	Ticket                *Ticket                                   `json:"ticket,omitempty"`
	TicketId              *int                                      `json:"ticketId,omitempty"`
	User                  *User_Customer                            `json:"user,omitempty"`
}

type Billing_Item_Cancellation_Request_Item struct {
	Entity

	BillingItem               *Billing_Item                      `json:"billingItem,omitempty"`
	BillingItemId             *int                               `json:"billingItemId,omitempty"`
	CancellationRequest       *Billing_Item_Cancellation_Request `json:"cancellationRequest,omitempty"`
	CancellationRequestId     *int                               `json:"cancellationRequestId,omitempty"`
	Id                        *int                               `json:"id,omitempty"`
	ImmediateCancellationFlag *bool                              `json:"immediateCancellationFlag,omitempty"`
	ScheduledCancellationDate *time.Time                         `json:"scheduledCancellationDate,omitempty"`
	ServiceReclaimStatusCode  *string                            `json:"serviceReclaimStatusCode,omitempty"`
}

type Billing_Item_Cancellation_Request_Status struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	KeyName     *string `json:"keyName,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Billing_Item_Ctc_Account struct {
	Billing_Item
}

type Billing_Item_Gateway_Appliance_Cluster struct {
	Billing_Item

	Resource *Resource_Group `json:"resource,omitempty"`
}

type Billing_Item_Hardware struct {
	Billing_Item

	BillingCycleBandwidthUsage             []Network_Bandwidth_Usage     `json:"billingCycleBandwidthUsage,omitempty"`
	BillingCycleBandwidthUsageCount        *uint                         `json:"billingCycleBandwidthUsageCount,omitempty"`
	BillingCyclePrivateBandwidthUsage      []Network_Bandwidth_Usage     `json:"billingCyclePrivateBandwidthUsage,omitempty"`
	BillingCyclePrivateBandwidthUsageCount *uint                         `json:"billingCyclePrivateBandwidthUsageCount,omitempty"`
	BillingCyclePrivateUsageIn             *float64                      `json:"billingCyclePrivateUsageIn,omitempty"`
	BillingCyclePrivateUsageOut            *float64                      `json:"billingCyclePrivateUsageOut,omitempty"`
	BillingCyclePrivateUsageTotal          *uint                         `json:"billingCyclePrivateUsageTotal,omitempty"`
	BillingCyclePublicBandwidthUsage       []Network_Bandwidth_Usage     `json:"billingCyclePublicBandwidthUsage,omitempty"`
	BillingCyclePublicBandwidthUsageCount  *uint                         `json:"billingCyclePublicBandwidthUsageCount,omitempty"`
	BillingCyclePublicUsageIn              *float64                      `json:"billingCyclePublicUsageIn,omitempty"`
	BillingCyclePublicUsageOut             *float64                      `json:"billingCyclePublicUsageOut,omitempty"`
	BillingCyclePublicUsageTotal           *uint                         `json:"billingCyclePublicUsageTotal,omitempty"`
	LockboxNetworkStorage                  *Billing_Item_Network_Storage `json:"lockboxNetworkStorage,omitempty"`
	MonitoringBillingItemCount             *uint                         `json:"monitoringBillingItemCount,omitempty"`
	MonitoringBillingItems                 []Billing_Item                `json:"monitoringBillingItems,omitempty"`
	Resource                               *Hardware_Server              `json:"resource,omitempty"`
	ResourceTableId                        *int                          `json:"resourceTableId,omitempty"`
}

type Billing_Item_Hardware_Colocation struct {
	Billing_Item_Hardware
}

type Billing_Item_Hardware_Component struct {
	Billing_Item

	Resource        []Hardware_Component `json:"resource,omitempty"`
	ResourceCount   *uint                `json:"resourceCount,omitempty"`
	ResourceTableId *int                 `json:"resourceTableId,omitempty"`
}

type Billing_Item_Hardware_Security_Module struct {
	Billing_Item_Hardware
}

type Billing_Item_Hardware_Server struct {
	Billing_Item_Hardware
}

type Billing_Item_Link_ThePlanet struct {
	Entity

	BillingItem     *Billing_Item     `json:"billingItem,omitempty"`
	ServiceProvider *Service_Provider `json:"serviceProvider,omitempty"`
}

type Billing_Item_Network_Application_Delivery_Controller struct {
	Billing_Item

	BandwidthAllotmentDetail *Network_Bandwidth_Version1_Allotment_Detail `json:"bandwidthAllotmentDetail,omitempty"`
	Resource                 *Network_Application_Delivery_Controller     `json:"resource,omitempty"`
}

type Billing_Item_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress struct {
	Billing_Item

	Resource *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"resource,omitempty"`
}

type Billing_Item_Network_Bandwidth struct {
	Billing_Item
}

type Billing_Item_Network_Firewall struct {
	Billing_Item

	Resource *Network_Component_Firewall `json:"resource,omitempty"`
}

type Billing_Item_Network_Firewall_Module_Context struct {
	Billing_Item
}

type Billing_Item_Network_Interconnect struct {
	Billing_Item
}

type Billing_Item_Network_LoadBalancer struct {
	Billing_Item
}

type Billing_Item_Network_LoadBalancer_Global struct {
	Billing_Item

	Resource *Network_LoadBalancer_Global_Account `json:"resource,omitempty"`
}

type Billing_Item_Network_LoadBalancer_VirtualIpAddress struct {
	Billing_Item

	Resource *Network_LoadBalancer_VirtualIpAddress `json:"resource,omitempty"`
}

type Billing_Item_Network_Message_Delivery struct {
	Billing_Item

	Resource *Network_Message_Delivery `json:"resource,omitempty"`
}

type Billing_Item_Network_Message_Queue struct {
	Billing_Item

	Resource *Network_Message_Queue `json:"resource,omitempty"`
}

type Billing_Item_Network_Message_Queue_Delivery struct {
	Billing_Item_Network_Message_Queue
}

type Billing_Item_Network_PerformanceStorage_Iscsi struct {
	Billing_Item_Network_Storage
}

type Billing_Item_Network_PerformanceStorage_Nfs struct {
	Billing_Item_Network_Storage
}

type Billing_Item_Network_Storage struct {
	Billing_Item

	Resource *Network_Storage `json:"resource,omitempty"`
}

type Billing_Item_Network_Storage_Hub struct {
	Billing_Item_Network_Storage
}

type Billing_Item_Network_Storage_Hub_Bandwidth struct {
	Billing_Item_Network_Storage
}

type Billing_Item_Network_Subnet struct {
	Billing_Item

	Resource        *Network_Subnet `json:"resource,omitempty"`
	ResourceName    *string         `json:"resourceName,omitempty"`
	ResourceTableId *int            `json:"resourceTableId,omitempty"`
}

type Billing_Item_Network_Subnet_IpAddress_Global struct {
	Billing_Item_Network_Subnet
}

type Billing_Item_Network_Tunnel struct {
	Billing_Item

	Resource *Network_Tunnel_Module_Context `json:"resource,omitempty"`
}

type Billing_Item_Network_Vlan struct {
	Billing_Item

	Resource *Network_Vlan `json:"resource,omitempty"`
}

type Billing_Item_Software_Component struct {
	Billing_Item

	Resource        *Software_Component `json:"resource,omitempty"`
	ResourceTableId *int                `json:"resourceTableId,omitempty"`
}

type Billing_Item_Software_Component_Analytics_Urchin struct {
	Billing_Item
}

type Billing_Item_Software_Component_ControlPanel struct {
	Billing_Item
}

type Billing_Item_Software_Component_ControlPanel_Parallels_Plesk_Billing struct {
	Billing_Item
}

type Billing_Item_Software_Component_OperatingSystem_Addon struct {
	Billing_Item
}

type Billing_Item_Software_Component_OperatingSystem_Addon_Citrix_Essentials struct {
	Billing_Item_Software_Component_OperatingSystem_Addon

	Resource *Software_Component `json:"resource,omitempty"`
}

type Billing_Item_Software_Component_Virtual_OperatingSystem struct {
	Billing_Item
}

type Billing_Item_Software_Component_Virtual_OperatingSystem_Microsoft struct {
	Billing_Item_Software_Component_Virtual_OperatingSystem

	Resource        *Software_VirtualLicense `json:"resource,omitempty"`
	ResourceTableId *int                     `json:"resourceTableId,omitempty"`
}

type Billing_Item_Software_Component_Virtual_OperatingSystem_Redhat struct {
	Billing_Item_Software_Component_Virtual_OperatingSystem

	Resource        *Software_Component `json:"resource,omitempty"`
	ResourceTableId *int                `json:"resourceTableId,omitempty"`
}

type Billing_Item_Software_License struct {
	Billing_Item

	Resource *Software_AccountLicense `json:"resource,omitempty"`
}

type Billing_Item_Support struct {
	Billing_Item
}

type Billing_Item_User_Customer_External_Binding struct {
	Billing_Item

	Resource *User_Customer_External_Binding `json:"resource,omitempty"`
}

type Billing_Item_Virtual_Dedicated_Rack struct {
	Billing_Item

	BillingCycleBandwidthUsage             []Network_Bandwidth_Usage             `json:"billingCycleBandwidthUsage,omitempty"`
	BillingCycleBandwidthUsageCount        *uint                                 `json:"billingCycleBandwidthUsageCount,omitempty"`
	BillingCyclePrivateBandwidthUsage      []Network_Bandwidth_Usage             `json:"billingCyclePrivateBandwidthUsage,omitempty"`
	BillingCyclePrivateBandwidthUsageCount *uint                                 `json:"billingCyclePrivateBandwidthUsageCount,omitempty"`
	BillingCyclePrivateUsageIn             *float64                              `json:"billingCyclePrivateUsageIn,omitempty"`
	BillingCyclePrivateUsageOut            *float64                              `json:"billingCyclePrivateUsageOut,omitempty"`
	BillingCyclePrivateUsageTotal          *uint                                 `json:"billingCyclePrivateUsageTotal,omitempty"`
	BillingCyclePublicBandwidthUsage       []Network_Bandwidth_Usage             `json:"billingCyclePublicBandwidthUsage,omitempty"`
	BillingCyclePublicBandwidthUsageCount  *uint                                 `json:"billingCyclePublicBandwidthUsageCount,omitempty"`
	BillingCyclePublicUsageIn              *float64                              `json:"billingCyclePublicUsageIn,omitempty"`
	BillingCyclePublicUsageOut             *float64                              `json:"billingCyclePublicUsageOut,omitempty"`
	BillingCyclePublicUsageTotal           *uint                                 `json:"billingCyclePublicUsageTotal,omitempty"`
	Resource                               *Network_Bandwidth_Version1_Allotment `json:"resource,omitempty"`
}

type Billing_Item_Virtual_Disk_Image struct {
	Billing_Item

	Resource        *Virtual_Disk_Image `json:"resource,omitempty"`
	ResourceTableId *int                `json:"resourceTableId,omitempty"`
}

type Billing_Item_Virtual_Guest struct {
	Billing_Item

	BillingCycleBandwidthUsage             []Network_Bandwidth_Usage `json:"billingCycleBandwidthUsage,omitempty"`
	BillingCycleBandwidthUsageCount        *uint                     `json:"billingCycleBandwidthUsageCount,omitempty"`
	BillingCyclePrivateBandwidthUsage      []Network_Bandwidth_Usage `json:"billingCyclePrivateBandwidthUsage,omitempty"`
	BillingCyclePrivateBandwidthUsageCount *uint                     `json:"billingCyclePrivateBandwidthUsageCount,omitempty"`
	BillingCyclePrivateUsageIn             *float64                  `json:"billingCyclePrivateUsageIn,omitempty"`
	BillingCyclePrivateUsageOut            *float64                  `json:"billingCyclePrivateUsageOut,omitempty"`
	BillingCyclePrivateUsageTotal          *uint                     `json:"billingCyclePrivateUsageTotal,omitempty"`
	BillingCyclePublicBandwidthUsage       []Network_Bandwidth_Usage `json:"billingCyclePublicBandwidthUsage,omitempty"`
	BillingCyclePublicBandwidthUsageCount  *uint                     `json:"billingCyclePublicBandwidthUsageCount,omitempty"`
	BillingCyclePublicUsageIn              *float64                  `json:"billingCyclePublicUsageIn,omitempty"`
	BillingCyclePublicUsageOut             *float64                  `json:"billingCyclePublicUsageOut,omitempty"`
	BillingCyclePublicUsageTotal           *uint                     `json:"billingCyclePublicUsageTotal,omitempty"`
	MonitoringBillingItemCount             *uint                     `json:"monitoringBillingItemCount,omitempty"`
	MonitoringBillingItems                 []Billing_Item            `json:"monitoringBillingItems,omitempty"`
	Resource                               *Virtual_Guest            `json:"resource,omitempty"`
	ResourceTableId                        *int                      `json:"resourceTableId,omitempty"`
}

type Billing_Item_Virtual_Host_Usage struct {
	Billing_Item

	Resource        *Hardware `json:"resource,omitempty"`
	ResourceTableId *int      `json:"resourceTableId,omitempty"`
}

type Billing_Item_Workspace struct {
	Billing_Item
}

type Billing_Order struct {
	Entity

	Account                      *Account                             `json:"account,omitempty"`
	AccountId                    *int                                 `json:"accountId,omitempty"`
	Brand                        *Brand                               `json:"brand,omitempty"`
	Cart                         *Billing_Order_Cart                  `json:"cart,omitempty"`
	CoreRestrictedItemCount      *uint                                `json:"coreRestrictedItemCount,omitempty"`
	CoreRestrictedItems          []Billing_Order_Item                 `json:"coreRestrictedItems,omitempty"`
	CreateDate                   *time.Time                           `json:"createDate,omitempty"`
	CreditCardTransactionCount   *uint                                `json:"creditCardTransactionCount,omitempty"`
	CreditCardTransactions       []Billing_Payment_Card_Transaction   `json:"creditCardTransactions,omitempty"`
	ExchangeRate                 *Billing_Currency_ExchangeRate       `json:"exchangeRate,omitempty"`
	Id                           *int                                 `json:"id,omitempty"`
	ImpersonatingUserRecordId    *int                                 `json:"impersonatingUserRecordId,omitempty"`
	InitialInvoice               *Billing_Invoice                     `json:"initialInvoice,omitempty"`
	ItemCount                    *uint                                `json:"itemCount,omitempty"`
	Items                        []Billing_Order_Item                 `json:"items,omitempty"`
	ModifyDate                   *time.Time                           `json:"modifyDate,omitempty"`
	OrderApprovalDate            *time.Time                           `json:"orderApprovalDate,omitempty"`
	OrderNonServerMonthlyAmount  *float64                             `json:"orderNonServerMonthlyAmount,omitempty"`
	OrderQuoteId                 *int                                 `json:"orderQuoteId,omitempty"`
	OrderServerMonthlyAmount     *float64                             `json:"orderServerMonthlyAmount,omitempty"`
	OrderTopLevelItemCount       *uint                                `json:"orderTopLevelItemCount,omitempty"`
	OrderTopLevelItems           []Billing_Order_Item                 `json:"orderTopLevelItems,omitempty"`
	OrderTotalAmount             *float64                             `json:"orderTotalAmount,omitempty"`
	OrderTotalOneTime            *float64                             `json:"orderTotalOneTime,omitempty"`
	OrderTotalOneTimeAmount      *float64                             `json:"orderTotalOneTimeAmount,omitempty"`
	OrderTotalOneTimeTaxAmount   *float64                             `json:"orderTotalOneTimeTaxAmount,omitempty"`
	OrderTotalRecurring          *float64                             `json:"orderTotalRecurring,omitempty"`
	OrderTotalRecurringAmount    *float64                             `json:"orderTotalRecurringAmount,omitempty"`
	OrderTotalRecurringTaxAmount *float64                             `json:"orderTotalRecurringTaxAmount,omitempty"`
	OrderTotalSetupAmount        *float64                             `json:"orderTotalSetupAmount,omitempty"`
	OrderType                    *Billing_Order_Type                  `json:"orderType,omitempty"`
	OrderTypeId                  *int                                 `json:"orderTypeId,omitempty"`
	PaypalTransactionCount       *uint                                `json:"paypalTransactionCount,omitempty"`
	PaypalTransactions           []Billing_Payment_PayPal_Transaction `json:"paypalTransactions,omitempty"`
	PresaleEvent                 *Sales_Presale_Event                 `json:"presaleEvent,omitempty"`
	PresaleEventId               *int                                 `json:"presaleEventId,omitempty"`
	PrivateCloudOrderFlag        *bool                                `json:"privateCloudOrderFlag,omitempty"`
	Quote                        *Billing_Order_Quote                 `json:"quote,omitempty"`
	ReferralPartner              *Account                             `json:"referralPartner,omitempty"`
	Status                       *string                              `json:"status,omitempty"`
	UpgradeRequestFlag           *bool                                `json:"upgradeRequestFlag,omitempty"`
	UserRecord                   *User_Customer                       `json:"userRecord,omitempty"`
	UserRecordId                 *int                                 `json:"userRecordId,omitempty"`
}

type Billing_Order_Cart struct {
	Billing_Order_Quote
}

type Billing_Order_Item struct {
	Entity

	BillingItem               *Billing_Item                        `json:"billingItem,omitempty"`
	BundledItemCount          *uint                                `json:"bundledItemCount,omitempty"`
	BundledItems              []Billing_Order_Item                 `json:"bundledItems,omitempty"`
	Category                  *Product_Item_Category               `json:"category,omitempty"`
	CategoryCode              *string                              `json:"categoryCode,omitempty"`
	Children                  []Billing_Order_Item                 `json:"children,omitempty"`
	ChildrenCount             *uint                                `json:"childrenCount,omitempty"`
	Description               *string                              `json:"description,omitempty"`
	DomainName                *string                              `json:"domainName,omitempty"`
	GlobalIdentifier          *string                              `json:"globalIdentifier,omitempty"`
	HardwareGenericComponent  *Hardware_Component_Model_Generic    `json:"hardwareGenericComponent,omitempty"`
	HostName                  *string                              `json:"hostName,omitempty"`
	HourlyRecurringFee        *float64                             `json:"hourlyRecurringFee,omitempty"`
	Id                        *int                                 `json:"id,omitempty"`
	Item                      *Product_Item                        `json:"item,omitempty"`
	ItemCategoryAnswerCount   *uint                                `json:"itemCategoryAnswerCount,omitempty"`
	ItemCategoryAnswers       []Billing_Order_Item_Category_Answer `json:"itemCategoryAnswers,omitempty"`
	ItemId                    *int                                 `json:"itemId,omitempty"`
	ItemPrice                 *Product_Item_Price                  `json:"itemPrice,omitempty"`
	ItemPriceId               *float64                             `json:"itemPriceId,omitempty"`
	LaborAfterTaxAmount       *float64                             `json:"laborAfterTaxAmount,omitempty"`
	LaborFee                  *float64                             `json:"laborFee,omitempty"`
	LaborFeeTaxRate           *float64                             `json:"laborFeeTaxRate,omitempty"`
	LaborTaxAmount            *float64                             `json:"laborTaxAmount,omitempty"`
	Location                  *Location                            `json:"location,omitempty"`
	NextOrderChildren         []Billing_Order_Item                 `json:"nextOrderChildren,omitempty"`
	NextOrderChildrenCount    *uint                                `json:"nextOrderChildrenCount,omitempty"`
	OldBillingItem            *Billing_Item                        `json:"oldBillingItem,omitempty"`
	OneTimeAfterTaxAmount     *float64                             `json:"oneTimeAfterTaxAmount,omitempty"`
	OneTimeFee                *float64                             `json:"oneTimeFee,omitempty"`
	OneTimeFeeTaxRate         *float64                             `json:"oneTimeFeeTaxRate,omitempty"`
	OneTimeTaxAmount          *float64                             `json:"oneTimeTaxAmount,omitempty"`
	Order                     *Billing_Order                       `json:"order,omitempty"`
	OrderApprovalDate         *time.Time                           `json:"orderApprovalDate,omitempty"`
	Package                   *Product_Package                     `json:"package,omitempty"`
	Parent                    *Billing_Order_Item                  `json:"parent,omitempty"`
	ParentId                  *int                                 `json:"parentId,omitempty"`
	PromoCodeId               *int                                 `json:"promoCodeId,omitempty"`
	Quantity                  *int                                 `json:"quantity,omitempty"`
	RecurringAfterTaxAmount   *float64                             `json:"recurringAfterTaxAmount,omitempty"`
	RecurringFee              *float64                             `json:"recurringFee,omitempty"`
	RecurringTaxAmount        *float64                             `json:"recurringTaxAmount,omitempty"`
	RedundantPowerSupplyCount *uint                                `json:"redundantPowerSupplyCount,omitempty"`
	SetupAfterTaxAmount       *float64                             `json:"setupAfterTaxAmount,omitempty"`
	SetupFee                  *float64                             `json:"setupFee,omitempty"`
	SetupFeeDeferralMonths    *int                                 `json:"setupFeeDeferralMonths,omitempty"`
	SetupFeeTaxRate           *float64                             `json:"setupFeeTaxRate,omitempty"`
	SetupTaxAmount            *float64                             `json:"setupTaxAmount,omitempty"`
	SoftwareDescription       *Software_Description                `json:"softwareDescription,omitempty"`
	StorageGroupCount         *uint                                `json:"storageGroupCount,omitempty"`
	StorageGroups             []Configuration_Storage_Group_Order  `json:"storageGroups,omitempty"`
	TotalRecurringAmount      *float64                             `json:"totalRecurringAmount,omitempty"`
	UpgradeItem               *Product_Item                        `json:"upgradeItem,omitempty"`
}

type Billing_Order_Item_Category_Answer struct {
	Entity

	Answer     *string                         `json:"answer,omitempty"`
	CreateDate *time.Time                      `json:"createDate,omitempty"`
	OrderItem  *Billing_Order_Item             `json:"orderItem,omitempty"`
	Question   *Product_Item_Category_Question `json:"question,omitempty"`
	QuestionId *int                            `json:"questionId,omitempty"`
}

type Billing_Order_Note struct {
	Entity

	CreateDate *time.Time     `json:"createDate,omitempty"`
	Employee   *User_Employee `json:"employee,omitempty"`
	Order      *Billing_Order `json:"order,omitempty"`
}

type Billing_Order_Quote struct {
	Entity

	Account                 *Account        `json:"account,omitempty"`
	AccountId               *int            `json:"accountId,omitempty"`
	CompletedPurchaseDataId *int            `json:"completedPurchaseDataId,omitempty"`
	CreateDate              *time.Time      `json:"createDate,omitempty"`
	ExpirationDate          *time.Time      `json:"expirationDate,omitempty"`
	Id                      *int            `json:"id,omitempty"`
	ModifyDate              *time.Time      `json:"modifyDate,omitempty"`
	Name                    *string         `json:"name,omitempty"`
	Order                   *Billing_Order  `json:"order,omitempty"`
	OrdersFromQuote         []Billing_Order `json:"ordersFromQuote,omitempty"`
	OrdersFromQuoteCount    *uint           `json:"ordersFromQuoteCount,omitempty"`
	PublicNote              *string         `json:"publicNote,omitempty"`
	QuoteKey                *string         `json:"quoteKey,omitempty"`
	Status                  *string         `json:"status,omitempty"`
}

type Billing_Order_Type struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	Type        *string `json:"type,omitempty"`
}

type Billing_Payment_Card_ChangeRequest struct {
	Entity

	Account                         *Account                          `json:"account,omitempty"`
	AccountId                       *int                              `json:"accountId,omitempty"`
	Amount                          *float64                          `json:"amount,omitempty"`
	AuthorizedCreditCardTransaction *Billing_Payment_Card_Transaction `json:"authorizedCreditCardTransaction,omitempty"`
	BillingAddressLine1             *string                           `json:"billingAddressLine1,omitempty"`
	BillingAddressLine2             *string                           `json:"billingAddressLine2,omitempty"`
	BillingCity                     *string                           `json:"billingCity,omitempty"`
	BillingCountryCode              *string                           `json:"billingCountryCode,omitempty"`
	BillingEmail                    *string                           `json:"billingEmail,omitempty"`
	BillingNameCompany              *string                           `json:"billingNameCompany,omitempty"`
	BillingNameFirst                *string                           `json:"billingNameFirst,omitempty"`
	BillingNameLast                 *string                           `json:"billingNameLast,omitempty"`
	BillingPhoneFax                 *string                           `json:"billingPhoneFax,omitempty"`
	BillingPhoneVoice               *string                           `json:"billingPhoneVoice,omitempty"`
	BillingPostalCode               *string                           `json:"billingPostalCode,omitempty"`
	BillingState                    *string                           `json:"billingState,omitempty"`
	CaptureCreditCardTransaction    *Billing_Payment_Card_Transaction `json:"captureCreditCardTransaction,omitempty"`
	CardAccountLast4                *string                           `json:"cardAccountLast4,omitempty"`
	CardAccountNumber               *string                           `json:"cardAccountNumber,omitempty"`
	CardExpirationMonth             *string                           `json:"cardExpirationMonth,omitempty"`
	CardExpirationYear              *string                           `json:"cardExpirationYear,omitempty"`
	CardNickname                    *string                           `json:"cardNickname,omitempty"`
	CardType                        *string                           `json:"cardType,omitempty"`
	CreditCardVerificationNumber    *string                           `json:"creditCardVerificationNumber,omitempty"`
	CurrencyShortName               *string                           `json:"currencyShortName,omitempty"`
	DeviceFingerprintId             *string                           `json:"deviceFingerprintId,omitempty"`
	Id                              *int                              `json:"id,omitempty"`
	Notes                           *string                           `json:"notes,omitempty"`
	PaymentRoleId                   *int                              `json:"paymentRoleId,omitempty"`
	PaymentType                     *string                           `json:"paymentType,omitempty"`
	TicketAttachmentReferenceCount  *uint                             `json:"ticketAttachmentReferenceCount,omitempty"`
	TicketAttachmentReferences      []Ticket_Attachment               `json:"ticketAttachmentReferences,omitempty"`
	TicketId                        *int                              `json:"ticketId,omitempty"`
}

type Billing_Payment_Card_ManualPayment struct {
	Entity

	Account                           *Account                            `json:"account,omitempty"`
	AccountId                         *int                                `json:"accountId,omitempty"`
	Amount                            *float64                            `json:"amount,omitempty"`
	AuthorizedCreditCardTransaction   *Billing_Payment_Card_Transaction   `json:"authorizedCreditCardTransaction,omitempty"`
	AuthorizedCreditCardTransactionId *int                                `json:"authorizedCreditCardTransactionId,omitempty"`
	AuthorizedPayPalTransaction       *Billing_Payment_PayPal_Transaction `json:"authorizedPayPalTransaction,omitempty"`
	AuthorizedPayPalTransactionId     *int                                `json:"authorizedPayPalTransactionId,omitempty"`
	BillingAddressLine1               *string                             `json:"billingAddressLine1,omitempty"`
	BillingAddressLine2               *string                             `json:"billingAddressLine2,omitempty"`
	BillingCity                       *string                             `json:"billingCity,omitempty"`
	BillingCountryCode                *string                             `json:"billingCountryCode,omitempty"`
	BillingEmail                      *string                             `json:"billingEmail,omitempty"`
	BillingNameCompany                *string                             `json:"billingNameCompany,omitempty"`
	BillingNameFirst                  *string                             `json:"billingNameFirst,omitempty"`
	BillingNameLast                   *string                             `json:"billingNameLast,omitempty"`
	BillingPhoneFax                   *string                             `json:"billingPhoneFax,omitempty"`
	BillingPhoneVoice                 *string                             `json:"billingPhoneVoice,omitempty"`
	BillingPostalCode                 *string                             `json:"billingPostalCode,omitempty"`
	BillingState                      *string                             `json:"billingState,omitempty"`
	CancelUrl                         *string                             `json:"cancelUrl,omitempty"`
	CaptureCreditCardTransaction      *Billing_Payment_Card_Transaction   `json:"captureCreditCardTransaction,omitempty"`
	CapturePayPalTransaction          *Billing_Payment_PayPal_Transaction `json:"capturePayPalTransaction,omitempty"`
	CardAccountHash                   *string                             `json:"cardAccountHash,omitempty"`
	CardAccountLast4                  *string                             `json:"cardAccountLast4,omitempty"`
	CardAccountNumber                 *string                             `json:"cardAccountNumber,omitempty"`
	CardExpirationMonth               *string                             `json:"cardExpirationMonth,omitempty"`
	CardExpirationYear                *string                             `json:"cardExpirationYear,omitempty"`
	CardType                          *string                             `json:"cardType,omitempty"`
	CreditCardVerificationNumber      *string                             `json:"creditCardVerificationNumber,omitempty"`
	CurrencyShortName                 *string                             `json:"currencyShortName,omitempty"`
	DeviceFingerprintId               *string                             `json:"deviceFingerprintId,omitempty"`
	FromIpAddress                     *string                             `json:"fromIpAddress,omitempty"`
	Id                                *int                                `json:"id,omitempty"`
	Notes                             *string                             `json:"notes,omitempty"`
	PaymentType                       *string                             `json:"paymentType,omitempty"`
	ReturnUrl                         *string                             `json:"returnUrl,omitempty"`
	TicketAttachmentReferenceCount    *uint                               `json:"ticketAttachmentReferenceCount,omitempty"`
	TicketAttachmentReferences        []Ticket_Attachment                 `json:"ticketAttachmentReferences,omitempty"`
	Type                              *string                             `json:"type,omitempty"`
}

type Billing_Payment_Card_Transaction struct {
	Billing_Payment_Transaction

	Account             *Account       `json:"account,omitempty"`
	AccountId           *int           `json:"accountId,omitempty"`
	Amount              *float64       `json:"amount,omitempty"`
	BillingAddressLine1 *string        `json:"billingAddressLine1,omitempty"`
	BillingAddressLine2 *string        `json:"billingAddressLine2,omitempty"`
	BillingCity         *string        `json:"billingCity,omitempty"`
	BillingCountryCode  *string        `json:"billingCountryCode,omitempty"`
	BillingEmail        *string        `json:"billingEmail,omitempty"`
	BillingNameCompany  *string        `json:"billingNameCompany,omitempty"`
	BillingNameFirst    *string        `json:"billingNameFirst,omitempty"`
	BillingNameLast     *string        `json:"billingNameLast,omitempty"`
	BillingPhoneFax     *string        `json:"billingPhoneFax,omitempty"`
	BillingPhoneVoice   *string        `json:"billingPhoneVoice,omitempty"`
	BillingPostalCode   *string        `json:"billingPostalCode,omitempty"`
	BillingState        *string        `json:"billingState,omitempty"`
	CardAccountLast4    *int           `json:"cardAccountLast4,omitempty"`
	CardExpirationMonth *int           `json:"cardExpirationMonth,omitempty"`
	CardExpirationYear  *int           `json:"cardExpirationYear,omitempty"`
	CardType            *string        `json:"cardType,omitempty"`
	CreateDate          *time.Time     `json:"createDate,omitempty"`
	Id                  *int           `json:"id,omitempty"`
	InvoiceId           *int           `json:"invoiceId,omitempty"`
	ModifyDate          *time.Time     `json:"modifyDate,omitempty"`
	Order               *Billing_Order `json:"order,omitempty"`
	OrderFromIpAddress  *string        `json:"orderFromIpAddress,omitempty"`
	ReferenceCode       *string        `json:"referenceCode,omitempty"`
	RequestId           *string        `json:"requestId,omitempty"`
	ReturnStatus        *int           `json:"returnStatus,omitempty"`
	SerializedReply     *string        `json:"serializedReply,omitempty"`
	SerializedRequest   *string        `json:"serializedRequest,omitempty"`
}

type Billing_Payment_PayPal_Transaction struct {
	Billing_Payment_Transaction

	Account              *Account       `json:"account,omitempty"`
	AccountId            *int           `json:"accountId,omitempty"`
	AddressCityName      *string        `json:"addressCityName,omitempty"`
	AddressCountry       *string        `json:"addressCountry,omitempty"`
	AddressName          *string        `json:"addressName,omitempty"`
	AddressPostalCode    *string        `json:"addressPostalCode,omitempty"`
	AddressStateProvence *string        `json:"addressStateProvence,omitempty"`
	AddressStatus        *string        `json:"addressStatus,omitempty"`
	AddressStreet1       *string        `json:"addressStreet1,omitempty"`
	AddressStreet2       *string        `json:"addressStreet2,omitempty"`
	ContactPhone         *string        `json:"contactPhone,omitempty"`
	CreateDate           *time.Time     `json:"createDate,omitempty"`
	ExchangeRate         *string        `json:"exchangeRate,omitempty"`
	FeeAmount            *float64       `json:"feeAmount,omitempty"`
	GrossAmount          *float64       `json:"grossAmount,omitempty"`
	Id                   *int           `json:"id,omitempty"`
	InvoiceId            *int           `json:"invoiceId,omitempty"`
	LastPaypalCommand    *string        `json:"lastPaypalCommand,omitempty"`
	ModifyDate           *time.Time     `json:"modifyDate,omitempty"`
	Order                *Billing_Order `json:"order,omitempty"`
	OrderFromIpAddress   *string        `json:"orderFromIpAddress,omitempty"`
	OrderTotal           *float64       `json:"orderTotal,omitempty"`
	Payer                *string        `json:"payer,omitempty"`
	PayerBusiness        *string        `json:"payerBusiness,omitempty"`
	PayerCountry         *string        `json:"payerCountry,omitempty"`
	PayerFirstName       *string        `json:"payerFirstName,omitempty"`
	PayerId              *string        `json:"payerId,omitempty"`
	PayerLastName        *string        `json:"payerLastName,omitempty"`
	PayerStatus          *string        `json:"payerStatus,omitempty"`
	PaymentDate          *time.Time     `json:"paymentDate,omitempty"`
	PaymentStatus        *string        `json:"paymentStatus,omitempty"`
	PaymentType          *string        `json:"paymentType,omitempty"`
	PendingReason        *string        `json:"pendingReason,omitempty"`
	SerializedReply      *string        `json:"serializedReply,omitempty"`
	SerializedRequest    *string        `json:"serializedRequest,omitempty"`
	SettleAmount         *float64       `json:"settleAmount,omitempty"`
	TaxAmount            *float64       `json:"taxAmount,omitempty"`
	Token                *string        `json:"token,omitempty"`
	TransactionId        *string        `json:"transactionId,omitempty"`
	TransactionType      *string        `json:"transactionType,omitempty"`
}

type Billing_Payment_Processor struct {
	Entity

	BrandAssignmentCount *uint                              `json:"brandAssignmentCount,omitempty"`
	BrandAssignments     []Brand_Payment_Processor          `json:"brandAssignments,omitempty"`
	Description          *string                            `json:"description,omitempty"`
	Name                 *string                            `json:"name,omitempty"`
	OwnerAccount         *Account                           `json:"ownerAccount,omitempty"`
	PaymentMethodCount   *uint                              `json:"paymentMethodCount,omitempty"`
	PaymentMethods       []Billing_Payment_Processor_Method `json:"paymentMethods,omitempty"`
	Type                 *Billing_Payment_Processor_Type    `json:"type,omitempty"`
}

type Billing_Payment_Processor_Method struct {
	Entity

	MethodKey            *string                    `json:"methodKey,omitempty"`
	MultipleCurrencyFlag *bool                      `json:"multipleCurrencyFlag,omitempty"`
	PaymentProcessor     *Billing_Payment_Processor `json:"paymentProcessor,omitempty"`
	PaymentType          *Billing_Payment_Type      `json:"paymentType,omitempty"`
}

type Billing_Payment_Processor_Type struct {
	Entity

	Description           *string                     `json:"description,omitempty"`
	KeyName               *string                     `json:"keyName,omitempty"`
	Name                  *string                     `json:"name,omitempty"`
	PaymentProcessorCount *uint                       `json:"paymentProcessorCount,omitempty"`
	PaymentProcessors     []Billing_Payment_Processor `json:"paymentProcessors,omitempty"`
}

type Billing_Payment_Transaction struct {
	Entity
}

type Billing_Payment_Type struct {
	Entity

	Description *string `json:"description,omitempty"`
	KeyName     *string `json:"keyName,omitempty"`
	Name        *string `json:"name,omitempty"`
}
