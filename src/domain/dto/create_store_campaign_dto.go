package dto

type CreateStoreCampaignDTO struct {
	CampaignID    uint    `json:"campaign_id" swaggerignore:"true"`
	StoreID       uint    `json:"store_id" swaggerignore:"true"`
	StartDate     string  `json:"start_date"`
	EndDate       string  `json:"end_date"`
	Operator      string  `json:"operator"`
	OperatorValue uint    `json:"operator_value"`
	MinAmount     float64 `json:"min_amount"`
} //@name storeCampaignRequest
