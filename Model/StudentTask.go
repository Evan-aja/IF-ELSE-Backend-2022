package Model

import (
	"time"
)

type StudentTask struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	TaskId      uint      `gorm:"notNull" json:"task_id"`
	StudentId   uint      `gorm:"notNull" json:"student_id_id"`
	Submission  string    `gorm:"type:text;notNull" json:"submission"`
	SubmittedAt time.Time `gorm:"default:current_timestamp;notNull" json:"submitted_at"`
}
