//Project : Hey Asha!
//Copyright (C) Balamurali Pandranki - All Rights Reserved
//Unauthorized copying of this file, via any medium is strictly prohibited
//Proprietary and confidential
//Written by Balamurali Pandranki <balamurali@live.com>, 25/4/2016 4:27 PM
package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/braintree/manners"
	"github.com/itsbalamurali/heyasha/controllers"
	"github.com/itsbalamurali/heyasha/controllers/platforms"
	"github.com/itsbalamurali/heyasha/core/database"
	"github.com/itsbalamurali/heyasha/middleware"
	"net/http"
	"os"
	"runtime"
	"time"
)

var logglyToken string = "09af9fc7-1db3-4c39-a452-f923467e3af1"

func init() {
	database.Connect()
}

func main() {
	// maximize CPU usage for maximum performance
	runtime.GOMAXPROCS(runtime.NumCPU())
	//gin.SetMode(gin.ReleaseMode)

	//Init Logger
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stderr)

	app_env := os.Getenv("ENV")
	if app_env == "production" {
		log.SetLevel(log.WarnLevel)
	} else {
		log.SetLevel(log.DebugLevel)
	}

	log.Infoln("Starting server...")

	//Port to Bind server to
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	//New Router
	router := gin.New()
	// Global middleware
	router.Use(middleware.Ginrus(log.StandardLogger(),time.RFC3339,true))
	router.Use(middleware.Connect())
	router.Use(gin.Recovery())
	router.Use(middleware.RequestIdMiddleware())

	//Favicon
	router.StaticFile("/favicon.ico", "./assets/img/favicon.ico")

	//Hello!!
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"version": "1.0", "message": "Hello, I'm listening!"})
	})

	v1 := router.Group("/v1", middleware.TokenAuthMiddleware())
	{
		//Core REST API routes
		v1.POST("/list", controllers.SpeechProcess)    //speech recognition
		v1.GET("/chat", controllers.Chat)              //chat with bot
		v1.GET("/extract", controllers.IntentExtract)  //Extract Intent from Text
		v1.GET("/suggest", controllers.SuggestQueries) //Autocomplete user queries
		v1.GET("/talk", controllers.SuggestQueries)    //Autocomplete user queries

		//Communication Platforms
		v1.POST("/chat/slack", platforms.SlackBot)              //SlackBot
		v1.POST("/chat/kik", platforms.KikBot)                  //Kik Bot
		v1.POST("/chat/telegram", platforms.TelegramBot)        //Telegram Bot
		v1.POST("/chat/skype", platforms.SkypeBot)              //Skype Bot
		v1.POST("/chat/messenger", platforms.MessengerBotChat)  //Messenger Bot
		v1.GET("/chat/messenger", platforms.MessengerBotVerify) //Facebook Callback Verification
		v1.POST("/chat/sms", platforms.SmsBot)                  //Sms Bot
		v1.POST("/chat/email", platforms.EmailBot)              //Email Bot

		//User REST API routes
		v1.POST("/users/", controllers.CreateUser)
		v1.POST("/users/login", controllers.LoginUser)
		v1.POST("/users/logout", controllers.LoginUser)
		v1.GET("/users/{UserId}", controllers.GetUserDetails)
		v1.GET("/users/me", controllers.GetUserDetails)
		v1.PUT("/users/{UserId}", controllers.UpdateUserDetails)
		v1.DELETE("/users/{UserId}", controllers.DeleteUser)
		v1.DELETE("/users/reset_password", controllers.ResetPassword)
	}

	//TODO Sessions & Files
	/*
		//Sync Adapters
		router.POST("/sync/contacts")
		router.POST("/sync/calender")
		router.POST("/sync/notes")
	*/

	//Method not allowed
	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"message": "method not allowed"})
	})

	//404 Handler
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "method not found"})
	})

	//Start server
	log.Infoln("Hi, I am running on port: " + port + " !!")
	manners.ListenAndServe(":" + port,router) //Graceful restarts
}
