package src

import (
	"github.com/gin-gonic/gin"
	"lucio.com/order-service/src/di"
)

func GetRoutes(app *gin.Engine) *gin.Engine {
	api := app.Group("api")
	{
		stores := api.Group("stores")
		{
			stores.POST("/v1", di.Container.StoreController.Create)
		}
	}

	return app
}
