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

// The SoftLayer_Account data type contains general information relating to a single SoftLayer customer account. Personal information in this type such as names, addresses, and phone numbers are assigned to the account only and not to users belonging to the account. The SoftLayer_Account data type contains a number of relational properties that are used by the SoftLayer customer portal to quickly present a variety of account related services to it's users.
//
// SoftLayer customers are unable to change their company account information in the portal or the API. If you need to change this information please open a sales ticket in our customer portal and our account management staff will assist you.
type Account struct {
	Entity

	// An email address that is responsible for abuse and legal inquiries on behalf of an account. For instance, new legal and abuse tickets are sent to this address.
	AbuseEmail *string `json:"abuseEmail,omitempty" xmlrpc:"abuseEmail"`

	// A count of email addresses that are responsible for abuse and legal inquiries on behalf of an account. For instance, new legal and abuse tickets are sent to these addresses.
	AbuseEmailCount *uint `json:"abuseEmailCount,omitempty" xmlrpc:"abuseEmailCount"`

	// Email addresses that are responsible for abuse and legal inquiries on behalf of an account. For instance, new legal and abuse tickets are sent to these addresses.
	AbuseEmails []Account_AbuseEmail `json:"abuseEmails,omitempty" xmlrpc:"abuseEmails"`

	// A count of the account contacts on an account.
	AccountContactCount *uint `json:"accountContactCount,omitempty" xmlrpc:"accountContactCount"`

	// The account contacts on an account.
	AccountContacts []Account_Contact `json:"accountContacts,omitempty" xmlrpc:"accountContacts"`

	// A count of the account software licenses owned by an account
	AccountLicenseCount *uint `json:"accountLicenseCount,omitempty" xmlrpc:"accountLicenseCount"`

	// The account software licenses owned by an account
	AccountLicenses []Software_AccountLicense `json:"accountLicenses,omitempty" xmlrpc:"accountLicenses"`

	// A count of
	AccountLinkCount *uint `json:"accountLinkCount,omitempty" xmlrpc:"accountLinkCount"`

	// no documentation yet
	AccountLinks []Account_Link `json:"accountLinks,omitempty" xmlrpc:"accountLinks"`

	// A flag indicating that the account has a managed resource.
	AccountManagedResourcesFlag *bool `json:"accountManagedResourcesFlag,omitempty" xmlrpc:"accountManagedResourcesFlag"`

	// An account's status presented in a more detailed data type.
	AccountStatus *Account_Status `json:"accountStatus,omitempty" xmlrpc:"accountStatus"`

	// A number reflecting the state of an account.
	AccountStatusId *int `json:"accountStatusId,omitempty" xmlrpc:"accountStatusId"`

	// The billing item associated with an account's monthly discount.
	ActiveAccountDiscountBillingItem *Billing_Item `json:"activeAccountDiscountBillingItem,omitempty" xmlrpc:"activeAccountDiscountBillingItem"`

	// A count of the active account software licenses owned by an account
	ActiveAccountLicenseCount *uint `json:"activeAccountLicenseCount,omitempty" xmlrpc:"activeAccountLicenseCount"`

	// The active account software licenses owned by an account
	ActiveAccountLicenses []Software_AccountLicense `json:"activeAccountLicenses,omitempty" xmlrpc:"activeAccountLicenses"`

	// A count of the active address(es) that belong to an account.
	ActiveAddressCount *uint `json:"activeAddressCount,omitempty" xmlrpc:"activeAddressCount"`

	// The active address(es) that belong to an account.
	ActiveAddresses []Account_Address `json:"activeAddresses,omitempty" xmlrpc:"activeAddresses"`

	// A count of all billing agreements for an account
	ActiveBillingAgreementCount *uint `json:"activeBillingAgreementCount,omitempty" xmlrpc:"activeBillingAgreementCount"`

	// All billing agreements for an account
	ActiveBillingAgreements []Account_Agreement `json:"activeBillingAgreements,omitempty" xmlrpc:"activeBillingAgreements"`

	// no documentation yet
	ActiveCatalystEnrollment *Catalyst_Enrollment `json:"activeCatalystEnrollment,omitempty" xmlrpc:"activeCatalystEnrollment"`

	// A count of the account's active top level colocation containers.
	ActiveColocationContainerCount *uint `json:"activeColocationContainerCount,omitempty" xmlrpc:"activeColocationContainerCount"`

	// The account's active top level colocation containers.
	ActiveColocationContainers []Billing_Item `json:"activeColocationContainers,omitempty" xmlrpc:"activeColocationContainers"`

	// Account's currently active Flexible Credit enrollment.
	ActiveFlexibleCreditEnrollment *FlexibleCredit_Enrollment `json:"activeFlexibleCreditEnrollment,omitempty" xmlrpc:"activeFlexibleCreditEnrollment"`

	// A count of
	ActiveNotificationSubscriberCount *uint `json:"activeNotificationSubscriberCount,omitempty" xmlrpc:"activeNotificationSubscriberCount"`

	// no documentation yet
	ActiveNotificationSubscribers []Notification_Subscriber `json:"activeNotificationSubscribers,omitempty" xmlrpc:"activeNotificationSubscribers"`

	// A count of an account's non-expired quotes.
	ActiveQuoteCount *uint `json:"activeQuoteCount,omitempty" xmlrpc:"activeQuoteCount"`

	// An account's non-expired quotes.
	ActiveQuotes []Billing_Order_Quote `json:"activeQuotes,omitempty" xmlrpc:"activeQuotes"`

	// A count of the virtual software licenses controlled by an account
	ActiveVirtualLicenseCount *uint `json:"activeVirtualLicenseCount,omitempty" xmlrpc:"activeVirtualLicenseCount"`

	// The virtual software licenses controlled by an account
	ActiveVirtualLicenses []Software_VirtualLicense `json:"activeVirtualLicenses,omitempty" xmlrpc:"activeVirtualLicenses"`

	// A count of an account's associated load balancers.
	AdcLoadBalancerCount *uint `json:"adcLoadBalancerCount,omitempty" xmlrpc:"adcLoadBalancerCount"`

	// An account's associated load balancers.
	AdcLoadBalancers []Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"adcLoadBalancers,omitempty" xmlrpc:"adcLoadBalancers"`

	// The first line of the mailing address belonging to an account.
	Address1 *string `json:"address1,omitempty" xmlrpc:"address1"`

	// The second line of the mailing address belonging to an account.
	Address2 *string `json:"address2,omitempty" xmlrpc:"address2"`

	// A count of all the address(es) that belong to an account.
	AddressCount *uint `json:"addressCount,omitempty" xmlrpc:"addressCount"`

	// All the address(es) that belong to an account.
	Addresses []Account_Address `json:"addresses,omitempty" xmlrpc:"addresses"`

	// An affiliate identifier associated with the customer account.
	AffiliateId *string `json:"affiliateId,omitempty" xmlrpc:"affiliateId"`

	// The billing items that will be on an account's next invoice.
	AllBillingItems []Billing_Item `json:"allBillingItems,omitempty" xmlrpc:"allBillingItems"`

	// A count of the billing items that will be on an account's next invoice.
	AllCommissionBillingItemCount *uint `json:"allCommissionBillingItemCount,omitempty" xmlrpc:"allCommissionBillingItemCount"`

	// The billing items that will be on an account's next invoice.
	AllCommissionBillingItems []Billing_Item `json:"allCommissionBillingItems,omitempty" xmlrpc:"allCommissionBillingItems"`

	// A count of the billing items that will be on an account's next invoice.
	AllRecurringTopLevelBillingItemCount *uint `json:"allRecurringTopLevelBillingItemCount,omitempty" xmlrpc:"allRecurringTopLevelBillingItemCount"`

	// The billing items that will be on an account's next invoice.
	AllRecurringTopLevelBillingItems []Billing_Item `json:"allRecurringTopLevelBillingItems,omitempty" xmlrpc:"allRecurringTopLevelBillingItems"`

	// The billing items that will be on an account's next invoice. Does not consider associated items.
	AllRecurringTopLevelBillingItemsUnfiltered []Billing_Item `json:"allRecurringTopLevelBillingItemsUnfiltered,omitempty" xmlrpc:"allRecurringTopLevelBillingItemsUnfiltered"`

	// A count of the billing items that will be on an account's next invoice. Does not consider associated items.
	AllRecurringTopLevelBillingItemsUnfilteredCount *uint `json:"allRecurringTopLevelBillingItemsUnfilteredCount,omitempty" xmlrpc:"allRecurringTopLevelBillingItemsUnfilteredCount"`

	// A count of the billing items that will be on an account's next invoice.
	AllSubnetBillingItemCount *uint `json:"allSubnetBillingItemCount,omitempty" xmlrpc:"allSubnetBillingItemCount"`

	// The billing items that will be on an account's next invoice.
	AllSubnetBillingItems []Billing_Item `json:"allSubnetBillingItems,omitempty" xmlrpc:"allSubnetBillingItems"`

	// A count of all billing items of an account.
	AllTopLevelBillingItemCount *uint `json:"allTopLevelBillingItemCount,omitempty" xmlrpc:"allTopLevelBillingItemCount"`

	// All billing items of an account.
	AllTopLevelBillingItems []Billing_Item `json:"allTopLevelBillingItems,omitempty" xmlrpc:"allTopLevelBillingItems"`

	// The billing items that will be on an account's next invoice. Does not consider associated items.
	AllTopLevelBillingItemsUnfiltered []Billing_Item `json:"allTopLevelBillingItemsUnfiltered,omitempty" xmlrpc:"allTopLevelBillingItemsUnfiltered"`

	// A count of the billing items that will be on an account's next invoice. Does not consider associated items.
	AllTopLevelBillingItemsUnfilteredCount *uint `json:"allTopLevelBillingItemsUnfilteredCount,omitempty" xmlrpc:"allTopLevelBillingItemsUnfilteredCount"`

	// The number of PPTP VPN users allowed on an account.
	AllowedPptpVpnQuantity *int `json:"allowedPptpVpnQuantity,omitempty" xmlrpc:"allowedPptpVpnQuantity"`

	// Flag indicating if this account can be linked with Bluemix.
	AllowsBluemixAccountLinkingFlag *bool `json:"allowsBluemixAccountLinkingFlag,omitempty" xmlrpc:"allowsBluemixAccountLinkingFlag"`

	// A secondary phone number assigned to an account.
	AlternatePhone *string `json:"alternatePhone,omitempty" xmlrpc:"alternatePhone"`

	// A count of an account's associated application delivery controller records.
	ApplicationDeliveryControllerCount *uint `json:"applicationDeliveryControllerCount,omitempty" xmlrpc:"applicationDeliveryControllerCount"`

	// An account's associated application delivery controller records.
	ApplicationDeliveryControllers []Network_Application_Delivery_Controller `json:"applicationDeliveryControllers,omitempty" xmlrpc:"applicationDeliveryControllers"`

	// A count of the account attribute values for a SoftLayer customer account.
	AttributeCount *uint `json:"attributeCount,omitempty" xmlrpc:"attributeCount"`

	// The account attribute values for a SoftLayer customer account.
	Attributes []Account_Attribute `json:"attributes,omitempty" xmlrpc:"attributes"`

	// A count of the public network VLANs assigned to an account.
	AvailablePublicNetworkVlanCount *uint `json:"availablePublicNetworkVlanCount,omitempty" xmlrpc:"availablePublicNetworkVlanCount"`

	// The public network VLANs assigned to an account.
	AvailablePublicNetworkVlans []Network_Vlan `json:"availablePublicNetworkVlans,omitempty" xmlrpc:"availablePublicNetworkVlans"`

	// The account balance of a SoftLayer customer account. An account's balance is the amount of money owed to SoftLayer by the account holder, returned as a floating point number with two decimal places, measured in US Dollars ($USD). A negative account balance means the account holder has overpaid and is owed money by SoftLayer.
	Balance *Float64 `json:"balance,omitempty" xmlrpc:"balance"`

	// A count of the bandwidth allotments for an account.
	BandwidthAllotmentCount *uint `json:"bandwidthAllotmentCount,omitempty" xmlrpc:"bandwidthAllotmentCount"`

	// The bandwidth allotments for an account.
	BandwidthAllotments []Network_Bandwidth_Version1_Allotment `json:"bandwidthAllotments,omitempty" xmlrpc:"bandwidthAllotments"`

	// The bandwidth allotments for an account currently over allocation.
	BandwidthAllotmentsOverAllocation []Network_Bandwidth_Version1_Allotment `json:"bandwidthAllotmentsOverAllocation,omitempty" xmlrpc:"bandwidthAllotmentsOverAllocation"`

	// A count of the bandwidth allotments for an account currently over allocation.
	BandwidthAllotmentsOverAllocationCount *uint `json:"bandwidthAllotmentsOverAllocationCount,omitempty" xmlrpc:"bandwidthAllotmentsOverAllocationCount"`

	// The bandwidth allotments for an account projected to go over allocation.
	BandwidthAllotmentsProjectedOverAllocation []Network_Bandwidth_Version1_Allotment `json:"bandwidthAllotmentsProjectedOverAllocation,omitempty" xmlrpc:"bandwidthAllotmentsProjectedOverAllocation"`

	// A count of the bandwidth allotments for an account projected to go over allocation.
	BandwidthAllotmentsProjectedOverAllocationCount *uint `json:"bandwidthAllotmentsProjectedOverAllocationCount,omitempty" xmlrpc:"bandwidthAllotmentsProjectedOverAllocationCount"`

	// A count of an account's associated bare metal server objects.
	BareMetalInstanceCount *uint `json:"bareMetalInstanceCount,omitempty" xmlrpc:"bareMetalInstanceCount"`

	// An account's associated bare metal server objects.
	BareMetalInstances []Hardware `json:"bareMetalInstances,omitempty" xmlrpc:"bareMetalInstances"`

	// A count of all billing agreements for an account
	BillingAgreementCount *uint `json:"billingAgreementCount,omitempty" xmlrpc:"billingAgreementCount"`

	// All billing agreements for an account
	BillingAgreements []Account_Agreement `json:"billingAgreements,omitempty" xmlrpc:"billingAgreements"`

	// An account's billing information.
	BillingInfo *Billing_Info `json:"billingInfo,omitempty" xmlrpc:"billingInfo"`

	// A count of private template group objects (parent and children) and the shared template group objects (parent only) for an account.
	BlockDeviceTemplateGroupCount *uint `json:"blockDeviceTemplateGroupCount,omitempty" xmlrpc:"blockDeviceTemplateGroupCount"`

	// Private template group objects (parent and children) and the shared template group objects (parent only) for an account.
	BlockDeviceTemplateGroups []Virtual_Guest_Block_Device_Template_Group `json:"blockDeviceTemplateGroups,omitempty" xmlrpc:"blockDeviceTemplateGroups"`

	// Indicates whether this account requires blue id authentication.
	BlueIdAuthenticationRequiredFlag *bool `json:"blueIdAuthenticationRequiredFlag,omitempty" xmlrpc:"blueIdAuthenticationRequiredFlag"`

	// Returns true if this account is linked to IBM Bluemix, false if not.
	BluemixLinkedFlag *bool `json:"bluemixLinkedFlag,omitempty" xmlrpc:"bluemixLinkedFlag"`

	// no documentation yet
	Brand *Brand `json:"brand,omitempty" xmlrpc:"brand"`

	// no documentation yet
	BrandAccountFlag *bool `json:"brandAccountFlag,omitempty" xmlrpc:"brandAccountFlag"`

	// The Brand tied to an account.
	BrandId *int `json:"brandId,omitempty" xmlrpc:"brandId"`

	// The brand keyName.
	BrandKeyName *string `json:"brandKeyName,omitempty" xmlrpc:"brandKeyName"`

	// Indicating whether this account can order additional Vlans.
	CanOrderAdditionalVlansFlag *bool `json:"canOrderAdditionalVlansFlag,omitempty" xmlrpc:"canOrderAdditionalVlansFlag"`

	// A count of an account's active carts.
	CartCount *uint `json:"cartCount,omitempty" xmlrpc:"cartCount"`

	// An account's active carts.
	Carts []Billing_Order_Quote `json:"carts,omitempty" xmlrpc:"carts"`

	// A count of
	CatalystEnrollmentCount *uint `json:"catalystEnrollmentCount,omitempty" xmlrpc:"catalystEnrollmentCount"`

	// no documentation yet
	CatalystEnrollments []Catalyst_Enrollment `json:"catalystEnrollments,omitempty" xmlrpc:"catalystEnrollments"`

	// A count of an account's associated CDN accounts.
	CdnAccountCount *uint `json:"cdnAccountCount,omitempty" xmlrpc:"cdnAccountCount"`

	// An account's associated CDN accounts.
	CdnAccounts []Network_ContentDelivery_Account `json:"cdnAccounts,omitempty" xmlrpc:"cdnAccounts"`

	// The city of the mailing address belonging to an account.
	City *string `json:"city,omitempty" xmlrpc:"city"`

	// Whether an account is exempt from taxes on their invoices.
	ClaimedTaxExemptTxFlag *bool `json:"claimedTaxExemptTxFlag,omitempty" xmlrpc:"claimedTaxExemptTxFlag"`

	// A count of all closed tickets associated with an account.
	ClosedTicketCount *uint `json:"closedTicketCount,omitempty" xmlrpc:"closedTicketCount"`

	// All closed tickets associated with an account.
	ClosedTickets []Ticket `json:"closedTickets,omitempty" xmlrpc:"closedTickets"`

	// The company name associated with an account.
	CompanyName *string `json:"companyName,omitempty" xmlrpc:"companyName"`

	// A two-letter abbreviation of the country in the mailing address belonging to an account.
	Country *string `json:"country,omitempty" xmlrpc:"country"`

	// The date an account was created.
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// A count of datacenters which contain subnets that the account has access to route.
	DatacentersWithSubnetAllocationCount *uint `json:"datacentersWithSubnetAllocationCount,omitempty" xmlrpc:"datacentersWithSubnetAllocationCount"`

	// Datacenters which contain subnets that the account has access to route.
	DatacentersWithSubnetAllocations []Location `json:"datacentersWithSubnetAllocations,omitempty" xmlrpc:"datacentersWithSubnetAllocations"`

	// Device Fingerprint Identifier - Used internally and can safely be ignored.
	DeviceFingerprintId *string `json:"deviceFingerprintId,omitempty" xmlrpc:"deviceFingerprintId"`

	// A flag indicating whether payments are processed for this account.
	DisablePaymentProcessingFlag *bool `json:"disablePaymentProcessingFlag,omitempty" xmlrpc:"disablePaymentProcessingFlag"`

	// A count of the SoftLayer employees that an account is assigned to.
	DisplaySupportRepresentativeAssignmentCount *uint `json:"displaySupportRepresentativeAssignmentCount,omitempty" xmlrpc:"displaySupportRepresentativeAssignmentCount"`

	// The SoftLayer employees that an account is assigned to.
	DisplaySupportRepresentativeAssignments []Account_Attachment_Employee `json:"displaySupportRepresentativeAssignments,omitempty" xmlrpc:"displaySupportRepresentativeAssignments"`

	// A count of the DNS domains associated with an account.
	DomainCount *uint `json:"domainCount,omitempty" xmlrpc:"domainCount"`

	// A count of
	DomainRegistrationCount *uint `json:"domainRegistrationCount,omitempty" xmlrpc:"domainRegistrationCount"`

	// no documentation yet
	DomainRegistrations []Dns_Domain_Registration `json:"domainRegistrations,omitempty" xmlrpc:"domainRegistrations"`

	// The DNS domains associated with an account.
	Domains []Dns_Domain `json:"domains,omitempty" xmlrpc:"domains"`

	// A count of the DNS domains associated with an account that were not created as a result of a secondary DNS zone transfer.
	DomainsWithoutSecondaryDnsRecordCount *uint `json:"domainsWithoutSecondaryDnsRecordCount,omitempty" xmlrpc:"domainsWithoutSecondaryDnsRecordCount"`

	// The DNS domains associated with an account that were not created as a result of a secondary DNS zone transfer.
	DomainsWithoutSecondaryDnsRecords []Dns_Domain `json:"domainsWithoutSecondaryDnsRecords,omitempty" xmlrpc:"domainsWithoutSecondaryDnsRecords"`

	// A general email address assigned to an account.
	Email *string `json:"email,omitempty" xmlrpc:"email"`

	// The total capacity of Legacy EVault Volumes on an account, in GB.
	EvaultCapacityGB *uint `json:"evaultCapacityGB,omitempty" xmlrpc:"evaultCapacityGB"`

	// A count of an account's master EVault user. This is only used when an account has EVault service.
	EvaultMasterUserCount *uint `json:"evaultMasterUserCount,omitempty" xmlrpc:"evaultMasterUserCount"`

	// An account's master EVault user. This is only used when an account has EVault service.
	EvaultMasterUsers []Account_Password `json:"evaultMasterUsers,omitempty" xmlrpc:"evaultMasterUsers"`

	// An account's associated EVault storage volumes.
	EvaultNetworkStorage []Network_Storage `json:"evaultNetworkStorage,omitempty" xmlrpc:"evaultNetworkStorage"`

	// A count of an account's associated EVault storage volumes.
	EvaultNetworkStorageCount *uint `json:"evaultNetworkStorageCount,omitempty" xmlrpc:"evaultNetworkStorageCount"`

	// A count of stored security certificates that are expired (ie. SSL)
	ExpiredSecurityCertificateCount *uint `json:"expiredSecurityCertificateCount,omitempty" xmlrpc:"expiredSecurityCertificateCount"`

	// Stored security certificates that are expired (ie. SSL)
	ExpiredSecurityCertificates []Security_Certificate `json:"expiredSecurityCertificates,omitempty" xmlrpc:"expiredSecurityCertificates"`

	// A count of logs of who entered a colocation area which is assigned to this account, or when a user under this account enters a datacenter.
	FacilityLogCount *uint `json:"facilityLogCount,omitempty" xmlrpc:"facilityLogCount"`

	// Logs of who entered a colocation area which is assigned to this account, or when a user under this account enters a datacenter.
	FacilityLogs []User_Access_Facility_Log `json:"facilityLogs,omitempty" xmlrpc:"facilityLogs"`

	// A fax phone number assigned to an account.
	FaxPhone *string `json:"faxPhone,omitempty" xmlrpc:"faxPhone"`

	// Each customer account is listed under a single individual. This is that individual's first name.
	FirstName *string `json:"firstName,omitempty" xmlrpc:"firstName"`

	// A count of all of the account's current and former Flexible Credit enrollments.
	FlexibleCreditEnrollmentCount *uint `json:"flexibleCreditEnrollmentCount,omitempty" xmlrpc:"flexibleCreditEnrollmentCount"`

	// All of the account's current and former Flexible Credit enrollments.
	FlexibleCreditEnrollments []FlexibleCredit_Enrollment `json:"flexibleCreditEnrollments,omitempty" xmlrpc:"flexibleCreditEnrollments"`

	// A count of
	GlobalIpRecordCount *uint `json:"globalIpRecordCount,omitempty" xmlrpc:"globalIpRecordCount"`

	// no documentation yet
	GlobalIpRecords []Network_Subnet_IpAddress_Global `json:"globalIpRecords,omitempty" xmlrpc:"globalIpRecords"`

	// A count of
	GlobalIpv4RecordCount *uint `json:"globalIpv4RecordCount,omitempty" xmlrpc:"globalIpv4RecordCount"`

	// no documentation yet
	GlobalIpv4Records []Network_Subnet_IpAddress_Global `json:"globalIpv4Records,omitempty" xmlrpc:"globalIpv4Records"`

	// A count of
	GlobalIpv6RecordCount *uint `json:"globalIpv6RecordCount,omitempty" xmlrpc:"globalIpv6RecordCount"`

	// no documentation yet
	GlobalIpv6Records []Network_Subnet_IpAddress_Global `json:"globalIpv6Records,omitempty" xmlrpc:"globalIpv6Records"`

	// A count of the global load balancer accounts for a softlayer customer account.
	GlobalLoadBalancerAccountCount *uint `json:"globalLoadBalancerAccountCount,omitempty" xmlrpc:"globalLoadBalancerAccountCount"`

	// The global load balancer accounts for a softlayer customer account.
	GlobalLoadBalancerAccounts []Network_LoadBalancer_Global_Account `json:"globalLoadBalancerAccounts,omitempty" xmlrpc:"globalLoadBalancerAccounts"`

	// An account's associated hardware objects.
	Hardware []Hardware `json:"hardware,omitempty" xmlrpc:"hardware"`

	// A count of an account's associated hardware objects.
	HardwareCount *uint `json:"hardwareCount,omitempty" xmlrpc:"hardwareCount"`

	// An account's associated hardware objects currently over bandwidth allocation.
	HardwareOverBandwidthAllocation []Hardware `json:"hardwareOverBandwidthAllocation,omitempty" xmlrpc:"hardwareOverBandwidthAllocation"`

	// A count of an account's associated hardware objects currently over bandwidth allocation.
	HardwareOverBandwidthAllocationCount *uint `json:"hardwareOverBandwidthAllocationCount,omitempty" xmlrpc:"hardwareOverBandwidthAllocationCount"`

	// An account's associated hardware objects projected to go over bandwidth allocation.
	HardwareProjectedOverBandwidthAllocation []Hardware `json:"hardwareProjectedOverBandwidthAllocation,omitempty" xmlrpc:"hardwareProjectedOverBandwidthAllocation"`

	// A count of an account's associated hardware objects projected to go over bandwidth allocation.
	HardwareProjectedOverBandwidthAllocationCount *uint `json:"hardwareProjectedOverBandwidthAllocationCount,omitempty" xmlrpc:"hardwareProjectedOverBandwidthAllocationCount"`

	// All hardware associated with an account that has the cPanel web hosting control panel installed.
	HardwareWithCpanel []Hardware `json:"hardwareWithCpanel,omitempty" xmlrpc:"hardwareWithCpanel"`

	// A count of all hardware associated with an account that has the cPanel web hosting control panel installed.
	HardwareWithCpanelCount *uint `json:"hardwareWithCpanelCount,omitempty" xmlrpc:"hardwareWithCpanelCount"`

	// All hardware associated with an account that has the Helm web hosting control panel installed.
	HardwareWithHelm []Hardware `json:"hardwareWithHelm,omitempty" xmlrpc:"hardwareWithHelm"`

	// A count of all hardware associated with an account that has the Helm web hosting control panel installed.
	HardwareWithHelmCount *uint `json:"hardwareWithHelmCount,omitempty" xmlrpc:"hardwareWithHelmCount"`

	// All hardware associated with an account that has McAfee Secure software components.
	HardwareWithMcafee []Hardware `json:"hardwareWithMcafee,omitempty" xmlrpc:"hardwareWithMcafee"`

	// All hardware associated with an account that has McAfee Secure AntiVirus for Redhat software components.
	HardwareWithMcafeeAntivirusRedhat []Hardware `json:"hardwareWithMcafeeAntivirusRedhat,omitempty" xmlrpc:"hardwareWithMcafeeAntivirusRedhat"`

	// A count of all hardware associated with an account that has McAfee Secure AntiVirus for Redhat software components.
	HardwareWithMcafeeAntivirusRedhatCount *uint `json:"hardwareWithMcafeeAntivirusRedhatCount,omitempty" xmlrpc:"hardwareWithMcafeeAntivirusRedhatCount"`

	// A count of all hardware associated with an account that has McAfee Secure AntiVirus for Windows software components.
	HardwareWithMcafeeAntivirusWindowCount *uint `json:"hardwareWithMcafeeAntivirusWindowCount,omitempty" xmlrpc:"hardwareWithMcafeeAntivirusWindowCount"`

	// All hardware associated with an account that has McAfee Secure AntiVirus for Windows software components.
	HardwareWithMcafeeAntivirusWindows []Hardware `json:"hardwareWithMcafeeAntivirusWindows,omitempty" xmlrpc:"hardwareWithMcafeeAntivirusWindows"`

	// A count of all hardware associated with an account that has McAfee Secure software components.
	HardwareWithMcafeeCount *uint `json:"hardwareWithMcafeeCount,omitempty" xmlrpc:"hardwareWithMcafeeCount"`

	// All hardware associated with an account that has McAfee Secure Intrusion Detection System software components.
	HardwareWithMcafeeIntrusionDetectionSystem []Hardware `json:"hardwareWithMcafeeIntrusionDetectionSystem,omitempty" xmlrpc:"hardwareWithMcafeeIntrusionDetectionSystem"`

	// A count of all hardware associated with an account that has McAfee Secure Intrusion Detection System software components.
	HardwareWithMcafeeIntrusionDetectionSystemCount *uint `json:"hardwareWithMcafeeIntrusionDetectionSystemCount,omitempty" xmlrpc:"hardwareWithMcafeeIntrusionDetectionSystemCount"`

	// All hardware associated with an account that has the Plesk web hosting control panel installed.
	HardwareWithPlesk []Hardware `json:"hardwareWithPlesk,omitempty" xmlrpc:"hardwareWithPlesk"`

	// A count of all hardware associated with an account that has the Plesk web hosting control panel installed.
	HardwareWithPleskCount *uint `json:"hardwareWithPleskCount,omitempty" xmlrpc:"hardwareWithPleskCount"`

	// All hardware associated with an account that has the QuantaStor storage system installed.
	HardwareWithQuantastor []Hardware `json:"hardwareWithQuantastor,omitempty" xmlrpc:"hardwareWithQuantastor"`

	// A count of all hardware associated with an account that has the QuantaStor storage system installed.
	HardwareWithQuantastorCount *uint `json:"hardwareWithQuantastorCount,omitempty" xmlrpc:"hardwareWithQuantastorCount"`

	// All hardware associated with an account that has the Urchin web traffic analytics package installed.
	HardwareWithUrchin []Hardware `json:"hardwareWithUrchin,omitempty" xmlrpc:"hardwareWithUrchin"`

	// A count of all hardware associated with an account that has the Urchin web traffic analytics package installed.
	HardwareWithUrchinCount *uint `json:"hardwareWithUrchinCount,omitempty" xmlrpc:"hardwareWithUrchinCount"`

	// A count of all hardware associated with an account that is running a version of the Microsoft Windows operating system.
	HardwareWithWindowCount *uint `json:"hardwareWithWindowCount,omitempty" xmlrpc:"hardwareWithWindowCount"`

	// All hardware associated with an account that is running a version of the Microsoft Windows operating system.
	HardwareWithWindows []Hardware `json:"hardwareWithWindows,omitempty" xmlrpc:"hardwareWithWindows"`

	// Return 1 if one of the account's hardware has the EVault Bare Metal Server Restore Plugin otherwise 0.
	HasEvaultBareMetalRestorePluginFlag *bool `json:"hasEvaultBareMetalRestorePluginFlag,omitempty" xmlrpc:"hasEvaultBareMetalRestorePluginFlag"`

	// Return 1 if one of the account's hardware has an installation of Idera Server Backup otherwise 0.
	HasIderaBareMetalRestorePluginFlag *bool `json:"hasIderaBareMetalRestorePluginFlag,omitempty" xmlrpc:"hasIderaBareMetalRestorePluginFlag"`

	// The number of orders in a PENDING status for a SoftLayer customer account.
	HasPendingOrder *uint `json:"hasPendingOrder,omitempty" xmlrpc:"hasPendingOrder"`

	// Return 1 if one of the account's hardware has an installation of R1Soft CDP otherwise 0.
	HasR1softBareMetalRestorePluginFlag *bool `json:"hasR1softBareMetalRestorePluginFlag,omitempty" xmlrpc:"hasR1softBareMetalRestorePluginFlag"`

	// A count of an account's associated hourly bare metal server objects.
	HourlyBareMetalInstanceCount *uint `json:"hourlyBareMetalInstanceCount,omitempty" xmlrpc:"hourlyBareMetalInstanceCount"`

	// An account's associated hourly bare metal server objects.
	HourlyBareMetalInstances []Hardware `json:"hourlyBareMetalInstances,omitempty" xmlrpc:"hourlyBareMetalInstances"`

	// A count of hourly service billing items that will be on an account's next invoice.
	HourlyServiceBillingItemCount *uint `json:"hourlyServiceBillingItemCount,omitempty" xmlrpc:"hourlyServiceBillingItemCount"`

	// Hourly service billing items that will be on an account's next invoice.
	HourlyServiceBillingItems []Billing_Item `json:"hourlyServiceBillingItems,omitempty" xmlrpc:"hourlyServiceBillingItems"`

	// A count of an account's associated hourly virtual guest objects.
	HourlyVirtualGuestCount *uint `json:"hourlyVirtualGuestCount,omitempty" xmlrpc:"hourlyVirtualGuestCount"`

	// An account's associated hourly virtual guest objects.
	HourlyVirtualGuests []Virtual_Guest `json:"hourlyVirtualGuests,omitempty" xmlrpc:"hourlyVirtualGuests"`

	// An account's associated Virtual Storage volumes.
	HubNetworkStorage []Network_Storage `json:"hubNetworkStorage,omitempty" xmlrpc:"hubNetworkStorage"`

	// A count of an account's associated Virtual Storage volumes.
	HubNetworkStorageCount *uint `json:"hubNetworkStorageCount,omitempty" xmlrpc:"hubNetworkStorageCount"`

	// A customer account's internal identifier. Account numbers are typically preceded by the string "SL" in the customer portal. Every SoftLayer account has at least one portal user whose username follows the "SL" + account number naming scheme.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// A count of
	InternalNoteCount *uint `json:"internalNoteCount,omitempty" xmlrpc:"internalNoteCount"`

	// no documentation yet
	InternalNotes []Account_Note `json:"internalNotes,omitempty" xmlrpc:"internalNotes"`

	// A count of an account's associated billing invoices.
	InvoiceCount *uint `json:"invoiceCount,omitempty" xmlrpc:"invoiceCount"`

	// An account's associated billing invoices.
	Invoices []Billing_Invoice `json:"invoices,omitempty" xmlrpc:"invoices"`

	// A count of
	IpAddressCount *uint `json:"ipAddressCount,omitempty" xmlrpc:"ipAddressCount"`

	// no documentation yet
	IpAddresses []Network_Subnet_IpAddress `json:"ipAddresses,omitempty" xmlrpc:"ipAddresses"`

	// A flag indicating if an account belongs to a reseller or not.
	IsReseller *int `json:"isReseller,omitempty" xmlrpc:"isReseller"`

	// An account's associated iSCSI storage volumes.
	IscsiNetworkStorage []Network_Storage `json:"iscsiNetworkStorage,omitempty" xmlrpc:"iscsiNetworkStorage"`

	// A count of an account's associated iSCSI storage volumes.
	IscsiNetworkStorageCount *uint `json:"iscsiNetworkStorageCount,omitempty" xmlrpc:"iscsiNetworkStorageCount"`

	// The most recently canceled billing item.
	LastCanceledBillingItem *Billing_Item `json:"lastCanceledBillingItem,omitempty" xmlrpc:"lastCanceledBillingItem"`

	// The most recent cancelled server billing item.
	LastCancelledServerBillingItem *Billing_Item `json:"lastCancelledServerBillingItem,omitempty" xmlrpc:"lastCancelledServerBillingItem"`

	// A count of the five most recently closed abuse tickets associated with an account.
	LastFiveClosedAbuseTicketCount *uint `json:"lastFiveClosedAbuseTicketCount,omitempty" xmlrpc:"lastFiveClosedAbuseTicketCount"`

	// The five most recently closed abuse tickets associated with an account.
	LastFiveClosedAbuseTickets []Ticket `json:"lastFiveClosedAbuseTickets,omitempty" xmlrpc:"lastFiveClosedAbuseTickets"`

	// A count of the five most recently closed accounting tickets associated with an account.
	LastFiveClosedAccountingTicketCount *uint `json:"lastFiveClosedAccountingTicketCount,omitempty" xmlrpc:"lastFiveClosedAccountingTicketCount"`

	// The five most recently closed accounting tickets associated with an account.
	LastFiveClosedAccountingTickets []Ticket `json:"lastFiveClosedAccountingTickets,omitempty" xmlrpc:"lastFiveClosedAccountingTickets"`

	// A count of the five most recently closed tickets that do not belong to the abuse, accounting, sales, or support groups associated with an account.
	LastFiveClosedOtherTicketCount *uint `json:"lastFiveClosedOtherTicketCount,omitempty" xmlrpc:"lastFiveClosedOtherTicketCount"`

	// The five most recently closed tickets that do not belong to the abuse, accounting, sales, or support groups associated with an account.
	LastFiveClosedOtherTickets []Ticket `json:"lastFiveClosedOtherTickets,omitempty" xmlrpc:"lastFiveClosedOtherTickets"`

	// A count of the five most recently closed sales tickets associated with an account.
	LastFiveClosedSalesTicketCount *uint `json:"lastFiveClosedSalesTicketCount,omitempty" xmlrpc:"lastFiveClosedSalesTicketCount"`

	// The five most recently closed sales tickets associated with an account.
	LastFiveClosedSalesTickets []Ticket `json:"lastFiveClosedSalesTickets,omitempty" xmlrpc:"lastFiveClosedSalesTickets"`

	// A count of the five most recently closed support tickets associated with an account.
	LastFiveClosedSupportTicketCount *uint `json:"lastFiveClosedSupportTicketCount,omitempty" xmlrpc:"lastFiveClosedSupportTicketCount"`

	// The five most recently closed support tickets associated with an account.
	LastFiveClosedSupportTickets []Ticket `json:"lastFiveClosedSupportTickets,omitempty" xmlrpc:"lastFiveClosedSupportTickets"`

	// A count of the five most recently closed tickets associated with an account.
	LastFiveClosedTicketCount *uint `json:"lastFiveClosedTicketCount,omitempty" xmlrpc:"lastFiveClosedTicketCount"`

	// The five most recently closed tickets associated with an account.
	LastFiveClosedTickets []Ticket `json:"lastFiveClosedTickets,omitempty" xmlrpc:"lastFiveClosedTickets"`

	// Each customer account is listed under a single individual. This is that individual's last name.
	LastName *string `json:"lastName,omitempty" xmlrpc:"lastName"`

	// Whether an account has late fee protection.
	LateFeeProtectionFlag *bool `json:"lateFeeProtectionFlag,omitempty" xmlrpc:"lateFeeProtectionFlag"`

	// An account's most recent billing date.
	LatestBillDate *Time `json:"latestBillDate,omitempty" xmlrpc:"latestBillDate"`

	// An account's latest recurring invoice.
	LatestRecurringInvoice *Billing_Invoice `json:"latestRecurringInvoice,omitempty" xmlrpc:"latestRecurringInvoice"`

	// An account's latest recurring pending invoice.
	LatestRecurringPendingInvoice *Billing_Invoice `json:"latestRecurringPendingInvoice,omitempty" xmlrpc:"latestRecurringPendingInvoice"`

	// A count of the legacy bandwidth allotments for an account.
	LegacyBandwidthAllotmentCount *uint `json:"legacyBandwidthAllotmentCount,omitempty" xmlrpc:"legacyBandwidthAllotmentCount"`

	// The legacy bandwidth allotments for an account.
	LegacyBandwidthAllotments []Network_Bandwidth_Version1_Allotment `json:"legacyBandwidthAllotments,omitempty" xmlrpc:"legacyBandwidthAllotments"`

	// The total capacity of Legacy iSCSI Volumes on an account, in GB.
	LegacyIscsiCapacityGB *uint `json:"legacyIscsiCapacityGB,omitempty" xmlrpc:"legacyIscsiCapacityGB"`

	// A count of an account's associated load balancers.
	LoadBalancerCount *uint `json:"loadBalancerCount,omitempty" xmlrpc:"loadBalancerCount"`

	// An account's associated load balancers.
	LoadBalancers []Network_LoadBalancer_VirtualIpAddress `json:"loadBalancers,omitempty" xmlrpc:"loadBalancers"`

	// The total capacity of Legacy lockbox Volumes on an account, in GB.
	LockboxCapacityGB *uint `json:"lockboxCapacityGB,omitempty" xmlrpc:"lockboxCapacityGB"`

	// An account's associated Lockbox storage volumes.
	LockboxNetworkStorage []Network_Storage `json:"lockboxNetworkStorage,omitempty" xmlrpc:"lockboxNetworkStorage"`

	// A count of an account's associated Lockbox storage volumes.
	LockboxNetworkStorageCount *uint `json:"lockboxNetworkStorageCount,omitempty" xmlrpc:"lockboxNetworkStorageCount"`

	// no documentation yet
	ManualPaymentsUnderReview []Billing_Payment_Card_ManualPayment `json:"manualPaymentsUnderReview,omitempty" xmlrpc:"manualPaymentsUnderReview"`

	// A count of
	ManualPaymentsUnderReviewCount *uint `json:"manualPaymentsUnderReviewCount,omitempty" xmlrpc:"manualPaymentsUnderReviewCount"`

	// An account's master user.
	MasterUser *User_Customer `json:"masterUser,omitempty" xmlrpc:"masterUser"`

	// A count of an account's media transfer service requests.
	MediaDataTransferRequestCount *uint `json:"mediaDataTransferRequestCount,omitempty" xmlrpc:"mediaDataTransferRequestCount"`

	// An account's media transfer service requests.
	MediaDataTransferRequests []Account_Media_Data_Transfer_Request `json:"mediaDataTransferRequests,omitempty" xmlrpc:"mediaDataTransferRequests"`

	// A count of an account's associated Message Queue accounts.
	MessageQueueAccountCount *uint `json:"messageQueueAccountCount,omitempty" xmlrpc:"messageQueueAccountCount"`

	// An account's associated Message Queue accounts.
	MessageQueueAccounts []Network_Message_Queue `json:"messageQueueAccounts,omitempty" xmlrpc:"messageQueueAccounts"`

	// The date an account was last modified.
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// A count of an account's associated monthly bare metal server objects.
	MonthlyBareMetalInstanceCount *uint `json:"monthlyBareMetalInstanceCount,omitempty" xmlrpc:"monthlyBareMetalInstanceCount"`

	// An account's associated monthly bare metal server objects.
	MonthlyBareMetalInstances []Hardware `json:"monthlyBareMetalInstances,omitempty" xmlrpc:"monthlyBareMetalInstances"`

	// A count of an account's associated monthly virtual guest objects.
	MonthlyVirtualGuestCount *uint `json:"monthlyVirtualGuestCount,omitempty" xmlrpc:"monthlyVirtualGuestCount"`

	// An account's associated monthly virtual guest objects.
	MonthlyVirtualGuests []Virtual_Guest `json:"monthlyVirtualGuests,omitempty" xmlrpc:"monthlyVirtualGuests"`

	// An account's associated NAS storage volumes.
	NasNetworkStorage []Network_Storage `json:"nasNetworkStorage,omitempty" xmlrpc:"nasNetworkStorage"`

	// A count of an account's associated NAS storage volumes.
	NasNetworkStorageCount *uint `json:"nasNetworkStorageCount,omitempty" xmlrpc:"nasNetworkStorageCount"`

	// Whether or not this account can define their own networks.
	NetworkCreationFlag *bool `json:"networkCreationFlag,omitempty" xmlrpc:"networkCreationFlag"`

	// A count of all network gateway devices on this account.
	NetworkGatewayCount *uint `json:"networkGatewayCount,omitempty" xmlrpc:"networkGatewayCount"`

	// All network gateway devices on this account.
	NetworkGateways []Network_Gateway `json:"networkGateways,omitempty" xmlrpc:"networkGateways"`

	// An account's associated network hardware.
	NetworkHardware []Hardware `json:"networkHardware,omitempty" xmlrpc:"networkHardware"`

	// A count of an account's associated network hardware.
	NetworkHardwareCount *uint `json:"networkHardwareCount,omitempty" xmlrpc:"networkHardwareCount"`

	// A count of
	NetworkMessageDeliveryAccountCount *uint `json:"networkMessageDeliveryAccountCount,omitempty" xmlrpc:"networkMessageDeliveryAccountCount"`

	// no documentation yet
	NetworkMessageDeliveryAccounts []Network_Message_Delivery `json:"networkMessageDeliveryAccounts,omitempty" xmlrpc:"networkMessageDeliveryAccounts"`

	// Hardware which is currently experiencing a service failure.
	NetworkMonitorDownHardware []Hardware `json:"networkMonitorDownHardware,omitempty" xmlrpc:"networkMonitorDownHardware"`

	// A count of hardware which is currently experiencing a service failure.
	NetworkMonitorDownHardwareCount *uint `json:"networkMonitorDownHardwareCount,omitempty" xmlrpc:"networkMonitorDownHardwareCount"`

	// A count of virtual guest which is currently experiencing a service failure.
	NetworkMonitorDownVirtualGuestCount *uint `json:"networkMonitorDownVirtualGuestCount,omitempty" xmlrpc:"networkMonitorDownVirtualGuestCount"`

	// Virtual guest which is currently experiencing a service failure.
	NetworkMonitorDownVirtualGuests []Virtual_Guest `json:"networkMonitorDownVirtualGuests,omitempty" xmlrpc:"networkMonitorDownVirtualGuests"`

	// Hardware which is currently recovering from a service failure.
	NetworkMonitorRecoveringHardware []Hardware `json:"networkMonitorRecoveringHardware,omitempty" xmlrpc:"networkMonitorRecoveringHardware"`

	// A count of hardware which is currently recovering from a service failure.
	NetworkMonitorRecoveringHardwareCount *uint `json:"networkMonitorRecoveringHardwareCount,omitempty" xmlrpc:"networkMonitorRecoveringHardwareCount"`

	// A count of virtual guest which is currently recovering from a service failure.
	NetworkMonitorRecoveringVirtualGuestCount *uint `json:"networkMonitorRecoveringVirtualGuestCount,omitempty" xmlrpc:"networkMonitorRecoveringVirtualGuestCount"`

	// Virtual guest which is currently recovering from a service failure.
	NetworkMonitorRecoveringVirtualGuests []Virtual_Guest `json:"networkMonitorRecoveringVirtualGuests,omitempty" xmlrpc:"networkMonitorRecoveringVirtualGuests"`

	// Hardware which is currently online.
	NetworkMonitorUpHardware []Hardware `json:"networkMonitorUpHardware,omitempty" xmlrpc:"networkMonitorUpHardware"`

	// A count of hardware which is currently online.
	NetworkMonitorUpHardwareCount *uint `json:"networkMonitorUpHardwareCount,omitempty" xmlrpc:"networkMonitorUpHardwareCount"`

	// A count of virtual guest which is currently online.
	NetworkMonitorUpVirtualGuestCount *uint `json:"networkMonitorUpVirtualGuestCount,omitempty" xmlrpc:"networkMonitorUpVirtualGuestCount"`

	// Virtual guest which is currently online.
	NetworkMonitorUpVirtualGuests []Virtual_Guest `json:"networkMonitorUpVirtualGuests,omitempty" xmlrpc:"networkMonitorUpVirtualGuests"`

	// An account's associated storage volumes. This includes Lockbox, NAS, EVault, and iSCSI volumes.
	NetworkStorage []Network_Storage `json:"networkStorage,omitempty" xmlrpc:"networkStorage"`

	// A count of an account's associated storage volumes. This includes Lockbox, NAS, EVault, and iSCSI volumes.
	NetworkStorageCount *uint `json:"networkStorageCount,omitempty" xmlrpc:"networkStorageCount"`

	// A count of an account's Network Storage groups.
	NetworkStorageGroupCount *uint `json:"networkStorageGroupCount,omitempty" xmlrpc:"networkStorageGroupCount"`

	// An account's Network Storage groups.
	NetworkStorageGroups []Network_Storage_Group `json:"networkStorageGroups,omitempty" xmlrpc:"networkStorageGroups"`

	// A count of iPSec network tunnels for an account.
	NetworkTunnelContextCount *uint `json:"networkTunnelContextCount,omitempty" xmlrpc:"networkTunnelContextCount"`

	// IPSec network tunnels for an account.
	NetworkTunnelContexts []Network_Tunnel_Module_Context `json:"networkTunnelContexts,omitempty" xmlrpc:"networkTunnelContexts"`

	// A count of all network VLANs assigned to an account.
	NetworkVlanCount *uint `json:"networkVlanCount,omitempty" xmlrpc:"networkVlanCount"`

	// Whether or not an account has automatic private VLAN spanning enabled.
	NetworkVlanSpan *Account_Network_Vlan_Span `json:"networkVlanSpan,omitempty" xmlrpc:"networkVlanSpan"`

	// All network VLANs assigned to an account.
	NetworkVlans []Network_Vlan `json:"networkVlans,omitempty" xmlrpc:"networkVlans"`

	// A count of dEPRECATED - This information can be pulled directly through tapping keys now - DEPRECATED. The allotments for this account and their servers for the next billing cycle. The public inbound and outbound bandwidth is calculated for each server in addition to the daily average network traffic since the last billing date.
	NextBillingPublicAllotmentHardwareBandwidthDetailCount *uint `json:"nextBillingPublicAllotmentHardwareBandwidthDetailCount,omitempty" xmlrpc:"nextBillingPublicAllotmentHardwareBandwidthDetailCount"`

	// DEPRECATED - This information can be pulled directly through tapping keys now - DEPRECATED. The allotments for this account and their servers for the next billing cycle. The public inbound and outbound bandwidth is calculated for each server in addition to the daily average network traffic since the last billing date.
	NextBillingPublicAllotmentHardwareBandwidthDetails []Network_Bandwidth_Version1_Allotment `json:"nextBillingPublicAllotmentHardwareBandwidthDetails,omitempty" xmlrpc:"nextBillingPublicAllotmentHardwareBandwidthDetails"`

	// The pre-tax total amount exempt from incubator credit for the account's next invoice. This field is now deprecated and will soon be removed. Please update all references to instead use nextInvoiceTotalAmount
	NextInvoiceIncubatorExemptTotal *Float64 `json:"nextInvoiceIncubatorExemptTotal,omitempty" xmlrpc:"nextInvoiceIncubatorExemptTotal"`

	// A count of the billing items that will be on an account's next invoice.
	NextInvoiceTopLevelBillingItemCount *uint `json:"nextInvoiceTopLevelBillingItemCount,omitempty" xmlrpc:"nextInvoiceTopLevelBillingItemCount"`

	// The billing items that will be on an account's next invoice.
	NextInvoiceTopLevelBillingItems []Billing_Item `json:"nextInvoiceTopLevelBillingItems,omitempty" xmlrpc:"nextInvoiceTopLevelBillingItems"`

	// The pre-tax total amount of an account's next invoice measured in US Dollars ($USD), assuming no changes or charges occur between now and time of billing.
	NextInvoiceTotalAmount *Float64 `json:"nextInvoiceTotalAmount,omitempty" xmlrpc:"nextInvoiceTotalAmount"`

	// The total one-time charge amount of an account's next invoice measured in US Dollars ($USD), assuming no changes or charges occur between now and time of billing.
	NextInvoiceTotalOneTimeAmount *Float64 `json:"nextInvoiceTotalOneTimeAmount,omitempty" xmlrpc:"nextInvoiceTotalOneTimeAmount"`

	// The total one-time tax amount of an account's next invoice measured in US Dollars ($USD), assuming no changes or charges occur between now and time of billing.
	NextInvoiceTotalOneTimeTaxAmount *Float64 `json:"nextInvoiceTotalOneTimeTaxAmount,omitempty" xmlrpc:"nextInvoiceTotalOneTimeTaxAmount"`

	// The total recurring charge amount of an account's next invoice measured in US Dollars ($USD), assuming no changes or charges occur between now and time of billing.
	NextInvoiceTotalRecurringAmount *Float64 `json:"nextInvoiceTotalRecurringAmount,omitempty" xmlrpc:"nextInvoiceTotalRecurringAmount"`

	// The total recurring charge amount of an account's next invoice measured in US Dollars ($USD), assuming no changes or charges occur between now and time of billing.
	NextInvoiceTotalRecurringAmountBeforeAccountDiscount *Float64 `json:"nextInvoiceTotalRecurringAmountBeforeAccountDiscount,omitempty" xmlrpc:"nextInvoiceTotalRecurringAmountBeforeAccountDiscount"`

	// The total recurring tax amount of an account's next invoice measured in US Dollars ($USD), assuming no changes or charges occur between now and time of billing.
	NextInvoiceTotalRecurringTaxAmount *Float64 `json:"nextInvoiceTotalRecurringTaxAmount,omitempty" xmlrpc:"nextInvoiceTotalRecurringTaxAmount"`

	// The total recurring charge amount of an account's next invoice measured in US Dollars ($USD), assuming no changes or charges occur between now and time of billing.
	NextInvoiceTotalTaxableRecurringAmount *Float64 `json:"nextInvoiceTotalTaxableRecurringAmount,omitempty" xmlrpc:"nextInvoiceTotalTaxableRecurringAmount"`

	// A count of
	NotificationSubscriberCount *uint `json:"notificationSubscriberCount,omitempty" xmlrpc:"notificationSubscriberCount"`

	// no documentation yet
	NotificationSubscribers []Notification_Subscriber `json:"notificationSubscribers,omitempty" xmlrpc:"notificationSubscribers"`

	// An office phone number assigned to an account.
	OfficePhone *string `json:"officePhone,omitempty" xmlrpc:"officePhone"`

	// A count of the open abuse tickets associated with an account.
	OpenAbuseTicketCount *uint `json:"openAbuseTicketCount,omitempty" xmlrpc:"openAbuseTicketCount"`

	// The open abuse tickets associated with an account.
	OpenAbuseTickets []Ticket `json:"openAbuseTickets,omitempty" xmlrpc:"openAbuseTickets"`

	// A count of the open accounting tickets associated with an account.
	OpenAccountingTicketCount *uint `json:"openAccountingTicketCount,omitempty" xmlrpc:"openAccountingTicketCount"`

	// The open accounting tickets associated with an account.
	OpenAccountingTickets []Ticket `json:"openAccountingTickets,omitempty" xmlrpc:"openAccountingTickets"`

	// A count of the open billing tickets associated with an account.
	OpenBillingTicketCount *uint `json:"openBillingTicketCount,omitempty" xmlrpc:"openBillingTicketCount"`

	// The open billing tickets associated with an account.
	OpenBillingTickets []Ticket `json:"openBillingTickets,omitempty" xmlrpc:"openBillingTickets"`

	// A count of an open ticket requesting cancellation of this server, if one exists.
	OpenCancellationRequestCount *uint `json:"openCancellationRequestCount,omitempty" xmlrpc:"openCancellationRequestCount"`

	// An open ticket requesting cancellation of this server, if one exists.
	OpenCancellationRequests []Billing_Item_Cancellation_Request `json:"openCancellationRequests,omitempty" xmlrpc:"openCancellationRequests"`

	// A count of the open tickets that do not belong to the abuse, accounting, sales, or support groups associated with an account.
	OpenOtherTicketCount *uint `json:"openOtherTicketCount,omitempty" xmlrpc:"openOtherTicketCount"`

	// The open tickets that do not belong to the abuse, accounting, sales, or support groups associated with an account.
	OpenOtherTickets []Ticket `json:"openOtherTickets,omitempty" xmlrpc:"openOtherTickets"`

	// A count of an account's recurring invoices.
	OpenRecurringInvoiceCount *uint `json:"openRecurringInvoiceCount,omitempty" xmlrpc:"openRecurringInvoiceCount"`

	// An account's recurring invoices.
	OpenRecurringInvoices []Billing_Invoice `json:"openRecurringInvoices,omitempty" xmlrpc:"openRecurringInvoices"`

	// A count of the open sales tickets associated with an account.
	OpenSalesTicketCount *uint `json:"openSalesTicketCount,omitempty" xmlrpc:"openSalesTicketCount"`

	// The open sales tickets associated with an account.
	OpenSalesTickets []Ticket `json:"openSalesTickets,omitempty" xmlrpc:"openSalesTickets"`

	// A count of
	OpenStackAccountLinkCount *uint `json:"openStackAccountLinkCount,omitempty" xmlrpc:"openStackAccountLinkCount"`

	// no documentation yet
	OpenStackAccountLinks []Account_Link `json:"openStackAccountLinks,omitempty" xmlrpc:"openStackAccountLinks"`

	// An account's associated Openstack related Object Storage accounts.
	OpenStackObjectStorage []Network_Storage `json:"openStackObjectStorage,omitempty" xmlrpc:"openStackObjectStorage"`

	// A count of an account's associated Openstack related Object Storage accounts.
	OpenStackObjectStorageCount *uint `json:"openStackObjectStorageCount,omitempty" xmlrpc:"openStackObjectStorageCount"`

	// A count of the open support tickets associated with an account.
	OpenSupportTicketCount *uint `json:"openSupportTicketCount,omitempty" xmlrpc:"openSupportTicketCount"`

	// The open support tickets associated with an account.
	OpenSupportTickets []Ticket `json:"openSupportTickets,omitempty" xmlrpc:"openSupportTickets"`

	// A count of all open tickets associated with an account.
	OpenTicketCount *uint `json:"openTicketCount,omitempty" xmlrpc:"openTicketCount"`

	// All open tickets associated with an account.
	OpenTickets []Ticket `json:"openTickets,omitempty" xmlrpc:"openTickets"`

	// All open tickets associated with an account last edited by an employee.
	OpenTicketsWaitingOnCustomer []Ticket `json:"openTicketsWaitingOnCustomer,omitempty" xmlrpc:"openTicketsWaitingOnCustomer"`

	// A count of all open tickets associated with an account last edited by an employee.
	OpenTicketsWaitingOnCustomerCount *uint `json:"openTicketsWaitingOnCustomerCount,omitempty" xmlrpc:"openTicketsWaitingOnCustomerCount"`

	// A count of an account's associated billing orders excluding upgrades.
	OrderCount *uint `json:"orderCount,omitempty" xmlrpc:"orderCount"`

	// An account's associated billing orders excluding upgrades.
	Orders []Billing_Order `json:"orders,omitempty" xmlrpc:"orders"`

	// A count of the billing items that have no parent billing item. These are items that don't necessarily belong to a single server.
	OrphanBillingItemCount *uint `json:"orphanBillingItemCount,omitempty" xmlrpc:"orphanBillingItemCount"`

	// The billing items that have no parent billing item. These are items that don't necessarily belong to a single server.
	OrphanBillingItems []Billing_Item `json:"orphanBillingItems,omitempty" xmlrpc:"orphanBillingItems"`

	// A count of
	OwnedBrandCount *uint `json:"ownedBrandCount,omitempty" xmlrpc:"ownedBrandCount"`

	// no documentation yet
	OwnedBrands []Brand `json:"ownedBrands,omitempty" xmlrpc:"ownedBrands"`

	// A count of
	OwnedHardwareGenericComponentModelCount *uint `json:"ownedHardwareGenericComponentModelCount,omitempty" xmlrpc:"ownedHardwareGenericComponentModelCount"`

	// no documentation yet
	OwnedHardwareGenericComponentModels []Hardware_Component_Model_Generic `json:"ownedHardwareGenericComponentModels,omitempty" xmlrpc:"ownedHardwareGenericComponentModels"`

	// A count of
	PaymentProcessorCount *uint `json:"paymentProcessorCount,omitempty" xmlrpc:"paymentProcessorCount"`

	// no documentation yet
	PaymentProcessors []Billing_Payment_Processor `json:"paymentProcessors,omitempty" xmlrpc:"paymentProcessors"`

	// A count of
	PendingEventCount *uint `json:"pendingEventCount,omitempty" xmlrpc:"pendingEventCount"`

	// no documentation yet
	PendingEvents []Notification_Occurrence_Event `json:"pendingEvents,omitempty" xmlrpc:"pendingEvents"`

	// An account's latest open (pending) invoice.
	PendingInvoice *Billing_Invoice `json:"pendingInvoice,omitempty" xmlrpc:"pendingInvoice"`

	// A count of a list of top-level invoice items that are on an account's currently pending invoice.
	PendingInvoiceTopLevelItemCount *uint `json:"pendingInvoiceTopLevelItemCount,omitempty" xmlrpc:"pendingInvoiceTopLevelItemCount"`

	// A list of top-level invoice items that are on an account's currently pending invoice.
	PendingInvoiceTopLevelItems []Billing_Invoice_Item `json:"pendingInvoiceTopLevelItems,omitempty" xmlrpc:"pendingInvoiceTopLevelItems"`

	// The total amount of an account's pending invoice, if one exists.
	PendingInvoiceTotalAmount *Float64 `json:"pendingInvoiceTotalAmount,omitempty" xmlrpc:"pendingInvoiceTotalAmount"`

	// The total one-time charges for an account's pending invoice, if one exists. In other words, it is the sum of one-time charges, setup fees, and labor fees. It does not include taxes.
	PendingInvoiceTotalOneTimeAmount *Float64 `json:"pendingInvoiceTotalOneTimeAmount,omitempty" xmlrpc:"pendingInvoiceTotalOneTimeAmount"`

	// The sum of all the taxes related to one time charges for an account's pending invoice, if one exists.
	PendingInvoiceTotalOneTimeTaxAmount *Float64 `json:"pendingInvoiceTotalOneTimeTaxAmount,omitempty" xmlrpc:"pendingInvoiceTotalOneTimeTaxAmount"`

	// The total recurring amount of an account's pending invoice, if one exists.
	PendingInvoiceTotalRecurringAmount *Float64 `json:"pendingInvoiceTotalRecurringAmount,omitempty" xmlrpc:"pendingInvoiceTotalRecurringAmount"`

	// The total amount of the recurring taxes on an account's pending invoice, if one exists.
	PendingInvoiceTotalRecurringTaxAmount *Float64 `json:"pendingInvoiceTotalRecurringTaxAmount,omitempty" xmlrpc:"pendingInvoiceTotalRecurringTaxAmount"`

	// A count of an account's permission groups.
	PermissionGroupCount *uint `json:"permissionGroupCount,omitempty" xmlrpc:"permissionGroupCount"`

	// An account's permission groups.
	PermissionGroups []User_Permission_Group `json:"permissionGroups,omitempty" xmlrpc:"permissionGroups"`

	// A count of an account's user roles.
	PermissionRoleCount *uint `json:"permissionRoleCount,omitempty" xmlrpc:"permissionRoleCount"`

	// An account's user roles.
	PermissionRoles []User_Permission_Role `json:"permissionRoles,omitempty" xmlrpc:"permissionRoles"`

	// A count of
	PortableStorageVolumeCount *uint `json:"portableStorageVolumeCount,omitempty" xmlrpc:"portableStorageVolumeCount"`

	// no documentation yet
	PortableStorageVolumes []Virtual_Disk_Image `json:"portableStorageVolumes,omitempty" xmlrpc:"portableStorageVolumes"`

	// A count of customer specified URIs that are downloaded onto a newly provisioned or reloaded server. If the URI is sent over https it will be executed directly on the server.
	PostProvisioningHookCount *uint `json:"postProvisioningHookCount,omitempty" xmlrpc:"postProvisioningHookCount"`

	// Customer specified URIs that are downloaded onto a newly provisioned or reloaded server. If the URI is sent over https it will be executed directly on the server.
	PostProvisioningHooks []Provisioning_Hook `json:"postProvisioningHooks,omitempty" xmlrpc:"postProvisioningHooks"`

	// The postal code of the mailing address belonging to an account.
	PostalCode *string `json:"postalCode,omitempty" xmlrpc:"postalCode"`

	// A count of an account's associated portal users with PPTP VPN access.
	PptpVpnUserCount *uint `json:"pptpVpnUserCount,omitempty" xmlrpc:"pptpVpnUserCount"`

	// An account's associated portal users with PPTP VPN access.
	PptpVpnUsers []User_Customer `json:"pptpVpnUsers,omitempty" xmlrpc:"pptpVpnUsers"`

	// The total recurring amount for an accounts previous revenue.
	PreviousRecurringRevenue *Float64 `json:"previousRecurringRevenue,omitempty" xmlrpc:"previousRecurringRevenue"`

	// A count of the item price that an account is restricted to.
	PriceRestrictionCount *uint `json:"priceRestrictionCount,omitempty" xmlrpc:"priceRestrictionCount"`

	// The item price that an account is restricted to.
	PriceRestrictions []Product_Item_Price_Account_Restriction `json:"priceRestrictions,omitempty" xmlrpc:"priceRestrictions"`

	// A count of all priority one tickets associated with an account.
	PriorityOneTicketCount *uint `json:"priorityOneTicketCount,omitempty" xmlrpc:"priorityOneTicketCount"`

	// All priority one tickets associated with an account.
	PriorityOneTickets []Ticket `json:"priorityOneTickets,omitempty" xmlrpc:"priorityOneTickets"`

	// A count of dEPRECATED - This information can be pulled directly through tapping keys now - DEPRECATED. The allotments for this account and their servers. The private inbound and outbound bandwidth is calculated for each server in addition to the daily average network traffic since the last billing date.
	PrivateAllotmentHardwareBandwidthDetailCount *uint `json:"privateAllotmentHardwareBandwidthDetailCount,omitempty" xmlrpc:"privateAllotmentHardwareBandwidthDetailCount"`

	// DEPRECATED - This information can be pulled directly through tapping keys now - DEPRECATED. The allotments for this account and their servers. The private inbound and outbound bandwidth is calculated for each server in addition to the daily average network traffic since the last billing date.
	PrivateAllotmentHardwareBandwidthDetails []Network_Bandwidth_Version1_Allotment `json:"privateAllotmentHardwareBandwidthDetails,omitempty" xmlrpc:"privateAllotmentHardwareBandwidthDetails"`

	// A count of private and shared template group objects (parent only) for an account.
	PrivateBlockDeviceTemplateGroupCount *uint `json:"privateBlockDeviceTemplateGroupCount,omitempty" xmlrpc:"privateBlockDeviceTemplateGroupCount"`

	// Private and shared template group objects (parent only) for an account.
	PrivateBlockDeviceTemplateGroups []Virtual_Guest_Block_Device_Template_Group `json:"privateBlockDeviceTemplateGroups,omitempty" xmlrpc:"privateBlockDeviceTemplateGroups"`

	// A count of
	PrivateIpAddressCount *uint `json:"privateIpAddressCount,omitempty" xmlrpc:"privateIpAddressCount"`

	// no documentation yet
	PrivateIpAddresses []Network_Subnet_IpAddress `json:"privateIpAddresses,omitempty" xmlrpc:"privateIpAddresses"`

	// A count of the private network VLANs assigned to an account.
	PrivateNetworkVlanCount *uint `json:"privateNetworkVlanCount,omitempty" xmlrpc:"privateNetworkVlanCount"`

	// The private network VLANs assigned to an account.
	PrivateNetworkVlans []Network_Vlan `json:"privateNetworkVlans,omitempty" xmlrpc:"privateNetworkVlans"`

	// A count of all private subnets associated with an account.
	PrivateSubnetCount *uint `json:"privateSubnetCount,omitempty" xmlrpc:"privateSubnetCount"`

	// All private subnets associated with an account.
	PrivateSubnets []Network_Subnet `json:"privateSubnets,omitempty" xmlrpc:"privateSubnets"`

	// A count of dEPRECATED - This information can be pulled directly through tapping keys now - DEPRECATED. The allotments for this account and their servers. The public inbound and outbound bandwidth is calculated for each server in addition to the daily average network traffic since the last billing date.
	PublicAllotmentHardwareBandwidthDetailCount *uint `json:"publicAllotmentHardwareBandwidthDetailCount,omitempty" xmlrpc:"publicAllotmentHardwareBandwidthDetailCount"`

	// DEPRECATED - This information can be pulled directly through tapping keys now - DEPRECATED. The allotments for this account and their servers. The public inbound and outbound bandwidth is calculated for each server in addition to the daily average network traffic since the last billing date.
	PublicAllotmentHardwareBandwidthDetails []Network_Bandwidth_Version1_Allotment `json:"publicAllotmentHardwareBandwidthDetails,omitempty" xmlrpc:"publicAllotmentHardwareBandwidthDetails"`

	// A count of
	PublicIpAddressCount *uint `json:"publicIpAddressCount,omitempty" xmlrpc:"publicIpAddressCount"`

	// no documentation yet
	PublicIpAddresses []Network_Subnet_IpAddress `json:"publicIpAddresses,omitempty" xmlrpc:"publicIpAddresses"`

	// A count of the public network VLANs assigned to an account.
	PublicNetworkVlanCount *uint `json:"publicNetworkVlanCount,omitempty" xmlrpc:"publicNetworkVlanCount"`

	// The public network VLANs assigned to an account.
	PublicNetworkVlans []Network_Vlan `json:"publicNetworkVlans,omitempty" xmlrpc:"publicNetworkVlans"`

	// A count of all public network subnets associated with an account.
	PublicSubnetCount *uint `json:"publicSubnetCount,omitempty" xmlrpc:"publicSubnetCount"`

	// All public network subnets associated with an account.
	PublicSubnets []Network_Subnet `json:"publicSubnets,omitempty" xmlrpc:"publicSubnets"`

	// A count of an account's quotes.
	QuoteCount *uint `json:"quoteCount,omitempty" xmlrpc:"quoteCount"`

	// An account's quotes.
	Quotes []Billing_Order_Quote `json:"quotes,omitempty" xmlrpc:"quotes"`

	// A count of
	RecentEventCount *uint `json:"recentEventCount,omitempty" xmlrpc:"recentEventCount"`

	// no documentation yet
	RecentEvents []Notification_Occurrence_Event `json:"recentEvents,omitempty" xmlrpc:"recentEvents"`

	// The Referral Partner for this account, if any.
	ReferralPartner *Account `json:"referralPartner,omitempty" xmlrpc:"referralPartner"`

	// A count of if this is a account is a referral partner, the accounts this referral partner has referred
	ReferredAccountCount *uint `json:"referredAccountCount,omitempty" xmlrpc:"referredAccountCount"`

	// If this is a account is a referral partner, the accounts this referral partner has referred
	ReferredAccounts []Account `json:"referredAccounts,omitempty" xmlrpc:"referredAccounts"`

	// A count of
	RegulatedWorkloadCount *uint `json:"regulatedWorkloadCount,omitempty" xmlrpc:"regulatedWorkloadCount"`

	// no documentation yet
	RegulatedWorkloads []Legal_RegulatedWorkload `json:"regulatedWorkloads,omitempty" xmlrpc:"regulatedWorkloads"`

	// A count of remote management command requests for an account
	RemoteManagementCommandRequestCount *uint `json:"remoteManagementCommandRequestCount,omitempty" xmlrpc:"remoteManagementCommandRequestCount"`

	// Remote management command requests for an account
	RemoteManagementCommandRequests []Hardware_Component_RemoteManagement_Command_Request `json:"remoteManagementCommandRequests,omitempty" xmlrpc:"remoteManagementCommandRequests"`

	// A count of the Replication events for all Network Storage volumes on an account.
	ReplicationEventCount *uint `json:"replicationEventCount,omitempty" xmlrpc:"replicationEventCount"`

	// The Replication events for all Network Storage volumes on an account.
	ReplicationEvents []Network_Storage_Event `json:"replicationEvents,omitempty" xmlrpc:"replicationEvents"`

	// A count of an account's associated top-level resource groups.
	ResourceGroupCount *uint `json:"resourceGroupCount,omitempty" xmlrpc:"resourceGroupCount"`

	// An account's associated top-level resource groups.
	ResourceGroups []Resource_Group `json:"resourceGroups,omitempty" xmlrpc:"resourceGroups"`

	// A count of all Routers that an accounts VLANs reside on
	RouterCount *uint `json:"routerCount,omitempty" xmlrpc:"routerCount"`

	// All Routers that an accounts VLANs reside on
	Routers []Hardware `json:"routers,omitempty" xmlrpc:"routers"`

	// An account's reverse WHOIS data. This data is used when making SWIP requests.
	RwhoisData *Network_Subnet_Rwhois_Data `json:"rwhoisData,omitempty" xmlrpc:"rwhoisData"`

	// no documentation yet
	SalesforceAccountLink *Account_Link `json:"salesforceAccountLink,omitempty" xmlrpc:"salesforceAccountLink"`

	// The SAML configuration for this account.
	SamlAuthentication *Account_Authentication_Saml `json:"samlAuthentication,omitempty" xmlrpc:"samlAuthentication"`

	// A count of all scale groups on this account.
	ScaleGroupCount *uint `json:"scaleGroupCount,omitempty" xmlrpc:"scaleGroupCount"`

	// All scale groups on this account.
	ScaleGroups []Scale_Group `json:"scaleGroups,omitempty" xmlrpc:"scaleGroups"`

	// A count of the secondary DNS records for a SoftLayer customer account.
	SecondaryDomainCount *uint `json:"secondaryDomainCount,omitempty" xmlrpc:"secondaryDomainCount"`

	// The secondary DNS records for a SoftLayer customer account.
	SecondaryDomains []Dns_Secondary `json:"secondaryDomains,omitempty" xmlrpc:"secondaryDomains"`

	// A count of stored security certificates (ie. SSL)
	SecurityCertificateCount *uint `json:"securityCertificateCount,omitempty" xmlrpc:"securityCertificateCount"`

	// Stored security certificates (ie. SSL)
	SecurityCertificates []Security_Certificate `json:"securityCertificates,omitempty" xmlrpc:"securityCertificates"`

	// A count of an account's vulnerability scan requests.
	SecurityScanRequestCount *uint `json:"securityScanRequestCount,omitempty" xmlrpc:"securityScanRequestCount"`

	// An account's vulnerability scan requests.
	SecurityScanRequests []Network_Security_Scanner_Request `json:"securityScanRequests,omitempty" xmlrpc:"securityScanRequests"`

	// A count of the service billing items that will be on an account's next invoice.
	ServiceBillingItemCount *uint `json:"serviceBillingItemCount,omitempty" xmlrpc:"serviceBillingItemCount"`

	// The service billing items that will be on an account's next invoice.
	ServiceBillingItems []Billing_Item `json:"serviceBillingItems,omitempty" xmlrpc:"serviceBillingItems"`

	// A count of shipments that belong to the customer's account.
	ShipmentCount *uint `json:"shipmentCount,omitempty" xmlrpc:"shipmentCount"`

	// Shipments that belong to the customer's account.
	Shipments []Account_Shipment `json:"shipments,omitempty" xmlrpc:"shipments"`

	// A count of customer specified SSH keys that can be implemented onto a newly provisioned or reloaded server.
	SshKeyCount *uint `json:"sshKeyCount,omitempty" xmlrpc:"sshKeyCount"`

	// Customer specified SSH keys that can be implemented onto a newly provisioned or reloaded server.
	SshKeys []Security_Ssh_Key `json:"sshKeys,omitempty" xmlrpc:"sshKeys"`

	// A count of an account's associated portal users with SSL VPN access.
	SslVpnUserCount *uint `json:"sslVpnUserCount,omitempty" xmlrpc:"sslVpnUserCount"`

	// An account's associated portal users with SSL VPN access.
	SslVpnUsers []User_Customer `json:"sslVpnUsers,omitempty" xmlrpc:"sslVpnUsers"`

	// A count of an account's virtual guest objects that are hosted on a user provisioned hypervisor.
	StandardPoolVirtualGuestCount *uint `json:"standardPoolVirtualGuestCount,omitempty" xmlrpc:"standardPoolVirtualGuestCount"`

	// An account's virtual guest objects that are hosted on a user provisioned hypervisor.
	StandardPoolVirtualGuests []Virtual_Guest `json:"standardPoolVirtualGuests,omitempty" xmlrpc:"standardPoolVirtualGuests"`

	// A two-letter abbreviation of the state in the mailing address belonging to an account. If an account does not reside in a province then this is typically blank.
	State *string `json:"state,omitempty" xmlrpc:"state"`

	// The date of an account's last status change.
	StatusDate *Time `json:"statusDate,omitempty" xmlrpc:"statusDate"`

	// A count of all network subnets associated with an account.
	SubnetCount *uint `json:"subnetCount,omitempty" xmlrpc:"subnetCount"`

	// A count of
	SubnetRegistrationCount *uint `json:"subnetRegistrationCount,omitempty" xmlrpc:"subnetRegistrationCount"`

	// A count of
	SubnetRegistrationDetailCount *uint `json:"subnetRegistrationDetailCount,omitempty" xmlrpc:"subnetRegistrationDetailCount"`

	// no documentation yet
	SubnetRegistrationDetails []Account_Regional_Registry_Detail `json:"subnetRegistrationDetails,omitempty" xmlrpc:"subnetRegistrationDetails"`

	// no documentation yet
	SubnetRegistrations []Network_Subnet_Registration `json:"subnetRegistrations,omitempty" xmlrpc:"subnetRegistrations"`

	// All network subnets associated with an account.
	Subnets []Network_Subnet `json:"subnets,omitempty" xmlrpc:"subnets"`

	// A count of the SoftLayer employees that an account is assigned to.
	SupportRepresentativeCount *uint `json:"supportRepresentativeCount,omitempty" xmlrpc:"supportRepresentativeCount"`

	// The SoftLayer employees that an account is assigned to.
	SupportRepresentatives []User_Employee `json:"supportRepresentatives,omitempty" xmlrpc:"supportRepresentatives"`

	// A count of the active support subscriptions for this account.
	SupportSubscriptionCount *uint `json:"supportSubscriptionCount,omitempty" xmlrpc:"supportSubscriptionCount"`

	// The active support subscriptions for this account.
	SupportSubscriptions []Billing_Item `json:"supportSubscriptions,omitempty" xmlrpc:"supportSubscriptions"`

	// no documentation yet
	SupportTier *string `json:"supportTier,omitempty" xmlrpc:"supportTier"`

	// A flag indicating to suppress invoices.
	SuppressInvoicesFlag *bool `json:"suppressInvoicesFlag,omitempty" xmlrpc:"suppressInvoicesFlag"`

	// A count of
	TagCount *uint `json:"tagCount,omitempty" xmlrpc:"tagCount"`

	// no documentation yet
	Tags []Tag `json:"tags,omitempty" xmlrpc:"tags"`

	// A count of an account's associated tickets.
	TicketCount *uint `json:"ticketCount,omitempty" xmlrpc:"ticketCount"`

	// An account's associated tickets.
	Tickets []Ticket `json:"tickets,omitempty" xmlrpc:"tickets"`

	// Tickets closed within the last 72 hours or last 10 tickets, whichever is less, associated with an account.
	TicketsClosedInTheLastThreeDays []Ticket `json:"ticketsClosedInTheLastThreeDays,omitempty" xmlrpc:"ticketsClosedInTheLastThreeDays"`

	// A count of tickets closed within the last 72 hours or last 10 tickets, whichever is less, associated with an account.
	TicketsClosedInTheLastThreeDaysCount *uint `json:"ticketsClosedInTheLastThreeDaysCount,omitempty" xmlrpc:"ticketsClosedInTheLastThreeDaysCount"`

	// Tickets closed today associated with an account.
	TicketsClosedToday []Ticket `json:"ticketsClosedToday,omitempty" xmlrpc:"ticketsClosedToday"`

	// A count of tickets closed today associated with an account.
	TicketsClosedTodayCount *uint `json:"ticketsClosedTodayCount,omitempty" xmlrpc:"ticketsClosedTodayCount"`

	// A count of an account's associated Transcode account.
	TranscodeAccountCount *uint `json:"transcodeAccountCount,omitempty" xmlrpc:"transcodeAccountCount"`

	// An account's associated Transcode account.
	TranscodeAccounts []Network_Media_Transcode_Account `json:"transcodeAccounts,omitempty" xmlrpc:"transcodeAccounts"`

	// A count of an account's associated upgrade requests.
	UpgradeRequestCount *uint `json:"upgradeRequestCount,omitempty" xmlrpc:"upgradeRequestCount"`

	// An account's associated upgrade requests.
	UpgradeRequests []Product_Upgrade_Request `json:"upgradeRequests,omitempty" xmlrpc:"upgradeRequests"`

	// A count of an account's portal users.
	UserCount *uint `json:"userCount,omitempty" xmlrpc:"userCount"`

	// An account's portal users.
	Users []User_Customer `json:"users,omitempty" xmlrpc:"users"`

	// A count of stored security certificates that are not expired (ie. SSL)
	ValidSecurityCertificateCount *uint `json:"validSecurityCertificateCount,omitempty" xmlrpc:"validSecurityCertificateCount"`

	// Stored security certificates that are not expired (ie. SSL)
	ValidSecurityCertificates []Security_Certificate `json:"validSecurityCertificates,omitempty" xmlrpc:"validSecurityCertificates"`

	// Return 0 if vpn updates are currently in progress on this account otherwise 1.
	VdrUpdatesInProgressFlag *bool `json:"vdrUpdatesInProgressFlag,omitempty" xmlrpc:"vdrUpdatesInProgressFlag"`

	// A count of the bandwidth pooling for this account.
	VirtualDedicatedRackCount *uint `json:"virtualDedicatedRackCount,omitempty" xmlrpc:"virtualDedicatedRackCount"`

	// The bandwidth pooling for this account.
	VirtualDedicatedRacks []Network_Bandwidth_Version1_Allotment `json:"virtualDedicatedRacks,omitempty" xmlrpc:"virtualDedicatedRacks"`

	// A count of an account's associated virtual server virtual disk images.
	VirtualDiskImageCount *uint `json:"virtualDiskImageCount,omitempty" xmlrpc:"virtualDiskImageCount"`

	// An account's associated virtual server virtual disk images.
	VirtualDiskImages []Virtual_Disk_Image `json:"virtualDiskImages,omitempty" xmlrpc:"virtualDiskImages"`

	// A count of an account's associated virtual guest objects.
	VirtualGuestCount *uint `json:"virtualGuestCount,omitempty" xmlrpc:"virtualGuestCount"`

	// An account's associated virtual guest objects.
	VirtualGuests []Virtual_Guest `json:"virtualGuests,omitempty" xmlrpc:"virtualGuests"`

	// An account's associated virtual guest objects currently over bandwidth allocation.
	VirtualGuestsOverBandwidthAllocation []Virtual_Guest `json:"virtualGuestsOverBandwidthAllocation,omitempty" xmlrpc:"virtualGuestsOverBandwidthAllocation"`

	// A count of an account's associated virtual guest objects currently over bandwidth allocation.
	VirtualGuestsOverBandwidthAllocationCount *uint `json:"virtualGuestsOverBandwidthAllocationCount,omitempty" xmlrpc:"virtualGuestsOverBandwidthAllocationCount"`

	// An account's associated virtual guest objects currently over bandwidth allocation.
	VirtualGuestsProjectedOverBandwidthAllocation []Virtual_Guest `json:"virtualGuestsProjectedOverBandwidthAllocation,omitempty" xmlrpc:"virtualGuestsProjectedOverBandwidthAllocation"`

	// A count of an account's associated virtual guest objects currently over bandwidth allocation.
	VirtualGuestsProjectedOverBandwidthAllocationCount *uint `json:"virtualGuestsProjectedOverBandwidthAllocationCount,omitempty" xmlrpc:"virtualGuestsProjectedOverBandwidthAllocationCount"`

	// All virtual guests associated with an account that has the cPanel web hosting control panel installed.
	VirtualGuestsWithCpanel []Virtual_Guest `json:"virtualGuestsWithCpanel,omitempty" xmlrpc:"virtualGuestsWithCpanel"`

	// A count of all virtual guests associated with an account that has the cPanel web hosting control panel installed.
	VirtualGuestsWithCpanelCount *uint `json:"virtualGuestsWithCpanelCount,omitempty" xmlrpc:"virtualGuestsWithCpanelCount"`

	// All virtual guests associated with an account that have McAfee Secure software components.
	VirtualGuestsWithMcafee []Virtual_Guest `json:"virtualGuestsWithMcafee,omitempty" xmlrpc:"virtualGuestsWithMcafee"`

	// All virtual guests associated with an account that have McAfee Secure AntiVirus for Redhat software components.
	VirtualGuestsWithMcafeeAntivirusRedhat []Virtual_Guest `json:"virtualGuestsWithMcafeeAntivirusRedhat,omitempty" xmlrpc:"virtualGuestsWithMcafeeAntivirusRedhat"`

	// A count of all virtual guests associated with an account that have McAfee Secure AntiVirus for Redhat software components.
	VirtualGuestsWithMcafeeAntivirusRedhatCount *uint `json:"virtualGuestsWithMcafeeAntivirusRedhatCount,omitempty" xmlrpc:"virtualGuestsWithMcafeeAntivirusRedhatCount"`

	// A count of all virtual guests associated with an account that has McAfee Secure AntiVirus for Windows software components.
	VirtualGuestsWithMcafeeAntivirusWindowCount *uint `json:"virtualGuestsWithMcafeeAntivirusWindowCount,omitempty" xmlrpc:"virtualGuestsWithMcafeeAntivirusWindowCount"`

	// All virtual guests associated with an account that has McAfee Secure AntiVirus for Windows software components.
	VirtualGuestsWithMcafeeAntivirusWindows []Virtual_Guest `json:"virtualGuestsWithMcafeeAntivirusWindows,omitempty" xmlrpc:"virtualGuestsWithMcafeeAntivirusWindows"`

	// A count of all virtual guests associated with an account that have McAfee Secure software components.
	VirtualGuestsWithMcafeeCount *uint `json:"virtualGuestsWithMcafeeCount,omitempty" xmlrpc:"virtualGuestsWithMcafeeCount"`

	// All virtual guests associated with an account that has McAfee Secure Intrusion Detection System software components.
	VirtualGuestsWithMcafeeIntrusionDetectionSystem []Virtual_Guest `json:"virtualGuestsWithMcafeeIntrusionDetectionSystem,omitempty" xmlrpc:"virtualGuestsWithMcafeeIntrusionDetectionSystem"`

	// A count of all virtual guests associated with an account that has McAfee Secure Intrusion Detection System software components.
	VirtualGuestsWithMcafeeIntrusionDetectionSystemCount *uint `json:"virtualGuestsWithMcafeeIntrusionDetectionSystemCount,omitempty" xmlrpc:"virtualGuestsWithMcafeeIntrusionDetectionSystemCount"`

	// All virtual guests associated with an account that has the Plesk web hosting control panel installed.
	VirtualGuestsWithPlesk []Virtual_Guest `json:"virtualGuestsWithPlesk,omitempty" xmlrpc:"virtualGuestsWithPlesk"`

	// A count of all virtual guests associated with an account that has the Plesk web hosting control panel installed.
	VirtualGuestsWithPleskCount *uint `json:"virtualGuestsWithPleskCount,omitempty" xmlrpc:"virtualGuestsWithPleskCount"`

	// All virtual guests associated with an account that have the QuantaStor storage system installed.
	VirtualGuestsWithQuantastor []Virtual_Guest `json:"virtualGuestsWithQuantastor,omitempty" xmlrpc:"virtualGuestsWithQuantastor"`

	// A count of all virtual guests associated with an account that have the QuantaStor storage system installed.
	VirtualGuestsWithQuantastorCount *uint `json:"virtualGuestsWithQuantastorCount,omitempty" xmlrpc:"virtualGuestsWithQuantastorCount"`

	// All virtual guests associated with an account that has the Urchin web traffic analytics package installed.
	VirtualGuestsWithUrchin []Virtual_Guest `json:"virtualGuestsWithUrchin,omitempty" xmlrpc:"virtualGuestsWithUrchin"`

	// A count of all virtual guests associated with an account that has the Urchin web traffic analytics package installed.
	VirtualGuestsWithUrchinCount *uint `json:"virtualGuestsWithUrchinCount,omitempty" xmlrpc:"virtualGuestsWithUrchinCount"`

	// The bandwidth pooling for this account.
	VirtualPrivateRack *Network_Bandwidth_Version1_Allotment `json:"virtualPrivateRack,omitempty" xmlrpc:"virtualPrivateRack"`

	// An account's associated virtual server archived storage repositories.
	VirtualStorageArchiveRepositories []Virtual_Storage_Repository `json:"virtualStorageArchiveRepositories,omitempty" xmlrpc:"virtualStorageArchiveRepositories"`

	// A count of an account's associated virtual server archived storage repositories.
	VirtualStorageArchiveRepositoryCount *uint `json:"virtualStorageArchiveRepositoryCount,omitempty" xmlrpc:"virtualStorageArchiveRepositoryCount"`

	// An account's associated virtual server public storage repositories.
	VirtualStoragePublicRepositories []Virtual_Storage_Repository `json:"virtualStoragePublicRepositories,omitempty" xmlrpc:"virtualStoragePublicRepositories"`

	// A count of an account's associated virtual server public storage repositories.
	VirtualStoragePublicRepositoryCount *uint `json:"virtualStoragePublicRepositoryCount,omitempty" xmlrpc:"virtualStoragePublicRepositoryCount"`
}

// An unfortunate facet of the hosting business is the necessity of with legal and network abuse inquiries. As these types of inquiries frequently contain sensitive information SoftLayer keeps a separate account contact email address for direct contact about legal and abuse matters, modeled by the SoftLayer_Account_AbuseEmail data type. SoftLayer will typically email an account's abuse email addresses in these types of cases, and an email is automatically sent to an account's abuse email addresses when a legal or abuse ticket is created or updated.
type Account_AbuseEmail struct {
	Entity

	// The account associated with an abuse email address.
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// A valid email address.
	Email *string `json:"email,omitempty" xmlrpc:"email"`
}

// The SoftLayer_Account_Address data type contains information on an address associated with a SoftLayer account.
type Account_Address struct {
	Entity

	// The account to which this address belongs.
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// no documentation yet
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// Line 1 of the address (normally the street address).
	Address1 *string `json:"address1,omitempty" xmlrpc:"address1"`

	// Line 2 of the address.
	Address2 *string `json:"address2,omitempty" xmlrpc:"address2"`

	// The city of the address.
	City *string `json:"city,omitempty" xmlrpc:"city"`

	// The contact name (person, office) of the address.
	ContactName *string `json:"contactName,omitempty" xmlrpc:"contactName"`

	// The country of the address.
	Country *string `json:"country,omitempty" xmlrpc:"country"`

	// The customer user who created this address.
	CreateUser *User_Customer `json:"createUser,omitempty" xmlrpc:"createUser"`

	// The description of the address.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// The unique id of the address.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// Flag to show whether the address is active.
	IsActive *int `json:"isActive,omitempty" xmlrpc:"isActive"`

	// The location of this address.
	Location *Location `json:"location,omitempty" xmlrpc:"location"`

	// The location id of the address.
	LocationId *int `json:"locationId,omitempty" xmlrpc:"locationId"`

	// The employee who last modified this address.
	ModifyEmployee *User_Employee `json:"modifyEmployee,omitempty" xmlrpc:"modifyEmployee"`

	// The customer user who last modified this address.
	ModifyUser *User_Customer `json:"modifyUser,omitempty" xmlrpc:"modifyUser"`

	// The postal (zip) code of the address.
	PostalCode *string `json:"postalCode,omitempty" xmlrpc:"postalCode"`

	// The state of the address.
	State *string `json:"state,omitempty" xmlrpc:"state"`

	// An account address' type.
	Type *Account_Address_Type `json:"type,omitempty" xmlrpc:"type"`
}

// no documentation yet
type Account_Address_Type struct {
	Entity

	// DEPRECATED
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// no documentation yet
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// This service allows for a unique identifier to be associated to an existing customer account.
type Account_Affiliation struct {
	Entity

	// The account that an affiliation belongs to.
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// A customer account's internal identifier.
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// An affiliate identifier associated with the customer account.
	AffiliateId *string `json:"affiliateId,omitempty" xmlrpc:"affiliateId"`

	// The date an account affiliation was created.
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// A customer affiliation internal identifier.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// The date an account affiliation was last modified.
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`
}

