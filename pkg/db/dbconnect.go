package db

import (
	"log"

	"github.com/ECOMMERCE_PROJECT/pkg/config"

	"github.com/ECOMMERCE_PROJECT/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := cfg.DBURL
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if dbErr != nil {
		log.Fatalf("Error connecting to the database: %v", dbErr)
	}

	db.AutoMigrate(
		&domain.Users{},
		&domain.UserInfo{},
		&domain.Admins{},
		&domain.Product{},
	)
	return db, dbErr
}
