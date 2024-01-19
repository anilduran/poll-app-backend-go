package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Email     string    `gorm:"uniqueIndex" json:"email"`
	Password  string
	Polls     []Poll    `json:"polls"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) BeforeCreate() (err error) {
	user.ID = uuid.New()
	return
}
