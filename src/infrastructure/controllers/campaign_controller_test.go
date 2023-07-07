package controllers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"lucio.com/order-service/src/domain/dto"
	"lucio.com/order-service/src/mocks"
)

func TestCampaignController_Create(t *testing.T) {
	type fields struct {
		CreateCampaignUC *mocks.CreateCampaignUC
	}

	tests := []struct {
		name       string
		fields     fields
		request    []byte
		statusCode int
		mocker     func(f fields)
	}{
		{
			name:       "should return an error when data is no valid",
			fields:     fields{},
			request:    []byte(`{"description":1}`),
			statusCode: http.StatusBadRequest,
			mocker:     func(f fields) {},
		},
		{
			name: "should create a campaign successfully",
			fields: fields{
				CreateCampaignUC: &mocks.CreateCampaignUC{},
			},
			request:    []byte(`{"description":"test"}`),
			statusCode: http.StatusCreated,
			mocker: func(f fields) {
				campaign := dto.CreateCampaignDTO{Description: "test"}

				f.CreateCampaignUC.On("Execute", campaign).
					Return(&dto.CampaignCreatedDTO{}, nil).
					Once()
			},
		},
		{
			name: "should return an error when campaign cannot be created",
			fields: fields{
				CreateCampaignUC: &mocks.CreateCampaignUC{},
			},
			request:    []byte(`{"description":"test"}`),
			statusCode: http.StatusInternalServerError,
			mocker: func(f fields) {
				f.CreateCampaignUC.On("Execute", mock.Anything).
					Return(nil, errors.New("error")).
					Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.fields)
			s := &CampaignController{
				CreateCampaignUC: tt.fields.CreateCampaignUC,
			}
			router := gin.Default()
			router.POST("/campaigns", s.Create)

			req, _ := http.NewRequest(
				http.MethodPost,
				"/campaigns",
				bytes.NewBuffer(tt.request),
			)

			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)
			assert.Equal(t, tt.statusCode, rec.Code)
		})
	}
}
