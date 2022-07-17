package Model

import "time"

type StudentAnswer struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	StudentId  int32     `gorm:"notNull" json:"student_id"`
	QuizId     int32     `gorm:"notNull" json:"quiz_id"`
	QuestionId int32     `gorm:"notNull" json:"question_id"`
	AnswerKey  string    `gorm:"size:255;null" json:"answer_key"`
	AnsweredAt time.Time `gorm:"default:current_timestamp;notNull" json:"answered_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
