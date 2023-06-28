package database

import "lucio.com/order-service/src/infrastructure/models"

func RunMigrations() error {
	return DB.AutoMigrate(
		&models.UserReward{},
		&models.User{},
		&models.Branch{},
		&models.Store{},
		&models.Campaign{},
		&models.Purchase{},
		&models.Reward{},
	)

}
