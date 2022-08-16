package Model

import "time"

type StudentPerizinan struct {
	ID          uint `gorm:"primaryKey" json:"id"`
	AgendaID    uint
	PerizinanID uint
	StudentID   uint
	LinkSurat   string    `gorm:"text;notNull" json:"link_surat"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
