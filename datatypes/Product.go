package datatypes

import "time"

type Product_Catalog struct {
	Entity

	BrandCount   *uint                `json:"brandCount,omitempty"`
	Brands       []Brand              `json:"brands,omitempty"`
	PackageCount *uint                `json:"packageCount,omitempty"`
	Packages     []Product_Package    `json:"packages,omitempty"`
	PriceCount   *uint                `json:"priceCount,omitempty"`
	Prices       []Product_Item_Price `json:"prices,omitempty"`
	ProductCount *uint                `json:"productCount,omitempty"`
	Products     []Product_Item       `json:"products,omitempty"`
}

type Product_Catalog_Item_Price struct {
	Entity

	Catalog    *Product_Catalog    `json:"catalog,omitempty"`
	CatalogId  *int                `json:"catalogId,omitempty"`
	CreateDate *time.Time          `json:"createDate,omitempty"`
	ModifyDate *time.Time          `json:"modifyDate,omitempty"`
	Price      *Product_Item_Price `json:"price,omitempty"`
	PriceId    *int                `json:"priceId,omitempty"`
}

type Product_Item struct {
	Entity

	ActivePresaleEventCount         *uint                             `json:"activePresaleEventCount,omitempty"`
	ActivePresaleEvents             []Sales_Presale_Event             `json:"activePresaleEvents,omitempty"`
	ActiveUsagePriceCount           *uint                             `json:"activeUsagePriceCount,omitempty"`
	ActiveUsagePrices               []Product_Item_Price              `json:"activeUsagePrices,omitempty"`
	AttributeCount                  *uint                             `json:"attributeCount,omitempty"`
	Attributes                      []Product_Item_Attribute          `json:"attributes,omitempty"`
	AvailabilityAttributeCount      *uint                             `json:"availabilityAttributeCount,omitempty"`
	AvailabilityAttributes          []Product_Item_Attribute          `json:"availabilityAttributes,omitempty"`
	BillingType                     *string                           `json:"billingType,omitempty"`
	Bundle                          []Product_Item_Bundles            `json:"bundle,omitempty"`
	BundleCount                     *uint                             `json:"bundleCount,omitempty"`
	Capacity                        *float64                          `json:"capacity,omitempty"`
	CapacityRestrictedProductFlag   *bool                             `json:"capacityRestrictedProductFlag,omitempty"`
	Categories                      []Product_Item_Category           `json:"categories,omitempty"`
	CategoryCount                   *uint                             `json:"categoryCount,omitempty"`
	ConfigurationTemplateCount      *uint                             `json:"configurationTemplateCount,omitempty"`
	ConfigurationTemplates          []Configuration_Template          `json:"configurationTemplates,omitempty"`
	ConflictCount                   *uint                             `json:"conflictCount,omitempty"`
	Conflicts                       []Product_Item_Resource_Conflict  `json:"conflicts,omitempty"`
	CoreRestrictedItemFlag          *bool                             `json:"coreRestrictedItemFlag,omitempty"`
	Description                     *string                           `json:"description,omitempty"`
	DowngradeItem                   *Product_Item                     `json:"downgradeItem,omitempty"`
	DowngradeItemCount              *uint                             `json:"downgradeItemCount,omitempty"`
	DowngradeItems                  []Product_Item                    `json:"downgradeItems,omitempty"`
	GlobalCategoryConflictCount     *uint                             `json:"globalCategoryConflictCount,omitempty"`
	GlobalCategoryConflicts         []Product_Item_Resource_Conflict  `json:"globalCategoryConflicts,omitempty"`
	HardwareGenericComponentModel   *Hardware_Component_Model_Generic `json:"hardwareGenericComponentModel,omitempty"`
	HideFromPortalFlag              *bool                             `json:"hideFromPortalFlag,omitempty"`
	Id                              *int                              `json:"id,omitempty"`
	Inventory                       []Product_Package_Inventory       `json:"inventory,omitempty"`
	InventoryCount                  *uint                             `json:"inventoryCount,omitempty"`
	IsEngineeredServerProduct       *bool                             `json:"isEngineeredServerProduct,omitempty"`
	ItemCategory                    *Product_Item_Category            `json:"itemCategory,omitempty"`
	ItemTaxCategoryId               *int                              `json:"itemTaxCategoryId,omitempty"`
	KeyName                         *string                           `json:"keyName,omitempty"`
	LocationConflictCount           *uint                             `json:"locationConflictCount,omitempty"`
	LocationConflicts               []Product_Item_Resource_Conflict  `json:"locationConflicts,omitempty"`
	LongDescription                 *string                           `json:"longDescription,omitempty"`
	ObjectStorageItemFlag           *bool                             `json:"objectStorageItemFlag,omitempty"`
	PackageCount                    *uint                             `json:"packageCount,omitempty"`
	Packages                        []Product_Package                 `json:"packages,omitempty"`
	PhysicalCoreCapacity            *string                           `json:"physicalCoreCapacity,omitempty"`
	PresaleEventCount               *uint                             `json:"presaleEventCount,omitempty"`
	PresaleEvents                   []Sales_Presale_Event             `json:"presaleEvents,omitempty"`
	PriceCount                      *uint                             `json:"priceCount,omitempty"`
	Prices                          []Product_Item_Price              `json:"prices,omitempty"`
	RequirementCount                *uint                             `json:"requirementCount,omitempty"`
	Requirements                    []Product_Item_Requirement        `json:"requirements,omitempty"`
	SoftwareDescription             *Software_Description             `json:"softwareDescription,omitempty"`
	SoftwareDescriptionId           *int                              `json:"softwareDescriptionId,omitempty"`
	TaxCategory                     *Product_Item_Tax_Category        `json:"taxCategory,omitempty"`
	ThirdPartyPolicyAssignmentCount *uint                             `json:"thirdPartyPolicyAssignmentCount,omitempty"`
	ThirdPartyPolicyAssignments     []Product_Item_Policy_Assignment  `json:"thirdPartyPolicyAssignments,omitempty"`
	ThirdPartySupportVendor         *string                           `json:"thirdPartySupportVendor,omitempty"`
	TotalPhysicalCoreCapacity       *int                              `json:"totalPhysicalCoreCapacity,omitempty"`
	TotalPhysicalCoreCount          *int                              `json:"totalPhysicalCoreCount,omitempty"`
	TotalProcessorCapacity          *int                              `json:"totalProcessorCapacity,omitempty"`
	Units                           *string                           `json:"units,omitempty"`
	UpgradeItem                     *Product_Item                     `json:"upgradeItem,omitempty"`
	UpgradeItemCount                *uint                             `json:"upgradeItemCount,omitempty"`
	UpgradeItemId                   *int                              `json:"upgradeItemId,omitempty"`
	UpgradeItems                    []Product_Item                    `json:"upgradeItems,omitempty"`
}

