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
type FlexibleCredit_Affiliate struct {
	Entity

	// Flexible Credit Program the affiliate belongs to.
	FlexibleCreditProgram *FlexibleCredit_Program `json:"flexibleCreditProgram,omitempty"`

	// Primary ID for the affiliate
	Id *int `json:"id,omitempty"`

	// Name of this affiliate
	Name *string `json:"name,omitempty"`
}

// no documentation yet
type FlexibleCredit_Company_Type struct {
	Entity

	// Description of the company type
	Description *string `json:"description,omitempty"`

	// Primary ID for the company type
	Id *int `json:"id,omitempty"`
}

// no documentation yet
type FlexibleCredit_Enrollment struct {
	Entity

	// Account the enrollment belongs to
	Account *Account `json:"account,omitempty"`

	// Account ID associated with this enrollment
	AccountId *int `json:"accountId,omitempty"`

	// Affiliate associated with the account enrollment
	Affiliate *FlexibleCredit_Affiliate `json:"affiliate,omitempty"`

	// ID of the corresponding Flexible Credit Program Affiliate
	AffiliateId *int `json:"affiliateId,omitempty"`

	// Indicates signing of Flexible Credit agreement (independent from MSA)
	AgreementCompleteFlag *int `json:"agreementCompleteFlag,omitempty"`

	// Brief description of the company
	CompanyDescription *string `json:"companyDescription,omitempty"`

	// Category which best describes the company
	CompanyType *FlexibleCredit_Company_Type `json:"companyType,omitempty"`

	// ID of the Flexible Credit Program Company classification for this enrollment
	CompanyTypeId *int `json:"companyTypeId,omitempty"`

	// Date when participation in the Flexible Credit program began
	EnrollmentDate *Time `json:"enrollmentDate,omitempty"`

	// Discount program the enrollment belongs to
	FlexibleCreditProgram *FlexibleCredit_Program `json:"flexibleCreditProgram,omitempty"`

	// Date Flexible Credit Program benefits end.
	GraduationDate *Time `json:"graduationDate,omitempty"`

	// Flag indicating whether an enrollment is active (true) or inactive (false)
	IsActiveFlag *bool `json:"isActiveFlag,omitempty"`

	// Amount of monthly credit (USD) given to the account
	MonthlyCreditAmount *float64 `json:"monthlyCreditAmount,omitempty"`

	// Employee overseeing the enrollment
	Representative *User_Employee `json:"representative,omitempty"`

	// ID of the employee representing this account.
	RepresentativeEmployeeId *int `json:"representativeEmployeeId,omitempty"`
}

// no documentation yet
type FlexibleCredit_Program struct {
	Entity

	// Primary ID of the Flexible Credit Program
	Id *int `json:"id,omitempty"`

	// Unique name for the Flexible Credit Program
	KeyName *string `json:"keyName,omitempty"`

	// Name of the Flexible Credit Program.
	Name *string `json:"name,omitempty"`
}
