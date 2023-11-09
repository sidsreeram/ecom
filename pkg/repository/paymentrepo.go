package repository

import (
	"fmt"

	"github.com/ECOMMERCE_PROJECT/pkg/domain"
	interfaces "github.com/ECOMMERCE_PROJECT/pkg/repository/interface"
	"gorm.io/gorm"
)

type PaymentDatabase struct {
	DB *gorm.DB
}

func NewPaymentRepository(DB *gorm.DB) interfaces.PaymentRepository {
	return &PaymentDatabase{DB}
}
func (c *PaymentDatabase) ViewPaymentDetails(orderID int) (domain.PaymentDetails, error) {
	var paymentDetails domain.PaymentDetails
	fetchPaymentDetailsQuery := `SELECT * FROM payment_details WHERE orders_id = $1;`
	err := c.DB.Raw(fetchPaymentDetailsQuery, orderID).Scan(&paymentDetails).Error
	fmt.Println("2", paymentDetails)
	return paymentDetails, err
}

func (c *PaymentDatabase) UpdatePaymentDetails(orderID int, paymentRef string) (domain.PaymentDetails, error) {
	var updatedPayment domain.PaymentDetails
	updatePaymentQuery := `	UPDATE payment_details SET payment_type_id = 2, payment_status_id = 1, payment_ref = $1, updated_at = NOW()
							WHERE orders_id = $2 RETURNING *;`
	err := c.DB.Raw(updatePaymentQuery, paymentRef, orderID).Scan(&updatedPayment).Error
	return updatedPayment, err
}
