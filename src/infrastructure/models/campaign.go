package models

import (
	"gorm.io/gorm"
)

type Campaign struct {
	gorm.Model
	Description     string `gorm:"not null"`
	Status          string `gorm:"default:'active'"`
	BranchCampaigns []BranchCampaign
}
