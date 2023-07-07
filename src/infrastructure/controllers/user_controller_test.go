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

func TestUserController_Create(t *testing.T) {
	type fields struct {
		CreateUserUC *mocks.CreateUserUC
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
			name: "should create a user successfully",
			fields: fields{
				CreateUserUC: &mocks.CreateUserUC{},
			},
			request:    []byte(`{"name":"test","identification":"1060873267"}`),
			statusCode: http.StatusCreated,
			mocker: func(f fields) {
				user := dto.CreateUserDTO{
					Name:           "test",
					Identification: "1060873267",
				}

				f.CreateUserUC.On("Execute", user).
					Return(&dto.UserCreatedDTO{}, nil).
					Once()
			},
		},
		{
			name: "should return an error when user cannot be created",
			fields: fields{
				CreateUserUC: &mocks.CreateUserUC{},
			},
			request:    []byte(`{"name":"test"}`),
			statusCode: http.StatusInternalServerError,
			mocker: func(f fields) {
				f.CreateUserUC.On("Execute", mock.Anything).
					Return(nil, errors.New("error")).
					Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.fields)
			u := &UserController{
				CreateUserUC: tt.fields.CreateUserUC,
			}
			router := gin.Default()
			router.POST("/users", u.Create)

			req, _ := http.NewRequest(
				http.MethodPost,
				"/users",
				bytes.NewBuffer(tt.request),
			)

			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)
			assert.Equal(t, tt.statusCode, rec.Code)
		})
	}
}

func TestUserController_RegisterTransaction(t *testing.T) {
	type fields struct {
		CreateTransactionUC *mocks.CreateTransactionUC
	}

	tests := []struct {
		name       string
		fields     fields
		userID     string
		branchID   string
		request    []byte
		statusCode int
		mocker     func(f fields)
	}{
		{
			name:       "should return an error when data is no valid",
			fields:     fields{},
			userID:     "1",
			branchID:   "1",
			request:    []byte(`{"amount":"test"}`),
			statusCode: http.StatusBadRequest,
			mocker:     func(f fields) {},
		},
		{
			name:       "should return an error when user_id is no valid",
			fields:     fields{},
			userID:     "fake",
			branchID:   "1",
			request:    []byte(`{"amount":1000}`),
			statusCode: http.StatusBadRequest,
			mocker:     func(f fields) {},
		},
		{
			name:       "should return an error when branch_id is no valid",
			fields:     fields{},
			userID:     "1",
			branchID:   "fake",
			request:    []byte(`{"amount":1000}`),
			statusCode: http.StatusBadRequest,
			mocker:     func(f fields) {},
		},
		{
			name: "should create a transaction successfully",
			fields: fields{
				CreateTransactionUC: &mocks.CreateTransactionUC{},
			},
			userID:     "1",
			branchID:   "1",
			request:    []byte(`{"amount":1000}`),
			statusCode: http.StatusCreated,
			mocker: func(f fields) {
				transaction := dto.CreateTransactionDTO{
					UserID:   1,
					BranchID: 1,
					Amount:   1000,
				}

				f.CreateTransactionUC.On("Execute", transaction).
					Return(&dto.TransactionCreatedDTO{}, nil).
					Once()
			},
		},
		{
			name: "should return an error when transaction cannot be created",
			fields: fields{
				CreateTransactionUC: &mocks.CreateTransactionUC{},
			},
			userID:     "1",
			branchID:   "1",
			request:    []byte(`{"amount":1000}`),
			statusCode: http.StatusInternalServerError,
			mocker: func(f fields) {
				f.CreateTransactionUC.On("Execute", mock.Anything).
					Return(nil, errors.New("error")).
					Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.fields)
			u := &UserController{
				CreateTransactionUC: tt.fields.CreateTransactionUC,
			}

			router := gin.Default()
			router.POST("/users/:user_id/transactions/branches/:branch_id", u.RegisterTransaction)

			req, _ := http.NewRequest(
				http.MethodPost,
				fmt.Sprintf("/users/%s/transactions/branches/%s", tt.userID, tt.branchID),
				bytes.NewBuffer(tt.request),
			)

			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)
			assert.Equal(t, tt.statusCode, rec.Code)
		})
	}
}
