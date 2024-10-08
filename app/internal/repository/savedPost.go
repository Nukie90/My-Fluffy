package repository

import (
	"github.com/Nukie90/my-fluffy/app/domain/entity"
	"github.com/Nukie90/my-fluffy/app/domain/model"
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

func (spr *SavedPostRepo) GetAllSavedPostsWithUsernames(userID ulid.ULID) ([]model.PostWithUsername, error) {
	var savedPosts []model.PostWithUsername

	result := spr.DB.Table("saved_posts").
		Select("posts.id, posts.title, posts.content, posts.status, posts.picture, posts.reward, posts.owner_id, posts.found_id, users.username").
		Joins("join posts on saved_posts.post_id = posts.id").
		Joins("join users on posts.owner_id = users.id").
		Where("saved_posts.user_id = ? AND saved_posts.deleted_at IS NULL AND posts.deleted_at IS NULL", userID).
		Find(&savedPosts)

	if result.Error != nil {
		return nil, result.Error
	}

	return savedPosts, nil
}

// Unsave removes a saved post
func (spr *SavedPostRepo) Unsave(userID ulid.ULID, postID uint) error {
	result := spr.DB.Where("user_id = ? AND post_id = ?", userID, postID).Delete(&entity.SavedPost{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
