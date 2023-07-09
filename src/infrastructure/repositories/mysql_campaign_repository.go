package repositories

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/vo"
	"lucio.com/order-service/src/infrastructure/models"
)

type MysqlCampaignRepository struct {
	DB     *gorm.DB
	Logger *logrus.Entry
}

func (m *MysqlCampaignRepository) Create(campaign entities.Campaign) (*entities.Campaign, error) {
	log := m.Logger.WithFields(logrus.Fields{
		"file":     "mysql_campaign_repository",
		"method":   "Create",
		"campaign": campaign,
	})

	campaignDB := models.Campaign{
		Description: campaign.Description,
		Status:      string(campaign.Status),
	}

	if result := m.DB.Create(&campaignDB); result.Error != nil {
		log.WithError(result.Error).Error("error creating a campaign")
		return nil, result.Error
	}

	campaign.ID = campaignDB.ID

	return &campaign, nil
}

func (m *MysqlCampaignRepository) FindByID(ID uint) *entities.Campaign {
	var campaignDB models.Campaign

	result := m.DB.Where("status=?", vo.ActiveStatus).Find(&campaignDB, ID)
	if result.RowsAffected == 0 {
		return nil
	}

	Campaign := entities.Campaign{
		ID:          campaignDB.ID,
		Description: campaignDB.Description,
		Status:      vo.Status(campaignDB.Status),
	}

	return &Campaign
}
