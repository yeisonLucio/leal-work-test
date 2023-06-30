package usecases

import (
	"encoding/json"
	"fmt"

	"lucio.com/order-service/src/domain/contracts/repositories"
	"lucio.com/order-service/src/domain/dto"
)

type GetBranchCampaignsUC struct {
	BranchCampaignRepository repositories.BranchCampaignRepository
	CacheRepository          repositories.CacheRepository
}

func (g *GetBranchCampaignsUC) Execute(branchID uint) []dto.BranchCampaignReportDTO {
	var branchCampaignsDTO []dto.BranchCampaignReportDTO
	key := fmt.Sprintf("branch-campaign-report-%d", branchID)

	result, err := g.CacheRepository.GetByKey(key)
	if err == nil {
		if err := json.Unmarshal([]byte(result), &branchCampaignsDTO); err == nil {
			return branchCampaignsDTO
		}
	}

	branchCampaignsDTO = g.BranchCampaignRepository.FindByBranchID(branchID)

	if len(branchCampaignsDTO) > 0 {
		object, _ := json.Marshal(branchCampaignsDTO)
		g.CacheRepository.SetByKey(key, string(object))
	}

	return branchCampaignsDTO
}
