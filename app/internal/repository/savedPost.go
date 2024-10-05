package repository

import (
	"fmt"

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

// FindAll finds all saved posts for a specific user
func (spr *SavedPostRepo) FindAllByUser(userID string) ([]entity.SavedPost, error) {
	var savedPosts []entity.SavedPost

	// Convert userID string to ulid.ULID
	userIDUlid, err := ulid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user_id format: %w", err)
	}

	fmt.Printf("Finding saved posts for userID: %s\n", userID)

	result := spr.DB.Where("user_id = ?", userIDUlid).Find(&savedPosts)
	if result.Error != nil {
		return nil, result.Error
	}

	fmt.Printf("Rows affected: %d\n", result.RowsAffected)

	return savedPosts, nil
}
