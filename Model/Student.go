package Model

type Student struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	NIM      string `gorm:"uniqueIndex;size:15;notNull" json:"nim"`
	Name     string `json:"nama"`
	NickName string `gorm:"default:Sobat IF ELSE" json:"nickname"`
	Address  string `json:"address"`
	UserId   string `json:"user_id"`
	GroupId  string `gorm:"null" json:"group_id"`
	WhatsApp string `gorm:"null" json:"whatsapp"`
	Line     string `gorm:"null" json:"line"`
	Avatar   string `gorm:"null" json:"avatar"`
}
