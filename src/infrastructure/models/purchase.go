package models

import "gorm.io/gorm"

type Purchase struct {
	gorm.Model
	UserID      uint    `gorm:"not null"`
	BranchID    uint    `gorm:"not null"`
	Amount      float64 `gorm:"not null"`
	Description string
	UserRewards []UserReward
}
