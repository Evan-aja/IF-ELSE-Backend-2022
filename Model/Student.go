package Model

type Student struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	GroupID  uint   `gorm:"default:1" json:"group_id"`
	Group    Group
	NIM      string `gorm:"uniqueIndex;size:15;notNull" json:"nim"`
	Name     string `gorm:"notNull;size:255" json:"nama"`
	Nickname string `gorm:"default:Sobat IF ELSE;notNull;size:255" json:"nickname"`
	Address  string `gorm:"size:255;notNull" json:"address"`
	Whatsapp string `gorm:"null;size:255" json:"whatsapp"`
	Line     string `gorm:"null;size:255" json:"line"`
	Avatar   string `gorm:"default:https://i.imgur.com/5HAGlzV.png;size:255" json:"avatar"`
	About    string `gorm:"size:255;null" json:"about"`
	StudentPerizinan []StudentPerizinan
	StudentTask []StudentTask
	StudentMarking []StudentMarking
}
