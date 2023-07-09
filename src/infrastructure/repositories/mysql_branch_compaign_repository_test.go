package repositories

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"lucio.com/order-service/src/config/database"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/vo"
	"lucio.com/order-service/src/infrastructure/models"
)

func TestMysqlBranchCampaignRepository_Create(t *testing.T) {
	db := database.GetTestDB(t)

	operator, _ := vo.NewOperator("%")
	amount := vo.NewAmountFromFloat(1000)

	branchCampaign := entities.BranchCampaign{
		BranchID:      1,
		CampaignID:    1,
		StartDate:     time.Now(),
		EndDate:       time.Now(),
		Operator:      operator,
		OperatorValue: 3,
		MinAmount:     amount,
	}

	repo := MysqlBranchCampaignRepository{
		DB: db,
	}

	result, err := repo.Create(branchCampaign)
	assert.NoError(t, err)
	assert.NotEmpty(t, result.ID)
}

func TestMysqlBranchCampaignRepository_FindByID(t *testing.T) {
	db := database.GetTestDB(t)

	branchCampaign := models.BranchCampaign{
		CampaignID: 1,
		BranchID:   1,
		MinAmount:  1000,
	}

	db.Create(&branchCampaign)

	repo := MysqlBranchCampaignRepository{
		DB: db,
	}

	result := repo.FindByID(branchCampaign.ID)

	assert.NotNil(t, result)
	assert.Equal(t, result.ID, branchCampaign.ID)
}
