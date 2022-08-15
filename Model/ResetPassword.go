package Model

import "time"

type PasswordReset struct {
	Email     string    `gorm:"notNull;index;size:255" json:"email"`
	Token     string    `gorm:"notNull;size:255" json:"token"`
	CreatedAt time.Time `gorm:"null" json:"created_at"`
}
