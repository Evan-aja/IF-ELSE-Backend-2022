package Model

import "time"

type Permission struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	StudentId      int32     `gorm:"notNull" json:"student_id"`
	AgendaId       int32     `gorm:"notNull" json:"agenda_id"`
	Status         bool      `gorm:"default:0;notNull" json:"status"`
	PermissionType int32     `gorm:"notNull" json:"permission_type"`
	Evidence       string    `gorm:"size:255;notNull" json:"evidence"`
	Details        string    `gorm:"type:text" json:"details"`
	CreatedAt      time.Time `gorm:"null" json:"created_at"`
	UpdatedAt      time.Time `gorm:"type:timestamp" json:"updated_at"`
}
