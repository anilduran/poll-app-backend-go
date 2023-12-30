package models

import "time"

type Vote struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UserID    uint      `json:"user_id" gorm:"index:idx_userid_pollid"`
	OptionID  uint      `json:"option_id" gorm:"index:idx_userid_pollid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
