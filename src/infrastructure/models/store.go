package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Name         string `gorm:"not null"`
	Status       string `gorm:"type:enum('active','inactive');default:active"`
	RewardPoints uint
	RewardCoins  uint
	PurchaseMin  float64
	Branches     []Branch
	Rewards      []Reward
	Transactions []Transaction
}
