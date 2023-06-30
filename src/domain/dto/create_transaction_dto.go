package dto

type CreateTransactionDTO struct {
	UserID   uint    `json:"user_id"`
	BranchID uint    `json:"branch_id"`
	Amount   float64 `json:"amount"`
}
