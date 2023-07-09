package repositories

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"lucio.com/order-service/src/domain/dto"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/vo"
	"lucio.com/order-service/src/infrastructure/models"
)

type MysqlBranchCampaignRepository struct {
	DB     *gorm.DB
	Logger *logrus.Entry
}

func (m *MysqlBranchCampaignRepository) Create(
	branchCampaign entities.BranchCampaign,
) (*entities.BranchCampaign, error) {
	log := m.Logger.WithFields(logrus.Fields{
		"file":           "mysql_branch_campaign_repository",
		"method":         "Create",
		"branchCampaign": branchCampaign,
	})

	branchDB := models.BranchCampaign{
		BranchID:      branchCampaign.BranchID,
		CampaignID:    branchCampaign.CampaignID,
		StartDate:     branchCampaign.StartDate,
		EndDate:       branchCampaign.EndDate,
		Operator:      branchCampaign.Operator.Value(),
		OperatorValue: branchCampaign.OperatorValue,
		MinAmount:     branchCampaign.MinAmount.Value(),
	}

	if result := m.DB.Create(&branchDB); result.Error != nil {
		log.WithError(result.Error).Error("error creating branchCampaign")
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

	operator, _ := vo.NewOperator(branchDB.Operator)

	branchCampaign := entities.BranchCampaign{
		ID:            branchDB.ID,
		CampaignID:    branchDB.CampaignID,
		BranchID:      branchDB.BranchID,
		StartDate:     branchDB.StartDate,
		EndDate:       branchDB.EndDate,
		Operator:      operator,
		OperatorValue: branchDB.OperatorValue,
		MinAmount:     vo.NewAmountFromFloat(branchDB.MinAmount),
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
