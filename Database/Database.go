package Database

import (
	"fmt"
	"ifelse/Model"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Open() *gorm.DB {
	var db *gorm.DB
	var err error

	// Buka Koneksi
	db, err = gorm.Open(
		mysql.Open(
			fmt.Sprintf(
				"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True",
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASS"),
				os.Getenv("DB_HOST"),
				os.Getenv("DB_NAME"),
			),
		),
		&gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	// Model
	if err = db.AutoMigrate(
		&Model.User{},
		&Model.Group{}, 
		&Model.Student{}, 
		&Model.Agenda{}, 
		&Model.Attendance{}, 
		&Model.Graduation{}, 
		&Model.GroupCompanions{}, 
		&Model.Marking{}, 
		&Model.News{}, 
		&Model.Pendataan{}, 
		&Model.Perizinan{}, 
		&Model.Permission{}, 
		&Model.QuizQuestion{}, 
		&Model.Quizs{}, 
		&Model.PasswordReset{}, 
		&Model.StudentAnswer{}, 
		&Model.StudentAttendance{}, 
		&Model.StudentPerizinan{}, 
		&Model.StudentQuiz{}, 
		&Model.StudentTask{}, 
		&Model.Task{});
		err != nil {
		log.Fatal(err.Error())
	}
	return db
}
