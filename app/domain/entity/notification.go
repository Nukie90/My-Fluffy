package entity

import (
	"github.com/oklog/ulid/v2"
	"time"
)

type Notification struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    ulid.ULID `json:"user_id" gorm:"primaryKey"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	DeleteAt  time.Time `json:"delete_at"`
}

func (Notification) TableName() string {
	return "notifications"
}
