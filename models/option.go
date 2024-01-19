package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Option struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Text      string    `json:"text"`
	PollID    uuid.UUID `json:"poll_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (option *Option) BeforeCreate(tx *gorm.DB) (err error) {
	option.ID = uuid.New()
	return
}
