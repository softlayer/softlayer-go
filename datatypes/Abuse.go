package datatypes

type Abuse_Lockdown_Resource struct {
	Entity

	Account     *Account              `json:"account:omitempty"`
	InvoiceItem *Billing_Invoice_Item `json:"invoiceItem:omitempty"`
}
