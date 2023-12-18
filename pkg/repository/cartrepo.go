package repository

import (
	"fmt"
	"time"

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

	// finding cart id corresponding to the user
	var cartId int
	findCartId := `SELECT id FROM carts WHERE user_id=? `
	err := tx.Raw(findCartId, userId).Scan(&cartId).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// Check whether the product exists in the cart_items
	var cartItemId int
	cartItemCheck := `SELECT id FROM cart_items WHERE carts_id = $1 AND product_item_id = $2 LIMIT 1`
	err = tx.Raw(cartItemCheck, cartId, productId).Scan(&cartItemId).Error
	fmt.Println(cartItemId)
	if err != nil {
		tx.Rollback()
		return err
	}

	if cartItemId == 0 {

		// Insert the new product into cart_items
		addToCart := `INSERT INTO cart_items (carts_id, product_item_id, quantity) VALUES ($1, $2, 1)`
		err = tx.Exec(addToCart, cartId, productId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		// Update the quantity of the existing product in cart_items
		updateCart := `UPDATE cart_items SET quantity = cart_items.quantity+1 WHERE id = $1`
		err = tx.Exec(updateCart, cartItemId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Finding the price of the product
	// ...

	// Finding the price of the product
	var proitems domain.ProductItem
	findPrice := `SELECT * FROM product_items WHERE id=$1`
	err = tx.Raw(findPrice, productId).Scan(&proitems).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	var productdetails domain.Product
	findcategory := `SELECT * FROM products WHERE id = ?`
	err = tx.Raw(findcategory, proitems.ProductID).Scan(&productdetails).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	var discount domain.Discount
	finddiscount := `SELECT * FROM discounts WHERE category_id = ?`
	err = tx.Raw(finddiscount, productdetails.CategoryID).Scan(&discount).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// if discount.Id == 0 {
	// 	tx.Rollback()
	// 	return fmt.Errorf("no offer found ")
	// }

	if discount.ExpirationDate.Before(time.Now()) {
		tx.Rollback()
		return fmt.Errorf("the offer has expired")
	}

	var total int
	updateSubTotal := `UPDATE carts SET sub_total=carts.sub_total+$1 WHERE user_id=$2 RETURNING sub_total`
	err = tx.Raw(updateSubTotal, proitems.Price, userId).Scan(&total).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if discount.MinimumPurchaseAmount > 0 && total < int(discount.MinimumPurchaseAmount) {
		tx.Rollback()
		return fmt.Errorf("the minimum purchase amount condition is not net for the offer")
	}

	// Calculating the discounted price
	discountedPrice := total - ((total * int(discount.DiscountPercent)) / 100)
	if discountedPrice < 0 {
		discountedPrice = 0
	}

	subtotal := discountedPrice

	// Updating the subtotal in the cart table

	// Updating the subtotal in the cart table

	// Check if any coupon is present inside the cart
	var couponId int
	findCoupon := `SELECT coupon_id FROM carts WHERE user_id=$1`
	tx.Raw(findCoupon, userId).Scan(&couponId)
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	// Apply the coupon to the total if present
	if couponId != 0 {
		// Find the coupon details
		var coupon domain.Coupons
		getCouponDetails := `SELECT * FROM coupons WHERE id=$1`
		err := tx.Raw(getCouponDetails, couponId).Scan(&coupon).Error
		if err != nil {
			tx.Rollback()
			return err
		}

		// Apply the coupon to the total
		discountAmount := (subtotal * int(coupon.DiscountPercent)) / 100
		if discountAmount > int(coupon.DiscountMaximumAmount) {
			discountAmount = int(coupon.DiscountMaximumAmount)
		}

		updateTotal := `UPDATE carts SET total=$1 WHERE id=$2`
		err = tx.Exec(updateTotal, subtotal-discountAmount, cartId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		// If no coupon, update the total without any discount
		updateTotal := `UPDATE carts SET total=$1 WHERE id=$2`
		err = tx.Exec(updateTotal, subtotal, cartId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction
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
