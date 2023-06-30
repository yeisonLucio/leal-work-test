package repositories

import (
	"gorm.io/gorm"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/infrastructure/models"
)

type MysqlTransactionRepository struct {
	DB *gorm.DB
}

func (m *MysqlTransactionRepository) Create(
	transaction entities.Transaction,
) (*entities.Transaction, error) {
	transactionDB := models.Transaction{
		UserID:   transaction.UserID,
		BranchID: transaction.BranchID,
		Amount:   transaction.Amount.GetValue(),
		Points:   transaction.Points,
		Coins:    transaction.Coins,
		Type:     string(transaction.Type),
	}

	if result := m.DB.Create(&transactionDB); result.Error != nil {
		return nil, result.Error
	}

	transaction.ID = transactionDB.ID

	return &transaction, nil
}
