package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func EmailBot(c echo.Context) error {
	return c.String(http.StatusOK,"Email Bot Response!!!")
}
