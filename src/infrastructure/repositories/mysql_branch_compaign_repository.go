package repositories

import (
	"gorm.io/gorm"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/valueobjects"
	"lucio.com/order-service/src/infrastructure/models"
)

type MysqlBranchCampaignRepository struct {
	DB *gorm.DB
}

func (m *MysqlBranchCampaignRepository) Create(
	branchCampaign entities.BranchCampaign,
) (*entities.BranchCampaign, error) {
	branchDB := models.BranchCampaign{
		BranchID:      branchCampaign.BranchID,
		CampaignID:    branchCampaign.CampaignID,
		StartDate:     branchCampaign.StartDate,
		EndDate:       branchCampaign.StartDate,
		Operator:      branchCampaign.Operator.GetValue(),
		OperatorValue: branchCampaign.OperationValue,
		MinAmount:     branchCampaign.MinAmount.GetValue(),
	}

	if result := m.DB.Create(&branchDB); result.Error != nil {
		return nil, result.Error
	}

	branchCampaign.ID = branchDB.ID

	return &branchCampaign, nil
}

func (m *MysqlBranchCampaignRepository) FindByID(ID uint) (*entities.BranchCampaign, error) {
	var branchDB models.BranchCampaign

	if result := m.DB.Find(&branchDB, ID); result.Error != nil {
		return nil, result.Error
	}

	var operator valueobjects.Operator
	operator.New(branchDB.Operator)

	var minAmount valueobjects.Amount
	minAmount.NewFromFloat(branchDB.MinAmount)

	branchCampaign := entities.BranchCampaign{
		ID:             branchDB.ID,
		CampaignID:     branchDB.CampaignID,
		BranchID:       branchDB.BranchID,
		StartDate:      branchDB.StartDate,
		EndDate:        branchDB.EndDate,
		Operator:       operator,
		OperationValue: branchDB.OperatorValue,
		MinAmount:      minAmount,
	}

	return &branchCampaign, nil
}
