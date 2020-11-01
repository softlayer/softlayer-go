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

// The SoftLayer_Monitoring_Robot data type contains general information relating to a monitoring robot.
type Monitoring_Robot struct {
	Entity

	// The account associated with the corresponding robot.
	Account *Account `json:"account,omitempty" xmlrpc:"account,omitempty"`

	// Internal identifier of a SoftLayer account that this robot belongs to
	AccountId *int `json:"accountId,omitempty" xmlrpc:"accountId,omitempty"`

	// Internal identifier of a monitoring robot
	Id *int `json:"id,omitempty" xmlrpc:"id,omitempty"`

	// Robot name
	Name *string `json:"name,omitempty" xmlrpc:"name,omitempty"`

	// The current status of the robot.
	RobotStatus *Monitoring_Robot_Status `json:"robotStatus,omitempty" xmlrpc:"robotStatus,omitempty"`

	// The SoftLayer_Software_Component that corresponds to the robot installation on the server.
	SoftwareComponent *Software_Component `json:"softwareComponent,omitempty" xmlrpc:"softwareComponent,omitempty"`

	// Internal identifier of a monitoring robot status
	StatusId *int `json:"statusId,omitempty" xmlrpc:"statusId,omitempty"`
}

// Your monitoring robot will be in "Active" status under normal circumstances. If you perform an OS reload, your robot will be in "Reclaim" status until it's reloaded on your server or virtual server.
//
// Advanced monitoring system requires "Nimsoft Monitoring (Advanced)" service running and TCP ports 48000 - 48020 to be open on your server or virtual server. Monitoring agents cannot be managed nor can the usage data be updated if these ports are closed. Your monitoring robot will be in "Limited Connectivity" status if our monitoring management system cannot communicate with your system.
//
// See [[SoftLayer_Monitoring_Robot::resetStatus|resetStatus]] service for more details.
type Monitoring_Robot_Status struct {
	Entity

	// Monitoring robot status description
	Description *string `json:"description,omitempty" xmlrpc:"description,omitempty"`

	// Internal identifier of a monitoring robot status
	Id *int `json:"id,omitempty" xmlrpc:"id,omitempty"`

	// Monitoring robot status name
	Name *string `json:"name,omitempty" xmlrpc:"name,omitempty"`
}
