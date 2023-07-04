package usecases

import (
	"lucio.com/order-service/src/domain/contracts/repositories"
	"lucio.com/order-service/src/domain/dto"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/vo"
)

type CreateUserUC struct {
	UserRepository repositories.UserRepository
}

func (c *CreateUserUC) Execute(createUserDTO dto.CreateUserDTO) (*dto.UserCreatedDTO, error) {
	createdUser, err := c.UserRepository.Create(entities.User{
		Name:           createUserDTO.Name,
		Identification: createUserDTO.Identification,
		Status:         vo.ActiveStatus,
	})

	if err != nil {
		return nil, err
	}

	response := dto.UserCreatedDTO{
		ID:             createdUser.ID,
		Name:           createdUser.Name,
		Identification: createUserDTO.Identification,
		Status:         string(createdUser.Status),
	}

	return &response, nil
}
