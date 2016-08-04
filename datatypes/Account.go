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

type Account struct {
	Entity

	AbuseEmail                                             *string                                                                 `json:"abuseEmail,omitempty"`
	AbuseEmailCount                                        *uint                                                                   `json:"abuseEmailCount,omitempty"`
	AbuseEmails                                            []Account_AbuseEmail                                                    `json:"abuseEmails,omitempty"`
	AccountContactCount                                    *uint                                                                   `json:"accountContactCount,omitempty"`
	AccountContacts                                        []Account_Contact                                                       `json:"accountContacts,omitempty"`
	AccountLicenseCount                                    *uint                                                                   `json:"accountLicenseCount,omitempty"`
	AccountLicenses                                        []Software_AccountLicense                                               `json:"accountLicenses,omitempty"`
	AccountLinkCount                                       *uint                                                                   `json:"accountLinkCount,omitempty"`
	AccountLinks                                           []Account_Link                                                          `json:"accountLinks,omitempty"`
	AccountManagedResourcesFlag                            *bool                                                                   `json:"accountManagedResourcesFlag,omitempty"`
	AccountStatus                                          *Account_Status                                                         `json:"accountStatus,omitempty"`
	AccountStatusId                                        *int                                                                    `json:"accountStatusId,omitempty"`
	ActiveAccountDiscountBillingItem                       *Billing_Item                                                           `json:"activeAccountDiscountBillingItem,omitempty"`
	ActiveAccountLicenseCount                              *uint                                                                   `json:"activeAccountLicenseCount,omitempty"`
	ActiveAccountLicenses                                  []Software_AccountLicense                                               `json:"activeAccountLicenses,omitempty"`
	ActiveAddressCount                                     *uint                                                                   `json:"activeAddressCount,omitempty"`
	ActiveAddresses                                        []Account_Address                                                       `json:"activeAddresses,omitempty"`
	ActiveBillingAgreementCount                            *uint                                                                   `json:"activeBillingAgreementCount,omitempty"`
	ActiveBillingAgreements                                []Account_Agreement                                                     `json:"activeBillingAgreements,omitempty"`
	ActiveCatalystEnrollment                               *Catalyst_Enrollment                                                    `json:"activeCatalystEnrollment,omitempty"`
	ActiveColocationContainerCount                         *uint                                                                   `json:"activeColocationContainerCount,omitempty"`
	ActiveColocationContainers                             []Billing_Item                                                          `json:"activeColocationContainers,omitempty"`
	ActiveFlexibleCreditEnrollment                         *FlexibleCredit_Enrollment                                              `json:"activeFlexibleCreditEnrollment,omitempty"`
	ActiveNotificationSubscriberCount                      *uint                                                                   `json:"activeNotificationSubscriberCount,omitempty"`
	ActiveNotificationSubscribers                          []Notification_Subscriber                                               `json:"activeNotificationSubscribers,omitempty"`
	ActiveQuoteCount                                       *uint                                                                   `json:"activeQuoteCount,omitempty"`
	ActiveQuotes                                           []Billing_Order_Quote                                                   `json:"activeQuotes,omitempty"`
	ActiveVirtualLicenseCount                              *uint                                                                   `json:"activeVirtualLicenseCount,omitempty"`
	ActiveVirtualLicenses                                  []Software_VirtualLicense                                               `json:"activeVirtualLicenses,omitempty"`
	AdcLoadBalancerCount                                   *uint                                                                   `json:"adcLoadBalancerCount,omitempty"`
	AdcLoadBalancers                                       []Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"adcLoadBalancers,omitempty"`
	Address1                                               *string                                                                 `json:"address1,omitempty"`
	Address2                                               *string                                                                 `json:"address2,omitempty"`
	AddressCount                                           *uint                                                                   `json:"addressCount,omitempty"`
	Addresses                                              []Account_Address                                                       `json:"addresses,omitempty"`
	AffiliateId                                            *string                                                                 `json:"affiliateId,omitempty"`
	AllBillingItems                                        []Billing_Item                                                          `json:"allBillingItems,omitempty"`
	AllCommissionBillingItemCount                          *uint                                                                   `json:"allCommissionBillingItemCount,omitempty"`
	AllCommissionBillingItems                              []Billing_Item                                                          `json:"allCommissionBillingItems,omitempty"`
	AllRecurringTopLevelBillingItemCount                   *uint                                                                   `json:"allRecurringTopLevelBillingItemCount,omitempty"`
	AllRecurringTopLevelBillingItems                       []Billing_Item                                                          `json:"allRecurringTopLevelBillingItems,omitempty"`
	AllRecurringTopLevelBillingItemsUnfiltered             []Billing_Item                                                          `json:"allRecurringTopLevelBillingItemsUnfiltered,omitempty"`
	AllRecurringTopLevelBillingItemsUnfilteredCount        *uint                                                                   `json:"allRecurringTopLevelBillingItemsUnfilteredCount,omitempty"`
	AllSubnetBillingItemCount                              *uint                                                                   `json:"allSubnetBillingItemCount,omitempty"`
	AllSubnetBillingItems                                  []Billing_Item                                                          `json:"allSubnetBillingItems,omitempty"`
	AllTopLevelBillingItemCount                            *uint                                                                   `json:"allTopLevelBillingItemCount,omitempty"`
	AllTopLevelBillingItems                                []Billing_Item                                                          `json:"allTopLevelBillingItems,omitempty"`
	AllTopLevelBillingItemsUnfiltered                      []Billing_Item                                                          `json:"allTopLevelBillingItemsUnfiltered,omitempty"`
	AllTopLevelBillingItemsUnfilteredCount                 *uint                                                                   `json:"allTopLevelBillingItemsUnfilteredCount,omitempty"`
	AllowedPptpVpnQuantity                                 *int                                                                    `json:"allowedPptpVpnQuantity,omitempty"`
	AllowsBluemixAccountLinkingFlag                        *bool                                                                   `json:"allowsBluemixAccountLinkingFlag,omitempty"`
	AlternatePhone                                         *string                                                                 `json:"alternatePhone,omitempty"`
	ApplicationDeliveryControllerCount                     *uint                                                                   `json:"applicationDeliveryControllerCount,omitempty"`
	ApplicationDeliveryControllers                         []Network_Application_Delivery_Controller                               `json:"applicationDeliveryControllers,omitempty"`
	AttributeCount                                         *uint                                                                   `json:"attributeCount,omitempty"`
	Attributes                                             []Account_Attribute                                                     `json:"attributes,omitempty"`
	AvailablePublicNetworkVlanCount                        *uint                                                                   `json:"availablePublicNetworkVlanCount,omitempty"`
	AvailablePublicNetworkVlans                            []Network_Vlan                                                          `json:"availablePublicNetworkVlans,omitempty"`
	Balance                                                *float64                                                                `json:"balance,omitempty"`
	BandwidthAllotmentCount                                *uint                                                                   `json:"bandwidthAllotmentCount,omitempty"`
	BandwidthAllotments                                    []Network_Bandwidth_Version1_Allotment                                  `json:"bandwidthAllotments,omitempty"`
	BandwidthAllotmentsOverAllocation                      []Network_Bandwidth_Version1_Allotment                                  `json:"bandwidthAllotmentsOverAllocation,omitempty"`
	BandwidthAllotmentsOverAllocationCount                 *uint                                                                   `json:"bandwidthAllotmentsOverAllocationCount,omitempty"`
	BandwidthAllotmentsProjectedOverAllocation             []Network_Bandwidth_Version1_Allotment                                  `json:"bandwidthAllotmentsProjectedOverAllocation,omitempty"`
	BandwidthAllotmentsProjectedOverAllocationCount        *uint                                                                   `json:"bandwidthAllotmentsProjectedOverAllocationCount,omitempty"`
	BareMetalInstanceCount                                 *uint                                                                   `json:"bareMetalInstanceCount,omitempty"`
	BareMetalInstances                                     []Hardware                                                              `json:"bareMetalInstances,omitempty"`
	BillingAgreementCount                                  *uint                                                                   `json:"billingAgreementCount,omitempty"`
	BillingAgreements                                      []Account_Agreement                                                     `json:"billingAgreements,omitempty"`
	BillingInfo                                            *Billing_Info                                                           `json:"billingInfo,omitempty"`
	BlockDeviceTemplateGroupCount                          *uint                                                                   `json:"blockDeviceTemplateGroupCount,omitempty"`
	BlockDeviceTemplateGroups                              []Virtual_Guest_Block_Device_Template_Group                             `json:"blockDeviceTemplateGroups,omitempty"`
	BlueIdAuthenticationRequiredFlag                       *bool                                                                   `json:"blueIdAuthenticationRequiredFlag,omitempty"`
	BluemixLinkedFlag                                      *bool                                                                   `json:"bluemixLinkedFlag,omitempty"`
	Brand                                                  *Brand                                                                  `json:"brand,omitempty"`
	BrandAccountFlag                                       *bool                                                                   `json:"brandAccountFlag,omitempty"`
	BrandId                                                *int                                                                    `json:"brandId,omitempty"`
	BrandKeyName                                           *string                                                                 `json:"brandKeyName,omitempty"`
	CanOrderAdditionalVlansFlag                            *bool                                                                   `json:"canOrderAdditionalVlansFlag,omitempty"`
	CartCount                                              *uint                                                                   `json:"cartCount,omitempty"`
	Carts                                                  []Billing_Order_Quote                                                   `json:"carts,omitempty"`
	CatalystEnrollmentCount                                *uint                                                                   `json:"catalystEnrollmentCount,omitempty"`
	CatalystEnrollments                                    []Catalyst_Enrollment                                                   `json:"catalystEnrollments,omitempty"`
	CdnAccountCount                                        *uint                                                                   `json:"cdnAccountCount,omitempty"`
	CdnAccounts                                            []Network_ContentDelivery_Account                                       `json:"cdnAccounts,omitempty"`
	City                                                   *string                                                                 `json:"city,omitempty"`
	ClaimedTaxExemptTxFlag                                 *bool                                                                   `json:"claimedTaxExemptTxFlag,omitempty"`
	ClosedTicketCount                                      *uint                                                                   `json:"closedTicketCount,omitempty"`
	ClosedTickets                                          []Ticket                                                                `json:"closedTickets,omitempty"`
	CompanyName                                            *string                                                                 `json:"companyName,omitempty"`
	Country                                                *string                                                                 `json:"country,omitempty"`
	CreateDate                                             *time.Time                                                              `json:"createDate,omitempty"`
	DatacentersWithSubnetAllocationCount                   *uint                                                                   `json:"datacentersWithSubnetAllocationCount,omitempty"`
	DatacentersWithSubnetAllocations                       []Location                                                              `json:"datacentersWithSubnetAllocations,omitempty"`
	DeviceFingerprintId                                    *string                                                                 `json:"deviceFingerprintId,omitempty"`
	DisablePaymentProcessingFlag                           *bool                                                                   `json:"disablePaymentProcessingFlag,omitempty"`
	DisplaySupportRepresentativeAssignmentCount            *uint                                                                   `json:"displaySupportRepresentativeAssignmentCount,omitempty"`
	DisplaySupportRepresentativeAssignments                []Account_Attachment_Employee                                           `json:"displaySupportRepresentativeAssignments,omitempty"`
	DomainCount                                            *uint                                                                   `json:"domainCount,omitempty"`
	DomainRegistrationCount                                *uint                                                                   `json:"domainRegistrationCount,omitempty"`
	DomainRegistrations                                    []Dns_Domain_Registration                                               `json:"domainRegistrations,omitempty"`
	Domains                                                []Dns_Domain                                                            `json:"domains,omitempty"`
	DomainsWithoutSecondaryDnsRecordCount                  *uint                                                                   `json:"domainsWithoutSecondaryDnsRecordCount,omitempty"`
	DomainsWithoutSecondaryDnsRecords                      []Dns_Domain                                                            `json:"domainsWithoutSecondaryDnsRecords,omitempty"`
	Email                                                  *string                                                                 `json:"email,omitempty"`
	EvaultCapacityGB                                       *uint                                                                   `json:"evaultCapacityGB,omitempty"`
	EvaultMasterUserCount                                  *uint                                                                   `json:"evaultMasterUserCount,omitempty"`
	EvaultMasterUsers                                      []Account_Password                                                      `json:"evaultMasterUsers,omitempty"`
	EvaultNetworkStorage                                   []Network_Storage                                                       `json:"evaultNetworkStorage,omitempty"`
	EvaultNetworkStorageCount                              *uint                                                                   `json:"evaultNetworkStorageCount,omitempty"`
	ExpiredSecurityCertificateCount                        *uint                                                                   `json:"expiredSecurityCertificateCount,omitempty"`
	ExpiredSecurityCertificates                            []Security_Certificate                                                  `json:"expiredSecurityCertificates,omitempty"`
	FacilityLogCount                                       *uint                                                                   `json:"facilityLogCount,omitempty"`
	FacilityLogs                                           []User_Access_Facility_Log                                              `json:"facilityLogs,omitempty"`
	FaxPhone                                               *string                                                                 `json:"faxPhone,omitempty"`
	FirstName                                              *string                                                                 `json:"firstName,omitempty"`
	FlexibleCreditEnrollmentCount                          *uint                                                                   `json:"flexibleCreditEnrollmentCount,omitempty"`
	FlexibleCreditEnrollments                              []FlexibleCredit_Enrollment                                             `json:"flexibleCreditEnrollments,omitempty"`
	GlobalIpRecordCount                                    *uint                                                                   `json:"globalIpRecordCount,omitempty"`
	GlobalIpRecords                                        []Network_Subnet_IpAddress_Global                                       `json:"globalIpRecords,omitempty"`
	GlobalIpv4RecordCount                                  *uint                                                                   `json:"globalIpv4RecordCount,omitempty"`
	GlobalIpv4Records                                      []Network_Subnet_IpAddress_Global                                       `json:"globalIpv4Records,omitempty"`
	GlobalIpv6RecordCount                                  *uint                                                                   `json:"globalIpv6RecordCount,omitempty"`
	GlobalIpv6Records                                      []Network_Subnet_IpAddress_Global                                       `json:"globalIpv6Records,omitempty"`
	GlobalLoadBalancerAccountCount                         *uint                                                                   `json:"globalLoadBalancerAccountCount,omitempty"`
	GlobalLoadBalancerAccounts                             []Network_LoadBalancer_Global_Account                                   `json:"globalLoadBalancerAccounts,omitempty"`
	Hardware                                               []Hardware                                                              `json:"hardware,omitempty"`
	HardwareCount                                          *uint                                                                   `json:"hardwareCount,omitempty"`
	HardwareOverBandwidthAllocation                        []Hardware                                                              `json:"hardwareOverBandwidthAllocation,omitempty"`
	HardwareOverBandwidthAllocationCount                   *uint                                                                   `json:"hardwareOverBandwidthAllocationCount,omitempty"`
	HardwareProjectedOverBandwidthAllocation               []Hardware                                                              `json:"hardwareProjectedOverBandwidthAllocation,omitempty"`
	HardwareProjectedOverBandwidthAllocationCount          *uint                                                                   `json:"hardwareProjectedOverBandwidthAllocationCount,omitempty"`
	HardwareWithCpanel                                     []Hardware                                                              `json:"hardwareWithCpanel,omitempty"`
	HardwareWithCpanelCount                                *uint                                                                   `json:"hardwareWithCpanelCount,omitempty"`
	HardwareWithHelm                                       []Hardware                                                              `json:"hardwareWithHelm,omitempty"`
	HardwareWithHelmCount                                  *uint                                                                   `json:"hardwareWithHelmCount,omitempty"`
	HardwareWithMcafee                                     []Hardware                                                              `json:"hardwareWithMcafee,omitempty"`
	HardwareWithMcafeeAntivirusRedhat                      []Hardware                                                              `json:"hardwareWithMcafeeAntivirusRedhat,omitempty"`
	HardwareWithMcafeeAntivirusRedhatCount                 *uint                                                                   `json:"hardwareWithMcafeeAntivirusRedhatCount,omitempty"`
	HardwareWithMcafeeAntivirusWindowCount                 *uint                                                                   `json:"hardwareWithMcafeeAntivirusWindowCount,omitempty"`
	HardwareWithMcafeeAntivirusWindows                     []Hardware                                                              `json:"hardwareWithMcafeeAntivirusWindows,omitempty"`
	HardwareWithMcafeeCount                                *uint                                                                   `json:"hardwareWithMcafeeCount,omitempty"`
	HardwareWithMcafeeIntrusionDetectionSystem             []Hardware                                                              `json:"hardwareWithMcafeeIntrusionDetectionSystem,omitempty"`
	HardwareWithMcafeeIntrusionDetectionSystemCount        *uint                                                                   `json:"hardwareWithMcafeeIntrusionDetectionSystemCount,omitempty"`
	HardwareWithPlesk                                      []Hardware                                                              `json:"hardwareWithPlesk,omitempty"`
	HardwareWithPleskCount                                 *uint                                                                   `json:"hardwareWithPleskCount,omitempty"`
	HardwareWithQuantastor                                 []Hardware                                                              `json:"hardwareWithQuantastor,omitempty"`
	HardwareWithQuantastorCount                            *uint                                                                   `json:"hardwareWithQuantastorCount,omitempty"`
	HardwareWithUrchin                                     []Hardware                                                              `json:"hardwareWithUrchin,omitempty"`
	HardwareWithUrchinCount                                *uint                                                                   `json:"hardwareWithUrchinCount,omitempty"`
	HardwareWithWindowCount                                *uint                                                                   `json:"hardwareWithWindowCount,omitempty"`
	HardwareWithWindows                                    []Hardware                                                              `json:"hardwareWithWindows,omitempty"`
	HasEvaultBareMetalRestorePluginFlag                    *bool                                                                   `json:"hasEvaultBareMetalRestorePluginFlag,omitempty"`
	HasIderaBareMetalRestorePluginFlag                     *bool                                                                   `json:"hasIderaBareMetalRestorePluginFlag,omitempty"`
	HasPendingOrder                                        *uint                                                                   `json:"hasPendingOrder,omitempty"`
	HasR1softBareMetalRestorePluginFlag                    *bool                                                                   `json:"hasR1softBareMetalRestorePluginFlag,omitempty"`
	HourlyBareMetalInstanceCount                           *uint                                                                   `json:"hourlyBareMetalInstanceCount,omitempty"`
	HourlyBareMetalInstances                               []Hardware                                                              `json:"hourlyBareMetalInstances,omitempty"`
	HourlyServiceBillingItemCount                          *uint                                                                   `json:"hourlyServiceBillingItemCount,omitempty"`
	HourlyServiceBillingItems                              []Billing_Item                                                          `json:"hourlyServiceBillingItems,omitempty"`
	HourlyVirtualGuestCount                                *uint                                                                   `json:"hourlyVirtualGuestCount,omitempty"`
	HourlyVirtualGuests                                    []Virtual_Guest                                                         `json:"hourlyVirtualGuests,omitempty"`
	HubNetworkStorage                                      []Network_Storage                                                       `json:"hubNetworkStorage,omitempty"`
	HubNetworkStorageCount                                 *uint                                                                   `json:"hubNetworkStorageCount,omitempty"`
	Id                                                     *int                                                                    `json:"id,omitempty"`
	InternalNoteCount                                      *uint                                                                   `json:"internalNoteCount,omitempty"`
	InternalNotes                                          []Account_Note                                                          `json:"internalNotes,omitempty"`
	InvoiceCount                                           *uint                                                                   `json:"invoiceCount,omitempty"`
	Invoices                                               []Billing_Invoice                                                       `json:"invoices,omitempty"`
	IpAddressCount                                         *uint                                                                   `json:"ipAddressCount,omitempty"`
	IpAddresses                                            []Network_Subnet_IpAddress                                              `json:"ipAddresses,omitempty"`
	IsReseller                                             *int                                                                    `json:"isReseller,omitempty"`
	IscsiNetworkStorage                                    []Network_Storage                                                       `json:"iscsiNetworkStorage,omitempty"`
	IscsiNetworkStorageCount                               *uint                                                                   `json:"iscsiNetworkStorageCount,omitempty"`
	LastCanceledBillingItem                                *Billing_Item                                                           `json:"lastCanceledBillingItem,omitempty"`
	LastCancelledServerBillingItem                         *Billing_Item                                                           `json:"lastCancelledServerBillingItem,omitempty"`
	LastFiveClosedAbuseTicketCount                         *uint                                                                   `json:"lastFiveClosedAbuseTicketCount,omitempty"`
	LastFiveClosedAbuseTickets                             []Ticket                                                                `json:"lastFiveClosedAbuseTickets,omitempty"`
	LastFiveClosedAccountingTicketCount                    *uint                                                                   `json:"lastFiveClosedAccountingTicketCount,omitempty"`
	LastFiveClosedAccountingTickets                        []Ticket                                                                `json:"lastFiveClosedAccountingTickets,omitempty"`
	LastFiveClosedOtherTicketCount                         *uint                                                                   `json:"lastFiveClosedOtherTicketCount,omitempty"`
	LastFiveClosedOtherTickets                             []Ticket                                                                `json:"lastFiveClosedOtherTickets,omitempty"`
	LastFiveClosedSalesTicketCount                         *uint                                                                   `json:"lastFiveClosedSalesTicketCount,omitempty"`
	LastFiveClosedSalesTickets                             []Ticket                                                                `json:"lastFiveClosedSalesTickets,omitempty"`
	LastFiveClosedSupportTicketCount                       *uint                                                                   `json:"lastFiveClosedSupportTicketCount,omitempty"`
	LastFiveClosedSupportTickets                           []Ticket                                                                `json:"lastFiveClosedSupportTickets,omitempty"`
	LastFiveClosedTicketCount                              *uint                                                                   `json:"lastFiveClosedTicketCount,omitempty"`
	LastFiveClosedTickets                                  []Ticket                                                                `json:"lastFiveClosedTickets,omitempty"`
	LastName                                               *string                                                                 `json:"lastName,omitempty"`
	LateFeeProtectionFlag                                  *bool                                                                   `json:"lateFeeProtectionFlag,omitempty"`
	LatestBillDate                                         *time.Time                                                              `json:"latestBillDate,omitempty"`
	LatestRecurringInvoice                                 *Billing_Invoice                                                        `json:"latestRecurringInvoice,omitempty"`
	LatestRecurringPendingInvoice                          *Billing_Invoice                                                        `json:"latestRecurringPendingInvoice,omitempty"`
	LegacyBandwidthAllotmentCount                          *uint                                                                   `json:"legacyBandwidthAllotmentCount,omitempty"`
	LegacyBandwidthAllotments                              []Network_Bandwidth_Version1_Allotment                                  `json:"legacyBandwidthAllotments,omitempty"`
	LegacyIscsiCapacityGB                                  *uint                                                                   `json:"legacyIscsiCapacityGB,omitempty"`
	LoadBalancerCount                                      *uint                                                                   `json:"loadBalancerCount,omitempty"`
	LoadBalancers                                          []Network_LoadBalancer_VirtualIpAddress                                 `json:"loadBalancers,omitempty"`
	LockboxCapacityGB                                      *uint                                                                   `json:"lockboxCapacityGB,omitempty"`
	LockboxNetworkStorage                                  []Network_Storage                                                       `json:"lockboxNetworkStorage,omitempty"`
	LockboxNetworkStorageCount                             *uint                                                                   `json:"lockboxNetworkStorageCount,omitempty"`
	ManualPaymentsUnderReview                              []Billing_Payment_Card_ManualPayment                                    `json:"manualPaymentsUnderReview,omitempty"`
	ManualPaymentsUnderReviewCount                         *uint                                                                   `json:"manualPaymentsUnderReviewCount,omitempty"`
	MasterUser                                             *User_Customer                                                          `json:"masterUser,omitempty"`
	MediaDataTransferRequestCount                          *uint                                                                   `json:"mediaDataTransferRequestCount,omitempty"`
	MediaDataTransferRequests                              []Account_Media_Data_Transfer_Request                                   `json:"mediaDataTransferRequests,omitempty"`
	MessageQueueAccountCount                               *uint                                                                   `json:"messageQueueAccountCount,omitempty"`
	MessageQueueAccounts                                   []Network_Message_Queue                                                 `json:"messageQueueAccounts,omitempty"`
	ModifyDate                                             *time.Time                                                              `json:"modifyDate,omitempty"`
	MonthlyBareMetalInstanceCount                          *uint                                                                   `json:"monthlyBareMetalInstanceCount,omitempty"`
	MonthlyBareMetalInstances                              []Hardware                                                              `json:"monthlyBareMetalInstances,omitempty"`
	MonthlyVirtualGuestCount                               *uint                                                                   `json:"monthlyVirtualGuestCount,omitempty"`
	MonthlyVirtualGuests                                   []Virtual_Guest                                                         `json:"monthlyVirtualGuests,omitempty"`
	NasNetworkStorage                                      []Network_Storage                                                       `json:"nasNetworkStorage,omitempty"`
	NasNetworkStorageCount                                 *uint                                                                   `json:"nasNetworkStorageCount,omitempty"`
	NetworkCreationFlag                                    *bool                                                                   `json:"networkCreationFlag,omitempty"`
	NetworkGatewayCount                                    *uint                                                                   `json:"networkGatewayCount,omitempty"`
	NetworkGateways                                        []Network_Gateway                                                       `json:"networkGateways,omitempty"`
	NetworkHardware                                        []Hardware                                                              `json:"networkHardware,omitempty"`
	NetworkHardwareCount                                   *uint                                                                   `json:"networkHardwareCount,omitempty"`
	NetworkMessageDeliveryAccountCount                     *uint                                                                   `json:"networkMessageDeliveryAccountCount,omitempty"`
	NetworkMessageDeliveryAccounts                         []Network_Message_Delivery                                              `json:"networkMessageDeliveryAccounts,omitempty"`
	NetworkMonitorDownHardware                             []Hardware                                                              `json:"networkMonitorDownHardware,omitempty"`
	NetworkMonitorDownHardwareCount                        *uint                                                                   `json:"networkMonitorDownHardwareCount,omitempty"`
	NetworkMonitorDownVirtualGuestCount                    *uint                                                                   `json:"networkMonitorDownVirtualGuestCount,omitempty"`
	NetworkMonitorDownVirtualGuests                        []Virtual_Guest                                                         `json:"networkMonitorDownVirtualGuests,omitempty"`
	NetworkMonitorRecoveringHardware                       []Hardware                                                              `json:"networkMonitorRecoveringHardware,omitempty"`
	NetworkMonitorRecoveringHardwareCount                  *uint                                                                   `json:"networkMonitorRecoveringHardwareCount,omitempty"`
	NetworkMonitorRecoveringVirtualGuestCount              *uint                                                                   `json:"networkMonitorRecoveringVirtualGuestCount,omitempty"`
	NetworkMonitorRecoveringVirtualGuests                  []Virtual_Guest                                                         `json:"networkMonitorRecoveringVirtualGuests,omitempty"`
	NetworkMonitorUpHardware                               []Hardware                                                              `json:"networkMonitorUpHardware,omitempty"`
	NetworkMonitorUpHardwareCount                          *uint                                                                   `json:"networkMonitorUpHardwareCount,omitempty"`
	NetworkMonitorUpVirtualGuestCount                      *uint                                                                   `json:"networkMonitorUpVirtualGuestCount,omitempty"`
	NetworkMonitorUpVirtualGuests                          []Virtual_Guest                                                         `json:"networkMonitorUpVirtualGuests,omitempty"`
	NetworkStorage                                         []Network_Storage                                                       `json:"networkStorage,omitempty"`
	NetworkStorageCount                                    *uint                                                                   `json:"networkStorageCount,omitempty"`
	NetworkStorageGroupCount                               *uint                                                                   `json:"networkStorageGroupCount,omitempty"`
	NetworkStorageGroups                                   []Network_Storage_Group                                                 `json:"networkStorageGroups,omitempty"`
	NetworkTunnelContextCount                              *uint                                                                   `json:"networkTunnelContextCount,omitempty"`
	NetworkTunnelContexts                                  []Network_Tunnel_Module_Context                                         `json:"networkTunnelContexts,omitempty"`
	NetworkVlanCount                                       *uint                                                                   `json:"networkVlanCount,omitempty"`
	NetworkVlanSpan                                        *Account_Network_Vlan_Span                                              `json:"networkVlanSpan,omitempty"`
	NetworkVlans                                           []Network_Vlan                                                          `json:"networkVlans,omitempty"`
	NextBillingPublicAllotmentHardwareBandwidthDetailCount *uint                                                                   `json:"nextBillingPublicAllotmentHardwareBandwidthDetailCount,omitempty"`
	NextBillingPublicAllotmentHardwareBandwidthDetails     []Network_Bandwidth_Version1_Allotment                                  `json:"nextBillingPublicAllotmentHardwareBandwidthDetails,omitempty"`
	NextInvoiceIncubatorExemptTotal                        *float64                                                                `json:"nextInvoiceIncubatorExemptTotal,omitempty"`
	NextInvoiceTopLevelBillingItemCount                    *uint                                                                   `json:"nextInvoiceTopLevelBillingItemCount,omitempty"`
	NextInvoiceTopLevelBillingItems                        []Billing_Item                                                          `json:"nextInvoiceTopLevelBillingItems,omitempty"`
	NextInvoiceTotalAmount                                 *float64                                                                `json:"nextInvoiceTotalAmount,omitempty"`
	NextInvoiceTotalOneTimeAmount                          *float64                                                                `json:"nextInvoiceTotalOneTimeAmount,omitempty"`
	NextInvoiceTotalOneTimeTaxAmount                       *float64                                                                `json:"nextInvoiceTotalOneTimeTaxAmount,omitempty"`
	NextInvoiceTotalRecurringAmount                        *float64                                                                `json:"nextInvoiceTotalRecurringAmount,omitempty"`
	NextInvoiceTotalRecurringAmountBeforeAccountDiscount   *float64                                                                `json:"nextInvoiceTotalRecurringAmountBeforeAccountDiscount,omitempty"`
	NextInvoiceTotalRecurringTaxAmount                     *float64                                                                `json:"nextInvoiceTotalRecurringTaxAmount,omitempty"`
	NextInvoiceTotalTaxableRecurringAmount                 *float64                                                                `json:"nextInvoiceTotalTaxableRecurringAmount,omitempty"`
	NotificationSubscriberCount                            *uint                                                                   `json:"notificationSubscriberCount,omitempty"`
	NotificationSubscribers                                []Notification_Subscriber                                               `json:"notificationSubscribers,omitempty"`
	OfficePhone                                            *string                                                                 `json:"officePhone,omitempty"`
	OpenAbuseTicketCount                                   *uint                                                                   `json:"openAbuseTicketCount,omitempty"`
	OpenAbuseTickets                                       []Ticket                                                                `json:"openAbuseTickets,omitempty"`
	OpenAccountingTicketCount                              *uint                                                                   `json:"openAccountingTicketCount,omitempty"`
	OpenAccountingTickets                                  []Ticket                                                                `json:"openAccountingTickets,omitempty"`
	OpenBillingTicketCount                                 *uint                                                                   `json:"openBillingTicketCount,omitempty"`
	OpenBillingTickets                                     []Ticket                                                                `json:"openBillingTickets,omitempty"`
	OpenCancellationRequestCount                           *uint                                                                   `json:"openCancellationRequestCount,omitempty"`
	OpenCancellationRequests                               []Billing_Item_Cancellation_Request                                     `json:"openCancellationRequests,omitempty"`
	OpenOtherTicketCount                                   *uint                                                                   `json:"openOtherTicketCount,omitempty"`
	OpenOtherTickets                                       []Ticket                                                                `json:"openOtherTickets,omitempty"`
	OpenRecurringInvoiceCount                              *uint                                                                   `json:"openRecurringInvoiceCount,omitempty"`
	OpenRecurringInvoices                                  []Billing_Invoice                                                       `json:"openRecurringInvoices,omitempty"`
	OpenSalesTicketCount                                   *uint                                                                   `json:"openSalesTicketCount,omitempty"`
	OpenSalesTickets                                       []Ticket                                                                `json:"openSalesTickets,omitempty"`
	OpenStackAccountLinkCount                              *uint                                                                   `json:"openStackAccountLinkCount,omitempty"`
	OpenStackAccountLinks                                  []Account_Link                                                          `json:"openStackAccountLinks,omitempty"`
	OpenStackObjectStorage                                 []Network_Storage                                                       `json:"openStackObjectStorage,omitempty"`
	OpenStackObjectStorageCount                            *uint                                                                   `json:"openStackObjectStorageCount,omitempty"`
	OpenSupportTicketCount                                 *uint                                                                   `json:"openSupportTicketCount,omitempty"`
	OpenSupportTickets                                     []Ticket                                                                `json:"openSupportTickets,omitempty"`
	OpenTicketCount                                        *uint                                                                   `json:"openTicketCount,omitempty"`
	OpenTickets                                            []Ticket                                                                `json:"openTickets,omitempty"`
	OpenTicketsWaitingOnCustomer                           []Ticket                                                                `json:"openTicketsWaitingOnCustomer,omitempty"`
	OpenTicketsWaitingOnCustomerCount                      *uint                                                                   `json:"openTicketsWaitingOnCustomerCount,omitempty"`
	OrderCount                                             *uint                                                                   `json:"orderCount,omitempty"`
	Orders                                                 []Billing_Order                                                         `json:"orders,omitempty"`
	OrphanBillingItemCount                                 *uint                                                                   `json:"orphanBillingItemCount,omitempty"`
	OrphanBillingItems                                     []Billing_Item                                                          `json:"orphanBillingItems,omitempty"`
	OwnedBrandCount                                        *uint                                                                   `json:"ownedBrandCount,omitempty"`
	OwnedBrands                                            []Brand                                                                 `json:"ownedBrands,omitempty"`
	OwnedHardwareGenericComponentModelCount                *uint                                                                   `json:"ownedHardwareGenericComponentModelCount,omitempty"`
	OwnedHardwareGenericComponentModels                    []Hardware_Component_Model_Generic                                      `json:"ownedHardwareGenericComponentModels,omitempty"`
	PaymentProcessorCount                                  *uint                                                                   `json:"paymentProcessorCount,omitempty"`
	PaymentProcessors                                      []Billing_Payment_Processor                                             `json:"paymentProcessors,omitempty"`
	PendingEventCount                                      *uint                                                                   `json:"pendingEventCount,omitempty"`
	PendingEvents                                          []Notification_Occurrence_Event                                         `json:"pendingEvents,omitempty"`
	PendingInvoice                                         *Billing_Invoice                                                        `json:"pendingInvoice,omitempty"`
	PendingInvoiceTopLevelItemCount                        *uint                                                                   `json:"pendingInvoiceTopLevelItemCount,omitempty"`
	PendingInvoiceTopLevelItems                            []Billing_Invoice_Item                                                  `json:"pendingInvoiceTopLevelItems,omitempty"`
	PendingInvoiceTotalAmount                              *float64                                                                `json:"pendingInvoiceTotalAmount,omitempty"`
	PendingInvoiceTotalOneTimeAmount                       *float64                                                                `json:"pendingInvoiceTotalOneTimeAmount,omitempty"`
	PendingInvoiceTotalOneTimeTaxAmount                    *float64                                                                `json:"pendingInvoiceTotalOneTimeTaxAmount,omitempty"`
	PendingInvoiceTotalRecurringAmount                     *float64                                                                `json:"pendingInvoiceTotalRecurringAmount,omitempty"`
	PendingInvoiceTotalRecurringTaxAmount                  *float64                                                                `json:"pendingInvoiceTotalRecurringTaxAmount,omitempty"`
	PermissionGroupCount                                   *uint                                                                   `json:"permissionGroupCount,omitempty"`
	PermissionGroups                                       []User_Permission_Group                                                 `json:"permissionGroups,omitempty"`
	PermissionRoleCount                                    *uint                                                                   `json:"permissionRoleCount,omitempty"`
	PermissionRoles                                        []User_Permission_Role                                                  `json:"permissionRoles,omitempty"`
	PortableStorageVolumeCount                             *uint                                                                   `json:"portableStorageVolumeCount,omitempty"`
	PortableStorageVolumes                                 []Virtual_Disk_Image                                                    `json:"portableStorageVolumes,omitempty"`
	PostProvisioningHookCount                              *uint                                                                   `json:"postProvisioningHookCount,omitempty"`
	PostProvisioningHooks                                  []Provisioning_Hook                                                     `json:"postProvisioningHooks,omitempty"`
	PostalCode                                             *string                                                                 `json:"postalCode,omitempty"`
	PptpVpnUserCount                                       *uint                                                                   `json:"pptpVpnUserCount,omitempty"`
	PptpVpnUsers                                           []User_Customer                                                         `json:"pptpVpnUsers,omitempty"`
	PreviousRecurringRevenue                               *float64                                                                `json:"previousRecurringRevenue,omitempty"`
	PriceRestrictionCount                                  *uint                                                                   `json:"priceRestrictionCount,omitempty"`
	PriceRestrictions                                      []Product_Item_Price_Account_Restriction                                `json:"priceRestrictions,omitempty"`
	PriorityOneTicketCount                                 *uint                                                                   `json:"priorityOneTicketCount,omitempty"`
	PriorityOneTickets                                     []Ticket                                                                `json:"priorityOneTickets,omitempty"`
	PrivateAllotmentHardwareBandwidthDetailCount           *uint                                                                   `json:"privateAllotmentHardwareBandwidthDetailCount,omitempty"`
	PrivateAllotmentHardwareBandwidthDetails               []Network_Bandwidth_Version1_Allotment                                  `json:"privateAllotmentHardwareBandwidthDetails,omitempty"`
	PrivateBlockDeviceTemplateGroupCount                   *uint                                                                   `json:"privateBlockDeviceTemplateGroupCount,omitempty"`
	PrivateBlockDeviceTemplateGroups                       []Virtual_Guest_Block_Device_Template_Group                             `json:"privateBlockDeviceTemplateGroups,omitempty"`
	PrivateIpAddressCount                                  *uint                                                                   `json:"privateIpAddressCount,omitempty"`
	PrivateIpAddresses                                     []Network_Subnet_IpAddress                                              `json:"privateIpAddresses,omitempty"`
	PrivateNetworkVlanCount                                *uint                                                                   `json:"privateNetworkVlanCount,omitempty"`
	PrivateNetworkVlans                                    []Network_Vlan                                                          `json:"privateNetworkVlans,omitempty"`
	PrivateSubnetCount                                     *uint                                                                   `json:"privateSubnetCount,omitempty"`
	PrivateSubnets                                         []Network_Subnet                                                        `json:"privateSubnets,omitempty"`
	PublicAllotmentHardwareBandwidthDetailCount            *uint                                                                   `json:"publicAllotmentHardwareBandwidthDetailCount,omitempty"`
	PublicAllotmentHardwareBandwidthDetails                []Network_Bandwidth_Version1_Allotment                                  `json:"publicAllotmentHardwareBandwidthDetails,omitempty"`
	PublicIpAddressCount                                   *uint                                                                   `json:"publicIpAddressCount,omitempty"`
	PublicIpAddresses                                      []Network_Subnet_IpAddress                                              `json:"publicIpAddresses,omitempty"`
	PublicNetworkVlanCount                                 *uint                                                                   `json:"publicNetworkVlanCount,omitempty"`
	PublicNetworkVlans                                     []Network_Vlan                                                          `json:"publicNetworkVlans,omitempty"`
	PublicSubnetCount                                      *uint                                                                   `json:"publicSubnetCount,omitempty"`
	PublicSubnets                                          []Network_Subnet                                                        `json:"publicSubnets,omitempty"`
	QuoteCount                                             *uint                                                                   `json:"quoteCount,omitempty"`
	Quotes                                                 []Billing_Order_Quote                                                   `json:"quotes,omitempty"`
	RecentEventCount                                       *uint                                                                   `json:"recentEventCount,omitempty"`
	RecentEvents                                           []Notification_Occurrence_Event                                         `json:"recentEvents,omitempty"`
	ReferralPartner                                        *Account                                                                `json:"referralPartner,omitempty"`
	ReferredAccountCount                                   *uint                                                                   `json:"referredAccountCount,omitempty"`
	ReferredAccounts                                       []Account                                                               `json:"referredAccounts,omitempty"`
	RegulatedWorkloadCount                                 *uint                                                                   `json:"regulatedWorkloadCount,omitempty"`
	RegulatedWorkloads                                     []Legal_RegulatedWorkload                                               `json:"regulatedWorkloads,omitempty"`
	RemoteManagementCommandRequestCount                    *uint                                                                   `json:"remoteManagementCommandRequestCount,omitempty"`
	RemoteManagementCommandRequests                        []Hardware_Component_RemoteManagement_Command_Request                   `json:"remoteManagementCommandRequests,omitempty"`
	ReplicationEventCount                                  *uint                                                                   `json:"replicationEventCount,omitempty"`
	ReplicationEvents                                      []Network_Storage_Event                                                 `json:"replicationEvents,omitempty"`
	ResourceGroupCount                                     *uint                                                                   `json:"resourceGroupCount,omitempty"`
	ResourceGroups                                         []Resource_Group                                                        `json:"resourceGroups,omitempty"`
	RouterCount                                            *uint                                                                   `json:"routerCount,omitempty"`
	Routers                                                []Hardware                                                              `json:"routers,omitempty"`
	RwhoisData                                             *Network_Subnet_Rwhois_Data                                             `json:"rwhoisData,omitempty"`
	SalesforceAccountLink                                  *Account_Link                                                           `json:"salesforceAccountLink,omitempty"`
	SamlAuthentication                                     *Account_Authentication_Saml                                            `json:"samlAuthentication,omitempty"`
	ScaleGroupCount                                        *uint                                                                   `json:"scaleGroupCount,omitempty"`
	ScaleGroups                                            []Scale_Group                                                           `json:"scaleGroups,omitempty"`
	SecondaryDomainCount                                   *uint                                                                   `json:"secondaryDomainCount,omitempty"`
	SecondaryDomains                                       []Dns_Secondary                                                         `json:"secondaryDomains,omitempty"`
	SecurityCertificateCount                               *uint                                                                   `json:"securityCertificateCount,omitempty"`
	SecurityCertificates                                   []Security_Certificate                                                  `json:"securityCertificates,omitempty"`
	SecurityScanRequestCount                               *uint                                                                   `json:"securityScanRequestCount,omitempty"`
	SecurityScanRequests                                   []Network_Security_Scanner_Request                                      `json:"securityScanRequests,omitempty"`
	ServiceBillingItemCount                                *uint                                                                   `json:"serviceBillingItemCount,omitempty"`
	ServiceBillingItems                                    []Billing_Item                                                          `json:"serviceBillingItems,omitempty"`
	ShipmentCount                                          *uint                                                                   `json:"shipmentCount,omitempty"`
	Shipments                                              []Account_Shipment                                                      `json:"shipments,omitempty"`
	SshKeyCount                                            *uint                                                                   `json:"sshKeyCount,omitempty"`
	SshKeys                                                []Security_Ssh_Key                                                      `json:"sshKeys,omitempty"`
	SslVpnUserCount                                        *uint                                                                   `json:"sslVpnUserCount,omitempty"`
	SslVpnUsers                                            []User_Customer                                                         `json:"sslVpnUsers,omitempty"`
	StandardPoolVirtualGuestCount                          *uint                                                                   `json:"standardPoolVirtualGuestCount,omitempty"`
	StandardPoolVirtualGuests                              []Virtual_Guest                                                         `json:"standardPoolVirtualGuests,omitempty"`
	State                                                  *string                                                                 `json:"state,omitempty"`
	StatusDate                                             *time.Time                                                              `json:"statusDate,omitempty"`
	SubnetCount                                            *uint                                                                   `json:"subnetCount,omitempty"`
	SubnetRegistrationCount                                *uint                                                                   `json:"subnetRegistrationCount,omitempty"`
	SubnetRegistrationDetailCount                          *uint                                                                   `json:"subnetRegistrationDetailCount,omitempty"`
	SubnetRegistrationDetails                              []Account_Regional_Registry_Detail                                      `json:"subnetRegistrationDetails,omitempty"`
	SubnetRegistrations                                    []Network_Subnet_Registration                                           `json:"subnetRegistrations,omitempty"`
	Subnets                                                []Network_Subnet                                                        `json:"subnets,omitempty"`
	SupportRepresentativeCount                             *uint                                                                   `json:"supportRepresentativeCount,omitempty"`
	SupportRepresentatives                                 []User_Employee                                                         `json:"supportRepresentatives,omitempty"`
	SupportSubscriptionCount                               *uint                                                                   `json:"supportSubscriptionCount,omitempty"`
	SupportSubscriptions                                   []Billing_Item                                                          `json:"supportSubscriptions,omitempty"`
	SupportTier                                            *string                                                                 `json:"supportTier,omitempty"`
	SuppressInvoicesFlag                                   *bool                                                                   `json:"suppressInvoicesFlag,omitempty"`
	TagCount                                               *uint                                                                   `json:"tagCount,omitempty"`
	Tags                                                   []Tag                                                                   `json:"tags,omitempty"`
	TicketCount                                            *uint                                                                   `json:"ticketCount,omitempty"`
	Tickets                                                []Ticket                                                                `json:"tickets,omitempty"`
	TicketsClosedInTheLastThreeDays                        []Ticket                                                                `json:"ticketsClosedInTheLastThreeDays,omitempty"`
	TicketsClosedInTheLastThreeDaysCount                   *uint                                                                   `json:"ticketsClosedInTheLastThreeDaysCount,omitempty"`
	TicketsClosedToday                                     []Ticket                                                                `json:"ticketsClosedToday,omitempty"`
	TicketsClosedTodayCount                                *uint                                                                   `json:"ticketsClosedTodayCount,omitempty"`
	TranscodeAccountCount                                  *uint                                                                   `json:"transcodeAccountCount,omitempty"`
	TranscodeAccounts                                      []Network_Media_Transcode_Account                                       `json:"transcodeAccounts,omitempty"`
	UpgradeRequestCount                                    *uint                                                                   `json:"upgradeRequestCount,omitempty"`
	UpgradeRequests                                        []Product_Upgrade_Request                                               `json:"upgradeRequests,omitempty"`
	UserCount                                              *uint                                                                   `json:"userCount,omitempty"`
	Users                                                  []User_Customer                                                         `json:"users,omitempty"`
	ValidSecurityCertificateCount                          *uint                                                                   `json:"validSecurityCertificateCount,omitempty"`
	ValidSecurityCertificates                              []Security_Certificate                                                  `json:"validSecurityCertificates,omitempty"`
	VdrUpdatesInProgressFlag                               *bool                                                                   `json:"vdrUpdatesInProgressFlag,omitempty"`
	VirtualDedicatedRackCount                              *uint                                                                   `json:"virtualDedicatedRackCount,omitempty"`
	VirtualDedicatedRacks                                  []Network_Bandwidth_Version1_Allotment                                  `json:"virtualDedicatedRacks,omitempty"`
	VirtualDiskImageCount                                  *uint                                                                   `json:"virtualDiskImageCount,omitempty"`
	VirtualDiskImages                                      []Virtual_Disk_Image                                                    `json:"virtualDiskImages,omitempty"`
	VirtualGuestCount                                      *uint                                                                   `json:"virtualGuestCount,omitempty"`
	VirtualGuests                                          []Virtual_Guest                                                         `json:"virtualGuests,omitempty"`
	VirtualGuestsOverBandwidthAllocation                   []Virtual_Guest                                                         `json:"virtualGuestsOverBandwidthAllocation,omitempty"`
	VirtualGuestsOverBandwidthAllocationCount              *uint                                                                   `json:"virtualGuestsOverBandwidthAllocationCount,omitempty"`
	VirtualGuestsProjectedOverBandwidthAllocation          []Virtual_Guest                                                         `json:"virtualGuestsProjectedOverBandwidthAllocation,omitempty"`
	VirtualGuestsProjectedOverBandwidthAllocationCount     *uint                                                                   `json:"virtualGuestsProjectedOverBandwidthAllocationCount,omitempty"`
	VirtualGuestsWithCpanel                                []Virtual_Guest                                                         `json:"virtualGuestsWithCpanel,omitempty"`
	VirtualGuestsWithCpanelCount                           *uint                                                                   `json:"virtualGuestsWithCpanelCount,omitempty"`
	VirtualGuestsWithMcafee                                []Virtual_Guest                                                         `json:"virtualGuestsWithMcafee,omitempty"`
	VirtualGuestsWithMcafeeAntivirusRedhat                 []Virtual_Guest                                                         `json:"virtualGuestsWithMcafeeAntivirusRedhat,omitempty"`
	VirtualGuestsWithMcafeeAntivirusRedhatCount            *uint                                                                   `json:"virtualGuestsWithMcafeeAntivirusRedhatCount,omitempty"`
	VirtualGuestsWithMcafeeAntivirusWindowCount            *uint                                                                   `json:"virtualGuestsWithMcafeeAntivirusWindowCount,omitempty"`
	VirtualGuestsWithMcafeeAntivirusWindows                []Virtual_Guest                                                         `json:"virtualGuestsWithMcafeeAntivirusWindows,omitempty"`
	VirtualGuestsWithMcafeeCount                           *uint                                                                   `json:"virtualGuestsWithMcafeeCount,omitempty"`
	VirtualGuestsWithMcafeeIntrusionDetectionSystem        []Virtual_Guest                                                         `json:"virtualGuestsWithMcafeeIntrusionDetectionSystem,omitempty"`
	VirtualGuestsWithMcafeeIntrusionDetectionSystemCount   *uint                                                                   `json:"virtualGuestsWithMcafeeIntrusionDetectionSystemCount,omitempty"`
	VirtualGuestsWithPlesk                                 []Virtual_Guest                                                         `json:"virtualGuestsWithPlesk,omitempty"`
	VirtualGuestsWithPleskCount                            *uint                                                                   `json:"virtualGuestsWithPleskCount,omitempty"`
	VirtualGuestsWithQuantastor                            []Virtual_Guest                                                         `json:"virtualGuestsWithQuantastor,omitempty"`
	VirtualGuestsWithQuantastorCount                       *uint                                                                   `json:"virtualGuestsWithQuantastorCount,omitempty"`
	VirtualGuestsWithUrchin                                []Virtual_Guest                                                         `json:"virtualGuestsWithUrchin,omitempty"`
	VirtualGuestsWithUrchinCount                           *uint                                                                   `json:"virtualGuestsWithUrchinCount,omitempty"`
	VirtualPrivateRack                                     *Network_Bandwidth_Version1_Allotment                                   `json:"virtualPrivateRack,omitempty"`
	VirtualStorageArchiveRepositories                      []Virtual_Storage_Repository                                            `json:"virtualStorageArchiveRepositories,omitempty"`
	VirtualStorageArchiveRepositoryCount                   *uint                                                                   `json:"virtualStorageArchiveRepositoryCount,omitempty"`
	VirtualStoragePublicRepositories                       []Virtual_Storage_Repository                                            `json:"virtualStoragePublicRepositories,omitempty"`
	VirtualStoragePublicRepositoryCount                    *uint                                                                   `json:"virtualStoragePublicRepositoryCount,omitempty"`
}

type Account_AbuseEmail struct {
	Entity

	Account *Account `json:"account,omitempty"`
	Email   *string  `json:"email,omitempty"`
}

type Account_Address struct {
	Entity

	Account        *Account              `json:"account,omitempty"`
	AccountId      *int                  `json:"accountId,omitempty"`
	Address1       *string               `json:"address1,omitempty"`
	Address2       *string               `json:"address2,omitempty"`
	City           *string               `json:"city,omitempty"`
	ContactName    *string               `json:"contactName,omitempty"`
	Country        *string               `json:"country,omitempty"`
	CreateUser     *User_Customer        `json:"createUser,omitempty"`
	Description    *string               `json:"description,omitempty"`
	Id             *int                  `json:"id,omitempty"`
	IsActive       *int                  `json:"isActive,omitempty"`
	Location       *Location             `json:"location,omitempty"`
	LocationId     *int                  `json:"locationId,omitempty"`
	ModifyEmployee *User_Employee        `json:"modifyEmployee,omitempty"`
	ModifyUser     *User_Customer        `json:"modifyUser,omitempty"`
	PostalCode     *string               `json:"postalCode,omitempty"`
	State          *string               `json:"state,omitempty"`
	Type           *Account_Address_Type `json:"type,omitempty"`
}

type Account_Address_Type struct {
	Entity

	CreateDate *time.Time `json:"createDate,omitempty"`
	Id         *int       `json:"id,omitempty"`
	KeyName    *string    `json:"keyName,omitempty"`
	Name       *string    `json:"name,omitempty"`
}

type Account_Affiliation struct {
	Entity

	Account     *Account   `json:"account,omitempty"`
	AccountId   *int       `json:"accountId,omitempty"`
	AffiliateId *string    `json:"affiliateId,omitempty"`
	CreateDate  *time.Time `json:"createDate,omitempty"`
	Id          *int       `json:"id,omitempty"`
	ModifyDate  *time.Time `json:"modifyDate,omitempty"`
}

type Account_Agreement struct {
	Entity

	Account                           *Account                         `json:"account,omitempty"`
	AgreementType                     *Account_Agreement_Type          `json:"agreementType,omitempty"`
	AgreementTypeId                   *int                             `json:"agreementTypeId,omitempty"`
	AttachedBillingAgreementFileCount *uint                            `json:"attachedBillingAgreementFileCount,omitempty"`
	AttachedBillingAgreementFiles     []Account_MasterServiceAgreement `json:"attachedBillingAgreementFiles,omitempty"`
	AutoRenew                         *int                             `json:"autoRenew,omitempty"`
	BillingItemCount                  *uint                            `json:"billingItemCount,omitempty"`
	BillingItems                      []Billing_Item                   `json:"billingItems,omitempty"`
	CancellationFee                   *int                             `json:"cancellationFee,omitempty"`
	CreateDate                        *time.Time                       `json:"createDate,omitempty"`
	DurationMonths                    *int                             `json:"durationMonths,omitempty"`
	EndDate                           *time.Time                       `json:"endDate,omitempty"`
	Id                                *int                             `json:"id,omitempty"`
	StartDate                         *time.Time                       `json:"startDate,omitempty"`
	Status                            *Account_Agreement_Status        `json:"status,omitempty"`
	StatusId                          *int                             `json:"statusId,omitempty"`
	Title                             *string                          `json:"title,omitempty"`
	TopLevelBillingItemCount          *uint                            `json:"topLevelBillingItemCount,omitempty"`
	TopLevelBillingItems              []Billing_Item                   `json:"topLevelBillingItems,omitempty"`
}

type Account_Agreement_Status struct {
	Entity

	Name *string `json:"name,omitempty"`
}

type Account_Agreement_Type struct {
	Entity

	Name *string `json:"name,omitempty"`
}

type Account_Attachment_Employee struct {
	Entity

	Account      *Account                          `json:"account,omitempty"`
	Employee     *User_Employee                    `json:"employee,omitempty"`
	EmployeeRole *Account_Attachment_Employee_Role `json:"employeeRole,omitempty"`
	RoleId       *int                              `json:"roleId,omitempty"`
}

type Account_Attachment_Employee_Role struct {
	Entity

	Keyname *string `json:"keyname,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type Account_Attribute struct {
	Entity

	Account                *Account                `json:"account,omitempty"`
	AccountAttributeType   *Account_Attribute_Type `json:"accountAttributeType,omitempty"`
	AccountAttributeTypeId *int                    `json:"accountAttributeTypeId,omitempty"`
	AccountId              *int                    `json:"accountId,omitempty"`
	Id                     *int                    `json:"id,omitempty"`
	Value                  *string                 `json:"value,omitempty"`
}

