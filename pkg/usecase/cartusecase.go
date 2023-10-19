package usecase
// import(
	
// )


// type CartUseCase struct {
// 	cartRepo interfaces.CartRepository
// }

// func NewCartUsecase(cartRepo interfaces.CartRepository) services.CartUsecase {
// 	return &CartUseCase{
// 		cartRepo: cartRepo,
// 	}
// }

// func (c *CartUseCase) CreateCart(id int) error {
// 	err := c.cartRepo.CreateCart(id)
// 	return err
// }

// func (c *CartUseCase) AddToCart(productId, userId int) error {
// 	err := c.cartRepo.AddToCart(productId, userId)
// 	return err
// }