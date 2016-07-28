package softlayer

import (
	"time"

	"github.ibm.com/riethm/gopherlayer/datatypes"
)

type Product_Catalog interface {
	Entity

	GetBrands() (datatypes.Brand, error)
	GetPackages() (datatypes.Product_Package, error)
	GetPrices() (datatypes.Product_Item_Price, error)
	GetProducts() (datatypes.Product_Item, error)
}

type Product_Catalog_Item_Price interface {
	Entity

	GetCatalog() (datatypes.Product_Catalog, error)
	GetPrice() (datatypes.Product_Item_Price, error)
}

type Product_Item interface {
	Entity

	GetActivePresaleEvents() (datatypes.Sales_Presale_Event, error)
	GetActiveUsagePrices() (datatypes.Product_Item_Price, error)
	GetAttributes() (datatypes.Product_Item_Attribute, error)
	GetAvailabilityAttributes() (datatypes.Product_Item_Attribute, error)
	GetBillingType() (string, error)
	GetBundle() (datatypes.Product_Item_Bundles, error)
	GetCapacityRestrictedProductFlag() (bool, error)
	GetCategories() (datatypes.Product_Item_Category, error)
	GetConfigurationTemplates() (datatypes.Configuration_Template, error)
	GetConflicts() (datatypes.Product_Item_Resource_Conflict, error)
	GetCoreRestrictedItemFlag() (bool, error)
	GetDowngradeItem() (datatypes.Product_Item, error)
	GetDowngradeItems() (datatypes.Product_Item, error)
	GetGlobalCategoryConflicts() (datatypes.Product_Item_Resource_Conflict, error)
	GetHardwareGenericComponentModel() (datatypes.Hardware_Component_Model_Generic, error)
	GetHideFromPortalFlag() (bool, error)
	GetInventory() (datatypes.Product_Package_Inventory, error)
	GetIsEngineeredServerProduct() (bool, error)
	GetItemCategory() (datatypes.Product_Item_Category, error)
	GetLocationConflicts() (datatypes.Product_Item_Resource_Conflict, error)
	GetObjectStorageItemFlag() (bool, error)
	GetPackages() (datatypes.Product_Package, error)
	GetPhysicalCoreCapacity() (string, error)
	GetPresaleEvents() (datatypes.Sales_Presale_Event, error)
	GetPrices() (datatypes.Product_Item_Price, error)
	GetRequirements() (datatypes.Product_Item_Requirement, error)
	GetSoftwareDescription() (datatypes.Software_Description, error)
	GetTaxCategory() (datatypes.Product_Item_Tax_Category, error)
	GetThirdPartyPolicyAssignments() (datatypes.Product_Item_Policy_Assignment, error)
	GetThirdPartySupportVendor() (string, error)
	GetTotalPhysicalCoreCapacity() (int, error)
	GetTotalPhysicalCoreCount() (int, error)
	GetTotalProcessorCapacity() (int, error)
	GetUpgradeItem() (datatypes.Product_Item, error)
	GetUpgradeItems() (datatypes.Product_Item, error)
}

type Product_Item_Attribute interface {
	Entity

	GetAttributeType() (datatypes.Product_Item_Attribute_Type, error)
	GetAttributeTypeKeyName() (string, error)
	GetItem() (datatypes.Product_Item, error)
}

type Product_Item_Attribute_Type interface {
	Entity
}

type Product_Item_Billing_Type interface {
	Entity
}

type Product_Item_Bundles interface {
	Entity

	GetBundleItem() (datatypes.Product_Item, error)
	GetCategory() (datatypes.Product_Item_Category, error)
	GetItemPrice() (datatypes.Product_Item_Price, error)
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

	GetBillingItems() (datatypes.Billing_Item, error)
	GetGroup() (datatypes.Product_Item_Category_Group, error)
	GetGroups() (datatypes.Product_Package_Item_Category_Group, error)
	GetOrderOptions() (datatypes.Product_Item_Category_Order_Option_Type, error)
	GetPackageConfigurations() (datatypes.Product_Package_Order_Configuration, error)
	GetPresetConfigurations() (datatypes.Product_Package_Preset_Configuration, error)
	GetQuestionReferences() (datatypes.Product_Item_Category_Question_Xref, error)
	GetQuestions() (datatypes.Product_Item_Category_Question, error)
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

	GetFieldType() (datatypes.Product_Item_Category_Question_Field_Type, error)
	GetItemCategoryReferences() (datatypes.Product_Item_Category_Question_Xref, error)
}

