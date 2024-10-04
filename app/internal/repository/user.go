package repository

import (
	"fmt"
	"github.com/Nukie90/my-fluffy/app/domain/entity"
	"gorm.io/gorm"
)

// UserRepo is the repository for users
type UserRepo struct {
	DB *gorm.DB
}

// NewUserRepo creates a new UserRepo
func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

// Create creates a new user
func (ur *UserRepo) Create(user *entity.User) error {
	result := ur.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// FindAll finds all users
func (ur *UserRepo) FindAll() ([]entity.User, error) {
	var users []entity.User
	result := ur.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	fmt.Println(result)
	return users, nil
}

// FindAdmin finds the admin user
func (ur *UserRepo) FindAdmin() ([]entity.User, error) {
	var admins []entity.User
	result := ur.DB.Where("role = ?", "admin").Find(&admins)
	if result.Error != nil {
		return nil, result.Error
	}

	return admins, nil
}

// StoreNotification stores a notification
func (ur *UserRepo) StoreNotification(notification *entity.Notification) error {
	result := ur.DB.Create(notification)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ur *UserRepo) Login(username, password string) (entity.User, error) {
	var user entity.User
	result := ur.DB.Where("username = ? AND password = ?", username, password).First(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}
