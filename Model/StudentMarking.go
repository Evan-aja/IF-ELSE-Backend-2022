package Model

import "time"

type StudentMarking struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	MarkingID uint      `json:"marking_id"`
	StudentID uint      `json:"student_id"`
	Mark      int32     `json:"mark"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
