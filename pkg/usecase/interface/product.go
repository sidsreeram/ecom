package interfaces

import (
	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
)

type ProductUsecase interface {
	CreateCategory(category helperstruct.Category) (response.Category, error)
	UpdateCategory(category helperstruct.Category,id int)(response.Category,error)
	DeleteCategory(id int)error
	ListCategories()([]response.Category,error)
}
