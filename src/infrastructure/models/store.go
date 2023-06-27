package models

type Store struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	Name   string
	Status string `gorm:"enum('active','inactive');default:active"`
}
