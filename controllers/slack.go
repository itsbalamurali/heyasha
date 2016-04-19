package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func SlackBot(c echo.Context) error {
	return c.String(http.StatusOK,"SlackBot response!!!")
}
