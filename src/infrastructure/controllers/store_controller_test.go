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

func TestStoreController_Create(t *testing.T) {
	type fields struct {
		CreateStoreUC *mocks.CreateStoreUC
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
			request:    []byte(`{"name":1}`),
			statusCode: http.StatusBadRequest,
			mocker:     func(f fields) {},
		},
		{
			name: "should create a store successfully",
			fields: fields{
				CreateStoreUC: &mocks.CreateStoreUC{},
			},
			request: []byte(`{
				"name":"test",
				"reward_points":20,
				"reward_coins":2,
				"min_amount":1000
			}`),
			statusCode: http.StatusCreated,
			mocker: func(f fields) {
				store := dto.CreateStoreDTO{
					Name:         "test",
					RewardPoints: 20,
					RewardCoins:  2,
					MinAmount:    1000,
				}

				f.CreateStoreUC.On("Execute", store).
					Return(&dto.StoreCreatedDTO{}, nil).
					Once()
			},
		},
		{
			name: "should return an error when store cannot be created",
			fields: fields{
				CreateStoreUC: &mocks.CreateStoreUC{},
			},
			request:    []byte(`{"name":"test"}`),
			statusCode: http.StatusInternalServerError,
			mocker: func(f fields) {
				f.CreateStoreUC.On("Execute", mock.Anything).
					Return(nil, errors.New("error")).
					Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.fields)
			s := &StoreController{
				CreateStoreUC: tt.fields.CreateStoreUC,
			}
			router := gin.Default()
			router.POST("/stores", s.Create)

			req, _ := http.NewRequest(
				http.MethodPost,
				"/stores",
				bytes.NewBuffer(tt.request),
			)

			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)
			assert.Equal(t, tt.statusCode, rec.Code)
		})
	}
}
