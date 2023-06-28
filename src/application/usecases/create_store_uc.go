package usecases

import (
	"lucio.com/order-service/src/domain/contracts/repositories"
	"lucio.com/order-service/src/domain/dto"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/valueobjects"
)

type CreateStoreUC struct {
	StoreRepository repositories.StoreRepository
}

func (c *CreateStoreUC) Execute(createStoreDTO dto.CreateStoreDTO) (*dto.StoreCreatedDTO, error) {
	storeDB, err := c.StoreRepository.Create(entities.Store{
		Name:   createStoreDTO.Name,
		Status: valueobjects.ActiveStatus,
	})

	if err != nil {
		return nil, err
	}

	response := dto.StoreCreatedDTO{
		ID:     storeDB.ID,
		Name:   storeDB.Name,
		Status: string(storeDB.Status),
	}

	return &response, nil
}
