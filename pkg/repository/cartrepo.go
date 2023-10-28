package repository

import (
	"fmt"

	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
	"github.com/ECOMMERCE_PROJECT/pkg/domain"
	interfaces "github.com/ECOMMERCE_PROJECT/pkg/repository/interface"
	"gorm.io/gorm"
)

type cartDatabase struct {
	DB *gorm.DB
}

func NewCartRepository(DB *gorm.DB) interfaces.CartRepository {
	return &cartDatabase{DB}
}
func (c *cartDatabase) CreateCart(id int) error {
    query := `INSERT INTO carts (user_id, sub_total, total, coupon_id) VALUES ($1, 0, 0, 0)`
    err := c.DB.Exec(query, id).Error
    return err
}

func (c *cartDatabase) AddtoCart(productId, userId int) error {
	tx := c.DB.Begin()

	var cartId int
	findcartid := `SELECT id from carts WHERE user_id=?`
	err := tx.Raw(findcartid, userId).Scan(&cartId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	var CartItemId int
	cartItemcheck := `SELECT id FROM cart_items WHERE carts_id =$1 AND product_item_id = $2 LIMIT 1`
	err = tx.Raw(cartItemcheck, cartId, productId).Scan(&CartItemId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if CartItemId == 0 {
		addtocart := `INSERT INTO cart_items(carts_id,product_item_id,quantity)VALUES($1,$2,1)`
		err = tx.Raw(addtocart, cartId, productId).Scan(&CartItemId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		updatecart := `UPDATE cart_items SET quantity = cart_items.quantity+1 WHERE id=$1`
		err = tx.Raw(updatecart, CartItemId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	var price int
	findprice := `SELECT price FROM product_items WHERE id=$1`
	err = tx.Raw(findprice, productId).Scan(&price).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	var subtotal int
	updateprice := `UPDATE carts SET sub_total=carts.sub_total+$1 WHERE user_id =$2 RETURNING sub_total`
	err = tx.Raw(updateprice, price, userId).Scan(&subtotal).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// var couponId int
	// findcoupon := `SELECT coupon_id FROM carts WHERE user_id=$1`
	// err = tx.Raw(findcoupon, userId).Scan(&couponId).Error
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	// if couponId != 0 {
	// 	var coupons domain.Coupons
	// 	getCouponDetails := `SELECT * FROM coupons WHERE id=$1`
	// 	err := tx.Raw(getCouponDetails, couponId).Scan(&coupons).Error
	// 	if err != nil {
	// 		tx.Rollback()
	// 		return err
	// 	}
	// 	discountAmount := (subtotal / 100) * int(coupons.DiscountPercent)
	// 	if discountAmount > int(coupons.MinimumPurchaseAmount) {
	// 		discountAmount = int(coupons.DiscountMaximumAmount)
	// 	}
	// 	updateTotal := `UPDATE carts SET total=$1 WHERE id=$2`
	// 	err = tx.Raw(updateTotal, subtotal-discountAmount, cartId).Error
	// 	if err != nil {
	// 		tx.Rollback()
	// 		return err
	// 	}
	// } else {
	// 	updateTotal := `UPDATE carts SET total=$1 WHERE id=$2`
	// 	err = tx.Raw(updateTotal, subtotal, cartId).Error
	// 	if err != nil {
	// 		tx.Rollback()
	// 		return err
	// 	}
	// }
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil

}
func (c *cartDatabase) RemoveFromCart(userId, productId int) error {
	tx := c.DB.Begin()

	var cartId int
	findcartid := `SELECT id from carts WHERE user_id=?`
	err := tx.Raw(findcartid, userId).Scan(&cartId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	var qty int
	findquantity := `SELECT quantity from cart_items WHERE carts_id=$1 AND product_item_id=$2`
	err = tx.Raw(findquantity, cartId, productId).Scan(&qty).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if qty == 0 {
		tx.Rollback()
		return fmt.Errorf("no items to remove from cart")
	}
	if qty == 1 {
		deleteitem := `DELETE FROM cart_items WHERE carts_id=$1 AND product_item_id=$2`
		err := tx.Raw(deleteitem, cartId, productId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		updateQty := `UPDATE cart_items SET quantity=cart_items.quantity-1 WHERE carts_id=$1 AND product_item_id=$2`
		err = tx.Raw(updateQty, cartId, productId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
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
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
func (c *cartDatabase) ListCart(userId int) (response.ViewCart, error) {
    tx := c.DB.Begin()
    
    // Get cart details
    type cartDetails struct {
        Id       int
        SubTotal float64
        Total    float64
    }
    var cart cartDetails
    getCartDetails := `SELECT
        c.id,
        c.sub_total,
        c.total
        FROM carts c WHERE c.user_id = $1`
    err := tx.Raw(getCartDetails, userId).Scan(&cart).Error

    if err != nil {
        tx.Rollback()
        return response.ViewCart{}, err
    }

    // Get cart_items details
    var cartItems []domain.CartItem
    getCartItemsDetails := `SELECT * FROM cart_items WHERE carts_id = $1`
    err = tx.Raw(getCartItemsDetails, cart.Id).Scan(&cartItems).Error
    if err != nil {
        tx.Rollback()
        return response.ViewCart{}, err
    }

    // Get the product details
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
