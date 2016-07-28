package softlayer

import (
	"time"

	"github.ibm.com/riethm/gopherlayer/datatypes"
)

type McAfee_Epolicy_Orchestrator_Version36_Agent_Details interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_AccessProtection interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_Filter_Description interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_BlockedApplicationEvent interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_Event_Signature interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_IPSEvent interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_BlockedApplicationEvent interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_Event_Signature interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version36_Policy_Object interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version36_Product_Properties interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version45_Agent_Details interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version45_Event interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version45_Event_Filter_Description interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version45_Event_Version7 interface {
	McAfee_Epolicy_Orchestrator_Version45_Event
}

type McAfee_Epolicy_Orchestrator_Version45_Event_Version8 interface {
	McAfee_Epolicy_Orchestrator_Version45_Event
}

type McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version7 interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version8 interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version45_Policy_Object interface {
	Entity
}

type McAfee_Epolicy_Orchestrator_Version45_Product_Properties interface {
	Entity
}

type Abuse_Lockdown_Resource interface {
	Entity
}

type Account interface {
	Entity

	ActivatePartner(accountId string, hashCode string) (datatypes.Account, error)
	AddAchInformation(achInformation datatypes.Container_Billing_Info_Ach) (bool, error)
	AddReferralPartnerPaymentOption(paymentOption datatypes.Container_Referral_Partner_Payment_Option) (bool, error)
	AreVdrUpdatesBlockedForBilling() (bool, error)
	CancelPayPalTransaction(token string, payerId string) (bool, error)
	CompletePayPalTransaction(token string, payerId string) (string, error)
	CountHourlyInstances() (int, error)
	GetAccountBackupHistory(startDate time.Time, endDate time.Time, backupStatus string) (datatypes.Container_Network_Storage_Evault_WebCc_JobDetails, error)
	GetAccountTraitValue(keyName string) (string, error)
	GetActiveAlarms() (datatypes.Container_Monitoring_Alarm_History, error)
	GetActiveOutletPackages() (datatypes.Product_Package, error)
	GetActivePackages() (datatypes.Product_Package, error)
	GetActivePackagesByAttribute(attributeKeyName string) (datatypes.Product_Package, error)
	GetActivePrivateHostedCloudPackages() (datatypes.Product_Package, error)
	GetAggregatedUptimeGraph(startDate time.Time, endDate time.Time) (datatypes.Container_Graph, error)
	GetAlternateCreditCardData() (datatypes.Container_Account_Payment_Method_CreditCard, error)
	GetAttributeByType(attributeType string) (datatypes.Account_Attribute, error)
	GetAuxiliaryNotifications() (datatypes.Container_Utility_Message, error)
	GetAverageArchiveUsageMetricDataByDate(startDateTime time.Time, endDateTime time.Time) (float64, error)
	GetAveragePublicUsageMetricDataByDate(startDateTime time.Time, endDateTime time.Time) (float64, error)
	GetCurrentBackupStatisticsGraph(detailedGraph bool) (datatypes.Container_Account_Graph_Outputs, error)
	GetCurrentTicketStatisticsGraph(detailedGraph bool) (datatypes.Container_Account_Graph_Outputs, error)
	GetCurrentUser() (datatypes.User_Customer, error)
	GetDiskUsageMetricDataByDate(startDateTime time.Time, endDateTime time.Time) (datatypes.Metric_Tracking_Object_Data, error)
	GetDiskUsageMetricDataFromLegacyByDate(startDateTime time.Time, endDateTime time.Time) (datatypes.Metric_Tracking_Object_Data, error)
	GetDiskUsageMetricDataFromMetricTrackingObjectSystemByDate(startDateTime time.Time, endDateTime time.Time) (datatypes.Metric_Tracking_Object_Data, error)
	GetDiskUsageMetricImageByDate(startDateTime time.Time, endDateTime time.Time) (datatypes.Container_Account_Graph_Outputs, error)
	GetExecutiveSummaryPdf(pdfType string, historicalType string, startDate string, endDate string) ([]byte, error)
	GetFlexibleCreditProgramInfo(forNextBillCycle bool) (datatypes.Container_Account_Discount_Program, error)
	GetHistoricalBackupGraph(startDate time.Time, endDate time.Time) (datatypes.Container_Account_Graph_Outputs, error)
	GetHistoricalBandwidthGraph(startDate time.Time, endDate time.Time) (datatypes.Container_Account_Graph_Outputs, error)
	GetHistoricalTicketGraph(startDate time.Time, endDate time.Time) (datatypes.Container_Account_Graph_Outputs, error)
	GetHistoricalUptimeGraph(startDate time.Time, endDate time.Time) (datatypes.Container_Account_Graph_Outputs, error)
	GetLargestAllowedSubnetCidr(numberOfHosts int, locationId int) (int, error)
	GetNextInvoiceExcel(documentCreateDate time.Time) ([]byte, error)
	GetNextInvoicePdf(documentCreateDate time.Time) ([]byte, error)
	GetNextInvoicePdfDetailed(documentCreateDate time.Time) ([]byte, error)
	GetNextInvoiceZeroFeeItemCounts() (datatypes.Container_Product_Item_Category_ZeroFee_Count, error)
	GetObject() (datatypes.Account, error)
	GetPendingCreditCardChangeRequestData() (datatypes.Container_Account_Payment_Method_CreditCard, error)
	GetReferralPartnerCommissionForecast() (datatypes.Container_Referral_Partner_Commission, error)
	GetReferralPartnerCommissionHistory() (datatypes.Container_Referral_Partner_Commission, error)
	GetReferralPartnerCommissionPending() (datatypes.Container_Referral_Partner_Commission, error)
	GetSharedBlockDeviceTemplateGroups() (datatypes.Virtual_Guest_Block_Device_Template_Group, error)
	GetTechIncubatorProgramInfo(forNextBillCycle bool) (datatypes.Container_Account_Discount_Program, error)
	GetThirdPartyPoliciesAcceptanceStatus() (datatypes.Container_Policy_Acceptance, error)
	GetValidSecurityCertificateEntries() (datatypes.Security_Certificate_Entry, error)
	GetVmWareActiveAccountLicenseKeys() (string, error)
	GetWindowsUpdateStatus() (datatypes.Container_Utility_Microsoft_Windows_UpdateServices_Status, error)
	HasAttribute(attributeType string) (bool, error)
	HourlyInstanceLimit() (int, error)
	HourlyServerLimit() (int, error)
	LinkExternalAccount(externalAccountId string, authorizationToken string, externalServiceProviderKey string) error
	RemoveAlternateCreditCard() (bool, error)
	RequestCreditCardChange(request datatypes.Billing_Payment_Card_ChangeRequest, vatId string, paymentRoleName string, onlyChangeNicknameFlag bool) (datatypes.Billing_Payment_Card_ChangeRequest, error)
	RequestManualPayment(request datatypes.Billing_Payment_Card_ManualPayment) (datatypes.Billing_Payment_Card_ManualPayment, error)
	RequestManualPaymentUsingCreditCardOnFile(amount string, payWithAlternateCardFlag bool, note string) (datatypes.Billing_Payment_Card_ManualPayment, error)
	SetAbuseEmails(emails string) (bool, error)
	SetVlanSpan(enabled bool) (bool, error)
	SwapCreditCards() (bool, error)
	UpdateVpnUsersForResource(objectId int, objectType string) (bool, error)
	Validate(account datatypes.Account) (string, error)
	ValidateManualPaymentAmount(amount string) (bool, error)
}

type Account_AbuseEmail interface {
	Entity
}

type Account_Address interface {
	Entity

	CreateObject(templateObject datatypes.Account_Address) (datatypes.Account_Address, error)
	EditObject(templateObject datatypes.Account_Address) (bool, error)
	GetAllDataCenters() (datatypes.Account_Address, error)
	GetNetworkAddress(name string) (datatypes.Account_Address, error)
	GetObject() (datatypes.Account_Address, error)
}

type Account_Address_Type interface {
	Entity

	GetObject() (datatypes.Account_Address_Type, error)
}

type Account_Affiliation interface {
	Entity

	CreateObject(templateObject datatypes.Account_Affiliation) (datatypes.Account_Affiliation, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Account_Affiliation) (bool, error)
	GetAccountAffiliationsByAffiliateId(affiliateId string) (datatypes.Account_Affiliation, error)
	GetObject() (datatypes.Account_Affiliation, error)
}

type Account_Agreement interface {
	Entity

	GetObject() (datatypes.Account_Agreement, error)
}

type Account_Agreement_Status interface {
	Entity
}

type Account_Agreement_Type interface {
	Entity
}

type Account_Attachment_Employee interface {
	Entity
}

type Account_Attachment_Employee_Role interface {
	Entity
}

type Account_Attribute interface {
	Entity
}

type Account_Attribute_Type interface {
	Entity
}

type Account_Authentication_Attribute interface {
	Entity

	GetObject() (datatypes.Account_Authentication_Attribute, error)
}

type Account_Authentication_Attribute_Type interface {
	Entity

	GetAllObjects() (datatypes.Account_Attribute_Type, error)
	GetObject() (datatypes.Account_Authentication_Attribute_Type, error)
}

type Account_Authentication_OpenIdConnect_Option interface {
	Entity
}

type Account_Authentication_OpenIdConnect_RegistrationInformation interface {
	Entity
}

type Account_Authentication_Saml interface {
	Entity

	CreateObject(templateObject datatypes.Account_Authentication_Saml) (datatypes.Account_Authentication_Saml, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Account_Authentication_Saml) (bool, error)
	GetMetadata() (string, error)
	GetObject() (datatypes.Account_Authentication_Saml, error)
}

type Account_Classification_Group_Type interface {
	Entity
}

type Account_Contact interface {
	Entity

	CreateObject(templateObject datatypes.Account_Contact) (datatypes.Account_Contact, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Account_Contact) (bool, error)
	GetAllContactTypes() (datatypes.Account_Contact_Type, error)
	GetObject() (datatypes.Account_Contact, error)
}

type Account_Contact_Type interface {
	Entity
}

type Account_Historical_Report interface {
	Entity

	GetAccountHostUptimeGraphData(startDate string, endDate string) (datatypes.Container_Graph, error)
	GetAccountHostUptimeSummary(startDateTime string, endDateTime string) (datatypes.Container_Account_Historical_Summary, error)
	GetAccountUrlUptimeGraphData(startDate string, endDate string) (datatypes.Container_Graph, error)
	GetAccountUrlUptimeSummary(startDateTime string, endDateTime string) (datatypes.Container_Account_Historical_Summary, error)
	GetHostUptimeDetail(configurationValueId int, startDateTime string, endDateTime string) (datatypes.Container_Account_Historical_Summary_Detail, error)
	GetHostUptimeGraphData(configurationValueId int, startDate string, endDate string) (datatypes.Container_Graph, error)
	GetUrlUptimeDetail(configurationValueId int, startDateTime string, endDateTime string) (datatypes.Container_Account_Historical_Summary_Detail, error)
	GetUrlUptimeGraphData(configurationValueId int, startDate string, endDate string) (datatypes.Container_Graph, error)
}

type Account_Link interface {
	Entity
}

type Account_Link_Bluemix interface {
	Account_Link

	GetObject() (datatypes.Account_Link_Bluemix, error)
	GetSupportTierType() (string, error)
}

type Account_Link_OpenStack interface {
	Account_Link

	CreateOSDomain(request datatypes.Account_Link_OpenStack_LinkRequest) (datatypes.Account_Link_OpenStack_DomainCreationDetails, error)
	CreateOSProject(request datatypes.Account_Link_OpenStack_LinkRequest) (datatypes.Account_Link_OpenStack_ProjectCreationDetails, error)
	DeleteOSDomain(domainId string) (bool, error)
	DeleteOSProject(projectId string) (bool, error)
	DeleteObject() (bool, error)
	GetOSProject(projectId string) (datatypes.Account_Link_OpenStack_ProjectDetails, error)
	GetObject() (datatypes.Account_Link_OpenStack, error)
	ListOSProjects() (datatypes.Account_Link_OpenStack_ProjectDetails, error)
}

type Account_Link_OpenStack_DomainCreationDetails interface {
	Entity
}

type Account_Link_OpenStack_LinkRequest interface {
	Entity
}

type Account_Link_OpenStack_ProjectCreationDetails interface {
	Entity
}

type Account_Link_OpenStack_ProjectDetails interface {
	Entity
}

type Account_Link_ThePlanet interface {
	Account_Link
}

type Account_Link_Vendor interface {
	Entity
}

type Account_Lockdown_Request interface {
	Entity

	CancelRequest() error
	DisableLockedAccount(disableDate string) (int, error)
	DisconnectCompute(accountId int, disconnectDate string) (int, error)
	GetAccountHistory(accountId int) (datatypes.Account_Lockdown_Request, error)
	GetObject() (datatypes.Account_Lockdown_Request, error)
	ReconnectCompute(reconnectDate string) (int, error)
}

type Account_MasterServiceAgreement interface {
	Entity

	GetFile() (datatypes.Container_Utility_File_Entity, error)
	GetObject() (datatypes.Account_MasterServiceAgreement, error)
}

type Account_Media interface {
	Entity

	EditObject(templateObject datatypes.Account_Media) (bool, error)
	GetAllMediaTypes() (datatypes.Account_Media_Type, error)
	GetObject() (datatypes.Account_Media, error)
	RemoveMediaFromList(mediaTemplate datatypes.Account_Media) (int, error)
}

type Account_Media_Data_Transfer_Request interface {
	Entity

	EditObject(templateObject datatypes.Account_Media_Data_Transfer_Request) (bool, error)
	GetAllRequestStatuses() (datatypes.Account_Media_Data_Transfer_Request_Status, error)
	GetObject() (datatypes.Account_Media_Data_Transfer_Request, error)
}

type Account_Media_Data_Transfer_Request_Status interface {
	Entity
}

type Account_Media_Type interface {
	Entity
}

type Account_Network_Vlan_Span interface {
	Entity
}

type Account_Note interface {
	Entity

	CreateObject(templateObject datatypes.Account_Note) (datatypes.Account_Note, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Account_Note) (bool, error)
	GetObject() (datatypes.Account_Note, error)
}

type Account_Note_History interface {
	Entity
}

type Account_Note_Type interface {
	Entity

	CreateObject(templateObject datatypes.Account_Note_Type) (datatypes.Account_Note_Type, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Account_Note_Type) (bool, error)
	GetAllObjects() (datatypes.Account_Note_Type, error)
	GetObject() (datatypes.Account_Note_Type, error)
}

type Account_Partner_Referral_Prospect interface {
	User_Customer_Prospect

	CreateProspect(templateObject datatypes.Container_Referral_Partner_Prospect, commit bool) (datatypes.Account_Partner_Referral_Prospect, error)
	GetObject() (datatypes.Account_Partner_Referral_Prospect, error)
	GetSurveyQuestions() (datatypes.Survey_Question, error)
}

type Account_Password interface {
	Entity

	EditObject(templateObject datatypes.Account_Password) (bool, error)
	GetObject() (datatypes.Account_Password, error)
}

type Account_Password_Type interface {
	Entity
}

type Account_Regional_Registry_Detail interface {
	Entity

	CreateObject(templateObject datatypes.Account_Regional_Registry_Detail) (datatypes.Account_Regional_Registry_Detail, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Account_Regional_Registry_Detail) (bool, error)
	GetObject() (datatypes.Account_Regional_Registry_Detail, error)
	UpdateReferencedRegistrations() (datatypes.Container_Network_Subnet_Registration_TransactionDetails, error)
}

type Account_Regional_Registry_Detail_Property interface {
	Entity

	CreateObject(templateObject datatypes.Account_Regional_Registry_Detail_Property) (datatypes.Account_Regional_Registry_Detail_Property, error)
	CreateObjects(templateObjects datatypes.Account_Regional_Registry_Detail_Property) (datatypes.Account_Regional_Registry_Detail_Property, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Account_Regional_Registry_Detail_Property) (bool, error)
	EditObjects(templateObjects datatypes.Account_Regional_Registry_Detail_Property) (bool, error)
	GetObject() (datatypes.Account_Regional_Registry_Detail_Property, error)
}

type Account_Regional_Registry_Detail_Property_Type interface {
	Entity

	GetAllObjects() (datatypes.Account_Regional_Registry_Detail_Property_Type, error)
	GetObject() (datatypes.Account_Regional_Registry_Detail_Property_Type, error)
}

type Account_Regional_Registry_Detail_Type interface {
	Entity

	GetAllObjects() (datatypes.Account_Regional_Registry_Detail_Type, error)
	GetObject() (datatypes.Account_Regional_Registry_Detail_Type, error)
}

type Account_Regional_Registry_Detail_Version4_Person_Default interface {
	Account_Regional_Registry_Detail
}

type Account_Reports_Request interface {
	Entity

	CreateRequest(contact datatypes.Account_Contact, reason string, reportType string) (datatypes.Account_Reports_Request, error)
	GetAllObjects() (datatypes.Account_Reports_Request, error)
	GetObject() (datatypes.Account_Reports_Request, error)
	GetRequestByRequestKey(requestKey string) (datatypes.Account_Reports_Request, error)
	SendReportEmail(request datatypes.Account_Reports_Request) (bool, error)
	UpdateTicketOnDecline(request datatypes.Account_Reports_Request) (bool, error)
}

type Account_Rwhois_Handle interface {
	Entity
}

type Account_Shipment interface {
	Entity

	EditObject(templateObject datatypes.Account_Shipment) (bool, error)
	GetAllCouriers() (datatypes.Auxiliary_Shipping_Courier, error)
	GetAllCouriersByType(courierTypeKeyName string) (datatypes.Auxiliary_Shipping_Courier, error)
	GetAllShipmentStatuses() (datatypes.Account_Shipment_Status, error)
	GetAllShipmentTypes() (datatypes.Account_Shipment_Type, error)
	GetObject() (datatypes.Account_Shipment, error)
}

type Account_Shipment_Item interface {
	Entity

	EditObject(templateObject datatypes.Account_Shipment_Item) (bool, error)
	GetObject() (datatypes.Account_Shipment_Item, error)
}

type Account_Shipment_Item_Type interface {
	Entity

	GetObject() (datatypes.Account_Shipment_Item_Type, error)
}

type Account_Shipment_Resource_Type interface {
	Entity

	GetObject() (datatypes.Account_Shipment_Resource_Type, error)
}

type Account_Shipment_Status interface {
	Entity

	GetObject() (datatypes.Account_Shipment_Status, error)
}

type Account_Shipment_Tracking_Data interface {
	Entity

	CreateObject(templateObject datatypes.Account_Shipment_Tracking_Data) (datatypes.Account_Shipment_Tracking_Data, error)
	CreateObjects(templateObjects datatypes.Account_Shipment_Tracking_Data) (datatypes.Account_Shipment_Tracking_Data, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Account_Shipment_Tracking_Data) (bool, error)
	GetObject() (datatypes.Account_Shipment_Tracking_Data, error)
}

type Account_Shipment_Type interface {
	Entity

	GetObject() (datatypes.Account_Shipment_Type, error)
}

type Account_Status interface {
	Entity
}

type Auxiliary_Marketing_Event interface {
	Entity

	GetMarketingEvents() (datatypes.Auxiliary_Marketing_Event, error)
	GetObject() (datatypes.Auxiliary_Marketing_Event, error)
}

type Auxiliary_Network_Status interface {
	Entity

	GetNetworkStatus(target string) (datatypes.Container_Auxiliary_Network_Status_Reading, error)
}

type Auxiliary_Notification_Emergency interface {
	Entity

	GetAllObjects() (datatypes.Auxiliary_Notification_Emergency, error)
	GetCurrentNotifications() (datatypes.Auxiliary_Notification_Emergency, error)
	GetObject() (datatypes.Auxiliary_Notification_Emergency, error)
}

type Auxiliary_Notification_Emergency_Signature interface {
	Entity
}

type Auxiliary_Notification_Emergency_Status interface {
	Entity
}

type Auxiliary_Press_Release interface {
	Entity

	GetAllObjects() (datatypes.Auxiliary_Press_Release, error)
	GetObject() (datatypes.Auxiliary_Press_Release, error)
	GetRenderedPressRelease() (datatypes.Auxiliary_Press_Release, error)
	GetRenderedPressReleases(resultLimit string, year string) (datatypes.Auxiliary_Press_Release, error)
	GetWebsiteHighlightPressReleases() (datatypes.Auxiliary_Press_Release, error)
}

type Auxiliary_Press_Release_About interface {
	Entity

	GetObject() (datatypes.Auxiliary_Press_Release_About, error)
}

type Auxiliary_Press_Release_About_Press_Release interface {
	Entity

	GetObject() (datatypes.Auxiliary_Press_Release_About_Press_Release, error)
}

type Auxiliary_Press_Release_Contact interface {
	Entity

	GetObject() (datatypes.Auxiliary_Press_Release_Contact, error)
}

type Auxiliary_Press_Release_Contact_Press_Release interface {
	Entity

	GetObject() (datatypes.Auxiliary_Press_Release_Contact_Press_Release, error)
}

type Auxiliary_Press_Release_Content interface {
	Entity

	GetObject() (datatypes.Auxiliary_Press_Release_Content, error)
}

type Auxiliary_Press_Release_Media_Partner interface {
	Entity

	GetObject() (datatypes.Auxiliary_Press_Release_Media_Partner, error)
}

type Auxiliary_Press_Release_Media_Partner_Press_Release interface {
	Entity

	GetObject() (datatypes.Auxiliary_Press_Release_Media_Partner_Press_Release, error)
}

type Auxiliary_Shipping_Courier interface {
	Entity
}

type Auxiliary_Shipping_Courier_Type interface {
	Entity

	GetObject() (datatypes.Auxiliary_Shipping_Courier_Type, error)
	GetTypeByKeyName(keyName string) (datatypes.Auxiliary_Shipping_Courier_Type, error)
}

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
}

type Billing_Info interface {
	Entity

	GetObject() (datatypes.Billing_Info, error)
}

type Billing_Info_Ach interface {
	Entity
}

type Billing_Info_Cycle interface {
	Entity
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
}

type Billing_Invoice_Item interface {
	Entity

	GetObject() (datatypes.Billing_Invoice_Item, error)
}

type Billing_Invoice_Item_Hardware interface {
	Billing_Invoice_Item
}

type Billing_Invoice_Item_Tax_Info interface {
	Entity
}

type Billing_Invoice_Next interface {
	Entity

	GetExcel(documentCreateDate time.Time) ([]byte, error)
	GetPdf(documentCreateDate time.Time) ([]byte, error)
	GetPdfDetailed(documentCreateDate time.Time) ([]byte, error)
}

type Billing_Invoice_Receivable_Payment interface {
	Entity
}

type Billing_Invoice_Tax_Info interface {
	Entity
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
}

type Billing_Item_Account_Media_Data_Transfer_Request interface {
	Billing_Item
}

type Billing_Item_Association_History interface {
	Entity
}

type Billing_Item_Cancellation_Reason interface {
	Entity

	GetAllCancellationReasons() (datatypes.Billing_Item_Cancellation_Reason, error)
	GetObject() (datatypes.Billing_Item_Cancellation_Reason, error)
}

type Billing_Item_Cancellation_Reason_Category interface {
	Entity

	GetAllCancellationReasonCategories() (datatypes.Billing_Item_Cancellation_Reason_Category, error)
	GetObject() (datatypes.Billing_Item_Cancellation_Reason_Category, error)
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
}

type Billing_Item_Cancellation_Request_Item interface {
	Entity
}

type Billing_Item_Cancellation_Request_Status interface {
	Entity
}

type Billing_Item_Ctc_Account interface {
	Billing_Item
}

type Billing_Item_Gateway_Appliance_Cluster interface {
	Billing_Item
}

type Billing_Item_Hardware interface {
	Billing_Item
}

type Billing_Item_Hardware_Colocation interface {
	Billing_Item_Hardware
}

type Billing_Item_Hardware_Component interface {
	Billing_Item
}

type Billing_Item_Hardware_Security_Module interface {
	Billing_Item_Hardware
}

type Billing_Item_Hardware_Server interface {
	Billing_Item_Hardware
}

type Billing_Item_Link_ThePlanet interface {
	Entity
}

type Billing_Item_Network_Application_Delivery_Controller interface {
	Billing_Item
}

type Billing_Item_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress interface {
	Billing_Item
}

type Billing_Item_Network_Bandwidth interface {
	Billing_Item
}

type Billing_Item_Network_Firewall interface {
	Billing_Item
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
}

type Billing_Item_Network_LoadBalancer_VirtualIpAddress interface {
	Billing_Item
}

type Billing_Item_Network_Message_Delivery interface {
	Billing_Item
}

type Billing_Item_Network_Message_Queue interface {
	Billing_Item
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
}

type Billing_Item_Network_Storage_Hub interface {
	Billing_Item_Network_Storage
}

type Billing_Item_Network_Storage_Hub_Bandwidth interface {
	Billing_Item_Network_Storage
}

type Billing_Item_Network_Subnet interface {
	Billing_Item
}

type Billing_Item_Network_Subnet_IpAddress_Global interface {
	Billing_Item_Network_Subnet
}

type Billing_Item_Network_Tunnel interface {
	Billing_Item
}

type Billing_Item_Network_Vlan interface {
	Billing_Item
}

type Billing_Item_Software_Component interface {
	Billing_Item
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
}

