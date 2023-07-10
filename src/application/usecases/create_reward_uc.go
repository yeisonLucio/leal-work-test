package usecases

import (
	"github.com/sirupsen/logrus"
	"lucio.com/order-service/src/domain/contracts/repositories"
	"lucio.com/order-service/src/domain/dto"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/vo"
)

type CreateRewardUC struct {
	RewardRepository repositories.RewardRepository
	StoreRepository  repositories.StoreRepository
	Logger           *logrus.Logger
}

func (c *CreateRewardUC) Execute(createRewardDTO dto.CreateRewardDTO) (*dto.RewardCreatedDTO, error) {
	log := c.Logger.WithFields(logrus.Fields{
		"file":            "create_reward_uc",
		"method":          "Execute",
		"createRewardDTO": createRewardDTO,
	})

	if store := c.StoreRepository.FindByID(createRewardDTO.StoreID); store == nil {
		log.Error(errStoreNotFound)
		return nil, errStoreNotFound
	}

	amountType, err := vo.NewAmountType(createRewardDTO.AmountType)
	if err != nil {
		log.WithError(err).Error("amount type invalid")
		return nil, err
	}

	rewardDB, err := c.RewardRepository.Create(entities.Reward{
		Reward:      createRewardDTO.Reward,
		Description: createRewardDTO.Description,
		MinAmount:   vo.NewAmountFromFloat(createRewardDTO.MinAmount),
		AmountType:  amountType,
		StoreID:     createRewardDTO.StoreID,
		Status:      vo.ActiveStatus,
	})

	if err != nil {
		log.WithError(err).Error("error creating reward")
		return nil, err
	}

	response := dto.RewardCreatedDTO{
		ID:          rewardDB.ID,
		Reward:      rewardDB.Reward,
		Description: rewardDB.Description,
		MinAmount:   rewardDB.MinAmount.Value(),
		AmountType:  rewardDB.AmountType.Value(),
		StoreID:     rewardDB.StoreID,
		Status:      string(rewardDB.Status),
	}

	return &response, nil
}
