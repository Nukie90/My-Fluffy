package entity

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type Payment struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Amount     float64   `json:"amount"`
	ReceiverID ulid.ULID `json:"receiver_id"`
	UserID     ulid.ULID `json:"owner_id"`
	CreatedAt  time.Time `json:"created_at"`
}

func (Payment) TableName() string {
	return "payments"
}
