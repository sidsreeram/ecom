// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/ECOMMERCE_PROJECT/pkg/api"
	"github.com/ECOMMERCE_PROJECT/pkg/api/handlers"
	"github.com/ECOMMERCE_PROJECT/pkg/config"
	"github.com/ECOMMERCE_PROJECT/pkg/db"
	"github.com/ECOMMERCE_PROJECT/pkg/repository"
	"github.com/ECOMMERCE_PROJECT/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(configConfig config.Config) (*api.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(configConfig)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	cartRepository := repository.NewCartRepository(gormDB)
	cartUseCase := usecase.NewCartUsecase(cartRepository)
	userHandler := handlers.NewUserHandelr(userUseCase, cartUseCase)
	adminRepository := repository.NewAdminRepository(gormDB)
	adminUsecase := usecase.NewAdminUsecase(adminRepository)
	adminHandler := handlers.NewAdminHandler(adminUsecase)
	productRepository := repository.NewProductRepostiory(gormDB)
	productUsecase := usecase.NewProductUsecase(productRepository)
	productHandler := handlers.NewProductHandler(productUsecase)
	cartHandler := handlers.NewCartHandler(cartUseCase)
	orderRepository := repository.NewOrderRepository(gormDB)
	orderUsercase := usecase.NewOrderUsecase(orderRepository)
	orderHandler := handlers.NewOrderHandler(orderUsercase)
	couponRepository := repository.NewCouponRepository(gormDB)
	couponusecase := usecase.NewCouponUsecase(couponRepository)
	couponHandler := handlers.NewCouponHandler(couponusecase)
	paymentRepository := repository.NewPaymentRepository(gormDB)
	paymentUsecase := usecase.NewPaymentuseCase(paymentRepository, orderRepository, configConfig)
	paymentHandler := handlers.NewPaymentHandler(paymentUsecase)
	wishlistRepository := repository.NewWishlistRepository(gormDB)
	wishlistUsecase := usecase.NewWishlistUsecase(wishlistRepository)
	wishlistHandler := handlers.NewWishlistHandler(wishlistUsecase)
	discountRepository := repository.NewDiscountRepository(gormDB)
	discountUsecase := usecase.NewDiscountUsecase(discountRepository)
	discountHandler := handlers.NewDiscountHandler(discountUsecase)
	serverHTTP := api.NewServerHTTP(userHandler, adminHandler, productHandler, cartHandler, orderHandler, couponHandler, paymentHandler, wishlistHandler, discountHandler)
	return serverHTTP, nil
}
