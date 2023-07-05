package usecases

type CalculateCampaignRewardsUC interface {
	Execute(brachID, storePoints, storeCoins uint, transactionAmount float64) (uint, uint)
}
