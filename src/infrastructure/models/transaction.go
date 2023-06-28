package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID  uint    `gorm:"not null"`
	StoreID uint    `gorm:"not null"`
	Amount  float64 `gorm:"not null"`
	Points  uint
	Coins   uint
	Type    string `gorm:"type:enum('add','subtract');not null"`
}
