package Model

import "time"

type StudentPasswordReset struct {
	Token     string    `gorm:"size:255;notNull" json:"token"`
	CreatedAt time.Time `gorm:"null" json:"created_at"`
}
