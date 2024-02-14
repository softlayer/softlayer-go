/**
 * Copyright 2016-2024 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
 * the License. You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed
 * on an "AS IS" BASIS,WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and limitations under the License.
 */

// AUTOMATICALLY GENERATED CODE - DO NOT MODIFY

package services

import (
	"fmt"
	"strings"

	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/sl"
)

// The SoftLayer_Billing_Invoice data type contains general information relating to an individual invoice applied to a SoftLayer customer account. Personal information in this type such as names, addresses, and phone numbers are taken from the account's contact information at the time the invoice is generated.
type Billing_Invoice struct {
	Session session.SLSession
	Options sl.Options
}

// GetBillingInvoiceService returns an instance of the Billing_Invoice SoftLayer service
func GetBillingInvoiceService(sess session.SLSession) Billing_Invoice {
	return Billing_Invoice{Session: sess}
}

func (r Billing_Invoice) Id(id int) Billing_Invoice {
	r.Options.Id = &id
	return r
}

func (r Billing_Invoice) Mask(mask string) Billing_Invoice {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Billing_Invoice) Filter(filter string) Billing_Invoice {
	r.Options.Filter = filter
	return r
}

func (r Billing_Invoice) Limit(limit int) Billing_Invoice {
	r.Options.Limit = &limit
	return r
}

func (r Billing_Invoice) Offset(offset int) Billing_Invoice {
	r.Options.Offset = &offset
	return r
}

// Retrieve A list of top-level invoice items that are on the currently pending invoice.
func (r Billing_Invoice) GetInvoiceTopLevelItems() (resp []datatypes.Billing_Invoice_Item, err error) {
	err = r.Session.DoRequest("SoftLayer_Billing_Invoice", "getInvoiceTopLevelItems", nil, &r.Options, &resp)
	return
}

// Every individual item that a SoftLayer customer is billed for is recorded in the SoftLayer_Billing_Item data type. Billing items range from server chassis to hard drives to control panels, bandwidth quota upgrades and port upgrade charges. Softlayer [[SoftLayer_Billing_Invoice|invoices]] are generated from the cost of a customer's billing items. Billing items are copied from the product catalog as they're ordered by customers to create a reference between an account and the billable items they own.
//
// Billing items exist in a tree relationship. Items are associated with each other by parent/child relationships. Component items such as CPU's, RAM, and software each have a parent billing item for the server chassis they're associated with. Billing Items with a null parent item do not have an associated parent item.
type Billing_Item struct {
	Session session.SLSession
	Options sl.Options
}

// GetBillingItemService returns an instance of the Billing_Item SoftLayer service
func GetBillingItemService(sess session.SLSession) Billing_Item {
	return Billing_Item{Session: sess}
}

func (r Billing_Item) Id(id int) Billing_Item {
	r.Options.Id = &id
	return r
}

func (r Billing_Item) Mask(mask string) Billing_Item {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Billing_Item) Filter(filter string) Billing_Item {
	r.Options.Filter = filter
	return r
}

func (r Billing_Item) Limit(limit int) Billing_Item {
	r.Options.Limit = &limit
	return r
}

func (r Billing_Item) Offset(offset int) Billing_Item {
	r.Options.Offset = &offset
	return r
}

// Cancel the resource or service for a billing Item. By default the billing item will be canceled on the next bill date and reclaim of the resource will begin shortly after the cancellation. Setting the "cancelImmediately" property to true will start the cancellation immediately if the item is eligible to be canceled immediately.
//
// The reason parameter could be from the list below:
// * "No longer needed"
// * "Business closing down"
// * "Server / Upgrade Costs"
// * "Migrating to larger server"
// * "Migrating to smaller server"
// * "Migrating to a different SoftLayer datacenter"
// * "Network performance / latency"
// * "Support response / timing"
// * "Sales process / upgrades"
// * "Moving to competitor"
func (r Billing_Item) CancelItem(cancelImmediately *bool, cancelAssociatedBillingItems *bool, reason *string, customerNote *string) (resp bool, err error) {
	params := []interface{}{
		cancelImmediately,
		cancelAssociatedBillingItems,
		reason,
		customerNote,
	}
	err = r.Session.DoRequest("SoftLayer_Billing_Item", "cancelItem", params, &r.Options, &resp)
	return
}

