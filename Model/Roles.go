package Model

import "time"

type Roles struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"default:student;size:255;notNull" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
