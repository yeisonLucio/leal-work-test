package di

import (
	"lucio.com/order-service/src/application/usecases"
	"lucio.com/order-service/src/database"
	"lucio.com/order-service/src/infrastructure/controllers"
	"lucio.com/order-service/src/infrastructure/repositories"
)

type Dependencies struct {
	StoreController *controllers.StoreController
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
}