type Product_Item_Attribute struct {
	Entity

	AttributeType        *Product_Item_Attribute_Type `json:"attributeType,omitempty"`
	AttributeTypeKeyName *string                      `json:"attributeTypeKeyName,omitempty"`
	Id                   *int                         `json:"id,omitempty"`
	Item                 *Product_Item                `json:"item,omitempty"`
	ItemAttributeTypeId  *int                         `json:"itemAttributeTypeId,omitempty"`
	ItemId               *int                         `json:"itemId,omitempty"`
	Value                *string                      `json:"value,omitempty"`
}

type Product_Item_Attribute_Type struct {
	Entity

	KeyName *string `json:"keyName,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type Product_Item_Billing_Type struct {
	Entity

	Name *string `json:"name,omitempty"`
}

type Product_Item_Bundles struct {
	Entity

	BundleItem   *Product_Item          `json:"bundleItem,omitempty"`
	BundleItemId *int                   `json:"bundleItemId,omitempty"`
	Category     *Product_Item_Category `json:"category,omitempty"`
	Id           *int                   `json:"id,omitempty"`
	ItemPrice    *Product_Item_Price    `json:"itemPrice,omitempty"`
	ItemPriceId  *int                   `json:"itemPriceId,omitempty"`
}

type Product_Item_Category struct {
	Entity

	BillingItemCount          *uint                                     `json:"billingItemCount,omitempty"`
	BillingItems              []Billing_Item                            `json:"billingItems,omitempty"`
	CategoryCode              *string                                   `json:"categoryCode,omitempty"`
	Group                     *Product_Item_Category_Group              `json:"group,omitempty"`
	GroupCount                *uint                                     `json:"groupCount,omitempty"`
	Groups                    []Product_Package_Item_Category_Group     `json:"groups,omitempty"`
	Id                        *int                                      `json:"id,omitempty"`
	Name                      *string                                   `json:"name,omitempty"`
	OrderOptionCount          *uint                                     `json:"orderOptionCount,omitempty"`
	OrderOptions              []Product_Item_Category_Order_Option_Type `json:"orderOptions,omitempty"`
	PackageConfigurationCount *uint                                     `json:"packageConfigurationCount,omitempty"`
	PackageConfigurations     []Product_Package_Order_Configuration     `json:"packageConfigurations,omitempty"`
	PresetConfigurationCount  *uint                                     `json:"presetConfigurationCount,omitempty"`
	PresetConfigurations      []Product_Package_Preset_Configuration    `json:"presetConfigurations,omitempty"`
	QuantityLimit             *int                                      `json:"quantityLimit,omitempty"`
	QuestionCount             *uint                                     `json:"questionCount,omitempty"`
	QuestionReferenceCount    *uint                                     `json:"questionReferenceCount,omitempty"`
	QuestionReferences        []Product_Item_Category_Question_Xref     `json:"questionReferences,omitempty"`
	Questions                 []Product_Item_Category_Question          `json:"questions,omitempty"`
}

type Product_Item_Category_Group struct {
	Entity

	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Product_Item_Category_Order_Option_Type struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	Keyname     *string `json:"keyname,omitempty"`
	Name        *string `json:"name,omitempty"`
	Value       *string `json:"value,omitempty"`
}

type Product_Item_Category_Question struct {
	Entity

	AnswerValueExpression      *string                                    `json:"answerValueExpression,omitempty"`
	Description                *string                                    `json:"description,omitempty"`
	FieldType                  *Product_Item_Category_Question_Field_Type `json:"fieldType,omitempty"`
	FieldTypeId                *int                                       `json:"fieldTypeId,omitempty"`
	Id                         *int                                       `json:"id,omitempty"`
	ItemCategoryReferenceCount *uint                                      `json:"itemCategoryReferenceCount,omitempty"`
	ItemCategoryReferences     []Product_Item_Category_Question_Xref      `json:"itemCategoryReferences,omitempty"`
	KeyName                    *string                                    `json:"keyName,omitempty"`
	Question                   *string                                    `json:"question,omitempty"`
	ValueExample               *string                                    `json:"valueExample,omitempty"`
}

type Product_Item_Category_Question_Field_Type struct {
	Entity

	Id      *int    `json:"id,omitempty"`
	KeyName *string `json:"keyName,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type Product_Item_Category_Question_Xref struct {
	Entity

	Id             *int                            `json:"id,omitempty"`
	ItemCategory   *Product_Item_Category          `json:"itemCategory,omitempty"`
	ItemCategoryId *int                            `json:"itemCategoryId,omitempty"`
	LocationId     *int                            `json:"locationId,omitempty"`
	Question       *Product_Item_Category_Question `json:"question,omitempty"`
	QuestionId     *int                            `json:"questionId,omitempty"`
	Required       *bool                           `json:"required,omitempty"`
}