type Account_Attribute_Type struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	KeyName     *string `json:"keyName,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Account_Authentication_Attribute struct {
	Entity

	Account              *Account                               `json:"account,omitempty"`
	AccountId            *int                                   `json:"accountId,omitempty"`
	AuthenticationRecord *Account_Authentication_Saml           `json:"authenticationRecord,omitempty"`
	Id                   *int                                   `json:"id,omitempty"`
	Type                 *Account_Authentication_Attribute_Type `json:"type,omitempty"`
	TypeId               *int                                   `json:"typeId,omitempty"`
	Value                *string                                `json:"value,omitempty"`
}

type Account_Authentication_Attribute_Type struct {
	Entity

	Description  *string `json:"description,omitempty"`
	Id           *int    `json:"id,omitempty"`
	KeyName      *string `json:"keyName,omitempty"`
	Name         *string `json:"name,omitempty"`
	ValueExample *string `json:"valueExample,omitempty"`
}

type Account_Authentication_OpenIdConnect_Option struct {
	Entity

	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

type Account_Authentication_OpenIdConnect_RegistrationInformation struct {
	Entity

	ExistingBlueIdFlag       *bool          `json:"existingBlueIdFlag,omitempty"`
	FederatedEmailDomainFlag *bool          `json:"federatedEmailDomainFlag,omitempty"`
	User                     *User_Customer `json:"user,omitempty"`
}

type Account_Authentication_Saml struct {
	Entity

	Account                             *Account                           `json:"account,omitempty"`
	AccountId                           *string                            `json:"accountId,omitempty"`
	AttributeCount                      *uint                              `json:"attributeCount,omitempty"`
	Attributes                          []Account_Authentication_Attribute `json:"attributes,omitempty"`
	Certificate                         *string                            `json:"certificate,omitempty"`
	CertificateFingerprint              *string                            `json:"certificateFingerprint,omitempty"`
	EntityId                            *string                            `json:"entityId,omitempty"`
	Id                                  *int                               `json:"id,omitempty"`
	ServiceProviderCertificate          *string                            `json:"serviceProviderCertificate,omitempty"`
	ServiceProviderEntityId             *string                            `json:"serviceProviderEntityId,omitempty"`
	ServiceProviderPublicKey            *string                            `json:"serviceProviderPublicKey,omitempty"`
	ServiceProviderSingleLogoutEncoding *string                            `json:"serviceProviderSingleLogoutEncoding,omitempty"`
	ServiceProviderSingleLogoutUrl      *string                            `json:"serviceProviderSingleLogoutUrl,omitempty"`
	ServiceProviderSingleSignOnEncoding *string                            `json:"serviceProviderSingleSignOnEncoding,omitempty"`
	ServiceProviderSingleSignOnUrl      *string                            `json:"serviceProviderSingleSignOnUrl,omitempty"`
	SingleLogoutEncoding                *string                            `json:"singleLogoutEncoding,omitempty"`
	SingleLogoutUrl                     *string                            `json:"singleLogoutUrl,omitempty"`
	SingleSignOnEncoding                *string                            `json:"singleSignOnEncoding,omitempty"`
	SingleSignOnUrl                     *string                            `json:"singleSignOnUrl,omitempty"`
}

type Account_Classification_Group_Type struct {
	Entity

	KeyName *string `json:"keyName,omitempty"`
}

type Account_Contact struct {
	Entity

	Account        *Account              `json:"account,omitempty"`
	AccountId      *int                  `json:"accountId,omitempty"`
	Address1       *string               `json:"address1,omitempty"`
	Address2       *string               `json:"address2,omitempty"`
	AlternatePhone *string               `json:"alternatePhone,omitempty"`
	City           *string               `json:"city,omitempty"`
	CompanyName    *string               `json:"companyName,omitempty"`
	Country        *string               `json:"country,omitempty"`
	CreateDate     *time.Time            `json:"createDate,omitempty"`
	Email          *string               `json:"email,omitempty"`
	FaxPhone       *string               `json:"faxPhone,omitempty"`
	FirstName      *string               `json:"firstName,omitempty"`
	Id             *int                  `json:"id,omitempty"`
	JobTitle       *string               `json:"jobTitle,omitempty"`
	LastName       *string               `json:"lastName,omitempty"`
	ModifyDate     *time.Time            `json:"modifyDate,omitempty"`
	OfficePhone    *string               `json:"officePhone,omitempty"`
	PostalCode     *string               `json:"postalCode,omitempty"`
	ProfileName    *string               `json:"profileName,omitempty"`
	State          *string               `json:"state,omitempty"`
	Type           *Account_Contact_Type `json:"type,omitempty"`
	TypeId         *int                  `json:"typeId,omitempty"`
	Url            *string               `json:"url,omitempty"`
}

type Account_Contact_Type struct {
	Entity

	CreateDate  *time.Time `json:"createDate,omitempty"`
	Description *string    `json:"description,omitempty"`
	Id          *int       `json:"id,omitempty"`
	KeyName     *string    `json:"keyName,omitempty"`
	ModifyDate  *time.Time `json:"modifyDate,omitempty"`
	Name        *string    `json:"name,omitempty"`
}

type Account_Historical_Report struct {
	Entity
}

type Account_Link struct {
	Entity

	Account                          *Account          `json:"account,omitempty"`
	AccountId                        *int              `json:"accountId,omitempty"`
	CreateDate                       *time.Time        `json:"createDate,omitempty"`
	DestinationAccountAlphanumericId *string           `json:"destinationAccountAlphanumericId,omitempty"`
	DestinationAccountId             *int              `json:"destinationAccountId,omitempty"`
	Id                               *int              `json:"id,omitempty"`
	ServiceProvider                  *Service_Provider `json:"serviceProvider,omitempty"`
	ServiceProviderId                *int              `json:"serviceProviderId,omitempty"`
}

type Account_Link_Bluemix struct {
	Account_Link
}

type Account_Link_OpenStack struct {
	Account_Link

	DomainId *string `json:"domainId,omitempty"`
}

type Account_Link_OpenStack_DomainCreationDetails struct {
	Entity

	DomainId *string `json:"domainId,omitempty"`
	UserId   *string `json:"userId,omitempty"`
	UserName *string `json:"userName,omitempty"`
}

type Account_Link_OpenStack_LinkRequest struct {
	Entity

	DesiredPassword    *string `json:"desiredPassword,omitempty"`
	DesiredProjectName *string `json:"desiredProjectName,omitempty"`
	DesiredUsername    *string `json:"desiredUsername,omitempty"`
}

type Account_Link_OpenStack_ProjectCreationDetails struct {
	Entity

	DomainId    *string `json:"domainId,omitempty"`
	ProjectId   *string `json:"projectId,omitempty"`
	ProjectName *string `json:"projectName,omitempty"`
	UserId      *string `json:"userId,omitempty"`
	UserName    *string `json:"userName,omitempty"`
}

type Account_Link_OpenStack_ProjectDetails struct {
	Entity

	ProjectId   *string `json:"projectId,omitempty"`
	ProjectName *string `json:"projectName,omitempty"`
}

type Account_Link_ThePlanet struct {
	Account_Link
}

type Account_Link_Vendor struct {
	Entity

	KeyName *string `json:"keyName,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type Account_Lockdown_Request struct {
	Entity

	AccountId  *int       `json:"accountId,omitempty"`
	Action     *string    `json:"action,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	Id         *int       `json:"id,omitempty"`
	ModifyDate *time.Time `json:"modifyDate,omitempty"`
	Status     *string    `json:"status,omitempty"`
}

type Account_MasterServiceAgreement struct {
	Entity

	Account   *Account `json:"account,omitempty"`
	AccountId *int     `json:"accountId,omitempty"`
	Guid      *string  `json:"guid,omitempty"`
	Id        *int     `json:"id,omitempty"`
	Name      *string  `json:"name,omitempty"`
}

type Account_Media struct {
	Entity

	Account        *Account                             `json:"account,omitempty"`
	CreateUser     *User_Customer                       `json:"createUser,omitempty"`
	Datacenter     *Location                            `json:"datacenter,omitempty"`
	Description    *string                              `json:"description,omitempty"`
	Id             *int                                 `json:"id,omitempty"`
	ModifyEmployee *User_Employee                       `json:"modifyEmployee,omitempty"`
	ModifyUser     *User_Customer                       `json:"modifyUser,omitempty"`
	Request        *Account_Media_Data_Transfer_Request `json:"request,omitempty"`
	RequestId      *int                                 `json:"requestId,omitempty"`
	SerialNumber   *string                              `json:"serialNumber,omitempty"`
	Type           *Account_Media_Type                  `json:"type,omitempty"`
	TypeId         *int                                 `json:"typeId,omitempty"`
	Volume         *Network_Storage                     `json:"volume,omitempty"`
}

type Account_Media_Data_Transfer_Request struct {
	Entity

	Account           *Account                                    `json:"account,omitempty"`
	AccountId         *int                                        `json:"accountId,omitempty"`
	ActiveTicketCount *uint                                       `json:"activeTicketCount,omitempty"`
	ActiveTickets     []Ticket                                    `json:"activeTickets,omitempty"`
	BillingItem       *Billing_Item                               `json:"billingItem,omitempty"`
	CreateUser        *User_Customer                              `json:"createUser,omitempty"`
	CreateUserId      *int                                        `json:"createUserId,omitempty"`
	EndDate           *time.Time                                  `json:"endDate,omitempty"`
	Id                *int                                        `json:"id,omitempty"`
	Media             *Account_Media                              `json:"media,omitempty"`
	ModifyEmployee    *User_Employee                              `json:"modifyEmployee,omitempty"`
	ModifyUser        *User_Customer                              `json:"modifyUser,omitempty"`
	ModifyUserId      *int                                        `json:"modifyUserId,omitempty"`
	ShipmentCount     *uint                                       `json:"shipmentCount,omitempty"`
	Shipments         []Account_Shipment                          `json:"shipments,omitempty"`
	StartDate         *time.Time                                  `json:"startDate,omitempty"`
	Status            *Account_Media_Data_Transfer_Request_Status `json:"status,omitempty"`
	StatusId          *int                                        `json:"statusId,omitempty"`
	TicketCount       *uint                                       `json:"ticketCount,omitempty"`
	Tickets           []Ticket                                    `json:"tickets,omitempty"`
}

type Account_Media_Data_Transfer_Request_Status struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	KeyName     *string `json:"keyName,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Account_Media_Type struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	KeyName     *string `json:"keyName,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Account_Network_Vlan_Span struct {
	Entity

	Account          *Account   `json:"account,omitempty"`
	EnabledFlag      *bool      `json:"enabledFlag,omitempty"`
	Id               *int       `json:"id,omitempty"`
	LastAppliedDate  *time.Time `json:"lastAppliedDate,omitempty"`
	LastVerifiedDate *time.Time `json:"lastVerifiedDate,omitempty"`
	ModifyDate       *time.Time `json:"modifyDate,omitempty"`
}

