package Model

import (
	"time"
)

type User struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Name          string    `gorm:"notNull;size:255" json:"name"`
	Username      string    `gorm:"size:255;notNull" json:"username"`
	Password      string    `gorm:"notNull;size:255" json:"password"`
	ApiToken      string    `gorm:"null;uniqueIndex;size:80;" json:"api_token"`
	RoleId        string    `gorm:"default:1;notNull;size:255" json:"role_id"`
	StudentId     uint32    `gorm:"null"`
	PermittedFor  string    `gorm:"default:news,tasks,quizs;size:255;notNull"`
	RememberToken string    `gorm:"size:100"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
