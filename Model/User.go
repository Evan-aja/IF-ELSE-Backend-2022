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

type ResetPassword struct {
	Email     string    `gorm:"" json:"email"`
	Token     string    `json:"token"`
	CreatedAt time.Time `gorm:"null" json:"created_at"`
}

type Student struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	NIM      string `gorm:"uniqueIndex;size:15;notNull" json:"nim"`
	Name     string `json:"nama"`
	NickName string `gorm:"default:Sobat IF ELSE" json:"nickname"`
	Address  string `json:"address"`
	UserId   string `json:"user_id"`
	GroupId  string `gorm:"null" json:"group_id"`
	WhatsApp string `gorm:"null" json:"whatsapp"`
	Line     string `gorm:"null" json:"line"`
	Avatar   string `gorm:"null" json:"avatar"`
}
