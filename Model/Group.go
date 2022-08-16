package Model

type Group struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	GroupName     string    `gorm:"size:255" json:"group_name"`
	LineGroup     string    `gorm:"size:255;default:#" json:"line_group"`
	CompanionName string    `gorm:"size:255" json:"companion_name"`
	IDLine        string    `gorm:"size:255" json:"id_line"`
	LinkFoto      string    `gorm:"size:255" json:"link_foto"`
	Student       []Student `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
