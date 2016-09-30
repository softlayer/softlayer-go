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

// CheckOrderItemProvisionTransactionStatus returns true if the status of the
// provision transaction for the provided order item is in the list of
// provided statuses. Returns false otherwise.
//
// As a side-effect, the original passed order item is updated with the latest
// billing item.
//
// Returns an error if the order item has no billing item or the billing item
// has no provision transaction (this probably just means that the billing
// item has not been created yet)
func CheckOrderItemProvisionStatus(sess *session.Session, orderItem *datatypes.Billing_Order_Item, statuses []string) (bool, error) {
	service := services.GetBillingOrderItemService(sess)

	billingItem, err := service.
		Id(*orderItem.Id).
		Mask("mask[id,provisionTransaction[id,transactionStatus[name]]]").
		GetBillingItem()

	if err != nil {
		return false, err
	}

	if billingItem.Id == nil {
		return false, fmt.Errorf("The billing item is nil - unable to proceed with transaction check")
	}

	// Set the updated billing item in the order item, for the caller's convenience
	orderItem.BillingItem = &billingItem

	if billingItem.ProvisionTransaction == nil {
		return false, fmt.Errorf("Order item %d has no associated provision transaction", *orderItem.Id)
	}

	for _, status := range statuses {
		if *billingItem.ProvisionTransaction.TransactionStatus.Name == status {
			return true, nil
		}
	}

	return false, nil
}

// CheckOrderItemProvisionTransactionStatus returns true if the status of the
// provision transaction for the provided order item is "COMPLETE". Returns
// false otherwise.
//
// As a side-effect, the original passed order item is updated with the latest
// billing item.
//
// Returns an error if the order item has no billing item or the billing item
// has no provision transaction (this probably just means that the billing
// item has not been created yet)
func CheckOrderItemProvisionComplete(sess *session.Session, orderItem *datatypes.Billing_Order_Item) (bool, error) {
	return CheckOrderItemProvisionStatus(sess, orderItem, []string{"COMPLETE"})
}
