package usecases

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"lucio.com/order-service/src/domain/dto"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/domain/valueobjects"
)

func TestCreateUserUC_Execute(t *testing.T) {
	type fields struct {
		UserRepository *userRepositoryMock
	}
	type args struct {
		createUserDTO dto.CreateUserDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mocker  func(a args, f fields)
		want    *dto.UserCreatedDTO
		wantErr bool
	}{
		{
			name: "should return an error when user cannot be created",
			fields: fields{
				UserRepository: &userRepositoryMock{},
			},
			args: args{
				createUserDTO: dto.CreateUserDTO{},
			},
			mocker: func(a args, f fields) {
				user := entities.User{
					Status: valueobjects.ActiveStatus,
				}
				f.UserRepository.On("Create", user).Return(&user, errors.New("user create error")).Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should create user successfully",
			fields: fields{
				UserRepository: &userRepositoryMock{},
			},
			args: args{
				createUserDTO: dto.CreateUserDTO{
					Name:           "test",
					Identification: "123",
				},
			},
			mocker: func(a args, f fields) {
				user := entities.User{
					Status:         valueobjects.ActiveStatus,
					Name:           a.createUserDTO.Name,
					Identification: a.createUserDTO.Identification,
				}
				userResponse := user
				userResponse.ID = 1
				f.UserRepository.On("Create", user).Return(&userResponse, nil).Once()
			},
			want: &dto.UserCreatedDTO{
				ID:             1,
				Name:           "test",
				Identification: "123",
				Status:         string(valueobjects.ActiveStatus),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.args, tt.fields)
			c := &CreateUserUC{
				UserRepository: tt.fields.UserRepository,
			}
			got, err := c.Execute(tt.args.createUserDTO)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUserUC.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUserUC.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Mocks

type userRepositoryMock struct {
	mock.Mock
}

func (m *userRepositoryMock) Create(user entities.User) (*entities.User, error) {
	args := m.Called(user)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *userRepositoryMock) FindByID(ID uint) *entities.User {
	args := m.Called(ID)
	return args.Get(0).(*entities.User)
}