type Billing_Item_Software_Component_Virtual_OperatingSystem interface {
	Billing_Item
}

type Billing_Item_Software_Component_Virtual_OperatingSystem_Microsoft interface {
	Billing_Item_Software_Component_Virtual_OperatingSystem
}

type Billing_Item_Software_Component_Virtual_OperatingSystem_Redhat interface {
	Billing_Item_Software_Component_Virtual_OperatingSystem
}

type Billing_Item_Software_License interface {
	Billing_Item
}

type Billing_Item_Support interface {
	Billing_Item
}

type Billing_Item_User_Customer_External_Binding interface {
	Billing_Item
}

type Billing_Item_Virtual_Dedicated_Rack interface {
	Billing_Item
}

type Billing_Item_Virtual_Disk_Image interface {
	Billing_Item
}

type Billing_Item_Virtual_Guest interface {
	Billing_Item
}

type Billing_Item_Virtual_Host_Usage interface {
	Billing_Item
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
}

type Billing_Order_Item_Category_Answer interface {
	Entity
}

type Billing_Order_Note interface {
	Entity
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
}

type Billing_Order_Type interface {
	Entity
}

type Billing_Payment_Card_ChangeRequest interface {
	Entity
}

type Billing_Payment_Card_ManualPayment interface {
	Entity
}

type Billing_Payment_Card_Transaction interface {
	Billing_Payment_Transaction
}

type Billing_Payment_PayPal_Transaction interface {
	Billing_Payment_Transaction
}

type Billing_Payment_Processor interface {
	Entity
}

type Billing_Payment_Processor_Method interface {
	Entity
}

type Billing_Payment_Processor_Type interface {
	Entity
}

type Billing_Payment_Transaction interface {
	Entity
}

type Billing_Payment_Type interface {
	Entity
}

type Brand interface {
	Entity

	CreateCustomerAccount(account datatypes.Account, bypassDuplicateAccountCheck bool) (datatypes.Account, error)
	CreateObject(templateObject datatypes.Brand) (datatypes.Brand, error)
	GetAllTicketSubjects(account datatypes.Account) (datatypes.Ticket_Subject, error)
	GetContactInformation() (datatypes.Brand_Contact, error)
	GetMerchantName() (string, error)
	GetObject() (datatypes.Brand, error)
	GetToken(userId int) (string, error)
}

type Brand_Attribute interface {
	Entity
}

type Brand_Contact interface {
	Entity
}

type Brand_Contact_Type interface {
	Entity
}

type Brand_Payment_Processor interface {
	Entity
}

type Brand_Restriction_Location_CustomerCountry interface {
	Entity

	GetAllObjects() (datatypes.Brand_Restriction_Location_CustomerCountry, error)
	GetObject() (datatypes.Brand_Restriction_Location_CustomerCountry, error)
}

type Catalyst_Affiliate interface {
	Entity
}

type Catalyst_Company_Type interface {
	Entity

	GetAllObjects() (datatypes.Catalyst_Company_Type, error)
	GetObject() (datatypes.Catalyst_Company_Type, error)
}

type Catalyst_Enrollment interface {
	Entity

	GetAffiliates() (datatypes.Catalyst_Affiliate, error)
	GetCompanyTypes() (datatypes.Catalyst_Company_Type, error)
	GetEnrollmentRequestAnnualRevenueOptions() (datatypes.Catalyst_Enrollment_Request_Container_AnswerOption, error)
	GetEnrollmentRequestUserCountOptions() (datatypes.Catalyst_Enrollment_Request_Container_AnswerOption, error)
	GetEnrollmentRequestYearsInOperationOptions() (datatypes.Catalyst_Enrollment_Request_Container_AnswerOption, error)
	GetObject() (datatypes.Catalyst_Enrollment, error)
	RequestManualEnrollment(request datatypes.Container_Catalyst_ManualEnrollmentRequest) error
	RequestSelfEnrollment(enrollmentRequest datatypes.Catalyst_Enrollment_Request) (datatypes.Account, error)
}

type Catalyst_Enrollment_Request interface {
	Entity
}

type Catalyst_Enrollment_Request_Container_AnswerOption interface {
	Entity
}

type Compliance_Report_Type interface {
	Entity

	GetAllObjects() (datatypes.Compliance_Report_Type, error)
	GetObject() (datatypes.Compliance_Report_Type, error)
}

type Configuration_Storage_Filesystem_Type interface {
	Entity
}

type Configuration_Storage_Group_Array_Type interface {
	Entity

	GetAllObjects() (datatypes.Configuration_Storage_Group_Array_Type, error)
	GetObject() (datatypes.Configuration_Storage_Group_Array_Type, error)
}

type Configuration_Storage_Group_Order interface {
	Entity
}

type Configuration_Storage_Group_Template_Group interface {
	Entity
}

type Configuration_Template interface {
	Entity

	CopyTemplate(templateObject datatypes.Configuration_Template) (datatypes.Configuration_Template, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Configuration_Template) (bool, error)
	GetAllObjects() (datatypes.Configuration_Template, error)
	GetObject() (datatypes.Configuration_Template, error)
	UpdateDefaultValues(configurationValues datatypes.Configuration_Template_Section_Definition_Value) (bool, error)
}

type Configuration_Template_Attribute interface {
	Entity
}

type Configuration_Template_Section interface {
	Entity

	GetObject() (datatypes.Configuration_Template_Section, error)
	HasSubSections() (bool, error)
}

type Configuration_Template_Section_Attribute interface {
	Entity
}

type Configuration_Template_Section_Definition interface {
	Entity

	GetObject() (datatypes.Configuration_Template_Section_Definition, error)
}

type Configuration_Template_Section_Definition_Attribute interface {
	Entity
}

type Configuration_Template_Section_Definition_Attribute_Type interface {
	Entity
}

type Configuration_Template_Section_Definition_Group interface {
	Entity

	GetAllGroups() (datatypes.Configuration_Template_Section_Definition_Group, error)
	GetObject() (datatypes.Configuration_Template_Section_Definition_Group, error)
}

type Configuration_Template_Section_Definition_Type interface {
	Entity

	GetObject() (datatypes.Configuration_Template_Section_Definition_Type, error)
}

type Configuration_Template_Section_Definition_Value interface {
	Entity

	GetObject() (datatypes.Configuration_Template_Section_Definition_Value, error)
}

type Configuration_Template_Section_Profile interface {
	Entity

	GetObject() (datatypes.Configuration_Template_Section_Profile, error)
}

type Configuration_Template_Section_Reference interface {
	Entity

	GetObject() (datatypes.Configuration_Template_Section_Reference, error)
}

type Configuration_Template_Section_Type interface {
	Entity

	GetObject() (datatypes.Configuration_Template_Section_Type, error)
}

type Configuration_Template_Type interface {
	Entity

	GetObject() (datatypes.Configuration_Template_Type, error)
}

type Container_Account_Discount_Program interface {
	Entity
}

type Container_Account_Graph_Outputs interface {
	Entity
}

type Container_Account_Historical_Summary interface {
	Entity
}

type Container_Account_Historical_Summary_Detail interface {
	Entity
}

type Container_Account_Historical_Summary_Detail_Uptime interface {
	Container_Account_Historical_Summary_Detail
}

type Container_Account_Historical_Summary_Uptime interface {
	Container_Account_Historical_Summary
}

type Container_Account_Payment_Method_CreditCard interface {
	Entity
}

type Container_Auxiliary_Network_Status_Reading interface {
	Entity
}

type Container_Bandwidth_GraphInputs interface {
	Entity
}

type Container_Bandwidth_GraphOutputs interface {
	Entity
}

type Container_Bandwidth_GraphOutputsExtended interface {
	Entity
}

type Container_Bandwidth_Projection interface {
	Entity
}

type Container_Billing_Currency_Format interface {
	Entity
}

type Container_Billing_Info_Ach interface {
	Entity
}

type Container_Billing_Invoice_Email interface {
	Entity
}

type Container_Billing_Order_Status interface {
	Entity
}

type Container_Catalyst_ManualEnrollmentRequest interface {
	Entity
}

type Container_Collection_Locale_CountryCode interface {
	Entity
}

type Container_Collection_Locale_StateCode interface {
	Entity
}

type Container_Disk_Image_Capture_Template interface {
	Entity
}

type Container_Disk_Image_Capture_Template_Volume interface {
	Entity
}

type Container_Disk_Image_Capture_Template_Volume_Partition interface {
	Entity
}

type Container_Dns_Domain_Registration_Contact interface {
	Entity
}

type Container_Dns_Domain_Registration_ExtendedAttribute interface {
	Entity
}

type Container_Dns_Domain_Registration_ExtendedAttribute_Configuration interface {
	Entity
}

type Container_Dns_Domain_Registration_ExtendedAttribute_Option interface {
	Entity
}

type Container_Dns_Domain_Registration_ExtendedAttribute_Option_Require interface {
	Entity
}

type Container_Dns_Domain_Registration_Information interface {
	Entity
}

type Container_Dns_Domain_Registration_List interface {
	Entity
}

type Container_Dns_Domain_Registration_Lookup interface {
	Entity
}

type Container_Dns_Domain_Registration_Lookup_Items interface {
	Entity
}

type Container_Dns_Domain_Registration_Nameserver interface {
	Entity
}

type Container_Dns_Domain_Registration_Nameserver_List interface {
	Entity
}

type Container_Dns_Domain_Registration_Registrant_Verification_StatusDetail interface {
	Entity
}

type Container_Dns_Domain_Registration_Transfer_Information interface {
	Entity
}

type Container_Exception interface {
	Entity
}

type Container_Graph interface {
	Entity
}

type Container_Graph_Option interface {
	Entity
}

type Container_Graph_Plot interface {
	Entity
}

type Container_Graph_Plot_Coordinate interface {
	Entity
}

type Container_Hardware_Configuration interface {
	Entity
}

type Container_Hardware_Configuration_Option interface {
	Entity
}

type Container_Hardware_MassUpdate interface {
	Entity
}

type Container_Hardware_Server_Configuration interface {
	Entity
}

type Container_Hardware_Server_Details interface {
	Entity
}

type Container_KnowledgeLayer_QuestionAnswer interface {
	Entity
}

type Container_Message interface {
	Entity
}

type Container_Metric_Data_Type interface {
	Entity
}

type Container_Metric_Tracking_Object_Details interface {
	Entity
}

type Container_Metric_Tracking_Object_Summary interface {
	Entity
}

type Container_Metric_Tracking_Object_Virtual_Host_Details interface {
	Container_Metric_Tracking_Object_Details
}

type Container_Metric_Tracking_Object_Virtual_Host_Summary interface {
	Container_Metric_Tracking_Object_Summary
}

type Container_Monitoring_Alarm_History interface {
	Entity
}

type Container_Monitoring_Graph_Outputs interface {
	Entity
}

type Container_Network_Authentication_Data interface {
	Entity
}

type Container_Network_Bandwidth_Data_Summary interface {
	Entity
}

type Container_Network_Bandwidth_Version1_Usage interface {
	Entity
}

type Container_Network_ContentDelivery_Authentication_Directory interface {
	Entity
}

type Container_Network_ContentDelivery_Authentication_Parameter interface {
	Entity
}

type Container_Network_ContentDelivery_Authentication_ServiceEndpoint interface {
	Entity
}

type Container_Network_ContentDelivery_Bandwidth_PointsOfPresence_Summary interface {
	Entity
}

type Container_Network_ContentDelivery_Bandwidth_Summary interface {
	Entity
}

type Container_Network_ContentDelivery_Bandwidth_Summary_Detail interface {
	Container_Network_ContentDelivery_Bandwidth_Summary
}

type Container_Network_ContentDelivery_OriginPull_Mapping interface {
	Entity
}

type Container_Network_ContentDelivery_PointsOfPresence interface {
	Entity
}

type Container_Network_ContentDelivery_PurgeService_Response interface {
	Entity
}

type Container_Network_ContentDelivery_Report_Usage interface {
	Entity
}

type Container_Network_ContentDelivery_SupportedProtocol interface {
	Entity
}

type Container_Network_Directory_Listing interface {
	Entity
}

type Container_Network_IntrusionProtection_Event interface {
	Entity
}

type Container_Network_IntrusionProtection_Statistic interface {
	Entity
}

type Container_Network_IntrusionProtection_Statistics interface {
	Entity
}

type Container_Network_IntrusionProtection_SubnetReport interface {
	Entity
}

type Container_Network_LoadBalancer_StatusEntry interface {
	Entity
}

type Container_Network_Media_Information interface {
	Entity
}

type Container_Network_Media_Transcode_Job_Watermark interface {
	Entity
}

type Container_Network_Media_Transcode_Job_Watermark_Position interface {
	Entity
}

type Container_Network_Media_Transcode_Preset interface {
	Entity
}

type Container_Network_Media_Transcode_Preset_Element interface {
	Entity
}

type Container_Network_Media_Transcode_Preset_Element_Option interface {
	Entity
}

type Container_Network_Message_Delivery_Email interface {
	Entity
}

type Container_Network_Message_Delivery_Email_Sendgrid_Account_Overview interface {
	Entity
}

type Container_Network_Message_Delivery_Email_Sendgrid_Customer_Profile interface {
	Entity
}

type Container_Network_Message_Delivery_Email_Sendgrid_List_Entry interface {
	Entity
}

type Container_Network_Message_Delivery_Email_Sendgrid_Statistics interface {
	Entity
}

type Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Graph interface {
	Entity
}

type Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Options interface {
	Entity
}

type Container_Network_Port_Statistic interface {
	Entity
}

type Container_Network_Service_Resource_ObjectStorage_ConnectionInformation interface {
	Entity
}

type Container_Network_Storage_Backup_Evault_WebCc_Authentication_Details interface {
	Entity
}

type Container_Network_Storage_Evault_Vault_Task interface {
	Entity
}

type Container_Network_Storage_Evault_WebCc_AgentStatus interface {
	Entity
}

type Container_Network_Storage_Evault_WebCc_BackupResults interface {
	Entity
}

type Container_Network_Storage_Evault_WebCc_JobDetails interface {
	Entity
}

type Container_Network_Storage_Host interface {
	Entity
}

type Container_Network_Storage_Hub_ObjectStorage_ContentDeliveryUrl interface {
	Entity
}

type Container_Network_Storage_Hub_ObjectStorage_Endpoint interface {
	Entity
}

type Container_Network_Storage_Hub_ObjectStorage_File interface {
	Container_Utility_File_Entity
}

type Container_Network_Storage_Hub_ObjectStorage_Folder interface {
	Entity
}

type Container_Network_Storage_Hub_ObjectStorage_Node interface {
	Entity
}

type Container_Network_Storage_NetworkConnectionInformation interface {
	Entity
}

type Container_Network_Subnet_IpAddress interface {
	Entity
}

type Container_Network_Subnet_Registration_SubnetReference interface {
	Entity
}

type Container_Network_Subnet_Registration_TransactionDetails interface {
	Entity
}

type Container_Notification_Mass_Filter_TemplateKey interface {
	Entity
}

type Container_Notification_Mass_Filter_TemplateValue interface {
	Entity
}

type Container_Policy_Acceptance interface {
	Entity
}

type Container_Product_Item_Category interface {
	Entity
}

type Container_Product_Item_Category_Question_Answer interface {
	Entity
}

type Container_Product_Item_Category_ZeroFee_Count interface {
	Entity
}

type Container_Product_Item_Discount_Program interface {
	Entity
}

type Container_Product_Order interface {
	Entity
}

type Container_Product_Order_Account_Media_Data_Transfer_Request interface {
	Container_Product_Order
}

type Container_Product_Order_Attribute_Address interface {
	Entity
}

type Container_Product_Order_Attribute_Contact interface {
	Entity
}

type Container_Product_Order_Attribute_Organization interface {
	Entity
}

type Container_Product_Order_Billing_Information interface {
	Entity
}

type Container_Product_Order_Dns_Domain_Registration interface {
	Container_Product_Order
}

type Container_Product_Order_Dns_Domain_Reseller interface {
	Container_Product_Order
}

type Container_Product_Order_Gateway_Appliance_Cluster interface {
	Container_Product_Order
}

type Container_Product_Order_Hardware_Security_Module interface {
	Container_Product_Order_Hardware_Server
}

type Container_Product_Order_Hardware_Server interface {
	Container_Product_Order
}

type Container_Product_Order_Hardware_Server_Colocation interface {
	Container_Product_Order_Hardware_Server
}

type Container_Product_Order_Hardware_Server_Gateway_Appliance interface {
	Container_Product_Order_Hardware_Server
}

type Container_Product_Order_Hardware_Server_Upgrade interface {
	Container_Product_Order_Hardware_Server
}

type Container_Product_Order_Monitoring_Package interface {
	Container_Product_Order
}

type Container_Product_Order_MultiConfiguration interface {
	Container_Product_Order
}

type Container_Product_Order_MultiConfiguration_Tornado interface {
	Container_Product_Order_MultiConfiguration
}

type Container_Product_Order_Network interface {
	Entity
}

type Container_Product_Order_Network_Application_Delivery_Controller interface {
	Container_Product_Order
}

type Container_Product_Order_Network_ContentDelivery_Account interface {
	Container_Product_Order
}

type Container_Product_Order_Network_ContentDelivery_Account_Upgrade interface {
	Container_Product_Order
}

type Container_Product_Order_Network_LoadBalancer interface {
	Container_Product_Order
}

type Container_Product_Order_Network_LoadBalancer_Global interface {
	Container_Product_Order
}

type Container_Product_Order_Network_Message_Delivery interface {
	Container_Product_Order
}

type Container_Product_Order_Network_Message_Queue interface {
	Container_Product_Order
}

type Container_Product_Order_Network_PerformanceStorage interface {
	Container_Product_Order
}

type Container_Product_Order_Network_PerformanceStorage_Iscsi interface {
	Container_Product_Order_Network_PerformanceStorage
}

type Container_Product_Order_Network_PerformanceStorage_Nfs interface {
	Container_Product_Order_Network_PerformanceStorage
}

type Container_Product_Order_Network_Protection_Firewall interface {
	Container_Product_Order
}

type Container_Product_Order_Network_Protection_Firewall_Dedicated interface {
	Container_Product_Order
}

type Container_Product_Order_Network_Storage_Backup_Evault_Plugin interface {
	Container_Product_Order
}

type Container_Product_Order_Network_Storage_Backup_Evault_Vault interface {
	Container_Product_Order
}

type Container_Product_Order_Network_Storage_Enterprise interface {
	Container_Product_Order
}

type Container_Product_Order_Network_Storage_Enterprise_SnapshotSpace interface {
	Container_Product_Order
}

type Container_Product_Order_Network_Storage_Enterprise_SnapshotSpace_Upgrade interface {
	Container_Product_Order_Network_Storage_Enterprise_SnapshotSpace
}

type Container_Product_Order_Network_Storage_Hub interface {
	Container_Product_Order
}

type Container_Product_Order_Network_Storage_Hub_Datacenter interface {
	Entity
}

type Container_Product_Order_Network_Storage_Iscsi interface {
	Container_Product_Order
}

type Container_Product_Order_Network_Storage_Iscsi_Replication interface {
	Container_Product_Order
}

type Container_Product_Order_Network_Storage_Iscsi_SnapshotSpace interface {
	Container_Product_Order
}

type Container_Product_Order_Network_Storage_Modification interface {
	Container_Product_Order
}

type Container_Product_Order_Network_Storage_Nas interface {
	Container_Product_Order
}

type Container_Product_Order_Network_Storage_Object interface {
	Container_Product_Order
}

type Container_Product_Order_Network_Subnet interface {
	Container_Product_Order
}

type Container_Product_Order_Network_Tunnel_Ipsec interface {
	Container_Product_Order
}

type Container_Product_Order_Network_Vlan interface {
	Container_Product_Order
}

type Container_Product_Order_Network_Vlans interface {
	Entity
}

type Container_Product_Order_Property interface {
	Entity
}

type Container_Product_Order_Receipt interface {
	Entity
}

type Container_Product_Order_Security_Certificate interface {
	Container_Product_Order
}

type Container_Product_Order_Software_Component_Virtual interface {
	Container_Product_Order
}

type Container_Product_Order_Software_License interface {
	Container_Product_Order
}

type Container_Product_Order_SshKeys interface {
	Entity
}

type Container_Product_Order_Storage_Group interface {
	Entity
}

type Container_Product_Order_Storage_Group_Partition interface {
	Entity
}

type Container_Product_Order_User_Customer_External_Binding interface {
	Container_Product_Order
}

type Container_Product_Order_Virtual_Disk_Image interface {
	Container_Product_Order
}

type Container_Product_Order_Virtual_Guest interface {
	Container_Product_Order_Hardware_Server
}

type Container_Product_Order_Virtual_Guest_Upgrade interface {
	Container_Product_Order_Virtual_Guest
}

type Container_Provisioning_Maintenance_Window interface {
	Entity
}

type Container_Referral_Partner_Commission interface {
	Entity
}

type Container_Referral_Partner_Payment_Option interface {
	Entity
}

type Container_Referral_Partner_Prospect interface {
	Entity
}

type Container_RemoteManagement_Graphs_SensorSpeed interface {
	Entity
}

type Container_RemoteManagement_Graphs_SensorTemperature interface {
	Entity
}

type Container_RemoteManagement_PmInfo interface {
	Entity
}

type Container_RemoteManagement_SensorReading interface {
	Entity
}

type Container_RemoteManagement_SensorReadingsWithGraphs interface {
	Entity
}

type Container_Resource_Metadata_ServiceResource interface {
	Entity
}

type Container_Search_ObjectType interface {
	Entity
}

type Container_Search_ObjectType_Property interface {
	Entity
}

type Container_Search_Result interface {
	Entity
}

type Container_Software_Component_HostIps_Policy interface {
	Entity
}

type Container_Tax_Cache interface {
	Entity
}

type Container_Tax_Cache_Item interface {
	Entity
}

type Container_Tax_Rates interface {
	Entity
}

type Container_Ticket_GraphInputs interface {
	Entity
}

type Container_Ticket_GraphOutputs interface {
	Entity
}

type Container_Ticket_Priority interface {
	Entity
}

type Container_Ticket_Survey_Preference interface {
	Entity
}

type Container_User_Authentication_Token interface {
	Entity
}

type Container_User_Customer_External_Binding interface {
	Entity
}

type Container_User_Customer_External_Binding_Phone interface {
	Container_User_Customer_External_Binding
}

type Container_User_Customer_External_Binding_Phone_Mode interface {
	Entity
}

type Container_User_Customer_External_Binding_Totp interface {
	Container_User_Customer_External_Binding
}

type Container_User_Customer_External_Binding_Vendor interface {
	Entity
}

type Container_User_Customer_External_Binding_Verisign interface {
	Container_User_Customer_External_Binding
}

type Container_User_Customer_PasswordSet interface {
	Entity
}

type Container_User_Customer_Portal_MobileToken interface {
	Container_User_Customer_Portal_Token
}

type Container_User_Customer_Portal_Token interface {
	Entity
}

type Container_User_Data_Phone interface {
	Entity
}

type Container_User_Employee_External_Binding_Verisign interface {
	Entity
}

type Container_Utility_File_Attachment interface {
	Entity
}

type Container_Utility_File_Descriptor interface {
	Entity
}

type Container_Utility_File_Entity interface {
	Entity
}

type Container_Utility_Message interface {
	Entity
}

type Container_Utility_Microsoft_Windows_UpdateServices_Status interface {
	Entity
}

type Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem interface {
	Entity
}

type Container_Utility_Network_Firewall_Rule_Attribute interface {
	Entity
}

type Container_Utility_Network_Subnet_Mask_Generic_Detail interface {
	Entity
}

type Container_Virtual_Guest_Block_Device_Template_Configuration interface {
	Entity
}

type Container_Virtual_Guest_Configuration interface {
	Entity
}

type Container_Virtual_Guest_Configuration_Option interface {
	Entity
}

type Dns_Domain interface {
	Entity

	CreateARecord(host string, data string, ttl int) (datatypes.Dns_Domain_ResourceRecord_AType, error)
	CreateAaaaRecord(host string, data string, ttl int) (datatypes.Dns_Domain_ResourceRecord_AaaaType, error)
	CreateCnameRecord(host string, data string, ttl int) (datatypes.Dns_Domain_ResourceRecord_CnameType, error)
	CreateMxRecord(host string, data string, ttl int, mxPriority int) (datatypes.Dns_Domain_ResourceRecord_MxType, error)
	CreateNsRecord(host string, data string, ttl int) (datatypes.Dns_Domain_ResourceRecord_NsType, error)
	CreateObject(templateObject datatypes.Dns_Domain) (datatypes.Dns_Domain, error)
	CreateObjects(templateObjects datatypes.Dns_Domain) (datatypes.Dns_Domain, error)
	CreatePtrRecord(ipAddress string, ptrRecord string, ttl int) (datatypes.Dns_Domain_ResourceRecord, error)
	CreateSpfRecord(host string, data string, ttl int) (datatypes.Dns_Domain_ResourceRecord_SpfType, error)
	CreateTxtRecord(host string, data string, ttl int) (datatypes.Dns_Domain_ResourceRecord_TxtType, error)
	DeleteObject() (bool, error)
	GetByDomainName(name string) (datatypes.Dns_Domain, error)
	GetObject() (datatypes.Dns_Domain, error)
	GetZoneFileContents() (string, error)
}

type Dns_Domain_Forward interface {
	Dns_Domain
}

