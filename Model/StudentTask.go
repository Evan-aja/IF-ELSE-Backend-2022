package Model

import (
	"time"
)

type StudentTask struct {
	ID             uint `gorm:"primaryKey" json:"id"`
	TaskID		   uint
	StudentID      uint `json:"student_id"`
	SubmissionLink string `json:"submission_link"`
	CreatedAt      time.Time 
	SubmittedAt    time.Time `json:"submitted_at"`
	Status         bool `json:"status"`
}
