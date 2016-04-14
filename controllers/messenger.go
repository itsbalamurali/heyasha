package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func VerifyMessengerBot(c echo.Context) error {
	verify_token := "er7Wq4yREXBKpdRKjhAg"
	hub_mode := c.QueryParam("hub.mode")
	hub_challenge := c.QueryParam("hub.challenge")
	hub_verify_token := c.QueryParam("hub.verify_token")
	if hub_verify_token == verify_token && hub_challenge != "" {
		c.Request().Header().Set("Hub Mode",hub_mode)
		return c.String(http.StatusOK,hub_challenge)
	} else {
		return c.String(http.StatusBadRequest, "Something went wrong!\n")
	}
}

func MessengerBot(c echo.Context) error {
	return c.String(http.StatusOK, "MessengerBot Response!\n")

}