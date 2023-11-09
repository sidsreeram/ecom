package usecase

import (
	"fmt"

	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/config"
	"github.com/ECOMMERCE_PROJECT/pkg/domain"
	interfaces "github.com/ECOMMERCE_PROJECT/pkg/repository/interface"
	services "github.com/ECOMMERCE_PROJECT/pkg/usecase/interface"
	"github.com/razorpay/razorpay-go"
)

type PaymentUsecase struct {
	paymentrepo interfaces.PaymentRepository
	orderrepo   interfaces.OrderRepository
	cfg         config.Config
}

func NewPaymentuseCase(paymentrepo interfaces.PaymentRepository, orderrepo interfaces.OrderRepository, cfg config.Config) services.PaymentUsecase {
	return &PaymentUsecase{
		paymentrepo: paymentrepo,
		orderrepo:   orderrepo,
		cfg:         cfg,
	}
}

func (c *PaymentUsecase) CreateRazorpayPayment(userId, orderId int) (domain.Orders, string, error) {
	paymentDetails, err := c.paymentrepo.ViewPaymentDetails(orderId)
	if err != nil {
		return domain.Orders{}, "", err
	}
	if paymentDetails.PaymentStatusID == 3 {
		return domain.Orders{}, "", fmt.Errorf("payment already completed")
	}
	//fetch order details from the db
	order, err := c.orderrepo.ListAorder(userId, orderId)
	if err != nil {
		return domain.Orders{}, "", err
	}
	if order.Id == 0 {
		return domain.Orders{}, "", fmt.Errorf("no such order found")
	}
	client := razorpay.NewClient(c.cfg.RAZORPAYID, c.cfg.RAZORPAYSECRET)

	data := map[string]interface{}{
		"amount":   order.OrderTotal * 100,
		"currency": "INR",
		"receipt":  "test_receipt_id",
	}

	body, err := client.Order.Create(data, nil)
	if err != nil {
		return domain.Orders{}, "", err
	}

	value := body["id"]
	razorpayID := value.(string)
	return order, razorpayID, err
}

func (c *PaymentUsecase) UpdatePaymentDetails(paymentVerifier helperstruct.PaymentVerification) error {
	paymentDetails, err := c.paymentrepo.ViewPaymentDetails(paymentVerifier.OrderID)
	if err != nil {
		return err
	}
	if paymentDetails.ID == 0 {
		return fmt.Errorf("no order found")
	}

	if paymentDetails.OrderTotal != paymentVerifier.Total {
		return fmt.Errorf("payment amount and order amount does not match")
	}
	updatedPayment, err := c.paymentrepo.UpdatePaymentDetails(paymentVerifier.OrderID, paymentVerifier.PaymentRef)
	if err != nil {
		return err
	}
	if updatedPayment.ID == 0 {
		return fmt.Errorf("failed to update payment details")
	}
	return nil
}
