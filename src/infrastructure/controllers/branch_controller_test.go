package controllers

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"lucio.com/order-service/src/domain/dto"
	"lucio.com/order-service/src/mocks"
)

func TestBranchController_Create(t *testing.T) {
	type fields struct {
		CreateBranchUC *mocks.CreateBranchUC
	}

	tests := []struct {
		name       string
		fields     fields
		request    []byte
		storeID    string
		mocker     func(f fields, storeID string)
		statusCode int
	}{
		{
			name:       "should return an error when data is invalid",
			request:    []byte(`{"name":12}`),
			storeID:    "1",
			statusCode: http.StatusBadRequest,
			mocker:     func(f fields, storeID string) {},
		},
		{
			name:       "should return an error when storeID is not a number",
			request:    []byte(`{"name":"test"}`),
			storeID:    "fake",
			statusCode: http.StatusBadRequest,
			mocker:     func(f fields, storeID string) {},
		},
		{
			name: "should create a branch successfully when data is ok",
			fields: fields{
				CreateBranchUC: &mocks.CreateBranchUC{},
			},
			request:    []byte(`{"name":"test"}`),
			storeID:    "1",
			statusCode: http.StatusCreated,
			mocker: func(f fields, storeID string) {
				branchDTO := dto.CreateBranchDTO{
					Name:    "test",
					StoreID: 1,
				}
				f.CreateBranchUC.On(
					"Execute",
					branchDTO,
				).Return(&dto.BranchCreatedDTO{}, nil).Once()
			},
		},
		{
			name: "should return an error when branch cannot be created successfully",
			fields: fields{
				CreateBranchUC: &mocks.CreateBranchUC{},
			},
			request:    []byte(`{"name":"test"}`),
			storeID:    "1",
			statusCode: http.StatusInternalServerError,
			mocker: func(f fields, storeID string) {
				f.CreateBranchUC.On(
					"Execute",
					mock.Anything,
				).Return(nil, errors.New("error")).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.fields, tt.storeID)
			b := &BranchController{
				CreateBranchUC: tt.fields.CreateBranchUC,
			}

			router := gin.Default()
			router.POST("/stores/:store_id/branches", b.Create)
			path := fmt.Sprintf("/stores/%s/branches", tt.storeID)
			req, _ := http.NewRequest(
				http.MethodPost,
				path,
				bytes.NewBuffer(tt.request),
			)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)
			assert.Equal(t, tt.statusCode, rec.Code)
		})
	}
}

func TestBranchController_CreateBranchCampaign(t *testing.T) {
	type fields struct {
		CreateBranchCampaignUC *mocks.CreateBranchCampaignUC
	}

	tests := []struct {
		name       string
		fields     fields
		request    []byte
		campaignID string
		branchID   string
		statusCode int
		mocker     func(f fields)
	}{
		{
			name:       "should return an error when data is invalid",
			request:    []byte(`{"operator":1}`),
			campaignID: "1",
			branchID:   "1",
			statusCode: http.StatusBadRequest,
			mocker:     func(f fields) {},
		},
		{
			name:       "should return an error when campaign_id is not a number",
			request:    []byte(`{"operator":"%"}`),
			campaignID: "fake",
			branchID:   "1",
			statusCode: http.StatusBadRequest,
			mocker:     func(f fields) {},
		},
		{
			name:       "should return an error when branch_id is not a number",
			request:    []byte(`{"operator":"%"}`),
			campaignID: "1",
			branchID:   "fake",
			statusCode: http.StatusBadRequest,
			mocker:     func(f fields) {},
		},
		{
			name: "should create campaign successfully when data is ok",
			fields: fields{
				CreateBranchCampaignUC: &mocks.CreateBranchCampaignUC{},
			},
			request: []byte(`{
				"operator":"%",
				"start_date":"2023-07-01 00:00:00",
				"end_date":"2023-07-01 12:00:00",
				"operator_value":3,
				"min_amount":1000
			}`),
			campaignID: "1",
			branchID:   "1",
			statusCode: http.StatusCreated,
			mocker: func(f fields) {
				campaign := dto.CreateBranchCampaignDTO{
					BranchID:      1,
					CampaignID:    1,
					StartDate:     "2023-07-01 00:00:00",
					EndDate:       "2023-07-01 12:00:00",
					Operator:      "%",
					OperatorValue: 3,
					MinAmount:     1000,
				}
				f.CreateBranchCampaignUC.On("Execute", campaign).
					Return(&dto.BranchCampaignCreatedDTO{}, nil).
					Once()
			},
		},
		{
			name: "should return an error when branch campaign cannot be created",
			fields: fields{
				CreateBranchCampaignUC: &mocks.CreateBranchCampaignUC{},
			},
			request:    []byte(`{}`),
			campaignID: "1",
			branchID:   "1",
			statusCode: http.StatusInternalServerError,
			mocker: func(f fields) {
				f.CreateBranchCampaignUC.On("Execute", mock.Anything).
					Return(nil, errors.New("error")).
					Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.fields)
			b := &BranchController{
				CreateBranchCampaignUC: tt.fields.CreateBranchCampaignUC,
			}

			router := gin.Default()
			router.POST("/campaigns/:campaign_id/branches/:branch_id", b.CreateBranchCampaign)
			path := fmt.Sprintf(
				"/campaigns/%s/branches/%s",
				tt.campaignID,
				tt.branchID,
			)
			req, _ := http.NewRequest(
				http.MethodPost,
				path,
				bytes.NewBuffer(tt.request),
			)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)
			assert.Equal(t, tt.statusCode, rec.Code)
		})
	}
}

