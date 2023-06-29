package repositories

import "lucio.com/order-service/src/domain/entities"

type UserRepository interface {
	Create(user entities.User) (*entities.User, error)
	FindByID(ID uint) *entities.User
}
