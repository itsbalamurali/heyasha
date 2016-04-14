package controllers

import (
	"net/http"
	"github.com/labstack/echo"
)

func TelegramBot(c echo.Context) error {
	c.Request().Body()
	return c.String(http.StatusOK, "Telegram Bot Response!\n")

}
