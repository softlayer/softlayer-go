package datatypes

import "time"

type Catalyst_Affiliate struct {
	Entity

	Id                             *int    `json:"id:omitempty"`
	Name                           *string `json:"name:omitempty"`
	SkipCreditCardVerificationFlag *bool   `json:"skipCreditCardVerificationFlag:omitempty"`
}

type Catalyst_Company_Type struct {
	Entity

	Description *string `json:"description:omitempty"`
	Id          *int    `json:"id:omitempty"`
}

type Catalyst_Enrollment struct {
	Entity

	Account                  *Account               `json:"account:omitempty"`
	AccountId                *int                   `json:"accountId:omitempty"`
	Affiliate                *Catalyst_Affiliate    `json:"affiliate:omitempty"`
	AffiliateId              *int                   `json:"affiliateId:omitempty"`
	AgreementCompleteFlag    *int                   `json:"agreementCompleteFlag:omitempty"`
	CompanyDescription       *string                `json:"companyDescription:omitempty"`
	CompanyType              *Catalyst_Company_Type `json:"companyType:omitempty"`
	CompanyTypeId            *int                   `json:"companyTypeId:omitempty"`
	EnrollmentDate           *time.Time             `json:"enrollmentDate:omitempty"`
	GraduationDate           *time.Time             `json:"graduationDate:omitempty"`
	IsActiveFlag             *bool                  `json:"isActiveFlag:omitempty"`
	MonthlyCreditAmount      *float64               `json:"monthlyCreditAmount:omitempty"`
	Representative           *User_Employee         `json:"representative:omitempty"`
	RepresentativeEmployeeId *int                   `json:"representativeEmployeeId:omitempty"`
}

type Catalyst_Enrollment_Request struct {
	Entity

	Address1                    *string                `json:"address1:omitempty"`
	Address2                    *string                `json:"address2:omitempty"`
	Affiliate                   *Catalyst_Affiliate    `json:"affiliate:omitempty"`
	AffiliateId                 *int                   `json:"affiliateId:omitempty"`
	AgreementCompleteFlag       *bool                  `json:"agreementCompleteFlag:omitempty"`
	ApplyToGepFlag              *bool                  `json:"applyToGepFlag:omitempty"`
	CardAccountNumber           *string                `json:"cardAccountNumber:omitempty"`
	CardExpirationMonth         *string                `json:"cardExpirationMonth:omitempty"`
	CardExpirationYear          *string                `json:"cardExpirationYear:omitempty"`
	CardType                    *string                `json:"cardType:omitempty"`
	CardVerificationNumber      *string                `json:"cardVerificationNumber:omitempty"`
	City                        *string                `json:"city:omitempty"`
	CompanyDescription          *string                `json:"companyDescription:omitempty"`
	CompanyName                 *string                `json:"companyName:omitempty"`
	CompanyType                 *Catalyst_Company_Type `json:"companyType:omitempty"`
	CompanyTypeId               *int                   `json:"companyTypeId:omitempty"`
	CompanyUrl                  *string                `json:"companyUrl:omitempty"`
	Country                     *string                `json:"country:omitempty"`
	CurrentUserChoice           *int                   `json:"currentUserChoice:omitempty"`
	DeviceFingerprintId         *string                `json:"deviceFingerprintId:omitempty"`
	Email                       *string                `json:"email:omitempty"`
	FirstName                   *string                `json:"firstName:omitempty"`
	FutureUserChoice            *int                   `json:"futureUserChoice:omitempty"`
	IncubatorName               *string                `json:"incubatorName:omitempty"`
	InvestorName                *string                `json:"investorName:omitempty"`
	IpAddress                   *string                `json:"ipAddress:omitempty"`
	LastName                    *string                `json:"lastName:omitempty"`
	OfficePhone                 *string                `json:"officePhone:omitempty"`
	OverFiveYearsOldFlag        *bool                  `json:"overFiveYearsOldFlag:omitempty"`
	PostalCode                  *string                `json:"postalCode:omitempty"`
	ReferralCode                *string                `json:"referralCode:omitempty"`
	RevenueOverOneMillionFlag   *bool                  `json:"revenueOverOneMillionFlag:omitempty"`
	SkipCatalystApplicationFlag *bool                  `json:"skipCatalystApplicationFlag:omitempty"`
	State                       *string                `json:"state:omitempty"`
	VatId                       *string                `json:"vatId:omitempty"`
}

type Catalyst_Enrollment_Request_Container_AnswerOption struct {
	Entity

	Answer *string `json:"answer:omitempty"`
	Index  *int    `json:"index:omitempty"`
}
