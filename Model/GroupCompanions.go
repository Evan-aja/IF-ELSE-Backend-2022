package Model

type GroupCompanions struct {
	ID          uint  `gorm:"primaryKey" json:"id"`
	CompanionID int32 `gorm:"notNull" json:"companion_id"`
	GroupID     int32 `gorm:"notNull" json:"group_id"`
}
