package usecases

import (
	"lucio.com/order-service/src/domain/contracts/repositories"
	"lucio.com/order-service/src/domain/dto"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/valueobjects"
)

type CreateStoreUC struct {
	StoreRepository repositories.StoreRepository
}

func (c *CreateStoreUC) Execute(createStoreDTO dto.CreateStoreDTO) (*dto.StoreCreatedDTO, error) {
	var minAmount valueobjects.Amount
	minAmount.NewFromFloat(createStoreDTO.MinAmount)

	createdStore, err := c.StoreRepository.Create(entities.Store{
		Name:         createStoreDTO.Name,
		Status:       valueobjects.ActiveStatus,
		RewardPoints: createStoreDTO.RewardPoints,
		RewardCoins:  createStoreDTO.RewardCoins,
		MinAmount:    minAmount,
	})

	if err != nil {
		return nil, err
	}

	response := dto.StoreCreatedDTO{
		ID:           createdStore.ID,
		Name:         createdStore.Name,
		RewardPoints: createdStore.RewardPoints,
		RewardCoins:  createdStore.RewardCoins,
		MinAmount:    createdStore.MinAmount.GetValue(),
		Status:       string(createdStore.Status),
	}

	return &response, nil
}
