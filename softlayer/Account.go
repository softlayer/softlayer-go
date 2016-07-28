package softlayer

import (
	"time"

	"github.ibm.com/riethm/gopherlayer/datatypes"
)

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

	GetAbuseEmail() (string, error)
	GetAbuseEmails() (datatypes.Account_AbuseEmail, error)
	GetAccountContacts() (datatypes.Account_Contact, error)
	GetAccountLicenses() (datatypes.Software_AccountLicense, error)
	GetAccountLinks() (datatypes.Account_Link, error)
	GetAccountStatus() (datatypes.Account_Status, error)
	GetActiveAccountDiscountBillingItem() (datatypes.Billing_Item, error)
	GetActiveAccountLicenses() (datatypes.Software_AccountLicense, error)
	GetActiveAddresses() (datatypes.Account_Address, error)
	GetActiveBillingAgreements() (datatypes.Account_Agreement, error)
	GetActiveCatalystEnrollment() (datatypes.Catalyst_Enrollment, error)
	GetActiveColocationContainers() (datatypes.Billing_Item, error)
	GetActiveFlexibleCreditEnrollment() (datatypes.FlexibleCredit_Enrollment, error)
	GetActiveNotificationSubscribers() (datatypes.Notification_Subscriber, error)
	GetActiveQuotes() (datatypes.Billing_Order_Quote, error)
	GetActiveVirtualLicenses() (datatypes.Software_VirtualLicense, error)
	GetAdcLoadBalancers() (datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress, error)
	GetAddresses() (datatypes.Account_Address, error)
	GetAffiliateId() (string, error)
	GetAllBillingItems() (datatypes.Billing_Item, error)
	GetAllCommissionBillingItems() (datatypes.Billing_Item, error)
	GetAllRecurringTopLevelBillingItems() (datatypes.Billing_Item, error)
	GetAllRecurringTopLevelBillingItemsUnfiltered() (datatypes.Billing_Item, error)
	GetAllSubnetBillingItems() (datatypes.Billing_Item, error)
	GetAllTopLevelBillingItems() (datatypes.Billing_Item, error)
	GetAllTopLevelBillingItemsUnfiltered() (datatypes.Billing_Item, error)
	GetAllowsBluemixAccountLinkingFlag() (bool, error)
	GetApplicationDeliveryControllers() (datatypes.Network_Application_Delivery_Controller, error)
	GetAttributes() (datatypes.Account_Attribute, error)
	GetAvailablePublicNetworkVlans() (datatypes.Network_Vlan, error)
	GetBalance() (float64, error)
	GetBandwidthAllotments() (datatypes.Network_Bandwidth_Version1_Allotment, error)
	GetBandwidthAllotmentsOverAllocation() (datatypes.Network_Bandwidth_Version1_Allotment, error)
	GetBandwidthAllotmentsProjectedOverAllocation() (datatypes.Network_Bandwidth_Version1_Allotment, error)
	GetBareMetalInstances() (datatypes.Hardware, error)
	GetBillingAgreements() (datatypes.Account_Agreement, error)
	GetBillingInfo() (datatypes.Billing_Info, error)
	GetBlockDeviceTemplateGroups() (datatypes.Virtual_Guest_Block_Device_Template_Group, error)
	GetBlueIdAuthenticationRequiredFlag() (bool, error)
	GetBluemixLinkedFlag() (bool, error)
	GetBrand() (datatypes.Brand, error)
	GetBrandAccountFlag() (bool, error)
	GetBrandKeyName() (string, error)
	GetCanOrderAdditionalVlansFlag() (bool, error)
	GetCarts() (datatypes.Billing_Order_Quote, error)
	GetCatalystEnrollments() (datatypes.Catalyst_Enrollment, error)
	GetCdnAccounts() (datatypes.Network_ContentDelivery_Account, error)
	GetClosedTickets() (datatypes.Ticket, error)
	GetDatacentersWithSubnetAllocations() (datatypes.Location, error)
	GetDisablePaymentProcessingFlag() (bool, error)
	GetDisplaySupportRepresentativeAssignments() (datatypes.Account_Attachment_Employee, error)
	GetDomainRegistrations() (datatypes.Dns_Domain_Registration, error)
	GetDomains() (datatypes.Dns_Domain, error)
	GetDomainsWithoutSecondaryDnsRecords() (datatypes.Dns_Domain, error)
	GetEvaultCapacityGB() (uint, error)
	GetEvaultMasterUsers() (datatypes.Account_Password, error)
	GetEvaultNetworkStorage() (datatypes.Network_Storage, error)
	GetExpiredSecurityCertificates() (datatypes.Security_Certificate, error)
	GetFacilityLogs() (datatypes.User_Access_Facility_Log, error)
	GetFlexibleCreditEnrollments() (datatypes.FlexibleCredit_Enrollment, error)
	GetGlobalIpRecords() (datatypes.Network_Subnet_IpAddress_Global, error)
	GetGlobalIpv4Records() (datatypes.Network_Subnet_IpAddress_Global, error)
	GetGlobalIpv6Records() (datatypes.Network_Subnet_IpAddress_Global, error)
	GetGlobalLoadBalancerAccounts() (datatypes.Network_LoadBalancer_Global_Account, error)
	GetHardware() (datatypes.Hardware, error)
	GetHardwareOverBandwidthAllocation() (datatypes.Hardware, error)
	GetHardwareProjectedOverBandwidthAllocation() (datatypes.Hardware, error)
	GetHardwareWithCpanel() (datatypes.Hardware, error)
	GetHardwareWithHelm() (datatypes.Hardware, error)
	GetHardwareWithMcafee() (datatypes.Hardware, error)
	GetHardwareWithMcafeeAntivirusRedhat() (datatypes.Hardware, error)
	GetHardwareWithMcafeeAntivirusWindows() (datatypes.Hardware, error)
	GetHardwareWithMcafeeIntrusionDetectionSystem() (datatypes.Hardware, error)
	GetHardwareWithPlesk() (datatypes.Hardware, error)
	GetHardwareWithQuantastor() (datatypes.Hardware, error)
	GetHardwareWithUrchin() (datatypes.Hardware, error)
	GetHardwareWithWindows() (datatypes.Hardware, error)
	GetHasEvaultBareMetalRestorePluginFlag() (bool, error)
	GetHasIderaBareMetalRestorePluginFlag() (bool, error)
	GetHasPendingOrder() (uint, error)
	GetHasR1softBareMetalRestorePluginFlag() (bool, error)
	GetHourlyBareMetalInstances() (datatypes.Hardware, error)
	GetHourlyServiceBillingItems() (datatypes.Billing_Item, error)
	GetHourlyVirtualGuests() (datatypes.Virtual_Guest, error)
	GetHubNetworkStorage() (datatypes.Network_Storage, error)
	GetInternalNotes() (datatypes.Account_Note, error)
	GetInvoices() (datatypes.Billing_Invoice, error)
	GetIpAddresses() (datatypes.Network_Subnet_IpAddress, error)
	GetIscsiNetworkStorage() (datatypes.Network_Storage, error)
	GetLastCanceledBillingItem() (datatypes.Billing_Item, error)
	GetLastCancelledServerBillingItem() (datatypes.Billing_Item, error)
	GetLastFiveClosedAbuseTickets() (datatypes.Ticket, error)
	GetLastFiveClosedAccountingTickets() (datatypes.Ticket, error)
	GetLastFiveClosedOtherTickets() (datatypes.Ticket, error)
	GetLastFiveClosedSalesTickets() (datatypes.Ticket, error)
	GetLastFiveClosedSupportTickets() (datatypes.Ticket, error)
	GetLastFiveClosedTickets() (datatypes.Ticket, error)
	GetLatestBillDate() (time.Time, error)
	GetLatestRecurringInvoice() (datatypes.Billing_Invoice, error)
	GetLatestRecurringPendingInvoice() (datatypes.Billing_Invoice, error)
	GetLegacyBandwidthAllotments() (datatypes.Network_Bandwidth_Version1_Allotment, error)
	GetLegacyIscsiCapacityGB() (uint, error)
	GetLoadBalancers() (datatypes.Network_LoadBalancer_VirtualIpAddress, error)
	GetLockboxCapacityGB() (uint, error)
	GetLockboxNetworkStorage() (datatypes.Network_Storage, error)
	GetManualPaymentsUnderReview() (datatypes.Billing_Payment_Card_ManualPayment, error)
	GetMasterUser() (datatypes.User_Customer, error)
	GetMediaDataTransferRequests() (datatypes.Account_Media_Data_Transfer_Request, error)
	GetMessageQueueAccounts() (datatypes.Network_Message_Queue, error)
	GetMonthlyBareMetalInstances() (datatypes.Hardware, error)
	GetMonthlyVirtualGuests() (datatypes.Virtual_Guest, error)
	GetNasNetworkStorage() (datatypes.Network_Storage, error)
	GetNetworkCreationFlag() (bool, error)
	GetNetworkGateways() (datatypes.Network_Gateway, error)
	GetNetworkHardware() (datatypes.Hardware, error)
	GetNetworkMessageDeliveryAccounts() (datatypes.Network_Message_Delivery, error)
	GetNetworkMonitorDownHardware() (datatypes.Hardware, error)
	GetNetworkMonitorDownVirtualGuests() (datatypes.Virtual_Guest, error)
	GetNetworkMonitorRecoveringHardware() (datatypes.Hardware, error)
	GetNetworkMonitorRecoveringVirtualGuests() (datatypes.Virtual_Guest, error)
	GetNetworkMonitorUpHardware() (datatypes.Hardware, error)
	GetNetworkMonitorUpVirtualGuests() (datatypes.Virtual_Guest, error)
	GetNetworkStorage() (datatypes.Network_Storage, error)
	GetNetworkStorageGroups() (datatypes.Network_Storage_Group, error)
	GetNetworkTunnelContexts() (datatypes.Network_Tunnel_Module_Context, error)
	GetNetworkVlanSpan() (datatypes.Account_Network_Vlan_Span, error)
	GetNetworkVlans() (datatypes.Network_Vlan, error)
	GetNextBillingPublicAllotmentHardwareBandwidthDetails() (datatypes.Network_Bandwidth_Version1_Allotment, error)
	GetNextInvoiceIncubatorExemptTotal() (float64, error)
	GetNextInvoiceTopLevelBillingItems() (datatypes.Billing_Item, error)
	GetNextInvoiceTotalAmount() (float64, error)
	GetNextInvoiceTotalOneTimeAmount() (float64, error)
	GetNextInvoiceTotalOneTimeTaxAmount() (float64, error)
	GetNextInvoiceTotalRecurringAmount() (float64, error)
	GetNextInvoiceTotalRecurringAmountBeforeAccountDiscount() (float64, error)
	GetNextInvoiceTotalRecurringTaxAmount() (float64, error)
	GetNextInvoiceTotalTaxableRecurringAmount() (float64, error)
	GetNotificationSubscribers() (datatypes.Notification_Subscriber, error)
	GetOpenAbuseTickets() (datatypes.Ticket, error)
	GetOpenAccountingTickets() (datatypes.Ticket, error)
	GetOpenBillingTickets() (datatypes.Ticket, error)
	GetOpenCancellationRequests() (datatypes.Billing_Item_Cancellation_Request, error)
	GetOpenOtherTickets() (datatypes.Ticket, error)
	GetOpenRecurringInvoices() (datatypes.Billing_Invoice, error)
	GetOpenSalesTickets() (datatypes.Ticket, error)
	GetOpenStackAccountLinks() (datatypes.Account_Link, error)
	GetOpenStackObjectStorage() (datatypes.Network_Storage, error)
	GetOpenSupportTickets() (datatypes.Ticket, error)
	GetOpenTickets() (datatypes.Ticket, error)
	GetOpenTicketsWaitingOnCustomer() (datatypes.Ticket, error)
	GetOrders() (datatypes.Billing_Order, error)
	GetOrphanBillingItems() (datatypes.Billing_Item, error)
	GetOwnedBrands() (datatypes.Brand, error)
	GetOwnedHardwareGenericComponentModels() (datatypes.Hardware_Component_Model_Generic, error)
	GetPaymentProcessors() (datatypes.Billing_Payment_Processor, error)
	GetPendingEvents() (datatypes.Notification_Occurrence_Event, error)
	GetPendingInvoice() (datatypes.Billing_Invoice, error)
	GetPendingInvoiceTopLevelItems() (datatypes.Billing_Invoice_Item, error)
	GetPendingInvoiceTotalAmount() (float64, error)
	GetPendingInvoiceTotalOneTimeAmount() (float64, error)
	GetPendingInvoiceTotalOneTimeTaxAmount() (float64, error)
	GetPendingInvoiceTotalRecurringAmount() (float64, error)
	GetPendingInvoiceTotalRecurringTaxAmount() (float64, error)
	GetPermissionGroups() (datatypes.User_Permission_Group, error)
	GetPermissionRoles() (datatypes.User_Permission_Role, error)
	GetPortableStorageVolumes() (datatypes.Virtual_Disk_Image, error)
	GetPostProvisioningHooks() (datatypes.Provisioning_Hook, error)
	GetPptpVpnUsers() (datatypes.User_Customer, error)
	GetPreviousRecurringRevenue() (float64, error)
	GetPriceRestrictions() (datatypes.Product_Item_Price_Account_Restriction, error)
	GetPriorityOneTickets() (datatypes.Ticket, error)
	GetPrivateAllotmentHardwareBandwidthDetails() (datatypes.Network_Bandwidth_Version1_Allotment, error)
	GetPrivateBlockDeviceTemplateGroups() (datatypes.Virtual_Guest_Block_Device_Template_Group, error)
	GetPrivateIpAddresses() (datatypes.Network_Subnet_IpAddress, error)
	GetPrivateNetworkVlans() (datatypes.Network_Vlan, error)
	GetPrivateSubnets() (datatypes.Network_Subnet, error)
	GetPublicAllotmentHardwareBandwidthDetails() (datatypes.Network_Bandwidth_Version1_Allotment, error)
	GetPublicIpAddresses() (datatypes.Network_Subnet_IpAddress, error)
	GetPublicNetworkVlans() (datatypes.Network_Vlan, error)
	GetPublicSubnets() (datatypes.Network_Subnet, error)
	GetQuotes() (datatypes.Billing_Order_Quote, error)
	GetRecentEvents() (datatypes.Notification_Occurrence_Event, error)
	GetReferralPartner() (datatypes.Account, error)
	GetReferredAccounts() (datatypes.Account, error)
	GetRegulatedWorkloads() (datatypes.Legal_RegulatedWorkload, error)
	GetRemoteManagementCommandRequests() (datatypes.Hardware_Component_RemoteManagement_Command_Request, error)
	GetReplicationEvents() (datatypes.Network_Storage_Event, error)
	GetResourceGroups() (datatypes.Resource_Group, error)
	GetRouters() (datatypes.Hardware, error)
	GetRwhoisData() (datatypes.Network_Subnet_Rwhois_Data, error)
	GetSalesforceAccountLink() (datatypes.Account_Link, error)
	GetSamlAuthentication() (datatypes.Account_Authentication_Saml, error)
	GetScaleGroups() (datatypes.Scale_Group, error)
	GetSecondaryDomains() (datatypes.Dns_Secondary, error)
	GetSecurityCertificates() (datatypes.Security_Certificate, error)
	GetSecurityScanRequests() (datatypes.Network_Security_Scanner_Request, error)
	GetServiceBillingItems() (datatypes.Billing_Item, error)
	GetShipments() (datatypes.Account_Shipment, error)
	GetSshKeys() (datatypes.Security_Ssh_Key, error)
	GetSslVpnUsers() (datatypes.User_Customer, error)
	GetStandardPoolVirtualGuests() (datatypes.Virtual_Guest, error)
	GetSubnetRegistrationDetails() (datatypes.Account_Regional_Registry_Detail, error)
	GetSubnetRegistrations() (datatypes.Network_Subnet_Registration, error)
	GetSubnets() (datatypes.Network_Subnet, error)
	GetSupportRepresentatives() (datatypes.User_Employee, error)
	GetSupportSubscriptions() (datatypes.Billing_Item, error)
	GetSupportTier() (string, error)
	GetSuppressInvoicesFlag() (bool, error)
	GetTags() (datatypes.Tag, error)
	GetTickets() (datatypes.Ticket, error)
	GetTicketsClosedInTheLastThreeDays() (datatypes.Ticket, error)
	GetTicketsClosedToday() (datatypes.Ticket, error)
	GetTranscodeAccounts() (datatypes.Network_Media_Transcode_Account, error)
	GetUpgradeRequests() (datatypes.Product_Upgrade_Request, error)
	GetUsers() (datatypes.User_Customer, error)
	GetValidSecurityCertificates() (datatypes.Security_Certificate, error)
	GetVdrUpdatesInProgressFlag() (bool, error)
	GetVirtualDedicatedRacks() (datatypes.Network_Bandwidth_Version1_Allotment, error)
	GetVirtualDiskImages() (datatypes.Virtual_Disk_Image, error)
	GetVirtualGuests() (datatypes.Virtual_Guest, error)
	GetVirtualGuestsOverBandwidthAllocation() (datatypes.Virtual_Guest, error)
	GetVirtualGuestsProjectedOverBandwidthAllocation() (datatypes.Virtual_Guest, error)
	GetVirtualGuestsWithCpanel() (datatypes.Virtual_Guest, error)
	GetVirtualGuestsWithMcafee() (datatypes.Virtual_Guest, error)
	GetVirtualGuestsWithMcafeeAntivirusRedhat() (datatypes.Virtual_Guest, error)
	GetVirtualGuestsWithMcafeeAntivirusWindows() (datatypes.Virtual_Guest, error)
	GetVirtualGuestsWithMcafeeIntrusionDetectionSystem() (datatypes.Virtual_Guest, error)
	GetVirtualGuestsWithPlesk() (datatypes.Virtual_Guest, error)
	GetVirtualGuestsWithQuantastor() (datatypes.Virtual_Guest, error)
	GetVirtualGuestsWithUrchin() (datatypes.Virtual_Guest, error)
	GetVirtualPrivateRack() (datatypes.Network_Bandwidth_Version1_Allotment, error)
	GetVirtualStorageArchiveRepositories() (datatypes.Virtual_Storage_Repository, error)
	GetVirtualStoragePublicRepositories() (datatypes.Virtual_Storage_Repository, error)
}

