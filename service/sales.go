package service

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Sales_Presale_Event struct {
	Session *Session
	Options
}

func (r *Session) GetSalesPresaleEventService() Sales_Presale_Event {
	return Sales_Presale_Event{Session: r}
}

func (r *Sales_Presale_Event) GetAllObjects() (resp []datatypes.Sales_Presale_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Sales_Presale_Event) GetObject() (resp datatypes.Sales_Presale_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Sales_Presale_Event) GetActiveFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Sales_Presale_Event) GetExpiredFlag() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Sales_Presale_Event) GetItem() (resp datatypes.Product_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Sales_Presale_Event) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Sales_Presale_Event) GetOrders() (resp []datatypes.Billing_Order, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
