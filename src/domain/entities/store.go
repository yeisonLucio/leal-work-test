package entities

import "lucio.com/order-service/src/domain/valueobjects"

type Store struct {
	ID           uint
	Name         string
	RewardPoints uint
	RewardCoins  uint
	MinAmount    valueobjects.Amount
	Status       valueobjects.Status
}
