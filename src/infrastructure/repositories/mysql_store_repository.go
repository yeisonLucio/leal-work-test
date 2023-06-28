package repositories

import (
	"gorm.io/gorm"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/valueobjects"
	"lucio.com/order-service/src/infrastructure/models"
)

type MysqlStoreRepository struct {
	DB *gorm.DB
}

func (m *MysqlStoreRepository) Create(store entities.Store) (*entities.Store, error) {
	storeDB := models.Store{
		Name:         store.Name,
		Status:       string(store.Status),
		RewardPoints: store.RewardPoints,
		RewardCoins:  store.RewardCoins,
		MinAmount:    store.MinAmount.GetValue(),
	}

	if result := m.DB.Create(&storeDB); result.Error != nil {
		return nil, result.Error
	}

	store.ID = storeDB.ID

	return &store, nil
}

func (m *MysqlStoreRepository) FindByID(ID uint) (*entities.Store, error) {
	var storeDB models.Store

	if result := m.DB.Find(&storeDB, ID); result.Error != nil {
		return nil, result.Error
	}

	var minAmount valueobjects.Amount
	minAmount.NewFromFloat(storeDB.MinAmount)

	store := entities.Store{
		ID:           storeDB.ID,
		Name:         storeDB.Name,
		RewardPoints: storeDB.RewardPoints,
		RewardCoins:  storeDB.RewardCoins,
		MinAmount:    minAmount,
		Status:       valueobjects.Status(storeDB.Status),
	}

	return &store, nil
}
