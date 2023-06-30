package usecases

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"lucio.com/order-service/src/domain/dto"
	"lucio.com/order-service/src/domain/entities"
)

func TestGetBranchCampaignsUC_Execute(t *testing.T) {
	testData := []dto.BranchCampaignReportDTO{
		{
			ID:            1,
			BranchID:      1,
			CampaignID:    1,
			Description:   "test campaign",
			Status:        "active",
			StartDate:     "2023-02-25 12:00:00",
			EndDate:       "2023-02-28 12:00:00",
			Operator:      "%",
			OperatorValue: 30,
			MinAmount:     2000,
		},
	}

	objectData, _ := json.Marshal(testData)

	type fields struct {
		BranchCampaignRepository *branchCampaignRepositoryMock
		CacheRepository          *cacheRepositoryMock
	}
	type args struct {
		branchID uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		mocker func(a args, f fields)
		want   []dto.BranchCampaignReportDTO
	}{
		{
			name: "should return empty array from db",
			fields: fields{
				BranchCampaignRepository: &branchCampaignRepositoryMock{},
				CacheRepository:          &cacheRepositoryMock{},
			},
			args: args{
				branchID: 1,
			},
			mocker: func(a args, f fields) {
				f.BranchCampaignRepository.On(
					"FindByBranchID",
					a.branchID,
				).Return(testData).Once()

				key := fmt.Sprintf("branch-campaign-report-%d", a.branchID)

				f.CacheRepository.On("GetByKey", key).Return("", errors.New("data not found")).Once()
				f.CacheRepository.On("SetByKey", key, string(objectData)).Return(nil).Once()
			},
			want: testData,
		},
		{
			name: "should return data from cache",
			fields: fields{
				BranchCampaignRepository: &branchCampaignRepositoryMock{},
				CacheRepository:          &cacheRepositoryMock{},
			},
			args: args{
				branchID: 1,
			},
			mocker: func(a args, f fields) {

				key := fmt.Sprintf("branch-campaign-report-%d", a.branchID)

				f.CacheRepository.On("GetByKey", key).Return(string(objectData), nil).Once()
			},
			want: testData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.args, tt.fields)
			g := &GetBranchCampaignsUC{
				BranchCampaignRepository: tt.fields.BranchCampaignRepository,
				CacheRepository:          tt.fields.CacheRepository,
			}
			if got := g.Execute(tt.args.branchID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBranchCampaignsUC.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}

// mocks

type branchCampaignRepositoryMock struct {
	mock.Mock
}

func (m *branchCampaignRepositoryMock) Create(branchCampaign entities.BranchCampaign) (*entities.BranchCampaign, error) {
	args := m.Called(branchCampaign)
	return args.Get(0).(*entities.BranchCampaign), args.Error(1)
}

func (m *branchCampaignRepositoryMock) FindByID(ID uint) *entities.BranchCampaign {
	args := m.Called(ID)
	return args.Get(0).(*entities.BranchCampaign)
}

func (m *branchCampaignRepositoryMock) FindByBranchID(branchID uint) []dto.BranchCampaignReportDTO {
	args := m.Called(branchID)
	return args.Get(0).([]dto.BranchCampaignReportDTO)
}

func (m *branchCampaignRepositoryMock) GetActivesByBranchID(branchID uint) []dto.BranchCampaignCreatedDTO {
	args := m.Called(branchID)
	return args.Get(0).([]dto.BranchCampaignCreatedDTO)
}

type cacheRepositoryMock struct {
	mock.Mock
}

func (m *cacheRepositoryMock) SetByKey(key string, object string) error {
	return m.Called(key, object).Error(0)
}
func (m *cacheRepositoryMock) GetByKey(key string) (string, error) {
	args := m.Called(key)
	return args.Get(0).(string), args.Error(1)
}