type Account_AbuseEmail interface {
	Entity

	GetAccount() (datatypes.Account, error)
}

type Account_Address interface {
	Entity

	CreateObject(templateObject datatypes.Account_Address) (datatypes.Account_Address, error)
	EditObject(templateObject datatypes.Account_Address) (bool, error)
	GetAllDataCenters() (datatypes.Account_Address, error)
	GetNetworkAddress(name string) (datatypes.Account_Address, error)
	GetObject() (datatypes.Account_Address, error)

	GetAccount() (datatypes.Account, error)
	GetCreateUser() (datatypes.User_Customer, error)
	GetLocation() (datatypes.Location, error)
	GetModifyEmployee() (datatypes.User_Employee, error)
	GetModifyUser() (datatypes.User_Customer, error)
	GetType() (datatypes.Account_Address_Type, error)
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

	GetAccount() (datatypes.Account, error)
}

type Account_Agreement interface {
	Entity

	GetObject() (datatypes.Account_Agreement, error)

	GetAccount() (datatypes.Account, error)
	GetAgreementType() (datatypes.Account_Agreement_Type, error)
	GetAttachedBillingAgreementFiles() (datatypes.Account_MasterServiceAgreement, error)
	GetBillingItems() (datatypes.Billing_Item, error)
	GetStatus() (datatypes.Account_Agreement_Status, error)
	GetTopLevelBillingItems() (datatypes.Billing_Item, error)
}