type Dns_Domain_Registration interface {
	Entity

	AddNameserversToDomain(nameservers string) (bool, error)
	DeleteRegisteredNameserver(nameserver string) (bool, error)
	GetAuthenticationCode() (string, error)
	GetDomainInformation() (datatypes.Container_Dns_Domain_Registration_Information, error)
	GetDomainNameservers() (datatypes.Container_Dns_Domain_Registration_Nameserver, error)
	GetExtendedAttributes(domainName string) (datatypes.Container_Dns_Domain_Registration_ExtendedAttribute, error)
	GetObject() (datatypes.Dns_Domain_Registration, error)
	GetRegisteredNameserver() (datatypes.Container_Dns_Domain_Registration_Nameserver, error)
	GetRegistrantVerificationStatusDetail() (datatypes.Container_Dns_Domain_Registration_Registrant_Verification_StatusDetail, error)
	GetTransferInformation(domainName string) (datatypes.Container_Dns_Domain_Registration_Transfer_Information, error)
	LockDomain() (bool, error)
	LookupDomain(domainName string) (datatypes.Container_Dns_Domain_Registration_Lookup, error)
	ModifyContact(contact datatypes.Container_Dns_Domain_Registration_Contact) (bool, error)
	ModifyRegisteredNameserver(oldNameserver string, newNameserver string, ipAddress string) (bool, error)
	RegisterNameserver(nameserver string, ipAddress string) (bool, error)
	RemoveNameserversFromDomain(nameservers string) (bool, error)
	SendAuthenticationCode() (bool, error)
	SendRegistrantVerificationEmail() (bool, error)
	SendTransferApprovalEmail() (bool, error)
	SetAuthenticationCode(authenticationCode string) (bool, error)
	UnlockDomain() (bool, error)
}

type Dns_Domain_Registration_Registrant_Verification_Status interface {
	Entity

	GetAllObjects() (datatypes.Dns_Domain_Registration_Registrant_Verification_Status, error)
	GetObject() (datatypes.Dns_Domain_Registration_Registrant_Verification_Status, error)
}

type Dns_Domain_Registration_Status interface {
	Entity

	GetAllObjects() (datatypes.Dns_Domain_Registration_Status, error)
	GetObject() (datatypes.Dns_Domain_Registration_Status, error)
}

type Dns_Domain_ResourceRecord interface {
	Entity

	CreateObject(templateObject datatypes.Dns_Domain_ResourceRecord) (datatypes.Dns_Domain_ResourceRecord, error)
	CreateObjects(templateObjects datatypes.Dns_Domain_ResourceRecord) (datatypes.Dns_Domain_ResourceRecord, error)
	DeleteObject() (bool, error)
	DeleteObjects(templateObjects datatypes.Dns_Domain_ResourceRecord) (bool, error)
	EditObject(templateObject datatypes.Dns_Domain_ResourceRecord) (bool, error)
	EditObjects(templateObjects datatypes.Dns_Domain_ResourceRecord) (bool, error)
	GetObject() (datatypes.Dns_Domain_ResourceRecord, error)
}

type Dns_Domain_ResourceRecord_AType interface {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_ResourceRecord_AaaaType interface {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_ResourceRecord_CnameType interface {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_ResourceRecord_MxType interface {
	Dns_Domain_ResourceRecord

	CreateObject(templateObject datatypes.Dns_Domain_ResourceRecord_MxType) (datatypes.Dns_Domain_ResourceRecord_MxType, error)
	CreateObjects(templateObjects datatypes.Dns_Domain_ResourceRecord) (datatypes.Dns_Domain_ResourceRecord, error)
	DeleteObject() (bool, error)
	DeleteObjects(templateObjects datatypes.Dns_Domain_ResourceRecord_MxType) (bool, error)
	EditObject(templateObject datatypes.Dns_Domain_ResourceRecord_MxType) (bool, error)
	EditObjects(templateObjects datatypes.Dns_Domain_ResourceRecord_MxType) (bool, error)
	GetObject() (datatypes.Dns_Domain_ResourceRecord_MxType, error)
}

type Dns_Domain_ResourceRecord_NsType interface {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_ResourceRecord_PtrType interface {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_ResourceRecord_SoaType interface {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_ResourceRecord_SpfType interface {
	Dns_Domain_ResourceRecord_TxtType
}

type Dns_Domain_ResourceRecord_SrvType interface {
	Dns_Domain_ResourceRecord

	CreateObject(templateObject datatypes.Dns_Domain_ResourceRecord_SrvType) (datatypes.Dns_Domain_ResourceRecord_SrvType, error)
	CreateObjects(templateObjects datatypes.Dns_Domain_ResourceRecord) (datatypes.Dns_Domain_ResourceRecord, error)
	DeleteObject() (bool, error)
	DeleteObjects(templateObjects datatypes.Dns_Domain_ResourceRecord_SrvType) (bool, error)
	EditObject(templateObject datatypes.Dns_Domain_ResourceRecord_SrvType) (bool, error)
	EditObjects(templateObjects datatypes.Dns_Domain_ResourceRecord_SrvType) (bool, error)
	GetObject() (datatypes.Dns_Domain_ResourceRecord_SrvType, error)
}

type Dns_Domain_ResourceRecord_TxtType interface {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_Reverse interface {
	Dns_Domain
}

type Dns_Domain_Reverse_Version4 interface {
	Dns_Domain_Reverse
}

type Dns_Domain_Reverse_Version6 interface {
	Dns_Domain_Reverse
}

type Dns_Message interface {
	Entity
}

type Dns_Secondary interface {
	Entity

	ConvertToPrimary() (bool, error)
	CreateObject(templateObject datatypes.Dns_Secondary) (datatypes.Dns_Secondary, error)
	CreateObjects(templateObjects datatypes.Dns_Secondary) (datatypes.Dns_Secondary, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Dns_Secondary) (bool, error)
	GetByDomainName(name string) (datatypes.Dns_Secondary, error)
	GetObject() (datatypes.Dns_Secondary, error)
	TransferNow() (bool, error)
}

type Dns_Status interface {
	Entity
}

type Entity interface {
}

type Event_Log interface {
	Entity

	GetAllEventNames(objectName string) (string, error)
	GetAllEventObjectNames() (string, error)
	GetAllObjects() (datatypes.Event_Log, error)
	GetAllUserTypes() (string, error)
}

type FlexibleCredit_Affiliate interface {
	Entity
}

type FlexibleCredit_Company_Type interface {
	Entity
}

type FlexibleCredit_Enrollment interface {
	Entity
}

type FlexibleCredit_Program interface {
	Entity

	GetAffiliatesAvailableForSelfEnrollmentByVerificationType(verificationTypeKeyName string) (datatypes.FlexibleCredit_Affiliate, error)
	GetCompanyTypes() (datatypes.FlexibleCredit_Company_Type, error)
	GetObject() (datatypes.FlexibleCredit_Program, error)
	SelfEnrollNewAccount(accountTemplate datatypes.Account) (datatypes.Account, error)
}

type Hardware interface {
	Entity

	AllowAccessToNetworkStorage(networkStorageTemplateObject datatypes.Network_Storage) (bool, error)
	AllowAccessToNetworkStorageList(networkStorageTemplateObjects datatypes.Network_Storage) (bool, error)
	CaptureImage(captureTemplate datatypes.Container_Disk_Image_Capture_Template) (datatypes.Virtual_Guest_Block_Device_Template_Group, error)
	CloseAlarm(alarmId string) (bool, error)
	CreateObject(templateObject datatypes.Hardware) (datatypes.Hardware, error)
	DeleteObject() (bool, error)
	DeleteSoftwareComponentPasswords(softwareComponentPasswords datatypes.Software_Component_Password) (bool, error)
	EditSoftwareComponentPasswords(softwareComponentPasswords datatypes.Software_Component_Password) (bool, error)
	ExecuteRemoteScript(uri string) error
	FindByIpAddress(ipAddress string) (datatypes.Hardware, error)
	GenerateOrderTemplate(templateObject datatypes.Hardware) (datatypes.Container_Product_Order, error)
	GetAlarmHistory(startDate time.Time, endDate time.Time, alarmId string) (datatypes.Container_Monitoring_Alarm_History, error)
	GetAttachedNetworkStorages(nasType string) (datatypes.Network_Storage, error)
	GetAvailableNetworkStorages(nasType string) (datatypes.Network_Storage, error)
	GetBackendIncomingBandwidth(startDate time.Time, endDate time.Time) (float64, error)
	GetBackendOutgoingBandwidth(startDate time.Time, endDate time.Time) (float64, error)
	GetCreateObjectOptions() (datatypes.Container_Hardware_Configuration, error)
	GetCurrentBillingDetail() (datatypes.Billing_Item, error)
	GetCurrentBillingTotal() (float64, error)
	GetDailyAverage(startDate time.Time, endDate time.Time) (float64, error)
	GetFrontendIncomingBandwidth(startDate time.Time, endDate time.Time) (float64, error)
	GetFrontendOutgoingBandwidth(startDate time.Time, endDate time.Time) (float64, error)
	GetHourlyBandwidth(mode string, day time.Time) (datatypes.Metric_Tracking_Object_Data, error)
	GetMonitoringActiveAlarms(startDate time.Time, endDate time.Time) (datatypes.Container_Monitoring_Alarm_History, error)
	GetMonitoringClosedAlarms(startDate time.Time, endDate time.Time) (datatypes.Container_Monitoring_Alarm_History, error)
	GetObject() (datatypes.Hardware, error)
	GetPrivateBandwidthData(startTime int, endTime int) (datatypes.Metric_Tracking_Object_Data, error)
	GetPublicBandwidthData(startTime int, endTime int) (datatypes.Metric_Tracking_Object_Data, error)
	GetSensorData() (datatypes.Container_RemoteManagement_SensorReading, error)
	GetSensorDataWithGraphs() (datatypes.Container_RemoteManagement_SensorReadingsWithGraphs, error)
	GetServerFanSpeedGraphs() (datatypes.Container_RemoteManagement_Graphs_SensorSpeed, error)
	GetServerPowerState() (string, error)
	GetServerTemperatureGraphs() (datatypes.Container_RemoteManagement_Graphs_SensorTemperature, error)
	GetTransactionHistory() (datatypes.Provisioning_Version1_Transaction_History, error)
	GetUpgradeItemPrices() (datatypes.Product_Item_Price, error)
	ImportVirtualHost() (datatypes.Virtual_Host, error)
	IsPingable() (bool, error)
	Ping() (string, error)
	PowerCycle() (bool, error)
	PowerOff() (bool, error)
	PowerOn() (bool, error)
	RebootDefault() (bool, error)
	RebootHard() (bool, error)
	RebootSoft() (bool, error)
	RemoveAccessToNetworkStorage(networkStorageTemplateObject datatypes.Network_Storage) (bool, error)
	RemoveAccessToNetworkStorageList(networkStorageTemplateObjects datatypes.Network_Storage) (bool, error)
	SetTags(tags string) (bool, error)
}

type Hardware_Attribute interface {
	Entity
}

type Hardware_Attribute_Type interface {
	Entity
}

type Hardware_Benchmark_Certification interface {
	Entity

	GetObject() (datatypes.Hardware_Benchmark_Certification, error)
	GetResultFile() ([]byte, error)
}

type Hardware_Chassis interface {
	Entity
}

type Hardware_Component interface {
	Entity
}

type Hardware_Component_Attribute interface {
	Entity
}

type Hardware_Component_Attribute_Type interface {
	Entity
}

type Hardware_Component_DriveController interface {
	Hardware_Component
}

type Hardware_Component_HardDrive interface {
	Hardware_Component
}

type Hardware_Component_Model interface {
	Entity

	GetObject() (datatypes.Hardware_Component_Model, error)
}

type Hardware_Component_Model_Architecture_Type interface {
	Entity
}

type Hardware_Component_Model_Attribute interface {
	Entity
}

type Hardware_Component_Model_Attribute_Type interface {
	Entity
}

type Hardware_Component_Model_Generic interface {
	Entity
}

type Hardware_Component_Model_Generic_Attribute interface {
	Entity
}

type Hardware_Component_Model_Generic_MarketingFeature interface {
	Entity
}

type Hardware_Component_Motherboard interface {
	Hardware_Component
}

type Hardware_Component_Motherboard_Reboot_Time interface {
	Entity
}

type Hardware_Component_NetworkCard interface {
	Hardware_Component
}

type Hardware_Component_Partition interface {
	Entity
}

type Hardware_Component_Partition_OperatingSystem interface {
	Entity

	GetAllObjects() (datatypes.Hardware_Component_Partition_OperatingSystem, error)
	GetByDescription(description string) (datatypes.Hardware_Component_Partition_OperatingSystem, error)
	GetObject() (datatypes.Hardware_Component_Partition_OperatingSystem, error)
}

type Hardware_Component_Partition_Template interface {
	Entity

	GetObject() (datatypes.Hardware_Component_Partition_Template, error)
}

type Hardware_Component_Partition_Template_Partition interface {
	Entity
}

type Hardware_Component_Processor interface {
	Hardware_Component
}

type Hardware_Component_Ram interface {
	Hardware_Component
}

type Hardware_Component_RemoteManagement interface {
	Hardware_Component
}

type Hardware_Component_RemoteManagement_Command interface {
	Entity
}

type Hardware_Component_RemoteManagement_Command_Request interface {
	Entity
}

type Hardware_Component_RemoteManagement_User interface {
	Entity
}

type Hardware_Component_SecurityDevice interface {
	Hardware_Component
}

type Hardware_Component_SecurityDevice_Infineon interface {
	Hardware_Component_SecurityDevice
}

type Hardware_Component_Type interface {
	Entity
}

type Hardware_Firewall interface {
	Hardware_Switch
}

type Hardware_Function interface {
	Entity
}

type Hardware_Group interface {
	Entity
}

type Hardware_LoadBalancer interface {
	Hardware
}

type Hardware_Note interface {
	Entity
}

type Hardware_Note_Type interface {
	Entity
}

type Hardware_Power_Component interface {
	Entity
}

type Hardware_Router interface {
	Hardware_Switch

	GetObject() (datatypes.Hardware_Router, error)
}

type Hardware_Router_Backend interface {
	Hardware_Router
}

type Hardware_Router_Frontend interface {
	Hardware_Router
}

type Hardware_SecurityModule interface {
	Hardware_Server

	CreateObject(templateObject datatypes.Hardware_SecurityModule) (datatypes.Hardware_SecurityModule, error)
	GetObject() (datatypes.Hardware_SecurityModule, error)
}

type Hardware_Server interface {
	Hardware

	ActivatePrivatePort() (bool, error)
	ActivatePublicPort() (bool, error)
	BootToRescueLayer(noOsBootEnvironment string) (bool, error)
	CreateFirmwareUpdateTransaction(ipmi int, raidController int, bios int, harddrive int) (bool, error)
	CreateObject(templateObject datatypes.Hardware_Server) (datatypes.Hardware_Server, error)
	CreatePostSoftwareInstallTransaction(installCodes string, returnBoolean bool) (bool, error)
	EditObject(templateObject datatypes.Hardware_Server) (bool, error)
	GetBackendBandwidthUsage(startDate time.Time, endDate time.Time) (datatypes.Metric_Tracking_Object_Data, error)
	GetBackendBandwidthUse(startDate time.Time, endDate time.Time) (datatypes.Network_Bandwidth_Version1_Usage_Detail, error)
	GetBandwidthForDateRange(startDate time.Time, endDate time.Time) (datatypes.Metric_Tracking_Object_Data, error)
	GetBandwidthImage(networkType string, snapshotRange string, draw bool, dateSpecified time.Time, dateSpecifiedEnd time.Time) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetCurrentBenchmarkCertificationResultFile() ([]byte, error)
	GetCustomBandwidthDataByDate(graphData datatypes.Container_Graph) (datatypes.Container_Graph, error)
	GetFirewallProtectableSubnets() (datatypes.Network_Subnet, error)
	GetFrontendBandwidthUsage(startDate time.Time, endDate time.Time) (datatypes.Metric_Tracking_Object_Data, error)
	GetFrontendBandwidthUse(startDate time.Time, endDate time.Time) (datatypes.Network_Bandwidth_Version1_Usage_Detail, error)
	GetHardwareByIpAddress(ipAddress string) (datatypes.Hardware_Server, error)
	GetItemPricesFromSoftwareDescriptions(softwareDescriptions datatypes.Software_Description, includeTranslationsFlag bool, returnAllPricesFlag bool) (datatypes.Product_Item, error)
	GetManagementNetworkComponent() (datatypes.Network_Component, error)
	GetNetworkComponentFirewallProtectableIpAddresses() (datatypes.Network_Subnet_IpAddress, error)
	GetObject() (datatypes.Hardware_Server, error)
	GetPMInfo() (datatypes.Container_RemoteManagement_PmInfo, error)
	GetPrimaryDriveSize() (int, error)
	GetPrivateBandwidthDataSummary() (datatypes.Container_Network_Bandwidth_Data_Summary, error)
	GetPrivateBandwidthGraphImage(startTime string, endTime string) ([]byte, error)
	GetPrivateNetworkComponent() (datatypes.Network_Component, error)
	GetPrivateVlan() (datatypes.Network_Vlan, error)
	GetPrivateVlanByIpAddress(ipAddress string) (datatypes.Network_Vlan, error)
	GetProvisionDate() (time.Time, error)
	GetPublicBandwidthDataSummary() (datatypes.Container_Network_Bandwidth_Data_Summary, error)
	GetPublicBandwidthGraphImage(startTime time.Time, endTime time.Time) ([]byte, error)
	GetPublicBandwidthTotal(startTime int, endTime int) (uint, error)
	GetPublicNetworkComponent() (datatypes.Network_Component, error)
	GetPublicVlan() (datatypes.Network_Vlan, error)
	GetPublicVlanByHostname(hostname string) (datatypes.Network_Vlan, error)
	GetReverseDomainRecords() (datatypes.Dns_Domain, error)
	GetSensorData() (datatypes.Container_RemoteManagement_SensorReading, error)
	GetSensorDataWithGraphs() (datatypes.Container_RemoteManagement_SensorReadingsWithGraphs, error)
	GetServerDetails() (datatypes.Container_Hardware_Server_Details, error)
	GetServerFanSpeedGraphs() (datatypes.Container_RemoteManagement_Graphs_SensorSpeed, error)
	GetServerPowerState() (string, error)
	GetServerTemperatureGraphs() (datatypes.Container_RemoteManagement_Graphs_SensorTemperature, error)
	GetValidBlockDeviceTemplateGroups(visibility string) (datatypes.Virtual_Guest_Block_Device_Template_Group, error)
	GetWindowsUpdateAvailableUpdates() (datatypes.Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem, error)
	GetWindowsUpdateInstalledUpdates() (datatypes.Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem, error)
	GetWindowsUpdateStatus() (datatypes.Container_Utility_Microsoft_Windows_UpdateServices_Status, error)
	InitiateIderaBareMetalRestore() (bool, error)
	InitiateR1SoftBareMetalRestore() (bool, error)
	IsBackendPingable() (bool, error)
	IsPingable() (bool, error)
	IsWindowsServer() (bool, error)
	Ping() (string, error)
	PowerCycle() (bool, error)
	PowerOff() (bool, error)
	PowerOn() (bool, error)
	RebootDefault() (bool, error)
	RebootHard() (bool, error)
	RebootSoft() (bool, error)
	ReloadCurrentOperatingSystemConfiguration(token string) (string, error)
	ReloadOperatingSystem(token string, config datatypes.Container_Hardware_Server_Configuration) (string, error)
	RunPassmarkCertificationBenchmark() (bool, error)
	SetOperatingSystemPassword(newPassword string) (bool, error)
	SetPrivateNetworkInterfaceSpeed(newSpeed int) (bool, error)
	SetPublicNetworkInterfaceSpeed(newSpeed int) (bool, error)
	SetUserMetadata(metadata string) (datatypes.Hardware_Attribute, error)
	ShutdownPrivatePort() (bool, error)
	ShutdownPublicPort() (bool, error)
	SparePool(action string, newOrder bool) (bool, error)
	ValidatePartitionsForOperatingSystem(operatingSystem datatypes.Software_Description, partitions datatypes.Hardware_Component_Partition) (bool, error)
}

type Hardware_Status interface {
	Entity
}

type Hardware_Switch interface {
	Hardware
}

type Layout_Container interface {
	Entity

	GetAllObjects() (datatypes.Layout_Container, error)
	GetObject() (datatypes.Layout_Container, error)
}

type Layout_Container_Type interface {
	Entity
}

type Layout_Item interface {
	Entity

	GetObject() (datatypes.Layout_Item, error)
}

type Layout_Item_Type interface {
	Entity
}

type Layout_Preference interface {
	Entity
}

type Layout_Preference_Type interface {
	Entity
}

type Layout_Profile interface {
	Entity

	CreateObject(templateObject datatypes.Layout_Profile) (bool, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Layout_Profile) (bool, error)
	GetObject() (datatypes.Layout_Profile, error)
	ModifyPreference(templateObject datatypes.Layout_Profile_Preference) (datatypes.Layout_Profile_Preference, error)
	ModifyPreferences(layoutPreferenceObjects datatypes.Layout_Profile_Preference) (datatypes.Layout_Profile_Preference, error)
}

type Layout_Profile_Containers interface {
	Entity

	CreateObject(templateObject datatypes.Layout_Profile_Containers) (bool, error)
	EditObject(templateObject datatypes.Layout_Profile_Containers) (bool, error)
	GetObject() (datatypes.Layout_Profile_Containers, error)
}

type Layout_Profile_Customer interface {
	Layout_Profile

	GetObject() (datatypes.Layout_Profile_Customer, error)
}

type Layout_Profile_Preference interface {
	Entity

	GetObject() (datatypes.Layout_Profile_Preference, error)
}

type Legal_RegulatedWorkload interface {
	Entity
}

type Legal_RegulatedWorkload_Type interface {
	Entity
}

type Locale interface {
	Entity

	GetClosestToLanguageTag(languageTag string) (datatypes.Locale, error)
	GetObject() (datatypes.Locale, error)
}

type Locale_Country interface {
	Entity

	GetAvailableCountries() (datatypes.Locale_Country, error)
	GetCountries() (datatypes.Locale_Country, error)
	GetObject() (datatypes.Locale_Country, error)
}

type Locale_StateProvince interface {
	Entity
}

type Locale_Timezone interface {
	Entity

	GetAllObjects() (datatypes.Locale_Timezone, error)
	GetObject() (datatypes.Locale_Timezone, error)
}

type Location interface {
	Entity

	GetAvailableObjectStorageDatacenters() (datatypes.Location, error)
	GetDatacenters() (datatypes.Location, error)
	GetDatacentersWithVirtualImageStoreServiceResourceRecord() (datatypes.Location, error)
	GetObject() (datatypes.Location, error)
	GetViewableDatacenters() (datatypes.Location, error)
	GetViewablePopsAndDataCenters() (datatypes.Location, error)
	GetViewablepointOfPresence() (datatypes.Location, error)
	GetpointOfPresence() (datatypes.Location, error)
}

type Location_Datacenter interface {
	Location

	GetObject() (datatypes.Location_Datacenter, error)
	GetStatisticsGraphImage() ([]byte, error)
}

type Location_Group interface {
	Entity

	GetAllObjects() (datatypes.Location_Group, error)
	GetObject() (datatypes.Location_Group, error)
}

type Location_Group_Location_CrossReference interface {
	Entity
}

type Location_Group_Pricing interface {
	Location_Group

	GetAllObjects() (datatypes.Location_Group, error)
	GetObject() (datatypes.Location_Group_Pricing, error)
}

type Location_Group_Regional interface {
	Location_Group

	GetAllObjects() (datatypes.Location_Group, error)
	GetObject() (datatypes.Location_Group_Regional, error)
}

type Location_Group_Type interface {
	Entity
}

type Location_Inventory_Room interface {
	Location
}

type Location_Network_Operations_Center interface {
	Location
}

type Location_Office interface {
	Location
}

type Location_Rack interface {
	Location
}

type Location_Region interface {
	Entity
}

type Location_Region_Location interface {
	Entity
}

type Location_Reservation interface {
	Entity

	GetAccountReservations() (datatypes.Location_Reservation, error)
	GetObject() (datatypes.Location_Reservation, error)
}

type Location_Reservation_Rack interface {
	Entity

	GetObject() (datatypes.Location_Reservation_Rack, error)
}

type Location_Reservation_Rack_Member interface {
	Entity

	GetObject() (datatypes.Location_Reservation_Rack_Member, error)
}

type Location_Root interface {
	Location
}

type Location_Server_Room interface {
	Location
}

type Location_Slot interface {
	Location
}

type Location_Status interface {
	Entity
}

type Location_Storage_Room interface {
	Location
}

type Marketplace_EmailDistribution interface {
	Entity
}

type Marketplace_Partner interface {
	Entity

	GetAllObjects() (datatypes.Marketplace_Partner, error)
	GetAllPublishedPartners(searchTerm string) (datatypes.Marketplace_Partner, error)
	GetFeaturedPartners(non bool) (datatypes.Marketplace_Partner, error)
	GetFile(name string) (datatypes.Marketplace_Partner_File, error)
	GetObject() (datatypes.Marketplace_Partner, error)
	GetPartnerByUrlIdentifier(urlIdentifier string) (datatypes.Marketplace_Partner, error)
}

type Marketplace_Partner_Attachment interface {
	Entity
}

type Marketplace_Partner_Attachment_Type interface {
	Entity
}

type Marketplace_Partner_File interface {
	Entity
}

type Marketplace_Partner_File_Attributes interface {
	Entity
}

type Metric_Tracking_Object interface {
	Entity

	GetBackboneBandwidthGraph(graphTitle string) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetBandwidthData(startDateTime time.Time, endDateTime time.Time, typ string, rollupSeconds int) (datatypes.Metric_Tracking_Object_Data, error)
	GetBandwidthGraph(startDateTime time.Time, endDateTime time.Time, graphType string, fontSize int, graphWidth int, graphHeight int, doNotShowTimeZone bool) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetBandwidthTotal(startDateTime time.Time, endDateTime time.Time, direction string, typ string) (uint, error)
	GetCustomGraphData(graphContainer datatypes.Container_Graph) (datatypes.Container_Graph, error)
	GetDetailsForDateRange(startDate time.Time, endDate time.Time, graphType string) (datatypes.Container_Metric_Tracking_Object_Details, error)
	GetGraph(startDateTime time.Time, endDateTime time.Time, graphType string) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetMetricDataTypes() (datatypes.Container_Metric_Data_Type, error)
	GetObject() (datatypes.Metric_Tracking_Object, error)
	GetSummary(graphType string) (datatypes.Container_Metric_Tracking_Object_Summary, error)
	GetSummaryData(startDateTime time.Time, endDateTime time.Time, validTypes datatypes.Container_Metric_Data_Type, summaryPeriod int) (datatypes.Metric_Tracking_Object_Data, error)
}

type Metric_Tracking_Object_Abstract interface {
	Metric_Tracking_Object
}

type Metric_Tracking_Object_Bandwidth_Summary interface {
	Entity

	GetObject() (datatypes.Metric_Tracking_Object_Bandwidth_Summary, error)
}

type Metric_Tracking_Object_Data interface {
	Entity
}

type Metric_Tracking_Object_Data_Network_ContentDelivery_Account interface {
	Metric_Tracking_Object_Data
}

type Metric_Tracking_Object_HardwareServer interface {
	Metric_Tracking_Object_Abstract
}

type Metric_Tracking_Object_Type interface {
	Entity
}

type Metric_Tracking_Object_VirtualDedicatedRack interface {
	Metric_Tracking_Object_Abstract
}

type Metric_Tracking_Object_Virtual_Storage_Repository interface {
	Metric_Tracking_Object_Abstract
}

type Monitoring_Agent interface {
	Entity

	Activate() (bool, error)
	AddConfigurationProfile(configurationValues datatypes.Monitoring_Agent_Configuration_Value) (datatypes.Provisioning_Version1_Transaction, error)
	ApplyConfigurationValues(configurationValues datatypes.Monitoring_Agent_Configuration_Value) (datatypes.Provisioning_Version1_Transaction, error)
	Deactivate() (bool, error)
	DeleteConfigurationProfile(sectionId int, profileId int) (bool, error)
	DeployMonitoringAgent(configurationTemplateId int) (datatypes.Provisioning_Version1_Transaction, error)
	GetActiveAlarmSubscribers() (datatypes.Notification_User_Subscriber, error)
	GetAvailableConfigurationTemplates() (datatypes.Configuration_Template, error)
	GetAvailableConfigurationValues(configurationDefinitionId int, configValues datatypes.Monitoring_Agent_Configuration_Value) (datatypes.Monitoring_Agent_Configuration_Value, error)
	GetEligibleAlarmSubscibers() (datatypes.User_Customer, error)
	GetGraph(configurationValues datatypes.Monitoring_Agent_Configuration_Value, beginDate time.Time, endDate time.Time) (datatypes.Container_Monitoring_Graph_Outputs, error)
	GetGraphData(metricDataTypes datatypes.Container_Metric_Data_Type, startDate time.Time, endDate time.Time) (datatypes.Metric_Tracking_Object_Data, error)
	GetObject() (datatypes.Monitoring_Agent, error)
	RemoveActiveAlarmSubscriber(userRecordId int) (bool, error)
	RemoveAllAlarmSubscribers() (bool, error)
	RestartMonitoringAgent() (bool, error)
	SetActiveAlarmSubscriber(userRecordId int) (bool, error)
}

type Monitoring_Agent_Configuration_Template_Group interface {
	Entity

	CreateObject(templateObject datatypes.Monitoring_Agent_Configuration_Template_Group) (datatypes.Monitoring_Agent_Configuration_Template_Group, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Monitoring_Agent_Configuration_Template_Group) (bool, error)
	GetAllObjects() (datatypes.Monitoring_Agent_Configuration_Template_Group, error)
	GetConfigurationGroups(packageId int) (datatypes.Monitoring_Agent_Configuration_Template_Group, error)
	GetObject() (datatypes.Monitoring_Agent_Configuration_Template_Group, error)
}

type Monitoring_Agent_Configuration_Template_Group_Reference interface {
	Entity

	CreateObject(templateObject datatypes.Monitoring_Agent_Configuration_Template_Group_Reference) (datatypes.Monitoring_Agent_Configuration_Template_Group_Reference, error)
	CreateObjects(templateObjects datatypes.Monitoring_Agent_Configuration_Template_Group_Reference) (bool, error)
	EditObject(templateObject datatypes.Monitoring_Agent_Configuration_Template_Group_Reference) (bool, error)
	EditObjects(templateObjects datatypes.Monitoring_Agent_Configuration_Template_Group_Reference) (bool, error)
	GetAllObjects() (datatypes.Monitoring_Agent_Configuration_Template_Group_Reference, error)
	GetObject() (datatypes.Monitoring_Agent_Configuration_Template_Group_Reference, error)
}

type Monitoring_Agent_Configuration_Value interface {
	Entity

	GetObject() (datatypes.Monitoring_Agent_Configuration_Value, error)
}

type Monitoring_Agent_Status interface {
	Entity

	GetObject() (datatypes.Monitoring_Agent_Status, error)
}

type Monitoring_Robot interface {
	Entity

	CheckConnection() (bool, error)
	DeployMonitoringAgents(configurationTemplateGroup datatypes.Monitoring_Agent_Configuration_Template_Group) (datatypes.Provisioning_Version1_Transaction, error)
	GetAvailableConfigurationGroups() (datatypes.Monitoring_Agent_Configuration_Template_Group, error)
	GetObject() (datatypes.Monitoring_Robot, error)
	ResetStatus() (bool, error)
}

type Monitoring_Robot_Status interface {
	Entity
}

type Network interface {
	Entity

	CreateObject(templateObject datatypes.Network) (datatypes.Network, error)
	CreateSubnet(subnet datatypes.Network_Subnet, pod datatypes.Network_Pod) (datatypes.Network_Subnet, error)
	DeleteObject() (bool, error)
	DeleteSubnet(subnet datatypes.Network_Subnet) (bool, error)
	EditObject(templateObject datatypes.Network) (bool, error)
	GetAllObjects() (datatypes.Network, error)
	GetObject() (datatypes.Network, error)
}

type Network_Application_Delivery_Controller interface {
	Entity

	CreateLiveLoadBalancer(loadBalancer datatypes.Network_LoadBalancer_VirtualIpAddress) (bool, error)
	DeleteLiveLoadBalancer(loadBalancer datatypes.Network_LoadBalancer_VirtualIpAddress) (bool, error)
	DeleteLiveLoadBalancerService(service datatypes.Network_LoadBalancer_Service) (bool, error)
	EditObject(templateObject datatypes.Network_Application_Delivery_Controller) (bool, error)
	GetBandwidthDataByDate(startDateTime time.Time, endDateTime time.Time, networkType string) (datatypes.Metric_Tracking_Object_Data, error)
	GetBandwidthImageByDate(startDateTime time.Time, endDateTime time.Time, networkType string) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetCustomBandwidthDataByDate(graphData datatypes.Container_Graph) (datatypes.Container_Graph, error)
	GetLiveLoadBalancerServiceGraphImage(service datatypes.Network_LoadBalancer_Service, graphType string, metric string) ([]byte, error)
	GetObject() (datatypes.Network_Application_Delivery_Controller, error)
	RestoreBaseConfiguration() (bool, error)
	RestoreConfiguration(configurationHistoryId int) (bool, error)
	SaveCurrentConfiguration(notes string) (datatypes.Network_Application_Delivery_Controller_Configuration_History, error)
	UpdateLiveLoadBalancer(loadBalancer datatypes.Network_LoadBalancer_VirtualIpAddress) (bool, error)
	UpdateNetScalerLicense() (datatypes.Provisioning_Version1_Transaction, error)
}

type Network_Application_Delivery_Controller_Configuration_History interface {
	Entity

	DeleteObject() (bool, error)
	GetObject() (datatypes.Network_Application_Delivery_Controller_Configuration_History, error)
}

type Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute interface {
	Entity

	GetObject() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute, error)
}

type Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type interface {
	Entity

	GetAllObjects() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type, error)
	GetObject() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type, error)
}

type Network_Application_Delivery_Controller_LoadBalancer_Health_Check interface {
	Entity

	GetObject() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Check, error)
}

type Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type interface {
	Entity

	GetAllObjects() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type, error)
	GetObject() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type, error)
}

type Network_Application_Delivery_Controller_LoadBalancer_Routing_Method interface {
	Entity

	GetAllObjects() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Routing_Method, error)
	GetObject() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Routing_Method, error)
}

