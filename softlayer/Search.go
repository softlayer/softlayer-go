package softlayer

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Search interface {
	Entity

	AdvancedSearch(searchString string) (datatypes.Container_Search_Result, error)
	GetObjectTypes() (datatypes.Container_Search_ObjectType, error)
	Search(searchString string) (datatypes.Container_Search_Result, error)
}
