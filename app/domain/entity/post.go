package entity

import (
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	Status    string         `json:"status"`
	Picture   []byte         `json:"picture"`
	Reward    float64        `json:"reward" gorm:"default:0"`
	OwnerID   ulid.ULID      `json:"owner_id"`
	FoundID   ulid.ULID      `json:"found_id" gorm:"default:null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (Post) TableName() string {
	return "posts"
}

func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	p.Status = "Missing"
	p.FoundID = ulid.ULID{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	return nil
}
