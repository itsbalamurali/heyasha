package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/itsbalamurali/bot/controllers"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello!!")

	// Echo instance
	e := echo.New()
	port := os.Getenv("PORT")
	if port == ""{
		port = "80"
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.Get("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	//speech recognition
	e.Post("/api/speech", func(c echo.Context) error {
		//c.File()
		return c.String(http.StatusOK,"Audio Stream")
	})

	//general purpose http api
	e.Get("/api/query", func(c echo.Context) error {
		return c.String(http.StatusOK, "Query API")
	})

	e.Get("/api/extract", func(c echo.Context) error {
		return c.String(http.StatusOK, "Extract Structured Intent in Text")
	})

	e.Post("/chat/slack",controllers.SlackBot) 	    //SlackBot
	e.Post("/chat/kik", controllers.KikBot)             //Kik Bot
	e.Post("/chat/telegram", controllers.TelegramBot)   //Telegram Bot
	e.Post("/chat/skype", controllers.SkypeBot)         //Skype Bot
	e.Post("/chat/messenger", controllers.MessengerBot) //Messenger Bot
	e.Get("/chat/messenger", controllers.VerifyMessengerBot) //Facebook Callback Verification
	e.Post("/chat/sms", controllers.SmsBot)             //Sms Bot
	e.Post("/chat/email", controllers.EmailBot) //Email Bot

	//User routes
	/*
		e.Post("/users/", controllers.CreateUser)
		e.Post("/users/login", controllers.LoginUser)
		e.Get("/users/:id",controllers.GetUserDetails)
		e.Delete("/users/:id", controllers.DeleteUser)
	*/
	// Start server
	e.Run(standard.New(":"+ port))
}
