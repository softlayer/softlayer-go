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

// no documentation yet
type Product_Order struct {
	Session session.SLSession
	Options sl.Options
}

// GetProductOrderService returns an instance of the Product_Order SoftLayer service
func GetProductOrderService(sess session.SLSession) Product_Order {
	return Product_Order{Session: sess}
}

func (r Product_Order) Id(id int) Product_Order {
	r.Options.Id = &id
	return r
}

func (r Product_Order) Mask(mask string) Product_Order {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Product_Order) Filter(filter string) Product_Order {
	r.Options.Filter = filter
	return r
}

func (r Product_Order) Limit(limit int) Product_Order {
	r.Options.Limit = &limit
	return r
}

func (r Product_Order) Offset(offset int) Product_Order {
	r.Options.Offset = &offset
	return r
}

// Use this method to place bare metal server, virtual server and additional service orders with SoftLayer.
// Upon success, your credit card or PayPal account will incur charges for the monthly order total
// (or prorated value if ordered mid billing cycle). If all products on the order are only billed hourly,
// you will be charged on your billing anniversary date, which occurs monthly on the day you ordered your first
// service with SoftLayer. For new customers, you are required to provide billing information when you place an order.
// For existing customers, the credit card on file will be charged. If you're a PayPal customer, a URL will be
// returned from the call to [[SoftLayer_Product_Order/placeOrder]] which is to be used to finish the authorization
// process. This authorization tells PayPal that you indeed want to place an order with SoftLayer.
// From PayPal's web site, you will be redirected back to SoftLayer for your order receipt.
//
// When an order is placed, your order will be in a "pending approval" state. When all internal checks pass,
// your order will be automatically approved. For orders that may need extra attention, a Sales representative
// will review the order and contact you if necessary. Once the order is approved, your server or service will
// be provisioned and available to you shortly thereafter. Depending on the type of server or service ordered,
// provisioning times will vary.
//
// ## Order Containers
//
// When placing API orders, it's important to order your server and services on the appropriate
// [[SoftLayer_Container_Product_Order]]. Failing to provide the correct container may delay your server or service
// from being provisioned in a timely manner. Some common order containers are included below.
//
// **Note:** `SoftLayer_Container_Product_Order_` has been removed from the containers in the table below for readability.
//
// | Product | Order Container | Package Type |
// | ------- | --------------- | ------------ |
// | Bare metal server by CPU | [[SoftLayer_Container_Product_Order_Hardware_Server]] | BARE_METAL_CPU |
// | Bare metal server by core | [[SoftLayer_Container_Product_Order_Hardware_Server]] | BARE_METAL_CORE |
// | Virtual server | [[SoftLayer_Container_Product_Order_Virtual_Guest]] | VIRTUAL_SERVER_INSTANCE |
// | Local & dedicated load balancers | [[SoftLayer_Container_Product_Order_Network_LoadBalancer]] | ADDITIONAL_SERVICES_LOAD_BALANCER |
// | Content delivery network | [[SoftLayer_Container_Product_Order_Network_ContentDelivery_Account]] | ADDITIONAL_SERVICES_CDN |
// | Content delivery network Addon | [[SoftLayer_Container_Product_Order_Network_ContentDelivery_Account_Addon]] | ADDITIONAL_SERVICES_CDN_ADDON |
// | Hardware & software firewalls | [[SoftLayer_Container_Product_Order_Network_Protection_Firewall]] | ADDITIONAL_SERVICES_FIREWALL |
// | Dedicated firewall | [[SoftLayer_Container_Product_Order_Network_Protection_Firewall_Dedicated]] | ADDITIONAL_SERVICES_FIREWALL |
// | Object storage | [[SoftLayer_Container_Product_Order_Network_Storage_Object]] | ADDITIONAL_SERVICES_OBJECT_STORAGE |
// | Object storage (hub) | [[SoftLayer_Container_Product_Order_Network_Storage_Hub]] | ADDITIONAL_SERVICES_OBJECT_STORAGE |
// | Network attached storage | [[SoftLayer_Container_Product_Order_Network_Storage_Nas]] | ADDITIONAL_SERVICES_NETWORK_ATTACHED_STORAGE |
// | Iscsi storage | [[SoftLayer_Container_Product_Order_Network_Storage_Iscsi]] | ADDITIONAL_SERVICES_ISCSI_STORAGE |
// | Evault | [[SoftLayer_Container_Product_Order_Network_Storage_Backup_Evault_Vault]] | ADDITIONAL_SERVICES |
// | Evault Plugin | [[SoftLayer_Container_Product_Order_Network_Storage_Backup_Evault_Plugin]] | ADDITIONAL_SERVICES |
// | Application delivery appliance | [[SoftLayer_Container_Product_Order_Network_Application_Delivery_Controller]] | ADDITIONAL_SERVICES_APPLICATION_DELIVERY_APPLIANCE |
// | Network subnet | [[SoftLayer_Container_Product_Order_Network_Subnet]] | ADDITIONAL_SERVICES |
// | Global IPv4 | [[SoftLayer_Container_Product_Order_Network_Subnet]] | ADDITIONAL_SERVICES_GLOBAL_IP_ADDRESSES |
// | Global IPv6 | [[SoftLayer_Container_Product_Order_Network_Subnet]] | ADDITIONAL_SERVICES_GLOBAL_IP_ADDRESSES |
// | Network VLAN | [[SoftLayer_Container_Product_Order_Network_Vlan]] | ADDITIONAL_SERVICES_NETWORK_VLAN |
// | Portable storage | [[SoftLayer_Container_Product_Order_Virtual_Disk_Image]] | ADDITIONAL_SERVICES_PORTABLE_STORAGE |
// | SSL certificate | [[SoftLayer_Container_Product_Order_Security_Certificate]] | ADDITIONAL_SERVICES_SSL_CERTIFICATE |
// | External authentication | [[SoftLayer_Container_Product_Order_User_Customer_External_Binding]] | ADDITIONAL_SERVICES |
// | Dedicated Host | [[SoftLayer_Container_Product_Order_Virtual_DedicatedHost]] | DEDICATED_HOST |
//
// ## Server example
//
// This example includes a single bare metal server being ordered with monthly billing.
//
// **Warning:** the price ids provided below may be outdated or unavailable, so you will need to determine the
//
// available prices from the bare metal server [[SoftLayer_Product_Package/getAllObjects]], which have a
// [[SoftLayer_Product_Package_Type]] of `BARE_METAL_CPU` or `BARE_METAL_CORE`. You can get a full list of
// package types with [[SoftLayer_Product_Package_Type/getAllObjects]].
//
// ### Bare Metal Ordering
//
// ```xml
// <SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="http://api.service.softlayer.com/soap/v3/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
//
//	<SOAP-ENV:Header>
//	  <ns1:authenticate>
//	    <username>your username</username>
//	    <apiKey>your api key</apiKey>
//	  </ns1:authenticate>
//	</SOAP-ENV:Header>
//	<SOAP-ENV:Body>
//	  <ns1:placeOrder>
//	    <orderData xsi:type="ns1:SoftLayer_Container_Product_Order_Hardware_Server">
//	      <hardware SOAP-ENC:arrayType="ns1:SoftLayer_Hardware[1]" xsi:type="ns1:SoftLayer_HardwareArray">
//	        <item xsi:type="ns1:SoftLayer_Hardware">
//	          <domain xsi:type="xsd:string">example.com</domain>
//	          <hostname xsi:type="xsd:string">server1</hostname>
//	        </item>
//	      </hardware>
//	      <location xsi:type="xsd:string">138124</location>
//	      <packageId xsi:type="xsd:int">142</packageId>
//	      <prices SOAP-ENC:arrayType="ns1:SoftLayer_Product_Item_Price[14]" xsi:type="ns1:SoftLayer_Product_Item_PriceArray">
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">58</id>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">22337</id>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">21189</id>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">876</id>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">57</id>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">55</id>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">21190</id>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">36381</id>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">21</id>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">22013</id>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">906</id>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">420</id>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">418</id>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">342</id>
//	        </item>
//	      </prices>
//	      <useHourlyPricing xsi:type="xsd:boolean">false</useHourlyPricing>
//	    </orderData>
//	    <saveAsQuote xsi:nil="true" />
//	  </ns1:placeOrder>
//	</SOAP-ENV:Body>
//
// </SOAP-ENV:Envelope>
// ```
//
// ## Virtual server example
//
// This example includes 2 identical virtual servers (except for hostname) being ordered for hourly billing.
// It includes an optional image template id and VLAN data specified on the virtualGuest objects -
// `primaryBackendNetworkComponent` and `primaryNetworkComponent`.
//
// **Warning:** the price ids provided below may be outdated or unavailable, so you will need to determine the
//
// available prices from the virtual server package with [[SoftLayer_Product_Package/getAllObjects]],
// which has a [[SoftLayer_Product_Package_Type]] of `VIRTUAL_SERVER_INSTANCE`.
//
// #### Virtual Ordering
//
// ```xml
// <SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="http://api.service.softlayer.com/soap/v3/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
//
//	<SOAP-ENV:Header>
//	  <ns1:authenticate>
//	    <username>your username</username>
//	    <apiKey>your api key</apiKey>
//	  </ns1:authenticate>
//	</SOAP-ENV:Header>
//	<SOAP-ENV:Body>
//	  <ns1:placeOrder>
//	    <orderData xsi:type="ns1:SoftLayer_Container_Product_Order_Virtual_Guest">
//	      <imageTemplateId xsi:type="xsd:int">13251</imageTemplateId>
//	      <location xsi:type="xsd:string">37473</location>
//	      <packageId xsi:type="xsd:int">46</packageId>
//	      <prices SOAP-ENC:arrayType="ns1:SoftLayer_Product_Item_Price[13]" xsi:type="ns1:SoftLayer_Product_Item_PriceArray">
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">2159</id>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">55</id>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">13754</id>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">1641</id>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">905</id>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">1800</id>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">58</id>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">21</id>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">1645</id>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">272</id>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">57</id>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">418</id>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">420</id>
//	        </item>
//	      </prices>
//	      <quantity xsi:type="xsd:int">2</quantity>
//	      <useHourlyPricing xsi:type="xsd:boolean">true</useHourlyPricing>
//	      <virtualGuests SOAP-ENC:arrayType="ns1:SoftLayer_Virtual_Guest[1]" xsi:type="ns1:SoftLayer_Virtual_GuestArray">
//	        <item xsi:type="ns1:SoftLayer_Virtual_Guest">
//	          <domain xsi:type="xsd:string">example.com</domain>
//	          <hostname xsi:type="xsd:string">server1</hostname>
//	          <primaryBackendNetworkComponent xsi:type="ns1:SoftLayer_Virtual_Guest_Network_Component">
//	            <networkVlan xsi:type="ns1:SoftLayer_Network_Vlan">
//	              <id xsi:type="xsd:int">12345</id>
//	            </networkVlan>
//	          </primaryBackendNetworkComponent>
//	          <primaryNetworkComponent xsi:type="ns1:SoftLayer_Virtual_Guest_Network_Component">
//	            <networkVlan xsi:type="ns1:SoftLayer_Network_Vlan">
//	              <id xsi:type="xsd:int">67890</id>
//	            </networkVlan>
//	          </primaryNetworkComponent>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Virtual_Guest">
//	          <domain xsi:type="xsd:string">example.com</domain>
//	          <hostname xsi:type="xsd:string">server2</hostname>
//	          <primaryBackendNetworkComponent xsi:type="ns1:SoftLayer_Virtual_Guest_Network_Component">
//	            <networkVlan xsi:type="ns1:SoftLayer_Network_Vlan">
//	              <id xsi:type="xsd:int">12345</id>
//	            </networkVlan>
//	          </primaryBackendNetworkComponent>
//	          <primaryNetworkComponent xsi:type="ns1:SoftLayer_Virtual_Guest_Network_Component">
//	            <networkVlan xsi:type="ns1:SoftLayer_Network_Vlan">
//	              <id xsi:type="xsd:int">67890</id>
//	            </networkVlan>
//	          </primaryNetworkComponent>
//	        </item>
//	      </virtualGuests>
//	    </orderData>
//	    <saveAsQuote xsi:nil="true" />
//	  </ns1:placeOrder>
//	</SOAP-ENV:Body>
//
// </SOAP-ENV:Envelope>
// ```
//
// ## VLAN example
//
// **Warning:** the price ids provided below may be outdated or unavailable, so you will need to determine the
//
// available prices from the additional services pacakge with [[SoftLayer_Product_Package/getAllObjects]],
// which has a [[SoftLayer_Product_Package_Type]] of `ADDITIONAL_SERVICES`.
// You can get a full list of [[SoftLayer_Product_Package_Type/getAllObjects|]] to find other available additional
// service packages.<br/><br/>
//
// ### VLAN Ordering
//
// ```xml
// <SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="http://api.service.softlayer.com/soap/v3/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
//
//	<SOAP-ENV:Header>
//	  <ns1:authenticate>
//	    <username>your username</username>
//	    <apiKey>your api key</apiKey>
//	  </ns1:authenticate>
//	</SOAP-ENV:Header>
//	<SOAP-ENV:Body>
//	  <ns1:placeOrder>
//	    <orderData xsi:type="ns1:SoftLayer_Container_Product_Order_Network_Vlan">
//	      <location xsi:type="xsd:string">154820</location>
//	      <packageId xsi:type="xsd:int">0</packageId>
//	      <prices SOAP-ENC:arrayType="ns1:SoftLayer_Product_Item_Price[2]" xsi:type="ns1:SoftLayer_Product_Item_PriceArray">
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">2021</id>
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Product_Item_Price">
//	          <id xsi:type="xsd:int">2018</id>
//	        </item>
//	      </prices>
//	      <useHourlyPricing xsi:type="xsd:boolean">true</useHourlyPricing>
//	    </orderData>
//	    <saveAsQuote xsi:nil="true" />
//	  </ns1:placeOrder>
//	</SOAP-ENV:Body>
//
// </SOAP-ENV:Envelope>
// ```
//
// ## Multiple products example
//
// This example includes a combination of the above examples in a single order. Note that all the configuration
// options for each individual order container are the same as above, except now we encapsulate each one within
// the `orderContainers` property on the base [[SoftLayer_Container_Product_Order]].
//
// **Warning:** not all products are available to be ordered with other products. For example, since
//
// SSL certificates require validation from a 3rd party, the approval process may take days or even weeks,
// and this would not be acceptable when you need your hourly virtual server right now. To better accommodate
// customers, we restrict several products to be ordered individually.
//
// ### Bare metal server + virtual server + VLAN
//
// ```xml
// <SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="http://api.service.softlayer.com/soap/v3/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
//
//	<SOAP-ENV:Header>
//	  <ns1:authenticate>
//	    <username>your username</username>
//	    <apiKey>your api key</apiKey>
//	  </ns1:authenticate>
//	</SOAP-ENV:Header>
//	<SOAP-ENV:Body>
//	  <ns1:placeOrder>
//	    <orderData xsi:type="ns1:SoftLayer_Container_Product_Order">
//	      <orderContainers SOAP-ENC:arrayType="ns1:SoftLayer_Container_Product_Order[3]" xsi:type="ns1:SoftLayer_Container_Product_OrderArray">
//	        <item xsi:type="ns1:SoftLayer_Container_Product_Order_Hardware_Server">
//	          ...
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Container_Product_Order_Virtual_Guest">
//	          ...
//	        </item>
//	        <item xsi:type="ns1:SoftLayer_Container_Product_Order_Network_Vlan">
//	          ...
//	        </item>
//	      </orderContainers>
//	    </orderData>
//	    <saveAsQuote xsi:nil="true" />
//	  </ns1:placeOrder>
//	</SOAP-ENV:Body>
//
// </SOAP-ENV:Envelope>
// ```
func (r Product_Order) PlaceOrder(orderData interface{}, saveAsQuote *bool) (resp datatypes.Container_Product_Order_Receipt, err error) {
	err = datatypes.SetComplexType(orderData)
	if err != nil {
		return
	}
	params := []interface{}{
		orderData,
		saveAsQuote,
	}
	err = r.Session.DoRequest("SoftLayer_Product_Order", "placeOrder", params, &r.Options, &resp)
	return
}

