package database

import "lucio.com/order-service/src/infrastructure/models"

func RunMigrations() (err error) {
	err = DB.AutoMigrate(&models.Store{})
	return
}
