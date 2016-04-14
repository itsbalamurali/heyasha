package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func MessengerBot(c echo.Context) error {
	return c.String(http.StatusOK, "Messenger Bot Response!\n")

}
