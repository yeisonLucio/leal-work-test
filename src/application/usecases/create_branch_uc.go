package usecases

import (
	"errors"

	"lucio.com/order-service/src/domain/contracts/repositories"
	"lucio.com/order-service/src/domain/dto"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/vo"
)

type CreateBranchUC struct {
	BranchRepository repositories.BranchRepository
	StoreRepository  repositories.StoreRepository
}

func (c *CreateBranchUC) Execute(createBranchDTO dto.CreateBranchDTO) (*dto.BranchCreatedDTO, error) {

	if store := c.StoreRepository.FindByID(createBranchDTO.StoreID); store == nil {
		return nil, errors.New("la tienda no existe")
	}

	branch := entities.Branch{
		Name:    createBranchDTO.Name,
		StoreID: createBranchDTO.StoreID,
		Status:  vo.ActiveStatus,
	}

	branchCreated, err := c.BranchRepository.Create(branch)
	if err != nil {
		return nil, err
	}

	response := dto.BranchCreatedDTO{
		ID:      branchCreated.ID,
		Name:    branchCreated.Name,
		Status:  string(branchCreated.Status),
		StoreID: branchCreated.StoreID,
	}

	return &response, nil

}
