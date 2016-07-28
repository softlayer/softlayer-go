package softlayer

import (
	"time"

	"github.ibm.com/riethm/gopherlayer/datatypes"
)

type Billing_Currency interface {
	Entity

	GetAllObjects() (datatypes.Billing_Currency, error)
	GetObject() (datatypes.Billing_Currency, error)
	GetPrice(price float64, formatOptions datatypes.Container_Billing_Currency_Format) (string, error)
}

type Billing_Currency_ExchangeRate interface {
	Entity

	GetAllCurrencyExchangeRates(stringDate string) (datatypes.Billing_Currency_ExchangeRate, error)
	GetCurrencies() (datatypes.Billing_Currency, error)
	GetExchangeRate(to string, from string, effectiveDate time.Time) (datatypes.Billing_Currency_ExchangeRate, error)
	GetObject() (datatypes.Billing_Currency_ExchangeRate, error)
	GetPrice(price float64, formatOptions datatypes.Container_Billing_Currency_Format) (string, error)

	GetFundingCurrency() (datatypes.Billing_Currency, error)
	GetLocalCurrency() (datatypes.Billing_Currency, error)
}

type Billing_Info interface {
	Entity

	GetObject() (datatypes.Billing_Info, error)

	GetAccount() (datatypes.Account, error)
	GetAchInformation() (datatypes.Billing_Info_Ach, error)
	GetCurrency() (datatypes.Billing_Currency, error)
	GetCurrentBillingCycle() (datatypes.Billing_Info_Cycle, error)
	GetLastBillDate() (time.Time, error)
	GetNextBillDate() (time.Time, error)
}

type Billing_Info_Ach interface {
	Entity

	GetAccount() (datatypes.Account, error)
}

type Billing_Info_Cycle interface {
	Entity

	GetAccount() (datatypes.Account, error)
}

type Billing_Invoice interface {
	Entity

	EmailInvoices(options datatypes.Container_Billing_Invoice_Email) error
	GetExcel() ([]byte, error)
	GetObject() (datatypes.Billing_Invoice, error)
	GetPdf() ([]byte, error)
	GetPdfDetailed() ([]byte, error)
	GetPdfDetailedFilename() (string, error)
	GetPdfFileSize() (int, error)
	GetPdfFilename() (string, error)
	GetPreliminaryExcel() ([]byte, error)
	GetPreliminaryPdf() ([]byte, error)
	GetPreliminaryPdfDetailed() ([]byte, error)
	GetXlsFilename() (string, error)
	GetZeroFeeItemCounts() (datatypes.Container_Product_Item_Category_ZeroFee_Count, error)

	GetAccount() (datatypes.Account, error)
	GetAmount() (float64, error)
	GetBrandAtInvoiceCreation() (datatypes.Brand, error)
	GetDetailedPdfGeneratedFlag() (bool, error)
	GetInvoiceTopLevelItems() (datatypes.Billing_Invoice_Item, error)
	GetInvoiceTotalAmount() (float64, error)
	GetInvoiceTotalOneTimeAmount() (float64, error)
	GetInvoiceTotalOneTimeTaxAmount() (float64, error)
	GetInvoiceTotalPreTaxAmount() (float64, error)
	GetInvoiceTotalRecurringAmount() (float64, error)
	GetInvoiceTotalRecurringTaxAmount() (float64, error)
	GetItems() (datatypes.Billing_Invoice_Item, error)
	GetPayment() (float64, error)
	GetPayments() (datatypes.Billing_Invoice_Receivable_Payment, error)
	GetSellerRegistration() (string, error)
	GetTaxInfo() (datatypes.Billing_Invoice_Tax_Info, error)
	GetTaxInfoHistory() (datatypes.Billing_Invoice_Tax_Info, error)
	GetTaxMessage() (string, error)
	GetTaxType() (datatypes.Billing_Invoice_Tax_Type, error)
}

