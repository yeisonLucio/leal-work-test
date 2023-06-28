package models

import (
	"time"

	"gorm.io/gorm"
)

type BranchCampaign struct {
	gorm.Model
	BranchID      uint      `gorm:"not null"`
	CampaignID    uint      `gorm:"not null"`
	StartDate     time.Time `gorm:"not null"`
	EndDate       time.Time `gorm:"not null"`
	Operator      string    `gorm:"not null"`
	OperatorValue uint      `gorm:"not null"`
	MinAmount     float64   `gorm:"not null"`
}
