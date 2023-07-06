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

func TestCreateStoreUC_Execute(t *testing.T) {
	type fields struct {
		StoreRepository *mocks.StoreRepository
	}
	type args struct {
		createStoreDTO dto.CreateStoreDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mocker  func(a args, f fields)
		want    *dto.StoreCreatedDTO
		wantErr bool
	}{
		{
			name: "should create a store successfully",
			fields: fields{
				StoreRepository: &mocks.StoreRepository{},
			},
			args: args{
				createStoreDTO: dto.CreateStoreDTO{
					Name:         "test",
					RewardPoints: 20,
					RewardCoins:  2,
					MinAmount:    10000,
				},
			},
			mocker: func(a args, f fields) {
				store := entities.Store{
					Name:         a.createStoreDTO.Name,
					Status:       vo.ActiveStatus,
					RewardPoints: a.createStoreDTO.RewardPoints,
					RewardCoins:  a.createStoreDTO.RewardCoins,
					MinAmount:    vo.NewAmountFromFloat(a.createStoreDTO.MinAmount),
				}

				f.StoreRepository.On("Create", store).Return(&store, nil).Once()
			},
			want: &dto.StoreCreatedDTO{
				Name:         "test",
				RewardPoints: 20,
				RewardCoins:  2,
				MinAmount:    10000,
				Status:       "active",
			},
		},
		{
			name: "should return an error when store cannot be created",
			fields: fields{
				StoreRepository: &mocks.StoreRepository{},
			},
			args: args{
				createStoreDTO: dto.CreateStoreDTO{},
			},
			mocker: func(a args, f fields) {
				f.StoreRepository.On("Create", mock.Anything).Return(nil, errors.New("error")).Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.args, tt.fields)
			c := &CreateStoreUC{
				StoreRepository: tt.fields.StoreRepository,
			}
			got, err := c.Execute(tt.args.createStoreDTO)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateStoreUC.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateStoreUC.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
