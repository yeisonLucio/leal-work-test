package repositories

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/infrastructure/models"
)

type MysqlTransactionRepository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func (m *MysqlTransactionRepository) Create(
	transaction entities.Transaction,
) (*entities.Transaction, error) {
	log := m.Logger.WithFields(logrus.Fields{
		"file":        "mysql_transaction_repository",
		"method":      "Create",
		"transaction": transaction,
	})

	transactionDB := models.Transaction{
		UserID:   transaction.UserID,
		BranchID: transaction.BranchID,
		Amount:   transaction.Amount.Value(),
		Points:   transaction.Points,
		Coins:    transaction.Coins,
		Type:     string(transaction.Type),
	}

	if result := m.DB.Create(&transactionDB); result.Error != nil {
		log.WithError(result.Error).Error("error creating a transaction")
		return nil, result.Error
	}

	transaction.ID = transactionDB.ID

	return &transaction, nil
}
