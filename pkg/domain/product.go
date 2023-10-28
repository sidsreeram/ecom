package domain

import "time"

type Category struct {
	ID           uint   `gorm:"primaryKey;unique;not null"`
	CategoryName string `gorm:"unique;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Product struct {
	ID          uint   `gorm:"primaryKey;unique;not null"`
	ProductName string `gorm:"unique;not null"`
	Description string
	Brand       string // Combined field for Brand and Manufacturer
	CategoryID  uint
	Category    Category `gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProductItem struct {
	ID         uint `gorm:"primaryKey;unique;not null"`
	ProductID  uint
	Product    Product `gorm:"foreignKey:ProductID"`
	SKU        string  `gorm:"not null"`
	QtyInStock int     // Quantity of the sport product in stock
	Price      float64 // Price of the sport product
	InStock    bool    // Indicates whether the sport product is in stock or not
	Color      string  // Color of the sport product
	Size       string  // Size of the sport product (e.g., "Small," "Medium," "Large," etc.)
	Rating     float64 // Customer rating of the sport product (e.g., 4.5 for a 4.5-star rating)
	Imag        string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ProductImage struct {
	ID            uint `gorm:"primaryKey;unique;not null"`
	ProductItemID uint
	ProductItem   ProductItem `gorm:"foreignKey:ProductItemID"`
	FileName      string
}
