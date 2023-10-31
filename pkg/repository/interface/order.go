package interfaces

import (
	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/domain"
)

type OrderRepository interface {
	OrderAll(id, paymentTypeId int) (domain.Orders, error)
	UserCancelOrder(orderId, userId int) error
	ListAorder(userId, orderId int) (domain.Orders, error)
	ListAllorder(userId int) (domain.Orders, error)
	ReturnOrder(userId, OrderId int) (int, error)
	UpdateOrder(updateorder helperstruct.UpdateOrder) error
}