type Account_Agreement_Status interface {
	Entity
}

type Account_Agreement_Type interface {
	Entity
}

type Account_Attachment_Employee interface {
	Entity

	GetAccount() (datatypes.Account, error)
	GetEmployee() (datatypes.User_Employee, error)
	GetEmployeeRole() (datatypes.Account_Attachment_Employee_Role, error)
}

type Account_Attachment_Employee_Role interface {
	Entity
}

type Account_Attribute interface {
	Entity

	GetAccount() (datatypes.Account, error)
	GetAccountAttributeType() (datatypes.Account_Attribute_Type, error)
}

type Account_Attribute_Type interface {
	Entity
}

type Account_Authentication_Attribute interface {
	Entity

	GetObject() (datatypes.Account_Authentication_Attribute, error)

	GetAccount() (datatypes.Account, error)
	GetAuthenticationRecord() (datatypes.Account_Authentication_Saml, error)
	GetType() (datatypes.Account_Authentication_Attribute_Type, error)
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

	GetAccount() (datatypes.Account, error)
	GetAttributes() (datatypes.Account_Authentication_Attribute, error)
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

	GetAccount() (datatypes.Account, error)
	GetType() (datatypes.Account_Contact_Type, error)
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

	GetAccount() (datatypes.Account, error)
	GetServiceProvider() (datatypes.Service_Provider, error)
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

	GetAccount() (datatypes.Account, error)
}