type Account_Note struct {
	Entity

	Account          *Account               `json:"account,omitempty"`
	AccountId        *int                   `json:"accountId,omitempty"`
	CreateDate       *time.Time             `json:"createDate,omitempty"`
	Customer         *User_Customer         `json:"customer,omitempty"`
	Id               *int                   `json:"id,omitempty"`
	ModifyDate       *time.Time             `json:"modifyDate,omitempty"`
	Note             *string                `json:"note,omitempty"`
	NoteHistory      []Account_Note_History `json:"noteHistory,omitempty"`
	NoteHistoryCount *uint                  `json:"noteHistoryCount,omitempty"`
	NoteType         *Account_Note_Type     `json:"noteType,omitempty"`
	NoteTypeId       *int                   `json:"noteTypeId,omitempty"`
	UserId           *int                   `json:"userId,omitempty"`
}

type Account_Note_History struct {
	Entity

	AccountNote   *Account_Note  `json:"accountNote,omitempty"`
	AccountNoteId *int           `json:"accountNoteId,omitempty"`
	CreateDate    *time.Time     `json:"createDate,omitempty"`
	Customer      *User_Customer `json:"customer,omitempty"`
	Id            *int           `json:"id,omitempty"`
	ModifyDate    *time.Time     `json:"modifyDate,omitempty"`
	Note          *string        `json:"note,omitempty"`
	UserId        *int           `json:"userId,omitempty"`
}

