package Model

import "time"

type Marking struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	TaskId           int32     `gorm:"notNull" json:"task_id"`
	StudentId        int32     `gorm:"notNull" json:"student_id"`
	Mark             int32     `gorm:"notNull" json:"mark"`
	UpdatedByAdminId int32     `gorm:"notNull" json:"updated_by_admin_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
