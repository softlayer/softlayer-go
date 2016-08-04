package service

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Legal_RegulatedWorkload struct {
	Session *Session
	Options
}

func (r *Session) GetLegalRegulatedWorkloadService() Legal_RegulatedWorkload {
	return Legal_RegulatedWorkload{Session: r}
}

func (r *Legal_RegulatedWorkload) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Legal_RegulatedWorkload) GetType() (resp datatypes.Legal_RegulatedWorkload_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Legal_RegulatedWorkload_Type struct {
	Session *Session
	Options
}

func (r *Session) GetLegalRegulatedWorkloadTypeService() Legal_RegulatedWorkload_Type {
	return Legal_RegulatedWorkload_Type{Session: r}
}