type Network_Application_Delivery_Controller_LoadBalancer_Routing_Type interface {
	Entity

	GetAllObjects() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Routing_Type, error)
	GetObject() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Routing_Type, error)
}

type Network_Application_Delivery_Controller_LoadBalancer_Service interface {
	Entity

	DeleteObject() (bool, error)
	GetGraphImage(graphType string, metric string) ([]byte, error)
	GetObject() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Service, error)
	ToggleStatus() (bool, error)
}

type Network_Application_Delivery_Controller_LoadBalancer_Service_Group interface {
	Entity

	GetGraphImage(graphType string, metric string) ([]byte, error)
	GetObject() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_Service_Group, error)
	KickAllConnections() (bool, error)
}

type Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference interface {
	Entity
}

type Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress interface {
	Entity

	EditObject(templateObject datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) (bool, error)
	GetAvailableSecureTransportCiphers() (datatypes.Security_SecureTransportCipher, error)
	GetAvailableSecureTransportProtocols() (datatypes.Security_SecureTransportProtocol, error)
	GetObject() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress, error)
	StartSsl() (bool, error)
	StopSsl() (bool, error)
	UpgradeConnectionLimit() (bool, error)
}

type Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportCipher interface {
	Entity
}

type Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportProtocol interface {
	Entity
}

type Network_Application_Delivery_Controller_LoadBalancer_VirtualServer interface {
	Entity

	DeleteObject() (bool, error)
	GetObject() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualServer, error)
	StartSsl() (bool, error)
	StopSsl() (bool, error)
}

type Network_Application_Delivery_Controller_Type interface {
	Entity
}

type Network_Backbone interface {
	Entity

	GetAllBackbones() (datatypes.Network_Backbone, error)
	GetBackbonesForLocationName(locationName string) (datatypes.Network_Backbone, error)
	GetGraphImage() ([]byte, error)
	GetObject() (datatypes.Network_Backbone, error)
}

type Network_Backbone_Location_Dependent interface {
	Entity

	GetAllObjects() (datatypes.Network_Backbone_Location_Dependent, error)
	GetObject() (datatypes.Network_Backbone_Location_Dependent, error)
	GetSourceDependentsByName(locationName string) (datatypes.Location, error)
}

type Network_Bandwidth_Usage interface {
	Entity
}

type Network_Bandwidth_Usage_Detail interface {
	Entity
}

type Network_Bandwidth_Version1_Allocation interface {
	Entity
}

type Network_Bandwidth_Version1_Allotment interface {
	Entity

	CreateObject(templateObject datatypes.Network_Bandwidth_Version1_Allotment) (datatypes.Network_Bandwidth_Version1_Allotment, error)
	EditObject(templateObject datatypes.Network_Bandwidth_Version1_Allotment) (bool, error)
	GetBackendBandwidthByHour(date time.Time) (datatypes.Container_Network_Bandwidth_Version1_Usage, error)
	GetBackendBandwidthUse(startDate time.Time, endDate time.Time) (datatypes.Network_Bandwidth_Version1_Usage_Detail, error)
	GetBandwidthForDateRange(startDate time.Time, endDate time.Time) (datatypes.Metric_Tracking_Object_Data, error)
	GetBandwidthImage(networkType string, snapshotRange string, draw bool, dateSpecified time.Time, dateSpecifiedEnd time.Time) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetCustomBandwidthDataByDate(graphData datatypes.Container_Graph) (datatypes.Container_Graph, error)
	GetFrontendBandwidthByHour(date time.Time) (datatypes.Container_Network_Bandwidth_Version1_Usage, error)
	GetFrontendBandwidthUse(startDate time.Time, endDate time.Time) (datatypes.Network_Bandwidth_Version1_Usage_Detail, error)
	GetObject() (datatypes.Network_Bandwidth_Version1_Allotment, error)
	ReassignServers(templateObjects datatypes.Hardware, newAllotmentId int) (bool, error)
	RequestVdrCancellation() (bool, error)
	RequestVdrContentUpdates(hardwareToAdd datatypes.Hardware, hardwareToRemove datatypes.Hardware, cloudsToAdd datatypes.Virtual_Guest, cloudsToRemove datatypes.Virtual_Guest, optionalAllotmentId int, adcToAdd datatypes.Network_Application_Delivery_Controller, adcToRemove datatypes.Network_Application_Delivery_Controller) (bool, error)
	SetVdrContent(hardware datatypes.Hardware, bareMetalServers datatypes.Hardware, virtualServerInstance datatypes.Virtual_Guest, adc datatypes.Network_Application_Delivery_Controller, optionalAllotmentId int) (bool, error)
	UnassignServers(templateObjects datatypes.Hardware) (bool, error)
	VoidPendingServerMove(id int, typ string) (bool, error)
	VoidPendingVdrCancellation() (bool, error)
}

type Network_Bandwidth_Version1_Allotment_Detail interface {
	Entity
}

type Network_Bandwidth_Version1_Host interface {
	Entity
}

type Network_Bandwidth_Version1_Interface interface {
	Entity
}

type Network_Bandwidth_Version1_Usage interface {
	Entity
}

type Network_Bandwidth_Version1_Usage_Detail interface {
	Entity
}

type Network_Bandwidth_Version1_Usage_Detail_Total interface {
	Entity
}

type Network_Bandwidth_Version1_Usage_Detail_Type interface {
	Entity
}

type Network_Component interface {
	Entity

	AddNetworkVlanTrunks(networkVlans datatypes.Network_Vlan) (datatypes.Network_Vlan, error)
	ClearNetworkVlanTrunks() (datatypes.Network_Vlan, error)
	GetCustomBandwidthDataByDate(graphData datatypes.Container_Graph) (datatypes.Container_Graph, error)
	GetObject() (datatypes.Network_Component, error)
	GetPortStatistics() (datatypes.Container_Network_Port_Statistic, error)
	RemoveNetworkVlanTrunks(networkVlans datatypes.Network_Vlan) (datatypes.Network_Vlan, error)
}

type Network_Component_Duplex_Mode interface {
	Entity
}

type Network_Component_Firewall interface {
	Entity

	GetObject() (datatypes.Network_Component_Firewall, error)
}

type Network_Component_Firewall_Rule interface {
	Entity
}

type Network_Component_Firewall_Subnets interface {
	Entity
}

type Network_Component_Group interface {
	Entity
}

type Network_Component_IpAddress interface {
	Entity
}

type Network_Component_Network_Vlan_Trunk interface {
	Entity
}

type Network_Component_RemoteManagement interface {
	Network_Component
}

type Network_Component_Uplink_Hardware interface {
	Entity
}

type Network_ContentDelivery_Account interface {
	Entity

	AuthenticateResourceRequest(parameter datatypes.Container_Network_ContentDelivery_Authentication_Parameter) (bool, error)
	CreateDirectory(directoryName string) (bool, error)
	CreateFtpUser(newPassword string) (bool, error)
	CreateOriginPullMapping(mappingObject datatypes.Container_Network_ContentDelivery_OriginPull_Mapping) (bool, error)
	CreateOriginPullRule(originDomain string, cnameRecord string) (bool, error)
	CreateTokenAuthenticationDirectory(directory string, mediaType string) (bool, error)
	DeleteFtpUser() (bool, error)
	DeleteOriginPullRule(originMappingId string) (bool, error)
	DisableLogging() (bool, error)
	EnableLogging() (bool, error)
	GetAllPopsBandwidthData(beginDateTime time.Time, endDateTime time.Time) (datatypes.Container_Network_ContentDelivery_Bandwidth_PointsOfPresence_Summary, error)
	GetAllPopsBandwidthImage(title string, beginDateTime time.Time, endDateTime time.Time, unit string) (datatypes.Container_Bandwidth_GraphOutputsExtended, error)
	GetAuthenticationServiceEndpoints() (datatypes.Container_Network_ContentDelivery_Authentication_ServiceEndpoint, error)
	GetBandwidthData(beginDateTime time.Time, endDateTime time.Time) (datatypes.Container_Network_ContentDelivery_Bandwidth_Summary, error)
	GetBandwidthDataWithTypes(beginDateTime time.Time, endDateTime time.Time, period string) (datatypes.Container_Network_ContentDelivery_Report_Usage, error)
	GetBandwidthImage(title string, beginDateTime time.Time, endDateTime time.Time) (datatypes.Container_Bandwidth_GraphOutputsExtended, error)
	GetCustomerOrigins(mediaType string) (datatypes.Container_Network_ContentDelivery_OriginPull_Mapping, error)
	GetDirectoryInformation(directoryName string) (datatypes.Container_Network_Directory_Listing, error)
	GetDiskSpaceUsageDataByDate(beginDateTime time.Time, endDateTime time.Time) (datatypes.Metric_Tracking_Object_Data, error)
	GetDiskSpaceUsageImageByDate(beginDateTime time.Time, endDateTime time.Time) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetFtpAttributes() (datatypes.Container_Network_Authentication_Data, error)
	GetMediaUrls() (datatypes.Container_Network_ContentDelivery_SupportedProtocol, error)
	GetObject() (datatypes.Network_ContentDelivery_Account, error)
	GetOriginPullMappingInformation() (datatypes.Container_Network_ContentDelivery_OriginPull_Mapping, error)
	GetOriginPullSupportedMediaUrls() (datatypes.Container_Network_ContentDelivery_SupportedProtocol, error)
	GetOriginPullUrl() (string, error)
	GetPopNames() (datatypes.Container_Network_ContentDelivery_PointsOfPresence, error)
	GetProviderPortalCredentials() (datatypes.Container_Network_Authentication_Data, error)
	GetTokenAuthenticationDirectories() (datatypes.Container_Network_Directory_Listing, error)
	GetVendorFtpAttributes() (datatypes.Container_Network_Authentication_Data, error)
	LoadContent(objectUrls string) (bool, error)
	ManageHttpCompression(enableFlag bool, mimeTypes string) (bool, error)
	PurgeCache(objectUrls string) (datatypes.Container_Network_ContentDelivery_PurgeService_Response, error)
	RemoveAuthenticationDirectory(directory string, mediaType string) (bool, error)
	RemoveFile(source string) (bool, error)
	SetAuthenticationServiceEndpoint(webserviceEndpoint string, protocol string) (bool, error)
	SetFtpPassword(newPassword string) (bool, error)
	UpdateNote(note string) (bool, error)
	UploadStream(source datatypes.Container_Utility_File_Attachment, target string) (bool, error)
}

type Network_ContentDelivery_Account_Status interface {
	Entity
}

type Network_ContentDelivery_Authentication_Address interface {
	Entity

	CreateObject(templateObject datatypes.Network_ContentDelivery_Authentication_Address) (datatypes.Network_ContentDelivery_Authentication_Address, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Network_ContentDelivery_Authentication_Address) (bool, error)
	GetObject() (datatypes.Network_ContentDelivery_Authentication_Address, error)
	RearrangeAuthenticationIp(cdnAccountId int, templateObjects datatypes.Network_ContentDelivery_Authentication_Address) (bool, error)
}

type Network_ContentDelivery_Authentication_Token interface {
	Entity

	CreateObject(templateObject datatypes.Network_ContentDelivery_Authentication_Token) (datatypes.Network_ContentDelivery_Authentication_Token, error)
	GetAllManagedTokens(cdnAccountId int) (datatypes.Network_ContentDelivery_Authentication_Token, error)
	GetObject() (datatypes.Network_ContentDelivery_Authentication_Token, error)
	GetTimedToken(cdnAccountId int, tokenLife int, clientIp string, referrer string, mediaType string) (string, error)
	RevokeAllManagedTokens(cdnAccountId int) (bool, error)
	RevokeAllTokens(cdnAccountId int, mediaType string) (bool, error)
	RevokeManagedToken(cdnAccountId int, token string) (bool, error)
	RevokeManagedTokens(templateObjects datatypes.Network_ContentDelivery_Authentication_Token) (bool, error)
}

type Network_Customer_Subnet interface {
	Entity

	CreateObject(templateObject datatypes.Network_Customer_Subnet) (datatypes.Network_Customer_Subnet, error)
	GetObject() (datatypes.Network_Customer_Subnet, error)
}

type Network_Customer_Subnet_IpAddress interface {
	Entity
}

type Network_Firewall_AccessControlList interface {
	Entity

	GetObject() (datatypes.Network_Firewall_AccessControlList, error)
}

type Network_Firewall_Interface interface {
	Network_Firewall_Module_Context_Interface

	GetObject() (datatypes.Network_Firewall_Interface, error)
}

type Network_Firewall_Module_Context_Interface interface {
	Entity

	GetObject() (datatypes.Network_Firewall_Module_Context_Interface, error)
}

type Network_Firewall_Template interface {
	Entity

	GetAllObjects() (datatypes.Network_Firewall_Template, error)
	GetObject() (datatypes.Network_Firewall_Template, error)
}

type Network_Firewall_Template_Rule interface {
	Entity
}

type Network_Firewall_Update_Request interface {
	Entity

	CreateObject(templateObject datatypes.Network_Firewall_Update_Request) (datatypes.Network_Firewall_Update_Request, error)
	GetFirewallUpdateRequestRuleAttributes() (datatypes.Container_Utility_Network_Firewall_Rule_Attribute, error)
	GetObject() (datatypes.Network_Firewall_Update_Request, error)
	UpdateRuleNote(fwRule datatypes.Network_Component_Firewall_Rule, note string) (bool, error)
}

type Network_Firewall_Update_Request_Customer interface {
	Network_Firewall_Update_Request
}

type Network_Firewall_Update_Request_Employee interface {
	Network_Firewall_Update_Request
}

type Network_Firewall_Update_Request_Rule interface {
	Entity

	CreateObject(templateObject datatypes.Network_Firewall_Update_Request_Rule) (datatypes.Network_Firewall_Update_Request_Rule, error)
	GetObject() (datatypes.Network_Firewall_Update_Request_Rule, error)
	ValidateRule(rule datatypes.Network_Firewall_Update_Request_Rule, applyToComponentId int, applyToAclId int) error
}

type Network_Firewall_Update_Request_Rule_Version6 interface {
	Network_Firewall_Update_Request_Rule
}

type Network_Gateway interface {
	Entity

	BypassAllVlans() error
	BypassVlans(vlans datatypes.Network_Gateway_Vlan) error
	CreateObject(templateObject datatypes.Network_Gateway) (datatypes.Network_Gateway, error)
	EditObject(templateObject datatypes.Network_Gateway) (bool, error)
	GetObject() (datatypes.Network_Gateway, error)
	GetPossibleInsideVlans() (datatypes.Network_Vlan, error)
	UnbypassAllVlans() error
	UnbypassVlans(vlans datatypes.Network_Gateway_Vlan) error
}

type Network_Gateway_Member interface {
	Entity

	CreateObject(templateObject datatypes.Network_Gateway_Member) (datatypes.Network_Gateway_Member, error)
	CreateObjects(templateObjects datatypes.Network_Gateway_Member) (datatypes.Network_Gateway_Member, error)
	GetObject() (datatypes.Network_Gateway_Member, error)
}

type Network_Gateway_Status interface {
	Entity

	GetObject() (datatypes.Network_Gateway_Status, error)
}

type Network_Gateway_Vlan interface {
	Entity

	Bypass() error
	CreateObject(templateObject datatypes.Network_Gateway_Vlan) (datatypes.Network_Gateway_Vlan, error)
	CreateObjects(templateObjects datatypes.Network_Gateway_Vlan) (datatypes.Network_Gateway_Vlan, error)
	DeleteObject() error
	DeleteObjects(templateObjects datatypes.Network_Gateway_Vlan) (bool, error)
	GetObject() (datatypes.Network_Gateway_Vlan, error)
	Unbypass() error
}

type Network_LoadBalancer_Global_Account interface {
	Entity

	AddNsRecord() (bool, error)
	EditObject(templateObject datatypes.Network_LoadBalancer_Global_Account) (bool, error)
	GetObject() (datatypes.Network_LoadBalancer_Global_Account, error)
	RemoveNsRecord() (bool, error)
}

type Network_LoadBalancer_Global_Host interface {
	Entity

	DeleteObject() (bool, error)
	GetObject() (datatypes.Network_LoadBalancer_Global_Host, error)
}

type Network_LoadBalancer_Global_Type interface {
	Entity
}

