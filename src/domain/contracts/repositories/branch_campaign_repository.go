package repositories

import (
	"lucio.com/order-service/src/domain/dto"
	"lucio.com/order-service/src/domain/entities"
)

type BranchCampaignRepository interface {
	Create(branchCampaign entities.BranchCampaign) (*entities.BranchCampaign, error)
	FindByID(ID uint) *entities.BranchCampaign
	FindByBranchID(branchID uint) []dto.BranchCampaignReportDTO
	GetActivesByBranchID(branchID uint) []dto.BranchCampaignCreatedDTO
}
