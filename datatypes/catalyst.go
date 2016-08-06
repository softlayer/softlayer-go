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

//
type Catalyst_Affiliate struct {
	Entity

	//
	Id *int `json:"id,omitempty"`

	//
	Name *string `json:"name,omitempty"`

	//
	SkipCreditCardVerificationFlag *bool `json:"skipCreditCardVerificationFlag,omitempty"`
}

//
type Catalyst_Company_Type struct {
	Entity

	//
	Description *string `json:"description,omitempty"`

	//
	Id *int `json:"id,omitempty"`
}

//
type Catalyst_Enrollment struct {
	Entity

	//
	Account *Account `json:"account,omitempty"`

	//
	AccountId *int `json:"accountId,omitempty"`

	//
	Affiliate *Catalyst_Affiliate `json:"affiliate,omitempty"`

	//
	AffiliateId *int `json:"affiliateId,omitempty"`

	//
	AgreementCompleteFlag *int `json:"agreementCompleteFlag,omitempty"`

	//
	CompanyDescription *string `json:"companyDescription,omitempty"`

	//
	CompanyType *Catalyst_Company_Type `json:"companyType,omitempty"`

	//
	CompanyTypeId *int `json:"companyTypeId,omitempty"`

	//
	EnrollmentDate *Time `json:"enrollmentDate,omitempty"`

	//
	GraduationDate *Time `json:"graduationDate,omitempty"`

	//
	IsActiveFlag *bool `json:"isActiveFlag,omitempty"`

	//
	MonthlyCreditAmount *float64 `json:"monthlyCreditAmount,omitempty"`

	//
	Representative *User_Employee `json:"representative,omitempty"`

	//
	RepresentativeEmployeeId *int `json:"representativeEmployeeId,omitempty"`
}

// Contains user information for Catalyst self-enrollment.
type Catalyst_Enrollment_Request struct {
	Entity

	// Applicant's address
	Address1 *string `json:"address1,omitempty"`

	// Additional field for extended address
	Address2 *string `json:"address2,omitempty"`

	//
	Affiliate *Catalyst_Affiliate `json:"affiliate,omitempty"`

	// Id of the affiliate who referred applicant's
	AffiliateId *int `json:"affiliateId,omitempty"`

	//
	AgreementCompleteFlag *bool `json:"agreementCompleteFlag,omitempty"`

	// Determines whether or not to also apply to the GEP program
	ApplyToGepFlag *bool `json:"applyToGepFlag,omitempty"`

	//
	CardAccountNumber *string `json:"cardAccountNumber,omitempty"`

	//
	CardExpirationMonth *string `json:"cardExpirationMonth,omitempty"`

	//
	CardExpirationYear *string `json:"cardExpirationYear,omitempty"`

	//
	CardType *string `json:"cardType,omitempty"`

	//
	CardVerificationNumber *string `json:"cardVerificationNumber,omitempty"`

	// Applicant's city
	City *string `json:"city,omitempty"`

	// Brief description of Startup's product and key differentiators
	CompanyDescription *string `json:"companyDescription,omitempty"`

	// Name of the applicant's company
	CompanyName *string `json:"companyName,omitempty"`

	//
	CompanyType *Catalyst_Company_Type `json:"companyType,omitempty"`

	// Id of the company type which best describes applicant's company
	CompanyTypeId *int `json:"companyTypeId,omitempty"`

	// URL to the Startup's site
	CompanyUrl *string `json:"companyUrl,omitempty"`

	// Applicant's country code
	Country *string `json:"country,omitempty"`

	// Index of answer chosen for how many current users question
	CurrentUserChoice *int `json:"currentUserChoice,omitempty"`

	// Id of the fingerprint
	DeviceFingerprintId *string `json:"deviceFingerprintId,omitempty"`

	// Applicant's email address
	Email *string `json:"email,omitempty"`

	// Applicant's first name
	FirstName *string `json:"firstName,omitempty"`

	// Index of answer chosen for how many future users question
	FutureUserChoice *int `json:"futureUserChoice,omitempty"`

	// Name of accelerator or incubator startup belongs to, if any
	IncubatorName *string `json:"incubatorName,omitempty"`

	// Name of the investor, if any
	InvestorName *string `json:"investorName,omitempty"`

	//
	IpAddress *string `json:"ipAddress,omitempty"`

	// Applicant's last name
	LastName *string `json:"lastName,omitempty"`

	// Applicant's primary phone number
	OfficePhone *string `json:"officePhone,omitempty"`

	// Whether or not the startup has been operating for more than five years
	OverFiveYearsOldFlag *bool `json:"overFiveYearsOldFlag,omitempty"`

	// Applicant's postal code
	PostalCode *string `json:"postalCode,omitempty"`

	// IBM referral code, if any
	ReferralCode *string `json:"referralCode,omitempty"`

	// Whether or not the startup has over one million in annual revenue
	RevenueOverOneMillionFlag *bool `json:"revenueOverOneMillionFlag,omitempty"`

	// Determines whether or not to apply to the Catalyst program
	SkipCatalystApplicationFlag *bool `json:"skipCatalystApplicationFlag,omitempty"`

	// Applicant's state/region code
	State *string `json:"state,omitempty"`

	// Applicant's vatId, if one exists
	VatId *string `json:"vatId,omitempty"`
}

//
type Catalyst_Enrollment_Request_Container_AnswerOption struct {
	Entity

	//
	Answer *string `json:"answer,omitempty"`

	//
	Index *int `json:"index,omitempty"`
}
