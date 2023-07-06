package usecases

import (
	"errors"
	"reflect"
	"testing"

	"lucio.com/order-service/src/domain/dto"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/mocks"
)

func TestCreateBranchCampaignUC_Execute(t *testing.T) {
	type fields struct {
		BranchRepository         *mocks.BranchRepository
		CampaignRepository       *mocks.CampaignRepository
		BranchCampaignRepository *mocks.BranchCampaignRepository
	}
	type args struct {
		createBranchCampaignDTO dto.CreateBranchCampaignDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mocker  func(a args, f fields)
		want    *dto.BranchCampaignCreatedDTO
		wantErr bool
	}{
		{
			name: "should return an error when campaign does not exist",
			fields: fields{
				CampaignRepository: &mocks.CampaignRepository{},
			},
			args: args{},
			mocker: func(a args, f fields) {
				f.CampaignRepository.On(
					"FindByID",
					a.createBranchCampaignDTO.CampaignID,
				).Return(nil).Once()
			},
			wantErr: true,
		},
		{
			name: "should return an error when branch does not exist",
			fields: fields{
				CampaignRepository: &mocks.CampaignRepository{},
				BranchRepository:   &mocks.BranchRepository{},
			},
			args: args{},
			mocker: func(a args, f fields) {
				f.CampaignRepository.On(
					"FindByID",
					a.createBranchCampaignDTO.CampaignID,
				).Return(&entities.Campaign{}).Once()

				f.BranchRepository.On(
					"FindByID",
					a.createBranchCampaignDTO.BranchID,
				).Return(nil).Once()
			},
			wantErr: true,
		},
		{
			name: "should return an error when start date is not valid",
			fields: fields{
				CampaignRepository: &mocks.CampaignRepository{},
				BranchRepository:   &mocks.BranchRepository{},
			},
			args: args{},
			mocker: func(a args, f fields) {
				f.CampaignRepository.On(
					"FindByID",
					a.createBranchCampaignDTO.CampaignID,
				).Return(&entities.Campaign{}).Once()

				f.BranchRepository.On(
					"FindByID",
					a.createBranchCampaignDTO.BranchID,
				).Return(&entities.Branch{}).Once()
			},
			wantErr: true,
		},
		{
			name: "should return an error when end date is not valid",
			fields: fields{
				CampaignRepository: &mocks.CampaignRepository{},
				BranchRepository:   &mocks.BranchRepository{},
			},
			args: args{
				dto.CreateBranchCampaignDTO{
					StartDate: "2023-01-01 00:00:00",
				},
			},
			mocker: func(a args, f fields) {
				f.CampaignRepository.On(
					"FindByID",
					a.createBranchCampaignDTO.CampaignID,
				).Return(&entities.Campaign{}).Once()

				f.BranchRepository.On(
					"FindByID",
					a.createBranchCampaignDTO.BranchID,
				).Return(&entities.Branch{}).Once()
			},
			wantErr: true,
		},
		{
			name: "should return an error when operator is not allowed",
			fields: fields{
				CampaignRepository: &mocks.CampaignRepository{},
				BranchRepository:   &mocks.BranchRepository{},
			},
			args: args{
				dto.CreateBranchCampaignDTO{
					StartDate: "2023-01-01 00:00:00",
					EndDate:   "2023-01-01 00:00:00",
					Operator:  "+",
				},
			},
			mocker: func(a args, f fields) {
				f.CampaignRepository.On(
					"FindByID",
					a.createBranchCampaignDTO.CampaignID,
				).Return(&entities.Campaign{}).Once()

				f.BranchRepository.On(
					"FindByID",
					a.createBranchCampaignDTO.BranchID,
				).Return(&entities.Branch{}).Once()
			},
			wantErr: true,
		},
		{
			name: "should return an error when campaign cannot be associated to branch",
			fields: fields{
				CampaignRepository:       &mocks.CampaignRepository{},
				BranchRepository:         &mocks.BranchRepository{},
				BranchCampaignRepository: &mocks.BranchCampaignRepository{},
			},
			args: args{
				dto.CreateBranchCampaignDTO{
					StartDate: "2023-01-01 00:00:00",
					EndDate:   "2023-01-01 00:00:00",
					Operator:  "%",
				},
			},
			mocker: func(a args, f fields) {
				f.CampaignRepository.On(
					"FindByID",
					a.createBranchCampaignDTO.CampaignID,
				).Return(&entities.Campaign{}).Once()

				f.BranchRepository.On(
					"FindByID",
					a.createBranchCampaignDTO.BranchID,
				).Return(&entities.Branch{}).Once()

				branchCampaign := entities.BranchCampaign{}
				branchCampaign.SetOperator(a.createBranchCampaignDTO.Operator)
				branchCampaign.SetStartDate(a.createBranchCampaignDTO.StartDate)
				branchCampaign.SetEndDate(a.createBranchCampaignDTO.EndDate)

				f.BranchCampaignRepository.On(
					"Create",
					branchCampaign,
				).Return(nil, errors.New("error"))
			},
			wantErr: true,
		},
		{
			name: "should associate campaign to branch when data is ok",
			fields: fields{
				CampaignRepository:       &mocks.CampaignRepository{},
				BranchRepository:         &mocks.BranchRepository{},
				BranchCampaignRepository: &mocks.BranchCampaignRepository{},
			},
			args: args{
				dto.CreateBranchCampaignDTO{
					StartDate: "2023-01-01 00:00:00",
					EndDate:   "2023-01-01 00:00:00",
					Operator:  "%",
				},
			},
			mocker: func(a args, f fields) {
				f.CampaignRepository.On(
					"FindByID",
					a.createBranchCampaignDTO.CampaignID,
				).Return(&entities.Campaign{}).Once()

				f.BranchRepository.On(
					"FindByID",
					a.createBranchCampaignDTO.BranchID,
				).Return(&entities.Branch{}).Once()

				branchCampaign := entities.BranchCampaign{}
				branchCampaign.SetOperator(a.createBranchCampaignDTO.Operator)
				branchCampaign.SetStartDate(a.createBranchCampaignDTO.StartDate)
				branchCampaign.SetEndDate(a.createBranchCampaignDTO.EndDate)

				f.BranchCampaignRepository.On(
					"Create",
					branchCampaign,
				).Return(&branchCampaign, nil)
			},
			want: &dto.BranchCampaignCreatedDTO{
				StartDate: "2023-01-01 00:00:00 +0000 UTC",
				EndDate:   "2023-01-01 00:00:00 +0000 UTC",
				Operator:  "%",
			},
		},
	}
	for _, tt := range tests {
		tt.mocker(tt.args, tt.fields)
		t.Run(tt.name, func(t *testing.T) {
			c := &CreateBranchCampaignUC{
				BranchRepository:         tt.fields.BranchRepository,
				CampaignRepository:       tt.fields.CampaignRepository,
				BranchCampaignRepository: tt.fields.BranchCampaignRepository,
			}
			got, err := c.Execute(tt.args.createBranchCampaignDTO)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateBranchCampaignUC.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateBranchCampaignUC.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
