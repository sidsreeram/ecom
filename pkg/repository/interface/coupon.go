package interfaces

import (
	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/domain"
)

type CouponRepository interface {
	AddCoupon(coupons helperstruct.Coupons) error
	UpdateCoupon(coupons helperstruct.Coupons, CouponId int) (domain.Coupons, error)
	DeleteCoupon(CouponId int) error
	ViewAllCoupons()([]domain.Coupons,error)
	ViewCoupon(couponId int)(domain.Coupons,error)
	ApplyCoupon(userId int, couponCode string )(int,error)
	RemoveCoupon(userId int)error
}
