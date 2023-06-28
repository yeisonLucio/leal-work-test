package models

import "gorm.io/gorm"

type UserReward struct {
	gorm.Model
	UserID     uint   `gorm:"not null"`
	PurchaseID uint   `gorm:"not null"`
	RewardID   uint   `gorm:"not null"`
	Name       string `gorm:"not null"`
	Points     uint
	Coins      uint
	Status     string `gorm:"enum('redeemed','available');default:available"`
}
