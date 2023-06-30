package usecases

import (
	"errors"

	"lucio.com/order-service/src/domain/contracts/repositories"
	"lucio.com/order-service/src/domain/dto"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/valueobjects"
)

type CreateTransactionUC struct {
	StoreRepository          repositories.StoreRepository
	TransactionRepository    repositories.TransactionRepository
	UserRepository           repositories.UserRepository
	BranchCampaignRepository repositories.BranchCampaignRepository
}

func (c *CreateTransactionUC) Execute(
	createTransactionUC dto.CreateTransactionDTO,
) (*dto.TransactionCreatedDTO, error) {
	if c.UserRepository.FindByID(createTransactionUC.UserID) == nil {
		return nil, errors.New("el usuario no existe")
	}

	store := c.StoreRepository.FindByBranchID(createTransactionUC.BranchID)
	if store == nil {
		return nil, errors.New("no se encuentra tienda para la sucursal ingresada")
	}

	var amount valueobjects.Amount
	amount.NewFromFloat(createTransactionUC.Amount)

	transaction := entities.Transaction{
		UserID:   createTransactionUC.UserID,
		BranchID: createTransactionUC.BranchID,
		Amount:   amount,
		Type:     valueobjects.TransactionType(valueobjects.AddType),
	}

	response := dto.TransactionCreatedDTO{
		UserID:   transaction.UserID,
		BranchID: transaction.BranchID,
		Amount:   transaction.Amount.GetValue(),
		Type:     string(transaction.Type),
	}

	campaigns := c.BranchCampaignRepository.GetActivesByBranchID(transaction.BranchID)

	if createTransactionUC.Amount < store.MinAmount.GetValue() || len(campaigns) == 0 {
		transactionCreated, err := c.TransactionRepository.Create(transaction)
		if err != nil {
			return nil, errors.New("la transacción no pudo ser creada")
		}

		response.ID = transactionCreated.ID

		return &response, nil
	}

	var rewardPoints, rewardCoins uint

	for _, campaign := range campaigns {
		points, coins := c.calculateCampaignRewards(
			campaign,
			store.RewardPoints,
			store.RewardCoins,
		)

		rewardPoints += points
		rewardCoins += coins
	}

	transaction.Points = store.RewardPoints + rewardPoints
	transaction.Coins = store.RewardCoins + rewardCoins

	transactionCreated, err := c.TransactionRepository.Create(transaction)
	if err != nil {
		return nil, errors.New("la transacción no pudo ser creada")
	}

	response.ID = transactionCreated.ID
	response.Coins = transaction.Coins
	response.Points = transaction.Points

	return &response, nil
}

func (c *CreateTransactionUC) calculateCampaignRewards(
	campaign dto.BranchCampaignCreatedDTO,
	storePoints uint,
	storeCoins uint,
) (uint, uint) {
	switch campaign.Operator {
	case "%":
		storePoints = (storePoints * campaign.OperatorValue) / 100
		storeCoins = (storeCoins * campaign.OperatorValue) / 100

	case "*":
		storeCoins = (storeCoins * campaign.OperatorValue) - storeCoins
		storePoints *= (storePoints * campaign.OperatorValue) - storePoints

	}

	return storePoints, storeCoins
}
