package Model

import "time"

type Marking struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	AgendaID  uint      `json:"agenda_id"`
	Agenda    Agenda    `gorm:"foreignKey:AgendaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	StudentID uint      `json:"student_id"`
	// Student   Student   `gorm:"foreignKey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Mark      int32     `gorm:"notNull" json:"mark"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
