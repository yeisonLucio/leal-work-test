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
			users.POST("/:user_id/transactions")
		}
	}

	return app
}