type Product_Item_Category_Question_Field_Type interface {
	Entity
}

type Product_Item_Category_Question_Xref interface {
	Entity

	GetItemCategory() (datatypes.Product_Item_Category, error)
	GetQuestion() (datatypes.Product_Item_Category_Question, error)
}

type Product_Item_Link_ThePlanet interface {
	Entity

	GetItem() (datatypes.Product_Item, error)
	GetServiceProvider() (datatypes.Service_Provider, error)
}

type Product_Item_Policy_Assignment interface {
	Entity

	AcceptFromTicket(ticketId int) (bool, error)
	GetObject() (datatypes.Product_Item_Policy_Assignment, error)
	GetPolicyDocumentContents() ([]byte, error)

	GetPolicyName() (string, error)
	GetProduct() (datatypes.Product_Item, error)
}

type Product_Item_Price interface {
	Entity

	GetObject() (datatypes.Product_Item_Price, error)
	GetUsageRatePrices(location datatypes.Location, items datatypes.Product_Item) (datatypes.Product_Item_Price, error)

	GetAccountRestrictions() (datatypes.Product_Item_Price_Account_Restriction, error)
	GetAttributes() (datatypes.Product_Item_Price_Attribute, error)
	GetBigDataOsJournalDiskFlag() (bool, error)
	GetBundleReferences() (datatypes.Product_Item_Bundles, error)
	GetCapacityRestrictionMaximum() (string, error)
	GetCapacityRestrictionMinimum() (string, error)
	GetCapacityRestrictionType() (string, error)
	GetCategories() (datatypes.Product_Item_Category, error)
	GetDefinedSoftwareLicenseFlag() (bool, error)
	GetInventory() (datatypes.Product_Package_Inventory, error)
	GetItem() (datatypes.Product_Item, error)
	GetOrderPremiums() (datatypes.Product_Item_Price_Premium, error)
	GetPackageReferences() (datatypes.Product_Package_Item_Prices, error)
	GetPackages() (datatypes.Product_Package, error)
	GetPresetConfigurations() (datatypes.Product_Package_Preset_Configuration, error)
	GetPricingLocationGroup() (datatypes.Location_Group_Pricing, error)
	GetRequiredCoreCount() (int, error)
}

type Product_Item_Price_Account_Restriction interface {
	Entity

	GetAccount() (datatypes.Account, error)
	GetItemPrice() (datatypes.Product_Item_Price, error)
}

type Product_Item_Price_Attribute interface {
	Entity

	GetItemPrice() (datatypes.Product_Item_Price, error)
	GetItemPriceAttributeType() (datatypes.Product_Item_Price_Attribute_Type, error)
}

type Product_Item_Price_Attribute_Type interface {
	Entity
}

type Product_Item_Price_Premium interface {
	Entity

	GetObject() (datatypes.Product_Item_Price_Premium, error)

	GetItemPrice() (datatypes.Product_Item_Price, error)
	GetLocation() (datatypes.Location, error)
	GetPackage() (datatypes.Product_Package, error)
}

type Product_Item_Requirement interface {
	Entity

	GetItem() (datatypes.Product_Item, error)
	GetProduct() (datatypes.Product_Item, error)
}

type Product_Item_Resource_Conflict interface {
	Entity

	GetItem() (datatypes.Product_Item, error)
	GetPackage() (datatypes.Product_Package, error)
}

type Product_Item_Resource_Conflict_Item interface {
	Product_Item_Resource_Conflict

	GetResource() (datatypes.Product_Item, error)
}

type Product_Item_Resource_Conflict_Item_Category interface {
	Product_Item_Resource_Conflict

	GetResource() (datatypes.Product_Item_Category, error)
}

