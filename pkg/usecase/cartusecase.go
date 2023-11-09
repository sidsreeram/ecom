package usecase

import (
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
	interfaces "github.com/ECOMMERCE_PROJECT/pkg/repository/interface"
	services "github.com/ECOMMERCE_PROJECT/pkg/usecase/interface"
)

type CartUseCase struct {
	cartRepo interfaces.CartRepository
}

func NewCartUsecase(cartRepo interfaces.CartRepository) services.CartUseCase {
	return &CartUseCase{
		cartRepo: cartRepo,
	}
}

func (c *CartUseCase) CreateCart(id int) error {
	err := c.cartRepo.CreateCart(id)
	return err
}

func (c *CartUseCase) AddToCart(productId, userId int) error {
	err := c.cartRepo.AddToCart(productId, userId)
	return err
}
func (c *CartUseCase) RemoveFromCart(userId, productId int) error {
	err := c.cartRepo.RemoveFromCart(userId, productId)
	return err
}
func (c *CartUseCase) ListCart(userId int) (response.ViewCart, error) {
    cart, err:=c.cartRepo.ListCart(userId)
    return cart,err
}
