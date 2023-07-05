package usecases

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"lucio.com/order-service/src/domain/dto"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/vo"
	"lucio.com/order-service/src/mocks"
)

func TestCreateTransactionUC_Execute(t *testing.T) {
	type fields struct {
		StoreRepository            *mocks.StoreRepository
		TransactionRepository      *mocks.TransactionRepository
		UserRepository             *mocks.UserRepository
		CalculateCampaignRewardsUC *mocks.CalculateCampaignRewardsUC
	}
	type args struct {
		createTransactionDTO dto.CreateTransactionDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mocker  func(a args, f fields)
		want    *dto.TransactionCreatedDTO
		wantErr bool
	}{
		{
			name: "should return an error when user does not exists",
			fields: fields{
				UserRepository: &mocks.UserRepository{},
			},
			args: args{
				createTransactionDTO: dto.CreateTransactionDTO{},
			},
			mocker: func(a args, f fields) {
				f.UserRepository.On(
					"FindByID",
					a.createTransactionDTO.UserID,
				).Return(nil).Once()
			},
			wantErr: true,
		},
		{
			name: "should return an error when store does not exists",
			fields: fields{
				UserRepository:  &mocks.UserRepository{},
				StoreRepository: &mocks.StoreRepository{},
			},
			args: args{
				createTransactionDTO: dto.CreateTransactionDTO{},
			},
			mocker: func(a args, f fields) {
				f.UserRepository.On(
					"FindByID",
					a.createTransactionDTO.UserID,
				).Return(&entities.User{}).Once()

				f.StoreRepository.On(
					"FindByBranchID",
					a.createTransactionDTO.BranchID,
				).Return(nil).Once()
			},
			wantErr: true,
		},
		{
			name: "should create a transaction without rewards",
			fields: fields{
				UserRepository:        &mocks.UserRepository{},
				StoreRepository:       &mocks.StoreRepository{},
				TransactionRepository: &mocks.TransactionRepository{},
			},
			args: args{
				createTransactionDTO: dto.CreateTransactionDTO{
					Amount: 10000,
				},
			},
			mocker: func(a args, f fields) {
				f.UserRepository.On(
					"FindByID",
					a.createTransactionDTO.UserID,
				).Return(&entities.User{}).Once()

				store := entities.Store{
					MinAmount: vo.NewAmountFromFloat(20000),
				}

				f.StoreRepository.On(
					"FindByBranchID",
					a.createTransactionDTO.BranchID,
				).Return(&store).Once()

				transaction := entities.Transaction{
					Amount: vo.NewAmountFromFloat(10000),
					Type:   vo.TransactionType(vo.AddType),
				}

				f.TransactionRepository.On("Create", transaction).Return(&transaction, nil)
			},
			want: &dto.TransactionCreatedDTO{
				Amount: 10000,
				Type:   string(vo.AddType),
			},
		},
		{
			name: "should return an error when rewards are zero and create transaction fails",
			fields: fields{
				UserRepository:        &mocks.UserRepository{},
				StoreRepository:       &mocks.StoreRepository{},
				TransactionRepository: &mocks.TransactionRepository{},
			},
			args: args{
				createTransactionDTO: dto.CreateTransactionDTO{
					Amount: 10000,
				},
			},
			mocker: func(a args, f fields) {
				f.UserRepository.On(
					"FindByID",
					a.createTransactionDTO.UserID,
				).Return(&entities.User{}).Once()

				store := entities.Store{
					MinAmount: vo.NewAmountFromFloat(20000),
				}

				f.StoreRepository.On(
					"FindByBranchID",
					a.createTransactionDTO.BranchID,
				).Return(&store).Once()

				f.TransactionRepository.On("Create", mock.Anything).Return(nil, errors.New("error"))
			},
			wantErr: true,
		},
		{
			name: "should return create transaction successfully when user has rewards",
			fields: fields{
				UserRepository:             &mocks.UserRepository{},
				StoreRepository:            &mocks.StoreRepository{},
				TransactionRepository:      &mocks.TransactionRepository{},
				CalculateCampaignRewardsUC: &mocks.CalculateCampaignRewardsUC{},
			},
			args: args{
				createTransactionDTO: dto.CreateTransactionDTO{
					Amount: 10000,
				},
			},
			mocker: func(a args, f fields) {
				f.UserRepository.On(
					"FindByID",
					a.createTransactionDTO.UserID,
				).Return(&entities.User{}).Once()

				store := entities.Store{
					MinAmount:    vo.NewAmountFromFloat(9000),
					RewardPoints: 20,
					RewardCoins:  2,
				}

				f.StoreRepository.On(
					"FindByBranchID",
					a.createTransactionDTO.BranchID,
				).Return(&store).Once()

				f.CalculateCampaignRewardsUC.On(
					"Execute",
					a.createTransactionDTO.BranchID,
					store.RewardPoints,
					store.RewardCoins,
					float64(10000),
				).Return(uint(0), uint(0)).Once()

				transaction := entities.Transaction{
					Amount: vo.NewAmountFromFloat(10000),
					Type:   vo.TransactionType(vo.AddType),
					Points: store.RewardPoints,
					Coins:  store.RewardCoins,
				}

				f.TransactionRepository.On("Create", transaction).Return(&transaction, nil)
			},
			want: &dto.TransactionCreatedDTO{
				Amount: 10000,
				Points: 20,
				Coins:  2,
				Type:   string(vo.AddType),
			},
		},
		{
			name: "should return error when user has rewards but create transaction fails",
			fields: fields{
				UserRepository:             &mocks.UserRepository{},
				StoreRepository:            &mocks.StoreRepository{},
				TransactionRepository:      &mocks.TransactionRepository{},
				CalculateCampaignRewardsUC: &mocks.CalculateCampaignRewardsUC{},
			},
			args: args{
				createTransactionDTO: dto.CreateTransactionDTO{
					Amount: 10000,
				},
			},
			mocker: func(a args, f fields) {
				f.UserRepository.On(
					"FindByID",
					a.createTransactionDTO.UserID,
				).Return(&entities.User{}).Once()

				store := entities.Store{
					MinAmount:    vo.NewAmountFromFloat(9000),
					RewardPoints: 20,
					RewardCoins:  2,
				}

				f.StoreRepository.On(
					"FindByBranchID",
					a.createTransactionDTO.BranchID,
				).Return(&store).Once()

				f.CalculateCampaignRewardsUC.On(
					"Execute",
					a.createTransactionDTO.BranchID,
					store.RewardPoints,
					store.RewardCoins,
					float64(10000),
				).Return(uint(0), uint(0)).Once()

				f.TransactionRepository.On("Create", mock.Anything).Return(nil, errors.New("error"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.args, tt.fields)
			c := &CreateTransactionUC{
				StoreRepository:            tt.fields.StoreRepository,
				TransactionRepository:      tt.fields.TransactionRepository,
				UserRepository:             tt.fields.UserRepository,
				CalculateCampaignRewardsUC: tt.fields.CalculateCampaignRewardsUC,
			}
			got, err := c.Execute(tt.args.createTransactionDTO)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTransactionUC.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTransactionUC.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
