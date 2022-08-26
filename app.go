package main

import (
	"fmt"

	server "example.com/book/src"
	"example.com/book/src/database"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Can't find env file")
	}

	database.Init()
	server.Init()
}
