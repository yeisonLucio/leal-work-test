package entities

import "lucio.com/order-service/src/domain/valueobjects"

type Reward struct {
	ID          uint
	Reward      string
	Description string
	MinAmount   valueobjects.Amount
	AmountType  valueobjects.AmountType
	StoreID     uint
	Status      valueobjects.Status
}