type Account_Note_Type struct {
	Entity

	BrandId         *int       `json:"brandId,omitempty"`
	CreateDate      *time.Time `json:"createDate,omitempty"`
	Description     *string    `json:"description,omitempty"`
	Id              *int       `json:"id,omitempty"`
	KeyName         *string    `json:"keyName,omitempty"`
	ModifyDate      *time.Time `json:"modifyDate,omitempty"`
	Name            *string    `json:"name,omitempty"`
	ValueExpression *string    `json:"valueExpression,omitempty"`
}

type Account_Partner_Referral_Prospect struct {
	User_Customer_Prospect

	CompanyName  *string `json:"companyName,omitempty"`
	EmailAddress *string `json:"emailAddress,omitempty"`
	FirstName    *string `json:"firstName,omitempty"`
	Id           *int    `json:"id,omitempty"`
	LastName     *string `json:"lastName,omitempty"`
}

type Account_Password struct {
	Entity

	Account   *Account               `json:"account,omitempty"`
	AccountId *int                   `json:"accountId,omitempty"`
	Id        *int                   `json:"id,omitempty"`
	Notes     *string                `json:"notes,omitempty"`
	Password  *string                `json:"password,omitempty"`
	Type      *Account_Password_Type `json:"type,omitempty"`
	TypeId    *int                   `json:"typeId,omitempty"`
	Username  *string                `json:"username,omitempty"`
}

