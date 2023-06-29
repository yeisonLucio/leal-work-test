package dto

type CreateStoreCampaignDTO struct {
	CampaignID    uint    `json:"campaign_id"`
	StoreID       uint    `json:"store_id"`
	StartDate     string  `json:"start_date"`
	EndDate       string  `json:"end_date"`
	Operator      string  `json:"operator"`
	OperatorValue uint    `json:"operator_value"`
	MinAmount     float64 `json:"min_amount"`
}