// Use this method for placing server quotes and additional services quotes. The same applies for this as with verifyOrder. Send in the SoftLayer_Container_Product_Order_Hardware_Server for server quotes. After placing the quote, you must go to this URL to finish the order process. After going to this URL, it will direct you back to a SoftLayer webpage that tells us you have finished the process. After this, it will go to sales for final approval.
func (r Product_Order) PlaceQuote(orderData interface{}) (resp datatypes.Container_Product_Order_Receipt, err error) {
	params := []interface{}{
		orderData,
	}
	err = r.Session.DoRequest("SoftLayer_Product_Order", "placeQuote", params, &r.Options, &resp)
	return
}

// This service is used to verify that an order meets all the necessary requirements to purchase a server, virtual server or service from SoftLayer. It will verify that the products requested do not conflict. For example, you cannot order a Windows firewall with a Linux operating system. It will also check to make sure you have provided all the products that are required for the [[SoftLayer_Product_Package_Order_Configuration]] associated with the [[SoftLayer_Product_Package]] on each of the [[SoftLayer_Container_Product_Order]] specified.<br/><br/>
//
// This service returns the same container that was provided, but with additional information that can be used for debugging or validation. It will also contain pricing information (prorated if applicable) for each of the products on the order. If an exception occurs during verification, a container with the <code>SoftLayer_Exception_Order</code> exception type will be specified in the result.<br/><br/>
//
// <code>verifyOrder</code> accepts the same [[SoftLayer_Container_Product_Order]] as <code>placeOrder</code>, so see [[SoftLayer_Product_Order/placeOrder]] for more details.
func (r Product_Order) VerifyOrder(orderData interface{}) (resp datatypes.Container_Product_Order, err error) {
	err = datatypes.SetComplexType(orderData)
	if err != nil {
		return
	}
	params := []interface{}{
		orderData,
	}
	err = r.Session.DoRequest("SoftLayer_Product_Order", "verifyOrder", params, &r.Options, &resp)
	return
}

