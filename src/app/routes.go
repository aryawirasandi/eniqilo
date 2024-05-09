package router

import (
	customerController "github.com/aryawirasandi/eniqilo/src/app/controller/user/customer"
	staffController "github.com/aryawirasandi/eniqilo/src/app/controller/user/staff"
	"github.com/labstack/echo/v4"
)

func Router() {
	e := echo.New()
	g := e.Group("/v1")
	staff := g.Group("/staff")
	customer := g.Group("/customer")

	staff.POST("/register", staffController.Register)
	staff.POST("/login", staffController.Login)
	customer.POST("/register", customerController.Register)
	customer.POST("/login", customerController.Login)

	e.Logger.Fatal(e.Start(":1324"))
}
