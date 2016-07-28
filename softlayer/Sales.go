package softlayer

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Sales_Presale_Event interface {
	Entity

	GetAllObjects() (datatypes.Sales_Presale_Event, error)
	GetObject() (datatypes.Sales_Presale_Event, error)

	GetActiveFlag() (bool, error)
	GetExpiredFlag() (bool, error)
	GetItem() (datatypes.Product_Item, error)
	GetLocation() (datatypes.Location, error)
	GetOrders() (datatypes.Billing_Order, error)
}
