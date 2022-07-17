package Model

import "time"

type StudentQuiz struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	StudentID uint      `gorm:"notNull" json:"student_id"`
	QuizID    uint      `gorm:"notNull" json:"quiz_id"`
	StartedAt time.Time `gorm:"notNull" json:"started_at"`
	EndedAt   time.Time `gorm:"null" json:"ended_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
