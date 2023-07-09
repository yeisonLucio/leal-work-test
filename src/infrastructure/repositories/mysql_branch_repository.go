package repositories

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/vo"
	"lucio.com/order-service/src/infrastructure/models"
)

type MysqlBranchRepository struct {
	DB     *gorm.DB
	Logger *logrus.Entry
}

func (m *MysqlBranchRepository) Create(branch entities.Branch) (*entities.Branch, error) {
	log := m.Logger.WithFields(logrus.Fields{
		"file":   "mysql_branch_repository",
		"method": "Create",
		"branch": branch,
	})

	branchDB := models.Branch{
		Name:    branch.Name,
		Status:  string(branch.Status),
		StoreID: branch.StoreID,
	}

	if result := m.DB.Create(&branchDB); result.Error != nil {
		log.WithError(result.Error).Error("error creating a branch")
		return nil, result.Error
	}

	branch.ID = branchDB.ID

	return &branch, nil
}

func (m *MysqlBranchRepository) FindByID(ID uint) *entities.Branch {
	var branchDB models.Branch

	if result := m.DB.Find(&branchDB, ID); result.RowsAffected == 0 {
		return nil
	}

	branch := entities.Branch{
		ID:     branchDB.ID,
		Name:   branchDB.Name,
		Status: vo.Status(branchDB.Status),
	}

	return &branch
}

func (m *MysqlBranchRepository) GetIdsByStoreID(storeID uint) []uint {
	var branchIds []uint
	m.DB.Model(&models.Branch{}).Where("store_id = ?", storeID).Pluck("id", &branchIds)

	return branchIds
}
