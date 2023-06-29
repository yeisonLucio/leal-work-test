package di

import (
	"lucio.com/order-service/src/application/usecases"
	"lucio.com/order-service/src/database"
	"lucio.com/order-service/src/infrastructure/controllers"
	"lucio.com/order-service/src/infrastructure/repositories"
)

type Dependencies struct {
	StoreController    *controllers.StoreController
	BranchController   *controllers.BranchController
	RewardController   *controllers.RewardController
	CampaignController *controllers.CampaignController
}

var Container Dependencies

func BuildContainer() {

	storeRepository := &repositories.MysqlStoreRepository{
		DB: database.DB,
	}

	createStoreUC := &usecases.CreateStoreUC{
		StoreRepository: storeRepository,
	}

	Container.StoreController = &controllers.StoreController{
		CreateStoreUC: createStoreUC,
	}

	branchRepository := &repositories.MysqlBranchRepository{
		DB: database.DB,
	}

	createBranchUC := &usecases.CreateBranchUC{
		BranchRepository: branchRepository,
		StoreRepository:  storeRepository,
	}

	campaignRepository := &repositories.MysqlCampaignRepository{
		DB: database.DB,
	}

	branchCampaignRepository := &repositories.MysqlBranchCampaignRepository{
		DB: database.DB,
	}

	createBranchCampaignUC := &usecases.CreateBranchCampaignUC{
		BranchRepository:         branchRepository,
		CampaignRepository:       campaignRepository,
		BranchCampaignRepository: branchCampaignRepository,
	}

	Container.BranchController = &controllers.BranchController{
		CreateBranchUC:         createBranchUC,
		CreateBranchCampaignUC: createBranchCampaignUC,
	}

	rewardRepository := &repositories.MysqlRewardRepository{
		DB: database.DB,
	}

	createRewardUC := &usecases.CreateRewardUC{
		RewardRepository: rewardRepository,
		StoreRepository:  storeRepository,
	}

	Container.RewardController = &controllers.RewardController{
		CreateRewardUC: createRewardUC,
	}

	createCampaignUC := &usecases.CreateCampaignUC{
		CampaignRepository: campaignRepository,
	}

	Container.CampaignController = &controllers.CampaignController{
		CreateCampaignUC: createCampaignUC,
	}

}
