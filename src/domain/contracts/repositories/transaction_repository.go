package repositories

import "lucio.com/order-service/src/domain/entities"

type TransactionRepository interface {
	Create(transaction entities.Transaction) (*entities.Transaction, error)
}
