package usecases

import (
	"github.com/sirupsen/logrus"
	"lucio.com/order-service/src/domain/contracts/repositories"
	"lucio.com/order-service/src/domain/dto"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/vo"
)

type CreateCampaignUC struct {
	CampaignRepository repositories.CampaignRepository
	Logger             *logrus.Logger
}

func (c *CreateCampaignUC) Execute(createCampaignDTO dto.CreateCampaignDTO) (*dto.CampaignCreatedDTO, error) {
	log := c.Logger.WithFields(logrus.Fields{
		"file":              "create_campaign_uc",
		"method":            "Execute",
		"createCampaignDTO": createCampaignDTO,
	})

	campaignDB, err := c.CampaignRepository.Create(entities.Campaign{
		Description: createCampaignDTO.Description,
		Status:      vo.ActiveStatus,
	})

	if err != nil {
		log.WithError(err).Error("erro creating campaign")
		return nil, err
	}

	response := dto.CampaignCreatedDTO{
		ID:          campaignDB.ID,
		Description: campaignDB.Description,
		Status:      string(campaignDB.Status),
	}

	return &response, nil
}
