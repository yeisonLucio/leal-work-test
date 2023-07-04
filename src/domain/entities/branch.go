package entities

import "lucio.com/order-service/src/domain/vo"

type Branch struct {
	ID      uint
	Name    string
	Status  vo.Status
	StoreID uint
}
