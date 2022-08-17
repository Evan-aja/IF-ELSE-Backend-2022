package Model

import "time"

type StudentTask struct {
	ID      uint `gorm:"primaryKey" json:"id"`
	LinksID uint `gorm:"notNull" json:"links_id"`
	// Task        Task      `gorm:"foreignKey:TaskID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Links       Links     
	StudentID   uint      `gorm:"notNull" json:"student_id"`
	Student     Student   `gorm:"foreignKey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Submission  string    `gorm:"type:text;null" json:"submission"`
	SubmittedAt time.Time `gorm:"default:current_timestamp;null" json:"submitted_at"`
}
