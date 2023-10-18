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

func InitializeAPI(config.Config) (*http.ServerHTTP, error) {                                  
	wire.Build(db.ConnectDatabase,
		repository.NewUserRepository,
		repository.NewAdminRepository,
		usecase.NewUserUseCase,
		usecase.NewAdminUsecase,
		handlers.NewUserHandelr,
		handlers.NewAdminHandler,
		http.NewServerHTTP,
	)
	return &http.ServerHTTP{}, nil
}
