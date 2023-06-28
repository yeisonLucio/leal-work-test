package dto

type CreateStoreDTO struct {
	Name         string `json:"name"`
	RewardPoints uint   `json:"reward_points"`
	RewardCoins  uint   `json:"reward_coins"`
	MinAmount    string `json:"min_amount"`
}