type Billing_Invoice_Item interface {
	Entity

	GetObject() (datatypes.Billing_Invoice_Item, error)

	GetAssociatedChildren() (datatypes.Billing_Invoice_Item, error)
	GetAssociatedInvoiceItem() (datatypes.Billing_Invoice_Item, error)
	GetBillingItem() (datatypes.Billing_Item, error)
	GetCategory() (datatypes.Product_Item_Category, error)
	GetChildren() (datatypes.Billing_Invoice_Item, error)
	GetFilteredAssociatedChildren() (datatypes.Billing_Invoice_Item, error)
	GetInvoice() (datatypes.Billing_Invoice, error)
	GetLocation() (datatypes.Location, error)
	GetNonZeroAssociatedChildren() (datatypes.Billing_Invoice_Item, error)
	GetParent() (datatypes.Billing_Invoice_Item, error)
	GetProduct() (datatypes.Product_Item, error)
	GetTotalOneTimeAmount() (float64, error)
	GetTotalOneTimeTaxAmount() (float64, error)
	GetTotalRecurringAmount() (float64, error)
	GetTotalRecurringTaxAmount() (float64, error)
}

type Billing_Invoice_Item_Hardware interface {
	Billing_Invoice_Item

	GetResource() (datatypes.Hardware, error)
}

type Billing_Invoice_Item_Tax_Info interface {
	Entity

	GetInvoiceItem() (datatypes.Billing_Invoice_Item, error)
	GetInvoiceTaxInfo() (datatypes.Billing_Invoice_Tax_Info, error)
	GetToCurrency() (datatypes.Billing_Currency, error)
}

type Billing_Invoice_Next interface {
	Entity

	GetExcel(documentCreateDate time.Time) ([]byte, error)
	GetPdf(documentCreateDate time.Time) ([]byte, error)
	GetPdfDetailed(documentCreateDate time.Time) ([]byte, error)
}

type Billing_Invoice_Receivable_Payment interface {
	Entity

	GetAccount() (datatypes.Account, error)
	GetCreditCardLastFourDigits() (int, error)
	GetCreditCardRequestId() (string, error)
	GetCreditCardTransaction() (datatypes.Billing_Payment_Card_Transaction, error)
	GetExchangeRate() (datatypes.Billing_Currency_ExchangeRate, error)
	GetInvoice() (datatypes.Billing_Invoice, error)
	GetPaypalTransaction() (datatypes.Billing_Payment_PayPal_Transaction, error)
}

type Billing_Invoice_Tax_Info interface {
	Entity

	GetCurrency() (datatypes.Billing_Currency, error)
	GetFunctionalCurrency() (datatypes.Billing_Currency, error)
	GetInvoice() (datatypes.Billing_Invoice, error)
	GetItemWithCurrencyInfo() (datatypes.Billing_Invoice_Item_Tax_Info, error)
	GetItems() (datatypes.Billing_Invoice_Item_Tax_Info, error)
	GetTotalTaxAmountToCurrency() (float64, error)
}

type Billing_Invoice_Tax_Status interface {
	Entity

	GetAllObjects() (datatypes.Billing_Invoice_Tax_Status, error)
	GetObject() (datatypes.Billing_Invoice_Tax_Status, error)
}

type Billing_Invoice_Tax_Type interface {
	Entity

	GetAllObjects() (datatypes.Billing_Invoice_Tax_Type, error)
	GetObject() (datatypes.Billing_Invoice_Tax_Type, error)
}

