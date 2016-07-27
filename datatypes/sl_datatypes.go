package datatypes

import (
	"time"
)

type McAfee_Epolicy_Orchestrator_Version36_Agent_Details struct {
	Entity

	AgentVersion  *string
	CurrentPolicy *McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details
	LastUpdate    *string
}

type McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details struct {
	Entity

	CurrentPolicy *McAfee_Epolicy_Orchestrator_Version36_Agent_Parent_Details
	Name          *string
}

type McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event struct {
	Entity

	EventLocalDateTime *time.Time
	Filename           *string
	VirusActionTaken   *McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_Filter_Description
	VirusName          *string
	VirusType          *string
}

type McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_AccessProtection struct {
	Entity

	EventLocalDateTime *time.Time
	Filename           *string
	ProcessName        *string
	RuleName           *string
	Source             *string
}

type McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_Filter_Description struct {
	Entity

	Name *string
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_BlockedApplicationEvent struct {
	Entity

	ApplicationDescription *string
	IncidentTime           *time.Time
	ProcessName            *string
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_Event_Signature struct {
	Entity

	SignatureName *string
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_IPSEvent struct {
	Entity

	IncidentTime    *time.Time
	ProcessName     *string
	ReactionText    *string
	RemoteIpAddress *string
	SeverityText    *string
	Signature       *McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_Event_Signature
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_BlockedApplicationEvent struct {
	Entity

	ApplicationDescription *string
	IncidentTime           *time.Time
	ProcessName            *string
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_Event_Signature struct {
	Entity

	SignatureName *string
}

type McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent struct {
	Entity

	IncidentTime    *time.Time
	ProcessName     *string
	ReactionText    *string
	RemoteIpAddress *string
	SeverityText    *string
	Signature       *McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_Event_Signature
}

type McAfee_Epolicy_Orchestrator_Version36_Policy_Object struct {
	Entity

	Name *string
}

type McAfee_Epolicy_Orchestrator_Version36_Product_Properties struct {
	Entity

	DatVersion *string
}

type McAfee_Epolicy_Orchestrator_Version45_Agent_Details struct {
	Entity

	AgentVersion *string
	LastUpdate   *time.Time
}

type McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details struct {
	Entity

	AgentDetails *McAfee_Epolicy_Orchestrator_Version45_Agent_Details
	Name         *string
	Policies     []McAfee_Epolicy_Orchestrator_Version45_Agent_Parent_Details
	PolicyCount  *uint
}

type McAfee_Epolicy_Orchestrator_Version45_Event struct {
	Entity

	AgentDetails        *McAfee_Epolicy_Orchestrator_Version45_Agent_Details
	DetectedUtc         *time.Time
	SourceIpv4          *string
	SourceProcessName   *string
	TargetFilename      *string
	ThreatActionTaken   *string
	ThreatName          *string
	ThreatSeverityLabel *string
	ThreatType          *string
	VirusActionTaken    *McAfee_Epolicy_Orchestrator_Version45_Event_Filter_Description
}

type McAfee_Epolicy_Orchestrator_Version45_Event_Filter_Description struct {
	Entity

	Name *string
}

type McAfee_Epolicy_Orchestrator_Version45_Event_Version7 struct {
	McAfee_Epolicy_Orchestrator_Version45_Event

	Signature *McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version7
}

type McAfee_Epolicy_Orchestrator_Version45_Event_Version8 struct {
	McAfee_Epolicy_Orchestrator_Version45_Event

	Signature *McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version8
}

type McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version7 struct {
	Entity

	SignatureName *string
}

type McAfee_Epolicy_Orchestrator_Version45_Hips_Event_Signature_Version8 struct {
	Entity

	SignatureName *string
}

type McAfee_Epolicy_Orchestrator_Version45_Policy_Object struct {
	Entity

	Name *string
}

type McAfee_Epolicy_Orchestrator_Version45_Product_Properties struct {
	Entity

	DatVersion *string
}

type Abuse_Lockdown_Resource struct {
	Entity

	Account     *Account
	InvoiceItem *Billing_Invoice_Item
}

type Account struct {
	Entity

	AbuseEmail                                             *string
	AbuseEmailCount                                        *uint
	AbuseEmails                                            []Account_AbuseEmail
	AccountContactCount                                    *uint
	AccountContacts                                        []Account_Contact
	AccountLicenseCount                                    *uint
	AccountLicenses                                        []Software_AccountLicense
	AccountLinkCount                                       *uint
	AccountLinks                                           []Account_Link
	AccountManagedResourcesFlag                            *bool
	AccountStatus                                          *Account_Status
	AccountStatusId                                        *int
	ActiveAccountDiscountBillingItem                       *Billing_Item
	ActiveAccountLicenseCount                              *uint
	ActiveAccountLicenses                                  []Software_AccountLicense
	ActiveAddressCount                                     *uint
	ActiveAddresses                                        []Account_Address
	ActiveBillingAgreementCount                            *uint
	ActiveBillingAgreements                                []Account_Agreement
	ActiveCatalystEnrollment                               *Catalyst_Enrollment
	ActiveColocationContainerCount                         *uint
	ActiveColocationContainers                             []Billing_Item
	ActiveFlexibleCreditEnrollment                         *FlexibleCredit_Enrollment
	ActiveNotificationSubscriberCount                      *uint
	ActiveNotificationSubscribers                          []Notification_Subscriber
	ActiveQuoteCount                                       *uint
	ActiveQuotes                                           []Billing_Order_Quote
	ActiveVirtualLicenseCount                              *uint
	ActiveVirtualLicenses                                  []Software_VirtualLicense
	AdcLoadBalancerCount                                   *uint
	AdcLoadBalancers                                       []Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress
	Address1                                               *string
	Address2                                               *string
	AddressCount                                           *uint
	Addresses                                              []Account_Address
	AffiliateId                                            *string
	AllBillingItems                                        []Billing_Item
	AllCommissionBillingItemCount                          *uint
	AllCommissionBillingItems                              []Billing_Item
	AllRecurringTopLevelBillingItemCount                   *uint
	AllRecurringTopLevelBillingItems                       []Billing_Item
	AllRecurringTopLevelBillingItemsUnfiltered             []Billing_Item
	AllRecurringTopLevelBillingItemsUnfilteredCount        *uint
	AllSubnetBillingItemCount                              *uint
	AllSubnetBillingItems                                  []Billing_Item
	AllTopLevelBillingItemCount                            *uint
	AllTopLevelBillingItems                                []Billing_Item
	AllTopLevelBillingItemsUnfiltered                      []Billing_Item
	AllTopLevelBillingItemsUnfilteredCount                 *uint
	AllowedPptpVpnQuantity                                 *int
	AllowsBluemixAccountLinkingFlag                        *bool
	AlternatePhone                                         *string
	ApplicationDeliveryControllerCount                     *uint
	ApplicationDeliveryControllers                         []Network_Application_Delivery_Controller
	AttributeCount                                         *uint
	Attributes                                             []Account_Attribute
	AvailablePublicNetworkVlanCount                        *uint
	AvailablePublicNetworkVlans                            []Network_Vlan
	Balance                                                *float64
	BandwidthAllotmentCount                                *uint
	BandwidthAllotments                                    []Network_Bandwidth_Version1_Allotment
	BandwidthAllotmentsOverAllocation                      []Network_Bandwidth_Version1_Allotment
	BandwidthAllotmentsOverAllocationCount                 *uint
	BandwidthAllotmentsProjectedOverAllocation             []Network_Bandwidth_Version1_Allotment
	BandwidthAllotmentsProjectedOverAllocationCount        *uint
	BareMetalInstanceCount                                 *uint
	BareMetalInstances                                     []Hardware
	BillingAgreementCount                                  *uint
	BillingAgreements                                      []Account_Agreement
	BillingInfo                                            *Billing_Info
	BlockDeviceTemplateGroupCount                          *uint
	BlockDeviceTemplateGroups                              []Virtual_Guest_Block_Device_Template_Group
	BlueIdAuthenticationRequiredFlag                       *bool
	BluemixLinkedFlag                                      *bool
	Brand                                                  *Brand
	BrandAccountFlag                                       *bool
	BrandId                                                *int
	BrandKeyName                                           *string
	CanOrderAdditionalVlansFlag                            *bool
	CartCount                                              *uint
	Carts                                                  []Billing_Order_Quote
	CatalystEnrollmentCount                                *uint
	CatalystEnrollments                                    []Catalyst_Enrollment
	CdnAccountCount                                        *uint
	CdnAccounts                                            []Network_ContentDelivery_Account
	City                                                   *string
	ClaimedTaxExemptTxFlag                                 *bool
	ClosedTicketCount                                      *uint
	ClosedTickets                                          []Ticket
	CompanyName                                            *string
	Country                                                *string
	CreateDate                                             *time.Time
	DatacentersWithSubnetAllocationCount                   *uint
	DatacentersWithSubnetAllocations                       []Location
	DeviceFingerprintId                                    *string
	DisablePaymentProcessingFlag                           *bool
	DisplaySupportRepresentativeAssignmentCount            *uint
	DisplaySupportRepresentativeAssignments                []Account_Attachment_Employee
	DomainCount                                            *uint
	DomainRegistrationCount                                *uint
	DomainRegistrations                                    []Dns_Domain_Registration
	Domains                                                []Dns_Domain
	DomainsWithoutSecondaryDnsRecordCount                  *uint
	DomainsWithoutSecondaryDnsRecords                      []Dns_Domain
	Email                                                  *string
	EvaultCapacityGB                                       *uint
	EvaultMasterUserCount                                  *uint
	EvaultMasterUsers                                      []Account_Password
	EvaultNetworkStorage                                   []Network_Storage
	EvaultNetworkStorageCount                              *uint
	ExpiredSecurityCertificateCount                        *uint
	ExpiredSecurityCertificates                            []Security_Certificate
	FacilityLogCount                                       *uint
	FacilityLogs                                           []User_Access_Facility_Log
	FaxPhone                                               *string
	FirstName                                              *string
	FlexibleCreditEnrollmentCount                          *uint
	FlexibleCreditEnrollments                              []FlexibleCredit_Enrollment
	GlobalIpRecordCount                                    *uint
	GlobalIpRecords                                        []Network_Subnet_IpAddress_Global
	GlobalIpv4RecordCount                                  *uint
	GlobalIpv4Records                                      []Network_Subnet_IpAddress_Global
	GlobalIpv6RecordCount                                  *uint
	GlobalIpv6Records                                      []Network_Subnet_IpAddress_Global
	GlobalLoadBalancerAccountCount                         *uint
	GlobalLoadBalancerAccounts                             []Network_LoadBalancer_Global_Account
	Hardware                                               []Hardware
	HardwareCount                                          *uint
	HardwareOverBandwidthAllocation                        []Hardware
	HardwareOverBandwidthAllocationCount                   *uint
	HardwareProjectedOverBandwidthAllocation               []Hardware
	HardwareProjectedOverBandwidthAllocationCount          *uint
	HardwareWithCpanel                                     []Hardware
	HardwareWithCpanelCount                                *uint
	HardwareWithHelm                                       []Hardware
	HardwareWithHelmCount                                  *uint
	HardwareWithMcafee                                     []Hardware
	HardwareWithMcafeeAntivirusRedhat                      []Hardware
	HardwareWithMcafeeAntivirusRedhatCount                 *uint
	HardwareWithMcafeeAntivirusWindowCount                 *uint
	HardwareWithMcafeeAntivirusWindows                     []Hardware
	HardwareWithMcafeeCount                                *uint
	HardwareWithMcafeeIntrusionDetectionSystem             []Hardware
	HardwareWithMcafeeIntrusionDetectionSystemCount        *uint
	HardwareWithPlesk                                      []Hardware
	HardwareWithPleskCount                                 *uint
	HardwareWithQuantastor                                 []Hardware
	HardwareWithQuantastorCount                            *uint
	HardwareWithUrchin                                     []Hardware
	HardwareWithUrchinCount                                *uint
	HardwareWithWindowCount                                *uint
	HardwareWithWindows                                    []Hardware
	HasEvaultBareMetalRestorePluginFlag                    *bool
	HasIderaBareMetalRestorePluginFlag                     *bool
	HasPendingOrder                                        *uint
	HasR1softBareMetalRestorePluginFlag                    *bool
	HourlyBareMetalInstanceCount                           *uint
	HourlyBareMetalInstances                               []Hardware
	HourlyServiceBillingItemCount                          *uint
	HourlyServiceBillingItems                              []Billing_Item
	HourlyVirtualGuestCount                                *uint
	HourlyVirtualGuests                                    []Virtual_Guest
	HubNetworkStorage                                      []Network_Storage
	HubNetworkStorageCount                                 *uint
	Id                                                     *int
	InternalNoteCount                                      *uint
	InternalNotes                                          []Account_Note
	InvoiceCount                                           *uint
	Invoices                                               []Billing_Invoice
	IpAddressCount                                         *uint
	IpAddresses                                            []Network_Subnet_IpAddress
	IsReseller                                             *int
	IscsiNetworkStorage                                    []Network_Storage
	IscsiNetworkStorageCount                               *uint
	LastCanceledBillingItem                                *Billing_Item
	LastCancelledServerBillingItem                         *Billing_Item
	LastFiveClosedAbuseTicketCount                         *uint
	LastFiveClosedAbuseTickets                             []Ticket
	LastFiveClosedAccountingTicketCount                    *uint
	LastFiveClosedAccountingTickets                        []Ticket
	LastFiveClosedOtherTicketCount                         *uint
	LastFiveClosedOtherTickets                             []Ticket
	LastFiveClosedSalesTicketCount                         *uint
	LastFiveClosedSalesTickets                             []Ticket
	LastFiveClosedSupportTicketCount                       *uint
	LastFiveClosedSupportTickets                           []Ticket
	LastFiveClosedTicketCount                              *uint
	LastFiveClosedTickets                                  []Ticket
	LastName                                               *string
	LateFeeProtectionFlag                                  *bool
	LatestBillDate                                         *time.Time
	LatestRecurringInvoice                                 *Billing_Invoice
	LatestRecurringPendingInvoice                          *Billing_Invoice
	LegacyBandwidthAllotmentCount                          *uint
	LegacyBandwidthAllotments                              []Network_Bandwidth_Version1_Allotment
	LegacyIscsiCapacityGB                                  *uint
	LoadBalancerCount                                      *uint
	LoadBalancers                                          []Network_LoadBalancer_VirtualIpAddress
	LockboxCapacityGB                                      *uint
	LockboxNetworkStorage                                  []Network_Storage
	LockboxNetworkStorageCount                             *uint
	ManualPaymentsUnderReview                              []Billing_Payment_Card_ManualPayment
	ManualPaymentsUnderReviewCount                         *uint
	MasterUser                                             *User_Customer
	MediaDataTransferRequestCount                          *uint
	MediaDataTransferRequests                              []Account_Media_Data_Transfer_Request
	MessageQueueAccountCount                               *uint
	MessageQueueAccounts                                   []Network_Message_Queue
	ModifyDate                                             *time.Time
	MonthlyBareMetalInstanceCount                          *uint
	MonthlyBareMetalInstances                              []Hardware
	MonthlyVirtualGuestCount                               *uint
	MonthlyVirtualGuests                                   []Virtual_Guest
	NasNetworkStorage                                      []Network_Storage
	NasNetworkStorageCount                                 *uint
	NetworkCreationFlag                                    *bool
	NetworkGatewayCount                                    *uint
	NetworkGateways                                        []Network_Gateway
	NetworkHardware                                        []Hardware
	NetworkHardwareCount                                   *uint
	NetworkMessageDeliveryAccountCount                     *uint
	NetworkMessageDeliveryAccounts                         []Network_Message_Delivery
	NetworkMonitorDownHardware                             []Hardware
	NetworkMonitorDownHardwareCount                        *uint
	NetworkMonitorDownVirtualGuestCount                    *uint
	NetworkMonitorDownVirtualGuests                        []Virtual_Guest
	NetworkMonitorRecoveringHardware                       []Hardware
	NetworkMonitorRecoveringHardwareCount                  *uint
	NetworkMonitorRecoveringVirtualGuestCount              *uint
	NetworkMonitorRecoveringVirtualGuests                  []Virtual_Guest
	NetworkMonitorUpHardware                               []Hardware
	NetworkMonitorUpHardwareCount                          *uint
	NetworkMonitorUpVirtualGuestCount                      *uint
	NetworkMonitorUpVirtualGuests                          []Virtual_Guest
	NetworkStorage                                         []Network_Storage
	NetworkStorageCount                                    *uint
	NetworkStorageGroupCount                               *uint
	NetworkStorageGroups                                   []Network_Storage_Group
	NetworkTunnelContextCount                              *uint
	NetworkTunnelContexts                                  []Network_Tunnel_Module_Context
	NetworkVlanCount                                       *uint
	NetworkVlanSpan                                        *Account_Network_Vlan_Span
	NetworkVlans                                           []Network_Vlan
	NextBillingPublicAllotmentHardwareBandwidthDetailCount *uint
	NextBillingPublicAllotmentHardwareBandwidthDetails     []Network_Bandwidth_Version1_Allotment
	NextInvoiceIncubatorExemptTotal                        *float64
	NextInvoiceTopLevelBillingItemCount                    *uint
	NextInvoiceTopLevelBillingItems                        []Billing_Item
	NextInvoiceTotalAmount                                 *float64
	NextInvoiceTotalOneTimeAmount                          *float64
	NextInvoiceTotalOneTimeTaxAmount                       *float64
	NextInvoiceTotalRecurringAmount                        *float64
	NextInvoiceTotalRecurringAmountBeforeAccountDiscount   *float64
	NextInvoiceTotalRecurringTaxAmount                     *float64
	NextInvoiceTotalTaxableRecurringAmount                 *float64
	NotificationSubscriberCount                            *uint
	NotificationSubscribers                                []Notification_Subscriber
	OfficePhone                                            *string
	OpenAbuseTicketCount                                   *uint
	OpenAbuseTickets                                       []Ticket
	OpenAccountingTicketCount                              *uint
	OpenAccountingTickets                                  []Ticket
	OpenBillingTicketCount                                 *uint
	OpenBillingTickets                                     []Ticket
	OpenCancellationRequestCount                           *uint
	OpenCancellationRequests                               []Billing_Item_Cancellation_Request
	OpenOtherTicketCount                                   *uint
	OpenOtherTickets                                       []Ticket
	OpenRecurringInvoiceCount                              *uint
	OpenRecurringInvoices                                  []Billing_Invoice
	OpenSalesTicketCount                                   *uint
	OpenSalesTickets                                       []Ticket
	OpenStackAccountLinkCount                              *uint
	OpenStackAccountLinks                                  []Account_Link
	OpenStackObjectStorage                                 []Network_Storage
	OpenStackObjectStorageCount                            *uint
	OpenSupportTicketCount                                 *uint
	OpenSupportTickets                                     []Ticket
	OpenTicketCount                                        *uint
	OpenTickets                                            []Ticket
	OpenTicketsWaitingOnCustomer                           []Ticket
	OpenTicketsWaitingOnCustomerCount                      *uint
	OrderCount                                             *uint
	Orders                                                 []Billing_Order
	OrphanBillingItemCount                                 *uint
	OrphanBillingItems                                     []Billing_Item
	OwnedBrandCount                                        *uint
	OwnedBrands                                            []Brand
	OwnedHardwareGenericComponentModelCount                *uint
	OwnedHardwareGenericComponentModels                    []Hardware_Component_Model_Generic
	PaymentProcessorCount                                  *uint
	PaymentProcessors                                      []Billing_Payment_Processor
	PendingEventCount                                      *uint
	PendingEvents                                          []Notification_Occurrence_Event
	PendingInvoice                                         *Billing_Invoice
	PendingInvoiceTopLevelItemCount                        *uint
	PendingInvoiceTopLevelItems                            []Billing_Invoice_Item
	PendingInvoiceTotalAmount                              *float64
	PendingInvoiceTotalOneTimeAmount                       *float64
	PendingInvoiceTotalOneTimeTaxAmount                    *float64
	PendingInvoiceTotalRecurringAmount                     *float64
	PendingInvoiceTotalRecurringTaxAmount                  *float64
	PermissionGroupCount                                   *uint
	PermissionGroups                                       []User_Permission_Group
	PermissionRoleCount                                    *uint
	PermissionRoles                                        []User_Permission_Role
	PortableStorageVolumeCount                             *uint
	PortableStorageVolumes                                 []Virtual_Disk_Image
	PostProvisioningHookCount                              *uint
	PostProvisioningHooks                                  []Provisioning_Hook
	PostalCode                                             *string
	PptpVpnUserCount                                       *uint
	PptpVpnUsers                                           []User_Customer
	PreviousRecurringRevenue                               *float64
	PriceRestrictionCount                                  *uint
	PriceRestrictions                                      []Product_Item_Price_Account_Restriction
	PriorityOneTicketCount                                 *uint
	PriorityOneTickets                                     []Ticket
	PrivateAllotmentHardwareBandwidthDetailCount           *uint
	PrivateAllotmentHardwareBandwidthDetails               []Network_Bandwidth_Version1_Allotment
	PrivateBlockDeviceTemplateGroupCount                   *uint
	PrivateBlockDeviceTemplateGroups                       []Virtual_Guest_Block_Device_Template_Group
	PrivateIpAddressCount                                  *uint
	PrivateIpAddresses                                     []Network_Subnet_IpAddress
	PrivateNetworkVlanCount                                *uint
	PrivateNetworkVlans                                    []Network_Vlan
	PrivateSubnetCount                                     *uint
	PrivateSubnets                                         []Network_Subnet
	PublicAllotmentHardwareBandwidthDetailCount            *uint
	PublicAllotmentHardwareBandwidthDetails                []Network_Bandwidth_Version1_Allotment
	PublicIpAddressCount                                   *uint
	PublicIpAddresses                                      []Network_Subnet_IpAddress
	PublicNetworkVlanCount                                 *uint
	PublicNetworkVlans                                     []Network_Vlan
	PublicSubnetCount                                      *uint
	PublicSubnets                                          []Network_Subnet
	QuoteCount                                             *uint
	Quotes                                                 []Billing_Order_Quote
	RecentEventCount                                       *uint
	RecentEvents                                           []Notification_Occurrence_Event
	ReferralPartner                                        *Account
	ReferredAccountCount                                   *uint
	ReferredAccounts                                       []Account
	RegulatedWorkloadCount                                 *uint
	RegulatedWorkloads                                     []Legal_RegulatedWorkload
	RemoteManagementCommandRequestCount                    *uint
	RemoteManagementCommandRequests                        []Hardware_Component_RemoteManagement_Command_Request
	ReplicationEventCount                                  *uint
	ReplicationEvents                                      []Network_Storage_Event
	ResourceGroupCount                                     *uint
	ResourceGroups                                         []Resource_Group
	RouterCount                                            *uint
	Routers                                                []Hardware
	RwhoisData                                             *Network_Subnet_Rwhois_Data
	SalesforceAccountLink                                  *Account_Link
	SamlAuthentication                                     *Account_Authentication_Saml
	ScaleGroupCount                                        *uint
	ScaleGroups                                            []Scale_Group
	SecondaryDomainCount                                   *uint
	SecondaryDomains                                       []Dns_Secondary
	SecurityCertificateCount                               *uint
	SecurityCertificates                                   []Security_Certificate
	SecurityScanRequestCount                               *uint
	SecurityScanRequests                                   []Network_Security_Scanner_Request
	ServiceBillingItemCount                                *uint
	ServiceBillingItems                                    []Billing_Item
	ShipmentCount                                          *uint
	Shipments                                              []Account_Shipment
	SshKeyCount                                            *uint
	SshKeys                                                []Security_Ssh_Key
	SslVpnUserCount                                        *uint
	SslVpnUsers                                            []User_Customer
	StandardPoolVirtualGuestCount                          *uint
	StandardPoolVirtualGuests                              []Virtual_Guest
	State                                                  *string
	StatusDate                                             *time.Time
	SubnetCount                                            *uint
	SubnetRegistrationCount                                *uint
	SubnetRegistrationDetailCount                          *uint
	SubnetRegistrationDetails                              []Account_Regional_Registry_Detail
	SubnetRegistrations                                    []Network_Subnet_Registration
	Subnets                                                []Network_Subnet
	SupportRepresentativeCount                             *uint
	SupportRepresentatives                                 []User_Employee
	SupportSubscriptionCount                               *uint
	SupportSubscriptions                                   []Billing_Item
	SupportTier                                            *string
	SuppressInvoicesFlag                                   *bool
	TagCount                                               *uint
	Tags                                                   []Tag
	TicketCount                                            *uint
	Tickets                                                []Ticket
	TicketsClosedInTheLastThreeDays                        []Ticket
	TicketsClosedInTheLastThreeDaysCount                   *uint
	TicketsClosedToday                                     []Ticket
	TicketsClosedTodayCount                                *uint
	TranscodeAccountCount                                  *uint
	TranscodeAccounts                                      []Network_Media_Transcode_Account
	UpgradeRequestCount                                    *uint
	UpgradeRequests                                        []Product_Upgrade_Request
	UserCount                                              *uint
	Users                                                  []User_Customer
	ValidSecurityCertificateCount                          *uint
	ValidSecurityCertificates                              []Security_Certificate
	VdrUpdatesInProgressFlag                               *bool
	VirtualDedicatedRackCount                              *uint
	VirtualDedicatedRacks                                  []Network_Bandwidth_Version1_Allotment
	VirtualDiskImageCount                                  *uint
	VirtualDiskImages                                      []Virtual_Disk_Image
	VirtualGuestCount                                      *uint
	VirtualGuests                                          []Virtual_Guest
	VirtualGuestsOverBandwidthAllocation                   []Virtual_Guest
	VirtualGuestsOverBandwidthAllocationCount              *uint
	VirtualGuestsProjectedOverBandwidthAllocation          []Virtual_Guest
	VirtualGuestsProjectedOverBandwidthAllocationCount     *uint
	VirtualGuestsWithCpanel                                []Virtual_Guest
	VirtualGuestsWithCpanelCount                           *uint
	VirtualGuestsWithMcafee                                []Virtual_Guest
	VirtualGuestsWithMcafeeAntivirusRedhat                 []Virtual_Guest
	VirtualGuestsWithMcafeeAntivirusRedhatCount            *uint
	VirtualGuestsWithMcafeeAntivirusWindowCount            *uint
	VirtualGuestsWithMcafeeAntivirusWindows                []Virtual_Guest
	VirtualGuestsWithMcafeeCount                           *uint
	VirtualGuestsWithMcafeeIntrusionDetectionSystem        []Virtual_Guest
	VirtualGuestsWithMcafeeIntrusionDetectionSystemCount   *uint
	VirtualGuestsWithPlesk                                 []Virtual_Guest
	VirtualGuestsWithPleskCount                            *uint
	VirtualGuestsWithQuantastor                            []Virtual_Guest
	VirtualGuestsWithQuantastorCount                       *uint
	VirtualGuestsWithUrchin                                []Virtual_Guest
	VirtualGuestsWithUrchinCount                           *uint
	VirtualPrivateRack                                     *Network_Bandwidth_Version1_Allotment
	VirtualStorageArchiveRepositories                      []Virtual_Storage_Repository
	VirtualStorageArchiveRepositoryCount                   *uint
	VirtualStoragePublicRepositories                       []Virtual_Storage_Repository
	VirtualStoragePublicRepositoryCount                    *uint
}

type Account_AbuseEmail struct {
	Entity

	Account *Account
	Email   *string
}

type Account_Address struct {
	Entity

	Account        *Account
	AccountId      *int
	Address1       *string
	Address2       *string
	City           *string
	ContactName    *string
	Country        *string
	CreateUser     *User_Customer
	Description    *string
	Id             *int
	IsActive       *int
	Location       *Location
	LocationId     *int
	ModifyEmployee *User_Employee
	ModifyUser     *User_Customer
	PostalCode     *string
	State          *string
	Type           *Account_Address_Type
}

type Account_Address_Type struct {
	Entity

	CreateDate *time.Time
	Id         *int
	KeyName    *string
	Name       *string
}

type Account_Affiliation struct {
	Entity

	Account     *Account
	AccountId   *int
	AffiliateId *string
	CreateDate  *time.Time
	Id          *int
	ModifyDate  *time.Time
}

type Account_Agreement struct {
	Entity

	Account                           *Account
	AgreementType                     *Account_Agreement_Type
	AgreementTypeId                   *int
	AttachedBillingAgreementFileCount *uint
	AttachedBillingAgreementFiles     []Account_MasterServiceAgreement
	AutoRenew                         *int
	BillingItemCount                  *uint
	BillingItems                      []Billing_Item
	CancellationFee                   *int
	CreateDate                        *time.Time
	DurationMonths                    *int
	EndDate                           *time.Time
	Id                                *int
	StartDate                         *time.Time
	Status                            *Account_Agreement_Status
	StatusId                          *int
	Title                             *string
	TopLevelBillingItemCount          *uint
	TopLevelBillingItems              []Billing_Item
}

type Account_Agreement_Status struct {
	Entity

	Name *string
}

type Account_Agreement_Type struct {
	Entity

	Name *string
}

type Account_Attachment_Employee struct {
	Entity

	Account      *Account
	Employee     *User_Employee
	EmployeeRole *Account_Attachment_Employee_Role
	RoleId       *int
}

type Account_Attachment_Employee_Role struct {
	Entity

	Keyname *string
	Name    *string
}

type Account_Attribute struct {
	Entity

	Account                *Account
	AccountAttributeType   *Account_Attribute_Type
	AccountAttributeTypeId *int
	AccountId              *int
	Id                     *int
	Value                  *string
}

type Account_Attribute_Type struct {
	Entity

	Description *string
	Id          *int
	KeyName     *string
	Name        *string
}

type Account_Authentication_Attribute struct {
	Entity

	Account              *Account
	AccountId            *int
	AuthenticationRecord *Account_Authentication_Saml
	Id                   *int
	Type                 *Account_Authentication_Attribute_Type
	TypeId               *int
	Value                *string
}

type Account_Authentication_Attribute_Type struct {
	Entity

	Description  *string
	Id           *int
	KeyName      *string
	Name         *string
	ValueExample *string
}

type Account_Authentication_OpenIdConnect_Option struct {
	Entity

	Key   *string
	Value *string
}

type Account_Authentication_OpenIdConnect_RegistrationInformation struct {
	Entity

	ExistingBlueIdFlag       *bool
	FederatedEmailDomainFlag *bool
	User                     *User_Customer
}

type Account_Authentication_Saml struct {
	Entity

	Account                             *Account
	AccountId                           *string
	AttributeCount                      *uint
	Attributes                          []Account_Authentication_Attribute
	Certificate                         *string
	CertificateFingerprint              *string
	EntityId                            *string
	Id                                  *int
	ServiceProviderCertificate          *string
	ServiceProviderEntityId             *string
	ServiceProviderPublicKey            *string
	ServiceProviderSingleLogoutEncoding *string
	ServiceProviderSingleLogoutUrl      *string
	ServiceProviderSingleSignOnEncoding *string
	ServiceProviderSingleSignOnUrl      *string
	SingleLogoutEncoding                *string
	SingleLogoutUrl                     *string
	SingleSignOnEncoding                *string
	SingleSignOnUrl                     *string
}

type Account_Classification_Group_Type struct {
	Entity

	KeyName *string
}

type Account_Contact struct {
	Entity

	Account        *Account
	AccountId      *int
	Address1       *string
	Address2       *string
	AlternatePhone *string
	City           *string
	CompanyName    *string
	Country        *string
	CreateDate     *time.Time
	Email          *string
	FaxPhone       *string
	FirstName      *string
	Id             *int
	JobTitle       *string
	LastName       *string
	ModifyDate     *time.Time
	OfficePhone    *string
	PostalCode     *string
	ProfileName    *string
	State          *string
	Type           *Account_Contact_Type
	TypeId         *int
	Url            *string
}

type Account_Contact_Type struct {
	Entity

	CreateDate  *time.Time
	Description *string
	Id          *int
	KeyName     *string
	ModifyDate  *time.Time
	Name        *string
}

type Account_Historical_Report struct {
	Entity
}

type Account_Link struct {
	Entity

	Account                          *Account
	AccountId                        *int
	CreateDate                       *time.Time
	DestinationAccountAlphanumericId *string
	DestinationAccountId             *int
	Id                               *int
	ServiceProvider                  *Service_Provider
	ServiceProviderId                *int
}

type Account_Link_Bluemix struct {
	Account_Link
}

type Account_Link_OpenStack struct {
	Account_Link

	DomainId *string
}

type Account_Link_OpenStack_DomainCreationDetails struct {
	Entity

	DomainId *string
	UserId   *string
	UserName *string
}

type Account_Link_OpenStack_LinkRequest struct {
	Entity

	DesiredPassword    *string
	DesiredProjectName *string
	DesiredUsername    *string
}

type Account_Link_OpenStack_ProjectCreationDetails struct {
	Entity

	DomainId    *string
	ProjectId   *string
	ProjectName *string
	UserId      *string
	UserName    *string
}

type Account_Link_OpenStack_ProjectDetails struct {
	Entity

	ProjectId   *string
	ProjectName *string
}

type Account_Link_ThePlanet struct {
	Account_Link
}

type Account_Link_Vendor struct {
	Entity

	KeyName *string
	Name    *string
}

type Account_Lockdown_Request struct {
	Entity

	AccountId  *int
	Action     *string
	CreateDate *time.Time
	Id         *int
	ModifyDate *time.Time
	Status     *string
}

type Account_MasterServiceAgreement struct {
	Entity

	Account   *Account
	AccountId *int
	Guid      *string
	Id        *int
	Name      *string
}

type Account_Media struct {
	Entity

	Account        *Account
	CreateUser     *User_Customer
	Datacenter     *Location
	Description    *string
	Id             *int
	ModifyEmployee *User_Employee
	ModifyUser     *User_Customer
	Request        *Account_Media_Data_Transfer_Request
	RequestId      *int
	SerialNumber   *string
	Type           *Account_Media_Type
	TypeId         *int
	Volume         *Network_Storage
}

type Account_Media_Data_Transfer_Request struct {
	Entity

	Account           *Account
	AccountId         *int
	ActiveTicketCount *uint
	ActiveTickets     []Ticket
	BillingItem       *Billing_Item
	CreateUser        *User_Customer
	CreateUserId      *int
	EndDate           *time.Time
	Id                *int
	Media             *Account_Media
	ModifyEmployee    *User_Employee
	ModifyUser        *User_Customer
	ModifyUserId      *int
	ShipmentCount     *uint
	Shipments         []Account_Shipment
	StartDate         *time.Time
	Status            *Account_Media_Data_Transfer_Request_Status
	StatusId          *int
	TicketCount       *uint
	Tickets           []Ticket
}

type Account_Media_Data_Transfer_Request_Status struct {
	Entity

	Description *string
	Id          *int
	KeyName     *string
	Name        *string
}

type Account_Media_Type struct {
	Entity

	Description *string
	Id          *int
	KeyName     *string
	Name        *string
}

type Account_Network_Vlan_Span struct {
	Entity

	Account          *Account
	EnabledFlag      *bool
	Id               *int
	LastAppliedDate  *time.Time
	LastVerifiedDate *time.Time
	ModifyDate       *time.Time
}

type Account_Note struct {
	Entity

	Account          *Account
	AccountId        *int
	CreateDate       *time.Time
	Customer         *User_Customer
	Id               *int
	ModifyDate       *time.Time
	Note             *string
	NoteHistory      []Account_Note_History
	NoteHistoryCount *uint
	NoteType         *Account_Note_Type
	NoteTypeId       *int
	UserId           *int
}

type Account_Note_History struct {
	Entity

	AccountNote   *Account_Note
	AccountNoteId *int
	CreateDate    *time.Time
	Customer      *User_Customer
	Id            *int
	ModifyDate    *time.Time
	Note          *string
	UserId        *int
}

type Account_Note_Type struct {
	Entity

	BrandId         *int
	CreateDate      *time.Time
	Description     *string
	Id              *int
	KeyName         *string
	ModifyDate      *time.Time
	Name            *string
	ValueExpression *string
}

type Account_Partner_Referral_Prospect struct {
	User_Customer_Prospect

	CompanyName  *string
	EmailAddress *string
	FirstName    *string
	Id           *int
	LastName     *string
}

type Account_Password struct {
	Entity

	Account   *Account
	AccountId *int
	Id        *int
	Notes     *string
	Password  *string
	Type      *Account_Password_Type
	TypeId    *int
	Username  *string
}

type Account_Password_Type struct {
	Entity

	Description *string
}

type Account_Regional_Registry_Detail struct {
	Entity

	Account                          *Account
	AccountId                        *int
	CreateDate                       *time.Time
	DetailCount                      *uint
	DetailType                       *Account_Regional_Registry_Detail_Type
	DetailTypeId                     *int
	Details                          []Network_Subnet_Registration_Details
	Id                               *int
	ModifyDate                       *time.Time
	Properties                       []Account_Regional_Registry_Detail_Property
	PropertyCount                    *uint
	RegionalInternetRegistryHandle   *Account_Rwhois_Handle
	RegionalInternetRegistryHandleId *int
}

type Account_Regional_Registry_Detail_Property struct {
	Entity

	CreateDate           *time.Time
	Detail               *Account_Regional_Registry_Detail
	Id                   *int
	ModifyDate           *time.Time
	PropertyType         *Account_Regional_Registry_Detail_Property_Type
	PropertyTypeId       *int
	RegistrationDetailId *int
	SequencePosition     *int
	Value                *string
}

type Account_Regional_Registry_Detail_Property_Type struct {
	Entity

	CreateDate      *time.Time
	Id              *int
	KeyName         *string
	ModifyDate      *time.Time
	Name            *string
	ValueExpression *string
}

type Account_Regional_Registry_Detail_Type struct {
	Entity

	CreateDate *time.Time
	Id         *int
	KeyName    *string
	ModifyDate *time.Time
	Name       *string
}

type Account_Regional_Registry_Detail_Version4_Person_Default struct {
	Account_Regional_Registry_Detail
}

type Account_Reports_Request struct {
	Entity

	Account                *Account
	AccountContact         *Account_Contact
	AccountContactId       *int
	AccountId              *int
	ComplianceReportTypeId *string
	CreateDate             *time.Time
	EmployeeRecordId       *int
	Id                     *int
	ModifyDate             *time.Time
	Nda                    *string
	Notes                  *string
	Report                 *string
	ReportType             *Compliance_Report_Type
	RequestKey             *string
	Status                 *string
	Ticket                 *Ticket
	TicketId               *int
	User                   *User_Customer
	UsrRecordId            *int
}

type Account_Rwhois_Handle struct {
	Entity

	Account    *Account
	AccountId  *int
	CreateDate *time.Time
	Handle     *string
	Id         *int
	ModifyDate *time.Time
}

type Account_Shipment struct {
	Entity

	Account              *Account
	AccountId            *int
	Courier              *Auxiliary_Shipping_Courier
	CourierId            *int
	CourierName          *string
	CreateEmployee       *User_Employee
	CreateUser           *User_Customer
	CreateUserId         *int
	DestinationAddress   *Account_Address
	DestinationAddressId *int
	DestinationDate      *time.Time
	Id                   *int
	ModifyEmployee       *User_Employee
	ModifyUser           *User_Customer
	ModifyUserId         *int
	Note                 *string
	OriginationAddress   *Account_Address
	OriginationAddressId *int
	OriginationDate      *time.Time
	ShipmentItemCount    *uint
	ShipmentItems        []Account_Shipment_Item
	Status               *Account_Shipment_Status
	StatusId             *int
	TrackingData         []Account_Shipment_Tracking_Data
	TrackingDataCount    *uint
	Type                 *Account_Shipment_Type
	TypeId               *int
}

type Account_Shipment_Item struct {
	Entity

	CreateDate         *time.Time
	Description        *string
	Id                 *int
	PackageId          *int
	Shipment           *Account_Shipment
	ShipmentId         *int
	ShipmentItemId     *int
	ShipmentItemType   *Account_Shipment_Item_Type
	ShipmentItemTypeId *int
}

type Account_Shipment_Item_Type struct {
	Entity

	CreateDate *time.Time
	Id         *int
	KeyName    *string
	Name       *string
}

type Account_Shipment_Resource_Type struct {
	Entity
}

type Account_Shipment_Status struct {
	Entity

	CreateDate *time.Time
	Id         *int
	KeyName    *string
	Name       *string
}

type Account_Shipment_Tracking_Data struct {
	Entity

	CreateEmployee *User_Employee
	CreateUser     *User_Customer
	CreateUserId   *int
	Id             *int
	ModifyEmployee *User_Employee
	ModifyUser     *User_Customer
	ModifyUserId   *int
	PackageId      *int
	Sequence       *int
	Shipment       *Account_Shipment
	ShipmentId     *int
	TrackingData   *string
}

type Account_Shipment_Type struct {
	Entity

	CreateDate  *time.Time
	Description *string
	Id          *int
	KeyName     *string
	Name        *string
}

type Account_Status struct {
	Entity

	Id   *int
	Name *string
}

type Auxiliary_Marketing_Event struct {
	Entity

	CreateDate  *time.Time
	EnabledFlag *int
	EndDate     *time.Time
	Location    *string
	ModifyDate  *time.Time
	StartDate   *time.Time
	Title       *string
	Url         *string
}

type Auxiliary_Network_Status struct {
	Entity
}

type Auxiliary_Notification_Emergency struct {
	Entity

	CreateDate       *time.Time
	Device           *string
	Duration         *string
	Id               *int
	Location         *string
	Message          *string
	ModifyDate       *time.Time
	ServicesAffected *string
	Signature        *Auxiliary_Notification_Emergency_Signature
	StartDate        *time.Time
	Status           *Auxiliary_Notification_Emergency_Status
	StatusId         *int
}

type Auxiliary_Notification_Emergency_Signature struct {
	Entity

	Name *string
}

type Auxiliary_Notification_Emergency_Status struct {
	Entity

	Name *string
}

type Auxiliary_Press_Release struct {
	Entity

	About                []Auxiliary_Press_Release_About_Press_Release
	AboutCount           *uint
	ContactCount         *uint
	Contacts             []Auxiliary_Press_Release_Contact_Press_Release
	Id                   *int
	MediaPartnerCount    *uint
	MediaPartners        []Auxiliary_Press_Release_Media_Partner_Press_Release
	PressReleaseContent  *Auxiliary_Press_Release_Content
	PublishDate          *time.Time
	ReleaseLocation      *string
	SubTitle             *string
	Title                *string
	WebsiteHighlightFlag *bool
}

type Auxiliary_Press_Release_About struct {
	Entity

	Content *string
	Id      *int
	Title   *string
}

type Auxiliary_Press_Release_About_Press_Release struct {
	Entity

	AboutParagraphCount *uint
	AboutParagraphs     []Auxiliary_Press_Release_About
	Id                  *int
	PressReleaseAboutId *int
	PressReleaseCount   *uint
	PressReleaseId      *int
	PressReleases       []Auxiliary_Press_Release
	SortOrder           *int
}

type Auxiliary_Press_Release_Contact struct {
	Entity

	Email             *string
	FirstName         *string
	Id                *int
	LastName          *string
	Phone             *string
	ProfessionalTitle *string
}

type Auxiliary_Press_Release_Contact_Press_Release struct {
	Entity

	ContactCount          *uint
	Contacts              []Auxiliary_Press_Release_Contact
	Id                    *int
	PressReleaseContactId *int
	PressReleaseCount     *uint
	PressReleaseId        *int
	PressReleases         []Auxiliary_Press_Release
	SortOrder             *int
}

type Auxiliary_Press_Release_Content struct {
	Entity

	Id             *int
	PressReleaseId *int
	Text           *string
}

type Auxiliary_Press_Release_Media_Partner struct {
	Entity

	Id   *int
	Name *string
}

type Auxiliary_Press_Release_Media_Partner_Press_Release struct {
	Entity

	Id                *int
	MediaPartnerCount *uint
	MediaPartnerId    *int
	MediaPartners     []Auxiliary_Press_Release_Media_Partner
	PressReleaseCount *uint
	PressReleaseId    *int
	PressReleases     []Auxiliary_Press_Release
}

type Auxiliary_Shipping_Courier struct {
	Entity

	Id      *int
	KeyName *string
	Name    *string
	Url     *string
}

type Auxiliary_Shipping_Courier_Type struct {
	Entity

	Courier      []Auxiliary_Shipping_Courier
	CourierCount *uint
	Description  *string
	Id           *int
	KeyName      *string
	Name         *string
}

type Billing_Currency struct {
	Entity

	Id      *int
	KeyName *string
	Name    *string
}

type Billing_Currency_ExchangeRate struct {
	Entity

	EffectiveDate   *time.Time
	ExpirationDate  *time.Time
	FundingCurrency *Billing_Currency
	Id              *int
	LocalCurrency   *Billing_Currency
	Rate            *float64
}

type Billing_Info struct {
	Entity

	Account                   *Account
	AccountId                 *int
	AchInformation            []Billing_Info_Ach
	AchInformationCount       *uint
	AnniversaryDayOfMonth     *int
	CardAccountNumber         *string
	CardExpirationMonth       *int
	CardExpirationYear        *int
	CardNickname              *string
	CardType                  *string
	CardVerificationNumber    *string
	CreateDate                *time.Time
	Currency                  *Billing_Currency
	CurrentBillingCycle       *Billing_Info_Cycle
	Id                        *int
	LastBillDate              *time.Time
	LastFourPaymentCardDigits *int
	LastPaymentDate           *time.Time
	ModifyDate                *time.Time
	NextBillDate              *time.Time
	PaymentTerms              *int
	PercentDiscountOnetime    *int
	PercentDiscountRecurring  *int
	SparePoolAmount           *int
	VatId                     *string
}

type Billing_Info_Ach struct {
	Entity

	Account           *Account
	AccountId         *int
	AccountNumber     *string
	AccountType       *string
	BankTransitNumber *string
	City              *string
	Country           *string
	FirstName         *string
	Id                *int
	LastName          *string
	PhoneNumber       *string
	Postalcode        *string
	State             *string
	Status            *string
	Street1           *string
	Street2           *string
	VerifiedDate      *time.Time
}

type Billing_Info_Cycle struct {
	Entity

	Account                *Account
	CurrentCycleEndDate    *time.Time
	CurrentCycleStartDate  *time.Time
	NextCycleStartDate     *time.Time
	PreviousCycleEndDate   *time.Time
	PreviousCycleStartDate *time.Time
}

type Billing_Invoice struct {
	Entity

	Account                        *Account
	AccountId                      *int
	Address1                       *string
	Address2                       *string
	Amount                         *float64
	BrandAtInvoiceCreation         *Brand
	City                           *string
	ClaimedTaxExemptTxFlag         *bool
	ClosedDate                     *time.Time
	CompanyName                    *string
	Country                        *string
	CreateDate                     *time.Time
	DetailedPdfGeneratedFlag       *bool
	DocumentsGeneratedFlag         *bool
	Email                          *string
	EndingBalance                  *float64
	FaxPhone                       *string
	FirstName                      *string
	Id                             *int
	InvoiceTopLevelItemCount       *uint
	InvoiceTopLevelItems           []Billing_Invoice_Item
	InvoiceTotalAmount             *float64
	InvoiceTotalOneTimeAmount      *float64
	InvoiceTotalOneTimeTaxAmount   *float64
	InvoiceTotalPreTaxAmount       *float64
	InvoiceTotalRecurringAmount    *float64
	InvoiceTotalRecurringTaxAmount *float64
	ItemCount                      *uint
	Items                          []Billing_Invoice_Item
	LastName                       *string
	ModifyDate                     *time.Time
	OfficePhone                    *string
	Payment                        *float64
	PaymentCount                   *uint
	Payments                       []Billing_Invoice_Receivable_Payment
	PostalCode                     *string
	PurchaseOrderNumber            *string
	SellerRegistration             *string
	StartingBalance                *float64
	State                          *string
	StatusCode                     *string
	TaxInfo                        *Billing_Invoice_Tax_Info
	TaxInfoHistory                 []Billing_Invoice_Tax_Info
	TaxInfoHistoryCount            *uint
	TaxMessage                     *string
	TaxStatusId                    *int
	TaxType                        *Billing_Invoice_Tax_Type
	TaxTypeId                      *int
	TypeCode                       *string
}

type Billing_Invoice_Item struct {
	Entity

	AssociatedChildren              []Billing_Invoice_Item
	AssociatedChildrenCount         *uint
	AssociatedInvoiceItem           *Billing_Invoice_Item
	AssociatedInvoiceItemId         *int
	BillingItem                     *Billing_Item
	BillingItemId                   *int
	Category                        *Product_Item_Category
	CategoryCode                    *string
	Children                        []Billing_Invoice_Item
	ChildrenCount                   *uint
	CreateDate                      *time.Time
	Description                     *string
	DomainName                      *string
	FilteredAssociatedChildren      []Billing_Invoice_Item
	FilteredAssociatedChildrenCount *uint
	HostName                        *string
	HourlyRecurringFee              *float64
	Id                              *int
	Invoice                         *Billing_Invoice
	InvoiceId                       *int
	LaborAfterTaxAmount             *float64
	LaborFee                        *float64
	LaborFeeTaxRate                 *float64
	LaborTaxAmount                  *float64
	Location                        *Location
	NonZeroAssociatedChildren       []Billing_Invoice_Item
	NonZeroAssociatedChildrenCount  *uint
	Notes                           *string
	OneTimeAfterTaxAmount           *float64
	OneTimeFee                      *float64
	OneTimeFeeTaxRate               *float64
	OneTimeTaxAmount                *float64
	Parent                          *Billing_Invoice_Item
	ParentId                        *int
	Product                         *Product_Item
	ProductItemId                   *int
	RecurringAfterTaxAmount         *float64
	RecurringFee                    *float64
	RecurringFeeTaxRate             *float64
	RecurringTaxAmount              *float64
	ResourceTableId                 *int
	SetupAfterTaxAmount             *float64
	SetupFee                        *float64
	SetupFeeTaxRate                 *float64
	SetupTaxAmount                  *float64
	TotalOneTimeAmount              *float64
	TotalOneTimeTaxAmount           *float64
	TotalRecurringAmount            *float64
	TotalRecurringTaxAmount         *float64
}

type Billing_Invoice_Item_Hardware struct {
	Billing_Invoice_Item

	Resource *Hardware
}

type Billing_Invoice_Item_Tax_Info struct {
	Entity

	CreateDate          *time.Time
	Description         *string
	EffectiveTaxRate    *float64
	ExemptAmount        *float64
	FeeProperty         *string
	Id                  *int
	InvoiceItem         *Billing_Invoice_Item
	InvoiceItemId       *int
	InvoiceTaxInfo      *Billing_Invoice_Tax_Info
	InvoiceTaxInfoId    *int
	ModifyDate          *time.Time
	NonTaxableBasis     *float64
	ReportedFlag        *bool
	SellerRegistration  *string
	TaxAmount           *float64
	TaxAmountToCurrency *float64
	TaxRate             *float64
	TaxableBasis        *float64
	ToCurrency          *Billing_Currency
	ToCurrencyId        *int
}

type Billing_Invoice_Next struct {
	Entity
}

type Billing_Invoice_Receivable_Payment struct {
	Entity

	Account                  *Account
	Amount                   *float64
	CreateDate               *time.Time
	CreditCardLastFourDigits *int
	CreditCardRequestId      *string
	CreditCardTransaction    *Billing_Payment_Card_Transaction
	ExchangeRate             *Billing_Currency_ExchangeRate
	Invoice                  *Billing_Invoice
	InvoiceId                *int
	PaypalTransaction        *Billing_Payment_PayPal_Transaction
	TypeCode                 *string
}

type Billing_Invoice_Tax_Info struct {
	Entity

	CreateDate               *time.Time
	Currency                 *Billing_Currency
	CurrencyId               *int
	FunctionalCurrency       *Billing_Currency
	Id                       *int
	Invoice                  *Billing_Invoice
	InvoiceId                *int
	ItemCount                *uint
	ItemWithCurrencyInfo     *Billing_Invoice_Item_Tax_Info
	Items                    []Billing_Invoice_Item_Tax_Info
	ModifyDate               *time.Time
	ReportedFlag             *bool
	TotalTaxAmountToCurrency *float64
}

type Billing_Invoice_Tax_Status struct {
	Entity

	CreateDate *time.Time
	Id         *int
	KeyName    *string
	ModifyDate *time.Time
	Name       *string
}

type Billing_Invoice_Tax_Type struct {
	Entity

	Id      *int
	KeyName *string
	Name    *string
}

type Billing_Item struct {
	Entity

	Account                                            *Account
	ActiveAgreement                                    *Account_Agreement
	ActiveAgreementFlag                                *Account_Agreement
	ActiveAssociatedChildren                           []Billing_Item
	ActiveAssociatedChildrenCount                      *uint
	ActiveAssociatedGuestDiskBillingItemCount          *uint
	ActiveAssociatedGuestDiskBillingItems              []Billing_Item
	ActiveBundledItemCount                             *uint
	ActiveBundledItems                                 []Billing_Item
	ActiveCancellationItem                             *Billing_Item_Cancellation_Request_Item
	ActiveChildren                                     []Billing_Item
	ActiveChildrenCount                                *uint
	ActiveFlag                                         *bool
	ActiveSparePoolAssociatedGuestDiskBillingItemCount *uint
	ActiveSparePoolAssociatedGuestDiskBillingItems     []Billing_Item
	ActiveSparePoolBundledItemCount                    *uint
	ActiveSparePoolBundledItems                        []Billing_Item
	AllowCancellationFlag                              *int
	AssociatedBillingItem                              *Billing_Item
	AssociatedBillingItemHistory                       []Billing_Item_Association_History
	AssociatedBillingItemHistoryCount                  *uint
	AssociatedBillingItemId                            *string
	AssociatedChildren                                 []Billing_Item
	AssociatedChildrenCount                            *uint
	AssociatedParent                                   []Billing_Item
	AssociatedParentCount                              *uint
	AvailableMatchingVlanCount                         *uint
	AvailableMatchingVlans                             []Network_Vlan
	BandwidthAllocation                                *Network_Bandwidth_Version1_Allocation
	BillableChildren                                   []Billing_Item
	BillableChildrenCount                              *uint
	BundleItemCount                                    *uint
	BundleItems                                        []Product_Item_Bundles
	BundledItemCount                                   *uint
	BundledItems                                       []Billing_Item
	CanceledChildren                                   []Billing_Item
	CanceledChildrenCount                              *uint
	CancellationDate                                   *time.Time
	CancellationReason                                 *Billing_Item_Cancellation_Reason
	CancellationRequestCount                           *uint
	CancellationRequests                               []Billing_Item_Cancellation_Request
	Category                                           *Product_Item_Category
	CategoryCode                                       *string
	Children                                           []Billing_Item
	ChildrenCount                                      *uint
	ChildrenWithActiveAgreement                        []Billing_Item
	ChildrenWithActiveAgreementCount                   *uint
	CreateDate                                         *time.Time
	CurrentHourlyCharge                                *string
	CycleStartDate                                     *time.Time
	Description                                        *string
	DomainName                                         *string
	DowngradeItemCount                                 *uint
	DowngradeItems                                     []Product_Item
	FilteredNextInvoiceChildren                        []Billing_Item
	FilteredNextInvoiceChildrenCount                   *uint
	HostName                                           *string
	HourlyFlag                                         *bool
	HourlyRecurringFee                                 *float64
	HoursUsed                                          *string
	Id                                                 *int
	InvoiceItem                                        *Billing_Invoice_Item
	InvoiceItemCount                                   *uint
	InvoiceItems                                       []Billing_Invoice_Item
	Item                                               *Product_Item
	LaborFee                                           *float64
	LaborFeeTaxRate                                    *float64
	LastBillDate                                       *time.Time
	Location                                           *Location
	ModifyDate                                         *time.Time
	NextBillDate                                       *time.Time
	NextInvoiceChildren                                []Billing_Item
	NextInvoiceChildrenCount                           *uint
	NextInvoiceTotalOneTimeAmount                      *float64
	NextInvoiceTotalOneTimeTaxAmount                   *float64
	NextInvoiceTotalRecurringAmount                    *float64
	NextInvoiceTotalRecurringTaxAmount                 *float64
	NonZeroNextInvoiceChildren                         []Billing_Item
	NonZeroNextInvoiceChildrenCount                    *uint
	Notes                                              *string
	OneTimeFee                                         *float64
	OneTimeFeeTaxRate                                  *float64
	OrderItem                                          *Billing_Order_Item
	OrderItemId                                        *int
	OriginalLocation                                   *Location
	Package                                            *Product_Package
	Parent                                             *Billing_Item
	ParentId                                           *int
	ParentVirtualGuestBillingItem                      *Billing_Item_Virtual_Guest
	PendingCancellationFlag                            *bool
	PendingOrderItem                                   *Billing_Order_Item
	ProvisionTransaction                               *Provisioning_Version1_Transaction
	RecurringFee                                       *float64
	RecurringFeeTaxRate                                *float64
	RecurringMonths                                    *int
	ServiceProviderId                                  *int
	SetupFee                                           *float64
	SetupFeeTaxRate                                    *float64
	SoftwareDescription                                *Software_Description
	UpgradeItem                                        *Product_Item
	UpgradeItemCount                                   *uint
	UpgradeItems                                       []Product_Item
}

type Billing_Item_Account_Media_Data_Transfer_Request struct {
	Billing_Item

	Resource *Account_Media_Data_Transfer_Request
}

type Billing_Item_Association_History struct {
	Entity

	AssociatedBillingItem   *Billing_Item
	AssociatedBillingItemId *int
	BillingItem             *Billing_Item
	BillingItemId           *int
	CreateDate              *time.Time
	Id                      *int
}

type Billing_Item_Cancellation_Reason struct {
	Entity

	BillingCancelReasonCategoryId     *int
	BillingCancellationReasonCategory *Billing_Item_Cancellation_Reason_Category
	BillingItemCount                  *uint
	BillingItems                      []Billing_Item
	Id                                *int
	KeyName                           *string
	Reason                            *string
	TranslatedReason                  *string
}

type Billing_Item_Cancellation_Reason_Category struct {
	Entity

	BillingCancellationReasonCount *uint
	BillingCancellationReasons     []Billing_Item_Cancellation_Reason
	Id                             *int
	Name                           *string
}

type Billing_Item_Cancellation_Request struct {
	Entity

	Account               *Account
	AccountId             *int
	BillingCancelReasonId *int
	CreateDate            *time.Time
	Id                    *int
	ItemCount             *uint
	Items                 []Billing_Item_Cancellation_Request_Item
	ModifyDate            *time.Time
	Notes                 *string
	Status                *Billing_Item_Cancellation_Request_Status
	StatusId              *int
	Ticket                *Ticket
	TicketId              *int
	User                  *User_Customer
}

type Billing_Item_Cancellation_Request_Item struct {
	Entity

	BillingItem               *Billing_Item
	BillingItemId             *int
	CancellationRequest       *Billing_Item_Cancellation_Request
	CancellationRequestId     *int
	Id                        *int
	ImmediateCancellationFlag *bool
	ScheduledCancellationDate *time.Time
	ServiceReclaimStatusCode  *string
}

type Billing_Item_Cancellation_Request_Status struct {
	Entity

	Description *string
	Id          *int
	KeyName     *string
	Name        *string
}

type Billing_Item_Ctc_Account struct {
	Billing_Item
}

type Billing_Item_Gateway_Appliance_Cluster struct {
	Billing_Item

	Resource *Resource_Group
}

type Billing_Item_Hardware struct {
	Billing_Item

	BillingCycleBandwidthUsage             []Network_Bandwidth_Usage
	BillingCycleBandwidthUsageCount        *uint
	BillingCyclePrivateBandwidthUsage      []Network_Bandwidth_Usage
	BillingCyclePrivateBandwidthUsageCount *uint
	BillingCyclePrivateUsageIn             *float64
	BillingCyclePrivateUsageOut            *float64
	BillingCyclePrivateUsageTotal          *uint
	BillingCyclePublicBandwidthUsage       []Network_Bandwidth_Usage
	BillingCyclePublicBandwidthUsageCount  *uint
	BillingCyclePublicUsageIn              *float64
	BillingCyclePublicUsageOut             *float64
	BillingCyclePublicUsageTotal           *uint
	LockboxNetworkStorage                  *Billing_Item_Network_Storage
	MonitoringBillingItemCount             *uint
	MonitoringBillingItems                 []Billing_Item
	Resource                               *Hardware_Server
	ResourceTableId                        *int
}

type Billing_Item_Hardware_Colocation struct {
	Billing_Item_Hardware
}

type Billing_Item_Hardware_Component struct {
	Billing_Item

	Resource        []Hardware_Component
	ResourceCount   *uint
	ResourceTableId *int
}

type Billing_Item_Hardware_Security_Module struct {
	Billing_Item_Hardware
}

type Billing_Item_Hardware_Server struct {
	Billing_Item_Hardware
}

type Billing_Item_Link_ThePlanet struct {
	Entity

	BillingItem     *Billing_Item
	ServiceProvider *Service_Provider
}

type Billing_Item_Network_Application_Delivery_Controller struct {
	Billing_Item

	BandwidthAllotmentDetail *Network_Bandwidth_Version1_Allotment_Detail
	Resource                 *Network_Application_Delivery_Controller
}

type Billing_Item_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress struct {
	Billing_Item

	Resource *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress
}

type Billing_Item_Network_Bandwidth struct {
	Billing_Item
}

type Billing_Item_Network_Firewall struct {
	Billing_Item

	Resource *Network_Component_Firewall
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

	Resource *Network_LoadBalancer_Global_Account
}

type Billing_Item_Network_LoadBalancer_VirtualIpAddress struct {
	Billing_Item

	Resource *Network_LoadBalancer_VirtualIpAddress
}

type Billing_Item_Network_Message_Delivery struct {
	Billing_Item

	Resource *Network_Message_Delivery
}

type Billing_Item_Network_Message_Queue struct {
	Billing_Item

	Resource *Network_Message_Queue
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

	Resource *Network_Storage
}

type Billing_Item_Network_Storage_Hub struct {
	Billing_Item_Network_Storage
}

type Billing_Item_Network_Storage_Hub_Bandwidth struct {
	Billing_Item_Network_Storage
}

type Billing_Item_Network_Subnet struct {
	Billing_Item

	Resource        *Network_Subnet
	ResourceName    *string
	ResourceTableId *int
}

type Billing_Item_Network_Subnet_IpAddress_Global struct {
	Billing_Item_Network_Subnet
}

type Billing_Item_Network_Tunnel struct {
	Billing_Item

	Resource *Network_Tunnel_Module_Context
}

type Billing_Item_Network_Vlan struct {
	Billing_Item

	Resource *Network_Vlan
}

type Billing_Item_Software_Component struct {
	Billing_Item

	Resource        *Software_Component
	ResourceTableId *int
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

	Resource *Software_Component
}

type Billing_Item_Software_Component_Virtual_OperatingSystem struct {
	Billing_Item
}

type Billing_Item_Software_Component_Virtual_OperatingSystem_Microsoft struct {
	Billing_Item_Software_Component_Virtual_OperatingSystem

	Resource        *Software_VirtualLicense
	ResourceTableId *int
}

type Billing_Item_Software_Component_Virtual_OperatingSystem_Redhat struct {
	Billing_Item_Software_Component_Virtual_OperatingSystem

	Resource        *Software_Component
	ResourceTableId *int
}

type Billing_Item_Software_License struct {
	Billing_Item

	Resource *Software_AccountLicense
}

type Billing_Item_Support struct {
	Billing_Item
}

type Billing_Item_User_Customer_External_Binding struct {
	Billing_Item

	Resource *User_Customer_External_Binding
}

type Billing_Item_Virtual_Dedicated_Rack struct {
	Billing_Item

	BillingCycleBandwidthUsage             []Network_Bandwidth_Usage
	BillingCycleBandwidthUsageCount        *uint
	BillingCyclePrivateBandwidthUsage      []Network_Bandwidth_Usage
	BillingCyclePrivateBandwidthUsageCount *uint
	BillingCyclePrivateUsageIn             *float64
	BillingCyclePrivateUsageOut            *float64
	BillingCyclePrivateUsageTotal          *uint
	BillingCyclePublicBandwidthUsage       []Network_Bandwidth_Usage
	BillingCyclePublicBandwidthUsageCount  *uint
	BillingCyclePublicUsageIn              *float64
	BillingCyclePublicUsageOut             *float64
	BillingCyclePublicUsageTotal           *uint
	Resource                               *Network_Bandwidth_Version1_Allotment
}

type Billing_Item_Virtual_Disk_Image struct {
	Billing_Item

	Resource        *Virtual_Disk_Image
	ResourceTableId *int
}

type Billing_Item_Virtual_Guest struct {
	Billing_Item

	BillingCycleBandwidthUsage             []Network_Bandwidth_Usage
	BillingCycleBandwidthUsageCount        *uint
	BillingCyclePrivateBandwidthUsage      []Network_Bandwidth_Usage
	BillingCyclePrivateBandwidthUsageCount *uint
	BillingCyclePrivateUsageIn             *float64
	BillingCyclePrivateUsageOut            *float64
	BillingCyclePrivateUsageTotal          *uint
	BillingCyclePublicBandwidthUsage       []Network_Bandwidth_Usage
	BillingCyclePublicBandwidthUsageCount  *uint
	BillingCyclePublicUsageIn              *float64
	BillingCyclePublicUsageOut             *float64
	BillingCyclePublicUsageTotal           *uint
	MonitoringBillingItemCount             *uint
	MonitoringBillingItems                 []Billing_Item
	Resource                               *Virtual_Guest
	ResourceTableId                        *int
}

type Billing_Item_Virtual_Host_Usage struct {
	Billing_Item

	Resource        *Hardware
	ResourceTableId *int
}

type Billing_Item_Workspace struct {
	Billing_Item
}

type Billing_Order struct {
	Entity

	Account                      *Account
	AccountId                    *int
	Brand                        *Brand
	Cart                         *Billing_Order_Cart
	CoreRestrictedItemCount      *uint
	CoreRestrictedItems          []Billing_Order_Item
	CreateDate                   *time.Time
	CreditCardTransactionCount   *uint
	CreditCardTransactions       []Billing_Payment_Card_Transaction
	ExchangeRate                 *Billing_Currency_ExchangeRate
	Id                           *int
	ImpersonatingUserRecordId    *int
	InitialInvoice               *Billing_Invoice
	ItemCount                    *uint
	Items                        []Billing_Order_Item
	ModifyDate                   *time.Time
	OrderApprovalDate            *time.Time
	OrderNonServerMonthlyAmount  *float64
	OrderQuoteId                 *int
	OrderServerMonthlyAmount     *float64
	OrderTopLevelItemCount       *uint
	OrderTopLevelItems           []Billing_Order_Item
	OrderTotalAmount             *float64
	OrderTotalOneTime            *float64
	OrderTotalOneTimeAmount      *float64
	OrderTotalOneTimeTaxAmount   *float64
	OrderTotalRecurring          *float64
	OrderTotalRecurringAmount    *float64
	OrderTotalRecurringTaxAmount *float64
	OrderTotalSetupAmount        *float64
	OrderType                    *Billing_Order_Type
	OrderTypeId                  *int
	PaypalTransactionCount       *uint
	PaypalTransactions           []Billing_Payment_PayPal_Transaction
	PresaleEvent                 *Sales_Presale_Event
	PresaleEventId               *int
	PrivateCloudOrderFlag        *bool
	Quote                        *Billing_Order_Quote
	ReferralPartner              *Account
	Status                       *string
	UpgradeRequestFlag           *bool
	UserRecord                   *User_Customer
	UserRecordId                 *int
}

type Billing_Order_Cart struct {
	Billing_Order_Quote
}

type Billing_Order_Item struct {
	Entity

	BillingItem               *Billing_Item
	BundledItemCount          *uint
	BundledItems              []Billing_Order_Item
	Category                  *Product_Item_Category
	CategoryCode              *string
	Children                  []Billing_Order_Item
	ChildrenCount             *uint
	Description               *string
	DomainName                *string
	GlobalIdentifier          *string
	HardwareGenericComponent  *Hardware_Component_Model_Generic
	HostName                  *string
	HourlyRecurringFee        *float64
	Id                        *int
	Item                      *Product_Item
	ItemCategoryAnswerCount   *uint
	ItemCategoryAnswers       []Billing_Order_Item_Category_Answer
	ItemId                    *int
	ItemPrice                 *Product_Item_Price
	ItemPriceId               *float64
	LaborAfterTaxAmount       *float64
	LaborFee                  *float64
	LaborFeeTaxRate           *float64
	LaborTaxAmount            *float64
	Location                  *Location
	NextOrderChildren         []Billing_Order_Item
	NextOrderChildrenCount    *uint
	OldBillingItem            *Billing_Item
	OneTimeAfterTaxAmount     *float64
	OneTimeFee                *float64
	OneTimeFeeTaxRate         *float64
	OneTimeTaxAmount          *float64
	Order                     *Billing_Order
	OrderApprovalDate         *time.Time
	Package                   *Product_Package
	Parent                    *Billing_Order_Item
	ParentId                  *int
	PromoCodeId               *int
	Quantity                  *int
	RecurringAfterTaxAmount   *float64
	RecurringFee              *float64
	RecurringTaxAmount        *float64
	RedundantPowerSupplyCount *uint
	SetupAfterTaxAmount       *float64
	SetupFee                  *float64
	SetupFeeDeferralMonths    *int
	SetupFeeTaxRate           *float64
	SetupTaxAmount            *float64
	SoftwareDescription       *Software_Description
	StorageGroupCount         *uint
	StorageGroups             []Configuration_Storage_Group_Order
	TotalRecurringAmount      *float64
	UpgradeItem               *Product_Item
}

type Billing_Order_Item_Category_Answer struct {
	Entity

	Answer     *string
	CreateDate *time.Time
	OrderItem  *Billing_Order_Item
	Question   *Product_Item_Category_Question
	QuestionId *int
}

type Billing_Order_Note struct {
	Entity

	CreateDate *time.Time
	Employee   *User_Employee
	Order      *Billing_Order
}

type Billing_Order_Quote struct {
	Entity

	Account                 *Account
	AccountId               *int
	CompletedPurchaseDataId *int
	CreateDate              *time.Time
	ExpirationDate          *time.Time
	Id                      *int
	ModifyDate              *time.Time
	Name                    *string
	Order                   *Billing_Order
	OrdersFromQuote         []Billing_Order
	OrdersFromQuoteCount    *uint
	PublicNote              *string
	QuoteKey                *string
	Status                  *string
}

type Billing_Order_Type struct {
	Entity

	Description *string
	Id          *int
	Type        *string
}

type Billing_Payment_Card_ChangeRequest struct {
	Entity

	Account                         *Account
	AccountId                       *int
	Amount                          *float64
	AuthorizedCreditCardTransaction *Billing_Payment_Card_Transaction
	BillingAddressLine1             *string
	BillingAddressLine2             *string
	BillingCity                     *string
	BillingCountryCode              *string
	BillingEmail                    *string
	BillingNameCompany              *string
	BillingNameFirst                *string
	BillingNameLast                 *string
	BillingPhoneFax                 *string
	BillingPhoneVoice               *string
	BillingPostalCode               *string
	BillingState                    *string
	CaptureCreditCardTransaction    *Billing_Payment_Card_Transaction
	CardAccountLast4                *string
	CardAccountNumber               *string
	CardExpirationMonth             *string
	CardExpirationYear              *string
	CardNickname                    *string
	CardType                        *string
	CreditCardVerificationNumber    *string
	CurrencyShortName               *string
	DeviceFingerprintId             *string
	Id                              *int
	Notes                           *string
	PaymentRoleId                   *int
	PaymentType                     *string
	TicketAttachmentReferenceCount  *uint
	TicketAttachmentReferences      []Ticket_Attachment
	TicketId                        *int
}

type Billing_Payment_Card_ManualPayment struct {
	Entity

	Account                           *Account
	AccountId                         *int
	Amount                            *float64
	AuthorizedCreditCardTransaction   *Billing_Payment_Card_Transaction
	AuthorizedCreditCardTransactionId *int
	AuthorizedPayPalTransaction       *Billing_Payment_PayPal_Transaction
	AuthorizedPayPalTransactionId     *int
	BillingAddressLine1               *string
	BillingAddressLine2               *string
	BillingCity                       *string
	BillingCountryCode                *string
	BillingEmail                      *string
	BillingNameCompany                *string
	BillingNameFirst                  *string
	BillingNameLast                   *string
	BillingPhoneFax                   *string
	BillingPhoneVoice                 *string
	BillingPostalCode                 *string
	BillingState                      *string
	CancelUrl                         *string
	CaptureCreditCardTransaction      *Billing_Payment_Card_Transaction
	CapturePayPalTransaction          *Billing_Payment_PayPal_Transaction
	CardAccountHash                   *string
	CardAccountLast4                  *string
	CardAccountNumber                 *string
	CardExpirationMonth               *string
	CardExpirationYear                *string
	CardType                          *string
	CreditCardVerificationNumber      *string
	CurrencyShortName                 *string
	DeviceFingerprintId               *string
	FromIpAddress                     *string
	Id                                *int
	Notes                             *string
	PaymentType                       *string
	ReturnUrl                         *string
	TicketAttachmentReferenceCount    *uint
	TicketAttachmentReferences        []Ticket_Attachment
	Type                              *string
}

type Billing_Payment_Card_Transaction struct {
	Billing_Payment_Transaction

	Account             *Account
	AccountId           *int
	Amount              *float64
	BillingAddressLine1 *string
	BillingAddressLine2 *string
	BillingCity         *string
	BillingCountryCode  *string
	BillingEmail        *string
	BillingNameCompany  *string
	BillingNameFirst    *string
	BillingNameLast     *string
	BillingPhoneFax     *string
	BillingPhoneVoice   *string
	BillingPostalCode   *string
	BillingState        *string
	CardAccountLast4    *int
	CardExpirationMonth *int
	CardExpirationYear  *int
	CardType            *string
	CreateDate          *time.Time
	Id                  *int
	InvoiceId           *int
	ModifyDate          *time.Time
	Order               *Billing_Order
	OrderFromIpAddress  *string
	ReferenceCode       *string
	RequestId           *string
	ReturnStatus        *int
	SerializedReply     *string
	SerializedRequest   *string
}

type Billing_Payment_PayPal_Transaction struct {
	Billing_Payment_Transaction

	Account              *Account
	AccountId            *int
	AddressCityName      *string
	AddressCountry       *string
	AddressName          *string
	AddressPostalCode    *string
	AddressStateProvence *string
	AddressStatus        *string
	AddressStreet1       *string
	AddressStreet2       *string
	ContactPhone         *string
	CreateDate           *time.Time
	ExchangeRate         *string
	FeeAmount            *float64
	GrossAmount          *float64
	Id                   *int
	InvoiceId            *int
	LastPaypalCommand    *string
	ModifyDate           *time.Time
	Order                *Billing_Order
	OrderFromIpAddress   *string
	OrderTotal           *float64
	Payer                *string
	PayerBusiness        *string
	PayerCountry         *string
	PayerFirstName       *string
	PayerId              *string
	PayerLastName        *string
	PayerStatus          *string
	PaymentDate          *time.Time
	PaymentStatus        *string
	PaymentType          *string
	PendingReason        *string
	SerializedReply      *string
	SerializedRequest    *string
	SettleAmount         *float64
	TaxAmount            *float64
	Token                *string
	TransactionId        *string
	TransactionType      *string
}

type Billing_Payment_Processor struct {
	Entity

	BrandAssignmentCount *uint
	BrandAssignments     []Brand_Payment_Processor
	Description          *string
	Name                 *string
	OwnerAccount         *Account
	PaymentMethodCount   *uint
	PaymentMethods       []Billing_Payment_Processor_Method
	Type                 *Billing_Payment_Processor_Type
}

type Billing_Payment_Processor_Method struct {
	Entity

	MethodKey            *string
	MultipleCurrencyFlag *bool
	PaymentProcessor     *Billing_Payment_Processor
	PaymentType          *Billing_Payment_Type
}

type Billing_Payment_Processor_Type struct {
	Entity

	Description           *string
	KeyName               *string
	Name                  *string
	PaymentProcessorCount *uint
	PaymentProcessors     []Billing_Payment_Processor
}

type Billing_Payment_Transaction struct {
	Entity
}

type Billing_Payment_Type struct {
	Entity

	Description *string
	KeyName     *string
	Name        *string
}

type Brand struct {
	Entity

	Account                                 *Account
	AllOwnedAccountCount                    *uint
	AllOwnedAccounts                        []Account
	AllowAccountCreationFlag                *bool
	Catalog                                 *Product_Catalog
	CatalogId                               *int
	ContactCount                            *uint
	Contacts                                []Brand_Contact
	CustomerCountryLocationRestrictionCount *uint
	CustomerCountryLocationRestrictions     []Brand_Restriction_Location_CustomerCountry
	Distributor                             *Brand
	DistributorChildFlag                    *bool
	DistributorFlag                         *string
	Hardware                                []Hardware
	HardwareCount                           *uint
	HasAgentSupportFlag                     *bool
	Id                                      *int
	KeyName                                 *string
	LongName                                *string
	Name                                    *string
	OpenTicketCount                         *uint
	OpenTickets                             []Ticket
	OwnedAccountCount                       *uint
	OwnedAccounts                           []Account
	TicketCount                             *uint
	TicketGroupCount                        *uint
	TicketGroups                            []Ticket_Group
	Tickets                                 []Ticket
	UserCount                               *uint
	Users                                   []User_Customer
	VirtualGuestCount                       *uint
	VirtualGuests                           []Virtual_Guest
}

type Brand_Attribute struct {
	Entity

	Brand *Brand
}

type Brand_Contact struct {
	Entity

	Address1           *string
	Address2           *string
	AlternatePhone     *string
	Brand              *Brand
	BrandContactType   *Brand_Contact_Type
	BrandContactTypeId *int
	City               *string
	Country            *string
	Email              *string
	FaxPhone           *string
	FirstName          *string
	LastName           *string
	OfficePhone        *string
	PostalCode         *string
	State              *string
}

type Brand_Contact_Type struct {
	Entity

	Description *string
	KeyName     *string
	Name        *string
}

type Brand_Payment_Processor struct {
	Entity

	Brand            *Brand
	PaymentProcessor *Billing_Payment_Processor
}

type Brand_Restriction_Location_CustomerCountry struct {
	Entity

	Brand               *Brand
	BrandId             *int
	CustomerCountryCode *string
	Location            *Location
	LocationId          *int
}

type Catalyst_Affiliate struct {
	Entity

	Id                             *int
	Name                           *string
	SkipCreditCardVerificationFlag *bool
}

type Catalyst_Company_Type struct {
	Entity

	Description *string
	Id          *int
}

type Catalyst_Enrollment struct {
	Entity

	Account                  *Account
	AccountId                *int
	Affiliate                *Catalyst_Affiliate
	AffiliateId              *int
	AgreementCompleteFlag    *int
	CompanyDescription       *string
	CompanyType              *Catalyst_Company_Type
	CompanyTypeId            *int
	EnrollmentDate           *time.Time
	GraduationDate           *time.Time
	IsActiveFlag             *bool
	MonthlyCreditAmount      *float64
	Representative           *User_Employee
	RepresentativeEmployeeId *int
}

type Catalyst_Enrollment_Request struct {
	Entity

	Address1                    *string
	Address2                    *string
	Affiliate                   *Catalyst_Affiliate
	AffiliateId                 *int
	AgreementCompleteFlag       *bool
	ApplyToGepFlag              *bool
	CardAccountNumber           *string
	CardExpirationMonth         *string
	CardExpirationYear          *string
	CardType                    *string
	CardVerificationNumber      *string
	City                        *string
	CompanyDescription          *string
	CompanyName                 *string
	CompanyType                 *Catalyst_Company_Type
	CompanyTypeId               *int
	CompanyUrl                  *string
	Country                     *string
	CurrentUserChoice           *int
	DeviceFingerprintId         *string
	Email                       *string
	FirstName                   *string
	FutureUserChoice            *int
	IncubatorName               *string
	InvestorName                *string
	IpAddress                   *string
	LastName                    *string
	OfficePhone                 *string
	OverFiveYearsOldFlag        *bool
	PostalCode                  *string
	ReferralCode                *string
	RevenueOverOneMillionFlag   *bool
	SkipCatalystApplicationFlag *bool
	State                       *string
	VatId                       *string
}

type Catalyst_Enrollment_Request_Container_AnswerOption struct {
	Entity

	Answer *string
	Index  *int
}

type Compliance_Report_Type struct {
	Entity

	Id      *int
	KeyName *string
	Name    *string
}

type Configuration_Storage_Filesystem_Type struct {
	Entity

	KeyName *string
	Name    *string
}

type Configuration_Storage_Group_Array_Type struct {
	Entity

	Description                 *string
	DriveMultiplier             *int
	HardwareComponentModelCount *uint
	HardwareComponentModels     []Hardware_Component_Model
	HotspareAllow               *bool
	Id                          *int
	KeyName                     *string
	MaximumDrives               *int
	MinimumDrives               *int
	Name                        *string
}

type Configuration_Storage_Group_Order struct {
	Entity

	ArrayNumber        *int
	ArraySize          *float64
	ArrayType          *Configuration_Storage_Group_Array_Type
	ArrayTypeId        *int
	BillingOrderItem   *Billing_Order_Item
	BillingOrderItemId *int
	HardDrives         []int
	HotSpareDrives     []int
	PartitionData      *string
}

type Configuration_Storage_Group_Template_Group struct {
	Entity

	Grow             *bool
	HardDrivesString *string
	OrderIndex       *int
	Size             *float64
	Type             *Configuration_Storage_Group_Array_Type
}

type Configuration_Template struct {
	Entity

	Account                             *Account
	AccountId                           *int
	ConfigurationSectionCount           *uint
	ConfigurationSections               []Configuration_Template_Section
	ConfigurationTemplateReference      []Monitoring_Agent_Configuration_Template_Group_Reference
	ConfigurationTemplateReferenceCount *uint
	CreateDate                          *time.Time
	DefaultValueCount                   *uint
	DefaultValues                       []Configuration_Template_Section_Definition_Value
	DefinitionCount                     *uint
	Definitions                         []Configuration_Template_Section_Definition
	Description                         *string
	Id                                  *int
	Item                                *Product_Item
	ItemId                              *int
	LinkedSectionReferences             *Configuration_Template_Section_Reference
	ModifyDate                          *time.Time
	Name                                *string
	Parent                              *Configuration_Template
	ParentId                            *int
	User                                *User_Customer
	UserRecordId                        *int
}

type Configuration_Template_Attribute struct {
	Entity

	ConfigurationTemplate *Configuration_Template
	Value                 *string
}

type Configuration_Template_Section struct {
	Entity

	CreateDate              *time.Time
	DefinitionCount         *uint
	Definitions             []Configuration_Template_Section_Definition
	Description             *string
	DisallowedDeletionFlag  *bool
	Id                      *int
	LinkedTemplate          *Configuration_Template
	LinkedTemplateId        *string
	LinkedTemplateReference *Configuration_Template_Section_Reference
	ModifyDate              *time.Time
	Name                    *string
	ParentId                *int
	ProfileCount            *uint
	Profiles                []Configuration_Template_Section_Profile
	SectionType             *Configuration_Template_Section_Type
	SectionTypeName         *string
	Sort                    *int
	SubSectionCount         *uint
	SubSections             []Configuration_Template_Section
	Template                *Configuration_Template
	TemplateId              *string
	TypeId                  *int
}

type Configuration_Template_Section_Attribute struct {
	Entity

	ConfigurationSection *Configuration_Template_Section
	Value                *string
}

type Configuration_Template_Section_Definition struct {
	Entity

	AttributeCount     *uint
	Attributes         []Configuration_Template_Section_Definition_Attribute
	CreateDate         *time.Time
	DefaultValue       *Configuration_Template_Section_Definition_Value
	Description        *string
	EnumerationValues  *string
	Group              *Configuration_Template_Section_Definition_Group
	GroupId            *string
	Id                 *int
	MaximumValue       *string
	MinimumValue       *string
	ModifyDate         *time.Time
	MonitoringDataFlag *bool
	Name               *string
	Path               *string
	RequireValueFlag   *int
	Section            *Configuration_Template_Section
	SectionId          *int
	ShortName          *string
	Sort               *int
	TypeId             *int
	ValueType          *Configuration_Template_Section_Definition_Type
}

type Configuration_Template_Section_Definition_Attribute struct {
	Entity

	AttributeType           *Configuration_Template_Section_Definition_Attribute_Type
	ConfigurationDefinition *Configuration_Template_Section_Definition
	Value                   *string
}

type Configuration_Template_Section_Definition_Attribute_Type struct {
	Entity

	Description *string
	Name        *string
}

type Configuration_Template_Section_Definition_Group struct {
	Entity

	CreateDate  *time.Time
	Description *string
	Id          *int
	Name        *string
	Parent      *Configuration_Template_Section_Definition_Group
	SortOrder   *int
}

type Configuration_Template_Section_Definition_Type struct {
	Entity

	Description *string
	Id          *int
	Name        *string
}

type Configuration_Template_Section_Definition_Value struct {
	Entity

	CreateDate   *time.Time
	Definition   *Configuration_Template_Section_Definition
	DefinitionId *int
	ModifyDate   *time.Time
	Template     *Configuration_Template
	TemplateId   *int
	Value        *string
}

type Configuration_Template_Section_Profile struct {
	Entity

	AgentId              *int
	ConfigurationSection *Configuration_Template_Section
	CreateDate           *time.Time
	Id                   *int
	MonitoringAgent      *Monitoring_Agent
	Name                 *string
	SectionId            *int
}

type Configuration_Template_Section_Reference struct {
	Entity

	CreateDate *time.Time
	Id         *int
	ModifyDate *time.Time
	Section    *Configuration_Template_Section
	SectionId  *int
	Template   *Configuration_Template
	TemplateId *int
}

type Configuration_Template_Section_Type struct {
	Entity

	Description *string
	Id          *int
	Name        *string
}

type Configuration_Template_Type struct {
	Entity

	CreateDate  *time.Time
	Description *string
	Id          *int
	Name        *string
}

type Container_Account_Discount_Program struct {
	Entity

	AppliedCredit           *float64
	IsParticipant           *bool
	LifetimeAppliedCredit   *float64
	LifetimeCredit          *float64
	LifetimeRemainingCredit *float64
	MaximumActiveOrders     *float64
	MonthlyCredit           *float64
	PostTaxRemainingCredit  *float64
	ProgramEndDate          *time.Time
	ProgramName             *string
	RemainingCredit         *float64
	RemainingCreditTax      *float64
}

type Container_Account_Graph_Outputs struct {
	Entity

	ClosedTickets                      *string
	CompletedBackupCount               *string
	ConflictBackupCount                *string
	EndDate                            *time.Time
	FailedBackupCount                  *string
	GraphError                         *string
	GraphImage                         *[]byte
	HardwareUptime                     *string
	InboundUsage                       *string
	OpenTickets                        *string
	OutboundUsage                      *string
	PendingCustomerResponseTicketCount *string
	StartDate                          *time.Time
	UrlUptime                          *string
	WaitingEmployeeResponseTicketCount *string
}

type Container_Account_Historical_Summary struct {
	Entity

	Details   []Container_Account_Historical_Summary_Detail
	EndDate   *time.Time
	StartDate *time.Time
}

type Container_Account_Historical_Summary_Detail struct {
	Entity

	EndDate   *time.Time
	StartDate *time.Time
}

type Container_Account_Historical_Summary_Detail_Uptime struct {
	Container_Account_Historical_Summary_Detail

	CloudComputingInstance *Virtual_Guest
	ConfigurationValue     *Monitoring_Agent_Configuration_Value
	Data                   []Metric_Tracking_Object_Data
	Hardware               *Hardware
}

type Container_Account_Historical_Summary_Uptime struct {
	Container_Account_Historical_Summary
}

type Container_Account_Payment_Method_CreditCard struct {
	Entity

	Address1                    *string
	Address2                    *string
	City                        *string
	Country                     *string
	CurrencyShortName           *string
	CybersourceAssignedCardType *string
	ExpireMonth                 *string
	ExpireYear                  *string
	FirstName                   *string
	LastFourDigits              *string
	LastName                    *string
	Nickname                    *string
	PaymentMethodRoleName       *string
	PaymentTypeId               *string
	PaymentTypeName             *string
	PostalCode                  *string
	State                       *string
}

type Container_Auxiliary_Network_Status_Reading struct {
	Entity

	AveragePing   *float64
	Fails         *int
	Frequency     *int
	Label         *string
	LastCheckDate *time.Time
	LastDownDate  *time.Time
	Latency       *float64
	Location      *string
	MaximumPing   *float64
	MinimumPing   *float64
	PingLoss      *float64
	StartDate     *time.Time
	StatusCode    *string
	StatusMessage *string
	Target        *string
	TargetType    *string
}

type Container_Bandwidth_GraphInputs struct {
	Entity

	EndDate            *time.Time
	NetworkInterfaceId *int
	Pod                *int
	ServerName         *string
	StartDate          *time.Time
}

type Container_Bandwidth_GraphOutputs struct {
	Entity

	GraphImage   *[]byte
	GraphTitle   *string
	MaxEndDate   *time.Time
	MinStartDate *time.Time
}

type Container_Bandwidth_GraphOutputsExtended struct {
	Entity

	GraphImage         *[]byte
	GraphTitle         *string
	InBoundTotalBytes  *uint
	MaxEndDate         *time.Time
	MinStartDate       *time.Time
	OutBoundTotalBytes *uint
}

type Container_Bandwidth_Projection struct {
	Entity

	AllowedUsage   *string
	EstimatedUsage *string
	HardwareId     *int
	ProjectedUsage *string
	ServerName     *string
	StartDate      *time.Time
}

type Container_Billing_Currency_Format struct {
	Entity

	Currency  *string
	Display   *int
	Format    *string
	Locale    *string
	Name      *string
	Position  *int
	Precision *int
	Script    *string
	Service   *string
	Symbol    *string
	Tag       *string
	Value     *float64
}

type Container_Billing_Info_Ach struct {
	Entity

	AccountNumber     *string
	AccountType       *string
	BankTransitNumber *string
	City              *string
	Country           *string
	FederalTaxId      *string
	FirstName         *string
	LastName          *string
	PhoneNumber       *string
	PostalCode        *string
	State             *string
	Street1           *string
	Street2           *string
}

type Container_Billing_Invoice_Email struct {
	Entity

	ExcelInvoiceIds       []int
	PdfDetailedInvoiceIds []int
	PdfInvoiceIds         []int
	Type                  *string
}

type Container_Billing_Order_Status struct {
	Entity

	Description *string
	Status      *string
}

type Container_Catalyst_ManualEnrollmentRequest struct {
	Entity

	CustomerEmail          *string
	CustomerName           *string
	StartupName            *string
	VentureAffiliationFlag *bool
	VentureFundName        *string
}

type Container_Collection_Locale_CountryCode struct {
	Entity

	LongName   *string
	ShortName  *string
	StateCodes []Container_Collection_Locale_StateCode
}

type Container_Collection_Locale_StateCode struct {
	Entity

	LongName  *string
	ShortName *string
}

type Container_Disk_Image_Capture_Template struct {
	Entity

	Description *string
	Name        *string
	Summary     *string
	Volumes     []Container_Disk_Image_Capture_Template_Volume
}

type Container_Disk_Image_Capture_Template_Volume struct {
	Entity

	Name       *string
	Partitions []Container_Disk_Image_Capture_Template_Volume_Partition
}

type Container_Disk_Image_Capture_Template_Volume_Partition struct {
	Entity

	Name *string
}

type Container_Dns_Domain_Registration_Contact struct {
	Entity

	Address1         *string
	Address2         *string
	Address3         *string
	City             *string
	Country          *string
	Email            *string
	Fax              *string
	FirstName        *string
	LastName         *string
	OrganizationName *string
	Phone            *string
	PostalCode       *string
	State            *string
	Type             *string
}

type Container_Dns_Domain_Registration_ExtendedAttribute struct {
	Entity

	ChildFlag       *bool
	Description     *string
	Name            *string
	Options         []Container_Dns_Domain_Registration_ExtendedAttribute_Option
	RequiredFlag    *int
	UserDefinedFlag *bool
}

type Container_Dns_Domain_Registration_ExtendedAttribute_Configuration struct {
	Entity

	Name  *string
	Value *string
}

type Container_Dns_Domain_Registration_ExtendedAttribute_Option struct {
	Entity

	Description               *string
	RequireExtendedAttributes []Container_Dns_Domain_Registration_ExtendedAttribute_Option_Require
	Title                     *string
	Value                     *string
}

type Container_Dns_Domain_Registration_ExtendedAttribute_Option_Require struct {
	Entity

	Name *string
}

type Container_Dns_Domain_Registration_Information struct {
	Entity

	Contacts           []Container_Dns_Domain_Registration_Contact
	ExpireDate         *time.Time
	Nameservers        []Container_Dns_Domain_Registration_Nameserver
	RegistryCreateDate *time.Time
	RegistryExpireDate *time.Time
	RegistryUpdateDate *time.Time
}

type Container_Dns_Domain_Registration_List struct {
	Entity

	DomainName                     *string
	EncodingType                   *string
	ExtendedAttributeConfiguration []Container_Dns_Domain_Registration_ExtendedAttribute_Configuration
	RegistrationPeriod             *int
}

type Container_Dns_Domain_Registration_Lookup struct {
	Entity

	Items []Container_Dns_Domain_Registration_Lookup_Items
}

type Container_Dns_Domain_Registration_Lookup_Items struct {
	Entity

	DomainName *string
	Status     *string
}

type Container_Dns_Domain_Registration_Nameserver struct {
	Entity

	Nameservers []Container_Dns_Domain_Registration_Nameserver_List
}

type Container_Dns_Domain_Registration_Nameserver_List struct {
	Entity

	Ipv4Address *string
	Ipv6Address *string
	Name        *string
	SortOrder   *int
}

type Container_Dns_Domain_Registration_Registrant_Verification_StatusDetail struct {
	Entity

	Status                   *Dns_Domain_Registration_Registrant_Verification_Status
	VerificationDeadlineDate *time.Time
}

type Container_Dns_Domain_Registration_Transfer_Information struct {
	Entity

	Reason          *string
	RegistrantEmail *string
	Status          *string
	TimeStamp       *time.Time
	Transferrable   *int
}

type Container_Exception struct {
	Entity

	ExceptionClass   *string
	ExceptionMessage *string
}

type Container_Graph struct {
	Entity

	BaseUnit      *string
	EndDatetime   *string
	Height        *int
	Image         *[]byte
	Interval      *int
	Metrics       []Container_Metric_Data_Type
	NormalizeFlag *[]byte
	Options       []Container_Graph_Option
	Plots         []Container_Graph_Plot
	ReturnImage   *bool
	StartDatetime *string
	Template      *string
	Title         *string
	Width         *int
}

type Container_Graph_Option struct {
	Entity

	Name  *string
	Value *string
}

type Container_Graph_Plot struct {
	Entity

	Data   []Container_Graph_Plot_Coordinate
	Metric *Container_Metric_Data_Type
	Unit   *string
}

type Container_Graph_Plot_Coordinate struct {
	Entity

	XValue *float64
	YValue *float64
	ZValue *float64
}

type Container_Hardware_Configuration struct {
	Entity

	Datacenters               []Container_Hardware_Configuration_Option
	FixedConfigurationPresets []Container_Hardware_Configuration_Option
	HardDrives                []Container_Hardware_Configuration_Option
	NetworkComponents         []Container_Hardware_Configuration_Option
	OperatingSystems          []Container_Hardware_Configuration_Option
	Processors                []Container_Hardware_Configuration_Option
}

type Container_Hardware_Configuration_Option struct {
	Entity

	ItemPrice *Product_Item_Price
	Preset    *Product_Package_Preset
	Template  *Hardware
}

type Container_Hardware_MassUpdate struct {
	Entity

	HardwareId  *int
	Message     *string
	SuccessFlag *string
}

type Container_Hardware_Server_Configuration struct {
	Entity

	AddToSparePoolAfterOsReload *int
	CustomProvisionScriptUri    *string
	DriveRetentionFlag          *bool
	EraseHardDrives             *int
	HardDrives                  []Hardware_Component
	ImageTemplateId             *int
	ItemPrices                  []Product_Item_Price
	ResetIpmiPassword           *int
	SshKeyIds                   []int
	UpgradeBios                 *int
	UpgradeHardDriveFirmware    *int
}

type Container_Hardware_Server_Details struct {
	Entity

	Components        []Hardware_Component
	NetworkComponents []Network_Component
	Software          []Software_Component
}

type Container_KnowledgeLayer_QuestionAnswer struct {
	Entity

	Answer   *string
	Link     *string
	Question *string
}

type Container_Message struct {
	Entity

	Message *string
	Type    *string
}

type Container_Metric_Data_Type struct {
	Entity

	KeyName     *string
	Name        *string
	SummaryType *string
	Unit        *string
}

type Container_Metric_Tracking_Object_Details struct {
	Entity

	MetricName *string
}

type Container_Metric_Tracking_Object_Summary struct {
	Entity

	MetricName *string
}

type Container_Metric_Tracking_Object_Virtual_Host_Details struct {
	Container_Metric_Tracking_Object_Details

	Day             *time.Time
	MaxInstances    *int
	MaxMemoryUsage  *int
	MeanInstances   *float64
	MeanMemoryUsage *float64
	MinInstances    *int
	MinMemoryUsage  *int
}

type Container_Metric_Tracking_Object_Virtual_Host_Summary struct {
	Container_Metric_Tracking_Object_Summary

	AvgMemoryUsageInBillingCycle *int
	CurrentBillCycleEnd          *time.Time
	CurrentBillCycleStart        *time.Time
	LastInstanceCount            *int
	LastMemoryUsageAmount        *int
	LastPollTime                 *time.Time
	MaxInstanceInBillingCycle    *int
	PreviousBillCycleEnd         *time.Time
	PreviousBillCycleStart       *time.Time
	VirtualPlatformName          *string
}

type Container_Monitoring_Alarm_History struct {
	Entity

	AccountId  *int
	AgentId    *int
	AlarmId    *string
	ClosedDate *time.Time
	CreateDate *time.Time
	Message    *string
	RobotId    *int
	Severity   *string
}

type Container_Monitoring_Graph_Outputs struct {
	Entity

	EndDate    *time.Time
	GraphError *string
	GraphImage *[]byte
	StartDate  *time.Time
}

type Container_Network_Authentication_Data struct {
	Entity

	Host     *string
	Password *string
	Port     *int
	Type     *string
	Username *string
}

type Container_Network_Bandwidth_Data_Summary struct {
	Entity

	AllowedUsage   *float64
	EstimatedUsage *float64
	ProjectedUsage *float64
	UsageUnits     *string
}

type Container_Network_Bandwidth_Version1_Usage struct {
	Entity

	IncomingAmount *float64
	OutgoingAmount *float64
	RecordedDate   *time.Time
}

type Container_Network_ContentDelivery_Authentication_Directory struct {
	Entity

	CreateDate *time.Time
	Name       *string
	Type       *string
}

type Container_Network_ContentDelivery_Authentication_Parameter struct {
	Entity

	CdnAccountName *string
	ClientIp       *string
	Referrer       *string
	SourceUrl      *string
	Token          *string
}

type Container_Network_ContentDelivery_Authentication_ServiceEndpoint struct {
	Entity

	Endpoint *string
	Protocol *string
}

type Container_Network_ContentDelivery_Bandwidth_PointsOfPresence_Summary struct {
	Entity

	Bandwidth     *uint
	EndDateTime   *time.Time
	PopName       *string
	StartDateTime *time.Time
	UsageUnits    *string
	ViewCount     *uint
}

type Container_Network_ContentDelivery_Bandwidth_Summary struct {
	Entity

	CdnAccountId  *int
	EndDateTime   *time.Time
	FileName      *string
	MediaType     *string
	StartDateTime *time.Time
	Usage         *float64
	UsageUnits    *string
}

type Container_Network_ContentDelivery_Bandwidth_Summary_Detail struct {
	Container_Network_ContentDelivery_Bandwidth_Summary

	Duration  *float64
	ViewCount *int
}

type Container_Network_ContentDelivery_OriginPull_Mapping struct {
	Entity

	Cname           *string
	Id              *string
	IsSecureContent *bool
	MediaType       *string
	OriginUrl       *string
}

type Container_Network_ContentDelivery_PointsOfPresence struct {
	Entity

	Id   *int
	Name *string
}

type Container_Network_ContentDelivery_PurgeService_Response struct {
	Entity

	StatusCode *string
	Url        *string
}

type Container_Network_ContentDelivery_Report_Usage struct {
	Entity

	ApplicationDeliveryNetwork    *float64
	ApplicationDeliveryNetworkSsl *float64
	DiskSpace                     *float64
	EndDate                       *time.Time
	Flash                         *float64
	Http                          *float64
	HttpSmall                     *float64
	Https                         *float64
	HttpsSmall                    *float64
	Region                        *string
	SslTotal                      *float64
	StandardTotal                 *float64
	StartDate                     *time.Time
	WindowsMedia                  *float64
}

type Container_Network_ContentDelivery_SupportedProtocol struct {
	Entity

	Host      *string
	MediaType *string
	Platform  *string
	Protocol  *string
}

type Container_Network_Directory_Listing struct {
	Entity

	FileCount *int
	Name      *string
	Type      *string
}

type Container_Network_IntrusionProtection_Event struct {
	Entity

	CVEId                 *string
	ActionTaken           *string
	AttackCount           *int
	AttackLongDescription *string
	AttackName            *string
	BeginTime             *string
	BugtraqId             *string
	Classification        *string
	DestinationIpAddress  *string
	DestinationPort       *int
	EndTime               *string
	Platform              *string
	Protocol              *string
	Severity              *string
	SignatureId           *string
	SourceIpAddress       *string
	SourcePort            *int
}

type Container_Network_IntrusionProtection_Statistic struct {
	Entity

	AttackCount *int
	Name        *string
}

type Container_Network_IntrusionProtection_Statistics struct {
	Entity

	Target       *string
	TargetType   *string
	TimeFrame    *string
	TopAttacks   []Container_Network_IntrusionProtection_Statistic
	TotalAttacks *int
}

type Container_Network_IntrusionProtection_SubnetReport struct {
	Entity

	Cidr            *int
	Direction       *string
	Events          []Container_Network_IntrusionProtection_Event
	SubnetIpAddress *string
}

type Container_Network_LoadBalancer_StatusEntry struct {
	Entity

	Content *string
	Label   *string
}

type Container_Network_Media_Information struct {
	Entity

	AudioBitRate     *int
	AudioChannelMode *string
	AudioChannels    *int
	AudioCodec       *string
	AudioSampleRate  *int
	Duration         *float64
	ErrorMessage     *string
	File             *string
	FileFormat       *string
	FileSize         *uint
	FrameRate        *float64
	SizeX            *int
	SizeY            *int
	TotalFrames      *uint
	VideoAspectX     *float64
	VideoAspectY     *int
	VideoCodec       *string
}

type Container_Network_Media_Transcode_Job_Watermark struct {
	Entity

	EndTime                *int
	FileName               *string
	Position               *Container_Network_Media_Transcode_Job_Watermark_Position
	StartTime              *int
	Text                   *string
	TransparencyPercentage *int
}

type Container_Network_Media_Transcode_Job_Watermark_Position struct {
	Entity

	X *int
	Y *int
}

type Container_Network_Media_Transcode_Preset struct {
	Entity

	GUID        *string
	Category    *string
	Description *string
	Name        *string
}

type Container_Network_Media_Transcode_Preset_Element struct {
	Entity

	AdditionalElements  []Container_Network_Media_Transcode_Preset_Element_Option
	DefaultValue        *string
	Description         *string
	Enabled             *bool
	ExtendedDescription *string
	Hidden              *bool
	MaximumValue        *int
	MinimumValue        *int
	Name                *string
	ParentName          *string
	Type                *string
}

type Container_Network_Media_Transcode_Preset_Element_Option struct {
	Entity

	Name  *string
	Value *string
}

type Container_Network_Message_Delivery_Email struct {
	Entity

	Body         *string
	ContainsHtml *bool
	From         *string
	Subject      *string
	To           *string
}

type Container_Network_Message_Delivery_Email_Sendgrid_Account_Overview struct {
	Entity

	CreditsAllowed *int
	CreditsOverage *int
	CreditsRemain  *int
	CreditsUsed    *int
	Package        *string
	Reputation     *int
	Requests       *int
}

type Container_Network_Message_Delivery_Email_Sendgrid_Customer_Profile struct {
	Entity

	Address   *string
	City      *string
	Country   *string
	Email     *string
	FirstName *string
	LastName  *string
	Phone     *string
	State     *string
	Website   *string
	Zip       *string
}

type Container_Network_Message_Delivery_Email_Sendgrid_List_Entry struct {
	Entity

	Created *string
	Email   *string
	Reason  *string
	Status  *string
}

type Container_Network_Message_Delivery_Email_Sendgrid_Statistics struct {
	Entity

	Blocks             *int
	Bounces            *int
	Clicks             *int
	Date               *string
	Delivered          *int
	InvalidEmail       *int
	Opens              *int
	RepeatBounces      *int
	RepeatSpamReports  *int
	RepeatUnsubscribes *int
	Requests           *int
	SpamReports        *int
	UniqueClicks       *int
	UniqueOpens        *int
	Unsubscribes       *int
}

type Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Graph struct {
	Entity

	GraphError *string
	GraphImage *[]byte
	GraphTitle *string
}

type Container_Network_Message_Delivery_Email_Sendgrid_Statistics_Options struct {
	Entity

	AggregatesOnly     *bool
	Category           *string
	Days               *int
	EndDate            *time.Time
	SelectedStatistics []string
	StartDate          *time.Time
}

type Container_Network_Port_Statistic struct {
	Entity

	AdministrativeStatus    *int
	InDiscardPackets        *uint
	InErrorPackets          *uint
	InOctets                *uint
	InUnicastPackets        *uint
	MaximumTransmissionUnit *uint
	OperationalStatus       *int
	OutDiscardPackets       *uint
	OutErrorPackets         *uint
	OutOctets               *uint
	OutUnicastPackets       *uint
	PortDuplex              *uint
	Speed                   *uint
}

type Container_Network_Service_Resource_ObjectStorage_ConnectionInformation struct {
	Entity

	Datacenter          *string
	DatacenterShortName *string
	PrivateEndpoint     *string
	PublicEndpoint      *string
}

type Container_Network_Storage_Backup_Evault_WebCc_Authentication_Details struct {
	Entity

	EventValidation *string
	ViewState       *string
	WebCcUrl        *string
}

type Container_Network_Storage_Evault_Vault_Task struct {
	Entity

	Id           *uint
	Name         *string
	UsedPoolsize *uint
}

type Container_Network_Storage_Evault_WebCc_AgentStatus struct {
	Entity

	LastBackup *time.Time
	Status     *string
}

type Container_Network_Storage_Evault_WebCc_BackupResults struct {
	Entity

	BeginTime *time.Time
	Conflict  *string
	EndTime   *time.Time
	Failed    *string
	Success   *string
}

type Container_Network_Storage_Evault_WebCc_JobDetails struct {
	Entity

	BytesUsed              *uint
	Description            *string
	HardwareId             *int
	LastRunDate            *time.Time
	Name                   *string
	OriginalSize           *uint
	PercentageOfTotalUsage *int
	Result                 *string
	VirtualGuestId         *int
}

type Container_Network_Storage_Host struct {
	Entity

	Id         *int
	ObjectType *string
}

type Container_Network_Storage_Hub_ObjectStorage_ContentDeliveryUrl struct {
	Entity

	Datacenter *string
	FlashUrl   *string
	HttpUrl    *string
}

type Container_Network_Storage_Hub_ObjectStorage_Endpoint struct {
	Entity

	Location *string
	Region   *string
	Type     *string
	Url      *string
}

type Container_Network_Storage_Hub_ObjectStorage_File struct {
	Container_Utility_File_Entity

	Folder *string
	Hash   *string
}

type Container_Network_Storage_Hub_ObjectStorage_Folder struct {
	Entity

	Bytes *uint
	Count *uint
	Name  *string
}

type Container_Network_Storage_Hub_ObjectStorage_Node struct {
	Entity

	DeviceName   *string
	ResourceName *string
	UserAuthUrl  *string
}

type Container_Network_Storage_NetworkConnectionInformation struct {
	Entity

	Id          *string
	IpAddress   *string
	StorageType *string
}

type Container_Network_Subnet_IpAddress struct {
	Entity

	Hardware           *Hardware
	IpAddress          *string
	IsBroadcastAddress *bool
	IsGatewayAddress   *bool
	IsNetworkAddress   *bool
}

type Container_Network_Subnet_Registration_SubnetReference struct {
	Entity

	RegistrationId *int
	SubnetCidr     *string
}

type Container_Network_Subnet_Registration_TransactionDetails struct {
	Entity

	SubnetReferences []Container_Network_Subnet_Registration_SubnetReference
	TransactionId    *int
}

type Container_Notification_Mass_Filter_TemplateKey struct {
	Entity
}

type Container_Notification_Mass_Filter_TemplateValue struct {
	Entity
}

type Container_Policy_Acceptance struct {
	Entity

	AcceptedFlag              *bool
	PolicyName                *string
	ProductPolicyAssignmentId *int
}

type Container_Product_Item_Category struct {
	Entity

	Id *int
}

type Container_Product_Item_Category_Question_Answer struct {
	Entity

	Answer       *string
	CategoryCode *string
	CategoryId   *int
	QuestionId   *int
}

type Container_Product_Item_Category_ZeroFee_Count struct {
	Entity

	CategoryCode *string
	CategoryId   *int
	CategoryName *string
	Count        *int
}

type Container_Product_Item_Discount_Program struct {
	Entity

	ApplicableQuantity      *int
	Item                    *Product_Item
	OneTimeAmount           *float64
	OneTimeTax              *float64
	Prices                  []Product_Item_Price
	ProratedOneTimeAmount   *float64
	ProratedOneTimeTax      *float64
	ProratedRecurringAmount *float64
	ProratedRecurringTax    *float64
	RecurringAmount         *float64
	RecurringTax            *float64
}

type Container_Product_Order struct {
	Entity

	BigDataOrderFlag              *bool
	BillingInformation            *Container_Product_Order_Billing_Information
	BillingOrderItemId            *int
	CancelUrl                     *string
	ContainerIdentifier           *string
	ContainerSplHash              *string
	CurrencyShortName             *string
	DeviceFingerprintId           *string
	DisplayLayerSessionId         *string
	ExtendedHardwareTesting       *bool
	FlexibleCreditProgramPrice    *Product_Item_Price
	Hardware                      []Hardware
	ImageTemplateGlobalIdentifier *string
	ImageTemplateId               *int
	IsManagedOrder                *int
	ItemCategoryQuestionAnswers   []Container_Product_Item_Category_Question_Answer
	Location                      *string
	LocationObject                *Location
	Message                       *string
	OrderContainers               []Container_Product_Order
	OrderHostnames                []string
	OrderVerificationExceptions   []Container_Exception
	PackageId                     *int
	PaymentType                   *string
	PostTaxRecurring              *float64
	PostTaxRecurringHourly        *float64
	PostTaxRecurringMonthly       *float64
	PostTaxSetup                  *float64
	PreTaxRecurring               *float64
	PreTaxRecurringHourly         *float64
	PreTaxRecurringMonthly        *float64
	PreTaxSetup                   *float64
	PresaleEvent                  *Sales_Presale_Event
	PresetId                      *int
	Prices                        []Product_Item_Price
	PrimaryDiskPartitionId        *int
	Priorities                    []string
	PrivateCloudOrderFlag         *bool
	PrivateCloudOrderType         *string
	PromotionCode                 *string
	Properties                    []Container_Product_Order_Property
	ProratedInitialCharge         *float64
	ProratedOrderTotal            *float64
	ProvisionScripts              []string
	Quantity                      *int
	QuoteName                     *string
	RegionalGroup                 *string
	ResourceGroupId               *int
	ResourceGroupName             *string
	ResourceGroupTemplateId       *int
	ReturnUrl                     *string
	SendQuoteEmailFlag            *bool
	ServerCoreCount               *int
	SourceVirtualGuestId          *int
	SshKeys                       []Container_Product_Order_SshKeys
	StepId                        *int
	StorageGroups                 []Container_Product_Order_Storage_Group
	TaxCacheHash                  *string
	TaxCompletedFlag              *bool
	TechIncubatorItemPrice        *Product_Item_Price
	TotalRecurringTax             *float64
	TotalSetupTax                 *float64
	UseHourlyPricing              *bool
	VirtualGuests                 []Virtual_Guest
}

type Container_Product_Order_Account_Media_Data_Transfer_Request struct {
	Container_Product_Order

	Request *Account_Media_Data_Transfer_Request
}

type Container_Product_Order_Attribute_Address struct {
	Entity

	AddressLine1 *string
	AddressLine2 *string
	City         *string
	CountryCode  *string
	NonUsState   *string
	PostalCode   *string
	State        *string
}

type Container_Product_Order_Attribute_Contact struct {
	Entity

	Address          *Container_Product_Order_Attribute_Address
	EmailAddress     *string
	FaxNumber        *string
	FirstName        *string
	LastName         *string
	OrganizationName *string
	PhoneNumber      *string
	Title            *string
}

type Container_Product_Order_Attribute_Organization struct {
	Entity

	Address          *Container_Product_Order_Attribute_Address
	FaxNumber        *string
	OrganizationName *string
	PhoneNumber      *string
}

type Container_Product_Order_Billing_Information struct {
	Entity

	BillingAddressLine1          *string
	BillingAddressLine2          *string
	BillingCity                  *string
	BillingCountryCode           *string
	BillingEmail                 *string
	BillingNameCompany           *string
	BillingNameFirst             *string
	BillingNameLast              *string
	BillingPhoneFax              *string
	BillingPhoneVoice            *string
	BillingPostalCode            *string
	BillingState                 *string
	CardAccountNumber            *string
	CardExpirationMonth          *int
	CardExpirationYear           *int
	CreditCardVerificationNumber *string
	TaxExempt                    *int
	VatId                        *string
}

type Container_Product_Order_Dns_Domain_Registration struct {
	Container_Product_Order

	AdministrativeContact  *Container_Dns_Domain_Registration_Contact
	BillingContact         *Container_Dns_Domain_Registration_Contact
	DomainRegistrationList []Container_Dns_Domain_Registration_List
	OwnerContact           *Container_Dns_Domain_Registration_Contact
	RegistrationType       *string
	TechnicalContact       *Container_Dns_Domain_Registration_Contact
}

type Container_Product_Order_Dns_Domain_Reseller struct {
	Container_Product_Order

	CreditAmount *float64
}

type Container_Product_Order_Gateway_Appliance_Cluster struct {
	Container_Product_Order

	ClusterIdentifier *string
	ClusterOrderType  *string
}

type Container_Product_Order_Hardware_Security_Module struct {
	Container_Product_Order_Hardware_Server
}

type Container_Product_Order_Hardware_Server struct {
	Container_Product_Order

	ClusterIdentifier                           *string
	ClusterOrderType                            *string
	ClusterResourceId                           *int
	MonitoringAgentConfigurationTemplateGroupId *int
	PrivateCloudServerRole                      *string
	RequiredUpstreamDeviceId                    *int
	Tags                                        []Container_Product_Order_Property
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

	ConfigurationTemplateGroups []Monitoring_Agent_Configuration_Template_Group
	ServerType                  *string
}

type Container_Product_Order_MultiConfiguration struct {
	Container_Product_Order
}

type Container_Product_Order_MultiConfiguration_Tornado struct {
	Container_Product_Order_MultiConfiguration
}

type Container_Product_Order_Network struct {
	Entity

	Network     *Network
	PublicVlans []Container_Product_Order
	Subnets     []Container_Product_Order
}

type Container_Product_Order_Network_Application_Delivery_Controller struct {
	Container_Product_Order

	ApplicationDeliveryControllerId *int
}

type Container_Product_Order_Network_ContentDelivery_Account struct {
	Container_Product_Order

	CdnAccountName *string
}

type Container_Product_Order_Network_ContentDelivery_Account_Upgrade struct {
	Container_Product_Order

	CdnAccountId *string
}

type Container_Product_Order_Network_LoadBalancer struct {
	Container_Product_Order
}

type Container_Product_Order_Network_LoadBalancer_Global struct {
	Container_Product_Order

	Domain   *string
	Hostname *string
}

type Container_Product_Order_Network_Message_Delivery struct {
	Container_Product_Order

	AccountPassword *string
	AccountUsername *string
	EmailAddress    *string
}

type Container_Product_Order_Network_Message_Queue struct {
	Container_Product_Order
}

type Container_Product_Order_Network_PerformanceStorage struct {
	Container_Product_Order
}

type Container_Product_Order_Network_PerformanceStorage_Iscsi struct {
	Container_Product_Order_Network_PerformanceStorage

	OsFormatType *Network_Storage_Iscsi_OS_Type
}

type Container_Product_Order_Network_PerformanceStorage_Nfs struct {
	Container_Product_Order_Network_PerformanceStorage
}

type Container_Product_Order_Network_Protection_Firewall struct {
	Container_Product_Order
}

type Container_Product_Order_Network_Protection_Firewall_Dedicated struct {
	Container_Product_Order

	Vlan   *Network_Vlan
	VlanId *int
}

type Container_Product_Order_Network_Storage_Backup_Evault_Plugin struct {
	Container_Product_Order
}

type Container_Product_Order_Network_Storage_Backup_Evault_Vault struct {
	Container_Product_Order
}

type Container_Product_Order_Network_Storage_Enterprise struct {
	Container_Product_Order

	OriginVolumeId         *int
	OriginVolumeScheduleId *int
	OsFormatType           *Network_Storage_Iscsi_OS_Type
}

type Container_Product_Order_Network_Storage_Enterprise_SnapshotSpace struct {
	Container_Product_Order

	VolumeId *int
}

type Container_Product_Order_Network_Storage_Enterprise_SnapshotSpace_Upgrade struct {
	Container_Product_Order_Network_Storage_Enterprise_SnapshotSpace
}

type Container_Product_Order_Network_Storage_Hub struct {
	Container_Product_Order
}

type Container_Product_Order_Network_Storage_Hub_Datacenter struct {
	Entity

	Location        *Location
	UsageRatePrices []Product_Item_Price
}

type Container_Product_Order_Network_Storage_Iscsi struct {
	Container_Product_Order
}

type Container_Product_Order_Network_Storage_Iscsi_Replication struct {
	Container_Product_Order

	VolumeId *int
}

type Container_Product_Order_Network_Storage_Iscsi_SnapshotSpace struct {
	Container_Product_Order

	VolumeId *int
}

type Container_Product_Order_Network_Storage_Modification struct {
	Container_Product_Order

	VolumeId *int
}

type Container_Product_Order_Network_Storage_Nas struct {
	Container_Product_Order
}

type Container_Product_Order_Network_Storage_Object struct {
	Container_Product_Order
}

type Container_Product_Order_Network_Subnet struct {
	Container_Product_Order

	Description         *string
	EndPointIpAddressId *int
	EndPointVlanId      *int
	Id                  *int
	RouterHostname      *string
}

type Container_Product_Order_Network_Tunnel_Ipsec struct {
	Container_Product_Order
}

type Container_Product_Order_Network_Vlan struct {
	Container_Product_Order

	Description        *string
	HostnameDatacenter *string
	HostnameRouter     *string
	Id                 *int
	Name               *string
	Router             *Hardware
	RouterId           *int
	Subnets            []Container_Product_Order
	VlanNumber         *int
}

type Container_Product_Order_Network_Vlans struct {
	Entity

	PrivateVlans []Container_Product_Order
	PublicVlans  []Container_Product_Order
}

type Container_Product_Order_Property struct {
	Entity

	Name  *string
	Value *string
}

type Container_Product_Order_Receipt struct {
	Entity

	ExternalPaymentCheckoutUrl *string
	ExternalPaymentToken       *string
	OrderDate                  *time.Time
	OrderDetails               *Container_Product_Order
	OrderId                    *int
	PaypalCheckoutUrl          *string
	PaypalToken                *string
	PlacedOrder                *Billing_Order
	Quote                      *Billing_Order_Quote
}

type Container_Product_Order_Security_Certificate struct {
	Container_Product_Order

	AdministrativeContact     *Container_Product_Order_Attribute_Contact
	BillingContact            *Container_Product_Order_Attribute_Contact
	CertificateSigningRequest *string
	OrderApproverEmailAddress *string
	OrganizationInformation   *Container_Product_Order_Attribute_Organization
	RenewalFlag               *bool
	ServerCount               *int
	ServerType                *string
	TechnicalContact          *Container_Product_Order_Attribute_Contact
	ValidityMonths            *int
}

type Container_Product_Order_Software_Component_Virtual struct {
	Container_Product_Order

	EndPointIpAddressIds []int
}

type Container_Product_Order_Software_License struct {
	Container_Product_Order
}

type Container_Product_Order_SshKeys struct {
	Entity

	SshKeyIds []int
}

type Container_Product_Order_Storage_Group struct {
	Entity

	ArraySize           *float64
	ArrayTypeId         *int
	HardDrives          []int
	HotSpareDrives      []int
	PartitionTemplateId *int
	Partitions          []Container_Product_Order_Storage_Group_Partition
}

type Container_Product_Order_Storage_Group_Partition struct {
	Entity

	IsGrow *bool
	Name   *string
	Size   *float64
}

type Container_Product_Order_User_Customer_External_Binding struct {
	Container_Product_Order

	ExternalId *string
	UserId     *int
	VendorId   *int
}

type Container_Product_Order_Virtual_Disk_Image struct {
	Container_Product_Order

	DiskDescription *string
}

type Container_Product_Order_Virtual_Guest struct {
	Container_Product_Order_Hardware_Server

	BootableDiskId *int
}

type Container_Product_Order_Virtual_Guest_Upgrade struct {
	Container_Product_Order_Virtual_Guest
}

type Container_Provisioning_Maintenance_Window struct {
	Entity

	ClassificationIds     []Provisioning_Maintenance_Classification
	ItemCategoryIds       []Product_Item_Category
	MaintenanceWindowId   *int
	TicketId              *int
	WindowMaintenanceDate *time.Time
}

type Container_Referral_Partner_Commission struct {
	Entity

	CommissionAmount         *float64
	CommissionRate           *float64
	CreateDate               *time.Time
	ReferralAccountId        *int
	ReferralCompanyName      *string
	ReferralPartnerAccountId *int
	ReferralRevenue          *float64
}

type Container_Referral_Partner_Payment_Option struct {
	Entity

	AccountNumber     *string
	AccountType       *string
	Address1          *string
	Address2          *string
	BankTransitNumber *string
	City              *string
	CompanyName       *string
	Country           *string
	FederalTaxId      *string
	FirstName         *string
	LastName          *string
	PaymentType       *string
	PaypalEmail       *string
	PhoneNumber       *string
	PostalCode        *string
	State             *string
}

type Container_Referral_Partner_Prospect struct {
	Entity

	Address1    *string
	Address2    *string
	City        *string
	CompanyName *string
	Country     *string
	Email       *string
	FirstName   *string
	LastName    *string
	OfficePhone *string
	PostalCode  *string
	Questions   []string
	Responses   []Survey_Response
	State       *string
	SurveyId    *string
}

type Container_RemoteManagement_Graphs_SensorSpeed struct {
	Entity

	Graph *[]byte
	Title *string
}

type Container_RemoteManagement_Graphs_SensorTemperature struct {
	Entity

	Graph *[]byte
	Title *string
}

type Container_RemoteManagement_PmInfo struct {
	Entity

	PmInfoId      *string
	PmInfoReading *string
}

type Container_RemoteManagement_SensorReading struct {
	Entity

	LowerCritical       *string
	LowerNonCritical    *string
	LowerNonRecoverable *string
	SensorId            *string
	SensorReading       *string
	SensorUnits         *string
	Status              *string
	UpperCritical       *string
	UpperNonCritical    *string
	UpperNonRecoverable *string
}

type Container_RemoteManagement_SensorReadingsWithGraphs struct {
	Entity

	RawData           []Container_RemoteManagement_SensorReading
	SpeedGraphs       []Container_RemoteManagement_Graphs_SensorSpeed
	TemperatureGraphs []Container_RemoteManagement_Graphs_SensorTemperature
}

type Container_Resource_Metadata_ServiceResource struct {
	Entity

	BackendIpAddress *string
	Type             *Network_Service_Resource_Type
}

type Container_Search_ObjectType struct {
	Entity

	Name       *string
	Properties []Container_Search_ObjectType_Property
}

type Container_Search_ObjectType_Property struct {
	Entity

	Name         *string
	SortableFlag *bool
	Type         *string
}

type Container_Search_Result struct {
	Entity

	MatchedTerms   []string
	RelevanceScore *float64
	Resource       *Entity
	ResourceType   *string
}

type Container_Software_Component_HostIps_Policy struct {
	Entity

	Policy      *string
	PolicyTitle *string
}

type Container_Tax_Cache struct {
	Entity

	EffectiveTaxRate *float64
	Items            []Container_Tax_Cache_Item
	Status           *string
	TotalTaxAmount   *float64
}

type Container_Tax_Cache_Item struct {
	Entity

	CategoryCode  *string
	ContainerHash *string
	ItemPriceId   *int
	TaxRates      *Container_Tax_Rates
}

type Container_Tax_Rates struct {
	Entity

	LaborTaxRate     *float64
	LocationId       *float64
	OneTimeTaxRate   *float64
	RecurringTaxRate *float64
	SetupTaxRate     *float64
}

type Container_Ticket_GraphInputs struct {
	Entity

	EndDate            *time.Time
	NetworkInterfaceId *int
	Pod                *int
	ServerName         *string
	StartDate          *time.Time
}

type Container_Ticket_GraphOutputs struct {
	Entity

	GraphImage   *[]byte
	GraphTitle   *string
	MaxEndDate   *time.Time
	MinStartDate *time.Time
}

type Container_Ticket_Priority struct {
	Entity

	Name  *string
	Value *int
}

type Container_Ticket_Survey_Preference struct {
	Entity

	Applicable   *bool
	OptedOut     *bool
	OptedOutDate *time.Time
}

type Container_User_Authentication_Token struct {
	Entity

	Hash   *string
	User   *User_Customer
	UserId *int
}

type Container_User_Customer_External_Binding struct {
	Entity

	AuthenticationToken      *string
	OpenIdConnectAccessToken *string
	OpenIdConnectAccountId   *int
	OpenIdConnectProvider    *int
	Password                 *string
	SecurityQuestionAnswer   *string
	SecurityQuestionId       *int
	Username                 *string
	Vendor                   *string
}

type Container_User_Customer_External_Binding_Phone struct {
	Container_User_Customer_External_Binding
}

type Container_User_Customer_External_Binding_Phone_Mode struct {
	Entity

	Mode    *string
	Pin     *string
	PinMode *string
}

type Container_User_Customer_External_Binding_Totp struct {
	Container_User_Customer_External_Binding

	SecurityCode *string
}

type Container_User_Customer_External_Binding_Vendor struct {
	Entity

	KeyName *string
	Name    *string
}

type Container_User_Customer_External_Binding_Verisign struct {
	Container_User_Customer_External_Binding

	SecondSecurityCode *string
	SecurityCode       *string
}

type Container_User_Customer_PasswordSet struct {
	Entity

	AnsweredSecurityQuestionId *int
	AuthenticationMethods      []int
	Key                        *string
	Password                   *string
	SecurityAnswer             *string
	SecurityQuestions          []User_Security_Question
	UserId                     *int
}

type Container_User_Customer_Portal_MobileToken struct {
	Container_User_Customer_Portal_Token

	HasExternalBinding *bool
}

type Container_User_Customer_Portal_Token struct {
	Entity

	Hash   *string
	User   *User_Customer
	UserId *int
}

type Container_User_Data_Phone struct {
	Entity

	CountryCode *int
	Extension   *string
	Phone       *string
	PhoneType   *string
}

type Container_User_Employee_External_Binding_Verisign struct {
	Entity
}

type Container_Utility_File_Attachment struct {
	Entity

	Data     *[]byte
	Filename *string
}

type Container_Utility_File_Descriptor struct {
	Entity

	FileName     *string
	FriendlyName *string
	ModifyDate   *time.Time
}

type Container_Utility_File_Entity struct {
	Entity

	Content     *[]byte
	ContentType *string
	CreateDate  *time.Time
	DeleteDate  *time.Time
	Id          *string
	IsShared    *int
	ModifyDate  *time.Time
	Name        *string
	Owner       *string
	Size        *uint
	Type        *string
	Version     *int
}

type Container_Utility_Message struct {
	Entity

	CreateDate *time.Time
	Id         *int
	Message    *string
	ModifyDate *time.Time
	Summary    *string
}

type Container_Utility_Microsoft_Windows_UpdateServices_Status struct {
	Entity

	LastRebootDate   *time.Time
	LastStatusDate   *time.Time
	LastSyncDate     *time.Time
	PrivateIPAddress *string
	SyncStatus       *string
	UpdateStatus     *string
}

type Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem struct {
	Entity

	Description     *string
	Failed          *bool
	KbArticleNumber *int
	Optional        *bool
	RequiresReboot  *bool
}

type Container_Utility_Network_Firewall_Rule_Attribute struct {
	Entity

	Actions             []string
	MaximumRuleCount    *int
	Protocols           []string
	SourceIpSubnetMasks []Container_Utility_Network_Subnet_Mask_Generic_Detail
}

type Container_Utility_Network_Subnet_Mask_Generic_Detail struct {
	Entity

	Cidr        *string
	Description *string
	Mask        *string
}

type Container_Virtual_Guest_Block_Device_Template_Configuration struct {
	Entity

	Name                         *string
	Note                         *string
	OperatingSystemReferenceCode *string
	Uri                          *string
}

type Container_Virtual_Guest_Configuration struct {
	Entity

	BlockDevices      []Container_Virtual_Guest_Configuration_Option
	Datacenters       []Container_Virtual_Guest_Configuration_Option
	Memory            []Container_Virtual_Guest_Configuration_Option
	NetworkComponents []Container_Virtual_Guest_Configuration_Option
	OperatingSystems  []Container_Virtual_Guest_Configuration_Option
	Processors        []Container_Virtual_Guest_Configuration_Option
}

type Container_Virtual_Guest_Configuration_Option struct {
	Entity

	ItemPrice *Product_Item_Price
	Template  *Virtual_Guest
}

type Dns_Domain struct {
	Entity

	Account             *Account
	Id                  *int
	ManagedResourceFlag *bool
	Name                *string
	ResourceRecordCount *uint
	ResourceRecords     []Dns_Domain_ResourceRecord
	Secondary           *Dns_Secondary
	Serial              *int
	UpdateDate          *time.Time
}

type Dns_Domain_Forward struct {
	Dns_Domain
}

type Dns_Domain_Registration struct {
	Entity

	Account                        *Account
	CreateDate                     *time.Time
	DomainRegistrationStatus       *Dns_Domain_Registration_Status
	DomainRegistrationStatusId     *int
	ExpireDate                     *time.Time
	Id                             *int
	LockedFlag                     *int
	ModifyDate                     *time.Time
	Name                           *string
	RegistrantVerificationStatus   *Dns_Domain_Registration_Registrant_Verification_Status
	RegistrantVerificationStatusId *int
	ServiceProvider                *Service_Provider
	ServiceProviderId              *int
}

type Dns_Domain_Registration_Registrant_Verification_Status struct {
	Entity

	Description *string
	Id          *int
	KeyName     *string
	Name        *string
}

type Dns_Domain_Registration_Status struct {
	Entity

	Description *string
	Id          *int
	KeyName     *string
	Name        *string
}

type Dns_Domain_ResourceRecord struct {
	Entity

	Data              *string
	Domain            *Dns_Domain
	DomainId          *int
	Expire            *int
	Host              *string
	Id                *int
	Minimum           *int
	MxPriority        *int
	Refresh           *int
	ResponsiblePerson *string
	Retry             *int
	Ttl               *int
	Type              *string
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

	IsGatewayAddress *bool
}

type Dns_Domain_ResourceRecord_SoaType struct {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_ResourceRecord_SpfType struct {
	Dns_Domain_ResourceRecord_TxtType
}

type Dns_Domain_ResourceRecord_SrvType struct {
	Dns_Domain_ResourceRecord

	Port     *int
	Priority *int
	Protocol *string
	Service  *string
	Weight   *int
}

type Dns_Domain_ResourceRecord_TxtType struct {
	Dns_Domain_ResourceRecord
}

type Dns_Domain_Reverse struct {
	Dns_Domain

	NetworkAddress *string
}

type Dns_Domain_Reverse_Version4 struct {
	Dns_Domain_Reverse
}

type Dns_Domain_Reverse_Version6 struct {
	Dns_Domain_Reverse
}

type Dns_Message struct {
	Entity

	CreateDate     *time.Time
	Domain         *Dns_Domain
	Id             *int
	Message        *string
	Priority       *string
	ResourceRecord *Dns_Domain_ResourceRecord
	Secondary      *Dns_Secondary
}

type Dns_Secondary struct {
	Entity

	Account           *Account
	CreateDate        *time.Time
	Domain            *Dns_Domain
	ErrorMessageCount *uint
	ErrorMessages     []Dns_Message
	Id                *int
	LastUpdate        *time.Time
	MasterIpAddress   *string
	Status            *Dns_Status
	StatusId          *int
	StatusText        *string
	TransferFrequency *int
	ZoneName          *string
}

type Dns_Status struct {
	Entity

	Id   *int
	Name *string
}

type Entity struct {
}

type Event_Log struct {
	Entity

	AccountId       *int
	EventCreateDate *time.Time
	EventName       *string
	IpAddress       *string
	Label           *string
	MetaData        *string
	ObjectId        *int
	ObjectName      *string
	Resource        *Entity
	TraceId         *string
	User            *User_Customer
	UserId          *int
	UserType        *string
	Username        *string
}

type FlexibleCredit_Affiliate struct {
	Entity

	FlexibleCreditProgram *FlexibleCredit_Program
	Id                    *int
	Name                  *string
}

type FlexibleCredit_Company_Type struct {
	Entity

	Description *string
	Id          *int
}

type FlexibleCredit_Enrollment struct {
	Entity

	Account                  *Account
	AccountId                *int
	Affiliate                *FlexibleCredit_Affiliate
	AffiliateId              *int
	AgreementCompleteFlag    *int
	CompanyDescription       *string
	CompanyType              *FlexibleCredit_Company_Type
	CompanyTypeId            *int
	EnrollmentDate           *time.Time
	FlexibleCreditProgram    *FlexibleCredit_Program
	GraduationDate           *time.Time
	IsActiveFlag             *bool
	MonthlyCreditAmount      *float64
	Representative           *User_Employee
	RepresentativeEmployeeId *int
}

type FlexibleCredit_Program struct {
	Entity

	Id      *int
	KeyName *string
	Name    *string
}

type Hardware struct {
	Entity

	Account                                     *Account
	AccountId                                   *int
	ActiveComponentCount                        *uint
	ActiveComponents                            []Hardware_Component
	ActiveNetworkMonitorIncident                []Network_Monitor_Version1_Incident
	ActiveNetworkMonitorIncidentCount           *uint
	AllPowerComponentCount                      *uint
	AllPowerComponents                          []Hardware_Power_Component
	AllowedHost                                 *Network_Storage_Allowed_Host
	AllowedNetworkStorage                       []Network_Storage
	AllowedNetworkStorageCount                  *uint
	AllowedNetworkStorageReplicaCount           *uint
	AllowedNetworkStorageReplicas               []Network_Storage
	AntivirusSpywareSoftwareComponent           *Software_Component
	AttributeCount                              *uint
	Attributes                                  []Hardware_Attribute
	AverageDailyPublicBandwidthUsage            *float64
	BackendNetworkComponentCount                *uint
	BackendNetworkComponents                    []Network_Component
	BackendRouterCount                          *uint
	BackendRouters                              []Hardware
	BandwidthAllocation                         *float64
	BandwidthAllotmentDetail                    *Network_Bandwidth_Version1_Allotment_Detail
	BareMetalInstanceFlag                       *int
	BenchmarkCertificationCount                 *uint
	BenchmarkCertifications                     []Hardware_Benchmark_Certification
	BillingItem                                 *Billing_Item_Hardware
	BillingItemFlag                             *bool
	BlockCancelBecauseDisconnectedFlag          *bool
	BusinessContinuanceInsuranceFlag            *bool
	ComponentCount                              *uint
	Components                                  []Hardware_Component
	ContinuousDataProtectionSoftwareComponent   *Software_Component
	CurrentBillableBandwidthUsage               *float64
	Datacenter                                  *Location
	DatacenterName                              *string
	Domain                                      *string
	DownlinkHardware                            []Hardware
	DownlinkHardwareCount                       *uint
	DownlinkNetworkHardware                     []Hardware
	DownlinkNetworkHardwareCount                *uint
	DownlinkServerCount                         *uint
	DownlinkServers                             []Hardware
	DownlinkVirtualGuestCount                   *uint
	DownlinkVirtualGuests                       []Virtual_Guest
	DownstreamHardwareBindingCount              *uint
	DownstreamHardwareBindings                  []Network_Component_Uplink_Hardware
	DownstreamNetworkHardware                   []Hardware
	DownstreamNetworkHardwareCount              *uint
	DownstreamNetworkHardwareWithIncidentCount  *uint
	DownstreamNetworkHardwareWithIncidents      []Hardware
	DownstreamServerCount                       *uint
	DownstreamServers                           []Hardware
	DownstreamVirtualGuestCount                 *uint
	DownstreamVirtualGuests                     []Virtual_Guest
	DriveControllerCount                        *uint
	DriveControllers                            []Hardware_Component
	EvaultNetworkStorage                        []Network_Storage
	EvaultNetworkStorageCount                   *uint
	FirewallServiceComponent                    *Network_Component_Firewall
	FixedConfigurationPreset                    *Product_Package_Preset
	FrontendNetworkComponentCount               *uint
	FrontendNetworkComponents                   []Network_Component
	FrontendRouterCount                         *uint
	FrontendRouters                             []Hardware
	FullyQualifiedDomainName                    *string
	GlobalIdentifier                            *string
	HardDriveCount                              *uint
	HardDrives                                  []Hardware_Component
	HardwareChassis                             *Hardware_Chassis
	HardwareFunction                            *Hardware_Function
	HardwareFunctionDescription                 *string
	HardwareStatus                              *Hardware_Status
	HardwareStatusId                            *int
	HasTrustedPlatformModuleBillingItemFlag     *bool
	HostIpsSoftwareComponent                    *Software_Component
	Hostname                                    *string
	HourlyBillingFlag                           *bool
	Id                                          *int
	InboundBandwidthUsage                       *float64
	InboundPublicBandwidthUsage                 *float64
	LastTransaction                             *Provisioning_Version1_Transaction
	LatestNetworkMonitorIncident                *Network_Monitor_Version1_Incident
	Location                                    *Location
	LocationPathString                          *string
	LockboxNetworkStorage                       *Network_Storage
	ManagedResourceFlag                         *bool
	ManufacturerSerialNumber                    *string
	Memory                                      []Hardware_Component
	MemoryCapacity                              *uint
	MemoryCount                                 *uint
	MetricTrackingObject                        *Metric_Tracking_Object_HardwareServer
	MonitoringAgentCount                        *uint
	MonitoringAgents                            []Monitoring_Agent
	MonitoringRobot                             *Monitoring_Robot
	MonitoringServiceComponent                  *Network_Monitor_Version1_Query_Host_Stratum
	MonitoringServiceEligibilityFlag            *bool
	MonitoringServiceFlag                       *bool
	Motherboard                                 *Hardware_Component
	NetworkCardCount                            *uint
	NetworkCards                                []Hardware_Component
	NetworkComponentCount                       *uint
	NetworkComponents                           []Network_Component
	NetworkGatewayMember                        *Network_Gateway_Member
	NetworkGatewayMemberFlag                    *bool
	NetworkManagementIpAddress                  *string
	NetworkMonitorAttachedDownHardware          []Hardware
	NetworkMonitorAttachedDownHardwareCount     *uint
	NetworkMonitorAttachedDownVirtualGuestCount *uint
	NetworkMonitorAttachedDownVirtualGuests     []Virtual_Guest
	NetworkMonitorCount                         *uint
	NetworkMonitorIncidentCount                 *uint
	NetworkMonitorIncidents                     []Network_Monitor_Version1_Incident
	NetworkMonitors                             []Network_Monitor_Version1_Query_Host
	NetworkStatus                               *string
	NetworkStatusAttribute                      *Hardware_Attribute
	NetworkStorage                              []Network_Storage
	NetworkStorageCount                         *uint
	NetworkVlanCount                            *uint
	NetworkVlans                                []Network_Vlan
	NextBillingCycleBandwidthAllocation         *float64
	Notes                                       *string
	NotesHistory                                []Hardware_Note
	NotesHistoryCount                           *uint
	OperatingSystem                             *Software_Component_OperatingSystem
	OperatingSystemReferenceCode                *string
	OutboundBandwidthUsage                      *float64
	OutboundPublicBandwidthUsage                *float64
	PointOfPresenceLocation                     *Location
	PostInstallScriptUri                        *string
	PowerComponentCount                         *uint
	PowerComponents                             []Hardware_Power_Component
	PowerSupply                                 []Hardware_Component
	PowerSupplyCount                            *uint
	PrimaryBackendIpAddress                     *string
	PrimaryBackendNetworkComponent              *Network_Component
	PrimaryIpAddress                            *string
	PrimaryNetworkComponent                     *Network_Component
	PrivateNetworkOnlyFlag                      *bool
	ProcessorCoreAmount                         *uint
	ProcessorCount                              *uint
	ProcessorPhysicalCoreAmount                 *uint
	Processors                                  []Hardware_Component
	ProvisionDate                               *time.Time
	Rack                                        *Location
	RaidControllerCount                         *uint
	RaidControllers                             []Hardware_Component
	RecentEventCount                            *uint
	RecentEvents                                []Notification_Occurrence_Event
	RemoteManagementAccountCount                *uint
	RemoteManagementAccounts                    []Hardware_Component_RemoteManagement_User
	RemoteManagementComponent                   *Network_Component
	ResourceGroupCount                          *uint
	ResourceGroupMemberReferenceCount           *uint
	ResourceGroupMemberReferences               []Resource_Group_Member
	ResourceGroupRoleCount                      *uint
	ResourceGroupRoles                          []Resource_Group_Role
	ResourceGroups                              []Resource_Group
	RouterCount                                 *uint
	Routers                                     []Hardware
	ScaleAssetCount                             *uint
	ScaleAssets                                 []Scale_Asset
	SecurityScanRequestCount                    *uint
	SecurityScanRequests                        []Network_Security_Scanner_Request
	SerialNumber                                *string
	ServerRoom                                  *Location
	ServiceProvider                             *Service_Provider
	ServiceProviderId                           *int
	ServiceProviderResourceId                   *int
	SoftwareComponentCount                      *uint
	SoftwareComponents                          []Software_Component
	SparePoolBillingItem                        *Billing_Item_Hardware
	SshKeyCount                                 *uint
	SshKeys                                     []Security_Ssh_Key
	StorageNetworkComponentCount                *uint
	StorageNetworkComponents                    []Network_Component
	TagReferenceCount                           *uint
	TagReferences                               []Tag_Reference
	TopLevelLocation                            *Location
	UpgradeRequest                              *Product_Upgrade_Request
	UplinkHardware                              *Hardware
	UplinkNetworkComponentCount                 *uint
	UplinkNetworkComponents                     []Network_Component
	UserData                                    []Hardware_Attribute
	UserDataCount                               *uint
	VirtualChassis                              *Hardware_Group
	VirtualChassisSiblingCount                  *uint
	VirtualChassisSiblings                      []Hardware
	VirtualHost                                 *Virtual_Host
	VirtualLicenseCount                         *uint
	VirtualLicenses                             []Software_VirtualLicense
	VirtualRack                                 *Network_Bandwidth_Version1_Allotment
	VirtualRackId                               *int
	VirtualRackName                             *string
	VirtualizationPlatform                      *Software_Component
}

type Hardware_Attribute struct {
	Entity

	HardwareAttributeType   *Hardware_Attribute_Type
	HardwareAttributeTypeId *int
	Id                      *int
	Value                   *string
}

type Hardware_Attribute_Type struct {
	Entity

	Keyname *string
	Name    *string
}

type Hardware_Benchmark_Certification struct {
	Entity

	Account    *Account
	AccountId  *int
	CreateDate *time.Time
	Hardware   *Hardware
	HardwareId *int
}

type Hardware_Chassis struct {
	Entity

	BackplaneCapacity       *string
	BayCapacity             *string
	DriveCapacity           *string
	DriveControllerCapacity *string
	FormFactorId            *int
	GpuCapacity             *string
	HardwareFunction        *Hardware_Function
	Id                      *int
	Manufacturer            *string
	Name                    *string
	PowerCapacity           *string
	UnitSize                *int
	Version                 *string
}

type Hardware_Component struct {
	Entity

	Capacity                       *float64
	Children                       []Hardware_Component
	ChildrenCount                  *uint
	DownlinkHardwareComponentCount *uint
	DownlinkHardwareComponents     []Hardware_Component
	Hardware                       *Hardware
	HardwareComponentModel         *Hardware_Component_Model
	HardwareComponentModelId       *int
	HardwareComponentType          *Hardware_Component_Type
	HardwareId                     *int
	Id                             *int
	ModifyDate                     *time.Time
	Name                           *string
	NetworkComponentCount          *uint
	NetworkComponents              []Network_Component
	Owner                          *Account
	Parent                         *Hardware_Component
	RaidMode                       *string
	SerialNumber                   *string
	ServiceProvider                *Service_Provider
	ServiceProviderId              *int
	UplinkHardwareComponentCount   *uint
	UplinkHardwareComponents       []Hardware_Component
}

type Hardware_Component_Attribute struct {
	Entity

	HardwareComponent                *Hardware_Component
	HardwareComponentAttributeType   *Hardware_Component_Attribute_Type
	HardwareComponentAttributeTypeId *int
	HardwareComponentId              *int
	Value                            *string
}

type Hardware_Component_Attribute_Type struct {
	Entity

	Description *string
	Id          *int
	KeyName     *string
	Name        *string
}

type Hardware_Component_DriveController struct {
	Hardware_Component
}

type Hardware_Component_HardDrive struct {
	Hardware_Component

	PartitionCount *uint
	Partitions     []Hardware_Component_Partition
}

type Hardware_Component_Model struct {
	Entity

	ArchitectureType                    *Hardware_Component_Model_Architecture_Type
	ArchitectureTypeId                  *string
	AttributeCount                      *uint
	Attributes                          []Hardware_Component_Model_Attribute
	Capacity                            *float64
	CompatibleArrayTypeCount            *uint
	CompatibleArrayTypes                []Configuration_Storage_Group_Array_Type
	CompatibleChildComponentModelCount  *uint
	CompatibleChildComponentModels      []Hardware_Component_Model
	CompatibleParentComponentModelCount *uint
	CompatibleParentComponentModels     []Hardware_Component_Model
	Description                         *string
	HardwareComponents                  []Hardware_Component
	HardwareGenericComponentModel       *Hardware_Component_Model_Generic
	HardwareGenericComponentModelId     *int
	Id                                  *int
	InfinibandCompatibleAttribute       *Hardware_Component_Model_Attribute
	IsInfinibandCompatible              *bool
	LongDescription                     *string
	Manufacturer                        *string
	Name                                *string
	RebootTime                          *Hardware_Component_Motherboard_Reboot_Time
	Type                                *string
	ValidAttributeTypeCount             *uint
	ValidAttributeTypes                 []Hardware_Component_Model_Attribute_Type
	Version                             *string
}

type Hardware_Component_Model_Architecture_Type struct {
	Entity

	Children      []Hardware_Component_Model_Architecture_Type
	ChildrenCount *uint
	Id            *int
	KeyName       *string
	Name          *string
	Parent        *Hardware_Component_Model_Architecture_Type
	ParentId      *string
}

type Hardware_Component_Model_Attribute struct {
	Entity

	AttributeTypeId                *int
	HardwareComponent              *Hardware_Component_Model
	HardwareComponentAttributeType *Hardware_Component_Model_Attribute_Type
	HardwareComponentModelId       *int
	Value                          *string
}

type Hardware_Component_Model_Attribute_Type struct {
	Entity

	Description             *string
	Id                      *int
	KeyName                 *string
	Name                    *string
	ValidComponentTypeCount *uint
	ValidComponentTypes     []Hardware_Component_Type
}

type Hardware_Component_Model_Generic struct {
	Entity

	Capacity                    *float64
	Description                 *string
	HardwareComponentModelCount *uint
	HardwareComponentModels     []Hardware_Component_Model
	HardwareComponentType       *Hardware_Component_Type
	HardwareComponentTypeId     *int
	Id                          *int
	MarketingFeatures           *Hardware_Component_Model_Generic_MarketingFeature
	Units                       *string
	UpgradePriority             *int
}

type Hardware_Component_Model_Generic_Attribute struct {
	Entity

	HardwareGenericComponentModel *Hardware_Component_Model_Generic
	Value                         *string
}

type Hardware_Component_Model_Generic_MarketingFeature struct {
	Entity

	Features                      *string
	HardwareGenericComponentModel *Hardware_Component_Model_Generic
	Price                         *string
}

type Hardware_Component_Motherboard struct {
	Hardware_Component
}

type Hardware_Component_Motherboard_Reboot_Time struct {
	Entity

	HardwareComponentModel *Hardware_Component_Model
	WithRaid               *int
	WithoutRaid            *int
}

type Hardware_Component_NetworkCard struct {
	Hardware_Component
}

type Hardware_Component_Partition struct {
	Entity

	DiskNumber          *int
	Grow                *int
	HardwareComponent   *Hardware_Component
	HardwareComponentId *int
	MinimumSize         *float64
	Name                *string
}

type Hardware_Component_Partition_OperatingSystem struct {
	Entity

	Description            *string
	Id                     *int
	Notes                  *string
	PartitionTemplateCount *uint
	PartitionTemplates     []Hardware_Component_Partition_Template
}

type Hardware_Component_Partition_Template struct {
	Entity

	Account                         *Account
	AccountId                       *int
	Data                            []Hardware_Component_Partition_Template_Partition
	DataCount                       *uint
	Description                     *string
	ExpireDate                      *string
	Id                              *int
	PartitionOperatingSystem        *Hardware_Component_Partition_OperatingSystem
	PartitionOperatingSystemId      *int
	PartitionTemplatePartition      []Hardware_Component_Partition_Template_Partition
	PartitionTemplatePartitionCount *uint
	StatusCode                      *string
	TemplateType                    *string
}

type Hardware_Component_Partition_Template_Partition struct {
	Entity

	FilesystemType      *Configuration_Storage_Filesystem_Type
	Id                  *int
	IsGrow              *bool
	PartitionName       *string
	PartitionSize       *float64
	PartitionTemplate   *Hardware_Component_Partition_Template
	PartitionTemplateId *int
}

type Hardware_Component_Processor struct {
	Hardware_Component
}

type Hardware_Component_Ram struct {
	Hardware_Component
}

type Hardware_Component_RemoteManagement struct {
	Hardware_Component

	NetworkComponent *Network_Component
}

type Hardware_Component_RemoteManagement_Command struct {
	Entity

	KeyName      *string
	RequestCount *uint
	Requests     []Hardware_Component_RemoteManagement_Command_Request
}

type Hardware_Component_RemoteManagement_Command_Request struct {
	Entity

	CreateDate              *time.Time
	Hardware                *Hardware
	HardwareId              *int
	ModifyDate              *time.Time
	NetworkComponent        *Network_Component
	Processed               *bool
	RemoteManagementCommand *Hardware_Component_RemoteManagement_Command
	User                    *User_Customer
}

type Hardware_Component_RemoteManagement_User struct {
	Entity

	Hardware         *Hardware
	NetworkComponent *Network_Component
	Password         *string
	Username         *string
}

type Hardware_Component_SecurityDevice struct {
	Hardware_Component
}

type Hardware_Component_SecurityDevice_Infineon struct {
	Hardware_Component_SecurityDevice
}

type Hardware_Component_Type struct {
	Entity

	HardwareGenericComponentModelCount *uint
	HardwareGenericComponentModels     []Hardware_Component_Model_Generic
	Id                                 *int
	KeyName                            *string
	Type                               *string
	TypeParent                         *Hardware_Component_Type
	TypeParentId                       *int
}

type Hardware_Firewall struct {
	Hardware_Switch

	UserCount *uint
	Users     []User_Customer
}

type Hardware_Function struct {
	Entity

	Code        *string
	Description *string
	Id          *int
}

type Hardware_Group struct {
	Entity

	Domain                                      *string
	DownlinkServerCount                         *uint
	DownlinkServers                             []Hardware
	DownlinkVirtualGuestCount                   *uint
	DownlinkVirtualGuests                       []Virtual_Guest
	DownstreamNetworkHardware                   []Hardware
	DownstreamNetworkHardwareCount              *uint
	DownstreamNetworkHardwareWithIncidentCount  *uint
	DownstreamNetworkHardwareWithIncidents      []Hardware
	HardwareChassis                             *Hardware_Chassis
	Hostname                                    *string
	Id                                          *int
	NetworkMonitorAttachedDownHardware          []Hardware
	NetworkMonitorAttachedDownHardwareCount     *uint
	NetworkMonitorAttachedDownVirtualGuestCount *uint
	NetworkMonitorAttachedDownVirtualGuests     []Virtual_Guest
	NetworkStatus                               *string
}

type Hardware_LoadBalancer struct {
	Hardware

	ModelFamily *string
	UserCount   *uint
	Users       []User_Customer
}

type Hardware_Note struct {
	Entity

	CreateDate   *time.Time
	Employee     *User_Employee
	Hardware     *Hardware
	HardwareId   *int
	Id           *int
	ModifyDate   *time.Time
	Note         *string
	Type         *Hardware_Note_Type
	TypeId       *int
	User         *User_Customer
	UserRecordId *int
}

type Hardware_Note_Type struct {
	Entity

	KeyName *string
}

type Hardware_Power_Component struct {
	Entity

	Hardware   *Hardware
	HardwareId *int
	Id         *int
}

type Hardware_Router struct {
	Hardware_Switch

	BoundSubnetCount               *uint
	BoundSubnets                   []Network_Subnet
	LocalDiskStorageCapabilityFlag *bool
	SanStorageCapabilityFlag       *bool
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

	ActiveNetworkFirewallBillingItem     *Billing_Item
	ActiveTicketCount                    *uint
	ActiveTickets                        []Ticket
	ActiveTransaction                    *Provisioning_Version1_Transaction
	ActiveTransactionCount               *uint
	ActiveTransactions                   []Provisioning_Version1_Transaction
	AvailableMonitoring                  []Network_Monitor_Version1_Query_Host_Stratum
	AvailableMonitoringCount             *uint
	AverageDailyBandwidthUsage           *float64
	AverageDailyPrivateBandwidthUsage    *float64
	BillingCycleBandwidthUsage           []Network_Bandwidth_Usage
	BillingCycleBandwidthUsageCount      *uint
	BillingCyclePrivateBandwidthUsage    *Network_Bandwidth_Usage
	BillingCyclePublicBandwidthUsage     *Network_Bandwidth_Usage
	ContainsSolidStateDrivesFlag         *bool
	ControlPanel                         *Software_Component_ControlPanel
	Cost                                 *float64
	CurrentBandwidthSummary              *Metric_Tracking_Object_Bandwidth_Summary
	CustomerInstalledOperatingSystemFlag *bool
	CustomerOwnedFlag                    *bool
	InboundPrivateBandwidthUsage         *float64
	LastOperatingSystemReload            *Provisioning_Version1_Transaction
	MetricTrackingObjectId               *int
	MonitoringUserNotification           []User_Customer_Notification_Hardware
	MonitoringUserNotificationCount      *uint
	OpenCancellationTicket               *Ticket
	OutboundPrivateBandwidthUsage        *float64
	OverBandwidthAllocationFlag          *int
	PrivateIpAddress                     *string
	ProjectedOverBandwidthAllocationFlag *int
	ProjectedPublicBandwidthUsage        *float64
	RecentRemoteManagementCommandCount   *uint
	RecentRemoteManagementCommands       []Hardware_Component_RemoteManagement_Command_Request
	RegionalInternetRegistry             *Network_Regional_Internet_Registry
	RemoteManagement                     *Hardware_Component_RemoteManagement
	RemoteManagementUserCount            *uint
	RemoteManagementUsers                []Hardware_Component_RemoteManagement_User
	StatisticsRemoteManagement           *Hardware_Component_RemoteManagement
	UserCount                            *uint
	Users                                []User_Customer
	VirtualGuestCount                    *uint
	VirtualGuests                        []Virtual_Guest
}

type Hardware_Status struct {
	Entity

	Id     *int
	Status *string
}

type Hardware_Switch struct {
	Hardware
}

type Layout_Container struct {
	Entity

	Id                    *int
	Keyname               *string
	LayoutContainerType   *Layout_Container_Type
	LayoutContainerTypeId *int
	LayoutItemCount       *uint
	LayoutItems           []Layout_Item
	Name                  *string
}

type Layout_Container_Type struct {
	Entity

	Id      *int
	Keyname *string
	Name    *string
}

type Layout_Item struct {
	Entity

	Id                        *int
	Keyname                   *string
	LayoutItemPreferenceCount *uint
	LayoutItemPreferences     []Layout_Preference
	LayoutItemType            *Layout_Item_Type
	LayoutItemTypeId          *int
	Name                      *string
}

type Layout_Item_Type struct {
	Entity

	Id      *int
	Keyname *string
	Name    *string
}

type Layout_Preference struct {
	Entity

	Id                     *int
	LayoutPreferenceType   *Layout_Preference_Type
	LayoutPreferenceTypeId *int
	Value                  *string
}

type Layout_Preference_Type struct {
	Entity

	Id              *int
	Keyname         *string
	Name            *string
	ValueExpression *string
}

type Layout_Profile struct {
	Entity

	ActiveFlag            *int
	CreateDate            *time.Time
	Id                    *int
	LayoutContainerCount  *uint
	LayoutContainers      []Layout_Container
	LayoutPreferenceCount *uint
	LayoutPreferences     []Layout_Profile_Preference
	ModifyDate            *time.Time
	Name                  *string
	UserRecordId          *int
}

type Layout_Profile_Containers struct {
	Entity

	CreateDate          *time.Time
	Id                  *int
	LayoutContainerId   *int
	LayoutContainerType *Layout_Container
	LayoutProfile       *Layout_Profile
	LayoutProfileId     *int
	ModifyDate          *time.Time
}

type Layout_Profile_Customer struct {
	Layout_Profile

	UserRecord *User_Customer
}

type Layout_Profile_Preference struct {
	Entity

	CreateDate         *time.Time
	DefaultValueFlag   *int
	LayoutContainer    *Layout_Container
	LayoutContainerId  *int
	LayoutItem         *Layout_Item
	LayoutItemId       *int
	LayoutPreference   *Layout_Preference
	LayoutPreferenceId *int
	LayoutProfile      *Layout_Profile
	LayoutProfileId    *int
	ModifyDate         *time.Time
	Value              *string
}

type Legal_RegulatedWorkload struct {
	Entity

	Account        *Account
	AccountId      *int
	EnabledFlag    *bool
	Id             *int
	Type           *Legal_RegulatedWorkload_Type
	WorkloadTypeId *int
}

type Legal_RegulatedWorkload_Type struct {
	Entity

	Id      *int
	KeyName *string
	Name    *string
}

type Locale struct {
	Entity

	FriendlyName *string
	Id           *int
	LanguageTag  *string
	Name         *string
}

type Locale_Country struct {
	Entity

	IsEuropeanUnionFlag *int
	LongName            *string
	ShortName           *string
	StateCount          *uint
	States              []Locale_StateProvince
}

type Locale_StateProvince struct {
	Entity

	LongName  *string
	ShortName *string
}

type Locale_Timezone struct {
	Entity

	Id        *int
	LongName  *string
	Name      *string
	Offset    *string
	ShortName *string
}

type Location struct {
	Entity

	BackboneDependentCount        *uint
	BackboneDependents            []Network_Backbone_Location_Dependent
	GroupCount                    *uint
	Groups                        []Location_Group
	HardwareFirewallCount         *uint
	HardwareFirewalls             []Hardware
	Id                            *int
	LocationAddress               *Account_Address
	LocationReservationMember     *Location_Reservation_Rack_Member
	LocationStatus                *Location_Status
	LongName                      *string
	Name                          *string
	NetworkConfigurationAttribute *Hardware_Attribute
	OnlinePptpVpnUserCount        *int
	OnlineSslVpnUserCount         *int
	PathString                    *string
	PriceGroupCount               *uint
	PriceGroups                   []Location_Group
	RegionCount                   *uint
	Regions                       []Location_Region
	StatusId                      *int
	Timezone                      *Locale_Timezone
	VdrGroup                      *Location_Group_Location_CrossReference
}

type Location_Datacenter struct {
	Location

	ActiveItemPresaleEventCount  *uint
	ActiveItemPresaleEvents      []Sales_Presale_Event
	ActivePresaleEventCount      *uint
	ActivePresaleEvents          []Sales_Presale_Event
	BackendHardwareRouterCount   *uint
	BackendHardwareRouters       []Hardware
	BoundSubnetCount             *uint
	BoundSubnets                 []Network_Subnet
	BrandCountryRestrictionCount *uint
	BrandCountryRestrictions     []Brand_Restriction_Location_CustomerCountry
	FrontendHardwareRouterCount  *uint
	FrontendHardwareRouters      []Hardware
	HardwareRouterCount          *uint
	HardwareRouters              []Hardware
	PresaleEventCount            *uint
	PresaleEvents                []Sales_Presale_Event
	RegionalGroup                *Location_Group_Regional
	RegionalInternetRegistry     *Network_Regional_Internet_Registry
	RoutableBoundSubnetCount     *uint
	RoutableBoundSubnets         []Network_Subnet
}

type Location_Group struct {
	Entity

	Description         *string
	Id                  *int
	LocationCount       *uint
	LocationGroupType   *Location_Group_Type
	LocationGroupTypeId *int
	Locations           []Location
	Name                *string
	SecurityLevelId     *int
}

type Location_Group_Location_CrossReference struct {
	Entity

	Location        *Location
	LocationGroup   *Location_Group
	LocationGroupId *int
	LocationId      *int
	Priority        *int
}

type Location_Group_Pricing struct {
	Location_Group

	PriceCount *uint
	Prices     []Product_Item_Price
}

type Location_Group_Regional struct {
	Location_Group

	DatacenterCount     *uint
	Datacenters         []Location
	PreferredDatacenter *Location_Datacenter
}

type Location_Group_Type struct {
	Entity

	Name *string
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

	Description *string
	Keyname     *string
	Location    *Location_Region_Location
	SortOrder   *int
}

type Location_Region_Location struct {
	Entity

	Location                   *Location
	LocationPackageDetailCount *uint
	LocationPackageDetails     []Product_Package_Locations
	Region                     *Location_Region
}

type Location_Reservation struct {
	Entity

	Account                 *Account
	Allotment               *Network_Bandwidth_Version1_Allotment
	AllotmentId             *int
	BillingItem             *Billing_Item
	Id                      *int
	Location                *Location
	LocationId              *int
	LocationReservationRack *Location_Reservation_Rack
	Name                    *string
	Notes                   *string
}

type Location_Reservation_Rack struct {
	Entity

	Allotment                    *Network_Bandwidth_Version1_Allotment
	Children                     []Location_Reservation_Rack_Member
	ChildrenCount                *uint
	Location                     *Location
	LocationId                   *int
	LocationReservation          *Location_Reservation
	LocationReservationId        *int
	NetworkConnectionCapacity    *int
	NetworkConnectionReservation *int
	PowerConnectionCapacity      *int
	PowerConnectionReservation   *int
	SlotCapacity                 *int
	SlotReservation              *int
}

type Location_Reservation_Rack_Member struct {
	Entity

	Id                      *int
	Location                *Location
	LocationId              *int
	LocationReservationRack *Location_Reservation
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

	Id     *int
	Status *string
}

type Location_Storage_Room struct {
	Location
}

type Marketplace_EmailDistribution struct {
	Entity

	Email *string
	Id    *int
}

type Marketplace_Partner struct {
	Entity

	AccountId               *int
	AttachedFiles           []Marketplace_Partner_Attachment
	AttachmentCount         *uint
	Attachments             []Marketplace_Partner_Attachment
	CompanyDescription      *string
	CompanyName             *string
	HeadlineDescription     *string
	Id                      *int
	LinkFreeTrial           *string
	LinkOrderPage           *string
	LinkWebsite             *string
	LogoMedium              *Marketplace_Partner_Attachment
	LogoMediumTemp          *Marketplace_Partner_Attachment
	LogoSmall               *Marketplace_Partner_Attachment
	LogoSmallTemp           *Marketplace_Partner_Attachment
	MetaDescription         *string
	MetaKeywords            *string
	ProductBenefits         *string
	ProductCategoryId       *int
	ProductDescriptionLong  *string
	ProductDescriptionShort *string
	ProductFeatures         *string
	ProductName             *string
	ProductTitle            *string
	UrlIdentifier           *string
}

type Marketplace_Partner_Attachment struct {
	Entity

	AttachmentType       *Marketplace_Partner_Attachment_Type
	AttachmentTypeId     *int
	BaseName             *string
	DisplayName          *string
	FileName             *string
	Id                   *int
	MarketplacePartnerId *int
	SaveAsName           *string
}

type Marketplace_Partner_Attachment_Type struct {
	Entity

	Id      *int
	KeyName *string
	Type    *string
}

type Marketplace_Partner_File struct {
	Entity

	Attributes *Marketplace_Partner_File_Attributes
	Contents   *[]byte
}

type Marketplace_Partner_File_Attributes struct {
	Entity

	Bits           *int
	Channels       *int
	Height         *int
	HtmlAttributes *string
	ImageType      *int
	IsImage        *bool
	MimeType       *string
	Width          *int
}

type Metric_Tracking_Object struct {
	Entity

	Data            []Metric_Tracking_Object_Data
	Id              *int
	Label           *string
	ResourceTableId *int
	StartDate       *time.Time
	Type            *Metric_Tracking_Object_Type
}

type Metric_Tracking_Object_Abstract struct {
	Metric_Tracking_Object
}

type Metric_Tracking_Object_Bandwidth_Summary struct {
	Entity

	AllocationAmount            *float64
	AllocationId                *int
	AmountOut                   *float64
	AverageDailyUsage           *float64
	CurrentlyOverAllocationFlag *int
	Id                          *int
	OutboundBandwidthAmount     *float64
	ProjectedBandwidthUsage     *float64
	ProjectedOverAllocationFlag *int
}

type Metric_Tracking_Object_Data struct {
	Entity

	Counter  *float64
	DateTime *time.Time
	Type     *string
}

type Metric_Tracking_Object_Data_Network_ContentDelivery_Account struct {
	Metric_Tracking_Object_Data

	FileName *string
	PopId    *int
}

type Metric_Tracking_Object_HardwareServer struct {
	Metric_Tracking_Object_Abstract

	BillingCycleBandwidthUsage             []Network_Bandwidth_Usage
	BillingCycleBandwidthUsageCount        *uint
	BillingCyclePrivateBandwidthUsage      []Network_Bandwidth_Usage
	BillingCyclePrivateBandwidthUsageCount *uint
	BillingCyclePrivateUsageIn             *float64
	BillingCyclePrivateUsageOut            *float64
	BillingCyclePrivateUsageTotal          *uint
	BillingCyclePublicBandwidthUsage       *Network_Bandwidth_Usage
	BillingCyclePublicUsageIn              *float64
	BillingCyclePublicUsageOut             *float64
	BillingCyclePublicUsageTotal           *uint
	Resource                               *Hardware_Server
}

type Metric_Tracking_Object_Type struct {
	Entity

	Keyname *string
	Name    *string
}

type Metric_Tracking_Object_VirtualDedicatedRack struct {
	Metric_Tracking_Object_Abstract

	BillingCycleBandwidthUsage             []Network_Bandwidth_Usage
	BillingCycleBandwidthUsageCount        *uint
	BillingCyclePrivateBandwidthUsage      []Network_Bandwidth_Usage
	BillingCyclePrivateBandwidthUsageCount *uint
	BillingCyclePrivateUsageIn             *float64
	BillingCyclePrivateUsageOut            *float64
	BillingCyclePrivateUsageTotal          *uint
	BillingCyclePublicBandwidthUsage       *Network_Bandwidth_Usage
	BillingCyclePublicUsageIn              *float64
	BillingCyclePublicUsageOut             *float64
	BillingCyclePublicUsageTotal           *uint
	Resource                               *Network_Bandwidth_Version1_Allotment
}

type Metric_Tracking_Object_Virtual_Storage_Repository struct {
	Metric_Tracking_Object_Abstract

	Resource *Virtual_Storage_Repository
}

type Monitoring_Agent struct {
	Entity

	AgentStatus               *Monitoring_Agent_Status
	ConfigurationProfileCount *uint
	ConfigurationProfiles     []Configuration_Template_Section_Profile
	ConfigurationTemplate     *Configuration_Template
	ConfigurationTemplateId   *int
	ConfigurationValueCount   *uint
	ConfigurationValues       []Monitoring_Agent_Configuration_Value
	CreateDate                *time.Time
	Hardware                  *Hardware
	Id                        *int
	ModifyDate                *time.Time
	Name                      *string
	ProductItem               *Product_Item
	RemoteMonitoringAgentFlag *bool
	RobotId                   *int
	SoftwareDescription       *Software_Description
	StatusId                  *int
	StatusName                *string
	VirtualGuest              *Virtual_Guest
}

type Monitoring_Agent_Configuration_Template_Group struct {
	Entity

	Account                             *Account
	AccountId                           *int
	ConfigurationTemplateCount          *uint
	ConfigurationTemplateReferenceCount *uint
	ConfigurationTemplateReferences     []Monitoring_Agent_Configuration_Template_Group_Reference
	ConfigurationTemplates              []Configuration_Template
	CreateDate                          *time.Time
	Description                         *string
	Id                                  *int
	Item                                *Product_Item
	ItemId                              *int
	ModifyDate                          *time.Time
	Name                                *string
}

type Monitoring_Agent_Configuration_Template_Group_Reference struct {
	Entity

	ConfigurationTemplate   *Configuration_Template
	ConfigurationTemplateId *int
	Id                      *int
	TemplateGroup           *Monitoring_Agent_Configuration_Template_Group
	TemplateGroupId         *int
}

type Monitoring_Agent_Configuration_Value struct {
	Entity

	AgentId                   *int
	ConfigurationDefinitionId *int
	Definition                *Configuration_Template_Section_Definition
	Description               *string
	Id                        *int
	MetricDataType            *Container_Metric_Data_Type
	MonitoringAgent           *Monitoring_Agent
	Profile                   *Configuration_Template_Section_Profile
	ProfileId                 *int
	Value                     *string
}

type Monitoring_Agent_Status struct {
	Entity

	Description *string
	Id          *int
	Name        *string
}

type Monitoring_Robot struct {
	Entity

	Account              *Account
	AccountId            *int
	Id                   *int
	MonitoringAgentCount *uint
	MonitoringAgents     []Monitoring_Agent
	Name                 *string
	RobotStatus          *Monitoring_Robot_Status
	SoftwareComponent    *Software_Component
	StatusId             *int
}

type Monitoring_Robot_Status struct {
	Entity

	Description *string
	Id          *int
	Name        *string
}

type Network struct {
	Entity

	AccountId         *int
	Cidr              *int
	Id                *int
	Name              *string
	NetworkIdentifier *string
	Notes             *string
	SubnetCount       *uint
	Subnets           []Network_Subnet
}

type Network_Application_Delivery_Controller struct {
	Entity

	Account                          *Account
	AccountId                        *int
	AverageDailyPublicBandwidthUsage *float64
	BillingItem                      *Billing_Item_Network_Application_Delivery_Controller
	ConfigurationHistory             []Network_Application_Delivery_Controller_Configuration_History
	ConfigurationHistoryCount        *uint
	CreateDate                       *time.Time
	Datacenter                       *Location
	Description                      *string
	Id                               *int
	LicenseExpirationDate            *time.Time
	LoadBalancerCount                *uint
	LoadBalancers                    []Network_LoadBalancer_VirtualIpAddress
	ManagedResourceFlag              *bool
	ManagementIpAddress              *string
	ModifyDate                       *time.Time
	Name                             *string
	NetworkVlan                      *Network_Vlan
	NetworkVlanCount                 *uint
	NetworkVlans                     []Network_Vlan
	Notes                            *string
	OutboundPublicBandwidthUsage     *float64
	Password                         *Software_Component_Password
	PrimaryIpAddress                 *string
	ProjectedPublicBandwidthUsage    *float64
	SubnetCount                      *uint
	Subnets                          []Network_Subnet
	TagReferenceCount                *uint
	TagReferences                    []Tag_Reference
	Type                             *Network_Application_Delivery_Controller_Type
	TypeId                           *int
	VirtualIpAddressCount            *uint
	VirtualIpAddresses               []Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress
}

type Network_Application_Delivery_Controller_Configuration_History struct {
	Entity

	Controller *Network_Application_Delivery_Controller
	CreateDate *time.Time
	Id         *int
	Notes      *string
}

type Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute struct {
	Entity

	HealthAttributeTypeId *int
	HealthCheck           *Network_Application_Delivery_Controller_LoadBalancer_Health_Check
	HealthCheckId         *int
	Id                    *int
	Type                  *Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type
	Value                 *string
}

type Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type struct {
	Entity

	Description     *string
	Id              *int
	Keyname         *string
	Name            *string
	ValueExpression *string
}

type Network_Application_Delivery_Controller_LoadBalancer_Health_Check struct {
	Entity

	AttributeCount         *uint
	Attributes             []Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute
	HealthCheckTypeId      *int
	Id                     *int
	Name                   *string
	Notes                  *string
	ScaleLoadBalancerCount *uint
	ScaleLoadBalancers     []Scale_LoadBalancer
	ServiceCount           *uint
	Services               []Network_Application_Delivery_Controller_LoadBalancer_Service
	Type                   *Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type
}

type Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type struct {
	Entity

	Id      *int
	Keyname *string
	Name    *string
}

type Network_Application_Delivery_Controller_LoadBalancer_Routing_Method struct {
	Entity

	Id      *int
	Keyname *string
	Name    *string
}

type Network_Application_Delivery_Controller_LoadBalancer_Routing_Type struct {
	Entity

	Id      *int
	Keyname *string
	Name    *string
}

type Network_Application_Delivery_Controller_LoadBalancer_Service struct {
	Entity

	Enabled             *int
	GroupCount          *uint
	GroupReferenceCount *uint
	GroupReferences     []Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference
	Groups              []Network_Application_Delivery_Controller_LoadBalancer_Service_Group
	HealthCheck         *Network_Application_Delivery_Controller_LoadBalancer_Health_Check
	HealthCheckCount    *uint
	HealthChecks        []Network_Application_Delivery_Controller_LoadBalancer_Health_Check
	Id                  *int
	IpAddress           *Network_Subnet_IpAddress
	IpAddressId         *int
	Name                *string
	Notes               *string
	Port                *int
	ServiceGroup        *Network_Application_Delivery_Controller_LoadBalancer_Service_Group
	Status              *string
}

type Network_Application_Delivery_Controller_LoadBalancer_Service_Group struct {
	Entity

	Id                    *int
	Name                  *string
	Notes                 *string
	RoutingMethod         *Network_Application_Delivery_Controller_LoadBalancer_Routing_Method
	RoutingMethodId       *int
	RoutingType           *Network_Application_Delivery_Controller_LoadBalancer_Routing_Type
	RoutingTypeId         *int
	ServiceCount          *uint
	ServiceReferenceCount *uint
	ServiceReferences     []Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference
	Services              []Network_Application_Delivery_Controller_LoadBalancer_Service
	Timeout               *int
	VirtualServer         *Network_Application_Delivery_Controller_LoadBalancer_VirtualServer
	VirtualServerCount    *uint
	VirtualServers        []Network_Application_Delivery_Controller_LoadBalancer_VirtualServer
}

type Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference struct {
	Entity

	Service        *Network_Application_Delivery_Controller_LoadBalancer_Service
	ServiceGroup   *Network_Application_Delivery_Controller_LoadBalancer_Service_Group
	ServiceGroupId *int
	ServiceId      *int
	Weight         *int
}

type Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress struct {
	Entity

	Account                            *Account
	AccountId                          *int
	ApplicationDeliveryController      *Network_Application_Delivery_Controller
	ApplicationDeliveryControllerCount *uint
	ApplicationDeliveryControllers     []Network_Application_Delivery_Controller
	BillingItem                        *Billing_Item
	ConnectionLimit                    *int
	ConnectionLimitUnits               *string
	DedicatedBillingItem               *Billing_Item_Network_LoadBalancer
	DedicatedFlag                      *bool
	HighAvailabilityFlag               *bool
	Id                                 *int
	IpAddress                          *Network_Subnet_IpAddress
	IpAddressId                        *int
	LoadBalancerHardware               []Hardware
	LoadBalancerHardwareCount          *uint
	ManagedResourceFlag                *bool
	Notes                              *string
	SecureTransportCipherCount         *uint
	SecureTransportCiphers             []Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportCipher
	SecureTransportProtocolCount       *uint
	SecureTransportProtocols           []Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportProtocol
	SecurityCertificate                *Security_Certificate
	SecurityCertificateEntry           *Security_Certificate_Entry
	SecurityCertificateId              *int
	SslActiveFlag                      *bool
	SslEnabledFlag                     *bool
	VirtualServerCount                 *uint
	VirtualServers                     []Network_Application_Delivery_Controller_LoadBalancer_VirtualServer
}

type Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportCipher struct {
	Entity

	Id                 *int
	KeyName            *string
	VirtualIpAddress   *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress
	VirtualIpAddressId *int
}

type Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportProtocol struct {
	Entity

	Id                 *int
	KeyName            *string
	VirtualIpAddress   *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress
	VirtualIpAddressId *int
}

type Network_Application_Delivery_Controller_LoadBalancer_VirtualServer struct {
	Entity

	Allocation             *int
	Id                     *int
	Name                   *string
	Notes                  *string
	Port                   *int
	RoutingMethod          *Network_Application_Delivery_Controller_LoadBalancer_Routing_Method
	RoutingMethodId        *int
	ScaleLoadBalancerCount *uint
	ScaleLoadBalancers     []Scale_LoadBalancer
	ServiceGroupCount      *uint
	ServiceGroups          []Network_Application_Delivery_Controller_LoadBalancer_Service_Group
	VirtualIpAddress       *Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress
	VirtualIpAddressId     *int
}

type Network_Application_Delivery_Controller_Type struct {
	Entity

	KeyName *string
	Name    *string
}

type Network_Backbone struct {
	Entity

	Capacity           *int
	CapacityUnits      *string
	Health             *string
	Id                 *int
	Location           *Location
	Name               *string
	NetworkComponent   *Network_Component
	NetworkComponentId *int
	Type               *string
}

type Network_Backbone_Location_Dependent struct {
	Entity

	DependentLocation   *Location
	DependentLocationId *int
	Id                  *int
	SourceLocation      *Location
	SourceLocationId    *int
}

type Network_Bandwidth_Usage struct {
	Entity

	AmountIn                   *float64
	AmountOut                  *float64
	BandwidthUsageDetailTypeId *float64
	TrackingObject             *Metric_Tracking_Object
	Type                       *Network_Bandwidth_Version1_Usage_Detail_Type
}

type Network_Bandwidth_Usage_Detail struct {
	Entity

	Account                    *Account
	AmountIn                   *float64
	AmountOut                  *float64
	BandwidthUsageDetailTypeId *float64
	TrackingObject             *Metric_Tracking_Object
	Type                       *Network_Bandwidth_Version1_Usage_Detail_Type
}

type Network_Bandwidth_Version1_Allocation struct {
	Entity

	AllotmentDetail *Network_Bandwidth_Version1_Allotment_Detail
	Amount          *float64
	BillingItem     *Billing_Item_Hardware
	Id              *int
}

type Network_Bandwidth_Version1_Allotment struct {
	Entity

	Account                              *Account
	AccountId                            *int
	ActiveDetailCount                    *uint
	ActiveDetails                        []Network_Bandwidth_Version1_Allotment_Detail
	ApplicationDeliveryControllerCount   *uint
	ApplicationDeliveryControllers       []Network_Application_Delivery_Controller
	AverageDailyPublicBandwidthUsage     *float64
	BandwidthAllotmentTypeId             *int
	BareMetalInstanceCount               *uint
	BareMetalInstances                   []Hardware
	BillingCycleBandwidthUsage           []Network_Bandwidth_Usage
	BillingCycleBandwidthUsageCount      *uint
	BillingCyclePrivateBandwidthUsage    *Network_Bandwidth_Usage
	BillingCyclePublicBandwidthUsage     *Network_Bandwidth_Usage
	BillingCyclePublicUsageTotal         *uint
	BillingItem                          *Billing_Item
	CreateDate                           *time.Time
	CurrentBandwidthSummary              *Metric_Tracking_Object_Bandwidth_Summary
	DetailCount                          *uint
	Details                              []Network_Bandwidth_Version1_Allotment_Detail
	EndDate                              *time.Time
	Hardware                             []Hardware
	HardwareCount                        *uint
	Id                                   *int
	InboundPublicBandwidthUsage          *float64
	LocationGroup                        *Location_Group
	LocationGroupId                      *int
	ManagedBareMetalInstanceCount        *uint
	ManagedBareMetalInstances            []Hardware
	ManagedHardware                      []Hardware
	ManagedHardwareCount                 *uint
	ManagedVirtualGuestCount             *uint
	ManagedVirtualGuests                 []Virtual_Guest
	MetricTrackingObject                 *Metric_Tracking_Object_VirtualDedicatedRack
	MetricTrackingObjectId               *int
	Name                                 *string
	OutboundPublicBandwidthUsage         *float64
	OverBandwidthAllocationFlag          *int
	PrivateNetworkOnlyHardware           []Hardware
	PrivateNetworkOnlyHardwareCount      *uint
	ProjectedOverBandwidthAllocationFlag *int
	ProjectedPublicBandwidthUsage        *float64
	ServiceProvider                      *Service_Provider
	ServiceProviderId                    *int
	TotalBandwidthAllocated              *uint
	VirtualGuestCount                    *uint
	VirtualGuests                        []Virtual_Guest
}

type Network_Bandwidth_Version1_Allotment_Detail struct {
	Entity

	Allocation           *Network_Bandwidth_Version1_Allocation
	AllocationId         *int
	BandwidthAllotment   *Network_Bandwidth_Version1_Allotment
	BandwidthAllotmentId *int
	BandwidthUsage       []Network_Bandwidth_Version1_Usage
	BandwidthUsageCount  *uint
	EffectiveDate        *time.Time
	EndEffectiveDate     *time.Time
	Id                   *int
	ServiceProviderId    *int
}

type Network_Bandwidth_Version1_Host struct {
	Entity

	PodId *int
}

type Network_Bandwidth_Version1_Interface struct {
	Entity

	Host               *Network_Bandwidth_Version1_Host
	HostId             *int
	NetworkComponent   *Network_Component
	NetworkComponentId *int
}

type Network_Bandwidth_Version1_Usage struct {
	Entity

	BandwidthAllotmentDetail  *Network_Bandwidth_Version1_Allotment_Detail
	BandwidthUsageDetail      []Network_Bandwidth_Version1_Usage_Detail
	BandwidthUsageDetailCount *uint
}

type Network_Bandwidth_Version1_Usage_Detail struct {
	Entity

	AmountIn                 *float64
	AmountOut                *float64
	BandwidthUsage           *Network_Bandwidth_Version1_Usage
	BandwidthUsageDetailType *Network_Bandwidth_Version1_Usage_Detail_Type
	Day                      *time.Time
}

type Network_Bandwidth_Version1_Usage_Detail_Total struct {
	Entity

	Account                    *Account
	AmountIn                   *float64
	AmountOut                  *float64
	BandwidthUsageDetailTypeId *float64
	TrackingObject             *Metric_Tracking_Object
	Type                       *Network_Bandwidth_Version1_Usage_Detail_Type
}

type Network_Bandwidth_Version1_Usage_Detail_Type struct {
	Entity

	Alias *string
}

type Network_Component struct {
	Entity

	ActiveCommand                  *Hardware_Component_RemoteManagement_Command_Request
	DownlinkComponent              *Network_Component
	DuplexMode                     *Network_Component_Duplex_Mode
	DuplexModeId                   *string
	Hardware                       *Hardware
	HardwareId                     *int
	HighAvailabilityFirewallFlag   *bool
	Id                             *int
	Interface                      *Network_Bandwidth_Version1_Interface
	IpAddressBindingCount          *uint
	IpAddressBindings              []Network_Component_IpAddress
	IpAddressCount                 *uint
	IpAddresses                    []Network_Subnet_IpAddress
	IpmiIpAddress                  *string
	IpmiMacAddress                 *string
	LastCommand                    *Hardware_Component_RemoteManagement_Command_Request
	MacAddress                     *string
	MaxSpeed                       *int
	MetricTrackingObject           *Metric_Tracking_Object
	ModifyDate                     *time.Time
	Name                           *string
	NetworkComponentFirewall       *Network_Component_Firewall
	NetworkComponentGroup          *Network_Component_Group
	NetworkHardware                []Hardware
	NetworkHardwareCount           *uint
	NetworkVlan                    *Network_Vlan
	NetworkVlanId                  *int
	NetworkVlanTrunkCount          *uint
	NetworkVlanTrunks              []Network_Component_Network_Vlan_Trunk
	Port                           *int
	PrimaryIpAddress               *string
	PrimaryIpAddressRecord         *Network_Subnet_IpAddress
	PrimarySubnet                  *Network_Subnet
	PrimaryVersion6IpAddressRecord *Network_Subnet_IpAddress
	RecentCommandCount             *uint
	RecentCommands                 []Hardware_Component_RemoteManagement_Command_Request
	RedundancyCapableFlag          *bool
	RedundancyEnabledFlag          *bool
	RemoteManagementUserCount      *uint
	RemoteManagementUsers          []Hardware_Component_RemoteManagement_User
	Router                         *Hardware
	Speed                          *int
	Status                         *string
	StorageNetworkFlag             *bool
	SubnetCount                    *uint
	Subnets                        []Network_Subnet
	UplinkComponent                *Network_Component
	UplinkDuplexMode               *Network_Component_Duplex_Mode
}

type Network_Component_Duplex_Mode struct {
	Entity

	Description *string
	Keyname     *string
	Name        *string
}

type Network_Component_Firewall struct {
	Entity

	ApplyServerRuleSubnetCount        *uint
	ApplyServerRuleSubnets            []Network_Subnet
	BillingItem                       *Billing_Item
	GuestNetworkComponent             *Virtual_Guest_Network_Component
	GuestNetworkComponentId           *int
	Id                                *int
	NetworkComponent                  *Network_Component
	NetworkComponentId                *int
	NetworkFirewallUpdateRequest      []Network_Firewall_Update_Request
	NetworkFirewallUpdateRequestCount *uint
	RuleCount                         *uint
	Rules                             []Network_Component_Firewall_Rule
	Status                            *string
	SubnetCount                       *uint
	Subnets                           []Network_Subnet
}

type Network_Component_Firewall_Rule struct {
	Entity

	Action                    *string
	DestinationIpAddress      *string
	DestinationIpCidr         *int
	DestinationIpSubnetMask   *string
	DestinationPortRangeEnd   *int
	DestinationPortRangeStart *int
	Id                        *int
	NetworkComponentFirewall  *Network_Component_Firewall
	Notes                     *string
	OrderValue                *int
	Protocol                  *string
	SourceIpAddress           *string
	SourceIpCidr              *int
	SourceIpSubnetMask        *string
	Status                    *string
	Version                   *int
}

type Network_Component_Firewall_Subnets struct {
	Entity

	ApplyServerRulesFlag     *bool
	NetworkComponentFirewall *Network_Component_Firewall
	Subnet                   *Network_Subnet
	SubnetId                 *int
}

type Network_Component_Group struct {
	Entity

	GroupTypeId           *int
	NetworkComponentCount *uint
	NetworkComponents     []Network_Component
}

type Network_Component_IpAddress struct {
	Entity

	IpAddress        *Network_Subnet_IpAddress
	NetworkComponent *Network_Component
}

type Network_Component_Network_Vlan_Trunk struct {
	Entity

	NetworkComponent   *Network_Component
	NetworkComponentId *int
	NetworkVlan        *Network_Vlan
	NetworkVlanId      *int
}

type Network_Component_RemoteManagement struct {
	Network_Component
}

type Network_Component_Uplink_Hardware struct {
	Entity

	Hardware         *Hardware
	NetworkComponent *Network_Component
}

type Network_ContentDelivery_Account struct {
	Entity

	Account                        *Account
	AccountId                      *int
	AssociatedCdnAccountId         *string
	AuthenticationIpAddressCount   *uint
	AuthenticationIpAddresses      []Network_ContentDelivery_Authentication_Address
	BillingItem                    *Billing_Item
	CdnAccountName                 *string
	CdnAccountNote                 *string
	CdnSolutionName                *string
	CreateDate                     *time.Time
	DependantServiceFlag           *bool
	Id                             *int
	LegacyCdnFlag                  *bool
	LogEnabledFlag                 *string
	ProviderPortalAccessFlag       *bool
	Status                         *Network_ContentDelivery_Account_Status
	StatusId                       *int
	TokenAuthenticationEnabledFlag *bool
}

type Network_ContentDelivery_Account_Status struct {
	Entity

	Description *string
	Id          *int
	Name        *string
}

type Network_ContentDelivery_Authentication_Address struct {
	Entity

	AccessType   *string
	CdnAccountId *int
	CreateDate   *time.Time
	Id           *int
	IpAddress    *string
	ModifyDate   *time.Time
	Name         *string
	Priority     *int
	UserId       *int
}

type Network_ContentDelivery_Authentication_Token struct {
	Entity

	CdnAccountId *int
	ClientIp     *string
	CreateDate   *time.Time
	Name         *string
	Referrer     *string
	Token        *string
}

type Network_Customer_Subnet struct {
	Entity

	AccountId         *int
	Cidr              *int
	Id                *int
	IpAddressCount    *uint
	IpAddresses       []Network_Customer_Subnet_IpAddress
	Netmask           *string
	NetworkIdentifier *string
	TotalIpAddresses  *int
}

type Network_Customer_Subnet_IpAddress struct {
	Entity

	Id               *int
	IpAddress        *string
	Notes            *string
	Subnet           *Network_Customer_Subnet
	SubnetId         *int
	TranslationCount *uint
	Translations     []Network_Tunnel_Module_Context_Address_Translation
}

type Network_Firewall_AccessControlList struct {
	Entity

	Direction                         *string
	FirewallContextInterfaceId        *int
	Id                                *int
	NetworkFirewallUpdateRequestCount *uint
	NetworkFirewallUpdateRequests     []Network_Firewall_Update_Request
	NetworkVlan                       *Network_Vlan
	RuleCount                         *uint
	Rules                             []Network_Vlan_Firewall_Rule
}

type Network_Firewall_Interface struct {
	Network_Firewall_Module_Context_Interface
}

type Network_Firewall_Module_Context_Interface struct {
	Entity

	FirewallContextAccessControlListCount *uint
	FirewallContextAccessControlLists     []Network_Firewall_AccessControlList
	Id                                    *int
	Name                                  *string
	NetworkVlan                           *Network_Vlan
}

type Network_Firewall_Template struct {
	Entity

	Id        *int
	Name      *string
	RuleCount *uint
	Rules     []Network_Firewall_Template_Rule
}

type Network_Firewall_Template_Rule struct {
	Entity

	Action                    *string
	DestinationIpAddress      *string
	DestinationIpSubnetMask   *string
	DestinationPortRangeEnd   *int
	DestinationPortRangeStart *int
	FirewallTemplate          *Network_Firewall_Template
	FirewallTemplateId        *int
	Id                        *int
	Notes                     *string
	OrderValue                *int
	Protocol                  *string
	SourceIpAddress           *string
	SourceIpSubnetMask        *string
}

type Network_Firewall_Update_Request struct {
	Entity

	ApplyDate                          *time.Time
	AuthorizingUser                    *User_Interface
	AuthorizingUserId                  *int
	AuthorizingUserType                *string
	BypassFlag                         *bool
	CreateDate                         *time.Time
	FirewallContextAccessControlListId *int
	Guest                              *Virtual_Guest
	Hardware                           *Hardware
	HardwareId                         *int
	Id                                 *int
	NetworkComponentFirewall           *Network_Component_Firewall
	NetworkComponentFirewallId         *int
	RuleCount                          *uint
	Rules                              []Network_Firewall_Update_Request_Rule
}

type Network_Firewall_Update_Request_Customer struct {
	Network_Firewall_Update_Request
}

type Network_Firewall_Update_Request_Employee struct {
	Network_Firewall_Update_Request
}

type Network_Firewall_Update_Request_Rule struct {
	Entity

	Action                    *string
	DestinationIpAddress      *string
	DestinationIpCidr         *int
	DestinationIpSubnetMask   *string
	DestinationPortRangeEnd   *int
	DestinationPortRangeStart *int
	FirewallUpdateRequest     *Network_Firewall_Update_Request
	FirewallUpdateRequestId   *int
	Id                        *int
	Notes                     *string
	OrderValue                *int
	Protocol                  *string
	SourceIpAddress           *string
	SourceIpCidr              *int
	SourceIpSubnetMask        *string
	Version                   *int
}

type Network_Firewall_Update_Request_Rule_Version6 struct {
	Network_Firewall_Update_Request_Rule
}

type Network_Gateway struct {
	Entity

	Account             *Account
	AccountId           *int
	GroupNumber         *int
	Id                  *int
	InsideVlanCount     *uint
	InsideVlans         []Network_Gateway_Vlan
	MemberCount         *uint
	Members             []Network_Gateway_Member
	Name                *string
	NetworkSpace        *string
	PrivateIpAddress    *Network_Subnet_IpAddress
	PrivateIpAddressId  *int
	PrivateVlan         *Network_Vlan
	PrivateVlanId       *int
	PublicIpAddress     *Network_Subnet_IpAddress
	PublicIpAddressId   *int
	PublicIpv6Address   *Network_Subnet_IpAddress
	PublicIpv6AddressId *int
	PublicVlan          *Network_Vlan
	PublicVlanId        *int
	Status              *Network_Gateway_Status
	StatusId            *int
}

type Network_Gateway_Member struct {
	Entity

	Hardware         *Hardware
	HardwareId       *int
	Id               *int
	NetworkGateway   *Network_Gateway
	NetworkGatewayId *int
	Priority         *int
}

type Network_Gateway_Status struct {
	Entity

	Description *string
	Id          *int
	KeyName     *string
	Name        *string
}

type Network_Gateway_Vlan struct {
	Entity

	BypassFlag       *bool
	Id               *int
	NetworkGateway   *Network_Gateway
	NetworkGatewayId *int
	NetworkVlan      *Network_Vlan
	NetworkVlanId    *int
}

type Network_LoadBalancer_Global_Account struct {
	Entity

	Account                     *Account
	AllowedNumberOfHosts        *int
	AverageConnectionsPerSecond *float64
	BillingItem                 *Billing_Item
	ConnectionsPerSecond        *int
	FallbackIp                  *string
	HostCount                   *uint
	Hostname                    *string
	Hosts                       []Network_LoadBalancer_Global_Host
	Id                          *int
	LoadBalanceType             *Network_LoadBalancer_Global_Type
	LoadBalanceTypeId           *int
	ManagedResourceFlag         *bool
	Notes                       *string
}

type Network_LoadBalancer_Global_Host struct {
	Entity

	DestinationIp       *string
	DestinationPort     *int
	Enabled             *int
	HealthCheck         *string
	Hits                *float64
	Id                  *int
	LoadBalanceOrder    *int
	LoadBalancerAccount *Network_LoadBalancer_Global_Account
	Location            *string
	Status              *string
	Weight              *int
}

type Network_LoadBalancer_Global_Type struct {
	Entity

	Id   *int
	Name *string
}

type Network_LoadBalancer_Service struct {
	Entity

	ConnectionLimit      *int
	CreateDate           *time.Time
	DestinationIpAddress *string
	DestinationPort      *int
	Enabled              *bool
	HealthCheck          *string
	HealthCheckURL       *string
	HealthResponse       *string
	Id                   *int
	ModifyDate           *time.Time
	Name                 *string
	Notes                *string
	PeakConnections      *int
	SourcePort           *int
	Type                 *string
	Vip                  *Network_LoadBalancer_VirtualIpAddress
	VipId                *int
	Weight               *int
}

type Network_LoadBalancer_VirtualIpAddress struct {
	Entity

	Account                     *Account
	BillingItem                 *Billing_Item
	ConnectionLimit             *int
	CustomerManagedFlag         *int
	Id                          *int
	LoadBalancingMethod         *string
	LoadBalancingMethodFullName *string
	ManagedResourceFlag         *bool
	ModifyDate                  *time.Time
	Name                        *string
	Notes                       *string
	SecurityCertificateId       *int
	ServiceCount                *uint
	Services                    []Network_LoadBalancer_Service
	SourcePort                  *int
	Type                        *string
	VirtualIpAddress            *string
}

type Network_Logging_Syslog struct {
	Entity

	CreateDate           *time.Time
	DestinationIpAddress *string
	DestinationPort      *int
	EventType            *string
	Message              *string
	Protocol             *string
	SourceIpAddress      *string
	SourcePort           *int
	TotalEvents          *int
}

type Network_Media_Transcode_Account struct {
	Entity

	Account           *Account
	AccountId         *int
	CreateDate        *time.Time
	Id                *int
	ModifyDate        *time.Time
	TranscodeJobCount *uint
	TranscodeJobs     []Network_Media_Transcode_Job
}

type Network_Media_Transcode_Job struct {
	Entity

	AutoDeleteDuration  *int
	ByteIn              *int
	CreateDate          *time.Time
	History             []Network_Media_Transcode_Job_History
	HistoryCount        *uint
	Id                  *int
	InputFile           *string
	ModifyDate          *time.Time
	Name                *string
	OutputFile          *string
	TranscodeAccount    *Network_Media_Transcode_Account
	TranscodeAccountId  *int
	TranscodeJobGuid    *string
	TranscodePresetGuid *string
	TranscodePresetName *string
	TranscodeStatus     *Network_Media_Transcode_Job_Status
	TranscodeStatusId   *int
	TranscodeStatusName *string
	User                *User_Customer
	UserId              *int
	Watermark           *Container_Network_Media_Transcode_Job_Watermark
}

type Network_Media_Transcode_Job_History struct {
	Entity

	CreateDate          *time.Time
	PublicNotes         *string
	TranscodeJobId      *int
	TranscodeStatusName *string
}

type Network_Media_Transcode_Job_Status struct {
	Entity

	Description *string
	Id          *int
	Name        *string
}

type Network_Message_Delivery struct {
	Entity

	Account     *Account
	AccountId   *int
	BillingItem *Billing_Item
	CreateDate  *time.Time
	Id          *int
	ModifyDate  *time.Time
	Password    *string
	Type        *Network_Message_Delivery_Type
	TypeId      *int
	Username    *string
	Vendor      *Network_Message_Delivery_Vendor
	VendorId    *int
}

type Network_Message_Delivery_Attribute struct {
	Entity

	NetworkMessageDelivery *Network_Message_Delivery
	Value                  *string
}

type Network_Message_Delivery_Email_Sendgrid struct {
	Network_Message_Delivery

	EmailAddress *string
	SmtpAccess   *string
}

type Network_Message_Delivery_Type struct {
	Entity

	Description *string
	Id          *int
	KeyName     *string
	Name        *string
}

type Network_Message_Delivery_Vendor struct {
	Entity

	Id      *int
	KeyName *string
	Name    *string
}

type Network_Message_Queue struct {
	Entity

	Account              *Account
	AccountId            *int
	BillingItem          *Billing_Item
	CreateDate           *time.Time
	Id                   *int
	MessageQueueStatusId *int
	Name                 *string
	NodeCount            *uint
	Nodes                []Network_Message_Queue_Node
	Notes                *string
	Status               *Network_Message_Queue_Status
}

type Network_Message_Queue_Node struct {
	Entity

	AccountName          *string
	Id                   *int
	MessageQueue         *Network_Message_Queue
	MessageQueueId       *int
	MetricTrackingObject *Metric_Tracking_Object
	Name                 *string
	Notes                *string
	ServiceResource      *Network_Service_Resource
}

type Network_Message_Queue_Status struct {
	Entity

	Description *string
	Id          *int
	Name        *string
}

type Network_Monitor struct {
	Entity
}

type Network_Monitor_Version1_Incident struct {
	Entity

	Status *string
}

type Network_Monitor_Version1_Query_Host struct {
	Entity

	Arg1Value        *string
	GuestId          *int
	Hardware         *Hardware
	HardwareId       *int
	HostId           *int
	Id               *int
	IpAddress        *string
	LastResult       *Network_Monitor_Version1_Query_Result
	QueryType        *Network_Monitor_Version1_Query_Type
	QueryTypeId      *int
	ResponseAction   *Network_Monitor_Version1_Query_ResponseType
	ResponseActionId *int
	Status           *string
	WaitCycles       *int
}

type Network_Monitor_Version1_Query_Host_Stratum struct {
	Entity

	Hardware      *Hardware
	MonitorLevel  *int
	ResponseLevel *int
}

type Network_Monitor_Version1_Query_ResponseType struct {
	Entity

	ActionDescription *string
	Id                *int
	Level             *int
}

type Network_Monitor_Version1_Query_Result struct {
	Entity

	FinishTime     *time.Time
	QueryHost      *Network_Monitor_Version1_Query_Host
	ResponseStatus *int
	ResponseTime   *float64
}

type Network_Monitor_Version1_Query_Type struct {
	Entity

	ArgumentDescription *string
	Description         *string
	Id                  *int
	MonitorLevel        *int
	Name                *string
}

type Network_Pod struct {
	Entity

	BackendRouterId    *int
	BackendRouterName  *string
	Capabilities       []string
	DatacenterLongName *string
	DatacenterName     *string
	FrontendRouterId   *int
	FrontendRouterName *string
	Name               *string
}

type Network_Protection_Address struct {
	Entity

	Account              *Account
	DepartmentId         *int
	IpAddress            *string
	Location             *Location
	ManagementMethodType *string
	ModifiedUser         *User_Employee
	PrimaryRouter        *Hardware_Router
	ServiceProvider      *Service_Provider
	Subnet               *Network_Subnet
	SubnetIpAddress      *Network_Subnet_IpAddress
	TerminatedUser       *User_Employee
	Ticket               *Ticket
	TransactionCount     *uint
	Transactions         []Provisioning_Version1_Transaction
	UserDepartment       *User_Employee_Department
	UserRecord           *User_Employee
}

type Network_Regional_Internet_Registry struct {
	Entity

	Id      *int
	KeyName *string
	Name    *string
}

type Network_Security_Scanner_Request struct {
	Entity

	Account            *Account
	AccountId          *int
	CreateDate         *time.Time
	Guest              *Virtual_Guest
	GuestId            *int
	Hardware           *Hardware
	HardwareId         *int
	HostId             *int
	Id                 *int
	IpAddress          *string
	ModifyDate         *time.Time
	RequestorOwnedFlag *bool
	Status             *Network_Security_Scanner_Request_Status
	StatusId           *int
}

type Network_Security_Scanner_Request_Status struct {
	Entity

	Id   *int
	Name *string
}

type Network_Service_Health struct {
	Entity

	CreateDate *time.Time
	Location   *Location
	LocationId *int
	ModifyDate *time.Time
	Status     *Network_Service_Health_Status
	StatusId   *int
}

type Network_Service_Health_Status struct {
	Entity

	Name *string
}

type Network_Service_Resource struct {
	Entity

	ApiHost           *string
	ApiPassword       *string
	ApiPath           *string
	ApiPort           *string
	ApiProtocol       *string
	ApiUsername       *string
	ApiVersion        *string
	AttributeCount    *uint
	Attributes        []Network_Service_Resource_Attribute
	BackendIpAddress  *string
	Datacenter        *Location
	FrontendIpAddress *string
	Id                *int
	Name              *string
	NetworkDevice     *Hardware
	SshUsername       *string
	Type              *Network_Service_Resource_Type
}

type Network_Service_Resource_Attribute struct {
	Entity

	AttributeType   *Network_Service_Resource_Attribute_Type
	ServiceResource *Network_Service_Resource
	Value           *string
}

type Network_Service_Resource_Attribute_Type struct {
	Entity

	Keyname *string
}

type Network_Service_Resource_Hub struct {
	Network_Service_Resource
}

type Network_Service_Resource_Hub_Swift struct {
	Network_Service_Resource_Hub
}

type Network_Service_Resource_MonitoringHub struct {
	Network_Service_Resource

	AdnServicesIp        *string
	HubAddress           *string
	HubConnectionTimeout *string
	RobotsCount          *string
	RobotsMax            *string
}

type Network_Service_Resource_NimsoftLandingHub struct {
	Network_Service_Resource_MonitoringHub
}

type Network_Service_Resource_Type struct {
	Entity

	ServiceResourceCount *uint
	ServiceResources     []Network_Service_Resource
	Type                 *string
}

type Network_Service_Vpn_Overrides struct {
	Entity

	Id       *int
	Subnet   *Network_Subnet
	SubnetId *int
	User     *User_Customer
	UserId   *int
}

type Network_Storage struct {
	Entity

	Account                             *Account
	AccountId                           *int
	AccountPassword                     *Account_Password
	ActiveTransactionCount              *uint
	ActiveTransactions                  []Provisioning_Version1_Transaction
	AllowedHardware                     []Hardware
	AllowedHardwareCount                *uint
	AllowedIpAddressCount               *uint
	AllowedIpAddresses                  []Network_Subnet_IpAddress
	AllowedReplicationHardware          []Hardware
	AllowedReplicationHardwareCount     *uint
	AllowedReplicationIpAddressCount    *uint
	AllowedReplicationIpAddresses       []Network_Subnet_IpAddress
	AllowedReplicationSubnetCount       *uint
	AllowedReplicationSubnets           []Network_Subnet
	AllowedReplicationVirtualGuestCount *uint
	AllowedReplicationVirtualGuests     []Virtual_Guest
	AllowedSubnetCount                  *uint
	AllowedSubnets                      []Network_Subnet
	AllowedVirtualGuestCount            *uint
	AllowedVirtualGuests                []Virtual_Guest
	BillingItem                         *Billing_Item
	BillingItemCategory                 *Product_Item_Category
	BytesUsed                           *string
	CapacityGb                          *int
	CreateDate                          *time.Time
	CreationScheduleId                  *string
	CredentialCount                     *uint
	Credentials                         []Network_Storage_Credential
	DailySchedule                       *Network_Storage_Schedule
	EventCount                          *uint
	Events                              []Network_Storage_Event
	GuestId                             *int
	Hardware                            *Hardware
	HardwareId                          *int
	HostId                              *int
	HourlySchedule                      *Network_Storage_Schedule
	Id                                  *int
	Iops                                *string
	IscsiLunCount                       *uint
	IscsiLuns                           []Network_Storage
	LunId                               *string
	ManualSnapshotCount                 *uint
	ManualSnapshots                     []Network_Storage
	MetricTrackingObject                *Metric_Tracking_Object
	MountableFlag                       *string
	NasType                             *string
	Notes                               *string
	NotificationSubscriberCount         *uint
	NotificationSubscribers             []Notification_User_Subscriber
	OsType                              *Network_Storage_Iscsi_OS_Type
	OsTypeId                            *string
	ParentPartnershipCount              *uint
	ParentPartnerships                  []Network_Storage_Partnership
	ParentVolume                        *Network_Storage
	PartnershipCount                    *uint
	Partnerships                        []Network_Storage_Partnership
	Password                            *string
	PermissionsGroupCount               *uint
	PermissionsGroups                   []Network_Storage_Group
	Properties                          []Network_Storage_Property
	PropertyCount                       *uint
	ReplicatingLunCount                 *uint
	ReplicatingLuns                     []Network_Storage
	ReplicatingVolume                   *Network_Storage
	ReplicationEventCount               *uint
	ReplicationEvents                   []Network_Storage_Event
	ReplicationPartnerCount             *uint
	ReplicationPartners                 []Network_Storage
	ReplicationSchedule                 *Network_Storage_Schedule
	ReplicationStatus                   *string
	ScheduleCount                       *uint
	Schedules                           []Network_Storage_Schedule
	ServiceProviderId                   *int
	ServiceResource                     *Network_Service_Resource
	ServiceResourceBackendIpAddress     *string
	ServiceResourceName                 *string
	SnapshotCapacityGb                  *string
	SnapshotCount                       *uint
	SnapshotCreationTimestamp           *string
	SnapshotDeletionThresholdPercentage *string
	SnapshotSizeBytes                   *string
	SnapshotSpaceAvailable              *string
	Snapshots                           []Network_Storage
	StorageGroupCount                   *uint
	StorageGroups                       []Network_Storage_Group
	StorageTierLevel                    *string
	StorageType                         *Network_Storage_Type
	StorageTypeId                       *string
	TotalBytesUsed                      *string
	TotalScheduleSnapshotRetentionCount *uint
	UpgradableFlag                      *bool
	UsageNotification                   *Notification
	Username                            *string
	VendorName                          *string
	VirtualGuest                        *Virtual_Guest
	VolumeHistory                       []Network_Storage_History
	VolumeHistoryCount                  *uint
	VolumeStatus                        *string
	WebccAccount                        *Account_Password
	WeeklySchedule                      *Network_Storage_Schedule
}

type Network_Storage_Allowed_Host struct {
	Entity

	AssignedGroupCount             *uint
	AssignedGroups                 []Network_Storage_Group
	AssignedReplicationVolumeCount *uint
	AssignedReplicationVolumes     []Network_Storage
	AssignedVolumeCount            *uint
	AssignedVolumes                []Network_Storage
	Credential                     *Network_Storage_Credential
	CredentialId                   *int
	Id                             *int
	Name                           *string
	ResourceTableId                *int
	ResourceTableName              *string
}

type Network_Storage_Allowed_Host_Hardware struct {
	Network_Storage_Allowed_Host

	Resource *Hardware
}

type Network_Storage_Allowed_Host_IpAddress struct {
	Network_Storage_Allowed_Host

	Resource *Network_Subnet_IpAddress
}

type Network_Storage_Allowed_Host_Subnet struct {
	Network_Storage_Allowed_Host

	Resource *Network_Subnet
}

type Network_Storage_Allowed_Host_VirtualGuest struct {
	Network_Storage_Allowed_Host

	Resource *Virtual_Guest
}

type Network_Storage_Backup struct {
	Network_Storage

	CurrentCyclePeakUsage  *uint
	PreviousCyclePeakUsage *uint
}

type Network_Storage_Backup_Evault struct {
	Network_Storage_Backup
}

type Network_Storage_Backup_Evault_Version6 struct {
	Network_Storage_Backup_Evault

	AgentStatusCount       *uint
	AgentStatuses          []Container_Network_Storage_Evault_WebCc_AgentStatus
	BackupJobDetailCount   *uint
	BackupJobDetails       []Container_Network_Storage_Evault_WebCc_JobDetails
	PluginBillingItemCount *uint
	PluginBillingItems     []Billing_Item
	RestoreJobDetailCount  *uint
	RestoreJobDetails      []Container_Network_Storage_Evault_WebCc_JobDetails
	SoftwareComponent      *Software_Component
	TaskCount              *uint
	Tasks                  []Container_Network_Storage_Evault_Vault_Task
}

type Network_Storage_Credential struct {
	Entity

	Account                    *Account
	AccountId                  *string
	CreateDate                 *time.Time
	Id                         *int
	ModifyDate                 *time.Time
	NasCredentialTypeId        *int
	NetworkStorageAllowedHosts *Network_Storage_Allowed_Host
	Password                   *string
	Type                       *Network_Storage_Credential_Type
	Username                   *string
	VolumeCount                *uint
	Volumes                    []Network_Storage
}

type Network_Storage_Credential_Type struct {
	Entity

	CreateDate  *time.Time
	Description *string
	KeyName     *string
	ModifyDate  *time.Time
	Name        *string
}

type Network_Storage_Daily_Usage struct {
	Entity

	BytesUsed          *uint
	CdnHttpBandwidth   *uint
	CreateDate         *time.Time
	NasVolume          *Network_Storage
	NasVolumeId        *int
	PublicBandwidthOut *uint
}

type Network_Storage_Event struct {
	Entity

	CreateDate *time.Time
	Message    *string
	Schedule   *Network_Storage_Schedule
	ScheduleId *int
	TypeId     *int
	Volume     *Network_Storage
	VolumeId   *int
}

type Network_Storage_Group struct {
	Entity

	Account             *Account
	AccountId           *int
	Alias               *string
	AllowedHostCount    *uint
	AllowedHosts        []Network_Storage_Allowed_Host
	AttachedVolumeCount *uint
	AttachedVolumes     []Network_Storage
	CreateDate          *time.Time
	GroupType           *Network_Storage_Group_Type
	GroupTypeId         *int
	Id                  *int
	ModifyDate          *time.Time
	OsType              *Network_Storage_Iscsi_OS_Type
	OsTypeId            *int
	ServiceResource     *Network_Service_Resource
	ServiceResourceId   *int
}

type Network_Storage_Group_Iscsi struct {
	Network_Storage_Group
}

type Network_Storage_Group_Nfs struct {
	Network_Storage_Group
}

type Network_Storage_Group_Type struct {
	Entity

	Id      *int
	KeyName *string
	Name    *string
}

type Network_Storage_History struct {
	Entity

	Account    *Account
	CreateDate *time.Time
	NasVolume  *Network_Storage
	Notes      *string
	Password   *string
	Username   *string
}

type Network_Storage_Hub struct {
	Network_Storage

	BandwidthBillingItemCount *uint
	BandwidthBillingItems     []Billing_Item
}

type Network_Storage_Hub_Cleversafe_Account struct {
	Entity

	Account              *Account
	AccountId            *int
	BillingItem          *Billing_Item
	CancelledBillingItem *Billing_Item
	CredentialCount      *uint
	Credentials          []Network_Storage_Credential
	EventCount           *uint
	Events               []Network_Storage_Event
	Id                   *int
	MetricTrackingObject *Metric_Tracking_Object
	NasType              *string
	Notes                *string
	StorageTypeId        *int
	Username             *string
	Uuid                 *string
}

type Network_Storage_Hub_Swift struct {
	Network_Storage_Hub

	StorageNodeCount *uint
	StorageNodes     []Network_Service_Resource
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

	FailbackInProgressFlag *bool
	VolumeName             *string
}

type Network_Storage_Iscsi_EqualLogic_Version3_Snapshot struct {
	Network_Storage_Iscsi_EqualLogic_Version3

	CreationSchedule *Network_Storage_Schedule
	VolumeName       *string
}

type Network_Storage_Iscsi_OS_Type struct {
	Entity

	CreateDate  *time.Time
	Description *string
	Id          *int
	KeyName     *string
	Name        *string
}

type Network_Storage_Nas struct {
	Network_Storage

	RecentBytesUsed *Network_Storage_Daily_Usage
}

type Network_Storage_OpenStack_Object struct {
	Network_Storage

	BandwidthBillingItemCount *uint
	BandwidthBillingItems     []Billing_Item
}

type Network_Storage_Partnership struct {
	Entity

	CreateDate      *time.Time
	ModifyDate      *time.Time
	PartnerVolume   *Network_Storage
	PartnerVolumeId *int
	Type            *Network_Storage_Partnership_Type
	Volume          *Network_Storage
	VolumeId        *int
}

type Network_Storage_Partnership_Type struct {
	Entity

	Description *string
	Keyname     *string
	Name        *string
}

type Network_Storage_Property struct {
	Entity

	CreateDate *time.Time
	ModifyDate *time.Time
	Type       *Network_Storage_Property_Type
	Value      *string
	Volume     *Network_Storage
	VolumeId   *int
}

type Network_Storage_Property_Type struct {
	Entity

	Description *string
	Keyname     *string
	Name        *string
}

type Network_Storage_Replicant struct {
	Network_Storage

	FailbackInProgressFlag *string
	VolumeName             *string
}

type Network_Storage_Schedule struct {
	Entity

	Active               *int
	CreateDate           *time.Time
	DayOfMonth           *string
	DayOfWeek            *string
	EventCount           *uint
	Events               []Network_Storage_Event
	Hour                 *string
	Id                   *int
	Minute               *string
	ModifyDate           *time.Time
	MonthOfYear          *string
	Name                 *string
	Partnership          *Network_Storage_Partnership
	PartnershipId        *int
	Properties           []Network_Storage_Schedule_Property
	PropertyCount        *uint
	ReplicaSnapshotCount *uint
	ReplicaSnapshots     []Network_Storage
	RetentionCount       *string
	SnapshotCount        *uint
	Snapshots            []Network_Storage
	Type                 *Network_Storage_Schedule_Type
	TypeId               *int
	Volume               *Network_Storage
	VolumeId             *int
}

type Network_Storage_Schedule_Property struct {
	Entity

	CreateDate *time.Time
	Id         *int
	ModifyDate *time.Time
	Schedule   *Network_Storage_Schedule
	Type       *Network_Storage_Schedule_Property_Type
	TypeId     *int
	Value      *string
}

type Network_Storage_Schedule_Property_Type struct {
	Entity

	Description *string
	Id          *int
	Keyname     *string
	Name        *string
	NasType     *string
}

type Network_Storage_Schedule_Type struct {
	Entity

	Id      *int
	Keyname *string
	Name    *string
}

type Network_Storage_Snapshot struct {
	Network_Storage

	CreationSchedule *Network_Storage_Schedule
	VolumeName       *string
}

type Network_Storage_Type struct {
	Entity

	Description *string
	Id          *int
	KeyName     *string
	VolumeCount *uint
	Volumes     []Network_Storage
}

type Network_Subnet struct {
	Entity

	Account                           *Account
	ActiveRegistration                *Network_Subnet_Registration
	ActiveSwipTransaction             *Network_Subnet_Swip_Transaction
	ActiveTransaction                 *Provisioning_Version1_Transaction
	AddressSpace                      *string
	AllowedHost                       *Network_Storage_Allowed_Host
	AllowedNetworkStorage             []Network_Storage
	AllowedNetworkStorageCount        *uint
	AllowedNetworkStorageReplicaCount *uint
	AllowedNetworkStorageReplicas     []Network_Storage
	BillingItem                       *Billing_Item
	BoundDescendantCount              *uint
	BoundDescendants                  []Network_Subnet
	BoundRouterCount                  *uint
	BoundRouterFlag                   *bool
	BoundRouters                      []Hardware
	BroadcastAddress                  *string
	Children                          []Network_Subnet
	ChildrenCount                     *uint
	Cidr                              *int
	Datacenter                        *Location_Datacenter
	DescendantCount                   *uint
	Descendants                       []Network_Subnet
	DisplayLabel                      *string
	EndPointIpAddress                 *Network_Subnet_IpAddress
	Gateway                           *string
	GlobalIpRecord                    *Network_Subnet_IpAddress_Global
	Hardware                          []Hardware
	HardwareCount                     *uint
	Id                                *int
	IpAddressCount                    *uint
	IpAddresses                       []Network_Subnet_IpAddress
	IsCustomerOwned                   *bool
	IsCustomerRoutable                *bool
	ModifyDate                        *time.Time
	Netmask                           *string
	NetworkComponent                  *Network_Component
	NetworkComponentFirewall          *Network_Component_Firewall
	NetworkIdentifier                 *string
	NetworkProtectionAddressCount     *uint
	NetworkProtectionAddresses        []Network_Protection_Address
	NetworkTunnelContextCount         *uint
	NetworkTunnelContexts             []Network_Tunnel_Module_Context
	NetworkVlan                       *Network_Vlan
	NetworkVlanId                     *int
	Note                              *string
	PodName                           *string
	ProtectedIpAddressCount           *uint
	ProtectedIpAddresses              []Network_Subnet_IpAddress
	RegionalInternetRegistry          *Network_Regional_Internet_Registry
	RegistrationCount                 *uint
	Registrations                     []Network_Subnet_Registration
	ResourceGroupCount                *uint
	ResourceGroups                    []Resource_Group
	ReverseDomain                     *Dns_Domain
	RoleKeyName                       *string
	RoleName                          *string
	RoutingTypeKeyName                *string
	RoutingTypeName                   *string
	SortOrder                         *string
	SubnetType                        *string
	SwipTransaction                   []Network_Subnet_Swip_Transaction
	SwipTransactionCount              *uint
	TotalIpAddresses                  *float64
	UnboundDescendantCount            *uint
	UnboundDescendants                []Network_Subnet
	UsableIpAddressCount              *float64
	Version                           *int
	VirtualGuestCount                 *uint
	VirtualGuests                     []Virtual_Guest
}

type Network_Subnet_IpAddress struct {
	Entity

	AllowedHost                                      *Network_Storage_Allowed_Host
	AllowedNetworkStorage                            []Network_Storage
	AllowedNetworkStorageCount                       *uint
	AllowedNetworkStorageReplicaCount                *uint
	AllowedNetworkStorageReplicas                    []Network_Storage
	ApplicationDeliveryController                    *Network_Application_Delivery_Controller
	ContextTunnelTranslationCount                    *uint
	ContextTunnelTranslations                        []Network_Tunnel_Module_Context_Address_Translation
	EndpointSubnetCount                              *uint
	EndpointSubnets                                  []Network_Subnet
	GuestNetworkComponent                            *Virtual_Guest_Network_Component
	GuestNetworkComponentBinding                     *Virtual_Guest_Network_Component_IpAddress
	Hardware                                         *Hardware
	Id                                               *int
	IpAddress                                        *string
	IsBroadcast                                      *bool
	IsGateway                                        *bool
	IsNetwork                                        *bool
	IsReserved                                       *bool
	NetworkComponent                                 *Network_Component
	Note                                             *string
	PrivateNetworkGateway                            *Network_Gateway
	ProtectionAddress                                []Network_Protection_Address
	ProtectionAddressCount                           *uint
	PublicNetworkGateway                             *Network_Gateway
	RemoteManagementNetworkComponent                 *Network_Component
	Subnet                                           *Network_Subnet
	SubnetId                                         *int
	SyslogEventsOneDay                               []Network_Logging_Syslog
	SyslogEventsOneDayCount                          *uint
	SyslogEventsSevenDayCount                        *uint
	SyslogEventsSevenDays                            []Network_Logging_Syslog
	TopTenSyslogEventsByDestinationPortOneDay        []Network_Logging_Syslog
	TopTenSyslogEventsByDestinationPortOneDayCount   *uint
	TopTenSyslogEventsByDestinationPortSevenDayCount *uint
	TopTenSyslogEventsByDestinationPortSevenDays     []Network_Logging_Syslog
	TopTenSyslogEventsByProtocolsOneDay              []Network_Logging_Syslog
	TopTenSyslogEventsByProtocolsOneDayCount         *uint
	TopTenSyslogEventsByProtocolsSevenDayCount       *uint
	TopTenSyslogEventsByProtocolsSevenDays           []Network_Logging_Syslog
	TopTenSyslogEventsBySourceIpOneDay               []Network_Logging_Syslog
	TopTenSyslogEventsBySourceIpOneDayCount          *uint
	TopTenSyslogEventsBySourceIpSevenDayCount        *uint
	TopTenSyslogEventsBySourceIpSevenDays            []Network_Logging_Syslog
	TopTenSyslogEventsBySourcePortOneDay             []Network_Logging_Syslog
	TopTenSyslogEventsBySourcePortOneDayCount        *uint
	TopTenSyslogEventsBySourcePortSevenDayCount      *uint
	TopTenSyslogEventsBySourcePortSevenDays          []Network_Logging_Syslog
	VirtualGuest                                     *Virtual_Guest
	VirtualLicenseCount                              *uint
	VirtualLicenses                                  []Software_VirtualLicense
}

type Network_Subnet_IpAddress_Global struct {
	Entity

	Account                *Account
	ActiveTransaction      *Provisioning_Version1_Transaction
	BillingItem            *Billing_Item_Network_Subnet_IpAddress_Global
	Description            *int
	DestinationIpAddress   *Network_Subnet_IpAddress
	DestinationIpAddressId *int
	Id                     *int
	IpAddress              *Network_Subnet_IpAddress
	IpAddressId            *int
	TypeId                 *int
}

type Network_Subnet_IpAddress_Version6 struct {
	Network_Subnet_IpAddress

	PublicVersion6NetworkGateway *Network_Gateway
}

type Network_Subnet_Registration struct {
	Entity

	Account                          *Account
	AccountId                        *int
	Cidr                             *int
	CreateDate                       *time.Time
	DetailReferenceCount             *uint
	DetailReferences                 []Network_Subnet_Registration_Details
	EventCount                       *uint
	Events                           []Network_Subnet_Registration_Event
	Id                               *int
	ModifyDate                       *time.Time
	NetworkDetail                    *Account_Regional_Registry_Detail
	NetworkHandle                    *string
	NetworkIdentifier                *string
	PersonDetail                     *Account_Regional_Registry_Detail
	RegionalInternetRegistry         *Network_Regional_Internet_Registry
	RegionalInternetRegistryHandle   *Account_Rwhois_Handle
	RegionalInternetRegistryHandleId *int
	RegionalInternetRegistryId       *int
	Status                           *Network_Subnet_Registration_Status
	StatusId                         *int
	Subnet                           *Network_Subnet
}

type Network_Subnet_Registration_Apnic struct {
	Network_Subnet_Registration
}

type Network_Subnet_Registration_Arin struct {
	Network_Subnet_Registration
}

type Network_Subnet_Registration_Details struct {
	Entity

	CreateDate     *time.Time
	Detail         *Account_Regional_Registry_Detail
	DetailId       *int
	Id             *int
	ModifyDate     *time.Time
	Registration   *Network_Subnet_Registration
	RegistrationId *int
}

type Network_Subnet_Registration_Event struct {
	Entity

	CreateDate     *time.Time
	Id             *int
	Message        *string
	ModifyDate     *time.Time
	Registration   *Network_Subnet_Registration
	RegistrationId *int
	Type           *Network_Subnet_Registration_Event_Type
	TypeId         *int
}

type Network_Subnet_Registration_Event_Type struct {
	Entity

	CreateDate *time.Time
	Id         *int
	KeyName    *string
	ModifyDate *time.Time
	Name       *string
}

type Network_Subnet_Registration_Ripe struct {
	Network_Subnet_Registration
}

type Network_Subnet_Registration_Status struct {
	Entity

	CreateDate *time.Time
	Id         *int
	KeyName    *string
	ModifyDate *time.Time
	Name       *string
}

type Network_Subnet_Rwhois_Data struct {
	Entity

	AbuseEmail           *string
	Account              *Account
	AccountId            *int
	Address1             *string
	Address2             *string
	City                 *string
	CompanyName          *string
	Country              *string
	CreateDate           *time.Time
	FirstName            *string
	Id                   *int
	LastName             *string
	ModifyDate           *time.Time
	PostalCode           *string
	PrivateResidenceFlag *bool
	State                *string
}

type Network_Subnet_Swip_Transaction struct {
	Entity

	Account    *Account
	Id         *int
	StatusName *string
	Subnet     *Network_Subnet
	SubnetId   *int
}

type Network_TippingPointReporting struct {
	Entity
}

type Network_Tunnel_Module_Context struct {
	Entity

	Account                        *Account
	AccountId                      *int
	ActiveTransaction              *Provisioning_Version1_Transaction
	AddressTranslationCount        *uint
	AddressTranslations            []Network_Tunnel_Module_Context_Address_Translation
	AdvancedConfigurationFlag      *int
	AllAvailableServiceSubnetCount *uint
	AllAvailableServiceSubnets     []Network_Subnet
	BillingItem                    *Billing_Item
	CreateDate                     *time.Time
	CustomerPeerIpAddress          *string
	CustomerSubnetCount            *uint
	CustomerSubnets                []Network_Customer_Subnet
	Datacenter                     *Location
	FriendlyName                   *string
	Id                             *int
	InternalPeerIpAddress          *string
	InternalSubnetCount            *uint
	InternalSubnets                []Network_Subnet
	ModifyDate                     *time.Time
	Name                           *string
	PhaseOneAuthentication         *string
	PhaseOneDiffieHellmanGroup     *int
	PhaseOneEncryption             *string
	PhaseOneKeylife                *int
	PhaseTwoAuthentication         *string
	PhaseTwoDiffieHellmanGroup     *int
	PhaseTwoEncryption             *string
	PhaseTwoKeylife                *int
	PhaseTwoPerfectForwardSecrecy  *int
	PresharedKey                   *string
	ServiceSubnetCount             *uint
	ServiceSubnets                 []Network_Subnet
	StaticRouteSubnetCount         *uint
	StaticRouteSubnets             []Network_Subnet
	TransactionHistory             []Provisioning_Version1_Transaction
	TransactionHistoryCount        *uint
}

type Network_Tunnel_Module_Context_Address_Translation struct {
	Entity

	CustomerIpAddress       *string
	CustomerIpAddressId     *int
	CustomerIpAddressRecord *Network_Customer_Subnet_IpAddress
	Id                      *int
	InternalIpAddress       *string
	InternalIpAddressId     *int
	InternalIpAddressRecord *Network_Subnet_IpAddress
	NetworkTunnelContext    *Network_Tunnel_Module_Context
	NetworkTunnelContextId  *int
	Notes                   *string
}

type Network_Vlan struct {
	Entity

	Account                            *Account
	AccountId                          *int
	AdditionalPrimarySubnetCount       *uint
	AdditionalPrimarySubnets           []Network_Subnet
	AttachedNetworkGateway             *Network_Gateway
	AttachedNetworkGatewayFlag         *bool
	AttachedNetworkGatewayVlan         *Network_Gateway_Vlan
	BillingItem                        *Billing_Item
	DedicatedFirewallFlag              *int
	ExtensionRouter                    *Hardware_Router
	FirewallGuestNetworkComponentCount *uint
	FirewallGuestNetworkComponents     []Network_Component_Firewall
	FirewallInterfaceCount             *uint
	FirewallInterfaces                 []Network_Firewall_Module_Context_Interface
	FirewallNetworkComponentCount      *uint
	FirewallNetworkComponents          []Network_Component_Firewall
	FirewallRuleCount                  *uint
	FirewallRules                      []Network_Vlan_Firewall_Rule
	GuestNetworkComponentCount         *uint
	GuestNetworkComponents             []Virtual_Guest_Network_Component
	Hardware                           []Hardware
	HardwareCount                      *uint
	HighAvailabilityFirewallFlag       *bool
	Id                                 *int
	LocalDiskStorageCapabilityFlag     *bool
	ModifyDate                         *time.Time
	Name                               *string
	Network                            *Network
	NetworkComponentCount              *uint
	NetworkComponentTrunkCount         *uint
	NetworkComponentTrunks             []Network_Component_Network_Vlan_Trunk
	NetworkComponents                  []Network_Component
	NetworkSpace                       *string
	NetworkVlanFirewall                *Network_Vlan_Firewall
	Note                               *string
	PrimaryRouter                      *Hardware_Router
	PrimarySubnet                      *Network_Subnet
	PrimarySubnetCount                 *uint
	PrimarySubnetId                    *int
	PrimarySubnetVersion6              *Network_Subnet
	PrimarySubnets                     []Network_Subnet
	PrivateNetworkGatewayCount         *uint
	PrivateNetworkGateways             []Network_Gateway
	ProtectedIpAddressCount            *uint
	ProtectedIpAddresses               []Network_Subnet_IpAddress
	PublicNetworkGatewayCount          *uint
	PublicNetworkGateways              []Network_Gateway
	ResourceGroupCount                 *uint
	ResourceGroupMember                []Resource_Group_Member
	ResourceGroupMemberCount           *uint
	ResourceGroups                     []Resource_Group
	SanStorageCapabilityFlag           *bool
	ScaleVlanCount                     *uint
	ScaleVlans                         []Scale_Network_Vlan
	SecondaryRouter                    *Hardware
	SecondarySubnetCount               *uint
	SecondarySubnets                   []Network_Subnet
	SubnetCount                        *uint
	Subnets                            []Network_Subnet
	TagReferenceCount                  *uint
	TagReferences                      []Tag_Reference
	TotalPrimaryIpAddressCount         *uint
	Type                               *Network_Vlan_Type
	VirtualGuestCount                  *uint
	VirtualGuests                      []Virtual_Guest
	VlanNumber                         *int
}

type Network_Vlan_Firewall struct {
	Entity

	AdministrativeBypassFlag          *string
	BillingItem                       *Billing_Item
	CustomerManagedFlag               *bool
	Datacenter                        *Location
	FirewallType                      *string
	FullyQualifiedDomainName          *string
	Id                                *int
	ManagementCredentials             *Software_Component_Password
	NetworkFirewallUpdateRequestCount *uint
	NetworkFirewallUpdateRequests     []Network_Firewall_Update_Request
	NetworkVlan                       *Network_Vlan
	NetworkVlanCount                  *uint
	NetworkVlans                      []Network_Vlan
	PrimaryIpAddress                  *string
	RuleCount                         *uint
	Rules                             []Network_Vlan_Firewall_Rule
	TagReferenceCount                 *uint
	TagReferences                     []Tag_Reference
}

type Network_Vlan_Firewall_Rule struct {
	Entity

	Action                    *string
	DestinationIpAddress      *string
	DestinationIpCidr         *int
	DestinationIpSubnetMask   *string
	DestinationPortRangeEnd   *int
	DestinationPortRangeStart *int
	Id                        *int
	NetworkComponentFirewall  *Network_Component_Firewall
	Notes                     *string
	OrderValue                *int
	Protocol                  *string
	SourceIpAddress           *string
	SourceIpCidr              *int
	SourceIpSubnetMask        *string
	Status                    *string
	Version                   *int
}

type Network_Vlan_Type struct {
	Entity

	Description *string
	Id          *int
	KeyName     *string
	Name        *string
}

type Notification struct {
	Entity

	Id                      *int
	KeyName                 *string
	Name                    *string
	PreferenceCount         *uint
	Preferences             []Notification_Preference
	RequiredPreferenceCount *uint
	RequiredPreferences     []Notification_Preference
}

type Notification_Delivery_Method struct {
	Entity

	Active      *int
	Description *string
	Id          *int
	KeyName     *string
	Name        *string
}

type Notification_Mobile struct {
	Notification
}

type Notification_Occurrence_Account struct {
	Entity

	Account                     *Account
	Active                      *int
	LastNotificationUpdate      *Notification_Occurrence_Update
	NotificationOccurrenceEvent *Notification_Occurrence_Event
}

type Notification_Occurrence_Event struct {
	Entity

	AcknowledgedFlag                *bool
	AttachmentCount                 *uint
	Attachments                     []Notification_Occurrence_Event_Attachment
	EndDate                         *time.Time
	FirstUpdate                     *Notification_Occurrence_Update
	Id                              *int
	ImpactedAccountCount            *uint
	ImpactedAccounts                []Notification_Occurrence_Account
	ImpactedResourceCount           *uint
	ImpactedResources               []Notification_Occurrence_Resource
	ImpactedUserCount               *uint
	ImpactedUsers                   []Notification_Occurrence_User
	LastImpactedUserCount           *int
	LastUpdate                      *Notification_Occurrence_Update
	ModifyDate                      *time.Time
	NotificationOccurrenceEventType *Notification_Occurrence_Event_Type
	RecoveryTime                    *int
	StartDate                       *time.Time
	StatusCode                      *Notification_Occurrence_Status_Code
	Subject                         *string
	Summary                         *string
	SystemTicketId                  *int
	UpdateCount                     *uint
	Updates                         []Notification_Occurrence_Update
}

type Notification_Occurrence_Event_Attachment struct {
	Entity

	CreateDate                    *time.Time
	FileName                      *string
	FileSize                      *string
	Id                            *int
	NotificationOccurrenceEvent   *Notification_Occurrence_Event
	NotificationOccurrenceEventId *int
}

type Notification_Occurrence_Event_Type struct {
	Entity

	KeyName *string
}

type Notification_Occurrence_Resource struct {
	Entity

	Active                        *int
	FilterLabel                   *string
	NotificationOccurrenceEvent   *Notification_Occurrence_Event
	NotificationOccurrenceEventId *int
	Resource                      *Entity
	ResourceAccountId             *int
	ResourceName                  *string
	ResourceTableId               *int
}

type Notification_Occurrence_Resource_Hardware struct {
	Notification_Occurrence_Resource

	Hostname     *string
	PrivateIp    *string
	PublicIp     *string
	ResourceType *string
}

type Notification_Occurrence_Resource_Network_Application_Delivery_Controller struct {
	Notification_Occurrence_Resource

	Hostname     *string
	PrivateIp    *string
	PublicIp     *string
	ResourceType *string
}

type Notification_Occurrence_Resource_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress struct {
	Notification_Occurrence_Resource

	Hostname     *string
	PublicIp     *string
	ResourceType *string
}

type Notification_Occurrence_Resource_Network_Storage_Iscsi_EqualLogic struct {
	Notification_Occurrence_Resource

	Hostname     *string
	PrivateIp    *string
	ResourceType *string
}

type Notification_Occurrence_Resource_Network_Storage_Iscsi_NetApp struct {
	Notification_Occurrence_Resource

	Hostname     *string
	PrivateIp    *string
	ResourceType *string
}

type Notification_Occurrence_Resource_Network_Storage_Lockbox struct {
	Notification_Occurrence_Resource

	Hostname     *string
	PrivateIp    *string
	ResourceType *string
}

type Notification_Occurrence_Resource_Network_Storage_Nas struct {
	Notification_Occurrence_Resource

	Hostname     *string
	PrivateIp    *string
	ResourceType *string
}

type Notification_Occurrence_Resource_Network_Storage_NetApp_Volume struct {
	Notification_Occurrence_Resource

	Hostname     *string
	PrivateIp    *string
	ResourceType *string
}

type Notification_Occurrence_Resource_Virtual struct {
	Notification_Occurrence_Resource

	Hostname     *string
	PrivateIp    *string
	PublicIp     *string
	ResourceType *string
}

type Notification_Occurrence_Status_Code struct {
	Entity

	Description *string
	KeyName     *string
	Name        *string
}

type Notification_Occurrence_Update struct {
	Entity

	Contents                    *string
	CreateDate                  *time.Time
	Employee                    *User_Employee
	EndDate                     *time.Time
	NotificationOccurrenceEvent *Notification_Occurrence_Event
	StartDate                   *time.Time
}

type Notification_Occurrence_User struct {
	Entity

	AcknowledgedFlag            *int
	Active                      *int
	Id                          *int
	ImpactedResourceCount       *uint
	ImpactedResources           []Notification_Occurrence_Resource
	NotificationOccurrenceEvent *Notification_Occurrence_Event
	User                        *User_Customer
	UsrRecordId                 *int
}

type Notification_Preference struct {
	Entity

	Description  *string
	Id           *int
	KeyName      *string
	MaximumValue *string
	MinimumValue *string
	Name         *string
	Units        *string
	Value        *string
}

type Notification_Subscriber struct {
	Entity

	Active                               *int
	CreateDate                           *time.Time
	DeliveryMethodCount                  *uint
	DeliveryMethods                      []Notification_Subscriber_Delivery_Method
	Id                                   *int
	ModifyDate                           *time.Time
	Notification                         *Notification
	NotificationId                       *int
	NotificationSubscriberTypeId         *int
	NotificationSubscriberTypeResourceId *int
}

type Notification_Subscriber_Customer struct {
	Notification_Subscriber

	SubscriberRecord *User_Customer
}

type Notification_Subscriber_Delivery_Method struct {
	Entity

	Active                       *int
	CreateDate                   *time.Time
	ModifyDate                   *time.Time
	NotificationDeliveryMethod   *Notification_Delivery_Method
	NotificationDeliveryMethodId *int
	NotificationSubscriber       *Notification_Subscriber
	NotificationSubscriberId     *int
}

type Notification_User_Subscriber struct {
	Entity

	Active                 *int
	DeliveryMethodCount    *uint
	DeliveryMethods        []Notification_Delivery_Method
	Id                     *int
	Notification           *Notification
	NotificationId         *int
	PreferenceCount        *uint
	Preferences            []Notification_User_Subscriber_Preference
	PreferencesDetailCount *uint
	PreferencesDetails     []Notification_Preference
	ResourceRecord         *Notification_User_Subscriber_Resource
	UserRecord             *User_Customer
	UserRecordId           *int
}

type Notification_User_Subscriber_Billing struct {
	Notification_User_Subscriber
}

type Notification_User_Subscriber_Delivery_Method struct {
	Entity

	Active                       *int
	DeliveryMethod               *Notification_Delivery_Method
	NotificationMethodId         *int
	NotificationUserSubscriber   *Notification_User_Subscriber
	NotificationUserSubscriberId *int
}

type Notification_User_Subscriber_Mobile struct {
	Notification_User_Subscriber
}

type Notification_User_Subscriber_Preference struct {
	Entity

	DefaultPreference            *Notification_Preference
	Id                           *int
	NotificationPreferenceId     *int
	NotificationUserSubscriber   *Notification_User_Subscriber
	NotificationUserSubscriberId *int
	Value                        *string
}

type Notification_User_Subscriber_Resource struct {
	Entity

	NotificationUserSubscriber   *Notification_User_Subscriber
	NotificationUserSubscriberId *int
	ResourceTableId              *int
}

type Product_Catalog struct {
	Entity

	BrandCount   *uint
	Brands       []Brand
	PackageCount *uint
	Packages     []Product_Package
	PriceCount   *uint
	Prices       []Product_Item_Price
	ProductCount *uint
	Products     []Product_Item
}

type Product_Catalog_Item_Price struct {
	Entity

	Catalog    *Product_Catalog
	CatalogId  *int
	CreateDate *time.Time
	ModifyDate *time.Time
	Price      *Product_Item_Price
	PriceId    *int
}

type Product_Item struct {
	Entity

	ActivePresaleEventCount         *uint
	ActivePresaleEvents             []Sales_Presale_Event
	ActiveUsagePriceCount           *uint
	ActiveUsagePrices               []Product_Item_Price
	AttributeCount                  *uint
	Attributes                      []Product_Item_Attribute
	AvailabilityAttributeCount      *uint
	AvailabilityAttributes          []Product_Item_Attribute
	BillingType                     *string
	Bundle                          []Product_Item_Bundles
	BundleCount                     *uint
	Capacity                        *float64
	CapacityRestrictedProductFlag   *bool
	Categories                      []Product_Item_Category
	CategoryCount                   *uint
	ConfigurationTemplateCount      *uint
	ConfigurationTemplates          []Configuration_Template
	ConflictCount                   *uint
	Conflicts                       []Product_Item_Resource_Conflict
	CoreRestrictedItemFlag          *bool
	Description                     *string
	DowngradeItem                   *Product_Item
	DowngradeItemCount              *uint
	DowngradeItems                  []Product_Item
	GlobalCategoryConflictCount     *uint
	GlobalCategoryConflicts         []Product_Item_Resource_Conflict
	HardwareGenericComponentModel   *Hardware_Component_Model_Generic
	HideFromPortalFlag              *bool
	Id                              *int
	Inventory                       []Product_Package_Inventory
	InventoryCount                  *uint
	IsEngineeredServerProduct       *bool
	ItemCategory                    *Product_Item_Category
	ItemTaxCategoryId               *int
	KeyName                         *string
	LocationConflictCount           *uint
	LocationConflicts               []Product_Item_Resource_Conflict
	LongDescription                 *string
	ObjectStorageItemFlag           *bool
	PackageCount                    *uint
	Packages                        []Product_Package
	PhysicalCoreCapacity            *string
	PresaleEventCount               *uint
	PresaleEvents                   []Sales_Presale_Event
	PriceCount                      *uint
	Prices                          []Product_Item_Price
	RequirementCount                *uint
	Requirements                    []Product_Item_Requirement
	SoftwareDescription             *Software_Description
	SoftwareDescriptionId           *int
	TaxCategory                     *Product_Item_Tax_Category
	ThirdPartyPolicyAssignmentCount *uint
	ThirdPartyPolicyAssignments     []Product_Item_Policy_Assignment
	ThirdPartySupportVendor         *string
	TotalPhysicalCoreCapacity       *int
	TotalPhysicalCoreCount          *int
	TotalProcessorCapacity          *int
	Units                           *string
	UpgradeItem                     *Product_Item
	UpgradeItemCount                *uint
	UpgradeItemId                   *int
	UpgradeItems                    []Product_Item
}

type Product_Item_Attribute struct {
	Entity

	AttributeType        *Product_Item_Attribute_Type
	AttributeTypeKeyName *string
	Id                   *int
	Item                 *Product_Item
	ItemAttributeTypeId  *int
	ItemId               *int
	Value                *string
}

type Product_Item_Attribute_Type struct {
	Entity

	KeyName *string
	Name    *string
}

type Product_Item_Billing_Type struct {
	Entity

	Name *string
}

type Product_Item_Bundles struct {
	Entity

	BundleItem   *Product_Item
	BundleItemId *int
	Category     *Product_Item_Category
	Id           *int
	ItemPrice    *Product_Item_Price
	ItemPriceId  *int
}

type Product_Item_Category struct {
	Entity

	BillingItemCount          *uint
	BillingItems              []Billing_Item
	CategoryCode              *string
	Group                     *Product_Item_Category_Group
	GroupCount                *uint
	Groups                    []Product_Package_Item_Category_Group
	Id                        *int
	Name                      *string
	OrderOptionCount          *uint
	OrderOptions              []Product_Item_Category_Order_Option_Type
	PackageConfigurationCount *uint
	PackageConfigurations     []Product_Package_Order_Configuration
	PresetConfigurationCount  *uint
	PresetConfigurations      []Product_Package_Preset_Configuration
	QuantityLimit             *int
	QuestionCount             *uint
	QuestionReferenceCount    *uint
	QuestionReferences        []Product_Item_Category_Question_Xref
	Questions                 []Product_Item_Category_Question
}

type Product_Item_Category_Group struct {
	Entity

	Id   *int
	Name *string
}

type Product_Item_Category_Order_Option_Type struct {
	Entity

	Description *string
	Id          *int
	Keyname     *string
	Name        *string
	Value       *string
}

type Product_Item_Category_Question struct {
	Entity

	AnswerValueExpression      *string
	Description                *string
	FieldType                  *Product_Item_Category_Question_Field_Type
	FieldTypeId                *int
	Id                         *int
	ItemCategoryReferenceCount *uint
	ItemCategoryReferences     []Product_Item_Category_Question_Xref
	KeyName                    *string
	Question                   *string
	ValueExample               *string
}

type Product_Item_Category_Question_Field_Type struct {
	Entity

	Id      *int
	KeyName *string
	Name    *string
}

type Product_Item_Category_Question_Xref struct {
	Entity

	Id             *int
	ItemCategory   *Product_Item_Category
	ItemCategoryId *int
	LocationId     *int
	Question       *Product_Item_Category_Question
	QuestionId     *int
	Required       *bool
}

type Product_Item_Link_ThePlanet struct {
	Entity

	Item            *Product_Item
	ServiceProvider *Service_Provider
}

type Product_Item_Policy_Assignment struct {
	Entity

	Id         *int
	PolicyName *string
	Product    *Product_Item
	ProductId  *int
}

type Product_Item_Price struct {
	Entity

	AccountRestrictionCount    *uint
	AccountRestrictions        []Product_Item_Price_Account_Restriction
	AttributeCount             *uint
	Attributes                 []Product_Item_Price_Attribute
	BigDataOsJournalDiskFlag   *bool
	BundleReferenceCount       *uint
	BundleReferences           []Product_Item_Bundles
	CapacityRestrictionMaximum *string
	CapacityRestrictionMinimum *string
	CapacityRestrictionType    *string
	Categories                 []Product_Item_Category
	CategoryCount              *uint
	CurrentPriceFlag           *bool
	DefinedSoftwareLicenseFlag *bool
	HourlyRecurringFee         *float64
	Id                         *int
	Inventory                  []Product_Package_Inventory
	InventoryCount             *uint
	Item                       *Product_Item
	ItemId                     *int
	LaborFee                   *float64
	LocationGroupId            *int
	OnSaleFlag                 *bool
	OneTimeFee                 *float64
	OneTimeFeeTax              *float64
	OrderOptions               []Product_Item_Category_Order_Option_Type
	OrderPremiumCount          *uint
	OrderPremiums              []Product_Item_Price_Premium
	PackageCount               *uint
	PackageReferenceCount      *uint
	PackageReferences          []Product_Package_Item_Prices
	Packages                   []Product_Package
	PresetConfigurationCount   *uint
	PresetConfigurations       []Product_Package_Preset_Configuration
	PricingLocationGroup       *Location_Group_Pricing
	ProratedRecurringFee       *float64
	ProratedRecurringFeeTax    *float64
	Quantity                   *int
	RecurringFee               *float64
	RecurringFeeTax            *float64
	RequiredCoreCount          *int
	SetupFee                   *float64
	Sort                       *int
	UsageRate                  *float64
}

type Product_Item_Price_Account_Restriction struct {
	Entity

	Account     *Account
	AccountId   *int
	Id          *int
	ItemPrice   *Product_Item_Price
	ItemPriceId *int
}

type Product_Item_Price_Attribute struct {
	Entity

	Id                       *int
	ItemPrice                *Product_Item_Price
	ItemPriceAttributeType   *Product_Item_Price_Attribute_Type
	ItemPriceAttributeTypeId *int
	ItemPriceId              *int
	Value                    *string
}

type Product_Item_Price_Attribute_Type struct {
	Entity

	Id      *int
	Keyname *string
}

type Product_Item_Price_Premium struct {
	Entity

	HourlyModifier  *float64
	ItemPrice       *Product_Item_Price
	ItemPriceId     *int
	Location        *Location
	LocationId      *int
	MonthlyModifier *float64
	Package         *Product_Package
	PackageId       *int
}

type Product_Item_Requirement struct {
	Entity

	Id             *int
	Item           *Product_Item
	ItemId         *int
	Message        *string
	Product        *Product_Item
	RequiredItemId *int
}

type Product_Item_Resource_Conflict struct {
	Entity

	Item            *Product_Item
	ItemId          *int
	Message         *string
	Package         *Product_Package
	PackageId       *int
	ResourceTableId *int
}

type Product_Item_Resource_Conflict_Item struct {
	Product_Item_Resource_Conflict

	Resource *Product_Item
}

type Product_Item_Resource_Conflict_Item_Category struct {
	Product_Item_Resource_Conflict

	Resource *Product_Item_Category
}

type Product_Item_Resource_Conflict_Location struct {
	Product_Item_Resource_Conflict

	Resource *Location
}

type Product_Item_Tax_Category struct {
	Entity

	Id         *int
	ItemCount  *uint
	Items      []Product_Item
	Name       *string
	StatusFlag *int
}

type Product_Order struct {
	Entity
}

type Product_Package struct {
	Entity

	AccountRestrictedCategories     []Product_Item_Category
	AccountRestrictedCategoryCount  *uint
	AccountRestrictedPricesFlag     *bool
	ActivePresetCount               *uint
	ActivePresets                   []Product_Package_Preset
	ActiveRamItemCount              *uint
	ActiveRamItems                  []Product_Item
	ActiveServerItemCount           *uint
	ActiveServerItems               []Product_Item
	ActiveSoftwareItemCount         *uint
	ActiveSoftwareItems             []Product_Item
	ActiveUsagePriceCount           *uint
	ActiveUsagePrices               []Product_Item_Price
	AdditionalServiceFlag           *bool
	AttributeCount                  *uint
	Attributes                      []Product_Package_Attribute
	AvailableLocationCount          *uint
	AvailableLocations              []Product_Package_Locations
	AvailableStorageUnits           *uint
	Categories                      []Product_Item_Category
	Configuration                   []Product_Package_Order_Configuration
	ConfigurationCount              *uint
	DefaultRamItemCount             *uint
	DefaultRamItems                 []Product_Item
	DeploymentCount                 *uint
	DeploymentNodeType              *string
	DeploymentPackageCount          *uint
	DeploymentPackages              []Product_Package
	DeploymentType                  *string
	Deployments                     []Product_Package
	Description                     *string
	DisallowCustomDiskPartitions    *bool
	FirstOrderStep                  *Product_Package_Order_Step
	FirstOrderStepId                *int
	GatewayApplianceFlag            *bool
	GpuFlag                         *bool
	HourlyBillingAvailableFlag      *bool
	Id                              *int
	IsActive                        *int
	ItemConflictCount               *uint
	ItemConflicts                   []Product_Item_Resource_Conflict
	ItemCount                       *uint
	ItemLocationConflictCount       *uint
	ItemLocationConflicts           []Product_Item_Resource_Conflict
	ItemPriceCount                  *uint
	ItemPriceReferenceCount         *uint
	ItemPriceReferences             []Product_Package_Item_Prices
	ItemPrices                      []Product_Item_Price
	Items                           []Product_Item
	KeyName                         *string
	LocationCount                   *uint
	Locations                       []Location
	LowestServerPrice               *Product_Item_Price
	MaximumPortSpeed                *uint
	MinimumPortSpeed                *uint
	MongoDbEngineeredFlag           *bool
	Name                            *string
	OrderPremiumCount               *uint
	OrderPremiums                   []Product_Item_Price_Premium
	PreconfiguredFlag               *bool
	PresetConfigurationRequiredFlag *bool
	PreventVlanSelectionFlag        *bool
	PrivateHostedCloudPackageFlag   *bool
	PrivateHostedCloudPackageType   *string
	PrivateNetworkOnlyFlag          *bool
	QuantaStorPackageFlag           *bool
	RaidDiskRestrictionFlag         *bool
	RedundantPowerFlag              *bool
	RegionCount                     *uint
	Regions                         []Location_Region
	ResourceGroupTemplate           *Resource_Group_Template
	SubDescription                  *string
	TopLevelItemCategoryCode        *string
	Type                            *Product_Package_Type
	UnitSize                        *int
}

type Product_Package_Attribute struct {
	Entity

	AttributeType *Product_Package_Attribute_Type
	Package       *Product_Package
	Value         *string
}

type Product_Package_Attribute_Type struct {
	Entity

	Description *string
	KeyName     *string
	Name        *string
}

type Product_Package_Inventory struct {
	Entity

	AvailableInventoryCount *int
	Item                    *Product_Item
	ItemId                  *int
	Location                *Location
	LocationId              *int
	ModifyDate              *time.Time
	OverstockFlag           *int
	Package                 *Product_Package
	PackageId               *int
}

type Product_Package_Item_Category_Group struct {
	Entity

	Category       *Product_Item_Category
	ItemCategoryId *int
	Package        *Product_Package
	PackageId      *int
	PriceCount     *uint
	Prices         []Product_Item_Price
	Sort           *int
	Title          *string
}

type Product_Package_Item_Prices struct {
	Entity

	Id          *int
	ItemPrice   *Product_Item_Price
	ItemPriceId *int
	Package     *Product_Package
	PackageId   *int
}

type Product_Package_Items struct {
	Entity

	Id        *string
	Item      *Product_Item
	ItemId    *int
	Package   *Product_Package
	PackageId *int
}

type Product_Package_Locations struct {
	Entity

	DeliveryTimeInformation *string
	IsAvailable             *int
	Location                *Location
	LocationId              *int
	Package                 *Product_Package
	PackageId               *int
}

type Product_Package_Order_Configuration struct {
	Entity

	ErrorMessage   *string
	Id             *int
	IsRequired     *int
	ItemCategory   *Product_Item_Category
	ItemCategoryId *int
	OrderStepId    *int
	Package        *Product_Package
	PackageId      *int
	Sort           *int
	Step           *Product_Package_Order_Step
}

type Product_Package_Order_Step struct {
	Entity

	Id                         *int
	InclusivePreviousStepCount *uint
	InclusivePreviousSteps     []Product_Package_Order_Step_Next
	NextStepCount              *uint
	NextSteps                  []Product_Package_Order_Step_Next
	PreviousStepCount          *uint
	PreviousSteps              []Product_Package_Order_Step_Next
	Step                       *string
}

type Product_Package_Order_Step_Next struct {
	Entity

	Id              *int
	NextOrderStepId *int
	OrderStepId     *int
	Step            *Product_Package_Order_Step
}

type Product_Package_Preset struct {
	Entity

	AvailableStorageUnits          *uint
	Categories                     []Product_Item_Category
	CategoryCount                  *uint
	Configuration                  []Product_Package_Preset_Configuration
	ConfigurationCount             *uint
	Description                    *string
	FixedConfigurationFlag         *bool
	Id                             *int
	IsActive                       *string
	KeyName                        *string
	LowestPresetServerPrice        *Product_Item_Price
	Name                           *string
	Package                        *Product_Package
	PackageConfiguration           []Product_Package_Order_Configuration
	PackageConfigurationCount      *uint
	PackageId                      *int
	PriceCount                     *uint
	Prices                         []Product_Item_Price
	StorageGroupTemplateArrayCount *uint
	StorageGroupTemplateArrays     []Configuration_Storage_Group_Template_Group
	TotalMinimumHourlyFee          *float64
	TotalMinimumRecurringFee       *float64
}

type Product_Package_Preset_Attribute struct {
	Entity

	AttributeType   *Product_Package_Preset_Attribute_Type
	AttributeTypeId *int
	Id              *int
	Preset          *Product_Package_Preset
	PresetId        *int
	Value           *string
}

type Product_Package_Preset_Attribute_Type struct {
	Entity

	Description *string
	Id          *int
	KeyName     *string
	Name        *string
}

type Product_Package_Preset_Configuration struct {
	Entity

	Category      *Product_Item_Category
	PackagePreset *Product_Package_Preset
	Price         *Product_Item_Price
}

type Product_Package_Server struct {
	Entity

	Catalog                *Product_Catalog
	CatalogId              *int
	Datacenters            *string
	DefaultRamCapacity     *float64
	DualPathNetworkFlag    *bool
	GpuFlag                *bool
	HourlyBillingFlag      *bool
	Id                     *int
	Item                   *Product_Item
	ItemId                 *int
	ItemPrice              *Product_Item_Price
	ItemPriceId            *int
	MaximumDriveCount      *int
	MaximumPortSpeed       *float64
	MaximumRamCapacity     *float64
	MinimumPortSpeed       *float64
	OutletFlag             *bool
	Package                *Product_Package
	PackageId              *int
	PackageType            *string
	PowerServerFlag        *bool
	Preset                 *Product_Package_Preset
	PresetId               *int
	PrivateNetworkOnlyFlag *bool
	ProcessorBusSpeed      *string
	ProcessorCache         *string
	ProcessorCores         *int
	ProcessorCount         *int
	ProcessorManufacturer  *string
	ProcessorModel         *string
	ProcessorName          *string
	ProcessorSpeed         *string
	ProductName            *string
	RedundantPowerFlag     *bool
	SapCertifiedServerFlag *bool
	StartingHourlyPrice    *float64
	StartingMonthlyPrice   *float64
	TotalCoreCount         *int
	TxtTpmFlag             *bool
	UnitSize               *int
}

type Product_Package_Server_Option struct {
	Entity

	CatalogId   *int
	Description *string
	Id          *int
	Type        *string
	Value       *string
}

type Product_Package_Type struct {
	Entity

	Id           *int
	KeyName      *string
	Name         *string
	PackageCount *uint
	Packages     []Product_Package
}

type Product_Upgrade_Request struct {
	Entity

	Account                 *Account
	AccountId               *int
	CompletedFlag           *bool
	CreateDate              *time.Time
	EmployeeId              *int
	GuestId                 *int
	HardwareId              *int
	Id                      *int
	Invoice                 *Billing_Invoice
	MaintenanceStartTimeUtc *time.Time
	ModifyDate              *time.Time
	Order                   *Billing_Order
	OrderId                 *int
	OrderTotal              *float64
	ProratedTotal           *float64
	Server                  *Hardware
	Status                  *Product_Upgrade_Request_Status
	StatusId                *int
	Ticket                  *Ticket
	TicketId                *int
	User                    *User_Customer
	UserId                  *int
	VirtualGuest            *Virtual_Guest
}

type Product_Upgrade_Request_Status struct {
	Entity

	Description *string
	Id          *int
	Name        *string
	StatusCode  *string
}

type Provisioning_Hook struct {
	Entity

	Account    *Account
	AccountId  *int
	CreateDate *time.Time
	HookType   *Provisioning_Hook_Type
	Id         *int
	ModifyDate *time.Time
	Name       *string
	TypeId     *int
	Uri        *string
}

type Provisioning_Hook_Type struct {
	Entity

	Description *string
	Id          *int
	KeyName     *string
	Name        *string
}

type Provisioning_Maintenance_Classification struct {
	Entity

	Id                *int
	ItemCategories    []Provisioning_Maintenance_Classification_Item_Category
	ItemCategoryCount *uint
	Slots             *int
	Type              *string
}

type Provisioning_Maintenance_Classification_Item_Category struct {
	Entity

	ItemCategoryId              *int
	MaintenanceClassification   *Provisioning_Maintenance_Classification
	MaintenanceClassificationId *int
}

type Provisioning_Maintenance_Slots struct {
	Entity

	AvailableSlots *int
}

type Provisioning_Maintenance_Ticket struct {
	Entity

	AvailableSlots   *Provisioning_Maintenance_Slots
	MaintClassId     *int
	MaintWindowId    *int
	MaintenanceClass *Provisioning_Maintenance_Classification
	MaintenanceDate  *time.Time
	Ticket           *Ticket
	TicketId         *int
}

type Provisioning_Maintenance_Window struct {
	Entity

	BeginDate  *time.Time
	DayOfWeek  *int
	EndDate    *time.Time
	Id         *int
	LocationId *int
	PortalTzId *int
}

type Provisioning_Version1_Transaction struct {
	Entity

	Account                             *Account
	CreateDate                          *time.Time
	ElapsedSeconds                      *int
	Guest                               *Virtual_Guest
	GuestId                             *int
	Hardware                            *Hardware
	HardwareId                          *int
	Id                                  *int
	Loopback                            []Provisioning_Version1_Transaction
	LoopbackCount                       *uint
	ModifyDate                          *time.Time
	PendingTransactionCount             *uint
	PendingTransactions                 []Provisioning_Version1_Transaction
	StatusChangeDate                    *time.Time
	TicketScheduledActionReference      []Ticket_Attachment
	TicketScheduledActionReferenceCount *uint
	TransactionGroup                    *Provisioning_Version1_Transaction_Group
	TransactionStatus                   *Provisioning_Version1_Transaction_Status
}

type Provisioning_Version1_Transaction_Group struct {
	Entity

	AverageTimeToComplete *float64
	Name                  *string
}

type Provisioning_Version1_Transaction_History struct {
	Entity

	FinishDate          *time.Time
	Guest               *Virtual_Guest
	GuestId             *int
	Hardware            *Hardware
	HardwareId          *int
	HostId              *int
	Id                  *int
	StartDate           *time.Time
	Transaction         *Provisioning_Version1_Transaction
	TransactionId       *int
	TransactionStatus   *Provisioning_Version1_Transaction_Status
	TransactionStatusId *int
}

type Provisioning_Version1_Transaction_Status struct {
	Entity

	AverageDuration              *float64
	FriendlyName                 *string
	Name                         *string
	NonCompletedTransactionCount *uint
	NonCompletedTransactions     []Provisioning_Version1_Transaction
}

type Provisioning_Version1_Transaction_SubnetMigration struct {
	Provisioning_Version1_Transaction
}

type Resource_Group struct {
	Entity

	AncestorGroupCount  *uint
	AncestorGroups      []Resource_Group
	AttributeCount      *uint
	Attributes          []Resource_Group_Attribute
	CreateDate          *time.Time
	Description         *string
	HardwareMemberCount *uint
	HardwareMembers     []Resource_Group_Member
	Id                  *int
	KeyName             *string
	MemberCount         *uint
	Members             []Resource_Group_Member
	Name                *string
	RootResourceGroup   *Resource_Group
	RootResourceGroupId *int
	SubnetMemberCount   *uint
	SubnetMembers       []Resource_Group_Member
	Template            *Resource_Group_Template
	TemplateId          *int
	VlanMemberCount     *uint
	VlanMembers         []Resource_Group_Member
}

type Resource_Group_Attribute struct {
	Entity

	CreateDate *time.Time
	Group      *Resource_Group
	Id         *int
	Type       *Resource_Group_Attribute_Type
	Value      *string
}

type Resource_Group_Attribute_Type struct {
	Entity

	Description *string
	Id          *int
	KeyName     *string
	Name        *string
}

type Resource_Group_Descendant_Reference struct {
	Entity

	Group       *Resource_Group
	GroupMember *Resource_Group_Member
}

type Resource_Group_Member struct {
	Entity

	AttributeCount        *uint
	Attributes            []Resource_Group_Member_Attribute
	CreateDate            *time.Time
	DescendantMemberCount *uint
	DescendantMembers     []Resource_Group_Member
	Group                 *Resource_Group
	Id                    *int
	RoleCount             *uint
	Roles                 []Resource_Group_Role
	Status                *string
	Type                  *Resource_Group_Member_Type
}

type Resource_Group_Member_Attribute struct {
	Entity

	CreateDate *time.Time
	Id         *int
	Member     *Resource_Group_Member
	Type       *Resource_Group_Member_Attribute_Type
	Value      *string
}

type Resource_Group_Member_Attribute_Type struct {
	Entity

	Description *string
	Id          *int
	KeyName     *string
	Name        *string
}

type Resource_Group_Member_CloudStack_Version3_Cluster struct {
	Resource_Group_Member

	Resource *Resource_Group
}

type Resource_Group_Member_CloudStack_Version3_Pod struct {
	Resource_Group_Member

	Resource *Resource_Group
}

type Resource_Group_Member_CloudStack_Version3_Zone struct {
	Resource_Group_Member

	Resource *Resource_Group
}

type Resource_Group_Member_Hardware struct {
	Resource_Group_Member

	Resource          *Hardware
	ServerArbiterOnly *Resource_Group_Member_Attribute
	ServerHidden      *Resource_Group_Member_Attribute
	ServerPriority    *Resource_Group_Member_Attribute
	ServerSlaveDelay  *Resource_Group_Member_Attribute
	ServerTags        *Resource_Group_Member_Attribute
	ServerVotes       *Resource_Group_Member_Attribute
}

type Resource_Group_Member_Network_Storage struct {
	Resource_Group_Member

	Resource *Network_Storage
}

type Resource_Group_Member_Network_Subnet struct {
	Resource_Group_Member

	Resource *Network_Subnet
}

type Resource_Group_Member_Network_Vlan struct {
	Resource_Group_Member

	Resource *Network_Vlan
}

type Resource_Group_Member_Resource_Group struct {
	Resource_Group_Member

	Resource *Resource_Group
}

type Resource_Group_Member_Role_Link struct {
	Entity

	GroupMemberId       *int
	GroupTemplateRoleId *int
}

type Resource_Group_Member_Software_Component_Password struct {
	Resource_Group_Member

	Resource *Software_Component_Password
}

type Resource_Group_Member_Type struct {
	Entity

	Description *string
	KeyName     *string
}

type Resource_Group_Member_Virtual_Host_Pool struct {
	Resource_Group_Member
}

type Resource_Group_Role struct {
	Entity

	Description     *string
	Id              *int
	KeyName         *string
	MemberLinkCount *uint
	MemberLinks     []Resource_Group_Member_Role_Link
}

type Resource_Group_Template struct {
	Entity

	Children      []Resource_Group_Template
	ChildrenCount *uint
	Description   *string
	Id            *int
	KeyName       *string
	MemberCount   *uint
	Members       []Resource_Group_Template_Member
	Package       *Product_Package
}

type Resource_Group_Template_Member struct {
	Entity

	MaxQuantity *int
	MinQuantity *int
	Role        *Resource_Group_Role
	RoleId      *int
	Template    *Resource_Group_Template
	TemplateId  *int
}

type Resource_Metadata struct {
	Entity
}

type Sales_Presale_Event struct {
	Entity

	ActiveFlag  *bool
	Description *string
	EndDate     *time.Time
	ExpiredFlag *bool
	Id          *int
	Item        *Product_Item
	ItemId      *int
	Location    *Location
	LocationId  *int
	OrderCount  *uint
	Orders      []Billing_Order
	StartDate   *time.Time
}

type Scale_Asset struct {
	Entity

	CreateDate   *time.Time
	DeleteFlag   *bool
	Id           *int
	ScaleGroup   *Scale_Group
	ScaleGroupId *int
}

type Scale_Asset_Hardware struct {
	Scale_Asset

	Hardware   *Hardware
	HardwareId *int
}

type Scale_Asset_Virtual_Guest struct {
	Scale_Asset

	VirtualGuest   *Virtual_Guest
	VirtualGuestId *int
}

type Scale_Group struct {
	Entity

	Account                    *Account
	AccountId                  *int
	BalancedTerminationFlag    *bool
	Cooldown                   *int
	CreateDate                 *time.Time
	DesiredMemberCount         *int
	Id                         *int
	LastActionDate             *time.Time
	LoadBalancerCount          *uint
	LoadBalancers              []Scale_LoadBalancer
	LogCount                   *uint
	Logs                       []Scale_Group_Log
	MaximumMemberCount         *int
	MinimumMemberCount         *int
	ModifyDate                 *time.Time
	Name                       *string
	NetworkVlanCount           *uint
	NetworkVlans               []Scale_Network_Vlan
	Policies                   []Scale_Policy
	PolicyCount                *uint
	RegionalGroup              *Location_Group_Regional
	RegionalGroupId            *int
	Status                     *Scale_Group_Status
	SuspendedFlag              *bool
	TerminationPolicy          *Scale_Termination_Policy
	TerminationPolicyId        *int
	VirtualGuestAssetCount     *uint
	VirtualGuestAssets         []Scale_Asset_Virtual_Guest
	VirtualGuestMemberCount    *uint
	VirtualGuestMemberTemplate *Virtual_Guest
	VirtualGuestMembers        []Scale_Member_Virtual_Guest
}

type Scale_Group_Log struct {
	Entity

	CreateDate   *time.Time
	Description  *string
	Id           *int
	ScaleGroup   *Scale_Group
	ScaleGroupId *int
}

type Scale_Group_Status struct {
	Entity

	Id      *int
	KeyName *string
	Name    *string
}

type Scale_LoadBalancer struct {
	Entity

	AllocationPercent  *int
	CreateDate         *time.Time
	DeleteFlag         *bool
	HealthCheck        *Network_Application_Delivery_Controller_LoadBalancer_Health_Check
	HealthCheckId      *int
	Id                 *int
	ModifyDate         *time.Time
	Port               *int
	RoutingMethod      *Network_Application_Delivery_Controller_LoadBalancer_Routing_Method
	RoutingType        *Network_Application_Delivery_Controller_LoadBalancer_Routing_Type
	ScaleGroup         *Scale_Group
	ScaleGroupId       *int
	VirtualIpAddressId *int
	VirtualServer      *Network_Application_Delivery_Controller_LoadBalancer_VirtualServer
	VirtualServerId    *int
	VirtualServerPort  *int
}

type Scale_Member struct {
	Entity

	CreateDate   *time.Time
	Id           *int
	ScaleGroup   *Scale_Group
	ScaleGroupId *int
}

type Scale_Member_Virtual_Guest struct {
	Scale_Member

	VirtualGuest   *Virtual_Guest
	VirtualGuestId *int
}

type Scale_Network_Vlan struct {
	Entity

	CreateDate    *time.Time
	DeleteFlag    *bool
	Id            *int
	NetworkVlan   *Network_Vlan
	NetworkVlanId *int
	ScaleGroup    *Scale_Group
	ScaleGroupId  *int
}

type Scale_Policy struct {
	Entity

	ActionCount             *uint
	Actions                 []Scale_Policy_Action
	Cooldown                *int
	CreateDate              *time.Time
	DeleteFlag              *bool
	Id                      *int
	ModifyDate              *time.Time
	Name                    *string
	OneTimeTriggerCount     *uint
	OneTimeTriggers         []Scale_Policy_Trigger_OneTime
	RepeatingTriggerCount   *uint
	RepeatingTriggers       []Scale_Policy_Trigger_Repeating
	ResourceUseTriggerCount *uint
	ResourceUseTriggers     []Scale_Policy_Trigger_ResourceUse
	ScaleActionCount        *uint
	ScaleActions            []Scale_Policy_Action_Scale
	ScaleGroup              *Scale_Group
	ScaleGroupId            *int
	TriggerCount            *uint
	Triggers                []Scale_Policy_Trigger
}

type Scale_Policy_Action struct {
	Entity

	CreateDate    *time.Time
	DeleteFlag    *bool
	Id            *int
	ModifyDate    *time.Time
	ScalePolicy   *Scale_Policy
	ScalePolicyId *int
	Type          *Scale_Policy_Action_Type
	TypeId        *int
}

type Scale_Policy_Action_Scale struct {
	Scale_Policy_Action

	Amount    *int
	ScaleType *string
}

type Scale_Policy_Action_Type struct {
	Entity

	Id      *int
	KeyName *string
	Name    *string
}

type Scale_Policy_Trigger struct {
	Entity

	CreateDate    *time.Time
	DeleteFlag    *bool
	Id            *int
	ModifyDate    *time.Time
	ScalePolicy   *Scale_Policy
	ScalePolicyId *int
	Type          *Scale_Policy_Trigger_Type
	TypeId        *int
}

type Scale_Policy_Trigger_OneTime struct {
	Scale_Policy_Trigger

	Date *time.Time
}

type Scale_Policy_Trigger_Repeating struct {
	Scale_Policy_Trigger

	Schedule *string
}

type Scale_Policy_Trigger_ResourceUse struct {
	Scale_Policy_Trigger

	WatchCount *uint
	Watches    []Scale_Policy_Trigger_ResourceUse_Watch
}

type Scale_Policy_Trigger_ResourceUse_Watch struct {
	Entity

	Algorithm            *string
	CreateDate           *time.Time
	DeleteFlag           *bool
	Id                   *int
	Metric               *string
	ModifyDate           *time.Time
	Operator             *string
	Period               *int
	ScalePolicyTrigger   *Scale_Policy_Trigger_ResourceUse
	ScalePolicyTriggerId *int
	Value                *string
}

type Scale_Policy_Trigger_Type struct {
	Entity

	Id      *int
	KeyName *string
	Name    *string
}

type Scale_Termination_Policy struct {
	Entity

	Id      *int
	KeyName *string
	Name    *string
}

type Search struct {
	Entity
}

type Security_Certificate struct {
	Entity

	AssociatedServiceCount            *int
	Certificate                       *string
	CertificateSigningRequest         *string
	CommonName                        *string
	CreateDate                        *time.Time
	Id                                *int
	IntermediateCertificate           *string
	KeySize                           *int
	LoadBalancerVirtualIpAddressCount *uint
	LoadBalancerVirtualIpAddresses    []Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress
	ModifyDate                        *time.Time
	Notes                             *string
	OrganizationName                  *string
	PrivateKey                        *string
	ValidityBegin                     *time.Time
	ValidityDays                      *int
	ValidityEnd                       *time.Time
}

type Security_Certificate_Entry struct {
	Entity

	CertificateId    *int
	CommonName       *string
	KeySize          *int
	OrganizationName *string
	ValidityBegin    *time.Time
	ValidityDays     *int
	ValidityEnd      *time.Time
}

type Security_Certificate_Request struct {
	Entity

	Account                      *Account
	AccountId                    *int
	ApproverEmailAddress         *string
	CertificateAuthorityName     *string
	CertificateSigningRequest    *string
	CommonName                   *string
	CreateDate                   *time.Time
	EffectiveDate                *time.Time
	ExpirationDate               *time.Time
	Id                           *int
	ModifyDate                   *time.Time
	Order                        *Billing_Order
	OrderItem                    *Billing_Order_Item
	Status                       *Security_Certificate_Request_Status
	StatusId                     *int
	TechnicalContactEmailAddress *string
}

type Security_Certificate_Request_ServerType struct {
	Entity

	Description *string
	Id          *int
	Name        *string
	Value       *int
}

type Security_Certificate_Request_Status struct {
	Entity

	Description *string
	Id          *int
	Name        *string
}

type Security_Directory_Service_Host_Xref_Hardware struct {
	Entity

	Host *Hardware
}

type Security_SecureTransportCipher struct {
	Entity

	KeyName *string
}

type Security_SecureTransportProtocol struct {
	Entity

	KeyName                         *string
	SupportedSecureTransportCiphers []Security_SecureTransportCipher
}

type Security_Ssh_Key struct {
	Entity

	Account                       *Account
	BlockDeviceTemplateGroupCount *uint
	BlockDeviceTemplateGroups     []Virtual_Guest_Block_Device_Template_Group
	CreateDate                    *time.Time
	Fingerprint                   *string
	Id                            *int
	Key                           *string
	Label                         *string
	ModifyDate                    *time.Time
	Notes                         *string
	SoftwarePasswordCount         *uint
	SoftwarePasswords             []Software_Component_Password
}

type Service_Provider struct {
	Entity

	Description *string
	Id          *int
	KeyName     *string
	Name        *string
}

type Software_AccountLicense struct {
	Entity

	Account             *Account
	AccountId           *int
	BillingItem         *Billing_Item
	Capacity            *string
	Key                 *string
	SoftwareDescription *Software_Description
	Units               *string
}

type Software_Component struct {
	Entity

	AverageInstallationDuration *uint
	BillingItem                 *Billing_Item
	Hardware                    *Hardware
	HardwareId                  *int
	Id                          *int
	ManufacturerActivationCode  *string
	ManufacturerLicenseInstance *string
	PasswordCount               *uint
	PasswordHistory             []Software_Component_Password_History
	PasswordHistoryCount        *uint
	Passwords                   []Software_Component_Password
	SoftwareDescription         *Software_Description
	SoftwareLicense             *Software_License
	VirtualGuest                *Virtual_Guest
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

	AgentDetails                     *McAfee_Epolicy_Orchestrator_Version36_Agent_Details
	CurrentAntivirusPolicy           *int
	DataFileVersion                  *McAfee_Epolicy_Orchestrator_Version36_Product_Properties
	EpoVersion                       *string
	LatestAccessProtectionEventCount *uint
	LatestAccessProtectionEvents     []McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_AccessProtection
	LatestAntivirusEventCount        *uint
	LatestAntivirusEvents            []McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event
	LatestSpywareEventCount          *uint
	LatestSpywareEvents              []McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event
	TransactionStatus                *string
}

type Software_Component_AntivirusSpyware_Mcafee_Epo_Version45 struct {
	Software_Component_AntivirusSpyware_Mcafee

	AgentDetails                     *McAfee_Epolicy_Orchestrator_Version45_Agent_Details
	CurrentAntivirusPolicy           *int
	DataFileVersion                  *McAfee_Epolicy_Orchestrator_Version45_Product_Properties
	EpoVersion                       *string
	LatestAccessProtectionEventCount *uint
	LatestAccessProtectionEvents     []McAfee_Epolicy_Orchestrator_Version45_Event
	LatestAntivirusEventCount        *uint
	LatestAntivirusEvents            []McAfee_Epolicy_Orchestrator_Version45_Event
	LatestSpywareEventCount          *uint
	LatestSpywareEvents              []McAfee_Epolicy_Orchestrator_Version45_Event
	TransactionStatus                *string
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

	AgentDetails                      *McAfee_Epolicy_Orchestrator_Version36_Agent_Details
	ApplicationModePolicyNameCount    *uint
	ApplicationModePolicyNames        []McAfee_Epolicy_Orchestrator_Version36_Policy_Object
	ApplicationRuleSetPolicyNameCount *uint
	ApplicationRuleSetPolicyNames     []McAfee_Epolicy_Orchestrator_Version36_Policy_Object
	EnforcementPolicyNameCount        *uint
	EnforcementPolicyNames            []McAfee_Epolicy_Orchestrator_Version36_Policy_Object
	EpoVersion                        *string
	FirewallModePolicyNameCount       *uint
	FirewallModePolicyNames           []McAfee_Epolicy_Orchestrator_Version36_Policy_Object
	FirewallRuleSetPolicyNameCount    *uint
	FirewallRuleSetPolicyNames        []McAfee_Epolicy_Orchestrator_Version36_Policy_Object
	IpsModePolicyNameCount            *uint
	IpsModePolicyNames                []McAfee_Epolicy_Orchestrator_Version36_Policy_Object
	IpsProtectionPolicyNameCount      *uint
	IpsProtectionPolicyNames          []McAfee_Epolicy_Orchestrator_Version36_Policy_Object
	TransactionStatus                 *string
}

type Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version6 struct {
	Software_Component_HostIps_Mcafee_Epo_Version36_Hips

	BlockedApplicationEventCount *uint
	BlockedApplicationEvents     []McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_BlockedApplicationEvent
	IpsEventCount                *uint
	IpsEvents                    []McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_IPSEvent
}

type Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version7 struct {
	Software_Component_HostIps_Mcafee_Epo_Version36_Hips

	BlockedApplicationEventCount *uint
	BlockedApplicationEvents     []McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_BlockedApplicationEvent
	IpsEventCount                *uint
	IpsEvents                    []McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent
}

type Software_Component_HostIps_Mcafee_Epo_Version45_Hips struct {
	Software_Component_HostIps_Mcafee

	AgentDetails                      *McAfee_Epolicy_Orchestrator_Version45_Agent_Details
	ApplicationModePolicyNameCount    *uint
	ApplicationModePolicyNames        []McAfee_Epolicy_Orchestrator_Version45_Policy_Object
	ApplicationRuleSetPolicyNameCount *uint
	ApplicationRuleSetPolicyNames     []McAfee_Epolicy_Orchestrator_Version45_Policy_Object
	BlockedApplicationEventCount      *uint
	BlockedApplicationEvents          []McAfee_Epolicy_Orchestrator_Version45_Event
	EnforcementPolicyNameCount        *uint
	EnforcementPolicyNames            []McAfee_Epolicy_Orchestrator_Version45_Policy_Object
	EpoVersion                        *string
	FirewallModePolicyNameCount       *uint
	FirewallModePolicyNames           []McAfee_Epolicy_Orchestrator_Version45_Policy_Object
	FirewallRuleSetPolicyNameCount    *uint
	FirewallRuleSetPolicyNames        []McAfee_Epolicy_Orchestrator_Version45_Policy_Object
	IpsEventCount                     *uint
	IpsEvents                         []McAfee_Epolicy_Orchestrator_Version45_Event
	IpsModePolicyNameCount            *uint
	IpsModePolicyNames                []McAfee_Epolicy_Orchestrator_Version45_Policy_Object
	IpsProtectionPolicyNameCount      *uint
	IpsProtectionPolicyNames          []McAfee_Epolicy_Orchestrator_Version45_Policy_Object
	TransactionStatus                 *string
}

type Software_Component_HostIps_Mcafee_Epo_Version45_Hips_Version7 struct {
	Software_Component_HostIps_Mcafee_Epo_Version45_Hips
}

type Software_Component_HostIps_Mcafee_Epo_Version45_Hips_Version8 struct {
	Software_Component_HostIps_Mcafee_Epo_Version45_Hips
}

type Software_Component_OperatingSystem struct {
	Software_Component

	LicenseExpirationDate  *time.Time
	PartitionTemplateCount *uint
	PartitionTemplates     []Hardware_Component_Partition_Template
	ReloadTransactionGroup *Provisioning_Version1_Transaction_Group
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

	CreateDate  *time.Time
	Id          *int
	ModifyDate  *time.Time
	Notes       *string
	Password    *string
	Port        *int
	Software    *Software_Component
	SoftwareId  *int
	SshKeyCount *uint
	SshKeys     []Security_Ssh_Key
	Username    *string
}

type Software_Component_Password_History struct {
	Entity

	CreateDate          *time.Time
	Notes               *string
	Password            *string
	SoftwareComponent   *Software_Component
	SoftwareComponentId *int
	Username            *string
}

type Software_Component_Security struct {
	Software_Component
}

type Software_Component_Security_SafeNet struct {
	Software_Component_Security
}

type Software_Description struct {
	Entity

	AttributeCount                     *uint
	Attributes                         []Software_Description_Attribute
	AverageInstallationDuration        *int
	CompatibleSoftwareDescriptionCount *uint
	CompatibleSoftwareDescriptions     []Software_Description
	ControlPanel                       *int
	FeatureCount                       *uint
	Features                           []Software_Description_Feature
	Id                                 *int
	LatestVersion                      []Software_Description
	LatestVersionCount                 *uint
	LicenseTermUnit                    *string
	LicenseTermValue                   *int
	LongDescription                    *string
	Manufacturer                       *string
	Name                               *string
	OperatingSystem                    *int
	ProductItemCount                   *uint
	ProductItems                       []Product_Item
	ProvisionTransactionGroup          *Provisioning_Version1_Transaction_Group
	ReferenceCode                      *string
	ReloadTransactionGroup             *Provisioning_Version1_Transaction_Group
	RequiredUser                       *string
	SoftwareLicenseCount               *uint
	SoftwareLicenses                   []Software_License
	UpgradeSoftwareDescription         *Software_Description
	UpgradeSoftwareDescriptionId       *int
	UpgradeSwDesc                      *Software_Description
	UpgradeSwDescId                    *int
	ValidFilesystemTypeCount           *uint
	ValidFilesystemTypes               []Configuration_Storage_Filesystem_Type
	Version                            *string
	VirtualLicense                     *int
	VirtualizationPlatform             *int
}

type Software_Description_Attribute struct {
	Entity

	SoftwareDescription *Software_Description
	Type                *Software_Description_Attribute_Type
	Value               *string
}

type Software_Description_Attribute_Type struct {
	Entity

	Keyname *string
}

type Software_Description_Feature struct {
	Entity

	Id      *int
	KeyName *string
	Name    *string
	Vendor  *string
}

type Software_Description_RequiredUser struct {
	Entity

	DefaultPassword *string
	Username        *string
}

type Software_License struct {
	Entity

	Account               *Account
	Id                    *int
	Owner                 *Account
	SoftwareDescription   *Software_Description
	SoftwareDescriptionId *int
}

type Software_VirtualLicense struct {
	Entity

	Account               *Account
	AccountId             *int
	BillingItem           *Billing_Item
	HostHardware          *Hardware_Server
	HostHardwareId        *int
	Id                    *int
	IpAddress             *string
	IpAddressRecord       *Network_Subnet_IpAddress
	Key                   *string
	Notes                 *string
	SoftwareDescription   *Software_Description
	SoftwareDescriptionId *int
	Subnet                *Network_Subnet
	SubnetId              *int
}

type Survey struct {
	Entity

	Active        *int
	CreateDate    *time.Time
	Id            *int
	Name          *string
	QuestionCount *uint
	Questions     []Survey_Question
	Status        *Survey_Status
	StatusId      *int
	Type          *Survey_Type
	TypeId        *int
}

type Survey_Answer struct {
	Entity

	Answer           *string
	AnswerOrder      *int
	Id               *int
	SurveyQuestion   *Survey_Question
	SurveyQuestionId *int
}

type Survey_Question struct {
	Entity

	AnswerCount   *uint
	Answers       []Survey_Answer
	Id            *int
	IsRequired    *int
	MultiAnswer   *int
	Question      *string
	QuestionOrder *int
	Survey        *Survey
	SurveyId      *int
}

type Survey_Response struct {
	Entity

	OtherAnswer    *string
	SurveyAnswer   *Survey_Answer
	SurveyAnswerId *int
}

type Survey_Status struct {
	Entity

	Description *string
	Id          *int
	Name        *string
}

type Survey_Type struct {
	Entity

	Description *string
	Id          *int
	Name        *string
}

type Tag struct {
	Entity

	Account        *Account
	AccountId      *int
	Id             *int
	Internal       *int
	Name           *string
	ReferenceCount *uint
	References     []Tag_Reference
}

type Tag_Reference struct {
	Entity

	Customer        *User_Customer
	EmpRecordId     *int
	Employee        *User_Employee
	Id              *int
	ResourceTableId *int
	Tag             *Tag
	TagId           *int
	TagType         *Tag_Type
	TagTypeId       *int
	UsrRecordId     *int
}

type Tag_Reference_Hardware struct {
	Tag_Reference

	Resource *Hardware
}

type Tag_Reference_Network_Application_Delivery_Controller struct {
	Tag_Reference

	Resource *Network_Application_Delivery_Controller
}

type Tag_Reference_Network_Vlan struct {
	Tag_Reference

	Resource *Network_Vlan
}

type Tag_Reference_Network_Vlan_Firewall struct {
	Tag_Reference

	Resource *Network_Vlan_Firewall
}

type Tag_Reference_Resource_Group struct {
	Tag_Reference

	Resource *Resource_Group
}

type Tag_Reference_Virtual_Guest struct {
	Tag_Reference

	Resource *Virtual_Guest
}

type Tag_Reference_Virtual_Guest_Block_Device_Template_Group struct {
	Tag_Reference

	Resource *Virtual_Guest_Block_Device_Template_Group
}

type Tag_Type struct {
	Entity

	Description *string
	KeyName     *string
}

type Ticket struct {
	Entity

	Account                              *Account
	AccountId                            *int
	AssignedAgentCount                   *uint
	AssignedAgents                       []User_Customer
	AssignedUser                         *User_Customer
	AssignedUserId                       *int
	AttachedAdditionalEmailCount         *uint
	AttachedAdditionalEmails             []User_Customer_AdditionalEmail
	AttachedFileCount                    *uint
	AttachedFiles                        []Ticket_Attachment_File
	AttachedHardware                     []Hardware
	AttachedHardwareCount                *uint
	AttachedResourceCount                *uint
	AttachedResources                    []Ticket_Attachment
	AttachedVirtualGuestCount            *uint
	AttachedVirtualGuests                []Virtual_Guest
	AwaitingUserResponseFlag             *bool
	BillableFlag                         *bool
	CancellationRequest                  *Billing_Item_Cancellation_Request
	ChangeOwnerFlag                      *bool
	CreateDate                           *time.Time
	EmployeeAttachmentCount              *uint
	EmployeeAttachments                  []User_Employee
	FinalComments                        *string
	FirstAttachedResource                *Ticket_Attachment
	FirstUpdate                          *Ticket_Update
	Group                                *Ticket_Group
	GroupId                              *int
	Id                                   *int
	InvoiceItemCount                     *uint
	InvoiceItems                         []Billing_Invoice_Item
	LastActivity                         *Ticket_Activity
	LastEditDate                         *time.Time
	LastEditType                         *string
	LastEditor                           *User_Interface
	LastResponseDate                     *time.Time
	LastUpdate                           *Ticket_Update
	LastViewedDate                       *time.Time
	Location                             *Location
	LocationId                           *int
	ModifyDate                           *time.Time
	NewUpdatesFlag                       *bool
	NotifyUserOnUpdateFlag               *bool
	OriginatingIpAddress                 *string
	Priority                             *int
	ResponsibleBrandId                   *int
	ScheduledActionCount                 *uint
	ScheduledActions                     []Provisioning_Version1_Transaction
	ServerAdministrationBillingAmount    *int
	ServerAdministrationBillingInvoice   *Billing_Invoice
	ServerAdministrationBillingInvoiceId *int
	ServerAdministrationFlag             *int
	ServerAdministrationRefundInvoice    *Billing_Invoice
	ServerAdministrationRefundInvoiceId  *int
	ServiceProvider                      *Service_Provider
	ServiceProviderId                    *int
	ServiceProviderResourceId            *int
	State                                []Ticket_State
	StateCount                           *uint
	Status                               *Ticket_Status
	StatusId                             *int
	Subject                              *Ticket_Subject
	SubjectId                            *int
	TagReferenceCount                    *uint
	TagReferences                        []Tag_Reference
	Title                                *string
	TotalUpdateCount                     *int
	UpdateCount                          *uint
	Updates                              []Ticket_Update
	UserEditableFlag                     *bool
}

type Ticket_Activity struct {
	Entity

	CreateDate      *time.Time
	CreateTimestamp *time.Time
	Editor          *User_Interface
	Id              *int
	Ticket          *Ticket
	TicketUpdate    *Ticket_Update
	Value           *string
}

type Ticket_Attachment struct {
	Entity

	AssignedAgent   *User_Customer
	AttachmentId    *int
	CreateDate      *time.Time
	Id              *int
	ScheduledAction *Provisioning_Version1_Transaction
	Ticket          *Ticket
	TicketId        *int
}

type Ticket_Attachment_Assigned_Agent struct {
	Ticket_Attachment

	AssignedAgentId *int
	Resource        *User_Customer
}

type Ticket_Attachment_CardChangeRequest struct {
	Ticket_Attachment

	Resource *Billing_Payment_Card_ChangeRequest
}

type Ticket_Attachment_File struct {
	Entity

	CreateDate   *time.Time
	FileName     *string
	FileSize     *string
	Id           *int
	ModifyDate   *time.Time
	Ticket       *Ticket
	TicketId     *int
	Update       *Ticket_Update
	UpdateId     *int
	UploaderId   *string
	UploaderType *string
}

type Ticket_Attachment_Hardware struct {
	Ticket_Attachment

	Hardware   *Hardware
	HardwareId *int
	Resource   *Hardware
}

type Ticket_Attachment_ManualPayment struct {
	Ticket_Attachment

	Resource *Billing_Payment_Card_ManualPayment
}

type Ticket_Attachment_Scheduled_Action struct {
	Ticket_Attachment

	Resource      *Provisioning_Version1_Transaction
	RunDate       *time.Time
	Transaction   *Provisioning_Version1_Transaction
	TransactionId *int
}

type Ticket_Attachment_Virtual_Guest struct {
	Ticket_Attachment

	Resource       *Virtual_Guest
	VirtualGuest   *Virtual_Guest
	VirtualGuestId *int
}

type Ticket_Chat struct {
	Entity

	Agent        *User_Employee
	Customer     *User_Customer
	CustomerId   *int
	EndDate      *time.Time
	StartDate    *time.Time
	TicketUpdate *Ticket_Update_Chat
	Transcript   *string
}

type Ticket_Chat_Liveperson struct {
	Ticket_Chat
}

type Ticket_Chat_TranscriptLine struct {
	Entity

	Speaker *User_Interface
}

type Ticket_Chat_TranscriptLine_Customer struct {
	Ticket_Chat_TranscriptLine
}

type Ticket_Chat_TranscriptLine_Employee struct {
	Ticket_Chat_TranscriptLine
}

type Ticket_Group struct {
	Entity

	AssignedBrandCount    *uint
	AssignedBrands        []Brand
	Category              *Ticket_Group_Category
	Id                    *int
	Name                  *string
	TicketGroupCategoryId *int
}

type Ticket_Group_Category struct {
	Entity

	Id   *int
	Name *string
}

type Ticket_Priority struct {
	Entity
}

type Ticket_State struct {
	Entity

	Id          *int
	StateType   *Ticket_State_Type
	StateTypeId *int
	Ticket      *Ticket
	TicketId    *int
}

type Ticket_State_Type struct {
	Entity

	Description *string
	Id          *int
	KeyName     *string
	Name        *string
}

type Ticket_Status struct {
	Entity

	Id   *int
	Name *string
}

type Ticket_Subject struct {
	Entity

	Group *Ticket_Group
	Id    *int
	Name  *string
}

type Ticket_Survey struct {
	Entity
}

type Ticket_Type struct {
	Entity

	Id      *int
	KeyName *string
}

type Ticket_Update struct {
	Entity

	ChangeOwnerActivity *string
	CreateDate          *time.Time
	Editor              *User_Interface
	EditorId            *int
	EditorType          *string
	Entry               *string
	FileAttachment      []Ticket_Attachment_File
	FileAttachmentCount *uint
	Id                  *int
	Ticket              *Ticket
	TicketId            *int
	Type                *Ticket_Update_Type
}

type Ticket_Update_Agent struct {
	Ticket_Update
}

type Ticket_Update_Chat struct {
	Ticket_Update

	Chat *Ticket_Chat_Liveperson
}

type Ticket_Update_Customer struct {
	Ticket_Update
}

type Ticket_Update_Employee struct {
	Ticket_Update

	ResponseRating *int
}

type Ticket_Update_Type struct {
	Entity

	Description *string
	KeyName     *string
	Ticket      *Ticket_Update
}

type User_Access_Facility_Log struct {
	Entity

	Account     *Account
	AccountId   *int
	Datacenter  *Location
	Description *string
	Hardware    *Hardware
	HardwareId  *int
	Id          *int
	LocationId  *int
	LogType     *User_Access_Facility_Log_Type
	TimeIn      *time.Time
	TimeOut     *time.Time
	Visitor     *Entity
}

type User_Access_Facility_Log_Type struct {
	Entity

	Id      *int
	KeyName *string
	Name    *string
}

type User_Access_Facility_Visitor struct {
	Entity

	CompanyName *string
	FirstName   *string
	LastName    *string
	TypeId      *int
	VisitorType *User_Access_Facility_Visitor_Type
}

type User_Access_Facility_Visitor_Type struct {
	Entity

	Id      *int
	KeyName *string
	Name    *string
}

type User_Customer struct {
	User_Interface

	Account                                  *Account
	AccountId                                *int
	ActionCount                              *uint
	Actions                                  []User_Permission_Action
	AdditionalEmailCount                     *uint
	AdditionalEmails                         []User_Customer_AdditionalEmail
	Address1                                 *string
	Address2                                 *string
	Aim                                      *string
	AlternatePhone                           *string
	ApiAuthenticationKeyCount                *uint
	ApiAuthenticationKeys                    []User_Customer_ApiAuthentication
	AuthenticationToken                      *Container_User_Authentication_Token
	CdnAccountCount                          *uint
	CdnAccounts                              []Network_ContentDelivery_Account
	ChildUserCount                           *uint
	ChildUsers                               []User_Customer
	City                                     *string
	ClosedTicketCount                        *uint
	ClosedTickets                            []Ticket
	CompanyName                              *string
	Country                                  *string
	CreateDate                               *time.Time
	DaylightSavingsTimeFlag                  *bool
	DenyAllResourceAccessOnCreateFlag        *bool
	DisplayName                              *string
	Email                                    *string
	ExternalBindingCount                     *uint
	ExternalBindings                         []User_External_Binding
	FirstName                                *string
	ForumPasswordHash                        *string
	Hardware                                 []Hardware
	HardwareCount                            *uint
	HardwareNotificationCount                *uint
	HardwareNotifications                    []User_Customer_Notification_Hardware
	HasAcknowledgedSupportPolicyFlag         *bool
	HasFullHardwareAccessFlag                *bool
	HasFullVirtualGuestAccessFlag            *bool
	Icq                                      *string
	Id                                       *int
	IpAddressRestriction                     *string
	LastName                                 *string
	LayoutProfileCount                       *uint
	LayoutProfiles                           []Layout_Profile
	Locale                                   *Locale
	LocaleId                                 *int
	LoginAttemptCount                        *uint
	LoginAttempts                            []User_Customer_Access_Authentication
	ManagedByFederationFlag                  *bool
	ManagedByOpenIdConnectFlag               *bool
	MobileDeviceCount                        *uint
	MobileDevices                            []User_Customer_MobileDevice
	ModifyDate                               *time.Time
	Msn                                      *string
	NameId                                   *string
	NotificationSubscriberCount              *uint
	NotificationSubscribers                  []Notification_Subscriber
	OfficePhone                              *string
	OpenIdConnectUserName                    *string
	OpenTicketCount                          *uint
	OpenTickets                              []Ticket
	OverrideCount                            *uint
	Overrides                                []Network_Service_Vpn_Overrides
	Parent                                   *User_Customer
	ParentId                                 *int
	PasswordExpireDate                       *time.Time
	PermissionCount                          *uint
	PermissionSystemVersion                  *int
	Permissions                              []User_Customer_CustomerPermission_Permission
	PostalCode                               *string
	PptpVpnAllowedFlag                       *bool
	PreferenceCount                          *uint
	Preferences                              []User_Preference
	RoleCount                                *uint
	Roles                                    []User_Permission_Role
	SalesforceUserLink                       *User_Customer_Link
	SavedId                                  *string
	SecondaryLoginManagementFlag             *bool
	SecondaryLoginRequiredFlag               *bool
	SecondaryPasswordModifyDate              *time.Time
	SecondaryPasswordTimeoutDays             *int
	SecurityAnswerCount                      *uint
	SecurityAnswers                          []User_Customer_Security_Answer
	Sms                                      *string
	SslVpnAllowedFlag                        *bool
	State                                    *string
	StatusDate                               *time.Time
	SubscriberCount                          *uint
	Subscribers                              []Notification_User_Subscriber
	SuccessfulLoginCount                     *uint
	SuccessfulLogins                         []User_Customer_Access_Authentication
	SupportPolicyAcknowledgementRequiredFlag *int
	SurveyCount                              *uint
	SurveyRequiredFlag                       *bool
	Surveys                                  []Survey
	TicketCount                              *uint
	Tickets                                  []Ticket
	Timezone                                 *Locale_Timezone
	TimezoneId                               *int
	UnsuccessfulLoginCount                   *uint
	UnsuccessfulLogins                       []User_Customer_Access_Authentication
	UserLinkCount                            *uint
	UserLinks                                []User_Customer_Link
	UserStatus                               *User_Customer_Status
	UserStatusId                             *int
	Username                                 *string
	VirtualGuestCount                        *uint
	VirtualGuests                            []Virtual_Guest
	VpnManualConfig                          *bool
	Yahoo                                    *string
}

type User_Customer_Access_Authentication struct {
	Entity

	CreateDate  *time.Time
	IpAddress   *string
	SuccessFlag *bool
	User        *User_Customer
	UserId      *int
	Username    *string
}

type User_Customer_AdditionalEmail struct {
	Entity

	Email  *string
	User   *User_Customer
	UserId *int
}

type User_Customer_ApiAuthentication struct {
	Entity

	AuthenticationKey    *string
	Id                   *int
	IpAddressRestriction *string
	TimestampKey         *int
	User                 *User_Customer
	UserId               *int
}

type User_Customer_CustomerPermission_Permission struct {
	Entity

	Key     *string
	KeyName *string
	Name    *string
}

type User_Customer_External_Binding struct {
	User_External_Binding

	User *User_Customer
}

type User_Customer_External_Binding_Attribute struct {
	User_External_Binding_Attribute
}

type User_Customer_External_Binding_Phone struct {
	User_Customer_External_Binding

	BindingStatus *string
	PinLength     *string
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

	CredentialExpirationDate *string
	CredentialLastUpdateDate *string
	CredentialState          *string
	CredentialType           *string
}

type User_Customer_Invitation struct {
	Entity

	Code                       *string
	CreateDate                 *time.Time
	CreatorId                  *int
	CreatorType                *string
	Email                      *string
	ExistingBlueIdFlag         *int
	ExpirationDate             *time.Time
	Id                         *int
	IsFederatedEmailDomainFlag *int
	ModifyDate                 *time.Time
	ResponseDate               *time.Time
	StatusId                   *int
	User                       *User_Customer
	UserId                     *int
}

type User_Customer_Link struct {
	Entity

	CreateDate                    *time.Time
	DefaultFlag                   *int
	DestinationUserAlphanumericId *string
	DestinationUserId             *int
	Id                            *int
	ServiceProvider               *Service_Provider
	ServiceProviderId             *int
	User                          *User_Customer
	UserId                        *int
}

type User_Customer_Link_ThePlanet struct {
	User_Customer_Link
}

type User_Customer_MobileDevice struct {
	Entity

	AvailablePushNotificationSubscriptionCount *uint
	AvailablePushNotificationSubscriptions     []Notification
	CreateDate                                 *time.Time
	Customer                                   *User_Customer
	DisplayResolutionXxY                       *string
	Id                                         *int
	MobileDeviceTypeId                         *int
	MobileOperatingSystemId                    *int
	ModelNumber                                *string
	ModifyDate                                 *time.Time
	Name                                       *string
	OperatingSystem                            *User_Customer_MobileDevice_OperatingSystem
	PhoneNumber                                *string
	PushNotificationSubscriptionCount          *uint
	PushNotificationSubscriptions              []Notification_User_Subscriber
	SerialNumber                               *string
	Token                                      *string
	Type                                       *User_Customer_MobileDevice_Type
	UserId                                     *int
}

type User_Customer_MobileDevice_OperatingSystem struct {
	Entity

	BuildVersion *int
	CreateDate   *time.Time
	Description  *string
	Id           *int
	MajorVersion *int
	MinorVersion *int
	ModifyDate   *time.Time
	Name         *string
}

type User_Customer_MobileDevice_Type struct {
	Entity

	CreateDate  *time.Time
	Description *string
	Id          *int
	ModifyDate  *time.Time
	Name        *string
}

type User_Customer_Notification_Hardware struct {
	Entity

	Hardware   *Hardware
	HardwareId *int
	Id         *int
	User       *User_Customer
	UserId     *int
}

type User_Customer_Notification_Virtual_Guest struct {
	Entity

	Guest   *Virtual_Guest
	GuestId *int
	Id      *int
	User    *User_Customer
	UserId  *int
}

type User_Customer_OpenIdConnect struct {
	User_Customer
}

type User_Customer_Prospect struct {
	Entity

	Account               *Account
	AssignedEmployeeCount *uint
	AssignedEmployees     []User_Employee
	QuoteCount            *uint
	Quotes                []Billing_Order_Quote
	Type                  *User_Customer_Prospect_Type
}

type User_Customer_Prospect_ServiceProvider_EnrollRequest struct {
	Entity

	AccountId                   *int
	Address1                    *string
	Address2                    *string
	CardAccountNumber           *string
	CardExpirationMonth         *string
	CardExpirationYear          *string
	CardType                    *string
	CardVerificationNumber      *string
	City                        *string
	CompanyName                 *string
	CompanyType                 *Catalyst_Company_Type
	CompanyTypeId               *int
	CompanyUrl                  *string
	ContactEmail                *string
	ContactFirstName            *string
	ContactLastName             *string
	ContactPhone                *string
	Country                     *string
	CustomerProspectId          *int
	DeviceFingerprintId         *string
	Email                       *string
	ExistingCustomerFlag        *bool
	FirstName                   *string
	IbmPartnerWorldId           *string
	IbmPartnerWorldMemberFlag   *bool
	LastName                    *string
	MasterAgreementCompleteFlag *bool
	OfficePhone                 *string
	PostalCode                  *string
	ServiceProviderAddendumFlag *bool
	State                       *string
	SurveyResponses             []Survey_Response
	VatId                       *string
}

type User_Customer_Prospect_Type struct {
	Entity

	CreateDate  *time.Time
	Description *string
	Id          *int
	KeyName     *string
	ModifyDate  *time.Time
	Name        *string
}

type User_Customer_Security_Answer struct {
	Entity

	Answer     *string
	Id         *int
	Question   *User_Security_Question
	QuestionId *int
	User       *User_Customer
	UserId     *int
}

type User_Customer_Status struct {
	Entity

	Id      *int
	KeyName *string
	Name    *string
}

type User_Employee struct {
	User_Interface

	ActionCount                    *uint
	Actions                        []User_Permission_Action
	ChatTranscript                 []Ticket_Chat
	ChatTranscriptCount            *uint
	DisplayName                    *string
	Email                          *string
	EmployeeDepartment             *User_Employee_Department
	EmployeeDepartmentId           *int
	FirstName                      *string
	LastName                       *string
	LayoutProfileCount             *uint
	LayoutProfiles                 []Layout_Profile
	MetricTrackingObject           *Metric_Tracking_Object
	OfficePhone                    *string
	RoleCount                      *uint
	Roles                          []User_Permission_Role
	TicketActivities               []Ticket_Activity
	TicketActivityCount            *uint
	TicketAttachmentReferenceCount *uint
	TicketAttachmentReferences     []Ticket_Attachment
	Username                       *string
}

type User_Employee_Department struct {
	Entity

	Name *string
}

type User_External_Binding struct {
	Entity

	Active         *bool
	AttributeCount *uint
	Attributes     []User_External_Binding_Attribute
	BillingItem    *Billing_Item
	CreateDate     *time.Time
	ExternalId     *string
	Id             *int
	Note           *string
	Password       *string
	Type           *User_External_Binding_Type
	TypeId         *int
	UserId         *int
	Vendor         *User_External_Binding_Vendor
	VendorId       *int
}

type User_External_Binding_Attribute struct {
	Entity

	ExternalBinding *User_External_Binding
	Value           *string
}

type User_External_Binding_Type struct {
	Entity

	KeyName *string
	Name    *string
}

type User_External_Binding_Vendor struct {
	Entity

	Id      *int
	KeyName *string
	Name    *string
}

type User_Interface struct {
	Entity
}

type User_Permission_Action struct {
	Entity

	CreateDate  *time.Time
	Description *string
	Id          *int
	KeyName     *string
	ModifyDate  *time.Time
	Name        *string
}

type User_Permission_Group struct {
	Entity

	Account        *Account
	AccountId      *int
	ActionCount    *uint
	Actions        []User_Permission_Action
	CreateDate     *time.Time
	Description    *string
	ExpirationDate *time.Time
	Id             *int
	ModifyDate     *time.Time
	Name           *string
	RoleCount      *uint
	Roles          []User_Permission_Role
	Type           *User_Permission_Group_Type
	TypeId         *int
}

type User_Permission_Group_Type struct {
	Entity

	CreateDate *time.Time
	GroupCount *uint
	Groups     []User_Permission_Group
	Id         *int
	KeyName    *string
	ModifyDate *time.Time
	Name       *string
}

type User_Permission_Role struct {
	Entity

	Account            *Account
	AccountId          *int
	ActionCount        *uint
	Actions            []User_Permission_Action
	CreateDate         *time.Time
	Description        *string
	GroupCount         *uint
	Groups             []User_Permission_Group
	Id                 *int
	ModifyDate         *time.Time
	Name               *string
	NewUserDefaultFlag *int
	SystemFlag         *int
	UserCount          *uint
	Users              []User_Customer
}

type User_Preference struct {
	Entity

	Description *string
	Type        *User_Preference_Type
	Value       *string
}

type User_Preference_Type struct {
	Entity

	Description  *string
	KeyName      *string
	Name         *string
	ValueExample *string
}

type User_Security_Question struct {
	Entity

	DisplayOrder *int
	Id           *int
	Question     *string
	Viewable     *int
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

	BillingItem             *Billing_Item_Virtual_Disk_Image
	BlockDeviceCount        *uint
	BlockDevices            []Virtual_Guest_Block_Device
	BootableVolumeFlag      *bool
	Capacity                *int
	Checksum                *string
	CoalescedDiskImageCount *uint
	CoalescedDiskImages     []Virtual_Disk_Image
	CopyOnWriteFlag         *bool
	CreateDate              *time.Time
	Description             *string
	Id                      *int
	LocalDiskFlag           *bool
	MetadataFlag            *bool
	ModifyDate              *time.Time
	Name                    *string
	ParentId                *int
	SoftwareReferenceCount  *uint
	SoftwareReferences      []Virtual_Disk_Image_Software
	SourceDiskImage         *Virtual_Disk_Image
	StorageRepository       *Virtual_Storage_Repository
	StorageRepositoryId     *int
	StorageRepositoryType   *Virtual_Storage_Repository_Type
	TemplateBlockDevice     *Virtual_Guest_Block_Device_Template
	Type                    *Virtual_Disk_Image_Type
	TypeId                  *int
	Units                   *string
	Uuid                    *string
}

type Virtual_Disk_Image_Software struct {
	Entity

	DiskImage             *Virtual_Disk_Image
	Id                    *int
	PasswordCount         *uint
	Passwords             []Virtual_Disk_Image_Software_Password
	SoftwareDescription   *Software_Description
	SoftwareDescriptionId *int
}

type Virtual_Disk_Image_Software_Password struct {
	Entity

	Password *string
	Software *Virtual_Disk_Image_Software
	Username *string
}

type Virtual_Disk_Image_Type struct {
	Entity

	Description *string
	KeyName     *string
	Name        *string
}

type Virtual_Guest struct {
	Entity

	Account                                   *Account
	AccountId                                 *int
	AccountOwnedPoolFlag                      *bool
	ActiveNetworkMonitorIncident              []Network_Monitor_Version1_Incident
	ActiveNetworkMonitorIncidentCount         *uint
	ActiveTicketCount                         *uint
	ActiveTickets                             []Ticket
	ActiveTransaction                         *Provisioning_Version1_Transaction
	ActiveTransactionCount                    *uint
	ActiveTransactions                        []Provisioning_Version1_Transaction
	AllowedHost                               *Network_Storage_Allowed_Host
	AllowedNetworkStorage                     []Network_Storage
	AllowedNetworkStorageCount                *uint
	AllowedNetworkStorageReplicaCount         *uint
	AllowedNetworkStorageReplicas             []Network_Storage
	AntivirusSpywareSoftwareComponent         *Software_Component
	ApplicationDeliveryController             *Network_Application_Delivery_Controller
	AttributeCount                            *uint
	Attributes                                []Virtual_Guest_Attribute
	AvailableMonitoring                       []Network_Monitor_Version1_Query_Host_Stratum
	AvailableMonitoringCount                  *uint
	AverageDailyPrivateBandwidthUsage         *float64
	AverageDailyPublicBandwidthUsage          *float64
	BackendNetworkComponentCount              *uint
	BackendNetworkComponents                  []Virtual_Guest_Network_Component
	BackendRouterCount                        *uint
	BackendRouters                            []Hardware
	BandwidthAllocation                       *float64
	BandwidthAllotmentDetail                  *Network_Bandwidth_Version1_Allotment_Detail
	BillingCycleBandwidthUsage                []Network_Bandwidth_Usage
	BillingCycleBandwidthUsageCount           *uint
	BillingCyclePrivateBandwidthUsage         *Network_Bandwidth_Usage
	BillingCyclePublicBandwidthUsage          *Network_Bandwidth_Usage
	BillingItem                               *Billing_Item_Virtual_Guest
	BlockCancelBecauseDisconnectedFlag        *bool
	BlockDeviceCount                          *uint
	BlockDeviceTemplateGroup                  *Virtual_Guest_Block_Device_Template_Group
	BlockDevices                              []Virtual_Guest_Block_Device
	ConsoleIpAddressFlag                      *bool
	ConsoleIpAddressRecord                    *Virtual_Guest_Network_Component_IpAddress
	ContinuousDataProtectionSoftwareComponent *Software_Component
	ControlPanel                              *Software_Component
	CreateDate                                *time.Time
	CurrentBandwidthSummary                   *Metric_Tracking_Object_Bandwidth_Summary
	Datacenter                                *Location
	DedicatedAccountHostOnlyFlag              *bool
	Domain                                    *string
	EvaultNetworkStorage                      []Network_Storage
	EvaultNetworkStorageCount                 *uint
	FirewallServiceComponent                  *Network_Component_Firewall
	FrontendNetworkComponentCount             *uint
	FrontendNetworkComponents                 []Virtual_Guest_Network_Component
	FrontendRouters                           *Hardware
	FullyQualifiedDomainName                  *string
	GlobalIdentifier                          *string
	GuestBootParameter                        *Virtual_Guest_Boot_Parameter
	Host                                      *Virtual_Host
	HostIpsSoftwareComponent                  *Software_Component
	Hostname                                  *string
	HourlyBillingFlag                         *bool
	Id                                        *int
	InboundPrivateBandwidthUsage              *float64
	InboundPublicBandwidthUsage               *float64
	InternalTagReferenceCount                 *uint
	InternalTagReferences                     []Tag_Reference
	LastKnownPowerState                       *Virtual_Guest_Power_State
	LastOperatingSystemReload                 *Provisioning_Version1_Transaction
	LastPowerStateId                          *int
	LastTransaction                           *Provisioning_Version1_Transaction
	LastVerifiedDate                          *time.Time
	LatestNetworkMonitorIncident              *Network_Monitor_Version1_Incident
	LocalDiskFlag                             *bool
	Location                                  *Location
	ManagedResourceFlag                       *bool
	MaxCpu                                    *int
	MaxCpuUnits                               *string
	MaxMemory                                 *int
	MetricPollDate                            *time.Time
	MetricTrackingObject                      *Metric_Tracking_Object
	MetricTrackingObjectId                    *int
	ModifyDate                                *time.Time
	MonitoringAgentCount                      *uint
	MonitoringAgents                          []Monitoring_Agent
	MonitoringRobot                           *Monitoring_Robot
	MonitoringServiceComponent                *Network_Monitor_Version1_Query_Host_Stratum
	MonitoringServiceEligibilityFlag          *bool
	MonitoringServiceFlag                     *bool
	MonitoringUserNotification                []User_Customer_Notification_Virtual_Guest
	MonitoringUserNotificationCount           *uint
	NetworkComponentCount                     *uint
	NetworkComponents                         []Virtual_Guest_Network_Component
	NetworkMonitorCount                       *uint
	NetworkMonitorIncidentCount               *uint
	NetworkMonitorIncidents                   []Network_Monitor_Version1_Incident
	NetworkMonitors                           []Network_Monitor_Version1_Query_Host
	NetworkStorage                            []Network_Storage
	NetworkStorageCount                       *uint
	NetworkVlanCount                          *uint
	NetworkVlans                              []Network_Vlan
	Notes                                     *string
	OpenCancellationTicket                    *Ticket
	OperatingSystem                           *Software_Component_OperatingSystem
	OperatingSystemReferenceCode              *string
	OrderedPackageId                          *string
	OutboundPrivateBandwidthUsage             *float64
	OutboundPublicBandwidthUsage              *float64
	OverBandwidthAllocationFlag               *int
	PostInstallScriptUri                      *string
	PowerState                                *Virtual_Guest_Power_State
	PrimaryBackendIpAddress                   *string
	PrimaryBackendNetworkComponent            *Virtual_Guest_Network_Component
	PrimaryIpAddress                          *string
	PrimaryNetworkComponent                   *Virtual_Guest_Network_Component
	PrivateNetworkOnlyFlag                    *bool
	ProjectedOverBandwidthAllocationFlag      *int
	ProjectedPublicBandwidthUsage             *float64
	ProvisionDate                             *time.Time
	RecentEventCount                          *uint
	RecentEvents                              []Notification_Occurrence_Event
	RegionalGroup                             *Location_Group_Regional
	RegionalInternetRegistry                  *Network_Regional_Internet_Registry
	ScaleAssetCount                           *uint
	ScaleAssets                               []Scale_Asset
	ScaleMember                               *Scale_Member_Virtual_Guest
	ScaledFlag                                *bool
	SecurityScanRequestCount                  *uint
	SecurityScanRequests                      []Network_Security_Scanner_Request
	ServerRoom                                *Location
	SoftwareComponentCount                    *uint
	SoftwareComponents                        []Software_Component
	SshKeyCount                               *uint
	SshKeys                                   []Security_Ssh_Key
	StartCpus                                 *int
	Status                                    *Virtual_Guest_Status
	StatusId                                  *int
	SupplementalCreateObjectOptions           *Virtual_Guest_SupplementalCreateObjectOptions
	TagReferenceCount                         *uint
	TagReferences                             []Tag_Reference
	UpgradeRequest                            *Product_Upgrade_Request
	UserCount                                 *uint
	UserData                                  []Virtual_Guest_Attribute
	UserDataCount                             *uint
	Users                                     []User_Customer
	Uuid                                      *string
	VirtualRack                               *Network_Bandwidth_Version1_Allotment
	VirtualRackId                             *int
	VirtualRackName                           *string
}

type Virtual_Guest_Attribute struct {
	Entity

	Guest *Virtual_Guest
	Type  *Virtual_Guest_Attribute_Type
	Value *string
}

type Virtual_Guest_Attribute_Type struct {
	Entity

	Keyname *string
	Name    *string
}

type Virtual_Guest_Attribute_UserData struct {
	Virtual_Guest_Attribute
}

type Virtual_Guest_Block_Device struct {
	Entity

	BootableFlag *int
	CreateDate   *time.Time
	Device       *string
	DiskImage    *Virtual_Disk_Image
	DiskImageId  *int
	Guest        *Virtual_Guest
	GuestId      *int
	HotPlugFlag  *int
	Id           *int
	ModifyDate   *time.Time
	MountMode    *string
	MountType    *string
	Status       *Virtual_Guest_Block_Device_Status
	StatusId     *int
	Uuid         *string
}

type Virtual_Guest_Block_Device_Status struct {
	Entity

	KeyName *string
	Name    *string
}

type Virtual_Guest_Block_Device_Template struct {
	Entity

	Device      *string
	DiskImage   *Virtual_Disk_Image
	DiskImageId *int
	DiskSpace   *float64
	Group       *Virtual_Guest_Block_Device_Template_Group
	GroupId     *int
	Id          *int
	Units       *string
}

type Virtual_Guest_Block_Device_Template_Group struct {
	Entity

	Account                    *Account
	AccountContactCount        *uint
	AccountContacts            []Account_Contact
	AccountId                  *int
	AccountReferenceCount      *uint
	AccountReferences          []Virtual_Guest_Block_Device_Template_Group_Accounts
	BlockDeviceCount           *uint
	BlockDevices               []Virtual_Guest_Block_Device_Template
	BlockDevicesDiskSpaceTotal *float64
	Children                   []Virtual_Guest_Block_Device_Template_Group
	ChildrenCount              *uint
	CreateDate                 *time.Time
	Datacenter                 *Location
	DatacenterCount            *uint
	Datacenters                []Location
	FlexImageFlag              *bool
	GlobalIdentifier           *string
	Id                         *int
	ImageType                  *string
	ImageTypeKeyName           *string
	Name                       *string
	Note                       *string
	Parent                     *Virtual_Guest_Block_Device_Template_Group
	ParentId                   *int
	PublicFlag                 *int
	SshKeyCount                *uint
	SshKeys                    []Security_Ssh_Key
	Status                     *Virtual_Guest_Block_Device_Template_Group_Status
	StatusId                   *int
	StorageRepository          *Virtual_Storage_Repository
	Summary                    *string
	TagReferenceCount          *uint
	TagReferences              []Tag_Reference
	Transaction                *Provisioning_Version1_Transaction
	TransactionId              *int
	UserRecordId               *int
}

type Virtual_Guest_Block_Device_Template_Group_Accounts struct {
	Entity

	Account    *Account
	AccountId  *int
	CreateDate *time.Time
	Group      *Virtual_Guest_Block_Device_Template_Group
	GroupId    *int
}

type Virtual_Guest_Block_Device_Template_Group_Status struct {
	Entity

	Description *string
	KeyName     *string
	Name        *string
}

type Virtual_Guest_Boot_Parameter struct {
	Entity

	CreateDate               *time.Time
	Guest                    *Virtual_Guest
	GuestBootParameterType   *Virtual_Guest_Boot_Parameter_Type
	GuestBootParameterTypeId *int
	GuestId                  *int
	Id                       *int
	ModifyDate               *time.Time
}

type Virtual_Guest_Boot_Parameter_Type struct {
	Entity

	BootOption  *string
	CreateDate  *time.Time
	Description *string
	Id          *int
	KeyName     *string
	ModifyDate  *time.Time
	Name        *string
	Value       *string
}

type Virtual_Guest_Network_Component struct {
	Entity

	CreateDate                     *time.Time
	Guest                          *Virtual_Guest
	GuestId                        *int
	HighAvailabilityFirewallFlag   *bool
	Id                             *int
	IpAddressBindingCount          *uint
	IpAddressBindings              []Virtual_Guest_Network_Component_IpAddress
	MacAddress                     *string
	MaxSpeed                       *int
	ModifyDate                     *time.Time
	Name                           *string
	NetworkComponentFirewall       *Network_Component_Firewall
	NetworkId                      *int
	NetworkVlan                    *Network_Vlan
	Port                           *int
	PrimaryIpAddress               *string
	PrimaryIpAddressRecord         *Network_Subnet_IpAddress
	PrimarySubnet                  *Network_Subnet
	PrimaryVersion6IpAddressRecord *Network_Subnet_IpAddress
	Router                         *Hardware_Router
	Speed                          *int
	Status                         *string
	SubnetCount                    *uint
	Subnets                        []Network_Subnet
	Uuid                           *string
}

type Virtual_Guest_Network_Component_IpAddress struct {
	Entity

	IpAddress        *Network_Subnet_IpAddress
	IpAddressId      *int
	NetworkComponent *Virtual_Guest_Network_Component
	Port             *int
	Type             *string
}

type Virtual_Guest_Power_State struct {
	Entity

	Description *string
	KeyName     *string
	Name        *string
}

type Virtual_Guest_Status struct {
	Entity

	KeyName *string
	Name    *string
}

type Virtual_Guest_SupplementalCreateObjectOptions struct {
	Entity

	ImmediateApprovalOnlyFlag *bool
	PostInstallScriptUri      *string
}

type Virtual_Host struct {
	Entity

	Account                  *Account
	AccountId                *int
	BilledPerGuestFlag       *bool
	BilledPerMemoryUsageFlag *bool
	CreateDate               *time.Time
	Description              *string
	EnabledFlag              *int
	GuestCount               *uint
	Guests                   []Virtual_Guest
	Hardware                 *Hardware_Server
	HardwareId               *int
	Id                       *int
	MetricTrackingObject     *Metric_Tracking_Object
	ModifyDate               *time.Time
	Name                     *string
	PhysicalMemoryCapacity   *int
	Uuid                     *string
}

type Virtual_Storage_Repository struct {
	Entity

	Account                *Account
	BillingItem            *Billing_Item
	Capacity               *float64
	Datacenter             *Location
	Description            *string
	DiskImageCount         *uint
	DiskImages             []Virtual_Disk_Image
	GuestCount             *uint
	Guests                 []Virtual_Guest
	Id                     *int
	MetricTrackingObject   *Metric_Tracking_Object_Virtual_Storage_Repository
	Name                   *string
	PublicFlag             *int
	PublicImageBillingItem *Billing_Item
	Type                   *Virtual_Storage_Repository_Type
	TypeId                 *int
}

type Virtual_Storage_Repository_Type struct {
	Entity

	Description            *string
	Name                   *string
	StorageRepositories    []Virtual_Storage_Repository
	StorageRepositoryCount *uint
}


