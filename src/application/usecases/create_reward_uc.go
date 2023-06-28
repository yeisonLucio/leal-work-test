package usecases

import (
	"errors"

	"lucio.com/order-service/src/domain/contracts/repositories"
	"lucio.com/order-service/src/domain/dto"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/valueobjects"
)

type CreateRewardUC struct {
	RewardRepository repositories.RewardRepository
	StoreRepository  repositories.StoreRepository
}

func (c *CreateRewardUC) Execute(createRewardDTO dto.CreateRewardDTO) (*dto.RewardCreatedDTO, error) {
	if _, err := c.StoreRepository.FindByID(createRewardDTO.StoreID); err != nil {
		return nil, errors.New("la tienda ingresada no existe")
	}

	var minAmount valueobjects.Amount
	if err := minAmount.NewFromString(createRewardDTO.MinAmount); err != nil {
		return nil, err
	}

	var amountType valueobjects.AmountType
	if err := amountType.New(createRewardDTO.AmountType); err != nil {
		return nil, err
	}

	rewardDB, err := c.RewardRepository.Create(entities.Reward{
		Reward:      createRewardDTO.Reward,
		Description: createRewardDTO.Description,
		MinAmount:   minAmount,
		AmountType:  amountType,
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
		MinAmount:   rewardDB.MinAmount.GetValue(),
		AmountType:  rewardDB.AmountType.GetValue(),
		StoreID:     rewardDB.StoreID,
		Status:      string(rewardDB.Status),
	}

	return &response, nil
}
