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

func TestCreateCampaignUC_Execute(t *testing.T) {
	type fields struct {
		CampaignRepository *mocks.CampaignRepository
	}
	type args struct {
		createCampaignDTO dto.CreateCampaignDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mocker  func(a args, f fields)
		want    *dto.CampaignCreatedDTO
		wantErr bool
	}{
		{
			name: "should create campaign successfully",
			fields: fields{
				CampaignRepository: &mocks.CampaignRepository{},
			},
			args: args{
				createCampaignDTO: dto.CreateCampaignDTO{
					Description: "test",
				},
			},
			mocker: func(a args, f fields) {
				campaign := entities.Campaign{
					Status:      vo.ActiveStatus,
					Description: "test",
				}
				f.CampaignRepository.On("Create", campaign).
					Return(&campaign, nil).
					Once()
			},
			want: &dto.CampaignCreatedDTO{
				Description: "test",
				Status:      string(vo.ActiveStatus),
			},
		},
		{
			name: "should return an error when campaign cannot be created",
			fields: fields{
				CampaignRepository: &mocks.CampaignRepository{},
			},
			args: args{
				createCampaignDTO: dto.CreateCampaignDTO{
					Description: "test",
				},
			},
			mocker: func(a args, f fields) {

				f.CampaignRepository.On("Create", mock.Anything).
					Return(nil, errors.New("error")).
					Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.args, tt.fields)
			c := &CreateCampaignUC{
				CampaignRepository: tt.fields.CampaignRepository,
			}
			got, err := c.Execute(tt.args.createCampaignDTO)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCampaignUC.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCampaignUC.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
