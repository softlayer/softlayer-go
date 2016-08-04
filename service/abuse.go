package service

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Abuse_Lockdown_Resource struct {
	Session *Session
	Options
}

func (r *Session) GetAbuseLockdownResourceService() Abuse_Lockdown_Resource {
	return Abuse_Lockdown_Resource{Session: r}
}

func (r *Abuse_Lockdown_Resource) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Abuse_Lockdown_Resource) GetInvoiceItem() (resp datatypes.Billing_Invoice_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
