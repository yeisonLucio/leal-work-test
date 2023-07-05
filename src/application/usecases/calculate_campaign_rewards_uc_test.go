package usecases

import (
	"testing"

	"lucio.com/order-service/src/domain/dto"
	"lucio.com/order-service/src/mocks"
)

func TestCalculateCampaignRewardsUC_Execute(t *testing.T) {
	type fields struct {
		BranchCampaignRepository *mocks.BranchCampaignRepository
	}
	type args struct {
		branchID          uint
		storePoints       uint
		storeCoins        uint
		transactionAmount float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		mocker func(a args, f fields)
		want   uint
		want1  uint
	}{
		{
			name: "should return 0 when branch campaign does not exist",
			fields: fields{
				BranchCampaignRepository: &mocks.BranchCampaignRepository{},
			},
			args: args{
				branchID:          1,
				storePoints:       2,
				storeCoins:        2,
				transactionAmount: 2000,
			},
			mocker: func(a args, f fields) {
				f.BranchCampaignRepository.
					On("GetActivesByBranchID", a.branchID).
					Return([]dto.BranchCampaignCreatedDTO{}).
					Once()

			},
			want:  uint(0),
			want1: uint(0),
		},
		{
			name: "should return rewards according to percentage operation",
			fields: fields{
				BranchCampaignRepository: &mocks.BranchCampaignRepository{},
			},
			args: args{
				branchID:          1,
				storePoints:       10,
				storeCoins:        2,
				transactionAmount: 3000,
			},
			mocker: func(a args, f fields) {
				f.BranchCampaignRepository.On("GetActivesByBranchID", a.branchID).
					Return([]dto.BranchCampaignCreatedDTO{
						{
							Operator:      "%",
							OperatorValue: 20,
							MinAmount:     2000,
						},
					}).
					Once()
			},
			want:  uint(2),
			want1: uint(0),
		},
		{
			name: "should return rewards according to multiplication operation",
			fields: fields{
				BranchCampaignRepository: &mocks.BranchCampaignRepository{},
			},
			args: args{
				branchID:          1,
				storePoints:       10,
				storeCoins:        2,
				transactionAmount: 3000,
			},
			mocker: func(a args, f fields) {
				f.BranchCampaignRepository.On("GetActivesByBranchID", a.branchID).
					Return([]dto.BranchCampaignCreatedDTO{
						{
							Operator:      "*",
							OperatorValue: 2,
							MinAmount:     2000,
						},
					}).
					Once()
			},
			want:  uint(10),
			want1: uint(2),
		},
		{
			name: "should return 0 when campaign cannot be applied",
			fields: fields{
				BranchCampaignRepository: &mocks.BranchCampaignRepository{},
			},
			args: args{
				branchID:          1,
				storePoints:       2,
				storeCoins:        2,
				transactionAmount: 1000,
			},
			mocker: func(a args, f fields) {
				f.BranchCampaignRepository.
					On("GetActivesByBranchID", a.branchID).
					Return([]dto.BranchCampaignCreatedDTO{
						{
							Operator:      "*",
							OperatorValue: 2,
							MinAmount:     2000,
						},
					}).
					Once()

			},
			want:  uint(0),
			want1: uint(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.args, tt.fields)
			uc := &CalculateCampaignRewardsUC{
				BranchCampaignRepository: tt.fields.BranchCampaignRepository,
			}
			got, got1 := uc.Execute(
				tt.args.branchID,
				tt.args.storePoints,
				tt.args.storeCoins,
				tt.args.transactionAmount,
			)
			if got != tt.want {
				t.Errorf("CalculateCampaignRewardsUC.Execute() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CalculateCampaignRewardsUC.Execute() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
