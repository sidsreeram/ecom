package usecase

import (
	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/domain"
	interfaces "github.com/ECOMMERCE_PROJECT/pkg/repository/interface"
	services "github.com/ECOMMERCE_PROJECT/pkg/usecase/interface"
)

type Couponusecase struct {
	couponrepository interfaces.CouponRepository
}

func NewCouponUsecase(couponrepo interfaces.CouponRepository) services.Couponusecase {
	return &Couponusecase{couponrepository: couponrepo}
}
func (c *Couponusecase) AddCoupon(coupons helperstruct.Coupons) error {
	err := c.couponrepository.AddCoupon(coupons)
	return err
}
func (c *Couponusecase) UpdateCoupon(coupons helperstruct.Coupons, CouponId int) (domain.Coupons, error){
	updatedcoupon,err:=c.couponrepository.UpdateCoupon(coupons,CouponId)
	return updatedcoupon,err
}
func (c*Couponusecase) DeleteCoupon(CouponId int) error{
	err:=c.couponrepository.DeleteCoupon(CouponId)
	return err
}
func (c*Couponusecase) ViewAllCoupons()([]domain.Coupons,error){
	coupons ,err:=c.couponrepository.ViewAllCoupons()
	return coupons,err
}
func (c*Couponusecase) ViewCoupon(couponId int)(domain.Coupons,error){
	coupon,err:=c.couponrepository.ViewCoupon(couponId)
	return coupon,err
}
