package usecases

import "lucio.com/order-service/src/domain/dto"

type GetBranchCampaignsUC interface {
	Execute(branchID uint) []dto.BranchCampaignReportDTO
}
