package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func SmsBot(c echo.Context) error {
	return c.String(http.StatusOK, "Sms Bot Response!\n")

}
