package business

import (
	"github.com/Nukie90/my-fluffy/app/domain/entity"
	"github.com/Nukie90/my-fluffy/app/domain/model"
	"github.com/Nukie90/my-fluffy/app/internal/repository"
	"github.com/oklog/ulid/v2"
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

func (spu *SavedPostUsecase) GetAllSavedPostsByUserWithDetails(userID string) ([]model.PostWithUsername, error) {
	useridUlid, err := ulid.Parse(userID)
	if err != nil {
		return nil, err
	}

	// Fetch saved posts with usernames from the repository
	savedPosts, err := spu.SavedPostRepo.GetAllSavedPostsWithUsernames(useridUlid)
	if err != nil {
		return nil, err
	}

	return savedPosts, nil
}

// Unsave removes a saved post
func (spu *SavedPostUsecase) Unsave(userID string, postID uint) error {
	useridUlid, err := ulid.Parse(userID)
	if err != nil {
		return err
	}

	// Call the repository to unsave the post
	err = spu.SavedPostRepo.Unsave(useridUlid, postID)
	if err != nil {
		return err
	}

	return nil
}
