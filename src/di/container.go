package di

import (
	"github.com/sirupsen/logrus"
	"lucio.com/order-service/src/application/usecases"
	"lucio.com/order-service/src/config/database"
	"lucio.com/order-service/src/config/redis"
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
	logger := logrus.New()

	//repositories
	storeRepository := &repositories.MysqlStoreRepository{
		DB:     database.DB,
		Logger: logger,
	}

	branchRepository := &repositories.MysqlBranchRepository{
		DB:     database.DB,
		Logger: logger,
	}

	campaignRepository := &repositories.MysqlCampaignRepository{
		DB:     database.DB,
		Logger: logger,
	}

	branchCampaignRepository := &repositories.MysqlBranchCampaignRepository{
		DB:     database.DB,
		Logger: logger,
	}

	rewardRepository := &repositories.MysqlRewardRepository{
		DB:     database.DB,
		Logger: logger,
	}

	cacheRepository := &repositories.RedisRepository{
		RedisClient: redis.RedisClient,
		Logger:      logger,
	}

	userRepository := &repositories.MysqlUserRepository{
		DB:     database.DB,
		Logger: logger,
	}

	transactionRepository := &repositories.MysqlTransactionRepository{
		DB:     database.DB,
		Logger: logger,
	}

	//use cases

	createStoreUC := &usecases.CreateStoreUC{
		StoreRepository: storeRepository,
		Logger:          logger,
	}

	createBranchUC := &usecases.CreateBranchUC{
		BranchRepository: branchRepository,
		StoreRepository:  storeRepository,
		Logger:           logger,
	}

	createBranchCampaignUC := &usecases.CreateBranchCampaignUC{
		BranchRepository:         branchRepository,
		CampaignRepository:       campaignRepository,
		BranchCampaignRepository: branchCampaignRepository,
		Logger:                   logger,
	}

	addCampaignToStoreUC := &usecases.AddCampaignToStoreUC{
		BranchRepository:       branchRepository,
		CreateBranchCampaignUC: createBranchCampaignUC,
		CampaignRepository:     campaignRepository,
		StoreRepository:        storeRepository,
		Logger:                 logger,
	}

	createRewardUC := &usecases.CreateRewardUC{
		RewardRepository: rewardRepository,
		StoreRepository:  storeRepository,
		Logger:           logger,
	}

	createCampaignUC := &usecases.CreateCampaignUC{
		CampaignRepository: campaignRepository,
		Logger:             logger,
	}

	getBranchCampaignUC := &usecases.GetBranchCampaignsUC{
		BranchCampaignRepository: branchCampaignRepository,
		CacheRepository:          cacheRepository,
	}

	createUserUC := &usecases.CreateUserUC{
		UserRepository: userRepository,
		Logger:         logger,
	}

	calculateCampaignRewardsUC := &usecases.CalculateCampaignRewardsUC{
		BranchCampaignRepository: branchCampaignRepository,
	}

	createTransactionUC := &usecases.CreateTransactionUC{
		StoreRepository:            storeRepository,
		TransactionRepository:      transactionRepository,
		UserRepository:             userRepository,
		CalculateCampaignRewardsUC: calculateCampaignRewardsUC,
		Logger:                     logger,
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
		CreateUserUC:        createUserUC,
		CreateTransactionUC: createTransactionUC,
	}
}
