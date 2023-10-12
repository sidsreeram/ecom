package api

import (
	"github.com/ECOMMERCE_PROJECT/pkg/api/handlers"
	"github.com/ECOMMERCE_PROJECT/pkg/api/middlewares"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(
	userHandler *handlers.UserHandler,
	otpHandler *handlers.OtpHandler,
	adminHandler *handlers.AdminHandler,
) *ServerHTTP {
	engine := gin.New()
	engine.Use(gin.Logger())

	user := engine.Group("/user")
	{
		user.POST("signup", userHandler.UserSignUp)
		user.POST("login", userHandler.UserLogin)

		otp := user.Group("/otp")
		{
			otp.POST("send", otpHandler.SendOtp)
			otp.POST("verfiy", otpHandler.ValidateOtp)
		}
	}
	admin := engine.Group("/admin")
	{
		admin.POST("login", adminHandler.AdminLoging)
		admin.Use(middlewares.AdminAuth)
		{
			admin.POST("createadmin", adminHandler.CreateAdmin)
			admin.POST("logout", adminHandler.AdminLogout)

		}
		adminUsers := admin.Group("/user")
		{

			adminUsers.PATCH("/block", adminHandler.BlockUser)
			adminUsers.PATCH("/unblock/:user_id", adminHandler.UnblockUser)
		}
	}
	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {

	sh.engine.Run(":3000")
}
