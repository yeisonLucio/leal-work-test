package entities

import "lucio.com/order-service/src/domain/valueobjects"

type Transaction struct {
	ID       uint
	UserID   uint
	BranchID uint
	Amount   valueobjects.Amount
	Points   uint
	Coins    uint
	Type     valueobjects.TransactionType
}