type Account_Password_Type struct {
	Entity

	Description *string `json:"description,omitempty"`
}

type Account_Regional_Registry_Detail struct {
	Entity

	Account                          *Account                                    `json:"account,omitempty"`
	AccountId                        *int                                        `json:"accountId,omitempty"`
	CreateDate                       *time.Time                                  `json:"createDate,omitempty"`
	DetailCount                      *uint                                       `json:"detailCount,omitempty"`
	DetailType                       *Account_Regional_Registry_Detail_Type      `json:"detailType,omitempty"`
	DetailTypeId                     *int                                        `json:"detailTypeId,omitempty"`
	Details                          []Network_Subnet_Registration_Details       `json:"details,omitempty"`
	Id                               *int                                        `json:"id,omitempty"`
	ModifyDate                       *time.Time                                  `json:"modifyDate,omitempty"`
	Properties                       []Account_Regional_Registry_Detail_Property `json:"properties,omitempty"`
	PropertyCount                    *uint                                       `json:"propertyCount,omitempty"`
	RegionalInternetRegistryHandle   *Account_Rwhois_Handle                      `json:"regionalInternetRegistryHandle,omitempty"`
	RegionalInternetRegistryHandleId *int                                        `json:"regionalInternetRegistryHandleId,omitempty"`
}

