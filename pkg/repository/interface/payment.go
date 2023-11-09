package interfaces

import "github.com/ECOMMERCE_PROJECT/pkg/domain"

type PaymentRepository interface {
	ViewPaymentDetails(orderID int) (domain.PaymentDetails, error)
	UpdatePaymentDetails(orderID int, paymentRef string) (domain.PaymentDetails, error)
}
