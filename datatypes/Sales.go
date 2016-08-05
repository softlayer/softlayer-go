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

package datatypes

type Sales_Presale_Event struct {
	Entity

	ActiveFlag  *bool           `json:"activeFlag,omitempty"`
	Description *string         `json:"description,omitempty"`
	EndDate     *Time           `json:"endDate,omitempty"`
	ExpiredFlag *bool           `json:"expiredFlag,omitempty"`
	Id          *int            `json:"id,omitempty"`
	Item        *Product_Item   `json:"item,omitempty"`
	ItemId      *int            `json:"itemId,omitempty"`
	Location    *Location       `json:"location,omitempty"`
	LocationId  *int            `json:"locationId,omitempty"`
	OrderCount  *uint           `json:"orderCount,omitempty"`
	Orders      []Billing_Order `json:"orders,omitempty"`
	StartDate   *Time           `json:"startDate,omitempty"`
}
