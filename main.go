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
	"github.com/itsbalamurali/heyasha/middleware"
	"net/http"
	"os"
	"runtime"
	"time"
	"github.com/itsbalamurali/heyasha/core/engine"
	"github.com/itsbalamurali/heyasha/core/database"
)

var logglyToken string = "09af9fc7-1db3-4c39-a452-f923467e3af1"

func init() {
	database.MysqlCon()
	engine.Boot() //Boot and train all intents
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
		gin.SetMode(gin.ReleaseMode)
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
	router.Use(middleware.MysqlConware())
	router.Use(gin.Recovery())
	router.Use(middleware.RequestIdMiddleware())


	//Static files
	router.Static("/assets", "./assets")

	//Favicon
	router.StaticFile("/favicon.ico", "./assets/img/favicon.ico")

	//Hello!!
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"version": "1.0", "message": "Hello, I'm listening!"})
	})

	/*
	admin := router.Group("/admin")
	{
		admin.GET("/")
	}*/

	//User API No Auth
	router.POST("/v1/users/", controllers.CreateUser)
	router.POST("/v1/users/login", controllers.LoginUser)
	router.POST("/v1/users/reset_password", controllers.ResetPassword)

	auth := router.Group("/",middleware.TokenAuthMiddleware())
	{
		//Core REST API routes
		auth.POST("/v1/stt", controllers.SpeechProcess)    //speech recognition
		auth.GET("/v1/text", controllers.Chat)              //chat with bot
		auth.GET("/v1/intent", controllers.IntentExtract)  //Extract Intent from Text
		auth.GET("/v1/suggest", controllers.SuggestQueries) //Autocomplete user queries
		auth.GET("/v1/tts", controllers.SuggestQueries)    //Text to speech

		//User REST API routes
		auth.POST("/v1/users/logout", controllers.LoginUser)
		auth.GET("/v1/users/me", controllers.GetUserDetails)
		auth.GET("/v1/users/profile/:userId", controllers.GetUserDetails)
		auth.PUT("/v1/users/:userId", controllers.UpdateUserDetails)
		auth.DELETE("/v1/users/:userId", controllers.DeleteUser)

		//Files Upload
		auth.POST("/v1/files/upload", controllers.FileUpload)
		auth.GET("/v1/files/:uuid", controllers.FileGetById)

		//TODO Sync Adapters
		auth.POST("/sync/contacts", controllers.SyncContacts)
		auth.POST("/sync/calender", controllers.SyncCalender)
		auth.POST("/sync/notes", controllers.SyncNotes)
	}

	//Communication Platforms
	router.POST("/chat/slack", platforms.SlackBot)              //SlackBot
	router.POST("/chat/kik", platforms.KikBot)                  //Kik Bot
	router.POST("/chat/telegram", platforms.TelegramBot)        //Telegram Bot
	router.POST("/chat/skype", platforms.SkypeBot)              //Skype Bot
	router.POST("/chat/messenger", platforms.MessengerBotChat)  //Messenger Bot
	router.GET("/chat/messenger", platforms.MessengerBotVerify) //Facebook Callback Verification
	router.POST("/chat/sms", platforms.SmsBot)                  //Sms Bot
	router.POST("/chat/email", platforms.EmailBot)              //Email Bot

	//Method not allowed
	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"status":"error","message": "method not allowed"})
	})

	//404 Handler
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"status":"error","message": "method not found"})
	})

	//Start server
	log.Infoln("Hi, I am running on port: " + port + " !!")
	manners.ListenAndServe(":" + port,router) //Graceful restarts
}