type Billing_Item interface {
	Entity

	CancelItem(cancelImmediately bool, cancelAssociatedBillingItems bool, reason string, customerNote string) (bool, error)
	CancelService() (bool, error)
	CancelServiceOnAnniversaryDate() (bool, error)
	GetObject() (datatypes.Billing_Item, error)
	GetServiceBillingItemsByCategory(categoryCode string, includeZeroRecurringFee bool) (datatypes.Billing_Item, error)
	RemoveAssociationId() (bool, error)
	SetAssociationId(associatedId int) (bool, error)
	VoidCancelService() (bool, error)

	GetAccount() (datatypes.Account, error)
	GetActiveAgreement() (datatypes.Account_Agreement, error)
	GetActiveAgreementFlag() (datatypes.Account_Agreement, error)
	GetActiveAssociatedChildren() (datatypes.Billing_Item, error)
	GetActiveAssociatedGuestDiskBillingItems() (datatypes.Billing_Item, error)
	GetActiveBundledItems() (datatypes.Billing_Item, error)
	GetActiveCancellationItem() (datatypes.Billing_Item_Cancellation_Request_Item, error)
	GetActiveChildren() (datatypes.Billing_Item, error)
	GetActiveFlag() (bool, error)
	GetActiveSparePoolAssociatedGuestDiskBillingItems() (datatypes.Billing_Item, error)
	GetActiveSparePoolBundledItems() (datatypes.Billing_Item, error)
	GetAssociatedBillingItem() (datatypes.Billing_Item, error)
	GetAssociatedBillingItemHistory() (datatypes.Billing_Item_Association_History, error)
	GetAssociatedChildren() (datatypes.Billing_Item, error)
	GetAssociatedParent() (datatypes.Billing_Item, error)
	GetAvailableMatchingVlans() (datatypes.Network_Vlan, error)
	GetBandwidthAllocation() (datatypes.Network_Bandwidth_Version1_Allocation, error)
	GetBillableChildren() (datatypes.Billing_Item, error)
	GetBundleItems() (datatypes.Product_Item_Bundles, error)
	GetBundledItems() (datatypes.Billing_Item, error)
	GetCanceledChildren() (datatypes.Billing_Item, error)
	GetCancellationReason() (datatypes.Billing_Item_Cancellation_Reason, error)
	GetCancellationRequests() (datatypes.Billing_Item_Cancellation_Request, error)
	GetCategory() (datatypes.Product_Item_Category, error)
	GetChildren() (datatypes.Billing_Item, error)
	GetChildrenWithActiveAgreement() (datatypes.Billing_Item, error)
	GetDowngradeItems() (datatypes.Product_Item, error)
	GetFilteredNextInvoiceChildren() (datatypes.Billing_Item, error)
	GetHourlyFlag() (bool, error)
	GetInvoiceItem() (datatypes.Billing_Invoice_Item, error)
	GetInvoiceItems() (datatypes.Billing_Invoice_Item, error)
	GetItem() (datatypes.Product_Item, error)
	GetLocation() (datatypes.Location, error)
	GetNextInvoiceChildren() (datatypes.Billing_Item, error)
	GetNextInvoiceTotalOneTimeAmount() (float64, error)
	GetNextInvoiceTotalOneTimeTaxAmount() (float64, error)
	GetNextInvoiceTotalRecurringAmount() (float64, error)
	GetNextInvoiceTotalRecurringTaxAmount() (float64, error)
	GetNonZeroNextInvoiceChildren() (datatypes.Billing_Item, error)
	GetOrderItem() (datatypes.Billing_Order_Item, error)
	GetOriginalLocation() (datatypes.Location, error)
	GetPackage() (datatypes.Product_Package, error)
	GetParent() (datatypes.Billing_Item, error)
	GetParentVirtualGuestBillingItem() (datatypes.Billing_Item_Virtual_Guest, error)
	GetPendingCancellationFlag() (bool, error)
	GetPendingOrderItem() (datatypes.Billing_Order_Item, error)
	GetProvisionTransaction() (datatypes.Provisioning_Version1_Transaction, error)
	GetSoftwareDescription() (datatypes.Software_Description, error)
	GetUpgradeItem() (datatypes.Product_Item, error)
	GetUpgradeItems() (datatypes.Product_Item, error)
}

