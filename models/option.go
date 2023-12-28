package models

import "time"

type Option struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Text      string    `json:"text"`
	PollID    uint      `json:"poll_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
