package Model

import "time"

type Perizinan struct {
	ID          uint `gorm:"primaryKey" json:"id"`
	AgendaID    uint
	Agenda Agenda
	StudentID   uint
	Student Student
	LinkSurat   string    `gorm:"text;notNull" json:"link_surat"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// student_id nya apa, agenda_id nya apa terus baru upload