package src

import (
	"github.com/gin-gonic/gin"
)

func GetApp() *gin.Engine {
	app := gin.Default()

	app = GetRoutes(app)

	return app
}