type Network_LoadBalancer_Service interface {
	Entity

	DeleteObject() (bool, error)
	GetGraphImage(graphType string, metric string) ([]byte, error)
	GetObject() (datatypes.Network_LoadBalancer_Service, error)
	GetStatus() (datatypes.Container_Network_LoadBalancer_StatusEntry, error)
	ResetPeakConnections() (bool, error)
}

type Network_LoadBalancer_VirtualIpAddress interface {
	Entity

	Disable() (bool, error)
	EditObject(templateObject datatypes.Network_LoadBalancer_VirtualIpAddress) (bool, error)
	Enable() (bool, error)
	GetObject() (datatypes.Network_LoadBalancer_VirtualIpAddress, error)
	KickAllConnections() (bool, error)
	UpgradeConnectionLimit() (bool, error)
}

type Network_Logging_Syslog interface {
	Entity
}

type Network_Media_Transcode_Account interface {
	Entity

	CreateTranscodeAccount() (bool, error)
	CreateTranscodeJob(newJob datatypes.Network_Media_Transcode_Job) (bool, error)
	GetDirectoryInformation(directoryName string, extensionFilter string) (datatypes.Container_Network_Directory_Listing, error)
	GetFileDetail(source string) (datatypes.Container_Network_Media_Information, error)
	GetFtpAttributes() (datatypes.Container_Network_Authentication_Data, error)
	GetObject() (datatypes.Network_Media_Transcode_Account, error)
	GetPresetDetail(guid string) (datatypes.Container_Network_Media_Transcode_Preset_Element, error)
	GetPresets() (datatypes.Container_Network_Media_Transcode_Preset, error)
}

type Network_Media_Transcode_Job interface {
	Entity

	CreateObject(templateObject datatypes.Network_Media_Transcode_Job) (datatypes.Network_Media_Transcode_Job, error)
	GetObject() (datatypes.Network_Media_Transcode_Job, error)
}

type Network_Media_Transcode_Job_History interface {
	Entity
}

type Network_Media_Transcode_Job_Status interface {
	Entity

	GetAllStatuses() (datatypes.Network_Media_Transcode_Job_Status, error)
	GetObject() (datatypes.Network_Media_Transcode_Job_Status, error)
}

type Network_Message_Delivery interface {
	Entity

	EditObject(templateObject datatypes.Network_Message_Delivery) (bool, error)
	GetObject() (datatypes.Network_Message_Delivery, error)
}

type Network_Message_Delivery_Attribute interface {
	Entity
}

type Network_Message_Delivery_Email_Sendgrid interface {
	Network_Message_Delivery

	AddUnsubscribeEmailAddress(emailAddress string) (bool, error)
	DeleteEmailListEntries(list string, entries string) (bool, error)
	DisableSmtpAccess() (bool, error)
	EnableSmtpAccess() (bool, error)
	GetAccountOverview() (datatypes.Container_Network_Message_Delivery_Email_Sendgrid_Account_Overview, error)
	GetCategoryList() (string, error)
	GetEmailList(list string) (datatypes.Container_Network_Message_Delivery_Email_Sendgrid_List_Entry, error)
	GetObject() (datatypes.Network_Message_Delivery_Email_Sendgrid, error)
	GetStatistics(options datatypes.Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Options) (datatypes.Container_Network_Message_Delivery_Email_Sendgrid_Statistics, error)
	GetStatisticsGraph(options datatypes.Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Options) (datatypes.Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Graph, error)
	GetVendorPortalUrl() (string, error)
	SendEmail(emailContainer datatypes.Container_Network_Message_Delivery_Email) (bool, error)
	UpdateEmailAddress(emailAddress string) (bool, error)
}

type Network_Message_Delivery_Type interface {
	Entity
}

type Network_Message_Delivery_Vendor interface {
	Entity
}

type Network_Message_Queue interface {
	Entity

	GetObject() (datatypes.Network_Message_Queue, error)
}

type Network_Message_Queue_Node interface {
	Entity

	AddUser(username string) (bool, error)
	DeleteUser(username string) (bool, error)
	GetAllUsers() (string, error)
	GetObject() (datatypes.Network_Message_Queue_Node, error)
	GetUsage(startDate time.Time, endDate time.Time) (datatypes.Metric_Tracking_Object_Data, error)
	GetUsageGraph(graphData datatypes.Container_Graph) (datatypes.Container_Graph, error)
}

type Network_Message_Queue_Status interface {
	Entity

	GetObject() (datatypes.Network_Message_Queue_Status, error)
}

type Network_Monitor interface {
	Entity

	GetIpAddressesByHardware(hardware datatypes.Hardware, partialIpAddress string) (datatypes.Network_Subnet_IpAddress, error)
	GetIpAddressesByVirtualGuest(guest datatypes.Virtual_Guest, partialIpAddress string) (datatypes.Network_Subnet_IpAddress, error)
}

type Network_Monitor_Version1_Incident interface {
	Entity
}

type Network_Monitor_Version1_Query_Host interface {
	Entity

	CreateObject(templateObject datatypes.Network_Monitor_Version1_Query_Host) (datatypes.Network_Monitor_Version1_Query_Host, error)
	CreateObjects(templateObjects datatypes.Network_Monitor_Version1_Query_Host) (datatypes.Network_Monitor_Version1_Query_Host, error)
	DeleteObject() (bool, error)
	DeleteObjects(templateObjects datatypes.Network_Monitor_Version1_Query_Host) (bool, error)
	EditObject(templateObject datatypes.Network_Monitor_Version1_Query_Host) (bool, error)
	EditObjects(templateObjects datatypes.Network_Monitor_Version1_Query_Host) (bool, error)
	FindByHardwareId(hardwareId int) (datatypes.Network_Monitor_Version1_Query_Host, error)
	GetObject() (datatypes.Network_Monitor_Version1_Query_Host, error)
}

type Network_Monitor_Version1_Query_Host_Stratum interface {
	Entity

	GetAllQueryTypes() (datatypes.Network_Monitor_Version1_Query_Type, error)
	GetAllResponseTypes() (datatypes.Network_Monitor_Version1_Query_ResponseType, error)
	GetObject() (datatypes.Network_Monitor_Version1_Query_Host_Stratum, error)
}

type Network_Monitor_Version1_Query_ResponseType interface {
	Entity
}

type Network_Monitor_Version1_Query_Result interface {
	Entity
}

type Network_Monitor_Version1_Query_Type interface {
	Entity
}

type Network_Pod interface {
	Entity

	GetAllObjects() (datatypes.Network_Pod, error)
	GetCapabilities() (string, error)
	GetObject() (datatypes.Network_Pod, error)
	ListCapabilities() (string, error)
}

type Network_Protection_Address interface {
	Entity
}

type Network_Regional_Internet_Registry interface {
	Entity
}

type Network_Security_Scanner_Request interface {
	Entity

	CreateObject(templateObject datatypes.Network_Security_Scanner_Request) (datatypes.Network_Security_Scanner_Request, error)
	GetObject() (datatypes.Network_Security_Scanner_Request, error)
	GetReport() (string, error)
}

type Network_Security_Scanner_Request_Status interface {
	Entity
}

type Network_Service_Health interface {
	Entity
}

type Network_Service_Health_Status interface {
	Entity
}

type Network_Service_Resource interface {
	Entity
}

type Network_Service_Resource_Attribute interface {
	Entity
}

type Network_Service_Resource_Attribute_Type interface {
	Entity
}

type Network_Service_Resource_Hub interface {
	Network_Service_Resource
}

type Network_Service_Resource_Hub_Swift interface {
	Network_Service_Resource_Hub
}

type Network_Service_Resource_MonitoringHub interface {
	Network_Service_Resource
}

type Network_Service_Resource_NimsoftLandingHub interface {
	Network_Service_Resource_MonitoringHub
}

type Network_Service_Resource_Type interface {
	Entity
}

type Network_Service_Vpn_Overrides interface {
	Entity

	CreateObjects(templateObjects datatypes.Network_Service_Vpn_Overrides) (bool, error)
	DeleteObject() (bool, error)
	DeleteObjects(templateObjects datatypes.Network_Service_Vpn_Overrides) (bool, error)
	GetObject() (datatypes.Network_Service_Vpn_Overrides, error)
}

type Network_Storage interface {
	Entity

	AllowAccessFromHardware(hardwareObjectTemplate datatypes.Hardware) (bool, error)
	AllowAccessFromHardwareList(hardwareObjectTemplates datatypes.Hardware) (bool, error)
	AllowAccessFromHost(typeClassName string, hostId int) (datatypes.Network_Storage_Allowed_Host, error)
	AllowAccessFromHostList(hostObjectTemplates datatypes.Container_Network_Storage_Host) (datatypes.Network_Storage_Allowed_Host, error)
	AllowAccessFromIpAddress(ipAddressObjectTemplate datatypes.Network_Subnet_IpAddress) (bool, error)
	AllowAccessFromIpAddressList(ipAddressObjectTemplates datatypes.Network_Subnet_IpAddress) (bool, error)
	AllowAccessFromSubnet(subnetObjectTemplate datatypes.Network_Subnet) (bool, error)
	AllowAccessFromSubnetList(subnetObjectTemplates datatypes.Network_Subnet) (bool, error)
	AllowAccessFromVirtualGuest(virtualGuestObjectTemplate datatypes.Virtual_Guest) (bool, error)
	AllowAccessFromVirtualGuestList(virtualGuestObjectTemplates datatypes.Virtual_Guest) (bool, error)
	AllowAccessToReplicantFromHardware(hardwareObjectTemplate datatypes.Hardware) (bool, error)
	AllowAccessToReplicantFromHardwareList(hardwareObjectTemplates datatypes.Hardware) (bool, error)
	AllowAccessToReplicantFromIpAddress(ipAddressObjectTemplate datatypes.Network_Subnet_IpAddress) (bool, error)
	AllowAccessToReplicantFromIpAddressList(ipAddressObjectTemplates datatypes.Network_Subnet_IpAddress) (bool, error)
	AllowAccessToReplicantFromSubnet(subnetObjectTemplate datatypes.Network_Subnet) (bool, error)
	AllowAccessToReplicantFromSubnetList(subnetObjectTemplates datatypes.Network_Subnet) (bool, error)
	AllowAccessToReplicantFromVirtualGuest(virtualGuestObjectTemplate datatypes.Virtual_Guest) (bool, error)
	AllowAccessToReplicantFromVirtualGuestList(virtualGuestObjectTemplates datatypes.Virtual_Guest) (bool, error)
	AssignCredential(username string) (bool, error)
	AssignNewCredential(typ string) (datatypes.Network_Storage_Credential, error)
	ChangePassword(username string, currentPassword string, newPassword string) (bool, error)
	CollectBandwidth(typ string, startDate time.Time, endDate time.Time) (uint, error)
	CollectBytesUsed() (uint, error)
	CreateFolder(folder string) (bool, error)
	CreateSnapshot(notes string) (datatypes.Network_Storage, error)
	DeleteAllFiles() (bool, error)
	DeleteFile(fileId string) (bool, error)
	DeleteFiles(fileIds string) (bool, error)
	DeleteFolder(folder string) (bool, error)
	DeleteObject() (bool, error)
	DisableSnapshots(scheduleType string) (bool, error)
	DownloadFile(fileId string) (datatypes.Container_Utility_File_Entity, error)
	EditCredential(username string, newPassword string) (bool, error)
	EditObject(templateObject datatypes.Network_Storage) (bool, error)
	EnableSnapshots(scheduleType string, retentionCount int, minute int, hour int, dayOfWeek string) (bool, error)
	FailbackFromReplicant() (bool, error)
	FailoverToReplicant(replicantId int) (bool, error)
	GetAllFiles() (datatypes.Container_Utility_File_Entity, error)
	GetAllFilesByFilter(filter datatypes.Container_Utility_File_Entity) (datatypes.Container_Utility_File_Entity, error)
	GetAllowableHardware(filterHostname string) (datatypes.Hardware, error)
	GetAllowableIpAddresses(subnetId int, filterIpAddress string) (datatypes.Network_Subnet_IpAddress, error)
	GetAllowableSubnets(filterNetworkIdentifier string) (datatypes.Network_Subnet, error)
	GetAllowableVirtualGuests(filterHostname string) (datatypes.Virtual_Guest, error)
	GetAllowedHostsLimit() (int, error)
	GetByUsername(username string, typ string) (datatypes.Network_Storage, error)
	GetCdnUrls() (datatypes.Container_Network_Storage_Hub_ObjectStorage_ContentDeliveryUrl, error)
	GetClusterResource() (datatypes.Network_Service_Resource, error)
	GetFileByIdentifier(identifier string) (datatypes.Container_Utility_File_Entity, error)
	GetFileCount() (int, error)
	GetFileList(folder string, path string) (datatypes.Container_Utility_File_Entity, error)
	GetFilePendingDeleteCount() (int, error)
	GetFilesPendingDelete() (datatypes.Container_Utility_File_Entity, error)
	GetFolderList() (datatypes.Container_Network_Storage_Hub_ObjectStorage_Folder, error)
	GetGraph(startDate time.Time, endDate time.Time, typ string) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetNetworkConnectionDetails() (datatypes.Container_Network_Storage_NetworkConnectionInformation, error)
	GetObject() (datatypes.Network_Storage, error)
	GetObjectStorageConnectionInformation() (datatypes.Container_Network_Service_Resource_ObjectStorage_ConnectionInformation, error)
	GetObjectsByCredential(credentialObject datatypes.Network_Storage_Credential) (datatypes.Network_Storage, error)
	GetRecycleBinFileByIdentifier(fileId string) (datatypes.Container_Utility_File_Entity, error)
	GetRemainingAllowedHosts() (int, error)
	GetSnapshotsForVolume() (datatypes.Network_Storage, error)
	GetStorageGroupsNetworkConnectionDetails() (datatypes.Container_Network_Storage_NetworkConnectionInformation, error)
	GetValidReplicationTargetDatacenterLocations() (datatypes.Location, error)
	ImmediateFailoverToReplicant(replicantId int) (bool, error)
	IsBlockingOperationInProgress(exemptStatusKeyNames string) (bool, error)
	RemoveAccessFromHardware(hardwareObjectTemplate datatypes.Hardware) (bool, error)
	RemoveAccessFromHardwareList(hardwareObjectTemplates datatypes.Hardware) (bool, error)
	RemoveAccessFromHost(typeClassName string, hostId int) (datatypes.Network_Storage_Allowed_Host, error)
	RemoveAccessFromHostList(hostObjectTemplates datatypes.Container_Network_Storage_Host) (datatypes.Network_Storage_Allowed_Host, error)
	RemoveAccessFromIpAddress(ipAddressObjectTemplate datatypes.Network_Subnet_IpAddress) (bool, error)
	RemoveAccessFromIpAddressList(ipAddressObjectTemplates datatypes.Network_Subnet_IpAddress) (bool, error)
	RemoveAccessFromSubnet(subnetObjectTemplate datatypes.Network_Subnet) (bool, error)
	RemoveAccessFromSubnetList(subnetObjectTemplates datatypes.Network_Subnet) (bool, error)
	RemoveAccessFromVirtualGuest(virtualGuestObjectTemplate datatypes.Virtual_Guest) (bool, error)
	RemoveAccessFromVirtualGuestList(virtualGuestObjectTemplates datatypes.Virtual_Guest) (bool, error)
	RemoveAccessToReplicantFromHardwareList(hardwareObjectTemplates datatypes.Hardware) (bool, error)
	RemoveAccessToReplicantFromIpAddressList(ipAddressObjectTemplates datatypes.Network_Subnet_IpAddress) (bool, error)
	RemoveAccessToReplicantFromSubnet(subnetObjectTemplate datatypes.Network_Subnet) (bool, error)
	RemoveAccessToReplicantFromSubnetList(subnetObjectTemplates datatypes.Network_Subnet) (bool, error)
	RemoveAccessToReplicantFromVirtualGuestList(virtualGuestObjectTemplates datatypes.Virtual_Guest) (bool, error)
	RemoveCredential(username string) (bool, error)
	RestoreFile(fileId string) (datatypes.Container_Utility_File_Entity, error)
	RestoreFromSnapshot(snapshotId int) (bool, error)
	SendPasswordReminderEmail(username string) (bool, error)
	SetMountable(mountable bool) (bool, error)
	SetSnapshotAllocation(capacityGb int) error
	UpgradeVolumeCapacity(itemId int) (bool, error)
	UploadFile(file datatypes.Container_Utility_File_Entity) (datatypes.Container_Utility_File_Entity, error)
}

type Network_Storage_Allowed_Host interface {
	Entity

	CreateObject(templateObject datatypes.Network_Storage_Allowed_Host) (bool, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Network_Storage_Allowed_Host) (bool, error)
	GetObject() (datatypes.Network_Storage_Allowed_Host, error)
	SetCredentialPassword(password string) (bool, error)
}

type Network_Storage_Allowed_Host_Hardware interface {
	Network_Storage_Allowed_Host

	GetObject() (datatypes.Network_Storage_Allowed_Host_Hardware, error)
}

type Network_Storage_Allowed_Host_IpAddress interface {
	Network_Storage_Allowed_Host

	GetObject() (datatypes.Network_Storage_Allowed_Host_IpAddress, error)
}

type Network_Storage_Allowed_Host_Subnet interface {
	Network_Storage_Allowed_Host

	GetObject() (datatypes.Network_Storage_Allowed_Host_Subnet, error)
}

type Network_Storage_Allowed_Host_VirtualGuest interface {
	Network_Storage_Allowed_Host

	GetObject() (datatypes.Network_Storage_Allowed_Host_VirtualGuest, error)
}

type Network_Storage_Backup interface {
	Network_Storage
}

type Network_Storage_Backup_Evault interface {
	Network_Storage_Backup

	DeleteTasks(tasks int) (bool, error)
	GetHardwareWithEvaultFirst(option string, exactMatch bool, criteria string, mode string) (datatypes.Hardware, error)
	GetObject() (datatypes.Network_Storage_Backup_Evault, error)
	GetWebCCAuthenticationDetails() (datatypes.Container_Network_Storage_Backup_Evault_WebCc_Authentication_Details, error)
	InitiateBareMetalRestore() (bool, error)
	InitiateBareMetalRestoreForServer(hardwareId int) (bool, error)
}

type Network_Storage_Backup_Evault_Version6 interface {
	Network_Storage_Backup_Evault
}

type Network_Storage_Credential interface {
	Entity
}

type Network_Storage_Credential_Type interface {
	Entity
}

type Network_Storage_Daily_Usage interface {
	Entity
}

type Network_Storage_Event interface {
	Entity
}

type Network_Storage_Group interface {
	Entity

	AddAllowedHost(allowedHost datatypes.Network_Storage_Allowed_Host) (bool, error)
	AttachToVolume(volume datatypes.Network_Storage) (bool, error)
	CreateObject(templateObject datatypes.Network_Storage_Group) (bool, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Network_Storage_Group) (bool, error)
	GetAllObjects() (datatypes.Network_Storage_Group, error)
	GetNetworkConnectionDetails() (datatypes.Container_Network_Storage_NetworkConnectionInformation, error)
	GetObject() (datatypes.Network_Storage_Group, error)
	RemoveAllowedHost(allowedHost datatypes.Network_Storage_Allowed_Host) (bool, error)
	RemoveFromVolume(volume datatypes.Network_Storage) (bool, error)
}

type Network_Storage_Group_Iscsi interface {
	Network_Storage_Group

	AddAllowedHost(allowedHost datatypes.Network_Storage_Allowed_Host) (bool, error)
	AttachToVolume(volume datatypes.Network_Storage) (bool, error)
	GetObject() (datatypes.Network_Storage_Group_Iscsi, error)
	RemoveAllowedHost(allowedHost datatypes.Network_Storage_Allowed_Host) (bool, error)
	RemoveFromVolume(volume datatypes.Network_Storage) (bool, error)
}

type Network_Storage_Group_Nfs interface {
	Network_Storage_Group

	AddAllowedHost(allowedHost datatypes.Network_Storage_Allowed_Host) (bool, error)
	AttachToVolume(volume datatypes.Network_Storage) (bool, error)
	GetObject() (datatypes.Network_Storage_Group_Nfs, error)
	RemoveAllowedHost(allowedHost datatypes.Network_Storage_Allowed_Host) (bool, error)
	RemoveFromVolume(volume datatypes.Network_Storage) (bool, error)
}

type Network_Storage_Group_Type interface {
	Entity

	GetAllObjects() (datatypes.Network_Storage_Group_Type, error)
	GetObject() (datatypes.Network_Storage_Group_Type, error)
}

type Network_Storage_History interface {
	Entity
}

type Network_Storage_Hub interface {
	Network_Storage
}

type Network_Storage_Hub_Cleversafe_Account interface {
	Entity

	CredentialCreate() (datatypes.Network_Storage_Credential, error)
	CredentialDelete(credential datatypes.Network_Storage_Credential) (bool, error)
	GetCapacityUsage() (int, error)
	GetCredentialLimit() (int, error)
	GetEndpoints() (datatypes.Container_Network_Storage_Hub_ObjectStorage_Endpoint, error)
	GetObject() (datatypes.Network_Storage_Hub_Cleversafe_Account, error)
}

type Network_Storage_Hub_Swift interface {
	Network_Storage_Hub
}

type Network_Storage_Hub_Swift_Container interface {
	Network_Storage_Hub_Swift
}

type Network_Storage_Hub_Swift_Share interface {
	Entity

	GetContainerList() (datatypes.Container_Network_Storage_Hub_ObjectStorage_Folder, error)
	GetFile(fileName string, container string) (datatypes.Container_Network_Storage_Hub_ObjectStorage_File, error)
	GetFileList(container string, path string) (datatypes.Container_Utility_File_Entity, error)
}

type Network_Storage_Hub_Swift_Version1 interface {
	Network_Storage_Hub_Swift
}

type Network_Storage_Iscsi interface {
	Network_Storage

	AllowAccessFromHardware(hardwareObjectTemplate datatypes.Hardware) (bool, error)
	AllowAccessFromIpAddress(ipAddressObjectTemplate datatypes.Network_Subnet_IpAddress) (bool, error)
	AllowAccessFromVirtualGuest(virtualGuestObjectTemplate datatypes.Virtual_Guest) (bool, error)
	AllowAccessToReplicantFromHardwareList(hardwareObjectTemplates datatypes.Hardware) (bool, error)
	AllowAccessToReplicantFromIpAddressList(ipAddressObjectTemplates datatypes.Network_Subnet_IpAddress) (bool, error)
	AllowAccessToReplicantFromVirtualGuestList(virtualGuestObjectTemplates datatypes.Virtual_Guest) (bool, error)
	GetObject() (datatypes.Network_Storage_Iscsi, error)
	GetSnapshotsForVolume() (datatypes.Network_Storage, error)
	RemoveAccessFromHardware(hardwareObjectTemplate datatypes.Hardware) (bool, error)
	RemoveAccessFromIpAddress(ipAddressObjectTemplate datatypes.Network_Subnet_IpAddress) (bool, error)
	RemoveAccessFromVirtualGuest(virtualGuestObjectTemplate datatypes.Virtual_Guest) (bool, error)
	RemoveAccessToReplicantFromHardwareList(hardwareObjectTemplates datatypes.Hardware) (bool, error)
	RemoveAccessToReplicantFromIpAddressList(ipAddressObjectTemplates datatypes.Network_Subnet_IpAddress) (bool, error)
	RemoveAccessToReplicantFromVirtualGuestList(virtualGuestObjectTemplates datatypes.Virtual_Guest) (bool, error)
}

type Network_Storage_Iscsi_EqualLogic_Version3 interface {
	Network_Storage_Iscsi
}

type Network_Storage_Iscsi_EqualLogic_Version3_Replicant interface {
	Network_Storage_Iscsi_EqualLogic_Version3
}

type Network_Storage_Iscsi_EqualLogic_Version3_Snapshot interface {
	Network_Storage_Iscsi_EqualLogic_Version3
}

type Network_Storage_Iscsi_OS_Type interface {
	Entity

	GetAllObjects() (datatypes.Network_Storage_Iscsi_OS_Type, error)
	GetObject() (datatypes.Network_Storage_Iscsi_OS_Type, error)
}

type Network_Storage_Nas interface {
	Network_Storage
}

type Network_Storage_OpenStack_Object interface {
	Network_Storage
}

type Network_Storage_Partnership interface {
	Entity
}

type Network_Storage_Partnership_Type interface {
	Entity
}

type Network_Storage_Property interface {
	Entity
}

type Network_Storage_Property_Type interface {
	Entity
}

type Network_Storage_Replicant interface {
	Network_Storage
}

type Network_Storage_Schedule interface {
	Entity

	CreateObject(templateObject datatypes.Network_Storage_Schedule) (datatypes.Network_Storage_Schedule, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Network_Storage_Schedule) (bool, error)
	GetObject() (datatypes.Network_Storage_Schedule, error)
}

type Network_Storage_Schedule_Property interface {
	Entity
}

type Network_Storage_Schedule_Property_Type interface {
	Entity

	GetAllObjects() (datatypes.Network_Storage_Schedule_Property_Type, error)
	GetObject() (datatypes.Network_Storage_Schedule_Property_Type, error)
}

type Network_Storage_Schedule_Type interface {
	Entity
}

type Network_Storage_Snapshot interface {
	Network_Storage
}

type Network_Storage_Type interface {
	Entity
}