// The SoftLayer_Product_Package data type contains information about packages from which orders can be generated. Packages contain general information regarding what is in them, where they are currently sold, availability, and pricing.
type Product_Package struct {
	Session session.SLSession
	Options sl.Options
}

// GetProductPackageService returns an instance of the Product_Package SoftLayer service
func GetProductPackageService(sess session.SLSession) Product_Package {
	return Product_Package{Session: sess}
}

func (r Product_Package) Id(id int) Product_Package {
	r.Options.Id = &id
	return r
}

func (r Product_Package) Mask(mask string) Product_Package {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Product_Package) Filter(filter string) Product_Package {
	r.Options.Filter = filter
	return r
}

func (r Product_Package) Limit(limit int) Product_Package {
	r.Options.Limit = &limit
	return r
}

func (r Product_Package) Offset(offset int) Product_Package {
	r.Options.Offset = &offset
	return r
}

// Retrieve The preset configurations available only for the authenticated account and this package.
func (r Product_Package) GetAccountRestrictedActivePresets() (resp []datatypes.Product_Package_Preset, err error) {
	err = r.Session.DoRequest("SoftLayer_Product_Package", "getAccountRestrictedActivePresets", nil, &r.Options, &resp)
	return
}

// Retrieve The available preset configurations for this package.
func (r Product_Package) GetActivePresets() (resp []datatypes.Product_Package_Preset, err error) {
	err = r.Session.DoRequest("SoftLayer_Product_Package", "getActivePresets", nil, &r.Options, &resp)
	return
}

