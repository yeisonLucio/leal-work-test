package entities

import (
	"lucio.com/order-service/src/domain/valueobjects"
)

type Campaign struct {
	ID          uint
	Description string
	Status      valueobjects.Status
}
