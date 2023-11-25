package interfaces

import (
	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/domain"
)

type PaymentUsecase interface {
	CreateRazorpayPayment(userId, orderId int) (domain.Orders, string, error)
	UpdatePaymentDetails(paymentVerifier helperstruct.PaymentVerification) error
	GetUserIdFromOrder(orderId int)(int, error)
}