type Network_Subnet interface {
	Entity

	AllowAccessToNetworkStorage(networkStorageTemplateObject datatypes.Network_Storage) (bool, error)
	AllowAccessToNetworkStorageList(networkStorageTemplateObjects datatypes.Network_Storage) (bool, error)
	CreateReverseDomainRecords() (datatypes.Dns_Domain_Reverse, error)
	CreateSubnetRouteUpdateTransaction(newEndPointIpAddress string) (bool, error)
	CreateSwipTransaction() (bool, error)
	EditNote(note string) (bool, error)
	FindAllSubnetsAndActiveSwipTransactionStatus() (datatypes.Network_Subnet, error)
	GetAttachedNetworkStorages(nasType string) (datatypes.Network_Storage, error)
	GetAvailableNetworkStorages(nasType string) (datatypes.Network_Storage, error)
	GetObject() (datatypes.Network_Subnet, error)
	GetReverseDomainRecords() (datatypes.Dns_Domain, error)
	GetRoutableEndpointIpAddresses() (datatypes.Network_Subnet_IpAddress, error)
	GetSubnetForIpAddress(ipAddress string) (datatypes.Network_Subnet, error)
	RemoveAccessToNetworkStorageList(networkStorageTemplateObjects datatypes.Network_Storage) (bool, error)
}

type Network_Subnet_IpAddress interface {
	Entity

	AllowAccessToNetworkStorage(networkStorageTemplateObject datatypes.Network_Storage) (bool, error)
	AllowAccessToNetworkStorageList(networkStorageTemplateObjects datatypes.Network_Storage) (bool, error)
	EditObject(templateObject datatypes.Network_Subnet_IpAddress) (bool, error)
	EditObjects(templateObjects datatypes.Network_Subnet_IpAddress) (bool, error)
	FindByIpv4Address(ipAddress string) (datatypes.Network_Subnet_IpAddress, error)
	GetAttachedNetworkStorages(nasType string) (datatypes.Network_Storage, error)
	GetAvailableNetworkStorages(nasType string) (datatypes.Network_Storage, error)
	GetByIpAddress(ipAddress string) (datatypes.Network_Subnet_IpAddress, error)
	GetObject() (datatypes.Network_Subnet_IpAddress, error)
	RemoveAccessToNetworkStorageList(networkStorageTemplateObjects datatypes.Network_Storage) (bool, error)
}

type Network_Subnet_IpAddress_Global interface {
	Entity

	GetObject() (datatypes.Network_Subnet_IpAddress_Global, error)
	Route(newEndPointIpAddress string) (datatypes.Provisioning_Version1_Transaction, error)
	Unroute() (datatypes.Provisioning_Version1_Transaction, error)
}

type Network_Subnet_IpAddress_Version6 interface {
	Network_Subnet_IpAddress
}

type Network_Subnet_Registration interface {
	Entity

	ClearRegistration() (bool, error)
	CreateObject(templateObject datatypes.Network_Subnet_Registration) (datatypes.Network_Subnet_Registration, error)
	EditObject(templateObject datatypes.Network_Subnet_Registration) (bool, error)
	EditRegistrationAttachedDetails(personObjectSkeleton datatypes.Network_Subnet_Registration_Details, networkObjectSkeleton datatypes.Network_Subnet_Registration_Details) (bool, error)
	GetObject() (datatypes.Network_Subnet_Registration, error)
}

type Network_Subnet_Registration_Apnic interface {
	Network_Subnet_Registration
}

type Network_Subnet_Registration_Arin interface {
	Network_Subnet_Registration
}

type Network_Subnet_Registration_Details interface {
	Entity

	CreateObject(templateObject datatypes.Network_Subnet_Registration_Details) (datatypes.Network_Subnet_Registration_Details, error)
	DeleteObject() (bool, error)
	GetObject() (datatypes.Network_Subnet_Registration_Details, error)
}

type Network_Subnet_Registration_Event interface {
	Entity
}

type Network_Subnet_Registration_Event_Type interface {
	Entity
}

type Network_Subnet_Registration_Ripe interface {
	Network_Subnet_Registration
}

type Network_Subnet_Registration_Status interface {
	Entity

	GetAllObjects() (datatypes.Network_Subnet_Registration_Status, error)
	GetObject() (datatypes.Network_Subnet_Registration_Status, error)
}

type Network_Subnet_Rwhois_Data interface {
	Entity

	EditObject(templateObject datatypes.Network_Subnet_Rwhois_Data) (bool, error)
	GetObject() (datatypes.Network_Subnet_Rwhois_Data, error)
}

type Network_Subnet_Swip_Transaction interface {
	Entity

	FindMyTransactions() (datatypes.Network_Subnet_Swip_Transaction, error)
	GetObject() (datatypes.Network_Subnet_Swip_Transaction, error)
	RemoveAllSubnetSwips() (int, error)
	RemoveSwipData() (bool, error)
	ResendSwipData() (bool, error)
	SwipAllSubnets() (int, error)
	UpdateAllSubnetSwips() (int, error)
}

type Network_TippingPointReporting interface {
	Entity

	DrillDownAttack(signatureId string, IpAddress string, subnetMask int, timeFrame int, direction string) (datatypes.Container_Network_IntrusionProtection_SubnetReport, error)
	GetMainStatistics(numberOfAttacks int) (datatypes.Container_Network_IntrusionProtection_Statistics, error)
	GetReportForIpAddressOrSubnet(IpAddress string, subnetMask int, timeFrame int, orderBy string, orderDirection string) (datatypes.Container_Network_IntrusionProtection_SubnetReport, error)
	GetSubnetReportForEntireAccount(timeFrame int, orderBy string, orderDirection string, returnSubnetGroups bool) (datatypes.Container_Network_IntrusionProtection_SubnetReport, error)
}

type Network_Tunnel_Module_Context interface {
	Entity

	AddCustomerSubnetToNetworkTunnel(subnetId int) (bool, error)
	AddPrivateSubnetToNetworkTunnel(subnetId int) (bool, error)
	AddServiceSubnetToNetworkTunnel(subnetId int) (bool, error)
	ApplyConfigurationsToDevice() (bool, error)
	CreateAddressTranslation(translation datatypes.Network_Tunnel_Module_Context_Address_Translation) (datatypes.Network_Tunnel_Module_Context_Address_Translation, error)
	CreateAddressTranslations(translations datatypes.Network_Tunnel_Module_Context_Address_Translation) (datatypes.Network_Tunnel_Module_Context_Address_Translation, error)
	DeleteAddressTranslation(translationId int) (bool, error)
	DownloadAddressTranslationConfigurations() (datatypes.Container_Utility_File_Entity, error)
	DownloadParameterConfigurations() (datatypes.Container_Utility_File_Entity, error)
	EditAddressTranslation(translation datatypes.Network_Tunnel_Module_Context_Address_Translation) (datatypes.Network_Tunnel_Module_Context_Address_Translation, error)
	EditAddressTranslations(translations datatypes.Network_Tunnel_Module_Context_Address_Translation) (datatypes.Network_Tunnel_Module_Context_Address_Translation, error)
	EditObject(templateObject datatypes.Network_Tunnel_Module_Context) (bool, error)
	GetAddressTranslationConfigurations() (string, error)
	GetAuthenticationDefault() (string, error)
	GetAuthenticationOptions() (string, error)
	GetDiffieHellmanGroupDefault() (int, error)
	GetDiffieHellmanGroupOptions() (int, error)
	GetEncryptionDefault() (string, error)
	GetEncryptionOptions() (string, error)
	GetKeylifeLimits() (int, error)
	GetObject() (datatypes.Network_Tunnel_Module_Context, error)
	GetParameterConfigurationsForCustomerView() (string, error)
	GetPhaseOneKeylifeDefault() (string, error)
	GetPhaseTwoKeylifeDefault() (string, error)
	RemoveCustomerSubnetFromNetworkTunnel(subnetId int) (bool, error)
	RemovePrivateSubnetFromNetworkTunnel(subnetId int) (bool, error)
	RemoveServiceSubnetFromNetworkTunnel(subnetId int) (bool, error)
}

type Network_Tunnel_Module_Context_Address_Translation interface {
	Entity
}

type Network_Vlan interface {
	Entity

	EditObject(templateObject datatypes.Network_Vlan) (bool, error)
	GetCancelFailureReasons() (string, error)
	GetFirewallProtectableIpAddresses() (datatypes.Network_Subnet_IpAddress, error)
	GetFirewallProtectableSubnets() (datatypes.Network_Subnet, error)
	GetObject() (datatypes.Network_Vlan, error)
	GetPrivateVlan() (datatypes.Network_Vlan, error)
	GetPrivateVlanByIpAddress(ipAddress string) (datatypes.Network_Vlan, error)
	GetPublicVlanByFqdn(fqdn string) (datatypes.Network_Vlan, error)
	GetReverseDomainRecords() (datatypes.Dns_Domain, error)
	GetVlanForIpAddress(ipAddress string) (datatypes.Network_Vlan, error)
	SetTags(tags string) (bool, error)
	UpdateFirewallIntraVlanCommunication(enabled bool) error
}

type Network_Vlan_Firewall interface {
	Entity

	GetObject() (datatypes.Network_Vlan_Firewall, error)
	RestoreDefaults() (datatypes.Provisioning_Version1_Transaction, error)
	SetTags(tags string) (bool, error)
	UpdateRouteBypass(bypass bool) (datatypes.Provisioning_Version1_Transaction, error)
}

type Network_Vlan_Firewall_Rule interface {
	Entity
}

type Network_Vlan_Type interface {
	Entity

	GetObject() (datatypes.Network_Vlan_Type, error)
}

type Notification interface {
	Entity

	GetAllObjects() (datatypes.Notification, error)
	GetObject() (datatypes.Notification, error)
}

type Notification_Delivery_Method interface {
	Entity
}

type Notification_Mobile interface {
	Notification

	CreateSubscriberForMobileDevice(keyName string, resourceTableId int, userRecordId int) (bool, error)
	GetObject() (datatypes.Notification_Mobile, error)
}

type Notification_Occurrence_Account interface {
	Entity
}

type Notification_Occurrence_Event interface {
	Entity

	AcknowledgeNotification() (bool, error)
	GetAllObjects() (datatypes.Notification_Occurrence_Event, error)
	GetAttachedFile(attachmentId int) ([]byte, error)
	GetImpactedAccountCount() (int, error)
	GetImpactedDeviceCount() (int, error)
	GetImpactedDevices() (datatypes.Notification_Occurrence_Resource, error)
	GetObject() (datatypes.Notification_Occurrence_Event, error)
}

type Notification_Occurrence_Event_Attachment interface {
	Entity
}

type Notification_Occurrence_Event_Type interface {
	Entity
}

type Notification_Occurrence_Resource interface {
	Entity
}

type Notification_Occurrence_Resource_Hardware interface {
	Notification_Occurrence_Resource
}

type Notification_Occurrence_Resource_Network_Application_Delivery_Controller interface {
	Notification_Occurrence_Resource
}

type Notification_Occurrence_Resource_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress interface {
	Notification_Occurrence_Resource
}

type Notification_Occurrence_Resource_Network_Storage_Iscsi_EqualLogic interface {
	Notification_Occurrence_Resource
}

type Notification_Occurrence_Resource_Network_Storage_Iscsi_NetApp interface {
	Notification_Occurrence_Resource
}

type Notification_Occurrence_Resource_Network_Storage_Lockbox interface {
	Notification_Occurrence_Resource
}

type Notification_Occurrence_Resource_Network_Storage_Nas interface {
	Notification_Occurrence_Resource
}

type Notification_Occurrence_Resource_Network_Storage_NetApp_Volume interface {
	Notification_Occurrence_Resource
}

type Notification_Occurrence_Resource_Virtual interface {
	Notification_Occurrence_Resource
}

type Notification_Occurrence_Status_Code interface {
	Entity
}

type Notification_Occurrence_Update interface {
	Entity
}

type Notification_Occurrence_User interface {
	Entity

	Acknowledge() (bool, error)
	GetAllObjects() (datatypes.Notification_Occurrence_User, error)
	GetImpactedDeviceCount() (int, error)
	GetObject() (datatypes.Notification_Occurrence_User, error)
}

type Notification_Preference interface {
	Entity
}

type Notification_Subscriber interface {
	Entity
}

type Notification_Subscriber_Customer interface {
	Notification_Subscriber
}

type Notification_Subscriber_Delivery_Method interface {
	Entity
}

type Notification_User_Subscriber interface {
	Entity

	CreateObject(templateObject datatypes.Notification_User_Subscriber) (bool, error)
	EditObject(templateObject datatypes.Notification_User_Subscriber) (bool, error)
	GetObject() (datatypes.Notification_User_Subscriber, error)
}

type Notification_User_Subscriber_Billing interface {
	Notification_User_Subscriber

	GetObject() (datatypes.Notification_User_Subscriber_Billing, error)
}

type Notification_User_Subscriber_Delivery_Method interface {
	Entity
}

type Notification_User_Subscriber_Mobile interface {
	Notification_User_Subscriber

	ClearSnoozeTimer() (bool, error)
	GetObject() (datatypes.Notification_User_Subscriber_Mobile, error)
	SetSnoozeTimer(start int, end int) (bool, error)
}

type Notification_User_Subscriber_Preference interface {
	Entity

	CreateObject(templateObject datatypes.Notification_User_Subscriber_Preference) (bool, error)
	EditObjects(templateObjects datatypes.Notification_User_Subscriber_Preference) (bool, error)
	GetObject() (datatypes.Notification_User_Subscriber_Preference, error)
}

type Notification_User_Subscriber_Resource interface {
	Entity
}

type Product_Catalog interface {
	Entity
}

type Product_Catalog_Item_Price interface {
	Entity
}

type Product_Item interface {
	Entity
}

type Product_Item_Attribute interface {
	Entity
}

type Product_Item_Attribute_Type interface {
	Entity
}

type Product_Item_Billing_Type interface {
	Entity
}

type Product_Item_Bundles interface {
	Entity
}

type Product_Item_Category interface {
	Entity

	GetAdditionalProductsForCategory() (datatypes.Product_Item, error)
	GetBandwidthCategories() (datatypes.Product_Item_Category, error)
	GetComputingCategories(resetCache bool) (datatypes.Product_Item_Category, error)
	GetCustomUsageRatesCategories() (datatypes.Product_Item_Category, error)
	GetObject() (datatypes.Product_Item_Category, error)
	GetSoftwareCategories() (datatypes.Product_Item_Category, error)
	GetSubnetCategories() (datatypes.Product_Item_Category, error)
	GetTopLevelCategories(resetCache bool) (datatypes.Product_Item_Category, error)
	GetValidCancelableServiceItemCategories() (datatypes.Product_Item_Category, error)
	GetVlanCategories() (datatypes.Product_Item_Category, error)
}

type Product_Item_Category_Group interface {
	Entity

	GetObject() (datatypes.Product_Item_Category_Group, error)
}

type Product_Item_Category_Order_Option_Type interface {
	Entity
}

type Product_Item_Category_Question interface {
	Entity
}

type Product_Item_Category_Question_Field_Type interface {
	Entity
}

type Product_Item_Category_Question_Xref interface {
	Entity
}

type Product_Item_Link_ThePlanet interface {
	Entity
}

type Product_Item_Policy_Assignment interface {
	Entity

	AcceptFromTicket(ticketId int) (bool, error)
	GetObject() (datatypes.Product_Item_Policy_Assignment, error)
	GetPolicyDocumentContents() ([]byte, error)
}

type Product_Item_Price interface {
	Entity

	GetObject() (datatypes.Product_Item_Price, error)
	GetUsageRatePrices(location datatypes.Location, items datatypes.Product_Item) (datatypes.Product_Item_Price, error)
}

type Product_Item_Price_Account_Restriction interface {
	Entity
}

type Product_Item_Price_Attribute interface {
	Entity
}

type Product_Item_Price_Attribute_Type interface {
	Entity
}

type Product_Item_Price_Premium interface {
	Entity

	GetObject() (datatypes.Product_Item_Price_Premium, error)
}

type Product_Item_Requirement interface {
	Entity
}

type Product_Item_Resource_Conflict interface {
	Entity
}

type Product_Item_Resource_Conflict_Item interface {
	Product_Item_Resource_Conflict
}

type Product_Item_Resource_Conflict_Item_Category interface {
	Product_Item_Resource_Conflict
}

type Product_Item_Resource_Conflict_Location interface {
	Product_Item_Resource_Conflict
}

type Product_Item_Tax_Category interface {
	Entity
}

type Product_Order interface {
	Entity

	CheckItemAvailability(itemPrices datatypes.Product_Item_Price, accountId int, availabilityTypeKeyNames string) (bool, error)
	CheckItemAvailabilityForImageTemplate(imageTemplateId int, accountId int, packageId int, availabilityTypeKeyNames string) (bool, error)
	CheckItemConflicts(itemPrices datatypes.Product_Item_Price) (bool, error)
	GetExternalPaymentAuthorizationReceipt(token string, payerId string) (datatypes.Container_Product_Order_Receipt, error)
	GetNetworks(locationId int, packageId int, accountId int) (datatypes.Container_Product_Order_Network, error)
	GetResellerOrder(orderContainer datatypes.Container_Product_Order) (datatypes.Container_Product_Order, error)
	GetTaxCalculationResult(orderHash string) (datatypes.Container_Tax_Cache, error)
	GetVlans(locationId int, packageId int, selectedItems string, vlanIds int, subnetIds int, accountId int, orderContainer datatypes.Container_Product_Order) (datatypes.Container_Product_Order_Network_Vlans, error)
	PlaceOrder(orderData datatypes.Container_Product_Order, saveAsQuote bool) (datatypes.Container_Product_Order_Receipt, error)
	PlaceQuote(orderData datatypes.Container_Product_Order) (datatypes.Container_Product_Order_Receipt, error)
	ProcessExternalPaymentAuthorization(token string, payerId string) (datatypes.Container_Product_Order, error)
	RequiredItems(itemPrices datatypes.Product_Item_Price) (datatypes.Product_Item, error)
	VerifyOrder(orderData datatypes.Container_Product_Order) (datatypes.Container_Product_Order, error)
}

type Product_Package interface {
	Entity

	GetActiveItems() (datatypes.Product_Item, error)
	GetActivePackagesByAttribute(attributeKeyName string) (datatypes.Product_Package, error)
	GetActivePrivateHostedCloudPackages() (datatypes.Product_Package, error)
	GetActiveUsageRatePrices(locationId int, categoryCode string) (datatypes.Product_Item_Price, error)
	GetAllObjects() (datatypes.Product_Package, error)
	GetAvailablePackagesForImageTemplate(imageTemplate datatypes.Virtual_Guest_Block_Device_Template_Group) (datatypes.Product_Package, error)
	GetCdnItems() (datatypes.Product_Item, error)
	GetCloudStorageItems(provider int) (datatypes.Product_Item, error)
	GetItemAvailabilityTypes() (datatypes.Product_Item_Attribute_Type, error)
	GetItemPricesFromSoftwareDescriptions(softwareDescriptions datatypes.Software_Description, includeTranslationsFlag bool, returnAllPricesFlag bool) (datatypes.Product_Item_Price, error)
	GetItemsFromImageTemplate(imageTemplate datatypes.Virtual_Guest_Block_Device_Template_Group) (datatypes.Product_Item, error)
	GetMessageQueueItems() (datatypes.Product_Item, error)
	GetObject() (datatypes.Product_Package, error)
	GetObjectStorageDatacenters() (datatypes.Container_Product_Order_Network_Storage_Hub_Datacenter, error)
	GetStandardCategories() (datatypes.Product_Item_Category, error)
}

type Product_Package_Attribute interface {
	Entity
}

type Product_Package_Attribute_Type interface {
	Entity
}

type Product_Package_Inventory interface {
	Entity
}

type Product_Package_Item_Category_Group interface {
	Entity
}

type Product_Package_Item_Prices interface {
	Entity
}

type Product_Package_Items interface {
	Entity
}

type Product_Package_Locations interface {
	Entity
}

type Product_Package_Order_Configuration interface {
	Entity
}

type Product_Package_Order_Step interface {
	Entity
}

type Product_Package_Order_Step_Next interface {
	Entity
}

type Product_Package_Preset interface {
	Entity

	GetAllObjects() (datatypes.Product_Package_Preset, error)
	GetObject() (datatypes.Product_Package_Preset, error)
}

type Product_Package_Preset_Attribute interface {
	Entity
}

type Product_Package_Preset_Attribute_Type interface {
	Entity
}

type Product_Package_Preset_Configuration interface {
	Entity
}

type Product_Package_Server interface {
	Entity

	GetAllObjects() (datatypes.Product_Package_Server, error)
	GetObject() (datatypes.Product_Package_Server, error)
}

type Product_Package_Server_Option interface {
	Entity

	GetAllOptions() (datatypes.Product_Package_Server_Option, error)
	GetObject() (datatypes.Product_Package_Server_Option, error)
	GetOptions(typ string) (datatypes.Product_Package_Server_Option, error)
}

type Product_Package_Type interface {
	Entity

	GetAllObjects() (datatypes.Product_Package_Type, error)
	GetObject() (datatypes.Product_Package_Type, error)
}

type Product_Upgrade_Request interface {
	Entity

	ApproveChanges() (bool, error)
	GetObject() (datatypes.Product_Upgrade_Request, error)
	UpdateMaintenanceWindow(maintenanceStartTime time.Time, maintenanceWindowId int) (bool, error)
}

type Product_Upgrade_Request_Status interface {
	Entity
}

type Provisioning_Hook interface {
	Entity

	CreateObject(templateObject datatypes.Provisioning_Hook) (datatypes.Provisioning_Hook, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Provisioning_Hook) (bool, error)
	GetObject() (datatypes.Provisioning_Hook, error)
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
}

type Provisioning_Maintenance_Classification_Item_Category interface {
	Entity

	GetObject() (datatypes.Provisioning_Maintenance_Classification_Item_Category, error)
}

type Provisioning_Maintenance_Slots interface {
	Entity

	GetObject() (datatypes.Provisioning_Maintenance_Slots, error)
}

type Provisioning_Maintenance_Ticket interface {
	Entity

	GetObject() (datatypes.Provisioning_Maintenance_Ticket, error)
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
}

type Provisioning_Version1_Transaction_Group interface {
	Entity

	GetAllObjects() (datatypes.Provisioning_Version1_Transaction_Group, error)
	GetObject() (datatypes.Provisioning_Version1_Transaction_Group, error)
}

type Provisioning_Version1_Transaction_History interface {
	Entity
}

type Provisioning_Version1_Transaction_Status interface {
	Entity
}

type Provisioning_Version1_Transaction_SubnetMigration interface {
	Provisioning_Version1_Transaction
}

type Resource_Group interface {
	Entity

	EditObject(templateObject datatypes.Resource_Group) (bool, error)
	GetObject() (datatypes.Resource_Group, error)
}

type Resource_Group_Attribute interface {
	Entity
}

type Resource_Group_Attribute_Type interface {
	Entity
}

type Resource_Group_Descendant_Reference interface {
	Entity
}

type Resource_Group_Member interface {
	Entity
}

type Resource_Group_Member_Attribute interface {
	Entity
}

type Resource_Group_Member_Attribute_Type interface {
	Entity
}

type Resource_Group_Member_CloudStack_Version3_Cluster interface {
	Resource_Group_Member
}

type Resource_Group_Member_CloudStack_Version3_Pod interface {
	Resource_Group_Member
}

type Resource_Group_Member_CloudStack_Version3_Zone interface {
	Resource_Group_Member
}

type Resource_Group_Member_Hardware interface {
	Resource_Group_Member
}

type Resource_Group_Member_Network_Storage interface {
	Resource_Group_Member
}

type Resource_Group_Member_Network_Subnet interface {
	Resource_Group_Member
}

type Resource_Group_Member_Network_Vlan interface {
	Resource_Group_Member
}

type Resource_Group_Member_Resource_Group interface {
	Resource_Group_Member
}

type Resource_Group_Member_Role_Link interface {
	Entity
}

type Resource_Group_Member_Software_Component_Password interface {
	Resource_Group_Member
}

type Resource_Group_Member_Type interface {
	Entity
}

type Resource_Group_Member_Virtual_Host_Pool interface {
	Resource_Group_Member
}

type Resource_Group_Role interface {
	Entity
}

type Resource_Group_Template interface {
	Entity

	GetAllObjects() (datatypes.Resource_Group_Template, error)
	GetObject() (datatypes.Resource_Group_Template, error)
}

type Resource_Group_Template_Member interface {
	Entity
}

type Resource_Metadata interface {
	Entity

	GetBackendMacAddresses() (string, error)
	GetDatacenter() (string, error)
	GetDatacenterId() (int, error)
	GetDomain() (string, error)
	GetFrontendMacAddresses() (string, error)
	GetFullyQualifiedDomainName() (string, error)
	GetHostname() (string, error)
	GetId() (int, error)
	GetPrimaryBackendIpAddress() (string, error)
	GetPrimaryIpAddress() (string, error)
	GetProvisionState() (string, error)
	GetRouter(macAddress string) (string, error)
	GetServiceResource(serviceName string, index int) (string, error)
	GetServiceResources() (datatypes.Container_Resource_Metadata_ServiceResource, error)
	GetTags() (string, error)
	GetUserMetadata() (string, error)
	GetVlanIds(macAddress string) (int, error)
	GetVlans(macAddress string) (int, error)
}

