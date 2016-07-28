package datatypes

import "time"

type FlexibleCredit_Affiliate struct {
	Entity

	FlexibleCreditProgram *FlexibleCredit_Program `json:"flexibleCreditProgram:omitempty"`
	Id                    *int                    `json:"id:omitempty"`
	Name                  *string                 `json:"name:omitempty"`
}

type FlexibleCredit_Company_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
}

type FlexibleCredit_Enrollment struct {
	Entity

	Account                  *Account                     `json:"account:omitempty"`
	AccountId                *int                         `json:"accountId:omitempty"`
	Affiliate                *FlexibleCredit_Affiliate    `json:"affiliate:omitempty"`
	AffiliateId              *int                         `json:"affiliateId:omitempty"`
	AgreementCompleteFlag    *int                         `json:"agreementCompleteFlag:omitempty"`
	CompanyDescription       *string                      `json:"companyDescription:omitempty"`
	CompanyType              *FlexibleCredit_Company_Type `json:"companyType:omitempty"`
	CompanyTypeId            *int                         `json:"companyTypeId:omitempty"`
	EnrollmentDate           *time.Time                   `json:"enrollmentDate:omitempty"`
	FlexibleCreditProgram    *FlexibleCredit_Program      `json:"flexibleCreditProgram:omitempty"`
	GraduationDate           *time.Time                   `json:"graduationDate:omitempty"`
	IsActiveFlag             *bool                        `json:"isActiveFlag:omitempty"`
	MonthlyCreditAmount      *float64                     `json:"monthlyCreditAmount:omitempty"`
	Representative           *User_Employee               `json:"representative:omitempty"`
	RepresentativeEmployeeId *int                         `json:"representativeEmployeeId:omitempty"`
}

type FlexibleCredit_Program struct {
	Entity

	Id      *int    `json:"id:omitempty"`
	KeyName *string `json:"keyName:omitempty"`
	Name    *string `json:"name:omitempty"`
}
