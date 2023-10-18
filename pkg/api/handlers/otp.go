package handlers

import (
	"fmt"
	"net/http"

	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
	"github.com/ECOMMERCE_PROJECT/pkg/config"
	services "github.com/ECOMMERCE_PROJECT/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type OtpHandler struct {
  otpUsecase services.OtpUseCase
  userUseCase services.UserUseCase
  cfg config.Config
}
func NewOtpHandler(cfg config.Config, otpUseCase services.OtpUseCase, userUseCase services.UserUseCase) *OtpHandler {
	return &OtpHandler{
		otpUsecase:  otpUseCase,
		userUseCase: userUseCase,
		cfg:         cfg,
	}
}

  // func (op *OtpHandler) SendOtp(c *gin.Context) {
  //   var phno helperstruct.OTPData
  //   err := c.Bind(&phno)
  //   if err != nil {
  //     fmt.Println("e1")
  //     c.JSON(http.StatusUnprocessableEntity, response.Response{
  //       StatusCode: 422,
  //       Message:    "unable to process the request",
  //       Data:       nil,
  //       Errors:     err.Error(),
  //     })
  //     fmt.Println("e2")
  //     return
  //   }
  
  //   isSignIn, err := op.userUseCase.IsSignIn(phno.PhoneNumber)
  //   if err != nil {
  //     c.JSON(http.StatusBadRequest, response.Response{
  //       StatusCode: 400,
  //       Message:    "No user with this phonenumber",
  //       Data:       nil,
  //       Errors:     err.Error(),
  //     })
  //     return
  //   }
  
  //   if !isSignIn {
  //     fmt.Println("login err")
  //     c.JSON(http.StatusBadRequest, response.Response{
  //       StatusCode: 400,
  //       Message:    "no user found",
  //       Data:       nil,
  //       Errors:     nil,
  //     })
  //     fmt.Println("login err2")
  //     return
  //   }
  
  //   err = op.otpUsecase.SendOtp(c.Request.Context(), phno)
  
  //   if err != nil {
  //     c.JSON(http.StatusBadRequest, response.Response{
  //       StatusCode: 400,
  //       Message:    "creatingfailed",
  //       Data:       nil,
  //       Errors:     err.Error(),
  //     })
  //     return
  //   }
  
  //   c.JSON(http.StatusCreated, response.Response{
  //     StatusCode: 201,
  //     Message:    "otp send",
  //     Data:       nil,
  //     Errors:     nil,
  //   })
  // }
  
  // func (op *OtpHandler) ValidateOtp(c *gin.Context) {
  //   var otpDetails helperstruct.VerifyOtp
  //   err := c.Bind(&otpDetails)
  
  //   if err != nil {
  //     c.JSON(http.StatusBadRequest, response.Response{
  //       StatusCode: 400,
  //       Message:    "can't bind",
  //       Data:       nil,
  //       Errors:     err.Error(),
  //     })
  //     return
  //   }
  //   resp, err := op.otpUsecase.ValidateOtp(otpDetails)
  
  //   if err != nil {
  //     c.JSON(http.StatusBadRequest, response.Response{
  //       StatusCode: 400,
  //       Message:    "can't bind",
  //       Data:       nil,
  //       Errors:     err.Error(),
  //     })
  //   } else if *resp.Status != "approved" {
  //     c.JSON(http.StatusBadRequest, response.Response{
  //       StatusCode: 400,
  //       Message:    "incorect",
  //       Data:       nil,
  //       Errors:     "incorect",
  //     })
  //     return
  //   }
  //   ss, err := op.userUseCase.OtpLogin(otpDetails.User.PhoneNumber)
  //   if err != nil {
  //     c.JSON(http.StatusBadRequest, response.Response{
  //       StatusCode: 400,
  //       Message:    "failed to login",
  //       Data:       nil,
  //       Errors:     err.Error(),
  //     })
  //     return
  //   }
  
  //   c.SetSameSite(http.SameSiteLaxMode)
  //   c.SetCookie("UserAuth", ss, 3600*24*30, "", "", false, true)
  //   c.JSON(http.StatusOK, response.Response{
  //     StatusCode: 200,
  //     Message:    "login successful",
  //     Data:       nil,
  //     Errors:     nil,
  //   })
  // }
  
  func (op*OtpHandler)GenerateEOTP(c *gin.Context){
    var email helperstruct.EMAILOTP
    
  }