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

package order

import (
	"fmt"

	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/session"
)

func getOrderItemIds(orderItems []datatypes.Billing_Order_Item) []interface{} {
	result := make([]interface{}, len(orderItems))
	for _, item := range orderItems {
		result = append(result, *item.Id)
	}
	return result
}

// CheckItemProvisionStatus returns true if the status of the provision
// transaction for ALL of the provided order items is in the list of provided
// statuses. Returns false otherwise.  The order items provided must be part of
// the same order.
//
// As a side-effect, the original passed order items are updated with the
// latest billing item details, in an effort to avoid the need for further
// trips to the API to fetch this info.
//
// Returns an error if an order item has no billing item or the billing item
// has no provision transaction (this probably just means that the billing
// item has not been created yet). Also returns an error if all order items are
// not part of the same overall order. In such cases, the function returns
// immediately, without checking remaining order items.
func CheckItemProvisionStatus(sess *session.Session, orderItems []datatypes.Billing_Order_Item, statuses []string) (bool, error) {
	if len(orderItems) == 0 {
		return false, fmt.Errorf("No order items provided")
	}

	order, err := services.GetBillingOrderItemService(sess).
		Id(*orderItems[0].Id).
		Mask("mask[id,items[billingItem[provisionTransaction[id,transactionStatus[name]]]]]").
		GetOrder()

	if err != nil {
		return false, fmt.Errorf("Unable to retrieve order data from SoftLayer: %s", err)
	}

	// Create a map of retrieved order item IDs to order items for lookup
	idMap := map[int]datatypes.Billing_Order_Item{}
	for _, item := range order.Items {
		idMap[*item.Id] = item
	}

	// Check updated transaction status for each provided order item in turn.
	for ordersIndex, orderItem := range orderItems {
		if item, ok := idMap[*orderItem.Id]; !ok {
			return false, fmt.Errorf(
				"No order item could be found with ID %d", *orderItem.Id)
		} else {
			// No billing item usually means the order has not yet
			// completed, or else the ordered item was cancelled.
			if item.BillingItem == nil {
				return false, fmt.Errorf(
					"The billing item for order item %d is nil - unable to proceed with transaction check",
					*item.Id)
			}

			// Set the updated billing item in the order item,
			// for the caller's convenience
			orderItems[ordersIndex].BillingItem = item.BillingItem

			// Not all billing items have provision transactions.
			// This may vary based on the type of order being
			// placed. It is up to the caller to pass the correct
			// order items to be checked.
			//
			// Therefore, no provision transaction means that the
			// item has not started provisioning.
			if item.BillingItem.ProvisionTransaction == nil {
				return false, fmt.Errorf(
					"Order item %d has no associated provision transaction", *orderItem.Id)
			}

			var itemComplete = false
			for _, status := range statuses {
				if *item.BillingItem.ProvisionTransaction.TransactionStatus.Name == status {
					itemComplete = true
					break
				}
			}

			if !itemComplete {
				return false, nil
			}
		}
	}

	return true, nil
}

// CheckItemProvisionStatus returns true if the status of the provision
// transaction for ALL of the provided order items is "COMPLETE". Returns false
// otherwise.  The order items provided must be part of the same order.
//
// See order.CheckItemProvisionStatus for a complete description of side-
// effects and error conditions.
func CheckItemProvisionComplete(sess *session.Session, orderItems []datatypes.Billing_Order_Item) (bool, error) {
	return CheckItemProvisionStatus(sess, orderItems, []string{"COMPLETE"})
}