type Product_Item_Link_ThePlanet struct {
	Entity

	Item            *Product_Item     `json:"item,omitempty"`
	ServiceProvider *Service_Provider `json:"serviceProvider,omitempty"`
}

type Product_Item_Policy_Assignment struct {
	Entity

	Id         *int          `json:"id,omitempty"`
	PolicyName *string       `json:"policyName,omitempty"`
	Product    *Product_Item `json:"product,omitempty"`
	ProductId  *int          `json:"productId,omitempty"`
}

type Product_Item_Price struct {
	Entity

	AccountRestrictionCount    *uint                                     `json:"accountRestrictionCount,omitempty"`
	AccountRestrictions        []Product_Item_Price_Account_Restriction  `json:"accountRestrictions,omitempty"`
	AttributeCount             *uint                                     `json:"attributeCount,omitempty"`
	Attributes                 []Product_Item_Price_Attribute            `json:"attributes,omitempty"`
	BigDataOsJournalDiskFlag   *bool                                     `json:"bigDataOsJournalDiskFlag,omitempty"`
	BundleReferenceCount       *uint                                     `json:"bundleReferenceCount,omitempty"`
	BundleReferences           []Product_Item_Bundles                    `json:"bundleReferences,omitempty"`
	CapacityRestrictionMaximum *string                                   `json:"capacityRestrictionMaximum,omitempty"`
	CapacityRestrictionMinimum *string                                   `json:"capacityRestrictionMinimum,omitempty"`
	CapacityRestrictionType    *string                                   `json:"capacityRestrictionType,omitempty"`
	Categories                 []Product_Item_Category                   `json:"categories,omitempty"`
	CategoryCount              *uint                                     `json:"categoryCount,omitempty"`
	CurrentPriceFlag           *bool                                     `json:"currentPriceFlag,omitempty"`
	DefinedSoftwareLicenseFlag *bool                                     `json:"definedSoftwareLicenseFlag,omitempty"`
	HourlyRecurringFee         *float64                                  `json:"hourlyRecurringFee,omitempty"`
	Id                         *int                                      `json:"id,omitempty"`
	Inventory                  []Product_Package_Inventory               `json:"inventory,omitempty"`
	InventoryCount             *uint                                     `json:"inventoryCount,omitempty"`
	Item                       *Product_Item                             `json:"item,omitempty"`
	ItemId                     *int                                      `json:"itemId,omitempty"`
	LaborFee                   *float64                                  `json:"laborFee,omitempty"`
	LocationGroupId            *int                                      `json:"locationGroupId,omitempty"`
	OnSaleFlag                 *bool                                     `json:"onSaleFlag,omitempty"`
	OneTimeFee                 *float64                                  `json:"oneTimeFee,omitempty"`
	OneTimeFeeTax              *float64                                  `json:"oneTimeFeeTax,omitempty"`
	OrderOptions               []Product_Item_Category_Order_Option_Type `json:"orderOptions,omitempty"`
	OrderPremiumCount          *uint                                     `json:"orderPremiumCount,omitempty"`
	OrderPremiums              []Product_Item_Price_Premium              `json:"orderPremiums,omitempty"`
	PackageCount               *uint                                     `json:"packageCount,omitempty"`
	PackageReferenceCount      *uint                                     `json:"packageReferenceCount,omitempty"`
	PackageReferences          []Product_Package_Item_Prices             `json:"packageReferences,omitempty"`
	Packages                   []Product_Package                         `json:"packages,omitempty"`
	PresetConfigurationCount   *uint                                     `json:"presetConfigurationCount,omitempty"`
	PresetConfigurations       []Product_Package_Preset_Configuration    `json:"presetConfigurations,omitempty"`
	PricingLocationGroup       *Location_Group_Pricing                   `json:"pricingLocationGroup,omitempty"`
	ProratedRecurringFee       *float64                                  `json:"proratedRecurringFee,omitempty"`
	ProratedRecurringFeeTax    *float64                                  `json:"proratedRecurringFeeTax,omitempty"`
	Quantity                   *int                                      `json:"quantity,omitempty"`
	RecurringFee               *float64                                  `json:"recurringFee,omitempty"`
	RecurringFeeTax            *float64                                  `json:"recurringFeeTax,omitempty"`
	RequiredCoreCount          *int                                      `json:"requiredCoreCount,omitempty"`
	SetupFee                   *float64                                  `json:"setupFee,omitempty"`
	Sort                       *int                                      `json:"sort,omitempty"`
	UsageRate                  *float64                                  `json:"usageRate,omitempty"`
}

