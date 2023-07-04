package entities

import "lucio.com/order-service/src/domain/vo"

type Reward struct {
	ID          uint
	Reward      string
	Description string
	MinAmount   vo.Amount
	AmountType  vo.AmountType
	StoreID     uint
	Status      vo.Status
}
