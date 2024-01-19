package models

import (
	"time"

	"github.com/google/uuid"
)

type Vote struct {
	ID        uuid.UUID `gorm:"primary_key" json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	OptionID  uuid.UUID `json:"option_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (vote *Vote) BeforeCreate() (err error) {
	vote.ID = uuid.New()
	return
}
