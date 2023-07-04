package usecases

import (
	"errors"

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
	BranchCampaignRepository   repositories.BranchCampaignRepository
	CalculateCampaignRewardsUC usecases.CalculateCampaignRewardsUC
}

func (c *CreateTransactionUC) Execute(
	createTransactionUC dto.CreateTransactionDTO,
) (*dto.TransactionCreatedDTO, error) {
	if c.UserRepository.FindByID(createTransactionUC.UserID) == nil {
		return nil, errUserDoesNotExist
	}

	store := c.StoreRepository.FindByBranchID(createTransactionUC.BranchID)
	if store == nil {
		return nil, errStoreDoesNotExist
	}

	transaction := entities.Transaction{
		UserID:   createTransactionUC.UserID,
		BranchID: createTransactionUC.BranchID,
		Amount:   vo.NewAmountFromFloat(createTransactionUC.Amount),
		Type:     vo.TransactionType(vo.AddType),
	}

	response := dto.TransactionCreatedDTO{
		UserID:   transaction.UserID,
		BranchID: transaction.BranchID,
		Amount:   transaction.Amount.Value(),
		Type:     string(transaction.Type),
	}

	if createTransactionUC.Amount < store.MinAmount.Value() {
		transactionCreated, err := c.TransactionRepository.Create(transaction)
		if err != nil {
			return nil, errTransactionCannotBeCreated
		}

		response.ID = transactionCreated.ID

		return &response, nil
	}

	points, coins := c.CalculateCampaignRewardsUC.Execute(
		createTransactionUC.BranchID,
		store.RewardPoints,
		store.RewardCoins,
	)

	transaction.Points = store.RewardPoints + points
	transaction.Coins = store.RewardCoins + coins

	transactionCreated, err := c.TransactionRepository.Create(transaction)
	if err != nil {
		return nil, errTransactionCannotBeCreated
	}

	response.ID = transactionCreated.ID
	response.Coins = transaction.Coins
	response.Points = transaction.Points

	return &response, nil
}
