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
	errCampaignNotFound = errors.New("campaña no encontrada")
	errBranchNotFound   = errors.New("sucursal no encontrada")
)

type CreateBranchCampaignUC struct {
	BranchRepository         repositories.BranchRepository
	CampaignRepository       repositories.CampaignRepository
	BranchCampaignRepository repositories.BranchCampaignRepository
	Logger                   *logrus.Logger
}

func (c *CreateBranchCampaignUC) Execute(
	createBranchCampaignDTO dto.CreateBranchCampaignDTO,
) (*dto.BranchCampaignCreatedDTO, error) {

	log := c.Logger.WithFields(logrus.Fields{
		"file":                    "create_branch_campaign_uc",
		"method":                  "Execute",
		"createBranchCampaignDTO": createBranchCampaignDTO,
	})

	if campaign := c.CampaignRepository.FindByID(createBranchCampaignDTO.CampaignID); campaign == nil {
		log.WithError(errCampaignNotFound).Error("campaign not found")
		return nil, errCampaignNotFound
	}

	if branch := c.BranchRepository.FindByID(createBranchCampaignDTO.BranchID); branch == nil {
		log.WithError(errBranchNotFound).Error("branch not found")
		return nil, errBranchNotFound
	}

	branchCampaign := entities.BranchCampaign{
		CampaignID:    createBranchCampaignDTO.CampaignID,
		BranchID:      createBranchCampaignDTO.BranchID,
		OperatorValue: createBranchCampaignDTO.OperatorValue,
		MinAmount:     vo.NewAmountFromFloat(createBranchCampaignDTO.MinAmount),
	}

	if err := branchCampaign.SetStartDate(createBranchCampaignDTO.StartDate); err != nil {
		log.WithError(err).Error("start date not valid")
		return nil, err
	}

	if err := branchCampaign.SetEndDate(createBranchCampaignDTO.EndDate); err != nil {
		log.WithError(err).Error("end date not valid")
		return nil, err
	}

	if err := branchCampaign.SetOperator(createBranchCampaignDTO.Operator); err != nil {
		log.WithError(err).Error("operator not valid")
		return nil, err
	}

	createdBranchCampaign, err := c.BranchCampaignRepository.Create(branchCampaign)
	if err != nil {
		log.WithError(err).Error("error creating branch campaign")
		return nil, err
	}

	response := dto.BranchCampaignCreatedDTO{
		ID:            createdBranchCampaign.ID,
		BranchID:      createdBranchCampaign.BranchID,
		CampaignID:    createdBranchCampaign.CampaignID,
		StartDate:     createdBranchCampaign.StartDate.String(),
		EndDate:       createdBranchCampaign.EndDate.String(),
		Operator:      createdBranchCampaign.Operator.Value(),
		OperatorValue: createdBranchCampaign.OperatorValue,
		MinAmount:     createdBranchCampaign.MinAmount.Value(),
	}

	return &response, nil
}
