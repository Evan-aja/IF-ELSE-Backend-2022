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
		&Model.Graduation{},
		&Model.Marking{},
		&Model.News{},
		&Model.Perizinan{},
		&Model.StudentTask{},
		&Model.Task{},
		&Model.Links{}); err != nil {
		log.Fatal(err.Error())
	}

	group := Model.Group{
        GroupName: "Belum diatur",
        LineGroup: "Belum diatur",
        CompanionName: "Belum diatur",
        IDLine: "Belum diatur",
        LinkFoto: "Belum diatur",
    }
	
	if result := db.Create(&group); result.Error != nil {
		fmt.Println(result.Error.Error())
		fmt.Println("this is expected, server will run normally")
	}

	return db
}
