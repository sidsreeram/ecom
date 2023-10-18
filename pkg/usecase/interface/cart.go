package interfaces
import(

)
type CartUseCase interface {
	CreateCart(id int) error
	AddToCart(productId, userId int) error
}