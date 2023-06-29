package usecases

import (
	"errors"
	"time"

	"lucio.com/order-service/src/domain/contracts/repositories"
	"lucio.com/order-service/src/domain/dto"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/valueobjects"
)

type CreateBranchCampaignUC struct {
	BranchRepository         repositories.BranchRepository
	CampaignRepository       repositories.CampaignRepository
	BranchCampaignRepository repositories.BranchCampaignRepository
}

func (c *CreateBranchCampaignUC) Execute(
	createBranchCampaignDTO dto.CreateBranchCampaignDTO,
) (*dto.BranchCampaignCreatedDTO, error) {

	if _, err := c.CampaignRepository.FindByID(createBranchCampaignDTO.CampaignID); err != nil {
		return nil, errors.New("el campaign_id no existe")
	}

	if _, err := c.BranchRepository.FindByID(createBranchCampaignDTO.BranchID); err != nil {
		return nil, errors.New("el branch_id no existe")
	}

	startDate, err := time.Parse(time.DateTime, createBranchCampaignDTO.StartDate)
	if err != nil {
		return nil, errors.New("el formato de start_date es incorrecto")
	}

	endDate, err := time.Parse(time.DateTime, createBranchCampaignDTO.EndDate)
	if err != nil {
		return nil, errors.New("el formato de end_date es incorrecto")
	}

	var operator valueobjects.Operator
	if err := operator.New(createBranchCampaignDTO.Operator); err != nil {
		return nil, err
	}

	var minAmount valueobjects.Amount
	minAmount.NewFromFloat(createBranchCampaignDTO.MinAmount)

	branchCampaign := entities.BranchCampaign{
		CampaignID:     createBranchCampaignDTO.CampaignID,
		BranchID:       createBranchCampaignDTO.BranchID,
		StartDate:      startDate,
		EndDate:        endDate,
		Operator:       operator,
		OperationValue: createBranchCampaignDTO.OperatorValue,
		MinAmount:      minAmount,
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
		Operator:      createdBranchCampaign.Operator.GetValue(),
		OperatorValue: createdBranchCampaign.OperationValue,
		MinAmount:     createdBranchCampaign.MinAmount.GetValue(),
	}

	return &response, nil
}
