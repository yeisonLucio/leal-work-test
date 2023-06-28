package models

import (
	"time"

	"gorm.io/gorm"
)

type Campaign struct {
	gorm.Model
	Description string    `gorm:"not null"`
	BranchID    uint      `gorm:"not null"`
	StartDate   time.Time `gorm:"not null"`
	EndDate     time.Time `gorm:"not null"`
	Operation   string    `gorm:"not null"`
	MinPurchase float64   `gorm:"not null"`
}
