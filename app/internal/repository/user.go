package repository

import (
	"fmt"
	"github.com/Nukie90/my-fluffy/app/domain/entity"
	"github.com/oklog/ulid/v2"
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
	// Create a new user but first check if the user already exists
	_, err := ur.FindByUsername(user.Username)
	if err == nil {
		return fmt.Errorf("user already exists")
	}
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

func (ur *UserRepo) FindByUsername(username string) (entity.User, error) {
	var user entity.User
	result := ur.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}

func (ur *UserRepo) FindByID(userID string) (entity.User, error) {
	var user entity.User
	userIDULID, err := ulid.Parse(userID)
	if err != nil {
		return entity.User{}, err
	}
	result := ur.DB.Where("id = ?", userIDULID).First(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}

func (ur *UserRepo) ListNotifications(userID ulid.ULID) ([]entity.Notification, error) {
	var notifications []entity.Notification
	result := ur.DB.Where("user_id = ?", userID).Find(&notifications)
	if result.Error != nil {
		return nil, result.Error
	}

	return notifications, nil
}

func (ur *UserRepo) DeleteNotification(notificationID uint) error {
	result := ur.DB.Delete(&entity.Notification{}, notificationID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
