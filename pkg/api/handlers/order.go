package handlers

import (
	"net/http"
	"strconv"

	"github.com/ECOMMERCE_PROJECT/pkg/api/handlerutils"
	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
	services "github.com/ECOMMERCE_PROJECT/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	Orderusecase services.OrderUsercase
}

func NewOrderHandler(orderUsercase services.OrderUsercase) *OrderHandler {
	return &OrderHandler{
		Orderusecase: orderUsercase,
	}
}
func (o *OrderHandler) OrderAll(c *gin.Context) {
	paramsId := c.Param("payment_id")
	paymentTypeId, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
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
	order, err := o.Orderusecase.OrderAll(userId, paymentTypeId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Place Order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 201,
		Message:    "Order Placed",
		Data:       order,
		Errors:     nil,
	})
}
func (o *OrderHandler) UserCancelOrder(c *gin.Context) {
	paramsId := c.Param("order_id")
	orderId, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind failed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
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
	err = o.Orderusecase.UserCancelOrder(orderId, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't cancel Order",
			Data:       nil,
			Errors:     err.Error(),
		})

		return
	}
	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 201,
		Message:    "Order canceld",
		Data:       nil,
		Errors:     nil,
	})
}
func (o *OrderHandler) ReturnOrder(c *gin.Context) {
	paramsId := c.Param("order_id")
	orderId, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind failed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
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
	returntotal, err := o.Orderusecase.ReturnOrder(userId, orderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Return order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, response.Response{
		StatusCode: 202,
		Message:    "Return Accepted",
		Data:       returntotal,
		Errors:     nil,
	})
}
func (o *OrderHandler) ListAorder(c *gin.Context) {
	paramsId := c.Param("order_id")
	orderId, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind failed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
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
	orders, err := o.Orderusecase.ListAorder(userId, orderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't fetch order details",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, response.Response{
		StatusCode: 202,
		Message:    "Order details are",
		Data:       orders,
		Errors:     nil,
	})
}
func (o *OrderHandler) ListAllorder(c *gin.Context) {

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
	orders, err := o.Orderusecase.ListAllorder(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't fetch order details",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, response.Response{
		StatusCode: 202,
		Message:    "Order details are",
		Data:       orders,
		Errors:     nil,
	})
}
func (o *OrderHandler) UpdateOrder(c *gin.Context) {
	var UpdateOrder helperstruct.UpdateOrder
	err := c.BindJSON(&UpdateOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Bind Json Data",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err =o.Orderusecase.UpdateOrder(UpdateOrder)
	if err!=nil{
		c.JSON(http.StatusBadRequest,response.Response{
			StatusCode: 400,
			Message: "Can't update Order",
			Data: nil,
			Errors: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,response.Response{
		StatusCode: 200,
		Message: "order Updated",
		Data: nil,
		Errors: nil,
	})
}
