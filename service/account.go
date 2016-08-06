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

type Account struct {
	Session *Session
	Options
}

func (r *Session) GetAccountService() Account {
	return Account{Session: r}
}

func (r *Account) ActivatePartner(accountId *string, hashCode *string) (resp datatypes.Account, err error) {
	params := []interface{}{
		accountId,
		hashCode,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) AddAchInformation(achInformation *datatypes.Container_Billing_Info_Ach) (resp bool, err error) {
	params := []interface{}{
		achInformation,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) AddReferralPartnerPaymentOption(paymentOption *datatypes.Container_Referral_Partner_Payment_Option) (resp bool, err error) {
	params := []interface{}{
		paymentOption,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) AreVdrUpdatesBlockedForBilling() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) CancelPayPalTransaction(token *string, payerId *string) (resp bool, err error) {
	params := []interface{}{
		token,
		payerId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) CompletePayPalTransaction(token *string, payerId *string) (resp string, err error) {
	params := []interface{}{
		token,
		payerId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) CountHourlyInstances() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAbuseEmail() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAbuseEmails() (resp []datatypes.Account_AbuseEmail, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAccountBackupHistory(startDate *datatypes.Time, endDate *datatypes.Time, backupStatus *string) (resp []datatypes.Container_Network_Storage_Evault_WebCc_JobDetails, err error) {
	params := []interface{}{
		startDate,
		endDate,
		backupStatus,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAccountContacts() (resp []datatypes.Account_Contact, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAccountLicenses() (resp []datatypes.Software_AccountLicense, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAccountLinks() (resp []datatypes.Account_Link, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAccountStatus() (resp datatypes.Account_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAccountTraitValue(keyName *string) (resp string, err error) {
	params := []interface{}{
		keyName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetActiveAccountDiscountBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetActiveAccountLicenses() (resp []datatypes.Software_AccountLicense, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetActiveAddresses() (resp []datatypes.Account_Address, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetActiveAlarms() (resp []datatypes.Container_Monitoring_Alarm_History, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetActiveBillingAgreements() (resp []datatypes.Account_Agreement, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetActiveCatalystEnrollment() (resp datatypes.Catalyst_Enrollment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetActiveColocationContainers() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetActiveFlexibleCreditEnrollment() (resp datatypes.FlexibleCredit_Enrollment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetActiveNotificationSubscribers() (resp []datatypes.Notification_Subscriber, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetActiveOutletPackages() (resp []datatypes.Product_Package, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetActivePackages() (resp []datatypes.Product_Package, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetActivePackagesByAttribute(attributeKeyName *string) (resp []datatypes.Product_Package, err error) {
	params := []interface{}{
		attributeKeyName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetActivePrivateHostedCloudPackages() (resp []datatypes.Product_Package, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetActiveQuotes() (resp []datatypes.Billing_Order_Quote, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetActiveVirtualLicenses() (resp []datatypes.Software_VirtualLicense, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAdcLoadBalancers() (resp []datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAddresses() (resp []datatypes.Account_Address, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAffiliateId() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAggregatedUptimeGraph(startDate *datatypes.Time, endDate *datatypes.Time) (resp datatypes.Container_Graph, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAllBillingItems() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAllCommissionBillingItems() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAllRecurringTopLevelBillingItems() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAllRecurringTopLevelBillingItemsUnfiltered() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAllSubnetBillingItems() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAllTopLevelBillingItems() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAllTopLevelBillingItemsUnfiltered() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAllowsBluemixAccountLinkingFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAlternateCreditCardData() (resp datatypes.Container_Account_Payment_Method_CreditCard, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetApplicationDeliveryControllers() (resp []datatypes.Network_Application_Delivery_Controller, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAttributeByType(attributeType *string) (resp datatypes.Account_Attribute, err error) {
	params := []interface{}{
		attributeType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAttributes() (resp []datatypes.Account_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAuxiliaryNotifications() (resp []datatypes.Container_Utility_Message, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAvailablePublicNetworkVlans() (resp []datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAverageArchiveUsageMetricDataByDate(startDateTime *datatypes.Time, endDateTime *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetAveragePublicUsageMetricDataByDate(startDateTime *datatypes.Time, endDateTime *datatypes.Time) (resp float64, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetBalance() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetBandwidthAllotments() (resp []datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetBandwidthAllotmentsOverAllocation() (resp []datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetBandwidthAllotmentsProjectedOverAllocation() (resp []datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetBareMetalInstances() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetBillingAgreements() (resp []datatypes.Account_Agreement, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetBillingInfo() (resp datatypes.Billing_Info, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetBlockDeviceTemplateGroups() (resp []datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetBlueIdAuthenticationRequiredFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetBluemixLinkedFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetBrand() (resp datatypes.Brand, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetBrandAccountFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetBrandKeyName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetCanOrderAdditionalVlansFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetCarts() (resp []datatypes.Billing_Order_Quote, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetCatalystEnrollments() (resp []datatypes.Catalyst_Enrollment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetCdnAccounts() (resp []datatypes.Network_ContentDelivery_Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetClosedTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetCurrentBackupStatisticsGraph(detailedGraph *bool) (resp datatypes.Container_Account_Graph_Outputs, err error) {
	params := []interface{}{
		detailedGraph,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetCurrentTicketStatisticsGraph(detailedGraph *bool) (resp datatypes.Container_Account_Graph_Outputs, err error) {
	params := []interface{}{
		detailedGraph,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetCurrentUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetDatacentersWithSubnetAllocations() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetDisablePaymentProcessingFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetDiskUsageMetricDataByDate(startDateTime *datatypes.Time, endDateTime *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetDiskUsageMetricDataFromLegacyByDate(startDateTime *datatypes.Time, endDateTime *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetDiskUsageMetricDataFromMetricTrackingObjectSystemByDate(startDateTime *datatypes.Time, endDateTime *datatypes.Time) (resp []datatypes.Metric_Tracking_Object_Data, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetDiskUsageMetricImageByDate(startDateTime *datatypes.Time, endDateTime *datatypes.Time) (resp datatypes.Container_Account_Graph_Outputs, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetDisplaySupportRepresentativeAssignments() (resp []datatypes.Account_Attachment_Employee, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetDomainRegistrations() (resp []datatypes.Dns_Domain_Registration, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetDomains() (resp []datatypes.Dns_Domain, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetDomainsWithoutSecondaryDnsRecords() (resp []datatypes.Dns_Domain, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetEvaultCapacityGB() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetEvaultMasterUsers() (resp []datatypes.Account_Password, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetEvaultNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetExecutiveSummaryPdf(pdfType *string, historicalType *string, startDate *string, endDate *string) (resp []byte, err error) {
	params := []interface{}{
		pdfType,
		historicalType,
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetExpiredSecurityCertificates() (resp []datatypes.Security_Certificate, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetFacilityLogs() (resp []datatypes.User_Access_Facility_Log, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetFlexibleCreditEnrollments() (resp []datatypes.FlexibleCredit_Enrollment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetFlexibleCreditProgramInfo(forNextBillCycle *bool) (resp datatypes.Container_Account_Discount_Program, err error) {
	params := []interface{}{
		forNextBillCycle,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetGlobalIpRecords() (resp []datatypes.Network_Subnet_IpAddress_Global, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetGlobalIpv4Records() (resp []datatypes.Network_Subnet_IpAddress_Global, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetGlobalIpv6Records() (resp []datatypes.Network_Subnet_IpAddress_Global, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetGlobalLoadBalancerAccounts() (resp []datatypes.Network_LoadBalancer_Global_Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetHardwareOverBandwidthAllocation() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetHardwareProjectedOverBandwidthAllocation() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetHardwareWithCpanel() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetHardwareWithHelm() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetHardwareWithMcafee() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetHardwareWithMcafeeAntivirusRedhat() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetHardwareWithMcafeeAntivirusWindows() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetHardwareWithMcafeeIntrusionDetectionSystem() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetHardwareWithPlesk() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetHardwareWithQuantastor() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetHardwareWithUrchin() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetHardwareWithWindows() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetHasEvaultBareMetalRestorePluginFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetHasIderaBareMetalRestorePluginFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetHasPendingOrder() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetHasR1softBareMetalRestorePluginFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetHistoricalBackupGraph(startDate *datatypes.Time, endDate *datatypes.Time) (resp datatypes.Container_Account_Graph_Outputs, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetHistoricalBandwidthGraph(startDate *datatypes.Time, endDate *datatypes.Time) (resp datatypes.Container_Account_Graph_Outputs, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetHistoricalTicketGraph(startDate *datatypes.Time, endDate *datatypes.Time) (resp datatypes.Container_Account_Graph_Outputs, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetHistoricalUptimeGraph(startDate *datatypes.Time, endDate *datatypes.Time) (resp datatypes.Container_Account_Graph_Outputs, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetHourlyBareMetalInstances() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetHourlyServiceBillingItems() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetHourlyVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetHubNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetInternalNotes() (resp []datatypes.Account_Note, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetInvoices() (resp []datatypes.Billing_Invoice, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetIpAddresses() (resp []datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetIscsiNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetLargestAllowedSubnetCidr(numberOfHosts *int, locationId *int) (resp int, err error) {
	params := []interface{}{
		numberOfHosts,
		locationId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetLastCanceledBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetLastCancelledServerBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetLastFiveClosedAbuseTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetLastFiveClosedAccountingTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetLastFiveClosedOtherTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetLastFiveClosedSalesTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetLastFiveClosedSupportTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetLastFiveClosedTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetLatestBillDate() (resp datatypes.Time, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetLatestRecurringInvoice() (resp datatypes.Billing_Invoice, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetLatestRecurringPendingInvoice() (resp datatypes.Billing_Invoice, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetLegacyBandwidthAllotments() (resp []datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetLegacyIscsiCapacityGB() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetLoadBalancers() (resp []datatypes.Network_LoadBalancer_VirtualIpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetLockboxCapacityGB() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetLockboxNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetManualPaymentsUnderReview() (resp []datatypes.Billing_Payment_Card_ManualPayment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetMasterUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetMediaDataTransferRequests() (resp []datatypes.Account_Media_Data_Transfer_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetMessageQueueAccounts() (resp []datatypes.Network_Message_Queue, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetMonthlyBareMetalInstances() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetMonthlyVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNasNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNetworkCreationFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNetworkGateways() (resp []datatypes.Network_Gateway, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNetworkHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNetworkMessageDeliveryAccounts() (resp []datatypes.Network_Message_Delivery, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNetworkMonitorDownHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNetworkMonitorDownVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNetworkMonitorRecoveringHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNetworkMonitorRecoveringVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNetworkMonitorUpHardware() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNetworkMonitorUpVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNetworkStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNetworkStorageGroups() (resp []datatypes.Network_Storage_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNetworkTunnelContexts() (resp []datatypes.Network_Tunnel_Module_Context, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNetworkVlanSpan() (resp datatypes.Account_Network_Vlan_Span, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNetworkVlans() (resp []datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNextBillingPublicAllotmentHardwareBandwidthDetails() (resp []datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNextInvoiceExcel(documentCreateDate *datatypes.Time) (resp []byte, err error) {
	params := []interface{}{
		documentCreateDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNextInvoiceIncubatorExemptTotal() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNextInvoicePdf(documentCreateDate *datatypes.Time) (resp []byte, err error) {
	params := []interface{}{
		documentCreateDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNextInvoicePdfDetailed(documentCreateDate *datatypes.Time) (resp []byte, err error) {
	params := []interface{}{
		documentCreateDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNextInvoiceTopLevelBillingItems() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNextInvoiceTotalAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNextInvoiceTotalOneTimeAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNextInvoiceTotalOneTimeTaxAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNextInvoiceTotalRecurringAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNextInvoiceTotalRecurringAmountBeforeAccountDiscount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNextInvoiceTotalRecurringTaxAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNextInvoiceTotalTaxableRecurringAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNextInvoiceZeroFeeItemCounts() (resp []datatypes.Container_Product_Item_Category_ZeroFee_Count, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetNotificationSubscribers() (resp []datatypes.Notification_Subscriber, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetObject() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetOpenAbuseTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetOpenAccountingTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetOpenBillingTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetOpenCancellationRequests() (resp []datatypes.Billing_Item_Cancellation_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetOpenOtherTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetOpenRecurringInvoices() (resp []datatypes.Billing_Invoice, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetOpenSalesTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetOpenStackAccountLinks() (resp []datatypes.Account_Link, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetOpenStackObjectStorage() (resp []datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetOpenSupportTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetOpenTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetOpenTicketsWaitingOnCustomer() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetOrders() (resp []datatypes.Billing_Order, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetOrphanBillingItems() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetOwnedBrands() (resp []datatypes.Brand, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetOwnedHardwareGenericComponentModels() (resp []datatypes.Hardware_Component_Model_Generic, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPaymentProcessors() (resp []datatypes.Billing_Payment_Processor, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPendingCreditCardChangeRequestData() (resp []datatypes.Container_Account_Payment_Method_CreditCard, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPendingEvents() (resp []datatypes.Notification_Occurrence_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPendingInvoice() (resp datatypes.Billing_Invoice, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPendingInvoiceTopLevelItems() (resp []datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPendingInvoiceTotalAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPendingInvoiceTotalOneTimeAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPendingInvoiceTotalOneTimeTaxAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPendingInvoiceTotalRecurringAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPendingInvoiceTotalRecurringTaxAmount() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPermissionGroups() (resp []datatypes.User_Permission_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPermissionRoles() (resp []datatypes.User_Permission_Role, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPortableStorageVolumes() (resp []datatypes.Virtual_Disk_Image, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPostProvisioningHooks() (resp []datatypes.Provisioning_Hook, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPptpVpnUsers() (resp []datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPreviousRecurringRevenue() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPriceRestrictions() (resp []datatypes.Product_Item_Price_Account_Restriction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPriorityOneTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPrivateAllotmentHardwareBandwidthDetails() (resp []datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPrivateBlockDeviceTemplateGroups() (resp []datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPrivateIpAddresses() (resp []datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPrivateNetworkVlans() (resp []datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPrivateSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPublicAllotmentHardwareBandwidthDetails() (resp []datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPublicIpAddresses() (resp []datatypes.Network_Subnet_IpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPublicNetworkVlans() (resp []datatypes.Network_Vlan, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetPublicSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetQuotes() (resp []datatypes.Billing_Order_Quote, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetRecentEvents() (resp []datatypes.Notification_Occurrence_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetReferralPartner() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetReferralPartnerCommissionForecast() (resp []datatypes.Container_Referral_Partner_Commission, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetReferralPartnerCommissionHistory() (resp []datatypes.Container_Referral_Partner_Commission, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetReferralPartnerCommissionPending() (resp []datatypes.Container_Referral_Partner_Commission, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetReferredAccounts() (resp []datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetRegulatedWorkloads() (resp []datatypes.Legal_RegulatedWorkload, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetRemoteManagementCommandRequests() (resp []datatypes.Hardware_Component_RemoteManagement_Command_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetReplicationEvents() (resp []datatypes.Network_Storage_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetResourceGroups() (resp []datatypes.Resource_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetRwhoisData() (resp datatypes.Network_Subnet_Rwhois_Data, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetSalesforceAccountLink() (resp datatypes.Account_Link, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetSamlAuthentication() (resp datatypes.Account_Authentication_Saml, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetScaleGroups() (resp []datatypes.Scale_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetSecondaryDomains() (resp []datatypes.Dns_Secondary, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetSecurityCertificates() (resp []datatypes.Security_Certificate, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetSecurityScanRequests() (resp []datatypes.Network_Security_Scanner_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetServiceBillingItems() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetSharedBlockDeviceTemplateGroups() (resp []datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetShipments() (resp []datatypes.Account_Shipment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetSshKeys() (resp []datatypes.Security_Ssh_Key, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetSslVpnUsers() (resp []datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetStandardPoolVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetSubnetRegistrationDetails() (resp []datatypes.Account_Regional_Registry_Detail, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetSubnetRegistrations() (resp []datatypes.Network_Subnet_Registration, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetSupportRepresentatives() (resp []datatypes.User_Employee, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetSupportSubscriptions() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetSupportTier() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetSuppressInvoicesFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetTags() (resp []datatypes.Tag, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetTechIncubatorProgramInfo(forNextBillCycle *bool) (resp datatypes.Container_Account_Discount_Program, err error) {
	params := []interface{}{
		forNextBillCycle,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetThirdPartyPoliciesAcceptanceStatus() (resp []datatypes.Container_Policy_Acceptance, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetTicketsClosedInTheLastThreeDays() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetTicketsClosedToday() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetTranscodeAccounts() (resp []datatypes.Network_Media_Transcode_Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetUpgradeRequests() (resp []datatypes.Product_Upgrade_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetUsers() (resp []datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetValidSecurityCertificateEntries() (resp []datatypes.Security_Certificate_Entry, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetValidSecurityCertificates() (resp []datatypes.Security_Certificate, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetVdrUpdatesInProgressFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetVirtualDedicatedRacks() (resp []datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetVirtualDiskImages() (resp []datatypes.Virtual_Disk_Image, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetVirtualGuests() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetVirtualGuestsOverBandwidthAllocation() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetVirtualGuestsProjectedOverBandwidthAllocation() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetVirtualGuestsWithCpanel() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetVirtualGuestsWithMcafee() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetVirtualGuestsWithMcafeeAntivirusRedhat() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetVirtualGuestsWithMcafeeAntivirusWindows() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetVirtualGuestsWithMcafeeIntrusionDetectionSystem() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetVirtualGuestsWithPlesk() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetVirtualGuestsWithQuantastor() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetVirtualGuestsWithUrchin() (resp []datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetVirtualPrivateRack() (resp datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetVirtualStorageArchiveRepositories() (resp []datatypes.Virtual_Storage_Repository, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetVirtualStoragePublicRepositories() (resp []datatypes.Virtual_Storage_Repository, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetVmWareActiveAccountLicenseKeys() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) GetWindowsUpdateStatus() (resp []datatypes.Container_Utility_Microsoft_Windows_UpdateServices_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) HasAttribute(attributeType *string) (resp bool, err error) {
	params := []interface{}{
		attributeType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) HourlyInstanceLimit() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) HourlyServerLimit() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) LinkExternalAccount(externalAccountId *string, authorizationToken *string, externalServiceProviderKey *string) (err error) {
	var resp datatypes.Void
	params := []interface{}{
		externalAccountId,
		authorizationToken,
		externalServiceProviderKey,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) RemoveAlternateCreditCard() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) RequestCreditCardChange(request *datatypes.Billing_Payment_Card_ChangeRequest, vatId *string, paymentRoleName *string, onlyChangeNicknameFlag *bool) (resp datatypes.Billing_Payment_Card_ChangeRequest, err error) {
	params := []interface{}{
		request,
		vatId,
		paymentRoleName,
		onlyChangeNicknameFlag,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) RequestManualPayment(request *datatypes.Billing_Payment_Card_ManualPayment) (resp datatypes.Billing_Payment_Card_ManualPayment, err error) {
	params := []interface{}{
		request,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) RequestManualPaymentUsingCreditCardOnFile(amount *string, payWithAlternateCardFlag *bool, note *string) (resp datatypes.Billing_Payment_Card_ManualPayment, err error) {
	params := []interface{}{
		amount,
		payWithAlternateCardFlag,
		note,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) SetAbuseEmails(emails []string) (resp bool, err error) {
	params := []interface{}{
		emails,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) SetVlanSpan(enabled *bool) (resp bool, err error) {
	params := []interface{}{
		enabled,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) SwapCreditCards() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account) UpdateVpnUsersForResource(objectId *int, objectType *string) (resp bool, err error) {
	params := []interface{}{
		objectId,
		objectType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) Validate(account *datatypes.Account) (resp []string, err error) {
	params := []interface{}{
		account,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account) ValidateManualPaymentAmount(amount *string) (resp bool, err error) {
	params := []interface{}{
		amount,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Account_Address struct {
	Session *Session
	Options
}

func (r *Session) GetAccountAddressService() Account_Address {
	return Account_Address{Session: r}
}

func (r *Account_Address) CreateObject(templateObject *datatypes.Account_Address) (resp datatypes.Account_Address, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Address) EditObject(templateObject *datatypes.Account_Address) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Address) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Address) GetAllDataCenters() (resp []datatypes.Account_Address, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Address) GetCreateUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Address) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Address) GetModifyEmployee() (resp datatypes.User_Employee, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Address) GetModifyUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Address) GetNetworkAddress(name *string) (resp []datatypes.Account_Address, err error) {
	params := []interface{}{
		name,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Address) GetObject() (resp datatypes.Account_Address, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Address) GetType() (resp datatypes.Account_Address_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Account_Address_Type struct {
	Session *Session
	Options
}

func (r *Session) GetAccountAddressTypeService() Account_Address_Type {
	return Account_Address_Type{Session: r}
}

func (r *Account_Address_Type) GetObject() (resp datatypes.Account_Address_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Account_Affiliation struct {
	Session *Session
	Options
}

func (r *Session) GetAccountAffiliationService() Account_Affiliation {
	return Account_Affiliation{Session: r}
}

func (r *Account_Affiliation) CreateObject(templateObject *datatypes.Account_Affiliation) (resp datatypes.Account_Affiliation, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Affiliation) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Affiliation) EditObject(templateObject *datatypes.Account_Affiliation) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Affiliation) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Affiliation) GetAccountAffiliationsByAffiliateId(affiliateId *string) (resp []datatypes.Account_Affiliation, err error) {
	params := []interface{}{
		affiliateId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Affiliation) GetObject() (resp datatypes.Account_Affiliation, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Account_Agreement struct {
	Session *Session
	Options
}

func (r *Session) GetAccountAgreementService() Account_Agreement {
	return Account_Agreement{Session: r}
}

func (r *Account_Agreement) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Agreement) GetAgreementType() (resp datatypes.Account_Agreement_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Agreement) GetAttachedBillingAgreementFiles() (resp []datatypes.Account_MasterServiceAgreement, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Agreement) GetBillingItems() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Agreement) GetObject() (resp datatypes.Account_Agreement, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Agreement) GetStatus() (resp datatypes.Account_Agreement_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Agreement) GetTopLevelBillingItems() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Account_Authentication_Attribute struct {
	Session *Session
	Options
}

func (r *Session) GetAccountAuthenticationAttributeService() Account_Authentication_Attribute {
	return Account_Authentication_Attribute{Session: r}
}

func (r *Account_Authentication_Attribute) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Authentication_Attribute) GetAuthenticationRecord() (resp datatypes.Account_Authentication_Saml, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Authentication_Attribute) GetObject() (resp datatypes.Account_Authentication_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Authentication_Attribute) GetType() (resp datatypes.Account_Authentication_Attribute_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Account_Authentication_Attribute_Type struct {
	Session *Session
	Options
}

func (r *Session) GetAccountAuthenticationAttributeTypeService() Account_Authentication_Attribute_Type {
	return Account_Authentication_Attribute_Type{Session: r}
}

func (r *Account_Authentication_Attribute_Type) GetAllObjects() (resp []datatypes.Account_Attribute_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Authentication_Attribute_Type) GetObject() (resp datatypes.Account_Authentication_Attribute_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Account_Authentication_Saml struct {
	Session *Session
	Options
}

func (r *Session) GetAccountAuthenticationSamlService() Account_Authentication_Saml {
	return Account_Authentication_Saml{Session: r}
}

func (r *Account_Authentication_Saml) CreateObject(templateObject *datatypes.Account_Authentication_Saml) (resp datatypes.Account_Authentication_Saml, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Authentication_Saml) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Authentication_Saml) EditObject(templateObject *datatypes.Account_Authentication_Saml) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Authentication_Saml) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Authentication_Saml) GetAttributes() (resp []datatypes.Account_Authentication_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Authentication_Saml) GetMetadata() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Authentication_Saml) GetObject() (resp datatypes.Account_Authentication_Saml, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Account_Contact struct {
	Session *Session
	Options
}

func (r *Session) GetAccountContactService() Account_Contact {
	return Account_Contact{Session: r}
}

func (r *Account_Contact) CreateObject(templateObject *datatypes.Account_Contact) (resp datatypes.Account_Contact, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Contact) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Contact) EditObject(templateObject *datatypes.Account_Contact) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Contact) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Contact) GetAllContactTypes() (resp []datatypes.Account_Contact_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Contact) GetObject() (resp datatypes.Account_Contact, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Contact) GetType() (resp datatypes.Account_Contact_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Account_Historical_Report struct {
	Session *Session
	Options
}

func (r *Session) GetAccountHistoricalReportService() Account_Historical_Report {
	return Account_Historical_Report{Session: r}
}

func (r *Account_Historical_Report) GetAccountHostUptimeGraphData(startDate *string, endDate *string) (resp datatypes.Container_Graph, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Historical_Report) GetAccountHostUptimeSummary(startDateTime *string, endDateTime *string) (resp datatypes.Container_Account_Historical_Summary, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Historical_Report) GetAccountUrlUptimeGraphData(startDate *string, endDate *string) (resp datatypes.Container_Graph, err error) {
	params := []interface{}{
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Historical_Report) GetAccountUrlUptimeSummary(startDateTime *string, endDateTime *string) (resp datatypes.Container_Account_Historical_Summary, err error) {
	params := []interface{}{
		startDateTime,
		endDateTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Historical_Report) GetHostUptimeDetail(configurationValueId *int, startDateTime *string, endDateTime *string) (resp datatypes.Container_Account_Historical_Summary_Detail, err error) {
	params := []interface{}{
		configurationValueId,
		startDateTime,
		endDateTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Historical_Report) GetHostUptimeGraphData(configurationValueId *int, startDate *string, endDate *string) (resp datatypes.Container_Graph, err error) {
	params := []interface{}{
		configurationValueId,
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Historical_Report) GetUrlUptimeDetail(configurationValueId *int, startDateTime *string, endDateTime *string) (resp datatypes.Container_Account_Historical_Summary_Detail, err error) {
	params := []interface{}{
		configurationValueId,
		startDateTime,
		endDateTime,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Historical_Report) GetUrlUptimeGraphData(configurationValueId *int, startDate *string, endDate *string) (resp datatypes.Container_Graph, err error) {
	params := []interface{}{
		configurationValueId,
		startDate,
		endDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Account_Link_Bluemix struct {
	Session *Session
	Options
}

func (r *Session) GetAccountLinkBluemixService() Account_Link_Bluemix {
	return Account_Link_Bluemix{Session: r}
}

func (r *Account_Link_Bluemix) GetObject() (resp datatypes.Account_Link_Bluemix, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Link_Bluemix) GetSupportTierType() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Account_Link_OpenStack struct {
	Session *Session
	Options
}

func (r *Session) GetAccountLinkOpenStackService() Account_Link_OpenStack {
	return Account_Link_OpenStack{Session: r}
}

func (r *Account_Link_OpenStack) CreateOSDomain(request *datatypes.Account_Link_OpenStack_LinkRequest) (resp datatypes.Account_Link_OpenStack_DomainCreationDetails, err error) {
	params := []interface{}{
		request,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Link_OpenStack) CreateOSProject(request *datatypes.Account_Link_OpenStack_LinkRequest) (resp datatypes.Account_Link_OpenStack_ProjectCreationDetails, err error) {
	params := []interface{}{
		request,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Link_OpenStack) DeleteOSDomain(domainId *string) (resp bool, err error) {
	params := []interface{}{
		domainId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Link_OpenStack) DeleteOSProject(projectId *string) (resp bool, err error) {
	params := []interface{}{
		projectId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Link_OpenStack) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Link_OpenStack) GetOSProject(projectId *string) (resp datatypes.Account_Link_OpenStack_ProjectDetails, err error) {
	params := []interface{}{
		projectId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Link_OpenStack) GetObject() (resp datatypes.Account_Link_OpenStack, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Link_OpenStack) ListOSProjects() (resp []datatypes.Account_Link_OpenStack_ProjectDetails, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Account_Lockdown_Request struct {
	Session *Session
	Options
}

func (r *Session) GetAccountLockdownRequestService() Account_Lockdown_Request {
	return Account_Lockdown_Request{Session: r}
}

func (r *Account_Lockdown_Request) CancelRequest() (err error) {
	var resp datatypes.Void
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Lockdown_Request) DisableLockedAccount(disableDate *string) (resp int, err error) {
	params := []interface{}{
		disableDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Lockdown_Request) DisconnectCompute(accountId *int, disconnectDate *string) (resp int, err error) {
	params := []interface{}{
		accountId,
		disconnectDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Lockdown_Request) GetAccountHistory(accountId *int) (resp []datatypes.Account_Lockdown_Request, err error) {
	params := []interface{}{
		accountId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Lockdown_Request) GetObject() (resp datatypes.Account_Lockdown_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Lockdown_Request) ReconnectCompute(reconnectDate *string) (resp int, err error) {
	params := []interface{}{
		reconnectDate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Account_MasterServiceAgreement struct {
	Session *Session
	Options
}

func (r *Session) GetAccountMasterServiceAgreementService() Account_MasterServiceAgreement {
	return Account_MasterServiceAgreement{Session: r}
}

func (r *Account_MasterServiceAgreement) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_MasterServiceAgreement) GetFile() (resp datatypes.Container_Utility_File_Entity, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_MasterServiceAgreement) GetObject() (resp datatypes.Account_MasterServiceAgreement, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Account_Media struct {
	Session *Session
	Options
}

func (r *Session) GetAccountMediaService() Account_Media {
	return Account_Media{Session: r}
}

func (r *Account_Media) EditObject(templateObject *datatypes.Account_Media) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Media) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Media) GetAllMediaTypes() (resp []datatypes.Account_Media_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Media) GetCreateUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Media) GetDatacenter() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Media) GetModifyEmployee() (resp datatypes.User_Employee, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Media) GetModifyUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Media) GetObject() (resp datatypes.Account_Media, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Media) GetRequest() (resp datatypes.Account_Media_Data_Transfer_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Media) GetType() (resp datatypes.Account_Media_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Media) GetVolume() (resp datatypes.Network_Storage, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Media) RemoveMediaFromList(mediaTemplate *datatypes.Account_Media) (resp int, err error) {
	params := []interface{}{
		mediaTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Account_Media_Data_Transfer_Request struct {
	Session *Session
	Options
}

func (r *Session) GetAccountMediaDataTransferRequestService() Account_Media_Data_Transfer_Request {
	return Account_Media_Data_Transfer_Request{Session: r}
}

func (r *Account_Media_Data_Transfer_Request) EditObject(templateObject *datatypes.Account_Media_Data_Transfer_Request) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Media_Data_Transfer_Request) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Media_Data_Transfer_Request) GetActiveTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Media_Data_Transfer_Request) GetAllRequestStatuses() (resp []datatypes.Account_Media_Data_Transfer_Request_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Media_Data_Transfer_Request) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Media_Data_Transfer_Request) GetCreateUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Media_Data_Transfer_Request) GetMedia() (resp datatypes.Account_Media, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Media_Data_Transfer_Request) GetModifyEmployee() (resp datatypes.User_Employee, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Media_Data_Transfer_Request) GetModifyUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Media_Data_Transfer_Request) GetObject() (resp datatypes.Account_Media_Data_Transfer_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Media_Data_Transfer_Request) GetShipments() (resp []datatypes.Account_Shipment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Media_Data_Transfer_Request) GetStatus() (resp datatypes.Account_Media_Data_Transfer_Request_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Media_Data_Transfer_Request) GetTickets() (resp []datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Account_Note struct {
	Session *Session
	Options
}

func (r *Session) GetAccountNoteService() Account_Note {
	return Account_Note{Session: r}
}

func (r *Account_Note) CreateObject(templateObject *datatypes.Account_Note) (resp datatypes.Account_Note, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Note) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Note) EditObject(templateObject *datatypes.Account_Note) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Note) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Note) GetCustomer() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Note) GetNoteHistory() (resp []datatypes.Account_Note_History, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Note) GetNoteType() (resp datatypes.Account_Note_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Note) GetObject() (resp datatypes.Account_Note, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Account_Note_Type struct {
	Session *Session
	Options
}

func (r *Session) GetAccountNoteTypeService() Account_Note_Type {
	return Account_Note_Type{Session: r}
}

func (r *Account_Note_Type) CreateObject(templateObject *datatypes.Account_Note_Type) (resp datatypes.Account_Note_Type, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Note_Type) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Note_Type) EditObject(templateObject *datatypes.Account_Note_Type) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Note_Type) GetAllObjects() (resp []datatypes.Account_Note_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Note_Type) GetObject() (resp datatypes.Account_Note_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Account_Partner_Referral_Prospect struct {
	Session *Session
	Options
}

func (r *Session) GetAccountPartnerReferralProspectService() Account_Partner_Referral_Prospect {
	return Account_Partner_Referral_Prospect{Session: r}
}

func (r *Account_Partner_Referral_Prospect) CreateProspect(templateObject *datatypes.Container_Referral_Partner_Prospect, commit *bool) (resp datatypes.Account_Partner_Referral_Prospect, err error) {
	params := []interface{}{
		templateObject,
		commit,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Partner_Referral_Prospect) GetObject() (resp datatypes.Account_Partner_Referral_Prospect, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Partner_Referral_Prospect) GetSurveyQuestions() (resp []datatypes.Survey_Question, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Account_Password struct {
	Session *Session
	Options
}

func (r *Session) GetAccountPasswordService() Account_Password {
	return Account_Password{Session: r}
}

func (r *Account_Password) EditObject(templateObject *datatypes.Account_Password) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Password) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Password) GetObject() (resp datatypes.Account_Password, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Password) GetType() (resp datatypes.Account_Password_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Account_Regional_Registry_Detail struct {
	Session *Session
	Options
}

func (r *Session) GetAccountRegionalRegistryDetailService() Account_Regional_Registry_Detail {
	return Account_Regional_Registry_Detail{Session: r}
}

func (r *Account_Regional_Registry_Detail) CreateObject(templateObject *datatypes.Account_Regional_Registry_Detail) (resp datatypes.Account_Regional_Registry_Detail, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Regional_Registry_Detail) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Regional_Registry_Detail) EditObject(templateObject *datatypes.Account_Regional_Registry_Detail) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Regional_Registry_Detail) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Regional_Registry_Detail) GetDetailType() (resp datatypes.Account_Regional_Registry_Detail_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Regional_Registry_Detail) GetDetails() (resp []datatypes.Network_Subnet_Registration_Details, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Regional_Registry_Detail) GetObject() (resp datatypes.Account_Regional_Registry_Detail, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Regional_Registry_Detail) GetProperties() (resp []datatypes.Account_Regional_Registry_Detail_Property, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Regional_Registry_Detail) GetRegionalInternetRegistryHandle() (resp datatypes.Account_Rwhois_Handle, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Regional_Registry_Detail) UpdateReferencedRegistrations() (resp datatypes.Container_Network_Subnet_Registration_TransactionDetails, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Account_Regional_Registry_Detail_Property struct {
	Session *Session
	Options
}

func (r *Session) GetAccountRegionalRegistryDetailPropertyService() Account_Regional_Registry_Detail_Property {
	return Account_Regional_Registry_Detail_Property{Session: r}
}

func (r *Account_Regional_Registry_Detail_Property) CreateObject(templateObject *datatypes.Account_Regional_Registry_Detail_Property) (resp datatypes.Account_Regional_Registry_Detail_Property, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Regional_Registry_Detail_Property) CreateObjects(templateObjects []datatypes.Account_Regional_Registry_Detail_Property) (resp []datatypes.Account_Regional_Registry_Detail_Property, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Regional_Registry_Detail_Property) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Regional_Registry_Detail_Property) EditObject(templateObject *datatypes.Account_Regional_Registry_Detail_Property) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Regional_Registry_Detail_Property) EditObjects(templateObjects []datatypes.Account_Regional_Registry_Detail_Property) (resp bool, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Regional_Registry_Detail_Property) GetDetail() (resp datatypes.Account_Regional_Registry_Detail, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Regional_Registry_Detail_Property) GetObject() (resp datatypes.Account_Regional_Registry_Detail_Property, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Regional_Registry_Detail_Property) GetPropertyType() (resp datatypes.Account_Regional_Registry_Detail_Property_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Account_Regional_Registry_Detail_Property_Type struct {
	Session *Session
	Options
}

func (r *Session) GetAccountRegionalRegistryDetailPropertyTypeService() Account_Regional_Registry_Detail_Property_Type {
	return Account_Regional_Registry_Detail_Property_Type{Session: r}
}

func (r *Account_Regional_Registry_Detail_Property_Type) GetAllObjects() (resp []datatypes.Account_Regional_Registry_Detail_Property_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Regional_Registry_Detail_Property_Type) GetObject() (resp datatypes.Account_Regional_Registry_Detail_Property_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Account_Regional_Registry_Detail_Type struct {
	Session *Session
	Options
}

func (r *Session) GetAccountRegionalRegistryDetailTypeService() Account_Regional_Registry_Detail_Type {
	return Account_Regional_Registry_Detail_Type{Session: r}
}

func (r *Account_Regional_Registry_Detail_Type) GetAllObjects() (resp []datatypes.Account_Regional_Registry_Detail_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Regional_Registry_Detail_Type) GetObject() (resp datatypes.Account_Regional_Registry_Detail_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Account_Reports_Request struct {
	Session *Session
	Options
}

func (r *Session) GetAccountReportsRequestService() Account_Reports_Request {
	return Account_Reports_Request{Session: r}
}

func (r *Account_Reports_Request) CreateRequest(contact *datatypes.Account_Contact, reason *string, reportType *string) (resp datatypes.Account_Reports_Request, err error) {
	params := []interface{}{
		contact,
		reason,
		reportType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Reports_Request) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Reports_Request) GetAccountContact() (resp datatypes.Account_Contact, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Reports_Request) GetAllObjects() (resp datatypes.Account_Reports_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Reports_Request) GetObject() (resp datatypes.Account_Reports_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Reports_Request) GetReportType() (resp datatypes.Compliance_Report_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Reports_Request) GetRequestByRequestKey(requestKey *string) (resp datatypes.Account_Reports_Request, err error) {
	params := []interface{}{
		requestKey,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Reports_Request) GetTicket() (resp datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Reports_Request) GetUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Reports_Request) SendReportEmail(request *datatypes.Account_Reports_Request) (resp bool, err error) {
	params := []interface{}{
		request,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Reports_Request) UpdateTicketOnDecline(request *datatypes.Account_Reports_Request) (resp bool, err error) {
	params := []interface{}{
		request,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Account_Shipment struct {
	Session *Session
	Options
}

func (r *Session) GetAccountShipmentService() Account_Shipment {
	return Account_Shipment{Session: r}
}

func (r *Account_Shipment) EditObject(templateObject *datatypes.Account_Shipment) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment) GetAllCouriers() (resp []datatypes.Auxiliary_Shipping_Courier, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment) GetAllCouriersByType(courierTypeKeyName *string) (resp []datatypes.Auxiliary_Shipping_Courier, err error) {
	params := []interface{}{
		courierTypeKeyName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment) GetAllShipmentStatuses() (resp []datatypes.Account_Shipment_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment) GetAllShipmentTypes() (resp []datatypes.Account_Shipment_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment) GetCourier() (resp datatypes.Auxiliary_Shipping_Courier, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment) GetCreateEmployee() (resp datatypes.User_Employee, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment) GetCreateUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment) GetDestinationAddress() (resp datatypes.Account_Address, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment) GetModifyEmployee() (resp datatypes.User_Employee, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment) GetModifyUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment) GetObject() (resp datatypes.Account_Shipment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment) GetOriginationAddress() (resp datatypes.Account_Address, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment) GetShipmentItems() (resp []datatypes.Account_Shipment_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment) GetStatus() (resp datatypes.Account_Shipment_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment) GetTrackingData() (resp []datatypes.Account_Shipment_Tracking_Data, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment) GetType() (resp datatypes.Account_Shipment_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Account_Shipment_Item struct {
	Session *Session
	Options
}

func (r *Session) GetAccountShipmentItemService() Account_Shipment_Item {
	return Account_Shipment_Item{Session: r}
}

func (r *Account_Shipment_Item) EditObject(templateObject *datatypes.Account_Shipment_Item) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment_Item) GetObject() (resp datatypes.Account_Shipment_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment_Item) GetShipment() (resp datatypes.Account_Shipment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment_Item) GetShipmentItemType() (resp datatypes.Account_Shipment_Item_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Account_Shipment_Item_Type struct {
	Session *Session
	Options
}

func (r *Session) GetAccountShipmentItemTypeService() Account_Shipment_Item_Type {
	return Account_Shipment_Item_Type{Session: r}
}

func (r *Account_Shipment_Item_Type) GetObject() (resp datatypes.Account_Shipment_Item_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Account_Shipment_Resource_Type struct {
	Session *Session
	Options
}

func (r *Session) GetAccountShipmentResourceTypeService() Account_Shipment_Resource_Type {
	return Account_Shipment_Resource_Type{Session: r}
}

func (r *Account_Shipment_Resource_Type) GetObject() (resp datatypes.Account_Shipment_Resource_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Account_Shipment_Status struct {
	Session *Session
	Options
}

func (r *Session) GetAccountShipmentStatusService() Account_Shipment_Status {
	return Account_Shipment_Status{Session: r}
}

func (r *Account_Shipment_Status) GetObject() (resp datatypes.Account_Shipment_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Account_Shipment_Tracking_Data struct {
	Session *Session
	Options
}

func (r *Session) GetAccountShipmentTrackingDataService() Account_Shipment_Tracking_Data {
	return Account_Shipment_Tracking_Data{Session: r}
}

func (r *Account_Shipment_Tracking_Data) CreateObject(templateObject *datatypes.Account_Shipment_Tracking_Data) (resp datatypes.Account_Shipment_Tracking_Data, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment_Tracking_Data) CreateObjects(templateObjects []datatypes.Account_Shipment_Tracking_Data) (resp []datatypes.Account_Shipment_Tracking_Data, err error) {
	params := []interface{}{
		templateObjects,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment_Tracking_Data) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment_Tracking_Data) EditObject(templateObject *datatypes.Account_Shipment_Tracking_Data) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment_Tracking_Data) GetCreateEmployee() (resp datatypes.User_Employee, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment_Tracking_Data) GetCreateUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment_Tracking_Data) GetModifyEmployee() (resp datatypes.User_Employee, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment_Tracking_Data) GetModifyUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment_Tracking_Data) GetObject() (resp datatypes.Account_Shipment_Tracking_Data, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Account_Shipment_Tracking_Data) GetShipment() (resp datatypes.Account_Shipment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Account_Shipment_Type struct {
	Session *Session
	Options
}

func (r *Session) GetAccountShipmentTypeService() Account_Shipment_Type {
	return Account_Shipment_Type{Session: r}
}

func (r *Account_Shipment_Type) GetObject() (resp datatypes.Account_Shipment_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
