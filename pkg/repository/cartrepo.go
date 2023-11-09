package repository

import (
	"fmt"
	

	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
	"github.com/ECOMMERCE_PROJECT/pkg/domain"
	interfaces "github.com/ECOMMERCE_PROJECT/pkg/repository/interface"
	"gorm.io/gorm"
)

type CartDatabase struct {
	DB *gorm.DB
}

func NewCartRepository(DB *gorm.DB) interfaces.CartRepository {
	return &CartDatabase{DB}
}

func (c *CartDatabase) CreateCart(id int) error {
	query := `INSERT INTO carts (user_id, sub_total,total,coupon_id) VALUES($1,0,0,0)`
	err := c.DB.Exec(query, id).Error
	return err
}		

func (c *CartDatabase) AddToCart(productId, userId int) error {
	tx := c.DB.Begin()
	//finding cart id coresponding to the user
	var cartId int
	findCartId := `SELECT id FROM carts WHERE user_id=? `
	err := tx.Raw(findCartId, userId).Scan(&cartId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//Check whether the product exists in the cart_items
	var cartItemId int
	cartItemCheck := `SELECT id FROM cart_items WHERE carts_id = $1 AND product_item_id = $2 LIMIT 1`
	err = tx.Raw(cartItemCheck, cartId, productId).Scan(&cartItemId).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if cartItemId == 0 {
		addToCart := `INSERT INTO cart_items (carts_id,product_item_id,quantity)VALUES($1,$2,1)`
		err = tx.Exec(addToCart, cartId, productId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		updatCart := `UPDATE cart_items SET quantity = cart_items.quantity+1 WHERE id = $1 `
		err = tx.Exec(updatCart, cartItemId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	//finding the price of the product
	var price int
	findPrice := `SELECT price FROM product_items WHERE id=$1`
	err = tx.Raw(findPrice, productId).Scan(&price).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	//Updating the subtotal in cart table
	var subtotal int
	updateSubTotal := `UPDATE carts SET sub_total=carts.sub_total+$1 WHERE user_id=$2 RETURNING sub_total`
	err = tx.Raw(updateSubTotal, price, userId).Scan(&subtotal).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	//check any coupon is present inside the cart
	var couponId int
	findCoupon := `SELECT coupon_id FROM carts WHERE user_id=$1`
	err = tx.Raw(findCoupon, userId).Scan(&couponId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if couponId != 0 {
		//find the coupon details
		var coupon domain.Coupons
		getCouponDetails := `SELECT * FROM coupons WHERE id=$1`
		err := tx.Raw(getCouponDetails, couponId).Scan(&coupon).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		//applay the coupon to the total
		discountAmount := (subtotal / 100) * int(coupon.DiscountPercent)
		if discountAmount > int(coupon.DiscountMaximumAmount) {
			discountAmount = int(coupon.DiscountMaximumAmount)
		}
		updateTotal := `UPDATE carts SET total=$1 where id=$2`
		err = tx.Exec(updateTotal, subtotal-discountAmount, cartId).Error
		if err != nil {
			tx.Rollback()
			return err
		}

	} else {
		updateTotal := `UPDATE carts SET total=$1 where id=$2`
		err = tx.Exec(updateTotal, subtotal, cartId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (c *CartDatabase) RemoveFromCart(userId, productId int) error {
	tx := c.DB.Begin()

	//Find cart id
	var cartId int
	findCartId := `SELECT id FROM carts WHERE user_id=? `
	err := tx.Raw(findCartId, userId).Scan(&cartId).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	//Find the qty of the product in cart
	var qty int
	findQty := `SELECT quantity FROM cart_items WHERE carts_id=$1 AND product_item_id=$2`
	err = tx.Raw(findQty, cartId, productId).Scan(&qty).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if qty == 0 {
		tx.Rollback()
		return fmt.Errorf("no items in cart to reomve")
	}

	//If the qty is 1 dlt the product from the cart
	if qty == 1 {
		dltItem := `DELETE FROM cart_items WHERE carts_id=$1 AND product_item_id=$2`
		err := tx.Exec(dltItem, cartId, productId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	} else { // If there is  more than one product reduce the qty by 1
		updateQty := `UPDATE cart_items SET quantity=cart_items.quantity-1 WHERE carts_id=$1 AND product_item_id=$2`
		err = tx.Exec(updateQty, cartId, productId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	//Find the price of the product item
	var price int
	productPrice := `SELECT price FROM product_items WHERE id=$1`
	err = tx.Raw(productPrice, productId).Scan(&price).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//Update the subtotal reduce the price of the cart total with price of the product
	var subTotal int
	updateSubTotal := `UPDATE carts SET sub_total=sub_total-$1 WHERE user_id=$2 RETURNING sub_total`
	err = tx.Raw(updateSubTotal, price, userId).Scan(&subTotal).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	//Check any coupon is added to the cart
	var couponId int
	findCoupon := `SELECT coupon_id from carts where id=$1`
	err = tx.Raw(findCoupon, cartId).Scan(&couponId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//If coupon is added check the subtotal meet the minimum value to applay coupon
	if couponId != 0 {
		//Find the coupon details
		var couponDetails domain.Coupons
		findCouponDetails := `SELECT * FROM coupons WHERE id=$1`
		err = tx.Raw(findCouponDetails, couponId).Scan(&couponDetails).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		if couponDetails.MinimumPurchaseAmount > float64(subTotal) {
			//if sub total is less than the minimum value needed set remove the coupon from the cart and set the subtoal as total
			updateTotal := `UPDATE carts SET total=$1 WHERE id=$2`
			err = tx.Exec(updateTotal, subTotal, cartId).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		} else {
			//applay the coupon to the total
			discountAmount := (subTotal / 100) * int(couponDetails.DiscountPercent)
			if discountAmount > int(couponDetails.DiscountMaximumAmount) {
				discountAmount = int(couponDetails.DiscountMaximumAmount)
			}
			updateTotal := `UPDATE carts SET total=$1 WHERE id=$2`
			err = tx.Exec(updateTotal, subTotal-discountAmount, cartId).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	//If yes applay the coupon to the subtotal and add it as total

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (c *CartDatabase) ListCart(userId int) (response.ViewCart, error) {
	tx := c.DB.Begin()
	//get cart details
	type cartDetails struct {
		Id         int
		SubTotal   float64
		Total      float64
		Couponcode string
	}
	var cart cartDetails
	getCartDetails := `SELECT
		c.id,
		c.sub_total,
		c.total,
		co.code AS couponcode
		FROM carts c LEFT JOIN coupons co ON c.coupon_id=co.id WHERE c.user_id=$1`
	err := tx.Raw(getCartDetails, userId).Scan(&cart).Error

	if err != nil {
		tx.Rollback()
		return response.ViewCart{}, err
	}
	//get cart_items details
	var cartItems domain.CartItem
	getCartItemsDetails := `SELECT * FROM cart_items WHERE carts_id=$1`
	err = tx.Raw(getCartItemsDetails, cart.Id).Scan(&cartItems).Error
	if err != nil {
		tx.Rollback()
		return response.ViewCart{}, err
	}
	//get the product details
	var details []response.DisplayCart
    getDetails := `SELECT p.brand, pi.sku AS productname, 
        pi.color,
        pi.size,
        pi.rating,
        ci.quantity,
        pi.price AS price_per_unit,
        (pi.price * ci.quantity) AS total
        FROM cart_items ci JOIN product_items pi  ON ci.product_item_id = pi.id
        JOIN products p ON pi.product_id = p.id WHERE ci.carts_id = $1`
	err = tx.Raw(getDetails, cart.Id).Scan(&details).Error
	if err != nil {
		tx.Rollback()
		return response.ViewCart{}, err
	}

	var carts response.ViewCart
	carts.Couponcode = cart.Couponcode
	carts.CartTotal = cart.Total
	carts.SubTotal = cart.SubTotal
	carts.CartItems = details
	carts.Discount = cart.SubTotal - cart.Total
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return response.ViewCart{}, err
	}
	return carts, nil
}
