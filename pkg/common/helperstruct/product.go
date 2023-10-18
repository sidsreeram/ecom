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
	Product_id  uint    `json:"productid"`
	Sku         string  `json:"sku"`
	Qty         int     `json:"quantity"`
	Color       string  `json:"colour"`
	Ram         int     `json:"ram"`
	Battery     int     `json:"battery"`
	Screen_size float64 `json:"screensize"`
	Storage     int     `json:"storage"`
	Camera      int     `json:"camera"`
	Price       int     `json:"price"`
	Imag        string  `json:"image"`
}
type QueryParams struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	Query    string `json:"query"`   //search key word
	Filter   string `json:"filter"`  //to specify the column name
	SortBy   string `json:"sort_by"` //to specify column to set the sorting
	SortDesc bool   `json:"sort_desc"`
}