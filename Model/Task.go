package Model

import "time"

type Task struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:255;notNull" json:"title"`
	Description string    `gorm:"type:longtext;notNull" json:"description"`
	Condition   string    `gorm:"type:text;notNull" json:"condition"`
	Step        string    `gorm:"type:text;notNull" json:"step"`
	JumlahLink  int32     `json:"jumlah_link"`
	Deadline    time.Time `json:"deadline"`
}

type Links struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	TaskID uint   `json:"task_id"`
	Title  string `gorm:"notNull" json:"title"`
	Task   Task   `gorm:"foreignKey:TaskID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Link   string `gorm:"size:255" json:"link"`
}

type NewTask struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Condition   string    `json:"condition"`
	Step        string    `json:"step"`
	JumlahLink  int32     `json:"jumlah_link"`
	Deadline    time.Time `json:"deadline"`
	Links 		[]string  `json:"links"`
}
