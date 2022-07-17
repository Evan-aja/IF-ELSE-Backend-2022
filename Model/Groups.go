package Model

type Groups struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"default:unnamed;size:255;notNull" json:"name"`
	LineGroup string `gorm:"default:#;size:255;notNull" json:"line_group"`
}
