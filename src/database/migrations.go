package database

import "lucio.com/order-service/src/infrastructure/models"

func RunMigrations() error {
	return DB.AutoMigrate(
		&models.User{},
		&models.Store{},
		&models.Branch{},
		&models.Campaign{},
		&models.Transaction{},
		&models.Reward{},
		&models.BranchCampaign{},
	)

}
