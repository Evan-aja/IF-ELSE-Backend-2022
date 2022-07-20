package Model

import "time"

type Graduation struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	StudentID   int32     `gorm:"notNull" json:"student_id"`
	IsGraduated bool      `gorm:"notNull" json:"is_graduated"`
	Score1      float64   `gorm:"notNull" json:"score_1"`
	Score2      float64   `gorm:"notNull" json:"score_2"`
	Score3      float64   `gorm:"notNull" json:"score_3"`
	Score4      float64   `gorm:"notNull" json:"score_4"`
	CertifUrl   string    `gorm:"size:255;null" json:"certif_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
