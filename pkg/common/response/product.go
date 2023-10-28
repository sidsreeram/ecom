package response

import "time"

type Category struct {
	Id           int
	CategoryName string
}

type Product struct {
	Id           int `json:",omitempty"`
	Name         string
	Description  string
	Brand        string
	CategoryName string
}
type ProductItem struct {
	ID         uint
	ProductID  uint
	Product    Product
	SKU        string
	QtyInStock int
	Price      float64 
	InStock    bool    
	Color      string  
	Size       string  
	Rating     float64 
	Image         []string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
