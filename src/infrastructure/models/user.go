package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name           string `gorm:"not null"`
	Identification string `gorm:"not null;index:idx_identification,unique"`
	Status         string `gorm:"type:enum('active','inactive');default:active"`
	Transactions   []Transaction
}
