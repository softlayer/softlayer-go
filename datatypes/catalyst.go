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
type Catalyst_Affiliate struct {
	Entity

	// no documentation yet
	Id *int `json:"id,omitempty"`

	// no documentation yet
	Name *string `json:"name,omitempty"`

	// no documentation yet
	SkipCreditCardVerificationFlag *bool `json:"skipCreditCardVerificationFlag,omitempty"`
}

// no documentation yet
type Catalyst_Company_Type struct {
	Entity

	// no documentation yet
	Description *string `json:"description,omitempty"`

	// no documentation yet
	Id *int `json:"id,omitempty"`
}

// no documentation yet
type Catalyst_Enrollment struct {
	Entity

	// no documentation yet
	Account *Account `json:"account,omitempty"`

	// no documentation yet
	AccountId *int `json:"accountId,omitempty"`

	// no documentation yet
	Affiliate *Catalyst_Affiliate `json:"affiliate,omitempty"`

	// no documentation yet
	AffiliateId *int `json:"affiliateId,omitempty"`

	// no documentation yet
	AgreementCompleteFlag *int `json:"agreementCompleteFlag,omitempty"`

	// no documentation yet
	CompanyDescription *string `json:"companyDescription,omitempty"`

	// no documentation yet
	CompanyType *Catalyst_Company_Type `json:"companyType,omitempty"`

	// no documentation yet
	CompanyTypeId *int `json:"companyTypeId,omitempty"`

	// no documentation yet
	EnrollmentDate *Time `json:"enrollmentDate,omitempty"`

	// no documentation yet
	GraduationDate *Time `json:"graduationDate,omitempty"`

	// no documentation yet
	IsActiveFlag *bool `json:"isActiveFlag,omitempty"`

	// no documentation yet
	MonthlyCreditAmount *Float64 `json:"monthlyCreditAmount,omitempty"`

	// no documentation yet
	Representative *User_Employee `json:"representative,omitempty"`

	// no documentation yet
	RepresentativeEmployeeId *int `json:"representativeEmployeeId,omitempty"`
}

// Contains user information for Catalyst self-enrollment.
type Catalyst_Enrollment_Request struct {
	Entity

	// Applicant's address
	Address1 *string `json:"address1,omitempty"`

	// Additional field for extended address
	Address2 *string `json:"address2,omitempty"`

	// no documentation yet
	Affiliate *Catalyst_Affiliate `json:"affiliate,omitempty"`

	// Id of the affiliate who referred applicant's
	AffiliateId *int `json:"affiliateId,omitempty"`

	// no documentation yet
	AgreementCompleteFlag *bool `json:"agreementCompleteFlag,omitempty"`

	// Determines whether or not to also apply to the GEP program
	ApplyToGepFlag *bool `json:"applyToGepFlag,omitempty"`

	// no documentation yet
	CardAccountNumber *string `json:"cardAccountNumber,omitempty"`

	// no documentation yet
	CardExpirationMonth *string `json:"cardExpirationMonth,omitempty"`

	// no documentation yet
	CardExpirationYear *string `json:"cardExpirationYear,omitempty"`

	// no documentation yet
	CardType *string `json:"cardType,omitempty"`

	// no documentation yet
	CardVerificationNumber *string `json:"cardVerificationNumber,omitempty"`

	// Applicant's city
	City *string `json:"city,omitempty"`

	// Brief description of Startup's product and key differentiators
	CompanyDescription *string `json:"companyDescription,omitempty"`

	// Name of the applicant's company
	CompanyName *string `json:"companyName,omitempty"`

	// no documentation yet
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

	// no documentation yet
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

// no documentation yet
type Catalyst_Enrollment_Request_Container_AnswerOption struct {
	Entity

	// no documentation yet
	Answer *string `json:"answer,omitempty"`

	// no documentation yet
	Index *int `json:"index,omitempty"`
}
