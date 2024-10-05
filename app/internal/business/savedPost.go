package business

import (
	"github.com/Nukie90/my-fluffy/app/domain/entity"
	"github.com/Nukie90/my-fluffy/app/domain/model"
	"github.com/Nukie90/my-fluffy/app/internal/repository"
)

// SavedPostUsecase is the usecase for saved posts
type SavedPostUsecase struct {
	SavedPostRepo *repository.SavedPostRepo
}

// NewSavedPostUsecase creates a new SavedPostUsecase
func NewSavedPostUsecase(spr *repository.SavedPostRepo) *SavedPostUsecase {
	return &SavedPostUsecase{SavedPostRepo: spr}
}

// Create creates a new saved post
func (spu *SavedPostUsecase) Create(savedPost *entity.SavedPost) error {
	err := spu.SavedPostRepo.Create(savedPost)
	if err != nil {
		return err
	}
	return nil
}

// GetAllByUser retrieves all saved posts for a user
func (spu *SavedPostUsecase) GetAllByUser(userID string) ([]model.SavedPost, error) {
	savedPosts, err := spu.SavedPostRepo.FindAllByUser(userID)
	if err != nil {
		return nil, err
	}

	var savedPostsModel []model.SavedPost
	for _, sp := range savedPosts {
		savedPostsModel = append(savedPostsModel, model.SavedPost{
			UserID: sp.UserID.String(),
			PostID: sp.PostID,
		})
	}

	return savedPostsModel, nil
}
