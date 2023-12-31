package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Name         string `gorm:"not null"`
	Status       string `gorm:"default:active"`
	RewardPoints uint
	RewardCoins  uint
	MinAmount    float64
	Branches     []Branch
	Rewards      []Reward
}
