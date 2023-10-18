package handlers

import (
	"net/http"
	"strconv"

	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
	services "github.com/ECOMMERCE_PROJECT/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductUsecase services.ProductUsecase
}

func NewProductHandler(Productusecase services.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		ProductUsecase: Productusecase,
	}

}
func (cr *ProductHandler) CreateCategory(c *gin.Context) {
	var category helperstruct.Category
	err := c.Bind(&category)
	if err != nil {
		c.JSON(http.StatusBadGateway, response.Response{
			StatusCode: 400,
			Message:    "Can't Bind ",
			Data:       nil,
			Errors:     err.Error(),
		})
	}
	Newcategory, err := cr.ProductUsecase.CreateCategory(category)
	if err != nil {
		c.JSON(http.StatusBadGateway, response.Response{
			StatusCode: 400,
			Message:    "Can't Create a New Category",
			Data:       nil,
			Errors:     err,
		})
		c.JSON(http.StatusCreated, response.Response{
			StatusCode: 201,
			Message:    "Category Created Successfully",
			Data:       Newcategory,
			Errors:     nil,
		})
	}

}
func (cr *ProductHandler) UpdateCategory(c *gin.Context) {
	var category helperstruct.Category
	err := c.Bind(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Bind",
			Data:       nil,
			Errors:     err.Error(),
		})
	}
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	updatecategory, err := cr.ProductUsecase.UpdateCategory(category, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Failed to Update category",
			Data:       nil,
			Errors:     err.Error(),
		})
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Category Updated Successfully",
		Data:       updatecategory,
		Errors:     nil,
	})
}
func (cr *ProductHandler) DeleteCategory(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err!=nil{
		c.JSON(http.StatusBadRequest,response.Response{
			StatusCode: 400,
			Message: "Can't Bind",
			Data: nil,
			Errors: err.Error(),
		})
	}
	err = cr.ProductUsecase.DeleteCategory(id)
	if err!=nil{
		c.JSON(http.StatusBadRequest,response.Response{
			StatusCode: 400,
			Message: "Failed To Delete",
			Data: nil,
			Errors: err,
		})
	}
	c.JSON(http.StatusOK,response.Response{
		StatusCode: 200,
		Message: "Category Deleted",
		Data: nil,
		Errors: nil,
	})
}
func (cr*ProductHandler)ListCategories(c *gin.Context){
	category,err:=cr.ProductUsecase.ListCategories()
	if err!=nil{
		c.JSON(http.StatusBadRequest,response.Response{
			StatusCode: 400,
			Message: "Can't Load categories",
			Data: nil,
			Errors: err.Error(),
		})
	}
	c.JSON(http.StatusOK,response.Response{
		StatusCode: 200,
		Message: "Categories are ",
		Data: category,
		Errors: nil,
	})
}