type Account_Media interface {
	Entity

	EditObject(templateObject datatypes.Account_Media) (bool, error)
	GetAllMediaTypes() (datatypes.Account_Media_Type, error)
	GetObject() (datatypes.Account_Media, error)
	RemoveMediaFromList(mediaTemplate datatypes.Account_Media) (int, error)

	GetAccount() (datatypes.Account, error)
	GetCreateUser() (datatypes.User_Customer, error)
	GetDatacenter() (datatypes.Location, error)
	GetModifyEmployee() (datatypes.User_Employee, error)
	GetModifyUser() (datatypes.User_Customer, error)
	GetRequest() (datatypes.Account_Media_Data_Transfer_Request, error)
	GetType() (datatypes.Account_Media_Type, error)
	GetVolume() (datatypes.Network_Storage, error)
}

type Account_Media_Data_Transfer_Request interface {
	Entity

	EditObject(templateObject datatypes.Account_Media_Data_Transfer_Request) (bool, error)
	GetAllRequestStatuses() (datatypes.Account_Media_Data_Transfer_Request_Status, error)
	GetObject() (datatypes.Account_Media_Data_Transfer_Request, error)

	GetAccount() (datatypes.Account, error)
	GetActiveTickets() (datatypes.Ticket, error)
	GetBillingItem() (datatypes.Billing_Item, error)
	GetCreateUser() (datatypes.User_Customer, error)
	GetMedia() (datatypes.Account_Media, error)
	GetModifyEmployee() (datatypes.User_Employee, error)
	GetModifyUser() (datatypes.User_Customer, error)
	GetShipments() (datatypes.Account_Shipment, error)
	GetStatus() (datatypes.Account_Media_Data_Transfer_Request_Status, error)
	GetTickets() (datatypes.Ticket, error)
}

