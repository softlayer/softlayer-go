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

package product

import (
	"fmt"

	"github.ibm.com/riethm/gopherlayer.git/datatypes"
	"github.ibm.com/riethm/gopherlayer.git/filter"
	"github.ibm.com/riethm/gopherlayer.git/services"
	"github.ibm.com/riethm/gopherlayer.git/session"
	"time"
)

// Get a list of product item prices for a specific product package type and
// a specific set of price category id / product item capacity combinations.
// These combinations are passed as a map of integers (category id) mapped
// to strings (capacity)
func GetProductPrices(
	sess session.Session,
	packageType string,
	options map[int]string,
) ([]datatypes.Product_Item_Price, error) {

	service := services.GetProductPackageService(sess)

	// 1. Get package id
	packages, err := service.
		Mask("id,name,description,isActive,type[keyName]").
		Filter(
			filter.Build(
				filter.Path("type.keyName").Eq(packageType),
				filter.Path("description").NotContains("OUTLET"),
				filter.Path("name").NotContains("OUTLET"),
			),
		).
		Limit(1).
		GetAllObjects()
	if err != nil {
		return nil, err
	}

	if len(packages) == 0 {
		return nil, fmt.Errorf("No product packages found for %s", packageType)
	}

	// 2. Get product items for package id
	productItems, err := service.
		Id(*packages[0].Id).
		Mask("id,capacity,description,units,keyName,prices[id,categories[id,name]]").
		GetItems()
	if err != nil {
		return nil, err
	}

	// 3. Filter product items based on sets of category ids and capacity numbers
	prices := []datatypes.Product_Item_Price{}
	for _, productItem := range productItems {
		for _, category := range productItem.Prices[0].Categories {
			for k, v := range options {
				if productItem.Capacity != nil && *productItem.Capacity == v && *category.Id == k {
					prices = append(prices, productItem.Prices[0])
				}
			}
		}
	}

	return prices, nil
}

// Upgrade a virtual guest specified with an id, and a set of features to
// upgrade. When the upgrade takes place can also be specified (`when`), but
// this is optional. The time set will be 'now' if left as nil.
// The features to upgrade are specified as the options used in
// GetProductPrices().
func UpgradeVirtualGuest(
	sess session.Session,
	id int,
	options map[int]string,
	when time.Time,
) (datatypes.Container_Product_Order_Receipt, error) {

	prices, err := GetProductPrices(sess, "VIRTUAL_SERVER_INSTANCE", options)
	if err != nil {
		return nil, err
	}

	if when == nil {
		when = time.Now().UTC().Format(time.RFC3339)
	}

	order := datatypes.Container_Product_Order{
		VirtualGuests: []datatypes.Virtual_Guest{
			{Id: id},
		},
		Prices: prices,
		Properties: []datatypes.Container_Product_Order_Property{
			{
				Name:  "MAINTENANCE_WINDOW",
				Value: when,
			},
		},
	}

	orderService := services.GetProductOrderService(sess)
	return orderService.PlaceOrder(order, false)
}
