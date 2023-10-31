package repository

import (
	"fmt"

	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	interfaces "github.com/ECOMMERCE_PROJECT/pkg/repository/interface"
	"gorm.io/gorm"
)

type WishlistDatabase struct {
	DB *gorm.DB
}

func NewWishlistRepository(DB *gorm.DB) interfaces.WishlistRepository {
	return &WishlistDatabase{DB}
}
func (c *WishlistDatabase) AddToWishlist(userID, productID int) error {
	tx := c.DB.Begin()
	var isPresent bool
	query := `SELECT EXISTS (SELECT 1 FROM wishlists WHERE user_id=$1 AND item_id=$2)`
	err:=tx.Raw(query,userID,productID).Scan(&isPresent).Error
	if err!=nil{
		tx.Rollback()
		return err
	}
    if isPresent {
		return fmt.Errorf("The Item is already Presented in your wishlist")
	}
	insertquery:=`INSERT INTO wishlists (user_id,item_id)VALUES($1,$2)`
	err = tx.Exec(insertquery,userID,productID).Error
	if err!=nil{
		tx.Rollback()
		return err
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
func (c*WishlistDatabase)RemoveFromWishlist(userId,productId int)error{
	tx := c.DB.Begin()
	var isPresent bool
	query := `SELECT EXISTS (SELECT 1 FROM wishlists WHERE user_id=$1 AND item_id=$2)`
	err:=tx.Raw(query,userId,productId).Scan(&isPresent).Error
	if err!=nil{
		tx.Rollback()
		return err
	}
	if !isPresent {
		return fmt.Errorf("Item is not present in your wishlist")
     }
     remove:=`DELETE FROM wishlists WHERE user_id=$1 AND item_id=$2`
	 err =tx.Exec(remove,userId,productId).Error
	 if err!=nil{
		tx.Rollback()
		return err
	 }
	 if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
func (c*WishlistDatabase)ViewAllWishlistItems(userId int)([]helperstruct.ProductItem,error){
	var wishlist []helperstruct.ProductItem
	viewwishlist:=`SELECT p.product_name,
	p.description,
	p.brand,
	c.category_name, 
	pi.*
	FROM wishlists w 
	JOIN product_items pi ON w.item_id = pi.id 
	JOIN products p ON pi.product_id = p.id 
	JOIN categories c ON p.category_id = c.id WHERE w.user_id=$1`
	err:=c.DB.Raw(viewwishlist,userId).Scan(&wishlist).Error
	return wishlist,err
	
}