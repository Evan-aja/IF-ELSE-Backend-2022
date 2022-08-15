package Model

type Attendance struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	AgendaId int32  `gorm:"default:1;notNull" json:"agenda_id"`
	Title    string `gorm:"default:'';size:255" json:"title"`
	Code     string `gorm:"notNull;size:255" json:"code"`
	Status   bool   `gorm:"default:false;notNull" json:"status"`
}
