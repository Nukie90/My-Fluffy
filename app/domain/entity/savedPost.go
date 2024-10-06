package entity

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

// SavedPost is the entity model for saved posts
type SavedPost struct {
	ID        ulid.ULID      `json:"id" gorm:"primaryKey"`
	UserID    ulid.ULID      `json:"user_id"`
	PostID    uint           `json:"post_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// TableName returns the table name for the saved post model
func (SavedPost) TableName() string {
	return "saved_posts"
}

// BeforeCreate generates the ULID for the saved post
func (sp *SavedPost) BeforeCreate(tx *gorm.DB) (err error) {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	sp.ID = ulid.MustNew(ulid.Now(), entropy)
	sp.CreatedAt = time.Now()
	sp.UpdatedAt = time.Now()
	return
}