// This method pulls all the active packages. This will give you a basic description of the packages that are currently active
func (r Product_Package) GetAllObjects() (resp []datatypes.Product_Package, err error) {
	err = r.Session.DoRequest("SoftLayer_Product_Package", "getAllObjects", nil, &r.Options, &resp)
	return
}

// Retrieve The item categories associated with a package, including information detailing which item categories are required as part of a SoftLayer product order.
func (r Product_Package) GetConfiguration() (resp []datatypes.Product_Package_Order_Configuration, err error) {
	err = r.Session.DoRequest("SoftLayer_Product_Package", "getConfiguration", nil, &r.Options, &resp)
	return
}

// Retrieve A collection of SoftLayer_Product_Item_Prices that are valid for this package.
func (r Product_Package) GetItemPrices() (resp []datatypes.Product_Item_Price, err error) {
	err = r.Session.DoRequest("SoftLayer_Product_Package", "getItemPrices", nil, &r.Options, &resp)
	return
}

// Retrieve A collection of valid items available for purchase in this package.
func (r Product_Package) GetItems() (resp []datatypes.Product_Item, err error) {
	err = r.Session.DoRequest("SoftLayer_Product_Package", "getItems", nil, &r.Options, &resp)
	return
}

// Retrieve The regional locations that a package is available in.
func (r Product_Package) GetRegions() (resp []datatypes.Location_Region, err error) {
	err = r.Session.DoRequest("SoftLayer_Product_Package", "getRegions", nil, &r.Options, &resp)
	return
}

