package interfaces

import "github.com/ECOMMERCE_PROJECT/pkg/common/response"

type CartRepository interface {
	CreateCart(id int) error
	AddToCart(productId, userId int) error
	RemoveFromCart(userId, productId int) error
	ListCart(userId int) (response.ViewCart, error)
}
