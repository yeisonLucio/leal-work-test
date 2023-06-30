package src

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "lucio.com/order-service/docs"
	"lucio.com/order-service/src/di"
)

func GetRoutes(app *gin.Engine) *gin.Engine {
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", os.Getenv("APP_PORT"))

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	api := app.Group("api/v1")
	{
		stores := api.Group("stores")
		{
			stores.POST("/", di.Container.StoreController.Create)
			stores.POST("/:store_id/rewards", di.Container.RewardController.Create)
			stores.POST("/:store_id/branches", di.Container.BranchController.Create)
		}

		campaigns := api.Group("campaigns")
		{
			campaigns.POST("/", di.Container.CampaignController.Create)

			campaigns.POST(
				"/:campaign_id/branches/:branch_id",
				di.Container.BranchController.CreateBranchCampaign,
			)

			campaigns.POST(
				"/:campaign_id/stores/:store_id",
				di.Container.BranchController.AddCampaignToBranches,
			)

			campaigns.GET(
				"/branches/:branch_id",
				di.Container.BranchController.GetBranchCampaignsByBranch,
			)
		}

		users := api.Group("users")
		{
			users.POST("/", di.Container.UserController.Create)
			users.POST(
				"/:user_id/transactions/branches/:branch_id",
				di.Container.UserController.RegisterTransaction,
			)
		}
	}

	return app
}
