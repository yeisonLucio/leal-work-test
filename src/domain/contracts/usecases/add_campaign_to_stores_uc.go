package usecases

import "lucio.com/order-service/src/domain/dto"

type AddCampaignToStoreUC interface {
	Execute(createStoreCampaignDTO dto.CreateStoreCampaignDTO) (*dto.StoreCampaignCreatedDTO, error)
}
