package entities

import "lucio.com/order-service/src/domain/valueobjects"

type Store struct {
	ID     uint
	Name   string
	Status valueobjects.Status
}
