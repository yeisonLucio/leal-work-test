package models

import "gorm.io/gorm"

type Reward struct {
	gorm.Model
	StoreID     uint    `gorm:"not null"`
	Reward      string  `gorm:"not null"`
	MinAmount   float64 `gorm:"not null"`
	Description string  `gorm:"not null"`
	AmountType  string  `gorm:"type:enum('coin','point');not null"`
	Status      string  `gorm:"type:enum('active','inactive');default:active"`
}