type Account_Media_Data_Transfer_Request_Status interface {
	Entity
}

type Account_Media_Type interface {
	Entity
}

type Account_Network_Vlan_Span interface {
	Entity

	GetAccount() (datatypes.Account, error)
}

type Account_Note interface {
	Entity

	CreateObject(templateObject datatypes.Account_Note) (datatypes.Account_Note, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Account_Note) (bool, error)
	GetObject() (datatypes.Account_Note, error)

	GetAccount() (datatypes.Account, error)
	GetCustomer() (datatypes.User_Customer, error)
	GetNoteHistory() (datatypes.Account_Note_History, error)
	GetNoteType() (datatypes.Account_Note_Type, error)
}

type Account_Note_History interface {
	Entity

	GetAccountNote() (datatypes.Account_Note, error)
	GetCustomer() (datatypes.User_Customer, error)
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

	GetAccount() (datatypes.Account, error)
	GetType() (datatypes.Account_Password_Type, error)
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

	GetAccount() (datatypes.Account, error)
	GetDetailType() (datatypes.Account_Regional_Registry_Detail_Type, error)
	GetDetails() (datatypes.Network_Subnet_Registration_Details, error)
	GetProperties() (datatypes.Account_Regional_Registry_Detail_Property, error)
	GetRegionalInternetRegistryHandle() (datatypes.Account_Rwhois_Handle, error)
}

