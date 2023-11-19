package usecase

import (
	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/domain"
	interfaces "github.com/ECOMMERCE_PROJECT/pkg/repository/interface"
	services "github.com/ECOMMERCE_PROJECT/pkg/usecase/interface"
)

type DiscountUsecase struct {
	discountrepo interfaces.DiscountRepository
}

func NewDiscountUsecase(discountrepoo interfaces.DiscountRepository) services.DiscountUsecase {
	return &DiscountUsecase{
		discountrepo: discountrepoo,
	}
}

func (c *DiscountUsecase) AddDiscount(discount helperstruct.Discount) error {
   err:=c.discountrepo.AddDiscount(discount)
   return err
}
func (c *DiscountUsecase) UpdateDiscount(id int, discount helperstruct.Discount) (domain.Discount, error) {
   discountdetails,err:=c.discountrepo.UpdateDiscount(id,discount)
   return discountdetails,err
}
func (c *DiscountUsecase) DeleteDiscount(id int) error {
  err:=c.discountrepo.DeleteDiscount(id)
  return err
}
func (c *DiscountUsecase) GetAllDiscount() ([]domain.Discount, error) {
   discount,err:=c.discountrepo.GetAllDiscount()
   return discount,err
}
func (c *DiscountUsecase) ViewDiscountbyID(id int) (domain.Discount, error) {
dis,err:=c.discountrepo.ViewDiscountbyID(id)
return dis,err
}
