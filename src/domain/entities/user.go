package entities

import "lucio.com/order-service/src/domain/vo"

type User struct {
	ID             uint
	Name           string
	Identification string
	Status         vo.Status
}
