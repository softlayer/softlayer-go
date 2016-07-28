package datatypes

import (
	"time"
)

type McAfee_Epolicy_Orchestrator_Version36_Agent_Details struct {
	Entity

	AgentVersion  *string                                                     `json:"agentVersion:omitempty"`
	CurrentPolicy *McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details `json:"currentPolicy:omitempty"`
	LastUpdate    *string                                                     `json:"lastUpdate:omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details struct {
	Entity

	CurrentPolicy *McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details `json:"currentPolicy:omitempty"`
	Name          *string                                                     `json:"name:omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event struct {
	Entity

	EventLocalDateTime *time.Time                                                                `json:"eventLocalDateTime:omitempty"`
	Filename           *string                                                                   `json:"filename:omitempty"`
	VirusActionTaken   *McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_Filter_Description `json:"virusActionTaken:omitempty"`
	VirusName          *string                                                                   `json:"virusName:omitempty"`
	VirusType          *string                                                                   `json:"virusType:omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_AccessProtection struct {
	Entity

	EventLocalDateTime *time.Time `json:"eventLocalDateTime:omitempty"`
	Filename           *string    `json:"filename:omitempty"`
	ProcessName        *string    `json:"processName:omitempty"`
	RuleName           *string    `json:"ruleName:omitempty"`
	Source             *string    `json:"source:omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_Filter_Description struct {
	Entity

	Name *string `json:"name:omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_BlockedApplicationEvent struct {
	Entity

	ApplicationDescription *string    `json:"applicationDescription:omitempty"`
	IncidentTime           *time.Time `json:"incidentTime:omitempty"`
	ProcessName            *string    `json:"processName:omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_Event_Signature struct {
	Entity

	SignatureName *string `json:"signatureName:omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_IPSEvent struct {
	Entity

	IncidentTime    *time.Time                                                           `json:"incidentTime:omitempty"`
	ProcessName     *string                                                              `json:"processName:omitempty"`
	ReactionText    *string                                                              `json:"reactionText:omitempty"`
	RemoteIpAddress *string                                                              `json:"remoteIpAddress:omitempty"`
	SeverityText    *string                                                              `json:"severityText:omitempty"`
	Signature       *McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_Event_Signature `json:"signature:omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_BlockedApplicationEvent struct {
	Entity

	ApplicationDescription *string    `json:"applicationDescription:omitempty"`
	IncidentTime           *time.Time `json:"incidentTime:omitempty"`
	ProcessName            *string    `json:"processName:omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_Event_Signature struct {
	Entity

	SignatureName *string `json:"signatureName:omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent struct {
	Entity

	IncidentTime    *time.Time                                                           `json:"incidentTime:omitempty"`
	ProcessName     *string                                                              `json:"processName:omitempty"`
	ReactionText    *string                                                              `json:"reactionText:omitempty"`
	RemoteIpAddress *string                                                              `json:"remoteIpAddress:omitempty"`
	SeverityText    *string                                                              `json:"severityText:omitempty"`
	Signature       *McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_Event_Signature `json:"signature:omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version36_Policy_Object struct {
	Entity

	Name *string `json:"name:omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version36_Product_Properties struct {
	Entity

	DatVersion *string `json:"datVersion:omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version45_Agent_Details struct {
	Entity

	AgentVersion *string    `json:"agentVersion:omitempty"`
	LastUpdate   *time.Time `json:"lastUpdate:omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details struct {
	Entity

	AgentDetails *McAfee_Epolicy_Orchestrator_Version45_Agent_Details         `json:"agentDetails:omitempty"`
	Name         *string                                                      `json:"name:omitempty"`
	Policies     []McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details `json:"policies:omitempty"`
	PolicyCount  *uint                                                        `json:"policyCount:omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version45_Event struct {
	Entity

	AgentDetails        *McAfee_Epolicy_Orchestrator_Version45_Agent_Details            `json:"agentDetails:omitempty"`
	DetectedUtc         *time.Time                                                      `json:"detectedUtc:omitempty"`
	SourceIpv4          *string                                                         `json:"sourceIpv4:omitempty"`
	SourceProcessName   *string                                                         `json:"sourceProcessName:omitempty"`
	TargetFilename      *string                                                         `json:"targetFilename:omitempty"`
	ThreatActionTaken   *string                                                         `json:"threatActionTaken:omitempty"`
	ThreatName          *string                                                         `json:"threatName:omitempty"`
	ThreatSeverityLabel *string                                                         `json:"threatSeverityLabel:omitempty"`
	ThreatType          *string                                                         `json:"threatType:omitempty"`
	VirusActionTaken    *McAfee_Epolicy_Orchestrator_Version45_Event_Filter_Description `json:"virusActionTaken:omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version45_Event_Filter_Description struct {
	Entity

	Name *string `json:"name:omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version45_Event_Version7 struct {
	McAfee_Epolicy_Orchestrator_Version45_Event

	Signature *McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version7 `json:"signature:omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version45_Event_Version8 struct {
	McAfee_Epolicy_Orchestrator_Version45_Event

	Signature *McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version8 `json:"signature:omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version7 struct {
	Entity

	SignatureName *string `json:"signatureName:omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version8 struct {
	Entity

	SignatureName *string `json:"signatureName:omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version45_Policy_Object struct {
	Entity

	Name *string `json:"name:omitempty"`
}

type McAfee_Epolicy_Orchestrator_Version45_Product_Properties struct {
	Entity

	DatVersion *string `json:"datVersion:omitempty"`
}

type Abuse_Lockdown_Resource struct {
	Entity

	Account     *Account              `json:"account:omitempty"`
	InvoiceItem *Billing_Invoice_Item `json:"invoiceItem:omitempty"`
}

type Account struct {
	Entity

	AbuseEmail                                             *string                                                                 `json:"abuseEmail:omitempty"`
	AbuseEmailCount                                        *uint                                                                   `json:"abuseEmailCount:omitempty"`
	AbuseEmails                                            []Account_AbuseEmail                                                    `json:"abuseEmails:omitempty"`
	AccountContactCount                                    *uint                                                                   `json:"accountContactCount:omitempty"`
	AccountContacts                                        []Account_Contact                                                       `json:"accountContacts:omitempty"`
	AccountLicenseCount                                    *uint                                                                   `json:"accountLicenseCount:omitempty"`
	AccountLicenses                                        []Software_AccountLicense                                               `json:"accountLicenses:omitempty"`
	AccountLinkCount                                       *uint                                                                   `json:"accountLinkCount:omitempty"`
	AccountLinks                                           []Account_Link                                                          `json:"accountLinks:omitempty"`
	AccountManagedResourcesFlag                            *bool                                                                   `json:"accountManagedResourcesFlag:omitempty"`
	AccountStatus                                          *Account_Status                                                         `json:"accountStatus:omitempty"`
	AccountStatusId                                        *int                                                                    `json:"accountStatusId:omitempty"`
	ActiveAccountDiscountBillingItem                       *Billing_Item                                                           `json:"activeAccountDiscountBillingItem:omitempty"`
	ActiveAccountLicenseCount                              *uint                                                                   `json:"activeAccountLicenseCount:omitempty"`
	ActiveAccountLicenses                                  []Software_AccountLicense                                               `json:"activeAccountLicenses:omitempty"`
	ActiveAddressCount                                     *uint                                                                   `json:"activeAddressCount:omitempty"`
	ActiveAddresses                                        []Account_Address                                                       `json:"activeAddresses:omitempty"`
	ActiveBillingAgreementCount                            *uint                                                                   `json:"activeBillingAgreementCount:omitempty"`
	ActiveBillingAgreements                                []Account_Agreement                                                     `json:"activeBillingAgreements:omitempty"`
	ActiveCatalystEnrollment                               *Catalyst_Enrollment                                                    `json:"activeCatalystEnrollment:omitempty"`
	ActiveColocationContainerCount                         *uint                                                                   `json:"activeColocationContainerCount:omitempty"`
	ActiveColocationContainers                             []Billing_Item                                                          `json:"activeColocationContainers:omitempty"`
	ActiveFlexibleCreditEnrollment                         *FlexibleCredit_Enrollment                                              `json:"activeFlexibleCreditEnrollment:omitempty"`
	ActiveNotificationSubscriberCount                      *uint                                                                   `json:"activeNotificationSubscriberCount:omitempty"`
	ActiveNotificationSubscribers                          []Notification_Subscriber                                               `json:"activeNotificationSubscribers:omitempty"`
	ActiveQuoteCount                                       *uint                                                                   `json:"activeQuoteCount:omitempty"`
	ActiveQuotes                                           []Billing_Order_Quote                                                   `json:"activeQuotes:omitempty"`
	ActiveVirtualLicenseCount                              *uint                                                                   `json:"activeVirtualLicenseCount:omitempty"`
	ActiveVirtualLicenses                                  []Software_VirtualLicense                                               `json:"activeVirtualLicenses:omitempty"`
	AdcLoadBalancerCount                                   *uint                                                                   `json:"adcLoadBalancerCount:omitempty"`
	AdcLoadBalancers                                       []Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"adcLoadBalancers:omitempty"`
	Address1                                               *string                                                                 `json:"address1:omitempty"`
	Address2                                               *string                                                                 `json:"address2:omitempty"`
	AddressCount                                           *uint                                                                   `json:"addressCount:omitempty"`
	Addresses                                              []Account_Address                                                       `json:"addresses:omitempty"`
	AffiliateId                                            *string                                                                 `json:"affiliateId:omitempty"`
	AllBillingItems                                        []Billing_Item                                                          `json:"allBillingItems:omitempty"`
	AllCommissionBillingItemCount                          *uint                                                                   `json:"allCommissionBillingItemCount:omitempty"`
	AllCommissionBillingItems                              []Billing_Item                                                          `json:"allCommissionBillingItems:omitempty"`
	AllRecurringTopLevelBillingItemCount                   *uint                                                                   `json:"allRecurringTopLevelBillingItemCount:omitempty"`
	AllRecurringTopLevelBillingItems                       []Billing_Item                                                          `json:"allRecurringTopLevelBillingItems:omitempty"`
	AllRecurringTopLevelBillingItemsUnfiltered             []Billing_Item                                                          `json:"allRecurringTopLevelBillingItemsUnfiltered:omitempty"`
	AllRecurringTopLevelBillingItemsUnfilteredCount        *uint                                                                   `json:"allRecurringTopLevelBillingItemsUnfilteredCount:omitempty"`
	AllSubnetBillingItemCount                              *uint                                                                   `json:"allSubnetBillingItemCount:omitempty"`
	AllSubnetBillingItems                                  []Billing_Item                                                          `json:"allSubnetBillingItems:omitempty"`
	AllTopLevelBillingItemCount                            *uint                                                                   `json:"allTopLevelBillingItemCount:omitempty"`
	AllTopLevelBillingItems                                []Billing_Item                                                          `json:"allTopLevelBillingItems:omitempty"`
	AllTopLevelBillingItemsUnfiltered                      []Billing_Item                                                          `json:"allTopLevelBillingItemsUnfiltered:omitempty"`
	AllTopLevelBillingItemsUnfilteredCount                 *uint                                                                   `json:"allTopLevelBillingItemsUnfilteredCount:omitempty"`
	AllowedPptpVpnQuantity                                 *int                                                                    `json:"allowedPptpVpnQuantity:omitempty"`
	AllowsBluemixAccountLinkingFlag                        *bool                                                                   `json:"allowsBluemixAccountLinkingFlag:omitempty"`
	AlternatePhone                                         *string                                                                 `json:"alternatePhone:omitempty"`
	ApplicationDeliveryControllerCount                     *uint                                                                   `json:"applicationDeliveryControllerCount:omitempty"`
	ApplicationDeliveryControllers                         []Network_Application_Delivery_Controller                               `json:"applicationDeliveryControllers:omitempty"`
	AttributeCount                                         *uint                                                                   `json:"attributeCount:omitempty"`
	Attributes                                             []Account_Attribute                                                     `json:"attributes:omitempty"`
	AvailablePublicNetworkVlanCount                        *uint                                                                   `json:"availablePublicNetworkVlanCount:omitempty"`
	AvailablePublicNetworkVlans                            []Network_Vlan                                                          `json:"availablePublicNetworkVlans:omitempty"`
	Balance                                                *float64                                                                `json:"balance:omitempty"`
	BandwidthAllotmentCount                                *uint                                                                   `json:"bandwidthAllotmentCount:omitempty"`
	BandwidthAllotments                                    []Network_Bandwidth_Version1_Allotment                                  `json:"bandwidthAllotments:omitempty"`
	BandwidthAllotmentsOverAllocation                      []Network_Bandwidth_Version1_Allotment                                  `json:"bandwidthAllotmentsOverAllocation:omitempty"`
	BandwidthAllotmentsOverAllocationCount                 *uint                                                                   `json:"bandwidthAllotmentsOverAllocationCount:omitempty"`
	BandwidthAllotmentsProjectedOverAllocation             []Network_Bandwidth_Version1_Allotment                                  `json:"bandwidthAllotmentsProjectedOverAllocation:omitempty"`
	BandwidthAllotmentsProjectedOverAllocationCount        *uint                                                                   `json:"bandwidthAllotmentsProjectedOverAllocationCount:omitempty"`
	BareMetalInstanceCount                                 *uint                                                                   `json:"bareMetalInstanceCount:omitempty"`
	BareMetalInstances                                     []Hardware                                                              `json:"bareMetalInstances:omitempty"`
	BillingAgreementCount                                  *uint                                                                   `json:"billingAgreementCount:omitempty"`
	BillingAgreements                                      []Account_Agreement                                                     `json:"billingAgreements:omitempty"`
	BillingInfo                                            *Billing_Info                                                           `json:"billingInfo:omitempty"`
	BlockDeviceTemplateGroupCount                          *uint                                                                   `json:"blockDeviceTemplateGroupCount:omitempty"`
	BlockDeviceTemplateGroups                              []Virtual_Guest_Block_Device_Template_Group                             `json:"blockDeviceTemplateGroups:omitempty"`
	BlueIdAuthenticationRequiredFlag                       *bool                                                                   `json:"blueIdAuthenticationRequiredFlag:omitempty"`
	BluemixLinkedFlag                                      *bool                                                                   `json:"bluemixLinkedFlag:omitempty"`
	Brand                                                  *Brand                                                                  `json:"brand:omitempty"`
	BrandAccountFlag                                       *bool                                                                   `json:"brandAccountFlag:omitempty"`
	BrandId                                                *int                                                                    `json:"brandId:omitempty"`
	BrandKeyName                                           *string                                                                 `json:"brandKeyName:omitempty"`
	CanOrderAdditionalVlansFlag                            *bool                                                                   `json:"canOrderAdditionalVlansFlag:omitempty"`
	CartCount                                              *uint                                                                   `json:"cartCount:omitempty"`
	Carts                                                  []Billing_Order_Quote                                                   `json:"carts:omitempty"`
	CatalystEnrollmentCount                                *uint                                                                   `json:"catalystEnrollmentCount:omitempty"`
	CatalystEnrollments                                    []Catalyst_Enrollment                                                   `json:"catalystEnrollments:omitempty"`
	CdnAccountCount                                        *uint                                                                   `json:"cdnAccountCount:omitempty"`
	CdnAccounts                                            []Network_ContentDelivery_Account                                       `json:"cdnAccounts:omitempty"`
	City                                                   *string                                                                 `json:"city:omitempty"`
	ClaimedTaxExemptTxFlag                                 *bool                                                                   `json:"claimedTaxExemptTxFlag:omitempty"`
	ClosedTicketCount                                      *uint                                                                   `json:"closedTicketCount:omitempty"`
	ClosedTickets                                          []Ticket                                                                `json:"closedTickets:omitempty"`
	CompanyName                                            *string                                                                 `json:"companyName:omitempty"`
	Country                                                *string                                                                 `json:"country:omitempty"`
	CreateDate                                             *time.Time                                                              `json:"createDate:omitempty"`
	DatacentersWithSubnetAllocationCount                   *uint                                                                   `json:"datacentersWithSubnetAllocationCount:omitempty"`
	DatacentersWithSubnetAllocations                       []Location                                                              `json:"datacentersWithSubnetAllocations:omitempty"`
	DeviceFingerprintId                                    *string                                                                 `json:"deviceFingerprintId:omitempty"`
	DisablePaymentProcessingFlag                           *bool                                                                   `json:"disablePaymentProcessingFlag:omitempty"`
	DisplaySupportRepresentativeAssignmentCount            *uint                                                                   `json:"displaySupportRepresentativeAssignmentCount:omitempty"`
	DisplaySupportRepresentativeAssignments                []Account_Attachment_Employee                                           `json:"displaySupportRepresentativeAssignments:omitempty"`
	DomainCount                                            *uint                                                                   `json:"domainCount:omitempty"`
	DomainRegistrationCount                                *uint                                                                   `json:"domainRegistrationCount:omitempty"`
	DomainRegistrations                                    []Dns_Domain_Registration                                               `json:"domainRegistrations:omitempty"`
	Domains                                                []Dns_Domain                                                            `json:"domains:omitempty"`
	DomainsWithoutSecondaryDnsRecordCount                  *uint                                                                   `json:"domainsWithoutSecondaryDnsRecordCount:omitempty"`
	DomainsWithoutSecondaryDnsRecords                      []Dns_Domain                                                            `json:"domainsWithoutSecondaryDnsRecords:omitempty"`
	Email                                                  *string                                                                 `json:"email:omitempty"`
	EvaultCapacityGB                                       *uint                                                                   `json:"evaultCapacityGB:omitempty"`
	EvaultMasterUserCount                                  *uint                                                                   `json:"evaultMasterUserCount:omitempty"`
	EvaultMasterUsers                                      []Account_Password                                                      `json:"evaultMasterUsers:omitempty"`
	EvaultNetworkStorage                                   []Network_Storage                                                       `json:"evaultNetworkStorage:omitempty"`
	EvaultNetworkStorageCount                              *uint                                                                   `json:"evaultNetworkStorageCount:omitempty"`
	ExpiredSecurityCertificateCount                        *uint                                                                   `json:"expiredSecurityCertificateCount:omitempty"`
	ExpiredSecurityCertificates                            []Security_Certificate                                                  `json:"expiredSecurityCertificates:omitempty"`
	FacilityLogCount                                       *uint                                                                   `json:"facilityLogCount:omitempty"`
	FacilityLogs                                           []User_Access_Facility_Log                                              `json:"facilityLogs:omitempty"`
	FaxPhone                                               *string                                                                 `json:"faxPhone:omitempty"`
	FirstName                                              *string                                                                 `json:"firstName:omitempty"`
	FlexibleCreditEnrollmentCount                          *uint                                                                   `json:"flexibleCreditEnrollmentCount:omitempty"`
	FlexibleCreditEnrollments                              []FlexibleCredit_Enrollment                                             `json:"flexibleCreditEnrollments:omitempty"`
	GlobalIpRecordCount                                    *uint                                                                   `json:"globalIpRecordCount:omitempty"`
	GlobalIpRecords                                        []Network_Subnet_IpAddress_Global                                       `json:"globalIpRecords:omitempty"`
	GlobalIpv4RecordCount                                  *uint                                                                   `json:"globalIpv4RecordCount:omitempty"`
	GlobalIpv4Records                                      []Network_Subnet_IpAddress_Global                                       `json:"globalIpv4Records:omitempty"`
	GlobalIpv6RecordCount                                  *uint                                                                   `json:"globalIpv6RecordCount:omitempty"`
	GlobalIpv6Records                                      []Network_Subnet_IpAddress_Global                                       `json:"globalIpv6Records:omitempty"`
	GlobalLoadBalancerAccountCount                         *uint                                                                   `json:"globalLoadBalancerAccountCount:omitempty"`
	GlobalLoadBalancerAccounts                             []Network_LoadBalancer_Global_Account                                   `json:"globalLoadBalancerAccounts:omitempty"`
	Hardware                                               []Hardware                                                              `json:"hardware:omitempty"`
	HardwareCount                                          *uint                                                                   `json:"hardwareCount:omitempty"`
	HardwareOverBandwidthAllocation                        []Hardware                                                              `json:"hardwareOverBandwidthAllocation:omitempty"`
	HardwareOverBandwidthAllocationCount                   *uint                                                                   `json:"hardwareOverBandwidthAllocationCount:omitempty"`
	HardwareProjectedOverBandwidthAllocation               []Hardware                                                              `json:"hardwareProjectedOverBandwidthAllocation:omitempty"`
	HardwareProjectedOverBandwidthAllocationCount          *uint                                                                   `json:"hardwareProjectedOverBandwidthAllocationCount:omitempty"`
	HardwareWithCpanel                                     []Hardware                                                              `json:"hardwareWithCpanel:omitempty"`
	HardwareWithCpanelCount                                *uint                                                                   `json:"hardwareWithCpanelCount:omitempty"`
	HardwareWithHelm                                       []Hardware                                                              `json:"hardwareWithHelm:omitempty"`
	HardwareWithHelmCount                                  *uint                                                                   `json:"hardwareWithHelmCount:omitempty"`
	HardwareWithMcafee                                     []Hardware                                                              `json:"hardwareWithMcafee:omitempty"`
	HardwareWithMcafeeAntivirusRedhat                      []Hardware                                                              `json:"hardwareWithMcafeeAntivirusRedhat:omitempty"`
	HardwareWithMcafeeAntivirusRedhatCount                 *uint                                                                   `json:"hardwareWithMcafeeAntivirusRedhatCount:omitempty"`
	HardwareWithMcafeeAntivirusWindowCount                 *uint                                                                   `json:"hardwareWithMcafeeAntivirusWindowCount:omitempty"`
	HardwareWithMcafeeAntivirusWindows                     []Hardware                                                              `json:"hardwareWithMcafeeAntivirusWindows:omitempty"`
	HardwareWithMcafeeCount                                *uint                                                                   `json:"hardwareWithMcafeeCount:omitempty"`
	HardwareWithMcafeeIntrusionDetectionSystem             []Hardware                                                              `json:"hardwareWithMcafeeIntrusionDetectionSystem:omitempty"`
	HardwareWithMcafeeIntrusionDetectionSystemCount        *uint                                                                   `json:"hardwareWithMcafeeIntrusionDetectionSystemCount:omitempty"`
	HardwareWithPlesk                                      []Hardware                                                              `json:"hardwareWithPlesk:omitempty"`
	HardwareWithPleskCount                                 *uint                                                                   `json:"hardwareWithPleskCount:omitempty"`
	HardwareWithQuantastor                                 []Hardware                                                              `json:"hardwareWithQuantastor:omitempty"`
	HardwareWithQuantastorCount                            *uint                                                                   `json:"hardwareWithQuantastorCount:omitempty"`
	HardwareWithUrchin                                     []Hardware                                                              `json:"hardwareWithUrchin:omitempty"`
	HardwareWithUrchinCount                                *uint                                                                   `json:"hardwareWithUrchinCount:omitempty"`
	HardwareWithWindowCount                                *uint                                                                   `json:"hardwareWithWindowCount:omitempty"`
	HardwareWithWindows                                    []Hardware                                                              `json:"hardwareWithWindows:omitempty"`
	HasEvaultBareMetalRestorePluginFlag                    *bool                                                                   `json:"hasEvaultBareMetalRestorePluginFlag:omitempty"`
	HasIderaBareMetalRestorePluginFlag                     *bool                                                                   `json:"hasIderaBareMetalRestorePluginFlag:omitempty"`
	HasPendingOrder                                        *uint                                                                   `json:"hasPendingOrder:omitempty"`
	HasR1softBareMetalRestorePluginFlag                    *bool                                                                   `json:"hasR1softBareMetalRestorePluginFlag:omitempty"`
	HourlyBareMetalInstanceCount                           *uint                                                                   `json:"hourlyBareMetalInstanceCount:omitempty"`
	HourlyBareMetalInstances                               []Hardware                                                              `json:"hourlyBareMetalInstances:omitempty"`
	HourlyServiceBillingItemCount                          *uint                                                                   `json:"hourlyServiceBillingItemCount:omitempty"`
	HourlyServiceBillingItems                              []Billing_Item                                                          `json:"hourlyServiceBillingItems:omitempty"`
	HourlyVirtualGuestCount                                *uint                                                                   `json:"hourlyVirtualGuestCount:omitempty"`
	HourlyVirtualGuests                                    []Virtual_Guest                                                         `json:"hourlyVirtualGuests:omitempty"`
	HubNetworkStorage                                      []Network_Storage                                                       `json:"hubNetworkStorage:omitempty"`
	HubNetworkStorageCount                                 *uint                                                                   `json:"hubNetworkStorageCount:omitempty"`
	Id                                                     *int                                                                    `json:"id:omitempty"`
	InternalNoteCount                                      *uint                                                                   `json:"internalNoteCount:omitempty"`
	InternalNotes                                          []Account_Note                                                          `json:"internalNotes:omitempty"`
	InvoiceCount                                           *uint                                                                   `json:"invoiceCount:omitempty"`
	Invoices                                               []Billing_Invoice                                                       `json:"invoices:omitempty"`
	IpAddressCount                                         *uint                                                                   `json:"ipAddressCount:omitempty"`
	IpAddresses                                            []Network_Subnet_IpAddress                                              `json:"ipAddresses:omitempty"`
	IsReseller                                             *int                                                                    `json:"isReseller:omitempty"`
	IscsiNetworkStorage                                    []Network_Storage                                                       `json:"iscsiNetworkStorage:omitempty"`
	IscsiNetworkStorageCount                               *uint                                                                   `json:"iscsiNetworkStorageCount:omitempty"`
	LastCanceledBillingItem                                *Billing_Item                                                           `json:"lastCanceledBillingItem:omitempty"`
	LastCancelledServerBillingItem                         *Billing_Item                                                           `json:"lastCancelledServerBillingItem:omitempty"`
	LastFiveClosedAbuseTicketCount                         *uint                                                                   `json:"lastFiveClosedAbuseTicketCount:omitempty"`
	LastFiveClosedAbuseTickets                             []Ticket                                                                `json:"lastFiveClosedAbuseTickets:omitempty"`
	LastFiveClosedAccountingTicketCount                    *uint                                                                   `json:"lastFiveClosedAccountingTicketCount:omitempty"`
	LastFiveClosedAccountingTickets                        []Ticket                                                                `json:"lastFiveClosedAccountingTickets:omitempty"`
	LastFiveClosedOtherTicketCount                         *uint                                                                   `json:"lastFiveClosedOtherTicketCount:omitempty"`
	LastFiveClosedOtherTickets                             []Ticket                                                                `json:"lastFiveClosedOtherTickets:omitempty"`
	LastFiveClosedSalesTicketCount                         *uint                                                                   `json:"lastFiveClosedSalesTicketCount:omitempty"`
	LastFiveClosedSalesTickets                             []Ticket                                                                `json:"lastFiveClosedSalesTickets:omitempty"`
	LastFiveClosedSupportTicketCount                       *uint                                                                   `json:"lastFiveClosedSupportTicketCount:omitempty"`
	LastFiveClosedSupportTickets                           []Ticket                                                                `json:"lastFiveClosedSupportTickets:omitempty"`
	LastFiveClosedTicketCount                              *uint                                                                   `json:"lastFiveClosedTicketCount:omitempty"`
	LastFiveClosedTickets                                  []Ticket                                                                `json:"lastFiveClosedTickets:omitempty"`
	LastName                                               *string                                                                 `json:"lastName:omitempty"`
	LateFeeProtectionFlag                                  *bool                                                                   `json:"lateFeeProtectionFlag:omitempty"`
	LatestBillDate                                         *time.Time                                                              `json:"latestBillDate:omitempty"`
	LatestRecurringInvoice                                 *Billing_Invoice                                                        `json:"latestRecurringInvoice:omitempty"`
	LatestRecurringPendingInvoice                          *Billing_Invoice                                                        `json:"latestRecurringPendingInvoice:omitempty"`
	LegacyBandwidthAllotmentCount                          *uint                                                                   `json:"legacyBandwidthAllotmentCount:omitempty"`
	LegacyBandwidthAllotments                              []Network_Bandwidth_Version1_Allotment                                  `json:"legacyBandwidthAllotments:omitempty"`
	LegacyIscsiCapacityGB                                  *uint                                                                   `json:"legacyIscsiCapacityGB:omitempty"`
	LoadBalancerCount                                      *uint                                                                   `json:"loadBalancerCount:omitempty"`
	LoadBalancers                                          []Network_LoadBalancer_VirtualIpAddress                                 `json:"loadBalancers:omitempty"`
	LockboxCapacityGB                                      *uint                                                                   `json:"lockboxCapacityGB:omitempty"`
	LockboxNetworkStorage                                  []Network_Storage                                                       `json:"lockboxNetworkStorage:omitempty"`
	LockboxNetworkStorageCount                             *uint                                                                   `json:"lockboxNetworkStorageCount:omitempty"`
	ManualPaymentsUnderReview                              []Billing_Payment_Card_ManualPayment                                    `json:"manualPaymentsUnderReview:omitempty"`
	ManualPaymentsUnderReviewCount                         *uint                                                                   `json:"manualPaymentsUnderReviewCount:omitempty"`
	MasterUser                                             *User_Customer                                                          `json:"masterUser:omitempty"`
	MediaDataTransferRequestCount                          *uint                                                                   `json:"mediaDataTransferRequestCount:omitempty"`
	MediaDataTransferRequests                              []Account_Media_Data_Transfer_Request                                   `json:"mediaDataTransferRequests:omitempty"`
	MessageQueueAccountCount                               *uint                                                                   `json:"messageQueueAccountCount:omitempty"`
	MessageQueueAccounts                                   []Network_Message_Queue                                                 `json:"messageQueueAccounts:omitempty"`
	ModifyDate                                             *time.Time                                                              `json:"modifyDate:omitempty"`
	MonthlyBareMetalInstanceCount                          *uint                                                                   `json:"monthlyBareMetalInstanceCount:omitempty"`
	MonthlyBareMetalInstances                              []Hardware                                                              `json:"monthlyBareMetalInstances:omitempty"`
	MonthlyVirtualGuestCount                               *uint                                                                   `json:"monthlyVirtualGuestCount:omitempty"`
	MonthlyVirtualGuests                                   []Virtual_Guest                                                         `json:"monthlyVirtualGuests:omitempty"`
	NasNetworkStorage                                      []Network_Storage                                                       `json:"nasNetworkStorage:omitempty"`
	NasNetworkStorageCount                                 *uint                                                                   `json:"nasNetworkStorageCount:omitempty"`
	NetworkCreationFlag                                    *bool                                                                   `json:"networkCreationFlag:omitempty"`
	NetworkGatewayCount                                    *uint                                                                   `json:"networkGatewayCount:omitempty"`
	NetworkGateways                                        []Network_Gateway                                                       `json:"networkGateways:omitempty"`
	NetworkHardware                                        []Hardware                                                              `json:"networkHardware:omitempty"`
	NetworkHardwareCount                                   *uint                                                                   `json:"networkHardwareCount:omitempty"`
	NetworkMessageDeliveryAccountCount                     *uint                                                                   `json:"networkMessageDeliveryAccountCount:omitempty"`
	NetworkMessageDeliveryAccounts                         []Network_Message_Delivery                                              `json:"networkMessageDeliveryAccounts:omitempty"`
	NetworkMonitorDownHardware                             []Hardware                                                              `json:"networkMonitorDownHardware:omitempty"`
	NetworkMonitorDownHardwareCount                        *uint                                                                   `json:"networkMonitorDownHardwareCount:omitempty"`
	NetworkMonitorDownVirtualGuestCount                    *uint                                                                   `json:"networkMonitorDownVirtualGuestCount:omitempty"`
	NetworkMonitorDownVirtualGuests                        []Virtual_Guest                                                         `json:"networkMonitorDownVirtualGuests:omitempty"`
	NetworkMonitorRecoveringHardware                       []Hardware                                                              `json:"networkMonitorRecoveringHardware:omitempty"`
	NetworkMonitorRecoveringHardwareCount                  *uint                                                                   `json:"networkMonitorRecoveringHardwareCount:omitempty"`
	NetworkMonitorRecoveringVirtualGuestCount              *uint                                                                   `json:"networkMonitorRecoveringVirtualGuestCount:omitempty"`
	NetworkMonitorRecoveringVirtualGuests                  []Virtual_Guest                                                         `json:"networkMonitorRecoveringVirtualGuests:omitempty"`
	NetworkMonitorUpHardware                               []Hardware                                                              `json:"networkMonitorUpHardware:omitempty"`
	NetworkMonitorUpHardwareCount                          *uint                                                                   `json:"networkMonitorUpHardwareCount:omitempty"`
	NetworkMonitorUpVirtualGuestCount                      *uint                                                                   `json:"networkMonitorUpVirtualGuestCount:omitempty"`
	NetworkMonitorUpVirtualGuests                          []Virtual_Guest                                                         `json:"networkMonitorUpVirtualGuests:omitempty"`
	NetworkStorage                                         []Network_Storage                                                       `json:"networkStorage:omitempty"`
	NetworkStorageCount                                    *uint                                                                   `json:"networkStorageCount:omitempty"`
	NetworkStorageGroupCount                               *uint                                                                   `json:"networkStorageGroupCount:omitempty"`
	NetworkStorageGroups                                   []Network_Storage_Group                                                 `json:"networkStorageGroups:omitempty"`
	NetworkTunnelContextCount                              *uint                                                                   `json:"networkTunnelContextCount:omitempty"`
	NetworkTunnelContexts                                  []Network_Tunnel_Module_Context                                         `json:"networkTunnelContexts:omitempty"`
	NetworkVlanCount                                       *uint                                                                   `json:"networkVlanCount:omitempty"`
	NetworkVlanSpan                                        *Account_Network_Vlan_Span                                              `json:"networkVlanSpan:omitempty"`
	NetworkVlans                                           []Network_Vlan                                                          `json:"networkVlans:omitempty"`
	NextBillingPublicAllotmentHardwareBandwidthDetailCount *uint                                                                   `json:"nextBillingPublicAllotmentHardwareBandwidthDetailCount:omitempty"`
	NextBillingPublicAllotmentHardwareBandwidthDetails     []Network_Bandwidth_Version1_Allotment                                  `json:"nextBillingPublicAllotmentHardwareBandwidthDetails:omitempty"`
	NextInvoiceIncubatorExemptTotal                        *float64                                                                `json:"nextInvoiceIncubatorExemptTotal:omitempty"`
	NextInvoiceTopLevelBillingItemCount                    *uint                                                                   `json:"nextInvoiceTopLevelBillingItemCount:omitempty"`
	NextInvoiceTopLevelBillingItems                        []Billing_Item                                                          `json:"nextInvoiceTopLevelBillingItems:omitempty"`
	NextInvoiceTotalAmount                                 *float64                                                                `json:"nextInvoiceTotalAmount:omitempty"`
	NextInvoiceTotalOneTimeAmount                          *float64                                                                `json:"nextInvoiceTotalOneTimeAmount:omitempty"`
	NextInvoiceTotalOneTimeTaxAmount                       *float64                                                                `json:"nextInvoiceTotalOneTimeTaxAmount:omitempty"`
	NextInvoiceTotalRecurringAmount                        *float64                                                                `json:"nextInvoiceTotalRecurringAmount:omitempty"`
	NextInvoiceTotalRecurringAmountBeforeAccountDiscount   *float64                                                                `json:"nextInvoiceTotalRecurringAmountBeforeAccountDiscount:omitempty"`
	NextInvoiceTotalRecurringTaxAmount                     *float64                                                                `json:"nextInvoiceTotalRecurringTaxAmount:omitempty"`
	NextInvoiceTotalTaxableRecurringAmount                 *float64                                                                `json:"nextInvoiceTotalTaxableRecurringAmount:omitempty"`
	NotificationSubscriberCount                            *uint                                                                   `json:"notificationSubscriberCount:omitempty"`
	NotificationSubscribers                                []Notification_Subscriber                                               `json:"notificationSubscribers:omitempty"`
	OfficePhone                                            *string                                                                 `json:"officePhone:omitempty"`
	OpenAbuseTicketCount                                   *uint                                                                   `json:"openAbuseTicketCount:omitempty"`
	OpenAbuseTickets                                       []Ticket                                                                `json:"openAbuseTickets:omitempty"`
	OpenAccountingTicketCount                              *uint                                                                   `json:"openAccountingTicketCount:omitempty"`
	OpenAccountingTickets                                  []Ticket                                                                `json:"openAccountingTickets:omitempty"`
	OpenBillingTicketCount                                 *uint                                                                   `json:"openBillingTicketCount:omitempty"`
	OpenBillingTickets                                     []Ticket                                                                `json:"openBillingTickets:omitempty"`
	OpenCancellationRequestCount                           *uint                                                                   `json:"openCancellationRequestCount:omitempty"`
	OpenCancellationRequests                               []Billing_Item_Cancellation_Request                                     `json:"openCancellationRequests:omitempty"`
	OpenOtherTicketCount                                   *uint                                                                   `json:"openOtherTicketCount:omitempty"`
	OpenOtherTickets                                       []Ticket                                                                `json:"openOtherTickets:omitempty"`
	OpenRecurringInvoiceCount                              *uint                                                                   `json:"openRecurringInvoiceCount:omitempty"`
	OpenRecurringInvoices                                  []Billing_Invoice                                                       `json:"openRecurringInvoices:omitempty"`
	OpenSalesTicketCount                                   *uint                                                                   `json:"openSalesTicketCount:omitempty"`
	OpenSalesTickets                                       []Ticket                                                                `json:"openSalesTickets:omitempty"`
	OpenStackAccountLinkCount                              *uint                                                                   `json:"openStackAccountLinkCount:omitempty"`
	OpenStackAccountLinks                                  []Account_Link                                                          `json:"openStackAccountLinks:omitempty"`
	OpenStackObjectStorage                                 []Network_Storage                                                       `json:"openStackObjectStorage:omitempty"`
	OpenStackObjectStorageCount                            *uint                                                                   `json:"openStackObjectStorageCount:omitempty"`
	OpenSupportTicketCount                                 *uint                                                                   `json:"openSupportTicketCount:omitempty"`
	OpenSupportTickets                                     []Ticket                                                                `json:"openSupportTickets:omitempty"`
	OpenTicketCount                                        *uint                                                                   `json:"openTicketCount:omitempty"`
	OpenTickets                                            []Ticket                                                                `json:"openTickets:omitempty"`
	OpenTicketsWaitingOnCustomer                           []Ticket                                                                `json:"openTicketsWaitingOnCustomer:omitempty"`
	OpenTicketsWaitingOnCustomerCount                      *uint                                                                   `json:"openTicketsWaitingOnCustomerCount:omitempty"`
	OrderCount                                             *uint                                                                   `json:"orderCount:omitempty"`
	Orders                                                 []Billing_Order                                                         `json:"orders:omitempty"`
	OrphanBillingItemCount                                 *uint                                                                   `json:"orphanBillingItemCount:omitempty"`
	OrphanBillingItems                                     []Billing_Item                                                          `json:"orphanBillingItems:omitempty"`
	OwnedBrandCount                                        *uint                                                                   `json:"ownedBrandCount:omitempty"`
	OwnedBrands                                            []Brand                                                                 `json:"ownedBrands:omitempty"`
	OwnedHardwareGenericComponentModelCount                *uint                                                                   `json:"ownedHardwareGenericComponentModelCount:omitempty"`
	OwnedHardwareGenericComponentModels                    []Hardware_Component_Model_Generic                                      `json:"ownedHardwareGenericComponentModels:omitempty"`
	PaymentProcessorCount                                  *uint                                                                   `json:"paymentProcessorCount:omitempty"`
	PaymentProcessors                                      []Billing_Payment_Processor                                             `json:"paymentProcessors:omitempty"`
	PendingEventCount                                      *uint                                                                   `json:"pendingEventCount:omitempty"`
	PendingEvents                                          []Notification_Occurrence_Event                                         `json:"pendingEvents:omitempty"`
	PendingInvoice                                         *Billing_Invoice                                                        `json:"pendingInvoice:omitempty"`
	PendingInvoiceTopLevelItemCount                        *uint                                                                   `json:"pendingInvoiceTopLevelItemCount:omitempty"`
	PendingInvoiceTopLevelItems                            []Billing_Invoice_Item                                                  `json:"pendingInvoiceTopLevelItems:omitempty"`
	PendingInvoiceTotalAmount                              *float64                                                                `json:"pendingInvoiceTotalAmount:omitempty"`
	PendingInvoiceTotalOneTimeAmount                       *float64                                                                `json:"pendingInvoiceTotalOneTimeAmount:omitempty"`
	PendingInvoiceTotalOneTimeTaxAmount                    *float64                                                                `json:"pendingInvoiceTotalOneTimeTaxAmount:omitempty"`
	PendingInvoiceTotalRecurringAmount                     *float64                                                                `json:"pendingInvoiceTotalRecurringAmount:omitempty"`
	PendingInvoiceTotalRecurringTaxAmount                  *float64                                                                `json:"pendingInvoiceTotalRecurringTaxAmount:omitempty"`
	PermissionGroupCount                                   *uint                                                                   `json:"permissionGroupCount:omitempty"`
	PermissionGroups                                       []User_Permission_Group                                                 `json:"permissionGroups:omitempty"`
	PermissionRoleCount                                    *uint                                                                   `json:"permissionRoleCount:omitempty"`
	PermissionRoles                                        []User_Permission_Role                                                  `json:"permissionRoles:omitempty"`
	PortableStorageVolumeCount                             *uint                                                                   `json:"portableStorageVolumeCount:omitempty"`
	PortableStorageVolumes                                 []Virtual_Disk_Image                                                    `json:"portableStorageVolumes:omitempty"`
	PostProvisioningHookCount                              *uint                                                                   `json:"postProvisioningHookCount:omitempty"`
	PostProvisioningHooks                                  []Provisioning_Hook                                                     `json:"postProvisioningHooks:omitempty"`
	PostalCode                                             *string                                                                 `json:"postalCode:omitempty"`
	PptpVpnUserCount                                       *uint                                                                   `json:"pptpVpnUserCount:omitempty"`
	PptpVpnUsers                                           []User_Customer                                                         `json:"pptpVpnUsers:omitempty"`
	PreviousRecurringRevenue                               *float64                                                                `json:"previousRecurringRevenue:omitempty"`
	PriceRestrictionCount                                  *uint                                                                   `json:"priceRestrictionCount:omitempty"`
	PriceRestrictions                                      []Product_Item_Price_Account_Restriction                                `json:"priceRestrictions:omitempty"`
	PriorityOneTicketCount                                 *uint                                                                   `json:"priorityOneTicketCount:omitempty"`
	PriorityOneTickets                                     []Ticket                                                                `json:"priorityOneTickets:omitempty"`
	PrivateAllotmentHardwareBandwidthDetailCount           *uint                                                                   `json:"privateAllotmentHardwareBandwidthDetailCount:omitempty"`
	PrivateAllotmentHardwareBandwidthDetails               []Network_Bandwidth_Version1_Allotment                                  `json:"privateAllotmentHardwareBandwidthDetails:omitempty"`
	PrivateBlockDeviceTemplateGroupCount                   *uint                                                                   `json:"privateBlockDeviceTemplateGroupCount:omitempty"`
	PrivateBlockDeviceTemplateGroups                       []Virtual_Guest_Block_Device_Template_Group                             `json:"privateBlockDeviceTemplateGroups:omitempty"`
	PrivateIpAddressCount                                  *uint                                                                   `json:"privateIpAddressCount:omitempty"`
	PrivateIpAddresses                                     []Network_Subnet_IpAddress                                              `json:"privateIpAddresses:omitempty"`
	PrivateNetworkVlanCount                                *uint                                                                   `json:"privateNetworkVlanCount:omitempty"`
	PrivateNetworkVlans                                    []Network_Vlan                                                          `json:"privateNetworkVlans:omitempty"`
	PrivateSubnetCount                                     *uint                                                                   `json:"privateSubnetCount:omitempty"`
	PrivateSubnets                                         []Network_Subnet                                                        `json:"privateSubnets:omitempty"`
	PublicAllotmentHardwareBandwidthDetailCount            *uint                                                                   `json:"publicAllotmentHardwareBandwidthDetailCount:omitempty"`
	PublicAllotmentHardwareBandwidthDetails                []Network_Bandwidth_Version1_Allotment                                  `json:"publicAllotmentHardwareBandwidthDetails:omitempty"`
	PublicIpAddressCount                                   *uint                                                                   `json:"publicIpAddressCount:omitempty"`
	PublicIpAddresses                                      []Network_Subnet_IpAddress                                              `json:"publicIpAddresses:omitempty"`
	PublicNetworkVlanCount                                 *uint                                                                   `json:"publicNetworkVlanCount:omitempty"`
	PublicNetworkVlans                                     []Network_Vlan                                                          `json:"publicNetworkVlans:omitempty"`
	PublicSubnetCount                                      *uint                                                                   `json:"publicSubnetCount:omitempty"`
	PublicSubnets                                          []Network_Subnet                                                        `json:"publicSubnets:omitempty"`
	QuoteCount                                             *uint                                                                   `json:"quoteCount:omitempty"`
	Quotes                                                 []Billing_Order_Quote                                                   `json:"quotes:omitempty"`
	RecentEventCount                                       *uint                                                                   `json:"recentEventCount:omitempty"`
	RecentEvents                                           []Notification_Occurrence_Event                                         `json:"recentEvents:omitempty"`
	ReferralPartner                                        *Account                                                                `json:"referralPartner:omitempty"`
	ReferredAccountCount                                   *uint                                                                   `json:"referredAccountCount:omitempty"`
	ReferredAccounts                                       []Account                                                               `json:"referredAccounts:omitempty"`
	RegulatedWorkloadCount                                 *uint                                                                   `json:"regulatedWorkloadCount:omitempty"`
	RegulatedWorkloads                                     []Legal_RegulatedWorkload                                               `json:"regulatedWorkloads:omitempty"`
	RemoteManagementCommandRequestCount                    *uint                                                                   `json:"remoteManagementCommandRequestCount:omitempty"`
	RemoteManagementCommandRequests                        []Hardware_Component_RemoteManagement_Command_Request                   `json:"remoteManagementCommandRequests:omitempty"`
	ReplicationEventCount                                  *uint                                                                   `json:"replicationEventCount:omitempty"`
	ReplicationEvents                                      []Network_Storage_Event                                                 `json:"replicationEvents:omitempty"`
	ResourceGroupCount                                     *uint                                                                   `json:"resourceGroupCount:omitempty"`
	ResourceGroups                                         []Resource_Group                                                        `json:"resourceGroups:omitempty"`
	RouterCount                                            *uint                                                                   `json:"routerCount:omitempty"`
	Routers                                                []Hardware                                                              `json:"routers:omitempty"`
	RwhoisData                                             *Network_Subnet_Rwhois_Data                                             `json:"rwhoisData:omitempty"`
	SalesforceAccountLink                                  *Account_Link                                                           `json:"salesforceAccountLink:omitempty"`
	SamlAuthentication                                     *Account_Authentication_Saml                                            `json:"samlAuthentication:omitempty"`
	ScaleGroupCount                                        *uint                                                                   `json:"scaleGroupCount:omitempty"`
	ScaleGroups                                            []Scale_Group                                                           `json:"scaleGroups:omitempty"`
	SecondaryDomainCount                                   *uint                                                                   `json:"secondaryDomainCount:omitempty"`
	SecondaryDomains                                       []Dns_Secondary                                                         `json:"secondaryDomains:omitempty"`
	SecurityCertificateCount                               *uint                                                                   `json:"securityCertificateCount:omitempty"`
	SecurityCertificates                                   []Security_Certificate                                                  `json:"securityCertificates:omitempty"`
	SecurityScanRequestCount                               *uint                                                                   `json:"securityScanRequestCount:omitempty"`
	SecurityScanRequests                                   []Network_Security_Scanner_Request                                      `json:"securityScanRequests:omitempty"`
	ServiceBillingItemCount                                *uint                                                                   `json:"serviceBillingItemCount:omitempty"`
	ServiceBillingItems                                    []Billing_Item                                                          `json:"serviceBillingItems:omitempty"`
	ShipmentCount                                          *uint                                                                   `json:"shipmentCount:omitempty"`
	Shipments                                              []Account_Shipment                                                      `json:"shipments:omitempty"`
	SshKeyCount                                            *uint                                                                   `json:"sshKeyCount:omitempty"`
	SshKeys                                                []Security_Ssh_Key                                                      `json:"sshKeys:omitempty"`
	SslVpnUserCount                                        *uint                                                                   `json:"sslVpnUserCount:omitempty"`
	SslVpnUsers                                            []User_Customer                                                         `json:"sslVpnUsers:omitempty"`
	StandardPoolVirtualGuestCount                          *uint                                                                   `json:"standardPoolVirtualGuestCount:omitempty"`
	StandardPoolVirtualGuests                              []Virtual_Guest                                                         `json:"standardPoolVirtualGuests:omitempty"`
	State                                                  *string                                                                 `json:"state:omitempty"`
	StatusDate                                             *time.Time                                                              `json:"statusDate:omitempty"`
	SubnetCount                                            *uint                                                                   `json:"subnetCount:omitempty"`
	SubnetRegistrationCount                                *uint                                                                   `json:"subnetRegistrationCount:omitempty"`
	SubnetRegistrationDetailCount                          *uint                                                                   `json:"subnetRegistrationDetailCount:omitempty"`
	SubnetRegistrationDetails                              []Account_Regional_Registry_Detail                                      `json:"subnetRegistrationDetails:omitempty"`
	SubnetRegistrations                                    []Network_Subnet_Registration                                           `json:"subnetRegistrations:omitempty"`
	Subnets                                                []Network_Subnet                                                        `json:"subnets:omitempty"`
	SupportRepresentativeCount                             *uint                                                                   `json:"supportRepresentativeCount:omitempty"`
	SupportRepresentatives                                 []User_Employee                                                         `json:"supportRepresentatives:omitempty"`
	SupportSubscriptionCount                               *uint                                                                   `json:"supportSubscriptionCount:omitempty"`
	SupportSubscriptions                                   []Billing_Item                                                          `json:"supportSubscriptions:omitempty"`
	SupportTier                                            *string                                                                 `json:"supportTier:omitempty"`
	SuppressInvoicesFlag                                   *bool                                                                   `json:"suppressInvoicesFlag:omitempty"`
	TagCount                                               *uint                                                                   `json:"tagCount:omitempty"`
	Tags                                                   []Tag                                                                   `json:"tags:omitempty"`
	TicketCount                                            *uint                                                                   `json:"ticketCount:omitempty"`
	Tickets                                                []Ticket                                                                `json:"tickets:omitempty"`
	TicketsClosedInTheLastThreeDays                        []Ticket                                                                `json:"ticketsClosedInTheLastThreeDays:omitempty"`
	TicketsClosedInTheLastThreeDaysCount                   *uint                                                                   `json:"ticketsClosedInTheLastThreeDaysCount:omitempty"`
	TicketsClosedToday                                     []Ticket                                                                `json:"ticketsClosedToday:omitempty"`
	TicketsClosedTodayCount                                *uint                                                                   `json:"ticketsClosedTodayCount:omitempty"`
	TranscodeAccountCount                                  *uint                                                                   `json:"transcodeAccountCount:omitempty"`
	TranscodeAccounts                                      []Network_Media_Transcode_Account                                       `json:"transcodeAccounts:omitempty"`
	UpgradeRequestCount                                    *uint                                                                   `json:"upgradeRequestCount:omitempty"`
	UpgradeRequests                                        []Product_Upgrade_Request                                               `json:"upgradeRequests:omitempty"`
	UserCount                                              *uint                                                                   `json:"userCount:omitempty"`
	Users                                                  []User_Customer                                                         `json:"users:omitempty"`
	ValidSecurityCertificateCount                          *uint                                                                   `json:"validSecurityCertificateCount:omitempty"`
	ValidSecurityCertificates                              []Security_Certificate                                                  `json:"validSecurityCertificates:omitempty"`
	VdrUpdatesInProgressFlag                               *bool                                                                   `json:"vdrUpdatesInProgressFlag:omitempty"`
	VirtualDedicatedRackCount                              *uint                                                                   `json:"virtualDedicatedRackCount:omitempty"`
	VirtualDedicatedRacks                                  []Network_Bandwidth_Version1_Allotment                                  `json:"virtualDedicatedRacks:omitempty"`
	VirtualDiskImageCount                                  *uint                                                                   `json:"virtualDiskImageCount:omitempty"`
	VirtualDiskImages                                      []Virtual_Disk_Image                                                    `json:"virtualDiskImages:omitempty"`
	VirtualGuestCount                                      *uint                                                                   `json:"virtualGuestCount:omitempty"`
	VirtualGuests                                          []Virtual_Guest                                                         `json:"virtualGuests:omitempty"`
	VirtualGuestsOverBandwidthAllocation                   []Virtual_Guest                                                         `json:"virtualGuestsOverBandwidthAllocation:omitempty"`
	VirtualGuestsOverBandwidthAllocationCount              *uint                                                                   `json:"virtualGuestsOverBandwidthAllocationCount:omitempty"`
	VirtualGuestsProjectedOverBandwidthAllocation          []Virtual_Guest                                                         `json:"virtualGuestsProjectedOverBandwidthAllocation:omitempty"`
	VirtualGuestsProjectedOverBandwidthAllocationCount     *uint                                                                   `json:"virtualGuestsProjectedOverBandwidthAllocationCount:omitempty"`
	VirtualGuestsWithCpanel                                []Virtual_Guest                                                         `json:"virtualGuestsWithCpanel:omitempty"`
	VirtualGuestsWithCpanelCount                           *uint                                                                   `json:"virtualGuestsWithCpanelCount:omitempty"`
	VirtualGuestsWithMcafee                                []Virtual_Guest                                                         `json:"virtualGuestsWithMcafee:omitempty"`
	VirtualGuestsWithMcafeeAntivirusRedhat                 []Virtual_Guest                                                         `json:"virtualGuestsWithMcafeeAntivirusRedhat:omitempty"`
	VirtualGuestsWithMcafeeAntivirusRedhatCount            *uint                                                                   `json:"virtualGuestsWithMcafeeAntivirusRedhatCount:omitempty"`
	VirtualGuestsWithMcafeeAntivirusWindowCount            *uint                                                                   `json:"virtualGuestsWithMcafeeAntivirusWindowCount:omitempty"`
	VirtualGuestsWithMcafeeAntivirusWindows                []Virtual_Guest                                                         `json:"virtualGuestsWithMcafeeAntivirusWindows:omitempty"`
	VirtualGuestsWithMcafeeCount                           *uint                                                                   `json:"virtualGuestsWithMcafeeCount:omitempty"`
	VirtualGuestsWithMcafeeIntrusionDetectionSystem        []Virtual_Guest                                                         `json:"virtualGuestsWithMcafeeIntrusionDetectionSystem:omitempty"`
	VirtualGuestsWithMcafeeIntrusionDetectionSystemCount   *uint                                                                   `json:"virtualGuestsWithMcafeeIntrusionDetectionSystemCount:omitempty"`
	VirtualGuestsWithPlesk                                 []Virtual_Guest                                                         `json:"virtualGuestsWithPlesk:omitempty"`
	VirtualGuestsWithPleskCount                            *uint                                                                   `json:"virtualGuestsWithPleskCount:omitempty"`
	VirtualGuestsWithQuantastor                            []Virtual_Guest                                                         `json:"virtualGuestsWithQuantastor:omitempty"`
	VirtualGuestsWithQuantastorCount                       *uint                                                                   `json:"virtualGuestsWithQuantastorCount:omitempty"`
	VirtualGuestsWithUrchin                                []Virtual_Guest                                                         `json:"virtualGuestsWithUrchin:omitempty"`
	VirtualGuestsWithUrchinCount                           *uint                                                                   `json:"virtualGuestsWithUrchinCount:omitempty"`
	VirtualPrivateRack                                     *Network_Bandwidth_Version1_Allotment                                   `json:"virtualPrivateRack:omitempty"`
	VirtualStorageArchiveRepositories                      []Virtual_Storage_Repository                                            `json:"virtualStorageArchiveRepositories:omitempty"`
	VirtualStorageArchiveRepositoryCount                   *uint                                                                   `json:"virtualStorageArchiveRepositoryCount:omitempty"`
	VirtualStoragePublicRepositories                       []Virtual_Storage_Repository                                            `json:"virtualStoragePublicRepositories:omitempty"`
	VirtualStoragePublicRepositoryCount                    *uint                                                                   `json:"virtualStoragePublicRepositoryCount:omitempty"`
}

type Account_AbuseEmail struct {
	Entity

	Account *Account `json:"account:omitempty"`
	Email   *string  `json:"email:omitempty"`
}

type Account_Address struct {
	Entity

	Account        *Account              `json:"account:omitempty"`
	AccountId      *int                  `json:"accountId:omitempty"`
	Address1       *string               `json:"address1:omitempty"`
	Address2       *string               `json:"address2:omitempty"`
	City           *string               `json:"city:omitempty"`
	ContactName    *string               `json:"contactName:omitempty"`
	Country        *string               `json:"country:omitempty"`
	CreateUser     *User_Customer        `json:"createUser:omitempty"`
	Description    *string               `json:"description:omitempty"`
	Id             *int                  `json:"id:omitempty"`
	IsActive       *int                  `json:"isActive:omitempty"`
	Location       *Location             `json:"location:omitempty"`
	LocationId     *int                  `json:"locationId:omitempty"`
	ModifyEmployee *User_Employee        `json:"modifyEmployee:omitempty"`
	ModifyUser     *User_Customer        `json:"modifyUser:omitempty"`
	PostalCode     *string               `json:"postalCode:omitempty"`
	State          *string               `json:"state:omitempty"`
	Type           *Account_Address_Type `json:"type:omitempty"`
}

type Account_Address_Type struct {
	Entity

	CreateDate *time.Time `json:"createDate:omitempty"`
	Id         *int       `json:"id:omitempty"`
	KeyName    *string    `json:"keyName:omitempty"`
	Name       *string    `json:"name:omitempty"`
}

type Account_Affiliation struct {
	Entity

	Account     *Account   `json:"account:omitempty"`
	AccountId   *int       `json:"accountId:omitempty"`
	AffiliateId *string    `json:"affiliateId:omitempty"`
	CreateDate  *time.Time `json:"createDate:omitempty"`
	Id          *int       `json:"id:omitempty"`
	ModifyDate  *time.Time `json:"modifyDate:omitempty"`
}

type Account_Agreement struct {
	Entity

	Account                           *Account                         `json:"account:omitempty"`
	AgreementType                     *Account_Agreement_Type          `json:"agreementType:omitempty"`
	AgreementTypeId                   *int                             `json:"agreementTypeId:omitempty"`
	AttachedBillingAgreementFileCount *uint                            `json:"attachedBillingAgreementFileCount:omitempty"`
	AttachedBillingAgreementFiles     []Account_MasterServiceAgreement `json:"attachedBillingAgreementFiles:omitempty"`
	AutoRenew                         *int                             `json:"autoRenew:omitempty"`
	BillingItemCount                  *uint                            `json:"billingItemCount:omitempty"`
	BillingItems                      []Billing_Item                   `json:"billingItems:omitempty"`
	CancellationFee                   *int                             `json:"cancellationFee:omitempty"`
	CreateDate                        *time.Time                       `json:"createDate:omitempty"`
	DurationMonths                    *int                             `json:"durationMonths:omitempty"`
	EndDate                           *time.Time                       `json:"endDate:omitempty"`
	Id                                *int                             `json:"id:omitempty"`
	StartDate                         *time.Time                       `json:"startDate:omitempty"`
	Status                            *Account_Agreement_Status        `json:"status:omitempty"`
	StatusId                          *int                             `json:"statusId:omitempty"`
	Title                             *string                          `json:"title:omitempty"`
	TopLevelBillingItemCount          *uint                            `json:"topLevelBillingItemCount:omitempty"`
	TopLevelBillingItems              []Billing_Item                   `json:"topLevelBillingItems:omitempty"`
}

type Account_Agreement_Status struct {
	Entity

	Name *string `json:"name:omitempty"`
}

type Account_Agreement_Type struct {
	Entity

	Name *string `json:"name:omitempty"`
}

type Account_Attachment_Employee struct {
	Entity

	Account      *Account                          `json:"account:omitempty"`
	Employee     *User_Employee                    `json:"employee:omitempty"`
	EmployeeRole *Account_Attachment_Employee_Role `json:"employeeRole:omitempty"`
	RoleId       *int                              `json:"roleId:omitempty"`
}

type Account_Attachment_Employee_Role struct {
	Entity

	Keyname *string `json:"keyname:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Account_Attribute struct {
	Entity

	Account                *Account                `json:"account:omitempty"`
	AccountAttributeType   *Account_Attribute_Type `json:"accountAttributeType:omitempty"`
	AccountAttributeTypeId *int                    `json:"accountAttributeTypeId:omitempty"`
	AccountId              *int                    `json:"accountId:omitempty"`
	Id                     *int                    `json:"id:omitempty"`
	Value                  *string                 `json:"value:omitempty"`
}

type Account_Attribute_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Account_Authentication_Attribute struct {
	Entity

	Account              *Account                               `json:"account:omitempty"`
	AccountId            *int                                   `json:"accountId:omitempty"`
	AuthenticationRecord *Account_Authentication_Saml           `json:"authenticationRecord:omitempty"`
	Id                   *int                                   `json:"id:omitempty"`
	Type                 *Account_Authentication_Attribute_Type `json:"type:omitempty"`
	TypeId               *int                                   `json:"typeId:omitempty"`
	Value                *string                                `json:"value:omitempty"`
}

type Account_Authentication_Attribute_Type struct {
	Entity

	Description  *string `json:"description:omitempty"`
	Id           *int    `json:"id:omitempty"`
	KeyName      *string `json:"keyName:omitempty"`
	Name         *string `json:"name:omitempty"`
	ValueExample *string `json:"valueExample:omitempty"`
}

type Account_Authentication_OpenIdConnect_Option struct {
	Entity

	Key   *string `json:"key:omitempty"`
	Value *string `json:"value:omitempty"`
}

type Account_Authentication_OpenIdConnect_RegistrationInformation struct {
	Entity

	ExistingBlueIdFlag       *bool          `json:"existingBlueIdFlag:omitempty"`
	FederatedEmailDomainFlag *bool          `json:"federatedEmailDomainFlag:omitempty"`
	User                     *User_Customer `json:"user:omitempty"`
}

type Account_Authentication_Saml struct {
	Entity

	Account                             *Account                           `json:"account:omitempty"`
	AccountId                           *string                            `json:"accountId:omitempty"`
	AttributeCount                      *uint                              `json:"attributeCount:omitempty"`
	Attributes                          []Account_Authentication_Attribute `json:"attributes:omitempty"`
	Certificate                         *string                            `json:"certificate:omitempty"`
	CertificateFingerprint              *string                            `json:"certificateFingerprint:omitempty"`
	EntityId                            *string                            `json:"entityId:omitempty"`
	Id                                  *int                               `json:"id:omitempty"`
	ServiceProviderCertificate          *string                            `json:"serviceProviderCertificate:omitempty"`
	ServiceProviderEntityId             *string                            `json:"serviceProviderEntityId:omitempty"`
	ServiceProviderPublicKey            *string                            `json:"serviceProviderPublicKey:omitempty"`
	ServiceProviderSingleLogoutEncoding *string                            `json:"serviceProviderSingleLogoutEncoding:omitempty"`
	ServiceProviderSingleLogoutUrl      *string                            `json:"serviceProviderSingleLogoutUrl:omitempty"`
	ServiceProviderSingleSignOnEncoding *string                            `json:"serviceProviderSingleSignOnEncoding:omitempty"`
	ServiceProviderSingleSignOnUrl      *string                            `json:"serviceProviderSingleSignOnUrl:omitempty"`
	SingleLogoutEncoding                *string                            `json:"singleLogoutEncoding:omitempty"`
	SingleLogoutUrl                     *string                            `json:"singleLogoutUrl:omitempty"`
	SingleSignOnEncoding                *string                            `json:"singleSignOnEncoding:omitempty"`
	SingleSignOnUrl                     *string                            `json:"singleSignOnUrl:omitempty"`
}

type Account_Classification_Group_Type struct {
	Entity

	KeyName *string `json:"keyName:omitempty"`
}

type Account_Contact struct {
	Entity

	Account        *Account              `json:"account:omitempty"`
	AccountId      *int                  `json:"accountId:omitempty"`
	Address1       *string               `json:"address1:omitempty"`
	Address2       *string               `json:"address2:omitempty"`
	AlternatePhone *string               `json:"alternatePhone:omitempty"`
	City           *string               `json:"city:omitempty"`
	CompanyName    *string               `json:"companyName:omitempty"`
	Country        *string               `json:"country:omitempty"`
	CreateDate     *time.Time            `json:"createDate:omitempty"`
	Email          *string               `json:"email:omitempty"`
	FaxPhone       *string               `json:"faxPhone:omitempty"`
	FirstName      *string               `json:"firstName:omitempty"`
	Id             *int                  `json:"id:omitempty"`
	JobTitle       *string               `json:"jobTitle:omitempty"`
	LastName       *string               `json:"lastName:omitempty"`
	ModifyDate     *time.Time            `json:"modifyDate:omitempty"`
	OfficePhone    *string               `json:"officePhone:omitempty"`
	PostalCode     *string               `json:"postalCode:omitempty"`
	ProfileName    *string               `json:"profileName:omitempty"`
	State          *string               `json:"state:omitempty"`
	Type           *Account_Contact_Type `json:"type:omitempty"`
	TypeId         *int                  `json:"typeId:omitempty"`
	Url            *string               `json:"url:omitempty"`
}

type Account_Contact_Type struct {
	Entity

	CreateDate  *time.Time `json:"createDate:omitempty"`
	Description *string    `json:"description:omitempty"`
	Id          *int       `json:"id:omitempty"`
	KeyName     *string    `json:"keyName:omitempty"`
	ModifyDate  *time.Time `json:"modifyDate:omitempty"`
	Name        *string    `json:"name:omitempty"`
}

type Account_Historical_Report struct {
	Entity
}

type Account_Link struct {
	Entity

	Account                          *Account          `json:"account:omitempty"`
	AccountId                        *int              `json:"accountId:omitempty"`
	CreateDate                       *time.Time        `json:"createDate:omitempty"`
	DestinationAccountAlphanumericId *string           `json:"destinationAccountAlphanumericId:omitempty"`
	DestinationAccountId             *int              `json:"destinationAccountId:omitempty"`
	Id                               *int              `json:"id:omitempty"`
	ServiceProvider                  *Service_Provider `json:"serviceProvider:omitempty"`
	ServiceProviderId                *int              `json:"serviceProviderId:omitempty"`
}

type Account_Link_Bluemix struct {
	Account_Link
}

type Account_Link_OpenStack struct {
	Account_Link

	DomainId *string `json:"domainId:omitempty"`
}

type Account_Link_OpenStack_DomainCreationDetails struct {
	Entity

	DomainId *string `json:"domainId:omitempty"`
	UserId   *string `json:"userId:omitempty"`
	UserName *string `json:"userName:omitempty"`
}

type Account_Link_OpenStack_LinkRequest struct {
	Entity

	DesiredPassword    *string `json:"desiredPassword:omitempty"`
	DesiredProjectName *string `json:"desiredProjectName:omitempty"`
	DesiredUsername    *string `json:"desiredUsername:omitempty"`
}

type Account_Link_OpenStack_ProjectCreationDetails struct {
	Entity

	DomainId    *string `json:"domainId:omitempty"`
	ProjectId   *string `json:"projectId:omitempty"`
	ProjectName *string `json:"projectName:omitempty"`
	UserId      *string `json:"userId:omitempty"`
	UserName    *string `json:"userName:omitempty"`
}

type Account_Link_OpenStack_ProjectDetails struct {
	Entity

	ProjectId   *string `json:"projectId:omitempty"`
	ProjectName *string `json:"projectName:omitempty"`
}

type Account_Link_ThePlanet struct {
	Account_Link
}

type Account_Link_Vendor struct {
	Entity

	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Account_Lockdown_Request struct {
	Entity

	AccountId  *int       `json:"accountId:omitempty"`
	Action     *string    `json:"action:omitempty"`
	CreateDate *time.Time `json:"createDate:omitempty"`
	Id         *int       `json:"id:omitempty"`
	ModifyDate *time.Time `json:"modifyDate:omitempty"`
	Status     *string    `json:"status:omitempty"`
}

type Account_MasterServiceAgreement struct {
	Entity

	Account   *Account `json:"account:omitempty"`
	AccountId *int     `json:"accountId:omitempty"`
	Guid      *string  `json:"guid:omitempty"`
	Id        *int     `json:"id:omitempty"`
	Name      *string  `json:"name:omitempty"`
}

type Account_Media struct {
	Entity

	Account        *Account                             `json:"account:omitempty"`
	CreateUser     *User_Customer                       `json:"createUser:omitempty"`
	Datacenter     *Location                            `json:"datacenter:omitempty"`
	Description    *string                              `json:"description:omitempty"`
	Id             *int                                 `json:"id:omitempty"`
	ModifyEmployee *User_Employee                       `json:"modifyEmployee:omitempty"`
	ModifyUser     *User_Customer                       `json:"modifyUser:omitempty"`
	Request        *Account_Media_Data_Transfer_Request `json:"request:omitempty"`
	RequestId      *int                                 `json:"requestId:omitempty"`
	SerialNumber   *string                              `json:"serialNumber:omitempty"`
	Type           *Account_Media_Type                  `json:"type:omitempty"`
	TypeId         *int                                 `json:"typeId:omitempty"`
	Volume         *Network_Storage                     `json:"volume:omitempty"`
}

type Account_Media_Data_Transfer_Request struct {
	Entity

	Account           *Account                                    `json:"account:omitempty"`
	AccountId         *int                                        `json:"accountId:omitempty"`
	ActiveTicketCount *uint                                       `json:"activeTicketCount:omitempty"`
	ActiveTickets     []Ticket                                    `json:"activeTickets:omitempty"`
	BillingItem       *Billing_Item                               `json:"billingItem:omitempty"`
	CreateUser        *User_Customer                              `json:"createUser:omitempty"`
	CreateUserId      *int                                        `json:"createUserId:omitempty"`
	EndDate           *time.Time                                  `json:"endDate:omitempty"`
	Id                *int                                        `json:"id:omitempty"`
	Media             *Account_Media                              `json:"media:omitempty"`
	ModifyEmployee    *User_Employee                              `json:"modifyEmployee:omitempty"`
	ModifyUser        *User_Customer                              `json:"modifyUser:omitempty"`
	ModifyUserId      *int                                        `json:"modifyUserId:omitempty"`
	ShipmentCount     *uint                                       `json:"shipmentCount:omitempty"`
	Shipments         []Account_Shipment                          `json:"shipments:omitempty"`
	StartDate         *time.Time                                  `json:"startDate:omitempty"`
	Status            *Account_Media_Data_Transfer_Request_Status `json:"status:omitempty"`
	StatusId          *int                                        `json:"statusId:omitempty"`
	TicketCount       *uint                                       `json:"ticketCount:omitempty"`
	Tickets           []Ticket                                    `json:"tickets:omitempty"`
}

type Account_Media_Data_Transfer_Request_Status struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Account_Media_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Account_Network_Vlan_Span struct {
	Entity

	Account          *Account   `json:"account:omitempty"`
	EnabledFlag      *bool      `json:"enabledFlag:omitempty"`
	Id               *int       `json:"id:omitempty"`
	LastAppliedDate  *time.Time `json:"lastAppliedDate:omitempty"`
	LastVerifiedDate *time.Time `json:"lastVerifiedDate:omitempty"`
	ModifyDate       *time.Time `json:"modifyDate:omitempty"`
}

type Account_Note struct {
	Entity

	Account          *Account               `json:"account:omitempty"`
	AccountId        *int                   `json:"accountId:omitempty"`
	CreateDate       *time.Time             `json:"createDate:omitempty"`
	Customer         *User_Customer         `json:"customer:omitempty"`
	Id               *int                   `json:"id:omitempty"`
	ModifyDate       *time.Time             `json:"modifyDate:omitempty"`
	Note             *string                `json:"note:omitempty"`
	NoteHistory      []Account_Note_History `json:"noteHistory:omitempty"`
	NoteHistoryCount *uint                  `json:"noteHistoryCount:omitempty"`
	NoteType         *Account_Note_Type     `json:"noteType:omitempty"`
	NoteTypeId       *int                   `json:"noteTypeId:omitempty"`
	UserId           *int                   `json:"userId:omitempty"`
}

type Account_Note_History struct {
	Entity

	AccountNote   *Account_Note  `json:"accountNote:omitempty"`
	AccountNoteId *int           `json:"accountNoteId:omitempty"`
	CreateDate    *time.Time     `json:"createDate:omitempty"`
	Customer      *User_Customer `json:"customer:omitempty"`
	Id            *int           `json:"id:omitempty"`
	ModifyDate    *time.Time     `json:"modifyDate:omitempty"`
	Note          *string        `json:"note:omitempty"`
	UserId        *int           `json:"userId:omitempty"`
}

type Account_Note_Type struct {
	Entity

	BrandId         *int       `json:"brandId:omitempty"`
	CreateDate      *time.Time `json:"createDate:omitempty"`
	Description     *string    `json:"description:omitempty"`
	Id              *int       `json:"id:omitempty"`
	KeyName         *string    `json:"keyName:omitempty"`
	ModifyDate      *time.Time `json:"modifyDate:omitempty"`
	Name            *string    `json:"name:omitempty"`
	ValueExpression *string    `json:"valueExpression:omitempty"`
}

type Account_Partner_Referral_Prospect struct {
	User_Customer_Prospect

	CompanyName  *string `json:"companyName:omitempty"`
	EmailAddress *string `json:"emailAddress:omitempty"`
	FirstName    *string `json:"firstName:omitempty"`
	Id           *int    `json:"id:omitempty"`
	LastName     *string `json:"lastName:omitempty"`
}

type Account_Password struct {
	Entity

	Account   *Account               `json:"account:omitempty"`
	AccountId *int                   `json:"accountId:omitempty"`
	Id        *int                   `json:"id:omitempty"`
	Notes     *string                `json:"notes:omitempty"`
	Password  *string                `json:"password:omitempty"`
	Type      *Account_Password_Type `json:"type:omitempty"`
	TypeId    *int                   `json:"typeId:omitempty"`
	Username  *string                `json:"username:omitempty"`
}

type Account_Password_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
}

type Account_Regional_Registry_Detail struct {
	Entity

	Account                          *Account                                    `json:"account:omitempty"`
	AccountId                        *int                                        `json:"accountId:omitempty"`
	CreateDate                       *time.Time                                  `json:"createDate:omitempty"`
	DetailCount                      *uint                                       `json:"detailCount:omitempty"`
	DetailType                       *Account_Regional_Registry_Detail_Type      `json:"detailType:omitempty"`
	DetailTypeId                     *int                                        `json:"detailTypeId:omitempty"`
	Details                          []Network_Subnet_Registration_Details       `json:"details:omitempty"`
	Id                               *int                                        `json:"id:omitempty"`
	ModifyDate                       *time.Time                                  `json:"modifyDate:omitempty"`
	Properties                       []Account_Regional_Registry_Detail_Property `json:"properties:omitempty"`
	PropertyCount                    *uint                                       `json:"propertyCount:omitempty"`
	RegionalInternetRegistryHandle   *Account_Rwhois_Handle                      `json:"regionalInternetRegistryHandle:omitempty"`
	RegionalInternetRegistryHandleId *int                                        `json:"regionalInternetRegistryHandleId:omitempty"`
}

type Account_Regional_Registry_Detail_Property struct {
	Entity

	CreateDate           *time.Time                                      `json:"createDate:omitempty"`
	Detail               *Account_Regional_Registry_Detail               `json:"detail:omitempty"`
	Id                   *int                                            `json:"id:omitempty"`
	ModifyDate           *time.Time                                      `json:"modifyDate:omitempty"`
	PropertyType         *Account_Regional_Registry_Detail_Property_Type `json:"propertyType:omitempty"`
	PropertyTypeId       *int                                            `json:"propertyTypeId:omitempty"`
	RegistrationDetailId *int                                            `json:"registrationDetailId:omitempty"`
	SequencePosition     *int                                            `json:"sequencePosition:omitempty"`
	Value                *string                                         `json:"value:omitempty"`
}

type Account_Regional_Registry_Detail_Property_Type struct {
	Entity

	CreateDate      *time.Time `json:"createDate:omitempty"`
	Id              *int       `json:"id:omitempty"`
	KeyName         *string    `json:"keyName:omitempty"`
	ModifyDate      *time.Time `json:"modifyDate:omitempty"`
	Name            *string    `json:"name:omitempty"`
	ValueExpression *string    `json:"valueExpression:omitempty"`
}

type Account_Regional_Registry_Detail_Type struct {
	Entity

	CreateDate *time.Time `json:"createDate:omitempty"`
	Id         *int       `json:"id:omitempty"`
	KeyName    *string    `json:"keyName:omitempty"`
	ModifyDate *time.Time `json:"modifyDate:omitempty"`
	Name       *string    `json:"name:omitempty"`
}

type Account_Regional_Registry_Detail_Version4_Person_Default struct {
	Account_Regional_Registry_Detail
}

type Account_Reports_Request struct {
	Entity

	Account                *Account                `json:"account:omitempty"`
	AccountContact         *Account_Contact        `json:"accountContact:omitempty"`
	AccountContactId       *int                    `json:"accountContactId:omitempty"`
	AccountId              *int                    `json:"accountId:omitempty"`
	ComplianceReportTypeId *string                 `json:"complianceReportTypeId:omitempty"`
	CreateDate             *time.Time              `json:"createDate:omitempty"`
	EmployeeRecordId       *int                    `json:"employeeRecordId:omitempty"`
	Id                     *int                    `json:"id:omitempty"`
	ModifyDate             *time.Time              `json:"modifyDate:omitempty"`
	Nda                    *string                 `json:"nda:omitempty"`
	Notes                  *string                 `json:"notes:omitempty"`
	Report                 *string                 `json:"report:omitempty"`
	ReportType             *Compliance_Report_Type `json:"reportType:omitempty"`
	RequestKey             *string                 `json:"requestKey:omitempty"`
	Status                 *string                 `json:"status:omitempty"`
	Ticket                 *Ticket                 `json:"ticket:omitempty"`
	TicketId               *int                    `json:"ticketId:omitempty"`
	User                   *User_Customer          `json:"user:omitempty"`
	UsrRecordId            *int                    `json:"usrRecordId:omitempty"`
}

type Account_Rwhois_Handle struct {
	Entity

	Account    *Account   `json:"account:omitempty"`
	AccountId  *int       `json:"accountId:omitempty"`
	CreateDate *time.Time `json:"createDate:omitempty"`
	Handle     *string    `json:"handle:omitempty"`
	Id         *int       `json:"id:omitempty"`
	ModifyDate *time.Time `json:"modifyDate:omitempty"`
}

type Account_Shipment struct {
	Entity

	Account              *Account                         `json:"account:omitempty"`
	AccountId            *int                             `json:"accountId:omitempty"`
	Courier              *Auxiliary_Shipping_Courier      `json:"courier:omitempty"`
	CourierId            *int                             `json:"courierId:omitempty"`
	CourierName          *string                          `json:"courierName:omitempty"`
	CreateEmployee       *User_Employee                   `json:"createEmployee:omitempty"`
	CreateUser           *User_Customer                   `json:"createUser:omitempty"`
	CreateUserId         *int                             `json:"createUserId:omitempty"`
	DestinationAddress   *Account_Address                 `json:"destinationAddress:omitempty"`
	DestinationAddressId *int                             `json:"destinationAddressId:omitempty"`
	DestinationDate      *time.Time                       `json:"destinationDate:omitempty"`
	Id                   *int                             `json:"id:omitempty"`
	ModifyEmployee       *User_Employee                   `json:"modifyEmployee:omitempty"`
	ModifyUser           *User_Customer                   `json:"modifyUser:omitempty"`
	ModifyUserId         *int                             `json:"modifyUserId:omitempty"`
	Note                 *string                          `json:"note:omitempty"`
	OriginationAddress   *Account_Address                 `json:"originationAddress:omitempty"`
	OriginationAddressId *int                             `json:"originationAddressId:omitempty"`
	OriginationDate      *time.Time                       `json:"originationDate:omitempty"`
	ShipmentItemCount    *uint                            `json:"shipmentItemCount:omitempty"`
	ShipmentItems        []Account_Shipment_Item          `json:"shipmentItems:omitempty"`
	Status               *Account_Shipment_Status         `json:"status:omitempty"`
	StatusId             *int                             `json:"statusId:omitempty"`
	TrackingData         []Account_Shipment_Tracking_Data `json:"trackingData:omitempty"`
	TrackingDataCount    *uint                            `json:"trackingDataCount:omitempty"`
	Type                 *Account_Shipment_Type           `json:"type:omitempty"`
	TypeId               *int                             `json:"typeId:omitempty"`
}

type Account_Shipment_Item struct {
	Entity

	CreateDate         *time.Time                  `json:"createDate:omitempty"`
	Description        *string                     `json:"description:omitempty"`
	Id                 *int                        `json:"id:omitempty"`
	PackageId          *int                        `json:"packageId:omitempty"`
	Shipment           *Account_Shipment           `json:"shipment:omitempty"`
	ShipmentId         *int                        `json:"shipmentId:omitempty"`
	ShipmentItemId     *int                        `json:"shipmentItemId:omitempty"`
	ShipmentItemType   *Account_Shipment_Item_Type `json:"shipmentItemType:omitempty"`
	ShipmentItemTypeId *int                        `json:"shipmentItemTypeId:omitempty"`
}

type Account_Shipment_Item_Type struct {
	Entity

	CreateDate *time.Time `json:"createDate:omitempty"`
	Id         *int       `json:"id:omitempty"`
	KeyName    *string    `json:"keyName:omitempty"`
	Name       *string    `json:"name:omitempty"`
}

type Account_Shipment_Resource_Type struct {
	Entity
}

type Account_Shipment_Status struct {
	Entity

	CreateDate *time.Time `json:"createDate:omitempty"`
	Id         *int       `json:"id:omitempty"`
	KeyName    *string    `json:"keyName:omitempty"`
	Name       *string    `json:"name:omitempty"`
}

type Account_Shipment_Tracking_Data struct {
	Entity

	CreateEmployee *User_Employee    `json:"createEmployee:omitempty"`
	CreateUser     *User_Customer    `json:"createUser:omitempty"`
	CreateUserId   *int              `json:"createUserId:omitempty"`
	Id             *int              `json:"id:omitempty"`
	ModifyEmployee *User_Employee    `json:"modifyEmployee:omitempty"`
	ModifyUser     *User_Customer    `json:"modifyUser:omitempty"`
	ModifyUserId   *int              `json:"modifyUserId:omitempty"`
	PackageId      *int              `json:"packageId:omitempty"`
	Sequence       *int              `json:"sequence:omitempty"`
	Shipment       *Account_Shipment `json:"shipment:omitempty"`
	ShipmentId     *int              `json:"shipmentId:omitempty"`
	TrackingData   *string           `json:"trackingData:omitempty"`
}

type Account_Shipment_Type struct {
	Entity

	CreateDate  *time.Time `json:"createDate:omitempty"`
	Description *string    `json:"description:omitempty"`
	Id          *int       `json:"id:omitempty"`
	KeyName     *string    `json:"keyName:omitempty"`
	Name        *string    `json:"name:omitempty"`
}

type Account_Status struct {
	Entity

	Id   *int    `json:"id:omitempty"`
	Name *string `json:"name:omitempty"`
}

type Auxiliary_Marketing_Event struct {
	Entity

	CreateDate  *time.Time `json:"createDate:omitempty"`
	EnabledFlag *int       `json:"enabledFlag:omitempty"`
	EndDate     *time.Time `json:"endDate:omitempty"`
	Location    *string    `json:"location:omitempty"`
	ModifyDate  *time.Time `json:"modifyDate:omitempty"`
	StartDate   *time.Time `json:"startDate:omitempty"`
	Title       *string    `json:"title:omitempty"`
	Url         *string    `json:"url:omitempty"`
}

type Auxiliary_Network_Status struct {
	Entity
}

type Auxiliary_Notification_Emergency struct {
	Entity

	CreateDate       *time.Time                                  `json:"createDate:omitempty"`
	Device           *string                                     `json:"device:omitempty"`
	Duration         *string                                     `json:"duration:omitempty"`
	Id               *int                                        `json:"id:omitempty"`
	Location         *string                                     `json:"location:omitempty"`
	Message          *string                                     `json:"message:omitempty"`
	ModifyDate       *time.Time                                  `json:"modifyDate:omitempty"`
	ServicesAffected *string                                     `json:"servicesAffected:omitempty"`
	Signature        *Auxiliary_Notification_Emergency_Signature `json:"signature:omitempty"`
	StartDate        *time.Time                                  `json:"startDate:omitempty"`
	Status           *Auxiliary_Notification_Emergency_Status    `json:"status:omitempty"`
	StatusId         *int                                        `json:"statusId:omitempty"`
}

type Auxiliary_Notification_Emergency_Signature struct {
	Entity

	Name *string `json:"name:omitempty"`
}

type Auxiliary_Notification_Emergency_Status struct {
	Entity

	Name *string `json:"name:omitempty"`
}

type Auxiliary_Press_Release struct {
	Entity

	About                []Auxiliary_Press_Release_About_Press_Release         `json:"about:omitempty"`
	AboutCount           *uint                                                 `json:"aboutCount:omitempty"`
	ContactCount         *uint                                                 `json:"contactCount:omitempty"`
	Contacts             []Auxiliary_Press_Release_Contact_Press_Release       `json:"contacts:omitempty"`
	Id                   *int                                                  `json:"id:omitempty"`
	MediaPartnerCount    *uint                                                 `json:"mediaPartnerCount:omitempty"`
	MediaPartners        []Auxiliary_Press_Release_Media_Partner_Press_Release `json:"mediaPartners:omitempty"`
	PressReleaseContent  *Auxiliary_Press_Release_Content                      `json:"pressReleaseContent:omitempty"`
	PublishDate          *time.Time                                            `json:"publishDate:omitempty"`
	ReleaseLocation      *string                                               `json:"releaseLocation:omitempty"`
	SubTitle             *string                                               `json:"subTitle:omitempty"`
	Title                *string                                               `json:"title:omitempty"`
	WebsiteHighlightFlag *bool                                                 `json:"websiteHighlightFlag:omitempty"`
}

type Auxiliary_Press_Release_About struct {
	Entity

	Content *string `json:"content:omitempty"`
	Id      *int    `json:"id:omitempty"`
	Title   *string `json:"title:omitempty"`
}

type Auxiliary_Press_Release_About_Press_Release struct {
	Entity

	AboutParagraphCount *uint                           `json:"aboutParagraphCount:omitempty"`
	AboutParagraphs     []Auxiliary_Press_Release_About `json:"aboutParagraphs:omitempty"`
	Id                  *int                            `json:"id:omitempty"`
	PressReleaseAboutId *int                            `json:"pressReleaseAboutId:omitempty"`
	PressReleaseCount   *uint                           `json:"pressReleaseCount:omitempty"`
	PressReleaseId      *int                            `json:"pressReleaseId:omitempty"`
	PressReleases       []Auxiliary_Press_Release       `json:"pressReleases:omitempty"`
	SortOrder           *int                            `json:"sortOrder:omitempty"`
}

type Auxiliary_Press_Release_Contact struct {
	Entity

	Email             *string `json:"email:omitempty"`
	FirstName         *string `json:"firstName:omitempty"`
	Id                *int    `json:"id:omitempty"`
	LastName          *string `json:"lastName:omitempty"`
	Phone             *string `json:"phone:omitempty"`
	ProfessionalTitle *string `json:"professionalTitle:omitempty"`
}

type Auxiliary_Press_Release_Contact_Press_Release struct {
	Entity

	ContactCount          *uint                             `json:"contactCount:omitempty"`
	Contacts              []Auxiliary_Press_Release_Contact `json:"contacts:omitempty"`
	Id                    *int                              `json:"id:omitempty"`
	PressReleaseContactId *int                              `json:"pressReleaseContactId:omitempty"`
	PressReleaseCount     *uint                             `json:"pressReleaseCount:omitempty"`
	PressReleaseId        *int                              `json:"pressReleaseId:omitempty"`
	PressReleases         []Auxiliary_Press_Release         `json:"pressReleases:omitempty"`
	SortOrder             *int                              `json:"sortOrder:omitempty"`
}

type Auxiliary_Press_Release_Content struct {
	Entity

	Id             *int    `json:"id:omitempty"`
	PressReleaseId *int    `json:"pressReleaseId:omitempty"`
	Text           *string `json:"text:omitempty"`
}

type Auxiliary_Press_Release_Media_Partner struct {
	Entity

	Id   *int    `json:"id:omitempty"`
	Name *string `json:"name:omitempty"`
}

type Auxiliary_Press_Release_Media_Partner_Press_Release struct {
	Entity

	Id                *int                                    `json:"id:omitempty"`
	MediaPartnerCount *uint                                   `json:"mediaPartnerCount:omitempty"`
	MediaPartnerId    *int                                    `json:"mediaPartnerId:omitempty"`
	MediaPartners     []Auxiliary_Press_Release_Media_Partner `json:"mediaPartners:omitempty"`
	PressReleaseCount *uint                                   `json:"pressReleaseCount:omitempty"`
	PressReleaseId    *int                                    `json:"pressReleaseId:omitempty"`
	PressReleases     []Auxiliary_Press_Release               `json:"pressReleases:omitempty"`
}

type Auxiliary_Shipping_Courier struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
	Url     *string `json:"url:omitempty"`
}

type Auxiliary_Shipping_Courier_Type struct {
	Entity

	Courier      []Auxiliary_Shipping_Courier `json:"courier:omitempty"`
	CourierCount *uint                        `json:"courierCount:omitempty"`
	Description  *string                      `json:"description:omitempty"`
	Id           *int                         `json:"id:omitempty"`
	KeyName      *string                      `json:"keyName:omitempty"`
	Name         *string                      `json:"name:omitempty"`
}

type Billing_Currency struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Billing_Currency_ExchangeRate struct {
	Entity

	EffectiveDate   *time.Time        `json:"effectiveDate:omitempty"`
	ExpirationDate  *time.Time        `json:"expirationDate:omitempty"`
	FundingCurrency *Billing_Currency `json:"fundingCurrency:omitempty"`
	Id              *int              `json:"id:omitempty"`
	LocalCurrency   *Billing_Currency `json:"localCurrency:omitempty"`
	Rate            *float64          `json:"rate:omitempty"`
}

type Billing_Info struct {
	Entity

	Account                   *Account            `json:"account:omitempty"`
	AccountId                 *int                `json:"accountId:omitempty"`
	AchInformation            []Billing_Info_Ach  `json:"achInformation:omitempty"`
	AchInformationCount       *uint               `json:"achInformationCount:omitempty"`
	AnniversaryDayOfMonth     *int                `json:"anniversaryDayOfMonth:omitempty"`
	CardAccountNumber         *string             `json:"cardAccountNumber:omitempty"`
	CardExpirationMonth       *int                `json:"cardExpirationMonth:omitempty"`
	CardExpirationYear        *int                `json:"cardExpirationYear:omitempty"`
	CardNickname              *string             `json:"cardNickname:omitempty"`
	CardType                  *string             `json:"cardType:omitempty"`
	CardVerificationNumber    *string             `json:"cardVerificationNumber:omitempty"`
	CreateDate                *time.Time          `json:"createDate:omitempty"`
	Currency                  *Billing_Currency   `json:"currency:omitempty"`
	CurrentBillingCycle       *Billing_Info_Cycle `json:"currentBillingCycle:omitempty"`
	Id                        *int                `json:"id:omitempty"`
	LastBillDate              *time.Time          `json:"lastBillDate:omitempty"`
	LastFourPaymentCardDigits *int                `json:"lastFourPaymentCardDigits:omitempty"`
	LastPaymentDate           *time.Time          `json:"lastPaymentDate:omitempty"`
	ModifyDate                *time.Time          `json:"modifyDate:omitempty"`
	NextBillDate              *time.Time          `json:"nextBillDate:omitempty"`
	PaymentTerms              *int                `json:"paymentTerms:omitempty"`
	PercentDiscountOnetime    *int                `json:"percentDiscountOnetime:omitempty"`
	PercentDiscountRecurring  *int                `json:"percentDiscountRecurring:omitempty"`
	SparePoolAmount           *int                `json:"sparePoolAmount:omitempty"`
	VatId                     *string             `json:"vatId:omitempty"`
}

type Billing_Info_Ach struct {
	Entity

	Account           *Account   `json:"account:omitempty"`
	AccountId         *int       `json:"accountId:omitempty"`
	AccountNumber     *string    `json:"accountNumber:omitempty"`
	AccountType       *string    `json:"accountType:omitempty"`
	BankTransitNumber *string    `json:"bankTransitNumber:omitempty"`
	City              *string    `json:"city:omitempty"`
	Country           *string    `json:"country:omitempty"`
	FirstName         *string    `json:"firstName:omitempty"`
	Id                *int       `json:"id:omitempty"`
	LastName          *string    `json:"lastName:omitempty"`
	PhoneNumber       *string    `json:"phoneNumber:omitempty"`
	Postalcode        *string    `json:"postalcode:omitempty"`
	State             *string    `json:"state:omitempty"`
	Status            *string    `json:"status:omitempty"`
	Street1           *string    `json:"street1:omitempty"`
	Street2           *string    `json:"street2:omitempty"`
	VerifiedDate      *time.Time `json:"verifiedDate:omitempty"`
}

type Billing_Info_Cycle struct {
	Entity

	Account                *Account   `json:"account:omitempty"`
	CurrentCycleEndDate    *time.Time `json:"currentCycleEndDate:omitempty"`
	CurrentCycleStartDate  *time.Time `json:"currentCycleStartDate:omitempty"`
	NextCycleStartDate     *time.Time `json:"nextCycleStartDate:omitempty"`
	PreviousCycleEndDate   *time.Time `json:"previousCycleEndDate:omitempty"`
	PreviousCycleStartDate *time.Time `json:"previousCycleStartDate:omitempty"`
}

type Billing_Invoice struct {
	Entity

	Account                        *Account                             `json:"account:omitempty"`
	AccountId                      *int                                 `json:"accountId:omitempty"`
	Address1                       *string                              `json:"address1:omitempty"`
	Address2                       *string                              `json:"address2:omitempty"`
	Amount                         *float64                             `json:"amount:omitempty"`
	BrandAtInvoiceCreation         *Brand                               `json:"brandAtInvoiceCreation:omitempty"`
	City                           *string                              `json:"city:omitempty"`
	ClaimedTaxExemptTxFlag         *bool                                `json:"claimedTaxExemptTxFlag:omitempty"`
	ClosedDate                     *time.Time                           `json:"closedDate:omitempty"`
	CompanyName                    *string                              `json:"companyName:omitempty"`
	Country                        *string                              `json:"country:omitempty"`
	CreateDate                     *time.Time                           `json:"createDate:omitempty"`
	DetailedPdfGeneratedFlag       *bool                                `json:"detailedPdfGeneratedFlag:omitempty"`
	DocumentsGeneratedFlag         *bool                                `json:"documentsGeneratedFlag:omitempty"`
	Email                          *string                              `json:"email:omitempty"`
	EndingBalance                  *float64                             `json:"endingBalance:omitempty"`
	FaxPhone                       *string                              `json:"faxPhone:omitempty"`
	FirstName                      *string                              `json:"firstName:omitempty"`
	Id                             *int                                 `json:"id:omitempty"`
	InvoiceTopLevelItemCount       *uint                                `json:"invoiceTopLevelItemCount:omitempty"`
	InvoiceTopLevelItems           []Billing_Invoice_Item               `json:"invoiceTopLevelItems:omitempty"`
	InvoiceTotalAmount             *float64                             `json:"invoiceTotalAmount:omitempty"`
	InvoiceTotalOneTimeAmount      *float64                             `json:"invoiceTotalOneTimeAmount:omitempty"`
	InvoiceTotalOneTimeTaxAmount   *float64                             `json:"invoiceTotalOneTimeTaxAmount:omitempty"`
	InvoiceTotalPreTaxAmount       *float64                             `json:"invoiceTotalPreTaxAmount:omitempty"`
	InvoiceTotalRecurringAmount    *float64                             `json:"invoiceTotalRecurringAmount:omitempty"`
	InvoiceTotalRecurringTaxAmount *float64                             `json:"invoiceTotalRecurringTaxAmount:omitempty"`
	ItemCount                      *uint                                `json:"itemCount:omitempty"`
	Items                          []Billing_Invoice_Item               `json:"items:omitempty"`
	LastName                       *string                              `json:"lastName:omitempty"`
	ModifyDate                     *time.Time                           `json:"modifyDate:omitempty"`
	OfficePhone                    *string                              `json:"officePhone:omitempty"`
	Payment                        *float64                             `json:"payment:omitempty"`
	PaymentCount                   *uint                                `json:"paymentCount:omitempty"`
	Payments                       []Billing_Invoice_Receivable_Payment `json:"payments:omitempty"`
	PostalCode                     *string                              `json:"postalCode:omitempty"`
	PurchaseOrderNumber            *string                              `json:"purchaseOrderNumber:omitempty"`
	SellerRegistration             *string                              `json:"sellerRegistration:omitempty"`
	StartingBalance                *float64                             `json:"startingBalance:omitempty"`
	State                          *string                              `json:"state:omitempty"`
	StatusCode                     *string                              `json:"statusCode:omitempty"`
	TaxInfo                        *Billing_Invoice_Tax_Info            `json:"taxInfo:omitempty"`
	TaxInfoHistory                 []Billing_Invoice_Tax_Info           `json:"taxInfoHistory:omitempty"`
	TaxInfoHistoryCount            *uint                                `json:"taxInfoHistoryCount:omitempty"`
	TaxMessage                     *string                              `json:"taxMessage:omitempty"`
	TaxStatusId                    *int                                 `json:"taxStatusId:omitempty"`
	TaxType                        *Billing_Invoice_Tax_Type            `json:"taxType:omitempty"`
	TaxTypeId                      *int                                 `json:"taxTypeId:omitempty"`
	TypeCode                       *string                              `json:"typeCode:omitempty"`
}

type Billing_Invoice_Item struct {
	Entity

	AssociatedChildren              []Billing_Invoice_Item `json:"associatedChildren:omitempty"`
	AssociatedChildrenCount         *uint                  `json:"associatedChildrenCount:omitempty"`
	AssociatedInvoiceItem           *Billing_Invoice_Item  `json:"associatedInvoiceItem:omitempty"`
	AssociatedInvoiceItemId         *int                   `json:"associatedInvoiceItemId:omitempty"`
	BillingItem                     *Billing_Item          `json:"billingItem:omitempty"`
	BillingItemId                   *int                   `json:"billingItemId:omitempty"`
	Category                        *Product_Item_Category `json:"category:omitempty"`
	CategoryCode                    *string                `json:"categoryCode:omitempty"`
	Children                        []Billing_Invoice_Item `json:"children:omitempty"`
	ChildrenCount                   *uint                  `json:"childrenCount:omitempty"`
	CreateDate                      *time.Time             `json:"createDate:omitempty"`
	Description                     *string                `json:"description:omitempty"`
	DomainName                      *string                `json:"domainName:omitempty"`
	FilteredAssociatedChildren      []Billing_Invoice_Item `json:"filteredAssociatedChildren:omitempty"`
	FilteredAssociatedChildrenCount *uint                  `json:"filteredAssociatedChildrenCount:omitempty"`
	HostName                        *string                `json:"hostName:omitempty"`
	HourlyRecurringFee              *float64               `json:"hourlyRecurringFee:omitempty"`
	Id                              *int                   `json:"id:omitempty"`
	Invoice                         *Billing_Invoice       `json:"invoice:omitempty"`
	InvoiceId                       *int                   `json:"invoiceId:omitempty"`
	LaborAfterTaxAmount             *float64               `json:"laborAfterTaxAmount:omitempty"`
	LaborFee                        *float64               `json:"laborFee:omitempty"`
	LaborFeeTaxRate                 *float64               `json:"laborFeeTaxRate:omitempty"`
	LaborTaxAmount                  *float64               `json:"laborTaxAmount:omitempty"`
	Location                        *Location              `json:"location:omitempty"`
	NonZeroAssociatedChildren       []Billing_Invoice_Item `json:"nonZeroAssociatedChildren:omitempty"`
	NonZeroAssociatedChildrenCount  *uint                  `json:"nonZeroAssociatedChildrenCount:omitempty"`
	Notes                           *string                `json:"notes:omitempty"`
	OneTimeAfterTaxAmount           *float64               `json:"oneTimeAfterTaxAmount:omitempty"`
	OneTimeFee                      *float64               `json:"oneTimeFee:omitempty"`
	OneTimeFeeTaxRate               *float64               `json:"oneTimeFeeTaxRate:omitempty"`
	OneTimeTaxAmount                *float64               `json:"oneTimeTaxAmount:omitempty"`
	Parent                          *Billing_Invoice_Item  `json:"parent:omitempty"`
	ParentId                        *int                   `json:"parentId:omitempty"`
	Product                         *Product_Item          `json:"product:omitempty"`
	ProductItemId                   *int                   `json:"productItemId:omitempty"`
	RecurringAfterTaxAmount         *float64               `json:"recurringAfterTaxAmount:omitempty"`
	RecurringFee                    *float64               `json:"recurringFee:omitempty"`
	RecurringFeeTaxRate             *float64               `json:"recurringFeeTaxRate:omitempty"`
	RecurringTaxAmount              *float64               `json:"recurringTaxAmount:omitempty"`
	ResourceTableId                 *int                   `json:"resourceTableId:omitempty"`
	SetupAfterTaxAmount             *float64               `json:"setupAfterTaxAmount:omitempty"`
	SetupFee                        *float64               `json:"setupFee:omitempty"`
	SetupFeeTaxRate                 *float64               `json:"setupFeeTaxRate:omitempty"`
	SetupTaxAmount                  *float64               `json:"setupTaxAmount:omitempty"`
	TotalOneTimeAmount              *float64               `json:"totalOneTimeAmount:omitempty"`
	TotalOneTimeTaxAmount           *float64               `json:"totalOneTimeTaxAmount:omitempty"`
	TotalRecurringAmount            *float64               `json:"totalRecurringAmount:omitempty"`
	TotalRecurringTaxAmount         *float64               `json:"totalRecurringTaxAmount:omitempty"`
}

type Billing_Invoice_Item_Hardware struct {
	Billing_Invoice_Item

	Resource *Hardware `json:"resource:omitempty"`
}

type Billing_Invoice_Item_Tax_Info struct {
	Entity

	CreateDate          *time.Time                `json:"createDate:omitempty"`
	Description         *string                   `json:"description:omitempty"`
	EffectiveTaxRate    *float64                  `json:"effectiveTaxRate:omitempty"`
	ExemptAmount        *float64                  `json:"exemptAmount:omitempty"`
	FeeProperty         *string                   `json:"feeProperty:omitempty"`
	Id                  *int                      `json:"id:omitempty"`
	InvoiceItem         *Billing_Invoice_Item     `json:"invoiceItem:omitempty"`
	InvoiceItemId       *int                      `json:"invoiceItemId:omitempty"`
	InvoiceTaxInfo      *Billing_Invoice_Tax_Info `json:"invoiceTaxInfo:omitempty"`
	InvoiceTaxInfoId    *int                      `json:"invoiceTaxInfoId:omitempty"`
	ModifyDate          *time.Time                `json:"modifyDate:omitempty"`
	NonTaxableBasis     *float64                  `json:"nonTaxableBasis:omitempty"`
	ReportedFlag        *bool                     `json:"reportedFlag:omitempty"`
	SellerRegistration  *string                   `json:"sellerRegistration:omitempty"`
	TaxAmount           *float64                  `json:"taxAmount:omitempty"`
	TaxAmountToCurrency *float64                  `json:"taxAmountToCurrency:omitempty"`
	TaxRate             *float64                  `json:"taxRate:omitempty"`
	TaxableBasis        *float64                  `json:"taxableBasis:omitempty"`
	ToCurrency          *Billing_Currency         `json:"toCurrency:omitempty"`
	ToCurrencyId        *int                      `json:"toCurrencyId:omitempty"`
}

type Billing_Invoice_Next struct {
	Entity
}

type Billing_Invoice_Receivable_Payment struct {
	Entity

	Account                  *Account                            `json:"account:omitempty"`
	Amount                   *float64                            `json:"amount:omitempty"`
	CreateDate               *time.Time                          `json:"createDate:omitempty"`
	CreditCardLastFourDigits *int                                `json:"creditCardLastFourDigits:omitempty"`
	CreditCardRequestId      *string                             `json:"creditCardRequestId:omitempty"`
	CreditCardTransaction    *Billing_Payment_Card_Transaction   `json:"creditCardTransaction:omitempty"`
	ExchangeRate             *Billing_Currency_ExchangeRate      `json:"exchangeRate:omitempty"`
	Invoice                  *Billing_Invoice                    `json:"invoice:omitempty"`
	InvoiceId                *int                                `json:"invoiceId:omitempty"`
	PaypalTransaction        *Billing_Payment_PayPal_Transaction `json:"paypalTransaction:omitempty"`
	TypeCode                 *string                             `json:"typeCode:omitempty"`
}

type Billing_Invoice_Tax_Info struct {
	Entity

	CreateDate               *time.Time                      `json:"createDate:omitempty"`
	Currency                 *Billing_Currency               `json:"currency:omitempty"`
	CurrencyId               *int                            `json:"currencyId:omitempty"`
	FunctionalCurrency       *Billing_Currency               `json:"functionalCurrency:omitempty"`
	Id                       *int                            `json:"id:omitempty"`
	Invoice                  *Billing_Invoice                `json:"invoice:omitempty"`
	InvoiceId                *int                            `json:"invoiceId:omitempty"`
	ItemCount                *uint                           `json:"itemCount:omitempty"`
	ItemWithCurrencyInfo     *Billing_Invoice_Item_Tax_Info  `json:"itemWithCurrencyInfo:omitempty"`
	Items                    []Billing_Invoice_Item_Tax_Info `json:"items:omitempty"`
	ModifyDate               *time.Time                      `json:"modifyDate:omitempty"`
	ReportedFlag             *bool                           `json:"reportedFlag:omitempty"`
	TotalTaxAmountToCurrency *float64                        `json:"totalTaxAmountToCurrency:omitempty"`
}

type Billing_Invoice_Tax_Status struct {
	Entity

	CreateDate *time.Time `json:"createDate:omitempty"`
	Id         *int       `json:"id:omitempty"`
	KeyName    *string    `json:"keyName:omitempty"`
	ModifyDate *time.Time `json:"modifyDate:omitempty"`
	Name       *string    `json:"name:omitempty"`
}

type Billing_Invoice_Tax_Type struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Billing_Item struct {
	Entity

	Account                                            *Account                                `json:"account:omitempty"`
	ActiveAgreement                                    *Account_Agreement                      `json:"activeAgreement:omitempty"`
	ActiveAgreementFlag                                *Account_Agreement                      `json:"activeAgreementFlag:omitempty"`
	ActiveAssociatedChildren                           []Billing_Item                          `json:"activeAssociatedChildren:omitempty"`
	ActiveAssociatedChildrenCount                      *uint                                   `json:"activeAssociatedChildrenCount:omitempty"`
	ActiveAssociatedGuestDiskBillingItemCount          *uint                                   `json:"activeAssociatedGuestDiskBillingItemCount:omitempty"`
	ActiveAssociatedGuestDiskBillingItems              []Billing_Item                          `json:"activeAssociatedGuestDiskBillingItems:omitempty"`
	ActiveBundledItemCount                             *uint                                   `json:"activeBundledItemCount:omitempty"`
	ActiveBundledItems                                 []Billing_Item                          `json:"activeBundledItems:omitempty"`
	ActiveCancellationItem                             *Billing_Item_Cancellation_Request_Item `json:"activeCancellationItem:omitempty"`
	ActiveChildren                                     []Billing_Item                          `json:"activeChildren:omitempty"`
	ActiveChildrenCount                                *uint                                   `json:"activeChildrenCount:omitempty"`
	ActiveFlag                                         *bool                                   `json:"activeFlag:omitempty"`
	ActiveSparePoolAssociatedGuestDiskBillingItemCount *uint                                   `json:"activeSparePoolAssociatedGuestDiskBillingItemCount:omitempty"`
	ActiveSparePoolAssociatedGuestDiskBillingItems     []Billing_Item                          `json:"activeSparePoolAssociatedGuestDiskBillingItems:omitempty"`
	ActiveSparePoolBundledItemCount                    *uint                                   `json:"activeSparePoolBundledItemCount:omitempty"`
	ActiveSparePoolBundledItems                        []Billing_Item                          `json:"activeSparePoolBundledItems:omitempty"`
	AllowCancellationFlag                              *int                                    `json:"allowCancellationFlag:omitempty"`
	AssociatedBillingItem                              *Billing_Item                           `json:"associatedBillingItem:omitempty"`
	AssociatedBillingItemHistory                       []Billing_Item_Association_History      `json:"associatedBillingItemHistory:omitempty"`
	AssociatedBillingItemHistoryCount                  *uint                                   `json:"associatedBillingItemHistoryCount:omitempty"`
	AssociatedBillingItemId                            *string                                 `json:"associatedBillingItemId:omitempty"`
	AssociatedChildren                                 []Billing_Item                          `json:"associatedChildren:omitempty"`
	AssociatedChildrenCount                            *uint                                   `json:"associatedChildrenCount:omitempty"`
	AssociatedParent                                   []Billing_Item                          `json:"associatedParent:omitempty"`
	AssociatedParentCount                              *uint                                   `json:"associatedParentCount:omitempty"`
	AvailableMatchingVlanCount                         *uint                                   `json:"availableMatchingVlanCount:omitempty"`
	AvailableMatchingVlans                             []Network_Vlan                          `json:"availableMatchingVlans:omitempty"`
	BandwidthAllocation                                *Network_Bandwidth_Version1_Allocation  `json:"bandwidthAllocation:omitempty"`
	BillableChildren                                   []Billing_Item                          `json:"billableChildren:omitempty"`
	BillableChildrenCount                              *uint                                   `json:"billableChildrenCount:omitempty"`
	BundleItemCount                                    *uint                                   `json:"bundleItemCount:omitempty"`
	BundleItems                                        []Product_Item_Bundles                  `json:"bundleItems:omitempty"`
	BundledItemCount                                   *uint                                   `json:"bundledItemCount:omitempty"`
	BundledItems                                       []Billing_Item                          `json:"bundledItems:omitempty"`
	CanceledChildren                                   []Billing_Item                          `json:"canceledChildren:omitempty"`
	CanceledChildrenCount                              *uint                                   `json:"canceledChildrenCount:omitempty"`
	CancellationDate                                   *time.Time                              `json:"cancellationDate:omitempty"`
	CancellationReason                                 *Billing_Item_Cancellation_Reason       `json:"cancellationReason:omitempty"`
	CancellationRequestCount                           *uint                                   `json:"cancellationRequestCount:omitempty"`
	CancellationRequests                               []Billing_Item_Cancellation_Request     `json:"cancellationRequests:omitempty"`
	Category                                           *Product_Item_Category                  `json:"category:omitempty"`
	CategoryCode                                       *string                                 `json:"categoryCode:omitempty"`
	Children                                           []Billing_Item                          `json:"children:omitempty"`
	ChildrenCount                                      *uint                                   `json:"childrenCount:omitempty"`
	ChildrenWithActiveAgreement                        []Billing_Item                          `json:"childrenWithActiveAgreement:omitempty"`
	ChildrenWithActiveAgreementCount                   *uint                                   `json:"childrenWithActiveAgreementCount:omitempty"`
	CreateDate                                         *time.Time                              `json:"createDate:omitempty"`
	CurrentHourlyCharge                                *string                                 `json:"currentHourlyCharge:omitempty"`
	CycleStartDate                                     *time.Time                              `json:"cycleStartDate:omitempty"`
	Description                                        *string                                 `json:"description:omitempty"`
	DomainName                                         *string                                 `json:"domainName:omitempty"`
	DowngradeItemCount                                 *uint                                   `json:"downgradeItemCount:omitempty"`
	DowngradeItems                                     []Product_Item                          `json:"downgradeItems:omitempty"`
	FilteredNextInvoiceChildren                        []Billing_Item                          `json:"filteredNextInvoiceChildren:omitempty"`
	FilteredNextInvoiceChildrenCount                   *uint                                   `json:"filteredNextInvoiceChildrenCount:omitempty"`
	HostName                                           *string                                 `json:"hostName:omitempty"`
	HourlyFlag                                         *bool                                   `json:"hourlyFlag:omitempty"`
	HourlyRecurringFee                                 *float64                                `json:"hourlyRecurringFee:omitempty"`
	HoursUsed                                          *string                                 `json:"hoursUsed:omitempty"`
	Id                                                 *int                                    `json:"id:omitempty"`
	InvoiceItem                                        *Billing_Invoice_Item                   `json:"invoiceItem:omitempty"`
	InvoiceItemCount                                   *uint                                   `json:"invoiceItemCount:omitempty"`
	InvoiceItems                                       []Billing_Invoice_Item                  `json:"invoiceItems:omitempty"`
	Item                                               *Product_Item                           `json:"item:omitempty"`
	LaborFee                                           *float64                                `json:"laborFee:omitempty"`
	LaborFeeTaxRate                                    *float64                                `json:"laborFeeTaxRate:omitempty"`
	LastBillDate                                       *time.Time                              `json:"lastBillDate:omitempty"`
	Location                                           *Location                               `json:"location:omitempty"`
	ModifyDate                                         *time.Time                              `json:"modifyDate:omitempty"`
	NextBillDate                                       *time.Time                              `json:"nextBillDate:omitempty"`
	NextInvoiceChildren                                []Billing_Item                          `json:"nextInvoiceChildren:omitempty"`
	NextInvoiceChildrenCount                           *uint                                   `json:"nextInvoiceChildrenCount:omitempty"`
	NextInvoiceTotalOneTimeAmount                      *float64                                `json:"nextInvoiceTotalOneTimeAmount:omitempty"`
	NextInvoiceTotalOneTimeTaxAmount                   *float64                                `json:"nextInvoiceTotalOneTimeTaxAmount:omitempty"`
	NextInvoiceTotalRecurringAmount                    *float64                                `json:"nextInvoiceTotalRecurringAmount:omitempty"`
	NextInvoiceTotalRecurringTaxAmount                 *float64                                `json:"nextInvoiceTotalRecurringTaxAmount:omitempty"`
	NonZeroNextInvoiceChildren                         []Billing_Item                          `json:"nonZeroNextInvoiceChildren:omitempty"`
	NonZeroNextInvoiceChildrenCount                    *uint                                   `json:"nonZeroNextInvoiceChildrenCount:omitempty"`
	Notes                                              *string                                 `json:"notes:omitempty"`
	OneTimeFee                                         *float64                                `json:"oneTimeFee:omitempty"`
	OneTimeFeeTaxRate                                  *float64                                `json:"oneTimeFeeTaxRate:omitempty"`
	OrderItem                                          *Billing_Order_Item                     `json:"orderItem:omitempty"`
	OrderItemId                                        *int                                    `json:"orderItemId:omitempty"`
	OriginalLocation                                   *Location                               `json:"originalLocation:omitempty"`
	Package                                            *Product_Package                        `json:"package:omitempty"`
	Parent                                             *Billing_Item                           `json:"parent:omitempty"`
	ParentId                                           *int                                    `json:"parentId:omitempty"`
	ParentVirtualGuestBillingItem                      *Billing_Item_Virtual_Guest             `json:"parentVirtualGuestBillingItem:omitempty"`
	PendingCancellationFlag                            *bool                                   `json:"pendingCancellationFlag:omitempty"`
	PendingOrderItem                                   *Billing_Order_Item                     `json:"pendingOrderItem:omitempty"`
	ProvisionTransaction                               *Provisioning_Version1_Transaction      `json:"provisionTransaction:omitempty"`
	RecurringFee                                       *float64                                `json:"recurringFee:omitempty"`
	RecurringFeeTaxRate                                *float64                                `json:"recurringFeeTaxRate:omitempty"`
	RecurringMonths                                    *int                                    `json:"recurringMonths:omitempty"`
	ServiceProviderId                                  *int                                    `json:"serviceProviderId:omitempty"`
	SetupFee                                           *float64                                `json:"setupFee:omitempty"`
	SetupFeeTaxRate                                    *float64                                `json:"setupFeeTaxRate:omitempty"`
	SoftwareDescription                                *Software_Description                   `json:"softwareDescription:omitempty"`
	UpgradeItem                                        *Product_Item                           `json:"upgradeItem:omitempty"`
	UpgradeItemCount                                   *uint                                   `json:"upgradeItemCount:omitempty"`
	UpgradeItems                                       []Product_Item                          `json:"upgradeItems:omitempty"`
}

type Billing_Item_Account_Media_Data_Transfer_Request struct {
	Billing_Item

	Resource *Account_Media_Data_Transfer_Request `json:"resource:omitempty"`
}

type Billing_Item_Association_History struct {
	Entity

	AssociatedBillingItem   *Billing_Item `json:"associatedBillingItem:omitempty"`
	AssociatedBillingItemId *int          `json:"associatedBillingItemId:omitempty"`
	BillingItem             *Billing_Item `json:"billingItem:omitempty"`
	BillingItemId           *int          `json:"billingItemId:omitempty"`
	CreateDate              *time.Time    `json:"createDate:omitempty"`
	Id                      *int          `json:"id:omitempty"`
}

type Billing_Item_Cancellation_Reason struct {
	Entity

	BillingCancelReasonCategoryId     *int                                       `json:"billingCancelReasonCategoryId:omitempty"`
	BillingCancellationReasonCategory *Billing_Item_Cancellation_Reason_Category `json:"billingCancellationReasonCategory:omitempty"`
	BillingItemCount                  *uint                                      `json:"billingItemCount:omitempty"`
	BillingItems                      []Billing_Item                             `json:"billingItems:omitempty"`
	Id                                *int                                       `json:"id:omitempty"`
	KeyName                           *string                                    `json:"keyName:omitempty"`
	Reason                            *string                                    `json:"reason:omitempty"`
	TranslatedReason                  *string                                    `json:"translatedReason:omitempty"`
}

type Billing_Item_Cancellation_Reason_Category struct {
	Entity

	BillingCancellationReasonCount *uint                              `json:"billingCancellationReasonCount:omitempty"`
	BillingCancellationReasons     []Billing_Item_Cancellation_Reason `json:"billingCancellationReasons:omitempty"`
	Id                             *int                               `json:"id:omitempty"`
	Name                           *string                            `json:"name:omitempty"`
}

type Billing_Item_Cancellation_Request struct {
	Entity

	Account               *Account                                  `json:"account:omitempty"`
	AccountId             *int                                      `json:"accountId:omitempty"`
	BillingCancelReasonId *int                                      `json:"billingCancelReasonId:omitempty"`
	CreateDate            *time.Time                                `json:"createDate:omitempty"`
	Id                    *int                                      `json:"id:omitempty"`
	ItemCount             *uint                                     `json:"itemCount:omitempty"`
	Items                 []Billing_Item_Cancellation_Request_Item  `json:"items:omitempty"`
	ModifyDate            *time.Time                                `json:"modifyDate:omitempty"`
	Notes                 *string                                   `json:"notes:omitempty"`
	Status                *Billing_Item_Cancellation_Request_Status `json:"status:omitempty"`
	StatusId              *int                                      `json:"statusId:omitempty"`
	Ticket                *Ticket                                   `json:"ticket:omitempty"`
	TicketId              *int                                      `json:"ticketId:omitempty"`
	User                  *User_Customer                            `json:"user:omitempty"`
}

type Billing_Item_Cancellation_Request_Item struct {
	Entity

	BillingItem               *Billing_Item                      `json:"billingItem:omitempty"`
	BillingItemId             *int                               `json:"billingItemId:omitempty"`
	CancellationRequest       *Billing_Item_Cancellation_Request `json:"cancellationRequest:omitempty"`
	CancellationRequestId     *int                               `json:"cancellationRequestId:omitempty"`
	Id                        *int                               `json:"id:omitempty"`
	ImmediateCancellationFlag *bool                              `json:"immediateCancellationFlag:omitempty"`
	ScheduledCancellationDate *time.Time                         `json:"scheduledCancellationDate:omitempty"`
	ServiceReclaimStatusCode  *string                            `json:"serviceReclaimStatusCode:omitempty"`
}

type Billing_Item_Cancellation_Request_Status struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Billing_Item_Ctc_Account struct {
	Billing_Item
}

type Billing_Item_Gateway_Appliance_Cluster struct {
	Billing_Item

	Resource *Resource_Group `json:"resource:omitempty"`
}

type Billing_Item_Hardware struct {
	Billing_Item

	BillingCycleBandwidthUsage             []Network_Bandwidth_Usage     `json:"billingCycleBandwidthUsage:omitempty"`
	BillingCycleBandwidthUsageCount        *uint                         `json:"billingCycleBandwidthUsageCount:omitempty"`
	BillingCyclePrivateBandwidthUsage      []Network_Bandwidth_Usage     `json:"billingCyclePrivateBandwidthUsage:omitempty"`
	BillingCyclePrivateBandwidthUsageCount *uint                         `json:"billingCyclePrivateBandwidthUsageCount:omitempty"`
	BillingCyclePrivateUsageIn             *float64                      `json:"billingCyclePrivateUsageIn:omitempty"`
	BillingCyclePrivateUsageOut            *float64                      `json:"billingCyclePrivateUsageOut:omitempty"`
	BillingCyclePrivateUsageTotal          *uint                         `json:"billingCyclePrivateUsageTotal:omitempty"`
	BillingCyclePublicBandwidthUsage       []Network_Bandwidth_Usage     `json:"billingCyclePublicBandwidthUsage:omitempty"`
	BillingCyclePublicBandwidthUsageCount  *uint                         `json:"billingCyclePublicBandwidthUsageCount:omitempty"`
	BillingCyclePublicUsageIn              *float64                      `json:"billingCyclePublicUsageIn:omitempty"`
	BillingCyclePublicUsageOut             *float64                      `json:"billingCyclePublicUsageOut:omitempty"`
	BillingCyclePublicUsageTotal           *uint                         `json:"billingCyclePublicUsageTotal:omitempty"`
	LockboxNetworkStorage                  *Billing_Item_Network_Storage `json:"lockboxNetworkStorage:omitempty"`
	MonitoringBillingItemCount             *uint                         `json:"monitoringBillingItemCount:omitempty"`
	MonitoringBillingItems                 []Billing_Item                `json:"monitoringBillingItems:omitempty"`
	Resource                               *Hardware_Server              `json:"resource:omitempty"`
	ResourceTableId                        *int                          `json:"resourceTableId:omitempty"`
}

type Billing_Item_Hardware_Colocation struct {
	Billing_Item_Hardware
}

type Billing_Item_Hardware_Component struct {
	Billing_Item

	Resource        []Hardware_Component `json:"resource:omitempty"`
	ResourceCount   *uint                `json:"resourceCount:omitempty"`
	ResourceTableId *int                 `json:"resourceTableId:omitempty"`
}

type Billing_Item_Hardware_Security_Module struct {
	Billing_Item_Hardware
}

type Billing_Item_Hardware_Server struct {
	Billing_Item_Hardware
}

type Billing_Item_Link_ThePlanet struct {
	Entity

	BillingItem     *Billing_Item     `json:"billingItem:omitempty"`
	ServiceProvider *Service_Provider `json:"serviceProvider:omitempty"`
}

type Billing_Item_Network_Application_Delivery_Controller struct {
	Billing_Item

	BandwidthAllotmentDetail *Network_Bandwidth_Version1_Allotment_Detail `json:"bandwidthAllotmentDetail:omitempty"`
	Resource                 *Network_Application_Delivery_Controller     `json:"resource:omitempty"`
}

type Billing_Item_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress struct {
	Billing_Item

	Resource *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"resource:omitempty"`
}

type Billing_Item_Network_Bandwidth struct {
	Billing_Item
}

type Billing_Item_Network_Firewall struct {
	Billing_Item

	Resource *Network_Component_Firewall `json:"resource:omitempty"`
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

	Resource *Network_LoadBalancer_Global_Account `json:"resource:omitempty"`
}

type Billing_Item_Network_LoadBalancer_VirtualIpAddress struct {
	Billing_Item

	Resource *Network_LoadBalancer_VirtualIpAddress `json:"resource:omitempty"`
}

type Billing_Item_Network_Message_Delivery struct {
	Billing_Item

	Resource *Network_Message_Delivery `json:"resource:omitempty"`
}

type Billing_Item_Network_Message_Queue struct {
	Billing_Item

	Resource *Network_Message_Queue `json:"resource:omitempty"`
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

	Resource *Network_Storage `json:"resource:omitempty"`
}

type Billing_Item_Network_Storage_Hub struct {
	Billing_Item_Network_Storage
}

type Billing_Item_Network_Storage_Hub_Bandwidth struct {
	Billing_Item_Network_Storage
}

type Billing_Item_Network_Subnet struct {
	Billing_Item

	Resource        *Network_Subnet `json:"resource:omitempty"`
	ResourceName    *string         `json:"resourceName:omitempty"`
	ResourceTableId *int            `json:"resourceTableId:omitempty"`
}

type Billing_Item_Network_Subnet_IpAddress_Global struct {
	Billing_Item_Network_Subnet
}

type Billing_Item_Network_Tunnel struct {
	Billing_Item

	Resource *Network_Tunnel_Module_Context `json:"resource:omitempty"`
}

type Billing_Item_Network_Vlan struct {
	Billing_Item

	Resource *Network_Vlan `json:"resource:omitempty"`
}

type Billing_Item_Software_Component struct {
	Billing_Item

	Resource        *Software_Component `json:"resource:omitempty"`
	ResourceTableId *int                `json:"resourceTableId:omitempty"`
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

	Resource *Software_Component `json:"resource:omitempty"`
}

type Billing_Item_Software_Component_Virtual_OperatingSystem struct {
	Billing_Item
}

type Billing_Item_Software_Component_Virtual_OperatingSystem_Microsoft struct {
	Billing_Item_Software_Component_Virtual_OperatingSystem

	Resource        *Software_VirtualLicense `json:"resource:omitempty"`
	ResourceTableId *int                     `json:"resourceTableId:omitempty"`
}

type Billing_Item_Software_Component_Virtual_OperatingSystem_Redhat struct {
	Billing_Item_Software_Component_Virtual_OperatingSystem

	Resource        *Software_Component `json:"resource:omitempty"`
	ResourceTableId *int                `json:"resourceTableId:omitempty"`
}

type Billing_Item_Software_License struct {
	Billing_Item

	Resource *Software_AccountLicense `json:"resource:omitempty"`
}

type Billing_Item_Support struct {
	Billing_Item
}

type Billing_Item_User_Customer_External_Binding struct {
	Billing_Item

	Resource *User_Customer_External_Binding `json:"resource:omitempty"`
}

type Billing_Item_Virtual_Dedicated_Rack struct {
	Billing_Item

	BillingCycleBandwidthUsage             []Network_Bandwidth_Usage             `json:"billingCycleBandwidthUsage:omitempty"`
	BillingCycleBandwidthUsageCount        *uint                                 `json:"billingCycleBandwidthUsageCount:omitempty"`
	BillingCyclePrivateBandwidthUsage      []Network_Bandwidth_Usage             `json:"billingCyclePrivateBandwidthUsage:omitempty"`
	BillingCyclePrivateBandwidthUsageCount *uint                                 `json:"billingCyclePrivateBandwidthUsageCount:omitempty"`
	BillingCyclePrivateUsageIn             *float64                              `json:"billingCyclePrivateUsageIn:omitempty"`
	BillingCyclePrivateUsageOut            *float64                              `json:"billingCyclePrivateUsageOut:omitempty"`
	BillingCyclePrivateUsageTotal          *uint                                 `json:"billingCyclePrivateUsageTotal:omitempty"`
	BillingCyclePublicBandwidthUsage       []Network_Bandwidth_Usage             `json:"billingCyclePublicBandwidthUsage:omitempty"`
	BillingCyclePublicBandwidthUsageCount  *uint                                 `json:"billingCyclePublicBandwidthUsageCount:omitempty"`
	BillingCyclePublicUsageIn              *float64                              `json:"billingCyclePublicUsageIn:omitempty"`
	BillingCyclePublicUsageOut             *float64                              `json:"billingCyclePublicUsageOut:omitempty"`
	BillingCyclePublicUsageTotal           *uint                                 `json:"billingCyclePublicUsageTotal:omitempty"`
	Resource                               *Network_Bandwidth_Version1_Allotment `json:"resource:omitempty"`
}

type Billing_Item_Virtual_Disk_Image struct {
	Billing_Item

	Resource        *Virtual_Disk_Image `json:"resource:omitempty"`
	ResourceTableId *int                `json:"resourceTableId:omitempty"`
}

type Billing_Item_Virtual_Guest struct {
	Billing_Item

	BillingCycleBandwidthUsage             []Network_Bandwidth_Usage `json:"billingCycleBandwidthUsage:omitempty"`
	BillingCycleBandwidthUsageCount        *uint                     `json:"billingCycleBandwidthUsageCount:omitempty"`
	BillingCyclePrivateBandwidthUsage      []Network_Bandwidth_Usage `json:"billingCyclePrivateBandwidthUsage:omitempty"`
	BillingCyclePrivateBandwidthUsageCount *uint                     `json:"billingCyclePrivateBandwidthUsageCount:omitempty"`
	BillingCyclePrivateUsageIn             *float64                  `json:"billingCyclePrivateUsageIn:omitempty"`
	BillingCyclePrivateUsageOut            *float64                  `json:"billingCyclePrivateUsageOut:omitempty"`
	BillingCyclePrivateUsageTotal          *uint                     `json:"billingCyclePrivateUsageTotal:omitempty"`
	BillingCyclePublicBandwidthUsage       []Network_Bandwidth_Usage `json:"billingCyclePublicBandwidthUsage:omitempty"`
	BillingCyclePublicBandwidthUsageCount  *uint                     `json:"billingCyclePublicBandwidthUsageCount:omitempty"`
	BillingCyclePublicUsageIn              *float64                  `json:"billingCyclePublicUsageIn:omitempty"`
	BillingCyclePublicUsageOut             *float64                  `json:"billingCyclePublicUsageOut:omitempty"`
	BillingCyclePublicUsageTotal           *uint                     `json:"billingCyclePublicUsageTotal:omitempty"`
	MonitoringBillingItemCount             *uint                     `json:"monitoringBillingItemCount:omitempty"`
	MonitoringBillingItems                 []Billing_Item            `json:"monitoringBillingItems:omitempty"`
	Resource                               *Virtual_Guest            `json:"resource:omitempty"`
	ResourceTableId                        *int                      `json:"resourceTableId:omitempty"`
}

type Billing_Item_Virtual_Host_Usage struct {
	Billing_Item

	Resource        *Hardware `json:"resource:omitempty"`
	ResourceTableId *int      `json:"resourceTableId:omitempty"`
}

type Billing_Item_Workspace struct {
	Billing_Item
}

type Billing_Order struct {
	Entity

	Account                      *Account                             `json:"account:omitempty"`
	AccountId                    *int                                 `json:"accountId:omitempty"`
	Brand                        *Brand                               `json:"brand:omitempty"`
	Cart                         *Billing_Order_Cart                  `json:"cart:omitempty"`
	CoreRestrictedItemCount      *uint                                `json:"coreRestrictedItemCount:omitempty"`
	CoreRestrictedItems          []Billing_Order_Item                 `json:"coreRestrictedItems:omitempty"`
	CreateDate                   *time.Time                           `json:"createDate:omitempty"`
	CreditCardTransactionCount   *uint                                `json:"creditCardTransactionCount:omitempty"`
	CreditCardTransactions       []Billing_Payment_Card_Transaction   `json:"creditCardTransactions:omitempty"`
	ExchangeRate                 *Billing_Currency_ExchangeRate       `json:"exchangeRate:omitempty"`
	Id                           *int                                 `json:"id:omitempty"`
	ImpersonatingUserRecordId    *int                                 `json:"impersonatingUserRecordId:omitempty"`
	InitialInvoice               *Billing_Invoice                     `json:"initialInvoice:omitempty"`
	ItemCount                    *uint                                `json:"itemCount:omitempty"`
	Items                        []Billing_Order_Item                 `json:"items:omitempty"`
	ModifyDate                   *time.Time                           `json:"modifyDate:omitempty"`
	OrderApprovalDate            *time.Time                           `json:"orderApprovalDate:omitempty"`
	OrderNonServerMonthlyAmount  *float64                             `json:"orderNonServerMonthlyAmount:omitempty"`
	OrderQuoteId                 *int                                 `json:"orderQuoteId:omitempty"`
	OrderServerMonthlyAmount     *float64                             `json:"orderServerMonthlyAmount:omitempty"`
	OrderTopLevelItemCount       *uint                                `json:"orderTopLevelItemCount:omitempty"`
	OrderTopLevelItems           []Billing_Order_Item                 `json:"orderTopLevelItems:omitempty"`
	OrderTotalAmount             *float64                             `json:"orderTotalAmount:omitempty"`
	OrderTotalOneTime            *float64                             `json:"orderTotalOneTime:omitempty"`
	OrderTotalOneTimeAmount      *float64                             `json:"orderTotalOneTimeAmount:omitempty"`
	OrderTotalOneTimeTaxAmount   *float64                             `json:"orderTotalOneTimeTaxAmount:omitempty"`
	OrderTotalRecurring          *float64                             `json:"orderTotalRecurring:omitempty"`
	OrderTotalRecurringAmount    *float64                             `json:"orderTotalRecurringAmount:omitempty"`
	OrderTotalRecurringTaxAmount *float64                             `json:"orderTotalRecurringTaxAmount:omitempty"`
	OrderTotalSetupAmount        *float64                             `json:"orderTotalSetupAmount:omitempty"`
	OrderType                    *Billing_Order_Type                  `json:"orderType:omitempty"`
	OrderTypeId                  *int                                 `json:"orderTypeId:omitempty"`
	PaypalTransactionCount       *uint                                `json:"paypalTransactionCount:omitempty"`
	PaypalTransactions           []Billing_Payment_PayPal_Transaction `json:"paypalTransactions:omitempty"`
	PresaleEvent                 *Sales_Presale_Event                 `json:"presaleEvent:omitempty"`
	PresaleEventId               *int                                 `json:"presaleEventId:omitempty"`
	PrivateCloudOrderFlag        *bool                                `json:"privateCloudOrderFlag:omitempty"`
	Quote                        *Billing_Order_Quote                 `json:"quote:omitempty"`
	ReferralPartner              *Account                             `json:"referralPartner:omitempty"`
	Status                       *string                              `json:"status:omitempty"`
	UpgradeRequestFlag           *bool                                `json:"upgradeRequestFlag:omitempty"`
	UserRecord                   *User_Customer                       `json:"userRecord:omitempty"`
	UserRecordId                 *int                                 `json:"userRecordId:omitempty"`
}

type Billing_Order_Cart struct {
	Billing_Order_Quote
}

type Billing_Order_Item struct {
	Entity

	BillingItem               *Billing_Item                        `json:"billingItem:omitempty"`
	BundledItemCount          *uint                                `json:"bundledItemCount:omitempty"`
	BundledItems              []Billing_Order_Item                 `json:"bundledItems:omitempty"`
	Category                  *Product_Item_Category               `json:"category:omitempty"`
	CategoryCode              *string                              `json:"categoryCode:omitempty"`
	Children                  []Billing_Order_Item                 `json:"children:omitempty"`
	ChildrenCount             *uint                                `json:"childrenCount:omitempty"`
	Description               *string                              `json:"description:omitempty"`
	DomainName                *string                              `json:"domainName:omitempty"`
	GlobalIdentifier          *string                              `json:"globalIdentifier:omitempty"`
	HardwareGenericComponent  *Hardware_Component_Model_Generic    `json:"hardwareGenericComponent:omitempty"`
	HostName                  *string                              `json:"hostName:omitempty"`
	HourlyRecurringFee        *float64                             `json:"hourlyRecurringFee:omitempty"`
	Id                        *int                                 `json:"id:omitempty"`
	Item                      *Product_Item                        `json:"item:omitempty"`
	ItemCategoryAnswerCount   *uint                                `json:"itemCategoryAnswerCount:omitempty"`
	ItemCategoryAnswers       []Billing_Order_Item_Category_Answer `json:"itemCategoryAnswers:omitempty"`
	ItemId                    *int                                 `json:"itemId:omitempty"`
	ItemPrice                 *Product_Item_Price                  `json:"itemPrice:omitempty"`
	ItemPriceId               *float64                             `json:"itemPriceId:omitempty"`
	LaborAfterTaxAmount       *float64                             `json:"laborAfterTaxAmount:omitempty"`
	LaborFee                  *float64                             `json:"laborFee:omitempty"`
	LaborFeeTaxRate           *float64                             `json:"laborFeeTaxRate:omitempty"`
	LaborTaxAmount            *float64                             `json:"laborTaxAmount:omitempty"`
	Location                  *Location                            `json:"location:omitempty"`
	NextOrderChildren         []Billing_Order_Item                 `json:"nextOrderChildren:omitempty"`
	NextOrderChildrenCount    *uint                                `json:"nextOrderChildrenCount:omitempty"`
	OldBillingItem            *Billing_Item                        `json:"oldBillingItem:omitempty"`
	OneTimeAfterTaxAmount     *float64                             `json:"oneTimeAfterTaxAmount:omitempty"`
	OneTimeFee                *float64                             `json:"oneTimeFee:omitempty"`
	OneTimeFeeTaxRate         *float64                             `json:"oneTimeFeeTaxRate:omitempty"`
	OneTimeTaxAmount          *float64                             `json:"oneTimeTaxAmount:omitempty"`
	Order                     *Billing_Order                       `json:"order:omitempty"`
	OrderApprovalDate         *time.Time                           `json:"orderApprovalDate:omitempty"`
	Package                   *Product_Package                     `json:"package:omitempty"`
	Parent                    *Billing_Order_Item                  `json:"parent:omitempty"`
	ParentId                  *int                                 `json:"parentId:omitempty"`
	PromoCodeId               *int                                 `json:"promoCodeId:omitempty"`
	Quantity                  *int                                 `json:"quantity:omitempty"`
	RecurringAfterTaxAmount   *float64                             `json:"recurringAfterTaxAmount:omitempty"`
	RecurringFee              *float64                             `json:"recurringFee:omitempty"`
	RecurringTaxAmount        *float64                             `json:"recurringTaxAmount:omitempty"`
	RedundantPowerSupplyCount *uint                                `json:"redundantPowerSupplyCount:omitempty"`
	SetupAfterTaxAmount       *float64                             `json:"setupAfterTaxAmount:omitempty"`
	SetupFee                  *float64                             `json:"setupFee:omitempty"`
	SetupFeeDeferralMonths    *int                                 `json:"setupFeeDeferralMonths:omitempty"`
	SetupFeeTaxRate           *float64                             `json:"setupFeeTaxRate:omitempty"`
	SetupTaxAmount            *float64                             `json:"setupTaxAmount:omitempty"`
	SoftwareDescription       *Software_Description                `json:"softwareDescription:omitempty"`
	StorageGroupCount         *uint                                `json:"storageGroupCount:omitempty"`
	StorageGroups             []Configuration_Storage_Group_Order  `json:"storageGroups:omitempty"`
	TotalRecurringAmount      *float64                             `json:"totalRecurringAmount:omitempty"`
	UpgradeItem               *Product_Item                        `json:"upgradeItem:omitempty"`
}

type Billing_Order_Item_Category_Answer struct {
	Entity

	Answer     *string                         `json:"answer:omitempty"`
	CreateDate *time.Time                      `json:"createDate:omitempty"`
	OrderItem  *Billing_Order_Item             `json:"orderItem:omitempty"`
	Question   *Product_Item_Category_Question `json:"question:omitempty"`
	QuestionId *int                            `json:"questionId:omitempty"`
}

type Billing_Order_Note struct {
	Entity

	CreateDate *time.Time     `json:"createDate:omitempty"`
	Employee   *User_Employee `json:"employee:omitempty"`
	Order      *Billing_Order `json:"order:omitempty"`
}

type Billing_Order_Quote struct {
	Entity

	Account                 *Account        `json:"account:omitempty"`
	AccountId               *int            `json:"accountId:omitempty"`
	CompletedPurchaseDataId *int            `json:"completedPurchaseDataId:omitempty"`
	CreateDate              *time.Time      `json:"createDate:omitempty"`
	ExpirationDate          *time.Time      `json:"expirationDate:omitempty"`
	Id                      *int            `json:"id:omitempty"`
	ModifyDate              *time.Time      `json:"modifyDate:omitempty"`
	Name                    *string         `json:"name:omitempty"`
	Order                   *Billing_Order  `json:"order:omitempty"`
	OrdersFromQuote         []Billing_Order `json:"ordersFromQuote:omitempty"`
	OrdersFromQuoteCount    *uint           `json:"ordersFromQuoteCount:omitempty"`
	PublicNote              *string         `json:"publicNote:omitempty"`
	QuoteKey                *string         `json:"quoteKey:omitempty"`
	Status                  *string         `json:"status:omitempty"`
}

type Billing_Order_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	Type        *string `json:"type:omitempty"`
}

type Billing_Payment_Card_ChangeRequest struct {
	Entity

	Account                         *Account                          `json:"account:omitempty"`
	AccountId                       *int                              `json:"accountId:omitempty"`
	Amount                          *float64                          `json:"amount:omitempty"`
	AuthorizedCreditCardTransaction *Billing_Payment_Card_Transaction `json:"authorizedCreditCardTransaction:omitempty"`
	BillingAddressLine1             *string                           `json:"billingAddressLine1:omitempty"`
	BillingAddressLine2             *string                           `json:"billingAddressLine2:omitempty"`
	BillingCity                     *string                           `json:"billingCity:omitempty"`
	BillingCountryCode              *string                           `json:"billingCountryCode:omitempty"`
	BillingEmail                    *string                           `json:"billingEmail:omitempty"`
	BillingNameCompany              *string                           `json:"billingNameCompany:omitempty"`
	BillingNameFirst                *string                           `json:"billingNameFirst:omitempty"`
	BillingNameLast                 *string                           `json:"billingNameLast:omitempty"`
	BillingPhoneFax                 *string                           `json:"billingPhoneFax:omitempty"`
	BillingPhoneVoice               *string                           `json:"billingPhoneVoice:omitempty"`
	BillingPostalCode               *string                           `json:"billingPostalCode:omitempty"`
	BillingState                    *string                           `json:"billingState:omitempty"`
	CaptureCreditCardTransaction    *Billing_Payment_Card_Transaction `json:"captureCreditCardTransaction:omitempty"`
	CardAccountLast4                *string                           `json:"cardAccountLast4:omitempty"`
	CardAccountNumber               *string                           `json:"cardAccountNumber:omitempty"`
	CardExpirationMonth             *string                           `json:"cardExpirationMonth:omitempty"`
	CardExpirationYear              *string                           `json:"cardExpirationYear:omitempty"`
	CardNickname                    *string                           `json:"cardNickname:omitempty"`
	CardType                        *string                           `json:"cardType:omitempty"`
	CreditCardVerificationNumber    *string                           `json:"creditCardVerificationNumber:omitempty"`
	CurrencyShortName               *string                           `json:"currencyShortName:omitempty"`
	DeviceFingerprintId             *string                           `json:"deviceFingerprintId:omitempty"`
	Id                              *int                              `json:"id:omitempty"`
	Notes                           *string                           `json:"notes:omitempty"`
	PaymentRoleId                   *int                              `json:"paymentRoleId:omitempty"`
	PaymentType                     *string                           `json:"paymentType:omitempty"`
	TicketAttachmentReferenceCount  *uint                             `json:"ticketAttachmentReferenceCount:omitempty"`
	TicketAttachmentReferences      []Ticket_Attachment               `json:"ticketAttachmentReferences:omitempty"`
	TicketId                        *int                              `json:"ticketId:omitempty"`
}

type Billing_Payment_Card_ManualPayment struct {
	Entity

	Account                           *Account                            `json:"account:omitempty"`
	AccountId                         *int                                `json:"accountId:omitempty"`
	Amount                            *float64                            `json:"amount:omitempty"`
	AuthorizedCreditCardTransaction   *Billing_Payment_Card_Transaction   `json:"authorizedCreditCardTransaction:omitempty"`
	AuthorizedCreditCardTransactionId *int                                `json:"authorizedCreditCardTransactionId:omitempty"`
	AuthorizedPayPalTransaction       *Billing_Payment_PayPal_Transaction `json:"authorizedPayPalTransaction:omitempty"`
	AuthorizedPayPalTransactionId     *int                                `json:"authorizedPayPalTransactionId:omitempty"`
	BillingAddressLine1               *string                             `json:"billingAddressLine1:omitempty"`
	BillingAddressLine2               *string                             `json:"billingAddressLine2:omitempty"`
	BillingCity                       *string                             `json:"billingCity:omitempty"`
	BillingCountryCode                *string                             `json:"billingCountryCode:omitempty"`
	BillingEmail                      *string                             `json:"billingEmail:omitempty"`
	BillingNameCompany                *string                             `json:"billingNameCompany:omitempty"`
	BillingNameFirst                  *string                             `json:"billingNameFirst:omitempty"`
	BillingNameLast                   *string                             `json:"billingNameLast:omitempty"`
	BillingPhoneFax                   *string                             `json:"billingPhoneFax:omitempty"`
	BillingPhoneVoice                 *string                             `json:"billingPhoneVoice:omitempty"`
	BillingPostalCode                 *string                             `json:"billingPostalCode:omitempty"`
	BillingState                      *string                             `json:"billingState:omitempty"`
	CancelUrl                         *string                             `json:"cancelUrl:omitempty"`
	CaptureCreditCardTransaction      *Billing_Payment_Card_Transaction   `json:"captureCreditCardTransaction:omitempty"`
	CapturePayPalTransaction          *Billing_Payment_PayPal_Transaction `json:"capturePayPalTransaction:omitempty"`
	CardAccountHash                   *string                             `json:"cardAccountHash:omitempty"`
	CardAccountLast4                  *string                             `json:"cardAccountLast4:omitempty"`
	CardAccountNumber                 *string                             `json:"cardAccountNumber:omitempty"`
	CardExpirationMonth               *string                             `json:"cardExpirationMonth:omitempty"`
	CardExpirationYear                *string                             `json:"cardExpirationYear:omitempty"`
	CardType                          *string                             `json:"cardType:omitempty"`
	CreditCardVerificationNumber      *string                             `json:"creditCardVerificationNumber:omitempty"`
	CurrencyShortName                 *string                             `json:"currencyShortName:omitempty"`
	DeviceFingerprintId               *string                             `json:"deviceFingerprintId:omitempty"`
	FromIpAddress                     *string                             `json:"fromIpAddress:omitempty"`
	Id                                *int                                `json:"id:omitempty"`
	Notes                             *string                             `json:"notes:omitempty"`
	PaymentType                       *string                             `json:"paymentType:omitempty"`
	ReturnUrl                         *string                             `json:"returnUrl:omitempty"`
	TicketAttachmentReferenceCount    *uint                               `json:"ticketAttachmentReferenceCount:omitempty"`
	TicketAttachmentReferences        []Ticket_Attachment                 `json:"ticketAttachmentReferences:omitempty"`
	Type                              *string                             `json:"type:omitempty"`
}

type Billing_Payment_Card_Transaction struct {
	Billing_Payment_Transaction

	Account             *Account       `json:"account:omitempty"`
	AccountId           *int           `json:"accountId:omitempty"`
	Amount              *float64       `json:"amount:omitempty"`
	BillingAddressLine1 *string        `json:"billingAddressLine1:omitempty"`
	BillingAddressLine2 *string        `json:"billingAddressLine2:omitempty"`
	BillingCity         *string        `json:"billingCity:omitempty"`
	BillingCountryCode  *string        `json:"billingCountryCode:omitempty"`
	BillingEmail        *string        `json:"billingEmail:omitempty"`
	BillingNameCompany  *string        `json:"billingNameCompany:omitempty"`
	BillingNameFirst    *string        `json:"billingNameFirst:omitempty"`
	BillingNameLast     *string        `json:"billingNameLast:omitempty"`
	BillingPhoneFax     *string        `json:"billingPhoneFax:omitempty"`
	BillingPhoneVoice   *string        `json:"billingPhoneVoice:omitempty"`
	BillingPostalCode   *string        `json:"billingPostalCode:omitempty"`
	BillingState        *string        `json:"billingState:omitempty"`
	CardAccountLast4    *int           `json:"cardAccountLast4:omitempty"`
	CardExpirationMonth *int           `json:"cardExpirationMonth:omitempty"`
	CardExpirationYear  *int           `json:"cardExpirationYear:omitempty"`
	CardType            *string        `json:"cardType:omitempty"`
	CreateDate          *time.Time     `json:"createDate:omitempty"`
	Id                  *int           `json:"id:omitempty"`
	InvoiceId           *int           `json:"invoiceId:omitempty"`
	ModifyDate          *time.Time     `json:"modifyDate:omitempty"`
	Order               *Billing_Order `json:"order:omitempty"`
	OrderFromIpAddress  *string        `json:"orderFromIpAddress:omitempty"`
	ReferenceCode       *string        `json:"referenceCode:omitempty"`
	RequestId           *string        `json:"requestId:omitempty"`
	ReturnStatus        *int           `json:"returnStatus:omitempty"`
	SerializedReply     *string        `json:"serializedReply:omitempty"`
	SerializedRequest   *string        `json:"serializedRequest:omitempty"`
}

type Billing_Payment_PayPal_Transaction struct {
	Billing_Payment_Transaction

	Account              *Account       `json:"account:omitempty"`
	AccountId            *int           `json:"accountId:omitempty"`
	AddressCityName      *string        `json:"addressCityName:omitempty"`
	AddressCountry       *string        `json:"addressCountry:omitempty"`
	AddressName          *string        `json:"addressName:omitempty"`
	AddressPostalCode    *string        `json:"addressPostalCode:omitempty"`
	AddressStateProvence *string        `json:"addressStateProvence:omitempty"`
	AddressStatus        *string        `json:"addressStatus:omitempty"`
	AddressStreet1       *string        `json:"addressStreet1:omitempty"`
	AddressStreet2       *string        `json:"addressStreet2:omitempty"`
	ContactPhone         *string        `json:"contactPhone:omitempty"`
	CreateDate           *time.Time     `json:"createDate:omitempty"`
	ExchangeRate         *string        `json:"exchangeRate:omitempty"`
	FeeAmount            *float64       `json:"feeAmount:omitempty"`
	GrossAmount          *float64       `json:"grossAmount:omitempty"`
	Id                   *int           `json:"id:omitempty"`
	InvoiceId            *int           `json:"invoiceId:omitempty"`
	LastPaypalCommand    *string        `json:"lastPaypalCommand:omitempty"`
	ModifyDate           *time.Time     `json:"modifyDate:omitempty"`
	Order                *Billing_Order `json:"order:omitempty"`
	OrderFromIpAddress   *string        `json:"orderFromIpAddress:omitempty"`
	OrderTotal           *float64       `json:"orderTotal:omitempty"`
	Payer                *string        `json:"payer:omitempty"`
	PayerBusiness        *string        `json:"payerBusiness:omitempty"`
	PayerCountry         *string        `json:"payerCountry:omitempty"`
	PayerFirstName       *string        `json:"payerFirstName:omitempty"`
	PayerId              *string        `json:"payerId:omitempty"`
	PayerLastName        *string        `json:"payerLastName:omitempty"`
	PayerStatus          *string        `json:"payerStatus:omitempty"`
	PaymentDate          *time.Time     `json:"paymentDate:omitempty"`
	PaymentStatus        *string        `json:"paymentStatus:omitempty"`
	PaymentType          *string        `json:"paymentType:omitempty"`
	PendingReason        *string        `json:"pendingReason:omitempty"`
	SerializedReply      *string        `json:"serializedReply:omitempty"`
	SerializedRequest    *string        `json:"serializedRequest:omitempty"`
	SettleAmount         *float64       `json:"settleAmount:omitempty"`
	TaxAmount            *float64       `json:"taxAmount:omitempty"`
	Token                *string        `json:"token:omitempty"`
	TransactionId        *string        `json:"transactionId:omitempty"`
	TransactionType      *string        `json:"transactionType:omitempty"`
}

type Billing_Payment_Processor struct {
	Entity

	BrandAssignmentCount *uint                              `json:"brandAssignmentCount:omitempty"`
	BrandAssignments     []Brand_Payment_Processor          `json:"brandAssignments:omitempty"`
	Description          *string                            `json:"description:omitempty"`
	Name                 *string                            `json:"name:omitempty"`
	OwnerAccount         *Account                           `json:"ownerAccount:omitempty"`
	PaymentMethodCount   *uint                              `json:"paymentMethodCount:omitempty"`
	PaymentMethods       []Billing_Payment_Processor_Method `json:"paymentMethods:omitempty"`
	Type                 *Billing_Payment_Processor_Type    `json:"type:omitempty"`
}

type Billing_Payment_Processor_Method struct {
	Entity

	MethodKey            *string                    `json:"methodKey:omitempty"`
	MultipleCurrencyFlag *bool                      `json:"multipleCurrencyFlag:omitempty"`
	PaymentProcessor     *Billing_Payment_Processor `json:"paymentProcessor:omitempty"`
	PaymentType          *Billing_Payment_Type      `json:"paymentType:omitempty"`
}

type Billing_Payment_Processor_Type struct {
	Entity

	Description           *string                     `json:"description:omitempty"`
	KeyName               *string                     `json:"keyName:omitempty"`
	Name                  *string                     `json:"name:omitempty"`
	PaymentProcessorCount *uint                       `json:"paymentProcessorCount:omitempty"`
	PaymentProcessors     []Billing_Payment_Processor `json:"paymentProcessors:omitempty"`
}

type Billing_Payment_Transaction struct {
	Entity
}

type Billing_Payment_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Brand struct {
	Entity

	Account                                 *Account                                     `json:"account:omitempty"`
	AllOwnedAccountCount                    *uint                                        `json:"allOwnedAccountCount:omitempty"`
	AllOwnedAccounts                        []Account                                    `json:"allOwnedAccounts:omitempty"`
	AllowAccountCreationFlag                *bool                                        `json:"allowAccountCreationFlag:omitempty"`
	Catalog                                 *Product_Catalog                             `json:"catalog:omitempty"`
	CatalogId                               *int                                         `json:"catalogId:omitempty"`
	ContactCount                            *uint                                        `json:"contactCount:omitempty"`
	Contacts                                []Brand_Contact                              `json:"contacts:omitempty"`
	CustomerCountryLocationRestrictionCount *uint                                        `json:"customerCountryLocationRestrictionCount:omitempty"`
	CustomerCountryLocationRestrictions     []Brand_Restriction_Location_CustomerCountry `json:"customerCountryLocationRestrictions:omitempty"`
	Distributor                             *Brand                                       `json:"distributor:omitempty"`
	DistributorChildFlag                    *bool                                        `json:"distributorChildFlag:omitempty"`
	DistributorFlag                         *string                                      `json:"distributorFlag:omitempty"`
	Hardware                                []Hardware                                   `json:"hardware:omitempty"`
	HardwareCount                           *uint                                        `json:"hardwareCount:omitempty"`
	HasAgentSupportFlag                     *bool                                        `json:"hasAgentSupportFlag:omitempty"`
	Id                                      *int                                         `json:"id:omitempty"`
	KeyName                                 *string                                      `json:"keyName:omitempty"`
	LongName                                *string                                      `json:"longName:omitempty"`
	Name                                    *string                                      `json:"name:omitempty"`
	OpenTicketCount                         *uint                                        `json:"openTicketCount:omitempty"`
	OpenTickets                             []Ticket                                     `json:"openTickets:omitempty"`
	OwnedAccountCount                       *uint                                        `json:"ownedAccountCount:omitempty"`
	OwnedAccounts                           []Account                                    `json:"ownedAccounts:omitempty"`
	TicketCount                             *uint                                        `json:"ticketCount:omitempty"`
	TicketGroupCount                        *uint                                        `json:"ticketGroupCount:omitempty"`
	TicketGroups                            []Ticket_Group                               `json:"ticketGroups:omitempty"`
	Tickets                                 []Ticket                                     `json:"tickets:omitempty"`
	UserCount                               *uint                                        `json:"userCount:omitempty"`
	Users                                   []User_Customer                              `json:"users:omitempty"`
	VirtualGuestCount                       *uint                                        `json:"virtualGuestCount:omitempty"`
	VirtualGuests                           []Virtual_Guest                              `json:"virtualGuests:omitempty"`
}

type Brand_Attribute struct {
	Entity

	Brand *Brand `json:"brand:omitempty"`
}

type Brand_Contact struct {
	Entity

	Address1           *string             `json:"address1:omitempty"`
	Address2           *string             `json:"address2:omitempty"`
	AlternatePhone     *string             `json:"alternatePhone:omitempty"`
	Brand              *Brand              `json:"brand:omitempty"`
	BrandContactType   *Brand_Contact_Type `json:"brandContactType:omitempty"`
	BrandContactTypeId *int                `json:"brandContactTypeId:omitempty"`
	City               *string             `json:"city:omitempty"`
	Country            *string             `json:"country:omitempty"`
	Email              *string             `json:"email:omitempty"`
	FaxPhone           *string             `json:"faxPhone:omitempty"`
	FirstName          *string             `json:"firstName:omitempty"`
	LastName           *string             `json:"lastName:omitempty"`
	OfficePhone        *string             `json:"officePhone:omitempty"`
	PostalCode         *string             `json:"postalCode:omitempty"`
	State              *string             `json:"state:omitempty"`
}

type Brand_Contact_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Brand_Payment_Processor struct {
	Entity

	Brand            *Brand                     `json:"brand:omitempty"`
	PaymentProcessor *Billing_Payment_Processor `json:"paymentProcessor:omitempty"`
}

type Brand_Restriction_Location_CustomerCountry struct {
	Entity

	Brand               *Brand    `json:"brand:omitempty"`
	BrandId             *int      `json:"brandId:omitempty"`
	CustomerCountryCode *string   `json:"customerCountryCode:omitempty"`
	Location            *Location `json:"location:omitempty"`
	LocationId          *int      `json:"locationId:omitempty"`
}

type Catalyst_Affiliate struct {
	Entity

	Id                             *int    `json:"id:omitempty"`
	Name                           *string `json:"name:omitempty"`
	SkipCreditCardVerificationFlag *bool   `json:"skipCreditCardVerificationFlag:omitempty"`
}

type Catalyst_Company_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
}

type Catalyst_Enrollment struct {
	Entity

	Account                  *Account               `json:"account:omitempty"`
	AccountId                *int                   `json:"accountId:omitempty"`
	Affiliate                *Catalyst_Affiliate    `json:"affiliate:omitempty"`
	AffiliateId              *int                   `json:"affiliateId:omitempty"`
	AgreementCompleteFlag    *int                   `json:"agreementCompleteFlag:omitempty"`
	CompanyDescription       *string                `json:"companyDescription:omitempty"`
	CompanyType              *Catalyst_Company_Type `json:"companyType:omitempty"`
	CompanyTypeId            *int                   `json:"companyTypeId:omitempty"`
	EnrollmentDate           *time.Time             `json:"enrollmentDate:omitempty"`
	GraduationDate           *time.Time             `json:"graduationDate:omitempty"`
	IsActiveFlag             *bool                  `json:"isActiveFlag:omitempty"`
	MonthlyCreditAmount      *float64               `json:"monthlyCreditAmount:omitempty"`
	Representative           *User_Employee         `json:"representative:omitempty"`
	RepresentativeEmployeeId *int                   `json:"representativeEmployeeId:omitempty"`
}

type Catalyst_Enrollment_Request struct {
	Entity

	Address1                    *string                `json:"address1:omitempty"`
	Address2                    *string                `json:"address2:omitempty"`
	Affiliate                   *Catalyst_Affiliate    `json:"affiliate:omitempty"`
	AffiliateId                 *int                   `json:"affiliateId:omitempty"`
	AgreementCompleteFlag       *bool                  `json:"agreementCompleteFlag:omitempty"`
	ApplyToGepFlag              *bool                  `json:"applyToGepFlag:omitempty"`
	CardAccountNumber           *string                `json:"cardAccountNumber:omitempty"`
	CardExpirationMonth         *string                `json:"cardExpirationMonth:omitempty"`
	CardExpirationYear          *string                `json:"cardExpirationYear:omitempty"`
	CardType                    *string                `json:"cardType:omitempty"`
	CardVerificationNumber      *string                `json:"cardVerificationNumber:omitempty"`
	City                        *string                `json:"city:omitempty"`
	CompanyDescription          *string                `json:"companyDescription:omitempty"`
	CompanyName                 *string                `json:"companyName:omitempty"`
	CompanyType                 *Catalyst_Company_Type `json:"companyType:omitempty"`
	CompanyTypeId               *int                   `json:"companyTypeId:omitempty"`
	CompanyUrl                  *string                `json:"companyUrl:omitempty"`
	Country                     *string                `json:"country:omitempty"`
	CurrentUserChoice           *int                   `json:"currentUserChoice:omitempty"`
	DeviceFingerprintId         *string                `json:"deviceFingerprintId:omitempty"`
	Email                       *string                `json:"email:omitempty"`
	FirstName                   *string                `json:"firstName:omitempty"`
	FutureUserChoice            *int                   `json:"futureUserChoice:omitempty"`
	IncubatorName               *string                `json:"incubatorName:omitempty"`
	InvestorName                *string                `json:"investorName:omitempty"`
	IpAddress                   *string                `json:"ipAddress:omitempty"`
	LastName                    *string                `json:"lastName:omitempty"`
	OfficePhone                 *string                `json:"officePhone:omitempty"`
	OverFiveYearsOldFlag        *bool                  `json:"overFiveYearsOldFlag:omitempty"`
	PostalCode                  *string                `json:"postalCode:omitempty"`
	ReferralCode                *string                `json:"referralCode:omitempty"`
	RevenueOverOneMillionFlag   *bool                  `json:"revenueOverOneMillionFlag:omitempty"`
	SkipCatalystApplicationFlag *bool                  `json:"skipCatalystApplicationFlag:omitempty"`
	State                       *string                `json:"state:omitempty"`
	VatId                       *string                `json:"vatId:omitempty"`
}

type Catalyst_Enrollment_Request_Container_AnswerOption struct {
	Entity

	Answer *string `json:"answer:omitempty"`
	Index  *int    `json:"index:omitempty"`
}

type Compliance_Report_Type struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Configuration_Storage_Filesystem_Type struct {
	Entity

	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Configuration_Storage_Group_Array_Type struct {
	Entity

	Description                 *string                    `json:"description:omitempty"`
	DriveMultiplier             *int                       `json:"driveMultiplier:omitempty"`
	HardwareComponentModelCount *uint                      `json:"hardwareComponentModelCount:omitempty"`
	HardwareComponentModels     []Hardware_Component_Model `json:"hardwareComponentModels:omitempty"`
	HotspareAllow               *bool                      `json:"hotspareAllow:omitempty"`
	Id                          *int                       `json:"id:omitempty"`
	KeyName                     *string                    `json:"keyName:omitempty"`
	MaximumDrives               *int                       `json:"maximumDrives:omitempty"`
	MinimumDrives               *int                       `json:"minimumDrives:omitempty"`
	Name                        *string                    `json:"name:omitempty"`
}

type Configuration_Storage_Group_Order struct {
	Entity

	ArrayNumber        *int                                    `json:"arrayNumber:omitempty"`
	ArraySize          *float64                                `json:"arraySize:omitempty"`
	ArrayType          *Configuration_Storage_Group_Array_Type `json:"arrayType:omitempty"`
	ArrayTypeId        *int                                    `json:"arrayTypeId:omitempty"`
	BillingOrderItem   *Billing_Order_Item                     `json:"billingOrderItem:omitempty"`
	BillingOrderItemId *int                                    `json:"billingOrderItemId:omitempty"`
	HardDrives         []int                                   `json:"hardDrives:omitempty"`
	HotSpareDrives     []int                                   `json:"hotSpareDrives:omitempty"`
	PartitionData      *string                                 `json:"partitionData:omitempty"`
}

type Configuration_Storage_Group_Template_Group struct {
	Entity

	Grow             *bool                                   `json:"grow:omitempty"`
	HardDrivesString *string                                 `json:"hardDrivesString:omitempty"`
	OrderIndex       *int                                    `json:"orderIndex:omitempty"`
	Size             *float64                                `json:"size:omitempty"`
	Type             *Configuration_Storage_Group_Array_Type `json:"type:omitempty"`
}

type Configuration_Template struct {
	Entity

	Account                             *Account                                                  `json:"account:omitempty"`
	AccountId                           *int                                                      `json:"accountId:omitempty"`
	ConfigurationSectionCount           *uint                                                     `json:"configurationSectionCount:omitempty"`
	ConfigurationSections               []Configuration_Template_Section                          `json:"configurationSections:omitempty"`
	ConfigurationTemplateReference      []Monitoring_Agent_Configuration_Template_Group_Reference `json:"configurationTemplateReference:omitempty"`
	ConfigurationTemplateReferenceCount *uint                                                     `json:"configurationTemplateReferenceCount:omitempty"`
	CreateDate                          *time.Time                                                `json:"createDate:omitempty"`
	DefaultValueCount                   *uint                                                     `json:"defaultValueCount:omitempty"`
	DefaultValues                       []Configuration_Template_Section_Definition_Value         `json:"defaultValues:omitempty"`
	DefinitionCount                     *uint                                                     `json:"definitionCount:omitempty"`
	Definitions                         []Configuration_Template_Section_Definition               `json:"definitions:omitempty"`
	Description                         *string                                                   `json:"description:omitempty"`
	Id                                  *int                                                      `json:"id:omitempty"`
	Item                                *Product_Item                                             `json:"item:omitempty"`
	ItemId                              *int                                                      `json:"itemId:omitempty"`
	LinkedSectionReferences             *Configuration_Template_Section_Reference                 `json:"linkedSectionReferences:omitempty"`
	ModifyDate                          *time.Time                                                `json:"modifyDate:omitempty"`
	Name                                *string                                                   `json:"name:omitempty"`
	Parent                              *Configuration_Template                                   `json:"parent:omitempty"`
	ParentId                            *int                                                      `json:"parentId:omitempty"`
	User                                *User_Customer                                            `json:"user:omitempty"`
	UserRecordId                        *int                                                      `json:"userRecordId:omitempty"`
}

type Configuration_Template_Attribute struct {
	Entity

	ConfigurationTemplate *Configuration_Template `json:"configurationTemplate:omitempty"`
	Value                 *string                 `json:"value:omitempty"`
}

type Configuration_Template_Section struct {
	Entity

	CreateDate              *time.Time                                  `json:"createDate:omitempty"`
	DefinitionCount         *uint                                       `json:"definitionCount:omitempty"`
	Definitions             []Configuration_Template_Section_Definition `json:"definitions:omitempty"`
	Description             *string                                     `json:"description:omitempty"`
	DisallowedDeletionFlag  *bool                                       `json:"disallowedDeletionFlag:omitempty"`
	Id                      *int                                        `json:"id:omitempty"`
	LinkedTemplate          *Configuration_Template                     `json:"linkedTemplate:omitempty"`
	LinkedTemplateId        *string                                     `json:"linkedTemplateId:omitempty"`
	LinkedTemplateReference *Configuration_Template_Section_Reference   `json:"linkedTemplateReference:omitempty"`
	ModifyDate              *time.Time                                  `json:"modifyDate:omitempty"`
	Name                    *string                                     `json:"name:omitempty"`
	ParentId                *int                                        `json:"parentId:omitempty"`
	ProfileCount            *uint                                       `json:"profileCount:omitempty"`
	Profiles                []Configuration_Template_Section_Profile    `json:"profiles:omitempty"`
	SectionType             *Configuration_Template_Section_Type        `json:"sectionType:omitempty"`
	SectionTypeName         *string                                     `json:"sectionTypeName:omitempty"`
	Sort                    *int                                        `json:"sort:omitempty"`
	SubSectionCount         *uint                                       `json:"subSectionCount:omitempty"`
	SubSections             []Configuration_Template_Section            `json:"subSections:omitempty"`
	Template                *Configuration_Template                     `json:"template:omitempty"`
	TemplateId              *string                                     `json:"templateId:omitempty"`
	TypeId                  *int                                        `json:"typeId:omitempty"`
}

type Configuration_Template_Section_Attribute struct {
	Entity

	ConfigurationSection *Configuration_Template_Section `json:"configurationSection:omitempty"`
	Value                *string                         `json:"value:omitempty"`
}

type Configuration_Template_Section_Definition struct {
	Entity

	AttributeCount     *uint                                                 `json:"attributeCount:omitempty"`
	Attributes         []Configuration_Template_Section_Definition_Attribute `json:"attributes:omitempty"`
	CreateDate         *time.Time                                            `json:"createDate:omitempty"`
	DefaultValue       *Configuration_Template_Section_Definition_Value      `json:"defaultValue:omitempty"`
	Description        *string                                               `json:"description:omitempty"`
	EnumerationValues  *string                                               `json:"enumerationValues:omitempty"`
	Group              *Configuration_Template_Section_Definition_Group      `json:"group:omitempty"`
	GroupId            *string                                               `json:"groupId:omitempty"`
	Id                 *int                                                  `json:"id:omitempty"`
	MaximumValue       *string                                               `json:"maximumValue:omitempty"`
	MinimumValue       *string                                               `json:"minimumValue:omitempty"`
	ModifyDate         *time.Time                                            `json:"modifyDate:omitempty"`
	MonitoringDataFlag *bool                                                 `json:"monitoringDataFlag:omitempty"`
	Name               *string                                               `json:"name:omitempty"`
	Path               *string                                               `json:"path:omitempty"`
	RequireValueFlag   *int                                                  `json:"requireValueFlag:omitempty"`
	Section            *Configuration_Template_Section                       `json:"section:omitempty"`
	SectionId          *int                                                  `json:"sectionId:omitempty"`
	ShortName          *string                                               `json:"shortName:omitempty"`
	Sort               *int                                                  `json:"sort:omitempty"`
	TypeId             *int                                                  `json:"typeId:omitempty"`
	ValueType          *Configuration_Template_Section_Definition_Type       `json:"valueType:omitempty"`
}

type Configuration_Template_Section_Definition_Attribute struct {
	Entity

	AttributeType           *Configuration_Template_Section_Definition_Attribute_Type `json:"attributeType:omitempty"`
	ConfigurationDefinition *Configuration_Template_Section_Definition                `json:"configurationDefinition:omitempty"`
	Value                   *string                                                   `json:"value:omitempty"`
}

type Configuration_Template_Section_Definition_Attribute_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Configuration_Template_Section_Definition_Group struct {
	Entity

	CreateDate  *time.Time                                       `json:"createDate:omitempty"`
	Description *string                                          `json:"description:omitempty"`
	Id          *int                                             `json:"id:omitempty"`
	Name        *string                                          `json:"name:omitempty"`
	Parent      *Configuration_Template_Section_Definition_Group `json:"parent:omitempty"`
	SortOrder   *int                                             `json:"sortOrder:omitempty"`
}

type Configuration_Template_Section_Definition_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Configuration_Template_Section_Definition_Value struct {
	Entity

	CreateDate   *time.Time                                 `json:"createDate:omitempty"`
	Definition   *Configuration_Template_Section_Definition `json:"definition:omitempty"`
	DefinitionId *int                                       `json:"definitionId:omitempty"`
	ModifyDate   *time.Time                                 `json:"modifyDate:omitempty"`
	Template     *Configuration_Template                    `json:"template:omitempty"`
	TemplateId   *int                                       `json:"templateId:omitempty"`
	Value        *string                                    `json:"value:omitempty"`
}

type Configuration_Template_Section_Profile struct {
	Entity

	AgentId              *int                            `json:"agentId:omitempty"`
	ConfigurationSection *Configuration_Template_Section `json:"configurationSection:omitempty"`
	CreateDate           *time.Time                      `json:"createDate:omitempty"`
	Id                   *int                            `json:"id:omitempty"`
	MonitoringAgent      *Monitoring_Agent               `json:"monitoringAgent:omitempty"`
	Name                 *string                         `json:"name:omitempty"`
	SectionId            *int                            `json:"sectionId:omitempty"`
}

type Configuration_Template_Section_Reference struct {
	Entity

	CreateDate *time.Time                      `json:"createDate:omitempty"`
	Id         *int                            `json:"id:omitempty"`
	ModifyDate *time.Time                      `json:"modifyDate:omitempty"`
	Section    *Configuration_Template_Section `json:"section:omitempty"`
	SectionId  *int                            `json:"sectionId:omitempty"`
	Template   *Configuration_Template         `json:"template:omitempty"`
	TemplateId *int                            `json:"templateId:omitempty"`
}

type Configuration_Template_Section_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Configuration_Template_Type struct {
	Entity

	CreateDate  *time.Time `json:"createDate:omitempty"`
	Description *string    `json:"description:omitempty"`
	Id          *int       `json:"id:omitempty"`
	Name        *string    `json:"name:omitempty"`
}

type Container_Account_Discount_Program struct {
	Entity

	AppliedCredit           *float64   `json:"appliedCredit:omitempty"`
	IsParticipant           *bool      `json:"isParticipant:omitempty"`
	LifetimeAppliedCredit   *float64   `json:"lifetimeAppliedCredit:omitempty"`
	LifetimeCredit          *float64   `json:"lifetimeCredit:omitempty"`
	LifetimeRemainingCredit *float64   `json:"lifetimeRemainingCredit:omitempty"`
	MaximumActiveOrders     *float64   `json:"maximumActiveOrders:omitempty"`
	MonthlyCredit           *float64   `json:"monthlyCredit:omitempty"`
	PostTaxRemainingCredit  *float64   `json:"postTaxRemainingCredit:omitempty"`
	ProgramEndDate          *time.Time `json:"programEndDate:omitempty"`
	ProgramName             *string    `json:"programName:omitempty"`
	RemainingCredit         *float64   `json:"remainingCredit:omitempty"`
	RemainingCreditTax      *float64   `json:"remainingCreditTax:omitempty"`
}

type Container_Account_Graph_Outputs struct {
	Entity

	ClosedTickets                      *string    `json:"closedTickets:omitempty"`
	CompletedBackupCount               *string    `json:"completedBackupCount:omitempty"`
	ConflictBackupCount                *string    `json:"conflictBackupCount:omitempty"`
	EndDate                            *time.Time `json:"endDate:omitempty"`
	FailedBackupCount                  *string    `json:"failedBackupCount:omitempty"`
	GraphError                         *string    `json:"graphError:omitempty"`
	GraphImage                         *[]byte    `json:"graphImage:omitempty"`
	HardwareUptime                     *string    `json:"hardwareUptime:omitempty"`
	InboundUsage                       *string    `json:"inboundUsage:omitempty"`
	OpenTickets                        *string    `json:"openTickets:omitempty"`
	OutboundUsage                      *string    `json:"outboundUsage:omitempty"`
	PendingCustomerResponseTicketCount *string    `json:"pendingCustomerResponseTicketCount:omitempty"`
	StartDate                          *time.Time `json:"startDate:omitempty"`
	UrlUptime                          *string    `json:"urlUptime:omitempty"`
	WaitingEmployeeResponseTicketCount *string    `json:"waitingEmployeeResponseTicketCount:omitempty"`
}

type Container_Account_Historical_Summary struct {
	Entity

	Details   []Container_Account_Historical_Summary_Detail `json:"details:omitempty"`
	EndDate   *time.Time                                    `json:"endDate:omitempty"`
	StartDate *time.Time                                    `json:"startDate:omitempty"`
}

type Container_Account_Historical_Summary_Detail struct {
	Entity

	EndDate   *time.Time `json:"endDate:omitempty"`
	StartDate *time.Time `json:"startDate:omitempty"`
}

type Container_Account_Historical_Summary_Detail_Uptime struct {
	Container_Account_Historical_Summary_Detail

	CloudComputingInstance *Virtual_Guest                        `json:"cloudComputingInstance:omitempty"`
	ConfigurationValue     *Monitoring_Agent_Configuration_Value `json:"configurationValue:omitempty"`
	Data                   []Metric_Tracking_Object_Data         `json:"data:omitempty"`
	Hardware               *Hardware                             `json:"hardware:omitempty"`
}

type Container_Account_Historical_Summary_Uptime struct {
	Container_Account_Historical_Summary
}

type Container_Account_Payment_Method_CreditCard struct {
	Entity

	Address1                    *string `json:"address1:omitempty"`
	Address2                    *string `json:"address2:omitempty"`
	City                        *string `json:"city:omitempty"`
	Country                     *string `json:"country:omitempty"`
	CurrencyShortName           *string `json:"currencyShortName:omitempty"`
	CybersourceAssignedCardType *string `json:"cybersourceAssignedCardType:omitempty"`
	ExpireMonth                 *string `json:"expireMonth:omitempty"`
	ExpireYear                  *string `json:"expireYear:omitempty"`
	FirstName                   *string `json:"firstName:omitempty"`
	LastFourDigits              *string `json:"lastFourDigits:omitempty"`
	LastName                    *string `json:"lastName:omitempty"`
	Nickname                    *string `json:"nickname:omitempty"`
	PaymentMethodRoleName       *string `json:"paymentMethodRoleName:omitempty"`
	PaymentTypeId               *string `json:"paymentTypeId:omitempty"`
	PaymentTypeName             *string `json:"paymentTypeName:omitempty"`
	PostalCode                  *string `json:"postalCode:omitempty"`
	State                       *string `json:"state:omitempty"`
}

type Container_Auxiliary_Network_Status_Reading struct {
	Entity

	AveragePing   *float64   `json:"averagePing:omitempty"`
	Fails         *int       `json:"fails:omitempty"`
	Frequency     *int       `json:"frequency:omitempty"`
	Label         *string    `json:"label:omitempty"`
	LastCheckDate *time.Time `json:"lastCheckDate:omitempty"`
	LastDownDate  *time.Time `json:"lastDownDate:omitempty"`
	Latency       *float64   `json:"latency:omitempty"`
	Location      *string    `json:"location:omitempty"`
	MaximumPing   *float64   `json:"maximumPing:omitempty"`
	MinimumPing   *float64   `json:"minimumPing:omitempty"`
	PingLoss      *float64   `json:"pingLoss:omitempty"`
	StartDate     *time.Time `json:"startDate:omitempty"`
	StatusCode    *string    `json:"statusCode:omitempty"`
	StatusMessage *string    `json:"statusMessage:omitempty"`
	Target        *string    `json:"target:omitempty"`
	TargetType    *string    `json:"targetType:omitempty"`
}

type Container_Bandwidth_GraphInputs struct {
	Entity

	EndDate            *time.Time `json:"endDate:omitempty"`
	NetworkInterfaceId *int       `json:"networkInterfaceId:omitempty"`
	Pod                *int       `json:"pod:omitempty"`
	ServerName         *string    `json:"serverName:omitempty"`
	StartDate          *time.Time `json:"startDate:omitempty"`
}

type Container_Bandwidth_GraphOutputs struct {
	Entity

	GraphImage   *[]byte    `json:"graphImage:omitempty"`
	GraphTitle   *string    `json:"graphTitle:omitempty"`
	MaxEndDate   *time.Time `json:"maxEndDate:omitempty"`
	MinStartDate *time.Time `json:"minStartDate:omitempty"`
}

type Container_Bandwidth_GraphOutputsExtended struct {
	Entity

	GraphImage         *[]byte    `json:"graphImage:omitempty"`
	GraphTitle         *string    `json:"graphTitle:omitempty"`
	InBoundTotalBytes  *uint      `json:"inBoundTotalBytes:omitempty"`
	MaxEndDate         *time.Time `json:"maxEndDate:omitempty"`
	MinStartDate       *time.Time `json:"minStartDate:omitempty"`
	OutBoundTotalBytes *uint      `json:"outBoundTotalBytes:omitempty"`
}

type Container_Bandwidth_Projection struct {
	Entity

	AllowedUsage   *string    `json:"allowedUsage:omitempty"`
	EstimatedUsage *string    `json:"estimatedUsage:omitempty"`
	HardwareId     *int       `json:"hardwareId:omitempty"`
	ProjectedUsage *string    `json:"projectedUsage:omitempty"`
	ServerName     *string    `json:"serverName:omitempty"`
	StartDate      *time.Time `json:"startDate:omitempty"`
}

type Container_Billing_Currency_Format struct {
	Entity

	Currency  *string  `json:"currency:omitempty"`
	Display   *int     `json:"display:omitempty"`
	Format    *string  `json:"format:omitempty"`
	Locale    *string  `json:"locale:omitempty"`
	Name      *string  `json:"name:omitempty"`
	Position  *int     `json:"position:omitempty"`
	Precision *int     `json:"precision:omitempty"`
	Script    *string  `json:"script:omitempty"`
	Service   *string  `json:"service:omitempty"`
	Symbol    *string  `json:"symbol:omitempty"`
	Tag       *string  `json:"tag:omitempty"`
	Value     *float64 `json:"value:omitempty"`
}

type Container_Billing_Info_Ach struct {
	Entity

	AccountNumber     *string `json:"accountNumber:omitempty"`
	AccountType       *string `json:"accountType:omitempty"`
	BankTransitNumber *string `json:"bankTransitNumber:omitempty"`
	City              *string `json:"city:omitempty"`
	Country           *string `json:"country:omitempty"`
	FederalTaxId      *string `json:"federalTaxId:omitempty"`
	FirstName         *string `json:"firstName:omitempty"`
	LastName          *string `json:"lastName:omitempty"`
	PhoneNumber       *string `json:"phoneNumber:omitempty"`
	PostalCode        *string `json:"postalCode:omitempty"`
	State             *string `json:"state:omitempty"`
	Street1           *string `json:"street1:omitempty"`
	Street2           *string `json:"street2:omitempty"`
}

type Container_Billing_Invoice_Email struct {
	Entity

	ExcelInvoiceIds       []int   `json:"excelInvoiceIds:omitempty"`
	PdfDetailedInvoiceIds []int   `json:"pdfDetailedInvoiceIds:omitempty"`
	PdfInvoiceIds         []int   `json:"pdfInvoiceIds:omitempty"`
	Type                  *string `json:"type:omitempty"`
}

type Container_Billing_Order_Status struct {
	Entity

	Description *string `json:"description:omitempty"`
	Status      *string `json:"status:omitempty"`
}

type Container_Catalyst_ManualEnrollmentRequest struct {
	Entity

	CustomerEmail          *string `json:"customerEmail:omitempty"`
	CustomerName           *string `json:"customerName:omitempty"`
	StartupName            *string `json:"startupName:omitempty"`
	VentureAffiliationFlag *bool   `json:"ventureAffiliationFlag:omitempty"`
	VentureFundName        *string `json:"ventureFundName:omitempty"`
}

type Container_Collection_Locale_CountryCode struct {
	Entity

	LongName   *string                                 `json:"longName:omitempty"`
	ShortName  *string                                 `json:"shortName:omitempty"`
	StateCodes []Container_Collection_Locale_StateCode `json:"stateCodes:omitempty"`
}

type Container_Collection_Locale_StateCode struct {
	Entity

	LongName  *string `json:"longName:omitempty"`
	ShortName *string `json:"shortName:omitempty"`
}

type Container_Disk_Image_Capture_Template struct {
	Entity

	Description *string                                        `json:"description:omitempty"`
	Name        *string                                        `json:"name:omitempty"`
	Summary     *string                                        `json:"summary:omitempty"`
	Volumes     []Container_Disk_Image_Capture_Template_Volume `json:"volumes:omitempty"`
}

type Container_Disk_Image_Capture_Template_Volume struct {
	Entity

	Name       *string                                                  `json:"name:omitempty"`
	Partitions []Container_Disk_Image_Capture_Template_Volume_Partition `json:"partitions:omitempty"`
}

type Container_Disk_Image_Capture_Template_Volume_Partition struct {
	Entity

	Name *string `json:"name:omitempty"`
}

type Container_Dns_Domain_Registration_Contact struct {
	Entity

	Address1         *string `json:"address1:omitempty"`
	Address2         *string `json:"address2:omitempty"`
	Address3         *string `json:"address3:omitempty"`
	City             *string `json:"city:omitempty"`
	Country          *string `json:"country:omitempty"`
	Email            *string `json:"email:omitempty"`
	Fax              *string `json:"fax:omitempty"`
	FirstName        *string `json:"firstName:omitempty"`
	LastName         *string `json:"lastName:omitempty"`
	OrganizationName *string `json:"organizationName:omitempty"`
	Phone            *string `json:"phone:omitempty"`
	PostalCode       *string `json:"postalCode:omitempty"`
	State            *string `json:"state:omitempty"`
	Type             *string `json:"type:omitempty"`
}

type Container_Dns_Domain_Registration_ExtendedAttribute struct {
	Entity

	ChildFlag       *bool                                                        `json:"childFlag:omitempty"`
	Description     *string                                                      `json:"description:omitempty"`
	Name            *string                                                      `json:"name:omitempty"`
	Options         []Container_Dns_Domain_Registration_ExtendedAttribute_Option `json:"options:omitempty"`
	RequiredFlag    *int                                                         `json:"requiredFlag:omitempty"`
	UserDefinedFlag *bool                                                        `json:"userDefinedFlag:omitempty"`
}

type Container_Dns_Domain_Registration_ExtendedAttribute_Configuration struct {
	Entity

	Name  *string `json:"name:omitempty"`
	Value *string `json:"value:omitempty"`
}

type Container_Dns_Domain_Registration_ExtendedAttribute_Option struct {
	Entity

	Description               *string                                                              `json:"description:omitempty"`
	RequireExtendedAttributes []Container_Dns_Domain_Registration_ExtendedAttribute_Option_Require `json:"requireExtendedAttributes:omitempty"`
	Title                     *string                                                              `json:"title:omitempty"`
	Value                     *string                                                              `json:"value:omitempty"`
}

type Container_Dns_Domain_Registration_ExtendedAttribute_Option_Require struct {
	Entity

	Name *string `json:"name:omitempty"`
}

type Container_Dns_Domain_Registration_Information struct {
	Entity

	Contacts           []Container_Dns_Domain_Registration_Contact    `json:"contacts:omitempty"`
	ExpireDate         *time.Time                                     `json:"expireDate:omitempty"`
	Nameservers        []Container_Dns_Domain_Registration_Nameserver `json:"nameservers:omitempty"`
	RegistryCreateDate *time.Time                                     `json:"registryCreateDate:omitempty"`
	RegistryExpireDate *time.Time                                     `json:"registryExpireDate:omitempty"`
	RegistryUpdateDate *time.Time                                     `json:"registryUpdateDate:omitempty"`
}

type Container_Dns_Domain_Registration_List struct {
	Entity

	DomainName                     *string                                                             `json:"domainName:omitempty"`
	EncodingType                   *string                                                             `json:"encodingType:omitempty"`
	ExtendedAttributeConfiguration []Container_Dns_Domain_Registration_ExtendedAttribute_Configuration `json:"extendedAttributeConfiguration:omitempty"`
	RegistrationPeriod             *int                                                                `json:"registrationPeriod:omitempty"`
}

type Container_Dns_Domain_Registration_Lookup struct {
	Entity

	Items []Container_Dns_Domain_Registration_Lookup_Items `json:"items:omitempty"`
}

type Container_Dns_Domain_Registration_Lookup_Items struct {
	Entity

	DomainName *string `json:"domainName:omitempty"`
	Status     *string `json:"status:omitempty"`
}

type Container_Dns_Domain_Registration_Nameserver struct {
	Entity

	Nameservers []Container_Dns_Domain_Registration_Nameserver_List `json:"nameservers:omitempty"`
}

type Container_Dns_Domain_Registration_Nameserver_List struct {
	Entity

	Ipv4Address *string `json:"ipv4Address:omitempty"`
	Ipv6Address *string `json:"ipv6Address:omitempty"`
	Name        *string `json:"name:omitempty"`
	SortOrder   *int    `json:"sortOrder:omitempty"`
}

type Container_Dns_Domain_Registration_Registrant_Verification_StatusDetail struct {
	Entity

	Status                   *Dns_Domain_Registration_Registrant_Verification_Status `json:"status:omitempty"`
	VerificationDeadlineDate *time.Time                                              `json:"verificationDeadlineDate:omitempty"`
}

type Container_Dns_Domain_Registration_Transfer_Information struct {
	Entity

	Reason          *string    `json:"reason:omitempty"`
	RegistrantEmail *string    `json:"registrantEmail:omitempty"`
	Status          *string    `json:"status:omitempty"`
	TimeStamp       *time.Time `json:"timeStamp:omitempty"`
	Transferrable   *int       `json:"transferrable:omitempty"`
}

type Container_Exception struct {
	Entity

	ExceptionClass   *string `json:"exceptionClass:omitempty"`
	ExceptionMessage *string `json:"exceptionMessage:omitempty"`
}

type Container_Graph struct {
	Entity

	BaseUnit      *string                      `json:"baseUnit:omitempty"`
	EndDatetime   *string                      `json:"endDatetime:omitempty"`
	Height        *int                         `json:"height:omitempty"`
	Image         *[]byte                      `json:"image:omitempty"`
	Interval      *int                         `json:"interval:omitempty"`
	Metrics       []Container_Metric_Data_Type `json:"metrics:omitempty"`
	NormalizeFlag *[]byte                      `json:"normalizeFlag:omitempty"`
	Options       []Container_Graph_Option     `json:"options:omitempty"`
	Plots         []Container_Graph_Plot       `json:"plots:omitempty"`
	ReturnImage   *bool                        `json:"returnImage:omitempty"`
	StartDatetime *string                      `json:"startDatetime:omitempty"`
	Template      *string                      `json:"template:omitempty"`
	Title         *string                      `json:"title:omitempty"`
	Width         *int                         `json:"width:omitempty"`
}

type Container_Graph_Option struct {
	Entity

	Name  *string `json:"name:omitempty"`
	Value *string `json:"value:omitempty"`
}

type Container_Graph_Plot struct {
	Entity

	Data   []Container_Graph_Plot_Coordinate `json:"data:omitempty"`
	Metric *Container_Metric_Data_Type       `json:"metric:omitempty"`
	Unit   *string                           `json:"unit:omitempty"`
}

type Container_Graph_Plot_Coordinate struct {
	Entity

	XValue *float64 `json:"xValue:omitempty"`
	YValue *float64 `json:"yValue:omitempty"`
	ZValue *float64 `json:"zValue:omitempty"`
}

type Container_Hardware_Configuration struct {
	Entity

	Datacenters               []Container_Hardware_Configuration_Option `json:"datacenters:omitempty"`
	FixedConfigurationPresets []Container_Hardware_Configuration_Option `json:"fixedConfigurationPresets:omitempty"`
	HardDrives                []Container_Hardware_Configuration_Option `json:"hardDrives:omitempty"`
	NetworkComponents         []Container_Hardware_Configuration_Option `json:"networkComponents:omitempty"`
	OperatingSystems          []Container_Hardware_Configuration_Option `json:"operatingSystems:omitempty"`
	Processors                []Container_Hardware_Configuration_Option `json:"processors:omitempty"`
}

type Container_Hardware_Configuration_Option struct {
	Entity

	ItemPrice *Product_Item_Price     `json:"itemPrice:omitempty"`
	Preset    *Product_Package_Preset `json:"preset:omitempty"`
	Template  *Hardware               `json:"template:omitempty"`
}

type Container_Hardware_MassUpdate struct {
	Entity

	HardwareId  *int    `json:"hardwareId:omitempty"`
	Message     *string `json:"message:omitempty"`
	SuccessFlag *string `json:"successFlag:omitempty"`
}

type Container_Hardware_Server_Configuration struct {
	Entity

	AddToSparePoolAfterOsReload *int                 `json:"addToSparePoolAfterOsReload:omitempty"`
	CustomProvisionScriptUri    *string              `json:"customProvisionScriptUri:omitempty"`
	DriveRetentionFlag          *bool                `json:"driveRetentionFlag:omitempty"`
	EraseHardDrives             *int                 `json:"eraseHardDrives:omitempty"`
	HardDrives                  []Hardware_Component `json:"hardDrives:omitempty"`
	ImageTemplateId             *int                 `json:"imageTemplateId:omitempty"`
	ItemPrices                  []Product_Item_Price `json:"itemPrices:omitempty"`
	ResetIpmiPassword           *int                 `json:"resetIpmiPassword:omitempty"`
	SshKeyIds                   []int                `json:"sshKeyIds:omitempty"`
	UpgradeBios                 *int                 `json:"upgradeBios:omitempty"`
	UpgradeHardDriveFirmware    *int                 `json:"upgradeHardDriveFirmware:omitempty"`
}

type Container_Hardware_Server_Details struct {
	Entity

	Components        []Hardware_Component `json:"components:omitempty"`
	NetworkComponents []Network_Component  `json:"networkComponents:omitempty"`
	Software          []Software_Component `json:"software:omitempty"`
}

type Container_KnowledgeLayer_QuestionAnswer struct {
	Entity

	Answer   *string `json:"answer:omitempty"`
	Link     *string `json:"link:omitempty"`
	Question *string `json:"question:omitempty"`
}

type Container_Message struct {
	Entity

	Message *string `json:"message:omitempty"`
	Type    *string `json:"type:omitempty"`
}

type Container_Metric_Data_Type struct {
	Entity

	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
	SummaryType *string `json:"summaryType:omitempty"`
	Unit        *string `json:"unit:omitempty"`
}

type Container_Metric_Tracking_Object_Details struct {
	Entity

	MetricName *string `json:"metricName:omitempty"`
}

type Container_Metric_Tracking_Object_Summary struct {
	Entity

	MetricName *string `json:"metricName:omitempty"`
}

type Container_Metric_Tracking_Object_Virtual_Host_Details struct {
	Container_Metric_Tracking_Object_Details

	Day             *time.Time `json:"day:omitempty"`
	MaxInstances    *int       `json:"maxInstances:omitempty"`
	MaxMemoryUsage  *int       `json:"maxMemoryUsage:omitempty"`
	MeanInstances   *float64   `json:"meanInstances:omitempty"`
	MeanMemoryUsage *float64   `json:"meanMemoryUsage:omitempty"`
	MinInstances    *int       `json:"minInstances:omitempty"`
	MinMemoryUsage  *int       `json:"minMemoryUsage:omitempty"`
}

type Container_Metric_Tracking_Object_Virtual_Host_Summary struct {
	Container_Metric_Tracking_Object_Summary

	AvgMemoryUsageInBillingCycle *int       `json:"avgMemoryUsageInBillingCycle:omitempty"`
	CurrentBillCycleEnd          *time.Time `json:"currentBillCycleEnd:omitempty"`
	CurrentBillCycleStart        *time.Time `json:"currentBillCycleStart:omitempty"`
	LastInstanceCount            *int       `json:"lastInstanceCount:omitempty"`
	LastMemoryUsageAmount        *int       `json:"lastMemoryUsageAmount:omitempty"`
	LastPollTime                 *time.Time `json:"lastPollTime:omitempty"`
	MaxInstanceInBillingCycle    *int       `json:"maxInstanceInBillingCycle:omitempty"`
	PreviousBillCycleEnd         *time.Time `json:"previousBillCycleEnd:omitempty"`
	PreviousBillCycleStart       *time.Time `json:"previousBillCycleStart:omitempty"`
	VirtualPlatformName          *string    `json:"virtualPlatformName:omitempty"`
}

type Container_Monitoring_Alarm_History struct {
	Entity

	AccountId  *int       `json:"accountId:omitempty"`
	AgentId    *int       `json:"agentId:omitempty"`
	AlarmId    *string    `json:"alarmId:omitempty"`
	ClosedDate *time.Time `json:"closedDate:omitempty"`
	CreateDate *time.Time `json:"createDate:omitempty"`
	Message    *string    `json:"message:omitempty"`
	RobotId    *int       `json:"robotId:omitempty"`
	Severity   *string    `json:"severity:omitempty"`
}

type Container_Monitoring_Graph_Outputs struct {
	Entity

	EndDate    *time.Time `json:"endDate:omitempty"`
	GraphError *string    `json:"graphError:omitempty"`
	GraphImage *[]byte    `json:"graphImage:omitempty"`
	StartDate  *time.Time `json:"startDate:omitempty"`
}

type Container_Network_Authentication_Data struct {
	Entity

	Host     *string `json:"host:omitempty"`
	Password *string `json:"password:omitempty"`
	Port     *int    `json:"port:omitempty"`
	Type     *string `json:"type:omitempty"`
	Username *string `json:"username:omitempty"`
}

type Container_Network_Bandwidth_Data_Summary struct {
	Entity

	AllowedUsage   *float64 `json:"allowedUsage:omitempty"`
	EstimatedUsage *float64 `json:"estimatedUsage:omitempty"`
	ProjectedUsage *float64 `json:"projectedUsage:omitempty"`
	UsageUnits     *string  `json:"usageUnits:omitempty"`
}

type Container_Network_Bandwidth_Version1_Usage struct {
	Entity

	IncomingAmount *float64   `json:"incomingAmount:omitempty"`
	OutgoingAmount *float64   `json:"outgoingAmount:omitempty"`
	RecordedDate   *time.Time `json:"recordedDate:omitempty"`
}

type Container_Network_ContentDelivery_Authentication_Directory struct {
	Entity

	CreateDate *time.Time `json:"createDate:omitempty"`
	Name       *string    `json:"name:omitempty"`
	Type       *string    `json:"type:omitempty"`
}

type Container_Network_ContentDelivery_Authentication_Parameter struct {
	Entity

	CdnAccountName *string `json:"cdnAccountName:omitempty"`
	ClientIp       *string `json:"clientIp:omitempty"`
	Referrer       *string `json:"referrer:omitempty"`
	SourceUrl      *string `json:"sourceUrl:omitempty"`
	Token          *string `json:"token:omitempty"`
}

type Container_Network_ContentDelivery_Authentication_ServiceEndpoint struct {
	Entity

	Endpoint *string `json:"endpoint:omitempty"`
	Protocol *string `json:"protocol:omitempty"`
}

type Container_Network_ContentDelivery_Bandwidth_PointsOfPresence_Summary struct {
	Entity

	Bandwidth     *uint      `json:"bandwidth:omitempty"`
	EndDateTime   *time.Time `json:"endDateTime:omitempty"`
	PopName       *string    `json:"popName:omitempty"`
	StartDateTime *time.Time `json:"startDateTime:omitempty"`
	UsageUnits    *string    `json:"usageUnits:omitempty"`
	ViewCount     *uint      `json:"viewCount:omitempty"`
}

type Container_Network_ContentDelivery_Bandwidth_Summary struct {
	Entity

	CdnAccountId  *int       `json:"cdnAccountId:omitempty"`
	EndDateTime   *time.Time `json:"endDateTime:omitempty"`
	FileName      *string    `json:"fileName:omitempty"`
	MediaType     *string    `json:"mediaType:omitempty"`
	StartDateTime *time.Time `json:"startDateTime:omitempty"`
	Usage         *float64   `json:"usage:omitempty"`
	UsageUnits    *string    `json:"usageUnits:omitempty"`
}

type Container_Network_ContentDelivery_Bandwidth_Summary_Detail struct {
	Container_Network_ContentDelivery_Bandwidth_Summary

	Duration  *float64 `json:"duration:omitempty"`
	ViewCount *int     `json:"viewCount:omitempty"`
}

type Container_Network_ContentDelivery_OriginPull_Mapping struct {
	Entity

	Cname           *string `json:"cname:omitempty"`
	Id              *string `json:"id:omitempty"`
	IsSecureContent *bool   `json:"isSecureContent:omitempty"`
	MediaType       *string `json:"mediaType:omitempty"`
	OriginUrl       *string `json:"originUrl:omitempty"`
}

type Container_Network_ContentDelivery_PointsOfPresence struct {
	Entity

	Id   *int    `json:"id:omitempty"`
	Name *string `json:"name:omitempty"`
}

type Container_Network_ContentDelivery_PurgeService_Response struct {
	Entity

	StatusCode *string `json:"statusCode:omitempty"`
	Url        *string `json:"url:omitempty"`
}

type Container_Network_ContentDelivery_Report_Usage struct {
	Entity

	ApplicationDeliveryNetwork    *float64   `json:"applicationDeliveryNetwork:omitempty"`
	ApplicationDeliveryNetworkSsl *float64   `json:"applicationDeliveryNetworkSsl:omitempty"`
	DiskSpace                     *float64   `json:"diskSpace:omitempty"`
	EndDate                       *time.Time `json:"endDate:omitempty"`
	Flash                         *float64   `json:"flash:omitempty"`
	Http                          *float64   `json:"http:omitempty"`
	HttpSmall                     *float64   `json:"httpSmall:omitempty"`
	Https                         *float64   `json:"https:omitempty"`
	HttpsSmall                    *float64   `json:"httpsSmall:omitempty"`
	Region                        *string    `json:"region:omitempty"`
	SslTotal                      *float64   `json:"sslTotal:omitempty"`
	StandardTotal                 *float64   `json:"standardTotal:omitempty"`
	StartDate                     *time.Time `json:"startDate:omitempty"`
	WindowsMedia                  *float64   `json:"windowsMedia:omitempty"`
}

type Container_Network_ContentDelivery_SupportedProtocol struct {
	Entity

	Host      *string `json:"host:omitempty"`
	MediaType *string `json:"mediaType:omitempty"`
	Platform  *string `json:"platform:omitempty"`
	Protocol  *string `json:"protocol:omitempty"`
}

type Container_Network_Directory_Listing struct {
	Entity

	FileCount *int    `json:"fileCount:omitempty"`
	Name      *string `json:"name:omitempty"`
	Type      *string `json:"type:omitempty"`
}

type Container_Network_IntrusionProtection_Event struct {
	Entity

	CVEId                 *string `json:"CVEId:omitempty"`
	ActionTaken           *string `json:"actionTaken:omitempty"`
	AttackCount           *int    `json:"attackCount:omitempty"`
	AttackLongDescription *string `json:"attackLongDescription:omitempty"`
	AttackName            *string `json:"attackName:omitempty"`
	BeginTime             *string `json:"beginTime:omitempty"`
	BugtraqId             *string `json:"bugtraqId:omitempty"`
	Classification        *string `json:"classification:omitempty"`
	DestinationIpAddress  *string `json:"destinationIpAddress:omitempty"`
	DestinationPort       *int    `json:"destinationPort:omitempty"`
	EndTime               *string `json:"endTime:omitempty"`
	Platform              *string `json:"platform:omitempty"`
	Protocol              *string `json:"protocol:omitempty"`
	Severity              *string `json:"severity:omitempty"`
	SignatureId           *string `json:"signatureId:omitempty"`
	SourceIpAddress       *string `json:"sourceIpAddress:omitempty"`
	SourcePort            *int    `json:"sourcePort:omitempty"`
}

type Container_Network_IntrusionProtection_Statistic struct {
	Entity

	AttackCount *int    `json:"attackCount:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Container_Network_IntrusionProtection_Statistics struct {
	Entity

	Target       *string                                           `json:"target:omitempty"`
	TargetType   *string                                           `json:"targetType:omitempty"`
	TimeFrame    *string                                           `json:"timeFrame:omitempty"`
	TopAttacks   []Container_Network_IntrusionProtection_Statistic `json:"topAttacks:omitempty"`
	TotalAttacks *int                                              `json:"totalAttacks:omitempty"`
}

type Container_Network_IntrusionProtection_SubnetReport struct {
	Entity

	Cidr            *int                                          `json:"cidr:omitempty"`
	Direction       *string                                       `json:"direction:omitempty"`
	Events          []Container_Network_IntrusionProtection_Event `json:"events:omitempty"`
	SubnetIpAddress *string                                       `json:"subnetIpAddress:omitempty"`
}

type Container_Network_LoadBalancer_StatusEntry struct {
	Entity

	Content *string `json:"content:omitempty"`
	Label   *string `json:"label:omitempty"`
}

type Container_Network_Media_Information struct {
	Entity

	AudioBitRate     *int     `json:"audioBitRate:omitempty"`
	AudioChannelMode *string  `json:"audioChannelMode:omitempty"`
	AudioChannels    *int     `json:"audioChannels:omitempty"`
	AudioCodec       *string  `json:"audioCodec:omitempty"`
	AudioSampleRate  *int     `json:"audioSampleRate:omitempty"`
	Duration         *float64 `json:"duration:omitempty"`
	ErrorMessage     *string  `json:"errorMessage:omitempty"`
	File             *string  `json:"file:omitempty"`
	FileFormat       *string  `json:"fileFormat:omitempty"`
	FileSize         *uint    `json:"fileSize:omitempty"`
	FrameRate        *float64 `json:"frameRate:omitempty"`
	SizeX            *int     `json:"sizeX:omitempty"`
	SizeY            *int     `json:"sizeY:omitempty"`
	TotalFrames      *uint    `json:"totalFrames:omitempty"`
	VideoAspectX     *float64 `json:"videoAspectX:omitempty"`
	VideoAspectY     *int     `json:"videoAspectY:omitempty"`
	VideoCodec       *string  `json:"videoCodec:omitempty"`
}

type Container_Network_Media_Transcode_Job_Watermark struct {
	Entity

	EndTime                *int                                                      `json:"endTime:omitempty"`
	FileName               *string                                                   `json:"fileName:omitempty"`
	Position               *Container_Network_Media_Transcode_Job_Watermark_Position `json:"position:omitempty"`
	StartTime              *int                                                      `json:"startTime:omitempty"`
	Text                   *string                                                   `json:"text:omitempty"`
	TransparencyPercentage *int                                                      `json:"transparencyPercentage:omitempty"`
}

type Container_Network_Media_Transcode_Job_Watermark_Position struct {
	Entity

	X *int `json:"x:omitempty"`
	Y *int `json:"y:omitempty"`
}

type Container_Network_Media_Transcode_Preset struct {
	Entity

	GUID        *string `json:"GUID:omitempty"`
	Category    *string `json:"category:omitempty"`
	Description *string `json:"description:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Container_Network_Media_Transcode_Preset_Element struct {
	Entity

	AdditionalElements  []Container_Network_Media_Transcode_Preset_Element_Option `json:"additionalElements:omitempty"`
	DefaultValue        *string                                                   `json:"defaultValue:omitempty"`
	Description         *string                                                   `json:"description:omitempty"`
	Enabled             *bool                                                     `json:"enabled:omitempty"`
	ExtendedDescription *string                                                   `json:"extendedDescription:omitempty"`
	Hidden              *bool                                                     `json:"hidden:omitempty"`
	MaximumValue        *int                                                      `json:"maximumValue:omitempty"`
	MinimumValue        *int                                                      `json:"minimumValue:omitempty"`
	Name                *string                                                   `json:"name:omitempty"`
	ParentName          *string                                                   `json:"parentName:omitempty"`
	Type                *string                                                   `json:"type:omitempty"`
}

type Container_Network_Media_Transcode_Preset_Element_Option struct {
	Entity

	Name  *string `json:"name:omitempty"`
	Value *string `json:"value:omitempty"`
}

type Container_Network_Message_Delivery_Email struct {
	Entity

	Body         *string `json:"body:omitempty"`
	ContainsHtml *bool   `json:"containsHtml:omitempty"`
	From         *string `json:"from:omitempty"`
	Subject      *string `json:"subject:omitempty"`
	To           *string `json:"to:omitempty"`
}

type Container_Network_Message_Delivery_Email_Sendgrid_Account_Overview struct {
	Entity

	CreditsAllowed *int    `json:"creditsAllowed:omitempty"`
	CreditsOverage *int    `json:"creditsOverage:omitempty"`
	CreditsRemain  *int    `json:"creditsRemain:omitempty"`
	CreditsUsed    *int    `json:"creditsUsed:omitempty"`
	Package        *string `json:"package:omitempty"`
	Reputation     *int    `json:"reputation:omitempty"`
	Requests       *int    `json:"requests:omitempty"`
}

type Container_Network_Message_Delivery_Email_Sendgrid_Customer_Profile struct {
	Entity

	Address   *string `json:"address:omitempty"`
	City      *string `json:"city:omitempty"`
	Country   *string `json:"country:omitempty"`
	Email     *string `json:"email:omitempty"`
	FirstName *string `json:"firstName:omitempty"`
	LastName  *string `json:"lastName:omitempty"`
	Phone     *string `json:"phone:omitempty"`
	State     *string `json:"state:omitempty"`
	Website   *string `json:"website:omitempty"`
	Zip       *string `json:"zip:omitempty"`
}

type Container_Network_Message_Delivery_Email_Sendgrid_List_Entry struct {
	Entity

	Created *string `json:"created:omitempty"`
	Email   *string `json:"email:omitempty"`
	Reason  *string `json:"reason:omitempty"`
	Status  *string `json:"status:omitempty"`
}

type Container_Network_Message_Delivery_Email_Sendgrid_Statistics struct {
	Entity

	Blocks             *int    `json:"blocks:omitempty"`
	Bounces            *int    `json:"bounces:omitempty"`
	Clicks             *int    `json:"clicks:omitempty"`
	Date               *string `json:"date:omitempty"`
	Delivered          *int    `json:"delivered:omitempty"`
	InvalidEmail       *int    `json:"invalidEmail:omitempty"`
	Opens              *int    `json:"opens:omitempty"`
	RepeatBounces      *int    `json:"repeatBounces:omitempty"`
	RepeatSpamReports  *int    `json:"repeatSpamReports:omitempty"`
	RepeatUnsubscribes *int    `json:"repeatUnsubscribes:omitempty"`
	Requests           *int    `json:"requests:omitempty"`
	SpamReports        *int    `json:"spamReports:omitempty"`
	UniqueClicks       *int    `json:"uniqueClicks:omitempty"`
	UniqueOpens        *int    `json:"uniqueOpens:omitempty"`
	Unsubscribes       *int    `json:"unsubscribes:omitempty"`
}

type Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Graph struct {
	Entity

	GraphError *string `json:"graphError:omitempty"`
	GraphImage *[]byte `json:"graphImage:omitempty"`
	GraphTitle *string `json:"graphTitle:omitempty"`
}

type Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Options struct {
	Entity

	AggregatesOnly     *bool      `json:"aggregatesOnly:omitempty"`
	Category           *string    `json:"category:omitempty"`
	Days               *int       `json:"days:omitempty"`
	EndDate            *time.Time `json:"endDate:omitempty"`
	SelectedStatistics []string   `json:"selectedStatistics:omitempty"`
	StartDate          *time.Time `json:"startDate:omitempty"`
}

type Container_Network_Port_Statistic struct {
	Entity

	AdministrativeStatus    *int  `json:"administrativeStatus:omitempty"`
	InDiscardPackets        *uint `json:"inDiscardPackets:omitempty"`
	InErrorPackets          *uint `json:"inErrorPackets:omitempty"`
	InOctets                *uint `json:"inOctets:omitempty"`
	InUnicastPackets        *uint `json:"inUnicastPackets:omitempty"`
	MaximumTransmissionUnit *uint `json:"maximumTransmissionUnit:omitempty"`
	OperationalStatus       *int  `json:"operationalStatus:omitempty"`
	OutDiscardPackets       *uint `json:"outDiscardPackets:omitempty"`
	OutErrorPackets         *uint `json:"outErrorPackets:omitempty"`
	OutOctets               *uint `json:"outOctets:omitempty"`
	OutUnicastPackets       *uint `json:"outUnicastPackets:omitempty"`
	PortDuplex              *uint `json:"portDuplex:omitempty"`
	Speed                   *uint `json:"speed:omitempty"`
}

type Container_Network_Service_Resource_ObjectStorage_ConnectionInformation struct {
	Entity

	Datacenter          *string `json:"datacenter:omitempty"`
	DatacenterShortName *string `json:"datacenterShortName:omitempty"`
	PrivateEndpoint     *string `json:"privateEndpoint:omitempty"`
	PublicEndpoint      *string `json:"publicEndpoint:omitempty"`
}

type Container_Network_Storage_Backup_Evault_WebCc_Authentication_Details struct {
	Entity

	EventValidation *string `json:"eventValidation:omitempty"`
	ViewState       *string `json:"viewState:omitempty"`
	WebCcUrl        *string `json:"webCcUrl:omitempty"`
}

type Container_Network_Storage_Evault_Vault_Task struct {
	Entity

	Id           *uint   `json:"id:omitempty"`
	Name         *string `json:"name:omitempty"`
	UsedPoolsize *uint   `json:"usedPoolsize:omitempty"`
}

type Container_Network_Storage_Evault_WebCc_AgentStatus struct {
	Entity

	LastBackup *time.Time `json:"lastBackup:omitempty"`
	Status     *string    `json:"status:omitempty"`
}

type Container_Network_Storage_Evault_WebCc_BackupResults struct {
	Entity

	BeginTime *time.Time `json:"beginTime:omitempty"`
	Conflict  *string    `json:"conflict:omitempty"`
	EndTime   *time.Time `json:"endTime:omitempty"`
	Failed    *string    `json:"failed:omitempty"`
	Success   *string    `json:"success:omitempty"`
}

type Container_Network_Storage_Evault_WebCc_JobDetails struct {
	Entity

	BytesUsed              *uint      `json:"bytesUsed:omitempty"`
	Description            *string    `json:"description:omitempty"`
	HardwareId             *int       `json:"hardwareId:omitempty"`
	LastRunDate            *time.Time `json:"lastRunDate:omitempty"`
	Name                   *string    `json:"name:omitempty"`
	OriginalSize           *uint      `json:"originalSize:omitempty"`
	PercentageOfTotalUsage *int       `json:"percentageOfTotalUsage:omitempty"`
	Result                 *string    `json:"result:omitempty"`
	VirtualGuestId         *int       `json:"virtualGuestId:omitempty"`
}

type Container_Network_Storage_Host struct {
	Entity

	Id         *int    `json:"id:omitempty"`
	ObjectType *string `json:"objectType:omitempty"`
}

type Container_Network_Storage_Hub_ObjectStorage_ContentDeliveryUrl struct {
	Entity

	Datacenter *string `json:"datacenter:omitempty"`
	FlashUrl   *string `json:"flashUrl:omitempty"`
	HttpUrl    *string `json:"httpUrl:omitempty"`
}

type Container_Network_Storage_Hub_ObjectStorage_Endpoint struct {
	Entity

	Location *string `json:"location:omitempty"`
	Region   *string `json:"region:omitempty"`
	Type     *string `json:"type:omitempty"`
	Url      *string `json:"url:omitempty"`
}

type Container_Network_Storage_Hub_ObjectStorage_File struct {
	Container_Utility_File_Entity

	Folder *string `json:"folder:omitempty"`
	Hash   *string `json:"hash:omitempty"`
}

type Container_Network_Storage_Hub_ObjectStorage_Folder struct {
	Entity

	Bytes *uint   `json:"bytes:omitempty"`
	Count *uint   `json:"count:omitempty"`
	Name  *string `json:"name:omitempty"`
}

type Container_Network_Storage_Hub_ObjectStorage_Node struct {
	Entity

	DeviceName   *string `json:"deviceName:omitempty"`
	ResourceName *string `json:"resourceName:omitempty"`
	UserAuthUrl  *string `json:"userAuthUrl:omitempty"`
}

type Container_Network_Storage_NetworkConnectionInformation struct {
	Entity

	Id          *string `json:"id:omitempty"`
	IpAddress   *string `json:"ipAddress:omitempty"`
	StorageType *string `json:"storageType:omitempty"`
}

type Container_Network_Subnet_IpAddress struct {
	Entity

	Hardware           *Hardware `json:"hardware:omitempty"`
	IpAddress          *string   `json:"ipAddress:omitempty"`
	IsBroadcastAddress *bool     `json:"isBroadcastAddress:omitempty"`
	IsGatewayAddress   *bool     `json:"isGatewayAddress:omitempty"`
	IsNetworkAddress   *bool     `json:"isNetworkAddress:omitempty"`
}

type Container_Network_Subnet_Registration_SubnetReference struct {
	Entity

	RegistrationId *int    `json:"registrationId:omitempty"`
	SubnetCidr     *string `json:"subnetCidr:omitempty"`
}

type Container_Network_Subnet_Registration_TransactionDetails struct {
	Entity

	SubnetReferences []Container_Network_Subnet_Registration_SubnetReference `json:"subnetReferences:omitempty"`
	TransactionId    *int                                                    `json:"transactionId:omitempty"`
}

type Container_Notification_Mass_Filter_TemplateKey struct {
	Entity
}

type Container_Notification_Mass_Filter_TemplateValue struct {
	Entity
}

type Container_Policy_Acceptance struct {
	Entity

	AcceptedFlag              *bool   `json:"acceptedFlag:omitempty"`
	PolicyName                *string `json:"policyName:omitempty"`
	ProductPolicyAssignmentId *int    `json:"productPolicyAssignmentId:omitempty"`
}

type Container_Product_Item_Category struct {
	Entity

	Id *int `json:"id:omitempty"`
}

type Container_Product_Item_Category_Question_Answer struct {
	Entity

	Answer       *string `json:"answer:omitempty"`
	CategoryCode *string `json:"categoryCode:omitempty"`
	CategoryId   *int    `json:"categoryId:omitempty"`
	QuestionId   *int    `json:"questionId:omitempty"`
}

type Container_Product_Item_Category_ZeroFee_Count struct {
	Entity

	CategoryCode *string `json:"categoryCode:omitempty"`
	CategoryId   *int    `json:"categoryId:omitempty"`
	CategoryName *string `json:"categoryName:omitempty"`
	Count        *int    `json:"count:omitempty"`
}

type Container_Product_Item_Discount_Program struct {
	Entity

	ApplicableQuantity      *int                 `json:"applicableQuantity:omitempty"`
	Item                    *Product_Item        `json:"item:omitempty"`
	OneTimeAmount           *float64             `json:"oneTimeAmount:omitempty"`
	OneTimeTax              *float64             `json:"oneTimeTax:omitempty"`
	Prices                  []Product_Item_Price `json:"prices:omitempty"`
	ProratedOneTimeAmount   *float64             `json:"proratedOneTimeAmount:omitempty"`
	ProratedOneTimeTax      *float64             `json:"proratedOneTimeTax:omitempty"`
	ProratedRecurringAmount *float64             `json:"proratedRecurringAmount:omitempty"`
	ProratedRecurringTax    *float64             `json:"proratedRecurringTax:omitempty"`
	RecurringAmount         *float64             `json:"recurringAmount:omitempty"`
	RecurringTax            *float64             `json:"recurringTax:omitempty"`
}

type Container_Product_Order struct {
	Entity

	BigDataOrderFlag              *bool                                             `json:"bigDataOrderFlag:omitempty"`
	BillingInformation            *Container_Product_Order_Billing_Information      `json:"billingInformation:omitempty"`
	BillingOrderItemId            *int                                              `json:"billingOrderItemId:omitempty"`
	CancelUrl                     *string                                           `json:"cancelUrl:omitempty"`
	ContainerIdentifier           *string                                           `json:"containerIdentifier:omitempty"`
	ContainerSplHash              *string                                           `json:"containerSplHash:omitempty"`
	CurrencyShortName             *string                                           `json:"currencyShortName:omitempty"`
	DeviceFingerprintId           *string                                           `json:"deviceFingerprintId:omitempty"`
	DisplayLayerSessionId         *string                                           `json:"displayLayerSessionId:omitempty"`
	ExtendedHardwareTesting       *bool                                             `json:"extendedHardwareTesting:omitempty"`
	FlexibleCreditProgramPrice    *Product_Item_Price                               `json:"flexibleCreditProgramPrice:omitempty"`
	Hardware                      []Hardware                                        `json:"hardware:omitempty"`
	ImageTemplateGlobalIdentifier *string                                           `json:"imageTemplateGlobalIdentifier:omitempty"`
	ImageTemplateId               *int                                              `json:"imageTemplateId:omitempty"`
	IsManagedOrder                *int                                              `json:"isManagedOrder:omitempty"`
	ItemCategoryQuestionAnswers   []Container_Product_Item_Category_Question_Answer `json:"itemCategoryQuestionAnswers:omitempty"`
	Location                      *string                                           `json:"location:omitempty"`
	LocationObject                *Location                                         `json:"locationObject:omitempty"`
	Message                       *string                                           `json:"message:omitempty"`
	OrderContainers               []Container_Product_Order                         `json:"orderContainers:omitempty"`
	OrderHostnames                []string                                          `json:"orderHostnames:omitempty"`
	OrderVerificationExceptions   []Container_Exception                             `json:"orderVerificationExceptions:omitempty"`
	PackageId                     *int                                              `json:"packageId:omitempty"`
	PaymentType                   *string                                           `json:"paymentType:omitempty"`
	PostTaxRecurring              *float64                                          `json:"postTaxRecurring:omitempty"`
	PostTaxRecurringHourly        *float64                                          `json:"postTaxRecurringHourly:omitempty"`
	PostTaxRecurringMonthly       *float64                                          `json:"postTaxRecurringMonthly:omitempty"`
	PostTaxSetup                  *float64                                          `json:"postTaxSetup:omitempty"`
	PreTaxRecurring               *float64                                          `json:"preTaxRecurring:omitempty"`
	PreTaxRecurringHourly         *float64                                          `json:"preTaxRecurringHourly:omitempty"`
	PreTaxRecurringMonthly        *float64                                          `json:"preTaxRecurringMonthly:omitempty"`
	PreTaxSetup                   *float64                                          `json:"preTaxSetup:omitempty"`
	PresaleEvent                  *Sales_Presale_Event                              `json:"presaleEvent:omitempty"`
	PresetId                      *int                                              `json:"presetId:omitempty"`
	Prices                        []Product_Item_Price                              `json:"prices:omitempty"`
	PrimaryDiskPartitionId        *int                                              `json:"primaryDiskPartitionId:omitempty"`
	Priorities                    []string                                          `json:"priorities:omitempty"`
	PrivateCloudOrderFlag         *bool                                             `json:"privateCloudOrderFlag:omitempty"`
	PrivateCloudOrderType         *string                                           `json:"privateCloudOrderType:omitempty"`
	PromotionCode                 *string                                           `json:"promotionCode:omitempty"`
	Properties                    []Container_Product_Order_Property                `json:"properties:omitempty"`
	ProratedInitialCharge         *float64                                          `json:"proratedInitialCharge:omitempty"`
	ProratedOrderTotal            *float64                                          `json:"proratedOrderTotal:omitempty"`
	ProvisionScripts              []string                                          `json:"provisionScripts:omitempty"`
	Quantity                      *int                                              `json:"quantity:omitempty"`
	QuoteName                     *string                                           `json:"quoteName:omitempty"`
	RegionalGroup                 *string                                           `json:"regionalGroup:omitempty"`
	ResourceGroupId               *int                                              `json:"resourceGroupId:omitempty"`
	ResourceGroupName             *string                                           `json:"resourceGroupName:omitempty"`
	ResourceGroupTemplateId       *int                                              `json:"resourceGroupTemplateId:omitempty"`
	ReturnUrl                     *string                                           `json:"returnUrl:omitempty"`
	SendQuoteEmailFlag            *bool                                             `json:"sendQuoteEmailFlag:omitempty"`
	ServerCoreCount               *int                                              `json:"serverCoreCount:omitempty"`
	SourceVirtualGuestId          *int                                              `json:"sourceVirtualGuestId:omitempty"`
	SshKeys                       []Container_Product_Order_SshKeys                 `json:"sshKeys:omitempty"`
	StepId                        *int                                              `json:"stepId:omitempty"`
	StorageGroups                 []Container_Product_Order_Storage_Group           `json:"storageGroups:omitempty"`
	TaxCacheHash                  *string                                           `json:"taxCacheHash:omitempty"`
	TaxCompletedFlag              *bool                                             `json:"taxCompletedFlag:omitempty"`
	TechIncubatorItemPrice        *Product_Item_Price                               `json:"techIncubatorItemPrice:omitempty"`
	TotalRecurringTax             *float64                                          `json:"totalRecurringTax:omitempty"`
	TotalSetupTax                 *float64                                          `json:"totalSetupTax:omitempty"`
	UseHourlyPricing              *bool                                             `json:"useHourlyPricing:omitempty"`
	VirtualGuests                 []Virtual_Guest                                   `json:"virtualGuests:omitempty"`
}

type Container_Product_Order_Account_Media_Data_Transfer_Request struct {
	Container_Product_Order

	Request *Account_Media_Data_Transfer_Request `json:"request:omitempty"`
}

type Container_Product_Order_Attribute_Address struct {
	Entity

	AddressLine1 *string `json:"addressLine1:omitempty"`
	AddressLine2 *string `json:"addressLine2:omitempty"`
	City         *string `json:"city:omitempty"`
	CountryCode  *string `json:"countryCode:omitempty"`
	NonUsState   *string `json:"nonUsState:omitempty"`
	PostalCode   *string `json:"postalCode:omitempty"`
	State        *string `json:"state:omitempty"`
}

type Container_Product_Order_Attribute_Contact struct {
	Entity

	Address          *Container_Product_Order_Attribute_Address `json:"address:omitempty"`
	EmailAddress     *string                                    `json:"emailAddress:omitempty"`
	FaxNumber        *string                                    `json:"faxNumber:omitempty"`
	FirstName        *string                                    `json:"firstName:omitempty"`
	LastName         *string                                    `json:"lastName:omitempty"`
	OrganizationName *string                                    `json:"organizationName:omitempty"`
	PhoneNumber      *string                                    `json:"phoneNumber:omitempty"`
	Title            *string                                    `json:"title:omitempty"`
}

type Container_Product_Order_Attribute_Organization struct {
	Entity

	Address          *Container_Product_Order_Attribute_Address `json:"address:omitempty"`
	FaxNumber        *string                                    `json:"faxNumber:omitempty"`
	OrganizationName *string                                    `json:"organizationName:omitempty"`
	PhoneNumber      *string                                    `json:"phoneNumber:omitempty"`
}

type Container_Product_Order_Billing_Information struct {
	Entity

	BillingAddressLine1          *string `json:"billingAddressLine1:omitempty"`
	BillingAddressLine2          *string `json:"billingAddressLine2:omitempty"`
	BillingCity                  *string `json:"billingCity:omitempty"`
	BillingCountryCode           *string `json:"billingCountryCode:omitempty"`
	BillingEmail                 *string `json:"billingEmail:omitempty"`
	BillingNameCompany           *string `json:"billingNameCompany:omitempty"`
	BillingNameFirst             *string `json:"billingNameFirst:omitempty"`
	BillingNameLast              *string `json:"billingNameLast:omitempty"`
	BillingPhoneFax              *string `json:"billingPhoneFax:omitempty"`
	BillingPhoneVoice            *string `json:"billingPhoneVoice:omitempty"`
	BillingPostalCode            *string `json:"billingPostalCode:omitempty"`
	BillingState                 *string `json:"billingState:omitempty"`
	CardAccountNumber            *string `json:"cardAccountNumber:omitempty"`
	CardExpirationMonth          *int    `json:"cardExpirationMonth:omitempty"`
	CardExpirationYear           *int    `json:"cardExpirationYear:omitempty"`
	CreditCardVerificationNumber *string `json:"creditCardVerificationNumber:omitempty"`
	TaxExempt                    *int    `json:"taxExempt:omitempty"`
	VatId                        *string `json:"vatId:omitempty"`
}

type Container_Product_Order_Dns_Domain_Registration struct {
	Container_Product_Order

	AdministrativeContact  *Container_Dns_Domain_Registration_Contact `json:"administrativeContact:omitempty"`
	BillingContact         *Container_Dns_Domain_Registration_Contact `json:"billingContact:omitempty"`
	DomainRegistrationList []Container_Dns_Domain_Registration_List   `json:"domainRegistrationList:omitempty"`
	OwnerContact           *Container_Dns_Domain_Registration_Contact `json:"ownerContact:omitempty"`
	RegistrationType       *string                                    `json:"registrationType:omitempty"`
	TechnicalContact       *Container_Dns_Domain_Registration_Contact `json:"technicalContact:omitempty"`
}

type Container_Product_Order_Dns_Domain_Reseller struct {
	Container_Product_Order

	CreditAmount *float64 `json:"creditAmount:omitempty"`
}

type Container_Product_Order_Gateway_Appliance_Cluster struct {
	Container_Product_Order

	ClusterIdentifier *string `json:"clusterIdentifier:omitempty"`
	ClusterOrderType  *string `json:"clusterOrderType:omitempty"`
}

type Container_Product_Order_Hardware_Security_Module struct {
	Container_Product_Order_Hardware_Server
}

type Container_Product_Order_Hardware_Server struct {
	Container_Product_Order

	ClusterIdentifier                           *string                            `json:"clusterIdentifier:omitempty"`
	ClusterOrderType                            *string                            `json:"clusterOrderType:omitempty"`
	ClusterResourceId                           *int                               `json:"clusterResourceId:omitempty"`
	MonitoringAgentConfigurationTemplateGroupId *int                               `json:"monitoringAgentConfigurationTemplateGroupId:omitempty"`
	PrivateCloudServerRole                      *string                            `json:"privateCloudServerRole:omitempty"`
	RequiredUpstreamDeviceId                    *int                               `json:"requiredUpstreamDeviceId:omitempty"`
	Tags                                        []Container_Product_Order_Property `json:"tags:omitempty"`
}

type Container_Product_Order_Hardware_Server_Colocation struct {
	Container_Product_Order_Hardware_Server
}

type Container_Product_Order_Hardware_Server_Gateway_Appliance struct {
	Container_Product_Order_Hardware_Server
}

type Container_Product_Order_Hardware_Server_Upgrade struct {
	Container_Product_Order_Hardware_Server
}

type Container_Product_Order_Monitoring_Package struct {
	Container_Product_Order

	ConfigurationTemplateGroups []Monitoring_Agent_Configuration_Template_Group `json:"configurationTemplateGroups:omitempty"`
	ServerType                  *string                                         `json:"serverType:omitempty"`
}

type Container_Product_Order_MultiConfiguration struct {
	Container_Product_Order
}

type Container_Product_Order_MultiConfiguration_Tornado struct {
	Container_Product_Order_MultiConfiguration
}

type Container_Product_Order_Network struct {
	Entity

	Network     *Network                  `json:"network:omitempty"`
	PublicVlans []Container_Product_Order `json:"publicVlans:omitempty"`
	Subnets     []Container_Product_Order `json:"subnets:omitempty"`
}

type Container_Product_Order_Network_Application_Delivery_Controller struct {
	Container_Product_Order

	ApplicationDeliveryControllerId *int `json:"applicationDeliveryControllerId:omitempty"`
}

type Container_Product_Order_Network_ContentDelivery_Account struct {
	Container_Product_Order

	CdnAccountName *string `json:"cdnAccountName:omitempty"`
}

type Container_Product_Order_Network_ContentDelivery_Account_Upgrade struct {
	Container_Product_Order

	CdnAccountId *string `json:"cdnAccountId:omitempty"`
}

type Container_Product_Order_Network_LoadBalancer struct {
	Container_Product_Order
}

type Container_Product_Order_Network_LoadBalancer_Global struct {
	Container_Product_Order

	Domain   *string `json:"domain:omitempty"`
	Hostname *string `json:"hostname:omitempty"`
}

type Container_Product_Order_Network_Message_Delivery struct {
	Container_Product_Order

	AccountPassword *string `json:"accountPassword:omitempty"`
	AccountUsername *string `json:"accountUsername:omitempty"`
	EmailAddress    *string `json:"emailAddress:omitempty"`
}

type Container_Product_Order_Network_Message_Queue struct {
	Container_Product_Order
}

type Container_Product_Order_Network_PerformanceStorage struct {
	Container_Product_Order
}

type Container_Product_Order_Network_PerformanceStorage_Iscsi struct {
	Container_Product_Order_Network_PerformanceStorage

	OsFormatType *Network_Storage_Iscsi_OS_Type `json:"osFormatType:omitempty"`
}

type Container_Product_Order_Network_PerformanceStorage_Nfs struct {
	Container_Product_Order_Network_PerformanceStorage
}

type Container_Product_Order_Network_Protection_Firewall struct {
	Container_Product_Order
}

type Container_Product_Order_Network_Protection_Firewall_Dedicated struct {
	Container_Product_Order

	Vlan   *Network_Vlan `json:"vlan:omitempty"`
	VlanId *int          `json:"vlanId:omitempty"`
}

type Container_Product_Order_Network_Storage_Backup_Evault_Plugin struct {
	Container_Product_Order
}

type Container_Product_Order_Network_Storage_Backup_Evault_Vault struct {
	Container_Product_Order
}

type Container_Product_Order_Network_Storage_Enterprise struct {
	Container_Product_Order

	OriginVolumeId         *int                           `json:"originVolumeId:omitempty"`
	OriginVolumeScheduleId *int                           `json:"originVolumeScheduleId:omitempty"`
	OsFormatType           *Network_Storage_Iscsi_OS_Type `json:"osFormatType:omitempty"`
}

type Container_Product_Order_Network_Storage_Enterprise_SnapshotSpace struct {
	Container_Product_Order

	VolumeId *int `json:"volumeId:omitempty"`
}

type Container_Product_Order_Network_Storage_Enterprise_SnapshotSpace_Upgrade struct {
	Container_Product_Order_Network_Storage_Enterprise_SnapshotSpace
}

type Container_Product_Order_Network_Storage_Hub struct {
	Container_Product_Order
}

type Container_Product_Order_Network_Storage_Hub_Datacenter struct {
	Entity

	Location        *Location            `json:"location:omitempty"`
	UsageRatePrices []Product_Item_Price `json:"usageRatePrices:omitempty"`
}

type Container_Product_Order_Network_Storage_Iscsi struct {
	Container_Product_Order
}

type Container_Product_Order_Network_Storage_Iscsi_Replication struct {
	Container_Product_Order

	VolumeId *int `json:"volumeId:omitempty"`
}

type Container_Product_Order_Network_Storage_Iscsi_SnapshotSpace struct {
	Container_Product_Order

	VolumeId *int `json:"volumeId:omitempty"`
}

type Container_Product_Order_Network_Storage_Modification struct {
	Container_Product_Order

	VolumeId *int `json:"volumeId:omitempty"`
}

type Container_Product_Order_Network_Storage_Nas struct {
	Container_Product_Order
}

type Container_Product_Order_Network_Storage_Object struct {
	Container_Product_Order
}

type Container_Product_Order_Network_Subnet struct {
	Container_Product_Order

	Description         *string `json:"description:omitempty"`
	EndPointIpAddressId *int    `json:"endPointIpAddressId:omitempty"`
	EndPointVlanId      *int    `json:"endPointVlanId:omitempty"`
	Id                  *int    `json:"id:omitempty"`
	RouterHostname      *string `json:"routerHostname:omitempty"`
}

type Container_Product_Order_Network_Tunnel_Ipsec struct {
	Container_Product_Order
}

type Container_Product_Order_Network_Vlan struct {
	Container_Product_Order

	Description        *string                   `json:"description:omitempty"`
	HostnameDatacenter *string                   `json:"hostnameDatacenter:omitempty"`
	HostnameRouter     *string                   `json:"hostnameRouter:omitempty"`
	Id                 *int                      `json:"id:omitempty"`
	Name               *string                   `json:"name:omitempty"`
	Router             *Hardware                 `json:"router:omitempty"`
	RouterId           *int                      `json:"routerId:omitempty"`
	Subnets            []Container_Product_Order `json:"subnets:omitempty"`
	VlanNumber         *int                      `json:"vlanNumber:omitempty"`
}

type Container_Product_Order_Network_Vlans struct {
	Entity

	PrivateVlans []Container_Product_Order `json:"privateVlans:omitempty"`
	PublicVlans  []Container_Product_Order `json:"publicVlans:omitempty"`
}

type Container_Product_Order_Property struct {
	Entity

	Name  *string `json:"name:omitempty"`
	Value *string `json:"value:omitempty"`
}

type Container_Product_Order_Receipt struct {
	Entity

	ExternalPaymentCheckoutUrl *string                  `json:"externalPaymentCheckoutUrl:omitempty"`
	ExternalPaymentToken       *string                  `json:"externalPaymentToken:omitempty"`
	OrderDate                  *time.Time               `json:"orderDate:omitempty"`
	OrderDetails               *Container_Product_Order `json:"orderDetails:omitempty"`
	OrderId                    *int                     `json:"orderId:omitempty"`
	PaypalCheckoutUrl          *string                  `json:"paypalCheckoutUrl:omitempty"`
	PaypalToken                *string                  `json:"paypalToken:omitempty"`
	PlacedOrder                *Billing_Order           `json:"placedOrder:omitempty"`
	Quote                      *Billing_Order_Quote     `json:"quote:omitempty"`
}

type Container_Product_Order_Security_Certificate struct {
	Container_Product_Order

	AdministrativeContact     *Container_Product_Order_Attribute_Contact      `json:"administrativeContact:omitempty"`
	BillingContact            *Container_Product_Order_Attribute_Contact      `json:"billingContact:omitempty"`
	CertificateSigningRequest *string                                         `json:"certificateSigningRequest:omitempty"`
	OrderApproverEmailAddress *string                                         `json:"orderApproverEmailAddress:omitempty"`
	OrganizationInformation   *Container_Product_Order_Attribute_Organization `json:"organizationInformation:omitempty"`
	RenewalFlag               *bool                                           `json:"renewalFlag:omitempty"`
	ServerCount               *int                                            `json:"serverCount:omitempty"`
	ServerType                *string                                         `json:"serverType:omitempty"`
	TechnicalContact          *Container_Product_Order_Attribute_Contact      `json:"technicalContact:omitempty"`
	ValidityMonths            *int                                            `json:"validityMonths:omitempty"`
}

type Container_Product_Order_Software_Component_Virtual struct {
	Container_Product_Order

	EndPointIpAddressIds []int `json:"endPointIpAddressIds:omitempty"`
}

type Container_Product_Order_Software_License struct {
	Container_Product_Order
}

type Container_Product_Order_SshKeys struct {
	Entity

	SshKeyIds []int `json:"sshKeyIds:omitempty"`
}

type Container_Product_Order_Storage_Group struct {
	Entity

	ArraySize           *float64                                          `json:"arraySize:omitempty"`
	ArrayTypeId         *int                                              `json:"arrayTypeId:omitempty"`
	HardDrives          []int                                             `json:"hardDrives:omitempty"`
	HotSpareDrives      []int                                             `json:"hotSpareDrives:omitempty"`
	PartitionTemplateId *int                                              `json:"partitionTemplateId:omitempty"`
	Partitions          []Container_Product_Order_Storage_Group_Partition `json:"partitions:omitempty"`
}

type Container_Product_Order_Storage_Group_Partition struct {
	Entity

	IsGrow *bool    `json:"isGrow:omitempty"`
	Name   *string  `json:"name:omitempty"`
	Size   *float64 `json:"size:omitempty"`
}

type Container_Product_Order_User_Customer_External_Binding struct {
	Container_Product_Order

	ExternalId *string `json:"externalId:omitempty"`
	UserId     *int    `json:"userId:omitempty"`
	VendorId   *int    `json:"vendorId:omitempty"`
}

type Container_Product_Order_Virtual_Disk_Image struct {
	Container_Product_Order

	DiskDescription *string `json:"diskDescription:omitempty"`
}

type Container_Product_Order_Virtual_Guest struct {
	Container_Product_Order_Hardware_Server

	BootableDiskId *int `json:"bootableDiskId:omitempty"`
}

type Container_Product_Order_Virtual_Guest_Upgrade struct {
	Container_Product_Order_Virtual_Guest
}

type Container_Provisioning_Maintenance_Window struct {
	Entity

	ClassificationIds     []Provisioning_Maintenance_Classification `json:"classificationIds:omitempty"`
	ItemCategoryIds       []Product_Item_Category                   `json:"itemCategoryIds:omitempty"`
	MaintenanceWindowId   *int                                      `json:"maintenanceWindowId:omitempty"`
	TicketId              *int                                      `json:"ticketId:omitempty"`
	WindowMaintenanceDate *time.Time                                `json:"windowMaintenanceDate:omitempty"`
}

type Container_Referral_Partner_Commission struct {
	Entity

	CommissionAmount         *float64   `json:"commissionAmount:omitempty"`
	CommissionRate           *float64   `json:"commissionRate:omitempty"`
	CreateDate               *time.Time `json:"createDate:omitempty"`
	ReferralAccountId        *int       `json:"referralAccountId:omitempty"`
	ReferralCompanyName      *string    `json:"referralCompanyName:omitempty"`
	ReferralPartnerAccountId *int       `json:"referralPartnerAccountId:omitempty"`
	ReferralRevenue          *float64   `json:"referralRevenue:omitempty"`
}

type Container_Referral_Partner_Payment_Option struct {
	Entity

	AccountNumber     *string `json:"accountNumber:omitempty"`
	AccountType       *string `json:"accountType:omitempty"`
	Address1          *string `json:"address1:omitempty"`
	Address2          *string `json:"address2:omitempty"`
	BankTransitNumber *string `json:"bankTransitNumber:omitempty"`
	City              *string `json:"city:omitempty"`
	CompanyName       *string `json:"companyName:omitempty"`
	Country           *string `json:"country:omitempty"`
	FederalTaxId      *string `json:"federalTaxId:omitempty"`
	FirstName         *string `json:"firstName:omitempty"`
	LastName          *string `json:"lastName:omitempty"`
	PaymentType       *string `json:"paymentType:omitempty"`
	PaypalEmail       *string `json:"paypalEmail:omitempty"`
	PhoneNumber       *string `json:"phoneNumber:omitempty"`
	PostalCode        *string `json:"postalCode:omitempty"`
	State             *string `json:"state:omitempty"`
}

type Container_Referral_Partner_Prospect struct {
	Entity

	Address1    *string           `json:"address1:omitempty"`
	Address2    *string           `json:"address2:omitempty"`
	City        *string           `json:"city:omitempty"`
	CompanyName *string           `json:"companyName:omitempty"`
	Country     *string           `json:"country:omitempty"`
	Email       *string           `json:"email:omitempty"`
	FirstName   *string           `json:"firstName:omitempty"`
	LastName    *string           `json:"lastName:omitempty"`
	OfficePhone *string           `json:"officePhone:omitempty"`
	PostalCode  *string           `json:"postalCode:omitempty"`
	Questions   []string          `json:"questions:omitempty"`
	Responses   []Survey_Response `json:"responses:omitempty"`
	State       *string           `json:"state:omitempty"`
	SurveyId    *string           `json:"surveyId:omitempty"`
}

type Container_RemoteManagement_Graphs_SensorSpeed struct {
	Entity

	Graph *[]byte `json:"graph:omitempty"`
	Title *string `json:"title:omitempty"`
}

type Container_RemoteManagement_Graphs_SensorTemperature struct {
	Entity

	Graph *[]byte `json:"graph:omitempty"`
	Title *string `json:"title:omitempty"`
}

type Container_RemoteManagement_PmInfo struct {
	Entity

	PmInfoId      *string `json:"pmInfoId:omitempty"`
	PmInfoReading *string `json:"pmInfoReading:omitempty"`
}

type Container_RemoteManagement_SensorReading struct {
	Entity

	LowerCritical       *string `json:"lowerCritical:omitempty"`
	LowerNonCritical    *string `json:"lowerNonCritical:omitempty"`
	LowerNonRecoverable *string `json:"lowerNonRecoverable:omitempty"`
	SensorId            *string `json:"sensorId:omitempty"`
	SensorReading       *string `json:"sensorReading:omitempty"`
	SensorUnits         *string `json:"sensorUnits:omitempty"`
	Status              *string `json:"status:omitempty"`
	UpperCritical       *string `json:"upperCritical:omitempty"`
	UpperNonCritical    *string `json:"upperNonCritical:omitempty"`
	UpperNonRecoverable *string `json:"upperNonRecoverable:omitempty"`
}

type Container_RemoteManagement_SensorReadingsWithGraphs struct {
	Entity

	RawData           []Container_RemoteManagement_SensorReading            `json:"rawData:omitempty"`
	SpeedGraphs       []Container_RemoteManagement_Graphs_SensorSpeed       `json:"speedGraphs:omitempty"`
	TemperatureGraphs []Container_RemoteManagement_Graphs_SensorTemperature `json:"temperatureGraphs:omitempty"`
}

type Container_Resource_Metadata_ServiceResource struct {
	Entity

	BackendIpAddress *string                        `json:"backendIpAddress:omitempty"`
	Type             *Network_Service_Resource_Type `json:"type:omitempty"`
}

type Container_Search_ObjectType struct {
	Entity

	Name       *string                                `json:"name:omitempty"`
	Properties []Container_Search_ObjectType_Property `json:"properties:omitempty"`
}

type Container_Search_ObjectType_Property struct {
	Entity

	Name         *string `json:"name:omitempty"`
	SortableFlag *bool   `json:"sortableFlag:omitempty"`
	Type         *string `json:"type:omitempty"`
}

type Container_Search_Result struct {
	Entity

	MatchedTerms   []string `json:"matchedTerms:omitempty"`
	RelevanceScore *float64 `json:"relevanceScore:omitempty"`
	Resource       *Entity  `json:"resource:omitempty"`
	ResourceType   *string  `json:"resourceType:omitempty"`
}

type Container_Software_Component_HostIps_Policy struct {
	Entity

	Policy      *string `json:"policy:omitempty"`
	PolicyTitle *string `json:"policyTitle:omitempty"`
}

type Container_Tax_Cache struct {
	Entity

	EffectiveTaxRate *float64                   `json:"effectiveTaxRate:omitempty"`
	Items            []Container_Tax_Cache_Item `json:"items:omitempty"`
	Status           *string                    `json:"status:omitempty"`
	TotalTaxAmount   *float64                   `json:"totalTaxAmount:omitempty"`
}

type Container_Tax_Cache_Item struct {
	Entity

	CategoryCode  *string              `json:"categoryCode:omitempty"`
	ContainerHash *string              `json:"containerHash:omitempty"`
	ItemPriceId   *int                 `json:"itemPriceId:omitempty"`
	TaxRates      *Container_Tax_Rates `json:"taxRates:omitempty"`
}

type Container_Tax_Rates struct {
	Entity

	LaborTaxRate     *float64 `json:"laborTaxRate:omitempty"`
	LocationId       *float64 `json:"locationId:omitempty"`
	OneTimeTaxRate   *float64 `json:"oneTimeTaxRate:omitempty"`
	RecurringTaxRate *float64 `json:"recurringTaxRate:omitempty"`
	SetupTaxRate     *float64 `json:"setupTaxRate:omitempty"`
}

type Container_Ticket_GraphInputs struct {
	Entity

	EndDate            *time.Time `json:"endDate:omitempty"`
	NetworkInterfaceId *int       `json:"networkInterfaceId:omitempty"`
	Pod                *int       `json:"pod:omitempty"`
	ServerName         *string    `json:"serverName:omitempty"`
	StartDate          *time.Time `json:"startDate:omitempty"`
}

type Container_Ticket_GraphOutputs struct {
	Entity

	GraphImage   *[]byte    `json:"graphImage:omitempty"`
	GraphTitle   *string    `json:"graphTitle:omitempty"`
	MaxEndDate   *time.Time `json:"maxEndDate:omitempty"`
	MinStartDate *time.Time `json:"minStartDate:omitempty"`
}

type Container_Ticket_Priority struct {
	Entity

	Name  *string `json:"name:omitempty"`
	Value *int    `json:"value:omitempty"`
}

type Container_Ticket_Survey_Preference struct {
	Entity

	Applicable   *bool      `json:"applicable:omitempty"`
	OptedOut     *bool      `json:"optedOut:omitempty"`
	OptedOutDate *time.Time `json:"optedOutDate:omitempty"`
}

type Container_User_Authentication_Token struct {
	Entity

	Hash   *string        `json:"hash:omitempty"`
	User   *User_Customer `json:"user:omitempty"`
	UserId *int           `json:"userId:omitempty"`
}

type Container_User_Customer_External_Binding struct {
	Entity

	AuthenticationToken      *string `json:"authenticationToken:omitempty"`
	OpenIdConnectAccessToken *string `json:"openIdConnectAccessToken:omitempty"`
	OpenIdConnectAccountId   *int    `json:"openIdConnectAccountId:omitempty"`
	OpenIdConnectProvider    *int    `json:"openIdConnectProvider:omitempty"`
	Password                 *string `json:"password:omitempty"`
	SecurityQuestionAnswer   *string `json:"securityQuestionAnswer:omitempty"`
	SecurityQuestionId       *int    `json:"securityQuestionId:omitempty"`
	Username                 *string `json:"username:omitempty"`
	Vendor                   *string `json:"vendor:omitempty"`
}

type Container_User_Customer_External_Binding_Phone struct {
	Container_User_Customer_External_Binding
}

type Container_User_Customer_External_Binding_Phone_Mode struct {
	Entity

	Mode    *string `json:"mode:omitempty"`
	Pin     *string `json:"pin:omitempty"`
	PinMode *string `json:"pinMode:omitempty"`
}

type Container_User_Customer_External_Binding_Totp struct {
	Container_User_Customer_External_Binding

	SecurityCode *string `json:"securityCode:omitempty"`
}

type Container_User_Customer_External_Binding_Vendor struct {
	Entity

	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Container_User_Customer_External_Binding_Verisign struct {
	Container_User_Customer_External_Binding

	SecondSecurityCode *string `json:"secondSecurityCode:omitempty"`
	SecurityCode       *string `json:"securityCode:omitempty"`
}

type Container_User_Customer_PasswordSet struct {
	Entity

	AnsweredSecurityQuestionId *int                     `json:"answeredSecurityQuestionId:omitempty"`
	AuthenticationMethods      []int                    `json:"authenticationMethods:omitempty"`
	Key                        *string                  `json:"key:omitempty"`
	Password                   *string                  `json:"password:omitempty"`
	SecurityAnswer             *string                  `json:"securityAnswer:omitempty"`
	SecurityQuestions          []User_Security_Question `json:"securityQuestions:omitempty"`
	UserId                     *int                     `json:"userId:omitempty"`
}

type Container_User_Customer_Portal_MobileToken struct {
	Container_User_Customer_Portal_Token

	HasExternalBinding *bool `json:"hasExternalBinding:omitempty"`
}

type Container_User_Customer_Portal_Token struct {
	Entity

	Hash   *string        `json:"hash:omitempty"`
	User   *User_Customer `json:"user:omitempty"`
	UserId *int           `json:"userId:omitempty"`
}

type Container_User_Data_Phone struct {
	Entity

	CountryCode *int    `json:"countryCode:omitempty"`
	Extension   *string `json:"extension:omitempty"`
	Phone       *string `json:"phone:omitempty"`
	PhoneType   *string `json:"phoneType:omitempty"`
}

type Container_User_Employee_External_Binding_Verisign struct {
	Entity
}

type Container_Utility_File_Attachment struct {
	Entity

	Data     *[]byte `json:"data:omitempty"`
	Filename *string `json:"filename:omitempty"`
}

type Container_Utility_File_Descriptor struct {
	Entity

	FileName     *string    `json:"fileName:omitempty"`
	FriendlyName *string    `json:"friendlyName:omitempty"`
	ModifyDate   *time.Time `json:"modifyDate:omitempty"`
}

type Container_Utility_File_Entity struct {
	Entity

	Content     *[]byte    `json:"content:omitempty"`
	ContentType *string    `json:"contentType:omitempty"`
	CreateDate  *time.Time `json:"createDate:omitempty"`
	DeleteDate  *time.Time `json:"deleteDate:omitempty"`
	Id          *string    `json:"id:omitempty"`
	IsShared    *int       `json:"isShared:omitempty"`
	ModifyDate  *time.Time `json:"modifyDate:omitempty"`
	Name        *string    `json:"name:omitempty"`
	Owner       *string    `json:"owner:omitempty"`
	Size        *uint      `json:"size:omitempty"`
	Type        *string    `json:"type:omitempty"`
	Version     *int       `json:"version:omitempty"`
}

type Container_Utility_Message struct {
	Entity

	CreateDate *time.Time `json:"createDate:omitempty"`
	Id         *int       `json:"id:omitempty"`
	Message    *string    `json:"message:omitempty"`
	ModifyDate *time.Time `json:"modifyDate:omitempty"`
	Summary    *string    `json:"summary:omitempty"`
}

type Container_Utility_Microsoft_Windows_UpdateServices_Status struct {
	Entity

	LastRebootDate   *time.Time `json:"lastRebootDate:omitempty"`
	LastStatusDate   *time.Time `json:"lastStatusDate:omitempty"`
	LastSyncDate     *time.Time `json:"lastSyncDate:omitempty"`
	PrivateIPAddress *string    `json:"privateIPAddress:omitempty"`
	SyncStatus       *string    `json:"syncStatus:omitempty"`
	UpdateStatus     *string    `json:"updateStatus:omitempty"`
}

type Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem struct {
	Entity

	Description     *string `json:"description:omitempty"`
	Failed          *bool   `json:"failed:omitempty"`
	KbArticleNumber *int    `json:"kbArticleNumber:omitempty"`
	Optional        *bool   `json:"optional:omitempty"`
	RequiresReboot  *bool   `json:"requiresReboot:omitempty"`
}

type Container_Utility_Network_Firewall_Rule_Attribute struct {
	Entity

	Actions             []string                                               `json:"actions:omitempty"`
	MaximumRuleCount    *int                                                   `json:"maximumRuleCount:omitempty"`
	Protocols           []string                                               `json:"protocols:omitempty"`
	SourceIpSubnetMasks []Container_Utility_Network_Subnet_Mask_Generic_Detail `json:"sourceIpSubnetMasks:omitempty"`
}

type Container_Utility_Network_Subnet_Mask_Generic_Detail struct {
	Entity

	Cidr        *string `json:"cidr:omitempty"`
	Description *string `json:"description:omitempty"`
	Mask        *string `json:"mask:omitempty"`
}

type Container_Virtual_Guest_Block_Device_Template_Configuration struct {
	Entity

	Name                         *string `json:"name:omitempty"`
	Note                         *string `json:"note:omitempty"`
	OperatingSystemReferenceCode *string `json:"operatingSystemReferenceCode:omitempty"`
	Uri                          *string `json:"uri:omitempty"`
}

type Container_Virtual_Guest_Configuration struct {
	Entity

	BlockDevices      []Container_Virtual_Guest_Configuration_Option `json:"blockDevices:omitempty"`
	Datacenters       []Container_Virtual_Guest_Configuration_Option `json:"datacenters:omitempty"`
	Memory            []Container_Virtual_Guest_Configuration_Option `json:"memory:omitempty"`
	NetworkComponents []Container_Virtual_Guest_Configuration_Option `json:"networkComponents:omitempty"`
	OperatingSystems  []Container_Virtual_Guest_Configuration_Option `json:"operatingSystems:omitempty"`
	Processors        []Container_Virtual_Guest_Configuration_Option `json:"processors:omitempty"`
}

type Container_Virtual_Guest_Configuration_Option struct {
	Entity

	ItemPrice *Product_Item_Price `json:"itemPrice:omitempty"`
	Template  *Virtual_Guest      `json:"template:omitempty"`
}

type Dns_Domain struct {
	Entity

	Account             *Account                    `json:"account:omitempty"`
	Id                  *int                        `json:"id:omitempty"`
	ManagedResourceFlag *bool                       `json:"managedResourceFlag:omitempty"`
	Name                *string                     `json:"name:omitempty"`
	ResourceRecordCount *uint                       `json:"resourceRecordCount:omitempty"`
	ResourceRecords     []Dns_Domain_ResourceRecord `json:"resourceRecords:omitempty"`
	Secondary           *Dns_Secondary              `json:"secondary:omitempty"`
	Serial              *int                        `json:"serial:omitempty"`
	UpdateDate          *time.Time                  `json:"updateDate:omitempty"`
}

type Dns_Domain_Forward struct {
	Dns_Domain
}

type Dns_Domain_Registration struct {
	Entity

	Account                        *Account                                                `json:"account:omitempty"`
	CreateDate                     *time.Time                                              `json:"createDate:omitempty"`
	DomainRegistrationStatus       *Dns_Domain_Registration_Status                         `json:"domainRegistrationStatus:omitempty"`
	DomainRegistrationStatusId     *int                                                    `json:"domainRegistrationStatusId:omitempty"`
	ExpireDate                     *time.Time                                              `json:"expireDate:omitempty"`
	Id                             *int                                                    `json:"id:omitempty"`
	LockedFlag                     *int                                                    `json:"lockedFlag:omitempty"`
	ModifyDate                     *time.Time                                              `json:"modifyDate:omitempty"`
	Name                           *string                                                 `json:"name:omitempty"`
	RegistrantVerificationStatus   *Dns_Domain_Registration_Registrant_Verification_Status `json:"registrantVerificationStatus:omitempty"`
	RegistrantVerificationStatusId *int                                                    `json:"registrantVerificationStatusId:omitempty"`
	ServiceProvider                *Service_Provider                                       `json:"serviceProvider:omitempty"`
	ServiceProviderId              *int                                                    `json:"serviceProviderId:omitempty"`
}

type Dns_Domain_Registration_Registrant_Verification_Status struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Dns_Domain_Registration_Status struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Dns_Domain_ResourceRecord struct {
	Entity

	Data              *string     `json:"data:omitempty"`
	Domain            *Dns_Domain `json:"domain:omitempty"`
	DomainId          *int        `json:"domainId:omitempty"`
	Expire            *int        `json:"expire:omitempty"`
	Host              *string     `json:"host:omitempty"`
	Id                *int        `json:"id:omitempty"`
	Minimum           *int        `json:"minimum:omitempty"`
	MxPriority        *int        `json:"mxPriority:omitempty"`
	Refresh           *int        `json:"refresh:omitempty"`
	ResponsiblePerson *string     `json:"responsiblePerson:omitempty"`
	Retry             *int        `json:"retry:omitempty"`
	Ttl               *int        `json:"ttl:omitempty"`
	Type              *string     `json:"type:omitempty"`
}

type Dns_Domain_ResourceRecord_AType struct {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_ResourceRecord_AaaaType struct {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_ResourceRecord_CnameType struct {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_ResourceRecord_MxType struct {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_ResourceRecord_NsType struct {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_ResourceRecord_PtrType struct {
	Dns_Domain_ResourceRecord

	IsGatewayAddress *bool `json:"isGatewayAddress:omitempty"`
}

type Dns_Domain_ResourceRecord_SoaType struct {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_ResourceRecord_SpfType struct {
	Dns_Domain_ResourceRecord_TxtType
}

type Dns_Domain_ResourceRecord_SrvType struct {
	Dns_Domain_ResourceRecord

	Port     *int    `json:"port:omitempty"`
	Priority *int    `json:"priority:omitempty"`
	Protocol *string `json:"protocol:omitempty"`
	Service  *string `json:"service:omitempty"`
	Weight   *int    `json:"weight:omitempty"`
}

type Dns_Domain_ResourceRecord_TxtType struct {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_Reverse struct {
	Dns_Domain

	NetworkAddress *string `json:"networkAddress:omitempty"`
}

type Dns_Domain_Reverse_Version4 struct {
	Dns_Domain_Reverse
}

type Dns_Domain_Reverse_Version6 struct {
	Dns_Domain_Reverse
}

type Dns_Message struct {
	Entity

	CreateDate     *time.Time                 `json:"createDate:omitempty"`
	Domain         *Dns_Domain                `json:"domain:omitempty"`
	Id             *int                       `json:"id:omitempty"`
	Message        *string                    `json:"message:omitempty"`
	Priority       *string                    `json:"priority:omitempty"`
	ResourceRecord *Dns_Domain_ResourceRecord `json:"resourceRecord:omitempty"`
	Secondary      *Dns_Secondary             `json:"secondary:omitempty"`
}

type Dns_Secondary struct {
	Entity

	Account           *Account      `json:"account:omitempty"`
	CreateDate        *time.Time    `json:"createDate:omitempty"`
	Domain            *Dns_Domain   `json:"domain:omitempty"`
	ErrorMessageCount *uint         `json:"errorMessageCount:omitempty"`
	ErrorMessages     []Dns_Message `json:"errorMessages:omitempty"`
	Id                *int          `json:"id:omitempty"`
	LastUpdate        *time.Time    `json:"lastUpdate:omitempty"`
	MasterIpAddress   *string       `json:"masterIpAddress:omitempty"`
	Status            *Dns_Status   `json:"status:omitempty"`
	StatusId          *int          `json:"statusId:omitempty"`
	StatusText        *string       `json:"statusText:omitempty"`
	TransferFrequency *int          `json:"transferFrequency:omitempty"`
	ZoneName          *string       `json:"zoneName:omitempty"`
}

type Dns_Status struct {
	Entity

	Id   *int    `json:"id:omitempty"`
	Name *string `json:"name:omitempty"`
}

type Entity struct {
}

type Event_Log struct {
	Entity

	AccountId       *int           `json:"accountId:omitempty"`
	EventCreateDate *time.Time     `json:"eventCreateDate:omitempty"`
	EventName       *string        `json:"eventName:omitempty"`
	IpAddress       *string        `json:"ipAddress:omitempty"`
	Label           *string        `json:"label:omitempty"`
	MetaData        *string        `json:"metaData:omitempty"`
	ObjectId        *int           `json:"objectId:omitempty"`
	ObjectName      *string        `json:"objectName:omitempty"`
	Resource        *Entity        `json:"resource:omitempty"`
	TraceId         *string        `json:"traceId:omitempty"`
	User            *User_Customer `json:"user:omitempty"`
	UserId          *int           `json:"userId:omitempty"`
	UserType        *string        `json:"userType:omitempty"`
	Username        *string        `json:"username:omitempty"`
}

type FlexibleCredit_Affiliate struct {
	Entity

	FlexibleCreditProgram *FlexibleCredit_Program `json:"flexibleCreditProgram:omitempty"`
	Id                    *int                    `json:"id:omitempty"`
	Name                  *string                 `json:"name:omitempty"`
}

type FlexibleCredit_Company_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
}

type FlexibleCredit_Enrollment struct {
	Entity

	Account                  *Account                     `json:"account:omitempty"`
	AccountId                *int                         `json:"accountId:omitempty"`
	Affiliate                *FlexibleCredit_Affiliate    `json:"affiliate:omitempty"`
	AffiliateId              *int                         `json:"affiliateId:omitempty"`
	AgreementCompleteFlag    *int                         `json:"agreementCompleteFlag:omitempty"`
	CompanyDescription       *string                      `json:"companyDescription:omitempty"`
	CompanyType              *FlexibleCredit_Company_Type `json:"companyType:omitempty"`
	CompanyTypeId            *int                         `json:"companyTypeId:omitempty"`
	EnrollmentDate           *time.Time                   `json:"enrollmentDate:omitempty"`
	FlexibleCreditProgram    *FlexibleCredit_Program      `json:"flexibleCreditProgram:omitempty"`
	GraduationDate           *time.Time                   `json:"graduationDate:omitempty"`
	IsActiveFlag             *bool                        `json:"isActiveFlag:omitempty"`
	MonthlyCreditAmount      *float64                     `json:"monthlyCreditAmount:omitempty"`
	Representative           *User_Employee               `json:"representative:omitempty"`
	RepresentativeEmployeeId *int                         `json:"representativeEmployeeId:omitempty"`
}

type FlexibleCredit_Program struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Hardware struct {
	Entity

	Account                                     *Account                                     `json:"account:omitempty"`
	AccountId                                   *int                                         `json:"accountId:omitempty"`
	ActiveComponentCount                        *uint                                        `json:"activeComponentCount:omitempty"`
	ActiveComponents                            []Hardware_Component                         `json:"activeComponents:omitempty"`
	ActiveNetworkMonitorIncident                []Network_Monitor_Version1_Incident          `json:"activeNetworkMonitorIncident:omitempty"`
	ActiveNetworkMonitorIncidentCount           *uint                                        `json:"activeNetworkMonitorIncidentCount:omitempty"`
	AllPowerComponentCount                      *uint                                        `json:"allPowerComponentCount:omitempty"`
	AllPowerComponents                          []Hardware_Power_Component                   `json:"allPowerComponents:omitempty"`
	AllowedHost                                 *Network_Storage_Allowed_Host                `json:"allowedHost:omitempty"`
	AllowedNetworkStorage                       []Network_Storage                            `json:"allowedNetworkStorage:omitempty"`
	AllowedNetworkStorageCount                  *uint                                        `json:"allowedNetworkStorageCount:omitempty"`
	AllowedNetworkStorageReplicaCount           *uint                                        `json:"allowedNetworkStorageReplicaCount:omitempty"`
	AllowedNetworkStorageReplicas               []Network_Storage                            `json:"allowedNetworkStorageReplicas:omitempty"`
	AntivirusSpywareSoftwareComponent           *Software_Component                          `json:"antivirusSpywareSoftwareComponent:omitempty"`
	AttributeCount                              *uint                                        `json:"attributeCount:omitempty"`
	Attributes                                  []Hardware_Attribute                         `json:"attributes:omitempty"`
	AverageDailyPublicBandwidthUsage            *float64                                     `json:"averageDailyPublicBandwidthUsage:omitempty"`
	BackendNetworkComponentCount                *uint                                        `json:"backendNetworkComponentCount:omitempty"`
	BackendNetworkComponents                    []Network_Component                          `json:"backendNetworkComponents:omitempty"`
	BackendRouterCount                          *uint                                        `json:"backendRouterCount:omitempty"`
	BackendRouters                              []Hardware                                   `json:"backendRouters:omitempty"`
	BandwidthAllocation                         *float64                                     `json:"bandwidthAllocation:omitempty"`
	BandwidthAllotmentDetail                    *Network_Bandwidth_Version1_Allotment_Detail `json:"bandwidthAllotmentDetail:omitempty"`
	BareMetalInstanceFlag                       *int                                         `json:"bareMetalInstanceFlag:omitempty"`
	BenchmarkCertificationCount                 *uint                                        `json:"benchmarkCertificationCount:omitempty"`
	BenchmarkCertifications                     []Hardware_Benchmark_Certification           `json:"benchmarkCertifications:omitempty"`
	BillingItem                                 *Billing_Item_Hardware                       `json:"billingItem:omitempty"`
	BillingItemFlag                             *bool                                        `json:"billingItemFlag:omitempty"`
	BlockCancelBecauseDisconnectedFlag          *bool                                        `json:"blockCancelBecauseDisconnectedFlag:omitempty"`
	BusinessContinuanceInsuranceFlag            *bool                                        `json:"businessContinuanceInsuranceFlag:omitempty"`
	ComponentCount                              *uint                                        `json:"componentCount:omitempty"`
	Components                                  []Hardware_Component                         `json:"components:omitempty"`
	ContinuousDataProtectionSoftwareComponent   *Software_Component                          `json:"continuousDataProtectionSoftwareComponent:omitempty"`
	CurrentBillableBandwidthUsage               *float64                                     `json:"currentBillableBandwidthUsage:omitempty"`
	Datacenter                                  *Location                                    `json:"datacenter:omitempty"`
	DatacenterName                              *string                                      `json:"datacenterName:omitempty"`
	Domain                                      *string                                      `json:"domain:omitempty"`
	DownlinkHardware                            []Hardware                                   `json:"downlinkHardware:omitempty"`
	DownlinkHardwareCount                       *uint                                        `json:"downlinkHardwareCount:omitempty"`
	DownlinkNetworkHardware                     []Hardware                                   `json:"downlinkNetworkHardware:omitempty"`
	DownlinkNetworkHardwareCount                *uint                                        `json:"downlinkNetworkHardwareCount:omitempty"`
	DownlinkServerCount                         *uint                                        `json:"downlinkServerCount:omitempty"`
	DownlinkServers                             []Hardware                                   `json:"downlinkServers:omitempty"`
	DownlinkVirtualGuestCount                   *uint                                        `json:"downlinkVirtualGuestCount:omitempty"`
	DownlinkVirtualGuests                       []Virtual_Guest                              `json:"downlinkVirtualGuests:omitempty"`
	DownstreamHardwareBindingCount              *uint                                        `json:"downstreamHardwareBindingCount:omitempty"`
	DownstreamHardwareBindings                  []Network_Component_Uplink_Hardware          `json:"downstreamHardwareBindings:omitempty"`
	DownstreamNetworkHardware                   []Hardware                                   `json:"downstreamNetworkHardware:omitempty"`
	DownstreamNetworkHardwareCount              *uint                                        `json:"downstreamNetworkHardwareCount:omitempty"`
	DownstreamNetworkHardwareWithIncidentCount  *uint                                        `json:"downstreamNetworkHardwareWithIncidentCount:omitempty"`
	DownstreamNetworkHardwareWithIncidents      []Hardware                                   `json:"downstreamNetworkHardwareWithIncidents:omitempty"`
	DownstreamServerCount                       *uint                                        `json:"downstreamServerCount:omitempty"`
	DownstreamServers                           []Hardware                                   `json:"downstreamServers:omitempty"`
	DownstreamVirtualGuestCount                 *uint                                        `json:"downstreamVirtualGuestCount:omitempty"`
	DownstreamVirtualGuests                     []Virtual_Guest                              `json:"downstreamVirtualGuests:omitempty"`
	DriveControllerCount                        *uint                                        `json:"driveControllerCount:omitempty"`
	DriveControllers                            []Hardware_Component                         `json:"driveControllers:omitempty"`
	EvaultNetworkStorage                        []Network_Storage                            `json:"evaultNetworkStorage:omitempty"`
	EvaultNetworkStorageCount                   *uint                                        `json:"evaultNetworkStorageCount:omitempty"`
	FirewallServiceComponent                    *Network_Component_Firewall                  `json:"firewallServiceComponent:omitempty"`
	FixedConfigurationPreset                    *Product_Package_Preset                      `json:"fixedConfigurationPreset:omitempty"`
	FrontendNetworkComponentCount               *uint                                        `json:"frontendNetworkComponentCount:omitempty"`
	FrontendNetworkComponents                   []Network_Component                          `json:"frontendNetworkComponents:omitempty"`
	FrontendRouterCount                         *uint                                        `json:"frontendRouterCount:omitempty"`
	FrontendRouters                             []Hardware                                   `json:"frontendRouters:omitempty"`
	FullyQualifiedDomainName                    *string                                      `json:"fullyQualifiedDomainName:omitempty"`
	GlobalIdentifier                            *string                                      `json:"globalIdentifier:omitempty"`
	HardDriveCount                              *uint                                        `json:"hardDriveCount:omitempty"`
	HardDrives                                  []Hardware_Component                         `json:"hardDrives:omitempty"`
	HardwareChassis                             *Hardware_Chassis                            `json:"hardwareChassis:omitempty"`
	HardwareFunction                            *Hardware_Function                           `json:"hardwareFunction:omitempty"`
	HardwareFunctionDescription                 *string                                      `json:"hardwareFunctionDescription:omitempty"`
	HardwareStatus                              *Hardware_Status                             `json:"hardwareStatus:omitempty"`
	HardwareStatusId                            *int                                         `json:"hardwareStatusId:omitempty"`
	HasTrustedPlatformModuleBillingItemFlag     *bool                                        `json:"hasTrustedPlatformModuleBillingItemFlag:omitempty"`
	HostIpsSoftwareComponent                    *Software_Component                          `json:"hostIpsSoftwareComponent:omitempty"`
	Hostname                                    *string                                      `json:"hostname:omitempty"`
	HourlyBillingFlag                           *bool                                        `json:"hourlyBillingFlag:omitempty"`
	Id                                          *int                                         `json:"id:omitempty"`
	InboundBandwidthUsage                       *float64                                     `json:"inboundBandwidthUsage:omitempty"`
	InboundPublicBandwidthUsage                 *float64                                     `json:"inboundPublicBandwidthUsage:omitempty"`
	LastTransaction                             *Provisioning_Version1_Transaction           `json:"lastTransaction:omitempty"`
	LatestNetworkMonitorIncident                *Network_Monitor_Version1_Incident           `json:"latestNetworkMonitorIncident:omitempty"`
	Location                                    *Location                                    `json:"location:omitempty"`
	LocationPathString                          *string                                      `json:"locationPathString:omitempty"`
	LockboxNetworkStorage                       *Network_Storage                             `json:"lockboxNetworkStorage:omitempty"`
	ManagedResourceFlag                         *bool                                        `json:"managedResourceFlag:omitempty"`
	ManufacturerSerialNumber                    *string                                      `json:"manufacturerSerialNumber:omitempty"`
	Memory                                      []Hardware_Component                         `json:"memory:omitempty"`
	MemoryCapacity                              *uint                                        `json:"memoryCapacity:omitempty"`
	MemoryCount                                 *uint                                        `json:"memoryCount:omitempty"`
	MetricTrackingObject                        *Metric_Tracking_Object_HardwareServer       `json:"metricTrackingObject:omitempty"`
	MonitoringAgentCount                        *uint                                        `json:"monitoringAgentCount:omitempty"`
	MonitoringAgents                            []Monitoring_Agent                           `json:"monitoringAgents:omitempty"`
	MonitoringRobot                             *Monitoring_Robot                            `json:"monitoringRobot:omitempty"`
	MonitoringServiceComponent                  *Network_Monitor_Version1_Query_Host_Stratum `json:"monitoringServiceComponent:omitempty"`
	MonitoringServiceEligibilityFlag            *bool                                        `json:"monitoringServiceEligibilityFlag:omitempty"`
	MonitoringServiceFlag                       *bool                                        `json:"monitoringServiceFlag:omitempty"`
	Motherboard                                 *Hardware_Component                          `json:"motherboard:omitempty"`
	NetworkCardCount                            *uint                                        `json:"networkCardCount:omitempty"`
	NetworkCards                                []Hardware_Component                         `json:"networkCards:omitempty"`
	NetworkComponentCount                       *uint                                        `json:"networkComponentCount:omitempty"`
	NetworkComponents                           []Network_Component                          `json:"networkComponents:omitempty"`
	NetworkGatewayMember                        *Network_Gateway_Member                      `json:"networkGatewayMember:omitempty"`
	NetworkGatewayMemberFlag                    *bool                                        `json:"networkGatewayMemberFlag:omitempty"`
	NetworkManagementIpAddress                  *string                                      `json:"networkManagementIpAddress:omitempty"`
	NetworkMonitorAttachedDownHardware          []Hardware                                   `json:"networkMonitorAttachedDownHardware:omitempty"`
	NetworkMonitorAttachedDownHardwareCount     *uint                                        `json:"networkMonitorAttachedDownHardwareCount:omitempty"`
	NetworkMonitorAttachedDownVirtualGuestCount *uint                                        `json:"networkMonitorAttachedDownVirtualGuestCount:omitempty"`
	NetworkMonitorAttachedDownVirtualGuests     []Virtual_Guest                              `json:"networkMonitorAttachedDownVirtualGuests:omitempty"`
	NetworkMonitorCount                         *uint                                        `json:"networkMonitorCount:omitempty"`
	NetworkMonitorIncidentCount                 *uint                                        `json:"networkMonitorIncidentCount:omitempty"`
	NetworkMonitorIncidents                     []Network_Monitor_Version1_Incident          `json:"networkMonitorIncidents:omitempty"`
	NetworkMonitors                             []Network_Monitor_Version1_Query_Host        `json:"networkMonitors:omitempty"`
	NetworkStatus                               *string                                      `json:"networkStatus:omitempty"`
	NetworkStatusAttribute                      *Hardware_Attribute                          `json:"networkStatusAttribute:omitempty"`
	NetworkStorage                              []Network_Storage                            `json:"networkStorage:omitempty"`
	NetworkStorageCount                         *uint                                        `json:"networkStorageCount:omitempty"`
	NetworkVlanCount                            *uint                                        `json:"networkVlanCount:omitempty"`
	NetworkVlans                                []Network_Vlan                               `json:"networkVlans:omitempty"`
	NextBillingCycleBandwidthAllocation         *float64                                     `json:"nextBillingCycleBandwidthAllocation:omitempty"`
	Notes                                       *string                                      `json:"notes:omitempty"`
	NotesHistory                                []Hardware_Note                              `json:"notesHistory:omitempty"`
	NotesHistoryCount                           *uint                                        `json:"notesHistoryCount:omitempty"`
	OperatingSystem                             *Software_Component_OperatingSystem          `json:"operatingSystem:omitempty"`
	OperatingSystemReferenceCode                *string                                      `json:"operatingSystemReferenceCode:omitempty"`
	OutboundBandwidthUsage                      *float64                                     `json:"outboundBandwidthUsage:omitempty"`
	OutboundPublicBandwidthUsage                *float64                                     `json:"outboundPublicBandwidthUsage:omitempty"`
	PointOfPresenceLocation                     *Location                                    `json:"pointOfPresenceLocation:omitempty"`
	PostInstallScriptUri                        *string                                      `json:"postInstallScriptUri:omitempty"`
	PowerComponentCount                         *uint                                        `json:"powerComponentCount:omitempty"`
	PowerComponents                             []Hardware_Power_Component                   `json:"powerComponents:omitempty"`
	PowerSupply                                 []Hardware_Component                         `json:"powerSupply:omitempty"`
	PowerSupplyCount                            *uint                                        `json:"powerSupplyCount:omitempty"`
	PrimaryBackendIpAddress                     *string                                      `json:"primaryBackendIpAddress:omitempty"`
	PrimaryBackendNetworkComponent              *Network_Component                           `json:"primaryBackendNetworkComponent:omitempty"`
	PrimaryIpAddress                            *string                                      `json:"primaryIpAddress:omitempty"`
	PrimaryNetworkComponent                     *Network_Component                           `json:"primaryNetworkComponent:omitempty"`
	PrivateNetworkOnlyFlag                      *bool                                        `json:"privateNetworkOnlyFlag:omitempty"`
	ProcessorCoreAmount                         *uint                                        `json:"processorCoreAmount:omitempty"`
	ProcessorCount                              *uint                                        `json:"processorCount:omitempty"`
	ProcessorPhysicalCoreAmount                 *uint                                        `json:"processorPhysicalCoreAmount:omitempty"`
	Processors                                  []Hardware_Component                         `json:"processors:omitempty"`
	ProvisionDate                               *time.Time                                   `json:"provisionDate:omitempty"`
	Rack                                        *Location                                    `json:"rack:omitempty"`
	RaidControllerCount                         *uint                                        `json:"raidControllerCount:omitempty"`
	RaidControllers                             []Hardware_Component                         `json:"raidControllers:omitempty"`
	RecentEventCount                            *uint                                        `json:"recentEventCount:omitempty"`
	RecentEvents                                []Notification_Occurrence_Event              `json:"recentEvents:omitempty"`
	RemoteManagementAccountCount                *uint                                        `json:"remoteManagementAccountCount:omitempty"`
	RemoteManagementAccounts                    []Hardware_Component_RemoteManagement_User   `json:"remoteManagementAccounts:omitempty"`
	RemoteManagementComponent                   *Network_Component                           `json:"remoteManagementComponent:omitempty"`
	ResourceGroupCount                          *uint                                        `json:"resourceGroupCount:omitempty"`
	ResourceGroupMemberReferenceCount           *uint                                        `json:"resourceGroupMemberReferenceCount:omitempty"`
	ResourceGroupMemberReferences               []Resource_Group_Member                      `json:"resourceGroupMemberReferences:omitempty"`
	ResourceGroupRoleCount                      *uint                                        `json:"resourceGroupRoleCount:omitempty"`
	ResourceGroupRoles                          []Resource_Group_Role                        `json:"resourceGroupRoles:omitempty"`
	ResourceGroups                              []Resource_Group                             `json:"resourceGroups:omitempty"`
	RouterCount                                 *uint                                        `json:"routerCount:omitempty"`
	Routers                                     []Hardware                                   `json:"routers:omitempty"`
	ScaleAssetCount                             *uint                                        `json:"scaleAssetCount:omitempty"`
	ScaleAssets                                 []Scale_Asset                                `json:"scaleAssets:omitempty"`
	SecurityScanRequestCount                    *uint                                        `json:"securityScanRequestCount:omitempty"`
	SecurityScanRequests                        []Network_Security_Scanner_Request           `json:"securityScanRequests:omitempty"`
	SerialNumber                                *string                                      `json:"serialNumber:omitempty"`
	ServerRoom                                  *Location                                    `json:"serverRoom:omitempty"`
	ServiceProvider                             *Service_Provider                            `json:"serviceProvider:omitempty"`
	ServiceProviderId                           *int                                         `json:"serviceProviderId:omitempty"`
	ServiceProviderResourceId                   *int                                         `json:"serviceProviderResourceId:omitempty"`
	SoftwareComponentCount                      *uint                                        `json:"softwareComponentCount:omitempty"`
	SoftwareComponents                          []Software_Component                         `json:"softwareComponents:omitempty"`
	SparePoolBillingItem                        *Billing_Item_Hardware                       `json:"sparePoolBillingItem:omitempty"`
	SshKeyCount                                 *uint                                        `json:"sshKeyCount:omitempty"`
	SshKeys                                     []Security_Ssh_Key                           `json:"sshKeys:omitempty"`
	StorageNetworkComponentCount                *uint                                        `json:"storageNetworkComponentCount:omitempty"`
	StorageNetworkComponents                    []Network_Component                          `json:"storageNetworkComponents:omitempty"`
	TagReferenceCount                           *uint                                        `json:"tagReferenceCount:omitempty"`
	TagReferences                               []Tag_Reference                              `json:"tagReferences:omitempty"`
	TopLevelLocation                            *Location                                    `json:"topLevelLocation:omitempty"`
	UpgradeRequest                              *Product_Upgrade_Request                     `json:"upgradeRequest:omitempty"`
	UplinkHardware                              *Hardware                                    `json:"uplinkHardware:omitempty"`
	UplinkNetworkComponentCount                 *uint                                        `json:"uplinkNetworkComponentCount:omitempty"`
	UplinkNetworkComponents                     []Network_Component                          `json:"uplinkNetworkComponents:omitempty"`
	UserData                                    []Hardware_Attribute                         `json:"userData:omitempty"`
	UserDataCount                               *uint                                        `json:"userDataCount:omitempty"`
	VirtualChassis                              *Hardware_Group                              `json:"virtualChassis:omitempty"`
	VirtualChassisSiblingCount                  *uint                                        `json:"virtualChassisSiblingCount:omitempty"`
	VirtualChassisSiblings                      []Hardware                                   `json:"virtualChassisSiblings:omitempty"`
	VirtualHost                                 *Virtual_Host                                `json:"virtualHost:omitempty"`
	VirtualLicenseCount                         *uint                                        `json:"virtualLicenseCount:omitempty"`
	VirtualLicenses                             []Software_VirtualLicense                    `json:"virtualLicenses:omitempty"`
	VirtualRack                                 *Network_Bandwidth_Version1_Allotment        `json:"virtualRack:omitempty"`
	VirtualRackId                               *int                                         `json:"virtualRackId:omitempty"`
	VirtualRackName                             *string                                      `json:"virtualRackName:omitempty"`
	VirtualizationPlatform                      *Software_Component                          `json:"virtualizationPlatform:omitempty"`
}

type Hardware_Attribute struct {
	Entity

	HardwareAttributeType   *Hardware_Attribute_Type `json:"hardwareAttributeType:omitempty"`
	HardwareAttributeTypeId *int                     `json:"hardwareAttributeTypeId:omitempty"`
	Id                      *int                     `json:"id:omitempty"`
	Value                   *string                  `json:"value:omitempty"`
}

type Hardware_Attribute_Type struct {
	Entity

	Keyname *string `json:"keyname:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Hardware_Benchmark_Certification struct {
	Entity

	Account    *Account   `json:"account:omitempty"`
	AccountId  *int       `json:"accountId:omitempty"`
	CreateDate *time.Time `json:"createDate:omitempty"`
	Hardware   *Hardware  `json:"hardware:omitempty"`
	HardwareId *int       `json:"hardwareId:omitempty"`
}

type Hardware_Chassis struct {
	Entity

	BackplaneCapacity       *string            `json:"backplaneCapacity:omitempty"`
	BayCapacity             *string            `json:"bayCapacity:omitempty"`
	DriveCapacity           *string            `json:"driveCapacity:omitempty"`
	DriveControllerCapacity *string            `json:"driveControllerCapacity:omitempty"`
	FormFactorId            *int               `json:"formFactorId:omitempty"`
	GpuCapacity             *string            `json:"gpuCapacity:omitempty"`
	HardwareFunction        *Hardware_Function `json:"hardwareFunction:omitempty"`
	Id                      *int               `json:"id:omitempty"`
	Manufacturer            *string            `json:"manufacturer:omitempty"`
	Name                    *string            `json:"name:omitempty"`
	PowerCapacity           *string            `json:"powerCapacity:omitempty"`
	UnitSize                *int               `json:"unitSize:omitempty"`
	Version                 *string            `json:"version:omitempty"`
}

type Hardware_Component struct {
	Entity

	Capacity                       *float64                  `json:"capacity:omitempty"`
	Children                       []Hardware_Component      `json:"children:omitempty"`
	ChildrenCount                  *uint                     `json:"childrenCount:omitempty"`
	DownlinkHardwareComponentCount *uint                     `json:"downlinkHardwareComponentCount:omitempty"`
	DownlinkHardwareComponents     []Hardware_Component      `json:"downlinkHardwareComponents:omitempty"`
	Hardware                       *Hardware                 `json:"hardware:omitempty"`
	HardwareComponentModel         *Hardware_Component_Model `json:"hardwareComponentModel:omitempty"`
	HardwareComponentModelId       *int                      `json:"hardwareComponentModelId:omitempty"`
	HardwareComponentType          *Hardware_Component_Type  `json:"hardwareComponentType:omitempty"`
	HardwareId                     *int                      `json:"hardwareId:omitempty"`
	Id                             *int                      `json:"id:omitempty"`
	ModifyDate                     *time.Time                `json:"modifyDate:omitempty"`
	Name                           *string                   `json:"name:omitempty"`
	NetworkComponentCount          *uint                     `json:"networkComponentCount:omitempty"`
	NetworkComponents              []Network_Component       `json:"networkComponents:omitempty"`
	Owner                          *Account                  `json:"owner:omitempty"`
	Parent                         *Hardware_Component       `json:"parent:omitempty"`
	RaidMode                       *string                   `json:"raidMode:omitempty"`
	SerialNumber                   *string                   `json:"serialNumber:omitempty"`
	ServiceProvider                *Service_Provider         `json:"serviceProvider:omitempty"`
	ServiceProviderId              *int                      `json:"serviceProviderId:omitempty"`
	UplinkHardwareComponentCount   *uint                     `json:"uplinkHardwareComponentCount:omitempty"`
	UplinkHardwareComponents       []Hardware_Component      `json:"uplinkHardwareComponents:omitempty"`
}

type Hardware_Component_Attribute struct {
	Entity

	HardwareComponent                *Hardware_Component                `json:"hardwareComponent:omitempty"`
	HardwareComponentAttributeType   *Hardware_Component_Attribute_Type `json:"hardwareComponentAttributeType:omitempty"`
	HardwareComponentAttributeTypeId *int                               `json:"hardwareComponentAttributeTypeId:omitempty"`
	HardwareComponentId              *int                               `json:"hardwareComponentId:omitempty"`
	Value                            *string                            `json:"value:omitempty"`
}

type Hardware_Component_Attribute_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Hardware_Component_DriveController struct {
	Hardware_Component
}

type Hardware_Component_HardDrive struct {
	Hardware_Component

	PartitionCount *uint                          `json:"partitionCount:omitempty"`
	Partitions     []Hardware_Component_Partition `json:"partitions:omitempty"`
}

type Hardware_Component_Model struct {
	Entity

	ArchitectureType                    *Hardware_Component_Model_Architecture_Type `json:"architectureType:omitempty"`
	ArchitectureTypeId                  *string                                     `json:"architectureTypeId:omitempty"`
	AttributeCount                      *uint                                       `json:"attributeCount:omitempty"`
	Attributes                          []Hardware_Component_Model_Attribute        `json:"attributes:omitempty"`
	Capacity                            *float64                                    `json:"capacity:omitempty"`
	CompatibleArrayTypeCount            *uint                                       `json:"compatibleArrayTypeCount:omitempty"`
	CompatibleArrayTypes                []Configuration_Storage_Group_Array_Type    `json:"compatibleArrayTypes:omitempty"`
	CompatibleChildComponentModelCount  *uint                                       `json:"compatibleChildComponentModelCount:omitempty"`
	CompatibleChildComponentModels      []Hardware_Component_Model                  `json:"compatibleChildComponentModels:omitempty"`
	CompatibleParentComponentModelCount *uint                                       `json:"compatibleParentComponentModelCount:omitempty"`
	CompatibleParentComponentModels     []Hardware_Component_Model                  `json:"compatibleParentComponentModels:omitempty"`
	Description                         *string                                     `json:"description:omitempty"`
	HardwareComponents                  []Hardware_Component                        `json:"hardwareComponents:omitempty"`
	HardwareGenericComponentModel       *Hardware_Component_Model_Generic           `json:"hardwareGenericComponentModel:omitempty"`
	HardwareGenericComponentModelId     *int                                        `json:"hardwareGenericComponentModelId:omitempty"`
	Id                                  *int                                        `json:"id:omitempty"`
	InfinibandCompatibleAttribute       *Hardware_Component_Model_Attribute         `json:"infinibandCompatibleAttribute:omitempty"`
	IsInfinibandCompatible              *bool                                       `json:"isInfinibandCompatible:omitempty"`
	LongDescription                     *string                                     `json:"longDescription:omitempty"`
	Manufacturer                        *string                                     `json:"manufacturer:omitempty"`
	Name                                *string                                     `json:"name:omitempty"`
	RebootTime                          *Hardware_Component_Motherboard_Reboot_Time `json:"rebootTime:omitempty"`
	Type                                *string                                     `json:"type:omitempty"`
	ValidAttributeTypeCount             *uint                                       `json:"validAttributeTypeCount:omitempty"`
	ValidAttributeTypes                 []Hardware_Component_Model_Attribute_Type   `json:"validAttributeTypes:omitempty"`
	Version                             *string                                     `json:"version:omitempty"`
}

type Hardware_Component_Model_Architecture_Type struct {
	Entity

	Children      []Hardware_Component_Model_Architecture_Type `json:"children:omitempty"`
	ChildrenCount *uint                                        `json:"childrenCount:omitempty"`
	Id            *int                                         `json:"id:omitempty"`
	KeyName       *string                                      `json:"keyName:omitempty"`
	Name          *string                                      `json:"name:omitempty"`
	Parent        *Hardware_Component_Model_Architecture_Type  `json:"parent:omitempty"`
	ParentId      *string                                      `json:"parentId:omitempty"`
}

type Hardware_Component_Model_Attribute struct {
	Entity

	AttributeTypeId                *int                                     `json:"attributeTypeId:omitempty"`
	HardwareComponent              *Hardware_Component_Model                `json:"hardwareComponent:omitempty"`
	HardwareComponentAttributeType *Hardware_Component_Model_Attribute_Type `json:"hardwareComponentAttributeType:omitempty"`
	HardwareComponentModelId       *int                                     `json:"hardwareComponentModelId:omitempty"`
	Value                          *string                                  `json:"value:omitempty"`
}

type Hardware_Component_Model_Attribute_Type struct {
	Entity

	Description             *string                   `json:"description:omitempty"`
	Id                      *int                      `json:"id:omitempty"`
	KeyName                 *string                   `json:"keyName:omitempty"`
	Name                    *string                   `json:"name:omitempty"`
	ValidComponentTypeCount *uint                     `json:"validComponentTypeCount:omitempty"`
	ValidComponentTypes     []Hardware_Component_Type `json:"validComponentTypes:omitempty"`
}

type Hardware_Component_Model_Generic struct {
	Entity

	Capacity                    *float64                                           `json:"capacity:omitempty"`
	Description                 *string                                            `json:"description:omitempty"`
	HardwareComponentModelCount *uint                                              `json:"hardwareComponentModelCount:omitempty"`
	HardwareComponentModels     []Hardware_Component_Model                         `json:"hardwareComponentModels:omitempty"`
	HardwareComponentType       *Hardware_Component_Type                           `json:"hardwareComponentType:omitempty"`
	HardwareComponentTypeId     *int                                               `json:"hardwareComponentTypeId:omitempty"`
	Id                          *int                                               `json:"id:omitempty"`
	MarketingFeatures           *Hardware_Component_Model_Generic_MarketingFeature `json:"marketingFeatures:omitempty"`
	Units                       *string                                            `json:"units:omitempty"`
	UpgradePriority             *int                                               `json:"upgradePriority:omitempty"`
}

type Hardware_Component_Model_Generic_Attribute struct {
	Entity

	HardwareGenericComponentModel *Hardware_Component_Model_Generic `json:"hardwareGenericComponentModel:omitempty"`
	Value                         *string                           `json:"value:omitempty"`
}

type Hardware_Component_Model_Generic_MarketingFeature struct {
	Entity

	Features                      *string                           `json:"features:omitempty"`
	HardwareGenericComponentModel *Hardware_Component_Model_Generic `json:"hardwareGenericComponentModel:omitempty"`
	Price                         *string                           `json:"price:omitempty"`
}

type Hardware_Component_Motherboard struct {
	Hardware_Component
}

type Hardware_Component_Motherboard_Reboot_Time struct {
	Entity

	HardwareComponentModel *Hardware_Component_Model `json:"hardwareComponentModel:omitempty"`
	WithRaid               *int                      `json:"withRaid:omitempty"`
	WithoutRaid            *int                      `json:"withoutRaid:omitempty"`
}

type Hardware_Component_NetworkCard struct {
	Hardware_Component
}

type Hardware_Component_Partition struct {
	Entity

	DiskNumber          *int                `json:"diskNumber:omitempty"`
	Grow                *int                `json:"grow:omitempty"`
	HardwareComponent   *Hardware_Component `json:"hardwareComponent:omitempty"`
	HardwareComponentId *int                `json:"hardwareComponentId:omitempty"`
	MinimumSize         *float64            `json:"minimumSize:omitempty"`
	Name                *string             `json:"name:omitempty"`
}

type Hardware_Component_Partition_OperatingSystem struct {
	Entity

	Description            *string                                 `json:"description:omitempty"`
	Id                     *int                                    `json:"id:omitempty"`
	Notes                  *string                                 `json:"notes:omitempty"`
	PartitionTemplateCount *uint                                   `json:"partitionTemplateCount:omitempty"`
	PartitionTemplates     []Hardware_Component_Partition_Template `json:"partitionTemplates:omitempty"`
}

type Hardware_Component_Partition_Template struct {
	Entity

	Account                         *Account                                          `json:"account:omitempty"`
	AccountId                       *int                                              `json:"accountId:omitempty"`
	Data                            []Hardware_Component_Partition_Template_Partition `json:"data:omitempty"`
	DataCount                       *uint                                             `json:"dataCount:omitempty"`
	Description                     *string                                           `json:"description:omitempty"`
	ExpireDate                      *string                                           `json:"expireDate:omitempty"`
	Id                              *int                                              `json:"id:omitempty"`
	PartitionOperatingSystem        *Hardware_Component_Partition_OperatingSystem     `json:"partitionOperatingSystem:omitempty"`
	PartitionOperatingSystemId      *int                                              `json:"partitionOperatingSystemId:omitempty"`
	PartitionTemplatePartition      []Hardware_Component_Partition_Template_Partition `json:"partitionTemplatePartition:omitempty"`
	PartitionTemplatePartitionCount *uint                                             `json:"partitionTemplatePartitionCount:omitempty"`
	StatusCode                      *string                                           `json:"statusCode:omitempty"`
	TemplateType                    *string                                           `json:"templateType:omitempty"`
}

type Hardware_Component_Partition_Template_Partition struct {
	Entity

	FilesystemType      *Configuration_Storage_Filesystem_Type `json:"filesystemType:omitempty"`
	Id                  *int                                   `json:"id:omitempty"`
	IsGrow              *bool                                  `json:"isGrow:omitempty"`
	PartitionName       *string                                `json:"partitionName:omitempty"`
	PartitionSize       *float64                               `json:"partitionSize:omitempty"`
	PartitionTemplate   *Hardware_Component_Partition_Template `json:"partitionTemplate:omitempty"`
	PartitionTemplateId *int                                   `json:"partitionTemplateId:omitempty"`
}

type Hardware_Component_Processor struct {
	Hardware_Component
}

type Hardware_Component_Ram struct {
	Hardware_Component
}

type Hardware_Component_RemoteManagement struct {
	Hardware_Component

	NetworkComponent *Network_Component `json:"networkComponent:omitempty"`
}

type Hardware_Component_RemoteManagement_Command struct {
	Entity

	KeyName      *string                                               `json:"keyName:omitempty"`
	RequestCount *uint                                                 `json:"requestCount:omitempty"`
	Requests     []Hardware_Component_RemoteManagement_Command_Request `json:"requests:omitempty"`
}

type Hardware_Component_RemoteManagement_Command_Request struct {
	Entity

	CreateDate              *time.Time                                   `json:"createDate:omitempty"`
	Hardware                *Hardware                                    `json:"hardware:omitempty"`
	HardwareId              *int                                         `json:"hardwareId:omitempty"`
	ModifyDate              *time.Time                                   `json:"modifyDate:omitempty"`
	NetworkComponent        *Network_Component                           `json:"networkComponent:omitempty"`
	Processed               *bool                                        `json:"processed:omitempty"`
	RemoteManagementCommand *Hardware_Component_RemoteManagement_Command `json:"remoteManagementCommand:omitempty"`
	User                    *User_Customer                               `json:"user:omitempty"`
}

type Hardware_Component_RemoteManagement_User struct {
	Entity

	Hardware         *Hardware          `json:"hardware:omitempty"`
	NetworkComponent *Network_Component `json:"networkComponent:omitempty"`
	Password         *string            `json:"password:omitempty"`
	Username         *string            `json:"username:omitempty"`
}

type Hardware_Component_SecurityDevice struct {
	Hardware_Component
}

type Hardware_Component_SecurityDevice_Infineon struct {
	Hardware_Component_SecurityDevice
}

type Hardware_Component_Type struct {
	Entity

	HardwareGenericComponentModelCount *uint                              `json:"hardwareGenericComponentModelCount:omitempty"`
	HardwareGenericComponentModels     []Hardware_Component_Model_Generic `json:"hardwareGenericComponentModels:omitempty"`
	Id                                 *int                               `json:"id:omitempty"`
	KeyName                            *string                            `json:"keyName:omitempty"`
	Type                               *string                            `json:"type:omitempty"`
	TypeParent                         *Hardware_Component_Type           `json:"typeParent:omitempty"`
	TypeParentId                       *int                               `json:"typeParentId:omitempty"`
}

type Hardware_Firewall struct {
	Hardware_Switch

	UserCount *uint           `json:"userCount:omitempty"`
	Users     []User_Customer `json:"users:omitempty"`
}

type Hardware_Function struct {
	Entity

	Code        *string `json:"code:omitempty"`
	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
}

type Hardware_Group struct {
	Entity

	Domain                                      *string           `json:"domain:omitempty"`
	DownlinkServerCount                         *uint             `json:"downlinkServerCount:omitempty"`
	DownlinkServers                             []Hardware        `json:"downlinkServers:omitempty"`
	DownlinkVirtualGuestCount                   *uint             `json:"downlinkVirtualGuestCount:omitempty"`
	DownlinkVirtualGuests                       []Virtual_Guest   `json:"downlinkVirtualGuests:omitempty"`
	DownstreamNetworkHardware                   []Hardware        `json:"downstreamNetworkHardware:omitempty"`
	DownstreamNetworkHardwareCount              *uint             `json:"downstreamNetworkHardwareCount:omitempty"`
	DownstreamNetworkHardwareWithIncidentCount  *uint             `json:"downstreamNetworkHardwareWithIncidentCount:omitempty"`
	DownstreamNetworkHardwareWithIncidents      []Hardware        `json:"downstreamNetworkHardwareWithIncidents:omitempty"`
	HardwareChassis                             *Hardware_Chassis `json:"hardwareChassis:omitempty"`
	Hostname                                    *string           `json:"hostname:omitempty"`
	Id                                          *int              `json:"id:omitempty"`
	NetworkMonitorAttachedDownHardware          []Hardware        `json:"networkMonitorAttachedDownHardware:omitempty"`
	NetworkMonitorAttachedDownHardwareCount     *uint             `json:"networkMonitorAttachedDownHardwareCount:omitempty"`
	NetworkMonitorAttachedDownVirtualGuestCount *uint             `json:"networkMonitorAttachedDownVirtualGuestCount:omitempty"`
	NetworkMonitorAttachedDownVirtualGuests     []Virtual_Guest   `json:"networkMonitorAttachedDownVirtualGuests:omitempty"`
	NetworkStatus                               *string           `json:"networkStatus:omitempty"`
}

type Hardware_LoadBalancer struct {
	Hardware

	ModelFamily *string         `json:"modelFamily:omitempty"`
	UserCount   *uint           `json:"userCount:omitempty"`
	Users       []User_Customer `json:"users:omitempty"`
}

type Hardware_Note struct {
	Entity

	CreateDate   *time.Time          `json:"createDate:omitempty"`
	Employee     *User_Employee      `json:"employee:omitempty"`
	Hardware     *Hardware           `json:"hardware:omitempty"`
	HardwareId   *int                `json:"hardwareId:omitempty"`
	Id           *int                `json:"id:omitempty"`
	ModifyDate   *time.Time          `json:"modifyDate:omitempty"`
	Note         *string             `json:"note:omitempty"`
	Type         *Hardware_Note_Type `json:"type:omitempty"`
	TypeId       *int                `json:"typeId:omitempty"`
	User         *User_Customer      `json:"user:omitempty"`
	UserRecordId *int                `json:"userRecordId:omitempty"`
}

type Hardware_Note_Type struct {
	Entity

	KeyName *string `json:"keyName:omitempty"`
}

type Hardware_Power_Component struct {
	Entity

	Hardware   *Hardware `json:"hardware:omitempty"`
	HardwareId *int      `json:"hardwareId:omitempty"`
	Id         *int      `json:"id:omitempty"`
}

type Hardware_Router struct {
	Hardware_Switch

	BoundSubnetCount               *uint            `json:"boundSubnetCount:omitempty"`
	BoundSubnets                   []Network_Subnet `json:"boundSubnets:omitempty"`
	LocalDiskStorageCapabilityFlag *bool            `json:"localDiskStorageCapabilityFlag:omitempty"`
	SanStorageCapabilityFlag       *bool            `json:"sanStorageCapabilityFlag:omitempty"`
}

type Hardware_Router_Backend struct {
	Hardware_Router
}

type Hardware_Router_Frontend struct {
	Hardware_Router
}

type Hardware_SecurityModule struct {
	Hardware_Server
}

type Hardware_Server struct {
	Hardware

	ActiveNetworkFirewallBillingItem     *Billing_Item                                         `json:"activeNetworkFirewallBillingItem:omitempty"`
	ActiveTicketCount                    *uint                                                 `json:"activeTicketCount:omitempty"`
	ActiveTickets                        []Ticket                                              `json:"activeTickets:omitempty"`
	ActiveTransaction                    *Provisioning_Version1_Transaction                    `json:"activeTransaction:omitempty"`
	ActiveTransactionCount               *uint                                                 `json:"activeTransactionCount:omitempty"`
	ActiveTransactions                   []Provisioning_Version1_Transaction                   `json:"activeTransactions:omitempty"`
	AvailableMonitoring                  []Network_Monitor_Version1_Query_Host_Stratum         `json:"availableMonitoring:omitempty"`
	AvailableMonitoringCount             *uint                                                 `json:"availableMonitoringCount:omitempty"`
	AverageDailyBandwidthUsage           *float64                                              `json:"averageDailyBandwidthUsage:omitempty"`
	AverageDailyPrivateBandwidthUsage    *float64                                              `json:"averageDailyPrivateBandwidthUsage:omitempty"`
	BillingCycleBandwidthUsage           []Network_Bandwidth_Usage                             `json:"billingCycleBandwidthUsage:omitempty"`
	BillingCycleBandwidthUsageCount      *uint                                                 `json:"billingCycleBandwidthUsageCount:omitempty"`
	BillingCyclePrivateBandwidthUsage    *Network_Bandwidth_Usage                              `json:"billingCyclePrivateBandwidthUsage:omitempty"`
	BillingCyclePublicBandwidthUsage     *Network_Bandwidth_Usage                              `json:"billingCyclePublicBandwidthUsage:omitempty"`
	ContainsSolidStateDrivesFlag         *bool                                                 `json:"containsSolidStateDrivesFlag:omitempty"`
	ControlPanel                         *Software_Component_ControlPanel                      `json:"controlPanel:omitempty"`
	Cost                                 *float64                                              `json:"cost:omitempty"`
	CurrentBandwidthSummary              *Metric_Tracking_Object_Bandwidth_Summary             `json:"currentBandwidthSummary:omitempty"`
	CustomerInstalledOperatingSystemFlag *bool                                                 `json:"customerInstalledOperatingSystemFlag:omitempty"`
	CustomerOwnedFlag                    *bool                                                 `json:"customerOwnedFlag:omitempty"`
	InboundPrivateBandwidthUsage         *float64                                              `json:"inboundPrivateBandwidthUsage:omitempty"`
	LastOperatingSystemReload            *Provisioning_Version1_Transaction                    `json:"lastOperatingSystemReload:omitempty"`
	MetricTrackingObjectId               *int                                                  `json:"metricTrackingObjectId:omitempty"`
	MonitoringUserNotification           []User_Customer_Notification_Hardware                 `json:"monitoringUserNotification:omitempty"`
	MonitoringUserNotificationCount      *uint                                                 `json:"monitoringUserNotificationCount:omitempty"`
	OpenCancellationTicket               *Ticket                                               `json:"openCancellationTicket:omitempty"`
	OutboundPrivateBandwidthUsage        *float64                                              `json:"outboundPrivateBandwidthUsage:omitempty"`
	OverBandwidthAllocationFlag          *int                                                  `json:"overBandwidthAllocationFlag:omitempty"`
	PrivateIpAddress                     *string                                               `json:"privateIpAddress:omitempty"`
	ProjectedOverBandwidthAllocationFlag *int                                                  `json:"projectedOverBandwidthAllocationFlag:omitempty"`
	ProjectedPublicBandwidthUsage        *float64                                              `json:"projectedPublicBandwidthUsage:omitempty"`
	RecentRemoteManagementCommandCount   *uint                                                 `json:"recentRemoteManagementCommandCount:omitempty"`
	RecentRemoteManagementCommands       []Hardware_Component_RemoteManagement_Command_Request `json:"recentRemoteManagementCommands:omitempty"`
	RegionalInternetRegistry             *Network_Regional_Internet_Registry                   `json:"regionalInternetRegistry:omitempty"`
	RemoteManagement                     *Hardware_Component_RemoteManagement                  `json:"remoteManagement:omitempty"`
	RemoteManagementUserCount            *uint                                                 `json:"remoteManagementUserCount:omitempty"`
	RemoteManagementUsers                []Hardware_Component_RemoteManagement_User            `json:"remoteManagementUsers:omitempty"`
	StatisticsRemoteManagement           *Hardware_Component_RemoteManagement                  `json:"statisticsRemoteManagement:omitempty"`
	UserCount                            *uint                                                 `json:"userCount:omitempty"`
	Users                                []User_Customer                                       `json:"users:omitempty"`
	VirtualGuestCount                    *uint                                                 `json:"virtualGuestCount:omitempty"`
	VirtualGuests                        []Virtual_Guest                                       `json:"virtualGuests:omitempty"`
}

type Hardware_Status struct {
	Entity

	Id     *int    `json:"id:omitempty"`
	Status *string `json:"status:omitempty"`
}

type Hardware_Switch struct {
	Hardware
}

type Layout_Container struct {
	Entity

	Id                    *int                   `json:"id:omitempty"`
	Keyname               *string                `json:"keyname:omitempty"`
	LayoutContainerType   *Layout_Container_Type `json:"layoutContainerType:omitempty"`
	LayoutContainerTypeId *int                   `json:"layoutContainerTypeId:omitempty"`
	LayoutItemCount       *uint                  `json:"layoutItemCount:omitempty"`
	LayoutItems           []Layout_Item          `json:"layoutItems:omitempty"`
	Name                  *string                `json:"name:omitempty"`
}

type Layout_Container_Type struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	Keyname *string `json:"keyname:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Layout_Item struct {
	Entity

	Id                        *int                `json:"id:omitempty"`
	Keyname                   *string             `json:"keyname:omitempty"`
	LayoutItemPreferenceCount *uint               `json:"layoutItemPreferenceCount:omitempty"`
	LayoutItemPreferences     []Layout_Preference `json:"layoutItemPreferences:omitempty"`
	LayoutItemType            *Layout_Item_Type   `json:"layoutItemType:omitempty"`
	LayoutItemTypeId          *int                `json:"layoutItemTypeId:omitempty"`
	Name                      *string             `json:"name:omitempty"`
}

type Layout_Item_Type struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	Keyname *string `json:"keyname:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Layout_Preference struct {
	Entity

	Id                     *int                    `json:"id:omitempty"`
	LayoutPreferenceType   *Layout_Preference_Type `json:"layoutPreferenceType:omitempty"`
	LayoutPreferenceTypeId *int                    `json:"layoutPreferenceTypeId:omitempty"`
	Value                  *string                 `json:"value:omitempty"`
}

type Layout_Preference_Type struct {
	Entity

	Id              *int    `json:"id:omitempty"`
	Keyname         *string `json:"keyname:omitempty"`
	Name            *string `json:"name:omitempty"`
	ValueExpression *string `json:"valueExpression:omitempty"`
}

type Layout_Profile struct {
	Entity

	ActiveFlag            *int                        `json:"activeFlag:omitempty"`
	CreateDate            *time.Time                  `json:"createDate:omitempty"`
	Id                    *int                        `json:"id:omitempty"`
	LayoutContainerCount  *uint                       `json:"layoutContainerCount:omitempty"`
	LayoutContainers      []Layout_Container          `json:"layoutContainers:omitempty"`
	LayoutPreferenceCount *uint                       `json:"layoutPreferenceCount:omitempty"`
	LayoutPreferences     []Layout_Profile_Preference `json:"layoutPreferences:omitempty"`
	ModifyDate            *time.Time                  `json:"modifyDate:omitempty"`
	Name                  *string                     `json:"name:omitempty"`
	UserRecordId          *int                        `json:"userRecordId:omitempty"`
}

type Layout_Profile_Containers struct {
	Entity

	CreateDate          *time.Time        `json:"createDate:omitempty"`
	Id                  *int              `json:"id:omitempty"`
	LayoutContainerId   *int              `json:"layoutContainerId:omitempty"`
	LayoutContainerType *Layout_Container `json:"layoutContainerType:omitempty"`
	LayoutProfile       *Layout_Profile   `json:"layoutProfile:omitempty"`
	LayoutProfileId     *int              `json:"layoutProfileId:omitempty"`
	ModifyDate          *time.Time        `json:"modifyDate:omitempty"`
}

type Layout_Profile_Customer struct {
	Layout_Profile

	UserRecord *User_Customer `json:"userRecord:omitempty"`
}

type Layout_Profile_Preference struct {
	Entity

	CreateDate         *time.Time         `json:"createDate:omitempty"`
	DefaultValueFlag   *int               `json:"defaultValueFlag:omitempty"`
	LayoutContainer    *Layout_Container  `json:"layoutContainer:omitempty"`
	LayoutContainerId  *int               `json:"layoutContainerId:omitempty"`
	LayoutItem         *Layout_Item       `json:"layoutItem:omitempty"`
	LayoutItemId       *int               `json:"layoutItemId:omitempty"`
	LayoutPreference   *Layout_Preference `json:"layoutPreference:omitempty"`
	LayoutPreferenceId *int               `json:"layoutPreferenceId:omitempty"`
	LayoutProfile      *Layout_Profile    `json:"layoutProfile:omitempty"`
	LayoutProfileId    *int               `json:"layoutProfileId:omitempty"`
	ModifyDate         *time.Time         `json:"modifyDate:omitempty"`
	Value              *string            `json:"value:omitempty"`
}

type Legal_RegulatedWorkload struct {
	Entity

	Account        *Account                      `json:"account:omitempty"`
	AccountId      *int                          `json:"accountId:omitempty"`
	EnabledFlag    *bool                         `json:"enabledFlag:omitempty"`
	Id             *int                          `json:"id:omitempty"`
	Type           *Legal_RegulatedWorkload_Type `json:"type:omitempty"`
	WorkloadTypeId *int                          `json:"workloadTypeId:omitempty"`
}

type Legal_RegulatedWorkload_Type struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Locale struct {
	Entity

	FriendlyName *string `json:"friendlyName:omitempty"`
	Id           *int    `json:"id:omitempty"`
	LanguageTag  *string `json:"languageTag:omitempty"`
	Name         *string `json:"name:omitempty"`
}

type Locale_Country struct {
	Entity

	IsEuropeanUnionFlag *int                   `json:"isEuropeanUnionFlag:omitempty"`
	LongName            *string                `json:"longName:omitempty"`
	ShortName           *string                `json:"shortName:omitempty"`
	StateCount          *uint                  `json:"stateCount:omitempty"`
	States              []Locale_StateProvince `json:"states:omitempty"`
}

type Locale_StateProvince struct {
	Entity

	LongName  *string `json:"longName:omitempty"`
	ShortName *string `json:"shortName:omitempty"`
}

type Locale_Timezone struct {
	Entity

	Id        *int    `json:"id:omitempty"`
	LongName  *string `json:"longName:omitempty"`
	Name      *string `json:"name:omitempty"`
	Offset    *string `json:"offset:omitempty"`
	ShortName *string `json:"shortName:omitempty"`
}

type Location struct {
	Entity

	BackboneDependentCount        *uint                                   `json:"backboneDependentCount:omitempty"`
	BackboneDependents            []Network_Backbone_Location_Dependent   `json:"backboneDependents:omitempty"`
	GroupCount                    *uint                                   `json:"groupCount:omitempty"`
	Groups                        []Location_Group                        `json:"groups:omitempty"`
	HardwareFirewallCount         *uint                                   `json:"hardwareFirewallCount:omitempty"`
	HardwareFirewalls             []Hardware                              `json:"hardwareFirewalls:omitempty"`
	Id                            *int                                    `json:"id:omitempty"`
	LocationAddress               *Account_Address                        `json:"locationAddress:omitempty"`
	LocationReservationMember     *Location_Reservation_Rack_Member       `json:"locationReservationMember:omitempty"`
	LocationStatus                *Location_Status                        `json:"locationStatus:omitempty"`
	LongName                      *string                                 `json:"longName:omitempty"`
	Name                          *string                                 `json:"name:omitempty"`
	NetworkConfigurationAttribute *Hardware_Attribute                     `json:"networkConfigurationAttribute:omitempty"`
	OnlinePptpVpnUserCount        *int                                    `json:"onlinePptpVpnUserCount:omitempty"`
	OnlineSslVpnUserCount         *int                                    `json:"onlineSslVpnUserCount:omitempty"`
	PathString                    *string                                 `json:"pathString:omitempty"`
	PriceGroupCount               *uint                                   `json:"priceGroupCount:omitempty"`
	PriceGroups                   []Location_Group                        `json:"priceGroups:omitempty"`
	RegionCount                   *uint                                   `json:"regionCount:omitempty"`
	Regions                       []Location_Region                       `json:"regions:omitempty"`
	StatusId                      *int                                    `json:"statusId:omitempty"`
	Timezone                      *Locale_Timezone                        `json:"timezone:omitempty"`
	VdrGroup                      *Location_Group_Location_CrossReference `json:"vdrGroup:omitempty"`
}

type Location_Datacenter struct {
	Location

	ActiveItemPresaleEventCount  *uint                                        `json:"activeItemPresaleEventCount:omitempty"`
	ActiveItemPresaleEvents      []Sales_Presale_Event                        `json:"activeItemPresaleEvents:omitempty"`
	ActivePresaleEventCount      *uint                                        `json:"activePresaleEventCount:omitempty"`
	ActivePresaleEvents          []Sales_Presale_Event                        `json:"activePresaleEvents:omitempty"`
	BackendHardwareRouterCount   *uint                                        `json:"backendHardwareRouterCount:omitempty"`
	BackendHardwareRouters       []Hardware                                   `json:"backendHardwareRouters:omitempty"`
	BoundSubnetCount             *uint                                        `json:"boundSubnetCount:omitempty"`
	BoundSubnets                 []Network_Subnet                             `json:"boundSubnets:omitempty"`
	BrandCountryRestrictionCount *uint                                        `json:"brandCountryRestrictionCount:omitempty"`
	BrandCountryRestrictions     []Brand_Restriction_Location_CustomerCountry `json:"brandCountryRestrictions:omitempty"`
	FrontendHardwareRouterCount  *uint                                        `json:"frontendHardwareRouterCount:omitempty"`
	FrontendHardwareRouters      []Hardware                                   `json:"frontendHardwareRouters:omitempty"`
	HardwareRouterCount          *uint                                        `json:"hardwareRouterCount:omitempty"`
	HardwareRouters              []Hardware                                   `json:"hardwareRouters:omitempty"`
	PresaleEventCount            *uint                                        `json:"presaleEventCount:omitempty"`
	PresaleEvents                []Sales_Presale_Event                        `json:"presaleEvents:omitempty"`
	RegionalGroup                *Location_Group_Regional                     `json:"regionalGroup:omitempty"`
	RegionalInternetRegistry     *Network_Regional_Internet_Registry          `json:"regionalInternetRegistry:omitempty"`
	RoutableBoundSubnetCount     *uint                                        `json:"routableBoundSubnetCount:omitempty"`
	RoutableBoundSubnets         []Network_Subnet                             `json:"routableBoundSubnets:omitempty"`
}

type Location_Group struct {
	Entity

	Description         *string              `json:"description:omitempty"`
	Id                  *int                 `json:"id:omitempty"`
	LocationCount       *uint                `json:"locationCount:omitempty"`
	LocationGroupType   *Location_Group_Type `json:"locationGroupType:omitempty"`
	LocationGroupTypeId *int                 `json:"locationGroupTypeId:omitempty"`
	Locations           []Location           `json:"locations:omitempty"`
	Name                *string              `json:"name:omitempty"`
	SecurityLevelId     *int                 `json:"securityLevelId:omitempty"`
}

type Location_Group_Location_CrossReference struct {
	Entity

	Location        *Location       `json:"location:omitempty"`
	LocationGroup   *Location_Group `json:"locationGroup:omitempty"`
	LocationGroupId *int            `json:"locationGroupId:omitempty"`
	LocationId      *int            `json:"locationId:omitempty"`
	Priority        *int            `json:"priority:omitempty"`
}

type Location_Group_Pricing struct {
	Location_Group

	PriceCount *uint                `json:"priceCount:omitempty"`
	Prices     []Product_Item_Price `json:"prices:omitempty"`
}

type Location_Group_Regional struct {
	Location_Group

	DatacenterCount     *uint                `json:"datacenterCount:omitempty"`
	Datacenters         []Location           `json:"datacenters:omitempty"`
	PreferredDatacenter *Location_Datacenter `json:"preferredDatacenter:omitempty"`
}

type Location_Group_Type struct {
	Entity

	Name *string `json:"name:omitempty"`
}

type Location_Inventory_Room struct {
	Location
}

type Location_Network_Operations_Center struct {
	Location
}

type Location_Office struct {
	Location
}

type Location_Rack struct {
	Location
}

type Location_Region struct {
	Entity

	Description *string                   `json:"description:omitempty"`
	Keyname     *string                   `json:"keyname:omitempty"`
	Location    *Location_Region_Location `json:"location:omitempty"`
	SortOrder   *int                      `json:"sortOrder:omitempty"`
}

type Location_Region_Location struct {
	Entity

	Location                   *Location                   `json:"location:omitempty"`
	LocationPackageDetailCount *uint                       `json:"locationPackageDetailCount:omitempty"`
	LocationPackageDetails     []Product_Package_Locations `json:"locationPackageDetails:omitempty"`
	Region                     *Location_Region            `json:"region:omitempty"`
}

type Location_Reservation struct {
	Entity

	Account                 *Account                              `json:"account:omitempty"`
	Allotment               *Network_Bandwidth_Version1_Allotment `json:"allotment:omitempty"`
	AllotmentId             *int                                  `json:"allotmentId:omitempty"`
	BillingItem             *Billing_Item                         `json:"billingItem:omitempty"`
	Id                      *int                                  `json:"id:omitempty"`
	Location                *Location                             `json:"location:omitempty"`
	LocationId              *int                                  `json:"locationId:omitempty"`
	LocationReservationRack *Location_Reservation_Rack            `json:"locationReservationRack:omitempty"`
	Name                    *string                               `json:"name:omitempty"`
	Notes                   *string                               `json:"notes:omitempty"`
}

type Location_Reservation_Rack struct {
	Entity

	Allotment                    *Network_Bandwidth_Version1_Allotment `json:"allotment:omitempty"`
	Children                     []Location_Reservation_Rack_Member    `json:"children:omitempty"`
	ChildrenCount                *uint                                 `json:"childrenCount:omitempty"`
	Location                     *Location                             `json:"location:omitempty"`
	LocationId                   *int                                  `json:"locationId:omitempty"`
	LocationReservation          *Location_Reservation                 `json:"locationReservation:omitempty"`
	LocationReservationId        *int                                  `json:"locationReservationId:omitempty"`
	NetworkConnectionCapacity    *int                                  `json:"networkConnectionCapacity:omitempty"`
	NetworkConnectionReservation *int                                  `json:"networkConnectionReservation:omitempty"`
	PowerConnectionCapacity      *int                                  `json:"powerConnectionCapacity:omitempty"`
	PowerConnectionReservation   *int                                  `json:"powerConnectionReservation:omitempty"`
	SlotCapacity                 *int                                  `json:"slotCapacity:omitempty"`
	SlotReservation              *int                                  `json:"slotReservation:omitempty"`
}

type Location_Reservation_Rack_Member struct {
	Entity

	Id                      *int                  `json:"id:omitempty"`
	Location                *Location             `json:"location:omitempty"`
	LocationId              *int                  `json:"locationId:omitempty"`
	LocationReservationRack *Location_Reservation `json:"locationReservationRack:omitempty"`
}

type Location_Root struct {
	Location
}

type Location_Server_Room struct {
	Location
}

type Location_Slot struct {
	Location
}

type Location_Status struct {
	Entity

	Id     *int    `json:"id:omitempty"`
	Status *string `json:"status:omitempty"`
}

type Location_Storage_Room struct {
	Location
}

type Marketplace_EmailDistribution struct {
	Entity

	Email *string `json:"email:omitempty"`
	Id    *int    `json:"id:omitempty"`
}

type Marketplace_Partner struct {
	Entity

	AccountId               *int                             `json:"accountId:omitempty"`
	AttachedFiles           []Marketplace_Partner_Attachment `json:"attachedFiles:omitempty"`
	AttachmentCount         *uint                            `json:"attachmentCount:omitempty"`
	Attachments             []Marketplace_Partner_Attachment `json:"attachments:omitempty"`
	CompanyDescription      *string                          `json:"companyDescription:omitempty"`
	CompanyName             *string                          `json:"companyName:omitempty"`
	HeadlineDescription     *string                          `json:"headlineDescription:omitempty"`
	Id                      *int                             `json:"id:omitempty"`
	LinkFreeTrial           *string                          `json:"linkFreeTrial:omitempty"`
	LinkOrderPage           *string                          `json:"linkOrderPage:omitempty"`
	LinkWebsite             *string                          `json:"linkWebsite:omitempty"`
	LogoMedium              *Marketplace_Partner_Attachment  `json:"logoMedium:omitempty"`
	LogoMediumTemp          *Marketplace_Partner_Attachment  `json:"logoMediumTemp:omitempty"`
	LogoSmall               *Marketplace_Partner_Attachment  `json:"logoSmall:omitempty"`
	LogoSmallTemp           *Marketplace_Partner_Attachment  `json:"logoSmallTemp:omitempty"`
	MetaDescription         *string                          `json:"metaDescription:omitempty"`
	MetaKeywords            *string                          `json:"metaKeywords:omitempty"`
	ProductBenefits         *string                          `json:"productBenefits:omitempty"`
	ProductCategoryId       *int                             `json:"productCategoryId:omitempty"`
	ProductDescriptionLong  *string                          `json:"productDescriptionLong:omitempty"`
	ProductDescriptionShort *string                          `json:"productDescriptionShort:omitempty"`
	ProductFeatures         *string                          `json:"productFeatures:omitempty"`
	ProductName             *string                          `json:"productName:omitempty"`
	ProductTitle            *string                          `json:"productTitle:omitempty"`
	UrlIdentifier           *string                          `json:"urlIdentifier:omitempty"`
}

type Marketplace_Partner_Attachment struct {
	Entity

	AttachmentType       *Marketplace_Partner_Attachment_Type `json:"attachmentType:omitempty"`
	AttachmentTypeId     *int                                 `json:"attachmentTypeId:omitempty"`
	BaseName             *string                              `json:"baseName:omitempty"`
	DisplayName          *string                              `json:"displayName:omitempty"`
	FileName             *string                              `json:"fileName:omitempty"`
	Id                   *int                                 `json:"id:omitempty"`
	MarketplacePartnerId *int                                 `json:"marketplacePartnerId:omitempty"`
	SaveAsName           *string                              `json:"saveAsName:omitempty"`
}

type Marketplace_Partner_Attachment_Type struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Type    *string `json:"type:omitempty"`
}

type Marketplace_Partner_File struct {
	Entity

	Attributes *Marketplace_Partner_File_Attributes `json:"attributes:omitempty"`
	Contents   *[]byte                              `json:"contents:omitempty"`
}

type Marketplace_Partner_File_Attributes struct {
	Entity

	Bits           *int    `json:"bits:omitempty"`
	Channels       *int    `json:"channels:omitempty"`
	Height         *int    `json:"height:omitempty"`
	HtmlAttributes *string `json:"htmlAttributes:omitempty"`
	ImageType      *int    `json:"imageType:omitempty"`
	IsImage        *bool   `json:"isImage:omitempty"`
	MimeType       *string `json:"mimeType:omitempty"`
	Width          *int    `json:"width:omitempty"`
}

type Metric_Tracking_Object struct {
	Entity

	Data            []Metric_Tracking_Object_Data `json:"data:omitempty"`
	Id              *int                          `json:"id:omitempty"`
	Label           *string                       `json:"label:omitempty"`
	ResourceTableId *int                          `json:"resourceTableId:omitempty"`
	StartDate       *time.Time                    `json:"startDate:omitempty"`
	Type            *Metric_Tracking_Object_Type  `json:"type:omitempty"`
}

type Metric_Tracking_Object_Abstract struct {
	Metric_Tracking_Object
}

type Metric_Tracking_Object_Bandwidth_Summary struct {
	Entity

	AllocationAmount            *float64 `json:"allocationAmount:omitempty"`
	AllocationId                *int     `json:"allocationId:omitempty"`
	AmountOut                   *float64 `json:"amountOut:omitempty"`
	AverageDailyUsage           *float64 `json:"averageDailyUsage:omitempty"`
	CurrentlyOverAllocationFlag *int     `json:"currentlyOverAllocationFlag:omitempty"`
	Id                          *int     `json:"id:omitempty"`
	OutboundBandwidthAmount     *float64 `json:"outboundBandwidthAmount:omitempty"`
	ProjectedBandwidthUsage     *float64 `json:"projectedBandwidthUsage:omitempty"`
	ProjectedOverAllocationFlag *int     `json:"projectedOverAllocationFlag:omitempty"`
}

type Metric_Tracking_Object_Data struct {
	Entity

	Counter  *float64   `json:"counter:omitempty"`
	DateTime *time.Time `json:"dateTime:omitempty"`
	Type     *string    `json:"type:omitempty"`
}

type Metric_Tracking_Object_Data_Network_ContentDelivery_Account struct {
	Metric_Tracking_Object_Data

	FileName *string `json:"fileName:omitempty"`
	PopId    *int    `json:"popId:omitempty"`
}

type Metric_Tracking_Object_HardwareServer struct {
	Metric_Tracking_Object_Abstract

	BillingCycleBandwidthUsage             []Network_Bandwidth_Usage `json:"billingCycleBandwidthUsage:omitempty"`
	BillingCycleBandwidthUsageCount        *uint                     `json:"billingCycleBandwidthUsageCount:omitempty"`
	BillingCyclePrivateBandwidthUsage      []Network_Bandwidth_Usage `json:"billingCyclePrivateBandwidthUsage:omitempty"`
	BillingCyclePrivateBandwidthUsageCount *uint                     `json:"billingCyclePrivateBandwidthUsageCount:omitempty"`
	BillingCyclePrivateUsageIn             *float64                  `json:"billingCyclePrivateUsageIn:omitempty"`
	BillingCyclePrivateUsageOut            *float64                  `json:"billingCyclePrivateUsageOut:omitempty"`
	BillingCyclePrivateUsageTotal          *uint                     `json:"billingCyclePrivateUsageTotal:omitempty"`
	BillingCyclePublicBandwidthUsage       *Network_Bandwidth_Usage  `json:"billingCyclePublicBandwidthUsage:omitempty"`
	BillingCyclePublicUsageIn              *float64                  `json:"billingCyclePublicUsageIn:omitempty"`
	BillingCyclePublicUsageOut             *float64                  `json:"billingCyclePublicUsageOut:omitempty"`
	BillingCyclePublicUsageTotal           *uint                     `json:"billingCyclePublicUsageTotal:omitempty"`
	Resource                               *Hardware_Server          `json:"resource:omitempty"`
}

type Metric_Tracking_Object_Type struct {
	Entity

	Keyname *string `json:"keyname:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Metric_Tracking_Object_VirtualDedicatedRack struct {
	Metric_Tracking_Object_Abstract

	BillingCycleBandwidthUsage             []Network_Bandwidth_Usage             `json:"billingCycleBandwidthUsage:omitempty"`
	BillingCycleBandwidthUsageCount        *uint                                 `json:"billingCycleBandwidthUsageCount:omitempty"`
	BillingCyclePrivateBandwidthUsage      []Network_Bandwidth_Usage             `json:"billingCyclePrivateBandwidthUsage:omitempty"`
	BillingCyclePrivateBandwidthUsageCount *uint                                 `json:"billingCyclePrivateBandwidthUsageCount:omitempty"`
	BillingCyclePrivateUsageIn             *float64                              `json:"billingCyclePrivateUsageIn:omitempty"`
	BillingCyclePrivateUsageOut            *float64                              `json:"billingCyclePrivateUsageOut:omitempty"`
	BillingCyclePrivateUsageTotal          *uint                                 `json:"billingCyclePrivateUsageTotal:omitempty"`
	BillingCyclePublicBandwidthUsage       *Network_Bandwidth_Usage              `json:"billingCyclePublicBandwidthUsage:omitempty"`
	BillingCyclePublicUsageIn              *float64                              `json:"billingCyclePublicUsageIn:omitempty"`
	BillingCyclePublicUsageOut             *float64                              `json:"billingCyclePublicUsageOut:omitempty"`
	BillingCyclePublicUsageTotal           *uint                                 `json:"billingCyclePublicUsageTotal:omitempty"`
	Resource                               *Network_Bandwidth_Version1_Allotment `json:"resource:omitempty"`
}

type Metric_Tracking_Object_Virtual_Storage_Repository struct {
	Metric_Tracking_Object_Abstract

	Resource *Virtual_Storage_Repository `json:"resource:omitempty"`
}

type Monitoring_Agent struct {
	Entity

	AgentStatus               *Monitoring_Agent_Status                 `json:"agentStatus:omitempty"`
	ConfigurationProfileCount *uint                                    `json:"configurationProfileCount:omitempty"`
	ConfigurationProfiles     []Configuration_Template_Section_Profile `json:"configurationProfiles:omitempty"`
	ConfigurationTemplate     *Configuration_Template                  `json:"configurationTemplate:omitempty"`
	ConfigurationTemplateId   *int                                     `json:"configurationTemplateId:omitempty"`
	ConfigurationValueCount   *uint                                    `json:"configurationValueCount:omitempty"`
	ConfigurationValues       []Monitoring_Agent_Configuration_Value   `json:"configurationValues:omitempty"`
	CreateDate                *time.Time                               `json:"createDate:omitempty"`
	Hardware                  *Hardware                                `json:"hardware:omitempty"`
	Id                        *int                                     `json:"id:omitempty"`
	ModifyDate                *time.Time                               `json:"modifyDate:omitempty"`
	Name                      *string                                  `json:"name:omitempty"`
	ProductItem               *Product_Item                            `json:"productItem:omitempty"`
	RemoteMonitoringAgentFlag *bool                                    `json:"remoteMonitoringAgentFlag:omitempty"`
	RobotId                   *int                                     `json:"robotId:omitempty"`
	SoftwareDescription       *Software_Description                    `json:"softwareDescription:omitempty"`
	StatusId                  *int                                     `json:"statusId:omitempty"`
	StatusName                *string                                  `json:"statusName:omitempty"`
	VirtualGuest              *Virtual_Guest                           `json:"virtualGuest:omitempty"`
}

type Monitoring_Agent_Configuration_Template_Group struct {
	Entity

	Account                             *Account                                                  `json:"account:omitempty"`
	AccountId                           *int                                                      `json:"accountId:omitempty"`
	ConfigurationTemplateCount          *uint                                                     `json:"configurationTemplateCount:omitempty"`
	ConfigurationTemplateReferenceCount *uint                                                     `json:"configurationTemplateReferenceCount:omitempty"`
	ConfigurationTemplateReferences     []Monitoring_Agent_Configuration_Template_Group_Reference `json:"configurationTemplateReferences:omitempty"`
	ConfigurationTemplates              []Configuration_Template                                  `json:"configurationTemplates:omitempty"`
	CreateDate                          *time.Time                                                `json:"createDate:omitempty"`
	Description                         *string                                                   `json:"description:omitempty"`
	Id                                  *int                                                      `json:"id:omitempty"`
	Item                                *Product_Item                                             `json:"item:omitempty"`
	ItemId                              *int                                                      `json:"itemId:omitempty"`
	ModifyDate                          *time.Time                                                `json:"modifyDate:omitempty"`
	Name                                *string                                                   `json:"name:omitempty"`
}

type Monitoring_Agent_Configuration_Template_Group_Reference struct {
	Entity

	ConfigurationTemplate   *Configuration_Template                        `json:"configurationTemplate:omitempty"`
	ConfigurationTemplateId *int                                           `json:"configurationTemplateId:omitempty"`
	Id                      *int                                           `json:"id:omitempty"`
	TemplateGroup           *Monitoring_Agent_Configuration_Template_Group `json:"templateGroup:omitempty"`
	TemplateGroupId         *int                                           `json:"templateGroupId:omitempty"`
}

type Monitoring_Agent_Configuration_Value struct {
	Entity

	AgentId                   *int                                       `json:"agentId:omitempty"`
	ConfigurationDefinitionId *int                                       `json:"configurationDefinitionId:omitempty"`
	Definition                *Configuration_Template_Section_Definition `json:"definition:omitempty"`
	Description               *string                                    `json:"description:omitempty"`
	Id                        *int                                       `json:"id:omitempty"`
	MetricDataType            *Container_Metric_Data_Type                `json:"metricDataType:omitempty"`
	MonitoringAgent           *Monitoring_Agent                          `json:"monitoringAgent:omitempty"`
	Profile                   *Configuration_Template_Section_Profile    `json:"profile:omitempty"`
	ProfileId                 *int                                       `json:"profileId:omitempty"`
	Value                     *string                                    `json:"value:omitempty"`
}

type Monitoring_Agent_Status struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Monitoring_Robot struct {
	Entity

	Account              *Account                 `json:"account:omitempty"`
	AccountId            *int                     `json:"accountId:omitempty"`
	Id                   *int                     `json:"id:omitempty"`
	MonitoringAgentCount *uint                    `json:"monitoringAgentCount:omitempty"`
	MonitoringAgents     []Monitoring_Agent       `json:"monitoringAgents:omitempty"`
	Name                 *string                  `json:"name:omitempty"`
	RobotStatus          *Monitoring_Robot_Status `json:"robotStatus:omitempty"`
	SoftwareComponent    *Software_Component      `json:"softwareComponent:omitempty"`
	StatusId             *int                     `json:"statusId:omitempty"`
}

type Monitoring_Robot_Status struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Network struct {
	Entity

	AccountId         *int             `json:"accountId:omitempty"`
	Cidr              *int             `json:"cidr:omitempty"`
	Id                *int             `json:"id:omitempty"`
	Name              *string          `json:"name:omitempty"`
	NetworkIdentifier *string          `json:"networkIdentifier:omitempty"`
	Notes             *string          `json:"notes:omitempty"`
	SubnetCount       *uint            `json:"subnetCount:omitempty"`
	Subnets           []Network_Subnet `json:"subnets:omitempty"`
}

type Network_Application_Delivery_Controller struct {
	Entity

	Account                          *Account                                                                `json:"account:omitempty"`
	AccountId                        *int                                                                    `json:"accountId:omitempty"`
	AverageDailyPublicBandwidthUsage *float64                                                                `json:"averageDailyPublicBandwidthUsage:omitempty"`
	BillingItem                      *Billing_Item_Network_Application_Delivery_Controller                   `json:"billingItem:omitempty"`
	ConfigurationHistory             []Network_Application_Delivery_Controller_Configuration_History         `json:"configurationHistory:omitempty"`
	ConfigurationHistoryCount        *uint                                                                   `json:"configurationHistoryCount:omitempty"`
	CreateDate                       *time.Time                                                              `json:"createDate:omitempty"`
	Datacenter                       *Location                                                               `json:"datacenter:omitempty"`
	Description                      *string                                                                 `json:"description:omitempty"`
	Id                               *int                                                                    `json:"id:omitempty"`
	LicenseExpirationDate            *time.Time                                                              `json:"licenseExpirationDate:omitempty"`
	LoadBalancerCount                *uint                                                                   `json:"loadBalancerCount:omitempty"`
	LoadBalancers                    []Network_LoadBalancer_VirtualIpAddress                                 `json:"loadBalancers:omitempty"`
	ManagedResourceFlag              *bool                                                                   `json:"managedResourceFlag:omitempty"`
	ManagementIpAddress              *string                                                                 `json:"managementIpAddress:omitempty"`
	ModifyDate                       *time.Time                                                              `json:"modifyDate:omitempty"`
	Name                             *string                                                                 `json:"name:omitempty"`
	NetworkVlan                      *Network_Vlan                                                           `json:"networkVlan:omitempty"`
	NetworkVlanCount                 *uint                                                                   `json:"networkVlanCount:omitempty"`
	NetworkVlans                     []Network_Vlan                                                          `json:"networkVlans:omitempty"`
	Notes                            *string                                                                 `json:"notes:omitempty"`
	OutboundPublicBandwidthUsage     *float64                                                                `json:"outboundPublicBandwidthUsage:omitempty"`
	Password                         *Software_Component_Password                                            `json:"password:omitempty"`
	PrimaryIpAddress                 *string                                                                 `json:"primaryIpAddress:omitempty"`
	ProjectedPublicBandwidthUsage    *float64                                                                `json:"projectedPublicBandwidthUsage:omitempty"`
	SubnetCount                      *uint                                                                   `json:"subnetCount:omitempty"`
	Subnets                          []Network_Subnet                                                        `json:"subnets:omitempty"`
	TagReferenceCount                *uint                                                                   `json:"tagReferenceCount:omitempty"`
	TagReferences                    []Tag_Reference                                                         `json:"tagReferences:omitempty"`
	Type                             *Network_Application_Delivery_Controller_Type                           `json:"type:omitempty"`
	TypeId                           *int                                                                    `json:"typeId:omitempty"`
	VirtualIpAddressCount            *uint                                                                   `json:"virtualIpAddressCount:omitempty"`
	VirtualIpAddresses               []Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"virtualIpAddresses:omitempty"`
}

type Network_Application_Delivery_Controller_Configuration_History struct {
	Entity

	Controller *Network_Application_Delivery_Controller `json:"controller:omitempty"`
	CreateDate *time.Time                               `json:"createDate:omitempty"`
	Id         *int                                     `json:"id:omitempty"`
	Notes      *string                                  `json:"notes:omitempty"`
}

type Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute struct {
	Entity

	HealthAttributeTypeId *int                                                                        `json:"healthAttributeTypeId:omitempty"`
	HealthCheck           *Network_Application_Delivery_Controller_LoadBalancer_Health_Check          `json:"healthCheck:omitempty"`
	HealthCheckId         *int                                                                        `json:"healthCheckId:omitempty"`
	Id                    *int                                                                        `json:"id:omitempty"`
	Type                  *Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type `json:"type:omitempty"`
	Value                 *string                                                                     `json:"value:omitempty"`
}

type Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type struct {
	Entity

	Description     *string `json:"description:omitempty"`
	Id              *int    `json:"id:omitempty"`
	Keyname         *string `json:"keyname:omitempty"`
	Name            *string `json:"name:omitempty"`
	ValueExpression *string `json:"valueExpression:omitempty"`
}

type Network_Application_Delivery_Controller_LoadBalancer_Health_Check struct {
	Entity

	AttributeCount         *uint                                                                   `json:"attributeCount:omitempty"`
	Attributes             []Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute `json:"attributes:omitempty"`
	HealthCheckTypeId      *int                                                                    `json:"healthCheckTypeId:omitempty"`
	Id                     *int                                                                    `json:"id:omitempty"`
	Name                   *string                                                                 `json:"name:omitempty"`
	Notes                  *string                                                                 `json:"notes:omitempty"`
	ScaleLoadBalancerCount *uint                                                                   `json:"scaleLoadBalancerCount:omitempty"`
	ScaleLoadBalancers     []Scale_LoadBalancer                                                    `json:"scaleLoadBalancers:omitempty"`
	ServiceCount           *uint                                                                   `json:"serviceCount:omitempty"`
	Services               []Network_Application_Delivery_Controller_LoadBalancer_Service          `json:"services:omitempty"`
	Type                   *Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type `json:"type:omitempty"`
}

type Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	Keyname *string `json:"keyname:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Network_Application_Delivery_Controller_LoadBalancer_Routing_Method struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	Keyname *string `json:"keyname:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Network_Application_Delivery_Controller_LoadBalancer_Routing_Type struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	Keyname *string `json:"keyname:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Network_Application_Delivery_Controller_LoadBalancer_Service struct {
	Entity

	Enabled             *int                                                                                `json:"enabled:omitempty"`
	GroupCount          *uint                                                                               `json:"groupCount:omitempty"`
	GroupReferenceCount *uint                                                                               `json:"groupReferenceCount:omitempty"`
	GroupReferences     []Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference `json:"groupReferences:omitempty"`
	Groups              []Network_Application_Delivery_Controller_LoadBalancer_Service_Group                `json:"groups:omitempty"`
	HealthCheck         *Network_Application_Delivery_Controller_LoadBalancer_Health_Check                  `json:"healthCheck:omitempty"`
	HealthCheckCount    *uint                                                                               `json:"healthCheckCount:omitempty"`
	HealthChecks        []Network_Application_Delivery_Controller_LoadBalancer_Health_Check                 `json:"healthChecks:omitempty"`
	Id                  *int                                                                                `json:"id:omitempty"`
	IpAddress           *Network_Subnet_IpAddress                                                           `json:"ipAddress:omitempty"`
	IpAddressId         *int                                                                                `json:"ipAddressId:omitempty"`
	Name                *string                                                                             `json:"name:omitempty"`
	Notes               *string                                                                             `json:"notes:omitempty"`
	Port                *int                                                                                `json:"port:omitempty"`
	ServiceGroup        *Network_Application_Delivery_Controller_LoadBalancer_Service_Group                 `json:"serviceGroup:omitempty"`
	Status              *string                                                                             `json:"status:omitempty"`
}

type Network_Application_Delivery_Controller_LoadBalancer_Service_Group struct {
	Entity

	Id                    *int                                                                                `json:"id:omitempty"`
	Name                  *string                                                                             `json:"name:omitempty"`
	Notes                 *string                                                                             `json:"notes:omitempty"`
	RoutingMethod         *Network_Application_Delivery_Controller_LoadBalancer_Routing_Method                `json:"routingMethod:omitempty"`
	RoutingMethodId       *int                                                                                `json:"routingMethodId:omitempty"`
	RoutingType           *Network_Application_Delivery_Controller_LoadBalancer_Routing_Type                  `json:"routingType:omitempty"`
	RoutingTypeId         *int                                                                                `json:"routingTypeId:omitempty"`
	ServiceCount          *uint                                                                               `json:"serviceCount:omitempty"`
	ServiceReferenceCount *uint                                                                               `json:"serviceReferenceCount:omitempty"`
	ServiceReferences     []Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference `json:"serviceReferences:omitempty"`
	Services              []Network_Application_Delivery_Controller_LoadBalancer_Service                      `json:"services:omitempty"`
	Timeout               *int                                                                                `json:"timeout:omitempty"`
	VirtualServer         *Network_Application_Delivery_Controller_LoadBalancer_VirtualServer                 `json:"virtualServer:omitempty"`
	VirtualServerCount    *uint                                                                               `json:"virtualServerCount:omitempty"`
	VirtualServers        []Network_Application_Delivery_Controller_LoadBalancer_VirtualServer                `json:"virtualServers:omitempty"`
}

type Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference struct {
	Entity

	Service        *Network_Application_Delivery_Controller_LoadBalancer_Service       `json:"service:omitempty"`
	ServiceGroup   *Network_Application_Delivery_Controller_LoadBalancer_Service_Group `json:"serviceGroup:omitempty"`
	ServiceGroupId *int                                                                `json:"serviceGroupId:omitempty"`
	ServiceId      *int                                                                `json:"serviceId:omitempty"`
	Weight         *int                                                                `json:"weight:omitempty"`
}

type Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress struct {
	Entity

	Account                            *Account                                                                                        `json:"account:omitempty"`
	AccountId                          *int                                                                                            `json:"accountId:omitempty"`
	ApplicationDeliveryController      *Network_Application_Delivery_Controller                                                        `json:"applicationDeliveryController:omitempty"`
	ApplicationDeliveryControllerCount *uint                                                                                           `json:"applicationDeliveryControllerCount:omitempty"`
	ApplicationDeliveryControllers     []Network_Application_Delivery_Controller                                                       `json:"applicationDeliveryControllers:omitempty"`
	BillingItem                        *Billing_Item                                                                                   `json:"billingItem:omitempty"`
	ConnectionLimit                    *int                                                                                            `json:"connectionLimit:omitempty"`
	ConnectionLimitUnits               *string                                                                                         `json:"connectionLimitUnits:omitempty"`
	DedicatedBillingItem               *Billing_Item_Network_LoadBalancer                                                              `json:"dedicatedBillingItem:omitempty"`
	DedicatedFlag                      *bool                                                                                           `json:"dedicatedFlag:omitempty"`
	HighAvailabilityFlag               *bool                                                                                           `json:"highAvailabilityFlag:omitempty"`
	Id                                 *int                                                                                            `json:"id:omitempty"`
	IpAddress                          *Network_Subnet_IpAddress                                                                       `json:"ipAddress:omitempty"`
	IpAddressId                        *int                                                                                            `json:"ipAddressId:omitempty"`
	LoadBalancerHardware               []Hardware                                                                                      `json:"loadBalancerHardware:omitempty"`
	LoadBalancerHardwareCount          *uint                                                                                           `json:"loadBalancerHardwareCount:omitempty"`
	ManagedResourceFlag                *bool                                                                                           `json:"managedResourceFlag:omitempty"`
	Notes                              *string                                                                                         `json:"notes:omitempty"`
	SecureTransportCipherCount         *uint                                                                                           `json:"secureTransportCipherCount:omitempty"`
	SecureTransportCiphers             []Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportCipher   `json:"secureTransportCiphers:omitempty"`
	SecureTransportProtocolCount       *uint                                                                                           `json:"secureTransportProtocolCount:omitempty"`
	SecureTransportProtocols           []Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportProtocol `json:"secureTransportProtocols:omitempty"`
	SecurityCertificate                *Security_Certificate                                                                           `json:"securityCertificate:omitempty"`
	SecurityCertificateEntry           *Security_Certificate_Entry                                                                     `json:"securityCertificateEntry:omitempty"`
	SecurityCertificateId              *int                                                                                            `json:"securityCertificateId:omitempty"`
	SslActiveFlag                      *bool                                                                                           `json:"sslActiveFlag:omitempty"`
	SslEnabledFlag                     *bool                                                                                           `json:"sslEnabledFlag:omitempty"`
	VirtualServerCount                 *uint                                                                                           `json:"virtualServerCount:omitempty"`
	VirtualServers                     []Network_Application_Delivery_Controller_LoadBalancer_VirtualServer                            `json:"virtualServers:omitempty"`
}

type Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportCipher struct {
	Entity

	Id                 *int                                                                   `json:"id:omitempty"`
	KeyName            *string                                                                `json:"keyName:omitempty"`
	VirtualIpAddress   *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"virtualIpAddress:omitempty"`
	VirtualIpAddressId *int                                                                   `json:"virtualIpAddressId:omitempty"`
}

type Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportProtocol struct {
	Entity

	Id                 *int                                                                   `json:"id:omitempty"`
	KeyName            *string                                                                `json:"keyName:omitempty"`
	VirtualIpAddress   *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"virtualIpAddress:omitempty"`
	VirtualIpAddressId *int                                                                   `json:"virtualIpAddressId:omitempty"`
}

type Network_Application_Delivery_Controller_LoadBalancer_VirtualServer struct {
	Entity

	Allocation             *int                                                                   `json:"allocation:omitempty"`
	Id                     *int                                                                   `json:"id:omitempty"`
	Name                   *string                                                                `json:"name:omitempty"`
	Notes                  *string                                                                `json:"notes:omitempty"`
	Port                   *int                                                                   `json:"port:omitempty"`
	RoutingMethod          *Network_Application_Delivery_Controller_LoadBalancer_Routing_Method   `json:"routingMethod:omitempty"`
	RoutingMethodId        *int                                                                   `json:"routingMethodId:omitempty"`
	ScaleLoadBalancerCount *uint                                                                  `json:"scaleLoadBalancerCount:omitempty"`
	ScaleLoadBalancers     []Scale_LoadBalancer                                                   `json:"scaleLoadBalancers:omitempty"`
	ServiceGroupCount      *uint                                                                  `json:"serviceGroupCount:omitempty"`
	ServiceGroups          []Network_Application_Delivery_Controller_LoadBalancer_Service_Group   `json:"serviceGroups:omitempty"`
	VirtualIpAddress       *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"virtualIpAddress:omitempty"`
	VirtualIpAddressId     *int                                                                   `json:"virtualIpAddressId:omitempty"`
}

type Network_Application_Delivery_Controller_Type struct {
	Entity

	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Network_Backbone struct {
	Entity

	Capacity           *int               `json:"capacity:omitempty"`
	CapacityUnits      *string            `json:"capacityUnits:omitempty"`
	Health             *string            `json:"health:omitempty"`
	Id                 *int               `json:"id:omitempty"`
	Location           *Location          `json:"location:omitempty"`
	Name               *string            `json:"name:omitempty"`
	NetworkComponent   *Network_Component `json:"networkComponent:omitempty"`
	NetworkComponentId *int               `json:"networkComponentId:omitempty"`
	Type               *string            `json:"type:omitempty"`
}

type Network_Backbone_Location_Dependent struct {
	Entity

	DependentLocation   *Location `json:"dependentLocation:omitempty"`
	DependentLocationId *int      `json:"dependentLocationId:omitempty"`
	Id                  *int      `json:"id:omitempty"`
	SourceLocation      *Location `json:"sourceLocation:omitempty"`
	SourceLocationId    *int      `json:"sourceLocationId:omitempty"`
}

type Network_Bandwidth_Usage struct {
	Entity

	AmountIn                   *float64                                      `json:"amountIn:omitempty"`
	AmountOut                  *float64                                      `json:"amountOut:omitempty"`
	BandwidthUsageDetailTypeId *float64                                      `json:"bandwidthUsageDetailTypeId:omitempty"`
	TrackingObject             *Metric_Tracking_Object                       `json:"trackingObject:omitempty"`
	Type                       *Network_Bandwidth_Version1_Usage_Detail_Type `json:"type:omitempty"`
}

type Network_Bandwidth_Usage_Detail struct {
	Entity

	Account                    *Account                                      `json:"account:omitempty"`
	AmountIn                   *float64                                      `json:"amountIn:omitempty"`
	AmountOut                  *float64                                      `json:"amountOut:omitempty"`
	BandwidthUsageDetailTypeId *float64                                      `json:"bandwidthUsageDetailTypeId:omitempty"`
	TrackingObject             *Metric_Tracking_Object                       `json:"trackingObject:omitempty"`
	Type                       *Network_Bandwidth_Version1_Usage_Detail_Type `json:"type:omitempty"`
}

type Network_Bandwidth_Version1_Allocation struct {
	Entity

	AllotmentDetail *Network_Bandwidth_Version1_Allotment_Detail `json:"allotmentDetail:omitempty"`
	Amount          *float64                                     `json:"amount:omitempty"`
	BillingItem     *Billing_Item_Hardware                       `json:"billingItem:omitempty"`
	Id              *int                                         `json:"id:omitempty"`
}

type Network_Bandwidth_Version1_Allotment struct {
	Entity

	Account                              *Account                                      `json:"account:omitempty"`
	AccountId                            *int                                          `json:"accountId:omitempty"`
	ActiveDetailCount                    *uint                                         `json:"activeDetailCount:omitempty"`
	ActiveDetails                        []Network_Bandwidth_Version1_Allotment_Detail `json:"activeDetails:omitempty"`
	ApplicationDeliveryControllerCount   *uint                                         `json:"applicationDeliveryControllerCount:omitempty"`
	ApplicationDeliveryControllers       []Network_Application_Delivery_Controller     `json:"applicationDeliveryControllers:omitempty"`
	AverageDailyPublicBandwidthUsage     *float64                                      `json:"averageDailyPublicBandwidthUsage:omitempty"`
	BandwidthAllotmentTypeId             *int                                          `json:"bandwidthAllotmentTypeId:omitempty"`
	BareMetalInstanceCount               *uint                                         `json:"bareMetalInstanceCount:omitempty"`
	BareMetalInstances                   []Hardware                                    `json:"bareMetalInstances:omitempty"`
	BillingCycleBandwidthUsage           []Network_Bandwidth_Usage                     `json:"billingCycleBandwidthUsage:omitempty"`
	BillingCycleBandwidthUsageCount      *uint                                         `json:"billingCycleBandwidthUsageCount:omitempty"`
	BillingCyclePrivateBandwidthUsage    *Network_Bandwidth_Usage                      `json:"billingCyclePrivateBandwidthUsage:omitempty"`
	BillingCyclePublicBandwidthUsage     *Network_Bandwidth_Usage                      `json:"billingCyclePublicBandwidthUsage:omitempty"`
	BillingCyclePublicUsageTotal         *uint                                         `json:"billingCyclePublicUsageTotal:omitempty"`
	BillingItem                          *Billing_Item                                 `json:"billingItem:omitempty"`
	CreateDate                           *time.Time                                    `json:"createDate:omitempty"`
	CurrentBandwidthSummary              *Metric_Tracking_Object_Bandwidth_Summary     `json:"currentBandwidthSummary:omitempty"`
	DetailCount                          *uint                                         `json:"detailCount:omitempty"`
	Details                              []Network_Bandwidth_Version1_Allotment_Detail `json:"details:omitempty"`
	EndDate                              *time.Time                                    `json:"endDate:omitempty"`
	Hardware                             []Hardware                                    `json:"hardware:omitempty"`
	HardwareCount                        *uint                                         `json:"hardwareCount:omitempty"`
	Id                                   *int                                          `json:"id:omitempty"`
	InboundPublicBandwidthUsage          *float64                                      `json:"inboundPublicBandwidthUsage:omitempty"`
	LocationGroup                        *Location_Group                               `json:"locationGroup:omitempty"`
	LocationGroupId                      *int                                          `json:"locationGroupId:omitempty"`
	ManagedBareMetalInstanceCount        *uint                                         `json:"managedBareMetalInstanceCount:omitempty"`
	ManagedBareMetalInstances            []Hardware                                    `json:"managedBareMetalInstances:omitempty"`
	ManagedHardware                      []Hardware                                    `json:"managedHardware:omitempty"`
	ManagedHardwareCount                 *uint                                         `json:"managedHardwareCount:omitempty"`
	ManagedVirtualGuestCount             *uint                                         `json:"managedVirtualGuestCount:omitempty"`
	ManagedVirtualGuests                 []Virtual_Guest                               `json:"managedVirtualGuests:omitempty"`
	MetricTrackingObject                 *Metric_Tracking_Object_VirtualDedicatedRack  `json:"metricTrackingObject:omitempty"`
	MetricTrackingObjectId               *int                                          `json:"metricTrackingObjectId:omitempty"`
	Name                                 *string                                       `json:"name:omitempty"`
	OutboundPublicBandwidthUsage         *float64                                      `json:"outboundPublicBandwidthUsage:omitempty"`
	OverBandwidthAllocationFlag          *int                                          `json:"overBandwidthAllocationFlag:omitempty"`
	PrivateNetworkOnlyHardware           []Hardware                                    `json:"privateNetworkOnlyHardware:omitempty"`
	PrivateNetworkOnlyHardwareCount      *uint                                         `json:"privateNetworkOnlyHardwareCount:omitempty"`
	ProjectedOverBandwidthAllocationFlag *int                                          `json:"projectedOverBandwidthAllocationFlag:omitempty"`
	ProjectedPublicBandwidthUsage        *float64                                      `json:"projectedPublicBandwidthUsage:omitempty"`
	ServiceProvider                      *Service_Provider                             `json:"serviceProvider:omitempty"`
	ServiceProviderId                    *int                                          `json:"serviceProviderId:omitempty"`
	TotalBandwidthAllocated              *uint                                         `json:"totalBandwidthAllocated:omitempty"`
	VirtualGuestCount                    *uint                                         `json:"virtualGuestCount:omitempty"`
	VirtualGuests                        []Virtual_Guest                               `json:"virtualGuests:omitempty"`
}

type Network_Bandwidth_Version1_Allotment_Detail struct {
	Entity

	Allocation           *Network_Bandwidth_Version1_Allocation `json:"allocation:omitempty"`
	AllocationId         *int                                   `json:"allocationId:omitempty"`
	BandwidthAllotment   *Network_Bandwidth_Version1_Allotment  `json:"bandwidthAllotment:omitempty"`
	BandwidthAllotmentId *int                                   `json:"bandwidthAllotmentId:omitempty"`
	BandwidthUsage       []Network_Bandwidth_Version1_Usage     `json:"bandwidthUsage:omitempty"`
	BandwidthUsageCount  *uint                                  `json:"bandwidthUsageCount:omitempty"`
	EffectiveDate        *time.Time                             `json:"effectiveDate:omitempty"`
	EndEffectiveDate     *time.Time                             `json:"endEffectiveDate:omitempty"`
	Id                   *int                                   `json:"id:omitempty"`
	ServiceProviderId    *int                                   `json:"serviceProviderId:omitempty"`
}

type Network_Bandwidth_Version1_Host struct {
	Entity

	PodId *int `json:"podId:omitempty"`
}

type Network_Bandwidth_Version1_Interface struct {
	Entity

	Host               *Network_Bandwidth_Version1_Host `json:"host:omitempty"`
	HostId             *int                             `json:"hostId:omitempty"`
	NetworkComponent   *Network_Component               `json:"networkComponent:omitempty"`
	NetworkComponentId *int                             `json:"networkComponentId:omitempty"`
}

type Network_Bandwidth_Version1_Usage struct {
	Entity

	BandwidthAllotmentDetail  *Network_Bandwidth_Version1_Allotment_Detail `json:"bandwidthAllotmentDetail:omitempty"`
	BandwidthUsageDetail      []Network_Bandwidth_Version1_Usage_Detail    `json:"bandwidthUsageDetail:omitempty"`
	BandwidthUsageDetailCount *uint                                        `json:"bandwidthUsageDetailCount:omitempty"`
}

type Network_Bandwidth_Version1_Usage_Detail struct {
	Entity

	AmountIn                 *float64                                      `json:"amountIn:omitempty"`
	AmountOut                *float64                                      `json:"amountOut:omitempty"`
	BandwidthUsage           *Network_Bandwidth_Version1_Usage             `json:"bandwidthUsage:omitempty"`
	BandwidthUsageDetailType *Network_Bandwidth_Version1_Usage_Detail_Type `json:"bandwidthUsageDetailType:omitempty"`
	Day                      *time.Time                                    `json:"day:omitempty"`
}

type Network_Bandwidth_Version1_Usage_Detail_Total struct {
	Entity

	Account                    *Account                                      `json:"account:omitempty"`
	AmountIn                   *float64                                      `json:"amountIn:omitempty"`
	AmountOut                  *float64                                      `json:"amountOut:omitempty"`
	BandwidthUsageDetailTypeId *float64                                      `json:"bandwidthUsageDetailTypeId:omitempty"`
	TrackingObject             *Metric_Tracking_Object                       `json:"trackingObject:omitempty"`
	Type                       *Network_Bandwidth_Version1_Usage_Detail_Type `json:"type:omitempty"`
}

type Network_Bandwidth_Version1_Usage_Detail_Type struct {
	Entity

	Alias *string `json:"alias:omitempty"`
}

type Network_Component struct {
	Entity

	ActiveCommand                  *Hardware_Component_RemoteManagement_Command_Request  `json:"activeCommand:omitempty"`
	DownlinkComponent              *Network_Component                                    `json:"downlinkComponent:omitempty"`
	DuplexMode                     *Network_Component_Duplex_Mode                        `json:"duplexMode:omitempty"`
	DuplexModeId                   *string                                               `json:"duplexModeId:omitempty"`
	Hardware                       *Hardware                                             `json:"hardware:omitempty"`
	HardwareId                     *int                                                  `json:"hardwareId:omitempty"`
	HighAvailabilityFirewallFlag   *bool                                                 `json:"highAvailabilityFirewallFlag:omitempty"`
	Id                             *int                                                  `json:"id:omitempty"`
	Interface                      *Network_Bandwidth_Version1_Interface                 `json:"interface:omitempty"`
	IpAddressBindingCount          *uint                                                 `json:"ipAddressBindingCount:omitempty"`
	IpAddressBindings              []Network_Component_IpAddress                         `json:"ipAddressBindings:omitempty"`
	IpAddressCount                 *uint                                                 `json:"ipAddressCount:omitempty"`
	IpAddresses                    []Network_Subnet_IpAddress                            `json:"ipAddresses:omitempty"`
	IpmiIpAddress                  *string                                               `json:"ipmiIpAddress:omitempty"`
	IpmiMacAddress                 *string                                               `json:"ipmiMacAddress:omitempty"`
	LastCommand                    *Hardware_Component_RemoteManagement_Command_Request  `json:"lastCommand:omitempty"`
	MacAddress                     *string                                               `json:"macAddress:omitempty"`
	MaxSpeed                       *int                                                  `json:"maxSpeed:omitempty"`
	MetricTrackingObject           *Metric_Tracking_Object                               `json:"metricTrackingObject:omitempty"`
	ModifyDate                     *time.Time                                            `json:"modifyDate:omitempty"`
	Name                           *string                                               `json:"name:omitempty"`
	NetworkComponentFirewall       *Network_Component_Firewall                           `json:"networkComponentFirewall:omitempty"`
	NetworkComponentGroup          *Network_Component_Group                              `json:"networkComponentGroup:omitempty"`
	NetworkHardware                []Hardware                                            `json:"networkHardware:omitempty"`
	NetworkHardwareCount           *uint                                                 `json:"networkHardwareCount:omitempty"`
	NetworkVlan                    *Network_Vlan                                         `json:"networkVlan:omitempty"`
	NetworkVlanId                  *int                                                  `json:"networkVlanId:omitempty"`
	NetworkVlanTrunkCount          *uint                                                 `json:"networkVlanTrunkCount:omitempty"`
	NetworkVlanTrunks              []Network_Component_Network_Vlan_Trunk                `json:"networkVlanTrunks:omitempty"`
	Port                           *int                                                  `json:"port:omitempty"`
	PrimaryIpAddress               *string                                               `json:"primaryIpAddress:omitempty"`
	PrimaryIpAddressRecord         *Network_Subnet_IpAddress                             `json:"primaryIpAddressRecord:omitempty"`
	PrimarySubnet                  *Network_Subnet                                       `json:"primarySubnet:omitempty"`
	PrimaryVersion6IpAddressRecord *Network_Subnet_IpAddress                             `json:"primaryVersion6IpAddressRecord:omitempty"`
	RecentCommandCount             *uint                                                 `json:"recentCommandCount:omitempty"`
	RecentCommands                 []Hardware_Component_RemoteManagement_Command_Request `json:"recentCommands:omitempty"`
	RedundancyCapableFlag          *bool                                                 `json:"redundancyCapableFlag:omitempty"`
	RedundancyEnabledFlag          *bool                                                 `json:"redundancyEnabledFlag:omitempty"`
	RemoteManagementUserCount      *uint                                                 `json:"remoteManagementUserCount:omitempty"`
	RemoteManagementUsers          []Hardware_Component_RemoteManagement_User            `json:"remoteManagementUsers:omitempty"`
	Router                         *Hardware                                             `json:"router:omitempty"`
	Speed                          *int                                                  `json:"speed:omitempty"`
	Status                         *string                                               `json:"status:omitempty"`
	StorageNetworkFlag             *bool                                                 `json:"storageNetworkFlag:omitempty"`
	SubnetCount                    *uint                                                 `json:"subnetCount:omitempty"`
	Subnets                        []Network_Subnet                                      `json:"subnets:omitempty"`
	UplinkComponent                *Network_Component                                    `json:"uplinkComponent:omitempty"`
	UplinkDuplexMode               *Network_Component_Duplex_Mode                        `json:"uplinkDuplexMode:omitempty"`
}

type Network_Component_Duplex_Mode struct {
	Entity

	Description *string `json:"description:omitempty"`
	Keyname     *string `json:"keyname:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Network_Component_Firewall struct {
	Entity

	ApplyServerRuleSubnetCount        *uint                             `json:"applyServerRuleSubnetCount:omitempty"`
	ApplyServerRuleSubnets            []Network_Subnet                  `json:"applyServerRuleSubnets:omitempty"`
	BillingItem                       *Billing_Item                     `json:"billingItem:omitempty"`
	GuestNetworkComponent             *Virtual_Guest_Network_Component  `json:"guestNetworkComponent:omitempty"`
	GuestNetworkComponentId           *int                              `json:"guestNetworkComponentId:omitempty"`
	Id                                *int                              `json:"id:omitempty"`
	NetworkComponent                  *Network_Component                `json:"networkComponent:omitempty"`
	NetworkComponentId                *int                              `json:"networkComponentId:omitempty"`
	NetworkFirewallUpdateRequest      []Network_Firewall_Update_Request `json:"networkFirewallUpdateRequest:omitempty"`
	NetworkFirewallUpdateRequestCount *uint                             `json:"networkFirewallUpdateRequestCount:omitempty"`
	RuleCount                         *uint                             `json:"ruleCount:omitempty"`
	Rules                             []Network_Component_Firewall_Rule `json:"rules:omitempty"`
	Status                            *string                           `json:"status:omitempty"`
	SubnetCount                       *uint                             `json:"subnetCount:omitempty"`
	Subnets                           []Network_Subnet                  `json:"subnets:omitempty"`
}

type Network_Component_Firewall_Rule struct {
	Entity

	Action                    *string                     `json:"action:omitempty"`
	DestinationIpAddress      *string                     `json:"destinationIpAddress:omitempty"`
	DestinationIpCidr         *int                        `json:"destinationIpCidr:omitempty"`
	DestinationIpSubnetMask   *string                     `json:"destinationIpSubnetMask:omitempty"`
	DestinationPortRangeEnd   *int                        `json:"destinationPortRangeEnd:omitempty"`
	DestinationPortRangeStart *int                        `json:"destinationPortRangeStart:omitempty"`
	Id                        *int                        `json:"id:omitempty"`
	NetworkComponentFirewall  *Network_Component_Firewall `json:"networkComponentFirewall:omitempty"`
	Notes                     *string                     `json:"notes:omitempty"`
	OrderValue                *int                        `json:"orderValue:omitempty"`
	Protocol                  *string                     `json:"protocol:omitempty"`
	SourceIpAddress           *string                     `json:"sourceIpAddress:omitempty"`
	SourceIpCidr              *int                        `json:"sourceIpCidr:omitempty"`
	SourceIpSubnetMask        *string                     `json:"sourceIpSubnetMask:omitempty"`
	Status                    *string                     `json:"status:omitempty"`
	Version                   *int                        `json:"version:omitempty"`
}

type Network_Component_Firewall_Subnets struct {
	Entity

	ApplyServerRulesFlag     *bool                       `json:"applyServerRulesFlag:omitempty"`
	NetworkComponentFirewall *Network_Component_Firewall `json:"networkComponentFirewall:omitempty"`
	Subnet                   *Network_Subnet             `json:"subnet:omitempty"`
	SubnetId                 *int                        `json:"subnetId:omitempty"`
}

type Network_Component_Group struct {
	Entity

	GroupTypeId           *int                `json:"groupTypeId:omitempty"`
	NetworkComponentCount *uint               `json:"networkComponentCount:omitempty"`
	NetworkComponents     []Network_Component `json:"networkComponents:omitempty"`
}

type Network_Component_IpAddress struct {
	Entity

	IpAddress        *Network_Subnet_IpAddress `json:"ipAddress:omitempty"`
	NetworkComponent *Network_Component        `json:"networkComponent:omitempty"`
}

type Network_Component_Network_Vlan_Trunk struct {
	Entity

	NetworkComponent   *Network_Component `json:"networkComponent:omitempty"`
	NetworkComponentId *int               `json:"networkComponentId:omitempty"`
	NetworkVlan        *Network_Vlan      `json:"networkVlan:omitempty"`
	NetworkVlanId      *int               `json:"networkVlanId:omitempty"`
}

type Network_Component_RemoteManagement struct {
	Network_Component
}

type Network_Component_Uplink_Hardware struct {
	Entity

	Hardware         *Hardware          `json:"hardware:omitempty"`
	NetworkComponent *Network_Component `json:"networkComponent:omitempty"`
}

type Network_ContentDelivery_Account struct {
	Entity

	Account                        *Account                                         `json:"account:omitempty"`
	AccountId                      *int                                             `json:"accountId:omitempty"`
	AssociatedCdnAccountId         *string                                          `json:"associatedCdnAccountId:omitempty"`
	AuthenticationIpAddressCount   *uint                                            `json:"authenticationIpAddressCount:omitempty"`
	AuthenticationIpAddresses      []Network_ContentDelivery_Authentication_Address `json:"authenticationIpAddresses:omitempty"`
	BillingItem                    *Billing_Item                                    `json:"billingItem:omitempty"`
	CdnAccountName                 *string                                          `json:"cdnAccountName:omitempty"`
	CdnAccountNote                 *string                                          `json:"cdnAccountNote:omitempty"`
	CdnSolutionName                *string                                          `json:"cdnSolutionName:omitempty"`
	CreateDate                     *time.Time                                       `json:"createDate:omitempty"`
	DependantServiceFlag           *bool                                            `json:"dependantServiceFlag:omitempty"`
	Id                             *int                                             `json:"id:omitempty"`
	LegacyCdnFlag                  *bool                                            `json:"legacyCdnFlag:omitempty"`
	LogEnabledFlag                 *string                                          `json:"logEnabledFlag:omitempty"`
	ProviderPortalAccessFlag       *bool                                            `json:"providerPortalAccessFlag:omitempty"`
	Status                         *Network_ContentDelivery_Account_Status          `json:"status:omitempty"`
	StatusId                       *int                                             `json:"statusId:omitempty"`
	TokenAuthenticationEnabledFlag *bool                                            `json:"tokenAuthenticationEnabledFlag:omitempty"`
}

type Network_ContentDelivery_Account_Status struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Network_ContentDelivery_Authentication_Address struct {
	Entity

	AccessType   *string    `json:"accessType:omitempty"`
	CdnAccountId *int       `json:"cdnAccountId:omitempty"`
	CreateDate   *time.Time `json:"createDate:omitempty"`
	Id           *int       `json:"id:omitempty"`
	IpAddress    *string    `json:"ipAddress:omitempty"`
	ModifyDate   *time.Time `json:"modifyDate:omitempty"`
	Name         *string    `json:"name:omitempty"`
	Priority     *int       `json:"priority:omitempty"`
	UserId       *int       `json:"userId:omitempty"`
}

type Network_ContentDelivery_Authentication_Token struct {
	Entity

	CdnAccountId *int       `json:"cdnAccountId:omitempty"`
	ClientIp     *string    `json:"clientIp:omitempty"`
	CreateDate   *time.Time `json:"createDate:omitempty"`
	Name         *string    `json:"name:omitempty"`
	Referrer     *string    `json:"referrer:omitempty"`
	Token        *string    `json:"token:omitempty"`
}

type Network_Customer_Subnet struct {
	Entity

	AccountId         *int                                `json:"accountId:omitempty"`
	Cidr              *int                                `json:"cidr:omitempty"`
	Id                *int                                `json:"id:omitempty"`
	IpAddressCount    *uint                               `json:"ipAddressCount:omitempty"`
	IpAddresses       []Network_Customer_Subnet_IpAddress `json:"ipAddresses:omitempty"`
	Netmask           *string                             `json:"netmask:omitempty"`
	NetworkIdentifier *string                             `json:"networkIdentifier:omitempty"`
	TotalIpAddresses  *int                                `json:"totalIpAddresses:omitempty"`
}

type Network_Customer_Subnet_IpAddress struct {
	Entity

	Id               *int                                                `json:"id:omitempty"`
	IpAddress        *string                                             `json:"ipAddress:omitempty"`
	Notes            *string                                             `json:"notes:omitempty"`
	Subnet           *Network_Customer_Subnet                            `json:"subnet:omitempty"`
	SubnetId         *int                                                `json:"subnetId:omitempty"`
	TranslationCount *uint                                               `json:"translationCount:omitempty"`
	Translations     []Network_Tunnel_Module_Context_Address_Translation `json:"translations:omitempty"`
}

type Network_Firewall_AccessControlList struct {
	Entity

	Direction                         *string                           `json:"direction:omitempty"`
	FirewallContextInterfaceId        *int                              `json:"firewallContextInterfaceId:omitempty"`
	Id                                *int                              `json:"id:omitempty"`
	NetworkFirewallUpdateRequestCount *uint                             `json:"networkFirewallUpdateRequestCount:omitempty"`
	NetworkFirewallUpdateRequests     []Network_Firewall_Update_Request `json:"networkFirewallUpdateRequests:omitempty"`
	NetworkVlan                       *Network_Vlan                     `json:"networkVlan:omitempty"`
	RuleCount                         *uint                             `json:"ruleCount:omitempty"`
	Rules                             []Network_Vlan_Firewall_Rule      `json:"rules:omitempty"`
}

type Network_Firewall_Interface struct {
	Network_Firewall_Module_Context_Interface
}

type Network_Firewall_Module_Context_Interface struct {
	Entity

	FirewallContextAccessControlListCount *uint                                `json:"firewallContextAccessControlListCount:omitempty"`
	FirewallContextAccessControlLists     []Network_Firewall_AccessControlList `json:"firewallContextAccessControlLists:omitempty"`
	Id                                    *int                                 `json:"id:omitempty"`
	Name                                  *string                              `json:"name:omitempty"`
	NetworkVlan                           *Network_Vlan                        `json:"networkVlan:omitempty"`
}

type Network_Firewall_Template struct {
	Entity

	Id        *int                             `json:"id:omitempty"`
	Name      *string                          `json:"name:omitempty"`
	RuleCount *uint                            `json:"ruleCount:omitempty"`
	Rules     []Network_Firewall_Template_Rule `json:"rules:omitempty"`
}

type Network_Firewall_Template_Rule struct {
	Entity

	Action                    *string                    `json:"action:omitempty"`
	DestinationIpAddress      *string                    `json:"destinationIpAddress:omitempty"`
	DestinationIpSubnetMask   *string                    `json:"destinationIpSubnetMask:omitempty"`
	DestinationPortRangeEnd   *int                       `json:"destinationPortRangeEnd:omitempty"`
	DestinationPortRangeStart *int                       `json:"destinationPortRangeStart:omitempty"`
	FirewallTemplate          *Network_Firewall_Template `json:"firewallTemplate:omitempty"`
	FirewallTemplateId        *int                       `json:"firewallTemplateId:omitempty"`
	Id                        *int                       `json:"id:omitempty"`
	Notes                     *string                    `json:"notes:omitempty"`
	OrderValue                *int                       `json:"orderValue:omitempty"`
	Protocol                  *string                    `json:"protocol:omitempty"`
	SourceIpAddress           *string                    `json:"sourceIpAddress:omitempty"`
	SourceIpSubnetMask        *string                    `json:"sourceIpSubnetMask:omitempty"`
}

type Network_Firewall_Update_Request struct {
	Entity

	ApplyDate                          *time.Time                             `json:"applyDate:omitempty"`
	AuthorizingUser                    *User_Interface                        `json:"authorizingUser:omitempty"`
	AuthorizingUserId                  *int                                   `json:"authorizingUserId:omitempty"`
	AuthorizingUserType                *string                                `json:"authorizingUserType:omitempty"`
	BypassFlag                         *bool                                  `json:"bypassFlag:omitempty"`
	CreateDate                         *time.Time                             `json:"createDate:omitempty"`
	FirewallContextAccessControlListId *int                                   `json:"firewallContextAccessControlListId:omitempty"`
	Guest                              *Virtual_Guest                         `json:"guest:omitempty"`
	Hardware                           *Hardware                              `json:"hardware:omitempty"`
	HardwareId                         *int                                   `json:"hardwareId:omitempty"`
	Id                                 *int                                   `json:"id:omitempty"`
	NetworkComponentFirewall           *Network_Component_Firewall            `json:"networkComponentFirewall:omitempty"`
	NetworkComponentFirewallId         *int                                   `json:"networkComponentFirewallId:omitempty"`
	RuleCount                          *uint                                  `json:"ruleCount:omitempty"`
	Rules                              []Network_Firewall_Update_Request_Rule `json:"rules:omitempty"`
}

type Network_Firewall_Update_Request_Customer struct {
	Network_Firewall_Update_Request
}

type Network_Firewall_Update_Request_Employee struct {
	Network_Firewall_Update_Request
}

type Network_Firewall_Update_Request_Rule struct {
	Entity

	Action                    *string                          `json:"action:omitempty"`
	DestinationIpAddress      *string                          `json:"destinationIpAddress:omitempty"`
	DestinationIpCidr         *int                             `json:"destinationIpCidr:omitempty"`
	DestinationIpSubnetMask   *string                          `json:"destinationIpSubnetMask:omitempty"`
	DestinationPortRangeEnd   *int                             `json:"destinationPortRangeEnd:omitempty"`
	DestinationPortRangeStart *int                             `json:"destinationPortRangeStart:omitempty"`
	FirewallUpdateRequest     *Network_Firewall_Update_Request `json:"firewallUpdateRequest:omitempty"`
	FirewallUpdateRequestId   *int                             `json:"firewallUpdateRequestId:omitempty"`
	Id                        *int                             `json:"id:omitempty"`
	Notes                     *string                          `json:"notes:omitempty"`
	OrderValue                *int                             `json:"orderValue:omitempty"`
	Protocol                  *string                          `json:"protocol:omitempty"`
	SourceIpAddress           *string                          `json:"sourceIpAddress:omitempty"`
	SourceIpCidr              *int                             `json:"sourceIpCidr:omitempty"`
	SourceIpSubnetMask        *string                          `json:"sourceIpSubnetMask:omitempty"`
	Version                   *int                             `json:"version:omitempty"`
}

type Network_Firewall_Update_Request_Rule_Version6 struct {
	Network_Firewall_Update_Request_Rule
}

type Network_Gateway struct {
	Entity

	Account             *Account                  `json:"account:omitempty"`
	AccountId           *int                      `json:"accountId:omitempty"`
	GroupNumber         *int                      `json:"groupNumber:omitempty"`
	Id                  *int                      `json:"id:omitempty"`
	InsideVlanCount     *uint                     `json:"insideVlanCount:omitempty"`
	InsideVlans         []Network_Gateway_Vlan    `json:"insideVlans:omitempty"`
	MemberCount         *uint                     `json:"memberCount:omitempty"`
	Members             []Network_Gateway_Member  `json:"members:omitempty"`
	Name                *string                   `json:"name:omitempty"`
	NetworkSpace        *string                   `json:"networkSpace:omitempty"`
	PrivateIpAddress    *Network_Subnet_IpAddress `json:"privateIpAddress:omitempty"`
	PrivateIpAddressId  *int                      `json:"privateIpAddressId:omitempty"`
	PrivateVlan         *Network_Vlan             `json:"privateVlan:omitempty"`
	PrivateVlanId       *int                      `json:"privateVlanId:omitempty"`
	PublicIpAddress     *Network_Subnet_IpAddress `json:"publicIpAddress:omitempty"`
	PublicIpAddressId   *int                      `json:"publicIpAddressId:omitempty"`
	PublicIpv6Address   *Network_Subnet_IpAddress `json:"publicIpv6Address:omitempty"`
	PublicIpv6AddressId *int                      `json:"publicIpv6AddressId:omitempty"`
	PublicVlan          *Network_Vlan             `json:"publicVlan:omitempty"`
	PublicVlanId        *int                      `json:"publicVlanId:omitempty"`
	Status              *Network_Gateway_Status   `json:"status:omitempty"`
	StatusId            *int                      `json:"statusId:omitempty"`
}

type Network_Gateway_Member struct {
	Entity

	Hardware         *Hardware        `json:"hardware:omitempty"`
	HardwareId       *int             `json:"hardwareId:omitempty"`
	Id               *int             `json:"id:omitempty"`
	NetworkGateway   *Network_Gateway `json:"networkGateway:omitempty"`
	NetworkGatewayId *int             `json:"networkGatewayId:omitempty"`
	Priority         *int             `json:"priority:omitempty"`
}

type Network_Gateway_Status struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Network_Gateway_Vlan struct {
	Entity

	BypassFlag       *bool            `json:"bypassFlag:omitempty"`
	Id               *int             `json:"id:omitempty"`
	NetworkGateway   *Network_Gateway `json:"networkGateway:omitempty"`
	NetworkGatewayId *int             `json:"networkGatewayId:omitempty"`
	NetworkVlan      *Network_Vlan    `json:"networkVlan:omitempty"`
	NetworkVlanId    *int             `json:"networkVlanId:omitempty"`
}

type Network_LoadBalancer_Global_Account struct {
	Entity

	Account                     *Account                           `json:"account:omitempty"`
	AllowedNumberOfHosts        *int                               `json:"allowedNumberOfHosts:omitempty"`
	AverageConnectionsPerSecond *float64                           `json:"averageConnectionsPerSecond:omitempty"`
	BillingItem                 *Billing_Item                      `json:"billingItem:omitempty"`
	ConnectionsPerSecond        *int                               `json:"connectionsPerSecond:omitempty"`
	FallbackIp                  *string                            `json:"fallbackIp:omitempty"`
	HostCount                   *uint                              `json:"hostCount:omitempty"`
	Hostname                    *string                            `json:"hostname:omitempty"`
	Hosts                       []Network_LoadBalancer_Global_Host `json:"hosts:omitempty"`
	Id                          *int                               `json:"id:omitempty"`
	LoadBalanceType             *Network_LoadBalancer_Global_Type  `json:"loadBalanceType:omitempty"`
	LoadBalanceTypeId           *int                               `json:"loadBalanceTypeId:omitempty"`
	ManagedResourceFlag         *bool                              `json:"managedResourceFlag:omitempty"`
	Notes                       *string                            `json:"notes:omitempty"`
}

type Network_LoadBalancer_Global_Host struct {
	Entity

	DestinationIp       *string                              `json:"destinationIp:omitempty"`
	DestinationPort     *int                                 `json:"destinationPort:omitempty"`
	Enabled             *int                                 `json:"enabled:omitempty"`
	HealthCheck         *string                              `json:"healthCheck:omitempty"`
	Hits                *float64                             `json:"hits:omitempty"`
	Id                  *int                                 `json:"id:omitempty"`
	LoadBalanceOrder    *int                                 `json:"loadBalanceOrder:omitempty"`
	LoadBalancerAccount *Network_LoadBalancer_Global_Account `json:"loadBalancerAccount:omitempty"`
	Location            *string                              `json:"location:omitempty"`
	Status              *string                              `json:"status:omitempty"`
	Weight              *int                                 `json:"weight:omitempty"`
}

type Network_LoadBalancer_Global_Type struct {
	Entity

	Id   *int    `json:"id:omitempty"`
	Name *string `json:"name:omitempty"`
}

type Network_LoadBalancer_Service struct {
	Entity

	ConnectionLimit      *int                                   `json:"connectionLimit:omitempty"`
	CreateDate           *time.Time                             `json:"createDate:omitempty"`
	DestinationIpAddress *string                                `json:"destinationIpAddress:omitempty"`
	DestinationPort      *int                                   `json:"destinationPort:omitempty"`
	Enabled              *bool                                  `json:"enabled:omitempty"`
	HealthCheck          *string                                `json:"healthCheck:omitempty"`
	HealthCheckURL       *string                                `json:"healthCheckURL:omitempty"`
	HealthResponse       *string                                `json:"healthResponse:omitempty"`
	Id                   *int                                   `json:"id:omitempty"`
	ModifyDate           *time.Time                             `json:"modifyDate:omitempty"`
	Name                 *string                                `json:"name:omitempty"`
	Notes                *string                                `json:"notes:omitempty"`
	PeakConnections      *int                                   `json:"peakConnections:omitempty"`
	SourcePort           *int                                   `json:"sourcePort:omitempty"`
	Type                 *string                                `json:"type:omitempty"`
	Vip                  *Network_LoadBalancer_VirtualIpAddress `json:"vip:omitempty"`
	VipId                *int                                   `json:"vipId:omitempty"`
	Weight               *int                                   `json:"weight:omitempty"`
}

type Network_LoadBalancer_VirtualIpAddress struct {
	Entity

	Account                     *Account                       `json:"account:omitempty"`
	BillingItem                 *Billing_Item                  `json:"billingItem:omitempty"`
	ConnectionLimit             *int                           `json:"connectionLimit:omitempty"`
	CustomerManagedFlag         *int                           `json:"customerManagedFlag:omitempty"`
	Id                          *int                           `json:"id:omitempty"`
	LoadBalancingMethod         *string                        `json:"loadBalancingMethod:omitempty"`
	LoadBalancingMethodFullName *string                        `json:"loadBalancingMethodFullName:omitempty"`
	ManagedResourceFlag         *bool                          `json:"managedResourceFlag:omitempty"`
	ModifyDate                  *time.Time                     `json:"modifyDate:omitempty"`
	Name                        *string                        `json:"name:omitempty"`
	Notes                       *string                        `json:"notes:omitempty"`
	SecurityCertificateId       *int                           `json:"securityCertificateId:omitempty"`
	ServiceCount                *uint                          `json:"serviceCount:omitempty"`
	Services                    []Network_LoadBalancer_Service `json:"services:omitempty"`
	SourcePort                  *int                           `json:"sourcePort:omitempty"`
	Type                        *string                        `json:"type:omitempty"`
	VirtualIpAddress            *string                        `json:"virtualIpAddress:omitempty"`
}

type Network_Logging_Syslog struct {
	Entity

	CreateDate           *time.Time `json:"createDate:omitempty"`
	DestinationIpAddress *string    `json:"destinationIpAddress:omitempty"`
	DestinationPort      *int       `json:"destinationPort:omitempty"`
	EventType            *string    `json:"eventType:omitempty"`
	Message              *string    `json:"message:omitempty"`
	Protocol             *string    `json:"protocol:omitempty"`
	SourceIpAddress      *string    `json:"sourceIpAddress:omitempty"`
	SourcePort           *int       `json:"sourcePort:omitempty"`
	TotalEvents          *int       `json:"totalEvents:omitempty"`
}

type Network_Media_Transcode_Account struct {
	Entity

	Account           *Account                      `json:"account:omitempty"`
	AccountId         *int                          `json:"accountId:omitempty"`
	CreateDate        *time.Time                    `json:"createDate:omitempty"`
	Id                *int                          `json:"id:omitempty"`
	ModifyDate        *time.Time                    `json:"modifyDate:omitempty"`
	TranscodeJobCount *uint                         `json:"transcodeJobCount:omitempty"`
	TranscodeJobs     []Network_Media_Transcode_Job `json:"transcodeJobs:omitempty"`
}

type Network_Media_Transcode_Job struct {
	Entity

	AutoDeleteDuration  *int                                             `json:"autoDeleteDuration:omitempty"`
	ByteIn              *int                                             `json:"byteIn:omitempty"`
	CreateDate          *time.Time                                       `json:"createDate:omitempty"`
	History             []Network_Media_Transcode_Job_History            `json:"history:omitempty"`
	HistoryCount        *uint                                            `json:"historyCount:omitempty"`
	Id                  *int                                             `json:"id:omitempty"`
	InputFile           *string                                          `json:"inputFile:omitempty"`
	ModifyDate          *time.Time                                       `json:"modifyDate:omitempty"`
	Name                *string                                          `json:"name:omitempty"`
	OutputFile          *string                                          `json:"outputFile:omitempty"`
	TranscodeAccount    *Network_Media_Transcode_Account                 `json:"transcodeAccount:omitempty"`
	TranscodeAccountId  *int                                             `json:"transcodeAccountId:omitempty"`
	TranscodeJobGuid    *string                                          `json:"transcodeJobGuid:omitempty"`
	TranscodePresetGuid *string                                          `json:"transcodePresetGuid:omitempty"`
	TranscodePresetName *string                                          `json:"transcodePresetName:omitempty"`
	TranscodeStatus     *Network_Media_Transcode_Job_Status              `json:"transcodeStatus:omitempty"`
	TranscodeStatusId   *int                                             `json:"transcodeStatusId:omitempty"`
	TranscodeStatusName *string                                          `json:"transcodeStatusName:omitempty"`
	User                *User_Customer                                   `json:"user:omitempty"`
	UserId              *int                                             `json:"userId:omitempty"`
	Watermark           *Container_Network_Media_Transcode_Job_Watermark `json:"watermark:omitempty"`
}

type Network_Media_Transcode_Job_History struct {
	Entity

	CreateDate          *time.Time `json:"createDate:omitempty"`
	PublicNotes         *string    `json:"publicNotes:omitempty"`
	TranscodeJobId      *int       `json:"transcodeJobId:omitempty"`
	TranscodeStatusName *string    `json:"transcodeStatusName:omitempty"`
}

type Network_Media_Transcode_Job_Status struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Network_Message_Delivery struct {
	Entity

	Account     *Account                         `json:"account:omitempty"`
	AccountId   *int                             `json:"accountId:omitempty"`
	BillingItem *Billing_Item                    `json:"billingItem:omitempty"`
	CreateDate  *time.Time                       `json:"createDate:omitempty"`
	Id          *int                             `json:"id:omitempty"`
	ModifyDate  *time.Time                       `json:"modifyDate:omitempty"`
	Password    *string                          `json:"password:omitempty"`
	Type        *Network_Message_Delivery_Type   `json:"type:omitempty"`
	TypeId      *int                             `json:"typeId:omitempty"`
	Username    *string                          `json:"username:omitempty"`
	Vendor      *Network_Message_Delivery_Vendor `json:"vendor:omitempty"`
	VendorId    *int                             `json:"vendorId:omitempty"`
}

type Network_Message_Delivery_Attribute struct {
	Entity

	NetworkMessageDelivery *Network_Message_Delivery `json:"networkMessageDelivery:omitempty"`
	Value                  *string                   `json:"value:omitempty"`
}

type Network_Message_Delivery_Email_Sendgrid struct {
	Network_Message_Delivery

	EmailAddress *string `json:"emailAddress:omitempty"`
	SmtpAccess   *string `json:"smtpAccess:omitempty"`
}

type Network_Message_Delivery_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Network_Message_Delivery_Vendor struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Network_Message_Queue struct {
	Entity

	Account              *Account                      `json:"account:omitempty"`
	AccountId            *int                          `json:"accountId:omitempty"`
	BillingItem          *Billing_Item                 `json:"billingItem:omitempty"`
	CreateDate           *time.Time                    `json:"createDate:omitempty"`
	Id                   *int                          `json:"id:omitempty"`
	MessageQueueStatusId *int                          `json:"messageQueueStatusId:omitempty"`
	Name                 *string                       `json:"name:omitempty"`
	NodeCount            *uint                         `json:"nodeCount:omitempty"`
	Nodes                []Network_Message_Queue_Node  `json:"nodes:omitempty"`
	Notes                *string                       `json:"notes:omitempty"`
	Status               *Network_Message_Queue_Status `json:"status:omitempty"`
}

type Network_Message_Queue_Node struct {
	Entity

	AccountName          *string                   `json:"accountName:omitempty"`
	Id                   *int                      `json:"id:omitempty"`
	MessageQueue         *Network_Message_Queue    `json:"messageQueue:omitempty"`
	MessageQueueId       *int                      `json:"messageQueueId:omitempty"`
	MetricTrackingObject *Metric_Tracking_Object   `json:"metricTrackingObject:omitempty"`
	Name                 *string                   `json:"name:omitempty"`
	Notes                *string                   `json:"notes:omitempty"`
	ServiceResource      *Network_Service_Resource `json:"serviceResource:omitempty"`
}

type Network_Message_Queue_Status struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Network_Monitor struct {
	Entity
}

type Network_Monitor_Version1_Incident struct {
	Entity

	Status *string `json:"status:omitempty"`
}

type Network_Monitor_Version1_Query_Host struct {
	Entity

	Arg1Value        *string                                      `json:"arg1Value:omitempty"`
	GuestId          *int                                         `json:"guestId:omitempty"`
	Hardware         *Hardware                                    `json:"hardware:omitempty"`
	HardwareId       *int                                         `json:"hardwareId:omitempty"`
	HostId           *int                                         `json:"hostId:omitempty"`
	Id               *int                                         `json:"id:omitempty"`
	IpAddress        *string                                      `json:"ipAddress:omitempty"`
	LastResult       *Network_Monitor_Version1_Query_Result       `json:"lastResult:omitempty"`
	QueryType        *Network_Monitor_Version1_Query_Type         `json:"queryType:omitempty"`
	QueryTypeId      *int                                         `json:"queryTypeId:omitempty"`
	ResponseAction   *Network_Monitor_Version1_Query_ResponseType `json:"responseAction:omitempty"`
	ResponseActionId *int                                         `json:"responseActionId:omitempty"`
	Status           *string                                      `json:"status:omitempty"`
	WaitCycles       *int                                         `json:"waitCycles:omitempty"`
}

type Network_Monitor_Version1_Query_Host_Stratum struct {
	Entity

	Hardware      *Hardware `json:"hardware:omitempty"`
	MonitorLevel  *int      `json:"monitorLevel:omitempty"`
	ResponseLevel *int      `json:"responseLevel:omitempty"`
}

type Network_Monitor_Version1_Query_ResponseType struct {
	Entity

	ActionDescription *string `json:"actionDescription:omitempty"`
	Id                *int    `json:"id:omitempty"`
	Level             *int    `json:"level:omitempty"`
}

type Network_Monitor_Version1_Query_Result struct {
	Entity

	FinishTime     *time.Time                           `json:"finishTime:omitempty"`
	QueryHost      *Network_Monitor_Version1_Query_Host `json:"queryHost:omitempty"`
	ResponseStatus *int                                 `json:"responseStatus:omitempty"`
	ResponseTime   *float64                             `json:"responseTime:omitempty"`
}

type Network_Monitor_Version1_Query_Type struct {
	Entity

	ArgumentDescription *string `json:"argumentDescription:omitempty"`
	Description         *string `json:"description:omitempty"`
	Id                  *int    `json:"id:omitempty"`
	MonitorLevel        *int    `json:"monitorLevel:omitempty"`
	Name                *string `json:"name:omitempty"`
}

type Network_Pod struct {
	Entity

	BackendRouterId    *int     `json:"backendRouterId:omitempty"`
	BackendRouterName  *string  `json:"backendRouterName:omitempty"`
	Capabilities       []string `json:"capabilities:omitempty"`
	DatacenterLongName *string  `json:"datacenterLongName:omitempty"`
	DatacenterName     *string  `json:"datacenterName:omitempty"`
	FrontendRouterId   *int     `json:"frontendRouterId:omitempty"`
	FrontendRouterName *string  `json:"frontendRouterName:omitempty"`
	Name               *string  `json:"name:omitempty"`
}

type Network_Protection_Address struct {
	Entity

	Account              *Account                            `json:"account:omitempty"`
	DepartmentId         *int                                `json:"departmentId:omitempty"`
	IpAddress            *string                             `json:"ipAddress:omitempty"`
	Location             *Location                           `json:"location:omitempty"`
	ManagementMethodType *string                             `json:"managementMethodType:omitempty"`
	ModifiedUser         *User_Employee                      `json:"modifiedUser:omitempty"`
	PrimaryRouter        *Hardware_Router                    `json:"primaryRouter:omitempty"`
	ServiceProvider      *Service_Provider                   `json:"serviceProvider:omitempty"`
	Subnet               *Network_Subnet                     `json:"subnet:omitempty"`
	SubnetIpAddress      *Network_Subnet_IpAddress           `json:"subnetIpAddress:omitempty"`
	TerminatedUser       *User_Employee                      `json:"terminatedUser:omitempty"`
	Ticket               *Ticket                             `json:"ticket:omitempty"`
	TransactionCount     *uint                               `json:"transactionCount:omitempty"`
	Transactions         []Provisioning_Version1_Transaction `json:"transactions:omitempty"`
	UserDepartment       *User_Employee_Department           `json:"userDepartment:omitempty"`
	UserRecord           *User_Employee                      `json:"userRecord:omitempty"`
}

type Network_Regional_Internet_Registry struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Network_Security_Scanner_Request struct {
	Entity

	Account            *Account                                 `json:"account:omitempty"`
	AccountId          *int                                     `json:"accountId:omitempty"`
	CreateDate         *time.Time                               `json:"createDate:omitempty"`
	Guest              *Virtual_Guest                           `json:"guest:omitempty"`
	GuestId            *int                                     `json:"guestId:omitempty"`
	Hardware           *Hardware                                `json:"hardware:omitempty"`
	HardwareId         *int                                     `json:"hardwareId:omitempty"`
	HostId             *int                                     `json:"hostId:omitempty"`
	Id                 *int                                     `json:"id:omitempty"`
	IpAddress          *string                                  `json:"ipAddress:omitempty"`
	ModifyDate         *time.Time                               `json:"modifyDate:omitempty"`
	RequestorOwnedFlag *bool                                    `json:"requestorOwnedFlag:omitempty"`
	Status             *Network_Security_Scanner_Request_Status `json:"status:omitempty"`
	StatusId           *int                                     `json:"statusId:omitempty"`
}

type Network_Security_Scanner_Request_Status struct {
	Entity

	Id   *int    `json:"id:omitempty"`
	Name *string `json:"name:omitempty"`
}

type Network_Service_Health struct {
	Entity

	CreateDate *time.Time                     `json:"createDate:omitempty"`
	Location   *Location                      `json:"location:omitempty"`
	LocationId *int                           `json:"locationId:omitempty"`
	ModifyDate *time.Time                     `json:"modifyDate:omitempty"`
	Status     *Network_Service_Health_Status `json:"status:omitempty"`
	StatusId   *int                           `json:"statusId:omitempty"`
}

type Network_Service_Health_Status struct {
	Entity

	Name *string `json:"name:omitempty"`
}

type Network_Service_Resource struct {
	Entity

	ApiHost           *string                              `json:"apiHost:omitempty"`
	ApiPassword       *string                              `json:"apiPassword:omitempty"`
	ApiPath           *string                              `json:"apiPath:omitempty"`
	ApiPort           *string                              `json:"apiPort:omitempty"`
	ApiProtocol       *string                              `json:"apiProtocol:omitempty"`
	ApiUsername       *string                              `json:"apiUsername:omitempty"`
	ApiVersion        *string                              `json:"apiVersion:omitempty"`
	AttributeCount    *uint                                `json:"attributeCount:omitempty"`
	Attributes        []Network_Service_Resource_Attribute `json:"attributes:omitempty"`
	BackendIpAddress  *string                              `json:"backendIpAddress:omitempty"`
	Datacenter        *Location                            `json:"datacenter:omitempty"`
	FrontendIpAddress *string                              `json:"frontendIpAddress:omitempty"`
	Id                *int                                 `json:"id:omitempty"`
	Name              *string                              `json:"name:omitempty"`
	NetworkDevice     *Hardware                            `json:"networkDevice:omitempty"`
	SshUsername       *string                              `json:"sshUsername:omitempty"`
	Type              *Network_Service_Resource_Type       `json:"type:omitempty"`
}

type Network_Service_Resource_Attribute struct {
	Entity

	AttributeType   *Network_Service_Resource_Attribute_Type `json:"attributeType:omitempty"`
	ServiceResource *Network_Service_Resource                `json:"serviceResource:omitempty"`
	Value           *string                                  `json:"value:omitempty"`
}

type Network_Service_Resource_Attribute_Type struct {
	Entity

	Keyname *string `json:"keyname:omitempty"`
}

type Network_Service_Resource_Hub struct {
	Network_Service_Resource
}

type Network_Service_Resource_Hub_Swift struct {
	Network_Service_Resource_Hub
}

type Network_Service_Resource_MonitoringHub struct {
	Network_Service_Resource

	AdnServicesIp        *string `json:"adnServicesIp:omitempty"`
	HubAddress           *string `json:"hubAddress:omitempty"`
	HubConnectionTimeout *string `json:"hubConnectionTimeout:omitempty"`
	RobotsCount          *string `json:"robotsCount:omitempty"`
	RobotsMax            *string `json:"robotsMax:omitempty"`
}

type Network_Service_Resource_NimsoftLandingHub struct {
	Network_Service_Resource_MonitoringHub
}

type Network_Service_Resource_Type struct {
	Entity

	ServiceResourceCount *uint                      `json:"serviceResourceCount:omitempty"`
	ServiceResources     []Network_Service_Resource `json:"serviceResources:omitempty"`
	Type                 *string                    `json:"type:omitempty"`
}

type Network_Service_Vpn_Overrides struct {
	Entity

	Id       *int            `json:"id:omitempty"`
	Subnet   *Network_Subnet `json:"subnet:omitempty"`
	SubnetId *int            `json:"subnetId:omitempty"`
	User     *User_Customer  `json:"user:omitempty"`
	UserId   *int            `json:"userId:omitempty"`
}

type Network_Storage struct {
	Entity

	Account                             *Account                            `json:"account:omitempty"`
	AccountId                           *int                                `json:"accountId:omitempty"`
	AccountPassword                     *Account_Password                   `json:"accountPassword:omitempty"`
	ActiveTransactionCount              *uint                               `json:"activeTransactionCount:omitempty"`
	ActiveTransactions                  []Provisioning_Version1_Transaction `json:"activeTransactions:omitempty"`
	AllowedHardware                     []Hardware                          `json:"allowedHardware:omitempty"`
	AllowedHardwareCount                *uint                               `json:"allowedHardwareCount:omitempty"`
	AllowedIpAddressCount               *uint                               `json:"allowedIpAddressCount:omitempty"`
	AllowedIpAddresses                  []Network_Subnet_IpAddress          `json:"allowedIpAddresses:omitempty"`
	AllowedReplicationHardware          []Hardware                          `json:"allowedReplicationHardware:omitempty"`
	AllowedReplicationHardwareCount     *uint                               `json:"allowedReplicationHardwareCount:omitempty"`
	AllowedReplicationIpAddressCount    *uint                               `json:"allowedReplicationIpAddressCount:omitempty"`
	AllowedReplicationIpAddresses       []Network_Subnet_IpAddress          `json:"allowedReplicationIpAddresses:omitempty"`
	AllowedReplicationSubnetCount       *uint                               `json:"allowedReplicationSubnetCount:omitempty"`
	AllowedReplicationSubnets           []Network_Subnet                    `json:"allowedReplicationSubnets:omitempty"`
	AllowedReplicationVirtualGuestCount *uint                               `json:"allowedReplicationVirtualGuestCount:omitempty"`
	AllowedReplicationVirtualGuests     []Virtual_Guest                     `json:"allowedReplicationVirtualGuests:omitempty"`
	AllowedSubnetCount                  *uint                               `json:"allowedSubnetCount:omitempty"`
	AllowedSubnets                      []Network_Subnet                    `json:"allowedSubnets:omitempty"`
	AllowedVirtualGuestCount            *uint                               `json:"allowedVirtualGuestCount:omitempty"`
	AllowedVirtualGuests                []Virtual_Guest                     `json:"allowedVirtualGuests:omitempty"`
	BillingItem                         *Billing_Item                       `json:"billingItem:omitempty"`
	BillingItemCategory                 *Product_Item_Category              `json:"billingItemCategory:omitempty"`
	BytesUsed                           *string                             `json:"bytesUsed:omitempty"`
	CapacityGb                          *int                                `json:"capacityGb:omitempty"`
	CreateDate                          *time.Time                          `json:"createDate:omitempty"`
	CreationScheduleId                  *string                             `json:"creationScheduleId:omitempty"`
	CredentialCount                     *uint                               `json:"credentialCount:omitempty"`
	Credentials                         []Network_Storage_Credential        `json:"credentials:omitempty"`
	DailySchedule                       *Network_Storage_Schedule           `json:"dailySchedule:omitempty"`
	EventCount                          *uint                               `json:"eventCount:omitempty"`
	Events                              []Network_Storage_Event             `json:"events:omitempty"`
	GuestId                             *int                                `json:"guestId:omitempty"`
	Hardware                            *Hardware                           `json:"hardware:omitempty"`
	HardwareId                          *int                                `json:"hardwareId:omitempty"`
	HostId                              *int                                `json:"hostId:omitempty"`
	HourlySchedule                      *Network_Storage_Schedule           `json:"hourlySchedule:omitempty"`
	Id                                  *int                                `json:"id:omitempty"`
	Iops                                *string                             `json:"iops:omitempty"`
	IscsiLunCount                       *uint                               `json:"iscsiLunCount:omitempty"`
	IscsiLuns                           []Network_Storage                   `json:"iscsiLuns:omitempty"`
	LunId                               *string                             `json:"lunId:omitempty"`
	ManualSnapshotCount                 *uint                               `json:"manualSnapshotCount:omitempty"`
	ManualSnapshots                     []Network_Storage                   `json:"manualSnapshots:omitempty"`
	MetricTrackingObject                *Metric_Tracking_Object             `json:"metricTrackingObject:omitempty"`
	MountableFlag                       *string                             `json:"mountableFlag:omitempty"`
	NasType                             *string                             `json:"nasType:omitempty"`
	Notes                               *string                             `json:"notes:omitempty"`
	NotificationSubscriberCount         *uint                               `json:"notificationSubscriberCount:omitempty"`
	NotificationSubscribers             []Notification_User_Subscriber      `json:"notificationSubscribers:omitempty"`
	OsType                              *Network_Storage_Iscsi_OS_Type      `json:"osType:omitempty"`
	OsTypeId                            *string                             `json:"osTypeId:omitempty"`
	ParentPartnershipCount              *uint                               `json:"parentPartnershipCount:omitempty"`
	ParentPartnerships                  []Network_Storage_Partnership       `json:"parentPartnerships:omitempty"`
	ParentVolume                        *Network_Storage                    `json:"parentVolume:omitempty"`
	PartnershipCount                    *uint                               `json:"partnershipCount:omitempty"`
	Partnerships                        []Network_Storage_Partnership       `json:"partnerships:omitempty"`
	Password                            *string                             `json:"password:omitempty"`
	PermissionsGroupCount               *uint                               `json:"permissionsGroupCount:omitempty"`
	PermissionsGroups                   []Network_Storage_Group             `json:"permissionsGroups:omitempty"`
	Properties                          []Network_Storage_Property          `json:"properties:omitempty"`
	PropertyCount                       *uint                               `json:"propertyCount:omitempty"`
	ReplicatingLunCount                 *uint                               `json:"replicatingLunCount:omitempty"`
	ReplicatingLuns                     []Network_Storage                   `json:"replicatingLuns:omitempty"`
	ReplicatingVolume                   *Network_Storage                    `json:"replicatingVolume:omitempty"`
	ReplicationEventCount               *uint                               `json:"replicationEventCount:omitempty"`
	ReplicationEvents                   []Network_Storage_Event             `json:"replicationEvents:omitempty"`
	ReplicationPartnerCount             *uint                               `json:"replicationPartnerCount:omitempty"`
	ReplicationPartners                 []Network_Storage                   `json:"replicationPartners:omitempty"`
	ReplicationSchedule                 *Network_Storage_Schedule           `json:"replicationSchedule:omitempty"`
	ReplicationStatus                   *string                             `json:"replicationStatus:omitempty"`
	ScheduleCount                       *uint                               `json:"scheduleCount:omitempty"`
	Schedules                           []Network_Storage_Schedule          `json:"schedules:omitempty"`
	ServiceProviderId                   *int                                `json:"serviceProviderId:omitempty"`
	ServiceResource                     *Network_Service_Resource           `json:"serviceResource:omitempty"`
	ServiceResourceBackendIpAddress     *string                             `json:"serviceResourceBackendIpAddress:omitempty"`
	ServiceResourceName                 *string                             `json:"serviceResourceName:omitempty"`
	SnapshotCapacityGb                  *string                             `json:"snapshotCapacityGb:omitempty"`
	SnapshotCount                       *uint                               `json:"snapshotCount:omitempty"`
	SnapshotCreationTimestamp           *string                             `json:"snapshotCreationTimestamp:omitempty"`
	SnapshotDeletionThresholdPercentage *string                             `json:"snapshotDeletionThresholdPercentage:omitempty"`
	SnapshotSizeBytes                   *string                             `json:"snapshotSizeBytes:omitempty"`
	SnapshotSpaceAvailable              *string                             `json:"snapshotSpaceAvailable:omitempty"`
	Snapshots                           []Network_Storage                   `json:"snapshots:omitempty"`
	StorageGroupCount                   *uint                               `json:"storageGroupCount:omitempty"`
	StorageGroups                       []Network_Storage_Group             `json:"storageGroups:omitempty"`
	StorageTierLevel                    *string                             `json:"storageTierLevel:omitempty"`
	StorageType                         *Network_Storage_Type               `json:"storageType:omitempty"`
	StorageTypeId                       *string                             `json:"storageTypeId:omitempty"`
	TotalBytesUsed                      *string                             `json:"totalBytesUsed:omitempty"`
	TotalScheduleSnapshotRetentionCount *uint                               `json:"totalScheduleSnapshotRetentionCount:omitempty"`
	UpgradableFlag                      *bool                               `json:"upgradableFlag:omitempty"`
	UsageNotification                   *Notification                       `json:"usageNotification:omitempty"`
	Username                            *string                             `json:"username:omitempty"`
	VendorName                          *string                             `json:"vendorName:omitempty"`
	VirtualGuest                        *Virtual_Guest                      `json:"virtualGuest:omitempty"`
	VolumeHistory                       []Network_Storage_History           `json:"volumeHistory:omitempty"`
	VolumeHistoryCount                  *uint                               `json:"volumeHistoryCount:omitempty"`
	VolumeStatus                        *string                             `json:"volumeStatus:omitempty"`
	WebccAccount                        *Account_Password                   `json:"webccAccount:omitempty"`
	WeeklySchedule                      *Network_Storage_Schedule           `json:"weeklySchedule:omitempty"`
}

type Network_Storage_Allowed_Host struct {
	Entity

	AssignedGroupCount             *uint                       `json:"assignedGroupCount:omitempty"`
	AssignedGroups                 []Network_Storage_Group     `json:"assignedGroups:omitempty"`
	AssignedReplicationVolumeCount *uint                       `json:"assignedReplicationVolumeCount:omitempty"`
	AssignedReplicationVolumes     []Network_Storage           `json:"assignedReplicationVolumes:omitempty"`
	AssignedVolumeCount            *uint                       `json:"assignedVolumeCount:omitempty"`
	AssignedVolumes                []Network_Storage           `json:"assignedVolumes:omitempty"`
	Credential                     *Network_Storage_Credential `json:"credential:omitempty"`
	CredentialId                   *int                        `json:"credentialId:omitempty"`
	Id                             *int                        `json:"id:omitempty"`
	Name                           *string                     `json:"name:omitempty"`
	ResourceTableId                *int                        `json:"resourceTableId:omitempty"`
	ResourceTableName              *string                     `json:"resourceTableName:omitempty"`
}

type Network_Storage_Allowed_Host_Hardware struct {
	Network_Storage_Allowed_Host

	Resource *Hardware `json:"resource:omitempty"`
}

type Network_Storage_Allowed_Host_IpAddress struct {
	Network_Storage_Allowed_Host

	Resource *Network_Subnet_IpAddress `json:"resource:omitempty"`
}

type Network_Storage_Allowed_Host_Subnet struct {
	Network_Storage_Allowed_Host

	Resource *Network_Subnet `json:"resource:omitempty"`
}

type Network_Storage_Allowed_Host_VirtualGuest struct {
	Network_Storage_Allowed_Host

	Resource *Virtual_Guest `json:"resource:omitempty"`
}

type Network_Storage_Backup struct {
	Network_Storage

	CurrentCyclePeakUsage  *uint `json:"currentCyclePeakUsage:omitempty"`
	PreviousCyclePeakUsage *uint `json:"previousCyclePeakUsage:omitempty"`
}

type Network_Storage_Backup_Evault struct {
	Network_Storage_Backup
}

type Network_Storage_Backup_Evault_Version6 struct {
	Network_Storage_Backup_Evault

	AgentStatusCount       *uint                                                `json:"agentStatusCount:omitempty"`
	AgentStatuses          []Container_Network_Storage_Evault_WebCc_AgentStatus `json:"agentStatuses:omitempty"`
	BackupJobDetailCount   *uint                                                `json:"backupJobDetailCount:omitempty"`
	BackupJobDetails       []Container_Network_Storage_Evault_WebCc_JobDetails  `json:"backupJobDetails:omitempty"`
	PluginBillingItemCount *uint                                                `json:"pluginBillingItemCount:omitempty"`
	PluginBillingItems     []Billing_Item                                       `json:"pluginBillingItems:omitempty"`
	RestoreJobDetailCount  *uint                                                `json:"restoreJobDetailCount:omitempty"`
	RestoreJobDetails      []Container_Network_Storage_Evault_WebCc_JobDetails  `json:"restoreJobDetails:omitempty"`
	SoftwareComponent      *Software_Component                                  `json:"softwareComponent:omitempty"`
	TaskCount              *uint                                                `json:"taskCount:omitempty"`
	Tasks                  []Container_Network_Storage_Evault_Vault_Task        `json:"tasks:omitempty"`
}

type Network_Storage_Credential struct {
	Entity

	Account                    *Account                         `json:"account:omitempty"`
	AccountId                  *string                          `json:"accountId:omitempty"`
	CreateDate                 *time.Time                       `json:"createDate:omitempty"`
	Id                         *int                             `json:"id:omitempty"`
	ModifyDate                 *time.Time                       `json:"modifyDate:omitempty"`
	NasCredentialTypeId        *int                             `json:"nasCredentialTypeId:omitempty"`
	NetworkStorageAllowedHosts *Network_Storage_Allowed_Host    `json:"networkStorageAllowedHosts:omitempty"`
	Password                   *string                          `json:"password:omitempty"`
	Type                       *Network_Storage_Credential_Type `json:"type:omitempty"`
	Username                   *string                          `json:"username:omitempty"`
	VolumeCount                *uint                            `json:"volumeCount:omitempty"`
	Volumes                    []Network_Storage                `json:"volumes:omitempty"`
}

type Network_Storage_Credential_Type struct {
	Entity

	CreateDate  *time.Time `json:"createDate:omitempty"`
	Description *string    `json:"description:omitempty"`
	KeyName     *string    `json:"keyName:omitempty"`
	ModifyDate  *time.Time `json:"modifyDate:omitempty"`
	Name        *string    `json:"name:omitempty"`
}

type Network_Storage_Daily_Usage struct {
	Entity

	BytesUsed          *uint            `json:"bytesUsed:omitempty"`
	CdnHttpBandwidth   *uint            `json:"cdnHttpBandwidth:omitempty"`
	CreateDate         *time.Time       `json:"createDate:omitempty"`
	NasVolume          *Network_Storage `json:"nasVolume:omitempty"`
	NasVolumeId        *int             `json:"nasVolumeId:omitempty"`
	PublicBandwidthOut *uint            `json:"publicBandwidthOut:omitempty"`
}

type Network_Storage_Event struct {
	Entity

	CreateDate *time.Time                `json:"createDate:omitempty"`
	Message    *string                   `json:"message:omitempty"`
	Schedule   *Network_Storage_Schedule `json:"schedule:omitempty"`
	ScheduleId *int                      `json:"scheduleId:omitempty"`
	TypeId     *int                      `json:"typeId:omitempty"`
	Volume     *Network_Storage          `json:"volume:omitempty"`
	VolumeId   *int                      `json:"volumeId:omitempty"`
}

type Network_Storage_Group struct {
	Entity

	Account             *Account                       `json:"account:omitempty"`
	AccountId           *int                           `json:"accountId:omitempty"`
	Alias               *string                        `json:"alias:omitempty"`
	AllowedHostCount    *uint                          `json:"allowedHostCount:omitempty"`
	AllowedHosts        []Network_Storage_Allowed_Host `json:"allowedHosts:omitempty"`
	AttachedVolumeCount *uint                          `json:"attachedVolumeCount:omitempty"`
	AttachedVolumes     []Network_Storage              `json:"attachedVolumes:omitempty"`
	CreateDate          *time.Time                     `json:"createDate:omitempty"`
	GroupType           *Network_Storage_Group_Type    `json:"groupType:omitempty"`
	GroupTypeId         *int                           `json:"groupTypeId:omitempty"`
	Id                  *int                           `json:"id:omitempty"`
	ModifyDate          *time.Time                     `json:"modifyDate:omitempty"`
	OsType              *Network_Storage_Iscsi_OS_Type `json:"osType:omitempty"`
	OsTypeId            *int                           `json:"osTypeId:omitempty"`
	ServiceResource     *Network_Service_Resource      `json:"serviceResource:omitempty"`
	ServiceResourceId   *int                           `json:"serviceResourceId:omitempty"`
}

type Network_Storage_Group_Iscsi struct {
	Network_Storage_Group
}

type Network_Storage_Group_Nfs struct {
	Network_Storage_Group
}

type Network_Storage_Group_Type struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Network_Storage_History struct {
	Entity

	Account    *Account         `json:"account:omitempty"`
	CreateDate *time.Time       `json:"createDate:omitempty"`
	NasVolume  *Network_Storage `json:"nasVolume:omitempty"`
	Notes      *string          `json:"notes:omitempty"`
	Password   *string          `json:"password:omitempty"`
	Username   *string          `json:"username:omitempty"`
}

type Network_Storage_Hub struct {
	Network_Storage

	BandwidthBillingItemCount *uint          `json:"bandwidthBillingItemCount:omitempty"`
	BandwidthBillingItems     []Billing_Item `json:"bandwidthBillingItems:omitempty"`
}

type Network_Storage_Hub_Cleversafe_Account struct {
	Entity

	Account              *Account                     `json:"account:omitempty"`
	AccountId            *int                         `json:"accountId:omitempty"`
	BillingItem          *Billing_Item                `json:"billingItem:omitempty"`
	CancelledBillingItem *Billing_Item                `json:"cancelledBillingItem:omitempty"`
	CredentialCount      *uint                        `json:"credentialCount:omitempty"`
	Credentials          []Network_Storage_Credential `json:"credentials:omitempty"`
	EventCount           *uint                        `json:"eventCount:omitempty"`
	Events               []Network_Storage_Event      `json:"events:omitempty"`
	Id                   *int                         `json:"id:omitempty"`
	MetricTrackingObject *Metric_Tracking_Object      `json:"metricTrackingObject:omitempty"`
	NasType              *string                      `json:"nasType:omitempty"`
	Notes                *string                      `json:"notes:omitempty"`
	StorageTypeId        *int                         `json:"storageTypeId:omitempty"`
	Username             *string                      `json:"username:omitempty"`
	Uuid                 *string                      `json:"uuid:omitempty"`
}

type Network_Storage_Hub_Swift struct {
	Network_Storage_Hub

	StorageNodeCount *uint                      `json:"storageNodeCount:omitempty"`
	StorageNodes     []Network_Service_Resource `json:"storageNodes:omitempty"`
}

type Network_Storage_Hub_Swift_Container struct {
	Network_Storage_Hub_Swift
}

type Network_Storage_Hub_Swift_Share struct {
	Entity
}

type Network_Storage_Hub_Swift_Version1 struct {
	Network_Storage_Hub_Swift
}

type Network_Storage_Iscsi struct {
	Network_Storage
}

type Network_Storage_Iscsi_EqualLogic_Version3 struct {
	Network_Storage_Iscsi
}

type Network_Storage_Iscsi_EqualLogic_Version3_Replicant struct {
	Network_Storage_Iscsi_EqualLogic_Version3

	FailbackInProgressFlag *bool   `json:"failbackInProgressFlag:omitempty"`
	VolumeName             *string `json:"volumeName:omitempty"`
}

type Network_Storage_Iscsi_EqualLogic_Version3_Snapshot struct {
	Network_Storage_Iscsi_EqualLogic_Version3

	CreationSchedule *Network_Storage_Schedule `json:"creationSchedule:omitempty"`
	VolumeName       *string                   `json:"volumeName:omitempty"`
}

type Network_Storage_Iscsi_OS_Type struct {
	Entity

	CreateDate  *time.Time `json:"createDate:omitempty"`
	Description *string    `json:"description:omitempty"`
	Id          *int       `json:"id:omitempty"`
	KeyName     *string    `json:"keyName:omitempty"`
	Name        *string    `json:"name:omitempty"`
}

type Network_Storage_Nas struct {
	Network_Storage

	RecentBytesUsed *Network_Storage_Daily_Usage `json:"recentBytesUsed:omitempty"`
}

type Network_Storage_OpenStack_Object struct {
	Network_Storage

	BandwidthBillingItemCount *uint          `json:"bandwidthBillingItemCount:omitempty"`
	BandwidthBillingItems     []Billing_Item `json:"bandwidthBillingItems:omitempty"`
}

type Network_Storage_Partnership struct {
	Entity

	CreateDate      *time.Time                        `json:"createDate:omitempty"`
	ModifyDate      *time.Time                        `json:"modifyDate:omitempty"`
	PartnerVolume   *Network_Storage                  `json:"partnerVolume:omitempty"`
	PartnerVolumeId *int                              `json:"partnerVolumeId:omitempty"`
	Type            *Network_Storage_Partnership_Type `json:"type:omitempty"`
	Volume          *Network_Storage                  `json:"volume:omitempty"`
	VolumeId        *int                              `json:"volumeId:omitempty"`
}

type Network_Storage_Partnership_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Keyname     *string `json:"keyname:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Network_Storage_Property struct {
	Entity

	CreateDate *time.Time                     `json:"createDate:omitempty"`
	ModifyDate *time.Time                     `json:"modifyDate:omitempty"`
	Type       *Network_Storage_Property_Type `json:"type:omitempty"`
	Value      *string                        `json:"value:omitempty"`
	Volume     *Network_Storage               `json:"volume:omitempty"`
	VolumeId   *int                           `json:"volumeId:omitempty"`
}

type Network_Storage_Property_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Keyname     *string `json:"keyname:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Network_Storage_Replicant struct {
	Network_Storage

	FailbackInProgressFlag *string `json:"failbackInProgressFlag:omitempty"`
	VolumeName             *string `json:"volumeName:omitempty"`
}

type Network_Storage_Schedule struct {
	Entity

	Active               *int                                `json:"active:omitempty"`
	CreateDate           *time.Time                          `json:"createDate:omitempty"`
	DayOfMonth           *string                             `json:"dayOfMonth:omitempty"`
	DayOfWeek            *string                             `json:"dayOfWeek:omitempty"`
	EventCount           *uint                               `json:"eventCount:omitempty"`
	Events               []Network_Storage_Event             `json:"events:omitempty"`
	Hour                 *string                             `json:"hour:omitempty"`
	Id                   *int                                `json:"id:omitempty"`
	Minute               *string                             `json:"minute:omitempty"`
	ModifyDate           *time.Time                          `json:"modifyDate:omitempty"`
	MonthOfYear          *string                             `json:"monthOfYear:omitempty"`
	Name                 *string                             `json:"name:omitempty"`
	Partnership          *Network_Storage_Partnership        `json:"partnership:omitempty"`
	PartnershipId        *int                                `json:"partnershipId:omitempty"`
	Properties           []Network_Storage_Schedule_Property `json:"properties:omitempty"`
	PropertyCount        *uint                               `json:"propertyCount:omitempty"`
	ReplicaSnapshotCount *uint                               `json:"replicaSnapshotCount:omitempty"`
	ReplicaSnapshots     []Network_Storage                   `json:"replicaSnapshots:omitempty"`
	RetentionCount       *string                             `json:"retentionCount:omitempty"`
	SnapshotCount        *uint                               `json:"snapshotCount:omitempty"`
	Snapshots            []Network_Storage                   `json:"snapshots:omitempty"`
	Type                 *Network_Storage_Schedule_Type      `json:"type:omitempty"`
	TypeId               *int                                `json:"typeId:omitempty"`
	Volume               *Network_Storage                    `json:"volume:omitempty"`
	VolumeId             *int                                `json:"volumeId:omitempty"`
}

type Network_Storage_Schedule_Property struct {
	Entity

	CreateDate *time.Time                              `json:"createDate:omitempty"`
	Id         *int                                    `json:"id:omitempty"`
	ModifyDate *time.Time                              `json:"modifyDate:omitempty"`
	Schedule   *Network_Storage_Schedule               `json:"schedule:omitempty"`
	Type       *Network_Storage_Schedule_Property_Type `json:"type:omitempty"`
	TypeId     *int                                    `json:"typeId:omitempty"`
	Value      *string                                 `json:"value:omitempty"`
}

type Network_Storage_Schedule_Property_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	Keyname     *string `json:"keyname:omitempty"`
	Name        *string `json:"name:omitempty"`
	NasType     *string `json:"nasType:omitempty"`
}

type Network_Storage_Schedule_Type struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	Keyname *string `json:"keyname:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Network_Storage_Snapshot struct {
	Network_Storage

	CreationSchedule *Network_Storage_Schedule `json:"creationSchedule:omitempty"`
	VolumeName       *string                   `json:"volumeName:omitempty"`
}

type Network_Storage_Type struct {
	Entity

	Description *string           `json:"description:omitempty"`
	Id          *int              `json:"id:omitempty"`
	KeyName     *string           `json:"keyName:omitempty"`
	VolumeCount *uint             `json:"volumeCount:omitempty"`
	Volumes     []Network_Storage `json:"volumes:omitempty"`
}

type Network_Subnet struct {
	Entity

	Account                           *Account                            `json:"account:omitempty"`
	ActiveRegistration                *Network_Subnet_Registration        `json:"activeRegistration:omitempty"`
	ActiveSwipTransaction             *Network_Subnet_Swip_Transaction    `json:"activeSwipTransaction:omitempty"`
	ActiveTransaction                 *Provisioning_Version1_Transaction  `json:"activeTransaction:omitempty"`
	AddressSpace                      *string                             `json:"addressSpace:omitempty"`
	AllowedHost                       *Network_Storage_Allowed_Host       `json:"allowedHost:omitempty"`
	AllowedNetworkStorage             []Network_Storage                   `json:"allowedNetworkStorage:omitempty"`
	AllowedNetworkStorageCount        *uint                               `json:"allowedNetworkStorageCount:omitempty"`
	AllowedNetworkStorageReplicaCount *uint                               `json:"allowedNetworkStorageReplicaCount:omitempty"`
	AllowedNetworkStorageReplicas     []Network_Storage                   `json:"allowedNetworkStorageReplicas:omitempty"`
	BillingItem                       *Billing_Item                       `json:"billingItem:omitempty"`
	BoundDescendantCount              *uint                               `json:"boundDescendantCount:omitempty"`
	BoundDescendants                  []Network_Subnet                    `json:"boundDescendants:omitempty"`
	BoundRouterCount                  *uint                               `json:"boundRouterCount:omitempty"`
	BoundRouterFlag                   *bool                               `json:"boundRouterFlag:omitempty"`
	BoundRouters                      []Hardware                          `json:"boundRouters:omitempty"`
	BroadcastAddress                  *string                             `json:"broadcastAddress:omitempty"`
	Children                          []Network_Subnet                    `json:"children:omitempty"`
	ChildrenCount                     *uint                               `json:"childrenCount:omitempty"`
	Cidr                              *int                                `json:"cidr:omitempty"`
	Datacenter                        *Location_Datacenter                `json:"datacenter:omitempty"`
	DescendantCount                   *uint                               `json:"descendantCount:omitempty"`
	Descendants                       []Network_Subnet                    `json:"descendants:omitempty"`
	DisplayLabel                      *string                             `json:"displayLabel:omitempty"`
	EndPointIpAddress                 *Network_Subnet_IpAddress           `json:"endPointIpAddress:omitempty"`
	Gateway                           *string                             `json:"gateway:omitempty"`
	GlobalIpRecord                    *Network_Subnet_IpAddress_Global    `json:"globalIpRecord:omitempty"`
	Hardware                          []Hardware                          `json:"hardware:omitempty"`
	HardwareCount                     *uint                               `json:"hardwareCount:omitempty"`
	Id                                *int                                `json:"id:omitempty"`
	IpAddressCount                    *uint                               `json:"ipAddressCount:omitempty"`
	IpAddresses                       []Network_Subnet_IpAddress          `json:"ipAddresses:omitempty"`
	IsCustomerOwned                   *bool                               `json:"isCustomerOwned:omitempty"`
	IsCustomerRoutable                *bool                               `json:"isCustomerRoutable:omitempty"`
	ModifyDate                        *time.Time                          `json:"modifyDate:omitempty"`
	Netmask                           *string                             `json:"netmask:omitempty"`
	NetworkComponent                  *Network_Component                  `json:"networkComponent:omitempty"`
	NetworkComponentFirewall          *Network_Component_Firewall         `json:"networkComponentFirewall:omitempty"`
	NetworkIdentifier                 *string                             `json:"networkIdentifier:omitempty"`
	NetworkProtectionAddressCount     *uint                               `json:"networkProtectionAddressCount:omitempty"`
	NetworkProtectionAddresses        []Network_Protection_Address        `json:"networkProtectionAddresses:omitempty"`
	NetworkTunnelContextCount         *uint                               `json:"networkTunnelContextCount:omitempty"`
	NetworkTunnelContexts             []Network_Tunnel_Module_Context     `json:"networkTunnelContexts:omitempty"`
	NetworkVlan                       *Network_Vlan                       `json:"networkVlan:omitempty"`
	NetworkVlanId                     *int                                `json:"networkVlanId:omitempty"`
	Note                              *string                             `json:"note:omitempty"`
	PodName                           *string                             `json:"podName:omitempty"`
	ProtectedIpAddressCount           *uint                               `json:"protectedIpAddressCount:omitempty"`
	ProtectedIpAddresses              []Network_Subnet_IpAddress          `json:"protectedIpAddresses:omitempty"`
	RegionalInternetRegistry          *Network_Regional_Internet_Registry `json:"regionalInternetRegistry:omitempty"`
	RegistrationCount                 *uint                               `json:"registrationCount:omitempty"`
	Registrations                     []Network_Subnet_Registration       `json:"registrations:omitempty"`
	ResourceGroupCount                *uint                               `json:"resourceGroupCount:omitempty"`
	ResourceGroups                    []Resource_Group                    `json:"resourceGroups:omitempty"`
	ReverseDomain                     *Dns_Domain                         `json:"reverseDomain:omitempty"`
	RoleKeyName                       *string                             `json:"roleKeyName:omitempty"`
	RoleName                          *string                             `json:"roleName:omitempty"`
	RoutingTypeKeyName                *string                             `json:"routingTypeKeyName:omitempty"`
	RoutingTypeName                   *string                             `json:"routingTypeName:omitempty"`
	SortOrder                         *string                             `json:"sortOrder:omitempty"`
	SubnetType                        *string                             `json:"subnetType:omitempty"`
	SwipTransaction                   []Network_Subnet_Swip_Transaction   `json:"swipTransaction:omitempty"`
	SwipTransactionCount              *uint                               `json:"swipTransactionCount:omitempty"`
	TotalIpAddresses                  *float64                            `json:"totalIpAddresses:omitempty"`
	UnboundDescendantCount            *uint                               `json:"unboundDescendantCount:omitempty"`
	UnboundDescendants                []Network_Subnet                    `json:"unboundDescendants:omitempty"`
	UsableIpAddressCount              *float64                            `json:"usableIpAddressCount:omitempty"`
	Version                           *int                                `json:"version:omitempty"`
	VirtualGuestCount                 *uint                               `json:"virtualGuestCount:omitempty"`
	VirtualGuests                     []Virtual_Guest                     `json:"virtualGuests:omitempty"`
}

type Network_Subnet_IpAddress struct {
	Entity

	AllowedHost                                      *Network_Storage_Allowed_Host                       `json:"allowedHost:omitempty"`
	AllowedNetworkStorage                            []Network_Storage                                   `json:"allowedNetworkStorage:omitempty"`
	AllowedNetworkStorageCount                       *uint                                               `json:"allowedNetworkStorageCount:omitempty"`
	AllowedNetworkStorageReplicaCount                *uint                                               `json:"allowedNetworkStorageReplicaCount:omitempty"`
	AllowedNetworkStorageReplicas                    []Network_Storage                                   `json:"allowedNetworkStorageReplicas:omitempty"`
	ApplicationDeliveryController                    *Network_Application_Delivery_Controller            `json:"applicationDeliveryController:omitempty"`
	ContextTunnelTranslationCount                    *uint                                               `json:"contextTunnelTranslationCount:omitempty"`
	ContextTunnelTranslations                        []Network_Tunnel_Module_Context_Address_Translation `json:"contextTunnelTranslations:omitempty"`
	EndpointSubnetCount                              *uint                                               `json:"endpointSubnetCount:omitempty"`
	EndpointSubnets                                  []Network_Subnet                                    `json:"endpointSubnets:omitempty"`
	GuestNetworkComponent                            *Virtual_Guest_Network_Component                    `json:"guestNetworkComponent:omitempty"`
	GuestNetworkComponentBinding                     *Virtual_Guest_Network_Component_IpAddress          `json:"guestNetworkComponentBinding:omitempty"`
	Hardware                                         *Hardware                                           `json:"hardware:omitempty"`
	Id                                               *int                                                `json:"id:omitempty"`
	IpAddress                                        *string                                             `json:"ipAddress:omitempty"`
	IsBroadcast                                      *bool                                               `json:"isBroadcast:omitempty"`
	IsGateway                                        *bool                                               `json:"isGateway:omitempty"`
	IsNetwork                                        *bool                                               `json:"isNetwork:omitempty"`
	IsReserved                                       *bool                                               `json:"isReserved:omitempty"`
	NetworkComponent                                 *Network_Component                                  `json:"networkComponent:omitempty"`
	Note                                             *string                                             `json:"note:omitempty"`
	PrivateNetworkGateway                            *Network_Gateway                                    `json:"privateNetworkGateway:omitempty"`
	ProtectionAddress                                []Network_Protection_Address                        `json:"protectionAddress:omitempty"`
	ProtectionAddressCount                           *uint                                               `json:"protectionAddressCount:omitempty"`
	PublicNetworkGateway                             *Network_Gateway                                    `json:"publicNetworkGateway:omitempty"`
	RemoteManagementNetworkComponent                 *Network_Component                                  `json:"remoteManagementNetworkComponent:omitempty"`
	Subnet                                           *Network_Subnet                                     `json:"subnet:omitempty"`
	SubnetId                                         *int                                                `json:"subnetId:omitempty"`
	SyslogEventsOneDay                               []Network_Logging_Syslog                            `json:"syslogEventsOneDay:omitempty"`
	SyslogEventsOneDayCount                          *uint                                               `json:"syslogEventsOneDayCount:omitempty"`
	SyslogEventsSevenDayCount                        *uint                                               `json:"syslogEventsSevenDayCount:omitempty"`
	SyslogEventsSevenDays                            []Network_Logging_Syslog                            `json:"syslogEventsSevenDays:omitempty"`
	TopTenSyslogEventsByDestinationPortOneDay        []Network_Logging_Syslog                            `json:"topTenSyslogEventsByDestinationPortOneDay:omitempty"`
	TopTenSyslogEventsByDestinationPortOneDayCount   *uint                                               `json:"topTenSyslogEventsByDestinationPortOneDayCount:omitempty"`
	TopTenSyslogEventsByDestinationPortSevenDayCount *uint                                               `json:"topTenSyslogEventsByDestinationPortSevenDayCount:omitempty"`
	TopTenSyslogEventsByDestinationPortSevenDays     []Network_Logging_Syslog                            `json:"topTenSyslogEventsByDestinationPortSevenDays:omitempty"`
	TopTenSyslogEventsByProtocolsOneDay              []Network_Logging_Syslog                            `json:"topTenSyslogEventsByProtocolsOneDay:omitempty"`
	TopTenSyslogEventsByProtocolsOneDayCount         *uint                                               `json:"topTenSyslogEventsByProtocolsOneDayCount:omitempty"`
	TopTenSyslogEventsByProtocolsSevenDayCount       *uint                                               `json:"topTenSyslogEventsByProtocolsSevenDayCount:omitempty"`
	TopTenSyslogEventsByProtocolsSevenDays           []Network_Logging_Syslog                            `json:"topTenSyslogEventsByProtocolsSevenDays:omitempty"`
	TopTenSyslogEventsBySourceIpOneDay               []Network_Logging_Syslog                            `json:"topTenSyslogEventsBySourceIpOneDay:omitempty"`
	TopTenSyslogEventsBySourceIpOneDayCount          *uint                                               `json:"topTenSyslogEventsBySourceIpOneDayCount:omitempty"`
	TopTenSyslogEventsBySourceIpSevenDayCount        *uint                                               `json:"topTenSyslogEventsBySourceIpSevenDayCount:omitempty"`
	TopTenSyslogEventsBySourceIpSevenDays            []Network_Logging_Syslog                            `json:"topTenSyslogEventsBySourceIpSevenDays:omitempty"`
	TopTenSyslogEventsBySourcePortOneDay             []Network_Logging_Syslog                            `json:"topTenSyslogEventsBySourcePortOneDay:omitempty"`
	TopTenSyslogEventsBySourcePortOneDayCount        *uint                                               `json:"topTenSyslogEventsBySourcePortOneDayCount:omitempty"`
	TopTenSyslogEventsBySourcePortSevenDayCount      *uint                                               `json:"topTenSyslogEventsBySourcePortSevenDayCount:omitempty"`
	TopTenSyslogEventsBySourcePortSevenDays          []Network_Logging_Syslog                            `json:"topTenSyslogEventsBySourcePortSevenDays:omitempty"`
	VirtualGuest                                     *Virtual_Guest                                      `json:"virtualGuest:omitempty"`
	VirtualLicenseCount                              *uint                                               `json:"virtualLicenseCount:omitempty"`
	VirtualLicenses                                  []Software_VirtualLicense                           `json:"virtualLicenses:omitempty"`
}

type Network_Subnet_IpAddress_Global struct {
	Entity

	Account                *Account                                      `json:"account:omitempty"`
	ActiveTransaction      *Provisioning_Version1_Transaction            `json:"activeTransaction:omitempty"`
	BillingItem            *Billing_Item_Network_Subnet_IpAddress_Global `json:"billingItem:omitempty"`
	Description            *int                                          `json:"description:omitempty"`
	DestinationIpAddress   *Network_Subnet_IpAddress                     `json:"destinationIpAddress:omitempty"`
	DestinationIpAddressId *int                                          `json:"destinationIpAddressId:omitempty"`
	Id                     *int                                          `json:"id:omitempty"`
	IpAddress              *Network_Subnet_IpAddress                     `json:"ipAddress:omitempty"`
	IpAddressId            *int                                          `json:"ipAddressId:omitempty"`
	TypeId                 *int                                          `json:"typeId:omitempty"`
}

type Network_Subnet_IpAddress_Version6 struct {
	Network_Subnet_IpAddress

	PublicVersion6NetworkGateway *Network_Gateway `json:"publicVersion6NetworkGateway:omitempty"`
}

type Network_Subnet_Registration struct {
	Entity

	Account                          *Account                              `json:"account:omitempty"`
	AccountId                        *int                                  `json:"accountId:omitempty"`
	Cidr                             *int                                  `json:"cidr:omitempty"`
	CreateDate                       *time.Time                            `json:"createDate:omitempty"`
	DetailReferenceCount             *uint                                 `json:"detailReferenceCount:omitempty"`
	DetailReferences                 []Network_Subnet_Registration_Details `json:"detailReferences:omitempty"`
	EventCount                       *uint                                 `json:"eventCount:omitempty"`
	Events                           []Network_Subnet_Registration_Event   `json:"events:omitempty"`
	Id                               *int                                  `json:"id:omitempty"`
	ModifyDate                       *time.Time                            `json:"modifyDate:omitempty"`
	NetworkDetail                    *Account_Regional_Registry_Detail     `json:"networkDetail:omitempty"`
	NetworkHandle                    *string                               `json:"networkHandle:omitempty"`
	NetworkIdentifier                *string                               `json:"networkIdentifier:omitempty"`
	PersonDetail                     *Account_Regional_Registry_Detail     `json:"personDetail:omitempty"`
	RegionalInternetRegistry         *Network_Regional_Internet_Registry   `json:"regionalInternetRegistry:omitempty"`
	RegionalInternetRegistryHandle   *Account_Rwhois_Handle                `json:"regionalInternetRegistryHandle:omitempty"`
	RegionalInternetRegistryHandleId *int                                  `json:"regionalInternetRegistryHandleId:omitempty"`
	RegionalInternetRegistryId       *int                                  `json:"regionalInternetRegistryId:omitempty"`
	Status                           *Network_Subnet_Registration_Status   `json:"status:omitempty"`
	StatusId                         *int                                  `json:"statusId:omitempty"`
	Subnet                           *Network_Subnet                       `json:"subnet:omitempty"`
}

type Network_Subnet_Registration_Apnic struct {
	Network_Subnet_Registration
}

type Network_Subnet_Registration_Arin struct {
	Network_Subnet_Registration
}

type Network_Subnet_Registration_Details struct {
	Entity

	CreateDate     *time.Time                        `json:"createDate:omitempty"`
	Detail         *Account_Regional_Registry_Detail `json:"detail:omitempty"`
	DetailId       *int                              `json:"detailId:omitempty"`
	Id             *int                              `json:"id:omitempty"`
	ModifyDate     *time.Time                        `json:"modifyDate:omitempty"`
	Registration   *Network_Subnet_Registration      `json:"registration:omitempty"`
	RegistrationId *int                              `json:"registrationId:omitempty"`
}

type Network_Subnet_Registration_Event struct {
	Entity

	CreateDate     *time.Time                              `json:"createDate:omitempty"`
	Id             *int                                    `json:"id:omitempty"`
	Message        *string                                 `json:"message:omitempty"`
	ModifyDate     *time.Time                              `json:"modifyDate:omitempty"`
	Registration   *Network_Subnet_Registration            `json:"registration:omitempty"`
	RegistrationId *int                                    `json:"registrationId:omitempty"`
	Type           *Network_Subnet_Registration_Event_Type `json:"type:omitempty"`
	TypeId         *int                                    `json:"typeId:omitempty"`
}

type Network_Subnet_Registration_Event_Type struct {
	Entity

	CreateDate *time.Time `json:"createDate:omitempty"`
	Id         *int       `json:"id:omitempty"`
	KeyName    *string    `json:"keyName:omitempty"`
	ModifyDate *time.Time `json:"modifyDate:omitempty"`
	Name       *string    `json:"name:omitempty"`
}

type Network_Subnet_Registration_Ripe struct {
	Network_Subnet_Registration
}

type Network_Subnet_Registration_Status struct {
	Entity

	CreateDate *time.Time `json:"createDate:omitempty"`
	Id         *int       `json:"id:omitempty"`
	KeyName    *string    `json:"keyName:omitempty"`
	ModifyDate *time.Time `json:"modifyDate:omitempty"`
	Name       *string    `json:"name:omitempty"`
}

type Network_Subnet_Rwhois_Data struct {
	Entity

	AbuseEmail           *string    `json:"abuseEmail:omitempty"`
	Account              *Account   `json:"account:omitempty"`
	AccountId            *int       `json:"accountId:omitempty"`
	Address1             *string    `json:"address1:omitempty"`
	Address2             *string    `json:"address2:omitempty"`
	City                 *string    `json:"city:omitempty"`
	CompanyName          *string    `json:"companyName:omitempty"`
	Country              *string    `json:"country:omitempty"`
	CreateDate           *time.Time `json:"createDate:omitempty"`
	FirstName            *string    `json:"firstName:omitempty"`
	Id                   *int       `json:"id:omitempty"`
	LastName             *string    `json:"lastName:omitempty"`
	ModifyDate           *time.Time `json:"modifyDate:omitempty"`
	PostalCode           *string    `json:"postalCode:omitempty"`
	PrivateResidenceFlag *bool      `json:"privateResidenceFlag:omitempty"`
	State                *string    `json:"state:omitempty"`
}

type Network_Subnet_Swip_Transaction struct {
	Entity

	Account    *Account        `json:"account:omitempty"`
	Id         *int            `json:"id:omitempty"`
	StatusName *string         `json:"statusName:omitempty"`
	Subnet     *Network_Subnet `json:"subnet:omitempty"`
	SubnetId   *int            `json:"subnetId:omitempty"`
}

type Network_TippingPointReporting struct {
	Entity
}

type Network_Tunnel_Module_Context struct {
	Entity

	Account                        *Account                                            `json:"account:omitempty"`
	AccountId                      *int                                                `json:"accountId:omitempty"`
	ActiveTransaction              *Provisioning_Version1_Transaction                  `json:"activeTransaction:omitempty"`
	AddressTranslationCount        *uint                                               `json:"addressTranslationCount:omitempty"`
	AddressTranslations            []Network_Tunnel_Module_Context_Address_Translation `json:"addressTranslations:omitempty"`
	AdvancedConfigurationFlag      *int                                                `json:"advancedConfigurationFlag:omitempty"`
	AllAvailableServiceSubnetCount *uint                                               `json:"allAvailableServiceSubnetCount:omitempty"`
	AllAvailableServiceSubnets     []Network_Subnet                                    `json:"allAvailableServiceSubnets:omitempty"`
	BillingItem                    *Billing_Item                                       `json:"billingItem:omitempty"`
	CreateDate                     *time.Time                                          `json:"createDate:omitempty"`
	CustomerPeerIpAddress          *string                                             `json:"customerPeerIpAddress:omitempty"`
	CustomerSubnetCount            *uint                                               `json:"customerSubnetCount:omitempty"`
	CustomerSubnets                []Network_Customer_Subnet                           `json:"customerSubnets:omitempty"`
	Datacenter                     *Location                                           `json:"datacenter:omitempty"`
	FriendlyName                   *string                                             `json:"friendlyName:omitempty"`
	Id                             *int                                                `json:"id:omitempty"`
	InternalPeerIpAddress          *string                                             `json:"internalPeerIpAddress:omitempty"`
	InternalSubnetCount            *uint                                               `json:"internalSubnetCount:omitempty"`
	InternalSubnets                []Network_Subnet                                    `json:"internalSubnets:omitempty"`
	ModifyDate                     *time.Time                                          `json:"modifyDate:omitempty"`
	Name                           *string                                             `json:"name:omitempty"`
	PhaseOneAuthentication         *string                                             `json:"phaseOneAuthentication:omitempty"`
	PhaseOneDiffieHellmanGroup     *int                                                `json:"phaseOneDiffieHellmanGroup:omitempty"`
	PhaseOneEncryption             *string                                             `json:"phaseOneEncryption:omitempty"`
	PhaseOneKeylife                *int                                                `json:"phaseOneKeylife:omitempty"`
	PhaseTwoAuthentication         *string                                             `json:"phaseTwoAuthentication:omitempty"`
	PhaseTwoDiffieHellmanGroup     *int                                                `json:"phaseTwoDiffieHellmanGroup:omitempty"`
	PhaseTwoEncryption             *string                                             `json:"phaseTwoEncryption:omitempty"`
	PhaseTwoKeylife                *int                                                `json:"phaseTwoKeylife:omitempty"`
	PhaseTwoPerfectForwardSecrecy  *int                                                `json:"phaseTwoPerfectForwardSecrecy:omitempty"`
	PresharedKey                   *string                                             `json:"presharedKey:omitempty"`
	ServiceSubnetCount             *uint                                               `json:"serviceSubnetCount:omitempty"`
	ServiceSubnets                 []Network_Subnet                                    `json:"serviceSubnets:omitempty"`
	StaticRouteSubnetCount         *uint                                               `json:"staticRouteSubnetCount:omitempty"`
	StaticRouteSubnets             []Network_Subnet                                    `json:"staticRouteSubnets:omitempty"`
	TransactionHistory             []Provisioning_Version1_Transaction                 `json:"transactionHistory:omitempty"`
	TransactionHistoryCount        *uint                                               `json:"transactionHistoryCount:omitempty"`
}

type Network_Tunnel_Module_Context_Address_Translation struct {
	Entity

	CustomerIpAddress       *string                            `json:"customerIpAddress:omitempty"`
	CustomerIpAddressId     *int                               `json:"customerIpAddressId:omitempty"`
	CustomerIpAddressRecord *Network_Customer_Subnet_IpAddress `json:"customerIpAddressRecord:omitempty"`
	Id                      *int                               `json:"id:omitempty"`
	InternalIpAddress       *string                            `json:"internalIpAddress:omitempty"`
	InternalIpAddressId     *int                               `json:"internalIpAddressId:omitempty"`
	InternalIpAddressRecord *Network_Subnet_IpAddress          `json:"internalIpAddressRecord:omitempty"`
	NetworkTunnelContext    *Network_Tunnel_Module_Context     `json:"networkTunnelContext:omitempty"`
	NetworkTunnelContextId  *int                               `json:"networkTunnelContextId:omitempty"`
	Notes                   *string                            `json:"notes:omitempty"`
}

type Network_Vlan struct {
	Entity

	Account                            *Account                                    `json:"account:omitempty"`
	AccountId                          *int                                        `json:"accountId:omitempty"`
	AdditionalPrimarySubnetCount       *uint                                       `json:"additionalPrimarySubnetCount:omitempty"`
	AdditionalPrimarySubnets           []Network_Subnet                            `json:"additionalPrimarySubnets:omitempty"`
	AttachedNetworkGateway             *Network_Gateway                            `json:"attachedNetworkGateway:omitempty"`
	AttachedNetworkGatewayFlag         *bool                                       `json:"attachedNetworkGatewayFlag:omitempty"`
	AttachedNetworkGatewayVlan         *Network_Gateway_Vlan                       `json:"attachedNetworkGatewayVlan:omitempty"`
	BillingItem                        *Billing_Item                               `json:"billingItem:omitempty"`
	DedicatedFirewallFlag              *int                                        `json:"dedicatedFirewallFlag:omitempty"`
	ExtensionRouter                    *Hardware_Router                            `json:"extensionRouter:omitempty"`
	FirewallGuestNetworkComponentCount *uint                                       `json:"firewallGuestNetworkComponentCount:omitempty"`
	FirewallGuestNetworkComponents     []Network_Component_Firewall                `json:"firewallGuestNetworkComponents:omitempty"`
	FirewallInterfaceCount             *uint                                       `json:"firewallInterfaceCount:omitempty"`
	FirewallInterfaces                 []Network_Firewall_Module_Context_Interface `json:"firewallInterfaces:omitempty"`
	FirewallNetworkComponentCount      *uint                                       `json:"firewallNetworkComponentCount:omitempty"`
	FirewallNetworkComponents          []Network_Component_Firewall                `json:"firewallNetworkComponents:omitempty"`
	FirewallRuleCount                  *uint                                       `json:"firewallRuleCount:omitempty"`
	FirewallRules                      []Network_Vlan_Firewall_Rule                `json:"firewallRules:omitempty"`
	GuestNetworkComponentCount         *uint                                       `json:"guestNetworkComponentCount:omitempty"`
	GuestNetworkComponents             []Virtual_Guest_Network_Component           `json:"guestNetworkComponents:omitempty"`
	Hardware                           []Hardware                                  `json:"hardware:omitempty"`
	HardwareCount                      *uint                                       `json:"hardwareCount:omitempty"`
	HighAvailabilityFirewallFlag       *bool                                       `json:"highAvailabilityFirewallFlag:omitempty"`
	Id                                 *int                                        `json:"id:omitempty"`
	LocalDiskStorageCapabilityFlag     *bool                                       `json:"localDiskStorageCapabilityFlag:omitempty"`
	ModifyDate                         *time.Time                                  `json:"modifyDate:omitempty"`
	Name                               *string                                     `json:"name:omitempty"`
	Network                            *Network                                    `json:"network:omitempty"`
	NetworkComponentCount              *uint                                       `json:"networkComponentCount:omitempty"`
	NetworkComponentTrunkCount         *uint                                       `json:"networkComponentTrunkCount:omitempty"`
	NetworkComponentTrunks             []Network_Component_Network_Vlan_Trunk      `json:"networkComponentTrunks:omitempty"`
	NetworkComponents                  []Network_Component                         `json:"networkComponents:omitempty"`
	NetworkSpace                       *string                                     `json:"networkSpace:omitempty"`
	NetworkVlanFirewall                *Network_Vlan_Firewall                      `json:"networkVlanFirewall:omitempty"`
	Note                               *string                                     `json:"note:omitempty"`
	PrimaryRouter                      *Hardware_Router                            `json:"primaryRouter:omitempty"`
	PrimarySubnet                      *Network_Subnet                             `json:"primarySubnet:omitempty"`
	PrimarySubnetCount                 *uint                                       `json:"primarySubnetCount:omitempty"`
	PrimarySubnetId                    *int                                        `json:"primarySubnetId:omitempty"`
	PrimarySubnetVersion6              *Network_Subnet                             `json:"primarySubnetVersion6:omitempty"`
	PrimarySubnets                     []Network_Subnet                            `json:"primarySubnets:omitempty"`
	PrivateNetworkGatewayCount         *uint                                       `json:"privateNetworkGatewayCount:omitempty"`
	PrivateNetworkGateways             []Network_Gateway                           `json:"privateNetworkGateways:omitempty"`
	ProtectedIpAddressCount            *uint                                       `json:"protectedIpAddressCount:omitempty"`
	ProtectedIpAddresses               []Network_Subnet_IpAddress                  `json:"protectedIpAddresses:omitempty"`
	PublicNetworkGatewayCount          *uint                                       `json:"publicNetworkGatewayCount:omitempty"`
	PublicNetworkGateways              []Network_Gateway                           `json:"publicNetworkGateways:omitempty"`
	ResourceGroupCount                 *uint                                       `json:"resourceGroupCount:omitempty"`
	ResourceGroupMember                []Resource_Group_Member                     `json:"resourceGroupMember:omitempty"`
	ResourceGroupMemberCount           *uint                                       `json:"resourceGroupMemberCount:omitempty"`
	ResourceGroups                     []Resource_Group                            `json:"resourceGroups:omitempty"`
	SanStorageCapabilityFlag           *bool                                       `json:"sanStorageCapabilityFlag:omitempty"`
	ScaleVlanCount                     *uint                                       `json:"scaleVlanCount:omitempty"`
	ScaleVlans                         []Scale_Network_Vlan                        `json:"scaleVlans:omitempty"`
	SecondaryRouter                    *Hardware                                   `json:"secondaryRouter:omitempty"`
	SecondarySubnetCount               *uint                                       `json:"secondarySubnetCount:omitempty"`
	SecondarySubnets                   []Network_Subnet                            `json:"secondarySubnets:omitempty"`
	SubnetCount                        *uint                                       `json:"subnetCount:omitempty"`
	Subnets                            []Network_Subnet                            `json:"subnets:omitempty"`
	TagReferenceCount                  *uint                                       `json:"tagReferenceCount:omitempty"`
	TagReferences                      []Tag_Reference                             `json:"tagReferences:omitempty"`
	TotalPrimaryIpAddressCount         *uint                                       `json:"totalPrimaryIpAddressCount:omitempty"`
	Type                               *Network_Vlan_Type                          `json:"type:omitempty"`
	VirtualGuestCount                  *uint                                       `json:"virtualGuestCount:omitempty"`
	VirtualGuests                      []Virtual_Guest                             `json:"virtualGuests:omitempty"`
	VlanNumber                         *int                                        `json:"vlanNumber:omitempty"`
}

type Network_Vlan_Firewall struct {
	Entity

	AdministrativeBypassFlag          *string                           `json:"administrativeBypassFlag:omitempty"`
	BillingItem                       *Billing_Item                     `json:"billingItem:omitempty"`
	CustomerManagedFlag               *bool                             `json:"customerManagedFlag:omitempty"`
	Datacenter                        *Location                         `json:"datacenter:omitempty"`
	FirewallType                      *string                           `json:"firewallType:omitempty"`
	FullyQualifiedDomainName          *string                           `json:"fullyQualifiedDomainName:omitempty"`
	Id                                *int                              `json:"id:omitempty"`
	ManagementCredentials             *Software_Component_Password      `json:"managementCredentials:omitempty"`
	NetworkFirewallUpdateRequestCount *uint                             `json:"networkFirewallUpdateRequestCount:omitempty"`
	NetworkFirewallUpdateRequests     []Network_Firewall_Update_Request `json:"networkFirewallUpdateRequests:omitempty"`
	NetworkVlan                       *Network_Vlan                     `json:"networkVlan:omitempty"`
	NetworkVlanCount                  *uint                             `json:"networkVlanCount:omitempty"`
	NetworkVlans                      []Network_Vlan                    `json:"networkVlans:omitempty"`
	PrimaryIpAddress                  *string                           `json:"primaryIpAddress:omitempty"`
	RuleCount                         *uint                             `json:"ruleCount:omitempty"`
	Rules                             []Network_Vlan_Firewall_Rule      `json:"rules:omitempty"`
	TagReferenceCount                 *uint                             `json:"tagReferenceCount:omitempty"`
	TagReferences                     []Tag_Reference                   `json:"tagReferences:omitempty"`
}

type Network_Vlan_Firewall_Rule struct {
	Entity

	Action                    *string                     `json:"action:omitempty"`
	DestinationIpAddress      *string                     `json:"destinationIpAddress:omitempty"`
	DestinationIpCidr         *int                        `json:"destinationIpCidr:omitempty"`
	DestinationIpSubnetMask   *string                     `json:"destinationIpSubnetMask:omitempty"`
	DestinationPortRangeEnd   *int                        `json:"destinationPortRangeEnd:omitempty"`
	DestinationPortRangeStart *int                        `json:"destinationPortRangeStart:omitempty"`
	Id                        *int                        `json:"id:omitempty"`
	NetworkComponentFirewall  *Network_Component_Firewall `json:"networkComponentFirewall:omitempty"`
	Notes                     *string                     `json:"notes:omitempty"`
	OrderValue                *int                        `json:"orderValue:omitempty"`
	Protocol                  *string                     `json:"protocol:omitempty"`
	SourceIpAddress           *string                     `json:"sourceIpAddress:omitempty"`
	SourceIpCidr              *int                        `json:"sourceIpCidr:omitempty"`
	SourceIpSubnetMask        *string                     `json:"sourceIpSubnetMask:omitempty"`
	Status                    *string                     `json:"status:omitempty"`
	Version                   *int                        `json:"version:omitempty"`
}

type Network_Vlan_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Notification struct {
	Entity

	Id                      *int                      `json:"id:omitempty"`
	KeyName                 *string                   `json:"keyName:omitempty"`
	Name                    *string                   `json:"name:omitempty"`
	PreferenceCount         *uint                     `json:"preferenceCount:omitempty"`
	Preferences             []Notification_Preference `json:"preferences:omitempty"`
	RequiredPreferenceCount *uint                     `json:"requiredPreferenceCount:omitempty"`
	RequiredPreferences     []Notification_Preference `json:"requiredPreferences:omitempty"`
}

type Notification_Delivery_Method struct {
	Entity

	Active      *int    `json:"active:omitempty"`
	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Notification_Mobile struct {
	Notification
}

type Notification_Occurrence_Account struct {
	Entity

	Account                     *Account                        `json:"account:omitempty"`
	Active                      *int                            `json:"active:omitempty"`
	LastNotificationUpdate      *Notification_Occurrence_Update `json:"lastNotificationUpdate:omitempty"`
	NotificationOccurrenceEvent *Notification_Occurrence_Event  `json:"notificationOccurrenceEvent:omitempty"`
}

type Notification_Occurrence_Event struct {
	Entity

	AcknowledgedFlag                *bool                                      `json:"acknowledgedFlag:omitempty"`
	AttachmentCount                 *uint                                      `json:"attachmentCount:omitempty"`
	Attachments                     []Notification_Occurrence_Event_Attachment `json:"attachments:omitempty"`
	EndDate                         *time.Time                                 `json:"endDate:omitempty"`
	FirstUpdate                     *Notification_Occurrence_Update            `json:"firstUpdate:omitempty"`
	Id                              *int                                       `json:"id:omitempty"`
	ImpactedAccountCount            *uint                                      `json:"impactedAccountCount:omitempty"`
	ImpactedAccounts                []Notification_Occurrence_Account          `json:"impactedAccounts:omitempty"`
	ImpactedResourceCount           *uint                                      `json:"impactedResourceCount:omitempty"`
	ImpactedResources               []Notification_Occurrence_Resource         `json:"impactedResources:omitempty"`
	ImpactedUserCount               *uint                                      `json:"impactedUserCount:omitempty"`
	ImpactedUsers                   []Notification_Occurrence_User             `json:"impactedUsers:omitempty"`
	LastImpactedUserCount           *int                                       `json:"lastImpactedUserCount:omitempty"`
	LastUpdate                      *Notification_Occurrence_Update            `json:"lastUpdate:omitempty"`
	ModifyDate                      *time.Time                                 `json:"modifyDate:omitempty"`
	NotificationOccurrenceEventType *Notification_Occurrence_Event_Type        `json:"notificationOccurrenceEventType:omitempty"`
	RecoveryTime                    *int                                       `json:"recoveryTime:omitempty"`
	StartDate                       *time.Time                                 `json:"startDate:omitempty"`
	StatusCode                      *Notification_Occurrence_Status_Code       `json:"statusCode:omitempty"`
	Subject                         *string                                    `json:"subject:omitempty"`
	Summary                         *string                                    `json:"summary:omitempty"`
	SystemTicketId                  *int                                       `json:"systemTicketId:omitempty"`
	UpdateCount                     *uint                                      `json:"updateCount:omitempty"`
	Updates                         []Notification_Occurrence_Update           `json:"updates:omitempty"`
}

type Notification_Occurrence_Event_Attachment struct {
	Entity

	CreateDate                    *time.Time                     `json:"createDate:omitempty"`
	FileName                      *string                        `json:"fileName:omitempty"`
	FileSize                      *string                        `json:"fileSize:omitempty"`
	Id                            *int                           `json:"id:omitempty"`
	NotificationOccurrenceEvent   *Notification_Occurrence_Event `json:"notificationOccurrenceEvent:omitempty"`
	NotificationOccurrenceEventId *int                           `json:"notificationOccurrenceEventId:omitempty"`
}

type Notification_Occurrence_Event_Type struct {
	Entity

	KeyName *string `json:"keyName:omitempty"`
}

type Notification_Occurrence_Resource struct {
	Entity

	Active                        *int                           `json:"active:omitempty"`
	FilterLabel                   *string                        `json:"filterLabel:omitempty"`
	NotificationOccurrenceEvent   *Notification_Occurrence_Event `json:"notificationOccurrenceEvent:omitempty"`
	NotificationOccurrenceEventId *int                           `json:"notificationOccurrenceEventId:omitempty"`
	Resource                      *Entity                        `json:"resource:omitempty"`
	ResourceAccountId             *int                           `json:"resourceAccountId:omitempty"`
	ResourceName                  *string                        `json:"resourceName:omitempty"`
	ResourceTableId               *int                           `json:"resourceTableId:omitempty"`
}

type Notification_Occurrence_Resource_Hardware struct {
	Notification_Occurrence_Resource

	Hostname     *string `json:"hostname:omitempty"`
	PrivateIp    *string `json:"privateIp:omitempty"`
	PublicIp     *string `json:"publicIp:omitempty"`
	ResourceType *string `json:"resourceType:omitempty"`
}

type Notification_Occurrence_Resource_Network_Application_Delivery_Controller struct {
	Notification_Occurrence_Resource

	Hostname     *string `json:"hostname:omitempty"`
	PrivateIp    *string `json:"privateIp:omitempty"`
	PublicIp     *string `json:"publicIp:omitempty"`
	ResourceType *string `json:"resourceType:omitempty"`
}

type Notification_Occurrence_Resource_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress struct {
	Notification_Occurrence_Resource

	Hostname     *string `json:"hostname:omitempty"`
	PublicIp     *string `json:"publicIp:omitempty"`
	ResourceType *string `json:"resourceType:omitempty"`
}

type Notification_Occurrence_Resource_Network_Storage_Iscsi_EqualLogic struct {
	Notification_Occurrence_Resource

	Hostname     *string `json:"hostname:omitempty"`
	PrivateIp    *string `json:"privateIp:omitempty"`
	ResourceType *string `json:"resourceType:omitempty"`
}

type Notification_Occurrence_Resource_Network_Storage_Iscsi_NetApp struct {
	Notification_Occurrence_Resource

	Hostname     *string `json:"hostname:omitempty"`
	PrivateIp    *string `json:"privateIp:omitempty"`
	ResourceType *string `json:"resourceType:omitempty"`
}

type Notification_Occurrence_Resource_Network_Storage_Lockbox struct {
	Notification_Occurrence_Resource

	Hostname     *string `json:"hostname:omitempty"`
	PrivateIp    *string `json:"privateIp:omitempty"`
	ResourceType *string `json:"resourceType:omitempty"`
}

type Notification_Occurrence_Resource_Network_Storage_Nas struct {
	Notification_Occurrence_Resource

	Hostname     *string `json:"hostname:omitempty"`
	PrivateIp    *string `json:"privateIp:omitempty"`
	ResourceType *string `json:"resourceType:omitempty"`
}

type Notification_Occurrence_Resource_Network_Storage_NetApp_Volume struct {
	Notification_Occurrence_Resource

	Hostname     *string `json:"hostname:omitempty"`
	PrivateIp    *string `json:"privateIp:omitempty"`
	ResourceType *string `json:"resourceType:omitempty"`
}

type Notification_Occurrence_Resource_Virtual struct {
	Notification_Occurrence_Resource

	Hostname     *string `json:"hostname:omitempty"`
	PrivateIp    *string `json:"privateIp:omitempty"`
	PublicIp     *string `json:"publicIp:omitempty"`
	ResourceType *string `json:"resourceType:omitempty"`
}

type Notification_Occurrence_Status_Code struct {
	Entity

	Description *string `json:"description:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Notification_Occurrence_Update struct {
	Entity

	Contents                    *string                        `json:"contents:omitempty"`
	CreateDate                  *time.Time                     `json:"createDate:omitempty"`
	Employee                    *User_Employee                 `json:"employee:omitempty"`
	EndDate                     *time.Time                     `json:"endDate:omitempty"`
	NotificationOccurrenceEvent *Notification_Occurrence_Event `json:"notificationOccurrenceEvent:omitempty"`
	StartDate                   *time.Time                     `json:"startDate:omitempty"`
}

type Notification_Occurrence_User struct {
	Entity

	AcknowledgedFlag            *int                               `json:"acknowledgedFlag:omitempty"`
	Active                      *int                               `json:"active:omitempty"`
	Id                          *int                               `json:"id:omitempty"`
	ImpactedResourceCount       *uint                              `json:"impactedResourceCount:omitempty"`
	ImpactedResources           []Notification_Occurrence_Resource `json:"impactedResources:omitempty"`
	NotificationOccurrenceEvent *Notification_Occurrence_Event     `json:"notificationOccurrenceEvent:omitempty"`
	User                        *User_Customer                     `json:"user:omitempty"`
	UsrRecordId                 *int                               `json:"usrRecordId:omitempty"`
}

type Notification_Preference struct {
	Entity

	Description  *string `json:"description:omitempty"`
	Id           *int    `json:"id:omitempty"`
	KeyName      *string `json:"keyName:omitempty"`
	MaximumValue *string `json:"maximumValue:omitempty"`
	MinimumValue *string `json:"minimumValue:omitempty"`
	Name         *string `json:"name:omitempty"`
	Units        *string `json:"units:omitempty"`
	Value        *string `json:"value:omitempty"`
}

type Notification_Subscriber struct {
	Entity

	Active                               *int                                      `json:"active:omitempty"`
	CreateDate                           *time.Time                                `json:"createDate:omitempty"`
	DeliveryMethodCount                  *uint                                     `json:"deliveryMethodCount:omitempty"`
	DeliveryMethods                      []Notification_Subscriber_Delivery_Method `json:"deliveryMethods:omitempty"`
	Id                                   *int                                      `json:"id:omitempty"`
	ModifyDate                           *time.Time                                `json:"modifyDate:omitempty"`
	Notification                         *Notification                             `json:"notification:omitempty"`
	NotificationId                       *int                                      `json:"notificationId:omitempty"`
	NotificationSubscriberTypeId         *int                                      `json:"notificationSubscriberTypeId:omitempty"`
	NotificationSubscriberTypeResourceId *int                                      `json:"notificationSubscriberTypeResourceId:omitempty"`
}

type Notification_Subscriber_Customer struct {
	Notification_Subscriber

	SubscriberRecord *User_Customer `json:"subscriberRecord:omitempty"`
}

type Notification_Subscriber_Delivery_Method struct {
	Entity

	Active                       *int                          `json:"active:omitempty"`
	CreateDate                   *time.Time                    `json:"createDate:omitempty"`
	ModifyDate                   *time.Time                    `json:"modifyDate:omitempty"`
	NotificationDeliveryMethod   *Notification_Delivery_Method `json:"notificationDeliveryMethod:omitempty"`
	NotificationDeliveryMethodId *int                          `json:"notificationDeliveryMethodId:omitempty"`
	NotificationSubscriber       *Notification_Subscriber      `json:"notificationSubscriber:omitempty"`
	NotificationSubscriberId     *int                          `json:"notificationSubscriberId:omitempty"`
}

type Notification_User_Subscriber struct {
	Entity

	Active                 *int                                      `json:"active:omitempty"`
	DeliveryMethodCount    *uint                                     `json:"deliveryMethodCount:omitempty"`
	DeliveryMethods        []Notification_Delivery_Method            `json:"deliveryMethods:omitempty"`
	Id                     *int                                      `json:"id:omitempty"`
	Notification           *Notification                             `json:"notification:omitempty"`
	NotificationId         *int                                      `json:"notificationId:omitempty"`
	PreferenceCount        *uint                                     `json:"preferenceCount:omitempty"`
	Preferences            []Notification_User_Subscriber_Preference `json:"preferences:omitempty"`
	PreferencesDetailCount *uint                                     `json:"preferencesDetailCount:omitempty"`
	PreferencesDetails     []Notification_Preference                 `json:"preferencesDetails:omitempty"`
	ResourceRecord         *Notification_User_Subscriber_Resource    `json:"resourceRecord:omitempty"`
	UserRecord             *User_Customer                            `json:"userRecord:omitempty"`
	UserRecordId           *int                                      `json:"userRecordId:omitempty"`
}

type Notification_User_Subscriber_Billing struct {
	Notification_User_Subscriber
}

type Notification_User_Subscriber_Delivery_Method struct {
	Entity

	Active                       *int                          `json:"active:omitempty"`
	DeliveryMethod               *Notification_Delivery_Method `json:"deliveryMethod:omitempty"`
	NotificationMethodId         *int                          `json:"notificationMethodId:omitempty"`
	NotificationUserSubscriber   *Notification_User_Subscriber `json:"notificationUserSubscriber:omitempty"`
	NotificationUserSubscriberId *int                          `json:"notificationUserSubscriberId:omitempty"`
}

type Notification_User_Subscriber_Mobile struct {
	Notification_User_Subscriber
}

type Notification_User_Subscriber_Preference struct {
	Entity

	DefaultPreference            *Notification_Preference      `json:"defaultPreference:omitempty"`
	Id                           *int                          `json:"id:omitempty"`
	NotificationPreferenceId     *int                          `json:"notificationPreferenceId:omitempty"`
	NotificationUserSubscriber   *Notification_User_Subscriber `json:"notificationUserSubscriber:omitempty"`
	NotificationUserSubscriberId *int                          `json:"notificationUserSubscriberId:omitempty"`
	Value                        *string                       `json:"value:omitempty"`
}

type Notification_User_Subscriber_Resource struct {
	Entity

	NotificationUserSubscriber   *Notification_User_Subscriber `json:"notificationUserSubscriber:omitempty"`
	NotificationUserSubscriberId *int                          `json:"notificationUserSubscriberId:omitempty"`
	ResourceTableId              *int                          `json:"resourceTableId:omitempty"`
}

type Product_Catalog struct {
	Entity

	BrandCount   *uint                `json:"brandCount:omitempty"`
	Brands       []Brand              `json:"brands:omitempty"`
	PackageCount *uint                `json:"packageCount:omitempty"`
	Packages     []Product_Package    `json:"packages:omitempty"`
	PriceCount   *uint                `json:"priceCount:omitempty"`
	Prices       []Product_Item_Price `json:"prices:omitempty"`
	ProductCount *uint                `json:"productCount:omitempty"`
	Products     []Product_Item       `json:"products:omitempty"`
}

type Product_Catalog_Item_Price struct {
	Entity

	Catalog    *Product_Catalog    `json:"catalog:omitempty"`
	CatalogId  *int                `json:"catalogId:omitempty"`
	CreateDate *time.Time          `json:"createDate:omitempty"`
	ModifyDate *time.Time          `json:"modifyDate:omitempty"`
	Price      *Product_Item_Price `json:"price:omitempty"`
	PriceId    *int                `json:"priceId:omitempty"`
}

type Product_Item struct {
	Entity

	ActivePresaleEventCount         *uint                             `json:"activePresaleEventCount:omitempty"`
	ActivePresaleEvents             []Sales_Presale_Event             `json:"activePresaleEvents:omitempty"`
	ActiveUsagePriceCount           *uint                             `json:"activeUsagePriceCount:omitempty"`
	ActiveUsagePrices               []Product_Item_Price              `json:"activeUsagePrices:omitempty"`
	AttributeCount                  *uint                             `json:"attributeCount:omitempty"`
	Attributes                      []Product_Item_Attribute          `json:"attributes:omitempty"`
	AvailabilityAttributeCount      *uint                             `json:"availabilityAttributeCount:omitempty"`
	AvailabilityAttributes          []Product_Item_Attribute          `json:"availabilityAttributes:omitempty"`
	BillingType                     *string                           `json:"billingType:omitempty"`
	Bundle                          []Product_Item_Bundles            `json:"bundle:omitempty"`
	BundleCount                     *uint                             `json:"bundleCount:omitempty"`
	Capacity                        *float64                          `json:"capacity:omitempty"`
	CapacityRestrictedProductFlag   *bool                             `json:"capacityRestrictedProductFlag:omitempty"`
	Categories                      []Product_Item_Category           `json:"categories:omitempty"`
	CategoryCount                   *uint                             `json:"categoryCount:omitempty"`
	ConfigurationTemplateCount      *uint                             `json:"configurationTemplateCount:omitempty"`
	ConfigurationTemplates          []Configuration_Template          `json:"configurationTemplates:omitempty"`
	ConflictCount                   *uint                             `json:"conflictCount:omitempty"`
	Conflicts                       []Product_Item_Resource_Conflict  `json:"conflicts:omitempty"`
	CoreRestrictedItemFlag          *bool                             `json:"coreRestrictedItemFlag:omitempty"`
	Description                     *string                           `json:"description:omitempty"`
	DowngradeItem                   *Product_Item                     `json:"downgradeItem:omitempty"`
	DowngradeItemCount              *uint                             `json:"downgradeItemCount:omitempty"`
	DowngradeItems                  []Product_Item                    `json:"downgradeItems:omitempty"`
	GlobalCategoryConflictCount     *uint                             `json:"globalCategoryConflictCount:omitempty"`
	GlobalCategoryConflicts         []Product_Item_Resource_Conflict  `json:"globalCategoryConflicts:omitempty"`
	HardwareGenericComponentModel   *Hardware_Component_Model_Generic `json:"hardwareGenericComponentModel:omitempty"`
	HideFromPortalFlag              *bool                             `json:"hideFromPortalFlag:omitempty"`
	Id                              *int                              `json:"id:omitempty"`
	Inventory                       []Product_Package_Inventory       `json:"inventory:omitempty"`
	InventoryCount                  *uint                             `json:"inventoryCount:omitempty"`
	IsEngineeredServerProduct       *bool                             `json:"isEngineeredServerProduct:omitempty"`
	ItemCategory                    *Product_Item_Category            `json:"itemCategory:omitempty"`
	ItemTaxCategoryId               *int                              `json:"itemTaxCategoryId:omitempty"`
	KeyName                         *string                           `json:"keyName:omitempty"`
	LocationConflictCount           *uint                             `json:"locationConflictCount:omitempty"`
	LocationConflicts               []Product_Item_Resource_Conflict  `json:"locationConflicts:omitempty"`
	LongDescription                 *string                           `json:"longDescription:omitempty"`
	ObjectStorageItemFlag           *bool                             `json:"objectStorageItemFlag:omitempty"`
	PackageCount                    *uint                             `json:"packageCount:omitempty"`
	Packages                        []Product_Package                 `json:"packages:omitempty"`
	PhysicalCoreCapacity            *string                           `json:"physicalCoreCapacity:omitempty"`
	PresaleEventCount               *uint                             `json:"presaleEventCount:omitempty"`
	PresaleEvents                   []Sales_Presale_Event             `json:"presaleEvents:omitempty"`
	PriceCount                      *uint                             `json:"priceCount:omitempty"`
	Prices                          []Product_Item_Price              `json:"prices:omitempty"`
	RequirementCount                *uint                             `json:"requirementCount:omitempty"`
	Requirements                    []Product_Item_Requirement        `json:"requirements:omitempty"`
	SoftwareDescription             *Software_Description             `json:"softwareDescription:omitempty"`
	SoftwareDescriptionId           *int                              `json:"softwareDescriptionId:omitempty"`
	TaxCategory                     *Product_Item_Tax_Category        `json:"taxCategory:omitempty"`
	ThirdPartyPolicyAssignmentCount *uint                             `json:"thirdPartyPolicyAssignmentCount:omitempty"`
	ThirdPartyPolicyAssignments     []Product_Item_Policy_Assignment  `json:"thirdPartyPolicyAssignments:omitempty"`
	ThirdPartySupportVendor         *string                           `json:"thirdPartySupportVendor:omitempty"`
	TotalPhysicalCoreCapacity       *int                              `json:"totalPhysicalCoreCapacity:omitempty"`
	TotalPhysicalCoreCount          *int                              `json:"totalPhysicalCoreCount:omitempty"`
	TotalProcessorCapacity          *int                              `json:"totalProcessorCapacity:omitempty"`
	Units                           *string                           `json:"units:omitempty"`
	UpgradeItem                     *Product_Item                     `json:"upgradeItem:omitempty"`
	UpgradeItemCount                *uint                             `json:"upgradeItemCount:omitempty"`
	UpgradeItemId                   *int                              `json:"upgradeItemId:omitempty"`
	UpgradeItems                    []Product_Item                    `json:"upgradeItems:omitempty"`
}

type Product_Item_Attribute struct {
	Entity

	AttributeType        *Product_Item_Attribute_Type `json:"attributeType:omitempty"`
	AttributeTypeKeyName *string                      `json:"attributeTypeKeyName:omitempty"`
	Id                   *int                         `json:"id:omitempty"`
	Item                 *Product_Item                `json:"item:omitempty"`
	ItemAttributeTypeId  *int                         `json:"itemAttributeTypeId:omitempty"`
	ItemId               *int                         `json:"itemId:omitempty"`
	Value                *string                      `json:"value:omitempty"`
}

type Product_Item_Attribute_Type struct {
	Entity

	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Product_Item_Billing_Type struct {
	Entity

	Name *string `json:"name:omitempty"`
}

type Product_Item_Bundles struct {
	Entity

	BundleItem   *Product_Item          `json:"bundleItem:omitempty"`
	BundleItemId *int                   `json:"bundleItemId:omitempty"`
	Category     *Product_Item_Category `json:"category:omitempty"`
	Id           *int                   `json:"id:omitempty"`
	ItemPrice    *Product_Item_Price    `json:"itemPrice:omitempty"`
	ItemPriceId  *int                   `json:"itemPriceId:omitempty"`
}

type Product_Item_Category struct {
	Entity

	BillingItemCount          *uint                                     `json:"billingItemCount:omitempty"`
	BillingItems              []Billing_Item                            `json:"billingItems:omitempty"`
	CategoryCode              *string                                   `json:"categoryCode:omitempty"`
	Group                     *Product_Item_Category_Group              `json:"group:omitempty"`
	GroupCount                *uint                                     `json:"groupCount:omitempty"`
	Groups                    []Product_Package_Item_Category_Group     `json:"groups:omitempty"`
	Id                        *int                                      `json:"id:omitempty"`
	Name                      *string                                   `json:"name:omitempty"`
	OrderOptionCount          *uint                                     `json:"orderOptionCount:omitempty"`
	OrderOptions              []Product_Item_Category_Order_Option_Type `json:"orderOptions:omitempty"`
	PackageConfigurationCount *uint                                     `json:"packageConfigurationCount:omitempty"`
	PackageConfigurations     []Product_Package_Order_Configuration     `json:"packageConfigurations:omitempty"`
	PresetConfigurationCount  *uint                                     `json:"presetConfigurationCount:omitempty"`
	PresetConfigurations      []Product_Package_Preset_Configuration    `json:"presetConfigurations:omitempty"`
	QuantityLimit             *int                                      `json:"quantityLimit:omitempty"`
	QuestionCount             *uint                                     `json:"questionCount:omitempty"`
	QuestionReferenceCount    *uint                                     `json:"questionReferenceCount:omitempty"`
	QuestionReferences        []Product_Item_Category_Question_Xref     `json:"questionReferences:omitempty"`
	Questions                 []Product_Item_Category_Question          `json:"questions:omitempty"`
}

type Product_Item_Category_Group struct {
	Entity

	Id   *int    `json:"id:omitempty"`
	Name *string `json:"name:omitempty"`
}

type Product_Item_Category_Order_Option_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	Keyname     *string `json:"keyname:omitempty"`
	Name        *string `json:"name:omitempty"`
	Value       *string `json:"value:omitempty"`
}

type Product_Item_Category_Question struct {
	Entity

	AnswerValueExpression      *string                                    `json:"answerValueExpression:omitempty"`
	Description                *string                                    `json:"description:omitempty"`
	FieldType                  *Product_Item_Category_Question_Field_Type `json:"fieldType:omitempty"`
	FieldTypeId                *int                                       `json:"fieldTypeId:omitempty"`
	Id                         *int                                       `json:"id:omitempty"`
	ItemCategoryReferenceCount *uint                                      `json:"itemCategoryReferenceCount:omitempty"`
	ItemCategoryReferences     []Product_Item_Category_Question_Xref      `json:"itemCategoryReferences:omitempty"`
	KeyName                    *string                                    `json:"keyName:omitempty"`
	Question                   *string                                    `json:"question:omitempty"`
	ValueExample               *string                                    `json:"valueExample:omitempty"`
}

type Product_Item_Category_Question_Field_Type struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Product_Item_Category_Question_Xref struct {
	Entity

	Id             *int                            `json:"id:omitempty"`
	ItemCategory   *Product_Item_Category          `json:"itemCategory:omitempty"`
	ItemCategoryId *int                            `json:"itemCategoryId:omitempty"`
	LocationId     *int                            `json:"locationId:omitempty"`
	Question       *Product_Item_Category_Question `json:"question:omitempty"`
	QuestionId     *int                            `json:"questionId:omitempty"`
	Required       *bool                           `json:"required:omitempty"`
}

type Product_Item_Link_ThePlanet struct {
	Entity

	Item            *Product_Item     `json:"item:omitempty"`
	ServiceProvider *Service_Provider `json:"serviceProvider:omitempty"`
}

type Product_Item_Policy_Assignment struct {
	Entity

	Id         *int          `json:"id:omitempty"`
	PolicyName *string       `json:"policyName:omitempty"`
	Product    *Product_Item `json:"product:omitempty"`
	ProductId  *int          `json:"productId:omitempty"`
}

type Product_Item_Price struct {
	Entity

	AccountRestrictionCount    *uint                                     `json:"accountRestrictionCount:omitempty"`
	AccountRestrictions        []Product_Item_Price_Account_Restriction  `json:"accountRestrictions:omitempty"`
	AttributeCount             *uint                                     `json:"attributeCount:omitempty"`
	Attributes                 []Product_Item_Price_Attribute            `json:"attributes:omitempty"`
	BigDataOsJournalDiskFlag   *bool                                     `json:"bigDataOsJournalDiskFlag:omitempty"`
	BundleReferenceCount       *uint                                     `json:"bundleReferenceCount:omitempty"`
	BundleReferences           []Product_Item_Bundles                    `json:"bundleReferences:omitempty"`
	CapacityRestrictionMaximum *string                                   `json:"capacityRestrictionMaximum:omitempty"`
	CapacityRestrictionMinimum *string                                   `json:"capacityRestrictionMinimum:omitempty"`
	CapacityRestrictionType    *string                                   `json:"capacityRestrictionType:omitempty"`
	Categories                 []Product_Item_Category                   `json:"categories:omitempty"`
	CategoryCount              *uint                                     `json:"categoryCount:omitempty"`
	CurrentPriceFlag           *bool                                     `json:"currentPriceFlag:omitempty"`
	DefinedSoftwareLicenseFlag *bool                                     `json:"definedSoftwareLicenseFlag:omitempty"`
	HourlyRecurringFee         *float64                                  `json:"hourlyRecurringFee:omitempty"`
	Id                         *int                                      `json:"id:omitempty"`
	Inventory                  []Product_Package_Inventory               `json:"inventory:omitempty"`
	InventoryCount             *uint                                     `json:"inventoryCount:omitempty"`
	Item                       *Product_Item                             `json:"item:omitempty"`
	ItemId                     *int                                      `json:"itemId:omitempty"`
	LaborFee                   *float64                                  `json:"laborFee:omitempty"`
	LocationGroupId            *int                                      `json:"locationGroupId:omitempty"`
	OnSaleFlag                 *bool                                     `json:"onSaleFlag:omitempty"`
	OneTimeFee                 *float64                                  `json:"oneTimeFee:omitempty"`
	OneTimeFeeTax              *float64                                  `json:"oneTimeFeeTax:omitempty"`
	OrderOptions               []Product_Item_Category_Order_Option_Type `json:"orderOptions:omitempty"`
	OrderPremiumCount          *uint                                     `json:"orderPremiumCount:omitempty"`
	OrderPremiums              []Product_Item_Price_Premium              `json:"orderPremiums:omitempty"`
	PackageCount               *uint                                     `json:"packageCount:omitempty"`
	PackageReferenceCount      *uint                                     `json:"packageReferenceCount:omitempty"`
	PackageReferences          []Product_Package_Item_Prices             `json:"packageReferences:omitempty"`
	Packages                   []Product_Package                         `json:"packages:omitempty"`
	PresetConfigurationCount   *uint                                     `json:"presetConfigurationCount:omitempty"`
	PresetConfigurations       []Product_Package_Preset_Configuration    `json:"presetConfigurations:omitempty"`
	PricingLocationGroup       *Location_Group_Pricing                   `json:"pricingLocationGroup:omitempty"`
	ProratedRecurringFee       *float64                                  `json:"proratedRecurringFee:omitempty"`
	ProratedRecurringFeeTax    *float64                                  `json:"proratedRecurringFeeTax:omitempty"`
	Quantity                   *int                                      `json:"quantity:omitempty"`
	RecurringFee               *float64                                  `json:"recurringFee:omitempty"`
	RecurringFeeTax            *float64                                  `json:"recurringFeeTax:omitempty"`
	RequiredCoreCount          *int                                      `json:"requiredCoreCount:omitempty"`
	SetupFee                   *float64                                  `json:"setupFee:omitempty"`
	Sort                       *int                                      `json:"sort:omitempty"`
	UsageRate                  *float64                                  `json:"usageRate:omitempty"`
}

type Product_Item_Price_Account_Restriction struct {
	Entity

	Account     *Account            `json:"account:omitempty"`
	AccountId   *int                `json:"accountId:omitempty"`
	Id          *int                `json:"id:omitempty"`
	ItemPrice   *Product_Item_Price `json:"itemPrice:omitempty"`
	ItemPriceId *int                `json:"itemPriceId:omitempty"`
}

type Product_Item_Price_Attribute struct {
	Entity

	Id                       *int                               `json:"id:omitempty"`
	ItemPrice                *Product_Item_Price                `json:"itemPrice:omitempty"`
	ItemPriceAttributeType   *Product_Item_Price_Attribute_Type `json:"itemPriceAttributeType:omitempty"`
	ItemPriceAttributeTypeId *int                               `json:"itemPriceAttributeTypeId:omitempty"`
	ItemPriceId              *int                               `json:"itemPriceId:omitempty"`
	Value                    *string                            `json:"value:omitempty"`
}

type Product_Item_Price_Attribute_Type struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	Keyname *string `json:"keyname:omitempty"`
}

type Product_Item_Price_Premium struct {
	Entity

	HourlyModifier  *float64            `json:"hourlyModifier:omitempty"`
	ItemPrice       *Product_Item_Price `json:"itemPrice:omitempty"`
	ItemPriceId     *int                `json:"itemPriceId:omitempty"`
	Location        *Location           `json:"location:omitempty"`
	LocationId      *int                `json:"locationId:omitempty"`
	MonthlyModifier *float64            `json:"monthlyModifier:omitempty"`
	Package         *Product_Package    `json:"package:omitempty"`
	PackageId       *int                `json:"packageId:omitempty"`
}

type Product_Item_Requirement struct {
	Entity

	Id             *int          `json:"id:omitempty"`
	Item           *Product_Item `json:"item:omitempty"`
	ItemId         *int          `json:"itemId:omitempty"`
	Message        *string       `json:"message:omitempty"`
	Product        *Product_Item `json:"product:omitempty"`
	RequiredItemId *int          `json:"requiredItemId:omitempty"`
}

type Product_Item_Resource_Conflict struct {
	Entity

	Item            *Product_Item    `json:"item:omitempty"`
	ItemId          *int             `json:"itemId:omitempty"`
	Message         *string          `json:"message:omitempty"`
	Package         *Product_Package `json:"package:omitempty"`
	PackageId       *int             `json:"packageId:omitempty"`
	ResourceTableId *int             `json:"resourceTableId:omitempty"`
}

type Product_Item_Resource_Conflict_Item struct {
	Product_Item_Resource_Conflict

	Resource *Product_Item `json:"resource:omitempty"`
}

type Product_Item_Resource_Conflict_Item_Category struct {
	Product_Item_Resource_Conflict

	Resource *Product_Item_Category `json:"resource:omitempty"`
}

type Product_Item_Resource_Conflict_Location struct {
	Product_Item_Resource_Conflict

	Resource *Location `json:"resource:omitempty"`
}

type Product_Item_Tax_Category struct {
	Entity

	Id         *int           `json:"id:omitempty"`
	ItemCount  *uint          `json:"itemCount:omitempty"`
	Items      []Product_Item `json:"items:omitempty"`
	Name       *string        `json:"name:omitempty"`
	StatusFlag *int           `json:"statusFlag:omitempty"`
}

type Product_Order struct {
	Entity
}

type Product_Package struct {
	Entity

	AccountRestrictedCategories     []Product_Item_Category               `json:"accountRestrictedCategories:omitempty"`
	AccountRestrictedCategoryCount  *uint                                 `json:"accountRestrictedCategoryCount:omitempty"`
	AccountRestrictedPricesFlag     *bool                                 `json:"accountRestrictedPricesFlag:omitempty"`
	ActivePresetCount               *uint                                 `json:"activePresetCount:omitempty"`
	ActivePresets                   []Product_Package_Preset              `json:"activePresets:omitempty"`
	ActiveRamItemCount              *uint                                 `json:"activeRamItemCount:omitempty"`
	ActiveRamItems                  []Product_Item                        `json:"activeRamItems:omitempty"`
	ActiveServerItemCount           *uint                                 `json:"activeServerItemCount:omitempty"`
	ActiveServerItems               []Product_Item                        `json:"activeServerItems:omitempty"`
	ActiveSoftwareItemCount         *uint                                 `json:"activeSoftwareItemCount:omitempty"`
	ActiveSoftwareItems             []Product_Item                        `json:"activeSoftwareItems:omitempty"`
	ActiveUsagePriceCount           *uint                                 `json:"activeUsagePriceCount:omitempty"`
	ActiveUsagePrices               []Product_Item_Price                  `json:"activeUsagePrices:omitempty"`
	AdditionalServiceFlag           *bool                                 `json:"additionalServiceFlag:omitempty"`
	AttributeCount                  *uint                                 `json:"attributeCount:omitempty"`
	Attributes                      []Product_Package_Attribute           `json:"attributes:omitempty"`
	AvailableLocationCount          *uint                                 `json:"availableLocationCount:omitempty"`
	AvailableLocations              []Product_Package_Locations           `json:"availableLocations:omitempty"`
	AvailableStorageUnits           *uint                                 `json:"availableStorageUnits:omitempty"`
	Categories                      []Product_Item_Category               `json:"categories:omitempty"`
	Configuration                   []Product_Package_Order_Configuration `json:"configuration:omitempty"`
	ConfigurationCount              *uint                                 `json:"configurationCount:omitempty"`
	DefaultRamItemCount             *uint                                 `json:"defaultRamItemCount:omitempty"`
	DefaultRamItems                 []Product_Item                        `json:"defaultRamItems:omitempty"`
	DeploymentCount                 *uint                                 `json:"deploymentCount:omitempty"`
	DeploymentNodeType              *string                               `json:"deploymentNodeType:omitempty"`
	DeploymentPackageCount          *uint                                 `json:"deploymentPackageCount:omitempty"`
	DeploymentPackages              []Product_Package                     `json:"deploymentPackages:omitempty"`
	DeploymentType                  *string                               `json:"deploymentType:omitempty"`
	Deployments                     []Product_Package                     `json:"deployments:omitempty"`
	Description                     *string                               `json:"description:omitempty"`
	DisallowCustomDiskPartitions    *bool                                 `json:"disallowCustomDiskPartitions:omitempty"`
	FirstOrderStep                  *Product_Package_Order_Step           `json:"firstOrderStep:omitempty"`
	FirstOrderStepId                *int                                  `json:"firstOrderStepId:omitempty"`
	GatewayApplianceFlag            *bool                                 `json:"gatewayApplianceFlag:omitempty"`
	GpuFlag                         *bool                                 `json:"gpuFlag:omitempty"`
	HourlyBillingAvailableFlag      *bool                                 `json:"hourlyBillingAvailableFlag:omitempty"`
	Id                              *int                                  `json:"id:omitempty"`
	IsActive                        *int                                  `json:"isActive:omitempty"`
	ItemConflictCount               *uint                                 `json:"itemConflictCount:omitempty"`
	ItemConflicts                   []Product_Item_Resource_Conflict      `json:"itemConflicts:omitempty"`
	ItemCount                       *uint                                 `json:"itemCount:omitempty"`
	ItemLocationConflictCount       *uint                                 `json:"itemLocationConflictCount:omitempty"`
	ItemLocationConflicts           []Product_Item_Resource_Conflict      `json:"itemLocationConflicts:omitempty"`
	ItemPriceCount                  *uint                                 `json:"itemPriceCount:omitempty"`
	ItemPriceReferenceCount         *uint                                 `json:"itemPriceReferenceCount:omitempty"`
	ItemPriceReferences             []Product_Package_Item_Prices         `json:"itemPriceReferences:omitempty"`
	ItemPrices                      []Product_Item_Price                  `json:"itemPrices:omitempty"`
	Items                           []Product_Item                        `json:"items:omitempty"`
	KeyName                         *string                               `json:"keyName:omitempty"`
	LocationCount                   *uint                                 `json:"locationCount:omitempty"`
	Locations                       []Location                            `json:"locations:omitempty"`
	LowestServerPrice               *Product_Item_Price                   `json:"lowestServerPrice:omitempty"`
	MaximumPortSpeed                *uint                                 `json:"maximumPortSpeed:omitempty"`
	MinimumPortSpeed                *uint                                 `json:"minimumPortSpeed:omitempty"`
	MongoDbEngineeredFlag           *bool                                 `json:"mongoDbEngineeredFlag:omitempty"`
	Name                            *string                               `json:"name:omitempty"`
	OrderPremiumCount               *uint                                 `json:"orderPremiumCount:omitempty"`
	OrderPremiums                   []Product_Item_Price_Premium          `json:"orderPremiums:omitempty"`
	PreconfiguredFlag               *bool                                 `json:"preconfiguredFlag:omitempty"`
	PresetConfigurationRequiredFlag *bool                                 `json:"presetConfigurationRequiredFlag:omitempty"`
	PreventVlanSelectionFlag        *bool                                 `json:"preventVlanSelectionFlag:omitempty"`
	PrivateHostedCloudPackageFlag   *bool                                 `json:"privateHostedCloudPackageFlag:omitempty"`
	PrivateHostedCloudPackageType   *string                               `json:"privateHostedCloudPackageType:omitempty"`
	PrivateNetworkOnlyFlag          *bool                                 `json:"privateNetworkOnlyFlag:omitempty"`
	QuantaStorPackageFlag           *bool                                 `json:"quantaStorPackageFlag:omitempty"`
	RaidDiskRestrictionFlag         *bool                                 `json:"raidDiskRestrictionFlag:omitempty"`
	RedundantPowerFlag              *bool                                 `json:"redundantPowerFlag:omitempty"`
	RegionCount                     *uint                                 `json:"regionCount:omitempty"`
	Regions                         []Location_Region                     `json:"regions:omitempty"`
	ResourceGroupTemplate           *Resource_Group_Template              `json:"resourceGroupTemplate:omitempty"`
	SubDescription                  *string                               `json:"subDescription:omitempty"`
	TopLevelItemCategoryCode        *string                               `json:"topLevelItemCategoryCode:omitempty"`
	Type                            *Product_Package_Type                 `json:"type:omitempty"`
	UnitSize                        *int                                  `json:"unitSize:omitempty"`
}

type Product_Package_Attribute struct {
	Entity

	AttributeType *Product_Package_Attribute_Type `json:"attributeType:omitempty"`
	Package       *Product_Package                `json:"package:omitempty"`
	Value         *string                         `json:"value:omitempty"`
}

type Product_Package_Attribute_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Product_Package_Inventory struct {
	Entity

	AvailableInventoryCount *int             `json:"availableInventoryCount:omitempty"`
	Item                    *Product_Item    `json:"item:omitempty"`
	ItemId                  *int             `json:"itemId:omitempty"`
	Location                *Location        `json:"location:omitempty"`
	LocationId              *int             `json:"locationId:omitempty"`
	ModifyDate              *time.Time       `json:"modifyDate:omitempty"`
	OverstockFlag           *int             `json:"overstockFlag:omitempty"`
	Package                 *Product_Package `json:"package:omitempty"`
	PackageId               *int             `json:"packageId:omitempty"`
}

type Product_Package_Item_Category_Group struct {
	Entity

	Category       *Product_Item_Category `json:"category:omitempty"`
	ItemCategoryId *int                   `json:"itemCategoryId:omitempty"`
	Package        *Product_Package       `json:"package:omitempty"`
	PackageId      *int                   `json:"packageId:omitempty"`
	PriceCount     *uint                  `json:"priceCount:omitempty"`
	Prices         []Product_Item_Price   `json:"prices:omitempty"`
	Sort           *int                   `json:"sort:omitempty"`
	Title          *string                `json:"title:omitempty"`
}

type Product_Package_Item_Prices struct {
	Entity

	Id          *int                `json:"id:omitempty"`
	ItemPrice   *Product_Item_Price `json:"itemPrice:omitempty"`
	ItemPriceId *int                `json:"itemPriceId:omitempty"`
	Package     *Product_Package    `json:"package:omitempty"`
	PackageId   *int                `json:"packageId:omitempty"`
}

type Product_Package_Items struct {
	Entity

	Id        *string          `json:"id:omitempty"`
	Item      *Product_Item    `json:"item:omitempty"`
	ItemId    *int             `json:"itemId:omitempty"`
	Package   *Product_Package `json:"package:omitempty"`
	PackageId *int             `json:"packageId:omitempty"`
}

type Product_Package_Locations struct {
	Entity

	DeliveryTimeInformation *string          `json:"deliveryTimeInformation:omitempty"`
	IsAvailable             *int             `json:"isAvailable:omitempty"`
	Location                *Location        `json:"location:omitempty"`
	LocationId              *int             `json:"locationId:omitempty"`
	Package                 *Product_Package `json:"package:omitempty"`
	PackageId               *int             `json:"packageId:omitempty"`
}

type Product_Package_Order_Configuration struct {
	Entity

	ErrorMessage   *string                     `json:"errorMessage:omitempty"`
	Id             *int                        `json:"id:omitempty"`
	IsRequired     *int                        `json:"isRequired:omitempty"`
	ItemCategory   *Product_Item_Category      `json:"itemCategory:omitempty"`
	ItemCategoryId *int                        `json:"itemCategoryId:omitempty"`
	OrderStepId    *int                        `json:"orderStepId:omitempty"`
	Package        *Product_Package            `json:"package:omitempty"`
	PackageId      *int                        `json:"packageId:omitempty"`
	Sort           *int                        `json:"sort:omitempty"`
	Step           *Product_Package_Order_Step `json:"step:omitempty"`
}

type Product_Package_Order_Step struct {
	Entity

	Id                         *int                              `json:"id:omitempty"`
	InclusivePreviousStepCount *uint                             `json:"inclusivePreviousStepCount:omitempty"`
	InclusivePreviousSteps     []Product_Package_Order_Step_Next `json:"inclusivePreviousSteps:omitempty"`
	NextStepCount              *uint                             `json:"nextStepCount:omitempty"`
	NextSteps                  []Product_Package_Order_Step_Next `json:"nextSteps:omitempty"`
	PreviousStepCount          *uint                             `json:"previousStepCount:omitempty"`
	PreviousSteps              []Product_Package_Order_Step_Next `json:"previousSteps:omitempty"`
	Step                       *string                           `json:"step:omitempty"`
}

type Product_Package_Order_Step_Next struct {
	Entity

	Id              *int                        `json:"id:omitempty"`
	NextOrderStepId *int                        `json:"nextOrderStepId:omitempty"`
	OrderStepId     *int                        `json:"orderStepId:omitempty"`
	Step            *Product_Package_Order_Step `json:"step:omitempty"`
}

type Product_Package_Preset struct {
	Entity

	AvailableStorageUnits          *uint                                        `json:"availableStorageUnits:omitempty"`
	Categories                     []Product_Item_Category                      `json:"categories:omitempty"`
	CategoryCount                  *uint                                        `json:"categoryCount:omitempty"`
	Configuration                  []Product_Package_Preset_Configuration       `json:"configuration:omitempty"`
	ConfigurationCount             *uint                                        `json:"configurationCount:omitempty"`
	Description                    *string                                      `json:"description:omitempty"`
	FixedConfigurationFlag         *bool                                        `json:"fixedConfigurationFlag:omitempty"`
	Id                             *int                                         `json:"id:omitempty"`
	IsActive                       *string                                      `json:"isActive:omitempty"`
	KeyName                        *string                                      `json:"keyName:omitempty"`
	LowestPresetServerPrice        *Product_Item_Price                          `json:"lowestPresetServerPrice:omitempty"`
	Name                           *string                                      `json:"name:omitempty"`
	Package                        *Product_Package                             `json:"package:omitempty"`
	PackageConfiguration           []Product_Package_Order_Configuration        `json:"packageConfiguration:omitempty"`
	PackageConfigurationCount      *uint                                        `json:"packageConfigurationCount:omitempty"`
	PackageId                      *int                                         `json:"packageId:omitempty"`
	PriceCount                     *uint                                        `json:"priceCount:omitempty"`
	Prices                         []Product_Item_Price                         `json:"prices:omitempty"`
	StorageGroupTemplateArrayCount *uint                                        `json:"storageGroupTemplateArrayCount:omitempty"`
	StorageGroupTemplateArrays     []Configuration_Storage_Group_Template_Group `json:"storageGroupTemplateArrays:omitempty"`
	TotalMinimumHourlyFee          *float64                                     `json:"totalMinimumHourlyFee:omitempty"`
	TotalMinimumRecurringFee       *float64                                     `json:"totalMinimumRecurringFee:omitempty"`
}

type Product_Package_Preset_Attribute struct {
	Entity

	AttributeType   *Product_Package_Preset_Attribute_Type `json:"attributeType:omitempty"`
	AttributeTypeId *int                                   `json:"attributeTypeId:omitempty"`
	Id              *int                                   `json:"id:omitempty"`
	Preset          *Product_Package_Preset                `json:"preset:omitempty"`
	PresetId        *int                                   `json:"presetId:omitempty"`
	Value           *string                                `json:"value:omitempty"`
}

type Product_Package_Preset_Attribute_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Product_Package_Preset_Configuration struct {
	Entity

	Category      *Product_Item_Category  `json:"category:omitempty"`
	PackagePreset *Product_Package_Preset `json:"packagePreset:omitempty"`
	Price         *Product_Item_Price     `json:"price:omitempty"`
}

type Product_Package_Server struct {
	Entity

	Catalog                *Product_Catalog        `json:"catalog:omitempty"`
	CatalogId              *int                    `json:"catalogId:omitempty"`
	Datacenters            *string                 `json:"datacenters:omitempty"`
	DefaultRamCapacity     *float64                `json:"defaultRamCapacity:omitempty"`
	DualPathNetworkFlag    *bool                   `json:"dualPathNetworkFlag:omitempty"`
	GpuFlag                *bool                   `json:"gpuFlag:omitempty"`
	HourlyBillingFlag      *bool                   `json:"hourlyBillingFlag:omitempty"`
	Id                     *int                    `json:"id:omitempty"`
	Item                   *Product_Item           `json:"item:omitempty"`
	ItemId                 *int                    `json:"itemId:omitempty"`
	ItemPrice              *Product_Item_Price     `json:"itemPrice:omitempty"`
	ItemPriceId            *int                    `json:"itemPriceId:omitempty"`
	MaximumDriveCount      *int                    `json:"maximumDriveCount:omitempty"`
	MaximumPortSpeed       *float64                `json:"maximumPortSpeed:omitempty"`
	MaximumRamCapacity     *float64                `json:"maximumRamCapacity:omitempty"`
	MinimumPortSpeed       *float64                `json:"minimumPortSpeed:omitempty"`
	OutletFlag             *bool                   `json:"outletFlag:omitempty"`
	Package                *Product_Package        `json:"package:omitempty"`
	PackageId              *int                    `json:"packageId:omitempty"`
	PackageType            *string                 `json:"packageType:omitempty"`
	PowerServerFlag        *bool                   `json:"powerServerFlag:omitempty"`
	Preset                 *Product_Package_Preset `json:"preset:omitempty"`
	PresetId               *int                    `json:"presetId:omitempty"`
	PrivateNetworkOnlyFlag *bool                   `json:"privateNetworkOnlyFlag:omitempty"`
	ProcessorBusSpeed      *string                 `json:"processorBusSpeed:omitempty"`
	ProcessorCache         *string                 `json:"processorCache:omitempty"`
	ProcessorCores         *int                    `json:"processorCores:omitempty"`
	ProcessorCount         *int                    `json:"processorCount:omitempty"`
	ProcessorManufacturer  *string                 `json:"processorManufacturer:omitempty"`
	ProcessorModel         *string                 `json:"processorModel:omitempty"`
	ProcessorName          *string                 `json:"processorName:omitempty"`
	ProcessorSpeed         *string                 `json:"processorSpeed:omitempty"`
	ProductName            *string                 `json:"productName:omitempty"`
	RedundantPowerFlag     *bool                   `json:"redundantPowerFlag:omitempty"`
	SapCertifiedServerFlag *bool                   `json:"sapCertifiedServerFlag:omitempty"`
	StartingHourlyPrice    *float64                `json:"startingHourlyPrice:omitempty"`
	StartingMonthlyPrice   *float64                `json:"startingMonthlyPrice:omitempty"`
	TotalCoreCount         *int                    `json:"totalCoreCount:omitempty"`
	TxtTpmFlag             *bool                   `json:"txtTpmFlag:omitempty"`
	UnitSize               *int                    `json:"unitSize:omitempty"`
}

type Product_Package_Server_Option struct {
	Entity

	CatalogId   *int    `json:"catalogId:omitempty"`
	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	Type        *string `json:"type:omitempty"`
	Value       *string `json:"value:omitempty"`
}

type Product_Package_Type struct {
	Entity

	Id           *int              `json:"id:omitempty"`
	KeyName      *string           `json:"keyName:omitempty"`
	Name         *string           `json:"name:omitempty"`
	PackageCount *uint             `json:"packageCount:omitempty"`
	Packages     []Product_Package `json:"packages:omitempty"`
}

type Product_Upgrade_Request struct {
	Entity

	Account                 *Account                        `json:"account:omitempty"`
	AccountId               *int                            `json:"accountId:omitempty"`
	CompletedFlag           *bool                           `json:"completedFlag:omitempty"`
	CreateDate              *time.Time                      `json:"createDate:omitempty"`
	EmployeeId              *int                            `json:"employeeId:omitempty"`
	GuestId                 *int                            `json:"guestId:omitempty"`
	HardwareId              *int                            `json:"hardwareId:omitempty"`
	Id                      *int                            `json:"id:omitempty"`
	Invoice                 *Billing_Invoice                `json:"invoice:omitempty"`
	MaintenanceStartTimeUtc *time.Time                      `json:"maintenanceStartTimeUtc:omitempty"`
	ModifyDate              *time.Time                      `json:"modifyDate:omitempty"`
	Order                   *Billing_Order                  `json:"order:omitempty"`
	OrderId                 *int                            `json:"orderId:omitempty"`
	OrderTotal              *float64                        `json:"orderTotal:omitempty"`
	ProratedTotal           *float64                        `json:"proratedTotal:omitempty"`
	Server                  *Hardware                       `json:"server:omitempty"`
	Status                  *Product_Upgrade_Request_Status `json:"status:omitempty"`
	StatusId                *int                            `json:"statusId:omitempty"`
	Ticket                  *Ticket                         `json:"ticket:omitempty"`
	TicketId                *int                            `json:"ticketId:omitempty"`
	User                    *User_Customer                  `json:"user:omitempty"`
	UserId                  *int                            `json:"userId:omitempty"`
	VirtualGuest            *Virtual_Guest                  `json:"virtualGuest:omitempty"`
}

type Product_Upgrade_Request_Status struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	Name        *string `json:"name:omitempty"`
	StatusCode  *string `json:"statusCode:omitempty"`
}

type Provisioning_Hook struct {
	Entity

	Account    *Account                `json:"account:omitempty"`
	AccountId  *int                    `json:"accountId:omitempty"`
	CreateDate *time.Time              `json:"createDate:omitempty"`
	HookType   *Provisioning_Hook_Type `json:"hookType:omitempty"`
	Id         *int                    `json:"id:omitempty"`
	ModifyDate *time.Time              `json:"modifyDate:omitempty"`
	Name       *string                 `json:"name:omitempty"`
	TypeId     *int                    `json:"typeId:omitempty"`
	Uri        *string                 `json:"uri:omitempty"`
}

type Provisioning_Hook_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Provisioning_Maintenance_Classification struct {
	Entity

	Id                *int                                                    `json:"id:omitempty"`
	ItemCategories    []Provisioning_Maintenance_Classification_Item_Category `json:"itemCategories:omitempty"`
	ItemCategoryCount *uint                                                   `json:"itemCategoryCount:omitempty"`
	Slots             *int                                                    `json:"slots:omitempty"`
	Type              *string                                                 `json:"type:omitempty"`
}

type Provisioning_Maintenance_Classification_Item_Category struct {
	Entity

	ItemCategoryId              *int                                     `json:"itemCategoryId:omitempty"`
	MaintenanceClassification   *Provisioning_Maintenance_Classification `json:"maintenanceClassification:omitempty"`
	MaintenanceClassificationId *int                                     `json:"maintenanceClassificationId:omitempty"`
}

type Provisioning_Maintenance_Slots struct {
	Entity

	AvailableSlots *int `json:"availableSlots:omitempty"`
}

type Provisioning_Maintenance_Ticket struct {
	Entity

	AvailableSlots   *Provisioning_Maintenance_Slots          `json:"availableSlots:omitempty"`
	MaintClassId     *int                                     `json:"maintClassId:omitempty"`
	MaintWindowId    *int                                     `json:"maintWindowId:omitempty"`
	MaintenanceClass *Provisioning_Maintenance_Classification `json:"maintenanceClass:omitempty"`
	MaintenanceDate  *time.Time                               `json:"maintenanceDate:omitempty"`
	Ticket           *Ticket                                  `json:"ticket:omitempty"`
	TicketId         *int                                     `json:"ticketId:omitempty"`
}

type Provisioning_Maintenance_Window struct {
	Entity

	BeginDate  *time.Time `json:"beginDate:omitempty"`
	DayOfWeek  *int       `json:"dayOfWeek:omitempty"`
	EndDate    *time.Time `json:"endDate:omitempty"`
	Id         *int       `json:"id:omitempty"`
	LocationId *int       `json:"locationId:omitempty"`
	PortalTzId *int       `json:"portalTzId:omitempty"`
}

type Provisioning_Version1_Transaction struct {
	Entity

	Account                             *Account                                  `json:"account:omitempty"`
	CreateDate                          *time.Time                                `json:"createDate:omitempty"`
	ElapsedSeconds                      *int                                      `json:"elapsedSeconds:omitempty"`
	Guest                               *Virtual_Guest                            `json:"guest:omitempty"`
	GuestId                             *int                                      `json:"guestId:omitempty"`
	Hardware                            *Hardware                                 `json:"hardware:omitempty"`
	HardwareId                          *int                                      `json:"hardwareId:omitempty"`
	Id                                  *int                                      `json:"id:omitempty"`
	Loopback                            []Provisioning_Version1_Transaction       `json:"loopback:omitempty"`
	LoopbackCount                       *uint                                     `json:"loopbackCount:omitempty"`
	ModifyDate                          *time.Time                                `json:"modifyDate:omitempty"`
	PendingTransactionCount             *uint                                     `json:"pendingTransactionCount:omitempty"`
	PendingTransactions                 []Provisioning_Version1_Transaction       `json:"pendingTransactions:omitempty"`
	StatusChangeDate                    *time.Time                                `json:"statusChangeDate:omitempty"`
	TicketScheduledActionReference      []Ticket_Attachment                       `json:"ticketScheduledActionReference:omitempty"`
	TicketScheduledActionReferenceCount *uint                                     `json:"ticketScheduledActionReferenceCount:omitempty"`
	TransactionGroup                    *Provisioning_Version1_Transaction_Group  `json:"transactionGroup:omitempty"`
	TransactionStatus                   *Provisioning_Version1_Transaction_Status `json:"transactionStatus:omitempty"`
}

type Provisioning_Version1_Transaction_Group struct {
	Entity

	AverageTimeToComplete *float64 `json:"averageTimeToComplete:omitempty"`
	Name                  *string  `json:"name:omitempty"`
}

type Provisioning_Version1_Transaction_History struct {
	Entity

	FinishDate          *time.Time                                `json:"finishDate:omitempty"`
	Guest               *Virtual_Guest                            `json:"guest:omitempty"`
	GuestId             *int                                      `json:"guestId:omitempty"`
	Hardware            *Hardware                                 `json:"hardware:omitempty"`
	HardwareId          *int                                      `json:"hardwareId:omitempty"`
	HostId              *int                                      `json:"hostId:omitempty"`
	Id                  *int                                      `json:"id:omitempty"`
	StartDate           *time.Time                                `json:"startDate:omitempty"`
	Transaction         *Provisioning_Version1_Transaction        `json:"transaction:omitempty"`
	TransactionId       *int                                      `json:"transactionId:omitempty"`
	TransactionStatus   *Provisioning_Version1_Transaction_Status `json:"transactionStatus:omitempty"`
	TransactionStatusId *int                                      `json:"transactionStatusId:omitempty"`
}

type Provisioning_Version1_Transaction_Status struct {
	Entity

	AverageDuration              *float64                            `json:"averageDuration:omitempty"`
	FriendlyName                 *string                             `json:"friendlyName:omitempty"`
	Name                         *string                             `json:"name:omitempty"`
	NonCompletedTransactionCount *uint                               `json:"nonCompletedTransactionCount:omitempty"`
	NonCompletedTransactions     []Provisioning_Version1_Transaction `json:"nonCompletedTransactions:omitempty"`
}

type Provisioning_Version1_Transaction_SubnetMigration struct {
	Provisioning_Version1_Transaction
}

type Resource_Group struct {
	Entity

	AncestorGroupCount  *uint                      `json:"ancestorGroupCount:omitempty"`
	AncestorGroups      []Resource_Group           `json:"ancestorGroups:omitempty"`
	AttributeCount      *uint                      `json:"attributeCount:omitempty"`
	Attributes          []Resource_Group_Attribute `json:"attributes:omitempty"`
	CreateDate          *time.Time                 `json:"createDate:omitempty"`
	Description         *string                    `json:"description:omitempty"`
	HardwareMemberCount *uint                      `json:"hardwareMemberCount:omitempty"`
	HardwareMembers     []Resource_Group_Member    `json:"hardwareMembers:omitempty"`
	Id                  *int                       `json:"id:omitempty"`
	KeyName             *string                    `json:"keyName:omitempty"`
	MemberCount         *uint                      `json:"memberCount:omitempty"`
	Members             []Resource_Group_Member    `json:"members:omitempty"`
	Name                *string                    `json:"name:omitempty"`
	RootResourceGroup   *Resource_Group            `json:"rootResourceGroup:omitempty"`
	RootResourceGroupId *int                       `json:"rootResourceGroupId:omitempty"`
	SubnetMemberCount   *uint                      `json:"subnetMemberCount:omitempty"`
	SubnetMembers       []Resource_Group_Member    `json:"subnetMembers:omitempty"`
	Template            *Resource_Group_Template   `json:"template:omitempty"`
	TemplateId          *int                       `json:"templateId:omitempty"`
	VlanMemberCount     *uint                      `json:"vlanMemberCount:omitempty"`
	VlanMembers         []Resource_Group_Member    `json:"vlanMembers:omitempty"`
}

type Resource_Group_Attribute struct {
	Entity

	CreateDate *time.Time                     `json:"createDate:omitempty"`
	Group      *Resource_Group                `json:"group:omitempty"`
	Id         *int                           `json:"id:omitempty"`
	Type       *Resource_Group_Attribute_Type `json:"type:omitempty"`
	Value      *string                        `json:"value:omitempty"`
}

type Resource_Group_Attribute_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Resource_Group_Descendant_Reference struct {
	Entity

	Group       *Resource_Group        `json:"group:omitempty"`
	GroupMember *Resource_Group_Member `json:"groupMember:omitempty"`
}

type Resource_Group_Member struct {
	Entity

	AttributeCount        *uint                             `json:"attributeCount:omitempty"`
	Attributes            []Resource_Group_Member_Attribute `json:"attributes:omitempty"`
	CreateDate            *time.Time                        `json:"createDate:omitempty"`
	DescendantMemberCount *uint                             `json:"descendantMemberCount:omitempty"`
	DescendantMembers     []Resource_Group_Member           `json:"descendantMembers:omitempty"`
	Group                 *Resource_Group                   `json:"group:omitempty"`
	Id                    *int                              `json:"id:omitempty"`
	RoleCount             *uint                             `json:"roleCount:omitempty"`
	Roles                 []Resource_Group_Role             `json:"roles:omitempty"`
	Status                *string                           `json:"status:omitempty"`
	Type                  *Resource_Group_Member_Type       `json:"type:omitempty"`
}

type Resource_Group_Member_Attribute struct {
	Entity

	CreateDate *time.Time                            `json:"createDate:omitempty"`
	Id         *int                                  `json:"id:omitempty"`
	Member     *Resource_Group_Member                `json:"member:omitempty"`
	Type       *Resource_Group_Member_Attribute_Type `json:"type:omitempty"`
	Value      *string                               `json:"value:omitempty"`
}

type Resource_Group_Member_Attribute_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Resource_Group_Member_CloudStack_Version3_Cluster struct {
	Resource_Group_Member

	Resource *Resource_Group `json:"resource:omitempty"`
}

type Resource_Group_Member_CloudStack_Version3_Pod struct {
	Resource_Group_Member

	Resource *Resource_Group `json:"resource:omitempty"`
}

type Resource_Group_Member_CloudStack_Version3_Zone struct {
	Resource_Group_Member

	Resource *Resource_Group `json:"resource:omitempty"`
}

type Resource_Group_Member_Hardware struct {
	Resource_Group_Member

	Resource          *Hardware                        `json:"resource:omitempty"`
	ServerArbiterOnly *Resource_Group_Member_Attribute `json:"serverArbiterOnly:omitempty"`
	ServerHidden      *Resource_Group_Member_Attribute `json:"serverHidden:omitempty"`
	ServerPriority    *Resource_Group_Member_Attribute `json:"serverPriority:omitempty"`
	ServerSlaveDelay  *Resource_Group_Member_Attribute `json:"serverSlaveDelay:omitempty"`
	ServerTags        *Resource_Group_Member_Attribute `json:"serverTags:omitempty"`
	ServerVotes       *Resource_Group_Member_Attribute `json:"serverVotes:omitempty"`
}

type Resource_Group_Member_Network_Storage struct {
	Resource_Group_Member

	Resource *Network_Storage `json:"resource:omitempty"`
}

type Resource_Group_Member_Network_Subnet struct {
	Resource_Group_Member

	Resource *Network_Subnet `json:"resource:omitempty"`
}

type Resource_Group_Member_Network_Vlan struct {
	Resource_Group_Member

	Resource *Network_Vlan `json:"resource:omitempty"`
}

type Resource_Group_Member_Resource_Group struct {
	Resource_Group_Member

	Resource *Resource_Group `json:"resource:omitempty"`
}

type Resource_Group_Member_Role_Link struct {
	Entity

	GroupMemberId       *int `json:"groupMemberId:omitempty"`
	GroupTemplateRoleId *int `json:"groupTemplateRoleId:omitempty"`
}

type Resource_Group_Member_Software_Component_Password struct {
	Resource_Group_Member

	Resource *Software_Component_Password `json:"resource:omitempty"`
}

type Resource_Group_Member_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
}

type Resource_Group_Member_Virtual_Host_Pool struct {
	Resource_Group_Member
}

type Resource_Group_Role struct {
	Entity

	Description     *string                           `json:"description:omitempty"`
	Id              *int                              `json:"id:omitempty"`
	KeyName         *string                           `json:"keyName:omitempty"`
	MemberLinkCount *uint                             `json:"memberLinkCount:omitempty"`
	MemberLinks     []Resource_Group_Member_Role_Link `json:"memberLinks:omitempty"`
}

type Resource_Group_Template struct {
	Entity

	Children      []Resource_Group_Template        `json:"children:omitempty"`
	ChildrenCount *uint                            `json:"childrenCount:omitempty"`
	Description   *string                          `json:"description:omitempty"`
	Id            *int                             `json:"id:omitempty"`
	KeyName       *string                          `json:"keyName:omitempty"`
	MemberCount   *uint                            `json:"memberCount:omitempty"`
	Members       []Resource_Group_Template_Member `json:"members:omitempty"`
	Package       *Product_Package                 `json:"package:omitempty"`
}

type Resource_Group_Template_Member struct {
	Entity

	MaxQuantity *int                     `json:"maxQuantity:omitempty"`
	MinQuantity *int                     `json:"minQuantity:omitempty"`
	Role        *Resource_Group_Role     `json:"role:omitempty"`
	RoleId      *int                     `json:"roleId:omitempty"`
	Template    *Resource_Group_Template `json:"template:omitempty"`
	TemplateId  *int                     `json:"templateId:omitempty"`
}

type Resource_Metadata struct {
	Entity
}

type Sales_Presale_Event struct {
	Entity

	ActiveFlag  *bool           `json:"activeFlag:omitempty"`
	Description *string         `json:"description:omitempty"`
	EndDate     *time.Time      `json:"endDate:omitempty"`
	ExpiredFlag *bool           `json:"expiredFlag:omitempty"`
	Id          *int            `json:"id:omitempty"`
	Item        *Product_Item   `json:"item:omitempty"`
	ItemId      *int            `json:"itemId:omitempty"`
	Location    *Location       `json:"location:omitempty"`
	LocationId  *int            `json:"locationId:omitempty"`
	OrderCount  *uint           `json:"orderCount:omitempty"`
	Orders      []Billing_Order `json:"orders:omitempty"`
	StartDate   *time.Time      `json:"startDate:omitempty"`
}

type Scale_Asset struct {
	Entity

	CreateDate   *time.Time   `json:"createDate:omitempty"`
	DeleteFlag   *bool        `json:"deleteFlag:omitempty"`
	Id           *int         `json:"id:omitempty"`
	ScaleGroup   *Scale_Group `json:"scaleGroup:omitempty"`
	ScaleGroupId *int         `json:"scaleGroupId:omitempty"`
}

type Scale_Asset_Hardware struct {
	Scale_Asset

	Hardware   *Hardware `json:"hardware:omitempty"`
	HardwareId *int      `json:"hardwareId:omitempty"`
}

type Scale_Asset_Virtual_Guest struct {
	Scale_Asset

	VirtualGuest   *Virtual_Guest `json:"virtualGuest:omitempty"`
	VirtualGuestId *int           `json:"virtualGuestId:omitempty"`
}

type Scale_Group struct {
	Entity

	Account                    *Account                     `json:"account:omitempty"`
	AccountId                  *int                         `json:"accountId:omitempty"`
	BalancedTerminationFlag    *bool                        `json:"balancedTerminationFlag:omitempty"`
	Cooldown                   *int                         `json:"cooldown:omitempty"`
	CreateDate                 *time.Time                   `json:"createDate:omitempty"`
	DesiredMemberCount         *int                         `json:"desiredMemberCount:omitempty"`
	Id                         *int                         `json:"id:omitempty"`
	LastActionDate             *time.Time                   `json:"lastActionDate:omitempty"`
	LoadBalancerCount          *uint                        `json:"loadBalancerCount:omitempty"`
	LoadBalancers              []Scale_LoadBalancer         `json:"loadBalancers:omitempty"`
	LogCount                   *uint                        `json:"logCount:omitempty"`
	Logs                       []Scale_Group_Log            `json:"logs:omitempty"`
	MaximumMemberCount         *int                         `json:"maximumMemberCount:omitempty"`
	MinimumMemberCount         *int                         `json:"minimumMemberCount:omitempty"`
	ModifyDate                 *time.Time                   `json:"modifyDate:omitempty"`
	Name                       *string                      `json:"name:omitempty"`
	NetworkVlanCount           *uint                        `json:"networkVlanCount:omitempty"`
	NetworkVlans               []Scale_Network_Vlan         `json:"networkVlans:omitempty"`
	Policies                   []Scale_Policy               `json:"policies:omitempty"`
	PolicyCount                *uint                        `json:"policyCount:omitempty"`
	RegionalGroup              *Location_Group_Regional     `json:"regionalGroup:omitempty"`
	RegionalGroupId            *int                         `json:"regionalGroupId:omitempty"`
	Status                     *Scale_Group_Status          `json:"status:omitempty"`
	SuspendedFlag              *bool                        `json:"suspendedFlag:omitempty"`
	TerminationPolicy          *Scale_Termination_Policy    `json:"terminationPolicy:omitempty"`
	TerminationPolicyId        *int                         `json:"terminationPolicyId:omitempty"`
	VirtualGuestAssetCount     *uint                        `json:"virtualGuestAssetCount:omitempty"`
	VirtualGuestAssets         []Scale_Asset_Virtual_Guest  `json:"virtualGuestAssets:omitempty"`
	VirtualGuestMemberCount    *uint                        `json:"virtualGuestMemberCount:omitempty"`
	VirtualGuestMemberTemplate *Virtual_Guest               `json:"virtualGuestMemberTemplate:omitempty"`
	VirtualGuestMembers        []Scale_Member_Virtual_Guest `json:"virtualGuestMembers:omitempty"`
}

type Scale_Group_Log struct {
	Entity

	CreateDate   *time.Time   `json:"createDate:omitempty"`
	Description  *string      `json:"description:omitempty"`
	Id           *int         `json:"id:omitempty"`
	ScaleGroup   *Scale_Group `json:"scaleGroup:omitempty"`
	ScaleGroupId *int         `json:"scaleGroupId:omitempty"`
}

type Scale_Group_Status struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Scale_LoadBalancer struct {
	Entity

	AllocationPercent  *int                                                                 `json:"allocationPercent:omitempty"`
	CreateDate         *time.Time                                                           `json:"createDate:omitempty"`
	DeleteFlag         *bool                                                                `json:"deleteFlag:omitempty"`
	HealthCheck        *Network_Application_Delivery_Controller_LoadBalancer_Health_Check   `json:"healthCheck:omitempty"`
	HealthCheckId      *int                                                                 `json:"healthCheckId:omitempty"`
	Id                 *int                                                                 `json:"id:omitempty"`
	ModifyDate         *time.Time                                                           `json:"modifyDate:omitempty"`
	Port               *int                                                                 `json:"port:omitempty"`
	RoutingMethod      *Network_Application_Delivery_Controller_LoadBalancer_Routing_Method `json:"routingMethod:omitempty"`
	RoutingType        *Network_Application_Delivery_Controller_LoadBalancer_Routing_Type   `json:"routingType:omitempty"`
	ScaleGroup         *Scale_Group                                                         `json:"scaleGroup:omitempty"`
	ScaleGroupId       *int                                                                 `json:"scaleGroupId:omitempty"`
	VirtualIpAddressId *int                                                                 `json:"virtualIpAddressId:omitempty"`
	VirtualServer      *Network_Application_Delivery_Controller_LoadBalancer_VirtualServer  `json:"virtualServer:omitempty"`
	VirtualServerId    *int                                                                 `json:"virtualServerId:omitempty"`
	VirtualServerPort  *int                                                                 `json:"virtualServerPort:omitempty"`
}

type Scale_Member struct {
	Entity

	CreateDate   *time.Time   `json:"createDate:omitempty"`
	Id           *int         `json:"id:omitempty"`
	ScaleGroup   *Scale_Group `json:"scaleGroup:omitempty"`
	ScaleGroupId *int         `json:"scaleGroupId:omitempty"`
}

type Scale_Member_Virtual_Guest struct {
	Scale_Member

	VirtualGuest   *Virtual_Guest `json:"virtualGuest:omitempty"`
	VirtualGuestId *int           `json:"virtualGuestId:omitempty"`
}

type Scale_Network_Vlan struct {
	Entity

	CreateDate    *time.Time    `json:"createDate:omitempty"`
	DeleteFlag    *bool         `json:"deleteFlag:omitempty"`
	Id            *int          `json:"id:omitempty"`
	NetworkVlan   *Network_Vlan `json:"networkVlan:omitempty"`
	NetworkVlanId *int          `json:"networkVlanId:omitempty"`
	ScaleGroup    *Scale_Group  `json:"scaleGroup:omitempty"`
	ScaleGroupId  *int          `json:"scaleGroupId:omitempty"`
}

type Scale_Policy struct {
	Entity

	ActionCount             *uint                              `json:"actionCount:omitempty"`
	Actions                 []Scale_Policy_Action              `json:"actions:omitempty"`
	Cooldown                *int                               `json:"cooldown:omitempty"`
	CreateDate              *time.Time                         `json:"createDate:omitempty"`
	DeleteFlag              *bool                              `json:"deleteFlag:omitempty"`
	Id                      *int                               `json:"id:omitempty"`
	ModifyDate              *time.Time                         `json:"modifyDate:omitempty"`
	Name                    *string                            `json:"name:omitempty"`
	OneTimeTriggerCount     *uint                              `json:"oneTimeTriggerCount:omitempty"`
	OneTimeTriggers         []Scale_Policy_Trigger_OneTime     `json:"oneTimeTriggers:omitempty"`
	RepeatingTriggerCount   *uint                              `json:"repeatingTriggerCount:omitempty"`
	RepeatingTriggers       []Scale_Policy_Trigger_Repeating   `json:"repeatingTriggers:omitempty"`
	ResourceUseTriggerCount *uint                              `json:"resourceUseTriggerCount:omitempty"`
	ResourceUseTriggers     []Scale_Policy_Trigger_ResourceUse `json:"resourceUseTriggers:omitempty"`
	ScaleActionCount        *uint                              `json:"scaleActionCount:omitempty"`
	ScaleActions            []Scale_Policy_Action_Scale        `json:"scaleActions:omitempty"`
	ScaleGroup              *Scale_Group                       `json:"scaleGroup:omitempty"`
	ScaleGroupId            *int                               `json:"scaleGroupId:omitempty"`
	TriggerCount            *uint                              `json:"triggerCount:omitempty"`
	Triggers                []Scale_Policy_Trigger             `json:"triggers:omitempty"`
}

type Scale_Policy_Action struct {
	Entity

	CreateDate    *time.Time                `json:"createDate:omitempty"`
	DeleteFlag    *bool                     `json:"deleteFlag:omitempty"`
	Id            *int                      `json:"id:omitempty"`
	ModifyDate    *time.Time                `json:"modifyDate:omitempty"`
	ScalePolicy   *Scale_Policy             `json:"scalePolicy:omitempty"`
	ScalePolicyId *int                      `json:"scalePolicyId:omitempty"`
	Type          *Scale_Policy_Action_Type `json:"type:omitempty"`
	TypeId        *int                      `json:"typeId:omitempty"`
}

type Scale_Policy_Action_Scale struct {
	Scale_Policy_Action

	Amount    *int    `json:"amount:omitempty"`
	ScaleType *string `json:"scaleType:omitempty"`
}

type Scale_Policy_Action_Type struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Scale_Policy_Trigger struct {
	Entity

	CreateDate    *time.Time                 `json:"createDate:omitempty"`
	DeleteFlag    *bool                      `json:"deleteFlag:omitempty"`
	Id            *int                       `json:"id:omitempty"`
	ModifyDate    *time.Time                 `json:"modifyDate:omitempty"`
	ScalePolicy   *Scale_Policy              `json:"scalePolicy:omitempty"`
	ScalePolicyId *int                       `json:"scalePolicyId:omitempty"`
	Type          *Scale_Policy_Trigger_Type `json:"type:omitempty"`
	TypeId        *int                       `json:"typeId:omitempty"`
}

type Scale_Policy_Trigger_OneTime struct {
	Scale_Policy_Trigger

	Date *time.Time `json:"date:omitempty"`
}

type Scale_Policy_Trigger_Repeating struct {
	Scale_Policy_Trigger

	Schedule *string `json:"schedule:omitempty"`
}

type Scale_Policy_Trigger_ResourceUse struct {
	Scale_Policy_Trigger

	WatchCount *uint                                    `json:"watchCount:omitempty"`
	Watches    []Scale_Policy_Trigger_ResourceUse_Watch `json:"watches:omitempty"`
}

type Scale_Policy_Trigger_ResourceUse_Watch struct {
	Entity

	Algorithm            *string                           `json:"algorithm:omitempty"`
	CreateDate           *time.Time                        `json:"createDate:omitempty"`
	DeleteFlag           *bool                             `json:"deleteFlag:omitempty"`
	Id                   *int                              `json:"id:omitempty"`
	Metric               *string                           `json:"metric:omitempty"`
	ModifyDate           *time.Time                        `json:"modifyDate:omitempty"`
	Operator             *string                           `json:"operator:omitempty"`
	Period               *int                              `json:"period:omitempty"`
	ScalePolicyTrigger   *Scale_Policy_Trigger_ResourceUse `json:"scalePolicyTrigger:omitempty"`
	ScalePolicyTriggerId *int                              `json:"scalePolicyTriggerId:omitempty"`
	Value                *string                           `json:"value:omitempty"`
}

type Scale_Policy_Trigger_Type struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Scale_Termination_Policy struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Search struct {
	Entity
}

type Security_Certificate struct {
	Entity

	AssociatedServiceCount            *int                                                                    `json:"associatedServiceCount:omitempty"`
	Certificate                       *string                                                                 `json:"certificate:omitempty"`
	CertificateSigningRequest         *string                                                                 `json:"certificateSigningRequest:omitempty"`
	CommonName                        *string                                                                 `json:"commonName:omitempty"`
	CreateDate                        *time.Time                                                              `json:"createDate:omitempty"`
	Id                                *int                                                                    `json:"id:omitempty"`
	IntermediateCertificate           *string                                                                 `json:"intermediateCertificate:omitempty"`
	KeySize                           *int                                                                    `json:"keySize:omitempty"`
	LoadBalancerVirtualIpAddressCount *uint                                                                   `json:"loadBalancerVirtualIpAddressCount:omitempty"`
	LoadBalancerVirtualIpAddresses    []Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"loadBalancerVirtualIpAddresses:omitempty"`
	ModifyDate                        *time.Time                                                              `json:"modifyDate:omitempty"`
	Notes                             *string                                                                 `json:"notes:omitempty"`
	OrganizationName                  *string                                                                 `json:"organizationName:omitempty"`
	PrivateKey                        *string                                                                 `json:"privateKey:omitempty"`
	ValidityBegin                     *time.Time                                                              `json:"validityBegin:omitempty"`
	ValidityDays                      *int                                                                    `json:"validityDays:omitempty"`
	ValidityEnd                       *time.Time                                                              `json:"validityEnd:omitempty"`
}

type Security_Certificate_Entry struct {
	Entity

	CertificateId    *int       `json:"certificateId:omitempty"`
	CommonName       *string    `json:"commonName:omitempty"`
	KeySize          *int       `json:"keySize:omitempty"`
	OrganizationName *string    `json:"organizationName:omitempty"`
	ValidityBegin    *time.Time `json:"validityBegin:omitempty"`
	ValidityDays     *int       `json:"validityDays:omitempty"`
	ValidityEnd      *time.Time `json:"validityEnd:omitempty"`
}

type Security_Certificate_Request struct {
	Entity

	Account                      *Account                             `json:"account:omitempty"`
	AccountId                    *int                                 `json:"accountId:omitempty"`
	ApproverEmailAddress         *string                              `json:"approverEmailAddress:omitempty"`
	CertificateAuthorityName     *string                              `json:"certificateAuthorityName:omitempty"`
	CertificateSigningRequest    *string                              `json:"certificateSigningRequest:omitempty"`
	CommonName                   *string                              `json:"commonName:omitempty"`
	CreateDate                   *time.Time                           `json:"createDate:omitempty"`
	EffectiveDate                *time.Time                           `json:"effectiveDate:omitempty"`
	ExpirationDate               *time.Time                           `json:"expirationDate:omitempty"`
	Id                           *int                                 `json:"id:omitempty"`
	ModifyDate                   *time.Time                           `json:"modifyDate:omitempty"`
	Order                        *Billing_Order                       `json:"order:omitempty"`
	OrderItem                    *Billing_Order_Item                  `json:"orderItem:omitempty"`
	Status                       *Security_Certificate_Request_Status `json:"status:omitempty"`
	StatusId                     *int                                 `json:"statusId:omitempty"`
	TechnicalContactEmailAddress *string                              `json:"technicalContactEmailAddress:omitempty"`
}

type Security_Certificate_Request_ServerType struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	Name        *string `json:"name:omitempty"`
	Value       *int    `json:"value:omitempty"`
}

type Security_Certificate_Request_Status struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Security_Directory_Service_Host_Xref_Hardware struct {
	Entity

	Host *Hardware `json:"host:omitempty"`
}

type Security_SecureTransportCipher struct {
	Entity

	KeyName *string `json:"keyName:omitempty"`
}

type Security_SecureTransportProtocol struct {
	Entity

	KeyName                         *string                          `json:"keyName:omitempty"`
	SupportedSecureTransportCiphers []Security_SecureTransportCipher `json:"supportedSecureTransportCiphers:omitempty"`
}

type Security_Ssh_Key struct {
	Entity

	Account                       *Account                                    `json:"account:omitempty"`
	BlockDeviceTemplateGroupCount *uint                                       `json:"blockDeviceTemplateGroupCount:omitempty"`
	BlockDeviceTemplateGroups     []Virtual_Guest_Block_Device_Template_Group `json:"blockDeviceTemplateGroups:omitempty"`
	CreateDate                    *time.Time                                  `json:"createDate:omitempty"`
	Fingerprint                   *string                                     `json:"fingerprint:omitempty"`
	Id                            *int                                        `json:"id:omitempty"`
	Key                           *string                                     `json:"key:omitempty"`
	Label                         *string                                     `json:"label:omitempty"`
	ModifyDate                    *time.Time                                  `json:"modifyDate:omitempty"`
	Notes                         *string                                     `json:"notes:omitempty"`
	SoftwarePasswordCount         *uint                                       `json:"softwarePasswordCount:omitempty"`
	SoftwarePasswords             []Software_Component_Password               `json:"softwarePasswords:omitempty"`
}

type Service_Provider struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Software_AccountLicense struct {
	Entity

	Account             *Account              `json:"account:omitempty"`
	AccountId           *int                  `json:"accountId:omitempty"`
	BillingItem         *Billing_Item         `json:"billingItem:omitempty"`
	Capacity            *string               `json:"capacity:omitempty"`
	Key                 *string               `json:"key:omitempty"`
	SoftwareDescription *Software_Description `json:"softwareDescription:omitempty"`
	Units               *string               `json:"units:omitempty"`
}

type Software_Component struct {
	Entity

	AverageInstallationDuration *uint                                 `json:"averageInstallationDuration:omitempty"`
	BillingItem                 *Billing_Item                         `json:"billingItem:omitempty"`
	Hardware                    *Hardware                             `json:"hardware:omitempty"`
	HardwareId                  *int                                  `json:"hardwareId:omitempty"`
	Id                          *int                                  `json:"id:omitempty"`
	ManufacturerActivationCode  *string                               `json:"manufacturerActivationCode:omitempty"`
	ManufacturerLicenseInstance *string                               `json:"manufacturerLicenseInstance:omitempty"`
	PasswordCount               *uint                                 `json:"passwordCount:omitempty"`
	PasswordHistory             []Software_Component_Password_History `json:"passwordHistory:omitempty"`
	PasswordHistoryCount        *uint                                 `json:"passwordHistoryCount:omitempty"`
	Passwords                   []Software_Component_Password         `json:"passwords:omitempty"`
	SoftwareDescription         *Software_Description                 `json:"softwareDescription:omitempty"`
	SoftwareLicense             *Software_License                     `json:"softwareLicense:omitempty"`
	VirtualGuest                *Virtual_Guest                        `json:"virtualGuest:omitempty"`
}

type Software_Component_Analytics struct {
	Software_Component
}

type Software_Component_Analytics_Urchin struct {
	Software_Component_Analytics
}

type Software_Component_AntivirusSpyware struct {
	Software_Component
}

type Software_Component_AntivirusSpyware_Mcafee struct {
	Software_Component_AntivirusSpyware
}

type Software_Component_AntivirusSpyware_Mcafee_Epo_Version36 struct {
	Software_Component_AntivirusSpyware_Mcafee

	AgentDetails                     *McAfee_Epolicy_Orchestrator_Version36_Agent_Details                     `json:"agentDetails:omitempty"`
	CurrentAntivirusPolicy           *int                                                                     `json:"currentAntivirusPolicy:omitempty"`
	DataFileVersion                  *McAfee_Epolicy_Orchestrator_Version36_Product_Properties                `json:"dataFileVersion:omitempty"`
	EpoVersion                       *string                                                                  `json:"epoVersion:omitempty"`
	LatestAccessProtectionEventCount *uint                                                                    `json:"latestAccessProtectionEventCount:omitempty"`
	LatestAccessProtectionEvents     []McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_AccessProtection `json:"latestAccessProtectionEvents:omitempty"`
	LatestAntivirusEventCount        *uint                                                                    `json:"latestAntivirusEventCount:omitempty"`
	LatestAntivirusEvents            []McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event                  `json:"latestAntivirusEvents:omitempty"`
	LatestSpywareEventCount          *uint                                                                    `json:"latestSpywareEventCount:omitempty"`
	LatestSpywareEvents              []McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event                  `json:"latestSpywareEvents:omitempty"`
	TransactionStatus                *string                                                                  `json:"transactionStatus:omitempty"`
}

type Software_Component_AntivirusSpyware_Mcafee_Epo_Version45 struct {
	Software_Component_AntivirusSpyware_Mcafee

	AgentDetails                     *McAfee_Epolicy_Orchestrator_Version45_Agent_Details      `json:"agentDetails:omitempty"`
	CurrentAntivirusPolicy           *int                                                      `json:"currentAntivirusPolicy:omitempty"`
	DataFileVersion                  *McAfee_Epolicy_Orchestrator_Version45_Product_Properties `json:"dataFileVersion:omitempty"`
	EpoVersion                       *string                                                   `json:"epoVersion:omitempty"`
	LatestAccessProtectionEventCount *uint                                                     `json:"latestAccessProtectionEventCount:omitempty"`
	LatestAccessProtectionEvents     []McAfee_Epolicy_Orchestrator_Version45_Event             `json:"latestAccessProtectionEvents:omitempty"`
	LatestAntivirusEventCount        *uint                                                     `json:"latestAntivirusEventCount:omitempty"`
	LatestAntivirusEvents            []McAfee_Epolicy_Orchestrator_Version45_Event             `json:"latestAntivirusEvents:omitempty"`
	LatestSpywareEventCount          *uint                                                     `json:"latestSpywareEventCount:omitempty"`
	LatestSpywareEvents              []McAfee_Epolicy_Orchestrator_Version45_Event             `json:"latestSpywareEvents:omitempty"`
	TransactionStatus                *string                                                   `json:"transactionStatus:omitempty"`
}

type Software_Component_ControlPanel struct {
	Software_Component
}

type Software_Component_ControlPanel_Cpanel struct {
	Software_Component
}

type Software_Component_ControlPanel_Idera struct {
	Software_Component
}

type Software_Component_ControlPanel_Idera_ServerBackup struct {
	Software_Component_ControlPanel_Idera
}

type Software_Component_ControlPanel_Microsoft struct {
	Software_Component
}

type Software_Component_ControlPanel_Microsoft_WebPlatform struct {
	Software_Component_ControlPanel_Microsoft
}

type Software_Component_ControlPanel_Parallels struct {
	Software_Component
}

type Software_Component_ControlPanel_Parallels_Plesk struct {
	Software_Component_ControlPanel_Parallels
}

type Software_Component_ControlPanel_R1soft struct {
	Software_Component
}

type Software_Component_ControlPanel_R1soft_Cdp struct {
	Software_Component_ControlPanel_R1soft
}

type Software_Component_ControlPanel_R1soft_ServerBackup struct {
	Software_Component_ControlPanel_R1soft
}

type Software_Component_ControlPanel_Swsoft struct {
	Software_Component
}

type Software_Component_ControlPanel_WebhostAutomation struct {
	Software_Component
}

type Software_Component_HostIps struct {
	Software_Component
}

type Software_Component_HostIps_Mcafee struct {
	Software_Component_HostIps
}

type Software_Component_HostIps_Mcafee_Epo_Version36_Hips struct {
	Software_Component_HostIps_Mcafee

	AgentDetails                      *McAfee_Epolicy_Orchestrator_Version36_Agent_Details  `json:"agentDetails:omitempty"`
	ApplicationModePolicyNameCount    *uint                                                 `json:"applicationModePolicyNameCount:omitempty"`
	ApplicationModePolicyNames        []McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"applicationModePolicyNames:omitempty"`
	ApplicationRuleSetPolicyNameCount *uint                                                 `json:"applicationRuleSetPolicyNameCount:omitempty"`
	ApplicationRuleSetPolicyNames     []McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"applicationRuleSetPolicyNames:omitempty"`
	EnforcementPolicyNameCount        *uint                                                 `json:"enforcementPolicyNameCount:omitempty"`
	EnforcementPolicyNames            []McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"enforcementPolicyNames:omitempty"`
	EpoVersion                        *string                                               `json:"epoVersion:omitempty"`
	FirewallModePolicyNameCount       *uint                                                 `json:"firewallModePolicyNameCount:omitempty"`
	FirewallModePolicyNames           []McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"firewallModePolicyNames:omitempty"`
	FirewallRuleSetPolicyNameCount    *uint                                                 `json:"firewallRuleSetPolicyNameCount:omitempty"`
	FirewallRuleSetPolicyNames        []McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"firewallRuleSetPolicyNames:omitempty"`
	IpsModePolicyNameCount            *uint                                                 `json:"ipsModePolicyNameCount:omitempty"`
	IpsModePolicyNames                []McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"ipsModePolicyNames:omitempty"`
	IpsProtectionPolicyNameCount      *uint                                                 `json:"ipsProtectionPolicyNameCount:omitempty"`
	IpsProtectionPolicyNames          []McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"ipsProtectionPolicyNames:omitempty"`
	TransactionStatus                 *string                                               `json:"transactionStatus:omitempty"`
}

type Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version6 struct {
	Software_Component_HostIps_Mcafee_Epo_Version36_Hips

	BlockedApplicationEventCount *uint                                                                         `json:"blockedApplicationEventCount:omitempty"`
	BlockedApplicationEvents     []McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_BlockedApplicationEvent `json:"blockedApplicationEvents:omitempty"`
	IpsEventCount                *uint                                                                         `json:"ipsEventCount:omitempty"`
	IpsEvents                    []McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_IPSEvent                `json:"ipsEvents:omitempty"`
}

type Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version7 struct {
	Software_Component_HostIps_Mcafee_Epo_Version36_Hips

	BlockedApplicationEventCount *uint                                                                         `json:"blockedApplicationEventCount:omitempty"`
	BlockedApplicationEvents     []McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_BlockedApplicationEvent `json:"blockedApplicationEvents:omitempty"`
	IpsEventCount                *uint                                                                         `json:"ipsEventCount:omitempty"`
	IpsEvents                    []McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent                `json:"ipsEvents:omitempty"`
}

type Software_Component_HostIps_Mcafee_Epo_Version45_Hips struct {
	Software_Component_HostIps_Mcafee

	AgentDetails                      *McAfee_Epolicy_Orchestrator_Version45_Agent_Details  `json:"agentDetails:omitempty"`
	ApplicationModePolicyNameCount    *uint                                                 `json:"applicationModePolicyNameCount:omitempty"`
	ApplicationModePolicyNames        []McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"applicationModePolicyNames:omitempty"`
	ApplicationRuleSetPolicyNameCount *uint                                                 `json:"applicationRuleSetPolicyNameCount:omitempty"`
	ApplicationRuleSetPolicyNames     []McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"applicationRuleSetPolicyNames:omitempty"`
	BlockedApplicationEventCount      *uint                                                 `json:"blockedApplicationEventCount:omitempty"`
	BlockedApplicationEvents          []McAfee_Epolicy_Orchestrator_Version45_Event         `json:"blockedApplicationEvents:omitempty"`
	EnforcementPolicyNameCount        *uint                                                 `json:"enforcementPolicyNameCount:omitempty"`
	EnforcementPolicyNames            []McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"enforcementPolicyNames:omitempty"`
	EpoVersion                        *string                                               `json:"epoVersion:omitempty"`
	FirewallModePolicyNameCount       *uint                                                 `json:"firewallModePolicyNameCount:omitempty"`
	FirewallModePolicyNames           []McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"firewallModePolicyNames:omitempty"`
	FirewallRuleSetPolicyNameCount    *uint                                                 `json:"firewallRuleSetPolicyNameCount:omitempty"`
	FirewallRuleSetPolicyNames        []McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"firewallRuleSetPolicyNames:omitempty"`
	IpsEventCount                     *uint                                                 `json:"ipsEventCount:omitempty"`
	IpsEvents                         []McAfee_Epolicy_Orchestrator_Version45_Event         `json:"ipsEvents:omitempty"`
	IpsModePolicyNameCount            *uint                                                 `json:"ipsModePolicyNameCount:omitempty"`
	IpsModePolicyNames                []McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"ipsModePolicyNames:omitempty"`
	IpsProtectionPolicyNameCount      *uint                                                 `json:"ipsProtectionPolicyNameCount:omitempty"`
	IpsProtectionPolicyNames          []McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"ipsProtectionPolicyNames:omitempty"`
	TransactionStatus                 *string                                               `json:"transactionStatus:omitempty"`
}

type Software_Component_HostIps_Mcafee_Epo_Version45_Hips_Version7 struct {
	Software_Component_HostIps_Mcafee_Epo_Version45_Hips
}

type Software_Component_HostIps_Mcafee_Epo_Version45_Hips_Version8 struct {
	Software_Component_HostIps_Mcafee_Epo_Version45_Hips
}

type Software_Component_OperatingSystem struct {
	Software_Component

	LicenseExpirationDate  *time.Time                               `json:"licenseExpirationDate:omitempty"`
	PartitionTemplateCount *uint                                    `json:"partitionTemplateCount:omitempty"`
	PartitionTemplates     []Hardware_Component_Partition_Template  `json:"partitionTemplates:omitempty"`
	ReloadTransactionGroup *Provisioning_Version1_Transaction_Group `json:"reloadTransactionGroup:omitempty"`
}

type Software_Component_Package struct {
	Software_Component
}

type Software_Component_Package_Management struct {
	Software_Component_Package
}

type Software_Component_Package_Management_Ksplice struct {
	Software_Component_Package_Management
}

type Software_Component_Password struct {
	Entity

	CreateDate  *time.Time          `json:"createDate:omitempty"`
	Id          *int                `json:"id:omitempty"`
	ModifyDate  *time.Time          `json:"modifyDate:omitempty"`
	Notes       *string             `json:"notes:omitempty"`
	Password    *string             `json:"password:omitempty"`
	Port        *int                `json:"port:omitempty"`
	Software    *Software_Component `json:"software:omitempty"`
	SoftwareId  *int                `json:"softwareId:omitempty"`
	SshKeyCount *uint               `json:"sshKeyCount:omitempty"`
	SshKeys     []Security_Ssh_Key  `json:"sshKeys:omitempty"`
	Username    *string             `json:"username:omitempty"`
}

type Software_Component_Password_History struct {
	Entity

	CreateDate          *time.Time          `json:"createDate:omitempty"`
	Notes               *string             `json:"notes:omitempty"`
	Password            *string             `json:"password:omitempty"`
	SoftwareComponent   *Software_Component `json:"softwareComponent:omitempty"`
	SoftwareComponentId *int                `json:"softwareComponentId:omitempty"`
	Username            *string             `json:"username:omitempty"`
}

type Software_Component_Security struct {
	Software_Component
}

type Software_Component_Security_SafeNet struct {
	Software_Component_Security
}

type Software_Description struct {
	Entity

	AttributeCount                     *uint                                    `json:"attributeCount:omitempty"`
	Attributes                         []Software_Description_Attribute         `json:"attributes:omitempty"`
	AverageInstallationDuration        *int                                     `json:"averageInstallationDuration:omitempty"`
	CompatibleSoftwareDescriptionCount *uint                                    `json:"compatibleSoftwareDescriptionCount:omitempty"`
	CompatibleSoftwareDescriptions     []Software_Description                   `json:"compatibleSoftwareDescriptions:omitempty"`
	ControlPanel                       *int                                     `json:"controlPanel:omitempty"`
	FeatureCount                       *uint                                    `json:"featureCount:omitempty"`
	Features                           []Software_Description_Feature           `json:"features:omitempty"`
	Id                                 *int                                     `json:"id:omitempty"`
	LatestVersion                      []Software_Description                   `json:"latestVersion:omitempty"`
	LatestVersionCount                 *uint                                    `json:"latestVersionCount:omitempty"`
	LicenseTermUnit                    *string                                  `json:"licenseTermUnit:omitempty"`
	LicenseTermValue                   *int                                     `json:"licenseTermValue:omitempty"`
	LongDescription                    *string                                  `json:"longDescription:omitempty"`
	Manufacturer                       *string                                  `json:"manufacturer:omitempty"`
	Name                               *string                                  `json:"name:omitempty"`
	OperatingSystem                    *int                                     `json:"operatingSystem:omitempty"`
	ProductItemCount                   *uint                                    `json:"productItemCount:omitempty"`
	ProductItems                       []Product_Item                           `json:"productItems:omitempty"`
	ProvisionTransactionGroup          *Provisioning_Version1_Transaction_Group `json:"provisionTransactionGroup:omitempty"`
	ReferenceCode                      *string                                  `json:"referenceCode:omitempty"`
	ReloadTransactionGroup             *Provisioning_Version1_Transaction_Group `json:"reloadTransactionGroup:omitempty"`
	RequiredUser                       *string                                  `json:"requiredUser:omitempty"`
	SoftwareLicenseCount               *uint                                    `json:"softwareLicenseCount:omitempty"`
	SoftwareLicenses                   []Software_License                       `json:"softwareLicenses:omitempty"`
	UpgradeSoftwareDescription         *Software_Description                    `json:"upgradeSoftwareDescription:omitempty"`
	UpgradeSoftwareDescriptionId       *int                                     `json:"upgradeSoftwareDescriptionId:omitempty"`
	UpgradeSwDesc                      *Software_Description                    `json:"upgradeSwDesc:omitempty"`
	UpgradeSwDescId                    *int                                     `json:"upgradeSwDescId:omitempty"`
	ValidFilesystemTypeCount           *uint                                    `json:"validFilesystemTypeCount:omitempty"`
	ValidFilesystemTypes               []Configuration_Storage_Filesystem_Type  `json:"validFilesystemTypes:omitempty"`
	Version                            *string                                  `json:"version:omitempty"`
	VirtualLicense                     *int                                     `json:"virtualLicense:omitempty"`
	VirtualizationPlatform             *int                                     `json:"virtualizationPlatform:omitempty"`
}

type Software_Description_Attribute struct {
	Entity

	SoftwareDescription *Software_Description                `json:"softwareDescription:omitempty"`
	Type                *Software_Description_Attribute_Type `json:"type:omitempty"`
	Value               *string                              `json:"value:omitempty"`
}

type Software_Description_Attribute_Type struct {
	Entity

	Keyname *string `json:"keyname:omitempty"`
}

type Software_Description_Feature struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
	Vendor  *string `json:"vendor:omitempty"`
}

type Software_Description_RequiredUser struct {
	Entity

	DefaultPassword *string `json:"defaultPassword:omitempty"`
	Username        *string `json:"username:omitempty"`
}

type Software_License struct {
	Entity

	Account               *Account              `json:"account:omitempty"`
	Id                    *int                  `json:"id:omitempty"`
	Owner                 *Account              `json:"owner:omitempty"`
	SoftwareDescription   *Software_Description `json:"softwareDescription:omitempty"`
	SoftwareDescriptionId *int                  `json:"softwareDescriptionId:omitempty"`
}

type Software_VirtualLicense struct {
	Entity

	Account               *Account                  `json:"account:omitempty"`
	AccountId             *int                      `json:"accountId:omitempty"`
	BillingItem           *Billing_Item             `json:"billingItem:omitempty"`
	HostHardware          *Hardware_Server          `json:"hostHardware:omitempty"`
	HostHardwareId        *int                      `json:"hostHardwareId:omitempty"`
	Id                    *int                      `json:"id:omitempty"`
	IpAddress             *string                   `json:"ipAddress:omitempty"`
	IpAddressRecord       *Network_Subnet_IpAddress `json:"ipAddressRecord:omitempty"`
	Key                   *string                   `json:"key:omitempty"`
	Notes                 *string                   `json:"notes:omitempty"`
	SoftwareDescription   *Software_Description     `json:"softwareDescription:omitempty"`
	SoftwareDescriptionId *int                      `json:"softwareDescriptionId:omitempty"`
	Subnet                *Network_Subnet           `json:"subnet:omitempty"`
	SubnetId              *int                      `json:"subnetId:omitempty"`
}

type Survey struct {
	Entity

	Active        *int              `json:"active:omitempty"`
	CreateDate    *time.Time        `json:"createDate:omitempty"`
	Id            *int              `json:"id:omitempty"`
	Name          *string           `json:"name:omitempty"`
	QuestionCount *uint             `json:"questionCount:omitempty"`
	Questions     []Survey_Question `json:"questions:omitempty"`
	Status        *Survey_Status    `json:"status:omitempty"`
	StatusId      *int              `json:"statusId:omitempty"`
	Type          *Survey_Type      `json:"type:omitempty"`
	TypeId        *int              `json:"typeId:omitempty"`
}

type Survey_Answer struct {
	Entity

	Answer           *string          `json:"answer:omitempty"`
	AnswerOrder      *int             `json:"answerOrder:omitempty"`
	Id               *int             `json:"id:omitempty"`
	SurveyQuestion   *Survey_Question `json:"surveyQuestion:omitempty"`
	SurveyQuestionId *int             `json:"surveyQuestionId:omitempty"`
}

type Survey_Question struct {
	Entity

	AnswerCount   *uint           `json:"answerCount:omitempty"`
	Answers       []Survey_Answer `json:"answers:omitempty"`
	Id            *int            `json:"id:omitempty"`
	IsRequired    *int            `json:"isRequired:omitempty"`
	MultiAnswer   *int            `json:"multiAnswer:omitempty"`
	Question      *string         `json:"question:omitempty"`
	QuestionOrder *int            `json:"questionOrder:omitempty"`
	Survey        *Survey         `json:"survey:omitempty"`
	SurveyId      *int            `json:"surveyId:omitempty"`
}

type Survey_Response struct {
	Entity

	OtherAnswer    *string        `json:"otherAnswer:omitempty"`
	SurveyAnswer   *Survey_Answer `json:"surveyAnswer:omitempty"`
	SurveyAnswerId *int           `json:"surveyAnswerId:omitempty"`
}

type Survey_Status struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Survey_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Tag struct {
	Entity

	Account        *Account        `json:"account:omitempty"`
	AccountId      *int            `json:"accountId:omitempty"`
	Id             *int            `json:"id:omitempty"`
	Internal       *int            `json:"internal:omitempty"`
	Name           *string         `json:"name:omitempty"`
	ReferenceCount *uint           `json:"referenceCount:omitempty"`
	References     []Tag_Reference `json:"references:omitempty"`
}

type Tag_Reference struct {
	Entity

	Customer        *User_Customer `json:"customer:omitempty"`
	EmpRecordId     *int           `json:"empRecordId:omitempty"`
	Employee        *User_Employee `json:"employee:omitempty"`
	Id              *int           `json:"id:omitempty"`
	ResourceTableId *int           `json:"resourceTableId:omitempty"`
	Tag             *Tag           `json:"tag:omitempty"`
	TagId           *int           `json:"tagId:omitempty"`
	TagType         *Tag_Type      `json:"tagType:omitempty"`
	TagTypeId       *int           `json:"tagTypeId:omitempty"`
	UsrRecordId     *int           `json:"usrRecordId:omitempty"`
}

type Tag_Reference_Hardware struct {
	Tag_Reference

	Resource *Hardware `json:"resource:omitempty"`
}

type Tag_Reference_Network_Application_Delivery_Controller struct {
	Tag_Reference

	Resource *Network_Application_Delivery_Controller `json:"resource:omitempty"`
}

type Tag_Reference_Network_Vlan struct {
	Tag_Reference

	Resource *Network_Vlan `json:"resource:omitempty"`
}

type Tag_Reference_Network_Vlan_Firewall struct {
	Tag_Reference

	Resource *Network_Vlan_Firewall `json:"resource:omitempty"`
}

type Tag_Reference_Resource_Group struct {
	Tag_Reference

	Resource *Resource_Group `json:"resource:omitempty"`
}

type Tag_Reference_Virtual_Guest struct {
	Tag_Reference

	Resource *Virtual_Guest `json:"resource:omitempty"`
}

type Tag_Reference_Virtual_Guest_Block_Device_Template_Group struct {
	Tag_Reference

	Resource *Virtual_Guest_Block_Device_Template_Group `json:"resource:omitempty"`
}

type Tag_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
}

type Ticket struct {
	Entity

	Account                              *Account                            `json:"account:omitempty"`
	AccountId                            *int                                `json:"accountId:omitempty"`
	AssignedAgentCount                   *uint                               `json:"assignedAgentCount:omitempty"`
	AssignedAgents                       []User_Customer                     `json:"assignedAgents:omitempty"`
	AssignedUser                         *User_Customer                      `json:"assignedUser:omitempty"`
	AssignedUserId                       *int                                `json:"assignedUserId:omitempty"`
	AttachedAdditionalEmailCount         *uint                               `json:"attachedAdditionalEmailCount:omitempty"`
	AttachedAdditionalEmails             []User_Customer_AdditionalEmail     `json:"attachedAdditionalEmails:omitempty"`
	AttachedFileCount                    *uint                               `json:"attachedFileCount:omitempty"`
	AttachedFiles                        []Ticket_Attachment_File            `json:"attachedFiles:omitempty"`
	AttachedHardware                     []Hardware                          `json:"attachedHardware:omitempty"`
	AttachedHardwareCount                *uint                               `json:"attachedHardwareCount:omitempty"`
	AttachedResourceCount                *uint                               `json:"attachedResourceCount:omitempty"`
	AttachedResources                    []Ticket_Attachment                 `json:"attachedResources:omitempty"`
	AttachedVirtualGuestCount            *uint                               `json:"attachedVirtualGuestCount:omitempty"`
	AttachedVirtualGuests                []Virtual_Guest                     `json:"attachedVirtualGuests:omitempty"`
	AwaitingUserResponseFlag             *bool                               `json:"awaitingUserResponseFlag:omitempty"`
	BillableFlag                         *bool                               `json:"billableFlag:omitempty"`
	CancellationRequest                  *Billing_Item_Cancellation_Request  `json:"cancellationRequest:omitempty"`
	ChangeOwnerFlag                      *bool                               `json:"changeOwnerFlag:omitempty"`
	CreateDate                           *time.Time                          `json:"createDate:omitempty"`
	EmployeeAttachmentCount              *uint                               `json:"employeeAttachmentCount:omitempty"`
	EmployeeAttachments                  []User_Employee                     `json:"employeeAttachments:omitempty"`
	FinalComments                        *string                             `json:"finalComments:omitempty"`
	FirstAttachedResource                *Ticket_Attachment                  `json:"firstAttachedResource:omitempty"`
	FirstUpdate                          *Ticket_Update                      `json:"firstUpdate:omitempty"`
	Group                                *Ticket_Group                       `json:"group:omitempty"`
	GroupId                              *int                                `json:"groupId:omitempty"`
	Id                                   *int                                `json:"id:omitempty"`
	InvoiceItemCount                     *uint                               `json:"invoiceItemCount:omitempty"`
	InvoiceItems                         []Billing_Invoice_Item              `json:"invoiceItems:omitempty"`
	LastActivity                         *Ticket_Activity                    `json:"lastActivity:omitempty"`
	LastEditDate                         *time.Time                          `json:"lastEditDate:omitempty"`
	LastEditType                         *string                             `json:"lastEditType:omitempty"`
	LastEditor                           *User_Interface                     `json:"lastEditor:omitempty"`
	LastResponseDate                     *time.Time                          `json:"lastResponseDate:omitempty"`
	LastUpdate                           *Ticket_Update                      `json:"lastUpdate:omitempty"`
	LastViewedDate                       *time.Time                          `json:"lastViewedDate:omitempty"`
	Location                             *Location                           `json:"location:omitempty"`
	LocationId                           *int                                `json:"locationId:omitempty"`
	ModifyDate                           *time.Time                          `json:"modifyDate:omitempty"`
	NewUpdatesFlag                       *bool                               `json:"newUpdatesFlag:omitempty"`
	NotifyUserOnUpdateFlag               *bool                               `json:"notifyUserOnUpdateFlag:omitempty"`
	OriginatingIpAddress                 *string                             `json:"originatingIpAddress:omitempty"`
	Priority                             *int                                `json:"priority:omitempty"`
	ResponsibleBrandId                   *int                                `json:"responsibleBrandId:omitempty"`
	ScheduledActionCount                 *uint                               `json:"scheduledActionCount:omitempty"`
	ScheduledActions                     []Provisioning_Version1_Transaction `json:"scheduledActions:omitempty"`
	ServerAdministrationBillingAmount    *int                                `json:"serverAdministrationBillingAmount:omitempty"`
	ServerAdministrationBillingInvoice   *Billing_Invoice                    `json:"serverAdministrationBillingInvoice:omitempty"`
	ServerAdministrationBillingInvoiceId *int                                `json:"serverAdministrationBillingInvoiceId:omitempty"`
	ServerAdministrationFlag             *int                                `json:"serverAdministrationFlag:omitempty"`
	ServerAdministrationRefundInvoice    *Billing_Invoice                    `json:"serverAdministrationRefundInvoice:omitempty"`
	ServerAdministrationRefundInvoiceId  *int                                `json:"serverAdministrationRefundInvoiceId:omitempty"`
	ServiceProvider                      *Service_Provider                   `json:"serviceProvider:omitempty"`
	ServiceProviderId                    *int                                `json:"serviceProviderId:omitempty"`
	ServiceProviderResourceId            *int                                `json:"serviceProviderResourceId:omitempty"`
	State                                []Ticket_State                      `json:"state:omitempty"`
	StateCount                           *uint                               `json:"stateCount:omitempty"`
	Status                               *Ticket_Status                      `json:"status:omitempty"`
	StatusId                             *int                                `json:"statusId:omitempty"`
	Subject                              *Ticket_Subject                     `json:"subject:omitempty"`
	SubjectId                            *int                                `json:"subjectId:omitempty"`
	TagReferenceCount                    *uint                               `json:"tagReferenceCount:omitempty"`
	TagReferences                        []Tag_Reference                     `json:"tagReferences:omitempty"`
	Title                                *string                             `json:"title:omitempty"`
	TotalUpdateCount                     *int                                `json:"totalUpdateCount:omitempty"`
	UpdateCount                          *uint                               `json:"updateCount:omitempty"`
	Updates                              []Ticket_Update                     `json:"updates:omitempty"`
	UserEditableFlag                     *bool                               `json:"userEditableFlag:omitempty"`
}

type Ticket_Activity struct {
	Entity

	CreateDate      *time.Time      `json:"createDate:omitempty"`
	CreateTimestamp *time.Time      `json:"createTimestamp:omitempty"`
	Editor          *User_Interface `json:"editor:omitempty"`
	Id              *int            `json:"id:omitempty"`
	Ticket          *Ticket         `json:"ticket:omitempty"`
	TicketUpdate    *Ticket_Update  `json:"ticketUpdate:omitempty"`
	Value           *string         `json:"value:omitempty"`
}

type Ticket_Attachment struct {
	Entity

	AssignedAgent   *User_Customer                     `json:"assignedAgent:omitempty"`
	AttachmentId    *int                               `json:"attachmentId:omitempty"`
	CreateDate      *time.Time                         `json:"createDate:omitempty"`
	Id              *int                               `json:"id:omitempty"`
	ScheduledAction *Provisioning_Version1_Transaction `json:"scheduledAction:omitempty"`
	Ticket          *Ticket                            `json:"ticket:omitempty"`
	TicketId        *int                               `json:"ticketId:omitempty"`
}

type Ticket_Attachment_Assigned_Agent struct {
	Ticket_Attachment

	AssignedAgentId *int           `json:"assignedAgentId:omitempty"`
	Resource        *User_Customer `json:"resource:omitempty"`
}

type Ticket_Attachment_CardChangeRequest struct {
	Ticket_Attachment

	Resource *Billing_Payment_Card_ChangeRequest `json:"resource:omitempty"`
}

type Ticket_Attachment_File struct {
	Entity

	CreateDate   *time.Time     `json:"createDate:omitempty"`
	FileName     *string        `json:"fileName:omitempty"`
	FileSize     *string        `json:"fileSize:omitempty"`
	Id           *int           `json:"id:omitempty"`
	ModifyDate   *time.Time     `json:"modifyDate:omitempty"`
	Ticket       *Ticket        `json:"ticket:omitempty"`
	TicketId     *int           `json:"ticketId:omitempty"`
	Update       *Ticket_Update `json:"update:omitempty"`
	UpdateId     *int           `json:"updateId:omitempty"`
	UploaderId   *string        `json:"uploaderId:omitempty"`
	UploaderType *string        `json:"uploaderType:omitempty"`
}

type Ticket_Attachment_Hardware struct {
	Ticket_Attachment

	Hardware   *Hardware `json:"hardware:omitempty"`
	HardwareId *int      `json:"hardwareId:omitempty"`
	Resource   *Hardware `json:"resource:omitempty"`
}

type Ticket_Attachment_ManualPayment struct {
	Ticket_Attachment

	Resource *Billing_Payment_Card_ManualPayment `json:"resource:omitempty"`
}

type Ticket_Attachment_Scheduled_Action struct {
	Ticket_Attachment

	Resource      *Provisioning_Version1_Transaction `json:"resource:omitempty"`
	RunDate       *time.Time                         `json:"runDate:omitempty"`
	Transaction   *Provisioning_Version1_Transaction `json:"transaction:omitempty"`
	TransactionId *int                               `json:"transactionId:omitempty"`
}

type Ticket_Attachment_Virtual_Guest struct {
	Ticket_Attachment

	Resource       *Virtual_Guest `json:"resource:omitempty"`
	VirtualGuest   *Virtual_Guest `json:"virtualGuest:omitempty"`
	VirtualGuestId *int           `json:"virtualGuestId:omitempty"`
}

type Ticket_Chat struct {
	Entity

	Agent        *User_Employee      `json:"agent:omitempty"`
	Customer     *User_Customer      `json:"customer:omitempty"`
	CustomerId   *int                `json:"customerId:omitempty"`
	EndDate      *time.Time          `json:"endDate:omitempty"`
	StartDate    *time.Time          `json:"startDate:omitempty"`
	TicketUpdate *Ticket_Update_Chat `json:"ticketUpdate:omitempty"`
	Transcript   *string             `json:"transcript:omitempty"`
}

type Ticket_Chat_Liveperson struct {
	Ticket_Chat
}

type Ticket_Chat_TranscriptLine struct {
	Entity

	Speaker *User_Interface `json:"speaker:omitempty"`
}

type Ticket_Chat_TranscriptLine_Customer struct {
	Ticket_Chat_TranscriptLine
}

type Ticket_Chat_TranscriptLine_Employee struct {
	Ticket_Chat_TranscriptLine
}

type Ticket_Group struct {
	Entity

	AssignedBrandCount    *uint                  `json:"assignedBrandCount:omitempty"`
	AssignedBrands        []Brand                `json:"assignedBrands:omitempty"`
	Category              *Ticket_Group_Category `json:"category:omitempty"`
	Id                    *int                   `json:"id:omitempty"`
	Name                  *string                `json:"name:omitempty"`
	TicketGroupCategoryId *int                   `json:"ticketGroupCategoryId:omitempty"`
}

type Ticket_Group_Category struct {
	Entity

	Id   *int    `json:"id:omitempty"`
	Name *string `json:"name:omitempty"`
}

type Ticket_Priority struct {
	Entity
}

type Ticket_State struct {
	Entity

	Id          *int               `json:"id:omitempty"`
	StateType   *Ticket_State_Type `json:"stateType:omitempty"`
	StateTypeId *int               `json:"stateTypeId:omitempty"`
	Ticket      *Ticket            `json:"ticket:omitempty"`
	TicketId    *int               `json:"ticketId:omitempty"`
}

type Ticket_State_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Ticket_Status struct {
	Entity

	Id   *int    `json:"id:omitempty"`
	Name *string `json:"name:omitempty"`
}

type Ticket_Subject struct {
	Entity

	Group *Ticket_Group `json:"group:omitempty"`
	Id    *int          `json:"id:omitempty"`
	Name  *string       `json:"name:omitempty"`
}

type Ticket_Survey struct {
	Entity
}

type Ticket_Type struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
}

type Ticket_Update struct {
	Entity

	ChangeOwnerActivity *string                  `json:"changeOwnerActivity:omitempty"`
	CreateDate          *time.Time               `json:"createDate:omitempty"`
	Editor              *User_Interface          `json:"editor:omitempty"`
	EditorId            *int                     `json:"editorId:omitempty"`
	EditorType          *string                  `json:"editorType:omitempty"`
	Entry               *string                  `json:"entry:omitempty"`
	FileAttachment      []Ticket_Attachment_File `json:"fileAttachment:omitempty"`
	FileAttachmentCount *uint                    `json:"fileAttachmentCount:omitempty"`
	Id                  *int                     `json:"id:omitempty"`
	Ticket              *Ticket                  `json:"ticket:omitempty"`
	TicketId            *int                     `json:"ticketId:omitempty"`
	Type                *Ticket_Update_Type      `json:"type:omitempty"`
}

type Ticket_Update_Agent struct {
	Ticket_Update
}

type Ticket_Update_Chat struct {
	Ticket_Update

	Chat *Ticket_Chat_Liveperson `json:"chat:omitempty"`
}

type Ticket_Update_Customer struct {
	Ticket_Update
}

type Ticket_Update_Employee struct {
	Ticket_Update

	ResponseRating *int `json:"responseRating:omitempty"`
}

type Ticket_Update_Type struct {
	Entity

	Description *string        `json:"description:omitempty"`
	KeyName     *string        `json:"keyName:omitempty"`
	Ticket      *Ticket_Update `json:"ticket:omitempty"`
}

type User_Access_Facility_Log struct {
	Entity

	Account     *Account                       `json:"account:omitempty"`
	AccountId   *int                           `json:"accountId:omitempty"`
	Datacenter  *Location                      `json:"datacenter:omitempty"`
	Description *string                        `json:"description:omitempty"`
	Hardware    *Hardware                      `json:"hardware:omitempty"`
	HardwareId  *int                           `json:"hardwareId:omitempty"`
	Id          *int                           `json:"id:omitempty"`
	LocationId  *int                           `json:"locationId:omitempty"`
	LogType     *User_Access_Facility_Log_Type `json:"logType:omitempty"`
	TimeIn      *time.Time                     `json:"timeIn:omitempty"`
	TimeOut     *time.Time                     `json:"timeOut:omitempty"`
	Visitor     *Entity                        `json:"visitor:omitempty"`
}

type User_Access_Facility_Log_Type struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type User_Access_Facility_Visitor struct {
	Entity

	CompanyName *string                            `json:"companyName:omitempty"`
	FirstName   *string                            `json:"firstName:omitempty"`
	LastName    *string                            `json:"lastName:omitempty"`
	TypeId      *int                               `json:"typeId:omitempty"`
	VisitorType *User_Access_Facility_Visitor_Type `json:"visitorType:omitempty"`
}

type User_Access_Facility_Visitor_Type struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type User_Customer struct {
	User_Interface

	Account                                  *Account                                      `json:"account:omitempty"`
	AccountId                                *int                                          `json:"accountId:omitempty"`
	ActionCount                              *uint                                         `json:"actionCount:omitempty"`
	Actions                                  []User_Permission_Action                      `json:"actions:omitempty"`
	AdditionalEmailCount                     *uint                                         `json:"additionalEmailCount:omitempty"`
	AdditionalEmails                         []User_Customer_AdditionalEmail               `json:"additionalEmails:omitempty"`
	Address1                                 *string                                       `json:"address1:omitempty"`
	Address2                                 *string                                       `json:"address2:omitempty"`
	Aim                                      *string                                       `json:"aim:omitempty"`
	AlternatePhone                           *string                                       `json:"alternatePhone:omitempty"`
	ApiAuthenticationKeyCount                *uint                                         `json:"apiAuthenticationKeyCount:omitempty"`
	ApiAuthenticationKeys                    []User_Customer_ApiAuthentication             `json:"apiAuthenticationKeys:omitempty"`
	AuthenticationToken                      *Container_User_Authentication_Token          `json:"authenticationToken:omitempty"`
	CdnAccountCount                          *uint                                         `json:"cdnAccountCount:omitempty"`
	CdnAccounts                              []Network_ContentDelivery_Account             `json:"cdnAccounts:omitempty"`
	ChildUserCount                           *uint                                         `json:"childUserCount:omitempty"`
	ChildUsers                               []User_Customer                               `json:"childUsers:omitempty"`
	City                                     *string                                       `json:"city:omitempty"`
	ClosedTicketCount                        *uint                                         `json:"closedTicketCount:omitempty"`
	ClosedTickets                            []Ticket                                      `json:"closedTickets:omitempty"`
	CompanyName                              *string                                       `json:"companyName:omitempty"`
	Country                                  *string                                       `json:"country:omitempty"`
	CreateDate                               *time.Time                                    `json:"createDate:omitempty"`
	DaylightSavingsTimeFlag                  *bool                                         `json:"daylightSavingsTimeFlag:omitempty"`
	DenyAllResourceAccessOnCreateFlag        *bool                                         `json:"denyAllResourceAccessOnCreateFlag:omitempty"`
	DisplayName                              *string                                       `json:"displayName:omitempty"`
	Email                                    *string                                       `json:"email:omitempty"`
	ExternalBindingCount                     *uint                                         `json:"externalBindingCount:omitempty"`
	ExternalBindings                         []User_External_Binding                       `json:"externalBindings:omitempty"`
	FirstName                                *string                                       `json:"firstName:omitempty"`
	ForumPasswordHash                        *string                                       `json:"forumPasswordHash:omitempty"`
	Hardware                                 []Hardware                                    `json:"hardware:omitempty"`
	HardwareCount                            *uint                                         `json:"hardwareCount:omitempty"`
	HardwareNotificationCount                *uint                                         `json:"hardwareNotificationCount:omitempty"`
	HardwareNotifications                    []User_Customer_Notification_Hardware         `json:"hardwareNotifications:omitempty"`
	HasAcknowledgedSupportPolicyFlag         *bool                                         `json:"hasAcknowledgedSupportPolicyFlag:omitempty"`
	HasFullHardwareAccessFlag                *bool                                         `json:"hasFullHardwareAccessFlag:omitempty"`
	HasFullVirtualGuestAccessFlag            *bool                                         `json:"hasFullVirtualGuestAccessFlag:omitempty"`
	Icq                                      *string                                       `json:"icq:omitempty"`
	Id                                       *int                                          `json:"id:omitempty"`
	IpAddressRestriction                     *string                                       `json:"ipAddressRestriction:omitempty"`
	LastName                                 *string                                       `json:"lastName:omitempty"`
	LayoutProfileCount                       *uint                                         `json:"layoutProfileCount:omitempty"`
	LayoutProfiles                           []Layout_Profile                              `json:"layoutProfiles:omitempty"`
	Locale                                   *Locale                                       `json:"locale:omitempty"`
	LocaleId                                 *int                                          `json:"localeId:omitempty"`
	LoginAttemptCount                        *uint                                         `json:"loginAttemptCount:omitempty"`
	LoginAttempts                            []User_Customer_Access_Authentication         `json:"loginAttempts:omitempty"`
	ManagedByFederationFlag                  *bool                                         `json:"managedByFederationFlag:omitempty"`
	ManagedByOpenIdConnectFlag               *bool                                         `json:"managedByOpenIdConnectFlag:omitempty"`
	MobileDeviceCount                        *uint                                         `json:"mobileDeviceCount:omitempty"`
	MobileDevices                            []User_Customer_MobileDevice                  `json:"mobileDevices:omitempty"`
	ModifyDate                               *time.Time                                    `json:"modifyDate:omitempty"`
	Msn                                      *string                                       `json:"msn:omitempty"`
	NameId                                   *string                                       `json:"nameId:omitempty"`
	NotificationSubscriberCount              *uint                                         `json:"notificationSubscriberCount:omitempty"`
	NotificationSubscribers                  []Notification_Subscriber                     `json:"notificationSubscribers:omitempty"`
	OfficePhone                              *string                                       `json:"officePhone:omitempty"`
	OpenIdConnectUserName                    *string                                       `json:"openIdConnectUserName:omitempty"`
	OpenTicketCount                          *uint                                         `json:"openTicketCount:omitempty"`
	OpenTickets                              []Ticket                                      `json:"openTickets:omitempty"`
	OverrideCount                            *uint                                         `json:"overrideCount:omitempty"`
	Overrides                                []Network_Service_Vpn_Overrides               `json:"overrides:omitempty"`
	Parent                                   *User_Customer                                `json:"parent:omitempty"`
	ParentId                                 *int                                          `json:"parentId:omitempty"`
	PasswordExpireDate                       *time.Time                                    `json:"passwordExpireDate:omitempty"`
	PermissionCount                          *uint                                         `json:"permissionCount:omitempty"`
	PermissionSystemVersion                  *int                                          `json:"permissionSystemVersion:omitempty"`
	Permissions                              []User_Customer_CustomerPermission_Permission `json:"permissions:omitempty"`
	PostalCode                               *string                                       `json:"postalCode:omitempty"`
	PptpVpnAllowedFlag                       *bool                                         `json:"pptpVpnAllowedFlag:omitempty"`
	PreferenceCount                          *uint                                         `json:"preferenceCount:omitempty"`
	Preferences                              []User_Preference                             `json:"preferences:omitempty"`
	RoleCount                                *uint                                         `json:"roleCount:omitempty"`
	Roles                                    []User_Permission_Role                        `json:"roles:omitempty"`
	SalesforceUserLink                       *User_Customer_Link                           `json:"salesforceUserLink:omitempty"`
	SavedId                                  *string                                       `json:"savedId:omitempty"`
	SecondaryLoginManagementFlag             *bool                                         `json:"secondaryLoginManagementFlag:omitempty"`
	SecondaryLoginRequiredFlag               *bool                                         `json:"secondaryLoginRequiredFlag:omitempty"`
	SecondaryPasswordModifyDate              *time.Time                                    `json:"secondaryPasswordModifyDate:omitempty"`
	SecondaryPasswordTimeoutDays             *int                                          `json:"secondaryPasswordTimeoutDays:omitempty"`
	SecurityAnswerCount                      *uint                                         `json:"securityAnswerCount:omitempty"`
	SecurityAnswers                          []User_Customer_Security_Answer               `json:"securityAnswers:omitempty"`
	Sms                                      *string                                       `json:"sms:omitempty"`
	SslVpnAllowedFlag                        *bool                                         `json:"sslVpnAllowedFlag:omitempty"`
	State                                    *string                                       `json:"state:omitempty"`
	StatusDate                               *time.Time                                    `json:"statusDate:omitempty"`
	SubscriberCount                          *uint                                         `json:"subscriberCount:omitempty"`
	Subscribers                              []Notification_User_Subscriber                `json:"subscribers:omitempty"`
	SuccessfulLoginCount                     *uint                                         `json:"successfulLoginCount:omitempty"`
	SuccessfulLogins                         []User_Customer_Access_Authentication         `json:"successfulLogins:omitempty"`
	SupportPolicyAcknowledgementRequiredFlag *int                                          `json:"supportPolicyAcknowledgementRequiredFlag:omitempty"`
	SurveyCount                              *uint                                         `json:"surveyCount:omitempty"`
	SurveyRequiredFlag                       *bool                                         `json:"surveyRequiredFlag:omitempty"`
	Surveys                                  []Survey                                      `json:"surveys:omitempty"`
	TicketCount                              *uint                                         `json:"ticketCount:omitempty"`
	Tickets                                  []Ticket                                      `json:"tickets:omitempty"`
	Timezone                                 *Locale_Timezone                              `json:"timezone:omitempty"`
	TimezoneId                               *int                                          `json:"timezoneId:omitempty"`
	UnsuccessfulLoginCount                   *uint                                         `json:"unsuccessfulLoginCount:omitempty"`
	UnsuccessfulLogins                       []User_Customer_Access_Authentication         `json:"unsuccessfulLogins:omitempty"`
	UserLinkCount                            *uint                                         `json:"userLinkCount:omitempty"`
	UserLinks                                []User_Customer_Link                          `json:"userLinks:omitempty"`
	UserStatus                               *User_Customer_Status                         `json:"userStatus:omitempty"`
	UserStatusId                             *int                                          `json:"userStatusId:omitempty"`
	Username                                 *string                                       `json:"username:omitempty"`
	VirtualGuestCount                        *uint                                         `json:"virtualGuestCount:omitempty"`
	VirtualGuests                            []Virtual_Guest                               `json:"virtualGuests:omitempty"`
	VpnManualConfig                          *bool                                         `json:"vpnManualConfig:omitempty"`
	Yahoo                                    *string                                       `json:"yahoo:omitempty"`
}

type User_Customer_Access_Authentication struct {
	Entity

	CreateDate  *time.Time     `json:"createDate:omitempty"`
	IpAddress   *string        `json:"ipAddress:omitempty"`
	SuccessFlag *bool          `json:"successFlag:omitempty"`
	User        *User_Customer `json:"user:omitempty"`
	UserId      *int           `json:"userId:omitempty"`
	Username    *string        `json:"username:omitempty"`
}

type User_Customer_AdditionalEmail struct {
	Entity

	Email  *string        `json:"email:omitempty"`
	User   *User_Customer `json:"user:omitempty"`
	UserId *int           `json:"userId:omitempty"`
}

type User_Customer_ApiAuthentication struct {
	Entity

	AuthenticationKey    *string        `json:"authenticationKey:omitempty"`
	Id                   *int           `json:"id:omitempty"`
	IpAddressRestriction *string        `json:"ipAddressRestriction:omitempty"`
	TimestampKey         *int           `json:"timestampKey:omitempty"`
	User                 *User_Customer `json:"user:omitempty"`
	UserId               *int           `json:"userId:omitempty"`
}

type User_Customer_CustomerPermission_Permission struct {
	Entity

	Key     *string `json:"key:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type User_Customer_External_Binding struct {
	User_External_Binding

	User *User_Customer `json:"user:omitempty"`
}

type User_Customer_External_Binding_Attribute struct {
	User_External_Binding_Attribute
}

type User_Customer_External_Binding_Phone struct {
	User_Customer_External_Binding

	BindingStatus *string `json:"bindingStatus:omitempty"`
	PinLength     *string `json:"pinLength:omitempty"`
}

type User_Customer_External_Binding_Totp struct {
	User_Customer_External_Binding
}

type User_Customer_External_Binding_Type struct {
	User_External_Binding_Type
}

type User_Customer_External_Binding_Vendor struct {
	User_External_Binding_Vendor
}

type User_Customer_External_Binding_Verisign struct {
	User_Customer_External_Binding

	CredentialExpirationDate *string `json:"credentialExpirationDate:omitempty"`
	CredentialLastUpdateDate *string `json:"credentialLastUpdateDate:omitempty"`
	CredentialState          *string `json:"credentialState:omitempty"`
	CredentialType           *string `json:"credentialType:omitempty"`
}

type User_Customer_Invitation struct {
	Entity

	Code                       *string        `json:"code:omitempty"`
	CreateDate                 *time.Time     `json:"createDate:omitempty"`
	CreatorId                  *int           `json:"creatorId:omitempty"`
	CreatorType                *string        `json:"creatorType:omitempty"`
	Email                      *string        `json:"email:omitempty"`
	ExistingBlueIdFlag         *int           `json:"existingBlueIdFlag:omitempty"`
	ExpirationDate             *time.Time     `json:"expirationDate:omitempty"`
	Id                         *int           `json:"id:omitempty"`
	IsFederatedEmailDomainFlag *int           `json:"isFederatedEmailDomainFlag:omitempty"`
	ModifyDate                 *time.Time     `json:"modifyDate:omitempty"`
	ResponseDate               *time.Time     `json:"responseDate:omitempty"`
	StatusId                   *int           `json:"statusId:omitempty"`
	User                       *User_Customer `json:"user:omitempty"`
	UserId                     *int           `json:"userId:omitempty"`
}

type User_Customer_Link struct {
	Entity

	CreateDate                    *time.Time        `json:"createDate:omitempty"`
	DefaultFlag                   *int              `json:"defaultFlag:omitempty"`
	DestinationUserAlphanumericId *string           `json:"destinationUserAlphanumericId:omitempty"`
	DestinationUserId             *int              `json:"destinationUserId:omitempty"`
	Id                            *int              `json:"id:omitempty"`
	ServiceProvider               *Service_Provider `json:"serviceProvider:omitempty"`
	ServiceProviderId             *int              `json:"serviceProviderId:omitempty"`
	User                          *User_Customer    `json:"user:omitempty"`
	UserId                        *int              `json:"userId:omitempty"`
}

type User_Customer_Link_ThePlanet struct {
	User_Customer_Link
}

type User_Customer_MobileDevice struct {
	Entity

	AvailablePushNotificationSubscriptionCount *uint                                       `json:"availablePushNotificationSubscriptionCount:omitempty"`
	AvailablePushNotificationSubscriptions     []Notification                              `json:"availablePushNotificationSubscriptions:omitempty"`
	CreateDate                                 *time.Time                                  `json:"createDate:omitempty"`
	Customer                                   *User_Customer                              `json:"customer:omitempty"`
	DisplayResolutionXxY                       *string                                     `json:"displayResolutionXxY:omitempty"`
	Id                                         *int                                        `json:"id:omitempty"`
	MobileDeviceTypeId                         *int                                        `json:"mobileDeviceTypeId:omitempty"`
	MobileOperatingSystemId                    *int                                        `json:"mobileOperatingSystemId:omitempty"`
	ModelNumber                                *string                                     `json:"modelNumber:omitempty"`
	ModifyDate                                 *time.Time                                  `json:"modifyDate:omitempty"`
	Name                                       *string                                     `json:"name:omitempty"`
	OperatingSystem                            *User_Customer_MobileDevice_OperatingSystem `json:"operatingSystem:omitempty"`
	PhoneNumber                                *string                                     `json:"phoneNumber:omitempty"`
	PushNotificationSubscriptionCount          *uint                                       `json:"pushNotificationSubscriptionCount:omitempty"`
	PushNotificationSubscriptions              []Notification_User_Subscriber              `json:"pushNotificationSubscriptions:omitempty"`
	SerialNumber                               *string                                     `json:"serialNumber:omitempty"`
	Token                                      *string                                     `json:"token:omitempty"`
	Type                                       *User_Customer_MobileDevice_Type            `json:"type:omitempty"`
	UserId                                     *int                                        `json:"userId:omitempty"`
}

type User_Customer_MobileDevice_OperatingSystem struct {
	Entity

	BuildVersion *int       `json:"buildVersion:omitempty"`
	CreateDate   *time.Time `json:"createDate:omitempty"`
	Description  *string    `json:"description:omitempty"`
	Id           *int       `json:"id:omitempty"`
	MajorVersion *int       `json:"majorVersion:omitempty"`
	MinorVersion *int       `json:"minorVersion:omitempty"`
	ModifyDate   *time.Time `json:"modifyDate:omitempty"`
	Name         *string    `json:"name:omitempty"`
}

type User_Customer_MobileDevice_Type struct {
	Entity

	CreateDate  *time.Time `json:"createDate:omitempty"`
	Description *string    `json:"description:omitempty"`
	Id          *int       `json:"id:omitempty"`
	ModifyDate  *time.Time `json:"modifyDate:omitempty"`
	Name        *string    `json:"name:omitempty"`
}

type User_Customer_Notification_Hardware struct {
	Entity

	Hardware   *Hardware      `json:"hardware:omitempty"`
	HardwareId *int           `json:"hardwareId:omitempty"`
	Id         *int           `json:"id:omitempty"`
	User       *User_Customer `json:"user:omitempty"`
	UserId     *int           `json:"userId:omitempty"`
}

type User_Customer_Notification_Virtual_Guest struct {
	Entity

	Guest   *Virtual_Guest `json:"guest:omitempty"`
	GuestId *int           `json:"guestId:omitempty"`
	Id      *int           `json:"id:omitempty"`
	User    *User_Customer `json:"user:omitempty"`
	UserId  *int           `json:"userId:omitempty"`
}

type User_Customer_OpenIdConnect struct {
	User_Customer
}

type User_Customer_Prospect struct {
	Entity

	Account               *Account                     `json:"account:omitempty"`
	AssignedEmployeeCount *uint                        `json:"assignedEmployeeCount:omitempty"`
	AssignedEmployees     []User_Employee              `json:"assignedEmployees:omitempty"`
	QuoteCount            *uint                        `json:"quoteCount:omitempty"`
	Quotes                []Billing_Order_Quote        `json:"quotes:omitempty"`
	Type                  *User_Customer_Prospect_Type `json:"type:omitempty"`
}

type User_Customer_Prospect_ServiceProvider_EnrollRequest struct {
	Entity

	AccountId                   *int                   `json:"accountId:omitempty"`
	Address1                    *string                `json:"address1:omitempty"`
	Address2                    *string                `json:"address2:omitempty"`
	CardAccountNumber           *string                `json:"cardAccountNumber:omitempty"`
	CardExpirationMonth         *string                `json:"cardExpirationMonth:omitempty"`
	CardExpirationYear          *string                `json:"cardExpirationYear:omitempty"`
	CardType                    *string                `json:"cardType:omitempty"`
	CardVerificationNumber      *string                `json:"cardVerificationNumber:omitempty"`
	City                        *string                `json:"city:omitempty"`
	CompanyName                 *string                `json:"companyName:omitempty"`
	CompanyType                 *Catalyst_Company_Type `json:"companyType:omitempty"`
	CompanyTypeId               *int                   `json:"companyTypeId:omitempty"`
	CompanyUrl                  *string                `json:"companyUrl:omitempty"`
	ContactEmail                *string                `json:"contactEmail:omitempty"`
	ContactFirstName            *string                `json:"contactFirstName:omitempty"`
	ContactLastName             *string                `json:"contactLastName:omitempty"`
	ContactPhone                *string                `json:"contactPhone:omitempty"`
	Country                     *string                `json:"country:omitempty"`
	CustomerProspectId          *int                   `json:"customerProspectId:omitempty"`
	DeviceFingerprintId         *string                `json:"deviceFingerprintId:omitempty"`
	Email                       *string                `json:"email:omitempty"`
	ExistingCustomerFlag        *bool                  `json:"existingCustomerFlag:omitempty"`
	FirstName                   *string                `json:"firstName:omitempty"`
	IbmPartnerWorldId           *string                `json:"ibmPartnerWorldId:omitempty"`
	IbmPartnerWorldMemberFlag   *bool                  `json:"ibmPartnerWorldMemberFlag:omitempty"`
	LastName                    *string                `json:"lastName:omitempty"`
	MasterAgreementCompleteFlag *bool                  `json:"masterAgreementCompleteFlag:omitempty"`
	OfficePhone                 *string                `json:"officePhone:omitempty"`
	PostalCode                  *string                `json:"postalCode:omitempty"`
	ServiceProviderAddendumFlag *bool                  `json:"serviceProviderAddendumFlag:omitempty"`
	State                       *string                `json:"state:omitempty"`
	SurveyResponses             []Survey_Response      `json:"surveyResponses:omitempty"`
	VatId                       *string                `json:"vatId:omitempty"`
}

type User_Customer_Prospect_Type struct {
	Entity

	CreateDate  *time.Time `json:"createDate:omitempty"`
	Description *string    `json:"description:omitempty"`
	Id          *int       `json:"id:omitempty"`
	KeyName     *string    `json:"keyName:omitempty"`
	ModifyDate  *time.Time `json:"modifyDate:omitempty"`
	Name        *string    `json:"name:omitempty"`
}

type User_Customer_Security_Answer struct {
	Entity

	Answer     *string                 `json:"answer:omitempty"`
	Id         *int                    `json:"id:omitempty"`
	Question   *User_Security_Question `json:"question:omitempty"`
	QuestionId *int                    `json:"questionId:omitempty"`
	User       *User_Customer          `json:"user:omitempty"`
	UserId     *int                    `json:"userId:omitempty"`
}

type User_Customer_Status struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type User_Employee struct {
	User_Interface

	ActionCount                    *uint                     `json:"actionCount:omitempty"`
	Actions                        []User_Permission_Action  `json:"actions:omitempty"`
	ChatTranscript                 []Ticket_Chat             `json:"chatTranscript:omitempty"`
	ChatTranscriptCount            *uint                     `json:"chatTranscriptCount:omitempty"`
	DisplayName                    *string                   `json:"displayName:omitempty"`
	Email                          *string                   `json:"email:omitempty"`
	EmployeeDepartment             *User_Employee_Department `json:"employeeDepartment:omitempty"`
	EmployeeDepartmentId           *int                      `json:"employeeDepartmentId:omitempty"`
	FirstName                      *string                   `json:"firstName:omitempty"`
	LastName                       *string                   `json:"lastName:omitempty"`
	LayoutProfileCount             *uint                     `json:"layoutProfileCount:omitempty"`
	LayoutProfiles                 []Layout_Profile          `json:"layoutProfiles:omitempty"`
	MetricTrackingObject           *Metric_Tracking_Object   `json:"metricTrackingObject:omitempty"`
	OfficePhone                    *string                   `json:"officePhone:omitempty"`
	RoleCount                      *uint                     `json:"roleCount:omitempty"`
	Roles                          []User_Permission_Role    `json:"roles:omitempty"`
	TicketActivities               []Ticket_Activity         `json:"ticketActivities:omitempty"`
	TicketActivityCount            *uint                     `json:"ticketActivityCount:omitempty"`
	TicketAttachmentReferenceCount *uint                     `json:"ticketAttachmentReferenceCount:omitempty"`
	TicketAttachmentReferences     []Ticket_Attachment       `json:"ticketAttachmentReferences:omitempty"`
	Username                       *string                   `json:"username:omitempty"`
}

type User_Employee_Department struct {
	Entity

	Name *string `json:"name:omitempty"`
}

type User_External_Binding struct {
	Entity

	Active         *bool                             `json:"active:omitempty"`
	AttributeCount *uint                             `json:"attributeCount:omitempty"`
	Attributes     []User_External_Binding_Attribute `json:"attributes:omitempty"`
	BillingItem    *Billing_Item                     `json:"billingItem:omitempty"`
	CreateDate     *time.Time                        `json:"createDate:omitempty"`
	ExternalId     *string                           `json:"externalId:omitempty"`
	Id             *int                              `json:"id:omitempty"`
	Note           *string                           `json:"note:omitempty"`
	Password       *string                           `json:"password:omitempty"`
	Type           *User_External_Binding_Type       `json:"type:omitempty"`
	TypeId         *int                              `json:"typeId:omitempty"`
	UserId         *int                              `json:"userId:omitempty"`
	Vendor         *User_External_Binding_Vendor     `json:"vendor:omitempty"`
	VendorId       *int                              `json:"vendorId:omitempty"`
}

type User_External_Binding_Attribute struct {
	Entity

	ExternalBinding *User_External_Binding `json:"externalBinding:omitempty"`
	Value           *string                `json:"value:omitempty"`
}

type User_External_Binding_Type struct {
	Entity

	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type User_External_Binding_Vendor struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type User_Interface struct {
	Entity
}

type User_Permission_Action struct {
	Entity

	CreateDate  *time.Time `json:"createDate:omitempty"`
	Description *string    `json:"description:omitempty"`
	Id          *int       `json:"id:omitempty"`
	KeyName     *string    `json:"keyName:omitempty"`
	ModifyDate  *time.Time `json:"modifyDate:omitempty"`
	Name        *string    `json:"name:omitempty"`
}

type User_Permission_Group struct {
	Entity

	Account        *Account                    `json:"account:omitempty"`
	AccountId      *int                        `json:"accountId:omitempty"`
	ActionCount    *uint                       `json:"actionCount:omitempty"`
	Actions        []User_Permission_Action    `json:"actions:omitempty"`
	CreateDate     *time.Time                  `json:"createDate:omitempty"`
	Description    *string                     `json:"description:omitempty"`
	ExpirationDate *time.Time                  `json:"expirationDate:omitempty"`
	Id             *int                        `json:"id:omitempty"`
	ModifyDate     *time.Time                  `json:"modifyDate:omitempty"`
	Name           *string                     `json:"name:omitempty"`
	RoleCount      *uint                       `json:"roleCount:omitempty"`
	Roles          []User_Permission_Role      `json:"roles:omitempty"`
	Type           *User_Permission_Group_Type `json:"type:omitempty"`
	TypeId         *int                        `json:"typeId:omitempty"`
}

type User_Permission_Group_Type struct {
	Entity

	CreateDate *time.Time              `json:"createDate:omitempty"`
	GroupCount *uint                   `json:"groupCount:omitempty"`
	Groups     []User_Permission_Group `json:"groups:omitempty"`
	Id         *int                    `json:"id:omitempty"`
	KeyName    *string                 `json:"keyName:omitempty"`
	ModifyDate *time.Time              `json:"modifyDate:omitempty"`
	Name       *string                 `json:"name:omitempty"`
}

type User_Permission_Role struct {
	Entity

	Account            *Account                 `json:"account:omitempty"`
	AccountId          *int                     `json:"accountId:omitempty"`
	ActionCount        *uint                    `json:"actionCount:omitempty"`
	Actions            []User_Permission_Action `json:"actions:omitempty"`
	CreateDate         *time.Time               `json:"createDate:omitempty"`
	Description        *string                  `json:"description:omitempty"`
	GroupCount         *uint                    `json:"groupCount:omitempty"`
	Groups             []User_Permission_Group  `json:"groups:omitempty"`
	Id                 *int                     `json:"id:omitempty"`
	ModifyDate         *time.Time               `json:"modifyDate:omitempty"`
	Name               *string                  `json:"name:omitempty"`
	NewUserDefaultFlag *int                     `json:"newUserDefaultFlag:omitempty"`
	SystemFlag         *int                     `json:"systemFlag:omitempty"`
	UserCount          *uint                    `json:"userCount:omitempty"`
	Users              []User_Customer          `json:"users:omitempty"`
}

type User_Preference struct {
	Entity

	Description *string               `json:"description:omitempty"`
	Type        *User_Preference_Type `json:"type:omitempty"`
	Value       *string               `json:"value:omitempty"`
}

type User_Preference_Type struct {
	Entity

	Description  *string `json:"description:omitempty"`
	KeyName      *string `json:"keyName:omitempty"`
	Name         *string `json:"name:omitempty"`
	ValueExample *string `json:"valueExample:omitempty"`
}

type User_Security_Question struct {
	Entity

	DisplayOrder *int    `json:"displayOrder:omitempty"`
	Id           *int    `json:"id:omitempty"`
	Question     *string `json:"question:omitempty"`
	Viewable     *int    `json:"viewable:omitempty"`
}

type Utility_Bandwidth_Graph struct {
	Entity
}

type Utility_Network struct {
	Entity
}

type Utility_ObjectFilter struct {
	Entity
}

type Utility_ObjectFilter_Operation struct {
	Entity
}

type Utility_ObjectFilter_Operation_Option struct {
	Entity
}

type Virtual_Disk_Image struct {
	Entity

	BillingItem             *Billing_Item_Virtual_Disk_Image     `json:"billingItem:omitempty"`
	BlockDeviceCount        *uint                                `json:"blockDeviceCount:omitempty"`
	BlockDevices            []Virtual_Guest_Block_Device         `json:"blockDevices:omitempty"`
	BootableVolumeFlag      *bool                                `json:"bootableVolumeFlag:omitempty"`
	Capacity                *int                                 `json:"capacity:omitempty"`
	Checksum                *string                              `json:"checksum:omitempty"`
	CoalescedDiskImageCount *uint                                `json:"coalescedDiskImageCount:omitempty"`
	CoalescedDiskImages     []Virtual_Disk_Image                 `json:"coalescedDiskImages:omitempty"`
	CopyOnWriteFlag         *bool                                `json:"copyOnWriteFlag:omitempty"`
	CreateDate              *time.Time                           `json:"createDate:omitempty"`
	Description             *string                              `json:"description:omitempty"`
	Id                      *int                                 `json:"id:omitempty"`
	LocalDiskFlag           *bool                                `json:"localDiskFlag:omitempty"`
	MetadataFlag            *bool                                `json:"metadataFlag:omitempty"`
	ModifyDate              *time.Time                           `json:"modifyDate:omitempty"`
	Name                    *string                              `json:"name:omitempty"`
	ParentId                *int                                 `json:"parentId:omitempty"`
	SoftwareReferenceCount  *uint                                `json:"softwareReferenceCount:omitempty"`
	SoftwareReferences      []Virtual_Disk_Image_Software        `json:"softwareReferences:omitempty"`
	SourceDiskImage         *Virtual_Disk_Image                  `json:"sourceDiskImage:omitempty"`
	StorageRepository       *Virtual_Storage_Repository          `json:"storageRepository:omitempty"`
	StorageRepositoryId     *int                                 `json:"storageRepositoryId:omitempty"`
	StorageRepositoryType   *Virtual_Storage_Repository_Type     `json:"storageRepositoryType:omitempty"`
	TemplateBlockDevice     *Virtual_Guest_Block_Device_Template `json:"templateBlockDevice:omitempty"`
	Type                    *Virtual_Disk_Image_Type             `json:"type:omitempty"`
	TypeId                  *int                                 `json:"typeId:omitempty"`
	Units                   *string                              `json:"units:omitempty"`
	Uuid                    *string                              `json:"uuid:omitempty"`
}

type Virtual_Disk_Image_Software struct {
	Entity

	DiskImage             *Virtual_Disk_Image                    `json:"diskImage:omitempty"`
	Id                    *int                                   `json:"id:omitempty"`
	PasswordCount         *uint                                  `json:"passwordCount:omitempty"`
	Passwords             []Virtual_Disk_Image_Software_Password `json:"passwords:omitempty"`
	SoftwareDescription   *Software_Description                  `json:"softwareDescription:omitempty"`
	SoftwareDescriptionId *int                                   `json:"softwareDescriptionId:omitempty"`
}

type Virtual_Disk_Image_Software_Password struct {
	Entity

	Password *string                      `json:"password:omitempty"`
	Software *Virtual_Disk_Image_Software `json:"software:omitempty"`
	Username *string                      `json:"username:omitempty"`
}

type Virtual_Disk_Image_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Virtual_Guest struct {
	Entity

	Account                                   *Account                                       `json:"account:omitempty"`
	AccountId                                 *int                                           `json:"accountId:omitempty"`
	AccountOwnedPoolFlag                      *bool                                          `json:"accountOwnedPoolFlag:omitempty"`
	ActiveNetworkMonitorIncident              []Network_Monitor_Version1_Incident            `json:"activeNetworkMonitorIncident:omitempty"`
	ActiveNetworkMonitorIncidentCount         *uint                                          `json:"activeNetworkMonitorIncidentCount:omitempty"`
	ActiveTicketCount                         *uint                                          `json:"activeTicketCount:omitempty"`
	ActiveTickets                             []Ticket                                       `json:"activeTickets:omitempty"`
	ActiveTransaction                         *Provisioning_Version1_Transaction             `json:"activeTransaction:omitempty"`
	ActiveTransactionCount                    *uint                                          `json:"activeTransactionCount:omitempty"`
	ActiveTransactions                        []Provisioning_Version1_Transaction            `json:"activeTransactions:omitempty"`
	AllowedHost                               *Network_Storage_Allowed_Host                  `json:"allowedHost:omitempty"`
	AllowedNetworkStorage                     []Network_Storage                              `json:"allowedNetworkStorage:omitempty"`
	AllowedNetworkStorageCount                *uint                                          `json:"allowedNetworkStorageCount:omitempty"`
	AllowedNetworkStorageReplicaCount         *uint                                          `json:"allowedNetworkStorageReplicaCount:omitempty"`
	AllowedNetworkStorageReplicas             []Network_Storage                              `json:"allowedNetworkStorageReplicas:omitempty"`
	AntivirusSpywareSoftwareComponent         *Software_Component                            `json:"antivirusSpywareSoftwareComponent:omitempty"`
	ApplicationDeliveryController             *Network_Application_Delivery_Controller       `json:"applicationDeliveryController:omitempty"`
	AttributeCount                            *uint                                          `json:"attributeCount:omitempty"`
	Attributes                                []Virtual_Guest_Attribute                      `json:"attributes:omitempty"`
	AvailableMonitoring                       []Network_Monitor_Version1_Query_Host_Stratum  `json:"availableMonitoring:omitempty"`
	AvailableMonitoringCount                  *uint                                          `json:"availableMonitoringCount:omitempty"`
	AverageDailyPrivateBandwidthUsage         *float64                                       `json:"averageDailyPrivateBandwidthUsage:omitempty"`
	AverageDailyPublicBandwidthUsage          *float64                                       `json:"averageDailyPublicBandwidthUsage:omitempty"`
	BackendNetworkComponentCount              *uint                                          `json:"backendNetworkComponentCount:omitempty"`
	BackendNetworkComponents                  []Virtual_Guest_Network_Component              `json:"backendNetworkComponents:omitempty"`
	BackendRouterCount                        *uint                                          `json:"backendRouterCount:omitempty"`
	BackendRouters                            []Hardware                                     `json:"backendRouters:omitempty"`
	BandwidthAllocation                       *float64                                       `json:"bandwidthAllocation:omitempty"`
	BandwidthAllotmentDetail                  *Network_Bandwidth_Version1_Allotment_Detail   `json:"bandwidthAllotmentDetail:omitempty"`
	BillingCycleBandwidthUsage                []Network_Bandwidth_Usage                      `json:"billingCycleBandwidthUsage:omitempty"`
	BillingCycleBandwidthUsageCount           *uint                                          `json:"billingCycleBandwidthUsageCount:omitempty"`
	BillingCyclePrivateBandwidthUsage         *Network_Bandwidth_Usage                       `json:"billingCyclePrivateBandwidthUsage:omitempty"`
	BillingCyclePublicBandwidthUsage          *Network_Bandwidth_Usage                       `json:"billingCyclePublicBandwidthUsage:omitempty"`
	BillingItem                               *Billing_Item_Virtual_Guest                    `json:"billingItem:omitempty"`
	BlockCancelBecauseDisconnectedFlag        *bool                                          `json:"blockCancelBecauseDisconnectedFlag:omitempty"`
	BlockDeviceCount                          *uint                                          `json:"blockDeviceCount:omitempty"`
	BlockDeviceTemplateGroup                  *Virtual_Guest_Block_Device_Template_Group     `json:"blockDeviceTemplateGroup:omitempty"`
	BlockDevices                              []Virtual_Guest_Block_Device                   `json:"blockDevices:omitempty"`
	ConsoleIpAddressFlag                      *bool                                          `json:"consoleIpAddressFlag:omitempty"`
	ConsoleIpAddressRecord                    *Virtual_Guest_Network_Component_IpAddress     `json:"consoleIpAddressRecord:omitempty"`
	ContinuousDataProtectionSoftwareComponent *Software_Component                            `json:"continuousDataProtectionSoftwareComponent:omitempty"`
	ControlPanel                              *Software_Component                            `json:"controlPanel:omitempty"`
	CreateDate                                *time.Time                                     `json:"createDate:omitempty"`
	CurrentBandwidthSummary                   *Metric_Tracking_Object_Bandwidth_Summary      `json:"currentBandwidthSummary:omitempty"`
	Datacenter                                *Location                                      `json:"datacenter:omitempty"`
	DedicatedAccountHostOnlyFlag              *bool                                          `json:"dedicatedAccountHostOnlyFlag:omitempty"`
	Domain                                    *string                                        `json:"domain:omitempty"`
	EvaultNetworkStorage                      []Network_Storage                              `json:"evaultNetworkStorage:omitempty"`
	EvaultNetworkStorageCount                 *uint                                          `json:"evaultNetworkStorageCount:omitempty"`
	FirewallServiceComponent                  *Network_Component_Firewall                    `json:"firewallServiceComponent:omitempty"`
	FrontendNetworkComponentCount             *uint                                          `json:"frontendNetworkComponentCount:omitempty"`
	FrontendNetworkComponents                 []Virtual_Guest_Network_Component              `json:"frontendNetworkComponents:omitempty"`
	FrontendRouters                           *Hardware                                      `json:"frontendRouters:omitempty"`
	FullyQualifiedDomainName                  *string                                        `json:"fullyQualifiedDomainName:omitempty"`
	GlobalIdentifier                          *string                                        `json:"globalIdentifier:omitempty"`
	GuestBootParameter                        *Virtual_Guest_Boot_Parameter                  `json:"guestBootParameter:omitempty"`
	Host                                      *Virtual_Host                                  `json:"host:omitempty"`
	HostIpsSoftwareComponent                  *Software_Component                            `json:"hostIpsSoftwareComponent:omitempty"`
	Hostname                                  *string                                        `json:"hostname:omitempty"`
	HourlyBillingFlag                         *bool                                          `json:"hourlyBillingFlag:omitempty"`
	Id                                        *int                                           `json:"id:omitempty"`
	InboundPrivateBandwidthUsage              *float64                                       `json:"inboundPrivateBandwidthUsage:omitempty"`
	InboundPublicBandwidthUsage               *float64                                       `json:"inboundPublicBandwidthUsage:omitempty"`
	InternalTagReferenceCount                 *uint                                          `json:"internalTagReferenceCount:omitempty"`
	InternalTagReferences                     []Tag_Reference                                `json:"internalTagReferences:omitempty"`
	LastKnownPowerState                       *Virtual_Guest_Power_State                     `json:"lastKnownPowerState:omitempty"`
	LastOperatingSystemReload                 *Provisioning_Version1_Transaction             `json:"lastOperatingSystemReload:omitempty"`
	LastPowerStateId                          *int                                           `json:"lastPowerStateId:omitempty"`
	LastTransaction                           *Provisioning_Version1_Transaction             `json:"lastTransaction:omitempty"`
	LastVerifiedDate                          *time.Time                                     `json:"lastVerifiedDate:omitempty"`
	LatestNetworkMonitorIncident              *Network_Monitor_Version1_Incident             `json:"latestNetworkMonitorIncident:omitempty"`
	LocalDiskFlag                             *bool                                          `json:"localDiskFlag:omitempty"`
	Location                                  *Location                                      `json:"location:omitempty"`
	ManagedResourceFlag                       *bool                                          `json:"managedResourceFlag:omitempty"`
	MaxCpu                                    *int                                           `json:"maxCpu:omitempty"`
	MaxCpuUnits                               *string                                        `json:"maxCpuUnits:omitempty"`
	MaxMemory                                 *int                                           `json:"maxMemory:omitempty"`
	MetricPollDate                            *time.Time                                     `json:"metricPollDate:omitempty"`
	MetricTrackingObject                      *Metric_Tracking_Object                        `json:"metricTrackingObject:omitempty"`
	MetricTrackingObjectId                    *int                                           `json:"metricTrackingObjectId:omitempty"`
	ModifyDate                                *time.Time                                     `json:"modifyDate:omitempty"`
	MonitoringAgentCount                      *uint                                          `json:"monitoringAgentCount:omitempty"`
	MonitoringAgents                          []Monitoring_Agent                             `json:"monitoringAgents:omitempty"`
	MonitoringRobot                           *Monitoring_Robot                              `json:"monitoringRobot:omitempty"`
	MonitoringServiceComponent                *Network_Monitor_Version1_Query_Host_Stratum   `json:"monitoringServiceComponent:omitempty"`
	MonitoringServiceEligibilityFlag          *bool                                          `json:"monitoringServiceEligibilityFlag:omitempty"`
	MonitoringServiceFlag                     *bool                                          `json:"monitoringServiceFlag:omitempty"`
	MonitoringUserNotification                []User_Customer_Notification_Virtual_Guest     `json:"monitoringUserNotification:omitempty"`
	MonitoringUserNotificationCount           *uint                                          `json:"monitoringUserNotificationCount:omitempty"`
	NetworkComponentCount                     *uint                                          `json:"networkComponentCount:omitempty"`
	NetworkComponents                         []Virtual_Guest_Network_Component              `json:"networkComponents:omitempty"`
	NetworkMonitorCount                       *uint                                          `json:"networkMonitorCount:omitempty"`
	NetworkMonitorIncidentCount               *uint                                          `json:"networkMonitorIncidentCount:omitempty"`
	NetworkMonitorIncidents                   []Network_Monitor_Version1_Incident            `json:"networkMonitorIncidents:omitempty"`
	NetworkMonitors                           []Network_Monitor_Version1_Query_Host          `json:"networkMonitors:omitempty"`
	NetworkStorage                            []Network_Storage                              `json:"networkStorage:omitempty"`
	NetworkStorageCount                       *uint                                          `json:"networkStorageCount:omitempty"`
	NetworkVlanCount                          *uint                                          `json:"networkVlanCount:omitempty"`
	NetworkVlans                              []Network_Vlan                                 `json:"networkVlans:omitempty"`
	Notes                                     *string                                        `json:"notes:omitempty"`
	OpenCancellationTicket                    *Ticket                                        `json:"openCancellationTicket:omitempty"`
	OperatingSystem                           *Software_Component_OperatingSystem            `json:"operatingSystem:omitempty"`
	OperatingSystemReferenceCode              *string                                        `json:"operatingSystemReferenceCode:omitempty"`
	OrderedPackageId                          *string                                        `json:"orderedPackageId:omitempty"`
	OutboundPrivateBandwidthUsage             *float64                                       `json:"outboundPrivateBandwidthUsage:omitempty"`
	OutboundPublicBandwidthUsage              *float64                                       `json:"outboundPublicBandwidthUsage:omitempty"`
	OverBandwidthAllocationFlag               *int                                           `json:"overBandwidthAllocationFlag:omitempty"`
	PostInstallScriptUri                      *string                                        `json:"postInstallScriptUri:omitempty"`
	PowerState                                *Virtual_Guest_Power_State                     `json:"powerState:omitempty"`
	PrimaryBackendIpAddress                   *string                                        `json:"primaryBackendIpAddress:omitempty"`
	PrimaryBackendNetworkComponent            *Virtual_Guest_Network_Component               `json:"primaryBackendNetworkComponent:omitempty"`
	PrimaryIpAddress                          *string                                        `json:"primaryIpAddress:omitempty"`
	PrimaryNetworkComponent                   *Virtual_Guest_Network_Component               `json:"primaryNetworkComponent:omitempty"`
	PrivateNetworkOnlyFlag                    *bool                                          `json:"privateNetworkOnlyFlag:omitempty"`
	ProjectedOverBandwidthAllocationFlag      *int                                           `json:"projectedOverBandwidthAllocationFlag:omitempty"`
	ProjectedPublicBandwidthUsage             *float64                                       `json:"projectedPublicBandwidthUsage:omitempty"`
	ProvisionDate                             *time.Time                                     `json:"provisionDate:omitempty"`
	RecentEventCount                          *uint                                          `json:"recentEventCount:omitempty"`
	RecentEvents                              []Notification_Occurrence_Event                `json:"recentEvents:omitempty"`
	RegionalGroup                             *Location_Group_Regional                       `json:"regionalGroup:omitempty"`
	RegionalInternetRegistry                  *Network_Regional_Internet_Registry            `json:"regionalInternetRegistry:omitempty"`
	ScaleAssetCount                           *uint                                          `json:"scaleAssetCount:omitempty"`
	ScaleAssets                               []Scale_Asset                                  `json:"scaleAssets:omitempty"`
	ScaleMember                               *Scale_Member_Virtual_Guest                    `json:"scaleMember:omitempty"`
	ScaledFlag                                *bool                                          `json:"scaledFlag:omitempty"`
	SecurityScanRequestCount                  *uint                                          `json:"securityScanRequestCount:omitempty"`
	SecurityScanRequests                      []Network_Security_Scanner_Request             `json:"securityScanRequests:omitempty"`
	ServerRoom                                *Location                                      `json:"serverRoom:omitempty"`
	SoftwareComponentCount                    *uint                                          `json:"softwareComponentCount:omitempty"`
	SoftwareComponents                        []Software_Component                           `json:"softwareComponents:omitempty"`
	SshKeyCount                               *uint                                          `json:"sshKeyCount:omitempty"`
	SshKeys                                   []Security_Ssh_Key                             `json:"sshKeys:omitempty"`
	StartCpus                                 *int                                           `json:"startCpus:omitempty"`
	Status                                    *Virtual_Guest_Status                          `json:"status:omitempty"`
	StatusId                                  *int                                           `json:"statusId:omitempty"`
	SupplementalCreateObjectOptions           *Virtual_Guest_SupplementalCreateObjectOptions `json:"supplementalCreateObjectOptions:omitempty"`
	TagReferenceCount                         *uint                                          `json:"tagReferenceCount:omitempty"`
	TagReferences                             []Tag_Reference                                `json:"tagReferences:omitempty"`
	UpgradeRequest                            *Product_Upgrade_Request                       `json:"upgradeRequest:omitempty"`
	UserCount                                 *uint                                          `json:"userCount:omitempty"`
	UserData                                  []Virtual_Guest_Attribute                      `json:"userData:omitempty"`
	UserDataCount                             *uint                                          `json:"userDataCount:omitempty"`
	Users                                     []User_Customer                                `json:"users:omitempty"`
	Uuid                                      *string                                        `json:"uuid:omitempty"`
	VirtualRack                               *Network_Bandwidth_Version1_Allotment          `json:"virtualRack:omitempty"`
	VirtualRackId                             *int                                           `json:"virtualRackId:omitempty"`
	VirtualRackName                           *string                                        `json:"virtualRackName:omitempty"`
}

type Virtual_Guest_Attribute struct {
	Entity

	Guest *Virtual_Guest                `json:"guest:omitempty"`
	Type  *Virtual_Guest_Attribute_Type `json:"type:omitempty"`
	Value *string                       `json:"value:omitempty"`
}

type Virtual_Guest_Attribute_Type struct {
	Entity

	Keyname *string `json:"keyname:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Virtual_Guest_Attribute_UserData struct {
	Virtual_Guest_Attribute
}

type Virtual_Guest_Block_Device struct {
	Entity

	BootableFlag *int                               `json:"bootableFlag:omitempty"`
	CreateDate   *time.Time                         `json:"createDate:omitempty"`
	Device       *string                            `json:"device:omitempty"`
	DiskImage    *Virtual_Disk_Image                `json:"diskImage:omitempty"`
	DiskImageId  *int                               `json:"diskImageId:omitempty"`
	Guest        *Virtual_Guest                     `json:"guest:omitempty"`
	GuestId      *int                               `json:"guestId:omitempty"`
	HotPlugFlag  *int                               `json:"hotPlugFlag:omitempty"`
	Id           *int                               `json:"id:omitempty"`
	ModifyDate   *time.Time                         `json:"modifyDate:omitempty"`
	MountMode    *string                            `json:"mountMode:omitempty"`
	MountType    *string                            `json:"mountType:omitempty"`
	Status       *Virtual_Guest_Block_Device_Status `json:"status:omitempty"`
	StatusId     *int                               `json:"statusId:omitempty"`
	Uuid         *string                            `json:"uuid:omitempty"`
}

type Virtual_Guest_Block_Device_Status struct {
	Entity

	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Virtual_Guest_Block_Device_Template struct {
	Entity

	Device      *string                                    `json:"device:omitempty"`
	DiskImage   *Virtual_Disk_Image                        `json:"diskImage:omitempty"`
	DiskImageId *int                                       `json:"diskImageId:omitempty"`
	DiskSpace   *float64                                   `json:"diskSpace:omitempty"`
	Group       *Virtual_Guest_Block_Device_Template_Group `json:"group:omitempty"`
	GroupId     *int                                       `json:"groupId:omitempty"`
	Id          *int                                       `json:"id:omitempty"`
	Units       *string                                    `json:"units:omitempty"`
}

type Virtual_Guest_Block_Device_Template_Group struct {
	Entity

	Account                    *Account                                             `json:"account:omitempty"`
	AccountContactCount        *uint                                                `json:"accountContactCount:omitempty"`
	AccountContacts            []Account_Contact                                    `json:"accountContacts:omitempty"`
	AccountId                  *int                                                 `json:"accountId:omitempty"`
	AccountReferenceCount      *uint                                                `json:"accountReferenceCount:omitempty"`
	AccountReferences          []Virtual_Guest_Block_Device_Template_Group_Accounts `json:"accountReferences:omitempty"`
	BlockDeviceCount           *uint                                                `json:"blockDeviceCount:omitempty"`
	BlockDevices               []Virtual_Guest_Block_Device_Template                `json:"blockDevices:omitempty"`
	BlockDevicesDiskSpaceTotal *float64                                             `json:"blockDevicesDiskSpaceTotal:omitempty"`
	Children                   []Virtual_Guest_Block_Device_Template_Group          `json:"children:omitempty"`
	ChildrenCount              *uint                                                `json:"childrenCount:omitempty"`
	CreateDate                 *time.Time                                           `json:"createDate:omitempty"`
	Datacenter                 *Location                                            `json:"datacenter:omitempty"`
	DatacenterCount            *uint                                                `json:"datacenterCount:omitempty"`
	Datacenters                []Location                                           `json:"datacenters:omitempty"`
	FlexImageFlag              *bool                                                `json:"flexImageFlag:omitempty"`
	GlobalIdentifier           *string                                              `json:"globalIdentifier:omitempty"`
	Id                         *int                                                 `json:"id:omitempty"`
	ImageType                  *string                                              `json:"imageType:omitempty"`
	ImageTypeKeyName           *string                                              `json:"imageTypeKeyName:omitempty"`
	Name                       *string                                              `json:"name:omitempty"`
	Note                       *string                                              `json:"note:omitempty"`
	Parent                     *Virtual_Guest_Block_Device_Template_Group           `json:"parent:omitempty"`
	ParentId                   *int                                                 `json:"parentId:omitempty"`
	PublicFlag                 *int                                                 `json:"publicFlag:omitempty"`
	SshKeyCount                *uint                                                `json:"sshKeyCount:omitempty"`
	SshKeys                    []Security_Ssh_Key                                   `json:"sshKeys:omitempty"`
	Status                     *Virtual_Guest_Block_Device_Template_Group_Status    `json:"status:omitempty"`
	StatusId                   *int                                                 `json:"statusId:omitempty"`
	StorageRepository          *Virtual_Storage_Repository                          `json:"storageRepository:omitempty"`
	Summary                    *string                                              `json:"summary:omitempty"`
	TagReferenceCount          *uint                                                `json:"tagReferenceCount:omitempty"`
	TagReferences              []Tag_Reference                                      `json:"tagReferences:omitempty"`
	Transaction                *Provisioning_Version1_Transaction                   `json:"transaction:omitempty"`
	TransactionId              *int                                                 `json:"transactionId:omitempty"`
	UserRecordId               *int                                                 `json:"userRecordId:omitempty"`
}

type Virtual_Guest_Block_Device_Template_Group_Accounts struct {
	Entity

	Account    *Account                                   `json:"account:omitempty"`
	AccountId  *int                                       `json:"accountId:omitempty"`
	CreateDate *time.Time                                 `json:"createDate:omitempty"`
	Group      *Virtual_Guest_Block_Device_Template_Group `json:"group:omitempty"`
	GroupId    *int                                       `json:"groupId:omitempty"`
}

type Virtual_Guest_Block_Device_Template_Group_Status struct {
	Entity

	Description *string `json:"description:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Virtual_Guest_Boot_Parameter struct {
	Entity

	CreateDate               *time.Time                         `json:"createDate:omitempty"`
	Guest                    *Virtual_Guest                     `json:"guest:omitempty"`
	GuestBootParameterType   *Virtual_Guest_Boot_Parameter_Type `json:"guestBootParameterType:omitempty"`
	GuestBootParameterTypeId *int                               `json:"guestBootParameterTypeId:omitempty"`
	GuestId                  *int                               `json:"guestId:omitempty"`
	Id                       *int                               `json:"id:omitempty"`
	ModifyDate               *time.Time                         `json:"modifyDate:omitempty"`
}

type Virtual_Guest_Boot_Parameter_Type struct {
	Entity

	BootOption  *string    `json:"bootOption:omitempty"`
	CreateDate  *time.Time `json:"createDate:omitempty"`
	Description *string    `json:"description:omitempty"`
	Id          *int       `json:"id:omitempty"`
	KeyName     *string    `json:"keyName:omitempty"`
	ModifyDate  *time.Time `json:"modifyDate:omitempty"`
	Name        *string    `json:"name:omitempty"`
	Value       *string    `json:"value:omitempty"`
}

type Virtual_Guest_Network_Component struct {
	Entity

	CreateDate                     *time.Time                                  `json:"createDate:omitempty"`
	Guest                          *Virtual_Guest                              `json:"guest:omitempty"`
	GuestId                        *int                                        `json:"guestId:omitempty"`
	HighAvailabilityFirewallFlag   *bool                                       `json:"highAvailabilityFirewallFlag:omitempty"`
	Id                             *int                                        `json:"id:omitempty"`
	IpAddressBindingCount          *uint                                       `json:"ipAddressBindingCount:omitempty"`
	IpAddressBindings              []Virtual_Guest_Network_Component_IpAddress `json:"ipAddressBindings:omitempty"`
	MacAddress                     *string                                     `json:"macAddress:omitempty"`
	MaxSpeed                       *int                                        `json:"maxSpeed:omitempty"`
	ModifyDate                     *time.Time                                  `json:"modifyDate:omitempty"`
	Name                           *string                                     `json:"name:omitempty"`
	NetworkComponentFirewall       *Network_Component_Firewall                 `json:"networkComponentFirewall:omitempty"`
	NetworkId                      *int                                        `json:"networkId:omitempty"`
	NetworkVlan                    *Network_Vlan                               `json:"networkVlan:omitempty"`
	Port                           *int                                        `json:"port:omitempty"`
	PrimaryIpAddress               *string                                     `json:"primaryIpAddress:omitempty"`
	PrimaryIpAddressRecord         *Network_Subnet_IpAddress                   `json:"primaryIpAddressRecord:omitempty"`
	PrimarySubnet                  *Network_Subnet                             `json:"primarySubnet:omitempty"`
	PrimaryVersion6IpAddressRecord *Network_Subnet_IpAddress                   `json:"primaryVersion6IpAddressRecord:omitempty"`
	Router                         *Hardware_Router                            `json:"router:omitempty"`
	Speed                          *int                                        `json:"speed:omitempty"`
	Status                         *string                                     `json:"status:omitempty"`
	SubnetCount                    *uint                                       `json:"subnetCount:omitempty"`
	Subnets                        []Network_Subnet                            `json:"subnets:omitempty"`
	Uuid                           *string                                     `json:"uuid:omitempty"`
}

type Virtual_Guest_Network_Component_IpAddress struct {
	Entity

	IpAddress        *Network_Subnet_IpAddress        `json:"ipAddress:omitempty"`
	IpAddressId      *int                             `json:"ipAddressId:omitempty"`
	NetworkComponent *Virtual_Guest_Network_Component `json:"networkComponent:omitempty"`
	Port             *int                             `json:"port:omitempty"`
	Type             *string                          `json:"type:omitempty"`
}

type Virtual_Guest_Power_State struct {
	Entity

	Description *string `json:"description:omitempty"`
	KeyName     *string `json:"keyName:omitempty"`
	Name        *string `json:"name:omitempty"`
}

type Virtual_Guest_Status struct {
	Entity

	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}

type Virtual_Guest_SupplementalCreateObjectOptions struct {
	Entity

	ImmediateApprovalOnlyFlag *bool   `json:"immediateApprovalOnlyFlag:omitempty"`
	PostInstallScriptUri      *string `json:"postInstallScriptUri:omitempty"`
}

type Virtual_Host struct {
	Entity

	Account                  *Account                `json:"account:omitempty"`
	AccountId                *int                    `json:"accountId:omitempty"`
	BilledPerGuestFlag       *bool                   `json:"billedPerGuestFlag:omitempty"`
	BilledPerMemoryUsageFlag *bool                   `json:"billedPerMemoryUsageFlag:omitempty"`
	CreateDate               *time.Time              `json:"createDate:omitempty"`
	Description              *string                 `json:"description:omitempty"`
	EnabledFlag              *int                    `json:"enabledFlag:omitempty"`
	GuestCount               *uint                   `json:"guestCount:omitempty"`
	Guests                   []Virtual_Guest         `json:"guests:omitempty"`
	Hardware                 *Hardware_Server        `json:"hardware:omitempty"`
	HardwareId               *int                    `json:"hardwareId:omitempty"`
	Id                       *int                    `json:"id:omitempty"`
	MetricTrackingObject     *Metric_Tracking_Object `json:"metricTrackingObject:omitempty"`
	ModifyDate               *time.Time              `json:"modifyDate:omitempty"`
	Name                     *string                 `json:"name:omitempty"`
	PhysicalMemoryCapacity   *int                    `json:"physicalMemoryCapacity:omitempty"`
	Uuid                     *string                 `json:"uuid:omitempty"`
}

type Virtual_Storage_Repository struct {
	Entity

	Account                *Account                                           `json:"account:omitempty"`
	BillingItem            *Billing_Item                                      `json:"billingItem:omitempty"`
	Capacity               *float64                                           `json:"capacity:omitempty"`
	Datacenter             *Location                                          `json:"datacenter:omitempty"`
	Description            *string                                            `json:"description:omitempty"`
	DiskImageCount         *uint                                              `json:"diskImageCount:omitempty"`
	DiskImages             []Virtual_Disk_Image                               `json:"diskImages:omitempty"`
	GuestCount             *uint                                              `json:"guestCount:omitempty"`
	Guests                 []Virtual_Guest                                    `json:"guests:omitempty"`
	Id                     *int                                               `json:"id:omitempty"`
	MetricTrackingObject   *Metric_Tracking_Object_Virtual_Storage_Repository `json:"metricTrackingObject:omitempty"`
	Name                   *string                                            `json:"name:omitempty"`
	PublicFlag             *int                                               `json:"publicFlag:omitempty"`
	PublicImageBillingItem *Billing_Item                                      `json:"publicImageBillingItem:omitempty"`
	Type                   *Virtual_Storage_Repository_Type                   `json:"type:omitempty"`
	TypeId                 *int                                               `json:"typeId:omitempty"`
}

type Virtual_Storage_Repository_Type struct {
	Entity

	Description            *string                      `json:"description:omitempty"`
	Name                   *string                      `json:"name:omitempty"`
	StorageRepositories    []Virtual_Storage_Repository `json:"storageRepositories:omitempty"`
	StorageRepositoryCount *uint                        `json:"storageRepositoryCount:omitempty"`
}


