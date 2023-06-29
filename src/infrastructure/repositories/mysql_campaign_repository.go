package repositories

import (
	"gorm.io/gorm"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/valueobjects"
	"lucio.com/order-service/src/infrastructure/models"
)

type MysqlCampaignRepository struct {
	DB *gorm.DB
}

func (m *MysqlCampaignRepository) Create(Campaign entities.Campaign) (*entities.Campaign, error) {
	campaignDB := models.Campaign{
		Description: Campaign.Description,
		Status:      string(Campaign.Status),
	}

	if result := m.DB.Create(&campaignDB); result.Error != nil {
		return nil, result.Error
	}

	Campaign.ID = campaignDB.ID

	return &Campaign, nil
}

func (m *MysqlCampaignRepository) FindByID(ID uint) *entities.Campaign {
	var campaignDB models.Campaign

	if result := m.DB.Find(&campaignDB, ID); result.RowsAffected == 0 {
		return nil
	}

	Campaign := entities.Campaign{
		ID:          campaignDB.ID,
		Description: campaignDB.Description,
		Status:      valueobjects.Status(campaignDB.Status),
	}

	return &Campaign
}
