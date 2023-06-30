package repositories

import "lucio.com/order-service/src/domain/entities"

type StoreRepository interface {
	Create(store entities.Store) (*entities.Store, error)
	FindByID(ID uint) *entities.Store
	FindByBranchID(branchID uint) *entities.Store
}
