package handlers

import (
	"fmt"
	"net/http"

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
	
	ss, err := u.userUsecase.UserLogin(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "failed to login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("UserAuth", ss, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "logined successfuly",
		Data:       nil,
		Errors:     nil,
	})
 return
}
