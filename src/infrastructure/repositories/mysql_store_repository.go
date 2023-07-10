package repositories

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/vo"
	"lucio.com/order-service/src/infrastructure/models"
)

type MysqlStoreRepository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func (m *MysqlStoreRepository) Create(store entities.Store) (*entities.Store, error) {
	log := m.Logger.WithFields(logrus.Fields{
		"file":   "mysql_store_repository",
		"method": "Create",
		"store":  store,
	})

	storeDB := models.Store{
		Name:         store.Name,
		Status:       string(store.Status),
		RewardPoints: store.RewardPoints,
		RewardCoins:  store.RewardCoins,
		MinAmount:    store.MinAmount.Value(),
	}

	if result := m.DB.Create(&storeDB); result.Error != nil {
		log.WithError(result.Error).Error("error creating store")
		return nil, result.Error
	}

	store.ID = storeDB.ID

	return &store, nil
}

func (m *MysqlStoreRepository) FindByID(ID uint) *entities.Store {
	var storeDB models.Store

	if result := m.DB.Find(&storeDB, ID); result.RowsAffected == 0 {
		return nil
	}

	store := entities.Store{
		ID:           storeDB.ID,
		Name:         storeDB.Name,
		RewardPoints: storeDB.RewardPoints,
		RewardCoins:  storeDB.RewardCoins,
		MinAmount:    vo.NewAmountFromFloat(storeDB.MinAmount),
		Status:       vo.Status(storeDB.Status),
	}

	return &store
}

func (m *MysqlStoreRepository) FindByBranchID(branchID uint) *entities.Store {
	var storeDB models.Store

	m.DB.Table("stores s").
		Select(`s.id, 
			s.name, 
			s.min_amount,
			s.reward_points, 
			s.reward_coins`).
		Joins("INNER JOIN branches b ON b.store_id = s.id").
		Where("b.id", branchID).
		First(&storeDB)

	store := entities.Store{
		ID:           storeDB.ID,
		Name:         storeDB.Name,
		RewardPoints: storeDB.RewardPoints,
		RewardCoins:  storeDB.RewardCoins,
		MinAmount:    vo.NewAmountFromFloat(storeDB.MinAmount),
		Status:       vo.Status(storeDB.Status),
	}

	return &store
}
