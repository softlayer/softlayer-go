package softlayer

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Legal_RegulatedWorkload interface {
	Entity

	GetAccount() (datatypes.Account, error)
	GetType() (datatypes.Legal_RegulatedWorkload_Type, error)
}

type Legal_RegulatedWorkload_Type interface {
	Entity
}
