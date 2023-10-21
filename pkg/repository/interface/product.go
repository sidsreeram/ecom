package interfaces

import (
	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
)

type ProductRepository interface {
	CreateCategory(category helperstruct.Category) (response.Category, error)
	UpdateCategory(category helperstruct.Category, id int) (response.Category, error)
	DeleteCategory(id int) error
	ListCategories() ([]response.Category, error)
	DisplayACategory(id int)(response.Category,error)
	AddProduct(product helperstruct.Product)(response.Product,error)
	UpdateProduct(id int, product helperstruct.Product) (response.Product, error) 
	DeleteProduct(id int)error
	AddProductitem(productItem helperstruct.ProductItem) (response.ProductItem, error)
	UpdateProductItem(id int, product helperstruct.ProductItem) (response.ProductItem, error)
	DeleteProductItem(id int)error
	DisplayAproductitem(id int)(response.ProductItem,error)
}
