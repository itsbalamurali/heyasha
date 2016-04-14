package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func SkypeBot(c echo.Context) error {
	return c.String(http.StatusOK, "Skype Bot Response!\n")

}
