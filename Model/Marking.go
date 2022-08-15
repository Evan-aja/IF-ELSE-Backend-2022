package Model

type Marking struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	RangkaianID      uint
	Rangkaian 		 Rangkaian
	StudentMarking  []StudentMarking
}
