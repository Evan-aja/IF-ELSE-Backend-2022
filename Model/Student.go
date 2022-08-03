package Model

type Student struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	NIM      string `gorm:"uniqueIndex;size:15;notNull" json:"nim"`
	Name     string `gorm:"notNull;size:255" json:"nama"`
	Nickname string `gorm:"default:Sobat IF ELSE;notNull;size:255" json:"nickname"`
	Address  string `gorm:"size:255;notNull" json:"address"`
	GroupId  string `gorm:"null;size:255" json:"group_id"`
	Whatsapp string `gorm:"null;size:255" json:"whatsapp"`
	Line     string `gorm:"null;size:255" json:"line"`
	Avatar   string `gorm:"null;size:255" json:"avatar"`
	About    string `gorm:"size:255;null" json:"about"`
}
