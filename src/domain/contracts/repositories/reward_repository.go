package repositories

import "lucio.com/order-service/src/domain/entities"

type RewardRepository interface {
	Create(reward entities.Reward) (*entities.Reward, error)
}
