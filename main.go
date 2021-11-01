package main

import (
	"GoCrud/pkg/db"
	"fmt"
)

func main() {
	db.ConnectDB()

	fmt.Println("Hello World!")
}
