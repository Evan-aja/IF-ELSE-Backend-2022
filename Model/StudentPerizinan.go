package Model

import "time"

type StudentPerizinan struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	AgendaID    int32     `gorm:"notNull" json:"agenda_id"`
	PerizinanID int32     `gorm:"notNull" json:"perizinan_id"`
	Residence   string    `gorm:"size:255;notNull" json:"residence"`
	CameraType  int32     `gorm:"notNull" json:"camera_type"`
	Details     string    `gorm:"type:text;notNull" json:"details"`
	Evidance    string    `gorm:"size:255;notNull" json:"evidance"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
