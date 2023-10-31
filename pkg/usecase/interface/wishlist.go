package interfaces

import "github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"

type WishlistUsecase interface {
	AddToWishlist(userID, productID int) error
	RemoveFromWishlist(userId,productId int)error
	ViewAllWishlistItems(userId int) ([]helperstruct.ProductItem, error)
}
