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

package virtual

import (
	"fmt"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/helpers/product"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/sl"
	"sync"
	"time"
)

// Upgrade a virtual guest to a specified set of features (e.g. cpu, ram).
// When the upgrade takes place can also be specified (`when`), but
// this is optional. The time set will be 'now' if left as nil.
// The features to upgrade are specified as the options used in
// GetProductPrices().
func UpgradeVirtualGuest(
	sess *session.Session,
	guest *datatypes.Virtual_Guest,
	options map[string]float64,
	when ...time.Time,
) (datatypes.Container_Product_Order_Receipt, error) {

	// Get the packageID of the guest and use same during upgrade of guest.
	service := services.GetVirtualGuestService(sess)
	guestPackage, err := service.Id(*guest.Id).Mask("privateNetworkOnlyFlag,dedicatedAccountHostOnlyFlag,billingItem[package[id]]").GetObject()
	if err != nil {
		return datatypes.Container_Product_Order_Receipt{}, err
	}

	if guest.PrivateNetworkOnlyFlag == nil || guest.DedicatedAccountHostOnlyFlag == nil {
		guest.PrivateNetworkOnlyFlag = guestPackage.PrivateNetworkOnlyFlag
		guest.DedicatedAccountHostOnlyFlag = guestPackage.DedicatedAccountHostOnlyFlag
	}

	var packageID int
	if guestPackage.BillingItem != nil {
		packageID = *guestPackage.BillingItem.Package.Id
	}

	productItems, err := product.GetPackageProducts(sess, packageID)
	if err != nil {
		return datatypes.Container_Product_Order_Receipt{}, err
	}

	prices := product.SelectProductPricesByCategory(productItems, options, !*guest.PrivateNetworkOnlyFlag, !*guest.DedicatedAccountHostOnlyFlag)

	upgradeTime := time.Now().UTC().Format(time.RFC3339)
	if len(when) > 0 {
		upgradeTime = when[0].UTC().Format(time.RFC3339)
	}

	order := datatypes.Container_Product_Order_Virtual_Guest_Upgrade{
		Container_Product_Order_Virtual_Guest: datatypes.Container_Product_Order_Virtual_Guest{
			Container_Product_Order_Hardware_Server: datatypes.Container_Product_Order_Hardware_Server{
				Container_Product_Order: datatypes.Container_Product_Order{
					PackageId: &packageID,
					VirtualGuests: []datatypes.Virtual_Guest{
						*guest,
					},
					Prices: prices,
					Properties: []datatypes.Container_Product_Order_Property{
						{
							Name:  sl.String("MAINTENANCE_WINDOW"),
							Value: &upgradeTime,
						},
					},
				},
			},
		},
	}

	orderService := services.GetProductOrderService(sess)
	return orderService.PlaceOrder(&order, sl.Bool(false))
}

// Upgrade a virtual guest with preset to a specified set of features (e.g. flavor,disks).
// When the upgrade takes place can also be specified (`when`), but
// this is optional. The time set will be 'now' if left as nil.
// The features to upgrade are specified as the options used in
// GetProductPrices().
func UpgradeVirtualGuestWithPreset(
	sess *session.Session,
	guest *datatypes.Virtual_Guest,
	presetKeyName string,
	options map[string]float64,
	when ...time.Time,
) (datatypes.Container_Product_Order_Receipt, error) {

	// Get the packageID of the guest and use same during upgrade of guest.
	service := services.GetVirtualGuestService(sess)
	guestPackage, err := service.Id(*guest.Id).Mask("privateNetworkOnlyFlag,dedicatedAccountHostOnlyFlag,billingItem[package[id]]").GetObject()
	if err != nil {
		return datatypes.Container_Product_Order_Receipt{}, err
	}

	if guest.PrivateNetworkOnlyFlag == nil || guest.DedicatedAccountHostOnlyFlag == nil {
		guest.PrivateNetworkOnlyFlag = guestPackage.PrivateNetworkOnlyFlag
		guest.DedicatedAccountHostOnlyFlag = guestPackage.DedicatedAccountHostOnlyFlag
	}

	var packageID int
	if guestPackage.BillingItem != nil {
		packageID = *guestPackage.BillingItem.Package.Id
	}

	preset, _ := product.GetPresetByKeyName(sess, packageID, presetKeyName)

	productItems, err := product.GetPackageProducts(sess, packageID)
	if err != nil {
		return datatypes.Container_Product_Order_Receipt{}, err
	}

	prices := product.SelectProductPricesByCategory(productItems, options, !*guest.PrivateNetworkOnlyFlag, !*guest.DedicatedAccountHostOnlyFlag)

	upgradeTime := time.Now().UTC().Format(time.RFC3339)
	if len(when) > 0 {
		upgradeTime = when[0].UTC().Format(time.RFC3339)
	}

	order := datatypes.Container_Product_Order_Virtual_Guest_Upgrade{
		Container_Product_Order_Virtual_Guest: datatypes.Container_Product_Order_Virtual_Guest{
			Container_Product_Order_Hardware_Server: datatypes.Container_Product_Order_Hardware_Server{
				Container_Product_Order: datatypes.Container_Product_Order{
					PackageId: &packageID,
					VirtualGuests: []datatypes.Virtual_Guest{
						*guest,
					},
					PresetId: preset.Id,
					Prices:   prices,
					Properties: []datatypes.Container_Product_Order_Property{
						{
							Name:  sl.String("MAINTENANCE_WINDOW"),
							Value: &upgradeTime,
						},
					},
				},
			},
		},
	}

	orderService := services.GetProductOrderService(sess)
	return orderService.PlaceOrder(&order, sl.Bool(false))
}

// Use go-routines to iterate through all virtual guest results.
// optoins should be any Mask or Filter you need, and a Limit if the default is too large.
// Any error in the subsequent API calls will be logged, but largely ignored
func GetVirtualGuestsIter(session session.SLSession, options *sl.Options) (resp []datatypes.Virtual_Guest, err error) {

	options.SetOffset(0)
	limit := options.ValidateLimit()

	// Can't call service.GetVirtualGuests because it passes a copy of options, not the address to options sadly.
	err = session.DoRequest("SoftLayer_Account", "getVirtualGuests", nil, options, &resp)
	if err != nil {
		return
	}
	apicalls := options.GetRemainingAPICalls()
	var wg sync.WaitGroup
	for x := 1; x <= apicalls; x++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			offset := i * limit
			this_resp := []datatypes.Virtual_Guest{}
			options.Offset = &offset
			err = session.DoRequest("SoftLayer_Account", "getVirtualGuests", nil, options, &this_resp)
			if err != nil {
				fmt.Printf("[ERROR] %v\n", err)
			}
			resp = append(resp, this_resp...)
		}(x)
	}
	wg.Wait()
	return resp, err
}
