package usecases

import (
	"errors"

	"github.com/sirupsen/logrus"
	"lucio.com/order-service/src/domain/contracts/repositories"
	"lucio.com/order-service/src/domain/contracts/usecases"
	"lucio.com/order-service/src/domain/dto"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/vo"
)

var (
	errUserDoesNotExist           = errors.New("el usuario no existe")
	errStoreDoesNotExist          = errors.New("no se encuentra tienda para la sucursal ingresada")
	errTransactionCannotBeCreated = errors.New("la transacción no pudo ser creada")
)

type CreateTransactionUC struct {
	StoreRepository            repositories.StoreRepository
	TransactionRepository      repositories.TransactionRepository
	UserRepository             repositories.UserRepository
	CalculateCampaignRewardsUC usecases.CalculateCampaignRewardsUC
	Logger                     *logrus.Entry
}

func (c *CreateTransactionUC) Execute(
	createTransactionDTO dto.CreateTransactionDTO,
) (*dto.TransactionCreatedDTO, error) {
	log := c.Logger.WithFields(logrus.Fields{
		"file":                 "create_transaction_uc",
		"method":               "Execute",
		"createTransactionDTO": createTransactionDTO,
	})

	if c.UserRepository.FindByID(createTransactionDTO.UserID) == nil {
		log.Error(errUserDoesNotExist)
		return nil, errUserDoesNotExist
	}

	store := c.StoreRepository.FindByBranchID(createTransactionDTO.BranchID)
	if store == nil {
		log.Error(errStoreDoesNotExist)
		return nil, errStoreDoesNotExist
	}

	transaction := entities.Transaction{
		UserID:   createTransactionDTO.UserID,
		BranchID: createTransactionDTO.BranchID,
		Amount:   vo.NewAmountFromFloat(createTransactionDTO.Amount),
		Type:     vo.TransactionType(vo.AddType),
	}

	response := dto.TransactionCreatedDTO{
		UserID:   transaction.UserID,
		BranchID: transaction.BranchID,
		Amount:   transaction.Amount.Value(),
		Type:     string(transaction.Type),
	}

	if createTransactionDTO.Amount < store.MinAmount.Value() {
		transactionCreated, err := c.TransactionRepository.Create(transaction)
		if err != nil {
			log.Error(errTransactionCannotBeCreated)
			return nil, errTransactionCannotBeCreated
		}

		response.ID = transactionCreated.ID

		return &response, nil
	}

	points, coins := c.CalculateCampaignRewardsUC.Execute(
		createTransactionDTO.BranchID,
		store.RewardPoints,
		store.RewardCoins,
		transaction.Amount.Value(),
	)

	transaction.Points = store.RewardPoints + points
	transaction.Coins = store.RewardCoins + coins

	transactionCreated, err := c.TransactionRepository.Create(transaction)
	if err != nil {
		log.Error(errTransactionCannotBeCreated)
		return nil, errTransactionCannotBeCreated
	}

	response.ID = transactionCreated.ID
	response.Coins = transaction.Coins
	response.Points = transaction.Points

	return &response, nil
}
