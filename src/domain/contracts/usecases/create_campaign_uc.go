package usecases

import "lucio.com/order-service/src/domain/dto"

type CreateCampaignUC interface {
	Execute(createCampaignDTO dto.CreateCampaignDTO) (*dto.CampaignCreatedDTO, error)
}
