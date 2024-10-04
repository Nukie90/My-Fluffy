package entity

import (
	"github.com/oklog/ulid/v2"
	"time"
)

type Payment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Amount    float64   `json:"amount"`
	Transaction string  `json:"transaction"`
	UserID   ulid.ULID `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (Payment) TableName() string {
	return "payments"
}