// Cancel the resource or service (excluding bare metal servers) for a billing Item. The billing item will be cancelled immediately and reclaim of the resource will begin shortly.
func (r Billing_Item) CancelService() (resp bool, err error) {
	err = r.Session.DoRequest("SoftLayer_Billing_Item", "cancelService", nil, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Billing_Item object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Billing_Item service. You can only retrieve billing items tied to the account that your portal user is assigned to. Billing items are an account's items of billable items. There are "parent" billing items and "child" billing items. The server billing item is generally referred to as a parent billing item. The items tied to a server, such as ram, harddrives, and operating systems are considered "child" billing items.
func (r Billing_Item) GetObject() (resp datatypes.Billing_Item, err error) {
	err = r.Session.DoRequest("SoftLayer_Billing_Item", "getObject", nil, &r.Options, &resp)
	return
}

// SoftLayer_Billing_Item_Cancellation_Request data type is used to cancel service billing items.
type Billing_Item_Cancellation_Request struct {
	Session session.SLSession
	Options sl.Options
}

// GetBillingItemCancellationRequestService returns an instance of the Billing_Item_Cancellation_Request SoftLayer service
func GetBillingItemCancellationRequestService(sess session.SLSession) Billing_Item_Cancellation_Request {
	return Billing_Item_Cancellation_Request{Session: sess}
}

func (r Billing_Item_Cancellation_Request) Id(id int) Billing_Item_Cancellation_Request {
	r.Options.Id = &id
	return r
}

func (r Billing_Item_Cancellation_Request) Mask(mask string) Billing_Item_Cancellation_Request {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Billing_Item_Cancellation_Request) Filter(filter string) Billing_Item_Cancellation_Request {
	r.Options.Filter = filter
	return r
}

func (r Billing_Item_Cancellation_Request) Limit(limit int) Billing_Item_Cancellation_Request {
	r.Options.Limit = &limit
	return r
}

func (r Billing_Item_Cancellation_Request) Offset(offset int) Billing_Item_Cancellation_Request {
	r.Options.Offset = &offset
	return r
}

// This method returns all service cancellation requests.
//
// Make sure to include the "resultLimit" in the SOAP request header for quicker response. If there is no result limit header is passed, it will return the latest 25 results by default.
func (r Billing_Item_Cancellation_Request) GetAllCancellationRequests() (resp []datatypes.Billing_Item_Cancellation_Request, err error) {
	err = r.Session.DoRequest("SoftLayer_Billing_Item_Cancellation_Request", "getAllCancellationRequests", nil, &r.Options, &resp)
	return
}

// The SoftLayer_Billing_Order data type contains general information relating to an individual order applied to a SoftLayer customer account or to a new customer. Personal information in this type such as names, addresses, and phone numbers are taken from the account's contact information at the time the order is generated for existing SoftLayer customer.
type Billing_Order struct {
	Session session.SLSession
	Options sl.Options
}

// GetBillingOrderService returns an instance of the Billing_Order SoftLayer service
func GetBillingOrderService(sess session.SLSession) Billing_Order {
	return Billing_Order{Session: sess}
}

func (r Billing_Order) Id(id int) Billing_Order {
	r.Options.Id = &id
	return r
}

func (r Billing_Order) Mask(mask string) Billing_Order {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Billing_Order) Filter(filter string) Billing_Order {
	r.Options.Filter = filter
	return r
}

func (r Billing_Order) Limit(limit int) Billing_Order {
	r.Options.Limit = &limit
	return r
}

func (r Billing_Order) Offset(offset int) Billing_Order {
	r.Options.Offset = &offset
	return r
}

// This will get all billing orders for your account.
func (r Billing_Order) GetAllObjects() (resp []datatypes.Billing_Order, err error) {
	err = r.Session.DoRequest("SoftLayer_Billing_Order", "getAllObjects", nil, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Billing_Order object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Billing_Order service. You can only retrieve orders that are assigned to your portal user's account.
func (r Billing_Order) GetObject() (resp datatypes.Billing_Order, err error) {
	err = r.Session.DoRequest("SoftLayer_Billing_Order", "getObject", nil, &r.Options, &resp)
	return
}

// Every individual item that a SoftLayer customer is billed for is recorded in the SoftLayer_Billing_Item data type. Billing items range from server chassis to hard drives to control panels, bandwidth quota upgrades and port upgrade charges. SoftLayer [[SoftLayer_Billing_Invoice|invoices]] are generated from the cost of a customer's billing items. Billing items are copied from the product catalog as they're ordered by customers to create a reference between an account and the billable items they own.
//
// Billing items exist in a tree relationship. Items are associated with each other by parent/child relationships. Component items such as CPU's, RAM, and software each have a parent billing item for the server chassis they're associated with. Billing Items with a null parent item do not have an associated parent item.
type Billing_Order_Item struct {
	Session session.SLSession
	Options sl.Options
}

// GetBillingOrderItemService returns an instance of the Billing_Order_Item SoftLayer service
func GetBillingOrderItemService(sess session.SLSession) Billing_Order_Item {
	return Billing_Order_Item{Session: sess}
}

func (r Billing_Order_Item) Id(id int) Billing_Order_Item {
	r.Options.Id = &id
	return r
}

func (r Billing_Order_Item) Mask(mask string) Billing_Order_Item {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Billing_Order_Item) Filter(filter string) Billing_Order_Item {
	r.Options.Filter = filter
	return r
}

func (r Billing_Order_Item) Limit(limit int) Billing_Order_Item {
	r.Options.Limit = &limit
	return r
}

func (r Billing_Order_Item) Offset(offset int) Billing_Order_Item {
	r.Options.Offset = &offset
	return r
}

// getObject retrieves the SoftLayer_Billing_Item object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Billing_Item service. You can only retrieve billing items tied to the account that your portal user is assigned to. Billing items are an account's items of billable items. There are "parent" billing items and "child" billing items. The server billing item is generally referred to as a parent billing item. The items tied to a server, such as ram, harddrives, and operating systems are considered "child" billing items.
func (r Billing_Order_Item) GetObject() (resp datatypes.Billing_Order_Item, err error) {
	err = r.Session.DoRequest("SoftLayer_Billing_Order_Item", "getObject", nil, &r.Options, &resp)
	return
}

// The SoftLayer_Billing_Oder_Quote data type contains general information relating to an individual order applied to a SoftLayer customer account or to a new customer. Personal information in this type such as names, addresses, and phone numbers are taken from the account's contact information at the time the quote is generated for existing SoftLayer customer.
type Billing_Order_Quote struct {
	Session session.SLSession
	Options sl.Options
}

// GetBillingOrderQuoteService returns an instance of the Billing_Order_Quote SoftLayer service
func GetBillingOrderQuoteService(sess session.SLSession) Billing_Order_Quote {
	return Billing_Order_Quote{Session: sess}
}

func (r Billing_Order_Quote) Id(id int) Billing_Order_Quote {
	r.Options.Id = &id
	return r
}

func (r Billing_Order_Quote) Mask(mask string) Billing_Order_Quote {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Billing_Order_Quote) Filter(filter string) Billing_Order_Quote {
	r.Options.Filter = filter
	return r
}

func (r Billing_Order_Quote) Limit(limit int) Billing_Order_Quote {
	r.Options.Limit = &limit
	return r
}

func (r Billing_Order_Quote) Offset(offset int) Billing_Order_Quote {
	r.Options.Offset = &offset
	return r
}

// Account master users and sub-users in the SoftLayer customer portal can delete the quote of an order.
func (r Billing_Order_Quote) DeleteQuote() (resp datatypes.Billing_Order_Quote, err error) {
	err = r.Session.DoRequest("SoftLayer_Billing_Order_Quote", "deleteQuote", nil, &r.Options, &resp)
	return
}

// getObject retrieves the SoftLayer_Billing_Order_Quote object whose ID number corresponds to the ID number of the init parameter passed to the SoftLayer_Billing_Order_Quote service. You can only retrieve quotes that are assigned to your portal user's account.
func (r Billing_Order_Quote) GetObject() (resp datatypes.Billing_Order_Quote, err error) {
	err = r.Session.DoRequest("SoftLayer_Billing_Order_Quote", "getObject", nil, &r.Options, &resp)
	return
}

// Generate an [[SoftLayer_Container_Product_Order|order container]] from the previously-created quote. This will take into account promotions, reseller status, estimated taxes and all other standard order verification processes.
func (r Billing_Order_Quote) GetRecalculatedOrderContainer(userOrderData *datatypes.Container_Product_Order, orderBeingPlacedFlag *bool) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		userOrderData,
		orderBeingPlacedFlag,
	}
	err = r.Session.DoRequest("SoftLayer_Billing_Order_Quote", "getRecalculatedOrderContainer", params, &r.Options, &resp)
	return
}

// Use this method for placing server orders and additional services orders. The same applies for this as with verifyOrder. Send in the SoftLayer_Container_Product_Order_Hardware_Server for server orders. In addition to verifying the order, placeOrder() also makes an initial authorization on the SoftLayer_Account tied to this order, if a credit card is on file. If the account tied to this order is a paypal customer, an URL will also be returned to the customer. After placing the order, you must go to this URL to finish the authorization process. This tells paypal that you indeed want to place the order. After going to this URL, it will direct you back to a SoftLayer webpage that tells us you have finished the process. After this, it will go to sales for final approval.
func (r Billing_Order_Quote) PlaceOrder(orderData interface{}) (resp datatypes.Container_Product_Order_Receipt, err error) {
	err = datatypes.SetComplexType(orderData)
	if err != nil {
		return
	}
	params := []interface{}{
		orderData,
	}
	err = r.Session.DoRequest("SoftLayer_Billing_Order_Quote", "placeOrder", params, &r.Options, &resp)
	return
}

// Use this method for placing server quotes and additional services quotes. The same applies for this as with verifyOrder. Send in the SoftLayer_Container_Product_Order_Hardware_Server for server quotes. In addition to verifying the quote, placeQuote() also makes an initial authorization on the SoftLayer_Account tied to this order, if a credit card is on file. If the account tied to this order is a paypal customer, an URL will also be returned to the customer. After placing the order, you must go to this URL to finish the authorization process. This tells paypal that you indeed want to place the order. After going to this URL, it will direct you back to a SoftLayer webpage that tells us you have finished the process.
func (r Billing_Order_Quote) PlaceQuote(orderData interface{}) (resp datatypes.Container_Product_Order, err error) {
	params := []interface{}{
		orderData,
	}
	err = r.Session.DoRequest("SoftLayer_Billing_Order_Quote", "placeQuote", params, &r.Options, &resp)
	return
}

// Account master users and sub-users in the SoftLayer customer portal can save the quote of an order to avoid its deletion after 5 days or its expiration after 2 days.
func (r Billing_Order_Quote) SaveQuote() (resp datatypes.Billing_Order_Quote, err error) {
	err = r.Session.DoRequest("SoftLayer_Billing_Order_Quote", "saveQuote", nil, &r.Options, &resp)
	return
}

// Use this method for placing server orders and additional services orders. The same applies for this as with verifyOrder. Send in the SoftLayer_Container_Product_Order_Hardware_Server for server orders. In addition to verifying the order, placeOrder() also makes an initial authorization on the SoftLayer_Account tied to this order, if a credit card is on file. If the account tied to this order is a paypal customer, an URL will also be returned to the customer. After placing the order, you must go to this URL to finish the authorization process. This tells paypal that you indeed want to place the order. After going to this URL, it will direct you back to a SoftLayer webpage that tells us you have finished the process. After this, it will go to sales for final approval.
func (r Billing_Order_Quote) VerifyOrder(orderData interface{}) (resp datatypes.Container_Product_Order, err error) {
	err = datatypes.SetComplexType(orderData)
	if err != nil {
		return
	}
	params := []interface{}{
		orderData,
	}
	err = r.Session.DoRequest("SoftLayer_Billing_Order_Quote", "verifyOrder", params, &r.Options, &resp)
	return
}
