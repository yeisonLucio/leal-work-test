package usecases

import "lucio.com/order-service/src/domain/dto"

type CreateBranchUC interface {
	Execute(createBranchDTO dto.CreateBranchDTO) (*dto.BranchCreatedDTO, error)
}