type Product_Item_Resource_Conflict_Location interface {
	Product_Item_Resource_Conflict

	GetResource() (datatypes.Location, error)
}

type Product_Item_Tax_Category interface {
	Entity

	GetItems() (datatypes.Product_Item, error)
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

	GetAccountRestrictedCategories() (datatypes.Product_Item_Category, error)
	GetAccountRestrictedPricesFlag() (bool, error)
	GetActivePresets() (datatypes.Product_Package_Preset, error)
	GetActiveRamItems() (datatypes.Product_Item, error)
	GetActiveServerItems() (datatypes.Product_Item, error)
	GetActiveSoftwareItems() (datatypes.Product_Item, error)
	GetActiveUsagePrices() (datatypes.Product_Item_Price, error)
	GetAdditionalServiceFlag() (bool, error)
	GetAttributes() (datatypes.Product_Package_Attribute, error)
	GetAvailableLocations() (datatypes.Product_Package_Locations, error)
	GetAvailableStorageUnits() (uint, error)
	GetCategories() (datatypes.Product_Item_Category, error)
	GetConfiguration() (datatypes.Product_Package_Order_Configuration, error)
	GetDefaultRamItems() (datatypes.Product_Item, error)
	GetDeploymentNodeType() (string, error)
	GetDeploymentPackages() (datatypes.Product_Package, error)
	GetDeploymentType() (string, error)
	GetDeployments() (datatypes.Product_Package, error)
	GetDisallowCustomDiskPartitions() (bool, error)
	GetFirstOrderStep() (datatypes.Product_Package_Order_Step, error)
	GetGatewayApplianceFlag() (bool, error)
	GetGpuFlag() (bool, error)
	GetHourlyBillingAvailableFlag() (bool, error)
	GetItemConflicts() (datatypes.Product_Item_Resource_Conflict, error)
	GetItemLocationConflicts() (datatypes.Product_Item_Resource_Conflict, error)
	GetItemPriceReferences() (datatypes.Product_Package_Item_Prices, error)
	GetItemPrices() (datatypes.Product_Item_Price, error)
	GetItems() (datatypes.Product_Item, error)
	GetLocations() (datatypes.Location, error)
	GetLowestServerPrice() (datatypes.Product_Item_Price, error)
	GetMaximumPortSpeed() (uint, error)
	GetMinimumPortSpeed() (uint, error)
	GetMongoDbEngineeredFlag() (bool, error)
	GetOrderPremiums() (datatypes.Product_Item_Price_Premium, error)
	GetPreconfiguredFlag() (bool, error)
	GetPresetConfigurationRequiredFlag() (bool, error)
	GetPreventVlanSelectionFlag() (bool, error)
	GetPrivateHostedCloudPackageFlag() (bool, error)
	GetPrivateHostedCloudPackageType() (string, error)
	GetPrivateNetworkOnlyFlag() (bool, error)
	GetQuantaStorPackageFlag() (bool, error)
	GetRaidDiskRestrictionFlag() (bool, error)
	GetRedundantPowerFlag() (bool, error)
	GetRegions() (datatypes.Location_Region, error)
	GetResourceGroupTemplate() (datatypes.Resource_Group_Template, error)
	GetTopLevelItemCategoryCode() (string, error)
	GetType() (datatypes.Product_Package_Type, error)
}

type Product_Package_Attribute interface {
	Entity

	GetAttributeType() (datatypes.Product_Package_Attribute_Type, error)
	GetPackage() (datatypes.Product_Package, error)
}

type Product_Package_Attribute_Type interface {
	Entity
}

type Product_Package_Inventory interface {
	Entity

	GetItem() (datatypes.Product_Item, error)
	GetLocation() (datatypes.Location, error)
	GetPackage() (datatypes.Product_Package, error)
}

type Product_Package_Item_Category_Group interface {
	Entity

	GetCategory() (datatypes.Product_Item_Category, error)
	GetPackage() (datatypes.Product_Package, error)
	GetPrices() (datatypes.Product_Item_Price, error)
}

