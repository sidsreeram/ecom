package di

import (
	"github.com/ECOMMERCE_PROJECT/pkg/api/handlers"
	config "github.com/ECOMMERCE_PROJECT/pkg/config"
	"github.com/ECOMMERCE_PROJECT/pkg/db"
	"github.com/ECOMMERCE_PROJECT/pkg/repository"
	"github.com/google/wire"

	http "github.com/ECOMMERCE_PROJECT/pkg/api"

	usecase "github.com/ECOMMERCE_PROJECT/pkg/usecase"
)

func InitializeAPI1(config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase,
		repository.NewUserRepository,
		repository.NewAdminRepository,
		repository.NewProductRepostiory,
		usecase.NewUserUseCase,
		usecase.NewAdminUsecase,
		usecase.NewProductUsecase,
		handlers.NewUserHandelr,
		handlers.NewAdminHandler,
		handlers.NewProductHandler,
		http.NewServerHTTP,
	)
	return &http.ServerHTTP{}, nil
}
