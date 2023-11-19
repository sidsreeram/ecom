package usecase

import (
	"log"

	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/domain"
	interfaces "github.com/ECOMMERCE_PROJECT/pkg/repository/interface"
	services "github.com/ECOMMERCE_PROJECT/pkg/usecase/interface"
)

type Orderusecase struct {
	Orderrepo interfaces.OrderRepository
}

func NewOrderUsecase(Orderrepo interfaces.OrderRepository) services.OrderUsercase {
	return &Orderusecase{
		Orderrepo: Orderrepo,
	}
}
func (o *Orderusecase) OrderAll(id, paymentTypeId int) (domain.Orders, error) {
	orders, err := o.Orderrepo.OrderAll(id, paymentTypeId)
	return orders, err
}
func (o *Orderusecase) UserCancelOrder(orderId, userId int) error {
	log.Println(orderId)
	log.Println(userId)
	err := o.Orderrepo.UserCancelOrder(orderId, userId)
	return err
}
func (o *Orderusecase) ReturnOrder(userId, OrderId int) (int, error) {
	ordertotal, err := o.Orderrepo.ReturnOrder(userId, OrderId)
	return ordertotal, err
}
func (o *Orderusecase) ListAorder(userId, orderId int) (domain.Orders, error) {
	order, err := o.Orderrepo.ListAorder(userId, orderId)
	return order, err
}
func (o *Orderusecase) ListAllorder(userId int) (domain.Orders, error) {
	order, err := o.Orderrepo.ListAllorder(userId)
	return order, err
}
func (o *Orderusecase) UpdateOrder(updateorder helperstruct.UpdateOrder) error {
	err := o.Orderrepo.UpdateOrder(updateorder)
	return err
}
func (o *Orderusecase) DownloadInvoice(orderId int) error {
	err := o.Orderrepo.DownloadInvoice(orderId)
	return err
}
