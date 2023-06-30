package dto

type CreateBranchDTO struct {
	Name    string `json:"name"`
	StoreID uint   `json:"store_id" swaggerignore:"true"`
} //@name branchRequest
