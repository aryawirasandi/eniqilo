package main

import (
	"fmt"

	router "github.com/aryawirasandi/eniqilo/src/app"
	"github.com/aryawirasandi/eniqilo/src/drivers"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	dbConnection, err := drivers.CreateConnection()
	if err != nil {
		fmt.Println("Error creating database connection:", err)
		return
	}

	defer func() {
		if err := dbConnection.Close(); err != nil {
			fmt.Println("Error closing database connection:", err)
		}
	}()

	router.Router()
}
