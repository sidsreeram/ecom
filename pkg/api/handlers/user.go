package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ECOMMERCE_PROJECT/pkg/api/handlerutils"
	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
	services "github.com/ECOMMERCE_PROJECT/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase services.UserUseCase
}

func NewUserHandelr(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUsecase: usecase,
	}
}

func (u *UserHandler) UserSignUp(c *gin.Context) {
	var user helperstruct.UserReq

	err := c.BindJSON(&user)
	fmt.Println(user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "Error in Binding data",
			Data:       nil,
			Errors:     err,
		})
	}
	userData, err := u.userUsecase.UserSignUp(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Signup failed",
			Data:       response.UserData{},
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 201,
		Message:    "user signup Successfully",
		Data:       userData,
		Errors:     nil,
	})

}

func (u *UserHandler) UserLogin(c *gin.Context) {
	var user helperstruct.LoginReq
	err := c.Bind(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "failed to read the request",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	err = u.userUsecase.UserLogin(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "failed to login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "otp send successfully,  Please enter otp",
		Data:       nil,
		Errors:     nil,
	})
	return

}

func (u *UserHandler) VerifyLogin(c *gin.Context) {
	var otp helperstruct.OTP

	err := c.Bind(&otp)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "failed to read the request",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ss, err := u.userUsecase.VerifyOTP(otp.Code)
	if err != nil {

		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "failed to login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Login Successfully",
		Data:       nil,
		Errors:     nil,
	})
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("UserAuth", ss, 3600*24*30, "", "", false, true)
	return

}

func (u *UserHandler) AddAddress(c *gin.Context) {
	Id, err := handlerutils.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find Id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	var Address helperstruct.Address
	err = c.Bind(&Address)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = u.userUsecase.AddAddress(Id, Address)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't add address",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "address added",
		Data:       nil,
		Errors:     nil,
	})
}

func (u *UserHandler) UpdateAddress(c *gin.Context) {
	paramsId := c.Param("addressId")
	addressId, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant Find ADDRESS ID",
		})
	}
	Id, err := handlerutils.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find Id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	var Address helperstruct.Address
	err = c.Bind(&Address)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = u.userUsecase.UpdateAddress(Id,addressId, Address)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't update address",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "address updated",
		Data:       nil,
		Errors:     nil,
	})

}
