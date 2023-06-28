package repositories

import (
	"gorm.io/gorm"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/valueobjects"
	"lucio.com/order-service/src/infrastructure/models"
)

type MysqlRewardRepository struct {
	DB *gorm.DB
}

func (m *MysqlRewardRepository) Create(reward entities.Reward) (*entities.Reward, error) {
	rewardDB := models.Reward{
		StoreID:     reward.StoreID,
		Reward:      reward.Reward,
		MinAmount:   reward.MinAmount.GetValue(),
		Description: reward.Description,
		AmountType:  reward.AmountType.GetValue(),
		Status:      string(reward.Status),
	}

	if result := m.DB.Create(&rewardDB); result.Error != nil {
		return nil, result.Error
	}

	reward.ID = rewardDB.ID

	return &reward, nil
}

func (m *MysqlRewardRepository) FindByID(ID uint) (*entities.Reward, error) {
	var rewardDB models.Reward

	if result := m.DB.Find(&rewardDB, ID); result.Error != nil {
		return nil, result.Error
	}

	var minAmount valueobjects.MinAmount
	minAmount.NewFromFloat(rewardDB.MinAmount)

	var amountType valueobjects.AmountType
	amountType.New(rewardDB.AmountType)

	reward := entities.Reward{
		ID:          rewardDB.ID,
		StoreID:     rewardDB.StoreID,
		Reward:      rewardDB.Reward,
		MinAmount:   minAmount,
		Description: rewardDB.Description,
		AmountType:  amountType,
		Status:      valueobjects.Status(rewardDB.Status),
	}

	return &reward, nil
}
