package entities

import "lucio.com/order-service/src/domain/valueobjects"

type Reward struct {
	ID          uint
	Reward      string
	Description string
	MinAmount   float64
	AmountType  string
	StoreID     uint
	Status      valueobjects.Status
}
