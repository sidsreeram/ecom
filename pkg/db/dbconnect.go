package db

import (

	// "github.com/ECOMMERCE_PROJECT/pkg/config"
	"github.com/ECOMMERCE_PROJECT/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	psqlInfo := "host=localhost user=postgres dbname=ecommerce_project port=5432 password=partner"
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	db.AutoMigrate(
		&domain.Users{},
		&domain.UserInfo{},
		&domain.Admins{},
		&domain.Product{},
	)
	return db, dbErr
}