type Sales_Presale_Event interface {
	Entity

	GetAllObjects() (datatypes.Sales_Presale_Event, error)
	GetObject() (datatypes.Sales_Presale_Event, error)
}

type Scale_Asset interface {
	Entity

	DeleteObject() (bool, error)
	GetObject() (datatypes.Scale_Asset, error)
}

type Scale_Asset_Hardware interface {
	Scale_Asset

	CreateObject(templateObject datatypes.Scale_Asset_Hardware) (datatypes.Scale_Asset_Hardware, error)
	GetObject() (datatypes.Scale_Asset_Hardware, error)
}

type Scale_Asset_Virtual_Guest interface {
	Scale_Asset

	CreateObject(templateObject datatypes.Scale_Asset_Virtual_Guest) (datatypes.Scale_Asset_Virtual_Guest, error)
	GetObject() (datatypes.Scale_Asset_Virtual_Guest, error)
}

type Scale_Group interface {
	Entity

	CreateObject(templateObject datatypes.Scale_Group) (datatypes.Scale_Group, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Scale_Group) (bool, error)
	ForceDeleteObject() (bool, error)
	GetAvailableHourlyInstanceLimit() (int, error)
	GetAvailableRegionalGroups() (datatypes.Location_Group, error)
	GetObject() (datatypes.Scale_Group, error)
	Resume() error
	Scale(delta int) (datatypes.Scale_Member, error)
	ScaleTo(number int) (datatypes.Scale_Member, error)
	Suspend() error
}

type Scale_Group_Log interface {
	Entity
}

type Scale_Group_Status interface {
	Entity

	GetAllObjects() (datatypes.Scale_Group_Status, error)
	GetObject() (datatypes.Scale_Group_Status, error)
}

type Scale_LoadBalancer interface {
	Entity

	CreateObject(templateObject datatypes.Scale_LoadBalancer) (datatypes.Scale_LoadBalancer, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Scale_LoadBalancer) (bool, error)
	GetObject() (datatypes.Scale_LoadBalancer, error)
}

type Scale_Member interface {
	Entity

	DeleteObject() (bool, error)
	GetObject() (datatypes.Scale_Member, error)
}

type Scale_Member_Virtual_Guest interface {
	Scale_Member

	DeleteObject() (bool, error)
	GetObject() (datatypes.Scale_Member_Virtual_Guest, error)
}

type Scale_Network_Vlan interface {
	Entity

	CreateObject(templateObject datatypes.Scale_Network_Vlan) (datatypes.Scale_Network_Vlan, error)
	DeleteObject() (bool, error)
	GetObject() (datatypes.Scale_Network_Vlan, error)
}

type Scale_Policy interface {
	Entity

	CreateObject(templateObject datatypes.Scale_Policy) (datatypes.Scale_Policy, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Scale_Policy) (bool, error)
	GetObject() (datatypes.Scale_Policy, error)
	Trigger() (datatypes.Scale_Member, error)
}

type Scale_Policy_Action interface {
	Entity

	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Scale_Policy_Action) (bool, error)
	GetObject() (datatypes.Scale_Policy_Action, error)
}

type Scale_Policy_Action_Scale interface {
	Scale_Policy_Action

	CreateObject(templateObject datatypes.Scale_Policy_Action_Scale) (datatypes.Scale_Policy_Action_Scale, error)
	GetObject() (datatypes.Scale_Policy_Action_Scale, error)
}

type Scale_Policy_Action_Type interface {
	Entity

	GetAllObjects() (datatypes.Scale_Policy_Action_Type, error)
	GetObject() (datatypes.Scale_Policy_Action_Type, error)
}

type Scale_Policy_Trigger interface {
	Entity

	CreateObject(templateObject datatypes.Scale_Policy_Trigger) (datatypes.Scale_Policy_Trigger, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Scale_Policy_Trigger) (bool, error)
	GetObject() (datatypes.Scale_Policy_Trigger, error)
}

type Scale_Policy_Trigger_OneTime interface {
	Scale_Policy_Trigger

	CreateObject(templateObject datatypes.Scale_Policy_Trigger_OneTime) (datatypes.Scale_Policy_Trigger_OneTime, error)
	GetObject() (datatypes.Scale_Policy_Trigger_OneTime, error)
}

type Scale_Policy_Trigger_Repeating interface {
	Scale_Policy_Trigger

	CreateObject(templateObject datatypes.Scale_Policy_Trigger_Repeating) (datatypes.Scale_Policy_Trigger_Repeating, error)
	GetObject() (datatypes.Scale_Policy_Trigger_Repeating, error)
	ValidateCronExpression(expression string) error
}

type Scale_Policy_Trigger_ResourceUse interface {
	Scale_Policy_Trigger

	CreateObject(templateObject datatypes.Scale_Policy_Trigger_ResourceUse) (datatypes.Scale_Policy_Trigger_ResourceUse, error)
	GetObject() (datatypes.Scale_Policy_Trigger_ResourceUse, error)
}

type Scale_Policy_Trigger_ResourceUse_Watch interface {
	Entity

	CreateObject(templateObject datatypes.Scale_Policy_Trigger_ResourceUse_Watch) (datatypes.Scale_Policy_Trigger_ResourceUse_Watch, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Scale_Policy_Trigger_ResourceUse_Watch) (bool, error)
	GetAllPossibleAlgorithms() (string, error)
	GetAllPossibleMetrics() (string, error)
	GetAllPossibleOperators() (string, error)
	GetObject() (datatypes.Scale_Policy_Trigger_ResourceUse_Watch, error)
}

type Scale_Policy_Trigger_Type interface {
	Entity

	GetAllObjects() (datatypes.Scale_Policy_Trigger_Type, error)
	GetObject() (datatypes.Scale_Policy_Trigger_Type, error)
}

type Scale_Termination_Policy interface {
	Entity

	GetAllObjects() (datatypes.Scale_Termination_Policy, error)
	GetObject() (datatypes.Scale_Termination_Policy, error)
}

type Search interface {
	Entity

	AdvancedSearch(searchString string) (datatypes.Container_Search_Result, error)
	GetObjectTypes() (datatypes.Container_Search_ObjectType, error)
	Search(searchString string) (datatypes.Container_Search_Result, error)
}

type Security_Certificate interface {
	Entity

	CreateObject(templateObject datatypes.Security_Certificate) (datatypes.Security_Certificate, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Security_Certificate) (bool, error)
	FindByCommonName(commonName string) (datatypes.Security_Certificate, error)
	GetObject() (datatypes.Security_Certificate, error)
	GetPemFormat() (string, error)
}

type Security_Certificate_Entry interface {
	Entity
}

type Security_Certificate_Request interface {
	Entity

	CancelSslOrder() (bool, error)
	GetAdministratorEmailDomains(commonName string) (string, error)
	GetAdministratorEmailPrefixes() (string, error)
	GetObject() (datatypes.Security_Certificate_Request, error)
	GetPreviousOrderData() (datatypes.Container_Product_Order_Security_Certificate, error)
	GetSslCertificateRequests(accountId int) (datatypes.Security_Certificate_Request, error)
	ResendEmail(emailType string) (bool, error)
	ValidateCsr(csr string, validityMonths int, itemId int, serverType string) (bool, error)
}

type Security_Certificate_Request_ServerType interface {
	Entity

	GetAllObjects() (datatypes.Security_Certificate_Request, error)
	GetObject() (datatypes.Security_Certificate_Request_ServerType, error)
}

type Security_Certificate_Request_Status interface {
	Entity

	GetObject() (datatypes.Security_Certificate_Request_Status, error)
	GetSslRequestStatuses() (datatypes.Security_Certificate_Request_Status, error)
}

type Security_Directory_Service_Host_Xref_Hardware interface {
	Entity
}

type Security_SecureTransportCipher interface {
	Entity
}

type Security_SecureTransportProtocol interface {
	Entity
}

type Security_Ssh_Key interface {
	Entity

	CreateObject(templateObject datatypes.Security_Ssh_Key) (datatypes.Security_Ssh_Key, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Security_Ssh_Key) (bool, error)
	GetObject() (datatypes.Security_Ssh_Key, error)
}

type Service_Provider interface {
	Entity
}

type Software_AccountLicense interface {
	Entity

	GetAllObjects() (datatypes.Software_AccountLicense, error)
	GetObject() (datatypes.Software_AccountLicense, error)
}

type Software_Component interface {
	Entity

	GetLicenseFile() (string, error)
	GetObject() (datatypes.Software_Component, error)
	GetVendorSetUpConfiguration() (string, error)
}

type Software_Component_Analytics interface {
	Software_Component
}

type Software_Component_Analytics_Urchin interface {
	Software_Component_Analytics
}

type Software_Component_AntivirusSpyware interface {
	Software_Component

	GetObject() (datatypes.Software_Component_AntivirusSpyware, error)
	UpdateAntivirusSpywarePolicy(newPolicy string, enforce bool) (bool, error)
}

type Software_Component_AntivirusSpyware_Mcafee interface {
	Software_Component_AntivirusSpyware
}

type Software_Component_AntivirusSpyware_Mcafee_Epo_Version36 interface {
	Software_Component_AntivirusSpyware_Mcafee
}

type Software_Component_AntivirusSpyware_Mcafee_Epo_Version45 interface {
	Software_Component_AntivirusSpyware_Mcafee
}

type Software_Component_ControlPanel interface {
	Software_Component
}

type Software_Component_ControlPanel_Cpanel interface {
	Software_Component
}

type Software_Component_ControlPanel_Idera interface {
	Software_Component
}

type Software_Component_ControlPanel_Idera_ServerBackup interface {
	Software_Component_ControlPanel_Idera
}

type Software_Component_ControlPanel_Microsoft interface {
	Software_Component
}

type Software_Component_ControlPanel_Microsoft_WebPlatform interface {
	Software_Component_ControlPanel_Microsoft
}

type Software_Component_ControlPanel_Parallels interface {
	Software_Component
}

type Software_Component_ControlPanel_Parallels_Plesk interface {
	Software_Component_ControlPanel_Parallels
}

type Software_Component_ControlPanel_R1soft interface {
	Software_Component
}

type Software_Component_ControlPanel_R1soft_Cdp interface {
	Software_Component_ControlPanel_R1soft
}

type Software_Component_ControlPanel_R1soft_ServerBackup interface {
	Software_Component_ControlPanel_R1soft
}

type Software_Component_ControlPanel_Swsoft interface {
	Software_Component
}

type Software_Component_ControlPanel_WebhostAutomation interface {
	Software_Component
}

type Software_Component_HostIps interface {
	Software_Component

	GetCurrentHostIpsPolicies() (datatypes.Container_Software_Component_HostIps_Policy, error)
	GetObject() (datatypes.Software_Component_HostIps, error)
	UpdateHipsPolicies(newIpsMode string, newIpsProtection string, newFirewallMode string, newFirewallRuleset string, newApplicationMode string, newApplicationRuleset string, newEnforcementPolicy string) (bool, error)
}

type Software_Component_HostIps_Mcafee interface {
	Software_Component_HostIps
}

type Software_Component_HostIps_Mcafee_Epo_Version36_Hips interface {
	Software_Component_HostIps_Mcafee
}

type Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version6 interface {
	Software_Component_HostIps_Mcafee_Epo_Version36_Hips
}

type Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version7 interface {
	Software_Component_HostIps_Mcafee_Epo_Version36_Hips
}

type Software_Component_HostIps_Mcafee_Epo_Version45_Hips interface {
	Software_Component_HostIps_Mcafee
}

type Software_Component_HostIps_Mcafee_Epo_Version45_Hips_Version7 interface {
	Software_Component_HostIps_Mcafee_Epo_Version45_Hips
}

type Software_Component_HostIps_Mcafee_Epo_Version45_Hips_Version8 interface {
	Software_Component_HostIps_Mcafee_Epo_Version45_Hips
}

type Software_Component_OperatingSystem interface {
	Software_Component
}

type Software_Component_Package interface {
	Software_Component
}

type Software_Component_Package_Management interface {
	Software_Component_Package
}

type Software_Component_Package_Management_Ksplice interface {
	Software_Component_Package_Management
}

type Software_Component_Password interface {
	Entity

	CreateObject(templateObject datatypes.Software_Component_Password) (datatypes.Software_Component_Password, error)
	CreateObjects(templateObjects datatypes.Software_Component_Password) (bool, error)
	DeleteObject() (bool, error)
	DeleteObjects(templateObjects datatypes.Software_Component_Password) (bool, error)
	EditObject(templateObject datatypes.Software_Component_Password) (bool, error)
	EditObjects(templateObjects datatypes.Software_Component_Password) (bool, error)
	GetObject() (datatypes.Software_Component_Password, error)
}

type Software_Component_Password_History interface {
	Entity
}

type Software_Component_Security interface {
	Software_Component
}

type Software_Component_Security_SafeNet interface {
	Software_Component_Security
}

type Software_Description interface {
	Entity

	GetAllObjects() (datatypes.Software_Description, error)
	GetObject() (datatypes.Software_Description, error)
}

type Software_Description_Attribute interface {
	Entity
}

type Software_Description_Attribute_Type interface {
	Entity
}

type Software_Description_Feature interface {
	Entity
}

type Software_Description_RequiredUser interface {
	Entity
}

type Software_License interface {
	Entity
}

type Software_VirtualLicense interface {
	Entity

	GetLicenseFile() ([]byte, error)
	GetObject() (datatypes.Software_VirtualLicense, error)
}

type Survey interface {
	Entity

	GetActiveSurveyByType(typ string) (datatypes.Survey, error)
	GetObject() (datatypes.Survey, error)
	TakeSurvey(responses datatypes.Survey_Response) (bool, error)
}

type Survey_Answer interface {
	Entity
}

type Survey_Question interface {
	Entity
}

type Survey_Response interface {
	Entity
}

type Survey_Status interface {
	Entity
}

type Survey_Type interface {
	Entity
}

type Tag interface {
	Entity

	AutoComplete(tag string) (datatypes.Tag, error)
	GetAllTagTypes() (datatypes.Tag_Type, error)
	GetObject() (datatypes.Tag, error)
	GetTagByTagName(tagList string) (datatypes.Tag, error)
	SetTags(tags string, keyName string, resourceTableId int) (bool, error)
}

type Tag_Reference interface {
	Entity
}

type Tag_Reference_Hardware interface {
	Tag_Reference
}

type Tag_Reference_Network_Application_Delivery_Controller interface {
	Tag_Reference
}

type Tag_Reference_Network_Vlan interface {
	Tag_Reference
}

type Tag_Reference_Network_Vlan_Firewall interface {
	Tag_Reference
}

type Tag_Reference_Resource_Group interface {
	Tag_Reference
}

type Tag_Reference_Virtual_Guest interface {
	Tag_Reference
}

type Tag_Reference_Virtual_Guest_Block_Device_Template_Group interface {
	Tag_Reference
}

type Tag_Type interface {
	Entity
}

type Ticket interface {
	Entity

	AddAssignedAgent(agentId int) error
	AddAttachedAdditionalEmails(emails string) (bool, error)
	AddAttachedFile(fileAttachment datatypes.Container_Utility_File_Attachment) (datatypes.Ticket_Attachment_File, error)
	AddAttachedHardware(hardwareId int) (datatypes.Ticket_Attachment_Hardware, error)
	AddAttachedVirtualGuest(guestId int) (datatypes.Ticket_Attachment_Virtual_Guest, error)
	AddFinalComments(finalComments string) (bool, error)
	AddScheduledAlert(activationTime string) error
	AddScheduledAutoClose(activationTime string) error
	AddUpdate(templateObject datatypes.Ticket_Update, attachedFiles datatypes.Container_Utility_File_Attachment) (datatypes.Ticket_Update, error)
	CreateAdministrativeTicket(templateObject datatypes.Ticket, contents string, attachmentId int, rootPassword string, controlPanelPassword string, accessPort string, attachedFiles datatypes.Container_Utility_File_Attachment, attachmentType string) (datatypes.Ticket, error)
	CreateCancelServerTicket(attachmentId int, reason string, content string, cancelAssociatedItems bool, attachmentType string) (datatypes.Ticket, error)
	CreateCancelServiceTicket(attachmentId int, reason string, content string, attachmentType string) (datatypes.Ticket, error)
	CreateStandardTicket(templateObject datatypes.Ticket, contents string, attachmentId int, rootPassword string, controlPanelPassword string, accessPort string, attachedFiles datatypes.Container_Utility_File_Attachment, attachmentType string) (datatypes.Ticket, error)
	CreateUpgradeTicket(attachmentId int, genericUpgrade string, upgradeMaintenanceWindow string, details string, attachmentType string) (datatypes.Ticket, error)
	Edit(templateObject datatypes.Ticket, contents string, attachedFiles datatypes.Container_Utility_File_Attachment) (datatypes.Ticket, error)
	GetAllTicketGroups() (datatypes.Ticket_Group, error)
	GetAllTicketStatuses() (datatypes.Ticket_Status, error)
	GetAttachedFile(attachmentId int) ([]byte, error)
	GetObject() (datatypes.Ticket, error)
	GetTicketsClosedSinceDate(closeDate time.Time) (datatypes.Ticket, error)
	MarkAsViewed() error
	RemoveAssignedAgent(agentId int) error
	RemoveAttachedAdditionalEmails(emails string) (bool, error)
	RemoveAttachedHardware(hardwareId int) (bool, error)
	RemoveAttachedVirtualGuest(guestId int) (bool, error)
	RemoveScheduledAlert() error
	RemoveScheduledAutoClose() error
	SetTags(tags string) (bool, error)
	SurveyEligible() (bool, error)
	UpdateAttachedAdditionalEmails(emails string) (bool, error)
}

type Ticket_Activity interface {
	Entity
}

type Ticket_Attachment interface {
	Entity
}

type Ticket_Attachment_Assigned_Agent interface {
	Ticket_Attachment
}

type Ticket_Attachment_CardChangeRequest interface {
	Ticket_Attachment
}

type Ticket_Attachment_File interface {
	Entity

	GetExtensionWhitelist() (string, error)
	GetObject() (datatypes.Ticket_Attachment_File, error)
}

type Ticket_Attachment_Hardware interface {
	Ticket_Attachment
}

type Ticket_Attachment_ManualPayment interface {
	Ticket_Attachment
}

type Ticket_Attachment_Scheduled_Action interface {
	Ticket_Attachment
}

type Ticket_Attachment_Virtual_Guest interface {
	Ticket_Attachment
}

type Ticket_Chat interface {
	Entity
}

type Ticket_Chat_Liveperson interface {
	Ticket_Chat
}

type Ticket_Chat_TranscriptLine interface {
	Entity
}

type Ticket_Chat_TranscriptLine_Customer interface {
	Ticket_Chat_TranscriptLine
}

type Ticket_Chat_TranscriptLine_Employee interface {
	Ticket_Chat_TranscriptLine
}

type Ticket_Group interface {
	Entity
}

type Ticket_Group_Category interface {
	Entity
}

type Ticket_Priority interface {
	Entity

	GetPriorities() (datatypes.Container_Ticket_Priority, error)
}

type Ticket_State interface {
	Entity
}

type Ticket_State_Type interface {
	Entity
}

type Ticket_Status interface {
	Entity
}

type Ticket_Subject interface {
	Entity

	GetAllObjects() (datatypes.Ticket_Subject, error)
	GetObject() (datatypes.Ticket_Subject, error)
	GetTopFiveKnowledgeLayerQuestions() (datatypes.Container_KnowledgeLayer_QuestionAnswer, error)
}

type Ticket_Survey interface {
	Entity

	GetPreference() (datatypes.Container_Ticket_Survey_Preference, error)
	OptIn() (datatypes.Container_Ticket_Survey_Preference, error)
	OptOut() (datatypes.Container_Ticket_Survey_Preference, error)
}

type Ticket_Type interface {
	Entity
}

type Ticket_Update interface {
	Entity
}

type Ticket_Update_Agent interface {
	Ticket_Update
}

type Ticket_Update_Chat interface {
	Ticket_Update
}

type Ticket_Update_Customer interface {
	Ticket_Update
}

type Ticket_Update_Employee interface {
	Ticket_Update

	AddResponseRating(responseRating int) (bool, error)
	GetObject() (datatypes.Ticket_Update_Employee, error)
}

type Ticket_Update_Type interface {
	Entity
}

type User_Access_Facility_Log interface {
	Entity
}

type User_Access_Facility_Log_Type interface {
	Entity
}

type User_Access_Facility_Visitor interface {
	Entity
}

type User_Access_Facility_Visitor_Type interface {
	Entity
}

type User_Customer interface {
	User_Interface

	AcknowledgeSupportPolicy() error
	AddApiAuthenticationKey() (string, error)
	AddBulkHardwareAccess(hardwareIds int) (bool, error)
	AddBulkPortalPermission(permissions datatypes.User_Customer_CustomerPermission_Permission) (bool, error)
	AddBulkRoles(roles datatypes.User_Permission_Role) error
	AddBulkVirtualGuestAccess(virtualGuestIds int) (bool, error)
	AddExternalBinding(externalBinding datatypes.User_External_Binding) (datatypes.User_Customer_External_Binding, error)
	AddHardwareAccess(hardwareId int) (bool, error)
	AddNotificationSubscriber(notificationKeyName string) (bool, error)
	AddPortalPermission(permission datatypes.User_Customer_CustomerPermission_Permission) (bool, error)
	AddRole(role datatypes.User_Permission_Role) error
	AddVirtualGuestAccess(virtualGuestId int) (bool, error)
	ChangePreference(preferenceTypeKeyName string, value string) (datatypes.User_Preference, error)
	CheckExternalAuthenticationStatus(authenticationContainer datatypes.Container_User_Customer_External_Binding) (datatypes.Container_User_Customer_Portal_Token, error)
	CheckPhoneFactorAuthenticationForPasswordSet(passwordSet datatypes.Container_User_Customer_PasswordSet, authenticationContainer datatypes.Container_User_Customer_External_Binding) (bool, error)
	CreateNotificationSubscriber(keyName string, resourceTableId int) (bool, error)
	CreateObject(templateObject datatypes.User_Customer, password string, vpnPassword string) (datatypes.User_Customer, error)
	CreateSubscriberDeliveryMethods(notificationKeyName string, deliveryMethodKeyNames string) (bool, error)
	DeactivateNotificationSubscriber(keyName string, resourceTableId int) (bool, error)
	EditObject(templateObject datatypes.User_Customer) (bool, error)
	EditObjects(templateObjects datatypes.User_Customer) (bool, error)
	FindUserPreference(profileName string, containerKeyname string, preferenceKeyname string) (datatypes.Layout_Profile, error)
	GetActiveExternalAuthenticationVendors() (datatypes.Container_User_Customer_External_Binding_Vendor, error)
	GetAllowedHardwareIds() (int, error)
	GetAllowedVirtualGuestIds() (int, error)
	GetAuthenticationToken(token datatypes.Container_User_Authentication_Token) (datatypes.Container_User_Authentication_Token, error)
	GetDefaultSecurityQuestions(key string) (datatypes.User_Security_Question, error)
	GetHardwareCount() (int, error)
	GetImpersonationToken() (string, error)
	GetObject() (datatypes.User_Customer, error)
	GetPortalLoginToken(username string, password string, securityQuestionId int, securityQuestionAnswer string) (datatypes.Container_User_Customer_Portal_Token, error)
	GetPreference(preferenceTypeKeyName string) (datatypes.User_Preference, error)
	GetPreferenceTypes() (datatypes.User_Preference_Type, error)
	GetRequirementsForPasswordSet(passwordSet datatypes.Container_User_Customer_PasswordSet) (datatypes.Container_User_Customer_PasswordSet, error)
	GetSupportPolicyDocument() ([]byte, error)
	GetSupportPolicyName() (string, error)
	GetSupportedLocales() (datatypes.Locale, error)
	GetUserFromLostPasswordRequest(key string) (datatypes.User_Security_Question, error)
	GetUserIdForPasswordSet(key string) (int, error)
	GetUserPreferences(profileName string, containerKeyname string) (datatypes.Layout_Profile, error)
	GetVirtualGuestCount() (int, error)
	InTerminalStatus() (bool, error)
	InitiateExternalAuthentication(authenticationContainer datatypes.Container_User_Customer_External_Binding) (string, error)
	InitiatePortalPasswordChange(username string) (bool, error)
	InitiatePortalPasswordChangeByBrandAgent(username string) (bool, error)
	InviteUserToLinkOpenIdConnect(providerType string) (bool, error)
	IsMasterUser() (bool, error)
	IsValidForumPassword(password string) (bool, error)
	IsValidPortalPassword(password string) (bool, error)
	LostPassword(username string, email string) (bool, error)
	PerformExternalAuthentication(authenticationContainer datatypes.Container_User_Customer_External_Binding) (datatypes.Container_User_Customer_Portal_Token, error)
	ProcessPasswordSetRequest(passwordSet datatypes.Container_User_Customer_PasswordSet, authenticationContainer datatypes.Container_User_Customer_External_Binding) (datatypes.Container_User_Customer_PasswordSet, error)
	RemoveAllHardwareAccessForThisUser() (bool, error)
	RemoveAllVirtualAccessForThisUser() (bool, error)
	RemoveApiAuthenticationKey(keyId int) (bool, error)
	RemoveBulkHardwareAccess(hardwareIds int) (bool, error)
	RemoveBulkPortalPermission(permissions datatypes.User_Customer_CustomerPermission_Permission) (bool, error)
	RemoveBulkRoles(roles datatypes.User_Permission_Role) error
	RemoveBulkVirtualGuestAccess(virtualGuestIds int) (bool, error)
	RemoveExternalBinding(externalBinding datatypes.User_External_Binding) (bool, error)
	RemoveHardwareAccess(hardwareId int) (bool, error)
	RemovePortalPermission(permission datatypes.User_Customer_CustomerPermission_Permission) (bool, error)
	RemoveRole(role datatypes.User_Permission_Role) error
	RemoveVirtualGuestAccess(virtualGuestId int) (bool, error)
	ResetExpiredPassword(username string, password string, newPassword string, securityQuestionId int, securityQuestionAnswer string) (bool, error)
	SamlAuthenticate(accountId string, samlResponse string) (datatypes.Container_User_Customer_Portal_Token, error)
	SamlBeginAuthentication(accountId int) (string, error)
	SamlBeginLogout() (string, error)
	SamlLogout(samlResponse string) error
	SetPasswordFromLostPasswordRequest(key string, password string, securityAnswers datatypes.User_Customer_Security_Answer) (bool, error)
	UpdateForumPassword(password string) (bool, error)
	UpdateNotificationSubscriber(notificationKeyName string, active int) (bool, error)
	UpdatePassword(password string) (bool, error)
	UpdateSecurityAnswers(questions datatypes.User_Security_Question, answers string) (bool, error)
	UpdateSubscriberDeliveryMethod(notificationKeyName string, deliveryMethodKeyNames string, active int) (bool, error)
	UpdateVpnPassword(password string) (bool, error)
	UpdateVpnUser() (bool, error)
	ValidateAuthenticationToken(authenticationToken datatypes.Container_User_Authentication_Token) (datatypes.Container_User_Customer_Portal_Token, error)
}

