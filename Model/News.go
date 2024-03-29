package Model

import (
	"time"
)

type News struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"notNull;size:255" json:"title"`
	Content     string    `gorm:"type:text;null" json:"content"`
	Image       string    `gorm:"null;size:255" json:"image"`
	IsPublished bool      `gorm:"default:true;notNull" json:"is_published"`
	CreatedAt   time.Time `gorm:"notNull;default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp" json:"updated_at"`
}