type Billing_Item_Account_Media_Data_Transfer_Request interface {
	Billing_Item

	GetResource() (datatypes.Account_Media_Data_Transfer_Request, error)
}

type Billing_Item_Association_History interface {
	Entity

	GetAssociatedBillingItem() (datatypes.Billing_Item, error)
	GetBillingItem() (datatypes.Billing_Item, error)
}

type Billing_Item_Cancellation_Reason interface {
	Entity

	GetAllCancellationReasons() (datatypes.Billing_Item_Cancellation_Reason, error)
	GetObject() (datatypes.Billing_Item_Cancellation_Reason, error)

	GetBillingCancellationReasonCategory() (datatypes.Billing_Item_Cancellation_Reason_Category, error)
	GetBillingItems() (datatypes.Billing_Item, error)
	GetTranslatedReason() (string, error)
}

type Billing_Item_Cancellation_Reason_Category interface {
	Entity

	GetAllCancellationReasonCategories() (datatypes.Billing_Item_Cancellation_Reason_Category, error)
	GetObject() (datatypes.Billing_Item_Cancellation_Reason_Category, error)

	GetBillingCancellationReasons() (datatypes.Billing_Item_Cancellation_Reason, error)
}

type Billing_Item_Cancellation_Request interface {
	Entity

	CreateObject(templateObject datatypes.Billing_Item_Cancellation_Request) (datatypes.Billing_Item_Cancellation_Request, error)
	GetAllCancellationRequests() (datatypes.Billing_Item_Cancellation_Request, error)
	GetCancellationCutoffDate(accountId int, categoryCode string) (time.Time, error)
	GetObject() (datatypes.Billing_Item_Cancellation_Request, error)
	RemoveCancellationItem(itemId int) (bool, error)
	ValidateBillingItemForCancellation(billingItemId int) (bool, error)
	Void(closeRelatedTicketFlag bool) (bool, error)

	GetAccount() (datatypes.Account, error)
	GetItems() (datatypes.Billing_Item_Cancellation_Request_Item, error)
	GetStatus() (datatypes.Billing_Item_Cancellation_Request_Status, error)
	GetTicket() (datatypes.Ticket, error)
	GetUser() (datatypes.User_Customer, error)
}

type Billing_Item_Cancellation_Request_Item interface {
	Entity

	GetBillingItem() (datatypes.Billing_Item, error)
	GetCancellationRequest() (datatypes.Billing_Item_Cancellation_Request, error)
}

type Billing_Item_Cancellation_Request_Status interface {
	Entity
}

type Billing_Item_Ctc_Account interface {
	Billing_Item
}

type Billing_Item_Gateway_Appliance_Cluster interface {
	Billing_Item

	GetResource() (datatypes.Resource_Group, error)
}

type Billing_Item_Hardware interface {
	Billing_Item

	GetBillingCycleBandwidthUsage() (datatypes.Network_Bandwidth_Usage, error)
	GetBillingCyclePrivateBandwidthUsage() (datatypes.Network_Bandwidth_Usage, error)
	GetBillingCyclePrivateUsageIn() (float64, error)
	GetBillingCyclePrivateUsageOut() (float64, error)
	GetBillingCyclePrivateUsageTotal() (uint, error)
	GetBillingCyclePublicBandwidthUsage() (datatypes.Network_Bandwidth_Usage, error)
	GetBillingCyclePublicUsageIn() (float64, error)
	GetBillingCyclePublicUsageOut() (float64, error)
	GetBillingCyclePublicUsageTotal() (uint, error)
	GetLockboxNetworkStorage() (datatypes.Billing_Item_Network_Storage, error)
	GetMonitoringBillingItems() (datatypes.Billing_Item, error)
	GetResource() (datatypes.Hardware_Server, error)
}

type Billing_Item_Hardware_Colocation interface {
	Billing_Item_Hardware
}

