package DB

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Get_db() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=36network dbname=new port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error Occured", err)
		return db, err
	}
	return db, nil
}

func handle_err(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type Person struct {
	gorm.Model
	Name  string
	Email string
}

func Migration() *gorm.DB {
	db, err := Get_db()
	handle_err(err)
	db.AutoMigrate(&Person{})
	return db
}
