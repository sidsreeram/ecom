package repository

import (
	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/domain"
	interfaces "github.com/ECOMMERCE_PROJECT/pkg/repository/interface"
	"gorm.io/gorm"
	
)
type DiscountDatabase struct{
	DB *gorm.DB
}
func NewDiscountRepository(DB *gorm.DB)interfaces.DiscountRepository{
	return &DiscountDatabase{DB}
}
func (c*DiscountDatabase) AddDiscount(discount helperstruct.Discount)error{
	add:=`INSERT INTO discounts(category_id,discount_percent,discount_maximum_amount,minimum_purchase_amount,expiration_date) VALUES ($1,$2,$3,$4,$5)`
	err:=c.DB.Exec(add,discount.CategoryId,discount.DiscountPercent,discount.DiscountMaximumAmount,discount.MinimumPurchaseAmount,discount.ExpirationDate).Error
	return err
}
func (c*DiscountDatabase) UpdateDiscount(id int, discount helperstruct.Discount) (domain.Discount,error) {
	var updatediscount domain.Discount
	updatequery:=`UPDATEdiscounts SET category_id=$1,discount_percent=$2,discount_maximum_amount=$3,minimum_purchase_amount=$4,expiration_date=$5 WHERE id=$6 RETURNING *`
	err:= c.DB.Raw(updatequery,discount.CategoryId,discount.DiscountPercent,discount.DiscountMaximumAmount,discount.MinimumPurchaseAmount,discount.ExpirationDate,discount.Id).Scan(&updatediscount).Error
	return updatediscount ,err
}
func(c*DiscountDatabase) DeleteDiscount(id int)error{
	deleteCoupon := `DELETE 	FROM discounts WHERE id=?`
	err := c.DB.Exec(deleteCoupon, id).Error
	return err
}
func (c*DiscountDatabase) GetAllDiscount()([]domain.Discount,error){
	var allDiscount []domain.Discount
	query := `SELECT *FROM discounts `
	err := c.DB.Raw(query).Scan(&allDiscount).Error
	return allDiscount, err
}
func (c*DiscountDatabase) ViewDiscountbyID(id int)(domain.Discount,error){
	var dis domain.Discount
	query:=`SELECT * FROM discounts WHERE id=?`
	err:=c.DB.Raw(query,id).Scan(&dis).Error
	return dis, err
}


