package dto

type BranchCampaignReportDTO struct {
	ID            uint    `json:"id"`
	BranchID      uint    `json:"branch_id"`
	CampaignID    uint    `json:"campaign_id"`
	Description   string  `json:"description"`
	Status        string  `json:"status"`
	StartDate     string  `json:"start_date"`
	EndDate       string  `json:"end_date"`
	Operator      string  `json:"operator"`
	OperatorValue uint    `json:"operator_value"`
	MinAmount     float64 `json:"min_amount"`
}