type Product_Item_Price_Account_Restriction struct {
	Entity

	Account     *Account            `json:"account,omitempty"`
	AccountId   *int                `json:"accountId,omitempty"`
	Id          *int                `json:"id,omitempty"`
	ItemPrice   *Product_Item_Price `json:"itemPrice,omitempty"`
	ItemPriceId *int                `json:"itemPriceId,omitempty"`
}

type Product_Item_Price_Attribute struct {
	Entity

	Id                       *int                               `json:"id,omitempty"`
	ItemPrice                *Product_Item_Price                `json:"itemPrice,omitempty"`
	ItemPriceAttributeType   *Product_Item_Price_Attribute_Type `json:"itemPriceAttributeType,omitempty"`
	ItemPriceAttributeTypeId *int                               `json:"itemPriceAttributeTypeId,omitempty"`
	ItemPriceId              *int                               `json:"itemPriceId,omitempty"`
	Value                    *string                            `json:"value,omitempty"`
}

type Product_Item_Price_Attribute_Type struct {
	Entity

	Id      *int    `json:"id,omitempty"`
	Keyname *string `json:"keyname,omitempty"`
}

type Product_Item_Price_Premium struct {
	Entity

	HourlyModifier  *float64            `json:"hourlyModifier,omitempty"`
	ItemPrice       *Product_Item_Price `json:"itemPrice,omitempty"`
	ItemPriceId     *int                `json:"itemPriceId,omitempty"`
	Location        *Location           `json:"location,omitempty"`
	LocationId      *int                `json:"locationId,omitempty"`
	MonthlyModifier *float64            `json:"monthlyModifier,omitempty"`
	Package         *Product_Package    `json:"package,omitempty"`
	PackageId       *int                `json:"packageId,omitempty"`
}

type Product_Item_Requirement struct {
	Entity

	Id             *int          `json:"id,omitempty"`
	Item           *Product_Item `json:"item,omitempty"`
	ItemId         *int          `json:"itemId,omitempty"`
	Message        *string       `json:"message,omitempty"`
	Product        *Product_Item `json:"product,omitempty"`
	RequiredItemId *int          `json:"requiredItemId,omitempty"`
}

type Product_Item_Resource_Conflict struct {
	Entity

	Item            *Product_Item    `json:"item,omitempty"`
	ItemId          *int             `json:"itemId,omitempty"`
	Message         *string          `json:"message,omitempty"`
	Package         *Product_Package `json:"package,omitempty"`
	PackageId       *int             `json:"packageId,omitempty"`
	ResourceTableId *int             `json:"resourceTableId,omitempty"`
}

type Product_Item_Resource_Conflict_Item struct {
	Product_Item_Resource_Conflict

	Resource *Product_Item `json:"resource,omitempty"`
}

type Product_Item_Resource_Conflict_Item_Category struct {
	Product_Item_Resource_Conflict

	Resource *Product_Item_Category `json:"resource,omitempty"`
}

type Product_Item_Resource_Conflict_Location struct {
	Product_Item_Resource_Conflict

	Resource *Location `json:"resource,omitempty"`
}

type Product_Item_Tax_Category struct {
	Entity

	Id         *int           `json:"id,omitempty"`
	ItemCount  *uint          `json:"itemCount,omitempty"`
	Items      []Product_Item `json:"items,omitempty"`
	Name       *string        `json:"name,omitempty"`
	StatusFlag *int           `json:"statusFlag,omitempty"`
}

type Product_Order struct {
	Entity
}

