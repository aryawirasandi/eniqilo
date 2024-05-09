package customerController

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	return c.String(http.StatusOK, "Create Customer")
}

func Login(c echo.Context) error {
	return c.String(http.StatusOK, "Login! Customer")
}
