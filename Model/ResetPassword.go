package Model

import "time"

type ResetPassword struct {
	Email     string    `gorm:"" json:"email"`
	Token     string    `json:"token"`
	CreatedAt time.Time `gorm:"null" json:"created_at"`
}