func TestBranchController_AddCampaignToBranches(t *testing.T) {
	type fields struct {
		AddCampaignToStoreUC *mocks.AddCampaignToStoreUC
	}

	tests := []struct {
		name       string
		fields     fields
		request    []byte
		storeID    string
		campaignID string
		statusCode int
		mocker     func(f fields)
	}{
		{
			name:       "should return an error when data is invalid",
			request:    []byte(`{"operator":1}`),
			campaignID: "1",
			storeID:    "1",
			statusCode: http.StatusBadRequest,
			mocker:     func(f fields) {},
		},
		{
			name:       "should return an error when campaign_id is not a number",
			request:    []byte(`{"operator":"%"}`),
			campaignID: "fake",
			storeID:    "1",
			statusCode: http.StatusBadRequest,
			mocker:     func(f fields) {},
		},
		{
			name:       "should return an error when branch_id is not a number",
			request:    []byte(`{"operator":"%"}`),
			campaignID: "1",
			storeID:    "fake",
			statusCode: http.StatusBadRequest,
			mocker:     func(f fields) {},
		},
		{
			name: "should create campaign successfully when data is ok",
			fields: fields{
				AddCampaignToStoreUC: &mocks.AddCampaignToStoreUC{},
			},
			request: []byte(`{
				"operator":"%",
				"start_date":"2023-07-01 00:00:00",
				"end_date":"2023-07-01 12:00:00",
				"operator_value":3,
				"min_amount":1000
			}`),
			campaignID: "1",
			storeID:    "1",
			statusCode: http.StatusCreated,
			mocker: func(f fields) {
				campaign := dto.CreateStoreCampaignDTO{
					StoreID:       1,
					CampaignID:    1,
					StartDate:     "2023-07-01 00:00:00",
					EndDate:       "2023-07-01 12:00:00",
					Operator:      "%",
					OperatorValue: 3,
					MinAmount:     1000,
				}
				f.AddCampaignToStoreUC.On("Execute", campaign).
					Return(&dto.StoreCampaignCreatedDTO{}, nil).
					Once()
			},
		},
		{
			name: "should return an error when branch campaign cannot be created",
			fields: fields{
				AddCampaignToStoreUC: &mocks.AddCampaignToStoreUC{},
			},
			request:    []byte(`{}`),
			campaignID: "1",
			storeID:    "1",
			statusCode: http.StatusInternalServerError,
			mocker: func(f fields) {
				f.AddCampaignToStoreUC.On("Execute", mock.Anything).
					Return(nil, errors.New("error")).
					Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.fields)
			b := &BranchController{
				AddCampaignToStoreUC: tt.fields.AddCampaignToStoreUC,
			}

			router := gin.Default()
			router.POST("/campaigns/:campaign_id/stores/:store_id", b.AddCampaignToBranches)

			path := fmt.Sprintf(
				"/campaigns/%s/stores/%s",
				tt.campaignID,
				tt.storeID,
			)

			req, _ := http.NewRequest(
				http.MethodPost,
				path,
				bytes.NewBuffer(tt.request),
			)

			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)
			assert.Equal(t, tt.statusCode, rec.Code)
		})
	}
}

func TestBranchController_GetBranchCampaignsByBranch(t *testing.T) {
	type fields struct {
		GetBranchCampaignsUC *mocks.GetBranchCampaignsUC
	}

	tests := []struct {
		name       string
		fields     fields
		branchID   string
		statusCode int
		mocker     func(f fields)
	}{
		{
			name:       "should return an error when branch id is not valid",
			branchID:   "fake",
			statusCode: http.StatusBadRequest,
			mocker:     func(f fields) {},
		},
		{
			name: "should return branch campaigns successfully",
			fields: fields{
				GetBranchCampaignsUC: &mocks.GetBranchCampaignsUC{},
			},
			branchID:   "1",
			statusCode: http.StatusOK,
			mocker: func(f fields) {
				f.GetBranchCampaignsUC.On("Execute", uint(1)).
					Return([]dto.BranchCampaignReportDTO{}).
					Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.fields)
			b := &BranchController{
				GetBranchCampaignsUC: tt.fields.GetBranchCampaignsUC,
			}

			router := gin.Default()
			router.GET("/campaigns/branches/:branch_id", b.GetBranchCampaignsByBranch)

			path := fmt.Sprintf(
				"/campaigns/branches/%s",
				tt.branchID,
			)

			req, _ := http.NewRequest(http.MethodGet, path, nil)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)
			assert.Equal(t, tt.statusCode, rec.Code)
		})
	}
}
