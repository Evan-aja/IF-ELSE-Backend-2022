package Model

type GroupCompanions struct {
	ID          uint `gorm:"primaryKey" json:"id"`
	CompanionID uint `gorm:"notNull" json:"companion_id"`
	GroupID     uint `gorm:"notNull" json:"group_id"`
}
