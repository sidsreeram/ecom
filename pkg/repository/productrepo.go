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
	query := `INSERT INTO category (category_name,created_at)VALUES($1,NOW())RETURNING 	ID,cATEGORY_name`
	err := c.DB.Raw(query, category.Name).Scan(&NewCategory).Error
	return NewCategory, err
}
func (c *productDatabase) UpdateCategory(category helperstruct.Category, id int) (response.Category, error) {
	var UpdateCategory response.Category
	query := `UPDATE category SET category_name=$1 ,updated_at=NOW() WHERE EXISTS (SELECT 1 FROM category WHERE ID =$2)RETURNING id,category_name`
	err := c.DB.Raw(query, category.Name, id).Scan(&UpdateCategory).Error
	if err != nil {
		return response.Category{}, err
	}
	if UpdateCategory.Id == 0 {
		return response.Category{}, fmt.Errorf("No such category")
	}
	return response.Category{},err
}
func (c*productDatabase)DeleteCategory(id int) error{
	var exists bool 
	query1:=`SELECT exists(SELECT 1 FROM category where id=?)`
	c.DB.Raw(query1,id).Scan(&exists)
	if !exists{
		return fmt.Errorf("There is no such category")
	}
	query:=`DELETE FROM category where id=$1`
  err:= c.DB.Exec(query,id).Error
  return err
}
func (c *productDatabase) ListCategories() ([]response.Category, error) {
	var Category []response.Category
	query := `SELECT * FROM category`
	err := c.DB.Raw(query).Scan(&Category).Error
	return Category, err
}