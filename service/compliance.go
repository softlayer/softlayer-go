package service

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Compliance_Report_Type struct {
	Session *Session
	Options
}

func (r *Session) GetComplianceReportTypeService() Compliance_Report_Type {
	return Compliance_Report_Type{Session: r}
}

func (r *Compliance_Report_Type) GetAllObjects() (resp []datatypes.Compliance_Report_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Compliance_Report_Type) GetObject() (resp datatypes.Compliance_Report_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
