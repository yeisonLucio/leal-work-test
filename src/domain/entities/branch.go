package entities

import "lucio.com/order-service/src/domain/valueobjects"

type Branch struct {
	ID      uint
	Name    string
	Status  valueobjects.Status
	StoreID uint
}
