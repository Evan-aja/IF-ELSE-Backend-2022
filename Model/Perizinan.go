package Model

import "time"

type Perizinan struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	AgendaID  uint      `json:"agenda_id"`
	Agenda    Agenda    `json:"agenda"`
	StudentID uint      `json:"student_id"`
	Student   Student   `json:"student"`
	LinkSurat string    `gorm:"text;notNull" json:"link_surat"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// student_id nya apa, agenda_id nya apa terus baru upload
