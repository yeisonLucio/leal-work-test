package entities

import "lucio.com/order-service/src/domain/vo"

type Campaign struct {
	ID          uint
	Description string
	Status      vo.Status
}
