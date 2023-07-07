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

func TestRewardController_Create(t *testing.T) {
	type fields struct {
		CreateRewardUC *mocks.CreateRewardUC
	}
	tests := []struct {
		name       string
		fields     fields
		storeID    string
		request    []byte
		statusCode int
		mocker     func(f fields)
	}{
		{
			name:       "should return an error when data is no valid",
			fields:     fields{},
			request:    []byte(`{"min_amount":"fake"}`),
			statusCode: http.StatusBadRequest,
			mocker:     func(f fields) {},
		},
		{
			name: "should create a reward successfully",
			fields: fields{
				CreateRewardUC: &mocks.CreateRewardUC{},
			},
			storeID: "1",
			request: []byte(`{
				"reward":"test",
				"description":"test",
				"min_amount":1000,
				"amount_type":"coins"
			}`),
			statusCode: http.StatusCreated,
			mocker: func(f fields) {
				reward := dto.CreateRewardDTO{
					Reward:      "test",
					MinAmount:   1000,
					AmountType:  "coins",
					StoreID:     1,
					Description: "test",
				}

				f.CreateRewardUC.On("Execute", reward).
					Return(&dto.RewardCreatedDTO{}, nil).
					Once()
			},
		},
		{
			name: "should return an error when reward cannot be created",
			fields: fields{
				CreateRewardUC: &mocks.CreateRewardUC{},
			},
			storeID:    "1",
			request:    []byte(`{"description":"test"}`),
			statusCode: http.StatusInternalServerError,
			mocker: func(f fields) {
				f.CreateRewardUC.On("Execute", mock.Anything).
					Return(nil, errors.New("error")).
					Once()
			},
		},
		{
			name:       "should return an error when store_id is no valid",
			storeID:    "fake",
			request:    []byte(`{"description":"test"}`),
			statusCode: http.StatusBadRequest,
			mocker:     func(f fields) {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.fields)
			r := &RewardController{
				CreateRewardUC: tt.fields.CreateRewardUC,
			}
			router := gin.Default()
			router.POST("/stores/:store_id/rewards", r.Create)

			req, _ := http.NewRequest(
				http.MethodPost,
				fmt.Sprintf("/stores/%s/rewards", tt.storeID),
				bytes.NewBuffer(tt.request),
			)

			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)
			assert.Equal(t, tt.statusCode, rec.Code)
		})
	}
}