type Billing_Item_Hardware_Component interface {
	Billing_Item

	GetResource() (datatypes.Hardware_Component, error)
}

type Billing_Item_Hardware_Security_Module interface {
	Billing_Item_Hardware
}

type Billing_Item_Hardware_Server interface {
	Billing_Item_Hardware
}

type Billing_Item_Link_ThePlanet interface {
	Entity

	GetBillingItem() (datatypes.Billing_Item, error)
	GetServiceProvider() (datatypes.Service_Provider, error)
}

type Billing_Item_Network_Application_Delivery_Controller interface {
	Billing_Item

	GetBandwidthAllotmentDetail() (datatypes.Network_Bandwidth_Version1_Allotment_Detail, error)
	GetResource() (datatypes.Network_Application_Delivery_Controller, error)
}

type Billing_Item_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress interface {
	Billing_Item

	GetResource() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress, error)
}

type Billing_Item_Network_Bandwidth interface {
	Billing_Item
}

type Billing_Item_Network_Firewall interface {
	Billing_Item

	GetResource() (datatypes.Network_Component_Firewall, error)
}

type Billing_Item_Network_Firewall_Module_Context interface {
	Billing_Item
}

type Billing_Item_Network_Interconnect interface {
	Billing_Item
}

type Billing_Item_Network_LoadBalancer interface {
	Billing_Item
}

type Billing_Item_Network_LoadBalancer_Global interface {
	Billing_Item

	GetResource() (datatypes.Network_LoadBalancer_Global_Account, error)
}

type Billing_Item_Network_LoadBalancer_VirtualIpAddress interface {
	Billing_Item

	GetResource() (datatypes.Network_LoadBalancer_VirtualIpAddress, error)
}

type Billing_Item_Network_Message_Delivery interface {
	Billing_Item

	GetResource() (datatypes.Network_Message_Delivery, error)
}

type Billing_Item_Network_Message_Queue interface {
	Billing_Item

	GetResource() (datatypes.Network_Message_Queue, error)
}

type Billing_Item_Network_Message_Queue_Delivery interface {
	Billing_Item_Network_Message_Queue
}

type Billing_Item_Network_PerformanceStorage_Iscsi interface {
	Billing_Item_Network_Storage
}

type Billing_Item_Network_PerformanceStorage_Nfs interface {
	Billing_Item_Network_Storage
}

type Billing_Item_Network_Storage interface {
	Billing_Item

	GetResource() (datatypes.Network_Storage, error)
}

type Billing_Item_Network_Storage_Hub interface {
	Billing_Item_Network_Storage
}

type Billing_Item_Network_Storage_Hub_Bandwidth interface {
	Billing_Item_Network_Storage
}

type Billing_Item_Network_Subnet interface {
	Billing_Item

	GetResource() (datatypes.Network_Subnet, error)
}

type Billing_Item_Network_Subnet_IpAddress_Global interface {
	Billing_Item_Network_Subnet
}

type Billing_Item_Network_Tunnel interface {
	Billing_Item

	GetResource() (datatypes.Network_Tunnel_Module_Context, error)
}

type Billing_Item_Network_Vlan interface {
	Billing_Item

	GetResource() (datatypes.Network_Vlan, error)
}

type Billing_Item_Software_Component interface {
	Billing_Item

	GetResource() (datatypes.Software_Component, error)
}

type Billing_Item_Software_Component_Analytics_Urchin interface {
	Billing_Item
}

type Billing_Item_Software_Component_ControlPanel interface {
	Billing_Item
}

type Billing_Item_Software_Component_ControlPanel_Parallels_Plesk_Billing interface {
	Billing_Item
}

type Billing_Item_Software_Component_OperatingSystem_Addon interface {
	Billing_Item
}

type Billing_Item_Software_Component_OperatingSystem_Addon_Citrix_Essentials interface {
	Billing_Item_Software_Component_OperatingSystem_Addon

	GetResource() (datatypes.Software_Component, error)
}

