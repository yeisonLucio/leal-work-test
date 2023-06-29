package usecases

import "lucio.com/order-service/src/domain/dto"

type CreateUserUC interface {
	Execute(createUserDTO dto.CreateUserDTO) (*dto.UserCreatedDTO, error)
}
