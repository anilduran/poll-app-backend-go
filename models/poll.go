package models

import "time"

type Poll struct {
	ID         uint        `json:"id" gorm:"primary_key"`
	Question   string      `json:"question"`
	UserID     uint        `json:"user_id"`
	Options    []Option    `json:"options"`
	Categories []*Category `json:"categories" gorm:"many2many:poll_categories;"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}
