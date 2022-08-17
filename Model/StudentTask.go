package Model

import (
	"time"
)

type StudentTask struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	TaskID      uint      `gorm:"notNull" json:"task_id"`
	Task        Task      `gorm:"foreignKey:TaskID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	StudentID   uint      `gorm:"notNull" json:"student_id"`
	Student     Student   `gorm:"foreignKey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Submission  []string  `gorm:"type:text;notNull" json:"submission"`
	SubmittedAt time.Time `gorm:"default:current_timestamp;notNull" json:"submitted_at"`
}
