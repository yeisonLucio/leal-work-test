package usecases

import (
	"lucio.com/order-service/src/domain/dto"
)

type CreateStoreUC interface {
	Execute(createStoreDTO dto.CreateStoreDTO) (*dto.StoreCreatedDTO, error)
}
