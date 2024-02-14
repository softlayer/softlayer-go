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

	"github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/sl"
)

// no documentation yet
type Resource_Metadata struct {
	Session session.SLSession
	Options sl.Options
}

// GetResourceMetadataService returns an instance of the Resource_Metadata SoftLayer service
func GetResourceMetadataService(sess session.SLSession) Resource_Metadata {
	return Resource_Metadata{Session: sess}
}

func (r Resource_Metadata) Id(id int) Resource_Metadata {
	r.Options.Id = &id
	return r
}

func (r Resource_Metadata) Mask(mask string) Resource_Metadata {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r Resource_Metadata) Filter(filter string) Resource_Metadata {
	r.Options.Filter = filter
	return r
}

func (r Resource_Metadata) Limit(limit int) Resource_Metadata {
	r.Options.Limit = &limit
	return r
}

func (r Resource_Metadata) Offset(offset int) Resource_Metadata {
	r.Options.Offset = &offset
	return r
}

// The getBackendMacAddresses method retrieves a list of backend MAC addresses for the resource
func (r Resource_Metadata) GetBackendMacAddresses() (resp []string, err error) {
	err = r.Session.DoRequest("SoftLayer_Resource_Metadata", "getBackendMacAddresses", nil, &r.Options, &resp)
	return
}

// The getDatacenter method retrieves the name of the datacenter in which the resource is located.
func (r Resource_Metadata) GetDatacenter() (resp string, err error) {
	err = r.Session.DoRequest("SoftLayer_Resource_Metadata", "getDatacenter", nil, &r.Options, &resp)
	return
}

// The getDatacenterId retrieves the ID for the datacenter in which the resource is located.
func (r Resource_Metadata) GetDatacenterId() (resp int, err error) {
	err = r.Session.DoRequest("SoftLayer_Resource_Metadata", "getDatacenterId", nil, &r.Options, &resp)
	return
}

// The getFrontendMacAddresses method retrieves a list of frontend MAC addresses for the resource
func (r Resource_Metadata) GetFrontendMacAddresses() (resp []string, err error) {
	err = r.Session.DoRequest("SoftLayer_Resource_Metadata", "getFrontendMacAddresses", nil, &r.Options, &resp)
	return
}

// The getFullyQualifiedDomainName method provides the user with a combined return which includes the hostname and domain for the resource. Because this method returns multiple pieces of information, it avoids the need to use multiple methods to return the desired information.
func (r Resource_Metadata) GetFullyQualifiedDomainName() (resp string, err error) {
	err = r.Session.DoRequest("SoftLayer_Resource_Metadata", "getFullyQualifiedDomainName", nil, &r.Options, &resp)
	return
}

// The getId method retrieves the ID for the resource
func (r Resource_Metadata) GetId() (resp int, err error) {
	err = r.Session.DoRequest("SoftLayer_Resource_Metadata", "getId", nil, &r.Options, &resp)
	return
}

// The getPrimaryBackendIpAddress method retrieves the primary backend IP address for the resource
func (r Resource_Metadata) GetPrimaryBackendIpAddress() (resp string, err error) {
	err = r.Session.DoRequest("SoftLayer_Resource_Metadata", "getPrimaryBackendIpAddress", nil, &r.Options, &resp)
	return
}

// The getPrimaryIpAddress method retrieves the primary IP address for the resource. For resources with a frontend network, the frontend IP address will be returned. For resources that have been provisioned with only a backend network, the backend IP address will be returned, as a frontend address will not exist.
func (r Resource_Metadata) GetPrimaryIpAddress() (resp string, err error) {
	err = r.Session.DoRequest("SoftLayer_Resource_Metadata", "getPrimaryIpAddress", nil, &r.Options, &resp)
	return
}

// The getProvisionState method retrieves the provision state of the resource. The provision state may be used to determine when it is considered safe to perform additional setup operations. The method returns 'PROCESSING' to indicate the provision is in progress and 'COMPLETE' when the provision is complete.
func (r Resource_Metadata) GetProvisionState() (resp string, err error) {
	err = r.Session.DoRequest("SoftLayer_Resource_Metadata", "getProvisionState", nil, &r.Options, &resp)
	return
}

// The getRouter method will return the router associated with a network component. When the router is redundant, the hostname of the redundant group will be returned, rather than the router hostname.
func (r Resource_Metadata) GetRouter(macAddress *string) (resp string, err error) {
	params := []interface{}{
		macAddress,
	}
	err = r.Session.DoRequest("SoftLayer_Resource_Metadata", "getRouter", params, &r.Options, &resp)
	return
}

// The getTags method retrieves all tags associated with the resource. Tags are single keywords assigned to a resource that assist the user in identifying the resource and its roles when performing a simple search. Tags are assigned by any user with access to the resource.
func (r Resource_Metadata) GetTags() (resp []string, err error) {
	err = r.Session.DoRequest("SoftLayer_Resource_Metadata", "getTags", nil, &r.Options, &resp)
	return
}

// The getUserMetadata method retrieves metadata completed by users who interact with the resource. Metadata gathered using this method is unique to parameters set using the ”'setUserMetadata”' method, which must be executed prior to completing this method. User metadata may also be provided while placing an order for a resource.
func (r Resource_Metadata) GetUserMetadata() (resp string, err error) {
	err = r.Session.DoRequest("SoftLayer_Resource_Metadata", "getUserMetadata", nil, &r.Options, &resp)
	return
}

// The getVlanIds method returns a list of VLAN IDs for the network component matching the provided MAC address associated with the resource. For each return, the native VLAN will appear first, followed by any trunked VLANs associated with the network component.
func (r Resource_Metadata) GetVlanIds(macAddress *string) (resp []int, err error) {
	params := []interface{}{
		macAddress,
	}
	err = r.Session.DoRequest("SoftLayer_Resource_Metadata", "getVlanIds", params, &r.Options, &resp)
	return
}

// The getVlans method returns a list of VLAN numbers for the network component matching the provided MAC address associated with the resource. For each return, the native VLAN will appear first, followed by any trunked VLANs associated with the network component.
func (r Resource_Metadata) GetVlans(macAddress *string) (resp []int, err error) {
	params := []interface{}{
		macAddress,
	}
	err = r.Session.DoRequest("SoftLayer_Resource_Metadata", "getVlans", params, &r.Options, &resp)
	return
}
