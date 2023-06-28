package models

import "gorm.io/gorm"

type Reward struct {
	gorm.Model
	StoreID     uint    `gorm:"not null"`
	Reward      uint    `gorm:"not null"`
	MinPurchase float64 `gorm:"not null"`
	Type        string  `gorm:"enum('coin','point');not null"`
	Status      string  `gorm:"enum('active','inactive');default:active"`
	UserRewards []UserReward
}
