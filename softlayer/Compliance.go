package softlayer

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Compliance_Report_Type interface {
	Entity

	GetAllObjects() (datatypes.Compliance_Report_Type, error)
	GetObject() (datatypes.Compliance_Report_Type, error)
}
