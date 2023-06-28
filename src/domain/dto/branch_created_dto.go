package dto

type BranchCreatedDTO struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	StoreID uint   `json:"store_id"`
}