// no documentation yet
type Account_Agreement struct {
	Entity

	// no documentation yet
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// The type of agreement.
	AgreementType *Account_Agreement_Type `json:"agreementType,omitempty" xmlrpc:"agreementType"`

	// The type of agreement identifier.
	AgreementTypeId *int `json:"agreementTypeId,omitempty" xmlrpc:"agreementTypeId"`

	// A count of the files attached to an agreement.
	AttachedBillingAgreementFileCount *uint `json:"attachedBillingAgreementFileCount,omitempty" xmlrpc:"attachedBillingAgreementFileCount"`

	// The files attached to an agreement.
	AttachedBillingAgreementFiles []Account_MasterServiceAgreement `json:"attachedBillingAgreementFiles,omitempty" xmlrpc:"attachedBillingAgreementFiles"`

	// no documentation yet
	AutoRenew *int `json:"autoRenew,omitempty" xmlrpc:"autoRenew"`

	// A count of the billing items associated with an agreement.
	BillingItemCount *uint `json:"billingItemCount,omitempty" xmlrpc:"billingItemCount"`

	// The billing items associated with an agreement.
	BillingItems []Billing_Item `json:"billingItems,omitempty" xmlrpc:"billingItems"`

	// no documentation yet
	CancellationFee *int `json:"cancellationFee,omitempty" xmlrpc:"cancellationFee"`

	// The date an agreement was created.
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// The duration in months of an agreement.
	DurationMonths *int `json:"durationMonths,omitempty" xmlrpc:"durationMonths"`

	// The end date of an agreement.
	EndDate *Time `json:"endDate,omitempty" xmlrpc:"endDate"`

	// An agreement's internal identifier.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// The effective start date of an agreement.
	StartDate *Time `json:"startDate,omitempty" xmlrpc:"startDate"`

	// The status of the agreement.
	Status *Account_Agreement_Status `json:"status,omitempty" xmlrpc:"status"`

	// The status identifier for an agreement.
	StatusId *int `json:"statusId,omitempty" xmlrpc:"statusId"`

	// The title of an agreement.
	Title *string `json:"title,omitempty" xmlrpc:"title"`

	// A count of the top level billing item associated with an agreement.
	TopLevelBillingItemCount *uint `json:"topLevelBillingItemCount,omitempty" xmlrpc:"topLevelBillingItemCount"`

	// The top level billing item associated with an agreement.
	TopLevelBillingItems []Billing_Item `json:"topLevelBillingItems,omitempty" xmlrpc:"topLevelBillingItems"`
}

