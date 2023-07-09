package usecases

import (
	"errors"

	"github.com/sirupsen/logrus"
	"lucio.com/order-service/src/domain/contracts/repositories"
	"lucio.com/order-service/src/domain/dto"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/vo"
)

var (
	errStoreNotFound = errors.New("tienda no encontrada")
)

type CreateBranchUC struct {
	BranchRepository repositories.BranchRepository
	StoreRepository  repositories.StoreRepository
	Logger           *logrus.Entry
}

func (c *CreateBranchUC) Execute(
	createBranchDTO dto.CreateBranchDTO,
) (*dto.BranchCreatedDTO, error) {

	log := c.Logger.WithFields(logrus.Fields{
		"file":            "create_branch_uc",
		"method":          "Execute",
		"createBranchDTO": createBranchDTO,
	})

	if store := c.StoreRepository.FindByID(createBranchDTO.StoreID); store == nil {
		log.WithError(errStoreNotFound).Error("store not found")
		return nil, errStoreNotFound
	}

	branch := entities.Branch{
		Name:    createBranchDTO.Name,
		StoreID: createBranchDTO.StoreID,
		Status:  vo.ActiveStatus,
	}

	branchCreated, err := c.BranchRepository.Create(branch)
	if err != nil {
		log.WithError(err).Error("error creating branch")
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
