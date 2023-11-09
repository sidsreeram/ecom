package di

import (
	"github.com/ECOMMERCE_PROJECT/pkg/api/handlers"
	config "github.com/ECOMMERCE_PROJECT/pkg/config"
	"github.com/ECOMMERCE_PROJECT/pkg/db"
	repository "github.com/ECOMMERCE_PROJECT/pkg/repository"
	"github.com/google/wire"

	http "github.com/ECOMMERCE_PROJECT/pkg/api"

	usecase "github.com/ECOMMERCE_PROJECT/pkg/usecase"
)

func InitializeAPI1(config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase,
		repository.NewUserRepository,
		repository.NewAdminRepository,
		repository.NewProductRepostiory,
		repository.NewCartRepository,
		repository.NewOrderRepository,
		repository.NewPaymentRepository,
		repository.NewCouponRepository,
		repository.NewWishlistRepository,
		usecase.NewUserUseCase,
		usecase.NewAdminUsecase,
		usecase.NewProductUsecase,
		usecase.NewCartUsecase,
		usecase.NewOrderUsecase,
		usecase.NewPaymentuseCase,
		usecase.NewCouponUsecase,
		usecase.NewWishlistUsecase,
		handlers.NewUserHandelr,
		handlers.NewAdminHandler,
		handlers.NewProductHandler,
		handlers.NewCartHandler,
		handlers.NewOrderHandler,
		handlers.NewPaymentHandler,
		handlers.NewCouponHandler,
		handlers.NewWishlistHandler,
		http.NewServerHTTP,
	)
	return &http.ServerHTTP{}, nil
}
