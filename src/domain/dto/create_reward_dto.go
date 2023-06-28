package dto

type CreateRewardDTO struct {
	Reward      string `json:"reward"`
	Description string `json:"description"`
	MinAmount   string `json:"min_amount"`
	AmountType  string `json:"amount_type"`
	StoreID     uint   `json:"store_id"`
}
