package usecases

import "lucio.com/order-service/src/domain/dto"

type CreateBranchCampaignUC interface {
	Execute(createBranchCampaignDTO dto.CreateBranchCampaignDTO) (*dto.BranchCampaignCreatedDTO, error)
}