type Account_Regional_Registry_Detail_Property interface {
	Entity

	CreateObject(templateObject datatypes.Account_Regional_Registry_Detail_Property) (datatypes.Account_Regional_Registry_Detail_Property, error)
	CreateObjects(templateObjects datatypes.Account_Regional_Registry_Detail_Property) (datatypes.Account_Regional_Registry_Detail_Property, error)
	DeleteObject() (bool, error)
	EditObject(templateObject datatypes.Account_Regional_Registry_Detail_Property) (bool, error)
	EditObjects(templateObjects datatypes.Account_Regional_Registry_Detail_Property) (bool, error)
	GetObject() (datatypes.Account_Regional_Registry_Detail_Property, error)

	GetDetail() (datatypes.Account_Regional_Registry_Detail, error)
	GetPropertyType() (datatypes.Account_Regional_Registry_Detail_Property_Type, error)
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

	GetAccount() (datatypes.Account, error)
	GetAccountContact() (datatypes.Account_Contact, error)
	GetReportType() (datatypes.Compliance_Report_Type, error)
	GetTicket() (datatypes.Ticket, error)
	GetUser() (datatypes.User_Customer, error)
}

type Account_Rwhois_Handle interface {
	Entity

	GetAccount() (datatypes.Account, error)
}

type Account_Shipment interface {
	Entity

	EditObject(templateObject datatypes.Account_Shipment) (bool, error)
	GetAllCouriers() (datatypes.Auxiliary_Shipping_Courier, error)
	GetAllCouriersByType(courierTypeKeyName string) (datatypes.Auxiliary_Shipping_Courier, error)
	GetAllShipmentStatuses() (datatypes.Account_Shipment_Status, error)
	GetAllShipmentTypes() (datatypes.Account_Shipment_Type, error)
	GetObject() (datatypes.Account_Shipment, error)

	GetAccount() (datatypes.Account, error)
	GetCourier() (datatypes.Auxiliary_Shipping_Courier, error)
	GetCreateEmployee() (datatypes.User_Employee, error)
	GetCreateUser() (datatypes.User_Customer, error)
	GetDestinationAddress() (datatypes.Account_Address, error)
	GetModifyEmployee() (datatypes.User_Employee, error)
	GetModifyUser() (datatypes.User_Customer, error)
	GetOriginationAddress() (datatypes.Account_Address, error)
	GetShipmentItems() (datatypes.Account_Shipment_Item, error)
	GetStatus() (datatypes.Account_Shipment_Status, error)
	GetTrackingData() (datatypes.Account_Shipment_Tracking_Data, error)
	GetType() (datatypes.Account_Shipment_Type, error)
}

type Account_Shipment_Item interface {
	Entity

	EditObject(templateObject datatypes.Account_Shipment_Item) (bool, error)
	GetObject() (datatypes.Account_Shipment_Item, error)

	GetShipment() (datatypes.Account_Shipment, error)
	GetShipmentItemType() (datatypes.Account_Shipment_Item_Type, error)
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

	GetCreateEmployee() (datatypes.User_Employee, error)
	GetCreateUser() (datatypes.User_Customer, error)
	GetModifyEmployee() (datatypes.User_Employee, error)
	GetModifyUser() (datatypes.User_Customer, error)
	GetShipment() (datatypes.Account_Shipment, error)
}

type Account_Shipment_Type interface {
	Entity

	GetObject() (datatypes.Account_Shipment_Type, error)
}

type Account_Status interface {
	Entity
}
