package dto

type CreateTransactionDTO struct {
	UserID   uint    `json:"user_id" swaggerignore:"true"`
	BranchID uint    `json:"branch_id" swaggerignore:"true"`
	Amount   float64 `json:"amount"`
} // @name transactionRequest
