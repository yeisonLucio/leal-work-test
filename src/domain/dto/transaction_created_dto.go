package dto

type TransactionCreatedDTO struct {
	ID       uint    `json:"id"`
	UserID   uint    `json:"user_id"`
	BranchID uint    `json:"branch_id"`
	Amount   float64 `json:"amount"`
	Points   uint    `json:"points"`
	Coins    uint    `json:"coins"`
	Type     string  `json:"type"`
} // @name transactionResponse
