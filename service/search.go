package service

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Search struct {
	Session *Session
	Options
}

func (r *Session) GetSearchService() Search {
	return Search{Session: r}
}

func (r *Search) AdvancedSearch(searchString *string) (resp []datatypes.Container_Search_Result, err error) {
	params := []interface{}{
		searchString,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Search) GetObjectTypes() (resp []datatypes.Container_Search_ObjectType, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Search) Search(searchString *string) (resp []datatypes.Container_Search_Result, err error) {
	params := []interface{}{
		searchString,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
