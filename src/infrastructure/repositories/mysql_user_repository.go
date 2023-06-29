package repositories

import (
	"gorm.io/gorm"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/infrastructure/models"
)

type MysqlUserRepository struct {
	DB *gorm.DB
}

func (m *MysqlUserRepository) Create(User entities.User) (*entities.User, error) {
	userDB := models.User{
		Name:           User.Name,
		Status:         string(User.Status),
		Identification: User.Identification,
	}

	if result := m.DB.Create(&userDB); result.Error != nil {
		return nil, result.Error
	}

	User.ID = userDB.ID

	return &User, nil
}

func (m *MysqlUserRepository) FindByID(ID uint) *entities.User {
	var user entities.User

	result := m.DB.Model(&models.User{}).Find(&user, ID)
	if result.RowsAffected == 0 {
		return nil
	}

	return &user
}
