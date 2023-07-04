package entities

import "lucio.com/order-service/src/domain/vo"

type Store struct {
	ID           uint
	Name         string
	RewardPoints uint
	RewardCoins  uint
	MinAmount    vo.Amount
	Status       vo.Status
}
