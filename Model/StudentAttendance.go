package Model

import "time"

type StudentAttendance struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	AttendanceId int32     `gorm:"notNull" json:"attendance_id"`
	StudentId    int32     `gorm:"notNull" json:"student_id"`
	AttendedAt   time.Time `gorm:"default:current_timestamp;notNull" json:"attended_at"`
}
