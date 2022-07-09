package Model

import (
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `json:"nama"`
	Username     string    `gorm:"uniqueIndex;size:20;notNull" json:"username"`
	Password     string    `gorm:"notNull" json:"password"`
	ApiToken     string    `gorm:"null;uniqueIndex" json:"api_token"`
	RoleId       string    `gorm:"default:1" json:"role_id"`
	StudentId    uint      `gorm:"null"`
	PermittedFor string    `gorm:"default:news,tasks,quizs"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
