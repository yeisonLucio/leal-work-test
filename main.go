package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"lucio.com/order-service/src"
	"lucio.com/order-service/src/config/database"
	"lucio.com/order-service/src/config/redis"
	"lucio.com/order-service/src/di"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if err := database.Connect(); err != nil {
		panic(err)
	}

	if err := database.RunMigrations(database.DB); err != nil {
		fmt.Println("error corriendo las migraciones")
	}

	redis.ConnectRedis()

	di.BuildContainer()

	app := src.GetApp()

	app.Run(":8080")
}
