package repositories

import (
	"time"

	"gorm.io/gorm"
	"lucio.com/order-service/src/domain/dto"
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
		EndDate:       branchCampaign.EndDate,
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

func (m *MysqlBranchCampaignRepository) FindByID(ID uint) *entities.BranchCampaign {
	var branchDB models.BranchCampaign

	if result := m.DB.Find(&branchDB, ID); result.RowsAffected == 0 {
		return nil
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

	return &branchCampaign
}

func (m *MysqlBranchCampaignRepository) FindByBranchID(
	branchID uint,
) []dto.BranchCampaignReportDTO {
	var branchCampaigns []dto.BranchCampaignReportDTO

	m.DB.Table("branch_campaigns bc").
		Select(`bc.branch_id,
			bc.campaign_id,
			c.description,
			bc.start_date,
			bc.end_date,
			bc.operator,
			bc.operator_value,
			bc.min_amount,
			c.status`).
		Joins("INNER JOIN campaigns c ON c.id=bc.campaign_id").
		Where("branch_id", branchID).Find(&branchCampaigns).Scan(&branchCampaigns)

	return branchCampaigns
}

func (m *MysqlBranchCampaignRepository) GetActivesByBranchID(
	branchID uint,
) []dto.BranchCampaignCreatedDTO {
	var branchCampaigns []dto.BranchCampaignCreatedDTO

	now := time.Now().Format(time.DateTime)

	m.DB.Table("branch_campaigns").
		Where("branch_id", branchID).
		Where("end_date > ?", now).
		Find(&branchCampaigns)

	return branchCampaigns
}
