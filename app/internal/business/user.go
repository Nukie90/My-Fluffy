package business

import (
	"github.com/Nukie90/my-fluffy/app/domain/entity"
	"github.com/Nukie90/my-fluffy/app/domain/model"
	"github.com/Nukie90/my-fluffy/app/internal/shared"
	"github.com/oklog/ulid/v2"

	"github.com/Nukie90/my-fluffy/app/internal/repository"
)

// UserUsecase is the usecase for users
type UserUsecase struct {
	UserRepo       *repository.UserRepo
	AdminNotifier  *shared.UserNotifier
	ClientNotifier *shared.UserNotifier
}

// NewUserUsecase creates a new UserUsecase
func NewUserUsecase(ur *repository.UserRepo, an *shared.UserNotifier, cn *shared.UserNotifier) *UserUsecase {
	return &UserUsecase{
		UserRepo:       ur,
		AdminNotifier:  an,
		ClientNotifier: cn,
	}
}

// Create creates a new user
func (uu *UserUsecase) Create(user *model.Signup) error {
	u := &entity.User{
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role,
	}

	err := uu.UserRepo.Create(u)
	if err != nil {
		return err
	}

	uu.AdminNotifier.NotifyObserver("", user.Username, "createUser")

	return nil
}

// GetAll gets all users
func (uu *UserUsecase) GetAll() ([]model.User, error) {
	users, err := uu.UserRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var usersModel []model.User
	for _, user := range users {
		usersModel = append(usersModel, model.User{
			ID:       user.ID.String(),
			Username: user.Username,
			Password: user.Password,
			Role:     user.Role,
		})
	}

	return usersModel, nil
}

// Login logs in a user
func (uu *UserUsecase) Login(loginInfo *model.Login) (model.User, error) {
	user, err := uu.UserRepo.Login(loginInfo.Username, loginInfo.Password)
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		ID:       user.ID.String(),
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role,
	}, nil
}

// GetNotifications gets all notifications for a user
func (uu *UserUsecase) GetNotifications(userID string) ([]model.Notification, error) {
	userIDUlid, err := ulid.Parse(userID)
	if err != nil {
		return nil, err
	}

	notifications, err := uu.UserRepo.ListNotifications(userIDUlid)
	if err != nil {
		return nil, err
	}

	var notificationsModel []model.Notification
	for _, notification := range notifications {
		timeStr := notification.CreatedAt.Format("2006-01-02 15:04:05")
		notificationsModel = append(notificationsModel, model.Notification{
			ID:       notification.ID,
			OwnerID:  notification.UserID.String(),
			Message:  notification.Message,
			CreateAt: timeStr,
		})
	}

	return notificationsModel, nil
}

// Delete Notification deletes a notification
func (uu *UserUsecase) DeleteNotification(notificationID uint) error {
	err := uu.UserRepo.DeleteNotification(notificationID)
	if err != nil {
		return err
	}

	return nil
}
