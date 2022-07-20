package Model

import "time"

type Pendataan struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	StudentID int32     `gorm:"notNull" json:"student_id"`
	AgendaID  int32     `gorm:"notNull" json:"agenda_id"`
	Reason    int32     `gorm:"notNull" json:"reason"`
	Details   string    `gorm:"type:text;notNull" json:"details"`
	Evidence1 string    `gorm:"type:text;notNull" json:"evidence_1"`
	Evidence2 string    `gorm:"type:text;notNull" json:"evidence_2"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