type Product_Package struct {
	Entity

	AccountRestrictedCategories     []Product_Item_Category               `json:"accountRestrictedCategories,omitempty"`
	AccountRestrictedCategoryCount  *uint                                 `json:"accountRestrictedCategoryCount,omitempty"`
	AccountRestrictedPricesFlag     *bool                                 `json:"accountRestrictedPricesFlag,omitempty"`
	ActivePresetCount               *uint                                 `json:"activePresetCount,omitempty"`
	ActivePresets                   []Product_Package_Preset              `json:"activePresets,omitempty"`
	ActiveRamItemCount              *uint                                 `json:"activeRamItemCount,omitempty"`
	ActiveRamItems                  []Product_Item                        `json:"activeRamItems,omitempty"`
	ActiveServerItemCount           *uint                                 `json:"activeServerItemCount,omitempty"`
	ActiveServerItems               []Product_Item                        `json:"activeServerItems,omitempty"`
	ActiveSoftwareItemCount         *uint                                 `json:"activeSoftwareItemCount,omitempty"`
	ActiveSoftwareItems             []Product_Item                        `json:"activeSoftwareItems,omitempty"`
	ActiveUsagePriceCount           *uint                                 `json:"activeUsagePriceCount,omitempty"`
	ActiveUsagePrices               []Product_Item_Price                  `json:"activeUsagePrices,omitempty"`
	AdditionalServiceFlag           *bool                                 `json:"additionalServiceFlag,omitempty"`
	AttributeCount                  *uint                                 `json:"attributeCount,omitempty"`
	Attributes                      []Product_Package_Attribute           `json:"attributes,omitempty"`
	AvailableLocationCount          *uint                                 `json:"availableLocationCount,omitempty"`
	AvailableLocations              []Product_Package_Locations           `json:"availableLocations,omitempty"`
	AvailableStorageUnits           *uint                                 `json:"availableStorageUnits,omitempty"`
	Categories                      []Product_Item_Category               `json:"categories,omitempty"`
	Configuration                   []Product_Package_Order_Configuration `json:"configuration,omitempty"`
	ConfigurationCount              *uint                                 `json:"configurationCount,omitempty"`
	DefaultRamItemCount             *uint                                 `json:"defaultRamItemCount,omitempty"`
	DefaultRamItems                 []Product_Item                        `json:"defaultRamItems,omitempty"`
	DeploymentCount                 *uint                                 `json:"deploymentCount,omitempty"`
	DeploymentNodeType              *string                               `json:"deploymentNodeType,omitempty"`
	DeploymentPackageCount          *uint                                 `json:"deploymentPackageCount,omitempty"`
	DeploymentPackages              []Product_Package                     `json:"deploymentPackages,omitempty"`
	DeploymentType                  *string                               `json:"deploymentType,omitempty"`
	Deployments                     []Product_Package                     `json:"deployments,omitempty"`
	Description                     *string                               `json:"description,omitempty"`
	DisallowCustomDiskPartitions    *bool                                 `json:"disallowCustomDiskPartitions,omitempty"`
	FirstOrderStep                  *Product_Package_Order_Step           `json:"firstOrderStep,omitempty"`
	FirstOrderStepId                *int                                  `json:"firstOrderStepId,omitempty"`
	GatewayApplianceFlag            *bool                                 `json:"gatewayApplianceFlag,omitempty"`
	GpuFlag                         *bool                                 `json:"gpuFlag,omitempty"`
	HourlyBillingAvailableFlag      *bool                                 `json:"hourlyBillingAvailableFlag,omitempty"`
	Id                              *int                                  `json:"id,omitempty"`
	IsActive                        *int                                  `json:"isActive,omitempty"`
	ItemConflictCount               *uint                                 `json:"itemConflictCount,omitempty"`
	ItemConflicts                   []Product_Item_Resource_Conflict      `json:"itemConflicts,omitempty"`
	ItemCount                       *uint                                 `json:"itemCount,omitempty"`
	ItemLocationConflictCount       *uint                                 `json:"itemLocationConflictCount,omitempty"`
	ItemLocationConflicts           []Product_Item_Resource_Conflict      `json:"itemLocationConflicts,omitempty"`
	ItemPriceCount                  *uint                                 `json:"itemPriceCount,omitempty"`
	ItemPriceReferenceCount         *uint                                 `json:"itemPriceReferenceCount,omitempty"`
	ItemPriceReferences             []Product_Package_Item_Prices         `json:"itemPriceReferences,omitempty"`
	ItemPrices                      []Product_Item_Price                  `json:"itemPrices,omitempty"`
	Items                           []Product_Item                        `json:"items,omitempty"`
	KeyName                         *string                               `json:"keyName,omitempty"`
	LocationCount                   *uint                                 `json:"locationCount,omitempty"`
	Locations                       []Location                            `json:"locations,omitempty"`
	LowestServerPrice               *Product_Item_Price                   `json:"lowestServerPrice,omitempty"`
	MaximumPortSpeed                *uint                                 `json:"maximumPortSpeed,omitempty"`
	MinimumPortSpeed                *uint                                 `json:"minimumPortSpeed,omitempty"`
	MongoDbEngineeredFlag           *bool                                 `json:"mongoDbEngineeredFlag,omitempty"`
	Name                            *string                               `json:"name,omitempty"`
	OrderPremiumCount               *uint                                 `json:"orderPremiumCount,omitempty"`
	OrderPremiums                   []Product_Item_Price_Premium          `json:"orderPremiums,omitempty"`
	PreconfiguredFlag               *bool                                 `json:"preconfiguredFlag,omitempty"`
	PresetConfigurationRequiredFlag *bool                                 `json:"presetConfigurationRequiredFlag,omitempty"`
	PreventVlanSelectionFlag        *bool                                 `json:"preventVlanSelectionFlag,omitempty"`
	PrivateHostedCloudPackageFlag   *bool                                 `json:"privateHostedCloudPackageFlag,omitempty"`
	PrivateHostedCloudPackageType   *string                               `json:"privateHostedCloudPackageType,omitempty"`
	PrivateNetworkOnlyFlag          *bool                                 `json:"privateNetworkOnlyFlag,omitempty"`
	QuantaStorPackageFlag           *bool                                 `json:"quantaStorPackageFlag,omitempty"`
	RaidDiskRestrictionFlag         *bool                                 `json:"raidDiskRestrictionFlag,omitempty"`
	RedundantPowerFlag              *bool                                 `json:"redundantPowerFlag,omitempty"`
	RegionCount                     *uint                                 `json:"regionCount,omitempty"`
	Regions                         []Location_Region                     `json:"regions,omitempty"`
	ResourceGroupTemplate           *Resource_Group_Template              `json:"resourceGroupTemplate,omitempty"`
	SubDescription                  *string                               `json:"subDescription,omitempty"`
	TopLevelItemCategoryCode        *string                               `json:"topLevelItemCategoryCode,omitempty"`
	Type                            *Product_Package_Type                 `json:"type,omitempty"`
	UnitSize                        *int                                  `json:"unitSize,omitempty"`
}

