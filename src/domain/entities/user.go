package entities

import "lucio.com/order-service/src/domain/valueobjects"

type User struct {
	ID             uint
	Name           string
	Identification string
	Status         valueobjects.Status
}
