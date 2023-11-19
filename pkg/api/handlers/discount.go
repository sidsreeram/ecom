package handlers

import (
	"net/http"
	"strconv"

	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
	services "github.com/ECOMMERCE_PROJECT/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type DiscountHandler struct {
	discountusecase services.DiscountUsecase
}

func NewDiscountHandler(discountusecase services.DiscountUsecase) *DiscountHandler {
	return &DiscountHandler{discountusecase: discountusecase}
}
func (cr *DiscountHandler) AddDiscount(c *gin.Context) {
	var newdiscount helperstruct.Discount
	err := c.Bind(&newdiscount)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Bind failed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.discountusecase.AddDiscount(newdiscount)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Failed to add discount",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, response.Response{
		StatusCode: 202,
		Message:    "Add discount success",
		Data:       nil,
		Errors:     nil,
	})
}
func (cr *DiscountHandler) UpdateDiscount(c *gin.Context) {
	id := c.Param("discountId")
	discountId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	var newdiscount helperstruct.Discount
	err = c.Bind(&newdiscount)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Bind failed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	updatedis,err := cr.discountusecase.UpdateDiscount(discountId, newdiscount)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Failed to update discount",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, response.Response{
		StatusCode: 202,
		Message:    "update discount success",
		Data:        updatedis,
		Errors:     nil,
	})
}
func (cr *DiscountHandler) DeleteDiscount(c *gin.Context) {
	id := c.Param("discountId")
	discountId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.discountusecase.DeleteDiscount(discountId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Failed to delete discount",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, response.Response{
		StatusCode: 202,
		Message:    "discount deleted",
		Data:       nil,
		Errors:     nil,
	})
}
func (cr *DiscountHandler) GetAllDiscount(c *gin.Context) {
	discount, err := cr.discountusecase.GetAllDiscount()

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find discount",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "all disounts are ",
		Data:       discount,
		Errors:     nil,
	})
}
func (cr *DiscountHandler) ViewDiscountbyID(c *gin.Context) {
	id := c.Param("discountId")
	discountId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	discount, err := cr.discountusecase.ViewDiscountbyID(discountId)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find discount",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "all disounts are ",
		Data:       discount,
		Errors:     nil,
	})
}