type Product_Package_Attribute struct {
	Entity

	AttributeType *Product_Package_Attribute_Type `json:"attributeType,omitempty"`
	Package       *Product_Package                `json:"package,omitempty"`
	Value         *string                         `json:"value,omitempty"`
}

type Product_Package_Attribute_Type struct {
	Entity

	Description *string `json:"description,omitempty"`
	KeyName     *string `json:"keyName,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Product_Package_Inventory struct {
	Entity

	AvailableInventoryCount *int             `json:"availableInventoryCount,omitempty"`
	Item                    *Product_Item    `json:"item,omitempty"`
	ItemId                  *int             `json:"itemId,omitempty"`
	Location                *Location        `json:"location,omitempty"`
	LocationId              *int             `json:"locationId,omitempty"`
	ModifyDate              *time.Time       `json:"modifyDate,omitempty"`
	OverstockFlag           *int             `json:"overstockFlag,omitempty"`
	Package                 *Product_Package `json:"package,omitempty"`
	PackageId               *int             `json:"packageId,omitempty"`
}

type Product_Package_Item_Category_Group struct {
	Entity

	Category       *Product_Item_Category `json:"category,omitempty"`
	ItemCategoryId *int                   `json:"itemCategoryId,omitempty"`
	Package        *Product_Package       `json:"package,omitempty"`
	PackageId      *int                   `json:"packageId,omitempty"`
	PriceCount     *uint                  `json:"priceCount,omitempty"`
	Prices         []Product_Item_Price   `json:"prices,omitempty"`
	Sort           *int                   `json:"sort,omitempty"`
	Title          *string                `json:"title,omitempty"`
}

type Product_Package_Item_Prices struct {
	Entity

	Id          *int                `json:"id,omitempty"`
	ItemPrice   *Product_Item_Price `json:"itemPrice,omitempty"`
	ItemPriceId *int                `json:"itemPriceId,omitempty"`
	Package     *Product_Package    `json:"package,omitempty"`
	PackageId   *int                `json:"packageId,omitempty"`
}

type Product_Package_Items struct {
	Entity

	Id        *string          `json:"id,omitempty"`
	Item      *Product_Item    `json:"item,omitempty"`
	ItemId    *int             `json:"itemId,omitempty"`
	Package   *Product_Package `json:"package,omitempty"`
	PackageId *int             `json:"packageId,omitempty"`
}

type Product_Package_Locations struct {
	Entity

	DeliveryTimeInformation *string          `json:"deliveryTimeInformation,omitempty"`
	IsAvailable             *int             `json:"isAvailable,omitempty"`
	Location                *Location        `json:"location,omitempty"`
	LocationId              *int             `json:"locationId,omitempty"`
	Package                 *Product_Package `json:"package,omitempty"`
	PackageId               *int             `json:"packageId,omitempty"`
}

type Product_Package_Order_Configuration struct {
	Entity

	ErrorMessage   *string                     `json:"errorMessage,omitempty"`
	Id             *int                        `json:"id,omitempty"`
	IsRequired     *int                        `json:"isRequired,omitempty"`
	ItemCategory   *Product_Item_Category      `json:"itemCategory,omitempty"`
	ItemCategoryId *int                        `json:"itemCategoryId,omitempty"`
	OrderStepId    *int                        `json:"orderStepId,omitempty"`
	Package        *Product_Package            `json:"package,omitempty"`
	PackageId      *int                        `json:"packageId,omitempty"`
	Sort           *int                        `json:"sort,omitempty"`
	Step           *Product_Package_Order_Step `json:"step,omitempty"`
}

type Product_Package_Order_Step struct {
	Entity

	Id                         *int                              `json:"id,omitempty"`
	InclusivePreviousStepCount *uint                             `json:"inclusivePreviousStepCount,omitempty"`
	InclusivePreviousSteps     []Product_Package_Order_Step_Next `json:"inclusivePreviousSteps,omitempty"`
	NextStepCount              *uint                             `json:"nextStepCount,omitempty"`
	NextSteps                  []Product_Package_Order_Step_Next `json:"nextSteps,omitempty"`
	PreviousStepCount          *uint                             `json:"previousStepCount,omitempty"`
	PreviousSteps              []Product_Package_Order_Step_Next `json:"previousSteps,omitempty"`
	Step                       *string                           `json:"step,omitempty"`
}

type Product_Package_Order_Step_Next struct {
	Entity

	Id              *int                        `json:"id,omitempty"`
	NextOrderStepId *int                        `json:"nextOrderStepId,omitempty"`
	OrderStepId     *int                        `json:"orderStepId,omitempty"`
	Step            *Product_Package_Order_Step `json:"step,omitempty"`
}

type Product_Package_Preset struct {
	Entity

	AvailableStorageUnits          *uint                                        `json:"availableStorageUnits,omitempty"`
	Categories                     []Product_Item_Category                      `json:"categories,omitempty"`
	CategoryCount                  *uint                                        `json:"categoryCount,omitempty"`
	Configuration                  []Product_Package_Preset_Configuration       `json:"configuration,omitempty"`
	ConfigurationCount             *uint                                        `json:"configurationCount,omitempty"`
	Description                    *string                                      `json:"description,omitempty"`
	FixedConfigurationFlag         *bool                                        `json:"fixedConfigurationFlag,omitempty"`
	Id                             *int                                         `json:"id,omitempty"`
	IsActive                       *string                                      `json:"isActive,omitempty"`
	KeyName                        *string                                      `json:"keyName,omitempty"`
	LowestPresetServerPrice        *Product_Item_Price                          `json:"lowestPresetServerPrice,omitempty"`
	Name                           *string                                      `json:"name,omitempty"`
	Package                        *Product_Package                             `json:"package,omitempty"`
	PackageConfiguration           []Product_Package_Order_Configuration        `json:"packageConfiguration,omitempty"`
	PackageConfigurationCount      *uint                                        `json:"packageConfigurationCount,omitempty"`
	PackageId                      *int                                         `json:"packageId,omitempty"`
	PriceCount                     *uint                                        `json:"priceCount,omitempty"`
	Prices                         []Product_Item_Price                         `json:"prices,omitempty"`
	StorageGroupTemplateArrayCount *uint                                        `json:"storageGroupTemplateArrayCount,omitempty"`
	StorageGroupTemplateArrays     []Configuration_Storage_Group_Template_Group `json:"storageGroupTemplateArrays,omitempty"`
	TotalMinimumHourlyFee          *float64                                     `json:"totalMinimumHourlyFee,omitempty"`
	TotalMinimumRecurringFee       *float64                                     `json:"totalMinimumRecurringFee,omitempty"`
}

type Product_Package_Preset_Attribute struct {
	Entity

	AttributeType   *Product_Package_Preset_Attribute_Type `json:"attributeType,omitempty"`
	AttributeTypeId *int                                   `json:"attributeTypeId,omitempty"`
	Id              *int                                   `json:"id,omitempty"`
	Preset          *Product_Package_Preset                `json:"preset,omitempty"`
	PresetId        *int                                   `json:"presetId,omitempty"`
	Value           *string                                `json:"value,omitempty"`
}

type Product_Package_Preset_Attribute_Type struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	KeyName     *string `json:"keyName,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Product_Package_Preset_Configuration struct {
	Entity

	Category      *Product_Item_Category  `json:"category,omitempty"`
	PackagePreset *Product_Package_Preset `json:"packagePreset,omitempty"`
	Price         *Product_Item_Price     `json:"price,omitempty"`
}

