package Model

import "time"

type Agenda struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	Title            string    `gorm:"size:255;notNull" json:"title"`
	Content          string    `gorm:"type:text;notNull" json:"content"`
	Image            string    `gorm:"size:255;notNull" json:"image"`
	StartAt          string    `gorm:"notNull" json:"start_at"`
	EndAt            string    `gorm:"notNull" json:"end_at"`
	PerizinanStartAt string    `gorm:"notNull" json:"perizinan_start_at"`
	PerizinanEndAt   string    `gorm:"notNull" json:"perizinan_end_at"`
	IsPublished      bool      `gorm:"default:false;notNull" json:"is_published"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Perizinan        []Perizinan `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
