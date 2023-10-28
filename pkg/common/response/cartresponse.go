package response

type DisplayCart struct {
	Productname  string
	Brand        string
	Color        string
    Size          int
	Quantity     uint
	PricePerUnit float64
	Total        float64
}

type ViewCart struct {
	CartItems  []DisplayCart `json:"cart_items"`
	Couponcode string        `json:"couponCode"`
	SubTotal   float64       `json:"sub_total"`
	Discount   float64       `json:"discount"`
	CartTotal  float64       `json:"cart_total"`
}