package domain

import "time"

type OTP struct {
	ID        uint `gorm:"primaryKey"`
	UserID    int
	UserEmail string
	Code      string
	ExpiresAt time.Time
	CreatedAt time.Time
}
