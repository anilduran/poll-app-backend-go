package models

import "time"

type User struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Email     string `gorm:"unique" json:"email"`
	Password  string
	Polls     []Poll    `json:"polls"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
