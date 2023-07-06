package usecases

import (
	"errors"

	"lucio.com/order-service/src/domain/contracts/repositories"
	"lucio.com/order-service/src/domain/dto"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/vo"
)

var (
	errCampaignNotFound = errors.New("campaña no encontrada")
	errBranchNotFound   = errors.New("sucursal no encontrada")
)

type CreateBranchCampaignUC struct {
	BranchRepository         repositories.BranchRepository
	CampaignRepository       repositories.CampaignRepository
	BranchCampaignRepository repositories.BranchCampaignRepository
}

func (c *CreateBranchCampaignUC) Execute(
	createBranchCampaignDTO dto.CreateBranchCampaignDTO,
) (*dto.BranchCampaignCreatedDTO, error) {

	if campaign := c.CampaignRepository.FindByID(createBranchCampaignDTO.CampaignID); campaign == nil {
		return nil, errCampaignNotFound
	}

	if branch := c.BranchRepository.FindByID(createBranchCampaignDTO.BranchID); branch == nil {
		return nil, errBranchNotFound
	}

	branchCampaign := entities.BranchCampaign{
		CampaignID:     createBranchCampaignDTO.CampaignID,
		BranchID:       createBranchCampaignDTO.BranchID,
		OperationValue: createBranchCampaignDTO.OperatorValue,
		MinAmount:      vo.NewAmountFromFloat(createBranchCampaignDTO.MinAmount),
	}

	if err := branchCampaign.SetStartDate(createBranchCampaignDTO.StartDate); err != nil {
		return nil, err
	}

	if err := branchCampaign.SetEndDate(createBranchCampaignDTO.EndDate); err != nil {
		return nil, err
	}

	if err := branchCampaign.SetOperator(createBranchCampaignDTO.Operator); err != nil {
		return nil, err
	}

	createdBranchCampaign, err := c.BranchCampaignRepository.Create(branchCampaign)
	if err != nil {
		return nil, err
	}

	response := dto.BranchCampaignCreatedDTO{
		ID:            createdBranchCampaign.ID,
		BranchID:      createdBranchCampaign.BranchID,
		CampaignID:    createdBranchCampaign.CampaignID,
		StartDate:     createdBranchCampaign.StartDate.String(),
		EndDate:       createdBranchCampaign.EndDate.String(),
		Operator:      createdBranchCampaign.Operator.Value(),
		OperatorValue: createdBranchCampaign.OperationValue,
		MinAmount:     createdBranchCampaign.MinAmount.Value(),
	}

	return &response, nil
}
