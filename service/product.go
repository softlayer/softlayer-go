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

type Product_Item_Category struct {
	Session *Session
	Options
}

func (r *Session) GetProductItemCategoryService() Product_Item_Category {
	return Product_Item_Category{Session: r}
}

func (r *Product_Item_Category) GetAdditionalProductsForCategory() (resp []datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Category) GetBandwidthCategories() (resp []datatypes.Product_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Category) GetBillingItems() (resp []datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Category) GetComputingCategories(resetCache *bool) (resp []datatypes.Product_Item_Category, err error) {
	params := []interface{}{
		resetCache,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Category) GetCustomUsageRatesCategories() (resp []datatypes.Product_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Category) GetGroup() (resp datatypes.Product_Item_Category_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Category) GetGroups() (resp []datatypes.Product_Package_Item_Category_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Category) GetObject() (resp datatypes.Product_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Category) GetOrderOptions() (resp []datatypes.Product_Item_Category_Order_Option_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Category) GetPackageConfigurations() (resp []datatypes.Product_Package_Order_Configuration, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Category) GetPresetConfigurations() (resp []datatypes.Product_Package_Preset_Configuration, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Category) GetQuestionReferences() (resp []datatypes.Product_Item_Category_Question_Xref, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Category) GetQuestions() (resp []datatypes.Product_Item_Category_Question, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Category) GetSoftwareCategories() (resp []datatypes.Product_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Category) GetSubnetCategories() (resp []datatypes.Product_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Category) GetTopLevelCategories(resetCache *bool) (resp []datatypes.Product_Item_Category, err error) {
	params := []interface{}{
		resetCache,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Category) GetValidCancelableServiceItemCategories() (resp []datatypes.Product_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Category) GetVlanCategories() (resp []datatypes.Product_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Product_Item_Category_Group struct {
	Session *Session
	Options
}

func (r *Session) GetProductItemCategoryGroupService() Product_Item_Category_Group {
	return Product_Item_Category_Group{Session: r}
}

func (r *Product_Item_Category_Group) GetObject() (resp datatypes.Product_Item_Category_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Product_Item_Policy_Assignment struct {
	Session *Session
	Options
}

func (r *Session) GetProductItemPolicyAssignmentService() Product_Item_Policy_Assignment {
	return Product_Item_Policy_Assignment{Session: r}
}

func (r *Product_Item_Policy_Assignment) AcceptFromTicket(ticketId *int) (resp bool, err error) {
	params := []interface{}{
		ticketId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Policy_Assignment) GetObject() (resp datatypes.Product_Item_Policy_Assignment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Policy_Assignment) GetPolicyDocumentContents() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Policy_Assignment) GetPolicyName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Policy_Assignment) GetProduct() (resp datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Product_Item_Price struct {
	Session *Session
	Options
}

func (r *Session) GetProductItemPriceService() Product_Item_Price {
	return Product_Item_Price{Session: r}
}

func (r *Product_Item_Price) GetAccountRestrictions() (resp []datatypes.Product_Item_Price_Account_Restriction, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Price) GetAttributes() (resp []datatypes.Product_Item_Price_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Price) GetBigDataOsJournalDiskFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Price) GetBundleReferences() (resp []datatypes.Product_Item_Bundles, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Price) GetCapacityRestrictionMaximum() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Price) GetCapacityRestrictionMinimum() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Price) GetCapacityRestrictionType() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Price) GetCategories() (resp []datatypes.Product_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Price) GetDefinedSoftwareLicenseFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Price) GetInventory() (resp []datatypes.Product_Package_Inventory, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Price) GetItem() (resp datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Price) GetObject() (resp datatypes.Product_Item_Price, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Price) GetOrderPremiums() (resp []datatypes.Product_Item_Price_Premium, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Price) GetPackageReferences() (resp []datatypes.Product_Package_Item_Prices, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Price) GetPackages() (resp []datatypes.Product_Package, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Price) GetPresetConfigurations() (resp []datatypes.Product_Package_Preset_Configuration, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Price) GetPricingLocationGroup() (resp datatypes.Location_Group_Pricing, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Price) GetRequiredCoreCount() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Price) GetUsageRatePrices(location *datatypes.Location, items []datatypes.Product_Item) (resp []datatypes.Product_Item_Price, err error) {
	params := []interface{}{
		location,
		items,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Product_Item_Price_Premium struct {
	Session *Session
	Options
}

func (r *Session) GetProductItemPricePremiumService() Product_Item_Price_Premium {
	return Product_Item_Price_Premium{Session: r}
}

func (r *Product_Item_Price_Premium) GetItemPrice() (resp datatypes.Product_Item_Price, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Price_Premium) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Price_Premium) GetObject() (resp datatypes.Product_Item_Price_Premium, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Item_Price_Premium) GetPackage() (resp datatypes.Product_Package, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Product_Order struct {
	Session *Session
	Options
}

func (r *Session) GetProductOrderService() Product_Order {
	return Product_Order{Session: r}
}

func (r *Product_Order) CheckItemAvailability(itemPrices []datatypes.Product_Item_Price, accountId *int, availabilityTypeKeyNames []string) (resp bool, err error) {
	params := []interface{}{
		itemPrices,
		accountId,
		availabilityTypeKeyNames,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Order) CheckItemAvailabilityForImageTemplate(imageTemplateId *int, accountId *int, packageId *int, availabilityTypeKeyNames []string) (resp bool, err error) {
	params := []interface{}{
		imageTemplateId,
		accountId,
		packageId,
		availabilityTypeKeyNames,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Order) CheckItemConflicts(itemPrices []datatypes.Product_Item_Price) (resp bool, err error) {
	params := []interface{}{
		itemPrices,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Order) GetExternalPaymentAuthorizationReceipt(token *string, payerId *string) (resp datatypes.Container_Product_Order_Receipt, err error) {
	params := []interface{}{
		token,
		payerId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Order) GetNetworks(locationId *int, packageId *int, accountId *int) (resp []datatypes.Container_Product_Order_Network, err error) {
	params := []interface{}{
		locationId,
		packageId,
		accountId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Order) GetResellerOrder(orderContainer *datatypes.Container_Product_Order) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		orderContainer,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Order) GetTaxCalculationResult(orderHash *string) (resp datatypes.Container_Tax_Cache, err error) {
	params := []interface{}{
		orderHash,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Order) GetVlans(locationId *int, packageId *int, selectedItems *string, vlanIds []int, subnetIds []int, accountId *int, orderContainer *datatypes.Container_Product_Order) (resp datatypes.Container_Product_Order_Network_Vlans, err error) {
	params := []interface{}{
		locationId,
		packageId,
		selectedItems,
		vlanIds,
		subnetIds,
		accountId,
		orderContainer,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Order) PlaceOrder(orderData *datatypes.Container_Product_Order, saveAsQuote *bool) (resp datatypes.Container_Product_Order_Receipt, err error) {
	params := []interface{}{
		orderData,
		saveAsQuote,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Order) PlaceQuote(orderData *datatypes.Container_Product_Order) (resp datatypes.Container_Product_Order_Receipt, err error) {
	params := []interface{}{
		orderData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Order) ProcessExternalPaymentAuthorization(token *string, payerId *string) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		token,
		payerId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Order) RequiredItems(itemPrices []datatypes.Product_Item_Price) (resp []datatypes.Product_Item, err error) {
	params := []interface{}{
		itemPrices,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Order) VerifyOrder(orderData *datatypes.Container_Product_Order) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		orderData,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Product_Package struct {
	Session *Session
	Options
}

func (r *Session) GetProductPackageService() Product_Package {
	return Product_Package{Session: r}
}

func (r *Product_Package) GetAccountRestrictedCategories() (resp []datatypes.Product_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetAccountRestrictedPricesFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetActiveItems() (resp []datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetActivePackagesByAttribute(attributeKeyName *string) (resp []datatypes.Product_Package, err error) {
	params := []interface{}{
		attributeKeyName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetActivePresets() (resp []datatypes.Product_Package_Preset, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetActivePrivateHostedCloudPackages() (resp []datatypes.Product_Package, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetActiveRamItems() (resp []datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetActiveServerItems() (resp []datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetActiveSoftwareItems() (resp []datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetActiveUsagePrices() (resp []datatypes.Product_Item_Price, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetActiveUsageRatePrices(locationId *int, categoryCode *string) (resp []datatypes.Product_Item_Price, err error) {
	params := []interface{}{
		locationId,
		categoryCode,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetAdditionalServiceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetAllObjects() (resp []datatypes.Product_Package, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetAttributes() (resp []datatypes.Product_Package_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetAvailableLocations() (resp []datatypes.Product_Package_Locations, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetAvailablePackagesForImageTemplate(imageTemplate *datatypes.Virtual_Guest_Block_Device_Template_Group) (resp []datatypes.Product_Package, err error) {
	params := []interface{}{
		imageTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetAvailableStorageUnits() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetCategories() (resp []datatypes.Product_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetCdnItems() (resp []datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetCloudStorageItems(provider *int) (resp []datatypes.Product_Item, err error) {
	params := []interface{}{
		provider,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetConfiguration() (resp []datatypes.Product_Package_Order_Configuration, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetDefaultRamItems() (resp []datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetDeploymentNodeType() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetDeploymentPackages() (resp []datatypes.Product_Package, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetDeploymentType() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetDeployments() (resp []datatypes.Product_Package, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetDisallowCustomDiskPartitions() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetFirstOrderStep() (resp datatypes.Product_Package_Order_Step, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetGatewayApplianceFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetGpuFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetHourlyBillingAvailableFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetItemAvailabilityTypes() (resp []datatypes.Product_Item_Attribute_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetItemConflicts() (resp []datatypes.Product_Item_Resource_Conflict, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetItemLocationConflicts() (resp []datatypes.Product_Item_Resource_Conflict, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetItemPriceReferences() (resp []datatypes.Product_Package_Item_Prices, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetItemPrices() (resp []datatypes.Product_Item_Price, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetItemPricesFromSoftwareDescriptions(softwareDescriptions []datatypes.Software_Description, includeTranslationsFlag *bool, returnAllPricesFlag *bool) (resp []datatypes.Product_Item_Price, err error) {
	params := []interface{}{
		softwareDescriptions,
		includeTranslationsFlag,
		returnAllPricesFlag,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetItems() (resp []datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetItemsFromImageTemplate(imageTemplate *datatypes.Virtual_Guest_Block_Device_Template_Group) (resp []datatypes.Product_Item, err error) {
	params := []interface{}{
		imageTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetLocations() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetLowestServerPrice() (resp datatypes.Product_Item_Price, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetMaximumPortSpeed() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetMessageQueueItems() (resp []datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetMinimumPortSpeed() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetMongoDbEngineeredFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetObject() (resp datatypes.Product_Package, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetObjectStorageDatacenters() (resp []datatypes.Container_Product_Order_Network_Storage_Hub_Datacenter, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetOrderPremiums() (resp []datatypes.Product_Item_Price_Premium, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetPreconfiguredFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetPresetConfigurationRequiredFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetPreventVlanSelectionFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetPrivateHostedCloudPackageFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetPrivateHostedCloudPackageType() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetPrivateNetworkOnlyFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetQuantaStorPackageFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetRaidDiskRestrictionFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetRedundantPowerFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetRegions() (resp []datatypes.Location_Region, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetResourceGroupTemplate() (resp datatypes.Resource_Group_Template, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetStandardCategories() (resp []datatypes.Product_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetTopLevelItemCategoryCode() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package) GetType() (resp datatypes.Product_Package_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Product_Package_Preset struct {
	Session *Session
	Options
}

func (r *Session) GetProductPackagePresetService() Product_Package_Preset {
	return Product_Package_Preset{Session: r}
}

func (r *Product_Package_Preset) GetAllObjects() (resp []datatypes.Product_Package_Preset, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package_Preset) GetAvailableStorageUnits() (resp uint, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package_Preset) GetCategories() (resp []datatypes.Product_Item_Category, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package_Preset) GetConfiguration() (resp []datatypes.Product_Package_Preset_Configuration, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package_Preset) GetFixedConfigurationFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package_Preset) GetLowestPresetServerPrice() (resp datatypes.Product_Item_Price, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package_Preset) GetObject() (resp datatypes.Product_Package_Preset, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package_Preset) GetPackage() (resp datatypes.Product_Package, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package_Preset) GetPackageConfiguration() (resp []datatypes.Product_Package_Order_Configuration, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package_Preset) GetPrices() (resp []datatypes.Product_Item_Price, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package_Preset) GetStorageGroupTemplateArrays() (resp []datatypes.Configuration_Storage_Group_Template_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package_Preset) GetTotalMinimumHourlyFee() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package_Preset) GetTotalMinimumRecurringFee() (resp float64, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Product_Package_Server struct {
	Session *Session
	Options
}

func (r *Session) GetProductPackageServerService() Product_Package_Server {
	return Product_Package_Server{Session: r}
}

func (r *Product_Package_Server) GetAllObjects() (resp []datatypes.Product_Package_Server, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package_Server) GetCatalog() (resp datatypes.Product_Catalog, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package_Server) GetItem() (resp datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package_Server) GetItemPrice() (resp datatypes.Product_Item_Price, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package_Server) GetObject() (resp datatypes.Product_Package_Server, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package_Server) GetPackage() (resp datatypes.Product_Package, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package_Server) GetPreset() (resp datatypes.Product_Package_Preset, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Product_Package_Server_Option struct {
	Session *Session
	Options
}

func (r *Session) GetProductPackageServerOptionService() Product_Package_Server_Option {
	return Product_Package_Server_Option{Session: r}
}

func (r *Product_Package_Server_Option) GetAllOptions() (resp []datatypes.Product_Package_Server_Option, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package_Server_Option) GetObject() (resp datatypes.Product_Package_Server_Option, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package_Server_Option) GetOptions(typ *string) (resp []datatypes.Product_Package_Server_Option, err error) {
	params := []interface{}{
		typ,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

type Product_Package_Type struct {
	Session *Session
	Options
}

func (r *Session) GetProductPackageTypeService() Product_Package_Type {
	return Product_Package_Type{Session: r}
}

func (r *Product_Package_Type) GetAllObjects() (resp []datatypes.Product_Package_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package_Type) GetObject() (resp datatypes.Product_Package_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Package_Type) GetPackages() (resp []datatypes.Product_Package, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Product_Upgrade_Request struct {
	Session *Session
	Options
}

func (r *Session) GetProductUpgradeRequestService() Product_Upgrade_Request {
	return Product_Upgrade_Request{Session: r}
}

func (r *Product_Upgrade_Request) ApproveChanges() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Upgrade_Request) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Upgrade_Request) GetCompletedFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Upgrade_Request) GetInvoice() (resp datatypes.Billing_Invoice, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Upgrade_Request) GetObject() (resp datatypes.Product_Upgrade_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Upgrade_Request) GetOrder() (resp datatypes.Billing_Order, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Upgrade_Request) GetServer() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Upgrade_Request) GetStatus() (resp datatypes.Product_Upgrade_Request_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Upgrade_Request) GetTicket() (resp datatypes.Ticket, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Upgrade_Request) GetUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Upgrade_Request) GetVirtualGuest() (resp datatypes.Virtual_Guest, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Product_Upgrade_Request) UpdateMaintenanceWindow(maintenanceStartTime *datatypes.Time, maintenanceWindowId *int) (resp bool, err error) {
	params := []interface{}{
		maintenanceStartTime,
		maintenanceWindowId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
