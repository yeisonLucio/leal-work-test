package dto

type StoreCampaignCreatedDTO struct {
	BranchCampaigns []BranchCampaignCreatedDTO `json:"branch_campaigns"`
	Errors          []ErroBranchCampaign       `json:"errors"`
} // @name storeCampaignResponse

type ErroBranchCampaign struct {
	Message  string `json:"message"`
	BranchId uint   `json:"branch_id"`
}