type Account_Regional_Registry_Detail_Property struct {
	Entity

	CreateDate           *time.Time                                      `json:"createDate,omitempty"`
	Detail               *Account_Regional_Registry_Detail               `json:"detail,omitempty"`
	Id                   *int                                            `json:"id,omitempty"`
	ModifyDate           *time.Time                                      `json:"modifyDate,omitempty"`
	PropertyType         *Account_Regional_Registry_Detail_Property_Type `json:"propertyType,omitempty"`
	PropertyTypeId       *int                                            `json:"propertyTypeId,omitempty"`
	RegistrationDetailId *int                                            `json:"registrationDetailId,omitempty"`
	SequencePosition     *int                                            `json:"sequencePosition,omitempty"`
	Value                *string                                         `json:"value,omitempty"`
}

type Account_Regional_Registry_Detail_Property_Type struct {
	Entity

	CreateDate      *time.Time `json:"createDate,omitempty"`
	Id              *int       `json:"id,omitempty"`
	KeyName         *string    `json:"keyName,omitempty"`
	ModifyDate      *time.Time `json:"modifyDate,omitempty"`
	Name            *string    `json:"name,omitempty"`
	ValueExpression *string    `json:"valueExpression,omitempty"`
}

type Account_Regional_Registry_Detail_Type struct {
	Entity

	CreateDate *time.Time `json:"createDate,omitempty"`
	Id         *int       `json:"id,omitempty"`
	KeyName    *string    `json:"keyName,omitempty"`
	ModifyDate *time.Time `json:"modifyDate,omitempty"`
	Name       *string    `json:"name,omitempty"`
}

