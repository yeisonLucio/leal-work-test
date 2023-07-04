package repositories

import (
	"gorm.io/gorm"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/vo"
	"lucio.com/order-service/src/infrastructure/models"
)

type MysqlRewardRepository struct {
	DB *gorm.DB
}

func (m *MysqlRewardRepository) Create(reward entities.Reward) (*entities.Reward, error) {
	rewardDB := models.Reward{
		StoreID:     reward.StoreID,
		Reward:      reward.Reward,
		MinAmount:   reward.MinAmount.Value(),
		Description: reward.Description,
		AmountType:  reward.AmountType.Value(),
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

	amountType, _ := vo.NewAmountType(rewardDB.AmountType)

	reward := entities.Reward{
		ID:          rewardDB.ID,
		StoreID:     rewardDB.StoreID,
		Reward:      rewardDB.Reward,
		MinAmount:   vo.NewAmountFromFloat(rewardDB.MinAmount),
		Description: rewardDB.Description,
		AmountType:  amountType,
		Status:      vo.Status(rewardDB.Status),
	}

	return &reward, nil
}
