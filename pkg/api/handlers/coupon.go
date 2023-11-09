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

type CouponHandler struct {
	couponusecase services.Couponusecase
}

func NewCouponHandler(couponusecase services.Couponusecase) *CouponHandler {
	return &CouponHandler{couponusecase: couponusecase}
}
func (cr *CouponHandler) AddCoupon(c *gin.Context) {
	var newcoupon helperstruct.Coupons
	err := c.Bind(&newcoupon)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Bind failed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.couponusecase.AddCoupon(newcoupon)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Failed to add coupon",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, response.Response{
		StatusCode: 202,
		Message:    "Add coupon success",
		Data:       nil,
		Errors:     nil,
	})
}
func (cr *CouponHandler) UpdateCoupon(c *gin.Context) {
	id := c.Param("couponId")
	couponId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	var newCoupon helperstruct.Coupons
	err = c.Bind(&newCoupon)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	updatedCoupon, err := cr.couponusecase.UpdateCoupon(newCoupon, couponId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't update coupon details",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "coupon updated",
		Data:       updatedCoupon,
		Errors:     nil,
	})
}

func (cr *CouponHandler) DeleteCoupon(c *gin.Context) {
	id := c.Param("couponId")
	couponId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.couponusecase.DeleteCoupon(couponId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Failed to delete coupon",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, response.Response{
		StatusCode: 202,
		Message:    "coupon deleted",
		Data:       nil,
		Errors:     nil,
	})
}
func (cr *CouponHandler) ViewCoupons(c *gin.Context) {
	coupons, err := cr.couponusecase.ViewAllCoupons()

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't finds coupon",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "coupons are ",
		Data:       coupons,
		Errors:     nil,
	})
}
func (cr *CouponHandler) ViewAcoupon(c *gin.Context) {
	id := c.Param("couponId")
	couponId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	coupondetails, err := cr.couponusecase.ViewCoupon(couponId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "coupons are ",
		Data:       coupondetails,
		Errors:     nil,
	})
}
func (cr *CouponHandler) ApplyCoupon(c *gin.Context) {
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
	couponCode := c.Param("code")

	discountRate, err := cr.couponusecase.ApplyCoupon(userId, couponCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't apply coupon",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 200,
			Message:    "coupon applied",
			Data:       []interface{}{"rate after coupon applied is ", discountRate},
			Errors:     nil,
		})
	}
}

func (cr *CouponHandler) RemoveCoupon(c *gin.Context) {
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
	err = cr.couponusecase.RemoveCoupon(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't remove coupon",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Coupon removed",
		Data:       nil,
		Errors:     nil,
	})
}
