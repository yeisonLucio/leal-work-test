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

func TestCreateBranchUC_Execute(t *testing.T) {
	type fields struct {
		BranchRepository *mocks.BranchRepository
		StoreRepository  *mocks.StoreRepository
	}
	type args struct {
		createBranchDTO dto.CreateBranchDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mocker  func(a args, f fields)
		want    *dto.BranchCreatedDTO
		wantErr bool
	}{
		{
			name: "should return an error when store does not exists",
			fields: fields{
				StoreRepository: &mocks.StoreRepository{},
			},
			args: args{
				createBranchDTO: dto.CreateBranchDTO{},
			},
			mocker: func(a args, f fields) {
				f.StoreRepository.On(
					"FindByID",
					a.createBranchDTO.StoreID,
				).Return(nil).Once()
			},
			wantErr: true,
		},
		{
			name: "should return an error when the branch cannot be created",
			fields: fields{
				StoreRepository:  &mocks.StoreRepository{},
				BranchRepository: &mocks.BranchRepository{},
			},
			args: args{
				createBranchDTO: dto.CreateBranchDTO{},
			},
			mocker: func(a args, f fields) {
				f.StoreRepository.On(
					"FindByID",
					a.createBranchDTO.StoreID,
				).Return(&entities.Store{}).Once()

				f.BranchRepository.On("Create", mock.Anything).
					Return(nil, errors.New("error")).
					Once()
			},
			wantErr: true,
		},
		{
			name: "should create a new branch when data is ok",
			fields: fields{
				StoreRepository:  &mocks.StoreRepository{},
				BranchRepository: &mocks.BranchRepository{},
			},
			args: args{
				createBranchDTO: dto.CreateBranchDTO{
					Name:    "test",
					StoreID: 1,
				},
			},
			mocker: func(a args, f fields) {
				f.StoreRepository.On(
					"FindByID",
					a.createBranchDTO.StoreID,
				).Return(&entities.Store{}).Once()

				branch := entities.Branch{
					Name:    a.createBranchDTO.Name,
					Status:  vo.ActiveStatus,
					StoreID: a.createBranchDTO.StoreID,
				}

				f.BranchRepository.On("Create", branch).
					Return(&branch, nil).
					Once()
			},
			want: &dto.BranchCreatedDTO{
				Name:    "test",
				Status:  "active",
				StoreID: 1,
			},
		},
	}
	for _, tt := range tests {
		tt.mocker(tt.args, tt.fields)
		t.Run(tt.name, func(t *testing.T) {
			c := &CreateBranchUC{
				BranchRepository: tt.fields.BranchRepository,
				StoreRepository:  tt.fields.StoreRepository,
			}
			got, err := c.Execute(tt.args.createBranchDTO)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateBranchUC.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateBranchUC.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
