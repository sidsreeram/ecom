package handlers

import (
	"net/http"
	"strconv"

	"github.com/ECOMMERCE_PROJECT/pkg/api/handlerutils"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
	services "github.com/ECOMMERCE_PROJECT/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartusecase services.CartUseCase
}

func NewCartHandler(cartusecase services.CartUseCase) *CartHandler {
	return &CartHandler{
		cartusecase: cartusecase,
	}
}
func (cr *CartHandler) AddToCart(c *gin.Context) {
	userId, err := handlerutils.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Find the user",
			Data:       nil,
			Errors:     nil,
		})
		return
	}
	paramID := c.Param("product_item_id")
	if paramID == "" {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Product ID is missing",
			Data:       nil,
			Errors:     nil,
		})
		return
	}
	productId, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Invalid product ID",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.cartusecase.AddToCart(productId, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Failed to add product into cart",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Product added to cart",
		Data:       nil,
		Errors:     nil,
	})
}

func (cr *CartHandler) RemoveFromCart(c *gin.Context) {
	userId, err := handlerutils.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Find the user",
			Data:       nil,
			Errors:     nil,
		})
		return
	}
	paramID := c.Param("product_item_id")
	productId, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find productid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.cartusecase.RemoveFromCart(userId, productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant remove product from cart",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "product removed from cart",
		Data:       nil,
		Errors:     nil,
	})

}
func (cr *CartHandler) ListCart(c *gin.Context) {
	userId, err := handlerutils.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Find the user",
			Data:       nil,
			Errors:     nil,
		})
		return
	}
	listcart, err := cr.cartusecase.ListCart(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find cart",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "cart items are",
		Data:       listcart,
		Errors:     nil,
	})
}
