package controllers

import (
	"net/http"
	"github.com/labstack/echo"
)

func KikBot(c echo.Context) error {
	return c.String(http.StatusOK, "Kik Bot Response!\n")

}
