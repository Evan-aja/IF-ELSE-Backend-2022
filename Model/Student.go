package Model

type Student struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"notNull" json:"name"`
	GroupID     uint   `gorm:"default:1;null" json:"group_id"`
	GroupName   string `json:"group_name"`
	NIM         string `gorm:"uniqueIndex;size:15;notNull" json:"nim"`
	Nickname    string `gorm:"default:Sobat IF ELSE;notNull;size:255" json:"nickname"`
	Address     string `gorm:"default:-;notNull" json:"address"`
	Whatsapp    string `gorm:"default:-;size:255" json:"whatsapp"`
	Line        string `gorm:"default:-;size:255" json:"line"`
	Avatar      string `gorm:"default:https://i.imgur.com/5HAGlzV.png;size:255" json:"avatar"`
	About       string `gorm:"default:-;size:255" json:"about"`
	Perizinan   []Perizinan `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Marking     []Marking `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StudentTask []StudentTask `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
