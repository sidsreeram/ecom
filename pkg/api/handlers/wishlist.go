package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ECOMMERCE_PROJECT/pkg/api/handlerutils"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
	services "github.com/ECOMMERCE_PROJECT/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type WishlistHandler struct {
	wishlistUsecase services.WishlistUsecase
}

func NewWishlistHandler(wishlistUsecase services.WishlistUsecase) *WishlistHandler {
	return &WishlistHandler{wishlistUsecase: wishlistUsecase}
}
func (w *WishlistHandler) AddToWishlist(c *gin.Context) {
	log.Println("hiii")
	userId, err := handlerutils.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find Id",
			Data:       nil,
			Errors:     err.Error(),
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
	log.Println("hii")
	err = w.wishlistUsecase.AddToWishlist(userId, productId)
	if err!=nil{
		c.JSON(http.StatusBadRequest,response.Response{
			StatusCode: 400,
			Message: "Failed to add ",
			Data: nil,
			Errors: err.Error(),
		})
		return
	}
    c.JSON(http.StatusOK,response.Response{
		StatusCode: 200,
		Message: "item added to wishlist successfully",
		Data: nil,
		Errors: nil,
	})
}
func (w *WishlistHandler)RemoveFromWishlist(c *gin.Context) {
	
	userId, err := handlerutils.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find Id",
			Data:       nil,
			Errors:     err.Error(),
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
	err = w.wishlistUsecase.RemoveFromWishlist(userId, productId)
	if err!=nil{
		c.JSON(http.StatusBadRequest,response.Response{
			StatusCode: 400,
			Message: "Failed to remove",
			Data: nil,
			Errors: err.Error(),
		})
		return
	}
    c.JSON(http.StatusOK,response.Response{
		StatusCode: 200,
		Message: "item removed from wishlist successfully",
		Data: nil,
		Errors: nil,
	})
}
func (w *WishlistHandler)ViewAllWishlistItems(c *gin.Context) {
	
	userId, err := handlerutils.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find Id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	wishlist,err := w.wishlistUsecase.ViewAllWishlistItems(userId)
	if err!=nil{
		c.JSON(http.StatusBadRequest,response.Response{
			StatusCode: 400,
			Message: "Failed to fetch wishlist items",
			Data: nil,
			Errors: err.Error(),
		})
		return
	}
    c.JSON(http.StatusOK,response.Response{
		StatusCode: 200,
		Message: "items in wishlist ",
		Data: wishlist,
		Errors: nil,
	})

}