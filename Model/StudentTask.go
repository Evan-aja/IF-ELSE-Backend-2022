package Model

type StudentTask struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	TaskID    uint    `gorm:"notNull" json:"task_id"`
	Task      Task    `gorm:"foreignKey:TaskID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	LinkID    uint    `gorm:"notNull" json:"link_id"`
	Links     Links   `gorm:"foreignKey:LinkID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	StudentID uint    `gorm:"notNull" json:"student_id"`
	Student   Student `gorm:"foreignKey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// Submission  string    `gorm:"type:text;null" json:"submission"`
	// SubmittedAt time.Time `gorm:"default:current_timestamp;null" json:"submitted_at"`
}
