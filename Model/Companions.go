package Model

type Companions struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"size:255;notNull" json:"name"`
	Line  string `gorm:"default:'';size:255;null" json:"line"`
	Phone string `gorm:"default:'';size:255;null" json:"phone"`
}
