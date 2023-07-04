package usecases

import (
	"errors"
	"time"

	"lucio.com/order-service/src/domain/contracts/repositories"
	"lucio.com/order-service/src/domain/dto"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/vo"
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
		return nil, errors.New("la campaña no existe")
	}

	if branch := c.BranchRepository.FindByID(createBranchCampaignDTO.BranchID); branch == nil {
		return nil, errors.New("la sucursal no existe")
	}

	startDate, err := time.Parse(time.DateTime, createBranchCampaignDTO.StartDate)
	if err != nil {
		return nil, errors.New("el formato de start_date es incorrecto")
	}

	endDate, err := time.Parse(time.DateTime, createBranchCampaignDTO.EndDate)
	if err != nil {
		return nil, errors.New("el formato de end_date es incorrecto")
	}

	operator, err := vo.NewOperator(createBranchCampaignDTO.Operator)
	if err != nil {
		return nil, err
	}

	branchCampaign := entities.BranchCampaign{
		CampaignID:     createBranchCampaignDTO.CampaignID,
		BranchID:       createBranchCampaignDTO.BranchID,
		StartDate:      startDate,
		EndDate:        endDate,
		Operator:       operator,
		OperationValue: createBranchCampaignDTO.OperatorValue,
		MinAmount:      vo.NewAmountFromFloat(createBranchCampaignDTO.MinAmount),
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
