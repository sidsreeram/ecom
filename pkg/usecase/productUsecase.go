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

func NewProductUsecase(productRepo interfaces.ProductRepository) services.ProductUsecase {
	return &ProductUsecase{
		productrepo: productRepo,
	}
}
func (c *ProductUsecase) CreateCategory(category helperstruct.Category) (response.Category, error) {
	newCatergory, err := c.productrepo.CreateCategory(category)
	return newCatergory, err
}
func (c *ProductUsecase) UpdateCategory(category helperstruct.Category, id int) (response.Category, error) {
	updateCategory, err := c.productrepo.UpdateCategory(category, id)
	return updateCategory, err
}
func (c *ProductUsecase) DeleteCategory(id int) error {
	err := c.productrepo.DeleteCategory(id)
	return err
}
func (c *ProductUsecase) ListCategories() ([]response.Category, error) {
	categories, err := c.productrepo.ListCategories()
	return categories, err
}
func (c *ProductUsecase) DisplayACategory(id int) (response.Category, error) {
	categories, err := c.productrepo.DisplayACategory(id)
	return categories, err
}
func (c *ProductUsecase) AddProduct(product helperstruct.Product) (response.Product, error) {
	newProduct, err := c.productrepo.AddProduct(product)
	return newProduct, err
}
func (c *ProductUsecase) UpdateProduct(id int, Product helperstruct.Product) (response.Product, error) {
	updateproduct, err := c.productrepo.UpdateProduct(id, Product)
	return updateproduct, err
}
func (c *ProductUsecase) DeleteProduct(id int) error {
	err := c.productrepo.DeleteProduct(id)
	return err
}
func (c *ProductUsecase) ListAllProduct(viewproduct helperstruct.QueryParams) ([]response.Product, error) {
	products, err := c.productrepo.ListAllProduct(viewproduct)
	return products, err
}
func (c*ProductUsecase) DisplayAProduct(id int) (response.Product, error) {
	product,err:=c.productrepo.DisplayAProduct(id)
	return product,err
}
func (c *ProductUsecase) AddProductitem(productItem helperstruct.ProductItem) (response.ProductItem, error) {
	newProductItem, err := c.productrepo.AddProductitem(productItem)
	return newProductItem, err
}
func (c *ProductUsecase) UpdateProductItem(id int, productItem helperstruct.ProductItem) (response.ProductItem, error) {
	updateProductItem, err := c.productrepo.UpdateProductItem(id, productItem)
	return updateProductItem, err
}
func (c *ProductUsecase) DeleteProductItem(id int) error {
	err := c.productrepo.DeleteProductItem(id)
	return err
}
func (c *ProductUsecase) DisplayAproductitem(id int) (response.ProductItem, error) {
	ProductItem, err := c.productrepo.DisplayAproductitem(id)
	return ProductItem, err
}
func (c*ProductUsecase) DisaplyaAllProductItems(viewProductitem helperstruct.QueryParams) ([]response.ProductItem, error){
	productitem,err:=c.productrepo.DisaplyaAllProductItems(viewProductitem)
	return productitem,err
}
func (c *ProductUsecase) UploadImage(filepath string, productId int) error {
	err := c.productrepo.UploadImage(filepath, productId)
	return err
}
func (c *ProductUsecase) UploadImageBinary(data []byte, filepath string, productId int) error {
    err := c.productrepo.UploadImageBinary(data, filepath, productId)
    return err
}
func (c *ProductUsecase) GetProductImages(productId int) ([]response.ProductImage,error){
	data,err:= c.productrepo.GetProductImages(productId)
	return data , err
}

