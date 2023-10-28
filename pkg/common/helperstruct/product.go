package helperstruct

type Category struct {
	Name string `json:"name" validate:"required"`
}

type Product struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Brand       string `json:"brand" validate:"required"`
	CategoryId  string `json:"categoryid" validate:"required"`
}

type ProductItem struct {
	Product_id uint    `json:"productid"`
	Sku        string  `json:"sku"`
	Qty        int     `json:"quantity"`
	Color      string  `json:"colour"`
	Size       string  `json:"size"`
	Price      int     `json:"price"`
	Instock    bool    `json:"instock"`
	Rating     float64 `json:"rating"`
	Imag       string  `json:"image"`
}
type QueryParams struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	Query    string `json:"query"`   //search key word
	Filter   string `json:"filter"`  //to specify the column name
	SortBy   string `json:"sort_by"` //to specify column to set the sorting
	SortDesc bool   `json:"sort_desc"`
}
