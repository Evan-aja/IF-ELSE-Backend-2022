package Model

import "time"

type Rangkaian struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	NamaRangkaian string `gorm:"size:255" json:"nama_rangkaian"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt 	  time.Time `json:"updated_at"`
}