type Account_Regional_Registry_Detail_Version4_Person_Default struct {
	Account_Regional_Registry_Detail
}

type Account_Reports_Request struct {
	Entity

	Account                *Account                `json:"account,omitempty"`
	AccountContact         *Account_Contact        `json:"accountContact,omitempty"`
	AccountContactId       *int                    `json:"accountContactId,omitempty"`
	AccountId              *int                    `json:"accountId,omitempty"`
	ComplianceReportTypeId *string                 `json:"complianceReportTypeId,omitempty"`
	CreateDate             *time.Time              `json:"createDate,omitempty"`
	EmployeeRecordId       *int                    `json:"employeeRecordId,omitempty"`
	Id                     *int                    `json:"id,omitempty"`
	ModifyDate             *time.Time              `json:"modifyDate,omitempty"`
	Nda                    *string                 `json:"nda,omitempty"`
	Notes                  *string                 `json:"notes,omitempty"`
	Report                 *string                 `json:"report,omitempty"`
	ReportType             *Compliance_Report_Type `json:"reportType,omitempty"`
	RequestKey             *string                 `json:"requestKey,omitempty"`
	Status                 *string                 `json:"status,omitempty"`
	Ticket                 *Ticket                 `json:"ticket,omitempty"`
	TicketId               *int                    `json:"ticketId,omitempty"`
	User                   *User_Customer          `json:"user,omitempty"`
	UsrRecordId            *int                    `json:"usrRecordId,omitempty"`
}

