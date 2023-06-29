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
	UserController     *controllers.UserController
}

var Container Dependencies

func BuildContainer() {

	//repositories
	storeRepository := &repositories.MysqlStoreRepository{
		DB: database.DB,
	}

	branchRepository := &repositories.MysqlBranchRepository{
		DB: database.DB,
	}

	campaignRepository := &repositories.MysqlCampaignRepository{
		DB: database.DB,
	}

	branchCampaignRepository := &repositories.MysqlBranchCampaignRepository{
		DB: database.DB,
	}

	rewardRepository := &repositories.MysqlRewardRepository{
		DB: database.DB,
	}

	cacheRepository := &repositories.RedisRepository{
		RedisClient: database.RedisClient,
	}

	userRepository := &repositories.MysqlUserRepository{
		DB: database.DB,
	}

	//use cases

	createStoreUC := &usecases.CreateStoreUC{
		StoreRepository: storeRepository,
	}

	createBranchUC := &usecases.CreateBranchUC{
		BranchRepository: branchRepository,
		StoreRepository:  storeRepository,
	}

	createBranchCampaignUC := &usecases.CreateBranchCampaignUC{
		BranchRepository:         branchRepository,
		CampaignRepository:       campaignRepository,
		BranchCampaignRepository: branchCampaignRepository,
	}

	addCampaignToStoreUC := &usecases.AddCampaignToStoreUC{
		BranchRepository:       branchRepository,
		CreateBranchCampaignUC: createBranchCampaignUC,
		CampaignRepository:     campaignRepository,
		StoreRepository:        storeRepository,
	}

	createRewardUC := &usecases.CreateRewardUC{
		RewardRepository: rewardRepository,
		StoreRepository:  storeRepository,
	}

	createCampaignUC := &usecases.CreateCampaignUC{
		CampaignRepository: campaignRepository,
	}

	getBranchCampaignUC := &usecases.GetBranchCampaignsUC{
		BranchCampaignRepository: branchCampaignRepository,
		CacheRepository:          cacheRepository,
	}

	createUserUC := &usecases.CreateUserUC{
		UserRepository: userRepository,
	}

	//controllers

	Container.StoreController = &controllers.StoreController{
		CreateStoreUC: createStoreUC,
	}

	Container.RewardController = &controllers.RewardController{
		CreateRewardUC: createRewardUC,
	}

	Container.CampaignController = &controllers.CampaignController{
		CreateCampaignUC: createCampaignUC,
	}

	Container.BranchController = &controllers.BranchController{
		CreateBranchUC:         createBranchUC,
		CreateBranchCampaignUC: createBranchCampaignUC,
		AddCampaignToStoreUC:   addCampaignToStoreUC,
		GetBranchCampaignsUC:   getBranchCampaignUC,
	}

	Container.UserController = &controllers.UserController{
		CreateUserUC: createUserUC,
	}
}
