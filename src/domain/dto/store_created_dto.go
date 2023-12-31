package dto

type StoreCreatedDTO struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	RewardPoints uint    `json:"reward_points"`
	RewardCoins  uint    `json:"reward_coins"`
	MinAmount    float64 `json:"min_amount"`
	Status       string  `json:"status"`
} // @name storeResponse
