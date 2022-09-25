package main

import (
	"Newapp/DB"
	"fmt"
	"os"
)

func handle_error(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	db := DB.Migration()
	p := DB.Person{Name: "Bapa", Email: "bapa@gmail.com"}
	result := db.Select("Name", "Email").Create(&p)
	fmt.Println(p.ID)
	user1 := DB.Person{Name: "Rama", Email: "rama@gmail.com"}
	db.Create(&user1)
	users := []DB.Person{{Name: "Sita", Email: "sita@gmail.com"}, {Name: "Geeta", Email: "geeta@gmail.com"}}
	db.Create(&users)
	fmt.Println(result)
}
