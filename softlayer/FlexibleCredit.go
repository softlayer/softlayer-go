package softlayer

import "github.ibm.com/riethm/gopherlayer/datatypes"

type FlexibleCredit_Affiliate interface {
	Entity

	GetFlexibleCreditProgram() (datatypes.FlexibleCredit_Program, error)
}

type FlexibleCredit_Company_Type interface {
	Entity
}

type FlexibleCredit_Enrollment interface {
	Entity

	GetAccount() (datatypes.Account, error)
	GetAffiliate() (datatypes.FlexibleCredit_Affiliate, error)
	GetCompanyType() (datatypes.FlexibleCredit_Company_Type, error)
	GetFlexibleCreditProgram() (datatypes.FlexibleCredit_Program, error)
	GetIsActiveFlag() (bool, error)
	GetRepresentative() (datatypes.User_Employee, error)
}

type FlexibleCredit_Program interface {
	Entity

	GetAffiliatesAvailableForSelfEnrollmentByVerificationType(verificationTypeKeyName string) (datatypes.FlexibleCredit_Affiliate, error)
	GetCompanyTypes() (datatypes.FlexibleCredit_Company_Type, error)
	GetObject() (datatypes.FlexibleCredit_Program, error)
	SelfEnrollNewAccount(accountTemplate datatypes.Account) (datatypes.Account, error)
}
