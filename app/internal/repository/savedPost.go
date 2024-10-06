package repository

import (
	"github.com/Nukie90/my-fluffy/app/domain/entity"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

// SavedPostRepo is the repository for saved posts
type SavedPostRepo struct {
	DB *gorm.DB
}

// NewSavedPostRepo creates a new SavedPostRepo
func NewSavedPostRepo(db *gorm.DB) *SavedPostRepo {
	return &SavedPostRepo{DB: db}
}

// Create creates a new saved post
func (spr *SavedPostRepo) Create(savedPost *entity.SavedPost) error {
	result := spr.DB.Create(savedPost)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (pr *SavedPostRepo) GetAllSavedPostsByUser(userID ulid.ULID) ([]entity.Post, error) {
	var posts []entity.Post
	// Ensure only posts saved by the specific userID are retrieved
	result := pr.DB.Table("saved_posts").
		Select("posts.*").
		Joins("join posts on saved_posts.post_id = posts.id").
		Where("saved_posts.user_id = ?", userID).
		Find(&posts)

	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}
