package usecases

import "lucio.com/order-service/src/domain/dto"

type CreateTransactionUC interface {
	Execute(createTransactionUC dto.CreateTransactionDTO) (*dto.TransactionCreatedDTO, error)
}
