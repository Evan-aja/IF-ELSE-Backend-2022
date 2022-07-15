package Model

import "time"

type Quizs struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Label       string    `gorm:"default:Quiz;notNull;size:255" json:"label"`
	StartAt     time.Time `gorm:"notNull" json:"start_at"`
	EndAt       time.Time `gorm:"notNull" json:"end_at"`
	AgendaId    int32     `gorm:"notNull" json:"agenda_id"`
	Total       int32     `gorm:"default:15;notNull" json:"total"`
	DurationSec int32     `gorm:"notNull;default:900;" json:"duration_sec"`
	IsPublished bool      `gorm:"notNull;default:false" json:"is_published"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
