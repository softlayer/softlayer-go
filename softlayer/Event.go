package softlayer

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Event_Log interface {
	Entity

	GetAllEventNames(objectName string) (string, error)
	GetAllEventObjectNames() (string, error)
	GetAllObjects() (datatypes.Event_Log, error)
	GetAllUserTypes() (string, error)

	GetUser() (datatypes.User_Customer, error)
}