type Billing_Item_Software_Component_Virtual_OperatingSystem interface {
	Billing_Item
}

type Billing_Item_Software_Component_Virtual_OperatingSystem_Microsoft interface {
	Billing_Item_Software_Component_Virtual_OperatingSystem

	GetResource() (datatypes.Software_VirtualLicense, error)
}

type Billing_Item_Software_Component_Virtual_OperatingSystem_Redhat interface {
	Billing_Item_Software_Component_Virtual_OperatingSystem

	GetResource() (datatypes.Software_Component, error)
}

type Billing_Item_Software_License interface {
	Billing_Item

	GetResource() (datatypes.Software_AccountLicense, error)
}

type Billing_Item_Support interface {
	Billing_Item
}

type Billing_Item_User_Customer_External_Binding interface {
	Billing_Item

	GetResource() (datatypes.User_Customer_External_Binding, error)
}

type Billing_Item_Virtual_Dedicated_Rack interface {
	Billing_Item

	GetBillingCycleBandwidthUsage() (datatypes.Network_Bandwidth_Usage, error)
	GetBillingCyclePrivateBandwidthUsage() (datatypes.Network_Bandwidth_Usage, error)
	GetBillingCyclePrivateUsageIn() (float64, error)
	GetBillingCyclePrivateUsageOut() (float64, error)
	GetBillingCyclePrivateUsageTotal() (uint, error)
	GetBillingCyclePublicBandwidthUsage() (datatypes.Network_Bandwidth_Usage, error)
	GetBillingCyclePublicUsageIn() (float64, error)
	GetBillingCyclePublicUsageOut() (float64, error)
	GetBillingCyclePublicUsageTotal() (uint, error)
	GetResource() (datatypes.Network_Bandwidth_Version1_Allotment, error)
}

type Billing_Item_Virtual_Disk_Image interface {
	Billing_Item

	GetResource() (datatypes.Virtual_Disk_Image, error)
}

type Billing_Item_Virtual_Guest interface {
	Billing_Item

	GetBillingCycleBandwidthUsage() (datatypes.Network_Bandwidth_Usage, error)
	GetBillingCyclePrivateBandwidthUsage() (datatypes.Network_Bandwidth_Usage, error)
	GetBillingCyclePrivateUsageIn() (float64, error)
	GetBillingCyclePrivateUsageOut() (float64, error)
	GetBillingCyclePrivateUsageTotal() (uint, error)
	GetBillingCyclePublicBandwidthUsage() (datatypes.Network_Bandwidth_Usage, error)
	GetBillingCyclePublicUsageIn() (float64, error)
	GetBillingCyclePublicUsageOut() (float64, error)
	GetBillingCyclePublicUsageTotal() (uint, error)
	GetMonitoringBillingItems() (datatypes.Billing_Item, error)
	GetResource() (datatypes.Virtual_Guest, error)
}

type Billing_Item_Virtual_Host_Usage interface {
	Billing_Item

	GetResource() (datatypes.Hardware, error)
}

type Billing_Item_Workspace interface {
	Billing_Item
}

