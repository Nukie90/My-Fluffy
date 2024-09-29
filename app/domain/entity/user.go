package entity

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"

	"gorm.io/gorm"
)

// User is the user model
type User struct {
	ID        ulid.ULID      `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
	Role      string         `json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeleteAt  gorm.DeletedAt `json:"delete_at" gorm:"index"`
}

// TableName returns the table name for the user model
func (User) TableName() string {
	return "users"
}

// BeforeCreate is a function to generate ULID before creating a new record
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	ulid := ulid.MustNew(ulid.Now(), entropy)
	u.ID = ulid

	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}

// BeforeUpdate is a function to update the updated_at field before updating a record
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
