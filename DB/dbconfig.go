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

func test_task(db *gorm.DB) {
	p := Person{Name: "Bapa", Email: "bapa@gmail.com"}
	result := db.Select("Name", "Email").Create(&p)
	fmt.Println(p.ID)
	user1 := Person{Name: "Rama", Email: "rama@gmail.com"}
	db.Create(&user1)
	users := []Person{{Name: "Sita", Email: "sita@gmail.com"}, {Name: "Geeta", Email: "geeta@gmail.com"}}
	db.Create(&users)
	fmt.Println(result)
}

func Create_map(db *gorm.DB) {
	db.Model(&Person{}).Create([]map[string]interface{}{{"Name": "Ravan", "Email": "ravan@gmail.com"}, {"Name": "Mintu", "Email": "mintu@gmail.com"}})
}
