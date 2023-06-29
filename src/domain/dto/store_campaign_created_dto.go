package dto

type StoreCampaignCreatedDTO struct {
	BranchCampaigns []BranchCampaignCreatedDTO `json:"branch_campaigns"`
	Errors          []ErroBranchCampaign       `json:"errors"`
}

type ErroBranchCampaign struct {
	Message  string `json:"message"`
	BranchId uint   `json:"branch_id"`
}
