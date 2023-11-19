package interfaces

import (
	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/domain"
)

type DiscountRepository interface {
	AddDiscount(discount helperstruct.Discount) error
	UpdateDiscount(id int, discount helperstruct.Discount) (domain.Discount, error)
	DeleteDiscount(id int)error
	GetAllDiscount()([]domain.Discount,error)
	ViewDiscountbyID(id int)(domain.Discount,error)
}
