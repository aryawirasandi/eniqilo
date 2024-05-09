package main

import (
	"fmt"
	"net/http"

	"github.com/aryawirasandi/eniqilo/src/drivers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/about", func(c echo.Context) error {
		return c.String(http.StatusOK, "About")
	})
	e.Logger.Fatal(e.Start(":1324"))
}
