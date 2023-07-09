package database

import (
	"gorm.io/gorm"
	"lucio.com/order-service/src/infrastructure/models"
)

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Store{},
		&models.Branch{},
		&models.Campaign{},
		&models.Transaction{},
		&models.Reward{},
		&models.BranchCampaign{},
	)

}
