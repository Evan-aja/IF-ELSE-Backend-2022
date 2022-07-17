package Model

type QuizQuestion struct {
	ID         uint  `gorm:"primaryKey" json:"id"`
	QuizId     int32 `gorm:"notNull" json:"quiz_id"`
	QuestionId int32 `gorm:"notNull" json:"question_id"`
}
