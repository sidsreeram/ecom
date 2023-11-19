package domain

import "time"

type Coupons struct {
	Id                    uint   `gorm:"primaryKey;unique;not null"`
	Code                  string ``
	DiscountPercent       float64
	DiscountMaximumAmount float64
	MinimumPurchaseAmount float64
	ExpirationDate        time.Time
}
type Discount struct {
	Id                    uint   `gorm:"primaryKey;unique;not null"`
	Category_id           int
	Category                Category `gorm:"foreignKey:category_id" json:"-"`
	DiscountPercent       float64
	DiscountMaximumAmount float64
	MinimumPurchaseAmount float64
	ExpirationDate        time.Time
}