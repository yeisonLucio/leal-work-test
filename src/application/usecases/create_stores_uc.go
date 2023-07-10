package usecases

import (
	"github.com/sirupsen/logrus"
	"lucio.com/order-service/src/domain/contracts/repositories"
	"lucio.com/order-service/src/domain/dto"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/vo"
)

type CreateStoreUC struct {
	StoreRepository repositories.StoreRepository
	Logger          *logrus.Logger
}

func (c *CreateStoreUC) Execute(createStoreDTO dto.CreateStoreDTO) (*dto.StoreCreatedDTO, error) {
	log := c.Logger.WithFields(logrus.Fields{
		"file":           "create_store_uc",
		"method":         "Execute",
		"createStoreDTO": createStoreDTO,
	})

	createdStore, err := c.StoreRepository.Create(entities.Store{
		Name:         createStoreDTO.Name,
		Status:       vo.ActiveStatus,
		RewardPoints: createStoreDTO.RewardPoints,
		RewardCoins:  createStoreDTO.RewardCoins,
		MinAmount:    vo.NewAmountFromFloat(createStoreDTO.MinAmount),
	})

	if err != nil {
		log.WithError(err).Error("Error creating store")
		return nil, err
	}

	response := dto.StoreCreatedDTO{
		ID:           createdStore.ID,
		Name:         createdStore.Name,
		RewardPoints: createdStore.RewardPoints,
		RewardCoins:  createdStore.RewardCoins,
		MinAmount:    createdStore.MinAmount.Value(),
		Status:       string(createdStore.Status),
	}

	return &response, nil
}
