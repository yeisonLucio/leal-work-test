package src

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"lucio.com/order-service/src/database"
)

func GetApp() *gin.Engine {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if err := database.Connect(); err != nil {
		panic(err)
	}

	if err := database.RunMigrations(); err != nil {
		fmt.Println("error corriendo las migraciones")
	}

	app := gin.Default()

	app = GetRoutes(app)

	return app
}
