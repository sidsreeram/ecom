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
	adminHandler *handlers.AdminHandler,
	productHandler *handlers.ProductHandler,
) *ServerHTTP {
	engine := gin.New()
	engine.Use(gin.Logger())

	user := engine.Group("/user")
	{
		user.POST("signup", userHandler.UserSignUp)
		user.POST("login", userHandler.UserLogin)
		user.POST("/otp", userHandler.VerifyLogin)

		category := user.Group("/categories")
		{
			category.GET("/categories", productHandler.ListCategories)
			category.GET("/showcategory/:category_id", productHandler.DisplayACategory)
		}
		//    product:=user.Group("/product")
		{
			// product.GET("/showproduct",productHandler.)
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
			adminUsers.GET("/finduser", adminHandler.FindUser)
			adminUsers.GET("listallusers", adminHandler.ListAllUsers)
		}
		admincategory := admin.Group("/category")
		{
			admincategory.POST("/addcategory", productHandler.CreateCategory)
			admincategory.PATCH("/updatecategory/:id", productHandler.UpdateCategory)
			admincategory.DELETE("/deletecategory/:category_id", productHandler.DeleteCategory)
			admincategory.GET("listallcategory", productHandler.ListCategories)
			admincategory.GET("/showcategory/:category_id", productHandler.DisplayACategory)
		}
		adminProduct := admin.Group("/product")
		{
			adminProduct.POST("/addproduct", productHandler.AddProduct)
			adminProduct.PATCH("/update/:id", productHandler.UpdateProduct)
			adminProduct.DELETE("/delete/:id", productHandler.DeleteProduct)
		}
		adminProductitem := admin.Group("/product-item")
		{
			adminProductitem.POST("add", productHandler.AddProductitem)
			adminProductitem.PATCH("update/:id", productHandler.UpdateProductitem)
			adminProductitem.DELETE("delete/:id", productHandler.DeleteProductItem)
		}
	}
	return &ServerHTTP{engine: engine}

}

func (s *ServerHTTP) Start() {
	s.engine.Run(":3000")
}