type Account_Rwhois_Handle struct {
	Entity

	Account    *Account   `json:"account,omitempty"`
	AccountId  *int       `json:"accountId,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	Handle     *string    `json:"handle,omitempty"`
	Id         *int       `json:"id,omitempty"`
	ModifyDate *time.Time `json:"modifyDate,omitempty"`
}

type Account_Shipment struct {
	Entity

	Account              *Account                         `json:"account,omitempty"`
	AccountId            *int                             `json:"accountId,omitempty"`
	Courier              *Auxiliary_Shipping_Courier      `json:"courier,omitempty"`
	CourierId            *int                             `json:"courierId,omitempty"`
	CourierName          *string                          `json:"courierName,omitempty"`
	CreateEmployee       *User_Employee                   `json:"createEmployee,omitempty"`
	CreateUser           *User_Customer                   `json:"createUser,omitempty"`
	CreateUserId         *int                             `json:"createUserId,omitempty"`
	DestinationAddress   *Account_Address                 `json:"destinationAddress,omitempty"`
	DestinationAddressId *int                             `json:"destinationAddressId,omitempty"`
	DestinationDate      *time.Time                       `json:"destinationDate,omitempty"`
	Id                   *int                             `json:"id,omitempty"`
	ModifyEmployee       *User_Employee                   `json:"modifyEmployee,omitempty"`
	ModifyUser           *User_Customer                   `json:"modifyUser,omitempty"`
	ModifyUserId         *int                             `json:"modifyUserId,omitempty"`
	Note                 *string                          `json:"note,omitempty"`
	OriginationAddress   *Account_Address                 `json:"originationAddress,omitempty"`
	OriginationAddressId *int                             `json:"originationAddressId,omitempty"`
	OriginationDate      *time.Time                       `json:"originationDate,omitempty"`
	ShipmentItemCount    *uint                            `json:"shipmentItemCount,omitempty"`
	ShipmentItems        []Account_Shipment_Item          `json:"shipmentItems,omitempty"`
	Status               *Account_Shipment_Status         `json:"status,omitempty"`
	StatusId             *int                             `json:"statusId,omitempty"`
	TrackingData         []Account_Shipment_Tracking_Data `json:"trackingData,omitempty"`
	TrackingDataCount    *uint                            `json:"trackingDataCount,omitempty"`
	Type                 *Account_Shipment_Type           `json:"type,omitempty"`
	TypeId               *int                             `json:"typeId,omitempty"`
}

type Account_Shipment_Item struct {
	Entity

	CreateDate         *time.Time                  `json:"createDate,omitempty"`
	Description        *string                     `json:"description,omitempty"`
	Id                 *int                        `json:"id,omitempty"`
	PackageId          *int                        `json:"packageId,omitempty"`
	Shipment           *Account_Shipment           `json:"shipment,omitempty"`
	ShipmentId         *int                        `json:"shipmentId,omitempty"`
	ShipmentItemId     *int                        `json:"shipmentItemId,omitempty"`
	ShipmentItemType   *Account_Shipment_Item_Type `json:"shipmentItemType,omitempty"`
	ShipmentItemTypeId *int                        `json:"shipmentItemTypeId,omitempty"`
}

type Account_Shipment_Item_Type struct {
	Entity

	CreateDate *time.Time `json:"createDate,omitempty"`
	Id         *int       `json:"id,omitempty"`
	KeyName    *string    `json:"keyName,omitempty"`
	Name       *string    `json:"name,omitempty"`
}

type Account_Shipment_Resource_Type struct {
	Entity
}

type Account_Shipment_Status struct {
	Entity

	CreateDate *time.Time `json:"createDate,omitempty"`
	Id         *int       `json:"id,omitempty"`
	KeyName    *string    `json:"keyName,omitempty"`
	Name       *string    `json:"name,omitempty"`
}

type Account_Shipment_Tracking_Data struct {
	Entity

	CreateEmployee *User_Employee    `json:"createEmployee,omitempty"`
	CreateUser     *User_Customer    `json:"createUser,omitempty"`
	CreateUserId   *int              `json:"createUserId,omitempty"`
	Id             *int              `json:"id,omitempty"`
	ModifyEmployee *User_Employee    `json:"modifyEmployee,omitempty"`
	ModifyUser     *User_Customer    `json:"modifyUser,omitempty"`
	ModifyUserId   *int              `json:"modifyUserId,omitempty"`
	PackageId      *int              `json:"packageId,omitempty"`
	Sequence       *int              `json:"sequence,omitempty"`
	Shipment       *Account_Shipment `json:"shipment,omitempty"`
	ShipmentId     *int              `json:"shipmentId,omitempty"`
	TrackingData   *string           `json:"trackingData,omitempty"`
}

type Account_Shipment_Type struct {
	Entity

	CreateDate  *time.Time `json:"createDate,omitempty"`
	Description *string    `json:"description,omitempty"`
	Id          *int       `json:"id,omitempty"`
	KeyName     *string    `json:"keyName,omitempty"`
	Name        *string    `json:"name,omitempty"`
}

type Account_Status struct {
	Entity

	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}
