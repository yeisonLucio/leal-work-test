package usecases

import (
	"lucio.com/order-service/src/domain/contracts/repositories"
	"lucio.com/order-service/src/domain/dto"
)

type CalculateCampaignRewardsUC struct {
	BranchCampaignRepository repositories.BranchCampaignRepository
}

func (uc *CalculateCampaignRewardsUC) Execute(
	branchID uint,
	storePoints uint,
	storeCoins uint,
	transactionAmount float64,
) (uint, uint) {
	campaigns := uc.BranchCampaignRepository.GetActivesByBranchID(branchID)

	var rewardPoints, rewardCoins uint

	for _, campaign := range campaigns {
		if transactionAmount < campaign.MinAmount {
			continue
		}

		points, coins := uc.calculateRewards(
			campaign,
			storePoints,
			storeCoins,
		)

		rewardPoints += points
		rewardCoins += coins
	}

	return rewardPoints, rewardCoins
}

func (uc *CalculateCampaignRewardsUC) calculateRewards(
	campaign dto.BranchCampaignCreatedDTO,
	storePoints uint,
	storeCoins uint,
) (uint, uint) {
	switch campaign.Operator {
	case "%":
		storePoints = (storePoints * campaign.OperatorValue) / 100
		storeCoins = (storeCoins * campaign.OperatorValue) / 100

	case "*":
		storeCoins = (storeCoins * campaign.OperatorValue) - storeCoins
		storePoints = (storePoints * campaign.OperatorValue) - storePoints
	}

	return storePoints, storeCoins
}