type Billing_Order interface {
	Entity

	ApproveModifiedOrder() (bool, error)
	GetAllObjects() (datatypes.Billing_Order, error)
	GetObject() (datatypes.Billing_Order, error)
	GetOrderStatuses() (datatypes.Container_Billing_Order_Status, error)
	GetPdf() ([]byte, error)
	GetPdfFilename() (string, error)
	GetRecalculatedOrderContainer(message string, ignoreDiscountsFlag bool) (datatypes.Container_Product_Order, error)
	GetReceipt() (datatypes.Container_Product_Order_Receipt, error)
	IsPendingEditApproval() (bool, error)

	GetAccount() (datatypes.Account, error)
	GetBrand() (datatypes.Brand, error)
	GetCart() (datatypes.Billing_Order_Cart, error)
	GetCoreRestrictedItems() (datatypes.Billing_Order_Item, error)
	GetCreditCardTransactions() (datatypes.Billing_Payment_Card_Transaction, error)
	GetExchangeRate() (datatypes.Billing_Currency_ExchangeRate, error)
	GetInitialInvoice() (datatypes.Billing_Invoice, error)
	GetItems() (datatypes.Billing_Order_Item, error)
	GetOrderApprovalDate() (time.Time, error)
	GetOrderNonServerMonthlyAmount() (float64, error)
	GetOrderServerMonthlyAmount() (float64, error)
	GetOrderTopLevelItems() (datatypes.Billing_Order_Item, error)
	GetOrderTotalAmount() (float64, error)
	GetOrderTotalOneTime() (float64, error)
	GetOrderTotalOneTimeAmount() (float64, error)
	GetOrderTotalOneTimeTaxAmount() (float64, error)
	GetOrderTotalRecurring() (float64, error)
	GetOrderTotalRecurringAmount() (float64, error)
	GetOrderTotalRecurringTaxAmount() (float64, error)
	GetOrderTotalSetupAmount() (float64, error)
	GetOrderType() (datatypes.Billing_Order_Type, error)
	GetPaypalTransactions() (datatypes.Billing_Payment_PayPal_Transaction, error)
	GetPresaleEvent() (datatypes.Sales_Presale_Event, error)
	GetQuote() (datatypes.Billing_Order_Quote, error)
	GetReferralPartner() (datatypes.Account, error)
	GetUpgradeRequestFlag() (bool, error)
	GetUserRecord() (datatypes.User_Customer, error)
}

type Billing_Order_Cart interface {
	Billing_Order_Quote

	CreateCart(orderData datatypes.Container_Product_Order) (int, error)
	DeleteCart() (bool, error)
	GetCartByCartKey(cartKey string) (datatypes.Billing_Order_Cart, error)
	GetObject() (datatypes.Billing_Order_Cart, error)
	GetPdf() ([]byte, error)
	GetRecalculatedOrderContainer(orderData datatypes.Container_Product_Order, orderBeingPlacedFlag bool) (datatypes.Container_Product_Order, error)
	UpdateCart(orderData datatypes.Container_Product_Order) (int, error)
}

type Billing_Order_Item interface {
	Entity

	GetObject() (datatypes.Billing_Order_Item, error)

	GetBillingItem() (datatypes.Billing_Item, error)
	GetBundledItems() (datatypes.Billing_Order_Item, error)
	GetCategory() (datatypes.Product_Item_Category, error)
	GetChildren() (datatypes.Billing_Order_Item, error)
	GetGlobalIdentifier() (string, error)
	GetHardwareGenericComponent() (datatypes.Hardware_Component_Model_Generic, error)
	GetItem() (datatypes.Product_Item, error)
	GetItemCategoryAnswers() (datatypes.Billing_Order_Item_Category_Answer, error)
	GetItemPrice() (datatypes.Product_Item_Price, error)
	GetLocation() (datatypes.Location, error)
	GetNextOrderChildren() (datatypes.Billing_Order_Item, error)
	GetOldBillingItem() (datatypes.Billing_Item, error)
	GetOrder() (datatypes.Billing_Order, error)
	GetOrderApprovalDate() (time.Time, error)
	GetPackage() (datatypes.Product_Package, error)
	GetParent() (datatypes.Billing_Order_Item, error)
	GetRedundantPowerSupplyCount() (uint, error)
	GetSoftwareDescription() (datatypes.Software_Description, error)
	GetStorageGroups() (datatypes.Configuration_Storage_Group_Order, error)
	GetTotalRecurringAmount() (float64, error)
	GetUpgradeItem() (datatypes.Product_Item, error)
}

