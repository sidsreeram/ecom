package repository

import (
	"fmt"

	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
	interfaces "github.com/ECOMMERCE_PROJECT/pkg/repository/interface"
	"gorm.io/gorm"
)

type productDatabase struct {
	DB *gorm.DB
}

func NewProductRepostiory(DB *gorm.DB) interfaces.ProductRepository {
	return &productDatabase{DB}
}
func (c *productDatabase) CreateCategory(category helperstruct.Category) (response.Category, error) {
	var NewCategory response.Category
	query := `INSERT INTO categories (category_name,created_at)VALUES($1,NOW())RETURNING 	ID,cATEGORY_name`
	err := c.DB.Raw(query, category.Name).Scan(&NewCategory).Error
	return NewCategory, err
}
func (c *productDatabase) UpdateCategory(category helperstruct.Category, id int) (response.Category, error) {
	var UpdateCategory response.Category
	query := `UPDATE categories SET category_name=$1 ,updated_at=NOW() WHERE EXISTS (SELECT 1 FROM categories WHERE ID =$2)RETURNING id,category_name`
	err := c.DB.Raw(query, category.Name, id).Scan(&UpdateCategory).Error
	if err != nil {
		return response.Category{}, err
	}
	if UpdateCategory.Id == 0 {
		return response.Category{}, fmt.Errorf("No such category")
	}
	return response.Category{}, err
}
func (c *productDatabase) DeleteCategory(id int) error {
	var exists bool
	query1 := `SELECT exists(SELECT 1 FROM categories where id=?)`
	c.DB.Raw(query1, id).Scan(&exists)
	if !exists {
		return fmt.Errorf("There is no such category")
	}
	query := `DELETE FROM categories where id=$1`
	err := c.DB.Exec(query, id).Error
	return err
}
func (c *productDatabase) ListCategories() ([]response.Category, error) {
	var Category []response.Category
	query := `SELECT * FROM categories`
	err := c.DB.Raw(query).Scan(&Category).Error
	return Category, err
}
func (c *productDatabase) DisplayACategory(id int) (response.Category, error) {
	var category response.Category
	var exists bool
	query1 := `SELECT exists(SELECT 1 from categories where id=?)`
	c.DB.Raw(query1, id).Scan(&exists)
	if !exists {
		return response.Category{}, fmt.Errorf("There is no such category")
	}
	query := `SELECT 	* FROM 	categories WHERE  id=$1`
	err := c.DB.Raw(query, id).Scan(&category).Error
	if err != nil {
		return response.Category{}, err
	}
	if category.Id == 0 {
		return response.Category{}, fmt.Errorf("There is no such category")
	}
	return category, nil

}

func (c *productDatabase) AddProduct(product helperstruct.Product) (response.Product, error) {
	var newProduct response.Product
	var exists bool
	query1 := `SELECT exists(SELECT 1 FROM categories WHERE id=?)`
	c.DB.Raw(query1, product.CategoryId).Scan(&exists)

	if !exists {
		return response.Product{}, fmt.Errorf("Category not Found")
	}
	query := `
        INSERT INTO products (product_name, description, brand, category_id, created_at)
        VALUES ($1, $2, $3, $4, NOW())
        RETURNING
            id,
            product_name AS name,
            description,
            brand,
            category_id,
            (SELECT category_name FROM categories WHERE id = $4) AS CategoryName
    `
	err := c.DB.Raw(query, product.Name, product.Description, product.Brand, product.CategoryId).Scan(&newProduct).Error
	return newProduct, err
}

func (c *productDatabase) UpdateProduct(id int, product helperstruct.Product) (response.Product, error) {
	var updateproduct response.Product
	query := `
        UPDATE products
        SET product_name = $1, description = $2, brand = $3, category_id = $4, updated_at = NOW()
        WHERE id = $5
        RETURNING
            id,
            product_name AS name,
            description,
            brand,
            category_id,
            (SELECT category_name FROM categories WHERE id = $4) AS CategoryName
    `
	err := c.DB.Raw(query, product.Name, product.Description, product.Brand, product.CategoryId, id).Scan(&updateproduct).Error
	if err != nil {
		return response.Product{}, err
	}
	if updateproduct.Id == 0 {
		return response.Product{}, fmt.Errorf("There is no such Product")
	}
	return updateproduct, nil
}

func (c *productDatabase) DeleteProduct(id int) error {
	var exists bool
	query1 := `SELECT exists(SELECT 1 FROM products where id=?)`
	c.DB.Raw(query1, id).Scan(&exists)
	if !exists {
		return fmt.Errorf("There is no such product")
	}
	query := `DELETE FROM products where id=$1`
	err := c.DB.Exec(query, id).Error
	return err
}
func (c *productDatabase) AddProductitem(productItem helperstruct.ProductItem) (response.ProductItem, error) {
	var newProductItem response.ProductItem
	query := `INSERT INTO product_items (
		product_id,sku,qty_in_stock,price,in_stock,color,size,rating,created_at)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,NOW()
	RETURNING id,product_id,sku,qty_in_stock,price,in_stock,color,size,rating,created_at )
	`
	err := c.DB.Raw(query, productItem.Product_id, productItem.Sku, productItem.Qty, productItem.Price, productItem.Instock, productItem.Color, productItem.Size, productItem.Rating).Scan(&newProductItem).Error

	return newProductItem, err

}
func (c *productDatabase) UpdateProductItem(id int, product helperstruct.ProductItem) (response.ProductItem, error) {
	var updatedProductItem response.ProductItem
// used product for productitem check that
	query := `
		UPDATE product_items
		SET product_id = $2, sku = $3, qty_in_stock = $4, price = $5, in_stock = $6, color = $7, size = $8, rating = $9
		WHERE id = $1
		RETURNING id, product_id, sku, qty_in_stock, price, in_stock, color, size, rating, created_at
	`

	err := c.DB.Raw(query,id,product.Product_id, product.Sku, product.Qty, product.Price, product.Instock, product.Color, product.Size, product.Rating).Scan(&updatedProductItem).Error

	return updatedProductItem, err
}
func(c*productDatabase) DeleteProductItem(id int)error{
	var exists bool
	query1 := `SELECT exists(SELECT 1 FROM product_items where id=?)`
	c.DB.Raw(query1, id).Scan(&exists)
	if !exists {
		return fmt.Errorf("There is no such product")
	}
	query := `DELETE FROM product_items where id=$1`
	err := c.DB.Exec(query, id).Error
	return err
}
func (c*productDatabase) DisplayAproductitem(id int)(response.ProductItem,error){
	var productItem response.ProductItem
	query:=`SELECT p.product_name,
	p.description,
	p.brand,
	c.category_name, 
	pi.*
	FROM products p 
	JOIN categories c ON p.category_id=c.id 
	JOIN product_items pi ON p.id=pi.product_id 
	WHERE pi.id=$1`
	err := c.DB.Raw(query, id).Scan(&productItem).Error
	if err != nil {
		return response.ProductItem{}, err
	}
	if productItem.Id == 0 {
		return response.ProductItem{}, fmt.Errorf("there is no such product item")
	}
	getImages := `SELECT file_name FROM images WHERE product_item_id=$1`
	err = c.DB.Raw(getImages, id).Scan(&productItem.Image).Error
	if err != nil {
		return response.ProductItem{}, err
	}
	return productItem, nil
}