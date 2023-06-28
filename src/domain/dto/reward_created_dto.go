package dto

type RewardCreatedDTO struct {
	ID          uint    `json:"id"`
	Reward      string  `json:"reward"`
	Description string  `json:"description"`
	MinAmount   float64 `json:"min_amount"`
	AmountType  string  `json:"amount_type"`
	StoreID     uint    `json:"store_id"`
	Status      string  `json:"status"`
}
