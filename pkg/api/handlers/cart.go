package handlers

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/ECOMMERCE_PROJECT/pkg/api/handlerutils"
// 	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
// 	"github.com/ECOMMERCE_PROJECT/pkg/usecase"
// 	"github.com/gin-gonic/gin"
// )

// type CartHandler struct {
// 	cartusecase usecase.CartUsecase
// }

// func NewCartHandler(cartusecase serivces.CartUsecase) {
// 	return &CartHandler{
// 		cartusecase: cartusecase,
// 	}

// }
// func (cr *CartHandler) AddToCART(c *gin.Context) {
// 	userId, err := handlerutils.GetUserIdFromContext(c)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, response.Response{
// 			StatusCode: 400,
// 			Message:    "Can't Find the user",
// 			Data:       nil,
// 			Errors:     nil,
// 		})
// 		return
// 	}
// 	paramID := c.Param("Product_item_id")
// 	productId, err := strconv.Atoi(paramID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, response.Response{
// 			StatusCode: 400,
// 			Message:    "cant find productid",
// 			Data:       nil,
// 			Errors:     err.Error(),
// 		})
// 		return
// 	}
// 	err := cr.cartusecase.AddToCART(productId, userId)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, response.Response{
// 			StatusCode: 400,
// 			Message:    "cant add product into cart",
// 			Data:       nil,
// 			Errors:     err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, response.Response{
// 		StatusCode: 200,
// 		Message:    "product added into cart",
// 		Data:       nil,
// 		Errors:     nil,
// 	})

// }
