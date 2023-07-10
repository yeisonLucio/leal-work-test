package repositories

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"lucio.com/order-service/src/domain/entities"
	"lucio.com/order-service/src/infrastructure/models"
)

type MysqlUserRepository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func (m *MysqlUserRepository) Create(user entities.User) (*entities.User, error) {
	log := m.Logger.WithFields(logrus.Fields{
		"file":   "mysql_user_repository",
		"method": "Create",
		"user":   user,
	})

	userDB := models.User{
		Name:           user.Name,
		Status:         string(user.Status),
		Identification: user.Identification,
	}

	if result := m.DB.Create(&userDB); result.Error != nil {
		log.WithError(result.Error).Error("error creating user")
		return nil, result.Error
	}

	user.ID = userDB.ID

	return &user, nil
}

func (m *MysqlUserRepository) FindByID(ID uint) *entities.User {
	var user entities.User

	result := m.DB.Model(&models.User{}).Find(&user, ID)
	if result.RowsAffected == 0 {
		return nil
	}

	return &user
}
