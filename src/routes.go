package src

import (
	"github.com/gin-gonic/gin"
	"lucio.com/order-service/src/di"
)

func GetRoutes(app *gin.Engine) *gin.Engine {
	api := app.Group("api/v1")
	{
		stores := api.Group("stores")
		{
			stores.POST("/", di.Container.StoreController.Create)
		}

		branches := api.Group("branches")
		{
			branches.POST("/", di.Container.BranchController.Create)
		}

		rewards := api.Group("rewards")
		{
			rewards.POST("/", di.Container.RewardController.Create)
		}
	}

	return app
}
