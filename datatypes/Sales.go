package datatypes

import "time"

type Sales_Presale_Event struct {
	Entity

	ActiveFlag  *bool           `json:"activeFlag:omitempty"`
	Description *string         `json:"description:omitempty"`
	EndDate     *time.Time      `json:"endDate:omitempty"`
	ExpiredFlag *bool           `json:"expiredFlag:omitempty"`
	Id          *int            `json:"id:omitempty"`
	Item        *Product_Item   `json:"item:omitempty"`
	ItemId      *int            `json:"itemId:omitempty"`
	Location    *Location       `json:"location:omitempty"`
	LocationId  *int            `json:"locationId:omitempty"`
	OrderCount  *uint           `json:"orderCount:omitempty"`
	Orders      []Billing_Order `json:"orders:omitempty"`
	StartDate   *time.Time      `json:"startDate:omitempty"`
}
