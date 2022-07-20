package Model

import (
	"gorm.io/datatypes"
	"time"
)

type StudentTask struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	TaskId      int32          `gorm:"notNull" json:"task_id"`
	StudentId   int32          `gorm:"notNull" json:"student_id_id"`
	Submission  datatypes.JSON `gorm:"notNull" json:"submission"`
	SubmittedAt time.Time      `gorm:"default:current_timestamp;notNull" json:"submitted_at"`
}
