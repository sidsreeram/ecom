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
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Bind ",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	Newcategory, err := cr.ProductUsecase.CreateCategory(category)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Create a New Category",
			Data:       nil,
			Errors:     err,
		})
		return
	}
	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 201,
		Message:    "Category Created Successfully",
		Data:       Newcategory,
		Errors:     nil,
	})

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
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Bind",
			Data:       nil,
			Errors:     err.Error(),
		})
	}
	err = cr.ProductUsecase.DeleteCategory(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Failed To Delete",
			Data:       nil,
			Errors:     err,
		})
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Category Deleted",
		Data:       nil,
		Errors:     nil,
	})
}
func (cr *ProductHandler) ListCategories(c *gin.Context) {
	category, err := cr.ProductUsecase.ListCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Load categories",
			Data:       nil,
			Errors:     err.Error(),
		})
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Categories are ",
		Data:       category,
		Errors:     nil,
	})
}
func (cr*ProductHandler)DisplayACategory(c*gin.Context){
	paramsId := c.Param("id")
	id,err:=strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Bind",
			Data:       nil,
			Errors:     err.Error(),
		})
	}
	category,err := cr.ProductUsecase.DisplayACategory(id)
	if err!=nil{
		c.JSON(http.StatusBadRequest,response.Response{
			StatusCode: 400,
			Message: "Can't fetch the data",
			Data: nil,
			Errors: err.Error(),
		})
	}
	c.JSON(http.StatusOK,response.Response{
		StatusCode: 200,
		Message: "Data fetched successfully",
		 Data: category,
		 Errors: nil,
	})
}

func (cr *ProductHandler)AddProduct(c *gin.Context) {
	var product helperstruct.Product
	err := c.Bind(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Bind ",
			Data:       nil,
			Errors:     err.Error(),
		})
	}
	Newproduct, err := cr.ProductUsecase.AddProduct(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Add new Product",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 201,
		Message:    "New Product added Successfully",
		Data:       Newproduct,
		Errors:     nil,
	})

}
func (cr *ProductHandler) UpdateProduct(c *gin.Context) {
	var Product helperstruct.Product
	err := c.Bind(&Product)
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
	updateProduct, err := cr.ProductUsecase.UpdateProduct(id,Product)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Failed to Update Porduct",
			Data:       nil,
			Errors:     err.Error(),
		})
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Product Updated Successfully",
		Data:       updateProduct,
		Errors:     nil,
	})
}
func (cr*ProductHandler)DeleteProduct(c *gin.Context){
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Bind",
			Data:       nil,
			Errors:     err.Error(),
		})
	}
	err = cr.ProductUsecase.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Failed To Delete",
			Data:       nil,
			Errors:     err,
		})
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Product Deleted",
		Data:       nil,
		Errors:     nil,
	})
}
//-----------------------------------------------------ProductItem------------------------------------------------------------
func (cr *ProductHandler)AddProductitem(c *gin.Context) {
	var productItem helperstruct.ProductItem
	err := c.Bind(&productItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Bind ",
			Data:       nil,
			Errors:     err.Error(),
		})
	}
	newProductItem, err := cr.ProductUsecase.AddProductitem(productItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Add new Product-item",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 201,
		Message:    "New Product-item added Successfully",
		Data:       newProductItem,
		Errors:     nil,
	})

}
//----------------------------------------------------update ProductItem-------------------------------------------------------------


func (cr *ProductHandler) UpdateProductitem(c *gin.Context) {
	var Productitem helperstruct.ProductItem
	err := c.Bind(&Productitem)
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
	updateProduct, err := cr.ProductUsecase.UpdateProductItem(id,Productitem)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Failed to Update Porduct",
			Data:       nil,
			Errors:     err.Error(),
		})
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Product Updated Successfully",
		Data:       updateProduct,
		Errors:     nil,
	})
}
//--------------------------------------Delete ProductItem---------------------------------------------------------------
func (cr*ProductHandler)DeleteProductItem(c *gin.Context){
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Bind",
			Data:       nil,
			Errors:     err.Error(),
		})
	}
	err = cr.ProductUsecase.DeleteProductItem(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Failed To Delete",
			Data:       nil,
			Errors:     err,
		})
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Product Deleted",
		Data:       nil,
		Errors:     nil,
	})
}
