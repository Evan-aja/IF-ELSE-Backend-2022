package Model

import (
	"time"
)

type User struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Name          string    `gorm:"notNull;size:255" json:"name"`
	Username      string    `gorm:"size:255;notNull" json:"username"`
	Email         string    `gorm:"size:255;notNull" json:"email"`
	Password      string    `gorm:"notNull;size:255" json:"password"`
	ApiToken      string    `gorm:"null;size:80;" json:"api_token"`
	RoleId        string    `gorm:"default:1;notNull;size:255" json:"role_id"`
	StudentId     uint      `json:"student_id"`
	Student       Student   `gorm:"ForeignKey:StudentId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PermittedFor  string    `gorm:"default:news,tasks,quizs;size:255;notNull"`
	RememberToken string    `gorm:"size:100"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ChangePassword struct {
	Password string `gorm:"notNull;size:255" json:"password"`
	Newpass1 string `gorm:"notNull;size:255" json:"newpass1"`
	Newpass2 string `gorm:"notNull;size:255" json:"newpass2"`
}
