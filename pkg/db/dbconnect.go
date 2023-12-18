package db

import (
	"log"

	"github.com/ECOMMERCE_PROJECT/pkg/config"

	"github.com/ECOMMERCE_PROJECT/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := cfg.DBURL
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})
	if dbErr != nil {
		log.Fatalf("Error connecting to the database: %v", dbErr)
	}

	db.AutoMigrate(
		&domain.Users{},
		&domain.UserInfo{},
		&domain.Address{},
		&domain.OTP{},
		&domain.Carts{},
		&domain.CartItem{},
		&domain.Admins{},
		&domain.Product{},
		&domain.Category{},
		&domain.ProductItem{},
		&domain.ProductImage{},
		&domain.PaymentType{},
		&domain.Orders{},
		&domain.OrderItem{},
		&domain.OrderStatus{},
		&domain.PaymentDetails{},
		&domain.PaymentStatus{},
		&domain.Wishlist{},
		&domain.Coupons{},
		&domain.Discount{},
		&domain.ProITemImage{},
	)
	return db, dbErr
}