type Product_Package_Item_Prices interface {
	Entity

	GetItemPrice() (datatypes.Product_Item_Price, error)
	GetPackage() (datatypes.Product_Package, error)
}

type Product_Package_Items interface {
	Entity

	GetItem() (datatypes.Product_Item, error)
	GetPackage() (datatypes.Product_Package, error)
}

type Product_Package_Locations interface {
	Entity

	GetLocation() (datatypes.Location, error)
	GetPackage() (datatypes.Product_Package, error)
}

type Product_Package_Order_Configuration interface {
	Entity

	GetItemCategory() (datatypes.Product_Item_Category, error)
	GetPackage() (datatypes.Product_Package, error)
	GetStep() (datatypes.Product_Package_Order_Step, error)
}

type Product_Package_Order_Step interface {
	Entity

	GetInclusivePreviousSteps() (datatypes.Product_Package_Order_Step_Next, error)
	GetNextSteps() (datatypes.Product_Package_Order_Step_Next, error)
	GetPreviousSteps() (datatypes.Product_Package_Order_Step_Next, error)
}

type Product_Package_Order_Step_Next interface {
	Entity

	GetStep() (datatypes.Product_Package_Order_Step, error)
}

type Product_Package_Preset interface {
	Entity

	GetAllObjects() (datatypes.Product_Package_Preset, error)
	GetObject() (datatypes.Product_Package_Preset, error)

	GetAvailableStorageUnits() (uint, error)
	GetCategories() (datatypes.Product_Item_Category, error)
	GetConfiguration() (datatypes.Product_Package_Preset_Configuration, error)
	GetFixedConfigurationFlag() (bool, error)
	GetLowestPresetServerPrice() (datatypes.Product_Item_Price, error)
	GetPackage() (datatypes.Product_Package, error)
	GetPackageConfiguration() (datatypes.Product_Package_Order_Configuration, error)
	GetPrices() (datatypes.Product_Item_Price, error)
	GetStorageGroupTemplateArrays() (datatypes.Configuration_Storage_Group_Template_Group, error)
	GetTotalMinimumHourlyFee() (float64, error)
	GetTotalMinimumRecurringFee() (float64, error)
}

type Product_Package_Preset_Attribute interface {
	Entity

	GetAttributeType() (datatypes.Product_Package_Preset_Attribute_Type, error)
	GetPreset() (datatypes.Product_Package_Preset, error)
}

type Product_Package_Preset_Attribute_Type interface {
	Entity
}

type Product_Package_Preset_Configuration interface {
	Entity

	GetCategory() (datatypes.Product_Item_Category, error)
	GetPackagePreset() (datatypes.Product_Package_Preset, error)
	GetPrice() (datatypes.Product_Item_Price, error)
}

type Product_Package_Server interface {
	Entity

	GetAllObjects() (datatypes.Product_Package_Server, error)
	GetObject() (datatypes.Product_Package_Server, error)

	GetCatalog() (datatypes.Product_Catalog, error)
	GetItem() (datatypes.Product_Item, error)
	GetItemPrice() (datatypes.Product_Item_Price, error)
	GetPackage() (datatypes.Product_Package, error)
	GetPreset() (datatypes.Product_Package_Preset, error)
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

	GetPackages() (datatypes.Product_Package, error)
}

type Product_Upgrade_Request interface {
	Entity

	ApproveChanges() (bool, error)
	GetObject() (datatypes.Product_Upgrade_Request, error)
	UpdateMaintenanceWindow(maintenanceStartTime time.Time, maintenanceWindowId int) (bool, error)

	GetAccount() (datatypes.Account, error)
	GetCompletedFlag() (bool, error)
	GetInvoice() (datatypes.Billing_Invoice, error)
	GetOrder() (datatypes.Billing_Order, error)
	GetServer() (datatypes.Hardware, error)
	GetStatus() (datatypes.Product_Upgrade_Request_Status, error)
	GetTicket() (datatypes.Ticket, error)
	GetUser() (datatypes.User_Customer, error)
	GetVirtualGuest() (datatypes.Virtual_Guest, error)
}

type Product_Upgrade_Request_Status interface {
	Entity
}
