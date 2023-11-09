package repository

import (
	"fmt"

	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/domain"
	interfaces "github.com/ECOMMERCE_PROJECT/pkg/repository/interface"
	"gorm.io/gorm"
)

type OrderDatabase struct {
	DB *gorm.DB
}

func NewOrderRepository(DB *gorm.DB) interfaces.OrderRepository {
	return &OrderDatabase{DB}
}
func (c *OrderDatabase) OrderAll(id, paymentTypeId int) (domain.Orders, error) {
	tx := c.DB.Begin()

	var cart domain.Carts
	findcart := `SELECT * FROM carts WHERE user_id=?`
	err := c.DB.Raw(findcart, id).Scan(&cart).Error
	if err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}
	if cart.Total == 0 {
		setTotal := `UPDATE carts SET total = carts.sub_total`
		err = c.DB.Exec(setTotal).Error
		if err != nil {
			tx.Rollback()
			return domain.Orders{}, err
		}
		cart.Total = cart.SubTotal
		if cart.SubTotal == 0 {
			tx.Rollback()
			return domain.Orders{}, fmt.Errorf("NO items in Cart")
		}

	}
	// adding address to the order
	var addressId int
	address := `SELECT id FROM addresses WHERE users_id=$1 AND is_default=true`
	err = tx.Raw(address, id).Scan(&addressId).Error
	if err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}
	if addressId == 0 {
		tx.Rollback()
		return domain.Orders{}, fmt.Errorf("Please add address")
	}
	// setting order id for the order
	var order domain.Orders
	setorder := `
    INSERT INTO orders(user_id, order_date, payment_type_id, shipping_address, order_total, order_status_id, coupon_id)
	VALUES($1, NOW(), $2, $3, $4, 1,$5)
	RETURNING *
               `
	err = tx.Raw(setorder, id, paymentTypeId, addressId, cart.Total,cart.CouponId).Scan(&order).Error
	if err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}

	// getting cart details of the user
	var cartitems []helperstruct.CartItems
	cartdetails := `select ci.product_item_id, ci.quantity, pi.price, pi.qty_in_stock
	  from cart_items ci
	  join product_items pi on ci.product_item_id = pi.id
	  where ci.carts_id = $1
	  `
	err = tx.Raw(cartdetails, cart.Id).Scan(&cartitems).Error
	if err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}
	for _, items := range cartitems {
		if items.Quantity > items.QtyInStock {
			return domain.Orders{}, fmt.Errorf("Item Out OF Stock")
		}
		insetOrderItems := `INSERT INTO order_items (orders_id,product_item_id,quantity,price) VALUES($1,$2,$3,$4)`
		err = tx.Exec(insetOrderItems, order.Id, items.ProductItemId, items.Quantity, items.Price).Error

		if err != nil {
			tx.Rollback()
			return domain.Orders{}, err
		}
	}
	updateCart := `UPDATE carts SET total=0,sub_total=0 WHERE user_id=?`
	err = tx.Exec(updateCart, id).Error
	if err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}
	for _, items := range cartitems {
		removeCartItems := `DELETE FROM cart_items WHERE carts_id =$1 AND product_item_id=$2`
		err = tx.Exec(removeCartItems, cart.Id, items.ProductItemId).Error
		if err != nil {
			tx.Rollback()
			return domain.Orders{}, err
		}
	}
	for _, items := range cartitems {
		updateQty := `UPDATE product_items SET qty_in_stock=product_items.qty_in_stock-$1 WHERE id=$2`
		err = tx.Exec(updateQty, items.Quantity, items.ProductItemId).Error
		if err != nil {
			tx.Rollback()
			return domain.Orders{}, err
		}
	}
	createPaymentDetails := `INSERT INTO payment_details
			(orders_id,
			order_total,
			payment_type_id,
			payment_status_id,
			updated_at)
			VALUES($1,$2,$3,$4,NOW())`
	if err = tx.Exec(createPaymentDetails, order.Id, order.OrderTotal, paymentTypeId, 1).Error; err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}
	return order, nil
}

func (c *OrderDatabase) UserCancelOrder(orderId, userId int) error {
	tx := c.DB.Begin()
	var items []helperstruct.CartItems
	finditem := `SELECT product_item_id,quantity FROM order_items WHERE orders_id =?`
	err := tx.Raw(finditem, orderId).Scan(&items).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if len(items) == 0 {
		return fmt.Errorf("No order found with this ID")
	}
	for _, item := range items {
		updateProductitem := `UPDATE product_items SET qty_in_stock=qty_in_stock+$1 WHERE id=$2`
		err := tx.Raw(updateProductitem, item.Quantity, item.ProductItemId).Error
		if err != nil {
			tx.Rollback()
			return err
		}

	}
	removeItems := `DELETE FROM order_items WHERE orders_id=$1`
	err = tx.Exec(removeItems, orderId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	cancelOrder := `UPDATE orders SET order_status_id=$1 WHERE id=$2 AND user_id=$3`
	err = tx.Exec(cancelOrder, 2, orderId, userId).Error
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
func (c *OrderDatabase) ListAorder(userId, orderId int) (domain.Orders, error) {
	var order domain.Orders
	query := `SELECT * FROM orders WHERE user_id=$1 AND id=$2`
	err := c.DB.Raw(query, userId, orderId).Scan(&order).Error
	return order, err
}
func (c *OrderDatabase) ListAllorder(userId int) (domain.Orders, error) {
	var order domain.Orders
	query := `SELECT * FROM orders WHERE user_id=?`
	err := c.DB.Raw(query, userId).Scan(&order).Error
	return order, err
}
func (c *OrderDatabase) ReturnOrder(userId, OrderId int) (int, error) {
	var order domain.Orders
	getorderdetails := `SELECT *from orders WHERE user_id=$1 AND id=$2`
	err := c.DB.Raw(getorderdetails, userId, OrderId).Scan(&order).Error
	if err != nil {
		return 0, err
	}
	if order.OrderStatusID != 3 {
		return 0, fmt.Errorf("the order is not deleverd")
	}
	returnOder := `UPDATE orders SET order_status_id=$1 WHERE id=$2`
	err = c.DB.Exec(returnOder, 5, OrderId).Error
	if err != nil {
		return 0, err
	}
	return order.OrderTotal, nil
}
func (c *OrderDatabase) UpdateOrder(updateorder helperstruct.UpdateOrder) error {
	var isexists bool
	query := `SELECT EXISTS (SELECT 1 FROM orders WHERE id=?)`
	err := c.DB.Raw(query, updateorder.OrderId).Scan(&isexists).Error
	if err != nil {
		return err
	}
	if !isexists {
		return fmt.Errorf("order not found with this id")
	}
	updateOrderQry := `UPDATE orders SET order_status_id=$1 WHERE id=$2`
	err = c.DB.Exec(updateOrderQry, updateorder.OrderStatusID, updateorder.OrderId).Error
	if err != nil {
		return err
	}
	return nil
}
