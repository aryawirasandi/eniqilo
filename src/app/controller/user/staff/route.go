package staffController

import (
	"net/http"

	dataSource "github.com/aryawirasandi/eniqilo/src/domain/datasource"
	"github.com/aryawirasandi/eniqilo/src/domain/entities"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	u := new(entities.CreateUser)

	if err := c.Bind(u); err != nil {
		return err
	}
	result, _ := dataSource.Create(u)
	return c.JSON(http.StatusCreated, result)

}

func Login(c echo.Context) error {

	return c.String(http.StatusOK, "Login! Staff")

}