type Product_Package_Server struct {
	Entity

	Catalog                *Product_Catalog        `json:"catalog,omitempty"`
	CatalogId              *int                    `json:"catalogId,omitempty"`
	Datacenters            *string                 `json:"datacenters,omitempty"`
	DefaultRamCapacity     *float64                `json:"defaultRamCapacity,omitempty"`
	DualPathNetworkFlag    *bool                   `json:"dualPathNetworkFlag,omitempty"`
	GpuFlag                *bool                   `json:"gpuFlag,omitempty"`
	HourlyBillingFlag      *bool                   `json:"hourlyBillingFlag,omitempty"`
	Id                     *int                    `json:"id,omitempty"`
	Item                   *Product_Item           `json:"item,omitempty"`
	ItemId                 *int                    `json:"itemId,omitempty"`
	ItemPrice              *Product_Item_Price     `json:"itemPrice,omitempty"`
	ItemPriceId            *int                    `json:"itemPriceId,omitempty"`
	MaximumDriveCount      *int                    `json:"maximumDriveCount,omitempty"`
	MaximumPortSpeed       *float64                `json:"maximumPortSpeed,omitempty"`
	MaximumRamCapacity     *float64                `json:"maximumRamCapacity,omitempty"`
	MinimumPortSpeed       *float64                `json:"minimumPortSpeed,omitempty"`
	OutletFlag             *bool                   `json:"outletFlag,omitempty"`
	Package                *Product_Package        `json:"package,omitempty"`
	PackageId              *int                    `json:"packageId,omitempty"`
	PackageType            *string                 `json:"packageType,omitempty"`
	PowerServerFlag        *bool                   `json:"powerServerFlag,omitempty"`
	Preset                 *Product_Package_Preset `json:"preset,omitempty"`
	PresetId               *int                    `json:"presetId,omitempty"`
	PrivateNetworkOnlyFlag *bool                   `json:"privateNetworkOnlyFlag,omitempty"`
	ProcessorBusSpeed      *string                 `json:"processorBusSpeed,omitempty"`
	ProcessorCache         *string                 `json:"processorCache,omitempty"`
	ProcessorCores         *int                    `json:"processorCores,omitempty"`
	ProcessorCount         *int                    `json:"processorCount,omitempty"`
	ProcessorManufacturer  *string                 `json:"processorManufacturer,omitempty"`
	ProcessorModel         *string                 `json:"processorModel,omitempty"`
	ProcessorName          *string                 `json:"processorName,omitempty"`
	ProcessorSpeed         *string                 `json:"processorSpeed,omitempty"`
	ProductName            *string                 `json:"productName,omitempty"`
	RedundantPowerFlag     *bool                   `json:"redundantPowerFlag,omitempty"`
	SapCertifiedServerFlag *bool                   `json:"sapCertifiedServerFlag,omitempty"`
	StartingHourlyPrice    *float64                `json:"startingHourlyPrice,omitempty"`
	StartingMonthlyPrice   *float64                `json:"startingMonthlyPrice,omitempty"`
	TotalCoreCount         *int                    `json:"totalCoreCount,omitempty"`
	TxtTpmFlag             *bool                   `json:"txtTpmFlag,omitempty"`
	UnitSize               *int                    `json:"unitSize,omitempty"`
}

