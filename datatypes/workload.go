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

/**
 * AUTOMATICALLY GENERATED CODE - DO NOT MODIFY
 */

package datatypes

// no documentation yet
type Workload_Citrix_Client struct {
	Entity
}

// no documentation yet
type Workload_Citrix_Client_Response struct {
	Entity

	// messageId of Citrix account validation response.
	MessageId *string `json:"messageId,omitempty" xmlrpc:"messageId,omitempty"`

	// status of Citrix account validation.
	Status *string `json:"status,omitempty" xmlrpc:"status,omitempty"`

	// status message of Citrix account validation.
	StatusMessage *string `json:"statusMessage,omitempty" xmlrpc:"statusMessage,omitempty"`
}

// no documentation yet
type Workload_Citrix_Client_Response_ResourceLocations struct {
	Workload_Citrix_Client_Response

	// no documentation yet
	ResourceLocations []string `json:"resourceLocations,omitempty" xmlrpc:"resourceLocations,omitempty"`
}

// no documentation yet
type Workload_Citrix_Request struct {
	Entity

	// no documentation yet
	ClientId *string `json:"clientId,omitempty" xmlrpc:"clientId,omitempty"`

	// no documentation yet
	ClientSecret *string `json:"clientSecret,omitempty" xmlrpc:"clientSecret,omitempty"`

	// no documentation yet
	CustomerId *string `json:"customerId,omitempty" xmlrpc:"customerId,omitempty"`
}

// no documentation yet
type Workload_Citrix_Request_CreateResourceLocation struct {
	Workload_Citrix_Request

	// no documentation yet
	ResourceLocationName *string `json:"resourceLocationName,omitempty" xmlrpc:"resourceLocationName,omitempty"`
}

// no documentation yet
type Workload_Citrix_Workspace_Order struct {
	Entity
}

// This is the datatype that needs to be populated and sent to SoftLayer_Workload_Citrix_Workspace_Order::placeWorkspaceOrder.
type Workload_Citrix_Workspace_Order_Container struct {
	Entity

	// The active directory domain name
	ActiveDirectoryDomainName *string `json:"activeDirectoryDomainName,omitempty" xmlrpc:"activeDirectoryDomainName,omitempty"`

	// The active directory netbios name (optional)
	ActiveDirectoryNetbiosName *string `json:"activeDirectoryNetbiosName,omitempty" xmlrpc:"activeDirectoryNetbiosName,omitempty"`

	// The active directory safe mode password
	ActiveDirectorySafeModePassword *string `json:"activeDirectorySafeModePassword,omitempty" xmlrpc:"activeDirectorySafeModePassword,omitempty"`

	// The Citrix API Client Id
	CitrixAPIClientId *string `json:"citrixAPIClientId,omitempty" xmlrpc:"citrixAPIClientId,omitempty"`

	// The Citrix API Client Secret
	CitrixAPIClientSecret *string `json:"citrixAPIClientSecret,omitempty" xmlrpc:"citrixAPIClientSecret,omitempty"`

	// The Citrix customer id
	CitrixCustomerId *string `json:"citrixCustomerId,omitempty" xmlrpc:"citrixCustomerId,omitempty"`

	// The Citrix resource location name
	CitrixResourceLocationName *string `json:"citrixResourceLocationName,omitempty" xmlrpc:"citrixResourceLocationName,omitempty"`

	// The default domain to be used for all server orders where the domain is not specified.
	Domain *string `json:"domain,omitempty" xmlrpc:"domain,omitempty"`

	// The specific [[SoftLayer_Location_Datacenter]] id where the order should be provisioned.
	Location *string `json:"location,omitempty" xmlrpc:"location,omitempty"`

	// There should be one child orderContainer for each component ordered.  The containerIdentifier should be set on each and have these exact values: proxy server bare metal server with hypervisor dhcp server citrix connector servers active directory server vlan subnet storage
	OrderContainers []Container_Product_Order `json:"orderContainers,omitempty" xmlrpc:"orderContainers,omitempty"`
}

// no documentation yet
type Workload_Citrix_Workspace_Response struct {
	Entity

	// messageId associated with any error
	MessageId *string `json:"messageId,omitempty" xmlrpc:"messageId,omitempty"`

	// status of service methods
	Status *string `json:"status,omitempty" xmlrpc:"status,omitempty"`

	// status message
	StatusMessage *string `json:"statusMessage,omitempty" xmlrpc:"statusMessage,omitempty"`
}
