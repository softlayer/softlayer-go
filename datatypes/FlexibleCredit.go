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

import "time"

type FlexibleCredit_Affiliate struct {
	Entity

	FlexibleCreditProgram *FlexibleCredit_Program `json:"flexibleCreditProgram,omitempty"`
	Id                    *int                    `json:"id,omitempty"`
	Name                  *string                 `json:"name,omitempty"`
}

type FlexibleCredit_Company_Type struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
}

type FlexibleCredit_Enrollment struct {
	Entity

	Account                  *Account                     `json:"account,omitempty"`
	AccountId                *int                         `json:"accountId,omitempty"`
	Affiliate                *FlexibleCredit_Affiliate    `json:"affiliate,omitempty"`
	AffiliateId              *int                         `json:"affiliateId,omitempty"`
	AgreementCompleteFlag    *int                         `json:"agreementCompleteFlag,omitempty"`
	CompanyDescription       *string                      `json:"companyDescription,omitempty"`
	CompanyType              *FlexibleCredit_Company_Type `json:"companyType,omitempty"`
	CompanyTypeId            *int                         `json:"companyTypeId,omitempty"`
	EnrollmentDate           *time.Time                   `json:"enrollmentDate,omitempty"`
	FlexibleCreditProgram    *FlexibleCredit_Program      `json:"flexibleCreditProgram,omitempty"`
	GraduationDate           *time.Time                   `json:"graduationDate,omitempty"`
	IsActiveFlag             *bool                        `json:"isActiveFlag,omitempty"`
	MonthlyCreditAmount      *float64                     `json:"monthlyCreditAmount,omitempty"`
	Representative           *User_Employee               `json:"representative,omitempty"`
	RepresentativeEmployeeId *int                         `json:"representativeEmployeeId,omitempty"`
}

type FlexibleCredit_Program struct {
	Entity

	Id      *int    `json:"id,omitempty"`
	KeyName *string `json:"keyName,omitempty"`
	Name    *string `json:"name,omitempty"`
}
