package entities

import "lucio.com/order-service/src/domain/vo"

type Transaction struct {
	ID       uint
	UserID   uint
	BranchID uint
	Amount   vo.Amount
	Points   uint
	Coins    uint
	Type     vo.TransactionType
}