// no documentation yet
type Account_Agreement_Status struct {
	Entity

	// The name of the agreement status.
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// no documentation yet
type Account_Agreement_Type struct {
	Entity

	// The name of the agreement type.
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// A SoftLayer_Account_Attachment_Employee models an assignment of a single [[SoftLayer_User_Employee|employee]] with a single [[SoftLayer_Account|account]]
type Account_Attachment_Employee struct {
	Entity

	// A [[SoftLayer_Account|account]] that is assigned to a [[SoftLayer_User_Employee|employee]].
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// A [[SoftLayer_User_Employee|employee]] that is assigned to a [[SoftLayer_Account|account]].
	Employee *User_Employee `json:"employee,omitempty" xmlrpc:"employee"`

	// A [[SoftLayer_User_Employee|employee]] that is assigned to a [[SoftLayer_Account|account]].
	EmployeeRole *Account_Attachment_Employee_Role `json:"employeeRole,omitempty" xmlrpc:"employeeRole"`

	// Role identifier.
	RoleId *int `json:"roleId,omitempty" xmlrpc:"roleId"`
}

// no documentation yet
type Account_Attachment_Employee_Role struct {
	Entity

	// no documentation yet
	Keyname *string `json:"keyname,omitempty" xmlrpc:"keyname"`

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// Many SoftLayer customer accounts have individual attributes assigned to them that describe features or special features for that account, such as special pricing, account statuses, and ordering instructions. The SoftLayer_Account_Attribute data type contains information relating to a single SoftLayer_Account attribute.
type Account_Attribute struct {
	Entity

	// The SoftLayer customer account that has an attribute.
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// The type of attribute assigned to a SoftLayer customer account.
	AccountAttributeType *Account_Attribute_Type `json:"accountAttributeType,omitempty" xmlrpc:"accountAttributeType"`

	// The internal identifier of the type of attribute that a SoftLayer customer account attribute belongs to.
	AccountAttributeTypeId *int `json:"accountAttributeTypeId,omitempty" xmlrpc:"accountAttributeTypeId"`

	// The internal identifier of the SoftLayer customer account that is assigned an account attribute.
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// A SoftLayer customer account attribute's internal identifier.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// A SoftLayer account attribute's value.
	Value *string `json:"value,omitempty" xmlrpc:"value"`
}

// SoftLayer_Account_Attribute_Type models the type of attribute that can be assigned to a SoftLayer customer account.
type Account_Attribute_Type struct {
	Entity

	// A brief description of a SoftLayer account attribute type.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// A SoftLayer account attribute type's internal identifier.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// A SoftLayer account attribute type's key name. This is typically a shorter version of an attribute type's name.
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// A SoftLayer account attribute type's name.
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// Account authentication has many different settings that can be set. This class allows the customer or employee to set these settigns.
type Account_Authentication_Attribute struct {
	Entity

	// The SoftLayer customer account.
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// The internal identifier of the SoftLayer customer account that is assigned an account authenction attribute.
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// The SoftLayer account authentication that has an attribute.
	AuthenticationRecord *Account_Authentication_Saml `json:"authenticationRecord,omitempty" xmlrpc:"authenticationRecord"`

	// A SoftLayer account authenction attribute's internal identifier.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// The type of attribute assigned to a SoftLayer account authentication.
	Type *Account_Authentication_Attribute_Type `json:"type,omitempty" xmlrpc:"type"`

	// The internal identifier of the type of attribute that a SoftLayer account authenction attribute belongs to.
	TypeId *int `json:"typeId,omitempty" xmlrpc:"typeId"`

	// A SoftLayer account authenction attribute's value.
	Value *string `json:"value,omitempty" xmlrpc:"value"`
}

// SoftLayer_Account_Authentication_Attribute_Type models the type of attribute that can be assigned to a SoftLayer customer account authentication.
type Account_Authentication_Attribute_Type struct {
	Entity

	// A brief description of a SoftLayer account authentication attribute type.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// A SoftLayer account authentication attribute type's internal identifier.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// A SoftLayer account authentication attribute type's key name. This is typically a shorter version of an attribute type's name.
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// A SoftLayer account authentication attribute type's name.
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// An example of what you can put in as your value.
	ValueExample *string `json:"valueExample,omitempty" xmlrpc:"valueExample"`
}

// no documentation yet
type Account_Authentication_OpenIdConnect_Option struct {
	Entity

	// no documentation yet
	Key *string `json:"key,omitempty" xmlrpc:"key"`

	// no documentation yet
	Value *string `json:"value,omitempty" xmlrpc:"value"`
}

// no documentation yet
type Account_Authentication_OpenIdConnect_RegistrationInformation struct {
	Entity

	// no documentation yet
	ExistingBlueIdFlag *bool `json:"existingBlueIdFlag,omitempty" xmlrpc:"existingBlueIdFlag"`

	// no documentation yet
	FederatedEmailDomainFlag *bool `json:"federatedEmailDomainFlag,omitempty" xmlrpc:"federatedEmailDomainFlag"`

	// no documentation yet
	User *User_Customer `json:"user,omitempty" xmlrpc:"user"`
}

// no documentation yet
type Account_Authentication_Saml struct {
	Entity

	// The account associated with this saml configuration.
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// The saml account id.
	AccountId *string `json:"accountId,omitempty" xmlrpc:"accountId"`

	// A count of the saml attribute values for a SoftLayer customer account.
	AttributeCount *uint `json:"attributeCount,omitempty" xmlrpc:"attributeCount"`

	// The saml attribute values for a SoftLayer customer account.
	Attributes []Account_Authentication_Attribute `json:"attributes,omitempty" xmlrpc:"attributes"`

	// The identity provider x509 certificate.
	Certificate *string `json:"certificate,omitempty" xmlrpc:"certificate"`

	// The identity provider x509 certificate fingerprint.
	CertificateFingerprint *string `json:"certificateFingerprint,omitempty" xmlrpc:"certificateFingerprint"`

	// The identity provider entity ID.
	EntityId *string `json:"entityId,omitempty" xmlrpc:"entityId"`

	// The saml internal identifying number.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// The service provider x509 certificate.
	ServiceProviderCertificate *string `json:"serviceProviderCertificate,omitempty" xmlrpc:"serviceProviderCertificate"`

	// The service provider entity IDs.
	ServiceProviderEntityId *string `json:"serviceProviderEntityId,omitempty" xmlrpc:"serviceProviderEntityId"`

	// The service provider public key.
	ServiceProviderPublicKey *string `json:"serviceProviderPublicKey,omitempty" xmlrpc:"serviceProviderPublicKey"`

	// The service provider signle logout encoding.
	ServiceProviderSingleLogoutEncoding *string `json:"serviceProviderSingleLogoutEncoding,omitempty" xmlrpc:"serviceProviderSingleLogoutEncoding"`

	// The service provider signle logout address.
	ServiceProviderSingleLogoutUrl *string `json:"serviceProviderSingleLogoutUrl,omitempty" xmlrpc:"serviceProviderSingleLogoutUrl"`

	// The service provider signle sign on encoding.
	ServiceProviderSingleSignOnEncoding *string `json:"serviceProviderSingleSignOnEncoding,omitempty" xmlrpc:"serviceProviderSingleSignOnEncoding"`

	// The service provider signle sign on address.
	ServiceProviderSingleSignOnUrl *string `json:"serviceProviderSingleSignOnUrl,omitempty" xmlrpc:"serviceProviderSingleSignOnUrl"`

	// The identity provider single logout encoding.
	SingleLogoutEncoding *string `json:"singleLogoutEncoding,omitempty" xmlrpc:"singleLogoutEncoding"`

	// The identity provider sigle logout address.
	SingleLogoutUrl *string `json:"singleLogoutUrl,omitempty" xmlrpc:"singleLogoutUrl"`

	// The identity provider single sign on encoding.
	SingleSignOnEncoding *string `json:"singleSignOnEncoding,omitempty" xmlrpc:"singleSignOnEncoding"`

	// The identity provider signle sign on address.
	SingleSignOnUrl *string `json:"singleSignOnUrl,omitempty" xmlrpc:"singleSignOnUrl"`
}

// no documentation yet
type Account_Classification_Group_Type struct {
	Entity

	// no documentation yet
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`
}

// no documentation yet
type Account_Contact struct {
	Entity

	// no documentation yet
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// no documentation yet
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// no documentation yet
	Address1 *string `json:"address1,omitempty" xmlrpc:"address1"`

	// no documentation yet
	Address2 *string `json:"address2,omitempty" xmlrpc:"address2"`

	// no documentation yet
	AlternatePhone *string `json:"alternatePhone,omitempty" xmlrpc:"alternatePhone"`

	// no documentation yet
	City *string `json:"city,omitempty" xmlrpc:"city"`

	// no documentation yet
	CompanyName *string `json:"companyName,omitempty" xmlrpc:"companyName"`

	// no documentation yet
	Country *string `json:"country,omitempty" xmlrpc:"country"`

	// no documentation yet
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// no documentation yet
	Email *string `json:"email,omitempty" xmlrpc:"email"`

	// no documentation yet
	FaxPhone *string `json:"faxPhone,omitempty" xmlrpc:"faxPhone"`

	// no documentation yet
	FirstName *string `json:"firstName,omitempty" xmlrpc:"firstName"`

	// no documentation yet
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	JobTitle *string `json:"jobTitle,omitempty" xmlrpc:"jobTitle"`

	// no documentation yet
	LastName *string `json:"lastName,omitempty" xmlrpc:"lastName"`

	// no documentation yet
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// no documentation yet
	OfficePhone *string `json:"officePhone,omitempty" xmlrpc:"officePhone"`

	// no documentation yet
	PostalCode *string `json:"postalCode,omitempty" xmlrpc:"postalCode"`

	// no documentation yet
	ProfileName *string `json:"profileName,omitempty" xmlrpc:"profileName"`

	// no documentation yet
	State *string `json:"state,omitempty" xmlrpc:"state"`

	// no documentation yet
	Type *Account_Contact_Type `json:"type,omitempty" xmlrpc:"type"`

	// no documentation yet
	TypeId *int `json:"typeId,omitempty" xmlrpc:"typeId"`

	// no documentation yet
	Url *string `json:"url,omitempty" xmlrpc:"url"`
}

// no documentation yet
type Account_Contact_Type struct {
	Entity

	// no documentation yet
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// no documentation yet
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// no documentation yet
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// no documentation yet
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// no documentation yet
type Account_Historical_Report struct {
	Entity
}

// no documentation yet
type Account_Link struct {
	Entity

	// no documentation yet
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// no documentation yet
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// no documentation yet
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// no documentation yet
	DestinationAccountAlphanumericId *string `json:"destinationAccountAlphanumericId,omitempty" xmlrpc:"destinationAccountAlphanumericId"`

	// no documentation yet
	DestinationAccountId *int `json:"destinationAccountId,omitempty" xmlrpc:"destinationAccountId"`

	// no documentation yet
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	ServiceProvider *Service_Provider `json:"serviceProvider,omitempty" xmlrpc:"serviceProvider"`

	// no documentation yet
	ServiceProviderId *int `json:"serviceProviderId,omitempty" xmlrpc:"serviceProviderId"`
}

// no documentation yet
type Account_Link_Bluemix struct {
	Account_Link
}

// no documentation yet
type Account_Link_OpenStack struct {
	Account_Link

	// Pseudonym for destinationAccountAlphanumericId
	DomainId *string `json:"domainId,omitempty" xmlrpc:"domainId"`
}

// OpenStack domain creation details
type Account_Link_OpenStack_DomainCreationDetails struct {
	Entity

	// Id for the domain this user was added to.
	DomainId *string `json:"domainId,omitempty" xmlrpc:"domainId"`

	// Id for the user given the Cloud Admin role for this domain.
	UserId *string `json:"userId,omitempty" xmlrpc:"userId"`

	// Name for the user given the Cloud Admin role for this domain.
	UserName *string `json:"userName,omitempty" xmlrpc:"userName"`
}

// Details required for OpenStack link request
type Account_Link_OpenStack_LinkRequest struct {
	Entity

	// Optional password
	DesiredPassword *string `json:"desiredPassword,omitempty" xmlrpc:"desiredPassword"`

	// Optional projectName
	DesiredProjectName *string `json:"desiredProjectName,omitempty" xmlrpc:"desiredProjectName"`

	// Required username
	DesiredUsername *string `json:"desiredUsername,omitempty" xmlrpc:"desiredUsername"`
}

// OpenStack project creation details
type Account_Link_OpenStack_ProjectCreationDetails struct {
	Entity

	// Id for the domain this project was added to.
	DomainId *string `json:"domainId,omitempty" xmlrpc:"domainId"`

	// Id for this project.
	ProjectId *string `json:"projectId,omitempty" xmlrpc:"projectId"`

	// Name for this project.
	ProjectName *string `json:"projectName,omitempty" xmlrpc:"projectName"`

	// Id for the user given the Project Admin role for this project.
	UserId *string `json:"userId,omitempty" xmlrpc:"userId"`

	// Name for the user given the Project Admin role for this project.
	UserName *string `json:"userName,omitempty" xmlrpc:"userName"`
}

// OpenStack project details
type Account_Link_OpenStack_ProjectDetails struct {
	Entity

	// Id for this project.
	ProjectId *string `json:"projectId,omitempty" xmlrpc:"projectId"`

	// Name for this project.
	ProjectName *string `json:"projectName,omitempty" xmlrpc:"projectName"`
}

// no documentation yet
type Account_Link_ThePlanet struct {
	Account_Link
}

// no documentation yet
type Account_Link_Vendor struct {
	Entity

	// no documentation yet
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// The SoftLayer_Account_Lockdown_Request data type holds information on API requests from brand customers.
type Account_Lockdown_Request struct {
	Entity

	// Account ID associated with this lockdown request.
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// Type of request.
	Action *string `json:"action,omitempty" xmlrpc:"action"`

	// Timestamp when the lockdown request was initially made.
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// ID of this lockdown request.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// Timestamp when the lockdown request was modified.
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// Status of the lockdown request denoting whether it's been completed.
	Status *string `json:"status,omitempty" xmlrpc:"status"`
}

// no documentation yet
type Account_MasterServiceAgreement struct {
	Entity

	// no documentation yet
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// no documentation yet
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// no documentation yet
	Guid *string `json:"guid,omitempty" xmlrpc:"guid"`

	// no documentation yet
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// The SoftLayer_Account_Media data type contains information on a single piece of media associated with a Data Transfer Service request.
type Account_Media struct {
	Entity

	// The account to which the media belongs.
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// The customer user who created the media object.
	CreateUser *User_Customer `json:"createUser,omitempty" xmlrpc:"createUser"`

	// The datacenter where the media resides.
	Datacenter *Location `json:"datacenter,omitempty" xmlrpc:"datacenter"`

	// The description of the media.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// The unique id of the media.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// The employee who last modified the media.
	ModifyEmployee *User_Employee `json:"modifyEmployee,omitempty" xmlrpc:"modifyEmployee"`

	// The customer user who last modified the media.
	ModifyUser *User_Customer `json:"modifyUser,omitempty" xmlrpc:"modifyUser"`

	// The request to which the media belongs.
	Request *Account_Media_Data_Transfer_Request `json:"request,omitempty" xmlrpc:"request"`

	// The request id of the media.
	RequestId *int `json:"requestId,omitempty" xmlrpc:"requestId"`

	// The manufacturer's serial number of the media.
	SerialNumber *string `json:"serialNumber,omitempty" xmlrpc:"serialNumber"`

	// The media's type.
	Type *Account_Media_Type `json:"type,omitempty" xmlrpc:"type"`

	// The type id of the media.
	TypeId *int `json:"typeId,omitempty" xmlrpc:"typeId"`

	// A guest's associated EVault network storage service account.
	Volume *Network_Storage `json:"volume,omitempty" xmlrpc:"volume"`
}

// The SoftLayer_Account_Media_Data_Transfer_Request data type contains information on a single Data Transfer Service request. Creation of these requests is limited to SoftLayer customers through the SoftLayer Customer Portal.
type Account_Media_Data_Transfer_Request struct {
	Entity

	// The account to which the request belongs.
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// The account id of the request.
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// A count of the active tickets that are attached to the data transfer request.
	ActiveTicketCount *uint `json:"activeTicketCount,omitempty" xmlrpc:"activeTicketCount"`

	// The active tickets that are attached to the data transfer request.
	ActiveTickets []Ticket `json:"activeTickets,omitempty" xmlrpc:"activeTickets"`

	// The billing item for the original request.
	BillingItem *Billing_Item `json:"billingItem,omitempty" xmlrpc:"billingItem"`

	// The customer user who created the request.
	CreateUser *User_Customer `json:"createUser,omitempty" xmlrpc:"createUser"`

	// The create user id of the request.
	CreateUserId *int `json:"createUserId,omitempty" xmlrpc:"createUserId"`

	// The end date of the request.
	EndDate *Time `json:"endDate,omitempty" xmlrpc:"endDate"`

	// The unique id of the request.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// The media of the request.
	Media *Account_Media `json:"media,omitempty" xmlrpc:"media"`

	// The employee who last modified the request.
	ModifyEmployee *User_Employee `json:"modifyEmployee,omitempty" xmlrpc:"modifyEmployee"`

	// The customer user who last modified the request.
	ModifyUser *User_Customer `json:"modifyUser,omitempty" xmlrpc:"modifyUser"`

	// The modify user id of the request.
	ModifyUserId *int `json:"modifyUserId,omitempty" xmlrpc:"modifyUserId"`

	// A count of the shipments of the request.
	ShipmentCount *uint `json:"shipmentCount,omitempty" xmlrpc:"shipmentCount"`

	// The shipments of the request.
	Shipments []Account_Shipment `json:"shipments,omitempty" xmlrpc:"shipments"`

	// The start date of the request.
	StartDate *Time `json:"startDate,omitempty" xmlrpc:"startDate"`

	// The status of the request.
	Status *Account_Media_Data_Transfer_Request_Status `json:"status,omitempty" xmlrpc:"status"`

	// The status id of the request.
	StatusId *int `json:"statusId,omitempty" xmlrpc:"statusId"`

	// A count of all tickets that are attached to the data transfer request.
	TicketCount *uint `json:"ticketCount,omitempty" xmlrpc:"ticketCount"`

	// All tickets that are attached to the data transfer request.
	Tickets []Ticket `json:"tickets,omitempty" xmlrpc:"tickets"`
}

// The SoftLayer_Account_Media_Data_Transfer_Request_Status data type contains general information relating to the statuses to which a Data Transfer Request may be set.
type Account_Media_Data_Transfer_Request_Status struct {
	Entity

	// The description of the request status.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// The unique id of the request status.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// The unique keyname of the request status.
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// The name of the request status.
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// The SoftLayer_Account_Media_Type data type contains general information relating to the different types of media devices that SoftLayer currently supports, as part of the Data Transfer Request Service. Such devices as USB hard drives and flash drives, as well as optical media such as CD and DVD are currently supported.
type Account_Media_Type struct {
	Entity

	// The description of the media type.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// The unique id of the media type.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// The unique keyname of the media type.
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// The name of the media type.
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// The SoftLayer_Account_Network_Vlan_Span data type exposes the setting which controls the automatic spanning of private VLANs attached to a given customers account.
type Account_Network_Vlan_Span struct {
	Entity

	// The SoftLayer customer account associated with a VLAN.
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// Flag indicating whether the customer wishes to have all private network VLANs associated with account automatically joined [0 or 1]
	EnabledFlag *bool `json:"enabledFlag,omitempty" xmlrpc:"enabledFlag"`

	// The unique internal identifier of the SoftLayer_Account_Network_Vlan_Span object.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// Timestamp of the last time the ACL for this account was applied.
	LastAppliedDate *Time `json:"lastAppliedDate,omitempty" xmlrpc:"lastAppliedDate"`

	// Timestamp of the last time the subnet hash was verified for this VLAN span record.
	LastVerifiedDate *Time `json:"lastVerifiedDate,omitempty" xmlrpc:"lastVerifiedDate"`

	// Timestamp of the last edit of the record.
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`
}

// no documentation yet
type Account_Note struct {
	Entity

	// no documentation yet
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// no documentation yet
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// no documentation yet
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// no documentation yet
	Customer *User_Customer `json:"customer,omitempty" xmlrpc:"customer"`

	// no documentation yet
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// no documentation yet
	Note *string `json:"note,omitempty" xmlrpc:"note"`

	// no documentation yet
	NoteHistory []Account_Note_History `json:"noteHistory,omitempty" xmlrpc:"noteHistory"`

	// A count of
	NoteHistoryCount *uint `json:"noteHistoryCount,omitempty" xmlrpc:"noteHistoryCount"`

	// no documentation yet
	NoteType *Account_Note_Type `json:"noteType,omitempty" xmlrpc:"noteType"`

	// no documentation yet
	NoteTypeId *int `json:"noteTypeId,omitempty" xmlrpc:"noteTypeId"`

	// no documentation yet
	UserId *int `json:"userId,omitempty" xmlrpc:"userId"`
}

// no documentation yet
type Account_Note_History struct {
	Entity

	// no documentation yet
	AccountNote *Account_Note `json:"accountNote,omitempty" xmlrpc:"accountNote"`

	// no documentation yet
	AccountNoteId *int `json:"accountNoteId,omitempty" xmlrpc:"accountNoteId"`

	// no documentation yet
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// no documentation yet
	Customer *User_Customer `json:"customer,omitempty" xmlrpc:"customer"`

	// no documentation yet
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// no documentation yet
	Note *string `json:"note,omitempty" xmlrpc:"note"`

	// no documentation yet
	UserId *int `json:"userId,omitempty" xmlrpc:"userId"`
}

// no documentation yet
type Account_Note_Type struct {
	Entity

	// no documentation yet
	BrandId *int `json:"brandId,omitempty" xmlrpc:"brandId"`

	// no documentation yet
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// no documentation yet
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// no documentation yet
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// no documentation yet
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// no documentation yet
	ValueExpression *string `json:"valueExpression,omitempty" xmlrpc:"valueExpression"`
}

// no documentation yet
type Account_Partner_Referral_Prospect struct {
	User_Customer_Prospect

	// no documentation yet
	CompanyName *string `json:"companyName,omitempty" xmlrpc:"companyName"`

	// no documentation yet
	EmailAddress *string `json:"emailAddress,omitempty" xmlrpc:"emailAddress"`

	// no documentation yet
	FirstName *string `json:"firstName,omitempty" xmlrpc:"firstName"`

	// no documentation yet
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	LastName *string `json:"lastName,omitempty" xmlrpc:"lastName"`
}

// The SoftLayer_Account_Password contains username, passwords and notes for services that may require for external applications such the Webcc interface for the EVault Storage service.
type Account_Password struct {
	Entity

	// no documentation yet
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// The SoftLayer customer account id that a username/password combination is associated with.
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// A username/password combination's internal identifier.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// A simple description of a username/password combination. These notes don't affect portal functionality.
	Notes *string `json:"notes,omitempty" xmlrpc:"notes"`

	// The password portion of a username/password combination.
	Password *string `json:"password,omitempty" xmlrpc:"password"`

	// The service that an account/password combination is tied to.
	Type *Account_Password_Type `json:"type,omitempty" xmlrpc:"type"`

	// An identifier relating to a username/password combinations's associated service.
	TypeId *int `json:"typeId,omitempty" xmlrpc:"typeId"`

	// The username portion of a username/password combination.
	Username *string `json:"username,omitempty" xmlrpc:"username"`
}

// Every username and password combination associated with a SoftLayer customer account belongs to a service that SoftLayer provides. The relationship between a username/password and it's service is provided by the SoftLayer_Account_Password_Type data type. Each username/password belongs to a single service type.
type Account_Password_Type struct {
	Entity

	// A description of the use for the account username/password combination.
	Description *string `json:"description,omitempty" xmlrpc:"description"`
}

//
//
//
//
//
type Account_Regional_Registry_Detail struct {
	Entity

	// The account that this detail object belongs to.
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// The detail object's associated [[SoftLayer_Account|account]] id
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// The date and time the detail object was created
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// A count of references to the [[SoftLayer_Network_Subnet_Registration|registration objects]] that consume this detail object.
	DetailCount *uint `json:"detailCount,omitempty" xmlrpc:"detailCount"`

	// The associated type of this detail object.
	DetailType *Account_Regional_Registry_Detail_Type `json:"detailType,omitempty" xmlrpc:"detailType"`

	// The detail object's associated [[SoftLayer_Account_Regional_Registry_Detail_Type|type]] id
	DetailTypeId *int `json:"detailTypeId,omitempty" xmlrpc:"detailTypeId"`

	// References to the [[SoftLayer_Network_Subnet_Registration|registration objects]] that consume this detail object.
	Details []Network_Subnet_Registration_Details `json:"details,omitempty" xmlrpc:"details"`

	// Unique ID of the detail object
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// The date and time the detail object was last modified
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// The individual properties that define this detail object's values.
	Properties []Account_Regional_Registry_Detail_Property `json:"properties,omitempty" xmlrpc:"properties"`

	// A count of the individual properties that define this detail object's values.
	PropertyCount *uint `json:"propertyCount,omitempty" xmlrpc:"propertyCount"`

	// The associated RWhois handle of this detail object. Used only when detailed reassignments are necessary.
	RegionalInternetRegistryHandle *Account_Rwhois_Handle `json:"regionalInternetRegistryHandle,omitempty" xmlrpc:"regionalInternetRegistryHandle"`

	// The detail object's associated [[SoftLayer_Account_Rwhois_Handle|RIR handle]] id
	RegionalInternetRegistryHandleId *int `json:"regionalInternetRegistryHandleId,omitempty" xmlrpc:"regionalInternetRegistryHandleId"`
}

// Subnet registration properties are used to define various attributes of the [[SoftLayer_Account_Regional_Registry_Detail|detail objects]]. These properties are defined by the [[SoftLayer_Account_Regional_Registry_Detail_Property_Type]] objects, which describe the available value formats.
type Account_Regional_Registry_Detail_Property struct {
	Entity

	// no documentation yet
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// The [[SoftLayer_Account_Regional_Registry_Detail]] object this property belongs to
	Detail *Account_Regional_Registry_Detail `json:"detail,omitempty" xmlrpc:"detail"`

	// Unique ID of the property object
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// The [[SoftLayer_Account_Regional_Registry_Detail_Property_Type]] object this property belongs to
	PropertyType *Account_Regional_Registry_Detail_Property_Type `json:"propertyType,omitempty" xmlrpc:"propertyType"`

	// The numeric ID of the related [[SoftLayer_Account_Regional_Registry_Detail_Property_Type|property type object]]
	PropertyTypeId *int `json:"propertyTypeId,omitempty" xmlrpc:"propertyTypeId"`

	// The numeric ID of the related [[SoftLayer_Account_Regional_Registry_Detail|detail object]]
	RegistrationDetailId *int `json:"registrationDetailId,omitempty" xmlrpc:"registrationDetailId"`

	// When multiple properties exist for a property type, defines the position in the sequence of those properties
	SequencePosition *int `json:"sequencePosition,omitempty" xmlrpc:"sequencePosition"`

	// The value of the property
	Value *string `json:"value,omitempty" xmlrpc:"value"`
}

// Subnet Registration Detail Property Type objects describe the nature of a [[SoftLayer_Account_Regional_Registry_Detail_Property]] object. These types use [http://php.net/pcre.pattern.php Perl-Compatible Regular Expressions] to validate the value of a property object.
type Account_Regional_Registry_Detail_Property_Type struct {
	Entity

	// no documentation yet
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// Unique numeric ID of the property type object
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// Code-friendly string name of the property type
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// no documentation yet
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// Human-readable name of the property type
	Name *string `json:"name,omitempty" xmlrpc:"name"`

	// A Perl-compatible regular expression used to describe the valid format of the property
	ValueExpression *string `json:"valueExpression,omitempty" xmlrpc:"valueExpression"`
}

// Subnet Registration Detail Type objects describe the nature of a [[SoftLayer_Account_Regional_Registry_Detail]] object.
//
// The standard values for these objects are as follows: <ul> <li><strong>NETWORK</strong> - The detail object represents the information for a [[SoftLayer_Network_Subnet|subnet]]</li> <li><strong>NETWORK6</strong> - The detail object represents the information for an [[SoftLayer_Network_Subnet_Version6|IPv6 subnet]]</li> <li><strong>PERSON</strong> - The detail object represents the information for a customer with the RIR</li> </ul>
type Account_Regional_Registry_Detail_Type struct {
	Entity

	// no documentation yet
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// Unique numeric ID of the detail type object
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// Code-friendly string name of the detail type
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// no documentation yet
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// Human-readable name of the detail type
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// The SoftLayer_Account_Regional_Registry_Detail_Version4_Person_Default data type contains general information relating to a single SoftLayer RIR account. RIR account information in this type such as names, addresses, and phone numbers are assigned to the registry only and not to users belonging to the account.
type Account_Regional_Registry_Detail_Version4_Person_Default struct {
	Account_Regional_Registry_Detail
}

// no documentation yet
type Account_Reports_Request struct {
	Entity

	// no documentation yet
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// A request's corresponding external contact, if one exists.
	AccountContact *Account_Contact `json:"accountContact,omitempty" xmlrpc:"accountContact"`

	// no documentation yet
	AccountContactId *int `json:"accountContactId,omitempty" xmlrpc:"accountContactId"`

	// no documentation yet
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// no documentation yet
	ComplianceReportTypeId *string `json:"complianceReportTypeId,omitempty" xmlrpc:"complianceReportTypeId"`

	// no documentation yet
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// no documentation yet
	EmployeeRecordId *int `json:"employeeRecordId,omitempty" xmlrpc:"employeeRecordId"`

	// no documentation yet
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`

	// no documentation yet
	Nda *string `json:"nda,omitempty" xmlrpc:"nda"`

	// no documentation yet
	Notes *string `json:"notes,omitempty" xmlrpc:"notes"`

	// no documentation yet
	Report *string `json:"report,omitempty" xmlrpc:"report"`

	// Type of the report customer is requesting for.
	ReportType *Compliance_Report_Type `json:"reportType,omitempty" xmlrpc:"reportType"`

	// no documentation yet
	RequestKey *string `json:"requestKey,omitempty" xmlrpc:"requestKey"`

	// no documentation yet
	Status *string `json:"status,omitempty" xmlrpc:"status"`

	// no documentation yet
	Ticket *Ticket `json:"ticket,omitempty" xmlrpc:"ticket"`

	// no documentation yet
	TicketId *int `json:"ticketId,omitempty" xmlrpc:"ticketId"`

	// The customer user that initiated a report request.
	User *User_Customer `json:"user,omitempty" xmlrpc:"user"`

	// no documentation yet
	UsrRecordId *int `json:"usrRecordId,omitempty" xmlrpc:"usrRecordId"`
}

// Provides a means of tracking handle identifiers at the various regional internet registries (RIRs). These objects are used by the [[SoftLayer_Network_Subnet_Registration (type)|SoftLayer_Network_Subnet_Registration]] objects to identify a customer or organization when a subnet is registered.
type Account_Rwhois_Handle struct {
	Entity

	// The account that this handle belongs to.
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// The handle object's associated [[SoftLayer_Account|account]] id
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// no documentation yet
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// The handle object's unique identifier as assigned by the RIR.
	Handle *string `json:"handle,omitempty" xmlrpc:"handle"`

	// Unique ID of the handle object
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	ModifyDate *Time `json:"modifyDate,omitempty" xmlrpc:"modifyDate"`
}

// The SoftLayer_Account_Shipment data type contains information relating to a shipment. Basic information such as addresses, the shipment courier, and any tracking information for as shipment is accessible with this data type.
type Account_Shipment struct {
	Entity

	// The account to which the shipment belongs.
	Account *Account `json:"account,omitempty" xmlrpc:"account"`

	// The account id of the shipment.
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId"`

	// The courier handling the shipment.
	Courier *Auxiliary_Shipping_Courier `json:"courier,omitempty" xmlrpc:"courier"`

	// The courier id of the shipment.
	CourierId *int `json:"courierId,omitempty" xmlrpc:"courierId"`

	// The courier name of the shipment.
	CourierName *string `json:"courierName,omitempty" xmlrpc:"courierName"`

	// The employee who created the shipment.
	CreateEmployee *User_Employee `json:"createEmployee,omitempty" xmlrpc:"createEmployee"`

	// The customer user who created the shipment.
	CreateUser *User_Customer `json:"createUser,omitempty" xmlrpc:"createUser"`

	// The create user id of the shipment.
	CreateUserId *int `json:"createUserId,omitempty" xmlrpc:"createUserId"`

	// The address at which the shipment is received.
	DestinationAddress *Account_Address `json:"destinationAddress,omitempty" xmlrpc:"destinationAddress"`

	// The destination address id of the shipment.
	DestinationAddressId *int `json:"destinationAddressId,omitempty" xmlrpc:"destinationAddressId"`

	// The destination date of the shipment.
	DestinationDate *Time `json:"destinationDate,omitempty" xmlrpc:"destinationDate"`

	// The unique id of the shipment.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// The employee who last modified the shipment.
	ModifyEmployee *User_Employee `json:"modifyEmployee,omitempty" xmlrpc:"modifyEmployee"`

	// The customer user who last modified the shipment.
	ModifyUser *User_Customer `json:"modifyUser,omitempty" xmlrpc:"modifyUser"`

	// The modify user id of the shipment.
	ModifyUserId *int `json:"modifyUserId,omitempty" xmlrpc:"modifyUserId"`

	// The shipment note (special handling instructions).
	Note *string `json:"note,omitempty" xmlrpc:"note"`

	// The address from which the shipment is sent.
	OriginationAddress *Account_Address `json:"originationAddress,omitempty" xmlrpc:"originationAddress"`

	// The origination address id of the shipment.
	OriginationAddressId *int `json:"originationAddressId,omitempty" xmlrpc:"originationAddressId"`

	// The origination date of the shipment.
	OriginationDate *Time `json:"originationDate,omitempty" xmlrpc:"originationDate"`

	// A count of the items in the shipment.
	ShipmentItemCount *uint `json:"shipmentItemCount,omitempty" xmlrpc:"shipmentItemCount"`

	// The items in the shipment.
	ShipmentItems []Account_Shipment_Item `json:"shipmentItems,omitempty" xmlrpc:"shipmentItems"`

	// The status of the shipment.
	Status *Account_Shipment_Status `json:"status,omitempty" xmlrpc:"status"`

	// The status id of the shipment.
	StatusId *int `json:"statusId,omitempty" xmlrpc:"statusId"`

	// The tracking data for the shipment.
	TrackingData []Account_Shipment_Tracking_Data `json:"trackingData,omitempty" xmlrpc:"trackingData"`

	// A count of the tracking data for the shipment.
	TrackingDataCount *uint `json:"trackingDataCount,omitempty" xmlrpc:"trackingDataCount"`

	// The type of shipment (e.g. for Data Transfer Service or Colocation Service).
	Type *Account_Shipment_Type `json:"type,omitempty" xmlrpc:"type"`

	// The type id of the shipment.
	TypeId *int `json:"typeId,omitempty" xmlrpc:"typeId"`
}

// The SoftLayer_Account_Shipment_Item data type contains information relating to a shipment's item. Basic information such as addresses, the shipment courier, and any tracking information for as shipment is accessible with this data type.
type Account_Shipment_Item struct {
	Entity

	// no documentation yet
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// The description of the shipping item.
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// The unique id of the shipping item.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// The package id of the shipping item.
	PackageId *int `json:"packageId,omitempty" xmlrpc:"packageId"`

	// The shipment to which this item belongs.
	Shipment *Account_Shipment `json:"shipment,omitempty" xmlrpc:"shipment"`

	// The shipment id of the shipping item.
	ShipmentId *int `json:"shipmentId,omitempty" xmlrpc:"shipmentId"`

	// The item id of the shipping item.
	ShipmentItemId *int `json:"shipmentItemId,omitempty" xmlrpc:"shipmentItemId"`

	// The type of this shipment item.
	ShipmentItemType *Account_Shipment_Item_Type `json:"shipmentItemType,omitempty" xmlrpc:"shipmentItemType"`

	// The item type id of the shipping item.
	ShipmentItemTypeId *int `json:"shipmentItemTypeId,omitempty" xmlrpc:"shipmentItemTypeId"`
}

// no documentation yet
type Account_Shipment_Item_Type struct {
	Entity

	// DEPRECATED
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// no documentation yet
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// no documentation yet
type Account_Shipment_Resource_Type struct {
	Entity
}

// no documentation yet
type Account_Shipment_Status struct {
	Entity

	// DEPRECATED
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// no documentation yet
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// The SoftLayer_Account_Shipment_Tracking_Data data type contains information on a single piece of tracking information pertaining to a shipment. This tracking information tracking numbers by which the shipment may be tracked through the shipping courier.
type Account_Shipment_Tracking_Data struct {
	Entity

	// The employee who created the tracking datum.
	CreateEmployee *User_Employee `json:"createEmployee,omitempty" xmlrpc:"createEmployee"`

	// The customer user who created the tracking datum.
	CreateUser *User_Customer `json:"createUser,omitempty" xmlrpc:"createUser"`

	// The create user id of the tracking data.
	CreateUserId *int `json:"createUserId,omitempty" xmlrpc:"createUserId"`

	// The unique id of the tracking data.
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// The employee who last modified the tracking datum.
	ModifyEmployee *User_Employee `json:"modifyEmployee,omitempty" xmlrpc:"modifyEmployee"`

	// The customer user who last modified the tracking datum.
	ModifyUser *User_Customer `json:"modifyUser,omitempty" xmlrpc:"modifyUser"`

	// The user id of the tracking data.
	ModifyUserId *int `json:"modifyUserId,omitempty" xmlrpc:"modifyUserId"`

	// The package id of the tracking data.
	PackageId *int `json:"packageId,omitempty" xmlrpc:"packageId"`

	// The sequence of the tracking data.
	Sequence *int `json:"sequence,omitempty" xmlrpc:"sequence"`

	// The shipment of the tracking datum.
	Shipment *Account_Shipment `json:"shipment,omitempty" xmlrpc:"shipment"`

	// The shipment id of the tracking data.
	ShipmentId *int `json:"shipmentId,omitempty" xmlrpc:"shipmentId"`

	// The tracking data (tracking number/reference number).
	TrackingData *string `json:"trackingData,omitempty" xmlrpc:"trackingData"`
}

// no documentation yet
type Account_Shipment_Type struct {
	Entity

	// DEPRECATED
	CreateDate *Time `json:"createDate,omitempty" xmlrpc:"createDate"`

	// no documentation yet
	Description *string `json:"description,omitempty" xmlrpc:"description"`

	// no documentation yet
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	KeyName *string `json:"keyName,omitempty" xmlrpc:"keyName"`

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}

// no documentation yet
type Account_Status struct {
	Entity

	// no documentation yet
	Id *int `json:"id,omitempty" xmlrpc:"id"`

	// no documentation yet
	Name *string `json:"name,omitempty" xmlrpc:"name"`
}
