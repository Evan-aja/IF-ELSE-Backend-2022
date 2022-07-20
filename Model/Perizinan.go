package Model

import "time"

type Perizinan struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	AgendaID  int32     `gorm:"notNull" json:"agenda_id"`
	StartedAt time.Time `gorm:"notNull" json:"started_at"`
	EndAt     time.Time `gorm:"notNull" json:"end_at"`
}
