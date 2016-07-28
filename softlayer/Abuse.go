package softlayer

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Abuse_Lockdown_Resource interface {
	Entity

	GetAccount() (datatypes.Account, error)
	GetInvoiceItem() (datatypes.Billing_Invoice_Item, error)
}
