package Model

import "time"

type Penugasan struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"notNull;size:255" json:"title"`
	Description string `gorm:"type:longtext;notNull" json:"description"`
	Condition   string `gorm:"type:text;notNull" json:"condition"`
	Step        string `gorm:"type:text;notNull" json:"step"`
	JumlahLink  int32  `json:"jumlah_link"`
	//Link        []string `gorm:"type:text" json:"link"`
	Links []Link `json:"links"`
	//Link1    string    `json:"link_1"`
	//Link2    string    `json:"link_2"`
	Fields   string    `json:"fields"`
	Deadline time.Time `json:"deadline"`
}

//
type Link struct {
	PenugasanID uint
	Link1       string `json:"link_1"`
	Link2       string `json:"link_2"`
}
