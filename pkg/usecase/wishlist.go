package usecase

import (
	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	interfaces "github.com/ECOMMERCE_PROJECT/pkg/repository/interface"
	services "github.com/ECOMMERCE_PROJECT/pkg/usecase/interface"
)

type WishlistUsecase struct {
	wishlistrepository interfaces.WishlistRepository
}

func NewWishlistUsecase(wishlistrepo interfaces.WishlistRepository) services.WishlistUsecase {
	return &WishlistUsecase{wishlistrepository: wishlistrepo}
}
func (w *WishlistUsecase) AddToWishlist(userID, productID int) error {
	err := w.wishlistrepository.AddToWishlist(userID, productID)
	return err
}
func (w *WishlistUsecase) RemoveFromWishlist(userId, productId int) error {
	err := w.wishlistrepository.RemoveFromWishlist(userId, productId)
	return err
}
func (w *WishlistUsecase) ViewAllWishlistItems(userId int) ([]helperstruct.ProductItem, error){
	wishlist,err:=w.wishlistrepository.ViewAllWishlistItems(userId)
	return wishlist,err
}
