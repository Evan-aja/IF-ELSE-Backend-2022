package Model

import "time"

type Questions struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	MCEQuestion string    `gorm:"size:255;notNull" json:"mce_question"`
	Choices     string    `gorm:"type:text;notNull" json:"choices"`
	AnswerKey   string    `gorm:"size:255;notNull" json:"answer_key"`
	CreatedAt   time.Time `gorm:"default:current_timestamp;notNull" json:"created_at"`
}