type Product_Package_Server_Option struct {
	Entity

	CatalogId   *int    `json:"catalogId,omitempty"`
	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	Type        *string `json:"type,omitempty"`
	Value       *string `json:"value,omitempty"`
}

type Product_Package_Type struct {
	Entity

	Id           *int              `json:"id,omitempty"`
	KeyName      *string           `json:"keyName,omitempty"`
	Name         *string           `json:"name,omitempty"`
	PackageCount *uint             `json:"packageCount,omitempty"`
	Packages     []Product_Package `json:"packages,omitempty"`
}

type Product_Upgrade_Request struct {
	Entity

	Account                 *Account                        `json:"account,omitempty"`
	AccountId               *int                            `json:"accountId,omitempty"`
	CompletedFlag           *bool                           `json:"completedFlag,omitempty"`
	CreateDate              *time.Time                      `json:"createDate,omitempty"`
	EmployeeId              *int                            `json:"employeeId,omitempty"`
	GuestId                 *int                            `json:"guestId,omitempty"`
	HardwareId              *int                            `json:"hardwareId,omitempty"`
	Id                      *int                            `json:"id,omitempty"`
	Invoice                 *Billing_Invoice                `json:"invoice,omitempty"`
	MaintenanceStartTimeUtc *time.Time                      `json:"maintenanceStartTimeUtc,omitempty"`
	ModifyDate              *time.Time                      `json:"modifyDate,omitempty"`
	Order                   *Billing_Order                  `json:"order,omitempty"`
	OrderId                 *int                            `json:"orderId,omitempty"`
	OrderTotal              *float64                        `json:"orderTotal,omitempty"`
	ProratedTotal           *float64                        `json:"proratedTotal,omitempty"`
	Server                  *Hardware                       `json:"server,omitempty"`
	Status                  *Product_Upgrade_Request_Status `json:"status,omitempty"`
	StatusId                *int                            `json:"statusId,omitempty"`
	Ticket                  *Ticket                         `json:"ticket,omitempty"`
	TicketId                *int                            `json:"ticketId,omitempty"`
	User                    *User_Customer                  `json:"user,omitempty"`
	UserId                  *int                            `json:"userId,omitempty"`
	VirtualGuest            *Virtual_Guest                  `json:"virtualGuest,omitempty"`
}

type Product_Upgrade_Request_Status struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	StatusCode  *string `json:"statusCode,omitempty"`
}
