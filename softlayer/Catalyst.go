package softlayer

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Catalyst_Affiliate interface {
	Entity
}

type Catalyst_Company_Type interface {
	Entity

	GetAllObjects() (datatypes.Catalyst_Company_Type, error)
	GetObject() (datatypes.Catalyst_Company_Type, error)
}

type Catalyst_Enrollment interface {
	Entity

	GetAffiliates() (datatypes.Catalyst_Affiliate, error)
	GetCompanyTypes() (datatypes.Catalyst_Company_Type, error)
	GetEnrollmentRequestAnnualRevenueOptions() (datatypes.Catalyst_Enrollment_Request_Container_AnswerOption, error)
	GetEnrollmentRequestUserCountOptions() (datatypes.Catalyst_Enrollment_Request_Container_AnswerOption, error)
	GetEnrollmentRequestYearsInOperationOptions() (datatypes.Catalyst_Enrollment_Request_Container_AnswerOption, error)
	GetObject() (datatypes.Catalyst_Enrollment, error)
	RequestManualEnrollment(request datatypes.Container_Catalyst_ManualEnrollmentRequest) error
	RequestSelfEnrollment(enrollmentRequest datatypes.Catalyst_Enrollment_Request) (datatypes.Account, error)

	GetAccount() (datatypes.Account, error)
	GetAffiliate() (datatypes.Catalyst_Affiliate, error)
	GetCompanyType() (datatypes.Catalyst_Company_Type, error)
	GetIsActiveFlag() (bool, error)
	GetRepresentative() (datatypes.User_Employee, error)
}

type Catalyst_Enrollment_Request interface {
	Entity

	GetAffiliate() (datatypes.Catalyst_Affiliate, error)
	GetCompanyType() (datatypes.Catalyst_Company_Type, error)
}

type Catalyst_Enrollment_Request_Container_AnswerOption interface {
	Entity
}
