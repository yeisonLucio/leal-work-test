package dto

type CreateStoreDTO struct {
	Name         string  `json:"name"`
	RewardPoints uint    `json:"reward_points"`
	RewardCoins  uint    `json:"reward_coins"`
	MinAmount    float64 `json:"min_amount"`
} // @name storeRequest
