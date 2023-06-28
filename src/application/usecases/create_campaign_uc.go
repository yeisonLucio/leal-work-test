package usecases

import (
	"lucio.com/order-service/src/domain/contracts/repositories"
	"lucio.com/order-service/src/domain/dto"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/valueobjects"
)

type CreateCampaignUC struct {
	CampaignRepository repositories.CampaignRepository
}

func (c *CreateCampaignUC) Execute(createCampaignDTO dto.CreateCampaignDTO) (*dto.CampaignCreatedDTO, error) {
	campaignDB, err := c.CampaignRepository.Create(entities.Campaign{
		Description: createCampaignDTO.Description,
		Status:      valueobjects.ActiveStatus,
	})

	if err != nil {
		return nil, err
	}

	response := dto.CampaignCreatedDTO{
		ID:          campaignDB.ID,
		Description: campaignDB.Description,
		Status:      string(campaignDB.Status),
	}

	return &response, nil
}
