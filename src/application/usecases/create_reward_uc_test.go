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

func TestCreateRewardUC_Execute(t *testing.T) {
	type fields struct {
		RewardRepository *mocks.RewardRepository
		StoreRepository  *mocks.StoreRepository
	}
	type args struct {
		createRewardDTO dto.CreateRewardDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mocker  func(a args, f fields)
		want    *dto.RewardCreatedDTO
		wantErr bool
	}{
		{
			name: "should return an error when store does not exist",
			fields: fields{
				StoreRepository: &mocks.StoreRepository{},
			},
			args: args{},
			mocker: func(a args, f fields) {
				f.StoreRepository.On("FindByID", mock.Anything).Return(nil).Once()
			},
			wantErr: true,
		},
		{
			name: "should return an error when amount type is not valid",
			fields: fields{
				StoreRepository: &mocks.StoreRepository{},
			},
			args: args{
				createRewardDTO: dto.CreateRewardDTO{
					AmountType: "tests",
				},
			},
			mocker: func(a args, f fields) {
				f.StoreRepository.On("FindByID", mock.Anything).Return(&entities.Store{}).Once()
			},
			wantErr: true,
		},
		{
			name: "should create a new reward when data is ok",
			fields: fields{
				StoreRepository:  &mocks.StoreRepository{},
				RewardRepository: &mocks.RewardRepository{},
			},
			args: args{
				createRewardDTO: dto.CreateRewardDTO{
					AmountType:  "coins",
					Reward:      "test",
					Description: "test",
					MinAmount:   2000,
					StoreID:     1,
				},
			},
			mocker: func(a args, f fields) {
				f.StoreRepository.On(
					"FindByID",
					a.createRewardDTO.StoreID,
				).Return(&entities.Store{}).Once()

				amountType, _ := vo.NewAmountType(a.createRewardDTO.AmountType)

				reward := entities.Reward{
					Reward:      a.createRewardDTO.Reward,
					Description: a.createRewardDTO.Description,
					MinAmount:   vo.NewAmountFromFloat(a.createRewardDTO.MinAmount),
					StoreID:     a.createRewardDTO.StoreID,
					Status:      vo.ActiveStatus,
					AmountType:  amountType,
				}

				f.RewardRepository.On("Create", reward).Return(&reward, nil).Once()
			},
			want: &dto.RewardCreatedDTO{
				AmountType:  "coins",
				Reward:      "test",
				Description: "test",
				MinAmount:   2000,
				StoreID:     1,
				Status:      "active",
			},
		},
		{
			name: "should return an error when reward cannot be created",
			fields: fields{
				StoreRepository:  &mocks.StoreRepository{},
				RewardRepository: &mocks.RewardRepository{},
			},
			args: args{
				createRewardDTO: dto.CreateRewardDTO{
					AmountType: "coins",
				},
			},
			mocker: func(a args, f fields) {
				f.StoreRepository.On(
					"FindByID",
					a.createRewardDTO.StoreID,
				).Return(&entities.Store{}).Once()
				f.RewardRepository.On(
					"Create",
					mock.Anything,
				).Return(nil, errors.New("error")).Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.args, tt.fields)
			c := &CreateRewardUC{
				RewardRepository: tt.fields.RewardRepository,
				StoreRepository:  tt.fields.StoreRepository,
			}
			got, err := c.Execute(tt.args.createRewardDTO)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateRewardUC.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRewardUC.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
