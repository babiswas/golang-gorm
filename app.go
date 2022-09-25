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
	DB.Create_map(db)
}
