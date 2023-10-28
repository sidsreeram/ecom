package repository

import (
	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/domain"
	interfaces "github.com/ECOMMERCE_PROJECT/pkg/repository/interface"
	"gorm.io/gorm"
)

type CouponDatatbase struct {
	DB *gorm.DB
}

func NewCouponRepository(DB *gorm.DB) interfaces.CouponRepository {
	return &CouponDatatbase{DB}
}
func (c *CouponDatatbase) AddCoupon(coupons helperstruct.Coupons) error {
	createCoupon := `INSERT INTO coupons(code, discount_precent,discount_maximum_amount,minimum_purchase_amount,expiration_date)VALUES($1,$2,$3,$4,$5)`
	err := c.DB.Exec(createCoupon, coupons.Code, coupons.DiscountPercent, coupons.DiscountMaximumAmount, coupons.MinimumPurchaseAmount, coupons.ExpirationDate).Error

	return err

}
func (c*CouponDatatbase) UpdateCoupon(coupons helperstruct.Coupons,CouponId int)(domain.Coupons,error){
	var updatedCoupon domain.Coupons
	updatequery:=`UPDATE coupons SET code=$1,discount_percent=$2,discount_maximum_amount=$3,minimum_purchase_amount=$4,expiration_date=$5 WHERE id=$6 RETURNING *`
	err:=c.DB.Raw(updatequery,coupons.Code,coupons.DiscountPercent,coupons.DiscountMaximumAmount,coupons.MinimumPurchaseAmount,coupons.ExpirationDate,CouponId).Scan(&updatedCoupon).Error
	return updatedCoupon,err
}
func (c*CouponDatatbase) DeleteCoupon(CouponId int)error{
	deleteCoupon:=`DELETE 	FROM coupons WHERE id=?`
	err:=c.DB.Raw(deleteCoupon,CouponId).Error
	return err

}
func (c*CouponDatatbase) ViewAllCoupons()([]domain.Coupons,error){
	var couponss []domain.Coupons
	query:=`SELECT *FROM coupons `
	err:=c.DB.Raw(query).Scan(&couponss).Error
	return couponss,err
}
func (c*CouponDatatbase) ViewCoupon(couponId int)(domain.Coupons,error){
	var coupon domain.Coupons
	query:=`SELECT * FROM coupons WHERE id=?`
	err:=c.DB.Raw(query,couponId).Scan(&coupon).Error
	return coupon,err
}
