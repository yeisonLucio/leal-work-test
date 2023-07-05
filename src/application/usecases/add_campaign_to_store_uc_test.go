package usecases

import (
	"errors"
	"reflect"
	"testing"

	"lucio.com/order-service/src/domain/dto"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/mocks"
)

func TestAddCampaignToStoreUC_Execute(t *testing.T) {
	type fields struct {
		BranchRepository       *mocks.BranchRepository
		CreateBranchCampaignUC *mocks.CreateBranchCampaignUC
		CampaignRepository     *mocks.CampaignRepository
		StoreRepository        *mocks.StoreRepository
	}
	type args struct {
		createStoreCampaignDTO dto.CreateStoreCampaignDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mocker  func(a args, f fields)
		want    *dto.StoreCampaignCreatedDTO
		wantErr bool
	}{
		{
			name: "should return an error when campaign does not exist",
			fields: fields{
				CampaignRepository: &mocks.CampaignRepository{},
			},
			args: args{
				createStoreCampaignDTO: dto.CreateStoreCampaignDTO{
					CampaignID: 1,
				},
			},
			mocker: func(a args, f fields) {
				f.CampaignRepository.On(
					"FindByID",
					a.createStoreCampaignDTO.CampaignID,
				).Return(nil).Once()
			},
			wantErr: true,
		},
		{
			name: "should return an error when store does not exists",
			fields: fields{
				CampaignRepository: &mocks.CampaignRepository{},
				StoreRepository:    &mocks.StoreRepository{},
			},
			args: args{
				createStoreCampaignDTO: dto.CreateStoreCampaignDTO{
					CampaignID: 1,
				},
			},
			mocker: func(a args, f fields) {
				f.CampaignRepository.On(
					"FindByID",
					a.createStoreCampaignDTO.CampaignID,
				).Return(&entities.Campaign{}).Once()

				f.StoreRepository.On(
					"FindByID",
					a.createStoreCampaignDTO.StoreID,
				).Return(nil).Once()

			},
			wantErr: true,
		},
		{
			name: "should not associate any campaigns when store does not has a branches",
			fields: fields{
				CampaignRepository: &mocks.CampaignRepository{},
				StoreRepository:    &mocks.StoreRepository{},
				BranchRepository:   &mocks.BranchRepository{},
			},
			args: args{
				createStoreCampaignDTO: dto.CreateStoreCampaignDTO{
					CampaignID: 1,
					StoreID:    1,
				},
			},
			mocker: func(a args, f fields) {
				f.CampaignRepository.On(
					"FindByID",
					a.createStoreCampaignDTO.CampaignID,
				).Return(&entities.Campaign{}).Once()

				f.StoreRepository.On(
					"FindByID",
					a.createStoreCampaignDTO.StoreID,
				).Return(&entities.Store{}).Once()

				f.BranchRepository.On(
					"GetIdsByStoreID",
					a.createStoreCampaignDTO.StoreID,
				).Return([]uint{}).Once()

			},
			want: &dto.StoreCampaignCreatedDTO{},
		},
		{
			name: "should associate campaign when the store has two branches",
			fields: fields{
				CampaignRepository:     &mocks.CampaignRepository{},
				StoreRepository:        &mocks.StoreRepository{},
				BranchRepository:       &mocks.BranchRepository{},
				CreateBranchCampaignUC: &mocks.CreateBranchCampaignUC{},
			},
			args: args{
				createStoreCampaignDTO: dto.CreateStoreCampaignDTO{
					CampaignID: 1,
					StoreID:    1,
				},
			},
			mocker: func(a args, f fields) {
				f.CampaignRepository.On(
					"FindByID",
					a.createStoreCampaignDTO.CampaignID,
				).Return(&entities.Campaign{}).Once()

				f.StoreRepository.On(
					"FindByID",
					a.createStoreCampaignDTO.StoreID,
				).Return(&entities.Store{}).Once()

				f.BranchRepository.On(
					"GetIdsByStoreID",
					a.createStoreCampaignDTO.StoreID,
				).Return([]uint{1, 2}).Once()

				branchCampaignDTO := dto.CreateBranchCampaignDTO{BranchID: 1, CampaignID: 1}

				f.CreateBranchCampaignUC.On("Execute", branchCampaignDTO).
					Return(&dto.BranchCampaignCreatedDTO{}, nil).
					Once()

				branchCampaignDTO.BranchID = 2
				f.CreateBranchCampaignUC.On("Execute", branchCampaignDTO).
					Return(&dto.BranchCampaignCreatedDTO{}, nil).
					Once()

			},
			want: &dto.StoreCampaignCreatedDTO{
				BranchCampaigns: []dto.BranchCampaignCreatedDTO{
					{},
					{},
				},
			},
		},
		{
			name: "should return an errors when campaign does cannot be associated to store branch",
			fields: fields{
				CampaignRepository:     &mocks.CampaignRepository{},
				StoreRepository:        &mocks.StoreRepository{},
				BranchRepository:       &mocks.BranchRepository{},
				CreateBranchCampaignUC: &mocks.CreateBranchCampaignUC{},
			},
			args: args{
				createStoreCampaignDTO: dto.CreateStoreCampaignDTO{
					CampaignID: 1,
					StoreID:    1,
				},
			},
			mocker: func(a args, f fields) {
				f.CampaignRepository.On(
					"FindByID",
					a.createStoreCampaignDTO.CampaignID,
				).Return(&entities.Campaign{}).Once()

				f.StoreRepository.On(
					"FindByID",
					a.createStoreCampaignDTO.StoreID,
				).Return(&entities.Store{}).Once()

				f.BranchRepository.On(
					"GetIdsByStoreID",
					a.createStoreCampaignDTO.StoreID,
				).Return([]uint{1}).Once()

				branchCampaignDTO := dto.CreateBranchCampaignDTO{BranchID: 1, CampaignID: 1}

				f.CreateBranchCampaignUC.On("Execute", branchCampaignDTO).
					Return(nil, errors.New("error")).
					Once()

			},
			want: &dto.StoreCampaignCreatedDTO{
				Errors: []dto.ErroBranchCampaign{
					{
						Message:  "error",
						BranchId: 1,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.args, tt.fields)
			a := &AddCampaignToStoreUC{
				BranchRepository:       tt.fields.BranchRepository,
				CreateBranchCampaignUC: tt.fields.CreateBranchCampaignUC,
				CampaignRepository:     tt.fields.CampaignRepository,
				StoreRepository:        tt.fields.StoreRepository,
			}
			got, err := a.Execute(tt.args.createStoreCampaignDTO)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddCampaignToStoreUC.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddCampaignToStoreUC.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
