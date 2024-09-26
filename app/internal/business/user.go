package business

import (
	"github.com/Nukie90/my-fluffy/app/domain/entity"
	"github.com/Nukie90/my-fluffy/app/domain/model"

	"github.com/Nukie90/my-fluffy/app/internal/repository"
)

// UserUsecase is the usecase for users
type UserUsecase struct {
	UserRepo *repository.UserRepo
}

// NewUserUsecase creates a new UserUsecase
func NewUserUsecase(ur *repository.UserRepo) *UserUsecase {
	return &UserUsecase{UserRepo: ur}
}

// Create creates a new user
func (uu *UserUsecase) Create(user *model.Signup) error {
	u := &entity.User{
		Username: user.Username,
		Password: user.Password,
	}

	return uu.UserRepo.Create(u)
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
		})
	}

	return usersModel, nil
}
