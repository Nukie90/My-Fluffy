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

func (pu *SavedPostUsecase) GetAllSavedPostsByUser(userID string) ([]model.Post, error) {
	useridUlid, err := ulid.Parse(userID)
	if err != nil {
		return nil, err
	}

	// Fetch full post details from repository
	posts, err := pu.SavedPostRepo.GetAllSavedPostsByUser(useridUlid)
	if err != nil {
		return nil, err
	}

	// Convert to model.Post type
	var postsModel []model.Post
	for _, post := range posts {
		postsModel = append(postsModel, model.Post{
			ID:      post.ID,
			Title:   post.Title,
			Content: post.Content,
			Picture: post.Picture,
			Reward:  post.Reward,
			OwnerID: post.OwnerID.String(),
			FoundID: post.FoundID.String(),
		})
	}

	return postsModel, nil
}
