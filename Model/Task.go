package Model

import "time"

type Task struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"size:255;null" json:"title"`
	Description string `gorm:"type:text;null" json:"description"`
	IsPublished bool   `gorm:"default:false;notNull" json:"is_published"`
	Fields      string `gorm:"type:longtext;null" json:"fields"`
	AgendaId    int32  `gorm:"default:1;notNull" json:"agenda_id"`
	//Condition, step (cara pengerjaan)
	Condition  string    `gorm:"notNull" json:"condition"`
	Step       string    `json:"step"`
	JumlahLink int32     `json:"jumlah_link"`
	Link1      string    `json:"link_1"`
	Linke2     string    `json:"linke_2"`
	StartAt    time.Time `gorm:"null" json:"start_at"`
	EndAt      time.Time `gorm:"null" json:"end_at"`
}
