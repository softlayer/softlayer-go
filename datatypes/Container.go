package datatypes

import "time"

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
