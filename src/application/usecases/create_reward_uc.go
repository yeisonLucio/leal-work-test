package usecases

import (
	"lucio.com/order-service/src/domain/contracts/repositories"
	"lucio.com/order-service/src/domain/dto"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/valueobjects"
)

type CreateRewardUC struct {
	RewardRepository repositories.RewardRepository
}

func (c *CreateRewardUC) Execute(createRewardDTO dto.CreateRewardDTO) (*dto.RewardCreatedDTO, error) {
	rewardDB, err := c.RewardRepository.Create(entities.Reward{
		Reward:      createRewardDTO.Reward,
		Description: createRewardDTO.Description,
		MinAmount:   createRewardDTO.MinAmount,
		AmountType:  createRewardDTO.AmountType,
		StoreID:     createRewardDTO.StoreID,
		Status:      valueobjects.ActiveStatus,
	})

	if err != nil {
		return nil, err
	}

	response := dto.RewardCreatedDTO{
		ID:          rewardDB.ID,
		Reward:      rewardDB.Reward,
		Description: rewardDB.Description,
		MinAmount:   rewardDB.MinAmount,
		AmountType:  rewardDB.AmountType,
		StoreID:     rewardDB.StoreID,
		Status:      string(rewardDB.Status),
	}

	return &response, nil
}
