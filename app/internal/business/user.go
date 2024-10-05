package business

import (
	"github.com/Nukie90/my-fluffy/app/domain/entity"
	"github.com/Nukie90/my-fluffy/app/domain/model"
	"github.com/Nukie90/my-fluffy/app/internal/shared"

	"github.com/Nukie90/my-fluffy/app/internal/repository"
)

// UserUsecase is the usecase for users
type UserUsecase struct {
	UserRepo     *repository.UserRepo
	UserNotifier *shared.UserNotifier
}

// NewUserUsecase creates a new UserUsecase
func NewUserUsecase(ur *repository.UserRepo, un *shared.UserNotifier) *UserUsecase {
	return &UserUsecase{
		UserRepo:     ur,
		UserNotifier: un,
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

	uu.UserNotifier.NotifyObserver("", user.Username, "createUser")

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
