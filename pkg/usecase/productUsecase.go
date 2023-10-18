package usecase

import (
	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
	interfaces "github.com/ECOMMERCE_PROJECT/pkg/repository/interface"
	services "github.com/ECOMMERCE_PROJECT/pkg/usecase/interface"
)

type ProductUsecase struct {
	productrepo interfaces.ProductRepository
}

func NewProductUsecase(productRepo interfaces.ProductRepository)services.ProductUsecase{
 return &ProductUsecase{
	productrepo:productRepo,
 }
}
func (c*ProductUsecase)CreateCategory(category helperstruct.Category)(response.Category,error){
	newCatergory ,err:=c.productrepo.CreateCategory(category)
	return newCatergory,err
}
func (c*ProductUsecase)UpdateCategory(category helperstruct.Category,id int)(response.Category,error){
	updateCategory,err:=c.productrepo.UpdateCategory(category,id)
	return updateCategory,err
}
func (c*ProductUsecase)DeleteCategory(id int)error{
	err:=c.productrepo.DeleteCategory(id)
	return err
}
func (c*ProductUsecase)ListCategories()([]response.Category,error){
	categories,err:=c.productrepo.ListCategories()
	return categories,err
}
