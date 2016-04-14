package main

import (
	"fmt"
	"github.com/itsbalamurali/bot/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	fmt.Println("Hello!!")

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.Get("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	e.Post("/chat/kik", controllers.KikBot)             //Kik Bot
	e.Post("/chat/telegram", controllers.TelegramBot)   //Telegram Bot
	e.Post("/chat/skype", controllers.SkypeBot)         //Skype Bot
	e.Post("/chat/messenger", controllers.MessengerBot) //Messenger Bot
	e.Post("/chat/sms", controllers.SmsBot)             //Sms Bot

	//User routes
	/*
		e.Post("/users/", controllers.CreateUser)
		e.Post("/users/login", controllers.LoginUser)
		e.Get("/users/:id",controllers.GetUserDetails)
		e.Delete("/users/:id", controllers.DeleteUser)
	*/
	// Start server
	e.Run(standard.New(":80"))
}
