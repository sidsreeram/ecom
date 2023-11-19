package repository

import (
	"fmt"

	"time"

	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/domain"
	interfaces "github.com/ECOMMERCE_PROJECT/pkg/repository/interface"
	"gorm.io/gorm"
)

type CouponDatabase struct {
	DB *gorm.DB
}

func NewCouponRepository(DB *gorm.DB) interfaces.CouponRepository {
	return &CouponDatabase{DB}
}
func (c *CouponDatabase) AddCoupon(coupons helperstruct.Coupons) error {
	createCoupon := `INSERT INTO coupons(code, discount_percent, discount_maximum_amount, minimum_purchase_amount, expiration_date) VALUES ($1, $2, $3, $4, $5)`
	err := c.DB.Exec(createCoupon, coupons.Code, coupons.DiscountPercent, coupons.DiscountMaximumAmount, coupons.MinimumPurchaseAmount, coupons.ExpirationDate).Error

	return err

}
func (c *CouponDatabase) UpdateCoupon(coupons helperstruct.Coupons, CouponId int) (domain.Coupons, error) {
	var updatedCoupon domain.Coupons
	updatequery := `UPDATE coupons SET code=$1,discount_percent=$2,discount_maximum_amount=$3,minimum_purchase_amount=$4,expiration_date=$5 WHERE id=$6 RETURNING *`
	err := c.DB.Raw(updatequery, coupons.Code, coupons.DiscountPercent, coupons.DiscountMaximumAmount, coupons.MinimumPurchaseAmount, coupons.ExpirationDate, CouponId).Scan(&updatedCoupon).Error
	return updatedCoupon, err
}
func (c *CouponDatabase) DeleteCoupon(CouponId int) error {
	deleteCoupon := `DELETE 	FROM coupons WHERE id=?`
	err := c.DB.Raw(deleteCoupon, CouponId).Error
	return err

}
func (c *CouponDatabase) ViewAllCoupons() ([]domain.Coupons, error) {
	var couponss []domain.Coupons
	query := `SELECT *FROM coupons `
	err := c.DB.Raw(query).Scan(&couponss).Error
	return couponss, err
}
func (c *CouponDatabase) ViewCoupon(couponId int) (domain.Coupons, error) {
	var coupon domain.Coupons
	query := `SELECT * FROM coupons WHERE id=?`
	err := c.DB.Raw(query, couponId).Scan(&coupon).Error
	return coupon, err
}
func (c *CouponDatabase) ApplyCoupon(userId int, couponCode string) (int, error) {
	tx := c.DB.Begin()
	var coupon domain.Coupons
	findcoupon := `SELECT * FROM coupons WHERE code=?`
	err := tx.Raw(findcoupon, couponCode).Scan(&coupon).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if coupon.Id == 0 {
		return 0, fmt.Errorf("No Coupon Found")
	}
	if coupon.ExpirationDate.Before(time.Now()) {
		tx.Rollback()
		return 0, fmt.Errorf("Coupon has expired")

	}
	var isUsed bool
	checkused := `SELECT EXISTS(SELECT 1 FROM orders WHERE user_id= $1 AND coupon_id=$2)`

	err = tx.Raw(checkused, userId, coupon.Id).Scan(&isUsed).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if isUsed {
		updatecart := `UPDATE carts SET coupon_id = 0 WHERE user_id = ?`
		err = tx.Exec(updatecart, userId).Error
		if err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("Failed to update cart: %v", err)
		}
		return 0, fmt.Errorf("The coupon was already userhtrytretrtrdtd")
	}
	
	var cartdetails domain.Carts
	Getcartdetails := `SELECT *FROM carts WHERE user_id=?`
	err = tx.Raw(Getcartdetails, userId).Scan(&cartdetails).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if cartdetails.CouponId == coupon.Id {
		tx.Rollback()
		return 0, fmt.Errorf("Coupon is Applied")
	}
	if cartdetails.SubTotal == 0 {
		tx.Rollback()
		return 0, fmt.Errorf("There is no products in your cart please add products to apply coupon")
	}
	if cartdetails.SubTotal <= int(coupon.MinimumPurchaseAmount) {
		topurchasemore := coupon.MinimumPurchaseAmount - float64(cartdetails.SubTotal)

		tx.Rollback()
		return 0, fmt.Errorf("Purchase %v more to apply coupon", topurchasemore)
	}

	discountAmt := int(float64(cartdetails.SubTotal) * (coupon.DiscountPercent / 100.0))
	if discountAmt > int(coupon.DiscountMaximumAmount) {
		discountAmt = int(coupon.DiscountMaximumAmount)
	}

	cartdetails.Total = cartdetails.SubTotal - discountAmt
	UpdateCart := `UPDATE carts SET total=$1 , coupon_id=$2  WHERE id=$3`
	err = tx.Exec(UpdateCart, cartdetails.Total, coupon.Id, cartdetails.Id).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	return cartdetails.Total, nil

}
func (c *CouponDatabase) RemoveCoupon(userId int) error {
	tx := c.DB.Begin()
	//get the details of the cart
	var cartDetails domain.Carts
	getCartDetails := `SELECT * FROM carts WHERE user_id=$1`
	err := tx.Raw(getCartDetails, userId).Scan(&cartDetails).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	//find any coupon is added
	if cartDetails.CouponId == 0 {
		tx.Rollback()
		return fmt.Errorf("no coupon to remove")
	}
	//if added remove the coupon
	removeCoupon := `UPDATE carts SET coupon_id=0, total = sub_total WHERE user_id=$1`
	err = tx.Exec(removeCoupon, userId).Error
	if cartDetails.CouponId == 0 {
		tx.Rollback()
		return err
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