type Billing_Order_Item_Category_Answer interface {
	Entity

	GetOrderItem() (datatypes.Billing_Order_Item, error)
	GetQuestion() (datatypes.Product_Item_Category_Question, error)
}

type Billing_Order_Note interface {
	Entity

	GetEmployee() (datatypes.User_Employee, error)
	GetOrder() (datatypes.Billing_Order, error)
}

type Billing_Order_Quote interface {
	Entity

	Claim(quoteKey string, quoteId int) (datatypes.Billing_Order_Quote, error)
	DeleteQuote() (datatypes.Billing_Order_Quote, error)
	GetObject() (datatypes.Billing_Order_Quote, error)
	GetPdf() ([]byte, error)
	GetQuoteByQuoteKey(quoteKey string) (datatypes.Billing_Order_Quote, error)
	GetRecalculatedOrderContainer(orderData datatypes.Container_Product_Order, orderBeingPlacedFlag bool) (datatypes.Container_Product_Order, error)
	PlaceOrder(orderData datatypes.Container_Product_Order) (datatypes.Container_Product_Order_Receipt, error)
	PlaceQuote(orderData datatypes.Container_Product_Order) (datatypes.Container_Product_Order, error)
	SaveQuote() (datatypes.Billing_Order_Quote, error)
	VerifyOrder(orderData datatypes.Container_Product_Order) (datatypes.Container_Product_Order, error)

	GetAccount() (datatypes.Account, error)
	GetOrder() (datatypes.Billing_Order, error)
	GetOrdersFromQuote() (datatypes.Billing_Order, error)
}

type Billing_Order_Type interface {
	Entity
}

type Billing_Payment_Card_ChangeRequest interface {
	Entity

	GetAccount() (datatypes.Account, error)
	GetAuthorizedCreditCardTransaction() (datatypes.Billing_Payment_Card_Transaction, error)
	GetCaptureCreditCardTransaction() (datatypes.Billing_Payment_Card_Transaction, error)
	GetTicketAttachmentReferences() (datatypes.Ticket_Attachment, error)
}

type Billing_Payment_Card_ManualPayment interface {
	Entity

	GetAccount() (datatypes.Account, error)
	GetAuthorizedCreditCardTransaction() (datatypes.Billing_Payment_Card_Transaction, error)
	GetAuthorizedPayPalTransaction() (datatypes.Billing_Payment_PayPal_Transaction, error)
	GetCaptureCreditCardTransaction() (datatypes.Billing_Payment_Card_Transaction, error)
	GetCapturePayPalTransaction() (datatypes.Billing_Payment_PayPal_Transaction, error)
	GetTicketAttachmentReferences() (datatypes.Ticket_Attachment, error)
}

type Billing_Payment_Card_Transaction interface {
	Billing_Payment_Transaction

	GetAccount() (datatypes.Account, error)
	GetOrder() (datatypes.Billing_Order, error)
}

type Billing_Payment_PayPal_Transaction interface {
	Billing_Payment_Transaction

	GetAccount() (datatypes.Account, error)
	GetOrder() (datatypes.Billing_Order, error)
}

type Billing_Payment_Processor interface {
	Entity

	GetBrandAssignments() (datatypes.Brand_Payment_Processor, error)
	GetOwnerAccount() (datatypes.Account, error)
	GetPaymentMethods() (datatypes.Billing_Payment_Processor_Method, error)
	GetType() (datatypes.Billing_Payment_Processor_Type, error)
}

type Billing_Payment_Processor_Method interface {
	Entity

	GetPaymentProcessor() (datatypes.Billing_Payment_Processor, error)
	GetPaymentType() (datatypes.Billing_Payment_Type, error)
}

type Billing_Payment_Processor_Type interface {
	Entity

	GetPaymentProcessors() (datatypes.Billing_Payment_Processor, error)
}

type Billing_Payment_Transaction interface {
	Entity
}

type Billing_Payment_Type interface {
	Entity
}
