package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Poll struct {
	ID         uuid.UUID   `gorm:"type:uuid;primary_key;"`
	Question   string      `json:"question"`
	UserID     uuid.UUID   `json:"user_id"`
	Options    []Option    `json:"options"`
	Categories []*Category `json:"categories" gorm:"many2many:poll_categories;"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

func (poll *Poll) BeforeCreate(tx *gorm.DB) (err error) {
	poll.ID = uuid.New()
	return
}
