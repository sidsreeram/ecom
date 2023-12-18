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
	carthandler *handlers.CartHandler,
	orderhandler *handlers.OrderHandler,
	couponhandelr *handlers.CouponHandler,
	paymenthandler *handlers.PaymentHandler,
	wishlisthandler *handlers.WishlistHandler,
	discounthandler *handlers.DiscountHandler) *ServerHTTP {

	engine := gin.New()
	engine.Use(gin.Logger())

	user := engine.Group("/user")

	{
		user.POST("signup", userHandler.UserSignUp)
		user.POST("login", userHandler.UserLogin)
		user.POST("/otp", userHandler.VerifyLogin)

		category := user.Group("/categories")
		{
			category.GET("/allcategories", productHandler.ListCategories)
			category.GET("/showcategory/:category_id", productHandler.DisplayACategory)
		}
		product := user.Group("/product")
		{
			product.GET("/displayproduct", productHandler.ListAllProduct)
			product.GET("/aproduct/:id", productHandler.DisplayAProduct)
		}
		productitem := user.Group("/productitem")
		{
			productitem.GET("allproductitem", productHandler.DisaplyaAllProductItems)
			productitem.GET("/aproductitem/:id", productHandler.DisplayAproductitem)
		}
		user.Use(middlewares.UserAuth)
		{

			profile := user.Group("/profile")
			{
				profile.GET("view", userHandler.Viewprofile)
				profile.PATCH("edit", userHandler.UserEditProfile)
			}

			address := user.Group("/address")
			{
				address.POST("add", userHandler.AddAddress)
				address.PATCH("update/:addressId", userHandler.UpdateAddress)
			}
			cart := user.Group("/cart")
			{
				cart.POST("/addtocart/:product_item_id", carthandler.AddToCart)
				cart.PATCH("/remove/:product_item_id", carthandler.RemoveFromCart)
				cart.GET("/cart", carthandler.ListCart)
			}
			order := user.Group("/order")
			{
				order.POST("/orderall/:payment_id", orderhandler.OrderAll)
				order.PATCH("/cancel/:order_id", orderhandler.UserCancelOrder)
				order.GET("/view/:order_id", orderhandler.ListAorder)
				order.GET("/viewall", orderhandler.ListAllorder)
				order.PATCH("/return/:order_id", orderhandler.ReturnOrder)
			}
			coupon := user.Group("coupon")
			{
				coupon.PATCH("apply/:code",couponhandelr.ApplyCoupon)
				coupon.PATCH("remove",couponhandelr.RemoveCoupon)
			}
			wishlist := user.Group("/wishlist")
			{
				wishlist.POST("/add/:product_item_id", wishlisthandler.AddToWishlist)
				wishlist.DELETE("/remove/:product_item_id", wishlisthandler.RemoveFromWishlist)
				wishlist.GET("/view", wishlisthandler.ViewAllWishlistItems)
			}
			passwrod := user.Group("/changepassword")
			{
				passwrod.POST("/change", userHandler.ChangePassword)
				passwrod.POST("/verify", userHandler.VerifyForPassword)
			}
			user.GET("order/onlinepayment/pay/:orderId",paymenthandler.CreateRazorpayPayment)
			user.GET("order/downloadinvoice/:orderId",orderhandler.DownloadInvoice)
			// user.GET("payment-handler", paymenthandler.)
			
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
			admincategory.POST("/add", productHandler.CreateCategory)
			admincategory.PATCH("/update/:id", productHandler.UpdateCategory)
			admincategory.DELETE("/delete/:id", productHandler.DeleteCategory)
			admincategory.GET("listall", productHandler.ListCategories)
			admincategory.GET("/show/:id", productHandler.DisplayACategory)
		}
		adminProduct := admin.Group("/product")
		{
			adminProduct.POST("/add", productHandler.AddProduct)
			adminProduct.PATCH("/update/:id", productHandler.UpdateProduct)
			adminProduct.DELETE("/delete/:id", productHandler.DeleteProduct)
			adminProduct.GET("/allproduct", productHandler.ListAllProduct)
			adminProduct.GET("/show/:id", productHandler.DisplayAProduct) //
		}
		adminProductitem := admin.Group("/product-item")
		{
			adminProductitem.POST("add", productHandler.AddProductitem)
			adminProductitem.PATCH("update/:id", productHandler.UpdateProductitem)
			adminProductitem.DELETE("delete/:id", productHandler.DeleteProductItem)
			adminProductitem.GET("/allproductitem", productHandler.DisaplyaAllProductItems)
			adminProductitem.GET("productitem/:id", productHandler.DisplayAproductitem)
			adminProductitem.POST("/uploadimage/:id",productHandler.UploadImage)
			adminProductitem.POST("/uploadimagebinary/:id",productHandler.UploadImageBinary)
			adminProductitem.GET("/image/:id",productHandler.GetProductImages)
		}
		discount:=admin.Group("/discount")
		{
			discount.POST("/add",discounthandler.AddDiscount)
			discount.PATCH("/update/:discountId",discounthandler.UpdateDiscount)
			discount.DELETE("/delete/:discountId",discounthandler.DeleteDiscount)
			discount.GET("/getall",discounthandler.GetAllDiscount)
			discount.GET("/getbyid/:discountId",discounthandler.ViewDiscountbyID)
		}
		coupon:=admin.Group("/coupons")
		{
			coupon.POST("/add",couponhandelr.AddCoupon)
			coupon.PATCH("/update/:couponId",couponhandelr.UpdateCoupon)
			coupon.DELETE("/delete/:couponId",couponhandelr.DeleteCoupon)
			coupon.GET("/view/:couponId",couponhandelr.ViewAcoupon)
			coupon.GET("/viewall",couponhandelr.ViewCoupons)
		}
		order := admin.Group("/order")
		{
			order.PATCH("/update", orderhandler.UpdateOrder)
		}
		sales := admin.Group("/sales")
		{
			sales.GET("/download",adminHandler.DownloadSalesReport)
			sales.GET("/daily", adminHandler.ViewDailySalesReport)
			sales.GET("/weekly", adminHandler.ViewWeelySalesReport)
			sales.GET("/monthly", adminHandler.ViewMonthlySalesReport)
			sales.GET("/yearly", adminHandler.ViewYearlySalesReport)
		}
		admin.GET("/dashboard",adminHandler.GetDashBoard)
	}
	return &ServerHTTP{engine: engine}

}

func (s *ServerHTTP) Start() {
	s.engine.LoadHTMLGlob("../template/*.html")
	s.engine.Run(":3000")
}
