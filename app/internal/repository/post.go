package repository

import (
	"github.com/Nukie90/my-fluffy/app/domain/entity"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type PostRepo struct {
	DB *gorm.DB
}

func NewPostRepo(db *gorm.DB) *PostRepo {
	return &PostRepo{DB: db}
}

func (pr *PostRepo) Create(post *entity.Post) error {
	result := pr.DB.Create(post)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (pr *PostRepo) GetPostFromSpecificUser(userID ulid.ULID) ([]entity.Post, error) {
	var posts []entity.Post
	result := pr.DB.Where("owner_id = ?", userID).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}
