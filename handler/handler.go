package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")
		return c.String(http.StatusOK, "Hello World "+ username)
	}
}
