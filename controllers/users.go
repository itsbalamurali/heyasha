package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func CreateUser(c echo.Context) error {
	return c.String(http.StatusOK,"")
}

func LoginUser(c echo.Context) error {
	return c.String(http.StatusOK,"")
}

func GetUserDetails(c echo.Context) error {
	return c.String(http.StatusOK,"")
}

func DeleteUser(c echo.Context) error {
	return c.String(http.StatusOK,"")
}