type User_Customer_Access_Authentication interface {
	Entity
}

type User_Customer_AdditionalEmail interface {
	Entity
}

type User_Customer_ApiAuthentication interface {
	Entity

	EditObject(templateObject datatypes.User_Customer_ApiAuthentication) (datatypes.User_Customer_ApiAuthentication, error)
	GetObject() (datatypes.User_Customer_ApiAuthentication, error)
}

type User_Customer_CustomerPermission_Permission interface {
	Entity

	GetAllObjects() (datatypes.User_Customer_CustomerPermission_Permission, error)
	GetObject() (datatypes.User_Customer_CustomerPermission_Permission, error)
}

type User_Customer_External_Binding interface {
	User_External_Binding

	Disable(reason string) (bool, error)
	Enable() (bool, error)
	GetObject() (datatypes.User_Customer_External_Binding, error)
}

type User_Customer_External_Binding_Attribute interface {
	User_External_Binding_Attribute
}

type User_Customer_External_Binding_Phone interface {
	User_Customer_External_Binding

	CheckPhoneValidationResult(token string) (bool, error)
	Disable(reason string) (bool, error)
	Enable() (bool, error)
	GetAllAuthenticationModes() (string, error)
	GetAllAuthenticationPinModes(authenticationModeKeyName string) (string, error)
	GetAuthenticationMode() (datatypes.Container_User_Customer_External_Binding_Phone_Mode, error)
	GetObject() (datatypes.User_Customer_External_Binding_Phone, error)
	GetPhoneAppActivationCode() (datatypes.User_External_Binding_Attribute, error)
	GetPhoneData() (datatypes.Container_User_Data_Phone, error)
	RequestPhoneValidation(phoneData datatypes.Container_User_Data_Phone) (string, error)
	UpdateAuthenticationMode(mode datatypes.Container_User_Customer_External_Binding_Phone_Mode) (bool, error)
	UpdatePhone(phoneData datatypes.Container_User_Data_Phone) (bool, error)
}

type User_Customer_External_Binding_Totp interface {
	User_Customer_External_Binding

	Activate() (bool, error)
	Deactivate() (bool, error)
	GenerateSecretKey() (string, error)
	GetObject() (datatypes.User_Customer_External_Binding_Totp, error)
}

type User_Customer_External_Binding_Type interface {
	User_External_Binding_Type
}

type User_Customer_External_Binding_Vendor interface {
	User_External_Binding_Vendor

	GetObject() (datatypes.User_Customer_External_Binding_Vendor, error)
}

type User_Customer_External_Binding_Verisign interface {
	User_Customer_External_Binding

	DeleteObject() (bool, error)
	Disable(reason string) (bool, error)
	Enable() (bool, error)
	GetActivationCodeForMobileClient() (string, error)
	GetObject() (datatypes.User_Customer_External_Binding_Verisign, error)
	Unlock(securityCode string) (bool, error)
	ValidateCredentialId(userId int, externalId string) (bool, error)
}

type User_Customer_Invitation interface {
	Entity

	GetObject() (datatypes.User_Customer_Invitation, error)
}

type User_Customer_Link interface {
	Entity
}

type User_Customer_Link_ThePlanet interface {
	User_Customer_Link
}

type User_Customer_MobileDevice interface {
	Entity

	CreateObject(templateObject datatypes.User_Customer_MobileDevice) (datatypes.User_Customer_MobileDevice, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.User_Customer_MobileDevice) (bool, error)
	GetObject() (datatypes.User_Customer_MobileDevice, error)
}

type User_Customer_MobileDevice_OperatingSystem interface {
	Entity

	GetAllObjects() (datatypes.User_Customer_MobileDevice_OperatingSystem, error)
	GetObject() (datatypes.User_Customer_MobileDevice_OperatingSystem, error)
}

type User_Customer_MobileDevice_Type interface {
	Entity

	GetAllObjects() (datatypes.User_Customer_MobileDevice_Type, error)
	GetObject() (datatypes.User_Customer_MobileDevice_Type, error)
}

type User_Customer_Notification_Hardware interface {
	Entity

	CreateObject(templateObject datatypes.User_Customer_Notification_Hardware) (datatypes.User_Customer_Notification_Hardware, error)
	CreateObjects(templateObjects datatypes.User_Customer_Notification_Hardware) (datatypes.Dns_Domain, error)
	DeleteObjects(templateObjects datatypes.User_Customer_Notification_Hardware) (bool, error)
	FindByHardwareId(hardwareId int) (datatypes.User_Customer_Notification_Hardware, error)
	GetObject() (datatypes.User_Customer_Notification_Hardware, error)
}

type User_Customer_Notification_Virtual_Guest interface {
	Entity

	CreateObject(templateObject datatypes.User_Customer_Notification_Virtual_Guest) (datatypes.User_Customer_Notification_Virtual_Guest, error)
	CreateObjects(templateObjects datatypes.User_Customer_Notification_Virtual_Guest) (datatypes.User_Customer_Notification_Virtual_Guest, error)
	DeleteObjects(templateObjects datatypes.User_Customer_Notification_Virtual_Guest) (bool, error)
	FindByGuestId(id int) (datatypes.User_Customer_Notification_Virtual_Guest, error)
	GetObject() (datatypes.User_Customer_Notification_Virtual_Guest, error)
}

type User_Customer_OpenIdConnect interface {
	User_Customer

	CompleteInvitationAfterLogin(providerType string, accessToken string, emailRegistrationCode string) error
	CreateOpenIdConnectUserAndCompleteInvitation(providerType string, user datatypes.User_Customer, password string, registrationCode string) (string, error)
	DeclineInvitation(providerType string, registrationCode string) error
	GetDefaultAccount(providerType string) (datatypes.Account, error)
	GetMappedAccounts(providerType string) (datatypes.Account, error)
	GetObject() (datatypes.User_Customer_OpenIdConnect, error)
	GetOpenIdRegistrationInfoFromCode(providerType string, registrationCode string) (datatypes.Account_Authentication_OpenIdConnect_RegistrationInformation, error)
	GetPortalLoginTokenOpenIdConnect(providerType string, accessToken string, accountId int) (datatypes.Container_User_Customer_Portal_Token, error)
	GetUserFromLostPasswordRequest(key string) (datatypes.User_Security_Question, error)
	IsValidPortalPassword(password string) (bool, error)
	LostPassword(username string, email string) (bool, error)
	ResetExpiredPassword(username string, password string, newPassword string, securityQuestionId int, securityQuestionAnswer string) (bool, error)
	SetDefaultAccount(providerType string, accountId int) (datatypes.Account, error)
	SetPasswordFromLostPasswordRequest(key string, password string, securityAnswers datatypes.User_Customer_Security_Answer) (bool, error)
	UpdatePassword(password string) (bool, error)
}

type User_Customer_Prospect interface {
	Entity
}

type User_Customer_Prospect_ServiceProvider_EnrollRequest interface {
	Entity

	Enroll(templateObject datatypes.User_Customer_Prospect_ServiceProvider_EnrollRequest) (datatypes.User_Customer_Prospect_ServiceProvider_EnrollRequest, error)
	GetObject() (datatypes.User_Customer_Prospect_ServiceProvider_EnrollRequest, error)
}

type User_Customer_Prospect_Type interface {
	Entity
}

type User_Customer_Security_Answer interface {
	Entity

	GetObject() (datatypes.User_Customer_Security_Answer, error)
}

type User_Customer_Status interface {
	Entity

	GetAllObjects() (datatypes.User_Customer_Status, error)
	GetObject() (datatypes.User_Customer_Status, error)
}

type User_Employee interface {
	User_Interface
}

type User_Employee_Department interface {
	Entity
}

type User_External_Binding interface {
	Entity

	DeleteObject() (bool, error)
	GetObject() (datatypes.User_External_Binding, error)
	UpdateNote(text string) (bool, error)
}

type User_External_Binding_Attribute interface {
	Entity
}

type User_External_Binding_Type interface {
	Entity
}

type User_External_Binding_Vendor interface {
	Entity

	GetAllObjects() (datatypes.User_External_Binding_Vendor, error)
	GetObject() (datatypes.User_External_Binding_Vendor, error)
}

type User_Interface interface {
	Entity
}

type User_Permission_Action interface {
	Entity

	GetAllObjects() (datatypes.User_Permission_Action, error)
	GetObject() (datatypes.User_Permission_Action, error)
}

type User_Permission_Group interface {
	Entity

	AddAction(action datatypes.User_Permission_Action) error
	AddBulkActions(actions datatypes.User_Permission_Action) error
	AddBulkResourceObjects(resourceObjects datatypes.Entity, resourceTypeKeyName string) (bool, error)
	AddResourceObject(resourceObject datatypes.Entity, resourceTypeKeyName string) (bool, error)
	CreateObject(templateObject datatypes.User_Permission_Group) (datatypes.User_Permission_Group, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.User_Permission_Group) (datatypes.User_Permission_Group, error)
	GetObject() (datatypes.User_Permission_Group, error)
	LinkRole(role datatypes.User_Permission_Role) error
	RemoveAction(action datatypes.User_Permission_Action) error
	RemoveBulkActions(actions datatypes.User_Permission_Action) error
	RemoveBulkResourceObjects(resourceObjects datatypes.Entity, resourceTypeKeyName string) (bool, error)
	RemoveResourceObject(resourceObject datatypes.Entity, resourceTypeKeyName string) (bool, error)
	UnlinkRole(role datatypes.User_Permission_Role) error
}

type User_Permission_Group_Type interface {
	Entity

	GetObject() (datatypes.User_Permission_Group_Type, error)
}

type User_Permission_Role interface {
	Entity

	AddUser(user datatypes.User_Customer) error
	CreateObject(templateObject datatypes.User_Permission_Role) (datatypes.User_Permission_Role, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.User_Permission_Role) (datatypes.User_Permission_Role, error)
	GetObject() (datatypes.User_Permission_Role, error)
	LinkGroup(group datatypes.User_Permission_Group) error
	RemoveUser(user datatypes.User_Customer) error
	UnlinkGroup(group datatypes.User_Permission_Group) error
}

type User_Preference interface {
	Entity
}

type User_Preference_Type interface {
	Entity
}

type User_Security_Question interface {
	Entity

	GetAllObjects() (datatypes.User_Security_Question, error)
	GetObject() (datatypes.User_Security_Question, error)
}

type Utility_Bandwidth_Graph interface {
	Entity
}

type Utility_Network interface {
	Entity

	NsLookup(address string, typ string) (string, error)
	Whois(address string) (string, error)
}

type Utility_ObjectFilter interface {
	Entity
}

type Utility_ObjectFilter_Operation interface {
	Entity
}

type Utility_ObjectFilter_Operation_Option interface {
	Entity
}

type Virtual_Disk_Image interface {
	Entity

	EditObject(templateObject datatypes.Virtual_Disk_Image) (bool, error)
	GetObject() (datatypes.Virtual_Disk_Image, error)
	GetPublicIsoImages() (datatypes.Virtual_Disk_Image, error)
}

type Virtual_Disk_Image_Software interface {
	Entity
}

type Virtual_Disk_Image_Software_Password interface {
	Entity
}

type Virtual_Disk_Image_Type interface {
	Entity
}

type Virtual_Guest interface {
	Entity

	ActivatePrivatePort() (bool, error)
	ActivatePublicPort() (bool, error)
	AllowAccessToNetworkStorage(networkStorageTemplateObject datatypes.Network_Storage) (bool, error)
	AllowAccessToNetworkStorageList(networkStorageTemplateObjects datatypes.Network_Storage) (bool, error)
	AttachDiskImage(imageId int) (datatypes.Provisioning_Version1_Transaction, error)
	CancelIsolationForDestructiveAction() error
	CaptureImage(captureTemplate datatypes.Container_Disk_Image_Capture_Template) (datatypes.Virtual_Guest_Block_Device_Template_Group, error)
	CheckHostDiskAvailability(diskCapacity int) (bool, error)
	CloseAlarm(alarmId string) (bool, error)
	ConfigureMetadataDisk() (datatypes.Provisioning_Version1_Transaction, error)
	CreateArchiveTransaction(groupName string, blockDevices datatypes.Virtual_Guest_Block_Device, note string) (datatypes.Provisioning_Version1_Transaction, error)
	CreateObject(templateObject datatypes.Virtual_Guest) (datatypes.Virtual_Guest, error)
	CreateObjects(templateObjects datatypes.Virtual_Guest) (datatypes.Virtual_Guest, error)
	CreatePostSoftwareInstallTransaction(data string, returnBoolean bool) (bool, error)
	DeleteObject() (bool, error)
	DetachDiskImage(imageId int) (datatypes.Provisioning_Version1_Transaction, error)
	EditObject(templateObject datatypes.Virtual_Guest) (bool, error)
	ExecuteIderaBareMetalRestore() (bool, error)
	ExecuteR1SoftBareMetalRestore() (bool, error)
	ExecuteRemoteScript(uri string) error
	ExecuteRescueLayer() (bool, error)
	FindByIpAddress(ipAddress string) (datatypes.Virtual_Guest, error)
	GenerateOrderTemplate(templateObject datatypes.Virtual_Guest) (datatypes.Container_Product_Order, error)
	GetAdditionalRequiredPricesForOsReload(config datatypes.Container_Hardware_Server_Configuration) (datatypes.Product_Item_Price, error)
	GetAlarmHistory(startDate time.Time, endDate time.Time, alarmId string) (datatypes.Container_Monitoring_Alarm_History, error)
	GetAttachedNetworkStorages(nasType string) (datatypes.Network_Storage, error)
	GetAvailableBlockDevicePositions() (string, error)
	GetAvailableNetworkStorages(nasType string) (datatypes.Network_Storage, error)
	GetBandwidthDataByDate(startDateTime time.Time, endDateTime time.Time, networkType string) (datatypes.Metric_Tracking_Object_Data, error)
	GetBandwidthForDateRange(startDate time.Time, endDate time.Time) (datatypes.Metric_Tracking_Object_Data, error)
	GetBandwidthImage(networkType string, snapshotRange string, dateSpecified time.Time, dateSpecifiedEnd time.Time) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetBandwidthImageByDate(startDateTime time.Time, endDateTime time.Time, networkType string) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetBandwidthTotal(startDateTime time.Time, endDateTime time.Time, direction string, side string) (uint, error)
	GetBootOrder() (string, error)
	GetConsoleAccessLog() (datatypes.Network_Logging_Syslog, error)
	GetCoreRestrictedOperatingSystemPrice() (datatypes.Product_Item_Price, error)
	GetCpuMetricDataByDate(startDateTime time.Time, endDateTime time.Time, cpuIndexes int) (datatypes.Metric_Tracking_Object_Data, error)
	GetCpuMetricImage(snapshotRange string, dateSpecified time.Time) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetCpuMetricImageByDate(startDateTime time.Time, endDateTime time.Time, cpuIndexes int) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetCreateObjectOptions() (datatypes.Container_Virtual_Guest_Configuration, error)
	GetCurrentBillingDetail() (datatypes.Billing_Item, error)
	GetCurrentBillingTotal() (float64, error)
	GetCustomBandwidthDataByDate(graphData datatypes.Container_Graph) (datatypes.Container_Graph, error)
	GetCustomMetricDataByDate(graphData datatypes.Container_Graph) (datatypes.Container_Graph, error)
	GetDriveRetentionItemPrice() (datatypes.Product_Item_Price, error)
	GetFirewallProtectableSubnets() (datatypes.Network_Subnet, error)
	GetFirstAvailableBlockDevicePosition() (string, error)
	GetIsoBootImage() (datatypes.Virtual_Disk_Image, error)
	GetItemPricesFromSoftwareDescriptions(softwareDescriptions datatypes.Software_Description, includeTranslationsFlag bool, returnAllPricesFlag bool) (datatypes.Product_Item, error)
	GetMemoryMetricDataByDate(startDateTime time.Time, endDateTime time.Time) (datatypes.Metric_Tracking_Object_Data, error)
	GetMemoryMetricImage(snapshotRange string, dateSpecified time.Time) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetMemoryMetricImageByDate(startDateTime time.Time, endDateTime time.Time) (datatypes.Container_Bandwidth_GraphOutputs, error)
	GetMonitoringActiveAlarms(startDate time.Time, endDate time.Time) (datatypes.Container_Monitoring_Alarm_History, error)
	GetMonitoringClosedAlarms(startDate time.Time, endDate time.Time) (datatypes.Container_Monitoring_Alarm_History, error)
	GetNetworkComponentFirewallProtectableIpAddresses() (datatypes.Network_Subnet_IpAddress, error)
	GetObject() (datatypes.Virtual_Guest, error)
	GetOrderTemplate(billingType string, orderPrices datatypes.Product_Item_Price) (datatypes.Container_Product_Order, error)
	GetProvisionDate() (time.Time, error)
	GetRecentMetricData(time uint) (datatypes.Metric_Tracking_Object, error)
	GetRemoteMonitoringActiveAlarms(startDate time.Time, endDate time.Time) (datatypes.Container_Monitoring_Alarm_History, error)
	GetRemoteMonitoringClosedAlarms(startDate time.Time, endDate time.Time) (datatypes.Container_Monitoring_Alarm_History, error)
	GetReverseDomainRecords() (datatypes.Dns_Domain, error)
	GetUpgradeItemPrices(includeDowngradeItemPrices bool) (datatypes.Product_Item_Price, error)
	GetValidBlockDeviceTemplateGroups(visibility string) (datatypes.Virtual_Guest_Block_Device_Template_Group, error)
	IsBackendPingable() (bool, error)
	IsPingable() (bool, error)
	IsolateInstanceForDestructiveAction() error
	MountIsoImage(diskImageId int) (datatypes.Provisioning_Version1_Transaction, error)
	Pause() (bool, error)
	PowerCycle() (bool, error)
	PowerOff() (bool, error)
	PowerOffSoft() (bool, error)
	PowerOn() (bool, error)
	RebootDefault() (bool, error)
	RebootHard() (bool, error)
	RebootSoft() (bool, error)
	ReloadCurrentOperatingSystemConfiguration() (datatypes.Provisioning_Version1_Transaction, error)
	ReloadOperatingSystem(token string, config datatypes.Container_Hardware_Server_Configuration) (string, error)
	RemoveAccessToNetworkStorage(networkStorageTemplateObject datatypes.Network_Storage) (bool, error)
	RemoveAccessToNetworkStorageList(networkStorageTemplateObjects datatypes.Network_Storage) (bool, error)
	Resume() (bool, error)
	SetPrivateNetworkInterfaceSpeed(newSpeed int) (bool, error)
	SetPublicNetworkInterfaceSpeed(newSpeed int) (bool, error)
	SetTags(tags string) (bool, error)
	SetUserMetadata(metadata string) (bool, error)
	ShutdownPrivatePort() (bool, error)
	ShutdownPublicPort() (bool, error)
	UnmountIsoImage() (datatypes.Provisioning_Version1_Transaction, error)
	ValidateImageTemplate(imageTemplateId int) (bool, error)
	VerifyReloadOperatingSystem(config datatypes.Container_Hardware_Server_Configuration) (bool, error)
}

type Virtual_Guest_Attribute interface {
	Entity
}

type Virtual_Guest_Attribute_Type interface {
	Entity
}

type Virtual_Guest_Attribute_UserData interface {
	Virtual_Guest_Attribute
}

type Virtual_Guest_Block_Device interface {
	Entity
}

type Virtual_Guest_Block_Device_Status interface {
	Entity
}

type Virtual_Guest_Block_Device_Template interface {
	Entity
}

type Virtual_Guest_Block_Device_Template_Group interface {
	Entity

	AddLocations(locations datatypes.Location) (bool, error)
	CopyToExternalSource(configuration datatypes.Container_Virtual_Guest_Block_Device_Template_Configuration) (bool, error)
	CreateFromExternalSource(configuration datatypes.Container_Virtual_Guest_Block_Device_Template_Configuration) (datatypes.Virtual_Guest_Block_Device_Template_Group, error)
	CreatePublicArchiveTransaction(groupName string, summary string, note string, locations datatypes.Location) (int, error)
	DeleteObject() (datatypes.Provisioning_Version1_Transaction, error)
	DenySharingAccess(accountId int) (bool, error)
	EditObject(templateObject datatypes.Virtual_Guest_Block_Device_Template_Group) (bool, error)
	GetObject() (datatypes.Virtual_Guest_Block_Device_Template_Group, error)
	GetPublicCustomerOwnedImages() (datatypes.Virtual_Guest_Block_Device_Template_Group, error)
	GetPublicImages() (datatypes.Virtual_Guest_Block_Device_Template_Group, error)
	GetStorageLocations() (datatypes.Location, error)
	GetVhdImportSoftwareDescriptions() (datatypes.Software_Description, error)
	PermitSharingAccess(accountId int) (bool, error)
	RemoveLocations(locations datatypes.Location) (bool, error)
	SetAvailableLocations(locations datatypes.Location) (bool, error)
	SetTags(tags string) (bool, error)
}

type Virtual_Guest_Block_Device_Template_Group_Accounts interface {
	Entity
}

type Virtual_Guest_Block_Device_Template_Group_Status interface {
	Entity
}

type Virtual_Guest_Boot_Parameter interface {
	Entity

	CreateObject(templateObject datatypes.Virtual_Guest_Boot_Parameter) (bool, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Virtual_Guest_Boot_Parameter) (bool, error)
	GetObject() (datatypes.Virtual_Guest_Boot_Parameter, error)
}

type Virtual_Guest_Boot_Parameter_Type interface {
	Entity

	GetAllObjects() (datatypes.Virtual_Guest_Boot_Parameter_Type, error)
	GetObject() (datatypes.Virtual_Guest_Boot_Parameter_Type, error)
}

type Virtual_Guest_Network_Component interface {
	Entity

	Disable() (bool, error)
	Enable() (bool, error)
	GetObject() (datatypes.Virtual_Guest_Network_Component, error)
	IsPingable() (bool, error)
}

type Virtual_Guest_Network_Component_IpAddress interface {
	Entity
}

type Virtual_Guest_Power_State interface {
	Entity
}

type Virtual_Guest_Status interface {
	Entity
}

type Virtual_Guest_SupplementalCreateObjectOptions interface {
	Entity
}

type Virtual_Host interface {
	Entity

	GetLiveGuestByUuid(uuid string) (datatypes.Virtual_Guest, error)
	GetLiveGuestList() (datatypes.Virtual_Guest, error)
	GetLiveGuestRecentMetricData(uuid string, time int, limit int, interval int) (datatypes.Metric_Tracking_Object, error)
	GetObject() (datatypes.Virtual_Host, error)
	PauseLiveGuest(uuid string) (bool, error)
	PowerCycleLiveGuest(uuid string) (bool, error)
	PowerOffLiveGuest(uuid string) (bool, error)
	PowerOnLiveGuest(uuid string) (bool, error)
	RebootSoftLiveGuest(uuid string) (bool, error)
	ResumeLiveGuest(uuid string) (bool, error)
}

type Virtual_Storage_Repository interface {
	Entity

	GetArchiveDiskUsageRatePerGb() (float64, error)
	GetAverageUsageMetricDataByDate(startDateTime time.Time, endDateTime time.Time) (float64, error)
	GetObject() (datatypes.Virtual_Storage_Repository, error)
	GetPublicImageDiskUsageRatePerGb() (float64, error)
	GetStorageLocations() (datatypes.Location, error)
	GetUsageMetricDataByDate(startDateTime time.Time, endDateTime time.Time) (datatypes.Metric_Tracking_Object_Data, error)
	GetUsageMetricImageByDate(startDateTime time.Time, endDateTime time.Time) (datatypes.Container_Bandwidth_GraphOutputs, error)
}

type Virtual_Storage_Repository_Type interface {
	Entity
}
