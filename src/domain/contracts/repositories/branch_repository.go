package repositories

import "lucio.com/order-service/src/domain/entities"

type BranchRepository interface {
	Create(branch entities.Branch) (*entities.Branch, error)
	FindByID(ID uint) *entities.Branch
	GetIdsByStoreID(storeID uint) []uint
}
