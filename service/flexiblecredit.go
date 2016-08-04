package service

import "github.ibm.com/riethm/gopherlayer/datatypes"

type FlexibleCredit_Affiliate struct {
	Session *Session
	Options
}

func (r *Session) GetFlexibleCreditAffiliateService() FlexibleCredit_Affiliate {
	return FlexibleCredit_Affiliate{Session: r}
}

func (r *FlexibleCredit_Affiliate) GetFlexibleCreditProgram() (resp datatypes.FlexibleCredit_Program, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type FlexibleCredit_Company_Type struct {
	Session *Session
	Options
}

func (r *Session) GetFlexibleCreditCompanyTypeService() FlexibleCredit_Company_Type {
	return FlexibleCredit_Company_Type{Session: r}
}

type FlexibleCredit_Enrollment struct {
	Session *Session
	Options
}

func (r *Session) GetFlexibleCreditEnrollmentService() FlexibleCredit_Enrollment {
	return FlexibleCredit_Enrollment{Session: r}
}

func (r *FlexibleCredit_Enrollment) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *FlexibleCredit_Enrollment) GetAffiliate() (resp datatypes.FlexibleCredit_Affiliate, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *FlexibleCredit_Enrollment) GetCompanyType() (resp datatypes.FlexibleCredit_Company_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *FlexibleCredit_Enrollment) GetFlexibleCreditProgram() (resp datatypes.FlexibleCredit_Program, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *FlexibleCredit_Enrollment) GetIsActiveFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *FlexibleCredit_Enrollment) GetRepresentative() (resp datatypes.User_Employee, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type FlexibleCredit_Program struct {
	Session *Session
	Options
}

func (r *Session) GetFlexibleCreditProgramService() FlexibleCredit_Program {
	return FlexibleCredit_Program{Session: r}
}

func (r *FlexibleCredit_Program) GetAffiliatesAvailableForSelfEnrollmentByVerificationType(verificationTypeKeyName *string) (resp []datatypes.FlexibleCredit_Affiliate, err error) {
	params := []interface{}{
		verificationTypeKeyName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *FlexibleCredit_Program) GetCompanyTypes() (resp []datatypes.FlexibleCredit_Company_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *FlexibleCredit_Program) GetObject() (resp datatypes.FlexibleCredit_Program, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *FlexibleCredit_Program) SelfEnrollNewAccount(accountTemplate *datatypes.Account) (resp datatypes.Account, err error) {
	params := []interface{}{
		accountTemplate,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
