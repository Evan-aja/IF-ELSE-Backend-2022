package Model

import (
	"time"
)

type Agenda struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	Title            string    `gorm:"size:255;notNull" json:"title"`
	Content          string    `gorm:"type:text;notNull" json:"content"`
	Image            string    `gorm:"size:255;notNull" json:"image"`
	StartAt          time.Time `gorm:"notNull" json:"start_at"`
	EndAt            time.Time `gorm:"notNull" json:"end_at"`
	PendataanStartAt time.Time `gorm:"default:current_timestamp;notNull" json:"pendataan_start_at"`
	PendataanEndAt   time.Time `gorm:"default:current_timestamp;notNull" json:"pendataan_end_at"`
	PerizinanStartAt time.Time `gorm:"default:current_timestamp;notNull" json:"perizinan_start_at"`
	PerizinanEndAt   time.Time `gorm:"default:current_timestamp;notNull" json:"perizinan_end_at"`
	IsPublished      bool      `gorm:"default:true;notNull" json:"is_published"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
