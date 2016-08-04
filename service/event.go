package service

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Event_Log struct {
	Session *Session
	Options
}

func (r *Session) GetEventLogService() Event_Log {
	return Event_Log{Session: r}
}

func (r *Event_Log) GetAllEventNames(objectName *string) (resp []string, err error) {
	params := []interface{}{
		objectName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Event_Log) GetAllEventObjectNames() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Event_Log) GetAllObjects() (resp []datatypes.Event_Log, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Event_Log) GetAllUserTypes() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Event_Log) GetUser() (resp datatypes.User_Customer, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
