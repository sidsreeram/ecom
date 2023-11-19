package helperstruct

import "time"

type Discount struct {
	Id                    uint      `json:"id"`
	CategoryId            uint      `json:"category_id"`
	Category              Category  `json:""`
	DiscountPercent       float64   `json:"discountpercent"`
	DiscountMaximumAmount float64   `json:"discountmaximumamount"`
	MinimumPurchaseAmount float64   `json:"minimumpurchaseamount"`
	ExpirationDate        time.Time `json:"expirationdate"`
}