// Package presets are used to simplify ordering by eliminating the need for price ids when submitting orders.
//
// Orders submitted with a preset id defined will use the prices included in the package preset. Prices submitted on an order with a preset id will replace the prices included in the package preset for that prices category. If the package preset has a fixed configuration flag <em>(fixedConfigurationFlag)</em> set then the prices included in the preset configuration cannot be replaced by prices submitted on the order. The only exception to the fixed configuration flag would be if a price submitted on the order is an account-restricted price for the same product item.
type Product_Package_Preset struct {
	Session session.SLSession
	Options sl.Options
}

// GetProductPackagePresetService returns an instance of the Product_Package_Preset SoftLayer service
func GetProductPackagePresetService(sess session.SLSession) Product_Package_Preset {
	return Product_Package_Preset{Session: sess}
}

func (r Product_Package_Preset) Id(id int) Product_Package_Preset {
	r.Options.Id = &id
	return r
}

func (r Product_Package_Preset) Mask(mask string) Product_Package_Preset {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Product_Package_Preset) Filter(filter string) Product_Package_Preset {
	r.Options.Filter = filter
	return r
}

func (r Product_Package_Preset) Limit(limit int) Product_Package_Preset {
	r.Options.Limit = &limit
	return r
}

func (r Product_Package_Preset) Offset(offset int) Product_Package_Preset {
	r.Options.Offset = &offset
	return r
}

// no documentation yet
func (r Product_Package_Preset) GetObject() (resp datatypes.Product_Package_Preset, err error) {
	err = r.Session.DoRequest("SoftLayer_Product_Package_Preset", "getObject", nil, &r.Options, &resp)
	return
}
