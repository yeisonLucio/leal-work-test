package models

import "gorm.io/gorm"

type Branch struct {
	gorm.Model
	StoreID   uint   `gorm:"not null"`
	Name      string `gorm:"not null"`
	Status    string `gorm:"type:enum('active','inactive');default:active"`
	Purchases []Purchase
	Campaigns []Campaign
}
